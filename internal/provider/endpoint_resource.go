// Copyright (c) 2025 Saviynt Inc.
// SPDX-License-Identifier: MPL-2.0

// saviynt_endpoint_resource manages endpoints in the Saviynt Security Manager.
// The resource implements the full Terraform lifecycle:
//   - Create: provisions a new endpoint using the supplied configuration.
//   - Read: fetches the current endpoint state from Saviynt to keep Terraform’s state in sync.
//   - Update: applies any configuration changes to an existing endpoint.
//   - Import: brings an existing endpoint under Terraform management by its name.

package provider

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
	"terraform-provider-Saviynt/internal/client"
	"terraform-provider-Saviynt/util"
	"terraform-provider-Saviynt/util/endpointsutil"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"

	"github.com/hashicorp/terraform-plugin-framework/types"

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

type EndpointResource struct {
	client          client.SaviyntClientInterface
	token           string
	provider        client.SaviyntProviderInterface
	endpointFactory client.EndpointFactoryInterface
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
	return &EndpointResource{
		endpointFactory: &client.DefaultEndpointFactory{},
	}

}
func NewEndpointResourceWithFactory(factory client.EndpointFactoryInterface) resource.Resource {
	return &EndpointResource{
		endpointFactory: factory,
	}
}

func (r *EndpointResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = "saviynt_endpoint_resource"
}

func (r *EndpointResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
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
				Computed:    true,
				Description: "A message indicating the outcome of the operation.",
			},
			"error_code": schema.StringAttribute{
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

func (r *EndpointResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Check if provider data is available.
	if req.ProviderData == nil {
		log.Println("ProviderData is nil, returning early.")
		return
	}

	// Cast provider data to your provider type.
	prov, ok := req.ProviderData.(*SaviyntProvider)
	if !ok {
		resp.Diagnostics.AddError("Unexpected Provider Data", "Expected *SaviyntProvider")
		return
	}

	// Set the client, token, and provider reference from the provider state
	r.client = &client.SaviyntClientWrapper{Client: prov.client}
	r.token = prov.accessToken
	r.provider = &client.SaviyntProviderWrapper{Provider: prov} // Store provider reference for retry logic
}

// SetClient sets the client for testing purposes
func (r *EndpointResource) SetClient(client client.SaviyntClientInterface) {
	r.client = client
}

// SetToken sets the token for testing purposes
func (r *EndpointResource) SetToken(token string) {
	r.token = token
}

// SetProvider sets the provider for testing purposes
func (r *EndpointResource) SetProvider(provider client.SaviyntProviderInterface) {
	r.provider = provider
}

func (r *EndpointResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan EndpointResourceModel

	planGetDiagnostics := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(planGetDiagnostics...)
	if resp.Diagnostics.HasError() {
		log.Printf("[ERROR]: Failed to get plan in Create Block: %v", resp.Diagnostics)
		return
	}

	// Use interface pattern instead of direct API client creation
	apiResp, err := r.CreateEndpoint(ctx, &plan)
	if err != nil {
		resp.Diagnostics.AddError("API Create Failed In Create Block", err.Error())
		return
	}

	// Use unified method to set fields after CREATE operation
	r.UpdateModelFromCreateResponse(&plan, apiResp)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		log.Printf("[ERROR]: Failed to set state in Create Block: %v", resp.Diagnostics)
		return
	}
}

func (r *EndpointResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state EndpointResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		log.Printf("[ERROR] Failed to get state in Read Block: %v", resp.Diagnostics)
		return
	}

	// Use interface pattern for read operation
	apiResp, err := r.GetEndpoints(ctx, &state)
	if err != nil {
		log.Printf("[ERROR] Failed to get endpoints in Read Block: %v", err)
		resp.Diagnostics.AddError("API Read Failed In Read Block", err.Error())
		return
	}

	// Update model from read response
	r.SetEndpointFields(&state, apiResp, &resp.Diagnostics)
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		log.Printf("[ERROR] Failed to set state in Read Block: %v", resp.Diagnostics)
		return
	}

}

func (r *EndpointResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan EndpointResourceModel
	var state EndpointResourceModel

	// Get current state and plan
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		log.Printf("[ERROR]: Failed to get state in Update Block: %v", resp.Diagnostics)
		return
	}

	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		log.Printf("[ERROR]: Failed to get plan in Update Block: %v", resp.Diagnostics)
		return
	}

	// Validate endpoint name cannot be changed
	if plan.EndpointName.ValueString() != state.EndpointName.ValueString() {
		log.Printf("[ERROR]: Endpoint name cannot be updated from %s to %s", state.EndpointName.ValueString(), plan.EndpointName.ValueString())
		resp.Diagnostics.AddError("Error", "Endpoint name cannot be updated")
		return
	}

	// Build UPDATE request using existing helper function with proper diagnostics
	updateReq := r.BuildEndpointRequest(ctx, &plan, &resp.Diagnostics, false).(openapi.UpdateEndpointRequest)
	if resp.Diagnostics.HasError() {
		log.Printf("[ERROR]: Failed to build update request in Update Block: %v", resp.Diagnostics)
		return
	}

	// Handle mapped endpoints with proper diagnostics
	mappedEndpoints, diags := r.BuildMappedEndpointsForUpdateWithDiags(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		log.Printf("[ERROR]: Failed to build mapped endpoints in Update Block: %v", resp.Diagnostics)
		return
	}

	if len(mappedEndpoints) > 0 {
		updateReq.MappedEndpoints = mappedEndpoints
	}

	// Handle requestable role types with proper diagnostics
	requestableRoleTypes, diags := r.BuildRequestableRoleTypesForUpdateWithDiags(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		log.Printf("[ERROR]: Failed to build requestable role types in Update Block: %v", resp.Diagnostics)
		return
	}

	if len(requestableRoleTypes) > 0 {
		updateReq.RequestableRoleType = requestableRoleTypes
	}

	var apiResp *openapi.UpdateEndpoint200Response

	// Execute the update with retry logic
	err := r.provider.AuthenticatedAPICallWithRetry(ctx, "update_endpoint_main", func(token string) error {
		endpointOps := r.endpointFactory.CreateEndpointOperations(r.client.APIBaseURL(), token)
		resp, httpResp, err := endpointOps.UpdateEndpoint(ctx, updateReq)
		if httpResp != nil && httpResp.StatusCode == 401 {
			return fmt.Errorf("401 unauthorized")
		}
		apiResp = resp
		return err
	})

	if err != nil {
		log.Printf("[ERROR]: Problem with update function in Update Block. Error: %v", err)
		resp.Diagnostics.AddError("API Update Failed In Update Block", fmt.Sprintf("Error: %v", err))
		return
	}

	// Check API response
	if apiResp != nil && apiResp.ErrorCode != nil && *apiResp.ErrorCode != "0" {
		errorCode := util.SafeDeref(apiResp.ErrorCode)
		msg := util.SafeDeref(apiResp.Msg)
		log.Printf("[ERROR]: Error Updating Endpoint: %v, Error code: %v", msg, errorCode)
		resp.Diagnostics.AddError("Error Updating Endpoint In Update Block", fmt.Sprintf("Error: %v, Error code: %v", msg, errorCode))
		return
	}

	// Refresh state after update
	apiResp2, err := r.GetEndpoints(ctx, &plan)
	if err != nil {
		resp.Diagnostics.AddError("API Read Failed After Update In Update Block", err.Error())
		return
	}

	// Update model from read response
	r.SetEndpointFields(&plan, apiResp2, &resp.Diagnostics)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *EndpointResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
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

func (r *EndpointResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	// Retrieve import ID and save to id attribute
	resource.ImportStatePassthroughID(ctx, path.Root("endpoint_name"), req, resp)
}

