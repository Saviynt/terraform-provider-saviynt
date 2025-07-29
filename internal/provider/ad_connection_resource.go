// Copyright (c) 2025 Saviynt Inc.
// SPDX-License-Identifier: MPL-2.0

// saviynt_ad_connection_resource manages AD connectors in the Saviynt Security Manager.
// The resource implements the full Terraform lifecycle:
//   - Create: provisions a new AD connector using the supplied configuration.
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

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	s "github.com/saviynt/saviynt-api-go-client"
	openapi "github.com/saviynt/saviynt-api-go-client/connections"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &adConnectionResource{}
var _ resource.ResourceWithImportState = &adConnectionResource{}

type ADConnectorResourceModel struct {
	BaseConnectorResourceModel
	ID                        types.String `tfsdk:"id"`
	URL                       types.String `tfsdk:"url"`
	Username                  types.String `tfsdk:"username"`
	Password                  types.String `tfsdk:"password"`
	LdapOrAd                  types.String `tfsdk:"ldap_or_ad"`
	EntitlementAttribute      types.String `tfsdk:"entitlement_attribute"`
	CheckForUnique            types.String `tfsdk:"check_for_unique"`
	GroupSearchBaseDN         types.String `tfsdk:"group_search_base_dn"`
	CreateUpdateMappings      types.String `tfsdk:"create_update_mappings"`
	IncrementalConfig         types.String `tfsdk:"incremental_config"`
	MaxChangeNumber           types.String `tfsdk:"max_changenumber"`
	ReadOperationalAttributes types.String `tfsdk:"read_operational_attributes"`
	Base                      types.String `tfsdk:"base"`
	DcLocator                 types.String `tfsdk:"dc_locator"`
	StatusThresholdConfig     types.String `tfsdk:"status_threshold_config"`
	RemoveAccountAction       types.String `tfsdk:"remove_account_action"`
	AccountAttribute          types.String `tfsdk:"account_attribute"`
	AccountNameRule           types.String `tfsdk:"account_name_rule"`
	Advsearch                 types.String `tfsdk:"advsearch"`
	Setdefaultpagesize        types.String `tfsdk:"setdefaultpagesize"`
	ResetAndChangePasswrdJson types.String `tfsdk:"reset_and_change_passwrd_json"`
	ReuseInactiveAccount      types.String `tfsdk:"reuse_inactive_account"`
	ImportJson                types.String `tfsdk:"import_json"`
	SupportEmptyString        types.String `tfsdk:"support_empty_string"`
	EnableAccountJson         types.String `tfsdk:"enable_account_json"`
	PageSize                  types.String `tfsdk:"page_size"`
	UserAttribute             types.String `tfsdk:"user_attribute"`
	DefaultUserRole           types.String `tfsdk:"default_user_role"`
	Searchfilter              types.String `tfsdk:"searchfilter"`
	EndpointsFilter           types.String `tfsdk:"endpoints_filter"`
	CreateAccountJson         types.String `tfsdk:"create_account_json"`
	UpdateAccountJson         types.String `tfsdk:"update_account_json"`
	ReuseAccountJson          types.String `tfsdk:"reuse_account_json"`
	EnforceTreeDeletion       types.String `tfsdk:"enforce_tree_deletion"`
	AdvanceFilterJson         types.String `tfsdk:"advance_filter_json"`
	Filter                    types.String `tfsdk:"filter"`
	Objectfilter              types.String `tfsdk:"objectfilter"`
	UpdateUserJson            types.String `tfsdk:"update_user_json"`
	Setrandompassword         types.String `tfsdk:"set_random_password"`
	PasswordMinLength         types.String `tfsdk:"password_min_length"`
	PasswordMaxLength         types.String `tfsdk:"password_max_length"`
	PasswordNoofcapsalpha     types.String `tfsdk:"password_noofcapsalpha"`
	PasswordNoofsplchars      types.String `tfsdk:"password_noofsplchars"`
	PasswordNoofdigits        types.String `tfsdk:"password_noofdigits"`
	GroupImportMapping        types.String `tfsdk:"group_import_mapping"`
	UnlockAccountJson         types.String `tfsdk:"unlock_account_json"`
	StatusKeyJson             types.String `tfsdk:"status_key_json"`
	DisableAccountJson        types.String `tfsdk:"disable_account_json"`
	ModifyUserdataJson        types.String `tfsdk:"modify_user_data_json"`
	OrgBase                   types.String `tfsdk:"org_base"`
	OrganizationAttribute     types.String `tfsdk:"organization_attribute"`
	Createorgjson             types.String `tfsdk:"create_org_json"`
	Updateorgjson             types.String `tfsdk:"update_org_json"`
	ConfigJson                types.String `tfsdk:"config_json"`
	PamConfig                 types.String `tfsdk:"pam_config"`
	EnableGroupManagement     types.String `tfsdk:"enable_group_management"`
	OrgImportJson             types.String `tfsdk:"org_import_json"`
}

type adConnectionResource struct {
	client *s.Client
	token  string
}

func NewADTestConnectionResource() resource.Resource {
	return &adConnectionResource{}
}

func (r *adConnectionResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = "saviynt_ad_connection_resource"
}

