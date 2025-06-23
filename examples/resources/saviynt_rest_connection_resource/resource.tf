// Copyright (c) Saviynt Inc.
// SPDX-License-Identifier: MPL-2.0

resource "saviynt_rest_connection_resource" "example" {
  connection_type = "REST"
  connection_name = "Terraform_Rest_Connector"
  connection_json = jsonencode({
    authentications = {
      acctAuth = {
        authType        = "Basic",
        url             = "@HOSTNAME@",
        httpMethod      = "POST",
        httpParams      = {},
        httpHeaders     = {},
        httpContentType = "text/html",
        properties = {
          userName = "@USERNAME@",
          password = "@PASSWORD@"
        },
        expiryError = "ExpiredAuthenticationToken",
        authError = [
          "InvalidAuthenticationToken",
          "AuthenticationFailed"
        ],
        timeOutError       = "Read timed out",
        errorPath          = "error.code",
        maxRefreshTryCount = 5,
        tokenResponsePath  = "access_token",
        tokenType          = "Basic",
        accessToken        = "Basic bmlzaGFyLmJhYnVAc2"
      }
    }
  })
  import_user_json = jsonencode({
    connection = "acctAuth",
    successResponses = {
      statusCode = [
        200,
        201,
        202,
        203,
        204,
        205
      ]
    },
    url        = "https://abc.zendesk.com/api//users.json",
    httpMethod = "GET",
    httpHeaders = {
      Authorization = "$${access_token}"
    },
    userResponsePath = "users",
    colsToPropsMap = {
      username       = "id~#~char",
      systemUserName = "id~#~char",
      displayname    = "name~#~char",
      email          = "email~#~char"
    }
  })
  import_account_ent_json = jsonencode({
    authentications = {
      userAuth = {
        authType   = "oauth2",
        url        = "https://<domain name>/api/v18.2/auth",
        httpMethod = "POST",
        httpParams = {
          username = "<Username>",
          password = "<PASSWORD>"
        },
        httpHeaders = {
          contentType = "application/x-www-form-urlencoded"
        },
        httpContentType = "application/x-www-form-urlencoded",
        expiryError     = "ExpiredAuthenticationToken",
        authError = [
          "InvalidAuthenticationToken",
          "AuthenticationFailed",
          "FAILURE",
          "INVALID_SESSION_ID"
        ],
        timeOutError       = "Read timed out",
        errorPath          = "errors.type",
        maxRefreshTryCount = 5,
        tokenResponsePath  = "sessionId",
        tokenType          = "Bearer",
        accessToken        = "<access token>"
      }
    }
  })

  status_threshold_config = jsonencode({
    statusAndThresholdConfig = {
      accountThresholdValue       = 50,
      correlateInactiveAccounts   = true,
      inactivateAccountsNotInFile = false,
      statusColumn                = "customproperty30",
      activeStatus = [
        "ENABLE",
        "PROVISIONED"
      ],
      deleteLinks             = true,
      inactivateEntsNotInFeed = true,
      entThresholdValue = {
        entType = {
          Group = {
            ent = 100
          },
          Role = {
            ent = 100
          }
        }
      }
    }
  })

  create_account_json = jsonencode({
    accountIdPath = "call1.message.user.id",
    dateFormat    = "yyyy-MM-dd'T'HH:mm:ssXXX",
    responseColsToPropsMap = {
      displayname = "call1.message.user.name~#~char"
    },
    call = [
      {
        name       = "call1",
        connection = "acctAuth",
        url        = "@HOSTNAME@/api/v2/users",
        httpMethod = "POST",
        httpParams = "{\"user\": {\"name\": \"$${user.firstname} $${user.lastname}\", \"email\": \"$${user.email}\", \"role\":\"agent\"}}",
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
      }
    ]
  })

  update_account_json = jsonencode({
    dateFormat = "yyyy-MM-dd'T'HH:mm:ssXXX",
    responseColsToPropsMap = {
      displayName = "call1.message.user.name~#~char"
    },
    call = [
      {
        name       = "Role",
        connection = "acctAuth",
        url        = "@HOSTNAME@/api/v2/users/$${account.accountID}",
        httpMethod = "PUT",
        httpParams = "{\"user\": {\"name\": \"$${user.firstname} $${user.lastname}\"}}",
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
      }
    ]
  })

  enable_account_json = jsonencode({
    call = [
      {
        name       = "call1",
        connection = "acctAuth",
        url        = "@HOSTNAME@/api/v2/users",
        httpMethod = "PUT",
        httpParams = "{\"user\":{\"suspended\": \"false\"}}",
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
      }
    ]
  })

  disable_account_json = jsonencode({
    call = [
      {
        name       = "call1",
        connection = "acctAuth",
        url        = "@HOSTNAME@/api/v2/users",
        httpMethod = "PUT",
        httpParams = "{\"user\":{\"suspended\": \"true\"}}",
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
      }
    ]
  })

  add_access_json = jsonencode({
    call = [
      {
        name       = "Group",
        connection = "acctAuth",
        url        = "@HOSTNAME@/api/v2/group_memberships",
        httpMethod = "POST",
        httpParams = "{\"group_membership\": {\"user_id\": \"$${account.accountID}\", \"group_id\": \"$${entitlementValue.entitlementID}\"}}",
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
        name       = "Role",
        connection = "acctAuth",
        url        = "@HOSTNAME@/api/v2/users/$${account.accountID}",
        httpMethod = "PUT",
        httpParams = "{\"user\": {\"custom_role_id\": $${entitlementValue.entitlementID}}}",
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
      }
    ]
  })

  remove_access_json = jsonencode({
    call = [
      {
        name       = "Group",
        connection = "acctAuth",
        url        = "@HOSTNAME@/api/v2/users/$${account.accountID}/group_memberships",
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
        name       = "Group",
        connection = "acctAuth",
        url        = "@HOSTNAME@/api/v2/group_memberships/$${for (Map map : response.Group1.message.group_memberships){if (map.group_id.toString().equals(entitlementValue.entitlementID)){return map.id;}}}",
        httpMethod = "DELETE",
        httpHeaders = {
          Authorization = "$${access_token}",
          Accept        = "application/json"
        },
        httpContentType = "application/json",
        successResponses = {
          statusCode = [
            200,
            204
          ]
        }
      },
      {
        name       = "Role",
        connection = "acctAuth",
        url        = "@HOSTNAME@/api/v2/users/$${account.accountID}",
        httpMethod = "PUT",
        httpParams = "{\"user\": {\"custom_role_id\": $${entitlementValue.entitlementID}}}",
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
      }
    ]
  })

  update_user_json = jsonencode({
    actions = {
      "Update Login" = {
        call = [
          {
            name       = "Update Login",
            callOrder  = 0,
            connection = "userAuth",
            url        = "https://<domain name>/odata/v2/User('$${user.employeeid}')",
            httpMethod = "POST",
            httpParams = "{\"__metadata\":{\"uri\":\"PerEmail(emailType='920',personIdExternal='$${user.employeeid}')\"},\"emailAddress\":\"$${user.email}\",\"isPrimary\":true}",
            httpHeaders = {
              x-http-method = "MERGE",
              Accept        = "application/json",
              Authorization = "$${access_token}"
            },
            httpContentType = "application/json",
            successResponses = {
              "statusCode" : [
                200,
                201
              ]
            }
          }
        ]
      }
    }
  })

  change_pass_json = jsonencode({
    call = [
      {
        name       = "call1",
        connection = "acctAuth",
        url        = "@HOSTNAME@/api/v2/users/$${account.accountID}/password.json",
        httpMethod = "POST",
        httpParams = "{\"password\": \"$${password}\"}",
        httpHeaders = {
          Authorization = "$${access_token}",
          Accept        = "application/json"
        },
        httpContentType = "application/json",
        successResponses = {
          statusCode = [
            200,
            201,
            204
          ]
        }
      }
    ]
  })

  remove_account_json = jsonencode({
    call = [
      {
        name       = "call1",
        connection = "acctAuth",
        url        = "@HOSTNAME@/api/v2/users/$${account.accountID}",
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
            204
          ]
        }
      }
    ]
  })

  ticket_status_json = jsonencode({
    call = [
      {
        name       = "call1",
        connection = "userAuth",
        url        = "<https://<domain-name>/api/now/table/sc_req_item?>  sysparm_query=request.number=$${ticketID}&sysparm_limit=1&sysparm_display_value=true",
        httpMethod = "GET",
        httpHeaders = {
          Authorization = "$${access_token}"
        },
        httpContentType  = "application/json",
        ticketStatusPath = "result[0].state",
        ticketStatusValue = [
          "Open",
          "OPEN",
          "open"
        ],
        successResponses = [
          {}
        ]
      }
    ]
  })

  create_ticket_json = jsonencode({
    call = [
      {
        name       = "call1",
        connection = "userAuth",
        url        = "https://<domain-name>/SaviyntRequest.do?Action=create_ritm",
        httpMethod = "POST",
        httpParams = "{\"bp_id\":\"$${user.username}\",\"name\":\"$${user.lastname},$${user.firstname\",\"$${taskIds}\",\"email\":\"$${user.email}\",\"permissions\": \"$${allEntitlementsValues}\"}",
        httpHeaders = {
          Authorization = "$${access_token}"
        },
        httpContentType = "application/json",
        ticketidPath    = "Request Number",
        unsuccessResponses = {
          message = "Failed"
        }
      }
    ]
  })

  endpoints_filter = jsonencode({
    "<endpoint_name>" : [
      {
        "<entitlementType>" : [
          "entitlementValue1",
          "entitlementValue2"
        ]
      }
    ]
    }
  )

  passwd_policy_json = jsonencode({
    minLength     = 8,
    maxLength     = 20,
    noOfCAPSAlpha = 3,
    noOfDigits    = 2,
    noOfSplChars  = 3
  })

  config_json = jsonencode({
    showLogs = false,
    provisioningLimit = {
      types = {
        NEWACCOUNT     = 100,
        UPDATEACCOUNT  = 100,
        ENABLEACCOUNT  = 100,
        DISABLEACCOUNT = 100,
        REMOVEACCOUNT  = 100,
        ADDACCESS      = 100,
        REMOVEACCESS   = 100,
        UPDATEUSER     = 100,
        CHANGEPASSWORD = 100
      },
      waitTimeMillis = 1000
    }
  })

  add_ffid_access_json = jsonencode({
    call = [
      {
        name       = "call1",
        connection = "acctAuth",
        url        = "<url>",
        httpMethod = "POST",
        httpParams = "{\"username\": \"vendormanager\", \"password\": \"$${arsTasks.password}\", \"changePasswordAssociatedAccounts\":\"false\", \"endpoint\":\"ABC\", \"validateagainstpolicy\":\"N\", \"updateUserPassword\":\"true\"}",
        httpHeaders = {
          Authorization = "$${access_token}"
        },
        httpContentType = "application/x-www-form-urlencoded",
        successResponses = {
          statusCode = [
            200,
            201
          ]
        }
      }
    ]
  })

  remove_ffid_access_json = jsonencode({
    call = [
      {
        name       = "call1",
        connection = "acctAuth",
        url        = "<url>",
        httpMethod = "POST",
        httpParams = "{\"username\": \"vendormanager\", \"password\": \"$${arsTasks.password}\", \"changePasswordAssociatedAccounts\":\"false\", \"endpoint\":\"ABC\", \"validateagainstpolicy\":\"N\", \"updateUserPassword\":\"true\"}",
        httpHeaders = {
          Authorization = "$${access_token}"
        },
        httpContentType = "application/x-www-form-urlencoded",
        successResponses = {
          statusCode = [
            200,
            201
          ]
        }
      }
    ]
  })

  send_otp_json = jsonencode({
    call = [
      {
        name       = "call1",
        connection = "acctAuth",
        url        = "<URL>",
        httpMethod = "POST",
        httpHeaders = {
          Accept       = "application/json",
          Content-Type = "application/json"
        },
        httpContentType = "application/json",
        httpParams = {
          senderid = "<SENDERID>",
          message  = "Your Saviynt verification code is $${otp}. It will be valid for the next 300 seconds.",
          phone    = "$${phone}"
        }
      }
    ]
  })

  validate_otp_json = jsonencode({
    call = [
      {
        name       = "call1",
        connection = "acctAuth",
        url        = "<URL>",
        httpHeaders = {
          Accept       = "application/json",
          Content-Type = "application/json"
        },
        httpContentType = "application/json",
        httpParams      = "{\"passCode\": \"$${passCode}\"}",
        unsuccessResponses = {
        statusCode = [403] },
        successResponses = {
          statusCode = [200]
        }
      }
    ]
  })
}
