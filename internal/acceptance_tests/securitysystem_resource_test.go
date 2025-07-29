// Copyright (c) 2025 Saviynt Inc.
// SPDX-License-Identifier: MPL-2.0

// securitysystem_resource_test.go contains the Terraform acceptance test suite for
// the `saviynt_security_system_resource`. It validates the resource’s full lifecycle against
// a live Saviynt instance by exercising:
//
//   - Create: provisions a new security system connection and asserts each attribute via JSONPath checks.
//   - Import: verifies the resource can be imported by its security system name.
//   - Update: applies configuration changes and confirms the updated attribute values.
//   - Negative Cases: ensures updates to `systemname` and `duplicate_systemname_name` are rejected.
//
// Test data is loaded from `securitysystem_resource_test_data.json` using `testutil.LoadConnectorData`.
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

func TestAccSaviyntSecuritySystemResource(t *testing.T) {
	filePath := testutil.GetTestDataPath(t, "./test_data/securitysystem_resource_test_data.json")
	filePath = testutil.PrepareTestDataWithEnv(t, filePath)
	createCfg := testutil.LoadConnectorData(t, filePath, "create")
	updateCfg := testutil.LoadConnectorData(t, filePath, "update")
	resourceName := "saviynt_security_system_resource.ss"
	securitySystemName := createCfg["systemname"]
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { TestAccPreCheck(t) },
		ProtoV6ProviderFactories: TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create
			{
				Config: testAccSecuritySystemConnectionResourceConfig(filePath, "create"),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("systemname"), knownvalue.StringExact(createCfg["systemname"])),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("display_name"), knownvalue.StringExact(createCfg["display_name"])),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("access_add_workflow"), knownvalue.StringExact(createCfg["access_add_workflow"])),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("access_remove_workflow"), knownvalue.StringExact(createCfg["access_remove_workflow"])),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("add_service_account_workflow"), knownvalue.StringExact(createCfg["add_service_account_workflow"])),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("remove_service_account_workflow"), knownvalue.StringExact(createCfg["remove_service_account_workflow"])),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("automated_provisioning"), knownvalue.StringExact(createCfg["automated_provisioning"])),
				},
			},
			// Import
			{
				ResourceName:      resourceName,
				ImportStateId:     createCfg["systemname"],
				ImportState:       true,
				ImportStateVerify: true,
			},
			// Update
			{
				Config: testAccSecuritySystemConnectionResourceConfig(filePath, "update"),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("systemname"), knownvalue.StringExact(updateCfg["systemname"])),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("display_name"), knownvalue.StringExact(updateCfg["display_name"])),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("access_add_workflow"), knownvalue.StringExact(updateCfg["access_add_workflow"])),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("access_remove_workflow"), knownvalue.StringExact(updateCfg["access_remove_workflow"])),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("add_service_account_workflow"), knownvalue.StringExact(updateCfg["add_service_account_workflow"])),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("remove_service_account_workflow"), knownvalue.StringExact(updateCfg["remove_service_account_workflow"])),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("automated_provisioning"), knownvalue.StringExact(updateCfg["automated_provisioning"])),
				},
			},
			// Update the Systemname to a new value
			{
				Config:      testAccSecuritySystemConnectionResourceConfig(filePath, "update_security_system_name"),
				ExpectError: regexp.MustCompile(`System name cannot be updated`),
			},
			// Create a new resource with the same Systemname
			{
				Config: testAccSecuritySecuritySystemWithSameNameConfig(filePath, "create_duplicate_security_system"),
				ExpectError: regexp.MustCompile(
					fmt.Sprintf(`systemname %s already exists`, securitySystemName),
				),
			},
		},
	})
}

func testAccSecuritySystemConnectionResourceConfig(jsonPath, operation string) string {
	return fmt.Sprintf(`
provider "saviynt" {
  server_url = "%s"
  username   = "%s"
  password   = "%s"
}

locals {
  cfg = jsondecode(file("%s"))["%s"]
}

resource "saviynt_security_system_resource" "ss" {
  systemname                         = local.cfg.systemname
  display_name                       = local.cfg.display_name
  access_add_workflow                = local.cfg.access_add_workflow
  access_remove_workflow             = local.cfg.access_remove_workflow
  add_service_account_workflow       = local.cfg.add_service_account_workflow
  remove_service_account_workflow    = local.cfg.remove_service_account_workflow
  automated_provisioning             = local.cfg.automated_provisioning
}
`, os.Getenv("SAVIYNT_URL"),
		os.Getenv("SAVIYNT_USERNAME"),
		os.Getenv("SAVIYNT_PASSWORD"),
		jsonPath,
		operation,
	)
}

func testAccSecuritySecuritySystemWithSameNameConfig(jsonPath, operation string) string {
	return fmt.Sprintf(`
provider "saviynt" {
  server_url = "%s"
  username   = "%s"
  password   = "%s"
}

locals {
  cfg = jsondecode(file("%s"))["%s"]
}

resource "saviynt_security_system_resource" "ss1" {
  systemname                         = local.cfg.systemname
  display_name                       = local.cfg.display_name
  access_add_workflow                = local.cfg.access_add_workflow
  access_remove_workflow             = local.cfg.access_remove_workflow
  add_service_account_workflow       = local.cfg.add_service_account_workflow
  remove_service_account_workflow    = local.cfg.remove_service_account_workflow
  automated_provisioning             = local.cfg.automated_provisioning
}
`, os.Getenv("SAVIYNT_URL"),
		os.Getenv("SAVIYNT_USERNAME"),
		os.Getenv("SAVIYNT_PASSWORD"),
		jsonPath,
		operation,
	)
}
