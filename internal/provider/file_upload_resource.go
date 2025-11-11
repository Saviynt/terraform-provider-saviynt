// Copyright (c) 2025 Saviynt Inc.
// SPDX-License-Identifier: MPL-2.0

// saviynt_file_upload_resource manages file uploads to the Saviynt instance.
// The resource implements the full Terraform lifecycle:
//   - Create: uploads a file to the specified location in Saviynt.
//   - Read: maintains the current state.
//   - Update: re-uploads the file if configuration changes.
//   - Delete: removes from Terraform state only (files remain in Saviynt).
//
// Supported file types and locations:
//   - CSV files: uploaded to "Datafiles" directory
//   - SAV files: uploaded to "SAV" directory
package provider

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"terraform-provider-Saviynt/internal/client"
	"terraform-provider-Saviynt/util"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	openapi "github.com/saviynt/saviynt-api-go-client/filedirectory"
)

var _ resource.Resource = &FileUploadResource{}

type FileUploadResource struct {
	client            client.SaviyntClientInterface
	token             string
	provider          client.SaviyntProviderInterface
	uploadFileFactory client.FileFactoryInterface
}

func NewFileUploadResource() resource.Resource {
	return &FileUploadResource{
		uploadFileFactory: &client.DefaultFileFactory{},
	}
}

func NewFileUploadResourceWithFactory(factory client.FileFactoryInterface) resource.Resource {
	return &FileUploadResource{
		uploadFileFactory: factory,
	}
}

type FileUploadResourceModel struct {
	ID           types.String `tfsdk:"id"`
	FilePath     types.String `tfsdk:"file_path"`
	PathLocation types.String `tfsdk:"path_location"`
	FileVersion  types.String `tfsdk:"file_version"`
	Message      types.String `tfsdk:"message"`
	ErrorCode    types.String `tfsdk:"error_code"`
}

func (r *FileUploadResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_file_upload_resource"
}

func (r *FileUploadResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: util.FileUploadDescription,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Unique identifier for the file upload",
			},
			"file_path": schema.StringAttribute{
				Required:            true,
				MarkdownDescription: "Path to the file to upload",
			},
			"path_location": schema.StringAttribute{
				Required:            true,
				MarkdownDescription: "Upload location: 'Datafiles' or 'SAV'",
			},
			"file_version": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "File version identifier. This acts as a change trigger - increment this value when you need to re-upload the same file with modifications.",
			},
			"message": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Response message from the upload",
			},
			"error_code": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Error code from the upload response",
			},
		},
	}
}

// Configure initializes the File upload resource with the provider's API client and access token.
func (r *FileUploadResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	tflog.Debug(ctx, "Starting File upload resource configuration")

	// Check if provider data is available.
	if req.ProviderData == nil {
		tflog.Debug(ctx, "ProviderData is nil, returning early")
		return
	}

	// Cast provider data to your provider type.
	prov, ok := req.ProviderData.(*SaviyntProvider)
	if !ok {
		tflog.Error(ctx, "Provider configuration failed - expected *SaviyntProvider, got different type")
		resp.Diagnostics.AddError(
			"Unexpected Provider Data",
			"Expected *SaviyntProvider, got different type",
		)
		return
	}

	// Set the client and token from the provider state using interface wrapper.
	r.client = &client.SaviyntClientWrapper{Client: prov.client}
	r.token = prov.accessToken
	r.provider = &client.SaviyntProviderWrapper{Provider: prov} // Store provider reference for retry logic
	tflog.Debug(ctx, "File upload resource configured successfully")
}

// SetClient sets the client for testing purposes
func (r *FileUploadResource) SetClient(client client.SaviyntClientInterface) {
	r.client = client
}

// SetToken sets the token for testing purposes
func (r *FileUploadResource) SetToken(token string) {
	r.token = token
}

// SetProvider sets the provider for testing purposes
func (r *FileUploadResource) SetProvider(provider client.SaviyntProviderInterface) {
	r.provider = provider
}

func (r *FileUploadResource) VerifyFilePathLocationCompatibility(ctx context.Context, fPath string, pathLocation string) error {
	fileExt := strings.ToLower(filepath.Ext(fPath))

	FilePathcompatibilityMap := map[string]string{
		".csv": "Datafiles",
		".sav": "SAV",
	}

	expectedLocation, exists := FilePathcompatibilityMap[fileExt]
	if !exists {
		return fmt.Errorf("unsupported file extension '%s'. Only .csv and .sav files are allowed", fileExt)
	}

	if pathLocation != expectedLocation {
		return fmt.Errorf("invalid path_location '%s' for file extension '%s'. Expected '%s'", pathLocation, fileExt, expectedLocation)
	}

	return nil
}

