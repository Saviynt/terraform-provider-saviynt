// Copyright (c) 2025 Saviynt Inc.
// SPDX-License-Identifier: MPL-2.0

// saviynt_workday_soap_connection_resource manages Workday SOAP connectors in the Saviynt Security Manager.
// The resource implements the full Terraform lifecycle:
//   - Create: provisions a new Workday SOAP connector using the supplied configuration.
//   - Read: fetches the current connector state from Saviynt to keep Terraform's state in sync.
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

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	openapi "github.com/saviynt/saviynt-api-go-client/connections"
)

// Ensure provider defined types fully satisfy framework interfaces
var _ resource.Resource = &WorkdaySOAPConnectionResource{}
var _ resource.ResourceWithImportState = &WorkdaySOAPConnectionResource{}

// Initialize error codes for Workday SOAP Connection operations
var workdaySOAPErrorCodes = errorsutil.NewConnectorErrorCodes(errorsutil.ConnectorTypeWorkdaySOAP)

type WorkdaySOAPConnectorResourceModel struct {
	BaseConnectorResourceModel
	ID                       types.String `tfsdk:"id"`
	Username                 types.String `tfsdk:"username"`
	Password                 types.String `tfsdk:"password"`
	PasswordWo               types.String `tfsdk:"password_wo"`
	SoapEndpoint             types.String `tfsdk:"soap_endpoint"`
	AccountsImportJson       types.String `tfsdk:"accounts_import_json"`
	ChangePassJson           types.String `tfsdk:"change_pass_json"`
	ChangePassJsonWo         types.String `tfsdk:"change_pass_json_wo"`
	CombinedCreateRequest    types.String `tfsdk:"combined_create_request"`
	ConnectionJson           types.String `tfsdk:"connection_json"`
	ConnectionJsonWo         types.String `tfsdk:"connection_json_wo"`
	CreateAccountJson        types.String `tfsdk:"create_account_json"`
	CustomConfig             types.String `tfsdk:"custom_config"`
	DataToImport             types.String `tfsdk:"data_to_import"`
	DateFormat               types.String `tfsdk:"date_format"`
	DeleteAccountJson        types.String `tfsdk:"delete_account_json"`
	DisableAccountJson       types.String `tfsdk:"disable_account_json"`
	EnableAccountJson        types.String `tfsdk:"enable_account_json"`
	GrantAccessJson          types.String `tfsdk:"grant_access_json"`
	HrImportJson             types.String `tfsdk:"hr_import_json"`
	ModifyUserDataJson       types.String `tfsdk:"modify_user_data_json"`
	PageSize                 types.String `tfsdk:"page_size"`
	PamConfig                types.String `tfsdk:"pam_config"`
	PasswordMaxLength        types.String `tfsdk:"password_max_length"`
	PasswordMinLength        types.String `tfsdk:"password_min_length"`
	PasswordNoofCapsAlpha    types.String `tfsdk:"password_noofcapsalpha"`
	PasswordNoofDigits       types.String `tfsdk:"password_noofdigits"`
	PasswordNoofSplChars     types.String `tfsdk:"password_noofsplchars"`
	PasswordType             types.String `tfsdk:"password_type"`
	ResponsePathPageResults  types.String `tfsdk:"responsepath_pageresults"`
	ResponsePathTotalResults types.String `tfsdk:"responsepath_totalresults"`
	ResponsePathUserList     types.String `tfsdk:"responsepath_userlist"`
	RevokeAccessJson         types.String `tfsdk:"revoke_access_json"`
	UpdateAccountJson        types.String `tfsdk:"update_account_json"`
	UpdateUserJson           types.String `tfsdk:"update_user_json"`
}

type WorkdaySOAPConnectionResource struct {
	client            client.SaviyntClientInterface
	token             string
	provider          client.SaviyntProviderInterface
	connectionFactory client.ConnectionFactoryInterface
}

func NewWorkdaySOAPConnectionResource() resource.Resource {
	return &WorkdaySOAPConnectionResource{
		connectionFactory: &client.DefaultConnectionFactory{},
	}
}

func NewWorkdaySOAPConnectionResourceWithFactory(factory client.ConnectionFactoryInterface) resource.Resource {
	return &WorkdaySOAPConnectionResource{
		connectionFactory: factory,
	}
}

func (r *WorkdaySOAPConnectionResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = "saviynt_workday_soap_connection_resource"
}

