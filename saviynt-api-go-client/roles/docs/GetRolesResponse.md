# GetRolesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Msg** | Pointer to **string** |  | [optional] 
**DisplayCount** | Pointer to **int32** |  | [optional] 
**ErrorCode** | Pointer to **string** |  | [optional] 
**TotalCount** | Pointer to **int32** |  | [optional] 
**Roledetails** | Pointer to [**[]GetRoleDetailsResponse**](GetRoleDetailsResponse.md) |  | [optional] 

## Methods

### NewGetRolesResponse

`func NewGetRolesResponse() *GetRolesResponse`

NewGetRolesResponse instantiates a new GetRolesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetRolesResponseWithDefaults

`func NewGetRolesResponseWithDefaults() *GetRolesResponse`

NewGetRolesResponseWithDefaults instantiates a new GetRolesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMsg

`func (o *GetRolesResponse) GetMsg() string`

GetMsg returns the Msg field if non-nil, zero value otherwise.

### GetMsgOk

`func (o *GetRolesResponse) GetMsgOk() (*string, bool)`

GetMsgOk returns a tuple with the Msg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsg

`func (o *GetRolesResponse) SetMsg(v string)`

SetMsg sets Msg field to given value.

### HasMsg

`func (o *GetRolesResponse) HasMsg() bool`

HasMsg returns a boolean if a field has been set.

### GetDisplayCount

`func (o *GetRolesResponse) GetDisplayCount() int32`

GetDisplayCount returns the DisplayCount field if non-nil, zero value otherwise.

### GetDisplayCountOk

`func (o *GetRolesResponse) GetDisplayCountOk() (*int32, bool)`

GetDisplayCountOk returns a tuple with the DisplayCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayCount

`func (o *GetRolesResponse) SetDisplayCount(v int32)`

SetDisplayCount sets DisplayCount field to given value.

### HasDisplayCount

`func (o *GetRolesResponse) HasDisplayCount() bool`

HasDisplayCount returns a boolean if a field has been set.

### GetErrorCode

`func (o *GetRolesResponse) GetErrorCode() string`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *GetRolesResponse) GetErrorCodeOk() (*string, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *GetRolesResponse) SetErrorCode(v string)`

SetErrorCode sets ErrorCode field to given value.

### HasErrorCode

`func (o *GetRolesResponse) HasErrorCode() bool`

HasErrorCode returns a boolean if a field has been set.

### GetTotalCount

`func (o *GetRolesResponse) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *GetRolesResponse) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *GetRolesResponse) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.

### HasTotalCount

`func (o *GetRolesResponse) HasTotalCount() bool`

HasTotalCount returns a boolean if a field has been set.

### GetRoledetails

`func (o *GetRolesResponse) GetRoledetails() []GetRoleDetailsResponse`

GetRoledetails returns the Roledetails field if non-nil, zero value otherwise.

### GetRoledetailsOk

`func (o *GetRolesResponse) GetRoledetailsOk() (*[]GetRoleDetailsResponse, bool)`

GetRoledetailsOk returns a tuple with the Roledetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoledetails

`func (o *GetRolesResponse) SetRoledetails(v []GetRoleDetailsResponse)`

SetRoledetails sets Roledetails field to given value.

### HasRoledetails

`func (o *GetRolesResponse) HasRoledetails() bool`

HasRoledetails returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


