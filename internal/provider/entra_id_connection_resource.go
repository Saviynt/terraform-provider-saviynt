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
	"net/http"
	"os"
	"terraform-provider-Saviynt/internal/client"
	"terraform-provider-Saviynt/util"

	connectionsutil "terraform-provider-Saviynt/util/connectionsutil"
	"terraform-provider-Saviynt/util/errorsutil"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	openapi "github.com/saviynt/saviynt-api-go-client/connections"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &EntraIdConnectionResource{}
var _ resource.ResourceWithImportState = &EntraIdConnectionResource{}

// Initialize error codes for EntraID Connection operations
var entraIdErrorCodes = errorsutil.NewConnectorErrorCodes(errorsutil.ConnectorTypeEntraID)

type EntraIdConnectorResourceModel struct {
	BaseConnectorResourceModel
	ID                              types.String `tfsdk:"id"`
	ClientId                        types.String `tfsdk:"client_id"`
	ClientSecret                    types.String `tfsdk:"client_secret"`
	ClientSecretWO                  types.String `tfsdk:"client_secret_wo"`
	AccessToken                     types.String `tfsdk:"access_token"`
	AccessTokenWO                   types.String `tfsdk:"access_token_wo"`
	AadTenantId                     types.String `tfsdk:"aad_tenant_id"`
	AzureMgmtAccessToken            types.String `tfsdk:"azure_mgmt_access_token"`
	AzureMgmtAccessTokenWO          types.String `tfsdk:"azure_mgmt_access_token_wo"`
	AuthenticationEndpoint          types.String `tfsdk:"authentication_endpoint"`
	MicrosoftGraphEndpoint          types.String `tfsdk:"microsoft_graph_endpoint"`
	AzureManagementEndpoint         types.String `tfsdk:"azure_management_endpoint"`
	ImportUserJson                  types.String `tfsdk:"import_user_json"`
	CreateUsers                     types.String `tfsdk:"create_users"`
	WindowsConnectorJson            types.String `tfsdk:"windows_connector_json"`
	WindowsConnectorJsonWO          types.String `tfsdk:"windows_connector_json_wo"`
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
	ConnectionJsonWO                types.String `tfsdk:"connection_json_wo"`
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

type EntraIdConnectionResource struct {
	client            client.SaviyntClientInterface
	token             string
	provider          client.SaviyntProviderInterface
	connectionFactory client.ConnectionFactoryInterface
}

func NewEntraIdConnectionResource() resource.Resource {
	return &EntraIdConnectionResource{
		connectionFactory: &client.DefaultConnectionFactory{},
	}
}

func NewEntraIdConnectionResourceWithFactory(factory client.ConnectionFactoryInterface) resource.Resource {
	return &EntraIdConnectionResource{
		connectionFactory: factory,
	}
}

func (r *EntraIdConnectionResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
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
			Description: "Client ID for authentication.",
		},
		"client_secret": schema.StringAttribute{
			Optional:    true,
			Sensitive:   true,
			Description: "Client Secret for authentication.",
			Validators: []validator.String{
				stringvalidator.ConflictsWith(path.MatchRoot("client_secret_wo")),
			},
		},
		"client_secret_wo": schema.StringAttribute{
			Optional:    true,
			WriteOnly:   true,
			Description: "Client Secret for authentication (write-only).",
			Validators: []validator.String{
				stringvalidator.ConflictsWith(path.MatchRoot("client_secret")),
			},
		},
		"access_token": schema.StringAttribute{
			Optional:    true,
			Sensitive:   true,
			Description: "Access token used for API calls.",
			Validators: []validator.String{
				stringvalidator.ConflictsWith(path.MatchRoot("access_token_wo")),
			},
		},
		"access_token_wo": schema.StringAttribute{
			Optional:    true,
			WriteOnly:   true,
			Description: "Access token used for API calls (write-only).",
			Validators: []validator.String{
				stringvalidator.ConflictsWith(path.MatchRoot("access_token")),
			},
		},
		"aad_tenant_id": schema.StringAttribute{
			Required:    true,
			Description: "Azure Active Directory tenant ID.",
		},
		"azure_mgmt_access_token": schema.StringAttribute{
			Optional:    true,
			Sensitive:   true,
			Description: "Access token for Azure management APIs.",
			Validators: []validator.String{
				stringvalidator.ConflictsWith(path.MatchRoot("azure_mgmt_access_token_wo")),
			},
		},
		"azure_mgmt_access_token_wo": schema.StringAttribute{
			Optional:    true,
			WriteOnly:   true,
			Description: "Access token for Azure management APIs (write-only).",
			Validators: []validator.String{
				stringvalidator.ConflictsWith(path.MatchRoot("azure_mgmt_access_token")),
			},
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
			Sensitive:   true,
			Description: "Windows connector JSON configuration.",
			Validators: []validator.String{
				stringvalidator.ConflictsWith(path.MatchRoot("windows_connector_json_wo")),
			},
		},
		"windows_connector_json_wo": schema.StringAttribute{
			Optional:    true,
			WriteOnly:   true,
			Description: "Windows connector JSON configuration (write-only).",
			Validators: []validator.String{
				stringvalidator.ConflictsWith(path.MatchRoot("windows_connector_json")),
			},
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
			Computed:    true,
			Description: "JSON template to change password.",
		},
		"remove_account_json": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "JSON template to remove account.",
		},
		"connection_json": schema.StringAttribute{
			Optional:    true,
			Sensitive:   true,
			Description: "Configuration for the connection in JSON format. Either the connection_json field or the connection_json_wo field must be populated to set the connection_json attribute.",
			Validators: []validator.String{
				stringvalidator.ConflictsWith(path.MatchRoot("connection_json_wo")),
			},
		},
		"connection_json_wo": schema.StringAttribute{
			Optional:    true,
			WriteOnly:   true,
			Description: "Connection JSON configuration (write-only). Either the connection_json field or the connection_json_wo field must be populated to set the connection_json attribute.",
			Validators: []validator.String{
				stringvalidator.ConflictsWith(path.MatchRoot("connection_json")),
			},
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

func (r *EntraIdConnectionResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: util.EntraIDConnDescription,
		Attributes:  connectionsutil.MergeResourceAttributes(BaseConnectorResourceSchema(), EntraIdConnectorResourceSchema()),
	}
}

