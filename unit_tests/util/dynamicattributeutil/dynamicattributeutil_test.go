// Copyright (c) 2025 Saviynt Inc.
// SPDX-License-Identifier: MPL-2.0

package dynamicattributeutil

import (
	"testing"

	"terraform-provider-Saviynt/util/dynamicattributeutil"
)

func TestTranslateValue(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		valueMap map[string]string
		expected string
	}{
		{
			name:     "Empty input returns empty string",
			input:    "",
			valueMap: map[string]string{"NUMBER": "NUMBER"},
			expected: "",
		},
		{
			name:     "Valid key returns mapped value",
			input:    "NUMBER",
			valueMap: map[string]string{"NUMBER": "NUMBER", "STRING": "STRING"},
			expected: "NUMBER",
		},
		{
			name:     "Invalid key returns original input",
			input:    "UNKNOWN_TYPE",
			valueMap: map[string]string{"NUMBER": "NUMBER", "STRING": "STRING"},
			expected: "UNKNOWN_TYPE",
		},
		{
			name:     "Empty map returns original input",
			input:    "test",
			valueMap: map[string]string{},
			expected: "test",
		},
		{
			name:     "Case sensitive key matching",
			input:    "number",
			valueMap: map[string]string{"NUMBER": "NUMBER", "number": "lowercase_number"},
			expected: "lowercase_number",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := dynamicattributeutil.TranslateValue(tt.input, tt.valueMap)
			if result != tt.expected {
				t.Errorf("TranslateValue(%q, %v) = %q, want %q", tt.input, tt.valueMap, result, tt.expected)
			}
		})
	}
}

func TestTranslateValueWithAttributeTypeMap(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "NUMBER maps correctly",
			input:    "NUMBER",
			expected: "NUMBER",
		},
		{
			name:     "STRING maps correctly",
			input:    "STRING",
			expected: "STRING",
		},
		{
			name:     "ENUM maps correctly",
			input:    "ENUM",
			expected: "ENUM",
		},
		{
			name:     "BOOLEAN maps correctly",
			input:    "BOOLEAN",
			expected: "BOOLEAN",
		},
		{
			name:     "MULTIPLE maps to full description",
			input:    "MULTIPLE",
			expected: "MULTIPLE SELECT FROM LIST",
		},
		{
			name:     "SQL MULTISELECT maps to full description",
			input:    "SQL MULTISELECT",
			expected: "MULTIPLE SELECT FROM SQL QUERY",
		},
		{
			name:     "SQL ENUM maps to full description",
			input:    "SQL ENUM",
			expected: "SINGLE SELECT FROM SQL QUERY",
		},
		{
			name:     "PASSWORD maps correctly",
			input:    "PASSWORD",
			expected: "PASSWORD",
		},
		{
			name:     "LARGE TEXT maps correctly",
			input:    "LARGE TEXT",
			expected: "LARGE TEXT",
		},
		{
			name:     "CHECKBOX maps to full description",
			input:    "CHECKBOX",
			expected: "CHECK BOX",
		},
		{
			name:     "DATE maps correctly",
			input:    "DATE",
			expected: "DATE",
		},
		{
			name:     "Unknown type returns original",
			input:    "UNKNOWN_TYPE",
			expected: "UNKNOWN_TYPE",
		},
		{
			name:     "Empty string returns empty",
			input:    "",
			expected: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := dynamicattributeutil.TranslateValue(tt.input, dynamicattributeutil.AttributeTypeMap)
			if result != tt.expected {
				t.Errorf("TranslateValue(%q, AttributeTypeMap) = %q, want %q", tt.input, result, tt.expected)
			}
		})
	}
}

func TestAttributeTypeMapCompleteness(t *testing.T) {
	expectedKeys := []string{
		"NUMBER", "STRING", "ENUM", "BOOLEAN", "MULTIPLE",
		"SQL MULTISELECT", "SQL ENUM", "PASSWORD", "LARGE TEXT",
		"CHECKBOX", "DATE",
	}

	for _, key := range expectedKeys {
		t.Run("AttributeTypeMap has key: "+key, func(t *testing.T) {
			if _, exists := dynamicattributeutil.AttributeTypeMap[key]; !exists {
				t.Errorf("AttributeTypeMap missing expected key: %q", key)
			}
		})
	}
}

func TestAttributeTypeMapValues(t *testing.T) {
	expectedMappings := map[string]string{
		"NUMBER":          "NUMBER",
		"STRING":          "STRING",
		"ENUM":            "ENUM",
		"BOOLEAN":         "BOOLEAN",
		"MULTIPLE":        "MULTIPLE SELECT FROM LIST",
		"SQL MULTISELECT": "MULTIPLE SELECT FROM SQL QUERY",
		"SQL ENUM":        "SINGLE SELECT FROM SQL QUERY",
		"PASSWORD":        "PASSWORD",
		"LARGE TEXT":      "LARGE TEXT",
		"CHECKBOX":        "CHECK BOX",
		"DATE":            "DATE",
	}

	for key, expectedValue := range expectedMappings {
		t.Run("AttributeTypeMap["+key+"] has correct value", func(t *testing.T) {
			if actualValue, exists := dynamicattributeutil.AttributeTypeMap[key]; !exists {
				t.Errorf("AttributeTypeMap missing key: %q", key)
			} else if actualValue != expectedValue {
				t.Errorf("AttributeTypeMap[%q] = %q, want %q", key, actualValue, expectedValue)
			}
		})
	}
}

func TestValidatorFactoryFunctions(t *testing.T) {
	t.Run("AttributeValueDisallowedForCertainAttributeTypes returns non-nil validator", func(t *testing.T) {
		validator := dynamicattributeutil.AttributeValueDisallowedForCertainAttributeTypes()
		if validator == nil {
			t.Error("AttributeValueDisallowedForCertainAttributeTypes() returned nil")
		}
	})

	t.Run("DescriptionDisallowedForCertainAttributeTypes returns non-nil validator", func(t *testing.T) {
		validator := dynamicattributeutil.DescriptionDisallowedForCertainAttributeTypes()
		if validator == nil {
			t.Error("DescriptionDisallowedForCertainAttributeTypes() returned nil")
		}
	})

	t.Run("DefaultValueDisallowedForCertainAttributeTypes returns non-nil validator", func(t *testing.T) {
		validator := dynamicattributeutil.DefaultValueDisallowedForCertainAttributeTypes()
		if validator == nil {
			t.Error("DefaultValueDisallowedForCertainAttributeTypes() returned nil")
		}
	})
}
