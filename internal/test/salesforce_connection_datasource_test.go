// Copyright (c) Saviynt Inc.
// SPDX-License-Identifier: MPL-2.0

// salesforce_connection_data_source_test.go contains the Terraform acceptance test suite for
// the `saviynt_salesforce_connection_datasource` data source. It verifies that once an Salesforce connection
// resource is created, the data source can successfully look it up and populate all expected
// attributes from Saviynt.
//
// The test flow is as follows:
//  1. Load test fixture data from `salesforce_connection_resource_test_data.json`.
//  2. Apply a resource configuration to create a `saviynt_salesforce_connection_resource`.
//  3. Use `data.saviynt_salesforce_connection_datasource.test` to read back the same connector by name.
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

func TestAccSaviyntSalesforceConnectionDataSource(t *testing.T) {
	filePath := testutil.GetTestDataPath(t, "./test_data/salesforce_connection_test_data.json")
	filePath = testutil.PrepareTestDataWithEnv(t, filePath)
	createCfg := testutil.LoadConnectorData(t, filePath, "ds")
	datasource := "data.saviynt_salesforce_connection_datasource.test"

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { TestAccPreCheck(t) },
		ProtoV6ProviderFactories: TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccSalesforceConnectionDataSourceConfig(filePath),
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
					resource.TestCheckResourceAttr(datasource, "connection_attributes.redirect_uri", createCfg["redirect_uri"]),
					resource.TestCheckResourceAttr(datasource, "connection_attributes.instance_url", createCfg["instance_url"]),
					resource.TestCheckResourceAttr(datasource, "connection_attributes.object_to_be_imported", createCfg["object_to_be_imported"]),
					resource.TestCheckResourceAttr(datasource, "connection_attributes.createaccountjson", createCfg["createaccountjson"]),
					resource.TestCheckResourceAttr(datasource, "connection_attributes.status_threshold_config", createCfg["status_threshold_config"]),
				),
			},
		},
	})
}

func testAccSalesforceConnectionDataSourceConfig(jsonPath string) string {
	return fmt.Sprintf(`
provider "saviynt" {
  server_url = "%s"
  username   = "%s"
  password   = "%s"
}

locals {
  cfg = jsondecode(file("%s"))["ds"]
}

resource "saviynt_salesforce_connection_resource" "salesforce" {
  connection_type    = local.cfg.connection_type
  connection_name    = local.cfg.connection_name
  client_id             = local.cfg.client_id
  redirect_uri          = local.cfg.redirect_uri
  instance_url          = local.cfg.instance_url
  object_to_be_imported = local.cfg.object_to_be_imported
  createaccountjson = jsonencode(local.cfg.createaccountjson)
  status_threshold_config=jsonencode(local.cfg.status_threshold_config)
}
  
data "saviynt_salesforce_connection_datasource" "test" {
	connection_name     = local.cfg.connection_name
	authenticate 		= true
	depends_on = [saviynt_salesforce_connection_resource.salesforce]
}
`, os.Getenv("SAVIYNT_URL"),
		os.Getenv("SAVIYNT_USERNAME"),
		os.Getenv("SAVIYNT_PASSWORD"),
		jsonPath,
	)
}
