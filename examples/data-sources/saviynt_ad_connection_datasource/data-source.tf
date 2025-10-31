// Copyright (c) 2025 Saviynt Inc.
// SPDX-License-Identifier: MPL-2.0

data "saviynt_ad_connection_datasource" "ad" {
  connection_key = "123"
  # connection_name="sample"          #Either one can be used

  authenticate = true
}