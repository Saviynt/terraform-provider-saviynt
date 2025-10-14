// Copyright (c) 2025 Saviynt Inc.
// SPDX-License-Identifier: MPL-2.0

// user_import_job_resource.go manages User Import Job triggers in Saviynt.
// The resource implements the full Terraform lifecycle:
//   - Create: provisions a new user import job trigger using the supplied configuration.
//   - Update: applies any configuration changes to an existing trigger.
//   - Delete: removes the trigger from Saviynt.
package provider

import (
	"context"
	"fmt"
	"net/http"
	"terraform-provider-Saviynt/internal/client"
	"terraform-provider-Saviynt/util"
	"terraform-provider-Saviynt/util/jobcontrolutil"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	openapi "github.com/saviynt/saviynt-api-go-client/job_control"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &UserImportJobResource{}

func NewUserImportJobResource() resource.Resource {
	return &UserImportJobResource{}
}

// UserImportJobResource defines the resource implementation.
type UserImportJobResource struct {
	client            client.SaviyntClientInterface
	provider          client.SaviyntProviderInterface
	jobControlFactory client.JobControlFactoryInterface
	token             string
}

// UserImportJobResourceModel describes the resource data model.
type UserImportJobResourceModel struct {
	Jobs types.List `tfsdk:"jobs"`
}

// UserImportJobModel describes individual job data model.
type UserImportJobModel struct {
	BaseJobControlResourceModel
	ExternalConn                       types.String `tfsdk:"external_conn"`
	FullOrIncremental                  types.String `tfsdk:"full_or_incremental"`
	UserNotInFeedAction                types.String `tfsdk:"user_not_in_feed_action"`
	UserOperationsAllowed              types.String `tfsdk:"user_operations_allowed"`
	ZeroDayProvisioning                types.String `tfsdk:"zero_day_provisioning"`
	GenerateSystemUsername             types.String `tfsdk:"generate_system_username"`
	GenerateEmail                      types.String `tfsdk:"generate_email"`
	CheckRules                         types.String `tfsdk:"check_rules"`
	BuildUserMap                       types.String `tfsdk:"build_user_map"`
	UserThreshold                      types.String `tfsdk:"user_threshold"`
	OnFailure                          types.String `tfsdk:"on_failure"`
	ZeroDayLimit                       types.String `tfsdk:"zero_day_limit"`
	TermUserLimit                      types.String `tfsdk:"term_user_limit"`
	ImportSavConnect                   types.String `tfsdk:"import_sav_connect"`
	ExportToSavCloud                   types.String `tfsdk:"export_to_sav_cloud"`
	UserReconciliationField            types.String `tfsdk:"user_reconciliation_field"`
	UserDefaultSavRole                 types.String `tfsdk:"user_default_sav_role"`
	UserStatusConfig                   types.String `tfsdk:"user_status_config"`
	EndpointsToAssociateOrphanAccounts types.String `tfsdk:"endpoints_to_associate_orphan_accounts"`
}

func UserImportJobResourceSchema() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"external_conn": schema.StringAttribute{
			Required:    true,
			Description: "External connection for the user import",
		},
		"full_or_incremental": schema.StringAttribute{
			Optional:    true,
			Description: "Full or incremental import type",
		},
		"user_not_in_feed_action": schema.StringAttribute{
			Optional:    true,
			Description: "Action to take for users not in feed",
		},
		"user_operations_allowed": schema.StringAttribute{
			Optional:    true,
			Description: "User operations allowed",
		},
		"zero_day_provisioning": schema.StringAttribute{
			Optional:    true,
			Description: "Zero day provisioning configuration",
		},
		"generate_system_username": schema.StringAttribute{
			Optional:    true,
			Description: "Generate system username configuration",
		},
		"generate_email": schema.StringAttribute{
			Optional:    true,
			Description: "Generate email configuration",
		},
		"check_rules": schema.StringAttribute{
			Optional:    true,
			Description: "Check rules configuration",
		},
		"build_user_map": schema.StringAttribute{
			Optional:    true,
			Description: "Build user map configuration",
		},
		"user_threshold": schema.StringAttribute{
			Optional:    true,
			Description: "User threshold configuration",
		},
		"on_failure": schema.StringAttribute{
			Optional:    true,
			Description: "Action to take on failure",
		},
		"zero_day_limit": schema.StringAttribute{
			Optional:    true,
			Description: "Zero day limit configuration",
		},
		"term_user_limit": schema.StringAttribute{
			Optional:    true,
			Description: "Term user limit configuration",
		},
		"import_sav_connect": schema.StringAttribute{
			Optional:    true,
			Description: "Import SAV connect configuration",
		},
		"export_to_sav_cloud": schema.StringAttribute{
			Optional:    true,
			Description: "Export to SAV cloud configuration",
		},
		"user_reconciliation_field": schema.StringAttribute{
			Optional:    true,
			Description: "User reconciliation field",
		},
		"user_default_sav_role": schema.StringAttribute{
			Optional:    true,
			Description: "User default SAV role",
		},
		"user_status_config": schema.StringAttribute{
			Optional:    true,
			Description: "User status configuration",
		},
		"endpoints_to_associate_orphan_accounts": schema.StringAttribute{
			Optional:    true,
			Description: "Endpoints to associate orphan accounts",
		},
	}
}