func ADConnectorResourceSchema() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:    true,
			Description: "Resource ID.",
		},
		"url": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "LDAP or target system URL. Example: \"ldap://uscentral.com:8972/\"",
		},
		"username": schema.StringAttribute{
			Optional:    true,
			WriteOnly:   true,
			Description: "System admin username.",
		},
		"password": schema.StringAttribute{
			Required:    true,
			WriteOnly:   true,
			Description: "Set the Password.",
		},
		"ldap_or_ad": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "Type of Endpoint - LDAP or AD. Default is 'AD'. Example: \"AD\"",
		},
		"entitlement_attribute": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "Attribute used for entitlements. Example: \"memberOf\"",
		},
		"check_for_unique": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "Uniqueness validation rule JSON. Example: '{\"sAMAccountName\":\"${task.accountName}\"}'",
		},
		"group_search_base_dn": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "Base DN for group search. Example: \"CN=Users,DC=Saviynt,DC=ABC,DC=Com\"",
		},
		"create_update_mappings": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "Mapping for group creation/updation (JSON string). Example: '{\"cn\":\"${role?.customproperty27}\",\"objectCategory\":\"CN=Group,CN=Schema,CN=Configuration,...}'",
		},
		"incremental_config": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "Incremental import configuration.",
		},
		"max_changenumber": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "Maximum change number. Example: \"4\"",
		},
		"read_operational_attributes": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "Flag for reading operational attributes. Example: \"FALSE\"",
		},
		"base": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "LDAP base DN. Example: \"CN=Users,DC=Saviynt,DC=ABC,DC=Com\"",
		},
		"dc_locator": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "Domain controller locator.",
		},
		"status_threshold_config": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "JSON configuration for status thresholds. Example: '{\"statusAndThresholdConfig\":{...}}'",
		},
		"remove_account_action": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "Action on account removal. Example: '{\"removeAction\":\"DELETE\"}'",
		},
		"account_attribute": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "Mapping for LDAP user to EIC account attribute. Example: '[\"ACCOUNTID::objectGUID#Binary\", \"NAME::sAMAccountName#String\", ...]'",
		},
		"account_name_rule": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "Rule to generate account name. Example: \"uid=${task.accountName.toString().toLowerCase()},ou=People,dc=racf,dc=com\"",
		},
		"advsearch": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "Advanced search settings.",
		},
		"setdefaultpagesize": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "Default page size setting. Example: \"FALSE\"",
		},
		"reset_and_change_passwrd_json": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "JSON for reset/change password actions. Example: '{\"RESET\":{\"pwdLastSet\":\"0\",\"title\":\"password reset\"},\"CHANGE\":{\"pwdLastSet\":\"-1\",\"title\":\"password changed\"}}'",
		},
		"reuse_inactive_account": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "Reuse inactive account flag. Example: \"TRUE\"",
		},
		"import_json": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "JSON import configuration. Example: '{\"envproperties\":{\"com.sun.jndi.ldap.connect.timeout\":\"10000\",...}}'",
		},
		"support_empty_string": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "Flag for sending empty values. Example: \"FALSE\"",
		},
		"enable_account_json": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "JSON configuration to enable account actions. Example: '{\"USEDNFROMACCOUNT\":\"NO\", ...}'",
		},
		"page_size": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "LDAP page size. Example: \"1000\"",
		},
		"user_attribute": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "Mapping for LDAP user to EIC user attribute. Example: '[\"USERNAME::sAMAccountName#String\", ...]'",
		},
		"default_user_role": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "Default SAV Role for imported users. Example: \"ROLE_TASK_ADMIN\"",
		},
		"searchfilter": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "LDAP search filter for users. Example: \"OU=Users,DC=domainname,DC=com\"",
		},
		"endpoints_filter": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "Configuration for child endpoints.",
		},
		"create_account_json": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "JSON to create an account. Example: '{\"cn\":\"${cn}\",\"displayname\":\"${user.displayname}\", ...}'",
		},
		"update_account_json": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "JSON to update an account. Example: '{\"uid\":\"${task.accountName.toString().toLowerCase()}\", ...}'",
		},
		"reuse_account_json": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "JSON to reuse an account. Example: '{\"ATTRIBUTESTOCHECK\":{\"userAccountControl\":\"514\",...}}'",
		},
		"enforce_tree_deletion": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "Enforce tree deletion flag. Example: \"TRUE\"",
		},
		"advance_filter_json": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "Advanced filter JSON configuration.",
		},
		"filter": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "Simple filter string.",
		},
		"objectfilter": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "LDAP object filter. Example: \"(objectClass=inetorgperson)\"",
		},
		"update_user_json": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "JSON to update a user. Example: '{\"mail\":\"${user.email}\", ...}'",
		},
		"set_random_password": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "Option to set a random password.",
		},
		"password_min_length": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "Minimum password length. Example: \"8\"",
		},
		"password_max_length": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "Maximum password length. Example: \"12\"",
		},
		"password_noofcapsalpha": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "Number of capital letters required. Example: \"2\"",
		},
		"password_noofsplchars": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "Number of special characters required. Example: \"1\"",
		},
		"password_noofdigits": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "Number of digits required. Example: \"5\"",
		},
		"group_import_mapping": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "JSON mapping for LDAP groups. Example: '{\"entitlementTypeName\":\"memberOf\", ...}'",
		},
		"unlock_account_json": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "JSON to unlock accounts. Example: '{\"lockoutTime\":\"0\"}'",
		},
		"status_key_json": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "JSON for account status keys. Example: '{\"STATUS_ACTIVE\":[\"512\",\"544\"], ...}'",
		},
		"disable_account_json": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "JSON to disable an account. Example: '{\"userAccountControl\":\"546\", ...}'",
		},
		"modify_user_data_json": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "JSON for inline user data transformation.",
		},
		"org_base": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "Organization BASE for provisioning.",
		},
		"organization_attribute": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "Organization attributes.",
		},
		"create_org_json": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "JSON for organization creation.",
		},
		"update_org_json": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "JSON for organization update.",
		},
		"config_json": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "JSON for connection timeout configuration. Example: '{\"connectionTimeoutConfig\":{\"connectionTimeout\":10,\"readTimeout\":50,\"retryWait\":2,\"retryCount\":3}}'",
		},
		"pam_config": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "JSON for PAM bootstrap configuration. Example: '{\"Connection\":\"AD\",...}'",
		},
		"enable_group_management": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "Enable group management. Example: \"TRUE\"",
		},
		"org_import_json": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "JSON for organization import configuration.",
		},
	}
}

func (r *adConnectionResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: util.ADConnDescription,
		Attributes:  connectionsutil.MergeResourceAttributes(BaseConnectorResourceSchema(), ADConnectorResourceSchema()),
	}
}

