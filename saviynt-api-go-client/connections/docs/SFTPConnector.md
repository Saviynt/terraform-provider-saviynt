# SFTPConnector

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HOST_NAME** | **string** | SFTP server hostname or IP address | 
**PORT_NUMBER** | **string** | SFTP server port number | 
**USERNAME** | **string** | Username for SFTP authentication | 
**AUTH_CREDENTIAL_TYPE** | **string** | Type of authentication (password, key, etc.) | 
**AUTH_CREDENTIAL_VALUE** | **string** | Authentication credential (password or private key path) | 
**PASSPHRASE** | Pointer to **string** | Passphrase for encrypted private key | [optional] 
**FILES_TO_GET** | Pointer to **string** | Files to download from SFTP server | [optional] 
**FILES_TO_PUT** | Pointer to **string** | Files to upload to SFTP server | [optional] 
**PAM_CONFIG** | Pointer to **string** | PAM configuration for SFTP connection | [optional] 

## Methods

### NewSFTPConnector

`func NewSFTPConnector(hOSTNAME string, pORTNUMBER string, uSERNAME string, aUTHCREDENTIALTYPE string, aUTHCREDENTIALVALUE string, ) *SFTPConnector`

NewSFTPConnector instantiates a new SFTPConnector object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSFTPConnectorWithDefaults

`func NewSFTPConnectorWithDefaults() *SFTPConnector`

NewSFTPConnectorWithDefaults instantiates a new SFTPConnector object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHOST_NAME

`func (o *SFTPConnector) GetHOST_NAME() string`

GetHOST_NAME returns the HOST_NAME field if non-nil, zero value otherwise.

### GetHOST_NAMEOk

`func (o *SFTPConnector) GetHOST_NAMEOk() (*string, bool)`

GetHOST_NAMEOk returns a tuple with the HOST_NAME field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHOST_NAME

`func (o *SFTPConnector) SetHOST_NAME(v string)`

SetHOST_NAME sets HOST_NAME field to given value.


### GetPORT_NUMBER

`func (o *SFTPConnector) GetPORT_NUMBER() string`

GetPORT_NUMBER returns the PORT_NUMBER field if non-nil, zero value otherwise.

### GetPORT_NUMBEROk

`func (o *SFTPConnector) GetPORT_NUMBEROk() (*string, bool)`

GetPORT_NUMBEROk returns a tuple with the PORT_NUMBER field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPORT_NUMBER

`func (o *SFTPConnector) SetPORT_NUMBER(v string)`

SetPORT_NUMBER sets PORT_NUMBER field to given value.


### GetUSERNAME

`func (o *SFTPConnector) GetUSERNAME() string`

GetUSERNAME returns the USERNAME field if non-nil, zero value otherwise.

### GetUSERNAMEOk

`func (o *SFTPConnector) GetUSERNAMEOk() (*string, bool)`

GetUSERNAMEOk returns a tuple with the USERNAME field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUSERNAME

`func (o *SFTPConnector) SetUSERNAME(v string)`

SetUSERNAME sets USERNAME field to given value.


### GetAUTH_CREDENTIAL_TYPE

`func (o *SFTPConnector) GetAUTH_CREDENTIAL_TYPE() string`

GetAUTH_CREDENTIAL_TYPE returns the AUTH_CREDENTIAL_TYPE field if non-nil, zero value otherwise.

### GetAUTH_CREDENTIAL_TYPEOk

`func (o *SFTPConnector) GetAUTH_CREDENTIAL_TYPEOk() (*string, bool)`

GetAUTH_CREDENTIAL_TYPEOk returns a tuple with the AUTH_CREDENTIAL_TYPE field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAUTH_CREDENTIAL_TYPE

`func (o *SFTPConnector) SetAUTH_CREDENTIAL_TYPE(v string)`

SetAUTH_CREDENTIAL_TYPE sets AUTH_CREDENTIAL_TYPE field to given value.


