// Copyright (c) 2025 Saviynt Inc.
// SPDX-License-Identifier: MPL-2.0

// saviynt_security_system_datasource retrieves security system details from the Saviynt Security Manager.
// The data source supports a single Read operation to look up an existing security system by name.
package provider

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"terraform-provider-Saviynt/util"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	s "github.com/saviynt/saviynt-api-go-client"
	openapi "github.com/saviynt/saviynt-api-go-client/securitysystems"
)

// SecuritySystemsDataSource defines the data source
type securitySystemsDataSource struct {
	client *s.Client
	token  string
}

type SecuritySystemsDataSourceModel struct {
	Systemname     types.String            `tfsdk:"systemname"`
	Max            types.Int64             `tfsdk:"max"`
	Offset         types.Int64             `tfsdk:"offset"`
	Connectionname types.String            `tfsdk:"connectionname"`
	ConnectionType types.String            `tfsdk:"connection_type"`
	Msg            types.String            `tfsdk:"msg"`
	DisplayCount   types.Int64             `tfsdk:"display_count"`
	ErrorCode      types.String            `tfsdk:"error_code"`
	TotalCount     types.Int64             `tfsdk:"total_count"`
	Results        []SecuritySystemDetails `tfsdk:"results"`
}

// SecuritySystemDetails represents a single security system details object.
type SecuritySystemDetails struct {
	Systemname                         types.String   `tfsdk:"systemname1"`
	ConnectionType                     types.String   `tfsdk:"connection_type_1"`
	DisplayName                        types.String   `tfsdk:"display_name"`
	Hostname                           types.String   `tfsdk:"hostname"`
	Port                               types.String   `tfsdk:"port"`
	AccessAddWorkflow                  types.String   `tfsdk:"access_add_workflow"`
	AccessRemoveWorkflow               types.String   `tfsdk:"access_remove_workflow"`
	AddServiceAccountWorkflow          types.String   `tfsdk:"add_service_account_workflow"`
	RemoveServiceAccountWorkflow       types.String   `tfsdk:"remove_service_account_workflow"`
	Connectionparameters               types.String   `tfsdk:"connection_parameters"`
	AutomatedProvisioning              types.String   `tfsdk:"automated_provisioning"`
	UseOpenConnector                   types.String   `tfsdk:"use_open_connector"`
	ManageEntity                       types.String   `tfsdk:"manage_entity"`
	PersistentData                     types.String   `tfsdk:"persistent_data"`
	DefaultSystem                      types.String   `tfsdk:"default_system"`
	ReconApplication                   types.String   `tfsdk:"recon_application"`
	InstantProvision                   types.String   `tfsdk:"instant_provision"`
	ProvisioningTries                  types.String   `tfsdk:"provisioning_tries"`
	ProvisioningComments               types.String   `tfsdk:"provisioning_comments"`
	ProposedAccountOwnersWorkflow      types.String   `tfsdk:"proposed_account_owners_workflow"`
	FirefighterIDWorkflow              types.String   `tfsdk:"firefighterid_workflow"`
	FirefighterIDRequestAccessWorkflow types.String   `tfsdk:"firefighterid_request_access_workflow"`
	PolicyRule                         types.String   `tfsdk:"policy_rule"`
	PolicyRuleServiceAccount           types.String   `tfsdk:"policy_rule_service_account"`
	Connectionname                     types.String   `tfsdk:"connectionname1"`
	ProvisioningConnection             types.String   `tfsdk:"provisioning_connection"`
	ServiceDeskConnection              types.String   `tfsdk:"service_desk_connection"`
	ExternalRiskConnectionJson         types.String   `tfsdk:"external_risk_connection_json"`
	Connection                         types.String   `tfsdk:"connection"`
	CreateDate                         types.String   `tfsdk:"create_date"`
	UpdateDate                         types.String   `tfsdk:"update_date"`
	Endpoints                          types.String   `tfsdk:"endpoints"`
	CreatedBy                          types.String   `tfsdk:"created_by"`
	UpdatedBy                          types.String   `tfsdk:"updated_by"`
	Status                             types.String   `tfsdk:"status"`
	CreatedFrom                        types.String   `tfsdk:"created_from"`
	InherentSodReportFields            []types.String `tfsdk:"inherent_sod_report_fields"`
}

