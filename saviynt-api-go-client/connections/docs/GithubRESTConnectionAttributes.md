# GithubRESTConnectionAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsTimeoutSupported** | Pointer to **bool** |  | [optional] 
**ConnectionJSON** | Pointer to **string** |  | [optional] 
**PAM_CONFIG** | Pointer to **string** |  | [optional] 
**ORGANIZATION_LIST** | Pointer to **string** |  | [optional] 
**ImportAccountEntJSON** | Pointer to **string** |  | [optional] 
**STATUS_THRESHOLD_CONFIG** | Pointer to **string** |  | [optional] 
**ACCESS_TOKENS** | Pointer to **string** |  | [optional] 
**ConnectionTimeoutConfig** | Pointer to [**ConnectionTimeoutConfig**](ConnectionTimeoutConfig.md) |  | [optional] 
**ConnectionType** | Pointer to **string** |  | [optional] 
**IsTimeoutConfigValidated** | Pointer to **bool** |  | [optional] 

## Methods

### NewGithubRESTConnectionAttributes

`func NewGithubRESTConnectionAttributes() *GithubRESTConnectionAttributes`

NewGithubRESTConnectionAttributes instantiates a new GithubRESTConnectionAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGithubRESTConnectionAttributesWithDefaults

`func NewGithubRESTConnectionAttributesWithDefaults() *GithubRESTConnectionAttributes`

NewGithubRESTConnectionAttributesWithDefaults instantiates a new GithubRESTConnectionAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsTimeoutSupported

`func (o *GithubRESTConnectionAttributes) GetIsTimeoutSupported() bool`

GetIsTimeoutSupported returns the IsTimeoutSupported field if non-nil, zero value otherwise.

### GetIsTimeoutSupportedOk

`func (o *GithubRESTConnectionAttributes) GetIsTimeoutSupportedOk() (*bool, bool)`

GetIsTimeoutSupportedOk returns a tuple with the IsTimeoutSupported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsTimeoutSupported

`func (o *GithubRESTConnectionAttributes) SetIsTimeoutSupported(v bool)`

SetIsTimeoutSupported sets IsTimeoutSupported field to given value.

### HasIsTimeoutSupported

`func (o *GithubRESTConnectionAttributes) HasIsTimeoutSupported() bool`

HasIsTimeoutSupported returns a boolean if a field has been set.

### GetConnectionJSON

`func (o *GithubRESTConnectionAttributes) GetConnectionJSON() string`

GetConnectionJSON returns the ConnectionJSON field if non-nil, zero value otherwise.

### GetConnectionJSONOk

`func (o *GithubRESTConnectionAttributes) GetConnectionJSONOk() (*string, bool)`

GetConnectionJSONOk returns a tuple with the ConnectionJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionJSON

`func (o *GithubRESTConnectionAttributes) SetConnectionJSON(v string)`

SetConnectionJSON sets ConnectionJSON field to given value.

### HasConnectionJSON

`func (o *GithubRESTConnectionAttributes) HasConnectionJSON() bool`

HasConnectionJSON returns a boolean if a field has been set.

### GetPAM_CONFIG

`func (o *GithubRESTConnectionAttributes) GetPAM_CONFIG() string`

GetPAM_CONFIG returns the PAM_CONFIG field if non-nil, zero value otherwise.

### GetPAM_CONFIGOk

`func (o *GithubRESTConnectionAttributes) GetPAM_CONFIGOk() (*string, bool)`

GetPAM_CONFIGOk returns a tuple with the PAM_CONFIG field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPAM_CONFIG

`func (o *GithubRESTConnectionAttributes) SetPAM_CONFIG(v string)`

SetPAM_CONFIG sets PAM_CONFIG field to given value.

### HasPAM_CONFIG

`func (o *GithubRESTConnectionAttributes) HasPAM_CONFIG() bool`

HasPAM_CONFIG returns a boolean if a field has been set.

### GetORGANIZATION_LIST

`func (o *GithubRESTConnectionAttributes) GetORGANIZATION_LIST() string`

GetORGANIZATION_LIST returns the ORGANIZATION_LIST field if non-nil, zero value otherwise.

### GetORGANIZATION_LISTOk

`func (o *GithubRESTConnectionAttributes) GetORGANIZATION_LISTOk() (*string, bool)`

GetORGANIZATION_LISTOk returns a tuple with the ORGANIZATION_LIST field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetORGANIZATION_LIST

`func (o *GithubRESTConnectionAttributes) SetORGANIZATION_LIST(v string)`

SetORGANIZATION_LIST sets ORGANIZATION_LIST field to given value.

### HasORGANIZATION_LIST

`func (o *GithubRESTConnectionAttributes) HasORGANIZATION_LIST() bool`

HasORGANIZATION_LIST returns a boolean if a field has been set.

### GetImportAccountEntJSON

`func (o *GithubRESTConnectionAttributes) GetImportAccountEntJSON() string`

GetImportAccountEntJSON returns the ImportAccountEntJSON field if non-nil, zero value otherwise.

### GetImportAccountEntJSONOk

`func (o *GithubRESTConnectionAttributes) GetImportAccountEntJSONOk() (*string, bool)`