func (r *adConnectionResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Check if provider data is available.
	if req.ProviderData == nil {
		log.Println("ProviderData is nil, returning early.")
		// resp.Diagnostics.AddError("Provider Data Not Found", "Provider data is not set. Please ensure the provider is configured correctly.")
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

func (r *adConnectionResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan, config ADConnectorResourceModel
	// Extract plan from request
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		log.Printf("[ERROR] Failed to get plan from request: %v", resp.Diagnostics)
		return
	}
	//Extract config from request
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		log.Printf("[ERROR] Failed to get config from request: %v", resp.Diagnostics)
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
	existingResource, _, _ := apiClient.ConnectionsAPI.GetConnectionDetails(ctx).GetConnectionDetailsRequest(reqParams).Execute()
	if existingResource != nil &&
		existingResource.ADConnectionResponse != nil &&
		existingResource.ADConnectionResponse.Errorcode != nil &&
		*existingResource.ADConnectionResponse.Errorcode == 0 {
		log.Printf("[ERROR] Connection name already exists. Please import or use a different name")
		resp.Diagnostics.AddError("API Create Failed", "Connection name already exists. Please import or use a different name")
		return
	}

	adConn := openapi.ADConnector{
		BaseConnector: openapi.BaseConnector{
			//required field
			Connectiontype: "AD",
			ConnectionName: plan.ConnectionName.ValueString(),
			//optional field
			Description:     util.StringPointerOrEmpty(plan.Description),
			Defaultsavroles: util.StringPointerOrEmpty(plan.DefaultSavRoles),
			EmailTemplate:   util.StringPointerOrEmpty(plan.EmailTemplate),
		},
		//required field
		PASSWORD: config.Password.ValueString(),
		//optional field
		URL:                         util.StringPointerOrEmpty(plan.URL),
		USERNAME:                    util.StringPointerOrEmpty(config.Username),
		LDAP_OR_AD:                  util.StringPointerOrEmpty(plan.LdapOrAd),
		ENTITLEMENT_ATTRIBUTE:       util.StringPointerOrEmpty(plan.EntitlementAttribute),
		CHECKFORUNIQUE:              util.StringPointerOrEmpty(plan.CheckForUnique),
		GroupSearchBaseDN:           util.StringPointerOrEmpty(plan.GroupSearchBaseDN),
		CreateUpdateMappings:        util.StringPointerOrEmpty(plan.CreateUpdateMappings),
		INCREMENTAL_CONFIG:          util.StringPointerOrEmpty(plan.IncrementalConfig),
		MAX_CHANGENUMBER:            util.StringPointerOrEmpty(plan.MaxChangeNumber),
		READ_OPERATIONAL_ATTRIBUTES: util.StringPointerOrEmpty(plan.ReadOperationalAttributes),
		BASE:                        util.StringPointerOrEmpty(plan.Base),
		DC_LOCATOR:                  util.StringPointerOrEmpty(plan.DcLocator),
		STATUS_THRESHOLD_CONFIG:     util.StringPointerOrEmpty(plan.StatusThresholdConfig),
		REMOVEACCOUNTACTION:         util.StringPointerOrEmpty(plan.RemoveAccountAction),
		ACCOUNT_ATTRIBUTE:           util.StringPointerOrEmpty(plan.AccountAttribute),
		ACCOUNTNAMERULE:             util.StringPointerOrEmpty(plan.AccountNameRule),
		ADVSEARCH:                   util.StringPointerOrEmpty(plan.Advsearch),
		SETDEFAULTPAGESIZE:          util.StringPointerOrEmpty(plan.Setdefaultpagesize),
		RESETANDCHANGEPASSWRDJSON:   util.StringPointerOrEmpty(plan.ResetAndChangePasswrdJson),
		REUSEINACTIVEACCOUNT:        util.StringPointerOrEmpty(plan.ReuseInactiveAccount),
		IMPORTJSON:                  util.StringPointerOrEmpty(plan.ImportJson),
		SUPPORTEMPTYSTRING:          util.StringPointerOrEmpty(plan.SupportEmptyString),
		ENABLEACCOUNTJSON:           util.StringPointerOrEmpty(plan.EnableAccountJson),
		PAGE_SIZE:                   util.StringPointerOrEmpty(plan.PageSize),
		USER_ATTRIBUTE:              util.StringPointerOrEmpty(plan.UserAttribute),
		DEFAULT_USER_ROLE:           util.StringPointerOrEmpty(plan.DefaultUserRole),
		SEARCHFILTER:                util.StringPointerOrEmpty(plan.Searchfilter),
		ENDPOINTS_FILTER:            util.StringPointerOrEmpty(plan.EndpointsFilter),
		CREATEACCOUNTJSON:           util.StringPointerOrEmpty(plan.CreateAccountJson),
		UPDATEACCOUNTJSON:           util.StringPointerOrEmpty(plan.UpdateAccountJson),
		REUSEACCOUNTJSON:            util.StringPointerOrEmpty(plan.ReuseAccountJson),
		ENFORCE_TREE_DELETION:       util.StringPointerOrEmpty(plan.EnforceTreeDeletion),
		ADVANCE_FILTER_JSON:         util.StringPointerOrEmpty(plan.AdvanceFilterJson),
		FILTER:                      util.StringPointerOrEmpty(plan.Filter),
		OBJECTFILTER:                util.StringPointerOrEmpty(plan.Objectfilter),
		UPDATEUSERJSON:              util.StringPointerOrEmpty(plan.UpdateUserJson),
		SETRANDOMPASSWORD:           util.StringPointerOrEmpty(plan.Setrandompassword),
		PASSWORD_MIN_LENGTH:         util.StringPointerOrEmpty(plan.PasswordMinLength),
		PASSWORD_MAX_LENGTH:         util.StringPointerOrEmpty(plan.PasswordMaxLength),
		PASSWORD_NOOFCAPSALPHA:      util.StringPointerOrEmpty(plan.PasswordNoofcapsalpha),
		PASSWORD_NOOFSPLCHARS:       util.StringPointerOrEmpty(plan.PasswordNoofsplchars),
		PASSWORD_NOOFDIGITS:         util.StringPointerOrEmpty(plan.PasswordNoofdigits),
		GroupImportMapping:          util.StringPointerOrEmpty(plan.GroupImportMapping),
		UNLOCKACCOUNTJSON:           util.StringPointerOrEmpty(plan.UnlockAccountJson),
		STATUSKEYJSON:               util.StringPointerOrEmpty(plan.StatusKeyJson),
		DISABLEACCOUNTJSON:          util.StringPointerOrEmpty(plan.DisableAccountJson),
		MODIFYUSERDATAJSON:          util.StringPointerOrEmpty(plan.ModifyUserdataJson),
		ORG_BASE:                    util.StringPointerOrEmpty(plan.OrgBase),
		ORGANIZATION_ATTRIBUTE:      util.StringPointerOrEmpty(plan.OrganizationAttribute),
		CREATEORGJSON:               util.StringPointerOrEmpty(plan.Createorgjson),
		UPDATEORGJSON:               util.StringPointerOrEmpty(plan.Updateorgjson),
		ConfigJSON:                  util.StringPointerOrEmpty(plan.ConfigJson),
		PAM_CONFIG:                  util.StringPointerOrEmpty(plan.PamConfig),
		ENABLEGROUPMANAGEMENT:       util.StringPointerOrEmpty(plan.EnableGroupManagement),
		ORGIMPORTJSON:               util.StringPointerOrEmpty(plan.OrgImportJson),
	}
	if plan.VaultConnection.ValueString() != "" {
		adConn.BaseConnector.VaultConnection = util.SafeStringConnector(plan.VaultConnection.ValueString())
		adConn.BaseConnector.VaultConfiguration = util.SafeStringConnector(plan.VaultConfiguration.ValueString())
		adConn.BaseConnector.Saveinvault = util.SafeStringConnector(plan.SaveInVault.ValueString())
	}

	adConnRequest := openapi.CreateOrUpdateRequest{
		ADConnector: &adConn,
	}
	// Initialize API client
	apiResp, _, err := apiClient.ConnectionsAPI.CreateOrUpdate(ctx).CreateOrUpdateRequest(adConnRequest).Execute()
	if err != nil {
		log.Printf("[ERROR] Failed to create API resource. Error: %v", err)
		resp.Diagnostics.AddError("API Create Failed", fmt.Sprintf("Error: %v", err))
		return
	}
	if apiResp != nil && *apiResp.ErrorCode != "0" {
		log.Printf("[ERROR]: Error in creating AD connection resource. Errorcode: %v, Message: %v", *apiResp.ErrorCode, *apiResp.Msg)
		resp.Diagnostics.AddError("Creation of AD connection failed", *apiResp.Msg)
		return
	}

	plan.ID = types.StringValue(fmt.Sprintf("%d", *apiResp.ConnectionKey))
	plan.ConnectionKey = types.Int64Value(int64(*apiResp.ConnectionKey))
	plan.Description = util.SafeStringDatasource(plan.Description.ValueStringPointer())
	plan.DefaultSavRoles = util.SafeStringDatasource(plan.DefaultSavRoles.ValueStringPointer())
	plan.EmailTemplate = util.SafeStringDatasource(plan.EmailTemplate.ValueStringPointer())
	plan.URL = util.SafeStringDatasource(plan.URL.ValueStringPointer())
	plan.Username = util.SafeStringDatasource(plan.Username.ValueStringPointer())
	plan.LdapOrAd = util.SafeStringDatasource(plan.LdapOrAd.ValueStringPointer())
	plan.EntitlementAttribute = util.SafeStringDatasource(plan.EntitlementAttribute.ValueStringPointer())
	plan.CheckForUnique = util.SafeStringDatasource(plan.CheckForUnique.ValueStringPointer())
	plan.GroupSearchBaseDN = util.SafeStringDatasource(plan.GroupSearchBaseDN.ValueStringPointer())
	plan.CreateUpdateMappings = util.SafeStringDatasource(plan.CreateUpdateMappings.ValueStringPointer())
	plan.IncrementalConfig = util.SafeStringDatasource(plan.IncrementalConfig.ValueStringPointer())
	plan.MaxChangeNumber = util.SafeStringDatasource(plan.MaxChangeNumber.ValueStringPointer())
	plan.ReadOperationalAttributes = util.SafeStringDatasource(plan.ReadOperationalAttributes.ValueStringPointer())
	plan.Base = util.SafeStringDatasource(plan.Base.ValueStringPointer())
	plan.DcLocator = util.SafeStringDatasource(plan.DcLocator.ValueStringPointer())
	plan.StatusThresholdConfig = util.SafeStringDatasource(plan.StatusThresholdConfig.ValueStringPointer())
	plan.RemoveAccountAction = util.SafeStringDatasource(plan.RemoveAccountAction.ValueStringPointer())
	plan.AccountAttribute = util.SafeStringDatasource(plan.AccountAttribute.ValueStringPointer())
	plan.AccountNameRule = util.SafeStringDatasource(plan.AccountNameRule.ValueStringPointer())
	plan.Advsearch = util.SafeStringDatasource(plan.Advsearch.ValueStringPointer())
	plan.Setdefaultpagesize = util.SafeStringDatasource(plan.Setdefaultpagesize.ValueStringPointer())
	plan.ResetAndChangePasswrdJson = util.SafeStringDatasource(plan.ResetAndChangePasswrdJson.ValueStringPointer())
	plan.ReuseInactiveAccount = util.SafeStringDatasource(plan.ReuseInactiveAccount.ValueStringPointer())
	plan.ImportJson = util.SafeStringDatasource(plan.ImportJson.ValueStringPointer())
	plan.SupportEmptyString = util.SafeStringDatasource(plan.SupportEmptyString.ValueStringPointer())
	plan.EnableAccountJson = util.SafeStringDatasource(plan.EnableAccountJson.ValueStringPointer())
	plan.PageSize = util.SafeStringDatasource(plan.PageSize.ValueStringPointer())
	plan.UserAttribute = util.SafeStringDatasource(plan.UserAttribute.ValueStringPointer())
	plan.DefaultUserRole = util.SafeStringDatasource(plan.DefaultUserRole.ValueStringPointer())
	plan.Searchfilter = util.SafeStringDatasource(plan.Searchfilter.ValueStringPointer())
	plan.EndpointsFilter = util.SafeStringDatasource(plan.EndpointsFilter.ValueStringPointer())
	plan.CreateAccountJson = util.SafeStringDatasource(plan.CreateAccountJson.ValueStringPointer())
	plan.UpdateAccountJson = util.SafeStringDatasource(plan.UpdateAccountJson.ValueStringPointer())
	plan.ReuseAccountJson = util.SafeStringDatasource(plan.ReuseAccountJson.ValueStringPointer())
	plan.EnforceTreeDeletion = util.SafeStringDatasource(plan.EnforceTreeDeletion.ValueStringPointer())
	plan.AdvanceFilterJson = util.SafeStringDatasource(plan.AdvanceFilterJson.ValueStringPointer())
	plan.Filter = util.SafeStringDatasource(plan.Filter.ValueStringPointer())
	plan.Objectfilter = util.SafeStringDatasource(plan.Objectfilter.ValueStringPointer())
	plan.UpdateUserJson = util.SafeStringDatasource(plan.UpdateUserJson.ValueStringPointer())
	plan.Setrandompassword = util.SafeStringDatasource(plan.Setrandompassword.ValueStringPointer())
	plan.PasswordMinLength = util.SafeStringDatasource(plan.PasswordMinLength.ValueStringPointer())
	plan.PasswordMaxLength = util.SafeStringDatasource(plan.PasswordMaxLength.ValueStringPointer())
	plan.PasswordNoofcapsalpha = util.SafeStringDatasource(plan.PasswordNoofcapsalpha.ValueStringPointer())
	plan.PasswordNoofsplchars = util.SafeStringDatasource(plan.PasswordNoofsplchars.ValueStringPointer())
	plan.PasswordNoofdigits = util.SafeStringDatasource(plan.PasswordNoofdigits.ValueStringPointer())
	plan.GroupImportMapping = util.SafeStringDatasource(plan.GroupImportMapping.ValueStringPointer())
	plan.UnlockAccountJson = util.SafeStringDatasource(plan.UnlockAccountJson.ValueStringPointer())
	plan.StatusKeyJson = util.SafeStringDatasource(plan.StatusKeyJson.ValueStringPointer())
	plan.DisableAccountJson = util.SafeStringDatasource(plan.DisableAccountJson.ValueStringPointer())
	plan.ModifyUserdataJson = util.SafeStringDatasource(plan.ModifyUserdataJson.ValueStringPointer())
	plan.OrgBase = util.SafeStringDatasource(plan.OrgBase.ValueStringPointer())
	plan.OrganizationAttribute = util.SafeStringDatasource(plan.OrganizationAttribute.ValueStringPointer())
	plan.Createorgjson = util.SafeStringDatasource(plan.Createorgjson.ValueStringPointer())
	plan.Updateorgjson = util.SafeStringDatasource(plan.Updateorgjson.ValueStringPointer())
	plan.ConfigJson = util.SafeStringDatasource(plan.ConfigJson.ValueStringPointer())
	plan.PamConfig = util.SafeStringDatasource(plan.PamConfig.ValueStringPointer())
	plan.EnableGroupManagement = util.SafeStringDatasource(plan.EnableGroupManagement.ValueStringPointer())
	plan.OrgImportJson = util.SafeStringDatasource(plan.OrgImportJson.ValueStringPointer())
	plan.Msg = types.StringValue(util.SafeDeref(apiResp.Msg))
	plan.ErrorCode = types.StringValue(util.SafeDeref(apiResp.ErrorCode))
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *adConnectionResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state ADConnectorResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		log.Printf("[ERROR] Failed to get state from request: %v", resp.Diagnostics)
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
		resp.Diagnostics.AddError("API Read Failed In Read Block", fmt.Sprintf("Error: %v", err))
		return
	}
	if apiResp != nil && apiResp.ADConnectionResponse != nil && *apiResp.ADConnectionResponse.Errorcode != 0 {
		log.Printf("[ERROR]: Error in reading AD connection resource. Errorcode: %v, Message: %v", *apiResp.ADConnectionResponse.Errorcode, *apiResp.ADConnectionResponse.Msg)
		resp.Diagnostics.AddError("Reading of AD connection resource failed", *apiResp.ADConnectionResponse.Msg)
		return
	}

	state.ConnectionKey = types.Int64Value(int64(*apiResp.ADConnectionResponse.Connectionkey))
	state.ID = types.StringValue(fmt.Sprintf("%d", *apiResp.ADConnectionResponse.Connectionkey))
	state.ConnectionName = util.SafeStringDatasource(apiResp.ADConnectionResponse.Connectionname)
	state.Description = util.SafeStringDatasource(apiResp.ADConnectionResponse.Description)
	state.DefaultSavRoles = util.SafeStringDatasource(apiResp.ADConnectionResponse.Defaultsavroles)
	state.EmailTemplate = util.SafeStringDatasource(apiResp.ADConnectionResponse.Emailtemplate)
	state.URL = util.SafeStringDatasource(apiResp.ADConnectionResponse.Connectionattributes.URL)
	state.Advsearch = util.SafeStringDatasource(apiResp.ADConnectionResponse.Connectionattributes.ADVSEARCH)
	state.CreateAccountJson = util.SafeStringDatasource(apiResp.ADConnectionResponse.Connectionattributes.CREATEACCOUNTJSON)
	state.DisableAccountJson = util.SafeStringDatasource(apiResp.ADConnectionResponse.Connectionattributes.DISABLEACCOUNTJSON)
	state.GroupSearchBaseDN = util.SafeStringDatasource(apiResp.ADConnectionResponse.Connectionattributes.GroupSearchBaseDN)
	state.PasswordNoofsplchars = util.SafeStringDatasource(apiResp.ADConnectionResponse.Connectionattributes.PASSWORD_NOOFSPLCHARS)
	state.PasswordNoofdigits = util.SafeStringDatasource(apiResp.ADConnectionResponse.Connectionattributes.PASSWORD_NOOFDIGITS)
	state.StatusKeyJson = util.SafeStringDatasource(apiResp.ADConnectionResponse.Connectionattributes.STATUSKEYJSON)
	state.Searchfilter = util.SafeStringDatasource(apiResp.ADConnectionResponse.Connectionattributes.SEARCHFILTER)
	state.ConfigJson = util.SafeStringDatasource(apiResp.ADConnectionResponse.Connectionattributes.ConfigJSON)
	state.RemoveAccountAction = util.SafeStringDatasource(apiResp.ADConnectionResponse.Connectionattributes.REMOVEACCOUNTACTION)
	state.AccountAttribute = util.SafeStringDatasource(apiResp.ADConnectionResponse.Connectionattributes.ACCOUNT_ATTRIBUTE)
	state.AccountNameRule = util.SafeStringDatasource(apiResp.ADConnectionResponse.Connectionattributes.ACCOUNTNAMERULE)
	state.Username = util.SafeStringDatasource(apiResp.ADConnectionResponse.Connectionattributes.USERNAME)
	state.LdapOrAd = util.SafeStringDatasource(apiResp.ADConnectionResponse.Connectionattributes.LDAP_OR_AD)
	state.EntitlementAttribute = util.SafeStringDatasource(apiResp.ADConnectionResponse.Connectionattributes.ENTITLEMENT_ATTRIBUTE)
	state.Setrandompassword = util.SafeStringDatasource(apiResp.ADConnectionResponse.Connectionattributes.SETRANDOMPASSWORD)
	state.PasswordMinLength = util.SafeStringDatasource(apiResp.ADConnectionResponse.Connectionattributes.PASSWORD_MIN_LENGTH)
	state.PasswordMaxLength = util.SafeStringDatasource(apiResp.ADConnectionResponse.Connectionattributes.PASSWORD_MAX_LENGTH)
	state.PasswordNoofcapsalpha = util.SafeStringDatasource(apiResp.ADConnectionResponse.Connectionattributes.PASSWORD_NOOFCAPSALPHA)
	state.Setdefaultpagesize = util.SafeStringDatasource(apiResp.ADConnectionResponse.Connectionattributes.SETDEFAULTPAGESIZE)
	state.ReuseInactiveAccount = util.SafeStringDatasource(apiResp.ADConnectionResponse.Connectionattributes.REUSEINACTIVEACCOUNT)
	state.ImportJson = util.SafeStringDatasource(apiResp.ADConnectionResponse.Connectionattributes.IMPORTJSON)
	state.CreateUpdateMappings = util.SafeStringDatasource(apiResp.ADConnectionResponse.Connectionattributes.CreateUpdateMappings)
	state.AdvanceFilterJson = util.SafeStringDatasource(apiResp.ADConnectionResponse.Connectionattributes.ADVANCE_FILTER_JSON)
	state.PamConfig = util.SafeStringDatasource(apiResp.ADConnectionResponse.Connectionattributes.PAM_CONFIG)
	state.PageSize = util.SafeStringDatasource(apiResp.ADConnectionResponse.Connectionattributes.PAGE_SIZE)
	state.Base = util.SafeStringDatasource(apiResp.ADConnectionResponse.Connectionattributes.BASE)
	state.DcLocator = util.SafeStringDatasource(apiResp.ADConnectionResponse.Connectionattributes.DC_LOCATOR)
	state.StatusThresholdConfig = util.SafeStringDatasource(apiResp.ADConnectionResponse.Connectionattributes.STATUS_THRESHOLD_CONFIG)
	state.ResetAndChangePasswrdJson = util.SafeStringDatasource(apiResp.ADConnectionResponse.Connectionattributes.RESETANDCHANGEPASSWRDJSON)
	state.SupportEmptyString = util.SafeStringDatasource(apiResp.ADConnectionResponse.Connectionattributes.SUPPORTEMPTYSTRING)
	state.ReadOperationalAttributes = util.SafeStringDatasource(apiResp.ADConnectionResponse.Connectionattributes.READ_OPERATIONAL_ATTRIBUTES)
	state.EnableAccountJson = util.SafeStringDatasource(apiResp.ADConnectionResponse.Connectionattributes.ENABLEACCOUNTJSON)
	state.UserAttribute = util.SafeStringDatasource(apiResp.ADConnectionResponse.Connectionattributes.USER_ATTRIBUTE)
	state.DefaultUserRole = util.SafeStringDatasource(apiResp.ADConnectionResponse.Connectionattributes.DEFAULT_USER_ROLE)
	state.EndpointsFilter = util.SafeStringDatasource(apiResp.ADConnectionResponse.Connectionattributes.ENDPOINTS_FILTER)
	state.UpdateAccountJson = util.SafeStringDatasource(apiResp.ADConnectionResponse.Connectionattributes.UPDATEACCOUNTJSON)
	state.ReuseAccountJson = util.SafeStringDatasource(apiResp.ADConnectionResponse.Connectionattributes.REUSEACCOUNTJSON)
	state.EnforceTreeDeletion = util.SafeStringDatasource(apiResp.ADConnectionResponse.Connectionattributes.ENFORCE_TREE_DELETION)
	state.Filter = util.SafeStringDatasource(apiResp.ADConnectionResponse.Connectionattributes.FILTER)
	state.Objectfilter = util.SafeStringDatasource(apiResp.ADConnectionResponse.Connectionattributes.OBJECTFILTER)
	state.UpdateUserJson = util.SafeStringDatasource(apiResp.ADConnectionResponse.Connectionattributes.UPDATEUSERJSON)
	state.GroupImportMapping = util.SafeStringDatasource(apiResp.ADConnectionResponse.Connectionattributes.GroupImportMapping)
	state.UnlockAccountJson = util.SafeStringDatasource(apiResp.ADConnectionResponse.Connectionattributes.UNLOCKACCOUNTJSON)
	state.ModifyUserdataJson = util.SafeStringDatasource(apiResp.ADConnectionResponse.Connectionattributes.MODIFYUSERDATAJSON)
	state.OrgBase = util.SafeStringDatasource(apiResp.ADConnectionResponse.Connectionattributes.ORG_BASE)
	state.OrganizationAttribute = util.SafeStringDatasource(apiResp.ADConnectionResponse.Connectionattributes.ORGANIZATION_ATTRIBUTE)
	state.Createorgjson = util.SafeStringDatasource(apiResp.ADConnectionResponse.Connectionattributes.CREATEORGJSON)
	state.Updateorgjson = util.SafeStringDatasource(apiResp.ADConnectionResponse.Connectionattributes.UPDATEORGJSON)
	state.MaxChangeNumber = util.SafeStringDatasource(apiResp.ADConnectionResponse.Connectionattributes.MAX_CHANGENUMBER)
	state.IncrementalConfig = util.SafeStringDatasource(apiResp.ADConnectionResponse.Connectionattributes.INCREMENTAL_CONFIG)
	state.CheckForUnique = util.SafeStringDatasource(apiResp.ADConnectionResponse.Connectionattributes.CHECKFORUNIQUE)
	state.EnableGroupManagement = util.SafeStringDatasource(apiResp.ADConnectionResponse.Connectionattributes.ENABLEGROUPMANAGEMENT)
	state.OrgImportJson = util.SafeStringDatasource(apiResp.ADConnectionResponse.Connectionattributes.ORGIMPORTJSON)
	stateDiagnostics := resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(stateDiagnostics...)
	if resp.Diagnostics.HasError() {
		log.Printf("[ERROR] Failed to set state in read block: %v", resp.Diagnostics)
		return
	}
}

