// Copyright (c) Saviynt Inc.
// SPDX-License-Identifier: MPL-2.0

// dynamic_attribute_resource_test.go contains the Terraform acceptance test suite for
// the `saviynt_dynamic_attribute_resource`. It validates the resource's full lifecycle against
// a live Saviynt instance by exercising:
//
//   - Create: provisions new dynamic attributes for an endpoint and asserts each attribute via JSONPath checks.
//   - Import: verifies the list of dynamic attributes can be imported by its endpoint name.
//   - Update: applies configuration changes and confirms the updated attribute values.
//
// Test data is loaded from `dynamic_attribute_test_data.json` using `testutil.LoadConnectorData`.
// Environment variables `SAVIYNT_URL`, `SAVIYNT_USERNAME`, and `SAVIYNT_PASSWORD` must be set
// to point at a valid Saviynt Security Manager before running these tests.
package testing

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"terraform-provider-Saviynt/util/testutil"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
)

func TestAccSaviyntDynamicAttributeResource(t *testing.T) {
	filePath := testutil.GetTestDataPath(t, "./test_data/dynamic_attribute_test_data.json")
	filePath = testutil.PrepareTestDataWithEnv(t, filePath)

	createReqCfg := testutil.LoadConnectorData(t, filePath, "createOnlyReq")
	createAllCfg := testutil.LoadConnectorData(t, filePath, "createWithAll")
	updateAllCfg := testutil.LoadConnectorData(t, filePath, "updateAll")
	addAttributeValueCfg := testutil.LoadConnectorData(t, filePath, "setAttributeValue")

	var dynAttrsCreateWithReq []map[string]string
	if err := json.Unmarshal([]byte(createReqCfg["dynamic_attributes"]), &dynAttrsCreateWithReq); err != nil {
		t.Fatalf("failed to unmarshal dynamic_attributes for create with required attributes only: %v", err)
	}

	var dynAttrsCreateWithAll []map[string]string
	if err := json.Unmarshal([]byte(createAllCfg["dynamic_attributes"]), &dynAttrsCreateWithAll); err != nil {
		t.Fatalf("failed to unmarshal dynamic_attributes for create with all attributes: %v", err)
	}

	var dynAttrsUpdateWithAll []map[string]string
	if err := json.Unmarshal([]byte(updateAllCfg["dynamic_attributes"]), &dynAttrsUpdateWithAll); err != nil {
		t.Fatalf("failed to unmarshal dynamic_attributes for updated with all attributes: %v", err)
	}
	var dynAttrsWithAttrValue []map[string]string
	if err := json.Unmarshal([]byte(addAttributeValueCfg["dynamic_attributes"]), &dynAttrsWithAttrValue); err != nil {
		t.Fatalf("failed to unmarshal dynamic_attributes for set attribute value: %v", err)
	}

	log.Print("Attribute 1: ", dynAttrsCreateWithReq[0]["attribute_name"])

	resourceNameWithReqOnly := "saviynt_dynamic_attribute_resource.da_req_attr"
	resourceNameWithAll := "saviynt_dynamic_attribute_resource.da_all_attr"
	resourceNameWithAttrValue := "saviynt_dynamic_attribute_resource.da_with_attr_value"

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { TestAccPreCheck(t) },
		ProtoV6ProviderFactories: TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create using only required attributes
			{
				Config: testAccDynamicAttributeResourceConfigWithReqAttr(filePath, "createOnlyReq"),
				Check: func(s *terraform.State) error {
					for resourceName, resourceState := range s.RootModule().Resources {
						log.Printf("[STATE DUMP] Resource: %s\n", resourceName)
						stateJSON, err := json.MarshalIndent(resourceState.Primary.Attributes, "", "  ")
						if err != nil {
							log.Printf("[ERROR] Failed to marshal state for %s: %v", resourceName, err)
						} else {
							log.Printf("[STATE ATTRIBUTES] %s\n", string(stateJSON))
						}
					}
					return nil
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceNameWithReqOnly, tfjsonpath.New("security_system"), knownvalue.StringExact(createReqCfg["security_system"])),
					statecheck.ExpectKnownValue(resourceNameWithReqOnly, tfjsonpath.New("endpoint"), knownvalue.StringExact(createReqCfg["endpoint"])),
					statecheck.ExpectKnownValue(resourceNameWithReqOnly, tfjsonpath.New("update_user"), knownvalue.StringExact(createReqCfg["update_user"])),
					// Check dynamic attributes map
					statecheck.ExpectKnownValue(resourceNameWithReqOnly, tfjsonpath.New("dynamic_attributes").AtMapKey(dynAttrsCreateWithReq[0]["attribute_name"]).AtMapKey("attribute_name"),
						knownvalue.StringExact(dynAttrsCreateWithReq[0]["attribute_name"])),
					statecheck.ExpectKnownValue(resourceNameWithReqOnly, tfjsonpath.New("dynamic_attributes").AtMapKey(dynAttrsCreateWithReq[0]["attribute_name"]).AtMapKey("request_type"),
						knownvalue.StringExact(dynAttrsCreateWithReq[0]["request_type"])),
					statecheck.ExpectKnownValue(resourceNameWithReqOnly, tfjsonpath.New("dynamic_attributes").AtMapKey(dynAttrsCreateWithReq[1]["attribute_name"]).AtMapKey("attribute_name"),
						knownvalue.StringExact(dynAttrsCreateWithReq[1]["attribute_name"])),
					statecheck.ExpectKnownValue(resourceNameWithReqOnly, tfjsonpath.New("dynamic_attributes").AtMapKey(dynAttrsCreateWithReq[1]["attribute_name"]).AtMapKey("request_type"),
						knownvalue.StringExact(dynAttrsCreateWithReq[1]["request_type"])),
				},
			},
			{
				Config: testAccDynamicAttributeResourceConfig(filePath, "createWithAll"),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceNameWithAll, tfjsonpath.New("security_system"), knownvalue.StringExact(createAllCfg["security_system"])),
					statecheck.ExpectKnownValue(resourceNameWithAll, tfjsonpath.New("endpoint"), knownvalue.StringExact(createAllCfg["endpoint"])),
					statecheck.ExpectKnownValue(resourceNameWithAll, tfjsonpath.New("update_user"), knownvalue.StringExact(createAllCfg["update_user"])),
					// Check dynamic attributes map
					statecheck.ExpectKnownValue(resourceNameWithAll, tfjsonpath.New("dynamic_attributes").AtMapKey(dynAttrsCreateWithAll[0]["attribute_name"]).AtMapKey("attribute_name"),
						knownvalue.StringExact(dynAttrsCreateWithAll[0]["attribute_name"])),

					statecheck.ExpectKnownValue(resourceNameWithAll, tfjsonpath.New("dynamic_attributes").AtMapKey(dynAttrsCreateWithAll[0]["attribute_name"]).AtMapKey("request_type"),
						knownvalue.StringExact(dynAttrsCreateWithAll[0]["request_type"])),

					statecheck.ExpectKnownValue(resourceNameWithAll, tfjsonpath.New("dynamic_attributes").AtMapKey(dynAttrsCreateWithAll[0]["attribute_name"]).AtMapKey("attribute_type"),
						knownvalue.StringExact(dynAttrsCreateWithAll[0]["attribute_type"])),

					statecheck.ExpectKnownValue(resourceNameWithAll, tfjsonpath.New("dynamic_attributes").AtMapKey(dynAttrsCreateWithAll[0]["attribute_name"]).AtMapKey("attribute_group"),
						knownvalue.StringExact(dynAttrsCreateWithAll[0]["attribute_group"])),

					statecheck.ExpectKnownValue(resourceNameWithAll, tfjsonpath.New("dynamic_attributes").AtMapKey(dynAttrsCreateWithAll[0]["attribute_name"]).AtMapKey("order_index"),
						knownvalue.StringExact(dynAttrsCreateWithAll[0]["order_index"])),

					statecheck.ExpectKnownValue(resourceNameWithAll, tfjsonpath.New("dynamic_attributes").AtMapKey(dynAttrsCreateWithAll[0]["attribute_name"]).AtMapKey("attribute_lable"),
						knownvalue.StringExact(dynAttrsCreateWithAll[0]["attribute_lable"])),

					statecheck.ExpectKnownValue(resourceNameWithAll, tfjsonpath.New("dynamic_attributes").AtMapKey(dynAttrsCreateWithAll[0]["attribute_name"]).AtMapKey("accounts_column"),
						knownvalue.StringExact(dynAttrsCreateWithAll[0]["accounts_column"])),

					statecheck.ExpectKnownValue(resourceNameWithAll, tfjsonpath.New("dynamic_attributes").AtMapKey(dynAttrsCreateWithAll[0]["attribute_name"]).AtMapKey("hide_on_create"),
						knownvalue.StringExact(dynAttrsCreateWithAll[0]["hide_on_create"])),

					statecheck.ExpectKnownValue(resourceNameWithAll, tfjsonpath.New("dynamic_attributes").AtMapKey(dynAttrsCreateWithAll[0]["attribute_name"]).AtMapKey("action_string"),
						knownvalue.StringExact(dynAttrsCreateWithAll[0]["action_string"])),

					statecheck.ExpectKnownValue(resourceNameWithAll, tfjsonpath.New("dynamic_attributes").AtMapKey(dynAttrsCreateWithAll[0]["attribute_name"]).AtMapKey("editable"),
						knownvalue.StringExact(dynAttrsCreateWithAll[0]["editable"])),

					statecheck.ExpectKnownValue(resourceNameWithAll, tfjsonpath.New("dynamic_attributes").AtMapKey(dynAttrsCreateWithAll[0]["attribute_name"]).AtMapKey("hide_on_update"),
						knownvalue.StringExact(dynAttrsCreateWithAll[0]["hide_on_update"])),

					statecheck.ExpectKnownValue(resourceNameWithAll, tfjsonpath.New("dynamic_attributes").AtMapKey(dynAttrsCreateWithAll[0]["attribute_name"]).AtMapKey("action_to_perform_when_parent_attribute_changes"),
						knownvalue.StringExact(dynAttrsCreateWithAll[0]["action_to_perform_when_parent_attribute_changes"])),

					statecheck.ExpectKnownValue(resourceNameWithAll, tfjsonpath.New("dynamic_attributes").AtMapKey(dynAttrsCreateWithAll[0]["attribute_name"]).AtMapKey("default_value"),
						knownvalue.StringExact(dynAttrsCreateWithAll[0]["default_value"])),

					statecheck.ExpectKnownValue(resourceNameWithAll, tfjsonpath.New("dynamic_attributes").AtMapKey(dynAttrsCreateWithAll[0]["attribute_name"]).AtMapKey("required"),
						knownvalue.StringExact(dynAttrsCreateWithAll[0]["required"])),

					statecheck.ExpectKnownValue(resourceNameWithAll, tfjsonpath.New("dynamic_attributes").AtMapKey(dynAttrsCreateWithAll[0]["attribute_name"]).AtMapKey("showonchild"),
						knownvalue.StringExact(dynAttrsCreateWithAll[0]["showonchild"])),

					statecheck.ExpectKnownValue(resourceNameWithAll, tfjsonpath.New("dynamic_attributes").AtMapKey(dynAttrsCreateWithAll[0]["attribute_name"]).AtMapKey("parent_attribute"),
						knownvalue.StringExact(dynAttrsCreateWithAll[0]["parent_attribute"])),

					statecheck.ExpectKnownValue(resourceNameWithAll, tfjsonpath.New("dynamic_attributes").AtMapKey(dynAttrsCreateWithAll[0]["attribute_name"]).AtMapKey("description_as_csv"),
						knownvalue.StringExact(dynAttrsCreateWithAll[0]["description_as_csv"])),

					statecheck.ExpectKnownValue(resourceNameWithAll, tfjsonpath.New("dynamic_attributes").AtMapKey(dynAttrsCreateWithAll[1]["attribute_name"]).AtMapKey("attribute_name"),
						knownvalue.StringExact(dynAttrsCreateWithAll[1]["attribute_name"])),

					statecheck.ExpectKnownValue(resourceNameWithAll, tfjsonpath.New("dynamic_attributes").AtMapKey(dynAttrsCreateWithAll[1]["attribute_name"]).AtMapKey("request_type"),
						knownvalue.StringExact(dynAttrsCreateWithAll[1]["request_type"])),

					statecheck.ExpectKnownValue(resourceNameWithAll, tfjsonpath.New("dynamic_attributes").AtMapKey(dynAttrsCreateWithAll[1]["attribute_name"]).AtMapKey("attribute_type"),
						knownvalue.StringExact(dynAttrsCreateWithAll[1]["attribute_type"])),

					statecheck.ExpectKnownValue(resourceNameWithAll, tfjsonpath.New("dynamic_attributes").AtMapKey(dynAttrsCreateWithAll[1]["attribute_name"]).AtMapKey("attribute_group"),
						knownvalue.StringExact(dynAttrsCreateWithAll[1]["attribute_group"])),

					statecheck.ExpectKnownValue(resourceNameWithAll, tfjsonpath.New("dynamic_attributes").AtMapKey(dynAttrsCreateWithAll[1]["attribute_name"]).AtMapKey("order_index"),
						knownvalue.StringExact(dynAttrsCreateWithAll[1]["order_index"])),

					statecheck.ExpectKnownValue(resourceNameWithAll, tfjsonpath.New("dynamic_attributes").AtMapKey(dynAttrsCreateWithAll[1]["attribute_name"]).AtMapKey("attribute_lable"),
						knownvalue.StringExact(dynAttrsCreateWithAll[1]["attribute_lable"])),

					statecheck.ExpectKnownValue(resourceNameWithAll, tfjsonpath.New("dynamic_attributes").AtMapKey(dynAttrsCreateWithAll[1]["attribute_name"]).AtMapKey("accounts_column"),
						knownvalue.StringExact(dynAttrsCreateWithAll[1]["accounts_column"])),

					statecheck.ExpectKnownValue(resourceNameWithAll, tfjsonpath.New("dynamic_attributes").AtMapKey(dynAttrsCreateWithAll[1]["attribute_name"]).AtMapKey("hide_on_create"),
						knownvalue.StringExact(dynAttrsCreateWithAll[1]["hide_on_create"])),

					statecheck.ExpectKnownValue(resourceNameWithAll, tfjsonpath.New("dynamic_attributes").AtMapKey(dynAttrsCreateWithAll[1]["attribute_name"]).AtMapKey("action_string"),
						knownvalue.StringExact(dynAttrsCreateWithAll[1]["action_string"])),

					statecheck.ExpectKnownValue(resourceNameWithAll, tfjsonpath.New("dynamic_attributes").AtMapKey(dynAttrsCreateWithAll[1]["attribute_name"]).AtMapKey("editable"),
						knownvalue.StringExact(dynAttrsCreateWithAll[1]["editable"])),

					statecheck.ExpectKnownValue(resourceNameWithAll, tfjsonpath.New("dynamic_attributes").AtMapKey(dynAttrsCreateWithAll[1]["attribute_name"]).AtMapKey("hide_on_update"),
						knownvalue.StringExact(dynAttrsCreateWithAll[1]["hide_on_update"])),

					statecheck.ExpectKnownValue(resourceNameWithAll, tfjsonpath.New("dynamic_attributes").AtMapKey(dynAttrsCreateWithAll[1]["attribute_name"]).AtMapKey("action_to_perform_when_parent_attribute_changes"),
						knownvalue.StringExact(dynAttrsCreateWithAll[1]["action_to_perform_when_parent_attribute_changes"])),

					statecheck.ExpectKnownValue(resourceNameWithAll, tfjsonpath.New("dynamic_attributes").AtMapKey(dynAttrsCreateWithAll[1]["attribute_name"]).AtMapKey("default_value"),
						knownvalue.StringExact(dynAttrsCreateWithAll[1]["default_value"])),

					statecheck.ExpectKnownValue(resourceNameWithAll, tfjsonpath.New("dynamic_attributes").AtMapKey(dynAttrsCreateWithAll[1]["attribute_name"]).AtMapKey("required"),
						knownvalue.StringExact(dynAttrsCreateWithAll[1]["required"])),

					statecheck.ExpectKnownValue(resourceNameWithAll, tfjsonpath.New("dynamic_attributes").AtMapKey(dynAttrsCreateWithAll[1]["attribute_name"]).AtMapKey("showonchild"),
						knownvalue.StringExact(dynAttrsCreateWithAll[1]["showonchild"])),

					statecheck.ExpectKnownValue(resourceNameWithAll, tfjsonpath.New("dynamic_attributes").AtMapKey(dynAttrsCreateWithAll[1]["attribute_name"]).AtMapKey("parent_attribute"),
						knownvalue.StringExact(dynAttrsCreateWithAll[1]["parent_attribute"])),

					statecheck.ExpectKnownValue(resourceNameWithAll, tfjsonpath.New("dynamic_attributes").AtMapKey(dynAttrsCreateWithAll[1]["attribute_name"]).AtMapKey("description_as_csv"),
						knownvalue.StringExact(dynAttrsCreateWithAll[1]["description_as_csv"])),
				},
			},
			// Import the create will dynamic attributes
			{
				ResourceName:      resourceNameWithAll,
				ImportStateId:     createAllCfg["endpoint"],
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{
					"update_user",
					"dynamic_attribute_errors",
					"msg",
					"regex",
					"attribute_value",
					fmt.Sprintf("dynamic_attributes.%s.regex", dynAttrsCreateWithAll[0]["attribute_name"]),
					fmt.Sprintf("dynamic_attributes.%s.regex", dynAttrsCreateWithAll[1]["attribute_name"]),
				},
			},
			// Update the create with all attributes
			{
				Config: testAccDynamicAttributeResourceConfig(filePath, "updateAll"),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceNameWithAll, tfjsonpath.New("security_system"), knownvalue.StringExact(updateAllCfg["security_system"])),
					statecheck.ExpectKnownValue(resourceNameWithAll, tfjsonpath.New("endpoint"), knownvalue.StringExact(updateAllCfg["endpoint"])),
					statecheck.ExpectKnownValue(resourceNameWithAll, tfjsonpath.New("update_user"), knownvalue.StringExact(updateAllCfg["update_user"])),
					// Check updated dynamic attributes
					statecheck.ExpectKnownValue(resourceNameWithAll, tfjsonpath.New("dynamic_attributes").AtMapKey(dynAttrsUpdateWithAll[0]["attribute_name"]).AtMapKey("attribute_name"),
						knownvalue.StringExact(dynAttrsUpdateWithAll[0]["attribute_name"])),

					statecheck.ExpectKnownValue(resourceNameWithAll, tfjsonpath.New("dynamic_attributes").AtMapKey(dynAttrsUpdateWithAll[0]["attribute_name"]).AtMapKey("request_type"),
						knownvalue.StringExact(dynAttrsUpdateWithAll[0]["request_type"])),

					statecheck.ExpectKnownValue(resourceNameWithAll, tfjsonpath.New("dynamic_attributes").AtMapKey(dynAttrsUpdateWithAll[0]["attribute_name"]).AtMapKey("attribute_type"),
						knownvalue.StringExact(dynAttrsUpdateWithAll[0]["attribute_type"])),

					statecheck.ExpectKnownValue(resourceNameWithAll, tfjsonpath.New("dynamic_attributes").AtMapKey(dynAttrsUpdateWithAll[0]["attribute_name"]).AtMapKey("attribute_group"),
						knownvalue.StringExact(dynAttrsUpdateWithAll[0]["attribute_group"])),

					statecheck.ExpectKnownValue(resourceNameWithAll, tfjsonpath.New("dynamic_attributes").AtMapKey(dynAttrsUpdateWithAll[0]["attribute_name"]).AtMapKey("order_index"),
						knownvalue.StringExact(dynAttrsUpdateWithAll[0]["order_index"])),

					statecheck.ExpectKnownValue(resourceNameWithAll, tfjsonpath.New("dynamic_attributes").AtMapKey(dynAttrsUpdateWithAll[0]["attribute_name"]).AtMapKey("attribute_lable"),
						knownvalue.StringExact(dynAttrsUpdateWithAll[0]["attribute_lable"])),

					statecheck.ExpectKnownValue(resourceNameWithAll, tfjsonpath.New("dynamic_attributes").AtMapKey(dynAttrsUpdateWithAll[0]["attribute_name"]).AtMapKey("accounts_column"),
						knownvalue.StringExact(dynAttrsUpdateWithAll[0]["accounts_column"])),

					statecheck.ExpectKnownValue(resourceNameWithAll, tfjsonpath.New("dynamic_attributes").AtMapKey(dynAttrsUpdateWithAll[0]["attribute_name"]).AtMapKey("hide_on_create"),
						knownvalue.StringExact(dynAttrsUpdateWithAll[0]["hide_on_create"])),

					statecheck.ExpectKnownValue(resourceNameWithAll, tfjsonpath.New("dynamic_attributes").AtMapKey(dynAttrsUpdateWithAll[0]["attribute_name"]).AtMapKey("action_string"),
						knownvalue.StringExact(dynAttrsUpdateWithAll[0]["action_string"])),

					statecheck.ExpectKnownValue(resourceNameWithAll, tfjsonpath.New("dynamic_attributes").AtMapKey(dynAttrsUpdateWithAll[0]["attribute_name"]).AtMapKey("editable"),
						knownvalue.StringExact(dynAttrsUpdateWithAll[0]["editable"])),

					statecheck.ExpectKnownValue(resourceNameWithAll, tfjsonpath.New("dynamic_attributes").AtMapKey(dynAttrsUpdateWithAll[0]["attribute_name"]).AtMapKey("hide_on_update"),
						knownvalue.StringExact(dynAttrsUpdateWithAll[0]["hide_on_update"])),

					statecheck.ExpectKnownValue(resourceNameWithAll, tfjsonpath.New("dynamic_attributes").AtMapKey(dynAttrsUpdateWithAll[0]["attribute_name"]).AtMapKey("action_to_perform_when_parent_attribute_changes"),
						knownvalue.StringExact(dynAttrsUpdateWithAll[0]["action_to_perform_when_parent_attribute_changes"])),

					statecheck.ExpectKnownValue(resourceNameWithAll, tfjsonpath.New("dynamic_attributes").AtMapKey(dynAttrsUpdateWithAll[0]["attribute_name"]).AtMapKey("default_value"),
						knownvalue.StringExact(dynAttrsUpdateWithAll[0]["default_value"])),

					statecheck.ExpectKnownValue(resourceNameWithAll, tfjsonpath.New("dynamic_attributes").AtMapKey(dynAttrsUpdateWithAll[0]["attribute_name"]).AtMapKey("required"),
						knownvalue.StringExact(dynAttrsUpdateWithAll[0]["required"])),

					statecheck.ExpectKnownValue(resourceNameWithAll, tfjsonpath.New("dynamic_attributes").AtMapKey(dynAttrsUpdateWithAll[0]["attribute_name"]).AtMapKey("showonchild"),
						knownvalue.StringExact(dynAttrsUpdateWithAll[0]["showonchild"])),

					statecheck.ExpectKnownValue(resourceNameWithAll, tfjsonpath.New("dynamic_attributes").AtMapKey(dynAttrsUpdateWithAll[0]["attribute_name"]).AtMapKey("parent_attribute"),
						knownvalue.StringExact(dynAttrsUpdateWithAll[0]["parent_attribute"])),

					statecheck.ExpectKnownValue(resourceNameWithAll, tfjsonpath.New("dynamic_attributes").AtMapKey(dynAttrsUpdateWithAll[0]["attribute_name"]).AtMapKey("description_as_csv"),
						knownvalue.StringExact(dynAttrsUpdateWithAll[0]["description_as_csv"])),

					statecheck.ExpectKnownValue(resourceNameWithAll, tfjsonpath.New("dynamic_attributes").AtMapKey(dynAttrsUpdateWithAll[1]["attribute_name"]).AtMapKey("attribute_name"),
						knownvalue.StringExact(dynAttrsUpdateWithAll[1]["attribute_name"])),

					statecheck.ExpectKnownValue(resourceNameWithAll, tfjsonpath.New("dynamic_attributes").AtMapKey(dynAttrsUpdateWithAll[1]["attribute_name"]).AtMapKey("request_type"),
						knownvalue.StringExact(dynAttrsUpdateWithAll[1]["request_type"])),

					statecheck.ExpectKnownValue(resourceNameWithAll, tfjsonpath.New("dynamic_attributes").AtMapKey(dynAttrsUpdateWithAll[1]["attribute_name"]).AtMapKey("attribute_type"),
						knownvalue.StringExact(dynAttrsUpdateWithAll[1]["attribute_type"])),

					statecheck.ExpectKnownValue(resourceNameWithAll, tfjsonpath.New("dynamic_attributes").AtMapKey(dynAttrsUpdateWithAll[1]["attribute_name"]).AtMapKey("attribute_group"),
						knownvalue.StringExact(dynAttrsUpdateWithAll[1]["attribute_group"])),

					statecheck.ExpectKnownValue(resourceNameWithAll, tfjsonpath.New("dynamic_attributes").AtMapKey(dynAttrsUpdateWithAll[1]["attribute_name"]).AtMapKey("order_index"),
						knownvalue.StringExact(dynAttrsUpdateWithAll[1]["order_index"])),

					statecheck.ExpectKnownValue(resourceNameWithAll, tfjsonpath.New("dynamic_attributes").AtMapKey(dynAttrsUpdateWithAll[1]["attribute_name"]).AtMapKey("attribute_lable"),
						knownvalue.StringExact(dynAttrsUpdateWithAll[1]["attribute_lable"])),

					statecheck.ExpectKnownValue(resourceNameWithAll, tfjsonpath.New("dynamic_attributes").AtMapKey(dynAttrsUpdateWithAll[1]["attribute_name"]).AtMapKey("accounts_column"),
						knownvalue.StringExact(dynAttrsUpdateWithAll[1]["accounts_column"])),

					statecheck.ExpectKnownValue(resourceNameWithAll, tfjsonpath.New("dynamic_attributes").AtMapKey(dynAttrsUpdateWithAll[1]["attribute_name"]).AtMapKey("hide_on_create"),
						knownvalue.StringExact(dynAttrsUpdateWithAll[1]["hide_on_create"])),

					statecheck.ExpectKnownValue(resourceNameWithAll, tfjsonpath.New("dynamic_attributes").AtMapKey(dynAttrsUpdateWithAll[1]["attribute_name"]).AtMapKey("action_string"),
						knownvalue.StringExact(dynAttrsUpdateWithAll[1]["action_string"])),

					statecheck.ExpectKnownValue(resourceNameWithAll, tfjsonpath.New("dynamic_attributes").AtMapKey(dynAttrsUpdateWithAll[1]["attribute_name"]).AtMapKey("editable"),
						knownvalue.StringExact(dynAttrsUpdateWithAll[1]["editable"])),

					statecheck.ExpectKnownValue(resourceNameWithAll, tfjsonpath.New("dynamic_attributes").AtMapKey(dynAttrsUpdateWithAll[1]["attribute_name"]).AtMapKey("hide_on_update"),
						knownvalue.StringExact(dynAttrsUpdateWithAll[1]["hide_on_update"])),

					statecheck.ExpectKnownValue(resourceNameWithAll, tfjsonpath.New("dynamic_attributes").AtMapKey(dynAttrsUpdateWithAll[1]["attribute_name"]).AtMapKey("action_to_perform_when_parent_attribute_changes"),
						knownvalue.StringExact(dynAttrsUpdateWithAll[1]["action_to_perform_when_parent_attribute_changes"])),

					statecheck.ExpectKnownValue(resourceNameWithAll, tfjsonpath.New("dynamic_attributes").AtMapKey(dynAttrsUpdateWithAll[1]["attribute_name"]).AtMapKey("default_value"),
						knownvalue.StringExact(dynAttrsUpdateWithAll[1]["default_value"])),

					statecheck.ExpectKnownValue(resourceNameWithAll, tfjsonpath.New("dynamic_attributes").AtMapKey(dynAttrsUpdateWithAll[1]["attribute_name"]).AtMapKey("required"),
						knownvalue.StringExact(dynAttrsUpdateWithAll[1]["required"])),

					statecheck.ExpectKnownValue(resourceNameWithAll, tfjsonpath.New("dynamic_attributes").AtMapKey(dynAttrsUpdateWithAll[1]["attribute_name"]).AtMapKey("showonchild"),
						knownvalue.StringExact(dynAttrsUpdateWithAll[1]["showonchild"])),

					statecheck.ExpectKnownValue(resourceNameWithAll, tfjsonpath.New("dynamic_attributes").AtMapKey(dynAttrsUpdateWithAll[1]["attribute_name"]).AtMapKey("parent_attribute"),
						knownvalue.StringExact(dynAttrsUpdateWithAll[1]["parent_attribute"])),

					statecheck.ExpectKnownValue(resourceNameWithAll, tfjsonpath.New("dynamic_attributes").AtMapKey(dynAttrsUpdateWithAll[1]["attribute_name"]).AtMapKey("description_as_csv"),
						knownvalue.StringExact(dynAttrsUpdateWithAll[1]["description_as_csv"])),
				},
			},
			{
				Config: testAccDynamicAttributeResourceConfigWithAttrValue(filePath, "setAttributeValue"),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(resourceNameWithAttrValue, tfjsonpath.New("security_system"), knownvalue.StringExact(addAttributeValueCfg["security_system"])),
					statecheck.ExpectKnownValue(resourceNameWithAttrValue, tfjsonpath.New("endpoint"), knownvalue.StringExact(addAttributeValueCfg["endpoint"])),
					statecheck.ExpectKnownValue(resourceNameWithAttrValue, tfjsonpath.New("update_user"), knownvalue.StringExact(addAttributeValueCfg["update_user"])),
					// Check updated dynamic attributes
					statecheck.ExpectKnownValue(resourceNameWithAttrValue, tfjsonpath.New("dynamic_attributes").AtMapKey(dynAttrsWithAttrValue[0]["attribute_name"]).AtMapKey("attribute_name"),
						knownvalue.StringExact(dynAttrsWithAttrValue[0]["attribute_name"])),

					statecheck.ExpectKnownValue(resourceNameWithAttrValue, tfjsonpath.New("dynamic_attributes").AtMapKey(dynAttrsWithAttrValue[0]["attribute_name"]).AtMapKey("request_type"),
						knownvalue.StringExact(dynAttrsWithAttrValue[0]["request_type"])),

					statecheck.ExpectKnownValue(resourceNameWithAttrValue, tfjsonpath.New("dynamic_attributes").AtMapKey(dynAttrsWithAttrValue[0]["attribute_name"]).AtMapKey("attribute_type"),
						knownvalue.StringExact(dynAttrsWithAttrValue[0]["attribute_type"])),

					statecheck.ExpectKnownValue(resourceNameWithAttrValue, tfjsonpath.New("dynamic_attributes").AtMapKey(dynAttrsWithAttrValue[0]["attribute_name"]).AtMapKey("attribute_group"),
						knownvalue.StringExact(dynAttrsWithAttrValue[0]["attribute_group"])),

					statecheck.ExpectKnownValue(resourceNameWithAttrValue, tfjsonpath.New("dynamic_attributes").AtMapKey(dynAttrsWithAttrValue[0]["attribute_name"]).AtMapKey("order_index"),
						knownvalue.StringExact(dynAttrsWithAttrValue[0]["order_index"])),

					statecheck.ExpectKnownValue(resourceNameWithAttrValue, tfjsonpath.New("dynamic_attributes").AtMapKey(dynAttrsWithAttrValue[0]["attribute_name"]).AtMapKey("attribute_lable"),
						knownvalue.StringExact(dynAttrsWithAttrValue[0]["attribute_lable"])),

					statecheck.ExpectKnownValue(resourceNameWithAttrValue, tfjsonpath.New("dynamic_attributes").AtMapKey(dynAttrsWithAttrValue[0]["attribute_name"]).AtMapKey("accounts_column"),
						knownvalue.StringExact(dynAttrsWithAttrValue[0]["accounts_column"])),

					statecheck.ExpectKnownValue(resourceNameWithAttrValue, tfjsonpath.New("dynamic_attributes").AtMapKey(dynAttrsWithAttrValue[0]["attribute_name"]).AtMapKey("hide_on_create"),
						knownvalue.StringExact(dynAttrsWithAttrValue[0]["hide_on_create"])),

					statecheck.ExpectKnownValue(resourceNameWithAttrValue, tfjsonpath.New("dynamic_attributes").AtMapKey(dynAttrsWithAttrValue[0]["attribute_name"]).AtMapKey("action_string"),
						knownvalue.StringExact(dynAttrsWithAttrValue[0]["action_string"])),

					statecheck.ExpectKnownValue(resourceNameWithAttrValue, tfjsonpath.New("dynamic_attributes").AtMapKey(dynAttrsWithAttrValue[0]["attribute_name"]).AtMapKey("editable"),
						knownvalue.StringExact(dynAttrsWithAttrValue[0]["editable"])),

					statecheck.ExpectKnownValue(resourceNameWithAttrValue, tfjsonpath.New("dynamic_attributes").AtMapKey(dynAttrsWithAttrValue[0]["attribute_name"]).AtMapKey("hide_on_update"),
						knownvalue.StringExact(dynAttrsWithAttrValue[0]["hide_on_update"])),

					statecheck.ExpectKnownValue(resourceNameWithAttrValue, tfjsonpath.New("dynamic_attributes").AtMapKey(dynAttrsWithAttrValue[0]["attribute_name"]).AtMapKey("action_to_perform_when_parent_attribute_changes"),
						knownvalue.StringExact(dynAttrsWithAttrValue[0]["action_to_perform_when_parent_attribute_changes"])),

					statecheck.ExpectKnownValue(resourceNameWithAttrValue, tfjsonpath.New("dynamic_attributes").AtMapKey(dynAttrsWithAttrValue[0]["attribute_name"]).AtMapKey("default_value"),
						knownvalue.StringExact(dynAttrsWithAttrValue[0]["default_value"])),

					statecheck.ExpectKnownValue(resourceNameWithAttrValue, tfjsonpath.New("dynamic_attributes").AtMapKey(dynAttrsWithAttrValue[0]["attribute_name"]).AtMapKey("required"),
						knownvalue.StringExact(dynAttrsWithAttrValue[0]["required"])),

					statecheck.ExpectKnownValue(resourceNameWithAttrValue, tfjsonpath.New("dynamic_attributes").AtMapKey(dynAttrsWithAttrValue[0]["attribute_name"]).AtMapKey("attribute_value"),
						knownvalue.StringExact(dynAttrsWithAttrValue[0]["attribute_value"])),

					statecheck.ExpectKnownValue(resourceNameWithAttrValue, tfjsonpath.New("dynamic_attributes").AtMapKey(dynAttrsWithAttrValue[0]["attribute_name"]).AtMapKey("showonchild"),
						knownvalue.StringExact(dynAttrsWithAttrValue[0]["showonchild"])),

					statecheck.ExpectKnownValue(resourceNameWithAttrValue, tfjsonpath.New("dynamic_attributes").AtMapKey(dynAttrsWithAttrValue[0]["attribute_name"]).AtMapKey("parent_attribute"),
						knownvalue.StringExact(dynAttrsWithAttrValue[0]["parent_attribute"])),

					statecheck.ExpectKnownValue(resourceNameWithAttrValue, tfjsonpath.New("dynamic_attributes").AtMapKey(dynAttrsWithAttrValue[0]["attribute_name"]).AtMapKey("description_as_csv"),
						knownvalue.StringExact(dynAttrsWithAttrValue[0]["description_as_csv"])),
				},
			},
		},
	})
}

