// Copyright (c) 2025 Saviynt Inc.
// SPDX-License-Identifier: MPL-2.0

// saviynt_ad_connection_datasource retrieves ad connections details from the Saviynt Security Manager.
// The data source supports a single Read operation to look up an existing ad connections by name.
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

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	s "github.com/saviynt/saviynt-api-go-client"

	connectionsutil "terraform-provider-Saviynt/util/connectionsutil"

	openapi "github.com/saviynt/saviynt-api-go-client/connections"
)

var _ datasource.DataSource = &adConnectionsDataSource{}

// ADConnectionsDataSource defines the data source
type adConnectionsDataSource struct {
	client *s.Client
	token  string
}

type ADConnectionDataSourceModel struct {
	ID types.String `tfsdk:"id"`
	BaseConnectionDataSourceModel
	ConnectionAttributes *ADConnectionAttributes `tfsdk:"connection_attributes"`
}

type ADConnectionAttributes struct {
	URL                        types.String             `tfsdk:"url"`
	ConnectionType             types.String             `tfsdk:"connection_type"`
	LastImportTime             types.String             `tfsdk:"last_import_time"`
	CreateAccountJSON          types.String             `tfsdk:"create_account_json"`
	DisableAccountJSON         types.String             `tfsdk:"disable_account_json"`
	GroupSearchBaseDN          types.String             `tfsdk:"group_search_base_dn"`
	PasswordNoOfSplChars       types.String             `tfsdk:"password_no_of_spl_chars"`
	PasswordNoOfDigits         types.String             `tfsdk:"password_no_of_digits"`
	StatusKeyJSON              types.String             `tfsdk:"status_key_json"`
	SearchFilter               types.String             `tfsdk:"search_filter"`
	ConfigJSON                 types.String             `tfsdk:"config_json"`
	RemoveAccountAction        types.String             `tfsdk:"remove_account_action"`
	AccountAttribute           types.String             `tfsdk:"account_attribute"`
	AccountNameRule            types.String             `tfsdk:"account_name_rule"`
	AdvSearch                  types.String             `tfsdk:"adv_search"`
	LDAPOrAD                   types.String             `tfsdk:"ldap_or_ad"`
	EntitlementAttribute       types.String             `tfsdk:"entitlement_attribute"`
	SetRandomPassword          types.String             `tfsdk:"set_random_password"`
	PasswordMinLength          types.String             `tfsdk:"password_min_length"`
	PasswordMaxLength          types.String             `tfsdk:"password_max_length"`
	PasswordNoOfCapsAlpha      types.String             `tfsdk:"password_no_of_caps_alpha"`
	SetDefaultPageSize         types.String             `tfsdk:"set_default_page_size"`
	IsTimeoutSupported         types.Bool               `tfsdk:"is_timeout_supported"`
	ReuseInactiveAccount       types.String             `tfsdk:"reuse_inactive_account"`
	ImportJSON                 types.String             `tfsdk:"import_json"`
	CreateUpdateMappings       types.String             `tfsdk:"create_update_mappings"`
	AdvanceFilterJSON          types.String             `tfsdk:"advance_filter_json"`
	OrgImportJSON              types.String             `tfsdk:"org_import_json"`
	PAMConfig                  types.String             `tfsdk:"pam_config"`
	PageSize                   types.String             `tfsdk:"page_size"`
	Base                       types.String             `tfsdk:"base"`
	DCLocator                  types.String             `tfsdk:"dc_locator"`
	StatusThresholdConfig      types.String             `tfsdk:"status_threshold_config"`
	ResetAndChangePasswordJSON types.String             `tfsdk:"reset_and_change_password_json"`
	SupportEmptyString         types.String             `tfsdk:"support_empty_string"`
	ReadOperationalAttributes  types.String             `tfsdk:"read_operational_attributes"`
	EnableAccountJSON          types.String             `tfsdk:"enable_account_json"`
	UserAttribute              types.String             `tfsdk:"user_attribute"`
	DefaultUserRole            types.String             `tfsdk:"default_user_role"`
	EndpointsFilter            types.String             `tfsdk:"endpoints_filter"`
	UpdateAccountJSON          types.String             `tfsdk:"update_account_json"`
	ReuseAccountJSON           types.String             `tfsdk:"reuse_account_json"`
	EnforceTreeDeletion        types.String             `tfsdk:"enforce_tree_deletion"`
	Filter                     types.String             `tfsdk:"filter"`
	ObjectFilter               types.String             `tfsdk:"object_filter"`
	UpdateUserJSON             types.String             `tfsdk:"update_user_json"`
	SaveConnection             types.String             `tfsdk:"save_connection"`
	SystemName                 types.String             `tfsdk:"system_name"`
	GroupImportMapping         types.String             `tfsdk:"group_import_mapping"`
	UnlockAccountJSON          types.String             `tfsdk:"unlock_account_json"`
	EnableGroupManagement      types.String             `tfsdk:"enable_group_management"`
	ModifyUserDataJSON         types.String             `tfsdk:"modify_user_data_json"`
	OrgBase                    types.String             `tfsdk:"org_base"`
	OrganizationAttribute      types.String             `tfsdk:"organization_attribute"`
	CreateOrgJSON              types.String             `tfsdk:"create_org_json"`
	UpdateOrgJSON              types.String             `tfsdk:"update_org_json"`
	MaxChangeNumber            types.String             `tfsdk:"max_change_number"`
	IncrementalConfig          types.String             `tfsdk:"incremental_config"`
	CheckForUnique             types.String             `tfsdk:"check_for_unique"`
	ConnectionTimeoutConfig    *ConnectionTimeoutConfig `tfsdk:"connection_timeout_config"`
	IsTimeoutConfigValidated   types.Bool               `tfsdk:"is_timeout_config_validated"`
	ResetAndChangePasswdJSON   types.String             `tfsdk:"reset_and_change_passwd_json"`
}

