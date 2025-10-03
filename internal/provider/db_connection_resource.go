// Copyright (c) 2025 Saviynt Inc.
// SPDX-License-Identifier: MPL-2.0

// saviynt_db_connection_resource manages DB connectors in the Saviynt Security Manager.
// The resource implements the full Terraform lifecycle:
//   - Create: provisions a new DB connector using the supplied configuration.
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

	openapi "github.com/saviynt/saviynt-api-go-client/connections"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &DBConnectionResource{}
var _ resource.ResourceWithImportState = &DBConnectionResource{}

// Initialize error codes for DB Connection operations
var dbErrorCodes = errorsutil.NewConnectorErrorCodes(errorsutil.ConnectorTypeDB)

type DBConnectorResourceModel struct {
	BaseConnectorResourceModel
	ID                     types.String `tfsdk:"id"`
	URL                    types.String `tfsdk:"url"`
	Username               types.String `tfsdk:"username"`
	Password               types.String `tfsdk:"password"`
	DriverName             types.String `tfsdk:"driver_name"`
	ConnectionProperties   types.String `tfsdk:"connection_properties"`
	PasswordMinLength      types.String `tfsdk:"password_min_length"`
	PasswordMaxLength      types.String `tfsdk:"password_max_length"`
	PasswordNoOfCapsAlpha  types.String `tfsdk:"password_no_of_caps_alpha"`
	PasswordNoOfDigits     types.String `tfsdk:"password_no_of_digits"`
	PasswordNoOfSplChars   types.String `tfsdk:"password_no_of_spl_chars"`
	CreateAccountJson      types.String `tfsdk:"create_account_json"`
	UpdateAccountJson      types.String `tfsdk:"update_account_json"`
	GrantAccessJson        types.String `tfsdk:"grant_access_json"`
	RevokeAccessJson       types.String `tfsdk:"revoke_access_json"`
	ChangePassJson         types.String `tfsdk:"change_pass_json"`
	DeleteAccountJson      types.String `tfsdk:"delete_account_json"`
	EnableAccountJson      types.String `tfsdk:"enable_account_json"`
	DisableAccountJson     types.String `tfsdk:"disable_account_json"`
	AccountExistsJson      types.String `tfsdk:"account_exists_json"`
	UpdateUserJson         types.String `tfsdk:"update_user_json"`
	AccountsImport         types.String `tfsdk:"accounts_import"`
	EntitlementValueImport types.String `tfsdk:"entitlement_value_import"`
	RoleOwnerImport        types.String `tfsdk:"role_owner_import"`
	RolesImport            types.String `tfsdk:"roles_import"`
	SystemImport           types.String `tfsdk:"system_import"`
	UserImport             types.String `tfsdk:"user_import"`
	ModifyUserDataJson     types.String `tfsdk:"modify_user_data_json"`
	StatusThresholdConfig  types.String `tfsdk:"status_threshold_config"`
	MaxPaginationSize      types.String `tfsdk:"max_pagination_size"`
	CliCommandJson         types.String `tfsdk:"cli_command_json"`
	//TER-176
	CreateEntitlementJson types.String `tfsdk:"create_entitlement_json"`
	DeleteEntitlementJson types.String `tfsdk:"delete_entitlement_json"`
	EntitlementExistJson  types.String `tfsdk:"entitlement_exist_json"`
	UpdateEntitlementJson types.String `tfsdk:"update_entitlement_json"`
}

type DBConnectionResource struct {
	client            client.SaviyntClientInterface
	token             string
	provider          client.SaviyntProviderInterface
	connectionFactory client.ConnectionFactoryInterface
}

func NewDBConnectionResource() resource.Resource {
	return &DBConnectionResource{
		connectionFactory: &client.DefaultConnectionFactory{},
	}
}

func NewDBConnectionResourceWithFactory(factory client.ConnectionFactoryInterface) resource.Resource {
	return &DBConnectionResource{
		connectionFactory: factory,
	}
}

func (r *DBConnectionResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = "saviynt_db_connection_resource"
}

