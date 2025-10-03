# GetUserResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Msg** | Pointer to **string** |  | [optional] 
**Displaycount** | Pointer to **string** |  | [optional] 
**Totalcount** | Pointer to **string** |  | [optional] 
**ErrorCode** | Pointer to **string** |  | [optional] 
**Userdetails** | Pointer to [**[]User**](User.md) |  | [optional] 

## Methods

### NewGetUserResponse

`func NewGetUserResponse() *GetUserResponse`

NewGetUserResponse instantiates a new GetUserResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetUserResponseWithDefaults

`func NewGetUserResponseWithDefaults() *GetUserResponse`

NewGetUserResponseWithDefaults instantiates a new GetUserResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMsg

`func (o *GetUserResponse) GetMsg() string`

GetMsg returns the Msg field if non-nil, zero value otherwise.

### GetMsgOk

`func (o *GetUserResponse) GetMsgOk() (*string, bool)`

GetMsgOk returns a tuple with the Msg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsg

`func (o *GetUserResponse) SetMsg(v string)`

SetMsg sets Msg field to given value.

### HasMsg

`func (o *GetUserResponse) HasMsg() bool`

HasMsg returns a boolean if a field has been set.

### GetDisplaycount

`func (o *GetUserResponse) GetDisplaycount() string`

GetDisplaycount returns the Displaycount field if non-nil, zero value otherwise.

### GetDisplaycountOk

`func (o *GetUserResponse) GetDisplaycountOk() (*string, bool)`

GetDisplaycountOk returns a tuple with the Displaycount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplaycount

`func (o *GetUserResponse) SetDisplaycount(v string)`

SetDisplaycount sets Displaycount field to given value.

### HasDisplaycount

`func (o *GetUserResponse) HasDisplaycount() bool`

HasDisplaycount returns a boolean if a field has been set.

### GetTotalcount

`func (o *GetUserResponse) GetTotalcount() string`

GetTotalcount returns the Totalcount field if non-nil, zero value otherwise.

### GetTotalcountOk

`func (o *GetUserResponse) GetTotalcountOk() (*string, bool)`

GetTotalcountOk returns a tuple with the Totalcount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalcount

`func (o *GetUserResponse) SetTotalcount(v string)`

SetTotalcount sets Totalcount field to given value.

### HasTotalcount

`func (o *GetUserResponse) HasTotalcount() bool`

HasTotalcount returns a boolean if a field has been set.

### GetErrorCode

`func (o *GetUserResponse) GetErrorCode() string`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *GetUserResponse) GetErrorCodeOk() (*string, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *GetUserResponse) SetErrorCode(v string)`

SetErrorCode sets ErrorCode field to given value.

### HasErrorCode

`func (o *GetUserResponse) HasErrorCode() bool`

HasErrorCode returns a boolean if a field has been set.

### GetUserdetails

`func (o *GetUserResponse) GetUserdetails() []User`

GetUserdetails returns the Userdetails field if non-nil, zero value otherwise.

### GetUserdetailsOk

`func (o *GetUserResponse) GetUserdetailsOk() (*[]User, bool)`

GetUserdetailsOk returns a tuple with the Userdetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserdetails

`func (o *GetUserResponse) SetUserdetails(v []User)`

SetUserdetails sets Userdetails field to given value.

### HasUserdetails

`func (o *GetUserResponse) HasUserdetails() bool`

HasUserdetails returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


