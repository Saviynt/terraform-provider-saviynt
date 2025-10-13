# GetUserRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Username** | Pointer to **string** | Specify the username for which you want to get the user attribute details | [optional] 
**Responsefields** | Pointer to **[]interface{}** | User attributes which you want to see in the response(for encrypted values, mention ecp&lt;1-5&gt;, and for hashed values, mention hcp&lt;1-5&gt;) | [optional] 
**Max** | Pointer to **int32** |  | [optional] 
**Offset** | Pointer to **int32** |  | [optional] 
**Order** | Pointer to **string** |  | [optional] 
**Manager** | Pointer to **string** | username | [optional] 
**Secondarymanager** | Pointer to **string** | userkey OR &#x60;secondaryManager&#x60; - username | [optional] 
**Showsecurityanswers** | Pointer to **string** | \&quot;0\&quot;/\&quot;1\&quot; to display encrypted security answers for the user | [optional] 
**Filtercriteria** | Pointer to **map[string]interface{}** |  | [optional] 
**SearchCriteria** | Pointer to **string** |  | [optional] 
**Advsearchcriteria** | Pointer to **string** |  | [optional] 
**UserQuery** | Pointer to **string** |  | [optional] 

## Methods

### NewGetUserRequest

`func NewGetUserRequest() *GetUserRequest`

NewGetUserRequest instantiates a new GetUserRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetUserRequestWithDefaults

`func NewGetUserRequestWithDefaults() *GetUserRequest`

NewGetUserRequestWithDefaults instantiates a new GetUserRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsername

`func (o *GetUserRequest) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *GetUserRequest) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *GetUserRequest) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *GetUserRequest) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetResponsefields

`func (o *GetUserRequest) GetResponsefields() []interface{}`

GetResponsefields returns the Responsefields field if non-nil, zero value otherwise.

### GetResponsefieldsOk

`func (o *GetUserRequest) GetResponsefieldsOk() (*[]interface{}, bool)`

GetResponsefieldsOk returns a tuple with the Responsefields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponsefields

`func (o *GetUserRequest) SetResponsefields(v []interface{})`

SetResponsefields sets Responsefields field to given value.

### HasResponsefields

`func (o *GetUserRequest) HasResponsefields() bool`

HasResponsefields returns a boolean if a field has been set.

### GetMax

`func (o *GetUserRequest) GetMax() int32`

GetMax returns the Max field if non-nil, zero value otherwise.

### GetMaxOk

`func (o *GetUserRequest) GetMaxOk() (*int32, bool)`

GetMaxOk returns a tuple with the Max field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMax

`func (o *GetUserRequest) SetMax(v int32)`

SetMax sets Max field to given value.

### HasMax

`func (o *GetUserRequest) HasMax() bool`

HasMax returns a boolean if a field has been set.

### GetOffset

`func (o *GetUserRequest) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *GetUserRequest) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *GetUserRequest) SetOffset(v int32)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *GetUserRequest) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### GetOrder

`func (o *GetUserRequest) GetOrder() string`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *GetUserRequest) GetOrderOk() (*string, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *GetUserRequest) SetOrder(v string)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *GetUserRequest) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetManager

`func (o *GetUserRequest) GetManager() string`

GetManager returns the Manager field if non-nil, zero value otherwise.

### GetManagerOk

`func (o *GetUserRequest) GetManagerOk() (*string, bool)`

GetManagerOk returns a tuple with the Manager field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManager

`func (o *GetUserRequest) SetManager(v string)`

SetManager sets Manager field to given value.

### HasManager

`func (o *GetUserRequest) HasManager() bool`

HasManager returns a boolean if a field has been set.

### GetSecondarymanager

`func (o *GetUserRequest) GetSecondarymanager() string`

GetSecondarymanager returns the Secondarymanager field if non-nil, zero value otherwise.

### GetSecondarymanagerOk

