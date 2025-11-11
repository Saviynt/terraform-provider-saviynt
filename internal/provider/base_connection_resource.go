// Copyright (c) 2025 Saviynt Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// BaseConnector holds all fields common to every connector resource.
type BaseConnectorResourceModel struct {
	ConnectionKey  types.Int64  `tfsdk:"connection_key"`
	ConnectionName types.String `tfsdk:"connection_name"`
	// Description maps to connectionDescription in the API
	Description        types.String `tfsdk:"description"`
	DefaultSavRoles    types.String `tfsdk:"defaultsavroles"`
	EmailTemplate      types.String `tfsdk:"email_template"`
	VaultConnection    types.String `tfsdk:"vault_connection"`
	VaultConfiguration types.String `tfsdk:"vault_configuration"`
	SaveInVault        types.String `tfsdk:"save_in_vault"`
	WriteOnlyVersion   types.String `tfsdk:"wo_version"`
	Msg                types.String `tfsdk:"msg"`
	ErrorCode          types.String `tfsdk:"error_code"`
}

func BaseConnectorResourceSchema() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"connection_key": schema.Int64Attribute{
			Computed:    true,
			Description: "Unique identifier of the connection returned by the API. Example: 1909",
		},
		"connection_name": schema.StringAttribute{
			Required:    true,
			Description: "Name of the connection. Example: \"Active Directory_Doc\"",
		},
		"description": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "Description for the connection. Example: \"ORG_AD\"",
		},
		"defaultsavroles": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "Default SAV roles for managing the connection. Example: \"ROLE_ORG\"",
		},
		"email_template": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "Email template for notifications. Example: \"New Account Task Creation\"",
		},
		"vault_connection": schema.StringAttribute{
			Optional:    true,
			Description: "Specifies the type of vault connection being used (e.g., 'Hashicorp'). Example: \"Hashicorp\"",
		},
		"vault_configuration": schema.StringAttribute{
			Optional:    true,
			Description: "JSON string specifying vault configuration.",
		},
		"save_in_vault": schema.StringAttribute{
			Optional:    true,
			Description: "Flag indicating whether the encrypted attribute should be saved in the configured vault. Example: \"false\"",
		},
		"wo_version": schema.StringAttribute{
			Optional:    true,
			Description: "Add/change the value of this attribute to update the writeonly attributes like username, password etc in connection resources",
		},
		"msg": schema.StringAttribute{
			Computed:    true,
			Description: "A message indicating the outcome of the operation.",
		},
		"error_code": schema.StringAttribute{
			Computed:    true,
			Description: "An error code where '0' signifies success and '1' signifies an unsuccessful operation.",
		},
	}
}
