// Copyright (c) 2025 Saviynt Inc.
// SPDX-License-Identifier: MPL-2.0

// saviynt_adsi_connection_resource manages ADSI connectors in the Saviynt Security Manager.
// The resource implements the full Terraform lifecycle:
//   - Create: provisions a new ADSI connector using the supplied configuration.
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
	"terraform-provider-Saviynt/util/errorsutil"

	connectionsutil "terraform-provider-Saviynt/util/connectionsutil"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	openapi "github.com/saviynt/saviynt-api-go-client/connections"
)

// Ensure provider defined types fully satisfy framework interfaces
var _ resource.Resource = &AdsiConnectionResource{}
var _ resource.ResourceWithImportState = &AdsiConnectionResource{}

// Initialize error codes for ADSI Connection operations
var adsiErrorCodes = errorsutil.NewConnectorErrorCodes(errorsutil.ConnectorTypeADSI)

type ADSIConnectorResourceModel struct {
	BaseConnectorResourceModel
	ID                          types.String `tfsdk:"id"`
	URL                         types.String `tfsdk:"url"`
	Username                    types.String `tfsdk:"username"`
	Password                    types.String `tfsdk:"password"`
	ConnectionUrl               types.String `tfsdk:"connection_url"`
	ProvisioningUrl             types.String `tfsdk:"provisioning_url"`
	ForestList                  types.String `tfsdk:"forestlist"`
	DefaultUserRole             types.String `tfsdk:"default_user_role"`
	UpdateUserJson              types.String `tfsdk:"updateuserjson"`
	EndpointsFilter             types.String `tfsdk:"endpoints_filter"`
	SearchFilter                types.String `tfsdk:"searchfilter"`
	ObjectFilter                types.String `tfsdk:"objectfilter"`
	AccountAttribute            types.String `tfsdk:"account_attribute"`
	StatusThresholdConfig       types.String `tfsdk:"status_threshold_config"`
	EntitlementAttribute        types.String `tfsdk:"entitlement_attribute"`
	UserAttribute               types.String `tfsdk:"user_attribute"`
	GroupSearchBaseDN           types.String `tfsdk:"group_search_base_dn"`
	CheckForUnique              types.String `tfsdk:"checkforunique"`
	StatusKeyJson               types.String `tfsdk:"statuskeyjson"`
	GroupImportMapping          types.String `tfsdk:"group_import_mapping"`
	ImportNestedMembership      types.String `tfsdk:"import_nested_membership"`
	PageSize                    types.String `tfsdk:"page_size"`
	AccountNameRule             types.String `tfsdk:"accountnamerule"`
	CreateAccountJson           types.String `tfsdk:"createaccountjson"`
	UpdateAccountJson           types.String `tfsdk:"updateaccountjson"`
	EnableAccountJson           types.String `tfsdk:"enableaccountjson"`
	DisableAccountJson          types.String `tfsdk:"disableaccountjson"`
	RemoveAccountJson           types.String `tfsdk:"removeaccountjson"`
	AddAccessJson               types.String `tfsdk:"addaccessjson"`
	RemoveAccessJson            types.String `tfsdk:"removeaccessjson"`
	ResetAndChangePasswrdJson   types.String `tfsdk:"resetandchangepasswrdjson"`
	CreateGroupJson             types.String `tfsdk:"creategroupjson"`
	UpdateGroupJson             types.String `tfsdk:"updategroupjson"`
	RemoveGroupJson             types.String `tfsdk:"removegroupjson"`
	AddAccessEntitlementJson    types.String `tfsdk:"addaccessentitlementjson"`
	CustomConfigJson            types.String `tfsdk:"customconfigjson"`
	RemoveAccessEntitlementJson types.String `tfsdk:"removeaccessentitlementjson"`
	CreateServiceAccountJson    types.String `tfsdk:"createserviceaccountjson"`
	UpdateServiceAccountJson    types.String `tfsdk:"updateserviceaccountjson"`
	RemoveServiceAccountJson    types.String `tfsdk:"removeserviceaccountjson"`
	PamConfig                   types.String `tfsdk:"pam_config"`
	ModifyUserDataJson          types.String `tfsdk:"modifyuserdatajson"`
}

type AdsiConnectionResource struct {
	client            client.SaviyntClientInterface
	token             string
	provider          client.SaviyntProviderInterface
	connectionFactory client.ConnectionFactoryInterface
}

func NewADSIConnectionResource() resource.Resource {
	return &AdsiConnectionResource{
		connectionFactory: &client.DefaultConnectionFactory{},
	}
}

func NewADSIConnectionResourceWithFactory(factory client.ConnectionFactoryInterface) resource.Resource {
	return &AdsiConnectionResource{
		connectionFactory: factory,
	}
}

func (r *AdsiConnectionResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = "saviynt_adsi_connection_resource"
}

