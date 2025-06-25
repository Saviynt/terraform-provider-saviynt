// Copyright (c) Saviynt Inc.
// SPDX-License-Identifier: MPL-2.0

variable "IP_ADDRESS" {
  type        = string
  description = "Saviynt host server"
}
variable "LDAP_PORT" {
  type        = string
  description = "Port for the connection"
}
variable "LDAP_PROTOCOL" {
  type        = string
  description = "Protocol type (e.g., LDAP, HTTP, etc.)"
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
variable "BASE_CONTAINER" {
  type        = string
  description = "Value of BASEDN"
}
variable "DOMAIN" {
  type        = string
  description = "Value of DOMCONTRDN"
}
resource "saviynt_ad_connection_resource" "example" {
  connection_type       = "AD"
  connection_name       = "Terraform_AD_Connector"
  url                   = format("%s://%s:%d", var.LDAP_PROTOCOL, var.IP_ADDRESS, var.LDAP_PORT)
  password              = var.PASSWORD
  username              = var.BIND_USER
  vault_connection      = var.VAULT_CONNECTION
  vault_configuration   = var.VAULT_CONFIG
  save_in_vault         = var.SAVE_IN_VAULT
  searchfilter          = var.BASE_CONTAINER
  base                  = var.BASE_CONTAINER
  group_search_base_dn  = "OU=Groups,DC=domainname,DC=com"
  ldap_or_ad            = "AD"
  objectfilter          = "(&(objectCategory=person)(objectClass=user))"
  account_attribute     = "[CUSTOMPROPERTY1::CN#String,CUSTOMPROPERTY2::userPrincipalName#String,LASTLOGONDATE::lastLogon#millisec,DISPLAYNAME::displayName#String,CUSTOMPROPERTY25::company#String,CUSTOMPROPERTY3::sn#String,COMMENTS::distinguishedName#String,CUSTOMPROPERTY4::homeDirectory#String,LASTPASSWORDCHANGE::pwdLastSet#millisec,CUSTOMPROPERTY5::co#String,CUSTOMPROPERTY6::cn#String,CUSTOMPROPERTY7::givenName#String,CUSTOMPROPERTY8::title#String,CUSTOMPROPERTY9::telephoneNumber#String,CUSTOMPROPERTY10::c#String,DESCRIPTION::description#String,CUSTOMPROPERTY11::uSNCreated#String,VALIDTHROUGH::accountExpires#millisec,CUSTOMPROPERTY12::logonCount#String,CUSTOMPROPERTY13::physicalDeliveryOfficeName#String,UPDATEDATE::whenChanged#date,CUSTOMPROPERTY14::extensionAttribute1#String,CUSTOMPROPERTY15::extensionAttribute2#String,CUSTOMPROPERTY16::streetAddress#String,CUSTOMPROPERTY17::mailNickname#String,CUSTOMPROPERTY18::department#String,CUSTOMPROPERTY19::countryCode#String,NAME::name#String,CUSTOMPROPERTY21::manager#String,CUSTOMPROPERTY22::homePhone#String,CUSTOMPROPERTY23::mobile#String,CREATED_ON::whenCreated#date,ACCOUNTCLASS::objectClass#String,ACCOUNTID::distinguishedName#String,CUSTOMPROPERTY24::userAccountControl#String,STATUS::userAccountControl#Number,CUSTOMPROPERTY26::objectGUID#Binary,CUSTOMPROPERTY27::objectSid#Binary]"
  entitlement_attribute = "memberOf"
  page_size             = "1000"
  user_attribute        = "[CUSTOMPROPERTY1::CN#String,USERNAME::name#String,DISPLAYNAME::displayName#String,CUSTOMPROPERTY25::company#String,CUSTOMPROPERTY3::sn#String,COMMENTS::distinguishedName#String,CUSTOMPROPERTY4::homeDirectory#String,CUSTOMPROPERTY5::co#String,CUSTOMPROPERTY6::cn#String,CUSTOMPROPERTY7::givenName#String,CUSTOMPROPERTY8::title#String,CUSTOMPROPERTY9::telephoneNumber#String,CUSTOMPROPERTY10::c#String,CUSTOMPROPERTY11::uSNCreated#String,ENDDATE::accountExpires#millisec,CUSTOMPROPERTY12::logonCount#String,CUSTOMPROPERTY13::physicalDeliveryOfficeName#String,UPDATEDATE::whenChanged#date,CUSTOMPROPERTY14::extensionAttribute1#String,CUSTOMPROPERTY15::extensionAttribute2#String,CUSTOMPROPERTY16::streetAddress#String,CUSTOMPROPERTY17::mailNickname#String,CUSTOMPROPERTY18::department#String,CUSTOMPROPERTY19::countryCode#String,CUSTOMPROPERTY2::sAMAccountName#String,CUSTOMPROPERTY20::userPrincipalName#String,CUSTOMPROPERTY21::manager#String,CUSTOMPROPERTY22::homePhone#String,CUSTOMPROPERTY23::mobile#String,CREATEDATE::whenCreated#date,customproperty24::userAccountControl#String,STATUSKEY::userAccountControl#Number]"
  filter                = "(sAMAccountName=*-adm)"
  endpoints_filter      = jsonencode({ AD_Child = [{ memberOf = ["CN=ADGroup15,DC=domainname,DC=com"] }] })
  create_account_json = jsonencode(
    { samaccountname = "$${task.accountName}", sn = "$${user.lastname}", displayName = "$${user.displayname}", cn = "$${cn}", objectclass = ["top", "person", "organizationalPerson", "user"], userAccountControl = "544", givenName = "$${user.firstname}", name = "$${user.displayname}" }
  )
  import_json = jsonencode(
    {
      envproperties = {
        "com.sun.jndi.ldap.connect.timeout" = "10000",
        "com.sun.jndi.ldap.read.timeout"    = "50000"
      }
  })
  advsearch = jsonencode(
    {
      params = [{
        name     = "uid",
        label    = "User ID",
        REQUIRED = false,
        FILTER   = true,
        ENDPOINT = "TestEndpoint"
        },
        {
          name  = "numofrecord",
          label = "NUMBER OF RECORD PER PAGE",
          value = "10"
      }]
  })
  check_for_unique = jsonencode(
    {
      userPrincipalName = "$${user.firstname}.$${user.lastname}@company.com###$${user.firstname}.$${user.lastname}1@company.com"
    }
  )
  incremental_config = jsonencode({
    incrementalImportType   = "NotUsed"
    changeLogBase           = "OU=TestOU,DC=domainname,DC=com"
    changeNumberFilter      = "&(changeNumber>=##MAX_CHANGENUMBER##)(targetDN=*OU=TestOU,DC=domainname,DC=com*)"
    dnAttributeName         = "targetDn"
    dnAttributeNameMappedTo = "username"
    changeNumberAttrName    = "uSNChanged"
    changeTypeAttrName      = "changeType"
    changedFeildsInScope    = "status,CUSTOMPROPERTY1,CUSTOMPROPERTY2,LASTLOGONDATE,DISPLAYNAME,CUSTOMPROPERTY25,CUSTOMPROPERTY3,COMMENTS,CUSTOMPROPERTY4,CUSTOMPROPERTY5,CUSTOMPROPERTY6,CUSTOMPROPERTY7,CUSTOMPROPERTY8,CUSTOMPROPERTY9,DESCRIPTION,CUSTOMPROPERTY14,CUSTOMPROPERTY15,CUSTOMPROPERTY16,CUSTOMPROPERTY17,CUSTOMPROPERTY18,NAME,CUSTOMPROPERTY20,CREATED_ON,ACCOUNTCLASS,customProperty29,dummy"
    changesLogAttrName      = "changes"
    searchAttribute         = "entrydn"
    searchOn                = "rdn"
  })
  reuse_account_json = jsonencode(
    {
      ATTRIBUTESTOCHECK = {
        userAccountControl = "514",
        sn                 = "$${user.lastname}",
      cn = "$${user.fistname}" },
      REUSEACCOUNTOU = "OU=ActiveUsers,DC=domainname,DC=com"
  })
  reset_and_change_passwrd_json = jsonencode({
    RESET = {
      pwdLastSet = "0"
      title      = "password reset"
    }
    CHANGE = {
      pwdLastSet = "-1"
      title      = "password changed"
    }
  })
  update_account_json = jsonencode({
    department = "$${user.departmentname ?: null}",
    division   = "$${user.siteid ?: null}"
  })
  update_user_json = jsonencode({
    sn = "$${user.lastname}"
  })
  enable_account_json = jsonencode({
    USEDNFROMACCOUNT = "YES",
    MOVEDN           = "NO",
    REMOVEGROUPS     = "NO",
    ENABLEACCOUNTOU  = "CN=Users,DC=corp,DC=AD,DC=com",
    AFTERMOVEACTIONS = {
      userAccountControl = "544",
      otherMailbox = [
        "$${user?.customproperty15.toString().replace(',','\",\"')}"
      ]
    }
  })
  reuse_inactive_account = "TRUE"
  account_name_rule      = "CN=$${user.username},CN=Users,DC=corp,DC=AD,DC=com"
  remove_account_action  = jsonencode({ removeAction = "SUSPEND", userAccountControl = "546" })
  set_random_password    = "FALSE"
  password_min_length    = "2"
  password_max_length    = "2"
  password_noofcapsalpha = "2"
  password_noofsplchars  = "2"
  password_noofdigits    = "2"
  group_import_mapping = jsonencode({
    importGroupHierarchy       = "true",
    entitlementTypeName        = "",
    performGroupAccountLinking = "true",
    incrementalTimeField       = "whenChanged",
    groupObjectClass           = "(objectclass=group)",
    mapping                    = "memberHash:member_char,customproperty1:sAMAccountType_char,customproperty16:memberOf_char,customproperty2:instanceType_char, customproperty3:uSNCreated_char,customproperty4:groupType_char,customproperty5:dSCorePropagationData_char,customproperty12:dn_char,customproperty13:cn_char,lastscandate:whenCreated_date,customproperty15:managedBy_char,entitlement_glossary:description_char,customproperty9:name_char,customproperty10:objectCategory_char,customproperty11:sAMAccountName_char,customproperty14:objectClass_char,status:isCriticalSystemObject_char,entitlement_value:distinguishedName_char,entitlementid:objectGUID_Binary,customproperty17:distinguishedName_char,updatedate:whenChanged_date,RECONCILATION_FIELD:entitlementid",
    entitlementOwnerAttribute  = "managedBy",
    tableFieldAttribute        = "COMMENTS"
  })
  create_update_mappings = "{\"cn\":\"$${role?.customproperty27}\",\"objectCategory\":\"CN=Group,CN=Schema,CN=Configuration,DC=corp,DC=domainname,DC=com\",\"distinguishedName\":\"$${role?.role_name}\",\"displayName\":\"$${role?.displayname}\",\"SamAccountName\":\"$${role?.customproperty27}\",\"description\":\"$${role?.description}\",\"objectClass\":\"group\",\"name\":\"$${role?.customproperty27}\",\"managedBy\":\"$${user?.comments}\"}"
  max_changenumber       = "1000"
  support_empty_string   = "TRUE"
  status_key_json = jsonencode({
    STATUS_ACTIVE = [
      "1", "ACTIVE", "true", "512", "544"
    ],
    STATUS_INACTIVE = [
      "0", "INACTIVE", "false", "546", "514"
    ]
  })
  status_threshold_config = jsonencode({
    statusAndThresholdConfig = {
      statusColumn                = "customproperty24",
      activeStatus                = ["512", "544"],
      deleteLinks                 = false,
      accountThresholdValue       = 1000,
      correlateInactiveAccounts   = true,
      inactivateAccountsNotInFile = false
    }
  })
  disable_account_json = jsonencode({
    userAccountControl = "546",
    deleteAllGroups    = "No"
  })
  config_json = jsonencode({
    connectionTimeoutConfig = {
      connectionTimeout = 10,
      readTimeout       = 50,
      retryWait         = 2,
      retryCount        = 3
    },
    ldapPolicy = {
      enforceNonLeafSearchContext = true
    }
  })
  read_operational_attributes = "FALSE"
  enforce_tree_deletion       = "FALSE"
  dc_locator                  = "Win-DC-Locator"
  advance_filter_json = jsonencode({
    AdvanceFilter = {
      "OU=Employees,DC=domainname,DC=com" = [
        "(&(objectCategory=person)(objectClass=user)(department=PM))"
      ]
      "OU=Vendors,DC=domainname,DC=com" = [
        "(&(objectCategory=person)(objectClass=user)(department=Vendor))"
      ]
    }
  })
  org_base = "DC=domainname,DC=com"
  organization_attribute = jsonencode({
    mapping = "CUSTOMERNAME::ou#String,CUSTOMPROPERTY1::name#String,CUSTOMPROPERTY11::whenChanged#String,CUSTOMPROPERTY2::st#String,CUSTOMPROPERTY3::postBoxOffice#String,CUSTOMPROPERTY4::postalAddress#String,CUSTOMPROPERTY5::postalCode#String,CUSTOMPROPERTY6::cn#String,DESCRIPTION::description#String,CREATEDATE::whenCreated#date,UPDATEDATE::whenChanged#date,ENTITYCLASS::objectClass#String,RECONCILATION_FIELD::CUSTOMPROPERTY10,CUSTOMPROPERTY10::objectGUID#Binary,CUSTOMERTYPE::1#SavData,STATUS::1#SavData,RISK::0#SavData,SCORE::0#SavData"
    attributes = [
      {
        name    = "locality"
        filter  = ["(L=*)"]
        mapping = "description#string,displayName#String,street#String"
      }
    ]
  })
  create_org_json = jsonencode(
    {
      name           = "$${customer.customername}"
      objectClass    = ["top", "organization"]
      objectCategory = "CN=Organization,CN=Schema,CN=Configuration,DC=domainname,DC=com"
      o              = "$${customer.customername}"
      description    = "$${customer.description}"
      attributes = [
        {
          name    = "locality"
          filter  = ["(L=*)"]
          mapping = "street,postalAddress"
        }
      ]
    }
  )
  update_org_json = jsonencode({
    o           = "$${customer.customername}"
    description = "$${customer.description}"
    attributes = [
      {
        name    = "locality"
        filter  = ["(L=*)"]
        mapping = "street,postalAddress"
      }
    ]
  })
  enable_group_management = "TRUE"
  unlock_account_json = jsonencode(
    {
      lockoutTime = "0"
    }
  )
}
