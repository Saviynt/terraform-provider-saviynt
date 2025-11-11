# GetEndpointsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Endpointname** | Pointer to **string** | Specify a name for the endpoint. Provide a logical name that will help you easily identify it. | [optional] 
**DisplayName** | Pointer to **string** |  Enter a user-friendly display name for the endpoint that will be displayed in the user interface. Display Name can be different from Endpoint Name. | [optional] 
**ConnectionType** | Pointer to **string** | Specify the Security system for which you want to create an endpoint. | [optional] 
**Endpointkey** | Pointer to **[]string** | Endpoint key. Specify the key(s) as an array of strings. | [optional] 
**Max** | Pointer to **string** | Description for the endpoint. | [optional] 
**Owner** | Pointer to **string** | Owner of the endpoint. If ownerType is User, specify the username of the owner. If ownerType is Usergroup, sepecify the name of the User group | [optional] 
**FilterCriteria** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewGetEndpointsRequest

`func NewGetEndpointsRequest() *GetEndpointsRequest`

NewGetEndpointsRequest instantiates a new GetEndpointsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetEndpointsRequestWithDefaults

`func NewGetEndpointsRequestWithDefaults() *GetEndpointsRequest`

NewGetEndpointsRequestWithDefaults instantiates a new GetEndpointsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEndpointname

`func (o *GetEndpointsRequest) GetEndpointname() string`

GetEndpointname returns the Endpointname field if non-nil, zero value otherwise.

### GetEndpointnameOk

`func (o *GetEndpointsRequest) GetEndpointnameOk() (*string, bool)`

GetEndpointnameOk returns a tuple with the Endpointname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpointname

`func (o *GetEndpointsRequest) SetEndpointname(v string)`

SetEndpointname sets Endpointname field to given value.

### HasEndpointname

`func (o *GetEndpointsRequest) HasEndpointname() bool`

HasEndpointname returns a boolean if a field has been set.

### GetDisplayName

`func (o *GetEndpointsRequest) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *GetEndpointsRequest) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *GetEndpointsRequest) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *GetEndpointsRequest) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetConnectionType

`func (o *GetEndpointsRequest) GetConnectionType() string`

GetConnectionType returns the ConnectionType field if non-nil, zero value otherwise.

### GetConnectionTypeOk

`func (o *GetEndpointsRequest) GetConnectionTypeOk() (*string, bool)`

GetConnectionTypeOk returns a tuple with the ConnectionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionType

`func (o *GetEndpointsRequest) SetConnectionType(v string)`

SetConnectionType sets ConnectionType field to given value.

### HasConnectionType

`func (o *GetEndpointsRequest) HasConnectionType() bool`

HasConnectionType returns a boolean if a field has been set.

### GetEndpointkey

`func (o *GetEndpointsRequest) GetEndpointkey() []string`

GetEndpointkey returns the Endpointkey field if non-nil, zero value otherwise.

### GetEndpointkeyOk

`func (o *GetEndpointsRequest) GetEndpointkeyOk() (*[]string, bool)`

GetEndpointkeyOk returns a tuple with the Endpointkey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpointkey

`func (o *GetEndpointsRequest) SetEndpointkey(v []string)`

SetEndpointkey sets Endpointkey field to given value.

### HasEndpointkey

`func (o *GetEndpointsRequest) HasEndpointkey() bool`

HasEndpointkey returns a boolean if a field has been set.

### GetMax

`func (o *GetEndpointsRequest) GetMax() string`

GetMax returns the Max field if non-nil, zero value otherwise.

### GetMaxOk

`func (o *GetEndpointsRequest) GetMaxOk() (*string, bool)`

GetMaxOk returns a tuple with the Max field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMax

`func (o *GetEndpointsRequest) SetMax(v string)`

SetMax sets Max field to given value.

### HasMax

`func (o *GetEndpointsRequest) HasMax() bool`

HasMax returns a boolean if a field has been set.

### GetOwner

`func (o *GetEndpointsRequest) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *GetEndpointsRequest) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *GetEndpointsRequest) SetOwner(v string)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *GetEndpointsRequest) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetFilterCriteria

`func (o *GetEndpointsRequest) GetFilterCriteria() map[string]interface{}`

GetFilterCriteria returns the FilterCriteria field if non-nil, zero value otherwise.

### GetFilterCriteriaOk

`func (o *GetEndpointsRequest) GetFilterCriteriaOk() (*map[string]interface{}, bool)`

GetFilterCriteriaOk returns a tuple with the FilterCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterCriteria

`func (o *GetEndpointsRequest) SetFilterCriteria(v map[string]interface{})`

SetFilterCriteria sets FilterCriteria field to given value.

### HasFilterCriteria

`func (o *GetEndpointsRequest) HasFilterCriteria() bool`

HasFilterCriteria returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