func ADSIConnectorResourceSchema() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:    true,
			Description: "Resource ID.",
		},
		"url": schema.StringAttribute{
			Required:    true,
			Description: "Primary/root domain URL list (comma Separated)",
		},
		"username": schema.StringAttribute{
			Required:    true,
			WriteOnly:   true,
			Description: "Service account username",
		},
		"password": schema.StringAttribute{
			Required:    true,
			WriteOnly:   true,
			Description: "Service account password",
		},
		"connection_url": schema.StringAttribute{
			Required:    true,
			Description: "ADSI remote agent Connection URL",
		},
		"forestlist": schema.StringAttribute{
			Required:    true,
			Description: "Forest List (Comma Separated) which we need to manage",
		},
		"provisioning_url": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "ADSI remote agent Provisioning URL",
		},
		"default_user_role": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "Default SAV Role to be assigned to all the new users that gets imported via User Import",
		},
		"updateuserjson": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "Specify the attribute Value which will be used to Update existing User",
		},
		"endpoints_filter": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "Provide the configuration to create Child Endpoints and import associated accounts and entitlements",
		},
		"searchfilter": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "Account Search Filter to specify the starting point of the directory from where the accounts needs to be imported. You can have multiple BaseDNs here separated by ###.",
		},
		"objectfilter": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "Object Filter is used to filter the objects that will be returned.This filter will be same for all domains.",
		},
		"account_attribute": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "Map EIC and AD attributes for account import (AD attributes must be in lower case)",
		},
		"status_threshold_config": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "Account status and threshold related config",
		},
		"entitlement_attribute": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "Account attribute that contains group membership",
		},
		"user_attribute": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "Map EIC and AD attributes for user import (AD attributes must be in lower case)",
		},
		"group_search_base_dn": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "Group Search Filter to specify the starting point of the directory from where the groups needs to be imported. You can have multiple BaseDNs here separated by ###.",
		},
		"checkforunique": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "Evaluate the uniqueness of an attribute",
		},
		"statuskeyjson": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "JSON configuration to specify Users status",
		},
		"group_import_mapping": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "Map AD group attribute to EIC entitlement attribute for import",
		},
		"import_nested_membership": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "Specify if you want the connector to import all indirect or nested membership of an account or a group during access import",
		},
		"page_size": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "Page size defines the number of objects to be returned from each AD operation.",
		},
		"accountnamerule": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "Rule to generate account name.",
		},
		"createaccountjson": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "Specify the attributes values which will be used to Create the New Account.",
		},
		"updateaccountjson": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "Specify the attributes values which will be used to Update existing Account.",
		},
		"enableaccountjson": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "Specify the actions and attribute updates to be performed for enabling an account.",
		},
		"disableaccountjson": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "Specify the actions and attributes updates to be performed for disabling an account.",
		},
		"removeaccountjson": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "Specify the actions to be performed for deleting an account.",
		},
		"addaccessjson": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "Configuration to ADD Access (cross domain/forest group membership) to an account.",
		},
		"removeaccessjson": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "Configuration to REMOVE Access (cross domain/forest group membership) to an account.",
		},
		"resetandchangepasswrdjson": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "Configuration to Reset and Change Password.",
		},
		"creategroupjson": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "Configuration to Create a Group",
		},
		"updategroupjson": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "Configuration to Update a Group",
		},
		"removegroupjson": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "Configuration to Delete a Group",
		},
		"addaccessentitlementjson": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "Configuration to Add nested group hierarchy",
		},
		"customconfigjson": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "Custom configuration JSON",
		},
		"removeaccessentitlementjson": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "Configuration to Remove nested group hierarchy",
		},
		"createserviceaccountjson": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "Specify the Field Value which will be used to Create the New Service Account.",
		},
		"updateserviceaccountjson": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "Specify the Field Value which will be used to update the existing Service Account.",
		},
		"removeserviceaccountjson": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "Specify the actions to be performed while deleting a service account.",
		},
		"pam_config": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "JSON to specify Bootstrap Config.",
		},
		"modifyuserdatajson": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "Specify this parameter to transform the data during user import.",
		},
	}
}

func (r *AdsiConnectionResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: util.ADSIConnDescription,
		Attributes:  connectionsutil.MergeResourceAttributes(BaseConnectorResourceSchema(), ADSIConnectorResourceSchema()),
	}
}

func (r *AdsiConnectionResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	opCtx := errorsutil.CreateOperationContext(errorsutil.ConnectorTypeADSI, "configure", "")
	ctx = opCtx.AddContextToLogger(ctx)

	opCtx.LogOperationStart(ctx, "Starting ADSI connection resource configuration")

	// Check if provider data is available.
	if req.ProviderData == nil {
		tflog.Debug(ctx, "ProviderData is nil, returning early")
		opCtx.LogOperationEnd(ctx, "ADSI connection resource configuration completed - no provider data")
		return
	}

	// Cast provider data to your provider type.
	prov, ok := req.ProviderData.(*SaviyntProvider)
	if !ok {
		errorCode := adsiErrorCodes.ProviderConfig()
		opCtx.LogOperationError(ctx, "Provider configuration failed", errorCode,
			fmt.Errorf("expected *SaviyntProvider, got different type"),
			map[string]interface{}{"expected_type": "*SaviyntProvider"})

		resp.Diagnostics.AddError(
			errorsutil.GetErrorMessage(errorsutil.ErrProviderConfig),
			fmt.Sprintf("[%s] Expected *SaviyntProvider, got different type", errorCode),
		)
		return
	}

	// Set the client and token from the provider state using interface wrapper.
	r.client = &client.SaviyntClientWrapper{Client: prov.client}
	r.token = prov.accessToken
	r.provider = &client.SaviyntProviderWrapper{Provider: prov} // Store provider reference for retry logic 

	opCtx.LogOperationEnd(ctx, "ADSI connection resource configured successfully")
}

// SetClient sets the client for testing purposes
func (r *AdsiConnectionResource) SetClient(client client.SaviyntClientInterface) {
	r.client = client
}

// SetToken sets the token for testing purposes
func (r *AdsiConnectionResource) SetToken(token string) {
	r.token = token
}

// SetProvider sets the provider for testing purposes
func (r *AdsiConnectionResource) SetProvider(provider client.SaviyntProviderInterface) {
	r.provider = provider
}

