// Copyright (c) 2025 Saviynt Inc.
// SPDX-License-Identifier: MPL-2.0

// saviynt_adsi_connection_datasource retrieves adsi connections details from the Saviynt Security Manager.
// The data source supports a single Read operation to look up an existing adsi connections by name.
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

var _ datasource.DataSource = &AdsiConnectionsDataSource{}

// Initialize error codes for ADSI Connection datasource operations
var adsiDatasourceErrorCodes = errorsutil.NewConnectorErrorCodes(errorsutil.ConnectorTypeADSI)

// ADSIConnectionsDataSource defines the data source
type AdsiConnectionsDataSource struct {
	client            client.SaviyntClientInterface
	token             string
	provider          client.SaviyntProviderInterface
	connectionFactory client.ConnectionFactoryInterface
}

type ADSIConnectionDataSourceModel struct {
	ID types.String `tfsdk:"id"`
	BaseConnectionDataSourceModel
	ConnectionAttributes *ADSIConnectionAttributes `tfsdk:"connection_attributes"`
}

type ADSIConnectionAttributes struct {
	ImportNestedMembership      types.String             `tfsdk:"import_nested_membership"`
	PASSWDPOLICYJSON            types.String             `tfsdk:"password_policy_json"`
	CREATEACCOUNTJSON           types.String             `tfsdk:"create_account_json"`
	ENDPOINTS_FILTER            types.String             `tfsdk:"endpoints_filter"`
	DISABLEACCOUNTJSON          types.String             `tfsdk:"disable_account_json"`
	REMOVEACCESSENTITLEMENTJSON types.String             `tfsdk:"remove_access_entitlement_json"`
	GroupSearchBaseDN           types.String             `tfsdk:"group_search_base_dn"`
	ConnectionType              types.String             `tfsdk:"connection_type"`
	STATUSKEYJSON               types.String             `tfsdk:"status_key_json"`
	DEFAULT_USER_ROLE           types.String             `tfsdk:"default_user_role"`
	FOREST_DETAILS              types.String             `tfsdk:"forest_details"`
	UPDATESERVICEACCOUNTJSON    types.String             `tfsdk:"update_service_account_json"`
	ADDACCESSJSON               types.String             `tfsdk:"add_access_json"`
	CREATESERVICEACCOUNTJSON    types.String             `tfsdk:"create_service_account_json"`
	ACCOUNTNAMERULE             types.String             `tfsdk:"account_name_rule"`
	CONNECTION_URL              types.String             `tfsdk:"connection_url"`
	IsTimeoutSupported          types.Bool               `tfsdk:"is_timeout_supported"`
	CreateUpdateMappings        types.String             `tfsdk:"create_update_mappings"`
	ACCOUNT_ATTRIBUTE           types.String             `tfsdk:"account_attribute"`
	PAM_CONFIG                  types.String             `tfsdk:"pam_config"`
	PAGE_SIZE                   types.String             `tfsdk:"page_size"`
	SEARCHFILTER                types.String             `tfsdk:"search_filter"`
	UPDATEGROUPJSON             types.String             `tfsdk:"update_group_json"`
	CREATEGROUPJSON             types.String             `tfsdk:"create_group_json"`
	ENTITLEMENT_ATTRIBUTE       types.String             `tfsdk:"entitlement_attribute"`
	CHECKFORUNIQUE              types.String             `tfsdk:"check_for_unique"`
	REMOVESERVICEACCOUNTJSON    types.String             `tfsdk:"remove_service_account_json"`
	ConnectionTimeoutConfig     *ConnectionTimeoutConfig `tfsdk:"connection_timeout_config"`
	UPDATEUSERJSON              types.String             `tfsdk:"update_user_json"`
	URL                         types.String             `tfsdk:"url"`
	MOVEACCOUNTJSON             types.String             `tfsdk:"move_account_json"`
	CUSTOMCONFIGJSON            types.String             `tfsdk:"custom_config_json"`
	STATUS_THRESHOLD_CONFIG     types.String             `tfsdk:"status_threshold_config"`
	GroupImportMapping          types.String             `tfsdk:"group_import_mapping"`
	PROVISIONING_URL            types.String             `tfsdk:"provisioning_url"`
	REMOVEGROUPJSON             types.String             `tfsdk:"remove_group_json"`
	REMOVEACCESSJSON            types.String             `tfsdk:"remove_access_json"`
	IMPORTDATACOOKIES           types.String             `tfsdk:"import_data_cookies"`
	RESETANDCHANGEPASSWRDJSON   types.String             `tfsdk:"reset_and_change_password_json"`
	USER_ATTRIBUTE              types.String             `tfsdk:"user_attribute"`
	ADDACCESSENTITLEMENTJSON    types.String             `tfsdk:"add_access_entitlement_json"`
	MODIFYUSERDATAJSON          types.String             `tfsdk:"modify_user_data_json"`
	IsTimeoutConfigValidated    types.Bool               `tfsdk:"is_timeout_config_validated"`
	ENABLEGROUPMANAGEMENT       types.String             `tfsdk:"enable_group_management"`
	ENABLEACCOUNTJSON           types.String             `tfsdk:"enable_account_json"`
	FORESTLIST                  types.String             `tfsdk:"forest_list"`
	OBJECTFILTER                types.String             `tfsdk:"object_filter"`
	UPDATEACCOUNTJSON           types.String             `tfsdk:"update_account_json"`
	REMOVEACCOUNTJSON           types.String             `tfsdk:"remove_account_json"`
}

