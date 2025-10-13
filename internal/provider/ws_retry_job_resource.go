// Copyright (c) 2025 Saviynt Inc.
// SPDX-License-Identifier: MPL-2.0

// ws_retry_job_resource.go manages WS Retry Job triggers in Saviynt.
// The resource implements the full Terraform lifecycle:
//   - Create: provisions a new WS retry job trigger using the supplied configuration.
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
var _ resource.Resource = &WSRetryJobResource{}

func NewWSRetryJobResource() resource.Resource {
	return &WSRetryJobResource{}
}

// WSRetryJobResource defines the resource implementation.
type WSRetryJobResource struct {
	client            client.SaviyntClientInterface
	provider          client.SaviyntProviderInterface
	jobControlFactory client.JobControlFactoryInterface
	token             string
}

// WSRetryJobResourceModel describes the resource data model.
type WSRetryJobResourceModel struct {
	Jobs types.List `tfsdk:"jobs"`
}

// WSRetryJobModel describes individual job data model.
type WSRetryJobModel struct {
	BaseJobControlResourceModel
	SecuritySystems types.List   `tfsdk:"security_systems"`
	TaskTypes       types.String `tfsdk:"task_types"`
}

func WSRetryJobResourceSchema() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"security_systems": schema.ListAttribute{
			Optional:    true,
			ElementType: types.StringType,
			Description: "List of security systems for the WS retry job",
		},
		"task_types": schema.StringAttribute{
			Optional:    true,
			Description: "Task types for the WS retry job",
		},
	}
}

func (r *WSRetryJobResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: util.WSRetryJobDescription,
		Attributes: map[string]schema.Attribute{
			"jobs": schema.ListNestedAttribute{
				Required:    true,
				Description: "List of WS Retry Jobs to create",
				NestedObject: schema.NestedAttributeObject{
					Attributes: jobcontrolutil.MergeJobResourceAttributes(
						BaseJobControlResourceSchema(),
						WSRetryJobResourceSchema(),
					),
				},
			},
		},
	}
}

func (r *WSRetryJobResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_ws_retry_job_resource"
}

func (r *WSRetryJobResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	tflog.Debug(ctx, "Starting WSRetryJobResource configuration")

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

	tflog.Info(ctx, "WSRetryJobResource configuration completed successfully")
}