func DBConnectorResourceSchema() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:    true,
			Description: "Resource ID.",
		},
		"url": schema.StringAttribute{
			Required:    true,
			Description: "Host Name for connection",
		},
		"username": schema.StringAttribute{
			Required:    true,
			WriteOnly:   true,
			Description: "Username for connection",
		},
		"password": schema.StringAttribute{
			Required:    true,
			WriteOnly:   true,
			Description: "Password for connection",
		},
		"driver_name": schema.StringAttribute{
			Required:    true,
			Description: "Driver name for the connection",
		},
		"connection_properties": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "Properties that need to be added when connecting to the database",
		},
		"password_min_length": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "Specify the minimum length for the random password",
		},
		"password_max_length": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "Specify the maximum length for the random password",
		},
		"password_no_of_caps_alpha": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "Specify the number of uppercase alphabets required for the random password",
		},
		"password_no_of_digits": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "Specify the number of digits required for the random password",
		},
		"password_no_of_spl_chars": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "Specify the number of special characters required for the random password",
		},
		"create_account_json": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "JSON to specify the queries/stored procedures used to create a new account (e.g., randomPassword, task, user, accountName, role, endpoint, etc.)",
		},
		"update_account_json": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "JSON to specify the queries/stored procedures used to update an existing account",
		},
		"grant_access_json": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "JSON to specify the queries/stored procedures used to provide access",
		},
		"revoke_access_json": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "JSON to specify the queries/stored procedures used to revoke access",
		},
		"change_pass_json": schema.StringAttribute{
			Optional:    true,
			WriteOnly:   true,
			Description: "JSON to specify the queries/stored procedures used to change a password",
		},
		"delete_account_json": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "JSON to specify the queries/stored procedures used to delete an account",
		},
		"enable_account_json": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "JSON to specify the queries/stored procedures used to enable an account",
		},
		"disable_account_json": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "JSON to specify the queries/stored procedures used to disable an account",
		},
		"account_exists_json": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "JSON to specify the query used to check whether an account exists",
		},
		"update_user_json": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "JSON to specify the queries/stored procedures used to update user information",
		},
		"accounts_import": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "Accounts Import XML file content",
		},
		"entitlement_value_import": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "Entitlement Value Import XML file content",
		},
		"role_owner_import": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "Role Owner Import XML file content",
		},
		"roles_import": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "Roles Import XML file content",
		},
		"system_import": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "System Import XML file content",
		},
		"user_import": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "User Import XML file content",
		},
		"modify_user_data_json": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "Property for MODIFYUSERDATAJSON",
		},
		"status_threshold_config": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "Configuration for status and threshold (e.g., statusColumn, activeStatus, accountThresholdValue, etc.)",
		},
		"max_pagination_size": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "Defines the maximum number of records to be processed per page",
		},
		"cli_command_json": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "JSON to specify commands executable on the target server",
		},
		//TER-176
		"create_entitlement_json": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: " JSON to specify the Queries/stored procedures which will be used to Create the New Entitlements. Objects Exposed - (entitlementMgmtObj, task, user, endpoint and all the objects defined in Dynamic Attributes).",
		},
		"delete_entitlement_json": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: " JSON to specify the Queries/stored procedures which will be used to Delete the Entitlements. Objects Exposed - (entitlementMgmtObj, task, user, endpoint and all the objects defined in Dynamic Attributes).",
		},
		"entitlement_exist_json": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "JSON to specify the Query which will be used to check whether an entitlement exists. Objects Exposed - (entitlementMgmtObj, task, user, endpoint and all the objects defined in Dynamic Attributes).",
		},
		"update_entitlement_json": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: " JSON to specify the Queries/stored procedures which will be used to Update the Entitlements. Objects Exposed - (entitlementMgmtObj, task, user, endpoint and all the objects defined in Dynamic Attributes).",
		},
	}
}

func (r *DBConnectionResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: util.DBConnDescription,
		Attributes:  connectionsutil.MergeResourceAttributes(BaseConnectorResourceSchema(), DBConnectorResourceSchema()),
	}
}

func (r *DBConnectionResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	opCtx := errorsutil.CreateOperationContext(errorsutil.ConnectorTypeDB, "configure", "")
	ctx = opCtx.AddContextToLogger(ctx)

	opCtx.LogOperationStart(ctx, "Starting DB connection resource configuration")

	// Check if provider data is available.
	if req.ProviderData == nil {
		tflog.Debug(ctx, "ProviderData is nil, returning early")
		opCtx.LogOperationEnd(ctx, "DB connection resource configuration completed - no provider data")
		return
	}

	// Cast provider data to your provider type.
	prov, ok := req.ProviderData.(*SaviyntProvider)
	if !ok {
		errorCode := dbErrorCodes.ProviderConfig()
		opCtx.LogOperationError(ctx, "Provider configuration failed", errorCode,
			fmt.Errorf("expected *SaviyntProvider, got different type"),
			map[string]interface{}{"expected_type": "*SaviyntProvider"})

		resp.Diagnostics.AddError(
			errorsutil.GetErrorMessage(errorsutil.ErrProviderConfig),
			fmt.Sprintf("[%s] Expected *SaviyntProvider, got different type", errorCode),
		)
		return
	}

	// Set the client and token from the provider state.
	r.client = &client.SaviyntClientWrapper{Client: prov.client}
	r.token = prov.accessToken
	r.provider = &client.SaviyntProviderWrapper{Provider: prov} // Store provider reference for retry logic

	opCtx.LogOperationEnd(ctx, "DB connection resource configuration completed successfully")
}

// SetClient sets the client for testing purposes
func (r *DBConnectionResource) SetClient(client client.SaviyntClientInterface) {
	r.client = client
}

// SetToken sets the token for testing purposes
func (r *DBConnectionResource) SetToken(token string) {
	r.token = token
}

