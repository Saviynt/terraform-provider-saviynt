// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"terraform-provider-Saviynt/util"
	"terraform-provider-Saviynt/util/endpointsutil"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"

	"github.com/hashicorp/terraform-plugin-framework/types"
	s "github.com/saviynt/saviynt-api-go-client"

	openapi "github.com/saviynt/saviynt-api-go-client/endpoints"
)

type EndpointResourceModel struct {
	ID                                      types.String `tfsdk:"id"`
	EndpointName                            types.String `tfsdk:"endpoint_name"`
	DisplayName                             types.String `tfsdk:"display_name"`
	SecuritySystem                          types.String `tfsdk:"security_system"`
	Description                             types.String `tfsdk:"description"`
	OwnerType                               types.String `tfsdk:"owner_type"`
	Owner                                   types.String `tfsdk:"owner"`
	ResourceOwnerType                       types.String `tfsdk:"resource_owner_type"`
	ResourceOwner                           types.String `tfsdk:"resource_owner"`
	AccessQuery                             types.String `tfsdk:"access_query"`
	EnableCopyAccess                        types.String `tfsdk:"enable_copy_access"`
	CreateEntTaskforRemoveAcc               types.String `tfsdk:"create_ent_task_for_remove_acc"`
	DisableNewAccountRequestIfAccountExists types.String `tfsdk:"disable_new_account_request_if_account_exists"`
	DisableRemoveAccount                    types.String `tfsdk:"disable_remove_account"`
	DisableModifyAccount                    types.String `tfsdk:"disable_modify_account"`
	OutOfBandAction                         types.String `tfsdk:"out_of_band_action"`
	UserAccountCorrelationRule              types.String `tfsdk:"user_account_correlation_rule"`
	ConnectionConfig                        types.String `tfsdk:"connection_config"`
	Requestable                             types.Bool   `tfsdk:"requestable"`
	ParentAccountPattern                    types.String `tfsdk:"parent_account_pattern"`
	ServiceAccountNameRule                  types.String `tfsdk:"service_account_name_rule"`
	ServiceAccountAccessQuery               types.String `tfsdk:"service_account_access_query"`
	ChangePasswordAccessQuery               types.String `tfsdk:"change_password_access_query"`
	BlockInflightRequest                    types.String `tfsdk:"block_inflight_request"`
	AccountNameRule                         types.String `tfsdk:"account_name_rule"`
	AllowChangePasswordSQLQuery             types.String `tfsdk:"allow_change_password_sql_query"`
	AccountNameValidatorRegex               types.String `tfsdk:"account_name_validator_regex"`
	StatusConfig                            types.String `tfsdk:"status_config"`
	PluginConfigs                           types.String `tfsdk:"plugin_configs"`
	PrimaryAccountType                      types.String `tfsdk:"primary_account_type"`
	AccountTypeNoPasswordChange             types.String `tfsdk:"account_type_no_password_change"`
	EndpointConfig                          types.String `tfsdk:"endpoint_config"`
	AllowRemoveAllRoleOnRequest             types.Bool   `tfsdk:"allow_remove_all_role_on_request"`

	CustomProperty1              types.String `tfsdk:"custom_property1"`
	CustomProperty2              types.String `tfsdk:"custom_property2"`
	CustomProperty3              types.String `tfsdk:"custom_property3"`
	CustomProperty4              types.String `tfsdk:"custom_property4"`
	CustomProperty5              types.String `tfsdk:"custom_property5"`
	CustomProperty6              types.String `tfsdk:"custom_property6"`
	CustomProperty7              types.String `tfsdk:"custom_property7"`
	CustomProperty8              types.String `tfsdk:"custom_property8"`
	CustomProperty9              types.String `tfsdk:"custom_property9"`
	CustomProperty10             types.String `tfsdk:"custom_property10"`
	CustomProperty11             types.String `tfsdk:"custom_property11"`
	CustomProperty12             types.String `tfsdk:"custom_property12"`
	CustomProperty13             types.String `tfsdk:"custom_property13"`
	CustomProperty14             types.String `tfsdk:"custom_property14"`
	CustomProperty15             types.String `tfsdk:"custom_property15"`
	CustomProperty16             types.String `tfsdk:"custom_property16"`
	CustomProperty17             types.String `tfsdk:"custom_property17"`
	CustomProperty18             types.String `tfsdk:"custom_property18"`
	CustomProperty19             types.String `tfsdk:"custom_property19"`
	CustomProperty20             types.String `tfsdk:"custom_property20"`
	CustomProperty21             types.String `tfsdk:"custom_property21"`
	CustomProperty22             types.String `tfsdk:"custom_property22"`
	CustomProperty23             types.String `tfsdk:"custom_property23"`
	CustomProperty24             types.String `tfsdk:"custom_property24"`
	CustomProperty25             types.String `tfsdk:"custom_property25"`
	CustomProperty26             types.String `tfsdk:"custom_property26"`
	CustomProperty27             types.String `tfsdk:"custom_property27"`
	CustomProperty28             types.String `tfsdk:"custom_property28"`
	CustomProperty29             types.String `tfsdk:"custom_property29"`
	CustomProperty30             types.String `tfsdk:"custom_property30"`
	CustomProperty31             types.String `tfsdk:"custom_property31"`
	CustomProperty32             types.String `tfsdk:"custom_property32"`
	CustomProperty33             types.String `tfsdk:"custom_property33"`
	CustomProperty34             types.String `tfsdk:"custom_property34"`
	CustomProperty35             types.String `tfsdk:"custom_property35"`
	CustomProperty36             types.String `tfsdk:"custom_property36"`
	CustomProperty37             types.String `tfsdk:"custom_property37"`
	CustomProperty38             types.String `tfsdk:"custom_property38"`
	CustomProperty39             types.String `tfsdk:"custom_property39"`
	CustomProperty40             types.String `tfsdk:"custom_property40"`
	CustomProperty41             types.String `tfsdk:"custom_property41"`
	CustomProperty42             types.String `tfsdk:"custom_property42"`
	CustomProperty43             types.String `tfsdk:"custom_property43"`
	CustomProperty44             types.String `tfsdk:"custom_property44"`
	CustomProperty45             types.String `tfsdk:"custom_property45"`
	AccountCustomProperty1Label  types.String `tfsdk:"account_custom_property_1_label"`
	AccountCustomProperty2Label  types.String `tfsdk:"account_custom_property_2_label"`
	AccountCustomProperty3Label  types.String `tfsdk:"account_custom_property_3_label"`
	AccountCustomProperty4Label  types.String `tfsdk:"account_custom_property_4_label"`
	AccountCustomProperty5Label  types.String `tfsdk:"account_custom_property_5_label"`
	AccountCustomProperty6Label  types.String `tfsdk:"account_custom_property_6_label"`
	AccountCustomProperty7Label  types.String `tfsdk:"account_custom_property_7_label"`
	AccountCustomProperty8Label  types.String `tfsdk:"account_custom_property_8_label"`
	AccountCustomProperty9Label  types.String `tfsdk:"account_custom_property_9_label"`
	AccountCustomProperty10Label types.String `tfsdk:"account_custom_property_10_label"`
	AccountCustomProperty11Label types.String `tfsdk:"account_custom_property_11_label"`
	AccountCustomProperty12Label types.String `tfsdk:"account_custom_property_12_label"`
	AccountCustomProperty13Label types.String `tfsdk:"account_custom_property_13_label"`
	AccountCustomProperty14Label types.String `tfsdk:"account_custom_property_14_label"`
	AccountCustomProperty15Label types.String `tfsdk:"account_custom_property_15_label"`
	AccountCustomProperty16Label types.String `tfsdk:"account_custom_property_16_label"`
	AccountCustomProperty17Label types.String `tfsdk:"account_custom_property_17_label"`
	AccountCustomProperty18Label types.String `tfsdk:"account_custom_property_18_label"`
	AccountCustomProperty19Label types.String `tfsdk:"account_custom_property_19_label"`
	AccountCustomProperty20Label types.String `tfsdk:"account_custom_property_20_label"`
	AccountCustomProperty21Label types.String `tfsdk:"account_custom_property_21_label"`
	AccountCustomProperty22Label types.String `tfsdk:"account_custom_property_22_label"`
	AccountCustomProperty23Label types.String `tfsdk:"account_custom_property_23_label"`
	AccountCustomProperty24Label types.String `tfsdk:"account_custom_property_24_label"`
	AccountCustomProperty25Label types.String `tfsdk:"account_custom_property_25_label"`
	AccountCustomProperty26Label types.String `tfsdk:"account_custom_property_26_label"`
	AccountCustomProperty27Label types.String `tfsdk:"account_custom_property_27_label"`
	AccountCustomProperty28Label types.String `tfsdk:"account_custom_property_28_label"`
	AccountCustomProperty29Label types.String `tfsdk:"account_custom_property_29_label"`
	AccountCustomProperty30Label types.String `tfsdk:"account_custom_property_30_label"`
	CustomProperty31Label        types.String `tfsdk:"custom_property31_label"`
	CustomProperty32Label        types.String `tfsdk:"custom_property32_label"`
	CustomProperty33Label        types.String `tfsdk:"custom_property33_label"`
	CustomProperty34Label        types.String `tfsdk:"custom_property34_label"`
	CustomProperty35Label        types.String `tfsdk:"custom_property35_label"`
	CustomProperty36Label        types.String `tfsdk:"custom_property36_label"`
	CustomProperty37Label        types.String `tfsdk:"custom_property37_label"`
	CustomProperty38Label        types.String `tfsdk:"custom_property38_label"`
	CustomProperty39Label        types.String `tfsdk:"custom_property39_label"`
	CustomProperty40Label        types.String `tfsdk:"custom_property40_label"`
	CustomProperty41Label        types.String `tfsdk:"custom_property41_label"`
	CustomProperty42Label        types.String `tfsdk:"custom_property42_label"`
	CustomProperty43Label        types.String `tfsdk:"custom_property43_label"`
	CustomProperty44Label        types.String `tfsdk:"custom_property44_label"`
	CustomProperty45Label        types.String `tfsdk:"custom_property45_label"`
	CustomProperty46Label        types.String `tfsdk:"custom_property46_label"`
	CustomProperty47Label        types.String `tfsdk:"custom_property47_label"`
	CustomProperty48Label        types.String `tfsdk:"custom_property48_label"`
	CustomProperty49Label        types.String `tfsdk:"custom_property49_label"`
	CustomProperty50Label        types.String `tfsdk:"custom_property50_label"`
	CustomProperty51Label        types.String `tfsdk:"custom_property51_label"`
	CustomProperty52Label        types.String `tfsdk:"custom_property52_label"`
	CustomProperty53Label        types.String `tfsdk:"custom_property53_label"`
	CustomProperty54Label        types.String `tfsdk:"custom_property54_label"`
	CustomProperty55Label        types.String `tfsdk:"custom_property55_label"`
	CustomProperty56Label        types.String `tfsdk:"custom_property56_label"`
	CustomProperty57Label        types.String `tfsdk:"custom_property57_label"`
	CustomProperty58Label        types.String `tfsdk:"custom_property58_label"`
	CustomProperty59Label        types.String `tfsdk:"custom_property59_label"`
	CustomProperty60Label        types.String `tfsdk:"custom_property60_label"`
	MappedEndpoints              types.List   `tfsdk:"mapped_endpoints"`
	EmailTemplates               types.List   `tfsdk:"email_templates"`
	RequestableRoleTypes         types.List   `tfsdk:"requestable_role_types"`
	Msg                          types.String `tfsdk:"msg"`
	ErrorCode                    types.String `tfsdk:"error_code"`
}

type endpointResource struct {
	client *s.Client
	token  string
}

type RequestableRoleType struct {
	RoleType       types.String `tfsdk:"role_type"`
	RequestOption  types.String `tfsdk:"request_option"`
	Required       types.Bool   `tfsdk:"required"`
	RequestedQuery types.String `tfsdk:"requested_query"`
	SelectedQuery  types.String `tfsdk:"selected_query"`
	ShowOn         types.String `tfsdk:"show_on"`
}

type EmailTemplate struct {
	EmailTemplateType types.String `tfsdk:"email_template_type"`
	TaskType          types.String `tfsdk:"task_type"`
	EmailTemplate     types.String `tfsdk:"email_template"`
}

type MappedEndpoint struct {
	SecuritySystem types.String `tfsdk:"security_system"`
	Endpoint       types.String `tfsdk:"endpoint"`
	Requestable    types.String `tfsdk:"requestable"`
	Operation      types.String `tfsdk:"operation"`
}

func NewEndpointResource() resource.Resource {
	return &endpointResource{}
}

func (r *endpointResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = "saviynt_endpoint_resource"
}

