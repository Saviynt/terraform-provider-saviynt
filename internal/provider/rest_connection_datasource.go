// Copyright (c) 2025 Saviynt Inc.
// SPDX-License-Identifier: MPL-2.0

// saviynt_rest_connection_datasource retrieves rest connections details from the Saviynt Security Manager.
// The data source supports a single Read operation to look up an existing rest connections by name.
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

var _ datasource.DataSource = &restConnectionDatasource{}

// Initialize error codes for REST Connection datasource operations
var restDatasourceErrorCodes = errorsutil.NewConnectorErrorCodes(errorsutil.ConnectorTypeREST)

// RESTConnectionsDataSource defines the data source
type restConnectionDatasource struct {
	client            client.SaviyntClientInterface
	token             string
	provider          client.SaviyntProviderInterface
	connectionFactory client.ConnectionFactoryInterface
}

type RESTConnectionDataSourceModel struct {
	ID types.String `tfsdk:"id"`
	BaseConnectionDataSourceModel
	ConnectionAttributes *RESTConnectionAttributes `tfsdk:"connection_attributes"`
}

type RESTConnectionAttributes struct {
	UpdateUserJSON           types.String             `tfsdk:"update_user_json"`
	RemoveAccountJSON        types.String             `tfsdk:"remove_account_json"`
	TicketStatusJSON         types.String             `tfsdk:"ticket_status_json"`
	CreateTicketJSON         types.String             `tfsdk:"create_ticket_json"`
	ConnectionType           types.String             `tfsdk:"connection_type"`
	EndpointsFilter          types.String             `tfsdk:"endpoints_filter"`
	PasswdPolicyJSON         types.String             `tfsdk:"passwd_policy_json"`
	ConfigJSON               types.String             `tfsdk:"config_json"`
	AddFFIDAccessJSON        types.String             `tfsdk:"add_ffid_access_json"`
	RemoveFFIDAccessJSON     types.String             `tfsdk:"remove_ffid_access_json"`
	StatusThresholdConfig    types.String             `tfsdk:"status_threshold_config"`
	ModifyUserDataJSON       types.String             `tfsdk:"modify_user_data_json"`
	SendOtpJSON              types.String             `tfsdk:"send_otp_json"`
	ValidateOtpJSON          types.String             `tfsdk:"validate_otp_json"`
	PamConfig                types.String             `tfsdk:"pam_config"`
	ConnectionTimeoutConfig  *ConnectionTimeoutConfig `tfsdk:"connection_timeout_config"`
	CreateAccountJSON        types.String             `tfsdk:"create_account_json"`
	UpdateAccountJSON        types.String             `tfsdk:"update_account_json"`
	EnableAccountJSON        types.String             `tfsdk:"enable_account_json"`
	DisableAccountJSON       types.String             `tfsdk:"disable_account_json"`
	AddAccessJSON            types.String             `tfsdk:"add_access_json"`
	RemoveAccessJSON         types.String             `tfsdk:"remove_access_json"`
	ImportUserJSON           types.String             `tfsdk:"import_user_json"`
	IsTimeoutSupported       types.Bool               `tfsdk:"is_timeout_supported"`
	ImportAccountEntJSON     types.String             `tfsdk:"import_account_ent_json"`
	IsTimeoutConfigValidated types.Bool               `tfsdk:"is_timeout_config_validated"`
	//TER-176
	ApplicationDiscoveryJson types.String `tfsdk:"application_discovery_json"`
	CreateEntitlementJson    types.String `tfsdk:"create_entitlement_json"`
	DeleteEntitlementJson    types.String `tfsdk:"delete_entitlement_json"`
	UpdateEntitlementJson    types.String `tfsdk:"update_entitlement_json"`
}

// NewRESTConnectionsDataSource creates a new REST connections data source with default factory
func NewRESTConnectionsDataSource() datasource.DataSource {
	return &restConnectionDatasource{
		connectionFactory: &client.DefaultConnectionFactory{},
	}
}

