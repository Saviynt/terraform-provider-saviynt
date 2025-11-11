# OktaConnectionAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IMPORTURL** | Pointer to **string** |  | [optional] 
**AUTHTOKEN** | Pointer to **string** |  | [optional] 
**IsTimeoutSupported** | Pointer to **bool** |  | [optional] 
**OKTA_GROUPS_FILTER** | Pointer to **string** |  | [optional] 
**ACCOUNTFIELDMAPPINGS** | Pointer to **string** |  | [optional] 
**IMPORT_INACTIVE_APPS** | Pointer to **string** |  | [optional] 
**AUDIT_FILTER** | Pointer to **string** |  | [optional] 
**USERFIELDMAPPINGS** | Pointer to **string** |  | [optional] 
**APPACCOUNTFIELDMAPPINGS** | Pointer to **string** |  | [optional] 
**MODIFYUSERDATAJSON** | Pointer to **string** |  | [optional] 
**ConnectionTimeoutConfig** | Pointer to [**ConnectionTimeoutConfig**](ConnectionTimeoutConfig.md) |  | [optional] 
**ConnectionType** | Pointer to **string** |  | [optional] 
**IsTimeoutConfigValidated** | Pointer to **bool** |  | [optional] 
**ACTIVATE_ENDPOINT** | Pointer to **string** |  | [optional] 
**ENTITLEMENTTYPESMAPPINGS** | Pointer to **string** |  | [optional] 
**OKTA_APPLICATION_SECURITYSYSTEM** | Pointer to **string** |  | [optional] 
**PAM_CONFIG** | Pointer to **string** |  | [optional] 
**ConfigJSON** | Pointer to **string** |  | [optional] 
**STATUS_THRESHOLD_CONFIG** | Pointer to **string** |  | [optional] 

## Methods

### NewOktaConnectionAttributes

`func NewOktaConnectionAttributes() *OktaConnectionAttributes`

NewOktaConnectionAttributes instantiates a new OktaConnectionAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOktaConnectionAttributesWithDefaults

`func NewOktaConnectionAttributesWithDefaults() *OktaConnectionAttributes`

NewOktaConnectionAttributesWithDefaults instantiates a new OktaConnectionAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIMPORTURL

`func (o *OktaConnectionAttributes) GetIMPORTURL() string`

GetIMPORTURL returns the IMPORTURL field if non-nil, zero value otherwise.

### GetIMPORTURLOk

`func (o *OktaConnectionAttributes) GetIMPORTURLOk() (*string, bool)`

GetIMPORTURLOk returns a tuple with the IMPORTURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIMPORTURL

`func (o *OktaConnectionAttributes) SetIMPORTURL(v string)`

SetIMPORTURL sets IMPORTURL field to given value.

### HasIMPORTURL

`func (o *OktaConnectionAttributes) HasIMPORTURL() bool`

HasIMPORTURL returns a boolean if a field has been set.

### GetAUTHTOKEN

`func (o *OktaConnectionAttributes) GetAUTHTOKEN() string`

GetAUTHTOKEN returns the AUTHTOKEN field if non-nil, zero value otherwise.

### GetAUTHTOKENOk

`func (o *OktaConnectionAttributes) GetAUTHTOKENOk() (*string, bool)`

GetAUTHTOKENOk returns a tuple with the AUTHTOKEN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAUTHTOKEN

`func (o *OktaConnectionAttributes) SetAUTHTOKEN(v string)`

SetAUTHTOKEN sets AUTHTOKEN field to given value.

### HasAUTHTOKEN

`func (o *OktaConnectionAttributes) HasAUTHTOKEN() bool`

HasAUTHTOKEN returns a boolean if a field has been set.

### GetIsTimeoutSupported

`func (o *OktaConnectionAttributes) GetIsTimeoutSupported() bool`

GetIsTimeoutSupported returns the IsTimeoutSupported field if non-nil, zero value otherwise.

### GetIsTimeoutSupportedOk

`func (o *OktaConnectionAttributes) GetIsTimeoutSupportedOk() (*bool, bool)`

