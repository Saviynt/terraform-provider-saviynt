// Copyright (c) Saviynt Inc.
// SPDX-License-Identifier: MPL-2.0

// adsi_connection_resource_test.go contains the Terraform acceptance test suite for
// the `saviynt_adsi_connection_resource`. It validates the resource’s full lifecycle against
// a live Saviynt instance by exercising:
//
//   • Create: provisions a new ADSI connection and asserts each attribute via JSONPath checks.
//   • Import: verifies the resource can be imported by its connection name.
//   • Update: applies configuration changes and confirms the updated attribute values.
//   • Negative Cases: ensures updates to `connection_name` and `connection_type` are rejected.
//
// Test data is loaded from `adsi_connection_resource_test_data.json` using `util.LoadConnectorData`.
// Environment variables `SAVIYNT_URL`, `SAVIYNT_USERNAME`, and `SAVIYNT_PASSWORD` must be set
// to point at a valid Saviynt Security Manager before running these tests.
package provider

import (
	"fmt"
	"os"
	"regexp"
	"terraform-provider-Saviynt/util"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
)

func TestAccSaviyntADSIConnectionResource(t *testing.T) {
	filePath := "adsi_connection_resource_test_data.json"
	createCfg := util.LoadConnectorData(t, filePath, "create")
	updateCfg := util.LoadConnectorData(t, filePath, "update")
	resourceName := "saviynt_adsi_connection_resource.adsi"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create
			{
				Config: testAccADSIConnectionResourceConfig("create"),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("connection_name"), knownvalue.StringExact(createCfg["connection_name"])),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("connection_type"), knownvalue.StringExact(createCfg["connection_type"])),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("url"), knownvalue.StringExact(createCfg["url"])),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("password"), knownvalue.StringExact(createCfg["password"])),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("username"), knownvalue.StringExact(createCfg["username"])),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("connection_url"), knownvalue.StringExact(createCfg["connection_url"])),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("provisioning_url"), knownvalue.StringExact(createCfg["provisioning_url"])),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("forestlist"), knownvalue.StringExact(createCfg["forestlist"])),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("updateaccountjson"), knownvalue.StringExact(createCfg["updateaccountjson"])),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("addaccessjson"), knownvalue.StringExact(createCfg["addaccessjson"])),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("searchfilter"), knownvalue.StringExact(createCfg["searchfilter"])),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("objectfilter"), knownvalue.StringExact(createCfg["objectfilter"])),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("account_attribute"), knownvalue.StringExact(createCfg["account_attribute"])),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("status_threshold_config"), knownvalue.StringExact(createCfg["status_threshold_config"])),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("entitlement_attribute"), knownvalue.StringExact(createCfg["entitlement_attribute"])),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("user_attribute"), knownvalue.StringExact(createCfg["user_attribute"])),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("group_search_base_dn"), knownvalue.StringExact(createCfg["group_search_base_dn"])),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("checkforunique"), knownvalue.StringExact(createCfg["checkforunique"])),
				},
			},
			// Import
			{
				ResourceName:            resourceName,
				ImportStateId:           createCfg["connection_name"],
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"msg", "error_code", "password"},
			},
			// Update
			{
				Config: testAccADSIConnectionResourceConfig("update"),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("connection_name"), knownvalue.StringExact(updateCfg["connection_name"])),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("connection_type"), knownvalue.StringExact(updateCfg["connection_type"])),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("url"), knownvalue.StringExact(updateCfg["url"])),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("password"), knownvalue.StringExact(updateCfg["password"])),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("username"), knownvalue.StringExact(updateCfg["username"])),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("connection_url"), knownvalue.StringExact(updateCfg["connection_url"])),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("provisioning_url"), knownvalue.StringExact(updateCfg["provisioning_url"])),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("forestlist"), knownvalue.StringExact(updateCfg["forestlist"])),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("updateaccountjson"), knownvalue.StringExact(updateCfg["updateaccountjson"])),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("addaccessjson"), knownvalue.StringExact(updateCfg["addaccessjson"])),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("searchfilter"), knownvalue.StringExact(updateCfg["searchfilter"])),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("objectfilter"), knownvalue.StringExact(updateCfg["objectfilter"])),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("account_attribute"), knownvalue.StringExact(updateCfg["account_attribute"])),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("status_threshold_config"), knownvalue.StringExact(updateCfg["status_threshold_config"])),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("entitlement_attribute"), knownvalue.StringExact(updateCfg["entitlement_attribute"])),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("user_attribute"), knownvalue.StringExact(updateCfg["user_attribute"])),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("group_search_base_dn"), knownvalue.StringExact(updateCfg["group_search_base_dn"])),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("checkforunique"), knownvalue.StringExact(updateCfg["checkforunique"])),
				},
			},
			// Update the Connectionname to a new value
			{
				Config:      testAccADSIConnectionResourceConfig("update_connection_name"),
				ExpectError: regexp.MustCompile(`Connection name cannot be updated`),
			},
			// Update the Connectiontype to a new value
			{
				Config:      testAccADSIConnectionResourceConfig("update_connection_type"),
				ExpectError: regexp.MustCompile(`Connection type cannot by updated`),
			},
		},
	})
}

