// Copyright (c) Saviynt Inc.
// SPDX-License-Identifier: MPL-2.0

// dynamic_attribute_datasource_test.go contains the Terraform acceptance test suite for
// the `saviynt_dynamic_attribute_datasource` data source. It verifies that once an endpoint
// resource is created, the data source can successfully look it up and populate all expected
// attributes from Saviynt.
//
// The test flow is as follows:
//  1. Load test fixture data from `dynamic_attribute_test_data.json`.
//  2. Apply a resource configuration to create a `saviynt_dyanamic_attribute_resource`.
//  3. Use `data.saviynt_dynamic_attribute_datasource.test` to read back the same dynamic attribute by name.
//  4. Assert that each field in the data source matches the values provided during creation.
//
// Test data is loaded from `dynamic_attribute_test_data.json` using `testutil.LoadConnectorData`.
// Environment variables `SAVIYNT_URL`, `SAVIYNT_USERNAME`, and `SAVIYNT_PASSWORD` must be set
// to point at a valid Saviynt Security Manager before running these tests.
package testing

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"terraform-provider-Saviynt/util/testutil"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

func checkDynamicAttribute(expected map[string]string) resource.TestCheckFunc {
    return func(s *terraform.State) error {
        res, ok := s.RootModule().Resources["data.saviynt_dynamic_attribute_datasource.test"]
        if !ok {
            return fmt.Errorf("data source not found")
        }

        fmt.Printf("\n=== DEBUG: Looking for attribute: %s ===\n", expected["attribute_name"])
        
        fmt.Println("--- All dynamic_attributes_list entries in state ---")
        attributeBlocks := make(map[string]map[string]string) 
        
        for k, v := range res.Primary.Attributes {
            if strings.HasPrefix(k, "dynamic_attributes_list.") {
                fmt.Printf("  %s = %s\n", k, v)
                
                parts := strings.Split(k, ".")
                if len(parts) == 3 {
                    index := parts[1]
                    attrName := parts[2]
                    
                    if attributeBlocks[index] == nil {
                        attributeBlocks[index] = make(map[string]string)
                    }
                    attributeBlocks[index][attrName] = v
                }
            }
        }
        
        var matchedIndex string
        expectedName := expected["attribute_name"]
        
        fmt.Printf("Searching for block with attribute_name: %q\n", expectedName)
        for index, attrs := range attributeBlocks {
            if attrs["attribute_name"] == expectedName {
                matchedIndex = index
                fmt.Printf("Found matching block at index: %s\n", index)
                fmt.Printf("Block contents: %+v\n", attrs)
                break
            }
        }
        
        if matchedIndex == "" {
            fmt.Printf("ERROR: No block found with attribute_name: %q\n", expectedName)
            fmt.Println("Available attribute_name values:")
            for index, attrs := range attributeBlocks {
                fmt.Printf("  Index %s: %q\n", index, attrs["attribute_name"])
            }
            return fmt.Errorf("attribute_name %q not found in any dynamic_attributes_list block", expectedName)
        }
        
        fmt.Printf("\n--- Validating attributes for block %s ---\n", matchedIndex)
        actualBlock := attributeBlocks[matchedIndex]
        
        for expectedKey, expectedVal := range expected {
            fmt.Printf("Checking: %s\n", expectedKey)
            fmt.Printf("Expected: %q\n", expectedVal)
            
            actualVal, exists := actualBlock[expectedKey]
            if !exists {
                fmt.Printf("ERROR: Key %q not found in state block\n", expectedKey)
                fmt.Printf("    Available keys: %v\n", func() []string {
                    keys := make([]string, 0, len(actualBlock))
                    for k := range actualBlock {
                        keys = append(keys, k)
                    }
                    return keys
                }())
                return fmt.Errorf("attribute %q not found in dynamic_attributes_list block %s", expectedKey, matchedIndex)
            }
            
            fmt.Printf("Actual: %q\n", actualVal)
            
            if actualVal != expectedVal {
                fmt.Printf("ERROR: Value mismatch\n")
                return fmt.Errorf("attribute %q in block %s: expected %q, got %q", expectedKey, matchedIndex, expectedVal, actualVal)
            } else {
                fmt.Printf("MATCH\n")
            }
        }
        
        fmt.Printf("All validations passed for attribute: %s\n\n", expectedName)
        return nil
    }
}

