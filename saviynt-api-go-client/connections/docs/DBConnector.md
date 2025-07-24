# DBConnector

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**URL** | **string** | Host Name for connection | 
**USERNAME** | **string** | Username for connection | 
**PASSWORD** | **string** | Password for connection | 
**DRIVERNAME** | **string** | Driver name for the connection | 
**CONNECTIONPROPERTIES** | Pointer to **string** | Properties that needs to be added when connecting to the database | [optional] 
**PASSWORD_MIN_LENGTH** | Pointer to **string** | Specify the Min length for the random password | [optional] 
**PASSWORD_MAX_LENGTH** | Pointer to **string** | Specify the Max length for the random password | [optional] 
**PASSWORD_NOOFCAPSALPHA** | Pointer to **string** | Specify the Number of Upper case alphabets required for the random password | [optional] 
**PASSWORD_NOOFDIGITS** | Pointer to **string** | Specify the Number of digits required for the random password | [optional] 
**PASSWORD_NOOFSPLCHARS** | Pointer to **string** | Specify the Number of special chars required for the random password | [optional] 
**CREATEACCOUNTJSON** | Pointer to **string** | JSON to specify the Queries/stored procedures which will be used to Create the New Account,Objects Exposed-(randomPassword,task,user,accountName,role,endpoint and all the objects defined in Dynamic Attributes ) | [optional] 
**UPDATEACCOUNTJSON** | Pointer to **string** | JSON to specify the Queries/stored procedures which will be used to Update an existing Account,Objects Exposed-(randomPassword,task,user,accountName,role,endpoint and all the objects defined in Dynamic Attributes ). | [optional] 
**GRANTACCESSJSON** | Pointer to **string** | JSON to specify the Queries/stored procedures which will be used to provide acccess,Objects Exposed-(task,user,accountName,role,endpoint and all the objects defined in Dynamic Attributes ).  | [optional] 
**REVOKEACCESSJSON** | Pointer to **string** | JSON to specify the Queries/stored procedures which will be used to revoke access,Objects Exposed-(task,user,accountName,role,endpoint and all the objects defined in Dynamic Attributes ). | [optional] 
**CHANGEPASSJSON** | Pointer to **string** | JSON to specify the Queries/stored procedures which will be used to change password,Objects Exposed-(randomPassword,task,user,accountName,role,endpoint and all the objects defined in Dynamic Attributes ).  | [optional] 
**DELETEACCOUNTJSON** | Pointer to **string** | JSON to specify the Queries/stored procedures which will be used to delete an account,Objects Exposed-(task,user,accountName,role,endpoint and all the objects defined in Dynamic Attributes ). | [optional] 
**ENABLEACCOUNTJSON** | Pointer to **string** | JSON to specify the Queries/stored procedures which will be used to Enable,Objects Exposed-(task,user,accountName,role,endpoint and all the objects defined in Dynamic Attributes ). | [optional] 
**DISABLEACCOUNTJSON** | Pointer to **string** | JSON to specify the Queries/stored procedures which will be used to Disable Account,Objects Exposed-(task,user,accountName,role,endpoint and all the objects defined in Dynamic Attributes ). | [optional] 
**ACCOUNTEXISTSJSON** | Pointer to **string** | JSON to specify the Query which will be used to check whether an account exists,Objects Exposed-(task,user,accountName,role,endpoint and all the objects defined in Dynamic Attributes ). | [optional] 
**UPDATEUSERJSON** | Pointer to **string** | JSON to specify the Queries/stored procedures which will be used to Update an existing Account,Objects Exposed-(randomPassword,task,user,accountName,role,updatetaskuser,endpoint and all the objects defined in Dynamic Attributes ). | [optional] 
**ACCOUNTSIMPORT** | Pointer to **string** | Accounts Import XML file content | [optional] 
**ENTITLEMENTVALUEIMPORT** | Pointer to **string** | Entitlement Value Import XML file content | [optional] 
**ROLEOWNERIMPORT** | Pointer to **string** | Role Owner Import XML file contentT | [optional] 
**ROLESIMPORT** | Pointer to **string** | Roles Import XML file content | [optional] 
**SYSTEMIMPORT** | Pointer to **string** | System Import XML file content | [optional] 
**USERIMPORT** | Pointer to **string** | User Import XML file content | [optional] 
**MODIFYUSERDATAJSON** | Pointer to **string** | Property for MODIFYUSERDATAJSON | [optional] 
**STATUS_THRESHOLD_CONFIG** | Pointer to **string** | Applicable for Accounts full import only.If this config isdefined with status/threshold values, it will take precedence over account_not_in_file_action defined in ACCOUNTIMPORT xml.If this config is defined with only correlateInactiveAccounts, then account_not_in_file_action will used as normal.The attributes of statusAndThresholdConfig json are:statusColumn: Property in saviynt which stores the status of target system.activeStatus:All possible values that denotes the Active status of the target system.accountThresholdValue: No. of accounts to be deleted in Saviynt &gt;&#x3D; accountThresholdValue, it performs NO ACTION, else it disables the accounts.inactivateAccountsNotInFile: If true,accounts not in file are marked as Inactive. If false, accounts not in file are marked as SUSPENDED FROM IMPORT SERVICE.CorrelateInactiveAccounts: If true, correlates disabled accounts as well with the users. | [optional] 
**MAX_PAGINATION_SIZE** | Pointer to **string** | Defines the max number of records from the target to be processed in each page | [optional] 
**CLI_COMMAND_JSON** | Pointer to **string** | JSON to specify the commands which can be executed in target server. | [optional] 
**CREATEENTITLEMENTJSON** | Pointer to **string** |  | [optional] 
**DELETEENTITLEMENTJSON** | Pointer to **string** |  | [optional] 
**ENTITLEMENTEXISTJSON** | Pointer to **string** |  | [optional] 
**UPDATEENTITLEMENTJSON** | Pointer to **string** |  | [optional] 

