// Copyright (c) 2025 Saviynt Inc.
// SPDX-License-Identifier: MPL-2.0

# Gets all entitlement types(display count is 50 by default)
data "saviynt_entitlement_type_datasource" "sample" {
  authenticate = true
}

# Gets entitlement types by name or endpoint
data "saviynt_entitlement_type_datasource" "sample" {
  entitlement_name = "sample_ent_type"
  #   endpoint_name="sample_endpoint"
  authenticate = true
}