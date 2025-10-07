# CreateDelegateResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Msg** | **string** |  | 
**Delegatekey** | Pointer to **string** |  | [optional] 
**ErrorCode** | **string** |  | 

## Methods

### NewCreateDelegateResponse

`func NewCreateDelegateResponse(msg string, errorCode string, ) *CreateDelegateResponse`

NewCreateDelegateResponse instantiates a new CreateDelegateResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateDelegateResponseWithDefaults

`func NewCreateDelegateResponseWithDefaults() *CreateDelegateResponse`

NewCreateDelegateResponseWithDefaults instantiates a new CreateDelegateResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMsg

`func (o *CreateDelegateResponse) GetMsg() string`

GetMsg returns the Msg field if non-nil, zero value otherwise.

### GetMsgOk

`func (o *CreateDelegateResponse) GetMsgOk() (*string, bool)`

GetMsgOk returns a tuple with the Msg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsg

`func (o *CreateDelegateResponse) SetMsg(v string)`

SetMsg sets Msg field to given value.


### GetDelegatekey

`func (o *CreateDelegateResponse) GetDelegatekey() string`

GetDelegatekey returns the Delegatekey field if non-nil, zero value otherwise.

### GetDelegatekeyOk

`func (o *CreateDelegateResponse) GetDelegatekeyOk() (*string, bool)`

GetDelegatekeyOk returns a tuple with the Delegatekey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelegatekey

`func (o *CreateDelegateResponse) SetDelegatekey(v string)`

SetDelegatekey sets Delegatekey field to given value.

### HasDelegatekey

`func (o *CreateDelegateResponse) HasDelegatekey() bool`

HasDelegatekey returns a boolean if a field has been set.

### GetErrorCode

`func (o *CreateDelegateResponse) GetErrorCode() string`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *CreateDelegateResponse) GetErrorCodeOk() (*string, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *CreateDelegateResponse) SetErrorCode(v string)`

SetErrorCode sets ErrorCode field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