## Methods

### NewDBConnector

`func NewDBConnector(uRL string, uSERNAME string, pASSWORD string, dRIVERNAME string, ) *DBConnector`

NewDBConnector instantiates a new DBConnector object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDBConnectorWithDefaults

`func NewDBConnectorWithDefaults() *DBConnector`

NewDBConnectorWithDefaults instantiates a new DBConnector object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetURL

`func (o *DBConnector) GetURL() string`

GetURL returns the URL field if non-nil, zero value otherwise.

### GetURLOk

`func (o *DBConnector) GetURLOk() (*string, bool)`

GetURLOk returns a tuple with the URL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetURL

`func (o *DBConnector) SetURL(v string)`

SetURL sets URL field to given value.


### GetUSERNAME

`func (o *DBConnector) GetUSERNAME() string`

GetUSERNAME returns the USERNAME field if non-nil, zero value otherwise.

### GetUSERNAMEOk

`func (o *DBConnector) GetUSERNAMEOk() (*string, bool)`

GetUSERNAMEOk returns a tuple with the USERNAME field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUSERNAME

`func (o *DBConnector) SetUSERNAME(v string)`

SetUSERNAME sets USERNAME field to given value.


### GetPASSWORD

`func (o *DBConnector) GetPASSWORD() string`

GetPASSWORD returns the PASSWORD field if non-nil, zero value otherwise.

### GetPASSWORDOk

`func (o *DBConnector) GetPASSWORDOk() (*string, bool)`

GetPASSWORDOk returns a tuple with the PASSWORD field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPASSWORD

`func (o *DBConnector) SetPASSWORD(v string)`

SetPASSWORD sets PASSWORD field to given value.


### GetDRIVERNAME

`func (o *DBConnector) GetDRIVERNAME() string`

GetDRIVERNAME returns the DRIVERNAME field if non-nil, zero value otherwise.

### GetDRIVERNAMEOk

`func (o *DBConnector) GetDRIVERNAMEOk() (*string, bool)`

GetDRIVERNAMEOk returns a tuple with the DRIVERNAME field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDRIVERNAME

`func (o *DBConnector) SetDRIVERNAME(v string)`

SetDRIVERNAME sets DRIVERNAME field to given value.


### GetCONNECTIONPROPERTIES

`func (o *DBConnector) GetCONNECTIONPROPERTIES() string`

GetCONNECTIONPROPERTIES returns the CONNECTIONPROPERTIES field if non-nil, zero value otherwise.

### GetCONNECTIONPROPERTIESOk

`func (o *DBConnector) GetCONNECTIONPROPERTIESOk() (*string, bool)`

GetCONNECTIONPROPERTIESOk returns a tuple with the CONNECTIONPROPERTIES field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCONNECTIONPROPERTIES

`func (o *DBConnector) SetCONNECTIONPROPERTIES(v string)`

SetCONNECTIONPROPERTIES sets CONNECTIONPROPERTIES field to given value.

