# WorkdaySOAPConnectionAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ACCOUNTS_IMPORT_JSON** | Pointer to **string** |  | [optional] 
**CHANGEPASSJSON** | Pointer to **string** |  | [optional] 
**COMBINEDCREATEREQUEST** | Pointer to **string** |  | [optional] 
**CONNECTIONJSON** | Pointer to **string** |  | [optional] 
**CREATEACCOUNTJSON** | Pointer to **string** |  | [optional] 
**CUSTOM_CONFIG** | Pointer to **string** |  | [optional] 
**DATA_TO_IMPORT** | Pointer to **string** |  | [optional] 
**DATEFORMAT** | Pointer to **string** |  | [optional] 
**DELETEACCOUNTJSON** | Pointer to **string** |  | [optional] 
**DISABLEACCOUNTJSON** | Pointer to **string** |  | [optional] 
**ENABLEACCOUNTJSON** | Pointer to **string** |  | [optional] 
**GRANTACCESSJSON** | Pointer to **string** |  | [optional] 
**HR_IMPORT_JSON** | Pointer to **string** |  | [optional] 
**MODIFYUSERDATAJSON** | Pointer to **string** |  | [optional] 
**PAGE_SIZE** | Pointer to **string** |  | [optional] 
**PAM_CONFIG** | Pointer to **string** |  | [optional] 
**PASSWORD** | Pointer to **string** |  | [optional] 
**PASSWORD_MAX_LENGTH** | Pointer to **string** |  | [optional] 
**PASSWORD_MIN_LENGTH** | Pointer to **string** |  | [optional] 
**PASSWORD_NOOFCAPSALPHA** | Pointer to **string** |  | [optional] 
**PASSWORD_NOOFDIGITS** | Pointer to **string** |  | [optional] 
**PASSWORD_NOOFSPLCHARS** | Pointer to **string** |  | [optional] 
**PASSWORD_TYPE** | Pointer to **string** |  | [optional] 
**RESPONSEPATH_PAGERESULTS** | Pointer to **string** |  | [optional] 
**RESPONSEPATH_TOTALRESULTS** | Pointer to **string** |  | [optional] 
**RESPONSEPATH_USERLIST** | Pointer to **string** |  | [optional] 
**REVOKEACCESSJSON** | Pointer to **string** |  | [optional] 
**SOAP_ENDPOINT** | Pointer to **string** |  | [optional] 
**UPDATEACCOUNTJSON** | Pointer to **string** |  | [optional] 
**UPDATEUSERJSON** | Pointer to **string** |  | [optional] 
**USERNAME** | Pointer to **string** |  | [optional] 
**ConnectionTimeoutConfig** | Pointer to [**ConnectionTimeoutConfig**](ConnectionTimeoutConfig.md) |  | [optional] 
**ConnectionType** | Pointer to **string** |  | [optional] 
**IsTimeoutConfigValidated** | Pointer to **bool** |  | [optional] 
**IsTimeoutSupported** | Pointer to **bool** |  | [optional] 

## Methods

### NewWorkdaySOAPConnectionAttributes

`func NewWorkdaySOAPConnectionAttributes() *WorkdaySOAPConnectionAttributes`

NewWorkdaySOAPConnectionAttributes instantiates a new WorkdaySOAPConnectionAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkdaySOAPConnectionAttributesWithDefaults

`func NewWorkdaySOAPConnectionAttributesWithDefaults() *WorkdaySOAPConnectionAttributes`

NewWorkdaySOAPConnectionAttributesWithDefaults instantiates a new WorkdaySOAPConnectionAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetACCOUNTS_IMPORT_JSON

`func (o *WorkdaySOAPConnectionAttributes) GetACCOUNTS_IMPORT_JSON() string`

GetACCOUNTS_IMPORT_JSON returns the ACCOUNTS_IMPORT_JSON field if non-nil, zero value otherwise.

### GetACCOUNTS_IMPORT_JSONOk

`func (o *WorkdaySOAPConnectionAttributes) GetACCOUNTS_IMPORT_JSONOk() (*string, bool)`

GetACCOUNTS_IMPORT_JSONOk returns a tuple with the ACCOUNTS_IMPORT_JSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetACCOUNTS_IMPORT_JSON

`func (o *WorkdaySOAPConnectionAttributes) SetACCOUNTS_IMPORT_JSON(v string)`

SetACCOUNTS_IMPORT_JSON sets ACCOUNTS_IMPORT_JSON field to given value.

### HasACCOUNTS_IMPORT_JSON

`func (o *WorkdaySOAPConnectionAttributes) HasACCOUNTS_IMPORT_JSON() bool`

HasACCOUNTS_IMPORT_JSON returns a boolean if a field has been set.

### GetCHANGEPASSJSON

