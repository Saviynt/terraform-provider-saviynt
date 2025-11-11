# OktaConnector

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IMPORTURL** | **string** |  | 
**AUTHTOKEN** | **string** |  | 
**ACCOUNTFIELDMAPPINGS** | Pointer to **string** |  | [optional] 
**USERFIELDMAPPINGS** | Pointer to **string** |  | [optional] 
**ENTITLEMENTTYPESMAPPINGS** | Pointer to **string** |  | [optional] 
**IMPORT_INACTIVE_APPS** | Pointer to **string** |  | [optional] 
**OKTA_APPLICATION_SECURITYSYSTEM** | **string** |  | 
**OKTA_GROUPS_FILTER** | Pointer to **string** |  | [optional] 
**APPACCOUNTFIELDMAPPINGS** | Pointer to **string** |  | [optional] 
**STATUS_THRESHOLD_CONFIG** | Pointer to **string** |  | [optional] 
**AUDIT_FILTER** | Pointer to **string** |  | [optional] 
**MODIFYUSERDATAJSON** | Pointer to **string** |  | [optional] 
**ACTIVATE_ENDPOINT** | Pointer to **string** |  | [optional] 
**ConfigJSON** | Pointer to **string** |  | [optional] 
**PAM_CONFIG** | Pointer to **string** |  | [optional] 

## Methods

### NewOktaConnector

`func NewOktaConnector(iMPORTURL string, aUTHTOKEN string, oKTAAPPLICATIONSECURITYSYSTEM string, ) *OktaConnector`

NewOktaConnector instantiates a new OktaConnector object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOktaConnectorWithDefaults

`func NewOktaConnectorWithDefaults() *OktaConnector`

NewOktaConnectorWithDefaults instantiates a new OktaConnector object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIMPORTURL

`func (o *OktaConnector) GetIMPORTURL() string`

GetIMPORTURL returns the IMPORTURL field if non-nil, zero value otherwise.

### GetIMPORTURLOk

`func (o *OktaConnector) GetIMPORTURLOk() (*string, bool)`

GetIMPORTURLOk returns a tuple with the IMPORTURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIMPORTURL

`func (o *OktaConnector) SetIMPORTURL(v string)`

SetIMPORTURL sets IMPORTURL field to given value.


### GetAUTHTOKEN

`func (o *OktaConnector) GetAUTHTOKEN() string`

GetAUTHTOKEN returns the AUTHTOKEN field if non-nil, zero value otherwise.

### GetAUTHTOKENOk

`func (o *OktaConnector) GetAUTHTOKENOk() (*string, bool)`

GetAUTHTOKENOk returns a tuple with the AUTHTOKEN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAUTHTOKEN

`func (o *OktaConnector) SetAUTHTOKEN(v string)`

SetAUTHTOKEN sets AUTHTOKEN field to given value.


### GetACCOUNTFIELDMAPPINGS

`func (o *OktaConnector) GetACCOUNTFIELDMAPPINGS() string`

GetACCOUNTFIELDMAPPINGS returns the ACCOUNTFIELDMAPPINGS field if non-nil, zero value otherwise.

### GetACCOUNTFIELDMAPPINGSOk

`func (o *OktaConnector) GetACCOUNTFIELDMAPPINGSOk() (*string, bool)`

GetACCOUNTFIELDMAPPINGSOk returns a tuple with the ACCOUNTFIELDMAPPINGS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetACCOUNTFIELDMAPPINGS

`func (o *OktaConnector) SetACCOUNTFIELDMAPPINGS(v string)`

SetACCOUNTFIELDMAPPINGS sets ACCOUNTFIELDMAPPINGS field to given value.

### HasACCOUNTFIELDMAPPINGS

`func (o *OktaConnector) HasACCOUNTFIELDMAPPINGS() bool`

HasACCOUNTFIELDMAPPINGS returns a boolean if a field has been set.

### GetUSERFIELDMAPPINGS

`func (o *OktaConnector) GetUSERFIELDMAPPINGS() string`

GetUSERFIELDMAPPINGS returns the USERFIELDMAPPINGS field if non-nil, zero value otherwise.

### GetUSERFIELDMAPPINGSOk

`func (o *OktaConnector) GetUSERFIELDMAPPINGSOk() (*string, bool)`

GetUSERFIELDMAPPINGSOk returns a tuple with the USERFIELDMAPPINGS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUSERFIELDMAPPINGS

`func (o *OktaConnector) SetUSERFIELDMAPPINGS(v string)`

SetUSERFIELDMAPPINGS sets USERFIELDMAPPINGS field to given value.

### HasUSERFIELDMAPPINGS

`func (o *OktaConnector) HasUSERFIELDMAPPINGS() bool`

HasUSERFIELDMAPPINGS returns a boolean if a field has been set.

### GetENTITLEMENTTYPESMAPPINGS

`func (o *OktaConnector) GetENTITLEMENTTYPESMAPPINGS() string`

GetENTITLEMENTTYPESMAPPINGS returns the ENTITLEMENTTYPESMAPPINGS field if non-nil, zero value otherwise.

