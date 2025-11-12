# SAPConnector

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MESSAGESERVER** | Pointer to **string** | set it to TRUE if the Message Server is going to be used for connecting to SAP | [optional] 
**JCO_ASHOST** | Pointer to **string** | HostName for connection for Import | [optional] 
**JCO_SYSNR** | Pointer to **string** | System Number for the SAP Instance for Import | [optional] 
**JCO_CLIENT** | Pointer to **string** | Client Number for the SAP Instance for Import | [optional] 
**JCO_USER** | Pointer to **string** | Username to connect to SAP using JCO for Import | [optional] 
**PASSWORD** | Pointer to **string** | Password for connection for Import | [optional] 
**JCO_LANG** | Pointer to **string** | LANGUAGE for the SAP Instance for Import | [optional] 
**JCOR3NAME** | Pointer to **string** | Message Server SystemID for connection for Import | [optional] 
**JCO_MSHOST** | Pointer to **string** | Message Server Hostname for connection for Import | [optional] 
**JCO_MSSERV** | Pointer to **string** | Message Server Port Number for connection for Import | [optional] 
**JCO_GROUP** | Pointer to **string** | Message Server logon Group connection for Import | [optional] 
**SNC** | Pointer to **string** | TRUE/FALSE, set it to TRUE if Secure Network Connection (SNC) is setup between ABAP and ASHOST | [optional] 
**JCO_SNC_MODE** | Pointer to **string** | Enable/Disable Secure Network Connection (SNC) Mode. 0 - Disabled, 1 - Enabled | [optional] 
**JCO_SNC_PARTNERNAME** | Pointer to **string** | String used to generate secured certificate in SAP server to be used by Saviynt | [optional] 
**JCO_SNC_MYNAME** | Pointer to **string** | String used to generate the secured certificate on the server where Saviynt is deployed | [optional] 
**JCO_SNC_LIBRARY** | Pointer to **string** | Location of SNC library on the Saviynt server | [optional] 
**JCO_SNC_QOP** | Pointer to **string** | The Quality of Protection Level | [optional] 
**TABLES** | Pointer to **string** | Tables to be Imported | [optional] 
**SYSTEMNAME** | Pointer to **string** |  | [optional] 
**TERMINATEDUSERGROUP** | Pointer to **string** | User group for all terminated users | [optional] 
**TERMINATED_USER_ROLE_ACTION** | Pointer to **string** | Action to take for the user roles when the user gets terminated or roles are removed | [optional] 
**CREATEACCOUNTJSON** | Pointer to **string** | JSON to specify the Field Value which will be used to Create the New Account,Objects | [optional] 
**PROV_JCO_ASHOST** | Pointer to **string** | HostName for connection for Provisioning | [optional] 
**PROV_JCO_SYSNR** | Pointer to **string** | System Number for the SAP Instance for Provisioning | [optional] 
**PROV_JCO_CLIENT** | Pointer to **string** | Client Number for the SAP Instance for Provisioning | [optional] 
**PROV_JCO_USER** | Pointer to **string** | Username to connect to to SAP using JCO for Provisioning | [optional] 
**PROV_PASSWORD** | Pointer to **string** | Password for connection for Provisioning | [optional] 
**PROV_JCO_LANG** | Pointer to **string** | LANGUAGE for the SAP Instance for Provisioning | [optional] 
**PROVJCOR3NAME** | Pointer to **string** | Message Server SystemID for connection for Provisioning | [optional] 
**PROV_JCO_MSHOST** | Pointer to **string** | Message Server Hostname for connection for Provisioning | [optional] 
**PROV_JCO_MSSERV** | Pointer to **string** | Message Server Port Number for connection for Provisioning | [optional] 
**PROV_JCO_GROUP** | Pointer to **string** | Message Server logon Group connection for Provisioning | [optional] 
**PROV_CUA_ENABLED** | Pointer to **string** | Is CUA Enabled | [optional] 
**PROV_CUA_SNC** | Pointer to **string** | Property for PROV_CUA_SNC | [optional] 
**RESET_PWD_FOR_NEWACCOUNT** | Pointer to **string** | For CUA enabled if the user already exists should the Password be reset for accounts for other systems | [optional] 
**ENFORCEPASSWORDCHANGE** | Pointer to **string** | Set it to FALSE if while changing the password the productive password should be set | [optional] 
**PASSWORD_MIN_LENGTH** | Pointer to **string** | Specify the Min length for the random password | [optional] 
**PASSWORD_MAX_LENGTH** | Pointer to **string** | Specify the Max length for the random password | [optional] 
**PASSWORD_NOOFCAPSALPHA** | Pointer to **string** | Specify the Number of Upper case alphabets required for the random password | [optional] 
**PASSWORD_NOOFDIGITS** | Pointer to **string** | Specify the Number of digits required for the random password  | [optional] 
**PASSWORD_NOOFSPLCHARS** | Pointer to **string** | Specify the Number of special chars required for the random password | [optional] 
**HANAREFTABLEJSON** | Pointer to **string** | JSON to specify the mapping for Reference table | [optional] 
**ENABLEACCOUNTJSON** | Pointer to **string** | JSON Similar to Create Account to mention the changes that needs to be made at the user level after enabling the account | [optional] 
**UPDATEACCOUNTJSON** | Pointer to **string** | JSON Similar to Create Account to mention the changes that needs to be made at the user level after updating the account | [optional] 
**USERIMPORTJSON** | Pointer to **string** | Property for USERIMPORTJSON | [optional] 
**STATUS_THRESHOLD_CONFIG** | Pointer to **string** | The attributes of statusAndThresholdConfig json | [optional] 
**SETCUASYSTEM** | Pointer to **string** | set it to FALSE if using an older CUA System to not support setting sub-systems | [optional] 
**FIREFIGHTERID_GRANT_ACCESS_JSON** | Pointer to **string** | JSON similar to Update Account to mention SNC and other related changes that needs to be done  | [optional] 
**FIREFIGHTERID_REVOKE_ACCESS_JSON** | Pointer to **string** | JSON similar to Update Account to reset SNC and other related changes that were done during GRANT | [optional] 
**MODIFYUSERDATAJSON** | Pointer to **string** | Property for MODIFYUSERDATAJSON | [optional] 
**EXTERNAL_SOD_EVAL_JSON** | Pointer to **string** | JSON to populate values for External SOD Evaluation | [optional] 
**EXTERNAL_SOD_EVAL_JSON_DETAIL** | Pointer to **string** | Property for EXTERNAL_SOD_EVAL_JSON_DETAIL | [optional] 
**LOGS_TABLE_FILTER** | Pointer to **string** | Property for LOGS_TABLE_FILTER | [optional] 
**PAM_CONFIG** | Pointer to **string** | JSON for PAM Bootstrap Configuration  | [optional] 
**SAPTABLE_FILTER_LANG** | Pointer to **string** | Property for SAPTABLE_FILTER_LANG | [optional] 
**ALTERNATE_OUTPUT_PARAMETER_ET_DATA** | Pointer to **string** | Based on the flag &#39;&#39;USE_ET_DATA_4_RETURN&#39;&#39; in SAP, set this value to make use of export parameter ET_DATA | [optional] 
**AUDIT_LOG_JSON** | Pointer to **string** | Property for AUDIT_LOG_JSON | [optional] 
**ECCORS4HANA** | Pointer to **string** | Property for ECC_OR_S4HANA | [optional] 
**DATA_IMPORT_FILTER** | Pointer to **string** | Property for DATA_IMPORT_FILTER | [optional] 
**ConfigJSON** | Pointer to **string** | Property for ConfigJSON | [optional] 
**ROLE_DEFAULT_DATE** | Pointer to **string** | Default end date for SAP roles. Format: yyyyMMdd (e.g., 20251231). If not configured, defaults to 99991231. | [optional] 