func (r *AdsiConnectionResource) CreateADSIConnection(ctx context.Context, plan *ADSIConnectorResourceModel, config *ADSIConnectorResourceModel) (*openapi.CreateOrUpdateResponse, error) {
	connectionName := plan.ConnectionName.ValueString()
	opCtx := errorsutil.CreateOperationContext(errorsutil.ConnectorTypeADSI, "create", connectionName)

	// Create logging context (separate from API context)
	logCtx := opCtx.AddContextToLogger(ctx)

	opCtx.LogOperationStart(logCtx, "Starting ADSI connection creation")

	// Check if connection already exists (idempotency check) with retry logic
	tflog.Debug(logCtx, "Checking if connection already exists")
	var existingResource *openapi.GetConnectionDetailsResponse
	var httpsResp *http.Response
	err := r.provider.AuthenticatedAPICallWithRetry(ctx, "get_connection_details_idempotency", func(token string) error {
		connectionOps := r.connectionFactory.CreateConnectionOperations(r.client.APIBaseURL(), token)
		resp, httpResp, err := connectionOps.GetConnectionDetails(ctx, connectionName)
		if httpResp != nil && httpResp.StatusCode == 401 {
			return fmt.Errorf("401 unauthorized")
		}
		existingResource = resp
		httpsResp = httpResp
		return err
	})

	if err != nil && httpsResp != nil && httpsResp.StatusCode != 412 {
		errorCode := adsiErrorCodes.ReadFailed()
		opCtx.LogOperationError(logCtx, "Failed to check existing connection", errorCode, err, nil)
		return nil, fmt.Errorf("[%s] Failed to check existing connection: %w", errorCode, err)
	}

	if existingResource != nil &&
		existingResource.ADSIConnectionResponse != nil &&
		existingResource.ADSIConnectionResponse.Errorcode != nil &&
		*existingResource.ADSIConnectionResponse.Errorcode == 0 {

		errorCode := adsiErrorCodes.DuplicateName()
		opCtx.LogOperationError(ctx, "Connection name already exists. Please import or use a different name", errorCode,
			fmt.Errorf("duplicate connection name"))
		return nil, errorsutil.CreateStandardError(errorsutil.ConnectorTypeADSI, errorCode, "create", connectionName, nil)
	}

	// Build ADSI connection create request
	tflog.Debug(ctx, "Building ADSI connection create request")

	adsiConn := r.BuildADSIConnector(plan, config)
	createReq := openapi.CreateOrUpdateRequest{
		ADSIConnector: &adsiConn,
	}

	var apiResp *openapi.CreateOrUpdateResponse

	// Execute create operation with retry logic
	tflog.Debug(ctx, "Executing create operation")
	err = r.provider.AuthenticatedAPICallWithRetry(ctx, "create_adsi_connection", func(token string) error {
		connectionOps := r.connectionFactory.CreateConnectionOperations(r.client.APIBaseURL(), token)
		resp, httpResp, err := connectionOps.CreateOrUpdateConnection(ctx, createReq)
		if httpResp != nil && httpResp.StatusCode == 401 {
			return fmt.Errorf("401 unauthorized")
		}
		apiResp = resp
		return err
	})

	if err != nil {
		errorCode := adsiErrorCodes.CreateFailed()
		opCtx.LogOperationError(ctx, "Failed to create ADSI connection", errorCode, err)
		return nil, errorsutil.CreateStandardError(errorsutil.ConnectorTypeADSI, errorCode, "create", connectionName, err)
	}

	if apiResp != nil && apiResp.ErrorCode != nil && *apiResp.ErrorCode != "0" {
		apiErr := fmt.Errorf("API returned error code %s: %s", *apiResp.ErrorCode, errorsutil.SanitizeMessage(apiResp.Msg))
		errorCode := adsiErrorCodes.APIError()
		opCtx.LogOperationError(ctx, "ADSI connection creation failed with API error", errorCode, apiErr,
			map[string]interface{}{
				"api_error_code": *apiResp.ErrorCode,
				"message":        errorsutil.SanitizeMessage(apiResp.Msg),
			})
		return nil, errorsutil.CreateStandardError(errorsutil.ConnectorTypeADSI, errorCode, "create", connectionName, apiErr)
	}

	opCtx.LogOperationEnd(logCtx, "ADSI connection created successfully",
		map[string]interface{}{"connection_key": func() interface{} {
			if apiResp.ConnectionKey != nil {
				return *apiResp.ConnectionKey
			}
			return "unknown"
		}()})

	return apiResp, nil
}

