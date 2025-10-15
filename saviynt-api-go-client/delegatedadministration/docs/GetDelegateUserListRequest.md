# GetDelegateUserListRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Parentusername** | **string** |  | 
**Max** | Pointer to **int32** |  | [optional] 
**Offset** | Pointer to **int32** |  | [optional] 
**SearchCriteria** | Pointer to **string** | example \&quot;t*\&quot; or \&quot;te\&quot; or \&quot;test\&quot; to search in username or firstname or lastname for the delegate user | [optional] 

## Methods

### NewGetDelegateUserListRequest

`func NewGetDelegateUserListRequest(parentusername string, ) *GetDelegateUserListRequest`

NewGetDelegateUserListRequest instantiates a new GetDelegateUserListRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetDelegateUserListRequestWithDefaults

`func NewGetDelegateUserListRequestWithDefaults() *GetDelegateUserListRequest`

NewGetDelegateUserListRequestWithDefaults instantiates a new GetDelegateUserListRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetParentusername

`func (o *GetDelegateUserListRequest) GetParentusername() string`

GetParentusername returns the Parentusername field if non-nil, zero value otherwise.

### GetParentusernameOk

`func (o *GetDelegateUserListRequest) GetParentusernameOk() (*string, bool)`

GetParentusernameOk returns a tuple with the Parentusername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentusername

`func (o *GetDelegateUserListRequest) SetParentusername(v string)`

SetParentusername sets Parentusername field to given value.


### GetMax

`func (o *GetDelegateUserListRequest) GetMax() int32`

GetMax returns the Max field if non-nil, zero value otherwise.

### GetMaxOk

`func (o *GetDelegateUserListRequest) GetMaxOk() (*int32, bool)`

GetMaxOk returns a tuple with the Max field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMax

`func (o *GetDelegateUserListRequest) SetMax(v int32)`

SetMax sets Max field to given value.

### HasMax

`func (o *GetDelegateUserListRequest) HasMax() bool`

HasMax returns a boolean if a field has been set.

### GetOffset

`func (o *GetDelegateUserListRequest) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *GetDelegateUserListRequest) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *GetDelegateUserListRequest) SetOffset(v int32)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *GetDelegateUserListRequest) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### GetSearchCriteria

`func (o *GetDelegateUserListRequest) GetSearchCriteria() string`

GetSearchCriteria returns the SearchCriteria field if non-nil, zero value otherwise.

### GetSearchCriteriaOk

`func (o *GetDelegateUserListRequest) GetSearchCriteriaOk() (*string, bool)`

GetSearchCriteriaOk returns a tuple with the SearchCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchCriteria

`func (o *GetDelegateUserListRequest) SetSearchCriteria(v string)`

SetSearchCriteria sets SearchCriteria field to given value.

### HasSearchCriteria

`func (o *GetDelegateUserListRequest) HasSearchCriteria() bool`

HasSearchCriteria returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


