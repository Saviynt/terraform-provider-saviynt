// Copyright (c) 2025 Saviynt Inc.
// SPDX-License-Identifier: MPL-2.0

data "saviynt_adsi_connection_datasource" "example" {
  connection_key = "123"
  # connection_name="sample"          #Either one can be used
}