var _ datasource.DataSource = &securitySystemsDataSource{}

func NewSecuritySystemsDataSource() datasource.DataSource {
	return &securitySystemsDataSource{}
}

func (d *securitySystemsDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "saviynt_security_system_datasource"
}

func (d *securitySystemsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: util.SecuritySystemDataSourceDescription,
		Attributes: map[string]schema.Attribute{
			"systemname": schema.StringAttribute{
				Optional:    true,
				Description: "Name of the security systeme.",
			},
			"max": schema.Int64Attribute{
				Optional:    true,
				Description: "Name for the security system that will be displayed in the user interface.",
			},
			"offset": schema.Int64Attribute{
				Optional:    true,
				Description: "Security system for which you want to create an endpoint.",
			},
			"connectionname": schema.StringAttribute{
				Optional:    true,
				Description: "Owner type of the endpoint. It could be User or Usergroup.",
			},
			"connection_type": schema.StringAttribute{
				Optional:    true,
				Description: "Owner of the endpoint. If ownerType is User, specify the username of the owner. If ownerType is Usergroup, sepecify the name of the User group.",
			},
			"msg": schema.StringAttribute{
				Computed:    true,
				Description: "A message indicating the outcome of the operation.",
			},
			"display_count": schema.Int64Attribute{
				Computed:    true,
				Description: "The number of items currently displayed (e.g., on the current page or view).",
			},
			"error_code": schema.StringAttribute{
				Computed:    true,
				Description: "An error code where '0' signifies success and '1' signifies an unsuccessful operation.",
			},
			"total_count": schema.Int64Attribute{
				Computed:    true,
				Description: "The total number of items available in the dataset, irrespective of the current display settings.",
			},
			"results": schema.ListNestedAttribute{
				Computed:    true,
				Description: "List of security systems retrieved",
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"display_name":                          schema.StringAttribute{Computed: true, Description: "Specify a user-friendly display name that is shown on the the user interface."},
						"hostname":                              schema.StringAttribute{Computed: true, Description: "Security system for which you want to create an endpoint."},
						"connection_type_1":                     schema.StringAttribute{Computed: true, Description: "Specify a connection type to view all connections in EIC for the connection type."},
						"systemname1":                           schema.StringAttribute{Computed: true, Description: "Specify the security system name."},
						"access_add_workflow":                   schema.StringAttribute{Computed: true, Description: "Specify the workflow used for approvals for an access request (account, entitlements, role, etc.)."},
						"access_remove_workflow":                schema.StringAttribute{Computed: true, Description: "Workflow used when revoking access from accounts, entitlements, or performing other de-provisioning tasks."},
						"add_service_account_workflow":          schema.StringAttribute{Computed: true, Description: "Workflow for adding a service account."},
						"remove_service_account_workflow":       schema.StringAttribute{Computed: true, Description: "Workflow for removing a service account."},
						"connection_parameters":                 schema.StringAttribute{Computed: true, Description: "Query or parameters to restrict endpoint access to specific users."},
						"automated_provisioning":                schema.StringAttribute{Computed: true, Description: "Enables automated provisioning if set to true."},
						"use_open_connector":                    schema.StringAttribute{Computed: true, Description: "Enables connectivity using open-source connectors such as REST if set to true."},
						"manage_entity":                         schema.StringAttribute{Computed: true, Description: "Indicates if entity management is enabled for the security system."},
						"persistent_data":                       schema.StringAttribute{Computed: true, Description: "Indicates whether persistent data storage is enabled for the security system."},
						"default_system":                        schema.StringAttribute{Computed: true, Description: "Sets this security system as the default system for account searches when set to true."},
						"recon_application":                     schema.StringAttribute{Computed: true, Description: "Enables importing data from endpoints associated with the security system."},
						"instant_provision":                     schema.StringAttribute{Computed: true, Description: "Prevents users from submitting duplicate provisioning requests if set to true."},
						"provisioning_tries":                    schema.StringAttribute{Computed: true, Description: "Number of attempts allowed for provisioning actions."},
						"provisioning_comments":                 schema.StringAttribute{Computed: true, Description: "Comments relevant to provisioning actions."},
						"proposed_account_owners_workflow":      schema.StringAttribute{Computed: true, Description: "Workflow for assigning proposed account owners."},
						"firefighterid_workflow":                schema.StringAttribute{Computed: true, Description: "Workflow for handling firefighter ID requests."},
						"firefighterid_request_access_workflow": schema.StringAttribute{Computed: true, Description: "Workflow for requesting access to firefighter IDs."},
						"policy_rule":                           schema.StringAttribute{Computed: true, Description: "Password policy assigned for the security system."},
						"policy_rule_service_account":           schema.StringAttribute{Computed: true, Description: "Password policy applied to service accounts."},
						"connectionname1":                       schema.StringAttribute{Computed: true, Description: "Name of connection used for reconciling identity objects from third-party applications."},
						"provisioning_connection":               schema.StringAttribute{Computed: true, Description: "Dedicated connection for provisioning and de-provisioning tasks."},
						"service_desk_connection":               schema.StringAttribute{Computed: true, Description: "Connection to service desk or ticketing system integration."},
						"external_risk_connection_json":         schema.StringAttribute{Computed: true, Description: "JSON configuration for external risk connections (e.g., SAP)."},
						"connection":                            schema.StringAttribute{Computed: true, Description: "Primary connection used by the security system."},
						"create_date":                           schema.StringAttribute{Computed: true, Description: "Timestamp indicating when the security system was created."},
						"update_date":                           schema.StringAttribute{Computed: true, Description: "Timestamp indicating the last update to the security system."},
						"endpoints":                             schema.StringAttribute{Computed: true, Description: "Endpoints associated with the security system."},
						"created_by":                            schema.StringAttribute{Computed: true, Description: "Identifier of the user who created the security system."},
						"updated_by":                            schema.StringAttribute{Computed: true, Description: "Identifier of the user who last updated the security system."},
						"status":                                schema.StringAttribute{Computed: true, Description: "Current status of the security system (e.g., enabled, disabled)."},
						"created_from":                          schema.StringAttribute{Computed: true, Description: "Origin or method through which the security system was created."},
						"port":                                  schema.StringAttribute{Computed: true, Description: "Port information or description for the endpoint."},
						"inherent_sod_report_fields": schema.ListAttribute{
							ElementType: types.StringType,
							Computed:    true,
							Description: "List of fields used in filtering Segregation of Duties (SOD) reports.",
						},
					},
				},
			},
		},
	}
}

