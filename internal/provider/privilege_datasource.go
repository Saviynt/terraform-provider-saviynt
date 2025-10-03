// Copyright (c) 2025 Saviynt Inc.
// SPDX-License-Identifier: MPL-2.0

// saviynt_privilege_datasource retrieves privilege details from the Saviynt Security Manager.
// The data source supports a single Read operation to look up an existing privilege with various filters like endpoint, entitlement type etc.

package provider

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"terraform-provider-Saviynt/internal/client"
	"terraform-provider-Saviynt/util"
	"terraform-provider-Saviynt/util/errorsutil"
	"terraform-provider-Saviynt/util/privilegeutil"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	openapi "github.com/saviynt/saviynt-api-go-client/privileges"
)

type privilegeDatasource struct {
	client           client.SaviyntClientInterface
	token            string
	provider         client.SaviyntProviderInterface
	privilegeFactory client.PrivilegeFactoryInterface
}

var _ datasource.DataSource = &privilegeDatasource{}
var _ datasource.DataSourceWithConfigure = &privilegeDatasource{}

func NewPrivilegeDataSource() datasource.DataSource {
	return &privilegeDatasource{
		privilegeFactory: &client.DefaultPrivilegeFactory{},
	}
}

// NewPrivilegeDataSourceWithFactory creates a new privilege data source with custom factory
// Used primarily for testing with mock factories
func NewPrivilegeDataSourceWithFactory(factory client.PrivilegeFactoryInterface) datasource.DataSource {
	return &privilegeDatasource{
		privilegeFactory: factory,
	}
}

// SetClient sets the client for testing purposes
func (d *privilegeDatasource) SetClient(client client.SaviyntClientInterface) {
	d.client = client
}

// SetToken sets the token for testing purposes
func (d *privilegeDatasource) SetToken(token string) {
	d.token = token
}

// SetProvider sets the provider for testing purposes
func (r *privilegeDatasource) SetProvider(provider client.SaviyntProviderInterface) {
	r.provider = provider
}

type PrivilegeDataSourceModel struct {
	// Input Filters
	Endpoint        types.String `tfsdk:"endpoint"`
	Entitlementtype types.String `tfsdk:"entitlement_type"`
	Offset          types.String `tfsdk:"offset"`
	Max             types.String `tfsdk:"max"`

	// Output
	Msg          types.String        `tfsdk:"msg"`
	ErrorCode    types.String        `tfsdk:"error_code"`
	DisplayCount types.Int32         `tfsdk:"display_count"`
	TotalCount   types.Int32         `tfsdk:"total_count"`
	Authenticate types.Bool          `tfsdk:"authenticate"`
	Privileges   []PrivilegeDataItem `tfsdk:"privileges_list"`
}

// PrivilegeDataItem represents a privilege item in the datasource (separate from resource Privilege struct)
type PrivilegeDataItem struct {
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

func (d *privilegeDatasource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "saviynt_privilege_datasource"
}

func (d *privilegeDatasource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: util.PrivilegeDataSourceDescription,
		Attributes: map[string]schema.Attribute{
			"endpoint": schema.StringAttribute{
				MarkdownDescription: "The name of the endpoint to query.",
				Required:            true,
			},
			"entitlement_type": schema.StringAttribute{
				MarkdownDescription: "The name of the entitlement type to query.",
				Optional:            true,
			},
			"authenticate": schema.BoolAttribute{
				Required:            true,
				MarkdownDescription: "If false, do not store sensitive attributes in state",
			},
			"display_count": schema.Int32Attribute{
				MarkdownDescription: "The number of records returned in the response.",
				Computed:            true,
			},
			"total_count": schema.Int32Attribute{
				MarkdownDescription: "The total number of privileges available.",
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
				MarkdownDescription: "Pagination offset for retrieving privileges.",
				Optional:            true,
			},
			"max": schema.StringAttribute{
				MarkdownDescription: "Maximum number of privileges to retrieve.",
				Optional:            true,
			},
			"privileges_list": schema.ListNestedAttribute{
				Computed: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"attribute_name": schema.StringAttribute{
							Computed:    true,
							Description: "Attribute name for the privilege",
						},
						"attribute_type": schema.StringAttribute{
							Computed:    true,
							Description: "Type of the attribute/privilege",
						},
						"order_index": schema.StringAttribute{
							Computed:    true,
							Description: "Order index",
						},
						"default_value": schema.StringAttribute{
							Computed:    true,
							Description: "Default value for the privilege",
						},
						"attribute_config": schema.StringAttribute{
							Computed:    true,
							Description: "Configuration type for the attribute",
						},
						"label": schema.StringAttribute{
							Computed:    true,
							Description: "Label for the privilege",
						},
						"attribute_group": schema.StringAttribute{
							Computed:    true,
							Description: "Attribute group",
						},
						"parent_attribute": schema.StringAttribute{
							Computed:    true,
							Description: "Parent attribute for the given privilege",
						},
						"child_action": schema.StringAttribute{
							Computed:    true,
							Description: "Child action",
						},
						"description": schema.StringAttribute{
							Computed:    true,
							Description: "Description for the privilege",
						},
						"required": schema.BoolAttribute{
							Computed:    true,
							Description: "Is required",
						},
						"requestable": schema.BoolAttribute{
							Computed:    true,
							Description: "Is requestable",
						},
						"hide_on_create": schema.BoolAttribute{
							Optional:    true,
							Computed:    true,
							Description: "Hide on create",
						},
						"hide_on_update": schema.BoolAttribute{
							Computed:    true,
							Description: "Hide on update",
						},
						"action_string": schema.StringAttribute{
							Computed:    true,
							Description: "Action string",
						},
					},
				},
			},
		},
	}
}

