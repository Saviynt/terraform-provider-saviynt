# WorkdaySOAPConnector

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ACCOUNTS_IMPORT_JSON** | Pointer to **string** | JSON configuration for accounts import | [optional] 
**CHANGEPASSJSON** | Pointer to **string** | JSON for password changes | [optional] 
**COMBINEDCREATEREQUEST** | Pointer to **string** | Combined create request configuration | [optional] 
**CONNECTIONJSON** | Pointer to **string** |  | [optional] 
**CREATEACCOUNTJSON** | Pointer to **string** | JSON for account creation | [optional] 
**CUSTOM_CONFIG** | Pointer to **string** | Custom configuration JSON | [optional] 
**DATA_TO_IMPORT** | Pointer to **string** | Specification of data types to import | [optional] 
**DATEFORMAT** | Pointer to **string** | Date format for data processing | [optional] 
**DELETEACCOUNTJSON** | Pointer to **string** | JSON for account deletion | [optional] 
**DISABLEACCOUNTJSON** | Pointer to **string** | JSON for account disabling | [optional] 
**ENABLEACCOUNTJSON** | Pointer to **string** | JSON for account enabling | [optional] 
**GRANTACCESSJSON** | Pointer to **string** | JSON for granting access | [optional] 
**HR_IMPORT_JSON** | Pointer to **string** | JSON configuration for HR data import | [optional] 
**MODIFYUSERDATAJSON** | Pointer to **string** | JSON for modifying user data | [optional] 
**PAGE_SIZE** | Pointer to **string** | Number of records per page | [optional] 
**PAM_CONFIG** | Pointer to **string** | PAM configuration JSON | [optional] 
**PASSWORD** | Pointer to **string** | Password for SOAP authentication | [optional] 
**PASSWORD_MAX_LENGTH** | Pointer to **string** | Maximum password length | [optional] 
**PASSWORD_MIN_LENGTH** | Pointer to **string** | Minimum password length | [optional] 
**PASSWORD_NOOFCAPSALPHA** | Pointer to **string** | Number of capital letters required | [optional] 
**PASSWORD_NOOFDIGITS** | Pointer to **string** | Number of digits required | [optional] 
**PASSWORD_NOOFSPLCHARS** | Pointer to **string** | Number of special characters required | [optional] 
**PASSWORD_TYPE** | Pointer to **string** | Type of password authentication | [optional] 
**RESPONSEPATH_PAGERESULTS** | Pointer to **string** | Response path for page results | [optional] 
**RESPONSEPATH_TOTALRESULTS** | Pointer to **string** | Response path for total results count | [optional] 
**RESPONSEPATH_USERLIST** | Pointer to **string** | Response path for user list | [optional] 
**REVOKEACCESSJSON** | Pointer to **string** | JSON for revoking access | [optional] 
**SOAP_ENDPOINT** | Pointer to **string** | SOAP endpoint URL for Workday | [optional] 
**UPDATEACCOUNTJSON** | Pointer to **string** | JSON for account updates | [optional] 
**UPDATEUSERJSON** | Pointer to **string** | JSON for user updates | [optional] 
**USERNAME** | Pointer to **string** | Username for SOAP authentication | [optional] 

## Methods

### NewWorkdaySOAPConnector

`func NewWorkdaySOAPConnector() *WorkdaySOAPConnector`

NewWorkdaySOAPConnector instantiates a new WorkdaySOAPConnector object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkdaySOAPConnectorWithDefaults

`func NewWorkdaySOAPConnectorWithDefaults() *WorkdaySOAPConnector`

NewWorkdaySOAPConnectorWithDefaults instantiates a new WorkdaySOAPConnector object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetACCOUNTS_IMPORT_JSON

`func (o *WorkdaySOAPConnector) GetACCOUNTS_IMPORT_JSON() string`

GetACCOUNTS_IMPORT_JSON returns the ACCOUNTS_IMPORT_JSON field if non-nil, zero value otherwise.

### GetACCOUNTS_IMPORT_JSONOk

`func (o *WorkdaySOAPConnector) GetACCOUNTS_IMPORT_JSONOk() (*string, bool)`

