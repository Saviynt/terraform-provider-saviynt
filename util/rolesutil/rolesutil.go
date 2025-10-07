// Copyright (c) 2025 Saviynt Inc.
// SPDX-License-Identifier: MPL-2.0

package rolesutil

var RoleTypeMap = map[string]string{
	"1": "ENABLER",
	"2": "TRANSACTIONAL",
	"3": "FIREFIGHTER",
	"4": "ENTERPRISE",
	"5": "APPLICATION",
	"6": "ENTITLEMENT",
}

var StatusMap = map[string]string{
	"0": "Inactive",
	"1": "Active",
}

var SoxCriticalityMap = map[string]string{
	"0": "None",
	"1": "Very Low",
	"2": "Low",
	"3": "Medium",
	"4": "High",
	"5": "Critical",
}

var SysCriticalMap = map[string]string{
	"0": "None",
	"1": "Very Low",
	"2": "Low",
	"3": "Medium",
	"4": "High",
	"5": "Critical",
}

var PrivilegedMap = map[string]string{
	"0": "None",
	"1": "Very Low",
	"2": "Low",
	"3": "Medium",
	"4": "High",
	"5": "Critical",
}

var ConfidentialityMap = map[string]string{
	"0": "None",
	"1": "Very Low",
	"2": "Low",
	"3": "Medium",
	"4": "High",
	"5": "Critical",
}

// ReverseTranslateValue converts human-readable values back to numeric keys using the provided map
func ReverseTranslateValue(input string, valueMap map[string]string) string {
	if input == "" {
		return ""
	}
	// Search for the input value in the map values and return the corresponding key
	for key, value := range valueMap {
		if value == input {
			return key
		}
	}
	// If not found in map, return the input as-is (might already be numeric)
	return input
}

var RiskMap = map[string]string{
	"0": "None",
	"1": "Very Low",
	"2": "Low",
	"3": "Medium",
	"4": "High",
	"5": "Critical",
}
