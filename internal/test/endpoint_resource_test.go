// Copyright (c) 2025 Saviynt Inc.
// SPDX-License-Identifier: MPL-2.0

// endpoint_resource_test.go contains the Terraform acceptance test suite for
// the `saviynt_endpoint_resource`. It validates the resource’s full lifecycle against
// a live Saviynt instance by exercising:
//
//   - Create: provisions a new endpoint and asserts each attribute via JSONPath checks.
//   - Import: verifies the resource can be imported by its endpoint name.
//   - Update: applies configuration changes and confirms the updated attribute values.
//   - Negative Cases: ensures updates to `endpoint_name` and `duplicate_endpoint_name` are rejected.
//
// Test data is loaded from `endpoint_resource_test_data.json` using `testutil.LoadConnectorData`.
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

func TestAccSaviyntEndpointResource(t *testing.T) {
	filePath := testutil.GetTestDataPath(t, "./test_data/endpoint_resource_test_data.json")
	filePath = testutil.PrepareTestDataWithEnv(t, filePath)
	createCfg := testutil.LoadConnectorData(t, filePath, "create")
	updateCfg := testutil.LoadConnectorData(t, filePath, "update")
	resourceName := "saviynt_endpoint_resource.endpoint"

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { TestAccPreCheck(t) },
		ProtoV6ProviderFactories: TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create
			{
				Config: testAccEndpointResourceConfig(filePath, "create"),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("endpoint_name"), knownvalue.StringExact(createCfg["endpoint_name"])),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("display_name"), knownvalue.StringExact(createCfg["display_name"])),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("security_system"), knownvalue.StringExact(createCfg["security_system"])),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("owner_type"), knownvalue.StringExact(createCfg["owner_type"])),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("custom_property1"), knownvalue.StringExact(createCfg["custom_property1"])),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("account_custom_property_1_label"), knownvalue.StringExact(createCfg["account_custom_property_1_label"])),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("custom_property31_label"), knownvalue.StringExact(createCfg["custom_property31_label"])),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("enable_copy_access"), knownvalue.StringExact(createCfg["enable_copy_access"])),
				},
			},
			// Import
			{
				ResourceName:            resourceName,
				ImportStateId:           createCfg["endpoint_name"],
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"endpoint_config", "connection_config"},
			},
			// Update
			{
				Config: testAccEndpointResourceConfig(filePath, "update"),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("endpoint_name"), knownvalue.StringExact(updateCfg["endpoint_name"])),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("display_name"), knownvalue.StringExact(updateCfg["display_name"])),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("security_system"), knownvalue.StringExact(updateCfg["security_system"])),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("owner_type"), knownvalue.StringExact(updateCfg["owner_type"])),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("custom_property1"), knownvalue.StringExact(updateCfg["custom_property1"])),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("account_custom_property_1_label"), knownvalue.StringExact(updateCfg["account_custom_property_1_label"])),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("custom_property31_label"), knownvalue.StringExact(updateCfg["custom_property31_label"])),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("enable_copy_access"), knownvalue.StringExact(updateCfg["enable_copy_access"])),
				},
			},
			// Update the Endpoint name to a new value
			{
				Config:      testAccEndpointResourceConfig(filePath, "update_endpoint_name"),
				ExpectError: regexp.MustCompile(`Endpoint name cannot be updated`),
			},
			// Create a new resource with the same Endpoint name
			{
				Config:      testAccEndpointWithSameNameConfig(filePath, "create_duplicate_endpoint"),
				ExpectError: regexp.MustCompile(`Endpoint name already exists`),
			},
		},
	})
}

func testAccEndpointResourceConfig(jsonPath, operation string) string {
	return fmt.Sprintf(`
provider "saviynt" {
  server_url = "%s"
  username   = "%s"
  password   = "%s"
}

locals {
  cfg = jsondecode(file("%s"))["%s"]
}

resource "saviynt_endpoint_resource" "endpoint" {
  endpoint_name                      = local.cfg.endpoint_name
  display_name                       = local.cfg.display_name
  security_system                    = local.cfg.security_system
  owner_type					   	 = local.cfg.owner_type	
  custom_property1            		 = local.cfg.custom_property1
  account_custom_property_1_label 	 = local.cfg.account_custom_property_1_label
  custom_property31_label 			 = local.cfg.custom_property31_label
  enable_copy_access				 = local.cfg.enable_copy_access
}
`, os.Getenv("SAVIYNT_URL"),
		os.Getenv("SAVIYNT_USERNAME"),
		os.Getenv("SAVIYNT_PASSWORD"),
		jsonPath,
		operation,
	)
}

func testAccEndpointWithSameNameConfig(jsonPath, operation string) string {
	return fmt.Sprintf(`
provider "saviynt" {
  server_url = "%s"
  username   = "%s"
  password   = "%s"
}

locals {
  cfg = jsondecode(file("%s"))["%s"]
}

resource "saviynt_endpoint_resource" "endpoint_1" {
  endpoint_name                      = local.cfg.endpoint_name
  display_name                       = local.cfg.display_name
  security_system                    = local.cfg.security_system
  owner_type					   	 = local.cfg.owner_type	
  custom_property1            		 = local.cfg.custom_property1
  account_custom_property_1_label 	 = local.cfg.account_custom_property_1_label
  custom_property31_label 			 = local.cfg.custom_property31_label
  enable_copy_access				 = local.cfg.enable_copy_access
}
`, os.Getenv("SAVIYNT_URL"),
		os.Getenv("SAVIYNT_USERNAME"),
		os.Getenv("SAVIYNT_PASSWORD"),
		jsonPath,
		operation,
	)
}
