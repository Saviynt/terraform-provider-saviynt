# SFTPConnectionAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HOST_NAME** | Pointer to **string** |  | [optional] 
**PORT_NUMBER** | Pointer to **string** |  | [optional] 
**USERNAME** | Pointer to **string** |  | [optional] 
**AUTH_CREDENTIAL_TYPE** | Pointer to **string** |  | [optional] 
**AUTH_CREDENTIAL_VALUE** | Pointer to **string** |  | [optional] 
**PASSPHRASE** | Pointer to **string** |  | [optional] 
**FILES_TO_GET** | Pointer to **string** |  | [optional] 
**FILES_TO_PUT** | Pointer to **string** |  | [optional] 
**PAM_CONFIG** | Pointer to **string** |  | [optional] 
**ConnectionType** | Pointer to **string** |  | [optional] 
**IsTimeoutSupported** | Pointer to **bool** |  | [optional] 
**IsTimeoutConfigValidated** | Pointer to **bool** |  | [optional] 
**ConnectionTimeoutConfig** | Pointer to [**ConnectionTimeoutConfig**](ConnectionTimeoutConfig.md) |  | [optional] 

## Methods

### NewSFTPConnectionAttributes

`func NewSFTPConnectionAttributes() *SFTPConnectionAttributes`

NewSFTPConnectionAttributes instantiates a new SFTPConnectionAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSFTPConnectionAttributesWithDefaults

`func NewSFTPConnectionAttributesWithDefaults() *SFTPConnectionAttributes`

NewSFTPConnectionAttributesWithDefaults instantiates a new SFTPConnectionAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHOST_NAME

`func (o *SFTPConnectionAttributes) GetHOST_NAME() string`

GetHOST_NAME returns the HOST_NAME field if non-nil, zero value otherwise.

### GetHOST_NAMEOk

`func (o *SFTPConnectionAttributes) GetHOST_NAMEOk() (*string, bool)`

GetHOST_NAMEOk returns a tuple with the HOST_NAME field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHOST_NAME

`func (o *SFTPConnectionAttributes) SetHOST_NAME(v string)`

SetHOST_NAME sets HOST_NAME field to given value.

### HasHOST_NAME

`func (o *SFTPConnectionAttributes) HasHOST_NAME() bool`

HasHOST_NAME returns a boolean if a field has been set.

### GetPORT_NUMBER

`func (o *SFTPConnectionAttributes) GetPORT_NUMBER() string`

GetPORT_NUMBER returns the PORT_NUMBER field if non-nil, zero value otherwise.

### GetPORT_NUMBEROk

`func (o *SFTPConnectionAttributes) GetPORT_NUMBEROk() (*string, bool)`

GetPORT_NUMBEROk returns a tuple with the PORT_NUMBER field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPORT_NUMBER

`func (o *SFTPConnectionAttributes) SetPORT_NUMBER(v string)`

SetPORT_NUMBER sets PORT_NUMBER field to given value.

### HasPORT_NUMBER

`func (o *SFTPConnectionAttributes) HasPORT_NUMBER() bool`

HasPORT_NUMBER returns a boolean if a field has been set.

### GetUSERNAME

`func (o *SFTPConnectionAttributes) GetUSERNAME() string`

GetUSERNAME returns the USERNAME field if non-nil, zero value otherwise.

### GetUSERNAMEOk

`func (o *SFTPConnectionAttributes) GetUSERNAMEOk() (*string, bool)`

GetUSERNAMEOk returns a tuple with the USERNAME field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUSERNAME

`func (o *SFTPConnectionAttributes) SetUSERNAME(v string)`

SetUSERNAME sets USERNAME field to given value.

### HasUSERNAME

`func (o *SFTPConnectionAttributes) HasUSERNAME() bool`

HasUSERNAME returns a boolean if a field has been set.

### GetAUTH_CREDENTIAL_TYPE

`func (o *SFTPConnectionAttributes) GetAUTH_CREDENTIAL_TYPE() string`

GetAUTH_CREDENTIAL_TYPE returns the AUTH_CREDENTIAL_TYPE field if non-nil, zero value otherwise.

### GetAUTH_CREDENTIAL_TYPEOk

`func (o *SFTPConnectionAttributes) GetAUTH_CREDENTIAL_TYPEOk() (*string, bool)`

