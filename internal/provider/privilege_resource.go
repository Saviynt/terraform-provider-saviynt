// Copyright (c) 2025 Saviynt Inc.
// SPDX-License-Identifier: MPL-2.0

// saviynt_privilege_resource manages privileges in the Saviynt Security Manager.
// The resource implements the full Terraform lifecycle:
//   - Create: provisions new privileges for a privilege using the supplied configuration.
//   - Read: fetches the current privilege state for a privilege from Saviynt to keep Terraform’s state in sync.
//   - Update: applies any configuration changes to existing privileges.
//.  - Delete: deletes the privileges for a given endpoint and entitlement type.
//   - Import: brings existing privileges under Terraform management by its name.

package provider

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"terraform-provider-Saviynt/internal/client"
	"terraform-provider-Saviynt/util"
	"terraform-provider-Saviynt/util/errorsutil"
	"terraform-provider-Saviynt/util/privilegeutil"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"

	"github.com/hashicorp/terraform-plugin-framework/types"

	endpoint "github.com/saviynt/saviynt-api-go-client/endpoints"
	openapi "github.com/saviynt/saviynt-api-go-client/privileges"
)

var _ resource.Resource = &PrivilegeResource{}
var _ resource.ResourceWithImportState = &PrivilegeResource{}

type PrivilegeResource struct {
	client           client.SaviyntClientInterface
	token            string
	provider         client.SaviyntProviderInterface
	privilegeFactory client.PrivilegeFactoryInterface
}

func NewPrivilegeResource() resource.Resource {
	return &PrivilegeResource{
		privilegeFactory: &client.DefaultPrivilegeFactory{},
	}
}

func NewPrivilegeResourceWithFactory(factory client.PrivilegeFactoryInterface) resource.Resource {
	return &PrivilegeResource{
		privilegeFactory: factory,
	}
}

type PrivilegeResourceModel struct {
	ID              types.String `tfsdk:"id"`
	SecuritySystem  types.String `tfsdk:"security_system"`
	Endpoint        types.String `tfsdk:"endpoint"`
	EntitlementType types.String `tfsdk:"entitlement_type"`
	Privileges      types.Map    `tfsdk:"privileges"`
}

type Privilege struct {
	AttributeName   types.String `tfsdk:"attribute_name"`
	AttributeType   types.String `tfsdk:"attribute_type"`
	OrderIndex      types.String `tfsdk:"order_index"`
	DefaultValue    types.String `tfsdk:"default_value"`
	AttributeConfig types.String `tfsdk:"attribute_config"`
	Label           types.String `tfsdk:"label"`
	AttributeGroup  types.String `tfsdk:"attribute_group"`
	ParentAttribute types.String `tfsdk:"parent_attribute"`
	ChildAction     types.String `tfsdk:"child_action"`
	Description     types.String `tfsdk:"description"`
	Required        types.Bool   `tfsdk:"required"`
	Requestable     types.Bool   `tfsdk:"requestable"`
	HideOnCreate    types.Bool   `tfsdk:"hide_on_create"`
	HideOnUpdate    types.Bool   `tfsdk:"hide_on_update"`
	ActionString    types.String `tfsdk:"action_string"`
}

func (r *PrivilegeResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = "saviynt_privilege_resource"
}

func (r *PrivilegeResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: util.PrivilegeDescription,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "Unique identifier of the privilege resource block. This is automatically generated as 'priv-' + endpoint + '-' + entitlementtype",
			},
			"security_system": schema.StringAttribute{
				Required:    true,
				Description: "Name of the security system to which the endpoint belongs",
			},
			"endpoint": schema.StringAttribute{
				Required:    true,
				Description: "Name of the endpoint to which the entitlement type belongs",
			},
			"entitlement_type": schema.StringAttribute{
				Required:    true,
				Description: "Name of the entitlement type for the privilege",
			},
			"privileges": schema.MapNestedAttribute{
				Required:    true,
				Description: "Map of privileges to create/manage",
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"attribute_name": schema.StringAttribute{
							Required:    true,
							Description: "Attribute name for the privilege",
						},
						"attribute_type": schema.StringAttribute{
							Required:    true,
							Description: "Type of the attribute/privilege",
						},
						"order_index": schema.StringAttribute{
							Optional:    true,
							Computed:    true,
							Description: "Order index",
						},
						"default_value": schema.StringAttribute{
							Optional:    true,
							Computed:    true,
							Description: "Default value for the privilege",
						},
						"attribute_config": schema.StringAttribute{
							Required:    true,
							Description: "Configuration type for the attribute",
						},
						"label": schema.StringAttribute{
							Optional:    true,
							Computed:    true,
							Description: "Label for the privilege",
						},
						"attribute_group": schema.StringAttribute{
							Optional:    true,
							Computed:    true,
							Description: "Attribute group",
						},
						"parent_attribute": schema.StringAttribute{
							Optional:    true,
							Computed:    true,
							Description: "Parent attribute for the given privilege",
						},
						"child_action": schema.StringAttribute{
							Optional:    true,
							Computed:    true,
							Description: "Child action",
						},
						"description": schema.StringAttribute{
							Optional:    true,
							Computed:    true,
							Description: "Description for the privilege",
						},
						"required": schema.BoolAttribute{
							Optional:    true,
							Computed:    true,
							Default:     booldefault.StaticBool(false),
							Description: "Is required",
						},
						"requestable": schema.BoolAttribute{
							Optional:    true,
							Computed:    true,
							Default:     booldefault.StaticBool(false),
							Description: "Is requestable",
						},
						"hide_on_create": schema.BoolAttribute{
							Optional:    true,
							Computed:    true,
							Description: "Hide on create",
						},
						"hide_on_update": schema.BoolAttribute{
							Optional:    true,
							Computed:    true,
							Description: "Hide on update",
						},
						"action_string": schema.StringAttribute{
							Optional:    true,
							Computed:    true,
							Description: "Action string",
						},
					},
				},
			},
		},
	}
}

