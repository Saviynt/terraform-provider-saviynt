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

// saviynt_workday_connection_resource manages Workday connectors in the Saviynt Security Manager.
// The resource implements the full Terraform lifecycle:
//   - Create: provisions a new Workday connector using the supplied configuration.
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
var _ resource.Resource = &workdayConnectionResource{}
var _ resource.ResourceWithImportState = &workdayConnectionResource{}

type WorkdayConnectorResourceModel struct {
	BaseConnectorResourceModel
	ID                     types.String `tfsdk:"id"`
	UsersLastImportTime    types.String `tfsdk:"users_last_import_time"`
	AccountsLastImportTime types.String `tfsdk:"accounts_last_import_time"`
	AccessLastImportTime   types.String `tfsdk:"access_last_import_time"`
	BaseURL                types.String `tfsdk:"base_url"`
	APIVersion             types.String `tfsdk:"api_version"`
	TenantName             types.String `tfsdk:"tenant_name"`
	ReportOwner            types.String `tfsdk:"report_owner"`
	UseOAuth               types.String `tfsdk:"use_oauth"`
	IncludeReferenceDesc   types.String `tfsdk:"include_reference_descriptors"`
	UseEnhancedOrgRole     types.String `tfsdk:"use_enhanced_orgrole"`
	UseX509AuthForSOAP     types.String `tfsdk:"use_x509auth_for_soap"`
	X509Key                types.String `tfsdk:"x509_key"`
	X509Cert               types.String `tfsdk:"x509_cert"`
	Username               types.String `tfsdk:"username"`
	Password               types.String `tfsdk:"password"`
	ClientID               types.String `tfsdk:"client_id"`
	ClientSecret           types.String `tfsdk:"client_secret"`
	RefreshToken           types.String `tfsdk:"refresh_token"`
	PageSize               types.String `tfsdk:"page_size"`
	UserImportPayload      types.String `tfsdk:"user_import_payload"`
	UserImportMapping      types.String `tfsdk:"user_import_mapping"`
	AccountImportPayload   types.String `tfsdk:"account_import_payload"`
	AccountImportMapping   types.String `tfsdk:"account_import_mapping"`
	AccessImportList       types.String `tfsdk:"access_import_list"`
	RAASMappingJSON        types.String `tfsdk:"raas_mapping_json"`
	AccessImportMapping    types.String `tfsdk:"access_import_mapping"`
	// OrgRoleImportPayload   types.String `tfsdk:"orgrole_import_payload"`
	StatusKeyJSON         types.String `tfsdk:"status_key_json"`
	UserAttributeJSON     types.String `tfsdk:"userattributejson"`
	CustomConfig          types.String `tfsdk:"custom_config"`
	PAMConfig             types.String `tfsdk:"pam_config"`
	ModifyUserDataJSON    types.String `tfsdk:"modify_user_data_json"`
	StatusThresholdConfig types.String `tfsdk:"status_threshold_config"`
	CreateAccountPayload  types.String `tfsdk:"create_account_payload"`
	UpdateAccountPayload  types.String `tfsdk:"update_account_payload"`
	UpdateUserPayload     types.String `tfsdk:"update_user_payload"`
	AssignOrgRolePayload  types.String `tfsdk:"assign_orgrole_payload"`
	RemoveOrgRolePayload  types.String `tfsdk:"remove_orgrole_payload"`
}

// workdayConnectionResource implements the resource.Resource interface.
type workdayConnectionResource struct {
	client *s.Client
	token  string
}

// NewTestConnectionResource returns a new instance of testConnectionResource.
func NewWorkdayTestConnectionResource() resource.Resource {
	return &workdayConnectionResource{}
}

func (r *workdayConnectionResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = "saviynt_workday_connection_resource"
}