// BuildEndpointRequest builds either CREATE or UPDATE request based on isCreate parameter
func (r *EndpointResource) BuildEndpointRequest(ctx context.Context, plan *EndpointResourceModel, diagnostics *diag.Diagnostics, isCreate bool) interface{} {
	if isCreate {
		// CREATE request with specific method patterns
		createReq := openapi.CreateEndpointRequest{
			// Core fields - CREATE uses ValueString() for required fields
			Endpointname:   plan.EndpointName.ValueString(),
			DisplayName:    plan.DisplayName.ValueString(),
			Securitysystem: plan.SecuritySystem.ValueString(),

			// Optional fields - CREATE uses util.StringPointerOrEmpty()
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

			// Boolean fields - CREATE uses util.BoolPointerOrEmpty()
			Requestable:                 util.BoolPointerOrEmpty(plan.Requestable),
			AllowRemoveAllRoleOnRequest: util.BoolPointerOrEmpty(plan.AllowRemoveAllRoleOnRequest),
		}

		// Add custom properties and labels for CREATE
		r.SetCustomProperties(&createReq, nil, plan, isCreate)
		r.SetCustomPropertyLabels(&createReq, nil, plan, isCreate)

		// Add email templates (unified logic)
		emailTemplates := r.BuildEmailTemplates(ctx, plan, diagnostics)
		if len(emailTemplates) > 0 {
			createReq.Taskemailtemplates = emailTemplates
		}

		return createReq
	} else {
		// UPDATE request with specific method patterns
		updateReq := openapi.UpdateEndpointRequest{
			// Core field - UPDATE uses ValueString() for endpoint name (required)
			Endpointname: plan.EndpointName.ValueString(),
			// Optional fields
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
			ParentAccountPattern:                    plan.ParentAccountPattern.ValueStringPointer(),
			ServiceAccountNameRule:                  plan.ServiceAccountNameRule.ValueStringPointer(),
			ServiceAccountAccessQuery:               plan.ServiceAccountAccessQuery.ValueStringPointer(),
			ChangePasswordAccessQuery:               plan.ChangePasswordAccessQuery.ValueStringPointer(),
			AccountNameValidatorRegex:               plan.AccountNameValidatorRegex.ValueStringPointer(),
			StatusConfig:                            plan.StatusConfig.ValueStringPointer(),
			PluginConfigs:                           plan.PluginConfigs.ValueStringPointer(),
			EndpointConfig:                          plan.EndpointConfig.ValueStringPointer(),

			// Boolean fields - UPDATE uses ValueBoolPointer()
			Requestable:                 plan.Requestable.ValueBoolPointer(),
			AllowRemoveAllRoleOnRequest: plan.AllowRemoveAllRoleOnRequest.ValueBoolPointer(),
		}

		// Add custom properties and labels for UPDATE
		r.SetCustomProperties(nil, &updateReq, plan, isCreate)
		r.SetCustomPropertyLabels(nil, &updateReq, plan, isCreate)

		// Add email templates (unified logic)
		emailTemplates := r.BuildEmailTemplates(ctx, plan, diagnostics)
		if len(emailTemplates) > 0 {
			updateReq.Taskemailtemplates = emailTemplates
		}

		return updateReq
	}
}

// SetCustomProperties sets custom properties 1-45 in either CREATE or UPDATE request
func (r *EndpointResource) SetCustomProperties(createReq *openapi.CreateEndpointRequest, updateReq *openapi.UpdateEndpointRequest, plan *EndpointResourceModel, isCreate bool) {
	// Custom properties use the same method (util.StringPointerOrEmpty) for both CREATE and UPDATE
	if isCreate {
		createReq.Customproperty1 = util.StringPointerOrEmpty(plan.CustomProperty1)
		createReq.Customproperty2 = util.StringPointerOrEmpty(plan.CustomProperty2)
		createReq.Customproperty3 = util.StringPointerOrEmpty(plan.CustomProperty3)
		createReq.Customproperty4 = util.StringPointerOrEmpty(plan.CustomProperty4)
		createReq.Customproperty5 = util.StringPointerOrEmpty(plan.CustomProperty5)
		createReq.Customproperty6 = util.StringPointerOrEmpty(plan.CustomProperty6)
		createReq.Customproperty7 = util.StringPointerOrEmpty(plan.CustomProperty7)
		createReq.Customproperty8 = util.StringPointerOrEmpty(plan.CustomProperty8)
		createReq.Customproperty9 = util.StringPointerOrEmpty(plan.CustomProperty9)
		createReq.Customproperty10 = util.StringPointerOrEmpty(plan.CustomProperty10)
		createReq.Customproperty11 = util.StringPointerOrEmpty(plan.CustomProperty11)
		createReq.Customproperty12 = util.StringPointerOrEmpty(plan.CustomProperty12)
		createReq.Customproperty13 = util.StringPointerOrEmpty(plan.CustomProperty13)
		createReq.Customproperty14 = util.StringPointerOrEmpty(plan.CustomProperty14)
		createReq.Customproperty15 = util.StringPointerOrEmpty(plan.CustomProperty15)
		createReq.Customproperty16 = util.StringPointerOrEmpty(plan.CustomProperty16)
		createReq.Customproperty17 = util.StringPointerOrEmpty(plan.CustomProperty17)
		createReq.Customproperty18 = util.StringPointerOrEmpty(plan.CustomProperty18)
		createReq.Customproperty19 = util.StringPointerOrEmpty(plan.CustomProperty19)
		createReq.Customproperty20 = util.StringPointerOrEmpty(plan.CustomProperty20)
		createReq.Customproperty21 = util.StringPointerOrEmpty(plan.CustomProperty21)
		createReq.Customproperty22 = util.StringPointerOrEmpty(plan.CustomProperty22)
		createReq.Customproperty23 = util.StringPointerOrEmpty(plan.CustomProperty23)
		createReq.Customproperty24 = util.StringPointerOrEmpty(plan.CustomProperty24)
		createReq.Customproperty25 = util.StringPointerOrEmpty(plan.CustomProperty25)
		createReq.Customproperty26 = util.StringPointerOrEmpty(plan.CustomProperty26)
		createReq.Customproperty27 = util.StringPointerOrEmpty(plan.CustomProperty27)
		createReq.Customproperty28 = util.StringPointerOrEmpty(plan.CustomProperty28)
		createReq.Customproperty29 = util.StringPointerOrEmpty(plan.CustomProperty29)
		createReq.Customproperty30 = util.StringPointerOrEmpty(plan.CustomProperty30)
		createReq.Customproperty31 = util.StringPointerOrEmpty(plan.CustomProperty31)
		createReq.Customproperty32 = util.StringPointerOrEmpty(plan.CustomProperty32)
		createReq.Customproperty33 = util.StringPointerOrEmpty(plan.CustomProperty33)
		createReq.Customproperty34 = util.StringPointerOrEmpty(plan.CustomProperty34)
		createReq.Customproperty35 = util.StringPointerOrEmpty(plan.CustomProperty35)
		createReq.Customproperty36 = util.StringPointerOrEmpty(plan.CustomProperty36)
		createReq.Customproperty37 = util.StringPointerOrEmpty(plan.CustomProperty37)
		createReq.Customproperty38 = util.StringPointerOrEmpty(plan.CustomProperty38)
		createReq.Customproperty39 = util.StringPointerOrEmpty(plan.CustomProperty39)
		createReq.Customproperty40 = util.StringPointerOrEmpty(plan.CustomProperty40)
		createReq.Customproperty41 = util.StringPointerOrEmpty(plan.CustomProperty41)
		createReq.Customproperty42 = util.StringPointerOrEmpty(plan.CustomProperty42)
		createReq.Customproperty43 = util.StringPointerOrEmpty(plan.CustomProperty43)
		createReq.Customproperty44 = util.StringPointerOrEmpty(plan.CustomProperty44)
		createReq.Customproperty45 = util.StringPointerOrEmpty(plan.CustomProperty45)
	} else {
		updateReq.Customproperty1 = util.StringPointerOrEmpty(plan.CustomProperty1)
		updateReq.Customproperty2 = util.StringPointerOrEmpty(plan.CustomProperty2)
		updateReq.Customproperty3 = util.StringPointerOrEmpty(plan.CustomProperty3)
		updateReq.Customproperty4 = util.StringPointerOrEmpty(plan.CustomProperty4)
		updateReq.Customproperty5 = util.StringPointerOrEmpty(plan.CustomProperty5)
		updateReq.Customproperty6 = util.StringPointerOrEmpty(plan.CustomProperty6)
		updateReq.Customproperty7 = util.StringPointerOrEmpty(plan.CustomProperty7)
		updateReq.Customproperty8 = util.StringPointerOrEmpty(plan.CustomProperty8)
		updateReq.Customproperty9 = util.StringPointerOrEmpty(plan.CustomProperty9)
		updateReq.Customproperty10 = util.StringPointerOrEmpty(plan.CustomProperty10)
		updateReq.Customproperty11 = util.StringPointerOrEmpty(plan.CustomProperty11)
		updateReq.Customproperty12 = util.StringPointerOrEmpty(plan.CustomProperty12)
		updateReq.Customproperty13 = util.StringPointerOrEmpty(plan.CustomProperty13)
		updateReq.Customproperty14 = util.StringPointerOrEmpty(plan.CustomProperty14)
		updateReq.Customproperty15 = util.StringPointerOrEmpty(plan.CustomProperty15)
		updateReq.Customproperty16 = util.StringPointerOrEmpty(plan.CustomProperty16)
		updateReq.Customproperty17 = util.StringPointerOrEmpty(plan.CustomProperty17)
		updateReq.Customproperty18 = util.StringPointerOrEmpty(plan.CustomProperty18)
		updateReq.Customproperty19 = util.StringPointerOrEmpty(plan.CustomProperty19)
		updateReq.Customproperty20 = util.StringPointerOrEmpty(plan.CustomProperty20)
		updateReq.Customproperty21 = util.StringPointerOrEmpty(plan.CustomProperty21)
		updateReq.Customproperty22 = util.StringPointerOrEmpty(plan.CustomProperty22)
		updateReq.Customproperty23 = util.StringPointerOrEmpty(plan.CustomProperty23)
		updateReq.Customproperty24 = util.StringPointerOrEmpty(plan.CustomProperty24)
		updateReq.Customproperty25 = util.StringPointerOrEmpty(plan.CustomProperty25)
		updateReq.Customproperty26 = util.StringPointerOrEmpty(plan.CustomProperty26)
		updateReq.Customproperty27 = util.StringPointerOrEmpty(plan.CustomProperty27)
		updateReq.Customproperty28 = util.StringPointerOrEmpty(plan.CustomProperty28)
		updateReq.Customproperty29 = util.StringPointerOrEmpty(plan.CustomProperty29)
		updateReq.Customproperty30 = util.StringPointerOrEmpty(plan.CustomProperty30)
		updateReq.Customproperty31 = util.StringPointerOrEmpty(plan.CustomProperty31)
		updateReq.Customproperty32 = util.StringPointerOrEmpty(plan.CustomProperty32)
		updateReq.Customproperty33 = util.StringPointerOrEmpty(plan.CustomProperty33)
		updateReq.Customproperty34 = util.StringPointerOrEmpty(plan.CustomProperty34)
		updateReq.Customproperty35 = util.StringPointerOrEmpty(plan.CustomProperty35)
		updateReq.Customproperty36 = util.StringPointerOrEmpty(plan.CustomProperty36)
		updateReq.Customproperty37 = util.StringPointerOrEmpty(plan.CustomProperty37)
		updateReq.Customproperty38 = util.StringPointerOrEmpty(plan.CustomProperty38)
		updateReq.Customproperty39 = util.StringPointerOrEmpty(plan.CustomProperty39)
		updateReq.Customproperty40 = util.StringPointerOrEmpty(plan.CustomProperty40)
		updateReq.Customproperty41 = util.StringPointerOrEmpty(plan.CustomProperty41)
		updateReq.Customproperty42 = util.StringPointerOrEmpty(plan.CustomProperty42)
		updateReq.Customproperty43 = util.StringPointerOrEmpty(plan.CustomProperty43)
		updateReq.Customproperty44 = util.StringPointerOrEmpty(plan.CustomProperty44)
		updateReq.Customproperty45 = util.StringPointerOrEmpty(plan.CustomProperty45)
	}
}

