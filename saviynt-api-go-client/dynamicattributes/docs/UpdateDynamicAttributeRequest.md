# UpdateDynamicAttributeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Securitysystem** | **string** | Name of the security systems | 
**Endpoint** | **string** | Name of the endpoint | 
**Updateuser** | **string** | Username | 
**Dynamicattributes** | [**[]UpdateDynamicAttributesInner**](UpdateDynamicAttributesInner.md) |  | 

## Methods

### NewUpdateDynamicAttributeRequest

`func NewUpdateDynamicAttributeRequest(securitysystem string, endpoint string, updateuser string, dynamicattributes []UpdateDynamicAttributesInner, ) *UpdateDynamicAttributeRequest`

NewUpdateDynamicAttributeRequest instantiates a new UpdateDynamicAttributeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateDynamicAttributeRequestWithDefaults

`func NewUpdateDynamicAttributeRequestWithDefaults() *UpdateDynamicAttributeRequest`

NewUpdateDynamicAttributeRequestWithDefaults instantiates a new UpdateDynamicAttributeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSecuritysystem

`func (o *UpdateDynamicAttributeRequest) GetSecuritysystem() string`

GetSecuritysystem returns the Securitysystem field if non-nil, zero value otherwise.

### GetSecuritysystemOk

`func (o *UpdateDynamicAttributeRequest) GetSecuritysystemOk() (*string, bool)`

GetSecuritysystemOk returns a tuple with the Securitysystem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecuritysystem

`func (o *UpdateDynamicAttributeRequest) SetSecuritysystem(v string)`

SetSecuritysystem sets Securitysystem field to given value.


### GetEndpoint

`func (o *UpdateDynamicAttributeRequest) GetEndpoint() string`

GetEndpoint returns the Endpoint field if non-nil, zero value otherwise.

### GetEndpointOk

`func (o *UpdateDynamicAttributeRequest) GetEndpointOk() (*string, bool)`

GetEndpointOk returns a tuple with the Endpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoint

`func (o *UpdateDynamicAttributeRequest) SetEndpoint(v string)`

SetEndpoint sets Endpoint field to given value.


### GetUpdateuser

`func (o *UpdateDynamicAttributeRequest) GetUpdateuser() string`

GetUpdateuser returns the Updateuser field if non-nil, zero value otherwise.

### GetUpdateuserOk

`func (o *UpdateDynamicAttributeRequest) GetUpdateuserOk() (*string, bool)`

GetUpdateuserOk returns a tuple with the Updateuser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateuser

`func (o *UpdateDynamicAttributeRequest) SetUpdateuser(v string)`

SetUpdateuser sets Updateuser field to given value.


### GetDynamicattributes

`func (o *UpdateDynamicAttributeRequest) GetDynamicattributes() []UpdateDynamicAttributesInner`

GetDynamicattributes returns the Dynamicattributes field if non-nil, zero value otherwise.

### GetDynamicattributesOk

`func (o *UpdateDynamicAttributeRequest) GetDynamicattributesOk() (*[]UpdateDynamicAttributesInner, bool)`

GetDynamicattributesOk returns a tuple with the Dynamicattributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDynamicattributes

`func (o *UpdateDynamicAttributeRequest) SetDynamicattributes(v []UpdateDynamicAttributesInner)`

SetDynamicattributes sets Dynamicattributes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


