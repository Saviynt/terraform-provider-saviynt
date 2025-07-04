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

// securitysystem_data_source_test.go contains the Terraform acceptance test suite for
// the `saviynt_security_systems_datasource` data source. It verifies that once an security system
// resource is created, the data source can successfully look it up and populate all expected
// attributes from Saviynt.
//
// The test flow is as follows:
//  1. Load test fixture data from `securitysystem_resource_test_data.json`.
//  2. Apply a resource configuration to create a `saviynt_security_system_resource`.
//  3. Use `data.saviynt_security_systems_datasource.test` to read back the same security system by name.
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

func TestAccSaviyntSecuritySystemDataSource(t *testing.T) {
	filePath := testutil.GetTestDataPath(t, "./test_data/securitysystem_resource_test_data.json")
	filePath = testutil.PrepareTestDataWithEnv(t, filePath)
	createCfg := testutil.LoadConnectorData(t, filePath, "ds")
	datasource := "data.saviynt_security_system_datasource.test"

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { TestAccPreCheck(t) },
		ProtoV6ProviderFactories: TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccSaviyntSecuritySystemDataSourceConfig(filePath),
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
					resource.TestCheckResourceAttr(datasource, "msg", "Success"),
					resource.TestCheckResourceAttr(datasource, "error_code", "0"),
					resource.TestCheckResourceAttr(datasource, "results.0.systemname1", createCfg["systemname"]),
					resource.TestCheckResourceAttr(datasource, "results.0.display_name", createCfg["display_name"]),
					resource.TestCheckResourceAttr(datasource, "results.0.access_add_workflow", createCfg["access_add_workflow"]),
					resource.TestCheckResourceAttr(datasource, "results.0.access_remove_workflow", createCfg["access_remove_workflow"]),
					resource.TestCheckResourceAttr(datasource, "results.0.add_service_account_workflow", createCfg["add_service_account_workflow"]),
					resource.TestCheckResourceAttr(datasource, "results.0.remove_service_account_workflow", createCfg["remove_service_account_workflow"]),
					resource.TestCheckResourceAttr(datasource, "results.0.automated_provisioning", createCfg["automated_provisioning"]),
				),
			},
		},
	})
}

func testAccSaviyntSecuritySystemDataSourceConfig(jsonPath string) string {
	return fmt.Sprintf(`
provider "saviynt" {
  server_url = "%s"
  username   = "%s"
  password   = "%s"
}

locals {
  cfg = jsondecode(file("%s"))["ds"]
}

resource "saviynt_security_system_resource" "ss" {
  systemname                         = local.cfg.systemname
  display_name                       = local.cfg.display_name
  access_add_workflow                = local.cfg.access_add_workflow
  access_remove_workflow             = local.cfg.access_remove_workflow
  add_service_account_workflow       = local.cfg.add_service_account_workflow
  remove_service_account_workflow    = local.cfg.remove_service_account_workflow
  automated_provisioning             = local.cfg.automated_provisioning
}
  
data "saviynt_security_system_datasource" "test" {
	systemname = local.cfg.systemname
	depends_on = [saviynt_security_system_resource.ss]
}
`, os.Getenv("SAVIYNT_URL"),
		os.Getenv("SAVIYNT_USERNAME"),
		os.Getenv("SAVIYNT_PASSWORD"),
		jsonPath,
	)
}