// SetCustomPropertyLabels sets custom property labels 1-60 in either CREATE or UPDATE request
func (r *EndpointResource) SetCustomPropertyLabels(createReq *openapi.CreateEndpointRequest, updateReq *openapi.UpdateEndpointRequest, plan *EndpointResourceModel, isCreate bool) {
	// Custom property labels use the same method (util.StringPointerOrEmpty) for both CREATE and UPDATE
	if isCreate {
		// Labels 1-30 use AccountCustomProperty pattern
		createReq.Customproperty1Label = util.StringPointerOrEmpty(plan.AccountCustomProperty1Label)
		createReq.Customproperty2Label = util.StringPointerOrEmpty(plan.AccountCustomProperty2Label)
		createReq.Customproperty3Label = util.StringPointerOrEmpty(plan.AccountCustomProperty3Label)
		createReq.Customproperty4Label = util.StringPointerOrEmpty(plan.AccountCustomProperty4Label)
		createReq.Customproperty5Label = util.StringPointerOrEmpty(plan.AccountCustomProperty5Label)
		createReq.Customproperty6Label = util.StringPointerOrEmpty(plan.AccountCustomProperty6Label)
		createReq.Customproperty7Label = util.StringPointerOrEmpty(plan.AccountCustomProperty7Label)
		createReq.Customproperty8Label = util.StringPointerOrEmpty(plan.AccountCustomProperty8Label)
		createReq.Customproperty9Label = util.StringPointerOrEmpty(plan.AccountCustomProperty9Label)
		createReq.Customproperty10Label = util.StringPointerOrEmpty(plan.AccountCustomProperty10Label)
		createReq.Customproperty11Label = util.StringPointerOrEmpty(plan.AccountCustomProperty11Label)
		createReq.Customproperty12Label = util.StringPointerOrEmpty(plan.AccountCustomProperty12Label)
		createReq.Customproperty13Label = util.StringPointerOrEmpty(plan.AccountCustomProperty13Label)
		createReq.Customproperty14Label = util.StringPointerOrEmpty(plan.AccountCustomProperty14Label)
		createReq.Customproperty15Label = util.StringPointerOrEmpty(plan.AccountCustomProperty15Label)
		createReq.Customproperty16Label = util.StringPointerOrEmpty(plan.AccountCustomProperty16Label)
		createReq.Customproperty17Label = util.StringPointerOrEmpty(plan.AccountCustomProperty17Label)
		createReq.Customproperty18Label = util.StringPointerOrEmpty(plan.AccountCustomProperty18Label)
		createReq.Customproperty19Label = util.StringPointerOrEmpty(plan.AccountCustomProperty19Label)
		createReq.Customproperty20Label = util.StringPointerOrEmpty(plan.AccountCustomProperty20Label)
		createReq.Customproperty21Label = util.StringPointerOrEmpty(plan.AccountCustomProperty21Label)
		createReq.Customproperty22Label = util.StringPointerOrEmpty(plan.AccountCustomProperty22Label)
		createReq.Customproperty23Label = util.StringPointerOrEmpty(plan.AccountCustomProperty23Label)
		createReq.Customproperty24Label = util.StringPointerOrEmpty(plan.AccountCustomProperty24Label)
		createReq.Customproperty25Label = util.StringPointerOrEmpty(plan.AccountCustomProperty25Label)
		createReq.Customproperty26Label = util.StringPointerOrEmpty(plan.AccountCustomProperty26Label)
		createReq.Customproperty27Label = util.StringPointerOrEmpty(plan.AccountCustomProperty27Label)
		createReq.Customproperty28Label = util.StringPointerOrEmpty(plan.AccountCustomProperty28Label)
		createReq.Customproperty29Label = util.StringPointerOrEmpty(plan.AccountCustomProperty29Label)
		createReq.Customproperty30Label = util.StringPointerOrEmpty(plan.AccountCustomProperty30Label)

		// Labels 31-60 use CustomProperty pattern
		createReq.Customproperty31Label = util.StringPointerOrEmpty(plan.CustomProperty31Label)
		createReq.Customproperty32Label = util.StringPointerOrEmpty(plan.CustomProperty32Label)
		createReq.Customproperty33Label = util.StringPointerOrEmpty(plan.CustomProperty33Label)
		createReq.Customproperty34Label = util.StringPointerOrEmpty(plan.CustomProperty34Label)
		createReq.Customproperty35Label = util.StringPointerOrEmpty(plan.CustomProperty35Label)
		createReq.Customproperty36Label = util.StringPointerOrEmpty(plan.CustomProperty36Label)
		createReq.Customproperty37Label = util.StringPointerOrEmpty(plan.CustomProperty37Label)
		createReq.Customproperty38Label = util.StringPointerOrEmpty(plan.CustomProperty38Label)
		createReq.Customproperty39Label = util.StringPointerOrEmpty(plan.CustomProperty39Label)
		createReq.Customproperty40Label = util.StringPointerOrEmpty(plan.CustomProperty40Label)
		createReq.Customproperty41Label = util.StringPointerOrEmpty(plan.CustomProperty41Label)
		createReq.Customproperty42Label = util.StringPointerOrEmpty(plan.CustomProperty42Label)
		createReq.Customproperty43Label = util.StringPointerOrEmpty(plan.CustomProperty43Label)
		createReq.Customproperty44Label = util.StringPointerOrEmpty(plan.CustomProperty44Label)
		createReq.Customproperty45Label = util.StringPointerOrEmpty(plan.CustomProperty45Label)
		createReq.Customproperty46Label = util.StringPointerOrEmpty(plan.CustomProperty46Label)
		createReq.Customproperty47Label = util.StringPointerOrEmpty(plan.CustomProperty47Label)
		createReq.Customproperty48Label = util.StringPointerOrEmpty(plan.CustomProperty48Label)
		createReq.Customproperty49Label = util.StringPointerOrEmpty(plan.CustomProperty49Label)
		createReq.Customproperty50Label = util.StringPointerOrEmpty(plan.CustomProperty50Label)
		createReq.Customproperty51Label = util.StringPointerOrEmpty(plan.CustomProperty51Label)
		createReq.Customproperty52Label = util.StringPointerOrEmpty(plan.CustomProperty52Label)
		createReq.Customproperty53Label = util.StringPointerOrEmpty(plan.CustomProperty53Label)
		createReq.Customproperty54Label = util.StringPointerOrEmpty(plan.CustomProperty54Label)
		createReq.Customproperty55Label = util.StringPointerOrEmpty(plan.CustomProperty55Label)
		createReq.Customproperty56Label = util.StringPointerOrEmpty(plan.CustomProperty56Label)
		createReq.Customproperty57Label = util.StringPointerOrEmpty(plan.CustomProperty57Label)
		createReq.Customproperty58Label = util.StringPointerOrEmpty(plan.CustomProperty58Label)
		createReq.Customproperty59Label = util.StringPointerOrEmpty(plan.CustomProperty59Label)
		createReq.Customproperty60Label = util.StringPointerOrEmpty(plan.CustomProperty60Label)
	} else {
		// Same assignments for UPDATE request
		updateReq.Customproperty1Label = util.StringPointerOrEmpty(plan.AccountCustomProperty1Label)
		updateReq.Customproperty2Label = util.StringPointerOrEmpty(plan.AccountCustomProperty2Label)
		updateReq.Customproperty3Label = util.StringPointerOrEmpty(plan.AccountCustomProperty3Label)
		updateReq.Customproperty4Label = util.StringPointerOrEmpty(plan.AccountCustomProperty4Label)
		updateReq.Customproperty5Label = util.StringPointerOrEmpty(plan.AccountCustomProperty5Label)
		updateReq.Customproperty6Label = util.StringPointerOrEmpty(plan.AccountCustomProperty6Label)
		updateReq.Customproperty7Label = util.StringPointerOrEmpty(plan.AccountCustomProperty7Label)
		updateReq.Customproperty8Label = util.StringPointerOrEmpty(plan.AccountCustomProperty8Label)
		updateReq.Customproperty9Label = util.StringPointerOrEmpty(plan.AccountCustomProperty9Label)
		updateReq.Customproperty10Label = util.StringPointerOrEmpty(plan.AccountCustomProperty10Label)
		updateReq.Customproperty11Label = util.StringPointerOrEmpty(plan.AccountCustomProperty11Label)
		updateReq.Customproperty12Label = util.StringPointerOrEmpty(plan.AccountCustomProperty12Label)
		updateReq.Customproperty13Label = util.StringPointerOrEmpty(plan.AccountCustomProperty13Label)
		updateReq.Customproperty14Label = util.StringPointerOrEmpty(plan.AccountCustomProperty14Label)
		updateReq.Customproperty15Label = util.StringPointerOrEmpty(plan.AccountCustomProperty15Label)
		updateReq.Customproperty16Label = util.StringPointerOrEmpty(plan.AccountCustomProperty16Label)
		updateReq.Customproperty17Label = util.StringPointerOrEmpty(plan.AccountCustomProperty17Label)
		updateReq.Customproperty18Label = util.StringPointerOrEmpty(plan.AccountCustomProperty18Label)
		updateReq.Customproperty19Label = util.StringPointerOrEmpty(plan.AccountCustomProperty19Label)
		updateReq.Customproperty20Label = util.StringPointerOrEmpty(plan.AccountCustomProperty20Label)
		updateReq.Customproperty21Label = util.StringPointerOrEmpty(plan.AccountCustomProperty21Label)
		updateReq.Customproperty22Label = util.StringPointerOrEmpty(plan.AccountCustomProperty22Label)
		updateReq.Customproperty23Label = util.StringPointerOrEmpty(plan.AccountCustomProperty23Label)
		updateReq.Customproperty24Label = util.StringPointerOrEmpty(plan.AccountCustomProperty24Label)
		updateReq.Customproperty25Label = util.StringPointerOrEmpty(plan.AccountCustomProperty25Label)
		updateReq.Customproperty26Label = util.StringPointerOrEmpty(plan.AccountCustomProperty26Label)
		updateReq.Customproperty27Label = util.StringPointerOrEmpty(plan.AccountCustomProperty27Label)
		updateReq.Customproperty28Label = util.StringPointerOrEmpty(plan.AccountCustomProperty28Label)
		updateReq.Customproperty29Label = util.StringPointerOrEmpty(plan.AccountCustomProperty29Label)
		updateReq.Customproperty30Label = util.StringPointerOrEmpty(plan.AccountCustomProperty30Label)

		// Labels 31-60 use CustomProperty pattern
		updateReq.Customproperty31Label = util.StringPointerOrEmpty(plan.CustomProperty31Label)
		updateReq.Customproperty32Label = util.StringPointerOrEmpty(plan.CustomProperty32Label)
		updateReq.Customproperty33Label = util.StringPointerOrEmpty(plan.CustomProperty33Label)
		updateReq.Customproperty34Label = util.StringPointerOrEmpty(plan.CustomProperty34Label)
		updateReq.Customproperty35Label = util.StringPointerOrEmpty(plan.CustomProperty35Label)
		updateReq.Customproperty36Label = util.StringPointerOrEmpty(plan.CustomProperty36Label)
		updateReq.Customproperty37Label = util.StringPointerOrEmpty(plan.CustomProperty37Label)
		updateReq.Customproperty38Label = util.StringPointerOrEmpty(plan.CustomProperty38Label)
		updateReq.Customproperty39Label = util.StringPointerOrEmpty(plan.CustomProperty39Label)
		updateReq.Customproperty40Label = util.StringPointerOrEmpty(plan.CustomProperty40Label)
		updateReq.Customproperty41Label = util.StringPointerOrEmpty(plan.CustomProperty41Label)
		updateReq.Customproperty42Label = util.StringPointerOrEmpty(plan.CustomProperty42Label)
		updateReq.Customproperty43Label = util.StringPointerOrEmpty(plan.CustomProperty43Label)
		updateReq.Customproperty44Label = util.StringPointerOrEmpty(plan.CustomProperty44Label)
		updateReq.Customproperty45Label = util.StringPointerOrEmpty(plan.CustomProperty45Label)
		updateReq.Customproperty46Label = util.StringPointerOrEmpty(plan.CustomProperty46Label)
		updateReq.Customproperty47Label = util.StringPointerOrEmpty(plan.CustomProperty47Label)
		updateReq.Customproperty48Label = util.StringPointerOrEmpty(plan.CustomProperty48Label)
		updateReq.Customproperty49Label = util.StringPointerOrEmpty(plan.CustomProperty49Label)
		updateReq.Customproperty50Label = util.StringPointerOrEmpty(plan.CustomProperty50Label)
		updateReq.Customproperty51Label = util.StringPointerOrEmpty(plan.CustomProperty51Label)
		updateReq.Customproperty52Label = util.StringPointerOrEmpty(plan.CustomProperty52Label)
		updateReq.Customproperty53Label = util.StringPointerOrEmpty(plan.CustomProperty53Label)
		updateReq.Customproperty54Label = util.StringPointerOrEmpty(plan.CustomProperty54Label)
		updateReq.Customproperty55Label = util.StringPointerOrEmpty(plan.CustomProperty55Label)
		updateReq.Customproperty56Label = util.StringPointerOrEmpty(plan.CustomProperty56Label)
		updateReq.Customproperty57Label = util.StringPointerOrEmpty(plan.CustomProperty57Label)
		updateReq.Customproperty58Label = util.StringPointerOrEmpty(plan.CustomProperty58Label)
		updateReq.Customproperty59Label = util.StringPointerOrEmpty(plan.CustomProperty59Label)
		updateReq.Customproperty60Label = util.StringPointerOrEmpty(plan.CustomProperty60Label)
	}
}

