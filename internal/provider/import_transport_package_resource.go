// Copyright (c) 2025 Saviynt Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"terraform-provider-Saviynt/internal/client"
	"terraform-provider-Saviynt/util"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	openapi "github.com/saviynt/saviynt-api-go-client/transports"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &ImportTransportPackageResource{}

// ImportTransportPackageResourceModel defines the resource data model.
type ImportTransportPackageResourceModel struct {
	ID                    types.String `tfsdk:"id"`
	PackagePath           types.String `tfsdk:"package_path"`
	UpdateUser            types.String `tfsdk:"update_user"`
	BusinessJustification types.String `tfsdk:"business_justification"`
	ImportPackageVersion  types.String `tfsdk:"import_package_version"`
	RequestID             types.String `tfsdk:"request_id"`
	Message               types.String `tfsdk:"message"`
	ErrorCode             types.String `tfsdk:"error_code"`
}

type ImportTransportPackageResource struct {
	client           client.SaviyntClientInterface
	token            string
	provider         client.SaviyntProviderInterface
	transportFactory client.TransportFactoryInterface
}

func NewImportTransportPackageResource() resource.Resource {
	return &ImportTransportPackageResource{
		transportFactory: &client.DefaultTransportFactory{},
	}
}

func NewImportTransportPackageResourceWithFactory(factory client.TransportFactoryInterface) resource.Resource {
	return &ImportTransportPackageResource{
		transportFactory: factory,
	}
}

func (r *ImportTransportPackageResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_import_transport_package_resource"
}

func (r *ImportTransportPackageResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: util.ImportTransportPackageDescription,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The unique ID of the resource.",
			},
			"package_path": schema.StringAttribute{
				Required:    true,
				Description: "Complete path of the package that needs to be imported.",
			},
			"update_user": schema.StringAttribute{
				Optional:    true,
				Description: "Username of the user importing the package.",
			},
			"business_justification": schema.StringAttribute{
				Optional:    true,
				Description: "Business justification for the import.",
			},
			"import_package_version": schema.StringAttribute{
				Optional:    true,
				Description: "Version identifier for the import package. Change this value to trigger re-import of the same package.",
			},
			"request_id": schema.StringAttribute{
				Computed:    true,
				Description: "Request ID generated during import submission.",
			},
			"message": schema.StringAttribute{
				Computed:    true,
				Description: "Response message from the import operation.",
			},
			"error_code": schema.StringAttribute{
				Computed:    true,
				Description: "Error code from the import operation.",
			},
		},
	}
}

func (r *ImportTransportPackageResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	tflog.Debug(ctx, "Starting ImportTransportPackageResource configuration")

	if req.ProviderData == nil {
		tflog.Debug(ctx, "Provider data is nil, skipping configuration")
		return
	}

	prov, ok := req.ProviderData.(*SaviyntProvider)
	if !ok {
		tflog.Error(ctx, "Type assertion failed", map[string]interface{}{
			"error": fmt.Sprintf("Expected *SaviyntProvider, got: %T", req.ProviderData),
		})
		resp.Diagnostics.AddError(
			"Unexpected Resource Configure Type",
			fmt.Sprintf("Expected *SaviyntProvider, got: %T", req.ProviderData),
		)
		return
	}

	r.client = &client.SaviyntClientWrapper{Client: prov.client}
	r.token = prov.accessToken
	r.provider = &client.SaviyntProviderWrapper{Provider: prov}

	tflog.Debug(ctx, "ImportTransportPackageResource configured successfully")
}