func WorkdayConnectorResourceSchema() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:    true,
			Description: "Resource ID.",
		},
		"users_last_import_time": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "Property for USERS_LAST_IMPORT_TIME.",
		},
		"accounts_last_import_time": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "Property for ACCOUNTS_LAST_IMPORT_TIME.",
		},
		"access_last_import_time": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "Property for ACCESS_LAST_IMPORT_TIME.",
		},
		"base_url": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "Base URL of the Workday tenant instance.",
		},
		"api_version": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "Version of the SOAP API used for the connection.",
		},
		"tenant_name": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "The name of your tenant.",
		},
		"report_owner": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "Account name of the report owner used to build default RaaS URLs.",
		},
		"use_oauth": schema.StringAttribute{
			Required:    true,
			Description: "Whether to use OAuth authentication.",
		},
		"include_reference_descriptors": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "Include descriptor attribute in response if set to TRUE.",
		},
		"use_enhanced_orgrole": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "Set TRUE to utilize enhanced Organizational Role setup.",
		},
		"use_x509auth_for_soap": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "Set TRUE to use certificate-based authentication for SOAP.",
		},
		"x509_key": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Sensitive:   true,
			Description: "Private key for x509-based SOAP authentication.",
		},
		"x509_cert": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Sensitive:   true,
			Description: "Certificate for x509-based SOAP authentication.",
		},
		"username": schema.StringAttribute{
			Optional:    true,
			WriteOnly:   true,
			Description: "Username for SOAP authentication.",
		},
		"password": schema.StringAttribute{
			Optional:    true,
			WriteOnly:   true,
			Description: "Password for SOAP authentication.",
		},
		"client_id": schema.StringAttribute{
			Optional:    true,
			WriteOnly:   true,
			Description: "OAuth client ID.",
		},
		"client_secret": schema.StringAttribute{
			Optional:    true,
			WriteOnly:   true,
			Description: "OAuth client secret.",
		},
		"refresh_token": schema.StringAttribute{
			Optional:    true,
			WriteOnly:   true,
			Description: "OAuth refresh token.",
		},
		"page_size": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "Number of objects to return per page during import.",
		},
		"user_import_payload": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "Request payload for importing users.",
		},
		"user_import_mapping": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "Mapping configuration for user import.",
		},
		"account_import_payload": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "Request payload for importing accounts.",
		},
		"account_import_mapping": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "Mapping configuration for account import.",
		},
		"access_import_list": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "Comma-separated list of access types to import.",
		},
		"raas_mapping_json": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "Overrides default report mapping for RaaS.",
		},
		"access_import_mapping": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "Additional access attribute mapping for Workday access objects.",
		},
		// "orgrole_import_payload": schema.StringAttribute{
		// 	Optional:    true,
		// 	Computed:    true,
		// 	Description: "Custom SOAP body for organization role import.",
		// },
		"status_key_json": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "Mapping of user status.",
		},
		"userattributejson": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "Specifies which job-related attributes are stored as user attributes.",
		},
		"custom_config": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "Custom configuration for Workday connector.",
		},
		"pam_config": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "Privileged Access Management configuration.",
		},
		"modify_user_data_json": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "Payload for modifying user data.",
		},
		"status_threshold_config": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "Config for reading and importing status of account and entitlement.",
		},
		"create_account_payload": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "Payload for creating an account.",
		},
		"update_account_payload": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "Payload for updating an account.",
		},
		"update_user_payload": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "Payload for updating a user.",
		},
		"assign_orgrole_payload": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "Payload for assigning org role.",
		},
		"remove_orgrole_payload": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "Payload for removing org role.",
		},
	}
}

func (r *workdayConnectionResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: util.WorkdayConnDescription,
		Attributes:  connectionsutil.MergeResourceAttributes(BaseConnectorResourceSchema(), WorkdayConnectorResourceSchema()),
	}
}

