// Copyright (c) 2025 Saviynt Inc.
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
	"fmt"
	"os"
	"strings"

	"terraform-provider-Saviynt/internal/client"
	"terraform-provider-Saviynt/util"
	endpointsutil "terraform-provider-Saviynt/util/endpointsutil"
	"terraform-provider-Saviynt/util/rolesutil"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	endpoint "github.com/saviynt/saviynt-api-go-client/endpoints"
	openapi "github.com/saviynt/saviynt-api-go-client/roles"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &RolesResource{}
var _ resource.ResourceWithImportState = &RolesResource{}

// Operation constants for update types
const (
	UpdateTypeAdd    = "ADD"
	UpdateTypeRemove = "REMOVE"
)

// ChangeProcessor is a generic processor for handling add/remove operations on collections
// T: the item type (e.g., Entitlement, Owner, ChildRoles)
// K: the key type used for comparison (typically string)
// P: the payload type for updates (e.g., UpdateEntitlementPayload)
type ChangeProcessor[T any, K comparable, P any] struct {
	// KeyExtractor extracts the comparison key from an item
	KeyExtractor func(T) K

	// PayloadBuilder creates an update payload from an item and operation type
	PayloadBuilder func(T, string) P

	// ErrorContext provides context for error messages
	ErrorContext string
}

// ProcessChanges processes changes between state and plan collections
// Returns a list of update payloads for items that were added or removed
func (cp *ChangeProcessor[T, K, P]) ProcessChanges(
	ctx context.Context,
	stateItems types.Set,
	planItems types.Set,
) ([]P, error) {
	// Extract state items
	var stateItemsList []T
	if !stateItems.IsNull() && !stateItems.IsUnknown() {
		if err := stateItems.ElementsAs(ctx, &stateItemsList, false); err != nil {
			return nil, fmt.Errorf("failed to extract state %s: %v", cp.ErrorContext, err)
		}
	}

	// Extract plan items
	var planItemsList []T
	if !planItems.IsNull() && !planItems.IsUnknown() {
		if err := planItems.ElementsAs(ctx, &planItemsList, false); err != nil {
			return nil, fmt.Errorf("failed to extract plan %s: %v", cp.ErrorContext, err)
		}
	}

	// Create maps for efficient comparison
	oldItemMap := make(map[K]T, len(stateItemsList))
	for _, item := range stateItemsList {
		key := cp.KeyExtractor(item)
		oldItemMap[key] = item
	}

	newItemMap := make(map[K]T, len(planItemsList))
	for _, item := range planItemsList {
		key := cp.KeyExtractor(item)
		newItemMap[key] = item
	}

	var payloads []P

	// Find removed items (in state but not in plan)
	for key, oldItem := range oldItemMap {
		if _, stillExists := newItemMap[key]; !stillExists {
			payload := cp.PayloadBuilder(oldItem, UpdateTypeRemove)
			payloads = append(payloads, payload)
		}
	}

	// Find added items (in plan but not in state)
	for key, newItem := range newItemMap {
		if _, existed := oldItemMap[key]; !existed {
			payload := cp.PayloadBuilder(newItem, UpdateTypeAdd)
			payloads = append(payloads, payload)
		}
	}

	return payloads, nil
}

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
	Privileged       types.String `tfsdk:"privileged"`
	Confidentiality  types.String `tfsdk:"confidentiality"`
	Requestable      types.String `tfsdk:"requestable"`
	ShowDynamicAttrs types.String `tfsdk:"show_dynamic_attrs"`
	CheckSod         types.String `tfsdk:"check_sod"`
	Entitlements     types.Set    `tfsdk:"entitlements"`
	ChildRoles       types.Set    `tfsdk:"child_roles"`
	Users            types.Set    `tfsdk:"users"`
}

// Entitlement defines the structure for entitlements associated with a role.
type Entitlement struct {
	EntitlementValue types.String `tfsdk:"entitlement_value"`
	EntitlementType  types.String `tfsdk:"entitlement_type"`
	Endpoint         types.String `tfsdk:"endpoint"`
}

// Owner defines the structure for owners associated with a role.
type Owner struct {
	OwnerName types.String `tfsdk:"owner_name"`
	Rank      types.String `tfsdk:"rank"`
}

// ChildRole defines the structure for childRole associated with a role.
type ChildRoles struct {
	RoleName types.String `tfsdk:"role_name"`
}

// Users defines the structure for users associated with a role.
type Users struct {
	UserName types.String `tfsdk:"user_name"`
}

// rolesResource implements the resource.Resource interface for managing Saviynt roles.
type RolesResource struct {
	client      client.SaviyntClientInterface
	token       string
	requestor   string
	roleFactory client.RoleFactoryInterface
}

// NewRolesResource creates a new instance of the roles resource.
func NewRolesResource() resource.Resource {
	return &RolesResource{
		roleFactory: &client.DefaultRoleFactory{},
	}
}

func NewRolesResourceWithFactory(factory client.RoleFactoryInterface) resource.Resource {
	return &RolesResource{
		roleFactory: factory,
	}
}

// Metadata returns the metadata for the roles resource.
func (r *RolesResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = "saviynt_enterprise_roles_resource"
}

