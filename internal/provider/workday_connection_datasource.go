// Copyright (c) 2025 Saviynt Inc.
// SPDX-License-Identifier: MPL-2.0

// saviynt_workday_connection_datasource retrieves workday connections details from the Saviynt Security Manager.
// The data source supports a single Read operation to look up an existing workday connections by name.
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

var _ datasource.DataSource = &WorkdayConnectionDataSource{}

// Initialize error codes for Workday Connection datasource operations
var workdayDatasourceErrorCodes = errorsutil.NewConnectorErrorCodes(errorsutil.ConnectorTypeWorkday)

// WorkdayConnectionDataSource defines the data source
type WorkdayConnectionDataSource struct {
	client            client.SaviyntClientInterface
	token             string
	provider          client.SaviyntProviderInterface
	connectionFactory client.ConnectionFactoryInterface
}

type WorkdayConnectionDataSourceModel struct {
	ID types.String `tfsdk:"id"`
	BaseConnectionDataSourceModel
	ConnectionAttributes *WorkdayConnectionAttributes `tfsdk:"connection_attributes"`
}

type WorkdayConnectionAttributes struct {
	UseOauth                    types.String             `tfsdk:"use_oauth"`
	UserImportMapping           types.String             `tfsdk:"user_import_mapping"`
	AccountsLastImportTime      types.String             `tfsdk:"accounts_last_import_time"`
	StatusKeyJson               types.String             `tfsdk:"status_key_json"`
	ConnectionTimeoutConfig     *ConnectionTimeoutConfig `tfsdk:"connection_timeout_config"`
	ConnectionType              types.String             `tfsdk:"connection_type"`
	RaasMappingJson             types.String             `tfsdk:"raas_mapping_json"`
	AccountImportPayload        types.String             `tfsdk:"account_import_payload"`
	UpdateAccountPayload        types.String             `tfsdk:"update_account_payload"`
	StatusThresholdConfig       types.String             `tfsdk:"status_threshold_config"`
	AccessImportList            types.String             `tfsdk:"access_import_list"`
	IsTimeoutSupported          types.Bool               `tfsdk:"is_timeout_supported"`
	AccountImportMapping        types.String             `tfsdk:"account_import_mapping"`
	AssignOrgrolePayload        types.String             `tfsdk:"assign_orgrole_payload"`
	AccessImportMapping         types.String             `tfsdk:"access_import_mapping"`
	ApiVersion                  types.String             `tfsdk:"api_version"`
	RemoveOrgrolePayload        types.String             `tfsdk:"remove_orgrole_payload"`
	IncludeReferenceDescriptors types.String             `tfsdk:"include_reference_descriptors"`
	ModifyUserDataJson          types.String             `tfsdk:"modifyuserdatajson"`
	IsTimeoutConfigValidated    types.Bool               `tfsdk:"is_timeout_config_validated"`
	UseX509AuthForSoap          types.String             `tfsdk:"use_x509auth_for_soap"`
	ReportOwner                 types.String             `tfsdk:"report_owner"`
	X509Key                     types.String             `tfsdk:"x509_key"`
	CustomConfig                types.String             `tfsdk:"custom_config"`
	UserAttributeJson           types.String             `tfsdk:"userattributejson"`
	X509Cert                    types.String             `tfsdk:"x509_cert"`
	UserImportPayload           types.String             `tfsdk:"user_import_payload"`
	PamConfig                   types.String             `tfsdk:"pam_config"`
	AccessLastImportTime        types.String             `tfsdk:"access_last_import_time"`
	UsersLastImportTime         types.String             `tfsdk:"users_last_import_time"`
	UpdateUserPayload           types.String             `tfsdk:"update_user_payload"`
	PageSize                    types.String             `tfsdk:"page_size"`
	TenantName                  types.String             `tfsdk:"tenant_name"`
	UseEnhancedOrgrole          types.String             `tfsdk:"use_enhanced_orgrole"`
	CreateAccountPayload        types.String             `tfsdk:"create_account_payload"`
	BaseUrl                     types.String             `tfsdk:"base_url"`
}

// NewWorkdayConnectionsDataSource creates a new Workday connections data source with default factory
func NewWorkdayConnectionsDataSource() datasource.DataSource {
	return &WorkdayConnectionDataSource{
		connectionFactory: &client.DefaultConnectionFactory{},
	}
}

