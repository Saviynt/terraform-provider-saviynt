// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// saviynt_entitlement_type_resource manages entitlement types in the Saviynt Security Manager.
// The resource implements the full Terraform lifecycle:
//   - Create: provisions new entitlement type for an endpoint using the supplied configuration.
//   - Read: fetches the current entitlement type state from Saviynt to keep Terraform’s state in sync.
//   - Update: applies any configuration changes to an existing entitlement type.
//   - Import: brings existing entitlement type under Terraform management by its name.

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
	"terraform-provider-Saviynt/util/entitlementtypeutil"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	openapi "github.com/saviynt/saviynt-api-go-client/entitlementtype"
)

var _ resource.Resource = &EntitlementTypeResource{}
var _ resource.ResourceWithImportState = &EntitlementTypeResource{}

type EntitlementTypeResourceModel struct {
	ID                                 types.String `tfsdk:"id"`
	EntitlementName                    types.String `tfsdk:"entitlement_name"`
	EndpointName                       types.String `tfsdk:"endpoint_name"`
	DisplayName                        types.String `tfsdk:"display_name"`
	EntitlementDescription             types.String `tfsdk:"entitlement_description"`
	Workflow                           types.String `tfsdk:"workflow"`
	EnableEntitlementToRoleSync        types.Bool   `tfsdk:"enable_entitlement_to_role_sync"`
	AvailableQueryServiceAccount       types.String `tfsdk:"available_query_service_account"`
	SelectedQueryServiceAccount        types.String `tfsdk:"selected_query_service_account"`
	ArsRequestableEntitlementQuery     types.String `tfsdk:"ars_requestable_entitlement_query"`
	ArsSelectedEntitlementQuery        types.String `tfsdk:"ars_selected_entitlement_query"`
	Certifiable                        types.Bool   `tfsdk:"certifiable"`
	CreateTaskAction                   types.Set    `tfsdk:"create_task_action"` // Set of strings
	RequestDatesConfJson               types.String `tfsdk:"request_dates_conf_json"`
	StartDateInRevokeRequest           types.String `tfsdk:"start_date_in_revoke_request"`
	StartEndDateInRequest              types.String `tfsdk:"start_end_date_in_request"`
	AllowRemoveAllEntitlementInRequest types.String `tfsdk:"allow_remove_all_entitlement_in_request"`
	OrderIndex                         types.Int32  `tfsdk:"order_index"`
	RequiredInRequest                  types.Bool   `tfsdk:"required_in_request"`
	RequiredInServiceRequest           types.Bool   `tfsdk:"required_in_service_request"`
	HierarchyRequired                  types.String `tfsdk:"hierarchy_required"` //defaults to "0" in get
	ShowEntTypeOn                      types.String `tfsdk:"show_ent_type_on"`
	EnableProvisioningPriority         types.Bool   `tfsdk:"enable_provisioning_priority"`
	RequestOption                      types.String `tfsdk:"request_option"`
	Recon                              types.Bool   `tfsdk:"recon"`
	ExcludeRuleAssgnEntsInRequest      types.Bool   `tfsdk:"exclude_rule_assgn_ents_in_req"`
	CustomProperty1                    types.String `tfsdk:"custom_property1"`
	CustomProperty2                    types.String `tfsdk:"custom_property2"`
	CustomProperty3                    types.String `tfsdk:"custom_property3"`
	CustomProperty4                    types.String `tfsdk:"custom_property4"`
	CustomProperty5                    types.String `tfsdk:"custom_property5"`
	CustomProperty6                    types.String `tfsdk:"custom_property6"`
	CustomProperty7                    types.String `tfsdk:"custom_property7"`
	CustomProperty8                    types.String `tfsdk:"custom_property8"`
	CustomProperty9                    types.String `tfsdk:"custom_property9"`
	CustomProperty10                   types.String `tfsdk:"custom_property10"`
	CustomProperty11                   types.String `tfsdk:"custom_property11"`
	CustomProperty12                   types.String `tfsdk:"custom_property12"`
	CustomProperty13                   types.String `tfsdk:"custom_property13"`
	CustomProperty14                   types.String `tfsdk:"custom_property14"`
	CustomProperty15                   types.String `tfsdk:"custom_property15"`
	CustomProperty16                   types.String `tfsdk:"custom_property16"`
	CustomProperty17                   types.String `tfsdk:"custom_property17"`
	CustomProperty18                   types.String `tfsdk:"custom_property18"`
	CustomProperty19                   types.String `tfsdk:"custom_property19"`
	CustomProperty20                   types.String `tfsdk:"custom_property20"`
	CustomProperty21                   types.String `tfsdk:"custom_property21"`
	CustomProperty22                   types.String `tfsdk:"custom_property22"`
	CustomProperty23                   types.String `tfsdk:"custom_property23"`
	CustomProperty24                   types.String `tfsdk:"custom_property24"`
	CustomProperty25                   types.String `tfsdk:"custom_property25"`
	CustomProperty26                   types.String `tfsdk:"custom_property26"`
	CustomProperty27                   types.String `tfsdk:"custom_property27"`
	CustomProperty28                   types.String `tfsdk:"custom_property28"`
	CustomProperty29                   types.String `tfsdk:"custom_property29"`
	CustomProperty30                   types.String `tfsdk:"custom_property30"`
	CustomProperty31                   types.String `tfsdk:"custom_property31"`
	CustomProperty32                   types.String `tfsdk:"custom_property32"`
	CustomProperty33                   types.String `tfsdk:"custom_property33"`
	CustomProperty34                   types.String `tfsdk:"custom_property34"`
	CustomProperty35                   types.String `tfsdk:"custom_property35"`
	CustomProperty36                   types.String `tfsdk:"custom_property36"`
	CustomProperty37                   types.String `tfsdk:"custom_property37"`
	CustomProperty38                   types.String `tfsdk:"custom_property38"`
	CustomProperty39                   types.String `tfsdk:"custom_property39"`
	CustomProperty40                   types.String `tfsdk:"custom_property40"`
}

type EntitlementTypeResource struct {
	client                 client.SaviyntClientInterface
	token                  string
	entitlementTypeFactory client.EntitlementTypeFactoryInterface
}

func NewEntitlementTypeResource() resource.Resource {
	return &EntitlementTypeResource{
		entitlementTypeFactory: &client.DefaultEntitlementTypeFactory{},
	}
}

func NewEntitlementTypeResourceWithFactory(factory client.EntitlementTypeFactoryInterface) resource.Resource {
	return &EntitlementTypeResource{
		entitlementTypeFactory: factory,
	}
}

func (r *EntitlementTypeResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = "saviynt_entitlement_type_resource"
}