// BuildEmailTemplates builds email templates for both CREATE and UPDATE operations
func (r *EndpointResource) BuildEmailTemplates(ctx context.Context, plan *EndpointResourceModel, diagnostics *diag.Diagnostics) []openapi.CreateEndpointRequestEmailTemplateInner {
	var emailTemplates []openapi.CreateEndpointRequestEmailTemplateInner
	var tfEmailTemplates []EmailTemplate

	diags := plan.EmailTemplates.ElementsAs(ctx, &tfEmailTemplates, true)
	diagnostics.Append(diags...)
	if diagnostics.HasError() {
		return emailTemplates
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

	return emailTemplates
}

// SetEndpointFields sets endpoint fields for both READ (state) and UPDATE (plan) operations
func (r *EndpointResource) SetEndpointFields(target *EndpointResourceModel, readResp *openapi.GetEndpoints200Response, diagnostics *diag.Diagnostics) {
	endpoint := readResp.Endpoints[0]

	target.ID = types.StringValue("endpoint-" + *endpoint.Endpointname)
	target.DisplayName = util.SafeString(endpoint.DisplayName)
	target.SecuritySystem = util.SafeString(endpoint.Securitysystem)
	target.Description = util.SafeString(endpoint.Description)
	target.OwnerType = util.SafeString(util.StringPtr(endpointsutil.TranslateValue(util.SafeDeref(endpoint.OwnerType), endpointsutil.OwnerTypeMap)))
	target.ResourceOwnerType = util.SafeString(util.StringPtr(endpointsutil.TranslateValue(util.SafeDeref(endpoint.Requestownertype), endpointsutil.OwnerTypeMap)))
	target.PrimaryAccountType = util.SafeString(endpoint.PrimaryAccountType)
	target.AccountTypeNoPasswordChange = util.SafeString(endpoint.AccountTypeNoPasswordChange)
	target.ServiceAccountNameRule = util.SafeString(endpoint.ServiceAccountNameRule)
	target.AccountNameValidatorRegex = util.SafeString(endpoint.AccountNameValidatorRegex)
	target.AllowChangePasswordSQLQuery = util.SafeString(endpoint.AllowChangePasswordSqlquery)
	target.ParentAccountPattern = util.SafeString(endpoint.ParentAccountPattern)
	target.EndpointName = util.SafeString(endpoint.Endpointname)
	target.AccessQuery = util.SafeString(endpoint.Accessquery)
	target.DisplayName = util.SafeString(endpoint.DisplayName)
	// Handle boolean fields with proper null checks
	if endpoint.Requestable != nil {
		if *endpoint.Requestable == "true" {
			target.Requestable = types.BoolValue(true)
		} else {
			target.Requestable = types.BoolValue(false)
		}
	} else {
		target.Requestable = types.BoolNull()
	}

	if endpoint.AllowRemoveAllRoleOnRequest != nil {
		if *endpoint.AllowRemoveAllRoleOnRequest == "true" {
			target.AllowRemoveAllRoleOnRequest = types.BoolValue(true)
		} else {
			target.AllowRemoveAllRoleOnRequest = types.BoolValue(false)
		}
	} else {
		target.AllowRemoveAllRoleOnRequest = types.BoolNull()
	}

	// Handle ConnectionConfig with JSON normalization
	if endpoint.ConnectionconfigAsJson != nil {
		normalized, err := endpointsutil.NormalizeJSON(*endpoint.ConnectionconfigAsJson)
		if err != nil {
			target.ConnectionConfig = types.StringNull()
		} else {
			target.ConnectionConfig = types.StringValue(normalized)
		}
	} else {
		target.ConnectionConfig = types.StringNull()
	}

	target.AccountNameRule = util.SafeString(endpoint.AccountNameRule)
	target.ChangePasswordAccessQuery = util.SafeString(endpoint.ChangePasswordAccessQuery)
	target.PluginConfigs = util.SafeString(endpoint.PluginConfigs)
	target.CreateEntTaskforRemoveAcc = util.SafeString(endpoint.CreateEntTaskforRemoveAcc)
	target.EnableCopyAccess = util.SafeString(endpoint.EnableCopyAccess)
	target.EndpointConfig = util.SafeString(endpoint.EndpointConfig)
	target.ServiceAccountAccessQuery = util.SafeString(endpoint.ServiceAccountAccessQuery)
	target.UserAccountCorrelationRule = util.SafeString(endpoint.UserAccountCorrelationRule)
	target.StatusConfig = util.SafeString(endpoint.StatusConfig)
	if endpoint.Disableaccountrequest != nil {
		disableAccountRequestStr := *endpoint.Disableaccountrequest

		var disableAccountRequestMap map[string]string

		err := json.Unmarshal([]byte(disableAccountRequestStr), &disableAccountRequestMap)
		if err != nil {
			log.Printf("Error parsing disableaccountrequest JSON: %v", err)
		} else {
			target.DisableNewAccountRequestIfAccountExists = types.StringValue(endpointsutil.NormalizeToStringBool(disableAccountRequestMap["DISABLENEWACCOUNT"]))
			target.DisableRemoveAccount = types.StringValue(endpointsutil.NormalizeToStringBool(disableAccountRequestMap["DISABLEREMOVEACCOUNT"]))
			target.DisableModifyAccount = types.StringValue(endpointsutil.NormalizeToStringBool(disableAccountRequestMap["DISABLEMODIFYACCOUNT"]))
			target.BlockInflightRequest = types.StringValue(endpointsutil.NormalizeToStringBool(disableAccountRequestMap["BLOCKINFLIGHTREQUEST"]))
		}
	}
	// Set custom properties 1-45
	r.SetCustomPropertiesFromAPI(target, &endpoint)

	// Set custom property labels 1-60
	r.SetCustomPropertyLabelsFromAPI(target, &endpoint)

	// Process complex fields - Email Templates
	target.EmailTemplates = r.BuildEmailTemplatesFromAPI(&endpoint, diagnostics)

	// Process complex fields - Requestable Role Types
	target.RequestableRoleTypes = r.BuildRequestableRoleTypesFromAPI(&endpoint, diagnostics)

	// Set response message fields
	msgValue := util.SafeDeref(readResp.Message)
	target.Msg = util.SafeString(&msgValue)
	target.ErrorCode = util.SafeString(readResp.ErrorCode)

}

// SetCustomPropertiesFromAPI sets custom properties 1-45 from API response
func (r *EndpointResource) SetCustomPropertiesFromAPI(target *EndpointResourceModel, apiResponse *openapi.GetEndpoints200ResponseEndpointsInner) {
	target.CustomProperty1 = util.SafeString(apiResponse.CustomProperty1)
	target.CustomProperty2 = util.SafeString(apiResponse.CustomProperty2)
	target.CustomProperty3 = util.SafeString(apiResponse.CustomProperty3)
	target.CustomProperty4 = util.SafeString(apiResponse.CustomProperty4)
	target.CustomProperty5 = util.SafeString(apiResponse.CustomProperty5)
	target.CustomProperty6 = util.SafeString(apiResponse.CustomProperty6)
	target.CustomProperty7 = util.SafeString(apiResponse.CustomProperty7)
	target.CustomProperty8 = util.SafeString(apiResponse.CustomProperty8)
	target.CustomProperty9 = util.SafeString(apiResponse.CustomProperty9)
	target.CustomProperty10 = util.SafeString(apiResponse.CustomProperty10)
	target.CustomProperty11 = util.SafeString(apiResponse.CustomProperty11)
	target.CustomProperty12 = util.SafeString(apiResponse.CustomProperty12)
	target.CustomProperty13 = util.SafeString(apiResponse.CustomProperty13)
	target.CustomProperty14 = util.SafeString(apiResponse.CustomProperty14)
	target.CustomProperty15 = util.SafeString(apiResponse.CustomProperty15)
	target.CustomProperty16 = util.SafeString(apiResponse.CustomProperty16)
	target.CustomProperty17 = util.SafeString(apiResponse.CustomProperty17)
	target.CustomProperty18 = util.SafeString(apiResponse.CustomProperty18)
	target.CustomProperty19 = util.SafeString(apiResponse.CustomProperty19)
	target.CustomProperty20 = util.SafeString(apiResponse.CustomProperty20)
	target.CustomProperty21 = util.SafeString(apiResponse.CustomProperty21)
	target.CustomProperty22 = util.SafeString(apiResponse.CustomProperty22)
	target.CustomProperty23 = util.SafeString(apiResponse.CustomProperty23)
	target.CustomProperty24 = util.SafeString(apiResponse.CustomProperty24)
	target.CustomProperty25 = util.SafeString(apiResponse.CustomProperty25)
	target.CustomProperty26 = util.SafeString(apiResponse.CustomProperty26)
	target.CustomProperty27 = util.SafeString(apiResponse.CustomProperty27)
	target.CustomProperty28 = util.SafeString(apiResponse.CustomProperty28)
	target.CustomProperty29 = util.SafeString(apiResponse.CustomProperty29)
	target.CustomProperty30 = util.SafeString(apiResponse.CustomProperty30)
	target.CustomProperty31 = util.SafeString(apiResponse.Customproperty31)
	target.CustomProperty32 = util.SafeString(apiResponse.Customproperty32)
	target.CustomProperty33 = util.SafeString(apiResponse.Customproperty33)
	target.CustomProperty34 = util.SafeString(apiResponse.Customproperty34)
	target.CustomProperty35 = util.SafeString(apiResponse.Customproperty35)
	target.CustomProperty36 = util.SafeString(apiResponse.Customproperty36)
	target.CustomProperty37 = util.SafeString(apiResponse.Customproperty37)
	target.CustomProperty38 = util.SafeString(apiResponse.Customproperty38)
	target.CustomProperty39 = util.SafeString(apiResponse.Customproperty39)
	target.CustomProperty40 = util.SafeString(apiResponse.Customproperty40)
	target.CustomProperty41 = util.SafeString(apiResponse.Customproperty41)
	target.CustomProperty42 = util.SafeString(apiResponse.Customproperty42)
	target.CustomProperty43 = util.SafeString(apiResponse.Customproperty43)
	target.CustomProperty44 = util.SafeString(apiResponse.Customproperty44)
	target.CustomProperty45 = util.SafeString(apiResponse.Customproperty45)
}

// SetCustomPropertyLabelsFromAPI sets custom property labels 1-60 from API response
func (r *EndpointResource) SetCustomPropertyLabelsFromAPI(target *EndpointResourceModel, apiResponse *openapi.GetEndpoints200ResponseEndpointsInner) {
	// Labels 1-30 use AccountCustomProperty pattern
	target.AccountCustomProperty1Label = util.SafeString(apiResponse.AccountCustomProperty1Label)
	target.AccountCustomProperty2Label = util.SafeString(apiResponse.AccountCustomProperty2Label)
	target.AccountCustomProperty3Label = util.SafeString(apiResponse.AccountCustomProperty3Label)
	target.AccountCustomProperty4Label = util.SafeString(apiResponse.AccountCustomProperty4Label)
	target.AccountCustomProperty5Label = util.SafeString(apiResponse.AccountCustomProperty5Label)
	target.AccountCustomProperty6Label = util.SafeString(apiResponse.AccountCustomProperty6Label)
	target.AccountCustomProperty7Label = util.SafeString(apiResponse.AccountCustomProperty7Label)
	target.AccountCustomProperty8Label = util.SafeString(apiResponse.AccountCustomProperty8Label)
	target.AccountCustomProperty9Label = util.SafeString(apiResponse.AccountCustomProperty9Label)
	target.AccountCustomProperty10Label = util.SafeString(apiResponse.AccountCustomProperty10Label)
	target.AccountCustomProperty11Label = util.SafeString(apiResponse.AccountCustomProperty11Label)
	target.AccountCustomProperty12Label = util.SafeString(apiResponse.AccountCustomProperty12Label)
	target.AccountCustomProperty13Label = util.SafeString(apiResponse.AccountCustomProperty13Label)
	target.AccountCustomProperty14Label = util.SafeString(apiResponse.AccountCustomProperty14Label)
	target.AccountCustomProperty15Label = util.SafeString(apiResponse.AccountCustomProperty15Label)
	target.AccountCustomProperty16Label = util.SafeString(apiResponse.AccountCustomProperty16Label)
	target.AccountCustomProperty17Label = util.SafeString(apiResponse.AccountCustomProperty17Label)
	target.AccountCustomProperty18Label = util.SafeString(apiResponse.AccountCustomProperty18Label)
	target.AccountCustomProperty19Label = util.SafeString(apiResponse.AccountCustomProperty19Label)
	target.AccountCustomProperty20Label = util.SafeString(apiResponse.AccountCustomProperty20Label)
	target.AccountCustomProperty21Label = util.SafeString(apiResponse.AccountCustomProperty21Label)
	target.AccountCustomProperty22Label = util.SafeString(apiResponse.AccountCustomProperty22Label)
	target.AccountCustomProperty23Label = util.SafeString(apiResponse.AccountCustomProperty23Label)
	target.AccountCustomProperty24Label = util.SafeString(apiResponse.AccountCustomProperty24Label)
	target.AccountCustomProperty25Label = util.SafeString(apiResponse.AccountCustomProperty25Label)
	target.AccountCustomProperty26Label = util.SafeString(apiResponse.AccountCustomProperty26Label)
	target.AccountCustomProperty27Label = util.SafeString(apiResponse.AccountCustomProperty27Label)
	target.AccountCustomProperty28Label = util.SafeString(apiResponse.AccountCustomProperty28Label)
	target.AccountCustomProperty29Label = util.SafeString(apiResponse.AccountCustomProperty29Label)
	target.AccountCustomProperty30Label = util.SafeString(apiResponse.AccountCustomProperty30Label)

	// Labels 31-60 use CustomProperty pattern
	target.CustomProperty31Label = util.SafeString(apiResponse.Customproperty31Label)
	target.CustomProperty32Label = util.SafeString(apiResponse.Customproperty32Label)
	target.CustomProperty33Label = util.SafeString(apiResponse.Customproperty33Label)
	target.CustomProperty34Label = util.SafeString(apiResponse.Customproperty34Label)
	target.CustomProperty35Label = util.SafeString(apiResponse.Customproperty35Label)
	target.CustomProperty36Label = util.SafeString(apiResponse.Customproperty36Label)
	target.CustomProperty37Label = util.SafeString(apiResponse.Customproperty37Label)
	target.CustomProperty38Label = util.SafeString(apiResponse.Customproperty38Label)
	target.CustomProperty39Label = util.SafeString(apiResponse.Customproperty39Label)
	target.CustomProperty40Label = util.SafeString(apiResponse.Customproperty40Label)
	target.CustomProperty41Label = util.SafeString(apiResponse.Customproperty41Label)
	target.CustomProperty42Label = util.SafeString(apiResponse.Customproperty42Label)
	target.CustomProperty43Label = util.SafeString(apiResponse.Customproperty43Label)
	target.CustomProperty44Label = util.SafeString(apiResponse.Customproperty44Label)
	target.CustomProperty45Label = util.SafeString(apiResponse.Customproperty45Label)
	target.CustomProperty46Label = util.SafeString(apiResponse.Customproperty46Label)
	target.CustomProperty47Label = util.SafeString(apiResponse.Customproperty47Label)
	target.CustomProperty48Label = util.SafeString(apiResponse.Customproperty48Label)
	target.CustomProperty49Label = util.SafeString(apiResponse.Customproperty49Label)
	target.CustomProperty50Label = util.SafeString(apiResponse.Customproperty50Label)
	target.CustomProperty51Label = util.SafeString(apiResponse.Customproperty51Label)
	target.CustomProperty52Label = util.SafeString(apiResponse.Customproperty52Label)
	target.CustomProperty53Label = util.SafeString(apiResponse.Customproperty53Label)
	target.CustomProperty54Label = util.SafeString(apiResponse.Customproperty54Label)
	target.CustomProperty55Label = util.SafeString(apiResponse.Customproperty55Label)
	target.CustomProperty56Label = util.SafeString(apiResponse.Customproperty56Label)
	target.CustomProperty57Label = util.SafeString(apiResponse.Customproperty57Label)
	target.CustomProperty58Label = util.SafeString(apiResponse.Customproperty58Label)
	target.CustomProperty59Label = util.SafeString(apiResponse.Customproperty59Label)
	target.CustomProperty60Label = util.SafeString(apiResponse.Customproperty60Label)
}

// BuildEmailTemplatesFromAPI processes email templates from API response
func (r *EndpointResource) BuildEmailTemplatesFromAPI(endpoint *openapi.GetEndpoints200ResponseEndpointsInner, diagnostics *diag.Diagnostics) types.List {
	if endpoint.Taskemailtemplates == nil || *endpoint.Taskemailtemplates == "" {
		return types.ListNull(types.ObjectType{
			AttrTypes: map[string]attr.Type{
				"email_template_type": types.StringType,
				"task_type":           types.StringType,
				"email_template":      types.StringType,
			},
		})
	}

	taskEmailTemplatesStr := *endpoint.Taskemailtemplates
	type ApiEmailTemplate struct {
		EmailTemplateType string `json:"emailTemplateType"`
		TaskType          string `json:"taskType"`
		EmailTemplate     string `json:"emailTemplate"`
	}

	var apiTemplates []ApiEmailTemplate
	if err := json.Unmarshal([]byte(taskEmailTemplatesStr), &apiTemplates); err != nil {
		diagnostics.AddError(
			"Client Error",
			fmt.Sprintf("Failed to parse taskemailtemplates JSON In buildEmailTemplatesFromAPI Block: %s", err),
		)
		return types.ListNull(types.ObjectType{
			AttrTypes: map[string]attr.Type{
				"email_template_type": types.StringType,
				"task_type":           types.StringType,
				"email_template":      types.StringType,
			},
		})
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
			diagnostics.Append(diags...)
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
		diagnostics.Append(diags...)
		return types.ListNull(types.ObjectType{
			AttrTypes: map[string]attr.Type{
				"email_template_type": types.StringType,
				"task_type":           types.StringType,
				"email_template":      types.StringType,
			},
		})
	}
	return listVal
}

// BuildRequestableRoleTypesFromAPI processes requestable role types from API response
func (r *EndpointResource) BuildRequestableRoleTypesFromAPI(endpoint *openapi.GetEndpoints200ResponseEndpointsInner, diagnostics *diag.Diagnostics) types.List {
	if endpoint.RoleTypeAsJson == nil || *endpoint.RoleTypeAsJson == "" {
		return types.ListNull(types.ObjectType{
			AttrTypes: map[string]attr.Type{
				"role_type":       types.StringType,
				"request_option":  types.StringType,
				"required":        types.BoolType,
				"requested_query": types.StringType,
				"selected_query":  types.StringType,
				"show_on":         types.StringType,
			},
		})
	}

	roleTypeJsonStr := *endpoint.RoleTypeAsJson
	var roleTypeMap map[string]string
	if err := json.Unmarshal([]byte(roleTypeJsonStr), &roleTypeMap); err != nil {
		diagnostics.AddError(
			"Client Error",
			fmt.Sprintf("Failed to parse RoleTypeAsJson outer JSON In buildRequestableRoleTypesFromAPI Block: %s", err),
		)
		return types.ListNull(types.ObjectType{
			AttrTypes: map[string]attr.Type{
				"role_type":       types.StringType,
				"request_option":  types.StringType,
				"required":        types.BoolType,
				"requested_query": types.StringType,
				"selected_query":  types.StringType,
				"show_on":         types.StringType,
			},
		})
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
				"request_option":  types.StringValue(endpointsutil.TranslateValue(get(0), endpointsutil.RequestOptionMap)),
				"required":        types.BoolValue(get(1) == "1"),
				"requested_query": types.StringValue(get(2)),
				"selected_query":  types.StringValue(get(3)),
				"show_on":         types.StringValue(endpointsutil.TranslateValue(get(4), endpointsutil.ShowOnMap)),
			},
		)
		if diags.HasError() {
			diagnostics.Append(diags...)
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
		diagnostics.Append(diags...)
		return types.ListNull(types.ObjectType{
			AttrTypes: map[string]attr.Type{
				"role_type":       types.StringType,
				"request_option":  types.StringType,
				"required":        types.BoolType,
				"requested_query": types.StringType,
				"selected_query":  types.StringType,
				"show_on":         types.StringType,
			},
		})
	}
	return listVal
}

