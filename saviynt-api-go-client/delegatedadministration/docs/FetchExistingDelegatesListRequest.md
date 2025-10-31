# FetchExistingDelegatesListRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserName** | **string** | this is the parentusername | 
**Max** | Pointer to **int32** |  | [optional] 
**Offset** | Pointer to **int32** |  | [optional] 
**Status** | Pointer to **string** | values can be ACTIVE / INACTIVE | [optional] 

## Methods

### NewFetchExistingDelegatesListRequest

`func NewFetchExistingDelegatesListRequest(userName string, ) *FetchExistingDelegatesListRequest`

NewFetchExistingDelegatesListRequest instantiates a new FetchExistingDelegatesListRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFetchExistingDelegatesListRequestWithDefaults

`func NewFetchExistingDelegatesListRequestWithDefaults() *FetchExistingDelegatesListRequest`

NewFetchExistingDelegatesListRequestWithDefaults instantiates a new FetchExistingDelegatesListRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserName

`func (o *FetchExistingDelegatesListRequest) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *FetchExistingDelegatesListRequest) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *FetchExistingDelegatesListRequest) SetUserName(v string)`

SetUserName sets UserName field to given value.


### GetMax

`func (o *FetchExistingDelegatesListRequest) GetMax() int32`

GetMax returns the Max field if non-nil, zero value otherwise.

### GetMaxOk

`func (o *FetchExistingDelegatesListRequest) GetMaxOk() (*int32, bool)`

GetMaxOk returns a tuple with the Max field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMax

`func (o *FetchExistingDelegatesListRequest) SetMax(v int32)`

SetMax sets Max field to given value.

### HasMax

`func (o *FetchExistingDelegatesListRequest) HasMax() bool`

HasMax returns a boolean if a field has been set.

### GetOffset

`func (o *FetchExistingDelegatesListRequest) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *FetchExistingDelegatesListRequest) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *FetchExistingDelegatesListRequest) SetOffset(v int32)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *FetchExistingDelegatesListRequest) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### GetStatus

`func (o *FetchExistingDelegatesListRequest) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *FetchExistingDelegatesListRequest) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *FetchExistingDelegatesListRequest) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *FetchExistingDelegatesListRequest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


