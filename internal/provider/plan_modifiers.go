// Copyright (c) 2025 Saviynt Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"context"
	"sort"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// PreserveOrderIfSemanticallyEqual compares the current state value with the new API value.
// If both contain the same comma-separated items (regardless of order), it returns
// the current state value to prevent unnecessary drift. Otherwise, it returns the new value.
// This prevents Terraform from showing a diff when only the ordering changed.
func PreserveOrderIfSemanticallyEqual(currentState types.String, newValue types.String) types.String {
	if currentState.IsNull() || currentState.IsUnknown() {
		return newValue
	}
	if newValue.IsNull() || newValue.IsUnknown() {
		return newValue
	}

	currentSorted := sortedCSV(currentState.ValueString())
	newSorted := sortedCSV(newValue.ValueString())

	if currentSorted == newSorted {
		// Same roles, different order — keep the current state to avoid drift
		return currentState
	}

	// Actually different values — return the new one
	return newValue
}

// sortedCSV splits a comma-separated string, trims whitespace, sorts, and rejoins.
func sortedCSV(s string) string {
	if s == "" {
		return ""
	}
	parts := strings.Split(s, ",")
	for i, p := range parts {
		parts[i] = strings.TrimSpace(p)
	}
	sort.Strings(parts)
	return strings.Join(parts, ",")
}

// connectionStableComputedAttributes are computed-only connection attributes
// whose value never changes once the connection exists.
var connectionStableComputedAttributes = map[string]bool{
	"id":             true,
	"connection_key": true,
}

// WithPriorStateForUnset keeps the prior value of connection attributes that the
// configuration does not set, instead of planning them as unknown.
//
// Without it, any change to a connection marks every Optional+Computed
// attribute that the configuration does not set as "(known after apply)", so
// `terraform plan` prints the old value of each one (often large JSON
// documents) next to the attribute that actually changed. Those attributes do
// not change during the update, so keeping the prior state value gives a plan
// that shows only the real change.
//
// It adds:
//   - UnsetKeepsPriorState to Optional+Computed string attributes that are not
//     sensitive and do not already declare plan modifiers, and
//   - UseStateForUnknown to the computed-only identifiers listed in
//     connectionStableComputedAttributes.
//
// Sensitive attributes are left alone so a value the API might mask on read
// is never planned back into an update request.
func WithPriorStateForUnset(attrs map[string]schema.Attribute) map[string]schema.Attribute {
	for name, attr := range attrs {
		switch a := attr.(type) {
		case schema.StringAttribute:
			if len(a.PlanModifiers) > 0 {
				continue
			}
			switch {
			case a.Optional && a.Computed && !a.Sensitive:
				a.PlanModifiers = []planmodifier.String{UnsetKeepsPriorState()}
				attrs[name] = a
			case a.Computed && !a.Optional && connectionStableComputedAttributes[name]:
				a.PlanModifiers = []planmodifier.String{stringplanmodifier.UseStateForUnknown()}
				attrs[name] = a
			}
		case schema.Int64Attribute:
			if len(a.PlanModifiers) == 0 && a.Computed && !a.Optional && connectionStableComputedAttributes[name] {
				a.PlanModifiers = []planmodifier.Int64{int64planmodifier.UseStateForUnknown()}
				attrs[name] = a
			}
		}
	}
	return attrs
}

// UnsetKeepsPriorState plans an Optional+Computed attribute that is not set in
// the configuration with its current state value on update, including when
// that value is null.
//
// stringplanmodifier.UseStateForUnknown skips null state values, so an
// attribute the API does not return would still show up as
// "(known after apply)" on every update. The provider omits null attributes
// from update requests, so they stay null after the update as well.
func UnsetKeepsPriorState() planmodifier.String {
	return unsetKeepsPriorStateModifier{}
}

type unsetKeepsPriorStateModifier struct{}

func (m unsetKeepsPriorStateModifier) Description(_ context.Context) string {
	return "When not set in the configuration, the value stays the same as the current state."
}

func (m unsetKeepsPriorStateModifier) MarkdownDescription(ctx context.Context) string {
	return m.Description(ctx)
}

func (m unsetKeepsPriorStateModifier) PlanModifyString(_ context.Context, req planmodifier.StringRequest, resp *planmodifier.StringResponse) {
	// Creating the resource: let the provider compute the value.
	if req.State.Raw.IsNull() {
		return
	}
	// The resource is being destroyed, or the value is already known.
	if req.Plan.Raw.IsNull() || !req.PlanValue.IsUnknown() {
		return
	}
	// Only attributes the configuration leaves unset keep their prior value.
	if !req.ConfigValue.IsNull() {
		return
	}
	resp.PlanValue = req.StateValue
}
