# DeletePrivilegeResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Msg** | Pointer to **string** | Response message | [optional] 
**Errorcode** | Pointer to **int32** | Error code | [optional] 
**Securitysystem** | Pointer to **string** | Security system validation error | [optional] 
**Endpoint** | Pointer to **string** | Endpoint validation error | [optional] 
**Entitlementtype** | Pointer to **string** | Entitlement type validation error | [optional] 
**Privilege** | Pointer to [**DeletePrivilegeResponsePrivilege**](DeletePrivilegeResponsePrivilege.md) |  | [optional] 

## Methods

### NewDeletePrivilegeResponse

`func NewDeletePrivilegeResponse() *DeletePrivilegeResponse`

NewDeletePrivilegeResponse instantiates a new DeletePrivilegeResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeletePrivilegeResponseWithDefaults

`func NewDeletePrivilegeResponseWithDefaults() *DeletePrivilegeResponse`

NewDeletePrivilegeResponseWithDefaults instantiates a new DeletePrivilegeResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMsg

`func (o *DeletePrivilegeResponse) GetMsg() string`

GetMsg returns the Msg field if non-nil, zero value otherwise.

### GetMsgOk

`func (o *DeletePrivilegeResponse) GetMsgOk() (*string, bool)`

GetMsgOk returns a tuple with the Msg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsg

`func (o *DeletePrivilegeResponse) SetMsg(v string)`

SetMsg sets Msg field to given value.

### HasMsg

`func (o *DeletePrivilegeResponse) HasMsg() bool`

HasMsg returns a boolean if a field has been set.

### GetErrorcode

`func (o *DeletePrivilegeResponse) GetErrorcode() int32`

GetErrorcode returns the Errorcode field if non-nil, zero value otherwise.

### GetErrorcodeOk

`func (o *DeletePrivilegeResponse) GetErrorcodeOk() (*int32, bool)`

GetErrorcodeOk returns a tuple with the Errorcode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorcode

`func (o *DeletePrivilegeResponse) SetErrorcode(v int32)`

SetErrorcode sets Errorcode field to given value.

### HasErrorcode

`func (o *DeletePrivilegeResponse) HasErrorcode() bool`

HasErrorcode returns a boolean if a field has been set.

### GetSecuritysystem

`func (o *DeletePrivilegeResponse) GetSecuritysystem() string`

GetSecuritysystem returns the Securitysystem field if non-nil, zero value otherwise.

### GetSecuritysystemOk

`func (o *DeletePrivilegeResponse) GetSecuritysystemOk() (*string, bool)`

GetSecuritysystemOk returns a tuple with the Securitysystem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecuritysystem

`func (o *DeletePrivilegeResponse) SetSecuritysystem(v string)`

SetSecuritysystem sets Securitysystem field to given value.

### HasSecuritysystem

`func (o *DeletePrivilegeResponse) HasSecuritysystem() bool`

HasSecuritysystem returns a boolean if a field has been set.

### GetEndpoint

`func (o *DeletePrivilegeResponse) GetEndpoint() string`

GetEndpoint returns the Endpoint field if non-nil, zero value otherwise.

### GetEndpointOk

`func (o *DeletePrivilegeResponse) GetEndpointOk() (*string, bool)`

GetEndpointOk returns a tuple with the Endpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoint

`func (o *DeletePrivilegeResponse) SetEndpoint(v string)`

SetEndpoint sets Endpoint field to given value.

### HasEndpoint

`func (o *DeletePrivilegeResponse) HasEndpoint() bool`

HasEndpoint returns a boolean if a field has been set.

### GetEntitlementtype

`func (o *DeletePrivilegeResponse) GetEntitlementtype() string`

GetEntitlementtype returns the Entitlementtype field if non-nil, zero value otherwise.

### GetEntitlementtypeOk

`func (o *DeletePrivilegeResponse) GetEntitlementtypeOk() (*string, bool)`

GetEntitlementtypeOk returns a tuple with the Entitlementtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitlementtype

`func (o *DeletePrivilegeResponse) SetEntitlementtype(v string)`

SetEntitlementtype sets Entitlementtype field to given value.

### HasEntitlementtype

`func (o *DeletePrivilegeResponse) HasEntitlementtype() bool`

HasEntitlementtype returns a boolean if a field has been set.

### GetPrivilege

`func (o *DeletePrivilegeResponse) GetPrivilege() DeletePrivilegeResponsePrivilege`

GetPrivilege returns the Privilege field if non-nil, zero value otherwise.

### GetPrivilegeOk

`func (o *DeletePrivilegeResponse) GetPrivilegeOk() (*DeletePrivilegeResponsePrivilege, bool)`

GetPrivilegeOk returns a tuple with the Privilege field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivilege

`func (o *DeletePrivilegeResponse) SetPrivilege(v DeletePrivilegeResponsePrivilege)`

SetPrivilege sets Privilege field to given value.

### HasPrivilege

`func (o *DeletePrivilegeResponse) HasPrivilege() bool`

HasPrivilege returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


