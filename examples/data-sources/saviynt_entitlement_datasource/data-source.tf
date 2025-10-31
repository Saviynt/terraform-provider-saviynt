// Copyright (c) 2025 Saviynt Inc.
// SPDX-License-Identifier: MPL-2.0

# Gets all entitlements (display count is 50 by default)
data "saviynt_entitlement_datasource" "all_entitlements" {
  authenticate = true
}

data "saviynt_entitlement_datasource" "filtered_entitlements" {
  # Required attribute
  authenticate = true

  # Optional filter attributes
  endpoint = "sample_endpoint"
  # entitlementtype      = "sample_ent_type"
  # entitlement_value    = "sample_ent"
  # ent_query           = "ent.id like '1'"
}
