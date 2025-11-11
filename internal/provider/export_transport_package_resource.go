// Copyright (c) 2025 Saviynt Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"terraform-provider-Saviynt/internal/client"
	"terraform-provider-Saviynt/util"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	openapi "github.com/saviynt/saviynt-api-go-client/transports"
)

var _ resource.Resource = &ExportTransportPackageResource{}

type ExportTransportPackageResourceModel struct {
	ID                    types.String `tfsdk:"id"`
	UpdateUser            types.String `tfsdk:"update_user"`
	TransportOwner        types.String `tfsdk:"transport_owner"`
	TransportMembers      types.String `tfsdk:"transport_members"`
	ExportOnline          types.String `tfsdk:"export_online"`
	ExportPath            types.String `tfsdk:"export_path"`
	EnvironmentName       types.String `tfsdk:"environment_name"`
	BusinessJustification types.String `tfsdk:"business_justification"`
	ExportPackageVersion  types.String `tfsdk:"export_package_version"`

	// Objects to export
	SavRoles        types.List `tfsdk:"sav_roles"`
	EmailTemplate   types.List `tfsdk:"email_template"`
	Roles           types.List `tfsdk:"roles"`
	AnalyticsV1     types.List `tfsdk:"analytics_v1"`
	AnalyticsV2     types.List `tfsdk:"analytics_v2"`
	GlobalConfig    types.List `tfsdk:"global_config"`
	Workflows       types.List `tfsdk:"workflows"`
	Connections     types.List `tfsdk:"connections"`
	AppOnboarding   types.List `tfsdk:"app_onboarding"`
	UserGroups      types.List `tfsdk:"user_groups"`
	ScanRules       types.List `tfsdk:"scan_rules"`
	Organizations   types.List `tfsdk:"organizations"`
	SecuritySystems types.List `tfsdk:"security_systems"`

	// Response fields
	Message        types.String `tfsdk:"message"`
	FileName       types.String `tfsdk:"file_name"`
	MsgDescription types.String `tfsdk:"msg_description"`
	ErrorCode      types.String `tfsdk:"error_code"`
}

type ExportTransportPackageResource struct {
	client           client.SaviyntClientInterface
	token            string
	provider         client.SaviyntProviderInterface
	transportFactory client.TransportFactoryInterface
}

func NewExportTransportPackageResource() resource.Resource {
	return &ExportTransportPackageResource{
		transportFactory: &client.DefaultTransportFactory{},
	}
}

func NewExportTransportPackageResourceWithFactory(factory client.TransportFactoryInterface) resource.Resource {
	return &ExportTransportPackageResource{
		transportFactory: factory,
	}
}

func (r *ExportTransportPackageResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_export_transport_package_resource"
}

