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

// ad_connection_data_source_test.go contains the Terraform acceptance test suite for
// the `saviynt_ad_connection_datasource` data source. It verifies that once an AD connection
// resource is created, the data source can successfully look it up and populate all expected
// attributes from Saviynt.
//
// The test flow is as follows:
//  1. Load test fixture data from `ad_connection_resource_test_data.json`.
//  2. Apply a resource configuration to create a `saviynt_ad_connection_resource`.
//  3. Use `data.saviynt_ad_connection_datasource.test` to read back the same connector by name.
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

func TestAccSaviyntADConnectionDataSource(t *testing.T) {
	filePath := testutil.GetTestDataPath(t, "./test_data/ad_connection_test_data.json")
	filePath = testutil.PrepareTestDataWithEnv(t, filePath)
	createCfg := testutil.LoadConnectorData(t, filePath, "ds")
	// createCfg[ds][connection_name]=createCfg[ds][conne]
	datasource := "data.saviynt_ad_connection_datasource.test"

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { TestAccPreCheck(t) },
		ProtoV6ProviderFactories: TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccADConnectionDataSourceConfig(filePath),
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
					resource.TestCheckResourceAttr(datasource, "connection_attributes.search_filter", createCfg["searchfilter"]),
					resource.TestCheckResourceAttr(datasource, "connection_attributes.base", createCfg["base"]),
					resource.TestCheckResourceAttr(datasource, "connection_attributes.group_search_base_dn", createCfg["group_search_base_dn"]),
					resource.TestCheckResourceAttr(datasource, "connection_attributes.ldap_or_ad", createCfg["ldap_or_ad"]),
					resource.TestCheckResourceAttr(datasource, "connection_attributes.object_filter", createCfg["objectfilter"]),
					resource.TestCheckResourceAttr(datasource, "connection_attributes.account_attribute", createCfg["account_attribute"]),
					resource.TestCheckResourceAttr(datasource, "connection_attributes.entitlement_attribute", createCfg["entitlement_attribute"]),
					resource.TestCheckResourceAttr(datasource, "connection_attributes.page_size", createCfg["page_size"]),
				),
			},
		},
	})
}

func testAccADConnectionDataSourceConfig(jsonPath string) string {
	return fmt.Sprintf(`
provider "saviynt" {
  server_url = "%s"
  username   = "%s"
  password   = "%s"
}

locals {
  cfg = jsondecode(file("%s"))["ds"]
}

resource "saviynt_ad_connection_resource" "ad" {
  connection_type     = local.cfg.connection_type
  connection_name     = local.cfg.connection_name
  url                 = local.cfg.url
  password            = local.cfg.password
  username            = local.cfg.username
  searchfilter        = local.cfg.searchfilter
  base                = local.cfg.base
  group_search_base_dn= local.cfg.group_search_base_dn
  ldap_or_ad          = local.cfg.ldap_or_ad
  objectfilter        = local.cfg.objectfilter
  account_attribute   = local.cfg.account_attribute
  entitlement_attribute = local.cfg.entitlement_attribute
  page_size           = local.cfg.page_size
  user_attribute      = local.cfg.user_attribute

  endpoints_filter    = jsonencode(local.cfg.endpoints_filter)
  create_account_json = jsonencode(local.cfg.create_account_json)
  update_account_json = jsonencode(local.cfg.update_account_json)
  update_user_json    = jsonencode(local.cfg.update_user_json)
  enable_account_json = jsonencode(local.cfg.enable_account_json)

  account_name_rule   = local.cfg.account_name_rule
  remove_account_action = jsonencode(local.cfg.remove_account_action)
  set_random_password = local.cfg.set_random_password
}
  
data "saviynt_ad_connection_datasource" "test" {
	connection_name     = local.cfg.connection_name
	depends_on = [saviynt_ad_connection_resource.ad]
}
`, os.Getenv("SAVIYNT_URL"),
		os.Getenv("SAVIYNT_USERNAME"),
		os.Getenv("SAVIYNT_PASSWORD"),
		jsonPath,
		// os.Getenv("CI_PIPELINE_ID"),
	)
}