func testAccADSIConnectionResourceConfig(operation string) string {
	jsonPath := "${filepath}/adsi_connection_resource_test_data.json"
	return fmt.Sprintf(`
provider "saviynt" {
  server_url = "%s"
  username   = "%s"
  password   = "%s"
}

locals {
  cfg = jsondecode(file("%s"))["%s"]
}

resource "saviynt_adsi_connection_resource" "adsi" {
  connection_type              = local.cfg.connection_type
  connection_name              = local.cfg.connection_name
  url                          = local.cfg.url
  password                     = local.cfg.password
  username                     = local.cfg.username
  connection_url               = local.cfg.connection_url
  provisioning_url             = local.cfg.provisioning_url
  forestlist                   = local.cfg.forestlist
  searchfilter                 = local.cfg.searchfilter
  group_search_base_dn         = local.cfg.group_search_base_dn
  objectfilter                 = local.cfg.objectfilter
  account_attribute            = local.cfg.account_attribute
  entitlement_attribute        = local.cfg.entitlement_attribute
  user_attribute               = local.cfg.user_attribute
  page_size                    = tostring(local.cfg.page_size)
  import_nested_membership     = tostring(local.cfg.import_nested_membership)
  statuskeyjson                = jsonencode(local.cfg.statuskeyjson)
  status_threshold_config      = jsonencode(local.cfg.status_threshold_config)
  checkforunique               = jsonencode(local.cfg.checkforunique)
  group_import_mapping         = jsonencode(local.cfg.group_import_mapping)
  createaccountjson          = jsonencode(local.cfg.createaccountjson)
  updateaccountjson          = jsonencode(local.cfg.updateaccountjson)
  enableaccountjson          = jsonencode(local.cfg.enableaccountjson)
  disableaccountjson         = jsonencode(local.cfg.disableaccountjson)
  removeaccountjson          = jsonencode(local.cfg.removeaccountjson)
  addaccessjson                = jsonencode(local.cfg.addaccessjson)
  removeaccessjson             = jsonencode(local.cfg.removeaccessjson)
  resetandchangepasswrdjson    = jsonencode(local.cfg.resetandchangepasswrdjson)
  creategroupjson            = jsonencode(local.cfg.creategroupjson)
  updategroupjson            = jsonencode(local.cfg.updategroupjson)
  removegroupjson            = jsonencode(local.cfg.removegroupjson)
  addaccessentitlementjson     = jsonencode(local.cfg.addaccessentitlementjson)
  removeaccessentitlementjson  = jsonencode(local.cfg.removeaccessentitlementjson)

}
`, os.Getenv("SAVIYNT_URL"),
		os.Getenv("SAVIYNT_USERNAME"),
		os.Getenv("SAVIYNT_PASSWORD"),
		jsonPath,
		operation,
	)
}