func (r *PrivilegeResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Check if provider data is available.
	if req.ProviderData == nil {
		log.Println("ProviderData is nil, returning early.")
		return
	}

	// Cast provider data to your provider type.
	prov, ok := req.ProviderData.(*SaviyntProvider)
	if !ok {
		log.Println("[ERROR] Provider: Unexpected provider data")
		resp.Diagnostics.AddError("Unexpected Provider Data", "Expected *SaviyntProvider")
		return
	}

	// Set the client and token from the provider state using interface wrapper.
	r.client = &client.SaviyntClientWrapper{Client: prov.client}
	r.token = prov.accessToken
	r.provider = &client.SaviyntProviderWrapper{Provider: prov} // Store provider reference for retry logic
	log.Println("[DEBUG] Privilege: Resource configured successfully")
}

// SetClient sets the client for testing purposes
func (r *PrivilegeResource) SetClient(client client.SaviyntClientInterface) {
	r.client = client
}

// SetToken sets the token for testing purposes
func (r *PrivilegeResource) SetToken(token string) {
	r.token = token
}

// SetProvider sets the provider for testing purposes
func (r *PrivilegeResource) SetProvider(provider client.SaviyntProviderInterface) {
	r.provider = provider
}

// ReadPrivilege reads privileges for a specific endpoint and entitlement type
func (r *PrivilegeResource) ReadPrivilege(ctx context.Context, endpointName, entitlementType string) (*openapi.GetPrivilegeListResponse, error) {
	log.Printf("[DEBUG] Privilege: Starting read operation for endpoint: %s, entitlement type: %s", endpointName, entitlementType)

	getReq := openapi.GetPrivilegeListRequest{
		Endpoint:        endpointName,
		Entitlementtype: &entitlementType,
	}

	getReqJson, _ := json.Marshal(getReq)
	log.Printf("[DEBUG] Privilege get request: %s", getReqJson)

	log.Printf("[DEBUG] Privilege: Making API call to fetch privileges for endpoint: %s", endpointName)
	var fetchResp *openapi.GetPrivilegeListResponse
	var finalHttpResp *http.Response
	err := r.provider.AuthenticatedAPICallWithRetry(ctx, "get_privilege", func(token string) error {
		privilegeOps := r.privilegeFactory.CreatePrivilegeOperations(r.client.APIBaseURL(), token)
		resp, httpResp, err := privilegeOps.GetPrivilege(ctx, getReq)
		if httpResp != nil && httpResp.StatusCode == 401 {
			return fmt.Errorf("401 unauthorized")
		}
		fetchResp = resp
		finalHttpResp = httpResp // Capture final HTTP response
		return err
	})

	if err != nil {
		log.Printf("[ERROR] Privilege: Problem with the get function in ReadPrivilege. Error: %v", err)
		err = errorsutil.HandleHTTPError(finalHttpResp, err, "Read")
		return nil, fmt.Errorf("fetch API call failed: %w", err)
	}

	getRespJson, _ := json.Marshal(fetchResp)
	log.Printf("[DEBUG] Privilege get response: %s", getRespJson)

	// Check for API-level errors in the response
	if fetchResp != nil {
		if err := r.ProcessPrivilegeListErrorResponse(fetchResp); err != nil {
			return nil, err
		}
	}

	log.Printf("[INFO] Privilege resource read successfully. Endpoint: %s", endpointName)
	return fetchResp, nil
}

