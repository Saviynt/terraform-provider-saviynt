// Copyright (c) 2025 Saviynt Inc.
// SPDX-License-Identifier: MPL-2.0

data "saviynt_roles_datasource" "example" {
  # Required parameter to control sensitive data visibility
  authenticate = false

  # Optional filters
  role_type   = "ENTERPRISE"
  requestable = "true"
  status      = "Active"
  max         = "10"
  offset      = "0"

  # Optional specific filters
  # username     = "john.doe"
  # role_name    = "Admin_Role"
  # description  = "Administrator"
  # display_name = "Admin Role"
  # risk         = "Medium"
  # sox_critical = "High"
  # sys_critical = "Medium"
  # privileged   = "Low"
  # confidentiality = "Medium"
  # level        = "2"

  # Optional advanced filters
  # role_query   = "r.role_name LIKE '%Admin%'"
  # hide_blank_values = "true"
  # requested_object = "entitlement"

  # Custom property filters (examples)
  # custom_property1 = "Department"
  # custom_property2 = "Business Unit"
}
