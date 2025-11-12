// Copyright (c) 2025 Saviynt Inc.
// SPDX-License-Identifier: MPL-2.0

// saviynt_entraid_connection_datasource retrieves entra id connections details from the Saviynt Security Manager.
// The data source supports a single Read operation to look up an existing entra id connections by name.
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

var _ datasource.DataSource = &EntraIdConnectionDataSource{}

// Initialize error codes for EntraID Connection datasource operations
var entraIdDatasourceErrorCodes = errorsutil.NewConnectorErrorCodes(errorsutil.ConnectorTypeEntraID)

// EntraIDConnectionDataSource defines the data source
type EntraIdConnectionDataSource struct {
	client            client.SaviyntClientInterface
	token             string
	provider          client.SaviyntProviderInterface
	connectionFactory client.ConnectionFactoryInterface
}

type EntraIdConnectionDataSourceModel struct {
	ID types.String `tfsdk:"id"`
	BaseConnectionDataSourceModel
	ConnectionAttributes *EntraIdConnectionAttributes `tfsdk:"connection_attributes"`
}

type EntraIdConnectionAttributes struct {
	UpdateUserJSON           types.String             `tfsdk:"update_user_json"`
	MicrosoftGraphEndpoint   types.String             `tfsdk:"microsoft_graph_endpoint"`
	EndpointsFilter          types.String             `tfsdk:"endpoints_filter"`
	ImportUserJSON           types.String             `tfsdk:"import_user_json"`
	ConnectionType           types.String             `tfsdk:"connection_type"`
	EnableAccountJSON        types.String             `tfsdk:"enable_account_json"`
	DeleteGroupJSON          types.String             `tfsdk:"delete_group_json"`
	ConfigJSON               types.String             `tfsdk:"config_json"`
	AddAccessJSON            types.String             `tfsdk:"add_access_json"`
	CreateChannelJSON        types.String             `tfsdk:"create_channel_json"`
	UpdateAccountJSON        types.String             `tfsdk:"update_account_json"`
	IsTimeoutSupported       types.Bool               `tfsdk:"is_timeout_supported"`
	CreateAccountJSON        types.String             `tfsdk:"create_account_json"`
	PamConfig                types.String             `tfsdk:"pam_config"`
	AzureManagementEndpoint  types.String             `tfsdk:"azure_management_endpoint"`
	EntitlementAttribute     types.String             `tfsdk:"entitlement_attribute"`
	AccountsFilter           types.String             `tfsdk:"accounts_filter"`
	DeltaTokensJSON          types.String             `tfsdk:"deltatokens_json"`
	CreateTeamJSON           types.String             `tfsdk:"create_team_json"`
	StatusThresholdConfig    types.String             `tfsdk:"status_threshold_config"`
	AccountImportFields      types.String             `tfsdk:"account_import_fields"`
	RemoveAccountJSON        types.String             `tfsdk:"remove_account_json"`
	EntitlementFilterJSON    types.String             `tfsdk:"entitlement_filter_json"`
	AuthenticationEndpoint   types.String             `tfsdk:"authentication_endpoint"`
	ModifyUserDataJSON       types.String             `tfsdk:"modifyuserdatajson"`
	IsTimeoutConfigValidated types.Bool               `tfsdk:"is_timeout_config_validated"`
	RemoveAccessJSON         types.String             `tfsdk:"remove_access_json"`
	CreateUsers              types.String             `tfsdk:"createusers"`
	DisableAccountJSON       types.String             `tfsdk:"disable_account_json"`
	CreateNewEndpoints       types.String             `tfsdk:"create_new_endpoints"`
	AccountAttributes        types.String             `tfsdk:"account_attributes"`
	AadTenantID              types.String             `tfsdk:"aad_tenant_id"`
	UpdateGroupJSON          types.String             `tfsdk:"update_group_json"`
	CreateGroupJSON          types.String             `tfsdk:"create_group_json"`
	ConnectionTimeoutConfig  *ConnectionTimeoutConfig `tfsdk:"connection_timeout_config"`
}

var _ datasource.DataSource = &EntraIdConnectionDataSource{}

func NewEntraIDConnectionsDataSource() datasource.DataSource {
	return &EntraIdConnectionDataSource{
		connectionFactory: &client.DefaultConnectionFactory{},
	}
}

