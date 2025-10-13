# \EntitlementAPI

All URIs are relative to *http://localhost:3000*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateUpdateEntitlement**](EntitlementAPI.md#CreateUpdateEntitlement) | **Post** /ECM/api/v5/createUpdateEntitlement | Create and update an entitlement
[**GetEntitlements**](EntitlementAPI.md#GetEntitlements) | **Post** /ECM/api/v5/getEntitlements | Get list of entitlements



## CreateUpdateEntitlement

> CreateOrUpdateEntitlementResponse CreateUpdateEntitlement(ctx).CreateUpdateEntitlementRequest(createUpdateEntitlementRequest).Execute()

Create and update an entitlement

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
	createUpdateEntitlementRequest := *openapiclient.NewCreateUpdateEntitlementRequest("sample-endpoint", "sample-entitlement-type", "sample-entitlement-value") // CreateUpdateEntitlementRequest | Request payload for creating/updating an entitlement

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EntitlementAPI.CreateUpdateEntitlement(context.Background()).CreateUpdateEntitlementRequest(createUpdateEntitlementRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EntitlementAPI.CreateUpdateEntitlement``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateUpdateEntitlement`: CreateOrUpdateEntitlementResponse
	fmt.Fprintf(os.Stdout, "Response from `EntitlementAPI.CreateUpdateEntitlement`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateUpdateEntitlementRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createUpdateEntitlementRequest** | [**CreateUpdateEntitlementRequest**](CreateUpdateEntitlementRequest.md) | Request payload for creating/updating an entitlement | 

### Return type

[**CreateOrUpdateEntitlementResponse**](CreateOrUpdateEntitlementResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEntitlements

> GetEntitlementResponse GetEntitlements(ctx).GetEntitlementRequest(getEntitlementRequest).Execute()

Get list of entitlements

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
	getEntitlementRequest := *openapiclient.NewGetEntitlementRequest() // GetEntitlementRequest | Request payload for getting a list of entitlements

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EntitlementAPI.GetEntitlements(context.Background()).GetEntitlementRequest(getEntitlementRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EntitlementAPI.GetEntitlements``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEntitlements`: GetEntitlementResponse
	fmt.Fprintf(os.Stdout, "Response from `EntitlementAPI.GetEntitlements`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetEntitlementsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **getEntitlementRequest** | [**GetEntitlementRequest**](GetEntitlementRequest.md) | Request payload for getting a list of entitlements | 

### Return type

[**GetEntitlementResponse**](GetEntitlementResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

