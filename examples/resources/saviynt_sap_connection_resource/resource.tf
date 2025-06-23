// Copyright (c) Saviynt Inc.
// SPDX-License-Identifier: MPL-2.0

resource "saviynt_sap_connection_resource" "ss" {
  connection_type             = "SAP"
  connection_name             = "Terraform_Sap_Connection"
  message_server              = "FALSE"
  jco_ashost                  = "<jco_ashost>"
  jco_sysnr                   = "<jco_sysnr>"
  jco_client                  = "<jco_client>"
  jco_user                    = "<jco_user>"
  password                    = "<password>"
  jco_lang                    = "EN"
  jco_r3name                  = "<jco_r3name>"
  jco_mshost                  = "<jco_mshost>"
  jco_msserv                  = "<jco_msserv>"
  jco_group                   = "<jco_group>"
  snc                         = "false"
  jco_snc_mode                = "0"
  jco_snc_partnername         = ""
  jco_snc_myname              = ""
  jco_snc_library             = ""
  jco_snc_qop                 = ""
  tables                      = "<tables>"
  system_name                 = "<system_name>"
  terminated_user_group       = "SAP_TERMINATED"
  terminated_user_role_action = "REMOVE"
  user_import_json = jsonencode({
    USERTABLEMAPPING = [
      {
        SAPTABLE = "PA0105"
        SAVTABLE = "PA0105"
        SAPMAP   = "PERNR,ENDDA,BEGDA,AEDTM,UNAME,USRTY,USRID,SUBTY,OBJPS,SEQNR"
      },
      {
        SAPTABLE = "PA0001"
        SAVTABLE = "PA0001"
        SAPMAP   = "PERNR,ENDDA,BEGDA,AEDTM,UNAME,BUKRS,WERKS,PERSG,PERSK,VDSK1,BTRTL,KOSTL,ORGEH,PLANS,STELL,SUBTY"
      },
      {
        SAPTABLE = "PA0002"
        SAVTABLE = "PA0002"
        SAPMAP   = "PERNR,ENDDA,BEGDA,AEDTM,UNAME,NACHN,VORNA,TITEL,CNAME"
      },
      {
        SAPTABLE = "PA0000"
        SAVTABLE = "PA0000"
        SAPMAP   = "PERNR,ENDDA,BEGDA,AEDTM,UNAME,STAT2"
      }
    ],
    USERIMPORTQRY          = "SELECT DISTINCT p0.PERNR username,p5.USRID employeeid,CASE WHEN (p0.STAT2 = 3) THEN 1 END statuskey,CASE WHEN (p0.STAT2 = 3) THEN 1 END enabled,p0.ENDDA end_date, p0.BEGDA start_date,p2.NACHN first_name,p2.VORNA last_name,p1.BUKRS company_code,p1.KOSTL cost_center,p1.PERSG EEGroup,p1.PERSK EESubgroup,p1.WERKS PersonnelArea,p1.BTRTL PersonnelSubArea,p1.STELL Job,p1.ORGEH Org_Unit,'SAPHR' as customproperty10 FROM PA0000 p0 left join PA0001 p1 on  p0.PERNR = p1.PERNR left join PA0002 p2 on p1.PERNR = p2.PERNR left join PA0105 p5 on p2.PERNR = p5.PERNR left join (SELECT p51.pernr FROM PA0105 p51 WHERE p51.USRTY = '0010' ) p51out on p51out.pernr = p5.pernr WHERE p0.STAT2 = 3 and p0.ENDDA > SYSDATE() and (p1.endda > SYSDATE() or p1.endda is null) and (p5.USRTY = '0001' or p5.usrty is null)",
    SAVCOLMAPPING          = "username,employeeid, statuskey, enabled, enddate, startdate, firstname, lastname, companyname, costcenter,customproperty1, customproperty2,customproperty3, customproperty4, jobcode,orgunitid,customproperty10",
    userNotInFileAction    = "NOACTION",
    zeroDayProvisioning    = "FALSE",
    generateSystemUsername = "FALSE",
    checkRules             = "FALSE",
    buildUserMap           = "FALSE"
  })
  create_account_json = jsonencode({
    ADDRESS = {
      LASTNAME  = "$${user.lastname}"
      FIRSTNAME = "$${user.firstname}"
      E_MAIL    = "$${user.email}"
    }
    DEFAULTS = {
      DATFM = "2"
      DCPFM = "X"
      SPLD  = "LOCL"
      SPDA  = "D"
    }
    SNC = {
      GUIFLAG = "U"
      PNAME   = "p:$${accountName}@saviynt.com"
    }
    LOGONDATA = {
      USTYP = "A"
    }
    UCLASS = {
      LIC_TYPE = "02"
    }
    PASSWORD = {
      BAPIPWD = "<specify password>"
    }
    REF_USER = {
      REF_USER = "$${referenceid}"
    }
  })
  prov_jco_ashost           = "<prov_jco_ashost>"
  prov_jco_sysnr            = "<prov_jco_sysnr>"
  prov_jco_client           = "<prov_jco_client>"
  prov_jco_user             = "<prov_jco_user>"
  prov_password             = "<prov_password>"
  prov_jco_lang             = "EN"
  prov_jco_r3name           = "QAS"
  prov_jco_mshost           = "<prov_jco_mshost>"
  prov_jco_msserv           = "<prov_jco_msserv>"
  prov_jco_group            = "PROV_GROUP"
  prov_cua_enabled          = "false"
  prov_cua_snc              = "false"
  reset_pwd_for_newaccount  = "true"
  enforce_password_change   = "true"
  password_min_length       = "8"
  password_max_length       = "16"
  password_no_of_caps_alpha = "2"
  password_no_of_digits     = "2"
  password_no_of_spl_chars  = "1"
  hanareftablejson          = "{\"hanaTable\":\"Z_TABLE\"}"
  config_json               = "{\"connectionTimeout\":30,\"retryCount\":3}"
  enable_account_json = jsonencode({
    ADDRESS = {
      LASTNAME  = "$${user.lastname}"
      FIRSTNAME = "$${user.firstname}"
    }
    PASSWORD = {
      BAPIPWD = "$${randomPassword}"
    }
  })
  update_account_json = jsonencode({
    ADDRESS = {
      LASTNAME  = "$${user.lastname}"
      FIRSTNAME = "$${user.firstname}"
      E_MAIL    = "$${user.email}"
    }
    DEFAULTS = {
      DATFM = "2"
      DCPFM = "X"
      SPLD  = "LOCL"
      SPDA  = "D"
    }
    SNC = {
      GUIFLAG = "U"
      PNAME   = "p:$${accountName}@abcdefgh.com"
    }
  })
  status_threshold_config = jsonencode({
    statusAndThresholdConfig = {
      statusColumn                  = "userlock"
      activeStatus                  = [512]
      deleteLinks                   = true
      accountThresholdValue         = 1000
      correlateInactiveAccounts     = true
      inactivateAccountsNotInFile   = false
      deleteAccEntForActiveAccounts = false
      lockedStatusColumn            = "userlock"
      lockedStatusMapping = {
        Locked   = [512]
        Unlocked = [0]
      }
    }
  })
  set_cua_system = "SAP_CUA"
  fire_fighter_id_grant_access_json = jsonencode({
    SNC = {
      GUIFLAG = "U"
      PNAME   = "p:$${user?.username?.toLowerCase()}@ABC1.AD.ABC.COM"
    }
  })
  fire_fighter_id_revoke_access_json = jsonencode({
    PASSWORD = {
      BAPIPWD = "$${randomPassword}"
    }
  })
  external_sod_eval_json        = "{\"EVAL\":\"true\"}"
  external_sod_eval_json_detail = "{\"DETAIL\":\"on\"}"
  logs_table_filter = jsonencode({
    GENERALFILTERS = {
      ACCOUNTTYPE = [
        "A",
      "B"]
      ACCOUNTNAME = [
        "testname1",
      "testname2"]
    }
  })
  saptable_filter_lang               = "EN"
  alternate_output_parameter_et_data = "custom_param"
  audit_log_json                     = "{\"LOG_LEVEL\":\"FULL\"}"
  ecc_or_s4hana                      = "S4HANA"
  data_import_filter = jsonencode({
    accountsFilter = "NAME LIKE 'A%%' AND ACCOUNTTYPE LIKE '%%S'"
  })
}
