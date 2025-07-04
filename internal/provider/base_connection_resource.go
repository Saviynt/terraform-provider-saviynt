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

package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// BaseConnector holds all fields common to every connector resource.
type BaseConnectorResourceModel struct {
	ConnectionKey      types.Int64  `tfsdk:"connection_key"`
	ConnectionName     types.String `tfsdk:"connection_name"`
	ConnectionType     types.String `tfsdk:"connection_type"`
	DefaultSavRoles    types.String `tfsdk:"defaultsavroles"`
	EmailTemplate      types.String `tfsdk:"email_template"`
	VaultConnection    types.String `tfsdk:"vault_connection"`
	VaultConfiguration types.String `tfsdk:"vault_configuration"`
	SaveInVault        types.String `tfsdk:"save_in_vault"`
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
		"connection_type": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "Connection type (e.g., 'AD' for Active Directory). Example: \"AD\"",
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
			Description: "JSON string specifying vault configuration. Example: '{\"path\":\"/secrets/data/kv-dev-intgn1/-AD_Credential\",\"keyMapping\":{\"PASSWORD\":\"AD_PASSWORD~#~None\"}}'",
		},
		"save_in_vault": schema.StringAttribute{
			Optional:    true,
			Description: "Flag indicating whether the encrypted attribute should be saved in the configured vault. Example: \"false\"",
		},
		"msg": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "A message indicating the outcome of the operation.",
		},
		"error_code": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "An error code where '0' signifies success and '1' signifies an unsuccessful operation.",
		},
	}
}
