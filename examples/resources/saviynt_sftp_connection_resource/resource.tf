// Copyright (c) 2025 Saviynt Inc.
// SPDX-License-Identifier: MPL-2.0

resource "saviynt_sftp_connection_resource" "example" {
  # Required attributes
  connection_name      = "Terraform_SFTP_Connector"
  host_name            = "hostname"
  username             = "username"
  auth_credential_type = "PASSWORD"

  # Optional attributes
  defaultsavroles       = "ROLE_ADMIN"
  email_template        = "SFTP Connection Notification"
  port_number           = "22"
  auth_credential_value = "cred_value" # Use either this or auth_credential_value_wo
  # auth_credential_value_wo = "cred_value"  # Alternative write-only credential
  passphrase = "passphrase" # Use either this or passphrase_wo
  # passphrase_wo       = "passphrase"       # Alternative write-only passphrase
  files_to_get = "*.csv,*.txt"
  files_to_put = "upload/*.json"

  wo_version = "v1.2"
}