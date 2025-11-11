// Copyright (c) 2025 Saviynt Inc.
// SPDX-License-Identifier: MPL-2.0

// saviynt_job_control_resource manages job running operations in the Saviynt Security Manager.
// The resource implements the full Terraform lifecycle:
//   - Create: executes run operations on specified jobs.
//   - Read: returns current state (stateless operations).
//   - Update: applies any configuration changes to job control operations.
//   - Delete: no-op as job control operations are stateless.
package provider

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
	"terraform-provider-Saviynt/internal/client"
	"terraform-provider-Saviynt/util"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	openapi "github.com/saviynt/saviynt-api-go-client/job_control"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &JobControlResource{}

func NewJobControlResource() resource.Resource {
	return &JobControlResource{}
}

// JobControlResource defines the resource implementation.
type JobControlResource struct {
	client            client.SaviyntClientInterface
	provider          client.SaviyntProviderInterface
	jobControlFactory client.JobControlFactoryInterface
	token             string
}

// JobControlResourceModel describes the resource data model.
type JobControlResourceModel struct {
	ID          types.String `tfsdk:"id"`
	RunJobs     types.Set    `tfsdk:"run_jobs"`     // List of job objects to run
	RunMessages types.Map    `tfsdk:"run_messages"` // Map of trigger name to run response message
}

// RunJobModel describes the structure for run job operations
type RunJobModel struct {
	JobName       types.String `tfsdk:"job_name"`
	TriggerName   types.String `tfsdk:"trigger_name"`
	JobGroup      types.String `tfsdk:"job_group"`
	RunJobVersion types.String `tfsdk:"run_job_version"`
}

func (r *JobControlResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_job_control_resource"
}

func (r *JobControlResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: util.JobControlDescription,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Job control resource identifier",
			},
			"run_jobs": schema.SetNestedAttribute{
				Optional:            true,
				MarkdownDescription: "List of job objects to run",
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"job_name": schema.StringAttribute{
							Required:            true,
							MarkdownDescription: "Name of the job to run",
						},
						"trigger_name": schema.StringAttribute{
							Required:            true,
							MarkdownDescription: "Trigger name for the job",
						},
						"job_group": schema.StringAttribute{
							Required:            true,
							MarkdownDescription: "Job group for the job",
						},
						"run_job_version": schema.StringAttribute{
							Optional:            true,
							MarkdownDescription: "Version identifier to trigger job run. Change this value to trigger re-run of the job",
						},
					},
				},
			},
			"run_messages": schema.MapAttribute{
				ElementType:         types.StringType,
				Computed:            true,
				MarkdownDescription: "Map of trigger names to run operation response messages",
			},
		},
	}
}

func (r *JobControlResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	tflog.Debug(ctx, "Starting JobControlResource configuration")

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

	tflog.Info(ctx, "JobControlResource configuration completed successfully")
}

// SetClient sets the client for testing purposes
func (r *JobControlResource) SetClient(client client.SaviyntClientInterface) {
	r.client = client
}

// SetToken sets the token for testing purposes
func (r *JobControlResource) SetToken(token string) {
	r.token = token
}

// SetProvider sets the provider for testing purposes
func (r *JobControlResource) SetProvider(provider client.SaviyntProviderInterface) {
	r.provider = provider
}

// SetJobControlFactory sets the job control factory for testing purposes
func (r *JobControlResource) SetJobControlFactory(factory client.JobControlFactoryInterface) {
	r.jobControlFactory = factory
}

// ExecuteJobControlOperations handles the business logic for running jobs
func (r *JobControlResource) ExecuteJobControlOperations(ctx context.Context, plan *JobControlResourceModel, operation string) error {
	tflog.Debug(ctx, "Starting job control operations", map[string]interface{}{
		"operation": operation,
	})

	// Maps to store results and errors
	runMessages := make(map[string]string)
	var errors []string

	// Process run jobs
	if !plan.RunJobs.IsNull() && !plan.RunJobs.IsUnknown() {
		var runJobsList []RunJobModel
		if err := plan.RunJobs.ElementsAs(ctx, &runJobsList, false); err != nil {
			return fmt.Errorf("failed to extract run jobs: %v", err)
		}

		tflog.Info(ctx, "Processing run jobs", map[string]interface{}{
			"job_count": len(runJobsList),
		})

		for _, job := range runJobsList {
			triggerName := job.TriggerName.ValueString()
			message, err := r.RunJob(ctx, job)
			if err != nil {
				errors = append(errors, fmt.Sprintf("Run job %s: %s", triggerName, err.Error()))
				runMessages[triggerName] = fmt.Sprintf("ERROR: %s", err.Error())
			} else {
				runMessages[triggerName] = message
			}
		}
	}

	// Set response messages in state (always set to known values)
	runMap, diags := types.MapValueFrom(context.Background(), types.StringType, runMessages)
	if diags.HasError() {
		return fmt.Errorf("failed to create run messages map")
	}
	plan.RunMessages = runMap

	// Set ID
	plan.ID = types.StringValue("job-control")

	// Return combined error if any failures occurred
	if len(errors) > 0 {
		return fmt.Errorf("job control operations failed: %s", strings.Join(errors, "; "))
	}

	return nil
}

