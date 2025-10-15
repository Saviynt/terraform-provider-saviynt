// Copyright (c) 2025 Saviynt Inc.
// SPDX-License-Identifier: MPL-2.0

// saviynt_endpoints_datasource retrieves endpoint details from the Saviynt Security Manager.
// The data source supports a single Read operation to look up existing endpoints with comprehensive filtering.
package provider

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"terraform-provider-Saviynt/internal/client"
	"terraform-provider-Saviynt/util"

	openapi "github.com/saviynt/saviynt-api-go-client/endpoints"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

var _ datasource.DataSource = &endpointsDataSource{}

type endpointsDataSource struct {
	client            client.SaviyntClientInterface
	token             string
	provider          client.SaviyntProviderInterface
	endpointFactory   client.EndpointFactoryInterface
}

// NewEndpointsDataSource creates a new endpoints data source with default factory
func NewEndpointsDataSource() datasource.DataSource {
	return &endpointsDataSource{
		endpointFactory: &client.DefaultEndpointFactory{},
	}
}

// NewEndpointsDataSourceWithFactory creates a new endpoints data source with custom factory
// Used primarily for testing with mock factories
func NewEndpointsDataSourceWithFactory(factory client.EndpointFactoryInterface) datasource.DataSource {
	return &endpointsDataSource{
		endpointFactory: factory,
	}
}

// SetClient sets the client for testing purposes
func (d *endpointsDataSource) SetClient(client client.SaviyntClientInterface) {
	d.client = client
}

// SetToken sets the token for testing purposes
func (d *endpointsDataSource) SetToken(token string) {
	d.token = token
}

// SetProvider sets the provider for testing purposes
func (d *endpointsDataSource) SetProvider(provider client.SaviyntProviderInterface) {
	d.provider = provider
}

type EndpointsDataSourceModel struct {
	Results        []Endpoint   `tfsdk:"results"`
	Authenticate   types.Bool   `tfsdk:"authenticate"`
	DisplayCount   types.Int64  `tfsdk:"display_count"`
	ErrorCode      types.String `tfsdk:"error_code"`
	TotalCount     types.Int64  `tfsdk:"total_count"`
	Message        types.String `tfsdk:"message"`
	EndpointName   types.String `tfsdk:"endpointname"`
	EndpointKey    types.List   `tfsdk:"endpointkey"`
	ConnectionType types.String `tfsdk:"connection_type"`
	Displayname    types.String `tfsdk:"displayname"`
	Owner          types.String `tfsdk:"owner"`
	FilterCriteria types.Map    `tfsdk:"filter_criteria"`
	Max            types.String `tfsdk:"max"`
}

type Endpoint struct {
	Id                                  types.String `tfsdk:"id"`
	Description                         types.String `tfsdk:"description"`
	StatusForUniqueAccount              types.String `tfsdk:"status_for_unique_account"`
	Requestowner                        types.String `tfsdk:"requestowner"`
	Requestable                         types.String `tfsdk:"requestable"`
	PrimaryAccountType                  types.String `tfsdk:"primary_account_type"`
	AccountTypeNoPasswordChange         types.String `tfsdk:"account_type_no_password_change"`
	ServiceAccountNameRule              types.String `tfsdk:"service_account_name_rule"`
	AccountNameValidatorRegex           types.String `tfsdk:"account_name_validator_regex"`
	AllowChangePasswordSqlquery         types.String `tfsdk:"allow_change_password_sqlquery"`
	ParentAccountPattern                types.String `tfsdk:"parent_account_pattern"`
	OwnerType                           types.String `tfsdk:"owner_type"`
	Securitysystem                      types.String `tfsdk:"securitysystem"`
	Endpointname                        types.String `tfsdk:"endpointname"`
	UpdatedBy                           types.String `tfsdk:"updated_by"`
	Accessquery                         types.String `tfsdk:"accessquery"`
	Status                              types.String `tfsdk:"status"`
	DisplayName                         types.String `tfsdk:"display_name"`
	UpdateDate                          types.String `tfsdk:"update_date"`
	AllowRemoveAllRoleOnRequest         types.String `tfsdk:"allow_remove_all_role_on_request"`
	RoleTypeAsJson                      types.String `tfsdk:"role_type_as_json"`
	EntsWithNewAccount                  types.String `tfsdk:"ents_with_new_account"`
	ConnectionconfigAsJson              types.String `tfsdk:"connectionconfig_as_json"`
	Connectionconfig                    types.String `tfsdk:"connectionconfig"`
	AccountNameRule                     types.String `tfsdk:"account_name_rule"`
	ChangePasswordAccessQuery           types.String `tfsdk:"change_password_access_query"`
	Disableaccountrequest               types.String `tfsdk:"disableaccountrequest"`
	PluginConfigs                       types.String `tfsdk:"plugin_configs"`
	DisableaccountrequestServiceAccount types.String `tfsdk:"disableaccountrequest_service_account"`
	Requestableapplication              types.String `tfsdk:"requestableapplication"`
	CreatedFrom                         types.String `tfsdk:"created_from"`
	CreatedBy                           types.String `tfsdk:"created_by"`
	CreateDate                          types.String `tfsdk:"create_date"`
	ParentEndpoint                      types.String `tfsdk:"parent_endpoint"`
	BaseLineConfig                      types.String `tfsdk:"base_line_config"`
	Requestownertype                    types.String `tfsdk:"requestownertype"`
	CreateEntTaskforRemoveAcc           types.String `tfsdk:"create_ent_taskfor_remove_acc"`
	EnableCopyAccess                    types.String `tfsdk:"enable_copy_access"`
	AccountTypeNoDeprovision            types.String `tfsdk:"account_type_no_deprovision"`
	EndpointConfig                      types.String `tfsdk:"endpoint_config"`
	Taskemailtemplates                  types.String `tfsdk:"taskemailtemplates"`
	Ownerkey                            types.String `tfsdk:"ownerkey"`
	ServiceAccountAccessQuery           types.String `tfsdk:"service_account_access_query"`
	UserAccountCorrelationRule          types.String `tfsdk:"user_account_correlation_rule"`
	StatusConfig                        types.String `tfsdk:"status_config"`
	CustomPropertyModel
	AccountCustomPropertyLabelModel
	CustomPropertyLabelModel
}

