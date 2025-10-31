// Copyright (c) 2025 Saviynt Inc.
// SPDX-License-Identifier: MPL-2.0

// saviynt_unix_connection_datasource retrieves unix connections details from the Saviynt Security Manager.
// The data source supports a single Read operation to look up an existing unix connections by name.
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

var _ datasource.DataSource = &UnixConnectionsDataSource{}

// Initialize error codes for Unix Connection datasource operations
var unixDatasourceErrorCodes = errorsutil.NewConnectorErrorCodes(errorsutil.ConnectorTypeUnix)

// UnixConnectionsDataSource defines the data source
type UnixConnectionsDataSource struct {
	client            client.SaviyntClientInterface
	token             string
	provider          client.SaviyntProviderInterface
	connectionFactory client.ConnectionFactoryInterface
}

type UnixConnectionDataSourceModel struct {
	ID types.String `tfsdk:"id"`
	BaseConnectionDataSourceModel
	ConnectionAttributes *UnixConnectionAttributes `tfsdk:"connection_attributes"`
}

type UnixConnectionAttributes struct {
	GroupsFile                       types.String             `tfsdk:"groups_file"`
	AccountEntitlementMappingCommand types.String             `tfsdk:"account_entitlement_mapping_command"`
	RemoveAccessCommand              types.String             `tfsdk:"remove_access_command"`
	PEMKeyFile                       types.String             `tfsdk:"pem_key_file"`
	PassThroughConnectionDetails     types.String             `tfsdk:"pass_through_connection_details"`
	DisableAccountCommand            types.String             `tfsdk:"disable_account_command"`
	ConnectionTimeoutConfig          *ConnectionTimeoutConfig `tfsdk:"connection_timeout_config"`
	PortNumber                       types.String             `tfsdk:"port_number"`
	ConnectionType                   types.String             `tfsdk:"connection_type"`
	CreateGroupCommand               types.String             `tfsdk:"create_group_command"`
	AccountsFile                     types.String             `tfsdk:"accounts_file"`
	DeleteGroupCommand               types.String             `tfsdk:"delete_group_command"`
	HostName                         types.String             `tfsdk:"host_name"`
	AddGroupOwnerCommand             types.String             `tfsdk:"add_group_owner_command"`
	StatusThresholdConfig            types.String             `tfsdk:"status_threshold_config"`
	InactiveLockAccount              types.String             `tfsdk:"inactive_lock_account"`
	AddAccessCommand                 types.String             `tfsdk:"add_access_command"`
	UpdateAccountCommand             types.String             `tfsdk:"update_account_command"`
	ShadowFile                       types.String             `tfsdk:"shadow_file"`
	IsTimeoutSupported               types.Bool               `tfsdk:"is_timeout_supported"`
	ProvisionAccountCommand          types.String             `tfsdk:"provision_account_command"`
	FirefighterIDGrantAccessCommand  types.String             `tfsdk:"firefighterid_grant_access_command"`
	UnlockAccountCommand             types.String             `tfsdk:"unlock_account_command"`
	DeprovisionAccountCommand        types.String             `tfsdk:"deprovision_account_command"`
	FirefighterIDRevokeAccessCommand types.String             `tfsdk:"firefighterid_revoke_access_command"`
	AddPrimaryGroupCommand           types.String             `tfsdk:"add_primary_group_command"`
	IsTimeoutConfigValidated         types.Bool               `tfsdk:"is_timeout_config_validated"`
	LockAccountCommand               types.String             `tfsdk:"lock_account_command"`
	CustomConfigJSON                 types.String             `tfsdk:"custom_config_json"`
	EnableAccountCommand             types.String             `tfsdk:"enable_account_command"`

	//25.B.1
	ServerType types.String `tfsdk:"server_type"`
}

// NewUnixConnectionsDataSource creates a new Unix connections data source with default factory
func NewUnixConnectionsDataSource() datasource.DataSource {
	return &UnixConnectionsDataSource{
		connectionFactory: &client.DefaultConnectionFactory{},
	}
}

// NewUnixConnectionsDataSourceWithFactory creates a new Unix connections data source with custom factory
// Used primarily for testing with mock factories
func NewUnixConnectionsDataSourceWithFactory(factory client.ConnectionFactoryInterface) datasource.DataSource {
	return &UnixConnectionsDataSource{
		connectionFactory: factory,
	}
}

