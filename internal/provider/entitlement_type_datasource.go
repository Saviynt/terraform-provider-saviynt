// Copyright (c) 2025 Saviynt Inc.
// SPDX-License-Identifier: MPL-2.0

// saviynt_entitlement_type_datasource retrieves entitlement type details from the Saviynt Security Manager.
// The data source supports a single Read operation to look up existing entitlement types by name or endpoint.
package provider

import (
	"context"
	"encoding/json"
	"fmt"
	"terraform-provider-Saviynt/internal/client"
	"terraform-provider-Saviynt/util"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	openapi "github.com/saviynt/saviynt-api-go-client/entitlementtype"
)

type entitlementTypeDataSource struct {
	client                 client.SaviyntClientInterface
	token                  string
	entitlementTypeFactory client.EntitlementTypeFactoryInterface
}

var _ datasource.DataSource = &entitlementTypeDataSource{}
var _ datasource.DataSourceWithConfigure = &entitlementTypeDataSource{}

func NewEntitlementTypeDataSource() datasource.DataSource {
	return &entitlementTypeDataSource{
		entitlementTypeFactory: &client.DefaultEntitlementTypeFactory{},
	}
}

// NewEntitlementTypeDataSourceWithFactory creates a new entitlement type data source with custom factory
// Used primarily for testing with mock factories
func NewEntitlementTypeDataSourceWithFactory(factory client.EntitlementTypeFactoryInterface) datasource.DataSource {
	return &entitlementTypeDataSource{
		entitlementTypeFactory: factory,
	}
}

// SetClient sets the client for testing purposes
func (d *entitlementTypeDataSource) SetClient(client client.SaviyntClientInterface) {
	d.client = client
}

// SetToken sets the token for testing purposes
func (d *entitlementTypeDataSource) SetToken(token string) {
	d.token = token
}

type EntitlementTypeDataSourceModel struct {
	DisplayCount    types.Int64       `tfsdk:"display_count"`
	ErrorCode       types.String      `tfsdk:"error_code"`
	TotalCount      types.Int64       `tfsdk:"total_count"`
	Msg             types.String      `tfsdk:"msg"`
	EntitlementName types.String      `tfsdk:"entitlement_name"`
	EndpointName    types.String      `tfsdk:"endpoint_name"`
	Offset          types.String      `tfsdk:"offset"`
	Max             types.String      `tfsdk:"max"`
	Authenticate    types.Bool        `tfsdk:"authenticate"`
	Results         []EntitlementType `tfsdk:"results"`
}

