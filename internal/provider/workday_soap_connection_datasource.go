// Copyright (c) 2025 Saviynt Inc.
// SPDX-License-Identifier: MPL-2.0

// saviynt_workday_soap_connection_datasource retrieves workday-soap connections details from the Saviynt Security Manager.
// The data source supports a single Read operation to look up an existing workday-soap connections by name.
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

var _ datasource.DataSource = &WorkdaySOAPConnectionDataSource{}

// Initialize error codes for Workday SOAP Connection datasource operations
var workdaySoapDatasourceErrorCodes = errorsutil.NewConnectorErrorCodes(errorsutil.ConnectorTypeWorkdaySOAP)

// WorkdaySOAPConnectionDataSource defines the data source
type WorkdaySOAPConnectionDataSource struct {
	client            client.SaviyntClientInterface
	token             string
	provider          client.SaviyntProviderInterface
	connectionFactory client.ConnectionFactoryInterface
}

type WorkdaySOAPConnectionDataSourceModel struct {
	ID types.String `tfsdk:"id"`
	BaseConnectionDataSourceModel
	ConnectionAttributes *WorkdaySOAPConnectionAttributes `tfsdk:"connection_attributes"`
}

type WorkdaySOAPConnectionAttributes struct {
	ResponsePathPageResults  types.String             `tfsdk:"responsepath_pageresults"`
	ChangePassJson           types.String             `tfsdk:"change_pass_json"`
	PasswordMinLength        types.String             `tfsdk:"password_min_length"`
	ConnectionJson           types.String             `tfsdk:"connection_json"`
	CreateAccountJson        types.String             `tfsdk:"create_account_json"`
	AccountsImportJson       types.String             `tfsdk:"accounts_import_json"`
	DisableAccountJson       types.String             `tfsdk:"disable_account_json"`
	ConnectionTimeoutConfig  *ConnectionTimeoutConfig `tfsdk:"connection_timeout_config"`
	UpdateUserJson           types.String             `tfsdk:"update_user_json"`
	RevokeAccessJson         types.String             `tfsdk:"revoke_access_json"`
	PasswordNoOfSplChars     types.String             `tfsdk:"password_noofsplchars"`
	ConnectionType           types.String             `tfsdk:"connection_type"`
	ResponsePathUserList     types.String             `tfsdk:"responsepath_userlist"`
	DeleteAccountJson        types.String             `tfsdk:"delete_account_json"`
	Username                 types.String             `tfsdk:"username"`
	DateFormat               types.String             `tfsdk:"date_format"`
	IsTimeoutSupported       types.Bool               `tfsdk:"is_timeout_supported"`
	HrImportJson             types.String             `tfsdk:"hr_import_json"`
	PasswordNoOfCapsAlpha    types.String             `tfsdk:"password_noofcapsalpha"`
	PasswordNoOfDigits       types.String             `tfsdk:"password_noofdigits"`
	SoapEndpoint             types.String             `tfsdk:"soap_endpoint"`
	ModifyUserDataJson       types.String             `tfsdk:"modify_user_data_json"`
	IsTimeoutConfigValidated types.Bool               `tfsdk:"is_timeout_config_validated"`
	CustomConfig             types.String             `tfsdk:"custom_config"`
	ResponsePathTotalResults types.String             `tfsdk:"responsepath_totalresults"`
	CombinedCreateRequest    types.String             `tfsdk:"combined_create_request"`
	Password                 types.String             `tfsdk:"password"`
	DataToImport             types.String             `tfsdk:"data_to_import"`
	PasswordType             types.String             `tfsdk:"password_type"`
	PamConfig                types.String             `tfsdk:"pam_config"`
	EnableAccountJson        types.String             `tfsdk:"enable_account_json"`
	PageSize                 types.String             `tfsdk:"page_size"`
	PasswordMaxLength        types.String             `tfsdk:"password_max_length"`
	UpdateAccountJson        types.String             `tfsdk:"update_account_json"`
	GrantAccessJson          types.String             `tfsdk:"grant_access_json"`
}

// NewWorkdaySOAPConnectionsDataSource creates a new Workday SOAP connections data source with default factory
func NewWorkdaySOAPConnectionsDataSource() datasource.DataSource {
	return &WorkdaySOAPConnectionDataSource{
		connectionFactory: &client.DefaultConnectionFactory{},
	}
}

// NewWorkdaySOAPConnectionsDataSourceWithFactory creates a new Workday SOAP connections data source with custom factory
// Used primarily for testing with mock factories
func NewWorkdaySOAPConnectionsDataSourceWithFactory(factory client.ConnectionFactoryInterface) datasource.DataSource {
	return &WorkdaySOAPConnectionDataSource{
		connectionFactory: factory,
	}
}

