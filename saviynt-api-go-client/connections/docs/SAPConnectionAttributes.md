# SAPConnectionAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CREATEACCOUNTJSON** | Pointer to **string** |  | [optional] 
**AUDIT_LOG_JSON** | Pointer to **string** |  | [optional] 
**ConnectionType** | Pointer to **string** |  | [optional] 
**SAPTABLE_FILTER_LANG** | Pointer to **string** |  | [optional] 
**PASSWORD_NOOFSPLCHARS** | Pointer to **string** |  | [optional] 
**TERMINATEDUSERGROUP** | Pointer to **string** |  | [optional] 
**LOGS_TABLE_FILTER** | Pointer to **string** |  | [optional] 
**ECCORS4HANA** | Pointer to **string** |  | [optional] 
**FIREFIGHTERID_REVOKE_ACCESS_JSON** | Pointer to **string** |  | [optional] 
**ConfigJSON** | Pointer to **string** |  | [optional] 
**FIREFIGHTERID_GRANT_ACCESS_JSON** | Pointer to **string** |  | [optional] 
**PROV_PASSWORD** | Pointer to **string** |  | [optional] 
**JCO_SNC_LIBRARY** | Pointer to **string** |  | [optional] 
**IsTimeoutSupported** | Pointer to **bool** |  | [optional] 
**JCOR3NAME** | Pointer to **string** |  | [optional] 
**EXTERNAL_SOD_EVAL_JSON** | Pointer to **string** |  | [optional] 
**JCO_ASHOST** | Pointer to **string** |  | [optional] 
**PASSWORD_NOOFDIGITS** | Pointer to **string** |  | [optional] 
**PROV_JCO_MSHOST** | Pointer to **string** |  | [optional] 
**PASSWORD** | Pointer to **string** |  | [optional] 
**PAM_CONFIG** | Pointer to **string** |  | [optional] 
**JCO_SNC_MYNAME** | Pointer to **string** |  | [optional] 
**ENFORCEPASSWORDCHANGE** | Pointer to **string** |  | [optional] 
**JCO_USER** | Pointer to **string** |  | [optional] 
**JCO_SNC_MODE** | Pointer to **string** |  | [optional] 
**PROV_JCO_MSSERV** | Pointer to **string** |  | [optional] 
**HANAREFTABLEJSON** | Pointer to **string** |  | [optional] 
**PASSWORD_MIN_LENGTH** | Pointer to **string** |  | [optional] 
**JCO_CLIENT** | Pointer to **string** |  | [optional] 
**TERMINATED_USER_ROLE_ACTION** | Pointer to **string** |  | [optional] 
**RESET_PWD_FOR_NEWACCOUNT** | Pointer to **string** |  | [optional] 
**PROV_JCO_CLIENT** | Pointer to **string** |  | [optional] 
**SNC** | Pointer to **string** |  | [optional] 
**JCO_MSSERV** | Pointer to **string** |  | [optional] 
**PROV_CUA_SNC** | Pointer to **string** |  | [optional] 
**ConnectionTimeoutConfig** | Pointer to [**ConnectionTimeoutConfig**](ConnectionTimeoutConfig.md) |  | [optional] 
**PROV_JCO_USER** | Pointer to **string** |  | [optional] 
**JCO_LANG** | Pointer to **string** |  | [optional] 
**JCO_SNC_PARTNERNAME** | Pointer to **string** |  | [optional] 
**STATUS_THRESHOLD_CONFIG** | Pointer to **string** |  | [optional] 
**PROV_JCO_SYSNR** | Pointer to **string** |  | [optional] 
**SETCUASYSTEM** | Pointer to **string** |  | [optional] 
**MESSAGESERVER** | Pointer to **string** |  | [optional] 
**PROV_JCO_ASHOST** | Pointer to **string** |  | [optional] 
**PROV_JCO_GROUP** | Pointer to **string** |  | [optional] 
**PROV_CUA_ENABLED** | Pointer to **string** |  | [optional] 
**JCO_MSHOST** | Pointer to **string** |  | [optional] 
**PROVJCOR3NAME** | Pointer to **string** |  | [optional] 
**PASSWORD_NOOFCAPSALPHA** | Pointer to **string** |  | [optional] 
**MODIFYUSERDATAJSON** | Pointer to **string** |  | [optional] 
**IsTimeoutConfigValidated** | Pointer to **bool** |  | [optional] 
**JCO_SNC_QOP** | Pointer to **string** |  | [optional] 
**TABLES** | Pointer to **string** |  | [optional] 
**PROV_JCO_LANG** | Pointer to **string** |  | [optional] 
**JCO_SYSNR** | Pointer to **string** |  | [optional] 
**EXTERNAL_SOD_EVAL_JSON_DETAIL** | Pointer to **string** |  | [optional] 
**DATA_IMPORT_FILTER** | Pointer to **string** |  | [optional] 
**ENABLEACCOUNTJSON** | Pointer to **string** |  | [optional] 
**ALTERNATE_OUTPUT_PARAMETER_ET_DATA** | Pointer to **string** |  | [optional] 
**JCO_GROUP** | Pointer to **string** |  | [optional] 
**PASSWORD_MAX_LENGTH** | Pointer to **string** |  | [optional] 
**USERIMPORTJSON** | Pointer to **string** |  | [optional] 
**SYSTEMNAME** | Pointer to **string** |  | [optional] 
**UPDATEACCOUNTJSON** | Pointer to **string** |  | [optional] 
**ROLE_DEFAULT_DATE** | Pointer to **string** | Default end date for SAP roles. Format: yyyyMMdd (e.g., 20251231). If not configured, defaults to 99991231. | [optional] 

## Methods

### NewSAPConnectionAttributes

`func NewSAPConnectionAttributes() *SAPConnectionAttributes`

NewSAPConnectionAttributes instantiates a new SAPConnectionAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSAPConnectionAttributesWithDefaults

`func NewSAPConnectionAttributesWithDefaults() *SAPConnectionAttributes`

NewSAPConnectionAttributesWithDefaults instantiates a new SAPConnectionAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCREATEACCOUNTJSON

`func (o *SAPConnectionAttributes) GetCREATEACCOUNTJSON() string`

GetCREATEACCOUNTJSON returns the CREATEACCOUNTJSON field if non-nil, zero value otherwise.

### GetCREATEACCOUNTJSONOk

`func (o *SAPConnectionAttributes) GetCREATEACCOUNTJSONOk() (*string, bool)`

GetCREATEACCOUNTJSONOk returns a tuple with the CREATEACCOUNTJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCREATEACCOUNTJSON

`func (o *SAPConnectionAttributes) SetCREATEACCOUNTJSON(v string)`

SetCREATEACCOUNTJSON sets CREATEACCOUNTJSON field to given value.

### HasCREATEACCOUNTJSON

`func (o *SAPConnectionAttributes) HasCREATEACCOUNTJSON() bool`

HasCREATEACCOUNTJSON returns a boolean if a field has been set.

### GetAUDIT_LOG_JSON

`func (o *SAPConnectionAttributes) GetAUDIT_LOG_JSON() string`

GetAUDIT_LOG_JSON returns the AUDIT_LOG_JSON field if non-nil, zero value otherwise.

### GetAUDIT_LOG_JSONOk

`func (o *SAPConnectionAttributes) GetAUDIT_LOG_JSONOk() (*string, bool)`

GetAUDIT_LOG_JSONOk returns a tuple with the AUDIT_LOG_JSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAUDIT_LOG_JSON

`func (o *SAPConnectionAttributes) SetAUDIT_LOG_JSON(v string)`

SetAUDIT_LOG_JSON sets AUDIT_LOG_JSON field to given value.

### HasAUDIT_LOG_JSON

`func (o *SAPConnectionAttributes) HasAUDIT_LOG_JSON() bool`

HasAUDIT_LOG_JSON returns a boolean if a field has been set.

### GetConnectionType

`func (o *SAPConnectionAttributes) GetConnectionType() string`

GetConnectionType returns the ConnectionType field if non-nil, zero value otherwise.

### GetConnectionTypeOk

`func (o *SAPConnectionAttributes) GetConnectionTypeOk() (*string, bool)`

GetConnectionTypeOk returns a tuple with the ConnectionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionType

