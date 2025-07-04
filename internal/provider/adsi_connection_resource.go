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

// saviynt_adsi_connection_resource manages ADSI connectors in the Saviynt Security Manager.
// The resource implements the full Terraform lifecycle:
//   - Create: provisions a new ADSI connector using the supplied configuration.
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

	openapi "github.com/saviynt/saviynt-api-go-client/connections"

	s "github.com/saviynt/saviynt-api-go-client"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &adsiConnectionResource{}
var _ resource.ResourceWithImportState = &adsiConnectionResource{}

type ADSIConnectorResourceModel struct {
	BaseConnectorResourceModel
	ID                          types.String `tfsdk:"id"`
	URL                         types.String `tfsdk:"url"`
	Username                    types.String `tfsdk:"username"`
	Password                    types.String `tfsdk:"password"`
	ConnectionUrl               types.String `tfsdk:"connection_url"`
	ProvisioningUrl             types.String `tfsdk:"provisioning_url"`
	ForestList                  types.String `tfsdk:"forestlist"`
	DefaultUserRole             types.String `tfsdk:"default_user_role"`
	UpdateUserJson              types.String `tfsdk:"updateuserjson"`
	EndpointsFilter             types.String `tfsdk:"endpoints_filter"`
	SearchFilter                types.String `tfsdk:"searchfilter"`
	ObjectFilter                types.String `tfsdk:"objectfilter"`
	AccountAttribute            types.String `tfsdk:"account_attribute"`
	StatusThresholdConfig       types.String `tfsdk:"status_threshold_config"`
	EntitlementAttribute        types.String `tfsdk:"entitlement_attribute"`
	UserAttribute               types.String `tfsdk:"user_attribute"`
	GroupSearchBaseDN           types.String `tfsdk:"group_search_base_dn"`
	CheckForUnique              types.String `tfsdk:"checkforunique"`
	StatusKeyJson               types.String `tfsdk:"statuskeyjson"`
	GroupImportMapping          types.String `tfsdk:"group_import_mapping"`
	ImportNestedMembership      types.String `tfsdk:"import_nested_membership"`
	PageSize                    types.String `tfsdk:"page_size"`
	AccountNameRule             types.String `tfsdk:"accountnamerule"`
	CreateAccountJson           types.String `tfsdk:"createaccountjson"`
	UpdateAccountJson           types.String `tfsdk:"updateaccountjson"`
	EnableAccountJson           types.String `tfsdk:"enableaccountjson"`
	DisableAccountJson          types.String `tfsdk:"disableaccountjson"`
	RemoveAccountJson           types.String `tfsdk:"removeaccountjson"`
	AddAccessJson               types.String `tfsdk:"addaccessjson"`
	RemoveAccessJson            types.String `tfsdk:"removeaccessjson"`
	ResetAndChangePasswrdJson   types.String `tfsdk:"resetandchangepasswrdjson"`
	CreateGroupJson             types.String `tfsdk:"creategroupjson"`
	UpdateGroupJson             types.String `tfsdk:"updategroupjson"`
	RemoveGroupJson             types.String `tfsdk:"removegroupjson"`
	AddAccessEntitlementJson    types.String `tfsdk:"addaccessentitlementjson"`
	CustomConfigJson            types.String `tfsdk:"customconfigjson"`
	RemoveAccessEntitlementJson types.String `tfsdk:"removeaccessentitlementjson"`
	CreateServiceAccountJson    types.String `tfsdk:"createserviceaccountjson"`
	UpdateServiceAccountJson    types.String `tfsdk:"updateserviceaccountjson"`
	RemoveServiceAccountJson    types.String `tfsdk:"removeserviceaccountjson"`
	PamConfig                   types.String `tfsdk:"pam_config"`
	ModifyUserDataJson          types.String `tfsdk:"modifyuserdatajson"`
}

type adsiConnectionResource struct {
	client *s.Client
	token  string
}

func NewADSITestConnectionResource() resource.Resource {
	return &adsiConnectionResource{}
}

func (r *adsiConnectionResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = "saviynt_adsi_connection_resource"
}