// SetClient sets the client for testing purposes
func (d *WorkdaySOAPConnectionDataSource) SetClient(client client.SaviyntClientInterface) {
	d.client = client
}

// SetToken sets the token for testing purposes
func (d *WorkdaySOAPConnectionDataSource) SetToken(token string) {
	d.token = token
}

// SetProvider sets the provider for testing purposes
func (d *WorkdaySOAPConnectionDataSource) SetProvider(provider client.SaviyntProviderInterface) {
	d.provider = provider
}

// Metadata sets the data source type name for Terraform
func (d *WorkdaySOAPConnectionDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "saviynt_workday_soap_connection_datasource"
}

func WorkdaySOAPConnectorsDataSourceSchema() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:    true,
			Description: "Resource ID.",
		},
		"connection_attributes": schema.SingleNestedAttribute{
			Computed: true,
			Attributes: map[string]schema.Attribute{
				"responsepath_pageresults":    schema.StringAttribute{Computed: true},
				"change_pass_json":            schema.StringAttribute{Computed: true},
				"password_min_length":         schema.StringAttribute{Computed: true},
				"connection_json":             schema.StringAttribute{Computed: true},
				"create_account_json":         schema.StringAttribute{Computed: true},
				"accounts_import_json":        schema.StringAttribute{Computed: true},
				"disable_account_json":        schema.StringAttribute{Computed: true},
				"update_user_json":            schema.StringAttribute{Computed: true},
				"revoke_access_json":          schema.StringAttribute{Computed: true},
				"password_noofsplchars":       schema.StringAttribute{Computed: true},
				"connection_type":             schema.StringAttribute{Computed: true},
				"responsepath_userlist":       schema.StringAttribute{Computed: true},
				"delete_account_json":         schema.StringAttribute{Computed: true},
				"username":                    schema.StringAttribute{Computed: true},
				"date_format":                 schema.StringAttribute{Computed: true},
				"is_timeout_supported":        schema.BoolAttribute{Computed: true},
				"hr_import_json":              schema.StringAttribute{Computed: true},
				"password_noofcapsalpha":      schema.StringAttribute{Computed: true},
				"password_noofdigits":         schema.StringAttribute{Computed: true},
				"soap_endpoint":               schema.StringAttribute{Computed: true},
				"modify_user_data_json":       schema.StringAttribute{Computed: true},
				"is_timeout_config_validated": schema.BoolAttribute{Computed: true},
				"custom_config":               schema.StringAttribute{Computed: true},
				"responsepath_totalresults":   schema.StringAttribute{Computed: true},
				"combined_create_request":     schema.StringAttribute{Computed: true},
				"password":                    schema.StringAttribute{Computed: true},
				"data_to_import":              schema.StringAttribute{Computed: true},
				"password_type":               schema.StringAttribute{Computed: true},
				"pam_config":                  schema.StringAttribute{Computed: true},
				"enable_account_json":         schema.StringAttribute{Computed: true},
				"page_size":                   schema.StringAttribute{Computed: true},
				"password_max_length":         schema.StringAttribute{Computed: true},
				"update_account_json":         schema.StringAttribute{Computed: true},
				"grant_access_json":           schema.StringAttribute{Computed: true},
				"connection_timeout_config": schema.SingleNestedAttribute{
					Computed:   true,
					Attributes: ConnectionTimeoutConfigeSchema(),
				},
			},
		},
	}
}

// Schema defines the structure and attributes available for the Workday SOAP connection data source
func (d *WorkdaySOAPConnectionDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: util.WorkdaySOAPConnDataSourceDescription,
		Attributes:  connectionsutil.MergeDataSourceAttributes(BaseConnectorDataSourceSchema(), WorkdaySOAPConnectorsDataSourceSchema()),
	}
}