type EntitlementType struct {
	OrderIndex                   types.String `tfsdk:"order_index"`
	Workflow                     types.String `tfsdk:"workflow"`
	EnableEntitlementToRoleSync  types.Bool   `tfsdk:"enable_entitlement_to_role_sync"`
	RequiredInRequest            types.String `tfsdk:"required_in_request"`
	RequiredInServiceRequest     types.String `tfsdk:"required_in_service_request"`
	DisplayName                  types.String `tfsdk:"display_name"`
	SecuritySystem               types.String `tfsdk:"security_system"`
	AvailableQueryServiceAccount types.String `tfsdk:"available_query_service_account"`
	CreateTaskAction             types.String `tfsdk:"create_task_action"`
	ArsRequestableEntitlementSQL types.String `tfsdk:"ars_requestable_entitlement_sql_query"`
	HierarchyRequired            types.String `tfsdk:"hierarchy_required"`
	EntitlementName              types.String `tfsdk:"entitlement_name"`
	EntitlementTypeKey           types.Int64  `tfsdk:"entitlement_type_key"`
	EndpointName                 types.String `tfsdk:"endpoint_name"`
	EndpointKey                  types.Int64  `tfsdk:"endpoint_key"`
	RequestOption                types.String `tfsdk:"request_option"`
	EnableProvisioningPriority   types.String `tfsdk:"enable_provisioning_priority"`
	ExcludeRuleAssignedEnts      types.String `tfsdk:"exclude_rule_assigned_ents_in_request"`
	ShowOnChild                  types.String `tfsdk:"show_on_child"`
	ShowEntTypeOn                types.String `tfsdk:"show_ent_type_on"`
	Recon                        types.String `tfsdk:"recon"`
	ArsSelectedEntitlementSQL    types.String `tfsdk:"ars_selected_entitlement_sql_query"`
	Certifiable                  types.String `tfsdk:"certifiable"`
	EntitlementDescription       types.String `tfsdk:"entitlement_description"`
	RequestDatesConfJson         types.String `tfsdk:"request_dates_conf_json"`
	SelectedQueryServiceAccount  types.String `tfsdk:"selected_query_service_account"`

	// Labels for custom properties
	CustomProperty1Label  types.String `tfsdk:"customproperty1_label"`
	CustomProperty2Label  types.String `tfsdk:"customproperty2_label"`
	CustomProperty3Label  types.String `tfsdk:"customproperty3_label"`
	CustomProperty4Label  types.String `tfsdk:"customproperty4_label"`
	CustomProperty5Label  types.String `tfsdk:"customproperty5_label"`
	CustomProperty6Label  types.String `tfsdk:"customproperty6_label"`
	CustomProperty7Label  types.String `tfsdk:"customproperty7_label"`
	CustomProperty8Label  types.String `tfsdk:"customproperty8_label"`
	CustomProperty9Label  types.String `tfsdk:"customproperty9_label"`
	CustomProperty10Label types.String `tfsdk:"customproperty10_label"`
	CustomProperty11Label types.String `tfsdk:"customproperty11_label"`
	CustomProperty12Label types.String `tfsdk:"customproperty12_label"`
	CustomProperty13Label types.String `tfsdk:"customproperty13_label"`
	CustomProperty14Label types.String `tfsdk:"customproperty14_label"`
	CustomProperty15Label types.String `tfsdk:"customproperty15_label"`
	CustomProperty16Label types.String `tfsdk:"customproperty16_label"`
	CustomProperty17Label types.String `tfsdk:"customproperty17_label"`
	CustomProperty18Label types.String `tfsdk:"customproperty18_label"`
	CustomProperty19Label types.String `tfsdk:"customproperty19_label"`
	CustomProperty20Label types.String `tfsdk:"customproperty20_label"`
	CustomProperty21Label types.String `tfsdk:"customproperty21_label"`
	CustomProperty22Label types.String `tfsdk:"customproperty22_label"`
	CustomProperty23Label types.String `tfsdk:"customproperty23_label"`
	CustomProperty24Label types.String `tfsdk:"customproperty24_label"`
	CustomProperty25Label types.String `tfsdk:"customproperty25_label"`
	CustomProperty26Label types.String `tfsdk:"customproperty26_label"`
	CustomProperty27Label types.String `tfsdk:"customproperty27_label"`
	CustomProperty28Label types.String `tfsdk:"customproperty28_label"`
	CustomProperty29Label types.String `tfsdk:"customproperty29_label"`
	CustomProperty30Label types.String `tfsdk:"customproperty30_label"`
	CustomProperty31Label types.String `tfsdk:"customproperty31_label"`
	CustomProperty32Label types.String `tfsdk:"customproperty32_label"`
	CustomProperty33Label types.String `tfsdk:"customproperty33_label"`
	CustomProperty34Label types.String `tfsdk:"customproperty34_label"`
	CustomProperty35Label types.String `tfsdk:"customproperty35_label"`
	CustomProperty36Label types.String `tfsdk:"customproperty36_label"`
	CustomProperty37Label types.String `tfsdk:"customproperty37_label"`
	CustomProperty38Label types.String `tfsdk:"customproperty38_label"`
	CustomProperty39Label types.String `tfsdk:"customproperty39_label"`
	CustomProperty40Label types.String `tfsdk:"customproperty40_label"`
}

func (d *entitlementTypeDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "saviynt_entitlement_type_datasource"
}