func ADSIConnectorResourceSchema() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:    true,
			Description: "Resource ID.",
		},
		"url": schema.StringAttribute{
			Required:    true,
			Description: "Primary/root domain URL list (comma Separated)",
		},
		"username": schema.StringAttribute{
			Required:    true,
			WriteOnly:   true,
			Description: "Service account username",
		},
		"password": schema.StringAttribute{
			Required:    true,
			WriteOnly:   true,
			Description: "Service account password",
		},
		"connection_url": schema.StringAttribute{
			Required:    true,
			Description: "ADSI remote agent Connection URL",
		},
		"forestlist": schema.StringAttribute{
			Required:    true,
			Description: "Forest List (Comma Separated) which we need to manage",
		},
		"provisioning_url": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "ADSI remote agent Provisioning URL",
		},
		"default_user_role": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "Default SAV Role to be assigned to all the new users that gets imported via User Import",
		},
		"updateuserjson": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "Specify the attribute Value which will be used to Update existing User",
		},
		"endpoints_filter": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "Provide the configuration to create Child Endpoints and import associated accounts and entitlements",
		},
		"searchfilter": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "Account Search Filter to specify the starting point of the directory from where the accounts needs to be imported. You can have multiple BaseDNs here separated by ###.",
		},
		"objectfilter": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "Object Filter is used to filter the objects that will be returned.This filter will be same for all domains.",
		},
		"account_attribute": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "Map EIC and AD attributes for account import (AD attributes must be in lower case)",
		},
		"status_threshold_config": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "Account status and threshold related config",
		},
		"entitlement_attribute": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "Account attribute that contains group membership",
		},
		"user_attribute": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "Map EIC and AD attributes for user import (AD attributes must be in lower case)",
		},
		"group_search_base_dn": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "Group Search Filter to specify the starting point of the directory from where the groups needs to be imported. You can have multiple BaseDNs here separated by ###.",
		},
		"checkforunique": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "Evaluate the uniqueness of an attribute",
		},
		"statuskeyjson": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "JSON configuration to specify Users status",
		},
		"group_import_mapping": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "Map AD group attribute to EIC entitlement attribute for import",
		},
		"import_nested_membership": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "Specify if you want the connector to import all indirect or nested membership of an account or a group during access import",
		},
		"page_size": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "Page size defines the number of objects to be returned from each AD operation.",
		},
		"accountnamerule": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "Rule to generate account name.",
		},
		"createaccountjson": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "Specify the attributes values which will be used to Create the New Account.",
		},
		"updateaccountjson": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "Specify the attributes values which will be used to Update existing Account.",
		},
		"enableaccountjson": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "Specify the actions and attribute updates to be performed for enabling an account.",
		},
		"disableaccountjson": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "Specify the actions and attributes updates to be performed for disabling an account.",
		},
		"removeaccountjson": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "Specify the actions to be performed for deleting an account.",
		},
		"addaccessjson": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "Configuration to ADD Access (cross domain/forest group membership) to an account.",
		},
		"removeaccessjson": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "Configuration to REMOVE Access (cross domain/forest group membership) to an account.",
		},
		"resetandchangepasswrdjson": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "Configuration to Reset and Change Password.",
		},
		"creategroupjson": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "Configuration to Create a Group",
		},
		"updategroupjson": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "Configuration to Update a Group",
		},
		"removegroupjson": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "Configuration to Delete a Group",
		},
		"addaccessentitlementjson": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "Configuration to Add nested group hierarchy",
		},
		"customconfigjson": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "Custom configuration JSON",
		},
		"removeaccessentitlementjson": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "Configuration to Remove nested group hierarchy",
		},
		"createserviceaccountjson": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "Specify the Field Value which will be used to Create the New Service Account.",
		},
		"updateserviceaccountjson": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "Specify the Field Value which will be used to update the existing Service Account.",
		},
		"removeserviceaccountjson": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "Specify the actions to be performed while deleting a service account.",
		},
		"pam_config": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "JSON to specify Bootstrap Config.",
		},
		"modifyuserdatajson": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "Specify this parameter to transform the data during user import.",
		},
	}
}

func (r *adsiConnectionResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: util.ADSIConnDescription,
		Attributes:  connectionsutil.MergeResourceAttributes(BaseConnectorResourceSchema(), ADSIConnectorResourceSchema()),
	}
}

