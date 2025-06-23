// Copyright (c) Saviynt Inc.
// SPDX-License-Identifier: MPL-2.0

// github_rest_connection_data_source_test.go contains the Terraform acceptance test suite for
// the `saviynt_github_rest_connection_datasource` data source. It verifies that once an GithubRest connection
// resource is created, the data source can successfully look it up and populate all expected
// attributes from Saviynt.
//
// The test flow is as follows:
//  1. Load test fixture data from `github_rest_connection_resource_test_data.json`.
//  2. Apply a resource configuration to create a `saviynt_github_rest_connection_resource`.
//  3. Use `data.saviynt_github_rest_connection_datasource.test` to read back the same connector by name.
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

func TestAccSaviyntGithubRestConnectionDataSource(t *testing.T) {
	filePath := "github_rest_connection_test_data.json"
	createCfg := util.LoadConnectorData(t, filePath, "create")
	datasource := "data.saviynt_github_rest_connection_datasource.test"

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccGithubRestConnectionDataSourceConfig(),
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
					resource.TestCheckResourceAttr(datasource, "connection_attributes.import_account_ent_json", createCfg["import_account_ent_json"]),
					resource.TestCheckResourceAttr(datasource, "connection_attributes.organization_list", createCfg["organization_list"]),
				),
			},
		},
	})
}

func testAccGithubRestConnectionDataSourceConfig() string {
	jsonPath := "${filepath}/github_rest_connection_test_data.json"
	return fmt.Sprintf(`
provider "saviynt" {
  server_url = "%s"
  username   = "%s"
  password   = "%s"
}

locals {
  cfg = jsondecode(file("%s"))["create"]
}

resource "saviynt_github_rest_connection_resource" "github_rest" {
  connection_type    				 = local.cfg.connection_type
  connection_name   				 = local.cfg.connection_name
  connection_json                    = jsonencode(local.cfg.connection_json)
  import_account_ent_json			 = jsonencode(local.cfg.import_account_ent_json)
  organization_list                  = local.cfg.organization_list
}
  
data "saviynt_github_rest_connection_datasource" "test" {
	connection_name     = local.cfg.connection_name
	depends_on = [saviynt_github_rest_connection_resource.github_rest]
}
`, os.Getenv("SAVIYNT_URL"),
		os.Getenv("SAVIYNT_USERNAME"),
		os.Getenv("SAVIYNT_PASSWORD"),
		jsonPath,
	)
}
