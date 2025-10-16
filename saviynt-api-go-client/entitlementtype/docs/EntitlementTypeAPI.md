# \EntitlementTypeAPI

All URIs are relative to *http://localhost:3000*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateEntitlementType**](EntitlementTypeAPI.md#CreateEntitlementType) | **Post** /ECM/api/v5/createEntitlementType | Create an entitlement type
[**GetEntitlementType**](EntitlementTypeAPI.md#GetEntitlementType) | **Get** /ECM/api/v5/getEntitlementTypes | Get a list of entitlement types
[**UpdateEntitlementType**](EntitlementTypeAPI.md#UpdateEntitlementType) | **Put** /ECM/api/v5/updateEntitlementType | Update an entitlement type



## CreateEntitlementType

> CreateOrUpdateEntitlementTypeResponse CreateEntitlementType(ctx).CreateEntitlementTypeRequest(createEntitlementTypeRequest).Execute()

Create an entitlement type

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
	createEntitlementTypeRequest := *openapiclient.NewCreateEntitlementTypeRequest("Ent-type-1", "Endpoint-1") // CreateEntitlementTypeRequest | Request payload for creating an entitlement type

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EntitlementTypeAPI.CreateEntitlementType(context.Background()).CreateEntitlementTypeRequest(createEntitlementTypeRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EntitlementTypeAPI.CreateEntitlementType``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateEntitlementType`: CreateOrUpdateEntitlementTypeResponse
	fmt.Fprintf(os.Stdout, "Response from `EntitlementTypeAPI.CreateEntitlementType`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateEntitlementTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createEntitlementTypeRequest** | [**CreateEntitlementTypeRequest**](CreateEntitlementTypeRequest.md) | Request payload for creating an entitlement type | 

### Return type

[**CreateOrUpdateEntitlementTypeResponse**](CreateOrUpdateEntitlementTypeResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEntitlementType

> GetEntitlementTypeResponse GetEntitlementType(ctx).Entitlementname(entitlementname).Max(max).Offset(offset).Endpointname(endpointname).Execute()

Get a list of entitlement types

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
	entitlementname := "entitlementname_example" // string | Name of entitlement type (optional)
	max := "max_example" // string | Maximum number of results to return (optional)
	offset := "offset_example" // string | Offset for pagination (optional)
	endpointname := "endpointname_example" // string | Name of the endpoint to get entitlement types for (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EntitlementTypeAPI.GetEntitlementType(context.Background()).Entitlementname(entitlementname).Max(max).Offset(offset).Endpointname(endpointname).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EntitlementTypeAPI.GetEntitlementType``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEntitlementType`: GetEntitlementTypeResponse
	fmt.Fprintf(os.Stdout, "Response from `EntitlementTypeAPI.GetEntitlementType`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetEntitlementTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **entitlementname** | **string** | Name of entitlement type | 
 **max** | **string** | Maximum number of results to return | 
 **offset** | **string** | Offset for pagination | 
 **endpointname** | **string** | Name of the endpoint to get entitlement types for | 

### Return type

[**GetEntitlementTypeResponse**](GetEntitlementTypeResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateEntitlementType

> CreateOrUpdateEntitlementTypeResponse UpdateEntitlementType(ctx).UpdateEntitlementTypeRequest(updateEntitlementTypeRequest).Execute()

Update an entitlement type

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
	updateEntitlementTypeRequest := *openapiclient.NewUpdateEntitlementTypeRequest("Ent-type-1", "Endpoint-1") // UpdateEntitlementTypeRequest | Request payload for updating an entitlement type

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EntitlementTypeAPI.UpdateEntitlementType(context.Background()).UpdateEntitlementTypeRequest(updateEntitlementTypeRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EntitlementTypeAPI.UpdateEntitlementType``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateEntitlementType`: CreateOrUpdateEntitlementTypeResponse
	fmt.Fprintf(os.Stdout, "Response from `EntitlementTypeAPI.UpdateEntitlementType`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateEntitlementTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateEntitlementTypeRequest** | [**UpdateEntitlementTypeRequest**](UpdateEntitlementTypeRequest.md) | Request payload for updating an entitlement type | 

### Return type

[**CreateOrUpdateEntitlementTypeResponse**](CreateOrUpdateEntitlementTypeResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

