// Copyright (c) 2025 Saviynt Inc.
// SPDX-License-Identifier: MPL-2.0

// accounts_import_incremental_job_resource.go manages Accounts Import Incremental Job triggers in Saviynt.
// The resource implements the full Terraform lifecycle:
//   - Create: provisions a new accounts import incremental job trigger using the supplied configuration.
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
var _ resource.Resource = &AccountsImportIncrementalJobResource{}

func NewAccountsImportIncrementalJobResource() resource.Resource {
	return &AccountsImportIncrementalJobResource{}
}

// AccountsImportIncrementalJobResource defines the resource implementation.
type AccountsImportIncrementalJobResource struct {
	client            client.SaviyntClientInterface
	provider          client.SaviyntProviderInterface
	jobControlFactory client.JobControlFactoryInterface
	token             string
}

// AccountsImportIncrementalJobResourceModel describes the resource data model.
type AccountsImportIncrementalJobResourceModel struct {
	Jobs types.List `tfsdk:"jobs"`
}

// AccountsImportIncrementalJobModel describes individual job data model.
type AccountsImportIncrementalJobModel struct {
	BaseJobTriggerResourceModel
	ConnectionName types.String `tfsdk:"connection_name"`
}

func AccountsImportIncrementalJobResourceSchema() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"jobs": schema.ListNestedAttribute{
			Required:    true,
			Description: "List of Accounts Import Incremental Job configurations",
			NestedObject: schema.NestedAttributeObject{
				Attributes: jobcontrolutil.MergeJobResourceAttributes(
					BaseJobTriggerResourceSchema(),
					map[string]schema.Attribute{
						"connection_name": schema.StringAttribute{
							Optional:    true,
							Description: "Connection name for the accounts import incremental job",
						},
					},
				),
			},
		},
	}
}

func (r *AccountsImportIncrementalJobResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: util.AccountsImportIncrementalJobDescription,
		Attributes:          AccountsImportIncrementalJobResourceSchema(),
	}
}

func (r *AccountsImportIncrementalJobResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_accounts_import_incremental_job_resource"
}

func (r *AccountsImportIncrementalJobResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	tflog.Debug(ctx, "Starting AccountsImportIncrementalJobResource configuration")

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

	tflog.Info(ctx, "AccountsImportIncrementalJobResource configuration completed successfully")
}

func (r *AccountsImportIncrementalJobResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan AccountsImportIncrementalJobResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var jobs []AccountsImportIncrementalJobModel
	resp.Diagnostics.Append(plan.Jobs.ElementsAs(ctx, &jobs, false)...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Starting Accounts Import Incremental Job triggers creation", map[string]interface{}{
		"job_count": len(jobs),
	})

	var jobTriggerItems []openapi.JobTriggerItem

	for i, job := range jobs {
		// Validate job name
		if job.JobName.ValueString() != "AccountsImportIncrementalJob" {
			resp.Diagnostics.AddError(
				"Invalid Job Name",
				fmt.Sprintf("Job %d: job_name must be 'AccountsImportIncrementalJob', got '%s'", i, job.JobName.ValueString()),
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

		// Create the value map
		valueMap := openapi.NewAccountsImportIncrementalJobAllOfValueMap()
		if !job.ConnectionName.IsNull() && job.ConnectionName.ValueString() != "" {
			valueMap.SetCONNECTION(job.ConnectionName.ValueString())
		}

		// Create the job trigger
		jobTrigger := openapi.NewAccountsImportIncrementalJob(
			job.Name.ValueString(),
			job.JobName.ValueString(),
			job.JobGroup.ValueString(),
			job.Group.ValueString(),
			job.CronExp.ValueString(),
		)
		jobTrigger.SetValueMap(*valueMap)

		// Create job trigger item
		jobTriggerItem := openapi.AccountsImportIncrementalJobAsJobTriggerItem(jobTrigger)
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
	err := r.provider.AuthenticatedAPICallWithRetry(ctx, "create_accounts_import_incremental_jobs", func(token string) error {
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
		resp.Diagnostics.AddError(
			"API Call Error",
			fmt.Sprintf("Error during API call to create Accounts Import Incremental Job triggers: %s", err.Error()),
		)
		return
	}

	// Handle API response errors - CreateTriggersResponse has different structure
	if apiResp != nil {
		// Check for error response
		if apiResp.CreateTriggersResponseOneOf != nil && apiResp.CreateTriggersResponseOneOf.Error != "" {
			tflog.Error(ctx, "API error during trigger creation", map[string]interface{}{
				"error": apiResp.CreateTriggersResponseOneOf.Error,
			})
			resp.Diagnostics.AddError("API Error", fmt.Sprintf("Error creating Accounts Import Incremental Job triggers: %s", apiResp.CreateTriggersResponseOneOf.Error))
			return
		}

		// Check for success response
		if apiResp.MapmapOfStringarrayOfString != nil {
			tflog.Debug(ctx, "Received success response from CreateTrigger API")
		}
	}

	tflog.Info(ctx, "Accounts Import Incremental Job triggers created successfully", map[string]interface{}{
		"job_count": len(jobs),
	})

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *AccountsImportIncrementalJobResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state AccountsImportIncrementalJobResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading Accounts Import Incremental Job triggers")

	// For now, we'll keep the state as-is since the API doesn't provide a direct read method
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *AccountsImportIncrementalJobResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan AccountsImportIncrementalJobResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var jobs []AccountsImportIncrementalJobModel
	resp.Diagnostics.Append(plan.Jobs.ElementsAs(ctx, &jobs, false)...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Starting Accounts Import Incremental Job triggers update", map[string]interface{}{
		"job_count": len(jobs),
	})

	var jobTriggerItems []openapi.JobTriggerItem

	for i, job := range jobs {
		// Validate job name
		if job.JobName.ValueString() != "AccountsImportIncrementalJob" {
			resp.Diagnostics.AddError(
				"Invalid Job Name",
				fmt.Sprintf("Job %d: job_name must be 'AccountsImportIncrementalJob', got '%s'", i, job.JobName.ValueString()),
			)
			return
		}

		// Create the value map
		valueMap := openapi.NewAccountsImportIncrementalJobAllOfValueMap()
		if !job.ConnectionName.IsNull() && job.ConnectionName.ValueString() != "" {
			valueMap.SetCONNECTION(job.ConnectionName.ValueString())
		}

		// Create the job trigger
		jobTrigger := openapi.NewAccountsImportIncrementalJob(
			job.Name.ValueString(),
			job.JobName.ValueString(),
			job.JobGroup.ValueString(),
			job.Group.ValueString(),
			job.CronExp.ValueString(),
		)
		jobTrigger.SetValueMap(*valueMap)

		// Create job trigger item
		jobTriggerItem := openapi.AccountsImportIncrementalJobAsJobTriggerItem(jobTrigger)
		jobTriggerItems = append(jobTriggerItems, jobTriggerItem)
	}

	// Create the request
	updateReq := []openapi.JobTriggerRequest{
		{
			Triggers: jobTriggerItems,
		},
	}

	// Make the API call
	var apiResp *openapi.CreateTriggersResponse
	var finalHttpResp *http.Response
	err := r.provider.AuthenticatedAPICallWithRetry(ctx, "update_accounts_import_incremental_jobs", func(token string) error {
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
		resp.Diagnostics.AddError(
			"API Call Error",
			fmt.Sprintf("Error during API call to update Accounts Import Incremental Job triggers: %s", err.Error()),
		)
		return
	}

	// Handle API response errors - CreateTriggersResponse has different structure
	if apiResp != nil {
		// Check for error response
		if apiResp.CreateTriggersResponseOneOf != nil && apiResp.CreateTriggersResponseOneOf.Error != "" {
			tflog.Error(ctx, "API error during trigger update", map[string]interface{}{
				"error": apiResp.CreateTriggersResponseOneOf.Error,
			})
			resp.Diagnostics.AddError("API Error", fmt.Sprintf("Error updating Accounts Import Incremental Job triggers: %s", apiResp.CreateTriggersResponseOneOf.Error))
			return
		}

		// Check for success response
		if apiResp.MapmapOfStringarrayOfString != nil {
			tflog.Debug(ctx, "Received success response from CreateTrigger API")
		}
	}

	tflog.Info(ctx, "Accounts Import Incremental Job triggers updated successfully", map[string]interface{}{
		"job_count": len(jobs),
	})

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *AccountsImportIncrementalJobResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state AccountsImportIncrementalJobResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var jobs []AccountsImportIncrementalJobModel
	resp.Diagnostics.Append(state.Jobs.ElementsAs(ctx, &jobs, false)...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Starting Accounts Import Incremental Job triggers deletion", map[string]interface{}{
		"job_count": len(jobs),
	})

	// Delete each job individually since DeleteTrigger API doesn't support bulk operations
	for _, job := range jobs {
		name := job.Name.ValueString()
		jobName := job.JobName.ValueString()
		jobGroup := job.JobGroup.ValueString()

		// Create delete request
		deleteReq := openapi.DeleteTriggerRequest{
			Triggername: name,
			Jobname:     jobName,
			Jobgroup:    jobGroup,
		}

		// Make the API call
		var apiResp *openapi.DeleteTriggerResponse
		var finalHttpResp *http.Response
		err := r.provider.AuthenticatedAPICallWithRetry(ctx, "delete_accounts_import_incremental_job", func(token string) error {
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
				fmt.Sprintf("Error during API call to delete Accounts Import Incremental Job trigger '%s': %s", name, err.Error()),
			)
			return
		}

		// Handle HTTP errors using job control error handler
		var diags diag.Diagnostics
		if jobcontrolutil.JobControlHandleHTTPError(ctx, finalHttpResp, err, fmt.Sprintf("deleting Accounts Import Incremental Job trigger '%s'", name), &diags) {
			tflog.Error(ctx, "Failed to delete Accounts Import Incremental Job trigger", map[string]interface{}{
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
			if jobcontrolutil.JobControlHandleAPIError(ctx, &errorCodeStr, &apiResp.Msg, fmt.Sprintf("deleting Accounts Import Incremental Job trigger '%s'", name), &diags) {
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

		tflog.Info(ctx, "Accounts Import Incremental Job trigger deleted successfully", map[string]interface{}{
			"trigger_name": name,
		})
	}

	tflog.Info(ctx, "All Accounts Import Incremental Job triggers deleted successfully", map[string]interface{}{
		"job_count": len(jobs),
	})
}