func WorkdaySOAPConnectorResourceSchema() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:    true,
			Description: "Resource ID.",
		},
		"username": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "Username for SOAP authentication.",
		},
		"password": schema.StringAttribute{
			Optional:    true,
			Sensitive:   true,
			Description: "Password for SOAP authentication. Either this or password_wo must be set.",
			Validators: []validator.String{
				stringvalidator.ConflictsWith(path.MatchRoot("password_wo")),
			},
		},
		"password_wo": schema.StringAttribute{
			Optional:    true,
			WriteOnly:   true,
			Description: "Write-only password for SOAP authentication. Either this or password must be set.",
			Validators: []validator.String{
				stringvalidator.ConflictsWith(path.MatchRoot("password")),
			},
		},
		"soap_endpoint": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "SOAP endpoint URL for Workday. Example: \"https://wd2-impl-services1.workday.com/ccx/service/tenant/Human_Resources/v35.0\"",
		},
		"accounts_import_json": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "JSON configuration for accounts import.",
		},
		"change_pass_json": schema.StringAttribute{
			Optional:    true,
			Sensitive:   true,
			Description: "JSON configuration for password changes. Either this or change_pass_json_wo must be set.",
			Validators: []validator.String{
				stringvalidator.ConflictsWith(path.MatchRoot("change_pass_json_wo")),
			},
		},
		"change_pass_json_wo": schema.StringAttribute{
			Optional:    true,
			WriteOnly:   true,
			Description: "Write-only JSON configuration for password changes. Either this or change_pass_json must be set.",
			Validators: []validator.String{
				stringvalidator.ConflictsWith(path.MatchRoot("change_pass_json")),
			},
		},
		"combined_create_request": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "Combined create request configuration.",
		},
		"connection_json": schema.StringAttribute{
			Optional:    true,
			Sensitive:   true,
			Description: "General connection JSON configuration. Either this or connection_json_wo must be set.",
			Validators: []validator.String{
				stringvalidator.ConflictsWith(path.MatchRoot("connection_json_wo")),
			},
		},
		"connection_json_wo": schema.StringAttribute{
			Optional:    true,
			WriteOnly:   true,
			Description: "Write-only general connection JSON configuration. Either this or connection_json must be set.",
			Validators: []validator.String{
				stringvalidator.ConflictsWith(path.MatchRoot("connection_json")),
			},
		},
		"create_account_json": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "JSON configuration for account creation.",
		},
		"custom_config": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "Custom configuration JSON.",
		},
		"data_to_import": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "Specification of data types to import. Example: \"Users,Accounts\"",
		},
		"date_format": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "Date format for data processing. Example: \"yyyy-MM-dd\"",
		},
		"delete_account_json": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "JSON configuration for account deletion.",
		},
		"disable_account_json": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "JSON configuration for account disabling.",
		},
		"enable_account_json": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "JSON configuration for account enabling.",
		},
		"grant_access_json": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "JSON configuration for granting access.",
		},
		"hr_import_json": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "JSON configuration for HR data import.",
		},
		"modify_user_data_json": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "JSON configuration for modifying user data.",
		},
		"page_size": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "Number of records per page. Example: \"100\"",
		},
		"pam_config": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "PAM configuration JSON.",
		},
		"password_max_length": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "Maximum password length. Example: \"20\"",
		},
		"password_min_length": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "Minimum password length. Example: \"8\"",
		},
		"password_noofcapsalpha": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "Number of capital letters required in password. Example: \"1\"",
		},
		"password_noofdigits": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "Number of digits required in password. Example: \"1\"",
		},
		"password_noofsplchars": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "Number of special characters required in password. Example: \"1\"",
		},
		"password_type": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "Type of password authentication. Example: \"BASIC\"",
		},
		"responsepath_pageresults": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "Response path for page results.",
		},
		"responsepath_totalresults": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "Response path for total results count.",
		},
		"responsepath_userlist": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "Response path for user list.",
		},
		"revoke_access_json": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "JSON configuration for revoking access.",
		},
		"update_account_json": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "JSON configuration for account updates.",
		},
		"update_user_json": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "JSON configuration for user updates.",
		},
	}
}

func (r *WorkdaySOAPConnectionResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: util.WorkdaySOAPConnDescription,
		Attributes:  connectionsutil.MergeResourceAttributes(BaseConnectorResourceSchema(), WorkdaySOAPConnectorResourceSchema()),
	}
}

func (r *WorkdaySOAPConnectionResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	opCtx := errorsutil.CreateOperationContext(errorsutil.ConnectorTypeWorkdaySOAP, "configure", "")
	ctx = opCtx.AddContextToLogger(ctx)

	opCtx.LogOperationStart(ctx, "Starting Workday SOAP connection resource configuration")

	// Check if provider data is available
	if req.ProviderData == nil {
		tflog.Debug(ctx, "ProviderData is nil, returning early")
		opCtx.LogOperationEnd(ctx, "Workday SOAP connection resource configuration completed - no provider data")
		return
	}

	// Cast provider data to your provider type
	prov, ok := req.ProviderData.(*SaviyntProvider)
	if !ok {
		errorCode := workdaySOAPErrorCodes.ProviderConfig()
		opCtx.LogOperationError(ctx, "Provider configuration failed", errorCode,
			fmt.Errorf("expected *SaviyntProvider, got different type"),
			map[string]interface{}{"expected_type": "*SaviyntProvider"})

		resp.Diagnostics.AddError(
			errorsutil.GetErrorMessage(errorsutil.ErrProviderConfig),
			fmt.Sprintf("[%s] Expected *SaviyntProvider, got different type", errorCode),
		)
		return
	}

	// Set the client and token from the provider state using interface wrapper
	r.client = &client.SaviyntClientWrapper{Client: prov.client}
	r.token = prov.accessToken
	r.provider = &client.SaviyntProviderWrapper{Provider: prov} // Store provider reference for retry logic

	opCtx.LogOperationEnd(ctx, "Workday SOAP connection resource configured successfully")
}

