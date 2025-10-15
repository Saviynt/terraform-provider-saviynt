// Copyright (c) 2025 Saviynt Inc.
// SPDX-License-Identifier: MPL-2.0

// saviynt_connections_datasource retrieves connections details from the Saviynt Security Manager.
// The data source supports a single Read operation to look up existing connections with filtering capabilities.
package provider

import (
	"context"
	"fmt"
	"terraform-provider-Saviynt/internal/client"
	"terraform-provider-Saviynt/util"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	openapi "github.com/saviynt/saviynt-api-go-client/connections"
)

var _ datasource.DataSource = &ConnectionsDataSource{}

// ConnectionsDataSource defines the data source
type ConnectionsDataSource struct {
	client            client.SaviyntClientInterface
	token             string
	provider          client.SaviyntProviderInterface
	connectionFactory client.ConnectionFactoryInterface
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
	Authenticate   types.Bool   `tfsdk:"authenticate"`
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

// NewConnectionsDataSource creates a new connections data source with default factory
func NewConnectionsDataSource() datasource.DataSource {
	return &ConnectionsDataSource{
		connectionFactory: &client.DefaultConnectionFactory{},
	}
}

// NewConnectionsDataSourceWithFactory creates a new connections data source with custom factory
// Used primarily for testing with mock factories
func NewConnectionsDataSourceWithFactory(factory client.ConnectionFactoryInterface) datasource.DataSource {
	return &ConnectionsDataSource{
		connectionFactory: factory,
	}
}

// SetClient sets the client for testing purposes
func (d *ConnectionsDataSource) SetClient(client client.SaviyntClientInterface) {
	d.client = client
}

// SetToken sets the token for testing purposes
func (d *ConnectionsDataSource) SetToken(token string) {
	d.token = token
}

// SetProvider sets the provider for testing purposes
func (d *ConnectionsDataSource) SetProvider(provider client.SaviyntProviderInterface) {
	d.provider = provider
}

// Metadata sets the data source type name for Terraform
func (d *ConnectionsDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "saviynt_connections_datasource"
}

// Schema defines the structure and attributes available for the connections data source
func (d *ConnectionsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: util.ConnDataSourceDescription,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "Resource ID.",
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
				Description: "Offset for pagination",
			},
			"display_count": schema.Int64Attribute{
				Computed:    true,
				Description: "Number of records returned in the response",
			},
			"authenticate": schema.BoolAttribute{
				Required:    true,
				Description: "If false, do not store connection results in state",
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

// Configure initializes the data source with provider configuration
// Sets up client and authentication token for API operations
func (d *ConnectionsDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	tflog.Debug(ctx, "Starting connections datasource configuration")

	// Check if provider data is available.
	if req.ProviderData == nil {
		tflog.Debug(ctx, "ProviderData is nil, returning early")
		return
	}

	// Cast provider data to your provider type.
	prov, ok := req.ProviderData.(*SaviyntProvider)
	if !ok {
		tflog.Error(ctx, "Provider configuration failed - expected *SaviyntProvider, got different type")
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

	tflog.Debug(ctx, "Connections datasource configured successfully")
}

// Read retrieves connections details from Saviynt and populates the Terraform state
// Supports filtering by connection name and type with comprehensive error handling
func (d *ConnectionsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var state ConnectionsDataSourceModel

	tflog.Debug(ctx, "Starting connections datasource read")

	// Extract configuration from request
	configDiagnostics := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(configDiagnostics...)
	if resp.Diagnostics.HasError() {
		tflog.Error(ctx, "Failed to get config from request")
		return
	}

	// Execute API call to get connections
	apiResp, err := d.ReadConnectionsDetails(ctx, &state)
	if err != nil {
		tflog.Error(ctx, "Failed to read connections details", map[string]interface{}{"error": err.Error()})
		resp.Diagnostics.AddError(
			"Connections Read Failed",
			fmt.Sprintf("Failed to read connections: %s", err.Error()),
		)
		return
	}

	// Map API response to state
	d.UpdateModelFromConnectionsResponse(&state, apiResp)

	// Handle authentication logic
	d.HandleConnectionsAuthenticationLogic(&state, resp)

	// Set final state
	stateDiagnostics := resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(stateDiagnostics...)
	if resp.Diagnostics.HasError() {
		tflog.Error(ctx, "Failed to set state")
		return
	}

	tflog.Debug(ctx, "Connections datasource read completed successfully", map[string]interface{}{
		"results_count": len(state.Results),
	})
}

// ReadConnectionsDetails retrieves connections details from Saviynt API
// Handles filtering parameters and returns standardized errors with proper correlation tracking
func (d *ConnectionsDataSource) ReadConnectionsDetails(ctx context.Context, state *ConnectionsDataSourceModel) (*openapi.GetConnectionsResponse, error) {
	tflog.Debug(ctx, "Starting connections API call")

	var apiResp *openapi.GetConnectionsResponse

	// Execute API request with retry logic
	err := d.provider.AuthenticatedAPICallWithRetry(ctx, "read_connections_datasource", func(token string) error {
		connectionOps := d.connectionFactory.CreateConnectionOperations(d.client.APIBaseURL(), token)

		// Build request with filtering parameters
		req := openapi.NewGetConnectionsRequest()

		if !state.ConnectionName.IsNull() && state.ConnectionName.ValueString() != "" {
			connectionName := state.ConnectionName.ValueString()
			req.Connectionname = &connectionName
		}

		if !state.ConnectionType.IsNull() && state.ConnectionType.ValueString() != "" {
			connectionType := state.ConnectionType.ValueString()
			req.Connectiontype = &connectionType
		}

		if !state.Offset.IsNull() && state.Offset.ValueString() != "" {
			offset := state.Offset.ValueString()
			req.Offset = &offset
		}

		if !state.Max.IsNull() && state.Max.ValueString() != "" {
			max := state.Max.ValueString()
			req.Max = &max
		}

		resp, httpResp, err := connectionOps.GetConnectionsDataSource(ctx, *req)
		if httpResp != nil && httpResp.StatusCode == 401 {
			return fmt.Errorf("401 unauthorized")
		}
		apiResp = resp
		return err
	})

	if err != nil {
		tflog.Error(ctx, "Failed to read connections details", map[string]interface{}{"error": err.Error()})
		return nil, fmt.Errorf("failed to read connections: %w", err)
	}

	tflog.Debug(ctx, "Connections API call completed successfully")
	return apiResp, nil
}

// UpdateModelFromConnectionsResponse maps API response data to the Terraform state model
func (d *ConnectionsDataSource) UpdateModelFromConnectionsResponse(state *ConnectionsDataSourceModel, apiResp *openapi.GetConnectionsResponse) {
	// Set ID for the datasource
	state.ID = types.StringValue("connections-datasource")

	// Map response metadata
	state.Msg = util.SafeStringDatasource(apiResp.Msg)
	state.DisplayCount = util.SafeInt64(apiResp.DisplayCount)
	state.ErrorCode = util.SafeStringDatasource(apiResp.ErrorCode)
	state.TotalCount = util.SafeInt64(apiResp.TotalCount)

	// Map connection results
	d.MapConnectionResults(state, apiResp)
}

// MapConnectionResults maps connection list from API response to state model
func (d *ConnectionsDataSource) MapConnectionResults(state *ConnectionsDataSourceModel, apiResp *openapi.GetConnectionsResponse) {
	if apiResp.ConnectionList == nil {
		state.Results = []Connection{}
		return
	}

	state.Results = make([]Connection, 0, len(apiResp.ConnectionList))
	for _, item := range apiResp.ConnectionList {
		connection := Connection{
			ConnectionName:        util.SafeString(item.CONNECTIONNAME),
			ConnectionType:        util.SafeString(item.CONNECTIONTYPE),
			ConnectionDescription: util.SafeString(item.CONNECTIONDESCRIPTION),
			Status:                util.SafeInt32(item.STATUS),
			CreatedBy:             util.SafeString(item.CREATEDBY),
			CreatedOn:             util.SafeString(item.CREATEDON),
			UpdatedBy:             util.SafeString(item.UPDATEDBY),
			UpdatedOn:             util.SafeString(item.UPDATEDON),
		}
		state.Results = append(state.Results, connection)
	}
}

// HandleConnectionsAuthenticationLogic processes the authenticate flag to control sensitive data visibility
// When authenticate=false, results are removed from state to prevent sensitive data exposure
// When authenticate=true, all results are returned in state
func (d *ConnectionsDataSource) HandleConnectionsAuthenticationLogic(state *ConnectionsDataSourceModel, resp *datasource.ReadResponse) {
	if !state.Authenticate.IsNull() && !state.Authenticate.IsUnknown() {
		if state.Authenticate.ValueBool() {
			tflog.Info(context.Background(), "Authentication enabled - returning all connections")
			resp.Diagnostics.AddWarning(
				"Authentication Enabled",
				"`authenticate` is true; all connections will be returned in state.",
			)
		} else {
			tflog.Info(context.Background(), "Authentication disabled - removing connections from state")
			resp.Diagnostics.AddWarning(
				"Authentication Disabled",
				"`authenticate` is false; connections will be removed from state.",
			)
			state.Results = nil
		}
	}
}
