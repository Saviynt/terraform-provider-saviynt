// Copyright (c) 2025 Saviynt Inc.
// SPDX-License-Identifier: MPL-2.0

# List all endpoints
data "saviynt_endpoints_datasource" "all" {}

# Get endpoint by name
data "saviynt_endpoints_datasource" "by_name" {
  endpointname = "sample"
  authenticate = true
}