func NewADConnectionsDataSource() datasource.DataSource {
	return &adConnectionsDataSource{}
}

func (d *adConnectionsDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "saviynt_ad_connection_datasource"
}

func ADConnectorsDataSourceSchema() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:    true,
			Description: "Resource ID.",
		},
		"connection_attributes": schema.SingleNestedAttribute{
			Computed: true,
			Attributes: map[string]schema.Attribute{
				"url":                            schema.StringAttribute{Computed: true},
				"connection_type":                schema.StringAttribute{Computed: true},
				"last_import_time":               schema.StringAttribute{Computed: true},
				"create_account_json":            schema.StringAttribute{Computed: true},
				"disable_account_json":           schema.StringAttribute{Computed: true},
				"group_search_base_dn":           schema.StringAttribute{Computed: true},
				"password_no_of_spl_chars":       schema.StringAttribute{Computed: true},
				"password_no_of_digits":          schema.StringAttribute{Computed: true},
				"status_key_json":                schema.StringAttribute{Computed: true},
				"search_filter":                  schema.StringAttribute{Computed: true},
				"config_json":                    schema.StringAttribute{Computed: true},
				"remove_account_action":          schema.StringAttribute{Computed: true},
				"account_attribute":              schema.StringAttribute{Computed: true},
				"account_name_rule":              schema.StringAttribute{Computed: true},
				"adv_search":                     schema.StringAttribute{Computed: true},
				"ldap_or_ad":                     schema.StringAttribute{Computed: true},
				"entitlement_attribute":          schema.StringAttribute{Computed: true},
				"set_random_password":            schema.StringAttribute{Computed: true},
				"password_min_length":            schema.StringAttribute{Computed: true},
				"password_max_length":            schema.StringAttribute{Computed: true},
				"password_no_of_caps_alpha":      schema.StringAttribute{Computed: true},
				"set_default_page_size":          schema.StringAttribute{Computed: true},
				"is_timeout_supported":           schema.BoolAttribute{Computed: true},
				"reuse_inactive_account":         schema.StringAttribute{Computed: true},
				"import_json":                    schema.StringAttribute{Computed: true},
				"create_update_mappings":         schema.StringAttribute{Computed: true},
				"advance_filter_json":            schema.StringAttribute{Computed: true},
				"org_import_json":                schema.StringAttribute{Computed: true},
				"pam_config":                     schema.StringAttribute{Computed: true},
				"page_size":                      schema.StringAttribute{Computed: true},
				"base":                           schema.StringAttribute{Computed: true},
				"dc_locator":                     schema.StringAttribute{Computed: true},
				"status_threshold_config":        schema.StringAttribute{Computed: true},
				"reset_and_change_password_json": schema.StringAttribute{Computed: true},
				"support_empty_string":           schema.StringAttribute{Computed: true},
				"read_operational_attributes":    schema.StringAttribute{Computed: true},
				"enable_account_json":            schema.StringAttribute{Computed: true},
				"user_attribute":                 schema.StringAttribute{Computed: true},
				"default_user_role":              schema.StringAttribute{Computed: true},
				"endpoints_filter":               schema.StringAttribute{Computed: true},
				"update_account_json":            schema.StringAttribute{Computed: true},
				"reuse_account_json":             schema.StringAttribute{Computed: true},
				"enforce_tree_deletion":          schema.StringAttribute{Computed: true},
				"filter":                         schema.StringAttribute{Computed: true},
				"object_filter":                  schema.StringAttribute{Computed: true},
				"update_user_json":               schema.StringAttribute{Computed: true},
				"save_connection":                schema.StringAttribute{Computed: true},
				"system_name":                    schema.StringAttribute{Computed: true},
				"group_import_mapping":           schema.StringAttribute{Computed: true},
				"unlock_account_json":            schema.StringAttribute{Computed: true},
				"enable_group_management":        schema.StringAttribute{Computed: true},
				"modify_user_data_json":          schema.StringAttribute{Computed: true},
				"org_base":                       schema.StringAttribute{Computed: true},
				"organization_attribute":         schema.StringAttribute{Computed: true},
				"create_org_json":                schema.StringAttribute{Computed: true},
				"update_org_json":                schema.StringAttribute{Computed: true},
				"max_change_number":              schema.StringAttribute{Computed: true},
				"incremental_config":             schema.StringAttribute{Computed: true},
				"check_for_unique":               schema.StringAttribute{Computed: true},
				"reset_and_change_passwd_json":   schema.StringAttribute{Computed: true},
				"connection_timeout_config": schema.SingleNestedAttribute{
					Computed:   true,
					Attributes: ConnectionTimeoutConfigeSchema(),
				},
				"is_timeout_config_validated": schema.BoolAttribute{Computed: true},
			},
		},
	}
}
func (d *adConnectionsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: util.ADConnDataSourceDescription,
		Attributes:  connectionsutil.MergeDataSourceAttributes(BaseConnectorDataSourceSchema(), ADConnectorsDataSourceSchema()),
	}
}