// SetClient sets the client for testing purposes
func (d *UnixConnectionsDataSource) SetClient(client client.SaviyntClientInterface) {
	d.client = client
}

// SetToken sets the token for testing purposes
func (d *UnixConnectionsDataSource) SetToken(token string) {
	d.token = token
}

// SetProvider sets the provider for testing purposes
func (d *UnixConnectionsDataSource) SetProvider(provider client.SaviyntProviderInterface) {
	d.provider = provider
}

// Metadata sets the data source type name for Terraform
func (d *UnixConnectionsDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "saviynt_unix_connection_datasource"
}

func UnixConnectorsDataSourceSchema() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:    true,
			Description: "Resource ID.",
		},
		"connection_attributes": schema.SingleNestedAttribute{
			Computed: true,
			Attributes: map[string]schema.Attribute{
				"groups_file":                         schema.StringAttribute{Computed: true},
				"account_entitlement_mapping_command": schema.StringAttribute{Computed: true},
				"remove_access_command":               schema.StringAttribute{Computed: true},
				"pem_key_file":                        schema.StringAttribute{Computed: true},
				"pass_through_connection_details":     schema.StringAttribute{Computed: true},
				"disable_account_command":             schema.StringAttribute{Computed: true},
				"port_number":                         schema.StringAttribute{Computed: true},
				"connection_type":                     schema.StringAttribute{Computed: true},
				"create_group_command":                schema.StringAttribute{Computed: true},
				"accounts_file":                       schema.StringAttribute{Computed: true},
				"delete_group_command":                schema.StringAttribute{Computed: true},
				"host_name":                           schema.StringAttribute{Computed: true},
				"add_group_owner_command":             schema.StringAttribute{Computed: true},
				"status_threshold_config":             schema.StringAttribute{Computed: true},
				"inactive_lock_account":               schema.StringAttribute{Computed: true},
				"add_access_command":                  schema.StringAttribute{Computed: true},
				"update_account_command":              schema.StringAttribute{Computed: true},
				"shadow_file":                         schema.StringAttribute{Computed: true},
				"is_timeout_supported":                schema.BoolAttribute{Computed: true},
				"provision_account_command":           schema.StringAttribute{Computed: true},
				"firefighterid_grant_access_command":  schema.StringAttribute{Computed: true},
				"unlock_account_command":              schema.StringAttribute{Computed: true},
				"deprovision_account_command":         schema.StringAttribute{Computed: true},
				"firefighterid_revoke_access_command": schema.StringAttribute{Computed: true},
				"add_primary_group_command":           schema.StringAttribute{Computed: true},
				"is_timeout_config_validated":         schema.BoolAttribute{Computed: true},
				"lock_account_command":                schema.StringAttribute{Computed: true},
				"custom_config_json":                  schema.StringAttribute{Computed: true},
				"enable_account_command":              schema.StringAttribute{Computed: true},
				"server_type":                         schema.StringAttribute{Computed: true},
				"connection_timeout_config": schema.SingleNestedAttribute{
					Computed:   true,
					Attributes: ConnectionTimeoutConfigeSchema(),
				},
			},
		},
	}
}

// Schema defines the structure and attributes available for the Unix connection data source
func (d *UnixConnectionsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: util.UnixConnDataSourceDescription,
		Attributes:  connectionsutil.MergeDataSourceAttributes(BaseConnectorDataSourceSchema(), UnixConnectorsDataSourceSchema()),
	}
}

