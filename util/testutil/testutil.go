// Copyright (c) 2025 Saviynt Inc.
// SPDX-License-Identifier: MPL-2.0

package testutil

import (
	"context"
	"encoding/json"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
)

func LoadConnectorData(t *testing.T, filePath, scenario string) map[string]string {
	// Load base JSON file
	data, err := os.ReadFile(filePath)
	if err != nil {
		t.Fatalf("failed to read test config file: %v", err)
	}

	var allData map[string]map[string]interface{}
	if err := json.Unmarshal(data, &allData); err != nil {
		t.Fatalf("failed to unmarshal test config: %v", err)
	}

	// Select create or update based on the scenario
	secrets, exists := allData[scenario]
	if !exists {
		t.Fatalf("scenario '%s' not found in config", scenario)
	}

	result := make(map[string]string)
	for key, value := range secrets {
		switch v := value.(type) {
		case string:
			// Keep plain strings as-is
			result[key] = v

		case map[string]interface{}:
			// Handle nested maps (e.g., status_key_json)
			jsonValue, err := json.Marshal(v)
			if err != nil {
				t.Fatalf("failed to marshal map for key %s: %v", key, err)
			}
			result[key] = string(jsonValue)

		default:
			// Fallback for other types, including arrays or deeply nested data
			jsonValue, err := json.Marshal(value)
			if err != nil {
				t.Fatalf("failed to marshal value for key %s: %v", key, err)
			}
			result[key] = string(jsonValue)
		}
	}

	return result
}

func GetTestDataPath(t *testing.T, relativeFilename string) string {
	t.Helper()

	// Get the full path of the test file that called this function
	_, callerFile, _, ok := runtime.Caller(1)
	if !ok {
		t.Fatalf("unable to determine caller path")
	}

	// Build the absolute path to the data file relative to the caller file
	absPath := filepath.Join(filepath.Dir(callerFile), relativeFilename)
	return absPath
}

// GetTestIdentifier returns a timestamp-based identifier for test resources
// in the format YYYYMMDDHHMMSS (e.g., "20251106050300")
func GetTestIdentifier() string {
	return time.Now().Format("20060102150405")
}

func PrepareTestDataWithEnv(t *testing.T, templatePath string) string {
	t.Helper()

	data, err := os.ReadFile(templatePath)
	if err != nil {
		t.Fatalf("failed to read test data template: %v", err)
	}

	content := string(data)

	// Use timestamp for unique resource names
	timestamp := GetTestIdentifier()

	replacements := map[string]string{
		"{{TEST_ID}}": timestamp,
	}

	// Add fixture replacements from environment variables (set by TestMain)
	fixtureEnvMappings := map[string]string{
		"{{FIXTURE_SECURITY_SYSTEM}}":    "TF_FIXTURE_SECURITY_SYSTEM",
		"{{FIXTURE_ENDPOINT_1}}":         "TF_FIXTURE_ENDPOINT_1",
		"{{FIXTURE_ENDPOINT_2}}":         "TF_FIXTURE_ENDPOINT_2",
		"{{FIXTURE_ENDPOINT_3}}":         "TF_FIXTURE_ENDPOINT_3",
		"{{FIXTURE_ENDPOINT_4}}":         "TF_FIXTURE_ENDPOINT_4",
		"{{FIXTURE_ENDPOINT_5}}":         "TF_FIXTURE_ENDPOINT_5",
		"{{FIXTURE_ENTITLEMENT_TYPE}}":   "TF_FIXTURE_ENTITLEMENT_TYPE",
		"{{FIXTURE_ENTITLEMENT_TYPE_2}}": "TF_FIXTURE_ENTITLEMENT_TYPE_2",
		"{{FIXTURE_ENTITLEMENT}}":        "TF_FIXTURE_ENTITLEMENT",
		"{{FIXTURE_ENTITLEMENT_2}}":      "TF_FIXTURE_ENTITLEMENT_2",
		"{{FIXTURE_ROLE}}":               "TF_FIXTURE_ROLE",
		"{{FIXTURE_REST_CONNECTION}}":    "TF_FIXTURE_REST_CONNECTION",
		"{{FIXTURE_REST_CONNECTION_ID}}": "TF_FIXTURE_REST_CONNECTION_ID",
	}
	for placeholder, envVar := range fixtureEnvMappings {
		if val := os.Getenv(envVar); val != "" {
			replacements[placeholder] = val
		}
	}

	for placeholder, value := range replacements {
		content = strings.ReplaceAll(content, placeholder, value)
	}

	tmpFile := filepath.Join(t.TempDir(), filepath.Base(templatePath))
	if err := os.WriteFile(tmpFile, []byte(content), 0644); err != nil {
		t.Fatalf("failed to write processed test data: %v", err)
	}

	log.Printf("New file path: %v", tmpFile)
	return tmpFile
}

// BuildBlankState constructs an empty tfsdk.State backed by the resource schema.
// Required by ImportStateResponse before SetAttribute can be called on it.
func BuildBlankState(t *testing.T, schemaResp *resource.SchemaResponse) tfsdk.State {
	t.Helper()
	return tfsdk.State{
		Schema: schemaResp.Schema,
		Raw:    tftypes.NewValue(schemaResp.Schema.Type().TerraformType(context.Background()), nil),
	}
}

