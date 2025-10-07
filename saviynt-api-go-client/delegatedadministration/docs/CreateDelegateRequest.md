# CreateDelegateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserName** | **string** | this is user who is creating the delegation | 
**Name** | **string** | name of delegate. | 
**Delegateusername** | **string** | This is the user who should be assigned as the delegate of the parent user. It must match &#x60;username&#x60; in &#x60;GetDelegateUserList&#x60; API response. | 
**Delegatestartdate** | **string** | in format MMDDYYYY | 
**Delegateenddate** | **string** | in format MMDDYYYY | 
**Parentusername** | Pointer to **string** | this is the parent username, if not passed it will consider &#x60;userName&#x60; as parentusername | [optional] 
**Description** | Pointer to **string** |  | [optional] 

## Methods

### NewCreateDelegateRequest

`func NewCreateDelegateRequest(userName string, name string, delegateusername string, delegatestartdate string, delegateenddate string, ) *CreateDelegateRequest`

NewCreateDelegateRequest instantiates a new CreateDelegateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateDelegateRequestWithDefaults

`func NewCreateDelegateRequestWithDefaults() *CreateDelegateRequest`

NewCreateDelegateRequestWithDefaults instantiates a new CreateDelegateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserName

`func (o *CreateDelegateRequest) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *CreateDelegateRequest) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *CreateDelegateRequest) SetUserName(v string)`

SetUserName sets UserName field to given value.


### GetName

`func (o *CreateDelegateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateDelegateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateDelegateRequest) SetName(v string)`

SetName sets Name field to given value.


### GetDelegateusername

`func (o *CreateDelegateRequest) GetDelegateusername() string`

GetDelegateusername returns the Delegateusername field if non-nil, zero value otherwise.

### GetDelegateusernameOk

`func (o *CreateDelegateRequest) GetDelegateusernameOk() (*string, bool)`

GetDelegateusernameOk returns a tuple with the Delegateusername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelegateusername

`func (o *CreateDelegateRequest) SetDelegateusername(v string)`

SetDelegateusername sets Delegateusername field to given value.


### GetDelegatestartdate

`func (o *CreateDelegateRequest) GetDelegatestartdate() string`

GetDelegatestartdate returns the Delegatestartdate field if non-nil, zero value otherwise.

### GetDelegatestartdateOk

`func (o *CreateDelegateRequest) GetDelegatestartdateOk() (*string, bool)`

GetDelegatestartdateOk returns a tuple with the Delegatestartdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelegatestartdate

`func (o *CreateDelegateRequest) SetDelegatestartdate(v string)`

SetDelegatestartdate sets Delegatestartdate field to given value.


### GetDelegateenddate

`func (o *CreateDelegateRequest) GetDelegateenddate() string`

GetDelegateenddate returns the Delegateenddate field if non-nil, zero value otherwise.

### GetDelegateenddateOk

`func (o *CreateDelegateRequest) GetDelegateenddateOk() (*string, bool)`

GetDelegateenddateOk returns a tuple with the Delegateenddate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelegateenddate

`func (o *CreateDelegateRequest) SetDelegateenddate(v string)`

SetDelegateenddate sets Delegateenddate field to given value.


### GetParentusername

`func (o *CreateDelegateRequest) GetParentusername() string`

GetParentusername returns the Parentusername field if non-nil, zero value otherwise.

### GetParentusernameOk

`func (o *CreateDelegateRequest) GetParentusernameOk() (*string, bool)`

GetParentusernameOk returns a tuple with the Parentusername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentusername

`func (o *CreateDelegateRequest) SetParentusername(v string)`

SetParentusername sets Parentusername field to given value.

### HasParentusername

`func (o *CreateDelegateRequest) HasParentusername() bool`

HasParentusername returns a boolean if a field has been set.

### GetDescription

`func (o *CreateDelegateRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateDelegateRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateDelegateRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateDelegateRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


