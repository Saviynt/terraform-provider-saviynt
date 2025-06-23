# DBConnectionAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PASSWORD_MIN_LENGTH** | Pointer to **string** | Minimum length of password | [optional] 
**CHANGEPASSJSON** | Pointer to **string** |  | [optional] 
**ACCOUNTEXISTSJSON** | Pointer to **string** |  | [optional] 
**ROLESIMPORT** | Pointer to **string** |  | [optional] 
**ROLEOWNERIMPORT** | Pointer to **string** |  | [optional] 
**CREATEACCOUNTJSON** | Pointer to **string** |  | [optional] 
**USERIMPORT** | Pointer to **string** |  | [optional] 
**DISABLEACCOUNTJSON** | Pointer to **string** |  | [optional] 
**ENTITLEMENTVALUEIMPORT** | Pointer to **string** |  | [optional] 
**ConnectionTimeoutConfig** | Pointer to [**ConnectionTimeoutConfig**](ConnectionTimeoutConfig.md) |  | [optional] 
**UPDATEUSERJSON** | Pointer to **string** |  | [optional] 
**PASSWORD_NOOFSPLCHARS** | Pointer to **string** |  | [optional] 
**REVOKEACCESSJSON** | Pointer to **string** |  | [optional] 
**ConnectionType** | Pointer to **string** |  | [optional] 
**URL** | Pointer to **string** |  | [optional] 
**SYSTEMIMPORT** | Pointer to **string** |  | [optional] 
**DRIVERNAME** | Pointer to **string** |  | [optional] 
**DELETEACCOUNTJSON** | Pointer to **string** |  | [optional] 
**STATUS_THRESHOLD_CONFIG** | Pointer to **string** |  | [optional] 
**USERNAME** | Pointer to **string** |  | [optional] 
**IsTimeoutSupported** | Pointer to **bool** |  | [optional] 
**PASSWORD_NOOFCAPSALPHA** | Pointer to **string** |  | [optional] 
**PASSWORD_NOOFDIGITS** | Pointer to **string** |  | [optional] 
**CONNECTIONPROPERTIES** | Pointer to **string** |  | [optional] 
**MODIFYUSERDATAJSON** | Pointer to **string** |  | [optional] 
**IsTimeoutConfigValidated** | Pointer to **bool** |  | [optional] 
**ACCOUNTSIMPORT** | Pointer to **string** |  | [optional] 
**PASSWORD** | Pointer to **string** |  | [optional] 
**ENABLEACCOUNTJSON** | Pointer to **string** |  | [optional] 
**PASSWORD_MAX_LENGTH** | Pointer to **string** |  | [optional] 
**MAX_PAGINATION_SIZE** | Pointer to **string** |  | [optional] 
**UPDATEACCOUNTJSON** | Pointer to **string** |  | [optional] 
**GRANTACCESSJSON** | Pointer to **string** |  | [optional] 
**CLI_COMMAND_JSON** | Pointer to **string** |  | [optional] 

## Methods

### NewDBConnectionAttributes

`func NewDBConnectionAttributes() *DBConnectionAttributes`

NewDBConnectionAttributes instantiates a new DBConnectionAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDBConnectionAttributesWithDefaults

`func NewDBConnectionAttributesWithDefaults() *DBConnectionAttributes`

NewDBConnectionAttributesWithDefaults instantiates a new DBConnectionAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPASSWORD_MIN_LENGTH

`func (o *DBConnectionAttributes) GetPASSWORD_MIN_LENGTH() string`

GetPASSWORD_MIN_LENGTH returns the PASSWORD_MIN_LENGTH field if non-nil, zero value otherwise.

### GetPASSWORD_MIN_LENGTHOk

`func (o *DBConnectionAttributes) GetPASSWORD_MIN_LENGTHOk() (*string, bool)`

GetPASSWORD_MIN_LENGTHOk returns a tuple with the PASSWORD_MIN_LENGTH field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPASSWORD_MIN_LENGTH

`func (o *DBConnectionAttributes) SetPASSWORD_MIN_LENGTH(v string)`

SetPASSWORD_MIN_LENGTH sets PASSWORD_MIN_LENGTH field to given value.

### HasPASSWORD_MIN_LENGTH

`func (o *DBConnectionAttributes) HasPASSWORD_MIN_LENGTH() bool`

