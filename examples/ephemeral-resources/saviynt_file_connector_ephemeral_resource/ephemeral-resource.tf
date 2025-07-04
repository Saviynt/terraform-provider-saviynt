// Copyright (c) Saviynt Inc.

ephemeral "saviynt_file_connector_ephemeral_resource" "example" {
  file_path = "creds.json"
}

# Can be used as below:
resource "saviynt_ad_connection_resource" "example" {
  connection_type = "AD"
  connection_name = "Terraform_AD_Connector"
  url             = format("%s://%s:%d", var.LDAP_PROTOCOL, var.IP_ADDRESS, var.LDAP_PORT)
  password        = ephemeral.saviynt_file_connector_ephemeral_resource.example.password
  username        = ephemeral.saviynt_file_connector_ephemeral_resource.example.username
}