// Copyright (c) Saviynt Inc.
// SPDX-License-Identifier: MPL-2.0

variable "LDAP_PROTOCOL" {
  type        = string
  description = "Protocol type (e.g., LDAP, HTTP, etc.)"
}
variable "IP_ADDRESS" {
  type        = string
  description = "Saviynt host server"
}
variable "PASSWORD" {
  type        = string
  description = "Connection password"
  sensitive   = true
}
variable "BIND_USER" {
  type        = string
  description = "Connection username"
}
variable "VAULT_CONNECTION" {
  type        = string
  description = "Vault connection"
}
variable "VAULT_CONFIG" {
  type        = string
  description = "Vault config"
}
variable "SAVE_IN_VAULT" {
  type        = string
  description = "Save in vault"
}
variable "CONNECTION_URL" {
  type        = string
  description = "Value of CONNECTION_URL"
}
variable "PROVISIONING_URL" {
  type        = string
  description = "Value of PROVISIONING_URL"
}
resource "saviynt_adsi_connection_resource" "example" {
  connection_type     = "ADSI"
  connection_name     = "Terraform_ADSI_Connector"
  email_template      = "Account Password Expiry Email"
  url                 = format("%s://%s", var.LDAP_PROTOCOL, var.IP_ADDRESS)
  password            = var.PASSWORD
  username            = var.BIND_USER
  vault_connection    = var.VAULT_CONNECTION
  vault_configuration = var.VAULT_CONFIG
  save_in_vault       = var.SAVE_IN_VAULT
  connection_url      = var.CONNECTION_URL
  provisioning_url    = var.PROVISIONING_URL
  forestlist          = "domain.org"
  searchfilter        = "DC=domain,DC=org"
  objectfilter        = "(&(objectCategory=person)(objectClass=user))"
  user_attribute      = "[FIRSTNAME::givenname#String,LASTNAME::sn#String,CUSTOMPROPERTY1::samaccountname#String,USERNAME::distinguishedname#String,DISPLAYNAME::cn#String,CUSTOMPROPERTY25::description#String,CUSTOMPROPERTY3::sn#String,COMMENTS::distinguishedname#String,CUSTOMPROPERTY4::homedirectory#String,CUSTOMPROPERTY5::co#String,CUSTOMPROPERTY6::cn#String,CUSTOMPROPERTY7::givenname#String,CUSTOMPROPERTY8::title#String,CUSTOMPROPERTY9::telephonenumber#String,CUSTOMPROPERTY10::c#String,CUSTOMPROPERTY11::uSNCreated#String,ENDDATE::accountExpires#millisec,CUSTOMPROPERTY12::logonCount#String,CUSTOMPROPERTY13::physicaldeliveryofficename#String,UPDATEDATE::whenchanged#date,CUSTOMPROPERTY14::extensionattribute1#String,CUSTOMPROPERTY15::extensionattribute2#String,CUSTOMPROPERTY16::streetaddress#String,CUSTOMPROPERTY17::mailnickname#String,CUSTOMPROPERTY18::department#String,CUSTOMPROPERTY19::countrycode#String,CUSTOMPROPERTY2::samaccountname#String,CUSTOMPROPERTY20::userprincipalname#String,CUSTOMPROPERTY21::manager#String,CUSTOMPROPERTY22::homephone#String,CUSTOMPROPERTY23::mobile#String,CREATEDATE::whencreated#date,customproperty24::useraccountcontrol#String,CUSTOMPROPERTY26::distinguishedname#String,statuskey::useraccountcontrol#String,CUSTOMPROPERTY27::objectguid#String,RECONCILATION_FIELD::CUSTOMPROPERTY27,CUSTOMPROPERTY28::forest#String,CUSTOMPROPERTY29::domain#string]"
  statuskeyjson = jsonencode({
    STATUS_ACTIVE = [
      "512",
      "544",
      "66048"
    ],
    STATUS_INACTIVE = [
      "546",
      "514"
    ]
  })
  account_attribute = "[CUSTOMPROPERTY1::samaccountname#String,CUSTOMPROPERTY2::userprincipalname#String,LASTLOGONDATE::lastLogon#millisec,DISPLAYNAME::cn#String,CUSTOMPROPERTY25::company#String,CUSTOMPROPERTY3::sn#String,COMMENTS::distinguishedname#String,CUSTOMPROPERTY4::homedirectory#String,LASTPASSWORDCHANGE::pwdlastset#millisec,CUSTOMPROPERTY5::co#String,CUSTOMPROPERTY6::cn#String,CUSTOMPROPERTY7::givenname#String,CUSTOMPROPERTY8::title#String,CUSTOMPROPERTY9::telephonenumber#String,CUSTOMPROPERTY10::c#String,DESCRIPTION::description#String,CUSTOMPROPERTY11::usncreated#String,VALIDTHROUGH::accountexpires#millisec,CUSTOMPROPERTY12::logoncount#String,CUSTOMPROPERTY13::physicaldeliveryofficename#String,UPDATEDATE::whenchanged#date,CUSTOMPROPERTY14::extensionattribute1#String,CUSTOMPROPERTY15::extensionattribute2#String,CUSTOMPROPERTY16::streetaddress#String,CUSTOMPROPERTY17::mailnickname#String,CUSTOMPROPERTY18::department#String,CUSTOMPROPERTY19::countrycode#String,NAME::distinguishedname#String,CUSTOMPROPERTY21::manager#String,CUSTOMPROPERTY22::homephone#String,CUSTOMPROPERTY23::mobile#String,ACCOUNTCLASS::objectclass#String,ACCOUNTID::objectguid#String,CUSTOMPROPERTY24::useraccountcontrol#String,CUSTOMPROPERTY28::forest#String,CUSTOMPROPERTY29::domain#string,CUSTOMPROPERTY30::objectclass#String,CUSTOMPROPERTY31::displayname#String,RECONCILATION_FIELD::ACCOUNTID]"

  status_threshold_config = jsonencode({
    statusAndThresholdConfig = {
      statusColumn = "customproperty24",
      activeStatus = [
        "512",
        "544",
        "66048"
      ],
      deleteLinks                   = false,
      accountThresholdValue         = 1000,
      correlateInactiveAccounts     = true,
      inactivateAccountsNotInFile   = false,
      deleteAccEntForActiveAccounts = false
    }
  })
  endpoints_filter = jsonencode(
    {
      "Test Endpoint Filter 1" : [
        {
          memberOf = [
            "CN=ChildSaviyntlabs\\,DomainLocal,OU=Groups,DC=child,DC=saviyntlabs,DC=org",
            "CN=Saviyntlabs-DomainLocal,OU=Groups,DC=saviyntlabs,DC=org"
          ]
        }
      ]
    }
  )
  entitlement_attribute = "memberOf"
  accountnamerule = jsonencode({
    AccountNameRule = {
      Attributes = [
        {
          cn        = "ACCOUNTID",
          baseDN    = "$${baseDN}",
          RuleCheck = "$${user.lastname}, $${user.firstname}###$${user.lastname}, $${user.firstname}1###$${user.lastname}, $${user.firstname}2###$${user.lastname}, $${user.firstname}3###$${user.firstname}4"
        }
      ]
    }
  })
  checkforunique = jsonencode(
    { CheckForUnique = { Attributes = [{ samaccountname = "customproperty1", RuleCheck = "$${user.lastname}###$${user.lastname}1###$${user.lastname}2###$${user.lastname}3###$${user.lastname}4###$${user.lastname}5###$${user.lastname}6###$${user.lastname}7###$${user.lastname}8" }, { userprincipalname = "customproperty2", RuleCheck = "$${user.lastname}@domainame.com###$${user.lastname}1@domainname.com###$${user.lastname}2@domainname.com###$${user.lastname}3@domainname.com###$${user.lastname}4@domainname.com" }, { displayname = "customproperty31", RuleCheck = "$${user.lastname}, $${user.firstname}###$${user.lastname}, $${user.firstname}1###$${user.lastname}, $${user.firstname}2###$${user.lastname}, $${user.firstname}3###$${user.lastname}, $${user.firstname}4" }] } }
  )
  group_search_base_dn = "DC=domain,DC=org"
  group_import_mapping = jsonencode({
    importGroupHierarchy       = "true",
    entitlementTypeName        = "memberOf",
    performGroupAccountLinking = "true",
    groupObjectClass           = "(objectclass=group)",
    entitlementOwnerAttribute  = "managedby",
    tableFieldAttribute        = "accountID",
    mapping                    = "memberHash:memberof_char,customProperty1:samaccounttype_char,customProperty2:instancetype_char,customProperty3:usncreated_char,customProperty4:grouptype_char,customProperty5:dscorepropagationdata_char,customProperty12:dn_char,customProperty13:cn_char,lastscandate:whencreated_date,customProperty15:managedBy_char,entitlement_glossary:description_char,description:description_char,displayname:name_char,customProperty9:name_char,customProperty10:objectcategory_char,customProperty11:samaccounttype_char,entitlement_value:distinguishedname_char,entitlementid:distinguishedname_char,customProperty14:objectclass_char,updatedate:whenchanged_date,customProperty17:distinguishedname_char,RECONCILATION_FIELD:customProperty18,customProperty18:objectguid_char"
  })
  import_nested_membership = "FALSE"
  page_size                = "1000"
  createaccountjson = jsonencode({
    objects = [
      {
        objectClasses = [
          "user",
          "top",
          "Person",
          "OrganizationalPerson"
        ],
        baseDn   = "CN=Users,DC=domain,DC=org",
        password = "*******",
        attributes = {
          sn                 = "$${user.lastname}",
          sAMAccountName     = "$${task.accountName}",
          cn                 = "$${task.accountName}",
          userAccountControl = 512,
          co                 = "",
          department         = "$${user.departmentname}",
          displayName        = "$${user.displayname}",
          employeeID         = "$${user.employeeid}",
          employeeNumber     = "1",
          employeeType       = "$${user.employeeType}",
          givenName          = "$${user.firstname}",
          l                  = "$${user.city}",
          mail               = "$${user.email}",
          proxyAddresses = [
            "SMTP:test@test.org",
            "SIP:Test@test1.org"
          ]
        }
      }
    ]
  })
  updateaccountjson = jsonencode({
    objects = [
      {
        objectClasses = [
          "user"
        ]
        distinguishedName = "$${account.accountID?.replace('\\', '\\\\')?.replace('/', '\\/')}"
        attributes = {
          displayName   = "$${user.displayname}",
          streetAddress = "$${user.street}",
          additionalAttributes = {
            departmentName = "dept",
            companyName    = "domain"
          }
        }
      }
    ]
  })
  addaccessjson = jsonencode({
    objects = [
      {
        objectClasses = [
          "user"
        ]
        distinguishedName = "$${accountID}",
        addGroup          = "$${entitlement_values}"
      }
    ]
    requestConfiguration = {
      grpMemExistenceChk = {
        enable = true
      }
    }
  })
  removeaccessjson = jsonencode({
    objects = [
      {
        objectClasses = [
          "group"
        ]
        distinguishedName = "$${accountID}",
        removeGroup       = "$${entitlement_values}"
      }
    ]
    requestConfiguration = {
      grpMemExistenceChk = {
        enable = true
      }
    }
  })
  enableaccountjson = jsonencode({
    objects = [
      {
        objectClasses = [
          "user"
        ],
        distinguishedName = "$${account.accountID}",
        deleteAllGroups   = false,
        attributes = {
          userAccountControl = 512
        }
      }
    ]
  })
  disableaccountjson = jsonencode({
    objects = [
      {
        objectClasses = [
          "user"
        ],
        distinguishedName = "$${account.accountID}",
        deleteAllGroups   = false,
        attributes = {
          userAccountControl = 514
        }
      }
    ]
  })
  removeaccountjson = jsonencode({
    objects = [
      {
        distinguishedName  = "$${account.accountID}",
        removeAction       = "DELETE",
        deleteChildObjects = false
      }
    ]
  })
  resetandchangepasswrdjson = jsonencode({
    objects = [
      {
        objectClasses = [
          "user"
        ],
        password          = "$${password}",
        distinguishedName = "",
        attributes = {
          pwdLastSet = "$${pwdLastSet}"
        }
      }
    ]
  })
  creategroupjson = jsonencode({
    objects = [
      {
        objectClasses = [
          "group"
        ],
        baseDn = "$${role.customproperty24}",
        attributes = {
          cn          = "$${role.displayname}",
          name        = "$${role.displayname}",
          description = "$${role.description}",
          displayName = "$${role.displayname}",
          groupType   = "$${role?.customproperty21 == 'Security' && role?.customproperty22 == 'Global'?'-2147483646' : role?.customproperty21=='Security'&&role?.customproperty22=='Universal'?'-2147483640' : role?.customproperty21== 'Security'&&role?.customproperty22=='Domain Local' ? '-2147483644':role?.customproperty21=='Distribution'&&role?.customproperty22=='Global' ? '2':role?.customproperty21== 'Distribution'&&role?.customproperty22=='Universal'?'8':role?.customproperty21=='Distribution'&& role?.customproperty22=='Domain Local'?'4':''}"
        }
      }
    ]
  })
  updategroupjson = jsonencode({
    objects = [
      {
        objectClasses = [
          "group"
        ],
        distinguishedName = "$${role.role_name}",
        attributes = {
          description    = "$${role.description}",
          proxyAddresses = "$${role.customproperty20}"
        }
      }
    ]
  })
  removegroupjson = jsonencode({
    objects = [
      {
        distinguishedName  = "$${role.role_name}",
        removeAction       = "DELETE",
        deleteChildObjects = false
      }
    ]
  })
  addaccessentitlementjson = jsonencode({
    objects = [
      {
        objectClasses = [
          "group"
        ],
        distinguishedName = "$${ent1Value.entitlement_value?.replace('\\', '\\\\')?.replace('/', '\\/')}",
        addGroup          = "$${ent2Value.entitlement_value?.replace('\\', '\\\\')?.replace('/', '\\/')}"
      }
    ]
    requestConfiguration = {
      grpMemExistenceChk = {
        enable = true
      }
    }
  })
  removeaccessentitlementjson = jsonencode({
    objects = [
      {
        objectClasses = [
          "group"
        ],
        distinguishedName = "$${ent1Value.entitlement_value?.replace('\\', '\\\\')?.replace('/', '\\/')}",
        removeGroup       = "$${ent2Value.entitlement_value?.replace('\\', '\\\\')?.replace('/', '\\/')}"
      }
    ]
    requestConfiguration = {
      grpMemExistenceChk = {
        enable = true
      }
    }
  })
  createserviceaccountjson = jsonencode({
    objects = [
      {
        objectClasses = [
          "user",
          "top",
          "Person",
          "OrganizationalPerson"
        ],
        baseDn   = "$${baseDN}",
        password = "$${password}",
        attributes = {
          sAMAccountName     = "$${task.accountName}",
          cn                 = "$${task.accountName}",
          displayname        = "testDP",
          userAccountControl = 512
        }
      }
    ]
  })
  updateserviceaccountjson = jsonencode({
    objects = [
      {
        objectClasses = [
          "user",
          "top",
          "Person",
          "OrganizationalPerson"
        ],
        distinguishedName = "$${account.accountID?.replace('\\', '\\\\')?.replace('/', '\\/')}",
        attributes = {
          pwdLastSet  = 0,
          displayName = "testDPUpdated"
        }
      }
    ]
  })
  removeserviceaccountjson = jsonencode({
    objects = [
      {
        distinguishedName  = "$${account.accountID?.replace('\\', '\\\\')?.replace('/', '\\/')}",
        removeAction       = "DELETE",
        deleteChildObjects = false
      }
    ]
  })
  customconfigjson = jsonencode(
    {
      connectionTimeoutConfig = {
        import = {
          timeout   = 300,
          retryWait = 500
        },
        provisioning = {
          timeou    = 300,
          retryWait = 500
        }
      }
    }
  )
  updateuserjson = jsonencode({
    objects = [
      {
        objectClasses     = ["user"]
        distinguishedName = "$${account.COMMENTS?.replace('', '\\')?.replace('/', '/')}"
        moveObjectToOU    = "OU=TestContainer,DC=domain1,DC=com"
        attributes = {
          displayName   = "$${user.displayname}"
          streetaddress = "$${user.street}"
        }
      }
    ]
  })
}