// SetProvider sets the provider for testing purposes
func (r *DBConnectionResource) SetProvider(provider client.SaviyntProviderInterface) {
	r.provider = provider
}
func (r *DBConnectionResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan, config DBConnectorResourceModel

	opCtx := errorsutil.CreateOperationContext(errorsutil.ConnectorTypeDB, "terraform_create", "")
	ctx = opCtx.AddContextToLogger(ctx)

	opCtx.LogOperationStart(ctx, "Starting DB connection resource creation")

	// Extract plan from request
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		errorCode := dbErrorCodes.PlanExtraction()
		opCtx.LogOperationError(ctx, "Failed to get plan from request", errorCode,
			fmt.Errorf("plan extraction failed"))
		resp.Diagnostics.AddError(
			errorsutil.GetErrorMessage(errorsutil.ErrPlanExtraction),
			fmt.Sprintf("[%s] Unable to extract Terraform plan from request", errorCode),
		)
		return
	}

	connectionName := plan.ConnectionName.ValueString()
	opCtx.ConnectionName = connectionName

	//Extract config from request
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		errorCode := dbErrorCodes.ConfigExtraction()
		opCtx.LogOperationError(ctx, "Failed to get config from request", errorCode,
			fmt.Errorf("config extraction failed"))
		resp.Diagnostics.AddError(
			errorsutil.GetErrorMessage(errorsutil.ErrConfigExtraction),
			fmt.Sprintf("[%s] Unable to extract Terraform configuration from request for connection '%s'", errorCode, connectionName),
		)
		return
	}

	// Use interface pattern instead of direct API client creation
	apiResp, err := r.CreateDBConnection(ctx, &plan, &config)
	if err != nil {
		opCtx.LogOperationError(ctx, "DB connection creation failed", "", err)
		resp.Diagnostics.AddError(
			"DB Connection Creation Failed",
			err.Error(),
		)
		return
	}

	r.UpdateModelFromCreateResponse(&plan, apiResp)

	diags := resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)

	opCtx.LogOperationEnd(ctx, "DB connection resource creation completed successfully",
		map[string]interface{}{"connection_name": connectionName})
}

// CreateDBConnection creates a new DB connection
func (r *DBConnectionResource) CreateDBConnection(ctx context.Context, plan *DBConnectorResourceModel, config *DBConnectorResourceModel) (*openapi.CreateOrUpdateResponse, error) {
	connectionName := plan.ConnectionName.ValueString()
	opCtx := errorsutil.CreateOperationContext(errorsutil.ConnectorTypeDB, "create", connectionName)

	// Create logging context (separate from API context)
	logCtx := opCtx.AddContextToLogger(ctx)

	opCtx.LogOperationStart(logCtx, "Starting DB connection creation",
		map[string]interface{}{"connection_name": connectionName})

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
		errorCode := dbErrorCodes.ReadFailed()
		opCtx.LogOperationError(logCtx, "Failed to check existing connection", errorCode, err)
		return nil, errorsutil.CreateStandardError(errorsutil.ConnectorTypeDB, errorCode, "create", connectionName, err)
	}

	if existingResource != nil &&
		existingResource.DBConnectionResponse != nil &&
		existingResource.DBConnectionResponse.Errorcode != nil &&
		*existingResource.DBConnectionResponse.Errorcode == 0 {

		errorCode := dbErrorCodes.DuplicateName()
		opCtx.LogOperationError(ctx, "Connection name already exists.Please import or use a different name", errorCode,
			fmt.Errorf("duplicate connection name"))
		return nil, errorsutil.CreateStandardError(errorsutil.ConnectorTypeDB, errorCode, "create", connectionName, nil)
	}

	// Build DB connector request
	dbConn := r.BuildDBConnector(plan, config)
	dbConnRequest := openapi.CreateOrUpdateRequest{
		DBConnector: &dbConn,
	}

	opCtx.LogOperationStart(logCtx, "Executing DB connection create API call")

	// Execute create operation with retry logic
	var apiResp *openapi.CreateOrUpdateResponse

	err = r.provider.AuthenticatedAPICallWithRetry(ctx, "create_db_connection", func(token string) error {
		connectionOps := r.connectionFactory.CreateConnectionOperations(r.client.APIBaseURL(), token)
		resp, httpResp, err := connectionOps.CreateOrUpdateConnection(ctx, dbConnRequest)
		if httpResp != nil && httpResp.StatusCode == 401 {
			return fmt.Errorf("401 unauthorized")
		}
		apiResp = resp
		return err
	})

	if err != nil {
		errorCode := dbErrorCodes.CreateFailed()
		opCtx.LogOperationError(logCtx, "Failed to create DB connection", errorCode, err)
		return nil, errorsutil.CreateStandardError(errorsutil.ConnectorTypeDB, errorCode, "create", connectionName, err)
	}

	if apiResp != nil && apiResp.ErrorCode != nil && *apiResp.ErrorCode != "0" {
		apiErr := fmt.Errorf("API returned error code %s: %s", *apiResp.ErrorCode, errorsutil.SanitizeMessage(apiResp.Msg))
		errorCode := dbErrorCodes.APIError()
		opCtx.LogOperationError(ctx, "DB connection creation failed with API error", errorCode, apiErr,
			map[string]interface{}{
				"api_error_code": *apiResp.ErrorCode,
				"message":        errorsutil.SanitizeMessage(apiResp.Msg),
			})
		return nil, errorsutil.CreateStandardError(errorsutil.ConnectorTypeDB, errorCode, "create", connectionName, apiErr)
	}

	opCtx.LogOperationEnd(logCtx, "DB connection created successfully",
		map[string]interface{}{
			"connection_name": connectionName,
			"connection_key":  apiResp.ConnectionKey,
		})

	return apiResp, nil
}

