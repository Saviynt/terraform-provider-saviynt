// Copyright (c) 2025 Saviynt Inc.
// SPDX-License-Identifier: MPL-2.0

package endpointsutil

import (
	"testing"

	"terraform-provider-Saviynt/util/endpointsutil"
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
			valueMap: map[string]string{"1": "User"},
			expected: "",
		},
		{
			name:     "Valid key returns mapped value",
			input:    "1",
			valueMap: map[string]string{"1": "User", "2": "Usergroup"},
			expected: "User",
		},
		{
			name:     "Invalid key returns original input",
			input:    "999",
			valueMap: map[string]string{"1": "User", "2": "Usergroup"},
			expected: "999",
		},
		{
			name:     "Zero key returns mapped value",
			input:    "0",
			valueMap: map[string]string{"0": "None", "1": "User"},
			expected: "None",
		},
		{
			name:     "Empty map returns original input",
			input:    "test",
			valueMap: map[string]string{},
			expected: "test",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := endpointsutil.TranslateValue(tt.input, tt.valueMap)
			if result != tt.expected {
				t.Errorf("TranslateValue(%q, %v) = %q, want %q", tt.input, tt.valueMap, result, tt.expected)
			}
		})
	}
}

func TestTranslateValueWithPredefinedMaps(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		valueMap map[string]string
		expected string
	}{
		// endpointsutil.OwnerTypeMap tests
		{
			name:     "endpointsutil.OwnerTypeMap - empty string",
			input:    "",
			valueMap: endpointsutil.OwnerTypeMap,
			expected: "",
		},
		{
			name:     "endpointsutil.OwnerTypeMap - None",
			input:    "0",
			valueMap: endpointsutil.OwnerTypeMap,
			expected: "None",
		},
		{
			name:     "endpointsutil.OwnerTypeMap - User",
			input:    "1",
			valueMap: endpointsutil.OwnerTypeMap,
			expected: "User",
		},
		{
			name:     "endpointsutil.OwnerTypeMap - Usergroup",
			input:    "2",
			valueMap: endpointsutil.OwnerTypeMap,
			expected: "Usergroup",
		},
		// endpointsutil.RoleTypeMap tests
		{
			name:     "endpointsutil.RoleTypeMap - Node",
			input:    "0",
			valueMap: endpointsutil.RoleTypeMap,
			expected: "Node",
		},
		{
			name:     "endpointsutil.RoleTypeMap - Enabler",
			input:    "1",
			valueMap: endpointsutil.RoleTypeMap,
			expected: "Enabler",
		},
		{
			name:     "endpointsutil.RoleTypeMap - Emergency Access",
			input:    "3",
			valueMap: endpointsutil.RoleTypeMap,
			expected: "Emergency Access",
		},
		// endpointsutil.RequestOptionMap tests
		{
			name:     "endpointsutil.RequestOptionMap - DropdownSingle",
			input:    "1",
			valueMap: endpointsutil.RequestOptionMap,
			expected: "DropdownSingle",
		},
		{
			name:     "endpointsutil.RequestOptionMap - Table",
			input:    "2",
			valueMap: endpointsutil.RequestOptionMap,
			expected: "Table",
		},
		// ShowOnMap tests
		{
			name:     "ShowOnMap - All",
			input:    "-1",
			valueMap: endpointsutil.ShowOnMap,
			expected: "All",
		},
		{
			name:     "ShowOnMap - ShowOnServiceAccountRequest",
			input:    "0",
			valueMap: endpointsutil.ShowOnMap,
			expected: "ShowOnServiceAccountRequest",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := endpointsutil.TranslateValue(tt.input, tt.valueMap)
			if result != tt.expected {
				t.Errorf("TranslateValue(%q, %v) = %q, want %q", tt.input, tt.valueMap, result, tt.expected)
			}
		})
	}
}

func TestNormalizeToStringBool(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		// True cases
		{
			name:     "String '1' returns 'true'",
			input:    "1",
			expected: "true",
		},
		{
			name:     "String 'on' returns 'true'",
			input:    "on",
			expected: "true",
		},
		{
			name:     "String 'true' returns 'true'",
			input:    "true",
			expected: "true",
		},
		{
			name:     "String 'ON' (uppercase) returns 'true'",
			input:    "ON",
			expected: "true",
		},
		{
			name:     "String 'TRUE' (uppercase) returns 'true'",
			input:    "TRUE",
			expected: "true",
		},
		{
			name:     "String 'True' (mixed case) returns 'true'",
			input:    "True",
			expected: "true",
		},
		// False cases
		{
			name:     "String '0' returns 'false'",
			input:    "0",
			expected: "false",
		},
		{
			name:     "String 'off' returns 'false'",
			input:    "off",
			expected: "false",
		},
		{
			name:     "String 'false' returns 'false'",
			input:    "false",
			expected: "false",
		},
		{
			name:     "String 'OFF' (uppercase) returns 'false'",
			input:    "OFF",
			expected: "false",
		},
		{
			name:     "String 'FALSE' (uppercase) returns 'false'",
			input:    "FALSE",
			expected: "false",
		},
		{
			name:     "String 'False' (mixed case) returns 'false'",
			input:    "False",
			expected: "false",
		},
		// Default cases
		{
			name:     "Empty string returns 'false'",
			input:    "",
			expected: "false",
		},
		{
			name:     "Unknown string returns 'false'",
			input:    "unknown",
			expected: "false",
		},
		{
			name:     "Random string returns 'false'",
			input:    "random",
			expected: "false",
		},
		{
			name:     "Number string returns 'false'",
			input:    "123",
			expected: "false",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := endpointsutil.NormalizeToStringBool(tt.input)
			if result != tt.expected {
				t.Errorf("NormalizeToStringBool(%q) = %q, want %q", tt.input, result, tt.expected)
			}
		})
	}
}

