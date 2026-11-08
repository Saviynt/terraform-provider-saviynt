// Copyright (c) 2025 Saviynt Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"sort"
	"strings"
	"sync"
	"testing"

	openapi "github.com/saviynt/saviynt-api-go-client/connections"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
)

// fakeEIC is a minimal stand-in for the two Saviynt endpoints a REST connection
// uses: create/update (testConnection) and read (getConnectionDetails). It
// stores connection attributes and returns them on read the way the API does.
type fakeEIC struct {
	mu          sync.Mutex
	attrs       map[string]string
	lastRequest map[string]string
}

func (f *fakeEIC) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	f.mu.Lock()
	defer f.mu.Unlock()
	w.Header().Set("Content-Type", "application/json")

	switch r.URL.Path {
	case "/ECM/api/v5/testConnection":
		// The client sends create/update requests as multipart form fields.
		if err := r.ParseMultipartForm(1 << 20); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		f.lastRequest = map[string]string{}
		for k, v := range r.MultipartForm.Value {
			if len(v) > 0 {
				f.lastRequest[k] = v[0]
				f.attrs[k] = v[0]
			}
		}
		_, _ = w.Write([]byte(`{"connectionKey":101,"msg":"Connection Updated Successfully","errorCode":"0"}`))

	case "/ECM/api/v5/getConnectionDetails":
		if _, ok := f.attrs["connectionName"]; !ok {
			w.WriteHeader(http.StatusPreconditionFailed)
			_, _ = w.Write([]byte(`{"errorcode":1,"msg":"Connection does not exist"}`))
			return
		}
		stored, _ := json.Marshal(f.attrs)
		var ca openapi.RESTConnectionAttributes // drops keys that are not connection attributes
		_ = json.Unmarshal(stored, &ca)
		resp := openapi.RESTConnectionResponse{
			Connectionkey:        openapi.PtrInt32(101),
			Connectionname:       openapi.PtrString(f.attrs["connectionName"]),
			Connectiontype:       openapi.PtrString("REST"),
			Errorcode:            openapi.PtrInt32(0),
			Msg:                  openapi.PtrString("Success"),
			Connectionattributes: &ca,
		}
		_ = json.NewEncoder(w).Encode(resp)

	default:
		http.NotFound(w, r)
	}
}

// unknownAttributes reports which attributes of a resource are planned as unknown.
type unknownAttributes struct {
	address string
	got     *[]string
}

func (c unknownAttributes) CheckPlan(_ context.Context, req plancheck.CheckPlanRequest, resp *plancheck.CheckPlanResponse) {
	for _, rc := range req.Plan.ResourceChanges {
		if rc.Address != c.address {
			continue
		}
		unknown, _ := rc.Change.AfterUnknown.(map[string]interface{})
		var names []string
		for k, v := range unknown {
			if b, ok := v.(bool); ok && b {
				names = append(names, k)
			}
		}
		sort.Strings(names)
		*c.got = names
		return
	}
	resp.Error = fmt.Errorf("resource %s not found in plan", c.address)
}

func restConnectionConfig(serverURL, importJSON string) string {
	return fmt.Sprintf(`
provider "saviynt" {
  server_url   = %q
  access_token = "test-token"
}

resource "saviynt_rest_connection_resource" "test" {
  connection_name         = "tf-plan-noise"
  import_account_ent_json = %q
  create_account_json     = jsonencode({ call = [{ name = "call1", url = "https://example.com/users" }] })
}
`, serverURL, importJSON)
}

// TestAccRestConnection_UpdatePlanShowsOnlyChangedAttributes runs a real
// Terraform import, plan and apply against a fake EIC server. The connection
// already exists in EIC with JSON attributes that the configuration does not
// manage. Changing one managed attribute must not plan those attributes, or
// the connection identifiers, as "(known after apply)" (issue #17).
func TestAccRestConnection_UpdatePlanShowsOnlyChangedAttributes(t *testing.T) {
	importV1 := `{"accountParams":{"connection":"acctAuth","processingType":"SequentialAndIterative"}}`
	importV2 := `{"accountParams":{"connection":"acctAuth","processingType":"httpMethod"}}`
	updateAccountJSON := `{"call":[{"name":"update","url":"https://example.com/users/${account.accountID}","httpMethod":"PUT"}]}`
	f := &fakeEIC{attrs: map[string]string{
		"connectionName":       "tf-plan-noise",
		"ImportAccountEntJSON": importV1,
		"CreateAccountJSON":    `{"call":[{"name":"call1","url":"https://example.com/users"}]}`,
		"UpdateAccountJSON":    updateAccountJSON,
		"EnableAccountJSON":    `{"call":[{"name":"enable","url":"https://example.com/users/${account.accountID}/enable"}]}`,
		"DisableAccountJSON":   `{"call":[{"name":"disable","url":"https://example.com/users/${account.accountID}/disable"}]}`,
		"AddAccessJSON":        `{"call":[{"name":"addAccess","url":"https://example.com/groups/${entitlementValue.entitlementID}/members"}]}`,
		"RemoveAccessJSON":     `{"call":[{"name":"removeAccess","url":"https://example.com/groups/${entitlementValue.entitlementID}/members/${account.accountID}"}]}`,
		"RemoveAccountJSON":    `{"call":[{"name":"remove","url":"https://example.com/users/${account.accountID}","httpMethod":"DELETE"}]}`,
	}}
	srv := httptest.NewTLSServer(f)
	defer srv.Close()

	// The provider talks to EIC through http.DefaultClient; trust the test server's certificate.
	prevTransport := http.DefaultTransport
	http.DefaultTransport = srv.Client().Transport
	defer func() { http.DefaultTransport = prevTransport }()

	host := strings.TrimPrefix(srv.URL, "https://")
	const addr = "saviynt_rest_connection_resource.test"
	var unknownOnUpdate []string

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: map[string]func() (tfprotov6.ProviderServer, error){
			"saviynt": providerserver.NewProtocol6WithError(New("test")()),
		},
		Steps: []resource.TestStep{
			{
				Config:             restConnectionConfig(host, importV1),
				ResourceName:       addr,
				ImportState:        true,
				ImportStateId:      "tf-plan-noise",
				ImportStatePersist: true,
			},
			{
				Config: restConnectionConfig(host, importV2),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(addr, plancheck.ResourceActionUpdate),
						unknownAttributes{address: addr, got: &unknownOnUpdate},
						plancheck.ExpectKnownValue(addr, tfjsonpath.New("update_account_json"), knownvalue.StringExact(updateAccountJSON)),
						plancheck.ExpectKnownValue(addr, tfjsonpath.New("connection_key"), knownvalue.Int64Exact(101)),
					},
				},
			},
		},
	})

	// Only the per-operation status fields may be unknown on an update.
	if want := []string{"error_code", "msg"}; strings.Join(unknownOnUpdate, ",") != strings.Join(want, ",") {
		t.Errorf("attributes planned as unknown on update = %v, want %v", unknownOnUpdate, want)
	}
	t.Logf("attributes planned as unknown on update: %d %v", len(unknownOnUpdate), unknownOnUpdate)

	// Attributes configured outside Terraform must be unchanged after the update.
	if got := f.attrs["UpdateAccountJSON"]; got != updateAccountJSON {
		t.Errorf("UpdateAccountJSON changed on the server: %q", got)
	}
	if got := f.attrs["ImportAccountEntJSON"]; got != importV2 {
		t.Errorf("ImportAccountEntJSON was not updated: %q", got)
	}
}