HasPASSWORD_MIN_LENGTH returns a boolean if a field has been set.

### GetCHANGEPASSJSON

`func (o *DBConnectionAttributes) GetCHANGEPASSJSON() string`

GetCHANGEPASSJSON returns the CHANGEPASSJSON field if non-nil, zero value otherwise.

### GetCHANGEPASSJSONOk

`func (o *DBConnectionAttributes) GetCHANGEPASSJSONOk() (*string, bool)`

GetCHANGEPASSJSONOk returns a tuple with the CHANGEPASSJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCHANGEPASSJSON

`func (o *DBConnectionAttributes) SetCHANGEPASSJSON(v string)`

SetCHANGEPASSJSON sets CHANGEPASSJSON field to given value.

### HasCHANGEPASSJSON

`func (o *DBConnectionAttributes) HasCHANGEPASSJSON() bool`

HasCHANGEPASSJSON returns a boolean if a field has been set.

### GetACCOUNTEXISTSJSON

`func (o *DBConnectionAttributes) GetACCOUNTEXISTSJSON() string`

GetACCOUNTEXISTSJSON returns the ACCOUNTEXISTSJSON field if non-nil, zero value otherwise.

### GetACCOUNTEXISTSJSONOk

`func (o *DBConnectionAttributes) GetACCOUNTEXISTSJSONOk() (*string, bool)`

GetACCOUNTEXISTSJSONOk returns a tuple with the ACCOUNTEXISTSJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetACCOUNTEXISTSJSON

`func (o *DBConnectionAttributes) SetACCOUNTEXISTSJSON(v string)`

SetACCOUNTEXISTSJSON sets ACCOUNTEXISTSJSON field to given value.

### HasACCOUNTEXISTSJSON

`func (o *DBConnectionAttributes) HasACCOUNTEXISTSJSON() bool`

HasACCOUNTEXISTSJSON returns a boolean if a field has been set.

### GetROLESIMPORT

`func (o *DBConnectionAttributes) GetROLESIMPORT() string`

GetROLESIMPORT returns the ROLESIMPORT field if non-nil, zero value otherwise.

### GetROLESIMPORTOk

`func (o *DBConnectionAttributes) GetROLESIMPORTOk() (*string, bool)`

GetROLESIMPORTOk returns a tuple with the ROLESIMPORT field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetROLESIMPORT

`func (o *DBConnectionAttributes) SetROLESIMPORT(v string)`

SetROLESIMPORT sets ROLESIMPORT field to given value.

### HasROLESIMPORT

`func (o *DBConnectionAttributes) HasROLESIMPORT() bool`

HasROLESIMPORT returns a boolean if a field has been set.

### GetROLEOWNERIMPORT

`func (o *DBConnectionAttributes) GetROLEOWNERIMPORT() string`

GetROLEOWNERIMPORT returns the ROLEOWNERIMPORT field if non-nil, zero value otherwise.

### GetROLEOWNERIMPORTOk

`func (o *DBConnectionAttributes) GetROLEOWNERIMPORTOk() (*string, bool)`

GetROLEOWNERIMPORTOk returns a tuple with the ROLEOWNERIMPORT field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetROLEOWNERIMPORT

`func (o *DBConnectionAttributes) SetROLEOWNERIMPORT(v string)`

SetROLEOWNERIMPORT sets ROLEOWNERIMPORT field to given value.

### HasROLEOWNERIMPORT

`func (o *DBConnectionAttributes) HasROLEOWNERIMPORT() bool`

HasROLEOWNERIMPORT returns a boolean if a field has been set.

### GetCREATEACCOUNTJSON

`func (o *DBConnectionAttributes) GetCREATEACCOUNTJSON() string`

GetCREATEACCOUNTJSON returns the CREATEACCOUNTJSON field if non-nil, zero value otherwise.

### GetCREATEACCOUNTJSONOk

`func (o *DBConnectionAttributes) GetCREATEACCOUNTJSONOk() (*string, bool)`

GetCREATEACCOUNTJSONOk returns a tuple with the CREATEACCOUNTJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCREATEACCOUNTJSON

`func (o *DBConnectionAttributes) SetCREATEACCOUNTJSON(v string)`