GetAUTH_CREDENTIAL_TYPEOk returns a tuple with the AUTH_CREDENTIAL_TYPE field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAUTH_CREDENTIAL_TYPE

`func (o *SFTPConnectionAttributes) SetAUTH_CREDENTIAL_TYPE(v string)`

SetAUTH_CREDENTIAL_TYPE sets AUTH_CREDENTIAL_TYPE field to given value.

### HasAUTH_CREDENTIAL_TYPE

`func (o *SFTPConnectionAttributes) HasAUTH_CREDENTIAL_TYPE() bool`

HasAUTH_CREDENTIAL_TYPE returns a boolean if a field has been set.

### GetAUTH_CREDENTIAL_VALUE

`func (o *SFTPConnectionAttributes) GetAUTH_CREDENTIAL_VALUE() string`

GetAUTH_CREDENTIAL_VALUE returns the AUTH_CREDENTIAL_VALUE field if non-nil, zero value otherwise.

### GetAUTH_CREDENTIAL_VALUEOk

`func (o *SFTPConnectionAttributes) GetAUTH_CREDENTIAL_VALUEOk() (*string, bool)`

GetAUTH_CREDENTIAL_VALUEOk returns a tuple with the AUTH_CREDENTIAL_VALUE field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAUTH_CREDENTIAL_VALUE

`func (o *SFTPConnectionAttributes) SetAUTH_CREDENTIAL_VALUE(v string)`

SetAUTH_CREDENTIAL_VALUE sets AUTH_CREDENTIAL_VALUE field to given value.

### HasAUTH_CREDENTIAL_VALUE

`func (o *SFTPConnectionAttributes) HasAUTH_CREDENTIAL_VALUE() bool`

HasAUTH_CREDENTIAL_VALUE returns a boolean if a field has been set.

### GetPASSPHRASE

`func (o *SFTPConnectionAttributes) GetPASSPHRASE() string`

GetPASSPHRASE returns the PASSPHRASE field if non-nil, zero value otherwise.

### GetPASSPHRASEOk

`func (o *SFTPConnectionAttributes) GetPASSPHRASEOk() (*string, bool)`

GetPASSPHRASEOk returns a tuple with the PASSPHRASE field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPASSPHRASE

`func (o *SFTPConnectionAttributes) SetPASSPHRASE(v string)`

SetPASSPHRASE sets PASSPHRASE field to given value.

### HasPASSPHRASE

`func (o *SFTPConnectionAttributes) HasPASSPHRASE() bool`

HasPASSPHRASE returns a boolean if a field has been set.

### GetFILES_TO_GET

`func (o *SFTPConnectionAttributes) GetFILES_TO_GET() string`

GetFILES_TO_GET returns the FILES_TO_GET field if non-nil, zero value otherwise.

### GetFILES_TO_GETOk

`func (o *SFTPConnectionAttributes) GetFILES_TO_GETOk() (*string, bool)`

GetFILES_TO_GETOk returns a tuple with the FILES_TO_GET field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFILES_TO_GET

`func (o *SFTPConnectionAttributes) SetFILES_TO_GET(v string)`

SetFILES_TO_GET sets FILES_TO_GET field to given value.

### HasFILES_TO_GET

`func (o *SFTPConnectionAttributes) HasFILES_TO_GET() bool`

HasFILES_TO_GET returns a boolean if a field has been set.

### GetFILES_TO_PUT

`func (o *SFTPConnectionAttributes) GetFILES_TO_PUT() string`

GetFILES_TO_PUT returns the FILES_TO_PUT field if non-nil, zero value otherwise.

### GetFILES_TO_PUTOk

`func (o *SFTPConnectionAttributes) GetFILES_TO_PUTOk() (*string, bool)`

GetFILES_TO_PUTOk returns a tuple with the FILES_TO_PUT field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFILES_TO_PUT

`func (o *SFTPConnectionAttributes) SetFILES_TO_PUT(v string)`

SetFILES_TO_PUT sets FILES_TO_PUT field to given value.

### HasFILES_TO_PUT

`func (o *SFTPConnectionAttributes) HasFILES_TO_PUT() bool`

HasFILES_TO_PUT returns a boolean if a field has been set.

