# ImportTransportPackageResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Msg** | **string** |  | 
**RequestId** | Pointer to **string** |  | [optional] 
**MsgDescription** | Pointer to **string** |  | [optional] 
**Errorcode** | **int32** |  | 

## Methods

### NewImportTransportPackageResponse

`func NewImportTransportPackageResponse(msg string, errorcode int32, ) *ImportTransportPackageResponse`

NewImportTransportPackageResponse instantiates a new ImportTransportPackageResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImportTransportPackageResponseWithDefaults

`func NewImportTransportPackageResponseWithDefaults() *ImportTransportPackageResponse`

NewImportTransportPackageResponseWithDefaults instantiates a new ImportTransportPackageResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMsg

`func (o *ImportTransportPackageResponse) GetMsg() string`

GetMsg returns the Msg field if non-nil, zero value otherwise.

### GetMsgOk

`func (o *ImportTransportPackageResponse) GetMsgOk() (*string, bool)`

GetMsgOk returns a tuple with the Msg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsg

`func (o *ImportTransportPackageResponse) SetMsg(v string)`

SetMsg sets Msg field to given value.


### GetRequestId

`func (o *ImportTransportPackageResponse) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *ImportTransportPackageResponse) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *ImportTransportPackageResponse) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *ImportTransportPackageResponse) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.

### GetMsgDescription

`func (o *ImportTransportPackageResponse) GetMsgDescription() string`

GetMsgDescription returns the MsgDescription field if non-nil, zero value otherwise.

### GetMsgDescriptionOk

`func (o *ImportTransportPackageResponse) GetMsgDescriptionOk() (*string, bool)`

GetMsgDescriptionOk returns a tuple with the MsgDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsgDescription

`func (o *ImportTransportPackageResponse) SetMsgDescription(v string)`

SetMsgDescription sets MsgDescription field to given value.

### HasMsgDescription

`func (o *ImportTransportPackageResponse) HasMsgDescription() bool`

HasMsgDescription returns a boolean if a field has been set.

### GetErrorcode

`func (o *ImportTransportPackageResponse) GetErrorcode() int32`

GetErrorcode returns the Errorcode field if non-nil, zero value otherwise.

### GetErrorcodeOk

`func (o *ImportTransportPackageResponse) GetErrorcodeOk() (*int32, bool)`

GetErrorcodeOk returns a tuple with the Errorcode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorcode

`func (o *ImportTransportPackageResponse) SetErrorcode(v int32)`

SetErrorcode sets Errorcode field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