`func (o *WorkdaySOAPConnectionAttributes) GetCHANGEPASSJSON() string`

GetCHANGEPASSJSON returns the CHANGEPASSJSON field if non-nil, zero value otherwise.

### GetCHANGEPASSJSONOk

`func (o *WorkdaySOAPConnectionAttributes) GetCHANGEPASSJSONOk() (*string, bool)`

GetCHANGEPASSJSONOk returns a tuple with the CHANGEPASSJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCHANGEPASSJSON

`func (o *WorkdaySOAPConnectionAttributes) SetCHANGEPASSJSON(v string)`

SetCHANGEPASSJSON sets CHANGEPASSJSON field to given value.

### HasCHANGEPASSJSON

`func (o *WorkdaySOAPConnectionAttributes) HasCHANGEPASSJSON() bool`

HasCHANGEPASSJSON returns a boolean if a field has been set.

### GetCOMBINEDCREATEREQUEST

`func (o *WorkdaySOAPConnectionAttributes) GetCOMBINEDCREATEREQUEST() string`

GetCOMBINEDCREATEREQUEST returns the COMBINEDCREATEREQUEST field if non-nil, zero value otherwise.

### GetCOMBINEDCREATEREQUESTOk

`func (o *WorkdaySOAPConnectionAttributes) GetCOMBINEDCREATEREQUESTOk() (*string, bool)`

GetCOMBINEDCREATEREQUESTOk returns a tuple with the COMBINEDCREATEREQUEST field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCOMBINEDCREATEREQUEST

`func (o *WorkdaySOAPConnectionAttributes) SetCOMBINEDCREATEREQUEST(v string)`

SetCOMBINEDCREATEREQUEST sets COMBINEDCREATEREQUEST field to given value.

### HasCOMBINEDCREATEREQUEST

`func (o *WorkdaySOAPConnectionAttributes) HasCOMBINEDCREATEREQUEST() bool`

HasCOMBINEDCREATEREQUEST returns a boolean if a field has been set.

### GetCONNECTIONJSON

`func (o *WorkdaySOAPConnectionAttributes) GetCONNECTIONJSON() string`

GetCONNECTIONJSON returns the CONNECTIONJSON field if non-nil, zero value otherwise.

### GetCONNECTIONJSONOk

`func (o *WorkdaySOAPConnectionAttributes) GetCONNECTIONJSONOk() (*string, bool)`

GetCONNECTIONJSONOk returns a tuple with the CONNECTIONJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCONNECTIONJSON

`func (o *WorkdaySOAPConnectionAttributes) SetCONNECTIONJSON(v string)`

SetCONNECTIONJSON sets CONNECTIONJSON field to given value.

### HasCONNECTIONJSON

`func (o *WorkdaySOAPConnectionAttributes) HasCONNECTIONJSON() bool`

HasCONNECTIONJSON returns a boolean if a field has been set.

### GetCREATEACCOUNTJSON

`func (o *WorkdaySOAPConnectionAttributes) GetCREATEACCOUNTJSON() string`

GetCREATEACCOUNTJSON returns the CREATEACCOUNTJSON field if non-nil, zero value otherwise.

### GetCREATEACCOUNTJSONOk

`func (o *WorkdaySOAPConnectionAttributes) GetCREATEACCOUNTJSONOk() (*string, bool)`

GetCREATEACCOUNTJSONOk returns a tuple with the CREATEACCOUNTJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCREATEACCOUNTJSON

`func (o *WorkdaySOAPConnectionAttributes) SetCREATEACCOUNTJSON(v string)`

SetCREATEACCOUNTJSON sets CREATEACCOUNTJSON field to given value.

### HasCREATEACCOUNTJSON

`func (o *WorkdaySOAPConnectionAttributes) HasCREATEACCOUNTJSON() bool`

HasCREATEACCOUNTJSON returns a boolean if a field has been set.

### GetCUSTOM_CONFIG

`func (o *WorkdaySOAPConnectionAttributes) GetCUSTOM_CONFIG() string`

GetCUSTOM_CONFIG returns the CUSTOM_CONFIG field if non-nil, zero value otherwise.

### GetCUSTOM_CONFIGOk

`func (o *WorkdaySOAPConnectionAttributes) GetCUSTOM_CONFIGOk() (*string, bool)`

GetCUSTOM_CONFIGOk returns a tuple with the CUSTOM_CONFIG field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCUSTOM_CONFIG

`func (o *WorkdaySOAPConnectionAttributes) SetCUSTOM_CONFIG(v string)`

SetCUSTOM_CONFIG sets CUSTOM_CONFIG field to given value.

### HasCUSTOM_CONFIG

`func (o *WorkdaySOAPConnectionAttributes) HasCUSTOM_CONFIG() bool`