func (d *adConnectionsDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (d *adConnectionsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var state ADConnectionDataSourceModel

	configDiagnostics := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(configDiagnostics...)
	if resp.Diagnostics.HasError() {
		return
	}

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
			log.Printf("[ERROR] HTTP error while creating AD Connector: %s", httpResp.Status)
			var fetchResp map[string]interface{}
			if err := json.NewDecoder(httpResp.Body).Decode(&fetchResp); err != nil {
				resp.Diagnostics.AddError("Failed to decode error response", err.Error())
				return
			}
			resp.Diagnostics.AddError(
				"HTTP Error",
				fmt.Sprintf("HTTP error while creating AD Connector for the reasons: %s", fetchResp["msg"]),
			)

		} else {
			log.Printf("[ERROR] API Call Failed: %v", err)
			resp.Diagnostics.AddError("API Call Failed", fmt.Sprintf("Error: %v", err))
		}
		return
	}

	if apiResp != nil && apiResp.ADConnectionResponse == nil {
		error := "Verify the connection type"
		log.Printf("[ERROR]: Verify the connection type given")
		resp.Diagnostics.AddError("Read of AD connection failed", error)
		return
	}

	log.Printf("[DEBUG] HTTP Status Code: %d", httpResp.StatusCode)

	state.Msg = util.SafeStringDatasource(apiResp.ADConnectionResponse.Msg)
	state.ErrorCode = util.SafeInt64(apiResp.ADConnectionResponse.Errorcode)
	state.ConnectionName = util.SafeStringDatasource(apiResp.ADConnectionResponse.Connectionname)
	state.ConnectionKey = util.SafeInt64(apiResp.ADConnectionResponse.Connectionkey)
	state.Description = util.SafeStringDatasource(apiResp.ADConnectionResponse.Description)
	state.DefaultSavRoles = util.SafeStringDatasource(apiResp.ADConnectionResponse.Defaultsavroles)
	state.ConnectionType = util.SafeStringDatasource(apiResp.ADConnectionResponse.Connectiontype)
	state.CreatedOn = util.SafeStringDatasource(apiResp.ADConnectionResponse.Createdon)
	state.CreatedBy = util.SafeStringDatasource(apiResp.ADConnectionResponse.Createdby)
	state.UpdatedBy = util.SafeStringDatasource(apiResp.ADConnectionResponse.Updatedby)
	state.Msg = util.SafeStringDatasource(apiResp.ADConnectionResponse.Msg)
	state.EmailTemplate = util.SafeStringDatasource(apiResp.ADConnectionResponse.Emailtemplate)

	if apiResp.ADConnectionResponse.Connectionattributes != nil {
		state.ConnectionAttributes = &ADConnectionAttributes{
			URL:                       util.SafeStringDatasource(apiResp.ADConnectionResponse.Connectionattributes.URL),
			ConnectionType:            util.SafeStringDatasource(apiResp.ADConnectionResponse.Connectionattributes.ConnectionType),
			AdvSearch:                 util.SafeStringDatasource(apiResp.ADConnectionResponse.Connectionattributes.ADVSEARCH),
			LastImportTime:            util.SafeStringDatasource(apiResp.ADConnectionResponse.Connectionattributes.LAST_IMPORT_TIME),
			CreateAccountJSON:         util.SafeStringDatasource(apiResp.ADConnectionResponse.Connectionattributes.CREATEACCOUNTJSON),
			DisableAccountJSON:        util.SafeStringDatasource(apiResp.ADConnectionResponse.Connectionattributes.DISABLEACCOUNTJSON),
			GroupSearchBaseDN:         util.SafeStringDatasource(apiResp.ADConnectionResponse.Connectionattributes.GroupSearchBaseDN),
			PasswordNoOfSplChars:      util.SafeStringDatasource(apiResp.ADConnectionResponse.Connectionattributes.PASSWORD_NOOFSPLCHARS),
			PasswordNoOfDigits:        util.SafeStringDatasource(apiResp.ADConnectionResponse.Connectionattributes.PASSWORD_NOOFDIGITS),
			StatusKeyJSON:             util.SafeStringDatasource(apiResp.ADConnectionResponse.Connectionattributes.STATUSKEYJSON),
			SearchFilter:              util.SafeStringDatasource(apiResp.ADConnectionResponse.Connectionattributes.SEARCHFILTER),
			ConfigJSON:                util.SafeStringDatasource(apiResp.ADConnectionResponse.Connectionattributes.ConfigJSON),
			RemoveAccountAction:       util.SafeStringDatasource(apiResp.ADConnectionResponse.Connectionattributes.REMOVEACCOUNTACTION),
			AccountAttribute:          util.SafeStringDatasource(apiResp.ADConnectionResponse.Connectionattributes.ACCOUNT_ATTRIBUTE),
			AccountNameRule:           util.SafeStringDatasource(apiResp.ADConnectionResponse.Connectionattributes.ACCOUNTNAMERULE),
			LDAPOrAD:                  util.SafeStringDatasource(apiResp.ADConnectionResponse.Connectionattributes.LDAP_OR_AD),
			EntitlementAttribute:      util.SafeStringDatasource(apiResp.ADConnectionResponse.Connectionattributes.ENTITLEMENT_ATTRIBUTE),
			SetRandomPassword:         util.SafeStringDatasource(apiResp.ADConnectionResponse.Connectionattributes.SETRANDOMPASSWORD),
			PasswordMinLength:         util.SafeStringDatasource(apiResp.ADConnectionResponse.Connectionattributes.PASSWORD_MIN_LENGTH),
			PasswordMaxLength:         util.SafeStringDatasource(apiResp.ADConnectionResponse.Connectionattributes.PASSWORD_MAX_LENGTH),
			PasswordNoOfCapsAlpha:     util.SafeStringDatasource(apiResp.ADConnectionResponse.Connectionattributes.PASSWORD_NOOFCAPSALPHA),
			SetDefaultPageSize:        util.SafeStringDatasource(apiResp.ADConnectionResponse.Connectionattributes.SETDEFAULTPAGESIZE),
			IsTimeoutSupported:        util.SafeBoolDatasource(apiResp.ADConnectionResponse.Connectionattributes.IsTimeoutSupported),
			ReuseInactiveAccount:      util.SafeStringDatasource(apiResp.ADConnectionResponse.Connectionattributes.REUSEINACTIVEACCOUNT),
			ImportJSON:                util.SafeStringDatasource(apiResp.ADConnectionResponse.Connectionattributes.IMPORTJSON),
			CreateUpdateMappings:      util.SafeStringDatasource(apiResp.ADConnectionResponse.Connectionattributes.CreateUpdateMappings),
			AdvanceFilterJSON:         util.SafeStringDatasource(apiResp.ADConnectionResponse.Connectionattributes.ADVANCE_FILTER_JSON),
			OrgImportJSON:             util.SafeStringDatasource(apiResp.ADConnectionResponse.Connectionattributes.ORGIMPORTJSON),
			PAMConfig:                 util.SafeStringDatasource(apiResp.ADConnectionResponse.Connectionattributes.PAM_CONFIG),
			PageSize:                  util.SafeStringDatasource(apiResp.ADConnectionResponse.Connectionattributes.PAGE_SIZE),
			Base:                      util.SafeStringDatasource(apiResp.ADConnectionResponse.Connectionattributes.BASE),
			DCLocator:                 util.SafeStringDatasource(apiResp.ADConnectionResponse.Connectionattributes.DC_LOCATOR),
			StatusThresholdConfig:     util.SafeStringDatasource(apiResp.ADConnectionResponse.Connectionattributes.STATUS_THRESHOLD_CONFIG),
			ResetAndChangePasswdJSON:  util.SafeStringDatasource(apiResp.ADConnectionResponse.Connectionattributes.RESETANDCHANGEPASSWRDJSON),
			SupportEmptyString:        util.SafeStringDatasource(apiResp.ADConnectionResponse.Connectionattributes.SUPPORTEMPTYSTRING),
			ReadOperationalAttributes: util.SafeStringDatasource(apiResp.ADConnectionResponse.Connectionattributes.READ_OPERATIONAL_ATTRIBUTES),
			EnableAccountJSON:         util.SafeStringDatasource(apiResp.ADConnectionResponse.Connectionattributes.ENABLEACCOUNTJSON),
			UserAttribute:             util.SafeStringDatasource(apiResp.ADConnectionResponse.Connectionattributes.USER_ATTRIBUTE),
			DefaultUserRole:           util.SafeStringDatasource(apiResp.ADConnectionResponse.Connectionattributes.DEFAULT_USER_ROLE),
			EndpointsFilter:           util.SafeStringDatasource(apiResp.ADConnectionResponse.Connectionattributes.ENDPOINTS_FILTER),
			UpdateAccountJSON:         util.SafeStringDatasource(apiResp.ADConnectionResponse.Connectionattributes.UPDATEACCOUNTJSON),
			ReuseAccountJSON:          util.SafeStringDatasource(apiResp.ADConnectionResponse.Connectionattributes.REUSEACCOUNTJSON),
			EnforceTreeDeletion:       util.SafeStringDatasource(apiResp.ADConnectionResponse.Connectionattributes.ENFORCE_TREE_DELETION),
			Filter:                    util.SafeStringDatasource(apiResp.ADConnectionResponse.Connectionattributes.FILTER),
			ObjectFilter:              util.SafeStringDatasource(apiResp.ADConnectionResponse.Connectionattributes.OBJECTFILTER),
			UpdateUserJSON:            util.SafeStringDatasource(apiResp.ADConnectionResponse.Connectionattributes.UPDATEUSERJSON),
			SaveConnection:            util.SafeStringDatasource(apiResp.ADConnectionResponse.Connectionattributes.Saveconnection),
			SystemName:                util.SafeStringDatasource(apiResp.ADConnectionResponse.Connectionattributes.Systemname),
			GroupImportMapping:        util.SafeStringDatasource(apiResp.ADConnectionResponse.Connectionattributes.GroupImportMapping),
			UnlockAccountJSON:         util.SafeStringDatasource(apiResp.ADConnectionResponse.Connectionattributes.UNLOCKACCOUNTJSON),
			EnableGroupManagement:     util.SafeStringDatasource(apiResp.ADConnectionResponse.Connectionattributes.ENABLEGROUPMANAGEMENT),
			ModifyUserDataJSON:        util.SafeStringDatasource(apiResp.ADConnectionResponse.Connectionattributes.MODIFYUSERDATAJSON),
			OrgBase:                   util.SafeStringDatasource(apiResp.ADConnectionResponse.Connectionattributes.ORG_BASE),
			OrganizationAttribute:     util.SafeStringDatasource(apiResp.ADConnectionResponse.Connectionattributes.ORGANIZATION_ATTRIBUTE),
			CreateOrgJSON:             util.SafeStringDatasource(apiResp.ADConnectionResponse.Connectionattributes.CREATEORGJSON),
			UpdateOrgJSON:             util.SafeStringDatasource(apiResp.ADConnectionResponse.Connectionattributes.UPDATEORGJSON),
			MaxChangeNumber:           util.SafeStringDatasource(apiResp.ADConnectionResponse.Connectionattributes.MAX_CHANGENUMBER),
			IncrementalConfig:         util.SafeStringDatasource(apiResp.ADConnectionResponse.Connectionattributes.INCREMENTAL_CONFIG),
			CheckForUnique:            util.SafeStringDatasource(apiResp.ADConnectionResponse.Connectionattributes.CHECKFORUNIQUE),
			IsTimeoutConfigValidated:  util.SafeBoolDatasource(apiResp.ADConnectionResponse.Connectionattributes.IsTimeoutConfigValidated),
		}
		if apiResp.ADConnectionResponse.Connectionattributes.ConnectionTimeoutConfig != nil {
			state.ConnectionAttributes.ConnectionTimeoutConfig = &ConnectionTimeoutConfig{
				RetryWait:               util.SafeInt64(apiResp.ADConnectionResponse.Connectionattributes.ConnectionTimeoutConfig.RetryWait),
				TokenRefreshMaxTryCount: util.SafeInt64(apiResp.ADConnectionResponse.Connectionattributes.ConnectionTimeoutConfig.TokenRefreshMaxTryCount),
				RetryWaitMaxValue:       util.SafeInt64(apiResp.ADConnectionResponse.Connectionattributes.ConnectionTimeoutConfig.RetryWaitMaxValue),
				RetryCount:              util.SafeInt64(apiResp.ADConnectionResponse.Connectionattributes.ConnectionTimeoutConfig.RetryCount),
				ReadTimeout:             util.SafeInt64(apiResp.ADConnectionResponse.Connectionattributes.ConnectionTimeoutConfig.ReadTimeout),
				ConnectionTimeout:       util.SafeInt64(apiResp.ADConnectionResponse.Connectionattributes.ConnectionTimeoutConfig.ConnectionTimeout),
				RetryFailureStatusCode:  util.SafeInt64(apiResp.ADConnectionResponse.Connectionattributes.ConnectionTimeoutConfig.RetryFailureStatusCode),
			}
		}
	}
	if apiResp.ADConnectionResponse.Connectionattributes == nil {
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
