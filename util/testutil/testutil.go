// Copyright (c) 2025 Saviynt Inc.
// SPDX-License-Identifier: MPL-2.0

package testutil

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
	"time"
)

func LoadConnectorData(t *testing.T, filePath, scenario string) map[string]string {
	// Load base JSON file
	data, err := os.ReadFile(filePath)
	if err != nil {
		t.Fatalf("failed to read test config file: %v", err)
	}

	var allData map[string]map[string]interface{}
	if err := json.Unmarshal(data, &allData); err != nil {
		t.Fatalf("failed to unmarshal test config: %v", err)
	}

	// Select create or update based on the scenario
	secrets, exists := allData[scenario]
	if !exists {
		t.Fatalf("scenario '%s' not found in config", scenario)
	}

	result := make(map[string]string)
	for key, value := range secrets {
		switch v := value.(type) {
		case string:
			// Keep plain strings as-is
			result[key] = v

		case map[string]interface{}:
			// Handle nested maps (e.g., status_key_json)
			jsonValue, err := json.Marshal(v)
			if err != nil {
				t.Fatalf("failed to marshal map for key %s: %v", key, err)
			}
			result[key] = string(jsonValue)

		default:
			// Fallback for other types, including arrays or deeply nested data
			jsonValue, err := json.Marshal(value)
			if err != nil {
				t.Fatalf("failed to marshal value for key %s: %v", key, err)
			}
			result[key] = string(jsonValue)
		}
	}

	return result
}

func GetTestDataPath(t *testing.T, relativeFilename string) string {
	t.Helper()

	// Get the full path of the test file that called this function
	_, callerFile, _, ok := runtime.Caller(1)
	if !ok {
		t.Fatalf("unable to determine caller path")
	}

	// Build the absolute path to the data file relative to the caller file
	absPath := filepath.Join(filepath.Dir(callerFile), relativeFilename)
	return absPath
}

// GetTestIdentifier returns a timestamp-based identifier for test resources
// in the format YYYYMMDDHHMMSS (e.g., "20251106050300")
func GetTestIdentifier() string {
	return time.Now().Format("20060102150405")
}

func PrepareTestDataWithEnv(t *testing.T, templatePath string) string {
	t.Helper()

	data, err := os.ReadFile(templatePath)
	if err != nil {
		t.Fatalf("failed to read test data template: %v", err)
	}

	content := string(data)

	// Use timestamp for unique resource names
	timestamp := GetTestIdentifier()

	replacements := map[string]string{
		"{{TEST_ID}}": timestamp,
	}

	// Add fixture replacements from environment variables (set by TestMain)
	fixtureEnvMappings := map[string]string{
		"{{FIXTURE_SECURITY_SYSTEM}}":    "TF_FIXTURE_SECURITY_SYSTEM",
		"{{FIXTURE_ENDPOINT_1}}":         "TF_FIXTURE_ENDPOINT_1",
		"{{FIXTURE_ENDPOINT_2}}":         "TF_FIXTURE_ENDPOINT_2",
		"{{FIXTURE_ENDPOINT_3}}":         "TF_FIXTURE_ENDPOINT_3",
		"{{FIXTURE_ENDPOINT_4}}":         "TF_FIXTURE_ENDPOINT_4",
		"{{FIXTURE_ENDPOINT_5}}":         "TF_FIXTURE_ENDPOINT_5",
		"{{FIXTURE_ENTITLEMENT_TYPE}}":   "TF_FIXTURE_ENTITLEMENT_TYPE",
		"{{FIXTURE_ENTITLEMENT_TYPE_2}}": "TF_FIXTURE_ENTITLEMENT_TYPE_2",
		"{{FIXTURE_ENTITLEMENT}}":        "TF_FIXTURE_ENTITLEMENT",
		"{{FIXTURE_ENTITLEMENT_2}}":      "TF_FIXTURE_ENTITLEMENT_2",
		"{{FIXTURE_ROLE}}":               "TF_FIXTURE_ROLE",
		"{{FIXTURE_REST_CONNECTION}}":    "TF_FIXTURE_REST_CONNECTION",
		"{{FIXTURE_REST_CONNECTION_ID}}": "TF_FIXTURE_REST_CONNECTION_ID",
	}
	for placeholder, envVar := range fixtureEnvMappings {
		if val := os.Getenv(envVar); val != "" {
			replacements[placeholder] = val
		}
	}

	for placeholder, value := range replacements {
		content = strings.ReplaceAll(content, placeholder, value)
	}

	tmpFile := filepath.Join(t.TempDir(), filepath.Base(templatePath))
	if err := os.WriteFile(tmpFile, []byte(content), 0644); err != nil {
		t.Fatalf("failed to write processed test data: %v", err)
	}

	log.Printf("New file path: %v", tmpFile)
	return tmpFile
}

// PrepareTestDataWithReplacements processes a test data template with custom replacements
// in addition to the standard {{TEST_ID}} replacement.
func PrepareTestDataWithReplacements(t *testing.T, templatePath string, extraReplacements map[string]string) string {
	t.Helper()

	data, err := os.ReadFile(templatePath)
	if err != nil {
		t.Fatalf("failed to read test data template: %v", err)
	}

	content := string(data)

	// Use timestamp for unique resource names
	timestamp := GetTestIdentifier()

	// Start with standard replacements
	replacements := map[string]string{
		"{{TEST_ID}}": timestamp,
	}

	// Add extra replacements
	for k, v := range extraReplacements {
		replacements[k] = v
	}

	for placeholder, value := range replacements {
		content = strings.ReplaceAll(content, placeholder, value)
	}

	tmpFile := filepath.Join(t.TempDir(), filepath.Base(templatePath))
	if err := os.WriteFile(tmpFile, []byte(content), 0644); err != nil {
		t.Fatalf("failed to write processed test data: %v", err)
	}

	log.Printf("New file path: %v", tmpFile)
	return tmpFile
}