// SetClient sets the client for testing purposes
func (r *WorkdaySOAPConnectionResource) SetClient(client client.SaviyntClientInterface) {
	r.client = client
}

// SetToken sets the token for testing purposes
func (r *WorkdaySOAPConnectionResource) SetToken(token string) {
	r.token = token
}

// SetProvider sets the provider for testing purposes
func (r *WorkdaySOAPConnectionResource) SetProvider(provider client.SaviyntProviderInterface) {
	r.provider = provider
}

func (r *WorkdaySOAPConnectionResource) BuildWorkdaySOAPConnector(plan *WorkdaySOAPConnectorResourceModel, config *WorkdaySOAPConnectorResourceModel) openapi.WorkdaySOAPConnector {
	var password string
	if !config.Password.IsNull() && !config.Password.IsUnknown() {
		password = config.Password.ValueString()
	} else if !config.PasswordWo.IsNull() && !config.PasswordWo.IsUnknown() {
		password = config.PasswordWo.ValueString()
	}

	var changePassJson string
	if !config.ChangePassJson.IsNull() && !config.ChangePassJson.IsUnknown() {
		changePassJson = config.ChangePassJson.ValueString()
	} else if !config.ChangePassJsonWo.IsNull() && !config.ChangePassJsonWo.IsUnknown() {
		changePassJson = config.ChangePassJsonWo.ValueString()
	}

	var connectionJson string
	if !config.ConnectionJson.IsNull() && !config.ConnectionJson.IsUnknown() {
		connectionJson = config.ConnectionJson.ValueString()
	} else if !config.ConnectionJsonWo.IsNull() && !config.ConnectionJsonWo.IsUnknown() {
		connectionJson = config.ConnectionJsonWo.ValueString()
	}

	workdaySOAPConn := openapi.WorkdaySOAPConnector{
		BaseConnector: openapi.BaseConnector{
			// Required fields
			Connectiontype: "Workday-SOAP",
			ConnectionName: plan.ConnectionName.ValueString(),
			// Optional fields
			Description:     util.StringPointerOrEmpty(plan.Description),
			Defaultsavroles: util.StringPointerOrEmpty(plan.DefaultSavRoles),
			EmailTemplate:   util.StringPointerOrEmpty(plan.EmailTemplate),
		},
		// Authentication
		USERNAME:      util.StringPointerOrEmpty(plan.Username),
		PASSWORD:      &password,
		SOAP_ENDPOINT: util.StringPointerOrEmpty(plan.SoapEndpoint),

		// Data Import Configuration
		ACCOUNTS_IMPORT_JSON: util.StringPointerOrEmpty(plan.AccountsImportJson),
		HR_IMPORT_JSON:       util.StringPointerOrEmpty(plan.HrImportJson),
		DATA_TO_IMPORT:       util.StringPointerOrEmpty(plan.DataToImport),
		PAGE_SIZE:            util.StringPointerOrEmpty(plan.PageSize),
		DATEFORMAT:           util.StringPointerOrEmpty(plan.DateFormat),

		// Account Management
		CREATEACCOUNTJSON:  util.StringPointerOrEmpty(plan.CreateAccountJson),
		UPDATEACCOUNTJSON:  util.StringPointerOrEmpty(plan.UpdateAccountJson),
		DELETEACCOUNTJSON:  util.StringPointerOrEmpty(plan.DeleteAccountJson),
		ENABLEACCOUNTJSON:  util.StringPointerOrEmpty(plan.EnableAccountJson),
		DISABLEACCOUNTJSON: util.StringPointerOrEmpty(plan.DisableAccountJson),

		// User Management
		UPDATEUSERJSON:     util.StringPointerOrEmpty(plan.UpdateUserJson),
		MODIFYUSERDATAJSON: util.StringPointerOrEmpty(plan.ModifyUserDataJson),

		// Access Management
		GRANTACCESSJSON:  util.StringPointerOrEmpty(plan.GrantAccessJson),
		REVOKEACCESSJSON: util.StringPointerOrEmpty(plan.RevokeAccessJson),

		// Password Management
		CHANGEPASSJSON:         &changePassJson,
		PASSWORD_TYPE:          util.StringPointerOrEmpty(plan.PasswordType),
		PASSWORD_MIN_LENGTH:    util.StringPointerOrEmpty(plan.PasswordMinLength),
		PASSWORD_MAX_LENGTH:    util.StringPointerOrEmpty(plan.PasswordMaxLength),
		PASSWORD_NOOFCAPSALPHA: util.StringPointerOrEmpty(plan.PasswordNoofCapsAlpha),
		PASSWORD_NOOFDIGITS:    util.StringPointerOrEmpty(plan.PasswordNoofDigits),
		PASSWORD_NOOFSPLCHARS:  util.StringPointerOrEmpty(plan.PasswordNoofSplChars),

		// Response Path Configuration
		RESPONSEPATH_USERLIST:     util.StringPointerOrEmpty(plan.ResponsePathUserList),
		RESPONSEPATH_PAGERESULTS:  util.StringPointerOrEmpty(plan.ResponsePathPageResults),
		RESPONSEPATH_TOTALRESULTS: util.StringPointerOrEmpty(plan.ResponsePathTotalResults),

		// Advanced Configuration
		CONNECTIONJSON:        &connectionJson,
		CUSTOM_CONFIG:         util.StringPointerOrEmpty(plan.CustomConfig),
		PAM_CONFIG:            util.StringPointerOrEmpty(plan.PamConfig),
		COMBINEDCREATEREQUEST: util.StringPointerOrEmpty(plan.CombinedCreateRequest),
	}

	return workdaySOAPConn
}