// UpdateModelFromReadResponse updates the model from read response
func (r *PrivilegeResource) UpdateModelFromReadResponse(state *PrivilegeResourceModel, apiResp *openapi.GetPrivilegeListResponse, ctx context.Context) error {
	log.Printf("[DEBUG] Read API Response: %+v", apiResp)
	isImport := state.ID.IsNull() || state.ID.IsUnknown()

	// Get current state privileges as a map
	currentPrivs := make(map[string]Privilege)
	if !state.Privileges.IsNull() {
		diags := state.Privileges.ElementsAs(ctx, &currentPrivs, false)
		if diags.HasError() {
			return fmt.Errorf("failed to get current privileges: %v", diags.Errors())
		}
	}

	// Create map to store API privileges by name
	apiPrivs := make(map[string]openapi.GetPrivilegeDetail)
	if apiResp != nil && apiResp.PrivilegeDetails != nil {
		for _, item := range apiResp.PrivilegeDetails {
			privName := util.SafeDeref(item.Attribute)
			if privName != "" {
				apiPrivs[privName] = item
			}
		}
	}

	// Build updated privileges map - only include privileges that are in BOTH state and API
	updatedPrivs := make(map[string]Privilege)
	removedPrivileges := make([]string, 0)

	// if len(currentPrivs) == 0 {
	if isImport {
		// IMPORT CASE: State is empty, bring in all privileges from API
		updatedPrivs = r.BuildUpdatedPrivilegesForImport(apiPrivs)
	} else {
		// REGULAR READ CASE: Only keep privileges that exist in both state and API
		for privName := range currentPrivs {
			if apiPriv, exists := apiPrivs[privName]; exists {
				// Privilege exists in both state and API - update with API values
				updatedPrivs[privName] = r.BuildUpdatedPrivilegeFromAPI(privName, apiPriv, false)
			} else {
				// Privilege exists in state but not in API - remove from state
				removedPrivileges = append(removedPrivileges, privName)
			}
		}
	}

	// Log removed privileges
	if len(removedPrivileges) > 0 {
		log.Printf("[INFO] Removing privileges not present in API response: %v", removedPrivileges)
	}

	// Convert to types.Map
	privilegesMap, diags := r.ConvertPrivilegesToTerraformMap(updatedPrivs, ctx)
	if diags.HasError() {
		return fmt.Errorf("failed to convert privileges to map: %v", diags.Errors())
	}

	state.Privileges = privilegesMap
	state.ID = types.StringValue("priv-" + state.Endpoint.ValueString() + "-" + state.EntitlementType.ValueString())

	log.Printf("[INFO] Privilege: Successfully read %d privileges for endpoint: %s", len(updatedPrivs), state.Endpoint.ValueString())
	return nil
}

// BuildUpdatedPrivilegesForImport builds updated privileges for import case
func (r *PrivilegeResource) BuildUpdatedPrivilegesForImport(apiPrivs map[string]openapi.GetPrivilegeDetail) map[string]Privilege {
	updatedPrivs := make(map[string]Privilege)
	for privName, apiPriv := range apiPrivs {
		updatedPrivs[privName] = r.BuildUpdatedPrivilegeFromAPI(privName, apiPriv, true)
	}
	log.Printf("[IMPORT] Importing all %d privileges from API", len(apiPrivs))
	return updatedPrivs
}

// BuildUpdatedPrivilegeFromAPI builds updated privilege from API response
func (r *PrivilegeResource) BuildUpdatedPrivilegeFromAPI(privName string, apiPriv openapi.GetPrivilegeDetail, isImport bool) Privilege {
	// Translate attribute type from API format to Terraform format
	attributeType := util.SafeDeref(apiPriv.AttributeType)
	translatedAttributeType := privilegeutil.TranslateValue(attributeType, privilegeutil.AttributeTypeMap)

	return Privilege{
		AttributeName:   types.StringValue(privName),
		AttributeType:   types.StringValue(translatedAttributeType),
		OrderIndex:      util.SafeString(apiPriv.Orderindex),
		DefaultValue:    util.SafeString(apiPriv.Defaultvalue),
		AttributeConfig: util.SafeString(apiPriv.AttributeConfig),
		Label:           util.SafeString(apiPriv.Label),
		AttributeGroup:  util.SafeString(apiPriv.Attributegroup),
		ParentAttribute: util.SafeString(apiPriv.Parentattribute),
		ChildAction:     util.SafeString(apiPriv.Childaction),
		Description:     util.SafeString(apiPriv.Descriptionascsv),
		Required:        util.SafeBoolDatasource(apiPriv.Required),
		Requestable:     util.SafeBoolDatasource(apiPriv.Requestablerequired),
		HideOnCreate:    util.SafeBoolDatasource(apiPriv.Hideoncreate),
		HideOnUpdate:    util.SafeBoolDatasource(apiPriv.Hideonupd),
		ActionString:    util.SafeString(apiPriv.ActionString),
	}
}

// ConvertPrivilegesToTerraformMap converts privileges map to Terraform Map type
func (r *PrivilegeResource) ConvertPrivilegesToTerraformMap(updatedPrivs map[string]Privilege, ctx context.Context) (types.Map, diag.Diagnostics) {
	privilegesMap, diags := types.MapValueFrom(ctx, types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"attribute_name":   types.StringType,
			"attribute_type":   types.StringType,
			"order_index":      types.StringType,
			"default_value":    types.StringType,
			"attribute_config": types.StringType,
			"label":            types.StringType,
			"attribute_group":  types.StringType,
			"parent_attribute": types.StringType,
			"child_action":     types.StringType,
			"description":      types.StringType,
			"required":         types.BoolType,
			"requestable":      types.BoolType,
			"hide_on_create":   types.BoolType,
			"hide_on_update":   types.BoolType,
			"action_string":    types.StringType,
		},
	}, updatedPrivs)

	return privilegesMap, diags
}