func (r *AdsiConnectionResource) BuildADSIConnector(plan *ADSIConnectorResourceModel, config *ADSIConnectorResourceModel) openapi.ADSIConnector {
	if plan.EntitlementAttribute.IsNull() || plan.EntitlementAttribute.IsUnknown() {
		plan.EntitlementAttribute = types.StringValue("memberOf")
	}

	adsiConn := openapi.ADSIConnector{
		BaseConnector: openapi.BaseConnector{
			Connectiontype:  "ADSI",
			ConnectionName:  plan.ConnectionName.ValueString(),
			Description:     util.StringPointerOrEmpty(plan.Description),
			Defaultsavroles: util.StringPointerOrEmpty(plan.DefaultSavRoles),
			EmailTemplate:   util.StringPointerOrEmpty(plan.EmailTemplate),
		},
		URL:                         plan.URL.ValueString(),
		USERNAME:                    config.Username.ValueString(),
		PASSWORD:                    config.Password.ValueString(),
		CONNECTION_URL:              plan.ConnectionUrl.ValueString(),
		FORESTLIST:                  plan.ForestList.ValueString(),
		PROVISIONING_URL:            util.StringPointerOrEmpty(plan.ProvisioningUrl),
		DEFAULT_USER_ROLE:           util.StringPointerOrEmpty(plan.DefaultUserRole),
		UPDATEUSERJSON:              util.StringPointerOrEmpty(plan.UpdateUserJson),
		ENDPOINTS_FILTER:            util.StringPointerOrEmpty(plan.EndpointsFilter),
		SEARCHFILTER:                util.StringPointerOrEmpty(plan.SearchFilter),
		OBJECTFILTER:                util.StringPointerOrEmpty(plan.ObjectFilter),
		ACCOUNT_ATTRIBUTE:           util.StringPointerOrEmpty(plan.AccountAttribute),
		STATUS_THRESHOLD_CONFIG:     util.StringPointerOrEmpty(plan.StatusThresholdConfig),
		ENTITLEMENT_ATTRIBUTE:       util.StringPointerOrEmpty(plan.EntitlementAttribute),
		USER_ATTRIBUTE:              util.StringPointerOrEmpty(plan.UserAttribute),
		GroupSearchBaseDN:           util.StringPointerOrEmpty(plan.GroupSearchBaseDN),
		CHECKFORUNIQUE:              util.StringPointerOrEmpty(plan.CheckForUnique),
		STATUSKEYJSON:               util.StringPointerOrEmpty(plan.StatusKeyJson),
		GroupImportMapping:          util.StringPointerOrEmpty(plan.GroupImportMapping),
		ImportNestedMembership:      util.StringPointerOrEmpty(plan.ImportNestedMembership),
		PAGE_SIZE:                   util.StringPointerOrEmpty(plan.PageSize),
		ACCOUNTNAMERULE:             util.StringPointerOrEmpty(plan.AccountNameRule),
		CREATEACCOUNTJSON:           util.StringPointerOrEmpty(plan.CreateAccountJson),
		UPDATEACCOUNTJSON:           util.StringPointerOrEmpty(plan.UpdateAccountJson),
		ENABLEACCOUNTJSON:           util.StringPointerOrEmpty(plan.EnableAccountJson),
		DISABLEACCOUNTJSON:          util.StringPointerOrEmpty(plan.DisableAccountJson),
		REMOVEACCOUNTJSON:           util.StringPointerOrEmpty(plan.RemoveAccountJson),
		ADDACCESSJSON:               util.StringPointerOrEmpty(plan.AddAccessJson),
		REMOVEACCESSJSON:            util.StringPointerOrEmpty(plan.RemoveAccessJson),
		RESETANDCHANGEPASSWRDJSON:   util.StringPointerOrEmpty(plan.ResetAndChangePasswrdJson),
		CREATEGROUPJSON:             util.StringPointerOrEmpty(plan.CreateGroupJson),
		UPDATEGROUPJSON:             util.StringPointerOrEmpty(plan.UpdateGroupJson),
		REMOVEGROUPJSON:             util.StringPointerOrEmpty(plan.RemoveGroupJson),
		ADDACCESSENTITLEMENTJSON:    util.StringPointerOrEmpty(plan.AddAccessEntitlementJson),
		CUSTOMCONFIGJSON:            util.StringPointerOrEmpty(plan.CustomConfigJson),
		REMOVEACCESSENTITLEMENTJSON: util.StringPointerOrEmpty(plan.RemoveAccessEntitlementJson),
		CREATESERVICEACCOUNTJSON:    util.StringPointerOrEmpty(plan.CreateServiceAccountJson),
		UPDATESERVICEACCOUNTJSON:    util.StringPointerOrEmpty(plan.UpdateServiceAccountJson),
		REMOVESERVICEACCOUNTJSON:    util.StringPointerOrEmpty(plan.RemoveServiceAccountJson),
		PAM_CONFIG:                  util.StringPointerOrEmpty(plan.PamConfig),
		MODIFYUSERDATAJSON:          util.StringPointerOrEmpty(plan.ModifyUserDataJson),
	}

	if plan.VaultConnection.ValueString() != "" {
		adsiConn.BaseConnector.VaultConnection = util.SafeStringConnector(plan.VaultConnection.ValueString())
		adsiConn.BaseConnector.VaultConfiguration = util.SafeStringConnector(plan.VaultConfiguration.ValueString())
		adsiConn.BaseConnector.Saveinvault = util.SafeStringConnector(plan.SaveInVault.ValueString())
	}

	return adsiConn
}

func (r *AdsiConnectionResource) UpdateModelFromCreateResponse(plan *ADSIConnectorResourceModel, apiResp *openapi.CreateOrUpdateResponse) {
	plan.ID = types.StringValue(fmt.Sprintf("%d", *apiResp.ConnectionKey))
	plan.ConnectionKey = types.Int64Value(int64(*apiResp.ConnectionKey))

	// Update all optional fields to maintain state
	plan.Description = util.SafeStringDatasource(plan.Description.ValueStringPointer())
	plan.DefaultSavRoles = util.SafeStringDatasource(plan.DefaultSavRoles.ValueStringPointer())
	plan.EmailTemplate = util.SafeStringDatasource(plan.EmailTemplate.ValueStringPointer())
	plan.ProvisioningUrl = util.SafeStringDatasource(plan.ProvisioningUrl.ValueStringPointer())
	plan.DefaultUserRole = util.SafeStringDatasource(plan.DefaultUserRole.ValueStringPointer())
	plan.UpdateUserJson = util.SafeStringDatasource(plan.UpdateUserJson.ValueStringPointer())
	plan.EndpointsFilter = util.SafeStringDatasource(plan.EndpointsFilter.ValueStringPointer())
	plan.SearchFilter = util.SafeStringDatasource(plan.SearchFilter.ValueStringPointer())
	plan.ObjectFilter = util.SafeStringDatasource(plan.ObjectFilter.ValueStringPointer())
	plan.AccountAttribute = util.SafeStringDatasource(plan.AccountAttribute.ValueStringPointer())
	plan.StatusThresholdConfig = util.SafeStringDatasource(plan.StatusThresholdConfig.ValueStringPointer())
	plan.EntitlementAttribute = util.SafeStringDatasource(plan.EntitlementAttribute.ValueStringPointer())
	plan.UserAttribute = util.SafeStringDatasource(plan.UserAttribute.ValueStringPointer())
	plan.GroupSearchBaseDN = util.SafeStringDatasource(plan.GroupSearchBaseDN.ValueStringPointer())
	plan.CheckForUnique = util.SafeStringDatasource(plan.CheckForUnique.ValueStringPointer())
	plan.StatusKeyJson = util.SafeStringDatasource(plan.StatusKeyJson.ValueStringPointer())
	plan.GroupImportMapping = util.SafeStringDatasource(plan.GroupImportMapping.ValueStringPointer())
	plan.ImportNestedMembership = util.SafeStringDatasource(plan.ImportNestedMembership.ValueStringPointer())
	plan.PageSize = util.SafeStringDatasource(plan.PageSize.ValueStringPointer())
	plan.AccountNameRule = util.SafeStringDatasource(plan.AccountNameRule.ValueStringPointer())
	plan.CreateAccountJson = util.SafeStringDatasource(plan.CreateAccountJson.ValueStringPointer())
	plan.UpdateAccountJson = util.SafeStringDatasource(plan.UpdateAccountJson.ValueStringPointer())
	plan.EnableAccountJson = util.SafeStringDatasource(plan.EnableAccountJson.ValueStringPointer())
	plan.DisableAccountJson = util.SafeStringDatasource(plan.DisableAccountJson.ValueStringPointer())
	plan.RemoveAccountJson = util.SafeStringDatasource(plan.RemoveAccountJson.ValueStringPointer())
	plan.AddAccessJson = util.SafeStringDatasource(plan.AddAccessJson.ValueStringPointer())
	plan.RemoveAccessJson = util.SafeStringDatasource(plan.RemoveAccessJson.ValueStringPointer())
	plan.ResetAndChangePasswrdJson = util.SafeStringDatasource(plan.ResetAndChangePasswrdJson.ValueStringPointer())
	plan.CreateGroupJson = util.SafeStringDatasource(plan.CreateGroupJson.ValueStringPointer())
	plan.UpdateGroupJson = util.SafeStringDatasource(plan.UpdateGroupJson.ValueStringPointer())
	plan.RemoveGroupJson = util.SafeStringDatasource(plan.RemoveGroupJson.ValueStringPointer())
	plan.AddAccessEntitlementJson = util.SafeStringDatasource(plan.AddAccessEntitlementJson.ValueStringPointer())
	plan.CustomConfigJson = util.SafeStringDatasource(plan.CustomConfigJson.ValueStringPointer())
	plan.RemoveAccessEntitlementJson = util.SafeStringDatasource(plan.RemoveAccessEntitlementJson.ValueStringPointer())
	plan.CreateServiceAccountJson = util.SafeStringDatasource(plan.CreateServiceAccountJson.ValueStringPointer())
	plan.UpdateServiceAccountJson = util.SafeStringDatasource(plan.UpdateServiceAccountJson.ValueStringPointer())
	plan.RemoveServiceAccountJson = util.SafeStringDatasource(plan.RemoveServiceAccountJson.ValueStringPointer())
	plan.PamConfig = util.SafeStringDatasource(plan.PamConfig.ValueStringPointer())
	plan.ModifyUserDataJson = util.SafeStringDatasource(plan.ModifyUserDataJson.ValueStringPointer())

	plan.Msg = types.StringValue(util.SafeDeref(apiResp.Msg))
	plan.ErrorCode = types.StringValue(util.SafeDeref(apiResp.ErrorCode))
}

