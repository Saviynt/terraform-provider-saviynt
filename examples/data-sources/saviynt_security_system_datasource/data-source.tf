// Copyright (c) 2025 Saviynt Inc.
// SPDX-License-Identifier: MPL-2.0

# List all the security systems
data "saviynt_security_system_datasource" "all" {}


# Get a security system by its name
data "saviynt_security_system_datasource" "by_name" {
  systemname = "sample"
}