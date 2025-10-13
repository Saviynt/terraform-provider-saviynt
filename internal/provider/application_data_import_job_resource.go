// Copyright (c) 2025 Saviynt Inc.
// SPDX-License-Identifier: MPL-2.0

// application_data_import_job_resource.go manages Application Data Import Job triggers in Saviynt.
// The resource implements the full Terraform lifecycle:
//   - Create: provisions a new application data import job trigger using the supplied configuration.
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
var _ resource.Resource = &ApplicationDataImportJobResource{}

func NewApplicationDataImportJobResource() resource.Resource {
	return &ApplicationDataImportJobResource{}
}

// ApplicationDataImportJobResource defines the resource implementation.
type ApplicationDataImportJobResource struct {
	client            client.SaviyntClientInterface
	provider          client.SaviyntProviderInterface
	jobControlFactory client.JobControlFactoryInterface
	token             string
}

// ApplicationDataImportJobResourceModel describes the resource data model.
type ApplicationDataImportJobResourceModel struct {
	Jobs types.List `tfsdk:"jobs"`
}

// ApplicationDataImportJobModel describes individual job data model.
type ApplicationDataImportJobModel struct {
	BaseJobControlResourceModel
	SecuritySystems   types.List   `tfsdk:"security_systems"`
	AccountsOrAccess  types.String `tfsdk:"accounts_or_access"`
	ExternalConn      types.String `tfsdk:"external_conn"`
	FullOrIncremental types.String `tfsdk:"full_or_incremental"`
}

func (r *ApplicationDataImportJobResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_application_data_import_job_resource"
}

func ApplicationDataImportJobResourceSchema() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"security_systems": schema.ListAttribute{
			Required:    true,
			ElementType: types.StringType,
			Description: "List of security systems for the application data import",
		},
		"accounts_or_access": schema.StringAttribute{
			Optional:    true,
			Description: "Accounts or access configuration",
		},
		"external_conn": schema.StringAttribute{
			Optional:    true,
			Description: "External connection configuration",
		},
		"full_or_incremental": schema.StringAttribute{
			Optional:    true,
			Description: "Full or incremental import type",
		},
	}
}

func (r *ApplicationDataImportJobResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: util.ApplicationDataImportJobDescription,
		Attributes: map[string]schema.Attribute{
			"jobs": schema.ListNestedAttribute{
				Required:    true,
				Description: "List of Application Data Import Jobs to create",
				NestedObject: schema.NestedAttributeObject{
					Attributes: jobcontrolutil.MergeJobResourceAttributes(
						BaseJobControlResourceSchema(),
						ApplicationDataImportJobResourceSchema(),
					),
				},
			},
		},
	}
}

func (r *ApplicationDataImportJobResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	tflog.Debug(ctx, "Starting ApplicationDataImportJobResource configuration")

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

	tflog.Info(ctx, "ApplicationDataImportJobResource configuration completed successfully")
}

// SetClient sets the client for testing purposes
func (r *ApplicationDataImportJobResource) SetClient(client client.SaviyntClientInterface) {
	r.client = client
}

// SetToken sets the token for testing purposes
func (r *ApplicationDataImportJobResource) SetToken(token string) {
	r.token = token
}