func (d *endpointsDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "saviynt_endpoints_datasource"
}

func ResultSchema() map[string]schema.Attribute {
	attrs := map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:    true,
			Description: "Unique ID of the endpoint",
		},
		"description": schema.StringAttribute{
			Computed:    true,
			Description: "Description for the endpoint",
		},
		"status_for_unique_account": schema.StringAttribute{
			Computed:    true,
			Description: "Status for unique account",
		},
		"requestowner": schema.StringAttribute{
			Computed:    true,
			Description: "Request owner",
		},
		"requestable": schema.StringAttribute{
			Computed:    true,
			Description: "Requestable flag",
		},
		"primary_account_type": schema.StringAttribute{
			Computed:    true,
			Description: "Primary account type",
		},
		"account_type_no_password_change": schema.StringAttribute{
			Computed:    true,
			Description: "Account types for which password change is not allowed",
		},
		"service_account_name_rule": schema.StringAttribute{
			Computed:    true,
			Description: "Rule for generating service account names",
		},
		"account_name_validator_regex": schema.StringAttribute{
			Computed:    true,
			Description: "Regex to validate account name",
		},
		"allow_change_password_sqlquery": schema.StringAttribute{
			Computed:    true,
			Description: "SQL query to allow change password",
		},
		"parent_account_pattern": schema.StringAttribute{
			Computed:    true,
			Description: "Pattern for parent account",
		},
		"owner_type": schema.StringAttribute{
			Computed:    true,
			Description: "Owner type of the endpoint (User/Usergroup)",
		},
		"securitysystem": schema.StringAttribute{
			Computed:    true,
			Description: "Security system associated with the endpoint",
		},
		"endpointname": schema.StringAttribute{
			Computed:    true,
			Description: "Logical name of the endpoint",
		},
		"updated_by": schema.StringAttribute{
			Computed:    true,
			Description: "User who last updated the endpoint",
		},
		"accessquery": schema.StringAttribute{
			Computed:    true,
			Description: "Query to restrict endpoint visibility",
		},
		"status": schema.StringAttribute{
			Computed:    true,
			Description: "Status of the endpoint",
		},
		"display_name": schema.StringAttribute{
			Computed:    true,
			Description: "User-friendly display name for the endpoint",
		},
		"update_date": schema.StringAttribute{
			Computed:    true,
			Description: "Date when the endpoint was last updated",
		},
		"allow_remove_all_role_on_request": schema.StringAttribute{
			Computed:    true,
			Description: "Whether remove all roles is allowed in request",
		},
		"role_type_as_json": schema.StringAttribute{
			Computed:    true,
			Description: "Role types in JSON format",
		},
		"ents_with_new_account": schema.StringAttribute{
			Computed:    true,
			Description: "Entitlements associated with new account",
		},
		"connectionconfig_as_json": schema.StringAttribute{
			Computed:    true,
			Description: "Connection configuration in JSON",
		},
		"connectionconfig": schema.StringAttribute{
			Computed:    true,
			Description: "Connection configuration",
		},
		"account_name_rule": schema.StringAttribute{
			Computed:    true,
			Description: "Rule to generate account names",
		},
		"change_password_access_query": schema.StringAttribute{
			Computed:    true,
			Description: "Query to restrict password change",
		},
		"disableaccountrequest": schema.StringAttribute{
			Computed:    true,
			Description: "Disable account request",
		},
		"plugin_configs": schema.StringAttribute{
			Computed:    true,
			Description: "Plugin configuration for SmartAssist",
		},
		"disableaccountrequest_service_account": schema.StringAttribute{
			Computed:    true,
			Description: "Disable account request for service accounts",
		},
		"requestableapplication": schema.StringAttribute{
			Computed:    true,
			Description: "Associated requestable application",
		},
		"created_from": schema.StringAttribute{
			Computed:    true,
			Description: "Source of creation",
		},
		"created_by": schema.StringAttribute{
			Computed:    true,
			Description: "User who created the endpoint",
		},
		"create_date": schema.StringAttribute{
			Computed:    true,
			Description: "Date of creation",
		},
		"parent_endpoint": schema.StringAttribute{
			Computed:    true,
			Description: "Parent endpoint",
		},
		"base_line_config": schema.StringAttribute{
			Computed:    true,
			Description: "Baseline configuration",
		},
		"requestownertype": schema.StringAttribute{
			Computed:    true,
			Description: "Type of request owner",
		},
		"create_ent_taskfor_remove_acc": schema.StringAttribute{
			Computed:    true,
			Description: "Whether entitlement task is created for remove account",
		},
		"enable_copy_access": schema.StringAttribute{
			Computed:    true,
			Description: "Whether copy access is enabled",
		},
		"account_type_no_deprovision": schema.StringAttribute{
			Computed:    true,
			Description: "Account types not allowed for deprovision",
		},
		"endpoint_config": schema.StringAttribute{
			Computed:    true,
			Description: "Endpoint configuration",
		},
		"taskemailtemplates": schema.StringAttribute{
			Computed:    true,
			Description: "Task email templates",
		},
		"ownerkey": schema.StringAttribute{
			Computed:    true,
			Description: "Key of the owner",
		},
		"service_account_access_query": schema.StringAttribute{
			Computed:    true,
			Description: "Query to filter service account access",
		},
		"user_account_correlation_rule": schema.StringAttribute{
			Computed:    true,
			Description: "Rule to correlate user and account",
		},
		"status_config": schema.StringAttribute{
			Computed:    true,
			Description: "Status configuration for account operations",
		},
	}

	for i := 1; i <= 45; i++ {
		key := fmt.Sprintf("custom_property%d", i)
		attrs[key] = schema.StringAttribute{
			Computed:    true,
			Description: fmt.Sprintf("Custom property %d value for the endpoint.", i),
		}
	}

	for i := 1; i <= 30; i++ {
		key := fmt.Sprintf("account_custom_property_%d_label", i)
		attrs[key] = schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: fmt.Sprintf("Label for account custom property %d.", i),
		}
	}

	for i := 31; i <= 60; i++ {
		key := fmt.Sprintf("custom_property%d_label", i)
		attrs[key] = schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: fmt.Sprintf("Label for custom property %d of accounts of this endpoint.", i),
		}
	}
	return attrs
}