func (r *endpointResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: util.EndpointDescription,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The unique ID of the resource.",
			},
			"endpoint_name": schema.StringAttribute{
				Required:    true,
				Description: "Specify a name for the endpoint. Provide a logical name that will help you easily identify it.",
			},
			"display_name": schema.StringAttribute{
				Required:    true,
				Description: "Enter a user-friendly display name for the endpoint that will be displayed in the user interface. Display Name can be different from Endpoint Name.",
			},
			"security_system": schema.StringAttribute{
				Required:    true,
				Description: "Specify the Security system for which you want to create an endpoint.",
			},
			"description": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Specify a description for the endpoint.",
			},
			"owner_type": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Specify the owner type for the endpoint. An endpoint can be owned by a User or Usergroup.",
			},
			"owner": schema.StringAttribute{
				Optional:    true,
				Description: "Specify the owner of the endpoint. If the ownerType is User, then specify the username of the owner, and If it is is Usergroup then specify the name of the user group.",
			},
			"resource_owner_type": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Specify the resource owner type for the endpoint. An endpoint can be owned by a User or Usergroup.",
			},
			"resource_owner": schema.StringAttribute{
				Optional:    true,
				Description: "Specify the resource owner of the endpoint. If the resourceOwnerType is User, then specify the username of the owner and If it is Usergroup, specify the name of the user group.",
			},
			"access_query": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Specify the query to filter the access and display of the endpoint to specific users. If you do not define a query, the endpoint is displayed for all users.",
			},
			"enable_copy_access": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Specify true to display the Copy Access from User option in the Request pages.",
			},
			"disable_new_account_request_if_account_exists": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Specify true to disable users from requesting additional accounts on applications where they already have active accounts.",
			},
			"disable_remove_account": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Specify true to disable users from removing their existing application accounts.",
			},
			"disable_modify_account": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Specify true to disable users from modifying their application accounts.",
			},
			"user_account_correlation_rule": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Specify rule to map users in EIC with the accounts during account import.",
			},
			"create_ent_task_for_remove_acc": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "If this is set to true, remove Access tasks will be created for entitlements (account entitlements and their dependent entitlements) when a user requests for removing an account.",
			},
			"out_of_band_action": schema.StringAttribute{
				Optional:    true,
				Description: "Use this parameter to determine if you need to remove the accesses which were granted outside Saviynt.",
			},
			"connection_config": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Use this configuration for processing the add access tasks and remove access tasks for AD and LDAP Connectors.",
			},
			"requestable": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Is this endpoint requestable.",
			},
			"parent_account_pattern": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Specify the parent and child relationship for the Active Directory endpoint. The specified value is used to filter the parent and child objects in the Request Access tile.",
			},
			"service_account_name_rule": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Rule to generate a name for this endpoint while creating a new service account.",
			},
			"service_account_access_query": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Specify the query to filter the access and display of the endpoint for specific users while managing service accounts.",
			},
			"block_inflight_request": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Specify true to prevent users from raising duplicate requests for the same applications.",
			},
			"account_name_rule": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Specify rule to generate an account name for this endpoint while creating a new account.",
			},
			"allow_change_password_sql_query": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "SQL query to configure the accounts for which you can change passwords.",
			},
			"account_name_validator_regex": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Specify the regular expression which will be used to validate the account name either generated by the rule or provided manually.",
			},
			"status_config": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enable the State and Status options (Enable, Disable, Lock, Unlock) that would be available to request for a user and service accounts.",
			},
			"plugin_configs": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The Plugin Configuration drives the functionality of the Saviynt SmartAssist (Browserplugin).",
			},
			"endpoint_config": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Option to copy data in Step 3 of the service account request will be enabled.",
			},
			"primary_account_type": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Type of primary account",
			},
			"account_type_no_password_change": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Account type no password change",
			},
			"mapped_endpoints": schema.ListAttribute{
				ElementType: types.ObjectType{
					AttrTypes: map[string]attr.Type{
						"security_system": types.StringType,
						"endpoint":        types.StringType,
						"requestable":     types.StringType,
						"operation":       types.StringType,
					},
				},
				Optional: true,
			},
			"email_templates": schema.ListAttribute{
				ElementType: types.ObjectType{
					AttrTypes: map[string]attr.Type{
						"email_template_type": types.StringType,
						"task_type":           types.StringType,
						"email_template":      types.StringType,
					},
				},
				Optional: true,
				Computed: true,
			},

			"requestable_role_types": schema.ListAttribute{
				ElementType: types.ObjectType{
					AttrTypes: map[string]attr.Type{
						"role_type":       types.StringType,
						"request_option":  types.StringType,
						"required":        types.BoolType,
						"requested_query": types.StringType,
						"selected_query":  types.StringType,
						"show_on":         types.StringType,
					},
				},
				Optional: true,
				Computed: true,
			},
			"msg": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "A message indicating the outcome of the operation.",
			},
			"error_code": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "An error code where '0' signifies success and '1' signifies an unsuccessful operation.",
			},
		},
	}

	for i := 1; i <= 45; i++ {
		key := fmt.Sprintf("custom_property%d", i)
		resp.Schema.Attributes[key] = schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: fmt.Sprintf("Custom Property %d.", i),
		}
	}

	for i := 1; i <= 30; i++ {
		key := fmt.Sprintf("account_custom_property_%d_label", i)
		resp.Schema.Attributes[key] = schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: fmt.Sprintf("Account Custom Property label %d.", i),
		}
	}

	for i := 31; i <= 60; i++ {
		key := fmt.Sprintf("custom_property%d_label", i)
		resp.Schema.Attributes[key] = schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: fmt.Sprintf("Label for the custom property %d of accounts of this endpoint.", i),
		}
	}

	resp.Schema.Attributes["allow_remove_all_role_on_request"] = schema.BoolAttribute{
		Optional:    true,
		Computed:    true,
		Description: "Specify true to displays the Remove All Roles option in the Request page that can be used to remove all the roles.",
	}

	resp.Schema.Attributes["change_password_access_query"] = schema.StringAttribute{
		Optional:    true,
		Computed:    true,
		Description: "Specify query to restrict the access for changing the account password of the endpoint.",
	}
}