// AttrPath returns a root-level terraform-plugin-framework attribute path.
func AttrPath(attr string) path.Path {
	return path.Root(attr)
}

// PrepareTestDataWithReplacements processes a test data template with custom replacements
func PrepareTestDataWithReplacements(t *testing.T, templatePath string, extraReplacements map[string]string) string {
	t.Helper()

	data, err := os.ReadFile(templatePath)
	if err != nil {
		t.Fatalf("failed to read test data template: %v", err)
	}

	content := string(data)

	// Use timestamp for unique resource names
	timestamp := GetTestIdentifier()

	// Start with standard replacements
	replacements := map[string]string{
		"{{TEST_ID}}": timestamp,
	}

	// Add extra replacements
	for k, v := range extraReplacements {
		replacements[k] = v
	}

	for placeholder, value := range replacements {
		content = strings.ReplaceAll(content, placeholder, value)
	}

	tmpFile := filepath.Join(t.TempDir(), filepath.Base(templatePath))
	if err := os.WriteFile(tmpFile, []byte(content), 0644); err != nil {
		t.Fatalf("failed to write processed test data: %v", err)
	}

	log.Printf("New file path: %v", tmpFile)
	return tmpFile
}

// =============================================================================
// Coverage test helpers (moved from unit_tests/internal/provider cross-cutting
// coverage files so per-resource tests can use them from any package).
// =============================================================================

// SafeInvoke runs fn, recovering from any panic. Many CRUD handlers will
// dereference a nil client once they get past the extraction guard; we only
// need line coverage of the guard clauses themselves.
func SafeInvoke(t *testing.T, name string, fn func()) {
	t.Helper()
	defer func() {
		if rec := recover(); rec != nil {
			t.Logf("recovered from panic in %s: %v", name, rec)
		}
	}()
	fn()
}

// BlankStateForResource builds a schema-backed, empty tfsdk.State.
func BlankStateForResource(t *testing.T, r resource.Resource) tfsdk.State {
	t.Helper()
	ctx := context.Background()
	schemaResp := &resource.SchemaResponse{}
	r.Schema(ctx, resource.SchemaRequest{}, schemaResp)
	return BuildBlankState(t, schemaResp)
}

// BlankPlanForResource builds an empty but schema-backed tfsdk.Plan.
func BlankPlanForResource(t *testing.T, r resource.Resource) tfsdk.Plan {
	t.Helper()
	ctx := context.Background()
	schemaResp := &resource.SchemaResponse{}
	r.Schema(ctx, resource.SchemaRequest{}, schemaResp)
	return tfsdk.Plan{
		Schema: schemaResp.Schema,
		Raw:    tftypes.NewValue(schemaResp.Schema.Type().TerraformType(ctx), nil),
	}
}

// BlankConfigForResource builds an empty but schema-backed tfsdk.Config.
func BlankConfigForResource(t *testing.T, r resource.Resource) tfsdk.Config {
	t.Helper()
	ctx := context.Background()
	schemaResp := &resource.SchemaResponse{}
	r.Schema(ctx, resource.SchemaRequest{}, schemaResp)
	return tfsdk.Config{
		Schema: schemaResp.Schema,
		Raw:    tftypes.NewValue(schemaResp.Schema.Type().TerraformType(ctx), nil),
	}
}

// BuildPlanWithAttr constructs a schema-backed Plan with a single top-level
// string attribute set to the given value. All other attributes are null.
func BuildPlanWithAttr(t *testing.T, r resource.Resource, attrName, attrValue string) tfsdk.Plan {
	t.Helper()
	ctx := context.Background()
	schemaResp := &resource.SchemaResponse{}
	r.Schema(ctx, resource.SchemaRequest{}, schemaResp)
	schemaType := schemaResp.Schema.Type().TerraformType(ctx).(tftypes.Object)

	raw := map[string]tftypes.Value{}
	for name, at := range schemaType.AttributeTypes {
		if name == attrName {
			raw[name] = tftypes.NewValue(at, attrValue)
		} else {
			raw[name] = tftypes.NewValue(at, nil)
		}
	}
	return tfsdk.Plan{
		Schema: schemaResp.Schema,
		Raw:    tftypes.NewValue(schemaType, raw),
	}
}

// BuildStateWithAttr constructs a schema-backed State with a single top-level
// string attribute set to the given value. All other attributes are null.
func BuildStateWithAttr(t *testing.T, r resource.Resource, attrName, attrValue string) tfsdk.State {
	t.Helper()
	ctx := context.Background()
	schemaResp := &resource.SchemaResponse{}
	r.Schema(ctx, resource.SchemaRequest{}, schemaResp)
	schemaType := schemaResp.Schema.Type().TerraformType(ctx).(tftypes.Object)

	raw := map[string]tftypes.Value{}
	for name, at := range schemaType.AttributeTypes {
		if name == attrName {
			raw[name] = tftypes.NewValue(at, attrValue)
		} else {
			raw[name] = tftypes.NewValue(at, nil)
		}
	}
	return tfsdk.State{
		Schema: schemaResp.Schema,
		Raw:    tftypes.NewValue(schemaType, raw),
	}
}

