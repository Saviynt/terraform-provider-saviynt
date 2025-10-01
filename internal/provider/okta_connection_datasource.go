// Copyright (c) 2025 Saviynt Inc.
// SPDX-License-Identifier: MPL-2.0

// saviynt_okta_connection_datasource retrieves Okta connections details from the Saviynt Security Manager.
// The data source supports a single Read operation to look up an existing Okta connections by name.
package provider

import (
	"context"
	"fmt"
	"strconv"
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

var _ datasource.DataSource = &OktaConnectionsDataSource{}

// Initialize error codes for Okta Connection datasource operations
var oktaDatasourceErrorCodes = errorsutil.NewConnectorErrorCodes(errorsutil.ConnectorTypeOkta)

// OktaConnectionsDataSource defines the data source
type OktaConnectionsDataSource struct {
	client            client.SaviyntClientInterface
	token             string
	connectionFactory client.ConnectionFactoryInterface
}

type OktaConnectionDataSourceModel struct {
	ID types.String `tfsdk:"id"`
	BaseConnectionDataSourceModel
	ConnectionAttributes *OktaConnectionAttributes `tfsdk:"connection_attributes"`
}
type OktaConnectionAttributes struct {
	ImportUrl                     types.String            `tfsdk:"import_url"`
	AccountFieldMappings          types.String            `tfsdk:"account_field_mappings"`
	UserFieldMappings             types.String            `tfsdk:"user_field_mappings"`
	EntitlementTypesMappings      types.String            `tfsdk:"entitlement_types_mappings"`
	ImportInactiveApps            types.String            `tfsdk:"import_inactive_apps"`
	OktaApplicationSecuritySystem types.String            `tfsdk:"okta_application_securitysystem"`
	OktaGroupsFilter              types.String            `tfsdk:"okta_groups_filter"`
	AppAccountFieldMappings       types.String            `tfsdk:"app_account_field_mappings"`
	StatusThresholdConfig         types.String            `tfsdk:"status_threshold_config"`
	AuditFilter                   types.String            `tfsdk:"audit_filter"`
	ModifyUserDataJson            types.String            `tfsdk:"modify_user_data_json"`
	ActivateEndpoint              types.String            `tfsdk:"activate_endpoint"`
	ConfigJson                    types.String            `tfsdk:"config_json"`
	PamConfig                     types.String            `tfsdk:"pam_config"`
	ConnectionTimeoutConfig       ConnectionTimeoutConfig `tfsdk:"connection_timeout_config"`
	IsTimeoutConfigValidated      types.Bool              `tfsdk:"is_timeout_config_validated"`
}

// NewOktaConnectionsDataSource creates a new Okta connections data source with default factory
func NewOktaConnectionsDataSource() datasource.DataSource {
	return &OktaConnectionsDataSource{
		connectionFactory: &client.DefaultConnectionFactory{},
	}
}

// NewOktaConnectionsDataSourceWithFactory creates a new Okta connections data source with custom factory
// Used primarily for testing with mock factories
func NewOktaConnectionsDataSourceWithFactory(factory client.ConnectionFactoryInterface) datasource.DataSource {
	return &OktaConnectionsDataSource{
		connectionFactory: factory,
	}
}

// SetClient sets the client for testing purposes
func (d *OktaConnectionsDataSource) SetClient(client client.SaviyntClientInterface) {
	d.client = client
}

// SetToken sets the token for testing purposes
func (d *OktaConnectionsDataSource) SetToken(token string) {
	d.token = token
}

func (d *OktaConnectionsDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "saviynt_okta_connection_datasource"
}

func OktaConnectorsDataSourceSchema() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:    true,
			Description: "Resource ID.",
		},
		"connection_attributes": schema.SingleNestedAttribute{
			Computed: true,
			Attributes: map[string]schema.Attribute{
				"import_url": schema.StringAttribute{
					Computed: true,
				},
				"account_field_mappings": schema.StringAttribute{
					Computed: true,
				},
				"user_field_mappings": schema.StringAttribute{
					Computed: true,
				},
				"entitlement_types_mappings": schema.StringAttribute{
					Computed: true,
				},
				"import_inactive_apps": schema.StringAttribute{
					Computed: true,
				},
				"okta_application_securitysystem": schema.StringAttribute{
					Computed: true,
				},
				"okta_groups_filter": schema.StringAttribute{
					Computed: true,
				},
				"app_account_field_mappings": schema.StringAttribute{
					Computed: true,
				},
				"status_threshold_config": schema.StringAttribute{
					Computed: true,
				},
				"audit_filter": schema.StringAttribute{
					Computed: true,
				},
				"modify_user_data_json": schema.StringAttribute{
					Computed: true,
				},
				"activate_endpoint": schema.StringAttribute{
					Computed: true,
				},
				"config_json": schema.StringAttribute{
					Computed: true,
				},
				"pam_config": schema.StringAttribute{
					Computed: true,
				},
				"connection_timeout_config": schema.SingleNestedAttribute{
					Computed:   true,
					Attributes: ConnectionTimeoutConfigeSchema(),
				},
				"is_timeout_config_validated": schema.BoolAttribute{Computed: true},
			},
		},
	}
}

