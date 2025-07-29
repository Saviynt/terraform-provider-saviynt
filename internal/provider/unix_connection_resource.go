// Copyright (c) 2025 Saviynt Inc.
// SPDX-License-Identifier: MPL-2.0

// saviynt_unix_connection_resource manages Unix connectors in the Saviynt Security Manager.
// The resource implements the full Terraform lifecycle:
//   - Create: provisions a new Unix connector using the supplied configuration.
//   - Read: fetches the current connector state from Saviynt to keep Terraform’s state in sync.
//   - Update: applies any configuration changes to an existing connector.
//   - Import: brings an existing connector under Terraform management by its name.
package provider

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"terraform-provider-Saviynt/util"
	connectionsutil "terraform-provider-Saviynt/util/connectionsutil"

	s "github.com/saviynt/saviynt-api-go-client"
	openapi "github.com/saviynt/saviynt-api-go-client/connections"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &unixConnectionResource{}
var _ resource.ResourceWithImportState = &unixConnectionResource{}

type UnixConnectorResourceModel struct {
	BaseConnectorResourceModel
	ID                            types.String `tfsdk:"id"`
	HostName                      types.String `tfsdk:"host_name"`
	PortNumber                    types.String `tfsdk:"port_number"`
	Username                      types.String `tfsdk:"username"`
	Password                      types.String `tfsdk:"password"`
	GroupsFile                    types.String `tfsdk:"groups_file"`
	AccountsFile                  types.String `tfsdk:"accounts_file"`
	ShadowFile                    types.String `tfsdk:"shadow_file"`
	ProvisionAccountCommand       types.String `tfsdk:"provision_account_command"`
	DeprovisionAccountCommand     types.String `tfsdk:"deprovision_account_command"`
	AddAccessCommand              types.String `tfsdk:"add_access_command"`
	RemoveAccessCommand           types.String `tfsdk:"remove_access_command"`
	ChangePasswordJSON            types.String `tfsdk:"change_password_json"`
	PemKeyFile                    types.String `tfsdk:"pem_key_file"`
	EnableAccountCommand          types.String `tfsdk:"enable_account_command"`
	DisableAccountCommand         types.String `tfsdk:"disable_account_command"`
	AccountEntitlementMappingCmd  types.String `tfsdk:"account_entitlement_mapping_command"`
	Passphrase                    types.String `tfsdk:"passphrase"`
	UpdateAccountCommand          types.String `tfsdk:"update_account_command"`
	CreateGroupCommand            types.String `tfsdk:"create_group_command"`
	DeleteGroupCommand            types.String `tfsdk:"delete_group_command"`
	AddGroupOwnerCommand          types.String `tfsdk:"add_group_owner_command"`
	AddPrimaryGroupCommand        types.String `tfsdk:"add_primary_group_command"`
	FirefighterGrantAccessCommand types.String `tfsdk:"fire_fighter_id_grant_access_command"`
	FirefighterRevokeAccessCmd    types.String `tfsdk:"fire_fighter_id_revoke_access_command"`
	InactiveLockAccount           types.String `tfsdk:"inactive_lock_account"`
	StatusThresholdConfig         types.String `tfsdk:"status_threshold_config"`
	CustomConfigJSON              types.String `tfsdk:"custom_config_json"`
	SSHKey                        types.String `tfsdk:"ssh_key"`
	LockAccountCommand            types.String `tfsdk:"lock_account_command"`
	UnlockAccountCommand          types.String `tfsdk:"unlock_account_command"`
	PassThroughConnectionDetails  types.String `tfsdk:"pass_through_connection_details"`
	SSHPassThroughPassword        types.String `tfsdk:"ssh_pass_through_password"`
	SSHPassThroughSSHKEY          types.String `tfsdk:"ssh_pass_through_sshkey"`
	SSHPassThroughPassphrase      types.String `tfsdk:"ssh_pass_through_passphrase"`
}

type unixConnectionResource struct {
	client *s.Client
	token  string
}

func NewUnixTestConnectionResource() resource.Resource {
	return &unixConnectionResource{}
}

func (r *unixConnectionResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = "saviynt_unix_connection_resource"
}

