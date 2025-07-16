// Copyright (c) Saviynt Inc.
// SPDX-License-Identifier: MPL-2.0

// saviynt_security_system_resource manages security systems in the Saviynt Security Manager.
// The resource implements the full Terraform lifecycle:
//   - Create: provisions a new security system using the supplied configuration.
//   - Read: fetches the current security system state from Saviynt to keep Terraform’s state in sync.
//   - Update: applies any configuration changes to an existing security system.
//   - Import: brings an existing security system under Terraform management by its name.
package provider

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"terraform-provider-Saviynt/util"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	s "github.com/saviynt/saviynt-api-go-client"
	openapi "github.com/saviynt/saviynt-api-go-client/securitysystems"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &securitySystemResource{}
var _ resource.ResourceWithImportState = &securitySystemResource{}

// securitySystemResourceModel defines the state for our security system resource.
type SecuritySystemResourceModel struct {
	ID                                 types.String `tfsdk:"id"`
	Systemname                         types.String `tfsdk:"systemname"`
	DisplayName                        types.String `tfsdk:"display_name"`
	Hostname                           types.String `tfsdk:"hostname"`
	Port                               types.String `tfsdk:"port"`
	AccessAddWorkflow                  types.String `tfsdk:"access_add_workflow"`
	AccessRemoveWorkflow               types.String `tfsdk:"access_remove_workflow"`
	AddServiceAccountWorkflow          types.String `tfsdk:"add_service_account_workflow"`
	RemoveServiceAccountWorkflow       types.String `tfsdk:"remove_service_account_workflow"`
	Connectionparameters               types.String `tfsdk:"connection_parameters"`
	AutomatedProvisioning              types.String `tfsdk:"automated_provisioning"`
	UseOpenConnector                   types.Bool   `tfsdk:"use_open_connector"`
	ReconApplication                   types.String `tfsdk:"recon_application"`
	InstantProvision                   types.String `tfsdk:"instant_provision"`
	ProvisioningTries                  types.String `tfsdk:"provisioning_tries"`
	Provisioningcomments               types.String `tfsdk:"provisioning_comments"`
	ProposedAccountOwnersWorkflow      types.String `tfsdk:"proposed_account_owners_workflow"`
	FirefighterIDWorkflow              types.String `tfsdk:"firefighterid_workflow"`
	FirefighterIDRequestAccessWorkflow types.String `tfsdk:"firefighterid_request_access_workflow"`
	PolicyRule                         types.String `tfsdk:"policy_rule"`
	PolicyRuleServiceAccount           types.String `tfsdk:"policy_rule_service_account"`
	Connectionname                     types.String `tfsdk:"connectionname"`
	ProvisioningConnection             types.String `tfsdk:"provisioning_connection"`
	ServiceDeskConnection              types.String `tfsdk:"service_desk_connection"`
	ExternalRiskConnectionJson         types.String `tfsdk:"external_risk_connection_json"`
	InherentSODReportFields            types.Set    `tfsdk:"inherent_sod_report_fields"`
	Msg                                types.String `tfsdk:"msg"`
	ErrorCode                          types.String `tfsdk:"error_code"`
}

type securitySystemResource struct {
	client *s.Client
	token  string
}

func NewSecuritySystemResource() resource.Resource {
	return &securitySystemResource{}
}

func (r *securitySystemResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = "saviynt_security_system_resource"
}

