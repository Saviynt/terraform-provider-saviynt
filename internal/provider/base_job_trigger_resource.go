// Copyright (c) 2025 Saviynt Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// BaseJobTriggerResourceModel holds all fields common to every job trigger resource that uses CreateTrigger API.
// This is different from BaseJobControlResourceModel which is for CreateOrUpdateTriggers API.
type BaseJobTriggerResourceModel struct {
	Name     types.String `tfsdk:"name"`
	JobName  types.String `tfsdk:"job_name"`
	JobGroup types.String `tfsdk:"job_group"`
	Group    types.String `tfsdk:"group"`
	CronExp  types.String `tfsdk:"cron_exp"`
}

func BaseJobTriggerResourceSchema() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"name": schema.StringAttribute{
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
		"group": schema.StringAttribute{
			Required:    true,
			Description: "Group classification for the trigger. Example: \"GRAILS_JOBS\"",
		},
		"cron_exp": schema.StringAttribute{
			Required:    true,
			Description: "Cron expression defining the schedule for the trigger. Example: \"0 0 2 * * ?\"",
		},
	}
}