// NewWorkdayConnectionsDataSourceWithFactory creates a new Workday connections data source with custom factory
// Used primarily for testing with mock factories
func NewWorkdayConnectionsDataSourceWithFactory(factory client.ConnectionFactoryInterface) datasource.DataSource {
	return &WorkdayConnectionDataSource{
		connectionFactory: factory,
	}
}

// SetClient sets the client for testing purposes
func (d *WorkdayConnectionDataSource) SetClient(client client.SaviyntClientInterface) {
	d.client = client
}

// SetToken sets the token for testing purposes
func (d *WorkdayConnectionDataSource) SetToken(token string) {
	d.token = token
}

// SetProvider sets the provider for testing purposes
func (d *WorkdayConnectionDataSource) SetProvider(provider client.SaviyntProviderInterface) {
	d.provider = provider
}

// Metadata sets the data source type name for Terraform
func (d *WorkdayConnectionDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "saviynt_workday_connection_datasource"
}

func WorkdayConnectorsDataSourceSchema() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:    true,
			Description: "Resource ID.",
		},
		"connection_attributes": schema.SingleNestedAttribute{
			Computed: true,
			Attributes: map[string]schema.Attribute{
				"use_oauth":                     schema.StringAttribute{Computed: true},
				"user_import_mapping":           schema.StringAttribute{Computed: true},
				"accounts_last_import_time":     schema.StringAttribute{Computed: true},
				"status_key_json":               schema.StringAttribute{Computed: true},
				"connection_type":               schema.StringAttribute{Computed: true},
				"raas_mapping_json":             schema.StringAttribute{Computed: true},
				"account_import_payload":        schema.StringAttribute{Computed: true},
				"update_account_payload":        schema.StringAttribute{Computed: true},
				"status_threshold_config":       schema.StringAttribute{Computed: true},
				"access_import_list":            schema.StringAttribute{Computed: true},
				"is_timeout_supported":          schema.BoolAttribute{Computed: true},
				"account_import_mapping":        schema.StringAttribute{Computed: true},
				"assign_orgrole_payload":        schema.StringAttribute{Computed: true},
				"access_import_mapping":         schema.StringAttribute{Computed: true},
				"api_version":                   schema.StringAttribute{Computed: true},
				"remove_orgrole_payload":        schema.StringAttribute{Computed: true},
				"include_reference_descriptors": schema.StringAttribute{Computed: true},
				"modifyuserdatajson":            schema.StringAttribute{Computed: true},
				"is_timeout_config_validated":   schema.BoolAttribute{Computed: true},
				"use_x509auth_for_soap":         schema.StringAttribute{Computed: true},
				"report_owner":                  schema.StringAttribute{Computed: true},
				"x509_key":                      schema.StringAttribute{Computed: true},
				"custom_config":                 schema.StringAttribute{Computed: true},
				"userattributejson":             schema.StringAttribute{Computed: true},
				"x509_cert":                     schema.StringAttribute{Computed: true},
				"user_import_payload":           schema.StringAttribute{Computed: true},
				"pam_config":                    schema.StringAttribute{Computed: true},
				"access_last_import_time":       schema.StringAttribute{Computed: true},
				"users_last_import_time":        schema.StringAttribute{Computed: true},
				"update_user_payload":           schema.StringAttribute{Computed: true},
				"page_size":                     schema.StringAttribute{Computed: true},
				"tenant_name":                   schema.StringAttribute{Computed: true},
				"use_enhanced_orgrole":          schema.StringAttribute{Computed: true},
				"create_account_payload":        schema.StringAttribute{Computed: true},
				"base_url":                      schema.StringAttribute{Computed: true},
				"connection_timeout_config": schema.SingleNestedAttribute{
					Computed:   true,
					Attributes: ConnectionTimeoutConfigeSchema(),
				},
			},
		},
	}
}

// Schema defines the structure and attributes available for the Workday connection data source
func (d *WorkdayConnectionDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: util.WorkdayConnDataSourceDescription,
		Attributes:  connectionsutil.MergeDataSourceAttributes(BaseConnectorDataSourceSchema(), WorkdayConnectorsDataSourceSchema()),
	}
}

