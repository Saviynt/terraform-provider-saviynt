/*
 * Copyright (c) 2025 Saviynt Inc.
 * All Rights Reserved.
 *
 * This software is the confidential and proprietary information of
 * Saviynt Inc. ("Confidential Information"). You shall not disclose,
 * use, or distribute such Confidential Information except in accordance
 * with the terms of the license agreement you entered into with Saviynt.
 *
 * SAVIYNT MAKES NO REPRESENTATIONS OR WARRANTIES ABOUT THE SUITABILITY OF
 * THE SOFTWARE, EITHER EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO
 * THE IMPLIED WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR
 * PURPOSE, OR NON-INFRINGEMENT.
 */

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
	client *s.Client
	token  string
}

type EntraIdConnectionDataSourceModel struct {
	ID types.String `tfsdk:"id"`
	BaseConnectionDataSourceModel
	ConnectionAttributes *EntraIdConnectionAttributes `tfsdk:"connection_attributes"`
}

type EntraIdConnectionAttributes struct {
	UpdateUserJSON         types.String `tfsdk:"update_user_json"`
	MicrosoftGraphEndpoint types.String `tfsdk:"microsoft_graph_endpoint"`
	EndpointsFilter        types.String `tfsdk:"endpoints_filter"`
	ImportUserJSON         types.String `tfsdk:"import_user_json"`
	ConnectionType         types.String `tfsdk:"connection_type"`
	EnableAccountJSON      types.String `tfsdk:"enable_account_json"`
	// ConnectionJSON         types.String `tfsdk:"connection_json"`
	// ClientID               types.String `tfsdk:"client_id"`
	DeleteGroupJSON        types.String `tfsdk:"delete_group_json"`
	ConfigJSON             types.String `tfsdk:"config_json"`
	// AccessToken            types.String `tfsdk:"access_token"`
	AddAccessJSON          types.String `tfsdk:"add_access_json"`
	CreateChannelJSON      types.String `tfsdk:"create_channel_json"`
	UpdateAccountJSON      types.String `tfsdk:"update_account_json"`
	IsTimeoutSupported     types.Bool   `tfsdk:"is_timeout_supported"`
	// RemoveServicePrincipalJSON      types.String `tfsdk:"remove_service_principal_json"`
	// ImportDepth                     types.String `tfsdk:"import_depth"`
	CreateAccountJSON types.String `tfsdk:"create_account_json"`
	PamConfig         types.String `tfsdk:"pam_config"`
	// UpdateServicePrincipalJSON      types.String `tfsdk:"update_service_principal_json"`
	AzureManagementEndpoint types.String `tfsdk:"azure_management_endpoint"`
	EntitlementAttribute    types.String `tfsdk:"entitlement_attribute"`
	AccountsFilter          types.String `tfsdk:"accounts_filter"`
	// WindowsConnectorJSON    types.String `tfsdk:"windows_connector_json"`
	DeltaTokensJSON         types.String `tfsdk:"deltatokens_json"`
	// AzureMgmtAccessToken    types.String `tfsdk:"azure_mgmt_access_token"`
	CreateTeamJSON          types.String `tfsdk:"create_team_json"`
	// EnhancedDirectoryRoles          types.String `tfsdk:"enhanceddirectoryroles"`
	StatusThresholdConfig types.String `tfsdk:"status_threshold_config"`
	AccountImportFields   types.String `tfsdk:"account_import_fields"`
	RemoveAccountJSON     types.String `tfsdk:"remove_account_json"`
	// ChangePassJSON        types.String `tfsdk:"change_pass_json"`
	// ClientSecret          types.String `tfsdk:"client_secret"`
	EntitlementFilterJSON types.String `tfsdk:"entitlement_filter_json"`
	// ServiceAccountAttributes        types.String `tfsdk:"service_account_attributes"`
	// AddAccessToEntitlementJSON      types.String `tfsdk:"add_access_to_entitlement_json"`
	AuthenticationEndpoint types.String `tfsdk:"authentication_endpoint"`
	// CreateServicePrincipalJSON      types.String `tfsdk:"create_service_principal_json"`
	ModifyUserDataJSON       types.String `tfsdk:"modifyuserdatajson"`
	IsTimeoutConfigValidated types.Bool   `tfsdk:"is_timeout_config_validated"`
	RemoveAccessJSON         types.String `tfsdk:"remove_access_json"`
	CreateUsers              types.String `tfsdk:"createusers"`
	// RemoveAccessFromEntitlementJSON types.String `tfsdk:"remove_access_from_entitlement_json"`
	DisableAccountJSON types.String `tfsdk:"disable_account_json"`
	CreateNewEndpoints types.String `tfsdk:"create_new_endpoints"`
	// ManagedAccountType              types.String `tfsdk:"managed_account_type"`
	AccountAttributes types.String `tfsdk:"account_attributes"`
	AadTenantID       types.String `tfsdk:"aad_tenant_id"`
	UpdateGroupJSON   types.String `tfsdk:"update_group_json"`
	CreateGroupJSON   types.String `tfsdk:"create_group_json"`
	// ConnectionTimeoutConfig         *ConnectionTimeoutConfig `tfsdk:"connection_timeout_config"`
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
				"update_user_json":         schema.StringAttribute{Computed: true},
				"microsoft_graph_endpoint": schema.StringAttribute{Computed: true},
				"endpoints_filter":         schema.StringAttribute{Computed: true},
				"import_user_json":         schema.StringAttribute{Computed: true},
				"connection_type":          schema.StringAttribute{Computed: true},
				"enable_account_json":      schema.StringAttribute{Computed: true},
				// "connection_json":          schema.StringAttribute{Computed: true},
				// "client_id":                schema.StringAttribute{Computed: true},
				"delete_group_json":        schema.StringAttribute{Computed: true},
				"config_json":              schema.StringAttribute{Computed: true},
				// "access_token":             schema.StringAttribute{Computed: true},
				"add_access_json":          schema.StringAttribute{Computed: true},
				"create_channel_json":      schema.StringAttribute{Computed: true},
				"update_account_json":      schema.StringAttribute{Computed: true},
				"is_timeout_supported":     schema.BoolAttribute{Computed: true},
				// "remove_service_principal_json":       schema.StringAttribute{Computed: true},
				// "import_depth":                        schema.StringAttribute{Computed: true},
				"create_account_json": schema.StringAttribute{Computed: true},
				"pam_config":          schema.StringAttribute{Computed: true},
				// "update_service_principal_json":       schema.StringAttribute{Computed: true},
				"azure_management_endpoint": schema.StringAttribute{Computed: true},
				"entitlement_attribute":     schema.StringAttribute{Computed: true},
				"accounts_filter":           schema.StringAttribute{Computed: true},
				// "windows_connector_json":    schema.StringAttribute{Computed: true},
				"deltatokens_json":          schema.StringAttribute{Computed: true},
				// "azure_mgmt_access_token":   schema.StringAttribute{Computed: true},
				"create_team_json":          schema.StringAttribute{Computed: true},
				// "enhanceddirectoryroles":              schema.StringAttribute{Computed: true},
				"status_threshold_config": schema.StringAttribute{Computed: true},
				"account_import_fields":   schema.StringAttribute{Computed: true},
				"remove_account_json":     schema.StringAttribute{Computed: true},
				// "change_pass_json":        schema.StringAttribute{Computed: true},
				// "client_secret":           schema.StringAttribute{Computed: true},
				"entitlement_filter_json": schema.StringAttribute{Computed: true},
				// "service_account_attributes":          schema.StringAttribute{Computed: true},
				// "add_access_to_entitlement_json":      schema.StringAttribute{Computed: true},
				"authentication_endpoint": schema.StringAttribute{Computed: true},
				// "create_service_principal_json":       schema.StringAttribute{Computed: true},
				"modifyuserdatajson":          schema.StringAttribute{Computed: true},
				"is_timeout_config_validated": schema.BoolAttribute{Computed: true},
				"remove_access_json":          schema.StringAttribute{Computed: true},
				"createusers":                 schema.StringAttribute{Computed: true},
				// "remove_access_from_entitlement_json": schema.StringAttribute{Computed: true},
				"disable_account_json": schema.StringAttribute{Computed: true},
				"create_new_endpoints": schema.StringAttribute{Computed: true},
				// "managed_account_type":                schema.StringAttribute{Computed: true},
				"account_attributes": schema.StringAttribute{Computed: true},
				"aad_tenant_id":      schema.StringAttribute{Computed: true},
				"update_group_json":  schema.StringAttribute{Computed: true},
				"create_group_json":  schema.StringAttribute{Computed: true},
				// "connection_timeout_config": schema.SingleNestedAttribute{
				// 	Computed:   true,
				// 	Attributes: ConnectionTimeoutConfigeSchema(),
				// },
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
	prov, ok := req.ProviderData.(*saviyntProvider)
	if !ok {
		resp.Diagnostics.AddError("Unexpected Provider Data", "Expected *saviyntProvider")
		return
	}

	// Set the client and token from the provider state.
	d.client = prov.client
	d.token = prov.accessToken
}