// NewADSIConnectionsDataSource creates a new ADSI connections data source with default factory
func NewADSIConnectionsDataSource() datasource.DataSource {
	return &AdsiConnectionsDataSource{
		connectionFactory: &client.DefaultConnectionFactory{},
	}
}

// NewADSIConnectionsDataSourceWithFactory creates a new ADSI connections data source with custom factory
// Used primarily for testing with mock factories
func NewADSIConnectionsDataSourceWithFactory(factory client.ConnectionFactoryInterface) datasource.DataSource {
	return &AdsiConnectionsDataSource{
		connectionFactory: factory,
	}
}

// SetClient sets the client for testing purposes
func (d *AdsiConnectionsDataSource) SetClient(client client.SaviyntClientInterface) {
	d.client = client
}

// SetToken sets the token for testing purposes
func (d *AdsiConnectionsDataSource) SetToken(token string) {
	d.token = token
}

// SetProvider sets the provider for testing purposes
func (d *AdsiConnectionsDataSource) SetProvider(provider client.SaviyntProviderInterface) {
	d.provider = provider
}

// Metadata sets the data source type name for Terraform
func (d *AdsiConnectionsDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "saviynt_adsi_connection_datasource"
}

func ADSIConnectorsDataSourceSchema() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:    true,
			Description: "Resource ID.",
		},
		"connection_attributes": schema.SingleNestedAttribute{
			Computed: true,
			Attributes: map[string]schema.Attribute{
				"import_nested_membership":       schema.StringAttribute{Computed: true},
				"password_policy_json":           schema.StringAttribute{Computed: true},
				"create_account_json":            schema.StringAttribute{Computed: true},
				"endpoints_filter":               schema.StringAttribute{Computed: true},
				"disable_account_json":           schema.StringAttribute{Computed: true},
				"remove_access_entitlement_json": schema.StringAttribute{Computed: true},
				"group_search_base_dn":           schema.StringAttribute{Computed: true},
				"connection_type":                schema.StringAttribute{Computed: true},
				"status_key_json":                schema.StringAttribute{Computed: true},
				"default_user_role":              schema.StringAttribute{Computed: true},
				"forest_details":                 schema.StringAttribute{Computed: true},
				"update_service_account_json":    schema.StringAttribute{Computed: true},
				"add_access_json":                schema.StringAttribute{Computed: true},
				"create_service_account_json":    schema.StringAttribute{Computed: true},
				"account_name_rule":              schema.StringAttribute{Computed: true},
				"connection_url":                 schema.StringAttribute{Computed: true},
				"is_timeout_supported":           schema.BoolAttribute{Computed: true},
				"create_update_mappings":         schema.StringAttribute{Computed: true},
				"account_attribute":              schema.StringAttribute{Computed: true},
				"pam_config":                     schema.StringAttribute{Computed: true},
				"page_size":                      schema.StringAttribute{Computed: true},
				"search_filter":                  schema.StringAttribute{Computed: true},
				"update_group_json":              schema.StringAttribute{Computed: true},
				"create_group_json":              schema.StringAttribute{Computed: true},
				"entitlement_attribute":          schema.StringAttribute{Computed: true},
				"check_for_unique":               schema.StringAttribute{Computed: true},
				"remove_service_account_json":    schema.StringAttribute{Computed: true},
				"update_user_json":               schema.StringAttribute{Computed: true},
				"url":                            schema.StringAttribute{Computed: true},
				"move_account_json":              schema.StringAttribute{Computed: true},
				"custom_config_json":             schema.StringAttribute{Computed: true},
				"status_threshold_config":        schema.StringAttribute{Computed: true},
				"group_import_mapping":           schema.StringAttribute{Computed: true},
				"provisioning_url":               schema.StringAttribute{Computed: true},
				"remove_group_json":              schema.StringAttribute{Computed: true},
				"remove_access_json":             schema.StringAttribute{Computed: true},
				"import_data_cookies":            schema.StringAttribute{Computed: true},
				"reset_and_change_password_json": schema.StringAttribute{Computed: true},
				"user_attribute":                 schema.StringAttribute{Computed: true},
				"add_access_entitlement_json":    schema.StringAttribute{Computed: true},
				"modify_user_data_json":          schema.StringAttribute{Computed: true},
				"is_timeout_config_validated":    schema.BoolAttribute{Computed: true},
				"enable_group_management":        schema.StringAttribute{Computed: true},
				"enable_account_json":            schema.StringAttribute{Computed: true},
				"forest_list":                    schema.StringAttribute{Computed: true},
				"object_filter":                  schema.StringAttribute{Computed: true},
				"update_account_json":            schema.StringAttribute{Computed: true},
				"remove_account_json":            schema.StringAttribute{Computed: true},
				"connection_timeout_config": schema.SingleNestedAttribute{
					Computed:   true,
					Attributes: ConnectionTimeoutConfigeSchema(),
				},
			},
		},
	}
}

