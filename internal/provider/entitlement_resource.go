// Copyright (c) 2025 Saviynt Inc.
// SPDX-License-Identifier: MPL-2.0

// saviynt_entitlement_resource manages entitlements in the Saviynt Security Manager.
// The resource implements the full Terraform lifecycle:
//   - Create: provisions a new entitlement using the supplied configuration.
//   - Read: fetches the current entitlement state from Saviynt to keep Terraform’s state in sync.
//   - Update: applies any configuration changes to an existing entitlement.
//   - Import: brings an existing entitlement under Terraform management by its import key.

package provider

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"reflect"
	"regexp"
	"strings"

	"terraform-provider-Saviynt/internal/client"
	"terraform-provider-Saviynt/util"
	"terraform-provider-Saviynt/util/errorsutil"

	"github.com/hashicorp/terraform-plugin-framework-validators/mapvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"

	openapi "github.com/saviynt/saviynt-api-go-client/entitlements"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &EntitlementResource{}
var _ resource.ResourceWithImportState = &EntitlementResource{}

type EntitlementResource struct {
	client             client.SaviyntClientInterface
	token              string
	provider           client.SaviyntProviderInterface
	entitlementFactory client.EntitlementFactoryInterface
}

func NewEntitlementResource() resource.Resource {
	return &EntitlementResource{
		entitlementFactory: &client.DefaultEntitlementFactory{},
	}
}

func NewEntitlementResourceWithFactory(factory client.EntitlementFactoryInterface) resource.Resource {
	return &EntitlementResource{
		entitlementFactory: factory,
	}
}

type EntitlementResourceModel struct {
	ID                  types.String `tfsdk:"id"`
	Endpoint            types.String `tfsdk:"endpoint"`
	Entitlementtype     types.String `tfsdk:"entitlement_type"`
	EntitlementValue    types.String `tfsdk:"entitlement_value"`
	EntitlementValuekey types.String `tfsdk:"entitlement_valuekey"`
	Displayname         types.String `tfsdk:"displayname"`
	Risk                types.Int32  `tfsdk:"risk"`
	Status              types.Int32  `tfsdk:"status"`
	Soxcritical         types.Int32  `tfsdk:"soxcritical"`
	Syscritical         types.Int32  `tfsdk:"syscritical"`
	EntitlementGlossary types.String `tfsdk:"entitlement_glossary"`
	Privileged          types.Int32  `tfsdk:"privileged"`
	Module              types.String `tfsdk:"module"`
	Access              types.String `tfsdk:"access"`
	Priority            types.Int32  `tfsdk:"priority"`
	Description         types.String `tfsdk:"description"`
	Confidentiality     types.Int32  `tfsdk:"confidentiality"`
	Customproperty1     types.String `tfsdk:"customproperty1"`
	Customproperty2     types.String `tfsdk:"customproperty2"`
	Customproperty3     types.String `tfsdk:"customproperty3"`
	Customproperty4     types.String `tfsdk:"customproperty4"`
	Customproperty5     types.String `tfsdk:"customproperty5"`
	Customproperty6     types.String `tfsdk:"customproperty6"`
	Customproperty7     types.String `tfsdk:"customproperty7"`
	Customproperty8     types.String `tfsdk:"customproperty8"`
	Customproperty9     types.String `tfsdk:"customproperty9"`
	Customproperty10    types.String `tfsdk:"customproperty10"`
	Customproperty11    types.String `tfsdk:"customproperty11"`
	Customproperty12    types.String `tfsdk:"customproperty12"`
	Customproperty13    types.String `tfsdk:"customproperty13"`
	Customproperty14    types.String `tfsdk:"customproperty14"`
	Customproperty15    types.String `tfsdk:"customproperty15"`
	Customproperty16    types.String `tfsdk:"customproperty16"`
	Customproperty17    types.String `tfsdk:"customproperty17"`
	Customproperty18    types.String `tfsdk:"customproperty18"`
	Customproperty19    types.String `tfsdk:"customproperty19"`
	Customproperty20    types.String `tfsdk:"customproperty20"`
	Customproperty21    types.String `tfsdk:"customproperty21"`
	Customproperty22    types.String `tfsdk:"customproperty22"`
	Customproperty23    types.String `tfsdk:"customproperty23"`
	Customproperty24    types.String `tfsdk:"customproperty24"`
	Customproperty25    types.String `tfsdk:"customproperty25"`
	Customproperty26    types.String `tfsdk:"customproperty26"`
	Customproperty27    types.String `tfsdk:"customproperty27"`
	Customproperty28    types.String `tfsdk:"customproperty28"`
	Customproperty29    types.String `tfsdk:"customproperty29"`
	Customproperty30    types.String `tfsdk:"customproperty30"`
	Customproperty31    types.String `tfsdk:"customproperty31"`
	Customproperty32    types.String `tfsdk:"customproperty32"`
	Customproperty33    types.String `tfsdk:"customproperty33"`
	Customproperty34    types.String `tfsdk:"customproperty34"`
	Customproperty35    types.String `tfsdk:"customproperty35"`
	Customproperty36    types.String `tfsdk:"customproperty36"`
	Customproperty37    types.String `tfsdk:"customproperty37"`
	Customproperty38    types.String `tfsdk:"customproperty38"`
	Customproperty39    types.String `tfsdk:"customproperty39"`
	Customproperty40    types.String `tfsdk:"customproperty40"`

	EntitlementOwners types.Map `tfsdk:"entitlement_owners"`
	EntitlementMap    types.Set `tfsdk:"entitlement_map"`
}

type EntitlementMapModel struct {
	EntitlementValue       types.String `tfsdk:"entitlement_value"`         // Maps to Primary in API
	EntitlementType        types.String `tfsdk:"entitlement_type"`          // Maps to PrimaryEntType in API
	Endpoint               types.String `tfsdk:"endpoint"`                  // Not returned in get api but required while createUpdate
	EntitlementKey         types.String `tfsdk:"entitlement_key"`           // Maps to PrimaryEntKey in API (computed)
	RequestFilter          types.Bool   `tfsdk:"request_filter"`            // Maps to RequestFilter in API
	ExcludeEntitlement     types.Bool   `tfsdk:"exclude_entitlement"`       // Maps to ExcludeEntitlement in API
	AddDependentTask       types.Bool   `tfsdk:"add_dependent_task"`        // Maps to AddDependentTask in API
	RemoveDependentEntTask types.Bool   `tfsdk:"remove_dependent_ent_task"` // Maps to RemoveDependentEntTask in API
}

// entitlementMapsEqual compares two slices of EntitlementMapModel for equality
func entitlementMapsEqual(a, b []EntitlementMapModel) bool {
	log.Printf("[DEBUG] entitlementMapsEqual - Comparing slice A (len=%d): %+v", len(a), a)
	log.Printf("[DEBUG] entitlementMapsEqual - Comparing slice B (len=%d): %+v", len(b), b)

	if len(a) != len(b) {
		log.Printf("[DEBUG] entitlementMapsEqual - Length mismatch: %d != %d", len(a), len(b))
		return false
	}

	// If both are empty, they are equal
	if len(a) == 0 && len(b) == 0 {
		log.Printf("[DEBUG] entitlementMapsEqual - Both slices are empty, returning true")
		return true
	}

	// Create maps for comparison using entitlement_key as unique identifier
	mapA := make(map[string]EntitlementMapModel)
	mapB := make(map[string]EntitlementMapModel)

	for _, entMap := range a {
		// Always use composite key for consistent comparison
		key := entMap.Endpoint.ValueString() + "|" + entMap.EntitlementType.ValueString() + "|" + entMap.EntitlementValue.ValueString()
		mapA[key] = entMap
	}

	for _, entMap := range b {
		// Always use composite key for consistent comparison
		key := entMap.Endpoint.ValueString() + "|" + entMap.EntitlementType.ValueString() + "|" + entMap.EntitlementValue.ValueString()
		mapB[key] = entMap
	}

	// Compare maps
	for key, entMapA := range mapA {
		entMapB, exists := mapB[key]
		if !exists {
			log.Printf("[DEBUG] entitlementMapsEqual - Key %s exists in A but not in B", key)
			return false
		}

		// Compare all fields
		if entMapA.EntitlementValue.ValueString() != entMapB.EntitlementValue.ValueString() ||
			entMapA.EntitlementType.ValueString() != entMapB.EntitlementType.ValueString() ||
			entMapA.Endpoint.ValueString() != entMapB.Endpoint.ValueString() ||
			entMapA.RequestFilter.ValueBool() != entMapB.RequestFilter.ValueBool() ||
			entMapA.ExcludeEntitlement.ValueBool() != entMapB.ExcludeEntitlement.ValueBool() ||
			entMapA.AddDependentTask.ValueBool() != entMapB.AddDependentTask.ValueBool() ||
			entMapA.RemoveDependentEntTask.ValueBool() != entMapB.RemoveDependentEntTask.ValueBool() {
			log.Printf("[DEBUG] entitlementMapsEqual - Field mismatch for key %s:", key)
			log.Printf("[DEBUG]   A: EntitlementValue=%s, EntitlementType=%s, Endpoint=%s, RequestFilter=%t, ExcludeEntitlement=%t, AddDependentTask=%t, RemoveDependentEntTask=%t",
				entMapA.EntitlementValue.ValueString(), entMapA.EntitlementType.ValueString(), entMapA.Endpoint.ValueString(),
				entMapA.RequestFilter.ValueBool(), entMapA.ExcludeEntitlement.ValueBool(), entMapA.AddDependentTask.ValueBool(), entMapA.RemoveDependentEntTask.ValueBool())
			log.Printf("[DEBUG]   B: EntitlementValue=%s, EntitlementType=%s, Endpoint=%s, RequestFilter=%t, ExcludeEntitlement=%t, AddDependentTask=%t, RemoveDependentEntTask=%t",
				entMapB.EntitlementValue.ValueString(), entMapB.EntitlementType.ValueString(), entMapB.Endpoint.ValueString(),
				entMapB.RequestFilter.ValueBool(), entMapB.ExcludeEntitlement.ValueBool(), entMapB.AddDependentTask.ValueBool(), entMapB.RemoveDependentEntTask.ValueBool())
			return false
		}
	}

	log.Printf("[DEBUG] entitlementMapsEqual - All comparisons passed, returning true")
	return true
}