func TestNormalizeJSON(t *testing.T) {
	tests := []struct {
		name        string
		input       string
		expected    string
		expectError bool
	}{
		{
			name:        "Valid JSON object",
			input:       `{"name": "test", "value": 123}`,
			expected:    `{"name":"test","value":123}`,
			expectError: false,
		},
		{
			name:        "Valid JSON array",
			input:       `[1, 2, 3]`,
			expected:    `[1,2,3]`,
			expectError: false,
		},
		{
			name:        "Valid JSON with whitespace",
			input:       `{ "key" : "value" , "number" : 42 }`,
			expected:    `{"key":"value","number":42}`,
			expectError: false,
		},
		{
			name:        "Valid JSON string",
			input:       `"simple string"`,
			expected:    `"simple string"`,
			expectError: false,
		},
		{
			name:        "Valid JSON number",
			input:       `42`,
			expected:    `42`,
			expectError: false,
		},
		{
			name:        "Valid JSON boolean",
			input:       `true`,
			expected:    `true`,
			expectError: false,
		},
		{
			name:        "Valid JSON null",
			input:       `null`,
			expected:    `null`,
			expectError: false,
		},
		{
			name:        "Nested JSON object",
			input:       `{"outer": {"inner": "value"}}`,
			expected:    `{"outer":{"inner":"value"}}`,
			expectError: false,
		},
		{
			name:        "Complex JSON with arrays and objects",
			input:       `{"users": [{"name": "John", "age": 30}, {"name": "Jane", "age": 25}]}`,
			expected:    `{"users":[{"name":"John","age":30},{"name":"Jane","age":25}]}`,
			expectError: false,
		},
		// Error cases
		{
			name:        "Invalid JSON - missing quote",
			input:       `{"name: "test"}`,
			expected:    "",
			expectError: true,
		},
		{
			name:        "Invalid JSON - trailing comma",
			input:       `{"name": "test",}`,
			expected:    "",
			expectError: true,
		},
		{
			name:        "Invalid JSON - malformed",
			input:       `{name: test}`,
			expected:    "",
			expectError: true,
		},
		{
			name:        "Empty string",
			input:       ``,
			expected:    "",
			expectError: true,
		},
		{
			name:        "Non-JSON string",
			input:       `this is not json`,
			expected:    "",
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := endpointsutil.NormalizeJSON(tt.input)

			if tt.expectError {
				if err == nil {
					t.Errorf("NormalizeJSON(%q) expected error but got none", tt.input)
				}
			} else {
				if err != nil {
					t.Errorf("NormalizeJSON(%q) unexpected error: %v", tt.input, err)
				}
				if result != tt.expected {
					t.Errorf("NormalizeJSON(%q) = %q, want %q", tt.input, result, tt.expected)
				}
			}
		})
	}
}

func TestRequiredMapValues(t *testing.T) {
	tests := []struct {
		name     string
		key      string
		expected bool
	}{
		{
			name:     "Empty string returns false",
			key:      "",
			expected: false,
		},
		{
			name:     "Key '0' returns false",
			key:      "0",
			expected: false,
		},
		{
			name:     "Key '1' returns true",
			key:      "1",
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, exists := endpointsutil.RequiredMap[tt.key]
			if !exists {
				t.Errorf("RequiredMap[%q] key does not exist", tt.key)
			}
			if result != tt.expected {
				t.Errorf("RequiredMap[%q] = %v, want %v", tt.key, result, tt.expected)
			}
		})
	}
}

func TestPredefinedMapsCompleteness(t *testing.T) {
	// Test that all predefined maps have expected keys
	t.Run("endpointsutil.OwnerTypeMap has expected keys", func(t *testing.T) {
		expectedKeys := []string{"", "0", "1", "2"}
		for _, key := range expectedKeys {
			if _, exists := endpointsutil.OwnerTypeMap[key]; !exists {
				t.Errorf("OwnerTypeMap missing expected key: %q", key)
			}
		}
	})

	t.Run("endpointsutil.RoleTypeMap has expected keys", func(t *testing.T) {
		expectedKeys := []string{"", "0", "1", "2", "3", "4", "5", "6"}
		for _, key := range expectedKeys {
			if _, exists := endpointsutil.RoleTypeMap[key]; !exists {
				t.Errorf("RoleTypeMap missing expected key: %q", key)
			}
		}
	})

	t.Run("endpointsutil.RequestOptionMap has expected keys", func(t *testing.T) {
		expectedKeys := []string{"0", "1", "2", "3"}
		for _, key := range expectedKeys {
			if _, exists := endpointsutil.RequestOptionMap[key]; !exists {
				t.Errorf("RequestOptionMap missing expected key: %q", key)
			}
		}
	})

	t.Run("RequiredMap has expected keys", func(t *testing.T) {
		expectedKeys := []string{"", "0", "1"}
		for _, key := range expectedKeys {
			if _, exists := endpointsutil.RequiredMap[key]; !exists {
				t.Errorf("RequiredMap missing expected key: %q", key)
			}
		}
	})

	t.Run("ShowOnMap has expected keys", func(t *testing.T) {
		expectedKeys := []string{"", "-1", "0", "1"}
		for _, key := range expectedKeys {
			if _, exists := endpointsutil.ShowOnMap[key]; !exists {
				t.Errorf("ShowOnMap missing expected key: %q", key)
			}
		}
	})
}
