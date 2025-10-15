// Copyright (c) 2025 Saviynt Inc.
// SPDX-License-Identifier: MPL-2.0

// saviynt_dynamic_attribute_datasource retrieves dynamic attribute details from the Saviynt Security Manager.
// The data source supports a single Read operation to look up existing dynamic attributes with various filters 
// like attribute name, endpoint, security system, etc.
package provider

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"terraform-provider-Saviynt/internal/client"
	"terraform-provider-Saviynt/util"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	openapi "github.com/saviynt/saviynt-api-go-client/dynamicattributes"
)

var _ datasource.DataSource = &DynamicAttributeDataSource{}

type DynamicAttributeDataSource struct {
	client                  client.SaviyntClientInterface
	token                   string
	provider                client.SaviyntProviderInterface
	dynamicAttributeFactory client.DynamicAttributeFactoryInterface
}

type DynamicAttributeDataSourceModel struct {
	// Input Filters
	SecuritySystem    types.List   `tfsdk:"securitysystem"`
	Endpoint          types.List   `tfsdk:"endpoint"`
	DynamicAttributes types.List   `tfsdk:"dynamic_attributes"`
	RequestType       types.List   `tfsdk:"requesttype"`
	Offset            types.String `tfsdk:"offset"`
	Max               types.String `tfsdk:"max"`
	LoggedInUser      types.String `tfsdk:"loggedinuser"`

	// Output
	Msg               types.String        `tfsdk:"msg"`
	ErrorCode         types.String        `tfsdk:"error_code"`
	DisplayCount      types.Int32         `tfsdk:"display_count"`
	TotalCount        types.Int32         `tfsdk:"total_count"`
	Authenticate      types.Bool          `tfsdk:"authenticate"`
	Dynamicattributes []DynamicAttributes `tfsdk:"dynamic_attributes_list"`
}

type DynamicAttributes struct {
	Attributename                                   types.String `tfsdk:"attribute_name"`
	Requesttype                                     types.String `tfsdk:"request_type"`
	Securitysystem                                  types.String `tfsdk:"security_system"`
	Endpoint                                        types.String `tfsdk:"endpoint"`
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
	Defaultvalue                                    types.String `tfsdk:"default_value"`
	Required                                        types.String `tfsdk:"required"`
	Attributevalue                                  types.String `tfsdk:"attribute_value"`
	Showonchild                                     types.String `tfsdk:"showonchild"`
	Parentattribute                                 types.String `tfsdk:"parentattribute"`
	Descriptionascsv                                types.String `tfsdk:"descriptionascsv"`

	// Not supported by Read API
	// Regex                                     types.String `tfsdk:"regex"`
}

// NewDynamicAttributeDataSource creates a new dynamic attribute data source with default factory
func NewDynamicAttributeDataSource() datasource.DataSource {
	return &DynamicAttributeDataSource{
		dynamicAttributeFactory: &client.DefaultDynamicAttributeFactory{},
	}
}

// NewDynamicAttributeDataSourceWithFactory creates a new dynamic attribute data source with custom factory
// Used primarily for testing with mock factories
func NewDynamicAttributeDataSourceWithFactory(factory client.DynamicAttributeFactoryInterface) datasource.DataSource {
	return &DynamicAttributeDataSource{
		dynamicAttributeFactory: factory,
	}
}

// SetClient sets the client for testing purposes
func (d *DynamicAttributeDataSource) SetClient(client client.SaviyntClientInterface) {
	d.client = client
}

// SetToken sets the token for testing purposes
func (d *DynamicAttributeDataSource) SetToken(token string) {
	d.token = token
}

// SetProvider sets the provider for testing purposes
func (d *DynamicAttributeDataSource) SetProvider(provider client.SaviyntProviderInterface) {
	d.provider = provider
}

func (d *DynamicAttributeDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "saviynt_dynamic_attribute_datasource"
}