func (r *workdayConnectionResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *workdayConnectionResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan, config WorkdayConnectorResourceModel
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
	if existingResource != nil && existingResource.WorkdayConnectionResponse != nil && existingResource.WorkdayConnectionResponse.Errorcode != nil && *existingResource.WorkdayConnectionResponse.Errorcode == 0 {
		log.Printf("[ERROR] Connection name already exists. Please import or use a different name")
		resp.Diagnostics.AddError("API Create Failed", "Connection name already exists. Please import or use a different name")
		return
	}

	workdayConn := openapi.WorkdayConnector{
		BaseConnector: openapi.BaseConnector{
			//required fields
			Connectiontype: "Workday",
			ConnectionName: plan.ConnectionName.ValueString(),
			//optional fields
			// Description:        util.StringPointerOrEmpty(plan.Description),
			Defaultsavroles:    util.StringPointerOrEmpty(plan.DefaultSavRoles),
			EmailTemplate:      util.StringPointerOrEmpty(plan.EmailTemplate),
			VaultConnection:    util.SafeStringConnector(plan.VaultConnection.ValueString()),
			VaultConfiguration: util.SafeStringConnector(plan.VaultConfiguration.ValueString()),
			Saveinvault:        util.SafeStringConnector(plan.SaveInVault.ValueString()),
		},
		//required fields
		USE_OAUTH: plan.UseOAuth.ValueString(),
		//optional fields
		USERS_LAST_IMPORT_TIME:        util.StringPointerOrEmpty(plan.UsersLastImportTime),
		ACCOUNTS_LAST_IMPORT_TIME:     util.StringPointerOrEmpty(plan.AccountsLastImportTime),
		ACCESS_LAST_IMPORT_TIME:       util.StringPointerOrEmpty(plan.AccessLastImportTime),
		BASE_URL:                      util.StringPointerOrEmpty(plan.BaseURL),
		API_VERSION:                   util.StringPointerOrEmpty(plan.APIVersion),
		TENANT_NAME:                   util.StringPointerOrEmpty(plan.TenantName),
		REPORT_OWNER:                  util.StringPointerOrEmpty(plan.ReportOwner),
		INCLUDE_REFERENCE_DESCRIPTORS: util.StringPointerOrEmpty(plan.IncludeReferenceDesc),
		USE_ENHANCED_ORGROLE:          util.StringPointerOrEmpty(plan.UseEnhancedOrgRole),
		USEX509AUTHFORSOAP:            util.StringPointerOrEmpty(plan.UseX509AuthForSOAP),
		X509KEY:                       util.StringPointerOrEmpty(plan.X509Key),
		X509CERT:                      util.StringPointerOrEmpty(plan.X509Cert),
		USERNAME:                      util.StringPointerOrEmpty(config.Username),
		PASSWORD:                      util.StringPointerOrEmpty(config.Password),
		CLIENT_ID:                     util.StringPointerOrEmpty(config.ClientID),
		CLIENT_SECRET:                 util.StringPointerOrEmpty(config.ClientSecret),
		REFRESH_TOKEN:                 util.StringPointerOrEmpty(config.RefreshToken),
		PAGE_SIZE:                     util.StringPointerOrEmpty(plan.PageSize),
		USER_IMPORT_PAYLOAD:           util.StringPointerOrEmpty(plan.UserImportPayload),
		USER_IMPORT_MAPPING:           util.StringPointerOrEmpty(plan.UserImportMapping),
		ACCOUNT_IMPORT_PAYLOAD:        util.StringPointerOrEmpty(plan.AccountImportPayload),
		ACCOUNT_IMPORT_MAPPING:        util.StringPointerOrEmpty(plan.AccountImportMapping),
		ACCESS_IMPORT_LIST:            util.StringPointerOrEmpty(plan.AccessImportList),
		RAAS_MAPPING_JSON:             util.StringPointerOrEmpty(plan.RAASMappingJSON),
		ACCESS_IMPORT_MAPPING:         util.StringPointerOrEmpty(plan.AccessImportMapping),
		// ORGROLE_IMPORT_PAYLOAD:        util.StringPointerOrEmpty(plan.OrgRoleImportPayload),
		STATUS_KEY_JSON:         util.StringPointerOrEmpty(plan.StatusKeyJSON),
		USERATTRIBUTEJSON:       util.StringPointerOrEmpty(plan.UserAttributeJSON),
		CUSTOM_CONFIG:           util.StringPointerOrEmpty(plan.CustomConfig),
		PAM_CONFIG:              util.StringPointerOrEmpty(plan.PAMConfig),
		MODIFYUSERDATAJSON:      util.StringPointerOrEmpty(plan.ModifyUserDataJSON),
		STATUS_THRESHOLD_CONFIG: util.StringPointerOrEmpty(plan.StatusThresholdConfig),
		CREATE_ACCOUNT_PAYLOAD:  util.StringPointerOrEmpty(plan.CreateAccountPayload),
		UPDATE_ACCOUNT_PAYLOAD:  util.StringPointerOrEmpty(plan.UpdateAccountPayload),
		UPDATE_USER_PAYLOAD:     util.StringPointerOrEmpty(plan.UpdateUserPayload),
		ASSIGN_ORGROLE_PAYLOAD:  util.StringPointerOrEmpty(plan.AssignOrgRolePayload),
		REMOVE_ORGROLE_PAYLOAD:  util.StringPointerOrEmpty(plan.RemoveOrgRolePayload),
	}
	workdayConnRequest := openapi.CreateOrUpdateRequest{
		WorkdayConnector: &workdayConn,
	}

	apiResp, _, err := apiClient.ConnectionsAPI.CreateOrUpdate(ctx).CreateOrUpdateRequest(workdayConnRequest).Execute()
	if err != nil {
		log.Printf("[ERROR] Failed to create API resource. Error: %v", err)
		resp.Diagnostics.AddError("API Create Failed", fmt.Sprintf("Error: %v", err))
		return
	}
	if apiResp != nil && *apiResp.ErrorCode != "0" {
		log.Printf("[ERROR]: Error in creating Workday connection resource. Errorcode: %v, Message: %v", *apiResp.ErrorCode, *apiResp.Msg)
		resp.Diagnostics.AddError("Creation of Workday connection failed", *apiResp.Msg)
		return
	}

	if plan.UseEnhancedOrgRole.IsNull() || plan.UseEnhancedOrgRole.ValueString() == "" {
		plan.UseEnhancedOrgRole = types.StringValue("TRUE")
	}
	plan.ID = types.StringValue(fmt.Sprintf("%d", *apiResp.ConnectionKey))
	plan.ConnectionType = types.StringValue("Workday")
	plan.ConnectionKey = types.Int64Value(int64(*apiResp.ConnectionKey))
	// plan.Description = util.SafeStringDatasource(plan.Description.ValueStringPointer())
	plan.DefaultSavRoles = util.SafeStringDatasource(plan.DefaultSavRoles.ValueStringPointer())
	plan.EmailTemplate = util.SafeStringDatasource(plan.EmailTemplate.ValueStringPointer())
	plan.UsersLastImportTime = util.SafeStringDatasource(plan.UsersLastImportTime.ValueStringPointer())
	plan.AccountsLastImportTime = util.SafeStringDatasource(plan.AccountsLastImportTime.ValueStringPointer())
	plan.AccessLastImportTime = util.SafeStringDatasource(plan.AccessLastImportTime.ValueStringPointer())
	plan.BaseURL = util.SafeStringDatasource(plan.BaseURL.ValueStringPointer())
	plan.APIVersion = util.SafeStringDatasource(plan.APIVersion.ValueStringPointer())
	plan.TenantName = util.SafeStringDatasource(plan.TenantName.ValueStringPointer())
	plan.ReportOwner = util.SafeStringDatasource(plan.ReportOwner.ValueStringPointer())
	plan.IncludeReferenceDesc = util.SafeStringDatasource(plan.IncludeReferenceDesc.ValueStringPointer())
	plan.UseEnhancedOrgRole = util.SafeStringDatasource(plan.UseEnhancedOrgRole.ValueStringPointer())
	plan.UseX509AuthForSOAP = util.SafeStringDatasource(plan.UseX509AuthForSOAP.ValueStringPointer())
	plan.X509Key = util.SafeStringDatasource(plan.X509Key.ValueStringPointer())
	plan.X509Cert = util.SafeStringDatasource(plan.X509Cert.ValueStringPointer())
	plan.Username = util.SafeStringDatasource(plan.Username.ValueStringPointer())
	plan.ClientID = util.SafeStringDatasource(plan.ClientID.ValueStringPointer())
	plan.PageSize = util.SafeStringDatasource(plan.PageSize.ValueStringPointer())
	plan.UserImportPayload = util.SafeStringDatasource(plan.UserImportPayload.ValueStringPointer())
	plan.UserImportMapping = util.SafeStringDatasource(plan.UserImportMapping.ValueStringPointer())
	plan.AccountImportPayload = util.SafeStringDatasource(plan.AccountImportPayload.ValueStringPointer())
	plan.AccountImportMapping = util.SafeStringDatasource(plan.AccountImportMapping.ValueStringPointer())
	plan.AccessImportList = util.SafeStringDatasource(plan.AccessImportList.ValueStringPointer())
	plan.RAASMappingJSON = util.SafeStringDatasource(plan.RAASMappingJSON.ValueStringPointer())
	plan.AccessImportMapping = util.SafeStringDatasource(plan.AccessImportMapping.ValueStringPointer())
	// plan.OrgRoleImportPayload = util.SafeStringDatasource(plan.OrgRoleImportPayload.ValueStringPointer())
	plan.StatusKeyJSON = util.SafeStringDatasource(plan.StatusKeyJSON.ValueStringPointer())
	plan.UserAttributeJSON = util.SafeStringDatasource(plan.UserAttributeJSON.ValueStringPointer())
	plan.CustomConfig = util.SafeStringDatasource(plan.CustomConfig.ValueStringPointer())
	plan.PAMConfig = util.SafeStringDatasource(plan.PAMConfig.ValueStringPointer())
	plan.ModifyUserDataJSON = util.SafeStringDatasource(plan.ModifyUserDataJSON.ValueStringPointer())
	plan.StatusThresholdConfig = util.SafeStringDatasource(plan.StatusThresholdConfig.ValueStringPointer())
	plan.CreateAccountPayload = util.SafeStringDatasource(plan.CreateAccountPayload.ValueStringPointer())
	plan.UpdateAccountPayload = util.SafeStringDatasource(plan.UpdateAccountPayload.ValueStringPointer())
	plan.UpdateUserPayload = util.SafeStringDatasource(plan.UpdateUserPayload.ValueStringPointer())
	plan.AssignOrgRolePayload = util.SafeStringDatasource(plan.AssignOrgRolePayload.ValueStringPointer())
	plan.RemoveOrgRolePayload = util.SafeStringDatasource(plan.RemoveOrgRolePayload.ValueStringPointer())
	plan.Msg = types.StringValue(util.SafeDeref(apiResp.Msg))
	plan.ErrorCode = types.StringValue(util.SafeDeref(apiResp.ErrorCode))
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *workdayConnectionResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state WorkdayConnectorResourceModel

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
		resp.Diagnostics.AddError("API Read Failed", fmt.Sprintf("Error: %v", err))
		return
	}
	if apiResp != nil && *apiResp.WorkdayConnectionResponse.Errorcode != 0 {
		log.Printf("[ERROR]: Error in reading Workday connection resource. Errorcode: %v, Message: %v", *apiResp.WorkdayConnectionResponse.Errorcode, *apiResp.WorkdayConnectionResponse.Msg)
		resp.Diagnostics.AddError("Reading Workday connection resource failed", *apiResp.WorkdayConnectionResponse.Msg)
		return
	}

	state.ConnectionKey = types.Int64Value(int64(*apiResp.WorkdayConnectionResponse.Connectionkey))
	state.ID = types.StringValue(fmt.Sprintf("%d", *apiResp.WorkdayConnectionResponse.Connectionkey))
	state.ConnectionName = util.SafeStringDatasource(apiResp.WorkdayConnectionResponse.Connectionname)
	// state.Description = util.SafeStringDatasource(apiResp.WorkdayConnectionResponse.Description)
	state.DefaultSavRoles = util.SafeStringDatasource(apiResp.WorkdayConnectionResponse.Defaultsavroles)
	state.ConnectionType = util.SafeStringDatasource(apiResp.WorkdayConnectionResponse.Connectiontype)
	state.EmailTemplate = util.SafeStringDatasource(apiResp.WorkdayConnectionResponse.Emailtemplate)
	state.UseOAuth = util.SafeStringDatasource(apiResp.WorkdayConnectionResponse.Connectionattributes.USE_OAUTH)
	state.UserImportMapping = util.SafeStringDatasource(apiResp.WorkdayConnectionResponse.Connectionattributes.USER_IMPORT_MAPPING)
	state.AccountsLastImportTime = util.SafeStringDatasource(apiResp.WorkdayConnectionResponse.Connectionattributes.ACCOUNTS_LAST_IMPORT_TIME)
	state.StatusKeyJSON = util.SafeStringDatasource(apiResp.WorkdayConnectionResponse.Connectionattributes.STATUS_KEY_JSON)
	state.ConnectionType = util.SafeStringDatasource(apiResp.WorkdayConnectionResponse.Connectiontype)
	state.RAASMappingJSON = util.SafeStringDatasource(apiResp.WorkdayConnectionResponse.Connectionattributes.RAAS_MAPPING_JSON)
	state.AccountImportPayload = util.SafeStringDatasource(apiResp.WorkdayConnectionResponse.Connectionattributes.ACCOUNT_IMPORT_PAYLOAD)
	state.UpdateAccountPayload = util.SafeStringDatasource(apiResp.WorkdayConnectionResponse.Connectionattributes.UPDATE_ACCOUNT_PAYLOAD)
	state.ClientID = util.SafeStringDatasource(apiResp.WorkdayConnectionResponse.Connectionattributes.CLIENT_ID)
	state.StatusThresholdConfig = util.SafeStringDatasource(apiResp.WorkdayConnectionResponse.Connectionattributes.STATUS_THRESHOLD_CONFIG)
	state.Username = util.SafeStringDatasource(apiResp.WorkdayConnectionResponse.Connectionattributes.USERNAME)
	state.AccessImportList = util.SafeStringDatasource(apiResp.WorkdayConnectionResponse.Connectionattributes.ACCESS_IMPORT_LIST)
	state.AccountImportMapping = util.SafeStringDatasource(apiResp.WorkdayConnectionResponse.Connectionattributes.ACCOUNT_IMPORT_MAPPING)
	// state.OrgRoleImportPayload = util.SafeStringDatasource(apiResp.WorkdayConnectionResponse.Connectionattributes.ORGROLE_IMPORT_PAYLOAD)
	state.AssignOrgRolePayload = util.SafeStringDatasource(apiResp.WorkdayConnectionResponse.Connectionattributes.ASSIGN_ORGROLE_PAYLOAD)
	state.AccessImportMapping = util.SafeStringDatasource(apiResp.WorkdayConnectionResponse.Connectionattributes.ACCESS_IMPORT_MAPPING)
	state.APIVersion = util.SafeStringDatasource(apiResp.WorkdayConnectionResponse.Connectionattributes.API_VERSION)
	state.RemoveOrgRolePayload = util.SafeStringDatasource(apiResp.WorkdayConnectionResponse.Connectionattributes.REMOVE_ORGROLE_PAYLOAD)
	state.IncludeReferenceDesc = util.SafeStringDatasource(apiResp.WorkdayConnectionResponse.Connectionattributes.INCLUDE_REFERENCE_DESCRIPTORS)
	state.ModifyUserDataJSON = util.SafeStringDatasource(apiResp.WorkdayConnectionResponse.Connectionattributes.MODIFYUSERDATAJSON)
	state.UseX509AuthForSOAP = util.SafeStringDatasource(apiResp.WorkdayConnectionResponse.Connectionattributes.USEX509AUTHFORSOAP)
	state.ReportOwner = util.SafeStringDatasource(apiResp.WorkdayConnectionResponse.Connectionattributes.REPORT_OWNER)
	state.X509Key = util.SafeStringDatasource(apiResp.WorkdayConnectionResponse.Connectionattributes.X509KEY)
	state.CustomConfig = util.SafeStringDatasource(apiResp.WorkdayConnectionResponse.Connectionattributes.CUSTOM_CONFIG)
	state.UserAttributeJSON = util.SafeStringDatasource(apiResp.WorkdayConnectionResponse.Connectionattributes.USERATTRIBUTEJSON)
	state.X509Cert = util.SafeStringDatasource(apiResp.WorkdayConnectionResponse.Connectionattributes.X509CERT)
	state.UserImportPayload = util.SafeStringDatasource(apiResp.WorkdayConnectionResponse.Connectionattributes.USER_IMPORT_PAYLOAD)
	state.PAMConfig = util.SafeStringDatasource(apiResp.WorkdayConnectionResponse.Connectionattributes.PAM_CONFIG)
	state.AccessLastImportTime = util.SafeStringDatasource(apiResp.WorkdayConnectionResponse.Connectionattributes.ACCESS_LAST_IMPORT_TIME)
	state.UsersLastImportTime = util.SafeStringDatasource(apiResp.WorkdayConnectionResponse.Connectionattributes.USERS_LAST_IMPORT_TIME)
	state.UpdateUserPayload = util.SafeStringDatasource(apiResp.WorkdayConnectionResponse.Connectionattributes.UPDATE_USER_PAYLOAD)
	state.PageSize = util.SafeStringDatasource(apiResp.WorkdayConnectionResponse.Connectionattributes.PAGE_SIZE)
	state.TenantName = util.SafeStringDatasource(apiResp.WorkdayConnectionResponse.Connectionattributes.TENANT_NAME)
	state.UseEnhancedOrgRole = util.SafeStringDatasource(apiResp.WorkdayConnectionResponse.Connectionattributes.USE_ENHANCED_ORGROLE)
	state.CreateAccountPayload = util.SafeStringDatasource(apiResp.WorkdayConnectionResponse.Connectionattributes.CREATE_ACCOUNT_PAYLOAD)
	state.BaseURL = util.SafeStringDatasource(apiResp.WorkdayConnectionResponse.Connectionattributes.BASE_URL)
	apiMessage := util.SafeDeref(apiResp.WorkdayConnectionResponse.Msg)
	if apiMessage == "success" {
		state.Msg = types.StringValue("Connection Successful")
	} else {
		state.Msg = types.StringValue(apiMessage)
	}
	state.ErrorCode = util.Int32PtrToTFString(apiResp.WorkdayConnectionResponse.Errorcode)
	stateDiagnostics := resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(stateDiagnostics...)
	if resp.Diagnostics.HasError() {
		return
	}
}
func (r *workdayConnectionResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state, config WorkdayConnectorResourceModel
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
	if plan.ConnectionType.ValueString() != state.ConnectionType.ValueString() {
		resp.Diagnostics.AddError("Error", "Connection type cannot by updated")
		log.Printf("[ERROR]: Connection type cannot by updated")
		return
	}

	cfg.HTTPClient = http.DefaultClient
	workdayConn := openapi.WorkdayConnector{
		BaseConnector: openapi.BaseConnector{
			//required fields
			Connectiontype: "Workday",
			ConnectionName: plan.ConnectionName.ValueString(),
			//optional fields
			// Description:        util.StringPointerOrEmpty(plan.Description),
			Defaultsavroles:    util.StringPointerOrEmpty(plan.DefaultSavRoles),
			EmailTemplate:      util.StringPointerOrEmpty(plan.EmailTemplate),
			VaultConnection:    util.SafeStringConnector(plan.VaultConnection.ValueString()),
			VaultConfiguration: util.SafeStringConnector(plan.VaultConfiguration.ValueString()),
			Saveinvault:        util.SafeStringConnector(plan.SaveInVault.ValueString()),
		},
		//required fields
		USE_OAUTH: plan.UseOAuth.ValueString(),
		//optional fields
		USERS_LAST_IMPORT_TIME:        util.StringPointerOrEmpty(plan.UsersLastImportTime),
		ACCOUNTS_LAST_IMPORT_TIME:     util.StringPointerOrEmpty(plan.AccountsLastImportTime),
		ACCESS_LAST_IMPORT_TIME:       util.StringPointerOrEmpty(plan.AccessLastImportTime),
		BASE_URL:                      util.StringPointerOrEmpty(plan.BaseURL),
		API_VERSION:                   util.StringPointerOrEmpty(plan.APIVersion),
		TENANT_NAME:                   util.StringPointerOrEmpty(plan.TenantName),
		REPORT_OWNER:                  util.StringPointerOrEmpty(plan.ReportOwner),
		INCLUDE_REFERENCE_DESCRIPTORS: util.StringPointerOrEmpty(plan.IncludeReferenceDesc),
		USE_ENHANCED_ORGROLE:          util.StringPointerOrEmpty(plan.UseEnhancedOrgRole),
		USEX509AUTHFORSOAP:            util.StringPointerOrEmpty(plan.UseX509AuthForSOAP),
		X509KEY:                       util.StringPointerOrEmpty(plan.X509Key),
		X509CERT:                      util.StringPointerOrEmpty(plan.X509Cert),
		USERNAME:                      util.StringPointerOrEmpty(config.Username),
		PASSWORD:                      util.StringPointerOrEmpty(config.Password),
		CLIENT_ID:                     util.StringPointerOrEmpty(config.ClientID),
		CLIENT_SECRET:                 util.StringPointerOrEmpty(config.ClientSecret),
		REFRESH_TOKEN:                 util.StringPointerOrEmpty(config.RefreshToken),
		PAGE_SIZE:                     util.StringPointerOrEmpty(plan.PageSize),
		USER_IMPORT_PAYLOAD:           util.StringPointerOrEmpty(plan.UserImportPayload),
		USER_IMPORT_MAPPING:           util.StringPointerOrEmpty(plan.UserImportMapping),
		ACCOUNT_IMPORT_PAYLOAD:        util.StringPointerOrEmpty(plan.AccountImportPayload),
		ACCOUNT_IMPORT_MAPPING:        util.StringPointerOrEmpty(plan.AccountImportMapping),
		ACCESS_IMPORT_LIST:            util.StringPointerOrEmpty(plan.AccessImportList),
		RAAS_MAPPING_JSON:             util.StringPointerOrEmpty(plan.RAASMappingJSON),
		ACCESS_IMPORT_MAPPING:         util.StringPointerOrEmpty(plan.AccessImportMapping),
		// ORGROLE_IMPORT_PAYLOAD:        util.StringPointerOrEmpty(plan.OrgRoleImportPayload),
		STATUS_KEY_JSON:         util.StringPointerOrEmpty(plan.StatusKeyJSON),
		USERATTRIBUTEJSON:       util.StringPointerOrEmpty(plan.UserAttributeJSON),
		CUSTOM_CONFIG:           util.StringPointerOrEmpty(plan.CustomConfig),
		PAM_CONFIG:              util.StringPointerOrEmpty(plan.PAMConfig),
		MODIFYUSERDATAJSON:      util.StringPointerOrEmpty(plan.ModifyUserDataJSON),
		STATUS_THRESHOLD_CONFIG: util.StringPointerOrEmpty(plan.StatusThresholdConfig),
		CREATE_ACCOUNT_PAYLOAD:  util.StringPointerOrEmpty(plan.CreateAccountPayload),
		UPDATE_ACCOUNT_PAYLOAD:  util.StringPointerOrEmpty(plan.UpdateAccountPayload),
		UPDATE_USER_PAYLOAD:     util.StringPointerOrEmpty(plan.UpdateUserPayload),
		ASSIGN_ORGROLE_PAYLOAD:  util.StringPointerOrEmpty(plan.AssignOrgRolePayload),
		REMOVE_ORGROLE_PAYLOAD:  util.StringPointerOrEmpty(plan.RemoveOrgRolePayload),
	}
	workdayConnRequest := openapi.CreateOrUpdateRequest{
		WorkdayConnector: &workdayConn,
	}

	// Initialize API client
	apiClient := openapi.NewAPIClient(cfg)
	apiResp, _, err := apiClient.ConnectionsAPI.CreateOrUpdate(ctx).CreateOrUpdateRequest(workdayConnRequest).Execute()
	if err != nil {
		resp.Diagnostics.AddError("API Create Failed", fmt.Sprintf("Error: %v", err))
		return
	}
	if apiResp != nil && *apiResp.ErrorCode != "0" {
		log.Printf("[ERROR]: Error in updating Workday connection resource. Errorcode: %v, Message: %v", *apiResp.ErrorCode, *apiResp.Msg)
		resp.Diagnostics.AddError("Updation of Workday connection failed", *apiResp.Msg)
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
	if getResp != nil && *getResp.WorkdayConnectionResponse.Errorcode != 0 {
		log.Printf("[ERROR]: Error in reading Workday connection resource after updation. Errorcode: %v, Message: %v", *getResp.WorkdayConnectionResponse.Errorcode, *getResp.WorkdayConnectionResponse.Msg)
		resp.Diagnostics.AddError("Reading Workday connection after updation failed", *getResp.WorkdayConnectionResponse.Msg)
		return
	}

	plan.ConnectionKey = types.Int64Value(int64(*getResp.WorkdayConnectionResponse.Connectionkey))
	plan.ID = types.StringValue(fmt.Sprintf("%d", *getResp.WorkdayConnectionResponse.Connectionkey))
	plan.ConnectionName = util.SafeStringDatasource(getResp.WorkdayConnectionResponse.Connectionname)
	// plan.Description = util.SafeStringDatasource(getResp.WorkdayConnectionResponse.Description)
	plan.DefaultSavRoles = util.SafeStringDatasource(getResp.WorkdayConnectionResponse.Defaultsavroles)
	plan.ConnectionType = util.SafeStringDatasource(getResp.WorkdayConnectionResponse.Connectiontype)
	plan.EmailTemplate = util.SafeStringDatasource(getResp.WorkdayConnectionResponse.Emailtemplate)
	plan.UseOAuth = util.SafeStringDatasource(getResp.WorkdayConnectionResponse.Connectionattributes.USE_OAUTH)
	plan.UserImportMapping = util.SafeStringDatasource(getResp.WorkdayConnectionResponse.Connectionattributes.USER_IMPORT_MAPPING)
	plan.AccountsLastImportTime = util.SafeStringDatasource(getResp.WorkdayConnectionResponse.Connectionattributes.ACCOUNTS_LAST_IMPORT_TIME)
	plan.StatusKeyJSON = util.SafeStringDatasource(getResp.WorkdayConnectionResponse.Connectionattributes.STATUS_KEY_JSON)
	plan.ConnectionType = util.SafeStringDatasource(getResp.WorkdayConnectionResponse.Connectiontype)
	plan.RAASMappingJSON = util.SafeStringDatasource(getResp.WorkdayConnectionResponse.Connectionattributes.RAAS_MAPPING_JSON)
	plan.AccountImportPayload = util.SafeStringDatasource(getResp.WorkdayConnectionResponse.Connectionattributes.ACCOUNT_IMPORT_PAYLOAD)
	plan.UpdateAccountPayload = util.SafeStringDatasource(getResp.WorkdayConnectionResponse.Connectionattributes.UPDATE_ACCOUNT_PAYLOAD)
	plan.ClientID = util.SafeStringDatasource(getResp.WorkdayConnectionResponse.Connectionattributes.CLIENT_ID)
	plan.StatusThresholdConfig = util.SafeStringDatasource(getResp.WorkdayConnectionResponse.Connectionattributes.STATUS_THRESHOLD_CONFIG)
	plan.Username = util.SafeStringDatasource(getResp.WorkdayConnectionResponse.Connectionattributes.USERNAME)
	plan.AccessImportList = util.SafeStringDatasource(getResp.WorkdayConnectionResponse.Connectionattributes.ACCESS_IMPORT_LIST)
	plan.AccountImportMapping = util.SafeStringDatasource(getResp.WorkdayConnectionResponse.Connectionattributes.ACCOUNT_IMPORT_MAPPING)
	// plan.OrgRoleImportPayload = util.SafeStringDatasource(getResp.WorkdayConnectionResponse.Connectionattributes.ORGROLE_IMPORT_PAYLOAD)
	plan.AssignOrgRolePayload = util.SafeStringDatasource(getResp.WorkdayConnectionResponse.Connectionattributes.ASSIGN_ORGROLE_PAYLOAD)
	plan.AccessImportMapping = util.SafeStringDatasource(getResp.WorkdayConnectionResponse.Connectionattributes.ACCESS_IMPORT_MAPPING)
	plan.APIVersion = util.SafeStringDatasource(getResp.WorkdayConnectionResponse.Connectionattributes.API_VERSION)
	plan.RemoveOrgRolePayload = util.SafeStringDatasource(getResp.WorkdayConnectionResponse.Connectionattributes.REMOVE_ORGROLE_PAYLOAD)
	plan.IncludeReferenceDesc = util.SafeStringDatasource(getResp.WorkdayConnectionResponse.Connectionattributes.INCLUDE_REFERENCE_DESCRIPTORS)
	plan.ModifyUserDataJSON = util.SafeStringDatasource(getResp.WorkdayConnectionResponse.Connectionattributes.MODIFYUSERDATAJSON)
	plan.UseX509AuthForSOAP = util.SafeStringDatasource(getResp.WorkdayConnectionResponse.Connectionattributes.USEX509AUTHFORSOAP)
	plan.ReportOwner = util.SafeStringDatasource(getResp.WorkdayConnectionResponse.Connectionattributes.REPORT_OWNER)
	plan.X509Key = util.SafeStringDatasource(getResp.WorkdayConnectionResponse.Connectionattributes.X509KEY)
	plan.CustomConfig = util.SafeStringDatasource(getResp.WorkdayConnectionResponse.Connectionattributes.CUSTOM_CONFIG)
	plan.UserAttributeJSON = util.SafeStringDatasource(getResp.WorkdayConnectionResponse.Connectionattributes.USERATTRIBUTEJSON)
	plan.X509Cert = util.SafeStringDatasource(getResp.WorkdayConnectionResponse.Connectionattributes.X509CERT)
	plan.UserImportPayload = util.SafeStringDatasource(getResp.WorkdayConnectionResponse.Connectionattributes.USER_IMPORT_PAYLOAD)
	plan.PAMConfig = util.SafeStringDatasource(getResp.WorkdayConnectionResponse.Connectionattributes.PAM_CONFIG)
	plan.AccessLastImportTime = util.SafeStringDatasource(getResp.WorkdayConnectionResponse.Connectionattributes.ACCESS_LAST_IMPORT_TIME)
	plan.UsersLastImportTime = util.SafeStringDatasource(getResp.WorkdayConnectionResponse.Connectionattributes.USERS_LAST_IMPORT_TIME)
	plan.UpdateUserPayload = util.SafeStringDatasource(getResp.WorkdayConnectionResponse.Connectionattributes.UPDATE_USER_PAYLOAD)
	plan.PageSize = util.SafeStringDatasource(getResp.WorkdayConnectionResponse.Connectionattributes.PAGE_SIZE)
	plan.TenantName = util.SafeStringDatasource(getResp.WorkdayConnectionResponse.Connectionattributes.TENANT_NAME)
	plan.UseEnhancedOrgRole = util.SafeStringDatasource(getResp.WorkdayConnectionResponse.Connectionattributes.USE_ENHANCED_ORGROLE)
	plan.CreateAccountPayload = util.SafeStringDatasource(getResp.WorkdayConnectionResponse.Connectionattributes.CREATE_ACCOUNT_PAYLOAD)
	plan.BaseURL = util.SafeStringDatasource(getResp.WorkdayConnectionResponse.Connectionattributes.BASE_URL)
	apiMessage := util.SafeDeref(getResp.WorkdayConnectionResponse.Msg)
	if apiMessage == "success" {
		plan.Msg = types.StringValue("Connection Successful")
	} else {
		plan.Msg = types.StringValue(apiMessage)
	}
	plan.ErrorCode = util.Int32PtrToTFString(getResp.WorkdayConnectionResponse.Errorcode)
	stateUpdateDiagnostics := resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(stateUpdateDiagnostics...)
}

func (r *workdayConnectionResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
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

func (r *workdayConnectionResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	// Retrieve import ID and save to id attribute
	resource.ImportStatePassthroughID(ctx, path.Root("connection_name"), req, resp)
}
