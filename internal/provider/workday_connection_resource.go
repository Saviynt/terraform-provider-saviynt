// Copyright (c) 2025 Saviynt Inc.
// SPDX-License-Identifier: MPL-2.0

// saviynt_workday_connection_resource manages Workday connectors in the Saviynt Security Manager.
// The resource implements the full Terraform lifecycle:
//   - Create: provisions a new Workday connector using the supplied configuration.
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

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &WorkdayConnectionResource{}
var _ resource.ResourceWithImportState = &WorkdayConnectionResource{}

// Initialize error codes for Workday Connection operations
var workdayErrorCodes = errorsutil.NewConnectorErrorCodes(errorsutil.ConnectorTypeWorkday)

type WorkdayConnectorResourceModel struct {
	BaseConnectorResourceModel
	ID                     types.String `tfsdk:"id"`
	UsersLastImportTime    types.String `tfsdk:"users_last_import_time"`
	AccountsLastImportTime types.String `tfsdk:"accounts_last_import_time"`
	AccessLastImportTime   types.String `tfsdk:"access_last_import_time"`
	BaseURL                types.String `tfsdk:"base_url"`
	APIVersion             types.String `tfsdk:"api_version"`
	TenantName             types.String `tfsdk:"tenant_name"`
	ReportOwner            types.String `tfsdk:"report_owner"`
	UseOAuth               types.String `tfsdk:"use_oauth"`
	IncludeReferenceDesc   types.String `tfsdk:"include_reference_descriptors"`
	UseEnhancedOrgRole     types.String `tfsdk:"use_enhanced_orgrole"`
	UseX509AuthForSOAP     types.String `tfsdk:"use_x509auth_for_soap"`
	X509Key                types.String `tfsdk:"x509_key"`
	X509Cert               types.String `tfsdk:"x509_cert"`
	Username               types.String `tfsdk:"username"`
	Password               types.String `tfsdk:"password"`
	ClientID               types.String `tfsdk:"client_id"`
	ClientSecret           types.String `tfsdk:"client_secret"`
	RefreshToken           types.String `tfsdk:"refresh_token"`
	PageSize               types.String `tfsdk:"page_size"`
	UserImportPayload      types.String `tfsdk:"user_import_payload"`
	UserImportMapping      types.String `tfsdk:"user_import_mapping"`
	AccountImportPayload   types.String `tfsdk:"account_import_payload"`
	AccountImportMapping   types.String `tfsdk:"account_import_mapping"`
	AccessImportList       types.String `tfsdk:"access_import_list"`
	RAASMappingJSON        types.String `tfsdk:"raas_mapping_json"`
	AccessImportMapping    types.String `tfsdk:"access_import_mapping"`
	OrgRoleImportPayload   types.String `tfsdk:"orgrole_import_payload"`
	StatusKeyJSON          types.String `tfsdk:"status_key_json"`
	UserAttributeJSON      types.String `tfsdk:"userattributejson"`
	CustomConfig           types.String `tfsdk:"custom_config"`
	PAMConfig              types.String `tfsdk:"pam_config"`
	ModifyUserDataJSON     types.String `tfsdk:"modify_user_data_json"`
	StatusThresholdConfig  types.String `tfsdk:"status_threshold_config"`
	CreateAccountPayload   types.String `tfsdk:"create_account_payload"`
	UpdateAccountPayload   types.String `tfsdk:"update_account_payload"`
	UpdateUserPayload      types.String `tfsdk:"update_user_payload"`
	AssignOrgRolePayload   types.String `tfsdk:"assign_orgrole_payload"`
	RemoveOrgRolePayload   types.String `tfsdk:"remove_orgrole_payload"`
}

// workdayConnectionResource implements the resource.Resource interface.
type WorkdayConnectionResource struct {
	client            client.SaviyntClientInterface
	token             string
	provider          client.SaviyntProviderInterface
	connectionFactory client.ConnectionFactoryInterface
}

// NewTestConnectionResource returns a new instance of testConnectionResource.
func NewWorkdayConnectionResource() resource.Resource {
	return &WorkdayConnectionResource{
		connectionFactory: &client.DefaultConnectionFactory{},
	}
}

func NewWorkdayConnectionResourceWithFactory(factory client.ConnectionFactoryInterface) resource.Resource {
	return &WorkdayConnectionResource{
		connectionFactory: factory,
	}
}

func (r *WorkdayConnectionResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = "saviynt_workday_connection_resource"
}

