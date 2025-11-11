# SalesforceConnectionResponse

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
**Connectionattributes** | Pointer to [**SalesforceConnectionAttributes**](SalesforceConnectionAttributes.md) |  | [optional] 

## Methods

### NewSalesforceConnectionResponse

`func NewSalesforceConnectionResponse() *SalesforceConnectionResponse`

NewSalesforceConnectionResponse instantiates a new SalesforceConnectionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSalesforceConnectionResponseWithDefaults

`func NewSalesforceConnectionResponseWithDefaults() *SalesforceConnectionResponse`

NewSalesforceConnectionResponseWithDefaults instantiates a new SalesforceConnectionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMsg

`func (o *SalesforceConnectionResponse) GetMsg() string`

GetMsg returns the Msg field if non-nil, zero value otherwise.

### GetMsgOk

`func (o *SalesforceConnectionResponse) GetMsgOk() (*string, bool)`

GetMsgOk returns a tuple with the Msg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsg

`func (o *SalesforceConnectionResponse) SetMsg(v string)`

SetMsg sets Msg field to given value.

### HasMsg

`func (o *SalesforceConnectionResponse) HasMsg() bool`

HasMsg returns a boolean if a field has been set.

### GetEmailtemplate

`func (o *SalesforceConnectionResponse) GetEmailtemplate() string`

GetEmailtemplate returns the Emailtemplate field if non-nil, zero value otherwise.

### GetEmailtemplateOk

`func (o *SalesforceConnectionResponse) GetEmailtemplateOk() (*string, bool)`

GetEmailtemplateOk returns a tuple with the Emailtemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailtemplate

`func (o *SalesforceConnectionResponse) SetEmailtemplate(v string)`

SetEmailtemplate sets Emailtemplate field to given value.

### HasEmailtemplate

`func (o *SalesforceConnectionResponse) HasEmailtemplate() bool`

HasEmailtemplate returns a boolean if a field has been set.

### GetUpdatedby

`func (o *SalesforceConnectionResponse) GetUpdatedby() string`

GetUpdatedby returns the Updatedby field if non-nil, zero value otherwise.

### GetUpdatedbyOk

`func (o *SalesforceConnectionResponse) GetUpdatedbyOk() (*string, bool)`

GetUpdatedbyOk returns a tuple with the Updatedby field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedby

`func (o *SalesforceConnectionResponse) SetUpdatedby(v string)`

SetUpdatedby sets Updatedby field to given value.

### HasUpdatedby

`func (o *SalesforceConnectionResponse) HasUpdatedby() bool`

HasUpdatedby returns a boolean if a field has been set.

### GetConnectionname

`func (o *SalesforceConnectionResponse) GetConnectionname() string`

GetConnectionname returns the Connectionname field if non-nil, zero value otherwise.

### GetConnectionnameOk

`func (o *SalesforceConnectionResponse) GetConnectionnameOk() (*string, bool)`

GetConnectionnameOk returns a tuple with the Connectionname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionname

`func (o *SalesforceConnectionResponse) SetConnectionname(v string)`

SetConnectionname sets Connectionname field to given value.

### HasConnectionname

`func (o *SalesforceConnectionResponse) HasConnectionname() bool`

HasConnectionname returns a boolean if a field has been set.

### GetConnectionkey

`func (o *SalesforceConnectionResponse) GetConnectionkey() int32`

GetConnectionkey returns the Connectionkey field if non-nil, zero value otherwise.

### GetConnectionkeyOk

`func (o *SalesforceConnectionResponse) GetConnectionkeyOk() (*int32, bool)`

GetConnectionkeyOk returns a tuple with the Connectionkey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionkey

`func (o *SalesforceConnectionResponse) SetConnectionkey(v int32)`

SetConnectionkey sets Connectionkey field to given value.

### HasConnectionkey

`func (o *SalesforceConnectionResponse) HasConnectionkey() bool`

HasConnectionkey returns a boolean if a field has been set.

### GetDescription

