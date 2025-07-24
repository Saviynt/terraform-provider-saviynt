// Copyright (c) 2025 Saviynt Inc.
// SPDX-License-Identifier: MPL-2.0

// db_connection_resource_test.go contains the Terraform acceptance test suite for
// the `saviynt_db_connection_resource`. It validates the resource’s full lifecycle against
// a live Saviynt instance by exercising:
//
//   - Create: provisions a new DB connection and asserts each attribute via JSONPath checks.
//   - Import: verifies the resource can be imported by its connection name.
//   - Update: applies configuration changes and confirms the updated attribute values.
//
// Test data is loaded from `db_connection_resource_test_data.json` using `testutil.LoadConnectorData`.
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

func TestAccSaviyntDBConnectionResource25B(t *testing.T) {
	filePath := testutil.GetTestDataPath(t, "./test_data/db_connection_test_data.json")
	filePath = testutil.PrepareTestDataWithEnv(t, filePath)
	createCfg := testutil.LoadConnectorData(t, filePath, "create")
	updateCfg := testutil.LoadConnectorData(t, filePath, "update")
	resourceName := "saviynt_db_connection_resource.db"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { TestAccPreCheck(t) },
		ProtoV6ProviderFactories: TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create
			{
				Config: testAccDBConnectionResourceConfig25B(filePath, "create"),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("connection_name"), knownvalue.StringExact(createCfg["connection_name"])),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("url"), knownvalue.StringExact(createCfg["url"])),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("driver_name"), knownvalue.StringExact(createCfg["driver_name"])),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("create_account_json"), knownvalue.StringExact(createCfg["create_account_json"])),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("update_account_json"), knownvalue.StringExact(createCfg["update_account_json"])),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("grant_access_json"), knownvalue.StringExact(createCfg["grant_access_json"])),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("revoke_access_json"), knownvalue.StringExact(createCfg["revoke_access_json"])),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("delete_account_json"), knownvalue.StringExact(createCfg["delete_account_json"])),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("enable_account_json"), knownvalue.StringExact(createCfg["enable_account_json"])),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("disable_account_json"), knownvalue.StringExact(createCfg["disable_account_json"])),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("account_exists_json"), knownvalue.StringExact(createCfg["account_exists_json"])),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("update_user_json"), knownvalue.StringExact(createCfg["update_user_json"])),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("accounts_import"), knownvalue.StringExact(createCfg["accounts_import"])),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("entitlement_value_import"), knownvalue.StringExact(createCfg["entitlement_value_import"])),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("user_import"), knownvalue.StringExact(createCfg["user_import"])),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("create_entitlement_json"), knownvalue.StringExact(createCfg["create_entitlement_json"])),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("delete_entitlement_json"), knownvalue.StringExact(createCfg["delete_entitlement_json"])),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("entitlement_exist_json"), knownvalue.StringExact(createCfg["entitlement_exist_json"])),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("update_entitlement_json"), knownvalue.StringExact(createCfg["update_entitlement_json"])),
				},
			},
			// Import
			{
				ResourceName:            resourceName,
				ImportStateId:           createCfg["connection_name"],
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"msg", "error_code", "password", "change_pass_json", "username"},
			},
			// Update
			{
				Config: testAccDBConnectionResourceConfig25B(filePath, "update"),
				ConfigStateChecks: []statecheck.StateCheck{
					//Encrypted Connection Attributes are removed
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("connection_name"), knownvalue.StringExact(updateCfg["connection_name"])),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("url"), knownvalue.StringExact(updateCfg["url"])),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("driver_name"), knownvalue.StringExact(updateCfg["driver_name"])),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("create_account_json"), knownvalue.StringExact(updateCfg["create_account_json"])),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("update_account_json"), knownvalue.StringExact(updateCfg["update_account_json"])),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("grant_access_json"), knownvalue.StringExact(updateCfg["grant_access_json"])),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("revoke_access_json"), knownvalue.StringExact(updateCfg["revoke_access_json"])),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("delete_account_json"), knownvalue.StringExact(updateCfg["delete_account_json"])),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("enable_account_json"), knownvalue.StringExact(updateCfg["enable_account_json"])),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("disable_account_json"), knownvalue.StringExact(updateCfg["disable_account_json"])),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("account_exists_json"), knownvalue.StringExact(updateCfg["account_exists_json"])),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("update_user_json"), knownvalue.StringExact(updateCfg["update_user_json"])),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("accounts_import"), knownvalue.StringExact(updateCfg["accounts_import"])),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("entitlement_value_import"), knownvalue.StringExact(updateCfg["entitlement_value_import"])),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("user_import"), knownvalue.StringExact(updateCfg["user_import"])),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("create_entitlement_json"), knownvalue.StringExact(createCfg["create_entitlement_json"])),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("delete_entitlement_json"), knownvalue.StringExact(createCfg["delete_entitlement_json"])),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("entitlement_exist_json"), knownvalue.StringExact(createCfg["entitlement_exist_json"])),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("update_entitlement_json"), knownvalue.StringExact(createCfg["update_entitlement_json"])),
				},
			},
			// Update the Connectionname to a new value
			{
				Config:      testAccDBConnectionResourceConfig25B(filePath, "update_connection_name"),
				ExpectError: regexp.MustCompile(`Connection name cannot be updated`),
			},
		},
	})
}

func testAccDBConnectionResourceConfig25B(jsonPath, operation string) string {
	return fmt.Sprintf(`
provider "saviynt" {
  server_url = "%s"
  username   = "%s"
  password   = "%s"
}

locals {
  cfg = jsondecode(file("%s"))["%s"]
}

resource "saviynt_db_connection_resource" "db" {
  connection_name           = local.cfg.connection_name
  url                       = local.cfg.url
  username                  = local.cfg.username
  password                  = local.cfg.password
  driver_name               = local.cfg.driver_name

  create_account_json       = jsonencode(local.cfg.create_account_json)
  update_account_json       = jsonencode(local.cfg.update_account_json)
  grant_access_json         = jsonencode(local.cfg.grant_access_json)
  revoke_access_json        = jsonencode(local.cfg.revoke_access_json)
  change_pass_json          = jsonencode(local.cfg.change_pass_json)
  delete_account_json       = jsonencode(local.cfg.delete_account_json)
  enable_account_json       = jsonencode(local.cfg.enable_account_json)
  disable_account_json      = jsonencode(local.cfg.disable_account_json)
  account_exists_json       = jsonencode(local.cfg.account_exists_json)
  update_user_json          = jsonencode(local.cfg.update_user_json)
  create_entitlement_json   =jsonencode(local.cfg.create_entitlement_json)
  delete_entitlement_json   =jsonencode(local.cfg.delete_entitlement_json)
  entitlement_exist_json    =jsonencode(local.cfg.entitlement_exist_json)
  update_entitlement_json   =jsonencode(local.cfg.update_entitlement_json)
  accounts_import           = local.cfg.accounts_import
  entitlement_value_import  = local.cfg.entitlement_value_import
  user_import               = local.cfg.user_import

}
`, os.Getenv("SAVIYNT_URL_25B"),
		os.Getenv("SAVIYNT_USERNAME_25B"),
		os.Getenv("SAVIYNT_PASSWORD_25B"),
		jsonPath,
		operation,
	)
}