### GetENTITLEMENTTYPESMAPPINGSOk

`func (o *OktaConnector) GetENTITLEMENTTYPESMAPPINGSOk() (*string, bool)`

GetENTITLEMENTTYPESMAPPINGSOk returns a tuple with the ENTITLEMENTTYPESMAPPINGS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetENTITLEMENTTYPESMAPPINGS

`func (o *OktaConnector) SetENTITLEMENTTYPESMAPPINGS(v string)`

SetENTITLEMENTTYPESMAPPINGS sets ENTITLEMENTTYPESMAPPINGS field to given value.

### HasENTITLEMENTTYPESMAPPINGS

`func (o *OktaConnector) HasENTITLEMENTTYPESMAPPINGS() bool`

HasENTITLEMENTTYPESMAPPINGS returns a boolean if a field has been set.

### GetIMPORT_INACTIVE_APPS

`func (o *OktaConnector) GetIMPORT_INACTIVE_APPS() string`

GetIMPORT_INACTIVE_APPS returns the IMPORT_INACTIVE_APPS field if non-nil, zero value otherwise.

### GetIMPORT_INACTIVE_APPSOk

`func (o *OktaConnector) GetIMPORT_INACTIVE_APPSOk() (*string, bool)`

GetIMPORT_INACTIVE_APPSOk returns a tuple with the IMPORT_INACTIVE_APPS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIMPORT_INACTIVE_APPS

`func (o *OktaConnector) SetIMPORT_INACTIVE_APPS(v string)`

SetIMPORT_INACTIVE_APPS sets IMPORT_INACTIVE_APPS field to given value.

### HasIMPORT_INACTIVE_APPS

`func (o *OktaConnector) HasIMPORT_INACTIVE_APPS() bool`

HasIMPORT_INACTIVE_APPS returns a boolean if a field has been set.

### GetOKTA_APPLICATION_SECURITYSYSTEM

`func (o *OktaConnector) GetOKTA_APPLICATION_SECURITYSYSTEM() string`

GetOKTA_APPLICATION_SECURITYSYSTEM returns the OKTA_APPLICATION_SECURITYSYSTEM field if non-nil, zero value otherwise.

### GetOKTA_APPLICATION_SECURITYSYSTEMOk

`func (o *OktaConnector) GetOKTA_APPLICATION_SECURITYSYSTEMOk() (*string, bool)`

GetOKTA_APPLICATION_SECURITYSYSTEMOk returns a tuple with the OKTA_APPLICATION_SECURITYSYSTEM field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOKTA_APPLICATION_SECURITYSYSTEM

`func (o *OktaConnector) SetOKTA_APPLICATION_SECURITYSYSTEM(v string)`

SetOKTA_APPLICATION_SECURITYSYSTEM sets OKTA_APPLICATION_SECURITYSYSTEM field to given value.


### GetOKTA_GROUPS_FILTER

`func (o *OktaConnector) GetOKTA_GROUPS_FILTER() string`

GetOKTA_GROUPS_FILTER returns the OKTA_GROUPS_FILTER field if non-nil, zero value otherwise.

### GetOKTA_GROUPS_FILTEROk

`func (o *OktaConnector) GetOKTA_GROUPS_FILTEROk() (*string, bool)`

GetOKTA_GROUPS_FILTEROk returns a tuple with the OKTA_GROUPS_FILTER field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOKTA_GROUPS_FILTER

`func (o *OktaConnector) SetOKTA_GROUPS_FILTER(v string)`

SetOKTA_GROUPS_FILTER sets OKTA_GROUPS_FILTER field to given value.

### HasOKTA_GROUPS_FILTER

`func (o *OktaConnector) HasOKTA_GROUPS_FILTER() bool`

HasOKTA_GROUPS_FILTER returns a boolean if a field has been set.

### GetAPPACCOUNTFIELDMAPPINGS

`func (o *OktaConnector) GetAPPACCOUNTFIELDMAPPINGS() string`

GetAPPACCOUNTFIELDMAPPINGS returns the APPACCOUNTFIELDMAPPINGS field if non-nil, zero value otherwise.

### GetAPPACCOUNTFIELDMAPPINGSOk

`func (o *OktaConnector) GetAPPACCOUNTFIELDMAPPINGSOk() (*string, bool)`

GetAPPACCOUNTFIELDMAPPINGSOk returns a tuple with the APPACCOUNTFIELDMAPPINGS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAPPACCOUNTFIELDMAPPINGS

`func (o *OktaConnector) SetAPPACCOUNTFIELDMAPPINGS(v string)`

SetAPPACCOUNTFIELDMAPPINGS sets APPACCOUNTFIELDMAPPINGS field to given value.

### HasAPPACCOUNTFIELDMAPPINGS

`func (o *OktaConnector) HasAPPACCOUNTFIELDMAPPINGS() bool`

HasAPPACCOUNTFIELDMAPPINGS returns a boolean if a field has been set.

### GetSTATUS_THRESHOLD_CONFIG