SetCREATEACCOUNTJSON sets CREATEACCOUNTJSON field to given value.

### HasCREATEACCOUNTJSON

`func (o *DBConnectionAttributes) HasCREATEACCOUNTJSON() bool`

HasCREATEACCOUNTJSON returns a boolean if a field has been set.

### GetUSERIMPORT

`func (o *DBConnectionAttributes) GetUSERIMPORT() string`

GetUSERIMPORT returns the USERIMPORT field if non-nil, zero value otherwise.

### GetUSERIMPORTOk

`func (o *DBConnectionAttributes) GetUSERIMPORTOk() (*string, bool)`

GetUSERIMPORTOk returns a tuple with the USERIMPORT field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUSERIMPORT

`func (o *DBConnectionAttributes) SetUSERIMPORT(v string)`

SetUSERIMPORT sets USERIMPORT field to given value.

### HasUSERIMPORT

`func (o *DBConnectionAttributes) HasUSERIMPORT() bool`

HasUSERIMPORT returns a boolean if a field has been set.

### GetDISABLEACCOUNTJSON

`func (o *DBConnectionAttributes) GetDISABLEACCOUNTJSON() string`

GetDISABLEACCOUNTJSON returns the DISABLEACCOUNTJSON field if non-nil, zero value otherwise.

### GetDISABLEACCOUNTJSONOk

`func (o *DBConnectionAttributes) GetDISABLEACCOUNTJSONOk() (*string, bool)`

GetDISABLEACCOUNTJSONOk returns a tuple with the DISABLEACCOUNTJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDISABLEACCOUNTJSON

`func (o *DBConnectionAttributes) SetDISABLEACCOUNTJSON(v string)`

SetDISABLEACCOUNTJSON sets DISABLEACCOUNTJSON field to given value.

### HasDISABLEACCOUNTJSON

`func (o *DBConnectionAttributes) HasDISABLEACCOUNTJSON() bool`

HasDISABLEACCOUNTJSON returns a boolean if a field has been set.

### GetENTITLEMENTVALUEIMPORT

`func (o *DBConnectionAttributes) GetENTITLEMENTVALUEIMPORT() string`

GetENTITLEMENTVALUEIMPORT returns the ENTITLEMENTVALUEIMPORT field if non-nil, zero value otherwise.

### GetENTITLEMENTVALUEIMPORTOk

`func (o *DBConnectionAttributes) GetENTITLEMENTVALUEIMPORTOk() (*string, bool)`

GetENTITLEMENTVALUEIMPORTOk returns a tuple with the ENTITLEMENTVALUEIMPORT field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetENTITLEMENTVALUEIMPORT

`func (o *DBConnectionAttributes) SetENTITLEMENTVALUEIMPORT(v string)`

SetENTITLEMENTVALUEIMPORT sets ENTITLEMENTVALUEIMPORT field to given value.

### HasENTITLEMENTVALUEIMPORT

`func (o *DBConnectionAttributes) HasENTITLEMENTVALUEIMPORT() bool`

HasENTITLEMENTVALUEIMPORT returns a boolean if a field has been set.

### GetConnectionTimeoutConfig

`func (o *DBConnectionAttributes) GetConnectionTimeoutConfig() ConnectionTimeoutConfig`

GetConnectionTimeoutConfig returns the ConnectionTimeoutConfig field if non-nil, zero value otherwise.

### GetConnectionTimeoutConfigOk

`func (o *DBConnectionAttributes) GetConnectionTimeoutConfigOk() (*ConnectionTimeoutConfig, bool)`

GetConnectionTimeoutConfigOk returns a tuple with the ConnectionTimeoutConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionTimeoutConfig

`func (o *DBConnectionAttributes) SetConnectionTimeoutConfig(v ConnectionTimeoutConfig)`

SetConnectionTimeoutConfig sets ConnectionTimeoutConfig field to given value.

### HasConnectionTimeoutConfig

`func (o *DBConnectionAttributes) HasConnectionTimeoutConfig() bool`

HasConnectionTimeoutConfig returns a boolean if a field has been set.

### GetUPDATEUSERJSON

`func (o *DBConnectionAttributes) GetUPDATEUSERJSON() string`

GetUPDATEUSERJSON returns the UPDATEUSERJSON field if non-nil, zero value otherwise.

