// Copyright (c) 2025 Saviynt Inc.
// SPDX-License-Identifier: MPL-2.0

// file_transfer_job_resource.go manages File Transfer Job triggers in Saviynt.
package provider

import (
	"context"
	"fmt"
	"net/http"
	"terraform-provider-Saviynt/internal/client"
	"terraform-provider-Saviynt/util"
	"terraform-provider-Saviynt/util/jobcontrolutil"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	openapi "github.com/saviynt/saviynt-api-go-client/job_control"
)

var _ resource.Resource = &FileTransferJobResource{}

func NewFileTransferJobResource() resource.Resource {
	return &FileTransferJobResource{}
}

type FileTransferJobResource struct {
	client            client.SaviyntClientInterface
	provider          client.SaviyntProviderInterface
	jobControlFactory client.JobControlFactoryInterface
	token             string
}

type FileTransferJobResourceModel struct {
	Jobs types.List `tfsdk:"jobs"`
}

// FileTransferJobModel describes individual job data model.
type FileTransferJobModel struct {
	BaseJobTriggerResourceModel
	ExternalConnectionKey types.String `tfsdk:"external_connection_key"`
	FileTransferAction    types.String `tfsdk:"file_transfer_action"`
}

func FileTransferJobResourceSchema() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"jobs": schema.ListNestedAttribute{
			Required:    true,
			Description: "List of File Transfer Job configurations",
			NestedObject: schema.NestedAttributeObject{
				Attributes: jobcontrolutil.MergeJobResourceAttributes(
					BaseJobTriggerResourceSchema(),
					map[string]schema.Attribute{
						"external_connection_key": schema.StringAttribute{
							Required:    true,
							Description: "External connection key for file transfer",
						},
						"file_transfer_action": schema.StringAttribute{
							Required:    true,
							Description: "File transfer action (UPLOAD or DOWNLOAD)",
							Validators: []validator.String{
								stringvalidator.OneOf("UPLOAD", "DOWNLOAD"),
							},
						},
					},
				),
			},
		},
	}
}

func (r *FileTransferJobResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: util.FileTransferJobDescription,
		Attributes:          FileTransferJobResourceSchema(),
	}
}

func (r *FileTransferJobResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_file_transfer_job_resource"
}

func (r *FileTransferJobResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	tflog.Debug(ctx, "Starting FileTransferJobResource configuration")

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

	tflog.Info(ctx, "FileTransferJobResource configuration completed successfully")
}

// SetClient sets the client for testing purposes
func (r *FileTransferJobResource) SetClient(client client.SaviyntClientInterface) {
	r.client = client
}

// SetToken sets the token for testing purposes
func (r *FileTransferJobResource) SetToken(token string) {
	r.token = token
}

// SetProvider sets the provider for testing purposes
func (r *FileTransferJobResource) SetProvider(provider client.SaviyntProviderInterface) {
	r.provider = provider
}

// SetJobControlFactory sets the job control factory for testing purposes
func (r *FileTransferJobResource) SetJobControlFactory(factory client.JobControlFactoryInterface) {
	r.jobControlFactory = factory
}