## Methods

### NewSAPConnector

`func NewSAPConnector() *SAPConnector`

NewSAPConnector instantiates a new SAPConnector object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSAPConnectorWithDefaults

`func NewSAPConnectorWithDefaults() *SAPConnector`

NewSAPConnectorWithDefaults instantiates a new SAPConnector object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMESSAGESERVER

`func (o *SAPConnector) GetMESSAGESERVER() string`

GetMESSAGESERVER returns the MESSAGESERVER field if non-nil, zero value otherwise.

### GetMESSAGESERVEROk

`func (o *SAPConnector) GetMESSAGESERVEROk() (*string, bool)`

GetMESSAGESERVEROk returns a tuple with the MESSAGESERVER field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMESSAGESERVER

`func (o *SAPConnector) SetMESSAGESERVER(v string)`

SetMESSAGESERVER sets MESSAGESERVER field to given value.

### HasMESSAGESERVER

`func (o *SAPConnector) HasMESSAGESERVER() bool`

HasMESSAGESERVER returns a boolean if a field has been set.

### GetJCO_ASHOST

`func (o *SAPConnector) GetJCO_ASHOST() string`

GetJCO_ASHOST returns the JCO_ASHOST field if non-nil, zero value otherwise.

### GetJCO_ASHOSTOk

`func (o *SAPConnector) GetJCO_ASHOSTOk() (*string, bool)`

GetJCO_ASHOSTOk returns a tuple with the JCO_ASHOST field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJCO_ASHOST

`func (o *SAPConnector) SetJCO_ASHOST(v string)`

SetJCO_ASHOST sets JCO_ASHOST field to given value.

### HasJCO_ASHOST

`func (o *SAPConnector) HasJCO_ASHOST() bool`

HasJCO_ASHOST returns a boolean if a field has been set.

### GetJCO_SYSNR

`func (o *SAPConnector) GetJCO_SYSNR() string`

GetJCO_SYSNR returns the JCO_SYSNR field if non-nil, zero value otherwise.

### GetJCO_SYSNROk

`func (o *SAPConnector) GetJCO_SYSNROk() (*string, bool)`

GetJCO_SYSNROk returns a tuple with the JCO_SYSNR field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJCO_SYSNR

`func (o *SAPConnector) SetJCO_SYSNR(v string)`

SetJCO_SYSNR sets JCO_SYSNR field to given value.

### HasJCO_SYSNR

`func (o *SAPConnector) HasJCO_SYSNR() bool`

HasJCO_SYSNR returns a boolean if a field has been set.

### GetJCO_CLIENT

`func (o *SAPConnector) GetJCO_CLIENT() string`

GetJCO_CLIENT returns the JCO_CLIENT field if non-nil, zero value otherwise.

### GetJCO_CLIENTOk

`func (o *SAPConnector) GetJCO_CLIENTOk() (*string, bool)`

GetJCO_CLIENTOk returns a tuple with the JCO_CLIENT field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJCO_CLIENT

`func (o *SAPConnector) SetJCO_CLIENT(v string)`

SetJCO_CLIENT sets JCO_CLIENT field to given value.

### HasJCO_CLIENT

`func (o *SAPConnector) HasJCO_CLIENT() bool`

HasJCO_CLIENT returns a boolean if a field has been set.

### GetJCO_USER

`func (o *SAPConnector) GetJCO_USER() string`

GetJCO_USER returns the JCO_USER field if non-nil, zero value otherwise.

### GetJCO_USEROk

`func (o *SAPConnector) GetJCO_USEROk() (*string, bool)`

GetJCO_USEROk returns a tuple with the JCO_USER field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJCO_USER

`func (o *SAPConnector) SetJCO_USER(v string)`

SetJCO_USER sets JCO_USER field to given value.

### HasJCO_USER

`func (o *SAPConnector) HasJCO_USER() bool`

HasJCO_USER returns a boolean if a field has been set.

### GetPASSWORD

`func (o *SAPConnector) GetPASSWORD() string`

GetPASSWORD returns the PASSWORD field if non-nil, zero value otherwise.

### GetPASSWORDOk

`func (o *SAPConnector) GetPASSWORDOk() (*string, bool)`

GetPASSWORDOk returns a tuple with the PASSWORD field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPASSWORD

`func (o *SAPConnector) SetPASSWORD(v string)`

SetPASSWORD sets PASSWORD field to given value.

### HasPASSWORD

`func (o *SAPConnector) HasPASSWORD() bool`

HasPASSWORD returns a boolean if a field has been set.

### GetJCO_LANG

`func (o *SAPConnector) GetJCO_LANG() string`

GetJCO_LANG returns the JCO_LANG field if non-nil, zero value otherwise.

### GetJCO_LANGOk

`func (o *SAPConnector) GetJCO_LANGOk() (*string, bool)`

GetJCO_LANGOk returns a tuple with the JCO_LANG field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJCO_LANG

`func (o *SAPConnector) SetJCO_LANG(v string)`

SetJCO_LANG sets JCO_LANG field to given value.

### HasJCO_LANG

`func (o *SAPConnector) HasJCO_LANG() bool`

HasJCO_LANG returns a boolean if a field has been set.

### GetJCOR3NAME

`func (o *SAPConnector) GetJCOR3NAME() string`

GetJCOR3NAME returns the JCOR3NAME field if non-nil, zero value otherwise.

### GetJCOR3NAMEOk

`func (o *SAPConnector) GetJCOR3NAMEOk() (*string, bool)`

GetJCOR3NAMEOk returns a tuple with the JCOR3NAME field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJCOR3NAME

`func (o *SAPConnector) SetJCOR3NAME(v string)`

SetJCOR3NAME sets JCOR3NAME field to given value.

### HasJCOR3NAME

`func (o *SAPConnector) HasJCOR3NAME() bool`

HasJCOR3NAME returns a boolean if a field has been set.

### GetJCO_MSHOST

`func (o *SAPConnector) GetJCO_MSHOST() string`

GetJCO_MSHOST returns the JCO_MSHOST field if non-nil, zero value otherwise.

### GetJCO_MSHOSTOk

`func (o *SAPConnector) GetJCO_MSHOSTOk() (*string, bool)`

GetJCO_MSHOSTOk returns a tuple with the JCO_MSHOST field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJCO_MSHOST

`func (o *SAPConnector) SetJCO_MSHOST(v string)`

SetJCO_MSHOST sets JCO_MSHOST field to given value.

### HasJCO_MSHOST

`func (o *SAPConnector) HasJCO_MSHOST() bool`

HasJCO_MSHOST returns a boolean if a field has been set.

### GetJCO_MSSERV

`func (o *SAPConnector) GetJCO_MSSERV() string`

GetJCO_MSSERV returns the JCO_MSSERV field if non-nil, zero value otherwise.

### GetJCO_MSSERVOk

`func (o *SAPConnector) GetJCO_MSSERVOk() (*string, bool)`

GetJCO_MSSERVOk returns a tuple with the JCO_MSSERV field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJCO_MSSERV

`func (o *SAPConnector) SetJCO_MSSERV(v string)`

SetJCO_MSSERV sets JCO_MSSERV field to given value.

### HasJCO_MSSERV

`func (o *SAPConnector) HasJCO_MSSERV() bool`

HasJCO_MSSERV returns a boolean if a field has been set.

### GetJCO_GROUP

`func (o *SAPConnector) GetJCO_GROUP() string`

GetJCO_GROUP returns the JCO_GROUP field if non-nil, zero value otherwise.

### GetJCO_GROUPOk

`func (o *SAPConnector) GetJCO_GROUPOk() (*string, bool)`