### GetUPDATEUSERJSONOk

`func (o *DBConnectionAttributes) GetUPDATEUSERJSONOk() (*string, bool)`

GetUPDATEUSERJSONOk returns a tuple with the UPDATEUSERJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUPDATEUSERJSON

`func (o *DBConnectionAttributes) SetUPDATEUSERJSON(v string)`

SetUPDATEUSERJSON sets UPDATEUSERJSON field to given value.

### HasUPDATEUSERJSON

`func (o *DBConnectionAttributes) HasUPDATEUSERJSON() bool`

HasUPDATEUSERJSON returns a boolean if a field has been set.

### GetPASSWORD_NOOFSPLCHARS

`func (o *DBConnectionAttributes) GetPASSWORD_NOOFSPLCHARS() string`

GetPASSWORD_NOOFSPLCHARS returns the PASSWORD_NOOFSPLCHARS field if non-nil, zero value otherwise.

### GetPASSWORD_NOOFSPLCHARSOk

`func (o *DBConnectionAttributes) GetPASSWORD_NOOFSPLCHARSOk() (*string, bool)`

GetPASSWORD_NOOFSPLCHARSOk returns a tuple with the PASSWORD_NOOFSPLCHARS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPASSWORD_NOOFSPLCHARS

`func (o *DBConnectionAttributes) SetPASSWORD_NOOFSPLCHARS(v string)`

SetPASSWORD_NOOFSPLCHARS sets PASSWORD_NOOFSPLCHARS field to given value.

### HasPASSWORD_NOOFSPLCHARS

`func (o *DBConnectionAttributes) HasPASSWORD_NOOFSPLCHARS() bool`

HasPASSWORD_NOOFSPLCHARS returns a boolean if a field has been set.

### GetREVOKEACCESSJSON

`func (o *DBConnectionAttributes) GetREVOKEACCESSJSON() string`

GetREVOKEACCESSJSON returns the REVOKEACCESSJSON field if non-nil, zero value otherwise.

### GetREVOKEACCESSJSONOk

`func (o *DBConnectionAttributes) GetREVOKEACCESSJSONOk() (*string, bool)`

GetREVOKEACCESSJSONOk returns a tuple with the REVOKEACCESSJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetREVOKEACCESSJSON

`func (o *DBConnectionAttributes) SetREVOKEACCESSJSON(v string)`

SetREVOKEACCESSJSON sets REVOKEACCESSJSON field to given value.

### HasREVOKEACCESSJSON

`func (o *DBConnectionAttributes) HasREVOKEACCESSJSON() bool`

HasREVOKEACCESSJSON returns a boolean if a field has been set.

### GetConnectionType

`func (o *DBConnectionAttributes) GetConnectionType() string`

GetConnectionType returns the ConnectionType field if non-nil, zero value otherwise.

### GetConnectionTypeOk

`func (o *DBConnectionAttributes) GetConnectionTypeOk() (*string, bool)`

GetConnectionTypeOk returns a tuple with the ConnectionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionType

`func (o *DBConnectionAttributes) SetConnectionType(v string)`

SetConnectionType sets ConnectionType field to given value.

### HasConnectionType

`func (o *DBConnectionAttributes) HasConnectionType() bool`

HasConnectionType returns a boolean if a field has been set.

### GetURL

`func (o *DBConnectionAttributes) GetURL() string`

GetURL returns the URL field if non-nil, zero value otherwise.

### GetURLOk

`func (o *DBConnectionAttributes) GetURLOk() (*string, bool)`

GetURLOk returns a tuple with the URL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetURL

`func (o *DBConnectionAttributes) SetURL(v string)`

SetURL sets URL field to given value.

### HasURL

`func (o *DBConnectionAttributes) HasURL() bool`

HasURL returns a boolean if a field has been set.

### GetSYSTEMIMPORT

`func (o *DBConnectionAttributes) GetSYSTEMIMPORT() string`

GetSYSTEMIMPORT returns the SYSTEMIMPORT field if non-nil, zero value otherwise.

### GetSYSTEMIMPORTOk

`func (o *DBConnectionAttributes) GetSYSTEMIMPORTOk() (*string, bool)`

GetSYSTEMIMPORTOk returns a tuple with the SYSTEMIMPORT field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSYSTEMIMPORT

