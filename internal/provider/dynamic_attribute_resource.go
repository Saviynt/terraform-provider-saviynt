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

// saviynt_dynamic_attribute_resource manages ADSI connectors in the Saviynt Security Manager.
// The resource implements the full Terraform lifecycle:
//   - Create: provisions new dynamic attributes for an endpoint using the supplied configuration.
//   - Read: fetches the current dynamic attribute state from Saviynt to keep Terraform’s state in sync.
//   - Update: applies any configuration changes to an existing dynamic attribute.
//   - Import: brings existing dynamic attributes for an endpoint under Terraform management by its name.

package provider

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"terraform-provider-Saviynt/util"
	"terraform-provider-Saviynt/util/dynamicattributeutil"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"

	s "github.com/saviynt/saviynt-api-go-client"

	openapi "github.com/saviynt/saviynt-api-go-client/dynamicattributes"
	endpoint "github.com/saviynt/saviynt-api-go-client/endpoints"
)

type DynamicAttributeResourceModel struct {
	ID             types.String `tfsdk:"id"`
	Securitysystem types.String `tfsdk:"security_system"`
	Endpoint       types.String `tfsdk:"endpoint"`
	Updateuser     types.String `tfsdk:"update_user"`
	// DynamicAttributes      types.Set    `tfsdk:"dynamic_attributes"`
	DynamicAttributes      types.Map    `tfsdk:"dynamic_attributes"`
	DynamicAttributesError types.String `tfsdk:"dynamic_attribute_errors"`
	Msg                    types.String `tfsdk:"msg"`
	ErrorCode              types.String `tfsdk:"error_code"`
}

type Dynamicattribute struct {
	Attributename                                   types.String `tfsdk:"attribute_name"`
	Requesttype                                     types.String `tfsdk:"request_type"`
	Attributetype                                   types.String `tfsdk:"attribute_type"`
	Attributegroup                                  types.String `tfsdk:"attribute_group"`
	Orderindex                                      types.String `tfsdk:"order_index"`
	Attributelable                                  types.String `tfsdk:"attribute_lable"`
	Accountscolumn                                  types.String `tfsdk:"accounts_column"`
	Hideoncreate                                    types.String `tfsdk:"hide_on_create"`
	Actionstring                                    types.String `tfsdk:"action_string"`
	Editable                                        types.String `tfsdk:"editable"`
	Hideonupdate                                    types.String `tfsdk:"hide_on_update"`
	Action_to_perform_when_parent_attribute_changes types.String `tfsdk:"action_to_perform_when_parent_attribute_changes"`
	Required                                        types.String `tfsdk:"required"`
	Attributevalue                                  types.String `tfsdk:"attribute_value"`
	Showonchild                                     types.String `tfsdk:"showonchild"`
	Descriptionascsv                                types.String `tfsdk:"description_as_csv"`
	Parentattribute                                 types.String `tfsdk:"parent_attribute"`

	// Removed due to lack support in v24.4
	// Regex                                           types.String `tfsdk:"regex"`
	// Defaultvalue                                    types.String `tfsdk:"default_value"`
}

type dynamicAttributeResource struct {
	client *s.Client
	token  string
}

func NewDynamicAttributeResource() resource.Resource {
	return &dynamicAttributeResource{}
}

func (r *dynamicAttributeResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = "saviynt_dynamic_attribute_resource"
}

func (r *dynamicAttributeResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: util.DynamicAttrDescription,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The unique ID of the resource, typically managed by the API.",
			},
			"security_system": schema.StringAttribute{
				Required:    true,
				Description: "Security system associated with the dynamic attribute.",
			},
			"endpoint": schema.StringAttribute{
				Required:    true,
				Description: "Endpoint associated with the dynamic attribute.",
			},
			"update_user": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "User who last updated the dynamic attribute.",
			},
			"msg": schema.StringAttribute{
				Computed:    true,
				Description: "Response message from the API.",
			},
			"error_code": schema.StringAttribute{
				Computed:    true,
				Description: "Error code from the API response.",
			},
			"dynamic_attributes": schema.MapNestedAttribute{
				Description: "Set of dynamic attribute configuration blocks.",
				Required:    true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"attribute_name": schema.StringAttribute{
							Required:    true,
							Description: "Specify the dynamic attribute name.",
						},
						"request_type": schema.StringAttribute{
							Required:    true,
							Description: "Type of request.",
						},
						"attribute_type": schema.StringAttribute{
							Optional:    true,
							Computed:    true,
							Description: "Attribute type used for filtering and display.",
						},
						"attribute_group": schema.StringAttribute{
							Optional:    true,
							Computed:    true,
							Description: "Group or categorize the attribute in the request form.",
						},
						"order_index": schema.StringAttribute{
							Optional:    true,
							Computed:    true,
							Description: "Sequence for display of the dynamic attribute.",
						},
						"attribute_lable": schema.StringAttribute{
							Optional:    true,
							Computed:    true,
							Description: "Name to be shown in the Access Requests form.",
						},
						"accounts_column": schema.StringAttribute{
							Optional:    true,
							Computed:    true,
							Description: "Accounts column mapping.",
						},
						"hide_on_create": schema.StringAttribute{
							Optional:    true,
							Computed:    true,
							Description: "Whether to hide this attribute on create.",
						},
						"action_string": schema.StringAttribute{
							Optional:    true,
							Computed:    true,
							Description: "Action string value.",
						},
						"editable": schema.StringAttribute{
							Optional:    true,
							Computed:    true,
							Description: "Whether the attribute is editable.",
						},
						"hide_on_update": schema.StringAttribute{
							Optional:    true,
							Computed:    true,
							Description: "Whether to hide this attribute on update.",
						},
						"action_to_perform_when_parent_attribute_changes": schema.StringAttribute{
							Optional:    true,
							Computed:    true,
							Description: "Action to perform when the parent attribute changes.",
						},
						// "default_value": schema.StringAttribute{
						// 	Optional:    true,
						// 	Computed:    true,
						// 	Description: "Default value for the attribute.",
						// 	Validators: []validator.String{
						// 		dynamicattributeutil.DefaultValueDisallowedForCertainAttributeTypes(),
						// 	},
						// },
						"required": schema.StringAttribute{
							Optional:    true,
							Computed:    true,
							Description: "Whether this attribute is required.",
						},
						"attribute_value": schema.StringAttribute{
							Optional:    true,
							Computed:    true,
							Description: "Value options or query for the attribute.",
							Validators: []validator.String{
								dynamicattributeutil.AttributeValueDisallowedForCertainAttributeTypes(),
							},
						},
						"showonchild": schema.StringAttribute{
							Optional:    true,
							Computed:    true,
							Description: "Whether to show this on child requests.",
						},
						"description_as_csv": schema.StringAttribute{
							Optional:    true,
							Computed:    true,
							Description: "Description of values as CSV.",
							Validators: []validator.String{
								dynamicattributeutil.DescriptionDisallowedForCertainAttributeTypes(),
							},
						},
						// "regex": schema.StringAttribute{
						// 	Optional:    true,
						// 	Description: "Regex for validation.",
						// 	Validators: []validator.String{
						// 		dynamicattributeutil.RegexDisallowedForCertainAttributeTypes(),
						// 	},
						// },
						"parent_attribute": schema.StringAttribute{
							Optional:    true,
							Computed:    true,
							Description: "Parent attribute this one depends on.",
						},
					},
				},
			},
			"dynamic_attribute_errors": schema.StringAttribute{
				Description: "Error string or structured error details as flattened text.",
				Computed:    true,
				Optional:    true,
			},
		},
	}
}