GetJCO_GROUPOk returns a tuple with the JCO_GROUP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJCO_GROUP

`func (o *SAPConnector) SetJCO_GROUP(v string)`

SetJCO_GROUP sets JCO_GROUP field to given value.

### HasJCO_GROUP

`func (o *SAPConnector) HasJCO_GROUP() bool`

HasJCO_GROUP returns a boolean if a field has been set.

### GetSNC

`func (o *SAPConnector) GetSNC() string`

GetSNC returns the SNC field if non-nil, zero value otherwise.

### GetSNCOk

`func (o *SAPConnector) GetSNCOk() (*string, bool)`

GetSNCOk returns a tuple with the SNC field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSNC

`func (o *SAPConnector) SetSNC(v string)`

SetSNC sets SNC field to given value.

### HasSNC

`func (o *SAPConnector) HasSNC() bool`

HasSNC returns a boolean if a field has been set.

### GetJCO_SNC_MODE

`func (o *SAPConnector) GetJCO_SNC_MODE() string`

GetJCO_SNC_MODE returns the JCO_SNC_MODE field if non-nil, zero value otherwise.

### GetJCO_SNC_MODEOk

`func (o *SAPConnector) GetJCO_SNC_MODEOk() (*string, bool)`

GetJCO_SNC_MODEOk returns a tuple with the JCO_SNC_MODE field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJCO_SNC_MODE

`func (o *SAPConnector) SetJCO_SNC_MODE(v string)`

SetJCO_SNC_MODE sets JCO_SNC_MODE field to given value.

### HasJCO_SNC_MODE

`func (o *SAPConnector) HasJCO_SNC_MODE() bool`

HasJCO_SNC_MODE returns a boolean if a field has been set.

### GetJCO_SNC_PARTNERNAME

`func (o *SAPConnector) GetJCO_SNC_PARTNERNAME() string`

GetJCO_SNC_PARTNERNAME returns the JCO_SNC_PARTNERNAME field if non-nil, zero value otherwise.

### GetJCO_SNC_PARTNERNAMEOk

`func (o *SAPConnector) GetJCO_SNC_PARTNERNAMEOk() (*string, bool)`

GetJCO_SNC_PARTNERNAMEOk returns a tuple with the JCO_SNC_PARTNERNAME field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJCO_SNC_PARTNERNAME

`func (o *SAPConnector) SetJCO_SNC_PARTNERNAME(v string)`

SetJCO_SNC_PARTNERNAME sets JCO_SNC_PARTNERNAME field to given value.

### HasJCO_SNC_PARTNERNAME

`func (o *SAPConnector) HasJCO_SNC_PARTNERNAME() bool`

HasJCO_SNC_PARTNERNAME returns a boolean if a field has been set.

### GetJCO_SNC_MYNAME

`func (o *SAPConnector) GetJCO_SNC_MYNAME() string`

GetJCO_SNC_MYNAME returns the JCO_SNC_MYNAME field if non-nil, zero value otherwise.

### GetJCO_SNC_MYNAMEOk

`func (o *SAPConnector) GetJCO_SNC_MYNAMEOk() (*string, bool)`

GetJCO_SNC_MYNAMEOk returns a tuple with the JCO_SNC_MYNAME field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJCO_SNC_MYNAME

`func (o *SAPConnector) SetJCO_SNC_MYNAME(v string)`

SetJCO_SNC_MYNAME sets JCO_SNC_MYNAME field to given value.

### HasJCO_SNC_MYNAME

`func (o *SAPConnector) HasJCO_SNC_MYNAME() bool`

HasJCO_SNC_MYNAME returns a boolean if a field has been set.

### GetJCO_SNC_LIBRARY

`func (o *SAPConnector) GetJCO_SNC_LIBRARY() string`

GetJCO_SNC_LIBRARY returns the JCO_SNC_LIBRARY field if non-nil, zero value otherwise.

### GetJCO_SNC_LIBRARYOk

`func (o *SAPConnector) GetJCO_SNC_LIBRARYOk() (*string, bool)`

GetJCO_SNC_LIBRARYOk returns a tuple with the JCO_SNC_LIBRARY field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJCO_SNC_LIBRARY

`func (o *SAPConnector) SetJCO_SNC_LIBRARY(v string)`

SetJCO_SNC_LIBRARY sets JCO_SNC_LIBRARY field to given value.

### HasJCO_SNC_LIBRARY

`func (o *SAPConnector) HasJCO_SNC_LIBRARY() bool`

HasJCO_SNC_LIBRARY returns a boolean if a field has been set.

### GetJCO_SNC_QOP

`func (o *SAPConnector) GetJCO_SNC_QOP() string`

GetJCO_SNC_QOP returns the JCO_SNC_QOP field if non-nil, zero value otherwise.

### GetJCO_SNC_QOPOk

`func (o *SAPConnector) GetJCO_SNC_QOPOk() (*string, bool)`

GetJCO_SNC_QOPOk returns a tuple with the JCO_SNC_QOP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJCO_SNC_QOP

`func (o *SAPConnector) SetJCO_SNC_QOP(v string)`

SetJCO_SNC_QOP sets JCO_SNC_QOP field to given value.

### HasJCO_SNC_QOP

`func (o *SAPConnector) HasJCO_SNC_QOP() bool`

HasJCO_SNC_QOP returns a boolean if a field has been set.

### GetTABLES

`func (o *SAPConnector) GetTABLES() string`

GetTABLES returns the TABLES field if non-nil, zero value otherwise.

### GetTABLESOk

`func (o *SAPConnector) GetTABLESOk() (*string, bool)`

GetTABLESOk returns a tuple with the TABLES field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTABLES

`func (o *SAPConnector) SetTABLES(v string)`

SetTABLES sets TABLES field to given value.

### HasTABLES

`func (o *SAPConnector) HasTABLES() bool`

HasTABLES returns a boolean if a field has been set.

### GetSYSTEMNAME

`func (o *SAPConnector) GetSYSTEMNAME() string`

GetSYSTEMNAME returns the SYSTEMNAME field if non-nil, zero value otherwise.

### GetSYSTEMNAMEOk

`func (o *SAPConnector) GetSYSTEMNAMEOk() (*string, bool)`

GetSYSTEMNAMEOk returns a tuple with the SYSTEMNAME field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSYSTEMNAME

`func (o *SAPConnector) SetSYSTEMNAME(v string)`

SetSYSTEMNAME sets SYSTEMNAME field to given value.

### HasSYSTEMNAME

`func (o *SAPConnector) HasSYSTEMNAME() bool`

HasSYSTEMNAME returns a boolean if a field has been set.

### GetTERMINATEDUSERGROUP

`func (o *SAPConnector) GetTERMINATEDUSERGROUP() string`

GetTERMINATEDUSERGROUP returns the TERMINATEDUSERGROUP field if non-nil, zero value otherwise.

### GetTERMINATEDUSERGROUPOk

`func (o *SAPConnector) GetTERMINATEDUSERGROUPOk() (*string, bool)`

GetTERMINATEDUSERGROUPOk returns a tuple with the TERMINATEDUSERGROUP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTERMINATEDUSERGROUP

`func (o *SAPConnector) SetTERMINATEDUSERGROUP(v string)`

SetTERMINATEDUSERGROUP sets TERMINATEDUSERGROUP field to given value.

### HasTERMINATEDUSERGROUP

`func (o *SAPConnector) HasTERMINATEDUSERGROUP() bool`

HasTERMINATEDUSERGROUP returns a boolean if a field has been set.

### GetTERMINATED_USER_ROLE_ACTION

`func (o *SAPConnector) GetTERMINATED_USER_ROLE_ACTION() string`

GetTERMINATED_USER_ROLE_ACTION returns the TERMINATED_USER_ROLE_ACTION field if non-nil, zero value otherwise.

