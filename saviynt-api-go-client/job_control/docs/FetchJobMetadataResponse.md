# FetchJobMetadataResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Msg** | **string** | Message indicating the result of the operation | 
**ErrorCode** | **int32** | Error code if the operation failed | 
**Result** | Pointer to **map[string]interface{}** | Metadata of the job run | [optional] 

## Methods

### NewFetchJobMetadataResponse

`func NewFetchJobMetadataResponse(msg string, errorCode int32, ) *FetchJobMetadataResponse`

NewFetchJobMetadataResponse instantiates a new FetchJobMetadataResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFetchJobMetadataResponseWithDefaults

`func NewFetchJobMetadataResponseWithDefaults() *FetchJobMetadataResponse`

NewFetchJobMetadataResponseWithDefaults instantiates a new FetchJobMetadataResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMsg

`func (o *FetchJobMetadataResponse) GetMsg() string`

GetMsg returns the Msg field if non-nil, zero value otherwise.

### GetMsgOk

`func (o *FetchJobMetadataResponse) GetMsgOk() (*string, bool)`

GetMsgOk returns a tuple with the Msg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsg

`func (o *FetchJobMetadataResponse) SetMsg(v string)`

SetMsg sets Msg field to given value.


### GetErrorCode

`func (o *FetchJobMetadataResponse) GetErrorCode() int32`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *FetchJobMetadataResponse) GetErrorCodeOk() (*int32, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *FetchJobMetadataResponse) SetErrorCode(v int32)`

SetErrorCode sets ErrorCode field to given value.


### GetResult

`func (o *FetchJobMetadataResponse) GetResult() map[string]interface{}`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *FetchJobMetadataResponse) GetResultOk() (*map[string]interface{}, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *FetchJobMetadataResponse) SetResult(v map[string]interface{})`

SetResult sets Result field to given value.

### HasResult

`func (o *FetchJobMetadataResponse) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


