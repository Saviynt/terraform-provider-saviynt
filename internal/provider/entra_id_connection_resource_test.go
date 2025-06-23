// Copyright (c) Saviynt Inc.
// SPDX-License-Identifier: MPL-2.0

// entra_id_connection_resource_test.go contains the Terraform acceptance test suite for
// the `saviynt_entraid_connection_resource`. It validates the resource’s full lifecycle against
// a live Saviynt instance by exercising:
//
//   - Create: provisions a new EntraId connection and asserts each attribute via JSONPath checks.
//   - Import: verifies the resource can be imported by its connection name.
//   - Update: applies configuration changes and confirms the updated attribute values.
//   - Negative Cases: ensures updates to `connection_name` and `connection_type` are rejected.
//
// Test data is loaded from `entra_id_connection_resource_test_data.json` using `util.LoadConnectorData`.
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

func TestAccSaviyntEntraIdConnectionResource(t *testing.T) {
	filePath := "entra_id_connection_resource_test_data.json"
	createCfg := util.LoadConnectorData(t, filePath, "create")
	updateCfg := util.LoadConnectorData(t, filePath, "update")
	resourceName := "saviynt_entraid_connection_resource.entraid"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create
			{
				Config: testAccEntraIdConnectionResourceConfig("create"),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("connection_type"), knownvalue.StringExact(createCfg["connection_type"])),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("connection_name"), knownvalue.StringExact(createCfg["connection_name"])),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("client_id"), knownvalue.StringExact(createCfg["client_id"])),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("client_secret"), knownvalue.StringExact(createCfg["client_secret"])),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("aad_tenant_id"), knownvalue.StringExact(createCfg["aad_tenant_id"])),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("authentication_endpoint"), knownvalue.StringExact(createCfg["authentication_endpoint"])),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("microsoft_graph_endpoint"), knownvalue.StringExact(createCfg["microsoft_graph_endpoint"])),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("azure_management_endpoint"), knownvalue.StringExact(createCfg["azure_management_endpoint"])),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("account_attributes"), knownvalue.StringExact(createCfg["account_attributes"])),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("account_import_fields"), knownvalue.StringExact(createCfg["account_import_fields"])),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("status_threshold_config"), knownvalue.StringExact(createCfg["status_threshold_config"])),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("endpoints_filter"), knownvalue.StringExact(createCfg["endpoints_filter"])),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("connection_json"), knownvalue.StringExact(createCfg["connection_json"])),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("import_user_json"), knownvalue.StringExact(createCfg["import_user_json"])),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("create_account_json"), knownvalue.StringExact(createCfg["create_account_json"])),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("update_account_json"), knownvalue.StringExact(createCfg["update_account_json"])),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("enable_account_json"), knownvalue.StringExact(createCfg["enable_account_json"])),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("disable_account_json"), knownvalue.StringExact(createCfg["disable_account_json"])),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("add_access_json"), knownvalue.StringExact(createCfg["add_access_json"])),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("remove_access_json"), knownvalue.StringExact(createCfg["remove_access_json"])),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("remove_account_json"), knownvalue.StringExact(createCfg["remove_account_json"])),
				},
			},
			// Import
			{
				ResourceName:            resourceName,
				ImportStateId:           createCfg["connection_name"],
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"msg", "error_code", "client_secret", "windows_connector_json", "connection_json"},
			},
			// Update
			{
				Config: testAccEntraIdConnectionResourceConfig("update"),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("connection_type"), knownvalue.StringExact(updateCfg["connection_type"])),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("connection_name"), knownvalue.StringExact(updateCfg["connection_name"])),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("client_id"), knownvalue.StringExact(updateCfg["client_id"])),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("aad_tenant_id"), knownvalue.StringExact(updateCfg["aad_tenant_id"])),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("authentication_endpoint"), knownvalue.StringExact(updateCfg["authentication_endpoint"])),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("microsoft_graph_endpoint"), knownvalue.StringExact(updateCfg["microsoft_graph_endpoint"])),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("azure_management_endpoint"), knownvalue.StringExact(updateCfg["azure_management_endpoint"])),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("account_attributes"), knownvalue.StringExact(updateCfg["account_attributes"])),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("account_import_fields"), knownvalue.StringExact(updateCfg["account_import_fields"])),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("status_threshold_config"), knownvalue.StringExact(updateCfg["status_threshold_config"])),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("endpoints_filter"), knownvalue.StringExact(updateCfg["endpoints_filter"])),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("import_user_json"), knownvalue.StringExact(updateCfg["import_user_json"])),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("create_account_json"), knownvalue.StringExact(updateCfg["create_account_json"])),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("update_account_json"), knownvalue.StringExact(updateCfg["update_account_json"])),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("enable_account_json"), knownvalue.StringExact(updateCfg["enable_account_json"])),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("disable_account_json"), knownvalue.StringExact(updateCfg["disable_account_json"])),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("add_access_json"), knownvalue.StringExact(updateCfg["add_access_json"])),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("remove_access_json"), knownvalue.StringExact(updateCfg["remove_access_json"])),
					statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("remove_account_json"), knownvalue.StringExact(updateCfg["remove_account_json"])),
				},
			},
			// Update the Connectionname to a new value
			{
				Config:      testAccEntraIdConnectionResourceConfig("update_connection_name"),
				ExpectError: regexp.MustCompile(`Connection name cannot be updated`),
			},
			// Update the Connectiontype to a new value
			{
				Config:      testAccEntraIdConnectionResourceConfig("update_connection_type"),
				ExpectError: regexp.MustCompile(`Connection type cannot by updated`),
			},
		},
	})
}

func testAccEntraIdConnectionResourceConfig(operation string) string {
	jsonPath := "${filepath}/entra_id_connection_resource_test_data.json"
	return fmt.Sprintf(`
provider "saviynt" {
  server_url = "%s"
  username   = "%s"
  password   = "%s"
}

locals {
  cfg = jsondecode(file("%s"))["%s"]
}

resource "saviynt_entraid_connection_resource" "entraid" {
  connection_type           = local.cfg.connection_type
  connection_name           = local.cfg.connection_name
  client_id                 = local.cfg.client_id
  client_secret             = local.cfg.client_secret
  aad_tenant_id             = local.cfg.aad_tenant_id
  authentication_endpoint   = local.cfg.authentication_endpoint
  microsoft_graph_endpoint  = local.cfg.microsoft_graph_endpoint
  azure_management_endpoint = local.cfg.azure_management_endpoint

  account_attributes        = jsonencode(local.cfg.account_attributes)
  account_import_fields     = local.cfg.account_import_fields
  status_threshold_config   = jsonencode(local.cfg.status_threshold_config)
  endpoints_filter          = jsonencode(local.cfg.endpoints_filter)

  connection_json           = jsonencode(local.cfg.connection_json)
  import_user_json          = jsonencode(local.cfg.import_user_json)
  create_account_json       = jsonencode(local.cfg.create_account_json)
  update_account_json       = jsonencode(local.cfg.update_account_json)
  enable_account_json       = jsonencode(local.cfg.enable_account_json)
  disable_account_json      = jsonencode(local.cfg.disable_account_json)
  add_access_json           = jsonencode(local.cfg.add_access_json)
  remove_access_json        = jsonencode(local.cfg.remove_access_json)
  remove_account_json       = jsonencode(local.cfg.remove_account_json)
}
`, os.Getenv("SAVIYNT_URL"),
		os.Getenv("SAVIYNT_USERNAME"),
		os.Getenv("SAVIYNT_PASSWORD"),
		jsonPath,
		operation,
	)
}
