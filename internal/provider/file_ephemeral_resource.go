// Copyright (c) Saviynt Inc.
// SPDX-License-Identifier: MPL-2.0

// file_ephemeral_resource.go defines the Terraform Plugin Framework ephemeral resource
// `saviynt_file_connector_ephemeral_resource`. This resource performs a single Open operation
// (no remote create/read/update/delete), loading credentials and connection data from a local file
// into Terraform state at plan time.
//
// On `Open`, the file is parsed via `testutil.LoadConnectorDataForEphemeral`, and each recognized
// key/value is mapped to its corresponding attribute in the model. This ephemeral resource is
// useful for decoupling secret management from provider API calls by sourcing runtime credentials
// from a local file.
package provider

import (
	"context"
	"strings"
	"terraform-provider-Saviynt/util"

	"github.com/hashicorp/terraform-plugin-framework/ephemeral"
	"github.com/hashicorp/terraform-plugin-framework/ephemeral/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ ephemeral.EphemeralResource = &FileCredentialsResource{}

func NewFileCredentialsResource() ephemeral.EphemeralResource {
	return &FileCredentialsResource{}
}

type FileCredentialsModel struct {
	FilePath                 types.String `tfsdk:"file_path"`
	Username                 types.String `tfsdk:"username"`
	Password                 types.String `tfsdk:"password"`
	Change_Pass_Json         types.String `tfsdk:"change_pass_json"`
	Window_Connector_Json    types.String `tfsdk:"windows_connector_json"`
	Azure_Mgmt_Access_Token  types.String `tfsdk:"azure_mgmt_access_token"`
	Client_Id                types.String `tfsdk:"client_id"`
	Client_Secret            types.String `tfsdk:"client_secret"`
	Connection_Json          types.String `tfsdk:"connection_json"`
	Access_Token             types.String `tfsdk:"access_token"`
	Refresh_Token            types.String `tfsdk:"refresh_token"`
	Prov_Password            types.String `tfsdk:"prov_password"`
	Passphrase               types.String `tfsdk:"passphrase"`
	SSH_Key                  types.String `tfsdk:"ssh_key"`
	SSHPassThroughPassword   types.String `tfsdk:"ssh_pass_through_password"`
	SSHPassThroughPassphrase types.String `tfsdk:"ssh_pass_through_passphrase"`
	SSHPassThroughSSHKey     types.String `tfsdk:"ssh_pass_through_ssh_key"`
}

type FileCredentialsResource struct{}

func (r *FileCredentialsResource) Metadata(_ context.Context, req ephemeral.MetadataRequest, resp *ephemeral.MetadataResponse) {
	resp.TypeName = "saviynt_file_connector_ephemeral_resource"
}

func (r *FileCredentialsResource) Schema(ctx context.Context, _ ephemeral.SchemaRequest, resp *ephemeral.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: util.FileEphemeralResourceDescription,
		Attributes: map[string]schema.Attribute{
			"file_path": schema.StringAttribute{
				Required:            true,
				MarkdownDescription: "Path to a JSON (or key-value) file containing credentials.",
			},
			"username": schema.StringAttribute{
				Computed:            true,
				Sensitive:           true,
				MarkdownDescription: "Username read from the file.",
			},
			"password": schema.StringAttribute{
				Computed:            true,
				Sensitive:           true,
				MarkdownDescription: "Password read from the file.",
			},
			"change_pass_json": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "change_pass_json read from the file.",
				Sensitive:           true,
			},
			"windows_connector_json": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "windows_connector_json read from the file.",
				Sensitive:           true,
			},
			"azure_mgmt_access_token": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "azure_mgmt_access_token read from the file.",
				Sensitive:           true,
			},
			"client_id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "client_id read from the file.",
				Sensitive:           true,
			},
			"client_secret": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "client_secret read from the file.",
				Sensitive:           true,
			},
			"connection_json": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "connection_json read from the file.",
				Sensitive:           true,
			},
			"access_token": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "access_token read from the file.",
				Sensitive:           true,
			},
			"refresh_token": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "refresh_token read from the file.",
				Sensitive:           true,
			},
			"prov_password": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "prov_password read from the file.",
				Sensitive:           true,
			},
			"passphrase": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "passphrase read from the file.",
				Sensitive:           true,
			},
			"ssh_key": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "ssh_key read from the file.",
				Sensitive:           true,
			},
			"ssh_pass_through_password": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "ssh_pass_through_password read from the file.",
				Sensitive:           true,
			},
			"ssh_pass_through_passphrase": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "ssh_pass_through_passphrase read from the file.",
				Sensitive:           true,
			},
			"ssh_pass_through_ssh_key": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "ssh_pass_through_ssh_key read from the file.",
				Sensitive:           true,
			},
		},
	}
}

func (r *FileCredentialsResource) Open(ctx context.Context, req ephemeral.OpenRequest, resp *ephemeral.OpenResponse) {
	var data FileCredentialsModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	path := data.FilePath.ValueString()
	createCfg := util.LoadConnectorDataForEphemeral(path)
	for key, val := range createCfg {
		key = strings.TrimSpace(key)
		val = strings.TrimSpace(val)
		switch key {
		case "USERNAME":
			data.Username = types.StringValue(val)
		case "PASSWORD":
			data.Password = types.StringValue(val)
		case "CHANGE_PASS_JSON":
			data.Change_Pass_Json = types.StringValue(val)
		case "WINDOWS_CONNECTOR_JSON":
			data.Window_Connector_Json = types.StringValue(val)
		case "AZURE_MGMT_ACCESS_TOKEN":
			data.Azure_Mgmt_Access_Token = types.StringValue(val)
		case "CLIENT_ID":
			data.Client_Id = types.StringValue(val)
		case "CLIENT_SECRET":
			data.Client_Secret = types.StringValue(val)
		case "CONNECTION_JSON":
			data.Connection_Json = types.StringValue(val)
		case "ACCESS_TOKEN":
			data.Access_Token = types.StringValue(val)
		case "REFRESH_TOKEN":
			data.Refresh_Token = types.StringValue(val)
		case "PROV_PASSWORD":
			data.Prov_Password = types.StringValue(val)
		case "PASSPHRASE":
			data.Passphrase = types.StringValue(val)
		case "SSH_KEY":
			data.SSH_Key = types.StringValue(val)
		case "SSH_PASSTHROUGH_PASSWORD":
			data.SSHPassThroughPassword = types.StringValue(val)
		case "SSH_PASSTHROUGH_PASSPHRASE":
			data.SSHPassThroughPassphrase = types.StringValue(val)
		case "SSH_PASSTHROUGH_SSH_KEY":
			data.SSHPassThroughSSHKey = types.StringValue(val)
		}
	}
	resp.Diagnostics.Append(resp.Result.Set(ctx, data)...)
}
