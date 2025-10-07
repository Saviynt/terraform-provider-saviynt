// Copyright (c) 2025 Saviynt Inc.
// SPDX-License-Identifier: MPL-2.0

// saviynt_entitlements_datasource retrieves entitlement details from the Saviynt Security Manager.
// The data source supports filtering entitlements by various criteria.
package provider

import (
	"context"
	"fmt"
	"net/http"
	"terraform-provider-Saviynt/internal/client"
	"terraform-provider-Saviynt/util"

	"terraform-provider-Saviynt/util/errorsutil"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	openapi "github.com/saviynt/saviynt-api-go-client/entitlements"
)

type entitlementDataSource struct {
	client             client.SaviyntClientInterface
	token              string
	provider           client.SaviyntProviderInterface
	entitlementFactory client.EntitlementFactoryInterface
}

var _ datasource.DataSource = &entitlementDataSource{}
var _ datasource.DataSourceWithConfigure = &entitlementDataSource{}

// NewEntitlementDataSource creates a new entitlement data source with default factory
func NewEntitlementDataSource() datasource.DataSource {
	return &entitlementDataSource{
		entitlementFactory: &client.DefaultEntitlementFactory{},
	}
}

// NewEntitlementDataSourceWithFactory creates a new entitlement data source with custom factory
// Used primarily for testing with mock factories
func NewEntitlementDataSourceWithFactory(factory client.EntitlementFactoryInterface) datasource.DataSource {
	return &entitlementDataSource{
		entitlementFactory: factory,
	}
}

// SetClient sets the client for testing purposes
func (d *entitlementDataSource) SetClient(client client.SaviyntClientInterface) {
	d.client = client
}

// SetToken sets the token for testing purposes
func (d *entitlementDataSource) SetToken(token string) {
	d.token = token
}

// SetProvider sets the provider for testing purposes
func (d *entitlementDataSource) SetProvider(provider client.SaviyntProviderInterface) {
	d.provider = provider
}

type EntitlementDataSourceModel struct {
	Results               []EntitlementDetails `tfsdk:"results"`
	TotalEntitlementCount types.Int32          `tfsdk:"total_entitlement_count"`
	EntitlementsCount     types.Int32          `tfsdk:"entitlements_count"`
	ErrorCode             types.String         `tfsdk:"error_code"`
	Message               types.String         `tfsdk:"message"`
	Endpoint              types.String         `tfsdk:"endpoint"`
	Entitlementtype       types.String         `tfsdk:"entitlementtype"`
	EntitlementValue      types.String         `tfsdk:"entitlement_value"`
	EntQuery              types.String         `tfsdk:"ent_query"`
	Authenticate          types.Bool           `tfsdk:"authenticate"`
}

type EntitlementDetails struct {
	EntitlementValuekey   types.String `tfsdk:"entitlement_valuekey"`
	EntitlementValue      types.String `tfsdk:"entitlement_value"`
	Endpoint              types.String `tfsdk:"endpoint"`
	EndpointKey           types.Int32  `tfsdk:"endpoint_key"`
	EntitlementType       types.String `tfsdk:"entitlement_type"`
	EntitlementTypeName   types.String `tfsdk:"entitlement_type_name"`
	EntitlementTypeKey    types.Int32  `tfsdk:"entitlement_type_key"`
	Displayname           types.String `tfsdk:"displayname"`
	Description           types.String `tfsdk:"description"`
	Status                types.String `tfsdk:"status"`
	Risk                  types.String `tfsdk:"risk"`
	Priority              types.String `tfsdk:"priority"`
	Soxcritical           types.String `tfsdk:"soxcritical"`
	Syscritical           types.String `tfsdk:"syscritical"`
	Privileged            types.String `tfsdk:"privileged"`
	Confidentiality       types.String `tfsdk:"confidentiality"`
	Module                types.String `tfsdk:"module"`
	Access                types.String `tfsdk:"access"`
	EntitlementGlossary   types.String `tfsdk:"entitlement_glossary"`
	RequestForm           types.String `tfsdk:"request_form"`
	Createdate            types.String `tfsdk:"createdate"`
	Updatedate            types.String `tfsdk:"updatedate"`
	ChildEntitlementCount types.Int32  `tfsdk:"child_entitlement_count"`
	EntitlementOwner      types.Map    `tfsdk:"entitlement_owner"`
	EntitlementMapDetails types.List   `tfsdk:"entitlement_map_details"`
}