`func (o *DBConnectionAttributes) SetSYSTEMIMPORT(v string)`

SetSYSTEMIMPORT sets SYSTEMIMPORT field to given value.

### HasSYSTEMIMPORT

`func (o *DBConnectionAttributes) HasSYSTEMIMPORT() bool`

HasSYSTEMIMPORT returns a boolean if a field has been set.

### GetDRIVERNAME

`func (o *DBConnectionAttributes) GetDRIVERNAME() string`

GetDRIVERNAME returns the DRIVERNAME field if non-nil, zero value otherwise.

### GetDRIVERNAMEOk

`func (o *DBConnectionAttributes) GetDRIVERNAMEOk() (*string, bool)`

GetDRIVERNAMEOk returns a tuple with the DRIVERNAME field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDRIVERNAME

`func (o *DBConnectionAttributes) SetDRIVERNAME(v string)`

SetDRIVERNAME sets DRIVERNAME field to given value.

### HasDRIVERNAME

`func (o *DBConnectionAttributes) HasDRIVERNAME() bool`

HasDRIVERNAME returns a boolean if a field has been set.

### GetDELETEACCOUNTJSON

`func (o *DBConnectionAttributes) GetDELETEACCOUNTJSON() string`

GetDELETEACCOUNTJSON returns the DELETEACCOUNTJSON field if non-nil, zero value otherwise.

### GetDELETEACCOUNTJSONOk

`func (o *DBConnectionAttributes) GetDELETEACCOUNTJSONOk() (*string, bool)`

GetDELETEACCOUNTJSONOk returns a tuple with the DELETEACCOUNTJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDELETEACCOUNTJSON

`func (o *DBConnectionAttributes) SetDELETEACCOUNTJSON(v string)`

SetDELETEACCOUNTJSON sets DELETEACCOUNTJSON field to given value.

### HasDELETEACCOUNTJSON

`func (o *DBConnectionAttributes) HasDELETEACCOUNTJSON() bool`

HasDELETEACCOUNTJSON returns a boolean if a field has been set.

### GetSTATUS_THRESHOLD_CONFIG

`func (o *DBConnectionAttributes) GetSTATUS_THRESHOLD_CONFIG() string`

GetSTATUS_THRESHOLD_CONFIG returns the STATUS_THRESHOLD_CONFIG field if non-nil, zero value otherwise.

### GetSTATUS_THRESHOLD_CONFIGOk

`func (o *DBConnectionAttributes) GetSTATUS_THRESHOLD_CONFIGOk() (*string, bool)`

GetSTATUS_THRESHOLD_CONFIGOk returns a tuple with the STATUS_THRESHOLD_CONFIG field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSTATUS_THRESHOLD_CONFIG

`func (o *DBConnectionAttributes) SetSTATUS_THRESHOLD_CONFIG(v string)`

SetSTATUS_THRESHOLD_CONFIG sets STATUS_THRESHOLD_CONFIG field to given value.

### HasSTATUS_THRESHOLD_CONFIG

`func (o *DBConnectionAttributes) HasSTATUS_THRESHOLD_CONFIG() bool`

HasSTATUS_THRESHOLD_CONFIG returns a boolean if a field has been set.

### GetUSERNAME

`func (o *DBConnectionAttributes) GetUSERNAME() string`

GetUSERNAME returns the USERNAME field if non-nil, zero value otherwise.

### GetUSERNAMEOk

`func (o *DBConnectionAttributes) GetUSERNAMEOk() (*string, bool)`

GetUSERNAMEOk returns a tuple with the USERNAME field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUSERNAME

`func (o *DBConnectionAttributes) SetUSERNAME(v string)`

SetUSERNAME sets USERNAME field to given value.

### HasUSERNAME

`func (o *DBConnectionAttributes) HasUSERNAME() bool`

HasUSERNAME returns a boolean if a field has been set.

### GetIsTimeoutSupported

`func (o *DBConnectionAttributes) GetIsTimeoutSupported() bool`

GetIsTimeoutSupported returns the IsTimeoutSupported field if non-nil, zero value otherwise.

### GetIsTimeoutSupportedOk

`func (o *DBConnectionAttributes) GetIsTimeoutSupportedOk() (*bool, bool)`

