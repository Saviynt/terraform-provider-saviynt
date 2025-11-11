// Copyright (c) 2025 Saviynt Inc.
// SPDX-License-Identifier: MPL-2.0

resource "saviynt_enterprise_roles_resource" "example" {
  # Required attributes
  role_name = "TF_Enterprise_Role_Example"
  role_type = "ENTERPRISE"
  requestor = "admin"
  owners = [
    # Role owners with different ranks
    {
      owner_name = "admin"
      rank       = "1"
    },
    {
      owner_name = "roleowner1"
      rank       = "2"
    },
    {
      owner_name = "certifier1"
      rank       = "26" # Primary Certifier
    }
  ]

  # Basic role information
  description   = "Example enterprise role created via Terraform"
  display_name  = "Terraform Enterprise Role Example"
  glossary      = "This is an example enterprise role for demonstration purposes"
  endpoint_name = "System1"

  # Time frame configuration
  default_time_frame = "168" # 7 days in hours

  # Role properties
  requestable        = "true"
  show_dynamic_attrs = "true"
  check_sod          = "true"

  # Criticality levels
  sox_critical    = "Medium"
  sys_critical    = "High"
  privileged      = "Low"
  confidentiality = "Medium"
  risk            = "Low"
  level           = "2"


  # Entitlements associated with the role
  entitlements = [
    {
      endpoint          = "System1"
      entitlement_type  = "Groups"
      entitlement_value = "HR_Group"
    },
    {
      endpoint          = "System1"
      entitlement_type  = "Groups"
      entitlement_value = "Finance_Read"
    }
  ]

  # Child roles associated with the role (available in Saviynt 25.B+)
  child_roles = [
    {
      role_name = "Junior_HR_Role"
    },
    {
      role_name = "Temp_Access_Role"
    }
  ]

  # Users assigned to the role
  users = [
    {
      user_name = "john.doe"
    },
    {
      user_name = "jane.smith"
    },
    {
      user_name = "bob.johnson"
    }
  ]

  # Custom properties for additional metadata
  custom_property1 = "Department: HR"
  custom_property2 = "Cost Center: 1001"
  custom_property3 = "Business Unit: Corporate"
  custom_property4 = "Approval Required: Yes"
  custom_property5 = "Review Frequency: Quarterly"
}