func (d *entitlementDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "saviynt_entitlement_datasource"
}

func (d *entitlementDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: util.EntitlementDataSourceDescription,
		Attributes: map[string]schema.Attribute{
			"results": schema.ListAttribute{
				ElementType: types.ObjectType{
					AttrTypes: map[string]attr.Type{
						"entitlement_valuekey":    types.StringType,
						"entitlement_value":       types.StringType,
						"endpoint":                types.StringType,
						"endpoint_key":            types.Int32Type,
						"entitlement_type":        types.StringType,
						"entitlement_type_name":   types.StringType,
						"entitlement_type_key":    types.Int32Type,
						"displayname":             types.StringType,
						"description":             types.StringType,
						"status":                  types.StringType,
						"risk":                    types.StringType,
						"priority":                types.StringType,
						"soxcritical":             types.StringType,
						"syscritical":             types.StringType,
						"privileged":              types.StringType,
						"confidentiality":         types.StringType,
						"module":                  types.StringType,
						"access":                  types.StringType,
						"entitlement_glossary":    types.StringType,
						"request_form":            types.StringType,
						"createdate":              types.StringType,
						"updatedate":              types.StringType,
						"child_entitlement_count": types.Int32Type,
						"entitlement_owner":       types.MapType{ElemType: types.ListType{ElemType: types.StringType}},
						"entitlement_map_details": types.ListType{ElemType: types.ObjectType{
							AttrTypes: map[string]attr.Type{
								"primary_ent_type":          types.StringType,
								"request_filter":            types.BoolType,
								"add_dependent_task":        types.BoolType,
								"remove_dependent_ent_task": types.BoolType,
								"exclude_entitlement":       types.BoolType,
								"primary":                   types.StringType,
								"primary_ent_key":           types.StringType,
								"export_primary":            types.StringType,
								"description":               types.StringType,
							},
						}},
					},
				},
				Computed:    true,
				Description: "List of entitlements",
			},
			"total_entitlement_count": schema.Int32Attribute{
				Computed:    true,
				Description: "Total number of entitlements",
			},
			"entitlements_count": schema.Int32Attribute{
				Computed:    true,
				Description: "Number of entitlements returned",
			},
			"error_code": schema.StringAttribute{
				Computed:    true,
				Description: "Error code from API response",
			},
			"message": schema.StringAttribute{
				Computed:    true,
				Description: "Message from API response",
			},
			"endpoint": schema.StringAttribute{
				Optional:    true,
				Description: "Filter by endpoint name",
			},
			"entitlementtype": schema.StringAttribute{
				Optional:    true,
				Description: "Filter by entitlement type",
			},
			"entitlement_value": schema.StringAttribute{
				Optional:    true,
				Description: "Filter by entitlement value",
			},
			"ent_query": schema.StringAttribute{
				Optional:    true,
				Description: "SQL-like query for filtering entitlements",
			},
			"authenticate": schema.BoolAttribute{
				Required:    true,
				Description: "Whether to authenticate and return sensitive data",
			},
		},
	}
}

func (d *entitlementDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	tflog.Debug(ctx, "Starting entitlement datasource configuration")

	// Check if provider data is available.
	if req.ProviderData == nil {
		tflog.Debug(ctx, "ProviderData is nil, returning early")
		return
	}

	// Cast provider data to your provider type.
	prov, ok := req.ProviderData.(*SaviyntProvider)
	if !ok {
		tflog.Error(ctx, "Provider configuration failed - unexpected provider data type")
		resp.Diagnostics.AddError(
			"Unexpected Provider Data",
			"Expected *SaviyntProvider, got different type",
		)
		return
	}

	// Set the client and token from the provider state using interface wrapper.
	d.client = &client.SaviyntClientWrapper{Client: prov.client}
	d.token = prov.accessToken
	d.provider = &client.SaviyntProviderWrapper{Provider: prov}

	tflog.Debug(ctx, "Entitlement datasource configured successfully")
}

