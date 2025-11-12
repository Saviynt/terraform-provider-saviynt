// Copyright (c) 2025 Saviynt Inc.
// SPDX-License-Identifier: MPL-2.0

resource "saviynt_export_transport_package_resource" "example" {
  export_online          = "false"                                         # required - Export mode (true/false)
  export_path            = "/saviynt_shared/testexport/transportPackage"   # required - Export destination path
  update_user            = "admin"                                         # optional - User performing the export
  business_justification = "Exporting configuration for Q1 release backup" # optional - Business reason for export

  # Export SAV roles
  sav_roles = ["ROLE_ADMIN", "ROLE_USER"] # optional - List of SAV roles to export

  # Export other objects (optional)
  roles            = ["CustomRole1", "CustomRole2"]
  connections      = ["AD_Connection", "DB_Connection"]
  email_template   = ["Account Password Expiry Email", "Analytics Execution Complete Email Default Template"]
  workflows        = ["AOM Manager Approval"]
  security_systems = ["Active_Directory", "Database_System"]
  user_groups      = ["IT_Team", "HR_Team"]
  organizations    = ["IT_Department", "Finance_Department"]
}

# Example with minimal configuration (SAV roles only)
resource "saviynt_export_transport_package_resource" "minimal" {
  export_online = "false"
  export_path   = "/saviynt_shared/testexport/transportPackage"
  sav_roles     = ["ROLE_ADMIN"]
}

# Example with comprehensive export
resource "saviynt_export_transport_package_resource" "comprehensive" {
  export_online          = "false"
  export_path            = "/saviynt_shared/testexport/transportPackage"
  update_user            = "admin"
  transport_owner        = "admin"
  transport_members      = "true"
  environment_name       = "production"
  business_justification = "Full system backup before major upgrade"
  export_package_version = "2.1"

  sav_roles        = ["ROLE_ADMIN", "ROLE_USER", "ROLE_AUDITOR"]
  roles            = ["Manager", "Employee", "Contractor"]
  connections      = ["LDAP_Prod", "DB_Prod", "API_Gateway"]
  email_template   = ["Onboarding", "Offboarding", "Access_Request"]
  workflows        = ["User_Provisioning", "Access_Approval", "Termination"]
  global_config    = ["Password_Policy", "Session_Config"]
  analytics_v1     = ["User_Report", "Access_Report"]
  analytics_v2     = ["Advanced_Analytics"]
  security_systems = ["AD_Primary", "DB_Primary"]
  user_groups      = ["Admins", "Users", "Guests"]
  scan_rules       = ["SOD_Rule1", "Compliance_Rule"]
  organizations    = ["Corporate", "IT", "HR", "Finance"]
  app_onboarding   = ["App1_Config", "App2_Config"]
}
