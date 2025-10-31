# CreateOrUpdateResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConnectionKey** | Pointer to **int32** |  | [optional] 
**Msg** | Pointer to **string** |  | [optional] 
**ErrorCode** | Pointer to **string** |  | [optional] 

## Methods

### NewCreateOrUpdateResponse

`func NewCreateOrUpdateResponse() *CreateOrUpdateResponse`

NewCreateOrUpdateResponse instantiates a new CreateOrUpdateResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateOrUpdateResponseWithDefaults

`func NewCreateOrUpdateResponseWithDefaults() *CreateOrUpdateResponse`

NewCreateOrUpdateResponseWithDefaults instantiates a new CreateOrUpdateResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConnectionKey

`func (o *CreateOrUpdateResponse) GetConnectionKey() int32`

GetConnectionKey returns the ConnectionKey field if non-nil, zero value otherwise.

### GetConnectionKeyOk

`func (o *CreateOrUpdateResponse) GetConnectionKeyOk() (*int32, bool)`

GetConnectionKeyOk returns a tuple with the ConnectionKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionKey

`func (o *CreateOrUpdateResponse) SetConnectionKey(v int32)`

SetConnectionKey sets ConnectionKey field to given value.

### HasConnectionKey

`func (o *CreateOrUpdateResponse) HasConnectionKey() bool`

HasConnectionKey returns a boolean if a field has been set.

### GetMsg

`func (o *CreateOrUpdateResponse) GetMsg() string`

GetMsg returns the Msg field if non-nil, zero value otherwise.

### GetMsgOk

`func (o *CreateOrUpdateResponse) GetMsgOk() (*string, bool)`

GetMsgOk returns a tuple with the Msg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsg

`func (o *CreateOrUpdateResponse) SetMsg(v string)`

SetMsg sets Msg field to given value.

### HasMsg

`func (o *CreateOrUpdateResponse) HasMsg() bool`

HasMsg returns a boolean if a field has been set.

### GetErrorCode

`func (o *CreateOrUpdateResponse) GetErrorCode() string`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *CreateOrUpdateResponse) GetErrorCodeOk() (*string, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *CreateOrUpdateResponse) SetErrorCode(v string)`

SetErrorCode sets ErrorCode field to given value.

### HasErrorCode

`func (o *CreateOrUpdateResponse) HasErrorCode() bool`

HasErrorCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


