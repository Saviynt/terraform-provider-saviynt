// Copyright (c) 2025 Saviynt Inc.
// SPDX-License-Identifier: MPL-2.0

data "saviynt_github_rest_connection_datasource" "g" {
  connection_key = "123"
  # connection_name="sample"          #Either one can be used
}