`func (o *SAPConnectionAttributes) SetConnectionType(v string)`

SetConnectionType sets ConnectionType field to given value.

### HasConnectionType

`func (o *SAPConnectionAttributes) HasConnectionType() bool`

HasConnectionType returns a boolean if a field has been set.

### GetSAPTABLE_FILTER_LANG

`func (o *SAPConnectionAttributes) GetSAPTABLE_FILTER_LANG() string`

GetSAPTABLE_FILTER_LANG returns the SAPTABLE_FILTER_LANG field if non-nil, zero value otherwise.

### GetSAPTABLE_FILTER_LANGOk

`func (o *SAPConnectionAttributes) GetSAPTABLE_FILTER_LANGOk() (*string, bool)`

GetSAPTABLE_FILTER_LANGOk returns a tuple with the SAPTABLE_FILTER_LANG field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSAPTABLE_FILTER_LANG

`func (o *SAPConnectionAttributes) SetSAPTABLE_FILTER_LANG(v string)`

SetSAPTABLE_FILTER_LANG sets SAPTABLE_FILTER_LANG field to given value.

### HasSAPTABLE_FILTER_LANG

`func (o *SAPConnectionAttributes) HasSAPTABLE_FILTER_LANG() bool`

HasSAPTABLE_FILTER_LANG returns a boolean if a field has been set.

### GetPASSWORD_NOOFSPLCHARS

`func (o *SAPConnectionAttributes) GetPASSWORD_NOOFSPLCHARS() string`

GetPASSWORD_NOOFSPLCHARS returns the PASSWORD_NOOFSPLCHARS field if non-nil, zero value otherwise.

### GetPASSWORD_NOOFSPLCHARSOk

`func (o *SAPConnectionAttributes) GetPASSWORD_NOOFSPLCHARSOk() (*string, bool)`

GetPASSWORD_NOOFSPLCHARSOk returns a tuple with the PASSWORD_NOOFSPLCHARS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPASSWORD_NOOFSPLCHARS

`func (o *SAPConnectionAttributes) SetPASSWORD_NOOFSPLCHARS(v string)`

SetPASSWORD_NOOFSPLCHARS sets PASSWORD_NOOFSPLCHARS field to given value.

### HasPASSWORD_NOOFSPLCHARS

`func (o *SAPConnectionAttributes) HasPASSWORD_NOOFSPLCHARS() bool`

HasPASSWORD_NOOFSPLCHARS returns a boolean if a field has been set.

### GetTERMINATEDUSERGROUP

`func (o *SAPConnectionAttributes) GetTERMINATEDUSERGROUP() string`

GetTERMINATEDUSERGROUP returns the TERMINATEDUSERGROUP field if non-nil, zero value otherwise.

### GetTERMINATEDUSERGROUPOk

`func (o *SAPConnectionAttributes) GetTERMINATEDUSERGROUPOk() (*string, bool)`

GetTERMINATEDUSERGROUPOk returns a tuple with the TERMINATEDUSERGROUP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTERMINATEDUSERGROUP

`func (o *SAPConnectionAttributes) SetTERMINATEDUSERGROUP(v string)`

SetTERMINATEDUSERGROUP sets TERMINATEDUSERGROUP field to given value.

### HasTERMINATEDUSERGROUP

`func (o *SAPConnectionAttributes) HasTERMINATEDUSERGROUP() bool`

HasTERMINATEDUSERGROUP returns a boolean if a field has been set.

### GetLOGS_TABLE_FILTER

`func (o *SAPConnectionAttributes) GetLOGS_TABLE_FILTER() string`

GetLOGS_TABLE_FILTER returns the LOGS_TABLE_FILTER field if non-nil, zero value otherwise.

### GetLOGS_TABLE_FILTEROk

`func (o *SAPConnectionAttributes) GetLOGS_TABLE_FILTEROk() (*string, bool)`

GetLOGS_TABLE_FILTEROk returns a tuple with the LOGS_TABLE_FILTER field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLOGS_TABLE_FILTER

`func (o *SAPConnectionAttributes) SetLOGS_TABLE_FILTER(v string)`

SetLOGS_TABLE_FILTER sets LOGS_TABLE_FILTER field to given value.

### HasLOGS_TABLE_FILTER

`func (o *SAPConnectionAttributes) HasLOGS_TABLE_FILTER() bool`

HasLOGS_TABLE_FILTER returns a boolean if a field has been set.

### GetECCORS4HANA

`func (o *SAPConnectionAttributes) GetECCORS4HANA() string`

GetECCORS4HANA returns the ECCORS4HANA field if non-nil, zero value otherwise.

### GetECCORS4HANAOk

`func (o *SAPConnectionAttributes) GetECCORS4HANAOk() (*string, bool)`

GetECCORS4HANAOk returns a tuple with the ECCORS4HANA field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetECCORS4HANA

`func (o *SAPConnectionAttributes) SetECCORS4HANA(v string)`

SetECCORS4HANA sets ECCORS4HANA field to given value.

### HasECCORS4HANA

`func (o *SAPConnectionAttributes) HasECCORS4HANA() bool`

HasECCORS4HANA returns a boolean if a field has been set.

### GetFIREFIGHTERID_REVOKE_ACCESS_JSON

`func (o *SAPConnectionAttributes) GetFIREFIGHTERID_REVOKE_ACCESS_JSON() string`

GetFIREFIGHTERID_REVOKE_ACCESS_JSON returns the FIREFIGHTERID_REVOKE_ACCESS_JSON field if non-nil, zero value otherwise.

### GetFIREFIGHTERID_REVOKE_ACCESS_JSONOk

`func (o *SAPConnectionAttributes) GetFIREFIGHTERID_REVOKE_ACCESS_JSONOk() (*string, bool)`

GetFIREFIGHTERID_REVOKE_ACCESS_JSONOk returns a tuple with the FIREFIGHTERID_REVOKE_ACCESS_JSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFIREFIGHTERID_REVOKE_ACCESS_JSON

`func (o *SAPConnectionAttributes) SetFIREFIGHTERID_REVOKE_ACCESS_JSON(v string)`

SetFIREFIGHTERID_REVOKE_ACCESS_JSON sets FIREFIGHTERID_REVOKE_ACCESS_JSON field to given value.

### HasFIREFIGHTERID_REVOKE_ACCESS_JSON

`func (o *SAPConnectionAttributes) HasFIREFIGHTERID_REVOKE_ACCESS_JSON() bool`

HasFIREFIGHTERID_REVOKE_ACCESS_JSON returns a boolean if a field has been set.

### GetConfigJSON

`func (o *SAPConnectionAttributes) GetConfigJSON() string`

GetConfigJSON returns the ConfigJSON field if non-nil, zero value otherwise.

### GetConfigJSONOk

`func (o *SAPConnectionAttributes) GetConfigJSONOk() (*string, bool)`

GetConfigJSONOk returns a tuple with the ConfigJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigJSON

`func (o *SAPConnectionAttributes) SetConfigJSON(v string)`

SetConfigJSON sets ConfigJSON field to given value.

### HasConfigJSON

`func (o *SAPConnectionAttributes) HasConfigJSON() bool`

HasConfigJSON returns a boolean if a field has been set.

### GetFIREFIGHTERID_GRANT_ACCESS_JSON

`func (o *SAPConnectionAttributes) GetFIREFIGHTERID_GRANT_ACCESS_JSON() string`

GetFIREFIGHTERID_GRANT_ACCESS_JSON returns the FIREFIGHTERID_GRANT_ACCESS_JSON field if non-nil, zero value otherwise.

### GetFIREFIGHTERID_GRANT_ACCESS_JSONOk

`func (o *SAPConnectionAttributes) GetFIREFIGHTERID_GRANT_ACCESS_JSONOk() (*string, bool)`

GetFIREFIGHTERID_GRANT_ACCESS_JSONOk returns a tuple with the FIREFIGHTERID_GRANT_ACCESS_JSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFIREFIGHTERID_GRANT_ACCESS_JSON

`func (o *SAPConnectionAttributes) SetFIREFIGHTERID_GRANT_ACCESS_JSON(v string)`

SetFIREFIGHTERID_GRANT_ACCESS_JSON sets FIREFIGHTERID_GRANT_ACCESS_JSON field to given value.