func (r *EntitlementResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = "saviynt_entitlement_resource"
}

func (r *EntitlementResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	attributes := map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:    true,
			Description: "Entitlement identifier",
		},
		"endpoint": schema.StringAttribute{
			Required:    true,
			Description: "Name of the endpoint for the entitlement",
		},
		"entitlement_type": schema.StringAttribute{
			Required:    true,
			Description: "Entitlement type for the entitlement",
		},
		"entitlement_value": schema.StringAttribute{
			Required:    true,
			Description: "Value of the entitlement",
		},
		"entitlement_valuekey": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "Key for the entitlement value",
		},
		"displayname": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "Display name of the entitlement",
		},
		"risk": schema.Int32Attribute{
			Optional:    true,
			Computed:    true,
			Description: "Indicates the risk score or level of the entitlement",
		},
		"status": schema.Int32Attribute{
			Optional:    true,
			Computed:    true,
			Description: "Status of the entitlement (e.g., active/inactive)",
		},
		"soxcritical": schema.Int32Attribute{
			Optional:    true,
			Computed:    true,
			Description: "SOX criticality flag",
		},
		"syscritical": schema.Int32Attribute{
			Optional:    true,
			Computed:    true,
			Description: "System criticality flag",
		},
		"entitlement_glossary": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "Glossary term or explanation for the entitlement",
		},
		"privileged": schema.Int32Attribute{
			Optional:    true,
			Computed:    true,
			Description: "Indicates if the entitlement is privileged",
		},
		"module": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "Functional module the entitlement belongs to",
		},
		"access": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "Access type or permission level",
		},
		"priority": schema.Int32Attribute{
			Optional:    true,
			Computed:    true,
			Description: "Priority level of the entitlement",
		},
		"description": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "Description of the entitlement",
		},
		"confidentiality": schema.Int32Attribute{
			Optional:    true,
			Computed:    true,
			Description: "Confidentiality classification level",
		},
		"entitlement_owners": schema.MapAttribute{
			ElementType: types.SetType{ElemType: types.StringType},
			Optional:    true,
			Description: "Map of owner ranks to list of usernames. Use 'rank_1', 'rank_2' etc",
			Validators: []validator.Map{
				// Validates rank keys follow 'rank_X' pattern where X is 1-27 (Saviynt's supported owner rank range)
				mapvalidator.KeysAre(stringvalidator.RegexMatches(
					regexp.MustCompile(`^rank_([1-9]|1[0-9]|2[0-7])$`),
					"rank must be in format 'rank_X' where X is between 1 and 27",
				)),
			},
		},
		"entitlement_map": schema.SetNestedAttribute{
			NestedObject: schema.NestedAttributeObject{
				Attributes: map[string]schema.Attribute{
					"entitlement_value": schema.StringAttribute{
						Required:    true,
						Description: "The entitlement value to map",
					},
					"entitlement_type": schema.StringAttribute{
						Required:    true,
						Description: "The entitlement type to map",
					},
					"endpoint": schema.StringAttribute{
						Required:    true,
						Description: "The endpoint for this mapping",
					},
					"entitlement_key": schema.StringAttribute{
						Computed:    true,
						Description: "The entitlement key from API (computed)",
					},
					"request_filter": schema.BoolAttribute{
						Optional:    true,
						Computed:    true,
						Description: "Request filter flag for the mapping",
					},
					"exclude_entitlement": schema.BoolAttribute{
						Optional:    true,
						Computed:    true,
						Description: "Exclude entitlement flag",
					},
					"add_dependent_task": schema.BoolAttribute{
						Optional:    true,
						Computed:    true,
						Description: "Add dependent task flag",
					},
					"remove_dependent_ent_task": schema.BoolAttribute{
						Optional:    true,
						Computed:    true,
						Description: "Remove dependent entitlement task flag",
					},
				},
			},
			Optional:    true,
			Description: "Set of entitlement mappings for hierarchical relationships",
		},
	}

	// Generate custom properties using reflection
	for i := 1; i <= 40; i++ {
		attributes[fmt.Sprintf("customproperty%d", i)] = schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: fmt.Sprintf("Custom property %d value", i),
		}
	}

	resp.Schema = schema.Schema{
		Description: util.EntitlementDescription,
		Attributes:  attributes,
	}
}

func (r *EntitlementResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		log.Println("[DEBUG] Entitlements: ProviderData is nil, returning early.")
		return
	}

	// Cast provider data to your provider type.
	prov, ok := req.ProviderData.(*SaviyntProvider)
	if !ok {
		log.Printf("[ERROR] Entitlements: Unexpected Provider Data")
		resp.Diagnostics.AddError("Unexpected Provider Data", "Expected *SaviyntProvider")
		return
	}

	// Set the client and token from the provider state.
	r.client = &client.SaviyntClientWrapper{Client: prov.client}
	r.token = prov.accessToken
	r.provider = &client.SaviyntProviderWrapper{Provider: prov} // Store provider reference for retry logic

	log.Printf("[DEBUG] Entitlements: Resource configured successfully.")
}

// SetClient sets the client for testing purposes
func (r *EntitlementResource) SetClient(client client.SaviyntClientInterface) {
	r.client = client
}

// SetToken sets the token for testing purposes
func (r *EntitlementResource) SetToken(token string) {
	r.token = token
}

// SetProvider sets the provider for testing purposes
func (r *EntitlementResource) SetProvider(provider client.SaviyntProviderInterface) {
	r.provider = provider
}

