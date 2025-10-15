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
	"net/http"
	"os"
	"terraform-provider-Saviynt/internal/client"
	"terraform-provider-Saviynt/util"
	connectionsutil "terraform-provider-Saviynt/util/connectionsutil"
	"terraform-provider-Saviynt/util/errorsutil"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	openapi "github.com/saviynt/saviynt-api-go-client/connections"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &UnixConnectionResource{}
var _ resource.ResourceWithImportState = &UnixConnectionResource{}

// Initialize error codes for Unix Connection operations
var unixErrorCodes = errorsutil.NewConnectorErrorCodes(errorsutil.ConnectorTypeUnix)

type UnixConnectorResourceModel struct {
	BaseConnectorResourceModel
	ID                            types.String `tfsdk:"id"`
	HostName                      types.String `tfsdk:"host_name"`
	PortNumber                    types.String `tfsdk:"port_number"`
	Username                      types.String `tfsdk:"username"`
	Password                      types.String `tfsdk:"password"`
	PasswordWo                    types.String `tfsdk:"password_wo"`
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
	PassphraseWo                  types.String `tfsdk:"passphrase_wo"`
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
	SSHKeyWo                      types.String `tfsdk:"ssh_key_wo"`
	LockAccountCommand            types.String `tfsdk:"lock_account_command"`
	UnlockAccountCommand          types.String `tfsdk:"unlock_account_command"`
	PassThroughConnectionDetails  types.String `tfsdk:"pass_through_connection_details"`
	SSHPassThroughPassword        types.String `tfsdk:"ssh_pass_through_password"`
	SSHPassThroughPasswordWo      types.String `tfsdk:"ssh_pass_through_password_wo"`
	SSHPassThroughSSHKEY          types.String `tfsdk:"ssh_pass_through_sshkey"`
	SSHPassThroughSSHKEYWo        types.String `tfsdk:"ssh_pass_through_sshkey_wo"`
	SSHPassThroughPassphrase      types.String `tfsdk:"ssh_pass_through_passphrase"`
	SSHPassThroughPassphraseWo    types.String `tfsdk:"ssh_pass_through_passphrase_wo"`
}

type UnixConnectionResource struct {
	client            client.SaviyntClientInterface
	token             string
	provider          client.SaviyntProviderInterface
	connectionFactory client.ConnectionFactoryInterface
}

func NewUnixConnectionResource() resource.Resource {
	return &UnixConnectionResource{
		connectionFactory: &client.DefaultConnectionFactory{},
	}
}

func NewUnixConnectionResourceWithFactory(factory client.ConnectionFactoryInterface) resource.Resource {
	return &UnixConnectionResource{
		connectionFactory: factory,
	}
}

func (r *UnixConnectionResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
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
			Description: "Property for USERNAME",
		},
		"password": schema.StringAttribute{
			Optional:    true,
			Sensitive:   true,
			Description: "Property for PASSWORD. Either this or password_wo need to be set to configure the password attribute.",
			Validators: []validator.String{
				stringvalidator.ConflictsWith(path.MatchRoot("password_wo")),
			},
		},
		"password_wo": schema.StringAttribute{
			Optional:    true,
			WriteOnly:   true,
			Description: "Password write-only attribute. Either this or password need to be set to configure the password attribute.",
			Validators: []validator.String{
				stringvalidator.ConflictsWith(path.MatchRoot("password")),
			},
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
			Computed:    true,
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
			Sensitive:   true,
			Description: "Property for PASSPHRASE. Either this or passphrase_wo need to be set to configure the passphrase attribute.",
			Validators: []validator.String{
				stringvalidator.ConflictsWith(path.MatchRoot("passphrase_wo")),
			},
		},
		"passphrase_wo": schema.StringAttribute{
			Optional:    true,
			WriteOnly:   true,
			Description: "Passphrase write-only attribute. Either this or passphrase need to be set to configure the passphrase attribute.",
			Validators: []validator.String{
				stringvalidator.ConflictsWith(path.MatchRoot("passphrase")),
			},
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
			Sensitive:   true,
			Description: "Property for SSH_KEY. Either this or ssh_key_wo need to be set to configure the ssh_key attribute.",
			Validators: []validator.String{
				stringvalidator.ConflictsWith(path.MatchRoot("ssh_key_wo")),
			},
		},
		"ssh_key_wo": schema.StringAttribute{
			Optional:    true,
			WriteOnly:   true,
			Description: "SSH key write-only attribute. Either this or ssh_key need to be set to configure the ssh_key attribute.",
			Validators: []validator.String{
				stringvalidator.ConflictsWith(path.MatchRoot("ssh_key")),
			},
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
			Sensitive:   true,
			Description: "Property for SSHPassThroughPassword. Either this or ssh_pass_through_password_wo need to be set to configure the ssh_pass_through_password attribute.",
			Validators: []validator.String{
				stringvalidator.ConflictsWith(path.MatchRoot("ssh_pass_through_password_wo")),
			},
		},
		"ssh_pass_through_password_wo": schema.StringAttribute{
			Optional:    true,
			WriteOnly:   true,
			Description: "SSH pass-through password write-only attribute. Either this or ssh_pass_through_password need to be set to configure the ssh_pass_through_password attribute.",
			Validators: []validator.String{
				stringvalidator.ConflictsWith(path.MatchRoot("ssh_pass_through_password")),
			},
		},
		"ssh_pass_through_sshkey": schema.StringAttribute{
			Optional:    true,
			Sensitive:   true,
			Description: "Property for SSHPassThroughSSHKEY. Either this or ssh_pass_through_sshkey_wo need to be set to configure the ssh_pass_through_sshkey attribute.",
			Validators: []validator.String{
				stringvalidator.ConflictsWith(path.MatchRoot("ssh_pass_through_sshkey_wo")),
			},
		},
		"ssh_pass_through_sshkey_wo": schema.StringAttribute{
			Optional:    true,
			WriteOnly:   true,
			Description: "Property for SSHPassThroughSSHKEY. Either this or ssh_pass_through_sshkey need to be set to configure the ssh_pass_through_sshkey attribute.",
			Validators: []validator.String{
				stringvalidator.ConflictsWith(path.MatchRoot("ssh_pass_through_sshkey")),
			},
		},
		"ssh_pass_through_passphrase": schema.StringAttribute{
			Optional:    true,
			Sensitive:   true,
			Description: "Property for SSHPassThroughPassphrase. Either this or ssh_pass_through_passphrase_wo need to be set to configure the ssh_pass_through_passphrase attribute.",
			Validators: []validator.String{
				stringvalidator.ConflictsWith(path.MatchRoot("ssh_pass_through_passphrase_wo")),
			},
		},
		"ssh_pass_through_passphrase_wo": schema.StringAttribute{
			Optional:    true,
			WriteOnly:   true,
			Description: "SSH pass-through passphrase write-only attribute. Either this or ssh_pass_through_passphrase need to be set to configure the ssh_pass_through_passphrase attribute.",
			Validators: []validator.String{
				stringvalidator.ConflictsWith(path.MatchRoot("ssh_pass_through_passphrase")),
			},
		},
	}
}

