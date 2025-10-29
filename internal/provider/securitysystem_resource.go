// Copyright (c) 2025 Saviynt Inc.
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
	"os"

	"terraform-provider-Saviynt/internal/client"
	"terraform-provider-Saviynt/util"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	openapi "github.com/saviynt/saviynt-api-go-client/securitysystems"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &SecuritySystemResource{}
var _ resource.ResourceWithImportState = &SecuritySystemResource{}

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

type SecuritySystemResource struct {
	client                client.SaviyntClientInterface
	token                 string
	saviyntVersion        string
	provider              client.SaviyntProviderInterface
	securitySystemFactory client.SecuritySystemFactoryInterface
}

func NewSecuritySystemResource() resource.Resource {
	return &SecuritySystemResource{
		securitySystemFactory: &client.DefaultSecuritySystemFactory{},
	}
}

func NewSecuritySystemResourceWithFactory(factory client.SecuritySystemFactoryInterface) resource.Resource {
	return &SecuritySystemResource{
		securitySystemFactory: factory,
	}
}

func (r *SecuritySystemResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = "saviynt_security_system_resource"
}

func (r *SecuritySystemResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
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
				Computed:    true,
				Description: "A message indicating the outcome of the operation.",
			},
			"error_code": schema.StringAttribute{
				Computed:    true,
				Description: "An error code where '0' signifies success and '1' signifies an unsuccessful operation.",
			},
		},
	}
}

func (r *SecuritySystemResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Check if provider data is available.
	if req.ProviderData == nil {
		log.Println("ProviderData is nil, returning early.")
		return
	}

	// Cast provider data to your provider type.
	prov, ok := req.ProviderData.(*SaviyntProvider)
	if !ok {
		resp.Diagnostics.AddError("Unexpected Provider Data", "Expected *SaviyntProvider")
		return
	}

	// Set the client, token, and provider reference from the provider state
	r.client = &client.SaviyntClientWrapper{Client: prov.client}
	r.token = prov.accessToken
	r.saviyntVersion = prov.saviyntVersion
	r.provider = &client.SaviyntProviderWrapper{Provider: prov} // Store provider reference for retry logic
}

// SetClient sets the client for testing purposes
func (r *SecuritySystemResource) SetClient(client client.SaviyntClientInterface) {
	r.client = client
}

// SetToken sets the token for testing purposes
func (r *SecuritySystemResource) SetToken(token string) {
	r.token = token
}

// SetProvider sets the provider for testing purposes
func (r *SecuritySystemResource) SetProvider(provider client.SaviyntProviderInterface) {
	r.provider = provider
}