### HasCONNECTIONPROPERTIES

`func (o *DBConnector) HasCONNECTIONPROPERTIES() bool`

HasCONNECTIONPROPERTIES returns a boolean if a field has been set.

### GetPASSWORD_MIN_LENGTH

`func (o *DBConnector) GetPASSWORD_MIN_LENGTH() string`

GetPASSWORD_MIN_LENGTH returns the PASSWORD_MIN_LENGTH field if non-nil, zero value otherwise.

### GetPASSWORD_MIN_LENGTHOk

`func (o *DBConnector) GetPASSWORD_MIN_LENGTHOk() (*string, bool)`

GetPASSWORD_MIN_LENGTHOk returns a tuple with the PASSWORD_MIN_LENGTH field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPASSWORD_MIN_LENGTH

`func (o *DBConnector) SetPASSWORD_MIN_LENGTH(v string)`

SetPASSWORD_MIN_LENGTH sets PASSWORD_MIN_LENGTH field to given value.

### HasPASSWORD_MIN_LENGTH

`func (o *DBConnector) HasPASSWORD_MIN_LENGTH() bool`

HasPASSWORD_MIN_LENGTH returns a boolean if a field has been set.

### GetPASSWORD_MAX_LENGTH

`func (o *DBConnector) GetPASSWORD_MAX_LENGTH() string`

GetPASSWORD_MAX_LENGTH returns the PASSWORD_MAX_LENGTH field if non-nil, zero value otherwise.

### GetPASSWORD_MAX_LENGTHOk

`func (o *DBConnector) GetPASSWORD_MAX_LENGTHOk() (*string, bool)`

GetPASSWORD_MAX_LENGTHOk returns a tuple with the PASSWORD_MAX_LENGTH field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPASSWORD_MAX_LENGTH

`func (o *DBConnector) SetPASSWORD_MAX_LENGTH(v string)`

SetPASSWORD_MAX_LENGTH sets PASSWORD_MAX_LENGTH field to given value.

### HasPASSWORD_MAX_LENGTH

`func (o *DBConnector) HasPASSWORD_MAX_LENGTH() bool`

HasPASSWORD_MAX_LENGTH returns a boolean if a field has been set.

### GetPASSWORD_NOOFCAPSALPHA

`func (o *DBConnector) GetPASSWORD_NOOFCAPSALPHA() string`

GetPASSWORD_NOOFCAPSALPHA returns the PASSWORD_NOOFCAPSALPHA field if non-nil, zero value otherwise.

### GetPASSWORD_NOOFCAPSALPHAOk

`func (o *DBConnector) GetPASSWORD_NOOFCAPSALPHAOk() (*string, bool)`

GetPASSWORD_NOOFCAPSALPHAOk returns a tuple with the PASSWORD_NOOFCAPSALPHA field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPASSWORD_NOOFCAPSALPHA

`func (o *DBConnector) SetPASSWORD_NOOFCAPSALPHA(v string)`

SetPASSWORD_NOOFCAPSALPHA sets PASSWORD_NOOFCAPSALPHA field to given value.

### HasPASSWORD_NOOFCAPSALPHA

`func (o *DBConnector) HasPASSWORD_NOOFCAPSALPHA() bool`

HasPASSWORD_NOOFCAPSALPHA returns a boolean if a field has been set.

### GetPASSWORD_NOOFDIGITS

`func (o *DBConnector) GetPASSWORD_NOOFDIGITS() string`

GetPASSWORD_NOOFDIGITS returns the PASSWORD_NOOFDIGITS field if non-nil, zero value otherwise.

### GetPASSWORD_NOOFDIGITSOk

`func (o *DBConnector) GetPASSWORD_NOOFDIGITSOk() (*string, bool)`

GetPASSWORD_NOOFDIGITSOk returns a tuple with the PASSWORD_NOOFDIGITS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPASSWORD_NOOFDIGITS

`func (o *DBConnector) SetPASSWORD_NOOFDIGITS(v string)`

SetPASSWORD_NOOFDIGITS sets PASSWORD_NOOFDIGITS field to given value.

### HasPASSWORD_NOOFDIGITS

`func (o *DBConnector) HasPASSWORD_NOOFDIGITS() bool`

HasPASSWORD_NOOFDIGITS returns a boolean if a field has been set.

### GetPASSWORD_NOOFSPLCHARS

`func (o *DBConnector) GetPASSWORD_NOOFSPLCHARS() string`