func (r *UnixConnectionResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: util.UnixConnDescription,
		Attributes:  connectionsutil.MergeResourceAttributes(BaseConnectorResourceSchema(), UnixConnectorResourceSchema()),
	}
}

func (r *UnixConnectionResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	opCtx := errorsutil.CreateOperationContext(errorsutil.ConnectorTypeUnix, "configure", "")
	ctx = opCtx.AddContextToLogger(ctx)

	opCtx.LogOperationStart(ctx, "Starting Unix connection resource configuration")

	// Check if provider data is available.
	if req.ProviderData == nil {
		tflog.Debug(ctx, "ProviderData is nil, returning early")
		opCtx.LogOperationEnd(ctx, "Unix connection resource configuration completed - no provider data")
		return
	}

	// Cast provider data to your provider type.
	prov, ok := req.ProviderData.(*SaviyntProvider)
	if !ok {
		errorCode := unixErrorCodes.ProviderConfig()
		opCtx.LogOperationError(ctx, "Provider configuration failed", errorCode,
			fmt.Errorf("expected *saviyntProvider, got different type"),
			map[string]interface{}{"expected_type": "*saviyntProvider"})

		resp.Diagnostics.AddError(
			errorsutil.GetErrorMessage(errorsutil.ErrProviderConfig),
			fmt.Sprintf("[%s] Expected *saviyntProvider, got different type", errorCode),
		)
		return
	}

	// Set the client and token from the provider state using interface wrapper.
	r.client = &client.SaviyntClientWrapper{Client: prov.client}
	r.token = prov.accessToken
	r.provider = &client.SaviyntProviderWrapper{Provider: prov} // Store provider reference for retry logic

	opCtx.LogOperationEnd(ctx, "Unix connection resource configured successfully")
}

// SetClient sets the client for testing purposes
func (r *UnixConnectionResource) SetClient(client client.SaviyntClientInterface) {
	r.client = client
}

// SetToken sets the token for testing purposes
func (r *UnixConnectionResource) SetToken(token string) {
	r.token = token
}

// SetProvider sets the provider for testing purposes
func (r *UnixConnectionResource) SetProvider(provider client.SaviyntProviderInterface) {
	r.provider = provider
}

