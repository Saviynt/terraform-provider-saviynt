// Copyright (c) 2025 Saviynt Inc.
// SPDX-License-Identifier: MPL-2.0

resource "saviynt_privilege_resource" "test_privilege" {
  security_system  = "sample_security_system"
  endpoint         = "sample_endpoint"
  entitlement_type = "sample_ent_type"

  privileges = {
    "test_privilege_1" = {
      attribute_name   = "test_privilege_1"
      attribute_type   = "STRING"
      attribute_config = "USER-BASED"

      // optional attributes
      order_index      = "1"
      default_value    = "default_value_1"
      label            = "Test String Privilege"
      attribute_group  = "Group1"
      parent_attribute = "parent_attr_1"
      child_action     = "child_action_1"
      description      = "A test string privilege"
      required         = false
      requestable      = true
      hide_on_create   = true
      hide_on_update   = true
      action_string    = "CREATE"
    }

    "test_privilege_2" = {
      attribute_name   = "test_privilege_2"
      attribute_type   = "SINGLE SELECT FROM SQL QUERY"
      attribute_config = "ENTITLEMENT-BASED"
      order_index      = "2"
      default_value    = "true"
      label            = "Test Privilege label"
      attribute_group  = "Group2"
      parent_attribute = ""
      description      = "A test privilege"
      required         = true
      requestable      = false
      hide_on_create   = true
      hide_on_update   = true
      action_string    = "UPDATE"
    }
  }
}