func (r *AdsiConnectionResource) ReadADSIConnection(ctx context.Context, connectionName string) (*openapi.GetConnectionDetailsResponse, error) {
	opCtx := errorsutil.CreateOperationContext(errorsutil.ConnectorTypeADSI, "read", connectionName)

	// Create logging context (separate from API context)
	logCtx := opCtx.AddContextToLogger(ctx)

	opCtx.LogOperationStart(logCtx, "Starting ADSI connection read operation")

	// Execute read operation with retry logic
	var apiResp *openapi.GetConnectionDetailsResponse

	err := r.provider.AuthenticatedAPICallWithRetry(ctx, "read_adsi_connection", func(token string) error {
		connectionOps := r.connectionFactory.CreateConnectionOperations(r.client.APIBaseURL(), token)
		resp, httpResp, err := connectionOps.GetConnectionDetails(ctx, connectionName)
		if httpResp != nil && httpResp.StatusCode == 401 {
			return fmt.Errorf("401 unauthorized")
		}
		apiResp = resp
		return err
	})

	if err != nil {
		errorCode := adsiErrorCodes.ReadFailed()
		opCtx.LogOperationError(logCtx, "Failed to read ADSI connection", errorCode, err)
		return nil, errorsutil.CreateStandardError(errorsutil.ConnectorTypeADSI, errorCode, "read", connectionName, err)
	}

	if err := r.ValidateADSIConnectionResponse(apiResp); err != nil {
		errorCode := adsiErrorCodes.APIError()
		opCtx.LogOperationError(ctx, "Invalid connection type for ADSI datasource", errorCode, err)
		return nil, fmt.Errorf("[%s] Unable to verify connection type for connection %q. The provider could not determine the type of this connection. Please ensure the connection name is correct and belongs to a supported connector type", errorCode, connectionName)
	}

	if apiResp != nil && apiResp.ADSIConnectionResponse != nil && apiResp.ADSIConnectionResponse.Errorcode != nil && *apiResp.ADSIConnectionResponse.Errorcode != 0 {
		apiErr := fmt.Errorf("API returned error code %d: %s", *apiResp.ADSIConnectionResponse.Errorcode, errorsutil.SanitizeMessage(apiResp.ADSIConnectionResponse.Msg))
		errorCode := adsiErrorCodes.APIError()
		opCtx.LogOperationError(ctx, "ADSI connection read failed with API error", errorCode, apiErr,
			map[string]interface{}{
				"api_error_code": *apiResp.ADSIConnectionResponse.Errorcode,
				"message":        errorsutil.SanitizeMessage(apiResp.ADSIConnectionResponse.Msg),
			})
		return nil, errorsutil.CreateStandardError(errorsutil.ConnectorTypeADSI, errorCode, "read", connectionName, apiErr)
	}

	opCtx.LogOperationEnd(logCtx, "ADSI connection read completed successfully",
		map[string]interface{}{"connection_key": func() interface{} {
			if apiResp.ADSIConnectionResponse != nil && apiResp.ADSIConnectionResponse.Connectionkey != nil {
				return *apiResp.ADSIConnectionResponse.Connectionkey
			}
			return "unknown"
		}()})

	return apiResp, nil
}