func (r *UnixConnectionResource) BuildUnixConnector(plan *UnixConnectorResourceModel, config *UnixConnectorResourceModel) openapi.UNIXConnector {
	var password string
	if !config.Password.IsNull() && !config.Password.IsUnknown() {
		password = config.Password.ValueString()
	} else if !config.PasswordWo.IsNull() && !config.PasswordWo.IsUnknown() {
		password = config.PasswordWo.ValueString()
	}

	var passphrase string
	if !config.Passphrase.IsNull() && !config.Passphrase.IsUnknown() {
		passphrase = config.Passphrase.ValueString()
	} else if !config.PassphraseWo.IsNull() && !config.PassphraseWo.IsUnknown() {
		passphrase = config.PassphraseWo.ValueString()
	}

	var sshKey string
	if !config.SSHKey.IsNull() && !config.SSHKey.IsUnknown() {
		sshKey = config.SSHKey.ValueString()
	} else if !config.SSHKeyWo.IsNull() && !config.SSHKeyWo.IsUnknown() {
		sshKey = config.SSHKeyWo.ValueString()
	}

	var sshPassthroughPassword string
	if !config.SSHPassThroughPassword.IsNull() && !config.SSHPassThroughPassword.IsUnknown() {
		sshPassthroughPassword = config.SSHPassThroughPassword.ValueString()
	} else if !config.SSHPassThroughPasswordWo.IsNull() && !config.SSHPassThroughPasswordWo.IsUnknown() {
		sshPassthroughPassword = config.SSHPassThroughPasswordWo.ValueString()
	}

	var SSHPassthroughPassphrase string
	if !config.SSHPassThroughPassphrase.IsNull() && !config.SSHPassThroughPassphrase.IsUnknown() {
		SSHPassthroughPassphrase = config.SSHPassThroughPassphrase.ValueString()
	} else if !config.SSHPassThroughPassphraseWo.IsNull() && !config.SSHPassThroughPassphraseWo.IsUnknown() {
		SSHPassthroughPassphrase = config.SSHPassThroughPassphraseWo.ValueString()
	}

	var sshPassthroughSSHKey string
	if !config.SSHPassThroughSSHKEY.IsNull() && !config.SSHPassThroughSSHKEY.IsUnknown() {
		sshPassthroughSSHKey = config.SSHPassThroughSSHKEY.ValueString()
	} else if !config.SSHPassThroughSSHKEYWo.IsNull() && !config.SSHPassThroughSSHKEYWo.IsUnknown() {
		sshPassthroughSSHKey = config.SSHPassThroughSSHKEYWo.ValueString()
	}

	unixConn := openapi.UNIXConnector{
		BaseConnector: openapi.BaseConnector{
			//required field
			Connectiontype: "Unix",
			ConnectionName: plan.ConnectionName.ValueString(),
			//optional field
			Description:     util.StringPointerOrEmpty(plan.Description),
			Defaultsavroles: util.StringPointerOrEmpty(plan.DefaultSavRoles),
			EmailTemplate:   util.StringPointerOrEmpty(plan.EmailTemplate),
		},
		//required field
		HOST_NAME:   plan.HostName.ValueString(),
		PORT_NUMBER: plan.PortNumber.ValueString(),
		USERNAME:    plan.Username.ValueString(),
		//optional field
		PASSWORD:                            util.StringPointerOrEmpty(types.StringValue(password)),
		GROUPS_FILE:                         util.StringPointerOrEmpty(plan.GroupsFile),
		ACCOUNTS_FILE:                       util.StringPointerOrEmpty(plan.AccountsFile),
		SHADOW_FILE:                         util.StringPointerOrEmpty(plan.ShadowFile),
		PROVISION_ACCOUNT_COMMAND:           util.StringPointerOrEmpty(plan.ProvisionAccountCommand),
		DEPROVISION_ACCOUNT_COMMAND:         util.StringPointerOrEmpty(plan.DeprovisionAccountCommand),
		ADD_ACCESS_COMMAND:                  util.StringPointerOrEmpty(plan.AddAccessCommand),
		REMOVE_ACCESS_COMMAND:               util.StringPointerOrEmpty(plan.RemoveAccessCommand),
		CHANGE_PASSWRD_JSON:                 util.StringPointerOrEmpty(plan.ChangePasswordJSON),
		PEM_KEY_FILE:                        util.StringPointerOrEmpty(plan.PemKeyFile),
		ENABLE_ACCOUNT_COMMAND:              util.StringPointerOrEmpty(plan.EnableAccountCommand),
		DISABLE_ACCOUNT_COMMAND:             util.StringPointerOrEmpty(plan.DisableAccountCommand),
		ACCOUNT_ENTITLEMENT_MAPPING_COMMAND: util.StringPointerOrEmpty(plan.AccountEntitlementMappingCmd),
		PASSPHRASE:                          util.StringPointerOrEmpty(types.StringValue(passphrase)),
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
		SSH_KEY:                             util.StringPointerOrEmpty(types.StringValue(sshKey)),
		LOCK_ACCOUNT_COMMAND:                util.StringPointerOrEmpty(plan.LockAccountCommand),
		UNLOCK_ACCOUNT_COMMAND:              util.StringPointerOrEmpty(plan.UnlockAccountCommand),
		PassThroughConnectionDetails:        util.StringPointerOrEmpty(plan.PassThroughConnectionDetails),
		SSHPassThroughPassword:              util.StringPointerOrEmpty(types.StringValue(sshPassthroughPassword)),
		SSHPassThroughSSHKEY:                util.StringPointerOrEmpty(types.StringValue(sshPassthroughSSHKey)),
		SSHPassThroughPassphrase:            util.StringPointerOrEmpty(types.StringValue(SSHPassthroughPassphrase)),
	}

	// Handle vault configuration
	if !plan.VaultConnection.IsNull() && plan.VaultConnection.ValueString() != "" {
		unixConn.BaseConnector.VaultConnection = util.SafeStringConnector(plan.VaultConnection.ValueString())
		unixConn.BaseConnector.VaultConfiguration = util.SafeStringConnector(plan.VaultConfiguration.ValueString())
		unixConn.BaseConnector.Saveinvault = util.SafeStringConnector(plan.SaveInVault.ValueString())
	}

	return unixConn
}

func (r *UnixConnectionResource) UpdateModelFromCreateResponse(plan *UnixConnectorResourceModel, apiResp *openapi.CreateOrUpdateResponse) {
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
}

