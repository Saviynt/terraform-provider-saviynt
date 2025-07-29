// Copyright (c) 2025 Saviynt Inc.
// SPDX-License-Identifier: MPL-2.0

// saviynt_connections_datasource retrieves connections details from the Saviynt Security Manager.
// The data source supports a single Read operation to look up an existing sap connections by name.
package provider

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strings"
	"terraform-provider-Saviynt/util"

	openapi "github.com/saviynt/saviynt-api-go-client/connections"

	s "github.com/saviynt/saviynt-api-go-client"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type connectionsDataSource struct {
	client *s.Client
	token  string
}

var _ datasource.DataSource = &connectionsDataSource{}
var _ datasource.DataSourceWithConfigure = &connectionsDataSource{}

func NewConnectionsDataSource() datasource.DataSource {
	return &connectionsDataSource{}
}

type ConnectionsDataSourceModel struct {
	ID             types.String `tfsdk:"id"`
	Results        []Connection `tfsdk:"results"`
	ConnectionName types.String `tfsdk:"connection_name"`
	Offset         types.String `tfsdk:"offset"`
	DisplayCount   types.Int64  `tfsdk:"display_count"`
	ErrorCode      types.String `tfsdk:"error_code"`
	TotalCount     types.Int64  `tfsdk:"total_count"`
	Msg            types.String `tfsdk:"msg"`
	ConnectionType types.String `tfsdk:"connection_type"`
	Max            types.String `tfsdk:"max"`
	Authenticate    types.Bool   `tfsdk:"authenticate"`
}

type Connection struct {
	ConnectionName        types.String `tfsdk:"connectionname"`
	ConnectionType        types.String `tfsdk:"connectiontype"`
	ConnectionDescription types.String `tfsdk:"connectiondescription"`
	Status                types.Int32  `tfsdk:"status"`
	CreatedBy             types.String `tfsdk:"createdby"`
	CreatedOn             types.String `tfsdk:"createdon"`
	UpdatedBy             types.String `tfsdk:"updatedby"`
	UpdatedOn             types.String `tfsdk:"updatedon"`
}

func (d *connectionsDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "saviynt_connections_datasource"
}

func (d *connectionsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: util.ConnDataSourceDescription,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"connection_name": schema.StringAttribute{
				Optional:    true,
				Description: "Filter by connection name",
			},
			"connection_type": schema.StringAttribute{
				Optional:    true,
				Description: "Filter by connection type",
			},
			"max": schema.StringAttribute{
				Optional:    true,
				Description: "Maximum number of connections to retrieve",
			},
			"offset": schema.StringAttribute{
				Optional:    true,
				Description: "Offset",
			},
			"display_count": schema.Int64Attribute{
				Computed:    true,
				Description: "Number of records returned in the response",
			},
			"authenticate": schema.BoolAttribute{
				Required:    true,
				Description: "If false, do not store connection_attributes in state",
			},
			"error_code": schema.StringAttribute{
				Computed:    true,
				Description: "Error code from the API",
			},
			"total_count": schema.Int64Attribute{
				Computed:    true,
				Description: "Total count of available connections",
			},
			"msg": schema.StringAttribute{
				Computed:    true,
				Description: "API response message",
			},
			"results": schema.ListNestedAttribute{
				Computed:    true,
				Description: "List of connections retrieved",
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"connectionname": schema.StringAttribute{
							Computed:    true,
							Description: "Connection Name",
						},
						"connectiontype": schema.StringAttribute{
							Computed:    true,
							Description: "Type of connection",
						},
						"connectiondescription": schema.StringAttribute{
							Computed:    true,
							Description: "Description of the connection",
						},
						"status": schema.Int64Attribute{
							Computed:    true,
							Description: "Status of the connection",
						},
						"createdby": schema.StringAttribute{
							Computed:    true,
							Description: "User who created the connection",
						},
						"createdon": schema.StringAttribute{
							Computed:    true,
							Description: "Timestamp when the connection was created",
						},
						"updatedby": schema.StringAttribute{
							Computed:    true,
							Description: "User who last updated the connection",
						},
						"updatedon": schema.StringAttribute{
							Computed:    true,
							Description: "Timestamp when the connection was last updated",
						},
					},
				},
			},
		},
	}
}

