/*
 * Copyright (c) 2025 Saviynt Inc.
 * All Rights Reserved.
 *
 * This software is the confidential and proprietary information of
 * Saviynt Inc. ("Confidential Information"). You shall not disclose,
 * use, or distribute such Confidential Information except in accordance
 * with the terms of the license agreement you entered into with Saviynt.
 *
 * SAVIYNT MAKES NO REPRESENTATIONS OR WARRANTIES ABOUT THE SUITABILITY OF
 * THE SOFTWARE, EITHER EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO
 * THE IMPLIED WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR
 * PURPOSE, OR NON-INFRINGEMENT.
 */

variable "ORG_LIST" {
  type = string
}
resource "saviynt_github_rest_connection_resource" "example" {
  connection_type = "GithubRest"
  connection_name = "Terraform_Github_Connection"
  connection_json = jsonencode({
    authentications = {
      acctAuth = {
        authType   = "oauth2",
        url        = "https://<DOMAIN>/login/oauth/access_token",
        httpMethod = "POST",
        httpHeaders = {
          contentType = "application/x-www-form-urlencoded"
        },
        httpContentType = "application/x-www-form-urlencoded",
        expiryError     = "ExpiredAuthenticationToken",
        authError = [
          "InvalidAuthenticationToken",
          "AuthenticationFailed"
        ],
        timeOutError       = "Read timed out",
        errorPath          = "error.code",
        maxRefreshTryCount = 5,
        tokenResponsePath  = "access_token",
        tokenType          = "Bearer",
        accessToken        = "<TOKEN>"
      }
    }
  })
  import_account_ent_json = jsonencode({
    accountParams = {
      connection     = "acctAuth"
      createUsers    = false
      adminName      = "admin"
      processingType = "SequentialAndIterative"
      call = {
        call1 = {
          callOrder   = 0
          stageNumber = 0
          http = {
            url             = "https://<DOMAIN>/orgs/<Org_Name>/members?per_page=100&page=1"
            httpHeaders     = { Authorization = "$${access_token}" }
            httpContentType = "application/x-www-form-urlencoded"
            httpMethod      = "GET"
          }
          listField = ""
          keyField  = "accountID"
          colsToPropsMap = {
            accountID   = "id~#~char"
            name        = "login~#~char"
            displayName = "login~#~char"
          }
          makeProcessingStatus   = true
          disableDeletedAccounts = true
          pagination = {
            nextUrl = {
              nextUrlPath = "$$${headers?.Link == null ? \"\" : headers?.Link?.contains(\"next\") ? headers?.Link?.split(\",\")?.size() == 2 ? headers?.Link?.split(\",\")[0]?.replace(\"<\", \"\")?.replace(\">; rel=\\\"next\\\"\", \"\")?.trim() : headers?.Link?.split(\",\")[1]?.replace(\"<\", \"\")?.replace(\">; rel=\\\"next\\\"\", \"\")?.trim() : \"\"}"
            }
          }
        }
        call2 = {
          callOrder   = 1
          stageNumber = 1
          http = {
            url             = "https://<DOMAIN>/users/$${accountName}"
            httpHeaders     = { Authorization = "$${access_token}" }
            httpContentType = "application/x-www-form-urlencoded"
            httpMethod      = "GET"
          }
          listField = ""
          keyField  = "accountID"
          colsToPropsMap = {
            accountID        = "id~#~char"
            name             = "login~#~char"
            displayName      = "login~#~char"
            customproperty1  = "avatar_url~#~char"
            customproperty2  = "gravatar_id~#~char"
            customproperty3  = "url~#~char"
            customproperty10 = "suspended_at~#~char"
          }
          multiTrigger = {
            multiTriggerType = "MultiTriggerByAccountBatching"
            triggersCount    = "20"
            params           = { accountsoraccess = "accounts" }
          }
        }
      }
    }
    entitlementParams = {
      connection                = "acctAuth"
      processingType            = "SequentialAndIterative"
      supportedEntitlementTypes = ["Organization", "Team", "Repository"]
      entTypes = {
        Organization = {
          entTypeOrder = 0
          call = {
            call1 = {
              callOrder   = 0
              stageNumber = 3
              inputParams = { statusFilter = "(ev.status = 6 or ev.status = 1)" }
              http = {
                httpHeaders     = { Authorization = "$${access_token}" }
                url             = "https://<DOMAIN>/orgs/$${entitlementValue}"
                httpContentType = "application/x-www-form-urlencoded"
                httpMethod      = "GET"
              }
              listField = ""
              keyField  = "entitlementID"
              colsToPropsMap = {
                entitlementID     = "id~#~char"
                entitlement_value = "login~#~char"
                customproperty6   = "url~#~char"
                customproperty7   = "description~#~char"
                customproperty8   = "public_repos~#~char"
                customproperty9   = "public_gists~#~char"
                customproperty10  = "followers~#~char"
              }
            }
            call2 = {
              callOrder   = 1
              dummyCall   = true
              stageNumber = 4
              inputParams = { statusFilter = "(ev.status = 6 or ev.status = 1)" }
              multiTrigger = {
                multiTriggerType = "MultiTriggerByEntitlementBatching"
                triggersCount    = "20"
                params           = { accountsoraccess = "access" }
              }
            }
          }
        }
        Team = {
          entTypeOrder = 1
          entitlementOwnerConfig = {
            maxNumberOfOwner  = 4
            typeOfImportOwner = ["maintainer"]
          }
          call = {
            call1 = {
              callOrder   = 0
              stageNumber = 6
              inputParams = { entitlementname = "Organization" }
              http = {
                httpHeaders     = { Authorization = "$${access_token}" }
                url             = "https://<DOMAIN>/orgs/$${entitlementValue}/teams?per_page=100&page=1"
                httpContentType = "application/x-www-form-urlencoded"
                httpMethod      = "GET"
              }
              pagination = {
                nextUrl = {
                  nextUrlPath = "$$${headers?.Link == null ? '' : headers?.Link?.contains('next') ? (headers?.Link?.split(',')?.size() == 2 ? headers?.Link?.split(',')[0]?.replace('<', '')?.replace('>; rel=\\\"next\\\"', '')?.trim() : headers?.Link?.split(',')[1]?.replace('<', '')?.replace('>; rel=\\\"next\\\"', '')?.trim()) : ''}"
                }
              }
              listField = ""
              keyField  = "entitlementID"
              colsToPropsMap = {
                entitlementID          = "id~#~char"
                entitlement_value      = "name~#~char"
                customproperty1        = "url~#~char"
                customproperty6        = "privacy~#~char"
                customproperty7        = "permission~#~char"
                customproperty20       = "STORE#ENT#MAPPINGINFO#PARENTID#TYPE##ENTMAP~#~char"
                entitlementMappingJson = "STORE#ENT#MAPPINGINFO#PARENTID~#~char"
              }
            }
          }
        }
        Repository = {
          entTypeOrder = 2
          call = {
            call1 = {
              callOrder   = 0
              stageNumber = 7
              inputParams = {
                entitlementname = "Organization"
                statusFilter    = "(ev.status = 6 or ev.status = 1)"
              }
              http = {
                httpHeaders     = { Authorization = "$${access_token}" }
                url             = "https://<DOMAIN>/orgs/$${entitlementValue}/repos?per_page=100&page=1"
                httpContentType = "application/x-www-form-urlencoded"
                httpMethod      = "GET"
              }
              pagination = {
                nextUrl = {
                  nextUrlPath = "$$${headers?.Link == null ? '' : headers?.Link?.contains('next') ? (headers?.Link?.split(',')?.size() == 2 ? headers?.Link?.split(',')[0]?.replace('<', '')?.replace('>; rel=\\\"next\\\"', '')?.trim() : headers?.Link?.split(',')[1]?.replace('<', '')?.replace('>; rel=\\\"next\\\"', '')?.trim()) : ''}"
                }
              }
              listField = ""
              keyField  = "entitlementID"
              colsToPropsMap = {
                entitlementID          = "id~#~char"
                entitlement_value      = "name~#~char"
                customproperty1        = "url~#~char"
                customproperty6        = "private~#~char"
                customproperty7        = "size~#~char"
                customproperty8        = "permissions.admin~#~char"
                customproperty9        = "permissions.push~#~char"
                customproperty10       = "permissions.pull~#~char"
                customproperty11       = "owner.login~#~char"
                customproperty20       = "STORE#ENT#MAPPINGINFO#PARENTID#TYPE##ENTMAP~#~char"
                entitlementMappingJson = "STORE#ENT#MAPPINGINFO#PARENTID~#~char"
              }
            }
          }
        }
      }
    }
    acctEntParams = {
      processingType = "httpEntToAcct"
      connection     = "acctAuth"
      initPrivigesMap = [
        {
          privName        = "Permission"
          attrType        = 3
          entType         = "Organization"
          cfgType         = 1
          attributeValues = ""
          defaultValue    = ""
        },
        {
          privName        = "Permission"
          attrType        = 3
          entType         = "Team"
          cfgType         = 1
          attributeValues = ""
          defaultValue    = ""
        },
        {
          privName        = "Permission"
          attrType        = 3
          entType         = "Repository"
          cfgType         = 1
          attributeValues = ""
          defaultValue    = ""
        }
      ]
      entTypes = {
        Organization = {
          call = {
            call1 = {
              callOrder   = 0
              stageNumber = 8
              http = {
                httpHeaders     = { Authorization = "$${access_token}" }
                url             = "https://<DOMAIN>/orgs/$${id}/members"
                httpContentType = "application/x-www-form-urlencoded"
                httpMethod      = "GET"
              }
              listField    = ""
              acctKeyField = "accountID"
              entKeyField  = "entitlement_value"
              acctIdPath   = "id"
              privilegeParams = {
                attrName  = "Permission"
                attrValue = "member"
              }
            }
            call2 = {
              callOrder   = 0
              stageNumber = 9
              http = {
                httpHeaders     = { Authorization = "$${access_token}" }
                url             = "https://<DOMAIN>/orgs/$${id}/members?role=admin"
                httpContentType = "application/x-www-form-urlencoded"
                httpMethod      = "GET"
              }
              listField    = ""
              acctKeyField = "accountID"
              entKeyField  = "entitlement_value"
              acctIdPath   = "id"
              privilegeParams = {
                attrName  = "Permission"
                attrValue = "admin"
              }
            }
            call3 = {
              callOrder   = 0,
              stageNumber = 10,
              http = {
                httpHeaders = {
                  Authorization = "$${access_token}"
                },
                url             = "https://<DOMAIN>/orgs/$${id}/outside_collaborators",
                httpContentType = "application/x-www-form-urlencoded",
                httpMethod      = "GET"
              },
              listField    = "",
              acctKeyField = "accountID",
              entKeyField  = "entitlement_value",
              acctIdPath   = "id",
              privilegeParams = {
                attrName  = "Permission",
                attrValue = "outside_collaborator"
              }
            }
          }
        },
        Team = {
          call = {
            call1 = {
              callOrder   = 0,
              stageNumber = 11,
              inputParams = {
                entitlementname = "Organization",
                statusFilter    = "(ev.status = 6 or ev.status = 1)"
              },
              http = {
                httpHeaders = {
                  Authorization = "$${access_token}"
                },
                url             = "https://<DOMAIN>/teams/$${id}/members?role=member",
                httpContentType = "application/x-www-form-urlencoded",
                httpMethod      = "GET"
              },
              listField    = "",
              acctKeyField = "accountID",
              entKeyField  = "entitlementID",
              acctIdPath   = "id",
              privilegeParams = {
                attrName  = "Permission",
                attrValue = "member"
              }
            },
            call2 = {
              callOrder   = 1,
              stageNumber = 12,
              inputParams = {
                entitlementname = "Organization",
                statusFilter    = "(ev.status = 6 or ev.status = 1)"
              },
              http = {
                httpHeaders = {
                  Authorization = "$${access_token}"
                },
                url             = "https://<DOMAIN>/teams/$${id}/members?role=maintainer",
                httpContentType = "application/x-www-form-urlencoded",
                httpMethod      = "GET"
              },
              listField    = "",
              acctKeyField = "accountID",
              entKeyField  = "entitlementID",
              acctIdPath   = "id",
              privilegeParams = {
                attrName  = "Permission",
                attrValue = "maintainer"
              }
            }
          }
        },
        Repository = {
          call = {
            call1 = {
              callOrder   = 0,
              stageNumber = 13,
              inputParams = {
                entitlementname = "Organization",
                statusFilter    = "(ev.status = 6 or ev.status = 1)"
              },
              http = {
                httpHeaders = {
                  Authorization = "$${access_token}"
                },
                url             = "$${id}/collaborators",
                httpContentType = "application/x-www-form-urlencoded",
                httpMethod      = "GET"
              },
              listField    = "",
              acctKeyField = "accountID",
              entKeyField  = "customproperty1",
              acctIdPath   = "id",
              privilegeParams = {
                attrName  = "Permission",
                attrValue = "member"
              }
            },
            call2 = {
              callOrder   = 1,
              stageNumber = 14,
              inputParams = {
                entitlementname = "Organization",
                statusFilter    = "(ev.status = 6 or ev.status = 1)"
              },
              acctEntMappingTypeParam = {
                mappingType     = "ENT2PRIVREVERSE",
                entitlementName = "Team"
              },
              http = {
                httpHeaders = {
                  Authorization = "$${access_token}"
                },
                url             = "${id}/teams",
                httpContentType = "application/x-www-form-urlencoded",
                httpMethod      = "GET"
              },
              listField   = "",
              entKeyField = "customproperty1",
              entIdPath   = "url",
              privilegeParams = {
                attrName      = "Permission",
                attrValuePath = "permission"
              }
            }
          }
        }
      }
    },
    globalSettingParams = {
      globalTriggerParams = {
        maxLoopCountForLastTrigger             = 5,
        lastTriggerCompletionSleep             = 2000,
        lastTriggerCompletionIntermediateSleep = 1000
      },
      supportedEntitlementTypes = [
        {
          Organization = {
            custompropertyLabels = {
              customproperty3 = "Tags",
              customproperty4 = "ID",
              customproperty6 = "Provision VM Agent",
              customproperty7 = "Enable Automatic Updates"
            }
          },
          Team = {
            custompropertyLabels = {
              customproperty3 = "Tags",
              customproperty4 = "ID",
              customproperty6 = "Provision VM Agent",
              customproperty7 = "Enable Automatic Updates"
            }
          },
          Repository = {
            custompropertyLabels = {
              customproperty3 = "Tags",
              customproperty4 = "ID",
              customproperty6 = "Provision VM Agent",
              customproperty7 = "Enable Automatic Updates"
            }
          }
        }
      ]
    }
  })
  organization_list = var.ORG_LIST
}
