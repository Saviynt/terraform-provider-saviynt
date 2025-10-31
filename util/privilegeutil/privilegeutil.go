// Copyright (c) 2025 Saviynt Inc.
// SPDX-License-Identifier: MPL-2.0

package privilegeutil

// AttributeTypeMap maps API response values to Terraform configuration values
// API returns: "String", "Boolean", "Multiple Select From List", etc.
// Terraform expects: "STRING", "BOOLEAN", "MULTIPLE SELECT FROM LIST", etc.
var AttributeTypeMap = map[string]string{
	"String":                         "STRING",
	"Boolean":                        "BOOLEAN",
	"Number":                         "NUMBER",
	"Multiple Select From List":      "MULTIPLE SELECT FROM LIST",
	"Single Select From List":        "SINGLE SELECT FROM LIST",
	"Multiple Select From SQL Query": "MULTIPLE SELECT FROM SQL QUERY",
	"Single Select From SQL Query":   "SINGLE SELECT FROM SQL QUERY",
	"Password":                       "PASSWORD",
	"Large Text":                     "LARGE TEXT",
	"Check Box":                      "CHECK BOX",
	"Date":                           "DATE",
	"Enum":                           "ENUM",
}

// ReverseAttributeTypeMap maps Terraform configuration values to API response values
// Used when we need to convert from Terraform format to API format
var ReverseAttributeTypeMap = map[string]string{
	"STRING":                         "String",
	"BOOLEAN":                        "Boolean",
	"NUMBER":                         "Number",
	"MULTIPLE SELECT FROM LIST":      "Multiple Select From List",
	"SINGLE SELECT FROM LIST":        "Single Select From List",
	"MULTIPLE SELECT FROM SQL QUERY": "Multiple Select From SQL Query",
	"SINGLE SELECT FROM SQL QUERY":   "Single Select From SQL Query",
	"PASSWORD":                       "Password",
	"LARGE TEXT":                     "Large Text",
	"CHECK BOX":                      "Check Box",
	"DATE":                           "Date",
	"ENUM":                           "Enum",
}

// TranslateValue translates a value using the provided map
func TranslateValue(input string, valueMap map[string]string) string {
	if input == "" {
		return ""
	}
	if val, ok := valueMap[input]; ok {
		return val
	}
	return input
}