func (d *entraIdConnectionDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var state EntraIdConnectionDataSourceModel

	configDiagnostics := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(configDiagnostics...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Configure API client
	cfg := openapi.NewConfiguration()
	apiBaseURL := strings.TrimPrefix(strings.TrimPrefix(d.client.APIBaseURL(), "https://"), "http://")
	cfg.Host = apiBaseURL
	cfg.Scheme = "https"
	cfg.AddDefaultHeader("Authorization", "Bearer "+d.token)
	cfg.HTTPClient = http.DefaultClient

	apiClient := openapi.NewAPIClient(cfg)
	reqParams := openapi.GetConnectionDetailsRequest{}

	// Set filters based on provided parameters
	if !state.ConnectionName.IsNull() && state.ConnectionName.ValueString() != "" {
		reqParams.SetConnectionname(state.ConnectionName.ValueString())
	}
	if !state.ConnectionKey.IsNull() {
		connectionKeyInt := state.ConnectionKey.ValueInt64()
		reqParams.SetConnectionkey(strconv.FormatInt(connectionKeyInt, 10))
	}
	apiReq := apiClient.ConnectionsAPI.GetConnectionDetails(ctx).GetConnectionDetailsRequest(reqParams)

	// Execute API request
	apiResp, httpResp, err := apiReq.Execute()
	if err != nil {
		if httpResp != nil && httpResp.StatusCode != 200 {
			log.Printf("[ERROR] HTTP error while creating EntraId Connector: %s", httpResp.Status)
			var fetchResp map[string]interface{}
			if err := json.NewDecoder(httpResp.Body).Decode(&fetchResp); err != nil {
				resp.Diagnostics.AddError("Failed to decode error response", err.Error())
				return
			}
			resp.Diagnostics.AddError(
				"HTTP Error",
				fmt.Sprintf("HTTP error while creating EntraId Connector for the reasons: %s", fetchResp["msg"]),
			)

		} else {
			log.Printf("[ERROR] API Call Failed: %v", err)
			resp.Diagnostics.AddError("API Call Failed", fmt.Sprintf("Error: %v", err))
		}
		return
	}

	// if apiResp != nil && *apiResp.EntraIDConnectionResponse.Errorcode != 0 {
	// 	log.Printf("[ERROR]: Error in reading EntraId connection. Errorcode: %v, Message: %v", *apiResp.EntraIDConnectionResponse.Errorcode, *apiResp.EntraIDConnectionResponse.Msg)
	// 	resp.Diagnostics.AddError("Read DB connection failed", *apiResp.EntraIDConnectionResponse.Msg)
	// 	return
	// }

	if apiResp != nil && apiResp.EntraIDConnectionResponse == nil {
		error := "Verify the connection type"
		log.Printf("[ERROR]: Verify the connection type given")
		resp.Diagnostics.AddError("Read of EntraID connection failed", error)
		return
	}

	log.Printf("[DEBUG] HTTP Status Code: %d", httpResp.StatusCode)

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
			UpdateUserJSON:         util.SafeStringDatasource(apiResp.EntraIDConnectionResponse.Connectionattributes.UpdateUserJSON),
			MicrosoftGraphEndpoint: util.SafeStringDatasource(apiResp.EntraIDConnectionResponse.Connectionattributes.MICROSOFT_GRAPH_ENDPOINT),
			EndpointsFilter:        util.SafeStringDatasource(apiResp.EntraIDConnectionResponse.Connectionattributes.ENDPOINTS_FILTER),
			ImportUserJSON:         util.SafeStringDatasource(apiResp.EntraIDConnectionResponse.Connectionattributes.ImportUserJSON),
			ConnectionType:         util.SafeStringDatasource(apiResp.EntraIDConnectionResponse.Connectionattributes.ConnectionType),
			EnableAccountJSON:      util.SafeStringDatasource(apiResp.EntraIDConnectionResponse.Connectionattributes.EnableAccountJSON),
			// ConnectionJSON:         util.SafeStringDatasource(apiResp.EntraIDConnectionResponse.Connectionattributes.ConnectionJSON),
			// ClientID:               util.SafeStringDatasource(apiResp.EntraIDConnectionResponse.Connectionattributes.CLIENT_ID),
			DeleteGroupJSON:        util.SafeStringDatasource(apiResp.EntraIDConnectionResponse.Connectionattributes.DeleteGroupJSON),
			ConfigJSON:             util.SafeStringDatasource(apiResp.EntraIDConnectionResponse.Connectionattributes.ConfigJSON),
			// AccessToken:            util.SafeStringDatasource(apiResp.EntraIDConnectionResponse.Connectionattributes.ACCESS_TOKEN),
			AddAccessJSON:          util.SafeStringDatasource(apiResp.EntraIDConnectionResponse.Connectionattributes.AddAccessJSON),
			CreateChannelJSON:      util.SafeStringDatasource(apiResp.EntraIDConnectionResponse.Connectionattributes.CreateChannelJSON),
			UpdateAccountJSON:      util.SafeStringDatasource(apiResp.EntraIDConnectionResponse.Connectionattributes.UpdateAccountJSON),
			IsTimeoutSupported:     util.SafeBoolDatasource(apiResp.EntraIDConnectionResponse.Connectionattributes.IsTimeoutSupported),
			// RemoveServicePrincipalJSON:      util.SafeStringDatasource(apiResp.EntraIDConnectionResponse.Connectionattributes.RemoveServicePrincipalJSON),
			// ImportDepth:                     util.SafeStringDatasource(apiResp.EntraIDConnectionResponse.Connectionattributes.IMPORT_DEPTH),
			CreateAccountJSON: util.SafeStringDatasource(apiResp.EntraIDConnectionResponse.Connectionattributes.CreateAccountJSON),
			PamConfig:         util.SafeStringDatasource(apiResp.EntraIDConnectionResponse.Connectionattributes.PAM_CONFIG),
			// UpdateServicePrincipalJSON:      util.SafeStringDatasource(apiResp.EntraIDConnectionResponse.Connectionattributes.UpdateServicePrincipalJSON),
			AzureManagementEndpoint: util.SafeStringDatasource(apiResp.EntraIDConnectionResponse.Connectionattributes.AZURE_MANAGEMENT_ENDPOINT),
			EntitlementAttribute:    util.SafeStringDatasource(apiResp.EntraIDConnectionResponse.Connectionattributes.ENTITLEMENT_ATTRIBUTE),
			AccountsFilter:          util.SafeStringDatasource(apiResp.EntraIDConnectionResponse.Connectionattributes.ACCOUNTS_FILTER),
			// WindowsConnectorJSON:    util.SafeStringDatasource(apiResp.EntraIDConnectionResponse.Connectionattributes.WINDOWS_CONNECTOR_JSON),
			DeltaTokensJSON:         util.SafeStringDatasource(apiResp.EntraIDConnectionResponse.Connectionattributes.DELTATOKENSJSON),
			// AzureMgmtAccessToken:    util.SafeStringDatasource(apiResp.EntraIDConnectionResponse.Connectionattributes.AZURE_MGMT_ACCESS_TOKEN),
			CreateTeamJSON:          util.SafeStringDatasource(apiResp.EntraIDConnectionResponse.Connectionattributes.CreateTeamJSON),
			// EnhancedDirectoryRoles:          util.SafeStringDatasource(apiResp.EntraIDConnectionResponse.Connectionattributes.ENHANCEDDIRECTORYROLES),
			StatusThresholdConfig: util.SafeStringDatasource(apiResp.EntraIDConnectionResponse.Connectionattributes.STATUS_THRESHOLD_CONFIG),
			AccountImportFields:   util.SafeStringDatasource(apiResp.EntraIDConnectionResponse.Connectionattributes.ACCOUNT_IMPORT_FIELDS),
			RemoveAccountJSON:     util.SafeStringDatasource(apiResp.EntraIDConnectionResponse.Connectionattributes.RemoveAccountJSON),
			// ChangePassJSON:        util.SafeStringDatasource(apiResp.EntraIDConnectionResponse.Connectionattributes.ChangePassJSON),
			// ClientSecret:          util.SafeStringDatasource(apiResp.EntraIDConnectionResponse.Connectionattributes.CLIENT_SECRET),
			EntitlementFilterJSON: util.SafeStringDatasource(apiResp.EntraIDConnectionResponse.Connectionattributes.ENTITLEMENT_FILTER_JSON),
			// ServiceAccountAttributes:        util.SafeStringDatasource(apiResp.EntraIDConnectionResponse.Connectionattributes.SERVICE_ACCOUNT_ATTRIBUTES),
			// AddAccessToEntitlementJSON:      util.SafeStringDatasource(apiResp.EntraIDConnectionResponse.Connectionattributes.AddAccessToEntitlementJSON),
			AuthenticationEndpoint: util.SafeStringDatasource(apiResp.EntraIDConnectionResponse.Connectionattributes.AUTHENTICATION_ENDPOINT),
			// CreateServicePrincipalJSON:      util.SafeStringDatasource(apiResp.EntraIDConnectionResponse.Connectionattributes.CreateServicePrincipalJSON),
			ModifyUserDataJSON:       util.SafeStringDatasource(apiResp.EntraIDConnectionResponse.Connectionattributes.MODIFYUSERDATAJSON),
			IsTimeoutConfigValidated: util.SafeBoolDatasource(apiResp.EntraIDConnectionResponse.Connectionattributes.IsTimeoutConfigValidated),
			RemoveAccessJSON:         util.SafeStringDatasource(apiResp.EntraIDConnectionResponse.Connectionattributes.RemoveAccessJSON),
			CreateUsers:              util.SafeStringDatasource(apiResp.EntraIDConnectionResponse.Connectionattributes.CREATEUSERS),
			// RemoveAccessFromEntitlementJSON: util.SafeStringDatasource(apiResp.EntraIDConnectionResponse.Connectionattributes.RemoveAccessFromEntitlementJSON),
			DisableAccountJSON: util.SafeStringDatasource(apiResp.EntraIDConnectionResponse.Connectionattributes.DisableAccountJSON),
			CreateNewEndpoints: util.SafeStringDatasource(apiResp.EntraIDConnectionResponse.Connectionattributes.CREATE_NEW_ENDPOINTS),
			// ManagedAccountType:              util.SafeStringDatasource(apiResp.EntraIDConnectionResponse.Connectionattributes.MANAGED_ACCOUNT_TYPE),
			AccountAttributes: util.SafeStringDatasource(apiResp.EntraIDConnectionResponse.Connectionattributes.ACCOUNT_ATTRIBUTES),
			AadTenantID:       util.SafeStringDatasource(apiResp.EntraIDConnectionResponse.Connectionattributes.AAD_TENANT_ID),
			UpdateGroupJSON:   util.SafeStringDatasource(apiResp.EntraIDConnectionResponse.Connectionattributes.UpdateGroupJSON),
			CreateGroupJSON:   util.SafeStringDatasource(apiResp.EntraIDConnectionResponse.Connectionattributes.CreateGroupJSON),
		}
		// if apiResp.EntraIDConnectionResponse.Connectionattributes.ConnectionTimeoutConfig != nil {
		// 	state.ConnectionAttributes.ConnectionTimeoutConfig = &ConnectionTimeoutConfig{
		// 		RetryWait:               util.SafeInt64(apiResp.EntraIDConnectionResponse.Connectionattributes.ConnectionTimeoutConfig.RetryWait),
		// 		TokenRefreshMaxTryCount: util.SafeInt64(apiResp.EntraIDConnectionResponse.Connectionattributes.ConnectionTimeoutConfig.TokenRefreshMaxTryCount),
		// 		RetryFailureStatusCode:  util.SafeInt64(apiResp.EntraIDConnectionResponse.Connectionattributes.ConnectionTimeoutConfig.RetryFailureStatusCode),
		// 		RetryWaitMaxValue:       util.SafeInt64(apiResp.EntraIDConnectionResponse.Connectionattributes.ConnectionTimeoutConfig.RetryWaitMaxValue),
		// 		RetryCount:              util.SafeInt64(apiResp.EntraIDConnectionResponse.Connectionattributes.ConnectionTimeoutConfig.RetryCount),
		// 		ReadTimeout:             util.SafeInt64(apiResp.EntraIDConnectionResponse.Connectionattributes.ConnectionTimeoutConfig.ReadTimeout),
		// 		ConnectionTimeout:       util.SafeInt64(apiResp.EntraIDConnectionResponse.Connectionattributes.ConnectionTimeoutConfig.ConnectionTimeout),
		// 	}
		// }
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
