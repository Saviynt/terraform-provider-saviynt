// Copyright (c) 2025 Saviynt Inc.
// SPDX-License-Identifier: MPL-2.0

// saviynt_dynamic_attribute_datasource retrieves dynamic attribute details from the Saviynt Security Manager.
// The data source supports a single Read operation to look up an existing dynamic attributes with various filters like attribute name, endpoint etc.

package provider

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strings"
	"terraform-provider-Saviynt/util"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	s "github.com/saviynt/saviynt-api-go-client"

	openapi "github.com/saviynt/saviynt-api-go-client/dynamicattributes"
)

type DynamicAttributeDataSource struct {
	client *s.Client
	token  string
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

var _ datasource.DataSource = &DynamicAttributeDataSource{}

func NewDynamicAttributeDataSource() datasource.DataSource {
	return &DynamicAttributeDataSource{}
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

func (d *DynamicAttributeDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
	d.client = prov.client
	d.token = prov.accessToken
}

func (d *DynamicAttributeDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var state DynamicAttributeDataSourceModel

	configDiagnostics := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(configDiagnostics...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Configure API client
	cfg := openapi.NewConfiguration()
	apiBaseURL := strings.TrimPrefix(strings.TrimPrefix(d.client.APIBaseURL(), "https://"), "http://")
	cfg.Host = apiBaseURL
	cfg.Scheme = "https"
	cfg.AddDefaultHeader("Authorization", "Bearer "+d.token)
	cfg.HTTPClient = http.DefaultClient

	apiClient := openapi.NewAPIClient(cfg)
	securitySystems := util.ConvertListToStringSlice(ctx, state.SecuritySystem)
	endpoints := util.ConvertListToStringSlice(ctx, state.Endpoint)
	dynamicAttributes := util.ConvertListToStringSlice(ctx, state.DynamicAttributes)
	requestTypes := util.ConvertListToStringSlice(ctx, state.RequestType)
	loggedInUser := util.StringPointerOrEmpty(state.LoggedInUser)
	offset := util.StringPointerOrEmpty(state.Offset)
	max := util.StringPointerOrEmpty(state.Max)

	apiReq := apiClient.DynamicAttributesAPI.FetchDynamicAttribute(ctx)
	if securitySystems != nil {
		apiReq = apiReq.Securitysystem(securitySystems)
	}
	if endpoints != nil {
		apiReq = apiReq.Endpoint(endpoints)
	}
	if dynamicAttributes != nil {
		apiReq = apiReq.Dynamicattributes(dynamicAttributes)
	}
	if requestTypes != nil {
		apiReq = apiReq.Requesttype(requestTypes)
	}
	if offset != nil {
		apiReq = apiReq.Offset(*offset)
	}
	if max != nil {
		apiReq = apiReq.Max(*max)
	}
	if loggedInUser != nil {
		apiReq = apiReq.Loggedinuser(*loggedInUser)
	}

	apiResp, httpResp, err := apiReq.Execute()
	if err != nil {
		log.Printf("[ERROR] API Call Failed: %v", err)
		resp.Diagnostics.AddError("API Call Failed", fmt.Sprintf("Error: %v", err))
		return
	}
	log.Printf("[DEBUG] HTTP Status Code: %d", httpResp.StatusCode)

	state.Msg = util.SafeStringDatasource(apiResp.Msg)
	state.ErrorCode = util.SafeStringDatasource(apiResp.Errorcode)
	state.DisplayCount = util.SafeInt32(apiResp.Displaycount)
	state.TotalCount = util.SafeInt32(apiResp.Totalcount)

	if apiResp.Dynamicattributes != nil {
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
				Defaultvalue: util.SafeStringDatasource(attr.Defaultvalue),
				Required:     util.SafeStringDatasource(attr.Required),
				// Regex:                                     util.SafeStringDatasource(attr.Regex),
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
	stateDiagnostics := resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(stateDiagnostics...)
}
