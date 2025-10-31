# WorkdaySOAPConnectionResponse

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
**Connectionattributes** | Pointer to [**WorkdaySOAPConnectionAttributes**](WorkdaySOAPConnectionAttributes.md) |  | [optional] 

## Methods

### NewWorkdaySOAPConnectionResponse

`func NewWorkdaySOAPConnectionResponse() *WorkdaySOAPConnectionResponse`

NewWorkdaySOAPConnectionResponse instantiates a new WorkdaySOAPConnectionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkdaySOAPConnectionResponseWithDefaults

`func NewWorkdaySOAPConnectionResponseWithDefaults() *WorkdaySOAPConnectionResponse`

NewWorkdaySOAPConnectionResponseWithDefaults instantiates a new WorkdaySOAPConnectionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMsg

`func (o *WorkdaySOAPConnectionResponse) GetMsg() string`

GetMsg returns the Msg field if non-nil, zero value otherwise.

### GetMsgOk

`func (o *WorkdaySOAPConnectionResponse) GetMsgOk() (*string, bool)`

GetMsgOk returns a tuple with the Msg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsg

`func (o *WorkdaySOAPConnectionResponse) SetMsg(v string)`

SetMsg sets Msg field to given value.

### HasMsg

`func (o *WorkdaySOAPConnectionResponse) HasMsg() bool`

HasMsg returns a boolean if a field has been set.

### GetEmailtemplate

`func (o *WorkdaySOAPConnectionResponse) GetEmailtemplate() string`

GetEmailtemplate returns the Emailtemplate field if non-nil, zero value otherwise.

### GetEmailtemplateOk

`func (o *WorkdaySOAPConnectionResponse) GetEmailtemplateOk() (*string, bool)`

GetEmailtemplateOk returns a tuple with the Emailtemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailtemplate

`func (o *WorkdaySOAPConnectionResponse) SetEmailtemplate(v string)`

SetEmailtemplate sets Emailtemplate field to given value.

### HasEmailtemplate

`func (o *WorkdaySOAPConnectionResponse) HasEmailtemplate() bool`

HasEmailtemplate returns a boolean if a field has been set.

### GetUpdatedby

`func (o *WorkdaySOAPConnectionResponse) GetUpdatedby() string`

GetUpdatedby returns the Updatedby field if non-nil, zero value otherwise.

### GetUpdatedbyOk

`func (o *WorkdaySOAPConnectionResponse) GetUpdatedbyOk() (*string, bool)`

GetUpdatedbyOk returns a tuple with the Updatedby field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedby

`func (o *WorkdaySOAPConnectionResponse) SetUpdatedby(v string)`

SetUpdatedby sets Updatedby field to given value.

### HasUpdatedby

`func (o *WorkdaySOAPConnectionResponse) HasUpdatedby() bool`

HasUpdatedby returns a boolean if a field has been set.

### GetConnectionname

`func (o *WorkdaySOAPConnectionResponse) GetConnectionname() string`

GetConnectionname returns the Connectionname field if non-nil, zero value otherwise.

### GetConnectionnameOk

`func (o *WorkdaySOAPConnectionResponse) GetConnectionnameOk() (*string, bool)`

GetConnectionnameOk returns a tuple with the Connectionname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionname

`func (o *WorkdaySOAPConnectionResponse) SetConnectionname(v string)`

SetConnectionname sets Connectionname field to given value.

### HasConnectionname

`func (o *WorkdaySOAPConnectionResponse) HasConnectionname() bool`

HasConnectionname returns a boolean if a field has been set.

### GetConnectionkey

`func (o *WorkdaySOAPConnectionResponse) GetConnectionkey() int32`

GetConnectionkey returns the Connectionkey field if non-nil, zero value otherwise.

### GetConnectionkeyOk

`func (o *WorkdaySOAPConnectionResponse) GetConnectionkeyOk() (*int32, bool)`

GetConnectionkeyOk returns a tuple with the Connectionkey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionkey

`func (o *WorkdaySOAPConnectionResponse) SetConnectionkey(v int32)`

SetConnectionkey sets Connectionkey field to given value.

### HasConnectionkey

`func (o *WorkdaySOAPConnectionResponse) HasConnectionkey() bool`

HasConnectionkey returns a boolean if a field has been set.

### GetDescription

`func (o *WorkdaySOAPConnectionResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *WorkdaySOAPConnectionResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *WorkdaySOAPConnectionResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *WorkdaySOAPConnectionResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetConnectiontype

`func (o *WorkdaySOAPConnectionResponse) GetConnectiontype() string`

GetConnectiontype returns the Connectiontype field if non-nil, zero value otherwise.

### GetConnectiontypeOk

`func (o *WorkdaySOAPConnectionResponse) GetConnectiontypeOk() (*string, bool)`

GetConnectiontypeOk returns a tuple with the Connectiontype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectiontype

`func (o *WorkdaySOAPConnectionResponse) SetConnectiontype(v string)`

SetConnectiontype sets Connectiontype field to given value.

### HasConnectiontype

`func (o *WorkdaySOAPConnectionResponse) HasConnectiontype() bool`

HasConnectiontype returns a boolean if a field has been set.

### GetCreatedon

`func (o *WorkdaySOAPConnectionResponse) GetCreatedon() string`

GetCreatedon returns the Createdon field if non-nil, zero value otherwise.

### GetCreatedonOk

`func (o *WorkdaySOAPConnectionResponse) GetCreatedonOk() (*string, bool)`

GetCreatedonOk returns a tuple with the Createdon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedon

`func (o *WorkdaySOAPConnectionResponse) SetCreatedon(v string)`

SetCreatedon sets Createdon field to given value.

### HasCreatedon

`func (o *WorkdaySOAPConnectionResponse) HasCreatedon() bool`

HasCreatedon returns a boolean if a field has been set.

### GetCreatedby

`func (o *WorkdaySOAPConnectionResponse) GetCreatedby() string`

GetCreatedby returns the Createdby field if non-nil, zero value otherwise.

### GetCreatedbyOk

`func (o *WorkdaySOAPConnectionResponse) GetCreatedbyOk() (*string, bool)`

GetCreatedbyOk returns a tuple with the Createdby field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedby

`func (o *WorkdaySOAPConnectionResponse) SetCreatedby(v string)`

SetCreatedby sets Createdby field to given value.

### HasCreatedby

`func (o *WorkdaySOAPConnectionResponse) HasCreatedby() bool`

HasCreatedby returns a boolean if a field has been set.

### GetErrorcode

`func (o *WorkdaySOAPConnectionResponse) GetErrorcode() int32`

GetErrorcode returns the Errorcode field if non-nil, zero value otherwise.

### GetErrorcodeOk

`func (o *WorkdaySOAPConnectionResponse) GetErrorcodeOk() (*int32, bool)`

GetErrorcodeOk returns a tuple with the Errorcode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorcode

`func (o *WorkdaySOAPConnectionResponse) SetErrorcode(v int32)`

SetErrorcode sets Errorcode field to given value.

### HasErrorcode

`func (o *WorkdaySOAPConnectionResponse) HasErrorcode() bool`

HasErrorcode returns a boolean if a field has been set.

### GetStatus

`func (o *WorkdaySOAPConnectionResponse) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *WorkdaySOAPConnectionResponse) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *WorkdaySOAPConnectionResponse) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *WorkdaySOAPConnectionResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetDefaultsavroles

`func (o *WorkdaySOAPConnectionResponse) GetDefaultsavroles() string`

GetDefaultsavroles returns the Defaultsavroles field if non-nil, zero value otherwise.

### GetDefaultsavrolesOk

`func (o *WorkdaySOAPConnectionResponse) GetDefaultsavrolesOk() (*string, bool)`

GetDefaultsavrolesOk returns a tuple with the Defaultsavroles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultsavroles

`func (o *WorkdaySOAPConnectionResponse) SetDefaultsavroles(v string)`

SetDefaultsavroles sets Defaultsavroles field to given value.

### HasDefaultsavroles

`func (o *WorkdaySOAPConnectionResponse) HasDefaultsavroles() bool`

HasDefaultsavroles returns a boolean if a field has been set.

### GetConnectionattributes

`func (o *WorkdaySOAPConnectionResponse) GetConnectionattributes() WorkdaySOAPConnectionAttributes`

GetConnectionattributes returns the Connectionattributes field if non-nil, zero value otherwise.

### GetConnectionattributesOk

`func (o *WorkdaySOAPConnectionResponse) GetConnectionattributesOk() (*WorkdaySOAPConnectionAttributes, bool)`

GetConnectionattributesOk returns a tuple with the Connectionattributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionattributes

`func (o *WorkdaySOAPConnectionResponse) SetConnectionattributes(v WorkdaySOAPConnectionAttributes)`

SetConnectionattributes sets Connectionattributes field to given value.

### HasConnectionattributes

`func (o *WorkdaySOAPConnectionResponse) HasConnectionattributes() bool`

HasConnectionattributes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


