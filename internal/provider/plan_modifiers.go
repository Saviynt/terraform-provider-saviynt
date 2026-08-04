// Copyright (c) 2025 Saviynt Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"sort"
	"strings"

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