func (r *EntitlementTypeResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: util.EntitlementTypeDescription,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The unique ID of the entitlement type resource, typically assigned by the API.",
			},
			"entitlement_name": schema.StringAttribute{
				Required:    true,
				Description: "Name of the entitlement type.",
			},
			"endpoint_name": schema.StringAttribute{
				Required:    true,
				Description: "Name of the endpoint with which the entitlement type is associated.",
			},
			"display_name": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Display name for the entitlement type",
			},
			"entitlement_description": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Description for the entitlement type.",
			},
			"workflow": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Workflow associated with the entitlement type.",
			},
			"enable_entitlement_to_role_sync": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enable entitlement to role sync. Can only be set to true when workflow is also specified.",
				PlanModifiers: []planmodifier.Bool{
					entitlementtypeutil.RequireWorkflowWhenEnabled(),
				},
			},
			"available_query_service_account": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Query to fetch available service accounts.",
			},
			"selected_query_service_account": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Query to fetch selected service accounts.",
			},
			"ars_requestable_entitlement_query": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Query used to determine requestable entitlements.",
			},
			"ars_selected_entitlement_query": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Query used to determine selected entitlements.",
			},
			"certifiable": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Default:     booldefault.StaticBool(true),
				Description: "Indicates if the entitlement is certifiable.",
			},
			"create_task_action": schema.SetAttribute{
				Optional:    true,
				ElementType: types.StringType,
				Computed:    true,
				Description: "Action(s) to be performed when a task is created.",
				PlanModifiers: []planmodifier.Set{
					entitlementtypeutil.CreateTaskActionDefault(),
				},
			},
			"request_dates_conf_json": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Configuration in JSON for handling request dates.",
			},
			"start_date_in_revoke_request": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Include start date in revoke requests.",
			},
			"start_end_date_in_request": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Include both start and end dates in the request.",
			},
			"allow_remove_all_entitlement_in_request": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Allow user to remove all entitlements in one request.",
			},

			"order_index": schema.Int32Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Index to determine the order of processing or display.",
			},
			"required_in_request": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Flag indicating if this field is required in the request.",
			},
			"required_in_service_request": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Flag indicating if this field is required in service request.",
			},
			"hierarchy_required": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Flag indicating if a hierarchy is required.",
			},
			"show_ent_type_on": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Show entitlement type on.",
			},
			"enable_provisioning_priority": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enable provisioning priority.",
			},
			"request_option": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Defines how the entitlement should be presented or requested.",
			},
			"recon": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Recon",
			},
			"exclude_rule_assgn_ents_in_req": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Exclude Entitlements Assigned via Rule while Request",
			},
			"custom_property1": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Custom property 1 associated with the entitlement.",
			},
			"custom_property2": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Custom property 2 associated with the entitlement.",
			},
			"custom_property3": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Custom property 3 associated with the entitlement.",
			},
			"custom_property4": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Custom property 4 associated with the entitlement.",
			},
			"custom_property5": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Custom property 5 associated with the entitlement.",
			},
			"custom_property6": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Custom property 6 associated with the entitlement.",
			},
			"custom_property7": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Custom property 7 associated with the entitlement.",
			},
			"custom_property8": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Custom property 8 associated with the entitlement.",
			},
			"custom_property9": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Custom property 9 associated with the entitlement.",
			},
			"custom_property10": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Custom property 10 associated with the entitlement.",
			},
			"custom_property11": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Custom property 11 associated with the entitlement.",
			},
			"custom_property12": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Custom property 12 associated with the entitlement.",
			},
			"custom_property13": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Custom property 13 associated with the entitlement.",
			},
			"custom_property14": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Custom property 14 associated with the entitlement.",
			},
			"custom_property15": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Custom property 15 associated with the entitlement.",
			},
			"custom_property16": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Custom property 16 associated with the entitlement.",
			},
			"custom_property17": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Custom property 17 associated with the entitlement.",
			},
			"custom_property18": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Custom property 18 associated with the entitlement.",
			},
			"custom_property19": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Custom property 19 associated with the entitlement.",
			},
			"custom_property20": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Custom property 20 associated with the entitlement.",
			},
			"custom_property21": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Custom property 21 associated with the entitlement.",
			},
			"custom_property22": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Custom property 22 associated with the entitlement.",
			},
			"custom_property23": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Custom property 23 associated with the entitlement.",
			},
			"custom_property24": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Custom property 24 associated with the entitlement.",
			},
			"custom_property25": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Custom property 25 associated with the entitlement.",
			},
			"custom_property26": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Custom property 26 associated with the entitlement.",
			},
			"custom_property27": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Custom property 27 associated with the entitlement.",
			},
			"custom_property28": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Custom property 28 associated with the entitlement.",
			},
			"custom_property29": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Custom property 29 associated with the entitlement.",
			},
			"custom_property30": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Custom property 30 associated with the entitlement.",
			},
			"custom_property31": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Custom property 31 associated with the entitlement.",
			},
			"custom_property32": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Custom property 32 associated with the entitlement.",
			},
			"custom_property33": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Custom property 33 associated with the entitlement.",
			},
			"custom_property34": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Custom property 34 associated with the entitlement.",
			},
			"custom_property35": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Custom property 35 associated with the entitlement.",
			},
			"custom_property36": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Custom property 36 associated with the entitlement.",
			},
			"custom_property37": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Custom property 37 associated with the entitlement.",
			},
			"custom_property38": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Custom property 38 associated with the entitlement.",
			},
			"custom_property39": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Custom property 39 associated with the entitlement.",
			},
			"custom_property40": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Custom property 40 associated with the entitlement.",
			},
		},
	}
}

func (r *EntitlementTypeResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Check if provider data is available.
	if req.ProviderData == nil {
		log.Println("ProviderData is nil, returning early.")
		return
	}

	// Cast provider data to your provider type.
	prov, ok := req.ProviderData.(*saviyntProvider)
	if !ok {
		log.Println("[ERROR] Provider: Unexpected provider data")
		resp.Diagnostics.AddError("Unexpected Provider Data", "Expected *saviyntProvider")
		return
	}

	// Set the client and token from the provider state using interface wrapper.
	r.client = &client.SaviyntClientWrapper{Client: prov.client}
	r.token = prov.accessToken
	log.Println("[DEBUG] EntitlementType: Resource configured successfully")
}

// SetClient sets the client for testing purposes
func (r *EntitlementTypeResource) SetClient(client client.SaviyntClientInterface) {
	r.client = client
}

// SetToken sets the token for testing purposes
func (r *EntitlementTypeResource) SetToken(token string) {
	r.token = token
}

