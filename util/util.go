// Copyright (c) 2025 Saviynt Inc.
// SPDX-License-Identifier: MPL-2.0

package util

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// safeString converts a *string to types.String safely.
func SafeString(s *string) types.String {
	if s == nil {
		return types.StringValue("")
	}
	return types.StringValue(*s)
}

func SafeBoolDatasource(b *bool) types.Bool {
	if b == nil {
		return types.BoolNull()
	}
	return types.BoolValue(*b)
}

func SafeStringDatasource(s *string) types.String {
	if s == nil {
		return types.StringNull()
	}
	return types.StringValue(*s)
}

// SafeDeref safely dereferences a *string, returning an empty string if nil.
func SafeDeref(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}

func SafeStringValue(s types.String) string {
	if s.IsNull() || s.IsUnknown() {
		return ""
	}
	return s.ValueString()
}

// safeList converts a []string to a Terraform types.List.
func SafeList(items []string) (types.List, diag.Diagnostics) {
	if len(items) == 0 {
		return types.ListValueMust(types.StringType, []attr.Value{}), nil
	}

	var values []attr.Value
	for _, item := range items {
		values = append(values, types.StringValue(item))
	}

	return types.ListValue(types.StringType, values)
}

// StringsToTypeStrings converts a slice of Go strings to a slice of types.String.
func StringsToTypeStrings(items []string) []types.String {
	var result []types.String
	for _, s := range items {
		result = append(result, types.StringValue(s))
	}
	return result
}

func StringsToSet(items []string) types.Set {
	var elements []attr.Value
	for _, item := range items {
		elements = append(elements, types.StringValue(item))
	}

	if len(elements) == 0 {
		return types.SetNull(types.StringType)
	}

	return types.SetValueMust(types.StringType, elements)
}

// ConvertStringsToTypesString converts a slice of Go strings to a slice of types.String.
func ConvertStringsToTypesString(items []string) []types.String {
	var result []types.String
	for _, item := range items {
		result = append(result, types.StringValue(item))
	}
	if len(result) == 0 {
		return nil
	}
	return result
}

