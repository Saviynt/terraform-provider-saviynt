# TransportPackageStatusResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Msg** | Pointer to **string** | Response message | [optional] 
**MsgDescription** | Pointer to **string** | Detailed description of the response | [optional] 
**ErrorCode** | Pointer to **int32** | Error code (0 for success) | [optional] 

## Methods

### NewTransportPackageStatusResponse

`func NewTransportPackageStatusResponse() *TransportPackageStatusResponse`

NewTransportPackageStatusResponse instantiates a new TransportPackageStatusResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransportPackageStatusResponseWithDefaults

`func NewTransportPackageStatusResponseWithDefaults() *TransportPackageStatusResponse`

NewTransportPackageStatusResponseWithDefaults instantiates a new TransportPackageStatusResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMsg

`func (o *TransportPackageStatusResponse) GetMsg() string`

GetMsg returns the Msg field if non-nil, zero value otherwise.

### GetMsgOk

`func (o *TransportPackageStatusResponse) GetMsgOk() (*string, bool)`

GetMsgOk returns a tuple with the Msg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsg

`func (o *TransportPackageStatusResponse) SetMsg(v string)`

SetMsg sets Msg field to given value.

### HasMsg

`func (o *TransportPackageStatusResponse) HasMsg() bool`

HasMsg returns a boolean if a field has been set.

### GetMsgDescription

`func (o *TransportPackageStatusResponse) GetMsgDescription() string`

GetMsgDescription returns the MsgDescription field if non-nil, zero value otherwise.

### GetMsgDescriptionOk

`func (o *TransportPackageStatusResponse) GetMsgDescriptionOk() (*string, bool)`

GetMsgDescriptionOk returns a tuple with the MsgDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsgDescription

`func (o *TransportPackageStatusResponse) SetMsgDescription(v string)`

SetMsgDescription sets MsgDescription field to given value.

### HasMsgDescription

`func (o *TransportPackageStatusResponse) HasMsgDescription() bool`

HasMsgDescription returns a boolean if a field has been set.

### GetErrorCode

`func (o *TransportPackageStatusResponse) GetErrorCode() int32`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *TransportPackageStatusResponse) GetErrorCodeOk() (*int32, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *TransportPackageStatusResponse) SetErrorCode(v int32)`

SetErrorCode sets ErrorCode field to given value.

### HasErrorCode

`func (o *TransportPackageStatusResponse) HasErrorCode() bool`

HasErrorCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


