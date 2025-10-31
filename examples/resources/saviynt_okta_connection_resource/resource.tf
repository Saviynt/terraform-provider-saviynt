// Copyright (c) 2025 Saviynt Inc.
// SPDX-License-Identifier: MPL-2.0

variable "OKTA_DOMAIN" {
  type        = string
  description = "Okta domain URL (e.g., https://dev-123456.okta.com)"
}

variable "OKTA_API_TOKEN" {
  type        = string
  description = "Okta API token for authentication"
  sensitive   = true
}

variable "OKTA_APP_SECURITY_SYSTEM" {
  type        = string
  description = "Saviynt security system name for Okta applications"
}

variable "VAULT_CONNECTION" {
  type        = string
  description = "Vault connection name"
  default     = ""
}

variable "VAULT_CONFIG" {
  type        = string
  description = "Vault configuration"
  default     = ""
}

variable "SAVE_IN_VAULT" {
  type        = string
  description = "Save credentials in vault (true/false)"
  default     = ""
}

resource "saviynt_okta_connection_resource" "example" {
  connection_name = "Terraform_Okta_Connector"

  # Required fields
  import_url = var.OKTA_DOMAIN
  auth_token = var.OKTA_API_TOKEN

  //Specify the name of the security system, which you want to automatically create. It is mandatory to specify a different name from the manually created connection.
  okta_application_securitysystem = var.OKTA_APP_SECURITY_SYSTEM

  # Optional vault configuration
  vault_connection    = var.VAULT_CONNECTION
  vault_configuration = var.VAULT_CONFIG
  save_in_vault       = var.SAVE_IN_VAULT

  # Account field mappings - Specify this parameter to map account attributes of Okta to account attributes of Saviynt Identity Cloud for account import
  account_field_mappings = jsonencode({
    "customproperty1" = "firstName"
    "customproperty2" = "lastName"
    "customproperty3" = "email"
    "customproperty4" = "secondEmail"
    "customproperty5" = "mobilePhone"
    "customproperty6" = "status"
  })

  # User field mappings - Specify this parameter to map user attributes of Okta to user attributes of Saviynt Identity Cloud for user import
  user_field_mappings = jsonencode({
    "customproperty11"  = "profile.department"
    "customproperty12"  = "profile.contractEndDate"
    "customproperty13"  = "profile.oktaId"
    "employeeType"      = "profile.userType"
    "job_function"      = "profile.businessUnit"
    "location"          = "profile.location"
    "preferedFirstName" = "profile.knownAs"
  })

  # Entitlement types mappings - Specify the mapping between application objects and Saviynt Identity Cloud objects
  entitlement_types_mappings = "Groups=groups"

  # Import inactive apps - Set this parameter to TRUE, if you want to import the inactive Okta applications. 
  import_inactive_apps = "TRUE"

  # Okta groups filter - Specify this parameter to pull the Okta groups in the following format.
  okta_groups_filter = "type eq \"OKTA_GROUP\""

  # App account field mappings - Specify this parameter to map application account attributes of Okta to endpoint account attributes of Saviynt Identity Cloud for application account import.
  app_account_field_mappings = jsonencode({
    "defaultMapping" = {
      "customproperty3"  = "created"
      "customproperty8"  = "lastUpdated"
      "customproperty9"  = "scope"
      "customproperty10" = "status"
      "customproperty11" = "statusChanged"
      "customproperty14" = "syncState"
    }
    "boxnet(BoxNew)" = {
      "customproperty20" = "firstName"
      "customproperty21" = "lastName"
      "customproperty22" = "role"
      "customproperty23" = "email"
    }
  })

  # Status threshold configuration - Specify the account attribute mapped with the account status along with the values to be considered for imported accounts in the
  status_threshold_config = jsonencode({
    "statusAndThresholdConfig" = {
      "accountThresholdValue"       = 100
      "deleteLinks"                 = true
      "correlateInactiveAccounts"   = true
      "accountsNotInImportAction"   = "SUSPEND"
      "inactivateAccountsNotInFile" = false
      "statusColumn"                = "customproperty6"
      "activeStatus" = [
        "ACTIVE",
        "STAGED",
        "PROVISIONED",
        "RECOVERY",
        "LOCKED_OUT",
        "PASSWORD_EXPIRED"
      ]
    }
  })

  # Audit filter - Specify the parameter to filter the object as per requirement
  audit_filter = "eventType eq \"user.authentication.sso\" and target.alternateId eq \"Wayfair\""

  # Activate endpoint - Enable or disable the endpoint based on the application status received during import job run. Select True to enable or False to disable as per your requirement.
  activate_endpoint = "true"

  # Config Json - Specify the maximum percentage of API requests that can be allowed, and configurations to pause the import and retry the connection after the wait time.
  config_json = jsonencode({
    "apiRateLimitConfig" = {
      "maxApiCapacityPercentage" = 70
      "maxRefreshTryCount"       = 5
      "retryWaitSeconds"         = 8
    }
  })
}