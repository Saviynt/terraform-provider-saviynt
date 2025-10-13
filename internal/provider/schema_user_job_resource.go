// Copyright (c) 2025 Saviynt Inc.
// SPDX-License-Identifier: MPL-2.0

// schema_user_job_resource.go manages Schema User Job triggers in Saviynt.
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

var _ resource.Resource = &SchemaUserJobResource{}

func NewSchemaUserJobResource() resource.Resource {
	return &SchemaUserJobResource{}
}

type SchemaUserJobResource struct {
	client            client.SaviyntClientInterface
	provider          client.SaviyntProviderInterface
	jobControlFactory client.JobControlFactoryInterface
	token             string
}

type SchemaUserJobResourceModel struct {
	Jobs types.List `tfsdk:"jobs"`
}

// SchemaUserJobModel describes individual job data model.
type SchemaUserJobModel struct {
	BaseJobTriggerResourceModel
	SchemaFileNames types.String `tfsdk:"schema_file_names"`
}

func SchemaUserJobResourceSchema() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"jobs": schema.ListNestedAttribute{
			Required:    true,
			Description: "List of Schema User Job configurations",
			NestedObject: schema.NestedAttributeObject{
				Attributes: jobcontrolutil.MergeJobResourceAttributes(
					BaseJobTriggerResourceSchema(),
					map[string]schema.Attribute{
						"schema_file_names": schema.StringAttribute{
							Optional:    true,
							Description: "Schema file names for the user job",
						},
					},
				),
			},
		},
	}
}

func (r *SchemaUserJobResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: util.SchemaUserJobDescription,
		Attributes:          SchemaUserJobResourceSchema(),
	}
}

func (r *SchemaUserJobResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_schema_user_job_resource"
}

func (r *SchemaUserJobResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	tflog.Debug(ctx, "Starting SchemaUserJobResource configuration")

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
		resp.Diagnostics.AddError("Unexpected Resource Configure Type", errMsg)
		return
	}

	if prov.client == nil {
		tflog.Error(ctx, "Provider client is nil", map[string]interface{}{
			"error": "SaviyntProvider.client is nil",
		})
		resp.Diagnostics.AddError("Provider Configuration Error", "Provider client is not initialized")
		return
	}

	if prov.accessToken == "" {
		tflog.Error(ctx, "Access token is empty", map[string]interface{}{
			"error": "SaviyntProvider.accessToken is empty",
		})
		resp.Diagnostics.AddError("Provider Authentication Error", "Access token is not available")
		return
	}

	r.provider = &client.SaviyntProviderWrapper{Provider: prov}
	r.client = &client.SaviyntClientWrapper{Client: prov.client}
	r.jobControlFactory = &client.DefaultJobControlFactory{}
	r.token = prov.accessToken

	tflog.Info(ctx, "SchemaUserJobResource configuration completed successfully")
}

