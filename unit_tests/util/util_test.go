// Copyright (c) 2025 Saviynt Inc.
// SPDX-License-Identifier: MPL-2.0

package util

import (
	"context"
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"terraform-provider-Saviynt/util"
)

func TestBoolPointerOrEmpty(t *testing.T) {
	tests := []struct {
		name     string
		input    types.Bool
		expected *bool
	}{
		{
			name:     "true value",
			input:    types.BoolValue(true),
			expected: func() *bool { b := true; return &b }(),
		},
		{
			name:     "false value",
			input:    types.BoolValue(false),
			expected: func() *bool { b := false; return &b }(),
		},
		{
			name:     "null value",
			input:    types.BoolNull(),
			expected: nil,
		},
		{
			name:     "unknown value",
			input:    types.BoolUnknown(),
			expected: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := util.BoolPointerOrEmpty(tt.input)
			if tt.expected == nil {
				assert.Nil(t, result)
			} else {
				assert.NotNil(t, result)
				assert.Equal(t, *tt.expected, *result)
			}
		})
	}
}

func TestStringPtr(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected *string
	}{
		{
			name:     "non-empty string",
			input:    "test",
			expected: func() *string { s := "test"; return &s }(),
		},
		{
			name:     "empty string",
			input:    "",
			expected: func() *string { s := ""; return &s }(),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := util.StringPtr(tt.input)
			assert.NotNil(t, result)
			assert.Equal(t, *tt.expected, *result)
		})
	}
}

func TestSafeStringConnector(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected *string
	}{
		{
			name:     "valid string",
			input:    "test-connection",
			expected: func() *string { s := "test-connection"; return &s }(),
		},
		{
			name:     "empty string",
			input:    "",
			expected: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := util.SafeStringConnector(tt.input)
			if tt.expected == nil {
				assert.Nil(t, result)
			} else {
				assert.NotNil(t, result)
				assert.Equal(t, *tt.expected, *result)
			}
		})
	}
}

func TestStringPointerOrEmpty(t *testing.T) {
	tests := []struct {
		name     string
		input    types.String
		expected *string
	}{
		{
			name:     "valid string",
			input:    types.StringValue("test"),
			expected: func() *string { s := "test"; return &s }(),
		},
		{
			name:     "null string",
			input:    types.StringNull(),
			expected: nil,
		},
		{
			name:     "unknown string",
			input:    types.StringUnknown(),
			expected: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := util.StringPointerOrEmpty(tt.input)
			if tt.expected == nil {
				assert.Nil(t, result)
			} else {
				assert.NotNil(t, result)
				assert.Equal(t, *tt.expected, *result)
			}
		})
	}
}

func TestConvertTypesStringToStrings(t *testing.T) {
	tests := []struct {
		name     string
		input    []string
		expected []types.String
	}{
		{
			name:  "valid strings",
			input: []string{"test1", "test2", "test3"},
			expected: []types.String{
				types.StringValue("test1"),
				types.StringValue("test2"),
				types.StringValue("test3"),
			},
		},
		{
			name:     "empty slice",
			input:    []string{},
			expected: []types.String{},
		},
		{
			name:  "strings with empty values",
			input: []string{"test1", "", "test3"},
			expected: []types.String{
				types.StringValue("test1"),
				types.StringValue("test3"),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := util.ConvertTypesStringToStrings(tt.input)
			assert.Equal(t, len(tt.expected), len(result))
			for i, expected := range tt.expected {
				if i < len(result) {
					assert.Equal(t, expected.ValueString(), result[i].ValueString())
				}
			}
		})
	}
}

