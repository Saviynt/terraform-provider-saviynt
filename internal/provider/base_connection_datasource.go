// Copyright (c) Saviynt Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// BaseConnector holds all fields common to every connector datasource
type BaseConnectionDataSourceModel struct {
	ConnectionName  types.String `tfsdk:"connection_name"`
	ConnectionKey   types.Int64  `tfsdk:"connection_key"`
	Description     types.String `tfsdk:"description"`
	DefaultSavRoles types.String `tfsdk:"default_sav_roles"`
	EmailTemplate   types.String `tfsdk:"email_template"`
	ConnectionType  types.String `tfsdk:"connection_type"`
	CreatedOn       types.String `tfsdk:"created_on"`
	CreatedBy       types.String `tfsdk:"created_by"`
	UpdatedBy       types.String `tfsdk:"updated_by"`
	Status          types.Int64  `tfsdk:"status"`
	Msg             types.String `tfsdk:"msg"`
	ErrorCode       types.Int64  `tfsdk:"error_code"`
}

func BaseConnectorDataSourceSchema() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"connection_name": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "The name of the connection.",
		},
		"connection_key": schema.Int64Attribute{
			Optional:    true,
			Computed:    true,
			Description: "The key of the connection.",
		},
		"description": schema.StringAttribute{
			Computed: true,
		},
		"default_sav_roles": schema.StringAttribute{
			Computed: true,
		},
		"email_template": schema.StringAttribute{
			Computed: true,
		},
		"connection_type": schema.StringAttribute{
			Computed: true,
		},
		"created_on": schema.StringAttribute{
			Computed: true,
		},
		"created_by": schema.StringAttribute{
			Computed: true,
		},
		"updated_by": schema.StringAttribute{
			Computed: true,
		},
		"status": schema.Int64Attribute{
			Computed: true,
		},
		"msg": schema.StringAttribute{
			Computed: true,
		},
		"error_code": schema.Int64Attribute{
			Computed: true,
		},
	}
}