func TestAccSaviyntDynamicAttributeDataSource(t *testing.T) {
	filePath := testutil.GetTestDataPath(t, "./test_data/dynamic_attribute_test_data.json")
	filePath = testutil.PrepareTestDataWithEnv(t, filePath)
	createCfg := testutil.LoadConnectorData(t, filePath, "ds")
	datasource := "data.saviynt_dynamic_attribute_datasource.test"

	var dynamicAttributes []map[string]string
	if err := json.Unmarshal([]byte(createCfg["dynamic_attributes"]), &dynamicAttributes); err != nil {
		t.Fatalf("Failed to unmarshal dynamic_attributes: %v", err)
	}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { TestAccPreCheck(t) },
		ProtoV6ProviderFactories: TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccSaviyntDynamicAttributeDataSourceConfig(filePath),
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

					checkDynamicAttribute(dynamicAttributes[0]),
					checkDynamicAttribute(dynamicAttributes[1]),
				),
			},
		},
	})
}

func testAccSaviyntDynamicAttributeDataSourceConfig(jsonPath string) string {
	return fmt.Sprintf(`
	provider "saviynt" {
		server_url = "%s"
		username   = "%s"
		password   = "%s"
	}

	locals {
		cfg = jsondecode(file("%s"))["ds"]
	}

	resource "saviynt_dynamic_attribute_resource" "da" {
		security_system = local.cfg.security_system
		endpoint        = local.cfg.endpoint
		update_user     = local.cfg.update_user
		
		dynamic_attributes = {
			(local.cfg.dynamic_attributes[0].attribute_name) = {
				attribute_name 							= local.cfg.dynamic_attributes[0].attribute_name
				request_type   							= local.cfg.dynamic_attributes[0].request_type
				attribute_type                            = local.cfg.dynamic_attributes[0].attribute_type
				attribute_group                           = local.cfg.dynamic_attributes[0].attribute_group
				order_index                               = local.cfg.dynamic_attributes[0].order_index
				attribute_lable                           = local.cfg.dynamic_attributes[0].attribute_lable
				accounts_column                           = local.cfg.dynamic_attributes[0].accounts_column
				hide_on_create                            = local.cfg.dynamic_attributes[0].hide_on_create
				action_string                             = local.cfg.dynamic_attributes[0].action_string
				editable                                  = local.cfg.dynamic_attributes[0].editable
				hide_on_update                            = local.cfg.dynamic_attributes[0].hide_on_update
				action_to_perform_when_parent_attribute_changes = local.cfg.dynamic_attributes[0].action_to_perform_when_parent_attribute_changes
				default_value                             = local.cfg.dynamic_attributes[0].default_value
				required                                  = local.cfg.dynamic_attributes[0].required
				showonchild                               = local.cfg.dynamic_attributes[0].showonchild
			},
			(local.cfg.dynamic_attributes[1].attribute_name) = {
				attribute_name 							= local.cfg.dynamic_attributes[1].attribute_name
				request_type   							= local.cfg.dynamic_attributes[1].request_type
				attribute_type                            = local.cfg.dynamic_attributes[1].attribute_type
				attribute_group                           = local.cfg.dynamic_attributes[1].attribute_group
				order_index                               = local.cfg.dynamic_attributes[1].order_index
				attribute_lable                           = local.cfg.dynamic_attributes[1].attribute_lable
				accounts_column                           = local.cfg.dynamic_attributes[1].accounts_column
				hide_on_create                            = local.cfg.dynamic_attributes[1].hide_on_create
				action_string                             = local.cfg.dynamic_attributes[1].action_string
				editable                                  = local.cfg.dynamic_attributes[1].editable
				hide_on_update                            = local.cfg.dynamic_attributes[1].hide_on_update
				action_to_perform_when_parent_attribute_changes = local.cfg.dynamic_attributes[1].action_to_perform_when_parent_attribute_changes
				default_value                             = local.cfg.dynamic_attributes[1].default_value
				required                                  = local.cfg.dynamic_attributes[1].required
				showonchild                               = local.cfg.dynamic_attributes[1].showonchild
			}
		}
	}
	
	data "saviynt_dynamic_attribute_datasource" "test" {
		endpoint           = [local.cfg.endpoint]
		authenticate       = true
		depends_on = [saviynt_dynamic_attribute_resource.da]
	}
`, os.Getenv("SAVIYNT_URL"),
		os.Getenv("SAVIYNT_USERNAME"),
		os.Getenv("SAVIYNT_PASSWORD"),
		jsonPath,
	)
}