### HasFIREFIGHTERID_GRANT_ACCESS_JSON

`func (o *SAPConnectionAttributes) HasFIREFIGHTERID_GRANT_ACCESS_JSON() bool`

HasFIREFIGHTERID_GRANT_ACCESS_JSON returns a boolean if a field has been set.

### GetPROV_PASSWORD

`func (o *SAPConnectionAttributes) GetPROV_PASSWORD() string`

GetPROV_PASSWORD returns the PROV_PASSWORD field if non-nil, zero value otherwise.

### GetPROV_PASSWORDOk

`func (o *SAPConnectionAttributes) GetPROV_PASSWORDOk() (*string, bool)`

GetPROV_PASSWORDOk returns a tuple with the PROV_PASSWORD field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPROV_PASSWORD

`func (o *SAPConnectionAttributes) SetPROV_PASSWORD(v string)`

SetPROV_PASSWORD sets PROV_PASSWORD field to given value.

### HasPROV_PASSWORD

`func (o *SAPConnectionAttributes) HasPROV_PASSWORD() bool`

HasPROV_PASSWORD returns a boolean if a field has been set.

### GetJCO_SNC_LIBRARY

`func (o *SAPConnectionAttributes) GetJCO_SNC_LIBRARY() string`

GetJCO_SNC_LIBRARY returns the JCO_SNC_LIBRARY field if non-nil, zero value otherwise.

### GetJCO_SNC_LIBRARYOk

`func (o *SAPConnectionAttributes) GetJCO_SNC_LIBRARYOk() (*string, bool)`

GetJCO_SNC_LIBRARYOk returns a tuple with the JCO_SNC_LIBRARY field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJCO_SNC_LIBRARY

`func (o *SAPConnectionAttributes) SetJCO_SNC_LIBRARY(v string)`

SetJCO_SNC_LIBRARY sets JCO_SNC_LIBRARY field to given value.

### HasJCO_SNC_LIBRARY

`func (o *SAPConnectionAttributes) HasJCO_SNC_LIBRARY() bool`

HasJCO_SNC_LIBRARY returns a boolean if a field has been set.

### GetIsTimeoutSupported

`func (o *SAPConnectionAttributes) GetIsTimeoutSupported() bool`

GetIsTimeoutSupported returns the IsTimeoutSupported field if non-nil, zero value otherwise.

### GetIsTimeoutSupportedOk

`func (o *SAPConnectionAttributes) GetIsTimeoutSupportedOk() (*bool, bool)`

GetIsTimeoutSupportedOk returns a tuple with the IsTimeoutSupported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsTimeoutSupported

`func (o *SAPConnectionAttributes) SetIsTimeoutSupported(v bool)`

SetIsTimeoutSupported sets IsTimeoutSupported field to given value.

### HasIsTimeoutSupported

`func (o *SAPConnectionAttributes) HasIsTimeoutSupported() bool`

HasIsTimeoutSupported returns a boolean if a field has been set.

### GetJCOR3NAME

`func (o *SAPConnectionAttributes) GetJCOR3NAME() string`

GetJCOR3NAME returns the JCOR3NAME field if non-nil, zero value otherwise.

### GetJCOR3NAMEOk

`func (o *SAPConnectionAttributes) GetJCOR3NAMEOk() (*string, bool)`

GetJCOR3NAMEOk returns a tuple with the JCOR3NAME field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJCOR3NAME

`func (o *SAPConnectionAttributes) SetJCOR3NAME(v string)`

SetJCOR3NAME sets JCOR3NAME field to given value.

### HasJCOR3NAME

`func (o *SAPConnectionAttributes) HasJCOR3NAME() bool`

HasJCOR3NAME returns a boolean if a field has been set.

### GetEXTERNAL_SOD_EVAL_JSON

`func (o *SAPConnectionAttributes) GetEXTERNAL_SOD_EVAL_JSON() string`

GetEXTERNAL_SOD_EVAL_JSON returns the EXTERNAL_SOD_EVAL_JSON field if non-nil, zero value otherwise.

### GetEXTERNAL_SOD_EVAL_JSONOk

`func (o *SAPConnectionAttributes) GetEXTERNAL_SOD_EVAL_JSONOk() (*string, bool)`

GetEXTERNAL_SOD_EVAL_JSONOk returns a tuple with the EXTERNAL_SOD_EVAL_JSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEXTERNAL_SOD_EVAL_JSON

`func (o *SAPConnectionAttributes) SetEXTERNAL_SOD_EVAL_JSON(v string)`

SetEXTERNAL_SOD_EVAL_JSON sets EXTERNAL_SOD_EVAL_JSON field to given value.

### HasEXTERNAL_SOD_EVAL_JSON

`func (o *SAPConnectionAttributes) HasEXTERNAL_SOD_EVAL_JSON() bool`

HasEXTERNAL_SOD_EVAL_JSON returns a boolean if a field has been set.

### GetJCO_ASHOST

`func (o *SAPConnectionAttributes) GetJCO_ASHOST() string`

GetJCO_ASHOST returns the JCO_ASHOST field if non-nil, zero value otherwise.

### GetJCO_ASHOSTOk

`func (o *SAPConnectionAttributes) GetJCO_ASHOSTOk() (*string, bool)`

GetJCO_ASHOSTOk returns a tuple with the JCO_ASHOST field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJCO_ASHOST

`func (o *SAPConnectionAttributes) SetJCO_ASHOST(v string)`

SetJCO_ASHOST sets JCO_ASHOST field to given value.

### HasJCO_ASHOST

`func (o *SAPConnectionAttributes) HasJCO_ASHOST() bool`

HasJCO_ASHOST returns a boolean if a field has been set.

### GetPASSWORD_NOOFDIGITS

`func (o *SAPConnectionAttributes) GetPASSWORD_NOOFDIGITS() string`

GetPASSWORD_NOOFDIGITS returns the PASSWORD_NOOFDIGITS field if non-nil, zero value otherwise.

### GetPASSWORD_NOOFDIGITSOk

`func (o *SAPConnectionAttributes) GetPASSWORD_NOOFDIGITSOk() (*string, bool)`

GetPASSWORD_NOOFDIGITSOk returns a tuple with the PASSWORD_NOOFDIGITS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPASSWORD_NOOFDIGITS

`func (o *SAPConnectionAttributes) SetPASSWORD_NOOFDIGITS(v string)`

SetPASSWORD_NOOFDIGITS sets PASSWORD_NOOFDIGITS field to given value.

### HasPASSWORD_NOOFDIGITS

`func (o *SAPConnectionAttributes) HasPASSWORD_NOOFDIGITS() bool`

HasPASSWORD_NOOFDIGITS returns a boolean if a field has been set.

### GetPROV_JCO_MSHOST

`func (o *SAPConnectionAttributes) GetPROV_JCO_MSHOST() string`

GetPROV_JCO_MSHOST returns the PROV_JCO_MSHOST field if non-nil, zero value otherwise.

### GetPROV_JCO_MSHOSTOk

`func (o *SAPConnectionAttributes) GetPROV_JCO_MSHOSTOk() (*string, bool)`

GetPROV_JCO_MSHOSTOk returns a tuple with the PROV_JCO_MSHOST field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPROV_JCO_MSHOST

`func (o *SAPConnectionAttributes) SetPROV_JCO_MSHOST(v string)`

SetPROV_JCO_MSHOST sets PROV_JCO_MSHOST field to given value.

### HasPROV_JCO_MSHOST

`func (o *SAPConnectionAttributes) HasPROV_JCO_MSHOST() bool`

HasPROV_JCO_MSHOST returns a boolean if a field has been set.

### GetPASSWORD

`func (o *SAPConnectionAttributes) GetPASSWORD() string`

GetPASSWORD returns the PASSWORD field if non-nil, zero value otherwise.

### GetPASSWORDOk

`func (o *SAPConnectionAttributes) GetPASSWORDOk() (*string, bool)`

GetPASSWORDOk returns a tuple with the PASSWORD field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPASSWORD

`func (o *SAPConnectionAttributes) SetPASSWORD(v string)`

SetPASSWORD sets PASSWORD field to given value.

### HasPASSWORD

`func (o *SAPConnectionAttributes) HasPASSWORD() bool`

HasPASSWORD returns a boolean if a field has been set.

