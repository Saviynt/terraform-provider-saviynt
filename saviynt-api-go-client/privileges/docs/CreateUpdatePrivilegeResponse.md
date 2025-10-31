# CreateUpdatePrivilegeResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Msg** | Pointer to **string** | A message indicating the outcome of the operation. | [optional] 
**Errorcode** | Pointer to **int32** | An error code where &#39;0&#39; signifies success and &#39;1&#39; signifies an unsuccessful operation. | [optional] 
**Entitlementtypeprivilegekey** | Pointer to **int32** | Privilege key for the entitlement type | [optional] 
**Securitysystem** | Pointer to **string** | Name of the security system | [optional] 
**Endpoint** | Pointer to **string** | Name of endpoint | [optional] 
**Entitlementtype** | Pointer to **string** | Name of entitltment type | [optional] 
**Privilege** | Pointer to [**CreateUpdatePrivilegeResponsePrivilege**](CreateUpdatePrivilegeResponsePrivilege.md) |  | [optional] 

## Methods

### NewCreateUpdatePrivilegeResponse

`func NewCreateUpdatePrivilegeResponse() *CreateUpdatePrivilegeResponse`

NewCreateUpdatePrivilegeResponse instantiates a new CreateUpdatePrivilegeResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateUpdatePrivilegeResponseWithDefaults

`func NewCreateUpdatePrivilegeResponseWithDefaults() *CreateUpdatePrivilegeResponse`

NewCreateUpdatePrivilegeResponseWithDefaults instantiates a new CreateUpdatePrivilegeResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMsg

`func (o *CreateUpdatePrivilegeResponse) GetMsg() string`

GetMsg returns the Msg field if non-nil, zero value otherwise.

### GetMsgOk

`func (o *CreateUpdatePrivilegeResponse) GetMsgOk() (*string, bool)`

GetMsgOk returns a tuple with the Msg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsg

`func (o *CreateUpdatePrivilegeResponse) SetMsg(v string)`

SetMsg sets Msg field to given value.

### HasMsg

`func (o *CreateUpdatePrivilegeResponse) HasMsg() bool`

HasMsg returns a boolean if a field has been set.

### GetErrorcode

`func (o *CreateUpdatePrivilegeResponse) GetErrorcode() int32`

GetErrorcode returns the Errorcode field if non-nil, zero value otherwise.

### GetErrorcodeOk

`func (o *CreateUpdatePrivilegeResponse) GetErrorcodeOk() (*int32, bool)`

GetErrorcodeOk returns a tuple with the Errorcode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorcode

`func (o *CreateUpdatePrivilegeResponse) SetErrorcode(v int32)`

SetErrorcode sets Errorcode field to given value.

### HasErrorcode

`func (o *CreateUpdatePrivilegeResponse) HasErrorcode() bool`

HasErrorcode returns a boolean if a field has been set.

### GetEntitlementtypeprivilegekey

`func (o *CreateUpdatePrivilegeResponse) GetEntitlementtypeprivilegekey() int32`

GetEntitlementtypeprivilegekey returns the Entitlementtypeprivilegekey field if non-nil, zero value otherwise.

### GetEntitlementtypeprivilegekeyOk

`func (o *CreateUpdatePrivilegeResponse) GetEntitlementtypeprivilegekeyOk() (*int32, bool)`

GetEntitlementtypeprivilegekeyOk returns a tuple with the Entitlementtypeprivilegekey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitlementtypeprivilegekey

`func (o *CreateUpdatePrivilegeResponse) SetEntitlementtypeprivilegekey(v int32)`

SetEntitlementtypeprivilegekey sets Entitlementtypeprivilegekey field to given value.

### HasEntitlementtypeprivilegekey

`func (o *CreateUpdatePrivilegeResponse) HasEntitlementtypeprivilegekey() bool`

HasEntitlementtypeprivilegekey returns a boolean if a field has been set.

### GetSecuritysystem

`func (o *CreateUpdatePrivilegeResponse) GetSecuritysystem() string`

GetSecuritysystem returns the Securitysystem field if non-nil, zero value otherwise.

### GetSecuritysystemOk

`func (o *CreateUpdatePrivilegeResponse) GetSecuritysystemOk() (*string, bool)`

GetSecuritysystemOk returns a tuple with the Securitysystem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecuritysystem

`func (o *CreateUpdatePrivilegeResponse) SetSecuritysystem(v string)`

SetSecuritysystem sets Securitysystem field to given value.

### HasSecuritysystem

`func (o *CreateUpdatePrivilegeResponse) HasSecuritysystem() bool`

HasSecuritysystem returns a boolean if a field has been set.

### GetEndpoint

`func (o *CreateUpdatePrivilegeResponse) GetEndpoint() string`

GetEndpoint returns the Endpoint field if non-nil, zero value otherwise.

### GetEndpointOk

`func (o *CreateUpdatePrivilegeResponse) GetEndpointOk() (*string, bool)`

GetEndpointOk returns a tuple with the Endpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoint

`func (o *CreateUpdatePrivilegeResponse) SetEndpoint(v string)`

SetEndpoint sets Endpoint field to given value.

### HasEndpoint

`func (o *CreateUpdatePrivilegeResponse) HasEndpoint() bool`

HasEndpoint returns a boolean if a field has been set.

### GetEntitlementtype

`func (o *CreateUpdatePrivilegeResponse) GetEntitlementtype() string`

GetEntitlementtype returns the Entitlementtype field if non-nil, zero value otherwise.

### GetEntitlementtypeOk

`func (o *CreateUpdatePrivilegeResponse) GetEntitlementtypeOk() (*string, bool)`

GetEntitlementtypeOk returns a tuple with the Entitlementtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitlementtype

`func (o *CreateUpdatePrivilegeResponse) SetEntitlementtype(v string)`

SetEntitlementtype sets Entitlementtype field to given value.

### HasEntitlementtype

`func (o *CreateUpdatePrivilegeResponse) HasEntitlementtype() bool`

HasEntitlementtype returns a boolean if a field has been set.

### GetPrivilege

`func (o *CreateUpdatePrivilegeResponse) GetPrivilege() CreateUpdatePrivilegeResponsePrivilege`

GetPrivilege returns the Privilege field if non-nil, zero value otherwise.

### GetPrivilegeOk

`func (o *CreateUpdatePrivilegeResponse) GetPrivilegeOk() (*CreateUpdatePrivilegeResponsePrivilege, bool)`

GetPrivilegeOk returns a tuple with the Privilege field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivilege

`func (o *CreateUpdatePrivilegeResponse) SetPrivilege(v CreateUpdatePrivilegeResponsePrivilege)`

SetPrivilege sets Privilege field to given value.

### HasPrivilege

`func (o *CreateUpdatePrivilegeResponse) HasPrivilege() bool`

HasPrivilege returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