GetPASSWORD_NOOFSPLCHARS returns the PASSWORD_NOOFSPLCHARS field if non-nil, zero value otherwise.

### GetPASSWORD_NOOFSPLCHARSOk

`func (o *DBConnector) GetPASSWORD_NOOFSPLCHARSOk() (*string, bool)`

GetPASSWORD_NOOFSPLCHARSOk returns a tuple with the PASSWORD_NOOFSPLCHARS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPASSWORD_NOOFSPLCHARS

`func (o *DBConnector) SetPASSWORD_NOOFSPLCHARS(v string)`

SetPASSWORD_NOOFSPLCHARS sets PASSWORD_NOOFSPLCHARS field to given value.

### HasPASSWORD_NOOFSPLCHARS

`func (o *DBConnector) HasPASSWORD_NOOFSPLCHARS() bool`

HasPASSWORD_NOOFSPLCHARS returns a boolean if a field has been set.

### GetCREATEACCOUNTJSON

`func (o *DBConnector) GetCREATEACCOUNTJSON() string`

GetCREATEACCOUNTJSON returns the CREATEACCOUNTJSON field if non-nil, zero value otherwise.

### GetCREATEACCOUNTJSONOk

`func (o *DBConnector) GetCREATEACCOUNTJSONOk() (*string, bool)`

GetCREATEACCOUNTJSONOk returns a tuple with the CREATEACCOUNTJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCREATEACCOUNTJSON

`func (o *DBConnector) SetCREATEACCOUNTJSON(v string)`

SetCREATEACCOUNTJSON sets CREATEACCOUNTJSON field to given value.

### HasCREATEACCOUNTJSON

`func (o *DBConnector) HasCREATEACCOUNTJSON() bool`

HasCREATEACCOUNTJSON returns a boolean if a field has been set.

### GetUPDATEACCOUNTJSON

`func (o *DBConnector) GetUPDATEACCOUNTJSON() string`

GetUPDATEACCOUNTJSON returns the UPDATEACCOUNTJSON field if non-nil, zero value otherwise.

### GetUPDATEACCOUNTJSONOk

`func (o *DBConnector) GetUPDATEACCOUNTJSONOk() (*string, bool)`

GetUPDATEACCOUNTJSONOk returns a tuple with the UPDATEACCOUNTJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUPDATEACCOUNTJSON

`func (o *DBConnector) SetUPDATEACCOUNTJSON(v string)`

SetUPDATEACCOUNTJSON sets UPDATEACCOUNTJSON field to given value.

### HasUPDATEACCOUNTJSON

`func (o *DBConnector) HasUPDATEACCOUNTJSON() bool`

HasUPDATEACCOUNTJSON returns a boolean if a field has been set.

### GetGRANTACCESSJSON

`func (o *DBConnector) GetGRANTACCESSJSON() string`

GetGRANTACCESSJSON returns the GRANTACCESSJSON field if non-nil, zero value otherwise.

### GetGRANTACCESSJSONOk

`func (o *DBConnector) GetGRANTACCESSJSONOk() (*string, bool)`

GetGRANTACCESSJSONOk returns a tuple with the GRANTACCESSJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGRANTACCESSJSON

`func (o *DBConnector) SetGRANTACCESSJSON(v string)`

SetGRANTACCESSJSON sets GRANTACCESSJSON field to given value.

### HasGRANTACCESSJSON

`func (o *DBConnector) HasGRANTACCESSJSON() bool`

HasGRANTACCESSJSON returns a boolean if a field has been set.

### GetREVOKEACCESSJSON

`func (o *DBConnector) GetREVOKEACCESSJSON() string`

GetREVOKEACCESSJSON returns the REVOKEACCESSJSON field if non-nil, zero value otherwise.

### GetREVOKEACCESSJSONOk

`func (o *DBConnector) GetREVOKEACCESSJSONOk() (*string, bool)`

GetREVOKEACCESSJSONOk returns a tuple with the REVOKEACCESSJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetREVOKEACCESSJSON

`func (o *DBConnector) SetREVOKEACCESSJSON(v string)`

SetREVOKEACCESSJSON sets REVOKEACCESSJSON field to given value.

### HasREVOKEACCESSJSON

`func (o *DBConnector) HasREVOKEACCESSJSON() bool`

HasREVOKEACCESSJSON returns a boolean if a field has been set.

### GetCHANGEPASSJSON

