// Copyright (c) Saviynt Inc.
// SPDX-License-Identifier: MPL-2.0

// entra_id_connection_data_source_test.go contains the Terraform acceptance test suite for
// the `saviynt_entraid_connection_datasource` data source. It verifies that once an EntraId connection
// resource is created, the data source can successfully look it up and populate all expected
// attributes from Saviynt.
//
// The test flow is as follows:
//  1. Load test fixture data from `entra_id_connection_resource_test_data.json`.
//  2. Apply a resource configuration to create a `saviynt_entraid_connection_resource`.
//  3. Use `data.saviynt_entraid_connection_datasource.test` to read back the same connector by name.
//  4. Assert that each field in the data source matches the values provided during creation.
//
// Note: Environment variables `SAVIYNT_URL`, `SAVIYNT_USERNAME`, and `SAVIYNT_PASSWORD` must be set
//
//	to point at a live Saviynt instance for the acceptance test to run successfully.
package provider

import (
	"fmt"
	"os"
	"terraform-provider-Saviynt/util"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

func TestAccSaviyntEntraIdConnectionDataSource(t *testing.T) {
	filePath := "entra_id_connection_resource_test_data.json"
	createCfg := util.LoadConnectorData(t, filePath, "create")
	datasource := "data.saviynt_entraid_connection_datasource.test"

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccEntraIdConnectionDataSourceConfig(),
				Check: resource.ComposeTestCheckFunc(
					func(s *terraform.State) error {
						res, ok := s.RootModule().Resources[datasource]
						if !ok {
							t.Fatalf("Resource %s not found in state", datasource)
						}
						t.Logf("Full data source attributes:\n%+v", res.Primary.Attributes)
						return nil
					},
					// Now assert values
					resource.TestCheckResourceAttr(datasource, "msg", "success"),
					resource.TestCheckResourceAttr(datasource, "error_code", "0"),
					resource.TestCheckResourceAttr(datasource, "connection_name", createCfg["connection_name"]),
					resource.TestCheckResourceAttr(datasource, "connection_type", createCfg["connection_type"]),
					resource.TestCheckResourceAttr(datasource, "connection_attributes.client_id", createCfg["client_id"]),
					resource.TestCheckResourceAttr(datasource, "connection_attributes.aad_tenant_id", createCfg["aad_tenant_id"]),
					resource.TestCheckResourceAttr(datasource, "connection_attributes.authentication_endpoint", createCfg["authentication_endpoint"]),
					resource.TestCheckResourceAttr(datasource, "connection_attributes.microsoft_graph_endpoint", createCfg["microsoft_graph_endpoint"]),
					resource.TestCheckResourceAttr(datasource, "connection_attributes.azure_management_endpoint", createCfg["azure_management_endpoint"]),
					resource.TestCheckResourceAttr(datasource, "connection_attributes.account_attributes", createCfg["account_attributes"]),
					resource.TestCheckResourceAttr(datasource, "connection_attributes.account_import_fields", createCfg["account_import_fields"]),
					resource.TestCheckResourceAttr(datasource, "connection_attributes.status_threshold_config", createCfg["status_threshold_config"]),
					resource.TestCheckResourceAttr(datasource, "connection_attributes.endpoints_filter", createCfg["endpoints_filter"]),
					resource.TestCheckResourceAttr(datasource, "connection_attributes.import_user_json", createCfg["import_user_json"]),
					resource.TestCheckResourceAttr(datasource, "connection_attributes.create_account_json", createCfg["create_account_json"]),
					resource.TestCheckResourceAttr(datasource, "connection_attributes.update_account_json", createCfg["update_account_json"]),
					resource.TestCheckResourceAttr(datasource, "connection_attributes.enable_account_json", createCfg["enable_account_json"]),
					resource.TestCheckResourceAttr(datasource, "connection_attributes.disable_account_json", createCfg["disable_account_json"]),
					resource.TestCheckResourceAttr(datasource, "connection_attributes.add_access_json", createCfg["add_access_json"]),
					resource.TestCheckResourceAttr(datasource, "connection_attributes.remove_access_json", createCfg["remove_access_json"]),
					resource.TestCheckResourceAttr(datasource, "connection_attributes.remove_account_json", createCfg["remove_account_json"]),
				),
			},
		},
	})
}

func testAccEntraIdConnectionDataSourceConfig() string {
	jsonPath := "${filepath}/entra_id_connection_resource_test_data.json"
	return fmt.Sprintf(`
provider "saviynt" {
  server_url = "%s"
  username   = "%s"
  password   = "%s"
}

locals {
  cfg = jsondecode(file("%s"))["create"]
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
  
data "saviynt_entraid_connection_datasource" "test" {
	connection_name     = local.cfg.connection_name
	depends_on = [saviynt_entraid_connection_resource.entraid]
}
`, os.Getenv("SAVIYNT_URL"),
		os.Getenv("SAVIYNT_USERNAME"),
		os.Getenv("SAVIYNT_PASSWORD"),
		jsonPath,
	)
}