GetACCOUNTS_IMPORT_JSONOk returns a tuple with the ACCOUNTS_IMPORT_JSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetACCOUNTS_IMPORT_JSON

`func (o *WorkdaySOAPConnector) SetACCOUNTS_IMPORT_JSON(v string)`

SetACCOUNTS_IMPORT_JSON sets ACCOUNTS_IMPORT_JSON field to given value.

### HasACCOUNTS_IMPORT_JSON

`func (o *WorkdaySOAPConnector) HasACCOUNTS_IMPORT_JSON() bool`

HasACCOUNTS_IMPORT_JSON returns a boolean if a field has been set.

### GetCHANGEPASSJSON

`func (o *WorkdaySOAPConnector) GetCHANGEPASSJSON() string`

GetCHANGEPASSJSON returns the CHANGEPASSJSON field if non-nil, zero value otherwise.

### GetCHANGEPASSJSONOk

`func (o *WorkdaySOAPConnector) GetCHANGEPASSJSONOk() (*string, bool)`

GetCHANGEPASSJSONOk returns a tuple with the CHANGEPASSJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCHANGEPASSJSON

`func (o *WorkdaySOAPConnector) SetCHANGEPASSJSON(v string)`

SetCHANGEPASSJSON sets CHANGEPASSJSON field to given value.

### HasCHANGEPASSJSON

`func (o *WorkdaySOAPConnector) HasCHANGEPASSJSON() bool`

HasCHANGEPASSJSON returns a boolean if a field has been set.

### GetCOMBINEDCREATEREQUEST

`func (o *WorkdaySOAPConnector) GetCOMBINEDCREATEREQUEST() string`

GetCOMBINEDCREATEREQUEST returns the COMBINEDCREATEREQUEST field if non-nil, zero value otherwise.

### GetCOMBINEDCREATEREQUESTOk

`func (o *WorkdaySOAPConnector) GetCOMBINEDCREATEREQUESTOk() (*string, bool)`

GetCOMBINEDCREATEREQUESTOk returns a tuple with the COMBINEDCREATEREQUEST field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCOMBINEDCREATEREQUEST

`func (o *WorkdaySOAPConnector) SetCOMBINEDCREATEREQUEST(v string)`

SetCOMBINEDCREATEREQUEST sets COMBINEDCREATEREQUEST field to given value.

### HasCOMBINEDCREATEREQUEST

`func (o *WorkdaySOAPConnector) HasCOMBINEDCREATEREQUEST() bool`

HasCOMBINEDCREATEREQUEST returns a boolean if a field has been set.

### GetCONNECTIONJSON

`func (o *WorkdaySOAPConnector) GetCONNECTIONJSON() string`

GetCONNECTIONJSON returns the CONNECTIONJSON field if non-nil, zero value otherwise.

### GetCONNECTIONJSONOk

`func (o *WorkdaySOAPConnector) GetCONNECTIONJSONOk() (*string, bool)`

GetCONNECTIONJSONOk returns a tuple with the CONNECTIONJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCONNECTIONJSON

`func (o *WorkdaySOAPConnector) SetCONNECTIONJSON(v string)`

SetCONNECTIONJSON sets CONNECTIONJSON field to given value.

### HasCONNECTIONJSON

`func (o *WorkdaySOAPConnector) HasCONNECTIONJSON() bool`

HasCONNECTIONJSON returns a boolean if a field has been set.

### GetCREATEACCOUNTJSON

`func (o *WorkdaySOAPConnector) GetCREATEACCOUNTJSON() string`

GetCREATEACCOUNTJSON returns the CREATEACCOUNTJSON field if non-nil, zero value otherwise.

### GetCREATEACCOUNTJSONOk

`func (o *WorkdaySOAPConnector) GetCREATEACCOUNTJSONOk() (*string, bool)`

GetCREATEACCOUNTJSONOk returns a tuple with the CREATEACCOUNTJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCREATEACCOUNTJSON

`func (o *WorkdaySOAPConnector) SetCREATEACCOUNTJSON(v string)`

SetCREATEACCOUNTJSON sets CREATEACCOUNTJSON field to given value.

