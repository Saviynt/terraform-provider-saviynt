/*
 * Copyright (c) 2025 Saviynt Inc.
 * All Rights Reserved.
 *
 * This software is the confidential and proprietary information of
 * Saviynt Inc. ("Confidential Information"). You shall not disclose,
 * use, or distribute such Confidential Information except in accordance
 * with the terms of the license agreement you entered into with Saviynt.
 *
 * SAVIYNT MAKES NO REPRESENTATIONS OR WARRANTIES ABOUT THE SUITABILITY OF
 * THE SOFTWARE, EITHER EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO
 * THE IMPLIED WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR
 * PURPOSE, OR NON-INFRINGEMENT.
 */

// endpoint_data_source_test.go contains the Terraform acceptance test suite for
// the `saviynt_endpoints_datasource` data source. It verifies that once an endpoint
// resource is created, the data source can successfully look it up and populate all expected
// attributes from Saviynt.
//
// The test flow is as follows:
//  1. Load test fixture data from `endpoint_resource_test_data.json`.
//  2. Apply a resource configuration to create a `saviynt_endpoint_resource`.
//  3. Use `data.saviynt_endpoints_datasource.test` to read back the same endpoint by name.
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

func TestAccEndpointDataSource(t *testing.T) {
	filePath := testutil.GetTestDataPath(t, "./test_data/endpoint_resource_test_data.json")
	filePath = testutil.PrepareTestDataWithEnv(t, filePath)
	createCfg := testutil.LoadConnectorData(t, filePath, "ds")
	datasource := "data.saviynt_endpoints_datasource.test"

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { TestAccPreCheck(t) },
		ProtoV6ProviderFactories: TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccSaviyntEndpointDataSourceConfig(filePath),
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
					resource.TestCheckResourceAttr(datasource, "message", "Success"),
					resource.TestCheckResourceAttr(datasource, "error_code", "0"),
					resource.TestCheckResourceAttr(datasource, "results.0.endpointname", createCfg["endpoint_name"]),
					resource.TestCheckResourceAttr(datasource, "results.0.display_name", createCfg["display_name"]),
					resource.TestCheckResourceAttr(datasource, "results.0.securitysystem", createCfg["security_system"]),
					resource.TestCheckResourceAttr(datasource, "results.0.custom_property1", createCfg["custom_property1"]),
					resource.TestCheckResourceAttr(datasource, "results.0.account_custom_property_1_label", createCfg["account_custom_property_1_label"]),
					resource.TestCheckResourceAttr(datasource, "results.0.custom_property31_label", createCfg["custom_property31_label"]),
					resource.TestCheckResourceAttr(datasource, "results.0.enable_copy_access", createCfg["enable_copy_access"]),
				),
			},
		},
	})
}

func testAccSaviyntEndpointDataSourceConfig(jsonPath string) string {
	return fmt.Sprintf(`
provider "saviynt" {
  server_url = "%s"
  username   = "%s"
  password   = "%s"
}

locals {
  cfg = jsondecode(file("%s"))["ds"]
}

resource "saviynt_endpoint_resource" "e" {
  endpoint_name                      = local.cfg.endpoint_name
  display_name                       = local.cfg.display_name
  security_system                    = local.cfg.security_system
  custom_property1            		 = local.cfg.custom_property1
  account_custom_property_1_label    = local.cfg.account_custom_property_1_label
  custom_property31_label 			 = local.cfg.custom_property31_label
  enable_copy_access				 = local.cfg.enable_copy_access
}
  
data "saviynt_endpoints_datasource" "test" {
	endpointname = local.cfg.endpoint_name
	depends_on = [saviynt_endpoint_resource.e]
}
`, os.Getenv("SAVIYNT_URL"),
		os.Getenv("SAVIYNT_USERNAME"),
		os.Getenv("SAVIYNT_PASSWORD"),
		jsonPath,
	)
}
