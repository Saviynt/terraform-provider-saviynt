// Copyright (c) Saviynt Inc.
// SPDX-License-Identifier: MPL-2.0

// workday_connection_resource_test.go contains the Terraform acceptance test suite for
// the `saviynt_workday_connection_resource`. It validates the resource’s full lifecycle against
// a live Saviynt instance by exercising:
//
//   - Create: provisions a new Workday connection and asserts each attribute via JSONPath checks.
//   - Import: verifies the resource can be imported by its connection name.
//   - Update: applies configuration changes and confirms the updated attribute values.
//   - Negative Cases: ensures updates to `connection_name` and `connection_type` are rejected.
//
// Test data is loaded from `workday _connection_resource_test_data.json` using `testutil.LoadConnectorData`.
// Environment variables `SAVIYNT_URL`, `SAVIYNT_USERNAME`, and `SAVIYNT_PASSWORD` must be set
// to point at a valid Saviynt Security Manager before running these tests.
package testing

import (
	"fmt"
	// "io"
	"os"
	// "path/filepath"
	"regexp"
	// "runtime"
	"terraform-provider-Saviynt/util/testutil"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
)

// func copyTestDataFile(t *testing.T, filename string) string {
// 	_, callerFile, _, _ := runtime.Caller(0)
// 	src := filepath.Join(filepath.Dir(callerFile), filename)
// 	dst := filepath.Join(os.TempDir(), filename)

// 	from, err := os.Open(src)
// 	if err != nil {
// 		t.Fatalf("unable to open source test data file: %v", err)
// 	}
// 	defer from.Close()

// 	to, err := os.Create(dst)
// 	if err != nil {
// 		t.Fatalf("unable to create test data file in temp directory: %v", err)
// 	}
// 	defer to.Close()

// 	if _, err := io.Copy(to, from); err != nil {
// 		t.Fatalf("error copying test data file: %v", err)
// 	}

// 	return dst
// }

func TestAccSaviyntWorkdayConnectionResource(t *testing.T) {
	filePath := testutil.GetTestDataPath(t, "./test_data/workday_connection_test_data.json")
	filePath = testutil.PrepareTestDataWithEnv(t, filePath)
	createCfg := testutil.LoadConnectorData(t, filePath, "create")
	updateCfg := testutil.LoadConnectorData(t, filePath, "update")
	resourceName := "saviynt_workday_connection_resource.w"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { TestAccPreCheck(t) },
		ProtoV6ProviderFactories: TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create Step
			{
				Config: testAccWorkdayConnectionResourceConfig(filePath, "create"),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("connection_name"), knownvalue.StringExact(createCfg["connection_name"])),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("connection_type"), knownvalue.StringExact(createCfg["connection_type"])),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("base_url"), knownvalue.StringExact(createCfg["base_url"])),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("api_version"), knownvalue.StringExact(createCfg["api_version"])),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("use_oauth"), knownvalue.StringExact(createCfg["use_oauth"])),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("access_import_list"), knownvalue.StringExact(createCfg["access_import_list"])),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("status_key_json"), knownvalue.StringExact(createCfg["status_key_json"])),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("user_import_payload"), knownvalue.StringExact(createCfg["user_import_payload"])),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("user_import_mapping"), knownvalue.StringExact(createCfg["user_import_mapping"])),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("error_code"), knownvalue.StringExact("0")),
				},
			},
			// Import
			{
				ResourceName:            resourceName,
				ImportStateId:           createCfg["connection_name"],
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"msg", "client_secret", "password", "refresh_token"},
			},
			// Update Step
			{
				Config: testAccWorkdayConnectionResourceConfig(filePath, "update"),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("connection_name"), knownvalue.StringExact(updateCfg["connection_name"])),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("connection_type"), knownvalue.StringExact(updateCfg["connection_type"])),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("base_url"), knownvalue.StringExact(updateCfg["base_url"])),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("api_version"), knownvalue.StringExact(updateCfg["api_version"])),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("use_oauth"), knownvalue.StringExact(updateCfg["use_oauth"])),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("access_import_list"), knownvalue.StringExact(updateCfg["access_import_list"])),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("status_key_json"), knownvalue.StringExact(updateCfg["status_key_json"])),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("user_import_payload"), knownvalue.StringExact(updateCfg["user_import_payload"])),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("user_import_mapping"), knownvalue.StringExact(updateCfg["user_import_mapping"])),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("error_code"), knownvalue.StringExact("0")),
				},
			},
			// Update the Connectionname to a new value
			{
				Config:      testAccWorkdayConnectionResourceConfig(filePath, "update_connection_name"),
				ExpectError: regexp.MustCompile(`Connection name cannot be updated`),
			},
			// Update the Connectiontype to a new value
			{
				Config:      testAccWorkdayConnectionResourceConfig(filePath, "update_connection_type"),
				ExpectError: regexp.MustCompile(`Connection type cannot by updated`),
			},
		},
	},
	)
}

func testAccWorkdayConnectionResourceConfig(jsonPath, operation string) string {
	// jsonPath := "{filepath}/workday_connection_test_data.json"
	return fmt.Sprintf(`
	provider "saviynt" {
  server_url = "%s"
  username   = "%s"
  password   = "%s"
}
  locals {
  cfg = jsondecode(file("%s"))["%s"]
}

  resource "saviynt_workday_connection_resource" "w" {
  connection_type    = local.cfg.connection_type
  connection_name    = local.cfg.connection_name
  base_url           = local.cfg.base_url
  api_version        = local.cfg.api_version
  use_oauth          = local.cfg.use_oauth
  username           = local.cfg.username
  client_id          = local.cfg.client_id
  access_import_list = local.cfg.access_import_list
  status_key_json = jsonencode(local.cfg.status_key_json)
  user_import_payload = local.cfg.user_import_payload
  user_import_mapping = jsonencode(local.cfg.user_import_mapping)
  }`,
		os.Getenv("SAVIYNT_URL"),
		os.Getenv("SAVIYNT_USERNAME"),
		os.Getenv("SAVIYNT_PASSWORD"),
		jsonPath,
		operation,
	)
}
