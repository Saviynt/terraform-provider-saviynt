/*
 * Copyright (c) 2025 Saviynt Inc.
 * All Rights Reserved.
 *
 * This software is the confidential and proprietary information of
 * Saviynt Inc. ("Confidential Information"). You shall not disclose,
 * use, or distribute such Confidential Information except in accordance
 * with the terms of the license agreement you entered into with Saviynt.
 *
 * SAVIYNT MAKES NO REPRESENTATIONS OR WARRANTIES ABOUT THE SUITABILITY OF
 * THE SOFTWARE, EITHER EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO
 * THE IMPLIED WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR
 * PURPOSE, OR NON-INFRINGEMENT.
 */

// env_ephemeral_resource.go defines the Terraform Plugin Framework ephemeral resource
// `saviynt_env_ephemeral_resource`. This resource performs a single Open operation
// (no remote create/read/update/delete), loading credentials and connection data from a env
// into Terraform state at plan time.
package provider

import (
	"context"
	"os"

	// "fmt"
	// "os"

	"github.com/hashicorp/terraform-plugin-framework/ephemeral"
	"github.com/hashicorp/terraform-plugin-framework/ephemeral/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ ephemeral.EphemeralResource = &EnvCredentialsResource{}

func NewEnvCredentialsResource() ephemeral.EphemeralResource {
	return &EnvCredentialsResource{}
}

type EnvCredentialsModel struct {
	Svnt_Username                 types.String `tfsdk:"svnt_username"`
	Svnt_Password                 types.String `tfsdk:"svnt_password"`
	Svnt_Azure_Mgmt_Access_Token  types.String `tfsdk:"svnt_azure_mgmt_access_token"`
	Svnt_Client_Id                types.String `tfsdk:"svnt_client_id"`
	Svnt_Client_Secret            types.String `tfsdk:"svnt_client_secret"`
	Svnt_Access_Token             types.String `tfsdk:"svnt_access_token"`
	Svnt_Refresh_Token            types.String `tfsdk:"svnt_refresh_token"`
	Svnt_Prov_Password            types.String `tfsdk:"svnt_prov_password"`
	Svnt_Passphrase               types.String `tfsdk:"svnt_passphrase"`
	Svnt_SSH_Key                  types.String `tfsdk:"svnt_ssh_key"`
	Svnt_SSHPassThroughPassword   types.String `tfsdk:"svnt_ssh_pass_through_password"`
	Svnt_SSHPassThroughPassphrase types.String `tfsdk:"svnt_ssh_pass_through_passphrase"`
	Svnt_SSHPassThroughSSHKey     types.String `tfsdk:"svnt_ssh_pass_through_ssh_key"`
}

type EnvCredentialsResource struct{}

func (r *EnvCredentialsResource) Metadata(_ context.Context, req ephemeral.MetadataRequest, resp *ephemeral.MetadataResponse) {
	resp.TypeName = "saviynt_env_ephemeral_resource"
}

func (r *EnvCredentialsResource) Schema(ctx context.Context, _ ephemeral.SchemaRequest, resp *ephemeral.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Read username/password from a local file",
		Attributes: map[string]schema.Attribute{
			"svnt_username": schema.StringAttribute{
				Computed:            true,
				Sensitive:           true,
				MarkdownDescription: "Username read from the file.",
			},
			"svnt_password": schema.StringAttribute{
				Computed:            true,
				Sensitive:           true,
				MarkdownDescription: "Password read from the file.",
			},
			"svnt_azure_mgmt_access_token": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "azure_mgmt_access_token read from the file.",
				Sensitive:           true,
			},
			"svnt_client_id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "client_id read from the file.",
				Sensitive:           true,
			},
			"svnt_client_secret": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "client_secret read from the file.",
				Sensitive:           true,
			},
			"svnt_connection_json": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "connection_json read from the file.",
				Sensitive:           true,
			},
			"svnt_access_token": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "access_token read from the file.",
				Sensitive:           true,
			},
			"svnt_refresh_token": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "refresh_token read from the file.",
				Sensitive:           true,
			},
			"svnt_prov_password": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "prov_password read from the file.",
				Sensitive:           true,
			},
			"svnt_passphrase": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "passphrase read from the file.",
				Sensitive:           true,
			},
			"svnt_ssh_key": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "ssh_key read from the file.",
				Sensitive:           true,
			},
			"svnt_ssh_pass_through_password": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "ssh_pass_through_password read from the file.",
				Sensitive:           true,
			},
			"svnt_ssh_pass_through_passphrase": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "ssh_pass_through_passphrase read from the file.",
				Sensitive:           true,
			},
			"svnt_ssh_pass_through_ssh_key": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "ssh_pass_through_ssh_key read from the file.",
				Sensitive:           true,
			},
		},
	}
}

func (r *EnvCredentialsResource) Open(ctx context.Context, req ephemeral.OpenRequest, resp *ephemeral.OpenResponse) {
	var data EnvCredentialsModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	data.Svnt_Username = types.StringValue(os.Getenv("svnt_username"))
	data.Svnt_Password = types.StringValue(os.Getenv("svnt_password"))
	data.Svnt_Azure_Mgmt_Access_Token = types.StringValue(os.Getenv("svnt_azure_mgmt_access_token"))
	data.Svnt_Client_Id = types.StringValue(os.Getenv("svnt_client_id"))
	data.Svnt_Client_Secret = types.StringValue(os.Getenv("svnt_client_secret"))
	data.Svnt_Access_Token = types.StringValue(os.Getenv("svnt_access_token"))
	data.Svnt_Refresh_Token = types.StringValue(os.Getenv("svnt_refresh_token"))
	data.Svnt_Prov_Password = types.StringValue(os.Getenv("svnt_prov_password"))
	data.Svnt_Passphrase = types.StringValue(os.Getenv("svnt_passphrase"))
	data.Svnt_SSH_Key = types.StringValue(os.Getenv("svnt_ssh_key"))
	data.Svnt_SSHPassThroughPassword = types.StringValue(os.Getenv("svnt_ssh_pass_through_password"))
	data.Svnt_SSHPassThroughPassphrase = types.StringValue(os.Getenv("svnt_ssh_pass_through_passphrase"))
	data.Svnt_SSHPassThroughSSHKey = types.StringValue(os.Getenv("svnt_ssh_pass_through_ssh_key"))
	resp.Diagnostics.Append(resp.Result.Set(ctx, data)...)
}
