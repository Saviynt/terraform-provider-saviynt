package testutil

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
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

func PrepareTestDataWithEnv(t *testing.T, templatePath string) string {
	t.Helper()

	data, err := os.ReadFile(templatePath)
	if err != nil {
		t.Fatalf("failed to read test data template: %v", err)
	}

	content := string(data)
	pipelineID := os.Getenv("CI_PIPELINE_ID")
	if pipelineID == "" {
		pipelineID = "local"
	}

	replacements := map[string]string{
		"{{PIPELINE_ID}}": pipelineID,
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
