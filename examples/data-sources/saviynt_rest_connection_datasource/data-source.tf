// Copyright (c) Saviynt Inc.
// SPDX-License-Identifier: MPL-2.0

data "saviynt_rest_connection_datasource" "rest" {
  connection_key = "123"
  # connection_name="sample"         #Either one can be used
}