// NewRESTConnectionsDataSourceWithFactory creates a new REST connections data source with custom factory
// Used primarily for testing with mock factories
func NewRESTConnectionsDataSourceWithFactory(factory client.ConnectionFactoryInterface) datasource.DataSource {
	return &restConnectionDatasource{
		connectionFactory: factory,
	}
}

// SetClient sets the client for testing purposes
func (d *restConnectionDatasource) SetClient(client client.SaviyntClientInterface) {
	d.client = client
}

// SetToken sets the token for testing purposes
func (d *restConnectionDatasource) SetToken(token string) {
	d.token = token
}

// SetProvider sets the provider for testing purposes
func (d *restConnectionDatasource) SetProvider(provider client.SaviyntProviderInterface) {
	d.provider = provider
}

// Metadata sets the data source type name for Terraform
func (d *restConnectionDatasource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "saviynt_rest_connection_datasource"
}

func RESTConnectorsDataSourceSchema() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:    true,
			Description: "Resource ID.",
		},
		"connection_attributes": schema.SingleNestedAttribute{
			Computed: true,
			Attributes: map[string]schema.Attribute{
				"update_user_json":            schema.StringAttribute{Computed: true},
				"remove_account_json":         schema.StringAttribute{Computed: true},
				"ticket_status_json":          schema.StringAttribute{Computed: true},
				"create_ticket_json":          schema.StringAttribute{Computed: true},
				"connection_type":             schema.StringAttribute{Computed: true},
				"endpoints_filter":            schema.StringAttribute{Computed: true},
				"passwd_policy_json":          schema.StringAttribute{Computed: true},
				"config_json":                 schema.StringAttribute{Computed: true},
				"add_ffid_access_json":        schema.StringAttribute{Computed: true},
				"remove_ffid_access_json":     schema.StringAttribute{Computed: true},
				"status_threshold_config":     schema.StringAttribute{Computed: true},
				"modify_user_data_json":       schema.StringAttribute{Computed: true},
				"send_otp_json":               schema.StringAttribute{Computed: true},
				"validate_otp_json":           schema.StringAttribute{Computed: true},
				"pam_config":                  schema.StringAttribute{Computed: true},
				"create_account_json":         schema.StringAttribute{Computed: true},
				"update_account_json":         schema.StringAttribute{Computed: true},
				"enable_account_json":         schema.StringAttribute{Computed: true},
				"disable_account_json":        schema.StringAttribute{Computed: true},
				"add_access_json":             schema.StringAttribute{Computed: true},
				"remove_access_json":          schema.StringAttribute{Computed: true},
				"import_user_json":            schema.StringAttribute{Computed: true},
				"is_timeout_supported":        schema.BoolAttribute{Computed: true},
				"import_account_ent_json":     schema.StringAttribute{Computed: true},
				"is_timeout_config_validated": schema.BoolAttribute{Computed: true},
				"connection_timeout_config": schema.SingleNestedAttribute{
					Computed:   true,
					Attributes: ConnectionTimeoutConfigeSchema(),
				},
				//TER-176
				"application_discovery_json": schema.StringAttribute{Computed: true},
				"create_entitlement_json":    schema.StringAttribute{Computed: true},
				"delete_entitlement_json":    schema.StringAttribute{Computed: true},
				"update_entitlement_json":    schema.StringAttribute{Computed: true},
			},
		},
	}
}

// Schema defines the structure and attributes available for the REST connection data source
func (d *restConnectionDatasource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: util.RestConnDataSourceDescription,
		Attributes:  connectionsutil.MergeDataSourceAttributes(BaseConnectorDataSourceSchema(), RESTConnectorsDataSourceSchema()),
	}
}

