// Copyright (c) 2025 Saviynt Inc.
// SPDX-License-Identifier: MPL-2.0

// saviynt_sap_connection_datasource retrieves sap connections details from the Saviynt Security Manager.
// The data source supports a single Read operation to look up an existing sap connections by name.
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

var _ datasource.DataSource = &SapConnectionDataSource{}

// Initialize error codes for SAP Connection datasource operations
var sapDatasourceErrorCodes = errorsutil.NewConnectorErrorCodes(errorsutil.ConnectorTypeSAP)

// SAPConnectionDataSource defines the data source
type SapConnectionDataSource struct {
	client            client.SaviyntClientInterface
	token             string
	provider          client.SaviyntProviderInterface
	connectionFactory client.ConnectionFactoryInterface
}

type SapConnectionDataSourceModel struct {
	ID types.String `tfsdk:"id"`
	BaseConnectionDataSourceModel
	ConnectionAttributes *SapConnectionAttributes `tfsdk:"connection_attributes"`
}

type SapConnectionAttributes struct {
	CreateAccountJson              types.String             `tfsdk:"create_account_json"`
	AuditLogJson                   types.String             `tfsdk:"audit_log_json"`
	ConnectionType                 types.String             `tfsdk:"connection_type"`
	SapTableFilterLang             types.String             `tfsdk:"saptable_filter_lang"`
	PasswordNoOfSplChars           types.String             `tfsdk:"password_noof_spl_chars"`
	TerminatedUserGroup            types.String             `tfsdk:"terminated_user_group"`
	LogsTableFilter                types.String             `tfsdk:"logs_table_filter"`
	EccOrS4Hana                    types.String             `tfsdk:"ecc_or_s4hana"`
	FirefighterIdRevokeAccessJson  types.String             `tfsdk:"firefighterid_revoke_access_json"`
	ConfigJson                     types.String             `tfsdk:"config_json"`
	FirefighterIdGrantAccessJson   types.String             `tfsdk:"firefighterid_grant_access_json"`
	JcoSncLibrary                  types.String             `tfsdk:"jco_snc_library"`
	IsTimeoutSupported             types.Bool               `tfsdk:"is_timeout_supported"`
	JcoR3Name                      types.String             `tfsdk:"jco_r3name"`
	ExternalSodEvalJson            types.String             `tfsdk:"external_sod_eval_json"`
	JcoAshost                      types.String             `tfsdk:"jco_ashost"`
	PasswordNoOfDigits             types.String             `tfsdk:"password_noof_digits"`
	ProvJcoMsHost                  types.String             `tfsdk:"prov_jco_mshost"`
	PamConfig                      types.String             `tfsdk:"pam_config"`
	JcoSncMyName                   types.String             `tfsdk:"jco_snc_myname"`
	EnforcePasswordChange          types.String             `tfsdk:"enforce_password_change"`
	JcoUser                        types.String             `tfsdk:"jco_user"`
	JcoSncMode                     types.String             `tfsdk:"jco_snc_mode"`
	ProvJcoMsServ                  types.String             `tfsdk:"prov_jco_msserv"`
	HanaRefTableJson               types.String             `tfsdk:"hana_ref_table_json"`
	PasswordMinLength              types.String             `tfsdk:"password_min_length"`
	JcoClient                      types.String             `tfsdk:"jco_client"`
	TerminatedUserRoleAction       types.String             `tfsdk:"terminated_user_role_action"`
	ResetPwdForNewAccount          types.String             `tfsdk:"reset_pwd_for_new_account"`
	ProvJcoClient                  types.String             `tfsdk:"prov_jco_client"`
	Snc                            types.String             `tfsdk:"snc"`
	JcoMsServ                      types.String             `tfsdk:"jco_msserv"`
	ProvCuaSnc                     types.String             `tfsdk:"prov_cua_snc"`
	ConnectionTimeoutConfig        *ConnectionTimeoutConfig `tfsdk:"connection_timeout_config"`
	ProvJcoUser                    types.String             `tfsdk:"prov_jco_user"`
	JcoLang                        types.String             `tfsdk:"jco_lang"`
	JcoSncPartnerName              types.String             `tfsdk:"jco_snc_partner_name"`
	StatusThresholdConfig          types.String             `tfsdk:"status_threshold_config"`
	ProvJcoSysNr                   types.String             `tfsdk:"prov_jco_sysnr"`
	SetCuaSystem                   types.String             `tfsdk:"set_cua_system"`
	MessageServer                  types.String             `tfsdk:"message_server"`
	ProvJcoAshost                  types.String             `tfsdk:"prov_jco_ashost"`
	ProvJcoGroup                   types.String             `tfsdk:"prov_jco_group"`
	ProvCuaEnabled                 types.String             `tfsdk:"prov_cua_enabled"`
	JcoMsHost                      types.String             `tfsdk:"jco_mshost"`
	ProvJcoR3Name                  types.String             `tfsdk:"prov_jco_r3name"`
	PasswordNoOfCapsAlpha          types.String             `tfsdk:"password_noof_caps_alpha"`
	ModifyUserDataJson             types.String             `tfsdk:"modify_user_data_json"`
	IsTimeoutConfigValidated       types.Bool               `tfsdk:"is_timeout_config_validated"`
	JcoSncQop                      types.String             `tfsdk:"jco_snc_qop"`
	Tables                         types.String             `tfsdk:"tables"`
	ProvJcoLang                    types.String             `tfsdk:"prov_jco_lang"`
	JcoSysNr                       types.String             `tfsdk:"jco_sysnr"`
	ExternalSodEvalJsonDetail      types.String             `tfsdk:"external_sod_eval_json_detail"`
	DataImportFilter               types.String             `tfsdk:"data_import_filter"`
	EnableAccountJson              types.String             `tfsdk:"enable_account_json"`
	AlternateOutputParameterEtData types.String             `tfsdk:"alternate_output_parameter_et_data"`
	JcoGroup                       types.String             `tfsdk:"jco_group"`
	PasswordMaxLength              types.String             `tfsdk:"password_max_length"`
	UserImportJson                 types.String             `tfsdk:"user_import_json"`
	SystemName                     types.String             `tfsdk:"system_name"`
	UpdateAccountJson              types.String             `tfsdk:"update_account_json"`
}

