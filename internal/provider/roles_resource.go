// Copyright (c) Saviynt Inc.
// SPDX-License-Identifier: MPL-2.0

// saviynt_roles_resource manages roles in the Saviynt Security Manager.
// The resource implements the full Terraform lifecycle:
//   - Create: provisions a new role using the supplied configuration.
//   - Read: fetches the current role state from Saviynt to keep Terraform's state in sync.
//   - Update: applies any configuration changes to an existing role.
//   - Import: brings an existing role under Terraform management by its name.
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
	endpointsutil "terraform-provider-Saviynt/util/endpointsutil"
	"terraform-provider-Saviynt/util/rolesutil"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	s "github.com/saviynt/saviynt-api-go-client"
	endpoint "github.com/saviynt/saviynt-api-go-client/endpoints"
	openapi "github.com/saviynt/saviynt-api-go-client/roles"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &rolesResource{}
var _ resource.ResourceWithImportState = &rolesResource{}

// RolesResourceModel defines the state for our roles resource.
type RolesResourceModel struct {
	ID               types.String `tfsdk:"id"`
	RoleType         types.String `tfsdk:"role_type"`
	RoleName         types.String `tfsdk:"role_name"`
	Requestor        types.String `tfsdk:"requestor"`
	Owners           types.Set    `tfsdk:"owners"`
	CustomProperty1  types.String `tfsdk:"custom_property1"`
	CustomProperty2  types.String `tfsdk:"custom_property2"`
	CustomProperty3  types.String `tfsdk:"custom_property3"`
	CustomProperty4  types.String `tfsdk:"custom_property4"`
	CustomProperty5  types.String `tfsdk:"custom_property5"`
	CustomProperty6  types.String `tfsdk:"custom_property6"`
	CustomProperty7  types.String `tfsdk:"custom_property7"`
	CustomProperty8  types.String `tfsdk:"custom_property8"`
	CustomProperty9  types.String `tfsdk:"custom_property9"`
	CustomProperty10 types.String `tfsdk:"custom_property10"`
	CustomProperty11 types.String `tfsdk:"custom_property11"`
	CustomProperty12 types.String `tfsdk:"custom_property12"`
	CustomProperty13 types.String `tfsdk:"custom_property13"`
	CustomProperty14 types.String `tfsdk:"custom_property14"`
	CustomProperty15 types.String `tfsdk:"custom_property15"`
	CustomProperty16 types.String `tfsdk:"custom_property16"`
	CustomProperty17 types.String `tfsdk:"custom_property17"`
	CustomProperty18 types.String `tfsdk:"custom_property18"`
	CustomProperty19 types.String `tfsdk:"custom_property19"`
	CustomProperty20 types.String `tfsdk:"custom_property20"`
	CustomProperty21 types.String `tfsdk:"custom_property21"`
	CustomProperty22 types.String `tfsdk:"custom_property22"`
	CustomProperty23 types.String `tfsdk:"custom_property23"`
	CustomProperty24 types.String `tfsdk:"custom_property24"`
	CustomProperty25 types.String `tfsdk:"custom_property25"`
	CustomProperty26 types.String `tfsdk:"custom_property26"`
	CustomProperty27 types.String `tfsdk:"custom_property27"`
	CustomProperty28 types.String `tfsdk:"custom_property28"`
	CustomProperty29 types.String `tfsdk:"custom_property29"`
	CustomProperty30 types.String `tfsdk:"custom_property30"`
	CustomProperty31 types.String `tfsdk:"custom_property31"`
	CustomProperty32 types.String `tfsdk:"custom_property32"`
	CustomProperty33 types.String `tfsdk:"custom_property33"`
	CustomProperty34 types.String `tfsdk:"custom_property34"`
	CustomProperty35 types.String `tfsdk:"custom_property35"`
	CustomProperty36 types.String `tfsdk:"custom_property36"`
	CustomProperty37 types.String `tfsdk:"custom_property37"`
	CustomProperty38 types.String `tfsdk:"custom_property38"`
	CustomProperty39 types.String `tfsdk:"custom_property39"`
	CustomProperty40 types.String `tfsdk:"custom_property40"`
	CustomProperty41 types.String `tfsdk:"custom_property41"`
	CustomProperty42 types.String `tfsdk:"custom_property42"`
	CustomProperty43 types.String `tfsdk:"custom_property43"`
	CustomProperty44 types.String `tfsdk:"custom_property44"`
	CustomProperty45 types.String `tfsdk:"custom_property45"`
	CustomProperty46 types.String `tfsdk:"custom_property46"`
	CustomProperty47 types.String `tfsdk:"custom_property47"`
	CustomProperty48 types.String `tfsdk:"custom_property48"`
	CustomProperty49 types.String `tfsdk:"custom_property49"`
	CustomProperty50 types.String `tfsdk:"custom_property50"`
	CustomProperty51 types.String `tfsdk:"custom_property51"`
	CustomProperty52 types.String `tfsdk:"custom_property52"`
	CustomProperty53 types.String `tfsdk:"custom_property53"`
	CustomProperty54 types.String `tfsdk:"custom_property54"`
	CustomProperty55 types.String `tfsdk:"custom_property55"`
	CustomProperty56 types.String `tfsdk:"custom_property56"`
	CustomProperty57 types.String `tfsdk:"custom_property57"`
	CustomProperty58 types.String `tfsdk:"custom_property58"`
	CustomProperty59 types.String `tfsdk:"custom_property59"`
	CustomProperty60 types.String `tfsdk:"custom_property60"`
	EndpointName     types.String `tfsdk:"endpoint_name"`
	DefaultTimeFrame types.String `tfsdk:"default_time_frame"`
	Description      types.String `tfsdk:"description"`
	DisplayName      types.String `tfsdk:"display_name"`
	Glossary         types.String `tfsdk:"glossary"`
	Risk             types.String `tfsdk:"risk"`
	Level            types.String `tfsdk:"level"`
	SoxCritical      types.String `tfsdk:"sox_critical"`
	SysCritical      types.String `tfsdk:"sys_critical"`
	Priviliged       types.String `tfsdk:"priviliged"`
	Confidentiality  types.String `tfsdk:"confidentiality"`
	Requestable      types.String `tfsdk:"requestable"`
	ShowDynamicAttrs types.String `tfsdk:"show_dynamic_attrs"`
	CheckSod         types.String `tfsdk:"check_sod"`
	Entitlements     types.Set    `tfsdk:"entitlements"`
	RoleResponse
}

// Entitlement defines the structure for entitlements associated with a role.
type Entitlement struct {
	EntitlementValue types.String `tfsdk:"entitlement_value"`
	EntitlementType  types.String `tfsdk:"entitlement_type"`
	Endpoint         types.String `tfsdk:"endpoint"`
}

// RoleResponse defines the structure for the response from Saviynt API when creating or updating a role.
type RoleResponse struct {
	RequestId  types.String `tfsdk:"request_id"`
	Requestkey types.String `tfsdk:"request_key"`
	Msg        types.String `tfsdk:"msg"`
	ErrorCode  types.String `tfsdk:"error_code"`
}

// Owner defines the structure for owners associated with a role.
type Owner struct {
	OwnerName types.String `tfsdk:"owner_name"`
	Rank      types.String `tfsdk:"rank"`
}

// rolesResource implements the resource.Resource interface for managing Saviynt roles.
type rolesResource struct {
	client *s.Client
	token  string
}

// NewRolesResource creates a new instance of the roles resource.
func NewRolesResource() resource.Resource {
	return &rolesResource{}
}

// Metadata returns the metadata for the roles resource.
func (r *rolesResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = "saviynt_roles_resource"
}

