// Copyright (c) 2025 Saviynt Inc.
// SPDX-License-Identifier: MPL-2.0

// saviynt_sap_connection_datasource retrieves sap connections details from the Saviynt Security Manager.
// The data source supports a single Read operation to look up an existing sap connections by name.
package provider

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"terraform-provider-Saviynt/util"
	connectionsutil "terraform-provider-Saviynt/util/connectionsutil"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	s "github.com/saviynt/saviynt-api-go-client"

	openapi "github.com/saviynt/saviynt-api-go-client/connections"
)

var _ datasource.DataSource = &sapConnectionDataSource{}

// SAPConnectionDataSource defines the data source
type sapConnectionDataSource struct {
	client *s.Client
	token  string
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

func NewSAPConnectionsDataSource() datasource.DataSource {
	return &sapConnectionDataSource{}
}

func (d *sapConnectionDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
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

// Schema defines the attributes for the data source
func (d *sapConnectionDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: util.SAPConnDataSourceDescription,
		Attributes:  connectionsutil.MergeDataSourceAttributes(BaseConnectorDataSourceSchema(), SapConnectorsDataSourceSchema()),
	}
}

func (d *sapConnectionDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	// Check if provider data is available.
	if req.ProviderData == nil {
		log.Println("ProviderData is nil, returning early.")
		return
	}

	// Cast provider data to your provider type.
	prov, ok := req.ProviderData.(*saviyntProvider)
	if !ok {
		resp.Diagnostics.AddError("Unexpected Provider Data", "Expected *saviyntProvider")
		return
	}

	// Set the client and token from the provider state.
	d.client = prov.client
	d.token = prov.accessToken
}

func (d *sapConnectionDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var state SapConnectionDataSourceModel

	configDiagnostics := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(configDiagnostics...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Configure API client
	cfg := openapi.NewConfiguration()
	apiBaseURL := strings.TrimPrefix(strings.TrimPrefix(d.client.APIBaseURL(), "https://"), "http://")
	cfg.Host = apiBaseURL
	cfg.Scheme = "https"
	cfg.AddDefaultHeader("Authorization", "Bearer "+d.token)
	cfg.HTTPClient = http.DefaultClient

	apiClient := openapi.NewAPIClient(cfg)
	reqParams := openapi.GetConnectionDetailsRequest{}

	// Set filters based on provided parameters
	if !state.ConnectionName.IsNull() && state.ConnectionName.ValueString() != "" {
		reqParams.SetConnectionname(state.ConnectionName.ValueString())
	}
	if !state.ConnectionKey.IsNull() {
		connectionKeyInt := state.ConnectionKey.ValueInt64()
		reqParams.SetConnectionkey(strconv.FormatInt(connectionKeyInt, 10))
	}
	apiReq := apiClient.ConnectionsAPI.GetConnectionDetails(ctx).GetConnectionDetailsRequest(reqParams)

	// Execute API request
	apiResp, httpResp, err := apiReq.Execute()
	if err != nil {
		if httpResp != nil && httpResp.StatusCode != 200 {
			log.Printf("[ERROR] HTTP error while creating Sap Connection: %s", httpResp.Status)
			var fetchResp map[string]interface{}
			if err := json.NewDecoder(httpResp.Body).Decode(&fetchResp); err != nil {
				resp.Diagnostics.AddError("Failed to decode error response", err.Error())
				return
			}
			resp.Diagnostics.AddError(
				"HTTP Error",
				fmt.Sprintf("HTTP error while creating Sap Connection for the reasons: %s", fetchResp["msg"]),
			)

		} else {
			log.Printf("[ERROR] API Call Failed: %v", err)
			resp.Diagnostics.AddError("API Call Failed", fmt.Sprintf("Error: %v", err))
		}
		return
	}
	if apiResp != nil && apiResp.SAPConnectionResponse == nil {
		error := "Verify the connection type"
		log.Printf("[ERROR]: Verify the connection type given")
		resp.Diagnostics.AddError("Read of Sap connection failed", error)
		return
	}
	log.Printf("[DEBUG] HTTP Status Code: %d", httpResp.StatusCode)

	state.Msg = util.SafeStringDatasource(apiResp.SAPConnectionResponse.Msg)
	state.ErrorCode = util.SafeInt64(apiResp.SAPConnectionResponse.Errorcode)
	state.ConnectionName = util.SafeStringDatasource(apiResp.SAPConnectionResponse.Connectionname)
	state.ConnectionKey = util.SafeInt64(apiResp.SAPConnectionResponse.Connectionkey)
	state.Description = util.SafeStringDatasource(apiResp.SAPConnectionResponse.Description)
	state.DefaultSavRoles = util.SafeStringDatasource(apiResp.SAPConnectionResponse.Defaultsavroles)
	state.ConnectionType = util.SafeStringDatasource(apiResp.SAPConnectionResponse.Connectiontype)
	state.CreatedOn = util.SafeStringDatasource(apiResp.SAPConnectionResponse.Createdon)
	state.CreatedBy = util.SafeStringDatasource(apiResp.SAPConnectionResponse.Createdby)
	state.UpdatedBy = util.SafeStringDatasource(apiResp.SAPConnectionResponse.Updatedby)
	state.EmailTemplate = util.SafeStringDatasource(apiResp.SAPConnectionResponse.Emailtemplate)

	if apiResp.SAPConnectionResponse.Connectionattributes != nil {
		state.ConnectionAttributes = &SapConnectionAttributes{
			CreateAccountJson:              util.SafeStringDatasource(apiResp.SAPConnectionResponse.Connectionattributes.CREATEACCOUNTJSON),
			AuditLogJson:                   util.SafeStringDatasource(apiResp.SAPConnectionResponse.Connectionattributes.AUDIT_LOG_JSON),
			ConnectionType:                 util.SafeStringDatasource(apiResp.SAPConnectionResponse.Connectionattributes.ConnectionType),
			SapTableFilterLang:             util.SafeStringDatasource(apiResp.SAPConnectionResponse.Connectionattributes.SAPTABLE_FILTER_LANG),
			PasswordNoOfSplChars:           util.SafeStringDatasource(apiResp.SAPConnectionResponse.Connectionattributes.PASSWORD_NOOFSPLCHARS),
			TerminatedUserGroup:            util.SafeStringDatasource(apiResp.SAPConnectionResponse.Connectionattributes.TERMINATEDUSERGROUP),
			LogsTableFilter:                util.SafeStringDatasource(apiResp.SAPConnectionResponse.Connectionattributes.LOGS_TABLE_FILTER),
			EccOrS4Hana:                    util.SafeStringDatasource(apiResp.SAPConnectionResponse.Connectionattributes.ECCORS4HANA),
			FirefighterIdRevokeAccessJson:  util.SafeStringDatasource(apiResp.SAPConnectionResponse.Connectionattributes.FIREFIGHTERID_REVOKE_ACCESS_JSON),
			ConfigJson:                     util.SafeStringDatasource(apiResp.SAPConnectionResponse.Connectionattributes.ConfigJSON),
			FirefighterIdGrantAccessJson:   util.SafeStringDatasource(apiResp.SAPConnectionResponse.Connectionattributes.FIREFIGHTERID_GRANT_ACCESS_JSON),
			JcoSncLibrary:                  util.SafeStringDatasource(apiResp.SAPConnectionResponse.Connectionattributes.JCO_SNC_LIBRARY),
			IsTimeoutSupported:             util.SafeBoolDatasource(apiResp.SAPConnectionResponse.Connectionattributes.IsTimeoutSupported),
			JcoR3Name:                      util.SafeStringDatasource(apiResp.SAPConnectionResponse.Connectionattributes.JCOR3NAME),
			ExternalSodEvalJson:            util.SafeStringDatasource(apiResp.SAPConnectionResponse.Connectionattributes.EXTERNAL_SOD_EVAL_JSON),
			JcoAshost:                      util.SafeStringDatasource(apiResp.SAPConnectionResponse.Connectionattributes.JCO_ASHOST),
			PasswordNoOfDigits:             util.SafeStringDatasource(apiResp.SAPConnectionResponse.Connectionattributes.PASSWORD_NOOFDIGITS),
			ProvJcoMsHost:                  util.SafeStringDatasource(apiResp.SAPConnectionResponse.Connectionattributes.PROV_JCO_MSHOST),
			PamConfig:                      util.SafeStringDatasource(apiResp.SAPConnectionResponse.Connectionattributes.PAM_CONFIG),
			JcoSncMyName:                   util.SafeStringDatasource(apiResp.SAPConnectionResponse.Connectionattributes.JCO_SNC_MYNAME),
			EnforcePasswordChange:          util.SafeStringDatasource(apiResp.SAPConnectionResponse.Connectionattributes.ENFORCEPASSWORDCHANGE),
			JcoUser:                        util.SafeStringDatasource(apiResp.SAPConnectionResponse.Connectionattributes.JCO_USER),
			JcoSncMode:                     util.SafeStringDatasource(apiResp.SAPConnectionResponse.Connectionattributes.JCO_SNC_MODE),
			ProvJcoMsServ:                  util.SafeStringDatasource(apiResp.SAPConnectionResponse.Connectionattributes.PROV_JCO_MSSERV),
			HanaRefTableJson:               util.SafeStringDatasource(apiResp.SAPConnectionResponse.Connectionattributes.HANAREFTABLEJSON),
			PasswordMinLength:              util.SafeStringDatasource(apiResp.SAPConnectionResponse.Connectionattributes.PASSWORD_MIN_LENGTH),
			JcoClient:                      util.SafeStringDatasource(apiResp.SAPConnectionResponse.Connectionattributes.JCO_CLIENT),
			TerminatedUserRoleAction:       util.SafeStringDatasource(apiResp.SAPConnectionResponse.Connectionattributes.TERMINATED_USER_ROLE_ACTION),
			ResetPwdForNewAccount:          util.SafeStringDatasource(apiResp.SAPConnectionResponse.Connectionattributes.RESET_PWD_FOR_NEWACCOUNT),
			ProvJcoClient:                  util.SafeStringDatasource(apiResp.SAPConnectionResponse.Connectionattributes.PROV_JCO_CLIENT),
			Snc:                            util.SafeStringDatasource(apiResp.SAPConnectionResponse.Connectionattributes.SNC),
			JcoMsServ:                      util.SafeStringDatasource(apiResp.SAPConnectionResponse.Connectionattributes.JCO_MSSERV),
			ProvCuaSnc:                     util.SafeStringDatasource(apiResp.SAPConnectionResponse.Connectionattributes.PROV_CUA_SNC),
			ProvJcoUser:                    util.SafeStringDatasource(apiResp.SAPConnectionResponse.Connectionattributes.PROV_JCO_USER),
			JcoLang:                        util.SafeStringDatasource(apiResp.SAPConnectionResponse.Connectionattributes.JCO_LANG),
			JcoSncPartnerName:              util.SafeStringDatasource(apiResp.SAPConnectionResponse.Connectionattributes.JCO_SNC_PARTNERNAME),
			StatusThresholdConfig:          util.SafeStringDatasource(apiResp.SAPConnectionResponse.Connectionattributes.STATUS_THRESHOLD_CONFIG),
			ProvJcoSysNr:                   util.SafeStringDatasource(apiResp.SAPConnectionResponse.Connectionattributes.PROV_JCO_SYSNR),
			SetCuaSystem:                   util.SafeStringDatasource(apiResp.SAPConnectionResponse.Connectionattributes.SETCUASYSTEM),
			MessageServer:                  util.SafeStringDatasource(apiResp.SAPConnectionResponse.Connectionattributes.MESSAGESERVER),
			ProvJcoAshost:                  util.SafeStringDatasource(apiResp.SAPConnectionResponse.Connectionattributes.PROV_JCO_ASHOST),
			ProvJcoGroup:                   util.SafeStringDatasource(apiResp.SAPConnectionResponse.Connectionattributes.PROV_JCO_GROUP),
			ProvCuaEnabled:                 util.SafeStringDatasource(apiResp.SAPConnectionResponse.Connectionattributes.PROV_CUA_ENABLED),
			JcoMsHost:                      util.SafeStringDatasource(apiResp.SAPConnectionResponse.Connectionattributes.JCO_MSHOST),
			ProvJcoR3Name:                  util.SafeStringDatasource(apiResp.SAPConnectionResponse.Connectionattributes.PROVJCOR3NAME),
			PasswordNoOfCapsAlpha:          util.SafeStringDatasource(apiResp.SAPConnectionResponse.Connectionattributes.PASSWORD_NOOFCAPSALPHA),
			ModifyUserDataJson:             util.SafeStringDatasource(apiResp.SAPConnectionResponse.Connectionattributes.MODIFYUSERDATAJSON),
			IsTimeoutConfigValidated:       util.SafeBoolDatasource(apiResp.SAPConnectionResponse.Connectionattributes.IsTimeoutConfigValidated),
			JcoSncQop:                      util.SafeStringDatasource(apiResp.SAPConnectionResponse.Connectionattributes.JCO_SNC_QOP),
			Tables:                         util.SafeStringDatasource(apiResp.SAPConnectionResponse.Connectionattributes.TABLES),
			ProvJcoLang:                    util.SafeStringDatasource(apiResp.SAPConnectionResponse.Connectionattributes.PROV_JCO_LANG),
			JcoSysNr:                       util.SafeStringDatasource(apiResp.SAPConnectionResponse.Connectionattributes.JCO_SYSNR),
			ExternalSodEvalJsonDetail:      util.SafeStringDatasource(apiResp.SAPConnectionResponse.Connectionattributes.EXTERNAL_SOD_EVAL_JSON_DETAIL),
			DataImportFilter:               util.SafeStringDatasource(apiResp.SAPConnectionResponse.Connectionattributes.DATA_IMPORT_FILTER),
			EnableAccountJson:              util.SafeStringDatasource(apiResp.SAPConnectionResponse.Connectionattributes.ENABLEACCOUNTJSON),
			AlternateOutputParameterEtData: util.SafeStringDatasource(apiResp.SAPConnectionResponse.Connectionattributes.ALTERNATE_OUTPUT_PARAMETER_ET_DATA),
			JcoGroup:                       util.SafeStringDatasource(apiResp.SAPConnectionResponse.Connectionattributes.JCO_GROUP),
			PasswordMaxLength:              util.SafeStringDatasource(apiResp.SAPConnectionResponse.Connectionattributes.PASSWORD_MAX_LENGTH),
			UserImportJson:                 util.SafeStringDatasource(apiResp.SAPConnectionResponse.Connectionattributes.USERIMPORTJSON),
			SystemName:                     util.SafeStringDatasource(apiResp.SAPConnectionResponse.Connectionattributes.SYSTEMNAME),
			UpdateAccountJson:              util.SafeStringDatasource(apiResp.SAPConnectionResponse.Connectionattributes.UPDATEACCOUNTJSON),
		}
		if apiResp.SAPConnectionResponse.Connectionattributes.ConnectionTimeoutConfig != nil {
			state.ConnectionAttributes.ConnectionTimeoutConfig = &ConnectionTimeoutConfig{
				RetryWait:               util.SafeInt64(apiResp.SAPConnectionResponse.Connectionattributes.ConnectionTimeoutConfig.RetryWait),
				TokenRefreshMaxTryCount: util.SafeInt64(apiResp.SAPConnectionResponse.Connectionattributes.ConnectionTimeoutConfig.TokenRefreshMaxTryCount),
				RetryFailureStatusCode:  util.SafeInt64(apiResp.SAPConnectionResponse.Connectionattributes.ConnectionTimeoutConfig.RetryFailureStatusCode),
				RetryWaitMaxValue:       util.SafeInt64(apiResp.SAPConnectionResponse.Connectionattributes.ConnectionTimeoutConfig.RetryWaitMaxValue),
				RetryCount:              util.SafeInt64(apiResp.SAPConnectionResponse.Connectionattributes.ConnectionTimeoutConfig.RetryCount),
				ReadTimeout:             util.SafeInt64(apiResp.SAPConnectionResponse.Connectionattributes.ConnectionTimeoutConfig.ReadTimeout),
				ConnectionTimeout:       util.SafeInt64(apiResp.SAPConnectionResponse.Connectionattributes.ConnectionTimeoutConfig.ConnectionTimeout),
			}
		}
	}

	if apiResp.SAPConnectionResponse.Connectionattributes == nil {
		state.ConnectionAttributes = nil
	}
	if !state.Authenticate.IsNull() && !state.Authenticate.IsUnknown() {
		if state.Authenticate.ValueBool() {
			resp.Diagnostics.AddWarning(
				"Authentication Enabled",
				"`authenticate` is true; all connection_attributes will be returned in state.",
			)
		} else {
			resp.Diagnostics.AddWarning(
				"Authentication Disabled",
				"`authenticate` is false; connection_attributes will be removed from state.",
			)
			state.ConnectionAttributes = nil
		}
	}
	stateDiagnostics := resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(stateDiagnostics...)
}