func (d *entitlementDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var state EntitlementDataSourceModel

	tflog.Debug(ctx, "Starting entitlement datasource read operation")

	// Extract configuration from request
	configDiagnostics := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(configDiagnostics...)
	if resp.Diagnostics.HasError() {
		tflog.Error(ctx, "Failed to get config from request")
		resp.Diagnostics.AddError(
			"Configuration Error",
			"Unable to extract Terraform configuration from request",
		)
		return
	}

	// Execute API call to get entitlement details
	apiResp, err := d.ReadEntitlementDetails(ctx, &state)
	if err != nil {
		tflog.Error(ctx, "Failed to read entitlement details", map[string]interface{}{
			"error": err.Error(),
		})
		resp.Diagnostics.AddError("API Call Failed", fmt.Sprintf("Error: %v", err))
		return
	}

	// Check if no entitlements were found and add user-visible warning
	if apiResp != nil && apiResp.Entitlementdetails != nil && len(apiResp.Entitlementdetails) == 0 {
		resp.Diagnostics.AddWarning(
			"No Entitlements Found",
			fmt.Sprintf("No entitlements found for the specified criteria. Retrieved count: %d",
				func() int32 {
					if apiResp.EntitlementsCount != nil {
						return *apiResp.EntitlementsCount
					}
					return 0
				}()),
		)
	}

	// Map API response to state
	d.UpdateModelFromAPIResponse(&state, apiResp)

	// Handle authentication logic for results
	d.HandleAuthenticationLogic(&state, resp)

	// Set final state
	stateDiagnostics := resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(stateDiagnostics...)
	if resp.Diagnostics.HasError() {
		tflog.Error(ctx, "Failed to set state")
		resp.Diagnostics.AddError(
			"State Update Error",
			"Unable to update Terraform state for entitlement datasource",
		)
		return
	}

	tflog.Debug(ctx, "Entitlement datasource read operation completed successfully")
}

// ReadEntitlementDetails retrieves entitlement details from Saviynt API
// Handles parameter preparation and API call execution using factory pattern
func (d *entitlementDataSource) ReadEntitlementDetails(ctx context.Context, state *EntitlementDataSourceModel) (*openapi.GetEntitlementResponse, error) {
	tflog.Debug(ctx, "Starting entitlement API call")

	// Prepare API request
	getReq := openapi.GetEntitlementRequest{}

	if !state.Endpoint.IsNull() {
		getReq.Endpoint = util.StringPointerOrEmpty(state.Endpoint)
	}
	if !state.Entitlementtype.IsNull() {
		getReq.Entitlementtype = util.StringPointerOrEmpty(state.Entitlementtype)
	}
	if !state.EntitlementValue.IsNull() {
		getReq.EntitlementValue = util.StringPointerOrEmpty(state.EntitlementValue)
	}
	if !state.EntQuery.IsNull() {
		getReq.EntQuery = util.StringPointerOrEmpty(state.EntQuery)
	}
	// Always include owner rank information
	getReq.Entownerwithrank = util.StringPtr("true")
	getReq.Returnentitlementmap = util.StringPtr("true")

	tflog.Debug(ctx, fmt.Sprintf("Executing API request: %+v", getReq))

	// Execute API call with retry logic
	var getResp *openapi.GetEntitlementResponse
	var finalHttpResp *http.Response
	err := d.provider.AuthenticatedAPICallWithRetry(ctx, "read_entitlement_datasource", func(token string) error {
		entitlementOps := d.entitlementFactory.CreateEntitlementOperations(d.client.APIBaseURL(), token)
		resp, hResp, err := entitlementOps.GetEntitlements(ctx, getReq)
		if hResp != nil && hResp.StatusCode == 401 {
			return fmt.Errorf("401 unauthorized")
		}
		getResp = resp
		finalHttpResp = hResp // Update on every call including retries
		return err
	})

	if err != nil {
		tflog.Error(ctx, "API call failed", map[string]interface{}{
			"error": err.Error(),
		})
		err = errorsutil.HandleHTTPError(finalHttpResp, err, "Read")
		return nil, fmt.Errorf("Entitlement Datasource: API call failed: %w", err)
	}

	// Log the API response for debugging
	tflog.Debug(ctx, "API Response received", map[string]interface{}{
		"response": fmt.Sprintf("%+v", getResp),
	})

	if getResp != nil && getResp.ErrorCode != nil && *getResp.ErrorCode != "0" {
		errorCode := util.SafeDeref(getResp.ErrorCode)
		msg := util.SafeDeref(getResp.Msg)
		return nil, fmt.Errorf("Entitlement Datasource: API returned error code: %s and error message: %s", errorCode, msg)
	}

	tflog.Debug(ctx, "Entitlement: API call successful")

	return getResp, nil
}