// CreateEndpoint - Interface method for endpoint creation following Endpoint pattern
func (r *EndpointResource) CreateEndpoint(ctx context.Context, plan *EndpointResourceModel) (*openapi.UpdateEndpoint200Response, error) {
	// Check if endpoint already exists (idempotency check) with retry logic
	endpointName := plan.EndpointName.ValueString()
	var existingResource *openapi.GetEndpoints200Response

	err := r.provider.AuthenticatedAPICallWithRetry(ctx, "get_endpoints_idempotency", func(token string) error {
		endpointOps := r.endpointFactory.CreateEndpointOperations(r.client.APIBaseURL(), token)
		getReq := openapi.GetEndpointsRequest{}
		getReq.SetEndpointname(endpointName)

		resp, httpResp, err := endpointOps.GetEndpoints(ctx, getReq)
		if httpResp != nil && httpResp.StatusCode == 401 {
			return fmt.Errorf("401 unauthorized")
		}
		existingResource = resp
		return err
	})

	if err != nil {
		return nil, fmt.Errorf("failed to check existing endpoint: %w", err)
	}

	if existingResource != nil &&
		existingResource.Endpoints != nil &&
		len(existingResource.Endpoints) > 0 &&
		existingResource.ErrorCode != nil &&
		*existingResource.ErrorCode == "0" {
		log.Printf("[ERROR]: Endpoint with name '%s' already exists. Skipping creation In CreateEndpoint Block.", endpointName)
		return nil, fmt.Errorf("Endpoint with name '%s' already exists In CreateEndpoint Block", endpointName)
	}

	// Create local diagnostics for proper error handling
	var localDiags diag.Diagnostics

	// Build create request using unified helper function with proper diagnostics
	createReq := r.BuildEndpointRequest(ctx, plan, &localDiags, true).(openapi.CreateEndpointRequest)

	// Check if there were any diagnostics errors during request building
	if localDiags.HasError() {
		log.Printf("[ERROR]: Failed to build create request In CreateEndpoint Block: %v", localDiags.Errors())
		return nil, fmt.Errorf("failed to build create request In CreateEndpoint Block: %v", localDiags.Errors())
	}

	var apiResp *openapi.UpdateEndpoint200Response

	// Execute create operation with retry logic
	err = r.provider.AuthenticatedAPICallWithRetry(ctx, "create_endpoint", func(token string) error {
		endpointOps := r.endpointFactory.CreateEndpointOperations(r.client.APIBaseURL(), token)
		resp, httpResp, err := endpointOps.CreateEndpoint(ctx, createReq)
		if httpResp != nil && httpResp.StatusCode == 401 {
			return fmt.Errorf("401 unauthorized")
		}
		apiResp = resp
		return err
	})

	if err != nil {
		log.Printf("[ERROR]: Error Creating Endpoint In CreateEndpoint Block: %v", err)
		return nil, fmt.Errorf("error Creating Endpoint In CreateEndpoint Block: %w", err)
	}

	if apiResp != nil && apiResp.ErrorCode != nil && *apiResp.ErrorCode != "0" {
		msg := util.SafeDeref(apiResp.Msg)
		log.Printf("[ERROR]:API error In CreateEndpoint Block: %s", msg)
		return nil, fmt.Errorf("API error In CreateEndpoint Block: %s", msg)
	}
	return apiResp, nil
}