### GetAUTH_CREDENTIAL_VALUE

`func (o *SFTPConnector) GetAUTH_CREDENTIAL_VALUE() string`

GetAUTH_CREDENTIAL_VALUE returns the AUTH_CREDENTIAL_VALUE field if non-nil, zero value otherwise.

### GetAUTH_CREDENTIAL_VALUEOk

`func (o *SFTPConnector) GetAUTH_CREDENTIAL_VALUEOk() (*string, bool)`

GetAUTH_CREDENTIAL_VALUEOk returns a tuple with the AUTH_CREDENTIAL_VALUE field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAUTH_CREDENTIAL_VALUE

`func (o *SFTPConnector) SetAUTH_CREDENTIAL_VALUE(v string)`

SetAUTH_CREDENTIAL_VALUE sets AUTH_CREDENTIAL_VALUE field to given value.


### GetPASSPHRASE

`func (o *SFTPConnector) GetPASSPHRASE() string`

GetPASSPHRASE returns the PASSPHRASE field if non-nil, zero value otherwise.

### GetPASSPHRASEOk

`func (o *SFTPConnector) GetPASSPHRASEOk() (*string, bool)`

GetPASSPHRASEOk returns a tuple with the PASSPHRASE field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPASSPHRASE

`func (o *SFTPConnector) SetPASSPHRASE(v string)`

SetPASSPHRASE sets PASSPHRASE field to given value.

### HasPASSPHRASE

`func (o *SFTPConnector) HasPASSPHRASE() bool`

HasPASSPHRASE returns a boolean if a field has been set.

### GetFILES_TO_GET

`func (o *SFTPConnector) GetFILES_TO_GET() string`

GetFILES_TO_GET returns the FILES_TO_GET field if non-nil, zero value otherwise.

### GetFILES_TO_GETOk

`func (o *SFTPConnector) GetFILES_TO_GETOk() (*string, bool)`

GetFILES_TO_GETOk returns a tuple with the FILES_TO_GET field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFILES_TO_GET

`func (o *SFTPConnector) SetFILES_TO_GET(v string)`

SetFILES_TO_GET sets FILES_TO_GET field to given value.

### HasFILES_TO_GET

`func (o *SFTPConnector) HasFILES_TO_GET() bool`

HasFILES_TO_GET returns a boolean if a field has been set.

### GetFILES_TO_PUT

`func (o *SFTPConnector) GetFILES_TO_PUT() string`

GetFILES_TO_PUT returns the FILES_TO_PUT field if non-nil, zero value otherwise.

### GetFILES_TO_PUTOk

`func (o *SFTPConnector) GetFILES_TO_PUTOk() (*string, bool)`

GetFILES_TO_PUTOk returns a tuple with the FILES_TO_PUT field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFILES_TO_PUT

`func (o *SFTPConnector) SetFILES_TO_PUT(v string)`

SetFILES_TO_PUT sets FILES_TO_PUT field to given value.

### HasFILES_TO_PUT

`func (o *SFTPConnector) HasFILES_TO_PUT() bool`

HasFILES_TO_PUT returns a boolean if a field has been set.

### GetPAM_CONFIG

`func (o *SFTPConnector) GetPAM_CONFIG() string`

GetPAM_CONFIG returns the PAM_CONFIG field if non-nil, zero value otherwise.

### GetPAM_CONFIGOk

`func (o *SFTPConnector) GetPAM_CONFIGOk() (*string, bool)`

GetPAM_CONFIGOk returns a tuple with the PAM_CONFIG field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPAM_CONFIG

`func (o *SFTPConnector) SetPAM_CONFIG(v string)`

SetPAM_CONFIG sets PAM_CONFIG field to given value.

### HasPAM_CONFIG

`func (o *SFTPConnector) HasPAM_CONFIG() bool`

HasPAM_CONFIG returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