// ProcessDeletePrivilegeErrorResponse processes error response from delete privilege operations
func (r *PrivilegeResource) ProcessDeletePrivilegeErrorResponse(resp *openapi.DeletePrivilegeResponse, privilegeName string) error {
	if resp == nil || resp.Errorcode == nil {
		return fmt.Errorf("unexpected API response - no error code provided")
	}

	if *resp.Errorcode != 0 {
		var errorMessages []string

		// Add main message
		if resp.Msg != nil {
			errorMessages = append(errorMessages, fmt.Sprintf("message: %s", *resp.Msg))
		}

		// Handle different error field types
		if resp.Securitysystem != nil {
			errorMessages = append(errorMessages, fmt.Sprintf("securitysystem: %s", *resp.Securitysystem))
		}

		if resp.Endpoint != nil {
			errorMessages = append(errorMessages, fmt.Sprintf("endpoint: %s", *resp.Endpoint))
		}

		if resp.Entitlementtype != nil {
			errorMessages = append(errorMessages, fmt.Sprintf("entitlementtype: %s", *resp.Entitlementtype))
		}

		// Handle privilege field
		if resp.Privilege != nil {
			rawJSON, err := json.Marshal(resp.Privilege)
			if err != nil {
				log.Printf("Error marshaling privilege error: %v", err)
				errorMessages = append(errorMessages, "privilege: error processing privilege details")
			} else {
				// Try as string first
				var str string
				if err := json.Unmarshal(rawJSON, &str); err == nil {
					errorMessages = append(errorMessages, fmt.Sprintf("privilege: %s", str))
				} else {
					// Try as object with nested fields
					var privMap map[string]interface{}
					if err := json.Unmarshal(rawJSON, &privMap); err == nil {
						for field, msg := range privMap {
							errorMessages = append(errorMessages, fmt.Sprintf("privilege.%s: %v", field, msg))
						}
					} else {
						errorMessages = append(errorMessages, "privilege: unknown error format")
					}
				}
			}
		}

		fullError := strings.Join(errorMessages, "\n")
		return fmt.Errorf("Privilege Delete Operation Failed for '%s':\n%s", privilegeName, fullError)
	}

	return nil
}

// ProcessPrivilegeErrorResponse processes error response from privilege operations
func (r *PrivilegeResource) ProcessPrivilegeErrorResponse(resp *openapi.CreateUpdatePrivilegeResponse, operationType string, privilegeName string) error {
	if resp == nil || resp.Errorcode == nil {
		return fmt.Errorf("unexpected API response - no error code provided")
	}

	if *resp.Errorcode != 0 {
		// Show full API response for debugging
		respJSON, _ := json.Marshal(resp)
		return fmt.Errorf("Privilege %s Operation Failed for '%s': Full API Response: %s", operationType, privilegeName, string(respJSON))
	}

	return nil
}

