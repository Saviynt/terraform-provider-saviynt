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

// saviynt_salesforce_connection_datasource retrieves salesforce connections details from the Saviynt Security Manager.
// The data source supports a single Read operation to look up an existing salesforce connections by name.
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

var _ datasource.DataSource = &salesforceConnectionDataSource{}

// SalesforceConnectionDataSource defines the data source
type salesforceConnectionDataSource struct {
	client *s.Client
	token  string
}

type SalesforceConnectionDataSourceModel struct {
	ID types.String `tfsdk:"id"`
	BaseConnectionDataSourceModel
	ConnectionAttributes *SalesforceConnectionAttributes `tfsdk:"connection_attributes"`
}

type SalesforceConnectionAttributes struct {
	IsTimeoutSupported       types.Bool               `tfsdk:"is_timeout_supported"`
	ObjectToBeImported       types.String             `tfsdk:"object_to_be_imported"`
	FeatureLicenseJson       types.String             `tfsdk:"feature_license_json"`
	CreateAccountJson        types.String             `tfsdk:"createaccountjson"`
	RedirectUri              types.String             `tfsdk:"redirect_uri"`
	ConnectionTimeoutConfig  *ConnectionTimeoutConfig `tfsdk:"connection_timeout_config"`
	ModifyAccountJson        types.String             `tfsdk:"modifyaccountjson"`
	ConnectionType           types.String             `tfsdk:"connection_type"`
	IsTimeoutConfigValidated types.Bool               `tfsdk:"is_timeout_config_validated"`
	PamConfig                types.String             `tfsdk:"pam_config"`
	CustomConfigJson         types.String             `tfsdk:"customconfigjson"`
	FieldMappingJson         types.String             `tfsdk:"field_mapping_json"`
	StatusThresholdConfig    types.String             `tfsdk:"status_threshold_config"`
	AccountFieldQuery        types.String             `tfsdk:"account_field_query"`
	CustomCreateAccountUrl   types.String             `tfsdk:"custom_createaccount_url"`
	AccountFilterQuery       types.String             `tfsdk:"account_filter_query"`
	InstanceUrl              types.String             `tfsdk:"instance_url"`
}

func NewSalesforceConnectionsDataSource() datasource.DataSource {
	return &salesforceConnectionDataSource{}
}

func (d *salesforceConnectionDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "saviynt_salesforce_connection_datasource"
}

func SalesforceConnectorsDataSourceSchema() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:    true,
			Description: "Resource ID.",
		},
		"connection_attributes": schema.SingleNestedAttribute{
			Computed: true,
			Attributes: map[string]schema.Attribute{
				"is_timeout_supported":        schema.BoolAttribute{Computed: true},
				"object_to_be_imported":       schema.StringAttribute{Computed: true},
				"feature_license_json":        schema.StringAttribute{Computed: true},
				"createaccountjson":           schema.StringAttribute{Computed: true},
				"redirect_uri":                schema.StringAttribute{Computed: true},
				"modifyaccountjson":           schema.StringAttribute{Computed: true},
				"connection_type":             schema.StringAttribute{Computed: true},
				"is_timeout_config_validated": schema.BoolAttribute{Computed: true},
				"pam_config":                  schema.StringAttribute{Computed: true},
				"customconfigjson":            schema.StringAttribute{Computed: true},
				"field_mapping_json":          schema.StringAttribute{Computed: true},
				"status_threshold_config":     schema.StringAttribute{Computed: true},
				"account_field_query":         schema.StringAttribute{Computed: true},
				"custom_createaccount_url":    schema.StringAttribute{Computed: true},
				"account_filter_query":        schema.StringAttribute{Computed: true},
				"instance_url":                schema.StringAttribute{Computed: true},
				"connection_timeout_config": schema.SingleNestedAttribute{
					Computed:   true,
					Attributes: ConnectionTimeoutConfigeSchema(),
				},
			},
		},
	}
}

// Schema defines the attributes for the data source
func (d *salesforceConnectionDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: util.SalesforceConnDataSourceDescription,
		Attributes:  connectionsutil.MergeDataSourceAttributes(BaseConnectorDataSourceSchema(), SalesforceConnectorsDataSourceSchema()),
	}
}

