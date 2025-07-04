// Copyright (c) Saviynt Inc.
// SPDX-License-Identifier: MPL-2.0

variable "CLIENT_ID" {
  type        = string
  description = "Client ID for Salesforce connection"
}
variable "CLIENT_SECRET" {
  type        = string
  description = "Client Secret for Salesforce connection"
}
variable "REFRESH_TOKEN " {
  type        = string
  description = "Refresh Token for Salesforce connection"
}
variable "REDIRECT_URI" {
  type        = string
  description = "Redirect URI for Salesforce connection"
}
variable "INSTANCE_URL" {
  type        = string
  description = "Instance URL for Salesforce connection"
}
resource "saviynt_salesforce_connection_resource" "ss" {
  connection_type       = "SalesForce"
  connection_name       = "Terraform_Salesforce_Connection"
  client_id             = var.CLIENT_ID
  client_secret         = var.CLIENT_SECRET
  refresh_token         = var.REFRESH_TOKEN
  redirect_uri          = var.REDIRECT_URI
  instance_url          = var.INSTANCE_URL
  object_to_be_imported = "Profile,Group,PermissionSet,Role"
  account_field_query   = "Id,Username,LastName,FirstName,Name,CompanyName,Email,IsActive,UserRoleId,ProfileId,UserType,ManagerId,LastLoginDate,LastPasswordChangeDate,CreatedDate,CreatedById,LastModifiedDate,LastModifiedById,SystemModstamp,ContactId,AccountId,FederationIdentifier,UserPermissionsSupportUser"
  field_mapping_json = jsonencode({
    accountfield_mapping = {
      accountID          = "Id~#~char",
      name               = "Username~#~char",
      customproperty2    = "LastName~#~char",
      customproperty1    = "FirstName~#~char",
      displayName        = "Name~#~char",
      customproperty3    = "CompanyName~#~char",
      customproperty4    = "Email~#~char",
      status             = "IsActive~#~bool",
      customproperty5    = "UserRoleId~#~char",
      customproperty6    = "ProfileId~#~char",
      accounttype        = "UserType~#~char",
      customproperty7    = "ManagerId~#~char",
      lastlogondate      = "LastLoginDate~#~date",
      lastpasswordchange = "LastPasswordChangeDate~#~date",
      CREATED_ON         = "CreatedDate~#~date",
      creator            = "CreatedById~#~char",
      customproperty8    = "LastModifiedDate~#~date",
      updateUser         = "LastModifiedById~#~char",
      updatedate         = "SystemModstamp~#~date",
      customproperty9    = "ContactId~#~char",
      customproperty10   = "AccountId~#~char",
      customproperty13   = "FederationIdentifier~#~char",
      customproperty20   = "UserPermissionsSupportUser~#~bool"
    }
  })
  createaccountjson = jsonencode([
    {
      Alias                = "$${user?.getFirstname()}",
      Email                = "$${user?.getEmail()}",
      Username             = "$${user?.getEmail()}",
      CommunityNickname    = "$${user?.getFirstname()}",
      FirstName            = "$${user?.getFirstname()}",
      LastName             = "$${user?.getLastname()}",
      TimeZoneSidKey       = "America/Los_Angeles",
      LocaleSidKey         = "en_US",
      EmailEncodingKey     = "ISO-8859-1",
      ProfileId            = "$${profileId}",
      LanguageLocaleKey    = "en_US",
      IsActive             = true,
      FederationIdentifier = "$${user?.getEmail()}"
    }
  ])
  customconfigjson = jsonencode([
    {
      disableAccountForRevokeTask = false,
      defaultEntitlementId        = "<default entitlement id>",
    }
  ])
}