// BuildCreateEntitlementRequest builds the create request for entitlement
func (r *EntitlementResource) BuildCreateEntitlementRequest(plan *EntitlementResourceModel) openapi.CreateUpdateEntitlementRequest {
	createReq := openapi.CreateUpdateEntitlementRequest{
		EntitlementValue: plan.EntitlementValue.ValueString(),
		Entitlementtype:  plan.Entitlementtype.ValueString(),
		Endpoint:         plan.Endpoint.ValueString(),

		// Optional string attributes
		Displayname:         util.StringPointerOrEmpty(plan.Displayname),
		EntitlementGlossary: util.StringPointerOrEmpty(plan.EntitlementGlossary),
		Module:              util.StringPointerOrEmpty(plan.Module),
		Access:              util.StringPointerOrEmpty(plan.Access),
		Description:         util.StringPointerOrEmpty(plan.Description),

		// Optional boolean attributes
		Entitlementcasecheck: util.StringPtr("true"),

		// Optional int32 attributes
		Risk:            util.Int32PointerOrEmpty(plan.Risk),
		Status:          util.Int32PointerOrEmpty(plan.Status),
		Soxcritical:     util.Int32PointerOrEmpty(plan.Soxcritical),
		Syscritical:     util.Int32PointerOrEmpty(plan.Syscritical),
		Priviliged:      util.Int32PointerOrEmpty(plan.Privileged),
		Priority:        util.Int32PointerOrEmpty(plan.Priority),
		Confidentiality: util.Int32PointerOrEmpty(plan.Confidentiality),

		// Custom properties
		Customproperty1:  util.StringPointerOrEmpty(plan.Customproperty1),
		Customproperty2:  util.StringPointerOrEmpty(plan.Customproperty2),
		Customproperty3:  util.StringPointerOrEmpty(plan.Customproperty3),
		Customproperty4:  util.StringPointerOrEmpty(plan.Customproperty4),
		Customproperty5:  util.StringPointerOrEmpty(plan.Customproperty5),
		Customproperty6:  util.StringPointerOrEmpty(plan.Customproperty6),
		Customproperty7:  util.StringPointerOrEmpty(plan.Customproperty7),
		Customproperty8:  util.StringPointerOrEmpty(plan.Customproperty8),
		Customproperty9:  util.StringPointerOrEmpty(plan.Customproperty9),
		Customproperty10: util.StringPointerOrEmpty(plan.Customproperty10),
		Customproperty11: util.StringPointerOrEmpty(plan.Customproperty11),
		Customproperty12: util.StringPointerOrEmpty(plan.Customproperty12),
		Customproperty13: util.StringPointerOrEmpty(plan.Customproperty13),
		Customproperty14: util.StringPointerOrEmpty(plan.Customproperty14),
		Customproperty15: util.StringPointerOrEmpty(plan.Customproperty15),
		Customproperty16: util.StringPointerOrEmpty(plan.Customproperty16),
		Customproperty17: util.StringPointerOrEmpty(plan.Customproperty17),
		Customproperty18: util.StringPointerOrEmpty(plan.Customproperty18),
		Customproperty19: util.StringPointerOrEmpty(plan.Customproperty19),
		Customproperty20: util.StringPointerOrEmpty(plan.Customproperty20),
		Customproperty21: util.StringPointerOrEmpty(plan.Customproperty21),
		Customproperty22: util.StringPointerOrEmpty(plan.Customproperty22),
		Customproperty23: util.StringPointerOrEmpty(plan.Customproperty23),
		Customproperty24: util.StringPointerOrEmpty(plan.Customproperty24),
		Customproperty25: util.StringPointerOrEmpty(plan.Customproperty25),
		Customproperty26: util.StringPointerOrEmpty(plan.Customproperty26),
		Customproperty27: util.StringPointerOrEmpty(plan.Customproperty27),
		Customproperty28: util.StringPointerOrEmpty(plan.Customproperty28),
		Customproperty29: util.StringPointerOrEmpty(plan.Customproperty29),
		Customproperty30: util.StringPointerOrEmpty(plan.Customproperty30),
		Customproperty31: util.StringPointerOrEmpty(plan.Customproperty31),
		Customproperty32: util.StringPointerOrEmpty(plan.Customproperty32),
		Customproperty33: util.StringPointerOrEmpty(plan.Customproperty33),
		Customproperty34: util.StringPointerOrEmpty(plan.Customproperty34),
		Customproperty35: util.StringPointerOrEmpty(plan.Customproperty35),
		Customproperty36: util.StringPointerOrEmpty(plan.Customproperty36),
		Customproperty37: util.StringPointerOrEmpty(plan.Customproperty37),
		Customproperty38: util.StringPointerOrEmpty(plan.Customproperty38),
		Customproperty39: util.StringPointerOrEmpty(plan.Customproperty39),
		Customproperty40: util.StringPointerOrEmpty(plan.Customproperty40),
	}

	return createReq
}

// ProcessEntitlementOwnersForEntitlementCreatee processes entitlement owners for create requests
func (r *EntitlementResource) ProcessEntitlementOwnersForEntitlementCreatee(ctx context.Context, plan *EntitlementResourceModel, createReq *openapi.CreateUpdateEntitlementRequest) {
	if !plan.EntitlementOwners.IsNull() {
		var ownersMap map[string][]string
		plan.EntitlementOwners.ElementsAs(ctx, &ownersMap, false)

		for rankKey, users := range ownersMap {
			var rankNum int
			fmt.Sscanf(rankKey, "rank_%d", &rankNum)

			var usersWithAdd []string
			for _, user := range users {
				usersWithAdd = append(usersWithAdd, user+"##add")
			}
			fieldName := fmt.Sprintf("Entitlementowner%d", rankNum)
			field := reflect.ValueOf(createReq).Elem().FieldByName(fieldName)
			if field.IsValid() && field.CanSet() {
				field.Set(reflect.ValueOf(usersWithAdd))
			}
		}
	}
}

// ProcessEntitlementMap processes entitlement map for create requests
func (r *EntitlementResource) ProcessEntitlementMapForEntitlementCreate(ctx context.Context, plan *EntitlementResourceModel, createReq *openapi.CreateUpdateEntitlementRequest) {
	if !plan.EntitlementMap.IsNull() {
		var entitlementMaps []EntitlementMapModel
		plan.EntitlementMap.ElementsAs(ctx, &entitlementMaps, false)

		var apiEntitlementMaps []openapi.CreateUpdateEntitlementRequestEntitlementmapInner
		for _, entMap := range entitlementMaps {
			apiEntitlementMaps = append(apiEntitlementMaps, r.BuildEntitlementMapAPIObject(entMap, "ADD"))
		}
		createReq.Entitlementmap = apiEntitlementMaps
	}
}

// CreateEntitlement creates an entitlement
func (r *EntitlementResource) CreateEntitlement(ctx context.Context, plan *EntitlementResourceModel) (*openapi.CreateOrUpdateEntitlementResponse, error) {
	log.Printf("[DEBUG] Entitlements: Starting creation for entitlement: %s", plan.EntitlementValue.ValueString())

	createReq := r.BuildCreateEntitlementRequest(plan)
	r.ProcessEntitlementOwnersForEntitlementCreatee(ctx, plan, &createReq)
	r.ProcessEntitlementMapForEntitlementCreate(ctx, plan, &createReq)

	log.Printf("[DEBUG] Entitlements: Final create request JSON: %+v", createReq)

	// Execute create operation with retry logic
	var createResp *openapi.CreateOrUpdateEntitlementResponse
	var createHttpResp *http.Response
	err := r.provider.AuthenticatedAPICallWithRetry(ctx, "create_entitlement", func(token string) error {
		entitlementOps := r.entitlementFactory.CreateEntitlementOperations(r.client.APIBaseURL(), token)
		resp, hResp, err := entitlementOps.CreateUpdateEntitlement(ctx, createReq)
		if hResp != nil && hResp.StatusCode == 401 {
			return fmt.Errorf("401 unauthorized")
		}
		createResp = resp
		createHttpResp = hResp
		return err
	})

	if err != nil {
		// Handle API call failures and HTTP errors during entitlement creation
		log.Printf("[ERROR] Entitlements: Error in creating entitlement: %v", err)
		err = errorsutil.HandleHTTPError(createHttpResp, err, "Create")
		return nil, fmt.Errorf("error creating entitlement: %v", err)
	}
	if createResp != nil && createResp.ErrorCode != nil && *createResp.ErrorCode == "1" {
		msg := util.SafeDeref(createResp.Msg)
		log.Printf("[ERROR] Entitlements: Error in creating entitlement: %v", msg)
		return nil, fmt.Errorf("error creating entitlement: %v", msg)
	}

	return createResp, nil
}