func (d *OktaConnectionsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: util.OktaConnDataSourceDescription,
		Attributes:  connectionsutil.MergeDataSourceAttributes(BaseConnectorDataSourceSchema(), OktaConnectorsDataSourceSchema()),
	}
}

// Configure initializes the data source with provider configuration
// Sets up client and authentication token for API operations
func (d *OktaConnectionsDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	opCtx := errorsutil.CreateOperationContext(errorsutil.ConnectorTypeOkta, "configure", "")
	ctx = opCtx.AddContextToLogger(ctx)

	opCtx.LogOperationStart(ctx, "Starting Okta connection datasource configuration")

	// Check if provider data is available.
	if req.ProviderData == nil {
		tflog.Debug(ctx, "ProviderData is nil, returning early")
		opCtx.LogOperationEnd(ctx, "Okta connection datasource configuration completed - no provider data")
		return
	}

	// Cast provider data to your provider type.
	prov, ok := req.ProviderData.(*saviyntProvider)
	if !ok {
		errorCode := oktaDatasourceErrorCodes.ProviderConfig()
		opCtx.LogOperationError(ctx, "Provider configuration failed", errorCode,
			fmt.Errorf("expected *saviyntProvider, got different type"),
			map[string]interface{}{"expected_type": "*saviyntProvider"})

		resp.Diagnostics.AddError(
			errorsutil.GetErrorMessage(errorsutil.ErrProviderConfig),
			fmt.Sprintf("[%s] Expected *saviyntProvider, got different type", errorCode),
		)
		return
	}

	// Set the client and token from the provider state using interface wrapper.
	d.client = &client.SaviyntClientWrapper{Client: prov.client}
	d.token = prov.accessToken

	opCtx.LogOperationEnd(ctx, "Okta connection datasource configured successfully")
}

// ReadOktaConnectionDetails retrieves Okta connection details from Saviynt API
// Handles both connection name and connection key based lookups using factory pattern
// Returns standardized errors with proper correlation tracking and sensitive data sanitization
func (d *OktaConnectionsDataSource) ReadOktaConnectionDetails(ctx context.Context, connectionName string, connectionKey *int64) (*openapi.GetConnectionDetailsResponse, error) {
	opCtx := errorsutil.CreateOperationContext(errorsutil.ConnectorTypeOkta, "api_read", connectionName)
	logCtx := opCtx.AddContextToLogger(ctx)

	opCtx.LogOperationStart(logCtx, "Starting Okta connection API call")

	// Use factory pattern instead of direct client creation
	connectionOps := d.connectionFactory.CreateConnectionOperations(d.client.APIBaseURL(), d.token)

	tflog.Debug(logCtx, "Executing API request to get Okta connection details")

	// Execute API request through interface - use original context for API calls
	reqParams := openapi.GetConnectionDetailsRequest{}
	if connectionName != "" {
		reqParams.SetConnectionname(connectionName)
	}
	if connectionKey != nil {
		reqParams.SetConnectionkey(strconv.FormatInt(*connectionKey, 10))
	}
	apiResp, _, err := connectionOps.GetConnectionDetailsDataSource(ctx, reqParams)
	if err != nil {
		errorCode := oktaDatasourceErrorCodes.ReadFailed()
		opCtx.LogOperationError(logCtx, "Failed to read Okta connection details", errorCode, err)
		return nil, errorsutil.CreateStandardError(errorsutil.ConnectorTypeOkta, errorCode, "api_read", connectionName, err)
	}

	opCtx.LogOperationEnd(logCtx, "Okta connection API call completed successfully")

	return apiResp, nil
}

