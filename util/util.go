/*
 * Copyright (c) 2025 Saviynt Inc.
 * All Rights Reserved.
 *
 * This software is the confidential and proprietary information of
 * Saviynt Inc. ("Confidential Information"). You shall not disclose,
 * use, or distribute such Confidential Information except in accordance
 * with the terms of the license agreement you entered into with Saviynt.
 *
 * SAVIYNT MAKES NO REPRESENTATIONS OR WARRANTIES ABOUT THE SUITABILITY OF
 * THE SOFTWARE, EITHER EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO
 * THE IMPLIED WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR
 * PURPOSE, OR NON-INFRINGEMENT.
 */

package util

import (
	"encoding/json"
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

func BoolPointerOrEmtpy(tfBool types.Bool) *bool {
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

func LoadConnectorDataForEphemeral(filePath string) map[string]string {
	data, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatalf("failed to read test config file: %v", err)
	}

	var allData map[string]interface{}
	if err := json.Unmarshal(data, &allData); err != nil {
		log.Fatalf("failed to unmarshal test config: %v", err)
	}

	result := make(map[string]string)
	for key, value := range allData {
		switch v := value.(type) {
		case string:
			result[key] = v

		case map[string]interface{}, []interface{}:
			jsonValue, err := json.Marshal(v)
			if err != nil {
				log.Fatalf("failed to marshal key %s: %v", key, err)
			}
			result[key] = string(jsonValue)

		default:
			// Marshal any other types safely (e.g., numbers, bool)
			jsonValue, err := json.Marshal(v)
			if err != nil {
				log.Fatalf("failed to marshal key %s: %v", key, err)
			}
			result[key] = string(jsonValue)
		}
	}

	return result
}

func TypesStringOrOriginal(original types.String, apiValue *string) types.String {
	if apiValue != nil && *apiValue != "" {
		return types.StringValue(*apiValue)
	}
	return original
}

func SafeStringAlt(s *string, replace string) types.String {
	if types.StringValue(*s) == types.StringValue("") {
		return types.StringValue(replace)
	}

	return types.StringValue(*s)
}

// func PreserveString(apiVal *string, oldVal types.String) types.String {
// 	if apiVal != nil {
// 		return types.StringValue(*apiVal)
// 	}
// 	if !oldVal.IsNull() && !oldVal.IsUnknown() {
// 		return oldVal
// 	}
// 	return types.StringNull()
// }
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