// ValidateFilePath validates the file path to prevent path traversal attacks
func (r *FileUploadResource) ValidateFilePath(filePath string) error {
	cleanPath := filepath.Clean(filePath)
	if strings.Contains(cleanPath, "..") {
		return fmt.Errorf("path traversal not allowed: %s", filePath)
	}
	return nil
}

func (r *FileUploadResource) UploadFile(ctx context.Context, plan FileUploadResourceModel, op string) (*openapi.UploadSchemaFileResponse, error) {
	filePath := plan.FilePath.ValueString()
	pathLocation := plan.PathLocation.ValueString()

	// Validate file path and location compatibility
	if err := r.VerifyFilePathLocationCompatibility(ctx, filePath, pathLocation); err != nil {
		return nil, err
	}

	tflog.Debug(ctx, "Uploading file", map[string]interface{}{
		"file_path":     filePath,
		"path_location": pathLocation,
	})

	// Validate file path for security
	if err := r.ValidateFilePath(filePath); err != nil {
		return nil, err
	}

	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open file '%s': %s", filePath, err.Error())
	}
	defer file.Close()

	var uploadResp *openapi.UploadSchemaFileResponse

	err = r.provider.AuthenticatedAPICallWithRetry(ctx, op, func(token string) error {
		fileuploadOps := r.uploadFileFactory.CreateFileOperations(r.client.APIBaseURL(), token)
		apiResp, httpResp, apiErr := fileuploadOps.UploadSchemaFile(ctx, file, pathLocation)

		if httpResp != nil {
			if httpResp.StatusCode == 401 {
				return fmt.Errorf("authentication failed: invalid or expired token")
			}
			if httpResp.StatusCode >= 400 {
				return fmt.Errorf("API request failed with status %d", httpResp.StatusCode)
			}
		}

		uploadResp = apiResp
		return apiErr
	})

	if err != nil {
		return nil, fmt.Errorf("file upload failed for '%s': %s", filepath.Base(filePath), err.Error())
	}

	uploadJson, _ := json.Marshal(uploadResp)
	log.Printf("[DEBUG] Upload resp: %s", string(uploadJson))

	return uploadResp, nil
}

func (r *FileUploadResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan FileUploadResourceModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	uploadResp, err := r.UploadFile(ctx, plan, "upload_file")
	if err != nil {
		resp.Diagnostics.AddError(
			"File Upload Failed",
			fmt.Sprintf("Error uploading file: %s", err.Error()),
		)
		return
	}

	// Set computed values
	plan.ID = types.StringValue(fmt.Sprintf("%s_%s", filepath.Base(plan.FilePath.ValueString()), plan.PathLocation.ValueString()))
	plan.Message = types.StringValue(uploadResp.GetMsg())
	plan.ErrorCode = types.StringValue(uploadResp.GetErrorCode())

	tflog.Debug(ctx, "File uploaded successfully", map[string]interface{}{
		"message":    uploadResp.GetMsg(),
		"error_code": uploadResp.GetErrorCode(),
	})

	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *FileUploadResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data FileUploadResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// File upload is a one-time operation, so we just keep the state as is
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *FileUploadResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan FileUploadResourceModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	uploadResp, err := r.UploadFile(ctx, plan, "upload_file_update")
	if err != nil {
		resp.Diagnostics.AddError(
			"Updated File Upload Failed",
			fmt.Sprintf("Error uploading updated file: %s", err.Error()),
		)
		return
	}

	// Set computed values
	plan.ID = types.StringValue(fmt.Sprintf("%s_%s", filepath.Base(plan.FilePath.ValueString()), plan.PathLocation.ValueString()))
	plan.Message = types.StringValue(uploadResp.GetMsg())
	plan.ErrorCode = types.StringValue(uploadResp.GetErrorCode())

	tflog.Debug(ctx, "Updated file uploaded successfully", map[string]interface{}{
		"message":    uploadResp.GetMsg(),
		"error_code": uploadResp.GetErrorCode(),
	})

	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *FileUploadResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// File upload doesn't support deletion - just remove from state
	resp.Diagnostics.AddWarning(
		"Removed the resource block from state",
		"Deleting the File upload resource only removes it from Terraform state; files remain in Saviynt. Files need to be manually deleted from the EIC UI.",
	)

	tflog.Debug(ctx, "File upload resource deleted from state")
	resp.State.RemoveResource(ctx)
}