// BuildDBConnector builds the DB connector object
func (r *DBConnectionResource) BuildDBConnector(plan *DBConnectorResourceModel, config *DBConnectorResourceModel) openapi.DBConnector {
	connector := openapi.DBConnector{
		BaseConnector: openapi.BaseConnector{
			Connectiontype:  "DB",
			ConnectionName:  plan.ConnectionName.ValueString(),
			Description:     util.StringPointerOrEmpty(plan.Description),
			Defaultsavroles: util.StringPointerOrEmpty(plan.DefaultSavRoles),
			EmailTemplate:   util.StringPointerOrEmpty(plan.EmailTemplate),
		},
		// Required fields
		URL:        plan.URL.ValueString(),
		USERNAME:   config.Username.ValueString(),
		PASSWORD:   config.Password.ValueString(),
		DRIVERNAME: plan.DriverName.ValueString(),

		// Optional configuration fields
		CONNECTIONPROPERTIES:   util.StringPointerOrEmpty(plan.ConnectionProperties),
		PASSWORD_MIN_LENGTH:    util.StringPointerOrEmpty(plan.PasswordMinLength),
		PASSWORD_MAX_LENGTH:    util.StringPointerOrEmpty(plan.PasswordMaxLength),
		PASSWORD_NOOFCAPSALPHA: util.StringPointerOrEmpty(plan.PasswordNoOfCapsAlpha),
		PASSWORD_NOOFDIGITS:    util.StringPointerOrEmpty(plan.PasswordNoOfDigits),
		PASSWORD_NOOFSPLCHARS:  util.StringPointerOrEmpty(plan.PasswordNoOfSplChars),

		// Account management JSON configurations
		CREATEACCOUNTJSON:  util.StringPointerOrEmpty(plan.CreateAccountJson),
		UPDATEACCOUNTJSON:  util.StringPointerOrEmpty(plan.UpdateAccountJson),
		GRANTACCESSJSON:    util.StringPointerOrEmpty(plan.GrantAccessJson),
		REVOKEACCESSJSON:   util.StringPointerOrEmpty(plan.RevokeAccessJson),
		CHANGEPASSJSON:     util.StringPointerOrEmpty(config.ChangePassJson),
		DELETEACCOUNTJSON:  util.StringPointerOrEmpty(plan.DeleteAccountJson),
		ENABLEACCOUNTJSON:  util.StringPointerOrEmpty(plan.EnableAccountJson),
		DISABLEACCOUNTJSON: util.StringPointerOrEmpty(plan.DisableAccountJson),
		ACCOUNTEXISTSJSON:  util.StringPointerOrEmpty(plan.AccountExistsJson),
		UPDATEUSERJSON:     util.StringPointerOrEmpty(plan.UpdateUserJson),
		MODIFYUSERDATAJSON: util.StringPointerOrEmpty(plan.ModifyUserDataJson),

		// Import configurations
		ACCOUNTSIMPORT:         util.StringPointerOrEmpty(plan.AccountsImport),
		ENTITLEMENTVALUEIMPORT: util.StringPointerOrEmpty(plan.EntitlementValueImport),
		ROLEOWNERIMPORT:        util.StringPointerOrEmpty(plan.RoleOwnerImport),
		ROLESIMPORT:            util.StringPointerOrEmpty(plan.RolesImport),
		SYSTEMIMPORT:           util.StringPointerOrEmpty(plan.SystemImport),
		USERIMPORT:             util.StringPointerOrEmpty(plan.UserImport),

		// Additional configurations
		STATUS_THRESHOLD_CONFIG: util.StringPointerOrEmpty(plan.StatusThresholdConfig),
		MAX_PAGINATION_SIZE:     util.StringPointerOrEmpty(plan.MaxPaginationSize),
		CLI_COMMAND_JSON:        util.StringPointerOrEmpty(plan.CliCommandJson),

		// Entitlement management (TER-176)
		CREATEENTITLEMENTJSON: util.StringPointerOrEmpty(plan.CreateEntitlementJson),
		DELETEENTITLEMENTJSON: util.StringPointerOrEmpty(plan.DeleteEntitlementJson),
		ENTITLEMENTEXISTJSON:  util.StringPointerOrEmpty(plan.EntitlementExistJson),
		UPDATEENTITLEMENTJSON: util.StringPointerOrEmpty(plan.UpdateEntitlementJson),
	}

	// Handle vault configuration
	if plan.VaultConnection.ValueString() != "" {
		connector.BaseConnector.VaultConnection = util.SafeStringConnector(plan.VaultConnection.ValueString())
		connector.BaseConnector.VaultConfiguration = util.SafeStringConnector(plan.VaultConfiguration.ValueString())
		connector.BaseConnector.Saveinvault = util.SafeStringConnector(plan.SaveInVault.ValueString())
	}

	return connector
}

