# GetPrivilegeListRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Endpoint** | **string** | Endpoint name to search privilege | 
**Entitlementtype** | Pointer to **string** | Entitlement type | [optional] 
**Max** | Pointer to **string** | Max number of results to return | [optional] 
**Offset** | Pointer to **string** | Offset for the returned results | [optional] 

## Methods

### NewGetPrivilegeListRequest

`func NewGetPrivilegeListRequest(endpoint string, ) *GetPrivilegeListRequest`

NewGetPrivilegeListRequest instantiates a new GetPrivilegeListRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetPrivilegeListRequestWithDefaults

`func NewGetPrivilegeListRequestWithDefaults() *GetPrivilegeListRequest`

NewGetPrivilegeListRequestWithDefaults instantiates a new GetPrivilegeListRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEndpoint

`func (o *GetPrivilegeListRequest) GetEndpoint() string`

GetEndpoint returns the Endpoint field if non-nil, zero value otherwise.

### GetEndpointOk

`func (o *GetPrivilegeListRequest) GetEndpointOk() (*string, bool)`

GetEndpointOk returns a tuple with the Endpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoint

`func (o *GetPrivilegeListRequest) SetEndpoint(v string)`

SetEndpoint sets Endpoint field to given value.


### GetEntitlementtype

`func (o *GetPrivilegeListRequest) GetEntitlementtype() string`

GetEntitlementtype returns the Entitlementtype field if non-nil, zero value otherwise.

### GetEntitlementtypeOk

`func (o *GetPrivilegeListRequest) GetEntitlementtypeOk() (*string, bool)`

GetEntitlementtypeOk returns a tuple with the Entitlementtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitlementtype

`func (o *GetPrivilegeListRequest) SetEntitlementtype(v string)`

SetEntitlementtype sets Entitlementtype field to given value.

### HasEntitlementtype

`func (o *GetPrivilegeListRequest) HasEntitlementtype() bool`

HasEntitlementtype returns a boolean if a field has been set.

### GetMax

`func (o *GetPrivilegeListRequest) GetMax() string`

GetMax returns the Max field if non-nil, zero value otherwise.

### GetMaxOk

`func (o *GetPrivilegeListRequest) GetMaxOk() (*string, bool)`

GetMaxOk returns a tuple with the Max field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMax

`func (o *GetPrivilegeListRequest) SetMax(v string)`

SetMax sets Max field to given value.

### HasMax

`func (o *GetPrivilegeListRequest) HasMax() bool`

HasMax returns a boolean if a field has been set.

### GetOffset

`func (o *GetPrivilegeListRequest) GetOffset() string`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *GetPrivilegeListRequest) GetOffsetOk() (*string, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *GetPrivilegeListRequest) SetOffset(v string)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *GetPrivilegeListRequest) HasOffset() bool`

HasOffset returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


