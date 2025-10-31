# CreateDynamicAttributeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Securitysystem** | **string** | Name of the security systems | 
**Endpoint** | **string** | Name of the endpoint | 
**Updateuser** | **string** | Username | 
**Dynamicattributes** | [**[]CreateDynamicAttributesInner**](CreateDynamicAttributesInner.md) |  | 

## Methods

### NewCreateDynamicAttributeRequest

`func NewCreateDynamicAttributeRequest(securitysystem string, endpoint string, updateuser string, dynamicattributes []CreateDynamicAttributesInner, ) *CreateDynamicAttributeRequest`

NewCreateDynamicAttributeRequest instantiates a new CreateDynamicAttributeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateDynamicAttributeRequestWithDefaults

`func NewCreateDynamicAttributeRequestWithDefaults() *CreateDynamicAttributeRequest`

NewCreateDynamicAttributeRequestWithDefaults instantiates a new CreateDynamicAttributeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSecuritysystem

`func (o *CreateDynamicAttributeRequest) GetSecuritysystem() string`

GetSecuritysystem returns the Securitysystem field if non-nil, zero value otherwise.

### GetSecuritysystemOk

`func (o *CreateDynamicAttributeRequest) GetSecuritysystemOk() (*string, bool)`

GetSecuritysystemOk returns a tuple with the Securitysystem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecuritysystem

`func (o *CreateDynamicAttributeRequest) SetSecuritysystem(v string)`

SetSecuritysystem sets Securitysystem field to given value.


### GetEndpoint

`func (o *CreateDynamicAttributeRequest) GetEndpoint() string`

GetEndpoint returns the Endpoint field if non-nil, zero value otherwise.

### GetEndpointOk

`func (o *CreateDynamicAttributeRequest) GetEndpointOk() (*string, bool)`

GetEndpointOk returns a tuple with the Endpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoint

`func (o *CreateDynamicAttributeRequest) SetEndpoint(v string)`

SetEndpoint sets Endpoint field to given value.


### GetUpdateuser

`func (o *CreateDynamicAttributeRequest) GetUpdateuser() string`

GetUpdateuser returns the Updateuser field if non-nil, zero value otherwise.

### GetUpdateuserOk

`func (o *CreateDynamicAttributeRequest) GetUpdateuserOk() (*string, bool)`

GetUpdateuserOk returns a tuple with the Updateuser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateuser

`func (o *CreateDynamicAttributeRequest) SetUpdateuser(v string)`

SetUpdateuser sets Updateuser field to given value.


### GetDynamicattributes

`func (o *CreateDynamicAttributeRequest) GetDynamicattributes() []CreateDynamicAttributesInner`

GetDynamicattributes returns the Dynamicattributes field if non-nil, zero value otherwise.

### GetDynamicattributesOk

`func (o *CreateDynamicAttributeRequest) GetDynamicattributesOk() (*[]CreateDynamicAttributesInner, bool)`

GetDynamicattributesOk returns a tuple with the Dynamicattributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDynamicattributes

`func (o *CreateDynamicAttributeRequest) SetDynamicattributes(v []CreateDynamicAttributesInner)`

SetDynamicattributes sets Dynamicattributes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