// Schema defines the schema for the roles resource.
func (r *rolesResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Create and manage Roles in Saviynt",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The unique ID of the resource.",
			},
			"role_type": schema.StringAttribute{
				Description: "Type of the role",
				Required:    true,
			},
			"role_name": schema.StringAttribute{
				Description: "Name of the role",
				Required:    true,
			},
			"requestor": schema.StringAttribute{
				Description: "Requester of the role",
				Required:    true,
			},
			"owners": schema.SetAttribute{
				Description: "Set of owners",
				Optional:    true,
				Computed:    true,
				ElementType: types.ObjectType{
					AttrTypes: map[string]attr.Type{
						"owner_name": types.StringType,
						"rank":       types.StringType,
					},
				},
			},
			"endpoint_name": schema.StringAttribute{
				Description: "Name of the endpoint associated with the role",
				Optional:    true,
				Computed:    true,
			},
			"default_time_frame": schema.StringAttribute{
				Description: "Default time frame for the role",
				Optional:    true,
				Computed:    true,
			},
			"description": schema.StringAttribute{
				Description: "Description of the role",
				Optional:    true,
				Computed:    true,
			},
			"display_name": schema.StringAttribute{
				Description: "Display name of the role",
				Optional:    true,
				Computed:    true,
			},
			"glossary": schema.StringAttribute{
				Description: "Glossary term associated with the role",
				Optional:    true,
				Computed:    true,
			},
			"risk": schema.StringAttribute{
				Description: "Risk level of the role",
				Optional:    true,
			},
			"level": schema.StringAttribute{
				Description: "Risk level of the role",
				Optional:    true,
			},
			"sox_critical": schema.StringAttribute{
				Description: "SOX criticality of the role",
				Optional:    true,
				Computed:    true,
			},
			"sys_critical": schema.StringAttribute{
				Description: "System criticality of the role",
				Optional:    true,
				Computed:    true,
			},
			"priviliged": schema.StringAttribute{
				Description: "Indicates if the role is privileged",
				Optional:    true,
				Computed:    true,
			},
			"confidentiality": schema.StringAttribute{
				Description: "Confidentiality level of the role",
				Optional:    true,
				Computed:    true,
			},
			"requestable": schema.StringAttribute{
				Description: "Indicates if the role is requestable",
				Optional:    true,
				Computed:    true,
			},
			"show_dynamic_attrs": schema.StringAttribute{
				Description: "Indicates if dynamic attributes should be shown",
				Optional:    true,
				Computed:    true,
			},
			"check_sod": schema.StringAttribute{
				Description: "Indicates if segregation of duties (SoD) checks should be performed",
				Optional:    true,
			},
			"entitlements": schema.SetAttribute{
				Description: "Set of entitlements associated with the role",
				Optional:    true,
				Computed:    true,
				ElementType: types.ObjectType{
					AttrTypes: map[string]attr.Type{
						"entitlement_value": types.StringType,
						"entitlement_type":  types.StringType,
						"endpoint":          types.StringType,
					},
				},
			},
			"request_id": schema.StringAttribute{
				Description: "Request ID for the role operation",
				Optional:    true,
				Computed:    true,
			},
			"request_key": schema.StringAttribute{
				Description: "Request key for the role operation",
				Optional:    true,
				Computed:    true,
			},
			"msg": schema.StringAttribute{
				Description: "Message returned from the API",
				Optional:    true,
				Computed:    true,
			},
			"error_code": schema.StringAttribute{
				Description: "Error code returned from the API",
				Optional:    true,
				Computed:    true,
			},
		},
	}
	// Add custom properties to the schema.
	for i := 1; i <= 60; i++ {
		key := fmt.Sprintf("custom_property%d", i)
		resp.Schema.Attributes[key] = schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: fmt.Sprintf("Custom Property %d.", i),
		}
	}
}