func (d *salesforceConnectionDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (d *salesforceConnectionDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var state SalesforceConnectionDataSourceModel

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
			log.Printf("[ERROR] HTTP error while creating Salesforce Connection: %s", httpResp.Status)
			var fetchResp map[string]interface{}
			if err := json.NewDecoder(httpResp.Body).Decode(&fetchResp); err != nil {
				resp.Diagnostics.AddError("Failed to decode error response", err.Error())
				return
			}
			resp.Diagnostics.AddError(
				"HTTP Error",
				fmt.Sprintf("HTTP error while creating Salesforce Connection for the reasons: %s", fetchResp["msg"]),
			)

		} else {
			log.Printf("[ERROR] API Call Failed: %v", err)
			resp.Diagnostics.AddError("API Call Failed", fmt.Sprintf("Error: %v", err))
		}
		return
	}

	if apiResp != nil && apiResp.SalesforceConnectionResponse == nil {
		error := "Verify the connection type"
		log.Printf("[ERROR]: Verify the connection type given")
		resp.Diagnostics.AddError("Read of Salesforce connection failed", error)
		return
	}
	log.Printf("[DEBUG] HTTP Status Code: %d", httpResp.StatusCode)

	state.Msg = util.SafeStringDatasource(apiResp.SalesforceConnectionResponse.Msg)
	state.ErrorCode = util.SafeInt64(apiResp.SalesforceConnectionResponse.Errorcode)
	state.ConnectionName = util.SafeStringDatasource(apiResp.SalesforceConnectionResponse.Connectionname)
	state.ConnectionKey = util.SafeInt64(apiResp.SalesforceConnectionResponse.Connectionkey)
	state.Description = util.SafeStringDatasource(apiResp.SalesforceConnectionResponse.Description)
	state.DefaultSavRoles = util.SafeStringDatasource(apiResp.SalesforceConnectionResponse.Defaultsavroles)
	state.ConnectionType = util.SafeStringDatasource(apiResp.SalesforceConnectionResponse.Connectiontype)
	state.CreatedOn = util.SafeStringDatasource(apiResp.SalesforceConnectionResponse.Createdon)
	state.CreatedBy = util.SafeStringDatasource(apiResp.SalesforceConnectionResponse.Createdby)
	state.UpdatedBy = util.SafeStringDatasource(apiResp.SalesforceConnectionResponse.Updatedby)
	state.EmailTemplate = util.SafeStringDatasource(apiResp.SalesforceConnectionResponse.Emailtemplate)

	if apiResp.SalesforceConnectionResponse.Connectionattributes != nil {
		state.ConnectionAttributes = &SalesforceConnectionAttributes{
			IsTimeoutSupported:       util.SafeBoolDatasource(apiResp.SalesforceConnectionResponse.Connectionattributes.IsTimeoutSupported),
			ObjectToBeImported:       util.SafeStringDatasource(apiResp.SalesforceConnectionResponse.Connectionattributes.OBJECT_TO_BE_IMPORTED),
			FeatureLicenseJson:       util.SafeStringDatasource(apiResp.SalesforceConnectionResponse.Connectionattributes.FEATURE_LICENSE_JSON),
			CreateAccountJson:        util.SafeStringDatasource(apiResp.SalesforceConnectionResponse.Connectionattributes.CREATEACCOUNTJSON),
			RedirectUri:              util.SafeStringDatasource(apiResp.SalesforceConnectionResponse.Connectionattributes.REDIRECT_URI),
			ConnectionType:           util.SafeStringDatasource(apiResp.SalesforceConnectionResponse.Connectionattributes.ConnectionType),
			ModifyAccountJson:        util.SafeStringDatasource(apiResp.SalesforceConnectionResponse.Connectionattributes.MODIFYACCOUNTJSON),
			IsTimeoutConfigValidated: util.SafeBoolDatasource(apiResp.SalesforceConnectionResponse.Connectionattributes.IsTimeoutConfigValidated),
			PamConfig:                util.SafeStringDatasource(apiResp.SalesforceConnectionResponse.Connectionattributes.PAM_CONFIG),
			CustomConfigJson:         util.SafeStringDatasource(apiResp.SalesforceConnectionResponse.Connectionattributes.CUSTOMCONFIGJSON),
			FieldMappingJson:         util.SafeStringDatasource(apiResp.SalesforceConnectionResponse.Connectionattributes.FIELD_MAPPING_JSON),
			StatusThresholdConfig:    util.SafeStringDatasource(apiResp.SalesforceConnectionResponse.Connectionattributes.STATUS_THRESHOLD_CONFIG),
			AccountFieldQuery:        util.SafeStringDatasource(apiResp.SalesforceConnectionResponse.Connectionattributes.ACCOUNT_FIELD_QUERY),
			CustomCreateAccountUrl:   util.SafeStringDatasource(apiResp.SalesforceConnectionResponse.Connectionattributes.CUSTOM_CREATEACCOUNT_URL),
			AccountFilterQuery:       util.SafeStringDatasource(apiResp.SalesforceConnectionResponse.Connectionattributes.ACCOUNT_FILTER_QUERY),
			InstanceUrl:              util.SafeStringDatasource(apiResp.SalesforceConnectionResponse.Connectionattributes.INSTANCE_URL),
		}

		if apiResp.SalesforceConnectionResponse.Connectionattributes.ConnectionTimeoutConfig != nil {
			state.ConnectionAttributes.ConnectionTimeoutConfig = &ConnectionTimeoutConfig{
				RetryWait:               util.SafeInt64(apiResp.SalesforceConnectionResponse.Connectionattributes.ConnectionTimeoutConfig.RetryWait),
				TokenRefreshMaxTryCount: util.SafeInt64(apiResp.SalesforceConnectionResponse.Connectionattributes.ConnectionTimeoutConfig.TokenRefreshMaxTryCount),
				RetryFailureStatusCode:  util.SafeInt64(apiResp.SalesforceConnectionResponse.Connectionattributes.ConnectionTimeoutConfig.RetryFailureStatusCode),
				RetryWaitMaxValue:       util.SafeInt64(apiResp.SalesforceConnectionResponse.Connectionattributes.ConnectionTimeoutConfig.RetryWaitMaxValue),
				RetryCount:              util.SafeInt64(apiResp.SalesforceConnectionResponse.Connectionattributes.ConnectionTimeoutConfig.RetryCount),
				ReadTimeout:             util.SafeInt64(apiResp.SalesforceConnectionResponse.Connectionattributes.ConnectionTimeoutConfig.ReadTimeout),
				ConnectionTimeout:       util.SafeInt64(apiResp.SalesforceConnectionResponse.Connectionattributes.ConnectionTimeoutConfig.ConnectionTimeout),
			}
		}
	}

	if apiResp.SalesforceConnectionResponse.Connectionattributes == nil {
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