### HasCREATEACCOUNTJSON

`func (o *WorkdaySOAPConnector) HasCREATEACCOUNTJSON() bool`

HasCREATEACCOUNTJSON returns a boolean if a field has been set.

### GetCUSTOM_CONFIG

`func (o *WorkdaySOAPConnector) GetCUSTOM_CONFIG() string`

GetCUSTOM_CONFIG returns the CUSTOM_CONFIG field if non-nil, zero value otherwise.

### GetCUSTOM_CONFIGOk

`func (o *WorkdaySOAPConnector) GetCUSTOM_CONFIGOk() (*string, bool)`

GetCUSTOM_CONFIGOk returns a tuple with the CUSTOM_CONFIG field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCUSTOM_CONFIG

`func (o *WorkdaySOAPConnector) SetCUSTOM_CONFIG(v string)`

SetCUSTOM_CONFIG sets CUSTOM_CONFIG field to given value.

### HasCUSTOM_CONFIG

`func (o *WorkdaySOAPConnector) HasCUSTOM_CONFIG() bool`

HasCUSTOM_CONFIG returns a boolean if a field has been set.

### GetDATA_TO_IMPORT

`func (o *WorkdaySOAPConnector) GetDATA_TO_IMPORT() string`

GetDATA_TO_IMPORT returns the DATA_TO_IMPORT field if non-nil, zero value otherwise.

### GetDATA_TO_IMPORTOk

`func (o *WorkdaySOAPConnector) GetDATA_TO_IMPORTOk() (*string, bool)`

GetDATA_TO_IMPORTOk returns a tuple with the DATA_TO_IMPORT field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDATA_TO_IMPORT

`func (o *WorkdaySOAPConnector) SetDATA_TO_IMPORT(v string)`

SetDATA_TO_IMPORT sets DATA_TO_IMPORT field to given value.

### HasDATA_TO_IMPORT

`func (o *WorkdaySOAPConnector) HasDATA_TO_IMPORT() bool`

HasDATA_TO_IMPORT returns a boolean if a field has been set.

### GetDATEFORMAT

`func (o *WorkdaySOAPConnector) GetDATEFORMAT() string`

GetDATEFORMAT returns the DATEFORMAT field if non-nil, zero value otherwise.

### GetDATEFORMATOk

`func (o *WorkdaySOAPConnector) GetDATEFORMATOk() (*string, bool)`

GetDATEFORMATOk returns a tuple with the DATEFORMAT field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDATEFORMAT

`func (o *WorkdaySOAPConnector) SetDATEFORMAT(v string)`

SetDATEFORMAT sets DATEFORMAT field to given value.

### HasDATEFORMAT

`func (o *WorkdaySOAPConnector) HasDATEFORMAT() bool`

HasDATEFORMAT returns a boolean if a field has been set.

### GetDELETEACCOUNTJSON

`func (o *WorkdaySOAPConnector) GetDELETEACCOUNTJSON() string`

GetDELETEACCOUNTJSON returns the DELETEACCOUNTJSON field if non-nil, zero value otherwise.

### GetDELETEACCOUNTJSONOk

`func (o *WorkdaySOAPConnector) GetDELETEACCOUNTJSONOk() (*string, bool)`

GetDELETEACCOUNTJSONOk returns a tuple with the DELETEACCOUNTJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDELETEACCOUNTJSON

`func (o *WorkdaySOAPConnector) SetDELETEACCOUNTJSON(v string)`

SetDELETEACCOUNTJSON sets DELETEACCOUNTJSON field to given value.

### HasDELETEACCOUNTJSON

`func (o *WorkdaySOAPConnector) HasDELETEACCOUNTJSON() bool`

HasDELETEACCOUNTJSON returns a boolean if a field has been set.

### GetDISABLEACCOUNTJSON

`func (o *WorkdaySOAPConnector) GetDISABLEACCOUNTJSON() string`

GetDISABLEACCOUNTJSON returns the DISABLEACCOUNTJSON field if non-nil, zero value otherwise.

### GetDISABLEACCOUNTJSONOk

`func (o *WorkdaySOAPConnector) GetDISABLEACCOUNTJSONOk() (*string, bool)`