func (r *UnixConnectionResource) CreateUnixConnection(ctx context.Context, plan *UnixConnectorResourceModel, config *UnixConnectorResourceModel) (*openapi.CreateOrUpdateResponse, error) {
	connectionName := plan.ConnectionName.ValueString()
	opCtx := errorsutil.CreateOperationContext(errorsutil.ConnectorTypeUnix, "create", connectionName)

	// Create logging context (separate from API context)
	logCtx := opCtx.AddContextToLogger(ctx)

	opCtx.LogOperationStart(logCtx, "Starting Unix connection creation")

	// Check if connection already exists (idempotency check) with retry logic
	tflog.Debug(logCtx, "Checking if connection already exists")
	var existingResource *openapi.GetConnectionDetailsResponse
	var finalHttpResp *http.Response

	err := r.provider.AuthenticatedAPICallWithRetry(ctx, "get_connection_details_idempotency", func(token string) error {
		connectionOps := r.connectionFactory.CreateConnectionOperations(r.client.APIBaseURL(), token)
		resp, httpResp, err := connectionOps.GetConnectionDetails(ctx, connectionName)
		if httpResp != nil && httpResp.StatusCode == 401 {
			return fmt.Errorf("401 unauthorized")
		}
		existingResource = resp
		finalHttpResp = httpResp // Update on every call including retries
		return err
	})

	if err != nil && finalHttpResp != nil && finalHttpResp.StatusCode != 412 {
		errorCode := unixErrorCodes.ReadFailed()
		opCtx.LogOperationError(logCtx, "Failed to check existing connection", errorCode, err)
		return nil, errorsutil.CreateStandardError(errorsutil.ConnectorTypeUnix, errorCode, "create", connectionName, err)
	}

	if existingResource != nil &&
		existingResource.UNIXConnectionResponse != nil &&
		existingResource.UNIXConnectionResponse.Errorcode != nil &&
		*existingResource.UNIXConnectionResponse.Errorcode == 0 {

		errorCode := unixErrorCodes.DuplicateName()
		opCtx.LogOperationError(ctx, "Connection name already exists.Please import or use a different name", errorCode,
			fmt.Errorf("duplicate connection name"))
		return nil, errorsutil.CreateStandardError(errorsutil.ConnectorTypeUnix, errorCode, "create", connectionName, nil)
	}

	// Build Unix connection create request
	tflog.Debug(ctx, "Building Unix connection create request")

	unixConn := r.BuildUnixConnector(plan, config)
	createReq := openapi.CreateOrUpdateRequest{
		UNIXConnector: &unixConn,
	}

	// Execute create operation through interface
	tflog.Debug(ctx, "Executing create operation")

	// Execute create operation with retry logic
	var apiResp *openapi.CreateOrUpdateResponse

	err = r.provider.AuthenticatedAPICallWithRetry(ctx, "create_unix_connection", func(token string) error {
		connectionOps := r.connectionFactory.CreateConnectionOperations(r.client.APIBaseURL(), token)
		resp, httpResp, err := connectionOps.CreateOrUpdateConnection(ctx, createReq)
		if httpResp != nil && httpResp.StatusCode == 401 {
			return fmt.Errorf("401 unauthorized")
		}
		apiResp = resp
		return err
	})
	if err != nil {
		errorCode := unixErrorCodes.CreateFailed()
		opCtx.LogOperationError(ctx, "Failed to create Unix connection", errorCode, err)
		return nil, errorsutil.CreateStandardError(errorsutil.ConnectorTypeUnix, errorCode, "create", connectionName, err)
	}

	if apiResp != nil && apiResp.ErrorCode != nil && *apiResp.ErrorCode != "0" {
		apiErr := fmt.Errorf("API returned error code %s: %s", *apiResp.ErrorCode, errorsutil.SanitizeMessage(apiResp.Msg))
		errorCode := unixErrorCodes.APIError()
		opCtx.LogOperationError(ctx, "Unix connection creation failed with API error", errorCode, apiErr,
			map[string]interface{}{
				"api_error_code": *apiResp.ErrorCode,
				"message":        errorsutil.SanitizeMessage(apiResp.Msg),
			})
		return nil, errorsutil.CreateStandardError(errorsutil.ConnectorTypeUnix, errorCode, "create", connectionName, apiErr)
	}

	opCtx.LogOperationEnd(logCtx, "Unix connection created successfully",
		map[string]interface{}{"connection_key": func() interface{} {
			if apiResp.ConnectionKey != nil {
				return *apiResp.ConnectionKey
			}
			return "unknown"
		}()})

	return apiResp, nil
}

