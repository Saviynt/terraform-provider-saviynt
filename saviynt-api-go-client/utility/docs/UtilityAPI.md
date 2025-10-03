# \UtilityAPI

All URIs are relative to *http://localhost:3000*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AccessToken**](UtilityAPI.md#AccessToken) | **Post** /ECM/oauth/access_token | This API call can be used for getting the new access token
[**GetEcmVersion**](UtilityAPI.md#GetEcmVersion) | **Get** /ECM/api/v5/getEcmVersion | This API call can be used for getting the ecm version



## AccessToken

> AccessTokenResponse AccessToken(ctx).GrantType(grantType).RefreshToken(refreshToken).Execute()

This API call can be used for getting the new access token

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	grantType := "grantType_example" // string | 
	refreshToken := "refreshToken_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UtilityAPI.AccessToken(context.Background()).GrantType(grantType).RefreshToken(refreshToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UtilityAPI.AccessToken``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AccessToken`: AccessTokenResponse
	fmt.Fprintf(os.Stdout, "Response from `UtilityAPI.AccessToken`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAccessTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **grantType** | **string** |  | 
 **refreshToken** | **string** |  | 

### Return type

[**AccessTokenResponse**](AccessTokenResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEcmVersion

> GetEcmVersionResponse GetEcmVersion(ctx).Execute()

This API call can be used for getting the ecm version

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UtilityAPI.GetEcmVersion(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UtilityAPI.GetEcmVersion``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEcmVersion`: GetEcmVersionResponse
	fmt.Fprintf(os.Stdout, "Response from `UtilityAPI.GetEcmVersion`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetEcmVersionRequest struct via the builder pattern


### Return type

[**GetEcmVersionResponse**](GetEcmVersionResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