func WorkdayConnectorResourceSchema() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:    true,
			Description: "Resource ID.",
		},
		"users_last_import_time": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "Property for USERS_LAST_IMPORT_TIME.",
		},
		"accounts_last_import_time": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "Property for ACCOUNTS_LAST_IMPORT_TIME.",
		},
		"access_last_import_time": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "Property for ACCESS_LAST_IMPORT_TIME.",
		},
		"base_url": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "Base URL of the Workday tenant instance.",
		},
		"api_version": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "Version of the SOAP API used for the connection.",
		},
		"tenant_name": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "The name of your tenant.",
		},
		"report_owner": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "Account name of the report owner used to build default RaaS URLs.",
		},
		"use_oauth": schema.StringAttribute{
			Required:    true,
			Description: "Whether to use OAuth authentication.Values can be TRUE/FALSE",
		},
		"include_reference_descriptors": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "Include descriptor attribute in response if set to TRUE.",
		},
		"use_enhanced_orgrole": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "Set TRUE to utilize enhanced Organizational Role setup.",
		},
		"use_x509auth_for_soap": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "Set TRUE to use certificate-based authentication for SOAP.",
		},
		"x509_key": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Sensitive:   true,
			Description: "Private key for x509-based SOAP authentication.",
		},
		"x509_cert": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Sensitive:   true,
			Description: "Certificate for x509-based SOAP authentication.",
		},
		"username": schema.StringAttribute{
			Optional:    true,
			WriteOnly:   true,
			Description: "Username for SOAP authentication.",
		},
		"password": schema.StringAttribute{
			Optional:    true,
			WriteOnly:   true,
			Description: "Password for SOAP authentication.",
		},
		"client_id": schema.StringAttribute{
			Optional:    true,
			WriteOnly:   true,
			Description: "OAuth client ID.",
		},
		"client_secret": schema.StringAttribute{
			Optional:    true,
			WriteOnly:   true,
			Description: "OAuth client secret.",
		},
		"refresh_token": schema.StringAttribute{
			Optional:    true,
			WriteOnly:   true,
			Description: "OAuth refresh token.",
		},
		"page_size": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "Number of objects to return per page during import.",
		},
		"user_import_payload": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "Request payload for importing users.",
		},
		"user_import_mapping": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "Mapping configuration for user import.",
		},
		"account_import_payload": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "Request payload for importing accounts.",
		},
		"account_import_mapping": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "Mapping configuration for account import.",
		},
		"access_import_list": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "Comma-separated list of access types to import.",
		},
		"raas_mapping_json": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "Overrides default report mapping for RaaS.",
		},
		"access_import_mapping": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "Additional access attribute mapping for Workday access objects.",
		},
		"orgrole_import_payload": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "Custom SOAP body for organization role import.",
		},
		"status_key_json": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "Mapping of user status.",
		},
		"userattributejson": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "Specifies which job-related attributes are stored as user attributes.",
		},
		"custom_config": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "Custom configuration for Workday connector.",
		},
		"pam_config": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "Privileged Access Management configuration.",
		},
		"modify_user_data_json": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "Payload for modifying user data.",
		},
		"status_threshold_config": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "Config for reading and importing status of account and entitlement.",
		},
		"create_account_payload": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "Payload for creating an account.",
		},
		"update_account_payload": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "Payload for updating an account.",
		},
		"update_user_payload": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "Payload for updating a user.",
		},
		"assign_orgrole_payload": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "Payload for assigning org role.",
		},
		"remove_orgrole_payload": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "Payload for removing org role.",
		},
	}
}

func (r *WorkdayConnectionResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: util.WorkdayConnDescription,
		Attributes:  connectionsutil.MergeResourceAttributes(BaseConnectorResourceSchema(), WorkdayConnectorResourceSchema()),
	}
}