// ProcessPrivilegeListErrorResponse processes error response from get privilege list operation
func (r *PrivilegeResource) ProcessPrivilegeListErrorResponse(resp *openapi.GetPrivilegeListResponse) error {
	if resp == nil || resp.ErrorCode == nil {
		return fmt.Errorf("unexpected API response - no error code provided")
	}

	if *resp.ErrorCode != "0" {
		message := "no message found"

		// Add main message
		if resp.Msg != nil {
			message = *resp.Msg
		}

		return fmt.Errorf("Privilege List Fetch Failed:\n error code: %s, message: %s", *resp.ErrorCode, message)
	}

	return nil
}
func (r *PrivilegeResource) UpdatePrivilege(ctx context.Context, plan *PrivilegeResourceModel) (map[string]Privilege, []string) {
	log.Printf("[DEBUG] Privilege: Starting update operation for endpoint: %s", plan.Endpoint.ValueString())

	// Get planned privileges
	var planPrivs map[string]Privilege
	diags := plan.Privileges.ElementsAs(ctx, &planPrivs, false)
	if diags.HasError() {
		log.Printf("[ERROR] Privilege: Failed to process planned privileges: %v", diags.Errors())
		return nil, []string{fmt.Sprintf("failed to process planned privileges: %v", diags.Errors())}
	}

	// Get current state from server
	fetchResp, err := r.ReadPrivilege(ctx, plan.Endpoint.ValueString(), plan.EntitlementType.ValueString())
	if err != nil {
		log.Printf("[ERROR] Privilege: Failed to fetch current privileges: %v", err)
		return nil, []string{fmt.Sprintf("failed to fetch current privileges: %v", err)}
	}

	// Build map of existing privileges
	existingPrivs := make(map[string]bool)
	if fetchResp != nil && fetchResp.PrivilegeDetails != nil {
		for _, item := range fetchResp.PrivilegeDetails {
			if item.Attribute != nil {
				existingPrivs[*item.Attribute] = true
			}
		}
	}

	var errors []string
	successfulPrivileges := make(map[string]Privilege)

	// Process each planned privilege
	for _, priv := range planPrivs {
		privName := priv.AttributeName.ValueString()

		privilege := openapi.CreateUpdatePrivilegeRequestPrivilege{
			Attributename:   priv.AttributeName.ValueStringPointer(),
			Attributetype:   priv.AttributeType.ValueStringPointer(),
			Attributeconfig: priv.AttributeConfig.ValueStringPointer(),
			Orderindex:      util.StringPointerOrEmpty(priv.OrderIndex),
			Defaultvalue:    util.StringPointerOrEmpty(priv.DefaultValue),
			Label:           util.StringPointerOrEmpty(priv.Label),
			Attributegroup:  util.StringPointerOrEmpty(priv.AttributeGroup),
			Parentattribute: util.StringPointerOrEmpty(priv.ParentAttribute),
			Childaction:     util.StringPointerOrEmpty(priv.ChildAction),
			Description:     util.StringPointerOrEmpty(priv.Description),
			Actionstring:    util.StringPointerOrEmpty(priv.ActionString),
			Required:        util.BoolPointerOrEmpty(priv.Required),
			Requestable:     util.BoolPointerOrEmpty(priv.Requestable),
			Hideoncreate:    util.BoolPointerOrEmpty(priv.HideOnCreate),
			Hideonupd:       util.BoolPointerOrEmpty(priv.HideOnUpdate),
		}

		req := openapi.CreateUpdatePrivilegeRequest{
			Securitysystem:  plan.SecuritySystem.ValueString(),
			Endpoint:        plan.Endpoint.ValueString(),
			Entitlementtype: plan.EntitlementType.ValueString(),
			Privilege:       privilege,
		}

		if !existingPrivs[privName] {
			// Create new privilege
			log.Printf("[DEBUG] Privilege: Creating new privilege: %s", privName)
			reqJSON, _ := json.Marshal(req)
			log.Printf("[DEBUG] Privilege CREATE API REQUEST for '%s': %s", privName, string(reqJSON))

			var createResp *openapi.CreateUpdatePrivilegeResponse
			var finalHttpResp *http.Response
			err := r.provider.AuthenticatedAPICallWithRetry(ctx, "create_privilege", func(token string) error {
				privilegeOps := r.privilegeFactory.CreatePrivilegeOperations(r.client.APIBaseURL(), token)
				resp, httpResp, err := privilegeOps.CreatePrivilege(ctx, req)
				if httpResp != nil && httpResp.StatusCode == 401 {
					return fmt.Errorf("401 unauthorized")
				}
				createResp = resp
				finalHttpResp = httpResp // Capture final HTTP response
				return err
			})

			respJSON, _ := json.Marshal(createResp)
			log.Printf("[DEBUG] Privilege CREATE API RESPONSE for '%s': %s", privName, string(respJSON))
			if err != nil {
				if finalHttpResp != nil {
					var respBody map[string]interface{}
					json.NewDecoder(finalHttpResp.Body).Decode(&respBody)
					errors = append(errors, fmt.Sprintf("Privilege '%s': %v", privName, respBody))
				} else {
					errors = append(errors, fmt.Sprintf("Privilege '%s': %v", privName, err))
				}
				continue
			}
			err = r.ProcessPrivilegeErrorResponse(createResp, "create", privName)
			if err != nil {
				errors = append(errors, fmt.Sprintf("Privilege '%s': %v", privName, err))
				continue
			}
		} else {
			// Update existing privilege
			log.Printf("[DEBUG] Privilege: Updating existing privilege: %s", privName)
			reqJSON, _ := json.Marshal(req)
			log.Printf("[DEBUG] Privilege UPDATE API REQUEST for '%s': %s", privName, string(reqJSON))

			var updateResp *openapi.CreateUpdatePrivilegeResponse
			var finalHttpResp *http.Response
			err := r.provider.AuthenticatedAPICallWithRetry(ctx, "update_privilege", func(token string) error {
				privilegeOps := r.privilegeFactory.CreatePrivilegeOperations(r.client.APIBaseURL(), token)
				resp, httpResp, err := privilegeOps.UpdatePrivilege(ctx, req)
				if httpResp != nil && httpResp.StatusCode == 401 {
					return fmt.Errorf("401 unauthorized")
				}
				updateResp = resp
				finalHttpResp = httpResp // Capture final HTTP response
				return err
			})

			respJSON, _ := json.Marshal(updateResp)
			log.Printf("[DEBUG] Privilege UPDATE API RESPONSE for '%s': %s", privName, string(respJSON))
			if err != nil {
				if finalHttpResp != nil {
					var respBody map[string]interface{}
					json.NewDecoder(finalHttpResp.Body).Decode(&respBody)
					errors = append(errors, fmt.Sprintf("Privilege '%s': %v", privName, respBody))
				} else {
					errors = append(errors, fmt.Sprintf("Privilege '%s': %v", privName, err))
				}
				continue
			}
			err = r.ProcessPrivilegeErrorResponse(updateResp, "update", privName)
			if err != nil {
				errors = append(errors, fmt.Sprintf("Privilege '%s': %v", privName, err))
				continue
			}
		}

		// Add to successful privileges
		successfulPrivileges[privName] = priv
	}

	if len(successfulPrivileges) > 0 {
		log.Printf("[INFO] Privilege: Successfully updated %d out of %d privileges for endpoint: %s",
			len(successfulPrivileges), len(planPrivs), plan.Endpoint.ValueString())
	}

	return successfulPrivileges, errors
}

