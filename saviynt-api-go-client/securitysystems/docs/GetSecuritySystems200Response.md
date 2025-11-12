# GetSecuritySystems200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Msg** | Pointer to **string** | A message indicating the outcome of the operation. | [optional] 
**SecuritySystemDetails** | Pointer to [**[]GetSecuritySystems200ResponseSecuritySystemDetailsInner**](GetSecuritySystems200ResponseSecuritySystemDetailsInner.md) |  | [optional] 
**DisplayCount** | Pointer to **int32** | The number of items currently displayed (e.g., on the current page or view). | [optional] 
**ErrorCode** | Pointer to **string** | An error code where &#39;0&#39; signifies success and &#39;1&#39; signifies an unsuccessful operation. | [optional] 
**TotalCount** | Pointer to **int32** | The total number of items available in the dataset, irrespective of the current display settings. | [optional] 

## Methods

### NewGetSecuritySystems200Response

`func NewGetSecuritySystems200Response() *GetSecuritySystems200Response`

NewGetSecuritySystems200Response instantiates a new GetSecuritySystems200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetSecuritySystems200ResponseWithDefaults

`func NewGetSecuritySystems200ResponseWithDefaults() *GetSecuritySystems200Response`

NewGetSecuritySystems200ResponseWithDefaults instantiates a new GetSecuritySystems200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMsg

`func (o *GetSecuritySystems200Response) GetMsg() string`

GetMsg returns the Msg field if non-nil, zero value otherwise.

### GetMsgOk

`func (o *GetSecuritySystems200Response) GetMsgOk() (*string, bool)`

GetMsgOk returns a tuple with the Msg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsg

`func (o *GetSecuritySystems200Response) SetMsg(v string)`

SetMsg sets Msg field to given value.

### HasMsg

`func (o *GetSecuritySystems200Response) HasMsg() bool`

HasMsg returns a boolean if a field has been set.

### GetSecuritySystemDetails

`func (o *GetSecuritySystems200Response) GetSecuritySystemDetails() []GetSecuritySystems200ResponseSecuritySystemDetailsInner`

GetSecuritySystemDetails returns the SecuritySystemDetails field if non-nil, zero value otherwise.

### GetSecuritySystemDetailsOk

`func (o *GetSecuritySystems200Response) GetSecuritySystemDetailsOk() (*[]GetSecuritySystems200ResponseSecuritySystemDetailsInner, bool)`

GetSecuritySystemDetailsOk returns a tuple with the SecuritySystemDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecuritySystemDetails

`func (o *GetSecuritySystems200Response) SetSecuritySystemDetails(v []GetSecuritySystems200ResponseSecuritySystemDetailsInner)`

SetSecuritySystemDetails sets SecuritySystemDetails field to given value.

### HasSecuritySystemDetails

`func (o *GetSecuritySystems200Response) HasSecuritySystemDetails() bool`

HasSecuritySystemDetails returns a boolean if a field has been set.

### GetDisplayCount

`func (o *GetSecuritySystems200Response) GetDisplayCount() int32`

GetDisplayCount returns the DisplayCount field if non-nil, zero value otherwise.

### GetDisplayCountOk

`func (o *GetSecuritySystems200Response) GetDisplayCountOk() (*int32, bool)`

GetDisplayCountOk returns a tuple with the DisplayCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayCount

`func (o *GetSecuritySystems200Response) SetDisplayCount(v int32)`

SetDisplayCount sets DisplayCount field to given value.

### HasDisplayCount

`func (o *GetSecuritySystems200Response) HasDisplayCount() bool`

HasDisplayCount returns a boolean if a field has been set.

### GetErrorCode

`func (o *GetSecuritySystems200Response) GetErrorCode() string`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *GetSecuritySystems200Response) GetErrorCodeOk() (*string, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *GetSecuritySystems200Response) SetErrorCode(v string)`

SetErrorCode sets ErrorCode field to given value.

### HasErrorCode

`func (o *GetSecuritySystems200Response) HasErrorCode() bool`

HasErrorCode returns a boolean if a field has been set.

### GetTotalCount

`func (o *GetSecuritySystems200Response) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *GetSecuritySystems200Response) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *GetSecuritySystems200Response) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.

### HasTotalCount

`func (o *GetSecuritySystems200Response) HasTotalCount() bool`

HasTotalCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