// Configure initializes the data source with provider configuration
// Sets up client and authentication token for API operations
func (d *restConnectionDatasource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	opCtx := errorsutil.CreateOperationContext(errorsutil.ConnectorTypeREST, "configure", "")
	ctx = opCtx.AddContextToLogger(ctx)

	opCtx.LogOperationStart(ctx, "Starting REST connection datasource configuration")

	// Check if provider data is available.
	if req.ProviderData == nil {
		tflog.Debug(ctx, "ProviderData is nil, returning early")
		opCtx.LogOperationEnd(ctx, "REST connection datasource configuration completed - no provider data")
		return
	}

	// Cast provider data to your provider type.
	prov, ok := req.ProviderData.(*SaviyntProvider)
	if !ok {
		errorCode := restDatasourceErrorCodes.ProviderConfig()
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

	opCtx.LogOperationEnd(ctx, "REST connection datasource configured successfully")
}

// Read retrieves REST connection details from Saviynt and populates the Terraform state
// Supports lookup by connection name or connection key with comprehensive error handling
func (d *restConnectionDatasource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var state RESTConnectionDataSourceModel

	opCtx := errorsutil.CreateOperationContext(errorsutil.ConnectorTypeREST, "datasource_read", "")
	ctx = opCtx.AddContextToLogger(ctx)

	opCtx.LogOperationStart(ctx, "Starting REST connection datasource read")

	// Extract configuration from request
	configDiagnostics := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(configDiagnostics...)
	if resp.Diagnostics.HasError() {
		errorCode := restDatasourceErrorCodes.ConfigExtraction()
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
		errorCode := restDatasourceErrorCodes.MissingIdentifier()
		opCtx.LogOperationError(ctx, "Missing connection identifier", errorCode,
			fmt.Errorf("either connection_name or connection_key must be provided"))
		resp.Diagnostics.AddError(
			errorsutil.GetErrorMessage(errorsutil.ErrMissingIdentifier),
			fmt.Sprintf("[%s] Either 'connection_name' or 'connection_key' must be provided to look up the REST connection.", errorCode),
		)
		return
	}

	// Execute API call to get REST connection details
	apiResp, err := d.ReadRESTConnectionDetails(ctx, connectionName, connectionKey)
	if err != nil {
		// Error is already sanitized in ReadRESTConnectionDetails method
		errorCode := restDatasourceErrorCodes.ReadFailed()
		opCtx.LogOperationError(ctx, "Failed to read REST connection details", errorCode, err)
		resp.Diagnostics.AddError(
			errorsutil.GetErrorMessage(errorsutil.ErrReadFailed),
			fmt.Sprintf("[%s] %s", errorCode, err.Error()),
		)
		return
	}

	// Validate API response
	if err := d.ValidateRESTConnectionResponse(apiResp); err != nil {
		errorCode := restDatasourceErrorCodes.APIError()
		opCtx.LogOperationError(ctx, "Invalid connection type for REST datasource", errorCode, err)
		resp.Diagnostics.AddError(
			errorsutil.GetErrorMessage(errorsutil.ErrAPIError),
			fmt.Sprintf("[%s] Unable to verify connection type for connection %q. The provider could not determine the type of this connection. Please ensure the connection name is correct and belongs to a supported connector type.", errorCode, connectionName),
		)
		return
	}

	// Map API response to state
	d.UpdateModelFromRESTConnectionResponse(&state, apiResp)

	// Handle authentication logic
	d.HandleRESTAuthenticationLogic(&state, resp)

	// Set final state
	stateDiagnostics := resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(stateDiagnostics...)
	if resp.Diagnostics.HasError() {
		errorCode := restDatasourceErrorCodes.StateUpdate()
		opCtx.LogOperationError(ctx, "Failed to set state", errorCode,
			fmt.Errorf("state update failed"))
		resp.Diagnostics.AddError(
			errorsutil.GetErrorMessage(errorsutil.ErrStateUpdate),
			fmt.Sprintf("[%s] Unable to update Terraform state for REST connection datasource", errorCode),
		)
		return
	}

	opCtx.LogOperationEnd(ctx, "REST connection datasource read completed successfully",
		map[string]interface{}{
			"connection_name": connectionName,
			"has_attributes":  state.ConnectionAttributes != nil,
		})
}

// ReadRESTConnectionDetails retrieves REST connection details from Saviynt API
// Handles both connection name and connection key based lookups using factory pattern
// Returns standardized errors with proper correlation tracking and sensitive data sanitization
func (d *restConnectionDatasource) ReadRESTConnectionDetails(ctx context.Context, connectionName string, connectionKey *int64) (*openapi.GetConnectionDetailsResponse, error) {
	opCtx := errorsutil.CreateOperationContext(errorsutil.ConnectorTypeREST, "api_read", connectionName)
	logCtx := opCtx.AddContextToLogger(ctx)

	opCtx.LogOperationStart(logCtx, "Starting REST connection API call")

	tflog.Debug(logCtx, "Executing API request to get REST connection details")

	var apiResp *openapi.GetConnectionDetailsResponse

	// Execute API request with retry logic
	err := d.provider.AuthenticatedAPICallWithRetry(ctx, "read_rest_connection_datasource", func(token string) error {
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
		errorCode := restDatasourceErrorCodes.ReadFailed()
		opCtx.LogOperationError(logCtx, "Failed to read REST connection details", errorCode, err)
		return nil, errorsutil.CreateStandardError(errorsutil.ConnectorTypeREST, errorCode, "api_read", connectionName, err)
	}

	opCtx.LogOperationEnd(logCtx, "REST connection API call completed successfully")

	return apiResp, nil
}

// ValidateRESTConnectionResponse validates that the API response contains valid REST connection data
// Returns standardized error if validation fails
func (d *restConnectionDatasource) ValidateRESTConnectionResponse(apiResp *openapi.GetConnectionDetailsResponse) error {
	if apiResp != nil && apiResp.RESTConnectionResponse == nil {
		return fmt.Errorf("verify the connection type - REST connection response is nil")
	}
	return nil
}

// UpdateModelFromRESTConnectionResponse maps API response data to the Terraform state model
// It handles both base connection fields and detailed connection attributes
func (d *restConnectionDatasource) UpdateModelFromRESTConnectionResponse(state *RESTConnectionDataSourceModel, apiResp *openapi.GetConnectionDetailsResponse) {
	// Map base connection fields
	d.MapBaseRESTConnectionFields(state, apiResp)

	// Map connection attributes
	d.MapRESTConnectionAttributes(state, apiResp)
}

// MapBaseRESTConnectionFields maps basic connection fields from API response to state model
// These are common fields available for all connection types
func (d *restConnectionDatasource) MapBaseRESTConnectionFields(state *RESTConnectionDataSourceModel, apiResp *openapi.GetConnectionDetailsResponse) {
	state.Msg = util.SafeStringDatasource(apiResp.RESTConnectionResponse.Msg)
	state.ErrorCode = util.SafeInt64(apiResp.RESTConnectionResponse.Errorcode)
	connectionKey := util.SafeInt64(apiResp.RESTConnectionResponse.Connectionkey)
	state.ID = types.StringValue(fmt.Sprintf("ds-rest-%d", connectionKey.ValueInt64()))
	state.ConnectionName = util.SafeStringDatasource(apiResp.RESTConnectionResponse.Connectionname)
	state.ConnectionKey = util.SafeInt64(apiResp.RESTConnectionResponse.Connectionkey)
	state.Description = util.SafeStringDatasource(apiResp.RESTConnectionResponse.Description)
	state.DefaultSavRoles = util.SafeStringDatasource(apiResp.RESTConnectionResponse.Defaultsavroles)
	state.ConnectionType = util.SafeStringDatasource(apiResp.RESTConnectionResponse.Connectiontype)
	state.CreatedOn = util.SafeStringDatasource(apiResp.RESTConnectionResponse.Createdon)
	state.CreatedBy = util.SafeStringDatasource(apiResp.RESTConnectionResponse.Createdby)
	state.UpdatedBy = util.SafeStringDatasource(apiResp.RESTConnectionResponse.Updatedby)
	state.EmailTemplate = util.SafeStringDatasource(apiResp.RESTConnectionResponse.Emailtemplate)
}

// MapRESTConnectionAttributes maps detailed REST connection attributes from API response to state model
// Returns nil if no connection attributes are present in the response
func (d *restConnectionDatasource) MapRESTConnectionAttributes(state *RESTConnectionDataSourceModel, apiResp *openapi.GetConnectionDetailsResponse) {
	if apiResp.RESTConnectionResponse.Connectionattributes == nil {
		state.ConnectionAttributes = nil
		return
	}

	attrs := apiResp.RESTConnectionResponse.Connectionattributes
	state.ConnectionAttributes = &RESTConnectionAttributes{
		UpdateUserJSON:           util.SafeStringDatasource(attrs.UpdateUserJSON),
		RemoveAccountJSON:        util.SafeStringDatasource(attrs.RemoveAccountJSON),
		TicketStatusJSON:         util.SafeStringDatasource(attrs.TicketStatusJSON),
		CreateTicketJSON:         util.SafeStringDatasource(attrs.CreateTicketJSON),
		ConnectionType:           util.SafeStringDatasource(attrs.ConnectionType),
		EndpointsFilter:          util.SafeStringDatasource(attrs.ENDPOINTS_FILTER),
		PasswdPolicyJSON:         util.SafeStringDatasource(attrs.PasswdPolicyJSON),
		ConfigJSON:               util.SafeStringDatasource(attrs.ConfigJSON),
		AddFFIDAccessJSON:        util.SafeStringDatasource(attrs.AddFFIDAccessJSON),
		RemoveFFIDAccessJSON:     util.SafeStringDatasource(attrs.RemoveFFIDAccessJSON),
		StatusThresholdConfig:    util.SafeStringDatasource(attrs.STATUS_THRESHOLD_CONFIG),
		ModifyUserDataJSON:       util.SafeStringDatasource(attrs.MODIFYUSERDATAJSON),
		SendOtpJSON:              util.SafeStringDatasource(attrs.SendOtpJSON),
		ValidateOtpJSON:          util.SafeStringDatasource(attrs.ValidateOtpJSON),
		PamConfig:                util.SafeStringDatasource(attrs.PAM_CONFIG),
		CreateAccountJSON:        util.SafeStringDatasource(attrs.CreateAccountJSON),
		UpdateAccountJSON:        util.SafeStringDatasource(attrs.UpdateAccountJSON),
		EnableAccountJSON:        util.SafeStringDatasource(attrs.EnableAccountJSON),
		DisableAccountJSON:       util.SafeStringDatasource(attrs.DisableAccountJSON),
		AddAccessJSON:            util.SafeStringDatasource(attrs.AddAccessJSON),
		RemoveAccessJSON:         util.SafeStringDatasource(attrs.RemoveAccessJSON),
		ImportUserJSON:           util.SafeStringDatasource(attrs.ImportUserJSON),
		IsTimeoutSupported:       util.SafeBoolDatasource(attrs.IsTimeoutSupported),
		ImportAccountEntJSON:     util.SafeStringDatasource(attrs.ImportAccountEntJSON),
		IsTimeoutConfigValidated: util.SafeBoolDatasource(attrs.IsTimeoutConfigValidated),
		// TER-176 entitlement management fields
		ApplicationDiscoveryJson: util.SafeStringDatasource(attrs.ApplicationDiscoveryJSON),
		CreateEntitlementJson:    util.SafeStringDatasource(attrs.CreateEntitlementJSON),
		DeleteEntitlementJson:    util.SafeStringDatasource(attrs.DeleteEntitlementJSON),
		UpdateEntitlementJson:    util.SafeStringDatasource(attrs.UpdateEntitlementJSON),
	}

	// Map connection timeout config if present
	d.MapRESTTimeoutConfig(attrs, state.ConnectionAttributes)
}

// MapRESTTimeoutConfig maps connection timeout configuration from API response to state model
// Only maps timeout config if it exists in the API response
func (d *restConnectionDatasource) MapRESTTimeoutConfig(attrs *openapi.RESTConnectionAttributes, connectionAttrs *RESTConnectionAttributes) {
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

// HandleRESTAuthenticationLogic processes the authenticate flag to control sensitive data visibility
// When authenticate=false, connection_attributes are removed from state to prevent sensitive data exposure
// When authenticate=true, all connection_attributes are returned in state
func (d *restConnectionDatasource) HandleRESTAuthenticationLogic(state *RESTConnectionDataSourceModel, resp *datasource.ReadResponse) {
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
