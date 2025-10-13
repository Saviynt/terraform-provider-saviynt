// Copyright (c) 2025 Saviynt Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// BaseJobControlResourceModel holds all fields common to every job control resource.
type BaseJobControlResourceModel struct {
	TriggerName    types.String `tfsdk:"trigger_name"`
	JobName        types.String `tfsdk:"job_name"`
	JobGroup       types.String `tfsdk:"job_group"`
	TriggerGroup   types.String `tfsdk:"trigger_group"`
	CronExpression types.String `tfsdk:"cron_expression"`
}

func BaseJobControlResourceSchema() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"trigger_name": schema.StringAttribute{
			Required:    true,
			Description: "Unique name of the trigger. Example: \"MyTrigger_001\"",
		},
		"job_name": schema.StringAttribute{
			Required:    true,
			Description: "Name of the job associated with the trigger. Example: \"WSRetryJob\"",
		},
		"job_group": schema.StringAttribute{
			Required:    true,
			Description: "Name of the job group associated with the trigger. Example: \"utility\"",
		},
		"trigger_group": schema.StringAttribute{
			Optional:    true,
			Description: "Group classification for the trigger. Example: \"GRAILS_JOBS\"",
		},
		"cron_expression": schema.StringAttribute{
			Required:    true,
			Description: "Cron expression defining the schedule for the trigger. Example: \"0 0 2 * * ?\"",
		},
	}
}
