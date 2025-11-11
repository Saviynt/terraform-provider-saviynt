# ExportTransportPackageResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Msg** | Pointer to **string** | Response message | [optional] 
**FileName** | Pointer to **string** | Name of the generated transport package file | [optional] 
**MsgDescription** | Pointer to **string** | Detailed description of the response | [optional] 
**Errorcode** | Pointer to **int32** | Error code (0 for success) | [optional] 

## Methods

### NewExportTransportPackageResponse

`func NewExportTransportPackageResponse() *ExportTransportPackageResponse`

NewExportTransportPackageResponse instantiates a new ExportTransportPackageResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExportTransportPackageResponseWithDefaults

`func NewExportTransportPackageResponseWithDefaults() *ExportTransportPackageResponse`

NewExportTransportPackageResponseWithDefaults instantiates a new ExportTransportPackageResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMsg

`func (o *ExportTransportPackageResponse) GetMsg() string`

GetMsg returns the Msg field if non-nil, zero value otherwise.

### GetMsgOk

`func (o *ExportTransportPackageResponse) GetMsgOk() (*string, bool)`

GetMsgOk returns a tuple with the Msg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsg

`func (o *ExportTransportPackageResponse) SetMsg(v string)`

SetMsg sets Msg field to given value.

### HasMsg

`func (o *ExportTransportPackageResponse) HasMsg() bool`

HasMsg returns a boolean if a field has been set.

### GetFileName

`func (o *ExportTransportPackageResponse) GetFileName() string`

GetFileName returns the FileName field if non-nil, zero value otherwise.

### GetFileNameOk

`func (o *ExportTransportPackageResponse) GetFileNameOk() (*string, bool)`

GetFileNameOk returns a tuple with the FileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileName

`func (o *ExportTransportPackageResponse) SetFileName(v string)`

SetFileName sets FileName field to given value.

### HasFileName

`func (o *ExportTransportPackageResponse) HasFileName() bool`

HasFileName returns a boolean if a field has been set.

### GetMsgDescription

`func (o *ExportTransportPackageResponse) GetMsgDescription() string`

GetMsgDescription returns the MsgDescription field if non-nil, zero value otherwise.

### GetMsgDescriptionOk

`func (o *ExportTransportPackageResponse) GetMsgDescriptionOk() (*string, bool)`

GetMsgDescriptionOk returns a tuple with the MsgDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsgDescription

`func (o *ExportTransportPackageResponse) SetMsgDescription(v string)`

SetMsgDescription sets MsgDescription field to given value.

### HasMsgDescription

`func (o *ExportTransportPackageResponse) HasMsgDescription() bool`

HasMsgDescription returns a boolean if a field has been set.

### GetErrorcode

`func (o *ExportTransportPackageResponse) GetErrorcode() int32`

GetErrorcode returns the Errorcode field if non-nil, zero value otherwise.

### GetErrorcodeOk

`func (o *ExportTransportPackageResponse) GetErrorcodeOk() (*int32, bool)`

GetErrorcodeOk returns a tuple with the Errorcode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorcode

`func (o *ExportTransportPackageResponse) SetErrorcode(v int32)`

SetErrorcode sets Errorcode field to given value.

### HasErrorcode

`func (o *ExportTransportPackageResponse) HasErrorcode() bool`

HasErrorcode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