func (r *endpointResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *endpointResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan EndpointResourceModel

	planGetDiagnostics := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(planGetDiagnostics...)
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

	reqParams := openapi.GetEndpointsRequest{}
	reqParams.SetEndpointname(plan.EndpointName.ValueString())
	existingResource, _, err := apiClient.EndpointsAPI.GetEndpoints(ctx).GetEndpointsRequest(reqParams).Execute()
	if err != nil {
		log.Printf("Problem with the get function in Create block")
		resp.Diagnostics.AddError("API Read Failed In Create Block", fmt.Sprintf("Error: %v", err))
		return
	}

	if existingResource != nil &&
		len(existingResource.Endpoints) != 0 &&
		*existingResource.ErrorCode == "0" {
		log.Printf("[ERROR] Endpoint name already exists. Please import or use a different name")
		resp.Diagnostics.AddError("Endpoint name already exists", "Please import or use a different name")
		return
	}

	if plan.CreateEntTaskforRemoveAcc.IsNull() || plan.CreateEntTaskforRemoveAcc.IsUnknown() || plan.CreateEntTaskforRemoveAcc.ValueString() == "" {
		plan.CreateEntTaskforRemoveAcc = types.StringValue("false")
	}
	createReq := openapi.CreateEndpointRequest{
		Endpointname:                            plan.EndpointName.ValueString(),
		DisplayName:                             plan.DisplayName.ValueString(),
		Securitysystem:                          plan.SecuritySystem.ValueString(),
		Description:                             util.StringPointerOrEmpty(plan.Description),
		OwnerType:                               util.StringPointerOrEmpty(plan.OwnerType),
		Owner:                                   util.StringPointerOrEmpty(plan.Owner),
		ResourceOwnerType:                       util.StringPointerOrEmpty(plan.ResourceOwnerType),
		ResourceOwner:                           util.StringPointerOrEmpty(plan.ResourceOwner),
		Accessquery:                             util.StringPointerOrEmpty(plan.AccessQuery),
		EnableCopyAccess:                        util.StringPointerOrEmpty(plan.EnableCopyAccess),
		DisableNewAccountRequestIfAccountExists: util.StringPointerOrEmpty(plan.DisableNewAccountRequestIfAccountExists),
		DisableRemoveAccount:                    util.StringPointerOrEmpty(plan.DisableRemoveAccount),
		DisableModifyAccount:                    util.StringPointerOrEmpty(plan.DisableModifyAccount),
		UserAccountCorrelationRule:              util.StringPointerOrEmpty(plan.UserAccountCorrelationRule),
		CreateEntTaskforRemoveAcc:               util.StringPointerOrEmpty(plan.CreateEntTaskforRemoveAcc),
		Outofbandaction:                         util.StringPointerOrEmpty(plan.OutOfBandAction),
		Connectionconfig:                        util.StringPointerOrEmpty(plan.ConnectionConfig),
		Requestable:                             util.BoolPointerOrEmtpy(plan.Requestable),
		ParentAccountPattern:                    util.StringPointerOrEmpty(plan.ParentAccountPattern),
		ServiceAccountNameRule:                  util.StringPointerOrEmpty(plan.ServiceAccountNameRule),
		ServiceAccountAccessQuery:               util.StringPointerOrEmpty(plan.ServiceAccountAccessQuery),
		BlockInflightRequest:                    util.StringPointerOrEmpty(plan.BlockInflightRequest),
		AccountNameRule:                         util.StringPointerOrEmpty(plan.AccountNameRule),
		AllowChangePasswordSqlquery:             util.StringPointerOrEmpty(plan.AllowChangePasswordSQLQuery),
		AccountNameValidatorRegex:               util.StringPointerOrEmpty(plan.AccountNameValidatorRegex),
		PrimaryAccountType:                      util.StringPointerOrEmpty(plan.PrimaryAccountType),
		AccountTypeNoPasswordChange:             util.StringPointerOrEmpty(plan.AccountTypeNoPasswordChange),
		ChangePasswordAccessQuery:               util.StringPointerOrEmpty(plan.ChangePasswordAccessQuery),
		StatusConfig:                            util.StringPointerOrEmpty(plan.StatusConfig),
		PluginConfigs:                           util.StringPointerOrEmpty(plan.PluginConfigs),
		EndpointConfig:                          util.StringPointerOrEmpty(plan.EndpointConfig),
		AllowRemoveAllRoleOnRequest:             util.BoolPointerOrEmtpy(plan.AllowRemoveAllRoleOnRequest),
		Customproperty1:                         util.StringPointerOrEmpty(plan.CustomProperty1),
		Customproperty2:                         util.StringPointerOrEmpty(plan.CustomProperty2),
		Customproperty3:                         util.StringPointerOrEmpty(plan.CustomProperty3),
		Customproperty4:                         util.StringPointerOrEmpty(plan.CustomProperty4),
		Customproperty5:                         util.StringPointerOrEmpty(plan.CustomProperty5),
		Customproperty6:                         util.StringPointerOrEmpty(plan.CustomProperty6),
		Customproperty7:                         util.StringPointerOrEmpty(plan.CustomProperty7),
		Customproperty8:                         util.StringPointerOrEmpty(plan.CustomProperty8),
		Customproperty9:                         util.StringPointerOrEmpty(plan.CustomProperty9),
		Customproperty10:                        util.StringPointerOrEmpty(plan.CustomProperty10),
		Customproperty11:                        util.StringPointerOrEmpty(plan.CustomProperty11),
		Customproperty12:                        util.StringPointerOrEmpty(plan.CustomProperty12),
		Customproperty13:                        util.StringPointerOrEmpty(plan.CustomProperty13),
		Customproperty14:                        util.StringPointerOrEmpty(plan.CustomProperty14),
		Customproperty15:                        util.StringPointerOrEmpty(plan.CustomProperty15),
		Customproperty16:                        util.StringPointerOrEmpty(plan.CustomProperty16),
		Customproperty17:                        util.StringPointerOrEmpty(plan.CustomProperty17),
		Customproperty18:                        util.StringPointerOrEmpty(plan.CustomProperty18),
		Customproperty19:                        util.StringPointerOrEmpty(plan.CustomProperty19),
		Customproperty20:                        util.StringPointerOrEmpty(plan.CustomProperty20),
		Customproperty21:                        util.StringPointerOrEmpty(plan.CustomProperty21),
		Customproperty22:                        util.StringPointerOrEmpty(plan.CustomProperty22),
		Customproperty23:                        util.StringPointerOrEmpty(plan.CustomProperty23),
		Customproperty24:                        util.StringPointerOrEmpty(plan.CustomProperty24),
		Customproperty25:                        util.StringPointerOrEmpty(plan.CustomProperty25),
		Customproperty26:                        util.StringPointerOrEmpty(plan.CustomProperty26),
		Customproperty27:                        util.StringPointerOrEmpty(plan.CustomProperty27),
		Customproperty28:                        util.StringPointerOrEmpty(plan.CustomProperty28),
		Customproperty29:                        util.StringPointerOrEmpty(plan.CustomProperty29),
		Customproperty30:                        util.StringPointerOrEmpty(plan.CustomProperty30),
		Customproperty31:                        util.StringPointerOrEmpty(plan.CustomProperty31),
		Customproperty32:                        util.StringPointerOrEmpty(plan.CustomProperty32),
		Customproperty33:                        util.StringPointerOrEmpty(plan.CustomProperty33),
		Customproperty34:                        util.StringPointerOrEmpty(plan.CustomProperty34),
		Customproperty35:                        util.StringPointerOrEmpty(plan.CustomProperty35),
		Customproperty36:                        util.StringPointerOrEmpty(plan.CustomProperty36),
		Customproperty37:                        util.StringPointerOrEmpty(plan.CustomProperty37),
		Customproperty38:                        util.StringPointerOrEmpty(plan.CustomProperty38),
		Customproperty39:                        util.StringPointerOrEmpty(plan.CustomProperty39),
		Customproperty40:                        util.StringPointerOrEmpty(plan.CustomProperty40),
		Customproperty41:                        util.StringPointerOrEmpty(plan.CustomProperty41),
		Customproperty42:                        util.StringPointerOrEmpty(plan.CustomProperty42),
		Customproperty43:                        util.StringPointerOrEmpty(plan.CustomProperty43),
		Customproperty44:                        util.StringPointerOrEmpty(plan.CustomProperty44),
		Customproperty45:                        util.StringPointerOrEmpty(plan.CustomProperty45),
		Customproperty1Label:                    util.StringPointerOrEmpty(plan.AccountCustomProperty1Label),
		Customproperty2Label:                    util.StringPointerOrEmpty(plan.AccountCustomProperty2Label),
		Customproperty3Label:                    util.StringPointerOrEmpty(plan.AccountCustomProperty3Label),
		Customproperty4Label:                    util.StringPointerOrEmpty(plan.AccountCustomProperty4Label),
		Customproperty5Label:                    util.StringPointerOrEmpty(plan.AccountCustomProperty5Label),
		Customproperty6Label:                    util.StringPointerOrEmpty(plan.AccountCustomProperty6Label),
		Customproperty7Label:                    util.StringPointerOrEmpty(plan.AccountCustomProperty7Label),
		Customproperty8Label:                    util.StringPointerOrEmpty(plan.AccountCustomProperty8Label),
		Customproperty9Label:                    util.StringPointerOrEmpty(plan.AccountCustomProperty9Label),
		Customproperty10Label:                   util.StringPointerOrEmpty(plan.AccountCustomProperty10Label),
		Customproperty11Label:                   util.StringPointerOrEmpty(plan.AccountCustomProperty11Label),
		Customproperty12Label:                   util.StringPointerOrEmpty(plan.AccountCustomProperty12Label),
		Customproperty13Label:                   util.StringPointerOrEmpty(plan.AccountCustomProperty13Label),
		Customproperty14Label:                   util.StringPointerOrEmpty(plan.AccountCustomProperty14Label),
		Customproperty15Label:                   util.StringPointerOrEmpty(plan.AccountCustomProperty15Label),
		Customproperty16Label:                   util.StringPointerOrEmpty(plan.AccountCustomProperty16Label),
		Customproperty17Label:                   util.StringPointerOrEmpty(plan.AccountCustomProperty17Label),
		Customproperty18Label:                   util.StringPointerOrEmpty(plan.AccountCustomProperty18Label),
		Customproperty19Label:                   util.StringPointerOrEmpty(plan.AccountCustomProperty19Label),
		Customproperty20Label:                   util.StringPointerOrEmpty(plan.AccountCustomProperty20Label),
		Customproperty21Label:                   util.StringPointerOrEmpty(plan.AccountCustomProperty21Label),
		Customproperty22Label:                   util.StringPointerOrEmpty(plan.AccountCustomProperty22Label),
		Customproperty23Label:                   util.StringPointerOrEmpty(plan.AccountCustomProperty23Label),
		Customproperty24Label:                   util.StringPointerOrEmpty(plan.AccountCustomProperty24Label),
		Customproperty25Label:                   util.StringPointerOrEmpty(plan.AccountCustomProperty25Label),
		Customproperty26Label:                   util.StringPointerOrEmpty(plan.AccountCustomProperty26Label),
		Customproperty27Label:                   util.StringPointerOrEmpty(plan.AccountCustomProperty27Label),
		Customproperty28Label:                   util.StringPointerOrEmpty(plan.AccountCustomProperty28Label),
		Customproperty29Label:                   util.StringPointerOrEmpty(plan.AccountCustomProperty29Label),
		Customproperty30Label:                   util.StringPointerOrEmpty(plan.AccountCustomProperty30Label),
		Customproperty31Label:                   util.StringPointerOrEmpty(plan.CustomProperty31Label),
		Customproperty32Label:                   util.StringPointerOrEmpty(plan.CustomProperty32Label),
		Customproperty33Label:                   util.StringPointerOrEmpty(plan.CustomProperty33Label),
		Customproperty34Label:                   util.StringPointerOrEmpty(plan.CustomProperty34Label),
		Customproperty35Label:                   util.StringPointerOrEmpty(plan.CustomProperty35Label),
		Customproperty36Label:                   util.StringPointerOrEmpty(plan.CustomProperty36Label),
		Customproperty37Label:                   util.StringPointerOrEmpty(plan.CustomProperty37Label),
		Customproperty38Label:                   util.StringPointerOrEmpty(plan.CustomProperty38Label),
		Customproperty39Label:                   util.StringPointerOrEmpty(plan.CustomProperty39Label),
		Customproperty40Label:                   util.StringPointerOrEmpty(plan.CustomProperty40Label),
		Customproperty41Label:                   util.StringPointerOrEmpty(plan.CustomProperty41Label),
		Customproperty42Label:                   util.StringPointerOrEmpty(plan.CustomProperty42Label),
		Customproperty43Label:                   util.StringPointerOrEmpty(plan.CustomProperty43Label),
		Customproperty44Label:                   util.StringPointerOrEmpty(plan.CustomProperty44Label),
		Customproperty45Label:                   util.StringPointerOrEmpty(plan.CustomProperty45Label),
		Customproperty46Label:                   util.StringPointerOrEmpty(plan.CustomProperty46Label),
		Customproperty47Label:                   util.StringPointerOrEmpty(plan.CustomProperty47Label),
		Customproperty48Label:                   util.StringPointerOrEmpty(plan.CustomProperty48Label),
		Customproperty49Label:                   util.StringPointerOrEmpty(plan.CustomProperty49Label),
		Customproperty50Label:                   util.StringPointerOrEmpty(plan.CustomProperty50Label),
		Customproperty51Label:                   util.StringPointerOrEmpty(plan.CustomProperty51Label),
		Customproperty52Label:                   util.StringPointerOrEmpty(plan.CustomProperty52Label),
		Customproperty53Label:                   util.StringPointerOrEmpty(plan.CustomProperty53Label),
		Customproperty54Label:                   util.StringPointerOrEmpty(plan.CustomProperty54Label),
		Customproperty55Label:                   util.StringPointerOrEmpty(plan.CustomProperty55Label),
		Customproperty56Label:                   util.StringPointerOrEmpty(plan.CustomProperty56Label),
		Customproperty57Label:                   util.StringPointerOrEmpty(plan.CustomProperty57Label),
		Customproperty58Label:                   util.StringPointerOrEmpty(plan.CustomProperty58Label),
		Customproperty59Label:                   util.StringPointerOrEmpty(plan.CustomProperty59Label),
		Customproperty60Label:                   util.StringPointerOrEmpty(plan.CustomProperty60Label),
	}

	var emailTemplates []openapi.CreateEndpointRequestEmailTemplateInner
	var diags diag.Diagnostics
	var tfEmailTemplates []EmailTemplate

	// Convert from Terraform types to Go struct (allowing unknown values)
	diags = plan.EmailTemplates.ElementsAs(ctx, &tfEmailTemplates, true)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	for _, tfTemplate := range tfEmailTemplates {
		if tfTemplate.EmailTemplateType.IsUnknown() &&
			tfTemplate.TaskType.IsUnknown() &&
			tfTemplate.EmailTemplate.IsUnknown() {
			continue
		}

		emailTemplate := openapi.CreateEndpointRequestEmailTemplateInner{}

		if !tfTemplate.EmailTemplateType.IsNull() {
			emailTemplate.EmailTemplateType = tfTemplate.EmailTemplateType.ValueStringPointer()
		}
		if !tfTemplate.TaskType.IsNull() {
			emailTemplate.TaskType = tfTemplate.TaskType.ValueStringPointer()
		}
		if !tfTemplate.EmailTemplate.IsNull() {
			emailTemplate.EmailTemplate = tfTemplate.EmailTemplate.ValueStringPointer()
		}

		emailTemplates = append(emailTemplates, emailTemplate)
	}

	if len(emailTemplates) > 0 {
		createReq.Taskemailtemplates = emailTemplates
	}

	apiResp, httpResp, err := apiClient.EndpointsAPI.
		CreateEndpoint(ctx).
		CreateEndpointRequest(createReq).
		Execute()

	if err != nil {
		log.Printf("Error Creating Endpoint: %v, HTTP Response: %v", err, httpResp)
		resp.Diagnostics.AddError(
			"Error Creating Endpoint",
			fmt.Sprintf("Error creating endpoint: %s", err),
		)
		return
	}
	if apiResp != nil && *apiResp.ErrorCode != "0" {
		log.Printf("Error Creating Endpoint: %v, Error code: %v", *apiResp.Msg, *apiResp.ErrorCode)
		resp.Diagnostics.AddError(
			"Error Creating Endpoint",
			fmt.Sprintf("Error: %v, Error code: %v", *apiResp.Msg, *apiResp.ErrorCode),
		)
		return
	}

	plan.ID = types.StringValue("endpoint-" + plan.EndpointName.ValueString())
	if plan.EnableCopyAccess.IsNull() || plan.EnableCopyAccess.IsUnknown() || plan.EnableCopyAccess.ValueString() == "" {
		plan.EnableCopyAccess = types.StringValue("false")
	}
	if plan.AllowRemoveAllRoleOnRequest.IsNull() || plan.AllowRemoveAllRoleOnRequest.IsUnknown() {
		plan.AllowRemoveAllRoleOnRequest = types.BoolValue(false)
	} else {
		plan.AllowRemoveAllRoleOnRequest = types.BoolValue(plan.AllowRemoveAllRoleOnRequest.ValueBool())
	}
	if plan.Requestable.IsNull() || plan.Requestable.IsUnknown() {
		plan.Requestable = types.BoolValue(true)
	} else {
		plan.Requestable = types.BoolValue(plan.Requestable.ValueBool())
	}
	if plan.DisableRemoveAccount.IsNull() || plan.DisableRemoveAccount.IsUnknown() || plan.DisableRemoveAccount.ValueString() == "" {
		plan.DisableRemoveAccount = types.StringValue("false")
	}
	if plan.BlockInflightRequest.IsNull() || plan.BlockInflightRequest.IsUnknown() || plan.BlockInflightRequest.ValueString() == "" {
		plan.BlockInflightRequest = types.StringValue("false")
	}
	if plan.DisableModifyAccount.IsNull() || plan.DisableModifyAccount.IsUnknown() || plan.DisableModifyAccount.ValueString() == "" {
		plan.DisableModifyAccount = types.StringValue("false")
	}
	if plan.DisableNewAccountRequestIfAccountExists.IsNull() || plan.DisableNewAccountRequestIfAccountExists.IsUnknown() || plan.DisableNewAccountRequestIfAccountExists.ValueString() == "" {
		plan.DisableNewAccountRequestIfAccountExists = types.StringValue("false")
	}

	if plan.EmailTemplates.IsNull() || plan.EmailTemplates.IsUnknown() {
		plan.EmailTemplates = types.ListValueMust(types.ObjectType{
			AttrTypes: map[string]attr.Type{
				"email_template_type": types.StringType,
				"task_type":           types.StringType,
				"email_template":      types.StringType,
			},
		},
			[]attr.Value{},
		)
	}
	if plan.RequestableRoleTypes.IsNull() || plan.RequestableRoleTypes.IsUnknown() {
		plan.RequestableRoleTypes = types.ListValueMust(types.ObjectType{
			AttrTypes: map[string]attr.Type{
				"role_type":       types.StringType,
				"request_option":  types.StringType,
				"required":        types.BoolType,
				"requested_query": types.StringType,
				"selected_query":  types.StringType,
				"show_on":         types.StringType,
			},
		},
			[]attr.Value{},
		)
	}

	plan.Description = util.SafeString(plan.Description.ValueStringPointer())
	plan.OwnerType = util.SafeStringDatasource(util.StringPtr(endpointsutil.TranslateValue(util.SafeStringValue(plan.OwnerType), endpointsutil.OwnerTypeMap)))
	plan.ResourceOwnerType = util.SafeStringDatasource(util.StringPtr(endpointsutil.TranslateValue(util.SafeStringValue(plan.ResourceOwnerType), endpointsutil.OwnerTypeMap)))
	plan.AccessQuery = util.SafeString(plan.AccessQuery.ValueStringPointer())
	plan.EnableCopyAccess = util.SafeString(plan.EnableCopyAccess.ValueStringPointer())
	plan.DisableNewAccountRequestIfAccountExists = util.SafeString(plan.DisableNewAccountRequestIfAccountExists.ValueStringPointer())
	plan.DisableRemoveAccount = util.SafeString(plan.DisableRemoveAccount.ValueStringPointer())
	plan.DisableModifyAccount = util.SafeString(plan.DisableModifyAccount.ValueStringPointer())
	plan.UserAccountCorrelationRule = util.SafeString(plan.UserAccountCorrelationRule.ValueStringPointer())
	plan.CreateEntTaskforRemoveAcc = util.SafeString(plan.CreateEntTaskforRemoveAcc.ValueStringPointer())
	plan.ConnectionConfig = util.SafeString(plan.ConnectionConfig.ValueStringPointer())
	plan.ParentAccountPattern = util.SafeString(plan.ParentAccountPattern.ValueStringPointer())
	plan.ServiceAccountNameRule = util.SafeString(plan.ServiceAccountNameRule.ValueStringPointer())
	plan.ServiceAccountAccessQuery = util.SafeString(plan.ServiceAccountAccessQuery.ValueStringPointer())
	plan.BlockInflightRequest = util.SafeString(plan.BlockInflightRequest.ValueStringPointer())
	plan.AccountNameRule = util.SafeString(plan.AccountNameRule.ValueStringPointer())
	plan.AllowChangePasswordSQLQuery = util.SafeString(plan.AllowChangePasswordSQLQuery.ValueStringPointer())
	plan.AccountNameValidatorRegex = util.SafeString(plan.AccountNameValidatorRegex.ValueStringPointer())
	plan.PrimaryAccountType = util.SafeString(plan.PrimaryAccountType.ValueStringPointer())
	plan.AccountTypeNoPasswordChange = util.SafeString(plan.AccountTypeNoPasswordChange.ValueStringPointer())
	plan.ChangePasswordAccessQuery = util.SafeString(plan.ChangePasswordAccessQuery.ValueStringPointer())
	plan.StatusConfig = util.SafeString(plan.StatusConfig.ValueStringPointer())
	plan.PluginConfigs = util.SafeString(plan.PluginConfigs.ValueStringPointer())
	plan.EndpointConfig = util.SafeString(plan.EndpointConfig.ValueStringPointer())
	plan.CustomProperty1 = util.SafeString(plan.CustomProperty1.ValueStringPointer())
	plan.CustomProperty2 = util.SafeString(plan.CustomProperty2.ValueStringPointer())
	plan.CustomProperty3 = util.SafeString(plan.CustomProperty3.ValueStringPointer())
	plan.CustomProperty4 = util.SafeString(plan.CustomProperty4.ValueStringPointer())
	plan.CustomProperty5 = util.SafeString(plan.CustomProperty5.ValueStringPointer())
	plan.CustomProperty6 = util.SafeString(plan.CustomProperty6.ValueStringPointer())
	plan.CustomProperty7 = util.SafeString(plan.CustomProperty7.ValueStringPointer())
	plan.CustomProperty8 = util.SafeString(plan.CustomProperty8.ValueStringPointer())
	plan.CustomProperty9 = util.SafeString(plan.CustomProperty9.ValueStringPointer())
	plan.CustomProperty10 = util.SafeString(plan.CustomProperty10.ValueStringPointer())
	plan.CustomProperty11 = util.SafeString(plan.CustomProperty11.ValueStringPointer())
	plan.CustomProperty12 = util.SafeString(plan.CustomProperty12.ValueStringPointer())
	plan.CustomProperty13 = util.SafeString(plan.CustomProperty13.ValueStringPointer())
	plan.CustomProperty14 = util.SafeString(plan.CustomProperty14.ValueStringPointer())
	plan.CustomProperty15 = util.SafeString(plan.CustomProperty15.ValueStringPointer())
	plan.CustomProperty16 = util.SafeString(plan.CustomProperty16.ValueStringPointer())
	plan.CustomProperty17 = util.SafeString(plan.CustomProperty17.ValueStringPointer())
	plan.CustomProperty18 = util.SafeString(plan.CustomProperty18.ValueStringPointer())
	plan.CustomProperty19 = util.SafeString(plan.CustomProperty19.ValueStringPointer())
	plan.CustomProperty20 = util.SafeString(plan.CustomProperty20.ValueStringPointer())
	plan.CustomProperty21 = util.SafeString(plan.CustomProperty21.ValueStringPointer())
	plan.CustomProperty22 = util.SafeString(plan.CustomProperty22.ValueStringPointer())
	plan.CustomProperty23 = util.SafeString(plan.CustomProperty23.ValueStringPointer())
	plan.CustomProperty24 = util.SafeString(plan.CustomProperty24.ValueStringPointer())
	plan.CustomProperty25 = util.SafeString(plan.CustomProperty25.ValueStringPointer())
	plan.CustomProperty26 = util.SafeString(plan.CustomProperty26.ValueStringPointer())
	plan.CustomProperty27 = util.SafeString(plan.CustomProperty27.ValueStringPointer())
	plan.CustomProperty28 = util.SafeString(plan.CustomProperty28.ValueStringPointer())
	plan.CustomProperty29 = util.SafeString(plan.CustomProperty29.ValueStringPointer())
	plan.CustomProperty30 = util.SafeString(plan.CustomProperty30.ValueStringPointer())
	plan.CustomProperty31 = util.SafeString(plan.CustomProperty31.ValueStringPointer())
	plan.CustomProperty32 = util.SafeString(plan.CustomProperty32.ValueStringPointer())
	plan.CustomProperty33 = util.SafeString(plan.CustomProperty33.ValueStringPointer())
	plan.CustomProperty34 = util.SafeString(plan.CustomProperty34.ValueStringPointer())
	plan.CustomProperty35 = util.SafeString(plan.CustomProperty35.ValueStringPointer())
	plan.CustomProperty36 = util.SafeString(plan.CustomProperty36.ValueStringPointer())
	plan.CustomProperty37 = util.SafeString(plan.CustomProperty37.ValueStringPointer())
	plan.CustomProperty38 = util.SafeString(plan.CustomProperty38.ValueStringPointer())
	plan.CustomProperty39 = util.SafeString(plan.CustomProperty39.ValueStringPointer())
	plan.CustomProperty40 = util.SafeString(plan.CustomProperty40.ValueStringPointer())
	plan.CustomProperty41 = util.SafeString(plan.CustomProperty41.ValueStringPointer())
	plan.CustomProperty42 = util.SafeString(plan.CustomProperty42.ValueStringPointer())
	plan.CustomProperty43 = util.SafeString(plan.CustomProperty43.ValueStringPointer())
	plan.CustomProperty44 = util.SafeString(plan.CustomProperty44.ValueStringPointer())
	plan.CustomProperty45 = util.SafeString(plan.CustomProperty45.ValueStringPointer())
	plan.AccountCustomProperty1Label = util.SafeString(plan.AccountCustomProperty1Label.ValueStringPointer())
	plan.AccountCustomProperty2Label = util.SafeString(plan.AccountCustomProperty2Label.ValueStringPointer())
	plan.AccountCustomProperty3Label = util.SafeString(plan.AccountCustomProperty3Label.ValueStringPointer())
	plan.AccountCustomProperty4Label = util.SafeString(plan.AccountCustomProperty4Label.ValueStringPointer())
	plan.AccountCustomProperty5Label = util.SafeString(plan.AccountCustomProperty5Label.ValueStringPointer())
	plan.AccountCustomProperty6Label = util.SafeString(plan.AccountCustomProperty6Label.ValueStringPointer())
	plan.AccountCustomProperty7Label = util.SafeString(plan.AccountCustomProperty7Label.ValueStringPointer())
	plan.AccountCustomProperty8Label = util.SafeString(plan.AccountCustomProperty8Label.ValueStringPointer())
	plan.AccountCustomProperty9Label = util.SafeString(plan.AccountCustomProperty9Label.ValueStringPointer())
	plan.AccountCustomProperty10Label = util.SafeString(plan.AccountCustomProperty10Label.ValueStringPointer())
	plan.AccountCustomProperty11Label = util.SafeString(plan.AccountCustomProperty11Label.ValueStringPointer())
	plan.AccountCustomProperty12Label = util.SafeString(plan.AccountCustomProperty12Label.ValueStringPointer())
	plan.AccountCustomProperty13Label = util.SafeString(plan.AccountCustomProperty13Label.ValueStringPointer())
	plan.AccountCustomProperty14Label = util.SafeString(plan.AccountCustomProperty14Label.ValueStringPointer())
	plan.AccountCustomProperty15Label = util.SafeString(plan.AccountCustomProperty15Label.ValueStringPointer())
	plan.AccountCustomProperty16Label = util.SafeString(plan.AccountCustomProperty16Label.ValueStringPointer())
	plan.AccountCustomProperty17Label = util.SafeString(plan.AccountCustomProperty17Label.ValueStringPointer())
	plan.AccountCustomProperty18Label = util.SafeString(plan.AccountCustomProperty18Label.ValueStringPointer())
	plan.AccountCustomProperty19Label = util.SafeString(plan.AccountCustomProperty19Label.ValueStringPointer())
	plan.AccountCustomProperty20Label = util.SafeString(plan.AccountCustomProperty20Label.ValueStringPointer())
	plan.AccountCustomProperty21Label = util.SafeString(plan.AccountCustomProperty21Label.ValueStringPointer())
	plan.AccountCustomProperty22Label = util.SafeString(plan.AccountCustomProperty22Label.ValueStringPointer())
	plan.AccountCustomProperty23Label = util.SafeString(plan.AccountCustomProperty23Label.ValueStringPointer())
	plan.AccountCustomProperty24Label = util.SafeString(plan.AccountCustomProperty24Label.ValueStringPointer())
	plan.AccountCustomProperty25Label = util.SafeString(plan.AccountCustomProperty25Label.ValueStringPointer())
	plan.AccountCustomProperty26Label = util.SafeString(plan.AccountCustomProperty26Label.ValueStringPointer())
	plan.AccountCustomProperty27Label = util.SafeString(plan.AccountCustomProperty27Label.ValueStringPointer())
	plan.AccountCustomProperty28Label = util.SafeString(plan.AccountCustomProperty28Label.ValueStringPointer())
	plan.AccountCustomProperty29Label = util.SafeString(plan.AccountCustomProperty29Label.ValueStringPointer())
	plan.AccountCustomProperty30Label = util.SafeString(plan.AccountCustomProperty30Label.ValueStringPointer())
	plan.CustomProperty31Label = util.SafeString(plan.CustomProperty31Label.ValueStringPointer())
	plan.CustomProperty32Label = util.SafeString(plan.CustomProperty32Label.ValueStringPointer())
	plan.CustomProperty33Label = util.SafeString(plan.CustomProperty33Label.ValueStringPointer())
	plan.CustomProperty34Label = util.SafeString(plan.CustomProperty34Label.ValueStringPointer())
	plan.CustomProperty35Label = util.SafeString(plan.CustomProperty35Label.ValueStringPointer())
	plan.CustomProperty36Label = util.SafeString(plan.CustomProperty36Label.ValueStringPointer())
	plan.CustomProperty37Label = util.SafeString(plan.CustomProperty37Label.ValueStringPointer())
	plan.CustomProperty38Label = util.SafeString(plan.CustomProperty38Label.ValueStringPointer())
	plan.CustomProperty39Label = util.SafeString(plan.CustomProperty39Label.ValueStringPointer())
	plan.CustomProperty40Label = util.SafeString(plan.CustomProperty40Label.ValueStringPointer())
	plan.CustomProperty41Label = util.SafeString(plan.CustomProperty41Label.ValueStringPointer())
	plan.CustomProperty42Label = util.SafeString(plan.CustomProperty42Label.ValueStringPointer())
	plan.CustomProperty43Label = util.SafeString(plan.CustomProperty43Label.ValueStringPointer())
	plan.CustomProperty44Label = util.SafeString(plan.CustomProperty44Label.ValueStringPointer())
	plan.CustomProperty45Label = util.SafeString(plan.CustomProperty45Label.ValueStringPointer())
	plan.CustomProperty46Label = util.SafeString(plan.CustomProperty46Label.ValueStringPointer())
	plan.CustomProperty47Label = util.SafeString(plan.CustomProperty47Label.ValueStringPointer())
	plan.CustomProperty48Label = util.SafeString(plan.CustomProperty48Label.ValueStringPointer())
	plan.CustomProperty49Label = util.SafeString(plan.CustomProperty49Label.ValueStringPointer())
	plan.CustomProperty50Label = util.SafeString(plan.CustomProperty50Label.ValueStringPointer())
	plan.CustomProperty51Label = util.SafeString(plan.CustomProperty51Label.ValueStringPointer())
	plan.CustomProperty52Label = util.SafeString(plan.CustomProperty52Label.ValueStringPointer())
	plan.CustomProperty53Label = util.SafeString(plan.CustomProperty53Label.ValueStringPointer())
	plan.CustomProperty54Label = util.SafeString(plan.CustomProperty54Label.ValueStringPointer())
	plan.CustomProperty55Label = util.SafeString(plan.CustomProperty55Label.ValueStringPointer())
	plan.CustomProperty56Label = util.SafeString(plan.CustomProperty56Label.ValueStringPointer())
	plan.CustomProperty57Label = util.SafeString(plan.CustomProperty57Label.ValueStringPointer())
	plan.CustomProperty58Label = util.SafeString(plan.CustomProperty58Label.ValueStringPointer())
	plan.CustomProperty59Label = util.SafeString(plan.CustomProperty59Label.ValueStringPointer())
	plan.CustomProperty60Label = util.SafeString(plan.CustomProperty60Label.ValueStringPointer())

	msgValue := util.SafeDeref(apiResp.Msg)
	errorCodeValue := util.SafeDeref(apiResp.ErrorCode)
	plan.Msg = types.StringValue(msgValue)
	plan.ErrorCode = types.StringValue(errorCodeValue)

	stateCreateDiagnostics := resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(stateCreateDiagnostics...)
}

