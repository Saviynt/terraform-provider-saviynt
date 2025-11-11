// Copyright (c) 2025 Saviynt Inc.
// SPDX-License-Identifier: MPL-2.0

// saviynt_sftp_connection_datasource retrieves SFTP connections details from the Saviynt Security Manager.
// The data source supports a single Read operation to look up an existing SFTP connections by name.
package provider

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
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

var _ datasource.DataSource = &SftpConnectionsDataSource{}

// Initialize error codes for SFTP Connection datasource operations
var sftpDatasourceErrorCodes = errorsutil.NewConnectorErrorCodes(errorsutil.ConnectorTypeSFTP)

// SFTPConnectionsDataSource defines the data source
type SftpConnectionsDataSource struct {
	client            client.SaviyntClientInterface
	token             string
	provider          client.SaviyntProviderInterface
	connectionFactory client.ConnectionFactoryInterface
}

type SFTPConnectionDataSourceModel struct {
	ID types.String `tfsdk:"id"`
	BaseConnectionDataSourceModel
	ConnectionAttributes *SFTPConnectionAttributes `tfsdk:"connection_attributes"`
}

type SFTPConnectionAttributes struct {
	HostName                 types.String             `tfsdk:"host_name"`
	PortNumber               types.String             `tfsdk:"port_number"`
	Username                 types.String             `tfsdk:"username"`
	AuthCredentialType       types.String             `tfsdk:"auth_credential_type"`
	FilesToGet               types.String             `tfsdk:"files_to_get"`
	FilesToPut               types.String             `tfsdk:"files_to_put"`
	PamConfig                types.String             `tfsdk:"pam_config"`
	ConnectionType           types.String             `tfsdk:"connection_type"`
	IsTimeoutSupported       types.Bool               `tfsdk:"is_timeout_supported"`
	IsTimeoutConfigValidated types.Bool               `tfsdk:"is_timeout_config_validated"`
	ConnectionTimeoutConfig  *ConnectionTimeoutConfig `tfsdk:"connection_timeout_config"`
}

// NewSFTPConnectionsDataSource creates a new SFTP connections data source with default factory
func NewSFTPConnectionsDataSource() datasource.DataSource {
	return &SftpConnectionsDataSource{
		connectionFactory: &client.DefaultConnectionFactory{},
	}
}

// NewSFTPConnectionsDataSourceWithFactory creates a new SFTP connections data source with custom factory
// Used primarily for testing with mock factories
func NewSFTPConnectionsDataSourceWithFactory(factory client.ConnectionFactoryInterface) datasource.DataSource {
	return &SftpConnectionsDataSource{
		connectionFactory: factory,
	}
}

// SetClient sets the client for testing purposes
func (d *SftpConnectionsDataSource) SetClient(client client.SaviyntClientInterface) {
	d.client = client
}

// SetToken sets the token for testing purposes
func (d *SftpConnectionsDataSource) SetToken(token string) {
	d.token = token
}

// SetProvider sets the provider for testing purposes
func (d *SftpConnectionsDataSource) SetProvider(provider client.SaviyntProviderInterface) {
	d.provider = provider
}

// Metadata sets the data source type name for Terraform
func (d *SftpConnectionsDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "saviynt_sftp_connection_datasource"
}

func SFTPConnectorsDataSourceSchema() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:    true,
			Description: "Resource ID.",
		},
		"connection_attributes": schema.SingleNestedAttribute{
			Computed: true,
			Attributes: map[string]schema.Attribute{
				"host_name":                   schema.StringAttribute{Computed: true},
				"port_number":                 schema.StringAttribute{Computed: true},
				"username":                    schema.StringAttribute{Computed: true},
				"auth_credential_type":        schema.StringAttribute{Computed: true},
				"files_to_get":                schema.StringAttribute{Computed: true},
				"files_to_put":                schema.StringAttribute{Computed: true},
				"pam_config":                  schema.StringAttribute{Computed: true},
				"connection_type":             schema.StringAttribute{Computed: true},
				"is_timeout_supported":        schema.BoolAttribute{Computed: true},
				"is_timeout_config_validated": schema.BoolAttribute{Computed: true},
				"connection_timeout_config": schema.SingleNestedAttribute{
					Computed:   true,
					Attributes: ConnectionTimeoutConfigeSchema(),
				},
			},
		},
	}
}