GetIsTimeoutSupportedOk returns a tuple with the IsTimeoutSupported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsTimeoutSupported

`func (o *OktaConnectionAttributes) SetIsTimeoutSupported(v bool)`

SetIsTimeoutSupported sets IsTimeoutSupported field to given value.

### HasIsTimeoutSupported

`func (o *OktaConnectionAttributes) HasIsTimeoutSupported() bool`

HasIsTimeoutSupported returns a boolean if a field has been set.

### GetOKTA_GROUPS_FILTER

`func (o *OktaConnectionAttributes) GetOKTA_GROUPS_FILTER() string`

GetOKTA_GROUPS_FILTER returns the OKTA_GROUPS_FILTER field if non-nil, zero value otherwise.

### GetOKTA_GROUPS_FILTEROk

`func (o *OktaConnectionAttributes) GetOKTA_GROUPS_FILTEROk() (*string, bool)`

GetOKTA_GROUPS_FILTEROk returns a tuple with the OKTA_GROUPS_FILTER field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOKTA_GROUPS_FILTER

`func (o *OktaConnectionAttributes) SetOKTA_GROUPS_FILTER(v string)`

SetOKTA_GROUPS_FILTER sets OKTA_GROUPS_FILTER field to given value.

### HasOKTA_GROUPS_FILTER

`func (o *OktaConnectionAttributes) HasOKTA_GROUPS_FILTER() bool`

HasOKTA_GROUPS_FILTER returns a boolean if a field has been set.

### GetACCOUNTFIELDMAPPINGS

`func (o *OktaConnectionAttributes) GetACCOUNTFIELDMAPPINGS() string`

GetACCOUNTFIELDMAPPINGS returns the ACCOUNTFIELDMAPPINGS field if non-nil, zero value otherwise.

### GetACCOUNTFIELDMAPPINGSOk

`func (o *OktaConnectionAttributes) GetACCOUNTFIELDMAPPINGSOk() (*string, bool)`

GetACCOUNTFIELDMAPPINGSOk returns a tuple with the ACCOUNTFIELDMAPPINGS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetACCOUNTFIELDMAPPINGS

`func (o *OktaConnectionAttributes) SetACCOUNTFIELDMAPPINGS(v string)`

SetACCOUNTFIELDMAPPINGS sets ACCOUNTFIELDMAPPINGS field to given value.

### HasACCOUNTFIELDMAPPINGS

`func (o *OktaConnectionAttributes) HasACCOUNTFIELDMAPPINGS() bool`

HasACCOUNTFIELDMAPPINGS returns a boolean if a field has been set.

### GetIMPORT_INACTIVE_APPS

`func (o *OktaConnectionAttributes) GetIMPORT_INACTIVE_APPS() string`

GetIMPORT_INACTIVE_APPS returns the IMPORT_INACTIVE_APPS field if non-nil, zero value otherwise.

### GetIMPORT_INACTIVE_APPSOk

`func (o *OktaConnectionAttributes) GetIMPORT_INACTIVE_APPSOk() (*string, bool)`

GetIMPORT_INACTIVE_APPSOk returns a tuple with the IMPORT_INACTIVE_APPS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIMPORT_INACTIVE_APPS

`func (o *OktaConnectionAttributes) SetIMPORT_INACTIVE_APPS(v string)`

SetIMPORT_INACTIVE_APPS sets IMPORT_INACTIVE_APPS field to given value.

### HasIMPORT_INACTIVE_APPS

`func (o *OktaConnectionAttributes) HasIMPORT_INACTIVE_APPS() bool`

HasIMPORT_INACTIVE_APPS returns a boolean if a field has been set.

### GetAUDIT_FILTER

`func (o *OktaConnectionAttributes) GetAUDIT_FILTER() string`

GetAUDIT_FILTER returns the AUDIT_FILTER field if non-nil, zero value otherwise.

### GetAUDIT_FILTEROk

`func (o *OktaConnectionAttributes) GetAUDIT_FILTEROk() (*string, bool)`

GetAUDIT_FILTEROk returns a tuple with the AUDIT_FILTER field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAUDIT_FILTER