func (r *WorkdayConnectionResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	opCtx := errorsutil.CreateOperationContext(errorsutil.ConnectorTypeWorkday, "configure", "")
	ctx = opCtx.AddContextToLogger(ctx)

	opCtx.LogOperationStart(ctx, "Starting Workday connection resource configuration")

	// Check if provider data is available.
	if req.ProviderData == nil {
		tflog.Debug(ctx, "ProviderData is nil, returning early")
		opCtx.LogOperationEnd(ctx, "Workday connection resource configuration completed - no provider data")
		return
	}

	// Cast provider data to your provider type.
	prov, ok := req.ProviderData.(*SaviyntProvider)
	if !ok {
		errorCode := workdayErrorCodes.ProviderConfig()
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

	opCtx.LogOperationEnd(ctx, "Workday connection resource configured successfully")
}

// SetClient sets the client for testing purposes
func (r *WorkdayConnectionResource) SetClient(client client.SaviyntClientInterface) {
	r.client = client
}

// SetToken sets the token for testing purposes
func (r *WorkdayConnectionResource) SetToken(token string) {
	r.token = token
}

// SetProvider sets the provider for testing purposes
func (r *WorkdayConnectionResource) SetProvider(provider client.SaviyntProviderInterface) {
	r.provider = provider
}

func (r *WorkdayConnectionResource) CreateWorkdayConnection(ctx context.Context, plan *WorkdayConnectorResourceModel, config *WorkdayConnectorResourceModel) (*openapi.CreateOrUpdateResponse, error) {
	connectionName := plan.ConnectionName.ValueString()
	opCtx := errorsutil.CreateOperationContext(errorsutil.ConnectorTypeWorkday, "create", connectionName)

	// Create logging context (separate from API context)
	logCtx := opCtx.AddContextToLogger(ctx)

	opCtx.LogOperationStart(logCtx, "Starting Workday connection creation")

	// Check if connection already exists (idempotency check) with retry logic
	tflog.Debug(logCtx, "Checking if connection already exists")
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
		errorCode := workdayErrorCodes.ReadFailed()
		opCtx.LogOperationError(logCtx, "Failed to check existing connection", errorCode, err)
		return nil, errorsutil.CreateStandardError(errorsutil.ConnectorTypeWorkday, errorCode, "create", connectionName, err)
	}

	if existingResource != nil &&
		existingResource.WorkdayConnectionResponse != nil &&
		existingResource.WorkdayConnectionResponse.Errorcode != nil &&
		*existingResource.WorkdayConnectionResponse.Errorcode == 0 {

		errorCode := workdayErrorCodes.DuplicateName()
		opCtx.LogOperationError(ctx, "Connection name already exists.Please import or use a different name", errorCode,
			fmt.Errorf("duplicate connection name"))
		return nil, errorsutil.CreateStandardError(errorsutil.ConnectorTypeWorkday, errorCode, "create", connectionName, nil)
	}

	// Build Workday connection create request
	tflog.Debug(ctx, "Building Workday connection create request")

	workdayConn := r.BuildWorkdayConnector(plan, config)
	createReq := openapi.CreateOrUpdateRequest{
		WorkdayConnector: &workdayConn,
	}

	// Execute create operation with retry logic
	tflog.Debug(ctx, "Executing create operation")
	var apiResp *openapi.CreateOrUpdateResponse

	err = r.provider.AuthenticatedAPICallWithRetry(ctx, "create_workday_connection", func(token string) error {
		connectionOps := r.connectionFactory.CreateConnectionOperations(r.client.APIBaseURL(), token)
		resp, httpResp, err := connectionOps.CreateOrUpdateConnection(ctx, createReq)
		if httpResp != nil && httpResp.StatusCode == 401 {
			return fmt.Errorf("401 unauthorized")
		}
		apiResp = resp
		return err
	})

	if err != nil {
		errorCode := workdayErrorCodes.CreateFailed()
		opCtx.LogOperationError(ctx, "Failed to create Workday connection", errorCode, err)
		return nil, errorsutil.CreateStandardError(errorsutil.ConnectorTypeWorkday, errorCode, "create", connectionName, err)
	}

	if apiResp != nil && apiResp.ErrorCode != nil && *apiResp.ErrorCode != "0" {
		apiErr := fmt.Errorf("API returned error code %s: %s", *apiResp.ErrorCode, errorsutil.SanitizeMessage(apiResp.Msg))
		errorCode := workdayErrorCodes.APIError()
		opCtx.LogOperationError(ctx, "Workday connection creation failed with API error", errorCode, apiErr,
			map[string]interface{}{
				"api_error_code": *apiResp.ErrorCode,
				"message":        errorsutil.SanitizeMessage(apiResp.Msg),
			})
		return nil, errorsutil.CreateStandardError(errorsutil.ConnectorTypeWorkday, errorCode, "create", connectionName, apiErr)
	}

	opCtx.LogOperationEnd(logCtx, "Workday connection created successfully",
		map[string]interface{}{"connection_key": func() interface{} {
			if apiResp.ConnectionKey != nil {
				return *apiResp.ConnectionKey
			}
			return "unknown"
		}()})

	return apiResp, nil
}

