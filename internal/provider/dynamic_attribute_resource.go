// Copyright (c) 2025 Saviynt Inc.
// SPDX-License-Identifier: MPL-2.0

// saviynt_dynamic_attribute_resource manages dynamic attributes in the Saviynt Security Manager.
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
	"terraform-provider-Saviynt/internal/client"
	"terraform-provider-Saviynt/util"
	"terraform-provider-Saviynt/util/dynamicattributeutil"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"

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
	Defaultvalue                                    types.String `tfsdk:"default_value"`

	// Removed due to lack support Read API
	// Regex                                           types.String `tfsdk:"regex"`
}

type DynamicAttributeResource struct {
	client                  client.SaviyntClientInterface
	token                   string
	username                string
	dynamicAttributeFactory client.DynamicAttributeFactoryInterface
}

// Ensure provider defined types fully satisfy framework interfaces
var _ resource.Resource = &DynamicAttributeResource{}
var _ resource.ResourceWithImportState = &DynamicAttributeResource{}

func NewDynamicAttributeResource() resource.Resource {
	return &DynamicAttributeResource{
		dynamicAttributeFactory: &client.DefaultDynamicAttributeFactory{},
	}
}

func NewDynamicAttributeResourceWithFactory(factory client.DynamicAttributeFactoryInterface) resource.Resource {
	return &DynamicAttributeResource{
		dynamicAttributeFactory: factory,
	}
}

func (r *DynamicAttributeResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = "saviynt_dynamic_attribute_resource"
}

func (r *DynamicAttributeResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
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
						"default_value": schema.StringAttribute{
							Optional:    true,
							Computed:    true,
							Description: "Default value for the attribute(Currently not configurable for BOOLEAN attribute type from Terraform).",
							Validators: []validator.String{
								dynamicattributeutil.DefaultValueDisallowedForCertainAttributeTypes(),
							},
						},
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

func (r *DynamicAttributeResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Check if provider data is available.
	if req.ProviderData == nil {
		log.Println("[DEBUG] DynamicAttribute: ProviderData is nil, returning early.")
		return
	}

	// Cast provider data to your provider type.
	prov, ok := req.ProviderData.(*saviyntProvider)
	if !ok {
		log.Printf("[ERROR] DynamicAttribute: Unexpected Provider Data")
		resp.Diagnostics.AddError("Unexpected Provider Data", "Expected *saviyntProvider")
		return
	}

	// Set the client and token from the provider state.
	r.client = &client.SaviyntClientWrapper{Client: prov.client}
	r.token = prov.accessToken
	// Store username for import functionality
	if prov.client.Username != nil {
		r.username = *prov.client.Username
	}
	log.Printf("[DEBUG] DynamicAttribute: Resource configured successfully.")
}

// SetClient sets the client for testing purposes
func (r *DynamicAttributeResource) SetClient(client client.SaviyntClientInterface) {
	r.client = client
}

// SetToken sets the token for testing purposes
func (r *DynamicAttributeResource) SetToken(token string) {
	r.token = token
}

// SetUsername sets the username for testing purposes
func (r *DynamicAttributeResource) SetUsername(username string) {
	r.username = username
}

// BuildCreateDynamicAttributes builds the dynamic attributes for creation
func (r *DynamicAttributeResource) BuildCreateDynamicAttributes(dynamicAttrMap map[string]Dynamicattribute) []openapi.CreateDynamicAttributesInner {
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
		dynamicAttr.Defaultvalue = util.StringPointerOrEmpty(attr.Defaultvalue)
		dynamicAttr.Required = util.StringPointerOrEmpty(attr.Required)
		// dynamicAttr.Regex = util.StringPointerOrEmpty(attr.Regex)
		dynamicAttr.Attributevalue = util.StringPointerOrEmpty(attr.Attributevalue)
		dynamicAttr.Showonchild = util.StringPointerOrEmpty(attr.Showonchild)
		dynamicAttr.Parentattribute = util.StringPointerOrEmpty(attr.Parentattribute)
		dynamicAttr.Descriptionascsv = util.StringPointerOrEmpty(attr.Descriptionascsv)
		dynamicAttrs = append(dynamicAttrs, *dynamicAttr)
	}

	return dynamicAttrs
}