func (r *endpointResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state EndpointResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
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
	reqParams := openapi.GetEndpointsRequest{}
	reqParams.SetEndpointname(state.EndpointName.ValueString())
	readResp, _, err := apiClient.EndpointsAPI.GetEndpoints(ctx).GetEndpointsRequest(reqParams).Execute()
	if err != nil {
		log.Printf("Problem with the get function in read block")
		resp.Diagnostics.AddError("API Read Failed In Read Block", fmt.Sprintf("Error: %s", err))
		return
	}
	if readResp != nil && *readResp.ErrorCode != "0" {
		log.Printf("Error Reading Endpoint: %v, Error code: %v", *readResp.Message, *readResp.ErrorCode)
		resp.Diagnostics.AddError(
			"Error Reading Endpoint",
			fmt.Sprintf("Error: %v, Error code: %v", *readResp.Message, *readResp.ErrorCode),
		)
		return
	}
	if len(readResp.Endpoints) == 0 {
		resp.Diagnostics.AddError(
			"Client Error",
			"API returned empty endpoints list",
		)
		return
	}

	state.ID = types.StringValue("endpoint-" + *readResp.Endpoints[0].Endpointname)
	state.DisplayName = util.SafeString(readResp.Endpoints[0].DisplayName)
	state.SecuritySystem = util.SafeString(readResp.Endpoints[0].Securitysystem)
	state.Description = util.SafeString(readResp.Endpoints[0].Description)
	state.OwnerType = util.SafeString(util.StringPtr(endpointsutil.TranslateValue(util.SafeDeref(readResp.Endpoints[0].OwnerType), endpointsutil.OwnerTypeMap)))
	state.ResourceOwnerType = util.SafeString(util.StringPtr(endpointsutil.TranslateValue(util.SafeDeref(readResp.Endpoints[0].Requestownertype), endpointsutil.OwnerTypeMap)))
	state.PrimaryAccountType = util.SafeString(readResp.Endpoints[0].PrimaryAccountType)
	state.AccountTypeNoPasswordChange = util.SafeString(readResp.Endpoints[0].AccountTypeNoPasswordChange)
	state.ServiceAccountNameRule = util.SafeString(readResp.Endpoints[0].ServiceAccountNameRule)
	state.AccountNameValidatorRegex = util.SafeString(readResp.Endpoints[0].AccountNameValidatorRegex)
	state.AllowChangePasswordSQLQuery = util.SafeString(readResp.Endpoints[0].AllowChangePasswordSqlquery)
	state.ParentAccountPattern = util.SafeString(readResp.Endpoints[0].ParentAccountPattern)
	state.SecuritySystem = util.SafeString(readResp.Endpoints[0].Securitysystem)
	state.EndpointName = util.SafeString(readResp.Endpoints[0].Endpointname)
	state.AccessQuery = util.SafeString(readResp.Endpoints[0].Accessquery)
	state.DisplayName = util.SafeString(readResp.Endpoints[0].DisplayName)
	if readResp.Endpoints[0].Requestable != nil {
		if *readResp.Endpoints[0].Requestable == "true" {
			state.Requestable = types.BoolValue(true)
		} else {
			state.Requestable = types.BoolValue(false)
		}
	} else {
		state.Requestable = types.BoolNull()
	}
	if readResp.Endpoints[0].AllowRemoveAllRoleOnRequest != nil {
		if *readResp.Endpoints[0].AllowRemoveAllRoleOnRequest == "true" {
			state.AllowRemoveAllRoleOnRequest = types.BoolValue(true)
		} else {
			state.AllowRemoveAllRoleOnRequest = types.BoolValue(false)
		}
	} else {
		state.AllowRemoveAllRoleOnRequest = types.BoolNull()
	}
	if readResp.Endpoints[0].ConnectionconfigAsJson != nil {
		normalized, err := endpointsutil.NormalizeJSON(*readResp.Endpoints[0].ConnectionconfigAsJson)
		if err != nil {
			state.ConnectionConfig = types.StringNull()
		} else {
			state.ConnectionConfig = util.SafeString(&normalized)
		}
	} else {
		state.ConnectionConfig = types.StringNull()
	}
	state.AccountNameRule = util.SafeString(readResp.Endpoints[0].AccountNameRule)
	state.ChangePasswordAccessQuery = util.SafeString(readResp.Endpoints[0].ChangePasswordAccessQuery)
	state.PluginConfigs = util.SafeString(readResp.Endpoints[0].PluginConfigs)
	state.CreateEntTaskforRemoveAcc = util.SafeString(readResp.Endpoints[0].CreateEntTaskforRemoveAcc)
	state.EnableCopyAccess = util.SafeString(readResp.Endpoints[0].EnableCopyAccess)
	state.EndpointConfig = util.SafeString(readResp.Endpoints[0].EndpointConfig)
	state.ServiceAccountAccessQuery = util.SafeString(readResp.Endpoints[0].ServiceAccountAccessQuery)
	state.UserAccountCorrelationRule = util.SafeString(readResp.Endpoints[0].UserAccountCorrelationRule)
	state.StatusConfig = util.SafeString(readResp.Endpoints[0].StatusConfig)
	state.CustomProperty1 = util.SafeString(readResp.Endpoints[0].CustomProperty1)
	state.CustomProperty2 = util.SafeString(readResp.Endpoints[0].CustomProperty2)
	state.CustomProperty3 = util.SafeString(readResp.Endpoints[0].CustomProperty3)
	state.CustomProperty4 = util.SafeString(readResp.Endpoints[0].CustomProperty4)
	state.CustomProperty5 = util.SafeString(readResp.Endpoints[0].CustomProperty5)
	state.CustomProperty6 = util.SafeString(readResp.Endpoints[0].CustomProperty6)
	state.CustomProperty7 = util.SafeString(readResp.Endpoints[0].CustomProperty7)
	state.CustomProperty8 = util.SafeString(readResp.Endpoints[0].CustomProperty8)
	state.CustomProperty9 = util.SafeString(readResp.Endpoints[0].CustomProperty9)
	state.CustomProperty10 = util.SafeString(readResp.Endpoints[0].CustomProperty10)
	state.CustomProperty11 = util.SafeString(readResp.Endpoints[0].CustomProperty11)
	state.CustomProperty12 = util.SafeString(readResp.Endpoints[0].CustomProperty12)
	state.CustomProperty13 = util.SafeString(readResp.Endpoints[0].CustomProperty13)
	state.CustomProperty14 = util.SafeString(readResp.Endpoints[0].CustomProperty14)
	state.CustomProperty15 = util.SafeString(readResp.Endpoints[0].CustomProperty15)
	state.CustomProperty16 = util.SafeString(readResp.Endpoints[0].CustomProperty16)
	state.CustomProperty17 = util.SafeString(readResp.Endpoints[0].CustomProperty17)
	state.CustomProperty18 = util.SafeString(readResp.Endpoints[0].CustomProperty18)
	state.CustomProperty19 = util.SafeString(readResp.Endpoints[0].CustomProperty19)
	state.CustomProperty20 = util.SafeString(readResp.Endpoints[0].CustomProperty20)
	state.CustomProperty21 = util.SafeString(readResp.Endpoints[0].CustomProperty21)
	state.CustomProperty22 = util.SafeString(readResp.Endpoints[0].CustomProperty22)
	state.CustomProperty23 = util.SafeString(readResp.Endpoints[0].CustomProperty23)
	state.CustomProperty24 = util.SafeString(readResp.Endpoints[0].CustomProperty24)
	state.CustomProperty25 = util.SafeString(readResp.Endpoints[0].CustomProperty25)
	state.CustomProperty26 = util.SafeString(readResp.Endpoints[0].CustomProperty26)
	state.CustomProperty27 = util.SafeString(readResp.Endpoints[0].CustomProperty27)
	state.CustomProperty28 = util.SafeString(readResp.Endpoints[0].CustomProperty28)
	state.CustomProperty29 = util.SafeString(readResp.Endpoints[0].CustomProperty29)
	state.CustomProperty30 = util.SafeString(readResp.Endpoints[0].CustomProperty30)
	state.CustomProperty31 = util.SafeString(readResp.Endpoints[0].Customproperty31)
	state.CustomProperty32 = util.SafeString(readResp.Endpoints[0].Customproperty32)
	state.CustomProperty33 = util.SafeString(readResp.Endpoints[0].Customproperty33)
	state.CustomProperty34 = util.SafeString(readResp.Endpoints[0].Customproperty34)
	state.CustomProperty35 = util.SafeString(readResp.Endpoints[0].Customproperty35)
	state.CustomProperty36 = util.SafeString(readResp.Endpoints[0].Customproperty36)
	state.CustomProperty37 = util.SafeString(readResp.Endpoints[0].Customproperty37)
	state.CustomProperty38 = util.SafeString(readResp.Endpoints[0].Customproperty38)
	state.CustomProperty39 = util.SafeString(readResp.Endpoints[0].Customproperty39)
	state.CustomProperty40 = util.SafeString(readResp.Endpoints[0].Customproperty40)
	state.CustomProperty41 = util.SafeString(readResp.Endpoints[0].Customproperty41)
	state.CustomProperty42 = util.SafeString(readResp.Endpoints[0].Customproperty42)
	state.CustomProperty43 = util.SafeString(readResp.Endpoints[0].Customproperty43)
	state.CustomProperty44 = util.SafeString(readResp.Endpoints[0].Customproperty44)
	state.CustomProperty45 = util.SafeString(readResp.Endpoints[0].Customproperty45)
	state.AccountCustomProperty1Label = util.SafeString(readResp.Endpoints[0].AccountCustomProperty1Label)
	state.AccountCustomProperty2Label = util.SafeString(readResp.Endpoints[0].AccountCustomProperty2Label)
	state.AccountCustomProperty3Label = util.SafeString(readResp.Endpoints[0].AccountCustomProperty3Label)
	state.AccountCustomProperty4Label = util.SafeString(readResp.Endpoints[0].AccountCustomProperty4Label)
	state.AccountCustomProperty5Label = util.SafeString(readResp.Endpoints[0].AccountCustomProperty5Label)
	state.AccountCustomProperty6Label = util.SafeString(readResp.Endpoints[0].AccountCustomProperty6Label)
	state.AccountCustomProperty7Label = util.SafeString(readResp.Endpoints[0].AccountCustomProperty7Label)
	state.AccountCustomProperty8Label = util.SafeString(readResp.Endpoints[0].AccountCustomProperty8Label)
	state.AccountCustomProperty9Label = util.SafeString(readResp.Endpoints[0].AccountCustomProperty9Label)
	state.AccountCustomProperty10Label = util.SafeString(readResp.Endpoints[0].AccountCustomProperty10Label)
	state.AccountCustomProperty11Label = util.SafeString(readResp.Endpoints[0].AccountCustomProperty11Label)
	state.AccountCustomProperty12Label = util.SafeString(readResp.Endpoints[0].AccountCustomProperty12Label)
	state.AccountCustomProperty13Label = util.SafeString(readResp.Endpoints[0].AccountCustomProperty13Label)
	state.AccountCustomProperty14Label = util.SafeString(readResp.Endpoints[0].AccountCustomProperty14Label)
	state.AccountCustomProperty15Label = util.SafeString(readResp.Endpoints[0].AccountCustomProperty15Label)
	state.AccountCustomProperty16Label = util.SafeString(readResp.Endpoints[0].AccountCustomProperty16Label)
	state.AccountCustomProperty17Label = util.SafeString(readResp.Endpoints[0].AccountCustomProperty17Label)
	state.AccountCustomProperty18Label = util.SafeString(readResp.Endpoints[0].AccountCustomProperty18Label)
	state.AccountCustomProperty19Label = util.SafeString(readResp.Endpoints[0].AccountCustomProperty19Label)
	state.AccountCustomProperty20Label = util.SafeString(readResp.Endpoints[0].AccountCustomProperty20Label)
	state.AccountCustomProperty21Label = util.SafeString(readResp.Endpoints[0].AccountCustomProperty21Label)
	state.AccountCustomProperty22Label = util.SafeString(readResp.Endpoints[0].AccountCustomProperty22Label)
	state.AccountCustomProperty23Label = util.SafeString(readResp.Endpoints[0].AccountCustomProperty23Label)
	state.AccountCustomProperty24Label = util.SafeString(readResp.Endpoints[0].AccountCustomProperty24Label)
	state.AccountCustomProperty25Label = util.SafeString(readResp.Endpoints[0].AccountCustomProperty25Label)
	state.AccountCustomProperty26Label = util.SafeString(readResp.Endpoints[0].AccountCustomProperty26Label)
	state.AccountCustomProperty27Label = util.SafeString(readResp.Endpoints[0].AccountCustomProperty27Label)
	state.AccountCustomProperty28Label = util.SafeString(readResp.Endpoints[0].AccountCustomProperty28Label)
	state.AccountCustomProperty29Label = util.SafeString(readResp.Endpoints[0].AccountCustomProperty29Label)
	state.AccountCustomProperty30Label = util.SafeString(readResp.Endpoints[0].AccountCustomProperty30Label)
	state.CustomProperty31Label = util.SafeString(readResp.Endpoints[0].Customproperty31Label)
	state.CustomProperty32Label = util.SafeString(readResp.Endpoints[0].Customproperty32Label)
	state.CustomProperty33Label = util.SafeString(readResp.Endpoints[0].Customproperty33Label)
	state.CustomProperty34Label = util.SafeString(readResp.Endpoints[0].Customproperty34Label)
	state.CustomProperty35Label = util.SafeString(readResp.Endpoints[0].Customproperty35Label)
	state.CustomProperty36Label = util.SafeString(readResp.Endpoints[0].Customproperty36Label)
	state.CustomProperty37Label = util.SafeString(readResp.Endpoints[0].Customproperty37Label)
	state.CustomProperty38Label = util.SafeString(readResp.Endpoints[0].Customproperty38Label)
	state.CustomProperty39Label = util.SafeString(readResp.Endpoints[0].Customproperty39Label)
	state.CustomProperty40Label = util.SafeString(readResp.Endpoints[0].Customproperty40Label)
	state.CustomProperty41Label = util.SafeString(readResp.Endpoints[0].Customproperty41Label)
	state.CustomProperty42Label = util.SafeString(readResp.Endpoints[0].Customproperty42Label)
	state.CustomProperty43Label = util.SafeString(readResp.Endpoints[0].Customproperty43Label)
	state.CustomProperty44Label = util.SafeString(readResp.Endpoints[0].Customproperty44Label)
	state.CustomProperty45Label = util.SafeString(readResp.Endpoints[0].Customproperty45Label)
	state.CustomProperty46Label = util.SafeString(readResp.Endpoints[0].Customproperty46Label)
	state.CustomProperty47Label = util.SafeString(readResp.Endpoints[0].Customproperty47Label)
	state.CustomProperty48Label = util.SafeString(readResp.Endpoints[0].Customproperty48Label)
	state.CustomProperty49Label = util.SafeString(readResp.Endpoints[0].Customproperty49Label)
	state.CustomProperty50Label = util.SafeString(readResp.Endpoints[0].Customproperty50Label)
	state.CustomProperty51Label = util.SafeString(readResp.Endpoints[0].Customproperty51Label)
	state.CustomProperty52Label = util.SafeString(readResp.Endpoints[0].Customproperty52Label)
	state.CustomProperty53Label = util.SafeString(readResp.Endpoints[0].Customproperty53Label)
	state.CustomProperty54Label = util.SafeString(readResp.Endpoints[0].Customproperty54Label)
	state.CustomProperty55Label = util.SafeString(readResp.Endpoints[0].Customproperty55Label)
	state.CustomProperty56Label = util.SafeString(readResp.Endpoints[0].Customproperty56Label)
	state.CustomProperty57Label = util.SafeString(readResp.Endpoints[0].Customproperty57Label)
	state.CustomProperty58Label = util.SafeString(readResp.Endpoints[0].Customproperty58Label)
	state.CustomProperty59Label = util.SafeString(readResp.Endpoints[0].Customproperty59Label)
	state.CustomProperty60Label = util.SafeString(readResp.Endpoints[0].Customproperty60Label)

	var stateEmailTemplates types.List

	if readResp.Endpoints[0].Taskemailtemplates == nil || *readResp.Endpoints[0].Taskemailtemplates == "" {
		stateEmailTemplates = types.ListNull(types.ObjectType{
			AttrTypes: map[string]attr.Type{
				"email_template_type": types.StringType,
				"task_type":           types.StringType,
				"email_template":      types.StringType,
			},
		})
	} else {
		taskEmailTemplatesStr := *readResp.Endpoints[0].Taskemailtemplates
		type ApiEmailTemplate struct {
			EmailTemplateType string `json:"emailTemplateType"`
			TaskType          string `json:"taskType"`
			EmailTemplate     string `json:"emailTemplate"`
		}

		var apiTemplates []ApiEmailTemplate
		if err := json.Unmarshal([]byte(taskEmailTemplatesStr), &apiTemplates); err != nil {
			resp.Diagnostics.AddError(
				"Client Error",
				fmt.Sprintf("Failed to parse taskemailtemplates JSON: %s", err),
			)
			return
		}

		emailTemplateObjects := make([]attr.Value, 0, len(apiTemplates))
		for _, t := range apiTemplates {
			obj, diags := types.ObjectValue(
				map[string]attr.Type{
					"email_template_type": types.StringType,
					"task_type":           types.StringType,
					"email_template":      types.StringType,
				},
				map[string]attr.Value{
					"email_template_type": types.StringValue(t.EmailTemplateType),
					"task_type":           types.StringValue(t.TaskType),
					"email_template":      types.StringValue(t.EmailTemplate),
				},
			)
			if diags.HasError() {
				resp.Diagnostics.Append(diags...)
				continue
			}
			emailTemplateObjects = append(emailTemplateObjects, obj)
		}

		listVal, diags := types.ListValue(
			types.ObjectType{
				AttrTypes: map[string]attr.Type{
					"email_template_type": types.StringType,
					"task_type":           types.StringType,
					"email_template":      types.StringType,
				},
			},
			emailTemplateObjects,
		)
		if diags.HasError() {
			resp.Diagnostics.Append(diags...)
			return
		}
		stateEmailTemplates = listVal
	}

	state.EmailTemplates = stateEmailTemplates

	var stateRequestableRoleTypes types.List

	if readResp.Endpoints[0].RoleTypeAsJson == nil || *readResp.Endpoints[0].RoleTypeAsJson == "" {
		stateRequestableRoleTypes = types.ListNull(types.ObjectType{
			AttrTypes: map[string]attr.Type{
				"role_type":       types.StringType,
				"request_option":  types.StringType,
				"required":        types.BoolType,
				"requested_query": types.StringType,
				"selected_query":  types.StringType,
				"show_on":         types.StringType,
			},
		})
	} else {
		roleTypeJsonStr := *readResp.Endpoints[0].RoleTypeAsJson
		var roleTypeMap map[string]string
		if err := json.Unmarshal([]byte(roleTypeJsonStr), &roleTypeMap); err != nil {
			resp.Diagnostics.AddError(
				"Client Error",
				fmt.Sprintf("Failed to parse RoleTypeAsJson outer JSON: %s", err),
			)
			return
		}

		roleTypeObjects := make([]attr.Value, 0, len(roleTypeMap))
		fmt.Println("Role type objects: ", roleTypeObjects)

		for roleType, roleData := range roleTypeMap {
			parts := strings.Split(roleData, "__")

			get := func(i int) string {
				if i < len(parts) {
					return parts[i]
				}
				return ""
			}

			obj, diags := types.ObjectValue(
				map[string]attr.Type{
					"role_type":       types.StringType,
					"request_option":  types.StringType,
					"required":        types.BoolType,
					"requested_query": types.StringType,
					"selected_query":  types.StringType,
					"show_on":         types.StringType,
				},
				map[string]attr.Value{
					"role_type":       types.StringValue(endpointsutil.TranslateValue(roleType, endpointsutil.RoleTypeMap)),
					"request_option":  types.StringValue(endpointsutil.TranslateValue(get(0), endpointsutil.RequestOptionMap)),
					"required":        types.BoolValue(get(1) == "1"),
					"requested_query": types.StringValue(get(2)),
					"selected_query":  types.StringValue(get(3)),
					"show_on":         types.StringValue(endpointsutil.TranslateValue(get(4), endpointsutil.ShowOnMap)),
				},
			)
			if diags.HasError() {
				resp.Diagnostics.Append(diags...)
				continue
			}
			roleTypeObjects = append(roleTypeObjects, obj)
		}

		listVal, diags := types.ListValue(
			types.ObjectType{
				AttrTypes: map[string]attr.Type{
					"role_type":       types.StringType,
					"request_option":  types.StringType,
					"required":        types.BoolType,
					"requested_query": types.StringType,
					"selected_query":  types.StringType,
					"show_on":         types.StringType,
				},
			},
			roleTypeObjects,
		)
		if diags.HasError() {
			resp.Diagnostics.Append(diags...)
			return
		}
		stateRequestableRoleTypes = listVal
	}

	state.RequestableRoleTypes = stateRequestableRoleTypes

	if readResp.Endpoints[0].Disableaccountrequest != nil {
		disableAccountRequestStr := *readResp.Endpoints[0].Disableaccountrequest
		var disableAccountRequestMap map[string]string

		err := json.Unmarshal([]byte(disableAccountRequestStr), &disableAccountRequestMap)
		if err != nil {
			log.Printf("Error parsing disableaccountrequest JSON: %v", err)
		} else {
			state.DisableNewAccountRequestIfAccountExists = types.StringValue(endpointsutil.NormalizeToStringBool(disableAccountRequestMap["DISABLENEWACCOUNT"]))
			state.DisableRemoveAccount = types.StringValue(endpointsutil.NormalizeToStringBool(disableAccountRequestMap["DISABLEREMOVEACCOUNT"]))
			state.DisableModifyAccount = types.StringValue(endpointsutil.NormalizeToStringBool(disableAccountRequestMap["DISABLEMODIFYACCOUNT"]))
			state.BlockInflightRequest = types.StringValue(endpointsutil.NormalizeToStringBool(disableAccountRequestMap["BLOCKINFLIGHTREQUEST"]))

		}
	}

	msgValue := util.SafeDeref(readResp.Message)
	state.Msg = util.SafeString(&msgValue)
	state.ErrorCode = util.SafeString(readResp.ErrorCode)
	stateDiagnostics := resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(stateDiagnostics...)

	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *endpointResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan EndpointResourceModel
	var state EndpointResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	planGetDiagnostics := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(planGetDiagnostics...)
	if resp.Diagnostics.HasError() {
		return
	}

	if plan.EndpointName.ValueString() != state.EndpointName.ValueString() {
		resp.Diagnostics.AddError("Error", "Endpoint name cannot be updated")
		log.Printf("[ERROR]: Endpoint name cannot be updated")
		return
	}

	cfg := openapi.NewConfiguration()
	apiBaseURL := strings.TrimPrefix(strings.TrimPrefix(r.client.APIBaseURL(), "https://"), "http://")
	cfg.Host = apiBaseURL
	cfg.Scheme = "https"
	cfg.AddDefaultHeader("Authorization", "Bearer "+r.token)
	cfg.HTTPClient = http.DefaultClient

	apiClient := openapi.NewAPIClient(cfg)

	updateReq := openapi.UpdateEndpointRequest{
		Endpointname:                            plan.EndpointName.ValueString(),
		PrimaryAccountType:                      plan.PrimaryAccountType.ValueStringPointer(),
		AccountTypeNoPasswordChange:             plan.AccountTypeNoPasswordChange.ValueStringPointer(),
		DisplayName:                             plan.DisplayName.ValueStringPointer(),
		Securitysystem:                          plan.SecuritySystem.ValueStringPointer(),
		Description:                             plan.Description.ValueStringPointer(),
		OwnerType:                               plan.OwnerType.ValueStringPointer(),
		Owner:                                   plan.Owner.ValueStringPointer(),
		ResourceOwnerType:                       plan.ResourceOwnerType.ValueStringPointer(),
		ResourceOwner:                           plan.ResourceOwner.ValueStringPointer(),
		Accessquery:                             plan.AccessQuery.ValueStringPointer(),
		EnableCopyAccess:                        plan.EnableCopyAccess.ValueStringPointer(),
		CreateEntTaskforRemoveAcc:               plan.CreateEntTaskforRemoveAcc.ValueStringPointer(),
		DisableNewAccountRequestIfAccountExists: plan.DisableNewAccountRequestIfAccountExists.ValueStringPointer(),
		DisableRemoveAccount:                    plan.DisableRemoveAccount.ValueStringPointer(),
		DisableModifyAccount:                    plan.DisableModifyAccount.ValueStringPointer(),
		UserAccountCorrelationRule:              plan.UserAccountCorrelationRule.ValueStringPointer(),
		Connectionconfig:                        plan.ConnectionConfig.ValueStringPointer(),
		BlockInflightRequest:                    plan.BlockInflightRequest.ValueStringPointer(),
		Outofbandaction:                         plan.OutOfBandAction.ValueStringPointer(),
		AccountNameRule:                         plan.AccountNameRule.ValueStringPointer(),
		AllowChangePasswordSqlquery:             plan.AllowChangePasswordSQLQuery.ValueStringPointer(),
		Requestable:                             plan.Requestable.ValueBoolPointer(),
		ParentAccountPattern:                    plan.ParentAccountPattern.ValueStringPointer(),
		ServiceAccountNameRule:                  plan.ServiceAccountNameRule.ValueStringPointer(),
		ServiceAccountAccessQuery:               plan.ServiceAccountAccessQuery.ValueStringPointer(),
		ChangePasswordAccessQuery:               plan.ChangePasswordAccessQuery.ValueStringPointer(),
		AccountNameValidatorRegex:               plan.AccountNameValidatorRegex.ValueStringPointer(),
		StatusConfig:                            plan.StatusConfig.ValueStringPointer(),
		PluginConfigs:                           plan.PluginConfigs.ValueStringPointer(),
		EndpointConfig:                          plan.EndpointConfig.ValueStringPointer(),
		AllowRemoveAllRoleOnRequest:             plan.AllowRemoveAllRoleOnRequest.ValueBoolPointer(),
		Customproperty1:                         util.StringPointerOrEmpty(plan.CustomProperty1),
		Customproperty2:                         util.StringPointerOrEmpty(plan.CustomProperty2),
		Customproperty3:                         util.StringPointerOrEmpty(plan.CustomProperty3),
		Customproperty4:                         util.StringPointerOrEmpty(plan.CustomProperty4),
		Customproperty5:                         util.StringPointerOrEmpty(plan.CustomProperty5),
		Customproperty6:                         util.StringPointerOrEmpty(plan.CustomProperty6),
		Customproperty7:                         util.StringPointerOrEmpty(plan.CustomProperty7),
		Customproperty8:                         util.StringPointerOrEmpty(plan.CustomProperty8),
		Customproperty9:                         util.StringPointerOrEmpty(plan.CustomProperty9),
		Customproperty10:                        util.StringPointerOrEmpty(plan.CustomProperty10),
		Customproperty11:                        util.StringPointerOrEmpty(plan.CustomProperty11),
		Customproperty12:                        util.StringPointerOrEmpty(plan.CustomProperty12),
		Customproperty13:                        util.StringPointerOrEmpty(plan.CustomProperty13),
		Customproperty14:                        util.StringPointerOrEmpty(plan.CustomProperty14),
		Customproperty15:                        util.StringPointerOrEmpty(plan.CustomProperty15),
		Customproperty16:                        util.StringPointerOrEmpty(plan.CustomProperty16),
		Customproperty17:                        util.StringPointerOrEmpty(plan.CustomProperty17),
		Customproperty18:                        util.StringPointerOrEmpty(plan.CustomProperty18),
		Customproperty19:                        util.StringPointerOrEmpty(plan.CustomProperty19),
		Customproperty20:                        util.StringPointerOrEmpty(plan.CustomProperty20),
		Customproperty21:                        util.StringPointerOrEmpty(plan.CustomProperty21),
		Customproperty22:                        util.StringPointerOrEmpty(plan.CustomProperty22),
		Customproperty23:                        util.StringPointerOrEmpty(plan.CustomProperty23),
		Customproperty24:                        util.StringPointerOrEmpty(plan.CustomProperty24),
		Customproperty25:                        util.StringPointerOrEmpty(plan.CustomProperty25),
		Customproperty26:                        util.StringPointerOrEmpty(plan.CustomProperty26),
		Customproperty27:                        util.StringPointerOrEmpty(plan.CustomProperty27),
		Customproperty28:                        util.StringPointerOrEmpty(plan.CustomProperty28),
		Customproperty29:                        util.StringPointerOrEmpty(plan.CustomProperty29),
		Customproperty30:                        util.StringPointerOrEmpty(plan.CustomProperty30),
		Customproperty31:                        util.StringPointerOrEmpty(plan.CustomProperty31),
		Customproperty32:                        util.StringPointerOrEmpty(plan.CustomProperty32),
		Customproperty33:                        util.StringPointerOrEmpty(plan.CustomProperty33),
		Customproperty34:                        util.StringPointerOrEmpty(plan.CustomProperty34),
		Customproperty35:                        util.StringPointerOrEmpty(plan.CustomProperty35),
		Customproperty36:                        util.StringPointerOrEmpty(plan.CustomProperty36),
		Customproperty37:                        util.StringPointerOrEmpty(plan.CustomProperty37),
		Customproperty38:                        util.StringPointerOrEmpty(plan.CustomProperty38),
		Customproperty39:                        util.StringPointerOrEmpty(plan.CustomProperty39),
		Customproperty40:                        util.StringPointerOrEmpty(plan.CustomProperty40),
		Customproperty41:                        util.StringPointerOrEmpty(plan.CustomProperty41),
		Customproperty42:                        util.StringPointerOrEmpty(plan.CustomProperty42),
		Customproperty43:                        util.StringPointerOrEmpty(plan.CustomProperty43),
		Customproperty44:                        util.StringPointerOrEmpty(plan.CustomProperty44),
		Customproperty45:                        util.StringPointerOrEmpty(plan.CustomProperty45),
		Customproperty1Label:                    util.StringPointerOrEmpty(plan.AccountCustomProperty1Label),
		Customproperty2Label:                    util.StringPointerOrEmpty(plan.AccountCustomProperty2Label),
		Customproperty3Label:                    util.StringPointerOrEmpty(plan.AccountCustomProperty3Label),
		Customproperty4Label:                    util.StringPointerOrEmpty(plan.AccountCustomProperty4Label),
		Customproperty5Label:                    util.StringPointerOrEmpty(plan.AccountCustomProperty5Label),
		Customproperty6Label:                    util.StringPointerOrEmpty(plan.AccountCustomProperty6Label),
		Customproperty7Label:                    util.StringPointerOrEmpty(plan.AccountCustomProperty7Label),
		Customproperty8Label:                    util.StringPointerOrEmpty(plan.AccountCustomProperty8Label),
		Customproperty9Label:                    util.StringPointerOrEmpty(plan.AccountCustomProperty9Label),
		Customproperty10Label:                   util.StringPointerOrEmpty(plan.AccountCustomProperty10Label),
		Customproperty11Label:                   util.StringPointerOrEmpty(plan.AccountCustomProperty11Label),
		Customproperty12Label:                   util.StringPointerOrEmpty(plan.AccountCustomProperty12Label),
		Customproperty13Label:                   util.StringPointerOrEmpty(plan.AccountCustomProperty13Label),
		Customproperty14Label:                   util.StringPointerOrEmpty(plan.AccountCustomProperty14Label),
		Customproperty15Label:                   util.StringPointerOrEmpty(plan.AccountCustomProperty15Label),
		Customproperty16Label:                   util.StringPointerOrEmpty(plan.AccountCustomProperty16Label),
		Customproperty17Label:                   util.StringPointerOrEmpty(plan.AccountCustomProperty17Label),
		Customproperty18Label:                   util.StringPointerOrEmpty(plan.AccountCustomProperty18Label),
		Customproperty19Label:                   util.StringPointerOrEmpty(plan.AccountCustomProperty19Label),
		Customproperty20Label:                   util.StringPointerOrEmpty(plan.AccountCustomProperty20Label),
		Customproperty21Label:                   util.StringPointerOrEmpty(plan.AccountCustomProperty21Label),
		Customproperty22Label:                   util.StringPointerOrEmpty(plan.AccountCustomProperty22Label),
		Customproperty23Label:                   util.StringPointerOrEmpty(plan.AccountCustomProperty23Label),
		Customproperty24Label:                   util.StringPointerOrEmpty(plan.AccountCustomProperty24Label),
		Customproperty25Label:                   util.StringPointerOrEmpty(plan.AccountCustomProperty25Label),
		Customproperty26Label:                   util.StringPointerOrEmpty(plan.AccountCustomProperty26Label),
		Customproperty27Label:                   util.StringPointerOrEmpty(plan.AccountCustomProperty27Label),
		Customproperty28Label:                   util.StringPointerOrEmpty(plan.AccountCustomProperty28Label),
		Customproperty29Label:                   util.StringPointerOrEmpty(plan.AccountCustomProperty29Label),
		Customproperty30Label:                   util.StringPointerOrEmpty(plan.AccountCustomProperty30Label),
		Customproperty31Label:                   util.StringPointerOrEmpty(plan.CustomProperty31Label),
		Customproperty32Label:                   util.StringPointerOrEmpty(plan.CustomProperty32Label),
		Customproperty33Label:                   util.StringPointerOrEmpty(plan.CustomProperty33Label),
		Customproperty34Label:                   util.StringPointerOrEmpty(plan.CustomProperty34Label),
		Customproperty35Label:                   util.StringPointerOrEmpty(plan.CustomProperty35Label),
		Customproperty36Label:                   util.StringPointerOrEmpty(plan.CustomProperty36Label),
		Customproperty37Label:                   util.StringPointerOrEmpty(plan.CustomProperty37Label),
		Customproperty38Label:                   util.StringPointerOrEmpty(plan.CustomProperty38Label),
		Customproperty39Label:                   util.StringPointerOrEmpty(plan.CustomProperty39Label),
		Customproperty40Label:                   util.StringPointerOrEmpty(plan.CustomProperty40Label),
		Customproperty41Label:                   util.StringPointerOrEmpty(plan.CustomProperty41Label),
		Customproperty42Label:                   util.StringPointerOrEmpty(plan.CustomProperty42Label),
		Customproperty43Label:                   util.StringPointerOrEmpty(plan.CustomProperty43Label),
		Customproperty44Label:                   util.StringPointerOrEmpty(plan.CustomProperty44Label),
		Customproperty45Label:                   util.StringPointerOrEmpty(plan.CustomProperty45Label),
		Customproperty46Label:                   util.StringPointerOrEmpty(plan.CustomProperty46Label),
		Customproperty47Label:                   util.StringPointerOrEmpty(plan.CustomProperty47Label),
		Customproperty48Label:                   util.StringPointerOrEmpty(plan.CustomProperty48Label),
		Customproperty49Label:                   util.StringPointerOrEmpty(plan.CustomProperty49Label),
		Customproperty50Label:                   util.StringPointerOrEmpty(plan.CustomProperty50Label),
		Customproperty51Label:                   util.StringPointerOrEmpty(plan.CustomProperty51Label),
		Customproperty52Label:                   util.StringPointerOrEmpty(plan.CustomProperty52Label),
		Customproperty53Label:                   util.StringPointerOrEmpty(plan.CustomProperty53Label),
		Customproperty54Label:                   util.StringPointerOrEmpty(plan.CustomProperty54Label),
		Customproperty55Label:                   util.StringPointerOrEmpty(plan.CustomProperty55Label),
		Customproperty56Label:                   util.StringPointerOrEmpty(plan.CustomProperty56Label),
		Customproperty57Label:                   util.StringPointerOrEmpty(plan.CustomProperty57Label),
		Customproperty58Label:                   util.StringPointerOrEmpty(plan.CustomProperty58Label),
		Customproperty59Label:                   util.StringPointerOrEmpty(plan.CustomProperty59Label),
		Customproperty60Label:                   util.StringPointerOrEmpty(plan.CustomProperty60Label),
	}

	var mappedEndpoints []openapi.UpdateEndpointRequestMappedEndpointsInner
	var diags diag.Diagnostics
	var tfMappedEndpoints []MappedEndpoint

	diags = plan.MappedEndpoints.ElementsAs(ctx, &tfMappedEndpoints, true)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	for _, tfTemplate := range tfMappedEndpoints {
		if tfTemplate.SecuritySystem.IsUnknown() &&
			tfTemplate.Endpoint.IsUnknown() &&
			tfTemplate.Requestable.IsUnknown() &&
			tfTemplate.Operation.IsUnknown() {
			continue
		}

		mappedEndpoint := openapi.UpdateEndpointRequestMappedEndpointsInner{}

		if !tfTemplate.SecuritySystem.IsNull() {
			mappedEndpoint.Securitysystem = tfTemplate.SecuritySystem.ValueStringPointer()
		}
		if !tfTemplate.Endpoint.IsNull() {
			mappedEndpoint.Endpoint = tfTemplate.Endpoint.ValueStringPointer()
		}
		if !tfTemplate.Requestable.IsNull() {
			mappedEndpoint.Requestable = tfTemplate.Requestable.ValueStringPointer()
		}
		if !tfTemplate.Operation.IsNull() {
			mappedEndpoint.Operation = tfTemplate.Operation.ValueStringPointer()
		}

		mappedEndpoints = append(mappedEndpoints, mappedEndpoint)
	}

	if len(mappedEndpoints) > 0 {
		updateReq.MappedEndpoints = mappedEndpoints
	}

	var emailTemplates []openapi.CreateEndpointRequestEmailTemplateInner
	var tfEmailTemplates []EmailTemplate

	diags = plan.EmailTemplates.ElementsAs(ctx, &tfEmailTemplates, true)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	for _, tfTemplate := range tfEmailTemplates {
		if tfTemplate.EmailTemplateType.IsUnknown() &&
			tfTemplate.TaskType.IsUnknown() &&
			tfTemplate.EmailTemplate.IsUnknown() {
			continue
		}

		emailTemplate := openapi.CreateEndpointRequestEmailTemplateInner{}

		if !tfTemplate.EmailTemplateType.IsNull() {
			emailTemplate.EmailTemplateType = tfTemplate.EmailTemplateType.ValueStringPointer()
		}
		if !tfTemplate.TaskType.IsNull() {
			emailTemplate.TaskType = tfTemplate.TaskType.ValueStringPointer()
		}
		if !tfTemplate.EmailTemplate.IsNull() {
			emailTemplate.EmailTemplate = tfTemplate.EmailTemplate.ValueStringPointer()
		}

		emailTemplates = append(emailTemplates, emailTemplate)
	}

	if len(emailTemplates) > 0 {
		updateReq.Taskemailtemplates = emailTemplates
	}

	var requestableRoleTypes []openapi.UpdateEndpointRequestRequestableRoleTypeInner
	var tfRequestableRoleTypes []RequestableRoleType

	diags = plan.RequestableRoleTypes.ElementsAs(ctx, &tfRequestableRoleTypes, true)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	for _, tfTemplate := range tfRequestableRoleTypes {
		if tfTemplate.RoleType.IsUnknown() &&
			tfTemplate.RequestOption.IsUnknown() &&
			tfTemplate.RequestedQuery.IsUnknown() &&
			tfTemplate.Required.IsUnknown() &&
			tfTemplate.SelectedQuery.IsUnknown() &&
			tfTemplate.ShowOn.IsUnknown() {
			continue
		}

		requestableRoleType := openapi.UpdateEndpointRequestRequestableRoleTypeInner{}

		if !tfTemplate.RoleType.IsNull() {
			requestableRoleType.RoleType = tfTemplate.RoleType.ValueStringPointer()
		}
		if !tfTemplate.RequestOption.IsNull() {
			requestableRoleType.RequestOption = tfTemplate.RequestOption.ValueStringPointer()
		}
		if !tfTemplate.RequestedQuery.IsNull() {
			requestableRoleType.RequestedQuery = tfTemplate.RequestedQuery.ValueStringPointer()
		}
		if !tfTemplate.Required.IsNull() {
			requestableRoleType.Required = tfTemplate.Required.ValueBoolPointer()
		}
		if !tfTemplate.SelectedQuery.IsNull() {
			requestableRoleType.SelectedQuery = tfTemplate.SelectedQuery.ValueStringPointer()
		}
		if !tfTemplate.ShowOn.IsNull() {
			requestableRoleType.ShowOn = tfTemplate.ShowOn.ValueStringPointer()
		}

		requestableRoleTypes = append(requestableRoleTypes, requestableRoleType)
	}

	if len(requestableRoleTypes) > 0 {
		updateReq.RequestableRoleType = requestableRoleTypes
	}

	apiResp, httpResp, err := apiClient.EndpointsAPI.
		UpdateEndpoint(ctx).
		UpdateEndpointRequest(updateReq).
		Execute()
	log.Printf("createaccountenc:update %v", plan.CreateEntTaskforRemoveAcc.ValueString())
	if err != nil {
		log.Printf("Error Updating Endpoint: %v, HTTP Response: %v", err, httpResp)
		resp.Diagnostics.AddError(
			"Error Updating Endpoint",
			fmt.Sprintf("Error updating endpoints: %v", err),
		)
		return
	}
	if apiResp != nil && *apiResp.ErrorCode != "0" {
		log.Printf("Error Updating Endpoint: %v, Error code: %v", *apiResp.Msg, *apiResp.ErrorCode)
		resp.Diagnostics.AddError(
			"Error Updating Endpoint",
			fmt.Sprintf("Error: %v, Error code: %v", *apiResp.Msg, *apiResp.ErrorCode),
		)
		return
	}

	reqParams := openapi.GetEndpointsRequest{}
	reqParams.SetEndpointname(state.EndpointName.ValueString())
	readResp, _, err := apiClient.EndpointsAPI.GetEndpoints(ctx).GetEndpointsRequest(reqParams).Execute()
	if err != nil {
		log.Printf("Problem with the get function in update block")
		resp.Diagnostics.AddError("API Read Failed after updation of endpoint", fmt.Sprintf("Error: %v", err))
		return
	}
	if readResp != nil && *readResp.ErrorCode != "0" {
		log.Printf("Error Reading Endpoint after updation: %v, Error code: %v", *readResp.Message, *readResp.ErrorCode)
		resp.Diagnostics.AddError(
			"Error Reading endpoint after updation",
			fmt.Sprintf("Error: %v, Error code: %v", *readResp.Message, *readResp.ErrorCode),
		)
		return
	}

	log.Printf("createaccountenc:read %v", *readResp.Endpoints[0].CreateEntTaskforRemoveAcc)

	plan.ID = types.StringValue("endpoint-" + *readResp.Endpoints[0].Endpointname)
	plan.DisplayName = util.SafeString(readResp.Endpoints[0].DisplayName)
	plan.SecuritySystem = util.SafeString(readResp.Endpoints[0].Securitysystem)
	plan.Description = util.SafeString(readResp.Endpoints[0].Description)
	plan.OwnerType = util.SafeString(util.StringPtr(endpointsutil.TranslateValue(util.SafeDeref(readResp.Endpoints[0].OwnerType), endpointsutil.OwnerTypeMap)))
	plan.ResourceOwnerType = util.SafeString(util.StringPtr(endpointsutil.TranslateValue(util.SafeDeref(readResp.Endpoints[0].Requestownertype), endpointsutil.OwnerTypeMap)))
	plan.PrimaryAccountType = util.SafeString(readResp.Endpoints[0].PrimaryAccountType)
	plan.AccountTypeNoPasswordChange = util.SafeString(readResp.Endpoints[0].AccountTypeNoPasswordChange)
	plan.ServiceAccountNameRule = util.SafeString(readResp.Endpoints[0].ServiceAccountNameRule)
	plan.AccountNameValidatorRegex = util.SafeString(readResp.Endpoints[0].AccountNameValidatorRegex)
	plan.AllowChangePasswordSQLQuery = util.SafeString(readResp.Endpoints[0].AllowChangePasswordSqlquery)
	plan.ParentAccountPattern = util.SafeString(readResp.Endpoints[0].ParentAccountPattern)
	plan.SecuritySystem = util.SafeString(readResp.Endpoints[0].Securitysystem)
	plan.EndpointName = util.SafeString(readResp.Endpoints[0].Endpointname)
	plan.AccessQuery = util.SafeString(readResp.Endpoints[0].Accessquery)
	plan.DisplayName = util.SafeString(readResp.Endpoints[0].DisplayName)
	if readResp.Endpoints[0].Requestable != nil {
		if *readResp.Endpoints[0].Requestable == "true" {
			plan.Requestable = types.BoolValue(true)
		} else {
			plan.Requestable = types.BoolValue(false)
		}
	} else {
		plan.Requestable = types.BoolNull()
	}
	if readResp.Endpoints[0].AllowRemoveAllRoleOnRequest != nil {
		if *readResp.Endpoints[0].AllowRemoveAllRoleOnRequest == "true" {
			plan.AllowRemoveAllRoleOnRequest = types.BoolValue(true)
		} else {
			plan.AllowRemoveAllRoleOnRequest = types.BoolValue(false)
		}
	} else {
		plan.AllowRemoveAllRoleOnRequest = types.BoolNull()
	}
	if readResp.Endpoints[0].ConnectionconfigAsJson != nil {
		normalized, err := endpointsutil.NormalizeJSON(*readResp.Endpoints[0].ConnectionconfigAsJson)
		if err != nil {
			plan.ConnectionConfig = types.StringNull()
		} else {
			plan.ConnectionConfig = util.SafeString(&normalized)
		}
	} else {
		plan.ConnectionConfig = types.StringNull()
	}
	plan.AccountNameRule = util.SafeString(readResp.Endpoints[0].AccountNameRule)
	plan.ChangePasswordAccessQuery = util.SafeString(readResp.Endpoints[0].ChangePasswordAccessQuery)

	plan.PluginConfigs = util.SafeString(readResp.Endpoints[0].PluginConfigs)

	plan.CreateEntTaskforRemoveAcc = util.SafeString(readResp.Endpoints[0].CreateEntTaskforRemoveAcc)
	plan.EnableCopyAccess = util.SafeString(readResp.Endpoints[0].EnableCopyAccess)
	plan.EndpointConfig = util.SafeString(readResp.Endpoints[0].EndpointConfig)
	plan.ServiceAccountAccessQuery = util.SafeString(readResp.Endpoints[0].ServiceAccountAccessQuery)
	plan.UserAccountCorrelationRule = util.SafeString(readResp.Endpoints[0].UserAccountCorrelationRule)
	plan.StatusConfig = util.SafeString(readResp.Endpoints[0].StatusConfig)

	plan.CustomProperty1 = util.SafeString(readResp.Endpoints[0].CustomProperty1)
	plan.CustomProperty2 = util.SafeString(readResp.Endpoints[0].CustomProperty2)
	plan.CustomProperty3 = util.SafeString(readResp.Endpoints[0].CustomProperty3)
	plan.CustomProperty4 = util.SafeString(readResp.Endpoints[0].CustomProperty4)
	plan.CustomProperty5 = util.SafeString(readResp.Endpoints[0].CustomProperty5)
	plan.CustomProperty6 = util.SafeString(readResp.Endpoints[0].CustomProperty6)
	plan.CustomProperty7 = util.SafeString(readResp.Endpoints[0].CustomProperty7)
	plan.CustomProperty8 = util.SafeString(readResp.Endpoints[0].CustomProperty8)
	plan.CustomProperty9 = util.SafeString(readResp.Endpoints[0].CustomProperty9)
	plan.CustomProperty10 = util.SafeString(readResp.Endpoints[0].CustomProperty10)
	plan.CustomProperty11 = util.SafeString(readResp.Endpoints[0].CustomProperty11)
	plan.CustomProperty12 = util.SafeString(readResp.Endpoints[0].CustomProperty12)
	plan.CustomProperty13 = util.SafeString(readResp.Endpoints[0].CustomProperty13)
	plan.CustomProperty14 = util.SafeString(readResp.Endpoints[0].CustomProperty14)
	plan.CustomProperty15 = util.SafeString(readResp.Endpoints[0].CustomProperty15)
	plan.CustomProperty16 = util.SafeString(readResp.Endpoints[0].CustomProperty16)
	plan.CustomProperty17 = util.SafeString(readResp.Endpoints[0].CustomProperty17)
	plan.CustomProperty18 = util.SafeString(readResp.Endpoints[0].CustomProperty18)
	plan.CustomProperty19 = util.SafeString(readResp.Endpoints[0].CustomProperty19)
	plan.CustomProperty20 = util.SafeString(readResp.Endpoints[0].CustomProperty20)
	plan.CustomProperty21 = util.SafeString(readResp.Endpoints[0].CustomProperty21)
	plan.CustomProperty22 = util.SafeString(readResp.Endpoints[0].CustomProperty22)
	plan.CustomProperty23 = util.SafeString(readResp.Endpoints[0].CustomProperty23)
	plan.CustomProperty24 = util.SafeString(readResp.Endpoints[0].CustomProperty24)
	plan.CustomProperty25 = util.SafeString(readResp.Endpoints[0].CustomProperty25)
	plan.CustomProperty26 = util.SafeString(readResp.Endpoints[0].CustomProperty26)
	plan.CustomProperty27 = util.SafeString(readResp.Endpoints[0].CustomProperty27)
	plan.CustomProperty28 = util.SafeString(readResp.Endpoints[0].CustomProperty28)
	plan.CustomProperty29 = util.SafeString(readResp.Endpoints[0].CustomProperty29)
	plan.CustomProperty30 = util.SafeString(readResp.Endpoints[0].CustomProperty30)
	plan.CustomProperty31 = util.SafeString(readResp.Endpoints[0].Customproperty31)
	plan.CustomProperty32 = util.SafeString(readResp.Endpoints[0].Customproperty32)
	plan.CustomProperty33 = util.SafeString(readResp.Endpoints[0].Customproperty33)
	plan.CustomProperty34 = util.SafeString(readResp.Endpoints[0].Customproperty34)
	plan.CustomProperty35 = util.SafeString(readResp.Endpoints[0].Customproperty35)
	plan.CustomProperty36 = util.SafeString(readResp.Endpoints[0].Customproperty36)
	plan.CustomProperty37 = util.SafeString(readResp.Endpoints[0].Customproperty37)
	plan.CustomProperty38 = util.SafeString(readResp.Endpoints[0].Customproperty38)
	plan.CustomProperty39 = util.SafeString(readResp.Endpoints[0].Customproperty39)
	plan.CustomProperty40 = util.SafeString(readResp.Endpoints[0].Customproperty40)
	plan.CustomProperty41 = util.SafeString(readResp.Endpoints[0].Customproperty41)
	plan.CustomProperty42 = util.SafeString(readResp.Endpoints[0].Customproperty42)
	plan.CustomProperty43 = util.SafeString(readResp.Endpoints[0].Customproperty43)
	plan.CustomProperty44 = util.SafeString(readResp.Endpoints[0].Customproperty44)
	plan.CustomProperty45 = util.SafeString(readResp.Endpoints[0].Customproperty45)

	plan.AccountCustomProperty1Label = util.SafeString(readResp.Endpoints[0].AccountCustomProperty1Label)
	plan.AccountCustomProperty2Label = util.SafeString(readResp.Endpoints[0].AccountCustomProperty2Label)
	plan.AccountCustomProperty3Label = util.SafeString(readResp.Endpoints[0].AccountCustomProperty3Label)
	plan.AccountCustomProperty4Label = util.SafeString(readResp.Endpoints[0].AccountCustomProperty4Label)
	plan.AccountCustomProperty5Label = util.SafeString(readResp.Endpoints[0].AccountCustomProperty5Label)
	plan.AccountCustomProperty6Label = util.SafeString(readResp.Endpoints[0].AccountCustomProperty6Label)
	plan.AccountCustomProperty7Label = util.SafeString(readResp.Endpoints[0].AccountCustomProperty7Label)
	plan.AccountCustomProperty8Label = util.SafeString(readResp.Endpoints[0].AccountCustomProperty8Label)
	plan.AccountCustomProperty9Label = util.SafeString(readResp.Endpoints[0].AccountCustomProperty9Label)
	plan.AccountCustomProperty10Label = util.SafeString(readResp.Endpoints[0].AccountCustomProperty10Label)
	plan.AccountCustomProperty11Label = util.SafeString(readResp.Endpoints[0].AccountCustomProperty11Label)
	plan.AccountCustomProperty12Label = util.SafeString(readResp.Endpoints[0].AccountCustomProperty12Label)
	plan.AccountCustomProperty13Label = util.SafeString(readResp.Endpoints[0].AccountCustomProperty13Label)
	plan.AccountCustomProperty14Label = util.SafeString(readResp.Endpoints[0].AccountCustomProperty14Label)
	plan.AccountCustomProperty15Label = util.SafeString(readResp.Endpoints[0].AccountCustomProperty15Label)
	plan.AccountCustomProperty16Label = util.SafeString(readResp.Endpoints[0].AccountCustomProperty16Label)
	plan.AccountCustomProperty17Label = util.SafeString(readResp.Endpoints[0].AccountCustomProperty17Label)
	plan.AccountCustomProperty18Label = util.SafeString(readResp.Endpoints[0].AccountCustomProperty18Label)
	plan.AccountCustomProperty19Label = util.SafeString(readResp.Endpoints[0].AccountCustomProperty19Label)
	plan.AccountCustomProperty20Label = util.SafeString(readResp.Endpoints[0].AccountCustomProperty20Label)
	plan.AccountCustomProperty21Label = util.SafeString(readResp.Endpoints[0].AccountCustomProperty21Label)
	plan.AccountCustomProperty22Label = util.SafeString(readResp.Endpoints[0].AccountCustomProperty22Label)
	plan.AccountCustomProperty23Label = util.SafeString(readResp.Endpoints[0].AccountCustomProperty23Label)
	plan.AccountCustomProperty24Label = util.SafeString(readResp.Endpoints[0].AccountCustomProperty24Label)
	plan.AccountCustomProperty25Label = util.SafeString(readResp.Endpoints[0].AccountCustomProperty25Label)
	plan.AccountCustomProperty26Label = util.SafeString(readResp.Endpoints[0].AccountCustomProperty26Label)
	plan.AccountCustomProperty27Label = util.SafeString(readResp.Endpoints[0].AccountCustomProperty27Label)
	plan.AccountCustomProperty28Label = util.SafeString(readResp.Endpoints[0].AccountCustomProperty28Label)
	plan.AccountCustomProperty29Label = util.SafeString(readResp.Endpoints[0].AccountCustomProperty29Label)
	plan.AccountCustomProperty30Label = util.SafeString(readResp.Endpoints[0].AccountCustomProperty30Label)

	plan.CustomProperty31Label = util.SafeString(readResp.Endpoints[0].Customproperty31Label)
	plan.CustomProperty32Label = util.SafeString(readResp.Endpoints[0].Customproperty32Label)
	plan.CustomProperty33Label = util.SafeString(readResp.Endpoints[0].Customproperty33Label)
	plan.CustomProperty34Label = util.SafeString(readResp.Endpoints[0].Customproperty34Label)
	plan.CustomProperty35Label = util.SafeString(readResp.Endpoints[0].Customproperty35Label)
	plan.CustomProperty36Label = util.SafeString(readResp.Endpoints[0].Customproperty36Label)
	plan.CustomProperty37Label = util.SafeString(readResp.Endpoints[0].Customproperty37Label)
	plan.CustomProperty38Label = util.SafeString(readResp.Endpoints[0].Customproperty38Label)
	plan.CustomProperty39Label = util.SafeString(readResp.Endpoints[0].Customproperty39Label)
	plan.CustomProperty40Label = util.SafeString(readResp.Endpoints[0].Customproperty40Label)
	plan.CustomProperty41Label = util.SafeString(readResp.Endpoints[0].Customproperty41Label)
	plan.CustomProperty42Label = util.SafeString(readResp.Endpoints[0].Customproperty42Label)
	plan.CustomProperty43Label = util.SafeString(readResp.Endpoints[0].Customproperty43Label)
	plan.CustomProperty44Label = util.SafeString(readResp.Endpoints[0].Customproperty44Label)
	plan.CustomProperty45Label = util.SafeString(readResp.Endpoints[0].Customproperty45Label)
	plan.CustomProperty46Label = util.SafeString(readResp.Endpoints[0].Customproperty46Label)
	plan.CustomProperty47Label = util.SafeString(readResp.Endpoints[0].Customproperty47Label)
	plan.CustomProperty48Label = util.SafeString(readResp.Endpoints[0].Customproperty48Label)
	plan.CustomProperty49Label = util.SafeString(readResp.Endpoints[0].Customproperty49Label)
	plan.CustomProperty50Label = util.SafeString(readResp.Endpoints[0].Customproperty50Label)
	plan.CustomProperty51Label = util.SafeString(readResp.Endpoints[0].Customproperty51Label)
	plan.CustomProperty52Label = util.SafeString(readResp.Endpoints[0].Customproperty52Label)
	plan.CustomProperty53Label = util.SafeString(readResp.Endpoints[0].Customproperty53Label)
	plan.CustomProperty54Label = util.SafeString(readResp.Endpoints[0].Customproperty54Label)
	plan.CustomProperty55Label = util.SafeString(readResp.Endpoints[0].Customproperty55Label)
	plan.CustomProperty56Label = util.SafeString(readResp.Endpoints[0].Customproperty56Label)
	plan.CustomProperty57Label = util.SafeString(readResp.Endpoints[0].Customproperty57Label)
	plan.CustomProperty58Label = util.SafeString(readResp.Endpoints[0].Customproperty58Label)
	plan.CustomProperty59Label = util.SafeString(readResp.Endpoints[0].Customproperty59Label)
	plan.CustomProperty60Label = util.SafeString(readResp.Endpoints[0].Customproperty60Label)

	var planEmailTemplates types.List

	if readResp.Endpoints[0].Taskemailtemplates == nil || *readResp.Endpoints[0].Taskemailtemplates == "" {
		planEmailTemplates = types.ListNull(types.ObjectType{
			AttrTypes: map[string]attr.Type{
				"email_template_type": types.StringType,
				"task_type":           types.StringType,
				"email_template":      types.StringType,
			},
		})
	} else {
		taskEmailTemplatesStr := *readResp.Endpoints[0].Taskemailtemplates
		type ApiEmailTemplate struct {
			EmailTemplateType string `json:"emailTemplateType"`
			TaskType          string `json:"taskType"`
			EmailTemplate     string `json:"emailTemplate"`
		}

		var apiTemplates []ApiEmailTemplate
		if err := json.Unmarshal([]byte(taskEmailTemplatesStr), &apiTemplates); err != nil {
			resp.Diagnostics.AddError(
				"Client Error",
				fmt.Sprintf("Failed to parse taskemailtemplates JSON: %s", err),
			)
			return
		}
		log.Println("Task email templat: ", apiTemplates)

		emailTemplateObjects := make([]attr.Value, 0, len(apiTemplates))
		for _, t := range apiTemplates {
			obj, diags := types.ObjectValue(
				map[string]attr.Type{
					"email_template_type": types.StringType,
					"task_type":           types.StringType,
					"email_template":      types.StringType,
				},
				map[string]attr.Value{
					"email_template_type": types.StringValue(t.EmailTemplateType),
					"task_type":           types.StringValue(t.TaskType),
					"email_template":      types.StringValue(t.EmailTemplate),
				},
			)
			if diags.HasError() {
				resp.Diagnostics.Append(diags...)
				continue
			}
			emailTemplateObjects = append(emailTemplateObjects, obj)
		}

		listVal, diags := types.ListValue(
			types.ObjectType{
				AttrTypes: map[string]attr.Type{
					"email_template_type": types.StringType,
					"task_type":           types.StringType,
					"email_template":      types.StringType,
				},
			},
			emailTemplateObjects,
		)
		if diags.HasError() {
			resp.Diagnostics.Append(diags...)
			return
		}
		planEmailTemplates = listVal
	}

	plan.EmailTemplates = planEmailTemplates

	var planRequestableRoleTypes types.List

	if readResp.Endpoints[0].RoleTypeAsJson == nil || *readResp.Endpoints[0].RoleTypeAsJson == "" {
		planRequestableRoleTypes = types.ListNull(types.ObjectType{
			AttrTypes: map[string]attr.Type{
				"role_type":       types.StringType,
				"request_option":  types.StringType,
				"required":        types.BoolType,
				"requested_query": types.StringType,
				"selected_query":  types.StringType,
				"show_on":         types.StringType,
			},
		})
	} else {
		roleTypeJsonStr := *readResp.Endpoints[0].RoleTypeAsJson
		var roleTypeMap map[string]string
		if err := json.Unmarshal([]byte(roleTypeJsonStr), &roleTypeMap); err != nil {
			resp.Diagnostics.AddError(
				"Client Error",
				fmt.Sprintf("Failed to parse RoleTypeAsJson outer JSON: %s", err),
			)
			return
		}

		roleTypeObjects := make([]attr.Value, 0, len(roleTypeMap))

		for roleType, roleData := range roleTypeMap {
			parts := strings.Split(roleData, "__")

			get := func(i int) string {
				if i < len(parts) {
					return parts[i]
				}
				return ""
			}

			obj, diags := types.ObjectValue(
				map[string]attr.Type{
					"role_type":       types.StringType,
					"request_option":  types.StringType,
					"required":        types.BoolType,
					"requested_query": types.StringType,
					"selected_query":  types.StringType,
					"show_on":         types.StringType,
				},
				map[string]attr.Value{
					"role_type":       types.StringValue(endpointsutil.TranslateValue(roleType, endpointsutil.RoleTypeMap)),
					"request_option":  types.StringValue(endpointsutil.TranslateValue(parts[0], endpointsutil.RequestOptionMap)),
					"required":        types.BoolValue(get(1) == "1"),
					"requested_query": types.StringValue(get(2)),
					"selected_query":  types.StringValue(get(3)),
					"show_on":         types.StringValue(endpointsutil.TranslateValue(get(4), endpointsutil.ShowOnMap)),
				},
			)
			if diags.HasError() {
				resp.Diagnostics.Append(diags...)
				continue
			}
			roleTypeObjects = append(roleTypeObjects, obj)
		}

		listVal, diags := types.ListValue(
			types.ObjectType{
				AttrTypes: map[string]attr.Type{
					"role_type":       types.StringType,
					"request_option":  types.StringType,
					"required":        types.BoolType,
					"requested_query": types.StringType,
					"selected_query":  types.StringType,
					"show_on":         types.StringType,
				},
			},
			roleTypeObjects,
		)
		if diags.HasError() {
			resp.Diagnostics.Append(diags...)
			return
		}
		planRequestableRoleTypes = listVal
	}

	plan.RequestableRoleTypes = planRequestableRoleTypes

	if readResp.Endpoints[0].Disableaccountrequest != nil {
		disableAccountRequestStr := *readResp.Endpoints[0].Disableaccountrequest

		var disableAccountRequestMap map[string]string

		err := json.Unmarshal([]byte(disableAccountRequestStr), &disableAccountRequestMap)
		if err != nil {
			log.Printf("Error parsing disableaccountrequest JSON: %v", err)
		} else {
			plan.DisableNewAccountRequestIfAccountExists = types.StringValue(endpointsutil.NormalizeToStringBool(disableAccountRequestMap["DISABLENEWACCOUNT"]))
			plan.DisableRemoveAccount = types.StringValue(endpointsutil.NormalizeToStringBool(disableAccountRequestMap["DISABLEREMOVEACCOUNT"]))
			plan.DisableModifyAccount = types.StringValue(endpointsutil.NormalizeToStringBool(disableAccountRequestMap["DISABLEMODIFYACCOUNT"]))
			plan.BlockInflightRequest = types.StringValue(endpointsutil.NormalizeToStringBool(disableAccountRequestMap["BLOCKINFLIGHTREQUEST"]))

		}
	}

	msgValue := util.SafeDeref(apiResp.Msg)
	errorCodeValue := util.SafeDeref(apiResp.ErrorCode)

	plan.Msg = types.StringValue(msgValue)
	plan.ErrorCode = types.StringValue(errorCodeValue)

	stateUpdateDiagnostics := resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(stateUpdateDiagnostics...)
}

func (r *endpointResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
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

func (r *endpointResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	// Retrieve import ID and save to id attribute
	resource.ImportStatePassthroughID(ctx, path.Root("endpoint_name"), req, resp)
}
