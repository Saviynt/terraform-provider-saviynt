// Copyright (c) 2025 Saviynt Inc.
// SPDX-License-Identifier: MPL-2.0

// saviynt_db_connection_datasource retrieves db connections details from the Saviynt Security Manager.
// The data source supports a single Read operation to look up an existing db connections by name.
package provider

import (
	"context"
	"fmt"
	"terraform-provider-Saviynt/internal/client"
	"terraform-provider-Saviynt/util"
	connectionsutil "terraform-provider-Saviynt/util/connectionsutil"
	"terraform-provider-Saviynt/util/errorsutil"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	openapi "github.com/saviynt/saviynt-api-go-client/connections"
)

var _ datasource.DataSource = &DbConnectionsDataSource{}

// Initialize error codes for DB Connection datasource operations
var dbDatasourceErrorCodes = errorsutil.NewConnectorErrorCodes(errorsutil.ConnectorTypeDB)

// DBConnectionsDataSource defines the data source
type DbConnectionsDataSource struct {
	client            client.SaviyntClientInterface
	token             string
	provider          client.SaviyntProviderInterface
	connectionFactory client.ConnectionFactoryInterface
}

type DBConnectionDataSourceModel struct {
	ID types.String `tfsdk:"id"`
	BaseConnectionDataSourceModel
	ConnectionAttributes *DBConnectionAttributes `tfsdk:"connection_attributes"`
}

type DBConnectionAttributes struct {
	PasswordMinLength        types.String             `tfsdk:"password_min_length"`
	AccountExistsJSON        types.String             `tfsdk:"accountexists_json"`
	RolesImport              types.String             `tfsdk:"roles_import"`
	RoleOwnerImport          types.String             `tfsdk:"roleowner_import"`
	CreateAccountJSON        types.String             `tfsdk:"createaccount_json"`
	UserImport               types.String             `tfsdk:"user_import"`
	DisableAccountJSON       types.String             `tfsdk:"disableaccount_json"`
	EntitlementValueImport   types.String             `tfsdk:"entitlementvalue_import"`
	ConnectionType           types.String             `tfsdk:"connection_type"`
	UpdateUserJSON           types.String             `tfsdk:"updateuser_json"`
	PasswordNoOfSplChars     types.String             `tfsdk:"password_noofsplchars"`
	RevokeAccessJSON         types.String             `tfsdk:"revokeaccess_json"`
	URL                      types.String             `tfsdk:"url"`
	SystemImport             types.String             `tfsdk:"system_import"`
	DriverName               types.String             `tfsdk:"drivername"`
	DeleteAccountJSON        types.String             `tfsdk:"deleteaccount_json"`
	StatusThresholdConfig    types.String             `tfsdk:"status_threshold_config"`
	IsTimeoutSupported       types.Bool               `tfsdk:"is_timeout_supported"`
	PasswordNoOfCapsAlpha    types.String             `tfsdk:"password_noofcapsalpha"`
	PasswordNoOfDigits       types.String             `tfsdk:"password_noofdigits"`
	ConnectionProperties     types.String             `tfsdk:"connectionproperties"`
	ModifyUserDataJSON       types.String             `tfsdk:"modifyuserdata_json"`
	IsTimeoutConfigValidated types.Bool               `tfsdk:"is_timeout_config_validated"`
	AccountsImport           types.String             `tfsdk:"accounts_import"`
	EnableAccountJSON        types.String             `tfsdk:"enableaccount_json"`
	PasswordMaxLength        types.String             `tfsdk:"password_max_length"`
	MaxPaginationSize        types.String             `tfsdk:"max_pagination_size"`
	UpdateAccountJSON        types.String             `tfsdk:"updateaccount_json"`
	GrantAccessJSON          types.String             `tfsdk:"grantaccess_json"`
	CliCommandJSON           types.String             `tfsdk:"cli_command_json"`
	ConnectionTimeoutConfig  *ConnectionTimeoutConfig `tfsdk:"connection_timeout_config"`
	//TER-176
	CreateEntitlementJson types.String `tfsdk:"create_entitlement_json"`
	DeleteEntitlementJson types.String `tfsdk:"delete_entitlement_json"`
	EntitlementExistJson  types.String `tfsdk:"entitlement_exist_json"`
	UpdateEntitlementJson types.String `tfsdk:"update_entitlement_json"`
}

