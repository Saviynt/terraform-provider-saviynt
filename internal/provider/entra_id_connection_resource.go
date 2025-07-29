// Copyright (c) 2025 Saviynt Inc.
// SPDX-License-Identifier: MPL-2.0

// saviynt_entraid_connection_resource manages EntraId connectors in the Saviynt Security Manager.
// The resource implements the full Terraform lifecycle:
//   - Create: provisions a new EntraId connector using the supplied configuration.
//   - Read: fetches the current connector state from Saviynt to keep Terraform’s state in sync.
//   - Update: applies any configuration changes to an existing connector.
//   - Import: brings an existing connector under Terraform management by its name.
package provider

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"terraform-provider-Saviynt/util"
	connectionsutil "terraform-provider-Saviynt/util/connectionsutil"

	openapi "github.com/saviynt/saviynt-api-go-client/connections"

	s "github.com/saviynt/saviynt-api-go-client"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &entraIdConnectionResource{}
var _ resource.ResourceWithImportState = &entraIdConnectionResource{}

type EntraIdConnectorResourceModel struct {
	BaseConnectorResourceModel
	ID                              types.String `tfsdk:"id"`
	ClientId                        types.String `tfsdk:"client_id"`
	ClientSecret                    types.String `tfsdk:"client_secret"`
	AccessToken                     types.String `tfsdk:"access_token"`
	AadTenantId                     types.String `tfsdk:"aad_tenant_id"`
	AzureMgmtAccessToken            types.String `tfsdk:"azure_mgmt_access_token"`
	AuthenticationEndpoint          types.String `tfsdk:"authentication_endpoint"`
	MicrosoftGraphEndpoint          types.String `tfsdk:"microsoft_graph_endpoint"`
	AzureManagementEndpoint         types.String `tfsdk:"azure_management_endpoint"`
	ImportUserJson                  types.String `tfsdk:"import_user_json"`
	CreateUsers                     types.String `tfsdk:"create_users"`
	WindowsConnectorJson            types.String `tfsdk:"windows_connector_json"`
	CreateNewEndpoints              types.String `tfsdk:"create_new_endpoints"`
	ManagedAccountType              types.String `tfsdk:"managed_account_type"`
	AccountAttributes               types.String `tfsdk:"account_attributes"`
	ServiceAccountAttributes        types.String `tfsdk:"service_account_attributes"`
	DeltaTokensJson                 types.String `tfsdk:"delta_tokens_json"`
	AccountImportFields             types.String `tfsdk:"account_import_fields"`
	ImportDepth                     types.String `tfsdk:"import_depth"`
	EntitlementAttribute            types.String `tfsdk:"entitlement_attribute"`
	CreateAccountJson               types.String `tfsdk:"create_account_json"`
	UpdateAccountJson               types.String `tfsdk:"update_account_json"`
	EnableAccountJson               types.String `tfsdk:"enable_account_json"`
	DisableAccountJson              types.String `tfsdk:"disable_account_json"`
	AddAccessJson                   types.String `tfsdk:"add_access_json"`
	RemoveAccessJson                types.String `tfsdk:"remove_access_json"`
	UpdateUserJson                  types.String `tfsdk:"update_user_json"`
	ChangePassJson                  types.String `tfsdk:"change_pass_json"`
	RemoveAccountJson               types.String `tfsdk:"remove_account_json"`
	ConnectionJson                  types.String `tfsdk:"connection_json"`
	CreateGroupJson                 types.String `tfsdk:"create_group_json"`
	UpdateGroupJson                 types.String `tfsdk:"update_group_json"`
	AddAccessToEntitlementJson      types.String `tfsdk:"add_access_to_entitlement_json"`
	RemoveAccessFromEntitlementJson types.String `tfsdk:"remove_access_from_entitlement_json"`
	DeleteGroupJson                 types.String `tfsdk:"delete_group_json"`
	CreateServicePrincipalJson      types.String `tfsdk:"create_service_principal_json"`
	UpdateServicePrincipalJson      types.String `tfsdk:"update_service_principal_json"`
	RemoveServicePrincipalJson      types.String `tfsdk:"remove_service_principal_json"`
	EntitlementFilterJson           types.String `tfsdk:"entitlement_filter_json"`
	CreateTeamJson                  types.String `tfsdk:"create_team_json"`
	CreateChannelJson               types.String `tfsdk:"create_channel_json"`
	StatusThresholdConfig           types.String `tfsdk:"status_threshold_config"`
	AccountsFilter                  types.String `tfsdk:"accounts_filter"`
	PamConfig                       types.String `tfsdk:"pam_config"`
	EndpointsFilter                 types.String `tfsdk:"endpoints_filter"`
	ConfigJson                      types.String `tfsdk:"config_json"`
	ModifyUserdataJson              types.String `tfsdk:"modify_user_data_json"`
	EnhancedDirectoryRoles          types.String `tfsdk:"enhanced_directory_roles"`
}

type entraIdConnectionResource struct {
	client *s.Client
	token  string
}

func NewEntraIdTestConnectionResource() resource.Resource {
	return &entraIdConnectionResource{}
}

func (r *entraIdConnectionResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = "saviynt_entraid_connection_resource"
}