func (d *DynamicAttributeDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Retrieve the list of Dynamic Attributes",
		Attributes: map[string]schema.Attribute{
			"msg": schema.StringAttribute{
				Computed: true,
			},
			"error_code": schema.StringAttribute{
				Computed: true,
			},
			"securitysystem": schema.ListAttribute{
				Optional:    true,
				Description: "Security System",
				Computed:    false,
				ElementType: types.StringType,
			},
			"endpoint": schema.ListAttribute{
				Optional:    true,
				Description: "Endpoint",
				Computed:    false,
				ElementType: types.StringType,
			},
			"dynamic_attributes": schema.ListAttribute{
				Optional:    true,
				Description: "Dynamic Attributes",
				Computed:    false,
				ElementType: types.StringType,
			},
			"requesttype": schema.ListAttribute{
				Optional:    true,
				Description: "Request Type",
				Computed:    false,
				ElementType: types.StringType,
			},
			"offset": schema.StringAttribute{
				Optional:    true,
				Description: "Offset",
				Computed:    false,
			},
			"max": schema.StringAttribute{
				Optional:    true,
				Description: "Max",
				Computed:    false,
			},
			"loggedinuser": schema.StringAttribute{
				Optional:    true,
				Description: "Logged In User",
				Computed:    false,
			},
			"authenticate": schema.BoolAttribute{
				Required:    true,
				Description: "If false, do not store connection_attributes in state",
			},
			"dynamic_attributes_list": schema.ListNestedAttribute{
				Computed: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"attribute_name":  schema.StringAttribute{Computed: true},
						"request_type":    schema.StringAttribute{Computed: true},
						"security_system": schema.StringAttribute{Computed: true},
						"endpoint":        schema.StringAttribute{Computed: true},
						"attribute_type":  schema.StringAttribute{Computed: true},
						"attribute_group": schema.StringAttribute{Computed: true},
						"order_index":     schema.StringAttribute{Computed: true},
						"attribute_lable": schema.StringAttribute{Computed: true},
						"accounts_column": schema.StringAttribute{Computed: true},
						"hide_on_create":  schema.StringAttribute{Computed: true},
						"action_string":   schema.StringAttribute{Computed: true},
						"editable":        schema.StringAttribute{Computed: true},
						"hide_on_update":  schema.StringAttribute{Computed: true},
						"action_to_perform_when_parent_attribute_changes": schema.StringAttribute{Computed: true},
						"default_value": schema.StringAttribute{Computed: true},
						"required":      schema.StringAttribute{Computed: true},
						// "regex":                                     schema.StringAttribute{Computed: true},
						"attribute_value":  schema.StringAttribute{Computed: true},
						"showonchild":      schema.StringAttribute{Computed: true},
						"parentattribute":  schema.StringAttribute{Computed: true},
						"descriptionascsv": schema.StringAttribute{Computed: true},
					},
				},
			},
			"display_count": schema.Int32Attribute{
				Computed: true,
			},
			"total_count": schema.Int32Attribute{
				Computed: true,
			},
		},
	}
}

// Configure initializes the data source with provider configuration
// Sets up client and authentication token for API operations
func (d *DynamicAttributeDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	// Check if provider data is available.
	if req.ProviderData == nil {
		log.Println("[DEBUG] DynamicAttribute: ProviderData is nil, returning early.")
		return
	}

	// Cast provider data to your provider type.
	prov, ok := req.ProviderData.(*SaviyntProvider)
	if !ok {
		log.Printf("[ERROR] DynamicAttribute: Unexpected Provider Data")
		resp.Diagnostics.AddError("Unexpected Provider Data", "Expected *SaviyntProvider")
		return
	}

	// Set the client and token from the provider state using interface wrapper.
	d.client = &client.SaviyntClientWrapper{Client: prov.client}
	d.token = prov.accessToken
	d.provider = &client.SaviyntProviderWrapper{Provider: prov}

	log.Printf("[DEBUG] DynamicAttribute: Datasource configured successfully.")
}