// CreateEntitlementType creates an entitlement type
func (r *EntitlementTypeResource) CreateEntitlementType(ctx context.Context, plan *EntitlementTypeResourceModel) (*openapi.CreateOrUpdateEntitlementTypeResponse, error) {
	log.Printf("[DEBUG] EntitlementType: Starting creation for entitlement type: %s", plan.EntitlementName.ValueString())

	entitlementTypeOps := r.entitlementTypeFactory.CreateEntitlementTypeOperations(r.client.APIBaseURL(), r.token)

	existingResource, _, _ := entitlementTypeOps.GetEntitlementType(ctx, plan.EntitlementName.ValueString(), "", "", plan.EndpointName.ValueString())
	if existingResource != nil &&
		existingResource.EntitlementTypeDetails != nil &&
		len(existingResource.EntitlementTypeDetails) > 0 {
		log.Printf("[ERROR]: Entitlement type with name '%s' already exists for the given endpoint. Skipping creation.", plan.EntitlementName.ValueString())
		return nil, fmt.Errorf("entitlement type with name '%s' already exists for the endpoint: %s", plan.EntitlementName.ValueString(), plan.EndpointName.ValueString())
	}

	// Build entitlement type create request
	createReq := r.BuildCreateEntitlementTypeRequest(plan)

	createResp, _, err := entitlementTypeOps.CreateEntitlementType(ctx, createReq)
	if err != nil {
		return nil, fmt.Errorf("API call failed: %w", err)
	}

	if createResp != nil && createResp.ErrorCode != nil && *createResp.ErrorCode != "0" {
		log.Printf("[ERROR]: Error in creating Entitlement Type resource. Errorcode: %v, Message: %v", *createResp.ErrorCode, *createResp.Msg)
		return nil, fmt.Errorf("creating of Entitlement Type resource failed: %s", *createResp.Msg)
	}

	return createResp, nil
}

// UpdateEntitlementType updates an entitlement type
func (r *EntitlementTypeResource) UpdateEntitlementType(ctx context.Context, plan *EntitlementTypeResourceModel) (*openapi.CreateOrUpdateEntitlementTypeResponse, error) {
	log.Printf("[DEBUG] EntitlementType: Starting update for entitlement type: %s in endpoint: %s", plan.EntitlementName.ValueString(), plan.EndpointName.ValueString())

	entitlementTypeOps := r.entitlementTypeFactory.CreateEntitlementTypeOperations(r.client.APIBaseURL(), r.token)

	updateReq := r.BuildUpdateEntitlementTypeRequest(plan)

	updateResp, _, err := entitlementTypeOps.UpdateEntitlementType(ctx, updateReq)
	if err != nil {
		return nil, fmt.Errorf("API call failed: %w", err)
	}

	if updateResp != nil && updateResp.ErrorCode != nil && *updateResp.ErrorCode == "1" && strings.Contains(strings.ToLower(util.SafeDeref(updateResp.Msg)), "not found") {
		return nil, fmt.Errorf("entitlement type %s doesn't exist for the given endpoint", plan.EntitlementName.ValueString())
	}

	if updateResp != nil && updateResp.ErrorCode != nil && *updateResp.ErrorCode != "0" {
		log.Printf("[ERROR]: Error in updating Entitlement Type resource. Errorcode: %v, Message: %v", *updateResp.ErrorCode, *updateResp.Msg)
		return nil, fmt.Errorf("updating of Entitlement Type resource failed: %s", *updateResp.Msg)
	}

	return updateResp, nil
}

// BuildUpdateEntitlementTypeRequest - Extracted update request building logic
func (r *EntitlementTypeResource) BuildUpdateEntitlementTypeRequest(plan *EntitlementTypeResourceModel) openapi.UpdateEntitlementTypeRequest {
	updateReq := openapi.UpdateEntitlementTypeRequest{
		Entitlementname:                    plan.EntitlementName.ValueString(),
		Endpointname:                       plan.EndpointName.ValueString(),
		DisplayName:                        util.StringPointerOrEmpty(plan.DisplayName),
		Entitlementdescription:             util.StringPointerOrEmpty(plan.EntitlementDescription),
		Workflow:                           util.StringPointerOrEmpty(plan.Workflow),
		Certifiable:                        util.BoolPointerOrEmtpy(plan.Certifiable),
		ShowEntTypeOn:                      util.StringPointerOrEmpty(plan.ShowEntTypeOn),
		ArsRequestableEntitlementQuery:     util.StringPointerOrEmpty(plan.ArsRequestableEntitlementQuery),
		ArsSelectedEntitlementQuery:        util.StringPointerOrEmpty(plan.ArsSelectedEntitlementQuery),
		AvailableQueryServiceAccount:       util.StringPointerOrEmpty(plan.AvailableQueryServiceAccount),
		SelectedQueryServiceAccount:        util.StringPointerOrEmpty(plan.SelectedQueryServiceAccount),
		Orderindex:                         util.Int32PointerOrEmpty(plan.OrderIndex),
		EnableEntitlementToRoleSync:        util.BoolPointerOrEmtpy(plan.EnableEntitlementToRoleSync),
		Hiearchyrequired:                   entitlementtypeutil.ConvertHierarchyRequiredForUpdate(plan.HierarchyRequired),
		RequestDatesConfJson:               util.StringPointerOrEmpty(plan.RequestDatesConfJson),
		Requiredinrequest:                  entitlementtypeutil.ConvertBoolToStringForUpdate(plan.RequiredInRequest),
		Requiredinservicerequest:           util.BoolPointerOrEmpty(plan.RequiredInServiceRequest),
		Requestoption:                      util.StringPointerOrEmpty(plan.RequestOption),
		CreateTaskAction:                   util.ConvertTFSetToGoStrings(plan.CreateTaskAction),
		Recon:                              util.BoolPointerOrEmtpy(plan.Recon),
		ExcludeRuleAssgnEntsInRequest:      util.BoolPointerOrEmtpy(plan.ExcludeRuleAssgnEntsInRequest),
		EnableProvisioningPriority:         util.BoolPointerOrEmpty(plan.EnableProvisioningPriority),
		StartDateInRevokeRequest:           util.StringPointerOrEmpty(plan.StartDateInRevokeRequest),
		StartEndDateInRequest:              util.StringPointerOrEmpty(plan.StartEndDateInRequest),
		AllowRemoveAllEntitlementInRequest: util.StringPointerOrEmpty(plan.AllowRemoveAllEntitlementInRequest),
		Customproperty1:                    util.StringPointerOrEmpty(plan.CustomProperty1),
		Customproperty2:                    util.StringPointerOrEmpty(plan.CustomProperty2),
		Customproperty3:                    util.StringPointerOrEmpty(plan.CustomProperty3),
		Customproperty4:                    util.StringPointerOrEmpty(plan.CustomProperty4),
		Customproperty5:                    util.StringPointerOrEmpty(plan.CustomProperty5),
		Customproperty6:                    util.StringPointerOrEmpty(plan.CustomProperty6),
		Customproperty7:                    util.StringPointerOrEmpty(plan.CustomProperty7),
		Customproperty8:                    util.StringPointerOrEmpty(plan.CustomProperty8),
		Customproperty9:                    util.StringPointerOrEmpty(plan.CustomProperty9),
		Customproperty10:                   util.StringPointerOrEmpty(plan.CustomProperty10),
		Customproperty11:                   util.StringPointerOrEmpty(plan.CustomProperty11),
		Customproperty12:                   util.StringPointerOrEmpty(plan.CustomProperty12),
		Customproperty13:                   util.StringPointerOrEmpty(plan.CustomProperty13),
		Customproperty14:                   util.StringPointerOrEmpty(plan.CustomProperty14),
		Customproperty15:                   util.StringPointerOrEmpty(plan.CustomProperty15),
		Customproperty16:                   util.StringPointerOrEmpty(plan.CustomProperty16),
		Customproperty17:                   util.StringPointerOrEmpty(plan.CustomProperty17),
		Customproperty18:                   util.StringPointerOrEmpty(plan.CustomProperty18),
		Customproperty19:                   util.StringPointerOrEmpty(plan.CustomProperty19),
		Customproperty20:                   util.StringPointerOrEmpty(plan.CustomProperty20),
		Customproperty21:                   util.StringPointerOrEmpty(plan.CustomProperty21),
		Customproperty22:                   util.StringPointerOrEmpty(plan.CustomProperty22),
		Customproperty23:                   util.StringPointerOrEmpty(plan.CustomProperty23),
		Customproperty24:                   util.StringPointerOrEmpty(plan.CustomProperty24),
		Customproperty25:                   util.StringPointerOrEmpty(plan.CustomProperty25),
		Customproperty26:                   util.StringPointerOrEmpty(plan.CustomProperty26),
		Customproperty27:                   util.StringPointerOrEmpty(plan.CustomProperty27),
		Customproperty28:                   util.StringPointerOrEmpty(plan.CustomProperty28),
		Customproperty29:                   util.StringPointerOrEmpty(plan.CustomProperty29),
		Customproperty30:                   util.StringPointerOrEmpty(plan.CustomProperty30),
		Customproperty31:                   util.StringPointerOrEmpty(plan.CustomProperty31),
		Customproperty32:                   util.StringPointerOrEmpty(plan.CustomProperty32),
		Customproperty33:                   util.StringPointerOrEmpty(plan.CustomProperty33),
		Customproperty34:                   util.StringPointerOrEmpty(plan.CustomProperty34),
		Customproperty35:                   util.StringPointerOrEmpty(plan.CustomProperty35),
		Customproperty36:                   util.StringPointerOrEmpty(plan.CustomProperty36),
		Customproperty37:                   util.StringPointerOrEmpty(plan.CustomProperty37),
		Customproperty38:                   util.StringPointerOrEmpty(plan.CustomProperty38),
		Customproperty39:                   util.StringPointerOrEmpty(plan.CustomProperty39),
		Customproperty40:                   util.StringPointerOrEmpty(plan.CustomProperty40),
	}

	if reqJson, err := json.MarshalIndent(updateReq, "", "  "); err == nil {
		log.Printf("[DEBUG] BuildUpdateEntitlementTypeRequest - Full request body: %s", string(reqJson))
	} else {
		log.Printf("[DEBUG] BuildUpdateEntitlementTypeRequest - Failed to marshal request: %v", err)
	}

	return updateReq
}

