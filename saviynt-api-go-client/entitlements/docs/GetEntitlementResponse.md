# GetEntitlementResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Msg** | Pointer to **string** | Response message indicating the result of the request. | [optional] 
**ErrorCode** | Pointer to **string** | Error code (0 for success). | [optional] 
**TotalEntitlementCount** | Pointer to **int32** | Total number of entitlements in the system. | [optional] 
**EntitlementsCount** | Pointer to **int32** | Number of entitlements returned in the response. | [optional] 
**Entitlementdetails** | Pointer to [**[]GetEntitlementResponseEntitlementdetailsInner**](GetEntitlementResponseEntitlementdetailsInner.md) | List of entitlements with full metadata. | [optional] 

## Methods

### NewGetEntitlementResponse

`func NewGetEntitlementResponse() *GetEntitlementResponse`

NewGetEntitlementResponse instantiates a new GetEntitlementResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetEntitlementResponseWithDefaults

`func NewGetEntitlementResponseWithDefaults() *GetEntitlementResponse`

NewGetEntitlementResponseWithDefaults instantiates a new GetEntitlementResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMsg

`func (o *GetEntitlementResponse) GetMsg() string`

GetMsg returns the Msg field if non-nil, zero value otherwise.

### GetMsgOk

`func (o *GetEntitlementResponse) GetMsgOk() (*string, bool)`

GetMsgOk returns a tuple with the Msg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsg

`func (o *GetEntitlementResponse) SetMsg(v string)`

SetMsg sets Msg field to given value.

### HasMsg

`func (o *GetEntitlementResponse) HasMsg() bool`

HasMsg returns a boolean if a field has been set.

### GetErrorCode

`func (o *GetEntitlementResponse) GetErrorCode() string`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *GetEntitlementResponse) GetErrorCodeOk() (*string, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *GetEntitlementResponse) SetErrorCode(v string)`

SetErrorCode sets ErrorCode field to given value.

### HasErrorCode

`func (o *GetEntitlementResponse) HasErrorCode() bool`

HasErrorCode returns a boolean if a field has been set.

### GetTotalEntitlementCount

`func (o *GetEntitlementResponse) GetTotalEntitlementCount() int32`

GetTotalEntitlementCount returns the TotalEntitlementCount field if non-nil, zero value otherwise.

### GetTotalEntitlementCountOk

`func (o *GetEntitlementResponse) GetTotalEntitlementCountOk() (*int32, bool)`

GetTotalEntitlementCountOk returns a tuple with the TotalEntitlementCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalEntitlementCount

`func (o *GetEntitlementResponse) SetTotalEntitlementCount(v int32)`

SetTotalEntitlementCount sets TotalEntitlementCount field to given value.

### HasTotalEntitlementCount

`func (o *GetEntitlementResponse) HasTotalEntitlementCount() bool`

HasTotalEntitlementCount returns a boolean if a field has been set.

### GetEntitlementsCount

`func (o *GetEntitlementResponse) GetEntitlementsCount() int32`

GetEntitlementsCount returns the EntitlementsCount field if non-nil, zero value otherwise.

### GetEntitlementsCountOk

`func (o *GetEntitlementResponse) GetEntitlementsCountOk() (*int32, bool)`

GetEntitlementsCountOk returns a tuple with the EntitlementsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitlementsCount

`func (o *GetEntitlementResponse) SetEntitlementsCount(v int32)`

SetEntitlementsCount sets EntitlementsCount field to given value.

### HasEntitlementsCount

`func (o *GetEntitlementResponse) HasEntitlementsCount() bool`

HasEntitlementsCount returns a boolean if a field has been set.

### GetEntitlementdetails

`func (o *GetEntitlementResponse) GetEntitlementdetails() []GetEntitlementResponseEntitlementdetailsInner`

GetEntitlementdetails returns the Entitlementdetails field if non-nil, zero value otherwise.

### GetEntitlementdetailsOk

`func (o *GetEntitlementResponse) GetEntitlementdetailsOk() (*[]GetEntitlementResponseEntitlementdetailsInner, bool)`

GetEntitlementdetailsOk returns a tuple with the Entitlementdetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitlementdetails

`func (o *GetEntitlementResponse) SetEntitlementdetails(v []GetEntitlementResponseEntitlementdetailsInner)`

SetEntitlementdetails sets Entitlementdetails field to given value.

### HasEntitlementdetails

`func (o *GetEntitlementResponse) HasEntitlementdetails() bool`

HasEntitlementdetails returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