func (r *ExportTransportPackageResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: util.ExportTransportPackageDescription,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The unique ID of the resource.",
			},
			"update_user": schema.StringAttribute{
				Optional:    true,
				Description: "Username of the user exporting the package.",
			},
			"transport_owner": schema.StringAttribute{
				Optional:    true,
				Description: "Option to transport owners for selected objects.",
			},
			"transport_members": schema.StringAttribute{
				Optional:    true,
				Description: "Option to transport members for selected objects such as SAV role.",
			},
			"export_online": schema.StringAttribute{
				Required:    true,
				Description: "Determines if package needs to be exported online (true/false).",
			},
			"export_path": schema.StringAttribute{
				Required:    true,
				Description: "Local path where export package will be generated (required if export_online is false).",
			},
			"environment_name": schema.StringAttribute{
				Optional:    true,
				Description: "Name of the environment (required if export_online is true).",
			},
			"business_justification": schema.StringAttribute{
				Optional:    true,
				Description: "Business justification for the export.",
			},
			"export_package_version": schema.StringAttribute{
				Optional:    true,
				Description: "Version identifier for the export package. Change this value to trigger re-export.",
			},

			// Objects to export
			"sav_roles": schema.ListAttribute{
				Optional:    true,
				ElementType: types.StringType,
				Description: "List of SAV roles to export.",
			},
			"email_template": schema.ListAttribute{
				Optional:    true,
				ElementType: types.StringType,
				Description: "List of email templates to export.",
			},
			"roles": schema.ListAttribute{
				Optional:    true,
				ElementType: types.StringType,
				Description: "List of roles to export.",
			},
			"analytics_v1": schema.ListAttribute{
				Optional:    true,
				ElementType: types.StringType,
				Description: "List of analytics v1 to export.",
			},
			"analytics_v2": schema.ListAttribute{
				Optional:    true,
				ElementType: types.StringType,
				Description: "List of analytics v2 to export.",
			},
			"global_config": schema.ListAttribute{
				Optional:    true,
				ElementType: types.StringType,
				Description: "List of global configurations to export.",
			},
			"workflows": schema.ListAttribute{
				Optional:    true,
				ElementType: types.StringType,
				Description: "List of workflows to export.",
			},
			"connections": schema.ListAttribute{
				Optional:    true,
				ElementType: types.StringType,
				Description: "List of connections to export.",
			},
			"app_onboarding": schema.ListAttribute{
				Optional:    true,
				ElementType: types.StringType,
				Description: "List of app onboarding configurations to export.",
			},
			"user_groups": schema.ListAttribute{
				Optional:    true,
				ElementType: types.StringType,
				Description: "List of user groups to export.",
			},
			"scan_rules": schema.ListAttribute{
				Optional:    true,
				ElementType: types.StringType,
				Description: "List of scan rules to export.",
			},
			"organizations": schema.ListAttribute{
				Optional:    true,
				ElementType: types.StringType,
				Description: "List of organizations to export.",
			},
			"security_systems": schema.ListAttribute{
				Optional:    true,
				ElementType: types.StringType,
				Description: "List of security systems to export.",
			},

			// Response fields
			"message": schema.StringAttribute{
				Computed:    true,
				Description: "Response message from the export operation.",
			},
			"file_name": schema.StringAttribute{
				Computed:    true,
				Description: "Name of the generated transport package file.",
			},
			"msg_description": schema.StringAttribute{
				Computed:    true,
				Description: "Detailed description of the response.",
			},
			"error_code": schema.StringAttribute{
				Computed:    true,
				Description: "Error code from the export operation.",
			},
		},
	}
}

func (r *ExportTransportPackageResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	tflog.Debug(ctx, "Starting ExportTransportPackageResource configuration")

	if req.ProviderData == nil {
		tflog.Debug(ctx, "Provider data is nil, skipping configuration")
		return
	}

	prov, ok := req.ProviderData.(*SaviyntProvider)
	if !ok {
		tflog.Error(ctx, "Type assertion failed", map[string]interface{}{
			"error": fmt.Sprintf("Expected *SaviyntProvider, got: %T", req.ProviderData),
		})
		resp.Diagnostics.AddError(
			"Unexpected Resource Configure Type",
			fmt.Sprintf("Expected *SaviyntProvider, got: %T", req.ProviderData),
		)
		return
	}

	r.client = &client.SaviyntClientWrapper{Client: prov.client}
	r.token = prov.accessToken
	r.provider = &client.SaviyntProviderWrapper{Provider: prov}

	tflog.Debug(ctx, "ExportTransportPackageResource configured successfully")
}

// SetClient sets the client for testing purposes
func (r *ExportTransportPackageResource) SetClient(client client.SaviyntClientInterface) {
	r.client = client
}

// SetToken sets the token for testing purposes
func (r *ExportTransportPackageResource) SetToken(token string) {
	r.token = token
}

// SetProvider sets the provider for testing purposes
func (r *ExportTransportPackageResource) SetProvider(provider client.SaviyntProviderInterface) {
	r.provider = provider
}

