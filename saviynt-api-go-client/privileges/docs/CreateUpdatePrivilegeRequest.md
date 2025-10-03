# CreateUpdatePrivilegeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Securitysystem** | **string** | Name of the security system to which the endpoint belongs | 
**Endpoint** | **string** | Name of the endpoint to which the entitlement type belongs | 
**Entitlementtype** | **string** | Name of the entitlement type for the privilege | 
**Privilege** | [**CreateUpdatePrivilegeRequestPrivilege**](CreateUpdatePrivilegeRequestPrivilege.md) |  | 

## Methods

### NewCreateUpdatePrivilegeRequest

`func NewCreateUpdatePrivilegeRequest(securitysystem string, endpoint string, entitlementtype string, privilege CreateUpdatePrivilegeRequestPrivilege, ) *CreateUpdatePrivilegeRequest`

NewCreateUpdatePrivilegeRequest instantiates a new CreateUpdatePrivilegeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateUpdatePrivilegeRequestWithDefaults

`func NewCreateUpdatePrivilegeRequestWithDefaults() *CreateUpdatePrivilegeRequest`

NewCreateUpdatePrivilegeRequestWithDefaults instantiates a new CreateUpdatePrivilegeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSecuritysystem

`func (o *CreateUpdatePrivilegeRequest) GetSecuritysystem() string`

GetSecuritysystem returns the Securitysystem field if non-nil, zero value otherwise.

### GetSecuritysystemOk

`func (o *CreateUpdatePrivilegeRequest) GetSecuritysystemOk() (*string, bool)`

GetSecuritysystemOk returns a tuple with the Securitysystem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecuritysystem

`func (o *CreateUpdatePrivilegeRequest) SetSecuritysystem(v string)`

SetSecuritysystem sets Securitysystem field to given value.


### GetEndpoint

`func (o *CreateUpdatePrivilegeRequest) GetEndpoint() string`

GetEndpoint returns the Endpoint field if non-nil, zero value otherwise.

### GetEndpointOk

`func (o *CreateUpdatePrivilegeRequest) GetEndpointOk() (*string, bool)`

GetEndpointOk returns a tuple with the Endpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoint

`func (o *CreateUpdatePrivilegeRequest) SetEndpoint(v string)`

SetEndpoint sets Endpoint field to given value.


### GetEntitlementtype

`func (o *CreateUpdatePrivilegeRequest) GetEntitlementtype() string`

GetEntitlementtype returns the Entitlementtype field if non-nil, zero value otherwise.

### GetEntitlementtypeOk

`func (o *CreateUpdatePrivilegeRequest) GetEntitlementtypeOk() (*string, bool)`

GetEntitlementtypeOk returns a tuple with the Entitlementtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitlementtype

`func (o *CreateUpdatePrivilegeRequest) SetEntitlementtype(v string)`

SetEntitlementtype sets Entitlementtype field to given value.


### GetPrivilege

`func (o *CreateUpdatePrivilegeRequest) GetPrivilege() CreateUpdatePrivilegeRequestPrivilege`

GetPrivilege returns the Privilege field if non-nil, zero value otherwise.

### GetPrivilegeOk

`func (o *CreateUpdatePrivilegeRequest) GetPrivilegeOk() (*CreateUpdatePrivilegeRequestPrivilege, bool)`

GetPrivilegeOk returns a tuple with the Privilege field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivilege

`func (o *CreateUpdatePrivilegeRequest) SetPrivilege(v CreateUpdatePrivilegeRequestPrivilege)`

SetPrivilege sets Privilege field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


