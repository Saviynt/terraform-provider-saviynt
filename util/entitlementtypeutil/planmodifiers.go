// Copyright (c) Saviynt Inc.
// SPDX-License-Identifier: MPL-2.0

package entitlementtypeutil

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// RequireWorkflowWhenEnabled is a plan modifier that validates that workflow is set
// when enable_entitlement_to_role_sync is set to true
func RequireWorkflowWhenEnabled() planmodifier.Bool {
	return requireWorkflowWhenEnabledModifier{}
}

// requireWorkflowWhenEnabledModifier implements the plan modifier
type requireWorkflowWhenEnabledModifier struct{}

// Description returns a human-readable description of the plan modifier.
func (m requireWorkflowWhenEnabledModifier) Description(_ context.Context) string {
	return "Validates that workflow is set when enable_entitlement_to_role_sync is enabled"
}

// MarkdownDescription returns a markdown description of the plan modifier.
func (m requireWorkflowWhenEnabledModifier) MarkdownDescription(ctx context.Context) string {
	return m.Description(ctx)
}

// PlanModifyBool implements the plan modification logic.
func (m requireWorkflowWhenEnabledModifier) PlanModifyBool(ctx context.Context, req planmodifier.BoolRequest, resp *planmodifier.BoolResponse) {
	// If the planned value is null, unknown, or false, no validation needed
	if req.PlanValue.IsNull() || req.PlanValue.IsUnknown() || !req.PlanValue.ValueBool() {
		return
	}

	// If enable_entitlement_to_role_sync is being set to true, check if workflow is set
	if req.PlanValue.ValueBool() {
		// Get the workflow attribute from the plan
		var workflowValue types.String
		diags := req.Plan.GetAttribute(ctx, path.Root("workflow"), &workflowValue)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

		// Check if workflow is null, unknown, or empty
		if workflowValue.IsNull() || workflowValue.IsUnknown() || workflowValue.ValueString() == "" {
			resp.Diagnostics.AddAttributeError(
				req.Path,
				"Invalid Configuration",
				fmt.Sprintf("enable_entitlement_to_role_sync can only be set to true when workflow is also specified. "+
					"Please provide a value for the workflow attribute."),
			)
			return
		}
	}
}

// CreateTaskActionDefault is a plan modifier that handles state management for createTaskAction
// This attribute is update-only in the API, so we need special handling for defaults and state consistency
func CreateTaskActionDefault() planmodifier.Set {
	return createTaskActionDefaultModifier{}
}

// createTaskActionDefaultModifier implements the plan modifier for createTaskAction
type createTaskActionDefaultModifier struct{}

// Description returns a human-readable description of the plan modifier.
func (m createTaskActionDefaultModifier) Description(_ context.Context) string {
	return "Handles state management for createTaskAction including default values and update-only constraints"
}

// MarkdownDescription returns a markdown description of the plan modifier.
func (m createTaskActionDefaultModifier) MarkdownDescription(ctx context.Context) string {
	return m.Description(ctx)
}

// PlanModifySet implements the plan modification logic for Set attributes.
func (m createTaskActionDefaultModifier) PlanModifySet(ctx context.Context, req planmodifier.SetRequest, resp *planmodifier.SetResponse) {
	// If the plan value is null or unknown, we need to determine the appropriate value
	if req.PlanValue.IsNull() || req.PlanValue.IsUnknown() {
		// Check if this is a create operation (no prior state)
		if req.StateValue.IsNull() {
			// For create operations, set default value that will be applied during follow-up update
			defaultValue := types.StringValue("noAction")
			defaultSet := basetypes.NewSetValueMust(types.StringType, []attr.Value{defaultValue})
			resp.PlanValue = defaultSet
			return
		}

		// For updates where the value becomes null/unknown, preserve the existing state value
		resp.PlanValue = req.StateValue
		return
	}

	// If the plan has a value, use it as-is (whether empty set or populated)
	resp.PlanValue = req.PlanValue
}