// ImportTransportPackage handles the business logic for importing transport packages
func (r *ImportTransportPackageResource) ImportTransportPackage(ctx context.Context, plan *ImportTransportPackageResourceModel, operation string) (*openapi.ImportTransportPackageResponse, error) {
	// Build import request
	importReq := openapi.ImportTransportPackageRequest{
		Packagetoimport: plan.PackagePath.ValueString(),
	}

	// Add optional fields
	if !plan.UpdateUser.IsNull() && plan.UpdateUser.ValueString() != "" {
		importReq.Updateuser = plan.UpdateUser.ValueStringPointer()
	}
	if !plan.BusinessJustification.IsNull() && plan.BusinessJustification.ValueString() != "" {
		importReq.Businessjustification = plan.BusinessJustification.ValueStringPointer()
	}

	var apiResp *openapi.ImportTransportPackageResponse
	var finalHttpResp *http.Response

	// Execute import operation with retry logic
	err := r.provider.AuthenticatedAPICallWithRetry(ctx, fmt.Sprintf("%s_transport_package", operation), func(token string) error {
		transportOps := r.transportFactory.CreateTransportOperations(r.client.APIBaseURL(), token)
		resp, httpResp, err := transportOps.ImportTransportPackage(ctx, importReq)
		if httpResp != nil && httpResp.StatusCode == 401 {
			return fmt.Errorf("401 unauthorized")
		}
		apiResp = resp
		finalHttpResp = httpResp
		return err
	})

	// Handle non-412 errors first
	if err != nil && finalHttpResp != nil && finalHttpResp.StatusCode != 412 {
		return nil, fmt.Errorf("API call failed: %w", err)
	}

	// Handle 412 precondition failed with response body decoding
	if err != nil && finalHttpResp != nil && finalHttpResp.StatusCode == 412 {
		var errorResp struct {
			Msg            string `json:"msg"`
			MsgDescription string `json:"msgDescription"`
			ErrorCode      int    `json:"errorcode"`
		}

		if decodeErr := json.NewDecoder(finalHttpResp.Body).Decode(&errorResp); decodeErr == nil {
			return nil, fmt.Errorf("precondition failed - ErrorCode: %d, Msg: %s, Description: %s",
				errorResp.ErrorCode, errorResp.Msg, errorResp.MsgDescription)
		}
		return nil, fmt.Errorf("precondition failed: package not found or invalid package path")
	}

	if err != nil {
		return nil, fmt.Errorf("API call failed: %w", err)
	}

	return apiResp, nil
}

// SetClient sets the client for testing purposes
func (r *ImportTransportPackageResource) SetClient(client client.SaviyntClientInterface) {
	r.client = client
}

// SetToken sets the token for testing purposes
func (r *ImportTransportPackageResource) SetToken(token string) {
	r.token = token
}

// SetProvider sets the provider for testing purposes
func (r *ImportTransportPackageResource) SetProvider(provider client.SaviyntProviderInterface) {
	r.provider = provider
}

func (r *ImportTransportPackageResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan ImportTransportPackageResourceModel

	tflog.Debug(ctx, "Starting import transport package resource creation")

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
	apiResp, err := r.ImportTransportPackage(ctx, &plan, "create")
	if err != nil {
		tflog.Error(ctx, "Import transport package creation failed", map[string]interface{}{
			"error": err.Error(),
		})
		resp.Diagnostics.AddError(
			"Import Transport Package Creation Failed",
			err.Error(),
		)
		return
	}

	// Update model from response
	r.UpdateModelFromResponse(&plan, apiResp)

	// Add success warning if operation completed successfully
	r.AddSuccessWarning(resp, apiResp)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		resp.Diagnostics.AddError(
			"State Update Failed",
			"Unable to save import transport package state",
		)
		return
	}

	tflog.Info(ctx, "Import transport package resource created successfully")
}