// CreateDynamicAttribute creates dynamic attributes
func (r *DynamicAttributeResource) CreateDynamicAttribute(ctx context.Context, plan *DynamicAttributeResourceModel) (*openapi.CreateOrUpdateOrDeleteDynamicAttributeResponse, error) {
	log.Printf("[DEBUG] DynamicAttribute: Starting creation for endpoint: %s ", plan.Endpoint.ValueString())

	dynAttrOps := r.dynamicAttributeFactory.CreateDynamicAttributeOperations(r.client.APIBaseURL(), r.token)

	// Convert dynamic attributes map to a slice of Dynamicattribute
	var dynamicAttrMap map[string]Dynamicattribute
	diags := plan.DynamicAttributes.ElementsAs(ctx, &dynamicAttrMap, false)
	if diags.HasError() {
		log.Printf("[ERROR] DynamicAttribute: Failed to process dynamic attributes: %v", diags.Errors())
		return nil, fmt.Errorf("failed to process dynamic attributes: %v", diags.Errors())
	}

	log.Printf("[DEBUG] DynamicAttribute: Processing %d attributes for creation", len(dynamicAttrMap))

	dynamicAttrs := r.BuildCreateDynamicAttributes(dynamicAttrMap)
	createReq := openapi.NewCreateDynamicAttributeRequest(
		plan.Securitysystem.ValueString(),
		plan.Endpoint.ValueString(),
		plan.Updateuser.ValueString(),
		dynamicAttrs,
	)

	// Execute create operation through interface
	log.Printf("[DEBUG] DynamicAttribute: Making API call to create %d dynamic attributes for endpoint: %s",
		len(dynamicAttrs), plan.Endpoint.ValueString())
	createResp, httpResp, err := dynAttrOps.CreateDynamicAttribute(ctx, *createReq)
	if err != nil {
		log.Printf("[ERROR] DynamicAttribute: Problem with the creating function in CreateDynamicAttribute. Error: %v, HTTP Response: %v", err, httpResp)
		return nil, fmt.Errorf("API call failed: %w", err)
	}

	if createResp != nil && createResp.Errorcode == nil {
		log.Println("[ERROR] DynamicAttribute: Unexpected API response - Errorcode is nil")
		return nil, fmt.Errorf("unexpected API response - Errorcode is nil")
	}

	log.Printf("[INFO] DynamicAttribute: Creation API call completed. Error code: %s", util.SafeDeref(createResp.Errorcode))

	log.Printf("[INFO] Dynamic attributes resource created successfully for endpoint: %s", plan.Endpoint.ValueString())

	return createResp, nil
}