### GetPAM_CONFIG

`func (o *SAPConnectionAttributes) GetPAM_CONFIG() string`

GetPAM_CONFIG returns the PAM_CONFIG field if non-nil, zero value otherwise.

### GetPAM_CONFIGOk

`func (o *SAPConnectionAttributes) GetPAM_CONFIGOk() (*string, bool)`

GetPAM_CONFIGOk returns a tuple with the PAM_CONFIG field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPAM_CONFIG

`func (o *SAPConnectionAttributes) SetPAM_CONFIG(v string)`

SetPAM_CONFIG sets PAM_CONFIG field to given value.

### HasPAM_CONFIG

`func (o *SAPConnectionAttributes) HasPAM_CONFIG() bool`

HasPAM_CONFIG returns a boolean if a field has been set.

### GetJCO_SNC_MYNAME

`func (o *SAPConnectionAttributes) GetJCO_SNC_MYNAME() string`

GetJCO_SNC_MYNAME returns the JCO_SNC_MYNAME field if non-nil, zero value otherwise.

### GetJCO_SNC_MYNAMEOk

`func (o *SAPConnectionAttributes) GetJCO_SNC_MYNAMEOk() (*string, bool)`

GetJCO_SNC_MYNAMEOk returns a tuple with the JCO_SNC_MYNAME field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJCO_SNC_MYNAME

`func (o *SAPConnectionAttributes) SetJCO_SNC_MYNAME(v string)`

SetJCO_SNC_MYNAME sets JCO_SNC_MYNAME field to given value.

### HasJCO_SNC_MYNAME

`func (o *SAPConnectionAttributes) HasJCO_SNC_MYNAME() bool`

HasJCO_SNC_MYNAME returns a boolean if a field has been set.

### GetENFORCEPASSWORDCHANGE

`func (o *SAPConnectionAttributes) GetENFORCEPASSWORDCHANGE() string`

GetENFORCEPASSWORDCHANGE returns the ENFORCEPASSWORDCHANGE field if non-nil, zero value otherwise.

### GetENFORCEPASSWORDCHANGEOk

`func (o *SAPConnectionAttributes) GetENFORCEPASSWORDCHANGEOk() (*string, bool)`

GetENFORCEPASSWORDCHANGEOk returns a tuple with the ENFORCEPASSWORDCHANGE field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetENFORCEPASSWORDCHANGE

`func (o *SAPConnectionAttributes) SetENFORCEPASSWORDCHANGE(v string)`

SetENFORCEPASSWORDCHANGE sets ENFORCEPASSWORDCHANGE field to given value.

### HasENFORCEPASSWORDCHANGE

`func (o *SAPConnectionAttributes) HasENFORCEPASSWORDCHANGE() bool`

HasENFORCEPASSWORDCHANGE returns a boolean if a field has been set.

### GetJCO_USER

`func (o *SAPConnectionAttributes) GetJCO_USER() string`

GetJCO_USER returns the JCO_USER field if non-nil, zero value otherwise.

### GetJCO_USEROk

`func (o *SAPConnectionAttributes) GetJCO_USEROk() (*string, bool)`

GetJCO_USEROk returns a tuple with the JCO_USER field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJCO_USER

`func (o *SAPConnectionAttributes) SetJCO_USER(v string)`

SetJCO_USER sets JCO_USER field to given value.

### HasJCO_USER

`func (o *SAPConnectionAttributes) HasJCO_USER() bool`

HasJCO_USER returns a boolean if a field has been set.

### GetJCO_SNC_MODE

`func (o *SAPConnectionAttributes) GetJCO_SNC_MODE() string`

GetJCO_SNC_MODE returns the JCO_SNC_MODE field if non-nil, zero value otherwise.

### GetJCO_SNC_MODEOk

`func (o *SAPConnectionAttributes) GetJCO_SNC_MODEOk() (*string, bool)`

GetJCO_SNC_MODEOk returns a tuple with the JCO_SNC_MODE field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJCO_SNC_MODE

`func (o *SAPConnectionAttributes) SetJCO_SNC_MODE(v string)`

SetJCO_SNC_MODE sets JCO_SNC_MODE field to given value.

### HasJCO_SNC_MODE

`func (o *SAPConnectionAttributes) HasJCO_SNC_MODE() bool`

HasJCO_SNC_MODE returns a boolean if a field has been set.

### GetPROV_JCO_MSSERV

`func (o *SAPConnectionAttributes) GetPROV_JCO_MSSERV() string`

GetPROV_JCO_MSSERV returns the PROV_JCO_MSSERV field if non-nil, zero value otherwise.

### GetPROV_JCO_MSSERVOk

`func (o *SAPConnectionAttributes) GetPROV_JCO_MSSERVOk() (*string, bool)`

GetPROV_JCO_MSSERVOk returns a tuple with the PROV_JCO_MSSERV field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPROV_JCO_MSSERV

`func (o *SAPConnectionAttributes) SetPROV_JCO_MSSERV(v string)`

SetPROV_JCO_MSSERV sets PROV_JCO_MSSERV field to given value.

### HasPROV_JCO_MSSERV

`func (o *SAPConnectionAttributes) HasPROV_JCO_MSSERV() bool`

HasPROV_JCO_MSSERV returns a boolean if a field has been set.

### GetHANAREFTABLEJSON

`func (o *SAPConnectionAttributes) GetHANAREFTABLEJSON() string`

GetHANAREFTABLEJSON returns the HANAREFTABLEJSON field if non-nil, zero value otherwise.

### GetHANAREFTABLEJSONOk

`func (o *SAPConnectionAttributes) GetHANAREFTABLEJSONOk() (*string, bool)`

GetHANAREFTABLEJSONOk returns a tuple with the HANAREFTABLEJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHANAREFTABLEJSON

`func (o *SAPConnectionAttributes) SetHANAREFTABLEJSON(v string)`

SetHANAREFTABLEJSON sets HANAREFTABLEJSON field to given value.

### HasHANAREFTABLEJSON

`func (o *SAPConnectionAttributes) HasHANAREFTABLEJSON() bool`

HasHANAREFTABLEJSON returns a boolean if a field has been set.

### GetPASSWORD_MIN_LENGTH

`func (o *SAPConnectionAttributes) GetPASSWORD_MIN_LENGTH() string`

GetPASSWORD_MIN_LENGTH returns the PASSWORD_MIN_LENGTH field if non-nil, zero value otherwise.

### GetPASSWORD_MIN_LENGTHOk

`func (o *SAPConnectionAttributes) GetPASSWORD_MIN_LENGTHOk() (*string, bool)`

GetPASSWORD_MIN_LENGTHOk returns a tuple with the PASSWORD_MIN_LENGTH field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPASSWORD_MIN_LENGTH

`func (o *SAPConnectionAttributes) SetPASSWORD_MIN_LENGTH(v string)`

SetPASSWORD_MIN_LENGTH sets PASSWORD_MIN_LENGTH field to given value.

### HasPASSWORD_MIN_LENGTH

`func (o *SAPConnectionAttributes) HasPASSWORD_MIN_LENGTH() bool`

HasPASSWORD_MIN_LENGTH returns a boolean if a field has been set.

### GetJCO_CLIENT

`func (o *SAPConnectionAttributes) GetJCO_CLIENT() string`

GetJCO_CLIENT returns the JCO_CLIENT field if non-nil, zero value otherwise.

### GetJCO_CLIENTOk

`func (o *SAPConnectionAttributes) GetJCO_CLIENTOk() (*string, bool)`

GetJCO_CLIENTOk returns a tuple with the JCO_CLIENT field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJCO_CLIENT

`func (o *SAPConnectionAttributes) SetJCO_CLIENT(v string)`

SetJCO_CLIENT sets JCO_CLIENT field to given value.

### HasJCO_CLIENT

`func (o *SAPConnectionAttributes) HasJCO_CLIENT() bool`

HasJCO_CLIENT returns a boolean if a field has been set.

### GetTERMINATED_USER_ROLE_ACTION

`func (o *SAPConnectionAttributes) GetTERMINATED_USER_ROLE_ACTION() string`

GetTERMINATED_USER_ROLE_ACTION returns the TERMINATED_USER_ROLE_ACTION field if non-nil, zero value otherwise.