func (d *endpointsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: util.EndpointDataSourceDescription,
		Attributes: map[string]schema.Attribute{
			"endpointname": schema.StringAttribute{
				Optional:    true,
				Description: "Filter by endpoint name",
			},
			"endpointkey": schema.ListAttribute{
				Optional:    true,
				ElementType: types.StringType,
				Description: "List of endpoint keys to filter",
			},
			"authenticate": schema.BoolAttribute{
				Required:    true,
				Description: "If false, do not store endpoints details in state",
			},
			"connection_type": schema.StringAttribute{
				Optional:    true,
				Description: "Filter by connection type",
			},
			"displayname": schema.StringAttribute{
				Optional:    true,
				Description: "Filter by display name",
			},
			"owner": schema.StringAttribute{
				Optional:    true,
				Description: "Filter by owner",
			},
			"filter_criteria": schema.MapAttribute{
				ElementType: types.StringType,
				Optional:    true,
				Description: "Filter criteria",
			},
			"max": schema.StringAttribute{
				Optional: true,
			},
			"message": schema.StringAttribute{
				Computed:    true,
				Description: "API response message",
			},
			"display_count": schema.Int64Attribute{
				Computed:    true,
				Description: "Number of records returned in the response",
			},
			"error_code": schema.StringAttribute{
				Computed:    true,
				Description: "Error code from the API",
			},
			"total_count": schema.Int64Attribute{
				Computed:    true,
				Description: "Total count of available records",
			},
			"results": schema.ListNestedAttribute{
				Computed:    true,
				Description: "List of endpoints retrieved",
				NestedObject: schema.NestedAttributeObject{
					Attributes: ResultSchema(),
				},
			},
		},
	}
}