// marshalDeterministic marshals a map[string]string into a JSON string
// with keys sorted in lexicographical order.
func MarshalDeterministic(m map[string]string) (string, error) {
	// Get the keys and sort them.
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	// Build an ordered map slice.
	ordered := make(map[string]string, len(m))
	for _, k := range keys {
		ordered[k] = m[k]
	}

	// Marshal the ordered map.
	b, err := json.Marshal(ordered)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func BoolPointerOrEmpty(tfBool types.Bool) *bool {
	if tfBool.IsNull() || tfBool.IsUnknown() {
		return nil
	}
	val := tfBool.ValueBool()
	return &val
}

func StringPtr(v string) *string {
	return &v
}

func SafeStringConnector(s string) *string {
	if s == "" {
		return nil
	}
	return &s
}

func StringPointerOrEmpty(tfStr types.String) *string {
	if tfStr.IsNull() || tfStr.IsUnknown() {
		return nil
	}
	val := tfStr.ValueString()
	return &val
}

func ConvertTypesStringToStrings(input []string) []types.String {
	var result []types.String
	for _, v := range input {
		if v != "" {
			result = append(result, types.StringValue(v))
		}
	}
	if len(result) == 0 {
		return nil
	}
	return result
}

func SanitizeTypesStringList(input []types.String) []types.String {
	var result []types.String
	for _, v := range input {
		if !v.IsNull() && !v.IsUnknown() && v.ValueString() != "" {
			result = append(result, v)
		}
	}
	if len(result) == 0 {
		return nil
	}
	return result
}

func StringsFromSet(input types.Set) []string {
	if input.IsNull() || input.IsUnknown() {
		return nil
	}

	var result []string

	for _, val := range input.Elements() {
		strVal, ok := val.(types.String)
		if !ok || strVal.IsNull() || strVal.IsUnknown() {
			continue
		}
		result = append(result, strVal.ValueString())
	}

	if len(result) == 0 {
		return nil
	}

	return result
}

func StringsFromList(input types.List) []string {
	if input.IsNull() || input.IsUnknown() {
		return nil
	}

	var result []string

	for _, val := range input.Elements() {
		strVal, ok := val.(types.String)
		if !ok || strVal.IsNull() || strVal.IsUnknown() {
			continue
		}
		result = append(result, strVal.ValueString())
	}

	if len(result) == 0 {
		return nil
	}

	return result
}

func NormalizeTFSetString(list types.Set) types.Set {
	if list.IsNull() || list.IsUnknown() || len(list.Elements()) == 0 {
		return types.SetNull(types.StringType)
	}
	return list
}

func SafeInt32(ptr *int32) types.Int32 {
	if ptr == nil {
		return types.Int32Null()
	}
	return types.Int32Value(*ptr)
}

func SafeInt64[T int32 | int64](value *T) types.Int64 {
	if value == nil {
		return types.Int64Null()
	}
	return types.Int64Value(int64(*value))
}

func Int32PtrToTFString(val *int32) types.String {
	if val != nil {
		str := strconv.Itoa(int(*val))
		return types.StringValue(str)
	}
	return types.StringNull()
}

func LoadConnectorDataForEphemeral(filePath string) (map[string]string, error) {
	// Validate input
	if strings.TrimSpace(filePath) == "" {
		return nil, fmt.Errorf("filepath cannot be empty")
	}
	// Check if file exists
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return nil, fmt.Errorf("file does not exist: %s", filePath)
	}
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read connector config file: %v", err)
	}

	var allData map[string]interface{}
	if err := json.Unmarshal(data, &allData); err != nil {
		return nil, fmt.Errorf("failed to unmarshal connector config: %v", err)
	}

	result := make(map[string]string)
	for key, value := range allData {
		switch v := value.(type) {
		case string:
			result[key] = v

		case map[string]interface{}, []interface{}:
			jsonValue, err := json.Marshal(v)
			if err != nil {
				return nil, fmt.Errorf("failed to marshal key %s: %v", key, err)
			}
			result[key] = string(jsonValue)

		default:
			// Marshal any other types safely (e.g., numbers, bool)
			jsonValue, err := json.Marshal(v)
			if err != nil {
				return nil, fmt.Errorf("failed to marshal key %s: %v", key, err)
			}
			result[key] = string(jsonValue)
		}
	}

	return result, nil
}

func TypesStringOrOriginal(original types.String, apiValue *string) types.String {
	if apiValue != nil && *apiValue != "" {
		return types.StringValue(*apiValue)
	}
	return original
}

func SafeStringAlt(s *string, replace string) types.String {
	if s == nil || *s == "" {
		return types.StringValue(replace)
	}

	return types.StringValue(*s)
}

func PreserveString(apiVal *string, oldVal types.String) types.String {
	if apiVal != nil && strings.TrimSpace(*apiVal) != "" {
		// Return the value from the API only if it's non-empty
		return types.StringValue(*apiVal)
	}
	if !oldVal.IsNull() && !oldVal.IsUnknown() {
		// Preserve value from state if it's valid
		return oldVal
	}
	// Fallback to null if API and old value are empty
	return types.StringNull()
}

func SafeStringPreserveNull(s *string) types.String {
	if s == nil {
		return types.StringNull()
	}
	return types.StringValue(*s)
}

func ConvertListToStringSlice(ctx context.Context, tfList types.List) []string {
	if tfList.IsNull() || tfList.IsUnknown() {
		return nil
	}

	var result []string
	diags := tfList.ElementsAs(ctx, &result, false)
	if diags.HasError() {
		log.Printf("[ERROR] ConvertListToStringSlice: failed to convert list: %v", diags)
		return nil
	}

	return result
}