// Retrieve user-defined filters from configuration.
func (d *securitySystemsDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

// Read fetches data from the API and converts it to Terraform state.
func (d *securitySystemsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var state SecuritySystemsDataSourceModel

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

	// Initialize API client.
	apiClient := openapi.NewAPIClient(cfg)

	// Use provided Max filter or default value.
	apiReq := apiClient.SecuritySystemsAPI.GetSecuritySystems(ctx)

	// Only set Systemname if a non-null value is provided.
	if !state.Systemname.IsNull() && state.Systemname.ValueString() != "" {
		apiReq = apiReq.Systemname(state.Systemname.ValueString())
	}

	// Only set Max if provided.
	if !state.Max.IsNull() {
		apiReq = apiReq.Max(int32(state.Max.ValueInt64()))
	}

	// Only set Offset if provided.
	if !state.Offset.IsNull() {
		apiReq = apiReq.Offset(int32(state.Offset.ValueInt64()))
	}

	// Only set Connectionname if provided.
	if !state.Connectionname.IsNull() && state.Connectionname.ValueString() != "" {
		apiReq = apiReq.Connectionname(state.Connectionname.ValueString())
	}

	// Only set ConnectionType if provided.
	if !state.ConnectionType.IsNull() && state.ConnectionType.ValueString() != "" {
		apiReq = apiReq.ConnectionType(state.ConnectionType.ValueString())
	}

	// Execute the API request.
	apiResp, httpResp, err := apiReq.Execute()
	if err != nil {
		if httpResp != nil && httpResp.StatusCode != 200 {
			log.Printf("[ERROR] HTTP error while creating Security System: %s", httpResp.Status)
			var fetchResp map[string]interface{}
			if err := json.NewDecoder(httpResp.Body).Decode(&fetchResp); err != nil {
				resp.Diagnostics.AddError("Failed to decode error response", err.Error())
				return
			}
			resp.Diagnostics.AddError(
				"HTTP Error",
				fmt.Sprintf("HTTP error while creating security system for the reasons: %s", fetchResp["msg"]),
			)

		} else {
			log.Printf("[ERROR] API Call Failed: %v", err)
			resp.Diagnostics.AddError("API Call Failed", fmt.Sprintf("Error: %v", err))
		}
		return
	}
	//if the response is null
	if len(apiResp.SecuritySystemDetails) == 0 {
		log.Println("[DEBUG] No data returned from API")
		resp.Diagnostics.AddError("API Response is empty", "No data found for the given filters in the environment.")
		return
	}

	log.Printf("[DEBUG] HTTP Status Code: %d", httpResp.StatusCode)

	// Transform API response to a slice of SecuritySystemDetails.
	state.Msg = types.StringValue(*apiResp.Msg)
	state.DisplayCount = types.Int64Value(int64(*apiResp.DisplayCount))
	state.ErrorCode = types.StringValue(*apiResp.ErrorCode)
	state.TotalCount = types.Int64Value(int64(*apiResp.TotalCount))

	if apiResp.SecuritySystemDetails != nil {
		for _, item := range apiResp.SecuritySystemDetails {
			securitySystemState := SecuritySystemDetails{
				ConnectionType:                     util.SafeString(item.ConnectionType),
				Connectionname:                     util.SafeString(item.Connectionname),
				DefaultSystem:                      util.SafeString(item.DefaultSystem),
				DisplayName:                        util.SafeString(item.DisplayName),
				Hostname:                           util.SafeString(item.Hostname),
				Systemname:                         util.SafeString(item.Systemname),
				AccessAddWorkflow:                  util.SafeString(item.AccessAddWorkflow),
				AccessRemoveWorkflow:               util.SafeString(item.AccessRemoveWorkflow),
				AddServiceAccountWorkflow:          util.SafeString(item.AddServiceAccountWorkflow),
				RemoveServiceAccountWorkflow:       util.SafeString(item.RemoveServiceAccountWorkflow),
				Connectionparameters:               util.SafeString(item.Connectionparameters),
				ProvisioningConnection:             util.SafeString(item.ProvisioningConnection),
				Connection:                         util.SafeString(item.Connection),
				CreateDate:                         util.SafeString(item.CreateDate),
				UpdateDate:                         util.SafeString(item.UpdateDate),
				Endpoints:                          util.SafeString(item.Endpoints),
				UseOpenConnector:                   util.SafeString(item.Useopenconnector),
				ReconApplication:                   util.SafeString(item.ReconApplication),
				AutomatedProvisioning:              util.SafeString(item.AutomatedProvisioning),
				InstantProvision:                   util.SafeString(item.Instantprovision),
				ProvisioningComments:               util.SafeString(item.Provisioningcomments),
				ProvisioningTries:                  util.SafeString(item.ProvisioningTries),
				ProposedAccountOwnersWorkflow:      util.SafeString(item.ProposedAccountOwnersworkflow),
				FirefighterIDWorkflow:              util.SafeString(item.FirefighteridWorkflow),
				FirefighterIDRequestAccessWorkflow: util.SafeString(item.FirefighteridRequestAccessWorkflow),
				PolicyRuleServiceAccount:           util.SafeString(item.PolicyRuleServiceAccount),
				ServiceDeskConnection:              util.SafeString(item.ServiceDeskConnection),
				ExternalRiskConnectionJson:         util.SafeString(item.ExternalRiskConnectionJson),
				CreatedBy:                          util.SafeString(item.CreatedBy),
				UpdatedBy:                          util.SafeString(item.UpdatedBy),
				Status:                             util.SafeString(item.Status),
				CreatedFrom:                        util.SafeString(item.CreatedFrom),
				PolicyRule:                         util.SafeString(item.PolicyRule),
				Port:                               util.SafeString(item.Port),
				InherentSodReportFields:            util.StringsToTypeStrings(item.InherentSODReportFields),
			}
			state.Results = append(state.Results, securitySystemState)
		}
	}

	stateDiagnostics := resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(stateDiagnostics...)

	if resp.Diagnostics.HasError() {
		return
	}
}