func (r *WorkdaySOAPConnectionResource) UpdateModelFromCreateResponse(plan *WorkdaySOAPConnectorResourceModel, apiResp *openapi.CreateOrUpdateResponse) {
	plan.ID = types.StringValue(fmt.Sprintf("%d", *apiResp.ConnectionKey))
	plan.ConnectionKey = types.Int64Value(int64(*apiResp.ConnectionKey))
	plan.Description = util.SafeStringDatasource(plan.Description.ValueStringPointer())
	plan.DefaultSavRoles = util.SafeStringDatasource(plan.DefaultSavRoles.ValueStringPointer())
	plan.EmailTemplate = util.SafeStringDatasource(plan.EmailTemplate.ValueStringPointer())
	plan.Username = util.SafeStringDatasource(plan.Username.ValueStringPointer())
	plan.SoapEndpoint = util.SafeStringDatasource(plan.SoapEndpoint.ValueStringPointer())
	plan.AccountsImportJson = util.SafeStringDatasource(plan.AccountsImportJson.ValueStringPointer())
	plan.CombinedCreateRequest = util.SafeStringDatasource(plan.CombinedCreateRequest.ValueStringPointer())
	plan.CreateAccountJson = util.SafeStringDatasource(plan.CreateAccountJson.ValueStringPointer())
	plan.CustomConfig = util.SafeStringDatasource(plan.CustomConfig.ValueStringPointer())
	plan.DataToImport = util.SafeStringDatasource(plan.DataToImport.ValueStringPointer())
	plan.DateFormat = util.SafeStringDatasource(plan.DateFormat.ValueStringPointer())
	plan.DeleteAccountJson = util.SafeStringDatasource(plan.DeleteAccountJson.ValueStringPointer())
	plan.DisableAccountJson = util.SafeStringDatasource(plan.DisableAccountJson.ValueStringPointer())
	plan.EnableAccountJson = util.SafeStringDatasource(plan.EnableAccountJson.ValueStringPointer())
	plan.GrantAccessJson = util.SafeStringDatasource(plan.GrantAccessJson.ValueStringPointer())
	plan.HrImportJson = util.SafeStringDatasource(plan.HrImportJson.ValueStringPointer())
	plan.ModifyUserDataJson = util.SafeStringDatasource(plan.ModifyUserDataJson.ValueStringPointer())
	plan.PageSize = util.SafeStringDatasource(plan.PageSize.ValueStringPointer())
	plan.PamConfig = util.SafeStringDatasource(plan.PamConfig.ValueStringPointer())
	plan.PasswordMaxLength = util.SafeStringDatasource(plan.PasswordMaxLength.ValueStringPointer())
	plan.PasswordMinLength = util.SafeStringDatasource(plan.PasswordMinLength.ValueStringPointer())
	plan.PasswordNoofCapsAlpha = util.SafeStringDatasource(plan.PasswordNoofCapsAlpha.ValueStringPointer())
	plan.PasswordNoofDigits = util.SafeStringDatasource(plan.PasswordNoofDigits.ValueStringPointer())
	plan.PasswordNoofSplChars = util.SafeStringDatasource(plan.PasswordNoofSplChars.ValueStringPointer())
	plan.PasswordType = util.SafeStringDatasource(plan.PasswordType.ValueStringPointer())
	plan.ResponsePathPageResults = util.SafeStringDatasource(plan.ResponsePathPageResults.ValueStringPointer())
	plan.ResponsePathTotalResults = util.SafeStringDatasource(plan.ResponsePathTotalResults.ValueStringPointer())
	plan.ResponsePathUserList = util.SafeStringDatasource(plan.ResponsePathUserList.ValueStringPointer())
	plan.RevokeAccessJson = util.SafeStringDatasource(plan.RevokeAccessJson.ValueStringPointer())
	plan.UpdateAccountJson = util.SafeStringDatasource(plan.UpdateAccountJson.ValueStringPointer())
	plan.UpdateUserJson = util.SafeStringDatasource(plan.UpdateUserJson.ValueStringPointer())
	plan.Msg = types.StringValue(util.SafeDeref(apiResp.Msg))
	plan.ErrorCode = types.StringValue(util.SafeDeref(apiResp.ErrorCode))

}

