// Copyright (c) 2025 Saviynt Inc.
// SPDX-License-Identifier: MPL-2.0

// saviynt_workday_connection_datasource retrieves workday connections details from the Saviynt Security Manager.
// The data source supports a single Read operation to look up an existing workday connections by name.
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

var _ datasource.DataSource = &workdayConnectionDataSource{}

// WorkdayConnectionDataSource defines the data source
type workdayConnectionDataSource struct {
	client *s.Client
	token  string
}

type WorkdayConnectionDataSourceModel struct {
	ID types.String `tfsdk:"id"`
	BaseConnectionDataSourceModel
	ConnectionAttributes *WorkdayConnectionAttributes `tfsdk:"connection_attributes"`
}

type WorkdayConnectionAttributes struct {
	UseOauth                    types.String             `tfsdk:"use_oauth"`
	UserImportMapping           types.String             `tfsdk:"user_import_mapping"`
	AccountsLastImportTime      types.String             `tfsdk:"accounts_last_import_time"`
	StatusKeyJson               types.String             `tfsdk:"status_key_json"`
	ConnectionTimeoutConfig     *ConnectionTimeoutConfig `tfsdk:"connection_timeout_config"`
	ConnectionType              types.String             `tfsdk:"connection_type"`
	RaasMappingJson             types.String             `tfsdk:"raas_mapping_json"`
	AccountImportPayload        types.String             `tfsdk:"account_import_payload"`
	UpdateAccountPayload        types.String             `tfsdk:"update_account_payload"`
	StatusThresholdConfig       types.String             `tfsdk:"status_threshold_config"`
	AccessImportList            types.String             `tfsdk:"access_import_list"`
	IsTimeoutSupported          types.Bool               `tfsdk:"is_timeout_supported"`
	AccountImportMapping        types.String             `tfsdk:"account_import_mapping"`
	AssignOrgrolePayload        types.String             `tfsdk:"assign_orgrole_payload"`
	AccessImportMapping         types.String             `tfsdk:"access_import_mapping"`
	ApiVersion                  types.String             `tfsdk:"api_version"`
	RemoveOrgrolePayload        types.String             `tfsdk:"remove_orgrole_payload"`
	IncludeReferenceDescriptors types.String             `tfsdk:"include_reference_descriptors"`
	ModifyUserDataJson          types.String             `tfsdk:"modifyuserdatajson"`
	IsTimeoutConfigValidated    types.Bool               `tfsdk:"is_timeout_config_validated"`
	UseX509AuthForSoap          types.String             `tfsdk:"use_x509auth_for_soap"`
	ReportOwner                 types.String             `tfsdk:"report_owner"`
	X509Key                     types.String             `tfsdk:"x509_key"`
	CustomConfig                types.String             `tfsdk:"custom_config"`
	UserAttributeJson           types.String             `tfsdk:"userattributejson"`
	X509Cert                    types.String             `tfsdk:"x509_cert"`
	UserImportPayload           types.String             `tfsdk:"user_import_payload"`
	PamConfig                   types.String             `tfsdk:"pam_config"`
	AccessLastImportTime        types.String             `tfsdk:"access_last_import_time"`
	UsersLastImportTime         types.String             `tfsdk:"users_last_import_time"`
	UpdateUserPayload           types.String             `tfsdk:"update_user_payload"`
	PageSize                    types.String             `tfsdk:"page_size"`
	TenantName                  types.String             `tfsdk:"tenant_name"`
	UseEnhancedOrgrole          types.String             `tfsdk:"use_enhanced_orgrole"`
	CreateAccountPayload        types.String             `tfsdk:"create_account_payload"`
	BaseUrl                     types.String             `tfsdk:"base_url"`
}

func NewWorkdayConnectionsDataSource() datasource.DataSource {
	return &workdayConnectionDataSource{}
}

func (d *workdayConnectionDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "saviynt_workday_connection_datasource"
}