// Configure initializes the data source with provider configuration
// Sets up client and authentication token for API operations
func (d *endpointsDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	tflog.Debug(ctx, "Starting endpoints datasource configuration")

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

	tflog.Debug(ctx, "Endpoints datasource configured successfully")
}

// Read retrieves endpoints details from Saviynt and populates the Terraform state
// Supports comprehensive filtering with structured error handling and operation tracking
func (d *endpointsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var state EndpointsDataSourceModel

	tflog.Debug(ctx, "Starting endpoints datasource read")

	// Extract configuration from request
	configDiagnostics := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(configDiagnostics...)
	if resp.Diagnostics.HasError() {
		tflog.Error(ctx, "Failed to extract configuration from request")
		return
	}

	// Execute API call to get endpoints details
	apiResp, err := d.ReadEndpointsDetails(ctx, &state)
	if err != nil {
		tflog.Error(ctx, "Failed to read endpoints details", map[string]interface{}{"error": err.Error()})
		resp.Diagnostics.AddError(
			"Endpoints Read Failed",
			fmt.Sprintf("Failed to read endpoints: %s", err.Error()),
		)
		return
	}

	// Map API response to state
	d.UpdateModelFromEndpointsResponse(&state, apiResp)

	// Handle authentication logic
	d.HandleEndpointsAuthenticationLogic(&state, resp)

	// Set final state
	stateDiagnostics := resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(stateDiagnostics...)
	if resp.Diagnostics.HasError() {
		tflog.Error(ctx, "Failed to set state")
		return
	}

	tflog.Debug(ctx, "Endpoints datasource read completed successfully", map[string]interface{}{
		"results_count": len(state.Results),
		"total_count":   state.TotalCount.ValueInt64(),
	})
}
				
// ReadEndpointsDetails retrieves endpoints details from Saviynt API
// Handles comprehensive filtering using factory pattern with standardized error handling
func (d *endpointsDataSource) ReadEndpointsDetails(ctx context.Context, state *EndpointsDataSourceModel) (*openapi.GetEndpoints200Response, error) {
	tflog.Debug(ctx, "Starting endpoints API call")

	var apiResp *openapi.GetEndpoints200Response
	var finalHttpResp *http.Response

	// Execute API request with retry logic
	err := d.provider.AuthenticatedAPICallWithRetry(ctx, "read_endpoints_datasource", func(token string) error {
		endpointOps := d.endpointFactory.CreateEndpointOperations(d.client.APIBaseURL(), token)

		// Build request with filters
		req := d.BuildEndpointsRequest(ctx, state)

		resp, httpResp, err := endpointOps.GetEndpoints(ctx, req)
		finalHttpResp = httpResp
		if httpResp != nil && httpResp.StatusCode == 401 {
			return fmt.Errorf("401 unauthorized")
		}
		apiResp = resp
		return err
	})

	if err != nil {
		// Handle specific HTTP errors
		if finalHttpResp != nil && finalHttpResp.StatusCode == 412 {
			tflog.Error(ctx, "HTTP 412 error while reading endpoints", map[string]interface{}{"status": finalHttpResp.Status})
			return nil, d.HandleHTTP412Error(finalHttpResp)
		}
		
		tflog.Error(ctx, "Failed to read endpoints details", map[string]interface{}{"error": err.Error()})
		return nil, fmt.Errorf("failed to read endpoints: %w", err)
	}

	// Validate API response
	if err := d.ValidateEndpointsResponse(apiResp); err != nil {
		tflog.Error(ctx, "API returned error response", map[string]interface{}{"error": err.Error()})
		return nil, fmt.Errorf("API error: %w", err)
	}

	tflog.Debug(ctx, "Endpoints API call completed successfully")
	return apiResp, nil
}