// Read retrieves dynamic attribute details from Saviynt and populates the Terraform state
// Supports lookup with various filters and comprehensive error handling
func (d *DynamicAttributeDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var state DynamicAttributeDataSourceModel

	log.Printf("[DEBUG] DynamicAttribute: Starting datasource read")

	// Extract configuration from request
	configDiagnostics := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(configDiagnostics...)
	if resp.Diagnostics.HasError() {
		log.Printf("[ERROR] DynamicAttribute: Failed to get config from request")
		resp.Diagnostics.AddError("Configuration Error", "Unable to extract Terraform configuration from request")
		return
	}

	// Prepare request parameters
	securitySystems := util.ConvertListToStringSlice(ctx, state.SecuritySystem)
	endpoints := util.ConvertListToStringSlice(ctx, state.Endpoint)
	dynamicAttributes := util.ConvertListToStringSlice(ctx, state.DynamicAttributes)
	requestTypes := util.ConvertListToStringSlice(ctx, state.RequestType)
	loggedInUser := util.StringPointerOrEmpty(state.LoggedInUser)
	offset := util.StringPointerOrEmpty(state.Offset)
	max := util.StringPointerOrEmpty(state.Max)

	// Execute API call to get dynamic attribute details
	apiResp, err := d.ReadDynamicAttributeDetails(ctx, securitySystems, endpoints, dynamicAttributes, requestTypes, loggedInUser, offset, max)
	if err != nil {
		log.Printf("[ERROR] DynamicAttribute: Failed to read dynamic attribute details: %v", err)
		resp.Diagnostics.AddError("API Read Failed", fmt.Sprintf("Error: %v", err))
		return
	}

	// Map API response to state
	d.UpdateModelFromDynamicAttributeResponse(&state, apiResp)

	// Handle authentication logic
	d.HandleDynamicAttributeAuthenticationLogic(&state, resp)

	// Set final state
	stateDiagnostics := resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(stateDiagnostics...)
	if resp.Diagnostics.HasError() {
		log.Printf("[ERROR] DynamicAttribute: Failed to set state")
		resp.Diagnostics.AddError("State Update Error", "Unable to update Terraform state for dynamic attribute datasource")
		return
	}

	log.Printf("[DEBUG] DynamicAttribute: Datasource read completed successfully with %d results", len(state.Dynamicattributes))
}
// ReadDynamicAttributeDetails retrieves dynamic attribute details from Saviynt API
// Handles various filters and returns standardized errors with proper correlation tracking
func (d *DynamicAttributeDataSource) ReadDynamicAttributeDetails(ctx context.Context, securitySystems, endpoints, dynamicAttributes, requestTypes []string, loggedInUser, offset, max *string) (*openapi.FetchDynamicAttributesResponse, error) {
	log.Printf("[DEBUG] DynamicAttribute: Starting API call to fetch dynamic attributes")

	var apiResp *openapi.FetchDynamicAttributesResponse
	var finalHttpResp *http.Response

	// Execute API request with retry logic
	err := d.provider.AuthenticatedAPICallWithRetry(ctx, "read_dynamic_attribute_datasource", func(token string) error {
		dynAttrOps := d.dynamicAttributeFactory.CreateDynamicAttributeOperations(d.client.APIBaseURL(), token)
		
		resp, httpResp, err := dynAttrOps.FetchDynamicAttributesForDataSource(ctx, securitySystems, endpoints, dynamicAttributes, requestTypes, loggedInUser, offset, max)
		if httpResp != nil && httpResp.StatusCode == 401 {
			return fmt.Errorf("401 unauthorized")
		}
		apiResp = resp
		finalHttpResp = httpResp // Update on every call including retries
		return err
	})

	if err != nil {
		log.Printf("[ERROR] DynamicAttribute: API call failed: %v", err)
		return nil, fmt.Errorf("API call failed: %w", err)
	}

	if finalHttpResp != nil {
		log.Printf("[DEBUG] DynamicAttribute: HTTP Status Code: %d", finalHttpResp.StatusCode)
	}

	log.Printf("[INFO] DynamicAttribute: API call completed successfully")
	return apiResp, nil
}