// SetDefaultComputedValues sets default values for computed fields
func (r *EntitlementResource) SetDefaultComputedValues(plan *EntitlementResourceModel) {
	if plan.Soxcritical.IsNull() || plan.Soxcritical.IsUnknown() {
		plan.Soxcritical = types.Int32Value(0)
	}
	if plan.Syscritical.IsNull() || plan.Syscritical.IsUnknown() {
		plan.Syscritical = types.Int32Value(0)
	}
	if plan.Status.IsNull() || plan.Status.IsUnknown() {
		plan.Status = types.Int32Value(1)
	}

	// Set computed fields to ensure they have known values
	plan.Displayname = util.SafeStringDatasource(plan.Displayname.ValueStringPointer())
	plan.Description = util.SafeStringDatasource(plan.Description.ValueStringPointer())
	plan.EntitlementGlossary = util.SafeStringDatasource(plan.EntitlementGlossary.ValueStringPointer())
	plan.Module = util.SafeStringDatasource(plan.Module.ValueStringPointer())
	plan.Access = util.SafeStringDatasource(plan.Access.ValueStringPointer())
	plan.Risk = util.SafeInt32(plan.Risk.ValueInt32Pointer())
	plan.Privileged = util.SafeInt32(plan.Privileged.ValueInt32Pointer())
	plan.Priority = util.SafeInt32(plan.Priority.ValueInt32Pointer())
	plan.Confidentiality = util.SafeInt32(plan.Confidentiality.ValueInt32Pointer())

	// Set custom properties to ensure they have known values
	plan.Customproperty1 = util.SafeStringDatasource(plan.Customproperty1.ValueStringPointer())
	plan.Customproperty2 = util.SafeStringDatasource(plan.Customproperty2.ValueStringPointer())
	plan.Customproperty3 = util.SafeStringDatasource(plan.Customproperty3.ValueStringPointer())
	plan.Customproperty4 = util.SafeStringDatasource(plan.Customproperty4.ValueStringPointer())
	plan.Customproperty5 = util.SafeStringDatasource(plan.Customproperty5.ValueStringPointer())
	plan.Customproperty6 = util.SafeStringDatasource(plan.Customproperty6.ValueStringPointer())
	plan.Customproperty7 = util.SafeStringDatasource(plan.Customproperty7.ValueStringPointer())
	plan.Customproperty8 = util.SafeStringDatasource(plan.Customproperty8.ValueStringPointer())
	plan.Customproperty9 = util.SafeStringDatasource(plan.Customproperty9.ValueStringPointer())
	plan.Customproperty10 = util.SafeStringDatasource(plan.Customproperty10.ValueStringPointer())
	plan.Customproperty11 = util.SafeStringDatasource(plan.Customproperty11.ValueStringPointer())
	plan.Customproperty12 = util.SafeStringDatasource(plan.Customproperty12.ValueStringPointer())
	plan.Customproperty13 = util.SafeStringDatasource(plan.Customproperty13.ValueStringPointer())
	plan.Customproperty14 = util.SafeStringDatasource(plan.Customproperty14.ValueStringPointer())
	plan.Customproperty15 = util.SafeStringDatasource(plan.Customproperty15.ValueStringPointer())
	plan.Customproperty16 = util.SafeStringDatasource(plan.Customproperty16.ValueStringPointer())
	plan.Customproperty17 = util.SafeStringDatasource(plan.Customproperty17.ValueStringPointer())
	plan.Customproperty18 = util.SafeStringDatasource(plan.Customproperty18.ValueStringPointer())
	plan.Customproperty19 = util.SafeStringDatasource(plan.Customproperty19.ValueStringPointer())
	plan.Customproperty20 = util.SafeStringDatasource(plan.Customproperty20.ValueStringPointer())
	plan.Customproperty21 = util.SafeStringDatasource(plan.Customproperty21.ValueStringPointer())
	plan.Customproperty22 = util.SafeStringDatasource(plan.Customproperty22.ValueStringPointer())
	plan.Customproperty23 = util.SafeStringDatasource(plan.Customproperty23.ValueStringPointer())
	plan.Customproperty24 = util.SafeStringDatasource(plan.Customproperty24.ValueStringPointer())
	plan.Customproperty25 = util.SafeStringDatasource(plan.Customproperty25.ValueStringPointer())
	plan.Customproperty26 = util.SafeStringDatasource(plan.Customproperty26.ValueStringPointer())
	plan.Customproperty27 = util.SafeStringDatasource(plan.Customproperty27.ValueStringPointer())
	plan.Customproperty28 = util.SafeStringDatasource(plan.Customproperty28.ValueStringPointer())
	plan.Customproperty29 = util.SafeStringDatasource(plan.Customproperty29.ValueStringPointer())
	plan.Customproperty30 = util.SafeStringDatasource(plan.Customproperty30.ValueStringPointer())
	plan.Customproperty31 = util.SafeStringDatasource(plan.Customproperty31.ValueStringPointer())
	plan.Customproperty32 = util.SafeStringDatasource(plan.Customproperty32.ValueStringPointer())
	plan.Customproperty33 = util.SafeStringDatasource(plan.Customproperty33.ValueStringPointer())
	plan.Customproperty34 = util.SafeStringDatasource(plan.Customproperty34.ValueStringPointer())
	plan.Customproperty35 = util.SafeStringDatasource(plan.Customproperty35.ValueStringPointer())
	plan.Customproperty36 = util.SafeStringDatasource(plan.Customproperty36.ValueStringPointer())
	plan.Customproperty37 = util.SafeStringDatasource(plan.Customproperty37.ValueStringPointer())
	plan.Customproperty38 = util.SafeStringDatasource(plan.Customproperty38.ValueStringPointer())
	plan.Customproperty39 = util.SafeStringDatasource(plan.Customproperty39.ValueStringPointer())
	plan.Customproperty40 = util.SafeStringDatasource(plan.Customproperty40.ValueStringPointer())
}

// ReadEntitlement reads an entitlement from the API
func (r *EntitlementResource) ReadEntitlement(ctx context.Context, state *EntitlementResourceModel) (*openapi.GetEntitlementResponse, error) {
	log.Printf("[DEBUG] Entitlements: Starting reading for entitlement id: %s", state.EntitlementValuekey.ValueString())

	readReq := openapi.GetEntitlementRequest{
		Endpoint:             util.StringPtr(state.Endpoint.ValueString()),
		Entitlementtype:      util.StringPtr(state.Entitlementtype.ValueString()),
		EntitlementValue:     util.StringPtr(state.EntitlementValue.ValueString()),
		Entownerwithrank:     util.StringPtr("true"),
		Returnentitlementmap: util.StringPtr("true"),
	}

	log.Printf("[DEBUG] Read request - Endpoint: %s, EntitlementType: %s, EntitlementValue: %s",
		state.Endpoint.ValueString(), state.Entitlementtype.ValueString(), state.EntitlementValue.ValueString())

	// Execute read operation with retry logic
	var readResp *openapi.GetEntitlementResponse
	var readHttpResp *http.Response
	err := r.provider.AuthenticatedAPICallWithRetry(ctx, "read_entitlement", func(token string) error {
		entitlementOps := r.entitlementFactory.CreateEntitlementOperations(r.client.APIBaseURL(), token)
		resp, hResp, err := entitlementOps.GetEntitlements(ctx, readReq)
		if hResp != nil && hResp.StatusCode == 401 {
			return fmt.Errorf("401 unauthorized")
		}
		readResp = resp
		readHttpResp = hResp
		return err
	})

	if err != nil {
		// Handle API call failures and HTTP errors during entitlement read
		log.Printf("[ERROR] Entitlements: Error in reading entitlement: %v", err)
		err = errorsutil.HandleHTTPError(readHttpResp, err, "Read")
		return nil, fmt.Errorf("error reading entitlement: %v", err)
	}
	if readResp != nil && readResp.ErrorCode != nil && *readResp.ErrorCode != "0" {
		errorCode := util.SafeDeref(readResp.ErrorCode)
		msg := util.SafeDeref(readResp.Msg)
		log.Printf("[ERROR] Entitlements: API returned error code 1: %v", msg)
		return nil, fmt.Errorf("error reading entitlement. Error code: %v, Msg: %v", errorCode, msg)
	}

	return readResp, nil
}

// PopulateStateFromAPI populates the state model from API response
func (r *EntitlementResource) PopulateStateFromAPI(ctx context.Context, state *EntitlementResourceModel, entitlement openapi.GetEntitlementResponseEntitlementdetailsInner) {
	// Set basic fields
	state.ID = util.SafeString(entitlement.EntitlementValuekey)
	state.Endpoint = util.SafeString(entitlement.Endpoint)
	state.Entitlementtype = util.SafeString(entitlement.EntitlementType)
	state.EntitlementValue = util.SafeString(entitlement.EntitlementValue)
	state.EntitlementValuekey = util.SafeString(entitlement.EntitlementValuekey)
	state.Displayname = util.SafeString(entitlement.Displayname)
	state.Description = util.SafeString(entitlement.Description)
	state.EntitlementGlossary = util.SafeString(entitlement.EntitlementGlossary)

	// Set numeric fields (convert string to int32)
	state.Status = util.SafeInt32(util.StringPtrToInt32Ptr(entitlement.Status))
	state.Soxcritical = util.SafeInt32(util.StringPtrToInt32Ptr(entitlement.Soxcritical))
	state.Syscritical = util.SafeInt32(util.StringPtrToInt32Ptr(entitlement.Syscritical))
	state.Risk = util.SafeInt32(util.StringPtrToInt32Ptr(entitlement.Risk))
	state.Priority = util.SafeInt32(util.StringPtrToInt32Ptr(entitlement.Priority))
	state.Privileged = util.SafeInt32(util.StringPtrToInt32Ptr(entitlement.Priviliged))
	state.Confidentiality = util.SafeInt32(util.StringPtrToInt32Ptr(entitlement.Confidentiality))

	// Set other string fields
	state.Module = util.SafeString(entitlement.Module)
	state.Access = util.SafeString(entitlement.Access)

	// Handle custom properties 1-20 (CustomProperty format)
	for i := 1; i <= 20; i++ {
		fieldName := fmt.Sprintf("CustomProperty%d", i)
		if field := reflect.ValueOf(&entitlement).Elem().FieldByName(fieldName); field.IsValid() && !field.IsNil() {
			planField := reflect.ValueOf(state).Elem().FieldByName(fmt.Sprintf("Customproperty%d", i))
			if planField.IsValid() {
				planField.Set(reflect.ValueOf(util.SafeString(field.Interface().(*string))))
			}
		}
	}

	// Handle custom properties 21-40 (Customproperty format)
	for i := 21; i <= 40; i++ {
		fieldName := fmt.Sprintf("Customproperty%d", i)
		if field := reflect.ValueOf(&entitlement).Elem().FieldByName(fieldName); field.IsValid() && !field.IsNil() {
			planField := reflect.ValueOf(state).Elem().FieldByName(fmt.Sprintf("Customproperty%d", i))
			if planField.IsValid() {
				planField.Set(reflect.ValueOf(util.SafeString(field.Interface().(*string))))
			}
		}
	}
}