func (r *WorkdayConnectionResource) BuildWorkdayConnector(plan *WorkdayConnectorResourceModel, config *WorkdayConnectorResourceModel) openapi.WorkdayConnector {
	workdayConn := openapi.WorkdayConnector{
		BaseConnector: openapi.BaseConnector{
			//required fields
			Connectiontype: "Workday",
			ConnectionName: plan.ConnectionName.ValueString(),
			//optional fields
			Description:     util.StringPointerOrEmpty(plan.Description),
			Defaultsavroles: util.StringPointerOrEmpty(plan.DefaultSavRoles),
			EmailTemplate:   util.StringPointerOrEmpty(plan.EmailTemplate),
		},
		//required fields
		USE_OAUTH: plan.UseOAuth.ValueString(),
		//optional fields
		USERS_LAST_IMPORT_TIME:        util.StringPointerOrEmpty(plan.UsersLastImportTime),
		ACCOUNTS_LAST_IMPORT_TIME:     util.StringPointerOrEmpty(plan.AccountsLastImportTime),
		ACCESS_LAST_IMPORT_TIME:       util.StringPointerOrEmpty(plan.AccessLastImportTime),
		BASE_URL:                      util.StringPointerOrEmpty(plan.BaseURL),
		API_VERSION:                   util.StringPointerOrEmpty(plan.APIVersion),
		TENANT_NAME:                   util.StringPointerOrEmpty(plan.TenantName),
		REPORT_OWNER:                  util.StringPointerOrEmpty(plan.ReportOwner),
		INCLUDE_REFERENCE_DESCRIPTORS: util.StringPointerOrEmpty(plan.IncludeReferenceDesc),
		USE_ENHANCED_ORGROLE:          util.StringPointerOrEmpty(plan.UseEnhancedOrgRole),
		USEX509AUTHFORSOAP:            util.StringPointerOrEmpty(plan.UseX509AuthForSOAP),
		X509KEY:                       util.StringPointerOrEmpty(plan.X509Key),
		X509CERT:                      util.StringPointerOrEmpty(plan.X509Cert),
		USERNAME:                      util.StringPointerOrEmpty(config.Username),
		PASSWORD:                      util.StringPointerOrEmpty(config.Password),
		CLIENT_ID:                     util.StringPointerOrEmpty(config.ClientID),
		CLIENT_SECRET:                 util.StringPointerOrEmpty(config.ClientSecret),
		REFRESH_TOKEN:                 util.StringPointerOrEmpty(config.RefreshToken),
		PAGE_SIZE:                     util.StringPointerOrEmpty(plan.PageSize),
		USER_IMPORT_PAYLOAD:           util.StringPointerOrEmpty(plan.UserImportPayload),
		USER_IMPORT_MAPPING:           util.StringPointerOrEmpty(plan.UserImportMapping),
		ACCOUNT_IMPORT_PAYLOAD:        util.StringPointerOrEmpty(plan.AccountImportPayload),
		ACCOUNT_IMPORT_MAPPING:        util.StringPointerOrEmpty(plan.AccountImportMapping),
		ACCESS_IMPORT_LIST:            util.StringPointerOrEmpty(plan.AccessImportList),
		RAAS_MAPPING_JSON:             util.StringPointerOrEmpty(plan.RAASMappingJSON),
		ACCESS_IMPORT_MAPPING:         util.StringPointerOrEmpty(plan.AccessImportMapping),
		ORGROLE_IMPORT_PAYLOAD:        util.StringPointerOrEmpty(plan.OrgRoleImportPayload),
		STATUS_KEY_JSON:               util.StringPointerOrEmpty(plan.StatusKeyJSON),
		USERATTRIBUTEJSON:             util.StringPointerOrEmpty(plan.UserAttributeJSON),
		CUSTOM_CONFIG:                 util.StringPointerOrEmpty(plan.CustomConfig),
		PAM_CONFIG:                    util.StringPointerOrEmpty(plan.PAMConfig),
		MODIFYUSERDATAJSON:            util.StringPointerOrEmpty(plan.ModifyUserDataJSON),
		STATUS_THRESHOLD_CONFIG:       util.StringPointerOrEmpty(plan.StatusThresholdConfig),
		CREATE_ACCOUNT_PAYLOAD:        util.StringPointerOrEmpty(plan.CreateAccountPayload),
		UPDATE_ACCOUNT_PAYLOAD:        util.StringPointerOrEmpty(plan.UpdateAccountPayload),
		UPDATE_USER_PAYLOAD:           util.StringPointerOrEmpty(plan.UpdateUserPayload),
		ASSIGN_ORGROLE_PAYLOAD:        util.StringPointerOrEmpty(plan.AssignOrgRolePayload),
		REMOVE_ORGROLE_PAYLOAD:        util.StringPointerOrEmpty(plan.RemoveOrgRolePayload),
	}

	if plan.VaultConnection.ValueString() != "" {
		workdayConn.BaseConnector.VaultConnection = util.SafeStringConnector(plan.VaultConnection.ValueString())
		workdayConn.BaseConnector.VaultConfiguration = util.SafeStringConnector(plan.VaultConfiguration.ValueString())
		workdayConn.BaseConnector.Saveinvault = util.SafeStringConnector(plan.SaveInVault.ValueString())
	}

	return workdayConn
}

func (r *WorkdayConnectionResource) UpdateModelFromCreateResponse(plan *WorkdayConnectorResourceModel, apiResp *openapi.CreateOrUpdateResponse) {
	plan.ID = types.StringValue(fmt.Sprintf("%d", *apiResp.ConnectionKey))
	plan.ConnectionKey = types.Int64Value(int64(*apiResp.ConnectionKey))

	plan.Description = util.SafeStringDatasource(plan.Description.ValueStringPointer())
	plan.DefaultSavRoles = util.SafeStringDatasource(plan.DefaultSavRoles.ValueStringPointer())
	plan.EmailTemplate = util.SafeStringDatasource(plan.EmailTemplate.ValueStringPointer())
	plan.UsersLastImportTime = util.SafeStringDatasource(plan.UsersLastImportTime.ValueStringPointer())
	plan.AccountsLastImportTime = util.SafeStringDatasource(plan.AccountsLastImportTime.ValueStringPointer())
	plan.AccessLastImportTime = util.SafeStringDatasource(plan.AccessLastImportTime.ValueStringPointer())
	plan.BaseURL = util.SafeStringDatasource(plan.BaseURL.ValueStringPointer())
	plan.APIVersion = util.SafeStringDatasource(plan.APIVersion.ValueStringPointer())
	plan.TenantName = util.SafeStringDatasource(plan.TenantName.ValueStringPointer())
	plan.ReportOwner = util.SafeStringDatasource(plan.ReportOwner.ValueStringPointer())
	plan.IncludeReferenceDesc = util.SafeStringDatasource(plan.IncludeReferenceDesc.ValueStringPointer())
	plan.UseEnhancedOrgRole = util.SafeStringDatasource(plan.UseEnhancedOrgRole.ValueStringPointer())
	plan.UseX509AuthForSOAP = util.SafeStringDatasource(plan.UseX509AuthForSOAP.ValueStringPointer())
	plan.X509Key = util.SafeStringDatasource(plan.X509Key.ValueStringPointer())
	plan.X509Cert = util.SafeStringDatasource(plan.X509Cert.ValueStringPointer())
	plan.Username = util.SafeStringDatasource(plan.Username.ValueStringPointer())
	plan.ClientID = util.SafeStringDatasource(plan.ClientID.ValueStringPointer())
	plan.PageSize = util.SafeStringDatasource(plan.PageSize.ValueStringPointer())
	plan.UserImportPayload = util.SafeStringDatasource(plan.UserImportPayload.ValueStringPointer())
	plan.UserImportMapping = util.SafeStringDatasource(plan.UserImportMapping.ValueStringPointer())
	plan.AccountImportPayload = util.SafeStringDatasource(plan.AccountImportPayload.ValueStringPointer())
	plan.AccountImportMapping = util.SafeStringDatasource(plan.AccountImportMapping.ValueStringPointer())
	plan.AccessImportList = util.SafeStringDatasource(plan.AccessImportList.ValueStringPointer())
	plan.RAASMappingJSON = util.SafeStringDatasource(plan.RAASMappingJSON.ValueStringPointer())
	plan.AccessImportMapping = util.SafeStringDatasource(plan.AccessImportMapping.ValueStringPointer())
	plan.OrgRoleImportPayload = util.SafeStringDatasource(plan.OrgRoleImportPayload.ValueStringPointer())
	plan.StatusKeyJSON = util.SafeStringDatasource(plan.StatusKeyJSON.ValueStringPointer())
	plan.UserAttributeJSON = util.SafeStringDatasource(plan.UserAttributeJSON.ValueStringPointer())
	plan.CustomConfig = util.SafeStringDatasource(plan.CustomConfig.ValueStringPointer())
	plan.PAMConfig = util.SafeStringDatasource(plan.PAMConfig.ValueStringPointer())
	plan.ModifyUserDataJSON = util.SafeStringDatasource(plan.ModifyUserDataJSON.ValueStringPointer())
	plan.StatusThresholdConfig = util.SafeStringDatasource(plan.StatusThresholdConfig.ValueStringPointer())
	plan.CreateAccountPayload = util.SafeStringDatasource(plan.CreateAccountPayload.ValueStringPointer())
	plan.UpdateAccountPayload = util.SafeStringDatasource(plan.UpdateAccountPayload.ValueStringPointer())
	plan.UpdateUserPayload = util.SafeStringDatasource(plan.UpdateUserPayload.ValueStringPointer())
	plan.AssignOrgRolePayload = util.SafeStringDatasource(plan.AssignOrgRolePayload.ValueStringPointer())
	plan.RemoveOrgRolePayload = util.SafeStringDatasource(plan.RemoveOrgRolePayload.ValueStringPointer())
	if plan.UseEnhancedOrgRole.IsNull() || plan.UseEnhancedOrgRole.ValueString() == "" {
		plan.UseEnhancedOrgRole = types.StringValue("TRUE")
	}

	plan.Msg = types.StringValue(util.SafeDeref(apiResp.Msg))
	plan.ErrorCode = types.StringValue(util.SafeDeref(apiResp.ErrorCode))
}

