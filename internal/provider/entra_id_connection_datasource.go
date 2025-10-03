// Copyright (c) 2025 Saviynt Inc.
// SPDX-License-Identifier: MPL-2.0

// saviynt_entraid_connection_datasource retrieves entra id connections details from the Saviynt Security Manager.
// The data source supports a single Read operation to look up an existing entra id connections by name.
package provider

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"terraform-provider-Saviynt/internal/client"
	"terraform-provider-Saviynt/util"
	connectionsutil "terraform-provider-Saviynt/util/connectionsutil"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	s "github.com/saviynt/saviynt-api-go-client"

	openapi "github.com/saviynt/saviynt-api-go-client/connections"
)

// EntraIDConnectionDataSource defines the data source
type entraIdConnectionDataSource struct {
	client   *s.Client
	token    string
	provider client.SaviyntProviderInterface
}

type EntraIdConnectionDataSourceModel struct {
	ID types.String `tfsdk:"id"`
	BaseConnectionDataSourceModel
	ConnectionAttributes *EntraIdConnectionAttributes `tfsdk:"connection_attributes"`
}

type EntraIdConnectionAttributes struct {
	UpdateUserJSON           types.String             `tfsdk:"update_user_json"`
	MicrosoftGraphEndpoint   types.String             `tfsdk:"microsoft_graph_endpoint"`
	EndpointsFilter          types.String             `tfsdk:"endpoints_filter"`
	ImportUserJSON           types.String             `tfsdk:"import_user_json"`
	ConnectionType           types.String             `tfsdk:"connection_type"`
	EnableAccountJSON        types.String             `tfsdk:"enable_account_json"`
	DeleteGroupJSON          types.String             `tfsdk:"delete_group_json"`
	ConfigJSON               types.String             `tfsdk:"config_json"`
	AddAccessJSON            types.String             `tfsdk:"add_access_json"`
	CreateChannelJSON        types.String             `tfsdk:"create_channel_json"`
	UpdateAccountJSON        types.String             `tfsdk:"update_account_json"`
	IsTimeoutSupported       types.Bool               `tfsdk:"is_timeout_supported"`
	CreateAccountJSON        types.String             `tfsdk:"create_account_json"`
	PamConfig                types.String             `tfsdk:"pam_config"`
	AzureManagementEndpoint  types.String             `tfsdk:"azure_management_endpoint"`
	EntitlementAttribute     types.String             `tfsdk:"entitlement_attribute"`
	AccountsFilter           types.String             `tfsdk:"accounts_filter"`
	DeltaTokensJSON          types.String             `tfsdk:"deltatokens_json"`
	CreateTeamJSON           types.String             `tfsdk:"create_team_json"`
	StatusThresholdConfig    types.String             `tfsdk:"status_threshold_config"`
	AccountImportFields      types.String             `tfsdk:"account_import_fields"`
	RemoveAccountJSON        types.String             `tfsdk:"remove_account_json"`
	EntitlementFilterJSON    types.String             `tfsdk:"entitlement_filter_json"`
	AuthenticationEndpoint   types.String             `tfsdk:"authentication_endpoint"`
	ModifyUserDataJSON       types.String             `tfsdk:"modifyuserdatajson"`
	IsTimeoutConfigValidated types.Bool               `tfsdk:"is_timeout_config_validated"`
	RemoveAccessJSON         types.String             `tfsdk:"remove_access_json"`
	CreateUsers              types.String             `tfsdk:"createusers"`
	DisableAccountJSON       types.String             `tfsdk:"disable_account_json"`
	CreateNewEndpoints       types.String             `tfsdk:"create_new_endpoints"`
	AccountAttributes        types.String             `tfsdk:"account_attributes"`
	AadTenantID              types.String             `tfsdk:"aad_tenant_id"`
	UpdateGroupJSON          types.String             `tfsdk:"update_group_json"`
	CreateGroupJSON          types.String             `tfsdk:"create_group_json"`
	ConnectionTimeoutConfig  *ConnectionTimeoutConfig `tfsdk:"connection_timeout_config"`
}

