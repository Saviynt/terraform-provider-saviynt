// Copyright (c) 2025 Saviynt Inc.
// SPDX-License-Identifier: MPL-2.0

// saviynt_salesforce_connection_datasource retrieves salesforce connections details from the Saviynt Security Manager.
// The data source supports a single Read operation to look up an existing salesforce connections by name.
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

var _ datasource.DataSource = &SalesforceConnectionDataSource{}

// Initialize error codes for Salesforce Connection datasource operations
var salesforceDatasourceErrorCodes = errorsutil.NewConnectorErrorCodes(errorsutil.ConnectorTypeSalesforce)

// SalesforceConnectionDataSource defines the data source
type SalesforceConnectionDataSource struct {
	client            client.SaviyntClientInterface
	token             string
	provider          client.SaviyntProviderInterface
	connectionFactory client.ConnectionFactoryInterface
}

type SalesforceConnectionDataSourceModel struct {
	ID types.String `tfsdk:"id"`
	BaseConnectionDataSourceModel
	ConnectionAttributes *SalesforceConnectionAttributes `tfsdk:"connection_attributes"`
}

type SalesforceConnectionAttributes struct {
	IsTimeoutSupported       types.Bool               `tfsdk:"is_timeout_supported"`
	ObjectToBeImported       types.String             `tfsdk:"object_to_be_imported"`
	FeatureLicenseJson       types.String             `tfsdk:"feature_license_json"`
	CreateAccountJson        types.String             `tfsdk:"createaccountjson"`
	RedirectUri              types.String             `tfsdk:"redirect_uri"`
	ConnectionTimeoutConfig  *ConnectionTimeoutConfig `tfsdk:"connection_timeout_config"`
	ModifyAccountJson        types.String             `tfsdk:"modifyaccountjson"`
	ConnectionType           types.String             `tfsdk:"connection_type"`
	IsTimeoutConfigValidated types.Bool               `tfsdk:"is_timeout_config_validated"`
	PamConfig                types.String             `tfsdk:"pam_config"`
	CustomConfigJson         types.String             `tfsdk:"customconfigjson"`
	FieldMappingJson         types.String             `tfsdk:"field_mapping_json"`
	StatusThresholdConfig    types.String             `tfsdk:"status_threshold_config"`
	AccountFieldQuery        types.String             `tfsdk:"account_field_query"`
	CustomCreateAccountUrl   types.String             `tfsdk:"custom_createaccount_url"`
	AccountFilterQuery       types.String             `tfsdk:"account_filter_query"`
	InstanceUrl              types.String             `tfsdk:"instance_url"`
}

// NewSalesforceConnectionsDataSource creates a new Salesforce connections data source with default factory
func NewSalesforceConnectionsDataSource() datasource.DataSource {
	return &SalesforceConnectionDataSource{
		connectionFactory: &client.DefaultConnectionFactory{},
	}
}

// NewSalesforceConnectionsDataSourceWithFactory creates a new Salesforce connections data source with custom factory
// Used primarily for testing with mock factories
func NewSalesforceConnectionsDataSourceWithFactory(factory client.ConnectionFactoryInterface) datasource.DataSource {
	return &SalesforceConnectionDataSource{
		connectionFactory: factory,
	}
}

// SetClient sets the client for testing purposes
func (d *SalesforceConnectionDataSource) SetClient(client client.SaviyntClientInterface) {
	d.client = client
}

// SetToken sets the token for testing purposes
func (d *SalesforceConnectionDataSource) SetToken(token string) {
	d.token = token
}

// SetProvider sets the provider for testing purposes
func (d *SalesforceConnectionDataSource) SetProvider(provider client.SaviyntProviderInterface) {
	d.provider = provider
}

func (d *SalesforceConnectionDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "saviynt_salesforce_connection_datasource"
}