// UpdateModelFromAPIResponse maps API response data to the Terraform state model
func (d *entitlementDataSource) UpdateModelFromAPIResponse(state *EntitlementDataSourceModel, apiResp *openapi.GetEntitlementResponse) {
	// Map basic response fields
	d.MapBasicResponseFields(state, apiResp)

	// Process entitlement details if available
	if apiResp != nil && len(apiResp.Entitlementdetails) > 0 {
		var results []EntitlementDetails
		for _, ent := range apiResp.Entitlementdetails {
			result := d.MapEntitlementDetails(&ent)
			results = append(results, result)
		}
		state.Results = results
	}
}

// MapBasicResponseFields maps basic response fields from API response to state model
func (d *entitlementDataSource) MapBasicResponseFields(state *EntitlementDataSourceModel, apiResp *openapi.GetEntitlementResponse) {
	state.ErrorCode = util.SafeString(apiResp.ErrorCode)
	state.Message = util.SafeString(apiResp.Msg)
	state.TotalEntitlementCount = util.SafeInt32(apiResp.TotalEntitlementCount)
	state.EntitlementsCount = util.SafeInt32(apiResp.EntitlementsCount)
}

// MapEntitlementDetails maps individual entitlement details from API response to state model
func (d *entitlementDataSource) MapEntitlementDetails(ent *openapi.GetEntitlementResponseEntitlementdetailsInner) EntitlementDetails {
	result := EntitlementDetails{
		EntitlementValuekey:   util.SafeString(ent.EntitlementValuekey),
		EntitlementValue:      util.SafeString(ent.EntitlementValue),
		Endpoint:              util.SafeString(ent.Endpoint),
		EndpointKey:           util.SafeInt32(ent.EndpointKey),
		EntitlementType:       util.SafeString(ent.EntitlementType),
		EntitlementTypeName:   util.SafeString(ent.EntitlementTypeName),
		EntitlementTypeKey:    util.SafeInt32(ent.EntitlementTypeKey),
		Displayname:           util.SafeString(ent.Displayname),
		Description:           util.SafeString(ent.Description),
		Status:                util.SafeString(ent.Status),
		Risk:                  util.SafeString(ent.Risk),
		Priority:              util.SafeString(ent.Priority),
		Soxcritical:           util.SafeString(ent.Soxcritical),
		Syscritical:           util.SafeString(ent.Syscritical),
		Privileged:            util.SafeString(ent.Priviliged),
		Confidentiality:       util.SafeString(ent.Confidentiality),
		Module:                util.SafeString(ent.Module),
		Access:                util.SafeString(ent.Access),
		EntitlementGlossary:   util.SafeString(ent.EntitlementGlossary),
		RequestForm:           util.SafeString(ent.RequestForm),
		Createdate:            util.SafeString(ent.Createdate),
		Updatedate:            util.SafeString(ent.Updatedate),
		ChildEntitlementCount: util.SafeInt32(ent.ChildEntitlementCount),
	}

	// Handle entitlement owners
	d.MapEntitlementOwner(ent, &result)

	// Handle entitlement map details
	d.MapEntitlementMapDetails(ent, &result)

	return result
}

// MapEntitlementOwner maps entitlement owner information from API response to state model
func (d *entitlementDataSource) MapEntitlementOwner(ent *openapi.GetEntitlementResponseEntitlementdetailsInner, result *EntitlementDetails) {
	if ent.EntitlementOwner != nil {
		if ent.EntitlementOwner.MapmapOfStringarrayOfString != nil {
			// Handle ranked owners (map of rank to array of owners)
			result.EntitlementOwner, _ = types.MapValueFrom(context.Background(), types.ListType{ElemType: types.StringType}, *ent.EntitlementOwner.MapmapOfStringarrayOfString)
		} else if ent.EntitlementOwner.ArrayOfString != nil {
			// Handle simple array of owners
			ownerMap := map[string][]string{"owners": *ent.EntitlementOwner.ArrayOfString}
			result.EntitlementOwner, _ = types.MapValueFrom(context.Background(), types.ListType{ElemType: types.StringType}, ownerMap)
		} else if ent.EntitlementOwner.String != nil {
			// Handle string case (including empty string)
			if *ent.EntitlementOwner.String == "" {
				result.EntitlementOwner = types.MapNull(types.ListType{ElemType: types.StringType})
			} else {
				ownerMap := map[string][]string{"owners": {*ent.EntitlementOwner.String}}
				result.EntitlementOwner, _ = types.MapValueFrom(context.Background(), types.ListType{ElemType: types.StringType}, ownerMap)
			}
		} else {
			// Handle empty cases (empty object {} or empty array [])
			result.EntitlementOwner = types.MapNull(types.ListType{ElemType: types.StringType})
		}
	} else {
		result.EntitlementOwner = types.MapNull(types.ListType{ElemType: types.StringType})
	}
}