### GetPAM_CONFIG

`func (o *SFTPConnectionAttributes) GetPAM_CONFIG() string`

GetPAM_CONFIG returns the PAM_CONFIG field if non-nil, zero value otherwise.

### GetPAM_CONFIGOk

`func (o *SFTPConnectionAttributes) GetPAM_CONFIGOk() (*string, bool)`

GetPAM_CONFIGOk returns a tuple with the PAM_CONFIG field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPAM_CONFIG

`func (o *SFTPConnectionAttributes) SetPAM_CONFIG(v string)`

SetPAM_CONFIG sets PAM_CONFIG field to given value.

### HasPAM_CONFIG

`func (o *SFTPConnectionAttributes) HasPAM_CONFIG() bool`

HasPAM_CONFIG returns a boolean if a field has been set.

### GetConnectionType

`func (o *SFTPConnectionAttributes) GetConnectionType() string`

GetConnectionType returns the ConnectionType field if non-nil, zero value otherwise.

### GetConnectionTypeOk

`func (o *SFTPConnectionAttributes) GetConnectionTypeOk() (*string, bool)`

GetConnectionTypeOk returns a tuple with the ConnectionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionType

`func (o *SFTPConnectionAttributes) SetConnectionType(v string)`

SetConnectionType sets ConnectionType field to given value.

### HasConnectionType

`func (o *SFTPConnectionAttributes) HasConnectionType() bool`

HasConnectionType returns a boolean if a field has been set.

### GetIsTimeoutSupported

`func (o *SFTPConnectionAttributes) GetIsTimeoutSupported() bool`

GetIsTimeoutSupported returns the IsTimeoutSupported field if non-nil, zero value otherwise.

### GetIsTimeoutSupportedOk

`func (o *SFTPConnectionAttributes) GetIsTimeoutSupportedOk() (*bool, bool)`

GetIsTimeoutSupportedOk returns a tuple with the IsTimeoutSupported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsTimeoutSupported

`func (o *SFTPConnectionAttributes) SetIsTimeoutSupported(v bool)`

SetIsTimeoutSupported sets IsTimeoutSupported field to given value.

### HasIsTimeoutSupported

`func (o *SFTPConnectionAttributes) HasIsTimeoutSupported() bool`

HasIsTimeoutSupported returns a boolean if a field has been set.

### GetIsTimeoutConfigValidated

`func (o *SFTPConnectionAttributes) GetIsTimeoutConfigValidated() bool`

GetIsTimeoutConfigValidated returns the IsTimeoutConfigValidated field if non-nil, zero value otherwise.

### GetIsTimeoutConfigValidatedOk

`func (o *SFTPConnectionAttributes) GetIsTimeoutConfigValidatedOk() (*bool, bool)`

GetIsTimeoutConfigValidatedOk returns a tuple with the IsTimeoutConfigValidated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsTimeoutConfigValidated

`func (o *SFTPConnectionAttributes) SetIsTimeoutConfigValidated(v bool)`

SetIsTimeoutConfigValidated sets IsTimeoutConfigValidated field to given value.

### HasIsTimeoutConfigValidated

`func (o *SFTPConnectionAttributes) HasIsTimeoutConfigValidated() bool`

HasIsTimeoutConfigValidated returns a boolean if a field has been set.

### GetConnectionTimeoutConfig

`func (o *SFTPConnectionAttributes) GetConnectionTimeoutConfig() ConnectionTimeoutConfig`

GetConnectionTimeoutConfig returns the ConnectionTimeoutConfig field if non-nil, zero value otherwise.

### GetConnectionTimeoutConfigOk

`func (o *SFTPConnectionAttributes) GetConnectionTimeoutConfigOk() (*ConnectionTimeoutConfig, bool)`

GetConnectionTimeoutConfigOk returns a tuple with the ConnectionTimeoutConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionTimeoutConfig

`func (o *SFTPConnectionAttributes) SetConnectionTimeoutConfig(v ConnectionTimeoutConfig)`

SetConnectionTimeoutConfig sets ConnectionTimeoutConfig field to given value.

### HasConnectionTimeoutConfig

`func (o *SFTPConnectionAttributes) HasConnectionTimeoutConfig() bool`

HasConnectionTimeoutConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