GetDISABLEACCOUNTJSONOk returns a tuple with the DISABLEACCOUNTJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDISABLEACCOUNTJSON

`func (o *WorkdaySOAPConnector) SetDISABLEACCOUNTJSON(v string)`

SetDISABLEACCOUNTJSON sets DISABLEACCOUNTJSON field to given value.

### HasDISABLEACCOUNTJSON

`func (o *WorkdaySOAPConnector) HasDISABLEACCOUNTJSON() bool`

HasDISABLEACCOUNTJSON returns a boolean if a field has been set.

### GetENABLEACCOUNTJSON

`func (o *WorkdaySOAPConnector) GetENABLEACCOUNTJSON() string`

GetENABLEACCOUNTJSON returns the ENABLEACCOUNTJSON field if non-nil, zero value otherwise.

### GetENABLEACCOUNTJSONOk

`func (o *WorkdaySOAPConnector) GetENABLEACCOUNTJSONOk() (*string, bool)`

GetENABLEACCOUNTJSONOk returns a tuple with the ENABLEACCOUNTJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetENABLEACCOUNTJSON

`func (o *WorkdaySOAPConnector) SetENABLEACCOUNTJSON(v string)`

SetENABLEACCOUNTJSON sets ENABLEACCOUNTJSON field to given value.

### HasENABLEACCOUNTJSON

`func (o *WorkdaySOAPConnector) HasENABLEACCOUNTJSON() bool`

HasENABLEACCOUNTJSON returns a boolean if a field has been set.

### GetGRANTACCESSJSON

`func (o *WorkdaySOAPConnector) GetGRANTACCESSJSON() string`

GetGRANTACCESSJSON returns the GRANTACCESSJSON field if non-nil, zero value otherwise.

### GetGRANTACCESSJSONOk

`func (o *WorkdaySOAPConnector) GetGRANTACCESSJSONOk() (*string, bool)`

GetGRANTACCESSJSONOk returns a tuple with the GRANTACCESSJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGRANTACCESSJSON

`func (o *WorkdaySOAPConnector) SetGRANTACCESSJSON(v string)`

SetGRANTACCESSJSON sets GRANTACCESSJSON field to given value.

### HasGRANTACCESSJSON

`func (o *WorkdaySOAPConnector) HasGRANTACCESSJSON() bool`

HasGRANTACCESSJSON returns a boolean if a field has been set.

### GetHR_IMPORT_JSON

`func (o *WorkdaySOAPConnector) GetHR_IMPORT_JSON() string`

GetHR_IMPORT_JSON returns the HR_IMPORT_JSON field if non-nil, zero value otherwise.

### GetHR_IMPORT_JSONOk

`func (o *WorkdaySOAPConnector) GetHR_IMPORT_JSONOk() (*string, bool)`

GetHR_IMPORT_JSONOk returns a tuple with the HR_IMPORT_JSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHR_IMPORT_JSON

`func (o *WorkdaySOAPConnector) SetHR_IMPORT_JSON(v string)`

SetHR_IMPORT_JSON sets HR_IMPORT_JSON field to given value.

### HasHR_IMPORT_JSON

`func (o *WorkdaySOAPConnector) HasHR_IMPORT_JSON() bool`

HasHR_IMPORT_JSON returns a boolean if a field has been set.

### GetMODIFYUSERDATAJSON

`func (o *WorkdaySOAPConnector) GetMODIFYUSERDATAJSON() string`

GetMODIFYUSERDATAJSON returns the MODIFYUSERDATAJSON field if non-nil, zero value otherwise.

### GetMODIFYUSERDATAJSONOk

`func (o *WorkdaySOAPConnector) GetMODIFYUSERDATAJSONOk() (*string, bool)`

GetMODIFYUSERDATAJSONOk returns a tuple with the MODIFYUSERDATAJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMODIFYUSERDATAJSON

`func (o *WorkdaySOAPConnector) SetMODIFYUSERDATAJSON(v string)`

SetMODIFYUSERDATAJSON sets MODIFYUSERDATAJSON field to given value.

### HasMODIFYUSERDATAJSON

