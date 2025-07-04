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

// saviynt_github_rest_connection_datasource retrieves github rest connections details from the Saviynt Security Manager.
// The data source supports a single Read operation to look up an existing github rest connections by name.
package provider

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"terraform-provider-Saviynt/util"
	connectionsutil "terraform-provider-Saviynt/util/connectionsutil"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	s "github.com/saviynt/saviynt-api-go-client"

	openapi "github.com/saviynt/saviynt-api-go-client/connections"
)

var _ datasource.DataSource = &githubRestConnectionDataSource{}

// GithubRestConnectionDataSource defines the data source
type githubRestConnectionDataSource struct {
	client *s.Client
	token  string
}

type GithubRestConnectionDataSourceModel struct {
	ID types.String `tfsdk:"id"`
	BaseConnectionDataSourceModel
	ConnectionAttributes *GithubRestConnectionAttributes `tfsdk:"connection_attributes"`
}

type GithubRestConnectionAttributes struct {
	IsTimeoutSupported types.Bool `tfsdk:"is_timeout_supported"`
	// ConnectionJSON       types.String `tfsdk:"connection_json"`
	OrganizationList     types.String `tfsdk:"organization_list"`
	ImportAccountEntJSON types.String `tfsdk:"import_account_ent_json"`
	// StatusThresholdConfig    types.String             `tfsdk:"status_threshold_config"`
	// AccessTokens types.String `tfsdk:"access_tokens"`
	// ConnectionTimeoutConfig  *ConnectionTimeoutConfig `tfsdk:"connection_timeout_config"`
	ConnectionType           types.String `tfsdk:"connection_type"`
	IsTimeoutConfigValidated types.Bool   `tfsdk:"is_timeout_config_validated"`
}

func NewGithubRestConnectionsDataSource() datasource.DataSource {
	return &githubRestConnectionDataSource{}
}

func (d *githubRestConnectionDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "saviynt_github_rest_connection_datasource"
}

func GithubRestConnectorsDataSourceSchema() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:    true,
			Description: "Resource ID.",
		},
		"connection_attributes": schema.SingleNestedAttribute{
			Computed: true,
			Attributes: map[string]schema.Attribute{
				"is_timeout_supported": schema.BoolAttribute{Computed: true},
				// "connection_json":         schema.StringAttribute{Computed: true},
				"organization_list":       schema.StringAttribute{Computed: true},
				"connection_type":         schema.StringAttribute{Computed: true},
				"import_account_ent_json": schema.StringAttribute{Computed: true},
				// "status_threshold_config":     schema.StringAttribute{Computed: true},
				// "access_tokens":               schema.StringAttribute{Computed: true},
				"is_timeout_config_validated": schema.BoolAttribute{Computed: true},
				// "connection_timeout_config": schema.SingleNestedAttribute{
				// 	Computed:   true,
				// 	Attributes: ConnectionTimeoutConfigeSchema(),
				// },
			},
		},
	}
}

// Schema defines the attributes for the data source
func (d *githubRestConnectionDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: util.GithubRestConnDataSourceDescription,
		Attributes:  connectionsutil.MergeDataSourceAttributes(BaseConnectorDataSourceSchema(), GithubRestConnectorsDataSourceSchema()),
	}
}