GetImportAccountEntJSONOk returns a tuple with the ImportAccountEntJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportAccountEntJSON

`func (o *GithubRESTConnectionAttributes) SetImportAccountEntJSON(v string)`

SetImportAccountEntJSON sets ImportAccountEntJSON field to given value.

### HasImportAccountEntJSON

`func (o *GithubRESTConnectionAttributes) HasImportAccountEntJSON() bool`

HasImportAccountEntJSON returns a boolean if a field has been set.

### GetSTATUS_THRESHOLD_CONFIG

`func (o *GithubRESTConnectionAttributes) GetSTATUS_THRESHOLD_CONFIG() string`

GetSTATUS_THRESHOLD_CONFIG returns the STATUS_THRESHOLD_CONFIG field if non-nil, zero value otherwise.

### GetSTATUS_THRESHOLD_CONFIGOk

`func (o *GithubRESTConnectionAttributes) GetSTATUS_THRESHOLD_CONFIGOk() (*string, bool)`

GetSTATUS_THRESHOLD_CONFIGOk returns a tuple with the STATUS_THRESHOLD_CONFIG field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSTATUS_THRESHOLD_CONFIG

`func (o *GithubRESTConnectionAttributes) SetSTATUS_THRESHOLD_CONFIG(v string)`

SetSTATUS_THRESHOLD_CONFIG sets STATUS_THRESHOLD_CONFIG field to given value.

### HasSTATUS_THRESHOLD_CONFIG

`func (o *GithubRESTConnectionAttributes) HasSTATUS_THRESHOLD_CONFIG() bool`

HasSTATUS_THRESHOLD_CONFIG returns a boolean if a field has been set.

### GetACCESS_TOKENS

`func (o *GithubRESTConnectionAttributes) GetACCESS_TOKENS() string`

GetACCESS_TOKENS returns the ACCESS_TOKENS field if non-nil, zero value otherwise.

### GetACCESS_TOKENSOk

`func (o *GithubRESTConnectionAttributes) GetACCESS_TOKENSOk() (*string, bool)`

GetACCESS_TOKENSOk returns a tuple with the ACCESS_TOKENS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetACCESS_TOKENS

`func (o *GithubRESTConnectionAttributes) SetACCESS_TOKENS(v string)`

SetACCESS_TOKENS sets ACCESS_TOKENS field to given value.

### HasACCESS_TOKENS

`func (o *GithubRESTConnectionAttributes) HasACCESS_TOKENS() bool`

HasACCESS_TOKENS returns a boolean if a field has been set.

### GetConnectionTimeoutConfig

`func (o *GithubRESTConnectionAttributes) GetConnectionTimeoutConfig() ConnectionTimeoutConfig`

GetConnectionTimeoutConfig returns the ConnectionTimeoutConfig field if non-nil, zero value otherwise.

### GetConnectionTimeoutConfigOk

`func (o *GithubRESTConnectionAttributes) GetConnectionTimeoutConfigOk() (*ConnectionTimeoutConfig, bool)`

GetConnectionTimeoutConfigOk returns a tuple with the ConnectionTimeoutConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionTimeoutConfig

`func (o *GithubRESTConnectionAttributes) SetConnectionTimeoutConfig(v ConnectionTimeoutConfig)`

SetConnectionTimeoutConfig sets ConnectionTimeoutConfig field to given value.

### HasConnectionTimeoutConfig

`func (o *GithubRESTConnectionAttributes) HasConnectionTimeoutConfig() bool`

HasConnectionTimeoutConfig returns a boolean if a field has been set.

### GetConnectionType

`func (o *GithubRESTConnectionAttributes) GetConnectionType() string`

GetConnectionType returns the ConnectionType field if non-nil, zero value otherwise.

### GetConnectionTypeOk

`func (o *GithubRESTConnectionAttributes) GetConnectionTypeOk() (*string, bool)`

GetConnectionTypeOk returns a tuple with the ConnectionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionType

`func (o *GithubRESTConnectionAttributes) SetConnectionType(v string)`

SetConnectionType sets ConnectionType field to given value.

### HasConnectionType

`func (o *GithubRESTConnectionAttributes) HasConnectionType() bool`

HasConnectionType returns a boolean if a field has been set.

### GetIsTimeoutConfigValidated

`func (o *GithubRESTConnectionAttributes) GetIsTimeoutConfigValidated() bool`

GetIsTimeoutConfigValidated returns the IsTimeoutConfigValidated field if non-nil, zero value otherwise.

### GetIsTimeoutConfigValidatedOk

`func (o *GithubRESTConnectionAttributes) GetIsTimeoutConfigValidatedOk() (*bool, bool)`

GetIsTimeoutConfigValidatedOk returns a tuple with the IsTimeoutConfigValidated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsTimeoutConfigValidated

`func (o *GithubRESTConnectionAttributes) SetIsTimeoutConfigValidated(v bool)`

SetIsTimeoutConfigValidated sets IsTimeoutConfigValidated field to given value.

### HasIsTimeoutConfigValidated

`func (o *GithubRESTConnectionAttributes) HasIsTimeoutConfigValidated() bool`

HasIsTimeoutConfigValidated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


