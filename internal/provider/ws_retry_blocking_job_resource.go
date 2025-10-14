// Copyright (c) 2025 Saviynt Inc.
// SPDX-License-Identifier: MPL-2.0

// ws_retry_blocking_job_resource.go manages WS Blocking Retry Job triggers in Saviynt.
// The resource implements the full Terraform lifecycle:
//   - Create: provisions a new WS blocking retry job trigger using the supplied configuration.
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
var _ resource.Resource = &WSRetryBlockingJobResource{}

func NewWSRetryBlockingJobResource() resource.Resource {
	return &WSRetryBlockingJobResource{}
}

// WSRetryBlockingJobResource defines the resource implementation.
type WSRetryBlockingJobResource struct {
	client            client.SaviyntClientInterface
	provider          client.SaviyntProviderInterface
	jobControlFactory client.JobControlFactoryInterface
	token             string
}

// WSRetryBlockingJobResourceModel describes the resource data model.
type WSRetryBlockingJobResourceModel struct {
	Jobs types.List `tfsdk:"jobs"`
}

// WSRetryBlockingJobModel describes individual job data model.
type WSRetryBlockingJobModel struct {
	BaseJobControlResourceModel
	SecuritySystems types.List   `tfsdk:"security_systems"`
	TaskTypes       types.String `tfsdk:"task_types"`
}

func WSRetryBlockingJobResourceSchema() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"security_systems": schema.ListAttribute{
			Optional:    true,
			ElementType: types.StringType,
			Description: "List of security systems for the WS blocking retry job",
		},
		"task_types": schema.StringAttribute{
			Optional:    true,
			Description: "Task types for the WS blocking retry job",
		},
	}
}

func (r *WSRetryBlockingJobResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: util.WSRetryJobDescription,
		Attributes: map[string]schema.Attribute{
			"jobs": schema.ListNestedAttribute{
				Required:    true,
				Description: "List of WS Blocking Retry Jobs to create",
				NestedObject: schema.NestedAttributeObject{
					Attributes: jobcontrolutil.MergeJobResourceAttributes(
						BaseJobControlResourceSchema(),
						WSRetryBlockingJobResourceSchema(),
					),
				},
			},
		},
	}
}

func (r *WSRetryBlockingJobResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_ws_retry_blocking_job_resource"
}

func (r *WSRetryBlockingJobResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	tflog.Debug(ctx, "Starting WSRetryBlockingJobResource configuration")

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

	tflog.Info(ctx, "WSRetryBlockingJobResource configuration completed successfully")
}

// SetClient sets the client for testing purposes
func (r *WSRetryBlockingJobResource) SetClient(client client.SaviyntClientInterface) {
	r.client = client
}

// SetToken sets the token for testing purposes
func (r *WSRetryBlockingJobResource) SetToken(token string) {
	r.token = token
}

// SetProvider sets the provider for testing purposes
func (r *WSRetryBlockingJobResource) SetProvider(provider client.SaviyntProviderInterface) {
	r.provider = provider
}

// SetJobControlFactory sets the job control factory for testing purposes
func (r *WSRetryBlockingJobResource) SetJobControlFactory(factory client.JobControlFactoryInterface) {
	r.jobControlFactory = factory
}