// MapEntitlementMapDetails maps entitlement map details from API response to state model
func (d *entitlementDataSource) MapEntitlementMapDetails(ent *openapi.GetEntitlementResponseEntitlementdetailsInner, result *EntitlementDetails) {
	if len(ent.EntitlementMapDetails) > 0 {
		mapDetailsAttrTypes := map[string]attr.Type{
			"primary_ent_type":          types.StringType,
			"request_filter":            types.BoolType,
			"add_dependent_task":        types.BoolType,
			"remove_dependent_ent_task": types.BoolType,
			"exclude_entitlement":       types.BoolType,
			"primary":                   types.StringType,
			"primary_ent_key":           types.StringType,
			"export_primary":            types.StringType,
			"description":               types.StringType,
		}

		var mapDetailsList []attr.Value
		for _, mapDetail := range ent.EntitlementMapDetails {
			mapDetailAttrs := map[string]attr.Value{
				"primary_ent_type": types.StringValue(func() string {
					if mapDetail.PrimaryEntType != nil {
						return *mapDetail.PrimaryEntType
					}
					return ""
				}()),
				"request_filter":            types.BoolValue(mapDetail.RequestFilter != nil && *mapDetail.RequestFilter),
				"add_dependent_task":        types.BoolValue(mapDetail.AddDependentTask != nil && *mapDetail.AddDependentTask),
				"remove_dependent_ent_task": types.BoolValue(mapDetail.RemoveDependentEntTask != nil && *mapDetail.RemoveDependentEntTask),
				"exclude_entitlement":       types.BoolValue(mapDetail.ExcludeEntitlement != nil && *mapDetail.ExcludeEntitlement),
				"primary": types.StringValue(func() string {
					if mapDetail.Primary != nil {
						return *mapDetail.Primary
					}
					return ""
				}()),
				"primary_ent_key": types.StringValue(func() string {
					if mapDetail.PrimaryEntKey != nil {
						return *mapDetail.PrimaryEntKey
					}
					return ""
				}()),
				"export_primary": types.StringValue(func() string {
					if mapDetail.ExportPrimary != nil {
						return *mapDetail.ExportPrimary
					}
					return ""
				}()),
				"description": types.StringValue(func() string {
					if mapDetail.Description != nil {
						return *mapDetail.Description
					}
					return ""
				}()),
			}
			mapDetailObj, _ := types.ObjectValue(mapDetailsAttrTypes, mapDetailAttrs)
			mapDetailsList = append(mapDetailsList, mapDetailObj)
		}
		result.EntitlementMapDetails, _ = types.ListValue(types.ObjectType{AttrTypes: mapDetailsAttrTypes}, mapDetailsList)
	} else {
		result.EntitlementMapDetails = types.ListNull(types.ObjectType{AttrTypes: map[string]attr.Type{
			"primary_ent_type":          types.StringType,
			"request_filter":            types.BoolType,
			"add_dependent_task":        types.BoolType,
			"remove_dependent_ent_task": types.BoolType,
			"exclude_entitlement":       types.BoolType,
			"primary":                   types.StringType,
			"primary_ent_key":           types.StringType,
			"export_primary":            types.StringType,
			"description":               types.StringType,
		}})
	}
}

// HandleAuthenticationLogic processes the authenticate flag to control sensitive data visibility
// When authenticate=false, results are removed from state to prevent sensitive data exposure
// When authenticate=true, all entitlement results are returned in state
func (d *entitlementDataSource) HandleAuthenticationLogic(state *EntitlementDataSourceModel, resp *datasource.ReadResponse) {
	if !state.Authenticate.IsNull() && !state.Authenticate.IsUnknown() {
		if state.Authenticate.ValueBool() {
			tflog.Info(context.Background(), "Authentication enabled - returning all entitlement details")
			resp.Diagnostics.AddWarning(
				"Authentication Enabled",
				"`authenticate` is true; all entitlement details will be returned in state.",
			)
		} else {
			tflog.Info(context.Background(), "Authentication disabled - removing entitlement details from state")
			resp.Diagnostics.AddWarning(
				"Authentication Disabled",
				"`authenticate` is false; entitlement details will be removed from state.",
			)
			state.Results = nil
		}
	}
}