// UpdateModelFromResponse updates the model with API response data
func (r *ImportTransportPackageResource) UpdateModelFromResponse(plan *ImportTransportPackageResourceModel, apiResp *openapi.ImportTransportPackageResponse) {
	plan.ID = types.StringValue("import-transport-" + plan.PackagePath.ValueString())

	// Handle nil pointer dereference safely
	if apiResp != nil {
		if apiResp.Msg != nil {
			plan.Message = types.StringValue(*apiResp.Msg)
		} else {
			plan.Message = types.StringValue("")
		}
		if apiResp.Errorcode != nil {
			plan.ErrorCode = types.StringValue(fmt.Sprintf("%d", *apiResp.Errorcode))
		} else {
			plan.ErrorCode = types.StringValue("0")
		}
		if apiResp.RequestId != nil {
			plan.RequestID = types.StringValue(*apiResp.RequestId)
		} else {
			plan.RequestID = types.StringValue("")
		}
	} else {
		// Set null values when apiResp is nil to prevent framework errors
		plan.Message = types.StringNull()
		plan.ErrorCode = types.StringNull()
		plan.RequestID = types.StringNull()
	}
}

// AddSuccessWarning adds a warning for successful import operations
func (r *ImportTransportPackageResource) AddSuccessWarning(resp interface{}, apiResp *openapi.ImportTransportPackageResponse) {
	// Check if this is a successful operation (error_code = 0 and message = "success")
	if apiResp != nil && apiResp.Errorcode != nil && *apiResp.Errorcode == 0 {
		var message, requestID, errorCode string

		if apiResp.Msg != nil {
			message = *apiResp.Msg
		}
		if apiResp.RequestId != nil {
			requestID = *apiResp.RequestId
		}
		if apiResp.Errorcode != nil {
			errorCode = fmt.Sprintf("%d", *apiResp.Errorcode)
		}

		// Add warning for successful import
		if createResp, ok := resp.(*resource.CreateResponse); ok {
			createResp.Diagnostics.AddWarning(
				"Transport Package Import Completed Successfully",
				fmt.Sprintf(
					"Transport package imported successfully.\nMessage: %s\nError Code: %s\nRequest ID: %s\n\n"+
						"⚠️  IMPORTANT: The transport package has been imported into Saviynt. "+
						"Please verify the imported components in the Saviynt UI to ensure they are configured correctly.",
					message, errorCode, requestID,
				),
			)
		} else if updateResp, ok := resp.(*resource.UpdateResponse); ok {
			updateResp.Diagnostics.AddWarning(
				"Transport Package Import Completed Successfully",
				fmt.Sprintf(
					"Transport package imported successfully.\nMessage: %s\nError Code: %s\nRequest ID: %s\n\n"+
						"⚠️  IMPORTANT: The transport package has been imported into Saviynt. "+
						"Please verify the imported components in the Saviynt UI to ensure they are configured correctly.",
					message, errorCode, requestID,
				),
			)
		}
	}
}

func (r *ImportTransportPackageResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state ImportTransportPackageResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Import operations are one-time, so we just keep the existing state
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *ImportTransportPackageResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan ImportTransportPackageResourceModel

	tflog.Debug(ctx, "Starting import transport package resource update")

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
	apiResp, err := r.ImportTransportPackage(ctx, &plan, "update")
	if err != nil {
		tflog.Error(ctx, "Import transport package update failed", map[string]interface{}{
			"error": err.Error(),
		})
		resp.Diagnostics.AddError(
			"Import Transport Package Update Failed",
			err.Error(),
		)
		return
	}

	// Update model from response
	r.UpdateModelFromResponse(&plan, apiResp)

	// Add success warning if operation completed successfully
	r.AddSuccessWarning(resp, apiResp)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		resp.Diagnostics.AddError(
			"State Update Failed",
			"Unable to save import transport package state",
		)
		return
	}

	tflog.Info(ctx, "Import transport package resource updated successfully")
}

func (r *ImportTransportPackageResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// For acceptance tests only
	if os.Getenv("TF_ACC") == "1" {
		resp.State.RemoveResource(ctx)
		return
	}

	resp.Diagnostics.AddError(
		"Delete Not Supported",
		"Import transport package operations cannot be undone. Please remove the resource manually if required, or contact your administrator.",
	)
}