func (d *entitlementTypeDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: util.EntitlementTypeDataSourceDescription,
		Attributes: map[string]schema.Attribute{
			"entitlement_name": schema.StringAttribute{
				MarkdownDescription: "The name of the entitlement type to query.",
				Optional:            true,
			},
			"endpoint_name": schema.StringAttribute{
				MarkdownDescription: "The endpoint name associated with the entitlement type.",
				Optional:            true,
			},
			"authenticate": schema.BoolAttribute{
				Required:            true,
				MarkdownDescription: "If false, do not store sensitive attributes in state",
			},
			"display_count": schema.Int64Attribute{
				MarkdownDescription: "The number of records returned in the response.",
				Computed:            true,
			},
			"total_count": schema.Int64Attribute{
				MarkdownDescription: "The total number of entitlement types available.",
				Computed:            true,
			},
			"error_code": schema.StringAttribute{
				MarkdownDescription: "Error code returned by the API, if any.",
				Computed:            true,
			},
			"msg": schema.StringAttribute{
				MarkdownDescription: "Response message returned by the API.",
				Computed:            true,
			},
			"offset": schema.StringAttribute{
				MarkdownDescription: "Pagination offset for retrieving entitlement types.",
				Optional:            true,
			},
			"max": schema.StringAttribute{
				MarkdownDescription: "Maximum number of entitlement types to retrieve.",
				Optional:            true,
			},
			"results": schema.ListNestedAttribute{
				MarkdownDescription: "List of entitlement types returned by the query.",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: func() map[string]schema.Attribute {
						attrs := map[string]schema.Attribute{
							"order_index":                           schema.StringAttribute{MarkdownDescription: "The order index for the entitlement.", Computed: true},
							"workflow":                              schema.StringAttribute{MarkdownDescription: "The workflow associated with the entitlement.", Computed: true},
							"enable_entitlement_to_role_sync":       schema.BoolAttribute{MarkdownDescription: "Enable entitlement to role sync", Computed: true},
							"required_in_request":                   schema.StringAttribute{MarkdownDescription: "Indicates if the entitlement is required in the request.", Computed: true},
							"required_in_service_request":           schema.StringAttribute{MarkdownDescription: "Indicates if the entitlement is required in service account request.", Computed: true},
							"display_name":                          schema.StringAttribute{MarkdownDescription: "Display name for the entitlement.", Computed: true},
							"security_system":                       schema.StringAttribute{MarkdownDescription: "Associated security system.", Computed: true},
							"available_query_service_account":       schema.StringAttribute{MarkdownDescription: "Query to retrieve available service accounts.", Computed: true},
							"create_task_action":                    schema.StringAttribute{MarkdownDescription: "JSON string representing the list of task actions.", Computed: true},
							"ars_requestable_entitlement_sql_query": schema.StringAttribute{MarkdownDescription: "SQL query for requestable entitlements.", Computed: true},
							"hierarchy_required":                    schema.StringAttribute{MarkdownDescription: "Indicates if hierarchy is required (0 or 1).", Computed: true},
							"entitlement_name":                      schema.StringAttribute{MarkdownDescription: "Internal name for the entitlement.", Computed: true},
							"entitlement_type_key":                  schema.Int64Attribute{MarkdownDescription: "Unique key of the entitlement type.", Computed: true},
							"endpoint_name":                         schema.StringAttribute{MarkdownDescription: "Endpoint name associated with the entitlement.", Computed: true},
							"endpoint_key":                          schema.Int64Attribute{MarkdownDescription: "Unique key of the associated endpoint.", Computed: true},
							"request_option":                        schema.StringAttribute{MarkdownDescription: "Request option UI behavior.", Computed: true},
							"enable_provisioning_priority":          schema.StringAttribute{MarkdownDescription: "Enable provisioning priority.", Computed: true},
							"exclude_rule_assigned_ents_in_request": schema.StringAttribute{MarkdownDescription: "Exclude Entitlements Assigned via Rule while Request.", Computed: true},
							"show_on_child":                         schema.StringAttribute{MarkdownDescription: "Indicates if entitlement should show on child accounts.", Computed: true},
							"show_ent_type_on":                      schema.StringAttribute{MarkdownDescription: "Show entitlement type on.", Computed: true},
							"recon":                                 schema.StringAttribute{MarkdownDescription: "Indicates if the entitlement is part of reconciliation.", Computed: true},
							"ars_selected_entitlement_sql_query":    schema.StringAttribute{MarkdownDescription: "SQL query for selected entitlements.", Computed: true},
							"certifiable":                           schema.StringAttribute{MarkdownDescription: "Indicates if the entitlement is certifiable.", Computed: true},
							"entitlement_description":               schema.StringAttribute{MarkdownDescription: "Description of the entitlement.", Computed: true},
							"request_dates_conf_json":               schema.StringAttribute{MarkdownDescription: "JSON configuration for request dates.", Computed: true},
							"selected_query_service_account":        schema.StringAttribute{MarkdownDescription: "Query to retrieve selected service accounts.", Computed: true},
						}

						for i := 1; i <= 40; i++ {
							label := fmt.Sprintf("customproperty%d_label", i)
							attrs[label] = schema.StringAttribute{
								MarkdownDescription: fmt.Sprintf("Label for custom property %d.", i),
								Computed:            true,
							}
						}

						return attrs
					}(),
				},
			},
		},
	}
}

