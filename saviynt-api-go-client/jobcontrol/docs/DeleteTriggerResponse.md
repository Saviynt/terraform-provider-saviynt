# DeleteTriggerResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Msg** | Pointer to **string** |  | [optional] 
**ErrorCode** | Pointer to **string** |  | [optional] 

## Methods

### NewDeleteTriggerResponse

`func NewDeleteTriggerResponse() *DeleteTriggerResponse`

NewDeleteTriggerResponse instantiates a new DeleteTriggerResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeleteTriggerResponseWithDefaults

`func NewDeleteTriggerResponseWithDefaults() *DeleteTriggerResponse`

NewDeleteTriggerResponseWithDefaults instantiates a new DeleteTriggerResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMsg

`func (o *DeleteTriggerResponse) GetMsg() string`

GetMsg returns the Msg field if non-nil, zero value otherwise.

### GetMsgOk

`func (o *DeleteTriggerResponse) GetMsgOk() (*string, bool)`

GetMsgOk returns a tuple with the Msg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsg

`func (o *DeleteTriggerResponse) SetMsg(v string)`

SetMsg sets Msg field to given value.

### HasMsg

`func (o *DeleteTriggerResponse) HasMsg() bool`

HasMsg returns a boolean if a field has been set.

### GetErrorCode

`func (o *DeleteTriggerResponse) GetErrorCode() string`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *DeleteTriggerResponse) GetErrorCodeOk() (*string, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *DeleteTriggerResponse) SetErrorCode(v string)`

SetErrorCode sets ErrorCode field to given value.

### HasErrorCode

`func (o *DeleteTriggerResponse) HasErrorCode() bool`

HasErrorCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


