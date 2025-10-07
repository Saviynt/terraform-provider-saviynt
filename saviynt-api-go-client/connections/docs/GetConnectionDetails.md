# GetConnectionDetails

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

## Methods

### NewGetConnectionDetails

`func NewGetConnectionDetails() *GetConnectionDetails`

NewGetConnectionDetails instantiates a new GetConnectionDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetConnectionDetailsWithDefaults

`func NewGetConnectionDetailsWithDefaults() *GetConnectionDetails`

NewGetConnectionDetailsWithDefaults instantiates a new GetConnectionDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMsg

`func (o *GetConnectionDetails) GetMsg() string`

GetMsg returns the Msg field if non-nil, zero value otherwise.

### GetMsgOk

`func (o *GetConnectionDetails) GetMsgOk() (*string, bool)`

GetMsgOk returns a tuple with the Msg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsg

`func (o *GetConnectionDetails) SetMsg(v string)`

SetMsg sets Msg field to given value.

### HasMsg

`func (o *GetConnectionDetails) HasMsg() bool`

HasMsg returns a boolean if a field has been set.

### GetEmailtemplate

`func (o *GetConnectionDetails) GetEmailtemplate() string`

GetEmailtemplate returns the Emailtemplate field if non-nil, zero value otherwise.

### GetEmailtemplateOk

`func (o *GetConnectionDetails) GetEmailtemplateOk() (*string, bool)`

GetEmailtemplateOk returns a tuple with the Emailtemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailtemplate

`func (o *GetConnectionDetails) SetEmailtemplate(v string)`

SetEmailtemplate sets Emailtemplate field to given value.

### HasEmailtemplate

`func (o *GetConnectionDetails) HasEmailtemplate() bool`

HasEmailtemplate returns a boolean if a field has been set.

### GetUpdatedby

`func (o *GetConnectionDetails) GetUpdatedby() string`

GetUpdatedby returns the Updatedby field if non-nil, zero value otherwise.

### GetUpdatedbyOk

`func (o *GetConnectionDetails) GetUpdatedbyOk() (*string, bool)`

GetUpdatedbyOk returns a tuple with the Updatedby field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedby

`func (o *GetConnectionDetails) SetUpdatedby(v string)`

SetUpdatedby sets Updatedby field to given value.

### HasUpdatedby

`func (o *GetConnectionDetails) HasUpdatedby() bool`

HasUpdatedby returns a boolean if a field has been set.

### GetConnectionname

`func (o *GetConnectionDetails) GetConnectionname() string`

GetConnectionname returns the Connectionname field if non-nil, zero value otherwise.

### GetConnectionnameOk

`func (o *GetConnectionDetails) GetConnectionnameOk() (*string, bool)`

GetConnectionnameOk returns a tuple with the Connectionname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionname

`func (o *GetConnectionDetails) SetConnectionname(v string)`

SetConnectionname sets Connectionname field to given value.

### HasConnectionname

`func (o *GetConnectionDetails) HasConnectionname() bool`

HasConnectionname returns a boolean if a field has been set.

### GetConnectionkey

`func (o *GetConnectionDetails) GetConnectionkey() int32`

GetConnectionkey returns the Connectionkey field if non-nil, zero value otherwise.

### GetConnectionkeyOk

`func (o *GetConnectionDetails) GetConnectionkeyOk() (*int32, bool)`

GetConnectionkeyOk returns a tuple with the Connectionkey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionkey

`func (o *GetConnectionDetails) SetConnectionkey(v int32)`

SetConnectionkey sets Connectionkey field to given value.

### HasConnectionkey

`func (o *GetConnectionDetails) HasConnectionkey() bool`

HasConnectionkey returns a boolean if a field has been set.

### GetDescription

`func (o *GetConnectionDetails) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GetConnectionDetails) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GetConnectionDetails) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GetConnectionDetails) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetConnectiontype

`func (o *GetConnectionDetails) GetConnectiontype() string`

GetConnectiontype returns the Connectiontype field if non-nil, zero value otherwise.

### GetConnectiontypeOk

`func (o *GetConnectionDetails) GetConnectiontypeOk() (*string, bool)`

GetConnectiontypeOk returns a tuple with the Connectiontype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectiontype

`func (o *GetConnectionDetails) SetConnectiontype(v string)`

SetConnectiontype sets Connectiontype field to given value.

### HasConnectiontype

`func (o *GetConnectionDetails) HasConnectiontype() bool`

HasConnectiontype returns a boolean if a field has been set.

### GetCreatedon

`func (o *GetConnectionDetails) GetCreatedon() string`

GetCreatedon returns the Createdon field if non-nil, zero value otherwise.

### GetCreatedonOk

`func (o *GetConnectionDetails) GetCreatedonOk() (*string, bool)`

GetCreatedonOk returns a tuple with the Createdon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedon

`func (o *GetConnectionDetails) SetCreatedon(v string)`

SetCreatedon sets Createdon field to given value.

### HasCreatedon

`func (o *GetConnectionDetails) HasCreatedon() bool`

HasCreatedon returns a boolean if a field has been set.

### GetCreatedby

`func (o *GetConnectionDetails) GetCreatedby() string`

GetCreatedby returns the Createdby field if non-nil, zero value otherwise.

### GetCreatedbyOk

`func (o *GetConnectionDetails) GetCreatedbyOk() (*string, bool)`

GetCreatedbyOk returns a tuple with the Createdby field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedby

`func (o *GetConnectionDetails) SetCreatedby(v string)`

SetCreatedby sets Createdby field to given value.

### HasCreatedby

`func (o *GetConnectionDetails) HasCreatedby() bool`

HasCreatedby returns a boolean if a field has been set.

### GetErrorcode

`func (o *GetConnectionDetails) GetErrorcode() int32`

GetErrorcode returns the Errorcode field if non-nil, zero value otherwise.

### GetErrorcodeOk

`func (o *GetConnectionDetails) GetErrorcodeOk() (*int32, bool)`

GetErrorcodeOk returns a tuple with the Errorcode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorcode

`func (o *GetConnectionDetails) SetErrorcode(v int32)`

SetErrorcode sets Errorcode field to given value.

### HasErrorcode

`func (o *GetConnectionDetails) HasErrorcode() bool`

HasErrorcode returns a boolean if a field has been set.

### GetStatus

`func (o *GetConnectionDetails) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetConnectionDetails) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetConnectionDetails) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetConnectionDetails) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetDefaultsavroles

`func (o *GetConnectionDetails) GetDefaultsavroles() string`

GetDefaultsavroles returns the Defaultsavroles field if non-nil, zero value otherwise.

### GetDefaultsavrolesOk

`func (o *GetConnectionDetails) GetDefaultsavrolesOk() (*string, bool)`

GetDefaultsavrolesOk returns a tuple with the Defaultsavroles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultsavroles

`func (o *GetConnectionDetails) SetDefaultsavroles(v string)`

SetDefaultsavroles sets Defaultsavroles field to given value.

### HasDefaultsavroles

`func (o *GetConnectionDetails) HasDefaultsavroles() bool`

HasDefaultsavroles returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


