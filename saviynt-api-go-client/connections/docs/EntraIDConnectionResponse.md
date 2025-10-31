# EntraIDConnectionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Msg** | Pointer to **string** | API response message | [optional] 
**Emailtemplate** | Pointer to **string** | Email template for the connection | [optional] 
**Updatedby** | Pointer to **string** | Updator account for the connection | [optional] 
**Connectionname** | Pointer to **string** | Name of the connection | [optional] 
**Connectionkey** | Pointer to **int32** | Connection key | [optional] 
**Description** | Pointer to **string** | Description for the connection | [optional] 
**Connectiontype** | Pointer to **string** | Connection type | [optional] 
**Createdon** | Pointer to **string** | Connection creation time | [optional] 
**Createdby** | Pointer to **string** | Creator account for the connection | [optional] 
**Errorcode** | Pointer to **int32** | Error code | [optional] 
**Status** | Pointer to **int32** |  | [optional] 
**Defaultsavroles** | Pointer to **string** |  | [optional] 
**Connectionattributes** | Pointer to [**EntraIDConnectionAttributes**](EntraIDConnectionAttributes.md) |  | [optional] 

## Methods

### NewEntraIDConnectionResponse

`func NewEntraIDConnectionResponse() *EntraIDConnectionResponse`

NewEntraIDConnectionResponse instantiates a new EntraIDConnectionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEntraIDConnectionResponseWithDefaults

`func NewEntraIDConnectionResponseWithDefaults() *EntraIDConnectionResponse`

NewEntraIDConnectionResponseWithDefaults instantiates a new EntraIDConnectionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMsg

`func (o *EntraIDConnectionResponse) GetMsg() string`

GetMsg returns the Msg field if non-nil, zero value otherwise.

### GetMsgOk

`func (o *EntraIDConnectionResponse) GetMsgOk() (*string, bool)`

GetMsgOk returns a tuple with the Msg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsg

`func (o *EntraIDConnectionResponse) SetMsg(v string)`

SetMsg sets Msg field to given value.

### HasMsg

`func (o *EntraIDConnectionResponse) HasMsg() bool`

HasMsg returns a boolean if a field has been set.

### GetEmailtemplate

`func (o *EntraIDConnectionResponse) GetEmailtemplate() string`

GetEmailtemplate returns the Emailtemplate field if non-nil, zero value otherwise.

### GetEmailtemplateOk

`func (o *EntraIDConnectionResponse) GetEmailtemplateOk() (*string, bool)`

GetEmailtemplateOk returns a tuple with the Emailtemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailtemplate

`func (o *EntraIDConnectionResponse) SetEmailtemplate(v string)`

SetEmailtemplate sets Emailtemplate field to given value.

### HasEmailtemplate

`func (o *EntraIDConnectionResponse) HasEmailtemplate() bool`

HasEmailtemplate returns a boolean if a field has been set.

### GetUpdatedby

`func (o *EntraIDConnectionResponse) GetUpdatedby() string`

GetUpdatedby returns the Updatedby field if non-nil, zero value otherwise.

### GetUpdatedbyOk

`func (o *EntraIDConnectionResponse) GetUpdatedbyOk() (*string, bool)`

GetUpdatedbyOk returns a tuple with the Updatedby field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedby

`func (o *EntraIDConnectionResponse) SetUpdatedby(v string)`

SetUpdatedby sets Updatedby field to given value.

### HasUpdatedby

`func (o *EntraIDConnectionResponse) HasUpdatedby() bool`

HasUpdatedby returns a boolean if a field has been set.

### GetConnectionname

`func (o *EntraIDConnectionResponse) GetConnectionname() string`

GetConnectionname returns the Connectionname field if non-nil, zero value otherwise.

### GetConnectionnameOk

`func (o *EntraIDConnectionResponse) GetConnectionnameOk() (*string, bool)`

GetConnectionnameOk returns a tuple with the Connectionname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionname

`func (o *EntraIDConnectionResponse) SetConnectionname(v string)`

SetConnectionname sets Connectionname field to given value.

### HasConnectionname

`func (o *EntraIDConnectionResponse) HasConnectionname() bool`

HasConnectionname returns a boolean if a field has been set.

### GetConnectionkey

`func (o *EntraIDConnectionResponse) GetConnectionkey() int32`

GetConnectionkey returns the Connectionkey field if non-nil, zero value otherwise.

### GetConnectionkeyOk

`func (o *EntraIDConnectionResponse) GetConnectionkeyOk() (*int32, bool)`

GetConnectionkeyOk returns a tuple with the Connectionkey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionkey

`func (o *EntraIDConnectionResponse) SetConnectionkey(v int32)`