HasCUSTOM_CONFIG returns a boolean if a field has been set.

### GetDATA_TO_IMPORT

`func (o *WorkdaySOAPConnectionAttributes) GetDATA_TO_IMPORT() string`

GetDATA_TO_IMPORT returns the DATA_TO_IMPORT field if non-nil, zero value otherwise.

### GetDATA_TO_IMPORTOk

`func (o *WorkdaySOAPConnectionAttributes) GetDATA_TO_IMPORTOk() (*string, bool)`

GetDATA_TO_IMPORTOk returns a tuple with the DATA_TO_IMPORT field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDATA_TO_IMPORT

`func (o *WorkdaySOAPConnectionAttributes) SetDATA_TO_IMPORT(v string)`

SetDATA_TO_IMPORT sets DATA_TO_IMPORT field to given value.

### HasDATA_TO_IMPORT

`func (o *WorkdaySOAPConnectionAttributes) HasDATA_TO_IMPORT() bool`

HasDATA_TO_IMPORT returns a boolean if a field has been set.

### GetDATEFORMAT

`func (o *WorkdaySOAPConnectionAttributes) GetDATEFORMAT() string`

GetDATEFORMAT returns the DATEFORMAT field if non-nil, zero value otherwise.

### GetDATEFORMATOk

`func (o *WorkdaySOAPConnectionAttributes) GetDATEFORMATOk() (*string, bool)`

GetDATEFORMATOk returns a tuple with the DATEFORMAT field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDATEFORMAT

`func (o *WorkdaySOAPConnectionAttributes) SetDATEFORMAT(v string)`

SetDATEFORMAT sets DATEFORMAT field to given value.

### HasDATEFORMAT

`func (o *WorkdaySOAPConnectionAttributes) HasDATEFORMAT() bool`

HasDATEFORMAT returns a boolean if a field has been set.

### GetDELETEACCOUNTJSON

`func (o *WorkdaySOAPConnectionAttributes) GetDELETEACCOUNTJSON() string`

GetDELETEACCOUNTJSON returns the DELETEACCOUNTJSON field if non-nil, zero value otherwise.

### GetDELETEACCOUNTJSONOk

`func (o *WorkdaySOAPConnectionAttributes) GetDELETEACCOUNTJSONOk() (*string, bool)`

GetDELETEACCOUNTJSONOk returns a tuple with the DELETEACCOUNTJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDELETEACCOUNTJSON

`func (o *WorkdaySOAPConnectionAttributes) SetDELETEACCOUNTJSON(v string)`

SetDELETEACCOUNTJSON sets DELETEACCOUNTJSON field to given value.

### HasDELETEACCOUNTJSON

`func (o *WorkdaySOAPConnectionAttributes) HasDELETEACCOUNTJSON() bool`

HasDELETEACCOUNTJSON returns a boolean if a field has been set.

### GetDISABLEACCOUNTJSON

`func (o *WorkdaySOAPConnectionAttributes) GetDISABLEACCOUNTJSON() string`

GetDISABLEACCOUNTJSON returns the DISABLEACCOUNTJSON field if non-nil, zero value otherwise.

### GetDISABLEACCOUNTJSONOk

`func (o *WorkdaySOAPConnectionAttributes) GetDISABLEACCOUNTJSONOk() (*string, bool)`

GetDISABLEACCOUNTJSONOk returns a tuple with the DISABLEACCOUNTJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDISABLEACCOUNTJSON

`func (o *WorkdaySOAPConnectionAttributes) SetDISABLEACCOUNTJSON(v string)`

SetDISABLEACCOUNTJSON sets DISABLEACCOUNTJSON field to given value.

### HasDISABLEACCOUNTJSON

`func (o *WorkdaySOAPConnectionAttributes) HasDISABLEACCOUNTJSON() bool`

HasDISABLEACCOUNTJSON returns a boolean if a field has been set.

### GetENABLEACCOUNTJSON

`func (o *WorkdaySOAPConnectionAttributes) GetENABLEACCOUNTJSON() string`

GetENABLEACCOUNTJSON returns the ENABLEACCOUNTJSON field if non-nil, zero value otherwise.

### GetENABLEACCOUNTJSONOk

`func (o *WorkdaySOAPConnectionAttributes) GetENABLEACCOUNTJSONOk() (*string, bool)`

GetENABLEACCOUNTJSONOk returns a tuple with the ENABLEACCOUNTJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetENABLEACCOUNTJSON

`func (o *WorkdaySOAPConnectionAttributes) SetENABLEACCOUNTJSON(v string)`

SetENABLEACCOUNTJSON sets ENABLEACCOUNTJSON field to given value.

### HasENABLEACCOUNTJSON

`func (o *WorkdaySOAPConnectionAttributes) HasENABLEACCOUNTJSON() bool`

