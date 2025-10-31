# CreateOrUpdateEntitlementResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Msg** | Pointer to **string** | Response message indicating the result of the request. | [optional] 
**ErrorCode** | Pointer to **string** | Error code (0 for success). | [optional] 
**EntitlementObj** | Pointer to [**CreateOrUpdateEntitlementResponseEntitlementObj**](CreateOrUpdateEntitlementResponseEntitlementObj.md) |  | [optional] 

## Methods

### NewCreateOrUpdateEntitlementResponse

`func NewCreateOrUpdateEntitlementResponse() *CreateOrUpdateEntitlementResponse`

NewCreateOrUpdateEntitlementResponse instantiates a new CreateOrUpdateEntitlementResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateOrUpdateEntitlementResponseWithDefaults

`func NewCreateOrUpdateEntitlementResponseWithDefaults() *CreateOrUpdateEntitlementResponse`

NewCreateOrUpdateEntitlementResponseWithDefaults instantiates a new CreateOrUpdateEntitlementResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMsg

`func (o *CreateOrUpdateEntitlementResponse) GetMsg() string`

GetMsg returns the Msg field if non-nil, zero value otherwise.

### GetMsgOk

`func (o *CreateOrUpdateEntitlementResponse) GetMsgOk() (*string, bool)`

GetMsgOk returns a tuple with the Msg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsg

`func (o *CreateOrUpdateEntitlementResponse) SetMsg(v string)`

SetMsg sets Msg field to given value.

### HasMsg

`func (o *CreateOrUpdateEntitlementResponse) HasMsg() bool`

HasMsg returns a boolean if a field has been set.

### GetErrorCode

`func (o *CreateOrUpdateEntitlementResponse) GetErrorCode() string`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *CreateOrUpdateEntitlementResponse) GetErrorCodeOk() (*string, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *CreateOrUpdateEntitlementResponse) SetErrorCode(v string)`

SetErrorCode sets ErrorCode field to given value.

### HasErrorCode

`func (o *CreateOrUpdateEntitlementResponse) HasErrorCode() bool`

HasErrorCode returns a boolean if a field has been set.

### GetEntitlementObj

`func (o *CreateOrUpdateEntitlementResponse) GetEntitlementObj() CreateOrUpdateEntitlementResponseEntitlementObj`

GetEntitlementObj returns the EntitlementObj field if non-nil, zero value otherwise.

### GetEntitlementObjOk

`func (o *CreateOrUpdateEntitlementResponse) GetEntitlementObjOk() (*CreateOrUpdateEntitlementResponseEntitlementObj, bool)`

GetEntitlementObjOk returns a tuple with the EntitlementObj field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitlementObj

`func (o *CreateOrUpdateEntitlementResponse) SetEntitlementObj(v CreateOrUpdateEntitlementResponseEntitlementObj)`

SetEntitlementObj sets EntitlementObj field to given value.

### HasEntitlementObj

`func (o *CreateOrUpdateEntitlementResponse) HasEntitlementObj() bool`

HasEntitlementObj returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