// ReadEntitlementTypeState reads the current state from the API
func (r *EntitlementTypeResource) ReadEntitlementTypeState(ctx context.Context, plan *EntitlementTypeResourceModel) error {
	entitlementTypeOps := r.entitlementTypeFactory.CreateEntitlementTypeOperations(r.client.APIBaseURL(), r.token)

	readResp, _, err := entitlementTypeOps.GetEntitlementType(ctx, plan.EntitlementName.ValueString(), "", "", plan.EndpointName.ValueString())
	if err != nil {
		return fmt.Errorf("API call failed: %w", err)
	}

	if readResp != nil && readResp.DisplayCount != nil && *readResp.DisplayCount == 0 {
		return fmt.Errorf("entitlement Type not found")
	}

	if readResp != nil && readResp.ErrorCode != nil && *readResp.ErrorCode != "0" {
		return fmt.Errorf("API Read Failed with error code %s and message: %s", *readResp.ErrorCode, *readResp.Msg)
	}

	if readResp != nil && len(readResp.EntitlementTypeDetails) == 0 {
		log.Printf("[DEBUG] EntitlementType: No entitlement type details found for endpoint %s", plan.EndpointName.ValueString())
		return fmt.Errorf("no entitlement type %s found for endpoint %s", plan.EntitlementName.ValueString(), plan.EndpointName.ValueString())
	}

	// Update plan with API response data
	plan.ID = types.StringValue(fmt.Sprintf("ent-type-%s-%s", plan.EndpointName.ValueString(), plan.EntitlementName.ValueString()))

	var enable *bool
	var workflow *string

	if readResp.EntitlementTypeDetails[0].Workflow != nil && *readResp.EntitlementTypeDetails[0].Workflow != "" {
		var parsed map[string]json.RawMessage
		err := json.Unmarshal([]byte(*readResp.EntitlementTypeDetails[0].Workflow), &parsed)
		if err == nil {
			// Newer format: JSON object with nested fields
			if rawEnable, ok := parsed["enableEntitlementToRoleSync"]; ok {
				var val bool
				if err := json.Unmarshal(rawEnable, &val); err == nil {
					enable = &val
				}
			}
			if rawWorkflow, ok := parsed["workflow"]; ok {
				var val string
				if err := json.Unmarshal(rawWorkflow, &val); err == nil {
					workflow = &val
				}
			}
		} else {
			// Older format: simple string value
			workflow = readResp.EntitlementTypeDetails[0].Workflow
			enable = &[]bool{false}[0]
		}
	}

	plan.EnableEntitlementToRoleSync = util.SafeBoolDatasource(enable)
	plan.Workflow = util.SafeString(workflow)

	plan.DisplayName = util.SafeString(readResp.EntitlementTypeDetails[0].DisplayName)
	plan.AvailableQueryServiceAccount = util.SafeString(readResp.EntitlementTypeDetails[0].AvailableQueryServiceAccount)
	plan.SelectedQueryServiceAccount = util.SafeString(readResp.EntitlementTypeDetails[0].SelectedQueryServiceAccount)
	plan.ArsRequestableEntitlementQuery = util.SafeString(readResp.EntitlementTypeDetails[0].ArsReqEntSqlquerey)
	plan.ArsSelectedEntitlementQuery = util.SafeString(readResp.EntitlementTypeDetails[0].ArsSelectEntSqlquerey)
	plan.RequestDatesConfJson = util.SafeString(readResp.EntitlementTypeDetails[0].RequestDatesConfJson)
	plan.RequestOption = util.SafeString(entitlementtypeutil.TranslateValueWithDefault(*readResp.EntitlementTypeDetails[0].Requestoption))
	plan.RequiredInRequest = util.SafeBoolDatasource(util.ParseBoolPointerFromStringPointer(readResp.EntitlementTypeDetails[0].Requiredinrequest))
	plan.RequiredInServiceRequest = util.SafeBoolDatasource(util.ParseBoolPointerFromStringPointer(readResp.EntitlementTypeDetails[0].Requiredinservicerequest))
	plan.HierarchyRequired = util.SafeStringDatasource(readResp.EntitlementTypeDetails[0].Hiearchyrequired)
	plan.OrderIndex = util.SafeInt32(util.StringPtrToInt32Ptr(readResp.EntitlementTypeDetails[0].Orderindex))
	plan.EnableProvisioningPriority = util.SafeBoolDatasource(util.ParseBoolPointerFromStringPointer(readResp.EntitlementTypeDetails[0].EnableProvisioningPriority))

	parsedActions := entitlementtypeutil.ParseCreateTaskActionForState(readResp.EntitlementTypeDetails[0].CreateTaskAction)
	if parsedActions != nil {
		// Field returned by API - use actual values (could be empty array or populated)
		plan.CreateTaskAction = basetypes.NewSetValueMust(types.StringType, parsedActions)
	}
	// If parsedActions is nil, leave the field as-is - plan modifier handles defaults

	plan.Certifiable = util.SafeBoolDatasource(util.ParseBoolPointerFromStringPointer(readResp.EntitlementTypeDetails[0].Certifiable))
	plan.ShowEntTypeOn = util.SafeStringDatasource(readResp.EntitlementTypeDetails[0].ShowEntTypeOn)
	plan.EntitlementDescription = util.SafeString(readResp.EntitlementTypeDetails[0].Entitlementdescription)
	plan.Recon = util.SafeBoolDatasource(util.ParseBoolPointerFromStringPointer(readResp.EntitlementTypeDetails[0].Recon))
	plan.ExcludeRuleAssgnEntsInRequest = util.SafeBoolDatasource(util.ParseBoolPointerFromStringPointer(readResp.EntitlementTypeDetails[0].ExcludeRuleAssgnEntsInRequest))
	plan.StartDateInRevokeRequest = util.SafeStringDatasource(plan.StartDateInRevokeRequest.ValueStringPointer())
	plan.StartEndDateInRequest = util.SafeStringDatasource(plan.StartEndDateInRequest.ValueStringPointer())
	plan.AllowRemoveAllEntitlementInRequest = util.SafeStringDatasource(plan.AllowRemoveAllEntitlementInRequest.ValueStringPointer())

	plan.CustomProperty1 = util.SafeString(readResp.EntitlementTypeDetails[0].Customproperty1Label)
	plan.CustomProperty2 = util.SafeString(readResp.EntitlementTypeDetails[0].Customproperty2Label)
	plan.CustomProperty3 = util.SafeString(readResp.EntitlementTypeDetails[0].Customproperty3Label)
	plan.CustomProperty4 = util.SafeString(readResp.EntitlementTypeDetails[0].Customproperty4Label)
	plan.CustomProperty5 = util.SafeString(readResp.EntitlementTypeDetails[0].Customproperty5Label)
	plan.CustomProperty6 = util.SafeString(readResp.EntitlementTypeDetails[0].Customproperty6Label)
	plan.CustomProperty7 = util.SafeString(readResp.EntitlementTypeDetails[0].Customproperty7Label)
	plan.CustomProperty8 = util.SafeString(readResp.EntitlementTypeDetails[0].Customproperty8Label)
	plan.CustomProperty9 = util.SafeString(readResp.EntitlementTypeDetails[0].Customproperty9Label)
	plan.CustomProperty10 = util.SafeString(readResp.EntitlementTypeDetails[0].Customproperty10Label)
	plan.CustomProperty11 = util.SafeString(readResp.EntitlementTypeDetails[0].Customproperty11Label)
	plan.CustomProperty12 = util.SafeString(readResp.EntitlementTypeDetails[0].Customproperty12Label)
	plan.CustomProperty13 = util.SafeString(readResp.EntitlementTypeDetails[0].Customproperty13Label)
	plan.CustomProperty14 = util.SafeString(readResp.EntitlementTypeDetails[0].Customproperty14Label)
	plan.CustomProperty15 = util.SafeString(readResp.EntitlementTypeDetails[0].Customproperty15Label)
	plan.CustomProperty16 = util.SafeString(readResp.EntitlementTypeDetails[0].Customproperty16Label)
	plan.CustomProperty17 = util.SafeString(readResp.EntitlementTypeDetails[0].Customproperty17Label)
	plan.CustomProperty18 = util.SafeString(readResp.EntitlementTypeDetails[0].Customproperty18Label)
	plan.CustomProperty19 = util.SafeString(readResp.EntitlementTypeDetails[0].Customproperty19Label)
	plan.CustomProperty20 = util.SafeString(readResp.EntitlementTypeDetails[0].Customproperty20Label)
	plan.CustomProperty21 = util.SafeString(readResp.EntitlementTypeDetails[0].Customproperty21Label)
	plan.CustomProperty22 = util.SafeString(readResp.EntitlementTypeDetails[0].Customproperty22Label)
	plan.CustomProperty23 = util.SafeString(readResp.EntitlementTypeDetails[0].Customproperty23Label)
	plan.CustomProperty24 = util.SafeString(readResp.EntitlementTypeDetails[0].Customproperty24Label)
	plan.CustomProperty25 = util.SafeString(readResp.EntitlementTypeDetails[0].Customproperty25Label)
	plan.CustomProperty26 = util.SafeString(readResp.EntitlementTypeDetails[0].Customproperty26Label)
	plan.CustomProperty27 = util.SafeString(readResp.EntitlementTypeDetails[0].Customproperty27Label)
	plan.CustomProperty28 = util.SafeString(readResp.EntitlementTypeDetails[0].Customproperty28Label)
	plan.CustomProperty29 = util.SafeString(readResp.EntitlementTypeDetails[0].Customproperty29Label)
	plan.CustomProperty30 = util.SafeString(readResp.EntitlementTypeDetails[0].Customproperty30Label)
	plan.CustomProperty31 = util.SafeString(readResp.EntitlementTypeDetails[0].Customproperty31Label)
	plan.CustomProperty32 = util.SafeString(readResp.EntitlementTypeDetails[0].Customproperty32Label)
	plan.CustomProperty33 = util.SafeString(readResp.EntitlementTypeDetails[0].Customproperty33Label)
	plan.CustomProperty34 = util.SafeString(readResp.EntitlementTypeDetails[0].Customproperty34Label)
	plan.CustomProperty35 = util.SafeString(readResp.EntitlementTypeDetails[0].Customproperty35Label)
	plan.CustomProperty36 = util.SafeString(readResp.EntitlementTypeDetails[0].Customproperty36Label)
	plan.CustomProperty37 = util.SafeString(readResp.EntitlementTypeDetails[0].Customproperty37Label)
	plan.CustomProperty38 = util.SafeString(readResp.EntitlementTypeDetails[0].Customproperty38Label)
	plan.CustomProperty39 = util.SafeString(readResp.EntitlementTypeDetails[0].Customproperty39Label)
	plan.CustomProperty40 = util.SafeString(readResp.EntitlementTypeDetails[0].Customproperty40Label)

	return nil
}