func (r *WorkdayConnectionResource) ReadWorkdayConnection(ctx context.Context, connectionName string) (*openapi.GetConnectionDetailsResponse, error) {
	opCtx := errorsutil.CreateOperationContext(errorsutil.ConnectorTypeWorkday, "read", connectionName)

	// Create logging context (separate from API context)
	logCtx := opCtx.AddContextToLogger(ctx)

	opCtx.LogOperationStart(logCtx, "Starting Workday connection read operation")

	// Execute read operation with retry logic
	var apiResp *openapi.GetConnectionDetailsResponse

	err := r.provider.AuthenticatedAPICallWithRetry(ctx, "read_workday_connection", func(token string) error {
		connectionOps := r.connectionFactory.CreateConnectionOperations(r.client.APIBaseURL(), token)
		resp, httpResp, err := connectionOps.GetConnectionDetails(ctx, connectionName)
		if httpResp != nil && httpResp.StatusCode == 401 {
			return fmt.Errorf("401 unauthorized")
		}
		apiResp = resp
		return err
	})

	if err != nil {
		errorCode := workdayErrorCodes.ReadFailed()
		opCtx.LogOperationError(logCtx, "Failed to read Workday connection", errorCode, err)
		return nil, errorsutil.CreateStandardError(errorsutil.ConnectorTypeWorkday, errorCode, "read", connectionName, err)
	}

	if apiResp != nil && apiResp.WorkdayConnectionResponse != nil && apiResp.WorkdayConnectionResponse.Errorcode != nil && *apiResp.WorkdayConnectionResponse.Errorcode != 0 {
		apiErr := fmt.Errorf("API returned error code %d: %s", *apiResp.WorkdayConnectionResponse.Errorcode, errorsutil.SanitizeMessage(apiResp.WorkdayConnectionResponse.Msg))
		errorCode := workdayErrorCodes.APIError()
		opCtx.LogOperationError(ctx, "Workday connection read failed with API error", errorCode, apiErr,
			map[string]interface{}{
				"api_error_code": *apiResp.WorkdayConnectionResponse.Errorcode,
				"message":        errorsutil.SanitizeMessage(apiResp.WorkdayConnectionResponse.Msg),
			})
		return nil, errorsutil.CreateStandardError(errorsutil.ConnectorTypeWorkday, errorCode, "read", connectionName, apiErr)
	}

	opCtx.LogOperationEnd(logCtx, "Workday connection read completed successfully",
		map[string]interface{}{"connection_key": func() interface{} {
			if apiResp.WorkdayConnectionResponse != nil && apiResp.WorkdayConnectionResponse.Connectionkey != nil {
				return *apiResp.WorkdayConnectionResponse.Connectionkey
			}
			return "unknown"
		}()})

	return apiResp, nil
}