`func (o *DBConnector) GetCHANGEPASSJSON() string`

GetCHANGEPASSJSON returns the CHANGEPASSJSON field if non-nil, zero value otherwise.

### GetCHANGEPASSJSONOk

`func (o *DBConnector) GetCHANGEPASSJSONOk() (*string, bool)`

GetCHANGEPASSJSONOk returns a tuple with the CHANGEPASSJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCHANGEPASSJSON

`func (o *DBConnector) SetCHANGEPASSJSON(v string)`

SetCHANGEPASSJSON sets CHANGEPASSJSON field to given value.

### HasCHANGEPASSJSON

`func (o *DBConnector) HasCHANGEPASSJSON() bool`

HasCHANGEPASSJSON returns a boolean if a field has been set.

### GetDELETEACCOUNTJSON

`func (o *DBConnector) GetDELETEACCOUNTJSON() string`

GetDELETEACCOUNTJSON returns the DELETEACCOUNTJSON field if non-nil, zero value otherwise.

### GetDELETEACCOUNTJSONOk

`func (o *DBConnector) GetDELETEACCOUNTJSONOk() (*string, bool)`

GetDELETEACCOUNTJSONOk returns a tuple with the DELETEACCOUNTJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDELETEACCOUNTJSON

`func (o *DBConnector) SetDELETEACCOUNTJSON(v string)`

SetDELETEACCOUNTJSON sets DELETEACCOUNTJSON field to given value.

### HasDELETEACCOUNTJSON

`func (o *DBConnector) HasDELETEACCOUNTJSON() bool`

HasDELETEACCOUNTJSON returns a boolean if a field has been set.

### GetENABLEACCOUNTJSON

`func (o *DBConnector) GetENABLEACCOUNTJSON() string`

GetENABLEACCOUNTJSON returns the ENABLEACCOUNTJSON field if non-nil, zero value otherwise.

### GetENABLEACCOUNTJSONOk

`func (o *DBConnector) GetENABLEACCOUNTJSONOk() (*string, bool)`

GetENABLEACCOUNTJSONOk returns a tuple with the ENABLEACCOUNTJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetENABLEACCOUNTJSON

`func (o *DBConnector) SetENABLEACCOUNTJSON(v string)`

SetENABLEACCOUNTJSON sets ENABLEACCOUNTJSON field to given value.

### HasENABLEACCOUNTJSON

`func (o *DBConnector) HasENABLEACCOUNTJSON() bool`

HasENABLEACCOUNTJSON returns a boolean if a field has been set.

### GetDISABLEACCOUNTJSON

`func (o *DBConnector) GetDISABLEACCOUNTJSON() string`

GetDISABLEACCOUNTJSON returns the DISABLEACCOUNTJSON field if non-nil, zero value otherwise.

### GetDISABLEACCOUNTJSONOk

`func (o *DBConnector) GetDISABLEACCOUNTJSONOk() (*string, bool)`

GetDISABLEACCOUNTJSONOk returns a tuple with the DISABLEACCOUNTJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDISABLEACCOUNTJSON

`func (o *DBConnector) SetDISABLEACCOUNTJSON(v string)`

SetDISABLEACCOUNTJSON sets DISABLEACCOUNTJSON field to given value.

### HasDISABLEACCOUNTJSON

`func (o *DBConnector) HasDISABLEACCOUNTJSON() bool`

HasDISABLEACCOUNTJSON returns a boolean if a field has been set.

### GetACCOUNTEXISTSJSON

`func (o *DBConnector) GetACCOUNTEXISTSJSON() string`

GetACCOUNTEXISTSJSON returns the ACCOUNTEXISTSJSON field if non-nil, zero value otherwise.

### GetACCOUNTEXISTSJSONOk

`func (o *DBConnector) GetACCOUNTEXISTSJSONOk() (*string, bool)`

GetACCOUNTEXISTSJSONOk returns a tuple with the ACCOUNTEXISTSJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetACCOUNTEXISTSJSON

`func (o *DBConnector) SetACCOUNTEXISTSJSON(v string)`

SetACCOUNTEXISTSJSON sets ACCOUNTEXISTSJSON field to given value.

### HasACCOUNTEXISTSJSON

`func (o *DBConnector) HasACCOUNTEXISTSJSON() bool`

HasACCOUNTEXISTSJSON returns a boolean if a field has been set.

### GetUPDATEUSERJSON

`func (o *DBConnector) GetUPDATEUSERJSON() string`