func (r *dynamicAttributeResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

	// Set the client and token from the provider state.
	r.client = prov.client
	r.token = prov.accessToken
	log.Println("[DEBUG] DynamicAttribute: Resource configured successfully")
}

func (r *dynamicAttributeResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan DynamicAttributeResourceModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		log.Println("[ERROR] DynamicAttribute: Error getting plan data")
		return
	}

	cfg := openapi.NewConfiguration()
	apiBaseURL := strings.TrimPrefix(strings.TrimPrefix(r.client.APIBaseURL(), "https://"), "http://")
	cfg.Host = apiBaseURL
	cfg.Scheme = "https"
	cfg.AddDefaultHeader("Authorization", "Bearer "+r.token)
	cfg.HTTPClient = http.DefaultClient
	apiClient := openapi.NewAPIClient(cfg)

	// Convert dynamic attributes map to a slice of Dynamicattribute
	var dynamicAttrMap map[string]Dynamicattribute
	resp.Diagnostics.Append(plan.DynamicAttributes.ElementsAs(ctx, &dynamicAttrMap, false)...)
	if resp.Diagnostics.HasError() {
		log.Println("[ERROR] DynamicAttribute: Failed to process dynamic attributes")
		return
	}

	var dynamicAttrs []openapi.CreateDynamicAttributesInner
	for _, attr := range dynamicAttrMap {
		dynamicAttr := openapi.NewCreateDynamicAttributesInner(
			attr.Attributename.ValueString(),
			attr.Requesttype.ValueString(),
		)

		dynamicAttr.Attributetype = util.StringPointerOrEmpty(attr.Attributetype)
		dynamicAttr.Attributegroup = util.StringPointerOrEmpty(attr.Attributegroup)
		dynamicAttr.Orderindex = util.StringPointerOrEmpty(attr.Orderindex)
		dynamicAttr.Attributelable = util.StringPointerOrEmpty(attr.Attributelable)
		dynamicAttr.Accountscolumn = util.StringPointerOrEmpty(attr.Accountscolumn)
		dynamicAttr.Hideoncreate = util.StringPointerOrEmpty(attr.Hideoncreate)
		dynamicAttr.Actionstring = util.StringPointerOrEmpty(attr.Actionstring)
		dynamicAttr.Editable = util.StringPointerOrEmpty(attr.Editable)
		dynamicAttr.Hideonupdate = util.StringPointerOrEmpty(attr.Hideonupdate)
		dynamicAttr.Actiontoperformwhenparentattributechanges = util.StringPointerOrEmpty(attr.Action_to_perform_when_parent_attribute_changes)
		// dynamicAttr.Defaultvalue = util.StringPointerOrEmpty(attr.Defaultvalue)
		dynamicAttr.Required = util.StringPointerOrEmpty(attr.Required)
		// dynamicAttr.Regex = util.StringPointerOrEmpty(attr.Regex)
		dynamicAttr.Attributevalue = util.StringPointerOrEmpty(attr.Attributevalue)
		dynamicAttr.Showonchild = util.StringPointerOrEmpty(attr.Showonchild)
		dynamicAttr.Parentattribute = util.StringPointerOrEmpty(attr.Parentattribute)
		dynamicAttr.Descriptionascsv = util.StringPointerOrEmpty(attr.Descriptionascsv)
		dynamicAttrs = append(dynamicAttrs, *dynamicAttr)
	}

	createReq := openapi.NewCreateDynamicAttributeRequest(
		plan.Securitysystem.ValueString(),
		plan.Endpoint.ValueString(),
		plan.Updateuser.ValueString(),
		dynamicAttrs,
	)

	log.Println("[DEBUG] DynamicAttribute: Making API call to create dynamic attributes")
	createResp, httpResp, err := apiClient.DynamicAttributesAPI.
		CreateDynamicAttribute(ctx).
		CreateDynamicAttributeRequest(*createReq).
		Execute()

	if err != nil {
		log.Printf("[ERROR] Creating Dynamic attribute: %v, HTTP Response: %v", err, httpResp)
		resp.Diagnostics.AddError(
			"Error Creating Dynamic Attribute",
			fmt.Sprintf("API Error: %v", err),
		)
		return
	}

	if createResp.Errorcode == nil {
		log.Println("[ERROR] DynamicAttribute: Unexpected API response - Errorcode is nil")
		resp.Diagnostics.AddError(
			"Unexpected API Response",
			"Errorcode is nil in create response",
		)
		return
	}

	var errorMessages []string

	if createResp.Errorcode != nil && *createResp.Errorcode == "1" {
		plan.ErrorCode = types.StringPointerValue(createResp.Errorcode)
		plan.Msg = types.StringPointerValue(createResp.Msg)

		if createResp.Securitysystem != nil {
			errorMessages = append(errorMessages, fmt.Sprintf("securitysystem: %s", *createResp.Securitysystem))
		}

		if createResp.Endpoint != nil {
			errorMessages = append(errorMessages, fmt.Sprintf("endpoint: %s", *createResp.Endpoint))
		}

		if createResp.Updateuser != nil {
			errorMessages = append(errorMessages, fmt.Sprintf("updateuser: %s", *createResp.Updateuser))
		}

		if createResp.Dynamicattributes != nil {
			rawJSON, err := json.Marshal(createResp.Dynamicattributes)
			if err != nil {
				log.Printf("Error marshaling dynamicattributes: %v", err)
			} else {
				// Try as string
				var str string
				if err := json.Unmarshal(rawJSON, &str); err == nil {
					errorMessages = append(errorMessages, fmt.Sprintf("dynamicattributes: %s", str))
				} else {
					// Try as map
					var daMap map[string]map[string]string
					if err := json.Unmarshal(rawJSON, &daMap); err == nil {
						for attrID, fieldErrors := range daMap {
							for field, msg := range fieldErrors {
								errorMessages = append(errorMessages, fmt.Sprintf("dynamicattributes.%s.%s: %s", attrID, field, msg))
							}
						}
					} else {
						errorMessages = append(errorMessages, "dynamicattributes: unknown format")
					}
				}
			}
		}

		fullError := strings.Join(errorMessages, "\n")
		if strings.Contains(fullError, "attributename already exists") {
			fullError += "\n\nTry importing the resource or use a different name."
		}

		plan.DynamicAttributesError = types.StringValue(fullError)
		resp.Diagnostics.AddError("Dynamic Attribute Create Operation Failed", fullError)
		return
	} else {
		plan.DynamicAttributesError = types.StringNull()
	}

	// Update the dynamic attributes in state with SafeString handling
	updatedAttrs := make(map[string]Dynamicattribute)
	for attrName, attr := range dynamicAttrMap {
		// attrType := strings.ToLower(attr.Attributetype.ValueString())
		updatedAttr := Dynamicattribute{
			Attributename:  util.SafeString(attr.Attributename.ValueStringPointer()),
			Requesttype:    util.SafeString(attr.Requesttype.ValueStringPointer()),
			Attributetype:  util.SafeString(attr.Attributetype.ValueStringPointer()),
			Attributegroup: util.SafeString(attr.Attributegroup.ValueStringPointer()),
			Orderindex:     util.SafeStringAlt(attr.Orderindex.ValueStringPointer(), "0"),
			Attributelable: util.SafeString(attr.Attributelable.ValueStringPointer()),
			Accountscolumn: util.SafeString(attr.Accountscolumn.ValueStringPointer()),
			Hideoncreate:   util.SafeStringAlt(attr.Hideoncreate.ValueStringPointer(), "false"),
			Actionstring:   util.SafeString(attr.Actionstring.ValueStringPointer()),
			Editable:       util.SafeStringAlt(attr.Editable.ValueStringPointer(), "false"),
			Hideonupdate:   util.SafeStringAlt(attr.Hideoncreate.ValueStringPointer(), "false"),
			Action_to_perform_when_parent_attribute_changes: util.SafeString(attr.Action_to_perform_when_parent_attribute_changes.ValueStringPointer()),
			// Defaultvalue: util.SafeString(attr.Defaultvalue.ValueStringPointer()),
			Required:     util.SafeStringAlt(attr.Required.ValueStringPointer(), "false"),
			// Regex:            util.SafeStringDatasource(attr.Regex.ValueStringPointer()),
			Showonchild: util.SafeStringAlt(attr.Showonchild.ValueStringPointer(), "false"),
			Parentattribute:  util.SafeString(attr.Parentattribute.ValueStringPointer()),
			Descriptionascsv: util.SafeString(attr.Descriptionascsv.ValueStringPointer()),
		}
		if attr.Attributevalue.IsNull() || attr.Attributevalue.IsUnknown() {
			updatedAttr.Attributevalue = types.StringNull()
		} else {
			updatedAttr.Attributevalue = types.StringValue(attr.Attributevalue.ValueString())
		}
		updatedAttrs[attrName] = updatedAttr
		log.Printf("%s: %+v", attrName, updatedAttr)
	}

	// Convert the updated map back to types.Map
	updatedMap, diags := types.MapValueFrom(ctx, types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"attribute_name":  types.StringType,
			"request_type":    types.StringType,
			"attribute_type":  types.StringType,
			"attribute_group": types.StringType,
			"order_index":     types.StringType,
			"attribute_lable": types.StringType,
			"accounts_column": types.StringType,
			"hide_on_create":  types.StringType,
			"action_string":   types.StringType,
			"editable":        types.StringType,
			"hide_on_update":  types.StringType,
			"action_to_perform_when_parent_attribute_changes": types.StringType,
			// "default_value": types.StringType,
			"required":      types.StringType,
			// "regex":              types.StringType,
			"attribute_value":    types.StringType,
			"showonchild":        types.StringType,
			"parent_attribute":   types.StringType,
			"description_as_csv": types.StringType,
		},
	}, updatedAttrs)

	if diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}

	plan.DynamicAttributes = updatedMap
	plan.ID = types.StringValue("dynamic-attr-" + plan.Endpoint.ValueString())
	plan.Msg = types.StringValue(util.SafeDeref(createResp.Msg))
	plan.ErrorCode = types.StringValue(util.SafeDeref(createResp.Errorcode))
	plan.DynamicAttributesError = types.StringNull()
	plan.Updateuser = types.StringValue(plan.Updateuser.ValueString())

	// Set the final state
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
	log.Printf("[DEBUG] DynamicAttribute: Successfully created dynamic attributes for endpoint: %s", plan.Endpoint.ValueString())
}