func (r *WSRetryJobResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan WSRetryJobResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Extract jobs from the plan
	var jobs []WSRetryJobModel
	resp.Diagnostics.Append(plan.Jobs.ElementsAs(ctx, &jobs, false)...)
	if resp.Diagnostics.HasError() {
		return
	}

	if len(jobs) == 0 {
		resp.Diagnostics.AddError("Validation Error", "At least one job must be specified")
		return
	}

	tflog.Debug(ctx, "Starting WS Retry Job triggers creation", map[string]interface{}{
		"job_count": len(jobs),
	})

	var triggers []openapi.TriggerItem

	// Process each job
	for i, job := range jobs {
		// Validate job name is WSRetryJob
		if job.JobName.IsNull() || job.JobName.ValueString() != "WSRetryJob" {
			resp.Diagnostics.AddError(
				"Validation Error", 
				fmt.Sprintf("Job %d: job_name must be 'WSRetryJob', got '%s'", i+1, job.JobName.ValueString()),
			)
			return
		}

		// Validate required fields
		if job.TriggerName.IsNull() || job.TriggerName.ValueString() == "" {
			resp.Diagnostics.AddError("Validation Error", fmt.Sprintf("Job %d: trigger_name is required", i+1))
			return
		}
		if job.JobGroup.IsNull() || job.JobGroup.ValueString() == "" {
			resp.Diagnostics.AddError("Validation Error", fmt.Sprintf("Job %d: job_group is required", i+1))
			return
		}
		if job.CronExpression.IsNull() || job.CronExpression.ValueString() == "" {
			resp.Diagnostics.AddError("Validation Error", fmt.Sprintf("Job %d: cron_expression is required", i+1))
			return
		}

		// Create the job trigger
		jobTrigger := openapi.NewWSRetryJob(
			job.TriggerName.ValueString(),
			job.JobName.ValueString(),
			job.JobGroup.ValueString(),
			job.CronExpression.ValueString(),
		)

		// Set optional security systems if provided
		if !job.SecuritySystems.IsNull() && len(job.SecuritySystems.Elements()) > 0 {
			var securitySystems []string
			resp.Diagnostics.Append(job.SecuritySystems.ElementsAs(ctx, &securitySystems, false)...)
			if resp.Diagnostics.HasError() {
				return
			}
			jobTrigger.SetSecuritysystems(securitySystems)
		}

		// Set optional task types if provided
		if !job.TaskTypes.IsNull() && job.TaskTypes.ValueString() != "" {
			jobTrigger.SetTasktypes(job.TaskTypes.ValueString())
		}

		// Set optional trigger group if provided
		if !job.TriggerGroup.IsNull() && job.TriggerGroup.ValueString() != "" {
			jobTrigger.SetTriggergroup(job.TriggerGroup.ValueString())
		}

		triggers = append(triggers, openapi.WSRetryJobAsTriggerItem(jobTrigger))
	}

	// Create the request with all triggers
	createReq := openapi.CreateOrUpdateTriggersRequest{
		Triggers: triggers,
	}

	// Make the API call
	var apiResp *openapi.CreateOrUpdateTriggersResponse
	var finalHttpResp *http.Response
	err := r.provider.AuthenticatedAPICallWithRetry(ctx, "create_ws_retry_jobs", func(token string) error {
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
		resp.Diagnostics.AddError(
			"API Call Error",
			fmt.Sprintf("Error during API call to create WS Retry Job triggers: %s", err.Error()),
		)
		return
	}

	// Handle HTTP errors
	var diags diag.Diagnostics
	if jobcontrolutil.JobControlHandleHTTPError(ctx, finalHttpResp, err, "creating WS Retry Job triggers", &diags) {
		tflog.Error(ctx, "Failed to create WS Retry Job triggers", map[string]interface{}{
			"error": fmt.Sprintf("%v", diags.Errors()),
		})
		for _, diagnostic := range diags {
			if diagnostic.Severity() == diag.SeverityError {
				resp.Diagnostics.AddError(diagnostic.Summary(), diagnostic.Detail())
			}
		}
		return
	}

	// Handle API response errors
	if apiResp != nil && jobcontrolutil.JobControlHandleAPIError(ctx, &apiResp.ErrorCode, &apiResp.Msg, "creating WS Retry Job triggers", &diags) {
		tflog.Error(ctx, "API error during trigger creation", map[string]interface{}{
			"error": fmt.Sprintf("%v", diags.Errors()),
		})
		for _, diagnostic := range diags {
			if diagnostic.Severity() == diag.SeverityError {
				resp.Diagnostics.AddError(diagnostic.Summary(), diagnostic.Detail())
			}
		}
		return
	}

	tflog.Info(ctx, "WS Retry Job triggers created successfully", map[string]interface{}{
		"job_count": len(jobs),
	})

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *WSRetryJobResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state WSRetryJobResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading WS Retry Job triggers", map[string]interface{}{
		"job_count": len(state.Jobs.Elements()),
	})

	// For now, we'll keep the state as-is since the API doesn't provide a direct read method
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *WSRetryJobResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan WSRetryJobResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Extract jobs from the plan
	var jobs []WSRetryJobModel
	resp.Diagnostics.Append(plan.Jobs.ElementsAs(ctx, &jobs, false)...)
	if resp.Diagnostics.HasError() {
		return
	}

	if len(jobs) == 0 {
		resp.Diagnostics.AddError("Validation Error", "At least one job must be specified")
		return
	}

	tflog.Debug(ctx, "Starting WS Retry Job triggers update", map[string]interface{}{
		"job_count": len(jobs),
	})

	var triggers []openapi.TriggerItem

	// Process each job
	for i, job := range jobs {
		// Validate job name is WSRetryJob
		if job.JobName.IsNull() || job.JobName.ValueString() != "WSRetryJob" {
			resp.Diagnostics.AddError(
				"Validation Error", 
				fmt.Sprintf("Job %d: job_name must be 'WSRetryJob', got '%s'", i+1, job.JobName.ValueString()),
			)
			return
		}

		// Create the job trigger
		jobTrigger := openapi.NewWSRetryJob(
			job.TriggerName.ValueString(),
			job.JobName.ValueString(),
			job.JobGroup.ValueString(),
			job.CronExpression.ValueString(),
		)

		// Set optional security systems if provided
		if !job.SecuritySystems.IsNull() && len(job.SecuritySystems.Elements()) > 0 {
			var securitySystems []string
			resp.Diagnostics.Append(job.SecuritySystems.ElementsAs(ctx, &securitySystems, false)...)
			if resp.Diagnostics.HasError() {
				return
			}
			jobTrigger.SetSecuritysystems(securitySystems)
		}

		// Set optional task types if provided
		if !job.TaskTypes.IsNull() && job.TaskTypes.ValueString() != "" {
			jobTrigger.SetTasktypes(job.TaskTypes.ValueString())
		}

		// Set optional trigger group if provided
		if !job.TriggerGroup.IsNull() && job.TriggerGroup.ValueString() != "" {
			jobTrigger.SetTriggergroup(job.TriggerGroup.ValueString())
		}

		triggers = append(triggers, openapi.WSRetryJobAsTriggerItem(jobTrigger))
	}

	// Create the request
	updateReq := openapi.CreateOrUpdateTriggersRequest{
		Triggers: triggers,
	}

	// Make the API call
	var apiResp *openapi.CreateOrUpdateTriggersResponse
	var finalHttpResp *http.Response
	err := r.provider.AuthenticatedAPICallWithRetry(ctx, "update_ws_retry_jobs", func(token string) error {
		jobOps := r.jobControlFactory.CreateJobControlOperations(r.client.APIBaseURL(), token)
		apiResponse, httpResp, err := jobOps.CreateOrUpdateTriggers(ctx, updateReq)
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
		resp.Diagnostics.AddError(
			"API Call Error",
			fmt.Sprintf("Error during API call to update WS Retry Job triggers: %s", err.Error()),
		)
		return
	}

	// Handle HTTP errors
	var diags diag.Diagnostics
	if jobcontrolutil.JobControlHandleHTTPError(ctx, finalHttpResp, err, "updating WS Retry Job triggers", &diags) {
		tflog.Error(ctx, "Failed to update WS Retry Job triggers", map[string]interface{}{
			"error": fmt.Sprintf("%v", diags.Errors()),
		})
		for _, diagnostic := range diags {
			if diagnostic.Severity() == diag.SeverityError {
				resp.Diagnostics.AddError(diagnostic.Summary(), diagnostic.Detail())
			}
		}
		return
	}

	// Handle API response errors
	if apiResp != nil && jobcontrolutil.JobControlHandleAPIError(ctx, &apiResp.ErrorCode, &apiResp.Msg, "updating WS Retry Job triggers", &diags) {
		tflog.Error(ctx, "API error during trigger update", map[string]interface{}{
			"error": fmt.Sprintf("%v", diags.Errors()),
		})
		for _, diagnostic := range diags {
			if diagnostic.Severity() == diag.SeverityError {
				resp.Diagnostics.AddError(diagnostic.Summary(), diagnostic.Detail())
			}
		}
		return
	}

	tflog.Info(ctx, "WS Retry Job triggers updated successfully", map[string]interface{}{
		"job_count": len(jobs),
	})

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *WSRetryJobResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state WSRetryJobResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Extract jobs from the state
	var jobs []WSRetryJobModel
	resp.Diagnostics.Append(state.Jobs.ElementsAs(ctx, &jobs, false)...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Starting WS Retry Job triggers deletion", map[string]interface{}{
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
		err := r.provider.AuthenticatedAPICallWithRetry(ctx, "delete_ws_retry_job", func(token string) error {
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
			resp.Diagnostics.AddError(
				"API Call Error",
				fmt.Sprintf("Error during API call to delete WS Retry Job trigger '%s': %s", triggerName, err.Error()),
			)
			return
		}

		// Handle HTTP errors
		var diags diag.Diagnostics
		if jobcontrolutil.JobControlHandleHTTPError(ctx, finalHttpResp, err, fmt.Sprintf("deleting WS Retry Job trigger '%s'", triggerName), &diags) {
			tflog.Error(ctx, "Failed to delete WS Retry Job trigger", map[string]interface{}{
				"error":        fmt.Sprintf("%v", diags.Errors()),
				"trigger_name": triggerName,
			})
			for _, diagnostic := range diags {
				if diagnostic.Severity() == diag.SeverityError {
					resp.Diagnostics.AddError(diagnostic.Summary(), diagnostic.Detail())
				}
			}
			return
		}

		// Handle API response errors
		if apiResp != nil {
			errorCodeStr := fmt.Sprintf("%d", apiResp.ErrorCode)
			if jobcontrolutil.JobControlHandleAPIError(ctx, &errorCodeStr, &apiResp.Msg, fmt.Sprintf("deleting WS Retry Job trigger '%s'", triggerName), &diags) {
				tflog.Error(ctx, "API error during trigger deletion", map[string]interface{}{
					"error":        fmt.Sprintf("%v", diags.Errors()),
					"trigger_name": triggerName,
				})
				for _, diagnostic := range diags {
					if diagnostic.Severity() == diag.SeverityError {
						resp.Diagnostics.AddError(diagnostic.Summary(), diagnostic.Detail())
					}
				}
				return
			}
		}

		tflog.Info(ctx, "WS Retry Job trigger deleted successfully", map[string]interface{}{
			"trigger_name": triggerName,
		})
	}

	tflog.Info(ctx, "All WS Retry Job triggers deleted successfully", map[string]interface{}{
		"job_count": len(jobs),
	})
}