var _ datasource.DataSource = &DbConnectionsDataSource{}

// NewDBConnectionsDataSource creates a new DB connections data source with default factory
func NewDBConnectionsDataSource() datasource.DataSource {
	return &DbConnectionsDataSource{
		connectionFactory: &client.DefaultConnectionFactory{},
	}
}

// NewDBConnectionsDataSourceWithFactory creates a new DB connections data source with custom factory
// Used primarily for testing with mock factories
func NewDBConnectionsDataSourceWithFactory(factory client.ConnectionFactoryInterface) datasource.DataSource {
	return &DbConnectionsDataSource{
		connectionFactory: factory,
	}
}

// SetClient sets the client for testing purposes
func (d *DbConnectionsDataSource) SetClient(client client.SaviyntClientInterface) {
	d.client = client
}

// SetToken sets the token for testing purposes
func (d *DbConnectionsDataSource) SetToken(token string) {
	d.token = token
}

// SetProvider sets the provider for testing purposes
func (d *DbConnectionsDataSource) SetProvider(provider client.SaviyntProviderInterface) {
	d.provider = provider
}

// Metadata sets the data source type name for Terraform
func (d *DbConnectionsDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "saviynt_db_connection_datasource"
}

func DBConnectorsDataSourceSchema() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:    true,
			Description: "Resource ID.",
		},
		"connection_attributes": schema.SingleNestedAttribute{
			Computed: true,
			Attributes: map[string]schema.Attribute{
				"password_min_length":         schema.StringAttribute{Computed: true},
				"accountexists_json":          schema.StringAttribute{Computed: true},
				"roles_import":                schema.StringAttribute{Computed: true},
				"roleowner_import":            schema.StringAttribute{Computed: true},
				"createaccount_json":          schema.StringAttribute{Computed: true},
				"user_import":                 schema.StringAttribute{Computed: true},
				"disableaccount_json":         schema.StringAttribute{Computed: true},
				"entitlementvalue_import":     schema.StringAttribute{Computed: true},
				"connection_type":             schema.StringAttribute{Computed: true},
				"updateuser_json":             schema.StringAttribute{Computed: true},
				"password_noofsplchars":       schema.StringAttribute{Computed: true},
				"revokeaccess_json":           schema.StringAttribute{Computed: true},
				"url":                         schema.StringAttribute{Computed: true},
				"system_import":               schema.StringAttribute{Computed: true},
				"drivername":                  schema.StringAttribute{Computed: true},
				"deleteaccount_json":          schema.StringAttribute{Computed: true},
				"status_threshold_config":     schema.StringAttribute{Computed: true},
				"is_timeout_supported":        schema.BoolAttribute{Computed: true},
				"password_noofcapsalpha":      schema.StringAttribute{Computed: true},
				"password_noofdigits":         schema.StringAttribute{Computed: true},
				"connectionproperties":        schema.StringAttribute{Computed: true},
				"modifyuserdata_json":         schema.StringAttribute{Computed: true},
				"is_timeout_config_validated": schema.BoolAttribute{Computed: true},
				"accounts_import":             schema.StringAttribute{Computed: true},
				"enableaccount_json":          schema.StringAttribute{Computed: true},
				"password_max_length":         schema.StringAttribute{Computed: true},
				"max_pagination_size":         schema.StringAttribute{Computed: true},
				"updateaccount_json":          schema.StringAttribute{Computed: true},
				"grantaccess_json":            schema.StringAttribute{Computed: true},
				"cli_command_json":            schema.StringAttribute{Computed: true},
				"connection_timeout_config": schema.SingleNestedAttribute{
					Computed:   true,
					Attributes: ConnectionTimeoutConfigeSchema(),
				},
				//TER-176
				"create_entitlement_json": schema.StringAttribute{Computed: true},
				"delete_entitlement_json": schema.StringAttribute{Computed: true},
				"entitlement_exist_json":  schema.StringAttribute{Computed: true},
				"update_entitlement_json": schema.StringAttribute{Computed: true},
			},
		},
	}
}

// Schema defines the structure and attributes available for the DB connection data source
func (d *DbConnectionsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: util.DBConnDataSourceDescription,
		Attributes:  connectionsutil.MergeDataSourceAttributes(BaseConnectorDataSourceSchema(), DBConnectorsDataSourceSchema()),
	}
}