### GetTERMINATED_USER_ROLE_ACTIONOk

`func (o *SAPConnectionAttributes) GetTERMINATED_USER_ROLE_ACTIONOk() (*string, bool)`

GetTERMINATED_USER_ROLE_ACTIONOk returns a tuple with the TERMINATED_USER_ROLE_ACTION field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTERMINATED_USER_ROLE_ACTION

`func (o *SAPConnectionAttributes) SetTERMINATED_USER_ROLE_ACTION(v string)`

SetTERMINATED_USER_ROLE_ACTION sets TERMINATED_USER_ROLE_ACTION field to given value.

### HasTERMINATED_USER_ROLE_ACTION

`func (o *SAPConnectionAttributes) HasTERMINATED_USER_ROLE_ACTION() bool`

HasTERMINATED_USER_ROLE_ACTION returns a boolean if a field has been set.

### GetRESET_PWD_FOR_NEWACCOUNT

`func (o *SAPConnectionAttributes) GetRESET_PWD_FOR_NEWACCOUNT() string`

GetRESET_PWD_FOR_NEWACCOUNT returns the RESET_PWD_FOR_NEWACCOUNT field if non-nil, zero value otherwise.

### GetRESET_PWD_FOR_NEWACCOUNTOk

`func (o *SAPConnectionAttributes) GetRESET_PWD_FOR_NEWACCOUNTOk() (*string, bool)`

GetRESET_PWD_FOR_NEWACCOUNTOk returns a tuple with the RESET_PWD_FOR_NEWACCOUNT field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRESET_PWD_FOR_NEWACCOUNT

`func (o *SAPConnectionAttributes) SetRESET_PWD_FOR_NEWACCOUNT(v string)`

SetRESET_PWD_FOR_NEWACCOUNT sets RESET_PWD_FOR_NEWACCOUNT field to given value.

### HasRESET_PWD_FOR_NEWACCOUNT

`func (o *SAPConnectionAttributes) HasRESET_PWD_FOR_NEWACCOUNT() bool`

HasRESET_PWD_FOR_NEWACCOUNT returns a boolean if a field has been set.

### GetPROV_JCO_CLIENT

`func (o *SAPConnectionAttributes) GetPROV_JCO_CLIENT() string`

GetPROV_JCO_CLIENT returns the PROV_JCO_CLIENT field if non-nil, zero value otherwise.

### GetPROV_JCO_CLIENTOk

`func (o *SAPConnectionAttributes) GetPROV_JCO_CLIENTOk() (*string, bool)`

GetPROV_JCO_CLIENTOk returns a tuple with the PROV_JCO_CLIENT field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPROV_JCO_CLIENT

`func (o *SAPConnectionAttributes) SetPROV_JCO_CLIENT(v string)`

SetPROV_JCO_CLIENT sets PROV_JCO_CLIENT field to given value.

### HasPROV_JCO_CLIENT

`func (o *SAPConnectionAttributes) HasPROV_JCO_CLIENT() bool`

HasPROV_JCO_CLIENT returns a boolean if a field has been set.

### GetSNC

`func (o *SAPConnectionAttributes) GetSNC() string`

GetSNC returns the SNC field if non-nil, zero value otherwise.

### GetSNCOk

`func (o *SAPConnectionAttributes) GetSNCOk() (*string, bool)`

GetSNCOk returns a tuple with the SNC field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSNC

`func (o *SAPConnectionAttributes) SetSNC(v string)`

SetSNC sets SNC field to given value.

### HasSNC

`func (o *SAPConnectionAttributes) HasSNC() bool`

HasSNC returns a boolean if a field has been set.

### GetJCO_MSSERV

`func (o *SAPConnectionAttributes) GetJCO_MSSERV() string`

GetJCO_MSSERV returns the JCO_MSSERV field if non-nil, zero value otherwise.

### GetJCO_MSSERVOk

`func (o *SAPConnectionAttributes) GetJCO_MSSERVOk() (*string, bool)`

GetJCO_MSSERVOk returns a tuple with the JCO_MSSERV field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJCO_MSSERV

`func (o *SAPConnectionAttributes) SetJCO_MSSERV(v string)`

SetJCO_MSSERV sets JCO_MSSERV field to given value.

### HasJCO_MSSERV

`func (o *SAPConnectionAttributes) HasJCO_MSSERV() bool`

HasJCO_MSSERV returns a boolean if a field has been set.

### GetPROV_CUA_SNC

`func (o *SAPConnectionAttributes) GetPROV_CUA_SNC() string`

GetPROV_CUA_SNC returns the PROV_CUA_SNC field if non-nil, zero value otherwise.

### GetPROV_CUA_SNCOk

`func (o *SAPConnectionAttributes) GetPROV_CUA_SNCOk() (*string, bool)`

GetPROV_CUA_SNCOk returns a tuple with the PROV_CUA_SNC field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPROV_CUA_SNC

`func (o *SAPConnectionAttributes) SetPROV_CUA_SNC(v string)`

SetPROV_CUA_SNC sets PROV_CUA_SNC field to given value.

### HasPROV_CUA_SNC

`func (o *SAPConnectionAttributes) HasPROV_CUA_SNC() bool`

HasPROV_CUA_SNC returns a boolean if a field has been set.

### GetConnectionTimeoutConfig

`func (o *SAPConnectionAttributes) GetConnectionTimeoutConfig() ConnectionTimeoutConfig`

GetConnectionTimeoutConfig returns the ConnectionTimeoutConfig field if non-nil, zero value otherwise.

### GetConnectionTimeoutConfigOk

`func (o *SAPConnectionAttributes) GetConnectionTimeoutConfigOk() (*ConnectionTimeoutConfig, bool)`

GetConnectionTimeoutConfigOk returns a tuple with the ConnectionTimeoutConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionTimeoutConfig

`func (o *SAPConnectionAttributes) SetConnectionTimeoutConfig(v ConnectionTimeoutConfig)`

SetConnectionTimeoutConfig sets ConnectionTimeoutConfig field to given value.

### HasConnectionTimeoutConfig

`func (o *SAPConnectionAttributes) HasConnectionTimeoutConfig() bool`

HasConnectionTimeoutConfig returns a boolean if a field has been set.

### GetPROV_JCO_USER

`func (o *SAPConnectionAttributes) GetPROV_JCO_USER() string`

GetPROV_JCO_USER returns the PROV_JCO_USER field if non-nil, zero value otherwise.

### GetPROV_JCO_USEROk

`func (o *SAPConnectionAttributes) GetPROV_JCO_USEROk() (*string, bool)`

GetPROV_JCO_USEROk returns a tuple with the PROV_JCO_USER field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPROV_JCO_USER

`func (o *SAPConnectionAttributes) SetPROV_JCO_USER(v string)`

SetPROV_JCO_USER sets PROV_JCO_USER field to given value.

### HasPROV_JCO_USER

`func (o *SAPConnectionAttributes) HasPROV_JCO_USER() bool`

HasPROV_JCO_USER returns a boolean if a field has been set.

### GetJCO_LANG

`func (o *SAPConnectionAttributes) GetJCO_LANG() string`

GetJCO_LANG returns the JCO_LANG field if non-nil, zero value otherwise.

### GetJCO_LANGOk

`func (o *SAPConnectionAttributes) GetJCO_LANGOk() (*string, bool)`

GetJCO_LANGOk returns a tuple with the JCO_LANG field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJCO_LANG

`func (o *SAPConnectionAttributes) SetJCO_LANG(v string)`

SetJCO_LANG sets JCO_LANG field to given value.

### HasJCO_LANG

`func (o *SAPConnectionAttributes) HasJCO_LANG() bool`

HasJCO_LANG returns a boolean if a field has been set.

### GetJCO_SNC_PARTNERNAME

`func (o *SAPConnectionAttributes) GetJCO_SNC_PARTNERNAME() string`

GetJCO_SNC_PARTNERNAME returns the JCO_SNC_PARTNERNAME field if non-nil, zero value otherwise.

### GetJCO_SNC_PARTNERNAMEOk

`func (o *SAPConnectionAttributes) GetJCO_SNC_PARTNERNAMEOk() (*string, bool)`

GetJCO_SNC_PARTNERNAMEOk returns a tuple with the JCO_SNC_PARTNERNAME field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJCO_SNC_PARTNERNAME