// CreateOrUpdateWSRetryBlockingJobs handles the business logic for creating or updating WS Retry Blocking jobs
func (r *WSRetryBlockingJobResource) CreateOrUpdateWSRetryBlockingJobs(ctx context.Context, jobs []WSRetryBlockingJobModel, operation string) (*openapi.CreateOrUpdateTriggersResponse, error) {
	if len(jobs) == 0 {
		return nil, fmt.Errorf("at least one job must be specified")
	}

	tflog.Debug(ctx, "Starting WS Blocking Retry Job triggers creation", map[string]interface{}{
		"job_count": len(jobs),
	})

	var triggers []openapi.TriggerItem

	// Process each job
	for i, job := range jobs {
		// Validate job name is WSBlockingRetryJob
		if job.JobName.IsNull() || job.JobName.ValueString() != "WSBlockingRetryJob" {
			return nil, fmt.Errorf("job %d: job_name must be 'WSBlockingRetryJob', got '%s'", i+1, job.JobName.ValueString())
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

		// Create the job trigger using WSRetryJob (since WSBlockingRetryJob API model doesn't exist)
		jobTrigger := openapi.NewWSRetryJob(
			job.TriggerName.ValueString(),
			job.JobName.ValueString(),
			job.JobGroup.ValueString(),
			job.CronExpression.ValueString(),
		)

		// Set optional security systems if provided
		if !job.SecuritySystems.IsNull() && len(job.SecuritySystems.Elements()) > 0 {
			var securitySystems []string
			diags := job.SecuritySystems.ElementsAs(ctx, &securitySystems, false)
			if diags.HasError() {
				return nil, fmt.Errorf("job %d: failed to extract security systems", i+1)
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
	err := r.provider.AuthenticatedAPICallWithRetry(ctx, fmt.Sprintf("%s_ws_blocking_retry_jobs", operation), func(token string) error {
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
	if jobcontrolutil.JobControlHandleHTTPError(ctx, finalHttpResp, err, "creating WS Blocking Retry Job triggers", &diags) {
		tflog.Error(ctx, "Failed to create WS Blocking Retry Job triggers", map[string]interface{}{
			"error": fmt.Sprintf("%v", diags.Errors()),
		})
		if diags.HasError() {
			return nil, fmt.Errorf("HTTP error: %s", diags.Errors()[0].Detail())
		}
	}

	// Handle API response errors
	if apiResp != nil && jobcontrolutil.JobControlHandleAPIError(ctx, &apiResp.ErrorCode, &apiResp.Msg, "creating WS Blocking Retry Job triggers", &diags) {
		tflog.Error(ctx, "API error during trigger creation", map[string]interface{}{
			"error": fmt.Sprintf("%v", diags.Errors()),
		})
		if diags.HasError() {
			return nil, fmt.Errorf("API error: %s", diags.Errors()[0].Detail())
		}
	}

	tflog.Info(ctx, "WS Blocking Retry Job triggers created successfully", map[string]interface{}{
		"job_count": len(jobs),
	})

	return apiResp, nil
}

// DeleteWSRetryBlockingJobs handles the business logic for deleting WS Retry Blocking jobs
func (r *WSRetryBlockingJobResource) DeleteWSRetryBlockingJobs(ctx context.Context, jobs []WSRetryBlockingJobModel) error {
	if len(jobs) == 0 {
		return nil
	}

	tflog.Debug(ctx, "Starting WS Blocking Retry Job triggers deletion", map[string]interface{}{
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
		err := r.provider.AuthenticatedAPICallWithRetry(ctx, "delete_ws_blocking_retry_job", func(token string) error {
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
		if jobcontrolutil.JobControlHandleHTTPError(ctx, finalHttpResp, err, fmt.Sprintf("deleting WS Blocking Retry Job trigger '%s'", triggerName), &diags) {
			tflog.Error(ctx, "Failed to delete WS Blocking Retry Job trigger", map[string]interface{}{
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
			if jobcontrolutil.JobControlHandleAPIError(ctx, &errorCodeStr, &apiResp.Msg, fmt.Sprintf("deleting WS Blocking Retry Job trigger '%s'", triggerName), &diags) {
				tflog.Error(ctx, "API error during trigger deletion", map[string]interface{}{
					"error":        fmt.Sprintf("%v", diags.Errors()),
					"trigger_name": triggerName,
				})
				if diags.HasError() {
					return fmt.Errorf("API error for trigger '%s': %s", triggerName, diags.Errors()[0].Detail())
				}
			}
		}

		tflog.Info(ctx, "WS Blocking Retry Job trigger deleted successfully", map[string]interface{}{
			"trigger_name": triggerName,
		})
	}

	tflog.Info(ctx, "All WS Blocking Retry Job triggers deleted successfully", map[string]interface{}{
		"job_count": len(jobs),
	})

	return nil
}

func (r *WSRetryBlockingJobResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan WSRetryBlockingJobResourceModel

	tflog.Debug(ctx, "Starting WS Blocking Retry Job resource creation")

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
	var jobs []WSRetryBlockingJobModel
	resp.Diagnostics.Append(plan.Jobs.ElementsAs(ctx, &jobs, false)...)
	if resp.Diagnostics.HasError() {
		resp.Diagnostics.AddError(
			"Jobs Extraction Failed",
			"Unable to extract jobs from Terraform plan",
		)
		return
	}

	// Call the business logic method
	_, err := r.CreateOrUpdateWSRetryBlockingJobs(ctx, jobs, "create")
	if err != nil {
		tflog.Error(ctx, "WS Blocking Retry Job creation failed", map[string]interface{}{
			"error": err.Error(),
		})
		resp.Diagnostics.AddError(
			"WS Blocking Retry Job Creation Failed",
			err.Error(),
		)
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		resp.Diagnostics.AddError(
			"State Update Failed",
			"Unable to save WS Blocking Retry Job state",
		)
		return
	}

	tflog.Info(ctx, "WS Blocking Retry Job resource created successfully", map[string]interface{}{
		"job_count": len(jobs),
	})
}

func (r *WSRetryBlockingJobResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state WSRetryBlockingJobResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading WS Blocking Retry Job triggers")

	// For now, we'll keep the state as-is since the API doesn't provide a direct read method
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *WSRetryBlockingJobResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan WSRetryBlockingJobResourceModel

	tflog.Debug(ctx, "Starting WS Blocking Retry Job resource update")

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
	var jobs []WSRetryBlockingJobModel
	resp.Diagnostics.Append(plan.Jobs.ElementsAs(ctx, &jobs, false)...)
	if resp.Diagnostics.HasError() {
		resp.Diagnostics.AddError(
			"Jobs Extraction Failed",
			"Unable to extract jobs from Terraform plan",
		)
		return
	}

	// Call the business logic method
	_, err := r.CreateOrUpdateWSRetryBlockingJobs(ctx, jobs, "update")
	if err != nil {
		tflog.Error(ctx, "WS Blocking Retry Job update failed", map[string]interface{}{
			"error": err.Error(),
		})
		resp.Diagnostics.AddError(
			"WS Blocking Retry Job Update Failed",
			err.Error(),
		)
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		resp.Diagnostics.AddError(
			"State Update Failed",
			"Unable to save WS Blocking Retry Job state",
		)
		return
	}

	tflog.Info(ctx, "WS Blocking Retry Job resource updated successfully", map[string]interface{}{
		"job_count": len(jobs),
	})
}

func (r *WSRetryBlockingJobResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state WSRetryBlockingJobResourceModel

	tflog.Debug(ctx, "Starting WS Blocking Retry Job resource deletion")

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
	var jobs []WSRetryBlockingJobModel
	resp.Diagnostics.Append(state.Jobs.ElementsAs(ctx, &jobs, false)...)
	if resp.Diagnostics.HasError() {
		resp.Diagnostics.AddError(
			"Jobs Extraction Failed",
			"Unable to extract jobs from Terraform state",
		)
		return
	}

	// Call the business logic method
	err := r.DeleteWSRetryBlockingJobs(ctx, jobs)
	if err != nil {
		tflog.Error(ctx, "WS Blocking Retry Job deletion failed", map[string]interface{}{
			"error": err.Error(),
		})
		resp.Diagnostics.AddError(
			"WS Blocking Retry Job Deletion Failed",
			err.Error(),
		)
		return
	}

	tflog.Info(ctx, "WS Blocking Retry Job resource deleted successfully", map[string]interface{}{
		"job_count": len(jobs),
	})
}