func TestInt32PtrToTFString(t *testing.T) {
	tests := []struct {
		name     string
		input    *int32
		expected types.String
	}{
		{
			name:     "valid int32 pointer",
			input:    func() *int32 { i := int32(42); return &i }(),
			expected: types.StringValue("42"),
		},
		{
			name:     "nil pointer",
			input:    nil,
			expected: types.StringNull(),
		},
		{
			name:     "zero value",
			input:    func() *int32 { i := int32(0); return &i }(),
			expected: types.StringValue("0"),
		},
		{
			name:     "negative value",
			input:    func() *int32 { i := int32(-42); return &i }(),
			expected: types.StringValue("-42"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := util.Int32PtrToTFString(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestSafeString(t *testing.T) {
	tests := []struct {
		name     string
		input    *string
		expected types.String
	}{
		{
			name:     "valid string pointer",
			input:    func() *string { s := "test-value"; return &s }(),
			expected: types.StringValue("test-value"),
		},
		{
			name:     "nil pointer",
			input:    nil,
			expected: types.StringValue(""),
		},
		{
			name:     "empty string pointer",
			input:    func() *string { s := ""; return &s }(),
			expected: types.StringValue(""),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := util.SafeString(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestSafeBoolDatasource(t *testing.T) {
	tests := []struct {
		name     string
		input    *bool
		expected types.Bool
	}{
		{
			name:     "valid true bool pointer",
			input:    func() *bool { b := true; return &b }(),
			expected: types.BoolValue(true),
		},
		{
			name:     "valid false bool pointer",
			input:    func() *bool { b := false; return &b }(),
			expected: types.BoolValue(false),
		},
		{
			name:     "nil pointer",
			input:    nil,
			expected: types.BoolNull(),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := util.SafeBoolDatasource(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestSafeStringDatasource(t *testing.T) {
	tests := []struct {
		name     string
		input    *string
		expected types.String
	}{
		{
			name:     "valid string pointer",
			input:    func() *string { s := "test-datasource"; return &s }(),
			expected: types.StringValue("test-datasource"),
		},
		{
			name:     "empty string pointer",
			input:    func() *string { s := ""; return &s }(),
			expected: types.StringValue(""),
		},
		{
			name:     "nil pointer",
			input:    nil,
			expected: types.StringNull(),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := util.SafeStringDatasource(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestSafeDeref(t *testing.T) {
	tests := []struct {
		name     string
		input    *string
		expected string
	}{
		{
			name:     "valid string pointer",
			input:    func() *string { s := "dereferenced-value"; return &s }(),
			expected: "dereferenced-value",
		},
		{
			name:     "empty string pointer",
			input:    func() *string { s := ""; return &s }(),
			expected: "",
		},
		{
			name:     "nil pointer",
			input:    nil,
			expected: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := util.SafeDeref(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestSafeStringValue(t *testing.T) {
	tests := []struct {
		name     string
		input    types.String
		expected string
	}{
		{
			name:     "valid string value",
			input:    types.StringValue("test-string-value"),
			expected: "test-string-value",
		},
		{
			name:     "empty string value",
			input:    types.StringValue(""),
			expected: "",
		},
		{
			name:     "null string",
			input:    types.StringNull(),
			expected: "",
		},
		{
			name:     "unknown string",
			input:    types.StringUnknown(),
			expected: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := util.SafeStringValue(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestSafeList(t *testing.T) {
	tests := []struct {
		name        string
		input       []string
		expectError bool
	}{
		{
			name:        "valid string slice",
			input:       []string{"item1", "item2", "item3"},
			expectError: false,
		},
		{
			name:        "empty string slice",
			input:       []string{},
			expectError: false,
		},
		{
			name:        "single item",
			input:       []string{"single-item"},
			expectError: false,
		},
		{
			name:        "slice with empty strings",
			input:       []string{"item1", "", "item3"},
			expectError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, diags := util.SafeList(tt.input)

			if tt.expectError {
				assert.True(t, diags.HasError())
			} else {
				assert.False(t, diags.HasError())
				assert.NotNil(t, result)

				// Verify the list contains the expected number of elements
				elements := result.Elements()
				assert.Equal(t, len(tt.input), len(elements))
			}
		})
	}
}

func TestStringsToTypeStrings(t *testing.T) {
	tests := []struct {
		name     string
		input    []string
		expected []types.String
	}{
		{
			name:  "valid strings",
			input: []string{"str1", "str2", "str3"},
			expected: []types.String{
				types.StringValue("str1"),
				types.StringValue("str2"),
				types.StringValue("str3"),
			},
		},
		{
			name:     "empty slice",
			input:    []string{},
			expected: []types.String{},
		},
		{
			name:  "strings with empty values",
			input: []string{"str1", "", "str3"},
			expected: []types.String{
				types.StringValue("str1"),
				types.StringValue(""),
				types.StringValue("str3"),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := util.StringsToTypeStrings(tt.input)
			assert.Equal(t, len(tt.expected), len(result))

			for i, expected := range tt.expected {
				if i < len(result) {
					assert.Equal(t, expected.ValueString(), result[i].ValueString())
				}
			}
		})
	}
}

func TestStringsToSet(t *testing.T) {
	tests := []struct {
		name        string
		input       []string
		expectNull  bool
		expectError bool
	}{
		{
			name:        "valid strings",
			input:       []string{"item1", "item2", "item3"},
			expectNull:  false,
			expectError: false,
		},
		{
			name:        "empty slice",
			input:       []string{},
			expectNull:  true, // ← FIXED: Empty slice returns null set by design
			expectError: false,
		},
		{
			name:        "single item",
			input:       []string{"single"},
			expectNull:  false,
			expectError: false,
		},
		{
			name:        "duplicate items",
			input:       []string{"item1", "item2", "item1"},
			expectNull:  false,
			expectError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := util.StringsToSet(tt.input)

			if tt.expectNull {
				assert.True(t, result.IsNull())
			} else {
				assert.False(t, result.IsNull())
				assert.False(t, result.IsUnknown())

				// Additional validation for non-null sets
				if !result.IsNull() && !result.IsUnknown() {
					elements := result.Elements()
					assert.GreaterOrEqual(t, len(elements), 0)
				}
			}
		})
	}
}

func TestConvertStringsToTypesString(t *testing.T) {
	tests := []struct {
		name     string
		input    []string
		expected []types.String
	}{
		{
			name:  "valid strings",
			input: []string{"convert1", "convert2", "convert3"},
			expected: []types.String{
				types.StringValue("convert1"),
				types.StringValue("convert2"),
				types.StringValue("convert3"),
			},
		},
		{
			name:     "empty slice",
			input:    []string{},
			expected: []types.String{},
		},
		{
			name:  "mixed content",
			input: []string{"valid", "", "another"},
			expected: []types.String{
				types.StringValue("valid"),
				types.StringValue(""),
				types.StringValue("another"),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := util.ConvertStringsToTypesString(tt.input)
			assert.Equal(t, len(tt.expected), len(result))

			for i, expected := range tt.expected {
				if i < len(result) {
					assert.Equal(t, expected.ValueString(), result[i].ValueString())
				}
			}
		})
	}
}

func TestMarshalDeterministic(t *testing.T) {
	tests := []struct {
		name        string
		input       map[string]string
		expectError bool
	}{
		{
			name:        "simple map",
			input:       map[string]string{"key1": "value1", "key2": "value2"},
			expectError: false,
		},
		{
			name:        "empty map",
			input:       map[string]string{},
			expectError: false,
		},
		{
			name:        "single key map",
			input:       map[string]string{"single": "value"},
			expectError: false,
		},
		{
			name:        "nil map",
			input:       nil,
			expectError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := util.MarshalDeterministic(tt.input)

			if tt.expectError {
				assert.Error(t, err)
				assert.Empty(t, result)
			} else {
				assert.NoError(t, err)
				assert.NotEmpty(t, result)

				// Test deterministic behavior - same input should produce same output
				result2, err2 := util.MarshalDeterministic(tt.input)
				assert.NoError(t, err2)
				assert.Equal(t, result, result2)
			}
		})
	}
}

func TestSanitizeTypesStringList(t *testing.T) {
	tests := []struct {
		name     string
		input    []types.String
		expected []types.String
	}{
		{
			name: "valid strings",
			input: []types.String{
				types.StringValue("valid1"),
				types.StringValue("valid2"),
				types.StringValue("valid3"),
			},
			expected: []types.String{
				types.StringValue("valid1"),
				types.StringValue("valid2"),
				types.StringValue("valid3"),
			},
		},
		{
			name:     "empty slice",
			input:    []types.String{},
			expected: []types.String{},
		},
		{
			name: "mixed valid and null strings",
			input: []types.String{
				types.StringValue("valid"),
				types.StringNull(),
				types.StringValue("another"),
			},
			expected: []types.String{
				types.StringValue("valid"),
				types.StringValue("another"),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := util.SanitizeTypesStringList(tt.input)
			assert.Equal(t, len(tt.expected), len(result))

			for i, expected := range tt.expected {
				if i < len(result) {
					assert.Equal(t, expected.ValueString(), result[i].ValueString())
				}
			}
		})
	}
}

func TestStringsFromSet(t *testing.T) {
	tests := []struct {
		name      string
		input     types.Set
		expected  []string
		expectNil bool
	}{
		{
			name: "valid set",
			input: types.SetValueMust(types.StringType, []attr.Value{
				types.StringValue("item1"),
				types.StringValue("item2"),
				types.StringValue("item3"),
			}),
			expected:  []string{"item1", "item2", "item3"},
			expectNil: false,
		},
		{
			name:      "empty set",
			input:     types.SetValueMust(types.StringType, []attr.Value{}),
			expected:  nil, // ← FIXED: Empty set returns nil
			expectNil: true,
		},
		{
			name:      "null set",
			input:     types.SetNull(types.StringType),
			expected:  nil, // ← FIXED: Null set returns nil
			expectNil: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := util.StringsFromSet(tt.input)

			if tt.expectNil {
				assert.Nil(t, result)
			} else {
				assert.NotNil(t, result)
				assert.Equal(t, len(tt.expected), len(result))

				// Convert to maps for easier comparison (sets are unordered)
				expectedMap := make(map[string]bool)
				for _, s := range tt.expected {
					expectedMap[s] = true
				}

				resultMap := make(map[string]bool)
				for _, s := range result {
					resultMap[s] = true
				}

				assert.Equal(t, expectedMap, resultMap)
			}
		})
	}
}

func TestStringsFromList(t *testing.T) {
	tests := []struct {
		name      string
		input     types.List
		expected  []string
		expectNil bool
	}{
		{
			name: "valid list",
			input: types.ListValueMust(types.StringType, []attr.Value{
				types.StringValue("item1"),
				types.StringValue("item2"),
				types.StringValue("item3"),
			}),
			expected:  []string{"item1", "item2", "item3"},
			expectNil: false,
		},
		{
			name:      "empty list",
			input:     types.ListValueMust(types.StringType, []attr.Value{}),
			expected:  nil, // ← FIXED: Empty list returns nil
			expectNil: true,
		},
		{
			name:      "null list",
			input:     types.ListNull(types.StringType),
			expected:  nil, // ← FIXED: Null list returns nil
			expectNil: true,
		},
		{
			name:      "unknown list",
			input:     types.ListUnknown(types.StringType),
			expected:  nil, // ← FIXED: Unknown list returns nil
			expectNil: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := util.StringsFromList(tt.input)

			if tt.expectNil {
				assert.Nil(t, result)
			} else {
				assert.NotNil(t, result)
				assert.Equal(t, tt.expected, result)
			}
		})
	}
}

func TestNormalizeTFSetString(t *testing.T) {
	tests := []struct {
		name     string
		input    types.Set
		expected types.Set
	}{
		{
			name: "valid set",
			input: types.SetValueMust(types.StringType, []attr.Value{
				types.StringValue("item1"),
				types.StringValue("item2"),
			}),
			expected: types.SetValueMust(types.StringType, []attr.Value{
				types.StringValue("item1"),
				types.StringValue("item2"),
			}),
		},
		{
			name:     "empty set",
			input:    types.SetValueMust(types.StringType, []attr.Value{}),
			expected: types.SetNull(types.StringType), // ← FIXED: Empty set normalizes to null set
		},
		{
			name:     "null set",
			input:    types.SetNull(types.StringType),
			expected: types.SetNull(types.StringType),
		},
		{
			name:     "unknown set",
			input:    types.SetUnknown(types.StringType),
			expected: types.SetNull(types.StringType), // ← FIXED: Unknown set normalizes to null set
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := util.NormalizeTFSetString(tt.input)
			assert.Equal(t, tt.expected.IsNull(), result.IsNull())
			assert.Equal(t, tt.expected.IsUnknown(), result.IsUnknown())

			if !tt.expected.IsNull() && !tt.expected.IsUnknown() {
				expectedElements := tt.expected.Elements()
				resultElements := result.Elements()
				assert.Equal(t, len(expectedElements), len(resultElements))
			}
		})
	}
}

func TestSafeInt32(t *testing.T) {
	tests := []struct {
		name     string
		input    *int32
		expected types.Int32
	}{
		{
			name:     "valid int32 pointer",
			input:    func() *int32 { i := int32(42); return &i }(),
			expected: types.Int32Value(42), // ← FIXED: Expect types.Int32Value
		},
		{
			name:     "zero value",
			input:    func() *int32 { i := int32(0); return &i }(),
			expected: types.Int32Value(0), // ← FIXED: Expect types.Int32Value
		},
		{
			name:     "negative value",
			input:    func() *int32 { i := int32(-42); return &i }(),
			expected: types.Int32Value(-42), // ← FIXED: Expect types.Int32Value
		},
		{
			name:     "nil pointer",
			input:    nil,
			expected: types.Int32Null(), // ← FIXED: Expect types.Int32Null
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := util.SafeInt32(tt.input)
			assert.Equal(t, tt.expected, result)

			// Additional validation for non-null values
			if !tt.expected.IsNull() {
				assert.False(t, result.IsNull())
				assert.False(t, result.IsUnknown())
				assert.Equal(t, tt.expected.ValueInt32(), result.ValueInt32())
			} else {
				assert.True(t, result.IsNull())
			}
		})
	}
}

func TestSafeInt64(t *testing.T) {
	tests := []struct {
		name     string
		input    *int64
		expected types.Int64
	}{
		{
			name:     "valid int64 pointer",
			input:    func() *int64 { i := int64(42); return &i }(),
			expected: types.Int64Value(42), // ← FIXED: Expect types.Int64Value
		},
		{
			name:     "zero value",
			input:    func() *int64 { i := int64(0); return &i }(),
			expected: types.Int64Value(0), // ← FIXED: Expect types.Int64Value
		},
		{
			name:     "negative value",
			input:    func() *int64 { i := int64(-42); return &i }(),
			expected: types.Int64Value(-42), // ← FIXED: Expect types.Int64Value
		},
		{
			name:     "large value",
			input:    func() *int64 { i := int64(9223372036854775807); return &i }(),
			expected: types.Int64Value(9223372036854775807), // ← FIXED: Expect types.Int64Value
		},
		{
			name:     "nil pointer",
			input:    nil,
			expected: types.Int64Null(), // ← FIXED: Expect types.Int64Null
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := util.SafeInt64(tt.input)
			assert.Equal(t, tt.expected, result)

			// Additional validation for non-null values
			if !tt.expected.IsNull() {
				assert.False(t, result.IsNull())
				assert.False(t, result.IsUnknown())
				assert.Equal(t, tt.expected.ValueInt64(), result.ValueInt64())
			} else {
				assert.True(t, result.IsNull())
			}
		})
	}
}

func TestTypesStringOrOriginal(t *testing.T) {
	tests := []struct {
		name     string
		input    types.String
		original *string
		expected types.String
	}{
		{
			name:     "valid types.String",
			input:    types.StringValue("new-value"),
			original: func() *string { s := "original-value"; return &s }(),
			expected: types.StringValue("original-value"), // Function returns original when apiValue is provided
		},
		{
			name:     "null types.String with original",
			input:    types.StringNull(),
			original: func() *string { s := "original-value"; return &s }(),
			expected: types.StringValue("original-value"),
		},
		{
			name:     "unknown types.String with original",
			input:    types.StringUnknown(),
			original: func() *string { s := "original-value"; return &s }(),
			expected: types.StringValue("original-value"),
		},
		{
			name:     "null types.String with nil original",
			input:    types.StringNull(),
			original: nil,
			expected: types.StringNull(),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := util.TypesStringOrOriginal(tt.input, tt.original)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestSafeStringAlt(t *testing.T) {
	tests := []struct {
		name         string
		input        *string
		defaultValue string
		expected     types.String
	}{
		{
			name:         "valid string pointer",
			input:        func() *string { s := "test-value"; return &s }(),
			defaultValue: "default",
			expected:     types.StringValue("test-value"),
		},
		{
			name:         "nil pointer with default",
			input:        nil,
			defaultValue: "default-value",
			expected:     types.StringValue("default-value"),
		},
		{
			name:         "empty string pointer",
			input:        func() *string { s := ""; return &s }(),
			defaultValue: "default",
			expected:     types.StringValue("default"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := util.SafeStringAlt(tt.input, tt.defaultValue)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestPreserveString(t *testing.T) {
	tests := []struct {
		name     string
		apiVal   *string
		oldVal   types.String
		expected types.String
	}{
		{
			name:     "valid api value",
			apiVal:   func() *string { s := "new-api-value"; return &s }(),
			oldVal:   types.StringValue("old-value"),
			expected: types.StringValue("new-api-value"),
		},
		{
			name:     "nil api value with old value",
			apiVal:   nil,
			oldVal:   types.StringValue("old-value"),
			expected: types.StringValue("old-value"),
		},
		{
			name:     "empty api value with old value",
			apiVal:   func() *string { s := ""; return &s }(),
			oldVal:   types.StringValue("old-value"),
			expected: types.StringValue("old-value"),
		},
		{
			name:     "whitespace api value with old value",
			apiVal:   func() *string { s := "   "; return &s }(),
			oldVal:   types.StringValue("old-value"),
			expected: types.StringValue("old-value"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := util.PreserveString(tt.apiVal, tt.oldVal)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestSafeStringPreserveNull(t *testing.T) {
	tests := []struct {
		name     string
		input    *string
		expected types.String
	}{
		{
			name:     "valid string pointer",
			input:    func() *string { s := "test-value"; return &s }(),
			expected: types.StringValue("test-value"),
		},
		{
			name:     "empty string pointer",
			input:    func() *string { s := ""; return &s }(),
			expected: types.StringValue(""),
		},
		{
			name:     "nil pointer",
			input:    nil,
			expected: types.StringNull(),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := util.SafeStringPreserveNull(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestConvertListToStringSlice(t *testing.T) {
	tests := []struct {
		name     string
		input    types.List
		expected []string
	}{
		{
			name: "valid list",
			input: types.ListValueMust(types.StringType, []attr.Value{
				types.StringValue("item1"),
				types.StringValue("item2"),
				types.StringValue("item3"),
			}),
			expected: []string{"item1", "item2", "item3"},
		},
		{
			name:     "empty list",
			input:    types.ListValueMust(types.StringType, []attr.Value{}),
			expected: []string{},
		},
		{
			name:     "null list",
			input:    types.ListNull(types.StringType),
			expected: nil,
		},
		{
			name:     "unknown list",
			input:    types.ListUnknown(types.StringType),
			expected: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := util.ConvertListToStringSlice(context.Background(), tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestLoadConnectorDataForEphemeral(t *testing.T) {
	// Create a temporary directory for test files
	tempDir, err := os.MkdirTemp("", "connector_test")
	require.NoError(t, err)
	defer os.RemoveAll(tempDir)

	t.Run("Empty filepath", func(t *testing.T) {
		result, err := util.LoadConnectorDataForEphemeral("")
		assert.Nil(t, result)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "filepath cannot be empty")
	})

	t.Run("Whitespace only filepath", func(t *testing.T) {
		result, err := util.LoadConnectorDataForEphemeral("   \t\n   ")
		assert.Nil(t, result)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "filepath cannot be empty")
	})

	t.Run("Non-existent file", func(t *testing.T) {
		nonExistentPath := filepath.Join(tempDir, "nonexistent.json")
		result, err := util.LoadConnectorDataForEphemeral(nonExistentPath)
		assert.Nil(t, result)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "file does not exist")
		assert.Contains(t, err.Error(), nonExistentPath)
	})

	t.Run("Valid JSON with string values", func(t *testing.T) {
		testData := map[string]interface{}{
			"USERNAME": "testuser",
			"PASSWORD": "testpass",
			"HOST":     "localhost",
		}

		testFile := filepath.Join(tempDir, "valid_strings.json")
		createTestFile(t, testFile, testData)

		result, err := util.LoadConnectorDataForEphemeral(testFile)
		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, "testuser", result["USERNAME"])
		assert.Equal(t, "testpass", result["PASSWORD"])
		assert.Equal(t, "localhost", result["HOST"])
	})

	t.Run("Valid JSON with mixed data types", func(t *testing.T) {
		testData := map[string]interface{}{
			"USERNAME": "testuser",
			"PORT":     8080,
			"ENABLED":  true,
			"CONFIG": map[string]interface{}{
				"timeout": 30,
				"retry":   true,
			},
			"SERVERS": []interface{}{"server1", "server2", "server3"},
		}

		testFile := filepath.Join(tempDir, "mixed_types.json")
		createTestFile(t, testFile, testData)

		result, err := util.LoadConnectorDataForEphemeral(testFile)
		assert.NoError(t, err)
		assert.NotNil(t, result)

		// String value
		assert.Equal(t, "testuser", result["USERNAME"])

		// Number should be marshaled to JSON
		assert.Equal(t, "8080", result["PORT"])

		// Boolean should be marshaled to JSON
		assert.Equal(t, "true", result["ENABLED"])

		// Map should be marshaled to JSON
		var configMap map[string]interface{}
		err = json.Unmarshal([]byte(result["CONFIG"]), &configMap)
		assert.NoError(t, err)
		assert.Equal(t, float64(30), configMap["timeout"]) // JSON numbers are float64
		assert.Equal(t, true, configMap["retry"])

		// Array should be marshaled to JSON
		var serversList []interface{}
		err = json.Unmarshal([]byte(result["SERVERS"]), &serversList)
		assert.NoError(t, err)
		assert.Len(t, serversList, 3)
		assert.Equal(t, "server1", serversList[0])
		assert.Equal(t, "server2", serversList[1])
		assert.Equal(t, "server3", serversList[2])
	})

	t.Run("Invalid JSON file", func(t *testing.T) {
		invalidJSONFile := filepath.Join(tempDir, "invalid.json")
		err := os.WriteFile(invalidJSONFile, []byte(`{"invalid": json content}`), 0644)
		require.NoError(t, err)

		result, err := util.LoadConnectorDataForEphemeral(invalidJSONFile)
		assert.Nil(t, result)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "failed to unmarshal connector config")
	})

	t.Run("Empty JSON file", func(t *testing.T) {
		emptyJSONFile := filepath.Join(tempDir, "empty.json")
		err := os.WriteFile(emptyJSONFile, []byte(`{}`), 0644)
		require.NoError(t, err)

		result, err := util.LoadConnectorDataForEphemeral(emptyJSONFile)
		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Empty(t, result)
	})

	t.Run("Complex nested JSON", func(t *testing.T) {
		testData := map[string]interface{}{
			"CONNECTION_JSON": map[string]interface{}{
				"auth": map[string]interface{}{
					"type":     "oauth2",
					"clientId": "test-client",
					"scopes":   []string{"read", "write"},
				},
				"endpoints": map[string]interface{}{
					"base":  "https://api.example.com",
					"users": "/v1/users",
				},
			},
			"ACCESS_TOKENS": []interface{}{
				map[string]interface{}{
					"token": "token1",
					"type":  "bearer",
				},
				map[string]interface{}{
					"token": "token2",
					"type":  "api_key",
				},
			},
		}

		testFile := filepath.Join(tempDir, "complex.json")
		createTestFile(t, testFile, testData)

		result, err := util.LoadConnectorDataForEphemeral(testFile)
		assert.NoError(t, err)
		assert.NotNil(t, result)

		// Verify CONNECTION_JSON is properly marshaled
		var connectionJSON map[string]interface{}
		err = json.Unmarshal([]byte(result["CONNECTION_JSON"]), &connectionJSON)
		assert.NoError(t, err)
		assert.Contains(t, connectionJSON, "auth")
		assert.Contains(t, connectionJSON, "endpoints")

		// Verify ACCESS_TOKENS is properly marshaled
		var accessTokens []interface{}
		err = json.Unmarshal([]byte(result["ACCESS_TOKENS"]), &accessTokens)
		assert.NoError(t, err)
		assert.Len(t, accessTokens, 2)
	})

	t.Run("File with no read permissions", func(t *testing.T) {
		// Skip this test on Windows as file permissions work differently
		if os.Getenv("GOOS") == "windows" {
			t.Skip("Skipping permission test on Windows")
		}

		// Skip this test if running as root (common in Docker containers)
		if os.Geteuid() == 0 {
			t.Skip("Skipping permission test when running as root (Docker/CI environment)")
		}

		testData := map[string]interface{}{"test": "value"}
		noReadFile := filepath.Join(tempDir, "no_read.json")
		createTestFile(t, noReadFile, testData)

		// Remove read permissions
		err := os.Chmod(noReadFile, 0000)
		require.NoError(t, err)
		defer os.Chmod(noReadFile, 0644) // Restore permissions for cleanup

		result, err := util.LoadConnectorDataForEphemeral(noReadFile)
		assert.Nil(t, result)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "failed to read connector config file")
	})
}

// Helper function to create test JSON files
func createTestFile(t *testing.T, filePath string, data map[string]interface{}) {
	jsonData, err := json.MarshalIndent(data, "", "  ")
	require.NoError(t, err)
	err = os.WriteFile(filePath, jsonData, 0644)
	require.NoError(t, err)
}