// BuildEndpointsRequest constructs the API request with all filters from state
func (d *endpointsDataSource) BuildEndpointsRequest(ctx context.Context, state *EndpointsDataSourceModel) openapi.GetEndpointsRequest {
	req := openapi.GetEndpointsRequest{}

	// Apply filters
	if !state.EndpointName.IsNull() && state.EndpointName.ValueString() != "" {
		req.SetEndpointname(state.EndpointName.ValueString())
	}

	if !state.EndpointKey.IsNull() && len(state.EndpointKey.Elements()) > 0 {
		endpointkeys := util.StringsFromList(state.EndpointKey)
		req.SetEndpointkey(endpointkeys)
	}

	if !state.ConnectionType.IsNull() && state.ConnectionType.ValueString() != "" {
		req.SetConnectionType(state.ConnectionType.ValueString())
	}

	if !state.Displayname.IsNull() && state.Displayname.ValueString() != "" {
		req.SetDisplayName(state.Displayname.ValueString())
	}

	if !state.Owner.IsNull() && state.Owner.ValueString() != "" {
		req.SetOwner(state.Owner.ValueString())
	}

	if !state.Max.IsNull() && state.Max.ValueString() != "" {
		req.SetMax(state.Max.ValueString())
	}

	if !state.FilterCriteria.IsNull() {
		var filterMap map[string]string
		diags := state.FilterCriteria.ElementsAs(ctx, &filterMap, true)
		if !diags.HasError() {
			filterCriteria := make(map[string]interface{}, len(filterMap))
			for k, v := range filterMap {
				filterCriteria[k] = v
			}
			req.SetFilterCriteria(filterCriteria)
		}
	}

	return req
}

// HandleHTTP412Error handles 412 Precondition Failed errors with response body parsing
func (d *endpointsDataSource) HandleHTTP412Error(httpResp *http.Response) error {
	if httpResp.Body != nil {
		var fetchResp map[string]interface{}
		if err := json.NewDecoder(httpResp.Body).Decode(&fetchResp); err != nil {
			tflog.Error(context.Background(), "Failed to decode 412 error response", map[string]interface{}{"decode_error": err.Error()})
			return fmt.Errorf("412 Precondition Failed - failed to decode error response: %w", err)
		}
		return fmt.Errorf("412 Precondition Failed: %v", fetchResp)
	}
	return fmt.Errorf("412 Precondition Failed - no response body")
}

// ValidateEndpointsResponse validates that the API response is successful
func (d *endpointsDataSource) ValidateEndpointsResponse(apiResp *openapi.GetEndpoints200Response) error {
	if apiResp != nil && apiResp.ErrorCode != nil && *apiResp.ErrorCode != "0" {
		errorCode := util.SafeDeref(apiResp.ErrorCode)
		message := util.SafeDeref(apiResp.Message)
		tflog.Error(context.Background(), "API returned error response", map[string]interface{}{
			"error_code": errorCode,
			"message":    message,
		})
		return fmt.Errorf("API error - ErrorCode: %s, Message: %s", errorCode, message)
	}
	return nil
}

// UpdateModelFromEndpointsResponse maps API response data to the Terraform state model
func (d *endpointsDataSource) UpdateModelFromEndpointsResponse(state *EndpointsDataSourceModel, apiResp *openapi.GetEndpoints200Response) {
	// Map response metadata
	state.Message = util.SafeStringDatasource(apiResp.Message)
	state.ErrorCode = util.SafeStringDatasource(apiResp.ErrorCode)
	
	// Handle count fields with defaults
	if apiResp.DisplayCount != nil {
		state.DisplayCount = types.Int64Value(int64(*apiResp.DisplayCount))
	} else {
		state.DisplayCount = types.Int64Value(0)
	}
	
	if apiResp.TotalCount != nil {
		state.TotalCount = types.Int64Value(int64(*apiResp.TotalCount))
	} else {
		state.TotalCount = types.Int64Value(0)
	}

	// Map endpoints data
	d.MapEndpointsData(state, apiResp)
}