func (r *EntraIdConnectionResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	opCtx := errorsutil.CreateOperationContext(errorsutil.ConnectorTypeEntraID, "configure", "")
	ctx = opCtx.AddContextToLogger(ctx)

	opCtx.LogOperationStart(ctx, "Starting EntraID connection resource configuration")

	// Check if provider data is available.
	if req.ProviderData == nil {
		tflog.Debug(ctx, "ProviderData is nil, returning early")
		opCtx.LogOperationEnd(ctx, "EntraID connection resource configuration completed - no provider data")
		return
	}

	// Cast provider data to your provider type.
	prov, ok := req.ProviderData.(*SaviyntProvider)
	if !ok {
		errorCode := entraIdErrorCodes.ProviderConfig()
		opCtx.LogOperationError(ctx, "Provider configuration failed", errorCode,
			fmt.Errorf("expected *saviyntProvider, got different type"),
			map[string]interface{}{"expected_type": "*saviyntProvider"})

		resp.Diagnostics.AddError(
			errorsutil.GetErrorMessage(errorsutil.ErrProviderConfig),
			fmt.Sprintf("[%s] Expected *saviyntProvider, got different type", errorCode),
		)
		return
	}

	// Set the client and token from the provider state.
	r.client = &client.SaviyntClientWrapper{Client: prov.client}
	r.token = prov.accessToken
	r.provider = &client.SaviyntProviderWrapper{Provider: prov} // Store provider reference for retry logic

	opCtx.LogOperationEnd(ctx, "EntraID connection resource configured successfully")
}

// SetClient sets the client for testing purposes
func (r *EntraIdConnectionResource) SetClient(client client.SaviyntClientInterface) {
	r.client = client
}

// SetToken sets the token for testing purposes
func (r *EntraIdConnectionResource) SetToken(token string) {
	r.token = token
}

// SetProvider sets the provider for testing purposes
func (r *EntraIdConnectionResource) SetProvider(provider client.SaviyntProviderInterface) {
	r.provider = provider
}

