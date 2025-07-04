// Copyright (c) Saviynt Inc.
// SPDX-License-Identifier: MPL-2.0

// saviynt_db_connection_datasource retrieves db connections details from the Saviynt Security Manager.
// The data source supports a single Read operation to look up an existing db connections by name.
package provider

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"terraform-provider-Saviynt/util"
	connectionsutil "terraform-provider-Saviynt/util/connectionsutil"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	s "github.com/saviynt/saviynt-api-go-client"

	openapi "github.com/saviynt/saviynt-api-go-client/connections"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.ResourceWithImportState = &dbConnectionResource{}

// DBConnectionsDataSource defines the data source
type dbConnectionsDataSource struct {
	client *s.Client
	token  string
}

type DBConnectionDataSourceModel struct {
	ID types.String `tfsdk:"id"`
	BaseConnectionDataSourceModel
	ConnectionAttributes *DBConnectionAttributes `tfsdk:"connection_attributes"`
}

type DBConnectionAttributes struct {
	PasswordMinLength        types.String             `tfsdk:"password_min_length"`
	ChangePassJSON           types.String             `tfsdk:"changepass_json"`
	AccountExistsJSON        types.String             `tfsdk:"accountexists_json"`
	RolesImport              types.String             `tfsdk:"roles_import"`
	RoleOwnerImport          types.String             `tfsdk:"roleowner_import"`
	CreateAccountJSON        types.String             `tfsdk:"createaccount_json"`
	UserImport               types.String             `tfsdk:"user_import"`
	DisableAccountJSON       types.String             `tfsdk:"disableaccount_json"`
	EntitlementValueImport   types.String             `tfsdk:"entitlementvalue_import"`
	ConnectionType           types.String             `tfsdk:"connection_type"`
	UpdateUserJSON           types.String             `tfsdk:"updateuser_json"`
	PasswordNoOfSplChars     types.String             `tfsdk:"password_noofsplchars"`
	RevokeAccessJSON         types.String             `tfsdk:"revokeaccess_json"`
	URL                      types.String             `tfsdk:"url"`
	SystemImport             types.String             `tfsdk:"system_import"`
	DriverName               types.String             `tfsdk:"drivername"`
	DeleteAccountJSON        types.String             `tfsdk:"deleteaccount_json"`
	StatusThresholdConfig    types.String             `tfsdk:"status_threshold_config"`
	Username                 types.String             `tfsdk:"username"`
	IsTimeoutSupported       types.Bool               `tfsdk:"is_timeout_supported"`
	PasswordNoOfCapsAlpha    types.String             `tfsdk:"password_noofcapsalpha"`
	PasswordNoOfDigits       types.String             `tfsdk:"password_noofdigits"`
	ConnectionProperties     types.String             `tfsdk:"connectionproperties"`
	ModifyUserDataJSON       types.String             `tfsdk:"modifyuserdata_json"`
	IsTimeoutConfigValidated types.Bool               `tfsdk:"is_timeout_config_validated"`
	AccountsImport           types.String             `tfsdk:"accounts_import"`
	Password                 types.String             `tfsdk:"password"`
	EnableAccountJSON        types.String             `tfsdk:"enableaccount_json"`
	PasswordMaxLength        types.String             `tfsdk:"password_max_length"`
	MaxPaginationSize        types.String             `tfsdk:"max_pagination_size"`
	UpdateAccountJSON        types.String             `tfsdk:"updateaccount_json"`
	GrantAccessJSON          types.String             `tfsdk:"grantaccess_json"`
	CliCommandJSON           types.String             `tfsdk:"cli_command_json"`
	ConnectionTimeoutConfig  *ConnectionTimeoutConfig `tfsdk:"connection_timeout_config"`
}

var _ datasource.DataSource = &dbConnectionsDataSource{}

func NewDBConnectionsDataSource() datasource.DataSource {
	return &dbConnectionsDataSource{}
}

func (d *dbConnectionsDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "saviynt_db_connection_datasource"
}