### GetTERMINATED_USER_ROLE_ACTIONOk

`func (o *SAPConnector) GetTERMINATED_USER_ROLE_ACTIONOk() (*string, bool)`

GetTERMINATED_USER_ROLE_ACTIONOk returns a tuple with the TERMINATED_USER_ROLE_ACTION field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTERMINATED_USER_ROLE_ACTION

`func (o *SAPConnector) SetTERMINATED_USER_ROLE_ACTION(v string)`

SetTERMINATED_USER_ROLE_ACTION sets TERMINATED_USER_ROLE_ACTION field to given value.

### HasTERMINATED_USER_ROLE_ACTION

`func (o *SAPConnector) HasTERMINATED_USER_ROLE_ACTION() bool`

HasTERMINATED_USER_ROLE_ACTION returns a boolean if a field has been set.

### GetCREATEACCOUNTJSON

`func (o *SAPConnector) GetCREATEACCOUNTJSON() string`

GetCREATEACCOUNTJSON returns the CREATEACCOUNTJSON field if non-nil, zero value otherwise.

### GetCREATEACCOUNTJSONOk

`func (o *SAPConnector) GetCREATEACCOUNTJSONOk() (*string, bool)`

GetCREATEACCOUNTJSONOk returns a tuple with the CREATEACCOUNTJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCREATEACCOUNTJSON

`func (o *SAPConnector) SetCREATEACCOUNTJSON(v string)`

SetCREATEACCOUNTJSON sets CREATEACCOUNTJSON field to given value.

### HasCREATEACCOUNTJSON

`func (o *SAPConnector) HasCREATEACCOUNTJSON() bool`

HasCREATEACCOUNTJSON returns a boolean if a field has been set.

### GetPROV_JCO_ASHOST

`func (o *SAPConnector) GetPROV_JCO_ASHOST() string`

GetPROV_JCO_ASHOST returns the PROV_JCO_ASHOST field if non-nil, zero value otherwise.

### GetPROV_JCO_ASHOSTOk

`func (o *SAPConnector) GetPROV_JCO_ASHOSTOk() (*string, bool)`

GetPROV_JCO_ASHOSTOk returns a tuple with the PROV_JCO_ASHOST field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPROV_JCO_ASHOST

`func (o *SAPConnector) SetPROV_JCO_ASHOST(v string)`

SetPROV_JCO_ASHOST sets PROV_JCO_ASHOST field to given value.

### HasPROV_JCO_ASHOST

`func (o *SAPConnector) HasPROV_JCO_ASHOST() bool`

HasPROV_JCO_ASHOST returns a boolean if a field has been set.

### GetPROV_JCO_SYSNR

`func (o *SAPConnector) GetPROV_JCO_SYSNR() string`

GetPROV_JCO_SYSNR returns the PROV_JCO_SYSNR field if non-nil, zero value otherwise.

### GetPROV_JCO_SYSNROk

`func (o *SAPConnector) GetPROV_JCO_SYSNROk() (*string, bool)`

GetPROV_JCO_SYSNROk returns a tuple with the PROV_JCO_SYSNR field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPROV_JCO_SYSNR

`func (o *SAPConnector) SetPROV_JCO_SYSNR(v string)`

SetPROV_JCO_SYSNR sets PROV_JCO_SYSNR field to given value.

### HasPROV_JCO_SYSNR

`func (o *SAPConnector) HasPROV_JCO_SYSNR() bool`

HasPROV_JCO_SYSNR returns a boolean if a field has been set.

### GetPROV_JCO_CLIENT

`func (o *SAPConnector) GetPROV_JCO_CLIENT() string`

GetPROV_JCO_CLIENT returns the PROV_JCO_CLIENT field if non-nil, zero value otherwise.

### GetPROV_JCO_CLIENTOk

`func (o *SAPConnector) GetPROV_JCO_CLIENTOk() (*string, bool)`

GetPROV_JCO_CLIENTOk returns a tuple with the PROV_JCO_CLIENT field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPROV_JCO_CLIENT

`func (o *SAPConnector) SetPROV_JCO_CLIENT(v string)`

SetPROV_JCO_CLIENT sets PROV_JCO_CLIENT field to given value.

### HasPROV_JCO_CLIENT

`func (o *SAPConnector) HasPROV_JCO_CLIENT() bool`

HasPROV_JCO_CLIENT returns a boolean if a field has been set.

### GetPROV_JCO_USER

`func (o *SAPConnector) GetPROV_JCO_USER() string`

GetPROV_JCO_USER returns the PROV_JCO_USER field if non-nil, zero value otherwise.

### GetPROV_JCO_USEROk

`func (o *SAPConnector) GetPROV_JCO_USEROk() (*string, bool)`

GetPROV_JCO_USEROk returns a tuple with the PROV_JCO_USER field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPROV_JCO_USER

`func (o *SAPConnector) SetPROV_JCO_USER(v string)`

SetPROV_JCO_USER sets PROV_JCO_USER field to given value.

### HasPROV_JCO_USER

`func (o *SAPConnector) HasPROV_JCO_USER() bool`

HasPROV_JCO_USER returns a boolean if a field has been set.

### GetPROV_PASSWORD

`func (o *SAPConnector) GetPROV_PASSWORD() string`

GetPROV_PASSWORD returns the PROV_PASSWORD field if non-nil, zero value otherwise.

### GetPROV_PASSWORDOk

`func (o *SAPConnector) GetPROV_PASSWORDOk() (*string, bool)`

GetPROV_PASSWORDOk returns a tuple with the PROV_PASSWORD field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPROV_PASSWORD

`func (o *SAPConnector) SetPROV_PASSWORD(v string)`

SetPROV_PASSWORD sets PROV_PASSWORD field to given value.

### HasPROV_PASSWORD

`func (o *SAPConnector) HasPROV_PASSWORD() bool`

HasPROV_PASSWORD returns a boolean if a field has been set.

### GetPROV_JCO_LANG

`func (o *SAPConnector) GetPROV_JCO_LANG() string`

GetPROV_JCO_LANG returns the PROV_JCO_LANG field if non-nil, zero value otherwise.

### GetPROV_JCO_LANGOk

`func (o *SAPConnector) GetPROV_JCO_LANGOk() (*string, bool)`

GetPROV_JCO_LANGOk returns a tuple with the PROV_JCO_LANG field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPROV_JCO_LANG

`func (o *SAPConnector) SetPROV_JCO_LANG(v string)`

SetPROV_JCO_LANG sets PROV_JCO_LANG field to given value.

### HasPROV_JCO_LANG

`func (o *SAPConnector) HasPROV_JCO_LANG() bool`

HasPROV_JCO_LANG returns a boolean if a field has been set.

### GetPROVJCOR3NAME

`func (o *SAPConnector) GetPROVJCOR3NAME() string`

GetPROVJCOR3NAME returns the PROVJCOR3NAME field if non-nil, zero value otherwise.

### GetPROVJCOR3NAMEOk

`func (o *SAPConnector) GetPROVJCOR3NAMEOk() (*string, bool)`

GetPROVJCOR3NAMEOk returns a tuple with the PROVJCOR3NAME field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPROVJCOR3NAME

`func (o *SAPConnector) SetPROVJCOR3NAME(v string)`

SetPROVJCOR3NAME sets PROVJCOR3NAME field to given value.

### HasPROVJCOR3NAME

`func (o *SAPConnector) HasPROVJCOR3NAME() bool`

HasPROVJCOR3NAME returns a boolean if a field has been set.

### GetPROV_JCO_MSHOST

`func (o *SAPConnector) GetPROV_JCO_MSHOST() string`

GetPROV_JCO_MSHOST returns the PROV_JCO_MSHOST field if non-nil, zero value otherwise.