// UpdateModelFromUpdateResponse updates the model from update response
func (r *PrivilegeResource) UpdateModelFromUpdateResponse(plan *PrivilegeResourceModel, planPrivs map[string]Privilege, ctx context.Context) error {
	// Fetch updated state after update
	postUpdateFetchResp, err := r.ReadPrivilege(ctx, plan.Endpoint.ValueString(), plan.EntitlementType.ValueString())
	if err != nil {
		return fmt.Errorf("failed to fetch updated privileges: %w", err)
	}

	postUpdateApiPrivs := make(map[string]openapi.GetPrivilegeDetail)
	if postUpdateFetchResp != nil && postUpdateFetchResp.PrivilegeDetails != nil {
		for _, item := range postUpdateFetchResp.PrivilegeDetails {
			privName := util.SafeDeref(item.Attribute)
			if privName != "" {
				postUpdateApiPrivs[privName] = item
			}
		}
	}

	updatedPrivs := make(map[string]Privilege)
	for privName := range planPrivs {
		if apiPriv, exists := postUpdateApiPrivs[privName]; exists {
			updatedPrivs[privName] = r.BuildUpdatedPrivilegeFromAPI(privName, apiPriv, false)
		}
	}

	// Convert to types.Map
	privilegesMap, diags := r.ConvertPrivilegesToTerraformMap(updatedPrivs, ctx)
	if diags.HasError() {
		return fmt.Errorf("failed to convert updated privileges to map: %v", diags.Errors())
	}

	plan.Privileges = privilegesMap
	plan.ID = types.StringValue("priv-" + plan.Endpoint.ValueString() + "-" + plan.EntitlementType.ValueString())

	return nil
}

func (r *PrivilegeResource) UpdatePrivilegeDelete(ctx context.Context, plan *PrivilegeResourceModel, state *PrivilegeResourceModel) (map[string]Privilege, []string) {
	log.Printf("[DEBUG] Privilege: Starting update-delete operation for endpoint: %s", plan.Endpoint.ValueString())

	// Get planned privileges
	var planPrivs map[string]Privilege
	diags := plan.Privileges.ElementsAs(ctx, &planPrivs, false)
	if diags.HasError() {
		log.Printf("[ERROR] Privilege: Failed to process planned privileges: %v", diags.Errors())
		return nil, []string{fmt.Sprintf("failed to process planned privileges: %v", diags.Errors())}
	}

	var statePrivs map[string]Privilege
	diags = state.Privileges.ElementsAs(ctx, &statePrivs, false)
	if diags.HasError() {
		log.Printf("[ERROR] Privilege: Failed to process state privileges: %v", diags.Errors())
		return nil, []string{fmt.Sprintf("failed to process state privileges: %v", diags.Errors())}
	}

	existingPrivs := make(map[string]bool)
	for _, priv := range planPrivs {
		if namePtr := priv.AttributeName.ValueStringPointer(); namePtr != nil {
			existingPrivs[*namePtr] = true
		}
	}

	var errors []string
	deletedPrivileges := make(map[string]Privilege)

	if len(statePrivs) >= len(planPrivs) {
		for _, priv := range statePrivs {
			privName := priv.AttributeName.ValueString()
			if !existingPrivs[privName] {
				if !priv.AttributeName.IsNull() && !priv.AttributeName.IsUnknown() {

					deleteReq := openapi.DeletePrivilegeRequest{
						Securitysystem:  state.SecuritySystem.ValueString(),
						Endpoint:        state.Endpoint.ValueString(),
						Entitlementtype: state.EntitlementType.ValueString(),
						Privilege:       privName,
					}

					deleteReqJson, _ := json.Marshal(deleteReq)
					log.Printf("[DEBUG] UpdatePrivilegeDelete: Delete request for %s: %s", privName, string(deleteReqJson))

					log.Printf("[DEBUG] Privilege: Deleting privilege: %s", privName)
					var deleteResp *openapi.DeletePrivilegeResponse
					err := r.provider.AuthenticatedAPICallWithRetry(ctx, "delete_privilege", func(token string) error {
						privilegeOps := r.privilegeFactory.CreatePrivilegeOperations(r.client.APIBaseURL(), token)
						resp, httpResp, err := privilegeOps.DeletePrivilege(ctx, deleteReq)
						if httpResp != nil && httpResp.StatusCode == 401 {
							return fmt.Errorf("401 unauthorized")
						}
						deleteResp = resp
						return err
					})
					if err != nil {
						log.Printf("[ERROR] Privilege: Problem deleting privilege: %s. Error: %v", privName, err)
						return nil, errors
					}

					err = r.ProcessDeletePrivilegeErrorResponse(deleteResp, privName)
					if err != nil {
						return nil, []string{err.Error()}
					}

					deleteRespJson, _ := json.Marshal(deleteResp)
					log.Printf("[DEBUG] UpdatePrivilegeDelete: Delete response for %s: %s", privName, string(deleteRespJson))

					log.Printf("[INFO] Privilege: Successfully deleted privilege: %s", privName)
				}
				deletedPrivileges[privName] = priv
			}
		}
	}

	return deletedPrivileges, errors
}

// ValidatePrivilegeMapKeys validates that map keys match attribute_name values
func (r *PrivilegeResource) ValidatePrivilegeMapKeys(ctx context.Context, plan *PrivilegeResourceModel) error {
	var privilegeMap map[string]Privilege
	diags := plan.Privileges.ElementsAs(ctx, &privilegeMap, false)
	if diags.HasError() {
		return fmt.Errorf("failed to convert privileges map")
	}

	for mapKey, privilege := range privilegeMap {
		if mapKey != privilege.AttributeName.ValueString() {
			return fmt.Errorf("map key '%s' does not match attribute_name '%s'", mapKey, privilege.AttributeName.ValueString())
		}
	}
	return nil
}