// ValidateOktaConnectionResponse validates that the API response contains valid Okta connection data
// Returns standardized error if validation fails
func (d *OktaConnectionsDataSource) ValidateOktaConnectionResponse(apiResp *openapi.GetConnectionDetailsResponse) error {
	if apiResp != nil && apiResp.OktaConnectionResponse == nil {
		return fmt.Errorf("verify the connection type - Okta connection response is nil")
	}
	return nil
}

// UpdateModelFromOktaConnectionResponse maps API response data to the Terraform state model
// It handles both base connection fields and detailed connection attributes
func (d *OktaConnectionsDataSource) UpdateModelFromOktaConnectionResponse(state *OktaConnectionDataSourceModel, apiResp *openapi.GetConnectionDetailsResponse) {
	// Map base connection fields
	d.MapBaseOktaConnectionFields(state, apiResp)

	// Map connection attributes
	d.MapOktaConnectionAttributes(state, apiResp)
}

// MapBaseOktaConnectionFields maps basic connection fields from API response to state model
// These are common fields available for all connection types
func (d *OktaConnectionsDataSource) MapBaseOktaConnectionFields(state *OktaConnectionDataSourceModel, apiResp *openapi.GetConnectionDetailsResponse) {
	state.Msg = util.SafeStringDatasource(apiResp.OktaConnectionResponse.Msg)
	state.ErrorCode = util.SafeInt64(apiResp.OktaConnectionResponse.Errorcode)
	state.ID = types.StringValue(fmt.Sprintf("ds-okta-%d", *apiResp.OktaConnectionResponse.Connectionkey))
	state.ConnectionName = util.SafeStringDatasource(apiResp.OktaConnectionResponse.Connectionname)
	state.ConnectionKey = util.SafeInt64(apiResp.OktaConnectionResponse.Connectionkey)
	state.Description = util.SafeStringDatasource(apiResp.OktaConnectionResponse.Description)
	state.DefaultSavRoles = util.SafeStringDatasource(apiResp.OktaConnectionResponse.Defaultsavroles)
	state.ConnectionType = util.SafeStringDatasource(apiResp.OktaConnectionResponse.Connectiontype)
	state.CreatedOn = util.SafeStringDatasource(apiResp.OktaConnectionResponse.Createdon)
	state.CreatedBy = util.SafeStringDatasource(apiResp.OktaConnectionResponse.Createdby)
	state.UpdatedBy = util.SafeStringDatasource(apiResp.OktaConnectionResponse.Updatedby)
	state.EmailTemplate = util.SafeStringDatasource(apiResp.OktaConnectionResponse.Emailtemplate)
}

// MapOktaConnectionAttributes maps detailed Okta connection attributes from API response to state model
// Returns nil if no connection attributes are present in the response
func (d *OktaConnectionsDataSource) MapOktaConnectionAttributes(state *OktaConnectionDataSourceModel, apiResp *openapi.GetConnectionDetailsResponse) {
	if apiResp.OktaConnectionResponse.Connectionattributes == nil {
		state.ConnectionAttributes = nil
		return
	}

	attrs := apiResp.OktaConnectionResponse.Connectionattributes
	state.ConnectionAttributes = &OktaConnectionAttributes{
		ImportUrl:                     util.SafeStringDatasource(attrs.IMPORTURL),
		AccountFieldMappings:          util.SafeStringDatasource(attrs.ACCOUNTFIELDMAPPINGS),
		UserFieldMappings:             util.SafeStringDatasource(attrs.USERFIELDMAPPINGS),
		EntitlementTypesMappings:      util.SafeStringDatasource(attrs.ENTITLEMENTTYPESMAPPINGS),
		ImportInactiveApps:            util.SafeStringDatasource(attrs.IMPORT_INACTIVE_APPS),
		OktaApplicationSecuritySystem: util.SafeStringDatasource(attrs.OKTA_APPLICATION_SECURITYSYSTEM),
		OktaGroupsFilter:              util.SafeStringDatasource(attrs.OKTA_GROUPS_FILTER),
		AppAccountFieldMappings:       util.SafeStringDatasource(attrs.APPACCOUNTFIELDMAPPINGS),
		StatusThresholdConfig:         util.SafeStringDatasource(attrs.STATUS_THRESHOLD_CONFIG),
		AuditFilter:                   util.SafeStringDatasource(attrs.AUDIT_FILTER),
		ModifyUserDataJson:            util.SafeStringDatasource(attrs.MODIFYUSERDATAJSON),
		ActivateEndpoint:              util.SafeStringDatasource(attrs.ACTIVATE_ENDPOINT),
		ConfigJson:                    util.SafeStringDatasource(attrs.ConfigJSON),
		PamConfig:                     util.SafeStringDatasource(attrs.PAM_CONFIG),
		IsTimeoutConfigValidated:      util.SafeBoolDatasource(attrs.IsTimeoutConfigValidated),
	}

	// Map connection timeout config if present
	d.MapOktaTimeoutConfig(attrs, state.ConnectionAttributes)
}