func (r *AdsiConnectionResource) UpdateModelFromReadResponse(state *ADSIConnectorResourceModel, apiResp *openapi.GetConnectionDetailsResponse) {
	state.ConnectionKey = types.Int64Value(int64(*apiResp.ADSIConnectionResponse.Connectionkey))
	state.ID = types.StringValue(fmt.Sprintf("%d", *apiResp.ADSIConnectionResponse.Connectionkey))

	// Update all fields from API response
	state.ConnectionName = util.SafeStringDatasource(apiResp.ADSIConnectionResponse.Connectionname)
	state.Description = util.SafeStringDatasource(apiResp.ADSIConnectionResponse.Description)
	state.DefaultSavRoles = util.SafeStringDatasource(apiResp.ADSIConnectionResponse.Defaultsavroles)
	state.EmailTemplate = util.SafeStringDatasource(apiResp.ADSIConnectionResponse.Emailtemplate)
	state.ImportNestedMembership = util.SafeStringDatasource(apiResp.ADSIConnectionResponse.Connectionattributes.ImportNestedMembership)
	state.CreateAccountJson = util.SafeStringDatasource(apiResp.ADSIConnectionResponse.Connectionattributes.CREATEACCOUNTJSON)
	state.EndpointsFilter = util.SafeStringDatasource(apiResp.ADSIConnectionResponse.Connectionattributes.ENDPOINTS_FILTER)
	state.DisableAccountJson = util.SafeStringDatasource(apiResp.ADSIConnectionResponse.Connectionattributes.DISABLEACCOUNTJSON)
	state.RemoveAccessEntitlementJson = util.SafeStringDatasource(apiResp.ADSIConnectionResponse.Connectionattributes.REMOVEACCESSENTITLEMENTJSON)
	state.GroupSearchBaseDN = util.SafeStringDatasource(apiResp.ADSIConnectionResponse.Connectionattributes.GroupSearchBaseDN)
	state.StatusKeyJson = util.SafeStringDatasource(apiResp.ADSIConnectionResponse.Connectionattributes.STATUSKEYJSON)
	state.DefaultUserRole = util.SafeStringDatasource(apiResp.ADSIConnectionResponse.Connectionattributes.DEFAULT_USER_ROLE)
	state.Username = util.SafeStringDatasource(apiResp.ADSIConnectionResponse.Connectionattributes.USERNAME)
	state.UpdateServiceAccountJson = util.SafeStringDatasource(apiResp.ADSIConnectionResponse.Connectionattributes.UPDATESERVICEACCOUNTJSON)
	state.AddAccessJson = util.SafeStringDatasource(apiResp.ADSIConnectionResponse.Connectionattributes.ADDACCESSJSON)
	state.CreateServiceAccountJson = util.SafeStringDatasource(apiResp.ADSIConnectionResponse.Connectionattributes.CREATESERVICEACCOUNTJSON)
	state.AccountNameRule = util.SafeStringDatasource(apiResp.ADSIConnectionResponse.Connectionattributes.ACCOUNTNAMERULE)
	state.ConnectionUrl = util.SafeStringDatasource(apiResp.ADSIConnectionResponse.Connectionattributes.CONNECTION_URL)
	state.AccountAttribute = util.SafeStringDatasource(apiResp.ADSIConnectionResponse.Connectionattributes.ACCOUNT_ATTRIBUTE)
	state.PamConfig = util.SafeStringDatasource(apiResp.ADSIConnectionResponse.Connectionattributes.PAM_CONFIG)
	state.PageSize = util.SafeStringDatasource(apiResp.ADSIConnectionResponse.Connectionattributes.PAGE_SIZE)
	state.SearchFilter = util.SafeStringDatasource(apiResp.ADSIConnectionResponse.Connectionattributes.SEARCHFILTER)
	state.UpdateGroupJson = util.SafeStringDatasource(apiResp.ADSIConnectionResponse.Connectionattributes.UPDATEGROUPJSON)
	state.CreateGroupJson = util.SafeStringDatasource(apiResp.ADSIConnectionResponse.Connectionattributes.CREATEGROUPJSON)
	state.EntitlementAttribute = util.SafeStringDatasource(apiResp.ADSIConnectionResponse.Connectionattributes.ENTITLEMENT_ATTRIBUTE)
	state.CheckForUnique = util.SafeStringDatasource(apiResp.ADSIConnectionResponse.Connectionattributes.CHECKFORUNIQUE)
	state.RemoveServiceAccountJson = util.SafeStringDatasource(apiResp.ADSIConnectionResponse.Connectionattributes.REMOVESERVICEACCOUNTJSON)
	state.UpdateUserJson = util.SafeStringDatasource(apiResp.ADSIConnectionResponse.Connectionattributes.UPDATEUSERJSON)
	state.URL = util.SafeStringDatasource(apiResp.ADSIConnectionResponse.Connectionattributes.URL)
	state.CustomConfigJson = util.SafeStringDatasource(apiResp.ADSIConnectionResponse.Connectionattributes.CUSTOMCONFIGJSON)
	state.StatusThresholdConfig = util.SafeStringDatasource(apiResp.ADSIConnectionResponse.Connectionattributes.STATUS_THRESHOLD_CONFIG)
	state.GroupImportMapping = util.SafeStringDatasource(apiResp.ADSIConnectionResponse.Connectionattributes.GroupImportMapping)
	state.ProvisioningUrl = util.SafeStringDatasource(apiResp.ADSIConnectionResponse.Connectionattributes.PROVISIONING_URL)
	state.RemoveGroupJson = util.SafeStringDatasource(apiResp.ADSIConnectionResponse.Connectionattributes.REMOVEGROUPJSON)
	state.RemoveAccessJson = util.SafeStringDatasource(apiResp.ADSIConnectionResponse.Connectionattributes.REMOVEACCESSJSON)
	state.ResetAndChangePasswrdJson = util.SafeStringDatasource(apiResp.ADSIConnectionResponse.Connectionattributes.RESETANDCHANGEPASSWRDJSON)
	state.UserAttribute = util.SafeStringDatasource(apiResp.ADSIConnectionResponse.Connectionattributes.USER_ATTRIBUTE)
	state.AddAccessEntitlementJson = util.SafeStringDatasource(apiResp.ADSIConnectionResponse.Connectionattributes.ADDACCESSENTITLEMENTJSON)
	state.ModifyUserDataJson = util.SafeStringDatasource(apiResp.ADSIConnectionResponse.Connectionattributes.MODIFYUSERDATAJSON)
	state.EnableAccountJson = util.SafeStringDatasource(apiResp.ADSIConnectionResponse.Connectionattributes.ENABLEACCOUNTJSON)
	state.ForestList = util.SafeStringDatasource(apiResp.ADSIConnectionResponse.Connectionattributes.FORESTLIST)
	state.ObjectFilter = util.SafeStringDatasource(apiResp.ADSIConnectionResponse.Connectionattributes.OBJECTFILTER)
	state.UpdateAccountJson = util.SafeStringDatasource(apiResp.ADSIConnectionResponse.Connectionattributes.UPDATEACCOUNTJSON)
	state.RemoveAccountJson = util.SafeStringDatasource(apiResp.ADSIConnectionResponse.Connectionattributes.REMOVEACCOUNTJSON)

	apiMessage := util.SafeDeref(apiResp.ADSIConnectionResponse.Msg)
	if apiMessage == "success" {
		state.Msg = types.StringValue("Connection Successful")
	} else {
		state.Msg = types.StringValue(apiMessage)
	}
	state.ErrorCode = util.Int32PtrToTFString(apiResp.ADSIConnectionResponse.Errorcode)
}