func (r *EntraIdConnectionResource) BuildEntraIdConnector(plan *EntraIdConnectorResourceModel, config *EntraIdConnectorResourceModel) openapi.EntraIDConnector {
	var clientSecret string
	if !config.ClientSecret.IsNull() && !config.ClientSecret.IsUnknown() {
		clientSecret = config.ClientSecret.ValueString()
	} else if !config.ClientSecretWO.IsNull() && !config.ClientSecretWO.IsUnknown() {
		clientSecret = config.ClientSecretWO.ValueString()
	}

	var accessToken string
	if !config.AccessToken.IsNull() && !config.AccessToken.IsUnknown() {
		accessToken = config.AccessToken.ValueString()
	} else if !config.AccessTokenWO.IsNull() && !config.AccessTokenWO.IsUnknown() {
		accessToken = config.AccessTokenWO.ValueString()
	}

	var azureMgmtAccessToken string
	if !config.AzureMgmtAccessToken.IsNull() && !config.AzureMgmtAccessToken.IsUnknown() {
		azureMgmtAccessToken = config.AzureMgmtAccessToken.ValueString()
	} else if !config.AzureMgmtAccessTokenWO.IsNull() && !config.AzureMgmtAccessTokenWO.IsUnknown() {
		azureMgmtAccessToken = config.AzureMgmtAccessTokenWO.ValueString()
	}

	var windowsConnectorJson string
	if !config.WindowsConnectorJson.IsNull() && !config.WindowsConnectorJson.IsUnknown() {
		windowsConnectorJson = config.WindowsConnectorJson.ValueString()
	} else if !config.WindowsConnectorJsonWO.IsNull() && !config.WindowsConnectorJsonWO.IsUnknown() {
		windowsConnectorJson = config.WindowsConnectorJsonWO.ValueString()
	}

	var connectionJson string
	if !config.ConnectionJson.IsNull() && !config.ConnectionJson.IsUnknown() {
		connectionJson = config.ConnectionJson.ValueString()
	} else if !config.ConnectionJsonWO.IsNull() && !config.ConnectionJsonWO.IsUnknown() {
		connectionJson = config.ConnectionJsonWO.ValueString()
	}

	entraidConn := openapi.EntraIDConnector{
		BaseConnector: openapi.BaseConnector{
			//required fields
			Connectiontype: "AzureAD",
			ConnectionName: plan.ConnectionName.ValueString(),
			//optional fields
			ConnectionDescription: util.StringPointerOrEmpty(plan.Description),
			Defaultsavroles:       util.StringPointerOrEmpty(plan.DefaultSavRoles),
			EmailTemplate:         util.StringPointerOrEmpty(plan.EmailTemplate),
		},
		//required fields
		CLIENT_ID:     plan.ClientId.ValueString(),
		CLIENT_SECRET: clientSecret,
		AAD_TENANT_ID: plan.AadTenantId.ValueString(),
		//optional fields
		ACCESS_TOKEN:                    util.StringPointerOrEmpty(types.StringValue(accessToken)),
		AZURE_MGMT_ACCESS_TOKEN:         util.StringPointerOrEmpty(types.StringValue(azureMgmtAccessToken)),
		AUTHENTICATION_ENDPOINT:         util.StringPointerOrEmpty(plan.AuthenticationEndpoint),
		MICROSOFT_GRAPH_ENDPOINT:        util.StringPointerOrEmpty(plan.MicrosoftGraphEndpoint),
		AZURE_MANAGEMENT_ENDPOINT:       util.StringPointerOrEmpty(plan.AzureManagementEndpoint),
		ImportUserJSON:                  util.StringPointerOrEmpty(plan.ImportUserJson),
		CREATEUSERS:                     util.StringPointerOrEmpty(plan.CreateUsers),
		WINDOWS_CONNECTOR_JSON:          util.StringPointerOrEmpty(types.StringValue(windowsConnectorJson)),
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
		ChangePassJSON:                  util.StringPointerOrEmpty(plan.ChangePassJson),
		RemoveAccountJSON:               util.StringPointerOrEmpty(plan.RemoveAccountJson),
		ConnectionJSON:                  util.StringPointerOrEmpty(types.StringValue(connectionJson)),
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

	if plan.VaultConnection.ValueString() != "" {
		entraidConn.BaseConnector.VaultConnection = util.SafeStringConnector(plan.VaultConnection.ValueString())
		entraidConn.BaseConnector.VaultConfiguration = util.SafeStringConnector(plan.VaultConfiguration.ValueString())
		entraidConn.BaseConnector.Saveinvault = util.SafeStringConnector(plan.SaveInVault.ValueString())
	}

	return entraidConn
}

func (r *EntraIdConnectionResource) CreateEntraIdConnection(ctx context.Context, plan *EntraIdConnectorResourceModel, config *EntraIdConnectorResourceModel) (*openapi.CreateOrUpdateResponse, error) {
	connectionName := plan.ConnectionName.ValueString()
	opCtx := errorsutil.CreateOperationContext(errorsutil.ConnectorTypeEntraID, "create", connectionName)
	logCtx := opCtx.AddContextToLogger(ctx)

	opCtx.LogOperationStart(logCtx, "Starting EntraID connection creation")

	// Check if connection already exists (idempotency check) with retry logic
	var existingResource *openapi.GetConnectionDetailsResponse
	var finalHttpResp *http.Response

	err := r.provider.AuthenticatedAPICallWithRetry(ctx, "get_connection_details_idempotency", func(token string) error {
		connectionOps := r.connectionFactory.CreateConnectionOperations(r.client.APIBaseURL(), token)
		resp, httpResp, err := connectionOps.GetConnectionDetails(ctx, connectionName)
		if httpResp != nil && httpResp.StatusCode == 401 {
			return fmt.Errorf("401 unauthorized")
		}
		existingResource = resp
		finalHttpResp = httpResp // Update on every call including retries
		return err
	})

	if err != nil && finalHttpResp != nil && finalHttpResp.StatusCode != 412 {
		errorCode := entraIdErrorCodes.ReadFailed()
		opCtx.LogOperationError(logCtx, "Failed to check existing connection", errorCode, err)
		return nil, errorsutil.CreateStandardError(errorsutil.ConnectorTypeEntraID, errorCode, "create", connectionName, err)
	}

	if existingResource != nil &&
		existingResource.EntraIDConnectionResponse != nil &&
		existingResource.EntraIDConnectionResponse.Errorcode != nil &&
		*existingResource.EntraIDConnectionResponse.Errorcode == 0 {

		errorCode := entraIdErrorCodes.DuplicateName()
		opCtx.LogOperationError(logCtx, "Connection name already exists. Please import or use a different name", errorCode,
			fmt.Errorf("duplicate connection name"))
		return nil, errorsutil.CreateStandardError(errorsutil.ConnectorTypeEntraID, errorCode, "create", connectionName, nil)
	}

	// Build EntraID connection create request
	tflog.Debug(logCtx, "Building EntraID connection create request")

	// if (config.ClientSecret.IsNull() || config.ClientSecret.IsUnknown()) && (config.ClientSecretWO.IsNull() || config.ClientSecretWO.IsUnknown()) {
	// 	return nil, fmt.Errorf("either client_secret or client_secret_wo must be set")
	// }

	entraIdConn := r.BuildEntraIdConnector(plan, config)
	createReq := openapi.CreateOrUpdateRequest{
		EntraIDConnector: &entraIdConn,
	}

	// Execute create operation with retry logic
	var apiResp *openapi.CreateOrUpdateResponse

	err = r.provider.AuthenticatedAPICallWithRetry(ctx, "create_entraid_connection", func(token string) error {
		connectionOps := r.connectionFactory.CreateConnectionOperations(r.client.APIBaseURL(), token)
		resp, httpResp, err := connectionOps.CreateOrUpdateConnection(ctx, createReq)
		if httpResp != nil && httpResp.StatusCode == 401 {
			return fmt.Errorf("401 unauthorized")
		}
		apiResp = resp
		return err
	})

	if err != nil {
		errorCode := entraIdErrorCodes.CreateFailed()
		opCtx.LogOperationError(logCtx, "Failed to create EntraID connection", errorCode, err)
		return nil, errorsutil.CreateStandardError(errorsutil.ConnectorTypeEntraID, errorCode, "create", connectionName, err)
	}

	if apiResp != nil && apiResp.ErrorCode != nil && *apiResp.ErrorCode != "0" {
		apiErr := fmt.Errorf("API returned error code %s: %s", *apiResp.ErrorCode, errorsutil.SanitizeMessage(apiResp.Msg))
		errorCode := entraIdErrorCodes.APIError()
		opCtx.LogOperationError(logCtx, "EntraID connection creation failed with API error", errorCode, apiErr,
			map[string]interface{}{
				"api_error_code": *apiResp.ErrorCode,
				"message":        errorsutil.SanitizeMessage(apiResp.Msg),
			})
		return nil, errorsutil.CreateStandardError(errorsutil.ConnectorTypeEntraID, errorCode, "create", connectionName, apiErr)
	}

	opCtx.LogOperationEnd(logCtx, "EntraID connection created successfully",
		map[string]interface{}{"connection_key": func() interface{} {
			if apiResp.ConnectionKey != nil {
				return *apiResp.ConnectionKey
			}
			return nil
		}()})

	return apiResp, nil
}

func (r *EntraIdConnectionResource) ReadEntraIdConnection(ctx context.Context, connectionName string) (*openapi.GetConnectionDetailsResponse, error) {
	opCtx := errorsutil.CreateOperationContext(errorsutil.ConnectorTypeEntraID, "read", connectionName)
	logCtx := opCtx.AddContextToLogger(ctx)

	opCtx.LogOperationStart(logCtx, "Starting EntraID connection read")

	// Execute read operation with retry logic
	var apiResp *openapi.GetConnectionDetailsResponse

	err := r.provider.AuthenticatedAPICallWithRetry(ctx, "read_entraid_connection", func(token string) error {
		connectionOps := r.connectionFactory.CreateConnectionOperations(r.client.APIBaseURL(), token)
		resp, httpResp, err := connectionOps.GetConnectionDetails(ctx, connectionName)
		if httpResp != nil && httpResp.StatusCode == 401 {
			return fmt.Errorf("401 unauthorized")
		}
		apiResp = resp
		return err
	})

	if err != nil {
		errorCode := entraIdErrorCodes.ReadFailed()
		opCtx.LogOperationError(logCtx, "Failed to read EntraID connection", errorCode, err)
		return nil, errorsutil.CreateStandardError(errorsutil.ConnectorTypeEntraID, errorCode, "read", connectionName, err)
	}

	if err := r.ValidateEntraIdConnectionResponse(apiResp); err != nil {
		errorCode := entraIdErrorCodes.APIError()
		opCtx.LogOperationError(ctx, "Invalid connection type for EntraID datasource", errorCode, err)
		return nil, fmt.Errorf("[%s] Unable to verify connection type for connection %q. The provider could not determine the type of this connection. Please ensure the connection name is correct and belongs to a supported connector type", errorCode, connectionName)
	}

	if apiResp != nil && apiResp.EntraIDConnectionResponse != nil && apiResp.EntraIDConnectionResponse.Errorcode != nil && *apiResp.EntraIDConnectionResponse.Errorcode != 0 {
		apiErr := fmt.Errorf("API returned error code %d: %s", *apiResp.EntraIDConnectionResponse.Errorcode, errorsutil.SanitizeMessage(apiResp.EntraIDConnectionResponse.Msg))
		errorCode := entraIdErrorCodes.APIError()
		opCtx.LogOperationError(logCtx, "EntraID connection read failed with API error", errorCode, apiErr,
			map[string]interface{}{
				"api_error_code": *apiResp.EntraIDConnectionResponse.Errorcode,
				"message":        errorsutil.SanitizeMessage(apiResp.EntraIDConnectionResponse.Msg),
			})
		return nil, errorsutil.CreateStandardError(errorsutil.ConnectorTypeEntraID, errorCode, "read", connectionName, apiErr)
	}

	opCtx.LogOperationEnd(logCtx, "EntraID connection read completed successfully")
	return apiResp, nil
}

func (r *EntraIdConnectionResource) UpdateEntraIdConnection(ctx context.Context, plan *EntraIdConnectorResourceModel, config *EntraIdConnectorResourceModel) (*openapi.CreateOrUpdateResponse, error) {
	connectionName := plan.ConnectionName.ValueString()
	opCtx := errorsutil.CreateOperationContext(errorsutil.ConnectorTypeEntraID, "update", connectionName)
	logCtx := opCtx.AddContextToLogger(ctx)

	opCtx.LogOperationStart(logCtx, "Starting EntraID connection update")

	// Build EntraID connection update request
	tflog.Debug(logCtx, "Building EntraID connection update request")

	// if (config.ClientSecret.IsNull() || config.ClientSecret.IsUnknown()) && (config.ClientSecretWO.IsNull() || config.ClientSecretWO.IsUnknown()) {
	// 	return nil, fmt.Errorf("either client_secret or client_secret_wo must be set")
	// }

	entraIdConn := r.BuildEntraIdConnector(plan, config) // Reuse the same request builder

	updateReq := openapi.CreateOrUpdateRequest{
		EntraIDConnector: &entraIdConn,
	}

	// Execute update operation with retry logic
	var apiResp *openapi.CreateOrUpdateResponse

	err := r.provider.AuthenticatedAPICallWithRetry(ctx, "update_entraid_connection", func(token string) error {
		connectionOps := r.connectionFactory.CreateConnectionOperations(r.client.APIBaseURL(), token)
		resp, httpResp, err := connectionOps.CreateOrUpdateConnection(ctx, updateReq)
		if httpResp != nil && httpResp.StatusCode == 401 {
			return fmt.Errorf("401 unauthorized")
		}
		apiResp = resp
		return err
	})

	if err != nil {
		errorCode := entraIdErrorCodes.UpdateFailed()
		opCtx.LogOperationError(logCtx, "Failed to update EntraID connection", errorCode, err)
		return nil, errorsutil.CreateStandardError(errorsutil.ConnectorTypeEntraID, errorCode, "update", connectionName, err)
	}

	if apiResp != nil && apiResp.ErrorCode != nil && *apiResp.ErrorCode != "0" {
		apiErr := fmt.Errorf("API returned error code %s: %s", *apiResp.ErrorCode, errorsutil.SanitizeMessage(apiResp.Msg))
		errorCode := entraIdErrorCodes.APIError()
		opCtx.LogOperationError(logCtx, "EntraID connection update failed with API error", errorCode, apiErr,
			map[string]interface{}{
				"api_error_code": *apiResp.ErrorCode,
				"message":        errorsutil.SanitizeMessage(apiResp.Msg),
			})
		return nil, errorsutil.CreateStandardError(errorsutil.ConnectorTypeEntraID, errorCode, "update", connectionName, apiErr)
	}

	opCtx.LogOperationEnd(logCtx, "EntraID connection updated successfully",
		map[string]interface{}{"connection_key": func() interface{} {
			if apiResp.ConnectionKey != nil {
				return *apiResp.ConnectionKey
			}
			return nil
		}()})

	return apiResp, nil
}

func (r *EntraIdConnectionResource) UpdateModelFromCreateResponse(plan *EntraIdConnectorResourceModel, apiResp *openapi.CreateOrUpdateResponse) {
	plan.ID = types.StringValue(fmt.Sprintf("%d", *apiResp.ConnectionKey))
	plan.ConnectionKey = types.Int64Value(int64(*apiResp.ConnectionKey))
	plan.Description = util.SafeStringDatasource(plan.Description.ValueStringPointer())
	plan.DefaultSavRoles = util.SafeStringDatasource(plan.DefaultSavRoles.ValueStringPointer())
	plan.EmailTemplate = util.SafeStringDatasource(plan.EmailTemplate.ValueStringPointer())
	plan.AadTenantId = util.SafeStringDatasource(plan.AadTenantId.ValueStringPointer())
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
	plan.ChangePassJson = util.SafeStringDatasource(plan.ChangePassJson.ValueStringPointer())

	// Set response fields
	plan.ErrorCode = util.SafeStringDatasource(apiResp.ErrorCode)
	plan.Msg = util.SafeStringDatasource(apiResp.Msg)
}

func (r *EntraIdConnectionResource) UpdateModelFromReadResponse(state *EntraIdConnectorResourceModel, apiResp *openapi.GetConnectionDetailsResponse) {
	// Set basic connection info
	state.ConnectionKey = types.Int64Value(int64(*apiResp.EntraIDConnectionResponse.Connectionkey))
	state.ID = types.StringValue(fmt.Sprintf("%d", *apiResp.EntraIDConnectionResponse.Connectionkey))
	state.ConnectionName = util.SafeStringDatasource(apiResp.EntraIDConnectionResponse.Connectionname)
	state.Description = util.SafeStringDatasource(apiResp.EntraIDConnectionResponse.Description)
	state.DefaultSavRoles = util.SafeStringDatasource(apiResp.EntraIDConnectionResponse.Defaultsavroles)
	state.EmailTemplate = util.SafeStringDatasource(apiResp.EntraIDConnectionResponse.Emailtemplate)

	// Map all EntraID-specific attributes from the connection attributes
	if apiResp.EntraIDConnectionResponse.Connectionattributes != nil {
		attrs := apiResp.EntraIDConnectionResponse.Connectionattributes

		// Authentication and endpoint configuration
		state.ClientId = util.SafeStringDatasource(attrs.CLIENT_ID)
		state.AadTenantId = util.SafeStringDatasource(attrs.AAD_TENANT_ID)
		state.AuthenticationEndpoint = util.SafeStringDatasource(attrs.AUTHENTICATION_ENDPOINT)
		state.MicrosoftGraphEndpoint = util.SafeStringDatasource(attrs.MICROSOFT_GRAPH_ENDPOINT)
		state.AzureManagementEndpoint = util.SafeStringDatasource(attrs.AZURE_MANAGEMENT_ENDPOINT)

		// User and account management
		state.ImportUserJson = util.SafeStringDatasource(attrs.ImportUserJSON)
		state.CreateUsers = util.SafeStringDatasource(attrs.CREATEUSERS)
		state.CreateNewEndpoints = util.SafeStringDatasource(attrs.CREATE_NEW_ENDPOINTS)
		state.ManagedAccountType = util.SafeStringDatasource(attrs.MANAGED_ACCOUNT_TYPE)
		state.AccountAttributes = util.SafeStringDatasource(attrs.ACCOUNT_ATTRIBUTES)
		state.ServiceAccountAttributes = util.SafeStringDatasource(attrs.SERVICE_ACCOUNT_ATTRIBUTES)
		state.DeltaTokensJson = util.SafeStringDatasource(attrs.DELTATOKENSJSON)
		state.AccountImportFields = util.SafeStringDatasource(attrs.ACCOUNT_IMPORT_FIELDS)
		state.ImportDepth = util.SafeStringDatasource(attrs.IMPORT_DEPTH)
		state.EntitlementAttribute = util.SafeStringDatasource(attrs.ENTITLEMENT_ATTRIBUTE)

		// Account lifecycle operations
		state.CreateAccountJson = util.SafeStringDatasource(attrs.CreateAccountJSON)
		state.UpdateAccountJson = util.SafeStringDatasource(attrs.UpdateAccountJSON)
		state.EnableAccountJson = util.SafeStringDatasource(attrs.EnableAccountJSON)
		state.DisableAccountJson = util.SafeStringDatasource(attrs.DisableAccountJSON)
		state.RemoveAccountJson = util.SafeStringDatasource(attrs.RemoveAccountJSON)

		// Access management
		state.AddAccessJson = util.SafeStringDatasource(attrs.AddAccessJSON)
		state.RemoveAccessJson = util.SafeStringDatasource(attrs.RemoveAccessJSON)
		state.AddAccessToEntitlementJson = util.SafeStringDatasource(attrs.AddAccessToEntitlementJSON)
		state.RemoveAccessFromEntitlementJson = util.SafeStringDatasource(attrs.RemoveAccessFromEntitlementJSON)

		// User management
		state.UpdateUserJson = util.SafeStringDatasource(attrs.UpdateUserJSON)
		state.ChangePassJson = util.SafeStringDatasource(attrs.ChangePassJSON)
		state.ModifyUserdataJson = util.SafeStringDatasource(attrs.MODIFYUSERDATAJSON)

		// Group management
		state.CreateGroupJson = util.SafeStringDatasource(attrs.CreateGroupJSON)
		state.UpdateGroupJson = util.SafeStringDatasource(attrs.UpdateGroupJSON)
		state.DeleteGroupJson = util.SafeStringDatasource(attrs.DeleteGroupJSON)

		// Service principal management
		state.CreateServicePrincipalJson = util.SafeStringDatasource(attrs.CreateServicePrincipalJSON)
		state.UpdateServicePrincipalJson = util.SafeStringDatasource(attrs.UpdateServicePrincipalJSON)
		state.RemoveServicePrincipalJson = util.SafeStringDatasource(attrs.RemoveServicePrincipalJSON)

		// Teams and channels
		state.CreateTeamJson = util.SafeStringDatasource(attrs.CreateTeamJSON)
		state.CreateChannelJson = util.SafeStringDatasource(attrs.CreateChannelJSON)

		// Filtering and configuration
		state.EntitlementFilterJson = util.SafeStringDatasource(attrs.ENTITLEMENT_FILTER_JSON)
		state.StatusThresholdConfig = util.SafeStringDatasource(attrs.STATUS_THRESHOLD_CONFIG)
		state.AccountsFilter = util.SafeStringDatasource(attrs.ACCOUNTS_FILTER)
		state.PamConfig = util.SafeStringDatasource(attrs.PAM_CONFIG)
		state.EndpointsFilter = util.SafeStringDatasource(attrs.ENDPOINTS_FILTER)
		state.ConfigJson = util.SafeStringDatasource(attrs.ConfigJSON)
		state.EnhancedDirectoryRoles = util.SafeStringDatasource(attrs.ENHANCEDDIRECTORYROLES)
	}
}

func (r *EntraIdConnectionResource) ValidateEntraIdConnectionResponse(apiResp *openapi.GetConnectionDetailsResponse) error {
	if apiResp != nil && apiResp.EntraIDConnectionResponse == nil {
		return fmt.Errorf("verify the connection type - EntraId connection response is nil")
	}
	return nil
}

func (r *EntraIdConnectionResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan, config EntraIdConnectorResourceModel

	opCtx := errorsutil.CreateOperationContext(errorsutil.ConnectorTypeEntraID, "terraform_create", "")
	ctx = opCtx.AddContextToLogger(ctx)

	opCtx.LogOperationStart(ctx, "Starting EntraID connection Terraform create")

	// Extract plan from request
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		errorCode := entraIdErrorCodes.PlanExtraction()
		opCtx.LogOperationError(ctx, "Failed to get plan from request", errorCode,
			fmt.Errorf("plan extraction failed"))
		resp.Diagnostics.AddError(
			errorsutil.GetErrorMessage(errorsutil.ErrPlanExtraction),
			fmt.Sprintf("[%s] Unable to extract Terraform plan from request", errorCode),
		)
		return
	}

	// Extract config from request
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		errorCode := entraIdErrorCodes.ConfigExtraction()
		opCtx.LogOperationError(ctx, "Failed to get config from request", errorCode,
			fmt.Errorf("config extraction failed"))
		resp.Diagnostics.AddError(
			errorsutil.GetErrorMessage(errorsutil.ErrConfigExtraction),
			fmt.Sprintf("[%s] Unable to extract Terraform configuration from request", errorCode),
		)
		return
	}

	// Update operation context with connection name
	connectionName := plan.ConnectionName.ValueString()
	opCtx.ConnectionName = connectionName
	ctx = opCtx.AddContextToLogger(ctx)

	// Use the helper method to create the connection
	apiResp, err := r.CreateEntraIdConnection(ctx, &plan, &config)
	if err != nil {
		opCtx.LogOperationError(ctx, "EntraID connection creation failed", "", err)
		resp.Diagnostics.AddError(
			"EntraID Connection Creation Failed",
			fmt.Sprintf("Unable to create EntraID connection: %s", err.Error()),
		)
		return
	}

	// Update the model with the response data
	r.UpdateModelFromCreateResponse(&plan, apiResp)

	// Set the state
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		errorCode := entraIdErrorCodes.StateUpdate()
		opCtx.LogOperationError(ctx, "Failed to set state", errorCode,
			fmt.Errorf("state update failed"))
		resp.Diagnostics.AddError(
			errorsutil.GetErrorMessage(errorsutil.ErrStateUpdate),
			fmt.Sprintf("[%s] Unable to update Terraform state for EntraID connection", errorCode),
		)
		return
	}

	opCtx.LogOperationEnd(ctx, "EntraID connection Terraform create completed successfully",
		map[string]interface{}{
			"connection_name": connectionName,
			"connection_key": func() interface{} {
				if apiResp.ConnectionKey != nil {
					return *apiResp.ConnectionKey
				}
				return nil
			}(),
		})
}