func WorkdayConnectorsDataSourceSchema() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:    true,
			Description: "Resource ID.",
		},
		"connection_attributes": schema.SingleNestedAttribute{
			Computed: true,
			Attributes: map[string]schema.Attribute{
				"use_oauth":                     schema.StringAttribute{Computed: true},
				"user_import_mapping":           schema.StringAttribute{Computed: true},
				"accounts_last_import_time":     schema.StringAttribute{Computed: true},
				"status_key_json":               schema.StringAttribute{Computed: true},
				"connection_type":               schema.StringAttribute{Computed: true},
				"raas_mapping_json":             schema.StringAttribute{Computed: true},
				"account_import_payload":        schema.StringAttribute{Computed: true},
				"update_account_payload":        schema.StringAttribute{Computed: true},
				"status_threshold_config":       schema.StringAttribute{Computed: true},
				"access_import_list":            schema.StringAttribute{Computed: true},
				"is_timeout_supported":          schema.BoolAttribute{Computed: true},
				"account_import_mapping":        schema.StringAttribute{Computed: true},
				"assign_orgrole_payload":        schema.StringAttribute{Computed: true},
				"access_import_mapping":         schema.StringAttribute{Computed: true},
				"api_version":                   schema.StringAttribute{Computed: true},
				"remove_orgrole_payload":        schema.StringAttribute{Computed: true},
				"include_reference_descriptors": schema.StringAttribute{Computed: true},
				"modifyuserdatajson":            schema.StringAttribute{Computed: true},
				"is_timeout_config_validated":   schema.BoolAttribute{Computed: true},
				"use_x509auth_for_soap":         schema.StringAttribute{Computed: true},
				"report_owner":                  schema.StringAttribute{Computed: true},
				"x509_key":                      schema.StringAttribute{Computed: true},
				"custom_config":                 schema.StringAttribute{Computed: true},
				"userattributejson":             schema.StringAttribute{Computed: true},
				"x509_cert":                     schema.StringAttribute{Computed: true},
				"user_import_payload":           schema.StringAttribute{Computed: true},
				"pam_config":                    schema.StringAttribute{Computed: true},
				"access_last_import_time":       schema.StringAttribute{Computed: true},
				"users_last_import_time":        schema.StringAttribute{Computed: true},
				"update_user_payload":           schema.StringAttribute{Computed: true},
				"page_size":                     schema.StringAttribute{Computed: true},
				"tenant_name":                   schema.StringAttribute{Computed: true},
				"use_enhanced_orgrole":          schema.StringAttribute{Computed: true},
				"create_account_payload":        schema.StringAttribute{Computed: true},
				"base_url":                      schema.StringAttribute{Computed: true},
				"connection_timeout_config": schema.SingleNestedAttribute{
					Computed:   true,
					Attributes: ConnectionTimeoutConfigeSchema(),
				},
			},
		},
	}
}

// Schema defines the attributes for the data source
func (d *workdayConnectionDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: util.WorkdayConnDataSourceDescription,
		Attributes:  connectionsutil.MergeDataSourceAttributes(BaseConnectorDataSourceSchema(), WorkdayConnectorsDataSourceSchema()),
	}
}