// ProcessDynamicAttributeErrorResponse processes error response from any dynamic attribute operation
func (r *DynamicAttributeResource) ProcessDynamicAttributeErrorResponse(plan *DynamicAttributeResourceModel, resp *openapi.CreateOrUpdateOrDeleteDynamicAttributeResponse, operationType string) error {
	if resp.Errorcode != nil && *resp.Errorcode != "0" && operationType == "update" {
		return fmt.Errorf("dynamic Attribute Update Failed - Error: %s, Message: %s", *resp.Errorcode, *resp.Msg)
	}

	if resp.Errorcode != nil && *resp.Errorcode == "1" {
		var errorMessages []string

		// Only update plan fields if plan is provided (not for delete operations)
		if plan != nil {
			plan.ErrorCode = types.StringPointerValue(resp.Errorcode)
			plan.Msg = types.StringPointerValue(resp.Msg)
		}

		if resp.Securitysystem != nil {
			errorMessages = append(errorMessages, fmt.Sprintf("securitysystem: %s", *resp.Securitysystem))
		}

		if resp.Endpoint != nil {
			errorMessages = append(errorMessages, fmt.Sprintf("endpoint: %s", *resp.Endpoint))
		}

		if resp.Updateuser != nil {
			errorMessages = append(errorMessages, fmt.Sprintf("updateuser: %s", *resp.Updateuser))
		}

		if resp.Dynamicattributes != nil {
			rawJSON, err := json.Marshal(resp.Dynamicattributes)
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

		// Add specific message for create operations
		if operationType == "create" && strings.Contains(fullError, "attributename already exists") {
			fullError += "\n\nTry importing the resource or use a different name."
		}

		// Update plan error field if plan is provided
		if plan != nil {
			plan.DynamicAttributesError = types.StringValue(fullError)
		}

		return fmt.Errorf("Dynamic Attribute %s Operation Failed:\n%s", strings.Title(operationType), fullError)
	} else if operationType == "create" && plan != nil {
		plan.DynamicAttributesError = types.StringNull()
	}

	return nil
}

// UpdateModelFromCreateResponse updates the model from create response
func (r *DynamicAttributeResource) UpdateModelFromCreateResponse(plan *DynamicAttributeResourceModel, createResp *openapi.CreateOrUpdateOrDeleteDynamicAttributeResponse, dynamicAttrMap map[string]Dynamicattribute, ctx context.Context) error {
	// Update the dynamic attributes in state with SafeString handling
	updatedAttrs := make(map[string]Dynamicattribute)
	for attrName, attr := range dynamicAttrMap {
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
			Hideonupdate:   util.SafeStringAlt(attr.Hideonupdate.ValueStringPointer(), "false"),
			Action_to_perform_when_parent_attribute_changes: util.SafeString(attr.Action_to_perform_when_parent_attribute_changes.ValueStringPointer()),
			Defaultvalue:     util.SafeString(attr.Defaultvalue.ValueStringPointer()),
			Required:         util.SafeStringAlt(attr.Required.ValueStringPointer(), "false"),
			Showonchild:      util.SafeStringAlt(attr.Showonchild.ValueStringPointer(), "false"),
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
	updatedMap, diags := r.ConvertAttributesToTerraformMap(updatedAttrs, ctx)
	if diags.HasError() {
		return fmt.Errorf("failed to convert updated attributes to map: %v", diags.Errors())
	}

	plan.DynamicAttributes = updatedMap
	plan.ID = types.StringValue("dynamic-attr-" + plan.Endpoint.ValueString())
	plan.Msg = types.StringValue(util.SafeDeref(createResp.Msg))
	plan.ErrorCode = types.StringValue(util.SafeDeref(createResp.Errorcode))
	plan.DynamicAttributesError = types.StringNull()
	plan.Updateuser = types.StringValue(plan.Updateuser.ValueString())

	return nil
}

// ProcessReadErrorResponse processes 412 error response from read operation
func (r *DynamicAttributeResource) ProcessReadErrorResponse(httpResp *http.Response) error {
	var fetchResp map[string]interface{}
	if err := json.NewDecoder(httpResp.Body).Decode(&fetchResp); err != nil {
		return fmt.Errorf("failed to decode error response: %w", err)
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

	// Combine and return error
	fullError := strings.Join(errorMessages, "\n")
	return fmt.Errorf("dynamic Attributes Fetch Failed: %s", fullError)
}

// ReadDynamicAttribute reads dynamic attributes
func (r *DynamicAttributeResource) ReadDynamicAttribute(ctx context.Context, endpointName string) (*openapi.FetchDynamicAttributesResponse, error) {
	log.Printf("[DEBUG] DynamicAttribute: Starting read operation for endpoint: %s", endpointName)

	dynAttrOps := r.dynamicAttributeFactory.CreateDynamicAttributeOperations(r.client.APIBaseURL(), r.token)

	log.Printf("[DEBUG] DynamicAttribute: Making API call to fetch attributes for endpoint: %s", endpointName)
	fetchResp, httpResp, err := dynAttrOps.FetchDynamicAttribute(ctx, endpointName)

	if httpResp != nil && httpResp.StatusCode == 412 {
		log.Printf("[ERROR] DynamicAttribute: Received 412 Precondition Failed for endpoint: %s", endpointName)
		return nil, r.ProcessReadErrorResponse(httpResp)
	}

	if err != nil {
		log.Printf("[ERROR] DynamicAttribute: Problem with the get function in ReadDynamicAttribute. Error: %v", err)
		return nil, fmt.Errorf("fetch API call failed: %w", err)
	}

	log.Printf("[INFO] Dynamic attribute resource read successfully. Endpoint: %s", endpointName)

	return fetchResp, nil
}

// ConvertAttributesToTerraformMap converts attributes map to Terraform Map type
func (r *DynamicAttributeResource) ConvertAttributesToTerraformMap(updatedAttrs map[string]Dynamicattribute, ctx context.Context) (types.Map, diag.Diagnostics) {
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
			"default_value":      types.StringType,
			"required":           types.StringType,
			"attribute_value":    types.StringType,
			"showonchild":        types.StringType,
			"parent_attribute":   types.StringType,
			"description_as_csv": types.StringType,
		},
	}, updatedAttrs)

	return dynamicAttributesMap, diags
}
func (r *DynamicAttributeResource) BuildUpdatedAttributeFromAPI(attrName string, apiAttr openapi.FetchDynamicAttributeResponseInner, isImport bool) Dynamicattribute {
	updatedAttr := Dynamicattribute{
		Attributename:  types.StringValue(attrName),
		Requesttype:    util.SafeStringDatasource(apiAttr.Requesttype),
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
		Defaultvalue:     util.SafeStringDatasource(apiAttr.Defaultvalue),
		Attributevalue:   util.SafeStringPreserveNull(apiAttr.Attributevalue),
		Showonchild:      util.SafeStringDatasource(apiAttr.Showonchild),
		Parentattribute:  util.SafeStringDatasource(apiAttr.Parentattribute),
		Descriptionascsv: util.SafeStringDatasource(apiAttr.Descriptionascsv),
	}

	// Handle attribute type differently for import vs regular read
	if isImport {
		updatedAttr.Attributetype = util.SafeStringDatasource(apiAttr.Attributetype)
	} else {
		updatedAttr.Attributetype = types.StringValue(dynamicattributeutil.TranslateValue(*apiAttr.Attributetype, dynamicattributeutil.AttributeTypeMap))
	}

	return updatedAttr
}

