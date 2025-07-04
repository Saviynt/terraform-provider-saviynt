# ConnectionTimeoutConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RetryWait** | Pointer to **int32** | Wait time before retrying a failed connection. | [optional] 
**TokenRefreshMaxTryCount** | Pointer to **int32** | Maximum number of retries for token refresh. | [optional] 
**RetryFailureStatusCode** | Pointer to **int32** |  | [optional] 
**RetryWaitMaxValue** | Pointer to **int32** | Maximum wait time for retries. | [optional] 
**RetryCount** | Pointer to **int32** | Number of retry attempts allowed. | [optional] 
**ReadTimeout** | Pointer to **int32** | Read timeout duration (in seconds). | [optional] 
**ConnectionTimeout** | Pointer to **int32** | Connection timeout duration (in seconds). | [optional] 

## Methods

### NewConnectionTimeoutConfig

`func NewConnectionTimeoutConfig() *ConnectionTimeoutConfig`

NewConnectionTimeoutConfig instantiates a new ConnectionTimeoutConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectionTimeoutConfigWithDefaults

`func NewConnectionTimeoutConfigWithDefaults() *ConnectionTimeoutConfig`

NewConnectionTimeoutConfigWithDefaults instantiates a new ConnectionTimeoutConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRetryWait

`func (o *ConnectionTimeoutConfig) GetRetryWait() int32`

GetRetryWait returns the RetryWait field if non-nil, zero value otherwise.

### GetRetryWaitOk

`func (o *ConnectionTimeoutConfig) GetRetryWaitOk() (*int32, bool)`

GetRetryWaitOk returns a tuple with the RetryWait field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetryWait

`func (o *ConnectionTimeoutConfig) SetRetryWait(v int32)`

SetRetryWait sets RetryWait field to given value.

### HasRetryWait

`func (o *ConnectionTimeoutConfig) HasRetryWait() bool`

HasRetryWait returns a boolean if a field has been set.

### GetTokenRefreshMaxTryCount

`func (o *ConnectionTimeoutConfig) GetTokenRefreshMaxTryCount() int32`

GetTokenRefreshMaxTryCount returns the TokenRefreshMaxTryCount field if non-nil, zero value otherwise.

### GetTokenRefreshMaxTryCountOk

`func (o *ConnectionTimeoutConfig) GetTokenRefreshMaxTryCountOk() (*int32, bool)`

GetTokenRefreshMaxTryCountOk returns a tuple with the TokenRefreshMaxTryCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenRefreshMaxTryCount

`func (o *ConnectionTimeoutConfig) SetTokenRefreshMaxTryCount(v int32)`

SetTokenRefreshMaxTryCount sets TokenRefreshMaxTryCount field to given value.

### HasTokenRefreshMaxTryCount

`func (o *ConnectionTimeoutConfig) HasTokenRefreshMaxTryCount() bool`

HasTokenRefreshMaxTryCount returns a boolean if a field has been set.

### GetRetryFailureStatusCode

`func (o *ConnectionTimeoutConfig) GetRetryFailureStatusCode() int32`

GetRetryFailureStatusCode returns the RetryFailureStatusCode field if non-nil, zero value otherwise.

### GetRetryFailureStatusCodeOk

`func (o *ConnectionTimeoutConfig) GetRetryFailureStatusCodeOk() (*int32, bool)`

GetRetryFailureStatusCodeOk returns a tuple with the RetryFailureStatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetryFailureStatusCode

`func (o *ConnectionTimeoutConfig) SetRetryFailureStatusCode(v int32)`

SetRetryFailureStatusCode sets RetryFailureStatusCode field to given value.

### HasRetryFailureStatusCode

`func (o *ConnectionTimeoutConfig) HasRetryFailureStatusCode() bool`

HasRetryFailureStatusCode returns a boolean if a field has been set.

### GetRetryWaitMaxValue

`func (o *ConnectionTimeoutConfig) GetRetryWaitMaxValue() int32`

GetRetryWaitMaxValue returns the RetryWaitMaxValue field if non-nil, zero value otherwise.

### GetRetryWaitMaxValueOk

`func (o *ConnectionTimeoutConfig) GetRetryWaitMaxValueOk() (*int32, bool)`

GetRetryWaitMaxValueOk returns a tuple with the RetryWaitMaxValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetryWaitMaxValue

`func (o *ConnectionTimeoutConfig) SetRetryWaitMaxValue(v int32)`

SetRetryWaitMaxValue sets RetryWaitMaxValue field to given value.

### HasRetryWaitMaxValue

`func (o *ConnectionTimeoutConfig) HasRetryWaitMaxValue() bool`

HasRetryWaitMaxValue returns a boolean if a field has been set.

### GetRetryCount

`func (o *ConnectionTimeoutConfig) GetRetryCount() int32`

GetRetryCount returns the RetryCount field if non-nil, zero value otherwise.

### GetRetryCountOk

`func (o *ConnectionTimeoutConfig) GetRetryCountOk() (*int32, bool)`

GetRetryCountOk returns a tuple with the RetryCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetryCount

`func (o *ConnectionTimeoutConfig) SetRetryCount(v int32)`

SetRetryCount sets RetryCount field to given value.

### HasRetryCount

`func (o *ConnectionTimeoutConfig) HasRetryCount() bool`

HasRetryCount returns a boolean if a field has been set.

### GetReadTimeout

`func (o *ConnectionTimeoutConfig) GetReadTimeout() int32`

GetReadTimeout returns the ReadTimeout field if non-nil, zero value otherwise.

### GetReadTimeoutOk

`func (o *ConnectionTimeoutConfig) GetReadTimeoutOk() (*int32, bool)`

GetReadTimeoutOk returns a tuple with the ReadTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadTimeout

`func (o *ConnectionTimeoutConfig) SetReadTimeout(v int32)`

SetReadTimeout sets ReadTimeout field to given value.

### HasReadTimeout

`func (o *ConnectionTimeoutConfig) HasReadTimeout() bool`

HasReadTimeout returns a boolean if a field has been set.

### GetConnectionTimeout

`func (o *ConnectionTimeoutConfig) GetConnectionTimeout() int32`

GetConnectionTimeout returns the ConnectionTimeout field if non-nil, zero value otherwise.

### GetConnectionTimeoutOk

`func (o *ConnectionTimeoutConfig) GetConnectionTimeoutOk() (*int32, bool)`

GetConnectionTimeoutOk returns a tuple with the ConnectionTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionTimeout

`func (o *ConnectionTimeoutConfig) SetConnectionTimeout(v int32)`

SetConnectionTimeout sets ConnectionTimeout field to given value.

### HasConnectionTimeout

`func (o *ConnectionTimeoutConfig) HasConnectionTimeout() bool`

HasConnectionTimeout returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


