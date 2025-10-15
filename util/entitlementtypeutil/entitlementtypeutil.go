// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package entitlementtypeutil

import (
	"encoding/json"
	"log"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// "strings"

func ParseCreateTaskActionForState(raw *string) []attr.Value {
	if raw == nil || *raw == "" {
		return nil
	}

	var parsed map[string][]string
	err := json.Unmarshal([]byte(*raw), &parsed)
	if err != nil {
		log.Print("Error parsing create task action: ", err)
		return nil
	}

	taskActions := parsed["taskActions"]
	result := make([]attr.Value, len(taskActions))
	for i, v := range taskActions {
		result[i] = types.StringValue(v)
	}

	return result
}

var requestFormDisplayToUpdateValue = map[string]string{
	"Request Form NotRequestable Single":   "SHOW_BUT_NOTREUESTABLESINGLE",
	"Request Form NotRequestable Multiple": "SHOW_BUT_NOTREUESTABLEMULTIPLE",
	"Request Form None":                    "NONE",
	"Request Form Single":                  "SINGLE",
	"Request Form Multiple":                "MULTIPLE",
	"Request Form Table":                   "TABLE",
	"Request Form Free From Text":          "FREEFORMTEXT",
	"Request Form Table No Remove":         "TABLENOREMOVE",
	"Request Form Radio Button":            "RADIOBUTN",
	"Request Form CheckBox":                "CHECKBOXN",
	"Request Form Read Only Table":         "READONLYTABLE",
	// "NONE_BUT_CREATETASK" (int value 9) is not explicitly mapped in response, so not included
}

func TranslateValueWithDefault(input string) *string {
	return TranslateValue(input, nil)
}

func TranslateValue(input string, valueMap map[string]string) *string {
	if valueMap == nil {
		valueMap = requestFormDisplayToUpdateValue
	}
	if input == "" {
		return nil
	}
	if val, ok := valueMap[input]; ok {
		return &val
	}
	return &input
}

// ConvertHierarchyRequiredForUpdate converts the HierarchyRequired field for update requests
// Create API expects integers (1, 0) but Update API expects strings ("true", "false")
func ConvertHierarchyRequiredForUpdate(hierarchyRequired types.String) *string {
	if hierarchyRequired.IsNull() || hierarchyRequired.IsUnknown() {
		return nil
	}

	value := hierarchyRequired.ValueString()
	switch value {
	case "1":
		result := "true"
		return &result
	case "0":
		result := "false"
		return &result
	default:
		// If it's already "true" or "false", pass it through
		if value == "true" || value == "false" {
			return &value
		}
		// Default to "false" for any other value
		result := "false"
		return &result
	}
}

// ConvertBoolToStringForUpdate converts boolean fields to strings for update requests
// Some fields expect boolean in create API but strings ("true"/"false") in update API
func ConvertBoolToStringForUpdate(boolValue types.Bool) *string {
	if boolValue.IsNull() || boolValue.IsUnknown() {
		return nil
	}

	if boolValue.ValueBool() {
		result := "true"
		return &result
	} else {
		result := "false"
		return &result
	}
}
