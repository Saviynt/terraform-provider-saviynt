// Copyright (c) 2025 Saviynt Inc.
// SPDX-License-Identifier: MPL-2.0

// saviynt_github_rest_connection_datasource retrieves github rest connections details from the Saviynt Security Manager.
// The data source supports a single Read operation to look up an existing github rest connections by name.
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

var _ datasource.DataSource = &GithubRestConnectionDataSource{}

// Initialize error codes for GitHub REST Connection datasource operations
var githubRestDatasourceErrorCodes = errorsutil.NewConnectorErrorCodes(errorsutil.ConnectorTypeGithubREST)

// GithubRestConnectionDataSource defines the data source
type GithubRestConnectionDataSource struct {
	client            client.SaviyntClientInterface
	token             string
	provider          client.SaviyntProviderInterface
	connectionFactory client.ConnectionFactoryInterface
}

type GithubRestConnectionDataSourceModel struct {
	ID types.String `tfsdk:"id"`
	BaseConnectionDataSourceModel
	ConnectionAttributes *GithubRestConnectionAttributes `tfsdk:"connection_attributes"`
}

type GithubRestConnectionAttributes struct {
	IsTimeoutSupported       types.Bool               `tfsdk:"is_timeout_supported"`
	OrganizationList         types.String             `tfsdk:"organization_list"`
	ImportAccountEntJSON     types.String             `tfsdk:"import_account_ent_json"`
	ConnectionTimeoutConfig  *ConnectionTimeoutConfig `tfsdk:"connection_timeout_config"`
	ConnectionType           types.String             `tfsdk:"connection_type"`
	IsTimeoutConfigValidated types.Bool               `tfsdk:"is_timeout_config_validated"`
}

// NewGithubRestConnectionsDataSource creates a new GitHub REST connections data source with default factory
func NewGithubRestConnectionsDataSource() datasource.DataSource {
	return &GithubRestConnectionDataSource{
		connectionFactory: &client.DefaultConnectionFactory{},
	}
}

// NewGithubRestConnectionsDataSourceWithFactory creates a new GitHub REST connections data source with custom factory
// Used primarily for testing with mock factories
func NewGithubRestConnectionsDataSourceWithFactory(factory client.ConnectionFactoryInterface) datasource.DataSource {
	return &GithubRestConnectionDataSource{
		connectionFactory: factory,
	}
}

// SetClient sets the client for testing purposes
func (d *GithubRestConnectionDataSource) SetClient(client client.SaviyntClientInterface) {
	d.client = client
}

// SetToken sets the token for testing purposes
func (d *GithubRestConnectionDataSource) SetToken(token string) {
	d.token = token
}

// SetProvider sets the provider for testing purposes
func (d *GithubRestConnectionDataSource) SetProvider(provider client.SaviyntProviderInterface) {
	d.provider = provider
}

func (d *GithubRestConnectionDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "saviynt_github_rest_connection_datasource"
}

func GithubRestConnectorsDataSourceSchema() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:    true,
			Description: "Resource ID.",
		},
		"connection_attributes": schema.SingleNestedAttribute{
			Computed: true,
			Attributes: map[string]schema.Attribute{
				"is_timeout_supported":        schema.BoolAttribute{Computed: true},
				"organization_list":           schema.StringAttribute{Computed: true},
				"connection_type":             schema.StringAttribute{Computed: true},
				"import_account_ent_json":     schema.StringAttribute{Computed: true},
				"is_timeout_config_validated": schema.BoolAttribute{Computed: true},
				"connection_timeout_config": schema.SingleNestedAttribute{
					Computed:   true,
					Attributes: ConnectionTimeoutConfigeSchema(),
				},
			},
		},
	}
}

// Schema defines the attributes for the data source
func (d *GithubRestConnectionDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: util.GithubRestConnDataSourceDescription,
		Attributes:  connectionsutil.MergeDataSourceAttributes(BaseConnectorDataSourceSchema(), GithubRestConnectorsDataSourceSchema()),
	}
}

