# GetEntitlementTypeResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Msg** | Pointer to **string** | A message indicating the outcome of the operation. | [optional] 
**ErrorCode** | Pointer to **string** | An error code where &#39;0&#39; signifies success and &#39;1&#39; signifies an unsuccessful operation. | [optional] 
**DisplayCount** | Pointer to **int32** | Total number of records displayed. | [optional] 
**TotalCount** | Pointer to **int32** | Total number of records available. | [optional] 
**EntitlementTypeDetails** | Pointer to [**[]GetEntitlementTypeResponseEntitlementTypeDetailsInner**](GetEntitlementTypeResponseEntitlementTypeDetailsInner.md) |  | [optional] 

## Methods

### NewGetEntitlementTypeResponse

`func NewGetEntitlementTypeResponse() *GetEntitlementTypeResponse`

NewGetEntitlementTypeResponse instantiates a new GetEntitlementTypeResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetEntitlementTypeResponseWithDefaults

`func NewGetEntitlementTypeResponseWithDefaults() *GetEntitlementTypeResponse`

NewGetEntitlementTypeResponseWithDefaults instantiates a new GetEntitlementTypeResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMsg

`func (o *GetEntitlementTypeResponse) GetMsg() string`

GetMsg returns the Msg field if non-nil, zero value otherwise.

### GetMsgOk

`func (o *GetEntitlementTypeResponse) GetMsgOk() (*string, bool)`

GetMsgOk returns a tuple with the Msg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsg

`func (o *GetEntitlementTypeResponse) SetMsg(v string)`

SetMsg sets Msg field to given value.

### HasMsg

`func (o *GetEntitlementTypeResponse) HasMsg() bool`

HasMsg returns a boolean if a field has been set.

### GetErrorCode

`func (o *GetEntitlementTypeResponse) GetErrorCode() string`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *GetEntitlementTypeResponse) GetErrorCodeOk() (*string, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *GetEntitlementTypeResponse) SetErrorCode(v string)`

SetErrorCode sets ErrorCode field to given value.

### HasErrorCode

`func (o *GetEntitlementTypeResponse) HasErrorCode() bool`

HasErrorCode returns a boolean if a field has been set.

### GetDisplayCount

`func (o *GetEntitlementTypeResponse) GetDisplayCount() int32`

GetDisplayCount returns the DisplayCount field if non-nil, zero value otherwise.

### GetDisplayCountOk

`func (o *GetEntitlementTypeResponse) GetDisplayCountOk() (*int32, bool)`

GetDisplayCountOk returns a tuple with the DisplayCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayCount

`func (o *GetEntitlementTypeResponse) SetDisplayCount(v int32)`

SetDisplayCount sets DisplayCount field to given value.

### HasDisplayCount

`func (o *GetEntitlementTypeResponse) HasDisplayCount() bool`

HasDisplayCount returns a boolean if a field has been set.

### GetTotalCount

`func (o *GetEntitlementTypeResponse) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *GetEntitlementTypeResponse) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *GetEntitlementTypeResponse) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.

### HasTotalCount

`func (o *GetEntitlementTypeResponse) HasTotalCount() bool`

HasTotalCount returns a boolean if a field has been set.

### GetEntitlementTypeDetails

`func (o *GetEntitlementTypeResponse) GetEntitlementTypeDetails() []GetEntitlementTypeResponseEntitlementTypeDetailsInner`

GetEntitlementTypeDetails returns the EntitlementTypeDetails field if non-nil, zero value otherwise.

### GetEntitlementTypeDetailsOk

`func (o *GetEntitlementTypeResponse) GetEntitlementTypeDetailsOk() (*[]GetEntitlementTypeResponseEntitlementTypeDetailsInner, bool)`

GetEntitlementTypeDetailsOk returns a tuple with the EntitlementTypeDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitlementTypeDetails

`func (o *GetEntitlementTypeResponse) SetEntitlementTypeDetails(v []GetEntitlementTypeResponseEntitlementTypeDetailsInner)`

SetEntitlementTypeDetails sets EntitlementTypeDetails field to given value.

### HasEntitlementTypeDetails

`func (o *GetEntitlementTypeResponse) HasEntitlementTypeDetails() bool`

HasEntitlementTypeDetails returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