func (r *UserImportJobResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: util.UserImportJobDescription,
		Attributes: map[string]schema.Attribute{
			"jobs": schema.ListNestedAttribute{
				Required:    true,
				Description: "List of User Import Jobs to create",
				NestedObject: schema.NestedAttributeObject{
					Attributes: jobcontrolutil.MergeJobResourceAttributes(
						BaseJobControlResourceSchema(),
						UserImportJobResourceSchema(),
					),
				},
			},
		},
	}
}

func (r *UserImportJobResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_user_import_job_resource"
}

func (r *UserImportJobResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	tflog.Debug(ctx, "Starting UserImportJobResource configuration")

	if req.ProviderData == nil {
		tflog.Debug(ctx, "Provider data is nil, skipping configuration")
		return
	}

	prov, ok := req.ProviderData.(*SaviyntProvider)
	if !ok {
		errMsg := fmt.Sprintf("Expected *SaviyntProvider, got: %T", req.ProviderData)
		tflog.Error(ctx, "Type assertion failed", map[string]interface{}{
			"error": errMsg,
		})
		resp.Diagnostics.AddError(
			"Unexpected Resource Configure Type",
			fmt.Sprintf("Expected *SaviyntProvider, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)
		return
	}

	if prov.client == nil {
		tflog.Error(ctx, "Provider client is nil", map[string]interface{}{
			"error": "SaviyntProvider.client is nil",
		})
		resp.Diagnostics.AddError(
			"Provider Configuration Error",
			"Provider client is not initialized. Please check provider configuration.",
		)
		return
	}

	if prov.accessToken == "" {
		tflog.Error(ctx, "Access token is empty", map[string]interface{}{
			"error": "SaviyntProvider.accessToken is empty",
		})
		resp.Diagnostics.AddError(
			"Provider Authentication Error",
			"Access token is not available. Please check provider authentication.",
		)
		return
	}

	r.provider = &client.SaviyntProviderWrapper{Provider: prov}
	r.client = &client.SaviyntClientWrapper{Client: prov.client}
	r.jobControlFactory = &client.DefaultJobControlFactory{}
	r.token = prov.accessToken

	tflog.Info(ctx, "UserImportJobResource configuration completed successfully")
}

// SetClient sets the client for testing purposes
func (r *UserImportJobResource) SetClient(client client.SaviyntClientInterface) {
	r.client = client
}

// SetToken sets the token for testing purposes
func (r *UserImportJobResource) SetToken(token string) {
	r.token = token
}

// SetProvider sets the provider for testing purposes
func (r *UserImportJobResource) SetProvider(provider client.SaviyntProviderInterface) {
	r.provider = provider
}

// SetJobControlFactory sets the job control factory for testing purposes
func (r *UserImportJobResource) SetJobControlFactory(factory client.JobControlFactoryInterface) {
	r.jobControlFactory = factory
}

// buildUserImportJobValueMap creates a value map from job model
func buildUserImportJobValueMap(job UserImportJobModel) *openapi.UserImportJobAllOfValueMap {
	valueMap := openapi.NewUserImportJobAllOfValueMap(job.ExternalConn.ValueString())

	if !job.FullOrIncremental.IsNull() && job.FullOrIncremental.ValueString() != "" {
		valueMap.SetFullorincremental(job.FullOrIncremental.ValueString())
	}
	if !job.UserNotInFeedAction.IsNull() && job.UserNotInFeedAction.ValueString() != "" {
		valueMap.SetUserNotInFeedAction(job.UserNotInFeedAction.ValueString())
	}
	if !job.UserOperationsAllowed.IsNull() && job.UserOperationsAllowed.ValueString() != "" {
		valueMap.SetUserOperationsAllowed(job.UserOperationsAllowed.ValueString())
	}
	if !job.ZeroDayProvisioning.IsNull() && job.ZeroDayProvisioning.ValueString() != "" {
		valueMap.SetZeroDayProvisioning(job.ZeroDayProvisioning.ValueString())
	}
	if !job.GenerateSystemUsername.IsNull() && job.GenerateSystemUsername.ValueString() != "" {
		valueMap.SetGenerateSystemUsername(job.GenerateSystemUsername.ValueString())
	}
	if !job.GenerateEmail.IsNull() && job.GenerateEmail.ValueString() != "" {
		valueMap.SetGenerateEmail(job.GenerateEmail.ValueString())
	}
	if !job.CheckRules.IsNull() && job.CheckRules.ValueString() != "" {
		valueMap.SetCheckRules(job.CheckRules.ValueString())
	}
	if !job.BuildUserMap.IsNull() && job.BuildUserMap.ValueString() != "" {
		valueMap.SetBuildUserMap(job.BuildUserMap.ValueString())
	}
	if !job.UserThreshold.IsNull() && job.UserThreshold.ValueString() != "" {
		valueMap.SetUserThreshold(job.UserThreshold.ValueString())
	}
	if !job.OnFailure.IsNull() && job.OnFailure.ValueString() != "" {
		valueMap.SetOnFailure(job.OnFailure.ValueString())
	}
	if !job.ZeroDayLimit.IsNull() && job.ZeroDayLimit.ValueString() != "" {
		valueMap.SetZeroDayLimit(job.ZeroDayLimit.ValueString())
	}
	if !job.TermUserLimit.IsNull() && job.TermUserLimit.ValueString() != "" {
		valueMap.SetTermUserLimit(job.TermUserLimit.ValueString())
	}
	if !job.ImportSavConnect.IsNull() && job.ImportSavConnect.ValueString() != "" {
		valueMap.SetImportsavconnect(job.ImportSavConnect.ValueString())
	}
	if !job.ExportToSavCloud.IsNull() && job.ExportToSavCloud.ValueString() != "" {
		valueMap.SetExporttosavcloud(job.ExportToSavCloud.ValueString())
	}
	if !job.UserReconciliationField.IsNull() && job.UserReconciliationField.ValueString() != "" {
		valueMap.SetUserReconcillationField(job.UserReconciliationField.ValueString())
	}
	if !job.UserDefaultSavRole.IsNull() && job.UserDefaultSavRole.ValueString() != "" {
		valueMap.SetUserDefaultSavRole(job.UserDefaultSavRole.ValueString())
	}
	if !job.UserStatusConfig.IsNull() && job.UserStatusConfig.ValueString() != "" {
		valueMap.SetUserStatusConfig(job.UserStatusConfig.ValueString())
	}
	if !job.EndpointsToAssociateOrphanAccounts.IsNull() && job.EndpointsToAssociateOrphanAccounts.ValueString() != "" {
		valueMap.SetEndpointsToAssociateOrphanAccounts(job.EndpointsToAssociateOrphanAccounts.ValueString())
	}

	return valueMap
}

// CreateOrUpdateUserImportJobs handles the business logic for creating or updating User Import jobs
func (r *UserImportJobResource) CreateOrUpdateUserImportJobs(ctx context.Context, jobs []UserImportJobModel, operation string) (*openapi.CreateOrUpdateTriggersResponse, error) {
	if len(jobs) == 0 {
		return nil, fmt.Errorf("at least one job must be specified")
	}

	tflog.Debug(ctx, "Starting User Import Job triggers creation", map[string]interface{}{
		"job_count": len(jobs),
	})

	var triggers []openapi.TriggerItem

	// Process each job
	for i, job := range jobs {
		// Validate job name is UserImportJob
		if job.JobName.IsNull() || job.JobName.ValueString() != "UserImportJob" {
			return nil, fmt.Errorf("job %d: job_name must be 'UserImportJob', got '%s'", i+1, job.JobName.ValueString())
		}

		// Validate required fields
		if job.TriggerName.IsNull() || job.TriggerName.ValueString() == "" {
			return nil, fmt.Errorf("job %d: trigger_name is required", i+1)
		}
		if job.JobGroup.IsNull() || job.JobGroup.ValueString() == "" {
			return nil, fmt.Errorf("job %d: job_group is required", i+1)
		}
		if job.CronExpression.IsNull() || job.CronExpression.ValueString() == "" {
			return nil, fmt.Errorf("job %d: cron_expression is required", i+1)
		}
		if job.ExternalConn.IsNull() || job.ExternalConn.ValueString() == "" {
			return nil, fmt.Errorf("job %d: external_conn is required", i+1)
		}

		// Create the value map
		valueMap := buildUserImportJobValueMap(job)

		// Create the job trigger
		jobTrigger := openapi.NewUserImportJob(
			job.TriggerName.ValueString(),
			job.JobName.ValueString(),
			job.JobGroup.ValueString(),
			job.CronExpression.ValueString(),
		)
		jobTrigger.SetValueMap(*valueMap)

		// Set optional trigger group if provided
		if !job.TriggerGroup.IsNull() && job.TriggerGroup.ValueString() != "" {
			jobTrigger.SetTriggergroup(job.TriggerGroup.ValueString())
		}

		triggers = append(triggers, openapi.UserImportJobAsTriggerItem(jobTrigger))
	}

	// Create the request with all triggers
	createReq := openapi.CreateOrUpdateTriggersRequest{
		Triggers: triggers,
	}

	// Make the API call
	var apiResp *openapi.CreateOrUpdateTriggersResponse
	var finalHttpResp *http.Response
	err := r.provider.AuthenticatedAPICallWithRetry(ctx, fmt.Sprintf("%s_user_import_jobs", operation), func(token string) error {
		jobOps := r.jobControlFactory.CreateJobControlOperations(r.client.APIBaseURL(), token)
		apiResponse, httpResp, err := jobOps.CreateOrUpdateTriggers(ctx, createReq)
		if httpResp != nil && httpResp.StatusCode == 401 {
			return fmt.Errorf("401 unauthorized")
		}
		apiResp = apiResponse
		finalHttpResp = httpResp
		return err
	})

	if err != nil && finalHttpResp != nil && finalHttpResp.StatusCode != http.StatusPreconditionFailed {
		tflog.Error(ctx, "Error during API call", map[string]interface{}{
			"error": err.Error(),
		})
		return nil, fmt.Errorf("API call error: %s", err.Error())
	}

	// Handle HTTP errors
	var diags diag.Diagnostics
	if jobcontrolutil.JobControlHandleHTTPError(ctx, finalHttpResp, err, "creating User Import Job triggers", &diags) {
		tflog.Error(ctx, "Failed to create User Import Job triggers", map[string]interface{}{
			"error": fmt.Sprintf("%v", diags.Errors()),
		})
		if diags.HasError() {
			return nil, fmt.Errorf("HTTP error: %s", diags.Errors()[0].Detail())
		}
	}

	// Handle API response errors
	if apiResp != nil && jobcontrolutil.JobControlHandleAPIError(ctx, &apiResp.ErrorCode, &apiResp.Msg, "creating User Import Job triggers", &diags) {
		tflog.Error(ctx, "API error during trigger creation", map[string]interface{}{
			"error": fmt.Sprintf("%v", diags.Errors()),
		})
		if diags.HasError() {
			return nil, fmt.Errorf("API error: %s", diags.Errors()[0].Detail())
		}
	}

	tflog.Info(ctx, "User Import Job triggers created successfully", map[string]interface{}{
		"job_count": len(jobs),
	})

	return apiResp, nil
}

// DeleteUserImportJobs handles the business logic for deleting User Import jobs
func (r *UserImportJobResource) DeleteUserImportJobs(ctx context.Context, jobs []UserImportJobModel) error {
	if len(jobs) == 0 {
		return nil
	}

	tflog.Debug(ctx, "Starting User Import Job triggers deletion", map[string]interface{}{
		"job_count": len(jobs),
	})

	// Delete each job individually
	for i, job := range jobs {
		triggerName := job.TriggerName.ValueString()
		jobName := job.JobName.ValueString()
		jobGroup := job.JobGroup.ValueString()

		tflog.Debug(ctx, "Deleting job trigger", map[string]interface{}{
			"job_index":    i + 1,
			"trigger_name": triggerName,
		})

		// Create delete request
		deleteReq := openapi.DeleteTriggerRequest{
			Triggername: triggerName,
			Jobname:     jobName,
			Jobgroup:    jobGroup,
		}

		// Make the API call
		var apiResp *openapi.DeleteTriggerResponse
		var finalHttpResp *http.Response
		err := r.provider.AuthenticatedAPICallWithRetry(ctx, "delete_user_import_job", func(token string) error {
			jobOps := r.jobControlFactory.CreateJobControlOperations(r.client.APIBaseURL(), token)
			apiResponse, httpResp, err := jobOps.DeleteTrigger(ctx, deleteReq)
			if httpResp != nil && httpResp.StatusCode == 401 {
				return fmt.Errorf("401 unauthorized")
			}
			apiResp = apiResponse
			finalHttpResp = httpResp
			return err
		})

		if err != nil && finalHttpResp != nil && finalHttpResp.StatusCode != http.StatusPreconditionFailed {
			tflog.Error(ctx, "Error during API call", map[string]interface{}{
				"error":        err.Error(),
				"trigger_name": triggerName,
			})
			return fmt.Errorf("API call error for trigger '%s': %s", triggerName, err.Error())
		}

		// Handle HTTP errors
		var diags diag.Diagnostics
		if jobcontrolutil.JobControlHandleHTTPError(ctx, finalHttpResp, err, fmt.Sprintf("deleting User Import Job trigger '%s'", triggerName), &diags) {
			tflog.Error(ctx, "Failed to delete User Import Job trigger", map[string]interface{}{
				"error":        fmt.Sprintf("%v", diags.Errors()),
				"trigger_name": triggerName,
			})
			if diags.HasError() {
				return fmt.Errorf("HTTP error for trigger '%s': %s", triggerName, diags.Errors()[0].Detail())
			}
		}

		// Handle API response errors
		if apiResp != nil {
			errorCodeStr := fmt.Sprintf("%d", apiResp.ErrorCode)
			if jobcontrolutil.JobControlHandleAPIError(ctx, &errorCodeStr, &apiResp.Msg, fmt.Sprintf("deleting User Import Job trigger '%s'", triggerName), &diags) {
				tflog.Error(ctx, "API error during trigger deletion", map[string]interface{}{
					"error":        fmt.Sprintf("%v", diags.Errors()),
					"trigger_name": triggerName,
				})
				if diags.HasError() {
					return fmt.Errorf("API error for trigger '%s': %s", triggerName, diags.Errors()[0].Detail())
				}
			}
		}

		tflog.Info(ctx, "User Import Job trigger deleted successfully", map[string]interface{}{
			"trigger_name": triggerName,
		})
	}

	tflog.Info(ctx, "All User Import Job triggers deleted successfully", map[string]interface{}{
		"job_count": len(jobs),
	})

	return nil
}

func (r *UserImportJobResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan UserImportJobResourceModel

	tflog.Debug(ctx, "Starting User Import Job resource creation")

	// Extract plan from request
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		resp.Diagnostics.AddError(
			"Plan Extraction Failed",
			"Unable to extract Terraform plan from request",
		)
		return
	}

	// Extract jobs from the plan
	var jobs []UserImportJobModel
	resp.Diagnostics.Append(plan.Jobs.ElementsAs(ctx, &jobs, false)...)
	if resp.Diagnostics.HasError() {
		resp.Diagnostics.AddError(
			"Jobs Extraction Failed",
			"Unable to extract jobs from Terraform plan",
		)
		return
	}

	// Call the business logic method
	_, err := r.CreateOrUpdateUserImportJobs(ctx, jobs, "create")
	if err != nil {
		tflog.Error(ctx, "User Import Job creation failed", map[string]interface{}{
			"error": err.Error(),
		})
		resp.Diagnostics.AddError(
			"User Import Job Creation Failed",
			err.Error(),
		)
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		resp.Diagnostics.AddError(
			"State Update Failed",
			"Unable to save User Import Job state",
		)
		return
	}

	tflog.Info(ctx, "User Import Job resource created successfully", map[string]interface{}{
		"job_count": len(jobs),
	})
}

func (r *UserImportJobResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state UserImportJobResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading User Import Job triggers", map[string]interface{}{
		"job_count": len(state.Jobs.Elements()),
	})

	// For now, we'll keep the state as-is since the API doesn't provide a direct read method
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *UserImportJobResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan UserImportJobResourceModel

	tflog.Debug(ctx, "Starting User Import Job resource update")

	// Extract plan from request
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		resp.Diagnostics.AddError(
			"Plan Extraction Failed",
			"Unable to extract Terraform plan from request",
		)
		return
	}

	// Extract jobs from the plan
	var jobs []UserImportJobModel
	resp.Diagnostics.Append(plan.Jobs.ElementsAs(ctx, &jobs, false)...)
	if resp.Diagnostics.HasError() {
		resp.Diagnostics.AddError(
			"Jobs Extraction Failed",
			"Unable to extract jobs from Terraform plan",
		)
		return
	}

	// Call the business logic method
	_, err := r.CreateOrUpdateUserImportJobs(ctx, jobs, "update")
	if err != nil {
		tflog.Error(ctx, "User Import Job update failed", map[string]interface{}{
			"error": err.Error(),
		})
		resp.Diagnostics.AddError(
			"User Import Job Update Failed",
			err.Error(),
		)
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		resp.Diagnostics.AddError(
			"State Update Failed",
			"Unable to save User Import Job state",
		)
		return
	}

	tflog.Info(ctx, "User Import Job resource updated successfully", map[string]interface{}{
		"job_count": len(jobs),
	})
}

func (r *UserImportJobResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state UserImportJobResourceModel

	tflog.Debug(ctx, "Starting User Import Job resource deletion")

	// Extract state from request
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		resp.Diagnostics.AddError(
			"State Extraction Failed",
			"Unable to extract Terraform state from request",
		)
		return
	}

	// Extract jobs from the state
	var jobs []UserImportJobModel
	resp.Diagnostics.Append(state.Jobs.ElementsAs(ctx, &jobs, false)...)
	if resp.Diagnostics.HasError() {
		resp.Diagnostics.AddError(
			"Jobs Extraction Failed",
			"Unable to extract jobs from Terraform state",
		)
		return
	}

	// Call the business logic method
	err := r.DeleteUserImportJobs(ctx, jobs)
	if err != nil {
		tflog.Error(ctx, "User Import Job deletion failed", map[string]interface{}{
			"error": err.Error(),
		})
		resp.Diagnostics.AddError(
			"User Import Job Deletion Failed",
			err.Error(),
		)
		return
	}

	tflog.Info(ctx, "User Import Job resource deleted successfully", map[string]interface{}{
		"job_count": len(jobs),
	})
}