func (r *PrivilegeResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan PrivilegeResourceModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		log.Println("[ERROR] Privilege: Error getting plan data")
		return
	}

	// Validate that map keys match attribute_name values
	if err := r.ValidatePrivilegeMapKeys(ctx, &plan); err != nil {
		resp.Diagnostics.AddError("Validation Error", err.Error())
		return
	}

	// Use the same logic as Update - check if privileges exist and create/update accordingly
	successfulPrivileges, errors := r.UpdatePrivilege(ctx, &plan)

	// Create new state with only successful privileges
	var state PrivilegeResourceModel
	state.SecuritySystem = plan.SecuritySystem
	state.Endpoint = plan.Endpoint
	state.EntitlementType = plan.EntitlementType
	state.ID = types.StringValue("priv-" + plan.Endpoint.ValueString() + "-" + plan.EntitlementType.ValueString())

	if len(successfulPrivileges) > 0 {
		err := r.UpdateModelFromUpdateResponse(&state, successfulPrivileges, ctx)
		if err != nil {
			resp.Diagnostics.AddError("Failed to update model from response", err.Error())
			return
		}
		log.Printf("[DEBUG] Privilege: Set state for %d successful privileges", len(successfulPrivileges))
	} else {
		// Initialize with empty map to avoid type conversion errors
		emptyPrivileges := make(map[string]Privilege)
		emptyMap, diags := r.ConvertPrivilegesToTerraformMap(emptyPrivileges, ctx)
		if diags.HasError() {
			resp.Diagnostics.AddError("Failed to create empty privileges map", diags.Errors()[0].Summary())
			return
		}
		state.Privileges = emptyMap
	}

	// Set the state with successful privileges
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)

	// Fail at the end if there were errors
	if len(errors) > 0 {
		errorMsg := "Failed to create these privileges:\n"
		for _, err := range errors {
			errorMsg += "• " + err + "\n"
		}
		resp.Diagnostics.AddError("Privileges failed to create", errorMsg)
		return
	}
}

func (r *PrivilegeResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state PrivilegeResourceModel

	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	apiResp, err := r.ReadPrivilege(ctx, state.Endpoint.ValueString(), state.EntitlementType.ValueString())
	if err != nil {
		log.Printf("[ERROR] Privilege: API Read Failed In Read Block: %v", err)
		resp.Diagnostics.AddError("API Read Failed In Read Block", err.Error())
		return
	}

	err = r.UpdateModelFromReadResponse(&state, apiResp, ctx)
	if err != nil {
		log.Printf("[ERROR] Privilege: Failed to update model from read response: %v", err)
		resp.Diagnostics.AddError("Failed to update model from read response", err.Error())
		return
	}

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
}

func (r *PrivilegeResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state PrivilegeResourceModel

	// Get the plan
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Validate that map keys match attribute_name values
	if err := r.ValidatePrivilegeMapKeys(ctx, &plan); err != nil {
		resp.Diagnostics.AddError("Validation Error", err.Error())
		return
	}

	_, deleteErrors := r.UpdatePrivilegeDelete(ctx, &plan, &state)
	successfulPrivileges, errors := r.UpdatePrivilege(ctx, &plan)
	errors = append(errors, deleteErrors...)

	// Create new state with only successful privileges
	var newState PrivilegeResourceModel
	newState.SecuritySystem = plan.SecuritySystem
	newState.Endpoint = plan.Endpoint
	newState.EntitlementType = plan.EntitlementType
	newState.ID = types.StringValue("priv-" + plan.Endpoint.ValueString() + "-" + plan.EntitlementType.ValueString())

	if len(successfulPrivileges) > 0 {
		err := r.UpdateModelFromUpdateResponse(&newState, successfulPrivileges, ctx)
		if err != nil {
			resp.Diagnostics.AddError("Failed to update model from response", err.Error())
			return
		}
		log.Printf("[DEBUG] Privilege: Set state for %d successful privileges", len(successfulPrivileges))
	} else {
		// Initialize with empty map to avoid type conversion errors
		emptyPrivileges := make(map[string]Privilege)
		emptyMap, diags := r.ConvertPrivilegesToTerraformMap(emptyPrivileges, ctx)
		if diags.HasError() {
			resp.Diagnostics.AddError("Failed to create empty privileges map", diags.Errors()[0].Summary())
			return
		}
		newState.Privileges = emptyMap
	}

	// Set the state with successful privileges
	resp.Diagnostics.Append(resp.State.Set(ctx, &newState)...)

	// Fail at the end if there were errors
	if len(errors) > 0 {
		errorMsg := "Failed to update some privileges:\n"
		for _, err := range errors {
			errorMsg += "• " + err + "\n"
		}
		resp.Diagnostics.AddError("Some privileges failed to update", errorMsg)
		return
	}
}