HasENABLEACCOUNTJSON returns a boolean if a field has been set.

### GetGRANTACCESSJSON

`func (o *WorkdaySOAPConnectionAttributes) GetGRANTACCESSJSON() string`

GetGRANTACCESSJSON returns the GRANTACCESSJSON field if non-nil, zero value otherwise.

### GetGRANTACCESSJSONOk

`func (o *WorkdaySOAPConnectionAttributes) GetGRANTACCESSJSONOk() (*string, bool)`

GetGRANTACCESSJSONOk returns a tuple with the GRANTACCESSJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGRANTACCESSJSON

`func (o *WorkdaySOAPConnectionAttributes) SetGRANTACCESSJSON(v string)`

SetGRANTACCESSJSON sets GRANTACCESSJSON field to given value.

### HasGRANTACCESSJSON

`func (o *WorkdaySOAPConnectionAttributes) HasGRANTACCESSJSON() bool`

HasGRANTACCESSJSON returns a boolean if a field has been set.

### GetHR_IMPORT_JSON

`func (o *WorkdaySOAPConnectionAttributes) GetHR_IMPORT_JSON() string`

GetHR_IMPORT_JSON returns the HR_IMPORT_JSON field if non-nil, zero value otherwise.

### GetHR_IMPORT_JSONOk

`func (o *WorkdaySOAPConnectionAttributes) GetHR_IMPORT_JSONOk() (*string, bool)`

GetHR_IMPORT_JSONOk returns a tuple with the HR_IMPORT_JSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHR_IMPORT_JSON

`func (o *WorkdaySOAPConnectionAttributes) SetHR_IMPORT_JSON(v string)`

SetHR_IMPORT_JSON sets HR_IMPORT_JSON field to given value.

### HasHR_IMPORT_JSON

`func (o *WorkdaySOAPConnectionAttributes) HasHR_IMPORT_JSON() bool`

HasHR_IMPORT_JSON returns a boolean if a field has been set.

### GetMODIFYUSERDATAJSON

`func (o *WorkdaySOAPConnectionAttributes) GetMODIFYUSERDATAJSON() string`

GetMODIFYUSERDATAJSON returns the MODIFYUSERDATAJSON field if non-nil, zero value otherwise.

### GetMODIFYUSERDATAJSONOk

`func (o *WorkdaySOAPConnectionAttributes) GetMODIFYUSERDATAJSONOk() (*string, bool)`

GetMODIFYUSERDATAJSONOk returns a tuple with the MODIFYUSERDATAJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMODIFYUSERDATAJSON

`func (o *WorkdaySOAPConnectionAttributes) SetMODIFYUSERDATAJSON(v string)`

SetMODIFYUSERDATAJSON sets MODIFYUSERDATAJSON field to given value.

### HasMODIFYUSERDATAJSON

`func (o *WorkdaySOAPConnectionAttributes) HasMODIFYUSERDATAJSON() bool`

HasMODIFYUSERDATAJSON returns a boolean if a field has been set.

### GetPAGE_SIZE

`func (o *WorkdaySOAPConnectionAttributes) GetPAGE_SIZE() string`

GetPAGE_SIZE returns the PAGE_SIZE field if non-nil, zero value otherwise.

### GetPAGE_SIZEOk

`func (o *WorkdaySOAPConnectionAttributes) GetPAGE_SIZEOk() (*string, bool)`

GetPAGE_SIZEOk returns a tuple with the PAGE_SIZE field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPAGE_SIZE

`func (o *WorkdaySOAPConnectionAttributes) SetPAGE_SIZE(v string)`

SetPAGE_SIZE sets PAGE_SIZE field to given value.

### HasPAGE_SIZE

`func (o *WorkdaySOAPConnectionAttributes) HasPAGE_SIZE() bool`

HasPAGE_SIZE returns a boolean if a field has been set.

### GetPAM_CONFIG

`func (o *WorkdaySOAPConnectionAttributes) GetPAM_CONFIG() string`

GetPAM_CONFIG returns the PAM_CONFIG field if non-nil, zero value otherwise.

### GetPAM_CONFIGOk

`func (o *WorkdaySOAPConnectionAttributes) GetPAM_CONFIGOk() (*string, bool)`

GetPAM_CONFIGOk returns a tuple with the PAM_CONFIG field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPAM_CONFIG

`func (o *WorkdaySOAPConnectionAttributes) SetPAM_CONFIG(v string)`

SetPAM_CONFIG sets PAM_CONFIG field to given value.

### HasPAM_CONFIG

`func (o *WorkdaySOAPConnectionAttributes) HasPAM_CONFIG() bool`

HasPAM_CONFIG returns a boolean if a field has been set.

### GetPASSWORD

`func (o *WorkdaySOAPConnectionAttributes) GetPASSWORD() string`

