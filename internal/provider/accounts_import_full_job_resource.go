// Copyright (c) 2025 Saviynt Inc.
// SPDX-License-Identifier: MPL-2.0

// accounts_import_full_job_resource.go manages Accounts Import Full Job triggers in Saviynt.
// The resource implements the full Terraform lifecycle:
//   - Create: provisions a new accounts import full job trigger using the supplied configuration.
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
var _ resource.Resource = &AccountsImportFullJobResource{}

func NewAccountsImportFullJobResource() resource.Resource {
	return &AccountsImportFullJobResource{}
}

// AccountsImportFullJobResource defines the resource implementation.
type AccountsImportFullJobResource struct {
	client            client.SaviyntClientInterface
	provider          client.SaviyntProviderInterface
	jobControlFactory client.JobControlFactoryInterface
	token             string
}

// AccountsImportFullJobResourceModel describes the resource data model.
type AccountsImportFullJobResourceModel struct {
	Jobs types.List `tfsdk:"jobs"`
}

// AccountsImportFullJobModel describes individual job data model.
type AccountsImportFullJobModel struct {
	BaseJobControlResourceModel
	ConnectionName types.String `tfsdk:"connection_name"`
}

func AccountsImportFullJobResourceSchema() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"jobs": schema.ListNestedAttribute{
			Required:    true,
			Description: "List of Accounts Import Full Job configurations",
			NestedObject: schema.NestedAttributeObject{
				Attributes: jobcontrolutil.MergeJobResourceAttributes(
					BaseJobControlResourceSchema(),
					map[string]schema.Attribute{
						"connection_name": schema.StringAttribute{
							Required:    true,
							Description: "Name of the connection for the accounts import",
						},
					},
				),
			},
		},
	}
}

func (r *AccountsImportFullJobResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: util.AccountsImportFullJobDescription,
		Attributes:          AccountsImportFullJobResourceSchema(),
	}
}

func (r *AccountsImportFullJobResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_accounts_import_full_job_resource"
}

func (r *AccountsImportFullJobResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	tflog.Debug(ctx, "Starting AccountsImportFullJobResource configuration")

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

	tflog.Info(ctx, "AccountsImportFullJobResource configuration completed successfully")
}

// SetClient sets the client for testing purposes
func (r *AccountsImportFullJobResource) SetClient(client client.SaviyntClientInterface) {
	r.client = client
}

// SetToken sets the token for testing purposes
func (r *AccountsImportFullJobResource) SetToken(token string) {
	r.token = token
}

// SetProvider sets the provider for testing purposes
func (r *AccountsImportFullJobResource) SetProvider(provider client.SaviyntProviderInterface) {
	r.provider = provider
}

// SetJobControlFactory sets the job control factory for testing purposes
func (r *AccountsImportFullJobResource) SetJobControlFactory(factory client.JobControlFactoryInterface) {
	r.jobControlFactory = factory
}