// GetEndpoints - Interface method for endpoint creation following Endpoint pattern
func (r *EndpointResource) GetEndpoints(ctx context.Context, plan *EndpointResourceModel) (*openapi.GetEndpoints200Response, error) {
	endpointName := plan.EndpointName.ValueString()
	var apiResp *openapi.GetEndpoints200Response

	// Execute get operation with retry logic
	err := r.provider.AuthenticatedAPICallWithRetry(ctx, "get_endpoints", func(token string) error {
		endpointOps := r.endpointFactory.CreateEndpointOperations(r.client.APIBaseURL(), token)
		apiReq := openapi.GetEndpointsRequest{}
		apiReq.SetEndpointname(endpointName)

		resp, httpResp, err := endpointOps.GetEndpoints(ctx, apiReq)
		if httpResp != nil && httpResp.StatusCode == 401 {
			return fmt.Errorf("401 unauthorized")
		}
		apiResp = resp
		return err
	})

	if err != nil {
		log.Printf("Problem with the get function in GetEndpoints block. Error: %v", err)
		return nil, fmt.Errorf("API Read Failed In GetEndpoints Block: %w", err)
	}

	if apiResp != nil && apiResp.ErrorCode != nil && *apiResp.ErrorCode != "0" {
		errorCode := util.SafeDeref(apiResp.ErrorCode)
		message := util.SafeDeref(apiResp.Message)
		log.Printf("Error Reading Endpoint In GetEndpoints Block: %v, Error code: %v", message, errorCode)
		return nil, fmt.Errorf("reading of Endpoint resource failed In GetEndpoints Block: %s", message)
	}

	if len(apiResp.Endpoints) == 0 {
		return nil, fmt.Errorf("client Error: API returned empty endpoints list")
	}
	return apiResp, nil
}