// BuildCreateEntitlementTypeRequest - Extracted request building logic
func (r *EntitlementTypeResource) BuildCreateEntitlementTypeRequest(plan *EntitlementTypeResourceModel) openapi.CreateEntitlementTypeRequest {
	createReq := openapi.CreateEntitlementTypeRequest{
		Entitlementname:               plan.EntitlementName.ValueString(),
		Endpointname:                  plan.EndpointName.ValueString(),
		DisplayName:                   plan.DisplayName.ValueStringPointer(),
		Entitlementdescription:        plan.EntitlementDescription.ValueStringPointer(),
		Workflow:                      plan.Workflow.ValueStringPointer(),
		Certifiable:                   plan.Certifiable.ValueBoolPointer(),
		Orderindex:                    plan.OrderIndex.ValueInt32Pointer(),
		EnableEntitlementToRoleSync:   plan.EnableEntitlementToRoleSync.ValueBoolPointer(),
		EnableProvisioningPriority:    plan.EnableProvisioningPriority.ValueBoolPointer(),
		ShowEntTypeOn:                 util.StringPointerOrEmpty(plan.ShowEntTypeOn),
		AvailableQueryServiceAccount:  util.StringPointerOrEmpty(plan.AvailableQueryServiceAccount),
		SelectedQueryServiceAccount:   util.StringPointerOrEmpty(plan.SelectedQueryServiceAccount),
		Hiearchyrequired:              util.StringPtrToInt32Ptr(plan.HierarchyRequired.ValueStringPointer()),
		RequestDatesConfJson:          util.StringPointerOrEmpty(plan.RequestDatesConfJson),
		Requiredinrequest:             util.BoolPointerOrEmtpy(plan.RequiredInRequest),
		Requiredinservicerequest:      util.BoolPointerOrEmpty(plan.RequiredInServiceRequest),
		Recon:                         util.BoolPointerOrEmtpy(plan.Recon),
		ExcludeRuleAssgnEntsInRequest: util.BoolPointerOrEmtpy(plan.ExcludeRuleAssgnEntsInRequest),
		// Include all custom properties (these are supported in create)
		Customproperty1:  plan.CustomProperty1.ValueStringPointer(),
		Customproperty2:  plan.CustomProperty2.ValueStringPointer(),
		Customproperty3:  plan.CustomProperty3.ValueStringPointer(),
		Customproperty4:  plan.CustomProperty4.ValueStringPointer(),
		Customproperty5:  plan.CustomProperty5.ValueStringPointer(),
		Customproperty6:  plan.CustomProperty6.ValueStringPointer(),
		Customproperty7:  plan.CustomProperty7.ValueStringPointer(),
		Customproperty8:  plan.CustomProperty8.ValueStringPointer(),
		Customproperty9:  plan.CustomProperty9.ValueStringPointer(),
		Customproperty10: plan.CustomProperty10.ValueStringPointer(),
		Customproperty11: plan.CustomProperty11.ValueStringPointer(),
		Customproperty12: plan.CustomProperty12.ValueStringPointer(),
		Customproperty13: plan.CustomProperty13.ValueStringPointer(),
		Customproperty14: plan.CustomProperty14.ValueStringPointer(),
		Customproperty15: plan.CustomProperty15.ValueStringPointer(),
		Customproperty16: plan.CustomProperty16.ValueStringPointer(),
		Customproperty17: plan.CustomProperty17.ValueStringPointer(),
		Customproperty18: plan.CustomProperty18.ValueStringPointer(),
		Customproperty19: plan.CustomProperty19.ValueStringPointer(),
		Customproperty20: plan.CustomProperty20.ValueStringPointer(),
		Customproperty21: plan.CustomProperty21.ValueStringPointer(),
		Customproperty22: plan.CustomProperty22.ValueStringPointer(),
		Customproperty23: plan.CustomProperty23.ValueStringPointer(),
		Customproperty24: plan.CustomProperty24.ValueStringPointer(),
		Customproperty25: plan.CustomProperty25.ValueStringPointer(),
		Customproperty26: plan.CustomProperty26.ValueStringPointer(),
		Customproperty27: plan.CustomProperty27.ValueStringPointer(),
		Customproperty28: plan.CustomProperty28.ValueStringPointer(),
		Customproperty29: plan.CustomProperty29.ValueStringPointer(),
		Customproperty30: plan.CustomProperty30.ValueStringPointer(),
		Customproperty31: plan.CustomProperty31.ValueStringPointer(),
		Customproperty32: plan.CustomProperty32.ValueStringPointer(),
		Customproperty33: plan.CustomProperty33.ValueStringPointer(),
		Customproperty34: plan.CustomProperty34.ValueStringPointer(),
		Customproperty35: plan.CustomProperty35.ValueStringPointer(),
		Customproperty36: plan.CustomProperty36.ValueStringPointer(),
		Customproperty37: plan.CustomProperty37.ValueStringPointer(),
		Customproperty38: plan.CustomProperty38.ValueStringPointer(),
		Customproperty39: plan.CustomProperty39.ValueStringPointer(),
		Customproperty40: plan.CustomProperty40.ValueStringPointer(),
	}
	if reqJson, err := json.MarshalIndent(createReq, "", "  "); err == nil {
		log.Printf("[DEBUG] BuildCreateEntitlementTypeRequest - Full request body: %s", string(reqJson))
	} else {
		log.Printf("[DEBUG] BuildCreateEntitlementTypeRequest - Failed to marshal request: %v", err)
	}
	return createReq
}

