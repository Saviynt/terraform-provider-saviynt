/*
 * Copyright (c) 2025 Saviynt Inc.
 * All Rights Reserved.
 *
 * This software is the confidential and proprietary information of
 * Saviynt Inc. ("Confidential Information"). You shall not disclose,
 * use, or distribute such Confidential Information except in accordance
 * with the terms of the license agreement you entered into with Saviynt.
 *
 * SAVIYNT MAKES NO REPRESENTATIONS OR WARRANTIES ABOUT THE SUITABILITY OF
 * THE SOFTWARE, EITHER EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO
 * THE IMPLIED WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR
 * PURPOSE, OR NON-INFRINGEMENT.
 */

// saviynt_unix_connection_datasource retrieves unix connections details from the Saviynt Security Manager.
// The data source supports a single Read operation to look up an existing unix connections by name.
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

var _ datasource.DataSource = &unixConnectionDataSource{}

// UnixConnectionDataSource defines the data source
type unixConnectionDataSource struct {
	client *s.Client
	token  string
}

type UnixConnectionDataSourceModel struct {
	ID types.String `tfsdk:"id"`
	BaseConnectionDataSourceModel
	ConnectionAttributes *UnixConnectionAttributes `tfsdk:"connection_attributes"`
}

type UnixConnectionAttributes struct {
	GroupsFile types.String `tfsdk:"groups_file"`
	// SSHKey                           types.String `tfsdk:"ssh_key"`
	AccountEntitlementMappingCommand types.String `tfsdk:"account_entitlement_mapping_command"`
	RemoveAccessCommand              types.String `tfsdk:"remove_access_command"`
	PEMKeyFile                       types.String `tfsdk:"pem_key_file"`
	PassThroughConnectionDetails     types.String `tfsdk:"pass_through_connection_details"`
	DisableAccountCommand            types.String `tfsdk:"disable_account_command"`
	// ConnectionTimeoutConfig          *ConnectionTimeoutConfig `tfsdk:"connection_timeout_config"`
	PortNumber         types.String `tfsdk:"port_number"`
	ConnectionType     types.String `tfsdk:"connection_type"`
	CreateGroupCommand types.String `tfsdk:"create_group_command"`
	AccountsFile       types.String `tfsdk:"accounts_file"`
	// Passphrase                       types.String `tfsdk:"passphrase"`
	DeleteGroupCommand    types.String `tfsdk:"delete_group_command"`
	HostName              types.String `tfsdk:"host_name"`
	AddGroupOwnerCommand  types.String `tfsdk:"add_group_owner_command"`
	StatusThresholdConfig types.String `tfsdk:"status_threshold_config"`
	// Username                         types.String `tfsdk:"username"`
	InactiveLockAccount  types.String `tfsdk:"inactive_lock_account"`
	AddAccessCommand     types.String `tfsdk:"add_access_command"`
	UpdateAccountCommand types.String `tfsdk:"update_account_command"`
	// SSHPassThroughPassphrase         types.String `tfsdk:"ssh_pass_through_passphrase"`
	ShadowFile         types.String `tfsdk:"shadow_file"`
	IsTimeoutSupported types.Bool   `tfsdk:"is_timeout_supported"`
	// SSHPassThroughSSHKey             types.String `tfsdk:"ssh_pass_through_ssh_key"`
	ProvisionAccountCommand         types.String `tfsdk:"provision_account_command"`
	FirefighterIDGrantAccessCommand types.String `tfsdk:"firefighterid_grant_access_command"`
	UnlockAccountCommand            types.String `tfsdk:"unlock_account_command"`
	DeprovisionAccountCommand       types.String `tfsdk:"deprovision_account_command"`
	// ChangePasswordJSON               types.String `tfsdk:"change_passwrd_json"`
	// SSHPassThroughPassword           types.String `tfsdk:"ssh_pass_through_password"`
	FirefighterIDRevokeAccessCommand types.String `tfsdk:"firefighterid_revoke_access_command"`
	AddPrimaryGroupCommand           types.String `tfsdk:"add_primary_group_command"`
	IsTimeoutConfigValidated         types.Bool   `tfsdk:"is_timeout_config_validated"`
	LockAccountCommand               types.String `tfsdk:"lock_account_command"`
	// Password                         types.String `tfsdk:"password"`
	CustomConfigJSON     types.String `tfsdk:"custom_config_json"`
	EnableAccountCommand types.String `tfsdk:"enable_account_command"`
}

func NewUnixConnectionsDataSource() datasource.DataSource {
	return &unixConnectionDataSource{}
}

func (d *unixConnectionDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "saviynt_unix_connection_datasource"
}

