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
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ConnectionTimeoutConfig struct {
	RetryWait               types.Int64 `tfsdk:"retry_wait"`
	TokenRefreshMaxTryCount types.Int64 `tfsdk:"token_refresh_max_try_count"`
	RetryWaitMaxValue       types.Int64 `tfsdk:"retry_wait_max_value"`
	RetryCount              types.Int64 `tfsdk:"retry_count"`
	ReadTimeout             types.Int64 `tfsdk:"read_timeout"`
	ConnectionTimeout       types.Int64 `tfsdk:"connection_timeout"`
	RetryFailureStatusCode  types.Int64 `tfsdk:"retry_failure_status_code"`
}

func ConnectionTimeoutConfigeSchema() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"retry_wait":                  schema.Int64Attribute{Computed: true},
		"token_refresh_max_try_count": schema.Int64Attribute{Computed: true},
		"retry_wait_max_value":        schema.Int64Attribute{Computed: true},
		"retry_count":                 schema.Int64Attribute{Computed: true},
		"read_timeout":                schema.Int64Attribute{Computed: true},
		"connection_timeout":          schema.Int64Attribute{Computed: true},
		"retry_failure_status_code":   schema.Int64Attribute{Computed: true},
	}
}