func (d *privilegeDatasource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	tflog.Debug(ctx, "Starting privilege datasource configuration")

	// Check if provider data is available.
	if req.ProviderData == nil {
		tflog.Debug(ctx, "ProviderData is nil, returning early")
		return
	}

	// Cast provider data to your provider type.
	prov, ok := req.ProviderData.(*SaviyntProvider)
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
	d.provider = &client.SaviyntProviderWrapper{Provider: prov} // Store provider reference for retry logic
	tflog.Debug(ctx, "Privilege datasource configured successfully")
}

func (d *privilegeDatasource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var state PrivilegeDataSourceModel

	tflog.Debug(ctx, "Starting privilege datasource read operation")

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

	// Execute API call to get privilege details
	apiResp, err := d.ReadPrivilegeDetails(ctx, &state)
	if err != nil {
		tflog.Error(ctx, "Failed to read privilegs details", map[string]interface{}{
			"error": err.Error(),
		})
		resp.Diagnostics.AddError("API Call Failed", fmt.Sprintf("Error: %v", err))
		return
	}

	// Check if no privileges were found and add user-visible warning
	if apiResp != nil && apiResp.PrivilegeDetails != nil && len(apiResp.PrivilegeDetails) == 0 {
		resp.Diagnostics.AddWarning(
			"No Privileges Found",
			fmt.Sprintf("No privileges found for endpoint_name='%s' and the other filters.", state.Endpoint.ValueString()),
		)
	}

	// Map API response to state
	d.UpdatePrivilegeModelFromAPIResponse(&state, apiResp)

	// Handle authentication logic for results
	d.HandleAuthenticationLogic(&state, resp)

	// Set final state
	stateDiagnostics := resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(stateDiagnostics...)
	if resp.Diagnostics.HasError() {
		tflog.Error(ctx, "Failed to set state")
		resp.Diagnostics.AddError(
			"State Update Error",
			"Unable to update Terraform state for privilege datasource",
		)
		return
	}

	tflog.Debug(ctx, "Privilege datasource read operation completed successfully")
}

// ReadPrivilegeDetails retrieves privilege details from Saviynt API
// Handles parameter preparation and API call execution with refresh token retry logic
func (d *privilegeDatasource) ReadPrivilegeDetails(ctx context.Context, state *PrivilegeDataSourceModel) (*openapi.GetPrivilegeListResponse, error) {
	tflog.Debug(ctx, "Starting Privilege read API call in ReadPrivilegeDetails")

	var readResp *openapi.GetPrivilegeListResponse
	var finalHttpResp *http.Response
	// Execute API request with retry logic
	err := d.provider.AuthenticatedAPICallWithRetry(ctx, "read_privilege_datasource", func(token string) error {
		privilegeOps := d.privilegeFactory.CreatePrivilegeOperations(d.client.APIBaseURL(), token)

		getReq := openapi.GetPrivilegeListRequest{
			Endpoint:        util.SafeStringValue(state.Endpoint),
			Entitlementtype: util.StringPointerOrEmpty(state.Entitlementtype),
			Max:             util.StringPointerOrEmpty(state.Max),
			Offset:          util.StringPointerOrEmpty(state.Offset),
		}

		getReqJson, _ := json.Marshal(getReq)
		tflog.Debug(ctx, "Privilege Datasource: Get API REQUEST", map[string]interface{}{
			"request": string(getReqJson),
		})

		tflog.Debug(ctx, "Executing API request to get privilege details", map[string]interface{}{
			"endpoint":         state.Endpoint,
			"entitlement_type": state.Entitlementtype,
			"max":              state.Max,
			"offset":           state.Offset,
		})

		resp, httpResp, err := privilegeOps.GetPrivilege(ctx, getReq)
		if httpResp != nil && httpResp.StatusCode == 401 {
			return fmt.Errorf("401 unauthorized")
		}
		readResp = resp
		finalHttpResp = httpResp
		return err
	})

	if err != nil {
		log.Printf("[ERROR] Privileges: Error in reading privilege: %v", err)
		err = errorsutil.HandleHTTPError(finalHttpResp, err, "ReadPrivilegeDetails")
		return nil, fmt.Errorf("Privilege Datasource: API call failed: %w", err)
	}

	getRespJson, _ := json.Marshal(readResp)
	tflog.Debug(ctx, "Privilege Datasource: Get API Response", map[string]interface{}{
		"request": string(getRespJson),
	})

	if readResp != nil && readResp.ErrorCode != nil && *readResp.ErrorCode != "0" {
		return nil, fmt.Errorf("Privilege Datasource: API returned error code: %s and error message: %s", *readResp.ErrorCode, *readResp.Msg)
	}

	tflog.Debug(ctx, "Privilege Datasource: API call successful", map[string]interface{}{
		"status_code": finalHttpResp.StatusCode,
	})

	return readResp, nil
}

