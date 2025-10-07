# EditDelegateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | **string** | this is the delegatekey | 
**UserName** | **string** | this is user who is updating the delegation | 
**Name** | **string** | name of delegate | 
**Delegateusername** | **string** | this is the user who should be assigned as the delegate of the parent use | 
**Delegatestartdate** | **string** | in format MMDDYYYY | 
**Delegateenddate** | **string** | in format MMDDYYYY | 
**Parentusername** | Pointer to **string** | this is the parent username, if not passed it will consider userName as parentusername | [optional] 
**Description** | Pointer to **string** |  | [optional] 

## Methods

### NewEditDelegateRequest

`func NewEditDelegateRequest(key string, userName string, name string, delegateusername string, delegatestartdate string, delegateenddate string, ) *EditDelegateRequest`

NewEditDelegateRequest instantiates a new EditDelegateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEditDelegateRequestWithDefaults

`func NewEditDelegateRequestWithDefaults() *EditDelegateRequest`

NewEditDelegateRequestWithDefaults instantiates a new EditDelegateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *EditDelegateRequest) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *EditDelegateRequest) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *EditDelegateRequest) SetKey(v string)`

SetKey sets Key field to given value.


### GetUserName

`func (o *EditDelegateRequest) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *EditDelegateRequest) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *EditDelegateRequest) SetUserName(v string)`

SetUserName sets UserName field to given value.


### GetName

`func (o *EditDelegateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EditDelegateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EditDelegateRequest) SetName(v string)`

SetName sets Name field to given value.


### GetDelegateusername

`func (o *EditDelegateRequest) GetDelegateusername() string`

GetDelegateusername returns the Delegateusername field if non-nil, zero value otherwise.

### GetDelegateusernameOk

`func (o *EditDelegateRequest) GetDelegateusernameOk() (*string, bool)`

GetDelegateusernameOk returns a tuple with the Delegateusername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelegateusername

`func (o *EditDelegateRequest) SetDelegateusername(v string)`

SetDelegateusername sets Delegateusername field to given value.


### GetDelegatestartdate

`func (o *EditDelegateRequest) GetDelegatestartdate() string`

GetDelegatestartdate returns the Delegatestartdate field if non-nil, zero value otherwise.

### GetDelegatestartdateOk

`func (o *EditDelegateRequest) GetDelegatestartdateOk() (*string, bool)`

GetDelegatestartdateOk returns a tuple with the Delegatestartdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelegatestartdate

`func (o *EditDelegateRequest) SetDelegatestartdate(v string)`

SetDelegatestartdate sets Delegatestartdate field to given value.


### GetDelegateenddate

`func (o *EditDelegateRequest) GetDelegateenddate() string`

GetDelegateenddate returns the Delegateenddate field if non-nil, zero value otherwise.

### GetDelegateenddateOk

`func (o *EditDelegateRequest) GetDelegateenddateOk() (*string, bool)`

GetDelegateenddateOk returns a tuple with the Delegateenddate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelegateenddate

`func (o *EditDelegateRequest) SetDelegateenddate(v string)`

SetDelegateenddate sets Delegateenddate field to given value.


### GetParentusername

`func (o *EditDelegateRequest) GetParentusername() string`

GetParentusername returns the Parentusername field if non-nil, zero value otherwise.

### GetParentusernameOk

`func (o *EditDelegateRequest) GetParentusernameOk() (*string, bool)`

GetParentusernameOk returns a tuple with the Parentusername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentusername

`func (o *EditDelegateRequest) SetParentusername(v string)`

SetParentusername sets Parentusername field to given value.

### HasParentusername

`func (o *EditDelegateRequest) HasParentusername() bool`

HasParentusername returns a boolean if a field has been set.

### GetDescription

`func (o *EditDelegateRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *EditDelegateRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *EditDelegateRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *EditDelegateRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