var _ datasource.DataSource = &entraIdConnectionDataSource{}

func NewEntraIDConnectionsDataSource() datasource.DataSource {
	return &entraIdConnectionDataSource{}
}

func (d *entraIdConnectionDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "saviynt_entraid_connection_datasource"
}

func EntraIDConnectorsDataSourceSchema() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:    true,
			Description: "Resource ID.",
		},
		"connection_attributes": schema.SingleNestedAttribute{
			Computed: true,
			Attributes: map[string]schema.Attribute{
				"update_user_json":            schema.StringAttribute{Computed: true},
				"microsoft_graph_endpoint":    schema.StringAttribute{Computed: true},
				"endpoints_filter":            schema.StringAttribute{Computed: true},
				"import_user_json":            schema.StringAttribute{Computed: true},
				"connection_type":             schema.StringAttribute{Computed: true},
				"enable_account_json":         schema.StringAttribute{Computed: true},
				"delete_group_json":           schema.StringAttribute{Computed: true},
				"config_json":                 schema.StringAttribute{Computed: true},
				"add_access_json":             schema.StringAttribute{Computed: true},
				"create_channel_json":         schema.StringAttribute{Computed: true},
				"update_account_json":         schema.StringAttribute{Computed: true},
				"is_timeout_supported":        schema.BoolAttribute{Computed: true},
				"create_account_json":         schema.StringAttribute{Computed: true},
				"pam_config":                  schema.StringAttribute{Computed: true},
				"azure_management_endpoint":   schema.StringAttribute{Computed: true},
				"entitlement_attribute":       schema.StringAttribute{Computed: true},
				"accounts_filter":             schema.StringAttribute{Computed: true},
				"deltatokens_json":            schema.StringAttribute{Computed: true},
				"create_team_json":            schema.StringAttribute{Computed: true},
				"status_threshold_config":     schema.StringAttribute{Computed: true},
				"account_import_fields":       schema.StringAttribute{Computed: true},
				"remove_account_json":         schema.StringAttribute{Computed: true},
				"entitlement_filter_json":     schema.StringAttribute{Computed: true},
				"authentication_endpoint":     schema.StringAttribute{Computed: true},
				"modifyuserdatajson":          schema.StringAttribute{Computed: true},
				"is_timeout_config_validated": schema.BoolAttribute{Computed: true},
				"remove_access_json":          schema.StringAttribute{Computed: true},
				"createusers":                 schema.StringAttribute{Computed: true},
				"disable_account_json":        schema.StringAttribute{Computed: true},
				"create_new_endpoints":        schema.StringAttribute{Computed: true},
				"account_attributes":          schema.StringAttribute{Computed: true},
				"aad_tenant_id":               schema.StringAttribute{Computed: true},
				"update_group_json":           schema.StringAttribute{Computed: true},
				"create_group_json":           schema.StringAttribute{Computed: true},
				"connection_timeout_config": schema.SingleNestedAttribute{
					Computed:   true,
					Attributes: ConnectionTimeoutConfigeSchema(),
				},
			},
		},
	}
}

// Schema defines the attributes for the data source
func (d *entraIdConnectionDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: util.EntraIDConnDataSourceDescription,
		Attributes:  connectionsutil.MergeDataSourceAttributes(BaseConnectorDataSourceSchema(), EntraIDConnectorsDataSourceSchema()),
	}
}

func (d *entraIdConnectionDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	// Check if provider data is available.
	if req.ProviderData == nil {
		log.Println("ProviderData is nil, returning early.")
		return
	}

	// Cast provider data to your provider type.
	prov, ok := req.ProviderData.(*SaviyntProvider)
	if !ok {
		resp.Diagnostics.AddError("Unexpected Provider Data", "Expected *SaviyntProvider")
		return
	}

	// Set the client and token from the provider state.
	d.client = prov.client
	d.token = prov.accessToken
	d.provider = &client.SaviyntProviderWrapper{Provider: prov}
}

// SetProvider sets the provider for testing purposes
func (d *entraIdConnectionDataSource) SetProvider(provider client.SaviyntProviderInterface) {
	d.provider = provider
}