// Configure initializes the data source with provider configuration
// Sets up client and authentication token for API operations
func (d *UnixConnectionsDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	opCtx := errorsutil.CreateOperationContext(errorsutil.ConnectorTypeUnix, "configure", "")
	ctx = opCtx.AddContextToLogger(ctx)

	opCtx.LogOperationStart(ctx, "Starting Unix connection datasource configuration")

	// Check if provider data is available.
	if req.ProviderData == nil {
		tflog.Debug(ctx, "ProviderData is nil, returning early")
		opCtx.LogOperationEnd(ctx, "Unix connection datasource configuration completed - no provider data")
		return
	}

	// Cast provider data to your provider type.
	prov, ok := req.ProviderData.(*SaviyntProvider)
	if !ok {
		errorCode := unixDatasourceErrorCodes.ProviderConfig()
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

	opCtx.LogOperationEnd(ctx, "Unix connection datasource configured successfully")
}

// Read retrieves Unix connection details from Saviynt and populates the Terraform state
// Supports lookup by connection name or connection key with comprehensive error handling
func (d *UnixConnectionsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var state UnixConnectionDataSourceModel

	opCtx := errorsutil.CreateOperationContext(errorsutil.ConnectorTypeUnix, "datasource_read", "")
	ctx = opCtx.AddContextToLogger(ctx)

	opCtx.LogOperationStart(ctx, "Starting Unix connection datasource read")

	// Extract configuration from request
	configDiagnostics := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(configDiagnostics...)
	if resp.Diagnostics.HasError() {
		errorCode := unixDatasourceErrorCodes.ConfigExtraction()
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
		errorCode := unixDatasourceErrorCodes.MissingIdentifier()
		opCtx.LogOperationError(ctx, "Missing connection identifier", errorCode,
			fmt.Errorf("either connection_name or connection_key must be provided"))
		resp.Diagnostics.AddError(
			errorsutil.GetErrorMessage(errorsutil.ErrMissingIdentifier),
			fmt.Sprintf("[%s] Either 'connection_name' or 'connection_key' must be provided to look up the Unix connection.", errorCode),
		)
		return
	}

	// Execute API call to get Unix connection details
	apiResp, err := d.ReadUnixConnectionDetails(ctx, connectionName, connectionKey)
	if err != nil {
		// Error is already sanitized in ReadUnixConnectionDetails method
		errorCode := unixDatasourceErrorCodes.ReadFailed()
		opCtx.LogOperationError(ctx, "Failed to read Unix connection details", errorCode, err)
		resp.Diagnostics.AddError(
			errorsutil.GetErrorMessage(errorsutil.ErrReadFailed),
			fmt.Sprintf("[%s] %s", errorCode, err.Error()),
		)
		return
	}

	// Validate API response
	if err := d.ValidateUnixConnectionResponse(apiResp); err != nil {
		errorCode := unixDatasourceErrorCodes.APIError()
		opCtx.LogOperationError(ctx, "Invalid connection type for Unix datasource", errorCode, err)
		resp.Diagnostics.AddError(
			errorsutil.GetErrorMessage(errorsutil.ErrAPIError),
			fmt.Sprintf("[%s] Unable to verify connection type for connection %q. The provider could not determine the type of this connection. Please ensure the connection name is correct and belongs to a supported connector type.", errorCode, connectionName),
		)
		return
	}

	// Map API response to state
	d.UpdateModelFromUnixConnectionResponse(&state, apiResp)

	// Handle authentication logic
	d.HandleUnixAuthenticationLogic(&state, resp)

	// Set final state
	stateDiagnostics := resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(stateDiagnostics...)
	if resp.Diagnostics.HasError() {
		errorCode := unixDatasourceErrorCodes.StateUpdate()
		opCtx.LogOperationError(ctx, "Failed to set state", errorCode,
			fmt.Errorf("state update failed"))
		resp.Diagnostics.AddError(
			errorsutil.GetErrorMessage(errorsutil.ErrStateUpdate),
			fmt.Sprintf("[%s] Unable to update Terraform state for Unix connection datasource", errorCode),
		)
		return
	}

	opCtx.LogOperationEnd(ctx, "Unix connection datasource read completed successfully",
		map[string]interface{}{
			"connection_name": connectionName,
			"has_attributes":  state.ConnectionAttributes != nil,
		})
}

// ReadUnixConnectionDetails retrieves Unix connection details from Saviynt API
// Handles both connection name and connection key based lookups using factory pattern
// Returns standardized errors with proper correlation tracking and sensitive data sanitization
func (d *UnixConnectionsDataSource) ReadUnixConnectionDetails(ctx context.Context, connectionName string, connectionKey *int64) (*openapi.GetConnectionDetailsResponse, error) {
	opCtx := errorsutil.CreateOperationContext(errorsutil.ConnectorTypeUnix, "api_read", connectionName)
	logCtx := opCtx.AddContextToLogger(ctx)

	opCtx.LogOperationStart(logCtx, "Starting Unix connection API call")

	tflog.Debug(logCtx, "Executing API request to get Unix connection details")

	var apiResp *openapi.GetConnectionDetailsResponse

	// Execute API request with retry logic
	err := d.provider.AuthenticatedAPICallWithRetry(ctx, "read_unix_connection_datasource", func(token string) error {
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
		errorCode := unixDatasourceErrorCodes.ReadFailed()
		opCtx.LogOperationError(logCtx, "Failed to read Unix connection details", errorCode, err)
		return nil, errorsutil.CreateStandardError(errorsutil.ConnectorTypeUnix, errorCode, "api_read", connectionName, err)
	}

	opCtx.LogOperationEnd(logCtx, "Unix connection API call completed successfully")

	return apiResp, nil
}

// ValidateUnixConnectionResponse validates that the API response contains valid Unix connection data
// Returns standardized error if validation fails
func (d *UnixConnectionsDataSource) ValidateUnixConnectionResponse(apiResp *openapi.GetConnectionDetailsResponse) error {
	if apiResp != nil && apiResp.UNIXConnectionResponse == nil {
		return fmt.Errorf("verify the connection type - Unix connection response is nil")
	}
	return nil
}

// UpdateModelFromUnixConnectionResponse maps API response data to the Terraform state model
// It handles both base connection fields and detailed connection attributes
func (d *UnixConnectionsDataSource) UpdateModelFromUnixConnectionResponse(state *UnixConnectionDataSourceModel, apiResp *openapi.GetConnectionDetailsResponse) {
	// Map base connection fields
	d.MapBaseUnixConnectionFields(state, apiResp)

	// Map connection attributes
	d.MapUnixConnectionAttributes(state, apiResp)
}

// MapBaseUnixConnectionFields maps basic connection fields from API response to state model
// These are common fields available for all connection types
func (d *UnixConnectionsDataSource) MapBaseUnixConnectionFields(state *UnixConnectionDataSourceModel, apiResp *openapi.GetConnectionDetailsResponse) {
	state.Msg = util.SafeStringDatasource(apiResp.UNIXConnectionResponse.Msg)
	state.ErrorCode = util.SafeInt64(apiResp.UNIXConnectionResponse.Errorcode)
	connectionKey := util.SafeInt64(apiResp.UNIXConnectionResponse.Connectionkey)
	state.ID = types.StringValue(fmt.Sprintf("ds-unix-%d", connectionKey.ValueInt64()))
	state.ConnectionName = util.SafeStringDatasource(apiResp.UNIXConnectionResponse.Connectionname)
	state.ConnectionKey = util.SafeInt64(apiResp.UNIXConnectionResponse.Connectionkey)
	state.Description = util.SafeStringDatasource(apiResp.UNIXConnectionResponse.Description)
	state.DefaultSavRoles = util.SafeStringDatasource(apiResp.UNIXConnectionResponse.Defaultsavroles)
	state.ConnectionType = util.SafeStringDatasource(apiResp.UNIXConnectionResponse.Connectiontype)
	state.CreatedOn = util.SafeStringDatasource(apiResp.UNIXConnectionResponse.Createdon)
	state.CreatedBy = util.SafeStringDatasource(apiResp.UNIXConnectionResponse.Createdby)
	state.UpdatedBy = util.SafeStringDatasource(apiResp.UNIXConnectionResponse.Updatedby)
	state.EmailTemplate = util.SafeStringDatasource(apiResp.UNIXConnectionResponse.Emailtemplate)
}

// MapUnixConnectionAttributes maps detailed Unix connection attributes from API response to state model
// Returns nil if no connection attributes are present in the response
func (d *UnixConnectionsDataSource) MapUnixConnectionAttributes(state *UnixConnectionDataSourceModel, apiResp *openapi.GetConnectionDetailsResponse) {
	if apiResp.UNIXConnectionResponse.Connectionattributes == nil {
		state.ConnectionAttributes = nil
		return
	}

	attrs := apiResp.UNIXConnectionResponse.Connectionattributes
	state.ConnectionAttributes = &UnixConnectionAttributes{
		GroupsFile:                       util.SafeStringDatasource(attrs.GROUPS_FILE),
		AccountEntitlementMappingCommand: util.SafeStringDatasource(attrs.ACCOUNT_ENTITLEMENT_MAPPING_COMMAND),
		RemoveAccessCommand:              util.SafeStringDatasource(attrs.REMOVE_ACCESS_COMMAND),
		PEMKeyFile:                       util.SafeStringDatasource(attrs.PEM_KEY_FILE),
		PassThroughConnectionDetails:     util.SafeStringDatasource(attrs.PassThroughConnectionDetails),
		DisableAccountCommand:            util.SafeStringDatasource(attrs.DISABLE_ACCOUNT_COMMAND),
		PortNumber:                       util.SafeStringDatasource(attrs.PORT_NUMBER),
		ConnectionType:                   util.SafeStringDatasource(attrs.ConnectionType),
		CreateGroupCommand:               util.SafeStringDatasource(attrs.CREATE_GROUP_COMMAND),
		AccountsFile:                     util.SafeStringDatasource(attrs.ACCOUNTS_FILE),
		DeleteGroupCommand:               util.SafeStringDatasource(attrs.DELETE_GROUP_COMMAND),
		HostName:                         util.SafeStringDatasource(attrs.HOST_NAME),
		AddGroupOwnerCommand:             util.SafeStringDatasource(attrs.ADD_GROUP_OWNER_COMMAND),
		StatusThresholdConfig:            util.SafeStringDatasource(attrs.STATUS_THRESHOLD_CONFIG),
		InactiveLockAccount:              util.SafeStringDatasource(attrs.INACTIVE_LOCK_ACCOUNT),
		AddAccessCommand:                 util.SafeStringDatasource(attrs.ADD_ACCESS_COMMAND),
		UpdateAccountCommand:             util.SafeStringDatasource(attrs.UPDATE_ACCOUNT_COMMAND),
		ShadowFile:                       util.SafeStringDatasource(attrs.SHADOW_FILE),
		IsTimeoutSupported:               util.SafeBoolDatasource(attrs.IsTimeoutSupported),
		ProvisionAccountCommand:          util.SafeStringDatasource(attrs.PROVISION_ACCOUNT_COMMAND),
		FirefighterIDGrantAccessCommand:  util.SafeStringDatasource(attrs.FIREFIGHTERID_GRANT_ACCESS_COMMAND),
		UnlockAccountCommand:             util.SafeStringDatasource(attrs.UNLOCK_ACCOUNT_COMMAND),
		DeprovisionAccountCommand:        util.SafeStringDatasource(attrs.DEPROVISION_ACCOUNT_COMMAND),
		FirefighterIDRevokeAccessCommand: util.SafeStringDatasource(attrs.FIREFIGHTERID_REVOKE_ACCESS_COMMAND),
		AddPrimaryGroupCommand:           util.SafeStringDatasource(attrs.ADD_PRIMARY_GROUP_COMMAND),
		IsTimeoutConfigValidated:         util.SafeBoolDatasource(attrs.IsTimeoutConfigValidated),
		LockAccountCommand:               util.SafeStringDatasource(attrs.LOCK_ACCOUNT_COMMAND),
		CustomConfigJSON:                 util.SafeStringDatasource(attrs.CUSTOM_CONFIG_JSON),
		EnableAccountCommand:             util.SafeStringDatasource(attrs.ENABLE_ACCOUNT_COMMAND),

		ServerType: util.SafeStringDatasource(attrs.SERVER_TYPE),
	}

	// Map connection timeout config if present
	d.MapUnixTimeoutConfig(attrs, state.ConnectionAttributes)
}

// MapUnixTimeoutConfig maps connection timeout configuration from API response to state model
// Only maps timeout config if it exists in the API response
func (d *UnixConnectionsDataSource) MapUnixTimeoutConfig(attrs *openapi.UNIXConnectionAttributes, connectionAttrs *UnixConnectionAttributes) {
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

// HandleUnixAuthenticationLogic processes the authenticate flag to control sensitive data visibility
// When authenticate=false, connection_attributes are removed from state to prevent sensitive data exposure
// When authenticate=true, all connection_attributes are returned in state
func (d *UnixConnectionsDataSource) HandleUnixAuthenticationLogic(state *UnixConnectionDataSourceModel, resp *datasource.ReadResponse) {
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