// Configure initializes the data source with provider configuration
// Sets up client and authentication token for API operations
func (d *WorkdaySOAPConnectionDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	opCtx := errorsutil.CreateOperationContext(errorsutil.ConnectorTypeWorkdaySOAP, "configure", "")
	ctx = opCtx.AddContextToLogger(ctx)

	opCtx.LogOperationStart(ctx, "Starting Workday SOAP connection datasource configuration")

	// Check if provider data is available.
	if req.ProviderData == nil {
		tflog.Debug(ctx, "ProviderData is nil, returning early")
		opCtx.LogOperationEnd(ctx, "Workday SOAP connection datasource configuration completed - no provider data")
		return
	}

	// Cast provider data to your provider type.
	prov, ok := req.ProviderData.(*SaviyntProvider)
	if !ok {
		errorCode := workdaySoapDatasourceErrorCodes.ProviderConfig()
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

	opCtx.LogOperationEnd(ctx, "Workday SOAP connection datasource configured successfully")
}

// Read retrieves Workday SOAP connection details from Saviynt and populates the Terraform state
// Supports lookup by connection name or connection key with comprehensive error handling
func (d *WorkdaySOAPConnectionDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var state WorkdaySOAPConnectionDataSourceModel

	opCtx := errorsutil.CreateOperationContext(errorsutil.ConnectorTypeWorkdaySOAP, "datasource_read", "")
	ctx = opCtx.AddContextToLogger(ctx)

	opCtx.LogOperationStart(ctx, "Starting Workday SOAP connection datasource read")

	// Extract configuration from request
	configDiagnostics := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(configDiagnostics...)
	if resp.Diagnostics.HasError() {
		errorCode := workdaySoapDatasourceErrorCodes.ConfigExtraction()
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
		errorCode := workdaySoapDatasourceErrorCodes.MissingIdentifier()
		opCtx.LogOperationError(ctx, "Missing connection identifier", errorCode,
			fmt.Errorf("either connection_name or connection_key must be provided"))
		resp.Diagnostics.AddError(
			errorsutil.GetErrorMessage(errorsutil.ErrMissingIdentifier),
			fmt.Sprintf("[%s] Either 'connection_name' or 'connection_key' must be provided to look up the Workday SOAP connection.", errorCode),
		)
		return
	}

	// Execute API call to get Workday SOAP connection details
	apiResp, err := d.ReadWorkdaySOAPConnectionDetails(ctx, connectionName, connectionKey)
	if err != nil {
		// Error is already sanitized in ReadWorkdaySOAPConnectionDetails method
		errorCode := workdaySoapDatasourceErrorCodes.ReadFailed()
		opCtx.LogOperationError(ctx, "Failed to read Workday SOAP connection details", errorCode, err)
		resp.Diagnostics.AddError(
			errorsutil.GetErrorMessage(errorsutil.ErrReadFailed),
			fmt.Sprintf("[%s] %s", errorCode, err.Error()),
		)
		return
	}

	// Validate API response
	if err := d.ValidateWorkdaySOAPConnectionResponse(apiResp); err != nil {
		errorCode := workdaySoapDatasourceErrorCodes.APIError()
		opCtx.LogOperationError(ctx, "Invalid connection type for Workday SOAP datasource", errorCode, err)
		resp.Diagnostics.AddError(
			errorsutil.GetErrorMessage(errorsutil.ErrAPIError),
			fmt.Sprintf("[%s] Unable to verify connection type for connection %q. The provider could not determine the type of this connection. Please ensure the connection name is correct and belongs to a supported connector type.", errorCode, connectionName),
		)
		return
	}

	// Map API response to state
	d.UpdateModelFromWorkdaySOAPConnectionResponse(&state, apiResp)

	// Handle authentication logic
	d.HandleWorkdaySOAPAuthenticationLogic(&state, resp)

	// Set final state
	stateDiagnostics := resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(stateDiagnostics...)
	if resp.Diagnostics.HasError() {
		errorCode := workdaySoapDatasourceErrorCodes.StateUpdate()
		opCtx.LogOperationError(ctx, "Failed to set state", errorCode,
			fmt.Errorf("state update failed"))
		resp.Diagnostics.AddError(
			errorsutil.GetErrorMessage(errorsutil.ErrStateUpdate),
			fmt.Sprintf("[%s] Unable to update Terraform state for Workday SOAP connection datasource", errorCode),
		)
		return
	}

	opCtx.LogOperationEnd(ctx, "Workday SOAP connection datasource read completed successfully",
		map[string]interface{}{
			"connection_name": connectionName,
			"has_attributes":  state.ConnectionAttributes != nil,
		})
}

// ReadWorkdaySOAPConnectionDetails retrieves Workday SOAP connection details from Saviynt API
// Handles both connection name and connection key based lookups using factory pattern
// Returns standardized errors with proper correlation tracking and sensitive data sanitization
func (d *WorkdaySOAPConnectionDataSource) ReadWorkdaySOAPConnectionDetails(ctx context.Context, connectionName string, connectionKey *int64) (*openapi.GetConnectionDetailsResponse, error) {
	opCtx := errorsutil.CreateOperationContext(errorsutil.ConnectorTypeWorkdaySOAP, "api_read", connectionName)
	logCtx := opCtx.AddContextToLogger(ctx)

	opCtx.LogOperationStart(logCtx, "Starting Workday SOAP connection API call")

	tflog.Debug(logCtx, "Executing API request to get Workday SOAP connection details")

	var apiResp *openapi.GetConnectionDetailsResponse

	// Execute API request with retry logic
	err := d.provider.AuthenticatedAPICallWithRetry(ctx, "read_workday_soap_connection_datasource", func(token string) error {
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
		errorCode := workdaySoapDatasourceErrorCodes.ReadFailed()
		opCtx.LogOperationError(logCtx, "Failed to read Workday SOAP connection details", errorCode, err)
		return nil, errorsutil.CreateStandardError(errorsutil.ConnectorTypeWorkdaySOAP, errorCode, "api_read", connectionName, err)
	}

	opCtx.LogOperationEnd(logCtx, "Workday SOAP connection API call completed successfully")

	return apiResp, nil
}

// ValidateWorkdaySOAPConnectionResponse validates that the API response contains valid Workday SOAP connection data
// Returns standardized error if validation fails
func (d *WorkdaySOAPConnectionDataSource) ValidateWorkdaySOAPConnectionResponse(apiResp *openapi.GetConnectionDetailsResponse) error {
	if apiResp != nil && apiResp.WorkdaySOAPConnectionResponse == nil {
		return fmt.Errorf("verify the connection type - Workday SOAP connection response is nil")
	}
	return nil
}

// UpdateModelFromWorkdaySOAPConnectionResponse maps API response data to the Terraform state model
// It handles both base connection fields and detailed connection attributes
func (d *WorkdaySOAPConnectionDataSource) UpdateModelFromWorkdaySOAPConnectionResponse(state *WorkdaySOAPConnectionDataSourceModel, apiResp *openapi.GetConnectionDetailsResponse) {
	// Map base connection fields
	d.MapBaseWorkdaySOAPConnectionFields(state, apiResp)

	// Map connection attributes
	d.MapWorkdaySOAPConnectionAttributes(state, apiResp)
}

// MapBaseWorkdaySOAPConnectionFields maps basic connection fields from API response to state model
// These are common fields available for all connection types
func (d *WorkdaySOAPConnectionDataSource) MapBaseWorkdaySOAPConnectionFields(state *WorkdaySOAPConnectionDataSourceModel, apiResp *openapi.GetConnectionDetailsResponse) {
	state.Msg = util.SafeStringDatasource(apiResp.WorkdaySOAPConnectionResponse.Msg)
	state.ErrorCode = util.SafeInt64(apiResp.WorkdaySOAPConnectionResponse.Errorcode)
	connectionKey := util.SafeInt64(apiResp.WorkdaySOAPConnectionResponse.Connectionkey)
	state.ID = types.StringValue(fmt.Sprintf("ds-workday-soap-%d", connectionKey.ValueInt64()))
	state.ConnectionName = util.SafeStringDatasource(apiResp.WorkdaySOAPConnectionResponse.Connectionname)
	state.ConnectionKey = util.SafeInt64(apiResp.WorkdaySOAPConnectionResponse.Connectionkey)
	state.Description = util.SafeStringDatasource(apiResp.WorkdaySOAPConnectionResponse.Description)
	state.DefaultSavRoles = util.SafeStringDatasource(apiResp.WorkdaySOAPConnectionResponse.Defaultsavroles)
	state.ConnectionType = util.SafeStringDatasource(apiResp.WorkdaySOAPConnectionResponse.Connectiontype)
	state.CreatedOn = util.SafeStringDatasource(apiResp.WorkdaySOAPConnectionResponse.Createdon)
	state.CreatedBy = util.SafeStringDatasource(apiResp.WorkdaySOAPConnectionResponse.Createdby)
	state.UpdatedBy = util.SafeStringDatasource(apiResp.WorkdaySOAPConnectionResponse.Updatedby)
	state.EmailTemplate = util.SafeStringDatasource(apiResp.WorkdaySOAPConnectionResponse.Emailtemplate)
}

// MapWorkdaySOAPConnectionAttributes maps detailed Workday SOAP connection attributes from API response to state model
// Returns nil if no connection attributes are present in the response
func (d *WorkdaySOAPConnectionDataSource) MapWorkdaySOAPConnectionAttributes(state *WorkdaySOAPConnectionDataSourceModel, apiResp *openapi.GetConnectionDetailsResponse) {
	if apiResp.WorkdaySOAPConnectionResponse.Connectionattributes == nil {
		state.ConnectionAttributes = nil
		return
	}

	attrs := apiResp.WorkdaySOAPConnectionResponse.Connectionattributes
	state.ConnectionAttributes = &WorkdaySOAPConnectionAttributes{
		ResponsePathPageResults:  util.SafeStringDatasource(attrs.RESPONSEPATH_PAGERESULTS),
		ChangePassJson:           util.SafeStringDatasource(attrs.CHANGEPASSJSON),
		PasswordMinLength:        util.SafeStringDatasource(attrs.PASSWORD_MIN_LENGTH),
		ConnectionJson:           util.SafeStringDatasource(attrs.CONNECTIONJSON),
		CreateAccountJson:        util.SafeStringDatasource(attrs.CREATEACCOUNTJSON),
		AccountsImportJson:       util.SafeStringDatasource(attrs.ACCOUNTS_IMPORT_JSON),
		DisableAccountJson:       util.SafeStringDatasource(attrs.DISABLEACCOUNTJSON),
		UpdateUserJson:           util.SafeStringDatasource(attrs.UPDATEUSERJSON),
		RevokeAccessJson:         util.SafeStringDatasource(attrs.REVOKEACCESSJSON),
		PasswordNoOfSplChars:     util.SafeStringDatasource(attrs.PASSWORD_NOOFSPLCHARS),
		ConnectionType:           util.SafeStringDatasource(attrs.ConnectionType),
		ResponsePathUserList:     util.SafeStringDatasource(attrs.RESPONSEPATH_USERLIST),
		DeleteAccountJson:        util.SafeStringDatasource(attrs.DELETEACCOUNTJSON),
		Username:                 util.SafeStringDatasource(attrs.USERNAME),
		DateFormat:               util.SafeStringDatasource(attrs.DATEFORMAT),
		IsTimeoutSupported:       util.SafeBoolDatasource(attrs.IsTimeoutSupported),
		HrImportJson:             util.SafeStringDatasource(attrs.HR_IMPORT_JSON),
		PasswordNoOfCapsAlpha:    util.SafeStringDatasource(attrs.PASSWORD_NOOFCAPSALPHA),
		PasswordNoOfDigits:       util.SafeStringDatasource(attrs.PASSWORD_NOOFDIGITS),
		SoapEndpoint:             util.SafeStringDatasource(attrs.SOAP_ENDPOINT),
		ModifyUserDataJson:       util.SafeStringDatasource(attrs.MODIFYUSERDATAJSON),
		IsTimeoutConfigValidated: util.SafeBoolDatasource(attrs.IsTimeoutConfigValidated),
		CustomConfig:             util.SafeStringDatasource(attrs.CUSTOM_CONFIG),
		ResponsePathTotalResults: util.SafeStringDatasource(attrs.RESPONSEPATH_TOTALRESULTS),
		CombinedCreateRequest:    util.SafeStringDatasource(attrs.COMBINEDCREATEREQUEST),
		Password:                 util.SafeStringDatasource(attrs.PASSWORD),
		DataToImport:             util.SafeStringDatasource(attrs.DATA_TO_IMPORT),
		PasswordType:             util.SafeStringDatasource(attrs.PASSWORD_TYPE),
		PamConfig:                util.SafeStringDatasource(attrs.PAM_CONFIG),
		EnableAccountJson:        util.SafeStringDatasource(attrs.ENABLEACCOUNTJSON),
		PageSize:                 util.SafeStringDatasource(attrs.PAGE_SIZE),
		PasswordMaxLength:        util.SafeStringDatasource(attrs.PASSWORD_MAX_LENGTH),
		UpdateAccountJson:        util.SafeStringDatasource(attrs.UPDATEACCOUNTJSON),
		GrantAccessJson:          util.SafeStringDatasource(attrs.GRANTACCESSJSON),
	}

	// Map connection timeout config if present
	d.MapWorkdaySOAPTimeoutConfig(attrs, state.ConnectionAttributes)
}

// MapWorkdaySOAPTimeoutConfig maps connection timeout configuration from API response to state model
// Only maps timeout config if it exists in the API response
func (d *WorkdaySOAPConnectionDataSource) MapWorkdaySOAPTimeoutConfig(attrs *openapi.WorkdaySOAPConnectionAttributes, connectionAttrs *WorkdaySOAPConnectionAttributes) {
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

// HandleWorkdaySOAPAuthenticationLogic processes the authenticate flag to control sensitive data visibility
// When authenticate=false, connection_attributes are removed from state to prevent sensitive data exposure
// When authenticate=true, all connection_attributes are returned in state
func (d *WorkdaySOAPConnectionDataSource) HandleWorkdaySOAPAuthenticationLogic(state *WorkdaySOAPConnectionDataSourceModel, resp *datasource.ReadResponse) {
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