func (r *UnixConnectionResource) ReadUnixConnection(ctx context.Context, connectionName string) (*openapi.GetConnectionDetailsResponse, error) {
	opCtx := errorsutil.CreateOperationContext(errorsutil.ConnectorTypeUnix, "read", connectionName)

	// Create logging context (separate from API context)
	logCtx := opCtx.AddContextToLogger(ctx)

	opCtx.LogOperationStart(logCtx, "Starting Unix connection read operation")

	// Execute read operation with retry logic
	var apiResp *openapi.GetConnectionDetailsResponse

	err := r.provider.AuthenticatedAPICallWithRetry(ctx, "read_unix_connection", func(token string) error {
		connectionOps := r.connectionFactory.CreateConnectionOperations(r.client.APIBaseURL(), token)
		resp, httpResp, err := connectionOps.GetConnectionDetails(ctx, connectionName)
		if httpResp != nil && httpResp.StatusCode == 401 {
			return fmt.Errorf("401 unauthorized")
		}
		apiResp = resp
		return err
	})

	if err != nil {
		errorCode := unixErrorCodes.ReadFailed()
		opCtx.LogOperationError(logCtx, "Failed to read Unix connection", errorCode, err)
		return nil, errorsutil.CreateStandardError(errorsutil.ConnectorTypeUnix, errorCode, "read", connectionName, err)
	}

	if err := r.ValidateUnixConnectionResponse(apiResp); err != nil {
		errorCode := unixErrorCodes.APIError()
		opCtx.LogOperationError(ctx, "Invalid connection type for Unix datasource", errorCode, err)
		return nil, fmt.Errorf("[%s] Unable to verify connection type for connection %q. The provider could not determine the type of this connection. Please ensure the connection name is correct and belongs to a supported connector type", errorCode, connectionName)
	}

	if apiResp != nil && apiResp.UNIXConnectionResponse != nil && apiResp.UNIXConnectionResponse.Errorcode != nil && *apiResp.UNIXConnectionResponse.Errorcode != 0 {
		apiErr := fmt.Errorf("API returned error code %d: %s", *apiResp.UNIXConnectionResponse.Errorcode, errorsutil.SanitizeMessage(apiResp.UNIXConnectionResponse.Msg))
		errorCode := unixErrorCodes.APIError()
		opCtx.LogOperationError(ctx, "Unix connection read failed with API error", errorCode, apiErr,
			map[string]interface{}{
				"api_error_code": *apiResp.UNIXConnectionResponse.Errorcode,
				"message":        errorsutil.SanitizeMessage(apiResp.UNIXConnectionResponse.Msg),
			})
		return nil, errorsutil.CreateStandardError(errorsutil.ConnectorTypeUnix, errorCode, "read", connectionName, apiErr)
	}

	opCtx.LogOperationEnd(logCtx, "Unix connection read completed successfully",
		map[string]interface{}{"connection_key": func() interface{} {
			if apiResp.UNIXConnectionResponse != nil && apiResp.UNIXConnectionResponse.Connectionkey != nil {
				return *apiResp.UNIXConnectionResponse.Connectionkey
			}
			return "unknown"
		}()})

	return apiResp, nil
}

func (r *UnixConnectionResource) UpdateModelFromReadResponse(state *UnixConnectorResourceModel, apiResp *openapi.GetConnectionDetailsResponse) {
	state.ConnectionKey = types.Int64Value(int64(*apiResp.UNIXConnectionResponse.Connectionkey))
	state.ID = types.StringValue(fmt.Sprintf("%d", *apiResp.UNIXConnectionResponse.Connectionkey))
	state.ConnectionName = util.SafeStringDatasource(apiResp.UNIXConnectionResponse.Connectionname)
	state.Description = util.SafeStringDatasource(apiResp.UNIXConnectionResponse.Description)
	state.DefaultSavRoles = util.SafeStringDatasource(apiResp.UNIXConnectionResponse.Defaultsavroles)
	state.EmailTemplate = util.SafeStringDatasource(apiResp.UNIXConnectionResponse.Emailtemplate)

	// Map Unix-specific connection attributes
	if apiResp.UNIXConnectionResponse.Connectionattributes != nil {
		state.HostName = util.SafeStringDatasource(apiResp.UNIXConnectionResponse.Connectionattributes.HOST_NAME)
		state.PortNumber = util.SafeStringDatasource(apiResp.UNIXConnectionResponse.Connectionattributes.PORT_NUMBER)
		state.Username = util.SafeStringDatasource(apiResp.UNIXConnectionResponse.Connectionattributes.USERNAME)
		state.GroupsFile = util.SafeStringDatasource(apiResp.UNIXConnectionResponse.Connectionattributes.GROUPS_FILE)
		state.AccountsFile = util.SafeStringDatasource(apiResp.UNIXConnectionResponse.Connectionattributes.ACCOUNTS_FILE)
		state.ShadowFile = util.SafeStringDatasource(apiResp.UNIXConnectionResponse.Connectionattributes.SHADOW_FILE)
		state.ProvisionAccountCommand = util.SafeStringDatasource(apiResp.UNIXConnectionResponse.Connectionattributes.PROVISION_ACCOUNT_COMMAND)
		state.DeprovisionAccountCommand = util.SafeStringDatasource(apiResp.UNIXConnectionResponse.Connectionattributes.DEPROVISION_ACCOUNT_COMMAND)
		state.AddAccessCommand = util.SafeStringDatasource(apiResp.UNIXConnectionResponse.Connectionattributes.ADD_ACCESS_COMMAND)
		state.RemoveAccessCommand = util.SafeStringDatasource(apiResp.UNIXConnectionResponse.Connectionattributes.REMOVE_ACCESS_COMMAND)
		state.ChangePasswordJSON = util.SafeStringDatasource(apiResp.UNIXConnectionResponse.Connectionattributes.CHANGE_PASSWRD_JSON)
		state.PemKeyFile = util.SafeStringDatasource(apiResp.UNIXConnectionResponse.Connectionattributes.PEM_KEY_FILE)
		state.EnableAccountCommand = util.SafeStringDatasource(apiResp.UNIXConnectionResponse.Connectionattributes.ENABLE_ACCOUNT_COMMAND)
		state.DisableAccountCommand = util.SafeStringDatasource(apiResp.UNIXConnectionResponse.Connectionattributes.DISABLE_ACCOUNT_COMMAND)
		state.AccountEntitlementMappingCmd = util.SafeStringDatasource(apiResp.UNIXConnectionResponse.Connectionattributes.ACCOUNT_ENTITLEMENT_MAPPING_COMMAND)
		state.UpdateAccountCommand = util.SafeStringDatasource(apiResp.UNIXConnectionResponse.Connectionattributes.UPDATE_ACCOUNT_COMMAND)
		state.CreateGroupCommand = util.SafeStringDatasource(apiResp.UNIXConnectionResponse.Connectionattributes.CREATE_GROUP_COMMAND)
		state.DeleteGroupCommand = util.SafeStringDatasource(apiResp.UNIXConnectionResponse.Connectionattributes.DELETE_GROUP_COMMAND)
		state.AddGroupOwnerCommand = util.SafeStringDatasource(apiResp.UNIXConnectionResponse.Connectionattributes.ADD_GROUP_OWNER_COMMAND)
		state.AddPrimaryGroupCommand = util.SafeStringDatasource(apiResp.UNIXConnectionResponse.Connectionattributes.ADD_PRIMARY_GROUP_COMMAND)
		state.FirefighterGrantAccessCommand = util.SafeStringDatasource(apiResp.UNIXConnectionResponse.Connectionattributes.FIREFIGHTERID_GRANT_ACCESS_COMMAND)
		state.FirefighterRevokeAccessCmd = util.SafeStringDatasource(apiResp.UNIXConnectionResponse.Connectionattributes.FIREFIGHTERID_REVOKE_ACCESS_COMMAND)
		state.InactiveLockAccount = util.SafeStringDatasource(apiResp.UNIXConnectionResponse.Connectionattributes.INACTIVE_LOCK_ACCOUNT)
		state.StatusThresholdConfig = util.SafeStringDatasource(apiResp.UNIXConnectionResponse.Connectionattributes.STATUS_THRESHOLD_CONFIG)
		state.CustomConfigJSON = util.SafeStringDatasource(apiResp.UNIXConnectionResponse.Connectionattributes.CUSTOM_CONFIG_JSON)
		state.LockAccountCommand = util.SafeStringDatasource(apiResp.UNIXConnectionResponse.Connectionattributes.LOCK_ACCOUNT_COMMAND)
		state.UnlockAccountCommand = util.SafeStringDatasource(apiResp.UNIXConnectionResponse.Connectionattributes.UNLOCK_ACCOUNT_COMMAND)
		state.PassThroughConnectionDetails = util.SafeStringDatasource(apiResp.UNIXConnectionResponse.Connectionattributes.PassThroughConnectionDetails)
	}
}

