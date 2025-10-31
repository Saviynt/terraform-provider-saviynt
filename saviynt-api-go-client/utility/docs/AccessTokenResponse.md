# AccessTokenResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Username** | Pointer to **string** |  | [optional] 
**Roles** | Pointer to **[]string** |  | [optional] 
**TokenType** | Pointer to **string** |  | [optional] 
**AccessToken** | Pointer to **string** |  | [optional] 
**ExpiresIn** | Pointer to **int32** |  | [optional] 
**RefreshToken** | Pointer to **string** |  | [optional] 

## Methods

### NewAccessTokenResponse

`func NewAccessTokenResponse() *AccessTokenResponse`

NewAccessTokenResponse instantiates a new AccessTokenResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccessTokenResponseWithDefaults

`func NewAccessTokenResponseWithDefaults() *AccessTokenResponse`

NewAccessTokenResponseWithDefaults instantiates a new AccessTokenResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsername

`func (o *AccessTokenResponse) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *AccessTokenResponse) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *AccessTokenResponse) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *AccessTokenResponse) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetRoles

`func (o *AccessTokenResponse) GetRoles() []string`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *AccessTokenResponse) GetRolesOk() (*[]string, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *AccessTokenResponse) SetRoles(v []string)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *AccessTokenResponse) HasRoles() bool`

HasRoles returns a boolean if a field has been set.

### GetTokenType

`func (o *AccessTokenResponse) GetTokenType() string`

GetTokenType returns the TokenType field if non-nil, zero value otherwise.

### GetTokenTypeOk

`func (o *AccessTokenResponse) GetTokenTypeOk() (*string, bool)`

GetTokenTypeOk returns a tuple with the TokenType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenType

`func (o *AccessTokenResponse) SetTokenType(v string)`

SetTokenType sets TokenType field to given value.

### HasTokenType

`func (o *AccessTokenResponse) HasTokenType() bool`

HasTokenType returns a boolean if a field has been set.

### GetAccessToken

`func (o *AccessTokenResponse) GetAccessToken() string`

GetAccessToken returns the AccessToken field if non-nil, zero value otherwise.

### GetAccessTokenOk

`func (o *AccessTokenResponse) GetAccessTokenOk() (*string, bool)`

GetAccessTokenOk returns a tuple with the AccessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessToken

`func (o *AccessTokenResponse) SetAccessToken(v string)`

SetAccessToken sets AccessToken field to given value.

### HasAccessToken

`func (o *AccessTokenResponse) HasAccessToken() bool`

HasAccessToken returns a boolean if a field has been set.

### GetExpiresIn

`func (o *AccessTokenResponse) GetExpiresIn() int32`

GetExpiresIn returns the ExpiresIn field if non-nil, zero value otherwise.

### GetExpiresInOk

`func (o *AccessTokenResponse) GetExpiresInOk() (*int32, bool)`

GetExpiresInOk returns a tuple with the ExpiresIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresIn

`func (o *AccessTokenResponse) SetExpiresIn(v int32)`

SetExpiresIn sets ExpiresIn field to given value.

### HasExpiresIn

`func (o *AccessTokenResponse) HasExpiresIn() bool`

HasExpiresIn returns a boolean if a field has been set.

### GetRefreshToken

`func (o *AccessTokenResponse) GetRefreshToken() string`

GetRefreshToken returns the RefreshToken field if non-nil, zero value otherwise.

### GetRefreshTokenOk

`func (o *AccessTokenResponse) GetRefreshTokenOk() (*string, bool)`

GetRefreshTokenOk returns a tuple with the RefreshToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreshToken

`func (o *AccessTokenResponse) SetRefreshToken(v string)`

SetRefreshToken sets RefreshToken field to given value.

### HasRefreshToken

`func (o *AccessTokenResponse) HasRefreshToken() bool`

HasRefreshToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


