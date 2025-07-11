// Copyright (c) Saviynt Inc.
// SPDX-License-Identifier: MPL-2.0

// rest_connection_data_source_test.go contains the Terraform acceptance test suite for
// the `saviynt_rest_connection_datasource` data source. It verifies that once an Rest connection
// resource is created, the data source can successfully look it up and populate all expected
// attributes from Saviynt.
//
// The test flow is as follows:
//  1. Load test fixture data from `rest_connection_resource_test_data.json`.
//  2. Apply a resource configuration to create a `saviynt_rest_connection_resource`.
//  3. Use `data.saviynt_rest_connection_datasource.test` to read back the same connector by name.
//  4. Assert that each field in the data source matches the values provided during creation.
//
// Note: Environment variables `SAVIYNT_URL`, `SAVIYNT_USERNAME`, and `SAVIYNT_PASSWORD` must be set
//
//	to point at a live Saviynt instance for the acceptance test to run successfully.
package testing

import (
	"fmt"
	"os"
	"terraform-provider-Saviynt/util/testutil"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

func TestAccSaviyntRESTConnectionDataSource(t *testing.T) {
	filePath := testutil.GetTestDataPath(t, "./test_data/rest_connection_test_data.json")
	filePath = testutil.PrepareTestDataWithEnv(t, filePath)
	createCfg := testutil.LoadConnectorData(t, filePath, "ds")
	datasource := "data.saviynt_rest_connection_datasource.test"

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { TestAccPreCheck(t) },
		ProtoV6ProviderFactories: TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccRESTConnectionDataSourceConfig(filePath),
				Check: resource.ComposeTestCheckFunc(
					func(s *terraform.State) error {
						_, ok := s.RootModule().Resources[datasource]
						if !ok {
							t.Fatalf("Resource %s not found in state", datasource)
						}
						// t.Logf("Full data source attributes:\n%+v", res.Primary.Attributes)
						return nil
					},
					// Now assert values
					resource.TestCheckResourceAttr(datasource, "msg", "success"),
					resource.TestCheckResourceAttr(datasource, "error_code", "0"),
					resource.TestCheckResourceAttr(datasource, "connection_name", createCfg["connection_name"]),
					resource.TestCheckResourceAttr(datasource, "connection_type", createCfg["connection_type"]),
					resource.TestCheckResourceAttr(datasource, "connection_attributes.import_user_json", createCfg["import_user_json"]),
					resource.TestCheckResourceAttr(datasource, "connection_attributes.create_account_json", createCfg["create_account_json"]),
					resource.TestCheckResourceAttr(datasource, "connection_attributes.update_account_json", createCfg["update_account_json"]),
					resource.TestCheckResourceAttr(datasource, "connection_attributes.add_access_json", createCfg["add_access_json"]),
					resource.TestCheckResourceAttr(datasource, "connection_attributes.remove_access_json", createCfg["remove_access_json"]),
					resource.TestCheckResourceAttr(datasource, "connection_attributes.remove_account_json", createCfg["remove_account_json"]),
					resource.TestCheckResourceAttr(datasource, "connection_attributes.import_account_ent_json", createCfg["import_account_ent_json"]),
					resource.TestCheckResourceAttr(datasource, "connection_attributes.disable_account_json", createCfg["disable_account_json"]),
					resource.TestCheckResourceAttr(datasource, "connection_attributes.enable_account_json", createCfg["enable_account_json"]),
				),
			},
		},
	})
}

func testAccRESTConnectionDataSourceConfig(jsonPath string) string {
	return fmt.Sprintf(`
provider "saviynt" {
  server_url = "%s"
  username   = "%s"
  password   = "%s"
}

locals {
  cfg = jsondecode(file("%s"))["ds"]
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
  
data "saviynt_rest_connection_datasource" "test" {
	connection_name     = local.cfg.connection_name
	authenticate 		= true
	depends_on = [saviynt_rest_connection_resource.rest]
}
`, os.Getenv("SAVIYNT_URL"),
		os.Getenv("SAVIYNT_USERNAME"),
		os.Getenv("SAVIYNT_PASSWORD"),
		jsonPath,
	)
}