func (r *UnixConnectionResource) ValidateUnixConnectionResponse(apiResp *openapi.GetConnectionDetailsResponse) error {
	if apiResp != nil && apiResp.UNIXConnectionResponse == nil {
		return fmt.Errorf("verify the connection type - Unix connection response is nil")
	}
	return nil
}

func (r *UnixConnectionResource) UpdateUnixConnection(ctx context.Context, plan *UnixConnectorResourceModel, config *UnixConnectorResourceModel) (*openapi.CreateOrUpdateResponse, error) {
	connectionName := plan.ConnectionName.ValueString()
	opCtx := errorsutil.CreateOperationContext(errorsutil.ConnectorTypeUnix, "update", connectionName)

	// Create logging context (separate from API context)
	logCtx := opCtx.AddContextToLogger(ctx)

	opCtx.LogOperationStart(logCtx, "Starting Unix connection update")

	// Build Unix connection update request
	tflog.Debug(logCtx, "Building Unix connection update request")

	unixConn := r.BuildUnixConnector(plan, config)

	updateReq := openapi.CreateOrUpdateRequest{
		UNIXConnector: &unixConn,
	}

	// Execute update operation with retry logic
	tflog.Debug(logCtx, "Executing update operation")
	var apiResp *openapi.CreateOrUpdateResponse

	err := r.provider.AuthenticatedAPICallWithRetry(ctx, "update_unix_connection", func(token string) error {
		connectionOps := r.connectionFactory.CreateConnectionOperations(r.client.APIBaseURL(), token)
		resp, httpResp, err := connectionOps.CreateOrUpdateConnection(ctx, updateReq)
		if httpResp != nil && httpResp.StatusCode == 401 {
			return fmt.Errorf("401 unauthorized")
		}
		apiResp = resp
		return err
	})

	if err != nil {
		errorCode := unixErrorCodes.UpdateFailed()
		opCtx.LogOperationError(logCtx, "Failed to update Unix connection", errorCode, err)
		return nil, errorsutil.CreateStandardError(errorsutil.ConnectorTypeUnix, errorCode, "update", connectionName, err)
	}

	if apiResp != nil && apiResp.ErrorCode != nil && *apiResp.ErrorCode != "0" {
		apiErr := fmt.Errorf("API returned error code %s: %s", *apiResp.ErrorCode, errorsutil.SanitizeMessage(apiResp.Msg))
		errorCode := unixErrorCodes.APIError()
		opCtx.LogOperationError(logCtx, "Unix connection update failed with API error", errorCode, apiErr,
			map[string]interface{}{
				"api_error_code": *apiResp.ErrorCode,
				"message":        errorsutil.SanitizeMessage(apiResp.Msg),
			})
		return nil, errorsutil.CreateStandardError(errorsutil.ConnectorTypeUnix, errorCode, "update", connectionName, apiErr)
	}

	opCtx.LogOperationEnd(logCtx, "Unix connection updated successfully",
		map[string]interface{}{"connection_key": func() interface{} {
			if apiResp.ConnectionKey != nil {
				return *apiResp.ConnectionKey
			}
			return "unknown"
		}()})

	return apiResp, nil
}