func (r *adConnectionResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state, config ADConnectorResourceModel
	// Extract state from request
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		log.Printf("[ERROR] Failed to get state from request: %v", resp.Diagnostics)
		return
	}
	// Extract plan from request
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		log.Printf("[ERROR] Failed to get plan from request: %v", resp.Diagnostics)
		return
	}
	//Extract config from request
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		log.Printf("[ERROR] Failed to get config from request: %v", resp.Diagnostics)
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
	adConn := openapi.ADConnector{
		BaseConnector: openapi.BaseConnector{
			//required field
			Connectiontype: "AD",
			ConnectionName: plan.ConnectionName.ValueString(),
			//optional field
			Description:     util.StringPointerOrEmpty(plan.Description),
			Defaultsavroles: util.StringPointerOrEmpty(plan.DefaultSavRoles),
			EmailTemplate:   util.StringPointerOrEmpty(plan.EmailTemplate),
		},
		//required field
		PASSWORD: config.Password.ValueString(),
		//optional field
		URL:                         util.StringPointerOrEmpty(plan.URL),
		USERNAME:                    util.StringPointerOrEmpty(config.Username),
		LDAP_OR_AD:                  util.StringPointerOrEmpty(plan.LdapOrAd),
		ENTITLEMENT_ATTRIBUTE:       util.StringPointerOrEmpty(plan.EntitlementAttribute),
		CHECKFORUNIQUE:              util.StringPointerOrEmpty(plan.CheckForUnique),
		GroupSearchBaseDN:           util.StringPointerOrEmpty(plan.GroupSearchBaseDN),
		CreateUpdateMappings:        util.StringPointerOrEmpty(plan.CreateUpdateMappings),
		INCREMENTAL_CONFIG:          util.StringPointerOrEmpty(plan.IncrementalConfig),
		MAX_CHANGENUMBER:            util.StringPointerOrEmpty(plan.MaxChangeNumber),
		READ_OPERATIONAL_ATTRIBUTES: util.StringPointerOrEmpty(plan.ReadOperationalAttributes),
		BASE:                        util.StringPointerOrEmpty(plan.Base),
		DC_LOCATOR:                  util.StringPointerOrEmpty(plan.DcLocator),
		STATUS_THRESHOLD_CONFIG:     util.StringPointerOrEmpty(plan.StatusThresholdConfig),
		REMOVEACCOUNTACTION:         util.StringPointerOrEmpty(plan.RemoveAccountAction),
		ACCOUNT_ATTRIBUTE:           util.StringPointerOrEmpty(plan.AccountAttribute),
		ACCOUNTNAMERULE:             util.StringPointerOrEmpty(plan.AccountNameRule),
		ADVSEARCH:                   util.StringPointerOrEmpty(plan.Advsearch),
		SETDEFAULTPAGESIZE:          util.StringPointerOrEmpty(plan.Setdefaultpagesize),
		RESETANDCHANGEPASSWRDJSON:   util.StringPointerOrEmpty(plan.ResetAndChangePasswrdJson),
		REUSEINACTIVEACCOUNT:        util.StringPointerOrEmpty(plan.ReuseInactiveAccount),
		IMPORTJSON:                  util.StringPointerOrEmpty(plan.ImportJson),
		SUPPORTEMPTYSTRING:          util.StringPointerOrEmpty(plan.SupportEmptyString),
		ENABLEACCOUNTJSON:           util.StringPointerOrEmpty(plan.EnableAccountJson),
		PAGE_SIZE:                   util.StringPointerOrEmpty(plan.PageSize),
		USER_ATTRIBUTE:              util.StringPointerOrEmpty(plan.UserAttribute),
		DEFAULT_USER_ROLE:           util.StringPointerOrEmpty(plan.DefaultUserRole),
		SEARCHFILTER:                util.StringPointerOrEmpty(plan.Searchfilter),
		ENDPOINTS_FILTER:            util.StringPointerOrEmpty(plan.EndpointsFilter),
		CREATEACCOUNTJSON:           util.StringPointerOrEmpty(plan.CreateAccountJson),
		UPDATEACCOUNTJSON:           util.StringPointerOrEmpty(plan.UpdateAccountJson),
		REUSEACCOUNTJSON:            util.StringPointerOrEmpty(plan.ReuseAccountJson),
		ENFORCE_TREE_DELETION:       util.StringPointerOrEmpty(plan.EnforceTreeDeletion),
		ADVANCE_FILTER_JSON:         util.StringPointerOrEmpty(plan.AdvanceFilterJson),
		FILTER:                      util.StringPointerOrEmpty(plan.Filter),
		OBJECTFILTER:                util.StringPointerOrEmpty(plan.Objectfilter),
		UPDATEUSERJSON:              util.StringPointerOrEmpty(plan.UpdateUserJson),
		SETRANDOMPASSWORD:           util.StringPointerOrEmpty(plan.Setrandompassword),
		PASSWORD_MIN_LENGTH:         util.StringPointerOrEmpty(plan.PasswordMinLength),
		PASSWORD_MAX_LENGTH:         util.StringPointerOrEmpty(plan.PasswordMaxLength),
		PASSWORD_NOOFCAPSALPHA:      util.StringPointerOrEmpty(plan.PasswordNoofcapsalpha),
		PASSWORD_NOOFSPLCHARS:       util.StringPointerOrEmpty(plan.PasswordNoofsplchars),
		PASSWORD_NOOFDIGITS:         util.StringPointerOrEmpty(plan.PasswordNoofdigits),
		GroupImportMapping:          util.StringPointerOrEmpty(plan.GroupImportMapping),
		UNLOCKACCOUNTJSON:           util.StringPointerOrEmpty(plan.UnlockAccountJson),
		STATUSKEYJSON:               util.StringPointerOrEmpty(plan.StatusKeyJson),
		DISABLEACCOUNTJSON:          util.StringPointerOrEmpty(plan.DisableAccountJson),
		MODIFYUSERDATAJSON:          util.StringPointerOrEmpty(plan.ModifyUserdataJson),
		ORG_BASE:                    util.StringPointerOrEmpty(plan.OrgBase),
		ORGANIZATION_ATTRIBUTE:      util.StringPointerOrEmpty(plan.OrganizationAttribute),
		CREATEORGJSON:               util.StringPointerOrEmpty(plan.Createorgjson),
		UPDATEORGJSON:               util.StringPointerOrEmpty(plan.Updateorgjson),
		ConfigJSON:                  util.StringPointerOrEmpty(plan.ConfigJson),
		PAM_CONFIG:                  util.StringPointerOrEmpty(plan.PamConfig),
		ENABLEGROUPMANAGEMENT:       util.StringPointerOrEmpty(plan.EnableGroupManagement),
		ORGIMPORTJSON:               util.StringPointerOrEmpty(plan.OrgImportJson),
	}
	if plan.VaultConnection.ValueString() != "" {
		adConn.BaseConnector.VaultConnection = util.SafeStringConnector(plan.VaultConnection.ValueString())
		adConn.BaseConnector.VaultConfiguration = util.SafeStringConnector(plan.VaultConfiguration.ValueString())
		adConn.BaseConnector.Saveinvault = util.SafeStringConnector(plan.SaveInVault.ValueString())
	} else {
		emptyStr := ""
		adConn.BaseConnector.VaultConnection = &emptyStr
		adConn.BaseConnector.VaultConfiguration = &emptyStr
		adConn.BaseConnector.Saveinvault = &emptyStr
	}
	adConnRequest := openapi.CreateOrUpdateRequest{
		ADConnector: &adConn,
	}

	// Initialize API client
	apiClient := openapi.NewAPIClient(cfg)
	apiResp, _, err := apiClient.ConnectionsAPI.CreateOrUpdate(ctx).CreateOrUpdateRequest(adConnRequest).Execute()
	if err != nil {
		log.Printf("Problem with the update function")
		resp.Diagnostics.AddError("API Update Failed", fmt.Sprintf("Error: %v", err))
		return
	}
	if apiResp != nil && *apiResp.ErrorCode != "0" {
		log.Printf("[ERROR]: Error in updating AD connection resource. Errorcode: %v, Message: %v", *apiResp.ErrorCode, *apiResp.Msg)
		resp.Diagnostics.AddError("Updation of AD connection failed", *apiResp.Msg)
		return
	}

	reqParams := openapi.GetConnectionDetailsRequest{}

	reqParams.SetConnectionname(plan.ConnectionName.ValueString())
	getResp, _, err := apiClient.ConnectionsAPI.GetConnectionDetails(ctx).GetConnectionDetailsRequest(reqParams).Execute()
	if err != nil {
		log.Printf("Problem with the get function in update block")
		resp.Diagnostics.AddError("API Read Failed In Update Block", fmt.Sprintf("Error: %v", err))
		return
	}
	if getResp != nil && getResp.ADConnectionResponse != nil && *getResp.ADConnectionResponse.Errorcode != 0 {
		log.Printf("[ERROR]: Error in reading AD connection resource after updation. Errorcode: %v, Message: %v", *getResp.ADConnectionResponse.Errorcode, *getResp.ADConnectionResponse.Msg)
		resp.Diagnostics.AddError("Reading of AD connection after updation failed", *getResp.ADConnectionResponse.Msg)
		return
	}

	plan.ConnectionKey = types.Int64Value(int64(*getResp.ADConnectionResponse.Connectionkey))
	plan.ID = types.StringValue(fmt.Sprintf("%d", *getResp.ADConnectionResponse.Connectionkey))
	plan.ConnectionName = util.SafeStringDatasource(getResp.ADConnectionResponse.Connectionname)
	plan.Description = util.SafeStringDatasource(getResp.ADConnectionResponse.Description)
	plan.DefaultSavRoles = util.SafeStringDatasource(getResp.ADConnectionResponse.Defaultsavroles)
	plan.EmailTemplate = util.SafeStringDatasource(getResp.ADConnectionResponse.Emailtemplate)
	plan.URL = util.SafeStringDatasource(getResp.ADConnectionResponse.Connectionattributes.URL)
	plan.Advsearch = util.SafeStringDatasource(getResp.ADConnectionResponse.Connectionattributes.ADVSEARCH)
	plan.CreateAccountJson = util.SafeStringDatasource(getResp.ADConnectionResponse.Connectionattributes.CREATEACCOUNTJSON)
	plan.DisableAccountJson = util.SafeStringDatasource(getResp.ADConnectionResponse.Connectionattributes.DISABLEACCOUNTJSON)
	plan.GroupSearchBaseDN = util.SafeStringDatasource(getResp.ADConnectionResponse.Connectionattributes.GroupSearchBaseDN)
	plan.PasswordNoofsplchars = util.SafeStringDatasource(getResp.ADConnectionResponse.Connectionattributes.PASSWORD_NOOFSPLCHARS)
	plan.PasswordNoofdigits = util.SafeStringDatasource(getResp.ADConnectionResponse.Connectionattributes.PASSWORD_NOOFDIGITS)
	plan.StatusKeyJson = util.SafeStringDatasource(getResp.ADConnectionResponse.Connectionattributes.STATUSKEYJSON)
	plan.Searchfilter = util.SafeStringDatasource(getResp.ADConnectionResponse.Connectionattributes.SEARCHFILTER)
	plan.ConfigJson = util.SafeStringDatasource(getResp.ADConnectionResponse.Connectionattributes.ConfigJSON)
	plan.RemoveAccountAction = util.SafeStringDatasource(getResp.ADConnectionResponse.Connectionattributes.REMOVEACCOUNTACTION)
	plan.AccountAttribute = util.SafeStringDatasource(getResp.ADConnectionResponse.Connectionattributes.ACCOUNT_ATTRIBUTE)
	plan.AccountNameRule = util.SafeStringDatasource(getResp.ADConnectionResponse.Connectionattributes.ACCOUNTNAMERULE)
	plan.Username = util.SafeStringDatasource(getResp.ADConnectionResponse.Connectionattributes.USERNAME)
	plan.LdapOrAd = util.SafeStringDatasource(getResp.ADConnectionResponse.Connectionattributes.LDAP_OR_AD)
	plan.EntitlementAttribute = util.SafeStringDatasource(getResp.ADConnectionResponse.Connectionattributes.ENTITLEMENT_ATTRIBUTE)
	plan.Setrandompassword = util.SafeStringDatasource(getResp.ADConnectionResponse.Connectionattributes.SETRANDOMPASSWORD)
	plan.PasswordMinLength = util.SafeStringDatasource(getResp.ADConnectionResponse.Connectionattributes.PASSWORD_MIN_LENGTH)
	plan.PasswordMaxLength = util.SafeStringDatasource(getResp.ADConnectionResponse.Connectionattributes.PASSWORD_MAX_LENGTH)
	plan.PasswordNoofcapsalpha = util.SafeStringDatasource(getResp.ADConnectionResponse.Connectionattributes.PASSWORD_NOOFCAPSALPHA)
	plan.Setdefaultpagesize = util.SafeStringDatasource(getResp.ADConnectionResponse.Connectionattributes.SETDEFAULTPAGESIZE)
	plan.ReuseInactiveAccount = util.SafeStringDatasource(getResp.ADConnectionResponse.Connectionattributes.REUSEINACTIVEACCOUNT)
	plan.ImportJson = util.SafeStringDatasource(getResp.ADConnectionResponse.Connectionattributes.IMPORTJSON)
	plan.CreateUpdateMappings = util.SafeStringDatasource(getResp.ADConnectionResponse.Connectionattributes.CreateUpdateMappings)
	plan.AdvanceFilterJson = util.SafeStringDatasource(getResp.ADConnectionResponse.Connectionattributes.ADVANCE_FILTER_JSON)
	plan.PamConfig = util.SafeStringDatasource(getResp.ADConnectionResponse.Connectionattributes.PAM_CONFIG)
	plan.PageSize = util.SafeStringDatasource(getResp.ADConnectionResponse.Connectionattributes.PAGE_SIZE)
	plan.Base = util.SafeStringDatasource(getResp.ADConnectionResponse.Connectionattributes.BASE)
	plan.DcLocator = util.SafeStringDatasource(getResp.ADConnectionResponse.Connectionattributes.DC_LOCATOR)
	plan.StatusThresholdConfig = util.SafeStringDatasource(getResp.ADConnectionResponse.Connectionattributes.STATUS_THRESHOLD_CONFIG)
	plan.ResetAndChangePasswrdJson = util.SafeStringDatasource(getResp.ADConnectionResponse.Connectionattributes.RESETANDCHANGEPASSWRDJSON)
	plan.SupportEmptyString = util.SafeStringDatasource(getResp.ADConnectionResponse.Connectionattributes.SUPPORTEMPTYSTRING)
	plan.ReadOperationalAttributes = util.SafeStringDatasource(getResp.ADConnectionResponse.Connectionattributes.READ_OPERATIONAL_ATTRIBUTES)
	plan.EnableAccountJson = util.SafeStringDatasource(getResp.ADConnectionResponse.Connectionattributes.ENABLEACCOUNTJSON)
	plan.UserAttribute = util.SafeStringDatasource(getResp.ADConnectionResponse.Connectionattributes.USER_ATTRIBUTE)
	plan.DefaultUserRole = util.SafeStringDatasource(getResp.ADConnectionResponse.Connectionattributes.DEFAULT_USER_ROLE)
	plan.EndpointsFilter = util.SafeStringDatasource(getResp.ADConnectionResponse.Connectionattributes.ENDPOINTS_FILTER)
	plan.UpdateAccountJson = util.SafeStringDatasource(getResp.ADConnectionResponse.Connectionattributes.UPDATEACCOUNTJSON)
	plan.ReuseAccountJson = util.SafeStringDatasource(getResp.ADConnectionResponse.Connectionattributes.REUSEACCOUNTJSON)
	plan.EnforceTreeDeletion = util.SafeStringDatasource(getResp.ADConnectionResponse.Connectionattributes.ENFORCE_TREE_DELETION)
	plan.Filter = util.SafeStringDatasource(getResp.ADConnectionResponse.Connectionattributes.FILTER)
	plan.Objectfilter = util.SafeStringDatasource(getResp.ADConnectionResponse.Connectionattributes.OBJECTFILTER)
	plan.UpdateUserJson = util.SafeStringDatasource(getResp.ADConnectionResponse.Connectionattributes.UPDATEUSERJSON)
	plan.GroupImportMapping = util.SafeStringDatasource(getResp.ADConnectionResponse.Connectionattributes.GroupImportMapping)
	plan.UnlockAccountJson = util.SafeStringDatasource(getResp.ADConnectionResponse.Connectionattributes.UNLOCKACCOUNTJSON)
	plan.ModifyUserdataJson = util.SafeStringDatasource(getResp.ADConnectionResponse.Connectionattributes.MODIFYUSERDATAJSON)
	plan.OrgBase = util.SafeStringDatasource(getResp.ADConnectionResponse.Connectionattributes.ORG_BASE)
	plan.OrganizationAttribute = util.SafeStringDatasource(getResp.ADConnectionResponse.Connectionattributes.ORGANIZATION_ATTRIBUTE)
	plan.Createorgjson = util.SafeStringDatasource(getResp.ADConnectionResponse.Connectionattributes.CREATEORGJSON)
	plan.Updateorgjson = util.SafeStringDatasource(getResp.ADConnectionResponse.Connectionattributes.UPDATEORGJSON)
	plan.MaxChangeNumber = util.SafeStringDatasource(getResp.ADConnectionResponse.Connectionattributes.MAX_CHANGENUMBER)
	plan.IncrementalConfig = util.SafeStringDatasource(getResp.ADConnectionResponse.Connectionattributes.INCREMENTAL_CONFIG)
	plan.CheckForUnique = util.SafeStringDatasource(getResp.ADConnectionResponse.Connectionattributes.CHECKFORUNIQUE)
	plan.EnableGroupManagement = util.SafeStringDatasource(getResp.ADConnectionResponse.Connectionattributes.ENABLEGROUPMANAGEMENT)
	plan.OrgImportJson = util.SafeStringDatasource(getResp.ADConnectionResponse.Connectionattributes.ORGIMPORTJSON)
	plan.Msg = types.StringValue(util.SafeDeref(apiResp.Msg))
	plan.ErrorCode = types.StringValue(util.SafeDeref(apiResp.ErrorCode))
	stateUpdateDiagnostics := resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(stateUpdateDiagnostics...)
}

func (r *adConnectionResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
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
func (r *adConnectionResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	// Retrieve import ID and save to id attribute
	resource.ImportStatePassthroughID(ctx, path.Root("connection_name"), req, resp)
}