func (r *dynamicAttributeResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state DynamicAttributeResourceModel

	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
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

	fetchReq := apiClient.DynamicAttributesAPI.
		FetchDynamicAttribute(ctx).
		Endpoint([]string{state.Endpoint.ValueString()})
	log.Printf("[DEBUG] Fetch Request read: %+v", state.Endpoint.ValueString())

	log.Printf("[DEBUG] DynamicAttribute: Making API call to fetch attributes for endpoint: %s", state.Endpoint.ValueString())
	apiResp, httpResp, err := fetchReq.Execute()

	if httpResp != nil && httpResp.StatusCode == 412 {
		var fetchResp map[string]interface{}
		if err := json.NewDecoder(httpResp.Body).Decode(&fetchResp); err != nil {
			resp.Diagnostics.AddError("Failed to decode error response", err.Error())
			return
		}

		var errorMessages []string

		// Process securitysystems
		if secSystems, ok := fetchResp["securitysystems"].([]interface{}); ok {
			for _, item := range secSystems {
				if secMap, ok := item.(map[string]interface{}); ok {
					for key, val := range secMap {
						errorMessages = append(errorMessages, fmt.Sprintf("securitysystem: %s - %v", key, val))
					}
				}
			}
		}

		// Process endpoints
		if endpoints, ok := fetchResp["endpoints"].([]interface{}); ok {
			for _, item := range endpoints {
				if epMap, ok := item.(map[string]interface{}); ok {
					for key, val := range epMap {
						errorMessages = append(errorMessages, fmt.Sprintf("endpoint: %s - %v", key, val))
					}
				}
			}
		}

		if requestTypes, ok := fetchResp["requesttype"].([]interface{}); ok {
			for _, item := range requestTypes {
				if epMap, ok := item.(map[string]interface{}); ok {
					for key, val := range epMap {
						errorMessages = append(errorMessages, fmt.Sprintf("request type: %s - %v", key, val))
					}
				}
			}
		}

		// Process general msg and errorcode
		if msg, ok := fetchResp["msg"].(string); ok {
			errorMessages = append(errorMessages, fmt.Sprintf("message: %s", msg))
		}
		if ec, ok := fetchResp["errorcode"].(string); ok {
			errorMessages = append(errorMessages, fmt.Sprintf("errorcode: %s", ec))
		}

		// Combine and populate the Terraform state
		fullError := strings.Join(errorMessages, "\n")
		state.Securitysystem = types.StringValue("")
		state.Endpoint = types.StringValue("")
		state.DynamicAttributesError = types.StringValue(fullError)

		// Persist in state
		diags := resp.State.Set(ctx, &state)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

		// Display error to user
		resp.Diagnostics.AddError("Dynamic Attributes Fetch Failed", fullError)
		return
	}

	if err != nil {
		log.Printf("[ERROR] Fetch API Call Failed: %v", err)
		resp.Diagnostics.AddError("Fetch API Call Failed", fmt.Sprintf("Error: %v", err))
		return
	}

	log.Printf("[DEBUG] Read HTTP Status Code: %d", httpResp.StatusCode)
	log.Printf("[DEBUG] Read API Error code: %+v", util.SafeDeref(apiResp.Errorcode))
	log.Printf("[DEBUG] Read API Message: %+v", util.SafeDeref(apiResp.Msg))

	// Get current state attributes as a map
	currentAttrs := make(map[string]Dynamicattribute)
	if !state.DynamicAttributes.IsNull() {
		diags := state.DynamicAttributes.ElementsAs(ctx, &currentAttrs, false)
		if diags.HasError() {
			resp.Diagnostics.Append(diags...)
			return
		}
	}
	log.Printf("[DEBUG] Existing attributes (%d):", len(currentAttrs))
	for attrName, attr := range currentAttrs {
		log.Printf("  - %s: %+v", attrName, attr)
	}

	// Create map to store API attributes by name
	apiAttrs := make(map[string]openapi.FetchDynamicAttributeResponseInner)
	if apiResp != nil && apiResp.Dynamicattributes != nil && apiResp.Dynamicattributes.ArrayOfFetchDynamicAttributeResponseInner != nil {
		for _, item := range *apiResp.Dynamicattributes.ArrayOfFetchDynamicAttributeResponseInner {
			attrName := util.SafeDeref(item.Attributename)
			if attrName != "" {
				apiAttrs[attrName] = item
			}
		}
	}
	log.Printf("Api attributes: %v", apiAttrs)

	// Build updated attributes map - only include attributes that exist in API response
	updatedAttrs := make(map[string]Dynamicattribute)
	if len(currentAttrs) == 0 {
		// IMPORT CASE: State is empty, bring in all attributes from API
		for attrName, apiAttr := range apiAttrs {
			updatedAttr := Dynamicattribute{
				Attributename:  types.StringValue(attrName),
				Requesttype:    util.SafeStringDatasource(apiAttr.Requesttype),
				Attributetype:  util.SafeStringDatasource(apiAttr.Attributetype),
				Orderindex:     util.SafeStringDatasource(apiAttr.Orderindex),
				Required:       util.SafeStringDatasource(apiAttr.Required),
				Editable:       util.SafeStringDatasource(apiAttr.Editable),
				Hideoncreate:   util.SafeStringDatasource(apiAttr.Hideoncreate),
				Hideonupdate:   util.SafeStringDatasource(apiAttr.Hideonupdate),
				Attributegroup: util.SafeStringDatasource(apiAttr.Attributegroup),
				Attributelable: util.SafeStringDatasource(apiAttr.Attributelable),
				Accountscolumn: util.SafeStringDatasource(apiAttr.Accountscolumn),
				Actionstring:   util.SafeStringDatasource(apiAttr.Actionstring),
				Action_to_perform_when_parent_attribute_changes: util.SafeStringDatasource(apiAttr.Actiontoperformwhenparentattributechanges),
				// Defaultvalue: util.SafeStringDatasource(apiAttr.Defaultvalue),
				// Regex:            util.PreserveString(apiAttr.Regex, currentAttrs[attrName].Regex),
				// Attributevalue:   util.SafeStringDatasource(apiAttr.Attributevalue),
				Attributevalue: util.SafeStringPreserveNull(apiAttr.Attributevalue),
				Showonchild:    util.SafeStringDatasource(apiAttr.Showonchild),
				Parentattribute:  util.SafeStringDatasource(apiAttr.Parentattribute),
				Descriptionascsv: util.SafeStringDatasource(apiAttr.Descriptionascsv),
			}
			updatedAttrs[attrName] = updatedAttr
		}
		log.Printf("[IMPORT] Importing all %d dynamic attributes from API", len(apiAttrs))
	} else {
		removedAttributes := make([]string, 0)

		for attrName := range currentAttrs {
			if apiAttr, exists := apiAttrs[attrName]; exists {
				// Attribute exists in API response - update with API values
				updatedAttr := Dynamicattribute{
					Attributename: types.StringValue(attrName),
					Requesttype:   util.SafeStringDatasource(apiAttr.Requesttype),
					// Attributetype:  util.SafeStringDatasource(apiAttr.Attributetype),
					Attributetype:  types.StringValue(dynamicattributeutil.TranslateValue(*apiAttr.Attributetype, dynamicattributeutil.AttributeTypeMap)),
					Orderindex:     util.SafeStringDatasource(apiAttr.Orderindex),
					Required:       util.SafeStringDatasource(apiAttr.Required),
					Editable:       util.SafeStringDatasource(apiAttr.Editable),
					Hideoncreate:   util.SafeStringDatasource(apiAttr.Hideoncreate),
					Hideonupdate:   util.SafeStringDatasource(apiAttr.Hideonupdate),
					Attributegroup: util.SafeStringDatasource(apiAttr.Attributegroup),
					Attributelable: util.SafeStringDatasource(apiAttr.Attributelable),
					Accountscolumn: util.SafeStringDatasource(apiAttr.Accountscolumn),
					Actionstring:   util.SafeStringDatasource(apiAttr.Actionstring),
					Action_to_perform_when_parent_attribute_changes: util.SafeStringDatasource(apiAttr.Actiontoperformwhenparentattributechanges),
					// Defaultvalue: util.SafeStringDatasource(apiAttr.Defaultvalue),
					// Regex:            util.PreserveString(apiAttr.Regex, currentAttrs[attrName].Regex),
					// Attributevalue:   util.SafeStringDatasource(apiAttr.Attributevalue),
					Attributevalue:   util.SafeStringPreserveNull(apiAttr.Attributevalue),
					Showonchild:      util.SafeStringDatasource(apiAttr.Showonchild),
					Parentattribute:  util.SafeStringDatasource(apiAttr.Parentattribute),
					Descriptionascsv: util.SafeStringDatasource(apiAttr.Descriptionascsv),
				}
				updatedAttrs[attrName] = updatedAttr
			} else {
				// Attribute doesn't exist in API response - mark for removal
				removedAttributes = append(removedAttributes, attrName)
			}
		}
		log.Printf("Updated attributes: %q", updatedAttrs)
		// Log removed attributes
		if len(removedAttributes) > 0 {
			log.Printf("[INFO] Removing attributes not present in API response: %v", removedAttributes)
		}
	}
	// Convert to types.Map
	dynamicAttributesMap, diags := types.MapValueFrom(ctx, types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"attribute_name":  types.StringType,
			"request_type":    types.StringType,
			"attribute_type":  types.StringType,
			"attribute_group": types.StringType,
			"order_index":     types.StringType,
			"attribute_lable": types.StringType,
			"accounts_column": types.StringType,
			"hide_on_create":  types.StringType,
			"action_string":   types.StringType,
			"editable":        types.StringType,
			"hide_on_update":  types.StringType,
			"action_to_perform_when_parent_attribute_changes": types.StringType,
			// "default_value": types.StringType,
			"required":      types.StringType,
			// "regex":              types.StringType,
			"attribute_value":    types.StringType,
			"showonchild":        types.StringType,
			"parent_attribute":   types.StringType,
			"description_as_csv": types.StringType,
		},
	}, updatedAttrs)

	if diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}

	state.DynamicAttributes = dynamicAttributesMap
	state.ID = types.StringValue("dynamic-attr-" + state.Endpoint.ValueString())
	state.Endpoint = util.SafeString(state.Endpoint.ValueStringPointer())
	state.Updateuser = util.SafeString(state.Updateuser.ValueStringPointer())
	state.Msg = types.StringValue(util.SafeDeref(apiResp.Msg))
	state.ErrorCode = types.StringValue(util.SafeDeref(apiResp.Errorcode))

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	log.Printf("[INFO] DynamicAttribute: Successfully read %d attributes for endpoint: %s", len(updatedAttrs), state.Endpoint.ValueString())
}