func (r *SecuritySystemResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan SecuritySystemResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Validate version-specific attributes for Security System
	util.ValidateAttributeCompatibility(r.saviyntVersion, "SecuritySystem", "instant_provisioning", plan.InstantProvision.ValueStringPointer(), &resp.Diagnostics)

	if resp.Diagnostics.HasError() {
		return
	}

	// Use interface pattern instead of direct API client creation
	apiResp, err := r.CreateSecuritySystem(ctx, &plan)
	if err != nil {
		resp.Diagnostics.AddError("API Create Failed", err.Error())
		return
	}

	// Update model from create response
	r.UpdateModelFromCreateResponse(&plan, apiResp)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *SecuritySystemResource) CreateSecuritySystem(ctx context.Context, plan *SecuritySystemResourceModel) (*openapi.CreateSecuritySystem200Response, error) {
	var existingResource *openapi.GetSecuritySystems200Response

	// Check if security system already exists (idempotency check) with retry logic
	err := r.provider.AuthenticatedAPICallWithRetry(ctx, "get_security_systems", func(token string) error {
		securitySystemOps := r.securitySystemFactory.CreateSecuritySystemOperations(r.client.APIBaseURL(), token)
		resp, httpResp, err := securitySystemOps.GetSecuritySystems(ctx, plan.Systemname.ValueString())
		if httpResp != nil && httpResp.StatusCode == 401 {
			return fmt.Errorf("401 unauthorized")
		}
		existingResource = resp
		return err
	})

	if err != nil {
		return nil, fmt.Errorf("failed to check existing security system: %w", err)
	}

	if existingResource != nil &&
		existingResource.SecuritySystemDetails != nil &&
		len(existingResource.SecuritySystemDetails) > 0 {
		log.Printf("[ERROR]: Security system with name '%s' already exists. Skipping creation.", plan.Systemname.ValueString())
		return nil, fmt.Errorf("security system with name '%s' already exists", plan.Systemname.ValueString())
	}

	// Build security system create request
	createReq := r.BuildCreateSecuritySystemRequest(plan)

	var apiResp *openapi.CreateSecuritySystem200Response

	// Execute create operation with retry logic
	err = r.provider.AuthenticatedAPICallWithRetry(ctx, "create_security_system", func(token string) error {
		securitySystemOps := r.securitySystemFactory.CreateSecuritySystemOperations(r.client.APIBaseURL(), token)
		resp, httpResp, err := securitySystemOps.CreateSecuritySystem(ctx, createReq)
		if httpResp != nil && httpResp.StatusCode == 401 {
			return fmt.Errorf("401 unauthorized")
		}
		apiResp = resp
		return err
	})

	if err != nil {
		return nil, fmt.Errorf("API call failed: %w", err)
	}

	if apiResp != nil && apiResp.ErrorCode != nil && *apiResp.ErrorCode != "0" {
		message := ""
		if apiResp.Msg != nil {
			message = *apiResp.Msg
		}
		log.Printf("[ERROR]: Error in creating Security system resource. Errorcode: %v, Message: %v", *apiResp.ErrorCode, message)
		return nil, fmt.Errorf("creating of Security System resource failed In CreateSecuritySystem Block: %s", message)
	}

	// Execute update operation for additional fields with retry logic
	updateReq := r.BuildUpdateSecuritySystemRequest(plan)
	var updateResp *openapi.CreateSecuritySystem200Response

	err = r.provider.AuthenticatedAPICallWithRetry(ctx, "update_security_system", func(token string) error {
		securitySystemOps := r.securitySystemFactory.CreateSecuritySystemOperations(r.client.APIBaseURL(), token)
		resp, httpResp, err := securitySystemOps.UpdateSecuritySystem(ctx, updateReq)
		if httpResp != nil && httpResp.StatusCode == 401 {
			return fmt.Errorf("401 unauthorized")
		}
		updateResp = resp
		return err
	})

	if err != nil {
		log.Printf("Problem with the creating function")
		return nil, fmt.Errorf("API update call failed In CreateSecuritySystem: %w", err)
	}

	if updateResp != nil && updateResp.ErrorCode != nil && *updateResp.ErrorCode != "0" {
		message := ""
		if updateResp.Msg != nil {
			message = *updateResp.Msg
		}
		log.Printf("[ERROR]: Error in creating Security system resource. Errorcode: %v, Message: %v", *updateResp.ErrorCode, message)
		return nil, fmt.Errorf("API update error In CreateSecuritySystem Block: %s", message)
	}

	return apiResp, nil
}

