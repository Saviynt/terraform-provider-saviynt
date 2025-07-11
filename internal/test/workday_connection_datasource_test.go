// Copyright (c) Saviynt Inc.
// SPDX-License-Identifier: MPL-2.0

// workday_connection_data_source_test.go contains the Terraform acceptance test suite for
// the `saviynt_workday_connection_datasource` data source. It verifies that once an Workday connection
// resource is created, the data source can successfully look it up and populate all expected
// attributes from Saviynt.
//
// The test flow is as follows:
//  1. Load test fixture data from `workday_connection_resource_test_data.json`.
//  2. Apply a resource configuration to create a `saviynt_workday_connection_resource`.
//  3. Use `data.saviynt_workday_connection_datasource.test` to read back the same connector by name.
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

func TestAccSaviyntWorkdayConnectionDataSource(t *testing.T) {
	filePath := testutil.GetTestDataPath(t, "./test_data/workday_connection_test_data.json")
	filePath = testutil.PrepareTestDataWithEnv(t, filePath)
	createCfg := testutil.LoadConnectorData(t, filePath, "ds")
	datasource := "data.saviynt_workday_connection_datasource.test"

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { TestAccPreCheck(t) },
		ProtoV6ProviderFactories: TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccWorkdayConnectionDataSourceConfig(filePath),
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
					resource.TestCheckResourceAttr(datasource, "connection_attributes.base_url", createCfg["base_url"]),
					resource.TestCheckResourceAttr(datasource, "connection_attributes.api_version", createCfg["api_version"]),
					resource.TestCheckResourceAttr(datasource, "connection_attributes.use_oauth", createCfg["use_oauth"]),
					resource.TestCheckResourceAttr(datasource, "connection_attributes.access_import_list", createCfg["access_import_list"]),
					resource.TestCheckResourceAttr(datasource, "connection_attributes.status_key_json", createCfg["status_key_json"]),
					resource.TestCheckResourceAttr(datasource, "connection_attributes.user_import_payload", createCfg["user_import_payload"]),
					resource.TestCheckResourceAttr(datasource, "connection_attributes.user_import_mapping", createCfg["user_import_mapping"]),
				),
			},
		},
	})
}

func testAccWorkdayConnectionDataSourceConfig(jsonPath string) string {
	return fmt.Sprintf(`
provider "saviynt" {
  server_url = "%s"
  username   = "%s"
  password   = "%s"
}

locals {
  cfg = jsondecode(file("%s"))["ds"]
}

resource "saviynt_workday_connection_resource" "workday" {
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
}
  
data "saviynt_workday_connection_datasource" "test" {
	connection_name     = local.cfg.connection_name
	authenticate 		= true
	depends_on = [saviynt_workday_connection_resource.workday]
}
`, os.Getenv("SAVIYNT_URL"),
		os.Getenv("SAVIYNT_USERNAME"),
		os.Getenv("SAVIYNT_PASSWORD"),
		jsonPath,
	)
}
