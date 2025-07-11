// Copyright (c) Saviynt Inc.
// SPDX-License-Identifier: MPL-2.0

// sap_connection_resource_test.go contains the Terraform acceptance test suite for
// the `saviynt_sap_connection_resource`. It validates the resource’s full lifecycle against
// a live Saviynt instance by exercising:
//
//   - Create: provisions a new Sap connection and asserts each attribute via JSONPath checks.
//   - Import: verifies the resource can be imported by its connection name.
//   - Update: applies configuration changes and confirms the updated attribute values.
//   - Negative Cases: ensures updates to `connection_name` and `connection_type` are rejected.
//
// Test data is loaded from `sap_connection_resource_test_data.json` using `testutil.LoadConnectorData`.
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

func TestAccSaviyntSAPConnectionResource(t *testing.T) {
	filePath := testutil.GetTestDataPath(t, "./test_data/sap_connection_test_data.json")
	filePath = testutil.PrepareTestDataWithEnv(t, filePath)
	createCfg := testutil.LoadConnectorData(t, filePath, "create")
	updateCfg := testutil.LoadConnectorData(t, filePath, "update")
	resourceName := "saviynt_sap_connection_resource.sp"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { TestAccPreCheck(t) },
		ProtoV6ProviderFactories: TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create Step
			{
				Config: testAccSAPConnectionResourceConfig(filePath, "create"),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("connection_name"), knownvalue.StringExact(createCfg["connection_name"])),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("connection_type"), knownvalue.StringExact(createCfg["connection_type"])),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("message_server"), knownvalue.StringExact(createCfg["message_server"])),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("jco_snc_mode"), knownvalue.StringExact(createCfg["jco_snc_mode"])),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("system_name"), knownvalue.StringExact(createCfg["system_name"])),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("user_import_json"), knownvalue.StringExact(createCfg["user_import_json"])),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("password_max_length"), knownvalue.StringExact(createCfg["password_max_length"])),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("update_account_json"), knownvalue.StringExact(createCfg["update_account_json"])),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("set_cua_system"), knownvalue.StringExact(createCfg["set_cua_system"])),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("saptable_filter_lang"), knownvalue.StringExact(createCfg["saptable_filter_lang"])),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("error_code"), knownvalue.StringExact("0")),
				},
			},
			// Import
			{
				ResourceName:            resourceName,
				ImportStateId:           createCfg["connection_name"],
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"msg", "config_json"},
			},
			// Update Step
			{
				Config: testAccSAPConnectionResourceConfig(filePath, "update"),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("connection_name"), knownvalue.StringExact(updateCfg["connection_name"])),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("connection_type"), knownvalue.StringExact(updateCfg["connection_type"])),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("message_server"), knownvalue.StringExact(updateCfg["message_server"])),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("jco_snc_mode"), knownvalue.StringExact(updateCfg["jco_snc_mode"])),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("system_name"), knownvalue.StringExact(updateCfg["system_name"])),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("user_import_json"), knownvalue.StringExact(updateCfg["user_import_json"])),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("password_max_length"), knownvalue.StringExact(updateCfg["password_max_length"])),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("update_account_json"), knownvalue.StringExact(updateCfg["update_account_json"])),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("set_cua_system"), knownvalue.StringExact(updateCfg["set_cua_system"])),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("saptable_filter_lang"), knownvalue.StringExact(updateCfg["saptable_filter_lang"])),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("error_code"), knownvalue.StringExact("0")),
				},
			},
			// Update the Connectionname to a new value
			{
				Config:      testAccSAPConnectionResourceConfig(filePath, "update_connection_name"),
				ExpectError: regexp.MustCompile(`Connection name cannot be updated`),
			},
			// Update the Connectiontype to a new value
			{
				Config:      testAccSAPConnectionResourceConfig(filePath, "update_connection_type"),
				ExpectError: regexp.MustCompile(`Connection type cannot by updated`),
			},
		},
	})
}

func testAccSAPConnectionResourceConfig(jsonPath, operation string) string {
	return fmt.Sprintf(`
	provider "saviynt" {
  server_url = "%s"
  username   = "%s"
  password   = "%s"
}
  locals {
  cfg = jsondecode(file("%s"))["%s"]
}

  resource "saviynt_sap_connection_resource" "sp" {
  connection_type                    = local.cfg.connection_type
  connection_name                    = local.cfg.connection_name
  message_server                     = local.cfg.message_server
  jco_snc_mode                       = local.cfg.jco_snc_mode
  system_name                        = local.cfg.system_name
  user_import_json                   =jsonencode(local.cfg.user_import_json)
  password_max_length                = local.cfg.password_max_length
  update_account_json                = jsonencode(local.cfg.update_account_json)
  set_cua_system                     = local.cfg.set_cua_system
  saptable_filter_lang               = local.cfg.saptable_filter_lang
}
  `, os.Getenv("SAVIYNT_URL"),
		os.Getenv("SAVIYNT_USERNAME"),
		os.Getenv("SAVIYNT_PASSWORD"),
		jsonPath,
		operation,
	)
}