// UpdateModelFromDynamicAttributeResponse maps API response data to the Terraform state model
func (d *DynamicAttributeDataSource) UpdateModelFromDynamicAttributeResponse(state *DynamicAttributeDataSourceModel, apiResp *openapi.FetchDynamicAttributesResponse) {
	state.Msg = util.SafeStringDatasource(apiResp.Msg)
	state.ErrorCode = util.SafeStringDatasource(apiResp.Errorcode)
	
	// Default to 0 when API omits count fields (happens when no results found)
	if apiResp.Displaycount != nil {
		state.DisplayCount = types.Int32Value(*apiResp.Displaycount)
	} else {
		state.DisplayCount = types.Int32Value(0)
	}
	if apiResp.Totalcount != nil {
		state.TotalCount = types.Int32Value(*apiResp.Totalcount)
	} else {
		state.TotalCount = types.Int32Value(0)
	}

	if apiResp.Dynamicattributes != nil && apiResp.Dynamicattributes.ArrayOfFetchDynamicAttributeResponseInner != nil {
		dynamicAttributesList := make([]DynamicAttributes, len(*apiResp.Dynamicattributes.ArrayOfFetchDynamicAttributeResponseInner))
		for i, attr := range *apiResp.Dynamicattributes.ArrayOfFetchDynamicAttributeResponseInner {
			dynamicAttributesList[i] = DynamicAttributes{
				Attributename:  util.SafeStringDatasource(attr.Attributename),
				Requesttype:    util.SafeStringDatasource(attr.Requesttype),
				Securitysystem: util.SafeStringDatasource(attr.Securitysystem),
				Endpoint:       util.SafeStringDatasource(attr.Endpoint),
				Attributetype:  util.SafeStringDatasource(attr.Attributetype),
				Attributegroup: util.SafeStringDatasource(attr.Attributegroup),
				Orderindex:     util.SafeStringDatasource(attr.Orderindex),
				Attributelable: util.SafeStringDatasource(attr.Attributelable),
				Accountscolumn: util.SafeStringDatasource(attr.Accountscolumn),
				Hideoncreate:   util.SafeStringDatasource(attr.Hideoncreate),
				Actionstring:   util.SafeStringDatasource(attr.Actionstring),
				Editable:       util.SafeStringDatasource(attr.Editable),
				Hideonupdate:   util.SafeStringDatasource(attr.Hideonupdate),
				Action_to_perform_when_parent_attribute_changes: util.SafeStringDatasource(attr.Actiontoperformwhenparentattributechanges),
				Defaultvalue:     util.SafeStringDatasource(attr.Defaultvalue),
				Required:         util.SafeStringDatasource(attr.Required),
				Attributevalue:   util.SafeStringDatasource(attr.Attributevalue),
				Showonchild:      util.SafeStringDatasource(attr.Showonchild),
				Parentattribute:  util.SafeStringDatasource(attr.Parentattribute),
				Descriptionascsv: util.SafeStringDatasource(attr.Descriptionascsv),
			}
		}
		state.Dynamicattributes = dynamicAttributesList
	} else {
		state.Dynamicattributes = nil
	}
}

// HandleDynamicAttributeAuthenticationLogic processes the authenticate flag to control sensitive data visibility
// When authenticate=false, dynamic_attributes are removed from state to prevent sensitive data exposure
// When authenticate=true, all dynamic_attributes are returned in state
func (d *DynamicAttributeDataSource) HandleDynamicAttributeAuthenticationLogic(state *DynamicAttributeDataSourceModel, resp *datasource.ReadResponse) {
	if !state.Authenticate.IsNull() && !state.Authenticate.IsUnknown() {
		if state.Authenticate.ValueBool() {
			log.Printf("[INFO] DynamicAttribute: Authentication enabled - returning all dynamic attributes")
			resp.Diagnostics.AddWarning(
				"Authentication Enabled",
				"`authenticate` is true; all dynamic_attributes will be returned in state.",
			)
		} else {
			log.Printf("[INFO] DynamicAttribute: Authentication disabled - removing dynamic attributes from state")
			resp.Diagnostics.AddWarning(
				"Authentication Disabled",
				"`authenticate` is false; dynamic_attributes will be removed from state.",
			)
			state.Dynamicattributes = nil
		}
	}
}