// UpdateModelFromCreateResponse updates the model from create response
func (r *DBConnectionResource) UpdateModelFromCreateResponse(plan *DBConnectorResourceModel, apiResp *openapi.CreateOrUpdateResponse) {
	plan.ID = types.StringValue(fmt.Sprintf("%d", *apiResp.ConnectionKey))
	plan.ConnectionKey = types.Int64Value(int64(*apiResp.ConnectionKey))
	plan.ErrorCode = types.StringValue(util.SafeDeref(apiResp.ErrorCode))
	plan.Msg = types.StringValue(util.SafeDeref(apiResp.Msg))

	// Set computed values for optional fields
	plan.Description = util.SafeStringDatasource(plan.Description.ValueStringPointer())
	plan.DefaultSavRoles = util.SafeStringDatasource(plan.DefaultSavRoles.ValueStringPointer())
	plan.EmailTemplate = util.SafeStringDatasource(plan.EmailTemplate.ValueStringPointer())
	plan.ConnectionProperties = util.SafeStringDatasource(plan.ConnectionProperties.ValueStringPointer())
	plan.PasswordMinLength = util.SafeStringDatasource(plan.PasswordMinLength.ValueStringPointer())
	plan.PasswordMaxLength = util.SafeStringDatasource(plan.PasswordMaxLength.ValueStringPointer())
	plan.PasswordNoOfCapsAlpha = util.SafeStringDatasource(plan.PasswordNoOfCapsAlpha.ValueStringPointer())
	plan.PasswordNoOfDigits = util.SafeStringDatasource(plan.PasswordNoOfDigits.ValueStringPointer())
	plan.PasswordNoOfSplChars = util.SafeStringDatasource(plan.PasswordNoOfSplChars.ValueStringPointer())
	plan.CreateAccountJson = util.SafeStringDatasource(plan.CreateAccountJson.ValueStringPointer())
	plan.UpdateAccountJson = util.SafeStringDatasource(plan.UpdateAccountJson.ValueStringPointer())
	plan.GrantAccessJson = util.SafeStringDatasource(plan.GrantAccessJson.ValueStringPointer())
	plan.RevokeAccessJson = util.SafeStringDatasource(plan.RevokeAccessJson.ValueStringPointer())
	plan.DeleteAccountJson = util.SafeStringDatasource(plan.DeleteAccountJson.ValueStringPointer())
	plan.EnableAccountJson = util.SafeStringDatasource(plan.EnableAccountJson.ValueStringPointer())
	plan.DisableAccountJson = util.SafeStringDatasource(plan.DisableAccountJson.ValueStringPointer())
	plan.AccountExistsJson = util.SafeStringDatasource(plan.AccountExistsJson.ValueStringPointer())
	plan.UpdateUserJson = util.SafeStringDatasource(plan.UpdateUserJson.ValueStringPointer())
	plan.ModifyUserDataJson = util.SafeStringDatasource(plan.ModifyUserDataJson.ValueStringPointer())
	plan.AccountsImport = util.SafeStringDatasource(plan.AccountsImport.ValueStringPointer())
	plan.EntitlementValueImport = util.SafeStringDatasource(plan.EntitlementValueImport.ValueStringPointer())
	plan.RoleOwnerImport = util.SafeStringDatasource(plan.RoleOwnerImport.ValueStringPointer())
	plan.RolesImport = util.SafeStringDatasource(plan.RolesImport.ValueStringPointer())
	plan.SystemImport = util.SafeStringDatasource(plan.SystemImport.ValueStringPointer())
	plan.UserImport = util.SafeStringDatasource(plan.UserImport.ValueStringPointer())
	plan.StatusThresholdConfig = util.SafeStringDatasource(plan.StatusThresholdConfig.ValueStringPointer())
	plan.MaxPaginationSize = util.SafeStringDatasource(plan.MaxPaginationSize.ValueStringPointer())
	plan.CliCommandJson = util.SafeStringDatasource(plan.CliCommandJson.ValueStringPointer())
	plan.CreateEntitlementJson = util.SafeStringDatasource(plan.CreateEntitlementJson.ValueStringPointer())
	plan.DeleteEntitlementJson = util.SafeStringDatasource(plan.DeleteEntitlementJson.ValueStringPointer())
	plan.EntitlementExistJson = util.SafeStringDatasource(plan.EntitlementExistJson.ValueStringPointer())
	plan.UpdateEntitlementJson = util.SafeStringDatasource(plan.UpdateEntitlementJson.ValueStringPointer())
}