// CreateOrUpdateAccountsImportFullJobs handles the business logic for creating or updating accounts import full jobs
func (r *AccountsImportFullJobResource) CreateOrUpdateAccountsImportFullJobs(ctx context.Context, jobs []AccountsImportFullJobModel, operation string) (*openapi.CreateOrUpdateTriggersResponse, error) {
	if len(jobs) == 0 {
		return nil, fmt.Errorf("at least one job must be specified")
	}

	tflog.Debug(ctx, "Starting Accounts Import Full Job triggers creation", map[string]interface{}{
		"job_count": len(jobs),
	})

	var triggers []openapi.TriggerItem

	for i, job := range jobs {
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
		if job.ConnectionName.IsNull() || job.ConnectionName.ValueString() == "" {
			return nil, fmt.Errorf("job %d: connection_name is required", i+1)
		}

		// Create the value map
		valueMap := openapi.NewAccountsImportFullJobAllOfValueMap(job.ConnectionName.ValueString())

		// Create the job trigger
		jobTrigger := openapi.NewAccountsImportFullJob(
			job.TriggerName.ValueString(),
			"AccountsImportFullJob",
			job.JobGroup.ValueString(),
			job.CronExpression.ValueString(),
		)
		jobTrigger.SetValueMap(*valueMap)

		// Set optional trigger group if provided
		if !job.TriggerGroup.IsNull() && job.TriggerGroup.ValueString() != "" {
			jobTrigger.SetTriggergroup(job.TriggerGroup.ValueString())
		}

		// Create trigger item
		triggerItem := openapi.AccountsImportFullJobAsTriggerItem(jobTrigger)
		triggers = append(triggers, triggerItem)
	}

	// Create the request
	createReq := openapi.CreateOrUpdateTriggersRequest{
		Triggers: triggers,
	}

	// Make the API call
	var apiResp *openapi.CreateOrUpdateTriggersResponse
	var finalHttpResp *http.Response
	err := r.provider.AuthenticatedAPICallWithRetry(ctx, fmt.Sprintf("%s_accounts_import_full_jobs", operation), func(token string) error {
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

	// Handle HTTP errors using job control error handler
	var diags diag.Diagnostics
	if jobcontrolutil.JobControlHandleHTTPError(ctx, finalHttpResp, err, "creating Accounts Import Full Job triggers", &diags) {
		tflog.Error(ctx, "Failed to create Accounts Import Full Job triggers", map[string]interface{}{
			"error": fmt.Sprintf("%v", diags.Errors()),
		})
		if diags.HasError() {
			return nil, fmt.Errorf("HTTP error: %s", diags.Errors()[0].Detail())
		}
	}

	// Handle API response errors
	if apiResp != nil && jobcontrolutil.JobControlHandleAPIError(ctx, &apiResp.ErrorCode, &apiResp.Msg, "creating Accounts Import Full Job triggers", &diags) {
		tflog.Error(ctx, "API error during trigger creation", map[string]interface{}{
			"error": fmt.Sprintf("%v", diags.Errors()),
		})
		if diags.HasError() {
			return nil, fmt.Errorf("API error: %s", diags.Errors()[0].Detail())
		}
	}

	tflog.Info(ctx, "Accounts Import Full Job triggers created successfully", map[string]interface{}{
		"job_count": len(jobs),
	})

	return apiResp, nil
}

func (r *AccountsImportFullJobResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan AccountsImportFullJobResourceModel

	tflog.Debug(ctx, "Starting Accounts Import Full Job resource creation")

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
	var jobs []AccountsImportFullJobModel
	resp.Diagnostics.Append(plan.Jobs.ElementsAs(ctx, &jobs, false)...)
	if resp.Diagnostics.HasError() {
		resp.Diagnostics.AddError(
			"Jobs Extraction Failed",
			"Unable to extract jobs from Terraform plan",
		)
		return
	}

	// Call the business logic method
	_, err := r.CreateOrUpdateAccountsImportFullJobs(ctx, jobs, "create")
	if err != nil {
		tflog.Error(ctx, "Accounts Import Full Job creation failed", map[string]interface{}{
			"error": err.Error(),
		})
		resp.Diagnostics.AddError(
			"Accounts Import Full Job Creation Failed",
			err.Error(),
		)
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		resp.Diagnostics.AddError(
			"State Update Failed",
			"Unable to save Accounts Import Full Job state",
		)
		return
	}

	tflog.Info(ctx, "Accounts Import Full Job resource created successfully", map[string]interface{}{
		"job_count": len(jobs),
	})
}

func (r *AccountsImportFullJobResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state AccountsImportFullJobResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading Accounts Import Full Job triggers")

	// For now, we'll keep the state as-is since the API doesn't provide a direct read method
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *AccountsImportFullJobResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan AccountsImportFullJobResourceModel

	tflog.Debug(ctx, "Starting Accounts Import Full Job resource update")

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
	var jobs []AccountsImportFullJobModel
	resp.Diagnostics.Append(plan.Jobs.ElementsAs(ctx, &jobs, false)...)
	if resp.Diagnostics.HasError() {
		resp.Diagnostics.AddError(
			"Jobs Extraction Failed",
			"Unable to extract jobs from Terraform plan",
		)
		return
	}

	// Call the business logic method
	_, err := r.CreateOrUpdateAccountsImportFullJobs(ctx, jobs, "update")
	if err != nil {
		tflog.Error(ctx, "Accounts Import Full Job update failed", map[string]interface{}{
			"error": err.Error(),
		})
		resp.Diagnostics.AddError(
			"Accounts Import Full Job Update Failed",
			err.Error(),
		)
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		resp.Diagnostics.AddError(
			"State Update Failed",
			"Unable to save Accounts Import Full Job state",
		)
		return
	}

	tflog.Info(ctx, "Accounts Import Full Job resource updated successfully", map[string]interface{}{
		"job_count": len(jobs),
	})
}

// DeleteAccountsImportFullJobs handles the business logic for deleting accounts import full jobs
func (r *AccountsImportFullJobResource) DeleteAccountsImportFullJobs(ctx context.Context, jobs []AccountsImportFullJobModel) error {
	tflog.Debug(ctx, "Starting Accounts Import Full Job triggers deletion", map[string]interface{}{
		"job_count": len(jobs),
	})

	// Delete each job individually since DeleteTrigger API doesn't support bulk operations
	for i, job := range jobs {
		triggerName := job.TriggerName.ValueString()
		jobGroup := job.JobGroup.ValueString()

		tflog.Debug(ctx, "Deleting job trigger", map[string]interface{}{
			"job_index":    i + 1,
			"trigger_name": triggerName,
		})

		// Create delete request
		deleteReq := openapi.DeleteTriggerRequest{
			Triggername: triggerName,
			Jobname:     "AccountsImportFullJob",
			Jobgroup:    jobGroup,
		}

		// Make the API call
		var apiResp *openapi.DeleteTriggerResponse
		var finalHttpResp *http.Response
		err := r.provider.AuthenticatedAPICallWithRetry(ctx, "delete_accounts_import_full_job", func(token string) error {
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

		// Handle HTTP errors using job control error handler
		var diags diag.Diagnostics
		if jobcontrolutil.JobControlHandleHTTPError(ctx, finalHttpResp, err, fmt.Sprintf("deleting Accounts Import Full Job trigger '%s'", triggerName), &diags) {
			tflog.Error(ctx, "Failed to delete Accounts Import Full Job trigger", map[string]interface{}{
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
			if jobcontrolutil.JobControlHandleAPIError(ctx, &errorCodeStr, &apiResp.Msg, fmt.Sprintf("deleting Accounts Import Full Job trigger '%s'", triggerName), &diags) {
				tflog.Error(ctx, "API error during trigger deletion", map[string]interface{}{
					"error":        fmt.Sprintf("%v", diags.Errors()),
					"trigger_name": triggerName,
				})
				if diags.HasError() {
					return fmt.Errorf("API error for trigger '%s': %s", triggerName, diags.Errors()[0].Detail())
				}
			}
		}

		tflog.Info(ctx, "Accounts Import Full Job trigger deleted successfully", map[string]interface{}{
			"trigger_name": triggerName,
		})
	}

	tflog.Info(ctx, "All Accounts Import Full Job triggers deleted successfully", map[string]interface{}{
		"job_count": len(jobs),
	})

	return nil
}

func (r *AccountsImportFullJobResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state AccountsImportFullJobResourceModel

	tflog.Debug(ctx, "Starting Accounts Import Full Job resource deletion")

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
	var jobs []AccountsImportFullJobModel
	resp.Diagnostics.Append(state.Jobs.ElementsAs(ctx, &jobs, false)...)
	if resp.Diagnostics.HasError() {
		resp.Diagnostics.AddError(
			"Jobs Extraction Failed",
			"Unable to extract jobs from Terraform state",
		)
		return
	}

	// Call the business logic method
	err := r.DeleteAccountsImportFullJobs(ctx, jobs)
	if err != nil {
		tflog.Error(ctx, "Accounts Import Full Job deletion failed", map[string]interface{}{
			"error": err.Error(),
		})
		resp.Diagnostics.AddError(
			"Accounts Import Full Job Deletion Failed",
			err.Error(),
		)
		return
	}

	tflog.Info(ctx, "Accounts Import Full Job resource deleted successfully", map[string]interface{}{
		"job_count": len(jobs),
	})
}