func DBConnectorsDataSourceSchema() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:    true,
			Description: "Resource ID.",
		},
		"connection_attributes": schema.SingleNestedAttribute{
			Computed: true,
			Attributes: map[string]schema.Attribute{
				"password_min_length":         schema.StringAttribute{Computed: true},
				"changepass_json":             schema.StringAttribute{Computed: true},
				"accountexists_json":          schema.StringAttribute{Computed: true},
				"roles_import":                schema.StringAttribute{Computed: true},
				"roleowner_import":            schema.StringAttribute{Computed: true},
				"createaccount_json":          schema.StringAttribute{Computed: true},
				"user_import":                 schema.StringAttribute{Computed: true},
				"disableaccount_json":         schema.StringAttribute{Computed: true},
				"entitlementvalue_import":     schema.StringAttribute{Computed: true},
				"connection_type":             schema.StringAttribute{Computed: true},
				"updateuser_json":             schema.StringAttribute{Computed: true},
				"password_noofsplchars":       schema.StringAttribute{Computed: true},
				"revokeaccess_json":           schema.StringAttribute{Computed: true},
				"url":                         schema.StringAttribute{Computed: true},
				"system_import":               schema.StringAttribute{Computed: true},
				"drivername":                  schema.StringAttribute{Computed: true},
				"deleteaccount_json":          schema.StringAttribute{Computed: true},
				"status_threshold_config":     schema.StringAttribute{Computed: true},
				"username":                    schema.StringAttribute{Computed: true},
				"is_timeout_supported":        schema.BoolAttribute{Computed: true},
				"password_noofcapsalpha":      schema.StringAttribute{Computed: true},
				"password_noofdigits":         schema.StringAttribute{Computed: true},
				"connectionproperties":        schema.StringAttribute{Computed: true},
				"modifyuserdata_json":         schema.StringAttribute{Computed: true},
				"is_timeout_config_validated": schema.BoolAttribute{Computed: true},
				"accounts_import":             schema.StringAttribute{Computed: true},
				"password":                    schema.StringAttribute{Computed: true},
				"enableaccount_json":          schema.StringAttribute{Computed: true},
				"password_max_length":         schema.StringAttribute{Computed: true},
				"max_pagination_size":         schema.StringAttribute{Computed: true},
				"updateaccount_json":          schema.StringAttribute{Computed: true},
				"grantaccess_json":            schema.StringAttribute{Computed: true},
				"cli_command_json":            schema.StringAttribute{Computed: true},
				"connection_timeout_config": schema.SingleNestedAttribute{
					Computed:   true,
					Attributes: ConnectionTimeoutConfigeSchema(),
				},
			},
		},
	}
}

// Schema defines the attributes for the data source
func (d *dbConnectionsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: util.DBConnDataSourceDescription,
		Attributes:  connectionsutil.MergeDataSourceAttributes(BaseConnectorDataSourceSchema(), DBConnectorsDataSourceSchema()),
	}
}