func (d *githubRestConnectionDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (d *githubRestConnectionDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var state GithubRestConnectionDataSourceModel

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
	reqParams := openapi.GetConnectionDetailsRequest{}

	// Set filters based on provided parameters
	if !state.ConnectionName.IsNull() && state.ConnectionName.ValueString() != "" {
		reqParams.SetConnectionname(state.ConnectionName.ValueString())
	}
	if !state.ConnectionKey.IsNull() {
		connectionKeyInt := state.ConnectionKey.ValueInt64()
		reqParams.SetConnectionkey(strconv.FormatInt(connectionKeyInt, 10))
	}
	apiReq := apiClient.ConnectionsAPI.GetConnectionDetails(ctx).GetConnectionDetailsRequest(reqParams)

	// Execute API request
	apiResp, httpResp, err := apiReq.Execute()
	if err != nil {
		if httpResp != nil && httpResp.StatusCode != 200 {
			log.Printf("[ERROR] HTTP error while creating GithubRest Connector: %s", httpResp.Status)
			var fetchResp map[string]interface{}
			if err := json.NewDecoder(httpResp.Body).Decode(&fetchResp); err != nil {
				resp.Diagnostics.AddError("Failed to decode error response", err.Error())
				return
			}
			resp.Diagnostics.AddError(
				"HTTP Error",
				fmt.Sprintf("HTTP error while creating GithubRest Connector for the reasons: %s", fetchResp["msg"]),
			)

		} else {
			log.Printf("[ERROR] API Call Failed: %v", err)
			resp.Diagnostics.AddError("API Call Failed", fmt.Sprintf("Error: %v", err))
		}
		return
	}
	// if apiResp != nil && *apiResp.GithubRESTConnectionResponse.Errorcode != 0 {
	// 	log.Printf("[ERROR]: Error in reading Github Rest connection. Errorcode: %v, Message: %v", *apiResp.GithubRESTConnectionResponse.Errorcode, *apiResp.GithubRESTConnectionResponse.Msg)
	// 	resp.Diagnostics.AddError("Read Github Rest connection failed", *apiResp.GithubRESTConnectionResponse.Msg)
	// 	return
	// }

	if apiResp != nil && apiResp.GithubRESTConnectionResponse == nil {
		error := "Verify the connection type"
		log.Printf("[ERROR]: Verify the connection type given")
		resp.Diagnostics.AddError("Read of Github_Rest connection failed", error)
		return
	}
	log.Printf("[DEBUG] HTTP Status Code: %d", httpResp.StatusCode)

	state.Msg = util.SafeStringDatasource(apiResp.GithubRESTConnectionResponse.Msg)
	state.ErrorCode = util.SafeInt64(apiResp.GithubRESTConnectionResponse.Errorcode)
	state.ConnectionName = util.SafeStringDatasource(apiResp.GithubRESTConnectionResponse.Connectionname)
	state.ConnectionKey = util.SafeInt64(apiResp.GithubRESTConnectionResponse.Connectionkey)
	state.Description = util.SafeStringDatasource(apiResp.GithubRESTConnectionResponse.Description)
	state.DefaultSavRoles = util.SafeStringDatasource(apiResp.GithubRESTConnectionResponse.Defaultsavroles)
	state.ConnectionType = util.SafeStringDatasource(apiResp.GithubRESTConnectionResponse.Connectiontype)
	state.CreatedOn = util.SafeStringDatasource(apiResp.GithubRESTConnectionResponse.Createdon)
	state.CreatedBy = util.SafeStringDatasource(apiResp.GithubRESTConnectionResponse.Createdby)
	state.UpdatedBy = util.SafeStringDatasource(apiResp.GithubRESTConnectionResponse.Updatedby)
	state.EmailTemplate = util.SafeStringDatasource(apiResp.GithubRESTConnectionResponse.Emailtemplate)

	if apiResp.GithubRESTConnectionResponse.Connectionattributes != nil {

		state.ConnectionAttributes = &GithubRestConnectionAttributes{
			IsTimeoutSupported: util.SafeBoolDatasource(apiResp.GithubRESTConnectionResponse.Connectionattributes.IsTimeoutSupported),
			// ConnectionJSON:       util.SafeStringDatasource(apiResp.GithubRESTConnectionResponse.Connectionattributes.ConnectionJSON),
			OrganizationList:     util.SafeStringDatasource(apiResp.GithubRESTConnectionResponse.Connectionattributes.ORGANIZATION_LIST),
			ImportAccountEntJSON: util.SafeStringDatasource(apiResp.GithubRESTConnectionResponse.Connectionattributes.ImportAccountEntJSON),
			// StatusThresholdConfig:    util.SafeStringDatasource(apiResp.GithubRESTConnectionResponse.Connectionattributes.STATUS_THRESHOLD_CONFIG),
			// AccessTokens:             util.SafeStringDatasource(apiResp.GithubRESTConnectionResponse.Connectionattributes.ACCESS_TOKENS),
			ConnectionType:           util.SafeStringDatasource(apiResp.GithubRESTConnectionResponse.Connectionattributes.ConnectionType),
			IsTimeoutConfigValidated: util.SafeBoolDatasource(apiResp.GithubRESTConnectionResponse.Connectionattributes.IsTimeoutConfigValidated),
		}
		// if apiResp.GithubRESTConnectionResponse.Connectionattributes.ConnectionTimeoutConfig != nil {
		// 	state.ConnectionAttributes.ConnectionTimeoutConfig = &ConnectionTimeoutConfig{
		// 		RetryWait:               util.SafeInt64(apiResp.GithubRESTConnectionResponse.Connectionattributes.ConnectionTimeoutConfig.RetryWait),
		// 		TokenRefreshMaxTryCount: util.SafeInt64(apiResp.GithubRESTConnectionResponse.Connectionattributes.ConnectionTimeoutConfig.TokenRefreshMaxTryCount),
		// 		RetryFailureStatusCode:  util.SafeInt64(apiResp.GithubRESTConnectionResponse.Connectionattributes.ConnectionTimeoutConfig.RetryFailureStatusCode),
		// 		RetryWaitMaxValue:       util.SafeInt64(apiResp.GithubRESTConnectionResponse.Connectionattributes.ConnectionTimeoutConfig.RetryWaitMaxValue),
		// 		RetryCount:              util.SafeInt64(apiResp.GithubRESTConnectionResponse.Connectionattributes.ConnectionTimeoutConfig.RetryCount),
		// 		ReadTimeout:             util.SafeInt64(apiResp.GithubRESTConnectionResponse.Connectionattributes.ConnectionTimeoutConfig.ReadTimeout),
		// 		ConnectionTimeout:       util.SafeInt64(apiResp.GithubRESTConnectionResponse.Connectionattributes.ConnectionTimeoutConfig.ConnectionTimeout),
		// 	}
		// }
	}

	if apiResp.GithubRESTConnectionResponse.Connectionattributes == nil {
		state.ConnectionAttributes = nil
	}
	if !state.Authenticate.IsNull() && !state.Authenticate.IsUnknown() {
		if state.Authenticate.ValueBool() {
			resp.Diagnostics.AddWarning(
				"Authentication Enabled",
				"`authenticate` is true; all connection_attributes will be returned in state.",
			)
		} else {
			resp.Diagnostics.AddWarning(
				"Authentication Disabled",
				"`authenticate` is false; connection_attributes will be removed from state.",
			)
			state.ConnectionAttributes = nil
		}
	}
	stateDiagnostics := resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(stateDiagnostics...)
}