func (r *DBConnectionResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state DBConnectorResourceModel

	opCtx := errorsutil.CreateOperationContext(errorsutil.ConnectorTypeDB, "terraform_read", "")
	ctx = opCtx.AddContextToLogger(ctx)

	opCtx.LogOperationStart(ctx, "Starting DB connection resource read")

	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		errorCode := dbErrorCodes.StateExtraction()
		opCtx.LogOperationError(ctx, "Failed to get state from request", errorCode,
			fmt.Errorf("state extraction failed"))
		resp.Diagnostics.AddError(
			errorsutil.GetErrorMessage(errorsutil.ErrStateExtraction),
			fmt.Sprintf("[%s] Unable to extract Terraform state from request", errorCode),
		)
		return
	}

	connectionName := state.ConnectionName.ValueString()
	opCtx.ConnectionName = connectionName

	opCtx.LogOperationStart(ctx, "Reading DB connection details",
		map[string]interface{}{"connection_name": connectionName})

	// Use interface pattern instead of direct API client creation
	apiResp, err := r.ReadDBConnection(ctx, connectionName)
	if err != nil {
		opCtx.LogOperationError(ctx, "DB connection read failed", "", err)
		resp.Diagnostics.AddError(
			"DB Connection Read Failed",
			err.Error(),
		)
		return
	}

	r.UpdateModelFromReadResponse(&state, apiResp)

	stateDiagnostics := resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(stateDiagnostics...)
	if resp.Diagnostics.HasError() {
		errorCode := dbErrorCodes.StateUpdate()
		opCtx.LogOperationError(ctx, "Failed to set state", errorCode,
			fmt.Errorf("state update failed"))
		resp.Diagnostics.AddError(
			errorsutil.GetErrorMessage(errorsutil.ErrStateUpdate),
			fmt.Sprintf("[%s] Unable to update Terraform state for connection '%s'", errorCode, connectionName),
		)
		return
	}

	opCtx.LogOperationEnd(ctx, "DB connection resource read completed successfully",
		map[string]interface{}{"connection_name": connectionName})
}

// ReadDBConnection reads a DB connection
func (r *DBConnectionResource) ReadDBConnection(ctx context.Context, connectionName string) (*openapi.GetConnectionDetailsResponse, error) {
	opCtx := errorsutil.CreateOperationContext(errorsutil.ConnectorTypeDB, "read", connectionName)

	// Create logging context (separate from API context)
	logCtx := opCtx.AddContextToLogger(ctx)

	opCtx.LogOperationStart(logCtx, "Starting DB connection read",
		map[string]interface{}{"connection_name": connectionName})

	// Execute read operation with retry logic
	var apiResp *openapi.GetConnectionDetailsResponse

	err := r.provider.AuthenticatedAPICallWithRetry(ctx, "read_db_connection", func(token string) error {
		connectionOps := r.connectionFactory.CreateConnectionOperations(r.client.APIBaseURL(), token)
		resp, httpResp, err := connectionOps.GetConnectionDetails(ctx, connectionName)
		if httpResp != nil && httpResp.StatusCode == 401 {
			return fmt.Errorf("401 unauthorized")
		}
		apiResp = resp
		return err
	})

	if err != nil {
		errorCode := dbErrorCodes.ReadFailed()
		opCtx.LogOperationError(logCtx, "Failed to read DB connection", errorCode, err)
		return nil, errorsutil.CreateStandardError(errorsutil.ConnectorTypeDB, errorCode, "read", connectionName, err)
	}

	if apiResp != nil && apiResp.DBConnectionResponse != nil && apiResp.DBConnectionResponse.Errorcode != nil && *apiResp.DBConnectionResponse.Errorcode != 0 {
		apiErr := fmt.Errorf("API returned error code %d: %s", *apiResp.DBConnectionResponse.Errorcode, errorsutil.SanitizeMessage(apiResp.DBConnectionResponse.Msg))
		errorCode := dbErrorCodes.APIError()
		opCtx.LogOperationError(ctx, "DB connection read failed with API error", errorCode, apiErr,
			map[string]interface{}{
				"api_error_code": *apiResp.DBConnectionResponse.Errorcode,
				"message":        errorsutil.SanitizeMessage(apiResp.DBConnectionResponse.Msg),
			})
		return nil, errorsutil.CreateStandardError(errorsutil.ConnectorTypeDB, errorCode, "read", connectionName, apiErr)
	}

	opCtx.LogOperationEnd(logCtx, "DB connection read completed successfully",
		map[string]interface{}{
			"connection_name": connectionName,
			"connection_key":  apiResp.DBConnectionResponse.Connectionkey,
		})

	return apiResp, nil
}