`func (o *WorkdaySOAPConnector) HasMODIFYUSERDATAJSON() bool`

HasMODIFYUSERDATAJSON returns a boolean if a field has been set.

### GetPAGE_SIZE

`func (o *WorkdaySOAPConnector) GetPAGE_SIZE() string`

GetPAGE_SIZE returns the PAGE_SIZE field if non-nil, zero value otherwise.

### GetPAGE_SIZEOk

`func (o *WorkdaySOAPConnector) GetPAGE_SIZEOk() (*string, bool)`

GetPAGE_SIZEOk returns a tuple with the PAGE_SIZE field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPAGE_SIZE

`func (o *WorkdaySOAPConnector) SetPAGE_SIZE(v string)`

SetPAGE_SIZE sets PAGE_SIZE field to given value.

### HasPAGE_SIZE

`func (o *WorkdaySOAPConnector) HasPAGE_SIZE() bool`

HasPAGE_SIZE returns a boolean if a field has been set.

### GetPAM_CONFIG

`func (o *WorkdaySOAPConnector) GetPAM_CONFIG() string`

GetPAM_CONFIG returns the PAM_CONFIG field if non-nil, zero value otherwise.

### GetPAM_CONFIGOk

`func (o *WorkdaySOAPConnector) GetPAM_CONFIGOk() (*string, bool)`

GetPAM_CONFIGOk returns a tuple with the PAM_CONFIG field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPAM_CONFIG

`func (o *WorkdaySOAPConnector) SetPAM_CONFIG(v string)`

SetPAM_CONFIG sets PAM_CONFIG field to given value.

### HasPAM_CONFIG

`func (o *WorkdaySOAPConnector) HasPAM_CONFIG() bool`

HasPAM_CONFIG returns a boolean if a field has been set.

### GetPASSWORD

`func (o *WorkdaySOAPConnector) GetPASSWORD() string`

GetPASSWORD returns the PASSWORD field if non-nil, zero value otherwise.

### GetPASSWORDOk

`func (o *WorkdaySOAPConnector) GetPASSWORDOk() (*string, bool)`

GetPASSWORDOk returns a tuple with the PASSWORD field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPASSWORD

`func (o *WorkdaySOAPConnector) SetPASSWORD(v string)`

SetPASSWORD sets PASSWORD field to given value.

### HasPASSWORD

`func (o *WorkdaySOAPConnector) HasPASSWORD() bool`

HasPASSWORD returns a boolean if a field has been set.

### GetPASSWORD_MAX_LENGTH

`func (o *WorkdaySOAPConnector) GetPASSWORD_MAX_LENGTH() string`

GetPASSWORD_MAX_LENGTH returns the PASSWORD_MAX_LENGTH field if non-nil, zero value otherwise.

### GetPASSWORD_MAX_LENGTHOk

`func (o *WorkdaySOAPConnector) GetPASSWORD_MAX_LENGTHOk() (*string, bool)`

GetPASSWORD_MAX_LENGTHOk returns a tuple with the PASSWORD_MAX_LENGTH field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPASSWORD_MAX_LENGTH

`func (o *WorkdaySOAPConnector) SetPASSWORD_MAX_LENGTH(v string)`

SetPASSWORD_MAX_LENGTH sets PASSWORD_MAX_LENGTH field to given value.

### HasPASSWORD_MAX_LENGTH

`func (o *WorkdaySOAPConnector) HasPASSWORD_MAX_LENGTH() bool`

HasPASSWORD_MAX_LENGTH returns a boolean if a field has been set.

### GetPASSWORD_MIN_LENGTH

`func (o *WorkdaySOAPConnector) GetPASSWORD_MIN_LENGTH() string`

GetPASSWORD_MIN_LENGTH returns the PASSWORD_MIN_LENGTH field if non-nil, zero value otherwise.

### GetPASSWORD_MIN_LENGTHOk

`func (o *WorkdaySOAPConnector) GetPASSWORD_MIN_LENGTHOk() (*string, bool)`

GetPASSWORD_MIN_LENGTHOk returns a tuple with the PASSWORD_MIN_LENGTH field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPASSWORD_MIN_LENGTH