func (r *WorkdaySOAPConnectionResource) CreateWorkdaySOAPConnection(ctx context.Context, plan *WorkdaySOAPConnectorResourceModel, config *WorkdaySOAPConnectorResourceModel) (*openapi.CreateOrUpdateResponse, error) {
	connectionName := plan.ConnectionName.ValueString()
	opCtx := errorsutil.CreateOperationContext(errorsutil.ConnectorTypeWorkdaySOAP, "create", connectionName)

	// Create logging context (separate from API context)
	logCtx := opCtx.AddContextToLogger(ctx)

	opCtx.LogOperationStart(logCtx, "Starting Workday SOAP connection creation")

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
		errorCode := workdaySOAPErrorCodes.ReadFailed()
		opCtx.LogOperationError(logCtx, "Failed to check existing connection", errorCode, err, nil)
		return nil, fmt.Errorf("[%s] Failed to check existing connection: %w", errorCode, err)
	}

	if existingResource != nil &&
		existingResource.WorkdaySOAPConnectionResponse != nil &&
		existingResource.WorkdaySOAPConnectionResponse.Errorcode != nil &&
		*existingResource.WorkdaySOAPConnectionResponse.Errorcode == 0 {

		errorCode := workdaySOAPErrorCodes.DuplicateName()
		opCtx.LogOperationError(ctx, "Connection name already exists. Please import or use a different name", errorCode,
			fmt.Errorf("duplicate connection name"))
		return nil, errorsutil.CreateStandardError(errorsutil.ConnectorTypeWorkdaySOAP, errorCode, "create", connectionName, nil)
	}

	// Build Workday SOAP connection create request
	tflog.Debug(ctx, "Building Workday SOAP connection create request")
	workdaySOAPConn := r.BuildWorkdaySOAPConnector(plan, config)
	createReq := openapi.CreateOrUpdateRequest{
		WorkdaySOAPConnector: &workdaySOAPConn,
	}

	var apiResp *openapi.CreateOrUpdateResponse

	// Execute create operation with retry logic
	tflog.Debug(ctx, "Executing create operation")
	err = r.provider.AuthenticatedAPICallWithRetry(ctx, "create_workday_soap_connection", func(token string) error {
		connectionOps := r.connectionFactory.CreateConnectionOperations(r.client.APIBaseURL(), token)
		resp, httpResp, err := connectionOps.CreateOrUpdateConnection(ctx, createReq)
		if httpResp != nil && httpResp.StatusCode == 401 {
			return fmt.Errorf("401 unauthorized")
		}
		apiResp = resp
		return err
	})

	if err != nil {
		errorCode := workdaySOAPErrorCodes.CreateFailed()
		opCtx.LogOperationError(ctx, "Failed to create connection", errorCode, err, nil)
		return nil, fmt.Errorf("[%s] Failed to create Workday SOAP connection: %w", errorCode, err)
	}

	if apiResp != nil && apiResp.ErrorCode != nil && *apiResp.ErrorCode != "0" {
		apiErr := fmt.Errorf("API returned error code %s: %s", *apiResp.ErrorCode, errorsutil.SanitizeMessage(apiResp.Msg))
		errorCode := workdaySOAPErrorCodes.APIError()
		opCtx.LogOperationError(ctx, "Workday SOAP connection creation failed with API error", errorCode, apiErr,
			map[string]interface{}{
				"api_error_code": *apiResp.ErrorCode,
				"message":        errorsutil.SanitizeMessage(apiResp.Msg),
			})
		return nil, errorsutil.CreateStandardError(errorsutil.ConnectorTypeWorkdaySOAP, errorCode, "create", connectionName, apiErr)
	}

	opCtx.LogOperationEnd(logCtx, "Workday SOAP connection created successfully",
		map[string]interface{}{"connection_key": func() interface{} {
			if apiResp.ConnectionKey != nil {
				return *apiResp.ConnectionKey
			}
			return "unknown"
		}()})

	return apiResp, nil
}

func (r *WorkdaySOAPConnectionResource) ReadWorkdaySOAPConnection(ctx context.Context, connectionName string) (*openapi.GetConnectionDetailsResponse, error) {
	opCtx := errorsutil.CreateOperationContext(errorsutil.ConnectorTypeWorkdaySOAP, "read", connectionName)

	// Create logging context (separate from API context)
	logCtx := opCtx.AddContextToLogger(ctx)

	opCtx.LogOperationStart(logCtx, "Starting Workday SOAP connection read operation")

	var apiResp *openapi.GetConnectionDetailsResponse

	// Execute read operation with retry logic
	err := r.provider.AuthenticatedAPICallWithRetry(ctx, "read_workday_soap_connection", func(token string) error {
		connectionOps := r.connectionFactory.CreateConnectionOperations(r.client.APIBaseURL(), token)
		resp, httpResp, err := connectionOps.GetConnectionDetails(ctx, connectionName)
		if httpResp != nil && httpResp.StatusCode == 401 {
			return fmt.Errorf("401 unauthorized")
		}
		apiResp = resp
		return err
	})

	if err != nil {
		errorCode := workdaySOAPErrorCodes.ReadFailed()
		opCtx.LogOperationError(logCtx, "Failed to read Workday SOAP connection", errorCode, err)
		return nil, errorsutil.CreateStandardError(errorsutil.ConnectorTypeWorkdaySOAP, errorCode, "read", connectionName, err)
	}

	if err := r.ValidateWorkdaySOAPConnectionResponse(apiResp); err != nil {
		errorCode := workdaySOAPErrorCodes.APIError()
		opCtx.LogOperationError(ctx, "Invalid connection type for Workday SOAP datasource", errorCode, err)
		return nil, fmt.Errorf("[%s] Unable to verify connection type for connection %q. The provider could not determine the type of this connection. Please ensure the connection name is correct and belongs to a supported connector type", errorCode, connectionName)
	}

	if apiResp != nil && apiResp.WorkdaySOAPConnectionResponse != nil && apiResp.WorkdaySOAPConnectionResponse.Errorcode != nil && *apiResp.WorkdaySOAPConnectionResponse.Errorcode != 0 {
		apiErr := fmt.Errorf("API returned error code %d: %s", *apiResp.WorkdaySOAPConnectionResponse.Errorcode, errorsutil.SanitizeMessage(apiResp.WorkdaySOAPConnectionResponse.Msg))
		errorCode := workdaySOAPErrorCodes.APIError()
		opCtx.LogOperationError(ctx, "Workday SOAP connection read failed with API error", errorCode, apiErr,
			map[string]interface{}{
				"api_error_code": *apiResp.WorkdaySOAPConnectionResponse.Errorcode,
				"message":        errorsutil.SanitizeMessage(apiResp.WorkdaySOAPConnectionResponse.Msg),
			})
		return nil, errorsutil.CreateStandardError(errorsutil.ConnectorTypeWorkdaySOAP, errorCode, "read", connectionName, apiErr)
	}

	opCtx.LogOperationEnd(logCtx, "Workday SOAP connection read completed successfully",
		map[string]interface{}{"connection_key": func() interface{} {
			if apiResp.WorkdaySOAPConnectionResponse != nil && apiResp.WorkdaySOAPConnectionResponse.Connectionkey != nil {
				return *apiResp.WorkdaySOAPConnectionResponse.Connectionkey
			}
			return "unknown"
		}()})

	return apiResp, nil
}