func EntraIdConnectorResourceSchema() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:    true,
			Description: "Resource ID.",
		},
		"client_id": schema.StringAttribute{
			Required:    true,
			WriteOnly:   true,
			Description: "Client ID for authentication.",
		},
		"client_secret": schema.StringAttribute{
			Required:    true,
			WriteOnly:   true,
			Description: "Client Secret for authentication.",
		},
		"access_token": schema.StringAttribute{
			Optional:    true,
			WriteOnly:   true,
			Description: "Access token used for API calls.",
		},
		"aad_tenant_id": schema.StringAttribute{
			Required:    true,
			Description: "Azure Active Directory tenant ID.",
		},
		"azure_mgmt_access_token": schema.StringAttribute{
			Optional:    true,
			WriteOnly:   true,
			Description: "Access token for Azure management APIs.",
		},
		"authentication_endpoint": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "Authentication endpoint URL.",
		},
		"microsoft_graph_endpoint": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "Microsoft Graph API endpoint.",
		},
		"azure_management_endpoint": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "Azure management endpoint URL.",
		},
		"import_user_json": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "JSON configuration for importing users.",
		},
		"create_users": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "Flag or configuration for creating users.",
		},
		"windows_connector_json": schema.StringAttribute{
			Optional:    true,
			WriteOnly:   true,
			Description: "Windows connector JSON configuration.",
		},
		"create_new_endpoints": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "Configuration to create new endpoints.Value accpetd are YES/NO.",
		},
		"managed_account_type": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "Type of managed accounts.",
		},
		"account_attributes": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "Attributes for account configuration.",
		},
		"service_account_attributes": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "Attributes for service account configuration.",
		},
		"delta_tokens_json": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "Delta tokens JSON data.",
		},
		"account_import_fields": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "Fields to import for accounts.",
		},
		"import_depth": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "Depth level for import.",
		},
		"entitlement_attribute": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "Attribute used for entitlement.",
		},
		"create_account_json": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "JSON template to create an account.",
		},
		"update_account_json": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "JSON template to update an account.",
		},
		"enable_account_json": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "JSON template to enable an account.",
		},
		"disable_account_json": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "JSON template to disable an account.",
		},
		"add_access_json": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "JSON template to add access.",
		},
		"remove_access_json": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "JSON template to remove access.",
		},
		"update_user_json": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "JSON template to update user.",
		},
		"change_pass_json": schema.StringAttribute{
			Optional:    true,
			WriteOnly:   true,
			Description: "JSON template to change password.",
		},
		"remove_account_json": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "JSON template to remove account.",
		},
		"connection_json": schema.StringAttribute{
			Optional:    true,
			WriteOnly:   true,
			Description: "Connection JSON configuration.",
		},
		"create_group_json": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "JSON to create group.",
		},
		"update_group_json": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "JSON to update group.",
		},
		"add_access_to_entitlement_json": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "JSON to add access to entitlement.",
		},
		"remove_access_from_entitlement_json": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "JSON to remove access from entitlement.",
		},
		"delete_group_json": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "JSON to delete group.",
		},
		"create_service_principal_json": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "JSON to create service principal.",
		},
		"update_service_principal_json": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "JSON to update service principal.",
		},
		"remove_service_principal_json": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "JSON to remove service principal.",
		},
		"entitlement_filter_json": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "Filter JSON for entitlements.",
		},
		"create_team_json": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "JSON to create team.",
		},
		"create_channel_json": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "JSON to create channel.",
		},
		"status_threshold_config": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "Configuration for status thresholds.",
		},
		"accounts_filter": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "Filter for accounts.",
		},
		"pam_config": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "PAM configuration.",
		},
		"endpoints_filter": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "Endpoints filter configuration.",
		},
		"config_json": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "Main config JSON.",
		},
		"modify_user_data_json": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "JSON to modify user data.",
		},
		"enhanced_directory_roles": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "Configuration for enhanced directory roles.",
		},
	}
}

func (r *entraIdConnectionResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: util.EntraIDConnDescription,
		Attributes:  connectionsutil.MergeResourceAttributes(BaseConnectorResourceSchema(), EntraIdConnectorResourceSchema()),
	}
}

func (r *entraIdConnectionResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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
	r.client = prov.client
	r.token = prov.accessToken
}