// ExportTransportPackage handles the business logic for exporting transport packages
func (r *ExportTransportPackageResource) ExportTransportPackage(ctx context.Context, plan *ExportTransportPackageResourceModel, operation string) (*openapi.ExportTransportPackageResponse, error) {
	// Build objects to export
	objectsToExport := openapi.ExportTransportPackageRequestObjectstoexport{}

	// Convert lists to string slices
	if !plan.SavRoles.IsNull() {
		var savRoles []string
		plan.SavRoles.ElementsAs(ctx, &savRoles, false)
		objectsToExport.SavRoles = savRoles
	}
	if !plan.EmailTemplate.IsNull() {
		var emailTemplate []string
		plan.EmailTemplate.ElementsAs(ctx, &emailTemplate, false)
		objectsToExport.EmailTemplate = emailTemplate
	}
	if !plan.Roles.IsNull() {
		var roles []string
		plan.Roles.ElementsAs(ctx, &roles, false)
		objectsToExport.Roles = roles
	}
	if !plan.AnalyticsV1.IsNull() {
		var analyticsV1 []string
		plan.AnalyticsV1.ElementsAs(ctx, &analyticsV1, false)
		objectsToExport.AnalyticsV1 = analyticsV1
	}
	if !plan.AnalyticsV2.IsNull() {
		var analyticsV2 []string
		plan.AnalyticsV2.ElementsAs(ctx, &analyticsV2, false)
		objectsToExport.AnalyticsV2 = analyticsV2
	}
	if !plan.GlobalConfig.IsNull() {
		var globalConfig []string
		plan.GlobalConfig.ElementsAs(ctx, &globalConfig, false)
		objectsToExport.GlobalConfig = globalConfig
	}
	if !plan.Workflows.IsNull() {
		var workflows []string
		plan.Workflows.ElementsAs(ctx, &workflows, false)
		objectsToExport.Workflows = workflows
	}
	if !plan.Connections.IsNull() {
		var connections []string
		plan.Connections.ElementsAs(ctx, &connections, false)
		objectsToExport.Connection = connections
	}
	if !plan.AppOnboarding.IsNull() {
		var appOnboarding []string
		plan.AppOnboarding.ElementsAs(ctx, &appOnboarding, false)
		objectsToExport.AppOnboarding = appOnboarding
	}
	if !plan.UserGroups.IsNull() {
		var userGroups []string
		plan.UserGroups.ElementsAs(ctx, &userGroups, false)
		objectsToExport.UserGroups = userGroups
	}
	if !plan.ScanRules.IsNull() {
		var scanRules []string
		plan.ScanRules.ElementsAs(ctx, &scanRules, false)
		objectsToExport.ScanRules = scanRules
	}
	if !plan.Organizations.IsNull() {
		var organizations []string
		plan.Organizations.ElementsAs(ctx, &organizations, false)
		objectsToExport.Organizations = organizations
	}
	if !plan.SecuritySystems.IsNull() {
		var securitySystems []string
		plan.SecuritySystems.ElementsAs(ctx, &securitySystems, false)
		objectsToExport.SecuritySystems = securitySystems
	}

	// Build export request
	exportReq := openapi.ExportTransportPackageRequest{
		Exportonline:    plan.ExportOnline.ValueString(),
		Exportpath:      plan.ExportPath.ValueString(),
		Objectstoexport: objectsToExport,
	}

	// Add optional fields
	if !plan.UpdateUser.IsNull() && plan.UpdateUser.ValueString() != "" {
		exportReq.Updateuser = plan.UpdateUser.ValueStringPointer()
	}
	if !plan.TransportOwner.IsNull() && plan.TransportOwner.ValueString() != "" {
		exportReq.Transportowner = plan.TransportOwner.ValueStringPointer()
	}
	if !plan.TransportMembers.IsNull() && plan.TransportMembers.ValueString() != "" {
		exportReq.Transportmembers = plan.TransportMembers.ValueStringPointer()
	}
	if !plan.EnvironmentName.IsNull() && plan.EnvironmentName.ValueString() != "" {
		exportReq.Environmentname = plan.EnvironmentName.ValueStringPointer()
	}
	if !plan.BusinessJustification.IsNull() && plan.BusinessJustification.ValueString() != "" {
		exportReq.Businessjustification = plan.BusinessJustification.ValueStringPointer()
	}

	var apiResp *openapi.ExportTransportPackageResponse
	var finalHttpResp *http.Response

	// Execute export operation with retry logic
	err := r.provider.AuthenticatedAPICallWithRetry(ctx, fmt.Sprintf("%s_export_transport_package", operation), func(token string) error {
		transportOps := r.transportFactory.CreateTransportOperations(r.client.APIBaseURL(), token)
		resp, httpResp, err := transportOps.ExportTransportPackage(ctx, exportReq)
		if httpResp != nil && httpResp.StatusCode == 401 {
			return fmt.Errorf("401 unauthorized")
		}
		apiResp = resp
		finalHttpResp = httpResp
		return err
	})

	// Handle non-412 errors first
	if err != nil && finalHttpResp != nil && finalHttpResp.StatusCode != 412 {
		return nil, fmt.Errorf("API call failed: %w", err)
	}

	// Handle 412 precondition failed with response body decoding
	if err != nil && finalHttpResp != nil && finalHttpResp.StatusCode == 412 {
		var errorResp struct {
			Msg            string `json:"msg"`
			MsgDescription string `json:"msgDescription"`
			ErrorCode      int    `json:"errorcode"`
		}

		if decodeErr := json.NewDecoder(finalHttpResp.Body).Decode(&errorResp); decodeErr == nil {
			return nil, fmt.Errorf("precondition failed - ErrorCode: %d, Msg: %s, Description: %s",
				errorResp.ErrorCode, errorResp.Msg, errorResp.MsgDescription)
		}
		return nil, fmt.Errorf("precondition failed: ensure at least one object type is specified for export")
	}

	if err != nil {
		return nil, fmt.Errorf("API call failed: %w", err)
	}

	return apiResp, nil
}