// SetProvider sets the provider for testing purposes
func (r *ApplicationDataImportJobResource) SetProvider(provider client.SaviyntProviderInterface) {
	r.provider = provider
}
func (r *ApplicationDataImportJobResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan ApplicationDataImportJobResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Extract jobs from the plan
	var jobs []ApplicationDataImportJobModel
	resp.Diagnostics.Append(plan.Jobs.ElementsAs(ctx, &jobs, false)...)
	if resp.Diagnostics.HasError() {
		return
	}

	if len(jobs) == 0 {
		resp.Diagnostics.AddError("Validation Error", "At least one job must be specified")
		return
	}

	tflog.Debug(ctx, "Starting Application Data Import Job triggers creation", map[string]interface{}{
		"job_count": len(jobs),
	})

	var triggers []openapi.TriggerItem

	// Process each job
	for i, job := range jobs {
		// Validate job name is ApplicationDataImportJob
		if job.JobName.IsNull() || job.JobName.ValueString() != "ApplicationDataImportJob" {
			resp.Diagnostics.AddError(
				"Validation Error",
				fmt.Sprintf("Job %d: job_name must be 'ApplicationDataImportJob', got '%s'", i+1, job.JobName.ValueString()),
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

		// Convert security systems list
		var securitySystems []string
		resp.Diagnostics.Append(job.SecuritySystems.ElementsAs(ctx, &securitySystems, false)...)
		if resp.Diagnostics.HasError() {
			return
		}

		if len(securitySystems) == 0 {
			resp.Diagnostics.AddError("Validation Error", fmt.Sprintf("Job %d: security_systems is required", i+1))
			return
		}

		// Build the value map
		valueMap := openapi.NewApplicationDataImportJobAllOfValueMap(securitySystems)

		if !job.AccountsOrAccess.IsNull() {
			accountsOrAccess := job.AccountsOrAccess.ValueString()
			valueMap.Accountsoraccess = &accountsOrAccess
		}

		if !job.ExternalConn.IsNull() {
			externalConn := job.ExternalConn.ValueString()
			valueMap.Externalconn = &externalConn
		}

		if !job.FullOrIncremental.IsNull() {
			fullOrIncremental := job.FullOrIncremental.ValueString()
			valueMap.Fullorincremental = &fullOrIncremental
		}

		// Create the job
		appDataImportJob := openapi.NewApplicationDataImportJob(
			job.TriggerName.ValueString(),
			job.JobName.ValueString(),
			job.JobGroup.ValueString(),
			job.CronExpression.ValueString(),
		)
		appDataImportJob.SetValueMap(*valueMap)

		if !job.TriggerGroup.IsNull() {
			triggerGroup := job.TriggerGroup.ValueString()
			appDataImportJob.SetTriggergroup(triggerGroup)
		}

		triggers = append(triggers, openapi.ApplicationDataImportJobAsTriggerItem(appDataImportJob))
	}

	// Prepare the API request with all triggers
	createReq := openapi.CreateOrUpdateTriggersRequest{
		Triggers: triggers,
	}

	// Make the API call
	var apiResp *openapi.CreateOrUpdateTriggersResponse
	var finalHttpResp *http.Response
	err := r.provider.AuthenticatedAPICallWithRetry(ctx, "create_application_data_import_jobs", func(token string) error {
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
			fmt.Sprintf("Error during API call to create Application Data Import Job triggers: %s", err.Error()),
		)
		return
	}

	// Handle HTTP errors
	var diags diag.Diagnostics
	if jobcontrolutil.JobControlHandleHTTPError(ctx, finalHttpResp, err, "creating Application Data Import Job triggers", &diags) {
		tflog.Error(ctx, "Failed to create Application Data Import Job triggers", map[string]interface{}{
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
	if apiResp != nil && jobcontrolutil.JobControlHandleAPIError(ctx, &apiResp.ErrorCode, &apiResp.Msg, "creating Application Data Import Job triggers", &diags) {
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

	tflog.Info(ctx, "Application Data Import Job triggers created successfully", map[string]interface{}{
		"job_count": len(jobs),
	})

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *ApplicationDataImportJobResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state ApplicationDataImportJobResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading Application Data Import Job triggers", map[string]interface{}{
		"job_count": len(state.Jobs.Elements()),
	})

	// For now, we'll keep the state as-is since the API doesn't provide a direct read method
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *ApplicationDataImportJobResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan ApplicationDataImportJobResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Extract jobs from the plan
	var jobs []ApplicationDataImportJobModel
	resp.Diagnostics.Append(plan.Jobs.ElementsAs(ctx, &jobs, false)...)
	if resp.Diagnostics.HasError() {
		return
	}

	if len(jobs) == 0 {
		resp.Diagnostics.AddError("Validation Error", "At least one job must be specified")
		return
	}

	tflog.Debug(ctx, "Starting Application Data Import Job triggers update", map[string]interface{}{
		"job_count": len(jobs),
	})

	var triggers []openapi.TriggerItem

	// Process each job
	for i, job := range jobs {
		// Validate job name is ApplicationDataImportJob
		if job.JobName.IsNull() || job.JobName.ValueString() != "ApplicationDataImportJob" {
			resp.Diagnostics.AddError(
				"Validation Error", 
				fmt.Sprintf("Job %d: job_name must be 'ApplicationDataImportJob', got '%s'", i+1, job.JobName.ValueString()),
			)
			return
		}

		// Convert security systems list
		var securitySystems []string
		resp.Diagnostics.Append(job.SecuritySystems.ElementsAs(ctx, &securitySystems, false)...)
		if resp.Diagnostics.HasError() {
			return
		}

		// Build the value map
		valueMap := openapi.NewApplicationDataImportJobAllOfValueMap(securitySystems)

		if !job.AccountsOrAccess.IsNull() && job.AccountsOrAccess.ValueString() != "" {
			accountsOrAccess := job.AccountsOrAccess.ValueString()
			valueMap.Accountsoraccess = &accountsOrAccess
		}

		if !job.ExternalConn.IsNull() && job.ExternalConn.ValueString() != "" {
			externalConn := job.ExternalConn.ValueString()
			valueMap.Externalconn = &externalConn
		}

		if !job.FullOrIncremental.IsNull() && job.FullOrIncremental.ValueString() != "" {
			fullOrIncremental := job.FullOrIncremental.ValueString()
			valueMap.Fullorincremental = &fullOrIncremental
		}

		appDataImportJob := openapi.NewApplicationDataImportJob(
			job.TriggerName.ValueString(),
			job.JobName.ValueString(),
			job.JobGroup.ValueString(),
			job.CronExpression.ValueString(),
		)
		appDataImportJob.SetValueMap(*valueMap)

		if !job.TriggerGroup.IsNull() {
			triggerGroup := job.TriggerGroup.ValueString()
			appDataImportJob.SetTriggergroup(triggerGroup)
		}

		triggers = append(triggers, openapi.ApplicationDataImportJobAsTriggerItem(appDataImportJob))
	}

	updateReq := openapi.CreateOrUpdateTriggersRequest{
		Triggers: triggers,
	}

	var apiResp *openapi.CreateOrUpdateTriggersResponse
	var finalHttpResp *http.Response
	err := r.provider.AuthenticatedAPICallWithRetry(ctx, "update_application_data_import_jobs", func(token string) error {
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
			fmt.Sprintf("Error during API call to update Application Data Import Job triggers: %s", err.Error()),
		)
		return
	}

	// Handle HTTP errors
	var diags diag.Diagnostics
	if jobcontrolutil.JobControlHandleHTTPError(ctx, finalHttpResp, err, "updating Application Data Import Job triggers", &diags) {
		tflog.Error(ctx, "Failed to update Application Data Import Job triggers", map[string]interface{}{
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
	if apiResp != nil && jobcontrolutil.JobControlHandleAPIError(ctx, &apiResp.ErrorCode, &apiResp.Msg, "updating Application Data Import Job triggers", &diags) {
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

	tflog.Info(ctx, "Application Data Import Job triggers updated successfully", map[string]interface{}{
		"job_count": len(jobs),
	})

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *ApplicationDataImportJobResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state ApplicationDataImportJobResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Extract jobs from the state
	var jobs []ApplicationDataImportJobModel
	resp.Diagnostics.Append(state.Jobs.ElementsAs(ctx, &jobs, false)...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Starting Application Data Import Job triggers deletion", map[string]interface{}{
		"job_count": len(jobs),
	})

	// Delete each job individually since DeleteTrigger API doesn't support bulk operations
	for i, job := range jobs {
		triggerName := job.TriggerName.ValueString()
		jobName := job.JobName.ValueString()
		jobGroup := job.JobGroup.ValueString()

		tflog.Debug(ctx, "Deleting job trigger", map[string]interface{}{
			"job_index":    i + 1,
			"trigger_name": triggerName,
		})

		deleteReq := openapi.DeleteTriggerRequest{
			Triggername: triggerName,
			Jobname:     jobName,
			Jobgroup:    jobGroup,
		}

		var apiResp *openapi.DeleteTriggerResponse
		var finalHttpResp *http.Response
		err := r.provider.AuthenticatedAPICallWithRetry(ctx, "delete_application_data_import_job", func(token string) error {
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
				fmt.Sprintf("Error during API call to delete Application Data Import Job trigger '%s': %s", triggerName, err.Error()),
			)
			return
		}

		// Handle HTTP errors
		var diags diag.Diagnostics
		if jobcontrolutil.JobControlHandleHTTPError(ctx, finalHttpResp, err, fmt.Sprintf("deleting Application Data Import Job trigger '%s'", triggerName), &diags) {
			tflog.Error(ctx, "Failed to delete Application Data Import Job trigger", map[string]interface{}{
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
			if jobcontrolutil.JobControlHandleAPIError(ctx, &errorCodeStr, &apiResp.Msg, fmt.Sprintf("deleting Application Data Import Job trigger '%s'", triggerName), &diags) {
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

		tflog.Info(ctx, "Application Data Import Job trigger deleted successfully", map[string]interface{}{
			"trigger_name": triggerName,
		})
	}

	tflog.Info(ctx, "All Application Data Import Job triggers deleted successfully", map[string]interface{}{
		"job_count": len(jobs),
	})

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}
