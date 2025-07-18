// Copyright (c) 2025 Saviynt Inc.
// SPDX-License-Identifier: MPL-2.0

# For retrieving all dynamic attributes
data "saviynt_dynamic_attribute_datasource" "da" {}

data "saviynt_dynamic_attribute_datasource" "few" {
  securitysystem     = ["system1", "system2"]
  endpoint           = ["endpoint1", "endpoint2"]
  dynamic_attributes = ["da1", "da2"]
  requesttype        = ["req1", "req2"]
  offset             = "3"
  max                = "29"
  loggedinuser       = "username"
}