`func (o *OktaConnectionAttributes) SetAUDIT_FILTER(v string)`

SetAUDIT_FILTER sets AUDIT_FILTER field to given value.

### HasAUDIT_FILTER

`func (o *OktaConnectionAttributes) HasAUDIT_FILTER() bool`

HasAUDIT_FILTER returns a boolean if a field has been set.

### GetUSERFIELDMAPPINGS

`func (o *OktaConnectionAttributes) GetUSERFIELDMAPPINGS() string`

GetUSERFIELDMAPPINGS returns the USERFIELDMAPPINGS field if non-nil, zero value otherwise.

### GetUSERFIELDMAPPINGSOk

`func (o *OktaConnectionAttributes) GetUSERFIELDMAPPINGSOk() (*string, bool)`

GetUSERFIELDMAPPINGSOk returns a tuple with the USERFIELDMAPPINGS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUSERFIELDMAPPINGS

`func (o *OktaConnectionAttributes) SetUSERFIELDMAPPINGS(v string)`

SetUSERFIELDMAPPINGS sets USERFIELDMAPPINGS field to given value.

### HasUSERFIELDMAPPINGS

`func (o *OktaConnectionAttributes) HasUSERFIELDMAPPINGS() bool`

HasUSERFIELDMAPPINGS returns a boolean if a field has been set.

### GetAPPACCOUNTFIELDMAPPINGS

`func (o *OktaConnectionAttributes) GetAPPACCOUNTFIELDMAPPINGS() string`

GetAPPACCOUNTFIELDMAPPINGS returns the APPACCOUNTFIELDMAPPINGS field if non-nil, zero value otherwise.

### GetAPPACCOUNTFIELDMAPPINGSOk

`func (o *OktaConnectionAttributes) GetAPPACCOUNTFIELDMAPPINGSOk() (*string, bool)`

GetAPPACCOUNTFIELDMAPPINGSOk returns a tuple with the APPACCOUNTFIELDMAPPINGS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAPPACCOUNTFIELDMAPPINGS

`func (o *OktaConnectionAttributes) SetAPPACCOUNTFIELDMAPPINGS(v string)`

SetAPPACCOUNTFIELDMAPPINGS sets APPACCOUNTFIELDMAPPINGS field to given value.

### HasAPPACCOUNTFIELDMAPPINGS

`func (o *OktaConnectionAttributes) HasAPPACCOUNTFIELDMAPPINGS() bool`

HasAPPACCOUNTFIELDMAPPINGS returns a boolean if a field has been set.

### GetMODIFYUSERDATAJSON

`func (o *OktaConnectionAttributes) GetMODIFYUSERDATAJSON() string`

GetMODIFYUSERDATAJSON returns the MODIFYUSERDATAJSON field if non-nil, zero value otherwise.

### GetMODIFYUSERDATAJSONOk

`func (o *OktaConnectionAttributes) GetMODIFYUSERDATAJSONOk() (*string, bool)`

GetMODIFYUSERDATAJSONOk returns a tuple with the MODIFYUSERDATAJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMODIFYUSERDATAJSON

`func (o *OktaConnectionAttributes) SetMODIFYUSERDATAJSON(v string)`

SetMODIFYUSERDATAJSON sets MODIFYUSERDATAJSON field to given value.

### HasMODIFYUSERDATAJSON

`func (o *OktaConnectionAttributes) HasMODIFYUSERDATAJSON() bool`

HasMODIFYUSERDATAJSON returns a boolean if a field has been set.

### GetConnectionTimeoutConfig

`func (o *OktaConnectionAttributes) GetConnectionTimeoutConfig() ConnectionTimeoutConfig`

GetConnectionTimeoutConfig returns the ConnectionTimeoutConfig field if non-nil, zero value otherwise.

### GetConnectionTimeoutConfigOk

`func (o *OktaConnectionAttributes) GetConnectionTimeoutConfigOk() (*ConnectionTimeoutConfig, bool)`

GetConnectionTimeoutConfigOk returns a tuple with the ConnectionTimeoutConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionTimeoutConfig

`func (o *OktaConnectionAttributes) SetConnectionTimeoutConfig(v ConnectionTimeoutConfig)`

SetConnectionTimeoutConfig sets ConnectionTimeoutConfig field to given value.

### HasConnectionTimeoutConfig

`func (o *OktaConnectionAttributes) HasConnectionTimeoutConfig() bool`

HasConnectionTimeoutConfig returns a boolean if a field has been set.

### GetConnectionType

`func (o *OktaConnectionAttributes) GetConnectionType() string`

GetConnectionType returns the ConnectionType field if non-nil, zero value otherwise.

### GetConnectionTypeOk

`func (o *OktaConnectionAttributes) GetConnectionTypeOk() (*string, bool)`

GetConnectionTypeOk returns a tuple with the ConnectionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionType

`func (o *OktaConnectionAttributes) SetConnectionType(v string)`

SetConnectionType sets ConnectionType field to given value.

### HasConnectionType

`func (o *OktaConnectionAttributes) HasConnectionType() bool`

HasConnectionType returns a boolean if a field has been set.

### GetIsTimeoutConfigValidated

`func (o *OktaConnectionAttributes) GetIsTimeoutConfigValidated() bool`

GetIsTimeoutConfigValidated returns the IsTimeoutConfigValidated field if non-nil, zero value otherwise.

### GetIsTimeoutConfigValidatedOk

`func (o *OktaConnectionAttributes) GetIsTimeoutConfigValidatedOk() (*bool, bool)`

GetIsTimeoutConfigValidatedOk returns a tuple with the IsTimeoutConfigValidated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsTimeoutConfigValidated

`func (o *OktaConnectionAttributes) SetIsTimeoutConfigValidated(v bool)`

SetIsTimeoutConfigValidated sets IsTimeoutConfigValidated field to given value.

### HasIsTimeoutConfigValidated

`func (o *OktaConnectionAttributes) HasIsTimeoutConfigValidated() bool`

HasIsTimeoutConfigValidated returns a boolean if a field has been set.

### GetACTIVATE_ENDPOINT

`func (o *OktaConnectionAttributes) GetACTIVATE_ENDPOINT() string`

GetACTIVATE_ENDPOINT returns the ACTIVATE_ENDPOINT field if non-nil, zero value otherwise.

### GetACTIVATE_ENDPOINTOk

`func (o *OktaConnectionAttributes) GetACTIVATE_ENDPOINTOk() (*string, bool)`

GetACTIVATE_ENDPOINTOk returns a tuple with the ACTIVATE_ENDPOINT field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetACTIVATE_ENDPOINT

`func (o *OktaConnectionAttributes) SetACTIVATE_ENDPOINT(v string)`

SetACTIVATE_ENDPOINT sets ACTIVATE_ENDPOINT field to given value.

### HasACTIVATE_ENDPOINT

`func (o *OktaConnectionAttributes) HasACTIVATE_ENDPOINT() bool`

HasACTIVATE_ENDPOINT returns a boolean if a field has been set.

### GetENTITLEMENTTYPESMAPPINGS

`func (o *OktaConnectionAttributes) GetENTITLEMENTTYPESMAPPINGS() string`

GetENTITLEMENTTYPESMAPPINGS returns the ENTITLEMENTTYPESMAPPINGS field if non-nil, zero value otherwise.

### GetENTITLEMENTTYPESMAPPINGSOk

`func (o *OktaConnectionAttributes) GetENTITLEMENTTYPESMAPPINGSOk() (*string, bool)`

GetENTITLEMENTTYPESMAPPINGSOk returns a tuple with the ENTITLEMENTTYPESMAPPINGS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetENTITLEMENTTYPESMAPPINGS

`func (o *OktaConnectionAttributes) SetENTITLEMENTTYPESMAPPINGS(v string)`

SetENTITLEMENTTYPESMAPPINGS sets ENTITLEMENTTYPESMAPPINGS field to given value.

### HasENTITLEMENTTYPESMAPPINGS