`func (o *SAPConnectionAttributes) SetJCO_SNC_PARTNERNAME(v string)`

SetJCO_SNC_PARTNERNAME sets JCO_SNC_PARTNERNAME field to given value.

### HasJCO_SNC_PARTNERNAME

`func (o *SAPConnectionAttributes) HasJCO_SNC_PARTNERNAME() bool`

HasJCO_SNC_PARTNERNAME returns a boolean if a field has been set.

### GetSTATUS_THRESHOLD_CONFIG

`func (o *SAPConnectionAttributes) GetSTATUS_THRESHOLD_CONFIG() string`

GetSTATUS_THRESHOLD_CONFIG returns the STATUS_THRESHOLD_CONFIG field if non-nil, zero value otherwise.

### GetSTATUS_THRESHOLD_CONFIGOk

`func (o *SAPConnectionAttributes) GetSTATUS_THRESHOLD_CONFIGOk() (*string, bool)`

GetSTATUS_THRESHOLD_CONFIGOk returns a tuple with the STATUS_THRESHOLD_CONFIG field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSTATUS_THRESHOLD_CONFIG

`func (o *SAPConnectionAttributes) SetSTATUS_THRESHOLD_CONFIG(v string)`

SetSTATUS_THRESHOLD_CONFIG sets STATUS_THRESHOLD_CONFIG field to given value.

### HasSTATUS_THRESHOLD_CONFIG

`func (o *SAPConnectionAttributes) HasSTATUS_THRESHOLD_CONFIG() bool`

HasSTATUS_THRESHOLD_CONFIG returns a boolean if a field has been set.

### GetPROV_JCO_SYSNR

`func (o *SAPConnectionAttributes) GetPROV_JCO_SYSNR() string`

GetPROV_JCO_SYSNR returns the PROV_JCO_SYSNR field if non-nil, zero value otherwise.

### GetPROV_JCO_SYSNROk

`func (o *SAPConnectionAttributes) GetPROV_JCO_SYSNROk() (*string, bool)`

GetPROV_JCO_SYSNROk returns a tuple with the PROV_JCO_SYSNR field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPROV_JCO_SYSNR

`func (o *SAPConnectionAttributes) SetPROV_JCO_SYSNR(v string)`

SetPROV_JCO_SYSNR sets PROV_JCO_SYSNR field to given value.

### HasPROV_JCO_SYSNR

`func (o *SAPConnectionAttributes) HasPROV_JCO_SYSNR() bool`

HasPROV_JCO_SYSNR returns a boolean if a field has been set.

### GetSETCUASYSTEM

`func (o *SAPConnectionAttributes) GetSETCUASYSTEM() string`

GetSETCUASYSTEM returns the SETCUASYSTEM field if non-nil, zero value otherwise.

### GetSETCUASYSTEMOk

`func (o *SAPConnectionAttributes) GetSETCUASYSTEMOk() (*string, bool)`

GetSETCUASYSTEMOk returns a tuple with the SETCUASYSTEM field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSETCUASYSTEM

`func (o *SAPConnectionAttributes) SetSETCUASYSTEM(v string)`

SetSETCUASYSTEM sets SETCUASYSTEM field to given value.

### HasSETCUASYSTEM

`func (o *SAPConnectionAttributes) HasSETCUASYSTEM() bool`

HasSETCUASYSTEM returns a boolean if a field has been set.

### GetMESSAGESERVER

`func (o *SAPConnectionAttributes) GetMESSAGESERVER() string`

GetMESSAGESERVER returns the MESSAGESERVER field if non-nil, zero value otherwise.

### GetMESSAGESERVEROk

`func (o *SAPConnectionAttributes) GetMESSAGESERVEROk() (*string, bool)`

GetMESSAGESERVEROk returns a tuple with the MESSAGESERVER field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMESSAGESERVER

`func (o *SAPConnectionAttributes) SetMESSAGESERVER(v string)`

SetMESSAGESERVER sets MESSAGESERVER field to given value.

### HasMESSAGESERVER

`func (o *SAPConnectionAttributes) HasMESSAGESERVER() bool`

HasMESSAGESERVER returns a boolean if a field has been set.

### GetPROV_JCO_ASHOST

`func (o *SAPConnectionAttributes) GetPROV_JCO_ASHOST() string`

GetPROV_JCO_ASHOST returns the PROV_JCO_ASHOST field if non-nil, zero value otherwise.

### GetPROV_JCO_ASHOSTOk

`func (o *SAPConnectionAttributes) GetPROV_JCO_ASHOSTOk() (*string, bool)`

GetPROV_JCO_ASHOSTOk returns a tuple with the PROV_JCO_ASHOST field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPROV_JCO_ASHOST

`func (o *SAPConnectionAttributes) SetPROV_JCO_ASHOST(v string)`

SetPROV_JCO_ASHOST sets PROV_JCO_ASHOST field to given value.

### HasPROV_JCO_ASHOST

`func (o *SAPConnectionAttributes) HasPROV_JCO_ASHOST() bool`

HasPROV_JCO_ASHOST returns a boolean if a field has been set.

### GetPROV_JCO_GROUP

`func (o *SAPConnectionAttributes) GetPROV_JCO_GROUP() string`

GetPROV_JCO_GROUP returns the PROV_JCO_GROUP field if non-nil, zero value otherwise.

### GetPROV_JCO_GROUPOk

`func (o *SAPConnectionAttributes) GetPROV_JCO_GROUPOk() (*string, bool)`

GetPROV_JCO_GROUPOk returns a tuple with the PROV_JCO_GROUP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPROV_JCO_GROUP

`func (o *SAPConnectionAttributes) SetPROV_JCO_GROUP(v string)`

SetPROV_JCO_GROUP sets PROV_JCO_GROUP field to given value.

### HasPROV_JCO_GROUP

`func (o *SAPConnectionAttributes) HasPROV_JCO_GROUP() bool`

HasPROV_JCO_GROUP returns a boolean if a field has been set.

### GetPROV_CUA_ENABLED

`func (o *SAPConnectionAttributes) GetPROV_CUA_ENABLED() string`

GetPROV_CUA_ENABLED returns the PROV_CUA_ENABLED field if non-nil, zero value otherwise.

### GetPROV_CUA_ENABLEDOk

`func (o *SAPConnectionAttributes) GetPROV_CUA_ENABLEDOk() (*string, bool)`

GetPROV_CUA_ENABLEDOk returns a tuple with the PROV_CUA_ENABLED field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPROV_CUA_ENABLED

`func (o *SAPConnectionAttributes) SetPROV_CUA_ENABLED(v string)`

SetPROV_CUA_ENABLED sets PROV_CUA_ENABLED field to given value.

### HasPROV_CUA_ENABLED

`func (o *SAPConnectionAttributes) HasPROV_CUA_ENABLED() bool`

HasPROV_CUA_ENABLED returns a boolean if a field has been set.

### GetJCO_MSHOST

`func (o *SAPConnectionAttributes) GetJCO_MSHOST() string`

GetJCO_MSHOST returns the JCO_MSHOST field if non-nil, zero value otherwise.

### GetJCO_MSHOSTOk

`func (o *SAPConnectionAttributes) GetJCO_MSHOSTOk() (*string, bool)`

GetJCO_MSHOSTOk returns a tuple with the JCO_MSHOST field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJCO_MSHOST

`func (o *SAPConnectionAttributes) SetJCO_MSHOST(v string)`

SetJCO_MSHOST sets JCO_MSHOST field to given value.

### HasJCO_MSHOST

`func (o *SAPConnectionAttributes) HasJCO_MSHOST() bool`

HasJCO_MSHOST returns a boolean if a field has been set.

### GetPROVJCOR3NAME

`func (o *SAPConnectionAttributes) GetPROVJCOR3NAME() string`

GetPROVJCOR3NAME returns the PROVJCOR3NAME field if non-nil, zero value otherwise.

### GetPROVJCOR3NAMEOk

`func (o *SAPConnectionAttributes) GetPROVJCOR3NAMEOk() (*string, bool)`

GetPROVJCOR3NAMEOk returns a tuple with the PROVJCOR3NAME field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPROVJCOR3NAME

`func (o *SAPConnectionAttributes) SetPROVJCOR3NAME(v string)`