GetIsTimeoutSupportedOk returns a tuple with the IsTimeoutSupported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsTimeoutSupported

`func (o *DBConnectionAttributes) SetIsTimeoutSupported(v bool)`

SetIsTimeoutSupported sets IsTimeoutSupported field to given value.

### HasIsTimeoutSupported

`func (o *DBConnectionAttributes) HasIsTimeoutSupported() bool`

HasIsTimeoutSupported returns a boolean if a field has been set.

### GetPASSWORD_NOOFCAPSALPHA

`func (o *DBConnectionAttributes) GetPASSWORD_NOOFCAPSALPHA() string`

GetPASSWORD_NOOFCAPSALPHA returns the PASSWORD_NOOFCAPSALPHA field if non-nil, zero value otherwise.

### GetPASSWORD_NOOFCAPSALPHAOk

`func (o *DBConnectionAttributes) GetPASSWORD_NOOFCAPSALPHAOk() (*string, bool)`

GetPASSWORD_NOOFCAPSALPHAOk returns a tuple with the PASSWORD_NOOFCAPSALPHA field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPASSWORD_NOOFCAPSALPHA

`func (o *DBConnectionAttributes) SetPASSWORD_NOOFCAPSALPHA(v string)`

SetPASSWORD_NOOFCAPSALPHA sets PASSWORD_NOOFCAPSALPHA field to given value.

### HasPASSWORD_NOOFCAPSALPHA

`func (o *DBConnectionAttributes) HasPASSWORD_NOOFCAPSALPHA() bool`

HasPASSWORD_NOOFCAPSALPHA returns a boolean if a field has been set.

### GetPASSWORD_NOOFDIGITS

`func (o *DBConnectionAttributes) GetPASSWORD_NOOFDIGITS() string`

GetPASSWORD_NOOFDIGITS returns the PASSWORD_NOOFDIGITS field if non-nil, zero value otherwise.

### GetPASSWORD_NOOFDIGITSOk

`func (o *DBConnectionAttributes) GetPASSWORD_NOOFDIGITSOk() (*string, bool)`

GetPASSWORD_NOOFDIGITSOk returns a tuple with the PASSWORD_NOOFDIGITS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPASSWORD_NOOFDIGITS

`func (o *DBConnectionAttributes) SetPASSWORD_NOOFDIGITS(v string)`

SetPASSWORD_NOOFDIGITS sets PASSWORD_NOOFDIGITS field to given value.

### HasPASSWORD_NOOFDIGITS

`func (o *DBConnectionAttributes) HasPASSWORD_NOOFDIGITS() bool`

HasPASSWORD_NOOFDIGITS returns a boolean if a field has been set.

### GetCONNECTIONPROPERTIES

`func (o *DBConnectionAttributes) GetCONNECTIONPROPERTIES() string`

GetCONNECTIONPROPERTIES returns the CONNECTIONPROPERTIES field if non-nil, zero value otherwise.

### GetCONNECTIONPROPERTIESOk

`func (o *DBConnectionAttributes) GetCONNECTIONPROPERTIESOk() (*string, bool)`

GetCONNECTIONPROPERTIESOk returns a tuple with the CONNECTIONPROPERTIES field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCONNECTIONPROPERTIES

`func (o *DBConnectionAttributes) SetCONNECTIONPROPERTIES(v string)`

SetCONNECTIONPROPERTIES sets CONNECTIONPROPERTIES field to given value.

### HasCONNECTIONPROPERTIES

`func (o *DBConnectionAttributes) HasCONNECTIONPROPERTIES() bool`

HasCONNECTIONPROPERTIES returns a boolean if a field has been set.

### GetMODIFYUSERDATAJSON

`func (o *DBConnectionAttributes) GetMODIFYUSERDATAJSON() string`

GetMODIFYUSERDATAJSON returns the MODIFYUSERDATAJSON field if non-nil, zero value otherwise.

### GetMODIFYUSERDATAJSONOk

`func (o *DBConnectionAttributes) GetMODIFYUSERDATAJSONOk() (*string, bool)`

GetMODIFYUSERDATAJSONOk returns a tuple with the MODIFYUSERDATAJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMODIFYUSERDATAJSON

`func (o *DBConnectionAttributes) SetMODIFYUSERDATAJSON(v string)`