`func (o *WorkdaySOAPConnector) SetPASSWORD_MIN_LENGTH(v string)`

SetPASSWORD_MIN_LENGTH sets PASSWORD_MIN_LENGTH field to given value.

### HasPASSWORD_MIN_LENGTH

`func (o *WorkdaySOAPConnector) HasPASSWORD_MIN_LENGTH() bool`

HasPASSWORD_MIN_LENGTH returns a boolean if a field has been set.

### GetPASSWORD_NOOFCAPSALPHA

`func (o *WorkdaySOAPConnector) GetPASSWORD_NOOFCAPSALPHA() string`

GetPASSWORD_NOOFCAPSALPHA returns the PASSWORD_NOOFCAPSALPHA field if non-nil, zero value otherwise.

### GetPASSWORD_NOOFCAPSALPHAOk

`func (o *WorkdaySOAPConnector) GetPASSWORD_NOOFCAPSALPHAOk() (*string, bool)`

GetPASSWORD_NOOFCAPSALPHAOk returns a tuple with the PASSWORD_NOOFCAPSALPHA field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPASSWORD_NOOFCAPSALPHA

`func (o *WorkdaySOAPConnector) SetPASSWORD_NOOFCAPSALPHA(v string)`

SetPASSWORD_NOOFCAPSALPHA sets PASSWORD_NOOFCAPSALPHA field to given value.

### HasPASSWORD_NOOFCAPSALPHA

`func (o *WorkdaySOAPConnector) HasPASSWORD_NOOFCAPSALPHA() bool`

HasPASSWORD_NOOFCAPSALPHA returns a boolean if a field has been set.

### GetPASSWORD_NOOFDIGITS

`func (o *WorkdaySOAPConnector) GetPASSWORD_NOOFDIGITS() string`

GetPASSWORD_NOOFDIGITS returns the PASSWORD_NOOFDIGITS field if non-nil, zero value otherwise.

### GetPASSWORD_NOOFDIGITSOk

`func (o *WorkdaySOAPConnector) GetPASSWORD_NOOFDIGITSOk() (*string, bool)`

GetPASSWORD_NOOFDIGITSOk returns a tuple with the PASSWORD_NOOFDIGITS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPASSWORD_NOOFDIGITS

`func (o *WorkdaySOAPConnector) SetPASSWORD_NOOFDIGITS(v string)`

SetPASSWORD_NOOFDIGITS sets PASSWORD_NOOFDIGITS field to given value.

### HasPASSWORD_NOOFDIGITS

`func (o *WorkdaySOAPConnector) HasPASSWORD_NOOFDIGITS() bool`

HasPASSWORD_NOOFDIGITS returns a boolean if a field has been set.

### GetPASSWORD_NOOFSPLCHARS

`func (o *WorkdaySOAPConnector) GetPASSWORD_NOOFSPLCHARS() string`

GetPASSWORD_NOOFSPLCHARS returns the PASSWORD_NOOFSPLCHARS field if non-nil, zero value otherwise.

### GetPASSWORD_NOOFSPLCHARSOk

`func (o *WorkdaySOAPConnector) GetPASSWORD_NOOFSPLCHARSOk() (*string, bool)`

GetPASSWORD_NOOFSPLCHARSOk returns a tuple with the PASSWORD_NOOFSPLCHARS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPASSWORD_NOOFSPLCHARS

`func (o *WorkdaySOAPConnector) SetPASSWORD_NOOFSPLCHARS(v string)`

SetPASSWORD_NOOFSPLCHARS sets PASSWORD_NOOFSPLCHARS field to given value.

### HasPASSWORD_NOOFSPLCHARS

`func (o *WorkdaySOAPConnector) HasPASSWORD_NOOFSPLCHARS() bool`

HasPASSWORD_NOOFSPLCHARS returns a boolean if a field has been set.

### GetPASSWORD_TYPE

`func (o *WorkdaySOAPConnector) GetPASSWORD_TYPE() string`

GetPASSWORD_TYPE returns the PASSWORD_TYPE field if non-nil, zero value otherwise.

### GetPASSWORD_TYPEOk

