// Copyright (c) 2025 Saviynt Inc.
// SPDX-License-Identifier: MPL-2.0

# Variables for sensitive and configurable values
variable "WORKDAY_USERNAME" {
  type        = string
  description = "Username for Workday SOAP connection"
}

variable "WORKDAY_PASSWORD" {
  type        = string
  description = "Password for Workday SOAP connection"
  sensitive   = true
}

variable "SOAP_ENDPOINT_URL" {
  type        = string
  description = "Workday SOAP endpoint URL"
}

variable "TENANT_NAME" {
  type        = string
  description = "Workday tenant name"
}

variable "VAULT_CONNECTION" {
  type        = string
  description = "Vault connection name for credential management"
  default     = ""
}

variable "VAULT_CONFIG" {
  type        = string
  description = "Vault configuration"
  default     = ""
}


resource "saviynt_workday_soap_connection_resource" "example" {
  # Required attributes
  connection_name = "Terraform_Workday_SOAP_Connector"

  # Optional base connector attributes
  description         = "Workday SOAP connection for HR data integration - ${var.TENANT_NAME}"
  default_sav_roles   = "ROLE_ADMIN"
  email_template      = "default_template"
  vault_connection    = var.VAULT_CONNECTION
  vault_configuration = var.VAULT_CONFIG
  save_in_vault       = var.VAULT_CONNECTION != "" ? "true" : "false"
  wo_version          = "v1.0"

  # Authentication - using variables for sensitive data
  username = var.WORKDAY_USERNAME
  password = var.WORKDAY_PASSWORD
  # Alternative: Use write-only password for enhanced security
  # password_wo  = var.WORKDAY_PASSWORD
  soap_endpoint = var.SOAP_ENDPOINT_URL

  # Data Import Configuration
  accounts_import_json = jsonencode({
    "call1" : {
      "call" : [
        {
          "name" : "GetWorkers",
          "connection" : "userAuth",
          "url" : var.SOAP_ENDPOINT_URL,
          "httpMethod" : "POST",
          "httpParams" : "{}",
          "httpHeaders" : {
            "Authorization" : "$${access_token}",
            "Content-Type" : "text/xml; charset=utf-8",
            "SOAPAction" : "urn:com.workday/bsvc/Human_Resources/v35.0#Get_Workers"
          },
          "httpContentType" : "text/xml",
          "successResponses" : {
            "statusCode" : [200, 201]
          },
          "requestBody" : trimspace(<<-EOF
            <soapenv:Envelope xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/" xmlns:bsvc="urn:com.workday/bsvc">
              <soapenv:Header>
                <bsvc:Workday_Common_Header>
                  <bsvc:Include_Reference_Descriptors_In_Response>true</bsvc:Include_Reference_Descriptors_In_Response>
                </bsvc:Workday_Common_Header>
              </soapenv:Header>
              <soapenv:Body>
                <bsvc:Get_Workers_Request bsvc:version="v35.0">
                  <bsvc:Request_Criteria>
                    <bsvc:Exclude_Inactive_Workers>false</bsvc:Exclude_Inactive_Workers>
                    <bsvc:Exclude_Employees>false</bsvc:Exclude_Employees>
                    <bsvc:Exclude_Contingent_Workers>false</bsvc:Exclude_Contingent_Workers>
                  </bsvc:Request_Criteria>
                  <bsvc:Response_Filter>
                    <bsvc:Page>${PAGE_NUMBER}</bsvc:Page>
                    <bsvc:Count>${PAGE_SIZE}</bsvc:Count>
                  </bsvc:Response_Filter>
                  <bsvc:Response_Group>
                    <bsvc:Include_Reference>true</bsvc:Include_Reference>
                    <bsvc:Include_Personal_Information>true</bsvc:Include_Personal_Information>
                    <bsvc:Include_Employment_Information>true</bsvc:Include_Employment_Information>
                    <bsvc:Include_Organizations>true</bsvc:Include_Organizations>
                    <bsvc:Include_Roles>true</bsvc:Include_Roles>
                  </bsvc:Response_Group>
                </bsvc:Get_Workers_Request>
              </soapenv:Body>
            </soapenv:Envelope>
          EOF
          )
        }
      ]
    }
  })

  hr_import_json = jsonencode({
    "call1" : {
      "call" : [
        {
          "name" : "GetOrganizations",
          "connection" : "userAuth",
          "url" : var.SOAP_ENDPOINT_URL,
          "httpMethod" : "POST",
          "httpHeaders" : {
            "Content-Type" : "text/xml; charset=utf-8",
            "SOAPAction" : "urn:com.workday/bsvc/Human_Resources/v35.0#Get_Organizations"
          }
        }
      ]
    }
  })

  data_to_import = "Users,Accounts,Organizations"
  page_size      = "100"
  date_format    = "yyyy-MM-dd'T'HH:mm:ss.SSS'Z'"

  # Account Management with Workday-specific SOAP operations
  create_account_json = jsonencode({
    "accountIdPath" : "$.Response_Data.Worker.Worker_Reference.ID[?(@.type=='Employee_ID')].value",
    "processingType" : "SequentialAndIterative",
    "createAccountRequestPath" : "$.call1.call[0]",
    "successResponses" : {
      "statusCode" : [200, 201]
    }
  })

  update_account_json = jsonencode({
    "accountIdPath" : "$.Response_Data.Worker.Worker_Reference.ID[?(@.type=='Employee_ID')].value",
    "processingType" : "SequentialAndIterative",
    "updateAccountRequestPath" : "$.call1.call[0]",
    "successResponses" : {
      "statusCode" : [200, 201]
    }
  })

  delete_account_json = jsonencode({
    "accountIdPath" : "$.Response_Data.Worker.Worker_Reference.ID[?(@.type=='Employee_ID')].value",
    "processingType" : "SequentialAndIterative",
    "deleteAccountRequestPath" : "$.call1.call[0]"
  })

  enable_account_json = jsonencode({
    "accountIdPath" : "$.Response_Data.Worker.Worker_Reference.ID[?(@.type=='Employee_ID')].value",
    "processingType" : "SequentialAndIterative",
    "enableAccountRequestPath" : "$.call1.call[0]"
  })

  disable_account_json = jsonencode({
    "accountIdPath" : "$.Response_Data.Worker.Worker_Reference.ID[?(@.type=='Employee_ID')].value",
    "processingType" : "SequentialAndIterative",
    "disableAccountRequestPath" : "$.call1.call[0]"
  })

  # User Management
  update_user_json = jsonencode({
    "userIdPath" : "$.Response_Data.Worker.Worker_Reference.ID[?(@.type=='Employee_ID')].value",
    "processingType" : "SequentialAndIterative",
    "updateUserRequestPath" : "$.call1.call[0]"
  })

  modify_user_data_json = jsonencode({
    "userIdPath" : "$.Response_Data.Worker.Worker_Reference.ID[?(@.type=='Employee_ID')].value",
    "processingType" : "SequentialAndIterative",
    "modifyUserDataRequestPath" : "$.call1.call[0]"
  })

  # Access Management
  grant_access_json = jsonencode({
    "accountIdPath" : "$.Response_Data.Worker.Worker_Reference.ID[?(@.type=='Employee_ID')].value",
    "processingType" : "SequentialAndIterative",
    "grantAccessRequestPath" : "$.call1.call[0]"
  })

  revoke_access_json = jsonencode({
    "accountIdPath" : "$.Response_Data.Worker.Worker_Reference.ID[?(@.type=='Employee_ID')].value",
    "processingType" : "SequentialAndIterative",
    "revokeAccessRequestPath" : "$.call1.call[0]"
  })

  # Password Management
  change_pass_json = jsonencode({
    "accountIdPath" : "$.Response_Data.Worker.Worker_Reference.ID[?(@.type=='Employee_ID')].value",
    "processingType" : "SequentialAndIterative",
    "changePasswordRequestPath" : "$.call1.call[0]"
  })
  # Alternative: Use write-only for enhanced security
  # change_pass_json_wo = jsonencode({ ... })

  password_type          = "BASIC"
  password_min_length    = "8"
  password_max_length    = "20"
  password_noofcapsalpha = "1"
  password_noofdigits    = "1"
  password_noofsplchars  = "1"

  # Response Path Configuration for Workday SOAP responses
  responsepath_userlist     = "$.Response_Data.Worker"
  responsepath_pageresults  = "$.Response_Results.Page_Results"
  responsepath_totalresults = "$.Response_Results.Total_Results"

  # Advanced Configuration
  connection_json = jsonencode({
    "authType" : "basic",
    "url" : var.SOAP_ENDPOINT_URL,
    "httpMethod" : "POST",
    "httpHeaders" : {
      "Content-Type" : "text/xml; charset=utf-8"
    },
    "tenant" : var.TENANT_NAME,
    "timeout" : 30000,
    "retryCount" : 3
  })
  # Alternative: Use write-only for enhanced security
  # connection_json_wo = jsonencode({ ... })

  custom_config = jsonencode({
    "pagination" : {
      "enabled" : true,
      "pageSize" : 100,
      "maxPages" : 1000
    },
    "workdaySpecific" : {
      "tenant" : var.TENANT_NAME,
      "apiVersion" : "v35.0",
      "includeInactiveWorkers" : false
    },
    "errorHandling" : {
      "retryOnFailure" : true,
      "maxRetries" : 3,
      "retryDelay" : 5000
    }
  })

  pam_config = jsonencode({
    "Connection" : "Workday-SOAP",
    "ConnectorClass" : "sailpoint.connector.WorkdaySOAPConnector",
    "tenant" : var.TENANT_NAME,
    "apiVersion" : "v35.0"
  })

  combined_create_request = "true"
}