// UpdateModelFromResponse updates the model with API response data
func (r *ExportTransportPackageResource) UpdateModelFromResponse(plan *ExportTransportPackageResourceModel, apiResp *openapi.ExportTransportPackageResponse) {
	plan.ID = types.StringValue("export-transport-" + plan.ExportPath.ValueString())

	// Handle nil pointer dereference safely
	if apiResp != nil {
		if apiResp.Msg != nil {
			plan.Message = types.StringValue(*apiResp.Msg)
		} else {
			plan.Message = types.StringValue("")
		}
		if apiResp.FileName != nil {
			plan.FileName = types.StringValue(*apiResp.FileName)
		} else {
			plan.FileName = types.StringValue("")
		}
		if apiResp.MsgDescription != nil {
			plan.MsgDescription = types.StringValue(*apiResp.MsgDescription)
		} else {
			plan.MsgDescription = types.StringValue("")
		}
		if apiResp.Errorcode != nil {
			plan.ErrorCode = types.StringValue(fmt.Sprintf("%d", *apiResp.Errorcode))
		} else {
			plan.ErrorCode = types.StringValue("0")
		}
	} else {
		// Set null values when apiResp is nil to prevent framework errors
		plan.Message = types.StringNull()
		plan.FileName = types.StringNull()
		plan.MsgDescription = types.StringNull()
		plan.ErrorCode = types.StringNull()
	}
}