// NewEntraIDConnectionsDataSourceWithFactory creates a new EntraID connections data source with custom factory
// Used primarily for testing with mock factories
func NewEntraIDConnectionsDataSourceWithFactory(factory client.ConnectionFactoryInterface) datasource.DataSource {
	return &EntraIdConnectionDataSource{
		connectionFactory: factory,
	}
}

// SetClient sets the client for testing purposes
func (d *EntraIdConnectionDataSource) SetClient(client client.SaviyntClientInterface) {
	d.client = client
}

// SetToken sets the token for testing purposes
func (d *EntraIdConnectionDataSource) SetToken(token string) {
	d.token = token
}

// SetProvider sets the provider for testing purposes
func (d *EntraIdConnectionDataSource) SetProvider(provider client.SaviyntProviderInterface) {
	d.provider = provider
}

// Metadata sets the data source type name for Terraform
func (d *EntraIdConnectionDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "saviynt_entraid_connection_datasource"
}

func EntraIDConnectorsDataSourceSchema() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:    true,
			Description: "Resource ID.",
		},
		"connection_attributes": schema.SingleNestedAttribute{
			Computed: true,
			Attributes: map[string]schema.Attribute{
				"update_user_json":            schema.StringAttribute{Computed: true},
				"microsoft_graph_endpoint":    schema.StringAttribute{Computed: true},
				"endpoints_filter":            schema.StringAttribute{Computed: true},
				"import_user_json":            schema.StringAttribute{Computed: true},
				"connection_type":             schema.StringAttribute{Computed: true},
				"enable_account_json":         schema.StringAttribute{Computed: true},
				"delete_group_json":           schema.StringAttribute{Computed: true},
				"config_json":                 schema.StringAttribute{Computed: true},
				"add_access_json":             schema.StringAttribute{Computed: true},
				"create_channel_json":         schema.StringAttribute{Computed: true},
				"update_account_json":         schema.StringAttribute{Computed: true},
				"is_timeout_supported":        schema.BoolAttribute{Computed: true},
				"create_account_json":         schema.StringAttribute{Computed: true},
				"pam_config":                  schema.StringAttribute{Computed: true},
				"azure_management_endpoint":   schema.StringAttribute{Computed: true},
				"entitlement_attribute":       schema.StringAttribute{Computed: true},
				"accounts_filter":             schema.StringAttribute{Computed: true},
				"deltatokens_json":            schema.StringAttribute{Computed: true},
				"create_team_json":            schema.StringAttribute{Computed: true},
				"status_threshold_config":     schema.StringAttribute{Computed: true},
				"account_import_fields":       schema.StringAttribute{Computed: true},
				"remove_account_json":         schema.StringAttribute{Computed: true},
				"entitlement_filter_json":     schema.StringAttribute{Computed: true},
				"authentication_endpoint":     schema.StringAttribute{Computed: true},
				"modifyuserdatajson":          schema.StringAttribute{Computed: true},
				"is_timeout_config_validated": schema.BoolAttribute{Computed: true},
				"remove_access_json":          schema.StringAttribute{Computed: true},
				"createusers":                 schema.StringAttribute{Computed: true},
				"disable_account_json":        schema.StringAttribute{Computed: true},
				"create_new_endpoints":        schema.StringAttribute{Computed: true},
				"account_attributes":          schema.StringAttribute{Computed: true},
				"aad_tenant_id":               schema.StringAttribute{Computed: true},
				"update_group_json":           schema.StringAttribute{Computed: true},
				"create_group_json":           schema.StringAttribute{Computed: true},
				"connection_timeout_config": schema.SingleNestedAttribute{
					Computed:   true,
					Attributes: ConnectionTimeoutConfigeSchema(),
				},
			},
		},
	}
}

// Schema defines the structure and attributes available for the EntraID connection data source
func (d *EntraIdConnectionDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: util.EntraIDConnDataSourceDescription,
		Attributes:  connectionsutil.MergeDataSourceAttributes(BaseConnectorDataSourceSchema(), EntraIDConnectorsDataSourceSchema()),
	}
}

