# GetEndpoints200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | Pointer to **string** | A message indicating the outcome of the operation. | [optional] 
**ErrorCode** | Pointer to **string** | An error code where &#39;0&#39; signifies success and &#39;1&#39; signifies an unsuccessful operation. | [optional] 
**DisplayCount** | Pointer to **int32** | The number of items currently displayed (e.g., on the current page or view). | [optional] 
**TotalCount** | Pointer to **int32** | The total number of items available in the dataset, irrespective of the current display settings. | [optional] 
**Endpoints** | Pointer to [**[]GetEndpoints200ResponseEndpointsInner**](GetEndpoints200ResponseEndpointsInner.md) | A collection of endpoints. | [optional] 

## Methods

### NewGetEndpoints200Response

`func NewGetEndpoints200Response() *GetEndpoints200Response`

NewGetEndpoints200Response instantiates a new GetEndpoints200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetEndpoints200ResponseWithDefaults

`func NewGetEndpoints200ResponseWithDefaults() *GetEndpoints200Response`

NewGetEndpoints200ResponseWithDefaults instantiates a new GetEndpoints200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *GetEndpoints200Response) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *GetEndpoints200Response) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *GetEndpoints200Response) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *GetEndpoints200Response) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetErrorCode

`func (o *GetEndpoints200Response) GetErrorCode() string`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *GetEndpoints200Response) GetErrorCodeOk() (*string, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *GetEndpoints200Response) SetErrorCode(v string)`

SetErrorCode sets ErrorCode field to given value.

### HasErrorCode

`func (o *GetEndpoints200Response) HasErrorCode() bool`

HasErrorCode returns a boolean if a field has been set.

### GetDisplayCount

`func (o *GetEndpoints200Response) GetDisplayCount() int32`

GetDisplayCount returns the DisplayCount field if non-nil, zero value otherwise.

### GetDisplayCountOk

`func (o *GetEndpoints200Response) GetDisplayCountOk() (*int32, bool)`

GetDisplayCountOk returns a tuple with the DisplayCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayCount

`func (o *GetEndpoints200Response) SetDisplayCount(v int32)`

SetDisplayCount sets DisplayCount field to given value.

### HasDisplayCount

`func (o *GetEndpoints200Response) HasDisplayCount() bool`

HasDisplayCount returns a boolean if a field has been set.

### GetTotalCount

`func (o *GetEndpoints200Response) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *GetEndpoints200Response) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *GetEndpoints200Response) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.

### HasTotalCount

`func (o *GetEndpoints200Response) HasTotalCount() bool`

HasTotalCount returns a boolean if a field has been set.

### GetEndpoints

`func (o *GetEndpoints200Response) GetEndpoints() []GetEndpoints200ResponseEndpointsInner`

GetEndpoints returns the Endpoints field if non-nil, zero value otherwise.

### GetEndpointsOk

`func (o *GetEndpoints200Response) GetEndpointsOk() (*[]GetEndpoints200ResponseEndpointsInner, bool)`

GetEndpointsOk returns a tuple with the Endpoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoints

`func (o *GetEndpoints200Response) SetEndpoints(v []GetEndpoints200ResponseEndpointsInner)`

SetEndpoints sets Endpoints field to given value.

### HasEndpoints

`func (o *GetEndpoints200Response) HasEndpoints() bool`

HasEndpoints returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