func (r *AdsiConnectionResource) UpdateADSIConnection(ctx context.Context, plan *ADSIConnectorResourceModel, config *ADSIConnectorResourceModel) (*openapi.CreateOrUpdateResponse, error) {
	connectionName := plan.ConnectionName.ValueString()
	opCtx := errorsutil.CreateOperationContext(errorsutil.ConnectorTypeADSI, "update", connectionName)

	// Create logging context (separate from API context)
	logCtx := opCtx.AddContextToLogger(ctx)

	opCtx.LogOperationStart(logCtx, "Starting ADSI connection update")

	// Build ADSI connection update request
	tflog.Debug(logCtx, "Building ADSI connection update request")

	adsiConn := r.BuildADSIConnector(plan, config)
	if plan.VaultConnection.ValueString() == "" {
		emptyStr := ""
		adsiConn.BaseConnector.VaultConnection = &emptyStr
		adsiConn.BaseConnector.VaultConfiguration = &emptyStr
		adsiConn.BaseConnector.Saveinvault = &emptyStr
	}

	updateReq := openapi.CreateOrUpdateRequest{
		ADSIConnector: &adsiConn,
	}

	var apiResp *openapi.CreateOrUpdateResponse

	// Execute update operation with retry logic
	tflog.Debug(logCtx, "Executing update operation")
	err := r.provider.AuthenticatedAPICallWithRetry(ctx, "update_adsi_connection", func(token string) error {
		connectionOps := r.connectionFactory.CreateConnectionOperations(r.client.APIBaseURL(), token)
		resp, httpResp, err := connectionOps.CreateOrUpdateConnection(ctx, updateReq)
		if httpResp != nil && httpResp.StatusCode == 401 {
			return fmt.Errorf("401 unauthorized")
		}
		apiResp = resp
		return err
	})

	if err != nil {
		errorCode := adsiErrorCodes.UpdateFailed()
		opCtx.LogOperationError(logCtx, "Failed to update ADSI connection", errorCode, err)
		return nil, errorsutil.CreateStandardError(errorsutil.ConnectorTypeADSI, errorCode, "update", connectionName, err)
	}

	if apiResp != nil && apiResp.ErrorCode != nil && *apiResp.ErrorCode != "0" {
		apiErr := fmt.Errorf("API returned error code %s: %s", *apiResp.ErrorCode, errorsutil.SanitizeMessage(apiResp.Msg))
		errorCode := adsiErrorCodes.APIError()
		opCtx.LogOperationError(logCtx, "ADSI connection update failed with API error", errorCode, apiErr,
			map[string]interface{}{
				"api_error_code": *apiResp.ErrorCode,
				"message":        errorsutil.SanitizeMessage(apiResp.Msg),
			})
		return nil, errorsutil.CreateStandardError(errorsutil.ConnectorTypeADSI, errorCode, "update", connectionName, apiErr)
	}

	opCtx.LogOperationEnd(logCtx, "ADSI connection updated successfully",
		map[string]interface{}{"connection_key": func() interface{} {
			if apiResp.ConnectionKey != nil {
				return *apiResp.ConnectionKey
			}
			return "unknown"
		}()})

	return apiResp, nil
}

func (r *AdsiConnectionResource) ValidateADSIConnectionResponse(apiResp *openapi.GetConnectionDetailsResponse) error {
	if apiResp != nil && apiResp.ADSIConnectionResponse == nil {
		return fmt.Errorf("verify the connection type - ADSI connection response is nil")
	}
	return nil
}

func (r *AdsiConnectionResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan, config ADSIConnectorResourceModel

	opCtx := errorsutil.CreateOperationContext(errorsutil.ConnectorTypeADSI, "terraform_create", "")
	ctx = opCtx.AddContextToLogger(ctx)

	opCtx.LogOperationStart(ctx, "Starting ADSI connection resource creation")

	// Extract plan from request
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		errorCode := adsiErrorCodes.PlanExtraction()
		opCtx.LogOperationError(ctx, "Failed to get plan from request", errorCode,
			fmt.Errorf("plan extraction failed"))
		resp.Diagnostics.AddError(
			errorsutil.GetErrorMessage(errorsutil.ErrPlanExtraction),
			fmt.Sprintf("[%s] Unable to extract Terraform plan from request", errorCode),
		)
		return
	}

	connectionName := plan.ConnectionName.ValueString()
	// Update operation context with connection name
	opCtx.ConnectionName = connectionName
	ctx = opCtx.AddContextToLogger(ctx)

	//Extract config from request
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		errorCode := adsiErrorCodes.ConfigExtraction()
		opCtx.LogOperationError(ctx, "Failed to get config from request", errorCode,
			fmt.Errorf("config extraction failed"))
		resp.Diagnostics.AddError(
			errorsutil.GetErrorMessage(errorsutil.ErrConfigExtraction),
			fmt.Sprintf("[%s] Unable to extract Terraform configuration from request for connection '%s'", errorCode, connectionName),
		)
		return
	}

	// Use interface pattern instead of direct API client creation
	apiResp, err := r.CreateADSIConnection(ctx, &plan, &config)
	if err != nil {
		opCtx.LogOperationError(ctx, "ADSI connection creation failed", "", err)
		resp.Diagnostics.AddError(
			"ADSI Connection Creation Failed",
			err.Error(),
		)
		return
	}

	// Update model from create response
	r.UpdateModelFromCreateResponse(&plan, apiResp)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)

	opCtx.LogOperationEnd(ctx, "ADSI connection resource created successfully",
		map[string]interface{}{"connection_key": plan.ConnectionKey.ValueInt64()})
}

