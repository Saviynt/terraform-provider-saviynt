# DeletePrivilegeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Securitysystem** | **string** | Security system name | 
**Endpoint** | **string** | Endpoint Name | 
**Entitlementtype** | **string** | Entitlement type | 
**Privilege** | **string** | Name of the privilege(attribute name) to be deleted | 

## Methods

### NewDeletePrivilegeRequest

`func NewDeletePrivilegeRequest(securitysystem string, endpoint string, entitlementtype string, privilege string, ) *DeletePrivilegeRequest`

NewDeletePrivilegeRequest instantiates a new DeletePrivilegeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeletePrivilegeRequestWithDefaults

`func NewDeletePrivilegeRequestWithDefaults() *DeletePrivilegeRequest`

NewDeletePrivilegeRequestWithDefaults instantiates a new DeletePrivilegeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSecuritysystem

`func (o *DeletePrivilegeRequest) GetSecuritysystem() string`

GetSecuritysystem returns the Securitysystem field if non-nil, zero value otherwise.

### GetSecuritysystemOk

`func (o *DeletePrivilegeRequest) GetSecuritysystemOk() (*string, bool)`

GetSecuritysystemOk returns a tuple with the Securitysystem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecuritysystem

`func (o *DeletePrivilegeRequest) SetSecuritysystem(v string)`

SetSecuritysystem sets Securitysystem field to given value.


### GetEndpoint

`func (o *DeletePrivilegeRequest) GetEndpoint() string`

GetEndpoint returns the Endpoint field if non-nil, zero value otherwise.

### GetEndpointOk

`func (o *DeletePrivilegeRequest) GetEndpointOk() (*string, bool)`

GetEndpointOk returns a tuple with the Endpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoint

`func (o *DeletePrivilegeRequest) SetEndpoint(v string)`

SetEndpoint sets Endpoint field to given value.


### GetEntitlementtype

`func (o *DeletePrivilegeRequest) GetEntitlementtype() string`

GetEntitlementtype returns the Entitlementtype field if non-nil, zero value otherwise.

### GetEntitlementtypeOk

`func (o *DeletePrivilegeRequest) GetEntitlementtypeOk() (*string, bool)`

GetEntitlementtypeOk returns a tuple with the Entitlementtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitlementtype

`func (o *DeletePrivilegeRequest) SetEntitlementtype(v string)`

SetEntitlementtype sets Entitlementtype field to given value.


### GetPrivilege

`func (o *DeletePrivilegeRequest) GetPrivilege() string`

GetPrivilege returns the Privilege field if non-nil, zero value otherwise.

### GetPrivilegeOk

`func (o *DeletePrivilegeRequest) GetPrivilegeOk() (*string, bool)`

GetPrivilegeOk returns a tuple with the Privilege field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivilege

`func (o *DeletePrivilegeRequest) SetPrivilege(v string)`

SetPrivilege sets Privilege field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