func (r *WorkdayConnectionResource) UpdateModelFromReadResponse(state *WorkdayConnectorResourceModel, apiResp *openapi.GetConnectionDetailsResponse) {
	state.ConnectionKey = types.Int64Value(int64(*apiResp.WorkdayConnectionResponse.Connectionkey))
	state.ID = types.StringValue(fmt.Sprintf("%d", *apiResp.WorkdayConnectionResponse.Connectionkey))

	state.ConnectionName = util.SafeStringDatasource(apiResp.WorkdayConnectionResponse.Connectionname)
	state.Description = util.SafeStringDatasource(apiResp.WorkdayConnectionResponse.Description)
	state.DefaultSavRoles = util.SafeStringDatasource(apiResp.WorkdayConnectionResponse.Defaultsavroles)
	state.EmailTemplate = util.SafeStringDatasource(apiResp.WorkdayConnectionResponse.Emailtemplate)
	state.UseOAuth = util.SafeStringDatasource(apiResp.WorkdayConnectionResponse.Connectionattributes.USE_OAUTH)
	state.UserImportMapping = util.SafeStringDatasource(apiResp.WorkdayConnectionResponse.Connectionattributes.USER_IMPORT_MAPPING)
	state.AccountsLastImportTime = util.SafeStringDatasource(apiResp.WorkdayConnectionResponse.Connectionattributes.ACCOUNTS_LAST_IMPORT_TIME)
	state.StatusKeyJSON = util.SafeStringDatasource(apiResp.WorkdayConnectionResponse.Connectionattributes.STATUS_KEY_JSON)
	state.RAASMappingJSON = util.SafeStringDatasource(apiResp.WorkdayConnectionResponse.Connectionattributes.RAAS_MAPPING_JSON)
	state.AccountImportPayload = util.SafeStringDatasource(apiResp.WorkdayConnectionResponse.Connectionattributes.ACCOUNT_IMPORT_PAYLOAD)
	state.UpdateAccountPayload = util.SafeStringDatasource(apiResp.WorkdayConnectionResponse.Connectionattributes.UPDATE_ACCOUNT_PAYLOAD)
	state.ClientID = util.SafeStringDatasource(apiResp.WorkdayConnectionResponse.Connectionattributes.CLIENT_ID)
	state.StatusThresholdConfig = util.SafeStringDatasource(apiResp.WorkdayConnectionResponse.Connectionattributes.STATUS_THRESHOLD_CONFIG)
	state.Username = util.SafeStringDatasource(apiResp.WorkdayConnectionResponse.Connectionattributes.USERNAME)
	state.AccessImportList = util.SafeStringDatasource(apiResp.WorkdayConnectionResponse.Connectionattributes.ACCESS_IMPORT_LIST)
	state.AccountImportMapping = util.SafeStringDatasource(apiResp.WorkdayConnectionResponse.Connectionattributes.ACCOUNT_IMPORT_MAPPING)
	state.OrgRoleImportPayload = util.SafeStringDatasource(apiResp.WorkdayConnectionResponse.Connectionattributes.ORGROLE_IMPORT_PAYLOAD)
	state.AssignOrgRolePayload = util.SafeStringDatasource(apiResp.WorkdayConnectionResponse.Connectionattributes.ASSIGN_ORGROLE_PAYLOAD)
	state.AccessImportMapping = util.SafeStringDatasource(apiResp.WorkdayConnectionResponse.Connectionattributes.ACCESS_IMPORT_MAPPING)
	state.APIVersion = util.SafeStringDatasource(apiResp.WorkdayConnectionResponse.Connectionattributes.API_VERSION)
	state.RemoveOrgRolePayload = util.SafeStringDatasource(apiResp.WorkdayConnectionResponse.Connectionattributes.REMOVE_ORGROLE_PAYLOAD)
	state.IncludeReferenceDesc = util.SafeStringDatasource(apiResp.WorkdayConnectionResponse.Connectionattributes.INCLUDE_REFERENCE_DESCRIPTORS)
	state.ModifyUserDataJSON = util.SafeStringDatasource(apiResp.WorkdayConnectionResponse.Connectionattributes.MODIFYUSERDATAJSON)
	state.UseX509AuthForSOAP = util.SafeStringDatasource(apiResp.WorkdayConnectionResponse.Connectionattributes.USEX509AUTHFORSOAP)
	state.ReportOwner = util.SafeStringDatasource(apiResp.WorkdayConnectionResponse.Connectionattributes.REPORT_OWNER)
	state.X509Key = util.SafeStringDatasource(apiResp.WorkdayConnectionResponse.Connectionattributes.X509KEY)
	state.CustomConfig = util.SafeStringDatasource(apiResp.WorkdayConnectionResponse.Connectionattributes.CUSTOM_CONFIG)
	state.UserAttributeJSON = util.SafeStringDatasource(apiResp.WorkdayConnectionResponse.Connectionattributes.USERATTRIBUTEJSON)
	state.X509Cert = util.SafeStringDatasource(apiResp.WorkdayConnectionResponse.Connectionattributes.X509CERT)
	state.UserImportPayload = util.SafeStringDatasource(apiResp.WorkdayConnectionResponse.Connectionattributes.USER_IMPORT_PAYLOAD)
	state.PAMConfig = util.SafeStringDatasource(apiResp.WorkdayConnectionResponse.Connectionattributes.PAM_CONFIG)
	state.AccessLastImportTime = util.SafeStringDatasource(apiResp.WorkdayConnectionResponse.Connectionattributes.ACCESS_LAST_IMPORT_TIME)
	state.UsersLastImportTime = util.SafeStringDatasource(apiResp.WorkdayConnectionResponse.Connectionattributes.USERS_LAST_IMPORT_TIME)
	state.UpdateUserPayload = util.SafeStringDatasource(apiResp.WorkdayConnectionResponse.Connectionattributes.UPDATE_USER_PAYLOAD)
	state.PageSize = util.SafeStringDatasource(apiResp.WorkdayConnectionResponse.Connectionattributes.PAGE_SIZE)
	state.TenantName = util.SafeStringDatasource(apiResp.WorkdayConnectionResponse.Connectionattributes.TENANT_NAME)
	state.UseEnhancedOrgRole = util.SafeStringDatasource(apiResp.WorkdayConnectionResponse.Connectionattributes.USE_ENHANCED_ORGROLE)
	state.CreateAccountPayload = util.SafeStringDatasource(apiResp.WorkdayConnectionResponse.Connectionattributes.CREATE_ACCOUNT_PAYLOAD)
	state.BaseURL = util.SafeStringDatasource(apiResp.WorkdayConnectionResponse.Connectionattributes.BASE_URL)
	apiMessage := util.SafeDeref(apiResp.WorkdayConnectionResponse.Msg)
	if apiMessage == "success" {
		state.Msg = types.StringValue("Connection Successful")
	} else {
		state.Msg = types.StringValue(apiMessage)
	}
	state.ErrorCode = util.Int32PtrToTFString(apiResp.WorkdayConnectionResponse.Errorcode)
}