func (r *AdsiConnectionResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state ADSIConnectorResourceModel

	opCtx := errorsutil.CreateOperationContext(errorsutil.ConnectorTypeADSI, "terraform_read", "")
	ctx = opCtx.AddContextToLogger(ctx)

	opCtx.LogOperationStart(ctx, "Starting ADSI connection resource read")

	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		errorCode := adsiErrorCodes.StateExtraction()
		opCtx.LogOperationError(ctx, "Failed to get state from request", errorCode,
			fmt.Errorf("state extraction failed"))
		resp.Diagnostics.AddError(
			errorsutil.GetErrorMessage(errorsutil.ErrStateExtraction),
			fmt.Sprintf("[%s] Unable to extract Terraform state from request", errorCode),
		)
		return
	}

	connectionName := state.ConnectionName.ValueString()
	// Update operation context with connection name
	opCtx.ConnectionName = connectionName
	ctx = opCtx.AddContextToLogger(ctx)

	// Use interface pattern instead of direct API client creation
	apiResp, err := r.ReadADSIConnection(ctx, connectionName)
	if err != nil {
		opCtx.LogOperationError(ctx, "ADSI connection read failed", "", err)
		resp.Diagnostics.AddError(
			"ADSI Connection Read Failed",
			err.Error(),
		)
		return
	}

	// Update model from read response
	r.UpdateModelFromReadResponse(&state, apiResp)

	stateDiagnostics := resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(stateDiagnostics...)
	if resp.Diagnostics.HasError() {
		errorCode := adsiErrorCodes.StateUpdate()
		opCtx.LogOperationError(ctx, "Failed to set state", errorCode,
			fmt.Errorf("state update failed"))
		resp.Diagnostics.AddError(
			errorsutil.GetErrorMessage(errorsutil.ErrStateUpdate),
			fmt.Sprintf("[%s] Unable to update Terraform state for connection '%s'", errorCode, connectionName),
		)
		return
	}

	opCtx.LogOperationEnd(ctx, "ADSI connection resource read completed successfully")
}

func (r *AdsiConnectionResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state, config ADSIConnectorResourceModel

	opCtx := errorsutil.CreateOperationContext(errorsutil.ConnectorTypeADSI, "terraform_update", "")
	ctx = opCtx.AddContextToLogger(ctx)

	opCtx.LogOperationStart(ctx, "Starting ADSI connection resource update")

	// Extract state from request
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		errorCode := adsiErrorCodes.StateExtraction()
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
		errorCode := adsiErrorCodes.PlanExtraction()
		opCtx.LogOperationError(ctx, "Failed to get plan from request", errorCode,
			fmt.Errorf("plan extraction failed"))
		resp.Diagnostics.AddError(
			errorsutil.GetErrorMessage(errorsutil.ErrPlanExtraction),
			fmt.Sprintf("[%s] Unable to extract Terraform plan from request for connection '%s'", errorCode, state.ConnectionName.ValueString()),
		)
		return
	}

	//Extract config from request
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		errorCode := adsiErrorCodes.ConfigExtraction()
		opCtx.LogOperationError(ctx, "Failed to get config from request", errorCode,
			fmt.Errorf("config extraction failed"))
		resp.Diagnostics.AddError(
			errorsutil.GetErrorMessage(errorsutil.ErrConfigExtraction),
			fmt.Sprintf("[%s] Unable to extract Terraform configuration from request for connection '%s'", errorCode, plan.ConnectionName.ValueString()),
		)
		return
	}

	// Validate that connection name cannot be updated
	if plan.ConnectionName.ValueString() != state.ConnectionName.ValueString() {
		errorCode := adsiErrorCodes.NameImmutable()
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

	connectionName := plan.ConnectionName.ValueString()
	// Update operation context with connection name
	opCtx.ConnectionName = connectionName
	ctx = opCtx.AddContextToLogger(ctx)

	// Use interface pattern instead of direct API client creation
	_, err := r.UpdateADSIConnection(ctx, &plan, &config)
	if err != nil {
		opCtx.LogOperationError(ctx, "ADSI connection update failed", "", err)
		resp.Diagnostics.AddError(
			"ADSI Connection Update Failed",
			err.Error(),
		)
		return
	}

	// Read the updated connection to get the latest state
	getResp, err := r.ReadADSIConnection(ctx, connectionName)
	if err != nil {
		opCtx.LogOperationError(ctx, "Failed to read updated ADSI connection", "", err)
		resp.Diagnostics.AddError(
			"ADSI Connection Post-Update Read Failed",
			err.Error(),
		)
		return
	}

	// Update model from read response
	r.UpdateModelFromReadResponse(&plan, getResp)

	stateUpdateDiagnostics := resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(stateUpdateDiagnostics...)
	if resp.Diagnostics.HasError() {
		errorCode := adsiErrorCodes.StateUpdate()
		opCtx.LogOperationError(ctx, "Failed to update state after successful update", errorCode,
			fmt.Errorf("state update failed"))
		resp.Diagnostics.AddError(
			errorsutil.GetErrorMessage(errorsutil.ErrStateUpdate),
			fmt.Sprintf("[%s] Unable to update Terraform state after successful update for connection '%s'", errorCode, connectionName),
		)
		return
	}

	opCtx.LogOperationEnd(ctx, "ADSI connection resource updated successfully",
		map[string]interface{}{"connection_key": plan.ConnectionKey.ValueInt64()})
}

func (r *AdsiConnectionResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
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

func (r *AdsiConnectionResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	// Importing an ADSI connection resource requires the connection name
	connectionName := req.ID
	opCtx := errorsutil.CreateOperationContext(errorsutil.ConnectorTypeADSI, "terraform_import", connectionName)
	ctx = opCtx.AddContextToLogger(ctx)

	opCtx.LogOperationStart(ctx, "Starting ADSI connection resource import")

	// Retrieve import ID and save to connection_name attribute
	resource.ImportStatePassthroughID(ctx, path.Root("connection_name"), req, resp)

	opCtx.LogOperationEnd(ctx, "ADSI connection resource import completed successfully",
		map[string]interface{}{"import_id": connectionName})
}