SetConnectionkey sets Connectionkey field to given value.

### HasConnectionkey

`func (o *EntraIDConnectionResponse) HasConnectionkey() bool`

HasConnectionkey returns a boolean if a field has been set.

### GetDescription

`func (o *EntraIDConnectionResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *EntraIDConnectionResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *EntraIDConnectionResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *EntraIDConnectionResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetConnectiontype

`func (o *EntraIDConnectionResponse) GetConnectiontype() string`

GetConnectiontype returns the Connectiontype field if non-nil, zero value otherwise.

### GetConnectiontypeOk

`func (o *EntraIDConnectionResponse) GetConnectiontypeOk() (*string, bool)`

GetConnectiontypeOk returns a tuple with the Connectiontype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectiontype

`func (o *EntraIDConnectionResponse) SetConnectiontype(v string)`

SetConnectiontype sets Connectiontype field to given value.

### HasConnectiontype

`func (o *EntraIDConnectionResponse) HasConnectiontype() bool`

HasConnectiontype returns a boolean if a field has been set.

### GetCreatedon

`func (o *EntraIDConnectionResponse) GetCreatedon() string`

GetCreatedon returns the Createdon field if non-nil, zero value otherwise.

### GetCreatedonOk

`func (o *EntraIDConnectionResponse) GetCreatedonOk() (*string, bool)`

GetCreatedonOk returns a tuple with the Createdon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedon

`func (o *EntraIDConnectionResponse) SetCreatedon(v string)`

SetCreatedon sets Createdon field to given value.

### HasCreatedon

`func (o *EntraIDConnectionResponse) HasCreatedon() bool`

HasCreatedon returns a boolean if a field has been set.

### GetCreatedby

`func (o *EntraIDConnectionResponse) GetCreatedby() string`

GetCreatedby returns the Createdby field if non-nil, zero value otherwise.

### GetCreatedbyOk

`func (o *EntraIDConnectionResponse) GetCreatedbyOk() (*string, bool)`

GetCreatedbyOk returns a tuple with the Createdby field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedby

`func (o *EntraIDConnectionResponse) SetCreatedby(v string)`

SetCreatedby sets Createdby field to given value.

### HasCreatedby

`func (o *EntraIDConnectionResponse) HasCreatedby() bool`

HasCreatedby returns a boolean if a field has been set.

### GetErrorcode

`func (o *EntraIDConnectionResponse) GetErrorcode() int32`

GetErrorcode returns the Errorcode field if non-nil, zero value otherwise.

### GetErrorcodeOk

`func (o *EntraIDConnectionResponse) GetErrorcodeOk() (*int32, bool)`

GetErrorcodeOk returns a tuple with the Errorcode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorcode

`func (o *EntraIDConnectionResponse) SetErrorcode(v int32)`

SetErrorcode sets Errorcode field to given value.

### HasErrorcode

`func (o *EntraIDConnectionResponse) HasErrorcode() bool`

HasErrorcode returns a boolean if a field has been set.

### GetStatus

`func (o *EntraIDConnectionResponse) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *EntraIDConnectionResponse) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *EntraIDConnectionResponse) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *EntraIDConnectionResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetDefaultsavroles

`func (o *EntraIDConnectionResponse) GetDefaultsavroles() string`

GetDefaultsavroles returns the Defaultsavroles field if non-nil, zero value otherwise.

### GetDefaultsavrolesOk

`func (o *EntraIDConnectionResponse) GetDefaultsavrolesOk() (*string, bool)`

GetDefaultsavrolesOk returns a tuple with the Defaultsavroles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultsavroles

`func (o *EntraIDConnectionResponse) SetDefaultsavroles(v string)`

SetDefaultsavroles sets Defaultsavroles field to given value.

### HasDefaultsavroles

`func (o *EntraIDConnectionResponse) HasDefaultsavroles() bool`

HasDefaultsavroles returns a boolean if a field has been set.

### GetConnectionattributes

`func (o *EntraIDConnectionResponse) GetConnectionattributes() EntraIDConnectionAttributes`

GetConnectionattributes returns the Connectionattributes field if non-nil, zero value otherwise.

### GetConnectionattributesOk

`func (o *EntraIDConnectionResponse) GetConnectionattributesOk() (*EntraIDConnectionAttributes, bool)`

GetConnectionattributesOk returns a tuple with the Connectionattributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionattributes

`func (o *EntraIDConnectionResponse) SetConnectionattributes(v EntraIDConnectionAttributes)`

SetConnectionattributes sets Connectionattributes field to given value.

### HasConnectionattributes

`func (o *EntraIDConnectionResponse) HasConnectionattributes() bool`

HasConnectionattributes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