// BuildCreateSecuritySystemRequest - Extracted request building logic
func (r *SecuritySystemResource) BuildCreateSecuritySystemRequest(plan *SecuritySystemResourceModel) openapi.CreateSecuritySystemRequest {
	return openapi.CreateSecuritySystemRequest{
		// Required fields
		Systemname:  plan.Systemname.ValueString(),
		DisplayName: plan.DisplayName.ValueString(),
		// Optional fields
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
}

// BuildUpdateSecuritySystemRequest - Extracted update request building logic
func (r *SecuritySystemResource) BuildUpdateSecuritySystemRequest(plan *SecuritySystemResourceModel) openapi.UpdateSecuritySystemRequest {
	return openapi.UpdateSecuritySystemRequest{
		// Required fields
		Systemname:  plan.Systemname.ValueString(),
		DisplayName: plan.DisplayName.ValueString(),
		// Optional fields
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
}

// UpdateModelFromCreateResponse - Extracted state management logic
func (r *SecuritySystemResource) UpdateModelFromCreateResponse(plan *SecuritySystemResourceModel, apiResp *openapi.CreateSecuritySystem200Response) {
	// Set the resource ID
	plan.ID = types.StringValue("security-system-" + plan.Systemname.ValueString())

	// Set default values for computed fields
	if plan.UseOpenConnector.IsNull() || plan.UseOpenConnector.IsUnknown() {
		plan.UseOpenConnector = types.BoolValue(false)
	}

	if plan.ReconApplication.IsNull() || plan.ReconApplication.IsUnknown() || plan.ReconApplication.ValueString() == "" {
		plan.ReconApplication = types.StringValue("true")
	}

	if plan.InstantProvision.IsNull() || plan.InstantProvision.IsUnknown() || plan.InstantProvision.ValueString() == "" {
		plan.InstantProvision = types.StringValue("false")
	}

	// Update all optional fields to maintain state
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

	// Set API response fields
	plan.Msg = types.StringValue(util.SafeDeref(apiResp.Msg))
	plan.ErrorCode = types.StringValue(util.SafeDeref(apiResp.ErrorCode))
}

func (r *SecuritySystemResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state SecuritySystemResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Use interface pattern for read operation
	apiResp, err := r.ReadSecuritySystem(ctx, state.Systemname.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("API Read Failed", err.Error())
		return
	}

	// Update model from read response
	r.UpdateModelFromReadResponse(&state, apiResp)
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}
func (r *SecuritySystemResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state SecuritySystemResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Validate version-specific attributes for Security System
	util.ValidateAttributeCompatibility(r.saviyntVersion, "SecuritySystem", "instant_provisioning", plan.InstantProvision.ValueStringPointer(), &resp.Diagnostics)

	if resp.Diagnostics.HasError() {
		return
	}

	// Validate that system name cannot be changed
	if plan.Systemname.ValueString() != state.Systemname.ValueString() {
		resp.Diagnostics.AddError("Error", "System name cannot be updated")
		return
	}

	// Use interface pattern for update operation
	_, err := r.UpdateSecuritySystem(ctx, &plan)
	if err != nil {
		resp.Diagnostics.AddError("API Update Failed", err.Error())
		return
	}

	// Refresh state after update
	getResp, err := r.ReadSecuritySystem(ctx, plan.Systemname.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("API Read Failed After Update", err.Error())
		return
	}

	r.UpdateModelFromReadResponse(&plan, getResp)
	resp.Diagnostics.Append(resp.State.Set(ctx, plan)...)
}

func (r *SecuritySystemResource) UpdateSecuritySystem(ctx context.Context, plan *SecuritySystemResourceModel) (*openapi.CreateSecuritySystem200Response, error) {
	// Build update request
	updateReq := r.BuildUpdateSecuritySystemRequest(plan)
	var apiResp *openapi.CreateSecuritySystem200Response

	// Execute update operation with retry logic
	err := r.provider.AuthenticatedAPICallWithRetry(ctx, "update_security_system", func(token string) error {
		securitySystemOps := r.securitySystemFactory.CreateSecuritySystemOperations(r.client.APIBaseURL(), token)
		resp, httpResp, err := securitySystemOps.UpdateSecuritySystem(ctx, updateReq)
		if httpResp != nil && httpResp.StatusCode == 401 {
			return fmt.Errorf("401 unauthorized")
		}
		apiResp = resp
		return err
	})

	if err != nil {
		return nil, fmt.Errorf("API call failed: %w", err)
	}

	if apiResp != nil && apiResp.ErrorCode != nil && *apiResp.ErrorCode != "0" {
		message := ""
		if apiResp.Msg != nil {
			message = *apiResp.Msg
		}
		return nil, fmt.Errorf("API error: %s", message)
	}

	log.Printf("[INFO] Security system resource updated successfully. Response: %v", updateReq)

	return apiResp, nil
}

func (r *SecuritySystemResource) ReadSecuritySystem(ctx context.Context, systemname string) (*openapi.GetSecuritySystems200Response, error) {
	var apiResp *openapi.GetSecuritySystems200Response

	// Execute read operation with retry logic
	err := r.provider.AuthenticatedAPICallWithRetry(ctx, "read_security_system", func(token string) error {
		securitySystemOps := r.securitySystemFactory.CreateSecuritySystemOperations(r.client.APIBaseURL(), token)
		resp, httpResp, err := securitySystemOps.GetSecuritySystems(ctx, systemname)
		if httpResp != nil && httpResp.StatusCode == 401 {
			return fmt.Errorf("401 unauthorized")
		}
		apiResp = resp
		return err
	})

	if err != nil {
		log.Printf("[ERROR]: Problem with the get function in read block")
		return nil, fmt.Errorf("API Read Failed In ReadSecuritySystem Block: %w", err)
	}

	if apiResp != nil && apiResp.ErrorCode != nil && *apiResp.ErrorCode != "0" {
		message := ""
		if apiResp.Msg != nil {
			message = *apiResp.Msg
		}
		log.Printf("[ERROR]: Error in reading Security system resource In ReadSecuritySystem Block. Errorcode: %v, Message: %v", *apiResp.ErrorCode, message)
		return nil, fmt.Errorf("reading of Security System resource failed In ReadSecuritySystem Block: %s", message)
	}

	return apiResp, nil
}

// UpdateModelFromReadResponse - Extracted state management logic for read operations
func (r *SecuritySystemResource) UpdateModelFromReadResponse(plan *SecuritySystemResourceModel, apiResp *openapi.GetSecuritySystems200Response) {
	if len(apiResp.SecuritySystemDetails) == 0 {
		return
	}

	details := apiResp.SecuritySystemDetails[0]

	// Set the resource ID
	plan.ID = types.StringValue("security-system-" + *details.Systemname)

	// Update all fields from API response
	plan.Systemname = types.StringValue(util.SafeDeref(details.Systemname))
	plan.DisplayName = types.StringValue(util.SafeDeref(details.DisplayName))
	plan.Hostname = util.SafeString(details.Hostname)
	plan.Port = util.SafeString(details.Port)
	plan.AccessAddWorkflow = util.SafeString(details.AccessAddWorkflow)
	plan.AccessRemoveWorkflow = util.SafeString(details.AccessRemoveWorkflow)
	plan.AddServiceAccountWorkflow = util.SafeString(details.AddServiceAccountWorkflow)
	plan.RemoveServiceAccountWorkflow = util.SafeString(details.RemoveServiceAccountWorkflow)
	plan.Connectionparameters = util.SafeString(details.Connectionparameters)
	plan.AutomatedProvisioning = util.SafeString(details.AutomatedProvisioning)

	// Handle boolean field conversion
	if details.Useopenconnector == nil {
		plan.UseOpenConnector = types.BoolNull()
	} else {
		if *details.Useopenconnector == "true" {
			plan.UseOpenConnector = types.BoolValue(true)
		} else {
			plan.UseOpenConnector = types.BoolValue(false)
		}
	}

	plan.ReconApplication = util.SafeString(details.ReconApplication)
	plan.InstantProvision = util.SafeString(details.Instantprovision)
	plan.ProvisioningTries = util.SafeString(details.ProvisioningTries)
	plan.Provisioningcomments = util.SafeString(details.Provisioningcomments)
	plan.ProposedAccountOwnersWorkflow = util.SafeString(details.ProposedAccountOwnersworkflow)
	plan.FirefighterIDWorkflow = util.SafeString(details.FirefighteridWorkflow)
	plan.FirefighterIDRequestAccessWorkflow = util.SafeString(details.FirefighteridRequestAccessWorkflow)
	plan.PolicyRule = util.SafeString(details.PolicyRule)
	plan.PolicyRuleServiceAccount = util.SafeString(details.PolicyRuleServiceAccount)
	plan.Connectionname = util.SafeString(details.Connection)
	plan.ProvisioningConnection = util.SafeString(details.ProvisioningConnection)
	plan.ServiceDeskConnection = util.SafeString(details.ServiceDeskConnection)
	plan.ExternalRiskConnectionJson = util.SafeString(details.ExternalRiskConnectionJson)
	plan.InherentSODReportFields = util.StringsToSet(details.InherentSODReportFields)

	// Set API response fields
	plan.Msg = types.StringValue(util.SafeDeref(apiResp.Msg))
	plan.ErrorCode = types.StringValue(util.SafeDeref(apiResp.ErrorCode))
}

func (r *SecuritySystemResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
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

func (r *SecuritySystemResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	// Retrieve import ID and save to id attribute
	resource.ImportStatePassthroughID(ctx, path.Root("systemname"), req, resp)
}