func (r *WorkdaySOAPConnectionResource) ValidateWorkdaySOAPConnectionResponse(apiResp *openapi.GetConnectionDetailsResponse) error {
	if apiResp != nil && apiResp.WorkdaySOAPConnectionResponse == nil {
		return fmt.Errorf("verify the connection type - Workday SOAP connection response is nil")
	}
	return nil
}

func (r *WorkdaySOAPConnectionResource) UpdateWorkdaySOAPConnection(ctx context.Context, plan *WorkdaySOAPConnectorResourceModel, config *WorkdaySOAPConnectorResourceModel) (*openapi.CreateOrUpdateResponse, error) {
	connectionName := plan.ConnectionName.ValueString()
	opCtx := errorsutil.CreateOperationContext(errorsutil.ConnectorTypeWorkdaySOAP, "update", connectionName)

	// Create logging context (separate from API context)
	logCtx := opCtx.AddContextToLogger(ctx)

	opCtx.LogOperationStart(logCtx, "Starting Workday SOAP connection update")

	// Build Workday SOAP connection update request
	tflog.Debug(logCtx, "Building Workday SOAP connection update request")

	workdaySOAPConn := r.BuildWorkdaySOAPConnector(plan, config)

	updateReq := openapi.CreateOrUpdateRequest{
		WorkdaySOAPConnector: &workdaySOAPConn,
	}

	var apiResp *openapi.CreateOrUpdateResponse

	// Execute update operation with retry logic
	tflog.Debug(logCtx, "Executing update operation")
	err := r.provider.AuthenticatedAPICallWithRetry(ctx, "update_workday_soap_connection", func(token string) error {
		connectionOps := r.connectionFactory.CreateConnectionOperations(r.client.APIBaseURL(), token)
		resp, httpResp, err := connectionOps.CreateOrUpdateConnection(ctx, updateReq)
		if httpResp != nil && httpResp.StatusCode == 401 {
			return fmt.Errorf("401 unauthorized")
		}
		apiResp = resp
		return err
	})

	if err != nil {
		errorCode := workdaySOAPErrorCodes.UpdateFailed()
		opCtx.LogOperationError(logCtx, "Failed to update Workday SOAP connection", errorCode, err)
		return nil, errorsutil.CreateStandardError(errorsutil.ConnectorTypeWorkdaySOAP, errorCode, "update", connectionName, err)
	}

	if apiResp != nil && apiResp.ErrorCode != nil && *apiResp.ErrorCode != "0" {
		apiErr := fmt.Errorf("API returned error code %s: %s", *apiResp.ErrorCode, errorsutil.SanitizeMessage(apiResp.Msg))
		errorCode := workdaySOAPErrorCodes.APIError()
		opCtx.LogOperationError(logCtx, "Workday SOAP connection update failed with API error", errorCode, apiErr,
			map[string]interface{}{
				"api_error_code": *apiResp.ErrorCode,
				"message":        errorsutil.SanitizeMessage(apiResp.Msg),
			})
		return nil, errorsutil.CreateStandardError(errorsutil.ConnectorTypeWorkdaySOAP, errorCode, "update", connectionName, apiErr)
	}

	opCtx.LogOperationEnd(logCtx, "Workday SOAP connection updated successfully",
		map[string]interface{}{"connection_key": func() interface{} {
			if apiResp.ConnectionKey != nil {
				return *apiResp.ConnectionKey
			}
			return "unknown"
		}()})

	return apiResp, nil
}

