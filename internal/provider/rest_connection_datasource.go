// Copyright (c) 2025 Saviynt Inc.
// SPDX-License-Identifier: MPL-2.0

// saviynt_rest_connection_datasource retrieves rest connections details from the Saviynt Security Manager.
// The data source supports a single Read operation to look up an existing rest connections by name.
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

var _ datasource.DataSource = &restConnectionDatasource{}

type restConnectionDatasource struct {
	client *s.Client
	token  string
}

type RESTConnectionDataSourceModel struct {
	ID types.String `tfsdk:"id"`
	BaseConnectionDataSourceModel
	ConnectionAttributes *RESTConnectionAttributes `tfsdk:"connection_attributes"`
}

type RESTConnectionAttributes struct {
	UpdateUserJSON           types.String            `tfsdk:"update_user_json"`
	RemoveAccountJSON        types.String            `tfsdk:"remove_account_json"`
	TicketStatusJSON         types.String            `tfsdk:"ticket_status_json"`
	CreateTicketJSON         types.String            `tfsdk:"create_ticket_json"`
	ConnectionType           types.String            `tfsdk:"connection_type"`
	EndpointsFilter          types.String            `tfsdk:"endpoints_filter"`
	PasswdPolicyJSON         types.String            `tfsdk:"passwd_policy_json"`
	ConfigJSON               types.String            `tfsdk:"config_json"`
	AddFFIDAccessJSON        types.String            `tfsdk:"add_ffid_access_json"`
	RemoveFFIDAccessJSON     types.String            `tfsdk:"remove_ffid_access_json"`
	StatusThresholdConfig    types.String            `tfsdk:"status_threshold_config"`
	ModifyUserDataJSON       types.String            `tfsdk:"modify_user_data_json"`
	SendOtpJSON              types.String            `tfsdk:"send_otp_json"`
	ValidateOtpJSON          types.String            `tfsdk:"validate_otp_json"`
	PamConfig                types.String            `tfsdk:"pam_config"`
	ConnectionTimeoutConfig  ConnectionTimeoutConfig `tfsdk:"connection_timeout_config"`
	CreateAccountJSON        types.String            `tfsdk:"create_account_json"`
	UpdateAccountJSON        types.String            `tfsdk:"update_account_json"`
	EnableAccountJSON        types.String            `tfsdk:"enable_account_json"`
	DisableAccountJSON       types.String            `tfsdk:"disable_account_json"`
	AddAccessJSON            types.String            `tfsdk:"add_access_json"`
	RemoveAccessJSON         types.String            `tfsdk:"remove_access_json"`
	ImportUserJSON           types.String            `tfsdk:"import_user_json"`
	IsTimeoutSupported       types.Bool              `tfsdk:"is_timeout_supported"`
	ImportAccountEntJSON     types.String            `tfsdk:"import_account_ent_json"`
	IsTimeoutConfigValidated types.Bool              `tfsdk:"is_timeout_config_validated"`
	//TER-176
	ApplicationDiscoveryJson types.String `tfsdk:"application_discovery_json"`
	CreateEntitlementJson    types.String `tfsdk:"create_entitlement_json"`
	DeleteEntitlementJson    types.String `tfsdk:"delete_entitlement_json"`
	UpdateEntitlementJson    types.String `tfsdk:"update_entitlement_json"`
}

func NewRESTConnectionDatasource() datasource.DataSource {
	return &restConnectionDatasource{}
}

func (d *restConnectionDatasource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "saviynt_rest_connection_datasource"
}

func RESTConnectorsDataSourceSchema() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:    true,
			Description: "Resource ID.",
		},
		"connection_attributes": schema.SingleNestedAttribute{
			Computed: true,
			Attributes: map[string]schema.Attribute{
				"update_user_json":            schema.StringAttribute{Computed: true},
				"remove_account_json":         schema.StringAttribute{Computed: true},
				"ticket_status_json":          schema.StringAttribute{Computed: true},
				"create_ticket_json":          schema.StringAttribute{Computed: true},
				"connection_type":             schema.StringAttribute{Computed: true},
				"endpoints_filter":            schema.StringAttribute{Computed: true},
				"passwd_policy_json":          schema.StringAttribute{Computed: true},
				"config_json":                 schema.StringAttribute{Computed: true},
				"add_ffid_access_json":        schema.StringAttribute{Computed: true},
				"remove_ffid_access_json":     schema.StringAttribute{Computed: true},
				"status_threshold_config":     schema.StringAttribute{Computed: true},
				"modify_user_data_json":       schema.StringAttribute{Computed: true},
				"send_otp_json":               schema.StringAttribute{Computed: true},
				"validate_otp_json":           schema.StringAttribute{Computed: true},
				"pam_config":                  schema.StringAttribute{Computed: true},
				"create_account_json":         schema.StringAttribute{Computed: true},
				"update_account_json":         schema.StringAttribute{Computed: true},
				"enable_account_json":         schema.StringAttribute{Computed: true},
				"disable_account_json":        schema.StringAttribute{Computed: true},
				"add_access_json":             schema.StringAttribute{Computed: true},
				"remove_access_json":          schema.StringAttribute{Computed: true},
				"import_user_json":            schema.StringAttribute{Computed: true},
				"is_timeout_supported":        schema.BoolAttribute{Computed: true},
				"import_account_ent_json":     schema.StringAttribute{Computed: true},
				"is_timeout_config_validated": schema.BoolAttribute{Computed: true},
				"connection_timeout_config": schema.SingleNestedAttribute{
					Computed:   true,
					Attributes: ConnectionTimeoutConfigeSchema(),
				},
				//TER-176
				"application_discovery_json": schema.StringAttribute{Computed: true},
				"create_entitlement_json":    schema.StringAttribute{Computed: true},
				"delete_entitlement_json":    schema.StringAttribute{Computed: true},
				"update_entitlement_json":    schema.StringAttribute{Computed: true},
			},
		},
	}
}