// Configure initializes the data source with provider configuration
// Sets up client and authentication token for API operations
func (d *DbConnectionsDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	opCtx := errorsutil.CreateOperationContext(errorsutil.ConnectorTypeDB, "configure", "")
	ctx = opCtx.AddContextToLogger(ctx)

	opCtx.LogOperationStart(ctx, "Starting DB connection datasource configuration")

	// Check if provider data is available.
	if req.ProviderData == nil {
		tflog.Debug(ctx, "ProviderData is nil, returning early")
		opCtx.LogOperationEnd(ctx, "DB connection datasource configuration completed - no provider data")
		return
	}

	// Cast provider data to your provider type.
	prov, ok := req.ProviderData.(*SaviyntProvider)
	if !ok {
		errorCode := dbDatasourceErrorCodes.ProviderConfig()
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
	d.client = &client.SaviyntClientWrapper{Client: prov.client}
	d.token = prov.accessToken
	d.provider = &client.SaviyntProviderWrapper{Provider: prov}

	opCtx.LogOperationEnd(ctx, "DB connection datasource configured successfully")
}

// Read retrieves DB connection details from Saviynt and populates the Terraform state
// Supports lookup by connection name or connection key with comprehensive error handling
func (d *DbConnectionsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var state DBConnectionDataSourceModel

	opCtx := errorsutil.CreateOperationContext(errorsutil.ConnectorTypeDB, "datasource_read", "")
	ctx = opCtx.AddContextToLogger(ctx)

	opCtx.LogOperationStart(ctx, "Starting DB connection datasource read")

	// Extract configuration from request
	configDiagnostics := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(configDiagnostics...)
	if resp.Diagnostics.HasError() {
		errorCode := dbDatasourceErrorCodes.ConfigExtraction()
		opCtx.LogOperationError(ctx, "Failed to get config from request", errorCode,
			fmt.Errorf("config extraction failed"))
		resp.Diagnostics.AddError(
			errorsutil.GetErrorMessage(errorsutil.ErrConfigExtraction),
			fmt.Sprintf("[%s] Unable to extract Terraform configuration from request", errorCode),
		)
		return
	}

	// Update operation context with connection name if available
	connectionName := ""
	if !state.ConnectionName.IsNull() && state.ConnectionName.ValueString() != "" {
		connectionName = state.ConnectionName.ValueString()
		opCtx.ConnectionName = connectionName
		ctx = opCtx.AddContextToLogger(ctx)
	}

	// Prepare connection parameters
	var connectionKey *int64
	if !state.ConnectionKey.IsNull() {
		keyValue := state.ConnectionKey.ValueInt64()
		connectionKey = &keyValue
	}

	// Validate that at least one identifier is provided
	if connectionName == "" && connectionKey == nil {
		errorCode := dbDatasourceErrorCodes.MissingIdentifier()
		opCtx.LogOperationError(ctx, "Missing connection identifier", errorCode,
			fmt.Errorf("either connection_name or connection_key must be provided"))
		resp.Diagnostics.AddError(
			errorsutil.GetErrorMessage(errorsutil.ErrMissingIdentifier),
			fmt.Sprintf("[%s] Either 'connection_name' or 'connection_key' must be provided to look up the DB connection.", errorCode),
		)
		return
	}

	// Execute API call to get DB connection details
	apiResp, err := d.ReadDBConnectionDetails(ctx, connectionName, connectionKey)
	if err != nil {
		// Error is already sanitized in ReadDBConnectionDetails method
		errorCode := dbDatasourceErrorCodes.ReadFailed()
		opCtx.LogOperationError(ctx, "Failed to read DB connection details", errorCode, err)
		resp.Diagnostics.AddError(
			errorsutil.GetErrorMessage(errorsutil.ErrReadFailed),
			fmt.Sprintf("[%s] %s", errorCode, err.Error()),
		)
		return
	}

	// Validate API response
	if err := d.ValidateDBConnectionResponse(apiResp); err != nil {
		errorCode := dbDatasourceErrorCodes.APIError()
		opCtx.LogOperationError(ctx, "Invalid connection type for DB datasource", errorCode, err)
		resp.Diagnostics.AddError(
			errorsutil.GetErrorMessage(errorsutil.ErrAPIError),
			fmt.Sprintf("[%s] Unable to verify connection type for connection %q. The provider could not determine the type of this connection. Please ensure the connection name is correct and belongs to a supported connector type.", errorCode, connectionName),
		)
		return
	}

	// Map API response to state
	d.UpdateModelFromDBConnectionResponse(&state, apiResp)

	// Handle authentication logic
	d.HandleDBAuthenticationLogic(&state, resp)

	// Set final state
	stateDiagnostics := resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(stateDiagnostics...)
	if resp.Diagnostics.HasError() {
		errorCode := dbDatasourceErrorCodes.StateUpdate()
		opCtx.LogOperationError(ctx, "Failed to set state", errorCode,
			fmt.Errorf("state update failed"))
		resp.Diagnostics.AddError(
			errorsutil.GetErrorMessage(errorsutil.ErrStateUpdate),
			fmt.Sprintf("[%s] Unable to update Terraform state for DB connection datasource", errorCode),
		)
		return
	}

	opCtx.LogOperationEnd(ctx, "DB connection datasource read completed successfully",
		map[string]interface{}{
			"connection_name": connectionName,
			"has_attributes":  state.ConnectionAttributes != nil,
		})
}

// ReadDBConnectionDetails retrieves DB connection details from Saviynt API
// Handles both connection name and connection key based lookups using factory pattern
// Returns standardized errors with proper correlation tracking and sensitive data sanitization
func (d *DbConnectionsDataSource) ReadDBConnectionDetails(ctx context.Context, connectionName string, connectionKey *int64) (*openapi.GetConnectionDetailsResponse, error) {
	opCtx := errorsutil.CreateOperationContext(errorsutil.ConnectorTypeDB, "api_read", connectionName)
	logCtx := opCtx.AddContextToLogger(ctx)

	opCtx.LogOperationStart(logCtx, "Starting DB connection API call")

	tflog.Debug(logCtx, "Executing API request to get DB connection details")

	var apiResp *openapi.GetConnectionDetailsResponse

	// Execute API request with retry logic
	err := d.provider.AuthenticatedAPICallWithRetry(ctx, "read_db_connection_datasource", func(token string) error {
		connectionOps := d.connectionFactory.CreateConnectionOperations(d.client.APIBaseURL(), token)

		// Build request with both connectionname and connectionkey
		req := openapi.GetConnectionDetailsRequest{}
		if connectionName != "" {
			req.Connectionname = &connectionName
		}
		if connectionKey != nil {
			keyStr := fmt.Sprintf("%d", *connectionKey)
			req.Connectionkey = &keyStr
		}

		resp, httpResp, err := connectionOps.GetConnectionDetailsDataSource(ctx, req)
		if httpResp != nil && httpResp.StatusCode == 401 {
			return fmt.Errorf("401 unauthorized")
		}
		apiResp = resp
		return err
	})

	if err != nil {
		errorCode := dbDatasourceErrorCodes.ReadFailed()
		opCtx.LogOperationError(logCtx, "Failed to read DB connection details", errorCode, err)
		return nil, errorsutil.CreateStandardError(errorsutil.ConnectorTypeDB, errorCode, "api_read", connectionName, err)
	}

	opCtx.LogOperationEnd(logCtx, "DB connection API call completed successfully")

	return apiResp, nil
}

// ValidateDBConnectionResponse validates that the API response contains valid DB connection data
// Returns standardized error if validation fails
func (d *DbConnectionsDataSource) ValidateDBConnectionResponse(apiResp *openapi.GetConnectionDetailsResponse) error {
	if apiResp != nil && apiResp.DBConnectionResponse == nil {
		return fmt.Errorf("verify the connection type - DB connection response is nil")
	}
	return nil
}

// UpdateModelFromDBConnectionResponse maps API response data to the Terraform state model
// It handles both base connection fields and detailed connection attributes
func (d *DbConnectionsDataSource) UpdateModelFromDBConnectionResponse(state *DBConnectionDataSourceModel, apiResp *openapi.GetConnectionDetailsResponse) {
	// Map base connection fields
	d.MapBaseDBConnectionFields(state, apiResp)

	// Map connection attributes
	d.MapDBConnectionAttributes(state, apiResp)
}

// MapBaseDBConnectionFields maps basic connection fields from API response to state model
// These are common fields available for all connection types
func (d *DbConnectionsDataSource) MapBaseDBConnectionFields(state *DBConnectionDataSourceModel, apiResp *openapi.GetConnectionDetailsResponse) {
	state.Msg = util.SafeStringDatasource(apiResp.DBConnectionResponse.Msg)
	state.ErrorCode = util.SafeInt64(apiResp.DBConnectionResponse.Errorcode)
	connectionKey := util.SafeInt64(apiResp.DBConnectionResponse.Connectionkey)
	state.ID = types.StringValue(fmt.Sprintf("ds-db-%d", connectionKey.ValueInt64()))
	state.ConnectionName = util.SafeStringDatasource(apiResp.DBConnectionResponse.Connectionname)
	state.ConnectionKey = util.SafeInt64(apiResp.DBConnectionResponse.Connectionkey)
	state.Description = util.SafeStringDatasource(apiResp.DBConnectionResponse.Description)
	state.DefaultSavRoles = util.SafeStringDatasource(apiResp.DBConnectionResponse.Defaultsavroles)
	state.ConnectionType = util.SafeStringDatasource(apiResp.DBConnectionResponse.Connectiontype)
	state.CreatedOn = util.SafeStringDatasource(apiResp.DBConnectionResponse.Createdon)
	state.CreatedBy = util.SafeStringDatasource(apiResp.DBConnectionResponse.Createdby)
	state.UpdatedBy = util.SafeStringDatasource(apiResp.DBConnectionResponse.Updatedby)
	state.EmailTemplate = util.SafeStringDatasource(apiResp.DBConnectionResponse.Emailtemplate)
}

// MapDBConnectionAttributes maps detailed DB connection attributes from API response to state model
// Returns nil if no connection attributes are present in the response
func (d *DbConnectionsDataSource) MapDBConnectionAttributes(state *DBConnectionDataSourceModel, apiResp *openapi.GetConnectionDetailsResponse) {
	if apiResp.DBConnectionResponse.Connectionattributes == nil {
		state.ConnectionAttributes = nil
		return
	}

	attrs := apiResp.DBConnectionResponse.Connectionattributes
	state.ConnectionAttributes = &DBConnectionAttributes{
		PasswordMinLength:        util.SafeStringDatasource(attrs.PASSWORD_MIN_LENGTH),
		AccountExistsJSON:        util.SafeStringDatasource(attrs.ACCOUNTEXISTSJSON),
		RolesImport:              util.SafeStringDatasource(attrs.ROLESIMPORT),
		RoleOwnerImport:          util.SafeStringDatasource(attrs.ROLEOWNERIMPORT),
		CreateAccountJSON:        util.SafeStringDatasource(attrs.CREATEACCOUNTJSON),
		UserImport:               util.SafeStringDatasource(attrs.USERIMPORT),
		DisableAccountJSON:       util.SafeStringDatasource(attrs.DISABLEACCOUNTJSON),
		EntitlementValueImport:   util.SafeStringDatasource(attrs.ENTITLEMENTVALUEIMPORT),
		ConnectionType:           util.SafeStringDatasource(attrs.ConnectionType),
		UpdateUserJSON:           util.SafeStringDatasource(attrs.UPDATEUSERJSON),
		PasswordNoOfSplChars:     util.SafeStringDatasource(attrs.PASSWORD_NOOFSPLCHARS),
		RevokeAccessJSON:         util.SafeStringDatasource(attrs.REVOKEACCESSJSON),
		URL:                      util.SafeStringDatasource(attrs.URL),
		SystemImport:             util.SafeStringDatasource(attrs.SYSTEMIMPORT),
		DriverName:               util.SafeStringDatasource(attrs.DRIVERNAME),
		DeleteAccountJSON:        util.SafeStringDatasource(attrs.DELETEACCOUNTJSON),
		StatusThresholdConfig:    util.SafeStringDatasource(attrs.STATUS_THRESHOLD_CONFIG),
		IsTimeoutSupported:       util.SafeBoolDatasource(attrs.IsTimeoutSupported),
		PasswordNoOfCapsAlpha:    util.SafeStringDatasource(attrs.PASSWORD_NOOFCAPSALPHA),
		PasswordNoOfDigits:       util.SafeStringDatasource(attrs.PASSWORD_NOOFDIGITS),
		ConnectionProperties:     util.SafeStringDatasource(attrs.CONNECTIONPROPERTIES),
		ModifyUserDataJSON:       util.SafeStringDatasource(attrs.MODIFYUSERDATAJSON),
		IsTimeoutConfigValidated: util.SafeBoolDatasource(attrs.IsTimeoutConfigValidated),
		AccountsImport:           util.SafeStringDatasource(attrs.ACCOUNTSIMPORT),
		EnableAccountJSON:        util.SafeStringDatasource(attrs.ENABLEACCOUNTJSON),
		PasswordMaxLength:        util.SafeStringDatasource(attrs.PASSWORD_MAX_LENGTH),
		MaxPaginationSize:        util.SafeStringDatasource(attrs.MAX_PAGINATION_SIZE),
		UpdateAccountJSON:        util.SafeStringDatasource(attrs.UPDATEACCOUNTJSON),
		GrantAccessJSON:          util.SafeStringDatasource(attrs.GRANTACCESSJSON),
		CliCommandJSON:           util.SafeStringDatasource(attrs.CLI_COMMAND_JSON),
		// TER-176 entitlement management fields
		CreateEntitlementJson: util.SafeStringDatasource(attrs.CREATEENTITLEMENTJSON),
		DeleteEntitlementJson: util.SafeStringDatasource(attrs.DELETEENTITLEMENTJSON),
		EntitlementExistJson:  util.SafeStringDatasource(attrs.ENTITLEMENTEXISTJSON),
		UpdateEntitlementJson: util.SafeStringDatasource(attrs.UPDATEENTITLEMENTJSON),
	}

	// Map connection timeout config if present
	d.MapDBTimeoutConfig(attrs, state.ConnectionAttributes)
}

// MapDBTimeoutConfig maps connection timeout configuration from API response to state model
// Only maps timeout config if it exists in the API response
func (d *DbConnectionsDataSource) MapDBTimeoutConfig(attrs *openapi.DBConnectionAttributes, connectionAttrs *DBConnectionAttributes) {
	if attrs.ConnectionTimeoutConfig != nil {
		connectionAttrs.ConnectionTimeoutConfig = &ConnectionTimeoutConfig{
			RetryWait:               util.SafeInt64(attrs.ConnectionTimeoutConfig.RetryWait),
			TokenRefreshMaxTryCount: util.SafeInt64(attrs.ConnectionTimeoutConfig.TokenRefreshMaxTryCount),
			RetryWaitMaxValue:       util.SafeInt64(attrs.ConnectionTimeoutConfig.RetryWaitMaxValue),
			RetryCount:              util.SafeInt64(attrs.ConnectionTimeoutConfig.RetryCount),
			ReadTimeout:             util.SafeInt64(attrs.ConnectionTimeoutConfig.ReadTimeout),
			ConnectionTimeout:       util.SafeInt64(attrs.ConnectionTimeoutConfig.ConnectionTimeout),
			RetryFailureStatusCode:  util.SafeInt64(attrs.ConnectionTimeoutConfig.RetryFailureStatusCode),
		}
	}
}

// HandleDBAuthenticationLogic processes the authenticate flag to control sensitive data visibility
// When authenticate=false, connection_attributes are removed from state to prevent sensitive data exposure
// When authenticate=true, all connection_attributes are returned in state
func (d *DbConnectionsDataSource) HandleDBAuthenticationLogic(state *DBConnectionDataSourceModel, resp *datasource.ReadResponse) {
	if !state.Authenticate.IsNull() && !state.Authenticate.IsUnknown() {
		if state.Authenticate.ValueBool() {
			tflog.Info(context.Background(), "Authentication enabled - returning all connection attributes")
			resp.Diagnostics.AddWarning(
				"Authentication Enabled",
				"`authenticate` is true; all connection_attributes will be returned in state.",
			)
		} else {
			tflog.Info(context.Background(), "Authentication disabled - removing connection attributes from state")
			resp.Diagnostics.AddWarning(
				"Authentication Disabled",
				"`authenticate` is false; connection_attributes will be removed from state.",
			)
			state.ConnectionAttributes = nil
		}
	}
}