func (r *dynamicAttributeResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan DynamicAttributeResourceModel

	// Get the plan
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Initialize API client
	cfg := openapi.NewConfiguration()
	apiBaseURL := strings.TrimPrefix(strings.TrimPrefix(r.client.APIBaseURL(), "https://"), "http://")
	cfg.Host = apiBaseURL
	cfg.Scheme = "https"
	cfg.AddDefaultHeader("Authorization", "Bearer "+r.token)
	cfg.HTTPClient = http.DefaultClient
	apiClient := openapi.NewAPIClient(cfg)

	// Get planned attributes
	var planAttrs map[string]Dynamicattribute
	resp.Diagnostics.Append(plan.DynamicAttributes.ElementsAs(ctx, &planAttrs, false)...)
	if resp.Diagnostics.HasError() {
		return
	}
	log.Printf("Planned attributes before fetch: %v", len(planAttrs))

	// Get current state from server
	fetchReq := apiClient.DynamicAttributesAPI.
		FetchDynamicAttribute(ctx).
		Securitysystem([]string{plan.Securitysystem.ValueString()}).
		Endpoint([]string{plan.Endpoint.ValueString()})

	fetchResp, _, err := fetchReq.Execute()
	if err != nil {
		resp.Diagnostics.AddError("Failed to fetch current dynamic attributes", err.Error())
		return
	}
	log.Printf("Fetching existing attributes from api complete")

	// Build map of existing attributes
	existingAttrs := make(map[string]bool)
	if fetchResp.Dynamicattributes.ArrayOfFetchDynamicAttributeResponseInner != nil {
		for _, item := range *fetchResp.Dynamicattributes.ArrayOfFetchDynamicAttributeResponseInner {
			if item.Attributename != nil {
				existingAttrs[*item.Attributename] = true
			}
		}
	}
	log.Printf("Existing attributes from api: %v", len(existingAttrs))

	// Separate new attributes (need to be created) from existing ones (need to be updated)
	var newAttrs []openapi.CreateDynamicAttributesInner
	var updateAttrs []openapi.UpdateDynamicAttributesInner

	for _, attr := range planAttrs {
		attrName := attr.Attributename.ValueString()

		if !existingAttrs[attrName] {
			// This is a new attribute - prepare for creation
			newAttr := openapi.NewCreateDynamicAttributesInner(
				attrName,
				attr.Requesttype.ValueString(),
			)
			// Set all fields...
			newAttr.Attributetype = util.StringPointerOrEmpty(attr.Attributetype)
			newAttr.Attributegroup = util.StringPointerOrEmpty(attr.Attributegroup)
			newAttr.Orderindex = util.StringPointerOrEmpty(attr.Orderindex)
			newAttr.Accountscolumn = util.StringPointerOrEmpty(attr.Accountscolumn)
			newAttr.Attributelable = util.StringPointerOrEmpty(attr.Attributelable)
			newAttr.Hideoncreate = util.StringPointerOrEmpty(attr.Hideoncreate)
			newAttr.Actionstring = util.StringPointerOrEmpty(attr.Actionstring)
			newAttr.Editable = util.StringPointerOrEmpty(attr.Editable)
			newAttr.Hideonupdate = util.StringPointerOrEmpty(attr.Hideonupdate)
			newAttr.Actiontoperformwhenparentattributechanges = util.StringPointerOrEmpty(attr.Action_to_perform_when_parent_attribute_changes)
			// newAttr.Defaultvalue = util.StringPointerOrEmpty(attr.Defaultvalue)
			newAttr.Required = util.StringPointerOrEmpty(attr.Required)
			// newAttr.Regex = util.StringPointerOrEmpty(attr.Regex)
			newAttr.Attributevalue = util.StringPointerOrEmpty(attr.Attributevalue)
			newAttr.Showonchild = util.StringPointerOrEmpty(attr.Showonchild)
			newAttr.Parentattribute = util.StringPointerOrEmpty(attr.Parentattribute)
			newAttr.Descriptionascsv = util.StringPointerOrEmpty(attr.Descriptionascsv)
			newAttrs = append(newAttrs, *newAttr)
		} else {
			// Existing attribute - prepare for update
			updateAttr := openapi.NewUpdateDynamicAttributesInner(attrName)
			// Set all fields...
			// updateAttr := openapi.NewUpdateDynamicAttributesInner(attrName)
			updateAttr.Requesttype = util.StringPointerOrEmpty(attr.Requesttype)
			updateAttr.Attributetype = util.StringPointerOrEmpty(attr.Attributetype)
			updateAttr.Attributegroup = util.StringPointerOrEmpty(attr.Attributegroup)
			updateAttr.Orderindex = util.StringPointerOrEmpty(attr.Orderindex)
			updateAttr.Attributelable = util.StringPointerOrEmpty(attr.Attributelable)
			updateAttr.Accountscolumn = util.StringPointerOrEmpty(attr.Accountscolumn)
			updateAttr.Hideoncreate = util.StringPointerOrEmpty(attr.Hideoncreate)
			updateAttr.Actionstring = util.StringPointerOrEmpty(attr.Actionstring)
			updateAttr.Editable = util.StringPointerOrEmpty(attr.Editable)
			updateAttr.Hideonupdate = util.StringPointerOrEmpty(attr.Hideonupdate)
			updateAttr.Actiontoperformwhenparentattributechanges = util.StringPointerOrEmpty(attr.Action_to_perform_when_parent_attribute_changes)
			// updateAttr.Defaultvalue = util.StringPointerOrEmpty(attr.Defaultvalue)
			updateAttr.Required = util.StringPointerOrEmpty(attr.Required)
			// updateAttr.Regex = util.StringPointerOrEmpty(attr.Regex)
			updateAttr.Attributevalue = util.StringPointerOrEmpty(attr.Attributevalue)
			updateAttr.Showonchild = util.StringPointerOrEmpty(attr.Showonchild)
			updateAttr.Parentattribute = util.StringPointerOrEmpty(attr.Parentattribute)
			updateAttr.Descriptionascsv = util.StringPointerOrEmpty(attr.Descriptionascsv)
			updateAttrs = append(updateAttrs, *updateAttr)
		}
	}
	log.Printf("Attributes to create &+%v", len(newAttrs))
	log.Printf("Attributes to update: %v", len(updateAttrs))

	// First, create new attributes if any
	if len(newAttrs) > 0 {
		createReq := openapi.NewCreateDynamicAttributeRequest(
			plan.Securitysystem.ValueString(),
			plan.Endpoint.ValueString(),
			plan.Updateuser.ValueString(),
			newAttrs,
		)

		_, _, err := apiClient.DynamicAttributesAPI.CreateDynamicAttribute(ctx).CreateDynamicAttributeRequest(*createReq).Execute()
		if err != nil {
			resp.Diagnostics.AddError("Failed to create new dynamic attributes", err.Error())
			return
		}
	}

	// Then, update existing attributes if any
	if len(updateAttrs) > 0 {
		updateReq := openapi.NewUpdateDynamicAttributeRequest(
			plan.Securitysystem.ValueString(),
			plan.Endpoint.ValueString(),
			plan.Updateuser.ValueString(),
			updateAttrs,
		)

		updateResp, _, err := apiClient.DynamicAttributesAPI.UpdateDynamicAttribute(ctx).UpdateDynamicAttributeRequest(*updateReq).Execute()
		if err != nil {
			resp.Diagnostics.AddError("Failed to update dynamic attributes", err.Error())
			return
		}

		if updateResp.Errorcode != nil && *updateResp.Errorcode != "0" {
			resp.Diagnostics.AddError("Dynamic Attribute Update Failed",
				fmt.Sprintf("Error: %s, Message: %s", *updateResp.Errorcode, *updateResp.Msg))
			return
		}

		if updateResp.Errorcode != nil && *updateResp.Errorcode == "1" {
			var errorMessages []string
			plan.ErrorCode = types.StringPointerValue(updateResp.Errorcode)
			plan.Msg = types.StringPointerValue(updateResp.Msg)

			if updateResp.Securitysystem != nil {
				errorMessages = append(errorMessages, fmt.Sprintf("securitysystem: %s", *updateResp.Securitysystem))
			}
			if updateResp.Endpoint != nil {
				errorMessages = append(errorMessages, fmt.Sprintf("endpoint: %s", *updateResp.Endpoint))
			}
			if updateResp.Updateuser != nil {
				errorMessages = append(errorMessages, fmt.Sprintf("updateuser: %s", *updateResp.Updateuser))
			}

			if updateResp.Dynamicattributes != nil {
				rawJSON, err := json.Marshal(updateResp.Dynamicattributes)
				if err != nil {
					log.Printf("Error marshaling dynamicattributes: %v", err)
				} else {
					// Try as string
					var str string
					if err := json.Unmarshal(rawJSON, &str); err == nil {
						errorMessages = append(errorMessages, fmt.Sprintf("dynamicattributes: %s", str))
					} else {
						// Try as map
						var daMap map[string]map[string]string
						if err := json.Unmarshal(rawJSON, &daMap); err == nil {
							for attrID, fieldErrors := range daMap {
								for field, msg := range fieldErrors {
									errorMessages = append(errorMessages, fmt.Sprintf("dynamicattributes.%s.%s: %s", attrID, field, msg))
								}
							}
						} else {
							errorMessages = append(errorMessages, "dynamicattributes: unknown format")
						}
					}
				}
			}

			fullError := strings.Join(errorMessages, "\n")
			plan.DynamicAttributesError = types.StringValue(fullError)
			resp.Diagnostics.AddError("Dynamic Attribute Update Operation Failed", fullError)
			return
		}
	}

	postUpdateFetchReq := apiClient.DynamicAttributesAPI.
		FetchDynamicAttribute(ctx).
		Securitysystem([]string{plan.Securitysystem.ValueString()}).
		Endpoint([]string{plan.Endpoint.ValueString()})

	postUpdateFetchResp, httpResp, err := postUpdateFetchReq.Execute()
	if err != nil {
		resp.Diagnostics.AddError("Failed to fetch updated dynamic attributes", err.Error())
		return
	}

	if httpResp != nil && httpResp.StatusCode == 412 {
		var fetchResp map[string]interface{}
		if err := json.NewDecoder(httpResp.Body).Decode(&fetchResp); err != nil {
			resp.Diagnostics.AddError("Failed to decode error response", err.Error())
			return
		}

		var errorMessages []string
		if secSystems, ok := fetchResp["securitysystems"].([]interface{}); ok {
			for _, item := range secSystems {
				if secMap, ok := item.(map[string]interface{}); ok {
					for key, val := range secMap {
						errorMessages = append(errorMessages, fmt.Sprintf("securitysystem: %s - %v", key, val))
					}
				}
			}
		}
		if endpoints, ok := fetchResp["endpoints"].([]interface{}); ok {
			for _, item := range endpoints {
				if epMap, ok := item.(map[string]interface{}); ok {
					for key, val := range epMap {
						errorMessages = append(errorMessages, fmt.Sprintf("endpoint: %s - %v", key, val))
					}
				}
			}
		}
		if requestTypes, ok := fetchResp["requesttype"].([]interface{}); ok {
			for _, item := range requestTypes {
				if epMap, ok := item.(map[string]interface{}); ok {
					for key, val := range epMap {
						errorMessages = append(errorMessages, fmt.Sprintf("request type: %s - %v", key, val))
					}
				}
			}
		}
		if msg, ok := fetchResp["msg"].(string); ok {
			errorMessages = append(errorMessages, fmt.Sprintf("message: %s", msg))
		}
		if ec, ok := fetchResp["errorcode"].(string); ok {
			errorMessages = append(errorMessages, fmt.Sprintf("errorcode: %s", ec))
		}

		fullError := strings.Join(errorMessages, "\n")
		plan.DynamicAttributesError = types.StringValue(fullError)
		resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
		resp.Diagnostics.AddError("Dynamic Attributes Fetch after update Failed", fullError)
		return
	}
	log.Printf("Post updation read successful")

	postUpdateApiAttrs := make(map[string]openapi.FetchDynamicAttributeResponseInner)
	if postUpdateFetchResp != nil && postUpdateFetchResp.Dynamicattributes != nil && postUpdateFetchResp.Dynamicattributes.ArrayOfFetchDynamicAttributeResponseInner != nil {
		for _, item := range *postUpdateFetchResp.Dynamicattributes.ArrayOfFetchDynamicAttributeResponseInner {
			attrName := util.SafeDeref(item.Attributename)
			if attrName != "" {
				postUpdateApiAttrs[attrName] = item
			}
		}
	}
	log.Printf("Api attributes postupdate: %v", postUpdateApiAttrs)
	updatedAttrs := make(map[string]Dynamicattribute)
	log.Printf("Planned attributes: %v", len(planAttrs))
	// removedAttributes := make([]string, 0)

	for attrName := range planAttrs {
		if apiAttr, exists := postUpdateApiAttrs[attrName]; exists {
			// Attribute exists in API response - update with API values
			updatedAttr := Dynamicattribute{
				Attributename: types.StringValue(attrName),
				Requesttype:   util.SafeStringDatasource(apiAttr.Requesttype),
				// Attributetype:  util.SafeStringDatasource(apiAttr.Attributetype),
				Attributetype:  types.StringValue(dynamicattributeutil.TranslateValue(*apiAttr.Attributetype, dynamicattributeutil.AttributeTypeMap)),
				Orderindex:     util.SafeStringDatasource(apiAttr.Orderindex),
				Required:       util.SafeStringDatasource(apiAttr.Required),
				Editable:       util.SafeStringDatasource(apiAttr.Editable),
				Hideoncreate:   util.SafeStringDatasource(apiAttr.Hideoncreate),
				Hideonupdate:   util.SafeStringDatasource(apiAttr.Hideonupdate),
				Attributegroup: util.SafeStringDatasource(apiAttr.Attributegroup),
				Attributelable: util.SafeStringDatasource(apiAttr.Attributelable),
				Accountscolumn: util.SafeStringDatasource(apiAttr.Accountscolumn),
				Actionstring:   util.SafeStringDatasource(apiAttr.Actionstring),
				Action_to_perform_when_parent_attribute_changes: util.SafeStringDatasource(apiAttr.Actiontoperformwhenparentattributechanges),
				// Defaultvalue: util.SafeStringDatasource(apiAttr.Defaultvalue),
				// Regex:            util.PreserveString(apiAttr.Regex, currentAttrs[attrName].Regex),
				Attributevalue:   util.SafeStringPreserveNull(apiAttr.Attributevalue),
				Showonchild:      util.SafeStringDatasource(apiAttr.Showonchild),
				Parentattribute:  util.SafeStringDatasource(apiAttr.Parentattribute),
				Descriptionascsv: util.SafeStringDatasource(apiAttr.Descriptionascsv),
			}
			updatedAttrs[attrName] = updatedAttr
		}
	}

	log.Printf("[INFO] Processed %d dynamic attributes in state refresh", len(updatedAttrs))

	// Convert to types.Map
	// Convert the updated attributes map to a Terraform Map type
	dynamicAttributesMap, diags := types.MapValueFrom(ctx, types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"attribute_name":  types.StringType,
			"request_type":    types.StringType,
			"attribute_type":  types.StringType,
			"attribute_group": types.StringType,
			"order_index":     types.StringType,
			"attribute_lable": types.StringType,
			"accounts_column": types.StringType,
			"hide_on_create":  types.StringType,
			"action_string":   types.StringType,
			"editable":        types.StringType,
			"hide_on_update":  types.StringType,
			"action_to_perform_when_parent_attribute_changes": types.StringType,
			// "default_value": types.StringType,
			"required":      types.StringType,
			// "regex":           types.StringType,
			"attribute_value":    types.StringType,
			"showonchild":        types.StringType,
			"parent_attribute":   types.StringType,
			"description_as_csv": types.StringType,
		},
	}, updatedAttrs)

	if diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}

	plan.DynamicAttributes = dynamicAttributesMap
	plan.ID = types.StringValue("dynamic-attr-" + plan.Endpoint.ValueString())
	plan.DynamicAttributesError = types.StringNull()
	plan.Endpoint = util.SafeString(plan.Endpoint.ValueStringPointer())
	plan.Securitysystem = util.SafeString(plan.Securitysystem.ValueStringPointer())
	plan.Updateuser = util.SafeString(plan.Updateuser.ValueStringPointer())
	plan.Msg = types.StringValue(util.SafeDeref(fetchResp.Msg))
	plan.ErrorCode = types.StringValue(util.SafeDeref(fetchResp.Errorcode))

	// Set final state
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)

}