func (d *workdayConnectionDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (d *workdayConnectionDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var state WorkdayConnectionDataSourceModel

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
			log.Printf("[ERROR] HTTP error while creating Workday Connector: %s", httpResp.Status)
			var fetchResp map[string]interface{}
			if err := json.NewDecoder(httpResp.Body).Decode(&fetchResp); err != nil {
				resp.Diagnostics.AddError("Failed to decode error response", err.Error())
				return
			}
			resp.Diagnostics.AddError(
				"HTTP Error",
				fmt.Sprintf("HTTP error while creating Workday Connector for the reasons: %s", fetchResp["msg"]),
			)

		} else {
			log.Printf("[ERROR] API Call Failed: %v", err)
			resp.Diagnostics.AddError("API Call Failed", fmt.Sprintf("Error: %v", err))
		}
		return
	}

	if apiResp != nil && apiResp.WorkdayConnectionResponse == nil {
		error := "Verify the connection type"
		log.Printf("[ERROR]: Verify the connection type given")
		resp.Diagnostics.AddError("Read of Workday connection failed", error)
		return
	}
	log.Printf("[DEBUG] HTTP Status Code: %d", httpResp.StatusCode)

	state.Msg = util.SafeStringDatasource(apiResp.WorkdayConnectionResponse.Msg)
	state.ErrorCode = util.SafeInt64(apiResp.WorkdayConnectionResponse.Errorcode)
	state.ConnectionName = util.SafeStringDatasource(apiResp.WorkdayConnectionResponse.Connectionname)
	state.ConnectionKey = util.SafeInt64(apiResp.WorkdayConnectionResponse.Connectionkey)
	state.Description = util.SafeStringDatasource(apiResp.WorkdayConnectionResponse.Description)
	state.DefaultSavRoles = util.SafeStringDatasource(apiResp.WorkdayConnectionResponse.Defaultsavroles)
	state.ConnectionType = util.SafeStringDatasource(apiResp.WorkdayConnectionResponse.Connectiontype)
	state.CreatedOn = util.SafeStringDatasource(apiResp.WorkdayConnectionResponse.Createdon)
	state.CreatedBy = util.SafeStringDatasource(apiResp.WorkdayConnectionResponse.Createdby)
	state.UpdatedBy = util.SafeStringDatasource(apiResp.WorkdayConnectionResponse.Updatedby)
	state.EmailTemplate = util.SafeStringDatasource(apiResp.WorkdayConnectionResponse.Emailtemplate)

	if apiResp.WorkdayConnectionResponse.Connectionattributes != nil {
		state.ConnectionAttributes = &WorkdayConnectionAttributes{
			UseOauth:                    util.SafeStringDatasource(apiResp.WorkdayConnectionResponse.Connectionattributes.USE_OAUTH),
			UserImportMapping:           util.SafeStringDatasource(apiResp.WorkdayConnectionResponse.Connectionattributes.USER_IMPORT_MAPPING),
			AccountsLastImportTime:      util.SafeStringDatasource(apiResp.WorkdayConnectionResponse.Connectionattributes.ACCOUNTS_LAST_IMPORT_TIME),
			StatusKeyJson:               util.SafeStringDatasource(apiResp.WorkdayConnectionResponse.Connectionattributes.STATUS_KEY_JSON),
			ConnectionType:              util.SafeStringDatasource(apiResp.WorkdayConnectionResponse.Connectionattributes.ConnectionType),
			RaasMappingJson:             util.SafeStringDatasource(apiResp.WorkdayConnectionResponse.Connectionattributes.RAAS_MAPPING_JSON),
			AccountImportPayload:        util.SafeStringDatasource(apiResp.WorkdayConnectionResponse.Connectionattributes.ACCOUNT_IMPORT_PAYLOAD),
			UpdateAccountPayload:        util.SafeStringDatasource(apiResp.WorkdayConnectionResponse.Connectionattributes.UPDATE_ACCOUNT_PAYLOAD),
			StatusThresholdConfig:       util.SafeStringDatasource(apiResp.WorkdayConnectionResponse.Connectionattributes.STATUS_THRESHOLD_CONFIG),
			AccessImportList:            util.SafeStringDatasource(apiResp.WorkdayConnectionResponse.Connectionattributes.ACCESS_IMPORT_LIST),
			IsTimeoutSupported:          util.SafeBoolDatasource(apiResp.WorkdayConnectionResponse.Connectionattributes.IsTimeoutSupported),
			AccountImportMapping:        util.SafeStringDatasource(apiResp.WorkdayConnectionResponse.Connectionattributes.ACCOUNT_IMPORT_MAPPING),
			AssignOrgrolePayload:        util.SafeStringDatasource(apiResp.WorkdayConnectionResponse.Connectionattributes.ASSIGN_ORGROLE_PAYLOAD),
			AccessImportMapping:         util.SafeStringDatasource(apiResp.WorkdayConnectionResponse.Connectionattributes.ACCESS_IMPORT_MAPPING),
			ApiVersion:                  util.SafeStringDatasource(apiResp.WorkdayConnectionResponse.Connectionattributes.API_VERSION),
			RemoveOrgrolePayload:        util.SafeStringDatasource(apiResp.WorkdayConnectionResponse.Connectionattributes.REMOVE_ORGROLE_PAYLOAD),
			IncludeReferenceDescriptors: util.SafeStringDatasource(apiResp.WorkdayConnectionResponse.Connectionattributes.INCLUDE_REFERENCE_DESCRIPTORS),
			ModifyUserDataJson:          util.SafeStringDatasource(apiResp.WorkdayConnectionResponse.Connectionattributes.MODIFYUSERDATAJSON),
			IsTimeoutConfigValidated:    util.SafeBoolDatasource(apiResp.WorkdayConnectionResponse.Connectionattributes.IsTimeoutConfigValidated),
			UseX509AuthForSoap:          util.SafeStringDatasource(apiResp.WorkdayConnectionResponse.Connectionattributes.USEX509AUTHFORSOAP),
			ReportOwner:                 util.SafeStringDatasource(apiResp.WorkdayConnectionResponse.Connectionattributes.REPORT_OWNER),
			X509Key:                     util.SafeStringDatasource(apiResp.WorkdayConnectionResponse.Connectionattributes.X509KEY),
			CustomConfig:                util.SafeStringDatasource(apiResp.WorkdayConnectionResponse.Connectionattributes.CUSTOM_CONFIG),
			UserAttributeJson:           util.SafeStringDatasource(apiResp.WorkdayConnectionResponse.Connectionattributes.USERATTRIBUTEJSON),
			X509Cert:                    util.SafeStringDatasource(apiResp.WorkdayConnectionResponse.Connectionattributes.X509CERT),
			UserImportPayload:           util.SafeStringDatasource(apiResp.WorkdayConnectionResponse.Connectionattributes.USER_IMPORT_PAYLOAD),
			PamConfig:                   util.SafeStringDatasource(apiResp.WorkdayConnectionResponse.Connectionattributes.PAM_CONFIG),
			AccessLastImportTime:        util.SafeStringDatasource(apiResp.WorkdayConnectionResponse.Connectionattributes.ACCESS_LAST_IMPORT_TIME),
			UsersLastImportTime:         util.SafeStringDatasource(apiResp.WorkdayConnectionResponse.Connectionattributes.USERS_LAST_IMPORT_TIME),
			UpdateUserPayload:           util.SafeStringDatasource(apiResp.WorkdayConnectionResponse.Connectionattributes.UPDATE_USER_PAYLOAD),
			PageSize:                    util.SafeStringDatasource(apiResp.WorkdayConnectionResponse.Connectionattributes.PAGE_SIZE),
			TenantName:                  util.SafeStringDatasource(apiResp.WorkdayConnectionResponse.Connectionattributes.TENANT_NAME),
			UseEnhancedOrgrole:          util.SafeStringDatasource(apiResp.WorkdayConnectionResponse.Connectionattributes.USE_ENHANCED_ORGROLE),
			CreateAccountPayload:        util.SafeStringDatasource(apiResp.WorkdayConnectionResponse.Connectionattributes.CREATE_ACCOUNT_PAYLOAD),
			BaseUrl:                     util.SafeStringDatasource(apiResp.WorkdayConnectionResponse.Connectionattributes.BASE_URL),
		}
		if apiResp.WorkdayConnectionResponse.Connectionattributes.ConnectionTimeoutConfig != nil {
			state.ConnectionAttributes.ConnectionTimeoutConfig = &ConnectionTimeoutConfig{
				RetryWait:               util.SafeInt64(apiResp.WorkdayConnectionResponse.Connectionattributes.ConnectionTimeoutConfig.RetryWait),
				TokenRefreshMaxTryCount: util.SafeInt64(apiResp.WorkdayConnectionResponse.Connectionattributes.ConnectionTimeoutConfig.TokenRefreshMaxTryCount),
				RetryFailureStatusCode:  util.SafeInt64(apiResp.WorkdayConnectionResponse.Connectionattributes.ConnectionTimeoutConfig.RetryFailureStatusCode),
				RetryWaitMaxValue:       util.SafeInt64(apiResp.WorkdayConnectionResponse.Connectionattributes.ConnectionTimeoutConfig.RetryWaitMaxValue),
				RetryCount:              util.SafeInt64(apiResp.WorkdayConnectionResponse.Connectionattributes.ConnectionTimeoutConfig.RetryCount),
				ReadTimeout:             util.SafeInt64(apiResp.WorkdayConnectionResponse.Connectionattributes.ConnectionTimeoutConfig.ReadTimeout),
				ConnectionTimeout:       util.SafeInt64(apiResp.WorkdayConnectionResponse.Connectionattributes.ConnectionTimeoutConfig.ConnectionTimeout),
			}
		}
	}

	if apiResp.WorkdayConnectionResponse.Connectionattributes == nil {
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