func (d *connectionsDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (d *connectionsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var state ConnectionsDataSourceModel

	configDiagnostics := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(configDiagnostics...)

	if resp.Diagnostics.HasError() {
		return
	}

	cfg := openapi.NewConfiguration()
	apiBaseURL := strings.TrimPrefix(strings.TrimPrefix(d.client.APIBaseURL(), "https://"), "http://")

	cfg.Host = apiBaseURL
	cfg.Scheme = "https"
	cfg.AddDefaultHeader("Authorization", "Bearer "+d.token)
	cfg.HTTPClient = http.DefaultClient

	apiClient := openapi.NewAPIClient(cfg)

	areq := openapi.NewGetConnectionsRequest()

	if !state.ConnectionName.IsNull() && state.ConnectionName.ValueString() != "" {
		connectionName := state.ConnectionName.ValueString()
		areq.Connectionname = &connectionName
	}

	if !state.ConnectionType.IsNull() && state.ConnectionType.ValueString() != "" {
		connectionType := state.ConnectionType.ValueString()
		areq.Connectiontype = &connectionType
	}

	if !state.Offset.IsNull() && state.Offset.ValueString() != "" {
		offset := state.Offset.ValueString()
		areq.Offset = &offset
	}

	if !state.Max.IsNull() && state.Max.ValueString() != "" {
		max := state.Max.ValueString()
		areq.Max = &max
	}

	apiReq := apiClient.ConnectionsAPI.GetConnections(ctx).GetConnectionsRequest(*areq)

	connectionsResponse, httpResp, err := apiReq.Execute()
	if err != nil {
		log.Printf("[ERROR] API Call Failed: %v", err)
		resp.Diagnostics.AddError("API Call Failed", fmt.Sprintf("Error: %v", err))
		return
	}
	log.Printf("[DEBUG] HTTP Status Code: %d", httpResp.StatusCode)

	state.Msg = types.StringValue(*connectionsResponse.Msg)
	state.DisplayCount = types.Int64Value(int64(*connectionsResponse.DisplayCount))
	state.ErrorCode = types.StringValue(*connectionsResponse.ErrorCode)
	state.TotalCount = types.Int64Value(int64(*connectionsResponse.TotalCount))

	if connectionsResponse != nil && connectionsResponse.ConnectionList != nil {
		for _, item := range connectionsResponse.ConnectionList {
			resultState := Connection{
				ConnectionName:        util.SafeString(item.CONNECTIONNAME),
				ConnectionType:        util.SafeString(item.CONNECTIONTYPE),
				ConnectionDescription: util.SafeString(item.CONNECTIONDESCRIPTION),
				Status:                util.SafeInt32(item.STATUS),
				CreatedBy:             util.SafeString(item.CREATEDBY),
				CreatedOn:             util.SafeString(item.CREATEDON),
				UpdatedBy:             util.SafeString(item.UPDATEDBY),
				UpdatedOn:             util.SafeString(item.UPDATEDON),
			}

			state.Results = append(state.Results, resultState)
		}
	}

	if !state.Authenticate.IsNull() && !state.Authenticate.IsUnknown() {
		if state.Authenticate.ValueBool() {
			resp.Diagnostics.AddWarning(
				"Authentication Enabled",
				"`authenticate` is true; all connections will be returned in state.",
			)
		} else {
			resp.Diagnostics.AddWarning(
				"Authentication Disabled",
				"`authenticate` is false; connections will be removed from state.",
			)
			state.Results = nil
		}
	}

	stateDiagnostics := resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(stateDiagnostics...)

	if resp.Diagnostics.HasError() {
		return
	}
}