// UpdatePrivilegeModelFromAPIResponse maps API response data to the Terraform state model
func (d *privilegeDatasource) UpdatePrivilegeModelFromAPIResponse(state *PrivilegeDataSourceModel, apiResp *openapi.GetPrivilegeListResponse) {
	// Map basic response fields
	d.MapBasicPrivilegeResponseFields(state, apiResp)

	// Process privilege details if available
	if apiResp.PrivilegeDetails != nil {
		for _, item := range apiResp.PrivilegeDetails {
			privilege := d.MapPrivilegeDetails(&item)
			state.Privileges = append(state.Privileges, privilege)
		}
	}
}

// MapBasicPrivilegeResponseFields maps basic response fields from API response to state model
func (d *privilegeDatasource) MapBasicPrivilegeResponseFields(state *PrivilegeDataSourceModel, apiResp *openapi.GetPrivilegeListResponse) {
	state.Msg = util.SafeString(apiResp.Msg)
	state.DisplayCount = util.SafeInt32(apiResp.DisplayCount)
	state.ErrorCode = util.SafeString(apiResp.ErrorCode)
	state.TotalCount = util.SafeInt32(apiResp.TotalCount)
}

// HandleAuthenticationLogic processes the authenticate flag to control sensitive data visibility
// When authenticate=false, results are removed from state to prevent sensitive data exposure
// When authenticate=true, all privilege results are returned in state
func (d *privilegeDatasource) HandleAuthenticationLogic(state *PrivilegeDataSourceModel, resp *datasource.ReadResponse) {
	if !state.Authenticate.IsNull() && !state.Authenticate.IsUnknown() {
		if state.Authenticate.ValueBool() {
			tflog.Info(context.Background(), "Authentication enabled - returning all privilege details")
			resp.Diagnostics.AddWarning(
				"Authentication Enabled",
				"`authenticate` is true; all privilege details will be returned in state.",
			)
		} else {
			tflog.Info(context.Background(), "Authentication disabled - removing privilege details from state")
			resp.Diagnostics.AddWarning(
				"Authentication Disabled",
				"`authenticate` is false; privilege details will be removed from state.",
			)
			state.Privileges = nil
		}
	}
}

// MapPrivilegeDetails maps individual privilege details from API response to state model
func (d *privilegeDatasource) MapPrivilegeDetails(item *openapi.GetPrivilegeDetail) PrivilegeDataItem {
	// Normalize attribute type from API response format to Terraform format
	normalizedAttributeType := privilegeutil.TranslateValue(*item.AttributeType, privilegeutil.AttributeTypeMap)

	privilege := PrivilegeDataItem{
		AttributeName:   types.StringValue(*item.Attribute),
		AttributeType:   types.StringValue(normalizedAttributeType),
		OrderIndex:      util.SafeString(item.Orderindex),
		DefaultValue:    util.SafeString(item.Defaultvalue),
		AttributeConfig: util.SafeString(item.AttributeConfig),
		Label:           util.SafeString(item.Label),
		AttributeGroup:  util.SafeString(item.Attributegroup),
		ParentAttribute: util.SafeString(item.Parentattribute),
		ChildAction:     util.SafeString(item.Childaction),
		Description:     util.SafeString(item.Descriptionascsv),
		Required:        util.SafeBoolDatasource(item.Required),
		Requestable:     util.SafeBoolDatasource(item.Requestablerequired),
		HideOnCreate:    util.SafeBoolDatasource(item.Hideoncreate),
		HideOnUpdate:    util.SafeBoolDatasource(item.Hideonupd),
		ActionString:    util.SafeString(item.ActionString),
	}

	return privilege
}
