---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "saviynt_entraid_connection_resource Resource - saviynt"
subcategory: ""
description: |-
  Create and manage EntraID (AzureAD) connector in Saviynt
---

# saviynt_entraid_connection_resource (Resource)

Create and manage EntraID (AzureAD) connector in Saviynt

## Example Usage

```terraform
// Copyright (c) 2025 Saviynt Inc.
// SPDX-License-Identifier: MPL-2.0

variable "CLIENT_ID" {
  type        = string
  description = "Saviynt AzureAD CLIENT_ID"
  sensitive   = true
}

variable "CLIENT_SECRET" {
  type        = string
  description = "Saviynt AzureAD CLIENT_SECRET"
  sensitive   = true
}

variable "TENANT_ID" {
  type        = string
  description = "Saviynt AzureAD TENANT_ID"
  sensitive   = true
}

locals {
  import_user_json        = file("${path.module}/json/import_user.json")
  account_attributes      = file("${path.module}/json/account_attributes.json")
  entitlement_attribute   = file("${path.module}/json/entitlement_attribute.json")
  create_account_json     = file("${path.module}/json/create_account.json")
  add_access_json         = file("${path.module}/json/add_access.json")
  connection_json         = file("${path.module}/json/connection.json")
  status_threshold_config = file("${path.module}/json/status_threshold_config.json")
}

resource "saviynt_entraid_connection_resource" "example" {
  connection_type           = "AzureAD"
  connection_name           = "Terraform_EntraId_Connector"
  client_id                 = var.CLIENT_ID
  client_secret             = var.CLIENT_SECRET
  aad_tenant_id             = var.TENANT_ID
  authentication_endpoint   = "https://login.microsoftonline.com"
  microsoft_graph_endpoint  = "https://graph.microsoft.com"
  azure_management_endpoint = "https://management.azure.com"
  create_users              = "YES"
  create_new_endpoints      = "YES"
  managed_account_type      = "ACCOUNTS"
  import_depth              = "FINE GRAINED"

  import_user_json        = local.import_user_json
  account_attributes      = local.account_attributes
  entitlement_attribute   = local.entitlement_attribute
  create_account_json     = local.create_account_json
  add_access_json         = local.add_access_json
  connection_json         = local.connection_json
  status_threshold_config = local.status_threshold_config

  account_import_fields = "accountEnabled,mail,businessPhone,surname,givenName,displayName,userPrincipalName,id"

  update_account_json = jsonencode({
    call = [
      {
        name       = "call1",
        connection = "userAuth",
        url        = "https://graph.microsoft.com/v1.0/users/$${account.accountID}",
        httpMethod = "PATCH",
        httpParams = "{\"userprincipalname\": \"$${user.email}\"}",
        httpHeaders = {
          Authorization = "$${access_token}"
        },
        httpContentType = "application/json",
        successResponses = {
          statusCode = [
            200,
            201,
            204,
            205
          ]
        }
      }
    ]
  })

  enable_account_json = jsonencode({
    call = [{
      name       = "call1",
      connection = "userAuth",
      url        = "https://graph.microsoft.com/v1.0/users/$${account.accountID}",
      httpMethod = "PATCH",
      httpParams = "{\"accountEnabled\": true}",
      httpHeaders = {
        Authorization = "$${access_token}"
      },
      httpContentType = "application/json",
      successResponses = {
        statusCode = [200, 201, 204, 205]
      }
    }]
  })

  disable_account_json = jsonencode({
    call = [
      {
        name       = "call1",
        connection = "userAuth",
        url        = "https://graph.microsoft.com/v1.0/users/$${account.accountID}",
        httpMethod = "PATCH",
        httpParams = "{\"accountEnabled\": false}",
        httpHeaders = {
          Authorization = "$${access_token}"
        },
        httpContentType = "application/json",
        successResponses = {
          statusCode = [
            200,
            201,
            204,
            205
          ]
        }
      }
    ]
  })

  remove_access_json = jsonencode({
    call = [
      {
        name       = "SKU",
        connection = "userAuth",
        url        = "https://graph.microsoft.com/v1.0/users/$${account.accountID}/assignLicense",
        httpMethod = "POST",
        httpParams = "{\"addLicenses\": [],\"removeLicenses\": [\"$${entitlementValue.entitlementID}\"]}",
        httpHeaders = {
          Authorization = "$${access_token}"
        },
        httpContentType = "application/json",
        successResponses = {
          statusCode = [
            200,
            201,
            204,
            205
          ]
        }
      },
      {
        name       = "DirectoryRole",
        connection = "userAuth",
        url        = "https://graph.microsoft.com/v1.0/directoryRoles/$${entitlementValue.entitlementID}/members/$${account.accountID}/\\$ref",
        httpMethod = "DELETE",
        httpHeaders = {
          Authorization = "$${access_token}"
        },
        httpContentType = "application/json",
        successResponses = {
          statusCode = [
            200,
            201,
            204,
            205
          ]
        }
      },
      {
        name       = "AADGroup",
        connection = "userAuth",
        url        = "https://graph.microsoft.com/v1.0/groups/$${entitlementValue.entitlementID}/members/$${account.accountID}/\\$ref",
        httpMethod = "DELETE",
        httpHeaders = {
          Authorization = "$${access_token}"
        },
        httpContentType = "application/json",
        successResponses = {
          statusCode = [
            200,
            201,
            204,
            205
          ]
        }
      },
      {
        name       = "ApplicationInstance",
        connection = "entAuth",
        url        = "https://graph.windows.net/myOrganization/servicePrincipals/$${entitlementValue.entitlementID}/appRoleAssignedTo?api-version=1.6&\\$top=999",
        httpMethod = "GET",
        httpHeaders = {
          Authorization = "$${access_token}",
          Accept        = "application/json"
        },
        httpContentType = "application/json",
        successResponses = {
          statusCode = [
            200,
            201
          ]
        }
      },
      {
        name       = "ApplicationInstance",
        connection = "entAuth",
        url        = "https://graph.windows.net/myOrganization/servicePrincipals/$${entitlementValue.entitlementID}/appRoleAssignedTo/$${for (Map map : response.ApplicationInstance1.message.value){if (map.principalId.toString().equals(account.accountID)){return map.objectId;}}}?api-version=1.6",
        httpMethod = "DELETE",
        httpHeaders = {
          Authorization = "$${access_token}",
          Accept        = "application/json"
        },
        httpContentType = "application/json",
        successResponses = {
          statusCode = [
            200,
            201,
            204,
            205
          ]
        }
      }
    ]
  })

  update_user_json = jsonencode({
    actions = {
      "Disable User" = {
        call = [
          {
            callOrder       = 0,
            connection      = "$${connectionName}",
            httpContentType = "application/json",
            httpHeaders = {
              Authorization = "$${access_token}"
            },
            httpMethod = "PATCH",
            httpParams = "{\"accountEnabled\": false}",
            name       = "Disable User",
            successResponses = {
              statusCode = [
                200,
                204
              ]
            },
            url = "https://graph.microsoft.com/v1.0/users/$${user.customproperty4}"
          }
        ]
      },
      "Enable User" = {
        call = [
          {
            callOrder       = 0,
            connection      = "$${connectionName}",
            httpContentType = "application/json",
            httpHeaders = {
              Authorization = "$${access_token}"
            },
            httpMethod = "PATCH",
            httpParams = "{\"accountEnabled\": true}",
            name       = "Enable User",
            successResponses = {
              statusCode = [
                200,
                204
              ]
            },
            url = "https://graph.microsoft.com/v1.0/users/$${user.customproperty4}"
          }
        ]
      },
      "Update Login" = {
        call = [
          {
            callOrder       = 0,
            connection      = "$${connectionName}",
            httpContentType = "application/json",
            httpHeaders = {
              Authorization = "$${access_token}"
            },
            httpMethod = "PATCH",
            httpParams = "{\"mobilePhone\": \"$${user.phonenumber}\"}",
            name       = "Update Login",
            successResponses = {
              statusCode = [
                200,
                204
              ]
            },
            url = "https://graph.microsoft.com/v1.0/users/$${user.customproperty4}"
          }
        ]
      }
    }
  })

  change_pass_json = jsonencode({
    call = [
      {
        name       = "call1",
        connection = "$${connectionName}",
        url        = "https://graph.microsoft.com/v1.0/users/$${account.accountID}",
        httpMethod = "PATCH",
        httpParams = "{\"passwordPolicies\" :\"DisableStrongPassword\",\"passwordProfile\" : {\"password\":\"$${password}\",\"forceChangePasswordNextSignIn\": false}}",
        httpHeaders = {
          Authorization = "$${access_token}"
        },
        httpContentType = "application/json",
        successResponses = {
          statusCode = [
            200,
            201,
            204,
            205
          ]
        }
      }
    ]
  })

  remove_account_json = jsonencode({
    call = [
      {
        name       = "call1",
        connection = "userAuth",
        url        = "https://graph.microsoft.com/v1.0/users/$${account.accountID}",
        httpMethod = "DELETE",
        httpHeaders = {
          Authorization = "$${access_token}"
        },
        httpContentType = "application/json",
        successResponses = {
          statusCode = [
            200,
            201,
            204,
            205
          ]
        }
      }
    ]
  })

  endpoints_filter = jsonencode({
    APPLICATION_DEV = [
      {
        AADGROUP = [
          "GROUP_IN_ENGG",
          "GROUP_IN_FINANCE",
          "GROUP_IN_MARKETTING"
        ]
      }
    ]
  })

  accounts_filter = "(userType%20eq%20%27Member%27%20and%20(employeeType%20eq%20%27Employee%27%20or%20employeeType%20eq%20%27External%27%20or%20employeeType%20eq%20%27AdminAccount%27%20or%20employeeType%20eq%20%27Frontline%27)"

  config_json = jsonencode({
    connectionTimeoutConfig = {
      connectionTimeout = 10,
      readTimeout       = 60,
      writeTimeout      = 60,
      retryWait         = 5,
      retryCount        = 3
    }
  })

  windows_connector_json = jsonencode({
    http = {
      url = "http://<domain-name>/FIMAzure/PS/ExecutePSCommand",
      httpHeaders = {
        Content-Type  = "application/json",
        Authorization = "<authorization>",
        Username      = "<username>",
        Password      = "<encrypted password>",
        Command       = "<command>"
      },
      httpContentType = "application/json",
      httpMethod      = "GET"
    },
    listField = "value",
    keyField  = "accountID",
    colsToPropsMap = {
      accountID        = "ObjectId~#~char",
      customproperty17 = "WhenCreated~#~char",
      customproperty15 = "MFAState~#~char",
      customproperty16 = "MFADateTime~#~char"
    }
  })

  service_account_attributes = jsonencode({
    colsToPropsMap = {
      accountID        = "id~#~char",
      name             = "displayName~#~char",
      displayName      = "displayName~#~char",
      customproperty1  = "servicePrincipalNames~#~char",
      customproperty2  = "appId~#~char",
      status           = "accountEnabled~#~char",
      customproperty10 = "accountEnabled~#~char",
      customproperty3  = "appOwnerOrganizationId~#~char",
      customproperty4  = "appDescription~#~char",
      customproperty5  = "appDisplayName~#~char",
      customproperty6  = "accountEnabled~#~char",
      customproperty7  = "homepage~#~char",
      accountType      = "#CONST#Service Principal~#~char",
      accountClass     = "servicePrincipalType~#~char"
    }
  })
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `aad_tenant_id` (String) Azure Active Directory tenant ID.
- `client_id` (String) Client ID for authentication.
- `client_secret` (String) Client Secret for authentication.
- `connection_name` (String) Name of the connection. Example: "Active Directory_Doc"

### Optional

- `access_token` (String) Access token used for API calls.
- `account_attributes` (String) Attributes for account configuration.
- `account_import_fields` (String) Fields to import for accounts.
- `accounts_filter` (String) Filter for accounts.
- `add_access_json` (String) JSON template to add access.
- `authentication_endpoint` (String) Authentication endpoint URL.
- `azure_management_endpoint` (String) Azure management endpoint URL.
- `azure_mgmt_access_token` (String) Access token for Azure management APIs.
- `change_pass_json` (String) JSON template to change password.
- `config_json` (String) Main config JSON.
- `connection_json` (String) Connection JSON configuration.
- `connection_type` (String) Connection type (e.g., 'AD' for Active Directory). Example: "AD"
- `create_account_json` (String) JSON template to create an account.
- `create_channel_json` (String) JSON to create channel.
- `create_group_json` (String) JSON to create group.
- `create_new_endpoints` (String) Configuration to create new endpoints.
- `create_team_json` (String) JSON to create team.
- `create_users` (String) Flag or configuration for creating users.
- `defaultsavroles` (String) Default SAV roles for managing the connection. Example: "ROLE_ORG"
- `delete_group_json` (String) JSON to delete group.
- `delta_tokens_json` (String) Delta tokens JSON data.
- `disable_account_json` (String) JSON template to disable an account.
- `email_template` (String) Email template for notifications. Example: "New Account Task Creation"
- `enable_account_json` (String) JSON template to enable an account.
- `endpoints_filter` (String) Endpoints filter configuration.
- `entitlement_attribute` (String) Attribute used for entitlement.
- `entitlement_filter_json` (String) Filter JSON for entitlements.
- `error_code` (String) An error code where '0' signifies success and '1' signifies an unsuccessful operation.
- `import_user_json` (String) JSON configuration for importing users.
- `microsoft_graph_endpoint` (String) Microsoft Graph API endpoint.
- `modify_user_data_json` (String) JSON to modify user data.
- `msg` (String) A message indicating the outcome of the operation.
- `pam_config` (String) PAM configuration.
- `remove_access_json` (String) JSON template to remove access.
- `remove_account_json` (String) JSON template to remove account.
- `save_in_vault` (String) Flag indicating whether the encrypted attribute should be saved in the configured vault. Example: "false"
- `status_threshold_config` (String) Configuration for status thresholds.
- `update_account_json` (String) JSON template to update an account.
- `update_group_json` (String) JSON to update group.
- `update_user_json` (String) JSON template to update user.
- `vault_configuration` (String) JSON string specifying vault configuration. Example: '{"path":"/secrets/data/kv-dev-intgn1/-AD_Credential","keyMapping":{"PASSWORD":"AD_PASSWORD~#~None"}}'
- `vault_connection` (String) Specifies the type of vault connection being used (e.g., 'Hashicorp'). Example: "Hashicorp"
- `windows_connector_json` (String) Windows connector JSON configuration.

### Read-Only

- `connection_key` (Number) Unique identifier of the connection returned by the API. Example: 1909
- `id` (String) Resource ID.