GetUPDATEUSERJSON returns the UPDATEUSERJSON field if non-nil, zero value otherwise.

### GetUPDATEUSERJSONOk

`func (o *DBConnector) GetUPDATEUSERJSONOk() (*string, bool)`

GetUPDATEUSERJSONOk returns a tuple with the UPDATEUSERJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUPDATEUSERJSON

`func (o *DBConnector) SetUPDATEUSERJSON(v string)`

SetUPDATEUSERJSON sets UPDATEUSERJSON field to given value.

### HasUPDATEUSERJSON

`func (o *DBConnector) HasUPDATEUSERJSON() bool`

HasUPDATEUSERJSON returns a boolean if a field has been set.

### GetACCOUNTSIMPORT

`func (o *DBConnector) GetACCOUNTSIMPORT() string`

GetACCOUNTSIMPORT returns the ACCOUNTSIMPORT field if non-nil, zero value otherwise.

### GetACCOUNTSIMPORTOk

`func (o *DBConnector) GetACCOUNTSIMPORTOk() (*string, bool)`

GetACCOUNTSIMPORTOk returns a tuple with the ACCOUNTSIMPORT field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetACCOUNTSIMPORT

`func (o *DBConnector) SetACCOUNTSIMPORT(v string)`

SetACCOUNTSIMPORT sets ACCOUNTSIMPORT field to given value.

### HasACCOUNTSIMPORT

`func (o *DBConnector) HasACCOUNTSIMPORT() bool`

HasACCOUNTSIMPORT returns a boolean if a field has been set.

### GetENTITLEMENTVALUEIMPORT

`func (o *DBConnector) GetENTITLEMENTVALUEIMPORT() string`

GetENTITLEMENTVALUEIMPORT returns the ENTITLEMENTVALUEIMPORT field if non-nil, zero value otherwise.

### GetENTITLEMENTVALUEIMPORTOk

`func (o *DBConnector) GetENTITLEMENTVALUEIMPORTOk() (*string, bool)`

GetENTITLEMENTVALUEIMPORTOk returns a tuple with the ENTITLEMENTVALUEIMPORT field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetENTITLEMENTVALUEIMPORT

`func (o *DBConnector) SetENTITLEMENTVALUEIMPORT(v string)`

SetENTITLEMENTVALUEIMPORT sets ENTITLEMENTVALUEIMPORT field to given value.

### HasENTITLEMENTVALUEIMPORT

`func (o *DBConnector) HasENTITLEMENTVALUEIMPORT() bool`

HasENTITLEMENTVALUEIMPORT returns a boolean if a field has been set.

### GetROLEOWNERIMPORT

`func (o *DBConnector) GetROLEOWNERIMPORT() string`

GetROLEOWNERIMPORT returns the ROLEOWNERIMPORT field if non-nil, zero value otherwise.

### GetROLEOWNERIMPORTOk

`func (o *DBConnector) GetROLEOWNERIMPORTOk() (*string, bool)`

GetROLEOWNERIMPORTOk returns a tuple with the ROLEOWNERIMPORT field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetROLEOWNERIMPORT

`func (o *DBConnector) SetROLEOWNERIMPORT(v string)`

SetROLEOWNERIMPORT sets ROLEOWNERIMPORT field to given value.

### HasROLEOWNERIMPORT

`func (o *DBConnector) HasROLEOWNERIMPORT() bool`

HasROLEOWNERIMPORT returns a boolean if a field has been set.

### GetROLESIMPORT

`func (o *DBConnector) GetROLESIMPORT() string`

GetROLESIMPORT returns the ROLESIMPORT field if non-nil, zero value otherwise.

### GetROLESIMPORTOk

`func (o *DBConnector) GetROLESIMPORTOk() (*string, bool)`

GetROLESIMPORTOk returns a tuple with the ROLESIMPORT field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetROLESIMPORT

`func (o *DBConnector) SetROLESIMPORT(v string)`

SetROLESIMPORT sets ROLESIMPORT field to given value.

### HasROLESIMPORT

`func (o *DBConnector) HasROLESIMPORT() bool`

HasROLESIMPORT returns a boolean if a field has been set.

### GetSYSTEMIMPORT

`func (o *DBConnector) GetSYSTEMIMPORT() string`

GetSYSTEMIMPORT returns the SYSTEMIMPORT field if non-nil, zero value otherwise.

### GetSYSTEMIMPORTOk