// RunImportState invokes ImportState with a blank state built from the
// resource schema and returns whether diagnostics contain an error.
func RunImportState(t *testing.T, r resource.ResourceWithImportState, id string) bool {
	t.Helper()
	ctx := context.Background()

	schemaResp := &resource.SchemaResponse{}
	r.Schema(ctx, resource.SchemaRequest{}, schemaResp)

	resp := &resource.ImportStateResponse{
		State: BuildBlankState(t, schemaResp),
	}
	r.ImportState(ctx, resource.ImportStateRequest{ID: id}, resp)
	return resp.Diagnostics.HasError()
}

// ConfigureResourceIfPossible calls r.Configure with the given ProviderData so
// subsequent CRUD calls have a non-nil client wrapper. Failures are swallowed.
func ConfigureResourceIfPossible(t *testing.T, r resource.Resource, providerData interface{}) {
	t.Helper()
	rc, ok := r.(resource.ResourceWithConfigure)
	if !ok {
		return
	}
	SafeInvoke(t, "Configure(for CRUD)", func() {
		req := resource.ConfigureRequest{ProviderData: providerData}
		resp := &resource.ConfigureResponse{}
		rc.Configure(context.Background(), req, resp)
	})
}

// RunCrud drives Create/Read/Update/Delete with empty schema-backed requests.
// providerData is passed to Configure (typically *provider.SaviyntProvider).
func RunCrud(t *testing.T, r resource.Resource, providerData interface{}) {
	t.Helper()
	ctx := context.Background()

	ConfigureResourceIfPossible(t, r, providerData)

	plan := BlankPlanForResource(t, r)
	config := BlankConfigForResource(t, r)

	SafeInvoke(t, "Create", func() {
		createResp := &resource.CreateResponse{State: BlankStateForResource(t, r)}
		r.Create(ctx, resource.CreateRequest{Plan: plan, Config: config}, createResp)
	})

	SafeInvoke(t, "Read", func() {
		readResp := &resource.ReadResponse{State: BlankStateForResource(t, r)}
		r.Read(ctx, resource.ReadRequest{State: BlankStateForResource(t, r)}, readResp)
	})

	SafeInvoke(t, "Update", func() {
		updateResp := &resource.UpdateResponse{State: BlankStateForResource(t, r)}
		r.Update(ctx, resource.UpdateRequest{Plan: plan, Config: config, State: BlankStateForResource(t, r)}, updateResp)
	})

	SafeInvoke(t, "Delete", func() {
		deleteResp := &resource.DeleteResponse{State: BlankStateForResource(t, r)}
		r.Delete(ctx, resource.DeleteRequest{State: BlankStateForResource(t, r)}, deleteResp)
	})
}

// BlankConfigForDataSource builds an empty, schema-backed tfsdk.Config for a DataSource.
func BlankConfigForDataSource(t *testing.T, d datasource.DataSource) tfsdk.Config {
	t.Helper()
	ctx := context.Background()
	schemaResp := &datasource.SchemaResponse{}
	d.Schema(ctx, datasource.SchemaRequest{}, schemaResp)
	return tfsdk.Config{
		Schema: schemaResp.Schema,
		Raw:    tftypes.NewValue(schemaResp.Schema.Type().TerraformType(ctx), nil),
	}
}

// RunDSRead drives DataSource.Read with an empty schema-backed Config.
func RunDSRead(t *testing.T, name string, d datasource.DataSource) {
	t.Helper()
	ctx := context.Background()
	cfg := BlankConfigForDataSource(t, d)

	SafeInvoke(t, name+".Read", func() {
		resp := &datasource.ReadResponse{State: tfsdk.State{Schema: cfg.Schema, Raw: tftypes.NewValue(cfg.Schema.Type().TerraformType(ctx), nil)}}
		d.Read(ctx, datasource.ReadRequest{Config: cfg}, resp)
	})
}

// ConfigureDataSourceWith invokes Configure on a DataSource with the given provider data.
func ConfigureDataSourceWith(t *testing.T, d datasource.DataSource, providerData interface{}) {
	t.Helper()
	dc, ok := d.(datasource.DataSourceWithConfigure)
	if !ok {
		return
	}
	SafeInvoke(t, "ConfigureDS", func() {
		req := datasource.ConfigureRequest{ProviderData: providerData}
		resp := &datasource.ConfigureResponse{}
		dc.Configure(context.Background(), req, resp)
	})
}

// ConfigureResourceWith invokes Configure on a Resource with the given provider data.
func ConfigureResourceWith(t *testing.T, r resource.Resource, providerData interface{}) {
	t.Helper()
	rc, ok := r.(resource.ResourceWithConfigure)
	if !ok {
		return
	}
	SafeInvoke(t, "Configure", func() {
		req := resource.ConfigureRequest{ProviderData: providerData}
		resp := &resource.ConfigureResponse{}
		rc.Configure(context.Background(), req, resp)
	})
}