// AddSuccessWarning adds a warning for successful export operations
func (r *ExportTransportPackageResource) AddSuccessWarning(resp interface{}, apiResp *openapi.ExportTransportPackageResponse) {
	// Check if this is a successful operation (error_code = 0)
	if apiResp != nil && apiResp.Errorcode != nil && *apiResp.Errorcode == 0 {
		var message, fileName, errorCode string

		if apiResp.Msg != nil {
			message = *apiResp.Msg
		}
		if apiResp.FileName != nil {
			fileName = *apiResp.FileName
		}
		if apiResp.Errorcode != nil {
			errorCode = fmt.Sprintf("%d", *apiResp.Errorcode)
		}

		// Add warning for successful export
		if createResp, ok := resp.(*resource.CreateResponse); ok {
			createResp.Diagnostics.AddWarning(
				"Transport Package Export Completed Successfully",
				fmt.Sprintf(
					"Transport package exported successfully.\nMessage: %s\nError Code: %s\nFile Name: %s\n\n"+
						"⚠️  IMPORTANT: The transport package has been generated. "+
						"Please verify the exported file and use it for importing into target environments.",
					message, errorCode, fileName,
				),
			)
		} else if updateResp, ok := resp.(*resource.UpdateResponse); ok {
			updateResp.Diagnostics.AddWarning(
				"Transport Package Export Completed Successfully",
				fmt.Sprintf(
					"Transport package exported successfully.\nMessage: %s\nError Code: %s\nFile Name: %s\n\n"+
						"⚠️  IMPORTANT: The transport package has been generated. "+
						"Please verify the exported file and use it for importing into target environments.",
					message, errorCode, fileName,
				),
			)
		}
	}
}

func (r *ExportTransportPackageResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan ExportTransportPackageResourceModel

	tflog.Debug(ctx, "Starting export transport package resource creation")

	// Extract plan from request
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		resp.Diagnostics.AddError(
			"Plan Extraction Failed",
			"Unable to extract Terraform plan from request",
		)
		return
	}

	// Call the business logic method
	apiResp, err := r.ExportTransportPackage(ctx, &plan, "create")
	if err != nil {
		tflog.Error(ctx, "Export transport package creation failed", map[string]interface{}{
			"error": err.Error(),
		})
		resp.Diagnostics.AddError(
			"Export Transport Package Creation Failed",
			err.Error(),
		)
		return
	}

	// Update model from response
	r.UpdateModelFromResponse(&plan, apiResp)

	// Add success warning if operation completed successfully
	r.AddSuccessWarning(resp, apiResp)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		resp.Diagnostics.AddError(
			"State Update Failed",
			"Unable to save export transport package state",
		)
		return
	}

	tflog.Info(ctx, "Export transport package resource created successfully")
}

func (r *ExportTransportPackageResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state ExportTransportPackageResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Export operations are one-time, so we just keep the existing state
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *ExportTransportPackageResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan ExportTransportPackageResourceModel

	tflog.Debug(ctx, "Starting export transport package resource update")

	// Extract plan from request
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		resp.Diagnostics.AddError(
			"Plan Extraction Failed",
			"Unable to extract Terraform plan from request",
		)
		return
	}

	// Call the business logic method
	apiResp, err := r.ExportTransportPackage(ctx, &plan, "update")
	if err != nil {
		tflog.Error(ctx, "Export transport package update failed", map[string]interface{}{
			"error": err.Error(),
		})
		resp.Diagnostics.AddError(
			"Export Transport Package Update Failed",
			err.Error(),
		)
		return
	}

	// Update model from response
	r.UpdateModelFromResponse(&plan, apiResp)

	// Add success warning if operation completed successfully
	r.AddSuccessWarning(resp, apiResp)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		resp.Diagnostics.AddError(
			"State Update Failed",
			"Unable to save export transport package state",
		)
		return
	}

	tflog.Info(ctx, "Export transport package resource updated successfully")
}

func (r *ExportTransportPackageResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// For acceptance tests only
	if os.Getenv("TF_ACC") == "1" {
		resp.State.RemoveResource(ctx)
		return
	}

	resp.Diagnostics.AddError(
		"Delete Not Supported",
		"Export transport package operations cannot be undone. Please remove the resource manually if required, or contact your administrator.",
	)
}