GetPASSWORD returns the PASSWORD field if non-nil, zero value otherwise.

### GetPASSWORDOk

`func (o *WorkdaySOAPConnectionAttributes) GetPASSWORDOk() (*string, bool)`

GetPASSWORDOk returns a tuple with the PASSWORD field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPASSWORD

`func (o *WorkdaySOAPConnectionAttributes) SetPASSWORD(v string)`

SetPASSWORD sets PASSWORD field to given value.

### HasPASSWORD

`func (o *WorkdaySOAPConnectionAttributes) HasPASSWORD() bool`

HasPASSWORD returns a boolean if a field has been set.

### GetPASSWORD_MAX_LENGTH

`func (o *WorkdaySOAPConnectionAttributes) GetPASSWORD_MAX_LENGTH() string`

GetPASSWORD_MAX_LENGTH returns the PASSWORD_MAX_LENGTH field if non-nil, zero value otherwise.

### GetPASSWORD_MAX_LENGTHOk

`func (o *WorkdaySOAPConnectionAttributes) GetPASSWORD_MAX_LENGTHOk() (*string, bool)`

GetPASSWORD_MAX_LENGTHOk returns a tuple with the PASSWORD_MAX_LENGTH field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPASSWORD_MAX_LENGTH

`func (o *WorkdaySOAPConnectionAttributes) SetPASSWORD_MAX_LENGTH(v string)`

SetPASSWORD_MAX_LENGTH sets PASSWORD_MAX_LENGTH field to given value.

### HasPASSWORD_MAX_LENGTH

`func (o *WorkdaySOAPConnectionAttributes) HasPASSWORD_MAX_LENGTH() bool`

HasPASSWORD_MAX_LENGTH returns a boolean if a field has been set.

### GetPASSWORD_MIN_LENGTH

`func (o *WorkdaySOAPConnectionAttributes) GetPASSWORD_MIN_LENGTH() string`

GetPASSWORD_MIN_LENGTH returns the PASSWORD_MIN_LENGTH field if non-nil, zero value otherwise.

### GetPASSWORD_MIN_LENGTHOk

`func (o *WorkdaySOAPConnectionAttributes) GetPASSWORD_MIN_LENGTHOk() (*string, bool)`

GetPASSWORD_MIN_LENGTHOk returns a tuple with the PASSWORD_MIN_LENGTH field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPASSWORD_MIN_LENGTH

`func (o *WorkdaySOAPConnectionAttributes) SetPASSWORD_MIN_LENGTH(v string)`

SetPASSWORD_MIN_LENGTH sets PASSWORD_MIN_LENGTH field to given value.

### HasPASSWORD_MIN_LENGTH

`func (o *WorkdaySOAPConnectionAttributes) HasPASSWORD_MIN_LENGTH() bool`

HasPASSWORD_MIN_LENGTH returns a boolean if a field has been set.

### GetPASSWORD_NOOFCAPSALPHA

`func (o *WorkdaySOAPConnectionAttributes) GetPASSWORD_NOOFCAPSALPHA() string`

GetPASSWORD_NOOFCAPSALPHA returns the PASSWORD_NOOFCAPSALPHA field if non-nil, zero value otherwise.

### GetPASSWORD_NOOFCAPSALPHAOk

`func (o *WorkdaySOAPConnectionAttributes) GetPASSWORD_NOOFCAPSALPHAOk() (*string, bool)`

GetPASSWORD_NOOFCAPSALPHAOk returns a tuple with the PASSWORD_NOOFCAPSALPHA field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPASSWORD_NOOFCAPSALPHA

`func (o *WorkdaySOAPConnectionAttributes) SetPASSWORD_NOOFCAPSALPHA(v string)`

SetPASSWORD_NOOFCAPSALPHA sets PASSWORD_NOOFCAPSALPHA field to given value.

### HasPASSWORD_NOOFCAPSALPHA

`func (o *WorkdaySOAPConnectionAttributes) HasPASSWORD_NOOFCAPSALPHA() bool`

HasPASSWORD_NOOFCAPSALPHA returns a boolean if a field has been set.

### GetPASSWORD_NOOFDIGITS

`func (o *WorkdaySOAPConnectionAttributes) GetPASSWORD_NOOFDIGITS() string`

GetPASSWORD_NOOFDIGITS returns the PASSWORD_NOOFDIGITS field if non-nil, zero value otherwise.

### GetPASSWORD_NOOFDIGITSOk

`func (o *WorkdaySOAPConnectionAttributes) GetPASSWORD_NOOFDIGITSOk() (*string, bool)`

GetPASSWORD_NOOFDIGITSOk returns a tuple with the PASSWORD_NOOFDIGITS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPASSWORD_NOOFDIGITS