// MapOktaTimeoutConfig maps connection timeout configuration from API response to state model
// Only maps timeout config if it exists in the API response
func (d *OktaConnectionsDataSource) MapOktaTimeoutConfig(attrs *openapi.OktaConnectionAttributes, connectionAttrs *OktaConnectionAttributes) {
	if attrs.ConnectionTimeoutConfig != nil {
		connectionAttrs.ConnectionTimeoutConfig = ConnectionTimeoutConfig{
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

// HandleOktaAuthenticationLogic processes the authenticate flag to control sensitive data visibility
// When authenticate=false, connection_attributes are removed from state to prevent sensitive data exposure
// When authenticate=true, all connection_attributes are returned in state
func (d *OktaConnectionsDataSource) HandleOktaAuthenticationLogic(state *OktaConnectionDataSourceModel, resp *datasource.ReadResponse) {
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

// Read retrieves Okta connection details from Saviynt and populates the Terraform state
// Supports lookup by connection name or connection key with comprehensive error handling
func (d *OktaConnectionsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var state OktaConnectionDataSourceModel

	opCtx := errorsutil.CreateOperationContext(errorsutil.ConnectorTypeOkta, "datasource_read", "")
	ctx = opCtx.AddContextToLogger(ctx)

	opCtx.LogOperationStart(ctx, "Starting Okta connection datasource read")

	// Extract configuration from request
	configDiagnostics := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(configDiagnostics...)
	if resp.Diagnostics.HasError() {
		errorCode := oktaDatasourceErrorCodes.ConfigExtraction()
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

	// Execute API call to get Okta connection details
	apiResp, err := d.ReadOktaConnectionDetails(ctx, connectionName, connectionKey)
	if err != nil {
		// Error is already sanitized in ReadOktaConnectionDetails method
		errorCode := oktaDatasourceErrorCodes.ReadFailed()
		opCtx.LogOperationError(ctx, "Failed to read Okta connection details", errorCode, err)
		resp.Diagnostics.AddError(
			errorsutil.GetErrorMessage(errorsutil.ErrReadFailed),
			fmt.Sprintf("[%s] %s", errorCode, err.Error()),
		)
		return
	}

	// Validate API response
	if err := d.ValidateOktaConnectionResponse(apiResp); err != nil {
		errorCode := oktaDatasourceErrorCodes.APIError()
		opCtx.LogOperationError(ctx, "Invalid connection type for Okta datasource", errorCode, err)
		resp.Diagnostics.AddError(
			errorsutil.GetErrorMessage(errorsutil.ErrAPIError),
			fmt.Sprintf("[%s] Unable to verify connection type for connection %q. The provider could not determine the type of this connection. Please ensure the connection name is correct and belongs to a supported connector type.", errorCode, connectionName),
		)
		return
	}

	// Map API response to state
	d.UpdateModelFromOktaConnectionResponse(&state, apiResp)

	// Handle authentication logic
	d.HandleOktaAuthenticationLogic(&state, resp)

	// Set final state
	stateDiagnostics := resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(stateDiagnostics...)
	if resp.Diagnostics.HasError() {
		errorCode := oktaDatasourceErrorCodes.StateUpdate()
		opCtx.LogOperationError(ctx, "Failed to set state", errorCode,
			fmt.Errorf("state update failed"))
		resp.Diagnostics.AddError(
			errorsutil.GetErrorMessage(errorsutil.ErrStateUpdate),
			fmt.Sprintf("[%s] Unable to update Terraform state for Okta connection datasource", errorCode),
		)
		return
	}

	opCtx.LogOperationEnd(ctx, "Okta connection datasource read completed successfully",
		map[string]interface{}{
			"connection_name": connectionName,
			"has_attributes":  state.ConnectionAttributes != nil,
		})
}