// Configure initializes the roles resource with the provider's API client and access token.
func (r *rolesResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

// Create implements the resource.Resource interface for creating a new role in Saviynt.
func (r *rolesResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan RolesResourceModel
	// Extract plan from request
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Check if the client and token are set.
	cfg := openapi.NewConfiguration()
	apiBaseURL := strings.TrimPrefix(strings.TrimPrefix(r.client.APIBaseURL(), "https://"), "http://")
	cfg.Host = apiBaseURL
	cfg.Scheme = "https"
	cfg.AddDefaultHeader("Authorization", "Bearer "+r.token)
	cfg.HTTPClient = http.DefaultClient
	apiClient := openapi.NewAPIClient(cfg)

	// Setting up the owners for the role.
	var owners []openapi.CreateRoleOwnerPayload
	var ownerdiags diag.Diagnostics
	var tfOwnerTemplates []Owner
	ownerdiags = plan.Owners.ElementsAs(ctx, &tfOwnerTemplates, true)
	resp.Diagnostics.Append(ownerdiags...)
	if resp.Diagnostics.HasError() {
		return
	}
	if plan.Owners.IsNull() || plan.Owners.IsUnknown() {
		log.Println("No owners provided, skipping owner creation.")
		resp.Diagnostics.AddError(
			"Invalid Owners",
			"Owners cannot be null or unknown. Please provide a valid list of owners.",
		)
		return
	} else {
		if len(plan.Owners.Elements()) == 0 {
			log.Println("No owners provided, skipping owner creation.")
			resp.Diagnostics.AddError(
				"Invalid Owners",
				"Owners cannot be an empty list. Please provide at least one owner.",
			)
			return
		} else {
			for _, tfOwnerTemplate := range tfOwnerTemplates {
				owner := openapi.CreateRoleOwnerPayload{}
				if tfOwnerTemplate.OwnerName.IsNull() || tfOwnerTemplate.OwnerName.IsUnknown() {
					log.Println("Owner name is null or unknown, skipping this owner.")
					resp.Diagnostics.AddError(
						"Invalid Owner Name",
						"Owner name cannot be null or unknown. Please provide a valid owner name.",
					)
					return
				} else {
					owner.OwnerName = tfOwnerTemplate.OwnerName.ValueStringPointer()
				}
				if tfOwnerTemplate.Rank.IsNull() || tfOwnerTemplate.Rank.IsUnknown() {
					log.Println("Owner rank is null or unknown, skipping this owner.")
					resp.Diagnostics.AddError(
						"Invalid Owner Rank",
						"Owner rank cannot be null or unknown. Please provide a valid owner rank.",
					)
					return
				} else {
					owner.Rank = tfOwnerTemplate.Rank.ValueStringPointer()
				}
				owners = append(owners, owner)
			}
		}
	}

	// Setting the attributes for the role.
	createReq := openapi.CreateEnterpriseRoleRequest{
		// Required fields
		RoleName:  plan.RoleName.ValueString(),
		Roletype:  plan.RoleType.ValueString(),
		Requestor: plan.Requestor.ValueString(),
		Owner:     owners,
		// Optional fields
		Customproperty1:  util.StringPointerOrEmpty(plan.CustomProperty1),
		Customproperty2:  util.StringPointerOrEmpty(plan.CustomProperty2),
		Customproperty3:  util.StringPointerOrEmpty(plan.CustomProperty3),
		Customproperty4:  util.StringPointerOrEmpty(plan.CustomProperty4),
		Customproperty5:  util.StringPointerOrEmpty(plan.CustomProperty5),
		Customproperty6:  util.StringPointerOrEmpty(plan.CustomProperty6),
		Customproperty7:  util.StringPointerOrEmpty(plan.CustomProperty7),
		Customproperty8:  util.StringPointerOrEmpty(plan.CustomProperty8),
		Customproperty9:  util.StringPointerOrEmpty(plan.CustomProperty9),
		Customproperty10: util.StringPointerOrEmpty(plan.CustomProperty10),
		Customproperty11: util.StringPointerOrEmpty(plan.CustomProperty11),
		Customproperty12: util.StringPointerOrEmpty(plan.CustomProperty12),
		Customproperty13: util.StringPointerOrEmpty(plan.CustomProperty13),
		Customproperty14: util.StringPointerOrEmpty(plan.CustomProperty14),
		Customproperty15: util.StringPointerOrEmpty(plan.CustomProperty15),
		Customproperty16: util.StringPointerOrEmpty(plan.CustomProperty16),
		Customproperty17: util.StringPointerOrEmpty(plan.CustomProperty17),
		Customproperty18: util.StringPointerOrEmpty(plan.CustomProperty18),
		Customproperty19: util.StringPointerOrEmpty(plan.CustomProperty19),
		Customproperty20: util.StringPointerOrEmpty(plan.CustomProperty20),
		Customproperty21: util.StringPointerOrEmpty(plan.CustomProperty21),
		Customproperty22: util.StringPointerOrEmpty(plan.CustomProperty22),
		Customproperty23: util.StringPointerOrEmpty(plan.CustomProperty23),
		Customproperty24: util.StringPointerOrEmpty(plan.CustomProperty24),
		Customproperty25: util.StringPointerOrEmpty(plan.CustomProperty25),
		Customproperty26: util.StringPointerOrEmpty(plan.CustomProperty26),
		Customproperty27: util.StringPointerOrEmpty(plan.CustomProperty27),
		Customproperty28: util.StringPointerOrEmpty(plan.CustomProperty28),
		Customproperty29: util.StringPointerOrEmpty(plan.CustomProperty29),
		Customproperty30: util.StringPointerOrEmpty(plan.CustomProperty30),
		Customproperty31: util.StringPointerOrEmpty(plan.CustomProperty31),
		Customproperty32: util.StringPointerOrEmpty(plan.CustomProperty32),
		Customproperty33: util.StringPointerOrEmpty(plan.CustomProperty33),
		Customproperty34: util.StringPointerOrEmpty(plan.CustomProperty34),
		Customproperty35: util.StringPointerOrEmpty(plan.CustomProperty35),
		Customproperty36: util.StringPointerOrEmpty(plan.CustomProperty36),
		Customproperty37: util.StringPointerOrEmpty(plan.CustomProperty37),
		Customproperty38: util.StringPointerOrEmpty(plan.CustomProperty38),
		Customproperty39: util.StringPointerOrEmpty(plan.CustomProperty39),
		Customproperty40: util.StringPointerOrEmpty(plan.CustomProperty40),
		Customproperty41: util.StringPointerOrEmpty(plan.CustomProperty41),
		Customproperty42: util.StringPointerOrEmpty(plan.CustomProperty42),
		Customproperty43: util.StringPointerOrEmpty(plan.CustomProperty43),
		Customproperty44: util.StringPointerOrEmpty(plan.CustomProperty44),
		Customproperty45: util.StringPointerOrEmpty(plan.CustomProperty45),
		Customproperty46: util.StringPointerOrEmpty(plan.CustomProperty46),
		Customproperty47: util.StringPointerOrEmpty(plan.CustomProperty47),
		Customproperty48: util.StringPointerOrEmpty(plan.CustomProperty48),
		Customproperty49: util.StringPointerOrEmpty(plan.CustomProperty49),
		Customproperty50: util.StringPointerOrEmpty(plan.CustomProperty50),
		Customproperty51: util.StringPointerOrEmpty(plan.CustomProperty51),
		Customproperty52: util.StringPointerOrEmpty(plan.CustomProperty52),
		Customproperty53: util.StringPointerOrEmpty(plan.CustomProperty53),
		Customproperty54: util.StringPointerOrEmpty(plan.CustomProperty54),
		Customproperty55: util.StringPointerOrEmpty(plan.CustomProperty55),
		Customproperty56: util.StringPointerOrEmpty(plan.CustomProperty56),
		Customproperty57: util.StringPointerOrEmpty(plan.CustomProperty57),
		Customproperty58: util.StringPointerOrEmpty(plan.CustomProperty58),
		Customproperty59: util.StringPointerOrEmpty(plan.CustomProperty59),
		Customproperty60: util.StringPointerOrEmpty(plan.CustomProperty60),
		Endpointname:     util.StringPointerOrEmpty(plan.EndpointName),
		Defaulttimeframe: util.StringPointerOrEmpty(plan.DefaultTimeFrame),
		Description:      util.StringPointerOrEmpty(plan.Description),
		Displayname:      util.StringPointerOrEmpty(plan.DisplayName),
		Glossary:         util.StringPointerOrEmpty(plan.Glossary),
		Risk:             util.StringPointerOrEmpty(plan.Risk),
		Level:            util.StringPointerOrEmpty(plan.Level),
		Soxcritical:      util.StringPointerOrEmpty(plan.SoxCritical),
		Syscritical:      util.StringPointerOrEmpty(plan.SysCritical),
		Priviliged:       util.StringPointerOrEmpty(plan.Priviliged),
		Confidentiality:  util.StringPointerOrEmpty(plan.Confidentiality),
		Requestable:      util.StringPointerOrEmpty(plan.Requestable),
		ShowDynamicAttrs: util.StringPointerOrEmpty(plan.ShowDynamicAttrs),
		Checksod:         util.StringPointerOrEmpty(plan.CheckSod),
	}

	// Execute the API call.
	apiResp, https, err := apiClient.RolesAPI.CreateEnterpriseRoleRequest(ctx).CreateEnterpriseRoleRequest(createReq).Execute()

	// For the error other than the 200 OK response
	if https != nil && https.StatusCode != http.StatusOK {
		log.Printf("[ERROR] HTTP error while creating role: %s", https.Status)
		var fetchResp map[string]interface{}
		if err := json.NewDecoder(https.Body).Decode(&fetchResp); err != nil {
			resp.Diagnostics.AddError("Failed to decode error response", err.Error())
			return
		}
		resp.Diagnostics.AddError(
			"HTTP Error",
			fmt.Sprintf("HTTP error while creating role for the reasons: %s", fetchResp["message"]),
		)
		return
	}

	// If some other error
	if err != nil || *apiResp.ErrorCode != "0" {
		log.Printf("[ERROR] Failed to create API resource. Error: %v", err)
		resp.Diagnostics.AddError("API Create Failed", fmt.Sprintf("Error: %v", err))
		return
	}

	// Set the ID from the role name
	plan.ID = types.StringValue("roles-" + plan.RoleName.ValueString())

	// Populate the response model with the API response data.
	plan.RequestId = types.StringValue(*apiResp.Requestid)
	plan.Requestkey = types.StringValue(*apiResp.Requestkey)
	plan.Msg = types.StringValue(*apiResp.Message)
	plan.ErrorCode = types.StringValue(*apiResp.ErrorCode)

	// Set the plan with the request data.
	plan.RoleName = types.StringValue(createReq.RoleName)
	plan.RoleType = types.StringValue(createReq.Roletype)
	plan.Requestor = types.StringValue(createReq.Requestor)
	plan.EndpointName = types.StringValue(util.SafeDeref(createReq.Endpointname))
	plan.DefaultTimeFrame = types.StringValue(util.SafeDeref(createReq.Defaulttimeframe))
	plan.Description = util.SafeString(createReq.Description)
	plan.DisplayName = types.StringValue(util.SafeDeref(createReq.Displayname))
	plan.Glossary = util.SafeString(createReq.Glossary)
	plan.SoxCritical = util.SafeString(createReq.Soxcritical)
	plan.SysCritical = util.SafeString(createReq.Syscritical)
	plan.Priviliged = util.SafeString(createReq.Priviliged)
	plan.Confidentiality = util.SafeString(createReq.Confidentiality)
	plan.Requestable = util.SafeString(createReq.Requestable)
	plan.ShowDynamicAttrs = util.SafeString(createReq.ShowDynamicAttrs)

	// Set the owners to empty set if it is null or unknown
	if plan.Owners.IsNull() || plan.Owners.IsUnknown() {
		plan.Owners = types.SetValueMust(types.ObjectType{
			AttrTypes: map[string]attr.Type{
				"owner_name": types.StringType,
				"rank":       types.StringType,
			},
		}, []attr.Value{})

	}

	// Set the entitlements to empty set if it is null or unknown
	if plan.Entitlements.IsNull() || plan.Entitlements.IsUnknown() {
		plan.Entitlements = types.SetValueMust(types.ObjectType{
			AttrTypes: map[string]attr.Type{
				"entitlement_value": types.StringType,
				"entitlement_type":  types.StringType,
				"endpoint":          types.StringType,
			},
		}, []attr.Value{})

	}

	// Set the custom properties
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
	plan.CustomProperty46 = util.SafeString(plan.CustomProperty46.ValueStringPointer())
	plan.CustomProperty47 = util.SafeString(plan.CustomProperty47.ValueStringPointer())
	plan.CustomProperty48 = util.SafeString(plan.CustomProperty48.ValueStringPointer())
	plan.CustomProperty49 = util.SafeString(plan.CustomProperty49.ValueStringPointer())
	plan.CustomProperty50 = util.SafeString(plan.CustomProperty50.ValueStringPointer())
	plan.CustomProperty51 = util.SafeString(plan.CustomProperty51.ValueStringPointer())
	plan.CustomProperty52 = util.SafeString(plan.CustomProperty52.ValueStringPointer())
	plan.CustomProperty53 = util.SafeString(plan.CustomProperty53.ValueStringPointer())
	plan.CustomProperty54 = util.SafeString(plan.CustomProperty54.ValueStringPointer())
	plan.CustomProperty55 = util.SafeString(plan.CustomProperty55.ValueStringPointer())
	plan.CustomProperty56 = util.SafeString(plan.CustomProperty56.ValueStringPointer())
	plan.CustomProperty57 = util.SafeString(plan.CustomProperty57.ValueStringPointer())
	plan.CustomProperty58 = util.SafeString(plan.CustomProperty58.ValueStringPointer())
	plan.CustomProperty59 = util.SafeString(plan.CustomProperty59.ValueStringPointer())
	plan.CustomProperty60 = util.SafeString(plan.CustomProperty60.ValueStringPointer())
	stateCreateDiagnostics := resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(stateCreateDiagnostics...)
}

// Read implements the resource.Resource interface for reading a role from Saviynt.
func (r *rolesResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state RolesResourceModel

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
	reqParams := openapi.GetRolesRequest{}
	reqParams.SetRoleName(state.RoleName.ValueString())
	reqParams.SetRequestedObject("entitlements")
	apiResp, httpResp, err := apiClient.RolesAPI.GetRoles(ctx).GetRolesRequest(reqParams).Execute()
	//for the endpoint name
	if apiResp.Roledetails[0].Endpointkey != nil {
		endpointcfg := endpoint.NewConfiguration()
		endpointcfg.Host = apiBaseURL
		endpointcfg.Scheme = "https"
		endpointcfg.AddDefaultHeader("Authorization", "Bearer "+r.token)
		endpointcfg.HTTPClient = http.DefaultClient

		endpointClient := endpoint.NewAPIClient(endpointcfg)
		endpointReq := endpoint.GetEndpointsRequest{}
		var endpointKeys []string
		endpointKeys = []string{*apiResp.Roledetails[0].Endpointkey}
		endpointReq.Endpointkey = endpointKeys
		endpointResp, httpRespEndpoint, err := endpointClient.EndpointsAPI.GetEndpoints(ctx).GetEndpointsRequest(endpointReq).Execute()
		if err != nil {
			log.Printf("[ERROR] Failed to fetch endpoint details: %v", err)
			resp.Diagnostics.AddError("API Read Failed", fmt.Sprintf("Error fetching endpoint details: %v", err))
			return
		}
		if httpRespEndpoint != nil && httpRespEndpoint.StatusCode != http.StatusOK {
			log.Printf("[ERROR] HTTP error while reading endpoint: %s", httpRespEndpoint.Status)
			var fetchResp map[string]interface{}
			if err := json.NewDecoder(httpRespEndpoint.Body).Decode(&fetchResp); err != nil {
				resp.Diagnostics.AddError("Failed to decode error response", err.Error())
				return
			}
			resp.Diagnostics.AddError(
				"HTTP Error",
				fmt.Sprintf("HTTP error while reading endpoint: %s", fetchResp["message"]),
			)
			return
		}
		if len(endpointResp.Endpoints) > 0 {
			state.EndpointName = types.StringValue(*endpointResp.Endpoints[0].Endpointname)
		} else {
			log.Println("No endpoint found for the role")
			state.EndpointName = types.StringNull()
		}
	}
	// For the error other than the 200 OK response
	if httpResp != nil && httpResp.StatusCode != http.StatusOK {
		log.Printf("[ERROR] HTTP error while reading role: %s", httpResp.Status)
		var fetchResp map[string]interface{}
		if err := json.NewDecoder(httpResp.Body).Decode(&fetchResp); err != nil {
			resp.Diagnostics.AddError("Failed to decode error response", err.Error())
			return
		}
		resp.Diagnostics.AddError(
			"HTTP Error",
			fmt.Sprintf("HTTP error while reading role: %s", fetchResp["message"]),
		)
		return
	}

	if err != nil {
		log.Printf("Problem with the get function in read block: %v", err)
		resp.Diagnostics.AddError("API Read Failed", fmt.Sprintf("Error: %v", err))
		return
	}

	// Check if we got valid role details
	if apiResp.Roledetails == nil || len(apiResp.Roledetails) == 0 {
		resp.Diagnostics.AddError(
			"Client Error",
			"API returned empty role details list",
		)
		return
	}

	// Set the ID from the role name
	state.ID = types.StringValue("roles-" + state.RoleName.ValueString())
	roleDetails := apiResp.Roledetails[0]

	// Set basic role properties
	state.Msg = util.SafeString(apiResp.Msg)
	state.ErrorCode = util.SafeString(apiResp.ErrorCode)
	if roleDetails.Roletype != nil {
		state.RoleType = types.StringValue(endpointsutil.TranslateValue(*roleDetails.Roletype, rolesutil.RoleTypeMap))
	}
	state.RoleName = util.SafeString(roleDetails.RoleName)
	state.Description = util.SafeString(roleDetails.Description)
	state.DisplayName = util.SafeString(roleDetails.Displayname)
	state.DefaultTimeFrame = util.SafeString(roleDetails.DefaultTimeFrameHrs)
	state.Glossary = util.SafeString(roleDetails.Glossary)

	// Mapping of soxcritical, syscritical, and other role attributes
	if roleDetails.Soxcritical != nil {
		state.SoxCritical = types.StringValue(endpointsutil.TranslateValue(*roleDetails.Soxcritical, rolesutil.SoxCriticalityMap))
	}
	if roleDetails.Syscritical != nil {
		state.SysCritical = types.StringValue(endpointsutil.TranslateValue(*roleDetails.Syscritical, rolesutil.SysCriticalMap))
	}
	if roleDetails.Priviliged != nil {
		state.Priviliged = types.StringValue(endpointsutil.TranslateValue(*roleDetails.Priviliged, rolesutil.PrivilegedMap))
	}
	if roleDetails.Confidentiality != nil {
		state.Confidentiality = types.StringValue(endpointsutil.TranslateValue(*roleDetails.Confidentiality, rolesutil.ConfidentialityMap))
	}
	state.Requestable = util.SafeString(roleDetails.Requestable)
	state.ShowDynamicAttrs = util.SafeString(roleDetails.ShowDynamicAttrs)

	// Process owners if available in the response
	var ownersArr []openapi.GetRoleOwnersResponse
	ownerUnion := roleDetails.Owner
	if ownerUnion != nil {
		if ownerUnion.ArrayOfGetRoleOwnersResponse != nil {
			ownersArr = *ownerUnion.ArrayOfGetRoleOwnersResponse
		}
		if ownerUnion.String != nil {
			var tmp []openapi.GetRoleOwnersResponse
			ownersArr = tmp
		}
	}
	ownerObjects := make([]attr.Value, 0, len(ownersArr))
	for _, o := range ownersArr {
		ownerVal, ownerDiags := types.ObjectValue(
			map[string]attr.Type{
				"owner_name": types.StringType,
				"rank":       types.StringType,
			},
			map[string]attr.Value{
				"owner_name": types.StringValue(*o.Ownername),
				"rank":       types.StringValue(util.SafeDeref(o.Rank)),
			},
		)
		if ownerDiags.HasError() {
			resp.Diagnostics.Append(ownerDiags...)
			continue
		}
		ownerObjects = append(ownerObjects, ownerVal)
	}
	setOwnerVal, setOwnerDiags := types.SetValue(
		types.ObjectType{
			AttrTypes: map[string]attr.Type{
				"owner_name": types.StringType,
				"rank":       types.StringType,
			},
		}, ownerObjects,
	)
	if setOwnerDiags.HasError() {
		resp.Diagnostics.Append(setOwnerDiags...)
		return
	}
	state.Owners = setOwnerVal

	// Process entitlement if available in the response
	entitlementsArr := make([]openapi.GetEntitlementDetailsResponse, 0)
	entitlementUnion := roleDetails.EntitlementDetails
	if entitlementUnion != nil {
		entitlementsArr = entitlementUnion
	}
	entitlementObjects := make([]attr.Value, 0, len(entitlementsArr))
	for _, o := range entitlementsArr {
		entitlementVal, entitlementDiags := types.ObjectValue(
			map[string]attr.Type{
				"entitlement_value": types.StringType,
				"entitlement_type":  types.StringType,
				"endpoint":          types.StringType,
			},
			map[string]attr.Value{
				"entitlement_value": types.StringValue(*o.EntitlementValue),
				"entitlement_type":  types.StringValue(util.SafeDeref(o.EntitlementTypeName)),
				"endpoint":          types.StringValue(util.SafeDeref(o.Endpoint)),
			},
		)
		if entitlementDiags.HasError() {
			resp.Diagnostics.Append(entitlementDiags...)
			continue
		}
		entitlementObjects = append(entitlementObjects, entitlementVal)
	}
	setEntitlementVal, setEntitlementDiags := types.SetValue(
		types.ObjectType{
			AttrTypes: map[string]attr.Type{
				"entitlement_value": types.StringType,
				"entitlement_type":  types.StringType,
				"endpoint":          types.StringType,
			},
		}, entitlementObjects,
	)
	if setEntitlementDiags.HasError() {
		resp.Diagnostics.Append(setEntitlementDiags...)
		return
	}
	state.Entitlements = setEntitlementVal

	// Setting the Custom Properties
	state.CustomProperty1 = util.SafeString(roleDetails.CustomProperty1)
	state.CustomProperty2 = util.SafeString(roleDetails.CustomProperty2)
	state.CustomProperty3 = util.SafeString(roleDetails.CustomProperty3)
	state.CustomProperty4 = util.SafeString(roleDetails.CustomProperty4)
	state.CustomProperty5 = util.SafeString(roleDetails.CustomProperty5)
	state.CustomProperty6 = util.SafeString(roleDetails.CustomProperty6)
	state.CustomProperty7 = util.SafeString(roleDetails.CustomProperty7)
	state.CustomProperty8 = util.SafeString(roleDetails.CustomProperty8)
	state.CustomProperty9 = util.SafeString(roleDetails.CustomProperty9)
	state.CustomProperty10 = util.SafeString(roleDetails.CustomProperty10)
	state.CustomProperty11 = util.SafeString(roleDetails.CustomProperty11)
	state.CustomProperty12 = util.SafeString(roleDetails.CustomProperty12)
	state.CustomProperty13 = util.SafeString(roleDetails.CustomProperty13)
	state.CustomProperty14 = util.SafeString(roleDetails.CustomProperty14)
	state.CustomProperty15 = util.SafeString(roleDetails.CustomProperty15)
	state.CustomProperty16 = util.SafeString(roleDetails.CustomProperty16)
	state.CustomProperty17 = util.SafeString(roleDetails.CustomProperty17)
	state.CustomProperty18 = util.SafeString(roleDetails.CustomProperty18)
	state.CustomProperty19 = util.SafeString(roleDetails.CustomProperty19)
	state.CustomProperty20 = util.SafeString(roleDetails.CustomProperty20)
	state.CustomProperty21 = util.SafeString(roleDetails.CustomProperty21)
	state.CustomProperty22 = util.SafeString(roleDetails.CustomProperty22)
	state.CustomProperty23 = util.SafeString(roleDetails.CustomProperty23)
	state.CustomProperty24 = util.SafeString(roleDetails.CustomProperty24)
	state.CustomProperty25 = util.SafeString(roleDetails.CustomProperty25)
	state.CustomProperty26 = util.SafeString(roleDetails.CustomProperty26)
	state.CustomProperty27 = util.SafeString(roleDetails.CustomProperty27)
	state.CustomProperty28 = util.SafeString(roleDetails.CustomProperty28)
	state.CustomProperty29 = util.SafeString(roleDetails.CustomProperty29)
	state.CustomProperty30 = util.SafeString(roleDetails.CustomProperty30)
	state.CustomProperty31 = util.SafeString(roleDetails.CustomProperty31)
	state.CustomProperty32 = util.SafeString(roleDetails.CustomProperty32)
	state.CustomProperty33 = util.SafeString(roleDetails.CustomProperty33)
	state.CustomProperty34 = util.SafeString(roleDetails.CustomProperty34)
	state.CustomProperty35 = util.SafeString(roleDetails.CustomProperty35)
	state.CustomProperty36 = util.SafeString(roleDetails.CustomProperty36)
	state.CustomProperty37 = util.SafeString(roleDetails.CustomProperty37)
	state.CustomProperty38 = util.SafeString(roleDetails.CustomProperty38)
	state.CustomProperty39 = util.SafeString(roleDetails.CustomProperty39)
	state.CustomProperty40 = util.SafeString(roleDetails.CustomProperty40)
	state.CustomProperty41 = util.SafeString(roleDetails.CustomProperty41)
	state.CustomProperty42 = util.SafeString(roleDetails.CustomProperty42)
	state.CustomProperty43 = util.SafeString(roleDetails.CustomProperty43)
	state.CustomProperty44 = util.SafeString(roleDetails.CustomProperty44)
	state.CustomProperty45 = util.SafeString(roleDetails.CustomProperty45)
	state.CustomProperty46 = util.SafeString(roleDetails.CustomProperty46)
	state.CustomProperty47 = util.SafeString(roleDetails.CustomProperty47)
	state.CustomProperty48 = util.SafeString(roleDetails.CustomProperty48)
	state.CustomProperty49 = util.SafeString(roleDetails.CustomProperty49)
	state.CustomProperty50 = util.SafeString(roleDetails.CustomProperty50)
	state.CustomProperty51 = util.SafeString(roleDetails.CustomProperty51)
	state.CustomProperty52 = util.SafeString(roleDetails.CustomProperty52)
	state.CustomProperty53 = util.SafeString(roleDetails.CustomProperty53)
	state.CustomProperty54 = util.SafeString(roleDetails.CustomProperty54)
	state.CustomProperty55 = util.SafeString(roleDetails.CustomProperty55)
	state.CustomProperty56 = util.SafeString(roleDetails.CustomProperty56)
	state.CustomProperty57 = util.SafeString(roleDetails.CustomProperty57)
	state.CustomProperty58 = util.SafeString(roleDetails.CustomProperty58)
	state.CustomProperty59 = util.SafeString(roleDetails.CustomProperty59)
	state.CustomProperty60 = util.SafeString(roleDetails.CustomProperty60)
	stateDiagnostics := resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(stateDiagnostics...)

	if resp.Diagnostics.HasError() {
		return
	}
}

// Update implements the resource.Resource interface for updating an existing role in Saviynt.
func (r *rolesResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan RolesResourceModel
	var state RolesResourceModel

	// Extract state from request
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Extract plan from request
	planGetDiagnostics := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(planGetDiagnostics...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Validate that the role name has not changed.
	if plan.RoleName.ValueString() != state.RoleName.ValueString() {
		resp.Diagnostics.AddError("Error", "Role name cannot be updated")
		log.Printf("[ERROR]: Role name cannot be updated")
		return
	}

	// Initialize OpenAPI Client Configuration.
	cfg := openapi.NewConfiguration()
	apiBaseURL := strings.TrimPrefix(strings.TrimPrefix(r.client.APIBaseURL(), "https://"), "http://")
	cfg.Host = apiBaseURL
	cfg.Scheme = "https"
	cfg.AddDefaultHeader("Authorization", "Bearer "+r.token)
	cfg.HTTPClient = http.DefaultClient
	apiClient := openapi.NewAPIClient(cfg)

	// Preparing Entitlement list for the API call
	var stateEntitlements []Entitlement
	stateEntitlementsdiags := state.Entitlements.ElementsAs(ctx, &stateEntitlements, false)
	resp.Diagnostics.Append(stateEntitlementsdiags...)
	if resp.Diagnostics.HasError() {
		return
	}
	oldEntitlementMap := make(map[string]Entitlement, len(stateEntitlements))
	for _, o := range stateEntitlements {
		oldEntitlementMap[o.EntitlementValue.ValueString()] = o
	}
	var planEntitlements []Entitlement
	planEntitlementsdiags := plan.Entitlements.ElementsAs(ctx, &planEntitlements, false)
	resp.Diagnostics.Append(planEntitlementsdiags...)
	if resp.Diagnostics.HasError() {
		return
	}
	newEntitlementMap := make(map[string]Entitlement, len(planEntitlements))
	for _, o := range planEntitlements {
		newEntitlementMap[o.EntitlementValue.ValueString()] = o
	}
	var removedEntitlement []Entitlement
	var addEntitlement []Entitlement
	for name, oldO := range oldEntitlementMap {
		if _, stillThere := newEntitlementMap[name]; !stillThere {
			removedEntitlement = append(removedEntitlement, oldO)
		}
	}
	for name, newO := range newEntitlementMap {
		if _, stillThere := oldEntitlementMap[name]; !stillThere {
			addEntitlement = append(addEntitlement, newO)
		}
	}
	var entitlements []openapi.UpdateEntitlementPayload
	for _, removedEntitlement := range removedEntitlement {
		entitlement := openapi.UpdateEntitlementPayload{}
		entitlement.EntitlementValue = removedEntitlement.EntitlementValue.ValueStringPointer()
		entitlement.EntitlementType = removedEntitlement.EntitlementType.ValueStringPointer()
		entitlement.Endpoint = removedEntitlement.Endpoint.ValueStringPointer()
		update := "REMOVE"
		entitlement.UpdateType = &update
		entitlements = append(entitlements, entitlement)
	}
	for _, addEntitlement := range addEntitlement {
		entitlement := openapi.UpdateEntitlementPayload{}
		entitlement.EntitlementValue = addEntitlement.EntitlementValue.ValueStringPointer()
		entitlement.EntitlementType = addEntitlement.EntitlementType.ValueStringPointer()
		entitlement.Endpoint = addEntitlement.Endpoint.ValueStringPointer()
		update := "ADD"
		entitlement.UpdateType = &update
		entitlements = append(entitlements, entitlement)
	}

	// Preparing Owner list for the API call
	var stateOwners []Owner
	statediags := state.Owners.ElementsAs(ctx, &stateOwners, false)
	resp.Diagnostics.Append(statediags...)
	if resp.Diagnostics.HasError() {
		return
	}
	oldOwnerMap := make(map[string]Owner, len(stateOwners))
	for _, o := range stateOwners {
		oldOwnerMap[o.OwnerName.ValueString()] = o
	}
	var planOwners []Owner
	plandiags := plan.Owners.ElementsAs(ctx, &planOwners, false)
	resp.Diagnostics.Append(plandiags...)
	if resp.Diagnostics.HasError() {
		return
	}
	newOwnerMap := make(map[string]Owner, len(planOwners))
	for _, o := range planOwners {
		newOwnerMap[o.OwnerName.ValueString()] = o
	}
	var removedOwners []Owner
	var addOwners []Owner
	for name, oldO := range oldOwnerMap {
		if _, stillThere := newOwnerMap[name]; !stillThere {
			removedOwners = append(removedOwners, oldO)
		}
	}
	for name, newO := range newOwnerMap {
		if _, stillThere := oldOwnerMap[name]; !stillThere {
			addOwners = append(addOwners, newO)
		}
	}
	var owners []openapi.UpdateRoleOwnerPayload
	for _, removedOwner := range removedOwners {
		owner := openapi.UpdateRoleOwnerPayload{}
		owner.OwnerName = removedOwner.OwnerName.ValueStringPointer()
		owner.Rank = removedOwner.Rank.ValueStringPointer()
		update := "REMOVE"
		owner.UpdateType = &update
		owners = append(owners, owner)
	}
	for _, addOwner := range addOwners {
		owner := openapi.UpdateRoleOwnerPayload{}
		owner.OwnerName = addOwner.OwnerName.ValueStringPointer()
		owner.Rank = addOwner.Rank.ValueStringPointer()
		update := "ADD"
		owner.UpdateType = &update
		owners = append(owners, owner)
	}

	// Preparing the request for the API call
	updateReq := openapi.UpdateEnterpriseRoleRequest{
		// Required fields
		RoleName: plan.RoleName.ValueString(),
		Roletype: plan.RoleType.ValueString(),
		// Optional fields
		Requestor:        util.StringPointerOrEmpty(plan.Requestor),
		Owner:            owners,
		Entitlements:     entitlements,
		Endpointname:     util.StringPointerOrEmpty(plan.EndpointName),
		Defaulttimeframe: util.StringPointerOrEmpty(plan.DefaultTimeFrame),
		Description:      util.StringPointerOrEmpty(plan.Description),
		Displayname:      util.StringPointerOrEmpty(plan.DisplayName),
		Glossary:         util.StringPointerOrEmpty(plan.Glossary),
		Risk:             util.StringPointerOrEmpty(plan.Risk),
		Level:            util.StringPointerOrEmpty(plan.Level),
		Soxcritical:      util.StringPointerOrEmpty(plan.SoxCritical),
		Syscritical:      util.StringPointerOrEmpty(plan.SysCritical),
		Priviliged:       util.StringPointerOrEmpty(plan.Priviliged),
		Confidentiality:  util.StringPointerOrEmpty(plan.Confidentiality),
		Requestable:      util.StringPointerOrEmpty(plan.Requestable),
		ShowDynamicAttrs: util.StringPointerOrEmpty(plan.ShowDynamicAttrs),
		Customproperty1:  util.StringPointerOrEmpty(plan.CustomProperty1),
		Customproperty2:  util.StringPointerOrEmpty(plan.CustomProperty2),
		Customproperty3:  util.StringPointerOrEmpty(plan.CustomProperty3),
		Customproperty4:  util.StringPointerOrEmpty(plan.CustomProperty4),
		Customproperty5:  util.StringPointerOrEmpty(plan.CustomProperty5),
		Customproperty6:  util.StringPointerOrEmpty(plan.CustomProperty6),
		Customproperty7:  util.StringPointerOrEmpty(plan.CustomProperty7),
		Customproperty8:  util.StringPointerOrEmpty(plan.CustomProperty8),
		Customproperty9:  util.StringPointerOrEmpty(plan.CustomProperty9),
		Customproperty10: util.StringPointerOrEmpty(plan.CustomProperty10),
		Customproperty11: util.StringPointerOrEmpty(plan.CustomProperty11),
		Customproperty12: util.StringPointerOrEmpty(plan.CustomProperty12),
		Customproperty13: util.StringPointerOrEmpty(plan.CustomProperty13),
		Customproperty14: util.StringPointerOrEmpty(plan.CustomProperty14),
		Customproperty15: util.StringPointerOrEmpty(plan.CustomProperty15),
		Customproperty16: util.StringPointerOrEmpty(plan.CustomProperty16),
		Customproperty17: util.StringPointerOrEmpty(plan.CustomProperty17),
		Customproperty18: util.StringPointerOrEmpty(plan.CustomProperty18),
		Customproperty19: util.StringPointerOrEmpty(plan.CustomProperty19),
		Customproperty20: util.StringPointerOrEmpty(plan.CustomProperty20),
		Customproperty21: util.StringPointerOrEmpty(plan.CustomProperty21),
		Customproperty22: util.StringPointerOrEmpty(plan.CustomProperty22),
		Customproperty23: util.StringPointerOrEmpty(plan.CustomProperty23),
		Customproperty24: util.StringPointerOrEmpty(plan.CustomProperty24),
		Customproperty25: util.StringPointerOrEmpty(plan.CustomProperty25),
		Customproperty26: util.StringPointerOrEmpty(plan.CustomProperty26),
		Customproperty27: util.StringPointerOrEmpty(plan.CustomProperty27),
		Customproperty28: util.StringPointerOrEmpty(plan.CustomProperty28),
		Customproperty29: util.StringPointerOrEmpty(plan.CustomProperty29),
		Customproperty30: util.StringPointerOrEmpty(plan.CustomProperty30),
		Customproperty31: util.StringPointerOrEmpty(plan.CustomProperty31),
		Customproperty32: util.StringPointerOrEmpty(plan.CustomProperty32),
		Customproperty33: util.StringPointerOrEmpty(plan.CustomProperty33),
		Customproperty34: util.StringPointerOrEmpty(plan.CustomProperty34),
		Customproperty35: util.StringPointerOrEmpty(plan.CustomProperty35),
		Customproperty36: util.StringPointerOrEmpty(plan.CustomProperty36),
		Customproperty37: util.StringPointerOrEmpty(plan.CustomProperty37),
		Customproperty38: util.StringPointerOrEmpty(plan.CustomProperty38),
		Customproperty39: util.StringPointerOrEmpty(plan.CustomProperty39),
		Customproperty40: util.StringPointerOrEmpty(plan.CustomProperty40),
		Customproperty41: util.StringPointerOrEmpty(plan.CustomProperty41),
		Customproperty42: util.StringPointerOrEmpty(plan.CustomProperty42),
		Customproperty43: util.StringPointerOrEmpty(plan.CustomProperty43),
		Customproperty44: util.StringPointerOrEmpty(plan.CustomProperty44),
		Customproperty45: util.StringPointerOrEmpty(plan.CustomProperty45),
		Customproperty46: util.StringPointerOrEmpty(plan.CustomProperty46),
		Customproperty47: util.StringPointerOrEmpty(plan.CustomProperty47),
		Customproperty48: util.StringPointerOrEmpty(plan.CustomProperty48),
		Customproperty49: util.StringPointerOrEmpty(plan.CustomProperty49),
		Customproperty50: util.StringPointerOrEmpty(plan.CustomProperty50),
		Customproperty51: util.StringPointerOrEmpty(plan.CustomProperty51),
		Customproperty52: util.StringPointerOrEmpty(plan.CustomProperty52),
		Customproperty53: util.StringPointerOrEmpty(plan.CustomProperty53),
		Customproperty54: util.StringPointerOrEmpty(plan.CustomProperty54),
		Customproperty55: util.StringPointerOrEmpty(plan.CustomProperty55),
		Customproperty56: util.StringPointerOrEmpty(plan.CustomProperty56),
		Customproperty57: util.StringPointerOrEmpty(plan.CustomProperty57),
		Customproperty58: util.StringPointerOrEmpty(plan.CustomProperty58),
		Customproperty59: util.StringPointerOrEmpty(plan.CustomProperty59),
		Customproperty60: util.StringPointerOrEmpty(plan.CustomProperty60),
	}

	// Execute the update API call.
	apiResp, https, err := apiClient.RolesAPI.UpdateEnterpriseRoleRequest(ctx).UpdateEnterpriseRoleRequest(updateReq).Execute()

	// Check for the error other than the 200 OK response
	if https != nil && https.StatusCode != http.StatusOK {
		log.Printf("[ERROR] HTTP error while creating role: %s", https.Status)
		var fetchResp map[string]interface{}
		if err := json.NewDecoder(https.Body).Decode(&fetchResp); err != nil {
			resp.Diagnostics.AddError("Failed to decode error response", err.Error())
			return
		}
		resp.Diagnostics.AddError(
			"HTTP Error",
			fmt.Sprintf("HTTP error while creating role for the reasons: %s", fetchResp["message"]),
		)
		return
	}

	if err != nil || *apiResp.ErrorCode != "0" && *apiResp.Message != "" {
		b, err := json.MarshalIndent(updateReq, "", "  ")
		if err != nil {
			log.Printf("Failed to marshal updateReq: %v", err)
			return
		}
		fmt.Println("data to be printed", string(b))
		log.Printf("[ERROR] Failed to update API resource. Error: %v %v %v", err, *apiResp.Message, *apiResp.ErrorCode)
		resp.Diagnostics.AddError("API Update Failed", fmt.Sprintf("Error: %v", *apiResp.Message))
		return
	}

	reqParams := openapi.GetRolesRequest{}
	reqParams.SetRoleName(plan.RoleName.ValueString())
	reqParams.SetRequestedObject("entitlements")
	readResp, httpResp, err := apiClient.RolesAPI.GetRoles(ctx).GetRolesRequest(reqParams).Execute()

	// For the fetching endpoint name using the endpoint key
	if readResp.Roledetails[0].Endpointkey != nil {
		endpointcfg := endpoint.NewConfiguration()
		endpointcfg.Host = apiBaseURL
		endpointcfg.Scheme = "https"
		endpointcfg.AddDefaultHeader("Authorization", "Bearer "+r.token)
		endpointcfg.HTTPClient = http.DefaultClient

		endpointClient := endpoint.NewAPIClient(endpointcfg)
		endpointReq := endpoint.GetEndpointsRequest{}
		var endpointKeys []string
		endpointKeys = []string{*readResp.Roledetails[0].Endpointkey}
		endpointReq.Endpointkey = endpointKeys
		endpointResp, httpRespEndpoint, err := endpointClient.EndpointsAPI.GetEndpoints(ctx).GetEndpointsRequest(endpointReq).Execute()
		if err != nil {
			log.Printf("[ERROR] Failed to fetch endpoint details: %v", err)
			resp.Diagnostics.AddError("API Read Failed", fmt.Sprintf("Error fetching endpoint details: %v", err))
			return
		}
		if httpRespEndpoint != nil && httpRespEndpoint.StatusCode != http.StatusOK {
			log.Printf("[ERROR] HTTP error while reading endpoint: %s", httpRespEndpoint.Status)
			var fetchResp map[string]interface{}
			if err := json.NewDecoder(httpRespEndpoint.Body).Decode(&fetchResp); err != nil {
				resp.Diagnostics.AddError("Failed to decode error response", err.Error())
				return
			}
			resp.Diagnostics.AddError(
				"HTTP Error",
				fmt.Sprintf("HTTP error while reading endpoint: %s", fetchResp["message"]),
			)
			return
		}
		if len(endpointResp.Endpoints) > 0 {
			plan.EndpointName = types.StringValue(*endpointResp.Endpoints[0].Endpointname)
		} else {
			log.Println("No endpoint found for the role")
			plan.EndpointName = types.StringNull()
		}
	}

	// For the error other than the 200 OK response
	if httpResp != nil && httpResp.StatusCode != http.StatusOK {
		log.Printf("[ERROR] HTTP error while reading role: %s", httpResp.Status)
		var fetchResp map[string]interface{}
		if err := json.NewDecoder(httpResp.Body).Decode(&fetchResp); err != nil {
			resp.Diagnostics.AddError("Failed to decode error response", err.Error())
			return
		}
		resp.Diagnostics.AddError(
			"HTTP Error",
			fmt.Sprintf("HTTP error while reading role: %s", fetchResp["message"]),
		)
		return
	}

	if err != nil {
		log.Printf("Problem with the get function in read block: %v", err)
		resp.Diagnostics.AddError("API Read Failed", fmt.Sprintf("Error: %v", err))
		return
	}

	// Set the ID from the role name
	plan.ID = types.StringValue("roles-" + plan.RoleName.ValueString())
	roleDetails := readResp.Roledetails[0]

	// Set basic role properties
	plan.RequestId = util.SafeString(apiResp.Requestid)
	plan.Requestkey = util.SafeString(apiResp.Requestkey)
	plan.Msg = util.SafeString(apiResp.Message)
	plan.ErrorCode = util.SafeString(apiResp.ErrorCode)
	if roleDetails.Roletype != nil {
		plan.RoleType = types.StringValue(endpointsutil.TranslateValue(*roleDetails.Roletype, rolesutil.RoleTypeMap))
	}
	plan.RoleName = util.SafeString(roleDetails.RoleName)
	plan.Description = util.SafeString(roleDetails.Description)
	plan.DisplayName = util.SafeString(roleDetails.Displayname)
	plan.DefaultTimeFrame = util.SafeString(roleDetails.DefaultTimeFrameHrs)
	plan.Glossary = util.SafeString(roleDetails.Glossary)

	// Mapping of soxcritical, syscritical, and other role attributes
	if roleDetails.Soxcritical != nil {
		plan.SoxCritical = types.StringValue(endpointsutil.TranslateValue(*roleDetails.Soxcritical, rolesutil.SoxCriticalityMap))
	}
	if roleDetails.Soxcritical != nil {
		plan.SoxCritical = types.StringValue(endpointsutil.TranslateValue(*roleDetails.Soxcritical, rolesutil.SoxCriticalityMap))
	}
	if roleDetails.Syscritical != nil {
		plan.SysCritical = types.StringValue(endpointsutil.TranslateValue(*roleDetails.Syscritical, rolesutil.SysCriticalMap))
	}
	if roleDetails.Priviliged != nil {
		plan.Priviliged = types.StringValue(endpointsutil.TranslateValue(*roleDetails.Priviliged, rolesutil.PrivilegedMap))
	}
	if roleDetails.Confidentiality != nil {
		plan.Confidentiality = types.StringValue(endpointsutil.TranslateValue(*roleDetails.Confidentiality, rolesutil.ConfidentialityMap))
	}
	plan.Requestable = util.SafeString(roleDetails.Requestable)
	plan.ShowDynamicAttrs = util.SafeString(roleDetails.ShowDynamicAttrs)

	// Process owners if available in the response
	var ownersArr []openapi.GetRoleOwnersResponse
	ownerUnion := roleDetails.Owner
	if ownerUnion != nil {
		if ownerUnion.ArrayOfGetRoleOwnersResponse != nil {
			ownersArr = *ownerUnion.ArrayOfGetRoleOwnersResponse
		}
		if ownerUnion.String != nil {
			var tmp []openapi.GetRoleOwnersResponse
			ownersArr = tmp
		}
	}
	ownerObjects := make([]attr.Value, 0, len(ownersArr))
	for _, o := range ownersArr {
		objOwnerVal, objOwnerDiags := types.ObjectValue(
			map[string]attr.Type{
				"owner_name": types.StringType,
				"rank":       types.StringType,
			},
			map[string]attr.Value{
				"owner_name": types.StringValue(*o.Ownername),
				"rank":       types.StringValue(util.SafeDeref(o.Rank)),
			},
		)
		if objOwnerDiags.HasError() {
			resp.Diagnostics.Append(objOwnerDiags...)
			continue
		}
		ownerObjects = append(ownerObjects, objOwnerVal)
	}

	setOwnerVal, setOwnerDiags := types.SetValue(
		types.ObjectType{
			AttrTypes: map[string]attr.Type{
				"owner_name": types.StringType,
				"rank":       types.StringType,
			},
		}, ownerObjects,
	)
	if setOwnerDiags.HasError() {
		resp.Diagnostics.Append(setOwnerDiags...)
		return
	}
	plan.Owners = setOwnerVal

	// Process entitlement if available in the response
	entitlementsArr := make([]openapi.GetEntitlementDetailsResponse, 0)
	entitlementUnion := roleDetails.EntitlementDetails
	if entitlementUnion != nil {
		entitlementsArr = entitlementUnion
	}
	entitlementObjects := make([]attr.Value, 0, len(entitlementsArr))
	for _, o := range entitlementsArr {
		entitlementVal, entitlementDiags := types.ObjectValue(
			map[string]attr.Type{
				"entitlement_value": types.StringType,
				"entitlement_type":  types.StringType,
				"endpoint":          types.StringType,
			},
			map[string]attr.Value{
				"entitlement_value": types.StringValue(*o.EntitlementValue),
				"entitlement_type":  types.StringValue(util.SafeDeref(o.EntitlementTypeName)),
				"endpoint":          types.StringValue(util.SafeDeref(o.Endpoint)),
			},
		)
		if entitlementDiags.HasError() {
			resp.Diagnostics.Append(entitlementDiags...)
			continue
		}
		entitlementObjects = append(entitlementObjects, entitlementVal)
	}

	setEntitlementVal, setEntitlementDiags := types.SetValue(
		types.ObjectType{
			AttrTypes: map[string]attr.Type{
				"entitlement_value": types.StringType,
				"entitlement_type":  types.StringType,
				"endpoint":          types.StringType,
			},
		}, entitlementObjects,
	)
	if setEntitlementDiags.HasError() {
		resp.Diagnostics.Append(setEntitlementDiags...)
		return
	}
	plan.Entitlements = setEntitlementVal

	// Setting the Custom Properties
	plan.CustomProperty1 = util.SafeString(roleDetails.CustomProperty1)
	plan.CustomProperty2 = util.SafeString(roleDetails.CustomProperty2)
	plan.CustomProperty3 = util.SafeString(roleDetails.CustomProperty3)
	plan.CustomProperty4 = util.SafeString(roleDetails.CustomProperty4)
	plan.CustomProperty5 = util.SafeString(roleDetails.CustomProperty5)
	plan.CustomProperty6 = util.SafeString(roleDetails.CustomProperty6)
	plan.CustomProperty7 = util.SafeString(roleDetails.CustomProperty7)
	plan.CustomProperty8 = util.SafeString(roleDetails.CustomProperty8)
	plan.CustomProperty9 = util.SafeString(roleDetails.CustomProperty9)
	plan.CustomProperty10 = util.SafeString(roleDetails.CustomProperty10)
	plan.CustomProperty11 = util.SafeString(roleDetails.CustomProperty11)
	plan.CustomProperty12 = util.SafeString(roleDetails.CustomProperty12)
	plan.CustomProperty13 = util.SafeString(roleDetails.CustomProperty13)
	plan.CustomProperty14 = util.SafeString(roleDetails.CustomProperty14)
	plan.CustomProperty15 = util.SafeString(roleDetails.CustomProperty15)
	plan.CustomProperty16 = util.SafeString(roleDetails.CustomProperty16)
	plan.CustomProperty17 = util.SafeString(roleDetails.CustomProperty17)
	plan.CustomProperty18 = util.SafeString(roleDetails.CustomProperty18)
	plan.CustomProperty19 = util.SafeString(roleDetails.CustomProperty19)
	plan.CustomProperty20 = util.SafeString(roleDetails.CustomProperty20)
	plan.CustomProperty21 = util.SafeString(roleDetails.CustomProperty21)
	plan.CustomProperty22 = util.SafeString(roleDetails.CustomProperty22)
	plan.CustomProperty23 = util.SafeString(roleDetails.CustomProperty23)
	plan.CustomProperty24 = util.SafeString(roleDetails.CustomProperty24)
	plan.CustomProperty25 = util.SafeString(roleDetails.CustomProperty25)
	plan.CustomProperty26 = util.SafeString(roleDetails.CustomProperty26)
	plan.CustomProperty27 = util.SafeString(roleDetails.CustomProperty27)
	plan.CustomProperty28 = util.SafeString(roleDetails.CustomProperty28)
	plan.CustomProperty29 = util.SafeString(roleDetails.CustomProperty29)
	plan.CustomProperty30 = util.SafeString(roleDetails.CustomProperty30)
	plan.CustomProperty31 = util.SafeString(roleDetails.CustomProperty31)
	plan.CustomProperty32 = util.SafeString(roleDetails.CustomProperty32)
	plan.CustomProperty33 = util.SafeString(roleDetails.CustomProperty33)
	plan.CustomProperty34 = util.SafeString(roleDetails.CustomProperty34)
	plan.CustomProperty35 = util.SafeString(roleDetails.CustomProperty35)
	plan.CustomProperty36 = util.SafeString(roleDetails.CustomProperty36)
	plan.CustomProperty37 = util.SafeString(roleDetails.CustomProperty37)
	plan.CustomProperty38 = util.SafeString(roleDetails.CustomProperty38)
	plan.CustomProperty39 = util.SafeString(roleDetails.CustomProperty39)
	plan.CustomProperty40 = util.SafeString(roleDetails.CustomProperty40)
	plan.CustomProperty41 = util.SafeString(roleDetails.CustomProperty41)
	plan.CustomProperty42 = util.SafeString(roleDetails.CustomProperty42)
	plan.CustomProperty43 = util.SafeString(roleDetails.CustomProperty43)
	plan.CustomProperty44 = util.SafeString(roleDetails.CustomProperty44)
	plan.CustomProperty45 = util.SafeString(roleDetails.CustomProperty45)
	plan.CustomProperty46 = util.SafeString(roleDetails.CustomProperty46)
	plan.CustomProperty47 = util.SafeString(roleDetails.CustomProperty47)
	plan.CustomProperty48 = util.SafeString(roleDetails.CustomProperty48)
	plan.CustomProperty49 = util.SafeString(roleDetails.CustomProperty49)
	plan.CustomProperty50 = util.SafeString(roleDetails.CustomProperty50)
	plan.CustomProperty51 = util.SafeString(roleDetails.CustomProperty51)
	plan.CustomProperty52 = util.SafeString(roleDetails.CustomProperty52)
	plan.CustomProperty53 = util.SafeString(roleDetails.CustomProperty53)
	plan.CustomProperty54 = util.SafeString(roleDetails.CustomProperty54)
	plan.CustomProperty55 = util.SafeString(roleDetails.CustomProperty55)
	plan.CustomProperty56 = util.SafeString(roleDetails.CustomProperty56)
	plan.CustomProperty57 = util.SafeString(roleDetails.CustomProperty57)
	plan.CustomProperty58 = util.SafeString(roleDetails.CustomProperty58)
	plan.CustomProperty59 = util.SafeString(roleDetails.CustomProperty59)
	plan.CustomProperty60 = util.SafeString(roleDetails.CustomProperty60)
	stateUpdateDiagnostics := resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(stateUpdateDiagnostics...)
}

// We do not support deletion of roles in Saviynt, so this function is intentionally left empty.
func (r *rolesResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
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

// ImportState implements the resource.Resource interface for importing existing roles into the Terraform state.
func (r *rolesResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	// Retrieve import ID and save to id attribute
	resource.ImportStatePassthroughID(ctx, path.Root("role_name"), req, resp)
}