// UpdateModelAfterCreate updates the model from create response
func (r *EntitlementTypeResource) UpdateModelAfterCreate(plan *EntitlementTypeResourceModel, createResp *openapi.CreateOrUpdateEntitlementTypeResponse) {
	plan.ID = types.StringValue(fmt.Sprintf("ent-type-%s-%s", plan.EndpointName.ValueString(), plan.EntitlementName.ValueString()))

	// Normalize state after both create and potential update
	log.Printf("[DEBUG] Normalizing state after create/update...")
	r.NormalizeEntitlementTypeState(plan)

	// Set all attributes with proper defaults if they're not provided
	plan.EntitlementDescription = util.SafeString(plan.EntitlementDescription.ValueStringPointer())
	plan.EnableProvisioningPriority = util.SafeBoolDatasource(plan.EnableProvisioningPriority.ValueBoolPointer())
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
}

func (r *EntitlementTypeResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan EntitlementTypeResourceModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		log.Println("[ERROR] EntitlementType: Error getting plan data")
		return
	}

	apiResp, err := r.CreateEntitlementType(ctx, &plan)
	if err != nil {
		resp.Diagnostics.AddError("Entitlement Type Creation Failed", err.Error())
		return
	}

	// Handle follow-up update for create-unsupported attributes
	updateAttempted, updateErr := r.HandleFollowUpUpdateForCreate(ctx, plan)
	if updateAttempted && updateErr != nil {
		log.Printf("[WARNING] Entitlement Type: Follow-up update failed with error: %v, verifying current state", updateErr)

		// Read current state to see what actually got applied after failed update
		readErr := r.ReadEntitlementTypeState(ctx, &plan)
		if readErr != nil {
			log.Printf("[WARNING] Entitlement Type: Could not verify current state after failed follow up update: %v", readErr)
		} else {
			log.Printf("[INFO] Entitlement Type: Current state verified after failed follwo up update - Terraform state now reflects actual configuration")
		}

		resp.Diagnostics.AddWarning(
			"Partial Configuration Applied",
			fmt.Sprintf("Entitlement type was created successfully, but some advanced configuration could not be applied: %v. The current state has been verified and Terraform state updated accordingly. Run 'terraform apply' again to retry the configuration.", updateErr),
		)
	}

	// Update model from create response
	r.UpdateModelAfterCreate(&plan, apiResp)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
	log.Printf("[DEBUG] Entitlement type: Successfully created entitlement type: %s", plan.EntitlementName.ValueString())
}