func SalesforceConnectorsDataSourceSchema() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:    true,
			Description: "Resource ID.",
		},
		"connection_attributes": schema.SingleNestedAttribute{
			Computed: true,
			Attributes: map[string]schema.Attribute{
				"is_timeout_supported":        schema.BoolAttribute{Computed: true},
				"object_to_be_imported":       schema.StringAttribute{Computed: true},
				"feature_license_json":        schema.StringAttribute{Computed: true},
				"createaccountjson":           schema.StringAttribute{Computed: true},
				"redirect_uri":                schema.StringAttribute{Computed: true},
				"modifyaccountjson":           schema.StringAttribute{Computed: true},
				"connection_type":             schema.StringAttribute{Computed: true},
				"is_timeout_config_validated": schema.BoolAttribute{Computed: true},
				"pam_config":                  schema.StringAttribute{Computed: true},
				"customconfigjson":            schema.StringAttribute{Computed: true},
				"field_mapping_json":          schema.StringAttribute{Computed: true},
				"status_threshold_config":     schema.StringAttribute{Computed: true},
				"account_field_query":         schema.StringAttribute{Computed: true},
				"custom_createaccount_url":    schema.StringAttribute{Computed: true},
				"account_filter_query":        schema.StringAttribute{Computed: true},
				"instance_url":                schema.StringAttribute{Computed: true},
				"connection_timeout_config": schema.SingleNestedAttribute{
					Computed:   true,
					Attributes: ConnectionTimeoutConfigeSchema(),
				},
			},
		},
	}
}

// Schema defines the attributes for the data source
func (d *SalesforceConnectionDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: util.SalesforceConnDataSourceDescription,
		Attributes:  connectionsutil.MergeDataSourceAttributes(BaseConnectorDataSourceSchema(), SalesforceConnectorsDataSourceSchema()),
	}
}