// UpdateEndpoint - Clean business logic method for endpoint updates following clean pattern
func (r *EndpointResource) UpdateEndpoint(ctx context.Context, plan *EndpointResourceModel) error {
	// Build UPDATE request using existing helper function with proper diagnostics
	var localDiags diag.Diagnostics
	updateReq := r.BuildEndpointRequest(ctx, plan, &localDiags, false).(openapi.UpdateEndpointRequest)
	if localDiags.HasError() {
		return fmt.Errorf("failed to build endpoint request In UpdateEndpoint Block: %v", localDiags.Errors())
	}

	// Handle mapped endpoints
	mappedEndpoints, err := r.BuildMappedEndpointsForUpdate(ctx, plan)
	if err != nil {
		return fmt.Errorf("failed to build mapped endpoints In UpdateEndpoint Block: %w", err)
	}

	if len(mappedEndpoints) > 0 {
		updateReq.MappedEndpoints = mappedEndpoints
	}

	// Handle requestable role types
	requestableRoleTypes, err := r.BuildRequestableRoleTypesForUpdate(ctx, plan)
	if err != nil {
		return fmt.Errorf("failed to build requestable role types In UpdateEndpoint Block: %w", err)
	}

	if len(requestableRoleTypes) > 0 {
		updateReq.RequestableRoleType = requestableRoleTypes
	}

	var apiResp *openapi.UpdateEndpoint200Response

	// Execute the update with retry logic
	err = r.provider.AuthenticatedAPICallWithRetry(ctx, "update_endpoint", func(token string) error {
		endpointOps := r.endpointFactory.CreateEndpointOperations(r.client.APIBaseURL(), token)
		resp, httpResp, err := endpointOps.UpdateEndpoint(ctx, updateReq)
		if httpResp != nil && httpResp.StatusCode == 401 {
			return fmt.Errorf("401 unauthorized")
		}
		apiResp = resp
		return err
	})

	if err != nil {
		log.Printf("[ERROR]: Problem with update function in UpdateEndpoint block. Error: %v", err)
		return fmt.Errorf("API Update Failed In UpdateEndpoint Block: %w", err)
	}

	// Check API response
	if apiResp != nil && apiResp.ErrorCode != nil && *apiResp.ErrorCode != "0" {
		errorCode := util.SafeDeref(apiResp.ErrorCode)
		msg := util.SafeDeref(apiResp.Msg)
		log.Printf("[ERROR]: Error Updating Endpoint In UpdateEndpoint Block: %v, Error code: %v", msg, errorCode)
		return fmt.Errorf("Update of Endpoint resource failed In UpdateEndpoint Block: %s", msg)
	}

	return nil
}