// Schema defines the schema for the roles resource.
func (r *RolesResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: util.RoleDescription,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "Unique identifier of the role resource. This is automatically generated as 'roles-' + role_name.",
			},
			"role_type": schema.StringAttribute{
				Description: "Type of the role. Valid values: 'ENABLER', 'TRANSACTIONAL', 'FIREFIGHTER', 'ENTERPRISE', 'APPLICATION', 'ENTITLEMENT'.",
				Required:    true,
			},
			"role_name": schema.StringAttribute{
				Description: "Unique name of the role. This will be used as the identifier for the role in Saviynt and must be unique across all roles.",
				Required:    true,
			},
			"requestor": schema.StringAttribute{
				Description: "Username of the person requesting the role creation. This should be a valid Saviynt user who has permissions to create roles.",
				Required:    true,
			},
			"owners": schema.SetAttribute{
				Description: "Set of role owners with their respective ranks. Each owner must have 'owner_name' (valid Saviynt username) and 'rank' (1-27, where 1 is highest priority). The same owner can have up to 5 different ranks. To add owners, include them in the set; to remove owners, exclude them from the set.",
				Optional:    true,
				ElementType: types.ObjectType{
					AttrTypes: map[string]attr.Type{
						"owner_name": types.StringType,
						"rank":       types.StringType,
					},
				},
				Validators: []validator.Set{
					rolesutil.OwnerNameAddLimit(),
					rolesutil.OwnerRankValidator(),
				},
			},
			"endpoint_name": schema.StringAttribute{
				Description: "Name of the endpoint associated with this role. Must be an existing endpoint in Saviynt.",
				Optional:    true,
				Computed:    true,
			},
			"default_time_frame": schema.StringAttribute{
				Description: "Specify the default time frame (in hours) to request access for a role. This defines how long users will have access when assigned this role.",
				Optional:    true,
				Computed:    true,
			},
			"description": schema.StringAttribute{
				Description: "Displays the description of the role. You can change the description, as required. This helps users understand what the role is for and what permissions it grants.",
				Optional:    true,
				Computed:    true,
			},
			"display_name": schema.StringAttribute{
				Description: "Displays the display name of the role. This is a user-friendly name that can be different from the role_name.",
				Optional:    true,
				Computed:    true,
			},
			"glossary": schema.StringAttribute{
				Description: "Displays the Glossary about the role. This provides additional context and definitions related to the role.",
				Optional:    true,
				Computed:    true,
			},
			"risk": schema.StringAttribute{
				Description: "Displays the risk level specified for the role required during the separation of duties based on the roles. Valid options: 'None', 'Very Low', 'Low', 'Medium', 'High', 'Critical'.",
				Optional:    true,
				Computed:    true,
			},
			"level": schema.StringAttribute{
				Description: "Enter the hierarchy level of this role. This defines the role's position in the organizational hierarchy.",
				Optional:    true,
			},
			"sox_critical": schema.StringAttribute{
				Description: "Select the SOX criticality of the role. Valid options: 'None', 'Very Low', 'Low', 'Medium', 'High', 'Critical'.",
				Optional:    true,
				Computed:    true,
			},
			"sys_critical": schema.StringAttribute{
				Description: "Select the SYS criticality of the role. Valid options: 'None', 'Very Low', 'Low', 'Medium', 'High', 'Critical'.",
				Optional:    true,
				Computed:    true,
			},
			"privileged": schema.StringAttribute{
				Description: "Select the privileged criticality of the role which describes privileges assigned to the role and amount of risk to provide access to this role. Valid options: 'None', 'Very Low', 'Low', 'Medium', 'High', 'Critical'.",
				Optional:    true,
				Computed:    true,
			},
			"confidentiality": schema.StringAttribute{
				Description: "Select the confidentiality of this role. Valid options: 'None', 'Very Low', 'Low', 'Medium', 'High', 'Critical'.",
				Optional:    true,
				Computed:    true,
			},
			"requestable": schema.StringAttribute{
				Description: "Specify if you want the users to request for the role. Valid options: 'true' (makes the role requestable), 'false' (makes the role non-requestable). Defaults to 'true'.",
				Optional:    true,
				Computed:    true,
				Default:     stringdefault.StaticString("true"),
			},
			"show_dynamic_attrs": schema.StringAttribute{
				Description: "Displays the dynamic attributes associated with the role. For example, there is a Dynamic Attribute A, which is Boolean set as true and false.",
				Optional:    true,
				Computed:    true,
				Default:     stringdefault.StaticString("true"),
			},
			"check_sod": schema.StringAttribute{
				Description: "Indicates if segregation of duties (SoD) checks should be performed",
				Optional:    true,
			},
			"entitlements": schema.SetAttribute{
				Description: "Set of entitlements associated with the role. Entitlements dictate user (role assignee) responsibility in managing an application. To add entitlements, include them in the set; to remove entitlements, exclude them from the set. Each entitlement requires 'entitlement_value', 'entitlement_type', and 'endpoint'.",
				Optional:    true,
				ElementType: types.ObjectType{
					AttrTypes: map[string]attr.Type{
						"entitlement_value": types.StringType,
						"entitlement_type":  types.StringType,
						"endpoint":          types.StringType,
					},
				},
			},
			"child_roles": schema.SetAttribute{
				Description: "Set of child roles associated with the role. Child roles provide conditional access - when a user requests entitlements from a child role, the system checks if they have the parent role entitlements. Conversely, requesting parent role entitlements automatically grants child role entitlements. To add child roles, include them in the set; to remove child roles, exclude them from the set. Each child role requires 'role_name'. Note: This attribute is only available in Saviynt version 25.B and later.",
				Optional:    true,
				ElementType: types.ObjectType{
					AttrTypes: map[string]attr.Type{
						"role_name": types.StringType,
					},
				},
				Validators: []validator.Set{
					rolesutil.ChildRoleValidator(),
				},
			},
			"users": schema.SetAttribute{
				Description: "Set of users assigned to the role. To add users to the role, include them in the set; to remove users from the role, exclude them from the set. Each user requires 'user_name' (valid Saviynt username). When users are assigned to a role, they inherit all entitlements and permissions associated with that role.",
				Optional:    true,
				ElementType: types.ObjectType{
					AttrTypes: map[string]attr.Type{
						"user_name": types.StringType,
					},
				},
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
func (r *RolesResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	tflog.Debug(ctx, "Starting roles resource configuration")

	// Check if provider data is available.
	if req.ProviderData == nil {
		tflog.Debug(ctx, "ProviderData is nil, returning early")
		return
	}

	// Cast provider data to your provider type.
	prov, ok := req.ProviderData.(*saviyntProvider)
	if !ok {
		tflog.Error(ctx, "Provider configuration failed - expected *saviyntProvider, got different type")
		resp.Diagnostics.AddError(
			"Unexpected Provider Data",
			"Expected *saviyntProvider, got different type",
		)
		return
	}

	// Set the client and token from the provider state using interface wrapper.
	r.client = &client.SaviyntClientWrapper{Client: prov.client}
	r.token = prov.accessToken
	if prov.client.Username != nil {
		r.requestor = *prov.client.Username
	}
	tflog.Debug(ctx, "Roles resource configured successfully")
}

// SetClient sets the client for testing purposes
func (r *RolesResource) SetClient(client client.SaviyntClientInterface) {
	r.client = client
}

// SetToken sets the token for testing purposes
func (r *RolesResource) SetToken(token string) {
	r.token = token
}

// SetRequestor sets the requestor for testing purposes
func (r *RolesResource) SetRequestor(requestor string) {
	r.requestor = requestor
}

// buildCustomPropertiesForCreate sets all custom properties for the create request
func (r *RolesResource) buildCustomPropertiesForCreate(plan *RolesResourceModel, createReq *openapi.CreateEnterpriseRoleRequest) {
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
	createReq.Customproperty46 = util.StringPointerOrEmpty(plan.CustomProperty46)
	createReq.Customproperty47 = util.StringPointerOrEmpty(plan.CustomProperty47)
	createReq.Customproperty48 = util.StringPointerOrEmpty(plan.CustomProperty48)
	createReq.Customproperty49 = util.StringPointerOrEmpty(plan.CustomProperty49)
	createReq.Customproperty50 = util.StringPointerOrEmpty(plan.CustomProperty50)
	createReq.Customproperty51 = util.StringPointerOrEmpty(plan.CustomProperty51)
	createReq.Customproperty52 = util.StringPointerOrEmpty(plan.CustomProperty52)
	createReq.Customproperty53 = util.StringPointerOrEmpty(plan.CustomProperty53)
	createReq.Customproperty54 = util.StringPointerOrEmpty(plan.CustomProperty54)
	createReq.Customproperty55 = util.StringPointerOrEmpty(plan.CustomProperty55)
	createReq.Customproperty56 = util.StringPointerOrEmpty(plan.CustomProperty56)
	createReq.Customproperty57 = util.StringPointerOrEmpty(plan.CustomProperty57)
	createReq.Customproperty58 = util.StringPointerOrEmpty(plan.CustomProperty58)
	createReq.Customproperty59 = util.StringPointerOrEmpty(plan.CustomProperty59)
	createReq.Customproperty60 = util.StringPointerOrEmpty(plan.CustomProperty60)
}

// setCustomPropertiesInPlan sets all custom properties in the plan from their current values
func (r *RolesResource) setCustomPropertiesInPlan(plan *RolesResourceModel) {
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
}

// setCustomPropertiesFromResponse sets all custom properties from API response to the model
func (r *RolesResource) setCustomPropertiesFromResponse(model *RolesResourceModel, roleDetails *openapi.GetRoleDetailsResponse) {
	model.CustomProperty1 = util.SafeString(roleDetails.CustomProperty1)
	model.CustomProperty2 = util.SafeString(roleDetails.CustomProperty2)
	model.CustomProperty3 = util.SafeString(roleDetails.CustomProperty3)
	model.CustomProperty4 = util.SafeString(roleDetails.CustomProperty4)
	model.CustomProperty5 = util.SafeString(roleDetails.CustomProperty5)
	model.CustomProperty6 = util.SafeString(roleDetails.CustomProperty6)
	model.CustomProperty7 = util.SafeString(roleDetails.CustomProperty7)
	model.CustomProperty8 = util.SafeString(roleDetails.CustomProperty8)
	model.CustomProperty9 = util.SafeString(roleDetails.CustomProperty9)
	model.CustomProperty10 = util.SafeString(roleDetails.CustomProperty10)
	model.CustomProperty11 = util.SafeString(roleDetails.CustomProperty11)
	model.CustomProperty12 = util.SafeString(roleDetails.CustomProperty12)
	model.CustomProperty13 = util.SafeString(roleDetails.CustomProperty13)
	model.CustomProperty14 = util.SafeString(roleDetails.CustomProperty14)
	model.CustomProperty15 = util.SafeString(roleDetails.CustomProperty15)
	model.CustomProperty16 = util.SafeString(roleDetails.CustomProperty16)
	model.CustomProperty17 = util.SafeString(roleDetails.CustomProperty17)
	model.CustomProperty18 = util.SafeString(roleDetails.CustomProperty18)
	model.CustomProperty19 = util.SafeString(roleDetails.CustomProperty19)
	model.CustomProperty20 = util.SafeString(roleDetails.CustomProperty20)
	model.CustomProperty21 = util.SafeString(roleDetails.CustomProperty21)
	model.CustomProperty22 = util.SafeString(roleDetails.CustomProperty22)
	model.CustomProperty23 = util.SafeString(roleDetails.CustomProperty23)
	model.CustomProperty24 = util.SafeString(roleDetails.CustomProperty24)
	model.CustomProperty25 = util.SafeString(roleDetails.CustomProperty25)
	model.CustomProperty26 = util.SafeString(roleDetails.CustomProperty26)
	model.CustomProperty27 = util.SafeString(roleDetails.CustomProperty27)
	model.CustomProperty28 = util.SafeString(roleDetails.CustomProperty28)
	model.CustomProperty29 = util.SafeString(roleDetails.CustomProperty29)
	model.CustomProperty30 = util.SafeString(roleDetails.CustomProperty30)
	model.CustomProperty31 = util.SafeString(roleDetails.CustomProperty31)
	model.CustomProperty32 = util.SafeString(roleDetails.CustomProperty32)
	model.CustomProperty33 = util.SafeString(roleDetails.CustomProperty33)
	model.CustomProperty34 = util.SafeString(roleDetails.CustomProperty34)
	model.CustomProperty35 = util.SafeString(roleDetails.CustomProperty35)
	model.CustomProperty36 = util.SafeString(roleDetails.CustomProperty36)
	model.CustomProperty37 = util.SafeString(roleDetails.CustomProperty37)
	model.CustomProperty38 = util.SafeString(roleDetails.CustomProperty38)
	model.CustomProperty39 = util.SafeString(roleDetails.CustomProperty39)
	model.CustomProperty40 = util.SafeString(roleDetails.CustomProperty40)
	model.CustomProperty41 = util.SafeString(roleDetails.CustomProperty41)
	model.CustomProperty42 = util.SafeString(roleDetails.CustomProperty42)
	model.CustomProperty43 = util.SafeString(roleDetails.CustomProperty43)
	model.CustomProperty44 = util.SafeString(roleDetails.CustomProperty44)
	model.CustomProperty45 = util.SafeString(roleDetails.CustomProperty45)
	model.CustomProperty46 = util.SafeString(roleDetails.CustomProperty46)
	model.CustomProperty47 = util.SafeString(roleDetails.CustomProperty47)
	model.CustomProperty48 = util.SafeString(roleDetails.CustomProperty48)
	model.CustomProperty49 = util.SafeString(roleDetails.CustomProperty49)
	model.CustomProperty50 = util.SafeString(roleDetails.CustomProperty50)
	model.CustomProperty51 = util.SafeString(roleDetails.CustomProperty51)
	model.CustomProperty52 = util.SafeString(roleDetails.CustomProperty52)
	model.CustomProperty53 = util.SafeString(roleDetails.CustomProperty53)
	model.CustomProperty54 = util.SafeString(roleDetails.CustomProperty54)
	model.CustomProperty55 = util.SafeString(roleDetails.CustomProperty55)
	model.CustomProperty56 = util.SafeString(roleDetails.CustomProperty56)
	model.CustomProperty57 = util.SafeString(roleDetails.CustomProperty57)
	model.CustomProperty58 = util.SafeString(roleDetails.CustomProperty58)
	model.CustomProperty59 = util.SafeString(roleDetails.CustomProperty59)
	model.CustomProperty60 = util.SafeString(roleDetails.CustomProperty60)
}

// buildCustomPropertiesForUpdate sets all custom properties for the update request
func (r *RolesResource) buildCustomPropertiesForUpdate(plan *RolesResourceModel, updateReq *openapi.UpdateEnterpriseRoleRequest) {
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
	updateReq.Customproperty46 = util.StringPointerOrEmpty(plan.CustomProperty46)
	updateReq.Customproperty47 = util.StringPointerOrEmpty(plan.CustomProperty47)
	updateReq.Customproperty48 = util.StringPointerOrEmpty(plan.CustomProperty48)
	updateReq.Customproperty49 = util.StringPointerOrEmpty(plan.CustomProperty49)
	updateReq.Customproperty50 = util.StringPointerOrEmpty(plan.CustomProperty50)
	updateReq.Customproperty51 = util.StringPointerOrEmpty(plan.CustomProperty51)
	updateReq.Customproperty52 = util.StringPointerOrEmpty(plan.CustomProperty52)
	updateReq.Customproperty53 = util.StringPointerOrEmpty(plan.CustomProperty53)
	updateReq.Customproperty54 = util.StringPointerOrEmpty(plan.CustomProperty54)
	updateReq.Customproperty55 = util.StringPointerOrEmpty(plan.CustomProperty55)
	updateReq.Customproperty56 = util.StringPointerOrEmpty(plan.CustomProperty56)
	updateReq.Customproperty57 = util.StringPointerOrEmpty(plan.CustomProperty57)
	updateReq.Customproperty58 = util.StringPointerOrEmpty(plan.CustomProperty58)
	updateReq.Customproperty59 = util.StringPointerOrEmpty(plan.CustomProperty59)
	updateReq.Customproperty60 = util.StringPointerOrEmpty(plan.CustomProperty60)
}

// RoleResourceMapRoleDetailsToModel maps role details from API response to the Terraform model
func (r *RolesResource) RoleResourceMapRoleDetailsToModel(model *RolesResourceModel, roleDetails *openapi.GetRoleDetailsResponse, state *RolesResourceModel, diagnostics *diag.Diagnostics, forceImport ...bool) error {
	// Check if this is an import operation based on ID field or forceImport flag
	isImport := state.ID.IsNull() || state.ID.IsUnknown() || state.ID.ValueString() == ""
	if len(forceImport) > 0 && forceImport[0] {
		isImport = true
	}
	// Set the ID from the role name
	model.ID = types.StringValue("roles-" + model.RoleName.ValueString())

	// Set basic role properties
	if roleDetails.Roletype != nil {
		model.RoleType = types.StringValue(endpointsutil.TranslateValue(*roleDetails.Roletype, rolesutil.RoleTypeMap))
	}

	model.RoleName = util.SafeString(roleDetails.RoleName)
	model.Description = util.SafeString(roleDetails.Description)
	model.DisplayName = util.SafeString(roleDetails.Displayname)
	model.DefaultTimeFrame = util.SafeString(roleDetails.DefaultTimeFrameHrs)
	model.Glossary = util.SafeString(roleDetails.Glossary)

	// Mapping of soxcritical, syscritical, and other role attributes
	if roleDetails.Soxcritical != nil {
		model.SoxCritical = types.StringValue(endpointsutil.TranslateValue(*roleDetails.Soxcritical, rolesutil.SoxCriticalityMap))
	}
	if roleDetails.Syscritical != nil {
		model.SysCritical = types.StringValue(endpointsutil.TranslateValue(*roleDetails.Syscritical, rolesutil.SysCriticalMap))
	}
	if roleDetails.Priviliged != nil {
		model.Privileged = types.StringValue(endpointsutil.TranslateValue(*roleDetails.Priviliged, rolesutil.PrivilegedMap))
	}
	if roleDetails.Confidentiality != nil {
		model.Confidentiality = types.StringValue(endpointsutil.TranslateValue(*roleDetails.Confidentiality, rolesutil.ConfidentialityMap))
	}
	if roleDetails.Risk != nil {
		model.Risk = types.StringValue(endpointsutil.TranslateValue(*roleDetails.Risk, rolesutil.RiskMap))
	}

	model.Requestable = util.SafeString(roleDetails.Requestable)
	model.ShowDynamicAttrs = util.SafeString(roleDetails.ShowDynamicAttrs)
	if !isImport && model == state {
		// This is a normal read, check for drift
		r.RoleResourceDetectDrift(roleDetails, state, diagnostics)
		if diagnostics.HasError() {
			return fmt.Errorf("drift detected, cannot proceed")
		}
	}

	if isImport {
		var ownersArr []openapi.GetRoleOwnersResponse
		ownerUnion := roleDetails.Owner
		if ownerUnion != nil {
			if ownerUnion.ArrayOfGetRoleOwnersResponse != nil {
				ownersArr = *ownerUnion.ArrayOfGetRoleOwnersResponse
			}
		}

		// Import: Use all owners from API
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
				diagnostics.Append(ownerDiags...)
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
			diagnostics.Append(setOwnerDiags...)
			return fmt.Errorf("failed to set owners")
		}

		// Only set owners if API has data OR state was previously configured
		if len(ownerObjects) > 0 {
			model.Owners = setOwnerVal
		} else if !state.Owners.IsNull() {
			// API has no owners but state was configured - set empty
			model.Owners = setOwnerVal
		}
		// If state was null and API has no owners, leave it null
	}
	if isImport {
		// Import: Use all entitlements from API
		var entitlementsArr []openapi.GetEntitlementDetailsResponse
		if roleDetails.EntitlementDetails != nil {
			entitlementsArr = roleDetails.EntitlementDetails
		}

		entitlementObjects := make([]attr.Value, 0, len(entitlementsArr))
		for _, e := range entitlementsArr {
			entitlementVal, entitlementDiags := types.ObjectValue(
				map[string]attr.Type{
					"entitlement_value": types.StringType,
					"entitlement_type":  types.StringType,
					"endpoint":          types.StringType,
				},
				map[string]attr.Value{
					"entitlement_value": types.StringValue(*e.EntitlementValue),
					"entitlement_type":  types.StringValue(util.SafeDeref(e.EntitlementTypeName)),
					"endpoint":          types.StringValue(util.SafeDeref(e.Endpoint)),
				},
			)
			if entitlementDiags.HasError() {
				diagnostics.Append(entitlementDiags...)
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
			diagnostics.Append(setEntitlementDiags...)
			return fmt.Errorf("failed to set entitlements")
		}

		// Only set entitlements if API has data OR state was previously configured
		if len(entitlementObjects) > 0 {
			model.Entitlements = setEntitlementVal
		} else if !state.Entitlements.IsNull() {
			// API has no entitlements but state was configured - set empty
			model.Entitlements = setEntitlementVal
		}
		// If state was null and API has no entitlements, leave it null
	}

	// Process users: only set if import (state.Users is null)
	if isImport {
		// Import: Use all users from API
		var userObjects []attr.Value
		if roleDetails.UserDetails != nil {
			for _, u := range roleDetails.UserDetails {
				// Extract username from user details - check for nil GetUserDetailsResponse
				if u.GetUserDetailsResponse != nil && u.GetUserDetailsResponse.Username != nil {
					userVal, userDiags := types.ObjectValue(
						map[string]attr.Type{
							"user_name": types.StringType,
						},
						map[string]attr.Value{
							"user_name": types.StringValue(*u.GetUserDetailsResponse.Username),
						},
					)
					if userDiags.HasError() {
						diagnostics.Append(userDiags...)
						continue
					}
					userObjects = append(userObjects, userVal)
				}
			}
		}

		setUserVal, setUserDiags := types.SetValue(
			types.ObjectType{
				AttrTypes: map[string]attr.Type{
					"user_name": types.StringType,
				},
			}, userObjects,
		)
		if setUserDiags.HasError() {
			diagnostics.Append(setUserDiags...)
			return fmt.Errorf("failed to set users")
		}

		// Only set users if API has data OR state was previously configured
		if len(userObjects) > 0 {
			model.Users = setUserVal
		} else if !state.Users.IsNull() {
			// API has no users but state was configured - set empty
			model.Users = setUserVal
		}
		// If state was null and API has no users, leave it null
	}

	// Set all custom properties using helper function
	r.setCustomPropertiesFromResponse(model, roleDetails)

	return nil
}

// RoleResourceDetectDrift detects drift between state and API data and adds errors to block operations
func (r *RolesResource) RoleResourceDetectDrift(roleDetails *openapi.GetRoleDetailsResponse, state *RolesResourceModel, diagnostics *diag.Diagnostics) {
	// Check owner drift
	if !state.Owners.IsNull() && !state.Owners.IsUnknown() {
		var apiOwners []Owner
		ownerUnion := roleDetails.Owner
		if ownerUnion != nil && ownerUnion.ArrayOfGetRoleOwnersResponse != nil {
			for _, apiOwner := range *ownerUnion.ArrayOfGetRoleOwnersResponse {
				if apiOwner.Ownername != nil {
					apiOwners = append(apiOwners, Owner{
						OwnerName: types.StringValue(*apiOwner.Ownername),
						Rank:      types.StringValue(util.SafeDeref(apiOwner.Rank)),
					})
				}
			}
		}

		// Extract state owners
		var stateOwners []Owner
		state.Owners.ElementsAs(context.Background(), &stateOwners, false)

		// Create maps for comparison
		stateOwnerMap := make(map[string]bool)
		for _, owner := range stateOwners {
			key := owner.OwnerName.ValueString() + "|" + owner.Rank.ValueString()
			stateOwnerMap[key] = true
		}

		apiOwnerMap := make(map[string]bool)
		for _, owner := range apiOwners {
			key := owner.OwnerName.ValueString() + "|" + owner.Rank.ValueString()
			apiOwnerMap[key] = true
		}

		// Check for differences
		if len(stateOwnerMap) != len(apiOwnerMap) {
			diagnostics.AddError(
				"Owner Count Drift Detected",
				fmt.Sprintf("State has %d owners but API has %d owners. Please run 'terraform import' to sync the state before proceeding.", len(stateOwners), len(apiOwners)),
			)
		} else {
			// Check if all state owners exist in API
			for stateKey := range stateOwnerMap {
				if !apiOwnerMap[stateKey] {
					diagnostics.AddError(
						"Owner Value Drift Detected",
						fmt.Sprintf("Owner '%s' exists in state but not in API. Please run 'terraform import' to sync the state before proceeding.", stateKey),
					)
					break
				}
			}
		}
	}

	// Check entitlement drift
	if !state.Entitlements.IsNull() && !state.Entitlements.IsUnknown() {
		var apiEntitlements []Entitlement
		if roleDetails.EntitlementDetails != nil {
			for _, apiEnt := range roleDetails.EntitlementDetails {
				if apiEnt.EntitlementValue != nil && apiEnt.EntitlementTypeName != nil && apiEnt.Endpoint != nil {
					apiEntitlements = append(apiEntitlements, Entitlement{
						EntitlementValue: types.StringValue(*apiEnt.EntitlementValue),
						EntitlementType:  types.StringValue(*apiEnt.EntitlementTypeName),
						Endpoint:         types.StringValue(*apiEnt.Endpoint),
					})
				}
			}
		}

		// Extract state entitlements
		var stateEntitlements []Entitlement
		state.Entitlements.ElementsAs(context.Background(), &stateEntitlements, false)

		// Create maps for comparison
		stateEntitlementMap := make(map[string]bool)
		for _, ent := range stateEntitlements {
			key := ent.EntitlementValue.ValueString() + "|" + ent.EntitlementType.ValueString() + "|" + ent.Endpoint.ValueString()
			stateEntitlementMap[key] = true
		}

		apiEntitlementMap := make(map[string]bool)
		for _, ent := range apiEntitlements {
			key := ent.EntitlementValue.ValueString() + "|" + ent.EntitlementType.ValueString() + "|" + ent.Endpoint.ValueString()
			apiEntitlementMap[key] = true
		}

		// Check for differences
		if len(stateEntitlementMap) != len(apiEntitlementMap) {
			diagnostics.AddError(
				"Entitlement Count Drift Detected",
				fmt.Sprintf("State has %d entitlements but API has %d entitlements. Please run 'terraform import' to sync the state before proceeding.", len(stateEntitlements), len(apiEntitlements)),
			)
		} else {
			// Check if all state entitlements exist in API
			for stateKey := range stateEntitlementMap {
				if !apiEntitlementMap[stateKey] {
					diagnostics.AddError(
						"Entitlement Value Drift Detected",
						fmt.Sprintf("Entitlement '%s' exists in state but not in API. Please run 'terraform import' to sync the state before proceeding.", stateKey),
					)
					break
				}
			}
		}
	}

	// Check user drift
	if !state.Users.IsNull() && !state.Users.IsUnknown() {
		var apiUsers []Users
		if roleDetails.UserDetails != nil {
			for _, u := range roleDetails.UserDetails {
				if u.GetUserDetailsResponse != nil && u.GetUserDetailsResponse.Username != nil {
					apiUsers = append(apiUsers, Users{
						UserName: types.StringValue(*u.GetUserDetailsResponse.Username),
					})
				}
			}
		}

		// Extract state users
		var stateUsers []Users
		state.Users.ElementsAs(context.Background(), &stateUsers, false)

		// Create maps for comparison
		stateUserMap := make(map[string]bool)
		for _, user := range stateUsers {
			stateUserMap[user.UserName.ValueString()] = true
		}

		apiUserMap := make(map[string]bool)
		for _, user := range apiUsers {
			apiUserMap[user.UserName.ValueString()] = true
		}

		// Check for differences
		if len(stateUserMap) != len(apiUserMap) {
			diagnostics.AddError(
				"User Count Drift Detected",
				fmt.Sprintf("State has %d users but API has %d users. Please run 'terraform import' to sync the state before proceeding.", len(stateUsers), len(apiUsers)),
			)
		} else {
			// Check if all state users exist in API
			for stateKey := range stateUserMap {
				if !apiUserMap[stateKey] {
					diagnostics.AddError(
						"User Value Drift Detected",
						fmt.Sprintf("User '%s' exists in state but not in API. Please run 'terraform import' to sync the state before proceeding.", stateKey),
					)
					break
				}
			}
		}
	}
}

// CreateRole contains the business logic for creating a role
func (r *RolesResource) CreateRole(ctx context.Context, plan *RolesResourceModel, diagnostics *diag.Diagnostics) (*openapi.CreateEnterpriseRoleResponse, error) {
	roleName := plan.RoleName.ValueString()
	tflog.Debug(ctx, "Starting role creation business logic", map[string]interface{}{"role_name": roleName})

	// Use the factory to create role operations
	roleOps := r.roleFactory.CreateRoleOperations(r.client.APIBaseURL(), r.token)

	// Setting up the owners for the role
	var owners []openapi.CreateRoleOwnerPayload
	var tfOwnerTemplates []Owner
	if !plan.Owners.IsNull() && !plan.Owners.IsUnknown() {
		ownerdiags := plan.Owners.ElementsAs(ctx, &tfOwnerTemplates, true)
		if ownerdiags.HasError() {
			return nil, fmt.Errorf("failed to extract owners from plan")
		}

		if len(plan.Owners.Elements()) == 0 {
			return nil, fmt.Errorf("owners cannot be an empty list. Please provide at least one owner")
		}

		for _, tfOwnerTemplate := range tfOwnerTemplates {
			owner := openapi.CreateRoleOwnerPayload{}
			if tfOwnerTemplate.OwnerName.IsNull() || tfOwnerTemplate.OwnerName.IsUnknown() {
				return nil, fmt.Errorf("owner name cannot be null or unknown. Please provide a valid owner name")
			}
			owner.OwnerName = tfOwnerTemplate.OwnerName.ValueStringPointer()

			if tfOwnerTemplate.Rank.IsNull() || tfOwnerTemplate.Rank.IsUnknown() {
				return nil, fmt.Errorf("owner rank cannot be null or unknown. Please provide a valid owner rank")
			}
			owner.Rank = tfOwnerTemplate.Rank.ValueStringPointer()
			owners = append(owners, owner)
		}
	} else {
		return nil, fmt.Errorf("owners cannot be null or unknown. Please provide a valid list of owners")
	}

	// Build the create request
	createReq := openapi.CreateEnterpriseRoleRequest{
		// Required fields
		RoleName:  plan.RoleName.ValueString(),
		Roletype:  plan.RoleType.ValueString(),
		Requestor: plan.Requestor.ValueString(),
		Owner:     owners,
		// Optional fields
		Endpointname:     util.StringPointerOrEmpty(plan.EndpointName),
		Defaulttimeframe: util.StringPointerOrEmpty(plan.DefaultTimeFrame),
		Description:      util.StringPointerOrEmpty(plan.Description),
		Displayname:      util.StringPointerOrEmpty(plan.DisplayName),
		Glossary:         util.StringPointerOrEmpty(plan.Glossary),
		Risk:             util.StringPointerOrEmpty(plan.Risk),
		Level:            util.StringPointerOrEmpty(plan.Level),
		Soxcritical:      util.StringPointerOrEmpty(plan.SoxCritical),
		Syscritical:      util.StringPointerOrEmpty(plan.SysCritical),
		Priviliged:       util.StringPointerOrEmpty(plan.Privileged),
		Confidentiality:  util.StringPointerOrEmpty(plan.Confidentiality),
		Requestable:      util.StringPointerOrEmpty(plan.Requestable),
		ShowDynamicAttrs: util.StringPointerOrEmpty(plan.ShowDynamicAttrs),
		Checksod:         util.StringPointerOrEmpty(plan.CheckSod),
	}

	// Set all custom properties using helper function
	r.buildCustomPropertiesForCreate(plan, &createReq)

	// Execute the API call
	tflog.Debug(ctx, "Executing role creation API call")
	apiResp, httpResp, err := roleOps.CreateEnterpriseRole(ctx, createReq)

	// Handle errors using helper functions
	var diags diag.Diagnostics
	if rolesutil.RoleHandleHTTPError(ctx, httpResp, err, "creating role", &diags) {
		return nil, fmt.Errorf("failed to create role: %v", diags.Errors())
	}

	if apiResp != nil && rolesutil.RoleHandleAPIError(ctx, apiResp.ErrorCode, apiResp.Message, "creating role", &diags) {
		message := "Unknown error"
		if apiResp.Message != nil {
			message = *apiResp.Message
		}
		return nil, fmt.Errorf("API error during role creation: %s", message)
	}

	// Add entitlements and child roles if present
	hasEntitlements := !plan.Entitlements.IsNull() && !plan.Entitlements.IsUnknown() && len(plan.Entitlements.Elements()) > 0
	hasChildRoles := !plan.ChildRoles.IsNull() && !plan.ChildRoles.IsUnknown() && len(plan.ChildRoles.Elements()) > 0

	if hasEntitlements || hasChildRoles {
		_, updateErr := r.AddEntitlementsAndChildRoles(ctx, plan)
		if updateErr != nil {
			return nil, fmt.Errorf("failed to add entitlements/child roles after creation: %v", updateErr)
		}
	}

	logData := map[string]interface{}{"role_name": roleName}
	if apiResp != nil && apiResp.Requestid != nil && apiResp.Requestkey != nil {
		logData["request_id"] = util.SafeDeref(apiResp.Requestid)
		logData["request_key"] = util.SafeDeref(apiResp.Requestkey)
	}
	tflog.Info(ctx, "Role created successfully", logData)

	return apiResp, nil
}

// AddUsersToRole adds users to role during creation (used in create)
func (r *RolesResource) AddUsersToRole(ctx context.Context, plan *RolesResourceModel, diagnostics *diag.Diagnostics) ([]string, error) {
	roleName := plan.RoleName.ValueString()
	tflog.Debug(ctx, "Adding users to role", map[string]interface{}{"role_name": roleName})

	// Extract users from plan
	var planUsers []Users
	if !plan.Users.IsNull() && !plan.Users.IsUnknown() {
		if err := plan.Users.ElementsAs(ctx, &planUsers, false); err != nil {
			return nil, fmt.Errorf("failed to extract users from plan: %v", err)
		}
	}

	if len(planUsers) == 0 {
		return nil, nil // No users to add
	}

	roleOps := r.roleFactory.CreateRoleOperations(r.client.APIBaseURL(), r.token)

	// Collect errors and success messages
	var errors []string
	var successMessages []string
	var successCount int
	var successfulUsers []string

	// Add each user to the role
	for _, user := range planUsers {
		userName := user.UserName.ValueString()
		apiResp, httpResp, err := roleOps.AddUserToRole(ctx, userName, roleName)

		// Handle HTTP errors
		var diags diag.Diagnostics
		if rolesutil.RoleHandleHTTPError(ctx, httpResp, err, "adding user", &diags) {
			errors = append(errors, fmt.Sprintf("User %s: %v", userName, diags.Errors()))
			continue // Continue with next user
		}

		// Handle API responses - treat errorcode "1" as informational
		if apiResp != nil && apiResp.ErrorCode != nil && *apiResp.ErrorCode != "0" {
			message := "Unknown error"
			if apiResp.Message != nil {
				message = *apiResp.Message
			}

			if *apiResp.ErrorCode == "1" {
				errors = append(errors, fmt.Sprintf("User %s (Info): %s", userName, message))
				continue
			} else {
				// Actual error - collect error but continue
				errors = append(errors, fmt.Sprintf("User %s: %s", userName, message))
				continue
			}
		} else {
			// Success case (ErrorCode is "0" or nil)
			successCount++
			successfulUsers = append(successfulUsers, userName)

			// Collect success message with task IDs if available
			if apiResp != nil && apiResp.Message != nil && *apiResp.Message != "" {
				successMessages = append(successMessages, fmt.Sprintf("User %s: %s", userName, *apiResp.Message))
				tflog.Info(ctx, "User added to role successfully", map[string]interface{}{
					"user_name": userName,
					"role_name": roleName,
					"message":   *apiResp.Message,
				})
			}
		}

		tflog.Debug(ctx, "Added user to role", map[string]interface{}{
			"user_name": userName,
			"role_name": roleName,
		})
	}

	// Log summary
	tflog.Info(ctx, "User addition summary", map[string]interface{}{
		"role_name":   roleName,
		"total_users": len(planUsers),
		"successful":  successCount,
		"failed":      len(errors),
	})

	// Log success messages with task IDs for user visibility
	if len(successMessages) > 0 {
		tflog.Info(ctx, "User addition task details", map[string]interface{}{
			"role_name":        roleName,
			"success_messages": strings.Join(successMessages, "; "),
		})
	}

	// Add success warnings to show task IDs to users
	if len(successMessages) > 0 && diagnostics != nil {
		diagnostics.AddWarning(
			"User Addition Task Details - Manual UI Action Required",
			fmt.Sprintf("Users successfully added to role %s. Task details: %s\n\n"+
				"⚠️  IMPORTANT: These operations created tasks in Saviynt that require manual completion.\n"+
				"Please log into the Saviynt UI and navigate to the Pending Tasks section to complete these tasks.\n"+
				"Tasks will remain pending until manually approved/completed in the UI.",
				roleName, strings.Join(successMessages, "; ")),
		)
	}

	// Return combined error if any failures
	if len(errors) > 0 {
		return successfulUsers, fmt.Errorf("failed to add %d/%d users to role %s: %s",
			len(errors), len(planUsers), roleName, strings.Join(errors, "; "))
	}

	return successfulUsers, nil
}

// ReadRole contains the business logic for reading a role
func (r *RolesResource) ReadRole(ctx context.Context, roleName string) (*openapi.GetRolesResponse, error) {
	tflog.Debug(ctx, "Starting role read business logic", map[string]interface{}{"role_name": roleName})

	// Use the factory to create role operations
	roleOps := r.roleFactory.CreateRoleOperations(r.client.APIBaseURL(), r.token)

	reqParams := openapi.GetRolesRequest{}
	reqParams.SetRoleName(roleName)
	reqParams.SetRequestedObject("entitlements")

	tflog.Debug(ctx, "Executing role read API call")
	apiResp, httpResp, err := roleOps.GetRoles(ctx, reqParams)

	// Handle errors using helper functions
	var diags diag.Diagnostics
	if rolesutil.RoleHandleHTTPError(ctx, httpResp, err, "reading role", &diags) {
		return nil, fmt.Errorf("failed to read role: %v", diags.Errors())
	}

	if err != nil {
		tflog.Error(ctx, "Problem with the get function in read block", map[string]interface{}{"error": err.Error()})
		return nil, fmt.Errorf("error reading role: %v", err)
	}

	// Check if we got valid role details
	if apiResp.Roledetails == nil || len(apiResp.Roledetails) == 0 {
		return nil, fmt.Errorf("API returned empty role details list")
	}

	tflog.Info(ctx, "Role read successfully", map[string]interface{}{
		"role_name": roleName,
	})

	return apiResp, nil
}

// UpdateRole contains the business logic for updating a role
func (r *RolesResource) UpdateRole(ctx context.Context, plan *RolesResourceModel, state *RolesResourceModel, diagnostics *diag.Diagnostics) (*openapi.UpdateEnterpriseRoleResponse, error) {
	roleName := plan.RoleName.ValueString()
	tflog.Debug(ctx, "Starting role update business logic", map[string]interface{}{"role_name": roleName})

	// Use the factory to create role operations
	roleOps := r.roleFactory.CreateRoleOperations(r.client.APIBaseURL(), r.token)

	// Process entitlements, owners, and child roles changes
	entitlements, err := r.RoleResourceProcessEntitlementChanges(ctx, plan, state)
	if err != nil {
		return nil, fmt.Errorf("failed to process entitlement changes: %v", err)
	}

	owners, err := r.RoleResourceProcessOwnerChanges(ctx, plan, state)
	if err != nil {
		return nil, fmt.Errorf("failed to process owner changes: %v", err)
	}

	childRoles, err := r.processChildRoleChanges(ctx, plan, state)
	if err != nil {
		return nil, fmt.Errorf("failed to process child role changes: %v", err)
	}

	// Build the update request
	updateReq := openapi.UpdateEnterpriseRoleRequest{
		// Required fields
		RoleName: plan.RoleName.ValueString(),
		Roletype: plan.RoleType.ValueString(),
		// Optional fields
		Requestor:        util.StringPointerOrEmpty(plan.Requestor),
		Owner:            owners,
		Entitlements:     entitlements,
		ChildRoles:       childRoles,
		Endpointname:     util.StringPointerOrEmpty(plan.EndpointName),
		Defaulttimeframe: util.StringPointerOrEmpty(plan.DefaultTimeFrame),
		Description:      util.StringPointerOrEmpty(plan.Description),
		Displayname:      util.StringPointerOrEmpty(plan.DisplayName),
		Glossary:         util.StringPointerOrEmpty(plan.Glossary),
		Risk:             util.StringPointerOrEmpty(plan.Risk),
		Level:            util.StringPointerOrEmpty(plan.Level),
		Soxcritical:      util.StringPointerOrEmpty(plan.SoxCritical),
		Syscritical:      util.StringPointerOrEmpty(plan.SysCritical),
		Priviliged:       util.StringPointerOrEmpty(plan.Privileged),
		Confidentiality:  util.StringPointerOrEmpty(plan.Confidentiality),
		Requestable:      util.StringPointerOrEmpty(plan.Requestable),
		ShowDynamicAttrs: util.StringPointerOrEmpty(plan.ShowDynamicAttrs),
	}

	// Guard rails: Only include arrays if they have items
	if len(owners) > 0 {
		updateReq.Owner = owners
	}
	if len(entitlements) > 0 {
		updateReq.Entitlements = entitlements
	}
	if len(childRoles) > 0 {
		updateReq.ChildRoles = childRoles
	}

	// Set all custom properties using helper function
	r.buildCustomPropertiesForUpdate(plan, &updateReq)

	// Execute the update API call
	tflog.Debug(ctx, "Executing role update API call")
	apiResp, httpResp, err := roleOps.UpdateEnterpriseRole(ctx, updateReq)

	// Handle errors using helper functions
	var diags diag.Diagnostics
	if rolesutil.RoleHandleHTTPError(ctx, httpResp, err, "updating role", &diags) {
		return nil, fmt.Errorf("failed to update role: %v", diags.Errors())
	}

	if apiResp != nil && rolesutil.RoleHandleAPIError(ctx, apiResp.ErrorCode, apiResp.Message, "updating role", &diags) {
		message := "Unknown error"
		if apiResp.Message != nil {
			message = *apiResp.Message
		}
		return nil, fmt.Errorf("API error during role update: %s", message)
	}

	logData := map[string]interface{}{"role_name": roleName}
	if apiResp != nil && apiResp.Requestid != nil && apiResp.Requestkey != nil {
		logData["request_id"] = util.SafeDeref(apiResp.Requestid)
		logData["request_key"] = util.SafeDeref(apiResp.Requestkey)
	}
	tflog.Info(ctx, "Role updated successfully", logData)

	// Process user changes using specific APIs
	_, err = r.RoleResourceProcessUserChanges(ctx, plan, state, diagnostics)
	if err != nil {
		return nil, fmt.Errorf("failed to process user changes: %v", err)
	}
	return apiResp, nil
}

// AddEntitlementsAndChildRoles adds entitlements and child roles directly from plan (used in create)
func (r *RolesResource) AddEntitlementsAndChildRoles(ctx context.Context, plan *RolesResourceModel) (*openapi.UpdateEnterpriseRoleResponse, error) {
	roleName := plan.RoleName.ValueString()
	tflog.Debug(ctx, "Adding entitlements and child roles", map[string]interface{}{"role_name": roleName})

	// Use the factory to create role operations
	roleOps := r.roleFactory.CreateRoleOperations(r.client.APIBaseURL(), r.token)

	// Create entitlement payloads directly from plan
	var entitlements []openapi.UpdateEntitlementPayload
	if !plan.Entitlements.IsNull() && !plan.Entitlements.IsUnknown() {
		var tfEntitlements []Entitlement
		if err := plan.Entitlements.ElementsAs(ctx, &tfEntitlements, false); err != nil {
			return nil, fmt.Errorf("failed to extract entitlements: %v", err)
		}
		for _, e := range tfEntitlements {
			addType := UpdateTypeAdd
			entitlements = append(entitlements, openapi.UpdateEntitlementPayload{
				EntitlementValue: e.EntitlementValue.ValueStringPointer(),
				EntitlementType:  e.EntitlementType.ValueStringPointer(),
				Endpoint:         e.Endpoint.ValueStringPointer(),
				UpdateType:       &addType,
			})
		}
	}

	// Create child role payloads directly from plan
	var childRoles []openapi.UpdateChildRolePayload
	if !plan.ChildRoles.IsNull() && !plan.ChildRoles.IsUnknown() {
		var tfChildRoles []ChildRoles
		if err := plan.ChildRoles.ElementsAs(ctx, &tfChildRoles, false); err != nil {
			return nil, fmt.Errorf("failed to extract child roles: %v", err)
		}
		for _, cr := range tfChildRoles {
			addType := UpdateTypeAdd
			childRoles = append(childRoles, openapi.UpdateChildRolePayload{
				Rolename:   cr.RoleName.ValueStringPointer(),
				UpdateType: &addType,
			})
		}
	}

	// Build update request with only required fields
	updateReq := openapi.UpdateEnterpriseRoleRequest{
		RoleName:     plan.RoleName.ValueString(),
		Roletype:     plan.RoleType.ValueString(),
		Entitlements: entitlements,
		ChildRoles:   childRoles,
	}

	// Execute update API call
	tflog.Debug(ctx, "Executing entitlements and child roles add API call")
	apiResp, httpResp, err := roleOps.UpdateEnterpriseRole(ctx, updateReq)

	// Handle errors
	var diags diag.Diagnostics
	if rolesutil.RoleHandleHTTPError(ctx, httpResp, err, "adding entitlements/child roles", &diags) {
		return nil, fmt.Errorf("failed to add entitlements/child roles: %v", diags.Errors())
	}

	if apiResp != nil && rolesutil.RoleHandleAPIError(ctx, apiResp.ErrorCode, apiResp.Message, "adding entitlements/child roles", &diags) {
		message := "Unknown error"
		if apiResp.Message != nil {
			message = *apiResp.Message
		}
		return nil, fmt.Errorf("API error during entitlements/child roles add: %s", message)
	}

	tflog.Info(ctx, "Entitlements and child roles added successfully", map[string]interface{}{"role_name": roleName})
	return apiResp, nil
}

// RoleResourceProcessEntitlementChanges processes entitlement changes for update operations using the generic processor
func (r *RolesResource) RoleResourceProcessEntitlementChanges(ctx context.Context, plan *RolesResourceModel, state *RolesResourceModel) ([]openapi.UpdateEntitlementPayload, error) {
	processor := &ChangeProcessor[Entitlement, string, openapi.UpdateEntitlementPayload]{
		KeyExtractor: func(e Entitlement) string {
			return e.EntitlementValue.ValueString() + "|" + e.EntitlementType.ValueString() + "|" + e.Endpoint.ValueString()
		},
		PayloadBuilder: func(e Entitlement, updateType string) openapi.UpdateEntitlementPayload {
			return openapi.UpdateEntitlementPayload{
				EntitlementValue: e.EntitlementValue.ValueStringPointer(),
				EntitlementType:  e.EntitlementType.ValueStringPointer(),
				Endpoint:         e.Endpoint.ValueStringPointer(),
				UpdateType:       &updateType,
			}
		},
		ErrorContext: "entitlements",
	}

	return processor.ProcessChanges(ctx, state.Entitlements, plan.Entitlements)
}

// RoleResourceProcessOwnerChanges processes owner changes for update operations using the generic processor
func (r *RolesResource) RoleResourceProcessOwnerChanges(ctx context.Context, plan *RolesResourceModel, state *RolesResourceModel) ([]openapi.UpdateRoleOwnerPayload, error) {
	processor := &ChangeProcessor[Owner, string, openapi.UpdateRoleOwnerPayload]{
		KeyExtractor: func(o Owner) string {
			return o.OwnerName.ValueString() + "|" + o.Rank.ValueString()
		},
		PayloadBuilder: func(o Owner, updateType string) openapi.UpdateRoleOwnerPayload {
			return openapi.UpdateRoleOwnerPayload{
				OwnerName:  o.OwnerName.ValueStringPointer(),
				Rank:       o.Rank.ValueStringPointer(),
				UpdateType: &updateType,
			}
		},
		ErrorContext: "owners",
	}

	return processor.ProcessChanges(ctx, state.Owners, plan.Owners)
}

// processChildRoleChanges processes child role changes for update operations using the generic processor
func (r *RolesResource) processChildRoleChanges(ctx context.Context, plan *RolesResourceModel, state *RolesResourceModel) ([]openapi.UpdateChildRolePayload, error) {
	processor := &ChangeProcessor[ChildRoles, string, openapi.UpdateChildRolePayload]{
		KeyExtractor: func(cr ChildRoles) string {
			return cr.RoleName.ValueString()
		},
		PayloadBuilder: func(cr ChildRoles, updateType string) openapi.UpdateChildRolePayload {
			return openapi.UpdateChildRolePayload{
				Rolename:   cr.RoleName.ValueStringPointer(),
				UpdateType: &updateType,
			}
		},
		ErrorContext: "child roles",
	}

	return processor.ProcessChanges(ctx, state.ChildRoles, plan.ChildRoles)
}

// RoleResourceProcessUserChanges processes user changes by calling specific AddUserToRole/RemoveUserFromRole APIs
func (r *RolesResource) RoleResourceProcessUserChanges(ctx context.Context, plan *RolesResourceModel, state *RolesResourceModel, diagnostics *diag.Diagnostics) (*openapi.AddOrRemoveRoleResponse, error) {
	roleName := plan.RoleName.ValueString()

	// Variable to capture the last API response for message/errorcode
	var lastApiResp *openapi.AddOrRemoveRoleResponse

	// Extract state users
	var stateUsers []Users
	if !state.Users.IsNull() && !state.Users.IsUnknown() {
		if err := state.Users.ElementsAs(ctx, &stateUsers, false); err != nil {
			return nil, fmt.Errorf("failed to extract state users: %v", err)
		}
	}

	// Extract plan users
	var planUsers []Users
	if !plan.Users.IsNull() && !plan.Users.IsUnknown() {
		if err := plan.Users.ElementsAs(ctx, &planUsers, false); err != nil {
			return nil, fmt.Errorf("failed to extract plan users: %v", err)
		}
	}

	// Create maps for comparison (key = user_name)
	stateMap := make(map[string]bool)
	for _, u := range stateUsers {
		stateMap[u.UserName.ValueString()] = true
	}

	planMap := make(map[string]bool)
	for _, u := range planUsers {
		planMap[u.UserName.ValueString()] = true
	}

	roleOps := r.roleFactory.CreateRoleOperations(r.client.APIBaseURL(), r.token)

	// Collect errors and success messages
	var errors []string
	var successMessages []string
	var removeCount, addCount int

	// Remove users (in state but not in plan)
	for userName := range stateMap {
		if !planMap[userName] {
			apiResp, httpResp, err := roleOps.RemoveUserFromRole(ctx, userName, roleName)
			lastApiResp = apiResp // Capture for message/errorcode

			// Handle HTTP errors
			var diags diag.Diagnostics
			if rolesutil.RoleHandleHTTPError(ctx, httpResp, err, "removing user", &diags) {
				errors = append(errors, fmt.Sprintf("Remove user %s: %v", userName, diags.Errors()))
				continue
			}

			// Handle API responses - treat errorcode "1" as informational
			if apiResp != nil && apiResp.ErrorCode != nil && *apiResp.ErrorCode != "0" {
				message := "Unknown error"
				if apiResp.Message != nil {
					message = *apiResp.Message
				}

				if *apiResp.ErrorCode == "1" {
					errors = append(errors, fmt.Sprintf("Remove user %s (Info): %s", userName, message))
					continue
				} else {
					// Actual error - collect error but continue
					errors = append(errors, fmt.Sprintf("Remove user %s: %s", userName, message))
					continue
				}
			} else {
				// Success case
				removeCount++
				if apiResp != nil && apiResp.Message != nil && *apiResp.Message != "" {
					successMessages = append(successMessages, fmt.Sprintf("Removed user %s: %s", userName, *apiResp.Message))
				}
			}
		}
	}

	// Add users (in plan but not in state)
	for userName := range planMap {
		if !stateMap[userName] {
			apiResp, httpResp, err := roleOps.AddUserToRole(ctx, userName, roleName)
			lastApiResp = apiResp // Capture for message/errorcode

			// Handle HTTP errors
			var diags diag.Diagnostics
			if rolesutil.RoleHandleHTTPError(ctx, httpResp, err, "adding user", &diags) {
				errors = append(errors, fmt.Sprintf("Add user %s: %v", userName, diags.Errors()))
				continue
			}

			// Handle API responses - treat errorcode "1" as informational
			if apiResp != nil && apiResp.ErrorCode != nil && *apiResp.ErrorCode != "0" {
				message := "Unknown error"
				if apiResp.Message != nil {
					message = *apiResp.Message
				}

				if *apiResp.ErrorCode == "1" {
					errors = append(errors, fmt.Sprintf("Add user %s (Info): %s", userName, message))
					continue
				} else {
					// Actual error - collect error but continue
					errors = append(errors, fmt.Sprintf("Add user %s: %s", userName, message))
					continue
				}
			} else {
				// Success case
				addCount++
				if apiResp != nil && apiResp.Message != nil && *apiResp.Message != "" {
					successMessages = append(successMessages, fmt.Sprintf("Added user %s: %s", userName, *apiResp.Message))
				}
			}
		}
	}

	// Log summary
	tflog.Info(ctx, "User changes summary", map[string]interface{}{
		"role_name":     roleName,
		"users_added":   addCount,
		"users_removed": removeCount,
		"failed":        len(errors),
	})

	// Add success warnings to show task IDs to users
	if len(successMessages) > 0 && diagnostics != nil {
		diagnostics.AddWarning(
			"User Changes Task Details - Manual UI Action Required",
			fmt.Sprintf("User changes completed for role %s. Task details: %s\n\n"+
				"⚠️  IMPORTANT: These operations created tasks in Saviynt that require manual completion.\n"+
				"Please log into the Saviynt UI and navigate to the Pending Tasks section to complete these tasks.\n"+
				"Tasks will remain pending until manually approved/completed in the UI.",
				roleName, strings.Join(successMessages, "; ")),
		)
	}

	// Return combined error if any failures
	if len(errors) > 0 {
		return nil, fmt.Errorf("failed user changes for role %s: %s", roleName, strings.Join(errors, "; "))
	}

	return lastApiResp, nil
}

// RoleResourceHasRealNonUserChanges checks if there are meaningful non-user field changes (ignoring computed fields)
func (r *RolesResource) RoleResourceHasRealNonUserChanges(ctx context.Context, plan, state *RolesResourceModel) bool {
	// Helper to check if a field has a real change (not null/unknown -> computed)
	realChange := func(planVal, stateVal attr.Value) bool {
		if planVal.Equal(stateVal) {
			return false
		}
		// If state is null/unknown but plan has value, likely computed - ignore
		if (stateVal.IsNull() || stateVal.IsUnknown()) && !planVal.IsNull() && !planVal.IsUnknown() {
			return false
		}
		// If plan is unknown (computed), ignore regardless of state value
		if planVal.IsUnknown() {
			return false
		}
		return true
	}

	// Check ALL non-user fields
	return realChange(plan.RoleType, state.RoleType) ||
		realChange(plan.Description, state.Description) ||
		realChange(plan.DisplayName, state.DisplayName) ||
		realChange(plan.Requestor, state.Requestor) ||
		realChange(plan.Owners, state.Owners) ||
		realChange(plan.Entitlements, state.Entitlements) ||
		realChange(plan.ChildRoles, state.ChildRoles) ||
		realChange(plan.EndpointName, state.EndpointName) ||
		realChange(plan.DefaultTimeFrame, state.DefaultTimeFrame) ||
		realChange(plan.Glossary, state.Glossary) ||
		realChange(plan.Risk, state.Risk) ||
		realChange(plan.Level, state.Level) ||
		realChange(plan.SoxCritical, state.SoxCritical) ||
		realChange(plan.SysCritical, state.SysCritical) ||
		realChange(plan.Privileged, state.Privileged) ||
		realChange(plan.Confidentiality, state.Confidentiality) ||
		realChange(plan.Requestable, state.Requestable) ||
		realChange(plan.ShowDynamicAttrs, state.ShowDynamicAttrs) ||
		realChange(plan.CheckSod, state.CheckSod) ||
		// Computed fields that should be ignored
		realChange(plan.ID, state.ID) ||
		// Custom properties 1-60
		realChange(plan.CustomProperty1, state.CustomProperty1) ||
		realChange(plan.CustomProperty2, state.CustomProperty2) ||
		realChange(plan.CustomProperty3, state.CustomProperty3) ||
		realChange(plan.CustomProperty4, state.CustomProperty4) ||
		realChange(plan.CustomProperty5, state.CustomProperty5) ||
		realChange(plan.CustomProperty6, state.CustomProperty6) ||
		realChange(plan.CustomProperty7, state.CustomProperty7) ||
		realChange(plan.CustomProperty8, state.CustomProperty8) ||
		realChange(plan.CustomProperty9, state.CustomProperty9) ||
		realChange(plan.CustomProperty10, state.CustomProperty10) ||
		realChange(plan.CustomProperty11, state.CustomProperty11) ||
		realChange(plan.CustomProperty12, state.CustomProperty12) ||
		realChange(plan.CustomProperty13, state.CustomProperty13) ||
		realChange(plan.CustomProperty14, state.CustomProperty14) ||
		realChange(plan.CustomProperty15, state.CustomProperty15) ||
		realChange(plan.CustomProperty16, state.CustomProperty16) ||
		realChange(plan.CustomProperty17, state.CustomProperty17) ||
		realChange(plan.CustomProperty18, state.CustomProperty18) ||
		realChange(plan.CustomProperty19, state.CustomProperty19) ||
		realChange(plan.CustomProperty20, state.CustomProperty20) ||
		realChange(plan.CustomProperty21, state.CustomProperty21) ||
		realChange(plan.CustomProperty22, state.CustomProperty22) ||
		realChange(plan.CustomProperty23, state.CustomProperty23) ||
		realChange(plan.CustomProperty24, state.CustomProperty24) ||
		realChange(plan.CustomProperty25, state.CustomProperty25) ||
		realChange(plan.CustomProperty26, state.CustomProperty26) ||
		realChange(plan.CustomProperty27, state.CustomProperty27) ||
		realChange(plan.CustomProperty28, state.CustomProperty28) ||
		realChange(plan.CustomProperty29, state.CustomProperty29) ||
		realChange(plan.CustomProperty30, state.CustomProperty30) ||
		realChange(plan.CustomProperty31, state.CustomProperty31) ||
		realChange(plan.CustomProperty32, state.CustomProperty32) ||
		realChange(plan.CustomProperty33, state.CustomProperty33) ||
		realChange(plan.CustomProperty34, state.CustomProperty34) ||
		realChange(plan.CustomProperty35, state.CustomProperty35) ||
		realChange(plan.CustomProperty36, state.CustomProperty36) ||
		realChange(plan.CustomProperty37, state.CustomProperty37) ||
		realChange(plan.CustomProperty38, state.CustomProperty38) ||
		realChange(plan.CustomProperty39, state.CustomProperty39) ||
		realChange(plan.CustomProperty40, state.CustomProperty40) ||
		realChange(plan.CustomProperty41, state.CustomProperty41) ||
		realChange(plan.CustomProperty42, state.CustomProperty42) ||
		realChange(plan.CustomProperty43, state.CustomProperty43) ||
		realChange(plan.CustomProperty44, state.CustomProperty44) ||
		realChange(plan.CustomProperty45, state.CustomProperty45) ||
		realChange(plan.CustomProperty46, state.CustomProperty46) ||
		realChange(plan.CustomProperty47, state.CustomProperty47) ||
		realChange(plan.CustomProperty48, state.CustomProperty48) ||
		realChange(plan.CustomProperty49, state.CustomProperty49) ||
		realChange(plan.CustomProperty50, state.CustomProperty50) ||
		realChange(plan.CustomProperty51, state.CustomProperty51) ||
		realChange(plan.CustomProperty52, state.CustomProperty52) ||
		realChange(plan.CustomProperty53, state.CustomProperty53) ||
		realChange(plan.CustomProperty54, state.CustomProperty54) ||
		realChange(plan.CustomProperty55, state.CustomProperty55) ||
		realChange(plan.CustomProperty56, state.CustomProperty56) ||
		realChange(plan.CustomProperty57, state.CustomProperty57) ||
		realChange(plan.CustomProperty58, state.CustomProperty58) ||
		realChange(plan.CustomProperty59, state.CustomProperty59) ||
		realChange(plan.CustomProperty60, state.CustomProperty60)
}

// ReadRoleStateFromAPI reads current role state from API and updates the model (post-update sync)
func (r *RolesResource) ReadRoleStateFromAPI(ctx context.Context, plan *RolesResourceModel, diagnostics *diag.Diagnostics) error {
	// Read current role from API
	apiResp, err := r.ReadRole(ctx, plan.RoleName.ValueString())
	if err != nil {
		return fmt.Errorf("failed to read role from API: %w", err)
	}

	if len(apiResp.Roledetails) == 0 {
		return fmt.Errorf("no role details returned from API")
	}

	roleDetails := apiResp.Roledetails[0]

	// Use mapRoleDetailsToModel in IMPORT MODE to sync with API (like entitlements do)
	return r.RoleResourceMapRoleDetailsToModel(plan, &roleDetails, plan, diagnostics, true)
}

// updateModelFromCreateResponse updates the Terraform model from create API response
func (r *RolesResource) updateModelFromCreateResponse(plan *RolesResourceModel) {
	// Set the ID from the role name
	plan.ID = types.StringValue("roles-" + plan.RoleName.ValueString())

	// Set the plan with the request data.
	plan.RoleName = types.StringValue(plan.RoleName.ValueString())
	plan.RoleType = types.StringValue(plan.RoleType.ValueString())
	plan.Requestor = types.StringValue(plan.Requestor.ValueString())
	plan.EndpointName = types.StringValue(util.SafeDeref(plan.EndpointName.ValueStringPointer()))
	plan.DefaultTimeFrame = types.StringValue(util.SafeDeref(plan.DefaultTimeFrame.ValueStringPointer()))
	plan.Description = util.SafeString(plan.Description.ValueStringPointer())
	plan.DisplayName = types.StringValue(util.SafeDeref(plan.DisplayName.ValueStringPointer()))
	plan.Glossary = util.SafeString(plan.Glossary.ValueStringPointer())
	plan.SoxCritical = util.SafeString(plan.SoxCritical.ValueStringPointer())
	plan.SysCritical = util.SafeString(plan.SysCritical.ValueStringPointer())
	plan.Privileged = util.SafeString(plan.Privileged.ValueStringPointer())
	plan.Confidentiality = util.SafeString(plan.Confidentiality.ValueStringPointer())
	plan.Requestable = util.SafeString(plan.Requestable.ValueStringPointer())
	plan.ShowDynamicAttrs = util.SafeString(plan.ShowDynamicAttrs.ValueStringPointer())
	plan.Risk = util.SafeString(plan.Risk.ValueStringPointer())
	// Set all custom properties using helper function
	r.setCustomPropertiesInPlan(plan)
}

// RoleResourceFetchEndpointName fetches endpoint name using endpoint key
func (r *RolesResource) RoleResourceFetchEndpointName(ctx context.Context, endpointKey *string, diagnostics *diag.Diagnostics) types.String {
	if endpointKey == nil || *endpointKey == "" {
		return types.StringNull()
	}

	// Use the factory to create endpoint operations
	endpointOps := r.roleFactory.CreateEndpointOperations(r.client.APIBaseURL(), r.token)

	endpointReq := endpoint.GetEndpointsRequest{}
	endpointKeys := []string{*endpointKey}
	endpointReq.Endpointkey = endpointKeys

	tflog.Debug(ctx, "Fetching endpoint details", map[string]interface{}{"endpoint_key": *endpointKey})
	endpointResp, httpRespEndpoint, err := endpointOps.GetEndpoints(ctx, endpointReq)

	// Handle HTTP errors for endpoint operations using helper function
	if rolesutil.RoleHandleHTTPError(ctx, httpRespEndpoint, err, "reading endpoint", diagnostics) {
		return types.StringNull()
	}

	if len(endpointResp.Endpoints) > 0 {
		return types.StringValue(*endpointResp.Endpoints[0].Endpointname)
	} else {
		tflog.Debug(ctx, "No endpoint found for the role")
		return types.StringNull()
	}
}

// Create implements the resource.Resource interface for creating a new role in Saviynt.
func (r *RolesResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan RolesResourceModel

	tflog.Debug(ctx, "Starting role creation")

	// Extract plan from request
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	roleName := plan.RoleName.ValueString()
	tflog.Debug(ctx, "Creating role", map[string]interface{}{"role_name": roleName})

	// Use business logic method
	apiResp, err := r.CreateRole(ctx, &plan, &resp.Diagnostics)
	if err != nil {
		resp.Diagnostics.AddError(
			"Role Creation Failed",
			err.Error(),
		)
		return
	}
	if apiResp != nil && apiResp.Requestid != nil && *apiResp.Requestid == "" && apiResp.Requestkey != nil && *apiResp.Requestkey == "" {
		resp.Diagnostics.AddWarning(
			"Role Created - Manual UI Action Required",
			fmt.Sprintf(
				"Role created successfully.\nMessage: %s\nErrorCode: %s\nRequestID: %s\nRequestKey: %s\n\n"+
					"⚠️  IMPORTANT: Role creation has generated a pending request that requires manual approval.\n"+
					"Please log into the Saviynt UI and navigate to the Pending Requests page to approve this role creation.\n"+
					"The role will not be fully active until the request is approved in the UI.",
				*apiResp.Message,
				*apiResp.ErrorCode,
				*apiResp.Requestid,
				*apiResp.Requestkey,
			),
		)
	} else if apiResp != nil && apiResp.Message != nil && *apiResp.Message != "" && apiResp.ErrorCode != nil && *apiResp.ErrorCode != "" {
		resp.Diagnostics.AddWarning(
			"Role Created - Manual UI Action Required",
			fmt.Sprintf(
				"Role created successfully.\nMessage: %s\nErrorCode: %s\n\n"+
					"⚠️  IMPORTANT: Role creation has generated a pending request that requires manual approval.\n"+
					"Please log into the Saviynt UI and navigate to the Pending Requests page to approve this role creation.\n"+
					"The role will not be fully active until the request is approved in the UI.",
				*apiResp.Message,
				*apiResp.ErrorCode,
			),
		)
	} else {
		resp.Diagnostics.AddWarning(
			"Info",
			"Provider error: received unexpected response from Saviynt API.Please retry or contact support.",
		)
	}
	// Update model from create response
	r.updateModelFromCreateResponse(&plan)

	// Add users to role if present
	var userErr error
	var successfulUsers []string
	hasUsers := !plan.Users.IsNull() && !plan.Users.IsUnknown() && len(plan.Users.Elements()) > 0
	if hasUsers {
		successfulUsers, userErr = r.AddUsersToRole(ctx, &plan, &resp.Diagnostics)
		if userErr != nil {
			resp.Diagnostics.AddWarning(
				"User Operation Failed During Role Creation",
				fmt.Sprintf("Role '%s' was created successfully, but some users failed to be added: %s\n\n"+
					"The role exists and successful users have been added. "+
					"Please fix the failed users and run terraform apply again to add the remaining users.",
					roleName, userErr.Error()),
			)

			// Update plan to include only successful users, preserving null vs empty distinction
			if len(successfulUsers) > 0 {
				successUserObjects := make([]attr.Value, len(successfulUsers))
				for i, userName := range successfulUsers {
					successUserObjects[i] = types.ObjectValueMust(
						map[string]attr.Type{"user_name": types.StringType},
						map[string]attr.Value{"user_name": types.StringValue(userName)},
					)
				}
				plan.Users = types.SetValueMust(types.ObjectType{
					AttrTypes: map[string]attr.Type{"user_name": types.StringType},
				}, successUserObjects)
			} else {
				// No successful users - keep as empty array since original was not null
				plan.Users = types.SetValueMust(types.ObjectType{
					AttrTypes: map[string]attr.Type{"user_name": types.StringType},
				}, []attr.Value{})
			}
		}
	}

	// Set the state
	resp.Diagnostics.Append(resp.State.Set(ctx, plan)...)
	if userErr != nil {
		resp.Diagnostics.AddError(
			"User Operation Failed During Role Creation",
			fmt.Sprintf("Role '%s' was created successfully, but some users failed to be added: %s\n\n"+
				"The role exists and successful users have been added. "+
				"Please fix the failed users and run terraform apply again to add the remaining users.",
				roleName, userErr.Error()),
		)
	}

}

// Read implements the resource.Resource interface for reading a role from Saviynt.
func (r *RolesResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state RolesResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	roleName := state.RoleName.ValueString()
	tflog.Debug(ctx, "Starting role read", map[string]interface{}{"role_name": roleName})

	// Use business logic method
	apiResp, err := r.ReadRole(ctx, roleName)
	if err != nil {
		resp.Diagnostics.AddError("API Read Failed", fmt.Sprintf("Error: %v", err))
		return
	}

	// Fetch endpoint name if available
	if len(apiResp.Roledetails) > 0 && apiResp.Roledetails[0].Endpointkey != nil {
		state.EndpointName = r.RoleResourceFetchEndpointName(ctx, apiResp.Roledetails[0].Endpointkey, &resp.Diagnostics)
		if resp.Diagnostics.HasError() {
			return
		}
	}

	// Safe array access with bounds checking
	if len(apiResp.Roledetails) == 0 {
		resp.Diagnostics.AddError("API Error", "No role details returned from API")
		return
	}
	roleDetails := apiResp.Roledetails[0]
	resp.Diagnostics.AddWarning(
		"Info",
		fmt.Sprintf(
			"Role read successfully.\nMessage: %s\nErrorCode: %s\n",
			*apiResp.Msg,
			*apiResp.ErrorCode,
		),
	)

	// Map role details to model using helper function
	if err := r.RoleResourceMapRoleDetailsToModel(&state, &roleDetails, &state, &resp.Diagnostics); err != nil {
		resp.Diagnostics.AddError("Mapping Error", err.Error())
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

// Update implements the resource.Resource interface for updating an existing role in Saviynt.
func (r *RolesResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan RolesResourceModel
	var state RolesResourceModel

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

	roleName := plan.RoleName.ValueString()
	tflog.Debug(ctx, "Starting role update", map[string]interface{}{"role_name": roleName})

	// Validate that the role name has not changed
	if plan.RoleName.ValueString() != state.RoleName.ValueString() {
		resp.Diagnostics.AddError("Role Update Failed", "role name cannot be updated")
		return
	}

	// Detect what changed - use new simplified logic
	hasUserChanges := !plan.Users.Equal(state.Users)
	hasRealNonUserChanges := r.RoleResourceHasRealNonUserChanges(ctx, &plan, &state)

	// Declare apiResp at function level to be available for later use
	var apiResp *openapi.UpdateEnterpriseRoleResponse
	var err error
	var userErr error
	if hasUserChanges && !hasRealNonUserChanges {
		// Scenario 1: Only users changed - skip role API, call user API only
		tflog.Debug(ctx, "Only user fields changed, calling user APIs only")
		_, userErr = r.RoleResourceProcessUserChanges(ctx, &plan, &state, &resp.Diagnostics)
		if userErr != nil {
			// Use WARNING for user failures - these are non-critical
			resp.Diagnostics.AddWarning("User Operation Failed", userErr.Error())
		}
	} else if hasRealNonUserChanges {
		// Scenario 2: Non-user fields changed - call role API
		tflog.Debug(ctx, "Non-user fields changed, calling main role update API")
		apiResp, err = r.UpdateRole(ctx, &plan, &state, &resp.Diagnostics)
		if err != nil {
			resp.Diagnostics.AddError("Role Update Failed", err.Error())
			return
		}

		// Also call user API if users changed
		if hasUserChanges {
			tflog.Debug(ctx, "Both role and user fields changed, also processing user changes")
			_, userErr = r.RoleResourceProcessUserChanges(ctx, &plan, &state, &resp.Diagnostics)
			if userErr != nil {
				// Use WARNING for user failures - these are non-critical
				resp.Diagnostics.AddWarning("User Operation Failed", userErr.Error())
			}
		}
	} else {
		// Scenario 3: No real changes - skip both APIs
		tflog.Debug(ctx, "No meaningful changes detected, skipping API calls")
	}

	// Read the updated role to get the latest state
	readResp, err := r.ReadRole(ctx, roleName)
	if err != nil {
		resp.Diagnostics.AddError("API Read Failed", fmt.Sprintf("Error reading updated role: %v", err))
		return
	}

	// Fetch endpoint name if available
	if len(readResp.Roledetails) > 0 && readResp.Roledetails[0].Endpointkey != nil {
		plan.EndpointName = r.RoleResourceFetchEndpointName(ctx, readResp.Roledetails[0].Endpointkey, &resp.Diagnostics)
		if resp.Diagnostics.HasError() {
			return
		}
	}

	// Safe array access with bounds checking
	if len(readResp.Roledetails) == 0 {
		resp.Diagnostics.AddError("API Error", "No role details returned from API")
		return
	}
	roleDetails := readResp.Roledetails[0]

	// Map role details to model using helper function in IMPORT MODE
	if err := r.RoleResourceMapRoleDetailsToModel(&plan, &roleDetails, &state, &resp.Diagnostics, true); err != nil {
		return
	}

	// Add success message
	if apiResp != nil && apiResp.Requestid != nil && *apiResp.Requestid == "" && apiResp.Requestkey != nil && *apiResp.Requestkey == "" {
		resp.Diagnostics.AddWarning(
			"Role Updated - Manual UI Action Required",
			fmt.Sprintf(
				"Role updated successfully.\nMessage: %s\nErrorCode: %s\nRequestID: %s\nRequestKey: %s\n\n"+
					"⚠️  IMPORTANT: Role update has generated a pending request and created a new version.\n"+
					"Please log into the Saviynt UI and:\n"+
					"1. Navigate to the Pending Requests page to approve this role update\n"+
					"2. Check the Version Management page in the current role to review the new role version and make it active from composing\n"+
					"The role changes will not be fully active until the request is approved in the UI.",
				*apiResp.Message,
				*apiResp.ErrorCode,
				*apiResp.Requestid,
				*apiResp.Requestkey,
			),
		)
	} else if apiResp != nil && apiResp.Message != nil && *apiResp.Message != "" && apiResp.ErrorCode != nil && *apiResp.ErrorCode != "" {
		resp.Diagnostics.AddWarning(
			"Role Updated - Manual UI Action Required",
			fmt.Sprintf(
				"Role updated successfully.\nMessage: %s\nErrorCode: %s\n\n"+
					"⚠️  IMPORTANT: Role update has generated a pending request and created a new version.\n"+
					"Please log into the Saviynt UI and:\n"+
					"1. Navigate to the Pending Requests page to approve this role update\n"+
					"2. Check the Version Management page in the current role to review the new role version and make it active from composing\n"+
					"The role changes will not be fully active until the request is approved in the UI.",
				*apiResp.Message,
				*apiResp.ErrorCode,
			),
		)
	} else if hasUserChanges && !hasRealNonUserChanges {
		// User-only changes - role API was skipped, this is normal
		tflog.Info(ctx, "User-only changes completed successfully, role API was skipped")
	} else {
		resp.Diagnostics.AddWarning(
			"Info",
			"Provider error: received unexpected response from Saviynt API.Please retry or contact support.",
		)
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, plan)...)
	if userErr != nil {
		// Add a final warning if user changes had errors
		resp.Diagnostics.AddError("User Operation Failed", userErr.Error())
	}
}

// We do not support deletion of roles in Saviynt, so this function is intentionally left empty.
func (r *RolesResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
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
func (r *RolesResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	// Retrieve import ID and save to role_name attribute
	resource.ImportStatePassthroughID(ctx, path.Root("role_name"), req, resp)
	// Set the requestor to the current authenticated user
	resp.State.SetAttribute(ctx, path.Root("requestor"), r.requestor)
}