// UpdateModelFromReadResponse updates the model from read response
func (r *DBConnectionResource) UpdateModelFromReadResponse(state *DBConnectorResourceModel, apiResp *openapi.GetConnectionDetailsResponse) {
	if apiResp.DBConnectionResponse == nil {
		return
	}

	dbResp := apiResp.DBConnectionResponse

	// Update base connector fields
	state.ID = types.StringValue(fmt.Sprintf("%d", *dbResp.Connectionkey))
	state.ConnectionKey = types.Int64Value(int64(*dbResp.Connectionkey))
	state.ConnectionName = util.SafeStringDatasource(dbResp.Connectionname)
	state.Description = util.SafeStringDatasource(dbResp.Description)
	state.DefaultSavRoles = util.SafeStringDatasource(dbResp.Defaultsavroles)
	state.EmailTemplate = util.SafeStringDatasource(dbResp.Emailtemplate)

	// Update DB-specific fields from connection attributes
	if dbResp.Connectionattributes != nil {
		attrs := dbResp.Connectionattributes

		// Required fields
		state.URL = util.SafeStringDatasource(attrs.URL)
		state.Username = util.SafeStringDatasource(attrs.USERNAME)
		state.DriverName = util.SafeStringDatasource(attrs.DRIVERNAME)

		// Optional configuration fields
		state.ConnectionProperties = util.SafeStringDatasource(attrs.CONNECTIONPROPERTIES)
		state.PasswordMinLength = util.SafeStringDatasource(attrs.PASSWORD_MIN_LENGTH)
		state.PasswordMaxLength = util.SafeStringDatasource(attrs.PASSWORD_MAX_LENGTH)
		state.PasswordNoOfCapsAlpha = util.SafeStringDatasource(attrs.PASSWORD_NOOFCAPSALPHA)
		state.PasswordNoOfDigits = util.SafeStringDatasource(attrs.PASSWORD_NOOFDIGITS)
		state.PasswordNoOfSplChars = util.SafeStringDatasource(attrs.PASSWORD_NOOFSPLCHARS)

		// Account management JSON configurations
		state.CreateAccountJson = util.SafeStringDatasource(attrs.CREATEACCOUNTJSON)
		state.UpdateAccountJson = util.SafeStringDatasource(attrs.UPDATEACCOUNTJSON)
		state.GrantAccessJson = util.SafeStringDatasource(attrs.GRANTACCESSJSON)
		state.RevokeAccessJson = util.SafeStringDatasource(attrs.REVOKEACCESSJSON)
		state.DeleteAccountJson = util.SafeStringDatasource(attrs.DELETEACCOUNTJSON)
		state.EnableAccountJson = util.SafeStringDatasource(attrs.ENABLEACCOUNTJSON)
		state.DisableAccountJson = util.SafeStringDatasource(attrs.DISABLEACCOUNTJSON)
		state.AccountExistsJson = util.SafeStringDatasource(attrs.ACCOUNTEXISTSJSON)
		state.UpdateUserJson = util.SafeStringDatasource(attrs.UPDATEUSERJSON)
		state.ModifyUserDataJson = util.SafeStringDatasource(attrs.MODIFYUSERDATAJSON)

		// Import configurations
		state.AccountsImport = util.SafeStringDatasource(attrs.ACCOUNTSIMPORT)
		state.EntitlementValueImport = util.SafeStringDatasource(attrs.ENTITLEMENTVALUEIMPORT)
		state.RoleOwnerImport = util.SafeStringDatasource(attrs.ROLEOWNERIMPORT)
		state.RolesImport = util.SafeStringDatasource(attrs.ROLESIMPORT)
		state.SystemImport = util.SafeStringDatasource(attrs.SYSTEMIMPORT)
		state.UserImport = util.SafeStringDatasource(attrs.USERIMPORT)

		// Additional configurations
		state.StatusThresholdConfig = util.SafeStringDatasource(attrs.STATUS_THRESHOLD_CONFIG)
		state.MaxPaginationSize = util.SafeStringDatasource(attrs.MAX_PAGINATION_SIZE)
		state.CliCommandJson = util.SafeStringDatasource(attrs.CLI_COMMAND_JSON)

		// Entitlement management (TER-176)
		state.CreateEntitlementJson = util.SafeStringDatasource(attrs.CREATEENTITLEMENTJSON)
		state.DeleteEntitlementJson = util.SafeStringDatasource(attrs.DELETEENTITLEMENTJSON)
		state.EntitlementExistJson = util.SafeStringDatasource(attrs.ENTITLEMENTEXISTJSON)
		state.UpdateEntitlementJson = util.SafeStringDatasource(attrs.UPDATEENTITLEMENTJSON)
	}

	// Update response message and error code
	apiMessage := util.SafeDeref(dbResp.Msg)
	if apiMessage == "success" {
		state.Msg = types.StringValue("Connection Successful")
	} else {
		state.Msg = types.StringValue(apiMessage)
	}
	state.ErrorCode = util.Int32PtrToTFString(dbResp.Errorcode)
}