// BuildMappedEndpointsForUpdate - Helper method to build mapped endpoints
func (r *EndpointResource) BuildMappedEndpointsForUpdate(ctx context.Context, plan *EndpointResourceModel) ([]openapi.UpdateEndpointRequestMappedEndpointsInner, error) {
	var mappedEndpoints []openapi.UpdateEndpointRequestMappedEndpointsInner
	var tfMappedEndpoints []MappedEndpoint

	diags := plan.MappedEndpoints.ElementsAs(ctx, &tfMappedEndpoints, true)
	if diags.HasError() {
		return nil, fmt.Errorf("failed to parse mapped endpoints")
	}

	for _, tfTemplate := range tfMappedEndpoints {
		// Skip if all fields are unknown
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

	return mappedEndpoints, nil
}

// BuildRequestableRoleTypesForUpdate - Helper method to build requestable role types
func (r *EndpointResource) BuildRequestableRoleTypesForUpdate(ctx context.Context, plan *EndpointResourceModel) ([]openapi.UpdateEndpointRequestRequestableRoleTypeInner, error) {
	var requestableRoleTypes []openapi.UpdateEndpointRequestRequestableRoleTypeInner
	var tfRequestableRoleTypes []RequestableRoleType

	diags := plan.RequestableRoleTypes.ElementsAs(ctx, &tfRequestableRoleTypes, true)
	if diags.HasError() {
		return nil, fmt.Errorf("failed to parse requestable role types In buildRequestableRoleTypesForUpdate Block: %v", diags.Errors())
	}

	for _, tfTemplate := range tfRequestableRoleTypes {
		// Skip if all fields are unknown
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

	return requestableRoleTypes, nil
}

// BuildMappedEndpointsForUpdateWithDiags - Helper method that returns diagnostics instead of errors
func (r *EndpointResource) BuildMappedEndpointsForUpdateWithDiags(ctx context.Context, plan *EndpointResourceModel) ([]openapi.UpdateEndpointRequestMappedEndpointsInner, diag.Diagnostics) {
	var mappedEndpoints []openapi.UpdateEndpointRequestMappedEndpointsInner
	var tfMappedEndpoints []MappedEndpoint
	var diags diag.Diagnostics

	elementDiags := plan.MappedEndpoints.ElementsAs(ctx, &tfMappedEndpoints, true)
	diags.Append(elementDiags...)
	if diags.HasError() {
		log.Printf("[ERROR]: Failed to parse mapped endpoints In buildMappedEndpointsForUpdateWithDiags Block: %v", diags.Errors())
		return nil, diags
	}

	for _, tfTemplate := range tfMappedEndpoints {
		// Skip if all fields are unknown
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

	return mappedEndpoints, diags
}

// BuildRequestableRoleTypesForUpdateWithDiags - Helper method that returns diagnostics instead of errors
func (r *EndpointResource) BuildRequestableRoleTypesForUpdateWithDiags(ctx context.Context, plan *EndpointResourceModel) ([]openapi.UpdateEndpointRequestRequestableRoleTypeInner, diag.Diagnostics) {
	var requestableRoleTypes []openapi.UpdateEndpointRequestRequestableRoleTypeInner
	var tfRequestableRoleTypes []RequestableRoleType
	var diags diag.Diagnostics

	elementDiags := plan.RequestableRoleTypes.ElementsAs(ctx, &tfRequestableRoleTypes, true)
	diags.Append(elementDiags...)
	if diags.HasError() {
		log.Printf("[ERROR]: Failed to parse requestable role types In buildRequestableRoleTypesForUpdateWithDiags Block: %v", diags.Errors())
		return nil, diags
	}

	for _, tfTemplate := range tfRequestableRoleTypes {
		// Skip if all fields are unknown
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

	return requestableRoleTypes, diags
}

// UpdateModelFromCreateResponse - Extracted state management logic for CREATE operations
func (r *EndpointResource) UpdateModelFromCreateResponse(plan *EndpointResourceModel, apiResp *openapi.UpdateEndpoint200Response) {
	// Set the resource ID
	plan.ID = types.StringValue("endpoint-" + plan.EndpointName.ValueString())

	// Set default values for computed fields
	if plan.CreateEntTaskforRemoveAcc.IsNull() || plan.CreateEntTaskforRemoveAcc.IsUnknown() || plan.CreateEntTaskforRemoveAcc.ValueString() == "" {
		plan.CreateEntTaskforRemoveAcc = types.StringValue("false")
	}
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

	// Set default empty lists for complex fields
	if plan.EmailTemplates.IsNull() || plan.EmailTemplates.IsUnknown() {
		plan.EmailTemplates = types.ListValueMust(types.ObjectType{
			AttrTypes: map[string]attr.Type{
				"email_template_type": types.StringType,
				"task_type":           types.StringType,
				"email_template":      types.StringType,
			},
		}, []attr.Value{})
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
		}, []attr.Value{})
	}

	// Update all optional fields to maintain state (following SecuritySystem pattern)
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
	// Set API response fields
	plan.Msg = types.StringValue(util.SafeDeref(apiResp.Msg))
	plan.ErrorCode = types.StringValue(util.SafeDeref(apiResp.ErrorCode))
}