// ProcessEntitlementOwnersForEntitlementRead processes entitlement owners during read operations
func (r *EntitlementResource) ProcessEntitlementOwnersForEntitlementRead(ctx context.Context, state *EntitlementResourceModel, entitlement openapi.GetEntitlementResponseEntitlementdetailsInner, isImportOrPostUpdate bool) error {
	if isImportOrPostUpdate {
		// True import: Import all owners from API
		if entitlement.EntitlementOwner != nil && entitlement.EntitlementOwner.MapmapOfStringarrayOfString != nil {
			ownersMap := make(map[string][]string)
			for rank, users := range *entitlement.EntitlementOwner.MapmapOfStringarrayOfString {
				rankKey := strings.ToLower(strings.Replace(rank, " ", "_", -1))
				ownersMap[rankKey] = users
			}
			state.EntitlementOwners, _ = types.MapValueFrom(ctx, types.SetType{ElemType: types.StringType}, ownersMap)
		} else {
			// No owners in API - only set empty map if it was previously configured, otherwise keep null
			if !state.EntitlementOwners.IsNull() {
				state.EntitlementOwners, _ = types.MapValueFrom(ctx, types.SetType{ElemType: types.StringType}, make(map[string][]string))
			}
			// If it was null, leave it null
		}
	} else {
		// Regular operation: Always validate counts per rank
		// Get state owners (treat null as empty map)
		var stateOwners map[string][]string
		if !state.EntitlementOwners.IsNull() {
			state.EntitlementOwners.ElementsAs(ctx, &stateOwners, false)
		} else {
			stateOwners = make(map[string][]string)
		}

		// Get API owners
		apiOwners := make(map[string][]string)
		if entitlement.EntitlementOwner != nil && entitlement.EntitlementOwner.MapmapOfStringarrayOfString != nil {
			for rank, users := range *entitlement.EntitlementOwner.MapmapOfStringarrayOfString {
				rankKey := strings.ToLower(strings.Replace(rank, " ", "_", -1))
				apiOwners[rankKey] = users
			}
		}

		// Compare each rank individually
		allRanks := make(map[string]bool)
		for rank := range stateOwners {
			allRanks[rank] = true
		}
		for rank := range apiOwners {
			allRanks[rank] = true
		}

		// Check for drift in any rank
		driftDetected := false
		var driftDetails []string
		for rank := range allRanks {
			stateUsers := stateOwners[rank]
			apiUsers := apiOwners[rank]

			// Compare user lists (order-independent)
			if !util.StringSlicesEqual(stateUsers, apiUsers) {
				driftDetected = true
				driftDetails = append(driftDetails, fmt.Sprintf("rank %s: state has %v, API has %v", rank, stateUsers, apiUsers))
			}
		}

		if driftDetected {
			log.Printf("[DEBUG] Entitlements: Drift detected in entitlement owners: %+v", driftDetails)
			return fmt.Errorf("Entitlement Owner Drift Detected: Entitlement owners have been modified outside of Terraform. Please run 'terraform import' to sync the current state, or manually revert the changes in the UI")
		}

		// If no drift, sync the current API state to Terraform state
		if len(apiOwners) > 0 {
			state.EntitlementOwners, _ = types.MapValueFrom(ctx, types.SetType{ElemType: types.StringType}, apiOwners)
		} else if !state.EntitlementOwners.IsNull() {
			// API has no owners but state exists - set to empty
			state.EntitlementOwners, _ = types.MapValueFrom(ctx, types.SetType{ElemType: types.StringType}, make(map[string][]string))
		}
	}
	return nil
}

// ProcessEntitlementMapForEntitlementRead processes entitlement map during read operations
func (r *EntitlementResource) ProcessEntitlementMapForEntitlementRead(ctx context.Context, state *EntitlementResourceModel, entitlement openapi.GetEntitlementResponseEntitlementdetailsInner, isImportOrPostUpdate bool) error {
	if isImportOrPostUpdate {
		// True import: Import all maps from API
		if len(entitlement.EntitlementMapDetails) > 0 {
			var entitlementMaps []EntitlementMapModel
			for _, mapDetail := range entitlement.EntitlementMapDetails {
				// Get endpoint from primaryEntKey
				endpointValue := ""
				if mapDetail.PrimaryEntKey != nil {
					entQuery := fmt.Sprintf("ent.id like '%s'", *mapDetail.PrimaryEntKey)
					entReq := openapi.GetEntitlementRequest{
						EntQuery: &entQuery,
					}

					// Execute validation call with retry logic
					err := r.provider.AuthenticatedAPICallWithRetry(ctx, "validate_entitlement_map", func(token string) error {
						entitlementOps := r.entitlementFactory.CreateEntitlementOperations(r.client.APIBaseURL(), token)
						entResp, httpResp, err := entitlementOps.GetEntitlements(ctx, entReq)
						if httpResp != nil && httpResp.StatusCode == 401 {
							return fmt.Errorf("401 unauthorized")
						}
						if err == nil && entResp != nil && len(entResp.Entitlementdetails) > 0 {
							if entResp.Entitlementdetails[0].Endpoint != nil {
								endpointValue = *entResp.Entitlementdetails[0].Endpoint
							}
						}
						return err
					})

					if err != nil {
						log.Printf("[WARNING] Failed to validate entitlement map: %v", err)
					}
				}

				entMap := EntitlementMapModel{
					EntitlementValue:       util.SafeString(mapDetail.Primary),
					EntitlementType:        util.SafeString(mapDetail.PrimaryEntType),
					Endpoint:               types.StringValue(endpointValue),
					EntitlementKey:         util.SafeString(mapDetail.PrimaryEntKey),
					RequestFilter:          util.SafeBoolDatasource(mapDetail.RequestFilter),
					ExcludeEntitlement:     util.SafeBoolDatasource(mapDetail.ExcludeEntitlement),
					AddDependentTask:       util.SafeBoolDatasource(mapDetail.AddDependentTask),
					RemoveDependentEntTask: util.SafeBoolDatasource(mapDetail.RemoveDependentEntTask),
				}
				entitlementMaps = append(entitlementMaps, entMap)
			}
			state.EntitlementMap, _ = types.SetValueFrom(ctx, types.ObjectType{
				AttrTypes: map[string]attr.Type{
					"entitlement_value":         types.StringType,
					"entitlement_type":          types.StringType,
					"endpoint":                  types.StringType,
					"entitlement_key":           types.StringType,
					"request_filter":            types.BoolType,
					"exclude_entitlement":       types.BoolType,
					"add_dependent_task":        types.BoolType,
					"remove_dependent_ent_task": types.BoolType,
				},
			}, entitlementMaps)
		} else {
			// No maps in API - only set empty set if state previously had maps
			if !state.EntitlementMap.IsNull() {
				emptyMaps := []EntitlementMapModel{}
				state.EntitlementMap, _ = types.SetValueFrom(ctx, types.ObjectType{
					AttrTypes: map[string]attr.Type{
						"entitlement_value":         types.StringType,
						"entitlement_type":          types.StringType,
						"endpoint":                  types.StringType,
						"entitlement_key":           types.StringType,
						"request_filter":            types.BoolType,
						"exclude_entitlement":       types.BoolType,
						"add_dependent_task":        types.BoolType,
						"remove_dependent_ent_task": types.BoolType,
					},
				}, emptyMaps)
			}
			// If state was null, leave it null
		}
	} else {
		// Regular operation: Always validate counts per entitlement map
		// Get state maps (treat null as empty)
		var stateMaps []EntitlementMapModel
		if !state.EntitlementMap.IsNull() {
			state.EntitlementMap.ElementsAs(ctx, &stateMaps, false)
		}

		// Get API maps
		var apiMaps []EntitlementMapModel
		if len(entitlement.EntitlementMapDetails) > 0 {
			for _, mapDetail := range entitlement.EntitlementMapDetails {
				// Get endpoint from primaryEntKey
				endpointValue := ""
				if mapDetail.PrimaryEntKey != nil {
					entQuery := fmt.Sprintf("ent.id like '%s'", *mapDetail.PrimaryEntKey)
					entReq := openapi.GetEntitlementRequest{
						EntQuery: &entQuery,
					}

					// Execute validation call with retry logic
					err := r.provider.AuthenticatedAPICallWithRetry(ctx, "validate_entitlement_map_2", func(token string) error {
						entitlementOps := r.entitlementFactory.CreateEntitlementOperations(r.client.APIBaseURL(), token)
						entResp, httpResp, err := entitlementOps.GetEntitlements(ctx, entReq)
						if httpResp != nil && httpResp.StatusCode == 401 {
							return fmt.Errorf("401 unauthorized")
						}
						if err == nil && entResp != nil && len(entResp.Entitlementdetails) > 0 {
							if entResp.Entitlementdetails[0].Endpoint != nil {
								endpointValue = *entResp.Entitlementdetails[0].Endpoint
							}
						}
						return err
					})

					if err != nil {
						log.Printf("[WARNING] Failed to validate entitlement map: %v", err)
					}
				}

				entMap := EntitlementMapModel{
					EntitlementValue:       util.SafeString(mapDetail.Primary),
					EntitlementType:        util.SafeString(mapDetail.PrimaryEntType),
					Endpoint:               types.StringValue(endpointValue),
					EntitlementKey:         util.SafeString(mapDetail.PrimaryEntKey),
					RequestFilter:          util.SafeBoolDatasource(mapDetail.RequestFilter),
					ExcludeEntitlement:     util.SafeBoolDatasource(mapDetail.ExcludeEntitlement),
					AddDependentTask:       util.SafeBoolDatasource(mapDetail.AddDependentTask),
					RemoveDependentEntTask: util.SafeBoolDatasource(mapDetail.RemoveDependentEntTask),
				}
				apiMaps = append(apiMaps, entMap)
			}
		}

		// Compare actual values instead of just counts
		log.Printf("[DEBUG] Entitlement Map Comparison - State maps: %+v", stateMaps)
		log.Printf("[DEBUG] Entitlement Map Comparison - API maps: %+v", apiMaps)
		log.Printf("[DEBUG] Entitlement Map Comparison - State count: %d, API count: %d", len(stateMaps), len(apiMaps))

		if !entitlementMapsEqual(stateMaps, apiMaps) {
			log.Printf("[ERROR] Entitlement Map Drift - Maps are not equal")
			return fmt.Errorf("Entitlement Map Drift Detected: Entitlement map has been modified outside of Terraform. Please run 'terraform import' to sync the current state, or manually revert the changes in the UI")
		}

		log.Printf("[DEBUG] Entitlement Map Comparison - Maps are equal, no drift detected")

		// If no drift, sync the current API state
		if len(apiMaps) > 0 {
			state.EntitlementMap, _ = types.SetValueFrom(ctx, types.ObjectType{
				AttrTypes: map[string]attr.Type{
					"entitlement_value":         types.StringType,
					"entitlement_type":          types.StringType,
					"endpoint":                  types.StringType,
					"entitlement_key":           types.StringType,
					"request_filter":            types.BoolType,
					"exclude_entitlement":       types.BoolType,
					"add_dependent_task":        types.BoolType,
					"remove_dependent_ent_task": types.BoolType,
				},
			}, apiMaps)
		} else if !state.EntitlementMap.IsNull() {
			// API has no maps but state exists - set to empty only if state was not null
			emptyMaps := []EntitlementMapModel{}
			state.EntitlementMap, _ = types.SetValueFrom(ctx, types.ObjectType{
				AttrTypes: map[string]attr.Type{
					"entitlement_value":         types.StringType,
					"entitlement_type":          types.StringType,
					"endpoint":                  types.StringType,
					"entitlement_key":           types.StringType,
					"request_filter":            types.BoolType,
					"exclude_entitlement":       types.BoolType,
					"add_dependent_task":        types.BoolType,
					"remove_dependent_ent_task": types.BoolType,
				},
			}, emptyMaps)
		}
		// If state was null and API has no maps, leave it null
	}
	return nil
}