### GetPROV_JCO_MSHOSTOk

`func (o *SAPConnector) GetPROV_JCO_MSHOSTOk() (*string, bool)`

GetPROV_JCO_MSHOSTOk returns a tuple with the PROV_JCO_MSHOST field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPROV_JCO_MSHOST

`func (o *SAPConnector) SetPROV_JCO_MSHOST(v string)`

SetPROV_JCO_MSHOST sets PROV_JCO_MSHOST field to given value.

### HasPROV_JCO_MSHOST

`func (o *SAPConnector) HasPROV_JCO_MSHOST() bool`

HasPROV_JCO_MSHOST returns a boolean if a field has been set.

### GetPROV_JCO_MSSERV

`func (o *SAPConnector) GetPROV_JCO_MSSERV() string`

GetPROV_JCO_MSSERV returns the PROV_JCO_MSSERV field if non-nil, zero value otherwise.

### GetPROV_JCO_MSSERVOk

`func (o *SAPConnector) GetPROV_JCO_MSSERVOk() (*string, bool)`

GetPROV_JCO_MSSERVOk returns a tuple with the PROV_JCO_MSSERV field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPROV_JCO_MSSERV

`func (o *SAPConnector) SetPROV_JCO_MSSERV(v string)`

SetPROV_JCO_MSSERV sets PROV_JCO_MSSERV field to given value.

### HasPROV_JCO_MSSERV

`func (o *SAPConnector) HasPROV_JCO_MSSERV() bool`

HasPROV_JCO_MSSERV returns a boolean if a field has been set.

### GetPROV_JCO_GROUP

`func (o *SAPConnector) GetPROV_JCO_GROUP() string`

GetPROV_JCO_GROUP returns the PROV_JCO_GROUP field if non-nil, zero value otherwise.

### GetPROV_JCO_GROUPOk

`func (o *SAPConnector) GetPROV_JCO_GROUPOk() (*string, bool)`

GetPROV_JCO_GROUPOk returns a tuple with the PROV_JCO_GROUP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPROV_JCO_GROUP

`func (o *SAPConnector) SetPROV_JCO_GROUP(v string)`

SetPROV_JCO_GROUP sets PROV_JCO_GROUP field to given value.

### HasPROV_JCO_GROUP

`func (o *SAPConnector) HasPROV_JCO_GROUP() bool`

HasPROV_JCO_GROUP returns a boolean if a field has been set.

### GetPROV_CUA_ENABLED

`func (o *SAPConnector) GetPROV_CUA_ENABLED() string`

GetPROV_CUA_ENABLED returns the PROV_CUA_ENABLED field if non-nil, zero value otherwise.

### GetPROV_CUA_ENABLEDOk

`func (o *SAPConnector) GetPROV_CUA_ENABLEDOk() (*string, bool)`

GetPROV_CUA_ENABLEDOk returns a tuple with the PROV_CUA_ENABLED field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPROV_CUA_ENABLED

`func (o *SAPConnector) SetPROV_CUA_ENABLED(v string)`

SetPROV_CUA_ENABLED sets PROV_CUA_ENABLED field to given value.

### HasPROV_CUA_ENABLED

`func (o *SAPConnector) HasPROV_CUA_ENABLED() bool`

HasPROV_CUA_ENABLED returns a boolean if a field has been set.

### GetPROV_CUA_SNC

`func (o *SAPConnector) GetPROV_CUA_SNC() string`

GetPROV_CUA_SNC returns the PROV_CUA_SNC field if non-nil, zero value otherwise.

### GetPROV_CUA_SNCOk

`func (o *SAPConnector) GetPROV_CUA_SNCOk() (*string, bool)`

GetPROV_CUA_SNCOk returns a tuple with the PROV_CUA_SNC field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPROV_CUA_SNC

`func (o *SAPConnector) SetPROV_CUA_SNC(v string)`

SetPROV_CUA_SNC sets PROV_CUA_SNC field to given value.

### HasPROV_CUA_SNC

`func (o *SAPConnector) HasPROV_CUA_SNC() bool`

HasPROV_CUA_SNC returns a boolean if a field has been set.

### GetRESET_PWD_FOR_NEWACCOUNT

`func (o *SAPConnector) GetRESET_PWD_FOR_NEWACCOUNT() string`

GetRESET_PWD_FOR_NEWACCOUNT returns the RESET_PWD_FOR_NEWACCOUNT field if non-nil, zero value otherwise.

### GetRESET_PWD_FOR_NEWACCOUNTOk

`func (o *SAPConnector) GetRESET_PWD_FOR_NEWACCOUNTOk() (*string, bool)`

GetRESET_PWD_FOR_NEWACCOUNTOk returns a tuple with the RESET_PWD_FOR_NEWACCOUNT field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRESET_PWD_FOR_NEWACCOUNT

`func (o *SAPConnector) SetRESET_PWD_FOR_NEWACCOUNT(v string)`

SetRESET_PWD_FOR_NEWACCOUNT sets RESET_PWD_FOR_NEWACCOUNT field to given value.

### HasRESET_PWD_FOR_NEWACCOUNT

`func (o *SAPConnector) HasRESET_PWD_FOR_NEWACCOUNT() bool`

HasRESET_PWD_FOR_NEWACCOUNT returns a boolean if a field has been set.

### GetENFORCEPASSWORDCHANGE

`func (o *SAPConnector) GetENFORCEPASSWORDCHANGE() string`

GetENFORCEPASSWORDCHANGE returns the ENFORCEPASSWORDCHANGE field if non-nil, zero value otherwise.

### GetENFORCEPASSWORDCHANGEOk

`func (o *SAPConnector) GetENFORCEPASSWORDCHANGEOk() (*string, bool)`

GetENFORCEPASSWORDCHANGEOk returns a tuple with the ENFORCEPASSWORDCHANGE field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetENFORCEPASSWORDCHANGE

`func (o *SAPConnector) SetENFORCEPASSWORDCHANGE(v string)`

SetENFORCEPASSWORDCHANGE sets ENFORCEPASSWORDCHANGE field to given value.

### HasENFORCEPASSWORDCHANGE

`func (o *SAPConnector) HasENFORCEPASSWORDCHANGE() bool`

HasENFORCEPASSWORDCHANGE returns a boolean if a field has been set.

### GetPASSWORD_MIN_LENGTH

`func (o *SAPConnector) GetPASSWORD_MIN_LENGTH() string`

GetPASSWORD_MIN_LENGTH returns the PASSWORD_MIN_LENGTH field if non-nil, zero value otherwise.

### GetPASSWORD_MIN_LENGTHOk

`func (o *SAPConnector) GetPASSWORD_MIN_LENGTHOk() (*string, bool)`

GetPASSWORD_MIN_LENGTHOk returns a tuple with the PASSWORD_MIN_LENGTH field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPASSWORD_MIN_LENGTH

`func (o *SAPConnector) SetPASSWORD_MIN_LENGTH(v string)`

SetPASSWORD_MIN_LENGTH sets PASSWORD_MIN_LENGTH field to given value.

### HasPASSWORD_MIN_LENGTH

`func (o *SAPConnector) HasPASSWORD_MIN_LENGTH() bool`

HasPASSWORD_MIN_LENGTH returns a boolean if a field has been set.

### GetPASSWORD_MAX_LENGTH

`func (o *SAPConnector) GetPASSWORD_MAX_LENGTH() string`

GetPASSWORD_MAX_LENGTH returns the PASSWORD_MAX_LENGTH field if non-nil, zero value otherwise.

### GetPASSWORD_MAX_LENGTHOk

`func (o *SAPConnector) GetPASSWORD_MAX_LENGTHOk() (*string, bool)`