`func (o *WorkdaySOAPConnector) GetPASSWORD_TYPEOk() (*string, bool)`

GetPASSWORD_TYPEOk returns a tuple with the PASSWORD_TYPE field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPASSWORD_TYPE

`func (o *WorkdaySOAPConnector) SetPASSWORD_TYPE(v string)`

SetPASSWORD_TYPE sets PASSWORD_TYPE field to given value.

### HasPASSWORD_TYPE

`func (o *WorkdaySOAPConnector) HasPASSWORD_TYPE() bool`

HasPASSWORD_TYPE returns a boolean if a field has been set.

### GetRESPONSEPATH_PAGERESULTS

`func (o *WorkdaySOAPConnector) GetRESPONSEPATH_PAGERESULTS() string`

GetRESPONSEPATH_PAGERESULTS returns the RESPONSEPATH_PAGERESULTS field if non-nil, zero value otherwise.

### GetRESPONSEPATH_PAGERESULTSOk

`func (o *WorkdaySOAPConnector) GetRESPONSEPATH_PAGERESULTSOk() (*string, bool)`

GetRESPONSEPATH_PAGERESULTSOk returns a tuple with the RESPONSEPATH_PAGERESULTS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRESPONSEPATH_PAGERESULTS

`func (o *WorkdaySOAPConnector) SetRESPONSEPATH_PAGERESULTS(v string)`

SetRESPONSEPATH_PAGERESULTS sets RESPONSEPATH_PAGERESULTS field to given value.

### HasRESPONSEPATH_PAGERESULTS

`func (o *WorkdaySOAPConnector) HasRESPONSEPATH_PAGERESULTS() bool`

HasRESPONSEPATH_PAGERESULTS returns a boolean if a field has been set.

### GetRESPONSEPATH_TOTALRESULTS

`func (o *WorkdaySOAPConnector) GetRESPONSEPATH_TOTALRESULTS() string`

GetRESPONSEPATH_TOTALRESULTS returns the RESPONSEPATH_TOTALRESULTS field if non-nil, zero value otherwise.

### GetRESPONSEPATH_TOTALRESULTSOk

`func (o *WorkdaySOAPConnector) GetRESPONSEPATH_TOTALRESULTSOk() (*string, bool)`

GetRESPONSEPATH_TOTALRESULTSOk returns a tuple with the RESPONSEPATH_TOTALRESULTS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRESPONSEPATH_TOTALRESULTS

`func (o *WorkdaySOAPConnector) SetRESPONSEPATH_TOTALRESULTS(v string)`

SetRESPONSEPATH_TOTALRESULTS sets RESPONSEPATH_TOTALRESULTS field to given value.

### HasRESPONSEPATH_TOTALRESULTS

`func (o *WorkdaySOAPConnector) HasRESPONSEPATH_TOTALRESULTS() bool`

HasRESPONSEPATH_TOTALRESULTS returns a boolean if a field has been set.

### GetRESPONSEPATH_USERLIST

`func (o *WorkdaySOAPConnector) GetRESPONSEPATH_USERLIST() string`

GetRESPONSEPATH_USERLIST returns the RESPONSEPATH_USERLIST field if non-nil, zero value otherwise.

### GetRESPONSEPATH_USERLISTOk

`func (o *WorkdaySOAPConnector) GetRESPONSEPATH_USERLISTOk() (*string, bool)`

GetRESPONSEPATH_USERLISTOk returns a tuple with the RESPONSEPATH_USERLIST field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRESPONSEPATH_USERLIST

`func (o *WorkdaySOAPConnector) SetRESPONSEPATH_USERLIST(v string)`

SetRESPONSEPATH_USERLIST sets RESPONSEPATH_USERLIST field to given value.

### HasRESPONSEPATH_USERLIST

`func (o *WorkdaySOAPConnector) HasRESPONSEPATH_USERLIST() bool`

HasRESPONSEPATH_USERLIST returns a boolean if a field has been set.

### GetREVOKEACCESSJSON

`func (o *WorkdaySOAPConnector) GetREVOKEACCESSJSON() string`

GetREVOKEACCESSJSON returns the REVOKEACCESSJSON field if non-nil, zero value otherwise.