func (r *WorkdaySOAPConnectionResource) UpdateModelFromReadResponse(state *WorkdaySOAPConnectorResourceModel, apiResp *openapi.GetConnectionDetailsResponse) {
	state.ConnectionKey = types.Int64Value(int64(*apiResp.WorkdaySOAPConnectionResponse.Connectionkey))
	state.ID = types.StringValue(fmt.Sprintf("%d", *apiResp.WorkdaySOAPConnectionResponse.Connectionkey))
	state.ConnectionName = util.SafeStringDatasource(apiResp.WorkdaySOAPConnectionResponse.Connectionname)
	state.Description = util.SafeStringDatasource(apiResp.WorkdaySOAPConnectionResponse.Description)
	state.DefaultSavRoles = util.SafeStringDatasource(apiResp.WorkdaySOAPConnectionResponse.Defaultsavroles)
	state.EmailTemplate = util.SafeStringDatasource(apiResp.WorkdaySOAPConnectionResponse.Emailtemplate)

	if apiResp.WorkdaySOAPConnectionResponse.Connectionattributes != nil {
		attrs := apiResp.WorkdaySOAPConnectionResponse.Connectionattributes
		state.Username = util.SafeStringDatasource(attrs.USERNAME)
		state.SoapEndpoint = util.SafeStringDatasource(attrs.SOAP_ENDPOINT)
		state.AccountsImportJson = util.SafeStringDatasource(attrs.ACCOUNTS_IMPORT_JSON)
		state.CombinedCreateRequest = util.SafeStringDatasource(attrs.COMBINEDCREATEREQUEST)
		state.CreateAccountJson = util.SafeStringDatasource(attrs.CREATEACCOUNTJSON)
		state.CustomConfig = util.SafeStringDatasource(attrs.CUSTOM_CONFIG)
		state.DataToImport = util.SafeStringDatasource(attrs.DATA_TO_IMPORT)
		state.DateFormat = util.SafeStringDatasource(attrs.DATEFORMAT)
		state.DeleteAccountJson = util.SafeStringDatasource(attrs.DELETEACCOUNTJSON)
		state.DisableAccountJson = util.SafeStringDatasource(attrs.DISABLEACCOUNTJSON)
		state.EnableAccountJson = util.SafeStringDatasource(attrs.ENABLEACCOUNTJSON)
		state.GrantAccessJson = util.SafeStringDatasource(attrs.GRANTACCESSJSON)
		state.HrImportJson = util.SafeStringDatasource(attrs.HR_IMPORT_JSON)
		state.ModifyUserDataJson = util.SafeStringDatasource(attrs.MODIFYUSERDATAJSON)
		state.PageSize = util.SafeStringDatasource(attrs.PAGE_SIZE)
		state.PamConfig = util.SafeStringDatasource(attrs.PAM_CONFIG)
		state.PasswordMaxLength = util.SafeStringDatasource(attrs.PASSWORD_MAX_LENGTH)
		state.PasswordMinLength = util.SafeStringDatasource(attrs.PASSWORD_MIN_LENGTH)
		state.PasswordNoofCapsAlpha = util.SafeStringDatasource(attrs.PASSWORD_NOOFCAPSALPHA)
		state.PasswordNoofDigits = util.SafeStringDatasource(attrs.PASSWORD_NOOFDIGITS)
		state.PasswordNoofSplChars = util.SafeStringDatasource(attrs.PASSWORD_NOOFSPLCHARS)
		state.PasswordType = util.SafeStringDatasource(attrs.PASSWORD_TYPE)
		state.ResponsePathPageResults = util.SafeStringDatasource(attrs.RESPONSEPATH_PAGERESULTS)
		state.ResponsePathTotalResults = util.SafeStringDatasource(attrs.RESPONSEPATH_TOTALRESULTS)
		state.ResponsePathUserList = util.SafeStringDatasource(attrs.RESPONSEPATH_USERLIST)
		state.RevokeAccessJson = util.SafeStringDatasource(attrs.REVOKEACCESSJSON)
		state.UpdateAccountJson = util.SafeStringDatasource(attrs.UPDATEACCOUNTJSON)
		state.UpdateUserJson = util.SafeStringDatasource(attrs.UPDATEUSERJSON)
	}
}

func (r *WorkdaySOAPConnectionResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan, config WorkdaySOAPConnectorResourceModel

	opCtx := errorsutil.CreateOperationContext(errorsutil.ConnectorTypeWorkdaySOAP, "terraform_create", "")
	ctx = opCtx.AddContextToLogger(ctx)

	opCtx.LogOperationStart(ctx, "Starting Workday SOAP connection resource creation")

	// Extract plan from request
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		errorCode := workdaySOAPErrorCodes.PlanExtraction()
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
		errorCode := workdaySOAPErrorCodes.ConfigExtraction()
		opCtx.LogOperationError(ctx, "Failed to get config from request", errorCode,
			fmt.Errorf("config extraction failed"))
		resp.Diagnostics.AddError(
			errorsutil.GetErrorMessage(errorsutil.ErrConfigExtraction),
			fmt.Sprintf("[%s] Unable to extract Terraform configuration from request for connection '%s'", errorCode, connectionName),
		)
		return
	}

	// Use interface pattern instead of direct API client creation
	apiResp, err := r.CreateWorkdaySOAPConnection(ctx, &plan, &config)
	if err != nil {
		opCtx.LogOperationError(ctx, "Workday SOAP connection creation failed", "", err)
		resp.Diagnostics.AddError(
			"Workday SOAP Connection Creation Failed",
			err.Error(),
		)
		return
	}

	// Update model from create response
	r.UpdateModelFromCreateResponse(&plan, apiResp)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)

	opCtx.LogOperationEnd(ctx, "Workday SOAP connection resource created successfully",
		map[string]interface{}{"connection_key": plan.ConnectionKey.ValueInt64()})
}

func (r *WorkdaySOAPConnectionResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state WorkdaySOAPConnectorResourceModel

	opCtx := errorsutil.CreateOperationContext(errorsutil.ConnectorTypeWorkdaySOAP, "terraform_read", "")
	ctx = opCtx.AddContextToLogger(ctx)

	opCtx.LogOperationStart(ctx, "Starting Workday SOAP connection resource read")

	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		errorCode := workdaySOAPErrorCodes.StateExtraction()
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
	apiResp, err := r.ReadWorkdaySOAPConnection(ctx, connectionName)
	if err != nil {
		opCtx.LogOperationError(ctx, "Workday SOAP connection read failed", "", err)
		resp.Diagnostics.AddError(
			"Workday SOAP Connection Read Failed",
			err.Error(),
		)
		return
	}

	// Update model from read response
	r.UpdateModelFromReadResponse(&state, apiResp)
	apiMessage := util.SafeDeref(apiResp.WorkdaySOAPConnectionResponse.Msg)
	if apiMessage == "success" {
		state.Msg = types.StringValue("Connection Read Successful")
	} else {
		state.Msg = types.StringValue(apiMessage)
	}
	state.ErrorCode = util.Int32PtrToTFString(apiResp.WorkdaySOAPConnectionResponse.Errorcode)

	stateDiagnostics := resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(stateDiagnostics...)
	if resp.Diagnostics.HasError() {
		errorCode := workdaySOAPErrorCodes.StateUpdate()
		opCtx.LogOperationError(ctx, "Failed to set state", errorCode,
			fmt.Errorf("state update failed"))
		resp.Diagnostics.AddError(
			errorsutil.GetErrorMessage(errorsutil.ErrStateUpdate),
			fmt.Sprintf("[%s] Unable to update Terraform state for connection '%s'", errorCode, connectionName),
		)
		return
	}

	opCtx.LogOperationEnd(ctx, "Workday SOAP connection resource read completed successfully")
}

func (r *WorkdaySOAPConnectionResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state, config WorkdaySOAPConnectorResourceModel

	opCtx := errorsutil.CreateOperationContext(errorsutil.ConnectorTypeWorkdaySOAP, "terraform_update", "")
	ctx = opCtx.AddContextToLogger(ctx)

	opCtx.LogOperationStart(ctx, "Starting Workday SOAP connection resource update")

	// Extract state from request
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		errorCode := workdaySOAPErrorCodes.StateExtraction()
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
		errorCode := workdaySOAPErrorCodes.PlanExtraction()
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
		errorCode := workdaySOAPErrorCodes.ConfigExtraction()
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
		errorCode := workdaySOAPErrorCodes.NameImmutable()
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
	updateResp, err := r.UpdateWorkdaySOAPConnection(ctx, &plan, &config)
	if err != nil {
		opCtx.LogOperationError(ctx, "Workday SOAP connection update failed", "", err)
		resp.Diagnostics.AddError(
			"Workday SOAP Connection Update Failed",
			err.Error(),
		)
		return
	}

	// Read the updated connection to get the latest state
	getResp, err := r.ReadWorkdaySOAPConnection(ctx, connectionName)
	if err != nil {
		opCtx.LogOperationError(ctx, "Failed to read updated Workday SOAP connection", "", err)
		resp.Diagnostics.AddError(
			"Workday SOAP Connection Post-Update Read Failed",
			err.Error(),
		)
		return
	}

	// Update model from read response
	r.UpdateModelFromReadResponse(&plan, getResp)

	apiMessage := util.SafeDeref(updateResp.Msg)
	plan.Msg = types.StringValue(apiMessage)
	plan.ErrorCode = types.StringValue(*updateResp.ErrorCode)

	stateUpdateDiagnostics := resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(stateUpdateDiagnostics...)
	if resp.Diagnostics.HasError() {
		errorCode := workdaySOAPErrorCodes.StateUpdate()
		opCtx.LogOperationError(ctx, "Failed to update state after successful update", errorCode,
			fmt.Errorf("state update failed"))
		resp.Diagnostics.AddError(
			errorsutil.GetErrorMessage(errorsutil.ErrStateUpdate),
			fmt.Sprintf("[%s] Unable to update Terraform state after successful update for connection '%s'", errorCode, connectionName),
		)
		return
	}

	opCtx.LogOperationEnd(ctx, "Workday SOAP connection resource updated successfully",
		map[string]interface{}{"connection_key": plan.ConnectionKey.ValueInt64()})
}

func (r *WorkdaySOAPConnectionResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
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

func (r *WorkdaySOAPConnectionResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	// Retrieve import ID and save to connection_name attribute
	resource.ImportStatePassthroughID(ctx, path.Root("connection_name"), req, resp)
}