`func (o *WorkdaySOAPConnectionAttributes) SetPASSWORD_NOOFDIGITS(v string)`

SetPASSWORD_NOOFDIGITS sets PASSWORD_NOOFDIGITS field to given value.

### HasPASSWORD_NOOFDIGITS

`func (o *WorkdaySOAPConnectionAttributes) HasPASSWORD_NOOFDIGITS() bool`

HasPASSWORD_NOOFDIGITS returns a boolean if a field has been set.

### GetPASSWORD_NOOFSPLCHARS

`func (o *WorkdaySOAPConnectionAttributes) GetPASSWORD_NOOFSPLCHARS() string`

GetPASSWORD_NOOFSPLCHARS returns the PASSWORD_NOOFSPLCHARS field if non-nil, zero value otherwise.

### GetPASSWORD_NOOFSPLCHARSOk

`func (o *WorkdaySOAPConnectionAttributes) GetPASSWORD_NOOFSPLCHARSOk() (*string, bool)`

GetPASSWORD_NOOFSPLCHARSOk returns a tuple with the PASSWORD_NOOFSPLCHARS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPASSWORD_NOOFSPLCHARS

`func (o *WorkdaySOAPConnectionAttributes) SetPASSWORD_NOOFSPLCHARS(v string)`

SetPASSWORD_NOOFSPLCHARS sets PASSWORD_NOOFSPLCHARS field to given value.

### HasPASSWORD_NOOFSPLCHARS

`func (o *WorkdaySOAPConnectionAttributes) HasPASSWORD_NOOFSPLCHARS() bool`

HasPASSWORD_NOOFSPLCHARS returns a boolean if a field has been set.

### GetPASSWORD_TYPE

`func (o *WorkdaySOAPConnectionAttributes) GetPASSWORD_TYPE() string`

GetPASSWORD_TYPE returns the PASSWORD_TYPE field if non-nil, zero value otherwise.

### GetPASSWORD_TYPEOk

`func (o *WorkdaySOAPConnectionAttributes) GetPASSWORD_TYPEOk() (*string, bool)`

GetPASSWORD_TYPEOk returns a tuple with the PASSWORD_TYPE field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPASSWORD_TYPE

`func (o *WorkdaySOAPConnectionAttributes) SetPASSWORD_TYPE(v string)`

SetPASSWORD_TYPE sets PASSWORD_TYPE field to given value.

### HasPASSWORD_TYPE

`func (o *WorkdaySOAPConnectionAttributes) HasPASSWORD_TYPE() bool`

HasPASSWORD_TYPE returns a boolean if a field has been set.

### GetRESPONSEPATH_PAGERESULTS

`func (o *WorkdaySOAPConnectionAttributes) GetRESPONSEPATH_PAGERESULTS() string`

GetRESPONSEPATH_PAGERESULTS returns the RESPONSEPATH_PAGERESULTS field if non-nil, zero value otherwise.

### GetRESPONSEPATH_PAGERESULTSOk

`func (o *WorkdaySOAPConnectionAttributes) GetRESPONSEPATH_PAGERESULTSOk() (*string, bool)`

GetRESPONSEPATH_PAGERESULTSOk returns a tuple with the RESPONSEPATH_PAGERESULTS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRESPONSEPATH_PAGERESULTS

`func (o *WorkdaySOAPConnectionAttributes) SetRESPONSEPATH_PAGERESULTS(v string)`

SetRESPONSEPATH_PAGERESULTS sets RESPONSEPATH_PAGERESULTS field to given value.

### HasRESPONSEPATH_PAGERESULTS

`func (o *WorkdaySOAPConnectionAttributes) HasRESPONSEPATH_PAGERESULTS() bool`

HasRESPONSEPATH_PAGERESULTS returns a boolean if a field has been set.

### GetRESPONSEPATH_TOTALRESULTS

`func (o *WorkdaySOAPConnectionAttributes) GetRESPONSEPATH_TOTALRESULTS() string`

GetRESPONSEPATH_TOTALRESULTS returns the RESPONSEPATH_TOTALRESULTS field if non-nil, zero value otherwise.

### GetRESPONSEPATH_TOTALRESULTSOk

`func (o *WorkdaySOAPConnectionAttributes) GetRESPONSEPATH_TOTALRESULTSOk() (*string, bool)`

GetRESPONSEPATH_TOTALRESULTSOk returns a tuple with the RESPONSEPATH_TOTALRESULTS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRESPONSEPATH_TOTALRESULTS

`func (o *WorkdaySOAPConnectionAttributes) SetRESPONSEPATH_TOTALRESULTS(v string)`

SetRESPONSEPATH_TOTALRESULTS sets RESPONSEPATH_TOTALRESULTS field to given value.

### HasRESPONSEPATH_TOTALRESULTS