// DeletePrivilege deletes privileges
func (r *PrivilegeResource) DeletePrivilege(ctx context.Context, state *PrivilegeResourceModel) error {
	log.Printf("[DEBUG] Privilege: Starting delete operation for endpoint: %s", state.Endpoint.ValueString())

	// Get privilege names from the map
	var statePrivs map[string]Privilege
	diags := state.Privileges.ElementsAs(ctx, &statePrivs, false)
	if diags.HasError() {
		return fmt.Errorf("failed to process state privileges: %v", diags.Errors())
	}

	for _, priv := range statePrivs {
		if !priv.AttributeName.IsNull() && !priv.AttributeName.IsUnknown() {
			privName := priv.AttributeName.ValueString()

			deleteReq := openapi.DeletePrivilegeRequest{
				Securitysystem:  state.SecuritySystem.ValueString(),
				Endpoint:        state.Endpoint.ValueString(),
				Entitlementtype: state.EntitlementType.ValueString(),
				Privilege:       privName,
			}

			deleteReqJson, _ := json.Marshal(deleteReq)
			log.Printf("[DEBUG] DeletePrivilege: Delete request for %s: %s", privName, string(deleteReqJson))

			log.Printf("[DEBUG] Privilege: Deleting privilege: %s", privName)

			var deleteResp *openapi.DeletePrivilegeResponse
			err := r.provider.AuthenticatedAPICallWithRetry(ctx, "delete_privilege", func(token string) error {
				privilegeOps := r.privilegeFactory.CreatePrivilegeOperations(r.client.APIBaseURL(), token)
				resp, httpResp, err := privilegeOps.DeletePrivilege(ctx, deleteReq)
				if httpResp != nil && httpResp.StatusCode == 401 {
					return fmt.Errorf("401 unauthorized")
				}
				deleteResp = resp
				return err
			})
			if err != nil {
				log.Printf("[ERROR] Privilege: Problem deleting privilege: %s. Error: %v", privName, err)
				return fmt.Errorf("failed to delete privilege %s: %w", privName, err)
			}

			err = r.ProcessDeletePrivilegeErrorResponse(deleteResp, privName)
			if err != nil {
				return err
			}

			deleteRespJson, _ := json.Marshal(deleteResp)
			log.Printf("[DEBUG] DeletePrivilege: Delete response for %s: %s", privName, string(deleteRespJson))

			log.Printf("[INFO] Privilege: Successfully deleted privilege: %s", privName)
		}
	}

	log.Printf("[INFO] Privilege: Successfully deleted all privileges for endpoint: %s", state.Endpoint.ValueString())
	return nil
}

func (r *PrivilegeResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state PrivilegeResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	err := r.DeletePrivilege(ctx, &state)
	if err != nil {
		log.Printf("[ERROR] Privilege: API Delete Failed: %v", err)
		resp.Diagnostics.AddError("API Delete Failed", err.Error())
		return
	}
}

// ValidateImportKey validates that the endpoint exists and returns its associated security system
func (r *PrivilegeResource) ValidateImportKey(ctx context.Context, endpointName string) (string, error) {
	// Create the request object for GetEndpoints
	reqParams := endpoint.GetEndpointsRequest{}
	reqParams.SetEndpointname(endpointName)

	var endpointResp *endpoint.GetEndpoints200Response
	err := r.provider.AuthenticatedAPICallWithRetry(ctx, "get_endpoints", func(token string) error {
		endpointOps := r.privilegeFactory.CreateEndpointOperations(r.client.APIBaseURL(), token)
		resp, httpResp, err := endpointOps.GetEndpoints(ctx, reqParams)
		if httpResp != nil && httpResp.StatusCode == 401 {
			return fmt.Errorf("401 unauthorized")
		}
		endpointResp = resp
		return err
	})
	if err != nil {
		log.Printf("[ERROR] Privilege: API Call failed: Failed to fetch endpoint details. Error: %v", err)
		return "", fmt.Errorf("API Call failed: Failed to fetch endpoint details. Error: %v", err)
	}

	if endpointResp == nil || len(endpointResp.Endpoints) == 0 {
		log.Printf("[ERROR] Endpoint %s not found", endpointName)
		return "", fmt.Errorf("Endpoint %s not found", endpointName)
	}

	securitySystem := endpointResp.Endpoints[0].Securitysystem
	if securitySystem == nil || *securitySystem == "" {
		log.Printf("[ERROR] Security system is empty for endpoint %s", endpointName)
		return "", fmt.Errorf("endpoint %q has no associated security system", endpointName)
	}

	return *securitySystem, nil
}

func (r *PrivilegeResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	log.Printf("Import key received: %s", req.ID)
	idParts := strings.Split(req.ID, ":")
	if len(idParts) != 2 {
		resp.Diagnostics.AddError(
			"Invalid Import ID Format",
			fmt.Sprintf("Expected import ID format: 'endpoint:entitlementtype', got: %s\n"+
				"Example: terraform import saviynt_privilege_resource.example sample_endpoint:sample_ent_type", req.ID),
		)
		return
	}

	endpointName := strings.TrimSpace(idParts[0])
	entitlementTypeName := strings.TrimSpace(idParts[1])

	log.Printf("Starting import for privilege of entitlement type: %s for endpoint %s", entitlementTypeName, endpointName)

	if endpointName == "" || entitlementTypeName == "" {
		resp.Diagnostics.AddError(
			"Invalid Import ID Components",
			"Both endpoint and entitlementtype must be non-empty\n"+
				"Example: terraform import saviynt_privilege_resource.example sample_endpoint:sample_ent_type",
		)
		return
	}

	// Validate endpoint and get security system
	securitySystem, err := r.ValidateImportKey(ctx, endpointName)
	if err != nil {
		log.Printf("[ERROR] Privilege: Import key validation failed. %v", err)
		resp.Diagnostics.AddError("Import Key validation failed: ", err.Error())
		return
	}

	// Set the attributes in the state
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("endpoint"), endpointName)...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("entitlement_type"), entitlementTypeName)...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("security_system"), securitySystem)...)
}