`func (o *GetUserRequest) GetSecondarymanagerOk() (*string, bool)`

GetSecondarymanagerOk returns a tuple with the Secondarymanager field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondarymanager

`func (o *GetUserRequest) SetSecondarymanager(v string)`

SetSecondarymanager sets Secondarymanager field to given value.

### HasSecondarymanager

`func (o *GetUserRequest) HasSecondarymanager() bool`

HasSecondarymanager returns a boolean if a field has been set.

### GetShowsecurityanswers

`func (o *GetUserRequest) GetShowsecurityanswers() string`

GetShowsecurityanswers returns the Showsecurityanswers field if non-nil, zero value otherwise.

### GetShowsecurityanswersOk

`func (o *GetUserRequest) GetShowsecurityanswersOk() (*string, bool)`

GetShowsecurityanswersOk returns a tuple with the Showsecurityanswers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowsecurityanswers

`func (o *GetUserRequest) SetShowsecurityanswers(v string)`

SetShowsecurityanswers sets Showsecurityanswers field to given value.

### HasShowsecurityanswers

`func (o *GetUserRequest) HasShowsecurityanswers() bool`

HasShowsecurityanswers returns a boolean if a field has been set.

### GetFiltercriteria

`func (o *GetUserRequest) GetFiltercriteria() map[string]interface{}`

GetFiltercriteria returns the Filtercriteria field if non-nil, zero value otherwise.

### GetFiltercriteriaOk

`func (o *GetUserRequest) GetFiltercriteriaOk() (*map[string]interface{}, bool)`

GetFiltercriteriaOk returns a tuple with the Filtercriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiltercriteria

`func (o *GetUserRequest) SetFiltercriteria(v map[string]interface{})`

SetFiltercriteria sets Filtercriteria field to given value.

### HasFiltercriteria

`func (o *GetUserRequest) HasFiltercriteria() bool`

HasFiltercriteria returns a boolean if a field has been set.

### GetSearchCriteria

`func (o *GetUserRequest) GetSearchCriteria() string`

GetSearchCriteria returns the SearchCriteria field if non-nil, zero value otherwise.

### GetSearchCriteriaOk

`func (o *GetUserRequest) GetSearchCriteriaOk() (*string, bool)`

GetSearchCriteriaOk returns a tuple with the SearchCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchCriteria

`func (o *GetUserRequest) SetSearchCriteria(v string)`

SetSearchCriteria sets SearchCriteria field to given value.

### HasSearchCriteria

`func (o *GetUserRequest) HasSearchCriteria() bool`

HasSearchCriteria returns a boolean if a field has been set.

### GetAdvsearchcriteria

`func (o *GetUserRequest) GetAdvsearchcriteria() string`

GetAdvsearchcriteria returns the Advsearchcriteria field if non-nil, zero value otherwise.

### GetAdvsearchcriteriaOk

`func (o *GetUserRequest) GetAdvsearchcriteriaOk() (*string, bool)`

GetAdvsearchcriteriaOk returns a tuple with the Advsearchcriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdvsearchcriteria

`func (o *GetUserRequest) SetAdvsearchcriteria(v string)`

SetAdvsearchcriteria sets Advsearchcriteria field to given value.

### HasAdvsearchcriteria

`func (o *GetUserRequest) HasAdvsearchcriteria() bool`

HasAdvsearchcriteria returns a boolean if a field has been set.

### GetUserQuery

`func (o *GetUserRequest) GetUserQuery() string`

GetUserQuery returns the UserQuery field if non-nil, zero value otherwise.

### GetUserQueryOk

`func (o *GetUserRequest) GetUserQueryOk() (*string, bool)`

GetUserQueryOk returns a tuple with the UserQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserQuery

`func (o *GetUserRequest) SetUserQuery(v string)`

SetUserQuery sets UserQuery field to given value.

### HasUserQuery

`func (o *GetUserRequest) HasUserQuery() bool`

HasUserQuery returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