SetPROVJCOR3NAME sets PROVJCOR3NAME field to given value.

### HasPROVJCOR3NAME

`func (o *SAPConnectionAttributes) HasPROVJCOR3NAME() bool`

HasPROVJCOR3NAME returns a boolean if a field has been set.

### GetPASSWORD_NOOFCAPSALPHA

`func (o *SAPConnectionAttributes) GetPASSWORD_NOOFCAPSALPHA() string`

GetPASSWORD_NOOFCAPSALPHA returns the PASSWORD_NOOFCAPSALPHA field if non-nil, zero value otherwise.

### GetPASSWORD_NOOFCAPSALPHAOk

`func (o *SAPConnectionAttributes) GetPASSWORD_NOOFCAPSALPHAOk() (*string, bool)`

GetPASSWORD_NOOFCAPSALPHAOk returns a tuple with the PASSWORD_NOOFCAPSALPHA field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPASSWORD_NOOFCAPSALPHA

`func (o *SAPConnectionAttributes) SetPASSWORD_NOOFCAPSALPHA(v string)`

SetPASSWORD_NOOFCAPSALPHA sets PASSWORD_NOOFCAPSALPHA field to given value.

### HasPASSWORD_NOOFCAPSALPHA

`func (o *SAPConnectionAttributes) HasPASSWORD_NOOFCAPSALPHA() bool`

HasPASSWORD_NOOFCAPSALPHA returns a boolean if a field has been set.

### GetMODIFYUSERDATAJSON

`func (o *SAPConnectionAttributes) GetMODIFYUSERDATAJSON() string`

GetMODIFYUSERDATAJSON returns the MODIFYUSERDATAJSON field if non-nil, zero value otherwise.

### GetMODIFYUSERDATAJSONOk

`func (o *SAPConnectionAttributes) GetMODIFYUSERDATAJSONOk() (*string, bool)`

GetMODIFYUSERDATAJSONOk returns a tuple with the MODIFYUSERDATAJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMODIFYUSERDATAJSON

`func (o *SAPConnectionAttributes) SetMODIFYUSERDATAJSON(v string)`

SetMODIFYUSERDATAJSON sets MODIFYUSERDATAJSON field to given value.

### HasMODIFYUSERDATAJSON

`func (o *SAPConnectionAttributes) HasMODIFYUSERDATAJSON() bool`

HasMODIFYUSERDATAJSON returns a boolean if a field has been set.

### GetIsTimeoutConfigValidated

`func (o *SAPConnectionAttributes) GetIsTimeoutConfigValidated() bool`

GetIsTimeoutConfigValidated returns the IsTimeoutConfigValidated field if non-nil, zero value otherwise.

### GetIsTimeoutConfigValidatedOk

`func (o *SAPConnectionAttributes) GetIsTimeoutConfigValidatedOk() (*bool, bool)`

GetIsTimeoutConfigValidatedOk returns a tuple with the IsTimeoutConfigValidated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsTimeoutConfigValidated

`func (o *SAPConnectionAttributes) SetIsTimeoutConfigValidated(v bool)`

SetIsTimeoutConfigValidated sets IsTimeoutConfigValidated field to given value.

### HasIsTimeoutConfigValidated

`func (o *SAPConnectionAttributes) HasIsTimeoutConfigValidated() bool`

HasIsTimeoutConfigValidated returns a boolean if a field has been set.

### GetJCO_SNC_QOP

`func (o *SAPConnectionAttributes) GetJCO_SNC_QOP() string`

GetJCO_SNC_QOP returns the JCO_SNC_QOP field if non-nil, zero value otherwise.

### GetJCO_SNC_QOPOk

`func (o *SAPConnectionAttributes) GetJCO_SNC_QOPOk() (*string, bool)`

GetJCO_SNC_QOPOk returns a tuple with the JCO_SNC_QOP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJCO_SNC_QOP

`func (o *SAPConnectionAttributes) SetJCO_SNC_QOP(v string)`

SetJCO_SNC_QOP sets JCO_SNC_QOP field to given value.

### HasJCO_SNC_QOP

`func (o *SAPConnectionAttributes) HasJCO_SNC_QOP() bool`

HasJCO_SNC_QOP returns a boolean if a field has been set.

### GetTABLES

`func (o *SAPConnectionAttributes) GetTABLES() string`

GetTABLES returns the TABLES field if non-nil, zero value otherwise.

### GetTABLESOk

`func (o *SAPConnectionAttributes) GetTABLESOk() (*string, bool)`

GetTABLESOk returns a tuple with the TABLES field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTABLES

`func (o *SAPConnectionAttributes) SetTABLES(v string)`

SetTABLES sets TABLES field to given value.

### HasTABLES

`func (o *SAPConnectionAttributes) HasTABLES() bool`

HasTABLES returns a boolean if a field has been set.

### GetPROV_JCO_LANG

`func (o *SAPConnectionAttributes) GetPROV_JCO_LANG() string`

GetPROV_JCO_LANG returns the PROV_JCO_LANG field if non-nil, zero value otherwise.

### GetPROV_JCO_LANGOk

`func (o *SAPConnectionAttributes) GetPROV_JCO_LANGOk() (*string, bool)`

GetPROV_JCO_LANGOk returns a tuple with the PROV_JCO_LANG field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPROV_JCO_LANG

`func (o *SAPConnectionAttributes) SetPROV_JCO_LANG(v string)`

SetPROV_JCO_LANG sets PROV_JCO_LANG field to given value.

### HasPROV_JCO_LANG

`func (o *SAPConnectionAttributes) HasPROV_JCO_LANG() bool`

HasPROV_JCO_LANG returns a boolean if a field has been set.

### GetJCO_SYSNR

`func (o *SAPConnectionAttributes) GetJCO_SYSNR() string`

GetJCO_SYSNR returns the JCO_SYSNR field if non-nil, zero value otherwise.

### GetJCO_SYSNROk

`func (o *SAPConnectionAttributes) GetJCO_SYSNROk() (*string, bool)`

GetJCO_SYSNROk returns a tuple with the JCO_SYSNR field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJCO_SYSNR

`func (o *SAPConnectionAttributes) SetJCO_SYSNR(v string)`

SetJCO_SYSNR sets JCO_SYSNR field to given value.

### HasJCO_SYSNR

`func (o *SAPConnectionAttributes) HasJCO_SYSNR() bool`

HasJCO_SYSNR returns a boolean if a field has been set.

### GetEXTERNAL_SOD_EVAL_JSON_DETAIL

`func (o *SAPConnectionAttributes) GetEXTERNAL_SOD_EVAL_JSON_DETAIL() string`

GetEXTERNAL_SOD_EVAL_JSON_DETAIL returns the EXTERNAL_SOD_EVAL_JSON_DETAIL field if non-nil, zero value otherwise.

### GetEXTERNAL_SOD_EVAL_JSON_DETAILOk

`func (o *SAPConnectionAttributes) GetEXTERNAL_SOD_EVAL_JSON_DETAILOk() (*string, bool)`

GetEXTERNAL_SOD_EVAL_JSON_DETAILOk returns a tuple with the EXTERNAL_SOD_EVAL_JSON_DETAIL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEXTERNAL_SOD_EVAL_JSON_DETAIL

`func (o *SAPConnectionAttributes) SetEXTERNAL_SOD_EVAL_JSON_DETAIL(v string)`

SetEXTERNAL_SOD_EVAL_JSON_DETAIL sets EXTERNAL_SOD_EVAL_JSON_DETAIL field to given value.

### HasEXTERNAL_SOD_EVAL_JSON_DETAIL

`func (o *SAPConnectionAttributes) HasEXTERNAL_SOD_EVAL_JSON_DETAIL() bool`

HasEXTERNAL_SOD_EVAL_JSON_DETAIL returns a boolean if a field has been set.

### GetDATA_IMPORT_FILTER

`func (o *SAPConnectionAttributes) GetDATA_IMPORT_FILTER() string`

GetDATA_IMPORT_FILTER returns the DATA_IMPORT_FILTER field if non-nil, zero value otherwise.

### GetDATA_IMPORT_FILTEROk

`func (o *SAPConnectionAttributes) GetDATA_IMPORT_FILTEROk() (*string, bool)`

