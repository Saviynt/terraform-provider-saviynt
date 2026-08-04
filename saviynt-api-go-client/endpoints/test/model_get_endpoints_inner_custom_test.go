// Copyright (c) 2025 Saviynt Inc.
// SPDX-License-Identifier: MPL-2.0

package endpoints

import (
	"encoding/json"
	"testing"

	openapiclient "github.com/saviynt/saviynt-api-go-client/endpoints"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestUnmarshalJSON_SpacedKeys verifies that CP1–CP30 are correctly populated
// when the server returns spaced JSON keys ("Custom Property 1").
// This is the default server configuration and must continue to work.
func TestUnmarshalJSON_SpacedKeys(t *testing.T) {
	data := `{
		"Custom Property 1": "BusinessUnit",
		"Custom Property 2": "ApplicationName",
		"Custom Property 15": "Division",
		"Custom Property 30": "BadgeNumber"
	}`

	var result openapiclient.GetEndpoints200ResponseEndpointsInner
	err := json.Unmarshal([]byte(data), &result)
	require.NoError(t, err)

	assert.Equal(t, "BusinessUnit", *result.CustomProperty1)
	assert.Equal(t, "ApplicationName", *result.CustomProperty2)
	assert.Equal(t, "Division", *result.CustomProperty15)
	assert.Equal(t, "BadgeNumber", *result.CustomProperty30)
}

// TestUnmarshalJSON_LowercaseKeys verifies that CP1–CP30 are correctly populated
// when the server returns lowercase JSON keys ("customproperty1").
// This is the configuration that was previously broken.
func TestUnmarshalJSON_LowercaseKeys(t *testing.T) {
	data := `{
		"customproperty1": "BusinessUnit",
		"customproperty2": "ApplicationName",
		"customproperty15": "Division",
		"customproperty30": "BadgeNumber"
	}`

	var result openapiclient.GetEndpoints200ResponseEndpointsInner
	err := json.Unmarshal([]byte(data), &result)
	require.NoError(t, err)

	assert.Equal(t, "BusinessUnit", *result.CustomProperty1)
	assert.Equal(t, "ApplicationName", *result.CustomProperty2)
	assert.Equal(t, "Division", *result.CustomProperty15)
	assert.Equal(t, "BadgeNumber", *result.CustomProperty30)
}

// TestUnmarshalJSON_LowercaseKeys_CP31to45_Unaffected verifies that CP31–CP45
// continue to work with their existing lowercase tags and are not broken by
// the custom unmarshaler.
func TestUnmarshalJSON_LowercaseKeys_CP31to45_Unaffected(t *testing.T) {
	data := `{
		"customproperty31": "ContractorID",
		"customproperty45": "PriorityID"
	}`

	var result openapiclient.GetEndpoints200ResponseEndpointsInner
	err := json.Unmarshal([]byte(data), &result)
	require.NoError(t, err)

	assert.Equal(t, "ContractorID", *result.Customproperty31)
	assert.Equal(t, "PriorityID", *result.Customproperty45)
}

// TestUnmarshalJSON_SpacedKeyTakesPrecedence verifies that if both spaced and
// lowercase keys are present for the same field, the spaced key value wins
// (pass 1 fills it, pass 2 skips non-nil fields).
func TestUnmarshalJSON_SpacedKeyTakesPrecedence(t *testing.T) {
	data := `{
		"Custom Property 1": "from-spaced-key",
		"customproperty1": "from-lowercase-key"
	}`

	var result openapiclient.GetEndpoints200ResponseEndpointsInner
	err := json.Unmarshal([]byte(data), &result)
	require.NoError(t, err)

	assert.Equal(t, "from-spaced-key", *result.CustomProperty1)
}

// TestUnmarshalJSON_AbsentField verifies that fields absent from the JSON
// remain nil (omitempty behaviour preserved).
func TestUnmarshalJSON_AbsentField(t *testing.T) {
	data := `{"customproperty1": "BusinessUnit"}`

	var result openapiclient.GetEndpoints200ResponseEndpointsInner
	err := json.Unmarshal([]byte(data), &result)
	require.NoError(t, err)

	assert.Equal(t, "BusinessUnit", *result.CustomProperty1)
	assert.Nil(t, result.CustomProperty2)
}

// TestUnmarshalJSON_InvalidJSON verifies that malformed JSON returns an error.
func TestUnmarshalJSON_InvalidJSON(t *testing.T) {
	data := `{invalid json}`

	var result openapiclient.GetEndpoints200ResponseEndpointsInner
	err := json.Unmarshal([]byte(data), &result)

	assert.Error(t, err)
}

// TestUnmarshalJSON_OtherFields verifies that non-custom-property fields are
// still correctly populated and not disturbed by the custom unmarshaler.
func TestUnmarshalJSON_OtherFields(t *testing.T) {
	data := `{
		"endpointname": "MyEndpoint",
		"displayName": "My Endpoint",
		"customproperty1": "BusinessUnit"
	}`

	var result openapiclient.GetEndpoints200ResponseEndpointsInner
	err := json.Unmarshal([]byte(data), &result)
	require.NoError(t, err)

	assert.Equal(t, "MyEndpoint", *result.Endpointname)
	assert.Equal(t, "My Endpoint", *result.DisplayName)
	assert.Equal(t, "BusinessUnit", *result.CustomProperty1)
}