func testAccDynamicAttributeResourceConfigWithReqAttr(jsonPath, operation string) string {
	return fmt.Sprintf(`
		provider "saviynt" {
			server_url = "%s"
			username   = "%s"
			password   = "%s"
		}

		locals {
			cfg = jsondecode(file("%s"))["%s"]
		}

		resource "saviynt_dynamic_attribute_resource" "da_req_attr" {
			security_system = local.cfg.security_system
			endpoint        = local.cfg.endpoint
			update_user     = local.cfg.update_user
			
			dynamic_attributes = tomap({
			(local.cfg.dynamic_attributes[0].attribute_name)= {
				attribute_name = local.cfg.dynamic_attributes[0].attribute_name
				request_type   = local.cfg.dynamic_attributes[0].request_type
			},
			(local.cfg.dynamic_attributes[1].attribute_name)= {
				attribute_name = local.cfg.dynamic_attributes[1].attribute_name
				request_type   = local.cfg.dynamic_attributes[1].request_type
			}
			})
		}
	`, os.Getenv("SAVIYNT_URL"),
		os.Getenv("SAVIYNT_USERNAME"),
		os.Getenv("SAVIYNT_PASSWORD"),
		jsonPath,
		operation,
	)
}

func testAccDynamicAttributeResourceConfig(jsonPath, operation string) string {
	return fmt.Sprintf(`
		provider "saviynt" {
			server_url = "%s"
			username   = "%s"
			password   = "%s"
		}

		locals {
			cfg = jsondecode(file("%s"))["%s"]
		}

		resource "saviynt_dynamic_attribute_resource" "da_all_attr" {
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
					parent_attribute                          = local.cfg.dynamic_attributes[0].parent_attribute
					required                                  = local.cfg.dynamic_attributes[0].required
					regex                                     = local.cfg.dynamic_attributes[0].regex
					showonchild                               = local.cfg.dynamic_attributes[0].showonchild
					description_as_csv                        = local.cfg.dynamic_attributes[0].description_as_csv
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
					parent_attribute                          = local.cfg.dynamic_attributes[1].parent_attribute
					hide_on_update                            = local.cfg.dynamic_attributes[1].hide_on_update
					action_to_perform_when_parent_attribute_changes = local.cfg.dynamic_attributes[1].action_to_perform_when_parent_attribute_changes
					default_value                             = local.cfg.dynamic_attributes[1].default_value
					required                                  = local.cfg.dynamic_attributes[1].required
					regex                                     = local.cfg.dynamic_attributes[1].regex
					showonchild                               = local.cfg.dynamic_attributes[1].showonchild
					description_as_csv                        = local.cfg.dynamic_attributes[1].description_as_csv
				}
			}
		}
`, os.Getenv("SAVIYNT_URL"),
		os.Getenv("SAVIYNT_USERNAME"),
		os.Getenv("SAVIYNT_PASSWORD"),
		jsonPath,
		operation,
	)
}

