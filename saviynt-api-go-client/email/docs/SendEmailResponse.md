# SendEmailResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ErrorCode** | **string** | &#x60;\&quot;0\&quot;&#x60; indicates success. Other values indicate failure. | 
**Msg** | **string** |  | 

## Methods

### NewSendEmailResponse

`func NewSendEmailResponse(errorCode string, msg string, ) *SendEmailResponse`

NewSendEmailResponse instantiates a new SendEmailResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSendEmailResponseWithDefaults

`func NewSendEmailResponseWithDefaults() *SendEmailResponse`

NewSendEmailResponseWithDefaults instantiates a new SendEmailResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetErrorCode

`func (o *SendEmailResponse) GetErrorCode() string`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *SendEmailResponse) GetErrorCodeOk() (*string, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *SendEmailResponse) SetErrorCode(v string)`

SetErrorCode sets ErrorCode field to given value.


### GetMsg

`func (o *SendEmailResponse) GetMsg() string`

GetMsg returns the Msg field if non-nil, zero value otherwise.

### GetMsgOk

`func (o *SendEmailResponse) GetMsgOk() (*string, bool)`

GetMsgOk returns a tuple with the Msg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsg

`func (o *SendEmailResponse) SetMsg(v string)`

SetMsg sets Msg field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