`func (o *DBConnector) GetSYSTEMIMPORTOk() (*string, bool)`

GetSYSTEMIMPORTOk returns a tuple with the SYSTEMIMPORT field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSYSTEMIMPORT

`func (o *DBConnector) SetSYSTEMIMPORT(v string)`

SetSYSTEMIMPORT sets SYSTEMIMPORT field to given value.

### HasSYSTEMIMPORT

`func (o *DBConnector) HasSYSTEMIMPORT() bool`

HasSYSTEMIMPORT returns a boolean if a field has been set.

### GetUSERIMPORT

`func (o *DBConnector) GetUSERIMPORT() string`

GetUSERIMPORT returns the USERIMPORT field if non-nil, zero value otherwise.

### GetUSERIMPORTOk

`func (o *DBConnector) GetUSERIMPORTOk() (*string, bool)`

GetUSERIMPORTOk returns a tuple with the USERIMPORT field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUSERIMPORT

`func (o *DBConnector) SetUSERIMPORT(v string)`

SetUSERIMPORT sets USERIMPORT field to given value.

### HasUSERIMPORT

`func (o *DBConnector) HasUSERIMPORT() bool`

HasUSERIMPORT returns a boolean if a field has been set.

### GetMODIFYUSERDATAJSON

`func (o *DBConnector) GetMODIFYUSERDATAJSON() string`

GetMODIFYUSERDATAJSON returns the MODIFYUSERDATAJSON field if non-nil, zero value otherwise.

### GetMODIFYUSERDATAJSONOk

`func (o *DBConnector) GetMODIFYUSERDATAJSONOk() (*string, bool)`

GetMODIFYUSERDATAJSONOk returns a tuple with the MODIFYUSERDATAJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMODIFYUSERDATAJSON

`func (o *DBConnector) SetMODIFYUSERDATAJSON(v string)`

SetMODIFYUSERDATAJSON sets MODIFYUSERDATAJSON field to given value.

### HasMODIFYUSERDATAJSON

`func (o *DBConnector) HasMODIFYUSERDATAJSON() bool`

HasMODIFYUSERDATAJSON returns a boolean if a field has been set.

### GetSTATUS_THRESHOLD_CONFIG

`func (o *DBConnector) GetSTATUS_THRESHOLD_CONFIG() string`

GetSTATUS_THRESHOLD_CONFIG returns the STATUS_THRESHOLD_CONFIG field if non-nil, zero value otherwise.

### GetSTATUS_THRESHOLD_CONFIGOk

`func (o *DBConnector) GetSTATUS_THRESHOLD_CONFIGOk() (*string, bool)`

GetSTATUS_THRESHOLD_CONFIGOk returns a tuple with the STATUS_THRESHOLD_CONFIG field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSTATUS_THRESHOLD_CONFIG

`func (o *DBConnector) SetSTATUS_THRESHOLD_CONFIG(v string)`

SetSTATUS_THRESHOLD_CONFIG sets STATUS_THRESHOLD_CONFIG field to given value.

### HasSTATUS_THRESHOLD_CONFIG

`func (o *DBConnector) HasSTATUS_THRESHOLD_CONFIG() bool`

HasSTATUS_THRESHOLD_CONFIG returns a boolean if a field has been set.

### GetMAX_PAGINATION_SIZE

`func (o *DBConnector) GetMAX_PAGINATION_SIZE() string`

GetMAX_PAGINATION_SIZE returns the MAX_PAGINATION_SIZE field if non-nil, zero value otherwise.

### GetMAX_PAGINATION_SIZEOk

`func (o *DBConnector) GetMAX_PAGINATION_SIZEOk() (*string, bool)`

GetMAX_PAGINATION_SIZEOk returns a tuple with the MAX_PAGINATION_SIZE field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMAX_PAGINATION_SIZE

`func (o *DBConnector) SetMAX_PAGINATION_SIZE(v string)`

SetMAX_PAGINATION_SIZE sets MAX_PAGINATION_SIZE field to given value.

### HasMAX_PAGINATION_SIZE

`func (o *DBConnector) HasMAX_PAGINATION_SIZE() bool`

HasMAX_PAGINATION_SIZE returns a boolean if a field has been set.

### GetCLI_COMMAND_JSON

`func (o *DBConnector) GetCLI_COMMAND_JSON() string`

GetCLI_COMMAND_JSON returns the CLI_COMMAND_JSON field if non-nil, zero value otherwise.

### GetCLI_COMMAND_JSONOk