func (r *JobControlResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan JobControlResourceModel

	tflog.Debug(ctx, "Starting job control resource creation")

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
	err := r.ExecuteJobControlOperations(ctx, &plan, "create")

	// Save data into Terraform state first
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Show errors after state is set
	if err != nil {
		tflog.Error(ctx, "Job control creation failed", map[string]interface{}{
			"error": err.Error(),
		})
		resp.Diagnostics.AddError(
			"Job Control Creation Failed",
			err.Error(),
		)
	} else {
		// Add warning about job execution with run messages
		runMessagesStr := ""
		if !plan.RunMessages.IsNull() && !plan.RunMessages.IsUnknown() {
			runMessagesMap := make(map[string]string)
			plan.RunMessages.ElementsAs(ctx, &runMessagesMap, false)
			for triggerName, message := range runMessagesMap {
				runMessagesStr += fmt.Sprintf("\n- %s: %s", triggerName, message)
			}
		}

		resp.Diagnostics.AddWarning(
			"Job Control Operations Executed",
			fmt.Sprintf("Job control operations have been executed. Please check the Saviynt UI for job execution status and results.%s", runMessagesStr),
		)
	}

	tflog.Info(ctx, "Job control resource created successfully")
}

// RunJob runs a specific job and returns the response message
func (r *JobControlResource) RunJob(ctx context.Context, job RunJobModel) (string, error) {
	tflog.Debug(ctx, "Running job", map[string]interface{}{
		"job_name":     job.JobName.ValueString(),
		"trigger_name": job.TriggerName.ValueString(),
		"job_group":    job.JobGroup.ValueString(),
	})

	request := openapi.NewRunJobTriggerRequest(
		job.JobName.ValueString(),
		job.TriggerName.ValueString(),
		job.JobGroup.ValueString(),
	)

	var message string
	var apiResp *openapi.RunJobTriggerResponse
	var finalHttpResp *http.Response

	err := r.provider.AuthenticatedAPICallWithRetry(ctx, "run_job_trigger", func(token string) error {
		jobOps := r.jobControlFactory.CreateJobControlOperations(r.client.APIBaseURL(), token)
		resp, httpResp, err := jobOps.RunJobTrigger(ctx, *request)

		if httpResp != nil && httpResp.StatusCode == 401 {
			return fmt.Errorf("401 unauthorized")
		}

		apiResp = resp
		finalHttpResp = httpResp
		return err
	})

	// Handle non-412 errors first
	if err != nil && finalHttpResp != nil && finalHttpResp.StatusCode != 412 {
		return "", fmt.Errorf("API call failed: %w", err)
	}

	// Handle 412 precondition failed with response body decoding
	if err != nil && finalHttpResp != nil && finalHttpResp.StatusCode == 412 {
		var errorResp struct {
			Msg       string `json:"msg"`
			ErrorCode string `json:"errorCode"`
		}

		if decodeErr := json.NewDecoder(finalHttpResp.Body).Decode(&errorResp); decodeErr == nil {
			return "", fmt.Errorf("precondition failed - ErrorCode: %s, Msg: %s", errorResp.ErrorCode, errorResp.Msg)
		}
		return "", fmt.Errorf("precondition failed: some required condition was not met")
	}

	if err != nil {
		return "", fmt.Errorf("API call failed: %w", err)
	}

	if apiResp != nil {
		message = apiResp.Msg
	}

	tflog.Debug(ctx, "Run job completed", map[string]interface{}{
		"job_name": job.JobName.ValueString(),
		"message":  message,
	})

	return message, nil
}

func (r *JobControlResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state JobControlResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Job control is stateless - just return current state
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *JobControlResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan JobControlResourceModel

	tflog.Debug(ctx, "Starting job control resource update")

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
	err := r.ExecuteJobControlOperations(ctx, &plan, "update")

	// Save data into Terraform state first
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Show errors after state is set
	if err != nil {
		tflog.Error(ctx, "Job control update failed", map[string]interface{}{
			"error": err.Error(),
		})
		resp.Diagnostics.AddError(
			"Job Control Update Failed",
			err.Error(),
		)
	} else {
		// Add warning about job execution with run messages
		runMessagesStr := ""
		if !plan.RunMessages.IsNull() && !plan.RunMessages.IsUnknown() {
			runMessagesMap := make(map[string]string)
			plan.RunMessages.ElementsAs(ctx, &runMessagesMap, false)
			for triggerName, message := range runMessagesMap {
				runMessagesStr += fmt.Sprintf("\n- %s: %s", triggerName, message)
			}
		}

		resp.Diagnostics.AddWarning(
			"Job Control Operations Executed",
			fmt.Sprintf("Job control operations have been executed. Please check the Saviynt UI for job execution status and results.%s", runMessagesStr),
		)
	}

	tflog.Info(ctx, "Job control resource updated successfully")
}

func (r *JobControlResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	tflog.Debug(ctx, "Starting job control resource deletion")

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