func UnixConnectorResourceSchema() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:    true,
			Description: "Resource ID.",
		},
		"host_name": schema.StringAttribute{
			Required:    true,
			Description: "Property for HOST_NAME",
		},
		"port_number": schema.StringAttribute{
			Required:    true,
			Description: "Property for PORT_NUMBER",
		},
		"username": schema.StringAttribute{
			Required:    true,
			WriteOnly:   true,
			Description: "Property for USERNAME",
		},
		"password": schema.StringAttribute{
			Optional:    true,
			WriteOnly:   true,
			Description: "Property for PASSWORD",
		},
		"groups_file": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "Property for GROUPS_FILE",
		},
		"accounts_file": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "Property for ACCOUNTS_FILE",
		},
		"shadow_file": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "Property for SHADOW_FILE",
		},
		"provision_account_command": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "Property for PROVISION_ACCOUNT_COMMAND",
		},
		"deprovision_account_command": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "Property for DEPROVISION_ACCOUNT_COMMAND",
		},
		"add_access_command": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "Property for ADD_ACCESS_COMMAND",
		},
		"remove_access_command": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "Property for REMOVE_ACCESS_COMMAND",
		},
		"change_password_json": schema.StringAttribute{
			Optional:    true,
			WriteOnly:   true,
			Description: "Property for CHANGE_PASSWRD_JSON",
		},
		"pem_key_file": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "Property for PEM_KEY_FILE",
		},
		"enable_account_command": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "Property for ENABLE_ACCOUNT_COMMAND",
		},
		"disable_account_command": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "Property for DISABLE_ACCOUNT_COMMAND",
		},
		"account_entitlement_mapping_command": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "Property for ACCOUNT_ENTITLEMENT_MAPPING_COMMAND",
		},
		"passphrase": schema.StringAttribute{
			Optional:    true,
			WriteOnly:   true,
			Description: "Property for PASSPHRASE",
		},
		"update_account_command": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "Property for UPDATE_ACCOUNT_COMMAND",
		},
		"create_group_command": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "Property for CREATE_GROUP_COMMAND",
		},
		"delete_group_command": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "Property for DELETE_GROUP_COMMAND",
		},
		"add_group_owner_command": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "Property for ADD_GROUP_OWNER_COMMAND",
		},
		"add_primary_group_command": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "Property for ADD_PRIMARY_GROUP_COMMAND",
		},
		"fire_fighter_id_grant_access_command": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "Property for FIREFIGHTERID_GRANT_ACCESS_COMMAND",
		},
		"fire_fighter_id_revoke_access_command": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "Property for FIREFIGHTERID_REVOKE_ACCESS_COMMAND",
		},
		"inactive_lock_account": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "Property for INACTIVE_LOCK_ACCOUNT",
		},
		"status_threshold_config": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "Property for STATUS_THRESHOLD_CONFIG",
		},
		"custom_config_json": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "Property for CUSTOM_CONFIG_JSON",
		},
		"ssh_key": schema.StringAttribute{
			Optional:    true,
			WriteOnly:   true,
			Description: "Property for SSH_KEY",
		},
		"lock_account_command": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "Property for LOCK_ACCOUNT_COMMAND",
		},
		"unlock_account_command": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "Property for UNLOCK_ACCOUNT_COMMAND",
		},
		"pass_through_connection_details": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "Property for PassThroughConnectionDetails",
		},
		"ssh_pass_through_password": schema.StringAttribute{
			Optional:    true,
			WriteOnly:   true,
			Description: "Property for SSHPassThroughPassword",
		},
		"ssh_pass_through_sshkey": schema.StringAttribute{
			Optional:    true,
			WriteOnly:   true,
			Description: "Property for SSHPassThroughSSHKEY",
		},
		"ssh_pass_through_passphrase": schema.StringAttribute{
			Optional:    true,
			WriteOnly:   true,
			Description: "Property for SSHPassThroughPassphrase",
		},
	}
}

func (r *unixConnectionResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: util.UnixConnDescription,
		Attributes:  connectionsutil.MergeResourceAttributes(BaseConnectorResourceSchema(), UnixConnectorResourceSchema()),
	}
}

func (r *unixConnectionResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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
	r.client = prov.client
	r.token = prov.accessToken
}