func UnixConnectorsDataSourceSchema() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:    true,
			Description: "Resource ID.",
		},
		"connection_attributes": schema.SingleNestedAttribute{
			Computed: true,
			Attributes: map[string]schema.Attribute{
				"groups_file": schema.StringAttribute{Computed: true},
				// "ssh_key":                             schema.StringAttribute{Computed: true},
				"account_entitlement_mapping_command": schema.StringAttribute{Computed: true},
				"remove_access_command":               schema.StringAttribute{Computed: true},
				"pem_key_file":                        schema.StringAttribute{Computed: true},
				"pass_through_connection_details":     schema.StringAttribute{Computed: true},
				"disable_account_command":             schema.StringAttribute{Computed: true},
				"port_number":                         schema.StringAttribute{Computed: true},
				"connection_type":                     schema.StringAttribute{Computed: true},
				"create_group_command":                schema.StringAttribute{Computed: true},
				"accounts_file":                       schema.StringAttribute{Computed: true},
				// "passphrase":                          schema.StringAttribute{Computed: true},
				"delete_group_command":    schema.StringAttribute{Computed: true},
				"host_name":               schema.StringAttribute{Computed: true},
				"add_group_owner_command": schema.StringAttribute{Computed: true},
				"status_threshold_config": schema.StringAttribute{Computed: true},
				// "username":                schema.StringAttribute{Computed: true},
				"inactive_lock_account":  schema.StringAttribute{Computed: true},
				"add_access_command":     schema.StringAttribute{Computed: true},
				"update_account_command": schema.StringAttribute{Computed: true},
				// "ssh_pass_through_passphrase":         schema.StringAttribute{Computed: true},
				"shadow_file":          schema.StringAttribute{Computed: true},
				"is_timeout_supported": schema.BoolAttribute{Computed: true},
				// "ssh_pass_through_ssh_key":            schema.StringAttribute{Computed: true},
				"provision_account_command":          schema.StringAttribute{Computed: true},
				"firefighterid_grant_access_command": schema.StringAttribute{Computed: true},
				"unlock_account_command":             schema.StringAttribute{Computed: true},
				"deprovision_account_command":        schema.StringAttribute{Computed: true},
				// "change_passwrd_json":                 schema.StringAttribute{Computed: true},
				// "ssh_pass_through_password":           schema.StringAttribute{Computed: true},
				"firefighterid_revoke_access_command": schema.StringAttribute{Computed: true},
				"add_primary_group_command":           schema.StringAttribute{Computed: true},
				"is_timeout_config_validated":         schema.BoolAttribute{Computed: true},
				"lock_account_command":                schema.StringAttribute{Computed: true},
				// "password":                            schema.StringAttribute{Computed: true},
				"custom_config_json":     schema.StringAttribute{Computed: true},
				"enable_account_command": schema.StringAttribute{Computed: true},
				// "connection_timeout_config": schema.SingleNestedAttribute{
				// 	Computed:   true,
				// 	Attributes: ConnectionTimeoutConfigeSchema(),
				// },
			},
		},
	}
}

// Schema defines the attributes for the data source
func (d *unixConnectionDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: util.UnixConnDataSourceDescription,
		Attributes:  connectionsutil.MergeDataSourceAttributes(BaseConnectorDataSourceSchema(), UnixConnectorsDataSourceSchema()),
	}
}

