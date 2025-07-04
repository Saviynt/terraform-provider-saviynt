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

// adsi_connection_data_source_test.go contains the Terraform acceptance test suite for
// the `saviynt_adsi_connection_datasource` data source. It verifies that once an ADSI connection
// resource is created, the data source can successfully look it up and populate all expected
// attributes from Saviynt.
//
// The test flow is as follows:
//  1. Load test fixture data from `adsi_connection_resource_test_data.json`.
//  2. Apply a resource configuration to create a `saviynt_adsi_connection_resource`.
//  3. Use `data.saviynt_adsi_connection_datasource.test` to read back the same connector by name.
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

func TestAccSaviyntADSIConnectionDataSource(t *testing.T) {
	filePath := testutil.GetTestDataPath(t, "./test_data/adsi_connection_test_data.json")
	filePath = testutil.PrepareTestDataWithEnv(t, filePath)
	createCfg := testutil.LoadConnectorData(t, filePath, "ds")
	datasource := "data.saviynt_adsi_connection_datasource.test"

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { TestAccPreCheck(t) },
		ProtoV6ProviderFactories: TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccADSIConnectionDataSourceConfig(filePath),
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
					resource.TestCheckResourceAttr(datasource, "connection_attributes.url", createCfg["url"]),
					resource.TestCheckResourceAttr(datasource, "connection_attributes.username", createCfg["username"]),
					resource.TestCheckResourceAttr(datasource, "connection_attributes.connection_url", createCfg["connection_url"]),
					resource.TestCheckResourceAttr(datasource, "connection_attributes.provisioning_url", createCfg["provisioning_url"]),
					resource.TestCheckResourceAttr(datasource, "connection_attributes.forest_list", createCfg["forestlist"]),
					resource.TestCheckResourceAttr(datasource, "connection_attributes.update_account_json", createCfg["updateaccountjson"]),
					resource.TestCheckResourceAttr(datasource, "connection_attributes.add_access_json", createCfg["addaccessjson"]),
					resource.TestCheckResourceAttr(datasource, "connection_attributes.search_filter", createCfg["searchfilter"]),
					resource.TestCheckResourceAttr(datasource, "connection_attributes.object_filter", createCfg["objectfilter"]),
					resource.TestCheckResourceAttr(datasource, "connection_attributes.account_attribute", createCfg["account_attribute"]),
					resource.TestCheckResourceAttr(datasource, "connection_attributes.status_threshold_config", createCfg["status_threshold_config"]),
				),
			},
		},
	})
}

func testAccADSIConnectionDataSourceConfig(jsonPath string) string {
	return fmt.Sprintf(`
provider "saviynt" {
  server_url = "%s"
  username   = "%s"
  password   = "%s"
}

locals {
  cfg = jsondecode(file("%s"))["ds"]
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
  
data "saviynt_adsi_connection_datasource" "test" {
	connection_name     = local.cfg.connection_name
	depends_on = [saviynt_adsi_connection_resource.adsi]
}
`, os.Getenv("SAVIYNT_URL"),
		os.Getenv("SAVIYNT_USERNAME"),
		os.Getenv("SAVIYNT_PASSWORD"),
		jsonPath,
	)
}