`func (o *OktaConnectionAttributes) HasENTITLEMENTTYPESMAPPINGS() bool`

HasENTITLEMENTTYPESMAPPINGS returns a boolean if a field has been set.

### GetOKTA_APPLICATION_SECURITYSYSTEM

`func (o *OktaConnectionAttributes) GetOKTA_APPLICATION_SECURITYSYSTEM() string`

GetOKTA_APPLICATION_SECURITYSYSTEM returns the OKTA_APPLICATION_SECURITYSYSTEM field if non-nil, zero value otherwise.

### GetOKTA_APPLICATION_SECURITYSYSTEMOk

`func (o *OktaConnectionAttributes) GetOKTA_APPLICATION_SECURITYSYSTEMOk() (*string, bool)`

GetOKTA_APPLICATION_SECURITYSYSTEMOk returns a tuple with the OKTA_APPLICATION_SECURITYSYSTEM field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOKTA_APPLICATION_SECURITYSYSTEM

`func (o *OktaConnectionAttributes) SetOKTA_APPLICATION_SECURITYSYSTEM(v string)`

SetOKTA_APPLICATION_SECURITYSYSTEM sets OKTA_APPLICATION_SECURITYSYSTEM field to given value.

### HasOKTA_APPLICATION_SECURITYSYSTEM

`func (o *OktaConnectionAttributes) HasOKTA_APPLICATION_SECURITYSYSTEM() bool`

HasOKTA_APPLICATION_SECURITYSYSTEM returns a boolean if a field has been set.

### GetPAM_CONFIG

`func (o *OktaConnectionAttributes) GetPAM_CONFIG() string`

GetPAM_CONFIG returns the PAM_CONFIG field if non-nil, zero value otherwise.

### GetPAM_CONFIGOk

`func (o *OktaConnectionAttributes) GetPAM_CONFIGOk() (*string, bool)`

GetPAM_CONFIGOk returns a tuple with the PAM_CONFIG field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPAM_CONFIG

`func (o *OktaConnectionAttributes) SetPAM_CONFIG(v string)`

SetPAM_CONFIG sets PAM_CONFIG field to given value.

### HasPAM_CONFIG

`func (o *OktaConnectionAttributes) HasPAM_CONFIG() bool`

HasPAM_CONFIG returns a boolean if a field has been set.

### GetConfigJSON

`func (o *OktaConnectionAttributes) GetConfigJSON() string`

GetConfigJSON returns the ConfigJSON field if non-nil, zero value otherwise.

### GetConfigJSONOk

`func (o *OktaConnectionAttributes) GetConfigJSONOk() (*string, bool)`

GetConfigJSONOk returns a tuple with the ConfigJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigJSON

`func (o *OktaConnectionAttributes) SetConfigJSON(v string)`

SetConfigJSON sets ConfigJSON field to given value.

### HasConfigJSON

`func (o *OktaConnectionAttributes) HasConfigJSON() bool`

HasConfigJSON returns a boolean if a field has been set.

### GetSTATUS_THRESHOLD_CONFIG

`func (o *OktaConnectionAttributes) GetSTATUS_THRESHOLD_CONFIG() string`

GetSTATUS_THRESHOLD_CONFIG returns the STATUS_THRESHOLD_CONFIG field if non-nil, zero value otherwise.

### GetSTATUS_THRESHOLD_CONFIGOk

`func (o *OktaConnectionAttributes) GetSTATUS_THRESHOLD_CONFIGOk() (*string, bool)`

GetSTATUS_THRESHOLD_CONFIGOk returns a tuple with the STATUS_THRESHOLD_CONFIG field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSTATUS_THRESHOLD_CONFIG

`func (o *OktaConnectionAttributes) SetSTATUS_THRESHOLD_CONFIG(v string)`

SetSTATUS_THRESHOLD_CONFIG sets STATUS_THRESHOLD_CONFIG field to given value.

### HasSTATUS_THRESHOLD_CONFIG

`func (o *OktaConnectionAttributes) HasSTATUS_THRESHOLD_CONFIG() bool`

HasSTATUS_THRESHOLD_CONFIG returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