`func (o *OktaConnector) GetSTATUS_THRESHOLD_CONFIG() string`

GetSTATUS_THRESHOLD_CONFIG returns the STATUS_THRESHOLD_CONFIG field if non-nil, zero value otherwise.

### GetSTATUS_THRESHOLD_CONFIGOk

`func (o *OktaConnector) GetSTATUS_THRESHOLD_CONFIGOk() (*string, bool)`

GetSTATUS_THRESHOLD_CONFIGOk returns a tuple with the STATUS_THRESHOLD_CONFIG field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSTATUS_THRESHOLD_CONFIG

`func (o *OktaConnector) SetSTATUS_THRESHOLD_CONFIG(v string)`

SetSTATUS_THRESHOLD_CONFIG sets STATUS_THRESHOLD_CONFIG field to given value.

### HasSTATUS_THRESHOLD_CONFIG

`func (o *OktaConnector) HasSTATUS_THRESHOLD_CONFIG() bool`

HasSTATUS_THRESHOLD_CONFIG returns a boolean if a field has been set.

### GetAUDIT_FILTER

`func (o *OktaConnector) GetAUDIT_FILTER() string`

GetAUDIT_FILTER returns the AUDIT_FILTER field if non-nil, zero value otherwise.

### GetAUDIT_FILTEROk

`func (o *OktaConnector) GetAUDIT_FILTEROk() (*string, bool)`

GetAUDIT_FILTEROk returns a tuple with the AUDIT_FILTER field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAUDIT_FILTER

`func (o *OktaConnector) SetAUDIT_FILTER(v string)`

SetAUDIT_FILTER sets AUDIT_FILTER field to given value.

### HasAUDIT_FILTER

`func (o *OktaConnector) HasAUDIT_FILTER() bool`

HasAUDIT_FILTER returns a boolean if a field has been set.

### GetMODIFYUSERDATAJSON

`func (o *OktaConnector) GetMODIFYUSERDATAJSON() string`

GetMODIFYUSERDATAJSON returns the MODIFYUSERDATAJSON field if non-nil, zero value otherwise.

### GetMODIFYUSERDATAJSONOk

`func (o *OktaConnector) GetMODIFYUSERDATAJSONOk() (*string, bool)`

GetMODIFYUSERDATAJSONOk returns a tuple with the MODIFYUSERDATAJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMODIFYUSERDATAJSON

`func (o *OktaConnector) SetMODIFYUSERDATAJSON(v string)`

SetMODIFYUSERDATAJSON sets MODIFYUSERDATAJSON field to given value.

### HasMODIFYUSERDATAJSON

`func (o *OktaConnector) HasMODIFYUSERDATAJSON() bool`

HasMODIFYUSERDATAJSON returns a boolean if a field has been set.

### GetACTIVATE_ENDPOINT

`func (o *OktaConnector) GetACTIVATE_ENDPOINT() string`

GetACTIVATE_ENDPOINT returns the ACTIVATE_ENDPOINT field if non-nil, zero value otherwise.

### GetACTIVATE_ENDPOINTOk

`func (o *OktaConnector) GetACTIVATE_ENDPOINTOk() (*string, bool)`

GetACTIVATE_ENDPOINTOk returns a tuple with the ACTIVATE_ENDPOINT field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetACTIVATE_ENDPOINT

`func (o *OktaConnector) SetACTIVATE_ENDPOINT(v string)`

SetACTIVATE_ENDPOINT sets ACTIVATE_ENDPOINT field to given value.

### HasACTIVATE_ENDPOINT

`func (o *OktaConnector) HasACTIVATE_ENDPOINT() bool`

HasACTIVATE_ENDPOINT returns a boolean if a field has been set.

### GetConfigJSON

`func (o *OktaConnector) GetConfigJSON() string`

GetConfigJSON returns the ConfigJSON field if non-nil, zero value otherwise.

### GetConfigJSONOk

`func (o *OktaConnector) GetConfigJSONOk() (*string, bool)`

GetConfigJSONOk returns a tuple with the ConfigJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigJSON

`func (o *OktaConnector) SetConfigJSON(v string)`

SetConfigJSON sets ConfigJSON field to given value.

### HasConfigJSON

`func (o *OktaConnector) HasConfigJSON() bool`

HasConfigJSON returns a boolean if a field has been set.

### GetPAM_CONFIG

`func (o *OktaConnector) GetPAM_CONFIG() string`

GetPAM_CONFIG returns the PAM_CONFIG field if non-nil, zero value otherwise.

### GetPAM_CONFIGOk

`func (o *OktaConnector) GetPAM_CONFIGOk() (*string, bool)`

GetPAM_CONFIGOk returns a tuple with the PAM_CONFIG field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPAM_CONFIG

`func (o *OktaConnector) SetPAM_CONFIG(v string)`

SetPAM_CONFIG sets PAM_CONFIG field to given value.

### HasPAM_CONFIG

`func (o *OktaConnector) HasPAM_CONFIG() bool`

HasPAM_CONFIG returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


