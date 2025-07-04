// Copyright (c) 2025 Saviynt Inc.
// SPDX-License-Identifier: MPL-2.0

variable "URL" {
  type        = string
  description = "DB Connector URL"
}
variable "USERNAME" {
  type        = string
  description = "DB Connector USERNAME"
}
variable "DRIVERNAME" {
  type        = string
  description = "DB Connector DRIVERNAME"
}
variable "PASSWORD" {
  type        = string
  description = "DB Connector PASSWORD"
  sensitive   = true
}
resource "saviynt_db_connection_resource" "example" {
  connection_type           = "DB"
  connection_name           = "Terraform_DB_Connector"
  url                       = var.URL
  username                  = var.USERNAME
  password                  = var.PASSWORD
  driver_name               = var.DRIVERNAME
  password_min_length       = "2"
  password_max_length       = "2"
  password_no_of_caps_alpha = "2"
  password_no_of_digits     = "2"
  password_no_of_spl_chars  = "2"
  create_account_json = jsonencode({
    "CreateAccountQry" : [
      "CREATE USER $${accountName.toUpperCase()} PASSWORD $${randomPassword};",
      "ALTER USER $${accountName.toUpperCase()} SET PARAMETER 'FIRST NAME' = '$${user.firstname?.toUpperCase()}','LAST NAME' = '$${user.lastname?.toUpperCase()}',EMAIL ADDRESS = '$${user.email}'"
    ]
  })
  update_account_json = jsonencode({
    UpdateAccountQry = [
      "UPDATE mysqllocal.users SET firstname = $${user.firstname}, lastname = $${user.lastname}, departmentname = $${user.departmentname}, displayname = $${user.displayname}, manager = $${user.manager}, orgunitID = $${user.orgunitID}, updatedate = CURRENT_TIMESTAMP WHERE username = $${user.username}"
    ]
  })
  grant_access_json = jsonencode({
    "HANA_Role" : ["CALL GRANT_ACTIVATED_ROLE ('$${task.entitlement_valueKey.entitlement_value}','$${accountName.toUpperCase()}')"]
  })
  revoke_access_json = jsonencode({
    "HANA_Role" : ["CALL REVOKE_ACTIVATED_ROLE('$${task.entitlement_valueKey.entitlement_value}','$${accountName.toUpperCase()}')"]
  })
  change_pass_json = jsonencode({
    ChangePasswordQry = [
      "UPDATE mysqllocal.users SET password = $${user.password}, updatedate = CURRENT_TIMESTAMP WHERE username = $${user.username}"
    ]
  })
  delete_account_json = jsonencode({
    DeleteAccountQry = [
      "DELETE FROM mysqllocal.users WHERE username = $${user.username}",
      "DELETE FROM mysqllocal.accounts WHERE AccountName =$${user.username}"
    ]
  })
  enable_account_json = jsonencode({
    EnableAccountQry = [
      "UPDATE mysqllocal.users SET enabled = 1, updatedate = CURRENT_TIMESTAMP WHERE username = $${user.username}",
      "UPDATE mysqllocal.accounts SET Status = 1, UPDATEDATE = CURRENT_TIMESTAMP WHERE AccountName = $${user.username}"
    ]
  })
  disable_account_json = jsonencode({
    DisableAccountQry = [
      "UPDATE mysqllocal.users SET enabled = 0, updatedate = CURRENT_TIMESTAMP WHERE username = $${user.username}",
      "UPDATE mysqllocal.accounts SET Status = 0, UPDATEDATE = CURRENT_TIMESTAMP WHERE AccountName = $${user.username}"
    ]
  })
  account_exists_json = jsonencode({
    AccountExistQry = "SELECT username FROM mysqllocal.users WHERE username = $${user.username}"
  })
  update_user_json = jsonencode({
    UpdateUserQry = [
      "UPDATE mysqllocal.users SET firstname = $${user.firstname}, lastname = $${user.lastname}, departmentname = $${user.departmentname}, displayname = $${user.displayname}, manager = $${user.manager}, orgunitID = $${user.orgunitID}, updatedate = CURRENT_TIMESTAMP WHERE username = $${user.username}"
    ]
  })
  accounts_import = trimspace(<<EOF
  <dataMapping>
    <count-query description="This is the Source Endpoint DB Count Query">
        <![CDATA[
select count(*) as count, accountname, 'Test-DB-Connector-mysql-SS' as endpoint, 'Test-DB-Connector-mysql-SS' as securitysystem, "1" as status,updatedate from mysqllocal.accounts;
]]>
    </count-query>
    <sql-query description="This is the Source DB Query" uniquecolumnsascommaseparated="name">
        <![CDATA[select accountname,'Test-DB-Connector-mysql-SS' as securitysystem,'Test-DB-Connector-mysql-SS' as endpoint,entitlementtype,entitlementvalue,status,updatedate from mysqllocal.accounts]]>
    </sql-query>
    <incrementalcondition>
        <![CDATA['$${incrementalcolmaxval.format("yyy-mm-dd hh:mm:ss")}']]>
    </incrementalcondition>
    <mapper description="This is the mapping field for Saviynt Field name" accountnotinfileaction="Suspend" deleteaccountentitlement="true" ifusernotexists="noaction" incrementalcolumn="accounts.UPDATEDATE" systems="'Test-DB-Connector-mysql-SS'">
        <mapfield saviyntproperty="accounts.name" sourceproperty="accountName" type="character"></mapfield>
        <mapfield saviyntproperty="securitysystems.systemname" sourceproperty="securitysystem" type="character"></mapfield>
        <mapfield saviyntproperty="endpoints.endpointname" sourceproperty="endpoint" type="character"></mapfield>
        <mapfield saviyntproperty="entitlementtypes.entitlementname" sourceproperty="entitlementtype" type="character"></mapfield>
        <mapfield saviyntproperty="entitlementvalues.entitlementvalue" sourceproperty="entitlementvalue" type="character"></mapfield>
        <mapfield saviyntproperty="accounts.status" sourceproperty="status" type="character"></mapfield>
        <mapfield saviyntproperty="accounts.customproperty25" sourceproperty="updatedate" type="date"/>
    </mapper>
</dataMapping>
EOF
  )
  entitlement_value_import = trimspace(<<EOF
  <dataMapping>
    <count-query description="This is the Source Endpoint DB Count Query">
        <![CDATA[
select count(*) as count, endpointname,owner, 'Test-DB-Connector-mysql-SS' as endpoint, 'Test-DB-Connector-mysql-SS' as securitysystem, "1" as status from mysqllocal.accounts;
]]>
    </count-query>
    <sql-query description="This is the Source Database Query">
        <![CDATA[select 'Test-DB-Connector-mysql-SS' as systemname,dataowner1,'Test-DB-Connector-mysql-SS' as endpointname,entitlementtype,entitlementvalue,entID,1 as status from mysqllocal.entitlements]]>
    </sql-query>
    <mapper description="This is the mapping field for Saviynt Field name" deleteentitlementowner="true" entnotpresentaction="noaction" createentitlementtype="true" systems="'Test-DB-Connector-mysql-SS'">
        <mapfield saviyntproperty="securitysystems.systemname" sourceproperty="systemname" type="character"></mapfield>
        <mapfield saviyntproperty="endpoints.endpointname" sourceproperty="endpointname" type="character"></mapfield>
        <mapfield saviyntproperty="entitlementtypes.entitlementname" sourceproperty="entitlementtype" type="character"></mapfield>
        <mapfield saviyntproperty="entitlementvalues.entitlement_value" sourceproperty="entitlementvalue" type="character"></mapfield>
        <mapfield saviyntproperty="entitlementvalues.entitlementID" sourceproperty="entID" type="character"></mapfield>
        <mapfield saviyntproperty="entitlementvalues.status" sourceproperty="status" type="number"></mapfield>
        <mapfield saviyntproperty="entitlementvalues.entowner1" sourceproperty="dataowner1" type="number"></mapfield>
    </mapper>
</dataMapping>
EOF
  )
  user_import = trimspace(<<EOF
  <dataMapping>
    <count-query description="This is the Source Endpoint DB Count Query">
        <![CDATA[
select count(username) as count from mysqllocal.users;
]]>
    </count-query>
    <sql-query description="This is the Source DB Query">
        <![CDATA[select username,'Test-DB-Connector-mysql-SS' as endpoint,firstname,lastname, statuskey from mysqllocal.users;]]>
    </sql-query>
    <importsettings>
        <zeroDayProvisioning>false</zeroDayProvisioning>
        <generateEmail>true</generateEmail>
        <userNotInFileAction>NOACTION</userNotInFileAction>
        <checkRules>false</checkRules>
        <buildUserMap>false</buildUserMap>
        <generateSystemUsername>false</generateSystemUsername>
        <userOperationsAllowed>CREATE,UPDATE</userOperationsAllowed>
        <userReconcillationField>username</userReconcillationField>
    </importsettings>
    <mapper description="This is the mapping field for Saviynt Field name">
        <mapfield saviyntproperty="users.username" sourceproperty="username" type="character"/>
        <mapfield saviyntproperty="users.firstname" sourceproperty="firstname" type="character"></mapfield>
        <mapfield saviyntproperty="users.lastname" sourceproperty="lastname" type="character"></mapfield>
    </mapper>
</dataMapping>
EOF
  )
  role_owner_import = trimspace(<<EOF
  <dataMapping>
    <sql-query description="This is the Source Database Query">
        <![CDATA[select role_name,username,rank from roleowner]]>
    </sql-query>
    <mapper description="This is the mapping field for Saviynt Field name">
        <mapfield saviyntproperty="rolekey" sourceproperty="role_name" type="character"></mapfield>
        <mapfield saviyntproperty="userkey" sourceproperty="username" type="character"></mapfield>
        <mapfield saviyntproperty="rank" sourceproperty="rank" type="number"></mapfield>
    </mapper>
</dataMapping>
EOF
  )
  roles_import = trimspace(<<EOF
  <dataMapping>
    <sql-query description="This is the Source DB Query" uniquecolumnsascommaseparated="role_name">
        <![CDATA[select role_name,customproperty1,description,displayname,customproperty7,sox_critical,sys_critical,customproperty8,customproperty11,customproperty12,customproperty18,customproperty19,customproperty15,customproperty14,mininginstance,customproperty13,roletype,status,customproperty23,customproperty22,customproperty24,customproperty25 from roles]]>
    </sql-query>
    <mapper description="This is the mapping field for Saviynt Field name">
        <mapfield saviyntproperty="role_name" sourceproperty="role_name" type="character"></mapfield>
        <mapfield saviyntproperty="customproperty1" sourceproperty="customproperty1" type="character"></mapfield>
        <mapfield saviyntproperty="description" sourceproperty="description" type="character"></mapfield>
        <mapfield saviyntproperty="displayname" sourceproperty="displayname" type="character"></mapfield>
        <mapfield saviyntproperty="customproperty7" sourceproperty="customproperty7" type="character"></mapfield>
        <mapfield saviyntproperty="sox_critical" sourceproperty="sox_critical" type="character"></mapfield>
        <mapfield saviyntproperty="sys_critical" sourceproperty="sys_critical" type="character"></mapfield>
        <mapfield saviyntproperty="customproperty8" sourceproperty="customproperty8" type="character"></mapfield>
        <mapfield saviyntproperty="customproperty11" sourceproperty="customproperty11" type="character"></mapfield>
        <mapfield saviyntproperty="customproperty12" sourceproperty="customproperty12" type="character"></mapfield>
        <mapfield saviyntproperty="customproperty18" sourceproperty="customproperty18" type="character"></mapfield>
        <mapfield saviyntproperty="customproperty19" sourceproperty="customproperty19" type="character"></mapfield>
        <mapfield saviyntproperty="customproperty15" sourceproperty="customproperty15" type="character"></mapfield>
        <mapfield saviyntproperty="customproperty14" sourceproperty="customproperty14" type="number"></mapfield>
        <mapfield saviyntproperty="mininginstance" sourceproperty="mininginstance" type="character"></mapfield>
        <mapfield saviyntproperty="customproperty13" sourceproperty="customproperty13" type="number"></mapfield>
        <mapfield saviyntproperty="roletype" sourceproperty="roletype" type="number"></mapfield>
        <mapfield saviyntproperty="status" sourceproperty="status" type="number"></mapfield>
        <mapfield saviyntproperty="customproperty23" sourceproperty="customproperty23" type="character"></mapfield>
        <mapfield saviyntproperty="customproperty22" sourceproperty="customproperty22" type="character"></mapfield>
        <mapfield saviyntproperty="customproperty24" sourceproperty="customproperty24" type="character"></mapfield>
        <mapfield saviyntproperty="customproperty25" sourceproperty="customproperty25" type="date"></mapfield>
    </mapper>
</dataMapping>
EOF
  )
  system_import = trimspace(<<EOF
  <dataMapping>
    <sql-query description="This is the Source Database Query" uniquecolumnsascommaseparated="systemname">
        <![CDATA[select name,resourcename,attribute,description from securitysystems]]>
    </sql-query>
    <mapper description="This is the mapping field for Saviynt Field name">
        <mapfield saviyntproperty="securitysystems.systemname" sourceproperty="name" type="character" />
        <mapfield saviyntproperty="endpoints.endpointname" sourceproperty="resourcename" type="character" />
        <mapfield saviyntproperty="entitlementtype.entitlementname" sourceproperty="attribute" type="character" />
        <mapfield saviyntproperty="endpoints.endpointdescription" sourceproperty="description" type="character" />
    </mapper>
</dataMapping>
EOF
  )
  max_pagination_size = "1000"
  cli_command_json = jsonencode(
    {
      "launch" : "mysql -h$${hostip} -P$${port} -u$${hostuser} -p$${password}",
      "parserType" : "Array"
  })
}