`func (o *SalesforceConnectionResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SalesforceConnectionResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SalesforceConnectionResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SalesforceConnectionResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetConnectiontype

`func (o *SalesforceConnectionResponse) GetConnectiontype() string`

GetConnectiontype returns the Connectiontype field if non-nil, zero value otherwise.

### GetConnectiontypeOk

`func (o *SalesforceConnectionResponse) GetConnectiontypeOk() (*string, bool)`

GetConnectiontypeOk returns a tuple with the Connectiontype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectiontype

`func (o *SalesforceConnectionResponse) SetConnectiontype(v string)`

SetConnectiontype sets Connectiontype field to given value.

### HasConnectiontype

`func (o *SalesforceConnectionResponse) HasConnectiontype() bool`

HasConnectiontype returns a boolean if a field has been set.

### GetCreatedon

`func (o *SalesforceConnectionResponse) GetCreatedon() string`

GetCreatedon returns the Createdon field if non-nil, zero value otherwise.

### GetCreatedonOk

`func (o *SalesforceConnectionResponse) GetCreatedonOk() (*string, bool)`

GetCreatedonOk returns a tuple with the Createdon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedon

`func (o *SalesforceConnectionResponse) SetCreatedon(v string)`

SetCreatedon sets Createdon field to given value.

### HasCreatedon

`func (o *SalesforceConnectionResponse) HasCreatedon() bool`

HasCreatedon returns a boolean if a field has been set.

### GetCreatedby

`func (o *SalesforceConnectionResponse) GetCreatedby() string`

GetCreatedby returns the Createdby field if non-nil, zero value otherwise.

### GetCreatedbyOk

`func (o *SalesforceConnectionResponse) GetCreatedbyOk() (*string, bool)`

GetCreatedbyOk returns a tuple with the Createdby field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedby

`func (o *SalesforceConnectionResponse) SetCreatedby(v string)`

SetCreatedby sets Createdby field to given value.

### HasCreatedby

`func (o *SalesforceConnectionResponse) HasCreatedby() bool`

HasCreatedby returns a boolean if a field has been set.

### GetErrorcode

`func (o *SalesforceConnectionResponse) GetErrorcode() int32`

GetErrorcode returns the Errorcode field if non-nil, zero value otherwise.

### GetErrorcodeOk

`func (o *SalesforceConnectionResponse) GetErrorcodeOk() (*int32, bool)`

GetErrorcodeOk returns a tuple with the Errorcode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorcode

`func (o *SalesforceConnectionResponse) SetErrorcode(v int32)`

SetErrorcode sets Errorcode field to given value.

### HasErrorcode

`func (o *SalesforceConnectionResponse) HasErrorcode() bool`

HasErrorcode returns a boolean if a field has been set.

### GetStatus

`func (o *SalesforceConnectionResponse) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SalesforceConnectionResponse) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SalesforceConnectionResponse) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SalesforceConnectionResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetDefaultsavroles

`func (o *SalesforceConnectionResponse) GetDefaultsavroles() string`

GetDefaultsavroles returns the Defaultsavroles field if non-nil, zero value otherwise.

### GetDefaultsavrolesOk

`func (o *SalesforceConnectionResponse) GetDefaultsavrolesOk() (*string, bool)`

GetDefaultsavrolesOk returns a tuple with the Defaultsavroles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultsavroles

`func (o *SalesforceConnectionResponse) SetDefaultsavroles(v string)`

SetDefaultsavroles sets Defaultsavroles field to given value.

### HasDefaultsavroles

`func (o *SalesforceConnectionResponse) HasDefaultsavroles() bool`

HasDefaultsavroles returns a boolean if a field has been set.

### GetConnectionattributes

`func (o *SalesforceConnectionResponse) GetConnectionattributes() SalesforceConnectionAttributes`

GetConnectionattributes returns the Connectionattributes field if non-nil, zero value otherwise.

### GetConnectionattributesOk

`func (o *SalesforceConnectionResponse) GetConnectionattributesOk() (*SalesforceConnectionAttributes, bool)`

GetConnectionattributesOk returns a tuple with the Connectionattributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionattributes

`func (o *SalesforceConnectionResponse) SetConnectionattributes(v SalesforceConnectionAttributes)`

SetConnectionattributes sets Connectionattributes field to given value.

### HasConnectionattributes

`func (o *SalesforceConnectionResponse) HasConnectionattributes() bool`

HasConnectionattributes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


