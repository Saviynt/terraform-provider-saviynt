# FetchDynamicAttributesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Msg** | Pointer to **string** | A message indicating the outcome of the operation. | [optional] 
**Errorcode** | Pointer to **string** | An error code where &#39;0&#39; signifies success and &#39;1&#39; signifies an unsuccessful operation. | [optional] 
**Displaycount** | Pointer to **int32** | Total number of records displayed. | [optional] 
**Totalcount** | Pointer to **int32** | Total number of records available. | [optional] 
**Dynamicattributes** | Pointer to [**FetchDynamicAttributesResponseDynamicattributes**](FetchDynamicAttributesResponseDynamicattributes.md) |  | [optional] 

## Methods

### NewFetchDynamicAttributesResponse

`func NewFetchDynamicAttributesResponse() *FetchDynamicAttributesResponse`

NewFetchDynamicAttributesResponse instantiates a new FetchDynamicAttributesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFetchDynamicAttributesResponseWithDefaults

`func NewFetchDynamicAttributesResponseWithDefaults() *FetchDynamicAttributesResponse`

NewFetchDynamicAttributesResponseWithDefaults instantiates a new FetchDynamicAttributesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMsg

`func (o *FetchDynamicAttributesResponse) GetMsg() string`

GetMsg returns the Msg field if non-nil, zero value otherwise.

### GetMsgOk

`func (o *FetchDynamicAttributesResponse) GetMsgOk() (*string, bool)`

GetMsgOk returns a tuple with the Msg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsg

`func (o *FetchDynamicAttributesResponse) SetMsg(v string)`

SetMsg sets Msg field to given value.

### HasMsg

`func (o *FetchDynamicAttributesResponse) HasMsg() bool`

HasMsg returns a boolean if a field has been set.

### GetErrorcode

`func (o *FetchDynamicAttributesResponse) GetErrorcode() string`

GetErrorcode returns the Errorcode field if non-nil, zero value otherwise.

### GetErrorcodeOk

`func (o *FetchDynamicAttributesResponse) GetErrorcodeOk() (*string, bool)`

GetErrorcodeOk returns a tuple with the Errorcode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorcode

`func (o *FetchDynamicAttributesResponse) SetErrorcode(v string)`

SetErrorcode sets Errorcode field to given value.

### HasErrorcode

`func (o *FetchDynamicAttributesResponse) HasErrorcode() bool`

HasErrorcode returns a boolean if a field has been set.

### GetDisplaycount

`func (o *FetchDynamicAttributesResponse) GetDisplaycount() int32`

GetDisplaycount returns the Displaycount field if non-nil, zero value otherwise.

### GetDisplaycountOk

`func (o *FetchDynamicAttributesResponse) GetDisplaycountOk() (*int32, bool)`

GetDisplaycountOk returns a tuple with the Displaycount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplaycount

`func (o *FetchDynamicAttributesResponse) SetDisplaycount(v int32)`

SetDisplaycount sets Displaycount field to given value.

### HasDisplaycount

`func (o *FetchDynamicAttributesResponse) HasDisplaycount() bool`

HasDisplaycount returns a boolean if a field has been set.

### GetTotalcount

`func (o *FetchDynamicAttributesResponse) GetTotalcount() int32`

GetTotalcount returns the Totalcount field if non-nil, zero value otherwise.

### GetTotalcountOk

`func (o *FetchDynamicAttributesResponse) GetTotalcountOk() (*int32, bool)`

GetTotalcountOk returns a tuple with the Totalcount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalcount

`func (o *FetchDynamicAttributesResponse) SetTotalcount(v int32)`

SetTotalcount sets Totalcount field to given value.

### HasTotalcount

`func (o *FetchDynamicAttributesResponse) HasTotalcount() bool`

HasTotalcount returns a boolean if a field has been set.

### GetDynamicattributes

`func (o *FetchDynamicAttributesResponse) GetDynamicattributes() FetchDynamicAttributesResponseDynamicattributes`

GetDynamicattributes returns the Dynamicattributes field if non-nil, zero value otherwise.

### GetDynamicattributesOk

`func (o *FetchDynamicAttributesResponse) GetDynamicattributesOk() (*FetchDynamicAttributesResponseDynamicattributes, bool)`

GetDynamicattributesOk returns a tuple with the Dynamicattributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDynamicattributes

`func (o *FetchDynamicAttributesResponse) SetDynamicattributes(v FetchDynamicAttributesResponseDynamicattributes)`

SetDynamicattributes sets Dynamicattributes field to given value.

### HasDynamicattributes

`func (o *FetchDynamicAttributesResponse) HasDynamicattributes() bool`

HasDynamicattributes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