func (r *UnixConnectionResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan, config UnixConnectorResourceModel

	opCtx := errorsutil.CreateOperationContext(errorsutil.ConnectorTypeUnix, "terraform_create", "")
	ctx = opCtx.AddContextToLogger(ctx)

	opCtx.LogOperationStart(ctx, "Starting Unix connection resource creation")

	// Extract plan from request
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		errorCode := unixErrorCodes.PlanExtraction()
		opCtx.LogOperationError(ctx, "Failed to get plan from request", errorCode,
			fmt.Errorf("plan extraction failed"))
		resp.Diagnostics.AddError(
			errorsutil.GetErrorMessage(errorsutil.ErrPlanExtraction),
			fmt.Sprintf("[%s] Unable to extract Terraform plan from request", errorCode),
		)
		return
	}

	connectionName := plan.ConnectionName.ValueString()
	// Update operation context with connection name
	opCtx.ConnectionName = connectionName
	ctx = opCtx.AddContextToLogger(ctx)

	//Extract config from request
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		errorCode := unixErrorCodes.ConfigExtraction()
		opCtx.LogOperationError(ctx, "Failed to get config from request", errorCode,
			fmt.Errorf("config extraction failed"))
		resp.Diagnostics.AddError(
			errorsutil.GetErrorMessage(errorsutil.ErrConfigExtraction),
			fmt.Sprintf("[%s] Unable to extract Terraform configuration from request for connection '%s'", errorCode, connectionName),
		)
		return
	}

	// Use interface pattern instead of direct API client creation
	apiResp, err := r.CreateUnixConnection(ctx, &plan, &config)
	if err != nil {
		opCtx.LogOperationError(ctx, "Unix connection creation failed", "", err)
		resp.Diagnostics.AddError(
			"Unix Connection Creation Failed",
			err.Error(),
		)
		return
	}

	// Update model from create response
	r.UpdateModelFromCreateResponse(&plan, apiResp)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)

	opCtx.LogOperationEnd(ctx, "Unix connection resource created successfully",
		map[string]interface{}{"connection_key": plan.ConnectionKey.ValueInt64()})
}

func (r *UnixConnectionResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state UnixConnectorResourceModel

	opCtx := errorsutil.CreateOperationContext(errorsutil.ConnectorTypeUnix, "terraform_read", "")
	ctx = opCtx.AddContextToLogger(ctx)

	opCtx.LogOperationStart(ctx, "Starting Unix connection resource read")

	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		errorCode := unixErrorCodes.StateExtraction()
		opCtx.LogOperationError(ctx, "Failed to get state from request", errorCode,
			fmt.Errorf("state extraction failed"))
		resp.Diagnostics.AddError(
			errorsutil.GetErrorMessage(errorsutil.ErrStateExtraction),
			fmt.Sprintf("[%s] Unable to extract Terraform state from request", errorCode),
		)
		return
	}

	connectionName := state.ConnectionName.ValueString()
	// Update operation context with connection name
	opCtx.ConnectionName = connectionName
	ctx = opCtx.AddContextToLogger(ctx)

	// Use interface pattern instead of direct API client creation
	apiResp, err := r.ReadUnixConnection(ctx, connectionName)
	if err != nil {
		opCtx.LogOperationError(ctx, "Unix connection read failed", "", err)
		resp.Diagnostics.AddError(
			"Unix Connection Read Failed",
			err.Error(),
		)
		return
	}

	// Update model from read response
	r.UpdateModelFromReadResponse(&state, apiResp)

	apiMessage := util.SafeDeref(apiResp.UNIXConnectionResponse.Msg)
	if apiMessage == "success" {
		state.Msg = types.StringValue("Connection Read Successful")
	} else {
		state.Msg = types.StringValue(apiMessage)
	}
	state.ErrorCode = util.Int32PtrToTFString(apiResp.UNIXConnectionResponse.Errorcode)

	stateDiagnostics := resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(stateDiagnostics...)
	if resp.Diagnostics.HasError() {
		errorCode := unixErrorCodes.StateUpdate()
		opCtx.LogOperationError(ctx, "Failed to set state", errorCode,
			fmt.Errorf("state update failed"))
		resp.Diagnostics.AddError(
			errorsutil.GetErrorMessage(errorsutil.ErrStateUpdate),
			fmt.Sprintf("[%s] Unable to update Terraform state for connection '%s'", errorCode, connectionName),
		)
		return
	}

	opCtx.LogOperationEnd(ctx, "Unix connection resource read completed successfully")
}