`func (o *WorkdaySOAPConnectionAttributes) HasRESPONSEPATH_TOTALRESULTS() bool`

HasRESPONSEPATH_TOTALRESULTS returns a boolean if a field has been set.

### GetRESPONSEPATH_USERLIST

`func (o *WorkdaySOAPConnectionAttributes) GetRESPONSEPATH_USERLIST() string`

GetRESPONSEPATH_USERLIST returns the RESPONSEPATH_USERLIST field if non-nil, zero value otherwise.

### GetRESPONSEPATH_USERLISTOk

`func (o *WorkdaySOAPConnectionAttributes) GetRESPONSEPATH_USERLISTOk() (*string, bool)`

GetRESPONSEPATH_USERLISTOk returns a tuple with the RESPONSEPATH_USERLIST field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRESPONSEPATH_USERLIST

`func (o *WorkdaySOAPConnectionAttributes) SetRESPONSEPATH_USERLIST(v string)`

SetRESPONSEPATH_USERLIST sets RESPONSEPATH_USERLIST field to given value.

### HasRESPONSEPATH_USERLIST

`func (o *WorkdaySOAPConnectionAttributes) HasRESPONSEPATH_USERLIST() bool`

HasRESPONSEPATH_USERLIST returns a boolean if a field has been set.

### GetREVOKEACCESSJSON

`func (o *WorkdaySOAPConnectionAttributes) GetREVOKEACCESSJSON() string`

GetREVOKEACCESSJSON returns the REVOKEACCESSJSON field if non-nil, zero value otherwise.

### GetREVOKEACCESSJSONOk

`func (o *WorkdaySOAPConnectionAttributes) GetREVOKEACCESSJSONOk() (*string, bool)`

GetREVOKEACCESSJSONOk returns a tuple with the REVOKEACCESSJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetREVOKEACCESSJSON

`func (o *WorkdaySOAPConnectionAttributes) SetREVOKEACCESSJSON(v string)`

SetREVOKEACCESSJSON sets REVOKEACCESSJSON field to given value.

### HasREVOKEACCESSJSON

`func (o *WorkdaySOAPConnectionAttributes) HasREVOKEACCESSJSON() bool`

HasREVOKEACCESSJSON returns a boolean if a field has been set.

### GetSOAP_ENDPOINT

`func (o *WorkdaySOAPConnectionAttributes) GetSOAP_ENDPOINT() string`

GetSOAP_ENDPOINT returns the SOAP_ENDPOINT field if non-nil, zero value otherwise.

### GetSOAP_ENDPOINTOk

`func (o *WorkdaySOAPConnectionAttributes) GetSOAP_ENDPOINTOk() (*string, bool)`

GetSOAP_ENDPOINTOk returns a tuple with the SOAP_ENDPOINT field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSOAP_ENDPOINT

`func (o *WorkdaySOAPConnectionAttributes) SetSOAP_ENDPOINT(v string)`

SetSOAP_ENDPOINT sets SOAP_ENDPOINT field to given value.

### HasSOAP_ENDPOINT

`func (o *WorkdaySOAPConnectionAttributes) HasSOAP_ENDPOINT() bool`

HasSOAP_ENDPOINT returns a boolean if a field has been set.

### GetUPDATEACCOUNTJSON

`func (o *WorkdaySOAPConnectionAttributes) GetUPDATEACCOUNTJSON() string`

GetUPDATEACCOUNTJSON returns the UPDATEACCOUNTJSON field if non-nil, zero value otherwise.

### GetUPDATEACCOUNTJSONOk

`func (o *WorkdaySOAPConnectionAttributes) GetUPDATEACCOUNTJSONOk() (*string, bool)`

GetUPDATEACCOUNTJSONOk returns a tuple with the UPDATEACCOUNTJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUPDATEACCOUNTJSON

`func (o *WorkdaySOAPConnectionAttributes) SetUPDATEACCOUNTJSON(v string)`

SetUPDATEACCOUNTJSON sets UPDATEACCOUNTJSON field to given value.

### HasUPDATEACCOUNTJSON

`func (o *WorkdaySOAPConnectionAttributes) HasUPDATEACCOUNTJSON() bool`

HasUPDATEACCOUNTJSON returns a boolean if a field has been set.

### GetUPDATEUSERJSON

`func (o *WorkdaySOAPConnectionAttributes) GetUPDATEUSERJSON() string`

GetUPDATEUSERJSON returns the UPDATEUSERJSON field if non-nil, zero value otherwise.

### GetUPDATEUSERJSONOk

`func (o *WorkdaySOAPConnectionAttributes) GetUPDATEUSERJSONOk() (*string, bool)`

GetUPDATEUSERJSONOk returns a tuple with the UPDATEUSERJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUPDATEUSERJSON

`func (o *WorkdaySOAPConnectionAttributes) SetUPDATEUSERJSON(v string)`

SetUPDATEUSERJSON sets UPDATEUSERJSON field to given value.

### HasUPDATEUSERJSON

`func (o *WorkdaySOAPConnectionAttributes) HasUPDATEUSERJSON() bool`

HasUPDATEUSERJSON returns a boolean if a field has been set.

### GetUSERNAME

`func (o *WorkdaySOAPConnectionAttributes) GetUSERNAME() string`

GetUSERNAME returns the USERNAME field if non-nil, zero value otherwise.

### GetUSERNAMEOk

`func (o *WorkdaySOAPConnectionAttributes) GetUSERNAMEOk() (*string, bool)`

GetUSERNAMEOk returns a tuple with the USERNAME field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUSERNAME

`func (o *WorkdaySOAPConnectionAttributes) SetUSERNAME(v string)`

SetUSERNAME sets USERNAME field to given value.

### HasUSERNAME

`func (o *WorkdaySOAPConnectionAttributes) HasUSERNAME() bool`

HasUSERNAME returns a boolean if a field has been set.

### GetConnectionTimeoutConfig

`func (o *WorkdaySOAPConnectionAttributes) GetConnectionTimeoutConfig() ConnectionTimeoutConfig`

GetConnectionTimeoutConfig returns the ConnectionTimeoutConfig field if non-nil, zero value otherwise.

### GetConnectionTimeoutConfigOk

`func (o *WorkdaySOAPConnectionAttributes) GetConnectionTimeoutConfigOk() (*ConnectionTimeoutConfig, bool)`

GetConnectionTimeoutConfigOk returns a tuple with the ConnectionTimeoutConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionTimeoutConfig

`func (o *WorkdaySOAPConnectionAttributes) SetConnectionTimeoutConfig(v ConnectionTimeoutConfig)`

SetConnectionTimeoutConfig sets ConnectionTimeoutConfig field to given value.

### HasConnectionTimeoutConfig

`func (o *WorkdaySOAPConnectionAttributes) HasConnectionTimeoutConfig() bool`

HasConnectionTimeoutConfig returns a boolean if a field has been set.

### GetConnectionType

`func (o *WorkdaySOAPConnectionAttributes) GetConnectionType() string`

GetConnectionType returns the ConnectionType field if non-nil, zero value otherwise.

### GetConnectionTypeOk

`func (o *WorkdaySOAPConnectionAttributes) GetConnectionTypeOk() (*string, bool)`

GetConnectionTypeOk returns a tuple with the ConnectionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionType

`func (o *WorkdaySOAPConnectionAttributes) SetConnectionType(v string)`

SetConnectionType sets ConnectionType field to given value.

### HasConnectionType

`func (o *WorkdaySOAPConnectionAttributes) HasConnectionType() bool`

HasConnectionType returns a boolean if a field has been set.

### GetIsTimeoutConfigValidated

`func (o *WorkdaySOAPConnectionAttributes) GetIsTimeoutConfigValidated() bool`

GetIsTimeoutConfigValidated returns the IsTimeoutConfigValidated field if non-nil, zero value otherwise.

### GetIsTimeoutConfigValidatedOk

`func (o *WorkdaySOAPConnectionAttributes) GetIsTimeoutConfigValidatedOk() (*bool, bool)`

GetIsTimeoutConfigValidatedOk returns a tuple with the IsTimeoutConfigValidated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsTimeoutConfigValidated

`func (o *WorkdaySOAPConnectionAttributes) SetIsTimeoutConfigValidated(v bool)`

SetIsTimeoutConfigValidated sets IsTimeoutConfigValidated field to given value.

### HasIsTimeoutConfigValidated

`func (o *WorkdaySOAPConnectionAttributes) HasIsTimeoutConfigValidated() bool`

HasIsTimeoutConfigValidated returns a boolean if a field has been set.

### GetIsTimeoutSupported

`func (o *WorkdaySOAPConnectionAttributes) GetIsTimeoutSupported() bool`

GetIsTimeoutSupported returns the IsTimeoutSupported field if non-nil, zero value otherwise.

### GetIsTimeoutSupportedOk

`func (o *WorkdaySOAPConnectionAttributes) GetIsTimeoutSupportedOk() (*bool, bool)`

GetIsTimeoutSupportedOk returns a tuple with the IsTimeoutSupported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsTimeoutSupported

`func (o *WorkdaySOAPConnectionAttributes) SetIsTimeoutSupported(v bool)`

SetIsTimeoutSupported sets IsTimeoutSupported field to given value.

### HasIsTimeoutSupported

`func (o *WorkdaySOAPConnectionAttributes) HasIsTimeoutSupported() bool`

HasIsTimeoutSupported returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