// MapEndpointsData maps individual endpoint records from API response to state
func (d *endpointsDataSource) MapEndpointsData(state *EndpointsDataSourceModel, apiResp *openapi.GetEndpoints200Response) {
	if apiResp.Endpoints == nil {
		state.Results = nil
		return
	}

	state.Results = make([]Endpoint, 0, len(apiResp.Endpoints))
	for _, item := range apiResp.Endpoints {
		endpointState := Endpoint{
			Id:                                  util.SafeString(item.Id),
			Description:                         util.SafeString(item.Description),
			StatusForUniqueAccount:              util.SafeString(item.StatusForUniqueAccount),
			Requestowner:                        util.SafeString(item.Requestowner),
			Requestable:                         util.SafeString(item.Requestable),
			PrimaryAccountType:                  util.SafeString(item.PrimaryAccountType),
			AccountTypeNoPasswordChange:         util.SafeString(item.AccountTypeNoPasswordChange),
			ServiceAccountNameRule:              util.SafeString(item.ServiceAccountNameRule),
			AccountNameValidatorRegex:           util.SafeString(item.AccountNameValidatorRegex),
			AllowChangePasswordSqlquery:         util.SafeString(item.AllowChangePasswordSqlquery),
			ParentAccountPattern:                util.SafeString(item.ParentAccountPattern),
			OwnerType:                           util.SafeString(item.OwnerType),
			Securitysystem:                      util.SafeString(item.Securitysystem),
			Endpointname:                        util.SafeString(item.Endpointname),
			UpdatedBy:                           util.SafeString(item.UpdatedBy),
			Accessquery:                         util.SafeString(item.Accessquery),
			Status:                              util.SafeString(item.Status),
			DisplayName:                         util.SafeString(item.DisplayName),
			UpdateDate:                          util.SafeString(item.UpdateDate),
			AllowRemoveAllRoleOnRequest:         util.SafeString(item.AllowRemoveAllRoleOnRequest),
			RoleTypeAsJson:                      util.SafeString(item.RoleTypeAsJson),
			EntsWithNewAccount:                  util.SafeString(item.EntsWithNewAccount),
			ConnectionconfigAsJson:              util.SafeString(item.ConnectionconfigAsJson),
			Connectionconfig:                    util.SafeString(item.Connectionconfig),
			AccountNameRule:                     util.SafeString(item.AccountNameRule),
			ChangePasswordAccessQuery:           util.SafeString(item.ChangePasswordAccessQuery),
			Disableaccountrequest:               util.SafeString(item.Disableaccountrequest),
			PluginConfigs:                       util.SafeString(item.PluginConfigs),
			DisableaccountrequestServiceAccount: util.SafeString(item.DisableaccountrequestServiceAccount),
			Requestableapplication:              util.SafeString(item.Requestableapplication),
			CreatedFrom:                         util.SafeString(item.CreatedFrom),
			CreatedBy:                           util.SafeString(item.CreatedBy),
			CreateDate:                          util.SafeString(item.CreateDate),
			ParentEndpoint:                      util.SafeString(item.ParentEndpoint),
			BaseLineConfig:                      util.SafeString(item.BaseLineConfig),
			Requestownertype:                    util.SafeString(item.Requestownertype),
			CreateEntTaskforRemoveAcc:           util.SafeString(item.CreateEntTaskforRemoveAcc),
			EnableCopyAccess:                    util.SafeString(item.EnableCopyAccess),
			AccountTypeNoDeprovision:            util.SafeString(item.AccountTypeNoDeprovision),
			EndpointConfig:                      util.SafeString(item.EndpointConfig),
			Taskemailtemplates:                  util.SafeString(item.Taskemailtemplates),
			Ownerkey:                            util.SafeString(item.Ownerkey),
			ServiceAccountAccessQuery:           util.SafeString(item.ServiceAccountAccessQuery),
			UserAccountCorrelationRule:          util.SafeString(item.UserAccountCorrelationRule),
			StatusConfig:                        util.SafeString(item.StatusConfig),
		}

		// Map custom properties
		d.MapCustomProperties(&endpointState, &item)
		state.Results = append(state.Results, endpointState)
	}
}