func (d *unixConnectionDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (d *unixConnectionDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var state UnixConnectionDataSourceModel

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
			log.Printf("[ERROR] HTTP error while creating Unix Connector: %s", httpResp.Status)
			var fetchResp map[string]interface{}
			if err := json.NewDecoder(httpResp.Body).Decode(&fetchResp); err != nil {
				resp.Diagnostics.AddError("Failed to decode error response", err.Error())
				return
			}
			resp.Diagnostics.AddError(
				"HTTP Error",
				fmt.Sprintf("HTTP error while creating Unix Connector for the reasons: %s", fetchResp["msg"]),
			)

		} else {
			log.Printf("[ERROR] API Call Failed: %v", err)
			resp.Diagnostics.AddError("API Call Failed", fmt.Sprintf("Error: %v", err))
		}
		return
	}
	// if apiResp != nil && *apiResp.UNIXConnectionResponse.Errorcode != 0 {
	// 	log.Printf("[ERROR]: Error in reading Unix connection. Errorcode: %v, Message: %v", *apiResp.UNIXConnectionResponse.Errorcode, *apiResp.UNIXConnectionResponse.Msg)
	// 	resp.Diagnostics.AddError("Reading Unix connection failed", *apiResp.UNIXConnectionResponse.Msg)
	// 	return
	// }

	if apiResp != nil && apiResp.UNIXConnectionResponse == nil {
		error := "Verify the connection type"
		log.Printf("[ERROR]: Verify the connection type given")
		resp.Diagnostics.AddError("Read of Unix connection failed", error)
		return
	}
	log.Printf("[DEBUG] HTTP Status Code: %d", httpResp.StatusCode)

	state.Msg = util.SafeStringDatasource(apiResp.UNIXConnectionResponse.Msg)
	state.ErrorCode = util.SafeInt64(apiResp.UNIXConnectionResponse.Errorcode)
	state.ConnectionName = util.SafeStringDatasource(apiResp.UNIXConnectionResponse.Connectionname)
	state.ConnectionKey = util.SafeInt64(apiResp.UNIXConnectionResponse.Connectionkey)
	state.Description = util.SafeStringDatasource(apiResp.UNIXConnectionResponse.Description)
	state.DefaultSavRoles = util.SafeStringDatasource(apiResp.UNIXConnectionResponse.Defaultsavroles)
	state.ConnectionType = util.SafeStringDatasource(apiResp.UNIXConnectionResponse.Connectiontype)
	state.CreatedOn = util.SafeStringDatasource(apiResp.UNIXConnectionResponse.Createdon)
	state.CreatedBy = util.SafeStringDatasource(apiResp.UNIXConnectionResponse.Createdby)
	state.UpdatedBy = util.SafeStringDatasource(apiResp.UNIXConnectionResponse.Updatedby)
	state.EmailTemplate = util.SafeStringDatasource(apiResp.UNIXConnectionResponse.Emailtemplate)

	if apiResp.UNIXConnectionResponse.Connectionattributes != nil {
		state.ConnectionAttributes = &UnixConnectionAttributes{
			GroupsFile: util.SafeStringDatasource(apiResp.UNIXConnectionResponse.Connectionattributes.GROUPS_FILE),
			// SSHKey:                           util.SafeStringDatasource(apiResp.UNIXConnectionResponse.Connectionattributes.SSH_KEY),
			AccountEntitlementMappingCommand: util.SafeStringDatasource(apiResp.UNIXConnectionResponse.Connectionattributes.ACCOUNT_ENTITLEMENT_MAPPING_COMMAND),
			RemoveAccessCommand:              util.SafeStringDatasource(apiResp.UNIXConnectionResponse.Connectionattributes.REMOVE_ACCESS_COMMAND),
			PEMKeyFile:                       util.SafeStringDatasource(apiResp.UNIXConnectionResponse.Connectionattributes.PEM_KEY_FILE),
			PassThroughConnectionDetails:     util.SafeStringDatasource(apiResp.UNIXConnectionResponse.Connectionattributes.PassThroughConnectionDetails),
			DisableAccountCommand:            util.SafeStringDatasource(apiResp.UNIXConnectionResponse.Connectionattributes.DISABLE_ACCOUNT_COMMAND),
			PortNumber:                       util.SafeStringDatasource(apiResp.UNIXConnectionResponse.Connectionattributes.PORT_NUMBER),
			ConnectionType:                   util.SafeStringDatasource(apiResp.UNIXConnectionResponse.Connectionattributes.ConnectionType),
			CreateGroupCommand:               util.SafeStringDatasource(apiResp.UNIXConnectionResponse.Connectionattributes.CREATE_GROUP_COMMAND),
			AccountsFile:                     util.SafeStringDatasource(apiResp.UNIXConnectionResponse.Connectionattributes.ACCOUNTS_FILE),
			// Passphrase:                       util.SafeStringDatasource(apiResp.UNIXConnectionResponse.Connectionattributes.PASSPHRASE),
			DeleteGroupCommand:    util.SafeStringDatasource(apiResp.UNIXConnectionResponse.Connectionattributes.DELETE_GROUP_COMMAND),
			HostName:              util.SafeStringDatasource(apiResp.UNIXConnectionResponse.Connectionattributes.HOST_NAME),
			AddGroupOwnerCommand:  util.SafeStringDatasource(apiResp.UNIXConnectionResponse.Connectionattributes.ADD_GROUP_OWNER_COMMAND),
			StatusThresholdConfig: util.SafeStringDatasource(apiResp.UNIXConnectionResponse.Connectionattributes.STATUS_THRESHOLD_CONFIG),
			// Username:                         util.SafeStringDatasource(apiResp.UNIXConnectionResponse.Connectionattributes.USERNAME),
			InactiveLockAccount:  util.SafeStringDatasource(apiResp.UNIXConnectionResponse.Connectionattributes.INACTIVE_LOCK_ACCOUNT),
			AddAccessCommand:     util.SafeStringDatasource(apiResp.UNIXConnectionResponse.Connectionattributes.ADD_ACCESS_COMMAND),
			UpdateAccountCommand: util.SafeStringDatasource(apiResp.UNIXConnectionResponse.Connectionattributes.UPDATE_ACCOUNT_COMMAND),
			// SSHPassThroughPassphrase:         util.SafeStringDatasource(apiResp.UNIXConnectionResponse.Connectionattributes.SSHPassThroughPassphrase),
			ShadowFile:         util.SafeStringDatasource(apiResp.UNIXConnectionResponse.Connectionattributes.SHADOW_FILE),
			IsTimeoutSupported: util.SafeBoolDatasource(apiResp.UNIXConnectionResponse.Connectionattributes.IsTimeoutSupported),
			// SSHPassThroughSSHKey:             util.SafeStringDatasource(apiResp.UNIXConnectionResponse.Connectionattributes.SSHPassThroughSSHKEY),
			ProvisionAccountCommand:         util.SafeStringDatasource(apiResp.UNIXConnectionResponse.Connectionattributes.PROVISION_ACCOUNT_COMMAND),
			FirefighterIDGrantAccessCommand: util.SafeStringDatasource(apiResp.UNIXConnectionResponse.Connectionattributes.FIREFIGHTERID_GRANT_ACCESS_COMMAND),
			UnlockAccountCommand:            util.SafeStringDatasource(apiResp.UNIXConnectionResponse.Connectionattributes.UNLOCK_ACCOUNT_COMMAND),
			DeprovisionAccountCommand:       util.SafeStringDatasource(apiResp.UNIXConnectionResponse.Connectionattributes.DEPROVISION_ACCOUNT_COMMAND),
			// ChangePasswordJSON:               util.SafeStringDatasource(apiResp.UNIXConnectionResponse.Connectionattributes.CHANGE_PASSWRD_JSON),
			// SSHPassThroughPassword:           util.SafeStringDatasource(apiResp.UNIXConnectionResponse.Connectionattributes.SSHPassThroughPassword),
			FirefighterIDRevokeAccessCommand: util.SafeStringDatasource(apiResp.UNIXConnectionResponse.Connectionattributes.FIREFIGHTERID_REVOKE_ACCESS_COMMAND),
			AddPrimaryGroupCommand:           util.SafeStringDatasource(apiResp.UNIXConnectionResponse.Connectionattributes.ADD_PRIMARY_GROUP_COMMAND),
			IsTimeoutConfigValidated:         util.SafeBoolDatasource(apiResp.UNIXConnectionResponse.Connectionattributes.IsTimeoutConfigValidated),
			LockAccountCommand:               util.SafeStringDatasource(apiResp.UNIXConnectionResponse.Connectionattributes.LOCK_ACCOUNT_COMMAND),
			// Password:                         util.SafeStringDatasource(apiResp.UNIXConnectionResponse.Connectionattributes.PASSWORD),
			CustomConfigJSON:     util.SafeStringDatasource(apiResp.UNIXConnectionResponse.Connectionattributes.CUSTOM_CONFIG_JSON),
			EnableAccountCommand: util.SafeStringDatasource(apiResp.UNIXConnectionResponse.Connectionattributes.ENABLE_ACCOUNT_COMMAND),
		}
		// if apiResp.UNIXConnectionResponse.Connectionattributes.ConnectionTimeoutConfig != nil {
		// 	state.ConnectionAttributes.ConnectionTimeoutConfig = &ConnectionTimeoutConfig{
		// 		RetryWait:               util.SafeInt64(apiResp.UNIXConnectionResponse.Connectionattributes.ConnectionTimeoutConfig.RetryWait),
		// 		TokenRefreshMaxTryCount: util.SafeInt64(apiResp.UNIXConnectionResponse.Connectionattributes.ConnectionTimeoutConfig.TokenRefreshMaxTryCount),
		// 		RetryFailureStatusCode:  util.SafeInt64(apiResp.UNIXConnectionResponse.Connectionattributes.ConnectionTimeoutConfig.RetryFailureStatusCode),
		// 		RetryWaitMaxValue:       util.SafeInt64(apiResp.UNIXConnectionResponse.Connectionattributes.ConnectionTimeoutConfig.RetryWaitMaxValue),
		// 		RetryCount:              util.SafeInt64(apiResp.UNIXConnectionResponse.Connectionattributes.ConnectionTimeoutConfig.RetryCount),
		// 		ReadTimeout:             util.SafeInt64(apiResp.UNIXConnectionResponse.Connectionattributes.ConnectionTimeoutConfig.ReadTimeout),
		// 		ConnectionTimeout:       util.SafeInt64(apiResp.UNIXConnectionResponse.Connectionattributes.ConnectionTimeoutConfig.ConnectionTimeout),
		// 	}
		// }
	}

	if apiResp.UNIXConnectionResponse.Connectionattributes == nil {
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
