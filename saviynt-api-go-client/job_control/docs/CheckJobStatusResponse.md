# CheckJobStatusResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Msg** | **string** | Message indicating the result of the operation | 
**ErrorCode** | **string** | Error code if the operation failed | 

## Methods

### NewCheckJobStatusResponse

`func NewCheckJobStatusResponse(msg string, errorCode string, ) *CheckJobStatusResponse`

NewCheckJobStatusResponse instantiates a new CheckJobStatusResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCheckJobStatusResponseWithDefaults

`func NewCheckJobStatusResponseWithDefaults() *CheckJobStatusResponse`

NewCheckJobStatusResponseWithDefaults instantiates a new CheckJobStatusResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMsg

`func (o *CheckJobStatusResponse) GetMsg() string`

GetMsg returns the Msg field if non-nil, zero value otherwise.

### GetMsgOk

`func (o *CheckJobStatusResponse) GetMsgOk() (*string, bool)`

GetMsgOk returns a tuple with the Msg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsg

`func (o *CheckJobStatusResponse) SetMsg(v string)`

SetMsg sets Msg field to given value.


### GetErrorCode

`func (o *CheckJobStatusResponse) GetErrorCode() string`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *CheckJobStatusResponse) GetErrorCodeOk() (*string, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *CheckJobStatusResponse) SetErrorCode(v string)`

SetErrorCode sets ErrorCode field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