// Schema defines the attributes for the data source
func (d *restConnectionDatasource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: util.RestConnDataSourceDescription,
		Attributes:  connectionsutil.MergeDataSourceAttributes(BaseConnectorDataSourceSchema(), RESTConnectorsDataSourceSchema()),
	}
}

func (d *restConnectionDatasource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (d *restConnectionDatasource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var state RESTConnectionDataSourceModel

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
			log.Printf("[ERROR] HTTP error while creating Rest Connector: %s", httpResp.Status)
			var fetchResp map[string]interface{}
			if err := json.NewDecoder(httpResp.Body).Decode(&fetchResp); err != nil {
				resp.Diagnostics.AddError("Failed to decode error response", err.Error())
				return
			}
			resp.Diagnostics.AddError(
				"HTTP Error",
				fmt.Sprintf("HTTP error while creating Rest Connector for the reasons: %s", fetchResp["msg"]),
			)

		} else {
			log.Printf("[ERROR] API Call Failed: %v", err)
			resp.Diagnostics.AddError("API Call Failed", fmt.Sprintf("Error: %v", err))
		}
		return
	}

	if apiResp != nil && apiResp.RESTConnectionResponse == nil {
		error := "Verify the connection type"
		log.Printf("[ERROR]: Verify the connection type given")
		resp.Diagnostics.AddError("Read of REST connection failed", error)
		return
	}

	log.Printf("[DEBUG] HTTP Status Code: %d", httpResp.StatusCode)

	state.Msg = util.SafeStringDatasource(apiResp.RESTConnectionResponse.Msg)
	state.ErrorCode = util.SafeInt64(apiResp.RESTConnectionResponse.Errorcode)
	state.ConnectionName = util.SafeStringDatasource(apiResp.RESTConnectionResponse.Connectionname)
	state.ConnectionKey = util.SafeInt64(apiResp.RESTConnectionResponse.Connectionkey)
	state.Description = util.SafeStringDatasource(apiResp.RESTConnectionResponse.Description)
	state.DefaultSavRoles = util.SafeStringDatasource(apiResp.RESTConnectionResponse.Defaultsavroles)
	state.ConnectionType = util.SafeStringDatasource(apiResp.RESTConnectionResponse.Connectiontype)
	state.CreatedOn = util.SafeStringDatasource(apiResp.RESTConnectionResponse.Createdon)
	state.CreatedBy = util.SafeStringDatasource(apiResp.RESTConnectionResponse.Createdby)
	state.UpdatedBy = util.SafeStringDatasource(apiResp.RESTConnectionResponse.Updatedby)
	state.EmailTemplate = util.SafeStringDatasource(apiResp.RESTConnectionResponse.Emailtemplate)

	if apiResp.RESTConnectionResponse.Connectionattributes != nil {
		state.ConnectionAttributes = &RESTConnectionAttributes{
			UpdateUserJSON:           util.SafeStringDatasource(apiResp.RESTConnectionResponse.Connectionattributes.UpdateUserJSON),
			RemoveAccountJSON:        util.SafeStringDatasource(apiResp.RESTConnectionResponse.Connectionattributes.RemoveAccountJSON),
			TicketStatusJSON:         util.SafeStringDatasource(apiResp.RESTConnectionResponse.Connectionattributes.TicketStatusJSON),
			CreateTicketJSON:         util.SafeStringDatasource(apiResp.RESTConnectionResponse.Connectionattributes.CreateTicketJSON),
			ConnectionType:           util.SafeStringDatasource(apiResp.RESTConnectionResponse.Connectionattributes.ConnectionType),
			EndpointsFilter:          util.SafeStringDatasource(apiResp.RESTConnectionResponse.Connectionattributes.ENDPOINTS_FILTER),
			PasswdPolicyJSON:         util.SafeStringDatasource(apiResp.RESTConnectionResponse.Connectionattributes.PasswdPolicyJSON),
			ConfigJSON:               util.SafeStringDatasource(apiResp.RESTConnectionResponse.Connectionattributes.ConfigJSON),
			AddFFIDAccessJSON:        util.SafeStringDatasource(apiResp.RESTConnectionResponse.Connectionattributes.AddFFIDAccessJSON),
			RemoveFFIDAccessJSON:     util.SafeStringDatasource(apiResp.RESTConnectionResponse.Connectionattributes.RemoveFFIDAccessJSON),
			StatusThresholdConfig:    util.SafeStringDatasource(apiResp.RESTConnectionResponse.Connectionattributes.STATUS_THRESHOLD_CONFIG),
			ModifyUserDataJSON:       util.SafeStringDatasource(apiResp.RESTConnectionResponse.Connectionattributes.MODIFYUSERDATAJSON),
			SendOtpJSON:              util.SafeStringDatasource(apiResp.RESTConnectionResponse.Connectionattributes.SendOtpJSON),
			ValidateOtpJSON:          util.SafeStringDatasource(apiResp.RESTConnectionResponse.Connectionattributes.ValidateOtpJSON),
			PamConfig:                util.SafeStringDatasource(apiResp.RESTConnectionResponse.Connectionattributes.PAM_CONFIG),
			CreateAccountJSON:        util.SafeStringDatasource(apiResp.RESTConnectionResponse.Connectionattributes.CreateAccountJSON),
			UpdateAccountJSON:        util.SafeStringDatasource(apiResp.RESTConnectionResponse.Connectionattributes.UpdateAccountJSON),
			EnableAccountJSON:        util.SafeStringDatasource(apiResp.RESTConnectionResponse.Connectionattributes.EnableAccountJSON),
			DisableAccountJSON:       util.SafeStringDatasource(apiResp.RESTConnectionResponse.Connectionattributes.DisableAccountJSON),
			AddAccessJSON:            util.SafeStringDatasource(apiResp.RESTConnectionResponse.Connectionattributes.AddAccessJSON),
			RemoveAccessJSON:         util.SafeStringDatasource(apiResp.RESTConnectionResponse.Connectionattributes.RemoveAccessJSON),
			ImportUserJSON:           util.SafeStringDatasource(apiResp.RESTConnectionResponse.Connectionattributes.ImportUserJSON),
			IsTimeoutSupported:       util.SafeBoolDatasource(apiResp.RESTConnectionResponse.Connectionattributes.IsTimeoutSupported),
			ImportAccountEntJSON:     util.SafeStringDatasource(apiResp.RESTConnectionResponse.Connectionattributes.ImportAccountEntJSON),
			IsTimeoutConfigValidated: util.SafeBoolDatasource(apiResp.RESTConnectionResponse.Connectionattributes.IsTimeoutConfigValidated),
			ConnectionTimeoutConfig: ConnectionTimeoutConfig{
				RetryWait:               util.SafeInt64(apiResp.RESTConnectionResponse.Connectionattributes.ConnectionTimeoutConfig.RetryWait),
				TokenRefreshMaxTryCount: util.SafeInt64(apiResp.RESTConnectionResponse.Connectionattributes.ConnectionTimeoutConfig.TokenRefreshMaxTryCount),
				RetryWaitMaxValue:       util.SafeInt64(apiResp.RESTConnectionResponse.Connectionattributes.ConnectionTimeoutConfig.RetryWaitMaxValue),
				RetryCount:              util.SafeInt64(apiResp.RESTConnectionResponse.Connectionattributes.ConnectionTimeoutConfig.RetryCount),
				ReadTimeout:             util.SafeInt64(apiResp.RESTConnectionResponse.Connectionattributes.ConnectionTimeoutConfig.ReadTimeout),
				ConnectionTimeout:       util.SafeInt64(apiResp.RESTConnectionResponse.Connectionattributes.ConnectionTimeoutConfig.ConnectionTimeout),
				RetryFailureStatusCode:  util.SafeInt64(apiResp.RESTConnectionResponse.Connectionattributes.ConnectionTimeoutConfig.RetryFailureStatusCode),
			},
			ApplicationDiscoveryJson: util.SafeStringDatasource(apiResp.RESTConnectionResponse.Connectionattributes.ApplicationDiscoveryJSON),
			CreateEntitlementJson:    util.SafeStringDatasource(apiResp.RESTConnectionResponse.Connectionattributes.CreateEntitlementJSON),
			DeleteEntitlementJson:    util.SafeStringDatasource(apiResp.RESTConnectionResponse.Connectionattributes.DeleteEntitlementJSON),
			UpdateEntitlementJson:    util.SafeStringDatasource(apiResp.RESTConnectionResponse.Connectionattributes.UpdateEntitlementJSON),
		}
	}

	if apiResp.RESTConnectionResponse.Connectionattributes == nil {
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