GetPASSWORD_MAX_LENGTHOk returns a tuple with the PASSWORD_MAX_LENGTH field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPASSWORD_MAX_LENGTH

`func (o *SAPConnector) SetPASSWORD_MAX_LENGTH(v string)`

SetPASSWORD_MAX_LENGTH sets PASSWORD_MAX_LENGTH field to given value.

### HasPASSWORD_MAX_LENGTH

`func (o *SAPConnector) HasPASSWORD_MAX_LENGTH() bool`

HasPASSWORD_MAX_LENGTH returns a boolean if a field has been set.

### GetPASSWORD_NOOFCAPSALPHA

`func (o *SAPConnector) GetPASSWORD_NOOFCAPSALPHA() string`

GetPASSWORD_NOOFCAPSALPHA returns the PASSWORD_NOOFCAPSALPHA field if non-nil, zero value otherwise.

### GetPASSWORD_NOOFCAPSALPHAOk

`func (o *SAPConnector) GetPASSWORD_NOOFCAPSALPHAOk() (*string, bool)`

GetPASSWORD_NOOFCAPSALPHAOk returns a tuple with the PASSWORD_NOOFCAPSALPHA field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPASSWORD_NOOFCAPSALPHA

`func (o *SAPConnector) SetPASSWORD_NOOFCAPSALPHA(v string)`

SetPASSWORD_NOOFCAPSALPHA sets PASSWORD_NOOFCAPSALPHA field to given value.

### HasPASSWORD_NOOFCAPSALPHA

`func (o *SAPConnector) HasPASSWORD_NOOFCAPSALPHA() bool`

HasPASSWORD_NOOFCAPSALPHA returns a boolean if a field has been set.

### GetPASSWORD_NOOFDIGITS

`func (o *SAPConnector) GetPASSWORD_NOOFDIGITS() string`

GetPASSWORD_NOOFDIGITS returns the PASSWORD_NOOFDIGITS field if non-nil, zero value otherwise.

### GetPASSWORD_NOOFDIGITSOk

`func (o *SAPConnector) GetPASSWORD_NOOFDIGITSOk() (*string, bool)`

GetPASSWORD_NOOFDIGITSOk returns a tuple with the PASSWORD_NOOFDIGITS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPASSWORD_NOOFDIGITS

`func (o *SAPConnector) SetPASSWORD_NOOFDIGITS(v string)`

SetPASSWORD_NOOFDIGITS sets PASSWORD_NOOFDIGITS field to given value.

### HasPASSWORD_NOOFDIGITS

`func (o *SAPConnector) HasPASSWORD_NOOFDIGITS() bool`

HasPASSWORD_NOOFDIGITS returns a boolean if a field has been set.

### GetPASSWORD_NOOFSPLCHARS

`func (o *SAPConnector) GetPASSWORD_NOOFSPLCHARS() string`

GetPASSWORD_NOOFSPLCHARS returns the PASSWORD_NOOFSPLCHARS field if non-nil, zero value otherwise.

### GetPASSWORD_NOOFSPLCHARSOk

`func (o *SAPConnector) GetPASSWORD_NOOFSPLCHARSOk() (*string, bool)`

GetPASSWORD_NOOFSPLCHARSOk returns a tuple with the PASSWORD_NOOFSPLCHARS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPASSWORD_NOOFSPLCHARS

`func (o *SAPConnector) SetPASSWORD_NOOFSPLCHARS(v string)`

SetPASSWORD_NOOFSPLCHARS sets PASSWORD_NOOFSPLCHARS field to given value.

### HasPASSWORD_NOOFSPLCHARS

`func (o *SAPConnector) HasPASSWORD_NOOFSPLCHARS() bool`

HasPASSWORD_NOOFSPLCHARS returns a boolean if a field has been set.

### GetHANAREFTABLEJSON

`func (o *SAPConnector) GetHANAREFTABLEJSON() string`

GetHANAREFTABLEJSON returns the HANAREFTABLEJSON field if non-nil, zero value otherwise.

### GetHANAREFTABLEJSONOk

`func (o *SAPConnector) GetHANAREFTABLEJSONOk() (*string, bool)`

GetHANAREFTABLEJSONOk returns a tuple with the HANAREFTABLEJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHANAREFTABLEJSON

`func (o *SAPConnector) SetHANAREFTABLEJSON(v string)`

SetHANAREFTABLEJSON sets HANAREFTABLEJSON field to given value.

### HasHANAREFTABLEJSON

`func (o *SAPConnector) HasHANAREFTABLEJSON() bool`

HasHANAREFTABLEJSON returns a boolean if a field has been set.

### GetENABLEACCOUNTJSON

`func (o *SAPConnector) GetENABLEACCOUNTJSON() string`

GetENABLEACCOUNTJSON returns the ENABLEACCOUNTJSON field if non-nil, zero value otherwise.

### GetENABLEACCOUNTJSONOk

`func (o *SAPConnector) GetENABLEACCOUNTJSONOk() (*string, bool)`

GetENABLEACCOUNTJSONOk returns a tuple with the ENABLEACCOUNTJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetENABLEACCOUNTJSON

`func (o *SAPConnector) SetENABLEACCOUNTJSON(v string)`

SetENABLEACCOUNTJSON sets ENABLEACCOUNTJSON field to given value.

### HasENABLEACCOUNTJSON

`func (o *SAPConnector) HasENABLEACCOUNTJSON() bool`

HasENABLEACCOUNTJSON returns a boolean if a field has been set.

### GetUPDATEACCOUNTJSON

`func (o *SAPConnector) GetUPDATEACCOUNTJSON() string`

GetUPDATEACCOUNTJSON returns the UPDATEACCOUNTJSON field if non-nil, zero value otherwise.

### GetUPDATEACCOUNTJSONOk

`func (o *SAPConnector) GetUPDATEACCOUNTJSONOk() (*string, bool)`

GetUPDATEACCOUNTJSONOk returns a tuple with the UPDATEACCOUNTJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUPDATEACCOUNTJSON

`func (o *SAPConnector) SetUPDATEACCOUNTJSON(v string)`

SetUPDATEACCOUNTJSON sets UPDATEACCOUNTJSON field to given value.

### HasUPDATEACCOUNTJSON

`func (o *SAPConnector) HasUPDATEACCOUNTJSON() bool`

HasUPDATEACCOUNTJSON returns a boolean if a field has been set.

### GetUSERIMPORTJSON

`func (o *SAPConnector) GetUSERIMPORTJSON() string`

GetUSERIMPORTJSON returns the USERIMPORTJSON field if non-nil, zero value otherwise.

### GetUSERIMPORTJSONOk

`func (o *SAPConnector) GetUSERIMPORTJSONOk() (*string, bool)`

GetUSERIMPORTJSONOk returns a tuple with the USERIMPORTJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUSERIMPORTJSON

`func (o *SAPConnector) SetUSERIMPORTJSON(v string)`

SetUSERIMPORTJSON sets USERIMPORTJSON field to given value.

### HasUSERIMPORTJSON

`func (o *SAPConnector) HasUSERIMPORTJSON() bool`

HasUSERIMPORTJSON returns a boolean if a field has been set.

### GetSTATUS_THRESHOLD_CONFIG

`func (o *SAPConnector) GetSTATUS_THRESHOLD_CONFIG() string`

GetSTATUS_THRESHOLD_CONFIG returns the STATUS_THRESHOLD_CONFIG field if non-nil, zero value otherwise.

### GetSTATUS_THRESHOLD_CONFIGOk

`func (o *SAPConnector) GetSTATUS_THRESHOLD_CONFIGOk() (*string, bool)`

GetSTATUS_THRESHOLD_CONFIGOk returns a tuple with the STATUS_THRESHOLD_CONFIG field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSTATUS_THRESHOLD_CONFIG