// Schema defines the structure and attributes available for the SFTP connection data source
func (d *SftpConnectionsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: util.SFTPConnDataSourceDescription,
		Attributes:  connectionsutil.MergeDataSourceAttributes(BaseConnectorDataSourceSchema(), SFTPConnectorsDataSourceSchema()),
	}
}

// Configure sets up the data source with provider configuration
func (d *SftpConnectionsDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	opCtx := errorsutil.CreateOperationContext(errorsutil.ConnectorTypeSFTP, "configure", "")
	ctx = opCtx.AddContextToLogger(ctx)

	opCtx.LogOperationStart(ctx, "Starting SFTP connection datasource configuration")

	// Check if provider data is available.
	if req.ProviderData == nil {
		tflog.Debug(ctx, "ProviderData is nil, returning early")
		opCtx.LogOperationEnd(ctx, "SFTP connection datasource configuration completed - no provider data")
		return
	}

	// Cast provider data to your provider type.
	prov, ok := req.ProviderData.(*SaviyntProvider)
	if !ok {
		errorCode := sftpDatasourceErrorCodes.ProviderConfig()
		opCtx.LogOperationError(ctx, "Provider configuration failed", errorCode,
			fmt.Errorf("expected *SaviyntProvider, got different type"),
			map[string]interface{}{"expected_type": "*SaviyntProvider"})

		resp.Diagnostics.AddError(
			errorsutil.GetErrorMessage(errorsutil.ErrProviderConfig),
			fmt.Sprintf("[%s] Expected *SaviyntProvider, got different type", errorCode),
		)
		return
	}

	// Set the client, token, and provider reference from the provider state
	d.client = &client.SaviyntClientWrapper{Client: prov.client}
	d.token = prov.accessToken
	d.provider = &client.SaviyntProviderWrapper{Provider: prov}

	opCtx.LogOperationEnd(ctx, "SFTP connection datasource configured successfully")
}