// Configure initializes the data source with provider configuration
// Sets up client and authentication token for API operations
func (d *GithubRestConnectionDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	opCtx := errorsutil.CreateOperationContext(errorsutil.ConnectorTypeGithubREST, "configure", "")
	ctx = opCtx.AddContextToLogger(ctx)

	opCtx.LogOperationStart(ctx, "Starting GitHub REST connection datasource configuration")

	// Check if provider data is available.
	if req.ProviderData == nil {
		tflog.Debug(ctx, "ProviderData is nil, returning early")
		opCtx.LogOperationEnd(ctx, "GitHub REST connection datasource configuration completed - no provider data")
		return
	}

	// Cast provider data to your provider type.
	prov, ok := req.ProviderData.(*SaviyntProvider)
	if !ok {
		errorCode := githubRestDatasourceErrorCodes.ProviderConfig()
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

	opCtx.LogOperationEnd(ctx, "GitHub REST connection datasource configured successfully")
}

// Read retrieves GitHub REST connection details from Saviynt and populates the Terraform state
// Supports lookup by connection name or connection key with comprehensive error handling
func (d *GithubRestConnectionDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var state GithubRestConnectionDataSourceModel

	opCtx := errorsutil.CreateOperationContext(errorsutil.ConnectorTypeGithubREST, "datasource_read", "")
	ctx = opCtx.AddContextToLogger(ctx)

	opCtx.LogOperationStart(ctx, "Starting GitHub REST connection datasource read")

	// Extract configuration from request
	configDiagnostics := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(configDiagnostics...)
	if resp.Diagnostics.HasError() {
		errorCode := githubRestDatasourceErrorCodes.ConfigExtraction()
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
		errorCode := githubRestDatasourceErrorCodes.MissingIdentifier()
		opCtx.LogOperationError(ctx, "Missing connection identifier", errorCode,
			fmt.Errorf("either connection_name or connection_key must be provided"))
		resp.Diagnostics.AddError(
			errorsutil.GetErrorMessage(errorsutil.ErrMissingIdentifier),
			fmt.Sprintf("[%s] Either 'connection_name' or 'connection_key' must be provided to look up the GitHub REST connection.", errorCode),
		)
		return
	}

	// Execute API call to get GitHub REST connection details
	apiResp, err := d.ReadGithubRestConnectionDetails(ctx, connectionName, connectionKey)
	if err != nil {
		// Error is already sanitized in ReadGithubRestConnectionDetails method
		errorCode := githubRestDatasourceErrorCodes.ReadFailed()
		opCtx.LogOperationError(ctx, "Failed to read GitHub REST connection details", errorCode, err)
		resp.Diagnostics.AddError(
			errorsutil.GetErrorMessage(errorsutil.ErrReadFailed),
			fmt.Sprintf("[%s] %s", errorCode, err.Error()),
		)
		return
	}

	// Validate API response
	if err := d.ValidateGithubRestConnectionResponse(apiResp); err != nil {
		errorCode := githubRestDatasourceErrorCodes.APIError()
		opCtx.LogOperationError(ctx, "Invalid connection type for GitHub REST datasource", errorCode, err)
		resp.Diagnostics.AddError(
			errorsutil.GetErrorMessage(errorsutil.ErrAPIError),
			fmt.Sprintf("[%s] Unable to verify connection type for connection %q. The provider could not determine the type of this connection. Please ensure the connection name is correct and belongs to a supported connector type.", errorCode, connectionName),
		)
		return
	}

	// Map API response to state
	d.UpdateModelFromGithubRestConnectionResponse(&state, apiResp)

	// Handle authentication logic
	d.HandleGithubRestAuthenticationLogic(&state, resp)

	// Set final state
	stateDiagnostics := resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(stateDiagnostics...)
	if resp.Diagnostics.HasError() {
		errorCode := githubRestDatasourceErrorCodes.StateUpdate()
		opCtx.LogOperationError(ctx, "Failed to set state", errorCode,
			fmt.Errorf("state update failed"))
		resp.Diagnostics.AddError(
			errorsutil.GetErrorMessage(errorsutil.ErrStateUpdate),
			fmt.Sprintf("[%s] Unable to update Terraform state for GitHub REST connection datasource", errorCode),
		)
		return
	}

	opCtx.LogOperationEnd(ctx, "GitHub REST connection datasource read completed successfully",
		map[string]interface{}{
			"connection_name": connectionName,
			"has_attributes":  state.ConnectionAttributes != nil,
		})
}

// ReadGithubRestConnectionDetails retrieves GitHub REST connection details from Saviynt API
// Handles both connection name and connection key based lookups using factory pattern
// Returns standardized errors with proper correlation tracking and sensitive data sanitization
func (d *GithubRestConnectionDataSource) ReadGithubRestConnectionDetails(ctx context.Context, connectionName string, connectionKey *int64) (*openapi.GetConnectionDetailsResponse, error) {
	opCtx := errorsutil.CreateOperationContext(errorsutil.ConnectorTypeGithubREST, "api_read", connectionName)
	logCtx := opCtx.AddContextToLogger(ctx)

	opCtx.LogOperationStart(logCtx, "Starting GitHub REST connection API call")

	tflog.Debug(logCtx, "Executing API request to get GitHub REST connection details")

	var apiResp *openapi.GetConnectionDetailsResponse

	// Execute API request with retry logic
	err := d.provider.AuthenticatedAPICallWithRetry(ctx, "read_githubrest_connection_datasource", func(token string) error {
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
		errorCode := githubRestDatasourceErrorCodes.ReadFailed()
		opCtx.LogOperationError(logCtx, "Failed to read GitHub REST connection details", errorCode, err)
		return nil, errorsutil.CreateStandardError(errorsutil.ConnectorTypeGithubREST, errorCode, "api_read", connectionName, err)
	}

	opCtx.LogOperationEnd(logCtx, "GitHub REST connection API call completed successfully")

	return apiResp, nil
}

// ValidateGithubRestConnectionResponse validates that the API response contains valid GitHub REST connection data
// Returns standardized error if validation fails
func (d *GithubRestConnectionDataSource) ValidateGithubRestConnectionResponse(apiResp *openapi.GetConnectionDetailsResponse) error {
	if apiResp != nil && apiResp.GithubRESTConnectionResponse == nil {
		return fmt.Errorf("verify the connection type - GitHub REST connection response is nil")
	}
	return nil
}

// UpdateModelFromGithubRestConnectionResponse maps API response data to the Terraform state model
// It handles both base connection fields and detailed connection attributes
func (d *GithubRestConnectionDataSource) UpdateModelFromGithubRestConnectionResponse(state *GithubRestConnectionDataSourceModel, apiResp *openapi.GetConnectionDetailsResponse) {
	// Map base connection fields
	d.MapBaseGithubRestConnectionFields(state, apiResp)

	// Map connection attributes
	d.MapGithubRestConnectionAttributes(state, apiResp)
}

// MapBaseGithubRestConnectionFields maps basic connection fields from API response to state model
// These are common fields available for all connection types
func (d *GithubRestConnectionDataSource) MapBaseGithubRestConnectionFields(state *GithubRestConnectionDataSourceModel, apiResp *openapi.GetConnectionDetailsResponse) {
	state.Msg = util.SafeStringDatasource(apiResp.GithubRESTConnectionResponse.Msg)
	state.ErrorCode = util.SafeInt64(apiResp.GithubRESTConnectionResponse.Errorcode)
	connectionKey := util.SafeInt64(apiResp.GithubRESTConnectionResponse.Connectionkey)
	state.ID = types.StringValue(fmt.Sprintf("ds-githubrest-%d", connectionKey.ValueInt64()))
	state.ConnectionName = util.SafeStringDatasource(apiResp.GithubRESTConnectionResponse.Connectionname)
	state.ConnectionKey = util.SafeInt64(apiResp.GithubRESTConnectionResponse.Connectionkey)
	state.Description = util.SafeStringDatasource(apiResp.GithubRESTConnectionResponse.Description)
	state.DefaultSavRoles = util.SafeStringDatasource(apiResp.GithubRESTConnectionResponse.Defaultsavroles)
	state.ConnectionType = util.SafeStringDatasource(apiResp.GithubRESTConnectionResponse.Connectiontype)
	state.CreatedOn = util.SafeStringDatasource(apiResp.GithubRESTConnectionResponse.Createdon)
	state.CreatedBy = util.SafeStringDatasource(apiResp.GithubRESTConnectionResponse.Createdby)
	state.UpdatedBy = util.SafeStringDatasource(apiResp.GithubRESTConnectionResponse.Updatedby)
	state.EmailTemplate = util.SafeStringDatasource(apiResp.GithubRESTConnectionResponse.Emailtemplate)
}

// MapGithubRestConnectionAttributes maps detailed GitHub REST connection attributes from API response to state model
// Returns nil if no connection attributes are present in the response
func (d *GithubRestConnectionDataSource) MapGithubRestConnectionAttributes(state *GithubRestConnectionDataSourceModel, apiResp *openapi.GetConnectionDetailsResponse) {
	if apiResp.GithubRESTConnectionResponse.Connectionattributes == nil {
		state.ConnectionAttributes = nil
		return
	}

	attrs := apiResp.GithubRESTConnectionResponse.Connectionattributes
	state.ConnectionAttributes = &GithubRestConnectionAttributes{
		IsTimeoutSupported:       util.SafeBoolDatasource(attrs.IsTimeoutSupported),
		OrganizationList:         util.SafeStringDatasource(attrs.ORGANIZATION_LIST),
		ImportAccountEntJSON:     util.SafeStringDatasource(attrs.ImportAccountEntJSON),
		ConnectionType:           util.SafeStringDatasource(attrs.ConnectionType),
		IsTimeoutConfigValidated: util.SafeBoolDatasource(attrs.IsTimeoutConfigValidated),
	}

	// Map connection timeout config if present
	d.MapGithubRestTimeoutConfig(attrs, state.ConnectionAttributes)
}

// MapGithubRestTimeoutConfig maps connection timeout configuration from API response to state model
// Only maps timeout config if it exists in the API response
func (d *GithubRestConnectionDataSource) MapGithubRestTimeoutConfig(attrs *openapi.GithubRESTConnectionAttributes, connectionAttrs *GithubRestConnectionAttributes) {
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

// HandleGithubRestAuthenticationLogic processes the authenticate flag to control sensitive data visibility
// When authenticate=false, connection_attributes are removed from state to prevent sensitive data exposure
// When authenticate=true, all connection_attributes are returned in state
func (d *GithubRestConnectionDataSource) HandleGithubRestAuthenticationLogic(state *GithubRestConnectionDataSourceModel, resp *datasource.ReadResponse) {
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
