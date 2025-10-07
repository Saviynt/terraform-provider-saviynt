# RunJobTriggerResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Msg** | **string** | Message indicating the result of the operation | 
**ErrorCode** | **string** | Error code if the operation failed | 
**Timestamp** | Pointer to **string** | Timestamp of when the job was triggered | [optional] 

## Methods

### NewRunJobTriggerResponse

`func NewRunJobTriggerResponse(msg string, errorCode string, ) *RunJobTriggerResponse`

NewRunJobTriggerResponse instantiates a new RunJobTriggerResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRunJobTriggerResponseWithDefaults

`func NewRunJobTriggerResponseWithDefaults() *RunJobTriggerResponse`

NewRunJobTriggerResponseWithDefaults instantiates a new RunJobTriggerResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMsg

`func (o *RunJobTriggerResponse) GetMsg() string`

GetMsg returns the Msg field if non-nil, zero value otherwise.

### GetMsgOk

`func (o *RunJobTriggerResponse) GetMsgOk() (*string, bool)`

GetMsgOk returns a tuple with the Msg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsg

`func (o *RunJobTriggerResponse) SetMsg(v string)`

SetMsg sets Msg field to given value.


### GetErrorCode

`func (o *RunJobTriggerResponse) GetErrorCode() string`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *RunJobTriggerResponse) GetErrorCodeOk() (*string, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *RunJobTriggerResponse) SetErrorCode(v string)`

SetErrorCode sets ErrorCode field to given value.


### GetTimestamp

`func (o *RunJobTriggerResponse) GetTimestamp() string`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *RunJobTriggerResponse) GetTimestampOk() (*string, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *RunJobTriggerResponse) SetTimestamp(v string)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *RunJobTriggerResponse) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