func (r *DBConnectionResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state, config DBConnectorResourceModel

	opCtx := errorsutil.CreateOperationContext(errorsutil.ConnectorTypeDB, "terraform_update", "")
	ctx = opCtx.AddContextToLogger(ctx)

	opCtx.LogOperationStart(ctx, "Starting DB connection resource update")

	// Extract state from request
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		errorCode := dbErrorCodes.StateExtraction()
		opCtx.LogOperationError(ctx, "Failed to get state from request", errorCode,
			fmt.Errorf("state extraction failed"))
		resp.Diagnostics.AddError(
			errorsutil.GetErrorMessage(errorsutil.ErrStateExtraction),
			fmt.Sprintf("[%s] Unable to extract Terraform state from request", errorCode),
		)
		return
	}

	connectionName := state.ConnectionName.ValueString()
	opCtx.ConnectionName = connectionName

	// Extract plan from request
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		errorCode := dbErrorCodes.PlanExtraction()
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
		errorCode := dbErrorCodes.ConfigExtraction()
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
		errorCode := dbErrorCodes.NameImmutable()
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

	opCtx.LogOperationStart(ctx, "Executing DB connection update",
		map[string]interface{}{"connection_name": connectionName})

	// Use interface pattern instead of direct API client creation
	_, err := r.UpdateDBConnection(ctx, &plan, &config)
	if err != nil {
		opCtx.LogOperationError(ctx, "DB connection update failed", "", err)
		resp.Diagnostics.AddError(
			"DB Connection Update Failed",
			err.Error(),
		)
		return
	}

	// Read the updated connection to get the latest state
	getResp, err := r.ReadDBConnection(ctx, connectionName)
	if err != nil {
		opCtx.LogOperationError(ctx, "Failed to read updated DB connection", "", err)
		resp.Diagnostics.AddError(
			"DB Connection Post-Update Read Failed",
			err.Error(),
		)
		return
	}

	r.UpdateModelFromReadResponse(&plan, getResp)

	stateUpdateDiagnostics := resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(stateUpdateDiagnostics...)
	if resp.Diagnostics.HasError() {
		errorCode := dbErrorCodes.StateUpdate()
		opCtx.LogOperationError(ctx, "Failed to update state after successful update", errorCode,
			fmt.Errorf("state update failed"))
		resp.Diagnostics.AddError(
			errorsutil.GetErrorMessage(errorsutil.ErrStateUpdate),
			fmt.Sprintf("[%s] Unable to update Terraform state after successful update for connection '%s'", errorCode, connectionName),
		)
		return
	}

	opCtx.LogOperationEnd(ctx, "DB connection resource update completed successfully",
		map[string]interface{}{"connection_name": connectionName})
}

// UpdateDBConnection updates a DB connection
func (r *DBConnectionResource) UpdateDBConnection(ctx context.Context, plan *DBConnectorResourceModel, config *DBConnectorResourceModel) (*openapi.CreateOrUpdateResponse, error) {
	connectionName := plan.ConnectionName.ValueString()
	opCtx := errorsutil.CreateOperationContext(errorsutil.ConnectorTypeDB, "update", connectionName)

	// Create logging context (separate from API context)
	logCtx := opCtx.AddContextToLogger(ctx)

	opCtx.LogOperationStart(logCtx, "Starting DB connection update",
		map[string]interface{}{"connection_name": connectionName})

	// Build DB connector request
	dbConn := r.BuildDBConnector(plan, config)

	if plan.VaultConnection.ValueString() == "" {
		emptyStr := ""
		dbConn.BaseConnector.VaultConnection = &emptyStr
		dbConn.BaseConnector.VaultConfiguration = &emptyStr
		dbConn.BaseConnector.Saveinvault = &emptyStr
	}

	dbConnRequest := openapi.CreateOrUpdateRequest{
		DBConnector: &dbConn,
	}

	opCtx.LogOperationStart(logCtx, "Executing DB connection update API call")

	// Execute update operation with retry logic
	var apiResp *openapi.CreateOrUpdateResponse

	err := r.provider.AuthenticatedAPICallWithRetry(ctx, "update_db_connection", func(token string) error {
		connectionOps := r.connectionFactory.CreateConnectionOperations(r.client.APIBaseURL(), token)
		resp, httpResp, err := connectionOps.CreateOrUpdateConnection(ctx, dbConnRequest)
		if httpResp != nil && httpResp.StatusCode == 401 {
			return fmt.Errorf("401 unauthorized")
		}
		apiResp = resp
		return err
	})

	if err != nil {
		errorCode := dbErrorCodes.UpdateFailed()
		opCtx.LogOperationError(logCtx, "Failed to update DB connection", errorCode, err)
		return nil, errorsutil.CreateStandardError(errorsutil.ConnectorTypeDB, errorCode, "update", connectionName, err)
	}

	if apiResp != nil && apiResp.ErrorCode != nil && *apiResp.ErrorCode != "0" {
		apiErr := fmt.Errorf("API returned error code %s: %s", *apiResp.ErrorCode, errorsutil.SanitizeMessage(apiResp.Msg))
		errorCode := dbErrorCodes.APIError()
		opCtx.LogOperationError(logCtx, "DB connection update failed with API error", errorCode, apiErr,
			map[string]interface{}{
				"api_error_code": *apiResp.ErrorCode,
				"message":        errorsutil.SanitizeMessage(apiResp.Msg),
			})
		return nil, errorsutil.CreateStandardError(errorsutil.ConnectorTypeDB, errorCode, "update", connectionName, apiErr)
	}

	opCtx.LogOperationEnd(logCtx, "DB connection updated successfully",
		map[string]interface{}{
			"connection_name": connectionName,
			"connection_key":  apiResp.ConnectionKey,
		})

	return apiResp, nil
}

func (r *DBConnectionResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
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

func (r *DBConnectionResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	// Importing a DB connection resource requires the connection name
	connectionName := req.ID
	opCtx := errorsutil.CreateOperationContext(errorsutil.ConnectorTypeDB, "terraform_import", connectionName)
	ctx = opCtx.AddContextToLogger(ctx)

	opCtx.LogOperationStart(ctx, "Starting DB connection resource import")

	// Retrieve import ID and save to id attribute
	resource.ImportStatePassthroughID(ctx, path.Root("connection_name"), req, resp)

	opCtx.LogOperationEnd(ctx, "DB connection resource import completed successfully",
		map[string]interface{}{"connection_name": connectionName})
}