func (r *dynamicAttributeResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state DynamicAttributeResourceModel

	stateRetrievalDiagnostics := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(stateRetrievalDiagnostics...)
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

	// Get attribute names from the map
	var attributeNames []string
	var stateAttrs map[string]Dynamicattribute

	// Properly unmarshal the map
	diags := state.DynamicAttributes.ElementsAs(ctx, &stateAttrs, false)
	if diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}

	for _, attr := range stateAttrs {
		if !attr.Attributename.IsNull() && !attr.Attributename.IsUnknown() {
			attributeNames = append(attributeNames, attr.Attributename.ValueString())
		}
	}

	deleteReq := openapi.DeleteDynamicAttributeRequest{
		Securitysystem:    state.Securitysystem.ValueString(),
		Endpoint:          state.Endpoint.ValueString(),
		Updateuser:        state.Updateuser.ValueString(),
		Dynamicattributes: attributeNames,
	}

	deleteResp, httpResp, err := apiClient.DynamicAttributesAPI.
		DeleteDynamicAttribute(ctx).
		DeleteDynamicAttributeRequest(deleteReq).
		Execute()

	if err != nil {
		log.Printf("[ERROR] Delete API Call Failed: %v", err)
		resp.Diagnostics.AddError("Delete API Call Failed", fmt.Sprintf("Error: %v", err))
		return
	}

	var errorMessages []string

	if deleteResp.Errorcode != nil && *deleteResp.Errorcode == "1" {
		if deleteResp.Securitysystem != nil {
			errorMessages = append(errorMessages, fmt.Sprintf("securitysystem: %s", *deleteResp.Securitysystem))
		}

		if deleteResp.Endpoint != nil {
			errorMessages = append(errorMessages, fmt.Sprintf("endpoint: %s", *deleteResp.Endpoint))
		}

		if deleteResp.Updateuser != nil {
			errorMessages = append(errorMessages, fmt.Sprintf("updateuser: %s", *deleteResp.Updateuser))
		}

		if deleteResp.Dynamicattributes != nil {
			rawJSON, err := json.Marshal(deleteResp.Dynamicattributes)
			if err != nil {
				log.Printf("Error marshaling dynamicattributes: %v", err)
			} else {
				var str string
				if err := json.Unmarshal(rawJSON, &str); err == nil {
					errorMessages = append(errorMessages, fmt.Sprintf("dynamicattributes: %s", str))
				} else {
					var daMap map[string]map[string]string
					if err := json.Unmarshal(rawJSON, &daMap); err == nil {
						for attrID, fieldErrors := range daMap {
							for field, msg := range fieldErrors {
								errorMessages = append(errorMessages, fmt.Sprintf("dynamicattributes.%s.%s: %s", attrID, field, msg))
							}
						}
					} else {
						errorMessages = append(errorMessages, "dynamicattributes: unknown format")
					}
				}
			}
		}

		fullError := strings.Join(errorMessages, "\n")
		resp.Diagnostics.AddError("Dynamic Attribute Operation Failed", fullError)
		return
	}

	log.Printf("[DEBUG] Delete HTTP Status Code: %d", httpResp.StatusCode)
	log.Printf("[DEBUG] Delete API Response: %+v", deleteResp)
}

func (r *dynamicAttributeResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	endpointName := req.ID

	// Set the endpoint from import ID
	resp.State.SetAttribute(ctx, path.Root("endpoint"), endpointName)

	reqParams := endpoint.GetEndpointsRequest{}
	reqParams.SetEndpointname(endpointName)
	endpointResp, _, err := r.client.Endpoints.GetEndpoints(ctx).GetEndpointsRequest(reqParams).Execute()

	if err != nil || endpointResp == nil || len(endpointResp.Endpoints) == 0 {
		resp.Diagnostics.AddError(
			"Failed to fetch endpoint details for security system name while importing",
			fmt.Sprintf("Error retrieving endpoint %q for security system name while importing: %v", endpointName, err),
		)
		return
	}

	securitySystem := endpointResp.Endpoints[0].Securitysystem
	if *securitySystem == "" {
		resp.Diagnostics.AddWarning(
			"Security system is empty in import",
			fmt.Sprintf("Endpoint %q has no associated security system in import", endpointName),
		)
	}
	resp.State.SetAttribute(ctx, path.Root("security_system"), securitySystem)
	resp.State.SetAttribute(ctx, path.Root("update_user"), r.client.Username)
}