func (r *EntitlementTypeResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state EntitlementTypeResourceModel

	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Read the current state from the API
	err := r.ReadEntitlementTypeState(ctx, &state)
	if err != nil {
		resp.Diagnostics.AddError("API Read Failed", err.Error())
		return
	}

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
}

func (r *EntitlementTypeResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan EntitlementTypeResourceModel

	planGetDiagnostics := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(planGetDiagnostics...)
	if resp.Diagnostics.HasError() {
		return
	}

	_, err := r.UpdateEntitlementType(ctx, &plan)
	if err != nil {
		resp.Diagnostics.AddError("API Update Failed", err.Error())
		return
	}

	err = r.ReadEntitlementTypeState(ctx, &plan)
	if err != nil {
		resp.Diagnostics.AddError("API Read Failed After Update", err.Error())
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
	log.Printf("[DEBUG] Entitlement type: Successfully updated entitlement type: %s", plan.EntitlementName.ValueString())
}

func (r *EntitlementTypeResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// resp.State.RemoveResource(ctx)
	if os.Getenv("TF_ACC") == "1" {
		// During acceptance tests, skip deletion by just removing resource from state
		resp.State.RemoveResource(ctx)
		return
	}

	resp.Diagnostics.AddError(
		"Delete Not Supported",
		"Resource deletion is not supported by this provider. Please remove the resource manually if required, or contact your administrator.",
	)
}

func (r *EntitlementTypeResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	log.Printf("Import key received: %s", req.ID)
	idParts := strings.Split(req.ID, ":")
	if len(idParts) != 2 {
		resp.Diagnostics.AddError(
			"Invalid Import ID Format",
			fmt.Sprintf("Expected import ID format: 'endpoint_name:entitlement_name', got: %s\n"+
				"Example: terraform import saviynt_entitlement_type_resource.example sample-103:terraform_ent_type_1", req.ID),
		)
		return
	}

	endpointName := strings.TrimSpace(idParts[0])
	entitlementName := strings.TrimSpace(idParts[1])

	log.Printf("Starting import for entitltment type: %s for endpoint %s", entitlementName, endpointName)

	if endpointName == "" || entitlementName == "" {
		resp.Diagnostics.AddError(
			"Invalid Import ID Components",
			"Both endpoint_name and entitlement_name must be non-empty\n"+
				"Example: terraform import saviynt_entitlement_type_resource.example sample-103:terraform_ent_type_1",
		)
		return
	}

	// Set both endpoint_name and entitlement_name in the state
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("endpoint_name"), endpointName)...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("entitlement_name"), entitlementName)...)

	// Set the ID using the same format as used in Create/Update
	resourceID := fmt.Sprintf("ent-type-%s-%s", endpointName, entitlementName)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), resourceID)...)
}

// HandleFollowUpUpdateForCreate checks if follow-up update is needed after create and performs it
// Returns true if update was attempted, and error if the update failed
func (r *EntitlementTypeResource) HandleFollowUpUpdateForCreate(ctx context.Context, plan EntitlementTypeResourceModel) (bool, error) {
	hasUpdateOnlyAttrs := r.HasUpdateOnlyAttributes(plan)

	if !hasUpdateOnlyAttrs {
		log.Printf("[DEBUG] Entitlement Type: No follow-up update needed after create")
		return false, nil
	}

	log.Printf("[DEBUG] Entitlement type: Performing follow-up update for create-unsupported attributes")
	entitlementTypeOps := r.entitlementTypeFactory.CreateEntitlementTypeOperations(r.client.APIBaseURL(), r.token)

	err := r.PerformFollowUpUpdate(ctx, plan, entitlementTypeOps)
	if err != nil {
		log.Printf("[WARNING] Follow-up update failed: %v", err)
		return true, err
	}

	log.Printf("[DEBUG] Entitlement type: Follow-up update completed successfully")
	return true, nil
}

// HasUpdateOnlyAttributes checks if the plan contains attributes that are only supported during update
func (r *EntitlementTypeResource) HasUpdateOnlyAttributes(plan EntitlementTypeResourceModel) bool {
	log.Printf("[DEBUG] Checking hasUpdateOnlyAttributes...")

	if !plan.CreateTaskAction.IsNull() && !plan.CreateTaskAction.IsUnknown() {
		elements := plan.CreateTaskAction.Elements()
		log.Printf("[DEBUG] CreateTaskAction check - has %d elements", len(elements))
		if len(elements) > 0 {
			log.Printf("[DEBUG] CreateTaskAction needs update - returning true")
			return true
		}
	} else {
		log.Printf("[DEBUG] CreateTaskAction is null or unknown")
	}

	updateOnlyAttrs := []struct {
		name     string
		hasValue bool
	}{
		{"ArsRequestableEntitlementQuery", !plan.ArsRequestableEntitlementQuery.IsNull() && !plan.ArsRequestableEntitlementQuery.IsUnknown() && plan.ArsRequestableEntitlementQuery.ValueString() != ""},
		{"ArsSelectedEntitlementQuery", !plan.ArsSelectedEntitlementQuery.IsNull() && !plan.ArsSelectedEntitlementQuery.IsUnknown() && plan.ArsSelectedEntitlementQuery.ValueString() != ""},
		{"RequestOption", !plan.RequestOption.IsNull() && !plan.RequestOption.IsUnknown() && plan.RequestOption.ValueString() != "Request Form Table"},
		{"CreateTaskAction", !plan.CreateTaskAction.IsNull() && !plan.CreateTaskAction.IsUnknown()},
		{"StartDateInRevokeRequest", !plan.StartDateInRevokeRequest.IsNull() && !plan.StartDateInRevokeRequest.IsUnknown()},
		{"StartEndDateInRequest", !plan.StartEndDateInRequest.IsNull() && !plan.StartEndDateInRequest.IsUnknown()},
		{"AllowRemoveAllEntitlementInRequest", !plan.AllowRemoveAllEntitlementInRequest.IsNull() && !plan.AllowRemoveAllEntitlementInRequest.IsUnknown()},
	}

	for _, attr := range updateOnlyAttrs {
		if attr.hasValue {
			log.Printf("[DEBUG] Update-only attribute %s has non-default value - returning true", attr.name)
			return true
		}
	}

	log.Printf("[DEBUG] No update-only attributes found - returning false")
	return false
}