// Read fetches the SFTP connection data from Saviynt
func (d *SftpConnectionsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config SFTPConnectionDataSourceModel

	opCtx := errorsutil.CreateOperationContext(errorsutil.ConnectorTypeSFTP, "datasource_read", "")
	ctx = opCtx.AddContextToLogger(ctx)

	opCtx.LogOperationStart(ctx, "Starting SFTP connection datasource read")

	// Extract configuration from request
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		errorCode := sftpDatasourceErrorCodes.ConfigExtraction()
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
	if !config.ConnectionName.IsNull() && config.ConnectionName.ValueString() != "" {
		connectionName = config.ConnectionName.ValueString()
		opCtx.ConnectionName = connectionName
		ctx = opCtx.AddContextToLogger(ctx)
	}

	// Prepare connection parameters
	var connectionKey *int64
	if !config.ConnectionKey.IsNull() {
		keyValue := config.ConnectionKey.ValueInt64()
		connectionKey = &keyValue
	}

	// Validate that at least one identifier is provided
	if connectionName == "" && connectionKey == nil {
		errorCode := sftpDatasourceErrorCodes.MissingIdentifier()
		opCtx.LogOperationError(ctx, "Missing connection identifier", errorCode,
			fmt.Errorf("either connection_name or connection_key must be provided"))
		resp.Diagnostics.AddError(
			errorsutil.GetErrorMessage(errorsutil.ErrMissingIdentifier),
			fmt.Sprintf("[%s] Either 'connection_name' or 'connection_key' must be provided to look up the SFTP connection.", errorCode),
		)
		return
	}

	// Execute API call to get SFTP connection details
	apiResp, err := d.ReadSFTPConnectionDetails(ctx, connectionName, connectionKey)
	if err != nil {
		// Error is already sanitized in ReadSFTPConnectionDetails method
		errorCode := sftpDatasourceErrorCodes.ReadFailed()
		opCtx.LogOperationError(ctx, "Failed to read SFTP connection details", errorCode, err)
		resp.Diagnostics.AddError(
			errorsutil.GetErrorMessage(errorsutil.ErrReadFailed),
			fmt.Sprintf("[%s] %s", errorCode, err.Error()),
		)
		return
	}

	// Validate API response
	if err := d.ValidateSFTPConnectionResponse(apiResp); err != nil {
		errorCode := sftpDatasourceErrorCodes.APIError()
		connectionIdentifier := d.GetConnectionIdentifier(connectionName, connectionKey)
		opCtx.LogOperationError(ctx, "Invalid connection type for SFTP datasource", errorCode, err)
		resp.Diagnostics.AddError(
			errorsutil.GetErrorMessage(errorsutil.ErrAPIError),
			fmt.Sprintf("[%s] Unable to verify connection type for connection %q. The provider could not determine the type of this connection. Please ensure the connection name is correct and belongs to a supported connector type.", errorCode, connectionIdentifier),
		)
		return
	}

	// Map API response to state
	d.UpdateModelFromReadResponse(&config, apiResp)

	// Handle authentication logic
	d.HandleSFTPAuthenticationLogic(&config, resp)

	// Set final state
	stateDiagnostics := resp.State.Set(ctx, &config)
	resp.Diagnostics.Append(stateDiagnostics...)
	if resp.Diagnostics.HasError() {
		errorCode := sftpDatasourceErrorCodes.StateUpdate()
		opCtx.LogOperationError(ctx, "Failed to set state", errorCode,
			fmt.Errorf("state update failed"))
		resp.Diagnostics.AddError(
			errorsutil.GetErrorMessage(errorsutil.ErrStateUpdate),
			fmt.Sprintf("[%s] Unable to update Terraform state for SFTP connection datasource", errorCode),
		)
		return
	}

	opCtx.LogOperationEnd(ctx, "SFTP connection datasource read completed successfully",
		map[string]interface{}{
			"connection_name": connectionName,
			"has_attributes":  config.ConnectionAttributes != nil,
		})
}

// ValidateSFTPConnectionResponse validates that the API response is for an SFTP connection
func (d *SftpConnectionsDataSource) ValidateSFTPConnectionResponse(apiResp *openapi.GetConnectionDetailsResponse) error {
	if apiResp != nil && apiResp.SFTPConnectionResponse == nil {
		return fmt.Errorf("verify the connection type - SFTP connection response is nil")
	}
	return nil
}

// UpdateModelFromReadResponse updates the data source model from the API response
func (d *SftpConnectionsDataSource) UpdateModelFromReadResponse(state *SFTPConnectionDataSourceModel, apiResp *openapi.GetConnectionDetailsResponse) {
	// Update base connection attributes
	d.MapSFTPBaseAttributes(state, apiResp.SFTPConnectionResponse)

	// Initialize and update SFTP-specific attributes
	d.MapSFTPConnectionAttributes(state, apiResp.SFTPConnectionResponse.Connectionattributes)
}

// MapSFTPBaseAttributes maps base connection attributes from API response to state model
func (d *SftpConnectionsDataSource) MapSFTPBaseAttributes(state *SFTPConnectionDataSourceModel, response *openapi.SFTPConnectionResponse) {
	state.Msg = util.SafeStringDatasource(response.Msg)
	state.ErrorCode = util.SafeInt64(response.Errorcode)
	state.ConnectionKey = types.Int64Value(int64(*response.Connectionkey))
	state.ID = types.StringValue(fmt.Sprintf("ds-sftp-%d", *response.Connectionkey))
	state.ConnectionName = util.SafeStringDatasource(response.Connectionname)
	state.DefaultSavRoles = util.SafeStringDatasource(response.Defaultsavroles)
	state.EmailTemplate = util.SafeStringDatasource(response.Emailtemplate)
	state.Description = util.SafeStringDatasource(response.Description)
	state.ConnectionType = util.SafeStringDatasource(response.Connectiontype)
	state.CreatedBy = util.SafeStringDatasource(response.Createdby)
	state.CreatedOn = util.SafeStringDatasource(response.Createdon)
	state.UpdatedBy = util.SafeStringDatasource(response.Updatedby)
	state.Status = util.SafeInt64(response.Status)
}