func (r *securitySystemResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: util.SecuritySystemDescription,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Sensitive:   true,
				Description: "The unique ID of the resource.",
			},
			"systemname": schema.StringAttribute{
				Required:    true,
				Description: "Name of the security system.",
			},
			"display_name": schema.StringAttribute{
				Required:    true,
				Description: "Specify a user-friendly display name that is shown on the the user interface.",
			},
			"hostname": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Security system for which you want to create an endpoint.",
			},
			"port": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Description for the endpoint.",
			},
			"access_add_workflow": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Specify the workflow to be used for approvals for an access request, which can be for an account, entitlements, role, and so on.",
			},
			"access_remove_workflow": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Specify the workflow to be used when access has to be revoked, which can be for an account, entitlement, or any other de-provisioning task.",
			},
			"add_service_account_workflow": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Workflow for adding a service account.",
			},
			"remove_service_account_workflow": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Workflow for removing a service account.",
			},
			"proposed_account_owners_workflow": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Query to filter the access and display of the endpoint to specific users. If you do not define a query, the endpoint is displayed for all users",
			},
			"firefighterid_workflow": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Firefighter ID Workflow.",
			},
			"firefighterid_request_access_workflow": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Firefighter ID Request Access Workflow.",
			},
			"connection_parameters": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Query to filter the access and display of the endpoint to specific users. If you do not define a query, the endpoint is displayed for all users",
			},
			"automated_provisioning": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Specify true to enable automated provisioning.",
			},
			"provisioning_tries": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Specify the number of tries to be used for performing provisioning / de-provisioning to the third-party application. You can specify provisioningTries between 1 to 20 based on your requirement.",
			},
			"connectionname": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Select the connection name for performing reconciliation of identity objects from third-party application.",
			},
			"provisioning_connection": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "You can use a separate connection to an endpoint where you are performing provisioning or deprovisioning. Based on your requirement, you can specify a separate connection where you want to perform provisioning and de-provisioning.",
			},
			"service_desk_connection": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Specify the Service Desk Connection used for integration with a ticketing system, which can be a disconnected system too.",
			},
			"provisioning_comments": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Specify relevant comments for performing provisioning.",
			},
			"policy_rule": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Use this setting to assign the password policy for the security system.",
			},
			"policy_rule_service_account": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Use this setting to assign the password policy which will be used to set the service account passwords for the security system.",
			},
			"use_open_connector": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Specify true to enable the connectivity with any system over the open-source connectors such as REST.",
			},
			"recon_application": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Specify true to import data from the endpoint associated to the security system.",
			},
			"instant_provision": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Use this flag to prevent users from raising duplicate requests for the same applications.",
			},
			"external_risk_connection_json": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Contains JSON configuration for external risk connections and is applicable only for a few connections like SAP.",
			},
			"inherent_sod_report_fields": schema.SetAttribute{
				Computed:    true,
				ElementType: types.StringType,
				Optional:    true,
				Description: "You can use this option used to filter out columns in SOD.",
			},
			"msg": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "A message indicating the outcome of the operation.",
			},
			"error_code": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "An error code where '0' signifies success and '1' signifies an unsuccessful operation.",
			},
		},
	}
}

func (r *securitySystemResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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
	r.client = prov.client
	r.token = prov.accessToken
}