func (r *unixConnectionResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan, config UnixConnectorResourceModel
	// Extract plan from request
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}
	//Extract config from request
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}
	cfg := openapi.NewConfiguration()
	apiBaseURL := strings.TrimPrefix(strings.TrimPrefix(r.client.APIBaseURL(), "https://"), "http://")
	cfg.Host = apiBaseURL
	cfg.Scheme = "https"
	cfg.AddDefaultHeader("Authorization", "Bearer "+r.token)
	cfg.HTTPClient = http.DefaultClient
	apiClient := openapi.NewAPIClient(cfg)

	reqParams := openapi.GetConnectionDetailsRequest{}
	reqParams.SetConnectionname(plan.ConnectionName.ValueString())
	// reqParams.SetConnectionkey(state.ConnectionKey.String())
	existingResource, _, _ := apiClient.ConnectionsAPI.GetConnectionDetails(ctx).GetConnectionDetailsRequest(reqParams).Execute()
	if existingResource != nil && existingResource.UNIXConnectionResponse != nil && existingResource.UNIXConnectionResponse.Errorcode != nil && *existingResource.UNIXConnectionResponse.Errorcode == 0 {
		log.Printf("[ERROR] Connection name already exists. Please import or use a different name")
		resp.Diagnostics.AddError("API Create Failed", "Connection name already exists. Please import or use a different name")
		return
	}

	unixConn := openapi.UNIXConnector{
		BaseConnector: openapi.BaseConnector{
			//required values
			Connectiontype: "Unix",
			ConnectionName: plan.ConnectionName.ValueString(),
			//optional values
			Description:     util.StringPointerOrEmpty(plan.Description),
			Defaultsavroles: util.StringPointerOrEmpty(plan.DefaultSavRoles),
			EmailTemplate:   util.StringPointerOrEmpty(plan.EmailTemplate),
		},
		//required values
		HOST_NAME:   plan.HostName.ValueString(),
		PORT_NUMBER: plan.PortNumber.ValueString(),
		USERNAME:    config.Username.ValueString(),
		//optional values
		PASSWORD:                            util.StringPointerOrEmpty(config.Password),
		GROUPS_FILE:                         util.StringPointerOrEmpty(plan.GroupsFile),
		ACCOUNTS_FILE:                       util.StringPointerOrEmpty(plan.AccountsFile),
		SHADOW_FILE:                         util.StringPointerOrEmpty(plan.ShadowFile),
		PROVISION_ACCOUNT_COMMAND:           util.StringPointerOrEmpty(plan.ProvisionAccountCommand),
		DEPROVISION_ACCOUNT_COMMAND:         util.StringPointerOrEmpty(plan.DeprovisionAccountCommand),
		ADD_ACCESS_COMMAND:                  util.StringPointerOrEmpty(plan.AddAccessCommand),
		REMOVE_ACCESS_COMMAND:               util.StringPointerOrEmpty(plan.RemoveAccessCommand),
		CHANGE_PASSWRD_JSON:                 util.StringPointerOrEmpty(config.ChangePasswordJSON),
		PEM_KEY_FILE:                        util.StringPointerOrEmpty(plan.PemKeyFile),
		ENABLE_ACCOUNT_COMMAND:              util.StringPointerOrEmpty(plan.EnableAccountCommand),
		DISABLE_ACCOUNT_COMMAND:             util.StringPointerOrEmpty(plan.DisableAccountCommand),
		ACCOUNT_ENTITLEMENT_MAPPING_COMMAND: util.StringPointerOrEmpty(plan.AccountEntitlementMappingCmd),
		PASSPHRASE:                          util.StringPointerOrEmpty(config.Passphrase),
		UPDATE_ACCOUNT_COMMAND:              util.StringPointerOrEmpty(plan.UpdateAccountCommand),
		CREATE_GROUP_COMMAND:                util.StringPointerOrEmpty(plan.CreateGroupCommand),
		DELETE_GROUP_COMMAND:                util.StringPointerOrEmpty(plan.DeleteGroupCommand),
		ADD_GROUP_OWNER_COMMAND:             util.StringPointerOrEmpty(plan.AddGroupOwnerCommand),
		ADD_PRIMARY_GROUP_COMMAND:           util.StringPointerOrEmpty(plan.AddPrimaryGroupCommand),
		FIREFIGHTERID_GRANT_ACCESS_COMMAND:  util.StringPointerOrEmpty(plan.FirefighterGrantAccessCommand),
		FIREFIGHTERID_REVOKE_ACCESS_COMMAND: util.StringPointerOrEmpty(plan.FirefighterRevokeAccessCmd),
		INACTIVE_LOCK_ACCOUNT:               util.StringPointerOrEmpty(plan.InactiveLockAccount),
		STATUS_THRESHOLD_CONFIG:             util.StringPointerOrEmpty(plan.StatusThresholdConfig),
		CUSTOM_CONFIG_JSON:                  util.StringPointerOrEmpty(plan.CustomConfigJSON),
		SSH_KEY:                             util.StringPointerOrEmpty(config.SSHKey),
		LOCK_ACCOUNT_COMMAND:                util.StringPointerOrEmpty(plan.LockAccountCommand),
		UNLOCK_ACCOUNT_COMMAND:              util.StringPointerOrEmpty(plan.UnlockAccountCommand),
		PassThroughConnectionDetails:        util.StringPointerOrEmpty(plan.PassThroughConnectionDetails),
		SSHPassThroughPassword:              util.StringPointerOrEmpty(config.SSHPassThroughPassword),
		SSHPassThroughSSHKEY:                util.StringPointerOrEmpty(config.SSHPassThroughSSHKEY),
		SSHPassThroughPassphrase:            util.StringPointerOrEmpty(config.SSHPassThroughPassphrase),
	}
	if plan.VaultConnection.ValueString() != "" {
		unixConn.BaseConnector.VaultConnection = util.SafeStringConnector(plan.VaultConnection.ValueString())
		unixConn.BaseConnector.VaultConfiguration = util.SafeStringConnector(plan.VaultConfiguration.ValueString())
		unixConn.BaseConnector.Saveinvault = util.SafeStringConnector(plan.SaveInVault.ValueString())
	}
	unixRequest := openapi.CreateOrUpdateRequest{
		UNIXConnector: &unixConn,
	}

	apiResp, _, err := apiClient.ConnectionsAPI.CreateOrUpdate(ctx).CreateOrUpdateRequest(unixRequest).Execute()
	if err != nil {
		log.Printf("[ERROR] Failed to create API resource. Error: %v", err)
		resp.Diagnostics.AddError("API Create Failed", fmt.Sprintf("Error: %v", err))
		return
	}
	if apiResp != nil && *apiResp.ErrorCode != "0" {
		log.Printf("[ERROR]: Error in creating Unix connection resource. Errorcode: %v, Message: %v", *apiResp.ErrorCode, *apiResp.Msg)
		resp.Diagnostics.AddError("Creation of Unix connection failed", *apiResp.Msg)
		return
	}

	plan.ID = types.StringValue(fmt.Sprintf("%d", *apiResp.ConnectionKey))
	plan.ConnectionKey = types.Int64Value(int64(*apiResp.ConnectionKey))
	plan.Description = util.SafeStringDatasource(plan.Description.ValueStringPointer())
	plan.DefaultSavRoles = util.SafeStringDatasource(plan.DefaultSavRoles.ValueStringPointer())
	plan.EmailTemplate = util.SafeStringDatasource(plan.EmailTemplate.ValueStringPointer())
	plan.HostName = util.SafeStringDatasource(plan.HostName.ValueStringPointer())
	plan.PortNumber = util.SafeStringDatasource(plan.PortNumber.ValueStringPointer())
	plan.Username = util.SafeStringDatasource(plan.Username.ValueStringPointer())
	plan.GroupsFile = util.SafeStringDatasource(plan.GroupsFile.ValueStringPointer())
	plan.AccountsFile = util.SafeStringDatasource(plan.AccountsFile.ValueStringPointer())
	plan.ShadowFile = util.SafeStringDatasource(plan.ShadowFile.ValueStringPointer())
	plan.ProvisionAccountCommand = util.SafeStringDatasource(plan.ProvisionAccountCommand.ValueStringPointer())
	plan.DeprovisionAccountCommand = util.SafeStringDatasource(plan.DeprovisionAccountCommand.ValueStringPointer())
	plan.AddAccessCommand = util.SafeStringDatasource(plan.AddAccessCommand.ValueStringPointer())
	plan.RemoveAccessCommand = util.SafeStringDatasource(plan.RemoveAccessCommand.ValueStringPointer())
	plan.ChangePasswordJSON = util.SafeStringDatasource(plan.ChangePasswordJSON.ValueStringPointer())
	plan.PemKeyFile = util.SafeStringDatasource(plan.PemKeyFile.ValueStringPointer())
	plan.EnableAccountCommand = util.SafeStringDatasource(plan.EnableAccountCommand.ValueStringPointer())
	plan.DisableAccountCommand = util.SafeStringDatasource(plan.DisableAccountCommand.ValueStringPointer())
	plan.AccountEntitlementMappingCmd = util.SafeStringDatasource(plan.AccountEntitlementMappingCmd.ValueStringPointer())
	plan.UpdateAccountCommand = util.SafeStringDatasource(plan.UpdateAccountCommand.ValueStringPointer())
	plan.CreateGroupCommand = util.SafeStringDatasource(plan.CreateGroupCommand.ValueStringPointer())
	plan.DeleteGroupCommand = util.SafeStringDatasource(plan.DeleteGroupCommand.ValueStringPointer())
	plan.AddGroupOwnerCommand = util.SafeStringDatasource(plan.AddGroupOwnerCommand.ValueStringPointer())
	plan.AddPrimaryGroupCommand = util.SafeStringDatasource(plan.AddPrimaryGroupCommand.ValueStringPointer())
	plan.FirefighterGrantAccessCommand = util.SafeStringDatasource(plan.FirefighterGrantAccessCommand.ValueStringPointer())
	plan.FirefighterRevokeAccessCmd = util.SafeStringDatasource(plan.FirefighterRevokeAccessCmd.ValueStringPointer())
	plan.InactiveLockAccount = util.SafeStringDatasource(plan.InactiveLockAccount.ValueStringPointer())
	plan.StatusThresholdConfig = util.SafeStringDatasource(plan.StatusThresholdConfig.ValueStringPointer())
	plan.CustomConfigJSON = util.SafeStringDatasource(plan.CustomConfigJSON.ValueStringPointer())
	plan.LockAccountCommand = util.SafeStringDatasource(plan.LockAccountCommand.ValueStringPointer())
	plan.UnlockAccountCommand = util.SafeStringDatasource(plan.UnlockAccountCommand.ValueStringPointer())
	plan.PassThroughConnectionDetails = util.SafeStringDatasource(plan.PassThroughConnectionDetails.ValueStringPointer())
	plan.Msg = types.StringValue(util.SafeDeref(apiResp.Msg))
	plan.ErrorCode = types.StringValue(util.SafeDeref(apiResp.ErrorCode))
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *unixConnectionResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state UnixConnectorResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		log.Println("Diagnostics contain errors, returning early.")
		return
	}

	// Configure API client
	cfg := openapi.NewConfiguration()
	apiBaseURL := strings.TrimPrefix(strings.TrimPrefix(r.client.APIBaseURL(), "https://"), "http://")
	cfg.Host = apiBaseURL
	cfg.Scheme = "https"
	cfg.AddDefaultHeader("Authorization", "Bearer "+r.token)
	cfg.HTTPClient = http.DefaultClient

	apiClient := openapi.NewAPIClient(cfg)
	reqParams := openapi.GetConnectionDetailsRequest{}
	reqParams.SetConnectionname(state.ConnectionName.ValueString())
	apiResp, _, err := apiClient.ConnectionsAPI.GetConnectionDetails(ctx).GetConnectionDetailsRequest(reqParams).Execute()
	if err != nil {
		log.Printf("Problem with the get function in read block")
		resp.Diagnostics.AddError("API Read Failed", fmt.Sprintf("Error: %v", err))
		return
	}
	if apiResp != nil && apiResp.UNIXConnectionResponse != nil && *apiResp.UNIXConnectionResponse.Errorcode != 0 {
		log.Printf("[ERROR]: Error in reading Unix connection resource. Errorcode: %v, Message: %v", *apiResp.UNIXConnectionResponse.Errorcode, *apiResp.UNIXConnectionResponse.Msg)
		resp.Diagnostics.AddError("Reading Unix connection resource failed", *apiResp.UNIXConnectionResponse.Msg)
		return
	}

	state.ConnectionKey = types.Int64Value(int64(*apiResp.UNIXConnectionResponse.Connectionkey))
	state.ID = types.StringValue(fmt.Sprintf("%d", *apiResp.UNIXConnectionResponse.Connectionkey))
	state.ConnectionName = util.SafeStringDatasource(apiResp.UNIXConnectionResponse.Connectionname)
	state.Description = util.SafeStringDatasource(apiResp.UNIXConnectionResponse.Description)
	state.DefaultSavRoles = util.SafeStringDatasource(apiResp.UNIXConnectionResponse.Defaultsavroles)
	state.EmailTemplate = util.SafeStringDatasource(apiResp.UNIXConnectionResponse.Emailtemplate)
	state.GroupsFile = util.SafeStringDatasource(apiResp.UNIXConnectionResponse.Connectionattributes.GROUPS_FILE)
	state.AccountEntitlementMappingCmd = util.SafeStringDatasource(apiResp.UNIXConnectionResponse.Connectionattributes.ACCOUNT_ENTITLEMENT_MAPPING_COMMAND)
	state.RemoveAccessCommand = util.SafeStringDatasource(apiResp.UNIXConnectionResponse.Connectionattributes.REMOVE_ACCESS_COMMAND)
	state.PemKeyFile = util.SafeStringDatasource(apiResp.UNIXConnectionResponse.Connectionattributes.PEM_KEY_FILE)
	state.PassThroughConnectionDetails = util.SafeStringDatasource(apiResp.UNIXConnectionResponse.Connectionattributes.PassThroughConnectionDetails)
	state.DisableAccountCommand = util.SafeStringDatasource(apiResp.UNIXConnectionResponse.Connectionattributes.DISABLE_ACCOUNT_COMMAND)
	state.PortNumber = util.SafeStringDatasource(apiResp.UNIXConnectionResponse.Connectionattributes.PORT_NUMBER)
	state.CreateGroupCommand = util.SafeStringDatasource(apiResp.UNIXConnectionResponse.Connectionattributes.CREATE_GROUP_COMMAND)
	state.AccountsFile = util.SafeStringDatasource(apiResp.UNIXConnectionResponse.Connectionattributes.ACCOUNTS_FILE)
	state.DeleteGroupCommand = util.SafeStringDatasource(apiResp.UNIXConnectionResponse.Connectionattributes.DELETE_GROUP_COMMAND)
	state.HostName = util.SafeStringDatasource(apiResp.UNIXConnectionResponse.Connectionattributes.HOST_NAME)
	state.AddGroupOwnerCommand = util.SafeStringDatasource(apiResp.UNIXConnectionResponse.Connectionattributes.ADD_GROUP_OWNER_COMMAND)
	state.StatusThresholdConfig = util.SafeStringDatasource(apiResp.UNIXConnectionResponse.Connectionattributes.STATUS_THRESHOLD_CONFIG)
	state.Username = util.SafeStringDatasource(apiResp.UNIXConnectionResponse.Connectionattributes.USERNAME)
	state.InactiveLockAccount = util.SafeStringDatasource(apiResp.UNIXConnectionResponse.Connectionattributes.INACTIVE_LOCK_ACCOUNT)
	state.AddAccessCommand = util.SafeStringDatasource(apiResp.UNIXConnectionResponse.Connectionattributes.ADD_ACCESS_COMMAND)
	state.UpdateAccountCommand = util.SafeStringDatasource(apiResp.UNIXConnectionResponse.Connectionattributes.UPDATE_ACCOUNT_COMMAND)
	state.ShadowFile = util.SafeStringDatasource(apiResp.UNIXConnectionResponse.Connectionattributes.SHADOW_FILE)
	state.ProvisionAccountCommand = util.SafeStringDatasource(apiResp.UNIXConnectionResponse.Connectionattributes.PROVISION_ACCOUNT_COMMAND)
	state.FirefighterGrantAccessCommand = util.SafeStringDatasource(apiResp.UNIXConnectionResponse.Connectionattributes.FIREFIGHTERID_GRANT_ACCESS_COMMAND)
	state.UnlockAccountCommand = util.SafeStringDatasource(apiResp.UNIXConnectionResponse.Connectionattributes.UNLOCK_ACCOUNT_COMMAND)
	state.DeprovisionAccountCommand = util.SafeStringDatasource(apiResp.UNIXConnectionResponse.Connectionattributes.DEPROVISION_ACCOUNT_COMMAND)
	state.ChangePasswordJSON = util.SafeStringDatasource(apiResp.UNIXConnectionResponse.Connectionattributes.CHANGE_PASSWRD_JSON)
	state.FirefighterRevokeAccessCmd = util.SafeStringDatasource(apiResp.UNIXConnectionResponse.Connectionattributes.FIREFIGHTERID_REVOKE_ACCESS_COMMAND)
	state.AddPrimaryGroupCommand = util.SafeStringDatasource(apiResp.UNIXConnectionResponse.Connectionattributes.ADD_PRIMARY_GROUP_COMMAND)
	state.LockAccountCommand = util.SafeStringDatasource(apiResp.UNIXConnectionResponse.Connectionattributes.LOCK_ACCOUNT_COMMAND)
	state.CustomConfigJSON = util.SafeStringDatasource(apiResp.UNIXConnectionResponse.Connectionattributes.CUSTOM_CONFIG_JSON)
	state.EnableAccountCommand = util.SafeStringDatasource(apiResp.UNIXConnectionResponse.Connectionattributes.ENABLE_ACCOUNT_COMMAND)
	apiMessage := util.SafeDeref(apiResp.UNIXConnectionResponse.Msg)
	if apiMessage == "success" {
		state.Msg = types.StringValue("Connection Successful")
	} else {
		state.Msg = types.StringValue(apiMessage)
	}
	state.ErrorCode = util.Int32PtrToTFString(apiResp.UNIXConnectionResponse.Errorcode)
	stateDiagnostics := resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(stateDiagnostics...)
	if resp.Diagnostics.HasError() {
		log.Println("Diagnostics contain errors, returning early.")
		return
	}
}