func (r *WorkdayConnectionResource) UpdateWorkdayConnection(ctx context.Context, plan *WorkdayConnectorResourceModel, config *WorkdayConnectorResourceModel) (*openapi.CreateOrUpdateResponse, error) {
	connectionName := plan.ConnectionName.ValueString()
	opCtx := errorsutil.CreateOperationContext(errorsutil.ConnectorTypeWorkday, "update", connectionName)

	// Create logging context (separate from API context)
	logCtx := opCtx.AddContextToLogger(ctx)

	opCtx.LogOperationStart(logCtx, "Starting Workday connection update")

	// Build Workday connection update request
	tflog.Debug(logCtx, "Building Workday connection update request")

	workdayConn := r.BuildWorkdayConnector(plan, config)
	if plan.VaultConnection.ValueString() == "" {
		emptyStr := ""
		workdayConn.BaseConnector.VaultConnection = &emptyStr
		workdayConn.BaseConnector.VaultConfiguration = &emptyStr
		workdayConn.BaseConnector.Saveinvault = &emptyStr
	}

	updateReq := openapi.CreateOrUpdateRequest{
		WorkdayConnector: &workdayConn,
	}

	// Execute update operation with retry logic
	tflog.Debug(logCtx, "Executing update operation")
	var apiResp *openapi.CreateOrUpdateResponse

	err := r.provider.AuthenticatedAPICallWithRetry(ctx, "update_workday_connection", func(token string) error {
		connectionOps := r.connectionFactory.CreateConnectionOperations(r.client.APIBaseURL(), token)
		resp, httpResp, err := connectionOps.CreateOrUpdateConnection(ctx, updateReq)
		if httpResp != nil && httpResp.StatusCode == 401 {
			return fmt.Errorf("401 unauthorized")
		}
		apiResp = resp
		return err
	})

	if err != nil {
		errorCode := workdayErrorCodes.UpdateFailed()
		opCtx.LogOperationError(logCtx, "Failed to update Workday connection", errorCode, err)
		return nil, errorsutil.CreateStandardError(errorsutil.ConnectorTypeWorkday, errorCode, "update", connectionName, err)
	}

	if apiResp != nil && apiResp.ErrorCode != nil && *apiResp.ErrorCode != "0" {
		apiErr := fmt.Errorf("API returned error code %s: %s", *apiResp.ErrorCode, errorsutil.SanitizeMessage(apiResp.Msg))
		errorCode := workdayErrorCodes.APIError()
		opCtx.LogOperationError(logCtx, "Workday connection update failed with API error", errorCode, apiErr,
			map[string]interface{}{
				"api_error_code": *apiResp.ErrorCode,
				"message":        errorsutil.SanitizeMessage(apiResp.Msg),
			})
		return nil, errorsutil.CreateStandardError(errorsutil.ConnectorTypeWorkday, errorCode, "update", connectionName, apiErr)
	}

	opCtx.LogOperationEnd(logCtx, "Workday connection updated successfully",
		map[string]interface{}{"connection_key": func() interface{} {
			if apiResp.ConnectionKey != nil {
				return *apiResp.ConnectionKey
			}
			return "unknown"
		}()})

	return apiResp, nil
}