// BuildUpdatedAttributesForImport builds updated attributes for import case
func (r *DynamicAttributeResource) BuildUpdatedAttributesForImport(apiAttrs map[string]openapi.FetchDynamicAttributeResponseInner) map[string]Dynamicattribute {
	updatedAttrs := make(map[string]Dynamicattribute)
	for attrName, apiAttr := range apiAttrs {
		updatedAttrs[attrName] = r.BuildUpdatedAttributeFromAPI(attrName, apiAttr, true)
	}
	log.Printf("[IMPORT] Importing all %d dynamic attributes from API", len(apiAttrs))
	return updatedAttrs
}

// BuildUpdatedAttributesForRead builds updated attributes for regular read case
func (r *DynamicAttributeResource) BuildUpdatedAttributesForRead(currentAttrs map[string]Dynamicattribute, apiAttrs map[string]openapi.FetchDynamicAttributeResponseInner) map[string]Dynamicattribute {
	updatedAttrs := make(map[string]Dynamicattribute)
	removedAttributes := make([]string, 0)

	for attrName := range currentAttrs {
		if apiAttr, exists := apiAttrs[attrName]; exists {
			// Attribute exists in API response - update with API values
			updatedAttrs[attrName] = r.BuildUpdatedAttributeFromAPI(attrName, apiAttr, false)
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

	return updatedAttrs
}

// UpdateModelFromReadResponse updates the model from read response
func (r *DynamicAttributeResource) UpdateModelFromReadResponse(state *DynamicAttributeResourceModel, apiResp *openapi.FetchDynamicAttributesResponse, ctx context.Context) error {
	log.Printf("[DEBUG] Read HTTP Status Code: %d", 200) // Assuming success if we reach here
	log.Printf("[DEBUG] Read API Error code: %+v", util.SafeDeref(apiResp.Errorcode))
	log.Printf("[DEBUG] Read API Message: %+v", util.SafeDeref(apiResp.Msg))

	// Get current state attributes as a map
	currentAttrs := make(map[string]Dynamicattribute)
	if !state.DynamicAttributes.IsNull() {
		diags := state.DynamicAttributes.ElementsAs(ctx, &currentAttrs, false)
		if diags.HasError() {
			return fmt.Errorf("failed to get current attributes: %v", diags.Errors())
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

	// Build updated attributes map using helper functions
	var updatedAttrs map[string]Dynamicattribute
	if len(currentAttrs) == 0 {
		// IMPORT CASE: State is empty, bring in all attributes from API
		updatedAttrs = r.BuildUpdatedAttributesForImport(apiAttrs)
	} else {
		// REGULAR READ CASE: Update existing attributes with API values
		updatedAttrs = r.BuildUpdatedAttributesForRead(currentAttrs, apiAttrs)
	}

	// Convert to types.Map
	dynamicAttributesMap, diags := r.ConvertAttributesToTerraformMap(updatedAttrs, ctx)
	if diags.HasError() {
		return fmt.Errorf("failed to convert attributes to map: %v", diags.Errors())
	}

	state.DynamicAttributes = dynamicAttributesMap
	state.ID = types.StringValue("dynamic-attr-" + state.Endpoint.ValueString())
	state.Endpoint = util.SafeString(state.Endpoint.ValueStringPointer())
	state.Updateuser = util.SafeString(state.Updateuser.ValueStringPointer())
	state.Msg = types.StringValue(util.SafeDeref(apiResp.Msg))
	state.ErrorCode = types.StringValue(util.SafeDeref(apiResp.Errorcode))

	log.Printf("[INFO] DynamicAttribute: Successfully read %d attributes for endpoint: %s", len(updatedAttrs), state.Endpoint.ValueString())
	return nil
}

// BuildUpdateDynamicAttributes builds dynamic attributes for update
func (r *DynamicAttributeResource) BuildUpdateDynamicAttributes(attr Dynamicattribute) openapi.UpdateDynamicAttributesInner {
	attrName := attr.Attributename.ValueString()
	updateAttr := openapi.NewUpdateDynamicAttributesInner(attrName)

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
	updateAttr.Defaultvalue = util.StringPointerOrEmpty(attr.Defaultvalue)
	updateAttr.Required = util.StringPointerOrEmpty(attr.Required)
	updateAttr.Attributevalue = util.StringPointerOrEmpty(attr.Attributevalue)
	updateAttr.Showonchild = util.StringPointerOrEmpty(attr.Showonchild)
	updateAttr.Parentattribute = util.StringPointerOrEmpty(attr.Parentattribute)
	updateAttr.Descriptionascsv = util.StringPointerOrEmpty(attr.Descriptionascsv)

	return *updateAttr
}

// UpdateDynamicAttribute updates dynamic attributes
func (r *DynamicAttributeResource) UpdateDynamicAttribute(ctx context.Context, plan *DynamicAttributeResourceModel) error {
	log.Printf("[DEBUG] DynamicAttribute: Starting update operation for endpoint: %s", plan.Endpoint.ValueString())

	dynAttrOps := r.dynamicAttributeFactory.CreateDynamicAttributeOperations(r.client.APIBaseURL(), r.token)

	// Get planned attributes
	var planAttrs map[string]Dynamicattribute
	diags := plan.DynamicAttributes.ElementsAs(ctx, &planAttrs, false)
	if diags.HasError() {
		log.Printf("[ERROR] DynamicAttribute: Failed to process planned attributes: %v", diags.Errors())
		return fmt.Errorf("failed to process planned attributes: %v", diags.Errors())
	}
	log.Printf("Planned attributes before fetch: %v", len(planAttrs))

	// Get current state from server
	fetchResp, err := r.ReadDynamicAttribute(ctx, plan.Endpoint.ValueString())
	if err != nil {
		log.Printf("[ERROR] DynamicAttribute: Failed to fetch current dynamic attributes: %v", err)
		return fmt.Errorf("failed to fetch current dynamic attributes: %w", err)
	}
	log.Printf("Fetching existing attributes from api complete")

	// Build map of existing attributes
	existingAttrs := make(map[string]bool)
	if fetchResp != nil && fetchResp.Dynamicattributes != nil && fetchResp.Dynamicattributes.ArrayOfFetchDynamicAttributeResponseInner != nil {
		for _, item := range *fetchResp.Dynamicattributes.ArrayOfFetchDynamicAttributeResponseInner {
			if item.Attributename != nil {
				existingAttrs[*item.Attributename] = true
			}
		}
	}
	log.Printf("Length of existing attributes from api: %v", len(existingAttrs))

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
			// Set all fields using existing helper
			*newAttr = r.BuildCreateDynamicAttributes(map[string]Dynamicattribute{attrName: attr})[0]
			newAttrs = append(newAttrs, *newAttr)
		} else {
			// Existing attribute - prepare for update
			updateAttr := r.BuildUpdateDynamicAttributes(attr)
			updateAttrs = append(updateAttrs, updateAttr)
		}
	}
	log.Printf("Number of attributes to create: %v", len(newAttrs))
	log.Printf("Number of attributes to update: %v", len(updateAttrs))

	// First, create new attributes if any
	if len(newAttrs) > 0 {
		createReq := openapi.NewCreateDynamicAttributeRequest(
			plan.Securitysystem.ValueString(),
			plan.Endpoint.ValueString(),
			plan.Updateuser.ValueString(),
			newAttrs,
		)

		createResp, _, err := dynAttrOps.CreateDynamicAttribute(ctx, *createReq)
		if err != nil {
			log.Printf("[ERROR] DynamicAttribute: Failed to create new dynamic attributes: %v", err)
			return fmt.Errorf("failed to create new dynamic attributes: %w", err)
		}
		err = r.ProcessDynamicAttributeErrorResponse(plan, createResp, "create")
		if err != nil {
			return err
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

		// Execute update operation through interface
		updateResp, _, err := dynAttrOps.UpdateDynamicAttribute(ctx, *updateReq)
		if err != nil {
			log.Printf("[ERROR]: Problem with the updating function in UpdateDynamicAttribute. Error: %v", err)
			return fmt.Errorf("failed to update dynamic attributes: %w", err)
		}

		err = r.ProcessDynamicAttributeErrorResponse(plan, updateResp, "update")
		if err != nil {
			return err
		}
	}

	log.Printf("[INFO] Dynamic attribute resource updated successfully for endpoint: %s", plan.Endpoint.ValueString())

	return nil
}

// UpdateModelFromUpdateResponse updates the model from update response
func (r *DynamicAttributeResource) UpdateModelFromUpdateResponse(plan *DynamicAttributeResourceModel, planAttrs map[string]Dynamicattribute, ctx context.Context) error {
	// Fetch updated state after update
	postUpdateFetchResp, err := r.ReadDynamicAttribute(ctx, plan.Endpoint.ValueString())
	if err != nil {
		// Handle 412 error case
		if strings.Contains(err.Error(), "Dynamic Attributes Fetch Failed") {
			plan.DynamicAttributesError = types.StringValue(err.Error())
			return fmt.Errorf("dynamic Attributes Fetch after update Failed: %s", err.Error())
		}
		return fmt.Errorf("failed to fetch updated dynamic attributes: %w", err)
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

	for attrName := range planAttrs {
		if apiAttr, exists := postUpdateApiAttrs[attrName]; exists {
			// Reuse the existing helper function
			updatedAttrs[attrName] = r.BuildUpdatedAttributeFromAPI(attrName, apiAttr, false)
		}
	}

	log.Printf("[INFO] Processed %d dynamic attributes in state refresh", len(updatedAttrs))

	// Convert the updated attributes map to a Terraform Map type
	dynamicAttributesMap, diags := r.ConvertAttributesToTerraformMap(updatedAttrs, ctx)
	if diags.HasError() {
		return fmt.Errorf("failed to convert updated attributes to map: %v", diags.Errors())
	}

	plan.DynamicAttributes = dynamicAttributesMap
	plan.ID = types.StringValue("dynamic-attr-" + plan.Endpoint.ValueString())
	plan.DynamicAttributesError = types.StringNull()
	plan.Endpoint = util.SafeString(plan.Endpoint.ValueStringPointer())
	plan.Securitysystem = util.SafeString(plan.Securitysystem.ValueStringPointer())
	plan.Updateuser = util.SafeString(plan.Updateuser.ValueStringPointer())
	plan.Msg = types.StringValue(util.SafeDeref(postUpdateFetchResp.Msg))
	plan.ErrorCode = types.StringValue(util.SafeDeref(postUpdateFetchResp.Errorcode))

	return nil
}

// DeleteDynamicAttribute deletes dynamic attributes
func (r *DynamicAttributeResource) DeleteDynamicAttribute(ctx context.Context, state *DynamicAttributeResourceModel) error {
	log.Printf("[DEBUG] DynamicAttribute: Starting delete operation for endpoint: %s", state.Endpoint.ValueString())

	dynAttrOps := r.dynamicAttributeFactory.CreateDynamicAttributeOperations(r.client.APIBaseURL(), r.token)

	// Get attribute names from the map
	var attributeNames []string
	var stateAttrs map[string]Dynamicattribute

	// Properly unmarshal the map
	diags := state.DynamicAttributes.ElementsAs(ctx, &stateAttrs, false)
	if diags.HasError() {
		return fmt.Errorf("failed to process state attributes: %v", diags.Errors())
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

	// Execute delete operation through interface
	deleteResp, httpResp, err := dynAttrOps.DeleteDynamicAttribute(ctx, deleteReq)
	if err != nil {
		log.Printf("[ERROR] DynamicAttribute: Problem with the deleting function in DeleteDynamicAttribute. Error: %v", err)
		return err
	}

	err = r.ProcessDynamicAttributeErrorResponse(nil, deleteResp, "delete")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Delete HTTP Status Code: %d", httpResp.StatusCode)
	log.Printf("[DEBUG] Delete API Response: %+v", deleteResp)
	return nil
}

// ValidateImportKey validates that the endpoint exists and returns its associated security system
func (r *DynamicAttributeResource) ValidateImportKey(ctx context.Context, endpointName string) (string, error) {
	endpointOps := r.dynamicAttributeFactory.CreateEndpointOperations(r.client.APIBaseURL(), r.token)

	// Create the request object for GetEndpoints
	reqParams := endpoint.GetEndpointsRequest{}
	reqParams.SetEndpointname(endpointName)

	endpointResp, _, err := endpointOps.GetEndpoints(ctx, reqParams)
	if err != nil {
		log.Printf("[ERROR] DynamicAttribute: API Call failed: Failed to fetch endpoint details. Error: %v", err)
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

func (r *DynamicAttributeResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan DynamicAttributeResourceModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		log.Println("[ERROR] DynamicAttribute: Error getting plan data")
		return
	}

	// Convert dynamic attributes map to a slice of Dynamicattribute
	var dynamicAttrMap map[string]Dynamicattribute
	resp.Diagnostics.Append(plan.DynamicAttributes.ElementsAs(ctx, &dynamicAttrMap, false)...)
	if resp.Diagnostics.HasError() {
		log.Println("[ERROR] DynamicAttribute: Failed to process dynamic attributes")
		return
	}

	createResp, err := r.CreateDynamicAttribute(ctx, &plan)
	if err != nil {
		log.Printf("[ERROR] DynamicAttribute: Error Creating Dynamic Attribute: %v", err)
		resp.Diagnostics.AddError("Error Creating Dynamic Attribute", fmt.Sprintf("API Error: %v", err))
		return
	}

	// Process error response
	err = r.ProcessDynamicAttributeErrorResponse(&plan, createResp, "create")
	if err != nil {
		log.Printf("[ERROR] DynamicAttribute: Dynamic Attribute Create Operation Failed: %v", err)
		resp.Diagnostics.AddError("Dynamic Attribute Create Operation Failed", err.Error())
		return
	}

	// Update model from successful response
	err = r.UpdateModelFromCreateResponse(&plan, createResp, dynamicAttrMap, ctx)
	if err != nil {
		log.Printf("[ERROR] DynamicAttribute: Failed to update model from response: %v", err)
		resp.Diagnostics.AddError("Failed to update model from response", err.Error())
		return
	}

	// Set the final state
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
	log.Printf("[DEBUG] DynamicAttribute: Successfully created dynamic attributes for endpoint: %s", plan.Endpoint.ValueString())
}

func (r *DynamicAttributeResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state DynamicAttributeResourceModel

	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	apiResp, err := r.ReadDynamicAttribute(ctx, state.Endpoint.ValueString())
	if err != nil {
		// Handle 412 error case by setting error in state
		if strings.Contains(err.Error(), "Dynamic Attributes Fetch Failed") {
			state.Securitysystem = types.StringValue("")
			state.Endpoint = types.StringValue("")
			state.DynamicAttributesError = types.StringValue(err.Error())

			// Persist in state
			diags := resp.State.Set(ctx, &state)
			resp.Diagnostics.Append(diags...)
			if resp.Diagnostics.HasError() {
				return
			}

			// Display error to user
			resp.Diagnostics.AddError("Dynamic Attributes Fetch Failed", err.Error())
			return
		}
		resp.Diagnostics.AddError("API Read Failed In Read Block", err.Error())
		log.Printf("[ERROR] DynamicAttribute: API Read Failed In Read Block: %v", err)
		return
	}

	err = r.UpdateModelFromReadResponse(&state, apiResp, ctx)
	if err != nil {
		log.Printf("[ERROR] DynamicAttribute: Failed to update model from read response: %v", err)
		resp.Diagnostics.AddError("Failed to update model from read response", err.Error())
		return
	}

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
}

func (r *DynamicAttributeResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan DynamicAttributeResourceModel

	// Get the plan
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Get planned attributes
	var planAttrs map[string]Dynamicattribute
	resp.Diagnostics.Append(plan.DynamicAttributes.ElementsAs(ctx, &planAttrs, false)...)
	if resp.Diagnostics.HasError() {
		return
	}

	err := r.UpdateDynamicAttribute(ctx, &plan)
	if err != nil {
		log.Printf("[ERROR] DynamicAttribute: API Update Failed: %v", err)
		resp.Diagnostics.AddError("API Update Failed", err.Error())
		return
	}

	err = r.UpdateModelFromUpdateResponse(&plan, planAttrs, ctx)
	if err != nil {
		resp.Diagnostics.AddError("Failed to update model from response", err.Error())
		return
	}

	// Set final state
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *DynamicAttributeResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state DynamicAttributeResourceModel

	stateRetrievalDiagnostics := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(stateRetrievalDiagnostics...)
	if resp.Diagnostics.HasError() {
		return
	}

	err := r.DeleteDynamicAttribute(ctx, &state)
	if err != nil {
		log.Printf("[ERROR] DynamicAttribute: API Delete Failed: %v", err)
		resp.Diagnostics.AddError("API Delete Failed", err.Error())
		return
	}
}

func (r *DynamicAttributeResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	endpointName := req.ID

	// Set the endpoint from import ID
	resp.State.SetAttribute(ctx, path.Root("endpoint"), endpointName)

	securitySystem, err := r.ValidateImportKey(ctx, endpointName)
	if err != nil {
		log.Printf("[ERROR] DynamicAttribute: Import key validation failed. %v", err)
		resp.Diagnostics.AddError("Import Key validation failed: ", err.Error())
		return
	}

	resp.State.SetAttribute(ctx, path.Root("security_system"), securitySystem)
	resp.State.SetAttribute(ctx, path.Root("update_user"), r.username)
}