func (r *SchemaUserJobResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan SchemaUserJobResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var jobs []SchemaUserJobModel
	resp.Diagnostics.Append(plan.Jobs.ElementsAs(ctx, &jobs, false)...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Starting Schema User Job triggers creation", map[string]interface{}{
		"job_count": len(jobs),
	})

	var jobTriggerItems []openapi.JobTriggerItem

	for i, job := range jobs {
		// Validate job name
		if job.JobName.ValueString() != "SchemaUserJob" {
			resp.Diagnostics.AddError(
				"Invalid Job Name",
				fmt.Sprintf("Job %d: job_name must be 'SchemaUserJob', got '%s'", i, job.JobName.ValueString()),
			)
			return
		}

		// Validate required fields
		if job.Name.IsNull() || job.Name.ValueString() == "" {
			resp.Diagnostics.AddError("Validation Error", fmt.Sprintf("Job %d: name is required", i))
			return
		}
		if job.JobGroup.IsNull() || job.JobGroup.ValueString() == "" {
			resp.Diagnostics.AddError("Validation Error", fmt.Sprintf("Job %d: job_group is required", i))
			return
		}
		if job.Group.IsNull() || job.Group.ValueString() == "" {
			resp.Diagnostics.AddError("Validation Error", fmt.Sprintf("Job %d: group is required", i))
			return
		}
		if job.CronExp.IsNull() || job.CronExp.ValueString() == "" {
			resp.Diagnostics.AddError("Validation Error", fmt.Sprintf("Job %d: cron_exp is required", i))
			return
		}

		valueMap := openapi.NewSchemaUserJobAllOfValueMap()
		if !job.SchemaFileNames.IsNull() && job.SchemaFileNames.ValueString() != "" {
			valueMap.SetSchemaFileNames(job.SchemaFileNames.ValueString())
		}

		jobTrigger := openapi.NewSchemaUserJob(
			job.Name.ValueString(),
			job.JobName.ValueString(),
			job.JobGroup.ValueString(),
			job.Group.ValueString(),
			job.CronExp.ValueString(),
		)
		jobTrigger.SetValueMap(*valueMap)

		jobTriggerItem := openapi.SchemaUserJobAsJobTriggerItem(jobTrigger)
		jobTriggerItems = append(jobTriggerItems, jobTriggerItem)
	}

	createReq := []openapi.JobTriggerRequest{{Triggers: jobTriggerItems}}

	var apiResp *openapi.CreateTriggersResponse
	var finalHttpResp *http.Response
	err := r.provider.AuthenticatedAPICallWithRetry(ctx, "create_schema_user_jobs", func(token string) error {
		jobOps := r.jobControlFactory.CreateJobControlOperations(r.client.APIBaseURL(), token)
		apiResponse, httpResp, err := jobOps.CreateTrigger(ctx, createReq)
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
		resp.Diagnostics.AddError("API Call Error", fmt.Sprintf("Error creating Schema User Job triggers: %s", err.Error()))
		return
	}

	if apiResp != nil {
		if apiResp.CreateTriggersResponseOneOf != nil && apiResp.CreateTriggersResponseOneOf.Error != "" {
			tflog.Error(ctx, "API error during trigger creation", map[string]interface{}{
				"error": apiResp.CreateTriggersResponseOneOf.Error,
			})
			resp.Diagnostics.AddError("API Error", fmt.Sprintf("Error creating Schema User Job triggers: %s", apiResp.CreateTriggersResponseOneOf.Error))
			return
		}

		if apiResp.MapmapOfStringarrayOfString != nil {
			tflog.Debug(ctx, "Received success response from CreateTrigger API")
		}
	}

	tflog.Info(ctx, "Schema User Job triggers created successfully", map[string]interface{}{
		"job_count": len(jobs),
	})

	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *SchemaUserJobResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state SchemaUserJobResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading Schema User Job triggers")

	// For now, we'll keep the state as-is since the API doesn't provide a direct read method
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *SchemaUserJobResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan SchemaUserJobResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var jobs []SchemaUserJobModel
	resp.Diagnostics.Append(plan.Jobs.ElementsAs(ctx, &jobs, false)...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Starting Schema User Job triggers update", map[string]interface{}{
		"job_count": len(jobs),
	})

	var jobTriggerItems []openapi.JobTriggerItem

	for i, job := range jobs {
		// Validate job name
		if job.JobName.ValueString() != "SchemaUserJob" {
			resp.Diagnostics.AddError(
				"Invalid Job Name",
				fmt.Sprintf("Job %d: job_name must be 'SchemaUserJob', got '%s'", i, job.JobName.ValueString()),
			)
			return
		}

		valueMap := openapi.NewSchemaUserJobAllOfValueMap()
		if !job.SchemaFileNames.IsNull() && job.SchemaFileNames.ValueString() != "" {
			valueMap.SetSchemaFileNames(job.SchemaFileNames.ValueString())
		}

		jobTrigger := openapi.NewSchemaUserJob(
			job.Name.ValueString(),
			job.JobName.ValueString(),
			job.JobGroup.ValueString(),
			job.Group.ValueString(),
			job.CronExp.ValueString(),
		)
		jobTrigger.SetValueMap(*valueMap)

		jobTriggerItem := openapi.SchemaUserJobAsJobTriggerItem(jobTrigger)
		jobTriggerItems = append(jobTriggerItems, jobTriggerItem)
	}

	updateReq := []openapi.JobTriggerRequest{{Triggers: jobTriggerItems}}

	var apiResp *openapi.CreateTriggersResponse
	var finalHttpResp *http.Response
	err := r.provider.AuthenticatedAPICallWithRetry(ctx, "update_schema_user_jobs", func(token string) error {
		jobOps := r.jobControlFactory.CreateJobControlOperations(r.client.APIBaseURL(), token)
		apiResponse, httpResp, err := jobOps.CreateTrigger(ctx, updateReq)
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
		resp.Diagnostics.AddError("API Call Error", fmt.Sprintf("Error updating Schema User Job triggers: %s", err.Error()))
		return
	}

	if apiResp != nil {
		if apiResp.CreateTriggersResponseOneOf != nil && apiResp.CreateTriggersResponseOneOf.Error != "" {
			tflog.Error(ctx, "API error during trigger update", map[string]interface{}{
				"error": apiResp.CreateTriggersResponseOneOf.Error,
			})
			resp.Diagnostics.AddError("API Error", fmt.Sprintf("Error updating Schema User Job triggers: %s", apiResp.CreateTriggersResponseOneOf.Error))
			return
		}

		if apiResp.MapmapOfStringarrayOfString != nil {
			tflog.Debug(ctx, "Received success response from CreateTrigger API")
		}
	}

	tflog.Info(ctx, "Schema User Job triggers updated successfully", map[string]interface{}{
		"job_count": len(jobs),
	})

	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *SchemaUserJobResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state SchemaUserJobResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var jobs []SchemaUserJobModel
	resp.Diagnostics.Append(state.Jobs.ElementsAs(ctx, &jobs, false)...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Starting Schema User Job triggers deletion", map[string]interface{}{
		"job_count": len(jobs),
	})

	// Delete each job individually since DeleteTrigger API doesn't support bulk operations
	for _, job := range jobs {
		name := job.Name.ValueString()
		jobName := job.JobName.ValueString()
		jobGroup := job.JobGroup.ValueString()

		deleteReq := openapi.DeleteTriggerRequest{
			Triggername: name,
			Jobname:     jobName,
			Jobgroup:    jobGroup,
		}

		// Make the API call
		var apiResp *openapi.DeleteTriggerResponse
		var finalHttpResp *http.Response
		err := r.provider.AuthenticatedAPICallWithRetry(ctx, "delete_schema_user_job", func(token string) error {
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
				"trigger_name": name,
			})
			resp.Diagnostics.AddError(
				"API Call Error",
				fmt.Sprintf("Error during API call to delete Schema User Job trigger '%s': %s", name, err.Error()),
			)
			return
		}

		// Handle HTTP errors using job control error handler
		var diags diag.Diagnostics
		if jobcontrolutil.JobControlHandleHTTPError(ctx, finalHttpResp, err, fmt.Sprintf("deleting Schema User Job trigger '%s'", name), &diags) {
			tflog.Error(ctx, "Failed to delete Schema User Job trigger", map[string]interface{}{
				"error":        fmt.Sprintf("%v", diags.Errors()),
				"trigger_name": name,
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
			if jobcontrolutil.JobControlHandleAPIError(ctx, &errorCodeStr, &apiResp.Msg, fmt.Sprintf("deleting Schema User Job trigger '%s'", name), &diags) {
				tflog.Error(ctx, "API error during trigger deletion", map[string]interface{}{
					"error":        fmt.Sprintf("%v", diags.Errors()),
					"trigger_name": name,
				})
				for _, diagnostic := range diags {
					if diagnostic.Severity() == diag.SeverityError {
						resp.Diagnostics.AddError(diagnostic.Summary(), diagnostic.Detail())
					}
				}
				return
			}
		}

		tflog.Info(ctx, "Schema User Job trigger deleted successfully", map[string]interface{}{
			"trigger_name": name,
		})
	}

	tflog.Info(ctx, "All Schema User Job triggers deleted successfully", map[string]interface{}{
		"job_count": len(jobs),
	})
}