func (r *UnixConnectionResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state, config UnixConnectorResourceModel

	opCtx := errorsutil.CreateOperationContext(errorsutil.ConnectorTypeUnix, "terraform_update", "")
	ctx = opCtx.AddContextToLogger(ctx)

	opCtx.LogOperationStart(ctx, "Starting Unix connection resource update")

	// Extract state from request
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		errorCode := unixErrorCodes.StateExtraction()
		opCtx.LogOperationError(ctx, "Failed to get state from request", errorCode,
			fmt.Errorf("state extraction failed"))
		resp.Diagnostics.AddError(
			errorsutil.GetErrorMessage(errorsutil.ErrStateExtraction),
			fmt.Sprintf("[%s] Unable to extract Terraform state from request", errorCode),
		)
		return
	}

	// Extract plan from request
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		errorCode := unixErrorCodes.PlanExtraction()
		opCtx.LogOperationError(ctx, "Failed to get plan from request", errorCode,
			fmt.Errorf("plan extraction failed"))
		resp.Diagnostics.AddError(
			errorsutil.GetErrorMessage(errorsutil.ErrPlanExtraction),
			fmt.Sprintf("[%s] Unable to extract Terraform plan from request for connection '%s'", errorCode, state.ConnectionName.ValueString()),
		)
		return
	}

	//Extract config from request
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		errorCode := unixErrorCodes.ConfigExtraction()
		opCtx.LogOperationError(ctx, "Failed to get config from request", errorCode,
			fmt.Errorf("config extraction failed"))
		resp.Diagnostics.AddError(
			errorsutil.GetErrorMessage(errorsutil.ErrConfigExtraction),
			fmt.Sprintf("[%s] Unable to extract Terraform configuration from request for connection '%s'", errorCode, plan.ConnectionName.ValueString()),
		)
		return
	}

	// Validate that connection name cannot be updated
	if plan.ConnectionName.ValueString() != state.ConnectionName.ValueString() {
		errorCode := unixErrorCodes.NameImmutable()
		opCtx.LogOperationError(ctx, "Connection name cannot be updated", errorCode,
			fmt.Errorf("attempted to change connection name from '%s' to '%s'", state.ConnectionName.ValueString(), plan.ConnectionName.ValueString()),
			map[string]interface{}{
				"old_name": state.ConnectionName.ValueString(),
				"new_name": plan.ConnectionName.ValueString(),
			})
		resp.Diagnostics.AddError(
			errorsutil.GetErrorMessage(errorCode),
			fmt.Sprintf("[%s] Cannot change connection name from '%s' to '%s'", errorCode, state.ConnectionName.ValueString(), plan.ConnectionName.ValueString()),
		)
		return
	}

	connectionName := plan.ConnectionName.ValueString()
	// Update operation context with connection name
	opCtx.ConnectionName = connectionName
	ctx = opCtx.AddContextToLogger(ctx)

	// Use interface pattern instead of direct API client creation
	updateResp, err := r.UpdateUnixConnection(ctx, &plan, &config)
	if err != nil {
		opCtx.LogOperationError(ctx, "Unix connection update failed", "", err)
		resp.Diagnostics.AddError(
			"Unix Connection Update Failed",
			err.Error(),
		)
		return
	}

	// Read the updated connection to get the latest state
	getResp, err := r.ReadUnixConnection(ctx, connectionName)
	if err != nil {
		opCtx.LogOperationError(ctx, "Failed to read updated Unix connection", "", err)
		resp.Diagnostics.AddError(
			"Unix Connection Post-Update Read Failed",
			err.Error(),
		)
		return
	}

	// Update model from read response
	r.UpdateModelFromReadResponse(&plan, getResp)

	apiMessage := util.SafeDeref(updateResp.Msg)
	plan.Msg = types.StringValue(apiMessage)
	plan.ErrorCode = types.StringValue(*updateResp.ErrorCode)

	stateUpdateDiagnostics := resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(stateUpdateDiagnostics...)
	if resp.Diagnostics.HasError() {
		errorCode := unixErrorCodes.StateUpdate()
		opCtx.LogOperationError(ctx, "Failed to update state after successful update", errorCode,
			fmt.Errorf("state update failed"))
		resp.Diagnostics.AddError(
			errorsutil.GetErrorMessage(errorsutil.ErrStateUpdate),
			fmt.Sprintf("[%s] Unable to update Terraform state after successful update for connection '%s'", errorCode, connectionName),
		)
		return
	}

	opCtx.LogOperationEnd(ctx, "Unix connection resource updated successfully",
		map[string]interface{}{"connection_key": plan.ConnectionKey.ValueInt64()})
}

func (r *UnixConnectionResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
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

func (r *UnixConnectionResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	// Importing an Unix connection resource requires the connection name
	connectionName := req.ID
	opCtx := errorsutil.CreateOperationContext(errorsutil.ConnectorTypeUnix, "terraform_import", connectionName)
	ctx = opCtx.AddContextToLogger(ctx)

	opCtx.LogOperationStart(ctx, "Starting Unix connection resource import")
	// Retrieve import ID and save to id attribute
	resource.ImportStatePassthroughID(ctx, path.Root("connection_name"), req, resp)

	opCtx.LogOperationEnd(ctx, "Unix connection resource import completed successfully",
		map[string]interface{}{"import_id": connectionName})
}