`func (o *DBConnector) GetCLI_COMMAND_JSONOk() (*string, bool)`

GetCLI_COMMAND_JSONOk returns a tuple with the CLI_COMMAND_JSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCLI_COMMAND_JSON

`func (o *DBConnector) SetCLI_COMMAND_JSON(v string)`

SetCLI_COMMAND_JSON sets CLI_COMMAND_JSON field to given value.

### HasCLI_COMMAND_JSON

`func (o *DBConnector) HasCLI_COMMAND_JSON() bool`

HasCLI_COMMAND_JSON returns a boolean if a field has been set.

### GetCREATEENTITLEMENTJSON

`func (o *DBConnector) GetCREATEENTITLEMENTJSON() string`

GetCREATEENTITLEMENTJSON returns the CREATEENTITLEMENTJSON field if non-nil, zero value otherwise.

### GetCREATEENTITLEMENTJSONOk

`func (o *DBConnector) GetCREATEENTITLEMENTJSONOk() (*string, bool)`

GetCREATEENTITLEMENTJSONOk returns a tuple with the CREATEENTITLEMENTJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCREATEENTITLEMENTJSON

`func (o *DBConnector) SetCREATEENTITLEMENTJSON(v string)`

SetCREATEENTITLEMENTJSON sets CREATEENTITLEMENTJSON field to given value.

### HasCREATEENTITLEMENTJSON

`func (o *DBConnector) HasCREATEENTITLEMENTJSON() bool`

HasCREATEENTITLEMENTJSON returns a boolean if a field has been set.

### GetDELETEENTITLEMENTJSON

`func (o *DBConnector) GetDELETEENTITLEMENTJSON() string`

GetDELETEENTITLEMENTJSON returns the DELETEENTITLEMENTJSON field if non-nil, zero value otherwise.

### GetDELETEENTITLEMENTJSONOk

`func (o *DBConnector) GetDELETEENTITLEMENTJSONOk() (*string, bool)`

GetDELETEENTITLEMENTJSONOk returns a tuple with the DELETEENTITLEMENTJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDELETEENTITLEMENTJSON

`func (o *DBConnector) SetDELETEENTITLEMENTJSON(v string)`

SetDELETEENTITLEMENTJSON sets DELETEENTITLEMENTJSON field to given value.

### HasDELETEENTITLEMENTJSON

`func (o *DBConnector) HasDELETEENTITLEMENTJSON() bool`

HasDELETEENTITLEMENTJSON returns a boolean if a field has been set.

### GetENTITLEMENTEXISTJSON

`func (o *DBConnector) GetENTITLEMENTEXISTJSON() string`

GetENTITLEMENTEXISTJSON returns the ENTITLEMENTEXISTJSON field if non-nil, zero value otherwise.

### GetENTITLEMENTEXISTJSONOk

`func (o *DBConnector) GetENTITLEMENTEXISTJSONOk() (*string, bool)`

GetENTITLEMENTEXISTJSONOk returns a tuple with the ENTITLEMENTEXISTJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetENTITLEMENTEXISTJSON

`func (o *DBConnector) SetENTITLEMENTEXISTJSON(v string)`

SetENTITLEMENTEXISTJSON sets ENTITLEMENTEXISTJSON field to given value.

### HasENTITLEMENTEXISTJSON

`func (o *DBConnector) HasENTITLEMENTEXISTJSON() bool`

HasENTITLEMENTEXISTJSON returns a boolean if a field has been set.

### GetUPDATEENTITLEMENTJSON

`func (o *DBConnector) GetUPDATEENTITLEMENTJSON() string`

GetUPDATEENTITLEMENTJSON returns the UPDATEENTITLEMENTJSON field if non-nil, zero value otherwise.

### GetUPDATEENTITLEMENTJSONOk

`func (o *DBConnector) GetUPDATEENTITLEMENTJSONOk() (*string, bool)`

GetUPDATEENTITLEMENTJSONOk returns a tuple with the UPDATEENTITLEMENTJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUPDATEENTITLEMENTJSON

`func (o *DBConnector) SetUPDATEENTITLEMENTJSON(v string)`

SetUPDATEENTITLEMENTJSON sets UPDATEENTITLEMENTJSON field to given value.

### HasUPDATEENTITLEMENTJSON

`func (o *DBConnector) HasUPDATEENTITLEMENTJSON() bool`

HasUPDATEENTITLEMENTJSON returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


