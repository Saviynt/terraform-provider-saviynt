// Copyright (c) 2025 Saviynt Inc.
// SPDX-License-Identifier: MPL-2.0

// rest_connection_resource_test.go contains the Terraform acceptance test suite for
// the `saviynt_rest_connection_resource`. It validates the resource’s full lifecycle against
// a live Saviynt instance by exercising:
//
//   - Create: provisions a new Rest connection and asserts each attribute via JSONPath checks.
//   - Import: verifies the resource can be imported by its connection name.
//   - Update: applies configuration changes and confirms the updated attribute values.
//   - Negative Cases: ensures updates to `connection_name` and `connection_type` are rejected.
//
// Test data is loaded from `rest_connection_resource_test_data.json` using `testutil.LoadConnectorData`.
// Environment variables `SAVIYNT_URL`, `SAVIYNT_USERNAME`, and `SAVIYNT_PASSWORD` must be set
// to point at a valid Saviynt Security Manager before running these tests.
package testing

import (
	"fmt"
	"os"
	"regexp"
	"terraform-provider-Saviynt/util/testutil"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
)

func TestAccSaviyntRESTConnectionResource(t *testing.T) {
	filePath := testutil.GetTestDataPath(t, "./test_data/rest_connection_test_data.json")
	filePath = testutil.PrepareTestDataWithEnv(t, filePath)
	createCfg := testutil.LoadConnectorData(t, filePath, "create")
	updateCfg := testutil.LoadConnectorData(t, filePath, "update")
	resourceName := "saviynt_rest_connection_resource.rest"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { TestAccPreCheck(t) },
		ProtoV6ProviderFactories: TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create
			{
				Config: testAccRESTConnectionResourceConfig(filePath, "create"),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("connection_type"), knownvalue.StringExact(createCfg["connection_type"])),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("connection_name"), knownvalue.StringExact(createCfg["connection_name"])),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("import_user_json"), knownvalue.StringExact(createCfg["import_user_json"])),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("create_account_json"), knownvalue.StringExact(createCfg["create_account_json"])),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("update_account_json"), knownvalue.StringExact(createCfg["update_account_json"])),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("add_access_json"), knownvalue.StringExact(createCfg["add_access_json"])),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("remove_access_json"), knownvalue.StringExact(createCfg["remove_access_json"])),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("remove_account_json"), knownvalue.StringExact(createCfg["remove_account_json"])),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("import_account_ent_json"), knownvalue.StringExact(createCfg["import_account_ent_json"])),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("disable_account_json"), knownvalue.StringExact(createCfg["disable_account_json"])),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("enable_account_json"), knownvalue.StringExact(createCfg["enable_account_json"])),
				},
			},
			// Import
			{
				ResourceName:            resourceName,
				ImportStateId:           createCfg["connection_name"],
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"msg", "error_code", "connection_json", "change_pass_json"},
			},
			// Update
			{
				Config: testAccRESTConnectionResourceConfig(filePath, "update"),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("connection_type"), knownvalue.StringExact(updateCfg["connection_type"])),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("connection_name"), knownvalue.StringExact(updateCfg["connection_name"])),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("import_user_json"), knownvalue.StringExact(updateCfg["import_user_json"])),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("create_account_json"), knownvalue.StringExact(updateCfg["create_account_json"])),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("update_account_json"), knownvalue.StringExact(updateCfg["update_account_json"])),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("add_access_json"), knownvalue.StringExact(updateCfg["add_access_json"])),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("remove_access_json"), knownvalue.StringExact(updateCfg["remove_access_json"])),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("remove_account_json"), knownvalue.StringExact(updateCfg["remove_account_json"])),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("import_account_ent_json"), knownvalue.StringExact(updateCfg["import_account_ent_json"])),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("disable_account_json"), knownvalue.StringExact(updateCfg["disable_account_json"])),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("enable_account_json"), knownvalue.StringExact(updateCfg["enable_account_json"])),
				},
			},
			// Update the Connectionname to a new value
			{
				Config:      testAccRESTConnectionResourceConfig(filePath, "update_connection_name"),
				ExpectError: regexp.MustCompile(`Connection name cannot be updated`),
			},
			// Update the Connectiontype to a new value
			{
				Config:      testAccRESTConnectionResourceConfig(filePath, "update_connection_type"),
				ExpectError: regexp.MustCompile(`Connection type cannot by updated`),
			},
		},
	})
}

func testAccRESTConnectionResourceConfig(jsonPath, operation string) string {
	return fmt.Sprintf(`
provider "saviynt" {
  server_url = "%s"
  username   = "%s"
  password   = "%s"
}

locals {
  cfg = jsondecode(file("%s"))["%s"]
}

resource "saviynt_rest_connection_resource" "rest" {
  connection_type            = local.cfg.connection_type
  connection_name            = local.cfg.connection_name
  connection_json            = jsonencode(local.cfg.connection_json)
  import_user_json           = jsonencode(local.cfg.import_user_json)
  create_account_json        = jsonencode(local.cfg.create_account_json)
  update_account_json        = jsonencode(local.cfg.update_account_json)
  add_access_json            = jsonencode(local.cfg.add_access_json)
  remove_access_json         = jsonencode(local.cfg.remove_access_json)
  remove_account_json        = jsonencode(local.cfg.remove_account_json)
  import_account_ent_json    = jsonencode(local.cfg.import_account_ent_json)
  change_pass_json           = jsonencode(local.cfg.change_pass_json)
  disable_account_json       = jsonencode(local.cfg.disable_account_json)
  enable_account_json        = jsonencode(local.cfg.enable_account_json)
}
`, os.Getenv("SAVIYNT_URL"),
		os.Getenv("SAVIYNT_USERNAME"),
		os.Getenv("SAVIYNT_PASSWORD"),
		jsonPath,
		operation,
	)
}