func (r *entraIdConnectionResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan, config EntraIdConnectorResourceModel
	// Extract plan from request
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}
	// Extract config from request
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}
	cfg := openapi.NewConfiguration()
	apiBaseURL := strings.TrimPrefix(strings.TrimPrefix(r.client.APIBaseURL(), "https://"), "http://")
	cfg.Host = apiBaseURL
	cfg.Scheme = "https"
	cfg.AddDefaultHeader("Authorization", "Bearer "+r.token)
	cfg.HTTPClient = http.DefaultClient
	apiClient := openapi.NewAPIClient(cfg)
	reqParams := openapi.GetConnectionDetailsRequest{}
	reqParams.SetConnectionname(plan.ConnectionName.ValueString())
	existingResource, _, _ := apiClient.ConnectionsAPI.GetConnectionDetails(ctx).GetConnectionDetailsRequest(reqParams).Execute()
	if existingResource != nil &&
		existingResource.EntraIDConnectionResponse != nil &&
		existingResource.EntraIDConnectionResponse.Errorcode != nil &&
		*existingResource.EntraIDConnectionResponse.Errorcode == 0 {
		log.Printf("[ERROR] Connection name already exists. Please import or use a different name")
		resp.Diagnostics.AddError("API Create Failed", "Connection name already exists. Please import or use a different name")
		return
	}
	entraidConn := openapi.EntraIDConnector{
		BaseConnector: openapi.BaseConnector{
			//required fields
			Connectiontype: "AzureAD",
			ConnectionName: plan.ConnectionName.ValueString(),
			//optional fields
			Description:        util.StringPointerOrEmpty(plan.Description),
			Defaultsavroles:    util.StringPointerOrEmpty(plan.DefaultSavRoles),
			EmailTemplate:      util.StringPointerOrEmpty(plan.EmailTemplate),
			VaultConnection:    util.SafeStringConnector(plan.VaultConnection.ValueString()),
			VaultConfiguration: util.SafeStringConnector(plan.VaultConfiguration.ValueString()),
			Saveinvault:        util.SafeStringConnector(plan.SaveInVault.ValueString()),
		},
		//required fields
		CLIENT_ID:     config.ClientId.ValueString(),
		CLIENT_SECRET: config.ClientSecret.ValueString(),
		AAD_TENANT_ID: plan.AadTenantId.ValueString(),
		//optional fields
		ACCESS_TOKEN:                    util.StringPointerOrEmpty(config.AccessToken),
		AZURE_MGMT_ACCESS_TOKEN:         util.StringPointerOrEmpty(config.AzureMgmtAccessToken),
		AUTHENTICATION_ENDPOINT:         util.StringPointerOrEmpty(plan.AuthenticationEndpoint),
		MICROSOFT_GRAPH_ENDPOINT:        util.StringPointerOrEmpty(plan.MicrosoftGraphEndpoint),
		AZURE_MANAGEMENT_ENDPOINT:       util.StringPointerOrEmpty(plan.AzureManagementEndpoint),
		ImportUserJSON:                  util.StringPointerOrEmpty(plan.ImportUserJson),
		CREATEUSERS:                     util.StringPointerOrEmpty(plan.CreateUsers),
		WINDOWS_CONNECTOR_JSON:          util.StringPointerOrEmpty(config.WindowsConnectorJson),
		CREATE_NEW_ENDPOINTS:            util.StringPointerOrEmpty(plan.CreateNewEndpoints),
		MANAGED_ACCOUNT_TYPE:            util.StringPointerOrEmpty(plan.ManagedAccountType),
		ACCOUNT_ATTRIBUTES:              util.StringPointerOrEmpty(plan.AccountAttributes),
		SERVICE_ACCOUNT_ATTRIBUTES:      util.StringPointerOrEmpty(plan.ServiceAccountAttributes),
		DELTATOKENSJSON:                 util.StringPointerOrEmpty(plan.DeltaTokensJson),
		ACCOUNT_IMPORT_FIELDS:           util.StringPointerOrEmpty(plan.AccountImportFields),
		IMPORT_DEPTH:                    util.StringPointerOrEmpty(plan.ImportDepth),
		ENTITLEMENT_ATTRIBUTE:           util.StringPointerOrEmpty(plan.EntitlementAttribute),
		CreateAccountJSON:               util.StringPointerOrEmpty(plan.CreateAccountJson),
		UpdateAccountJSON:               util.StringPointerOrEmpty(plan.UpdateAccountJson),
		EnableAccountJSON:               util.StringPointerOrEmpty(plan.EnableAccountJson),
		DisableAccountJSON:              util.StringPointerOrEmpty(plan.DisableAccountJson),
		AddAccessJSON:                   util.StringPointerOrEmpty(plan.AddAccessJson),
		RemoveAccessJSON:                util.StringPointerOrEmpty(plan.RemoveAccessJson),
		UpdateUserJSON:                  util.StringPointerOrEmpty(plan.UpdateUserJson),
		ChangePassJSON:                  util.StringPointerOrEmpty(config.ChangePassJson),
		RemoveAccountJSON:               util.StringPointerOrEmpty(plan.RemoveAccountJson),
		ConnectionJSON:                  util.StringPointerOrEmpty(config.ConnectionJson),
		CreateGroupJSON:                 util.StringPointerOrEmpty(plan.CreateGroupJson),
		UpdateGroupJSON:                 util.StringPointerOrEmpty(plan.UpdateGroupJson),
		AddAccessToEntitlementJSON:      util.StringPointerOrEmpty(plan.AddAccessToEntitlementJson),
		RemoveAccessFromEntitlementJSON: util.StringPointerOrEmpty(plan.RemoveAccessFromEntitlementJson),
		DeleteGroupJSON:                 util.StringPointerOrEmpty(plan.DeleteGroupJson),
		CreateServicePrincipalJSON:      util.StringPointerOrEmpty(plan.CreateServicePrincipalJson),
		UpdateServicePrincipalJSON:      util.StringPointerOrEmpty(plan.UpdateServicePrincipalJson),
		RemoveServicePrincipalJSON:      util.StringPointerOrEmpty(plan.RemoveServicePrincipalJson),
		ENTITLEMENT_FILTER_JSON:         util.StringPointerOrEmpty(plan.EntitlementFilterJson),
		CreateTeamJSON:                  util.StringPointerOrEmpty(plan.CreateTeamJson),
		CreateChannelJSON:               util.StringPointerOrEmpty(plan.CreateChannelJson),
		STATUS_THRESHOLD_CONFIG:         util.StringPointerOrEmpty(plan.StatusThresholdConfig),
		ACCOUNTS_FILTER:                 util.StringPointerOrEmpty(plan.AccountsFilter),
		PAM_CONFIG:                      util.StringPointerOrEmpty(plan.PamConfig),
		ENDPOINTS_FILTER:                util.StringPointerOrEmpty(plan.EndpointsFilter),
		ConfigJSON:                      util.StringPointerOrEmpty(plan.ConfigJson),
		MODIFYUSERDATAJSON:              util.StringPointerOrEmpty(plan.ModifyUserdataJson),
		ENHANCEDDIRECTORYROLES:          util.StringPointerOrEmpty(plan.EnhancedDirectoryRoles),
	}

	entraidConnRequest := openapi.CreateOrUpdateRequest{
		EntraIDConnector: &entraidConn,
	}

	// Initialize API client
	apiResp, _, err := apiClient.ConnectionsAPI.CreateOrUpdate(ctx).CreateOrUpdateRequest(entraidConnRequest).Execute()
	if err != nil {
		log.Printf("[ERROR] Failed to create API resource. Error: %v", err)
		resp.Diagnostics.AddError("API Create Failed", fmt.Sprintf("Error: %v", err))
		return
	}
	if apiResp != nil && *apiResp.ErrorCode != "0" {
		log.Printf("[ERROR]: Error in creating EntraId connection resource. Errorcode: %v, Message: %v", *apiResp.ErrorCode, *apiResp.Msg)
		resp.Diagnostics.AddError("Creation of EntraId connection failed", *apiResp.Msg)
		return
	}

	plan.ID = types.StringValue(fmt.Sprintf("%d", *apiResp.ConnectionKey))
	plan.ConnectionKey = types.Int64Value(int64(*apiResp.ConnectionKey))
	plan.Description = util.SafeStringDatasource(plan.Description.ValueStringPointer())
	plan.DefaultSavRoles = util.SafeStringDatasource(plan.DefaultSavRoles.ValueStringPointer())
	plan.EmailTemplate = util.SafeStringDatasource(plan.EmailTemplate.ValueStringPointer())
	plan.AuthenticationEndpoint = util.SafeStringDatasource(plan.AuthenticationEndpoint.ValueStringPointer())
	plan.MicrosoftGraphEndpoint = util.SafeStringDatasource(plan.MicrosoftGraphEndpoint.ValueStringPointer())
	plan.AzureManagementEndpoint = util.SafeStringDatasource(plan.AzureManagementEndpoint.ValueStringPointer())
	plan.ImportUserJson = util.SafeStringDatasource(plan.ImportUserJson.ValueStringPointer())
	plan.CreateUsers = util.SafeStringDatasource(plan.CreateUsers.ValueStringPointer())
	plan.CreateNewEndpoints = util.SafeStringDatasource(plan.CreateNewEndpoints.ValueStringPointer())
	plan.ManagedAccountType = util.SafeStringDatasource(plan.ManagedAccountType.ValueStringPointer())
	plan.AccountAttributes = util.SafeStringDatasource(plan.AccountAttributes.ValueStringPointer())
	plan.ServiceAccountAttributes = util.SafeStringDatasource(plan.ServiceAccountAttributes.ValueStringPointer())
	plan.DeltaTokensJson = util.SafeStringDatasource(plan.DeltaTokensJson.ValueStringPointer())
	plan.AccountImportFields = util.SafeStringDatasource(plan.AccountImportFields.ValueStringPointer())
	plan.ImportDepth = util.SafeStringDatasource(plan.ImportDepth.ValueStringPointer())
	plan.EntitlementAttribute = util.SafeStringDatasource(plan.EntitlementAttribute.ValueStringPointer())
	plan.CreateAccountJson = util.SafeStringDatasource(plan.CreateAccountJson.ValueStringPointer())
	plan.UpdateAccountJson = util.SafeStringDatasource(plan.UpdateAccountJson.ValueStringPointer())
	plan.EnableAccountJson = util.SafeStringDatasource(plan.EnableAccountJson.ValueStringPointer())
	plan.DisableAccountJson = util.SafeStringDatasource(plan.DisableAccountJson.ValueStringPointer())
	plan.AddAccessJson = util.SafeStringDatasource(plan.AddAccessJson.ValueStringPointer())
	plan.RemoveAccessJson = util.SafeStringDatasource(plan.RemoveAccessJson.ValueStringPointer())
	plan.UpdateUserJson = util.SafeStringDatasource(plan.UpdateUserJson.ValueStringPointer())
	plan.ChangePassJson = util.SafeStringDatasource(plan.ChangePassJson.ValueStringPointer())
	plan.RemoveAccountJson = util.SafeStringDatasource(plan.RemoveAccountJson.ValueStringPointer())
	plan.CreateGroupJson = util.SafeStringDatasource(plan.CreateGroupJson.ValueStringPointer())
	plan.UpdateGroupJson = util.SafeStringDatasource(plan.UpdateGroupJson.ValueStringPointer())
	plan.AddAccessToEntitlementJson = util.SafeStringDatasource(plan.AddAccessToEntitlementJson.ValueStringPointer())
	plan.RemoveAccessFromEntitlementJson = util.SafeStringDatasource(plan.RemoveAccessFromEntitlementJson.ValueStringPointer())
	plan.DeleteGroupJson = util.SafeStringDatasource(plan.DeleteGroupJson.ValueStringPointer())
	plan.CreateServicePrincipalJson = util.SafeStringDatasource(plan.CreateServicePrincipalJson.ValueStringPointer())
	plan.UpdateServicePrincipalJson = util.SafeStringDatasource(plan.UpdateServicePrincipalJson.ValueStringPointer())
	plan.RemoveServicePrincipalJson = util.SafeStringDatasource(plan.RemoveServicePrincipalJson.ValueStringPointer())
	plan.EntitlementFilterJson = util.SafeStringDatasource(plan.EntitlementFilterJson.ValueStringPointer())
	plan.CreateTeamJson = util.SafeStringDatasource(plan.CreateTeamJson.ValueStringPointer())
	plan.CreateChannelJson = util.SafeStringDatasource(plan.CreateChannelJson.ValueStringPointer())
	plan.StatusThresholdConfig = util.SafeStringDatasource(plan.StatusThresholdConfig.ValueStringPointer())
	plan.AccountsFilter = util.SafeStringDatasource(plan.AccountsFilter.ValueStringPointer())
	plan.PamConfig = util.SafeStringDatasource(plan.PamConfig.ValueStringPointer())
	plan.EndpointsFilter = util.SafeStringDatasource(plan.EndpointsFilter.ValueStringPointer())
	plan.ConfigJson = util.SafeStringDatasource(plan.ConfigJson.ValueStringPointer())
	plan.ModifyUserdataJson = util.SafeStringDatasource(plan.ModifyUserdataJson.ValueStringPointer())
	plan.EnhancedDirectoryRoles = util.SafeStringDatasource(plan.EnhancedDirectoryRoles.ValueStringPointer())
	plan.Msg = types.StringValue(util.SafeDeref(apiResp.Msg))
	plan.ErrorCode = types.StringValue(util.SafeDeref(apiResp.ErrorCode))
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *entraIdConnectionResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state EntraIdConnectorResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Configure API client
	cfg := openapi.NewConfiguration()
	apiBaseURL := strings.TrimPrefix(strings.TrimPrefix(r.client.APIBaseURL(), "https://"), "http://")
	cfg.Host = apiBaseURL
	cfg.Scheme = "https"
	cfg.AddDefaultHeader("Authorization", "Bearer "+r.token)
	cfg.HTTPClient = http.DefaultClient

	apiClient := openapi.NewAPIClient(cfg)
	reqParams := openapi.GetConnectionDetailsRequest{}

	reqParams.SetConnectionname(state.ConnectionName.ValueString())
	apiResp, _, err := apiClient.ConnectionsAPI.GetConnectionDetails(ctx).GetConnectionDetailsRequest(reqParams).Execute()
	if err != nil {
		log.Printf("Problem with the get function in read block")
		resp.Diagnostics.AddError("API Read Failed In Read Block", fmt.Sprintf("Error: %v", err))
		return
	}
	if apiResp != nil && apiResp.EntraIDConnectionResponse != nil && *apiResp.EntraIDConnectionResponse.Errorcode != 0 {
		log.Printf("[ERROR]: Error in reading EntraId connection. Errorcode: %v, Message: %v", *apiResp.EntraIDConnectionResponse.Errorcode, *apiResp.EntraIDConnectionResponse.Msg)
		resp.Diagnostics.AddError("Read EntraId connection failed", *apiResp.EntraIDConnectionResponse.Msg)
		return
	}

	state.ConnectionKey = types.Int64Value(int64(*apiResp.EntraIDConnectionResponse.Connectionkey))
	state.ID = types.StringValue(fmt.Sprintf("%d", *apiResp.EntraIDConnectionResponse.Connectionkey))
	state.ConnectionName = util.SafeStringDatasource(apiResp.EntraIDConnectionResponse.Connectionname)
	state.ConnectionKey = util.SafeInt64(apiResp.EntraIDConnectionResponse.Connectionkey)
	state.Description = util.SafeStringDatasource(apiResp.EntraIDConnectionResponse.Description)
	state.DefaultSavRoles = util.SafeStringDatasource(apiResp.EntraIDConnectionResponse.Defaultsavroles)
	state.EmailTemplate = util.SafeStringDatasource(apiResp.EntraIDConnectionResponse.Emailtemplate)
	state.UpdateUserJson = util.SafeStringDatasource(apiResp.EntraIDConnectionResponse.Connectionattributes.UpdateUserJSON)
	state.MicrosoftGraphEndpoint = util.SafeStringDatasource(apiResp.EntraIDConnectionResponse.Connectionattributes.MICROSOFT_GRAPH_ENDPOINT)
	state.EndpointsFilter = util.SafeStringDatasource(apiResp.EntraIDConnectionResponse.Connectionattributes.ENDPOINTS_FILTER)
	state.ImportUserJson = util.SafeStringDatasource(apiResp.EntraIDConnectionResponse.Connectionattributes.ImportUserJSON)
	state.EnableAccountJson = util.SafeStringDatasource(apiResp.EntraIDConnectionResponse.Connectionattributes.EnableAccountJSON)
	state.ClientId = util.SafeStringDatasource(apiResp.EntraIDConnectionResponse.Connectionattributes.CLIENT_ID)
	state.DeleteGroupJson = util.SafeStringDatasource(apiResp.EntraIDConnectionResponse.Connectionattributes.DeleteGroupJSON)
	state.ConfigJson = util.SafeStringDatasource(apiResp.EntraIDConnectionResponse.Connectionattributes.ConfigJSON)
	state.AddAccessJson = util.SafeStringDatasource(apiResp.EntraIDConnectionResponse.Connectionattributes.AddAccessJSON)
	state.CreateChannelJson = util.SafeStringDatasource(apiResp.EntraIDConnectionResponse.Connectionattributes.CreateChannelJSON)
	state.UpdateAccountJson = util.SafeStringDatasource(apiResp.EntraIDConnectionResponse.Connectionattributes.UpdateAccountJSON)
	state.RemoveServicePrincipalJson = util.SafeStringDatasource(apiResp.EntraIDConnectionResponse.Connectionattributes.RemoveServicePrincipalJSON)
	state.ImportDepth = util.SafeStringDatasource(apiResp.EntraIDConnectionResponse.Connectionattributes.IMPORT_DEPTH)
	state.CreateAccountJson = util.SafeStringDatasource(apiResp.EntraIDConnectionResponse.Connectionattributes.CreateAccountJSON)
	state.PamConfig = util.SafeStringDatasource(apiResp.EntraIDConnectionResponse.Connectionattributes.PAM_CONFIG)
	state.UpdateServicePrincipalJson = util.SafeStringDatasource(apiResp.EntraIDConnectionResponse.Connectionattributes.UpdateServicePrincipalJSON)
	state.AzureManagementEndpoint = util.SafeStringDatasource(apiResp.EntraIDConnectionResponse.Connectionattributes.AZURE_MANAGEMENT_ENDPOINT)
	state.EntitlementAttribute = util.SafeStringDatasource(apiResp.EntraIDConnectionResponse.Connectionattributes.ENTITLEMENT_ATTRIBUTE)
	state.AccountsFilter = util.SafeStringDatasource(apiResp.EntraIDConnectionResponse.Connectionattributes.ACCOUNTS_FILTER)
	state.DeltaTokensJson = util.SafeStringDatasource(apiResp.EntraIDConnectionResponse.Connectionattributes.DELTATOKENSJSON)
	state.CreateTeamJson = util.SafeStringDatasource(apiResp.EntraIDConnectionResponse.Connectionattributes.CreateTeamJSON)
	state.EnhancedDirectoryRoles = util.SafeStringDatasource(apiResp.EntraIDConnectionResponse.Connectionattributes.ENHANCEDDIRECTORYROLES)
	state.StatusThresholdConfig = util.SafeStringDatasource(apiResp.EntraIDConnectionResponse.Connectionattributes.STATUS_THRESHOLD_CONFIG)
	state.AccountImportFields = util.SafeStringDatasource(apiResp.EntraIDConnectionResponse.Connectionattributes.ACCOUNT_IMPORT_FIELDS)
	state.RemoveAccountJson = util.SafeStringDatasource(apiResp.EntraIDConnectionResponse.Connectionattributes.RemoveAccountJSON)
	state.ChangePassJson = util.SafeStringDatasource(apiResp.EntraIDConnectionResponse.Connectionattributes.ChangePassJSON)
	state.EntitlementFilterJson = util.SafeStringDatasource(apiResp.EntraIDConnectionResponse.Connectionattributes.ENTITLEMENT_FILTER_JSON)
	state.ServiceAccountAttributes = util.SafeStringDatasource(apiResp.EntraIDConnectionResponse.Connectionattributes.SERVICE_ACCOUNT_ATTRIBUTES)
	state.AddAccessToEntitlementJson = util.SafeStringDatasource(apiResp.EntraIDConnectionResponse.Connectionattributes.AddAccessToEntitlementJSON)
	state.AuthenticationEndpoint = util.SafeStringDatasource(apiResp.EntraIDConnectionResponse.Connectionattributes.AUTHENTICATION_ENDPOINT)
	state.CreateServicePrincipalJson = util.SafeStringDatasource(apiResp.EntraIDConnectionResponse.Connectionattributes.CreateServicePrincipalJSON)
	state.ModifyUserdataJson = util.SafeStringDatasource(apiResp.EntraIDConnectionResponse.Connectionattributes.MODIFYUSERDATAJSON)
	state.RemoveAccessJson = util.SafeStringDatasource(apiResp.EntraIDConnectionResponse.Connectionattributes.RemoveAccessJSON)
	state.CreateUsers = util.SafeStringDatasource(apiResp.EntraIDConnectionResponse.Connectionattributes.CREATEUSERS)
	state.RemoveAccessFromEntitlementJson = util.SafeStringDatasource(apiResp.EntraIDConnectionResponse.Connectionattributes.RemoveAccessFromEntitlementJSON)
	state.DisableAccountJson = util.SafeStringDatasource(apiResp.EntraIDConnectionResponse.Connectionattributes.DisableAccountJSON)
	state.CreateNewEndpoints = util.SafeStringDatasource(apiResp.EntraIDConnectionResponse.Connectionattributes.CREATE_NEW_ENDPOINTS)
	state.ManagedAccountType = util.SafeStringDatasource(apiResp.EntraIDConnectionResponse.Connectionattributes.MANAGED_ACCOUNT_TYPE)
	state.AccountAttributes = util.SafeStringDatasource(apiResp.EntraIDConnectionResponse.Connectionattributes.ACCOUNT_ATTRIBUTES)
	state.AadTenantId = util.SafeStringDatasource(apiResp.EntraIDConnectionResponse.Connectionattributes.AAD_TENANT_ID)
	state.UpdateGroupJson = util.SafeStringDatasource(apiResp.EntraIDConnectionResponse.Connectionattributes.UpdateGroupJSON)
	state.CreateGroupJson = util.SafeStringDatasource(apiResp.EntraIDConnectionResponse.Connectionattributes.CreateGroupJSON)
	apiMessage := util.SafeDeref(apiResp.EntraIDConnectionResponse.Msg)
	if apiMessage == "success" {
		state.Msg = types.StringValue("Connection Successful")
	} else {
		state.Msg = types.StringValue(apiMessage)
	}
	state.ErrorCode = util.Int32PtrToTFString(apiResp.EntraIDConnectionResponse.Errorcode)
	stateDiagnostics := resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(stateDiagnostics...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *entraIdConnectionResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state, config EntraIdConnectorResourceModel
	// Extract state from request
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
	// Extract plan from request
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}
	//Extract config from request
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	cfg := openapi.NewConfiguration()
	apiBaseURL := strings.TrimPrefix(strings.TrimPrefix(r.client.APIBaseURL(), "https://"), "http://")
	cfg.Host = apiBaseURL
	cfg.Scheme = "https"
	cfg.AddDefaultHeader("Authorization", "Bearer "+r.token)
	if plan.ConnectionName.ValueString() != state.ConnectionName.ValueString() {
		resp.Diagnostics.AddError("Error", "Connection name cannot be updated")
		log.Printf("[ERROR]: Connection name cannot be updated")
		return
	}

	cfg.HTTPClient = http.DefaultClient
	entraidConn := openapi.EntraIDConnector{
		BaseConnector: openapi.BaseConnector{
			//required fields
			Connectiontype: "AzureAD",
			ConnectionName: plan.ConnectionName.ValueString(),
			//optional fields
			Description:        util.StringPointerOrEmpty(plan.Description),
			Defaultsavroles:    util.StringPointerOrEmpty(plan.DefaultSavRoles),
			EmailTemplate:      util.StringPointerOrEmpty(plan.EmailTemplate),
			VaultConnection:    util.SafeStringConnector(plan.VaultConnection.ValueString()),
			VaultConfiguration: util.SafeStringConnector(plan.VaultConfiguration.ValueString()),
			Saveinvault:        util.SafeStringConnector(plan.SaveInVault.ValueString()),
		},
		//required fields
		CLIENT_ID:     config.ClientId.ValueString(),
		CLIENT_SECRET: config.ClientSecret.ValueString(),
		AAD_TENANT_ID: plan.AadTenantId.ValueString(),
		//optional fields
		ACCESS_TOKEN:                    util.StringPointerOrEmpty(config.AccessToken),
		AZURE_MGMT_ACCESS_TOKEN:         util.StringPointerOrEmpty(config.AzureMgmtAccessToken),
		AUTHENTICATION_ENDPOINT:         util.StringPointerOrEmpty(plan.AuthenticationEndpoint),
		MICROSOFT_GRAPH_ENDPOINT:        util.StringPointerOrEmpty(plan.MicrosoftGraphEndpoint),
		AZURE_MANAGEMENT_ENDPOINT:       util.StringPointerOrEmpty(plan.AzureManagementEndpoint),
		ImportUserJSON:                  util.StringPointerOrEmpty(plan.ImportUserJson),
		CREATEUSERS:                     util.StringPointerOrEmpty(plan.CreateUsers),
		WINDOWS_CONNECTOR_JSON:          util.StringPointerOrEmpty(plan.WindowsConnectorJson),
		CREATE_NEW_ENDPOINTS:            util.StringPointerOrEmpty(plan.CreateNewEndpoints),
		MANAGED_ACCOUNT_TYPE:            util.StringPointerOrEmpty(plan.ManagedAccountType),
		ACCOUNT_ATTRIBUTES:              util.StringPointerOrEmpty(plan.AccountAttributes),
		SERVICE_ACCOUNT_ATTRIBUTES:      util.StringPointerOrEmpty(plan.ServiceAccountAttributes),
		DELTATOKENSJSON:                 util.StringPointerOrEmpty(plan.DeltaTokensJson),
		ACCOUNT_IMPORT_FIELDS:           util.StringPointerOrEmpty(plan.AccountImportFields),
		IMPORT_DEPTH:                    util.StringPointerOrEmpty(plan.ImportDepth),
		ENTITLEMENT_ATTRIBUTE:           util.StringPointerOrEmpty(plan.EntitlementAttribute),
		CreateAccountJSON:               util.StringPointerOrEmpty(plan.CreateAccountJson),
		UpdateAccountJSON:               util.StringPointerOrEmpty(plan.UpdateAccountJson),
		EnableAccountJSON:               util.StringPointerOrEmpty(plan.EnableAccountJson),
		DisableAccountJSON:              util.StringPointerOrEmpty(plan.DisableAccountJson),
		AddAccessJSON:                   util.StringPointerOrEmpty(plan.AddAccessJson),
		RemoveAccessJSON:                util.StringPointerOrEmpty(plan.RemoveAccessJson),
		UpdateUserJSON:                  util.StringPointerOrEmpty(plan.UpdateUserJson),
		ChangePassJSON:                  util.StringPointerOrEmpty(config.ChangePassJson),
		RemoveAccountJSON:               util.StringPointerOrEmpty(plan.RemoveAccountJson),
		ConnectionJSON:                  util.StringPointerOrEmpty(config.ConnectionJson),
		CreateGroupJSON:                 util.StringPointerOrEmpty(plan.CreateGroupJson),
		UpdateGroupJSON:                 util.StringPointerOrEmpty(plan.UpdateGroupJson),
		AddAccessToEntitlementJSON:      util.StringPointerOrEmpty(plan.AddAccessToEntitlementJson),
		RemoveAccessFromEntitlementJSON: util.StringPointerOrEmpty(plan.RemoveAccessFromEntitlementJson),
		DeleteGroupJSON:                 util.StringPointerOrEmpty(plan.DeleteGroupJson),
		CreateServicePrincipalJSON:      util.StringPointerOrEmpty(plan.CreateServicePrincipalJson),
		UpdateServicePrincipalJSON:      util.StringPointerOrEmpty(plan.UpdateServicePrincipalJson),
		RemoveServicePrincipalJSON:      util.StringPointerOrEmpty(plan.RemoveServicePrincipalJson),
		ENTITLEMENT_FILTER_JSON:         util.StringPointerOrEmpty(plan.EntitlementFilterJson),
		CreateTeamJSON:                  util.StringPointerOrEmpty(plan.CreateTeamJson),
		CreateChannelJSON:               util.StringPointerOrEmpty(plan.CreateChannelJson),
		STATUS_THRESHOLD_CONFIG:         util.StringPointerOrEmpty(plan.StatusThresholdConfig),
		ACCOUNTS_FILTER:                 util.StringPointerOrEmpty(plan.AccountsFilter),
		PAM_CONFIG:                      util.StringPointerOrEmpty(plan.PamConfig),
		ENDPOINTS_FILTER:                util.StringPointerOrEmpty(plan.EndpointsFilter),
		ConfigJSON:                      util.StringPointerOrEmpty(plan.ConfigJson),
		MODIFYUSERDATAJSON:              util.StringPointerOrEmpty(plan.ModifyUserdataJson),
		ENHANCEDDIRECTORYROLES:          util.StringPointerOrEmpty(plan.EnhancedDirectoryRoles),
	}

	entraidConnRequest := openapi.CreateOrUpdateRequest{
		EntraIDConnector: &entraidConn,
	}

	// Initialize API client
	apiClient := openapi.NewAPIClient(cfg)

	apiResp, _, err := apiClient.ConnectionsAPI.CreateOrUpdate(ctx).CreateOrUpdateRequest(entraidConnRequest).Execute()
	if err != nil || *apiResp.ErrorCode != "0" {
		log.Printf("Problem with the update function")
		resp.Diagnostics.AddError("API Update Failed", fmt.Sprintf("Error: %v", err))
		return
	}
	if apiResp != nil && *apiResp.ErrorCode != "0" {
		log.Printf("[ERROR]: Error in updation EntraId connection. Errorcode: %v, Message: %v", *apiResp.ErrorCode, *apiResp.Msg)
		resp.Diagnostics.AddError("Updation of EntraId connection failed", *apiResp.Msg)
		return
	}

	reqParams := openapi.GetConnectionDetailsRequest{}

	reqParams.SetConnectionname(plan.ConnectionName.ValueString())
	getResp, _, err := apiClient.ConnectionsAPI.GetConnectionDetails(ctx).GetConnectionDetailsRequest(reqParams).Execute()
	if err != nil {
		log.Printf("Problem with the get function in update block")
		resp.Diagnostics.AddError("API Read Failed In Update Block", fmt.Sprintf("Error: %v", *getResp.EntraIDConnectionResponse.Msg))
		return
	}
	if getResp != nil && getResp.EntraIDConnectionResponse != nil && *getResp.EntraIDConnectionResponse.Errorcode != 0 {
		log.Printf("[ERROR]: Error in reading EntraId connection after updation. Errorcode: %v, Message: %v", *getResp.EntraIDConnectionResponse.Errorcode, *getResp.EntraIDConnectionResponse.Msg)
		resp.Diagnostics.AddError("Read EntraId connection after updatio failed", *getResp.EntraIDConnectionResponse.Msg)
		return
	}

	plan.ConnectionKey = types.Int64Value(int64(*getResp.EntraIDConnectionResponse.Connectionkey))
	plan.ID = types.StringValue(fmt.Sprintf("%d", *getResp.EntraIDConnectionResponse.Connectionkey))
	plan.ConnectionName = util.SafeStringDatasource(getResp.EntraIDConnectionResponse.Connectionname)
	plan.ConnectionKey = util.SafeInt64(getResp.EntraIDConnectionResponse.Connectionkey)
	plan.Description = util.SafeStringDatasource(getResp.EntraIDConnectionResponse.Description)
	plan.DefaultSavRoles = util.SafeStringDatasource(getResp.EntraIDConnectionResponse.Defaultsavroles)
	plan.EmailTemplate = util.SafeStringDatasource(getResp.EntraIDConnectionResponse.Emailtemplate)
	plan.UpdateUserJson = util.SafeStringDatasource(getResp.EntraIDConnectionResponse.Connectionattributes.UpdateUserJSON)
	plan.MicrosoftGraphEndpoint = util.SafeStringDatasource(getResp.EntraIDConnectionResponse.Connectionattributes.MICROSOFT_GRAPH_ENDPOINT)
	plan.EndpointsFilter = util.SafeStringDatasource(getResp.EntraIDConnectionResponse.Connectionattributes.ENDPOINTS_FILTER)
	plan.ImportUserJson = util.SafeStringDatasource(getResp.EntraIDConnectionResponse.Connectionattributes.ImportUserJSON)
	plan.EnableAccountJson = util.SafeStringDatasource(getResp.EntraIDConnectionResponse.Connectionattributes.EnableAccountJSON)
	plan.ClientId = util.SafeStringDatasource(getResp.EntraIDConnectionResponse.Connectionattributes.CLIENT_ID)
	plan.DeleteGroupJson = util.SafeStringDatasource(getResp.EntraIDConnectionResponse.Connectionattributes.DeleteGroupJSON)
	plan.ConfigJson = util.SafeStringDatasource(getResp.EntraIDConnectionResponse.Connectionattributes.ConfigJSON)
	plan.AddAccessJson = util.SafeStringDatasource(getResp.EntraIDConnectionResponse.Connectionattributes.AddAccessJSON)
	plan.CreateChannelJson = util.SafeStringDatasource(getResp.EntraIDConnectionResponse.Connectionattributes.CreateChannelJSON)
	plan.UpdateAccountJson = util.SafeStringDatasource(getResp.EntraIDConnectionResponse.Connectionattributes.UpdateAccountJSON)
	plan.RemoveServicePrincipalJson = util.SafeStringDatasource(getResp.EntraIDConnectionResponse.Connectionattributes.RemoveServicePrincipalJSON)
	plan.ImportDepth = util.SafeStringDatasource(getResp.EntraIDConnectionResponse.Connectionattributes.IMPORT_DEPTH)
	plan.CreateAccountJson = util.SafeStringDatasource(getResp.EntraIDConnectionResponse.Connectionattributes.CreateAccountJSON)
	plan.PamConfig = util.SafeStringDatasource(getResp.EntraIDConnectionResponse.Connectionattributes.PAM_CONFIG)
	plan.UpdateServicePrincipalJson = util.SafeStringDatasource(getResp.EntraIDConnectionResponse.Connectionattributes.UpdateServicePrincipalJSON)
	plan.AzureManagementEndpoint = util.SafeStringDatasource(getResp.EntraIDConnectionResponse.Connectionattributes.AZURE_MANAGEMENT_ENDPOINT)
	plan.EntitlementAttribute = util.SafeStringDatasource(getResp.EntraIDConnectionResponse.Connectionattributes.ENTITLEMENT_ATTRIBUTE)
	plan.AccountsFilter = util.SafeStringDatasource(getResp.EntraIDConnectionResponse.Connectionattributes.ACCOUNTS_FILTER)
	plan.DeltaTokensJson = util.SafeStringDatasource(getResp.EntraIDConnectionResponse.Connectionattributes.DELTATOKENSJSON)
	plan.CreateTeamJson = util.SafeStringDatasource(getResp.EntraIDConnectionResponse.Connectionattributes.CreateTeamJSON)
	plan.EnhancedDirectoryRoles = util.SafeStringDatasource(getResp.EntraIDConnectionResponse.Connectionattributes.ENHANCEDDIRECTORYROLES)
	plan.StatusThresholdConfig = util.SafeStringDatasource(getResp.EntraIDConnectionResponse.Connectionattributes.STATUS_THRESHOLD_CONFIG)
	plan.AccountImportFields = util.SafeStringDatasource(getResp.EntraIDConnectionResponse.Connectionattributes.ACCOUNT_IMPORT_FIELDS)
	plan.RemoveAccountJson = util.SafeStringDatasource(getResp.EntraIDConnectionResponse.Connectionattributes.RemoveAccountJSON)
	plan.ChangePassJson = util.SafeStringDatasource(getResp.EntraIDConnectionResponse.Connectionattributes.ChangePassJSON)
	plan.EntitlementFilterJson = util.SafeStringDatasource(getResp.EntraIDConnectionResponse.Connectionattributes.ENTITLEMENT_FILTER_JSON)
	plan.ServiceAccountAttributes = util.SafeStringDatasource(getResp.EntraIDConnectionResponse.Connectionattributes.SERVICE_ACCOUNT_ATTRIBUTES)
	plan.AddAccessToEntitlementJson = util.SafeStringDatasource(getResp.EntraIDConnectionResponse.Connectionattributes.AddAccessToEntitlementJSON)
	plan.AuthenticationEndpoint = util.SafeStringDatasource(getResp.EntraIDConnectionResponse.Connectionattributes.AUTHENTICATION_ENDPOINT)
	plan.CreateServicePrincipalJson = util.SafeStringDatasource(getResp.EntraIDConnectionResponse.Connectionattributes.CreateServicePrincipalJSON)
	plan.ModifyUserdataJson = util.SafeStringDatasource(getResp.EntraIDConnectionResponse.Connectionattributes.MODIFYUSERDATAJSON)
	plan.RemoveAccessJson = util.SafeStringDatasource(getResp.EntraIDConnectionResponse.Connectionattributes.RemoveAccessJSON)
	plan.CreateUsers = util.SafeStringDatasource(getResp.EntraIDConnectionResponse.Connectionattributes.CREATEUSERS)
	plan.RemoveAccessFromEntitlementJson = util.SafeStringDatasource(getResp.EntraIDConnectionResponse.Connectionattributes.RemoveAccessFromEntitlementJSON)
	plan.DisableAccountJson = util.SafeStringDatasource(getResp.EntraIDConnectionResponse.Connectionattributes.DisableAccountJSON)
	plan.CreateNewEndpoints = util.SafeStringDatasource(getResp.EntraIDConnectionResponse.Connectionattributes.CREATE_NEW_ENDPOINTS)
	plan.ManagedAccountType = util.SafeStringDatasource(getResp.EntraIDConnectionResponse.Connectionattributes.MANAGED_ACCOUNT_TYPE)
	plan.AccountAttributes = util.SafeStringDatasource(getResp.EntraIDConnectionResponse.Connectionattributes.ACCOUNT_ATTRIBUTES)
	plan.AadTenantId = util.SafeStringDatasource(getResp.EntraIDConnectionResponse.Connectionattributes.AAD_TENANT_ID)
	plan.UpdateGroupJson = util.SafeStringDatasource(getResp.EntraIDConnectionResponse.Connectionattributes.UpdateGroupJSON)
	plan.CreateGroupJson = util.SafeStringDatasource(getResp.EntraIDConnectionResponse.Connectionattributes.CreateGroupJSON)
	apiMessage := util.SafeDeref(getResp.EntraIDConnectionResponse.Msg)
	if apiMessage == "success" {
		plan.Msg = types.StringValue("Connection Successful")
	} else {
		plan.Msg = types.StringValue(apiMessage)
	}
	plan.ErrorCode = util.Int32PtrToTFString(getResp.EntraIDConnectionResponse.Errorcode)
	stateUpdateDiagnostics := resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(stateUpdateDiagnostics...)
}
func (r *entraIdConnectionResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// resp.State.RemoveResource(ctx)
	if os.Getenv("TF_ACC") == "1" {
		resp.State.RemoveResource(ctx)
		return
	}
	resp.Diagnostics.AddError(
		"Delete Not Supported",
		"Resource deletion is not supported by this provider. Please remove the resource manually if required, or contact your administrator.",
	)
}

func (r *entraIdConnectionResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	// Retrieve import ID and save to id attribute
	resource.ImportStatePassthroughID(ctx, path.Root("connection_name"), req, resp)
}