// Schema defines the structure and attributes available for the ADSI connection data source
func (d *AdsiConnectionsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: util.ADSIConnDataSourceDescription,
		Attributes:  connectionsutil.MergeDataSourceAttributes(BaseConnectorDataSourceSchema(), ADSIConnectorsDataSourceSchema()),
	}
}

// Configure initializes the data source with provider configuration
// Sets up client and authentication token for API operations
func (d *AdsiConnectionsDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	opCtx := errorsutil.CreateOperationContext(errorsutil.ConnectorTypeADSI, "configure", "")
	ctx = opCtx.AddContextToLogger(ctx)

	opCtx.LogOperationStart(ctx, "Starting ADSI connection datasource configuration")

	// Check if provider data is available.
	if req.ProviderData == nil {
		tflog.Debug(ctx, "ProviderData is nil, returning early")
		opCtx.LogOperationEnd(ctx, "ADSI connection datasource configuration completed - no provider data")
		return
	}

	// Cast provider data to your provider type.
	prov, ok := req.ProviderData.(*SaviyntProvider)
	if !ok {
		errorCode := adsiDatasourceErrorCodes.ProviderConfig()
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

	opCtx.LogOperationEnd(ctx, "ADSI connection datasource configured successfully")
}

// Read retrieves ADSI connection details from Saviynt and populates the Terraform state
// Supports lookup by connection name or connection key with comprehensive error handling
func (d *AdsiConnectionsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var state ADSIConnectionDataSourceModel

	opCtx := errorsutil.CreateOperationContext(errorsutil.ConnectorTypeADSI, "datasource_read", "")
	ctx = opCtx.AddContextToLogger(ctx)

	opCtx.LogOperationStart(ctx, "Starting ADSI connection datasource read")

	// Extract configuration from request
	configDiagnostics := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(configDiagnostics...)
	if resp.Diagnostics.HasError() {
		errorCode := adsiDatasourceErrorCodes.ConfigExtraction()
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
		errorCode := adsiDatasourceErrorCodes.MissingIdentifier()
		opCtx.LogOperationError(ctx, "Missing connection identifier", errorCode,
			fmt.Errorf("either connection_name or connection_key must be provided"))
		resp.Diagnostics.AddError(
			errorsutil.GetErrorMessage(errorsutil.ErrMissingIdentifier),
			fmt.Sprintf("[%s] Either 'connection_name' or 'connection_key' must be provided to look up the ADSI connection.", errorCode),
		)
		return
	}

	// Execute API call to get ADSI connection details
	apiResp, err := d.ReadADSIConnectionDetails(ctx, connectionName, connectionKey)
	if err != nil {
		// Error is already sanitized in ReadADSIConnectionDetails method
		errorCode := adsiDatasourceErrorCodes.ReadFailed()
		opCtx.LogOperationError(ctx, "Failed to read ADSI connection details", errorCode, err)
		resp.Diagnostics.AddError(
			errorsutil.GetErrorMessage(errorsutil.ErrReadFailed),
			fmt.Sprintf("[%s] %s", errorCode, err.Error()),
		)
		return
	}

	// Validate API response
	if err := d.ValidateADSIConnectionResponse(apiResp); err != nil {
		errorCode := adsiDatasourceErrorCodes.APIError()
		opCtx.LogOperationError(ctx, "Invalid connection type for ADSI datasource", errorCode, err)
		resp.Diagnostics.AddError(
			errorsutil.GetErrorMessage(errorsutil.ErrAPIError),
			fmt.Sprintf("[%s] Unable to verify connection type for connection %q. The provider could not determine the type of this connection. Please ensure the connection name is correct and belongs to a supported connector type.", errorCode, connectionName),
		)
		return
	}

	// Map API response to state
	d.UpdateModelFromADSIConnectionResponse(&state, apiResp)

	// Handle authentication logic
	d.HandleADSIAuthenticationLogic(&state, resp)

	// Set final state
	stateDiagnostics := resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(stateDiagnostics...)
	if resp.Diagnostics.HasError() {
		errorCode := adsiDatasourceErrorCodes.StateUpdate()
		opCtx.LogOperationError(ctx, "Failed to set state", errorCode,
			fmt.Errorf("state update failed"))
		resp.Diagnostics.AddError(
			errorsutil.GetErrorMessage(errorsutil.ErrStateUpdate),
			fmt.Sprintf("[%s] Unable to update Terraform state for ADSI connection datasource", errorCode),
		)
		return
	}

	opCtx.LogOperationEnd(ctx, "ADSI connection datasource read completed successfully",
		map[string]interface{}{
			"connection_name": connectionName,
			"has_attributes":  state.ConnectionAttributes != nil,
		})
}

// ReadADSIConnectionDetails retrieves ADSI connection details from Saviynt API
// Handles both connection name and connection key based lookups using factory pattern
// Returns standardized errors with proper correlation tracking and sensitive data sanitization
func (d *AdsiConnectionsDataSource) ReadADSIConnectionDetails(ctx context.Context, connectionName string, connectionKey *int64) (*openapi.GetConnectionDetailsResponse, error) {
	opCtx := errorsutil.CreateOperationContext(errorsutil.ConnectorTypeADSI, "api_read", connectionName)
	logCtx := opCtx.AddContextToLogger(ctx)

	opCtx.LogOperationStart(logCtx, "Starting ADSI connection API call")

	tflog.Debug(logCtx, "Executing API request to get ADSI connection details")

	var apiResp *openapi.GetConnectionDetailsResponse

	// Execute API request with retry logic
	err := d.provider.AuthenticatedAPICallWithRetry(ctx, "read_adsi_connection_datasource", func(token string) error {
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
		errorCode := adsiDatasourceErrorCodes.ReadFailed()
		opCtx.LogOperationError(logCtx, "Failed to read ADSI connection details", errorCode, err)
		return nil, errorsutil.CreateStandardError(errorsutil.ConnectorTypeADSI, errorCode, "api_read", connectionName, err)
	}

	opCtx.LogOperationEnd(logCtx, "ADSI connection API call completed successfully")

	return apiResp, nil
}

// ValidateADSIConnectionResponse validates that the API response contains valid ADSI connection data
// Returns standardized error if validation fails
func (d *AdsiConnectionsDataSource) ValidateADSIConnectionResponse(apiResp *openapi.GetConnectionDetailsResponse) error {
	if apiResp != nil && apiResp.ADSIConnectionResponse == nil {
		return fmt.Errorf("verify the connection type - ADSI connection response is nil")
	}
	return nil
}

// UpdateModelFromADSIConnectionResponse maps API response data to the Terraform state model
// It handles both base connection fields and detailed connection attributes
func (d *AdsiConnectionsDataSource) UpdateModelFromADSIConnectionResponse(state *ADSIConnectionDataSourceModel, apiResp *openapi.GetConnectionDetailsResponse) {
	// Map base connection fields
	d.MapBaseADSIConnectionFields(state, apiResp)

	// Map connection attributes
	d.MapADSIConnectionAttributes(state, apiResp)
}

// MapBaseADSIConnectionFields maps basic connection fields from API response to state model
// These are common fields available for all connection types
func (d *AdsiConnectionsDataSource) MapBaseADSIConnectionFields(state *ADSIConnectionDataSourceModel, apiResp *openapi.GetConnectionDetailsResponse) {
	state.Msg = util.SafeStringDatasource(apiResp.ADSIConnectionResponse.Msg)
	state.ErrorCode = util.SafeInt64(apiResp.ADSIConnectionResponse.Errorcode)
	connectionKey := util.SafeInt64(apiResp.ADSIConnectionResponse.Connectionkey)
	state.ID = types.StringValue(fmt.Sprintf("ds-adsi-%d", connectionKey.ValueInt64()))
	state.ConnectionName = util.SafeStringDatasource(apiResp.ADSIConnectionResponse.Connectionname)
	state.ConnectionKey = util.SafeInt64(apiResp.ADSIConnectionResponse.Connectionkey)
	state.Description = util.SafeStringDatasource(apiResp.ADSIConnectionResponse.Description)
	state.DefaultSavRoles = util.SafeStringDatasource(apiResp.ADSIConnectionResponse.Defaultsavroles)
	state.ConnectionType = util.SafeStringDatasource(apiResp.ADSIConnectionResponse.Connectiontype)
	state.CreatedOn = util.SafeStringDatasource(apiResp.ADSIConnectionResponse.Createdon)
	state.CreatedBy = util.SafeStringDatasource(apiResp.ADSIConnectionResponse.Createdby)
	state.UpdatedBy = util.SafeStringDatasource(apiResp.ADSIConnectionResponse.Updatedby)
	state.EmailTemplate = util.SafeStringDatasource(apiResp.ADSIConnectionResponse.Emailtemplate)
}

// MapADSIConnectionAttributes maps detailed ADSI connection attributes from API response to state model
// Returns nil if no connection attributes are present in the response
func (d *AdsiConnectionsDataSource) MapADSIConnectionAttributes(state *ADSIConnectionDataSourceModel, apiResp *openapi.GetConnectionDetailsResponse) {
	if apiResp.ADSIConnectionResponse.Connectionattributes == nil {
		state.ConnectionAttributes = nil
		return
	}

	attrs := apiResp.ADSIConnectionResponse.Connectionattributes
	state.ConnectionAttributes = &ADSIConnectionAttributes{
		ImportNestedMembership:      util.SafeStringDatasource(attrs.ImportNestedMembership),
		PASSWDPOLICYJSON:            util.SafeStringDatasource(attrs.PASSWDPOLICYJSON),
		CREATEACCOUNTJSON:           util.SafeStringDatasource(attrs.CREATEACCOUNTJSON),
		ENDPOINTS_FILTER:            util.SafeStringDatasource(attrs.ENDPOINTS_FILTER),
		DISABLEACCOUNTJSON:          util.SafeStringDatasource(attrs.DISABLEACCOUNTJSON),
		REMOVEACCESSENTITLEMENTJSON: util.SafeStringDatasource(attrs.REMOVEACCESSENTITLEMENTJSON),
		GroupSearchBaseDN:           util.SafeStringDatasource(attrs.GroupSearchBaseDN),
		ConnectionType:              util.SafeStringDatasource(attrs.ConnectionType),
		STATUSKEYJSON:               util.SafeStringDatasource(attrs.STATUSKEYJSON),
		DEFAULT_USER_ROLE:           util.SafeStringDatasource(attrs.DEFAULT_USER_ROLE),
		FOREST_DETAILS:              util.SafeStringDatasource(attrs.FOREST_DETAILS),
		UPDATESERVICEACCOUNTJSON:    util.SafeStringDatasource(attrs.UPDATESERVICEACCOUNTJSON),
		ADDACCESSJSON:               util.SafeStringDatasource(attrs.ADDACCESSJSON),
		CREATESERVICEACCOUNTJSON:    util.SafeStringDatasource(attrs.CREATESERVICEACCOUNTJSON),
		ACCOUNTNAMERULE:             util.SafeStringDatasource(attrs.ACCOUNTNAMERULE),
		CONNECTION_URL:              util.SafeStringDatasource(attrs.CONNECTION_URL),
		IsTimeoutSupported:          util.SafeBoolDatasource(attrs.IsTimeoutSupported),
		CreateUpdateMappings:        util.SafeStringDatasource(attrs.CreateUpdateMappings),
		ACCOUNT_ATTRIBUTE:           util.SafeStringDatasource(attrs.ACCOUNT_ATTRIBUTE),
		PAM_CONFIG:                  util.SafeStringDatasource(attrs.PAM_CONFIG),
		PAGE_SIZE:                   util.SafeStringDatasource(attrs.PAGE_SIZE),
		SEARCHFILTER:                util.SafeStringDatasource(attrs.SEARCHFILTER),
		UPDATEGROUPJSON:             util.SafeStringDatasource(attrs.UPDATEGROUPJSON),
		CREATEGROUPJSON:             util.SafeStringDatasource(attrs.CREATEGROUPJSON),
		ENTITLEMENT_ATTRIBUTE:       util.SafeStringDatasource(attrs.ENTITLEMENT_ATTRIBUTE),
		CHECKFORUNIQUE:              util.SafeStringDatasource(attrs.CHECKFORUNIQUE),
		REMOVESERVICEACCOUNTJSON:    util.SafeStringDatasource(attrs.REMOVESERVICEACCOUNTJSON),
		UPDATEUSERJSON:              util.SafeStringDatasource(attrs.UPDATEUSERJSON),
		URL:                         util.SafeStringDatasource(attrs.URL),
		MOVEACCOUNTJSON:             util.SafeStringDatasource(attrs.MOVEACCOUNTJSON),
		CUSTOMCONFIGJSON:            util.SafeStringDatasource(attrs.CUSTOMCONFIGJSON),
		STATUS_THRESHOLD_CONFIG:     util.SafeStringDatasource(attrs.STATUS_THRESHOLD_CONFIG),
		GroupImportMapping:          util.SafeStringDatasource(attrs.GroupImportMapping),
		PROVISIONING_URL:            util.SafeStringDatasource(attrs.PROVISIONING_URL),
		REMOVEGROUPJSON:             util.SafeStringDatasource(attrs.REMOVEGROUPJSON),
		REMOVEACCESSJSON:            util.SafeStringDatasource(attrs.REMOVEACCESSJSON),
		IMPORTDATACOOKIES:           util.SafeStringDatasource(attrs.IMPORTDATACOOKIES),
		RESETANDCHANGEPASSWRDJSON:   util.SafeStringDatasource(attrs.RESETANDCHANGEPASSWRDJSON),
		USER_ATTRIBUTE:              util.SafeStringDatasource(attrs.USER_ATTRIBUTE),
		ADDACCESSENTITLEMENTJSON:    util.SafeStringDatasource(attrs.ADDACCESSENTITLEMENTJSON),
		MODIFYUSERDATAJSON:          util.SafeStringDatasource(attrs.MODIFYUSERDATAJSON),
		IsTimeoutConfigValidated:    util.SafeBoolDatasource(attrs.IsTimeoutConfigValidated),
		ENABLEGROUPMANAGEMENT:       util.SafeStringDatasource(attrs.ENABLEGROUPMANAGEMENT),
		ENABLEACCOUNTJSON:           util.SafeStringDatasource(attrs.ENABLEACCOUNTJSON),
		FORESTLIST:                  util.SafeStringDatasource(attrs.FORESTLIST),
		OBJECTFILTER:                util.SafeStringDatasource(attrs.OBJECTFILTER),
		UPDATEACCOUNTJSON:           util.SafeStringDatasource(attrs.UPDATEACCOUNTJSON),
		REMOVEACCOUNTJSON:           util.SafeStringDatasource(attrs.REMOVEACCOUNTJSON),
	}

	// Map connection timeout config if present
	d.MapADSITimeoutConfig(attrs, state.ConnectionAttributes)
}

// MapADSITimeoutConfig maps connection timeout configuration from API response to state model
// Only maps timeout config if it exists in the API response
func (d *AdsiConnectionsDataSource) MapADSITimeoutConfig(attrs *openapi.ADSIConnectionAttributes, connectionAttrs *ADSIConnectionAttributes) {
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

// HandleADSIAuthenticationLogic processes the authenticate flag to control sensitive data visibility
// When authenticate=false, connection_attributes are removed from state to prevent sensitive data exposure
// When authenticate=true, all connection_attributes are returned in state
func (d *AdsiConnectionsDataSource) HandleADSIAuthenticationLogic(state *ADSIConnectionDataSourceModel, resp *datasource.ReadResponse) {
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