GetDATA_IMPORT_FILTEROk returns a tuple with the DATA_IMPORT_FILTER field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDATA_IMPORT_FILTER

`func (o *SAPConnectionAttributes) SetDATA_IMPORT_FILTER(v string)`

SetDATA_IMPORT_FILTER sets DATA_IMPORT_FILTER field to given value.

### HasDATA_IMPORT_FILTER

`func (o *SAPConnectionAttributes) HasDATA_IMPORT_FILTER() bool`

HasDATA_IMPORT_FILTER returns a boolean if a field has been set.

### GetENABLEACCOUNTJSON

`func (o *SAPConnectionAttributes) GetENABLEACCOUNTJSON() string`

GetENABLEACCOUNTJSON returns the ENABLEACCOUNTJSON field if non-nil, zero value otherwise.

### GetENABLEACCOUNTJSONOk

`func (o *SAPConnectionAttributes) GetENABLEACCOUNTJSONOk() (*string, bool)`

GetENABLEACCOUNTJSONOk returns a tuple with the ENABLEACCOUNTJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetENABLEACCOUNTJSON

`func (o *SAPConnectionAttributes) SetENABLEACCOUNTJSON(v string)`

SetENABLEACCOUNTJSON sets ENABLEACCOUNTJSON field to given value.

### HasENABLEACCOUNTJSON

`func (o *SAPConnectionAttributes) HasENABLEACCOUNTJSON() bool`

HasENABLEACCOUNTJSON returns a boolean if a field has been set.

### GetALTERNATE_OUTPUT_PARAMETER_ET_DATA

`func (o *SAPConnectionAttributes) GetALTERNATE_OUTPUT_PARAMETER_ET_DATA() string`

GetALTERNATE_OUTPUT_PARAMETER_ET_DATA returns the ALTERNATE_OUTPUT_PARAMETER_ET_DATA field if non-nil, zero value otherwise.

### GetALTERNATE_OUTPUT_PARAMETER_ET_DATAOk

`func (o *SAPConnectionAttributes) GetALTERNATE_OUTPUT_PARAMETER_ET_DATAOk() (*string, bool)`

GetALTERNATE_OUTPUT_PARAMETER_ET_DATAOk returns a tuple with the ALTERNATE_OUTPUT_PARAMETER_ET_DATA field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetALTERNATE_OUTPUT_PARAMETER_ET_DATA

`func (o *SAPConnectionAttributes) SetALTERNATE_OUTPUT_PARAMETER_ET_DATA(v string)`

SetALTERNATE_OUTPUT_PARAMETER_ET_DATA sets ALTERNATE_OUTPUT_PARAMETER_ET_DATA field to given value.

### HasALTERNATE_OUTPUT_PARAMETER_ET_DATA

`func (o *SAPConnectionAttributes) HasALTERNATE_OUTPUT_PARAMETER_ET_DATA() bool`

HasALTERNATE_OUTPUT_PARAMETER_ET_DATA returns a boolean if a field has been set.

### GetJCO_GROUP

`func (o *SAPConnectionAttributes) GetJCO_GROUP() string`

GetJCO_GROUP returns the JCO_GROUP field if non-nil, zero value otherwise.

### GetJCO_GROUPOk

`func (o *SAPConnectionAttributes) GetJCO_GROUPOk() (*string, bool)`

GetJCO_GROUPOk returns a tuple with the JCO_GROUP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJCO_GROUP

`func (o *SAPConnectionAttributes) SetJCO_GROUP(v string)`

SetJCO_GROUP sets JCO_GROUP field to given value.

### HasJCO_GROUP

`func (o *SAPConnectionAttributes) HasJCO_GROUP() bool`

HasJCO_GROUP returns a boolean if a field has been set.

### GetPASSWORD_MAX_LENGTH

`func (o *SAPConnectionAttributes) GetPASSWORD_MAX_LENGTH() string`

GetPASSWORD_MAX_LENGTH returns the PASSWORD_MAX_LENGTH field if non-nil, zero value otherwise.

### GetPASSWORD_MAX_LENGTHOk

`func (o *SAPConnectionAttributes) GetPASSWORD_MAX_LENGTHOk() (*string, bool)`

GetPASSWORD_MAX_LENGTHOk returns a tuple with the PASSWORD_MAX_LENGTH field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPASSWORD_MAX_LENGTH

`func (o *SAPConnectionAttributes) SetPASSWORD_MAX_LENGTH(v string)`

SetPASSWORD_MAX_LENGTH sets PASSWORD_MAX_LENGTH field to given value.

### HasPASSWORD_MAX_LENGTH

`func (o *SAPConnectionAttributes) HasPASSWORD_MAX_LENGTH() bool`

HasPASSWORD_MAX_LENGTH returns a boolean if a field has been set.

### GetUSERIMPORTJSON

`func (o *SAPConnectionAttributes) GetUSERIMPORTJSON() string`

GetUSERIMPORTJSON returns the USERIMPORTJSON field if non-nil, zero value otherwise.

### GetUSERIMPORTJSONOk

`func (o *SAPConnectionAttributes) GetUSERIMPORTJSONOk() (*string, bool)`

GetUSERIMPORTJSONOk returns a tuple with the USERIMPORTJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUSERIMPORTJSON

`func (o *SAPConnectionAttributes) SetUSERIMPORTJSON(v string)`

SetUSERIMPORTJSON sets USERIMPORTJSON field to given value.

### HasUSERIMPORTJSON

`func (o *SAPConnectionAttributes) HasUSERIMPORTJSON() bool`

HasUSERIMPORTJSON returns a boolean if a field has been set.

### GetSYSTEMNAME

`func (o *SAPConnectionAttributes) GetSYSTEMNAME() string`

GetSYSTEMNAME returns the SYSTEMNAME field if non-nil, zero value otherwise.

### GetSYSTEMNAMEOk

`func (o *SAPConnectionAttributes) GetSYSTEMNAMEOk() (*string, bool)`

GetSYSTEMNAMEOk returns a tuple with the SYSTEMNAME field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSYSTEMNAME

`func (o *SAPConnectionAttributes) SetSYSTEMNAME(v string)`

SetSYSTEMNAME sets SYSTEMNAME field to given value.

### HasSYSTEMNAME

`func (o *SAPConnectionAttributes) HasSYSTEMNAME() bool`

HasSYSTEMNAME returns a boolean if a field has been set.

### GetUPDATEACCOUNTJSON

`func (o *SAPConnectionAttributes) GetUPDATEACCOUNTJSON() string`

GetUPDATEACCOUNTJSON returns the UPDATEACCOUNTJSON field if non-nil, zero value otherwise.

### GetUPDATEACCOUNTJSONOk

`func (o *SAPConnectionAttributes) GetUPDATEACCOUNTJSONOk() (*string, bool)`

GetUPDATEACCOUNTJSONOk returns a tuple with the UPDATEACCOUNTJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUPDATEACCOUNTJSON

`func (o *SAPConnectionAttributes) SetUPDATEACCOUNTJSON(v string)`

SetUPDATEACCOUNTJSON sets UPDATEACCOUNTJSON field to given value.

### HasUPDATEACCOUNTJSON

`func (o *SAPConnectionAttributes) HasUPDATEACCOUNTJSON() bool`

HasUPDATEACCOUNTJSON returns a boolean if a field has been set.

### GetROLE_DEFAULT_DATE

`func (o *SAPConnectionAttributes) GetROLE_DEFAULT_DATE() string`

GetROLE_DEFAULT_DATE returns the ROLE_DEFAULT_DATE field if non-nil, zero value otherwise.

### GetROLE_DEFAULT_DATEOk

`func (o *SAPConnectionAttributes) GetROLE_DEFAULT_DATEOk() (*string, bool)`

GetROLE_DEFAULT_DATEOk returns a tuple with the ROLE_DEFAULT_DATE field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetROLE_DEFAULT_DATE

`func (o *SAPConnectionAttributes) SetROLE_DEFAULT_DATE(v string)`

SetROLE_DEFAULT_DATE sets ROLE_DEFAULT_DATE field to given value.

### HasROLE_DEFAULT_DATE

`func (o *SAPConnectionAttributes) HasROLE_DEFAULT_DATE() bool`

HasROLE_DEFAULT_DATE returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