`func (o *SAPConnector) SetSTATUS_THRESHOLD_CONFIG(v string)`

SetSTATUS_THRESHOLD_CONFIG sets STATUS_THRESHOLD_CONFIG field to given value.

### HasSTATUS_THRESHOLD_CONFIG

`func (o *SAPConnector) HasSTATUS_THRESHOLD_CONFIG() bool`

HasSTATUS_THRESHOLD_CONFIG returns a boolean if a field has been set.

### GetSETCUASYSTEM

`func (o *SAPConnector) GetSETCUASYSTEM() string`

GetSETCUASYSTEM returns the SETCUASYSTEM field if non-nil, zero value otherwise.

### GetSETCUASYSTEMOk

`func (o *SAPConnector) GetSETCUASYSTEMOk() (*string, bool)`

GetSETCUASYSTEMOk returns a tuple with the SETCUASYSTEM field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSETCUASYSTEM

`func (o *SAPConnector) SetSETCUASYSTEM(v string)`

SetSETCUASYSTEM sets SETCUASYSTEM field to given value.

### HasSETCUASYSTEM

`func (o *SAPConnector) HasSETCUASYSTEM() bool`

HasSETCUASYSTEM returns a boolean if a field has been set.

### GetFIREFIGHTERID_GRANT_ACCESS_JSON

`func (o *SAPConnector) GetFIREFIGHTERID_GRANT_ACCESS_JSON() string`

GetFIREFIGHTERID_GRANT_ACCESS_JSON returns the FIREFIGHTERID_GRANT_ACCESS_JSON field if non-nil, zero value otherwise.

### GetFIREFIGHTERID_GRANT_ACCESS_JSONOk

`func (o *SAPConnector) GetFIREFIGHTERID_GRANT_ACCESS_JSONOk() (*string, bool)`

GetFIREFIGHTERID_GRANT_ACCESS_JSONOk returns a tuple with the FIREFIGHTERID_GRANT_ACCESS_JSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFIREFIGHTERID_GRANT_ACCESS_JSON

`func (o *SAPConnector) SetFIREFIGHTERID_GRANT_ACCESS_JSON(v string)`

SetFIREFIGHTERID_GRANT_ACCESS_JSON sets FIREFIGHTERID_GRANT_ACCESS_JSON field to given value.

### HasFIREFIGHTERID_GRANT_ACCESS_JSON

`func (o *SAPConnector) HasFIREFIGHTERID_GRANT_ACCESS_JSON() bool`

HasFIREFIGHTERID_GRANT_ACCESS_JSON returns a boolean if a field has been set.

### GetFIREFIGHTERID_REVOKE_ACCESS_JSON

`func (o *SAPConnector) GetFIREFIGHTERID_REVOKE_ACCESS_JSON() string`

GetFIREFIGHTERID_REVOKE_ACCESS_JSON returns the FIREFIGHTERID_REVOKE_ACCESS_JSON field if non-nil, zero value otherwise.

### GetFIREFIGHTERID_REVOKE_ACCESS_JSONOk

`func (o *SAPConnector) GetFIREFIGHTERID_REVOKE_ACCESS_JSONOk() (*string, bool)`

GetFIREFIGHTERID_REVOKE_ACCESS_JSONOk returns a tuple with the FIREFIGHTERID_REVOKE_ACCESS_JSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFIREFIGHTERID_REVOKE_ACCESS_JSON

`func (o *SAPConnector) SetFIREFIGHTERID_REVOKE_ACCESS_JSON(v string)`

SetFIREFIGHTERID_REVOKE_ACCESS_JSON sets FIREFIGHTERID_REVOKE_ACCESS_JSON field to given value.

### HasFIREFIGHTERID_REVOKE_ACCESS_JSON

`func (o *SAPConnector) HasFIREFIGHTERID_REVOKE_ACCESS_JSON() bool`

HasFIREFIGHTERID_REVOKE_ACCESS_JSON returns a boolean if a field has been set.

### GetMODIFYUSERDATAJSON

`func (o *SAPConnector) GetMODIFYUSERDATAJSON() string`

GetMODIFYUSERDATAJSON returns the MODIFYUSERDATAJSON field if non-nil, zero value otherwise.

### GetMODIFYUSERDATAJSONOk

`func (o *SAPConnector) GetMODIFYUSERDATAJSONOk() (*string, bool)`

GetMODIFYUSERDATAJSONOk returns a tuple with the MODIFYUSERDATAJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMODIFYUSERDATAJSON

`func (o *SAPConnector) SetMODIFYUSERDATAJSON(v string)`

SetMODIFYUSERDATAJSON sets MODIFYUSERDATAJSON field to given value.

### HasMODIFYUSERDATAJSON

`func (o *SAPConnector) HasMODIFYUSERDATAJSON() bool`

HasMODIFYUSERDATAJSON returns a boolean if a field has been set.

### GetEXTERNAL_SOD_EVAL_JSON

`func (o *SAPConnector) GetEXTERNAL_SOD_EVAL_JSON() string`

GetEXTERNAL_SOD_EVAL_JSON returns the EXTERNAL_SOD_EVAL_JSON field if non-nil, zero value otherwise.

### GetEXTERNAL_SOD_EVAL_JSONOk

`func (o *SAPConnector) GetEXTERNAL_SOD_EVAL_JSONOk() (*string, bool)`

GetEXTERNAL_SOD_EVAL_JSONOk returns a tuple with the EXTERNAL_SOD_EVAL_JSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEXTERNAL_SOD_EVAL_JSON

`func (o *SAPConnector) SetEXTERNAL_SOD_EVAL_JSON(v string)`

SetEXTERNAL_SOD_EVAL_JSON sets EXTERNAL_SOD_EVAL_JSON field to given value.

### HasEXTERNAL_SOD_EVAL_JSON

`func (o *SAPConnector) HasEXTERNAL_SOD_EVAL_JSON() bool`

HasEXTERNAL_SOD_EVAL_JSON returns a boolean if a field has been set.

### GetEXTERNAL_SOD_EVAL_JSON_DETAIL

`func (o *SAPConnector) GetEXTERNAL_SOD_EVAL_JSON_DETAIL() string`

GetEXTERNAL_SOD_EVAL_JSON_DETAIL returns the EXTERNAL_SOD_EVAL_JSON_DETAIL field if non-nil, zero value otherwise.

### GetEXTERNAL_SOD_EVAL_JSON_DETAILOk

`func (o *SAPConnector) GetEXTERNAL_SOD_EVAL_JSON_DETAILOk() (*string, bool)`

GetEXTERNAL_SOD_EVAL_JSON_DETAILOk returns a tuple with the EXTERNAL_SOD_EVAL_JSON_DETAIL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEXTERNAL_SOD_EVAL_JSON_DETAIL

`func (o *SAPConnector) SetEXTERNAL_SOD_EVAL_JSON_DETAIL(v string)`

SetEXTERNAL_SOD_EVAL_JSON_DETAIL sets EXTERNAL_SOD_EVAL_JSON_DETAIL field to given value.

### HasEXTERNAL_SOD_EVAL_JSON_DETAIL

`func (o *SAPConnector) HasEXTERNAL_SOD_EVAL_JSON_DETAIL() bool`

HasEXTERNAL_SOD_EVAL_JSON_DETAIL returns a boolean if a field has been set.

### GetLOGS_TABLE_FILTER

`func (o *SAPConnector) GetLOGS_TABLE_FILTER() string`

GetLOGS_TABLE_FILTER returns the LOGS_TABLE_FILTER field if non-nil, zero value otherwise.

### GetLOGS_TABLE_FILTEROk

`func (o *SAPConnector) GetLOGS_TABLE_FILTEROk() (*string, bool)`