// MapSFTPConnectionAttributes maps SFTP-specific connection attributes from API response to state model
func (d *SftpConnectionsDataSource) MapSFTPConnectionAttributes(state *SFTPConnectionDataSourceModel, attrs *openapi.SFTPConnectionAttributes) {
	// Initialize connection attributes if nil
	if state.ConnectionAttributes == nil {
		state.ConnectionAttributes = &SFTPConnectionAttributes{}
	}

	// Update SFTP-specific attributes if present
	if attrs != nil {
		state.ConnectionAttributes.HostName = util.SafeStringDatasource(attrs.HOST_NAME)
		state.ConnectionAttributes.PortNumber = util.SafeStringDatasource(attrs.PORT_NUMBER)
		state.ConnectionAttributes.Username = util.SafeStringDatasource(attrs.USERNAME)
		state.ConnectionAttributes.AuthCredentialType = util.SafeStringDatasource(attrs.AUTH_CREDENTIAL_TYPE)
		state.ConnectionAttributes.FilesToGet = util.SafeStringDatasource(attrs.FILES_TO_GET)
		state.ConnectionAttributes.FilesToPut = util.SafeStringDatasource(attrs.FILES_TO_PUT)
		state.ConnectionAttributes.PamConfig = util.SafeStringDatasource(attrs.PAM_CONFIG)
		state.ConnectionAttributes.ConnectionType = util.SafeStringDatasource(attrs.ConnectionType)
		state.ConnectionAttributes.IsTimeoutSupported = util.SafeBoolDatasource(attrs.IsTimeoutSupported)
		state.ConnectionAttributes.IsTimeoutConfigValidated = util.SafeBoolDatasource(attrs.IsTimeoutConfigValidated)

		// Map connection timeout config if present
		d.MapSFTPTimeoutConfig(attrs, state.ConnectionAttributes)
	}
}