func (r *securitySystemResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan SecuritySystemResourceModel
	// Extract plan from request
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	cfg := openapi.NewConfiguration()
	apiBaseURL := strings.TrimPrefix(strings.TrimPrefix(r.client.APIBaseURL(), "https://"), "http://")
	cfg.Host = apiBaseURL
	cfg.Scheme = "https"
	cfg.AddDefaultHeader("Authorization", "Bearer "+r.token)
	cfg.HTTPClient = http.DefaultClient
	apiClient := openapi.NewAPIClient(cfg)
	createReq := openapi.CreateSecuritySystemRequest{
		//required fields
		Systemname:  plan.Systemname.ValueString(),
		DisplayName: plan.DisplayName.ValueString(),
		//optional fields
		Hostname:                     util.StringPointerOrEmpty(plan.Hostname),
		Port:                         util.StringPointerOrEmpty(plan.Port),
		AccessAddWorkflow:            util.StringPointerOrEmpty(plan.AccessAddWorkflow),
		AccessRemoveWorkflow:         util.StringPointerOrEmpty(plan.AccessRemoveWorkflow),
		AddServiceAccountWorkflow:    util.StringPointerOrEmpty(plan.AddServiceAccountWorkflow),
		RemoveServiceAccountWorkflow: util.StringPointerOrEmpty(plan.RemoveServiceAccountWorkflow),
		Connectionparameters:         util.StringPointerOrEmpty(plan.Connectionparameters),
		AutomatedProvisioning:        util.StringPointerOrEmpty(plan.AutomatedProvisioning),
		Useopenconnector:             util.BoolPointerOrEmpty(plan.UseOpenConnector),
		ReconApplication:             util.StringPointerOrEmpty(plan.ReconApplication),
		Instantprovision:             util.StringPointerOrEmpty(plan.InstantProvision),
		ProvisioningTries:            util.StringPointerOrEmpty(plan.ProvisioningTries),
		Provisioningcomments:         util.StringPointerOrEmpty(plan.Provisioningcomments),
	}
	// Execute the API call.
	apiResp, _, err := apiClient.SecuritySystemsAPI.CreateSecuritySystem(ctx).CreateSecuritySystemRequest(createReq).Execute()

	if err != nil {
		log.Printf("[ERROR] Failed to create API resource. Error: %v", err)
		resp.Diagnostics.AddError("API Create Failed", fmt.Sprintf("Error: %v", err))
		return
	}
	if apiResp != nil && *apiResp.ErrorCode != "0" {
		log.Printf("[ERROR]: Error in creating Security system resource. Errorcode: %v, Message: %v", *apiResp.ErrorCode, *apiResp.Msg)
		resp.Diagnostics.AddError("Creating of Security System resource failed", *apiResp.Msg)
		return
	}
	log.Printf("[INFO] Security system resource created successfully. Response: %v", createReq)
	// if apiResp!=nil && *apiResp.ErrorCode == "1" && strings.Contains(strings.ToLower(util.SafeDeref(apiResp.Msg)), "already exists") {
	// 	message := fmt.Sprintf("Security System %s already exists. Import the existing resource into Terraform state.", plan.Systemname.ValueString())
	// 	resp.Diagnostics.AddError(
	// 		"Security System Already Exists",
	// 		message,
	// 	)
	// 	return
	// }
	updateReq := openapi.UpdateSecuritySystemRequest{
		//required fields
		Systemname:  plan.Systemname.ValueString(),
		DisplayName: plan.DisplayName.ValueString(),
		//optional fields
		Hostname:                           util.StringPointerOrEmpty(plan.Hostname),
		Port:                               util.StringPointerOrEmpty(plan.Port),
		AccessAddWorkflow:                  util.StringPointerOrEmpty(plan.AccessAddWorkflow),
		AccessRemoveWorkflow:               util.StringPointerOrEmpty(plan.AccessRemoveWorkflow),
		AddServiceAccountWorkflow:          util.StringPointerOrEmpty(plan.AddServiceAccountWorkflow),
		RemoveServiceAccountWorkflow:       util.StringPointerOrEmpty(plan.RemoveServiceAccountWorkflow),
		Connectionparameters:               util.StringPointerOrEmpty(plan.Connectionparameters),
		AutomatedProvisioning:              util.StringPointerOrEmpty(plan.AutomatedProvisioning),
		Useopenconnector:                   util.BoolPointerOrEmpty(plan.UseOpenConnector),
		ReconApplication:                   util.StringPointerOrEmpty(plan.ReconApplication),
		Instantprovision:                   util.StringPointerOrEmpty(plan.InstantProvision),
		ProvisioningTries:                  util.StringPointerOrEmpty(plan.ProvisioningTries),
		Provisioningcomments:               util.StringPointerOrEmpty(plan.Provisioningcomments),
		ProposedAccountOwnersworkflow:      util.StringPointerOrEmpty(plan.ProposedAccountOwnersWorkflow),
		FirefighteridWorkflow:              util.StringPointerOrEmpty(plan.FirefighterIDWorkflow),
		FirefighteridRequestAccessWorkflow: util.StringPointerOrEmpty(plan.FirefighterIDRequestAccessWorkflow),
		PolicyRule:                         util.StringPointerOrEmpty(plan.PolicyRule),
		PolicyRuleServiceAccount:           util.StringPointerOrEmpty(plan.PolicyRuleServiceAccount),
		Connectionname:                     util.StringPointerOrEmpty(plan.Connectionname),
		ProvisioningConnection:             util.StringPointerOrEmpty(plan.ProvisioningConnection),
		ServiceDeskConnection:              util.StringPointerOrEmpty(plan.ServiceDeskConnection),
		ExternalRiskConnectionJson:         util.StringPointerOrEmpty(plan.ExternalRiskConnectionJson),
		InherentSODReportFields:            util.StringsFromSet(plan.InherentSODReportFields),
	}
	// Execute the update API call.
	_, _, _ = apiClient.SecuritySystemsAPI.UpdateSecuritySystem(ctx).UpdateSecuritySystemRequest(updateReq).Execute()
	log.Printf("[INFO] Security system resource updated successfully. Response: %v", updateReq)

	// Set the resource ID and store the API response in state.
	plan.ID = types.StringValue("security-system-" + plan.Systemname.ValueString())
	if plan.UseOpenConnector.IsNull() || plan.UseOpenConnector.IsUnknown() {
		plan.UseOpenConnector = types.BoolValue(false)
	}

	if plan.ReconApplication.IsNull() || plan.ReconApplication.IsUnknown() || plan.ReconApplication.ValueString() == "" {
		plan.ReconApplication = types.StringValue("true")
	}

	if plan.InstantProvision.IsNull() || plan.InstantProvision.IsUnknown() || plan.InstantProvision.ValueString() == "" {
		plan.InstantProvision = types.StringValue("false")
	}

	plan.Hostname = util.SafeString(plan.Hostname.ValueStringPointer())
	plan.Port = util.SafeString(plan.Port.ValueStringPointer())
	plan.ProvisioningTries = util.SafeString(plan.ProvisioningTries.ValueStringPointer())
	plan.Connectionparameters = util.SafeString(plan.Connectionparameters.ValueStringPointer())
	plan.AccessAddWorkflow = util.SafeString(plan.AccessAddWorkflow.ValueStringPointer())
	plan.Provisioningcomments = util.SafeString(plan.Provisioningcomments.ValueStringPointer())
	plan.AccessRemoveWorkflow = util.SafeString(plan.AccessRemoveWorkflow.ValueStringPointer())
	plan.AddServiceAccountWorkflow = util.SafeString(plan.AddServiceAccountWorkflow.ValueStringPointer())
	plan.RemoveServiceAccountWorkflow = util.SafeString(plan.RemoveServiceAccountWorkflow.ValueStringPointer())
	plan.ProposedAccountOwnersWorkflow = util.SafeString(plan.ProposedAccountOwnersWorkflow.ValueStringPointer())
	plan.AutomatedProvisioning = util.SafeString(plan.AutomatedProvisioning.ValueStringPointer())
	plan.FirefighterIDWorkflow = util.SafeString(plan.FirefighterIDWorkflow.ValueStringPointer())
	plan.FirefighterIDRequestAccessWorkflow = util.SafeString(plan.FirefighterIDRequestAccessWorkflow.ValueStringPointer())
	plan.Connectionname = util.SafeString(plan.Connectionname.ValueStringPointer())
	plan.ProvisioningConnection = util.SafeString(plan.ProvisioningConnection.ValueStringPointer())
	plan.ServiceDeskConnection = util.SafeString(plan.ServiceDeskConnection.ValueStringPointer())
	plan.PolicyRule = util.SafeString(plan.PolicyRule.ValueStringPointer())
	plan.PolicyRuleServiceAccount = util.SafeString(plan.PolicyRuleServiceAccount.ValueStringPointer())
	plan.ExternalRiskConnectionJson = util.SafeString(plan.ExternalRiskConnectionJson.ValueStringPointer())
	plan.InherentSODReportFields = util.NormalizeTFSetString(plan.InherentSODReportFields)
	plan.Msg = types.StringValue(util.SafeDeref(apiResp.Msg))
	plan.ErrorCode = types.StringValue(util.SafeDeref(apiResp.ErrorCode))
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *securitySystemResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state SecuritySystemResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Configure API client
	cfg := openapi.NewConfiguration()
	apiBaseURL := strings.TrimPrefix(strings.TrimPrefix(r.client.APIBaseURL(), "https://"), "http://")
	cfg.Host = apiBaseURL
	cfg.Scheme = "https"
	cfg.AddDefaultHeader("Authorization", "Bearer "+r.token)
	cfg.HTTPClient = http.DefaultClient

	apiClient := openapi.NewAPIClient(cfg)
	apiResp, _, err := apiClient.SecuritySystemsAPI.GetSecuritySystems(ctx).Systemname(state.Systemname.ValueString()).Execute()
	if err != nil {
		log.Printf("Problem with the get function in read block")
		resp.Diagnostics.AddError("API Read Failed", fmt.Sprintf("Error: %v", err))
		return
	}
	if apiResp != nil && *apiResp.ErrorCode != "0" {
		log.Printf("[ERROR]: Error in reading Security system resource. Errorcode: %v, Message: %v", *apiResp.ErrorCode, *apiResp.Msg)
		resp.Diagnostics.AddError("Reading of Security System resource failed", *apiResp.Msg)
		return
	}
	state.ID = types.StringValue("security-system-" + *apiResp.SecuritySystemDetails[0].Systemname)
	state.DisplayName = types.StringValue(util.SafeDeref(apiResp.SecuritySystemDetails[0].DisplayName))
	state.Hostname = util.SafeString(apiResp.SecuritySystemDetails[0].Hostname)
	state.Port = util.SafeString(apiResp.SecuritySystemDetails[0].Port)
	state.AccessAddWorkflow = util.SafeString(apiResp.SecuritySystemDetails[0].AccessAddWorkflow)
	state.AccessRemoveWorkflow = util.SafeString(apiResp.SecuritySystemDetails[0].AccessRemoveWorkflow)
	state.AddServiceAccountWorkflow = util.SafeString(apiResp.SecuritySystemDetails[0].AddServiceAccountWorkflow)
	state.RemoveServiceAccountWorkflow = util.SafeString(apiResp.SecuritySystemDetails[0].RemoveServiceAccountWorkflow)
	state.Connectionparameters = util.SafeString(apiResp.SecuritySystemDetails[0].Connectionparameters)
	state.AutomatedProvisioning = util.SafeString(apiResp.SecuritySystemDetails[0].AutomatedProvisioning)
	if apiResp.SecuritySystemDetails[0].Useopenconnector == nil {
		state.UseOpenConnector = types.BoolNull()
	} else {
		if *apiResp.SecuritySystemDetails[0].Useopenconnector == "true" {
			state.UseOpenConnector = types.BoolValue(true)
		} else {
			state.UseOpenConnector = types.BoolValue(false)
		}
	}
	state.ReconApplication = util.SafeString(apiResp.SecuritySystemDetails[0].ReconApplication)
	state.InstantProvision = util.SafeString(apiResp.SecuritySystemDetails[0].Instantprovision)
	state.ProvisioningTries = util.SafeString(apiResp.SecuritySystemDetails[0].ProvisioningTries)
	state.Provisioningcomments = util.SafeString(apiResp.SecuritySystemDetails[0].Provisioningcomments)
	state.ProposedAccountOwnersWorkflow = util.SafeString(apiResp.SecuritySystemDetails[0].ProposedAccountOwnersworkflow)
	state.FirefighterIDWorkflow = util.SafeString(apiResp.SecuritySystemDetails[0].FirefighteridWorkflow)
	state.FirefighterIDRequestAccessWorkflow = util.SafeString(apiResp.SecuritySystemDetails[0].FirefighteridRequestAccessWorkflow)
	state.PolicyRule = util.SafeString(apiResp.SecuritySystemDetails[0].PolicyRule)
	state.PolicyRuleServiceAccount = util.SafeString(apiResp.SecuritySystemDetails[0].PolicyRuleServiceAccount)
	state.Connectionname = util.SafeString(apiResp.SecuritySystemDetails[0].Connection)
	state.ProvisioningConnection = util.SafeString(apiResp.SecuritySystemDetails[0].ProvisioningConnection)
	state.ServiceDeskConnection = util.SafeString(apiResp.SecuritySystemDetails[0].ServiceDeskConnection)
	state.ExternalRiskConnectionJson = util.SafeString(apiResp.SecuritySystemDetails[0].ExternalRiskConnectionJson)
	state.InherentSODReportFields = util.StringsToSet(apiResp.SecuritySystemDetails[0].InherentSODReportFields)
	state.Msg = types.StringValue(util.SafeDeref(apiResp.Msg))
	state.ErrorCode = types.StringValue(util.SafeDeref(apiResp.ErrorCode))
	stateDiagnostics := resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(stateDiagnostics...)
	if resp.Diagnostics.HasError() {
		return
	}
}
func (r *securitySystemResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan SecuritySystemResourceModel
	var state SecuritySystemResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Extract the desired state from the request.
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	if plan.Systemname.ValueString() != state.Systemname.ValueString() {
		resp.Diagnostics.AddError("Error", "System name cannot be updated")
		log.Printf("[ERROR]: System name cannot be updated")
		return
	}

	// Initialize OpenAPI Client Configuration.
	cfg := openapi.NewConfiguration()
	apiBaseURL := strings.TrimPrefix(strings.TrimPrefix(r.client.APIBaseURL(), "https://"), "http://")
	cfg.Host = apiBaseURL
	cfg.Scheme = "https"
	cfg.AddDefaultHeader("Authorization", "Bearer "+r.token)
	cfg.HTTPClient = http.DefaultClient
	apiClient := openapi.NewAPIClient(cfg)
	updateReq := openapi.UpdateSecuritySystemRequest{
		//required fields
		Systemname:  plan.Systemname.ValueString(),
		DisplayName: plan.DisplayName.ValueString(),
		//optional fields
		Hostname:                           util.StringPointerOrEmpty(plan.Hostname),
		Port:                               util.StringPointerOrEmpty(plan.Port),
		AccessAddWorkflow:                  util.StringPointerOrEmpty(plan.AccessAddWorkflow),
		AccessRemoveWorkflow:               util.StringPointerOrEmpty(plan.AccessRemoveWorkflow),
		AddServiceAccountWorkflow:          util.StringPointerOrEmpty(plan.AddServiceAccountWorkflow),
		RemoveServiceAccountWorkflow:       util.StringPointerOrEmpty(plan.RemoveServiceAccountWorkflow),
		Connectionparameters:               util.StringPointerOrEmpty(plan.Connectionparameters),
		AutomatedProvisioning:              util.StringPointerOrEmpty(plan.AutomatedProvisioning),
		Useopenconnector:                   util.BoolPointerOrEmpty(plan.UseOpenConnector),
		ReconApplication:                   util.StringPointerOrEmpty(plan.ReconApplication),
		Instantprovision:                   util.StringPointerOrEmpty(plan.InstantProvision),
		ProvisioningTries:                  util.StringPointerOrEmpty(plan.ProvisioningTries),
		Provisioningcomments:               util.StringPointerOrEmpty(plan.Provisioningcomments),
		ProposedAccountOwnersworkflow:      util.StringPointerOrEmpty(plan.ProposedAccountOwnersWorkflow),
		FirefighteridWorkflow:              util.StringPointerOrEmpty(plan.FirefighterIDWorkflow),
		FirefighteridRequestAccessWorkflow: util.StringPointerOrEmpty(plan.FirefighterIDRequestAccessWorkflow),
		PolicyRule:                         util.StringPointerOrEmpty(plan.PolicyRule),
		PolicyRuleServiceAccount:           util.StringPointerOrEmpty(plan.PolicyRuleServiceAccount),
		Connectionname:                     util.StringPointerOrEmpty(plan.Connectionname),
		ProvisioningConnection:             util.StringPointerOrEmpty(plan.ProvisioningConnection),
		ServiceDeskConnection:              util.StringPointerOrEmpty(plan.ServiceDeskConnection),
		ExternalRiskConnectionJson:         util.StringPointerOrEmpty(plan.ExternalRiskConnectionJson),
		InherentSODReportFields:            util.StringsFromSet(plan.InherentSODReportFields),
	}
	// Execute the update API call.
	apiResp, _, err := apiClient.SecuritySystemsAPI.UpdateSecuritySystem(ctx).UpdateSecuritySystemRequest(updateReq).Execute()
	log.Printf("[INFO] Security system resource update request: %v", updateReq)
	if err != nil {
		log.Printf("Problem with the update function")
		resp.Diagnostics.AddError("API Update Failed", fmt.Sprintf("Error: %v", err))
		return
	}
	if apiResp != nil && *apiResp.ErrorCode != "0" {
		log.Printf("[ERROR]: Error in updating Security system resource. Errorcode: %v, Message: %v", *apiResp.ErrorCode, *apiResp.Msg)
		resp.Diagnostics.AddError("Updation of Security System resource failed", *apiResp.Msg)
		return
	}

	getResp, _, err := apiClient.SecuritySystemsAPI.GetSecuritySystems(ctx).Systemname(plan.Systemname.ValueString()).Execute()
	if err != nil {
		log.Printf("Problem with the get function in update block")
		resp.Diagnostics.AddError("API Read Failed", fmt.Sprintf("Error: %v", err))
		return
	}
	if getResp != nil && *getResp.ErrorCode != "0" {
		log.Printf("[ERROR]: Error in reading Security system resource after updation. Errorcode: %v, Message: %v", *getResp.ErrorCode, *getResp.Msg)
		resp.Diagnostics.AddError("Reading of Security System resource after updation failed", *getResp.Msg)
		return
	}
	plan.ID = types.StringValue("security-system-" + *getResp.SecuritySystemDetails[0].Systemname)
	plan.DisplayName = types.StringValue(util.SafeDeref(getResp.SecuritySystemDetails[0].DisplayName))
	plan.Hostname = util.SafeString(getResp.SecuritySystemDetails[0].Hostname)
	plan.Port = util.SafeString(getResp.SecuritySystemDetails[0].Port)
	plan.AccessAddWorkflow = util.SafeString(getResp.SecuritySystemDetails[0].AccessAddWorkflow)
	plan.AccessRemoveWorkflow = util.SafeString(getResp.SecuritySystemDetails[0].AccessRemoveWorkflow)
	plan.AddServiceAccountWorkflow = util.SafeString(getResp.SecuritySystemDetails[0].AddServiceAccountWorkflow)
	plan.RemoveServiceAccountWorkflow = util.SafeString(getResp.SecuritySystemDetails[0].RemoveServiceAccountWorkflow)
	plan.Connectionparameters = util.SafeString(getResp.SecuritySystemDetails[0].Connectionparameters)
	plan.AutomatedProvisioning = util.SafeString(getResp.SecuritySystemDetails[0].AutomatedProvisioning)
	if getResp.SecuritySystemDetails[0].Useopenconnector == nil {
		plan.UseOpenConnector = types.BoolNull()
	} else {
		if *getResp.SecuritySystemDetails[0].Useopenconnector == "true" {
			plan.UseOpenConnector = types.BoolValue(true)
		} else {
			plan.UseOpenConnector = types.BoolValue(false)
		}
	}
	plan.ReconApplication = util.SafeString(getResp.SecuritySystemDetails[0].ReconApplication)
	plan.InstantProvision = util.SafeString(getResp.SecuritySystemDetails[0].Instantprovision)
	plan.ProvisioningTries = util.SafeString(getResp.SecuritySystemDetails[0].ProvisioningTries)
	plan.Provisioningcomments = util.SafeString(getResp.SecuritySystemDetails[0].Provisioningcomments)
	plan.ProposedAccountOwnersWorkflow = util.SafeString(getResp.SecuritySystemDetails[0].ProposedAccountOwnersworkflow)
	plan.FirefighterIDWorkflow = util.SafeString(getResp.SecuritySystemDetails[0].FirefighteridWorkflow)
	plan.FirefighterIDRequestAccessWorkflow = util.SafeString(getResp.SecuritySystemDetails[0].FirefighteridRequestAccessWorkflow)
	plan.PolicyRule = util.SafeString(getResp.SecuritySystemDetails[0].PolicyRule)
	plan.PolicyRuleServiceAccount = util.SafeString(getResp.SecuritySystemDetails[0].PolicyRuleServiceAccount)
	plan.Connectionname = util.SafeString(getResp.SecuritySystemDetails[0].Connection)
	plan.ProvisioningConnection = util.SafeString(getResp.SecuritySystemDetails[0].ProvisioningConnection)
	plan.ServiceDeskConnection = util.SafeString(getResp.SecuritySystemDetails[0].ServiceDeskConnection)
	plan.ExternalRiskConnectionJson = util.SafeString(getResp.SecuritySystemDetails[0].ExternalRiskConnectionJson)
	plan.InherentSODReportFields = util.StringsToSet(getResp.SecuritySystemDetails[0].InherentSODReportFields)
	plan.Msg = types.StringValue(util.SafeDeref(apiResp.Msg))
	plan.ErrorCode = types.StringValue(util.SafeDeref(apiResp.ErrorCode))
	stateUpdateDiagnostics := resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(stateUpdateDiagnostics...)
}

func (r *securitySystemResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// resp.State.RemoveResource(ctx)
	if os.Getenv("TF_ACC") == "1" {
		resp.State.RemoveResource(ctx)
		return
	}
	resp.Diagnostics.AddError(
		"Delete Not Supported",
		"Resource deletion is not supported by this provider. Please remove the resource manually if required, or contact your administrator.",
	)
}

func (r *securitySystemResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	// Retrieve import ID and save to id attribute
	resource.ImportStatePassthroughID(ctx, path.Root("systemname"), req, resp)
}