func (d *entitlementTypeDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	tflog.Debug(ctx, "Starting entitlement type datasource configuration")

	// Check if provider data is available.
	if req.ProviderData == nil {
		tflog.Debug(ctx, "ProviderData is nil, returning early")
		return
	}

	// Cast provider data to your provider type.
	prov, ok := req.ProviderData.(*saviyntProvider)
	if !ok {
		tflog.Error(ctx, "Provider configuration failed", map[string]interface{}{
			"expected_type": "*saviyntProvider",
		})
		resp.Diagnostics.AddError(
			"Unexpected Provider Data",
			"Expected *saviyntProvider, got different type",
		)
		return
	}

	// Set the client and token from the provider state using interface wrapper.
	d.client = &client.SaviyntClientWrapper{Client: prov.client}
	d.token = prov.accessToken

	tflog.Debug(ctx, "Entitlement type datasource configured successfully")
}

func (d *entitlementTypeDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var state EntitlementTypeDataSourceModel

	tflog.Debug(ctx, "Starting entitlement type datasource read operation")

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

	// Execute API call to get entitlement type details
	apiResp, err := d.ReadEntitlementTypeDetails(ctx, &state)
	if err != nil {
		tflog.Error(ctx, "Failed to read entitlement type details", map[string]interface{}{
			"error": err.Error(),
		})
		resp.Diagnostics.AddError("API Call Failed", fmt.Sprintf("Error: %v", err))
		return
	}

	// Check if no entitlements were found and add user-visible warning
	if apiResp != nil && apiResp.EntitlementTypeDetails!=nil && len(apiResp.EntitlementTypeDetails) == 0 {
		resp.Diagnostics.AddWarning(
			"No Entitlements Found",
			fmt.Sprintf("No entitlement types found for entitlement_name='%s' and endpoint_name='%s'. Retrieved count: %d",
				state.EntitlementName.ValueString(),
				state.EndpointName.ValueString(),
				*apiResp.DisplayCount),
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
			"Unable to update Terraform state for entitlement type datasource",
		)
		return
	}

	tflog.Debug(ctx, "Entitlement type datasource read operation completed successfully")
}

// ReadEntitlementTypeDetails retrieves entitlement type details from Saviynt API
// Handles parameter preparation and API call execution using factory pattern
func (d *entitlementTypeDataSource) ReadEntitlementTypeDetails(ctx context.Context, state *EntitlementTypeDataSourceModel) (*openapi.GetEntitlementTypeResponse, error) {
	tflog.Debug(ctx, "Starting entitlement type API call")

	// Create entitlement type operations using the factory
	entitlementTypeOps := d.entitlementTypeFactory.CreateEntitlementTypeOperations(d.client.APIBaseURL(), d.token)

	// Prepare parameters using existing utility function
	entitlementName := util.SafeStringValue(state.EntitlementName)
	endpointName := util.SafeStringValue(state.EndpointName)
	max := util.SafeStringValue(state.Max)
	offset := util.SafeStringValue(state.Offset)

	tflog.Debug(ctx, "Executing API request to get entitlement type details", map[string]interface{}{
		"entitlement_name": entitlementName,
		"endpoint_name":    endpointName,
		"max":              max,
		"offset":           offset,
	})

	// Execute API call
	readResp, httpResp, err := entitlementTypeOps.GetEntitlementType(ctx, entitlementName, max, offset, endpointName)
	if err != nil {
		return nil, fmt.Errorf("Entitlement Type Datasource: API call failed: %w", err)
	}

	if readResp!=nil && readResp.ErrorCode!=nil && *readResp.ErrorCode!="0"{
		return nil, fmt.Errorf("Entitlement Type Datasource: API returned error code: %s and error message: %s", *readResp.ErrorCode, *readResp.Msg)
	}

	tflog.Debug(ctx, "Entitltement type: API call successful", map[string]interface{}{
		"status_code": httpResp.StatusCode,
	})

	return readResp, nil
}

// UpdateModelFromAPIResponse maps API response data to the Terraform state model
func (d *entitlementTypeDataSource) UpdateModelFromAPIResponse(state *EntitlementTypeDataSourceModel, apiResp *openapi.GetEntitlementTypeResponse) {
	// Map basic response fields
	d.MapBasicResponseFields(state, apiResp)

	// Process entitlement type details if available
	if apiResp.EntitlementTypeDetails != nil {
		for _, item := range apiResp.EntitlementTypeDetails {
			entitlementType := d.MapEntitlementTypeDetails(&item)
			state.Results = append(state.Results, entitlementType)
		}
	}
}