func (r *adsiConnectionResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *adsiConnectionResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan, config ADSIConnectorResourceModel
	// Extract plan from request
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}
	// Extract config from request
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
	existingResource, _, _ := apiClient.ConnectionsAPI.GetConnectionDetails(ctx).GetConnectionDetailsRequest(reqParams).Execute()
	if existingResource != nil &&
		existingResource.ADSIConnectionResponse != nil &&
		existingResource.ADSIConnectionResponse.Errorcode != nil &&
		*existingResource.ADSIConnectionResponse.Errorcode == 0 {
		log.Printf("[ERROR] Connection name already exists. Please import or use a different name")
		resp.Diagnostics.AddError("API Create Failed", "Connection name already exists. Please import or use a different name")
		return
	}

	if plan.EntitlementAttribute.IsNull() || plan.EntitlementAttribute.IsUnknown() {
		plan.EntitlementAttribute = types.StringValue("memberOf")
	}
	adsiConn := openapi.ADSIConnector{
		BaseConnector: openapi.BaseConnector{
			//required values
			Connectiontype: "ADSI",
			ConnectionName: plan.ConnectionName.ValueString(),
			//optional values
			// Description:     util.StringPointerOrEmpty(plan.Description),
			Defaultsavroles: util.StringPointerOrEmpty(plan.DefaultSavRoles),
			EmailTemplate:   util.StringPointerOrEmpty(plan.EmailTemplate),
		},
		//required values
		URL:            plan.URL.ValueString(),
		USERNAME:       config.Username.ValueString(),
		PASSWORD:       config.Password.ValueString(),
		CONNECTION_URL: plan.ConnectionUrl.ValueString(),
		FORESTLIST:     plan.ForestList.ValueString(),
		//optional values
		PROVISIONING_URL:            util.StringPointerOrEmpty(plan.ProvisioningUrl),
		DEFAULT_USER_ROLE:           util.StringPointerOrEmpty(plan.DefaultUserRole),
		UPDATEUSERJSON:              util.StringPointerOrEmpty(plan.UpdateUserJson),
		ENDPOINTS_FILTER:            util.StringPointerOrEmpty(plan.EndpointsFilter),
		SEARCHFILTER:                util.StringPointerOrEmpty(plan.SearchFilter),
		OBJECTFILTER:                util.StringPointerOrEmpty(plan.ObjectFilter),
		ACCOUNT_ATTRIBUTE:           util.StringPointerOrEmpty(plan.AccountAttribute),
		STATUS_THRESHOLD_CONFIG:     util.StringPointerOrEmpty(plan.StatusThresholdConfig),
		ENTITLEMENT_ATTRIBUTE:       util.StringPointerOrEmpty(plan.EntitlementAttribute),
		USER_ATTRIBUTE:              util.StringPointerOrEmpty(plan.UserAttribute),
		GroupSearchBaseDN:           util.StringPointerOrEmpty(plan.GroupSearchBaseDN),
		CHECKFORUNIQUE:              util.StringPointerOrEmpty(plan.CheckForUnique),
		STATUSKEYJSON:               util.StringPointerOrEmpty(plan.StatusKeyJson),
		GroupImportMapping:          util.StringPointerOrEmpty(plan.GroupImportMapping),
		ImportNestedMembership:      util.StringPointerOrEmpty(plan.ImportNestedMembership),
		PAGE_SIZE:                   util.StringPointerOrEmpty(plan.PageSize),
		ACCOUNTNAMERULE:             util.StringPointerOrEmpty(plan.AccountNameRule),
		CREATEACCOUNTJSON:           util.StringPointerOrEmpty(plan.CreateAccountJson),
		UPDATEACCOUNTJSON:           util.StringPointerOrEmpty(plan.UpdateAccountJson),
		ENABLEACCOUNTJSON:           util.StringPointerOrEmpty(plan.EnableAccountJson),
		DISABLEACCOUNTJSON:          util.StringPointerOrEmpty(plan.DisableAccountJson),
		REMOVEACCOUNTJSON:           util.StringPointerOrEmpty(plan.RemoveAccountJson),
		ADDACCESSJSON:               util.StringPointerOrEmpty(plan.AddAccessJson),
		REMOVEACCESSJSON:            util.StringPointerOrEmpty(plan.RemoveAccessJson),
		RESETANDCHANGEPASSWRDJSON:   util.StringPointerOrEmpty(plan.ResetAndChangePasswrdJson),
		CREATEGROUPJSON:             util.StringPointerOrEmpty(plan.CreateGroupJson),
		UPDATEGROUPJSON:             util.StringPointerOrEmpty(plan.UpdateGroupJson),
		REMOVEGROUPJSON:             util.StringPointerOrEmpty(plan.RemoveGroupJson),
		ADDACCESSENTITLEMENTJSON:    util.StringPointerOrEmpty(plan.AddAccessEntitlementJson),
		CUSTOMCONFIGJSON:            util.StringPointerOrEmpty(plan.CustomConfigJson),
		REMOVEACCESSENTITLEMENTJSON: util.StringPointerOrEmpty(plan.RemoveAccessEntitlementJson),
		CREATESERVICEACCOUNTJSON:    util.StringPointerOrEmpty(plan.CreateServiceAccountJson),
		UPDATESERVICEACCOUNTJSON:    util.StringPointerOrEmpty(plan.UpdateServiceAccountJson),
		REMOVESERVICEACCOUNTJSON:    util.StringPointerOrEmpty(plan.RemoveServiceAccountJson),
		PAM_CONFIG:                  util.StringPointerOrEmpty(plan.PamConfig),
		MODIFYUSERDATAJSON:          util.StringPointerOrEmpty(plan.ModifyUserDataJson),
	}
	if plan.VaultConnection.ValueString() != "" {
		adsiConn.BaseConnector.VaultConnection = util.SafeStringConnector(plan.VaultConnection.ValueString())
		adsiConn.BaseConnector.VaultConfiguration = util.SafeStringConnector(plan.VaultConfiguration.ValueString())
		adsiConn.BaseConnector.Saveinvault = util.SafeStringConnector(plan.SaveInVault.ValueString())
	}
	adsiConnRequest := openapi.CreateOrUpdateRequest{
		ADSIConnector: &adsiConn,
	}

	// Initialize API client
	apiResp, _, err := apiClient.ConnectionsAPI.CreateOrUpdate(ctx).CreateOrUpdateRequest(adsiConnRequest).Execute()
	if err != nil {
		log.Printf("[ERROR] Failed to create API resource. Error: %v", err)
		resp.Diagnostics.AddError("API Create Failed", fmt.Sprintf("Error: %v", err))
		return
	}
	if apiResp!=nil && *apiResp.ErrorCode !="0"{
		log.Printf("[ERROR]: Error in creating ADSI connection resource. Errorcode: %v, Message: %v", *apiResp.ErrorCode, *apiResp.Msg)
		resp.Diagnostics.AddError("Creation of ADSI connection failed", *apiResp.Msg)
		return
	}

	plan.ID = types.StringValue(fmt.Sprintf("%d", *apiResp.ConnectionKey))
	plan.ConnectionType = types.StringValue("ADSI")
	plan.ConnectionKey = types.Int64Value(int64(*apiResp.ConnectionKey))
	// plan.Description = util.SafeStringDatasource(plan.Description.ValueStringPointer())
	plan.DefaultSavRoles = util.SafeStringDatasource(plan.DefaultSavRoles.ValueStringPointer())
	plan.EmailTemplate = util.SafeStringDatasource(plan.EmailTemplate.ValueStringPointer())
	plan.ProvisioningUrl = util.SafeStringDatasource(plan.ProvisioningUrl.ValueStringPointer())
	plan.DefaultUserRole = util.SafeStringDatasource(plan.DefaultUserRole.ValueStringPointer())
	plan.UpdateUserJson = util.SafeStringDatasource(plan.UpdateUserJson.ValueStringPointer())
	plan.EndpointsFilter = util.SafeStringDatasource(plan.EndpointsFilter.ValueStringPointer())
	plan.SearchFilter = util.SafeStringDatasource(plan.SearchFilter.ValueStringPointer())
	plan.ObjectFilter = util.SafeStringDatasource(plan.ObjectFilter.ValueStringPointer())
	plan.AccountAttribute = util.SafeStringDatasource(plan.AccountAttribute.ValueStringPointer())
	plan.StatusThresholdConfig = util.SafeStringDatasource(plan.StatusThresholdConfig.ValueStringPointer())
	plan.EntitlementAttribute = util.SafeStringDatasource(plan.EntitlementAttribute.ValueStringPointer())
	plan.UserAttribute = util.SafeStringDatasource(plan.UserAttribute.ValueStringPointer())
	plan.GroupSearchBaseDN = util.SafeStringDatasource(plan.GroupSearchBaseDN.ValueStringPointer())
	plan.CheckForUnique = util.SafeStringDatasource(plan.CheckForUnique.ValueStringPointer())
	plan.StatusKeyJson = util.SafeStringDatasource(plan.StatusKeyJson.ValueStringPointer())
	plan.GroupImportMapping = util.SafeStringDatasource(plan.GroupImportMapping.ValueStringPointer())
	plan.ImportNestedMembership = util.SafeStringDatasource(plan.ImportNestedMembership.ValueStringPointer())
	plan.PageSize = util.SafeStringDatasource(plan.PageSize.ValueStringPointer())
	plan.AccountNameRule = util.SafeStringDatasource(plan.AccountNameRule.ValueStringPointer())
	plan.CreateAccountJson = util.SafeStringDatasource(plan.CreateAccountJson.ValueStringPointer())
	plan.UpdateAccountJson = util.SafeStringDatasource(plan.UpdateAccountJson.ValueStringPointer())
	plan.EnableAccountJson = util.SafeStringDatasource(plan.EnableAccountJson.ValueStringPointer())
	plan.DisableAccountJson = util.SafeStringDatasource(plan.DisableAccountJson.ValueStringPointer())
	plan.RemoveAccountJson = util.SafeStringDatasource(plan.RemoveAccountJson.ValueStringPointer())
	plan.AddAccessJson = util.SafeStringDatasource(plan.AddAccessJson.ValueStringPointer())
	plan.RemoveAccessJson = util.SafeStringDatasource(plan.RemoveAccessJson.ValueStringPointer())
	plan.ResetAndChangePasswrdJson = util.SafeStringDatasource(plan.ResetAndChangePasswrdJson.ValueStringPointer())
	plan.CreateGroupJson = util.SafeStringDatasource(plan.CreateGroupJson.ValueStringPointer())
	plan.UpdateGroupJson = util.SafeStringDatasource(plan.UpdateGroupJson.ValueStringPointer())
	plan.RemoveGroupJson = util.SafeStringDatasource(plan.RemoveGroupJson.ValueStringPointer())
	plan.AddAccessEntitlementJson = util.SafeStringDatasource(plan.AddAccessEntitlementJson.ValueStringPointer())
	plan.CustomConfigJson = util.SafeStringDatasource(plan.CustomConfigJson.ValueStringPointer())
	plan.RemoveAccessEntitlementJson = util.SafeStringDatasource(plan.RemoveAccessEntitlementJson.ValueStringPointer())
	plan.CreateServiceAccountJson = util.SafeStringDatasource(plan.CreateServiceAccountJson.ValueStringPointer())
	plan.UpdateServiceAccountJson = util.SafeStringDatasource(plan.UpdateServiceAccountJson.ValueStringPointer())
	plan.RemoveServiceAccountJson = util.SafeStringDatasource(plan.RemoveServiceAccountJson.ValueStringPointer())
	plan.PamConfig = util.SafeStringDatasource(plan.PamConfig.ValueStringPointer())
	plan.ModifyUserDataJson = util.SafeStringDatasource(plan.ModifyUserDataJson.ValueStringPointer())
	plan.Msg = types.StringValue(util.SafeDeref(apiResp.Msg))
	plan.ErrorCode = types.StringValue(util.SafeDeref(apiResp.ErrorCode))
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *adsiConnectionResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state ADSIConnectorResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
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
	if apiResp!=nil && *apiResp.ADSIConnectionResponse.Errorcode !=0{
		log.Printf("[ERROR]: Error in reading ADSI connection resource. Errorcode: %v, Message: %v", *apiResp.ADSIConnectionResponse.Errorcode, *apiResp.ADSIConnectionResponse.Msg)
		resp.Diagnostics.AddError("Reading of ADSI connection resource failed", *apiResp.ADSIConnectionResponse.Msg)
		return
	}

	state.ConnectionKey = types.Int64Value(int64(*apiResp.ADSIConnectionResponse.Connectionkey))
	state.ID = types.StringValue(fmt.Sprintf("%d", *apiResp.ADSIConnectionResponse.Connectionkey))
	state.ConnectionName = util.SafeStringDatasource(apiResp.ADSIConnectionResponse.Connectionname)
	// state.Description = util.SafeStringDatasource(apiResp.ADSIConnectionResponse.Description)
	state.DefaultSavRoles = util.SafeStringDatasource(apiResp.ADSIConnectionResponse.Defaultsavroles)
	state.ConnectionType = util.SafeStringDatasource(apiResp.ADSIConnectionResponse.Connectiontype)
	state.EmailTemplate = util.SafeStringDatasource(apiResp.ADSIConnectionResponse.Emailtemplate)
	state.ImportNestedMembership = util.SafeStringDatasource(apiResp.ADSIConnectionResponse.Connectionattributes.ImportNestedMembership)
	state.CreateAccountJson = util.SafeStringDatasource(apiResp.ADSIConnectionResponse.Connectionattributes.CREATEACCOUNTJSON)
	state.EndpointsFilter = util.SafeStringDatasource(apiResp.ADSIConnectionResponse.Connectionattributes.ENDPOINTS_FILTER)
	state.DisableAccountJson = util.SafeStringDatasource(apiResp.ADSIConnectionResponse.Connectionattributes.DISABLEACCOUNTJSON)
	state.RemoveAccessEntitlementJson = util.SafeStringDatasource(apiResp.ADSIConnectionResponse.Connectionattributes.REMOVEACCESSENTITLEMENTJSON)
	state.GroupSearchBaseDN = util.SafeStringDatasource(apiResp.ADSIConnectionResponse.Connectionattributes.GroupSearchBaseDN)
	state.StatusKeyJson = util.SafeStringDatasource(apiResp.ADSIConnectionResponse.Connectionattributes.STATUSKEYJSON)
	state.DefaultUserRole = util.SafeStringDatasource(apiResp.ADSIConnectionResponse.Connectionattributes.DEFAULT_USER_ROLE)
	state.Username = util.SafeStringDatasource(apiResp.ADSIConnectionResponse.Connectionattributes.USERNAME)
	state.UpdateServiceAccountJson = util.SafeStringDatasource(apiResp.ADSIConnectionResponse.Connectionattributes.UPDATESERVICEACCOUNTJSON)
	state.AddAccessJson = util.SafeStringDatasource(apiResp.ADSIConnectionResponse.Connectionattributes.ADDACCESSJSON)
	state.CreateServiceAccountJson = util.SafeStringDatasource(apiResp.ADSIConnectionResponse.Connectionattributes.CREATESERVICEACCOUNTJSON)
	state.AccountNameRule = util.SafeStringDatasource(apiResp.ADSIConnectionResponse.Connectionattributes.ACCOUNTNAMERULE)
	state.ConnectionUrl = util.SafeStringDatasource(apiResp.ADSIConnectionResponse.Connectionattributes.CONNECTION_URL)
	state.AccountAttribute = util.SafeStringDatasource(apiResp.ADSIConnectionResponse.Connectionattributes.ACCOUNT_ATTRIBUTE)
	state.PamConfig = util.SafeStringDatasource(apiResp.ADSIConnectionResponse.Connectionattributes.PAM_CONFIG)
	state.PageSize = util.SafeStringDatasource(apiResp.ADSIConnectionResponse.Connectionattributes.PAGE_SIZE)
	state.SearchFilter = util.SafeStringDatasource(apiResp.ADSIConnectionResponse.Connectionattributes.SEARCHFILTER)
	state.UpdateGroupJson = util.SafeStringDatasource(apiResp.ADSIConnectionResponse.Connectionattributes.UPDATEGROUPJSON)
	state.CreateGroupJson = util.SafeStringDatasource(apiResp.ADSIConnectionResponse.Connectionattributes.CREATEGROUPJSON)
	state.EntitlementAttribute = util.SafeStringDatasource(apiResp.ADSIConnectionResponse.Connectionattributes.ENTITLEMENT_ATTRIBUTE)
	state.CheckForUnique = util.SafeStringDatasource(apiResp.ADSIConnectionResponse.Connectionattributes.CHECKFORUNIQUE)
	state.RemoveServiceAccountJson = util.SafeStringDatasource(apiResp.ADSIConnectionResponse.Connectionattributes.REMOVESERVICEACCOUNTJSON)
	state.UpdateUserJson = util.SafeStringDatasource(apiResp.ADSIConnectionResponse.Connectionattributes.UPDATEUSERJSON)
	state.URL = util.SafeStringDatasource(apiResp.ADSIConnectionResponse.Connectionattributes.URL)
	state.CustomConfigJson = util.SafeStringDatasource(apiResp.ADSIConnectionResponse.Connectionattributes.CUSTOMCONFIGJSON)
	state.StatusThresholdConfig = util.SafeStringDatasource(apiResp.ADSIConnectionResponse.Connectionattributes.STATUS_THRESHOLD_CONFIG)
	state.GroupImportMapping = util.SafeStringDatasource(apiResp.ADSIConnectionResponse.Connectionattributes.GroupImportMapping)
	state.ProvisioningUrl = util.SafeStringDatasource(apiResp.ADSIConnectionResponse.Connectionattributes.PROVISIONING_URL)
	state.RemoveGroupJson = util.SafeStringDatasource(apiResp.ADSIConnectionResponse.Connectionattributes.REMOVEGROUPJSON)
	state.RemoveAccessJson = util.SafeStringDatasource(apiResp.ADSIConnectionResponse.Connectionattributes.REMOVEACCESSJSON)
	state.ResetAndChangePasswrdJson = util.SafeStringDatasource(apiResp.ADSIConnectionResponse.Connectionattributes.RESETANDCHANGEPASSWRDJSON)
	state.UserAttribute = util.SafeStringDatasource(apiResp.ADSIConnectionResponse.Connectionattributes.USER_ATTRIBUTE)
	state.AddAccessEntitlementJson = util.SafeStringDatasource(apiResp.ADSIConnectionResponse.Connectionattributes.ADDACCESSENTITLEMENTJSON)
	state.ModifyUserDataJson = util.SafeStringDatasource(apiResp.ADSIConnectionResponse.Connectionattributes.MODIFYUSERDATAJSON)
	state.EnableAccountJson = util.SafeStringDatasource(apiResp.ADSIConnectionResponse.Connectionattributes.ENABLEACCOUNTJSON)
	state.ForestList = util.SafeStringDatasource(apiResp.ADSIConnectionResponse.Connectionattributes.FORESTLIST)
	state.ObjectFilter = util.SafeStringDatasource(apiResp.ADSIConnectionResponse.Connectionattributes.OBJECTFILTER)
	state.UpdateAccountJson = util.SafeStringDatasource(apiResp.ADSIConnectionResponse.Connectionattributes.UPDATEACCOUNTJSON)
	state.RemoveAccountJson = util.SafeStringDatasource(apiResp.ADSIConnectionResponse.Connectionattributes.REMOVEACCOUNTJSON)
	apiMessage := util.SafeDeref(apiResp.ADSIConnectionResponse.Msg)
	if apiMessage == "success" {
		state.Msg = types.StringValue("Connection Successful")
	} else {
		state.Msg = types.StringValue(apiMessage)
	}
	state.ErrorCode = util.Int32PtrToTFString(apiResp.ADSIConnectionResponse.Errorcode)
	stateDiagnostics := resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(stateDiagnostics...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *adsiConnectionResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state, config ADSIConnectorResourceModel
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
	if plan.ConnectionName.ValueString() != state.ConnectionName.ValueString() {
		resp.Diagnostics.AddError("Error", "Connection name cannot be updated")
		log.Printf("[ERROR]: Connection name cannot be updated")
		return
	}
	if plan.ConnectionType.ValueString() != state.ConnectionType.ValueString() {
		resp.Diagnostics.AddError("Error", "Connection type cannot by updated")
		log.Printf("[ERROR]: Connection type cannot by updated")
		return
	}
	cfg := openapi.NewConfiguration()
	apiBaseURL := strings.TrimPrefix(strings.TrimPrefix(r.client.APIBaseURL(), "https://"), "http://")
	cfg.Host = apiBaseURL
	cfg.Scheme = "https"
	cfg.AddDefaultHeader("Authorization", "Bearer "+r.token)

	cfg.HTTPClient = http.DefaultClient
	if plan.EntitlementAttribute.IsNull() || plan.EntitlementAttribute.IsUnknown() {
		plan.EntitlementAttribute = types.StringValue("memberOf")
	}
	adsiConn := openapi.ADSIConnector{
		BaseConnector: openapi.BaseConnector{
			//required values
			Connectiontype: "ADSI",
			ConnectionName: plan.ConnectionName.ValueString(),
			//optional values
			// Description:     util.StringPointerOrEmpty(plan.Description),
			Defaultsavroles: util.StringPointerOrEmpty(plan.DefaultSavRoles),
			EmailTemplate:   util.StringPointerOrEmpty(plan.EmailTemplate),
		},
		//required values
		URL:            plan.URL.ValueString(),
		USERNAME:       config.Username.ValueString(),
		PASSWORD:       config.Password.ValueString(),
		CONNECTION_URL: plan.ConnectionUrl.ValueString(),
		FORESTLIST:     plan.ForestList.ValueString(),
		//optional values
		PROVISIONING_URL:            util.StringPointerOrEmpty(plan.ProvisioningUrl),
		DEFAULT_USER_ROLE:           util.StringPointerOrEmpty(plan.DefaultUserRole),
		UPDATEUSERJSON:              util.StringPointerOrEmpty(plan.UpdateUserJson),
		ENDPOINTS_FILTER:            util.StringPointerOrEmpty(plan.EndpointsFilter),
		SEARCHFILTER:                util.StringPointerOrEmpty(plan.SearchFilter),
		OBJECTFILTER:                util.StringPointerOrEmpty(plan.ObjectFilter),
		ACCOUNT_ATTRIBUTE:           util.StringPointerOrEmpty(plan.AccountAttribute),
		STATUS_THRESHOLD_CONFIG:     util.StringPointerOrEmpty(plan.StatusThresholdConfig),
		ENTITLEMENT_ATTRIBUTE:       util.StringPointerOrEmpty(plan.EntitlementAttribute),
		USER_ATTRIBUTE:              util.StringPointerOrEmpty(plan.UserAttribute),
		GroupSearchBaseDN:           util.StringPointerOrEmpty(plan.GroupSearchBaseDN),
		CHECKFORUNIQUE:              util.StringPointerOrEmpty(plan.CheckForUnique),
		STATUSKEYJSON:               util.StringPointerOrEmpty(plan.StatusKeyJson),
		GroupImportMapping:          util.StringPointerOrEmpty(plan.GroupImportMapping),
		ImportNestedMembership:      util.StringPointerOrEmpty(plan.ImportNestedMembership),
		PAGE_SIZE:                   util.StringPointerOrEmpty(plan.PageSize),
		ACCOUNTNAMERULE:             util.StringPointerOrEmpty(plan.AccountNameRule),
		CREATEACCOUNTJSON:           util.StringPointerOrEmpty(plan.CreateAccountJson),
		UPDATEACCOUNTJSON:           util.StringPointerOrEmpty(plan.UpdateAccountJson),
		ENABLEACCOUNTJSON:           util.StringPointerOrEmpty(plan.EnableAccountJson),
		DISABLEACCOUNTJSON:          util.StringPointerOrEmpty(plan.DisableAccountJson),
		REMOVEACCOUNTJSON:           util.StringPointerOrEmpty(plan.RemoveAccountJson),
		ADDACCESSJSON:               util.StringPointerOrEmpty(plan.AddAccessJson),
		REMOVEACCESSJSON:            util.StringPointerOrEmpty(plan.RemoveAccessJson),
		RESETANDCHANGEPASSWRDJSON:   util.StringPointerOrEmpty(plan.ResetAndChangePasswrdJson),
		CREATEGROUPJSON:             util.StringPointerOrEmpty(plan.CreateGroupJson),
		UPDATEGROUPJSON:             util.StringPointerOrEmpty(plan.UpdateGroupJson),
		REMOVEGROUPJSON:             util.StringPointerOrEmpty(plan.RemoveGroupJson),
		ADDACCESSENTITLEMENTJSON:    util.StringPointerOrEmpty(plan.AddAccessEntitlementJson),
		CUSTOMCONFIGJSON:            util.StringPointerOrEmpty(plan.CustomConfigJson),
		REMOVEACCESSENTITLEMENTJSON: util.StringPointerOrEmpty(plan.RemoveAccessEntitlementJson),
		CREATESERVICEACCOUNTJSON:    util.StringPointerOrEmpty(plan.CreateServiceAccountJson),
		UPDATESERVICEACCOUNTJSON:    util.StringPointerOrEmpty(plan.UpdateServiceAccountJson),
		REMOVESERVICEACCOUNTJSON:    util.StringPointerOrEmpty(plan.RemoveServiceAccountJson),
		PAM_CONFIG:                  util.StringPointerOrEmpty(plan.PamConfig),
		MODIFYUSERDATAJSON:          util.StringPointerOrEmpty(plan.ModifyUserDataJson),
	}
	if plan.VaultConnection.ValueString() != "" {
		adsiConn.BaseConnector.VaultConnection = util.SafeStringConnector(plan.VaultConnection.ValueString())
		adsiConn.BaseConnector.VaultConfiguration = util.SafeStringConnector(plan.VaultConfiguration.ValueString())
		adsiConn.BaseConnector.Saveinvault = util.SafeStringConnector(plan.SaveInVault.ValueString())
	} else {
		emptyStr := ""
		adsiConn.BaseConnector.VaultConnection = &emptyStr
		adsiConn.BaseConnector.VaultConfiguration = &emptyStr
		adsiConn.BaseConnector.Saveinvault = &emptyStr
	}
	adsiConnRequest := openapi.CreateOrUpdateRequest{
		ADSIConnector: &adsiConn,
	}

	// Initialize API client
	apiClient := openapi.NewAPIClient(cfg)
	apiResp, _, err := apiClient.ConnectionsAPI.CreateOrUpdate(ctx).CreateOrUpdateRequest(adsiConnRequest).Execute()
	if err != nil {
		log.Printf("Problem with the update function")
		resp.Diagnostics.AddError("API Update Failed", fmt.Sprintf("Error: %v", err))
		return
	}
	if apiResp!=nil && *apiResp.ErrorCode !="0"{
		log.Printf("[ERROR]: Error in updating ADSI connection. Errorcode: %v, Message: %v", *apiResp.ErrorCode, *apiResp.Msg)
		resp.Diagnostics.AddError("Updating of ADSI connection failed", *apiResp.Msg)
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
	if getResp!=nil && *getResp.ADSIConnectionResponse.Errorcode !=0{
		log.Printf("[ERROR]: Error in reading ADSI connection after updation. Errorcode: %v, Message: %v", *getResp.ADSIConnectionResponse.Errorcode, *getResp.ADSIConnectionResponse.Msg)
		resp.Diagnostics.AddError("Reading after updation of ADSI connection failed", *getResp.ADSIConnectionResponse.Msg)
		return
	}

	plan.ConnectionKey = types.Int64Value(int64(*getResp.ADSIConnectionResponse.Connectionkey))
	plan.ID = types.StringValue(fmt.Sprintf("%d", *getResp.ADSIConnectionResponse.Connectionkey))
	plan.ConnectionName = util.SafeStringDatasource(getResp.ADSIConnectionResponse.Connectionname)
	// plan.Description = util.SafeStringDatasource(getResp.ADSIConnectionResponse.Description)
	plan.DefaultSavRoles = util.SafeStringDatasource(getResp.ADSIConnectionResponse.Defaultsavroles)
	plan.ConnectionType = util.SafeStringDatasource(getResp.ADSIConnectionResponse.Connectiontype)
	plan.EmailTemplate = util.SafeStringDatasource(getResp.ADSIConnectionResponse.Emailtemplate)
	plan.ImportNestedMembership = util.SafeStringDatasource(getResp.ADSIConnectionResponse.Connectionattributes.ImportNestedMembership)
	plan.CreateAccountJson = util.SafeStringDatasource(getResp.ADSIConnectionResponse.Connectionattributes.CREATEACCOUNTJSON)
	plan.EndpointsFilter = util.SafeStringDatasource(getResp.ADSIConnectionResponse.Connectionattributes.ENDPOINTS_FILTER)
	plan.DisableAccountJson = util.SafeStringDatasource(getResp.ADSIConnectionResponse.Connectionattributes.DISABLEACCOUNTJSON)
	plan.RemoveAccessEntitlementJson = util.SafeStringDatasource(getResp.ADSIConnectionResponse.Connectionattributes.REMOVEACCESSENTITLEMENTJSON)
	plan.GroupSearchBaseDN = util.SafeStringDatasource(getResp.ADSIConnectionResponse.Connectionattributes.GroupSearchBaseDN)
	plan.StatusKeyJson = util.SafeStringDatasource(getResp.ADSIConnectionResponse.Connectionattributes.STATUSKEYJSON)
	plan.DefaultUserRole = util.SafeStringDatasource(getResp.ADSIConnectionResponse.Connectionattributes.DEFAULT_USER_ROLE)
	plan.Username = util.SafeStringDatasource(getResp.ADSIConnectionResponse.Connectionattributes.USERNAME)
	plan.UpdateServiceAccountJson = util.SafeStringDatasource(getResp.ADSIConnectionResponse.Connectionattributes.UPDATESERVICEACCOUNTJSON)
	plan.AddAccessJson = util.SafeStringDatasource(getResp.ADSIConnectionResponse.Connectionattributes.ADDACCESSJSON)
	plan.CreateServiceAccountJson = util.SafeStringDatasource(getResp.ADSIConnectionResponse.Connectionattributes.CREATESERVICEACCOUNTJSON)
	plan.AccountNameRule = util.SafeStringDatasource(getResp.ADSIConnectionResponse.Connectionattributes.ACCOUNTNAMERULE)
	plan.ConnectionUrl = util.SafeStringDatasource(getResp.ADSIConnectionResponse.Connectionattributes.CONNECTION_URL)
	plan.AccountAttribute = util.SafeStringDatasource(getResp.ADSIConnectionResponse.Connectionattributes.ACCOUNT_ATTRIBUTE)
	plan.PamConfig = util.SafeStringDatasource(getResp.ADSIConnectionResponse.Connectionattributes.PAM_CONFIG)
	plan.PageSize = util.SafeStringDatasource(getResp.ADSIConnectionResponse.Connectionattributes.PAGE_SIZE)
	plan.SearchFilter = util.SafeStringDatasource(getResp.ADSIConnectionResponse.Connectionattributes.SEARCHFILTER)
	plan.UpdateGroupJson = util.SafeStringDatasource(getResp.ADSIConnectionResponse.Connectionattributes.UPDATEGROUPJSON)
	plan.CreateGroupJson = util.SafeStringDatasource(getResp.ADSIConnectionResponse.Connectionattributes.CREATEGROUPJSON)
	plan.EntitlementAttribute = util.SafeStringDatasource(getResp.ADSIConnectionResponse.Connectionattributes.ENTITLEMENT_ATTRIBUTE)
	plan.CheckForUnique = util.SafeStringDatasource(getResp.ADSIConnectionResponse.Connectionattributes.CHECKFORUNIQUE)
	plan.RemoveServiceAccountJson = util.SafeStringDatasource(getResp.ADSIConnectionResponse.Connectionattributes.REMOVESERVICEACCOUNTJSON)
	plan.UpdateUserJson = util.SafeStringDatasource(getResp.ADSIConnectionResponse.Connectionattributes.UPDATEUSERJSON)
	plan.URL = util.SafeStringDatasource(getResp.ADSIConnectionResponse.Connectionattributes.URL)
	plan.CustomConfigJson = util.SafeStringDatasource(getResp.ADSIConnectionResponse.Connectionattributes.CUSTOMCONFIGJSON)
	plan.StatusThresholdConfig = util.SafeStringDatasource(getResp.ADSIConnectionResponse.Connectionattributes.STATUS_THRESHOLD_CONFIG)
	plan.GroupImportMapping = util.SafeStringDatasource(getResp.ADSIConnectionResponse.Connectionattributes.GroupImportMapping)
	plan.ProvisioningUrl = util.SafeStringDatasource(getResp.ADSIConnectionResponse.Connectionattributes.PROVISIONING_URL)
	plan.RemoveGroupJson = util.SafeStringDatasource(getResp.ADSIConnectionResponse.Connectionattributes.REMOVEGROUPJSON)
	plan.RemoveAccessJson = util.SafeStringDatasource(getResp.ADSIConnectionResponse.Connectionattributes.REMOVEACCESSJSON)
	plan.ResetAndChangePasswrdJson = util.SafeStringDatasource(getResp.ADSIConnectionResponse.Connectionattributes.RESETANDCHANGEPASSWRDJSON)
	plan.UserAttribute = util.SafeStringDatasource(getResp.ADSIConnectionResponse.Connectionattributes.USER_ATTRIBUTE)
	plan.AddAccessEntitlementJson = util.SafeStringDatasource(getResp.ADSIConnectionResponse.Connectionattributes.ADDACCESSENTITLEMENTJSON)
	plan.ModifyUserDataJson = util.SafeStringDatasource(getResp.ADSIConnectionResponse.Connectionattributes.MODIFYUSERDATAJSON)
	plan.EnableAccountJson = util.SafeStringDatasource(getResp.ADSIConnectionResponse.Connectionattributes.ENABLEACCOUNTJSON)
	plan.ForestList = util.SafeStringDatasource(getResp.ADSIConnectionResponse.Connectionattributes.FORESTLIST)
	plan.ObjectFilter = util.SafeStringDatasource(getResp.ADSIConnectionResponse.Connectionattributes.OBJECTFILTER)
	plan.UpdateAccountJson = util.SafeStringDatasource(getResp.ADSIConnectionResponse.Connectionattributes.UPDATEACCOUNTJSON)
	plan.RemoveAccountJson = util.SafeStringDatasource(getResp.ADSIConnectionResponse.Connectionattributes.REMOVEACCOUNTJSON)
	apiMessage := util.SafeDeref(getResp.ADSIConnectionResponse.Msg)
	if apiMessage == "success" {
		plan.Msg = types.StringValue("Connection Successful")
	} else {
		plan.Msg = types.StringValue(apiMessage)
	}
	plan.ErrorCode = util.Int32PtrToTFString(getResp.ADSIConnectionResponse.Errorcode)
	stateUpdateDiagnostics := resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(stateUpdateDiagnostics...)
}

func (r *adsiConnectionResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
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
func (r *adsiConnectionResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	// Retrieve import ID and save to id attribute
	resource.ImportStatePassthroughID(ctx, path.Root("connection_name"), req, resp)
}