// BuildUpdateEntitlementRequest builds the update request for entitlement
func (r *EntitlementResource) BuildUpdateEntitlementRequest(plan *EntitlementResourceModel, state *EntitlementResourceModel) *openapi.CreateUpdateEntitlementRequest {
	updateReq := openapi.NewCreateUpdateEntitlementRequest(
		plan.Endpoint.ValueString(),
		plan.Entitlementtype.ValueString(),
		state.EntitlementValue.ValueString(), // Use current value for lookup
	)

	// Set optional fields
	updateReq.Displayname = util.StringPointerOrEmpty(plan.Displayname)
	updateReq.EntitlementGlossary = util.StringPointerOrEmpty(plan.EntitlementGlossary)
	updateReq.Module = util.StringPointerOrEmpty(plan.Module)
	updateReq.Access = util.StringPointerOrEmpty(plan.Access)
	updateReq.Description = util.StringPointerOrEmpty(plan.Description)
	updateReq.Entitlementcasecheck = util.StringPtr("true")
	updateReq.Risk = util.Int32PointerOrEmpty(plan.Risk)
	updateReq.Status = util.Int32PointerOrEmpty(plan.Status)
	updateReq.Soxcritical = util.Int32PointerOrEmpty(plan.Soxcritical)
	updateReq.Syscritical = util.Int32PointerOrEmpty(plan.Syscritical)
	updateReq.Priviliged = util.Int32PointerOrEmpty(plan.Privileged)
	updateReq.Priority = util.Int32PointerOrEmpty(plan.Priority)
	updateReq.Confidentiality = util.Int32PointerOrEmpty(plan.Confidentiality)

	// Custom properties
	updateReq.Customproperty1 = util.StringPointerOrEmpty(plan.Customproperty1)
	updateReq.Customproperty2 = util.StringPointerOrEmpty(plan.Customproperty2)
	updateReq.Customproperty3 = util.StringPointerOrEmpty(plan.Customproperty3)
	updateReq.Customproperty4 = util.StringPointerOrEmpty(plan.Customproperty4)
	updateReq.Customproperty5 = util.StringPointerOrEmpty(plan.Customproperty5)
	updateReq.Customproperty6 = util.StringPointerOrEmpty(plan.Customproperty6)
	updateReq.Customproperty7 = util.StringPointerOrEmpty(plan.Customproperty7)
	updateReq.Customproperty8 = util.StringPointerOrEmpty(plan.Customproperty8)
	updateReq.Customproperty9 = util.StringPointerOrEmpty(plan.Customproperty9)
	updateReq.Customproperty10 = util.StringPointerOrEmpty(plan.Customproperty10)
	updateReq.Customproperty11 = util.StringPointerOrEmpty(plan.Customproperty11)
	updateReq.Customproperty12 = util.StringPointerOrEmpty(plan.Customproperty12)
	updateReq.Customproperty13 = util.StringPointerOrEmpty(plan.Customproperty13)
	updateReq.Customproperty14 = util.StringPointerOrEmpty(plan.Customproperty14)
	updateReq.Customproperty15 = util.StringPointerOrEmpty(plan.Customproperty15)
	updateReq.Customproperty16 = util.StringPointerOrEmpty(plan.Customproperty16)
	updateReq.Customproperty17 = util.StringPointerOrEmpty(plan.Customproperty17)
	updateReq.Customproperty18 = util.StringPointerOrEmpty(plan.Customproperty18)
	updateReq.Customproperty19 = util.StringPointerOrEmpty(plan.Customproperty19)
	updateReq.Customproperty20 = util.StringPointerOrEmpty(plan.Customproperty20)
	updateReq.Customproperty21 = util.StringPointerOrEmpty(plan.Customproperty21)
	updateReq.Customproperty22 = util.StringPointerOrEmpty(plan.Customproperty22)
	updateReq.Customproperty23 = util.StringPointerOrEmpty(plan.Customproperty23)
	updateReq.Customproperty24 = util.StringPointerOrEmpty(plan.Customproperty24)
	updateReq.Customproperty25 = util.StringPointerOrEmpty(plan.Customproperty25)
	updateReq.Customproperty26 = util.StringPointerOrEmpty(plan.Customproperty26)
	updateReq.Customproperty27 = util.StringPointerOrEmpty(plan.Customproperty27)
	updateReq.Customproperty28 = util.StringPointerOrEmpty(plan.Customproperty28)
	updateReq.Customproperty29 = util.StringPointerOrEmpty(plan.Customproperty29)
	updateReq.Customproperty30 = util.StringPointerOrEmpty(plan.Customproperty30)
	updateReq.Customproperty31 = util.StringPointerOrEmpty(plan.Customproperty31)
	updateReq.Customproperty32 = util.StringPointerOrEmpty(plan.Customproperty32)
	updateReq.Customproperty33 = util.StringPointerOrEmpty(plan.Customproperty33)
	updateReq.Customproperty34 = util.StringPointerOrEmpty(plan.Customproperty34)
	updateReq.Customproperty35 = util.StringPointerOrEmpty(plan.Customproperty35)
	updateReq.Customproperty36 = util.StringPointerOrEmpty(plan.Customproperty36)
	updateReq.Customproperty37 = util.StringPointerOrEmpty(plan.Customproperty37)
	updateReq.Customproperty38 = util.StringPointerOrEmpty(plan.Customproperty38)
	updateReq.Customproperty39 = util.StringPointerOrEmpty(plan.Customproperty39)
	updateReq.Customproperty40 = util.StringPointerOrEmpty(plan.Customproperty40)

	if !plan.EntitlementValue.Equal(state.EntitlementValue) {
		updateReq.UpdatedentitlementValue = util.StringPointerOrEmpty(plan.EntitlementValue)
	}

	return updateReq
}