// CreateOrUpdateFileTransferJobs handles the business logic for creating or updating File Transfer jobs
func (r *FileTransferJobResource) CreateOrUpdateFileTransferJobs(ctx context.Context, jobs []FileTransferJobModel, operation string) (*openapi.CreateTriggersResponse, error) {
	if len(jobs) == 0 {
		return nil, fmt.Errorf("at least one job must be specified")
	}

	tflog.Debug(ctx, "Starting File Transfer Job triggers creation", map[string]interface{}{
		"job_count": len(jobs),
	})

	var jobTriggerItems []openapi.JobTriggerItem

	for i, job := range jobs {
		// Validate required fields
		if job.Name.IsNull() || job.Name.ValueString() == "" {
			return nil, fmt.Errorf("job %d: name is required", i+1)
		}
		if job.JobGroup.IsNull() || job.JobGroup.ValueString() == "" {
			return nil, fmt.Errorf("job %d: job_group is required", i+1)
		}
		if job.Group.IsNull() || job.Group.ValueString() == "" {
			return nil, fmt.Errorf("job %d: group is required", i+1)
		}
		if job.CronExp.IsNull() || job.CronExp.ValueString() == "" {
			return nil, fmt.Errorf("job %d: cron_exp is required", i+1)
		}
		if job.ExternalConnectionKey.IsNull() || job.ExternalConnectionKey.ValueString() == "" {
			return nil, fmt.Errorf("job %d: external_connection_key is required", i+1)
		}
		if job.FileTransferAction.IsNull() || job.FileTransferAction.ValueString() == "" {
			return nil, fmt.Errorf("job %d: file_transfer_action is required", i+1)
		}

		// Create the value map
		valueMap := openapi.NewFileTransferJobAllOfValueMap(
			job.ExternalConnectionKey.ValueString(),
			job.FileTransferAction.ValueString(),
		)

		// Create the job trigger
		jobTrigger := openapi.NewFileTransferJob(
			job.Name.ValueString(),
			"FileTransferJob",
			job.JobGroup.ValueString(),
			job.Group.ValueString(),
			job.CronExp.ValueString(),
		)
		jobTrigger.SetValueMap(*valueMap)

		// Create job trigger item
		jobTriggerItem := openapi.FileTransferJobAsJobTriggerItem(jobTrigger)
		jobTriggerItems = append(jobTriggerItems, jobTriggerItem)
	}

	// Create the request
	createReq := []openapi.JobTriggerRequest{
		{
			Triggers: jobTriggerItems,
		},
	}

	// Make the API call
	var apiResp *openapi.CreateTriggersResponse
	var finalHttpResp *http.Response
	err := r.provider.AuthenticatedAPICallWithRetry(ctx, fmt.Sprintf("%s_file_transfer_jobs", operation), func(token string) error {
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
		return nil, fmt.Errorf("API call error: %s", err.Error())
	}

	// Handle API response errors
	if apiResp != nil {
		if apiResp.CreateTriggersResponseOneOf != nil && apiResp.CreateTriggersResponseOneOf.Error != "" {
			tflog.Error(ctx, "API error during trigger creation", map[string]interface{}{
				"error": apiResp.CreateTriggersResponseOneOf.Error,
			})
			return nil, fmt.Errorf("API error: %s", apiResp.CreateTriggersResponseOneOf.Error)
		}

		if apiResp.MapmapOfStringarrayOfString != nil {
			tflog.Debug(ctx, "Received success response from CreateTrigger API")
		}
	}

	tflog.Info(ctx, "File Transfer Job triggers created successfully", map[string]interface{}{
		"job_count": len(jobs),
	})

	return apiResp, nil
}

// DeleteFileTransferJobs handles the business logic for deleting File Transfer jobs
func (r *FileTransferJobResource) DeleteFileTransferJobs(ctx context.Context, jobs []FileTransferJobModel) error {
	if len(jobs) == 0 {
		return nil
	}

	tflog.Debug(ctx, "Starting File Transfer Job triggers deletion", map[string]interface{}{
		"job_count": len(jobs),
	})

	for i, job := range jobs {
		name := job.Name.ValueString()
		jobGroup := job.JobGroup.ValueString()

		tflog.Debug(ctx, "Deleting job trigger", map[string]interface{}{
			"job_index":    i + 1,
			"trigger_name": name,
		})

		deleteReq := openapi.DeleteTriggerRequest{
			Triggername: name,
			Jobname:     "FileTransferJob",
			Jobgroup:    jobGroup,
		}

		var apiResp *openapi.DeleteTriggerResponse
		var finalHttpResp *http.Response
		err := r.provider.AuthenticatedAPICallWithRetry(ctx, "delete_file_transfer_job", func(token string) error {
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
			return fmt.Errorf("API call error for trigger '%s': %s", name, err.Error())
		}

		// Handle HTTP errors
		var diags diag.Diagnostics
		if jobcontrolutil.JobControlHandleHTTPError(ctx, finalHttpResp, err, fmt.Sprintf("deleting File Transfer Job trigger '%s'", name), &diags) {
			tflog.Error(ctx, "Failed to delete File Transfer Job trigger", map[string]interface{}{
				"error":        fmt.Sprintf("%v", diags.Errors()),
				"trigger_name": name,
			})
			if diags.HasError() {
				return fmt.Errorf("HTTP error for trigger '%s': %s", name, diags.Errors()[0].Detail())
			}
		}

		// Handle API response errors
		if apiResp != nil {
			errorCodeStr := fmt.Sprintf("%d", apiResp.ErrorCode)
			if jobcontrolutil.JobControlHandleAPIError(ctx, &errorCodeStr, &apiResp.Msg, fmt.Sprintf("deleting File Transfer Job trigger '%s'", name), &diags) {
				tflog.Error(ctx, "API error during trigger deletion", map[string]interface{}{
					"error":        fmt.Sprintf("%v", diags.Errors()),
					"trigger_name": name,
				})
				if diags.HasError() {
					return fmt.Errorf("API error for trigger '%s': %s", name, diags.Errors()[0].Detail())
				}
			}
		}

		tflog.Info(ctx, "File Transfer Job trigger deleted successfully", map[string]interface{}{
			"trigger_name": name,
		})
	}

	tflog.Info(ctx, "All File Transfer Job triggers deleted successfully", map[string]interface{}{
		"job_count": len(jobs),
	})

	return nil
}

func (r *FileTransferJobResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan FileTransferJobResourceModel

	tflog.Debug(ctx, "Starting File Transfer Job resource creation")

	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		resp.Diagnostics.AddError("Plan Extraction Failed", "Unable to extract Terraform plan from request")
		return
	}

	var jobs []FileTransferJobModel
	resp.Diagnostics.Append(plan.Jobs.ElementsAs(ctx, &jobs, false)...)
	if resp.Diagnostics.HasError() {
		resp.Diagnostics.AddError("Jobs Extraction Failed", "Unable to extract jobs from Terraform plan")
		return
	}

	_, err := r.CreateOrUpdateFileTransferJobs(ctx, jobs, "create")
	if err != nil {
		tflog.Error(ctx, "File Transfer Job creation failed", map[string]interface{}{
			"error": err.Error(),
		})
		resp.Diagnostics.AddError("File Transfer Job Creation Failed", err.Error())
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		resp.Diagnostics.AddError("State Update Failed", "Unable to save File Transfer Job state")
		return
	}

	tflog.Info(ctx, "File Transfer Job resource created successfully", map[string]interface{}{
		"job_count": len(jobs),
	})
}

func (r *FileTransferJobResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state FileTransferJobResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading File Transfer Job triggers")
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *FileTransferJobResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan FileTransferJobResourceModel

	tflog.Debug(ctx, "Starting File Transfer Job resource update")

	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		resp.Diagnostics.AddError("Plan Extraction Failed", "Unable to extract Terraform plan from request")
		return
	}

	var jobs []FileTransferJobModel
	resp.Diagnostics.Append(plan.Jobs.ElementsAs(ctx, &jobs, false)...)
	if resp.Diagnostics.HasError() {
		resp.Diagnostics.AddError("Jobs Extraction Failed", "Unable to extract jobs from Terraform plan")
		return
	}

	_, err := r.CreateOrUpdateFileTransferJobs(ctx, jobs, "update")
	if err != nil {
		tflog.Error(ctx, "File Transfer Job update failed", map[string]interface{}{
			"error": err.Error(),
		})
		resp.Diagnostics.AddError("File Transfer Job Update Failed", err.Error())
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		resp.Diagnostics.AddError("State Update Failed", "Unable to save File Transfer Job state")
		return
	}

	tflog.Info(ctx, "File Transfer Job resource updated successfully", map[string]interface{}{
		"job_count": len(jobs),
	})
}

func (r *FileTransferJobResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state FileTransferJobResourceModel

	tflog.Debug(ctx, "Starting File Transfer Job resource deletion")

	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		resp.Diagnostics.AddError("State Extraction Failed", "Unable to extract Terraform state from request")
		return
	}

	var jobs []FileTransferJobModel
	resp.Diagnostics.Append(state.Jobs.ElementsAs(ctx, &jobs, false)...)
	if resp.Diagnostics.HasError() {
		resp.Diagnostics.AddError("Jobs Extraction Failed", "Unable to extract jobs from Terraform state")
		return
	}

	err := r.DeleteFileTransferJobs(ctx, jobs)
	if err != nil {
		tflog.Error(ctx, "File Transfer Job deletion failed", map[string]interface{}{
			"error": err.Error(),
		})
		resp.Diagnostics.AddError("File Transfer Job Deletion Failed", err.Error())
		return
	}

	tflog.Info(ctx, "File Transfer Job resource deleted successfully", map[string]interface{}{
		"job_count": len(jobs),
	})
}