// NewSAPConnectionsDataSource creates a new SAP connections data source with default factory
func NewSAPConnectionsDataSource() datasource.DataSource {
	return &SapConnectionDataSource{
		connectionFactory: &client.DefaultConnectionFactory{},
	}
}

// NewSAPConnectionsDataSourceWithFactory creates a new SAP connections data source with custom factory
// Used primarily for testing with mock factories
func NewSAPConnectionsDataSourceWithFactory(factory client.ConnectionFactoryInterface) datasource.DataSource {
	return &SapConnectionDataSource{
		connectionFactory: factory,
	}
}

// SetClient sets the client for testing purposes
func (d *SapConnectionDataSource) SetClient(client client.SaviyntClientInterface) {
	d.client = client
}

// SetToken sets the token for testing purposes
func (d *SapConnectionDataSource) SetToken(token string) {
	d.token = token
}

// SetProvider sets the provider for testing purposes
func (d *SapConnectionDataSource) SetProvider(provider client.SaviyntProviderInterface) {
	d.provider = provider
}

// Metadata sets the data source type name for Terraform
func (d *SapConnectionDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "saviynt_sap_connection_datasource"
}

func SapConnectorsDataSourceSchema() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:    true,
			Description: "Resource ID.",
		},
		"connection_attributes": schema.SingleNestedAttribute{
			Computed: true,
			Attributes: map[string]schema.Attribute{
				"create_account_json":                schema.StringAttribute{Computed: true},
				"audit_log_json":                     schema.StringAttribute{Computed: true},
				"connection_type":                    schema.StringAttribute{Computed: true},
				"saptable_filter_lang":               schema.StringAttribute{Computed: true},
				"password_noof_spl_chars":            schema.StringAttribute{Computed: true},
				"terminated_user_group":              schema.StringAttribute{Computed: true},
				"logs_table_filter":                  schema.StringAttribute{Computed: true},
				"ecc_or_s4hana":                      schema.StringAttribute{Computed: true},
				"firefighterid_revoke_access_json":   schema.StringAttribute{Computed: true},
				"config_json":                        schema.StringAttribute{Computed: true},
				"firefighterid_grant_access_json":    schema.StringAttribute{Computed: true},
				"jco_snc_library":                    schema.StringAttribute{Computed: true},
				"is_timeout_supported":               schema.BoolAttribute{Computed: true},
				"jco_r3name":                         schema.StringAttribute{Computed: true},
				"external_sod_eval_json":             schema.StringAttribute{Computed: true},
				"jco_ashost":                         schema.StringAttribute{Computed: true},
				"password_noof_digits":               schema.StringAttribute{Computed: true},
				"prov_jco_mshost":                    schema.StringAttribute{Computed: true},
				"pam_config":                         schema.StringAttribute{Computed: true},
				"jco_snc_myname":                     schema.StringAttribute{Computed: true},
				"enforce_password_change":            schema.StringAttribute{Computed: true},
				"jco_user":                           schema.StringAttribute{Computed: true},
				"jco_snc_mode":                       schema.StringAttribute{Computed: true},
				"prov_jco_msserv":                    schema.StringAttribute{Computed: true},
				"hana_ref_table_json":                schema.StringAttribute{Computed: true},
				"password_min_length":                schema.StringAttribute{Computed: true},
				"jco_client":                         schema.StringAttribute{Computed: true},
				"terminated_user_role_action":        schema.StringAttribute{Computed: true},
				"reset_pwd_for_new_account":          schema.StringAttribute{Computed: true},
				"prov_jco_client":                    schema.StringAttribute{Computed: true},
				"snc":                                schema.StringAttribute{Computed: true},
				"jco_msserv":                         schema.StringAttribute{Computed: true},
				"prov_cua_snc":                       schema.StringAttribute{Computed: true},
				"prov_jco_user":                      schema.StringAttribute{Computed: true},
				"jco_lang":                           schema.StringAttribute{Computed: true},
				"jco_snc_partner_name":               schema.StringAttribute{Computed: true},
				"status_threshold_config":            schema.StringAttribute{Computed: true},
				"prov_jco_sysnr":                     schema.StringAttribute{Computed: true},
				"set_cua_system":                     schema.StringAttribute{Computed: true},
				"message_server":                     schema.StringAttribute{Computed: true},
				"prov_jco_ashost":                    schema.StringAttribute{Computed: true},
				"prov_jco_group":                     schema.StringAttribute{Computed: true},
				"prov_cua_enabled":                   schema.StringAttribute{Computed: true},
				"jco_mshost":                         schema.StringAttribute{Computed: true},
				"prov_jco_r3name":                    schema.StringAttribute{Computed: true},
				"password_noof_caps_alpha":           schema.StringAttribute{Computed: true},
				"modify_user_data_json":              schema.StringAttribute{Computed: true},
				"is_timeout_config_validated":        schema.BoolAttribute{Computed: true},
				"jco_snc_qop":                        schema.StringAttribute{Computed: true},
				"tables":                             schema.StringAttribute{Computed: true},
				"prov_jco_lang":                      schema.StringAttribute{Computed: true},
				"jco_sysnr":                          schema.StringAttribute{Computed: true},
				"external_sod_eval_json_detail":      schema.StringAttribute{Computed: true},
				"data_import_filter":                 schema.StringAttribute{Computed: true},
				"enable_account_json":                schema.StringAttribute{Computed: true},
				"alternate_output_parameter_et_data": schema.StringAttribute{Computed: true},
				"jco_group":                          schema.StringAttribute{Computed: true},
				"password_max_length":                schema.StringAttribute{Computed: true},
				"user_import_json":                   schema.StringAttribute{Computed: true},
				"system_name":                        schema.StringAttribute{Computed: true},
				"update_account_json":                schema.StringAttribute{Computed: true},
				"connection_timeout_config": schema.SingleNestedAttribute{
					Computed:   true,
					Attributes: ConnectionTimeoutConfigeSchema(),
				},
			},
		},
	}
}