func (r *EntraIdConnectionResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state EntraIdConnectorResourceModel

	opCtx := errorsutil.CreateOperationContext(errorsutil.ConnectorTypeEntraID, "terraform_read", "")
	ctx = opCtx.AddContextToLogger(ctx)

	opCtx.LogOperationStart(ctx, "Starting EntraID connection Terraform read")

	// Extract state from request
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		errorCode := entraIdErrorCodes.StateExtraction()
		opCtx.LogOperationError(ctx, "Failed to get state from request", errorCode,
			fmt.Errorf("state extraction failed"))
		resp.Diagnostics.AddError(
			errorsutil.GetErrorMessage(errorsutil.ErrStateExtraction),
			fmt.Sprintf("[%s] Unable to extract Terraform state from request", errorCode),
		)
		return
	}

	// Update operation context with connection name
	connectionName := state.ConnectionName.ValueString()
	opCtx.ConnectionName = connectionName
	ctx = opCtx.AddContextToLogger(ctx)

	// Use the helper method to read the connection
	apiResp, err := r.ReadEntraIdConnection(ctx, connectionName)
	if err != nil {
		opCtx.LogOperationError(ctx, "EntraID connection read failed", "", err)
		resp.Diagnostics.AddError(
			"EntraID Connection Read Failed",
			fmt.Sprintf("Unable to read EntraID connection: %s", err.Error()),
		)
		return
	}

	// Update the model with the response data
	r.UpdateModelFromReadResponse(&state, apiResp)

	apiMessage := util.SafeDeref(apiResp.EntraIDConnectionResponse.Msg)
	if apiMessage == "success" {
		state.Msg = types.StringValue("Connection Read Successful")
	} else {
		state.Msg = types.StringValue(apiMessage)
	}
	state.ErrorCode = util.Int32PtrToTFString(apiResp.EntraIDConnectionResponse.Errorcode)

	// Set the state
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		errorCode := entraIdErrorCodes.StateUpdate()
		opCtx.LogOperationError(ctx, "Failed to set state", errorCode,
			fmt.Errorf("state update failed"))
		resp.Diagnostics.AddError(
			errorsutil.GetErrorMessage(errorsutil.ErrStateUpdate),
			fmt.Sprintf("[%s] Unable to update Terraform state for EntraID connection", errorCode),
		)
		return
	}

	opCtx.LogOperationEnd(ctx, "EntraID connection Terraform read completed successfully",
		map[string]interface{}{"connection_name": connectionName})
}