// Configure initializes the data source with provider configuration
// Sets up client and authentication token for API operations
func (d *WorkdayConnectionDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	opCtx := errorsutil.CreateOperationContext(errorsutil.ConnectorTypeWorkday, "configure", "")
	ctx = opCtx.AddContextToLogger(ctx)

	opCtx.LogOperationStart(ctx, "Starting Workday connection datasource configuration")

	// Check if provider data is available.
	if req.ProviderData == nil {
		tflog.Debug(ctx, "ProviderData is nil, returning early")
		opCtx.LogOperationEnd(ctx, "Workday connection datasource configuration completed - no provider data")
		return
	}

	// Cast provider data to your provider type.
	prov, ok := req.ProviderData.(*SaviyntProvider)
	if !ok {
		errorCode := workdayDatasourceErrorCodes.ProviderConfig()
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

	opCtx.LogOperationEnd(ctx, "Workday connection datasource configured successfully")
}

// Read retrieves Workday connection details from Saviynt and populates the Terraform state
// Supports lookup by connection name or connection key with comprehensive error handling
func (d *WorkdayConnectionDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var state WorkdayConnectionDataSourceModel

	opCtx := errorsutil.CreateOperationContext(errorsutil.ConnectorTypeWorkday, "datasource_read", "")
	ctx = opCtx.AddContextToLogger(ctx)

	opCtx.LogOperationStart(ctx, "Starting Workday connection datasource read")

	// Extract configuration from request
	configDiagnostics := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(configDiagnostics...)
	if resp.Diagnostics.HasError() {
		errorCode := workdayDatasourceErrorCodes.ConfigExtraction()
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
		errorCode := workdayDatasourceErrorCodes.MissingIdentifier()
		opCtx.LogOperationError(ctx, "Missing connection identifier", errorCode,
			fmt.Errorf("either connection_name or connection_key must be provided"))
		resp.Diagnostics.AddError(
			errorsutil.GetErrorMessage(errorsutil.ErrMissingIdentifier),
			fmt.Sprintf("[%s] Either 'connection_name' or 'connection_key' must be provided to look up the Workday connection.", errorCode),
		)
		return
	}

	// Execute API call to get Workday connection details
	apiResp, err := d.ReadWorkdayConnectionDetails(ctx, connectionName, connectionKey)
	if err != nil {
		// Error is already sanitized in ReadWorkdayConnectionDetails method
		errorCode := workdayDatasourceErrorCodes.ReadFailed()
		opCtx.LogOperationError(ctx, "Failed to read Workday connection details", errorCode, err)
		resp.Diagnostics.AddError(
			errorsutil.GetErrorMessage(errorsutil.ErrReadFailed),
			fmt.Sprintf("[%s] %s", errorCode, err.Error()),
		)
		return
	}

	// Validate API response
	if err := d.ValidateWorkdayConnectionResponse(apiResp); err != nil {
		errorCode := workdayDatasourceErrorCodes.APIError()
		opCtx.LogOperationError(ctx, "Invalid connection type for Workday datasource", errorCode, err)
		resp.Diagnostics.AddError(
			errorsutil.GetErrorMessage(errorsutil.ErrAPIError),
			fmt.Sprintf("[%s] Unable to verify connection type for connection %q. The provider could not determine the type of this connection. Please ensure the connection name is correct and belongs to a supported connector type.", errorCode, connectionName),
		)
		return
	}

	// Map API response to state
	d.UpdateModelFromWorkdayConnectionResponse(&state, apiResp)

	// Handle authentication logic
	d.HandleWorkdayAuthenticationLogic(&state, resp)

	// Set final state
	stateDiagnostics := resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(stateDiagnostics...)
	if resp.Diagnostics.HasError() {
		errorCode := workdayDatasourceErrorCodes.StateUpdate()
		opCtx.LogOperationError(ctx, "Failed to set state", errorCode,
			fmt.Errorf("state update failed"))
		resp.Diagnostics.AddError(
			errorsutil.GetErrorMessage(errorsutil.ErrStateUpdate),
			fmt.Sprintf("[%s] Unable to update Terraform state for Workday connection datasource", errorCode),
		)
		return
	}

	opCtx.LogOperationEnd(ctx, "Workday connection datasource read completed successfully",
		map[string]interface{}{
			"connection_name": connectionName,
			"has_attributes":  state.ConnectionAttributes != nil,
		})
}

// ReadWorkdayConnectionDetails retrieves Workday connection details from Saviynt API
// Handles both connection name and connection key based lookups using factory pattern
// Returns standardized errors with proper correlation tracking and sensitive data sanitization
func (d *WorkdayConnectionDataSource) ReadWorkdayConnectionDetails(ctx context.Context, connectionName string, connectionKey *int64) (*openapi.GetConnectionDetailsResponse, error) {
	opCtx := errorsutil.CreateOperationContext(errorsutil.ConnectorTypeWorkday, "api_read", connectionName)
	logCtx := opCtx.AddContextToLogger(ctx)

	opCtx.LogOperationStart(logCtx, "Starting Workday connection API call")

	tflog.Debug(logCtx, "Executing API request to get Workday connection details")

	var apiResp *openapi.GetConnectionDetailsResponse

	// Execute API request with retry logic
	err := d.provider.AuthenticatedAPICallWithRetry(ctx, "read_workday_connection_datasource", func(token string) error {
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
		errorCode := workdayDatasourceErrorCodes.ReadFailed()
		opCtx.LogOperationError(logCtx, "Failed to read Workday connection details", errorCode, err)
		return nil, errorsutil.CreateStandardError(errorsutil.ConnectorTypeWorkday, errorCode, "api_read", connectionName, err)
	}

	opCtx.LogOperationEnd(logCtx, "Workday connection API call completed successfully")

	return apiResp, nil
}

// ValidateWorkdayConnectionResponse validates that the API response contains valid Workday connection data
// Returns standardized error if validation fails
func (d *WorkdayConnectionDataSource) ValidateWorkdayConnectionResponse(apiResp *openapi.GetConnectionDetailsResponse) error {
	if apiResp != nil && apiResp.WorkdayConnectionResponse == nil {
		return fmt.Errorf("verify the connection type - Workday connection response is nil")
	}
	return nil
}

// UpdateModelFromWorkdayConnectionResponse maps API response data to the Terraform state model
// It handles both base connection fields and detailed connection attributes
func (d *WorkdayConnectionDataSource) UpdateModelFromWorkdayConnectionResponse(state *WorkdayConnectionDataSourceModel, apiResp *openapi.GetConnectionDetailsResponse) {
	// Map base connection fields
	d.MapBaseWorkdayConnectionFields(state, apiResp)

	// Map connection attributes
	d.MapWorkdayConnectionAttributes(state, apiResp)
}

// MapBaseWorkdayConnectionFields maps basic connection fields from API response to state model
// These are common fields available for all connection types
func (d *WorkdayConnectionDataSource) MapBaseWorkdayConnectionFields(state *WorkdayConnectionDataSourceModel, apiResp *openapi.GetConnectionDetailsResponse) {
	state.Msg = util.SafeStringDatasource(apiResp.WorkdayConnectionResponse.Msg)
	state.ErrorCode = util.SafeInt64(apiResp.WorkdayConnectionResponse.Errorcode)
	connectionKey := util.SafeInt64(apiResp.WorkdayConnectionResponse.Connectionkey)
	state.ID = types.StringValue(fmt.Sprintf("ds-workday-%d", connectionKey.ValueInt64()))
	state.ConnectionName = util.SafeStringDatasource(apiResp.WorkdayConnectionResponse.Connectionname)
	state.ConnectionKey = util.SafeInt64(apiResp.WorkdayConnectionResponse.Connectionkey)
	state.Description = util.SafeStringDatasource(apiResp.WorkdayConnectionResponse.Description)
	state.DefaultSavRoles = util.SafeStringDatasource(apiResp.WorkdayConnectionResponse.Defaultsavroles)
	state.ConnectionType = util.SafeStringDatasource(apiResp.WorkdayConnectionResponse.Connectiontype)
	state.CreatedOn = util.SafeStringDatasource(apiResp.WorkdayConnectionResponse.Createdon)
	state.CreatedBy = util.SafeStringDatasource(apiResp.WorkdayConnectionResponse.Createdby)
	state.UpdatedBy = util.SafeStringDatasource(apiResp.WorkdayConnectionResponse.Updatedby)
	state.EmailTemplate = util.SafeStringDatasource(apiResp.WorkdayConnectionResponse.Emailtemplate)
}

// MapWorkdayConnectionAttributes maps detailed Workday connection attributes from API response to state model
// Returns nil if no connection attributes are present in the response
func (d *WorkdayConnectionDataSource) MapWorkdayConnectionAttributes(state *WorkdayConnectionDataSourceModel, apiResp *openapi.GetConnectionDetailsResponse) {
	if apiResp.WorkdayConnectionResponse.Connectionattributes == nil {
		state.ConnectionAttributes = nil
		return
	}

	attrs := apiResp.WorkdayConnectionResponse.Connectionattributes
	state.ConnectionAttributes = &WorkdayConnectionAttributes{
		UseOauth:                    util.SafeStringDatasource(attrs.USE_OAUTH),
		UserImportMapping:           util.SafeStringDatasource(attrs.USER_IMPORT_MAPPING),
		AccountsLastImportTime:      util.SafeStringDatasource(attrs.ACCOUNTS_LAST_IMPORT_TIME),
		StatusKeyJson:               util.SafeStringDatasource(attrs.STATUS_KEY_JSON),
		ConnectionType:              util.SafeStringDatasource(attrs.ConnectionType),
		RaasMappingJson:             util.SafeStringDatasource(attrs.RAAS_MAPPING_JSON),
		AccountImportPayload:        util.SafeStringDatasource(attrs.ACCOUNT_IMPORT_PAYLOAD),
		UpdateAccountPayload:        util.SafeStringDatasource(attrs.UPDATE_ACCOUNT_PAYLOAD),
		StatusThresholdConfig:       util.SafeStringDatasource(attrs.STATUS_THRESHOLD_CONFIG),
		AccessImportList:            util.SafeStringDatasource(attrs.ACCESS_IMPORT_LIST),
		IsTimeoutSupported:          util.SafeBoolDatasource(attrs.IsTimeoutSupported),
		AccountImportMapping:        util.SafeStringDatasource(attrs.ACCOUNT_IMPORT_MAPPING),
		AssignOrgrolePayload:        util.SafeStringDatasource(attrs.ASSIGN_ORGROLE_PAYLOAD),
		AccessImportMapping:         util.SafeStringDatasource(attrs.ACCESS_IMPORT_MAPPING),
		ApiVersion:                  util.SafeStringDatasource(attrs.API_VERSION),
		RemoveOrgrolePayload:        util.SafeStringDatasource(attrs.REMOVE_ORGROLE_PAYLOAD),
		IncludeReferenceDescriptors: util.SafeStringDatasource(attrs.INCLUDE_REFERENCE_DESCRIPTORS),
		ModifyUserDataJson:          util.SafeStringDatasource(attrs.MODIFYUSERDATAJSON),
		IsTimeoutConfigValidated:    util.SafeBoolDatasource(attrs.IsTimeoutConfigValidated),
		UseX509AuthForSoap:          util.SafeStringDatasource(attrs.USEX509AUTHFORSOAP),
		ReportOwner:                 util.SafeStringDatasource(attrs.REPORT_OWNER),
		X509Key:                     util.SafeStringDatasource(attrs.X509KEY),
		CustomConfig:                util.SafeStringDatasource(attrs.CUSTOM_CONFIG),
		UserAttributeJson:           util.SafeStringDatasource(attrs.USERATTRIBUTEJSON),
		X509Cert:                    util.SafeStringDatasource(attrs.X509CERT),
		UserImportPayload:           util.SafeStringDatasource(attrs.USER_IMPORT_PAYLOAD),
		PamConfig:                   util.SafeStringDatasource(attrs.PAM_CONFIG),
		AccessLastImportTime:        util.SafeStringDatasource(attrs.ACCESS_LAST_IMPORT_TIME),
		UsersLastImportTime:         util.SafeStringDatasource(attrs.USERS_LAST_IMPORT_TIME),
		UpdateUserPayload:           util.SafeStringDatasource(attrs.UPDATE_USER_PAYLOAD),
		PageSize:                    util.SafeStringDatasource(attrs.PAGE_SIZE),
		TenantName:                  util.SafeStringDatasource(attrs.TENANT_NAME),
		UseEnhancedOrgrole:          util.SafeStringDatasource(attrs.USE_ENHANCED_ORGROLE),
		CreateAccountPayload:        util.SafeStringDatasource(attrs.CREATE_ACCOUNT_PAYLOAD),
		BaseUrl:                     util.SafeStringDatasource(attrs.BASE_URL),
	}

	// Map connection timeout config if present
	d.MapWorkdayTimeoutConfig(attrs, state.ConnectionAttributes)
}

// MapWorkdayTimeoutConfig maps connection timeout configuration from API response to state model
// Only maps timeout config if it exists in the API response
func (d *WorkdayConnectionDataSource) MapWorkdayTimeoutConfig(attrs *openapi.WorkdayConnectionAttributes, connectionAttrs *WorkdayConnectionAttributes) {
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

// HandleWorkdayAuthenticationLogic processes the authenticate flag to control sensitive data visibility
// When authenticate=false, connection_attributes are removed from state to prevent sensitive data exposure
// When authenticate=true, all connection_attributes are returned in state
func (d *WorkdayConnectionDataSource) HandleWorkdayAuthenticationLogic(state *WorkdayConnectionDataSourceModel, resp *datasource.ReadResponse) {
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
