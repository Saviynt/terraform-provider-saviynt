# DeleteDynamicAttributeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Securitysystem** | **string** | Name of the security system | 
**Endpoint** | **string** | Name of the endpoint | 
**Updateuser** | **string** | Username of the user performing the update | 
**Dynamicattributes** | **[]string** | List of dynamic attribute names to be deleted | 

## Methods

### NewDeleteDynamicAttributeRequest

`func NewDeleteDynamicAttributeRequest(securitysystem string, endpoint string, updateuser string, dynamicattributes []string, ) *DeleteDynamicAttributeRequest`

NewDeleteDynamicAttributeRequest instantiates a new DeleteDynamicAttributeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeleteDynamicAttributeRequestWithDefaults

`func NewDeleteDynamicAttributeRequestWithDefaults() *DeleteDynamicAttributeRequest`

NewDeleteDynamicAttributeRequestWithDefaults instantiates a new DeleteDynamicAttributeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSecuritysystem

`func (o *DeleteDynamicAttributeRequest) GetSecuritysystem() string`

GetSecuritysystem returns the Securitysystem field if non-nil, zero value otherwise.

### GetSecuritysystemOk

`func (o *DeleteDynamicAttributeRequest) GetSecuritysystemOk() (*string, bool)`

GetSecuritysystemOk returns a tuple with the Securitysystem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecuritysystem

`func (o *DeleteDynamicAttributeRequest) SetSecuritysystem(v string)`

SetSecuritysystem sets Securitysystem field to given value.


### GetEndpoint

`func (o *DeleteDynamicAttributeRequest) GetEndpoint() string`

GetEndpoint returns the Endpoint field if non-nil, zero value otherwise.

### GetEndpointOk

`func (o *DeleteDynamicAttributeRequest) GetEndpointOk() (*string, bool)`

GetEndpointOk returns a tuple with the Endpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoint

`func (o *DeleteDynamicAttributeRequest) SetEndpoint(v string)`

SetEndpoint sets Endpoint field to given value.


### GetUpdateuser

`func (o *DeleteDynamicAttributeRequest) GetUpdateuser() string`

GetUpdateuser returns the Updateuser field if non-nil, zero value otherwise.

### GetUpdateuserOk

`func (o *DeleteDynamicAttributeRequest) GetUpdateuserOk() (*string, bool)`

GetUpdateuserOk returns a tuple with the Updateuser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateuser

`func (o *DeleteDynamicAttributeRequest) SetUpdateuser(v string)`

SetUpdateuser sets Updateuser field to given value.


### GetDynamicattributes

`func (o *DeleteDynamicAttributeRequest) GetDynamicattributes() []string`

GetDynamicattributes returns the Dynamicattributes field if non-nil, zero value otherwise.

### GetDynamicattributesOk

`func (o *DeleteDynamicAttributeRequest) GetDynamicattributesOk() (*[]string, bool)`

GetDynamicattributesOk returns a tuple with the Dynamicattributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDynamicattributes

`func (o *DeleteDynamicAttributeRequest) SetDynamicattributes(v []string)`

SetDynamicattributes sets Dynamicattributes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


