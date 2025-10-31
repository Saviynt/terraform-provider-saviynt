# FetchExistingDelegatesListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DelegateList** | [**[]Delegate**](Delegate.md) |  | 
**Msg** | **string** |  | 
**ErrorCode** | **string** |  | 
**TotalCount** | Pointer to **string** |  | [optional] 
**DelegateCount** | Pointer to **int32** |  | [optional] 

## Methods

### NewFetchExistingDelegatesListResponse

`func NewFetchExistingDelegatesListResponse(delegateList []Delegate, msg string, errorCode string, ) *FetchExistingDelegatesListResponse`

NewFetchExistingDelegatesListResponse instantiates a new FetchExistingDelegatesListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFetchExistingDelegatesListResponseWithDefaults

`func NewFetchExistingDelegatesListResponseWithDefaults() *FetchExistingDelegatesListResponse`

NewFetchExistingDelegatesListResponseWithDefaults instantiates a new FetchExistingDelegatesListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDelegateList

`func (o *FetchExistingDelegatesListResponse) GetDelegateList() []Delegate`

GetDelegateList returns the DelegateList field if non-nil, zero value otherwise.

### GetDelegateListOk

`func (o *FetchExistingDelegatesListResponse) GetDelegateListOk() (*[]Delegate, bool)`

GetDelegateListOk returns a tuple with the DelegateList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelegateList

`func (o *FetchExistingDelegatesListResponse) SetDelegateList(v []Delegate)`

SetDelegateList sets DelegateList field to given value.


### GetMsg

`func (o *FetchExistingDelegatesListResponse) GetMsg() string`

GetMsg returns the Msg field if non-nil, zero value otherwise.

### GetMsgOk

`func (o *FetchExistingDelegatesListResponse) GetMsgOk() (*string, bool)`

GetMsgOk returns a tuple with the Msg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsg

`func (o *FetchExistingDelegatesListResponse) SetMsg(v string)`

SetMsg sets Msg field to given value.


### GetErrorCode

`func (o *FetchExistingDelegatesListResponse) GetErrorCode() string`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *FetchExistingDelegatesListResponse) GetErrorCodeOk() (*string, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *FetchExistingDelegatesListResponse) SetErrorCode(v string)`

SetErrorCode sets ErrorCode field to given value.


### GetTotalCount

`func (o *FetchExistingDelegatesListResponse) GetTotalCount() string`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *FetchExistingDelegatesListResponse) GetTotalCountOk() (*string, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *FetchExistingDelegatesListResponse) SetTotalCount(v string)`

SetTotalCount sets TotalCount field to given value.

### HasTotalCount

`func (o *FetchExistingDelegatesListResponse) HasTotalCount() bool`

HasTotalCount returns a boolean if a field has been set.

### GetDelegateCount

`func (o *FetchExistingDelegatesListResponse) GetDelegateCount() int32`

GetDelegateCount returns the DelegateCount field if non-nil, zero value otherwise.

### GetDelegateCountOk

`func (o *FetchExistingDelegatesListResponse) GetDelegateCountOk() (*int32, bool)`

GetDelegateCountOk returns a tuple with the DelegateCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelegateCount

`func (o *FetchExistingDelegatesListResponse) SetDelegateCount(v int32)`

SetDelegateCount sets DelegateCount field to given value.

### HasDelegateCount

`func (o *FetchExistingDelegatesListResponse) HasDelegateCount() bool`

HasDelegateCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