SetMODIFYUSERDATAJSON sets MODIFYUSERDATAJSON field to given value.

### HasMODIFYUSERDATAJSON

`func (o *DBConnectionAttributes) HasMODIFYUSERDATAJSON() bool`

HasMODIFYUSERDATAJSON returns a boolean if a field has been set.

### GetIsTimeoutConfigValidated

`func (o *DBConnectionAttributes) GetIsTimeoutConfigValidated() bool`

GetIsTimeoutConfigValidated returns the IsTimeoutConfigValidated field if non-nil, zero value otherwise.

### GetIsTimeoutConfigValidatedOk

`func (o *DBConnectionAttributes) GetIsTimeoutConfigValidatedOk() (*bool, bool)`

GetIsTimeoutConfigValidatedOk returns a tuple with the IsTimeoutConfigValidated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsTimeoutConfigValidated

`func (o *DBConnectionAttributes) SetIsTimeoutConfigValidated(v bool)`

SetIsTimeoutConfigValidated sets IsTimeoutConfigValidated field to given value.

### HasIsTimeoutConfigValidated

`func (o *DBConnectionAttributes) HasIsTimeoutConfigValidated() bool`

HasIsTimeoutConfigValidated returns a boolean if a field has been set.

### GetACCOUNTSIMPORT

`func (o *DBConnectionAttributes) GetACCOUNTSIMPORT() string`

GetACCOUNTSIMPORT returns the ACCOUNTSIMPORT field if non-nil, zero value otherwise.

### GetACCOUNTSIMPORTOk

`func (o *DBConnectionAttributes) GetACCOUNTSIMPORTOk() (*string, bool)`

GetACCOUNTSIMPORTOk returns a tuple with the ACCOUNTSIMPORT field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetACCOUNTSIMPORT

`func (o *DBConnectionAttributes) SetACCOUNTSIMPORT(v string)`

SetACCOUNTSIMPORT sets ACCOUNTSIMPORT field to given value.

### HasACCOUNTSIMPORT

`func (o *DBConnectionAttributes) HasACCOUNTSIMPORT() bool`

HasACCOUNTSIMPORT returns a boolean if a field has been set.

### GetPASSWORD

`func (o *DBConnectionAttributes) GetPASSWORD() string`

GetPASSWORD returns the PASSWORD field if non-nil, zero value otherwise.

### GetPASSWORDOk

`func (o *DBConnectionAttributes) GetPASSWORDOk() (*string, bool)`

GetPASSWORDOk returns a tuple with the PASSWORD field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPASSWORD

`func (o *DBConnectionAttributes) SetPASSWORD(v string)`

SetPASSWORD sets PASSWORD field to given value.

### HasPASSWORD

`func (o *DBConnectionAttributes) HasPASSWORD() bool`

HasPASSWORD returns a boolean if a field has been set.

### GetENABLEACCOUNTJSON

`func (o *DBConnectionAttributes) GetENABLEACCOUNTJSON() string`

GetENABLEACCOUNTJSON returns the ENABLEACCOUNTJSON field if non-nil, zero value otherwise.

### GetENABLEACCOUNTJSONOk

`func (o *DBConnectionAttributes) GetENABLEACCOUNTJSONOk() (*string, bool)`

GetENABLEACCOUNTJSONOk returns a tuple with the ENABLEACCOUNTJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetENABLEACCOUNTJSON

`func (o *DBConnectionAttributes) SetENABLEACCOUNTJSON(v string)`

SetENABLEACCOUNTJSON sets ENABLEACCOUNTJSON field to given value.

### HasENABLEACCOUNTJSON

`func (o *DBConnectionAttributes) HasENABLEACCOUNTJSON() bool`

HasENABLEACCOUNTJSON returns a boolean if a field has been set.

### GetPASSWORD_MAX_LENGTH

`func (o *DBConnectionAttributes) GetPASSWORD_MAX_LENGTH() string`

GetPASSWORD_MAX_LENGTH returns the PASSWORD_MAX_LENGTH field if non-nil, zero value otherwise.

### GetPASSWORD_MAX_LENGTHOk

`func (o *DBConnectionAttributes) GetPASSWORD_MAX_LENGTHOk() (*string, bool)`