### GetREVOKEACCESSJSONOk

`func (o *WorkdaySOAPConnector) GetREVOKEACCESSJSONOk() (*string, bool)`

GetREVOKEACCESSJSONOk returns a tuple with the REVOKEACCESSJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetREVOKEACCESSJSON

`func (o *WorkdaySOAPConnector) SetREVOKEACCESSJSON(v string)`

SetREVOKEACCESSJSON sets REVOKEACCESSJSON field to given value.

### HasREVOKEACCESSJSON

`func (o *WorkdaySOAPConnector) HasREVOKEACCESSJSON() bool`

HasREVOKEACCESSJSON returns a boolean if a field has been set.

### GetSOAP_ENDPOINT

`func (o *WorkdaySOAPConnector) GetSOAP_ENDPOINT() string`

GetSOAP_ENDPOINT returns the SOAP_ENDPOINT field if non-nil, zero value otherwise.

### GetSOAP_ENDPOINTOk

`func (o *WorkdaySOAPConnector) GetSOAP_ENDPOINTOk() (*string, bool)`

GetSOAP_ENDPOINTOk returns a tuple with the SOAP_ENDPOINT field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSOAP_ENDPOINT

`func (o *WorkdaySOAPConnector) SetSOAP_ENDPOINT(v string)`

SetSOAP_ENDPOINT sets SOAP_ENDPOINT field to given value.

### HasSOAP_ENDPOINT

`func (o *WorkdaySOAPConnector) HasSOAP_ENDPOINT() bool`

HasSOAP_ENDPOINT returns a boolean if a field has been set.

### GetUPDATEACCOUNTJSON

`func (o *WorkdaySOAPConnector) GetUPDATEACCOUNTJSON() string`

GetUPDATEACCOUNTJSON returns the UPDATEACCOUNTJSON field if non-nil, zero value otherwise.

### GetUPDATEACCOUNTJSONOk

`func (o *WorkdaySOAPConnector) GetUPDATEACCOUNTJSONOk() (*string, bool)`

GetUPDATEACCOUNTJSONOk returns a tuple with the UPDATEACCOUNTJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUPDATEACCOUNTJSON

`func (o *WorkdaySOAPConnector) SetUPDATEACCOUNTJSON(v string)`

SetUPDATEACCOUNTJSON sets UPDATEACCOUNTJSON field to given value.

### HasUPDATEACCOUNTJSON

`func (o *WorkdaySOAPConnector) HasUPDATEACCOUNTJSON() bool`

HasUPDATEACCOUNTJSON returns a boolean if a field has been set.

### GetUPDATEUSERJSON

`func (o *WorkdaySOAPConnector) GetUPDATEUSERJSON() string`

GetUPDATEUSERJSON returns the UPDATEUSERJSON field if non-nil, zero value otherwise.

### GetUPDATEUSERJSONOk

`func (o *WorkdaySOAPConnector) GetUPDATEUSERJSONOk() (*string, bool)`

GetUPDATEUSERJSONOk returns a tuple with the UPDATEUSERJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUPDATEUSERJSON

`func (o *WorkdaySOAPConnector) SetUPDATEUSERJSON(v string)`

SetUPDATEUSERJSON sets UPDATEUSERJSON field to given value.

### HasUPDATEUSERJSON

`func (o *WorkdaySOAPConnector) HasUPDATEUSERJSON() bool`

HasUPDATEUSERJSON returns a boolean if a field has been set.

### GetUSERNAME

`func (o *WorkdaySOAPConnector) GetUSERNAME() string`

GetUSERNAME returns the USERNAME field if non-nil, zero value otherwise.

### GetUSERNAMEOk

`func (o *WorkdaySOAPConnector) GetUSERNAMEOk() (*string, bool)`

GetUSERNAMEOk returns a tuple with the USERNAME field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUSERNAME

`func (o *WorkdaySOAPConnector) SetUSERNAME(v string)`

SetUSERNAME sets USERNAME field to given value.

### HasUSERNAME

`func (o *WorkdaySOAPConnector) HasUSERNAME() bool`

HasUSERNAME returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


