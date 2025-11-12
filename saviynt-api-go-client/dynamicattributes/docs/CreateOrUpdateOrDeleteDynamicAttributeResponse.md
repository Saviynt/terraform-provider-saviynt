# CreateOrUpdateOrDeleteDynamicAttributeResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Msg** | Pointer to **string** | A message indicating the outcome of the operation. | [optional] 
**Errorcode** | Pointer to **string** | An error code where &#39;0&#39; signifies success and &#39;1&#39; signifies an unsuccessful operation. | [optional] 
**Securitysystem** | Pointer to **string** | Name of the security system | [optional] 
**Endpoint** | Pointer to **string** | Name of endpoint | [optional] 
**Updateuser** | Pointer to **string** | Username of the user performing the update | [optional] 
**Dynamicattributes** | Pointer to [**CreateOrUpdateOrDeleteDynamicAttributeResponseDynamicattributes**](CreateOrUpdateOrDeleteDynamicAttributeResponseDynamicattributes.md) |  | [optional] 

## Methods

### NewCreateOrUpdateOrDeleteDynamicAttributeResponse

`func NewCreateOrUpdateOrDeleteDynamicAttributeResponse() *CreateOrUpdateOrDeleteDynamicAttributeResponse`

NewCreateOrUpdateOrDeleteDynamicAttributeResponse instantiates a new CreateOrUpdateOrDeleteDynamicAttributeResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateOrUpdateOrDeleteDynamicAttributeResponseWithDefaults

`func NewCreateOrUpdateOrDeleteDynamicAttributeResponseWithDefaults() *CreateOrUpdateOrDeleteDynamicAttributeResponse`

NewCreateOrUpdateOrDeleteDynamicAttributeResponseWithDefaults instantiates a new CreateOrUpdateOrDeleteDynamicAttributeResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMsg

`func (o *CreateOrUpdateOrDeleteDynamicAttributeResponse) GetMsg() string`

GetMsg returns the Msg field if non-nil, zero value otherwise.

### GetMsgOk

`func (o *CreateOrUpdateOrDeleteDynamicAttributeResponse) GetMsgOk() (*string, bool)`

GetMsgOk returns a tuple with the Msg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsg

`func (o *CreateOrUpdateOrDeleteDynamicAttributeResponse) SetMsg(v string)`

SetMsg sets Msg field to given value.

### HasMsg

`func (o *CreateOrUpdateOrDeleteDynamicAttributeResponse) HasMsg() bool`

HasMsg returns a boolean if a field has been set.

### GetErrorcode

`func (o *CreateOrUpdateOrDeleteDynamicAttributeResponse) GetErrorcode() string`

GetErrorcode returns the Errorcode field if non-nil, zero value otherwise.

### GetErrorcodeOk

`func (o *CreateOrUpdateOrDeleteDynamicAttributeResponse) GetErrorcodeOk() (*string, bool)`

GetErrorcodeOk returns a tuple with the Errorcode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorcode

`func (o *CreateOrUpdateOrDeleteDynamicAttributeResponse) SetErrorcode(v string)`

SetErrorcode sets Errorcode field to given value.

### HasErrorcode

`func (o *CreateOrUpdateOrDeleteDynamicAttributeResponse) HasErrorcode() bool`

HasErrorcode returns a boolean if a field has been set.

### GetSecuritysystem

`func (o *CreateOrUpdateOrDeleteDynamicAttributeResponse) GetSecuritysystem() string`

GetSecuritysystem returns the Securitysystem field if non-nil, zero value otherwise.

### GetSecuritysystemOk

`func (o *CreateOrUpdateOrDeleteDynamicAttributeResponse) GetSecuritysystemOk() (*string, bool)`

GetSecuritysystemOk returns a tuple with the Securitysystem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecuritysystem

`func (o *CreateOrUpdateOrDeleteDynamicAttributeResponse) SetSecuritysystem(v string)`

SetSecuritysystem sets Securitysystem field to given value.

### HasSecuritysystem

`func (o *CreateOrUpdateOrDeleteDynamicAttributeResponse) HasSecuritysystem() bool`

HasSecuritysystem returns a boolean if a field has been set.

### GetEndpoint

`func (o *CreateOrUpdateOrDeleteDynamicAttributeResponse) GetEndpoint() string`

GetEndpoint returns the Endpoint field if non-nil, zero value otherwise.

### GetEndpointOk

`func (o *CreateOrUpdateOrDeleteDynamicAttributeResponse) GetEndpointOk() (*string, bool)`

GetEndpointOk returns a tuple with the Endpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoint

`func (o *CreateOrUpdateOrDeleteDynamicAttributeResponse) SetEndpoint(v string)`

SetEndpoint sets Endpoint field to given value.

### HasEndpoint

`func (o *CreateOrUpdateOrDeleteDynamicAttributeResponse) HasEndpoint() bool`

HasEndpoint returns a boolean if a field has been set.

### GetUpdateuser

`func (o *CreateOrUpdateOrDeleteDynamicAttributeResponse) GetUpdateuser() string`

GetUpdateuser returns the Updateuser field if non-nil, zero value otherwise.

### GetUpdateuserOk

`func (o *CreateOrUpdateOrDeleteDynamicAttributeResponse) GetUpdateuserOk() (*string, bool)`

GetUpdateuserOk returns a tuple with the Updateuser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateuser

`func (o *CreateOrUpdateOrDeleteDynamicAttributeResponse) SetUpdateuser(v string)`

SetUpdateuser sets Updateuser field to given value.

### HasUpdateuser

`func (o *CreateOrUpdateOrDeleteDynamicAttributeResponse) HasUpdateuser() bool`

HasUpdateuser returns a boolean if a field has been set.

### GetDynamicattributes

`func (o *CreateOrUpdateOrDeleteDynamicAttributeResponse) GetDynamicattributes() CreateOrUpdateOrDeleteDynamicAttributeResponseDynamicattributes`

GetDynamicattributes returns the Dynamicattributes field if non-nil, zero value otherwise.

### GetDynamicattributesOk

`func (o *CreateOrUpdateOrDeleteDynamicAttributeResponse) GetDynamicattributesOk() (*CreateOrUpdateOrDeleteDynamicAttributeResponseDynamicattributes, bool)`

GetDynamicattributesOk returns a tuple with the Dynamicattributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDynamicattributes

`func (o *CreateOrUpdateOrDeleteDynamicAttributeResponse) SetDynamicattributes(v CreateOrUpdateOrDeleteDynamicAttributeResponseDynamicattributes)`

SetDynamicattributes sets Dynamicattributes field to given value.

### HasDynamicattributes

`func (o *CreateOrUpdateOrDeleteDynamicAttributeResponse) HasDynamicattributes() bool`

HasDynamicattributes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