// MapSFTPTimeoutConfig maps connection timeout configuration from API response to state model
// Only maps timeout config if it exists in the API response
func (d *SftpConnectionsDataSource) MapSFTPTimeoutConfig(attrs *openapi.SFTPConnectionAttributes, connectionAttrs *SFTPConnectionAttributes) {
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

// ReadSFTPConnectionDetails retrieves SFTP connection details from Saviynt API
// Handles both connection name and connection key based lookups using factory pattern
// Returns standardized errors with proper correlation tracking and sensitive data sanitization
func (d *SftpConnectionsDataSource) ReadSFTPConnectionDetails(ctx context.Context, connectionName string, connectionKey *int64) (*openapi.GetConnectionDetailsResponse, error) {
	opCtx := errorsutil.CreateOperationContext(errorsutil.ConnectorTypeSFTP, "api_read", connectionName)
	logCtx := opCtx.AddContextToLogger(ctx)

	opCtx.LogOperationStart(logCtx, "Starting SFTP connection API read operation")

	tflog.Debug(logCtx, "Executing API request to get SFTP connection details")

	var apiResp *openapi.GetConnectionDetailsResponse
	var finalHttpResp *http.Response

	// Execute read operation with retry logic
	err := d.provider.AuthenticatedAPICallWithRetry(ctx, "read_sftp_connection_datasource", func(token string) error {
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
		finalHttpResp = httpResp // Update on every call including retries
		return err
	})

	// Handle non-412 errors first
	if err != nil && finalHttpResp != nil && finalHttpResp.StatusCode != 412 {
		connectionIdentifier := d.GetConnectionIdentifier(connectionName, connectionKey)
		errorCode := sftpDatasourceErrorCodes.ReadFailed()
		opCtx.LogOperationError(logCtx, "SFTP connection API read failed", errorCode, err)
		return nil, fmt.Errorf("[%s] Unable to read SFTP connection '%s': %w", errorCode, connectionIdentifier, err)
	}

	// Handle 412 precondition failed with response body decoding
	if err != nil && finalHttpResp != nil && finalHttpResp.StatusCode == 412 {
		connectionIdentifier := d.GetConnectionIdentifier(connectionName, connectionKey)
		var errorResp struct {
			Msg       string `json:"msg"`
			ErrorCode int    `json:"errorcode"`
		}

		if decodeErr := json.NewDecoder(finalHttpResp.Body).Decode(&errorResp); decodeErr == nil {
			errorCode := sftpDatasourceErrorCodes.ReadFailed()
			opCtx.LogOperationError(logCtx, "SFTP connection not found", errorCode,
				fmt.Errorf("connection not found - ErrorCode: %d, Msg: %s", errorResp.ErrorCode, errorResp.Msg))
			return nil, fmt.Errorf("[%s] SFTP connection '%s' not found - ErrorCode: %d, Msg: %s",
				errorCode, connectionIdentifier, errorResp.ErrorCode, errorResp.Msg)
		}
		errorCode := sftpDatasourceErrorCodes.ReadFailed()
		return nil, fmt.Errorf("[%s] SFTP connection '%s' not found", errorCode, connectionIdentifier)
	}

	if err != nil {
		connectionIdentifier := d.GetConnectionIdentifier(connectionName, connectionKey)
		errorCode := sftpDatasourceErrorCodes.ReadFailed()
		opCtx.LogOperationError(logCtx, "SFTP connection API read failed", errorCode, err)
		return nil, fmt.Errorf("[%s] Unable to read SFTP connection '%s': %w", errorCode, connectionIdentifier, err)
	}

	// Check for API errors
	if apiResp != nil && apiResp.SFTPConnectionResponse != nil && apiResp.SFTPConnectionResponse.Errorcode != nil && *apiResp.SFTPConnectionResponse.Errorcode != 0 {
		connectionIdentifier := d.GetConnectionIdentifier(connectionName, connectionKey)
		apiErr := fmt.Errorf("API returned error code %d: %s", *apiResp.SFTPConnectionResponse.Errorcode, errorsutil.SanitizeMessage(apiResp.SFTPConnectionResponse.Msg))
		errorCode := sftpDatasourceErrorCodes.APIError()
		opCtx.LogOperationError(logCtx, "SFTP connection read failed with API error", errorCode, apiErr,
			map[string]interface{}{
				"api_error_code": *apiResp.SFTPConnectionResponse.Errorcode,
				"message":        errorsutil.SanitizeMessage(apiResp.SFTPConnectionResponse.Msg),
			})
		return nil, fmt.Errorf("[%s] API error reading SFTP connection '%s': %w", errorCode, connectionIdentifier, apiErr)
	}

	opCtx.LogOperationEnd(logCtx, "SFTP connection API read completed successfully")
	return apiResp, nil
}

// HandleSFTPAuthenticationLogic handles authentication skip logic for SFTP datasource
// HandleSFTPAuthenticationLogic processes the authenticate flag to control sensitive data visibility
// When authenticate=false, connection_attributes are removed from state to prevent sensitive data exposure
// When authenticate=true, all connection_attributes are returned in state
func (d *SftpConnectionsDataSource) HandleSFTPAuthenticationLogic(state *SFTPConnectionDataSourceModel, resp *datasource.ReadResponse) {
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

// GetConnectionIdentifier returns appropriate connection identifier for logging and errors
func (d *SftpConnectionsDataSource) GetConnectionIdentifier(connectionName string, connectionKey *int64) string {
	if connectionName != "" {
		return connectionName
	}
	if connectionKey != nil {
		return fmt.Sprintf("%d", *connectionKey)
	}
	return "unknown"
}