func (d *dbConnectionsDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (d *dbConnectionsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var state DBConnectionDataSourceModel

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
		log.Printf("[ERROR] API Call Failed: %v", err)
		resp.Diagnostics.AddError("API Call Failed", fmt.Sprintf("Error: %v", err))
		return
	}
	log.Printf("[DEBUG] HTTP Status Code: %d", httpResp.StatusCode)

	state.Msg = util.SafeStringDatasource(apiResp.DBConnectionResponse.Msg)
	state.ErrorCode = util.SafeInt64(apiResp.DBConnectionResponse.Errorcode)
	state.ConnectionName = util.SafeStringDatasource(apiResp.DBConnectionResponse.Connectionname)
	state.ConnectionKey = util.SafeInt64(apiResp.DBConnectionResponse.Connectionkey)
	state.Description = util.SafeStringDatasource(apiResp.DBConnectionResponse.Description)
	state.DefaultSavRoles = util.SafeStringDatasource(apiResp.DBConnectionResponse.Defaultsavroles)
	state.ConnectionType = util.SafeStringDatasource(apiResp.DBConnectionResponse.Connectiontype)
	state.CreatedOn = util.SafeStringDatasource(apiResp.DBConnectionResponse.Createdon)
	state.CreatedBy = util.SafeStringDatasource(apiResp.DBConnectionResponse.Createdby)
	state.UpdatedBy = util.SafeStringDatasource(apiResp.DBConnectionResponse.Updatedby)
	state.Msg = util.SafeStringDatasource(apiResp.DBConnectionResponse.Msg)
	state.EmailTemplate = util.SafeStringDatasource(apiResp.DBConnectionResponse.Emailtemplate)

	if apiResp.DBConnectionResponse.Connectionattributes != nil {
		state.ConnectionAttributes = &DBConnectionAttributes{
			PasswordMinLength:        util.SafeStringDatasource(apiResp.DBConnectionResponse.Connectionattributes.PASSWORD_MIN_LENGTH),
			ChangePassJSON:           util.SafeStringDatasource(apiResp.DBConnectionResponse.Connectionattributes.CHANGEPASSJSON),
			AccountExistsJSON:        util.SafeStringDatasource(apiResp.DBConnectionResponse.Connectionattributes.ACCOUNTEXISTSJSON),
			RolesImport:              util.SafeStringDatasource(apiResp.DBConnectionResponse.Connectionattributes.ROLESIMPORT),
			RoleOwnerImport:          util.SafeStringDatasource(apiResp.DBConnectionResponse.Connectionattributes.ROLEOWNERIMPORT),
			CreateAccountJSON:        util.SafeStringDatasource(apiResp.DBConnectionResponse.Connectionattributes.CREATEACCOUNTJSON),
			UserImport:               util.SafeStringDatasource(apiResp.DBConnectionResponse.Connectionattributes.USERIMPORT),
			DisableAccountJSON:       util.SafeStringDatasource(apiResp.DBConnectionResponse.Connectionattributes.DISABLEACCOUNTJSON),
			EntitlementValueImport:   util.SafeStringDatasource(apiResp.DBConnectionResponse.Connectionattributes.ENTITLEMENTVALUEIMPORT),
			ConnectionType:           util.SafeStringDatasource(apiResp.DBConnectionResponse.Connectionattributes.ConnectionType),
			UpdateUserJSON:           util.SafeStringDatasource(apiResp.DBConnectionResponse.Connectionattributes.UPDATEUSERJSON),
			PasswordNoOfSplChars:     util.SafeStringDatasource(apiResp.DBConnectionResponse.Connectionattributes.PASSWORD_NOOFSPLCHARS),
			RevokeAccessJSON:         util.SafeStringDatasource(apiResp.DBConnectionResponse.Connectionattributes.REVOKEACCESSJSON),
			URL:                      util.SafeStringDatasource(apiResp.DBConnectionResponse.Connectionattributes.URL),
			SystemImport:             util.SafeStringDatasource(apiResp.DBConnectionResponse.Connectionattributes.SYSTEMIMPORT),
			DriverName:               util.SafeStringDatasource(apiResp.DBConnectionResponse.Connectionattributes.DRIVERNAME),
			DeleteAccountJSON:        util.SafeStringDatasource(apiResp.DBConnectionResponse.Connectionattributes.DELETEACCOUNTJSON),
			StatusThresholdConfig:    util.SafeStringDatasource(apiResp.DBConnectionResponse.Connectionattributes.STATUS_THRESHOLD_CONFIG),
			Username:                 util.SafeStringDatasource(apiResp.DBConnectionResponse.Connectionattributes.USERNAME),
			IsTimeoutSupported:       util.SafeBoolDatasource(apiResp.DBConnectionResponse.Connectionattributes.IsTimeoutSupported),
			PasswordNoOfCapsAlpha:    util.SafeStringDatasource(apiResp.DBConnectionResponse.Connectionattributes.PASSWORD_NOOFCAPSALPHA),
			PasswordNoOfDigits:       util.SafeStringDatasource(apiResp.DBConnectionResponse.Connectionattributes.PASSWORD_NOOFDIGITS),
			ConnectionProperties:     util.SafeStringDatasource(apiResp.DBConnectionResponse.Connectionattributes.CONNECTIONPROPERTIES),
			ModifyUserDataJSON:       util.SafeStringDatasource(apiResp.DBConnectionResponse.Connectionattributes.MODIFYUSERDATAJSON),
			IsTimeoutConfigValidated: util.SafeBoolDatasource(apiResp.DBConnectionResponse.Connectionattributes.IsTimeoutConfigValidated),
			AccountsImport:           util.SafeStringDatasource(apiResp.DBConnectionResponse.Connectionattributes.ACCOUNTSIMPORT),
			Password:                 util.SafeStringDatasource(apiResp.DBConnectionResponse.Connectionattributes.PASSWORD),
			EnableAccountJSON:        util.SafeStringDatasource(apiResp.DBConnectionResponse.Connectionattributes.ENABLEACCOUNTJSON),
			PasswordMaxLength:        util.SafeStringDatasource(apiResp.DBConnectionResponse.Connectionattributes.PASSWORD_MAX_LENGTH),
			MaxPaginationSize:        util.SafeStringDatasource(apiResp.DBConnectionResponse.Connectionattributes.MAX_PAGINATION_SIZE),
			UpdateAccountJSON:        util.SafeStringDatasource(apiResp.DBConnectionResponse.Connectionattributes.UPDATEACCOUNTJSON),
			GrantAccessJSON:          util.SafeStringDatasource(apiResp.DBConnectionResponse.Connectionattributes.GRANTACCESSJSON),
			CliCommandJSON:           util.SafeStringDatasource(apiResp.DBConnectionResponse.Connectionattributes.CLI_COMMAND_JSON),
		}
		if apiResp.DBConnectionResponse.Connectionattributes.ConnectionTimeoutConfig != nil {
			state.ConnectionAttributes.ConnectionTimeoutConfig = &ConnectionTimeoutConfig{
				RetryWait:               util.SafeInt64(apiResp.DBConnectionResponse.Connectionattributes.ConnectionTimeoutConfig.RetryWait),
				TokenRefreshMaxTryCount: util.SafeInt64(apiResp.DBConnectionResponse.Connectionattributes.ConnectionTimeoutConfig.TokenRefreshMaxTryCount),
				RetryFailureStatusCode:  util.SafeInt64(apiResp.DBConnectionResponse.Connectionattributes.ConnectionTimeoutConfig.RetryFailureStatusCode),
				RetryWaitMaxValue:       util.SafeInt64(apiResp.DBConnectionResponse.Connectionattributes.ConnectionTimeoutConfig.RetryWaitMaxValue),
				RetryCount:              util.SafeInt64(apiResp.DBConnectionResponse.Connectionattributes.ConnectionTimeoutConfig.RetryCount),
				ReadTimeout:             util.SafeInt64(apiResp.DBConnectionResponse.Connectionattributes.ConnectionTimeoutConfig.ReadTimeout),
				ConnectionTimeout:       util.SafeInt64(apiResp.DBConnectionResponse.Connectionattributes.ConnectionTimeoutConfig.ConnectionTimeout),
			}
		}
	}

	if apiResp.DBConnectionResponse.Connectionattributes == nil {
		state.ConnectionAttributes = nil
	}
	stateDiagnostics := resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(stateDiagnostics...)
}