// Configure initializes the data source with provider configuration
// Sets up client and authentication token for API operations
func (d *EntraIdConnectionDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	opCtx := errorsutil.CreateOperationContext(errorsutil.ConnectorTypeEntraID, "configure", "")
	ctx = opCtx.AddContextToLogger(ctx)

	opCtx.LogOperationStart(ctx, "Starting EntraID connection datasource configuration")

	// Check if provider data is available.
	if req.ProviderData == nil {
		tflog.Debug(ctx, "ProviderData is nil, returning early")
		opCtx.LogOperationEnd(ctx, "EntraID connection datasource configuration completed - no provider data")
		return
	}

	// Cast provider data to your provider type.
	prov, ok := req.ProviderData.(*SaviyntProvider)
	if !ok {
		errorCode := entraIdDatasourceErrorCodes.ProviderConfig()
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

	opCtx.LogOperationEnd(ctx, "EntraID connection datasource configured successfully")
}

// Read retrieves EntraID connection details from Saviynt and populates the Terraform state
// Supports lookup by connection name or connection key with comprehensive error handling
func (d *EntraIdConnectionDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var state EntraIdConnectionDataSourceModel

	opCtx := errorsutil.CreateOperationContext(errorsutil.ConnectorTypeEntraID, "datasource_read", "")
	ctx = opCtx.AddContextToLogger(ctx)

	opCtx.LogOperationStart(ctx, "Starting EntraID connection datasource read")

	// Extract configuration from request
	configDiagnostics := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(configDiagnostics...)
	if resp.Diagnostics.HasError() {
		errorCode := entraIdDatasourceErrorCodes.ConfigExtraction()
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
		errorCode := entraIdDatasourceErrorCodes.MissingIdentifier()
		opCtx.LogOperationError(ctx, "Missing connection identifier", errorCode,
			fmt.Errorf("either connection_name or connection_key must be provided"))
		resp.Diagnostics.AddError(
			errorsutil.GetErrorMessage(errorsutil.ErrMissingIdentifier),
			fmt.Sprintf("[%s] Either 'connection_name' or 'connection_key' must be provided to look up the EntraID connection.", errorCode),
		)
		return
	}

	// Execute API call to get EntraID connection details
	apiResp, err := d.ReadEntraIDConnectionDetails(ctx, connectionName, connectionKey)
	if err != nil {
		// Error is already sanitized in ReadEntraIDConnectionDetails method
		errorCode := entraIdDatasourceErrorCodes.ReadFailed()
		opCtx.LogOperationError(ctx, "Failed to read EntraID connection details", errorCode, err)
		resp.Diagnostics.AddError(
			errorsutil.GetErrorMessage(errorsutil.ErrReadFailed),
			fmt.Sprintf("[%s] %s", errorCode, err.Error()),
		)
		return
	}

	// Validate API response
	if err := d.ValidateEntraIDConnectionResponse(apiResp); err != nil {
		errorCode := entraIdDatasourceErrorCodes.APIError()
		opCtx.LogOperationError(ctx, "Invalid connection type for EntraID datasource", errorCode, err)
		resp.Diagnostics.AddError(
			errorsutil.GetErrorMessage(errorsutil.ErrAPIError),
			fmt.Sprintf("[%s] Unable to verify connection type for connection %q. The provider could not determine the type of this connection. Please ensure the connection name is correct and belongs to a supported connector type.", errorCode, connectionName),
		)
		return
	}

	// Map API response to state
	d.UpdateModelFromEntraIDConnectionResponse(&state, apiResp)

	// Handle authentication logic
	d.HandleEntraIDAuthenticationLogic(&state, resp)

	// Set final state
	stateDiagnostics := resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(stateDiagnostics...)
	if resp.Diagnostics.HasError() {
		errorCode := entraIdDatasourceErrorCodes.StateUpdate()
		opCtx.LogOperationError(ctx, "Failed to set state", errorCode,
			fmt.Errorf("state update failed"))
		resp.Diagnostics.AddError(
			errorsutil.GetErrorMessage(errorsutil.ErrStateUpdate),
			fmt.Sprintf("[%s] Unable to update Terraform state for EntraID connection datasource", errorCode),
		)
		return
	}

	opCtx.LogOperationEnd(ctx, "EntraID connection datasource read completed successfully",
		map[string]interface{}{
			"connection_name": connectionName,
			"has_attributes":  state.ConnectionAttributes != nil,
		})
}

// ReadEntraIDConnectionDetails retrieves EntraID connection details from Saviynt API
// Handles both connection name and connection key based lookups using factory pattern
// Returns standardized errors with proper correlation tracking and sensitive data sanitization
func (d *EntraIdConnectionDataSource) ReadEntraIDConnectionDetails(ctx context.Context, connectionName string, connectionKey *int64) (*openapi.GetConnectionDetailsResponse, error) {
	opCtx := errorsutil.CreateOperationContext(errorsutil.ConnectorTypeEntraID, "api_read", connectionName)
	logCtx := opCtx.AddContextToLogger(ctx)

	opCtx.LogOperationStart(logCtx, "Starting EntraID connection API call")

	tflog.Debug(logCtx, "Executing API request to get EntraID connection details")

	var apiResp *openapi.GetConnectionDetailsResponse

	// Execute API request with retry logic
	err := d.provider.AuthenticatedAPICallWithRetry(ctx, "read_entraid_connection_datasource", func(token string) error {
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
		errorCode := entraIdDatasourceErrorCodes.ReadFailed()
		opCtx.LogOperationError(logCtx, "Failed to read EntraID connection details", errorCode, err)
		return nil, errorsutil.CreateStandardError(errorsutil.ConnectorTypeEntraID, errorCode, "api_read", connectionName, err)
	}

	opCtx.LogOperationEnd(logCtx, "EntraID connection API call completed successfully")

	return apiResp, nil
}

// ValidateEntraIDConnectionResponse validates that the API response contains valid EntraID connection data
// Returns standardized error if validation fails
func (d *EntraIdConnectionDataSource) ValidateEntraIDConnectionResponse(apiResp *openapi.GetConnectionDetailsResponse) error {
	if apiResp != nil && apiResp.EntraIDConnectionResponse == nil {
		return fmt.Errorf("verify the connection type - EntraID connection response is nil")
	}
	return nil
}

// UpdateModelFromEntraIDConnectionResponse maps API response data to the Terraform state model
// It handles both base connection fields and detailed connection attributes
func (d *EntraIdConnectionDataSource) UpdateModelFromEntraIDConnectionResponse(state *EntraIdConnectionDataSourceModel, apiResp *openapi.GetConnectionDetailsResponse) {
	// Map base connection fields
	d.MapBaseEntraIDConnectionFields(state, apiResp)

	// Map connection attributes
	d.MapEntraIDConnectionAttributes(state, apiResp)
}

// MapBaseEntraIDConnectionFields maps basic connection fields from API response to state model
// These are common fields available for all connection types
func (d *EntraIdConnectionDataSource) MapBaseEntraIDConnectionFields(state *EntraIdConnectionDataSourceModel, apiResp *openapi.GetConnectionDetailsResponse) {
	state.Msg = util.SafeStringDatasource(apiResp.EntraIDConnectionResponse.Msg)
	state.ErrorCode = util.SafeInt64(apiResp.EntraIDConnectionResponse.Errorcode)
	connectionKey := util.SafeInt64(apiResp.EntraIDConnectionResponse.Connectionkey)
	state.ID = types.StringValue(fmt.Sprintf("ds-entraid-%d", connectionKey.ValueInt64()))
	state.ConnectionName = util.SafeStringDatasource(apiResp.EntraIDConnectionResponse.Connectionname)
	state.ConnectionKey = util.SafeInt64(apiResp.EntraIDConnectionResponse.Connectionkey)
	state.Description = util.SafeStringDatasource(apiResp.EntraIDConnectionResponse.Description)
	state.DefaultSavRoles = util.SafeStringDatasource(apiResp.EntraIDConnectionResponse.Defaultsavroles)
	state.ConnectionType = util.SafeStringDatasource(apiResp.EntraIDConnectionResponse.Connectiontype)
	state.CreatedOn = util.SafeStringDatasource(apiResp.EntraIDConnectionResponse.Createdon)
	state.CreatedBy = util.SafeStringDatasource(apiResp.EntraIDConnectionResponse.Createdby)
	state.UpdatedBy = util.SafeStringDatasource(apiResp.EntraIDConnectionResponse.Updatedby)
	state.EmailTemplate = util.SafeStringDatasource(apiResp.EntraIDConnectionResponse.Emailtemplate)
}

// MapEntraIDConnectionAttributes maps detailed EntraID connection attributes from API response to state model
// Returns nil if no connection attributes are present in the response
func (d *EntraIdConnectionDataSource) MapEntraIDConnectionAttributes(state *EntraIdConnectionDataSourceModel, apiResp *openapi.GetConnectionDetailsResponse) {
	if apiResp.EntraIDConnectionResponse.Connectionattributes == nil {
		state.ConnectionAttributes = nil
		return
	}

	attrs := apiResp.EntraIDConnectionResponse.Connectionattributes
	state.ConnectionAttributes = &EntraIdConnectionAttributes{
		UpdateUserJSON:           util.SafeStringDatasource(attrs.UpdateUserJSON),
		MicrosoftGraphEndpoint:   util.SafeStringDatasource(attrs.MICROSOFT_GRAPH_ENDPOINT),
		EndpointsFilter:          util.SafeStringDatasource(attrs.ENDPOINTS_FILTER),
		ImportUserJSON:           util.SafeStringDatasource(attrs.ImportUserJSON),
		ConnectionType:           util.SafeStringDatasource(attrs.ConnectionType),
		EnableAccountJSON:        util.SafeStringDatasource(attrs.EnableAccountJSON),
		DeleteGroupJSON:          util.SafeStringDatasource(attrs.DeleteGroupJSON),
		ConfigJSON:               util.SafeStringDatasource(attrs.ConfigJSON),
		AddAccessJSON:            util.SafeStringDatasource(attrs.AddAccessJSON),
		CreateChannelJSON:        util.SafeStringDatasource(attrs.CreateChannelJSON),
		UpdateAccountJSON:        util.SafeStringDatasource(attrs.UpdateAccountJSON),
		IsTimeoutSupported:       util.SafeBoolDatasource(attrs.IsTimeoutSupported),
		CreateAccountJSON:        util.SafeStringDatasource(attrs.CreateAccountJSON),
		PamConfig:                util.SafeStringDatasource(attrs.PAM_CONFIG),
		AzureManagementEndpoint:  util.SafeStringDatasource(attrs.AZURE_MANAGEMENT_ENDPOINT),
		EntitlementAttribute:     util.SafeStringDatasource(attrs.ENTITLEMENT_ATTRIBUTE),
		AccountsFilter:           util.SafeStringDatasource(attrs.ACCOUNTS_FILTER),
		DeltaTokensJSON:          util.SafeStringDatasource(attrs.DELTATOKENSJSON),
		CreateTeamJSON:           util.SafeStringDatasource(attrs.CreateTeamJSON),
		StatusThresholdConfig:    util.SafeStringDatasource(attrs.STATUS_THRESHOLD_CONFIG),
		AccountImportFields:      util.SafeStringDatasource(attrs.ACCOUNT_IMPORT_FIELDS),
		RemoveAccountJSON:        util.SafeStringDatasource(attrs.RemoveAccountJSON),
		EntitlementFilterJSON:    util.SafeStringDatasource(attrs.ENTITLEMENT_FILTER_JSON),
		AuthenticationEndpoint:   util.SafeStringDatasource(attrs.AUTHENTICATION_ENDPOINT),
		ModifyUserDataJSON:       util.SafeStringDatasource(attrs.MODIFYUSERDATAJSON),
		IsTimeoutConfigValidated: util.SafeBoolDatasource(attrs.IsTimeoutConfigValidated),
		RemoveAccessJSON:         util.SafeStringDatasource(attrs.RemoveAccessJSON),
		CreateUsers:              util.SafeStringDatasource(attrs.CREATEUSERS),
		DisableAccountJSON:       util.SafeStringDatasource(attrs.DisableAccountJSON),
		CreateNewEndpoints:       util.SafeStringDatasource(attrs.CREATE_NEW_ENDPOINTS),
		AccountAttributes:        util.SafeStringDatasource(attrs.ACCOUNT_ATTRIBUTES),
		AadTenantID:              util.SafeStringDatasource(attrs.AAD_TENANT_ID),
		UpdateGroupJSON:          util.SafeStringDatasource(attrs.UpdateGroupJSON),
		CreateGroupJSON:          util.SafeStringDatasource(attrs.CreateGroupJSON),
	}

	// Map connection timeout config if present
	d.MapEntraIDTimeoutConfig(attrs, state.ConnectionAttributes)
}

// MapEntraIDTimeoutConfig maps connection timeout configuration from API response to state model
// Only maps timeout config if it exists in the API response
func (d *EntraIdConnectionDataSource) MapEntraIDTimeoutConfig(attrs *openapi.EntraIDConnectionAttributes, connectionAttrs *EntraIdConnectionAttributes) {
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

// HandleEntraIDAuthenticationLogic processes the authenticate flag to control sensitive data visibility
// When authenticate=false, connection_attributes are removed from state to prevent sensitive data exposure
// When authenticate=true, all connection_attributes are returned in state
func (d *EntraIdConnectionDataSource) HandleEntraIDAuthenticationLogic(state *EntraIdConnectionDataSourceModel, resp *datasource.ReadResponse) {
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
