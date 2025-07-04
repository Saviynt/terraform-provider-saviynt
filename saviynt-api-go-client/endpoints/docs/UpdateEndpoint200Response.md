# UpdateEndpoint200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Msg** | Pointer to **string** | A message indicating the outcome of the operation. | [optional] 
**ErrorCode** | Pointer to **string** | An error code where &#39;0&#39; signifies success and &#39;1&#39; signifies an unsuccessful operation. | [optional] 

## Methods

### NewUpdateEndpoint200Response

`func NewUpdateEndpoint200Response() *UpdateEndpoint200Response`

NewUpdateEndpoint200Response instantiates a new UpdateEndpoint200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateEndpoint200ResponseWithDefaults

`func NewUpdateEndpoint200ResponseWithDefaults() *UpdateEndpoint200Response`

NewUpdateEndpoint200ResponseWithDefaults instantiates a new UpdateEndpoint200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMsg

`func (o *UpdateEndpoint200Response) GetMsg() string`

GetMsg returns the Msg field if non-nil, zero value otherwise.

### GetMsgOk

`func (o *UpdateEndpoint200Response) GetMsgOk() (*string, bool)`

GetMsgOk returns a tuple with the Msg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsg

`func (o *UpdateEndpoint200Response) SetMsg(v string)`

SetMsg sets Msg field to given value.

### HasMsg

`func (o *UpdateEndpoint200Response) HasMsg() bool`

HasMsg returns a boolean if a field has been set.

### GetErrorCode

`func (o *UpdateEndpoint200Response) GetErrorCode() string`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *UpdateEndpoint200Response) GetErrorCodeOk() (*string, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *UpdateEndpoint200Response) SetErrorCode(v string)`

SetErrorCode sets ErrorCode field to given value.

### HasErrorCode

`func (o *UpdateEndpoint200Response) HasErrorCode() bool`

HasErrorCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


