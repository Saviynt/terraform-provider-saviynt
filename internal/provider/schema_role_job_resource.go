// Copyright (c) 2025 Saviynt Inc.
// SPDX-License-Identifier: MPL-2.0

// schema_role_job_resource.go manages Schema Role Job triggers in Saviynt.
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

var _ resource.Resource = &SchemaRoleJobResource{}

func NewSchemaRoleJobResource() resource.Resource {
	return &SchemaRoleJobResource{}
}

type SchemaRoleJobResource struct {
	client            client.SaviyntClientInterface
	provider          client.SaviyntProviderInterface
	jobControlFactory client.JobControlFactoryInterface
	token             string
}

type SchemaRoleJobResourceModel struct {
	Jobs types.List `tfsdk:"jobs"`
}

// SchemaRoleJobModel describes individual job data model.
type SchemaRoleJobModel struct {
	BaseJobTriggerResourceModel
	SchemaFileNames types.String `tfsdk:"schema_file_names"`
}

func SchemaRoleJobResourceSchema() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"jobs": schema.ListNestedAttribute{
			Required:    true,
			Description: "List of Schema Role Job configurations",
			NestedObject: schema.NestedAttributeObject{
				Attributes: jobcontrolutil.MergeJobResourceAttributes(
					BaseJobTriggerResourceSchema(),
					map[string]schema.Attribute{
						"schema_file_names": schema.StringAttribute{
							Optional:    true,
							Description: "Schema file names for the role job",
						},
					},
				),
			},
		},
	}
}

func (r *SchemaRoleJobResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: util.SchemaRoleJobDescription,
		Attributes:          SchemaRoleJobResourceSchema(),
	}
}

func (r *SchemaRoleJobResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_schema_role_job_resource"
}

func (r *SchemaRoleJobResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	tflog.Debug(ctx, "Starting SchemaRoleJobResource configuration")
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

	tflog.Info(ctx, "SchemaRoleJobResource configuration completed successfully")
}

// SetClient sets the client for testing purposes
func (r *SchemaRoleJobResource) SetClient(client client.SaviyntClientInterface) {
	r.client = client
}

// SetToken sets the token for testing purposes
func (r *SchemaRoleJobResource) SetToken(token string) {
	r.token = token
}

// SetProvider sets the provider for testing purposes
func (r *SchemaRoleJobResource) SetProvider(provider client.SaviyntProviderInterface) {
	r.provider = provider
}

// SetJobControlFactory sets the job control factory for testing purposes
func (r *SchemaRoleJobResource) SetJobControlFactory(factory client.JobControlFactoryInterface) {
	r.jobControlFactory = factory
}

