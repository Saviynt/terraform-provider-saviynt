// Copyright (c) 2025 Saviynt Inc.
// SPDX-License-Identifier: MPL-2.0

resource "saviynt_entitlement_type_resource" "example" {
  # Required attributes
  entitlement_name = "example_ent_type"
  endpoint_name    = "example_endpoint"

  # Optional attributes
  display_name                      = "Test Entitlement Type"
  entitlement_description           = "This is a test entitlement type for demonstration"
  workflow                          = "Autoapprovalwf"
  enable_entitlement_to_role_sync   = true
  available_query_service_account   = "TestServiceAccount"
  selected_query_service_account    = "TestServiceAccount"
  ars_requestable_entitlement_query = "SELECT * FROM entitlements WHERE requestable = 1"
  ars_selected_entitlement_query    = "SELECT * FROM entitlements WHERE selected = 1"
  certifiable                       = true
  create_task_action                = ["enableRollback", "removeTaskForExistingEntitlements"]
  request_dates_conf_json = jsonencode({
    "startDate" : true,
    "endDate" : true,
    "justification" : true
  })
  order_index                    = 1
  required_in_request            = true
  required_in_service_request    = true
  hierarchy_required             = "0"
  show_ent_type_on               = "1"
  enable_provisioning_priority   = false
  request_option                 = "FREEFORMTEXT"
  recon                          = true
  exclude_rule_assgn_ents_in_req = true

  start_date_in_revoke_request            = "false"
  start_end_date_in_request               = "true"
  allow_remove_all_entitlement_in_request = "false"

  custom_property1 = "Custom Property 1 Value"
  custom_property2 = "Custom Property 2 Value"
  custom_property3 = "Custom Property 3 Value"
  custom_property4 = "Custom Property 4 Value"
  custom_property5 = "Custom Property 5 Value"
}