// Schema defines the structure and attributes available for the SAP connection data source
func (d *SapConnectionDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: util.SAPConnDataSourceDescription,
		Attributes:  connectionsutil.MergeDataSourceAttributes(BaseConnectorDataSourceSchema(), SapConnectorsDataSourceSchema()),
	}
}

// Configure initializes the data source with provider configuration
// Sets up client and authentication token for API operations
func (d *SapConnectionDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	opCtx := errorsutil.CreateOperationContext(errorsutil.ConnectorTypeSAP, "configure", "")
	ctx = opCtx.AddContextToLogger(ctx)

	opCtx.LogOperationStart(ctx, "Starting SAP connection datasource configuration")

	// Check if provider data is available.
	if req.ProviderData == nil {
		tflog.Debug(ctx, "ProviderData is nil, returning early")
		opCtx.LogOperationEnd(ctx, "SAP connection datasource configuration completed - no provider data")
		return
	}

	// Cast provider data to your provider type.
	prov, ok := req.ProviderData.(*SaviyntProvider)
	if !ok {
		errorCode := sapDatasourceErrorCodes.ProviderConfig()
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

	opCtx.LogOperationEnd(ctx, "SAP connection datasource configured successfully")
}

// Read retrieves SAP connection details from Saviynt and populates the Terraform state
// Supports lookup by connection name or connection key with comprehensive error handling
func (d *SapConnectionDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var state SapConnectionDataSourceModel

	opCtx := errorsutil.CreateOperationContext(errorsutil.ConnectorTypeSAP, "datasource_read", "")
	ctx = opCtx.AddContextToLogger(ctx)

	opCtx.LogOperationStart(ctx, "Starting SAP connection datasource read")

	// Extract configuration from request
	configDiagnostics := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(configDiagnostics...)
	if resp.Diagnostics.HasError() {
		errorCode := sapDatasourceErrorCodes.ConfigExtraction()
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
		errorCode := sapDatasourceErrorCodes.MissingIdentifier()
		opCtx.LogOperationError(ctx, "Missing connection identifier", errorCode,
			fmt.Errorf("either connection_name or connection_key must be provided"))
		resp.Diagnostics.AddError(
			errorsutil.GetErrorMessage(errorsutil.ErrMissingIdentifier),
			fmt.Sprintf("[%s] Either 'connection_name' or 'connection_key' must be provided to look up the SAP connection.", errorCode),
		)
		return
	}

	// Execute API call to get SAP connection details
	apiResp, err := d.ReadSAPConnectionDetails(ctx, connectionName, connectionKey)
	if err != nil {
		// Error is already sanitized in ReadSAPConnectionDetails method
		errorCode := sapDatasourceErrorCodes.ReadFailed()
		opCtx.LogOperationError(ctx, "Failed to read SAP connection details", errorCode, err)
		resp.Diagnostics.AddError(
			errorsutil.GetErrorMessage(errorsutil.ErrReadFailed),
			fmt.Sprintf("[%s] %s", errorCode, err.Error()),
		)
		return
	}

	// Validate API response
	if err := d.ValidateSAPConnectionResponse(apiResp); err != nil {
		errorCode := sapDatasourceErrorCodes.APIError()
		opCtx.LogOperationError(ctx, "Invalid connection type for SAP datasource", errorCode, err)
		resp.Diagnostics.AddError(
			errorsutil.GetErrorMessage(errorsutil.ErrAPIError),
			fmt.Sprintf("[%s] Unable to verify connection type for connection %q. The provider could not determine the type of this connection. Please ensure the connection name is correct and belongs to a supported connector type.", errorCode, connectionName),
		)
		return
	}

	// Map API response to state
	d.UpdateModelFromSAPConnectionResponse(&state, apiResp)

	// Handle authentication logic
	d.HandleSAPAuthenticationLogic(&state, resp)

	// Set final state
	stateDiagnostics := resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(stateDiagnostics...)
	if resp.Diagnostics.HasError() {
		errorCode := sapDatasourceErrorCodes.StateUpdate()
		opCtx.LogOperationError(ctx, "Failed to set state", errorCode,
			fmt.Errorf("state update failed"))
		resp.Diagnostics.AddError(
			errorsutil.GetErrorMessage(errorsutil.ErrStateUpdate),
			fmt.Sprintf("[%s] Unable to update Terraform state for SAP connection datasource", errorCode),
		)
		return
	}

	opCtx.LogOperationEnd(ctx, "SAP connection datasource read completed successfully",
		map[string]interface{}{
			"connection_name": connectionName,
			"has_attributes":  state.ConnectionAttributes != nil,
		})
}
// ReadSAPConnectionDetails retrieves SAP connection details from Saviynt API
// Handles both connection name and connection key based lookups using factory pattern
// Returns standardized errors with proper correlation tracking and sensitive data sanitization
func (d *SapConnectionDataSource) ReadSAPConnectionDetails(ctx context.Context, connectionName string, connectionKey *int64) (*openapi.GetConnectionDetailsResponse, error) {
	opCtx := errorsutil.CreateOperationContext(errorsutil.ConnectorTypeSAP, "api_read", connectionName)
	logCtx := opCtx.AddContextToLogger(ctx)

	opCtx.LogOperationStart(logCtx, "Starting SAP connection API call")

	tflog.Debug(logCtx, "Executing API request to get SAP connection details")

	var apiResp *openapi.GetConnectionDetailsResponse

	// Execute API request with retry logic
	err := d.provider.AuthenticatedAPICallWithRetry(ctx, "read_sap_connection_datasource", func(token string) error {
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
		errorCode := sapDatasourceErrorCodes.ReadFailed()
		opCtx.LogOperationError(logCtx, "Failed to read SAP connection details", errorCode, err)
		return nil, errorsutil.CreateStandardError(errorsutil.ConnectorTypeSAP, errorCode, "api_read", connectionName, err)
	}

	opCtx.LogOperationEnd(logCtx, "SAP connection API call completed successfully")

	return apiResp, nil
}

// ValidateSAPConnectionResponse validates that the API response contains valid SAP connection data
// Returns standardized error if validation fails
func (d *SapConnectionDataSource) ValidateSAPConnectionResponse(apiResp *openapi.GetConnectionDetailsResponse) error {
	if apiResp != nil && apiResp.SAPConnectionResponse == nil {
		return fmt.Errorf("verify the connection type - SAP connection response is nil")
	}
	return nil
}

// UpdateModelFromSAPConnectionResponse maps API response data to the Terraform state model
// It handles both base connection fields and detailed connection attributes
func (d *SapConnectionDataSource) UpdateModelFromSAPConnectionResponse(state *SapConnectionDataSourceModel, apiResp *openapi.GetConnectionDetailsResponse) {
	// Map base connection fields
	d.MapBaseSAPConnectionFields(state, apiResp)

	// Map connection attributes
	d.MapSAPConnectionAttributes(state, apiResp)
}

// MapBaseSAPConnectionFields maps basic connection fields from API response to state model
// These are common fields available for all connection types
func (d *SapConnectionDataSource) MapBaseSAPConnectionFields(state *SapConnectionDataSourceModel, apiResp *openapi.GetConnectionDetailsResponse) {
	state.Msg = util.SafeStringDatasource(apiResp.SAPConnectionResponse.Msg)
	state.ErrorCode = util.SafeInt64(apiResp.SAPConnectionResponse.Errorcode)
	connectionKey := util.SafeInt64(apiResp.SAPConnectionResponse.Connectionkey)
	state.ID = types.StringValue(fmt.Sprintf("ds-sap-%d", connectionKey.ValueInt64()))
	state.ConnectionName = util.SafeStringDatasource(apiResp.SAPConnectionResponse.Connectionname)
	state.ConnectionKey = util.SafeInt64(apiResp.SAPConnectionResponse.Connectionkey)
	state.Description = util.SafeStringDatasource(apiResp.SAPConnectionResponse.Description)
	state.DefaultSavRoles = util.SafeStringDatasource(apiResp.SAPConnectionResponse.Defaultsavroles)
	state.ConnectionType = util.SafeStringDatasource(apiResp.SAPConnectionResponse.Connectiontype)
	state.CreatedOn = util.SafeStringDatasource(apiResp.SAPConnectionResponse.Createdon)
	state.CreatedBy = util.SafeStringDatasource(apiResp.SAPConnectionResponse.Createdby)
	state.UpdatedBy = util.SafeStringDatasource(apiResp.SAPConnectionResponse.Updatedby)
	state.EmailTemplate = util.SafeStringDatasource(apiResp.SAPConnectionResponse.Emailtemplate)
}

// MapSAPConnectionAttributes maps detailed SAP connection attributes from API response to state model
// Returns nil if no connection attributes are present in the response
func (d *SapConnectionDataSource) MapSAPConnectionAttributes(state *SapConnectionDataSourceModel, apiResp *openapi.GetConnectionDetailsResponse) {
	if apiResp.SAPConnectionResponse.Connectionattributes == nil {
		state.ConnectionAttributes = nil
		return
	}

	attrs := apiResp.SAPConnectionResponse.Connectionattributes
	state.ConnectionAttributes = &SapConnectionAttributes{
		CreateAccountJson:              util.SafeStringDatasource(attrs.CREATEACCOUNTJSON),
		AuditLogJson:                   util.SafeStringDatasource(attrs.AUDIT_LOG_JSON),
		ConnectionType:                 util.SafeStringDatasource(attrs.ConnectionType),
		SapTableFilterLang:             util.SafeStringDatasource(attrs.SAPTABLE_FILTER_LANG),
		PasswordNoOfSplChars:           util.SafeStringDatasource(attrs.PASSWORD_NOOFSPLCHARS),
		TerminatedUserGroup:            util.SafeStringDatasource(attrs.TERMINATEDUSERGROUP),
		LogsTableFilter:                util.SafeStringDatasource(attrs.LOGS_TABLE_FILTER),
		EccOrS4Hana:                    util.SafeStringDatasource(attrs.ECCORS4HANA),
		FirefighterIdRevokeAccessJson:  util.SafeStringDatasource(attrs.FIREFIGHTERID_REVOKE_ACCESS_JSON),
		ConfigJson:                     util.SafeStringDatasource(attrs.ConfigJSON),
		FirefighterIdGrantAccessJson:   util.SafeStringDatasource(attrs.FIREFIGHTERID_GRANT_ACCESS_JSON),
		JcoSncLibrary:                  util.SafeStringDatasource(attrs.JCO_SNC_LIBRARY),
		IsTimeoutSupported:             util.SafeBoolDatasource(attrs.IsTimeoutSupported),
		JcoR3Name:                      util.SafeStringDatasource(attrs.JCOR3NAME),
		ExternalSodEvalJson:            util.SafeStringDatasource(attrs.EXTERNAL_SOD_EVAL_JSON),
		JcoAshost:                      util.SafeStringDatasource(attrs.JCO_ASHOST),
		PasswordNoOfDigits:             util.SafeStringDatasource(attrs.PASSWORD_NOOFDIGITS),
		ProvJcoMsHost:                  util.SafeStringDatasource(attrs.PROV_JCO_MSHOST),
		PamConfig:                      util.SafeStringDatasource(attrs.PAM_CONFIG),
		JcoSncMyName:                   util.SafeStringDatasource(attrs.JCO_SNC_MYNAME),
		EnforcePasswordChange:          util.SafeStringDatasource(attrs.ENFORCEPASSWORDCHANGE),
		JcoUser:                        util.SafeStringDatasource(attrs.JCO_USER),
		JcoSncMode:                     util.SafeStringDatasource(attrs.JCO_SNC_MODE),
		ProvJcoMsServ:                  util.SafeStringDatasource(attrs.PROV_JCO_MSSERV),
		HanaRefTableJson:               util.SafeStringDatasource(attrs.HANAREFTABLEJSON),
		PasswordMinLength:              util.SafeStringDatasource(attrs.PASSWORD_MIN_LENGTH),
		JcoClient:                      util.SafeStringDatasource(attrs.JCO_CLIENT),
		TerminatedUserRoleAction:       util.SafeStringDatasource(attrs.TERMINATED_USER_ROLE_ACTION),
		ResetPwdForNewAccount:          util.SafeStringDatasource(attrs.RESET_PWD_FOR_NEWACCOUNT),
		ProvJcoClient:                  util.SafeStringDatasource(attrs.PROV_JCO_CLIENT),
		Snc:                            util.SafeStringDatasource(attrs.SNC),
		JcoMsServ:                      util.SafeStringDatasource(attrs.JCO_MSSERV),
		ProvCuaSnc:                     util.SafeStringDatasource(attrs.PROV_CUA_SNC),
		ProvJcoUser:                    util.SafeStringDatasource(attrs.PROV_JCO_USER),
		JcoLang:                        util.SafeStringDatasource(attrs.JCO_LANG),
		JcoSncPartnerName:              util.SafeStringDatasource(attrs.JCO_SNC_PARTNERNAME),
		StatusThresholdConfig:          util.SafeStringDatasource(attrs.STATUS_THRESHOLD_CONFIG),
		ProvJcoSysNr:                   util.SafeStringDatasource(attrs.PROV_JCO_SYSNR),
		SetCuaSystem:                   util.SafeStringDatasource(attrs.SETCUASYSTEM),
		MessageServer:                  util.SafeStringDatasource(attrs.MESSAGESERVER),
		ProvJcoAshost:                  util.SafeStringDatasource(attrs.PROV_JCO_ASHOST),
		ProvJcoGroup:                   util.SafeStringDatasource(attrs.PROV_JCO_GROUP),
		ProvCuaEnabled:                 util.SafeStringDatasource(attrs.PROV_CUA_ENABLED),
		JcoMsHost:                      util.SafeStringDatasource(attrs.JCO_MSHOST),
		ProvJcoR3Name:                  util.SafeStringDatasource(attrs.PROVJCOR3NAME),
		PasswordNoOfCapsAlpha:          util.SafeStringDatasource(attrs.PASSWORD_NOOFCAPSALPHA),
		ModifyUserDataJson:             util.SafeStringDatasource(attrs.MODIFYUSERDATAJSON),
		IsTimeoutConfigValidated:       util.SafeBoolDatasource(attrs.IsTimeoutConfigValidated),
		JcoSncQop:                      util.SafeStringDatasource(attrs.JCO_SNC_QOP),
		Tables:                         util.SafeStringDatasource(attrs.TABLES),
		ProvJcoLang:                    util.SafeStringDatasource(attrs.PROV_JCO_LANG),
		JcoSysNr:                       util.SafeStringDatasource(attrs.JCO_SYSNR),
		ExternalSodEvalJsonDetail:      util.SafeStringDatasource(attrs.EXTERNAL_SOD_EVAL_JSON_DETAIL),
		DataImportFilter:               util.SafeStringDatasource(attrs.DATA_IMPORT_FILTER),
		EnableAccountJson:              util.SafeStringDatasource(attrs.ENABLEACCOUNTJSON),
		AlternateOutputParameterEtData: util.SafeStringDatasource(attrs.ALTERNATE_OUTPUT_PARAMETER_ET_DATA),
		JcoGroup:                       util.SafeStringDatasource(attrs.JCO_GROUP),
		PasswordMaxLength:              util.SafeStringDatasource(attrs.PASSWORD_MAX_LENGTH),
		UserImportJson:                 util.SafeStringDatasource(attrs.USERIMPORTJSON),
		SystemName:                     util.SafeStringDatasource(attrs.SYSTEMNAME),
		UpdateAccountJson:              util.SafeStringDatasource(attrs.UPDATEACCOUNTJSON),
	}

	// Map connection timeout config if present
	d.MapSAPTimeoutConfig(attrs, state.ConnectionAttributes)
}

// MapSAPTimeoutConfig maps connection timeout configuration from API response to state model
// Only maps timeout config if it exists in the API response
func (d *SapConnectionDataSource) MapSAPTimeoutConfig(attrs *openapi.SAPConnectionAttributes, connectionAttrs *SapConnectionAttributes) {
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

// HandleSAPAuthenticationLogic processes the authenticate flag to control sensitive data visibility
// When authenticate=false, connection_attributes are removed from state to prevent sensitive data exposure
// When authenticate=true, all connection_attributes are returned in state
func (d *SapConnectionDataSource) HandleSAPAuthenticationLogic(state *SapConnectionDataSourceModel, resp *datasource.ReadResponse) {
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