// MapCustomProperties maps custom properties and labels from API response to endpoint state
func (d *endpointsDataSource) MapCustomProperties(endpointState *Endpoint, item *openapi.GetEndpoints200ResponseEndpointsInner) {
	endpointState.CustomPropertyModel = CustomPropertyModel{
		CustomProperty1:  util.SafeString(item.CustomProperty1),
		CustomProperty2:  util.SafeString(item.CustomProperty2),
		CustomProperty3:  util.SafeString(item.CustomProperty3),
		CustomProperty4:  util.SafeString(item.CustomProperty4),
		CustomProperty5:  util.SafeString(item.CustomProperty5),
		CustomProperty6:  util.SafeString(item.CustomProperty6),
		CustomProperty7:  util.SafeString(item.CustomProperty7),
		CustomProperty8:  util.SafeString(item.CustomProperty8),
		CustomProperty9:  util.SafeString(item.CustomProperty9),
		CustomProperty10: util.SafeString(item.CustomProperty10),
		CustomProperty11: util.SafeString(item.CustomProperty11),
		CustomProperty12: util.SafeString(item.CustomProperty12),
		CustomProperty13: util.SafeString(item.CustomProperty13),
		CustomProperty14: util.SafeString(item.CustomProperty14),
		CustomProperty15: util.SafeString(item.CustomProperty15),
		CustomProperty16: util.SafeString(item.CustomProperty16),
		CustomProperty17: util.SafeString(item.CustomProperty17),
		CustomProperty18: util.SafeString(item.CustomProperty18),
		CustomProperty19: util.SafeString(item.CustomProperty19),
		CustomProperty20: util.SafeString(item.CustomProperty20),
		CustomProperty21: util.SafeString(item.CustomProperty21),
		CustomProperty22: util.SafeString(item.CustomProperty22),
		CustomProperty23: util.SafeString(item.CustomProperty23),
		CustomProperty24: util.SafeString(item.CustomProperty24),
		CustomProperty25: util.SafeString(item.CustomProperty25),
		CustomProperty26: util.SafeString(item.CustomProperty26),
		CustomProperty27: util.SafeString(item.CustomProperty27),
		CustomProperty28: util.SafeString(item.CustomProperty28),
		CustomProperty29: util.SafeString(item.CustomProperty29),
		CustomProperty30: util.SafeString(item.CustomProperty30),
		CustomProperty31: util.SafeString(item.Customproperty31),
		CustomProperty32: util.SafeString(item.Customproperty32),
		CustomProperty33: util.SafeString(item.Customproperty33),
		CustomProperty34: util.SafeString(item.Customproperty34),
		CustomProperty35: util.SafeString(item.Customproperty35),
		CustomProperty36: util.SafeString(item.Customproperty36),
		CustomProperty37: util.SafeString(item.Customproperty37),
		CustomProperty38: util.SafeString(item.Customproperty38),
		CustomProperty39: util.SafeString(item.Customproperty39),
		CustomProperty40: util.SafeString(item.Customproperty40),
		CustomProperty41: util.SafeString(item.Customproperty41),
		CustomProperty42: util.SafeString(item.Customproperty42),
		CustomProperty43: util.SafeString(item.Customproperty43),
		CustomProperty44: util.SafeString(item.Customproperty44),
		CustomProperty45: util.SafeString(item.Customproperty45),
	}

	endpointState.AccountCustomPropertyLabelModel = AccountCustomPropertyLabelModel{
		AccountCustomProperty1Label:  util.SafeString(item.AccountCustomProperty1Label),
		AccountCustomProperty2Label:  util.SafeString(item.AccountCustomProperty2Label),
		AccountCustomProperty3Label:  util.SafeString(item.AccountCustomProperty3Label),
		AccountCustomProperty4Label:  util.SafeString(item.AccountCustomProperty4Label),
		AccountCustomProperty5Label:  util.SafeString(item.AccountCustomProperty5Label),
		AccountCustomProperty6Label:  util.SafeString(item.AccountCustomProperty6Label),
		AccountCustomProperty7Label:  util.SafeString(item.AccountCustomProperty7Label),
		AccountCustomProperty8Label:  util.SafeString(item.AccountCustomProperty8Label),
		AccountCustomProperty9Label:  util.SafeString(item.AccountCustomProperty9Label),
		AccountCustomProperty10Label: util.SafeString(item.AccountCustomProperty10Label),
		AccountCustomProperty11Label: util.SafeString(item.AccountCustomProperty11Label),
		AccountCustomProperty12Label: util.SafeString(item.AccountCustomProperty12Label),
		AccountCustomProperty13Label: util.SafeString(item.AccountCustomProperty13Label),
		AccountCustomProperty14Label: util.SafeString(item.AccountCustomProperty14Label),
		AccountCustomProperty15Label: util.SafeString(item.AccountCustomProperty15Label),
		AccountCustomProperty16Label: util.SafeString(item.AccountCustomProperty16Label),
		AccountCustomProperty17Label: util.SafeString(item.AccountCustomProperty17Label),
		AccountCustomProperty18Label: util.SafeString(item.AccountCustomProperty18Label),
		AccountCustomProperty19Label: util.SafeString(item.AccountCustomProperty19Label),
		AccountCustomProperty20Label: util.SafeString(item.AccountCustomProperty20Label),
		AccountCustomProperty21Label: util.SafeString(item.AccountCustomProperty21Label),
		AccountCustomProperty22Label: util.SafeString(item.AccountCustomProperty22Label),
		AccountCustomProperty23Label: util.SafeString(item.AccountCustomProperty23Label),
		AccountCustomProperty24Label: util.SafeString(item.AccountCustomProperty24Label),
		AccountCustomProperty25Label: util.SafeString(item.AccountCustomProperty25Label),
		AccountCustomProperty26Label: util.SafeString(item.AccountCustomProperty26Label),
		AccountCustomProperty27Label: util.SafeString(item.AccountCustomProperty27Label),
		AccountCustomProperty28Label: util.SafeString(item.AccountCustomProperty28Label),
		AccountCustomProperty29Label: util.SafeString(item.AccountCustomProperty29Label),
		AccountCustomProperty30Label: util.SafeString(item.AccountCustomProperty30Label),
	}

	endpointState.CustomPropertyLabelModel = CustomPropertyLabelModel{
		CustomProperty31Label: util.SafeString(item.Customproperty31Label),
		CustomProperty32Label: util.SafeString(item.Customproperty32Label),
		CustomProperty33Label: util.SafeString(item.Customproperty33Label),
		CustomProperty34Label: util.SafeString(item.Customproperty34Label),
		CustomProperty35Label: util.SafeString(item.Customproperty35Label),
		CustomProperty36Label: util.SafeString(item.Customproperty36Label),
		CustomProperty37Label: util.SafeString(item.Customproperty37Label),
		CustomProperty38Label: util.SafeString(item.Customproperty38Label),
		CustomProperty39Label: util.SafeString(item.Customproperty39Label),
		CustomProperty40Label: util.SafeString(item.Customproperty40Label),
		CustomProperty41Label: util.SafeString(item.Customproperty41Label),
		CustomProperty42Label: util.SafeString(item.Customproperty42Label),
		CustomProperty43Label: util.SafeString(item.Customproperty43Label),
		CustomProperty44Label: util.SafeString(item.Customproperty44Label),
		CustomProperty45Label: util.SafeString(item.Customproperty45Label),
		CustomProperty46Label: util.SafeString(item.Customproperty46Label),
		CustomProperty47Label: util.SafeString(item.Customproperty47Label),
		CustomProperty48Label: util.SafeString(item.Customproperty48Label),
		CustomProperty49Label: util.SafeString(item.Customproperty49Label),
		CustomProperty50Label: util.SafeString(item.Customproperty50Label),
		CustomProperty51Label: util.SafeString(item.Customproperty51Label),
		CustomProperty52Label: util.SafeString(item.Customproperty52Label),
		CustomProperty53Label: util.SafeString(item.Customproperty53Label),
		CustomProperty54Label: util.SafeString(item.Customproperty54Label),
		CustomProperty55Label: util.SafeString(item.Customproperty55Label),
		CustomProperty56Label: util.SafeString(item.Customproperty56Label),
		CustomProperty57Label: util.SafeString(item.Customproperty57Label),
		CustomProperty58Label: util.SafeString(item.Customproperty58Label),
		CustomProperty59Label: util.SafeString(item.Customproperty59Label),
		CustomProperty60Label: util.SafeString(item.Customproperty60Label),
	}
}

// HandleEndpointsAuthenticationLogic processes the authenticate flag to control sensitive data visibility
// When authenticate=false, results are removed from state to prevent sensitive data exposure
// When authenticate=true, all endpoint details are returned in state
func (d *endpointsDataSource) HandleEndpointsAuthenticationLogic(state *EndpointsDataSourceModel, resp *datasource.ReadResponse) {
	if !state.Authenticate.IsNull() && !state.Authenticate.IsUnknown() {
		if state.Authenticate.ValueBool() {
			tflog.Info(context.Background(), "Authentication enabled - returning all endpoint details")
			resp.Diagnostics.AddWarning(
				"Authentication Enabled",
				"`authenticate` is true; all endpoint details will be returned in state.",
			)
		} else {
			tflog.Info(context.Background(), "Authentication disabled - removing endpoint details from state")
			resp.Diagnostics.AddWarning(
				"Authentication Disabled",
				"`authenticate` is false; endpoint details will be removed from state.",
			)
			state.Results = nil
		}
	}
}
