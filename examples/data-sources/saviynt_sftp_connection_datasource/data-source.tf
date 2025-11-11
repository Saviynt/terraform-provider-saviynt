// Copyright (c) 2025 Saviynt Inc.
// SPDX-License-Identifier: MPL-2.0

data "saviynt_sftp_connection_datasource" "sftp" {
  connection_name = "My_SFTP_Connection"
  # connection_key = "123"          # Either connection_name or connection_key can be used

  authenticate = true
}