// ProcessEntitlementOwnersForEntitlementUpdate processes entitlement owners for update requests
func (r *EntitlementResource) ProcessEntitlementOwnersForEntitlementUpdate(ctx context.Context, plan *EntitlementResourceModel, state *EntitlementResourceModel, updateReq *openapi.CreateUpdateEntitlementRequest) {
	if !plan.EntitlementOwners.IsNull() || !state.EntitlementOwners.IsNull() {
		var planOwners, stateOwners map[string][]string

		if !plan.EntitlementOwners.IsNull() {
			plan.EntitlementOwners.ElementsAs(ctx, &planOwners, false)
		}
		if !state.EntitlementOwners.IsNull() {
			state.EntitlementOwners.ElementsAs(ctx, &stateOwners, false)
		}

		allRanks := make(map[string]bool)
		for rank := range planOwners {
			allRanks[rank] = true
		}
		for rank := range stateOwners {
			allRanks[rank] = true
		}

		for rankKey := range allRanks {
			var rankNum int
			fmt.Sscanf(rankKey, "rank_%d", &rankNum)

			planUsers := planOwners[rankKey]
			stateUsers := stateOwners[rankKey]

			toAdd := util.Difference(planUsers, stateUsers)
			toRemove := util.Difference(stateUsers, planUsers)

			var finalUsers []string
			for _, user := range toAdd {
				finalUsers = append(finalUsers, user+"##add")
			}
			for _, user := range toRemove {
				finalUsers = append(finalUsers, user+"##remove")
			}

			if len(finalUsers) > 0 {
				fieldName := fmt.Sprintf("Entitlementowner%d", rankNum)
				field := reflect.ValueOf(updateReq).Elem().FieldByName(fieldName)
				if field.IsValid() && field.CanSet() {
					field.Set(reflect.ValueOf(finalUsers))
				}
			}
		}
	}
}

// ProcessEntitlementMapForEntitlementUpdate processes entitlement map for update requests
func (r *EntitlementResource) ProcessEntitlementMapForEntitlementUpdate(ctx context.Context, plan *EntitlementResourceModel, state *EntitlementResourceModel, updateReq *openapi.CreateUpdateEntitlementRequest) {
	if !plan.EntitlementMap.IsNull() || !state.EntitlementMap.IsNull() {
		var planMaps, stateMaps []EntitlementMapModel

		if !plan.EntitlementMap.IsNull() {
			plan.EntitlementMap.ElementsAs(ctx, &planMaps, false)
		}
		if !state.EntitlementMap.IsNull() {
			state.EntitlementMap.ElementsAs(ctx, &stateMaps, false)
		}

		// Create unique keys for comparison (entitlement_value + entitlement_type)
		planMapKeys := make(map[string]EntitlementMapModel)
		stateMapKeys := make(map[string]EntitlementMapModel)

		for _, entMap := range planMaps {
			key := entMap.EntitlementValue.ValueString() + "|" + entMap.EntitlementType.ValueString()
			planMapKeys[key] = entMap
		}

		for _, entMap := range stateMaps {
			key := entMap.EntitlementValue.ValueString() + "|" + entMap.EntitlementType.ValueString()
			stateMapKeys[key] = entMap
		}

		var apiEntitlementMaps []openapi.CreateUpdateEntitlementRequestEntitlementmapInner

		// Add new maps (in plan but not in state)
		for key, entMap := range planMapKeys {
			if _, exists := stateMapKeys[key]; !exists {
				apiEntitlementMaps = append(apiEntitlementMaps, r.BuildEntitlementMapAPIObject(entMap, "ADD"))
			}
		}

		// Remove old maps (in state but not in plan)
		for key, entMap := range stateMapKeys {
			if _, exists := planMapKeys[key]; !exists {
				apiEntitlementMaps = append(apiEntitlementMaps, r.BuildEntitlementMapAPIObject(entMap, "REMOVE"))
			}
		}

		// Update existing maps (in both plan and state but with different values)
		for key, planMap := range planMapKeys {
			if stateMap, exists := stateMapKeys[key]; exists {
				// Check if any boolean fields have changed
				if planMap.RequestFilter != stateMap.RequestFilter ||
					planMap.ExcludeEntitlement != stateMap.ExcludeEntitlement ||
					planMap.AddDependentTask != stateMap.AddDependentTask ||
					planMap.RemoveDependentEntTask != stateMap.RemoveDependentEntTask {

					apiEntitlementMaps = append(apiEntitlementMaps, r.BuildEntitlementMapAPIObject(planMap, "UPDATE"))
				}
			}
		}

		if len(apiEntitlementMaps) > 0 {
			updateReq.Entitlementmap = apiEntitlementMaps
		}
	}
}

// BuildEntitlementMapAPIObject creates an API object from EntitlementMapModel
func (r *EntitlementResource) BuildEntitlementMapAPIObject(entMap EntitlementMapModel, updateType string) openapi.CreateUpdateEntitlementRequestEntitlementmapInner {
	return openapi.CreateUpdateEntitlementRequestEntitlementmapInner{
		Entitlementvalue:       util.StringPointerOrEmpty(entMap.EntitlementValue),
		Entitlementtype:        util.StringPointerOrEmpty(entMap.EntitlementType),
		Endpoint:               util.StringPointerOrEmpty(entMap.Endpoint),
		Requestfilter:          util.BoolToStringPointer(entMap.RequestFilter),
		Excludeentitlement:     util.BoolToStringPointer(entMap.ExcludeEntitlement),
		Adddependenttask:       util.BoolToStringPointer(entMap.AddDependentTask),
		Removedependententtask: util.BoolToStringPointer(entMap.RemoveDependentEntTask),
		UpdateType:             util.StringPtr(updateType),
	}
}

// UpdateEntitlement updates an entitlement
func (r *EntitlementResource) UpdateEntitlement(ctx context.Context, plan *EntitlementResourceModel, state *EntitlementResourceModel) (*openapi.CreateOrUpdateEntitlementResponse, error) {
	log.Printf("[DEBUG] Entitlements: Starting updation for entitlement id: %s", plan.EntitlementValuekey.ValueString())

	updateReq := r.BuildUpdateEntitlementRequest(plan, state)
	r.ProcessEntitlementOwnersForEntitlementUpdate(ctx, plan, state, updateReq)
	r.ProcessEntitlementMapForEntitlementUpdate(ctx, plan, state, updateReq)

	// Debug logging with proper JSON marshaling
	if reqJSON, err := json.MarshalIndent(updateReq, "", "  "); err == nil {
		log.Printf("[DEBUG] Entitlements: Final update request JSON: %s", string(reqJSON))
	} else {
		log.Printf("[DEBUG] Entitlements: Final update request (struct): %+v", updateReq)
	}

	// Execute update operation with retry logic
	var updateResp *openapi.CreateOrUpdateEntitlementResponse
	var updateHttpResp *http.Response
	err := r.provider.AuthenticatedAPICallWithRetry(ctx, "update_entitlement", func(token string) error {
		entitlementOps := r.entitlementFactory.CreateEntitlementOperations(r.client.APIBaseURL(), token)
		resp, hResp, err := entitlementOps.CreateUpdateEntitlement(ctx, *updateReq)
		if hResp != nil && hResp.StatusCode == 401 {
			return fmt.Errorf("401 unauthorized")
		}
		updateResp = resp
		updateHttpResp = hResp
		return err
	})

	if err != nil {
		// Handle API call failures and HTTP errors during entitlement update
		log.Printf("[ERROR] Entitlements: Error in updating entitlement: %v", err)
		err = errorsutil.HandleHTTPError(updateHttpResp, err, "Update")
		return nil, fmt.Errorf("error updating entitlement: %v", err)
	}
	if updateResp != nil && updateResp.ErrorCode != nil && *updateResp.ErrorCode == "1" {
		msg := util.SafeDeref(updateResp.Msg)
		log.Printf("[ERROR] Entitlements: Error in updating entitlement: %v", msg)
		return nil, fmt.Errorf("error updating entitlement: %v", msg)
	}

	return updateResp, nil
}

// ReadEntitlementState reads the current state from the API for post-update operations
func (r *EntitlementResource) ReadEntitlementState(ctx context.Context, plan *EntitlementResourceModel) error {
	// Read entitlement from API using helper function
	readResp, err := r.ReadEntitlement(ctx, plan)
	if err != nil {
		// Handle failures when reading entitlement state for post-update operations
		return fmt.Errorf("error reading entitlement: %v", err)
	}

	if readResp != nil && len(readResp.Entitlementdetails) > 0 {
		entitlement := readResp.Entitlementdetails[0]

		// Populate basic fields from API response
		r.PopulateStateFromAPI(ctx, plan, entitlement)

		// Handle entitlement owners (post-update, so treat as import)
		err := r.ProcessEntitlementOwnersForEntitlementRead(ctx, plan, entitlement, true)
		if err != nil {
			// Handle failures when processing entitlement owners during post-update read
			return fmt.Errorf("error processing entitlement owners: %v", err)
		}

		// Handle entitlement map (post-update, so treat as import)
		err = r.ProcessEntitlementMapForEntitlementRead(ctx, plan, entitlement, true)
		if err != nil {
			// Handle failures when processing entitlement map during post-update read
			return fmt.Errorf("error processing entitlement map: %v", err)
		}
	}

	return nil
}

