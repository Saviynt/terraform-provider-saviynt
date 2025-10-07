# GetPrivilegeListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Msg** | Pointer to **string** | Response message | [optional] 
**ErrorCode** | Pointer to **string** | Error code | [optional] 
**DisplayCount** | Pointer to **int32** | Number of results returned | [optional] 
**TotalCount** | Pointer to **int32** | Total number of privileges for the given filter | [optional] 
**PrivilegeDetails** | Pointer to [**[]GetPrivilegeDetail**](GetPrivilegeDetail.md) |  | [optional] 

## Methods

### NewGetPrivilegeListResponse

`func NewGetPrivilegeListResponse() *GetPrivilegeListResponse`

NewGetPrivilegeListResponse instantiates a new GetPrivilegeListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetPrivilegeListResponseWithDefaults

`func NewGetPrivilegeListResponseWithDefaults() *GetPrivilegeListResponse`

NewGetPrivilegeListResponseWithDefaults instantiates a new GetPrivilegeListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMsg

`func (o *GetPrivilegeListResponse) GetMsg() string`

GetMsg returns the Msg field if non-nil, zero value otherwise.

### GetMsgOk

`func (o *GetPrivilegeListResponse) GetMsgOk() (*string, bool)`

GetMsgOk returns a tuple with the Msg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsg

`func (o *GetPrivilegeListResponse) SetMsg(v string)`

SetMsg sets Msg field to given value.

### HasMsg

`func (o *GetPrivilegeListResponse) HasMsg() bool`

HasMsg returns a boolean if a field has been set.

### GetErrorCode

`func (o *GetPrivilegeListResponse) GetErrorCode() string`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *GetPrivilegeListResponse) GetErrorCodeOk() (*string, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *GetPrivilegeListResponse) SetErrorCode(v string)`

SetErrorCode sets ErrorCode field to given value.

### HasErrorCode

`func (o *GetPrivilegeListResponse) HasErrorCode() bool`

HasErrorCode returns a boolean if a field has been set.

### GetDisplayCount

`func (o *GetPrivilegeListResponse) GetDisplayCount() int32`

GetDisplayCount returns the DisplayCount field if non-nil, zero value otherwise.

### GetDisplayCountOk

`func (o *GetPrivilegeListResponse) GetDisplayCountOk() (*int32, bool)`

GetDisplayCountOk returns a tuple with the DisplayCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayCount

`func (o *GetPrivilegeListResponse) SetDisplayCount(v int32)`

SetDisplayCount sets DisplayCount field to given value.

### HasDisplayCount

`func (o *GetPrivilegeListResponse) HasDisplayCount() bool`

HasDisplayCount returns a boolean if a field has been set.

### GetTotalCount

`func (o *GetPrivilegeListResponse) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *GetPrivilegeListResponse) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *GetPrivilegeListResponse) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.

### HasTotalCount

`func (o *GetPrivilegeListResponse) HasTotalCount() bool`

HasTotalCount returns a boolean if a field has been set.

### GetPrivilegeDetails

`func (o *GetPrivilegeListResponse) GetPrivilegeDetails() []GetPrivilegeDetail`

GetPrivilegeDetails returns the PrivilegeDetails field if non-nil, zero value otherwise.

### GetPrivilegeDetailsOk

`func (o *GetPrivilegeListResponse) GetPrivilegeDetailsOk() (*[]GetPrivilegeDetail, bool)`

GetPrivilegeDetailsOk returns a tuple with the PrivilegeDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivilegeDetails

`func (o *GetPrivilegeListResponse) SetPrivilegeDetails(v []GetPrivilegeDetail)`

SetPrivilegeDetails sets PrivilegeDetails field to given value.

### HasPrivilegeDetails

`func (o *GetPrivilegeListResponse) HasPrivilegeDetails() bool`

HasPrivilegeDetails returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