func testAccDynamicAttributeResourceConfigWithAttrValue(jsonPath, operation string) string {
	return fmt.Sprintf(`
		provider "saviynt" {
			server_url = "%s"
			username   = "%s"
			password   = "%s"
		}

		locals {
			cfg = jsondecode(file("%s"))["%s"]
		}

		resource "saviynt_dynamic_attribute_resource" "da_with_attr_value" {
			security_system = local.cfg.security_system
			endpoint        = local.cfg.endpoint
			update_user     = local.cfg.update_user
			
			dynamic_attributes = {
				(local.cfg.dynamic_attributes[0].attribute_name) = {
					attribute_name = local.cfg.dynamic_attributes[0].attribute_name
					request_type   = local.cfg.dynamic_attributes[0].request_type
					attribute_type                            = local.cfg.dynamic_attributes[0].attribute_type
					attribute_group                           = local.cfg.dynamic_attributes[0].attribute_group
					attribute_value							= local.cfg.dynamic_attributes[0].attribute_value
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
					regex                                     = local.cfg.dynamic_attributes[0].regex
					showonchild                               = local.cfg.dynamic_attributes[0].showonchild
					parent_attribute                          = local.cfg.dynamic_attributes[0].parent_attribute
					description_as_csv                        = local.cfg.dynamic_attributes[0].description_as_csv
				}
			}
		}
`, os.Getenv("SAVIYNT_URL"),
		os.Getenv("SAVIYNT_USERNAME"),
		os.Getenv("SAVIYNT_PASSWORD"),
		jsonPath,
		operation,
	)
}