GetLOGS_TABLE_FILTEROk returns a tuple with the LOGS_TABLE_FILTER field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLOGS_TABLE_FILTER

`func (o *SAPConnector) SetLOGS_TABLE_FILTER(v string)`

SetLOGS_TABLE_FILTER sets LOGS_TABLE_FILTER field to given value.

### HasLOGS_TABLE_FILTER

`func (o *SAPConnector) HasLOGS_TABLE_FILTER() bool`

HasLOGS_TABLE_FILTER returns a boolean if a field has been set.

### GetPAM_CONFIG

`func (o *SAPConnector) GetPAM_CONFIG() string`

GetPAM_CONFIG returns the PAM_CONFIG field if non-nil, zero value otherwise.

### GetPAM_CONFIGOk

`func (o *SAPConnector) GetPAM_CONFIGOk() (*string, bool)`

GetPAM_CONFIGOk returns a tuple with the PAM_CONFIG field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPAM_CONFIG

`func (o *SAPConnector) SetPAM_CONFIG(v string)`

SetPAM_CONFIG sets PAM_CONFIG field to given value.

### HasPAM_CONFIG

`func (o *SAPConnector) HasPAM_CONFIG() bool`

HasPAM_CONFIG returns a boolean if a field has been set.

### GetSAPTABLE_FILTER_LANG

`func (o *SAPConnector) GetSAPTABLE_FILTER_LANG() string`

GetSAPTABLE_FILTER_LANG returns the SAPTABLE_FILTER_LANG field if non-nil, zero value otherwise.

### GetSAPTABLE_FILTER_LANGOk

`func (o *SAPConnector) GetSAPTABLE_FILTER_LANGOk() (*string, bool)`

GetSAPTABLE_FILTER_LANGOk returns a tuple with the SAPTABLE_FILTER_LANG field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSAPTABLE_FILTER_LANG

`func (o *SAPConnector) SetSAPTABLE_FILTER_LANG(v string)`

SetSAPTABLE_FILTER_LANG sets SAPTABLE_FILTER_LANG field to given value.

### HasSAPTABLE_FILTER_LANG

`func (o *SAPConnector) HasSAPTABLE_FILTER_LANG() bool`

HasSAPTABLE_FILTER_LANG returns a boolean if a field has been set.

### GetALTERNATE_OUTPUT_PARAMETER_ET_DATA

`func (o *SAPConnector) GetALTERNATE_OUTPUT_PARAMETER_ET_DATA() string`

GetALTERNATE_OUTPUT_PARAMETER_ET_DATA returns the ALTERNATE_OUTPUT_PARAMETER_ET_DATA field if non-nil, zero value otherwise.

### GetALTERNATE_OUTPUT_PARAMETER_ET_DATAOk

`func (o *SAPConnector) GetALTERNATE_OUTPUT_PARAMETER_ET_DATAOk() (*string, bool)`

GetALTERNATE_OUTPUT_PARAMETER_ET_DATAOk returns a tuple with the ALTERNATE_OUTPUT_PARAMETER_ET_DATA field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetALTERNATE_OUTPUT_PARAMETER_ET_DATA

`func (o *SAPConnector) SetALTERNATE_OUTPUT_PARAMETER_ET_DATA(v string)`

SetALTERNATE_OUTPUT_PARAMETER_ET_DATA sets ALTERNATE_OUTPUT_PARAMETER_ET_DATA field to given value.

### HasALTERNATE_OUTPUT_PARAMETER_ET_DATA

`func (o *SAPConnector) HasALTERNATE_OUTPUT_PARAMETER_ET_DATA() bool`

HasALTERNATE_OUTPUT_PARAMETER_ET_DATA returns a boolean if a field has been set.

### GetAUDIT_LOG_JSON

`func (o *SAPConnector) GetAUDIT_LOG_JSON() string`

GetAUDIT_LOG_JSON returns the AUDIT_LOG_JSON field if non-nil, zero value otherwise.

### GetAUDIT_LOG_JSONOk

`func (o *SAPConnector) GetAUDIT_LOG_JSONOk() (*string, bool)`

GetAUDIT_LOG_JSONOk returns a tuple with the AUDIT_LOG_JSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAUDIT_LOG_JSON

`func (o *SAPConnector) SetAUDIT_LOG_JSON(v string)`

SetAUDIT_LOG_JSON sets AUDIT_LOG_JSON field to given value.

### HasAUDIT_LOG_JSON

`func (o *SAPConnector) HasAUDIT_LOG_JSON() bool`

HasAUDIT_LOG_JSON returns a boolean if a field has been set.

### GetECCORS4HANA

`func (o *SAPConnector) GetECCORS4HANA() string`

GetECCORS4HANA returns the ECCORS4HANA field if non-nil, zero value otherwise.

### GetECCORS4HANAOk

`func (o *SAPConnector) GetECCORS4HANAOk() (*string, bool)`

GetECCORS4HANAOk returns a tuple with the ECCORS4HANA field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetECCORS4HANA

`func (o *SAPConnector) SetECCORS4HANA(v string)`

SetECCORS4HANA sets ECCORS4HANA field to given value.

### HasECCORS4HANA

`func (o *SAPConnector) HasECCORS4HANA() bool`

HasECCORS4HANA returns a boolean if a field has been set.

### GetDATA_IMPORT_FILTER

`func (o *SAPConnector) GetDATA_IMPORT_FILTER() string`

GetDATA_IMPORT_FILTER returns the DATA_IMPORT_FILTER field if non-nil, zero value otherwise.

### GetDATA_IMPORT_FILTEROk

`func (o *SAPConnector) GetDATA_IMPORT_FILTEROk() (*string, bool)`

GetDATA_IMPORT_FILTEROk returns a tuple with the DATA_IMPORT_FILTER field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDATA_IMPORT_FILTER

`func (o *SAPConnector) SetDATA_IMPORT_FILTER(v string)`

SetDATA_IMPORT_FILTER sets DATA_IMPORT_FILTER field to given value.

### HasDATA_IMPORT_FILTER

`func (o *SAPConnector) HasDATA_IMPORT_FILTER() bool`

HasDATA_IMPORT_FILTER returns a boolean if a field has been set.

### GetConfigJSON

`func (o *SAPConnector) GetConfigJSON() string`

GetConfigJSON returns the ConfigJSON field if non-nil, zero value otherwise.

### GetConfigJSONOk

`func (o *SAPConnector) GetConfigJSONOk() (*string, bool)`

GetConfigJSONOk returns a tuple with the ConfigJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigJSON

`func (o *SAPConnector) SetConfigJSON(v string)`

SetConfigJSON sets ConfigJSON field to given value.

### HasConfigJSON

`func (o *SAPConnector) HasConfigJSON() bool`

HasConfigJSON returns a boolean if a field has been set.

### GetROLE_DEFAULT_DATE

`func (o *SAPConnector) GetROLE_DEFAULT_DATE() string`

GetROLE_DEFAULT_DATE returns the ROLE_DEFAULT_DATE field if non-nil, zero value otherwise.

### GetROLE_DEFAULT_DATEOk

`func (o *SAPConnector) GetROLE_DEFAULT_DATEOk() (*string, bool)`

GetROLE_DEFAULT_DATEOk returns a tuple with the ROLE_DEFAULT_DATE field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetROLE_DEFAULT_DATE

`func (o *SAPConnector) SetROLE_DEFAULT_DATE(v string)`

SetROLE_DEFAULT_DATE sets ROLE_DEFAULT_DATE field to given value.

### HasROLE_DEFAULT_DATE

`func (o *SAPConnector) HasROLE_DEFAULT_DATE() bool`

HasROLE_DEFAULT_DATE returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