// MapBasicResponseFields maps basic response fields from API response to state model
func (d *entitlementTypeDataSource) MapBasicResponseFields(state *EntitlementTypeDataSourceModel, apiResp *openapi.GetEntitlementTypeResponse) {
	state.Msg = util.SafeString(apiResp.Msg)
	state.DisplayCount = util.SafeInt64(apiResp.DisplayCount)
	state.ErrorCode = util.SafeString(apiResp.ErrorCode)
	state.TotalCount = util.SafeInt64(apiResp.TotalCount)
}

// HandleAuthenticationLogic processes the authenticate flag to control sensitive data visibility
// When authenticate=false, results are removed from state to prevent sensitive data exposure
// When authenticate=true, all entitlement type results are returned in state
func (d *entitlementTypeDataSource) HandleAuthenticationLogic(state *EntitlementTypeDataSourceModel, resp *datasource.ReadResponse) {
	if !state.Authenticate.IsNull() && !state.Authenticate.IsUnknown() {
		if state.Authenticate.ValueBool() {
			tflog.Info(context.Background(), "Authentication enabled - returning all entitlement type details")
			resp.Diagnostics.AddWarning(
				"Authentication Enabled",
				"`authenticate` is true; all entitlement type details will be returned in state.",
			)
		} else {
			tflog.Info(context.Background(), "Authentication disabled - removing entitlement type details from state")
			resp.Diagnostics.AddWarning(
				"Authentication Disabled",
				"`authenticate` is false; entitlement type details will be removed from state.",
			)
			state.Results = nil
		}
	}
}

// MapEntitlementTypeDetails maps individual entitlement type details from API response to state model
func (d *entitlementTypeDataSource) MapEntitlementTypeDetails(item *openapi.GetEntitlementTypeResponseEntitlementTypeDetailsInner) EntitlementType {
	// Parse workflow JSON to extract workflow name and enable_entitlement_to_role_sync (same logic as resource)
	var workflow *string
	var enableEntitlementToRoleSync bool = false // Default to false if not provided

	if item.Workflow != nil && *item.Workflow != "" {
		var parsed map[string]json.RawMessage
		err := json.Unmarshal([]byte(*item.Workflow), &parsed)
		if err == nil {
			// Extract workflow name
			if rawWorkflow, ok := parsed["workflow"]; ok {
				var val string
				if err := json.Unmarshal(rawWorkflow, &val); err == nil {
					workflow = &val
				}
			}
			// Extract enable_entitlement_to_role_sync
			if rawEnable, ok := parsed["enableEntitlementToRoleSync"]; ok {
				var val bool
				if err := json.Unmarshal(rawEnable, &val); err == nil {
					enableEntitlementToRoleSync = val
				}
			}
		} else {
			// Older format: simple string value
			workflow = item.Workflow
		}
	}

	entitlementType := EntitlementType{
		OrderIndex:                   util.SafeString(item.Orderindex),
		Workflow:                     util.SafeString(workflow),
		EnableEntitlementToRoleSync:  types.BoolValue(enableEntitlementToRoleSync),
		RequiredInRequest:            util.SafeString(item.Requiredinrequest),
		RequiredInServiceRequest:     util.SafeString(item.Requiredinservicerequest),
		DisplayName:                  util.SafeString(item.DisplayName),
		SecuritySystem:               util.SafeString(item.Securitysystem),
		AvailableQueryServiceAccount: util.SafeString(item.AvailableQueryServiceAccount),
		CreateTaskAction:             util.SafeString(item.CreateTaskAction),
		ArsRequestableEntitlementSQL: util.SafeString(item.ArsReqEntSqlquerey),
		HierarchyRequired:            util.SafeString(item.Hiearchyrequired),
		EntitlementName:              util.SafeString(item.Entitlementname),
		EntitlementTypeKey:           util.SafeInt64(item.EntitlementTypeKey),
		EndpointName:                 util.SafeString(item.Endpoint),
		EndpointKey:                  util.SafeInt64(item.EndpointKey),
		RequestOption:                util.SafeString(item.Requestoption),
		EnableProvisioningPriority:   util.SafeString(item.EnableProvisioningPriority),
		ExcludeRuleAssignedEnts:      util.SafeString(item.ExcludeRuleAssgnEntsInRequest),
		ShowOnChild:                  util.SafeString(item.Showonchild),
		ShowEntTypeOn:                util.SafeString(item.ShowEntTypeOn),
		Recon:                        util.SafeString(item.Recon),
		ArsSelectedEntitlementSQL:    util.SafeString(item.ArsSelectEntSqlquerey),
		Certifiable:                  util.SafeString(item.Certifiable),
		EntitlementDescription:       util.SafeString(item.Entitlementdescription),
		RequestDatesConfJson:         util.SafeString(item.RequestDatesConfJson),
		SelectedQueryServiceAccount:  util.SafeString(item.SelectedQueryServiceAccount),
	}

	// Map custom property labels
	d.MapCustomPropertyLabels(&entitlementType, item)

	return entitlementType
}

