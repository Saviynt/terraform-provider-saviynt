// Copyright (c) Saviynt Inc.
// SPDX-License-Identifier: MPL-2.0

resource "saviynt_dynamic_attribute_resource" "example" {
  security_system = "shaleenhuddle"
  endpoint        = "sample-106"
  update_user     = "admin"

  dynamic_attributes = {
    dynamic_attr_1 = { 
      attribute_name                            = "dynamic_attr_1"
      request_type                              = "ACCOUNT"
      attribute_type                            = "BOOLEAN"
      attribute_group                           = "Performance"
      order_index                               = "1"
      attribute_lable                           = "API Timeout (ms)"
      accounts_column                           = "false"
      hide_on_create                            = "false"
      action_string                             = "action_string"
      editable                                  = "true"
      hide_on_update                            = "false"
      actiontoperformwhenparentattributechanges = "action"
      default_value                             = "30000"
      required                                  = "true"
      attribute_value                           = "45000"
      showonchild                               = "false"
      description_as_csv                        = "timeout,maximum wait time"
    }
    dynamic_attr_2 = {
      attribute_name = "dynamic_attr_2"
      request_type   = "SERVICE ACCOUNT"
      # attribute_type  = "string"
      # attribute_group = "user_details"
      order_index = "1"
      required    = "true"
      editable    = "true"
      # hide_on_create  = "false"
      # hide_on_update  = "false"
    }
  }
}