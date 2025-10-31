// Copyright (c) 2025 Saviynt Inc.
// SPDX-License-Identifier: MPL-2.0

data "saviynt_privilege_datasource" "test" {
  // Required
  authenticate = true
  endpoint     = "sample_endpoint"

  // optional
  entitlement_type = "sample_ent_type"
  max              = "2"
  offset           = "3"
}