// CreateOrUpdateSchemaRoleJobs handles the business logic for creating or updating Schema Role jobs
func (r *SchemaRoleJobResource) CreateOrUpdateSchemaRoleJobs(ctx context.Context, jobs []SchemaRoleJobModel, operation string) (*openapi.CreateTriggersResponse, error) {
	if len(jobs) == 0 {
		return nil, fmt.Errorf("at least one job must be specified")
	}

	tflog.Debug(ctx, "Starting Schema Role Job triggers creation", map[string]interface{}{
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

		// Create the value map
		valueMap := openapi.NewSchemaRoleJobAllOfValueMap()
		if !job.SchemaFileNames.IsNull() && job.SchemaFileNames.ValueString() != "" {
			valueMap.SetSchemaFileNames(job.SchemaFileNames.ValueString())
		}

		// Create the job trigger
		jobTrigger := openapi.NewSchemaRoleJob(
			job.Name.ValueString(),
			"SchemaRoleJob",
			job.JobGroup.ValueString(),
			job.Group.ValueString(),
			job.CronExp.ValueString(),
		)
		jobTrigger.SetValueMap(*valueMap)

		// Create job trigger item
		jobTriggerItem := openapi.SchemaRoleJobAsJobTriggerItem(jobTrigger)
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
	err := r.provider.AuthenticatedAPICallWithRetry(ctx, fmt.Sprintf("%s_schema_role_jobs", operation), func(token string) error {
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

	// Handle API response errors - CreateTriggersResponse has different structure
	if apiResp != nil {
		// Check for error response
		if apiResp.CreateTriggersResponseOneOf != nil && apiResp.CreateTriggersResponseOneOf.Error != "" {
			tflog.Error(ctx, "API error during trigger creation", map[string]interface{}{
				"error": apiResp.CreateTriggersResponseOneOf.Error,
			})
			return nil, fmt.Errorf("API error: %s", apiResp.CreateTriggersResponseOneOf.Error)
		}

		// Check for success response
		if apiResp.MapmapOfStringarrayOfString != nil {
			tflog.Debug(ctx, "Received success response from CreateTrigger API")
		}
	}

	tflog.Info(ctx, "Schema Role Job triggers created successfully", map[string]interface{}{
		"job_count": len(jobs),
	})

	return apiResp, nil
}

// DeleteSchemaRoleJobs handles the business logic for deleting Schema Role jobs
func (r *SchemaRoleJobResource) DeleteSchemaRoleJobs(ctx context.Context, jobs []SchemaRoleJobModel) error {
	if len(jobs) == 0 {
		return nil
	}

	tflog.Debug(ctx, "Starting Schema Role Job triggers deletion", map[string]interface{}{
		"job_count": len(jobs),
	})

	// Delete each job individually
	for i, job := range jobs {
		name := job.Name.ValueString()
		jobGroup := job.JobGroup.ValueString()

		tflog.Debug(ctx, "Deleting job trigger", map[string]interface{}{
			"job_index":    i + 1,
			"trigger_name": name,
		})

		// Create delete request
		deleteReq := openapi.DeleteTriggerRequest{
			Triggername: name,
			Jobname:     "SchemaRoleJob",
			Jobgroup:    jobGroup,
		}

		// Make the API call
		var apiResp *openapi.DeleteTriggerResponse
		var finalHttpResp *http.Response
		err := r.provider.AuthenticatedAPICallWithRetry(ctx, "delete_schema_role_job", func(token string) error {
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
		if jobcontrolutil.JobControlHandleHTTPError(ctx, finalHttpResp, err, fmt.Sprintf("deleting Schema Role Job trigger '%s'", name), &diags) {
			tflog.Error(ctx, "Failed to delete Schema Role Job trigger", map[string]interface{}{
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
			if jobcontrolutil.JobControlHandleAPIError(ctx, &errorCodeStr, &apiResp.Msg, fmt.Sprintf("deleting Schema Role Job trigger '%s'", name), &diags) {
				tflog.Error(ctx, "API error during trigger deletion", map[string]interface{}{
					"error":        fmt.Sprintf("%v", diags.Errors()),
					"trigger_name": name,
				})
				if diags.HasError() {
					return fmt.Errorf("API error for trigger '%s': %s", name, diags.Errors()[0].Detail())
				}
			}
		}

		tflog.Info(ctx, "Schema Role Job trigger deleted successfully", map[string]interface{}{
			"trigger_name": name,
		})
	}

	tflog.Info(ctx, "All Schema Role Job triggers deleted successfully", map[string]interface{}{
		"job_count": len(jobs),
	})

	return nil
}

func (r *SchemaRoleJobResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan SchemaRoleJobResourceModel

	tflog.Debug(ctx, "Starting Schema Role Job resource creation")

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
	var jobs []SchemaRoleJobModel
	resp.Diagnostics.Append(plan.Jobs.ElementsAs(ctx, &jobs, false)...)
	if resp.Diagnostics.HasError() {
		resp.Diagnostics.AddError(
			"Jobs Extraction Failed",
			"Unable to extract jobs from Terraform plan",
		)
		return
	}

	// Call the business logic method
	_, err := r.CreateOrUpdateSchemaRoleJobs(ctx, jobs, "create")
	if err != nil {
		tflog.Error(ctx, "Schema Role Job creation failed", map[string]interface{}{
			"error": err.Error(),
		})
		resp.Diagnostics.AddError(
			"Schema Role Job Creation Failed",
			err.Error(),
		)
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		resp.Diagnostics.AddError(
			"State Update Failed",
			"Unable to save Schema Role Job state",
		)
		return
	}

	tflog.Info(ctx, "Schema Role Job resource created successfully", map[string]interface{}{
		"job_count": len(jobs),
	})
}

func (r *SchemaRoleJobResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state SchemaRoleJobResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading Schema Role Job triggers")

	// For now, we'll keep the state as-is since the API doesn't provide a direct read method
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *SchemaRoleJobResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan SchemaRoleJobResourceModel

	tflog.Debug(ctx, "Starting Schema Role Job resource update")

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
	var jobs []SchemaRoleJobModel
	resp.Diagnostics.Append(plan.Jobs.ElementsAs(ctx, &jobs, false)...)
	if resp.Diagnostics.HasError() {
		resp.Diagnostics.AddError(
			"Jobs Extraction Failed",
			"Unable to extract jobs from Terraform plan",
		)
		return
	}

	// Call the business logic method
	_, err := r.CreateOrUpdateSchemaRoleJobs(ctx, jobs, "update")
	if err != nil {
		tflog.Error(ctx, "Schema Role Job update failed", map[string]interface{}{
			"error": err.Error(),
		})
		resp.Diagnostics.AddError(
			"Schema Role Job Update Failed",
			err.Error(),
		)
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		resp.Diagnostics.AddError(
			"State Update Failed",
			"Unable to save Schema Role Job state",
		)
		return
	}

	tflog.Info(ctx, "Schema Role Job resource updated successfully", map[string]interface{}{
		"job_count": len(jobs),
	})
}

func (r *SchemaRoleJobResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state SchemaRoleJobResourceModel

	tflog.Debug(ctx, "Starting Schema Role Job resource deletion")

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
	var jobs []SchemaRoleJobModel
	resp.Diagnostics.Append(state.Jobs.ElementsAs(ctx, &jobs, false)...)
	if resp.Diagnostics.HasError() {
		resp.Diagnostics.AddError(
			"Jobs Extraction Failed",
			"Unable to extract jobs from Terraform state",
		)
		return
	}

	// Call the business logic method
	err := r.DeleteSchemaRoleJobs(ctx, jobs)
	if err != nil {
		tflog.Error(ctx, "Schema Role Job deletion failed", map[string]interface{}{
			"error": err.Error(),
		})
		resp.Diagnostics.AddError(
			"Schema Role Job Deletion Failed",
			err.Error(),
		)
		return
	}

	tflog.Info(ctx, "Schema Role Job resource deleted successfully", map[string]interface{}{
		"job_count": len(jobs),
	})
}