func (d *entraIdConnectionDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var state EntraIdConnectionDataSourceModel

	configDiagnostics := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(configDiagnostics...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Prepare request parameters
	reqParams := openapi.GetConnectionDetailsRequest{}

	// Set filters based on provided parameters
	if !state.ConnectionName.IsNull() && state.ConnectionName.ValueString() != "" {
		reqParams.SetConnectionname(state.ConnectionName.ValueString())
	}
	if !state.ConnectionKey.IsNull() {
		connectionKeyInt := state.ConnectionKey.ValueInt64()
		reqParams.SetConnectionkey(strconv.FormatInt(connectionKeyInt, 10))
	}

	var apiResp *openapi.GetConnectionDetailsResponse
	var finalHttpResp *http.Response

	// Execute API call with retry logic
	err := d.provider.AuthenticatedAPICallWithRetry(ctx, "read_entraid_connection_datasource", func(token string) error {
		// Configure API client with current token
		cfg := openapi.NewConfiguration()
		apiBaseURL := strings.TrimPrefix(strings.TrimPrefix(d.client.APIBaseURL(), "https://"), "http://")
		cfg.Host = apiBaseURL
		cfg.Scheme = "https"
		cfg.AddDefaultHeader("Authorization", "Bearer "+token)
		cfg.HTTPClient = http.DefaultClient

		apiClient := openapi.NewAPIClient(cfg)
		apiReq := apiClient.ConnectionsAPI.GetConnectionDetails(ctx).GetConnectionDetailsRequest(reqParams)

		// Execute API request
		resp, hResp, err := apiReq.Execute()
		if hResp != nil && hResp.StatusCode == 401 {
			return fmt.Errorf("401 unauthorized")
		}
		apiResp = resp
		finalHttpResp = hResp // Update on every call including retries
		return err
	})

	if err != nil {
		if finalHttpResp != nil && finalHttpResp.StatusCode != 200 {
			log.Printf("[ERROR] HTTP error while reading EntraId Connector: %s", finalHttpResp.Status)
			var fetchResp map[string]interface{}
			if err := json.NewDecoder(finalHttpResp.Body).Decode(&fetchResp); err != nil {
				resp.Diagnostics.AddError("Failed to decode error response", err.Error())
				return
			}
			resp.Diagnostics.AddError(
				"HTTP Error",
				fmt.Sprintf("HTTP error while reading EntraId Connector for the reasons: %s", fetchResp["msg"]),
			)

		} else {
			log.Printf("[ERROR] API Call Failed: %v", err)
			resp.Diagnostics.AddError("API Call Failed", fmt.Sprintf("Error: %v", err))
		}
		return
	}
	if apiResp != nil && apiResp.EntraIDConnectionResponse == nil {
		error := "Verify the connection type"
		log.Printf("[ERROR]: Verify the connection type given")
		resp.Diagnostics.AddError("Read of EntraID connection failed", error)
		return
	}

	log.Printf("[DEBUG] HTTP Status Code: %d", finalHttpResp.StatusCode)

	state.Msg = util.SafeStringDatasource(apiResp.EntraIDConnectionResponse.Msg)
	state.ErrorCode = util.SafeInt64(apiResp.EntraIDConnectionResponse.Errorcode)
	state.ConnectionName = util.SafeStringDatasource(apiResp.EntraIDConnectionResponse.Connectionname)
	state.ConnectionKey = util.SafeInt64(apiResp.EntraIDConnectionResponse.Connectionkey)
	state.Description = util.SafeStringDatasource(apiResp.EntraIDConnectionResponse.Description)
	state.DefaultSavRoles = util.SafeStringDatasource(apiResp.EntraIDConnectionResponse.Defaultsavroles)
	state.ConnectionType = util.SafeStringDatasource(apiResp.EntraIDConnectionResponse.Connectiontype)
	state.CreatedOn = util.SafeStringDatasource(apiResp.EntraIDConnectionResponse.Createdon)
	state.CreatedBy = util.SafeStringDatasource(apiResp.EntraIDConnectionResponse.Createdby)
	state.UpdatedBy = util.SafeStringDatasource(apiResp.EntraIDConnectionResponse.Updatedby)
	state.EmailTemplate = util.SafeStringDatasource(apiResp.EntraIDConnectionResponse.Emailtemplate)

	if apiResp.EntraIDConnectionResponse.Connectionattributes != nil {
		state.ConnectionAttributes = &EntraIdConnectionAttributes{
			UpdateUserJSON:           util.SafeStringDatasource(apiResp.EntraIDConnectionResponse.Connectionattributes.UpdateUserJSON),
			MicrosoftGraphEndpoint:   util.SafeStringDatasource(apiResp.EntraIDConnectionResponse.Connectionattributes.MICROSOFT_GRAPH_ENDPOINT),
			EndpointsFilter:          util.SafeStringDatasource(apiResp.EntraIDConnectionResponse.Connectionattributes.ENDPOINTS_FILTER),
			ImportUserJSON:           util.SafeStringDatasource(apiResp.EntraIDConnectionResponse.Connectionattributes.ImportUserJSON),
			ConnectionType:           util.SafeStringDatasource(apiResp.EntraIDConnectionResponse.Connectionattributes.ConnectionType),
			EnableAccountJSON:        util.SafeStringDatasource(apiResp.EntraIDConnectionResponse.Connectionattributes.EnableAccountJSON),
			DeleteGroupJSON:          util.SafeStringDatasource(apiResp.EntraIDConnectionResponse.Connectionattributes.DeleteGroupJSON),
			ConfigJSON:               util.SafeStringDatasource(apiResp.EntraIDConnectionResponse.Connectionattributes.ConfigJSON),
			AddAccessJSON:            util.SafeStringDatasource(apiResp.EntraIDConnectionResponse.Connectionattributes.AddAccessJSON),
			CreateChannelJSON:        util.SafeStringDatasource(apiResp.EntraIDConnectionResponse.Connectionattributes.CreateChannelJSON),
			UpdateAccountJSON:        util.SafeStringDatasource(apiResp.EntraIDConnectionResponse.Connectionattributes.UpdateAccountJSON),
			IsTimeoutSupported:       util.SafeBoolDatasource(apiResp.EntraIDConnectionResponse.Connectionattributes.IsTimeoutSupported),
			CreateAccountJSON:        util.SafeStringDatasource(apiResp.EntraIDConnectionResponse.Connectionattributes.CreateAccountJSON),
			PamConfig:                util.SafeStringDatasource(apiResp.EntraIDConnectionResponse.Connectionattributes.PAM_CONFIG),
			AzureManagementEndpoint:  util.SafeStringDatasource(apiResp.EntraIDConnectionResponse.Connectionattributes.AZURE_MANAGEMENT_ENDPOINT),
			EntitlementAttribute:     util.SafeStringDatasource(apiResp.EntraIDConnectionResponse.Connectionattributes.ENTITLEMENT_ATTRIBUTE),
			AccountsFilter:           util.SafeStringDatasource(apiResp.EntraIDConnectionResponse.Connectionattributes.ACCOUNTS_FILTER),
			DeltaTokensJSON:          util.SafeStringDatasource(apiResp.EntraIDConnectionResponse.Connectionattributes.DELTATOKENSJSON),
			CreateTeamJSON:           util.SafeStringDatasource(apiResp.EntraIDConnectionResponse.Connectionattributes.CreateTeamJSON),
			StatusThresholdConfig:    util.SafeStringDatasource(apiResp.EntraIDConnectionResponse.Connectionattributes.STATUS_THRESHOLD_CONFIG),
			AccountImportFields:      util.SafeStringDatasource(apiResp.EntraIDConnectionResponse.Connectionattributes.ACCOUNT_IMPORT_FIELDS),
			RemoveAccountJSON:        util.SafeStringDatasource(apiResp.EntraIDConnectionResponse.Connectionattributes.RemoveAccountJSON),
			EntitlementFilterJSON:    util.SafeStringDatasource(apiResp.EntraIDConnectionResponse.Connectionattributes.ENTITLEMENT_FILTER_JSON),
			AuthenticationEndpoint:   util.SafeStringDatasource(apiResp.EntraIDConnectionResponse.Connectionattributes.AUTHENTICATION_ENDPOINT),
			ModifyUserDataJSON:       util.SafeStringDatasource(apiResp.EntraIDConnectionResponse.Connectionattributes.MODIFYUSERDATAJSON),
			IsTimeoutConfigValidated: util.SafeBoolDatasource(apiResp.EntraIDConnectionResponse.Connectionattributes.IsTimeoutConfigValidated),
			RemoveAccessJSON:         util.SafeStringDatasource(apiResp.EntraIDConnectionResponse.Connectionattributes.RemoveAccessJSON),
			CreateUsers:              util.SafeStringDatasource(apiResp.EntraIDConnectionResponse.Connectionattributes.CREATEUSERS),
			DisableAccountJSON:       util.SafeStringDatasource(apiResp.EntraIDConnectionResponse.Connectionattributes.DisableAccountJSON),
			CreateNewEndpoints:       util.SafeStringDatasource(apiResp.EntraIDConnectionResponse.Connectionattributes.CREATE_NEW_ENDPOINTS),
			AccountAttributes:        util.SafeStringDatasource(apiResp.EntraIDConnectionResponse.Connectionattributes.ACCOUNT_ATTRIBUTES),
			AadTenantID:              util.SafeStringDatasource(apiResp.EntraIDConnectionResponse.Connectionattributes.AAD_TENANT_ID),
			UpdateGroupJSON:          util.SafeStringDatasource(apiResp.EntraIDConnectionResponse.Connectionattributes.UpdateGroupJSON),
			CreateGroupJSON:          util.SafeStringDatasource(apiResp.EntraIDConnectionResponse.Connectionattributes.CreateGroupJSON),
		}
		if apiResp.EntraIDConnectionResponse.Connectionattributes.ConnectionTimeoutConfig != nil {
			state.ConnectionAttributes.ConnectionTimeoutConfig = &ConnectionTimeoutConfig{
				RetryWait:               util.SafeInt64(apiResp.EntraIDConnectionResponse.Connectionattributes.ConnectionTimeoutConfig.RetryWait),
				TokenRefreshMaxTryCount: util.SafeInt64(apiResp.EntraIDConnectionResponse.Connectionattributes.ConnectionTimeoutConfig.TokenRefreshMaxTryCount),
				RetryFailureStatusCode:  util.SafeInt64(apiResp.EntraIDConnectionResponse.Connectionattributes.ConnectionTimeoutConfig.RetryFailureStatusCode),
				RetryWaitMaxValue:       util.SafeInt64(apiResp.EntraIDConnectionResponse.Connectionattributes.ConnectionTimeoutConfig.RetryWaitMaxValue),
				RetryCount:              util.SafeInt64(apiResp.EntraIDConnectionResponse.Connectionattributes.ConnectionTimeoutConfig.RetryCount),
				ReadTimeout:             util.SafeInt64(apiResp.EntraIDConnectionResponse.Connectionattributes.ConnectionTimeoutConfig.ReadTimeout),
				ConnectionTimeout:       util.SafeInt64(apiResp.EntraIDConnectionResponse.Connectionattributes.ConnectionTimeoutConfig.ConnectionTimeout),
			}
		}
	}

	if apiResp.EntraIDConnectionResponse.Connectionattributes == nil {
		state.ConnectionAttributes = nil
	}
	if !state.Authenticate.IsNull() && !state.Authenticate.IsUnknown() {
		if state.Authenticate.ValueBool() {
			resp.Diagnostics.AddWarning(
				"Authentication Enabled",
				"`authenticate` is true; all connection_attributes will be returned in state.",
			)
		} else {
			resp.Diagnostics.AddWarning(
				"Authentication Disabled",
				"`authenticate` is false; connection_attributes will be removed from state.",
			)
			state.ConnectionAttributes = nil
		}
	}
	stateDiagnostics := resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(stateDiagnostics...)
}