func (r *WorkdayConnectionResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan, config WorkdayConnectorResourceModel

	opCtx := errorsutil.CreateOperationContext(errorsutil.ConnectorTypeWorkday, "terraform_create", "")
	ctx = opCtx.AddContextToLogger(ctx)

	opCtx.LogOperationStart(ctx, "Starting Workday connection resource creation")

	// Extract plan from request
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		errorCode := workdayErrorCodes.PlanExtraction()
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
		errorCode := workdayErrorCodes.ConfigExtraction()
		opCtx.LogOperationError(ctx, "Failed to get config from request", errorCode,
			fmt.Errorf("config extraction failed"))
		resp.Diagnostics.AddError(
			errorsutil.GetErrorMessage(errorsutil.ErrConfigExtraction),
			fmt.Sprintf("[%s] Unable to extract Terraform configuration from request for connection '%s'", errorCode, connectionName),
		)
		return
	}

	// Use interface pattern instead of direct API client creation
	apiResp, err := r.CreateWorkdayConnection(ctx, &plan, &config)
	if err != nil {
		opCtx.LogOperationError(ctx, "Workday connection creation failed", "", err)
		resp.Diagnostics.AddError(
			"Workday Connection Creation Failed",
			err.Error(),
		)
		return
	}

	// Update model from create response
	r.UpdateModelFromCreateResponse(&plan, apiResp)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)

	opCtx.LogOperationEnd(ctx, "Workday connection resource created successfully",
		map[string]interface{}{"connection_key": plan.ConnectionKey.ValueInt64()})
}

func (r *WorkdayConnectionResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state WorkdayConnectorResourceModel

	opCtx := errorsutil.CreateOperationContext(errorsutil.ConnectorTypeWorkday, "terraform_read", "")
	ctx = opCtx.AddContextToLogger(ctx)

	opCtx.LogOperationStart(ctx, "Starting Workday connection resource read")

	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		errorCode := workdayErrorCodes.StateExtraction()
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
	apiResp, err := r.ReadWorkdayConnection(ctx, connectionName)
	if err != nil {
		opCtx.LogOperationError(ctx, "Workday connection read failed", "", err)
		resp.Diagnostics.AddError(
			"Workday Connection Read Failed",
			err.Error(),
		)
		return
	}

	// Update model from read response
	r.UpdateModelFromReadResponse(&state, apiResp)

	stateDiagnostics := resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(stateDiagnostics...)
	if resp.Diagnostics.HasError() {
		errorCode := workdayErrorCodes.StateUpdate()
		opCtx.LogOperationError(ctx, "Failed to set state", errorCode,
			fmt.Errorf("state update failed"))
		resp.Diagnostics.AddError(
			errorsutil.GetErrorMessage(errorsutil.ErrStateUpdate),
			fmt.Sprintf("[%s] Unable to update Terraform state for connection '%s'", errorCode, connectionName),
		)
		return
	}

	opCtx.LogOperationEnd(ctx, "Workday connection resource read completed successfully")
}
func (r *WorkdayConnectionResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state, config WorkdayConnectorResourceModel

	opCtx := errorsutil.CreateOperationContext(errorsutil.ConnectorTypeWorkday, "terraform_update", "")
	ctx = opCtx.AddContextToLogger(ctx)

	opCtx.LogOperationStart(ctx, "Starting Workday connection resource update")

	// Extract state from request
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		errorCode := workdayErrorCodes.StateExtraction()
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
		errorCode := workdayErrorCodes.PlanExtraction()
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
		errorCode := workdayErrorCodes.ConfigExtraction()
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
		errorCode := workdayErrorCodes.NameImmutable()
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
	_, err := r.UpdateWorkdayConnection(ctx, &plan, &config)
	if err != nil {
		opCtx.LogOperationError(ctx, "Workday connection update failed", "", err)
		resp.Diagnostics.AddError(
			"Workday Connection Update Failed",
			err.Error(),
		)
		return
	}

	// Read the updated connection to get the latest state
	getResp, err := r.ReadWorkdayConnection(ctx, connectionName)
	if err != nil {
		opCtx.LogOperationError(ctx, "Failed to read updated Workday connection", "", err)
		resp.Diagnostics.AddError(
			"Workday Connection Post-Update Read Failed",
			err.Error(),
		)
		return
	}

	// Update model from read response
	r.UpdateModelFromReadResponse(&plan, getResp)

	stateUpdateDiagnostics := resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(stateUpdateDiagnostics...)
	if resp.Diagnostics.HasError() {
		errorCode := workdayErrorCodes.StateUpdate()
		opCtx.LogOperationError(ctx, "Failed to update state after successful update", errorCode,
			fmt.Errorf("state update failed"))
		resp.Diagnostics.AddError(
			errorsutil.GetErrorMessage(errorsutil.ErrStateUpdate),
			fmt.Sprintf("[%s] Unable to update Terraform state after successful update for connection '%s'", errorCode, connectionName),
		)
		return
	}

	opCtx.LogOperationEnd(ctx, "Workday connection resource updated successfully",
		map[string]interface{}{"connection_key": plan.ConnectionKey.ValueInt64()})
}

func (r *WorkdayConnectionResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
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

func (r *WorkdayConnectionResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	// Importing a Workday connection resource requires the connection name
	connectionName := req.ID
	opCtx := errorsutil.CreateOperationContext(errorsutil.ConnectorTypeWorkday, "terraform_import", connectionName)
	ctx = opCtx.AddContextToLogger(ctx)

	opCtx.LogOperationStart(ctx, "Starting Workday connection resource import")

	// Retrieve import ID and save to connection_name attribute
	resource.ImportStatePassthroughID(ctx, path.Root("connection_name"), req, resp)

	opCtx.LogOperationEnd(ctx, "Workday connection resource import completed successfully",
		map[string]interface{}{"import_id": connectionName})
}