func (r *EntraIdConnectionResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state, config EntraIdConnectorResourceModel

	opCtx := errorsutil.CreateOperationContext(errorsutil.ConnectorTypeEntraID, "terraform_update", "")
	ctx = opCtx.AddContextToLogger(ctx)

	opCtx.LogOperationStart(ctx, "Starting EntraID connection Terraform update")

	// Extract state from request
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		errorCode := entraIdErrorCodes.StateExtraction()
		opCtx.LogOperationError(ctx, "Failed to get state from request", errorCode,
			fmt.Errorf("state extraction failed"))
		resp.Diagnostics.AddError(
			errorsutil.GetErrorMessage(errorsutil.ErrStateExtraction),
			fmt.Sprintf("[%s] Unable to extract Terraform state from request", errorCode),
		)
		return
	}

	// Extract plan from request
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		errorCode := entraIdErrorCodes.PlanExtraction()
		opCtx.LogOperationError(ctx, "Failed to get plan from request", errorCode,
			fmt.Errorf("plan extraction failed"))
		resp.Diagnostics.AddError(
			errorsutil.GetErrorMessage(errorsutil.ErrPlanExtraction),
			fmt.Sprintf("[%s] Unable to extract Terraform plan from request", errorCode),
		)
		return
	}

	// Extract config from request
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		errorCode := entraIdErrorCodes.ConfigExtraction()
		opCtx.LogOperationError(ctx, "Failed to get config from request", errorCode,
			fmt.Errorf("config extraction failed"))
		resp.Diagnostics.AddError(
			errorsutil.GetErrorMessage(errorsutil.ErrConfigExtraction),
			fmt.Sprintf("[%s] Unable to extract Terraform configuration from request", errorCode),
		)
		return
	}

	// Update operation context with connection name
	connectionName := plan.ConnectionName.ValueString()
	opCtx.ConnectionName = connectionName
	ctx = opCtx.AddContextToLogger(ctx)

	// Validate that connection name hasn't changed (not supported)
	if plan.ConnectionName.ValueString() != state.ConnectionName.ValueString() {
		errorCode := entraIdErrorCodes.NameImmutable()
		opCtx.LogOperationError(ctx, "Connection name cannot be updated", errorCode,
			fmt.Errorf("attempted to change connection name from '%s' to '%s'", state.ConnectionName.ValueString(), plan.ConnectionName.ValueString()),
			map[string]interface{}{
				"old_name": state.ConnectionName.ValueString(),
				"new_name": plan.ConnectionName.ValueString(),
			})
		resp.Diagnostics.AddError(
			errorsutil.GetErrorMessage(errorCode),
			fmt.Sprintf("[%s] Cannot change connection name from '%s' to '%s'", errorCode, state.ConnectionName.ValueString(), plan.ConnectionName.ValueString()),
		)
		return
	}

	// Use the helper method to update the connection
	updateResp, err := r.UpdateEntraIdConnection(ctx, &plan, &config)
	if err != nil {
		opCtx.LogOperationError(ctx, "EntraID connection update failed", "", err)
		resp.Diagnostics.AddError(
			"EntraID Connection Update Failed",
			fmt.Sprintf("Unable to update EntraID connection: %s", err.Error()),
		)
		return
	}

	// Refresh state after update
	getResp, err := r.ReadEntraIdConnection(ctx, connectionName)
	if err != nil {
		opCtx.LogOperationError(ctx, "Failed to read updated EntraID connection", "", err)
		resp.Diagnostics.AddError(
			"EntraID Connection Post-Update Read Failed",
			fmt.Sprintf("Update succeeded but failed to read updated connection state: %s", err.Error()),
		)
		return
	}

	// Update the model with the response data (includes read operation)
	r.UpdateModelFromReadResponse(&plan, getResp)

	apiMessage := util.SafeDeref(updateResp.Msg)
	plan.Msg = types.StringValue(apiMessage)
	plan.ErrorCode = types.StringValue(*updateResp.ErrorCode)

	// Set the state
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		errorCode := entraIdErrorCodes.StateUpdate()
		opCtx.LogOperationError(ctx, "Failed to update state after successful update", errorCode,
			fmt.Errorf("state update failed"))
		resp.Diagnostics.AddError(
			errorsutil.GetErrorMessage(errorsutil.ErrStateUpdate),
			fmt.Sprintf("[%s] Unable to update Terraform state after successful EntraID connection update", errorCode),
		)
		return
	}

	opCtx.LogOperationEnd(ctx, "EntraID connection Terraform update completed successfully",
		map[string]interface{}{"connection_name": connectionName})
}
func (r *EntraIdConnectionResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
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

func (r *EntraIdConnectionResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	// Importing an AD connection resource requires the connection name
	connectionName := req.ID
	opCtx := errorsutil.CreateOperationContext(errorsutil.ConnectorTypeEntraID, "terraform_import", connectionName)
	ctx = opCtx.AddContextToLogger(ctx)

	opCtx.LogOperationStart(ctx, "Starting EntraID connection resource import")

	// Retrieve import ID and save to id attribute
	resource.ImportStatePassthroughID(ctx, path.Root("connection_name"), req, resp)

	opCtx.LogOperationEnd(ctx, "EntraID connection resource import completed successfully",
		map[string]interface{}{"import_id": connectionName})
}