// Configure initializes the data source with provider configuration
// Sets up client and authentication token for API operations
func (d *SalesforceConnectionDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	opCtx := errorsutil.CreateOperationContext(errorsutil.ConnectorTypeSalesforce, "configure", "")
	ctx = opCtx.AddContextToLogger(ctx)

	opCtx.LogOperationStart(ctx, "Starting Salesforce connection datasource configuration")

	// Check if provider data is available.
	if req.ProviderData == nil {
		tflog.Debug(ctx, "ProviderData is nil, returning early")
		opCtx.LogOperationEnd(ctx, "Salesforce connection datasource configuration completed - no provider data")
		return
	}

	// Cast provider data to your provider type.
	prov, ok := req.ProviderData.(*SaviyntProvider)
	if !ok {
		errorCode := salesforceDatasourceErrorCodes.ProviderConfig()
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

	opCtx.LogOperationEnd(ctx, "Salesforce connection datasource configured successfully")
}

// Read retrieves Salesforce connection details from Saviynt and populates the Terraform state
// Supports lookup by connection name or connection key with comprehensive error handling
func (d *SalesforceConnectionDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var state SalesforceConnectionDataSourceModel

	opCtx := errorsutil.CreateOperationContext(errorsutil.ConnectorTypeSalesforce, "datasource_read", "")
	ctx = opCtx.AddContextToLogger(ctx)

	opCtx.LogOperationStart(ctx, "Starting Salesforce connection datasource read")

	// Extract configuration from request
	configDiagnostics := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(configDiagnostics...)
	if resp.Diagnostics.HasError() {
		errorCode := salesforceDatasourceErrorCodes.ConfigExtraction()
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
		errorCode := salesforceDatasourceErrorCodes.MissingIdentifier()
		opCtx.LogOperationError(ctx, "Missing connection identifier", errorCode,
			fmt.Errorf("either connection_name or connection_key must be provided"))
		resp.Diagnostics.AddError(
			errorsutil.GetErrorMessage(errorsutil.ErrMissingIdentifier),
			fmt.Sprintf("[%s] Either 'connection_name' or 'connection_key' must be provided to look up the Salesforce connection.", errorCode),
		)
		return
	}

	// Execute API call to get Salesforce connection details
	apiResp, err := d.ReadSalesforceConnectionDetails(ctx, connectionName, connectionKey)
	if err != nil {
		// Error is already sanitized in ReadSalesforceConnectionDetails method
		errorCode := salesforceDatasourceErrorCodes.ReadFailed()
		opCtx.LogOperationError(ctx, "Failed to read Salesforce connection details", errorCode, err)
		resp.Diagnostics.AddError(
			errorsutil.GetErrorMessage(errorsutil.ErrReadFailed),
			fmt.Sprintf("[%s] %s", errorCode, err.Error()),
		)
		return
	}

	// Validate API response
	if err := d.ValidateSalesforceConnectionResponse(apiResp); err != nil {
		errorCode := salesforceDatasourceErrorCodes.APIError()
		opCtx.LogOperationError(ctx, "Invalid connection type for Salesforce datasource", errorCode, err)
		resp.Diagnostics.AddError(
			errorsutil.GetErrorMessage(errorsutil.ErrAPIError),
			fmt.Sprintf("[%s] Unable to verify connection type for connection %q. The provider could not determine the type of this connection. Please ensure the connection name is correct and belongs to a supported connector type.", errorCode, connectionName),
		)
		return
	}

	// Map API response to state
	d.UpdateModelFromSalesforceConnectionResponse(&state, apiResp)

	// Handle authentication logic
	d.HandleSalesforceAuthenticationLogic(&state, resp)

	// Set final state
	stateDiagnostics := resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(stateDiagnostics...)
	if resp.Diagnostics.HasError() {
		errorCode := salesforceDatasourceErrorCodes.StateUpdate()
		opCtx.LogOperationError(ctx, "Failed to set state", errorCode,
			fmt.Errorf("state update failed"))
		resp.Diagnostics.AddError(
			errorsutil.GetErrorMessage(errorsutil.ErrStateUpdate),
			fmt.Sprintf("[%s] Unable to update Terraform state for Salesforce connection datasource", errorCode),
		)
		return
	}

	opCtx.LogOperationEnd(ctx, "Salesforce connection datasource read completed successfully",
		map[string]interface{}{
			"connection_name": connectionName,
			"has_attributes":  state.ConnectionAttributes != nil,
		})
}
// ReadSalesforceConnectionDetails retrieves Salesforce connection details from Saviynt API
// Handles both connection name and connection key based lookups using factory pattern
// Returns standardized errors with proper correlation tracking and sensitive data sanitization
func (d *SalesforceConnectionDataSource) ReadSalesforceConnectionDetails(ctx context.Context, connectionName string, connectionKey *int64) (*openapi.GetConnectionDetailsResponse, error) {
	opCtx := errorsutil.CreateOperationContext(errorsutil.ConnectorTypeSalesforce, "api_read", connectionName)
	logCtx := opCtx.AddContextToLogger(ctx)

	opCtx.LogOperationStart(logCtx, "Starting Salesforce connection API call")

	tflog.Debug(logCtx, "Executing API request to get Salesforce connection details")

	var apiResp *openapi.GetConnectionDetailsResponse

	// Execute API request with retry logic
	err := d.provider.AuthenticatedAPICallWithRetry(ctx, "read_salesforce_connection_datasource", func(token string) error {
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
		errorCode := salesforceDatasourceErrorCodes.ReadFailed()
		opCtx.LogOperationError(logCtx, "Failed to read Salesforce connection details", errorCode, err)
		return nil, errorsutil.CreateStandardError(errorsutil.ConnectorTypeSalesforce, errorCode, "api_read", connectionName, err)
	}

	opCtx.LogOperationEnd(logCtx, "Salesforce connection API call completed successfully")

	return apiResp, nil
}

// ValidateSalesforceConnectionResponse validates that the API response contains valid Salesforce connection data
// Returns standardized error if validation fails
func (d *SalesforceConnectionDataSource) ValidateSalesforceConnectionResponse(apiResp *openapi.GetConnectionDetailsResponse) error {
	if apiResp != nil && apiResp.SalesforceConnectionResponse == nil {
		return fmt.Errorf("verify the connection type - Salesforce connection response is nil")
	}
	return nil
}

// UpdateModelFromSalesforceConnectionResponse maps API response data to the Terraform state model
// It handles both base connection fields and detailed connection attributes
func (d *SalesforceConnectionDataSource) UpdateModelFromSalesforceConnectionResponse(state *SalesforceConnectionDataSourceModel, apiResp *openapi.GetConnectionDetailsResponse) {
	// Map base connection fields
	d.MapBaseSalesforceConnectionFields(state, apiResp)

	// Map connection attributes
	d.MapSalesforceConnectionAttributes(state, apiResp)
}

// MapBaseSalesforceConnectionFields maps basic connection fields from API response to state model
// These are common fields available for all connection types
func (d *SalesforceConnectionDataSource) MapBaseSalesforceConnectionFields(state *SalesforceConnectionDataSourceModel, apiResp *openapi.GetConnectionDetailsResponse) {
	state.Msg = util.SafeStringDatasource(apiResp.SalesforceConnectionResponse.Msg)
	state.ErrorCode = util.SafeInt64(apiResp.SalesforceConnectionResponse.Errorcode)
	connectionKey := util.SafeInt64(apiResp.SalesforceConnectionResponse.Connectionkey)
	state.ID = types.StringValue(fmt.Sprintf("ds-salesforce-%d", connectionKey.ValueInt64()))
	state.ConnectionName = util.SafeStringDatasource(apiResp.SalesforceConnectionResponse.Connectionname)
	state.ConnectionKey = util.SafeInt64(apiResp.SalesforceConnectionResponse.Connectionkey)
	state.Description = util.SafeStringDatasource(apiResp.SalesforceConnectionResponse.Description)
	state.DefaultSavRoles = util.SafeStringDatasource(apiResp.SalesforceConnectionResponse.Defaultsavroles)
	state.ConnectionType = util.SafeStringDatasource(apiResp.SalesforceConnectionResponse.Connectiontype)
	state.CreatedOn = util.SafeStringDatasource(apiResp.SalesforceConnectionResponse.Createdon)
	state.CreatedBy = util.SafeStringDatasource(apiResp.SalesforceConnectionResponse.Createdby)
	state.UpdatedBy = util.SafeStringDatasource(apiResp.SalesforceConnectionResponse.Updatedby)
	state.EmailTemplate = util.SafeStringDatasource(apiResp.SalesforceConnectionResponse.Emailtemplate)
}

// MapSalesforceConnectionAttributes maps detailed Salesforce connection attributes from API response to state model
// Returns nil if no connection attributes are present in the response
func (d *SalesforceConnectionDataSource) MapSalesforceConnectionAttributes(state *SalesforceConnectionDataSourceModel, apiResp *openapi.GetConnectionDetailsResponse) {
	if apiResp.SalesforceConnectionResponse.Connectionattributes == nil {
		state.ConnectionAttributes = nil
		return
	}

	attrs := apiResp.SalesforceConnectionResponse.Connectionattributes
	state.ConnectionAttributes = &SalesforceConnectionAttributes{
		IsTimeoutSupported:       util.SafeBoolDatasource(attrs.IsTimeoutSupported),
		ObjectToBeImported:       util.SafeStringDatasource(attrs.OBJECT_TO_BE_IMPORTED),
		FeatureLicenseJson:       util.SafeStringDatasource(attrs.FEATURE_LICENSE_JSON),
		CreateAccountJson:        util.SafeStringDatasource(attrs.CREATEACCOUNTJSON),
		RedirectUri:              util.SafeStringDatasource(attrs.REDIRECT_URI),
		ConnectionType:           util.SafeStringDatasource(attrs.ConnectionType),
		ModifyAccountJson:        util.SafeStringDatasource(attrs.MODIFYACCOUNTJSON),
		IsTimeoutConfigValidated: util.SafeBoolDatasource(attrs.IsTimeoutConfigValidated),
		PamConfig:                util.SafeStringDatasource(attrs.PAM_CONFIG),
		CustomConfigJson:         util.SafeStringDatasource(attrs.CUSTOMCONFIGJSON),
		FieldMappingJson:         util.SafeStringDatasource(attrs.FIELD_MAPPING_JSON),
		StatusThresholdConfig:    util.SafeStringDatasource(attrs.STATUS_THRESHOLD_CONFIG),
		AccountFieldQuery:        util.SafeStringDatasource(attrs.ACCOUNT_FIELD_QUERY),
		CustomCreateAccountUrl:   util.SafeStringDatasource(attrs.CUSTOM_CREATEACCOUNT_URL),
		AccountFilterQuery:       util.SafeStringDatasource(attrs.ACCOUNT_FILTER_QUERY),
		InstanceUrl:              util.SafeStringDatasource(attrs.INSTANCE_URL),
	}

	// Map connection timeout config if present
	d.MapSalesforceTimeoutConfig(attrs, state.ConnectionAttributes)
}

// MapSalesforceTimeoutConfig maps connection timeout configuration from API response to state model
// Only maps timeout config if it exists in the API response
func (d *SalesforceConnectionDataSource) MapSalesforceTimeoutConfig(attrs *openapi.SalesforceConnectionAttributes, connectionAttrs *SalesforceConnectionAttributes) {
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

// HandleSalesforceAuthenticationLogic processes the authenticate flag to control sensitive data visibility
// When authenticate=false, connection_attributes are removed from state to prevent sensitive data exposure
// When authenticate=true, all connection_attributes are returned in state
func (d *SalesforceConnectionDataSource) HandleSalesforceAuthenticationLogic(state *SalesforceConnectionDataSourceModel, resp *datasource.ReadResponse) {
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
