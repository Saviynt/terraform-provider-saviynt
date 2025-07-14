// Copyright (c) Saviynt Inc.
// SPDX-License-Identifier: MPL-2.0

resource "saviynt_security_system_resource" "example" {
  systemname                      = "Terraform_Security_System"
  display_name                    = "Terraform_Security_System"
  hostname                        = "EntitlementsOnly"
  port                            = "443"
  access_add_workflow             = "autoapprovalwf"
  access_remove_workflow          = "autoapprovalwf"
  add_service_account_workflow    = "autoapprovalwf"
  remove_service_account_workflow = "autoapprovalwf"
  automated_provisioning          = "true"
  use_open_connector              = "true"
  recon_application               = "true"
  instant_provision               = "true"
  provisioning_tries              = "3"
  provisioning_comments           = "Auto-provisioned by Terraform"
}