GetPASSWORD_MAX_LENGTHOk returns a tuple with the PASSWORD_MAX_LENGTH field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPASSWORD_MAX_LENGTH

`func (o *DBConnectionAttributes) SetPASSWORD_MAX_LENGTH(v string)`

SetPASSWORD_MAX_LENGTH sets PASSWORD_MAX_LENGTH field to given value.

### HasPASSWORD_MAX_LENGTH

`func (o *DBConnectionAttributes) HasPASSWORD_MAX_LENGTH() bool`

HasPASSWORD_MAX_LENGTH returns a boolean if a field has been set.

### GetMAX_PAGINATION_SIZE

`func (o *DBConnectionAttributes) GetMAX_PAGINATION_SIZE() string`

GetMAX_PAGINATION_SIZE returns the MAX_PAGINATION_SIZE field if non-nil, zero value otherwise.

### GetMAX_PAGINATION_SIZEOk

`func (o *DBConnectionAttributes) GetMAX_PAGINATION_SIZEOk() (*string, bool)`

GetMAX_PAGINATION_SIZEOk returns a tuple with the MAX_PAGINATION_SIZE field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMAX_PAGINATION_SIZE

`func (o *DBConnectionAttributes) SetMAX_PAGINATION_SIZE(v string)`

SetMAX_PAGINATION_SIZE sets MAX_PAGINATION_SIZE field to given value.

### HasMAX_PAGINATION_SIZE

`func (o *DBConnectionAttributes) HasMAX_PAGINATION_SIZE() bool`

HasMAX_PAGINATION_SIZE returns a boolean if a field has been set.

### GetUPDATEACCOUNTJSON

`func (o *DBConnectionAttributes) GetUPDATEACCOUNTJSON() string`

GetUPDATEACCOUNTJSON returns the UPDATEACCOUNTJSON field if non-nil, zero value otherwise.

### GetUPDATEACCOUNTJSONOk

`func (o *DBConnectionAttributes) GetUPDATEACCOUNTJSONOk() (*string, bool)`

GetUPDATEACCOUNTJSONOk returns a tuple with the UPDATEACCOUNTJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUPDATEACCOUNTJSON

`func (o *DBConnectionAttributes) SetUPDATEACCOUNTJSON(v string)`

SetUPDATEACCOUNTJSON sets UPDATEACCOUNTJSON field to given value.

### HasUPDATEACCOUNTJSON

`func (o *DBConnectionAttributes) HasUPDATEACCOUNTJSON() bool`

HasUPDATEACCOUNTJSON returns a boolean if a field has been set.

### GetGRANTACCESSJSON

`func (o *DBConnectionAttributes) GetGRANTACCESSJSON() string`

GetGRANTACCESSJSON returns the GRANTACCESSJSON field if non-nil, zero value otherwise.

### GetGRANTACCESSJSONOk

`func (o *DBConnectionAttributes) GetGRANTACCESSJSONOk() (*string, bool)`

GetGRANTACCESSJSONOk returns a tuple with the GRANTACCESSJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGRANTACCESSJSON

`func (o *DBConnectionAttributes) SetGRANTACCESSJSON(v string)`

SetGRANTACCESSJSON sets GRANTACCESSJSON field to given value.

### HasGRANTACCESSJSON

`func (o *DBConnectionAttributes) HasGRANTACCESSJSON() bool`

HasGRANTACCESSJSON returns a boolean if a field has been set.

### GetCLI_COMMAND_JSON

`func (o *DBConnectionAttributes) GetCLI_COMMAND_JSON() string`

GetCLI_COMMAND_JSON returns the CLI_COMMAND_JSON field if non-nil, zero value otherwise.

### GetCLI_COMMAND_JSONOk

`func (o *DBConnectionAttributes) GetCLI_COMMAND_JSONOk() (*string, bool)`

GetCLI_COMMAND_JSONOk returns a tuple with the CLI_COMMAND_JSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCLI_COMMAND_JSON

`func (o *DBConnectionAttributes) SetCLI_COMMAND_JSON(v string)`

SetCLI_COMMAND_JSON sets CLI_COMMAND_JSON field to given value.

### HasCLI_COMMAND_JSON

`func (o *DBConnectionAttributes) HasCLI_COMMAND_JSON() bool`

HasCLI_COMMAND_JSON returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


