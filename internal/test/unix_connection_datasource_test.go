// Copyright (c) 2025 Saviynt Inc.
// SPDX-License-Identifier: MPL-2.0

// unix_connection_data_source_test.go contains the Terraform acceptance test suite for
// the `saviynt_unix_connection_datasource` data source. It verifies that once an Unix connection
// resource is created, the data source can successfully look it up and populate all expected
// attributes from Saviynt.
//
// The test flow is as follows:
//  1. Load test fixture data from `unix_connection_resource_test_data.json`.
//  2. Apply a resource configuration to create a `saviynt_unix_connection_resource`.
//  3. Use `data.saviynt_unix_connection_datasource.test` to read back the same connector by name.
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

func TestAccSaviyntUnixConnectionDataSource(t *testing.T) {
	filePath := testutil.GetTestDataPath(t, "./test_data/unix_connection_test_data.json")
	filePath = testutil.PrepareTestDataWithEnv(t, filePath)
	createCfg := testutil.LoadConnectorData(t, filePath, "ds")
	datasource := "data.saviynt_unix_connection_datasource.test"

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { TestAccPreCheck(t) },
		ProtoV6ProviderFactories: TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccUnixConnectionDataSourceConfig(filePath),
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
					resource.TestCheckResourceAttr(datasource, "connection_attributes.host_name", createCfg["host_name"]),
					resource.TestCheckResourceAttr(datasource, "connection_attributes.port_number", createCfg["port_number"]),
					resource.TestCheckResourceAttr(datasource, "connection_attributes.groups_file", createCfg["groups_file"]),
					resource.TestCheckResourceAttr(datasource, "connection_attributes.status_threshold_config", createCfg["status_threshold_config"]),
				),
			},
		},
	})
}

func testAccUnixConnectionDataSourceConfig(jsonPath string) string {
	return fmt.Sprintf(`
provider "saviynt" {
  server_url = "%s"
  username   = "%s"
  password   = "%s"
}

locals {
  cfg = jsondecode(file("%s"))["ds"]
}

resource "saviynt_unix_connection_resource" "unix" {
  connection_name    				 = local.cfg.connection_name
  host_name       = local.cfg.host_name
  port_number     = local.cfg.port_number
  username        = local.cfg.username
  password        = local.cfg.password
  groups_file   = local.cfg.groups_file
  status_threshold_config = jsonencode(local.cfg.status_threshold_config)
  ssh_key = local.cfg.ssh_key
}
  
data "saviynt_unix_connection_datasource" "test" {
	connection_name     = local.cfg.connection_name
	authenticate 		= true
	depends_on = [saviynt_unix_connection_resource.unix]
}
`, os.Getenv("SAVIYNT_URL"),
		os.Getenv("SAVIYNT_USERNAME"),
		os.Getenv("SAVIYNT_PASSWORD"),
		jsonPath,
	)
}