// PerformFollowUpUpdate performs an update operation immediately after create to set update-only attributes
func (r *EntitlementTypeResource) PerformFollowUpUpdate(ctx context.Context, plan EntitlementTypeResourceModel, entitlementTypeOps client.EntitlementTypeOperationsInterface) error {
	log.Printf("[DEBUG] Starting performFollowUpUpdate for entitlement type: %s for endpoint: %s", plan.EntitlementName.ValueString(), plan.EndpointName.ValueString())

	if !plan.CreateTaskAction.IsNull() && !plan.CreateTaskAction.IsUnknown() {
		elements := plan.CreateTaskAction.Elements()
		log.Printf("[DEBUG] Follow-up update: CreateTaskAction has %d elements", len(elements))
		for i, elem := range elements {
			if str, ok := elem.(types.String); ok {
				log.Printf("[DEBUG] Follow-up update: CreateTaskAction[%d]: %s", i, str.ValueString())
			}
		}
	}

	updateReqBody := openapi.UpdateEntitlementTypeRequest{
		Entitlementname:                    plan.EntitlementName.ValueString(),
		Endpointname:                       plan.EndpointName.ValueString(),
		Workflow:                           util.StringPointerOrEmpty(plan.Workflow),
		ArsRequestableEntitlementQuery:     util.StringPointerOrEmpty(plan.ArsRequestableEntitlementQuery),
		ArsSelectedEntitlementQuery:        util.StringPointerOrEmpty(plan.ArsSelectedEntitlementQuery),
		Orderindex:                         util.Int32PointerOrEmpty(plan.OrderIndex),
		EnableEntitlementToRoleSync:        util.BoolPointerOrEmtpy(plan.EnableEntitlementToRoleSync),
		Requestoption:                      util.StringPointerOrEmpty(plan.RequestOption),
		CreateTaskAction:                   util.ConvertTFSetToGoStrings(plan.CreateTaskAction),
		StartDateInRevokeRequest:           util.StringPointerOrEmpty(plan.StartDateInRevokeRequest),
		StartEndDateInRequest:              util.StringPointerOrEmpty(plan.StartEndDateInRequest),
		AllowRemoveAllEntitlementInRequest: util.StringPointerOrEmpty(plan.AllowRemoveAllEntitlementInRequest),
	}

	if reqJson, err := json.MarshalIndent(updateReqBody, "", "  "); err == nil {
		log.Printf("[DEBUG] PerformFollowUpUpdate - Full request body: %s", string(reqJson))
	} else {
		log.Printf("[DEBUG] PerformFollowUpUpdate - Failed to marshal request: %v", err)
	}

	// Log the CreateTaskAction that will be sent in the update request
	if updateReqBody.CreateTaskAction != nil {
		log.Printf("[DEBUG] Follow-up update: Sending CreateTaskAction with %d items: %v", len(updateReqBody.CreateTaskAction), updateReqBody.CreateTaskAction)
	} else {
		log.Printf("[DEBUG] Follow-up update: CreateTaskAction is nil in update request")
	}

	log.Printf("[DEBUG] Executing follow-up update API call...")
	updateResp, httpResp, err := entitlementTypeOps.UpdateEntitlementType(ctx, updateReqBody)

	if err != nil {
		log.Printf("[ERROR] Follow-up update API call failed: %v", err)
		return fmt.Errorf("follow up update API call failed: %v", err)
	}

	log.Printf("[DEBUG] Follow-up update HTTP status: %d", httpResp.StatusCode)

	if updateResp != nil && updateResp.ErrorCode != nil {
		log.Printf("[DEBUG] Follow-up update response - ErrorCode: %s, Msg: %s",
			util.SafeDeref(updateResp.ErrorCode),
			util.SafeDeref(updateResp.Msg))

		if *updateResp.ErrorCode != "0" {
			return fmt.Errorf("follow up update returned error: %s - %s", util.SafeDeref(updateResp.ErrorCode), util.SafeDeref(updateResp.Msg))
		}
	}

	log.Printf("[DEBUG] Follow-up update completed successfully")
	return nil
}

// NormalizeEntitlementTypeState ensures consistent state representation across create/update operations
func (r *EntitlementTypeResource) NormalizeEntitlementTypeState(plan *EntitlementTypeResourceModel) {
	log.Printf("[DEBUG] Starting state normalization...")

	if plan.DisplayName.IsNull() || plan.DisplayName.IsUnknown() {
		plan.DisplayName = plan.EntitlementName
	}

	plan.Workflow = util.SafeString(plan.Workflow.ValueStringPointer())
	plan.AvailableQueryServiceAccount = util.SafeString(plan.AvailableQueryServiceAccount.ValueStringPointer())
	plan.SelectedQueryServiceAccount = util.SafeString(plan.SelectedQueryServiceAccount.ValueStringPointer())
	plan.ArsRequestableEntitlementQuery = util.SafeString(plan.ArsRequestableEntitlementQuery.ValueStringPointer())
	plan.ArsSelectedEntitlementQuery = util.SafeString(plan.ArsSelectedEntitlementQuery.ValueStringPointer())
	plan.RequestDatesConfJson = util.SafeString(plan.RequestDatesConfJson.ValueStringPointer())
	plan.StartDateInRevokeRequest = util.SafeString(plan.StartDateInRevokeRequest.ValueStringPointer())
	plan.StartEndDateInRequest = util.SafeString(plan.StartEndDateInRequest.ValueStringPointer())
	plan.AllowRemoveAllEntitlementInRequest = util.SafeString(plan.AllowRemoveAllEntitlementInRequest.ValueStringPointer())

	if plan.ShowEntTypeOn.IsNull() || plan.ShowEntTypeOn.IsUnknown() {
		plan.ShowEntTypeOn = types.StringValue("0")
	}
	if plan.RequiredInRequest.IsNull() || plan.RequiredInRequest.IsUnknown() {
		plan.RequiredInRequest = types.BoolValue(false)
	}
	if plan.RequiredInServiceRequest.IsNull() || plan.RequiredInServiceRequest.IsUnknown() {
		plan.RequiredInServiceRequest = types.BoolValue(false)
	}
	if plan.Certifiable.IsNull() || plan.Certifiable.IsUnknown() {
		plan.Certifiable = types.BoolValue(true)
	}
	if plan.EnableEntitlementToRoleSync.IsNull() || plan.EnableEntitlementToRoleSync.IsUnknown() {
		plan.EnableEntitlementToRoleSync = types.BoolValue(false)
	}
	if plan.Recon.IsNull() || plan.Recon.IsUnknown() {
		plan.Recon = types.BoolValue(true)
	}
	if plan.ExcludeRuleAssgnEntsInRequest.IsNull() || plan.ExcludeRuleAssgnEntsInRequest.IsUnknown() {
		plan.ExcludeRuleAssgnEntsInRequest = types.BoolValue(true)
	}

	if plan.OrderIndex.IsNull() || plan.OrderIndex.IsUnknown() {
		plan.OrderIndex = types.Int32Value(0)
	}

	if plan.HierarchyRequired.IsNull() || plan.HierarchyRequired.IsUnknown() {
		plan.HierarchyRequired = types.StringValue("0")
	}
	if plan.RequestOption.IsNull() || plan.RequestOption.IsUnknown() {
		plan.RequestOption = types.StringValue("Request Form Table")
	}

	if !plan.CreateTaskAction.IsNull() && !plan.CreateTaskAction.IsUnknown() {
		elements := plan.CreateTaskAction.Elements()
		log.Printf("[DEBUG] After normalization: CreateTaskAction has %d elements", len(elements))
	} else {
		log.Printf("[DEBUG] After normalization: CreateTaskAction is null or unknown")
	}

	log.Printf("[DEBUG] State normalization completed")
}