// CheckEntitlementExists checks if an entitlement already exists before creation
func (r *EntitlementResource) CheckEntitlementExists(ctx context.Context, plan *EntitlementResourceModel) error {
	log.Printf("[DEBUG] Checking if entitlement already exists: %s", plan.EntitlementValue.ValueString())

	getReq := openapi.GetEntitlementRequest{
		Endpoint:             util.StringPtr(plan.Endpoint.ValueString()),
		Entitlementtype:      util.StringPtr(plan.Entitlementtype.ValueString()),
		EntitlementValue:     util.StringPtr(plan.EntitlementValue.ValueString()),
		Entownerwithrank:     util.StringPtr("true"),
		Returnentitlementmap: util.StringPtr("true"),
	}

	log.Printf("[DEBUG] Making API call to check entitlement existence for endpoint: %s, type: %s, value: %s",
		plan.Endpoint.ValueString(), plan.Entitlementtype.ValueString(), plan.EntitlementValue.ValueString())

	// Execute import validation call with retry logic
	var existingEntitlement *openapi.GetEntitlementResponse
	err := r.provider.AuthenticatedAPICallWithRetry(ctx, "check_entitlement_exists", func(token string) error {
		entitlementOps := r.entitlementFactory.CreateEntitlementOperations(r.client.APIBaseURL(), token)
		resp, httpResp, err := entitlementOps.GetEntitlements(ctx, getReq)
		if httpResp != nil && httpResp.StatusCode == 401 {
			return fmt.Errorf("401 unauthorized")
		}
		existingEntitlement = resp
		return err
	})

	if err != nil {
		// Suppress API errors during existence check since create/update endpoints are the same
		log.Printf("[DEBUG] Could not check entitlement existence (suppressed): %v", err)
		return nil
	}
	if existingEntitlement != nil && existingEntitlement.ErrorCode != nil && *existingEntitlement.ErrorCode == "0" && len(existingEntitlement.Entitlementdetails) > 0 {
		log.Printf("[ERROR] Entitlement already exists: %s in endpoint %s", plan.EntitlementValue.ValueString(), plan.Endpoint.ValueString())
		return fmt.Errorf("Entitlement '%s' already exists in endpoint '%s' with entitlement type '%s'. Please use 'terraform import' to import the existing entitlement or use a different entitlement value",
			plan.EntitlementValue.ValueString(),
			plan.Endpoint.ValueString(),
			plan.Entitlementtype.ValueString(),
		)
	}

	log.Printf("[DEBUG] Entitlement does not exist, proceeding with creation")
	return nil
}

func (r *EntitlementResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan EntitlementResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		log.Printf("[ERROR] Failed to get plan from request: %v", resp.Diagnostics)
		return
	}

	// Check if entitlement already exists
	if err := r.CheckEntitlementExists(ctx, &plan); err != nil {
		resp.Diagnostics.AddError("Entitlement Already Exists", err.Error())
		return
	}

	// Create the entitlement using helper function
	createResp, err := r.CreateEntitlement(ctx, &plan)
	if err != nil {
		// Handle entitlement creation failures and add to Terraform diagnostics
		resp.Diagnostics.AddError("Error in creating entitlement", err.Error())
		return
	}

	// Check if EntitlementObj is not nil before accessing
	if createResp == nil || createResp.EntitlementObj == nil {
		log.Printf("[ERROR] Entitlements: Invalid response from API - EntitlementObj is nil")
		resp.Diagnostics.AddError("Invalid API Response", "Received nil EntitlementObj from API")
		return
	}

	plan.ID = util.SafeString(createResp.EntitlementObj.EntitlementValuekey)
	plan.EntitlementValuekey = util.SafeString(createResp.EntitlementObj.EntitlementValuekey)

	// Set default values for computed fields
	r.SetDefaultComputedValues(&plan)

	// Set EntitlementKey to empty for all entitlement maps since it's not available during create
	if !plan.EntitlementMap.IsNull() {
		var entitlementMaps []EntitlementMapModel
		plan.EntitlementMap.ElementsAs(ctx, &entitlementMaps, false)

		for i := range entitlementMaps {
			entitlementMaps[i].EntitlementKey = types.StringValue("")
			// Set default values for boolean fields if they are unknown
			if entitlementMaps[i].RequestFilter.IsUnknown() {
				entitlementMaps[i].RequestFilter = types.BoolValue(false)
			}
			if entitlementMaps[i].ExcludeEntitlement.IsUnknown() {
				entitlementMaps[i].ExcludeEntitlement = types.BoolValue(false)
			}
			if entitlementMaps[i].AddDependentTask.IsUnknown() {
				entitlementMaps[i].AddDependentTask = types.BoolValue(false)
			}
			if entitlementMaps[i].RemoveDependentEntTask.IsUnknown() {
				entitlementMaps[i].RemoveDependentEntTask = types.BoolValue(false)
			}
		}

		plan.EntitlementMap, _ = types.SetValueFrom(ctx, types.ObjectType{
			AttrTypes: map[string]attr.Type{
				"entitlement_value":         types.StringType,
				"entitlement_type":          types.StringType,
				"endpoint":                  types.StringType,
				"entitlement_key":           types.StringType,
				"request_filter":            types.BoolType,
				"exclude_entitlement":       types.BoolType,
				"add_dependent_task":        types.BoolType,
				"remove_dependent_ent_task": types.BoolType,
			},
		}, entitlementMaps)
	}

	stateCreateDiagnostics := resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(stateCreateDiagnostics...)
}

func (r *EntitlementResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state EntitlementResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		log.Printf("[ERROR] Failed to get plan from request: %v", resp.Diagnostics)
		return
	}

	// Check if this is an import before we modify state
	isImport := state.ID.IsNull()
	log.Printf("State: %+v", state)

	// Read entitlement from API using helper function
	readResp, err := r.ReadEntitlement(ctx, &state)
	if err != nil {
		// Handle entitlement read failures and add to Terraform diagnostics
		resp.Diagnostics.AddError("Error in reading entitlement", err.Error())
		return
	}

	if readResp != nil && len(readResp.Entitlementdetails) > 0 {
		entitlement := readResp.Entitlementdetails[0]

		// Populate basic fields from API response
		r.PopulateStateFromAPI(ctx, &state, entitlement)

		// Handle entitlement owners
		err := r.ProcessEntitlementOwnersForEntitlementRead(ctx, &state, entitlement, isImport)
		if err != nil {
			resp.Diagnostics.AddError("Entitlement Owner Drift Detected", err.Error())
			return
		}

		// Handle entitlement map
		err = r.ProcessEntitlementMapForEntitlementRead(ctx, &state, entitlement, isImport)
		if err != nil {
			resp.Diagnostics.AddError("Entitlement Map Drift Detected", err.Error())
			return
		}
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *EntitlementResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state EntitlementResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		log.Printf("[ERROR] Failed to get plan from request: %v", resp.Diagnostics)
		return
	}

	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		log.Printf("[ERROR] Failed to get state from request: %v", resp.Diagnostics)
		return
	}

	log.Printf("[DEBUG] Update - Plan: %+v", plan)
	log.Printf("[DEBUG] Update - State: %+v", state)

	// Update the entitlement using helper function
	updateResp, err := r.UpdateEntitlement(ctx, &plan, &state)
	if err != nil {
		resp.Diagnostics.AddError("Error in updating entitlement", err.Error())
		return
	}

	if updateResp == nil || updateResp.EntitlementObj == nil {
		resp.Diagnostics.AddError("Update Response Error", "API did not return entitlement object after update")
		return
	}

	plan.ID = util.SafeString(updateResp.EntitlementObj.EntitlementValuekey)
	plan.EntitlementValuekey = util.SafeString(updateResp.EntitlementObj.EntitlementValuekey)

	// Set default values for computed fields
	r.SetDefaultComputedValues(&plan)

	// Post-update read to refresh all values from API
	err = r.ReadEntitlementState(ctx, &plan)
	if err != nil {
		resp.Diagnostics.AddError("Post-Update Read Failed", err.Error())
		return
	}

	stateUpdateDiagnostics := resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(stateUpdateDiagnostics...)
}

func (r *EntitlementResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
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

func (r *EntitlementResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	log.Printf("Import key received: %s", req.ID)
	idParts := strings.Split(req.ID, ":")
	if len(idParts) != 3 {
		resp.Diagnostics.AddError(
			"Invalid Import ID Format",
			fmt.Sprintf("Expected import ID format: 'endpoint_name:entitlement_type:entitlement_value', got: %s\n"+
				"Example: terraform import saviynt_entitlement_resource.example sample-103:terraform_ent_type:terraform_ent", req.ID),
		)
		return
	}

	endpointName := strings.TrimSpace(idParts[0])
	entitlementType := strings.TrimSpace(idParts[1])
	entitlementName := strings.TrimSpace(idParts[2])

	log.Printf("Starting import for entitlement: %s, entitltment type: %s for endpoint %s", entitlementName, entitlementType, endpointName)

	if endpointName == "" || entitlementType == "" || entitlementName == "" {
		resp.Diagnostics.AddError(
			"Invalid Import ID Components",
			"endpoint_name, entitlement_type and entitlement_name must be non-empty\n"+
				"Example: terraform import saviynt_entitlement_resource.example sample-103:terraform_ent_type:terraform_ent",
		)
		return
	}

	// Set both endpoint_name and entitlement_name in the state
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("endpoint"), endpointName)...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("entitlement_type"), entitlementType)...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("entitlement_value"), entitlementName)...)
}