func (r *unixConnectionResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state, config UnixConnectorResourceModel
	// Extract state from request
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
	// Extract plan from request
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}
	//Extract config from request
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}
	// Configure API client
	cfg := openapi.NewConfiguration()
	apiBaseURL := strings.TrimPrefix(strings.TrimPrefix(r.client.APIBaseURL(), "https://"), "http://")
	cfg.Host = apiBaseURL
	cfg.Scheme = "https"
	cfg.AddDefaultHeader("Authorization", "Bearer "+r.token)
	if plan.ConnectionName.ValueString() != state.ConnectionName.ValueString() {
		resp.Diagnostics.AddError("Error", "Connection name cannot be updated")
		log.Printf("[ERROR]: Connection name cannot be updated")
		return
	}

	cfg.HTTPClient = http.DefaultClient
	unixConn := openapi.UNIXConnector{
		BaseConnector: openapi.BaseConnector{
			//required values
			Connectiontype: "Unix",
			ConnectionName: plan.ConnectionName.ValueString(),
			//optional values
			Description:     util.StringPointerOrEmpty(plan.Description),
			Defaultsavroles: util.StringPointerOrEmpty(plan.DefaultSavRoles),
			EmailTemplate:   util.StringPointerOrEmpty(plan.EmailTemplate),
		},
		//required values
		HOST_NAME:   plan.HostName.ValueString(),
		PORT_NUMBER: plan.PortNumber.ValueString(),
		USERNAME:    config.Username.ValueString(),
		//optional values
		PASSWORD:                            util.StringPointerOrEmpty(config.Password),
		GROUPS_FILE:                         util.StringPointerOrEmpty(plan.GroupsFile),
		ACCOUNTS_FILE:                       util.StringPointerOrEmpty(plan.AccountsFile),
		SHADOW_FILE:                         util.StringPointerOrEmpty(plan.ShadowFile),
		PROVISION_ACCOUNT_COMMAND:           util.StringPointerOrEmpty(plan.ProvisionAccountCommand),
		DEPROVISION_ACCOUNT_COMMAND:         util.StringPointerOrEmpty(plan.DeprovisionAccountCommand),
		ADD_ACCESS_COMMAND:                  util.StringPointerOrEmpty(plan.AddAccessCommand),
		REMOVE_ACCESS_COMMAND:               util.StringPointerOrEmpty(plan.RemoveAccessCommand),
		CHANGE_PASSWRD_JSON:                 util.StringPointerOrEmpty(config.ChangePasswordJSON),
		PEM_KEY_FILE:                        util.StringPointerOrEmpty(plan.PemKeyFile),
		ENABLE_ACCOUNT_COMMAND:              util.StringPointerOrEmpty(plan.EnableAccountCommand),
		DISABLE_ACCOUNT_COMMAND:             util.StringPointerOrEmpty(plan.DisableAccountCommand),
		ACCOUNT_ENTITLEMENT_MAPPING_COMMAND: util.StringPointerOrEmpty(plan.AccountEntitlementMappingCmd),
		PASSPHRASE:                          util.StringPointerOrEmpty(config.Passphrase),
		UPDATE_ACCOUNT_COMMAND:              util.StringPointerOrEmpty(plan.UpdateAccountCommand),
		CREATE_GROUP_COMMAND:                util.StringPointerOrEmpty(plan.CreateGroupCommand),
		DELETE_GROUP_COMMAND:                util.StringPointerOrEmpty(plan.DeleteGroupCommand),
		ADD_GROUP_OWNER_COMMAND:             util.StringPointerOrEmpty(plan.AddGroupOwnerCommand),
		ADD_PRIMARY_GROUP_COMMAND:           util.StringPointerOrEmpty(plan.AddPrimaryGroupCommand),
		FIREFIGHTERID_GRANT_ACCESS_COMMAND:  util.StringPointerOrEmpty(plan.FirefighterGrantAccessCommand),
		FIREFIGHTERID_REVOKE_ACCESS_COMMAND: util.StringPointerOrEmpty(plan.FirefighterRevokeAccessCmd),
		INACTIVE_LOCK_ACCOUNT:               util.StringPointerOrEmpty(plan.InactiveLockAccount),
		STATUS_THRESHOLD_CONFIG:             util.StringPointerOrEmpty(plan.StatusThresholdConfig),
		CUSTOM_CONFIG_JSON:                  util.StringPointerOrEmpty(plan.CustomConfigJSON),
		SSH_KEY:                             util.StringPointerOrEmpty(config.SSHKey),
		LOCK_ACCOUNT_COMMAND:                util.StringPointerOrEmpty(plan.LockAccountCommand),
		UNLOCK_ACCOUNT_COMMAND:              util.StringPointerOrEmpty(plan.UnlockAccountCommand),
		PassThroughConnectionDetails:        util.StringPointerOrEmpty(plan.PassThroughConnectionDetails),
		SSHPassThroughPassword:              util.StringPointerOrEmpty(config.SSHPassThroughPassword),
		SSHPassThroughSSHKEY:                util.StringPointerOrEmpty(config.SSHPassThroughSSHKEY),
		SSHPassThroughPassphrase:            util.StringPointerOrEmpty(config.SSHPassThroughPassphrase),
	}
	if plan.VaultConnection.ValueString() != "" {
		unixConn.BaseConnector.VaultConnection = util.SafeStringConnector(plan.VaultConnection.ValueString())
		unixConn.BaseConnector.VaultConfiguration = util.SafeStringConnector(plan.VaultConfiguration.ValueString())
		unixConn.BaseConnector.Saveinvault = util.SafeStringConnector(plan.SaveInVault.ValueString())
	} else {
		emptyStr := ""
		unixConn.BaseConnector.VaultConnection = &emptyStr
		unixConn.BaseConnector.VaultConfiguration = &emptyStr
		unixConn.BaseConnector.Saveinvault = &emptyStr
	}
	unixConnRequest := openapi.CreateOrUpdateRequest{
		UNIXConnector: &unixConn,
	}

	// Initialize API client
	apiClient := openapi.NewAPIClient(cfg)
	apiResp, _, err := apiClient.ConnectionsAPI.CreateOrUpdate(ctx).CreateOrUpdateRequest(unixConnRequest).Execute()
	if err != nil {
		log.Printf("Problem with the update function")
		resp.Diagnostics.AddError("API Update Failed", fmt.Sprintf("Error: %v", err))
		return
	}
	if apiResp != nil && *apiResp.ErrorCode != "0" {
		log.Printf("[ERROR]: Error in updating Unix connection resource. Errorcode: %v, Message: %v", *apiResp.ErrorCode, *apiResp.Msg)
		resp.Diagnostics.AddError("Updation of Unix connection failed", *apiResp.Msg)
		return
	}

	reqParams := openapi.GetConnectionDetailsRequest{}

	reqParams.SetConnectionname(plan.ConnectionName.ValueString())
	getResp, _, err := apiClient.ConnectionsAPI.GetConnectionDetails(ctx).GetConnectionDetailsRequest(reqParams).Execute()
	if err != nil {
		log.Printf("Problem with the get function in update block")
		resp.Diagnostics.AddError("API Read Failed", fmt.Sprintf("Error: %v", err))
		return
	}
	if getResp != nil && getResp.UNIXConnectionResponse != nil && *getResp.UNIXConnectionResponse.Errorcode != 0 {
		log.Printf("[ERROR]: Error in reading Unix connection resource after updation. Errorcode: %v, Message: %v", *getResp.UNIXConnectionResponse.Errorcode, *getResp.UNIXConnectionResponse.Msg)
		resp.Diagnostics.AddError("Reading Unix connection after updation failed", *getResp.UNIXConnectionResponse.Msg)
		return
	}

	plan.ConnectionKey = types.Int64Value(int64(*getResp.UNIXConnectionResponse.Connectionkey))
	plan.ID = types.StringValue(fmt.Sprintf("%d", *getResp.UNIXConnectionResponse.Connectionkey))
	plan.ConnectionName = util.SafeStringDatasource(getResp.UNIXConnectionResponse.Connectionname)
	plan.Description = util.SafeStringDatasource(getResp.UNIXConnectionResponse.Description)
	plan.DefaultSavRoles = util.SafeStringDatasource(getResp.UNIXConnectionResponse.Defaultsavroles)
	plan.EmailTemplate = util.SafeStringDatasource(getResp.UNIXConnectionResponse.Emailtemplate)
	plan.GroupsFile = util.SafeStringDatasource(getResp.UNIXConnectionResponse.Connectionattributes.GROUPS_FILE)
	plan.AccountEntitlementMappingCmd = util.SafeStringDatasource(getResp.UNIXConnectionResponse.Connectionattributes.ACCOUNT_ENTITLEMENT_MAPPING_COMMAND)
	plan.RemoveAccessCommand = util.SafeStringDatasource(getResp.UNIXConnectionResponse.Connectionattributes.REMOVE_ACCESS_COMMAND)
	plan.PemKeyFile = util.SafeStringDatasource(getResp.UNIXConnectionResponse.Connectionattributes.PEM_KEY_FILE)
	plan.PassThroughConnectionDetails = util.SafeStringDatasource(getResp.UNIXConnectionResponse.Connectionattributes.PassThroughConnectionDetails)
	plan.DisableAccountCommand = util.SafeStringDatasource(getResp.UNIXConnectionResponse.Connectionattributes.DISABLE_ACCOUNT_COMMAND)
	plan.PortNumber = util.SafeStringDatasource(getResp.UNIXConnectionResponse.Connectionattributes.PORT_NUMBER)
	plan.CreateGroupCommand = util.SafeStringDatasource(getResp.UNIXConnectionResponse.Connectionattributes.CREATE_GROUP_COMMAND)
	plan.AccountsFile = util.SafeStringDatasource(getResp.UNIXConnectionResponse.Connectionattributes.ACCOUNTS_FILE)
	plan.DeleteGroupCommand = util.SafeStringDatasource(getResp.UNIXConnectionResponse.Connectionattributes.DELETE_GROUP_COMMAND)
	plan.HostName = util.SafeStringDatasource(getResp.UNIXConnectionResponse.Connectionattributes.HOST_NAME)
	plan.AddGroupOwnerCommand = util.SafeStringDatasource(getResp.UNIXConnectionResponse.Connectionattributes.ADD_GROUP_OWNER_COMMAND)
	plan.StatusThresholdConfig = util.SafeStringDatasource(getResp.UNIXConnectionResponse.Connectionattributes.STATUS_THRESHOLD_CONFIG)
	plan.Username = util.SafeStringDatasource(getResp.UNIXConnectionResponse.Connectionattributes.USERNAME)
	plan.InactiveLockAccount = util.SafeStringDatasource(getResp.UNIXConnectionResponse.Connectionattributes.INACTIVE_LOCK_ACCOUNT)
	plan.AddAccessCommand = util.SafeStringDatasource(getResp.UNIXConnectionResponse.Connectionattributes.ADD_ACCESS_COMMAND)
	plan.UpdateAccountCommand = util.SafeStringDatasource(getResp.UNIXConnectionResponse.Connectionattributes.UPDATE_ACCOUNT_COMMAND)
	plan.ShadowFile = util.SafeStringDatasource(getResp.UNIXConnectionResponse.Connectionattributes.SHADOW_FILE)
	plan.ProvisionAccountCommand = util.SafeStringDatasource(getResp.UNIXConnectionResponse.Connectionattributes.PROVISION_ACCOUNT_COMMAND)
	plan.FirefighterGrantAccessCommand = util.SafeStringDatasource(getResp.UNIXConnectionResponse.Connectionattributes.FIREFIGHTERID_GRANT_ACCESS_COMMAND)
	plan.UnlockAccountCommand = util.SafeStringDatasource(getResp.UNIXConnectionResponse.Connectionattributes.UNLOCK_ACCOUNT_COMMAND)
	plan.DeprovisionAccountCommand = util.SafeStringDatasource(getResp.UNIXConnectionResponse.Connectionattributes.DEPROVISION_ACCOUNT_COMMAND)
	plan.ChangePasswordJSON = util.SafeStringDatasource(getResp.UNIXConnectionResponse.Connectionattributes.CHANGE_PASSWRD_JSON)
	plan.FirefighterRevokeAccessCmd = util.SafeStringDatasource(getResp.UNIXConnectionResponse.Connectionattributes.FIREFIGHTERID_REVOKE_ACCESS_COMMAND)
	plan.AddPrimaryGroupCommand = util.SafeStringDatasource(getResp.UNIXConnectionResponse.Connectionattributes.ADD_PRIMARY_GROUP_COMMAND)
	plan.LockAccountCommand = util.SafeStringDatasource(getResp.UNIXConnectionResponse.Connectionattributes.LOCK_ACCOUNT_COMMAND)
	plan.CustomConfigJSON = util.SafeStringDatasource(getResp.UNIXConnectionResponse.Connectionattributes.CUSTOM_CONFIG_JSON)
	plan.EnableAccountCommand = util.SafeStringDatasource(getResp.UNIXConnectionResponse.Connectionattributes.ENABLE_ACCOUNT_COMMAND)
	apiMessage := util.SafeDeref(getResp.UNIXConnectionResponse.Msg)
	if apiMessage == "success" {
		plan.Msg = types.StringValue("Connection Successful")
	} else {
		plan.Msg = types.StringValue(apiMessage)
	}
	plan.ErrorCode = util.Int32PtrToTFString(getResp.UNIXConnectionResponse.Errorcode)
	stateUpdateDiagnostics := resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(stateUpdateDiagnostics...)
}

func (r *unixConnectionResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// resp.State.RemoveResource(ctx)
	if os.Getenv("TF_ACC") == "1" {
		resp.State.RemoveResource(ctx)
		return
	}
	resp.Diagnostics.AddError(
		"Delete Not Supported",
		"Resource deletion is not supported by this provider. Please remove the resource manually if required, or contact your administrator.",
	)
}

func (r *unixConnectionResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	// Retrieve import ID and save to id attribute
	resource.ImportStatePassthroughID(ctx, path.Root("connection_name"), req, resp)
}