// MapCustomPropertyLabels maps all custom property labels from API response to entitlement type
func (d *entitlementTypeDataSource) MapCustomPropertyLabels(entitlementType *EntitlementType, item *openapi.GetEntitlementTypeResponseEntitlementTypeDetailsInner) {
	entitlementType.CustomProperty1Label = util.SafeString(item.Customproperty1Label)
	entitlementType.CustomProperty2Label = util.SafeString(item.Customproperty2Label)
	entitlementType.CustomProperty3Label = util.SafeString(item.Customproperty3Label)
	entitlementType.CustomProperty4Label = util.SafeString(item.Customproperty4Label)
	entitlementType.CustomProperty5Label = util.SafeString(item.Customproperty5Label)
	entitlementType.CustomProperty6Label = util.SafeString(item.Customproperty6Label)
	entitlementType.CustomProperty7Label = util.SafeString(item.Customproperty7Label)
	entitlementType.CustomProperty8Label = util.SafeString(item.Customproperty8Label)
	entitlementType.CustomProperty9Label = util.SafeString(item.Customproperty9Label)
	entitlementType.CustomProperty10Label = util.SafeString(item.Customproperty10Label)
	entitlementType.CustomProperty11Label = util.SafeString(item.Customproperty11Label)
	entitlementType.CustomProperty12Label = util.SafeString(item.Customproperty12Label)
	entitlementType.CustomProperty13Label = util.SafeString(item.Customproperty13Label)
	entitlementType.CustomProperty14Label = util.SafeString(item.Customproperty14Label)
	entitlementType.CustomProperty15Label = util.SafeString(item.Customproperty15Label)
	entitlementType.CustomProperty16Label = util.SafeString(item.Customproperty16Label)
	entitlementType.CustomProperty17Label = util.SafeString(item.Customproperty17Label)
	entitlementType.CustomProperty18Label = util.SafeString(item.Customproperty18Label)
	entitlementType.CustomProperty19Label = util.SafeString(item.Customproperty19Label)
	entitlementType.CustomProperty20Label = util.SafeString(item.Customproperty20Label)
	entitlementType.CustomProperty21Label = util.SafeString(item.Customproperty21Label)
	entitlementType.CustomProperty22Label = util.SafeString(item.Customproperty22Label)
	entitlementType.CustomProperty23Label = util.SafeString(item.Customproperty23Label)
	entitlementType.CustomProperty24Label = util.SafeString(item.Customproperty24Label)
	entitlementType.CustomProperty25Label = util.SafeString(item.Customproperty25Label)
	entitlementType.CustomProperty26Label = util.SafeString(item.Customproperty26Label)
	entitlementType.CustomProperty27Label = util.SafeString(item.Customproperty27Label)
	entitlementType.CustomProperty28Label = util.SafeString(item.Customproperty28Label)
	entitlementType.CustomProperty29Label = util.SafeString(item.Customproperty29Label)
	entitlementType.CustomProperty30Label = util.SafeString(item.Customproperty30Label)
	entitlementType.CustomProperty31Label = util.SafeString(item.Customproperty31Label)
	entitlementType.CustomProperty32Label = util.SafeString(item.Customproperty32Label)
	entitlementType.CustomProperty33Label = util.SafeString(item.Customproperty33Label)
	entitlementType.CustomProperty34Label = util.SafeString(item.Customproperty34Label)
	entitlementType.CustomProperty35Label = util.SafeString(item.Customproperty35Label)
	entitlementType.CustomProperty36Label = util.SafeString(item.Customproperty36Label)
	entitlementType.CustomProperty37Label = util.SafeString(item.Customproperty37Label)
	entitlementType.CustomProperty38Label = util.SafeString(item.Customproperty38Label)
	entitlementType.CustomProperty39Label = util.SafeString(item.Customproperty39Label)
	entitlementType.CustomProperty40Label = util.SafeString(item.Customproperty40Label)
}
