# \SecuritySystemsAPI

All URIs are relative to *http://localhost:3000*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSecuritySystem**](SecuritySystemsAPI.md#CreateSecuritySystem) | **Post** /ECM/api/v5/createSecuritySystem | Create Security System
[**GetSecuritySystems**](SecuritySystemsAPI.md#GetSecuritySystems) | **Get** /ECM/api/v5/getSecuritySystems | Get Security Systems
[**UpdateSecuritySystem**](SecuritySystemsAPI.md#UpdateSecuritySystem) | **Put** /ECM/api/v5/updateSecuritySystem | Update Security System



## CreateSecuritySystem

> CreateSecuritySystem200Response CreateSecuritySystem(ctx).CreateSecuritySystemRequest(createSecuritySystemRequest).Execute()

Create Security System



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
	createSecuritySystemRequest := *openapiclient.NewCreateSecuritySystemRequest("sys3", "sys3Disp") // CreateSecuritySystemRequest | Request payload for creating a Security System.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SecuritySystemsAPI.CreateSecuritySystem(context.Background()).CreateSecuritySystemRequest(createSecuritySystemRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecuritySystemsAPI.CreateSecuritySystem``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateSecuritySystem`: CreateSecuritySystem200Response
	fmt.Fprintf(os.Stdout, "Response from `SecuritySystemsAPI.CreateSecuritySystem`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateSecuritySystemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createSecuritySystemRequest** | [**CreateSecuritySystemRequest**](CreateSecuritySystemRequest.md) | Request payload for creating a Security System. | 

### Return type

[**CreateSecuritySystem200Response**](CreateSecuritySystem200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSecuritySystems

> GetSecuritySystems200Response GetSecuritySystems(ctx).Systemname(systemname).Max(max).Offset(offset).Connectionname(connectionname).ConnectionType(connectionType).Execute()

Get Security Systems



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
	systemname := "System1" // string | Filter by security system name. (optional)
	max := int32(4) // int32 | Maximum number of records to return. (optional)
	offset := int32(0) // int32 | Pagination offset. (optional)
	connectionname := "CN_SERP_ECC_A53" // string | Connection name used for reconciliation. (optional)
	connectionType := "SAP" // string | Type of the connection, e.g., SAP. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SecuritySystemsAPI.GetSecuritySystems(context.Background()).Systemname(systemname).Max(max).Offset(offset).Connectionname(connectionname).ConnectionType(connectionType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecuritySystemsAPI.GetSecuritySystems``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSecuritySystems`: GetSecuritySystems200Response
	fmt.Fprintf(os.Stdout, "Response from `SecuritySystemsAPI.GetSecuritySystems`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSecuritySystemsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **systemname** | **string** | Filter by security system name. | 
 **max** | **int32** | Maximum number of records to return. | 
 **offset** | **int32** | Pagination offset. | 
 **connectionname** | **string** | Connection name used for reconciliation. | 
 **connectionType** | **string** | Type of the connection, e.g., SAP. | 

### Return type

[**GetSecuritySystems200Response**](GetSecuritySystems200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSecuritySystem

> CreateSecuritySystem200Response UpdateSecuritySystem(ctx).UpdateSecuritySystemRequest(updateSecuritySystemRequest).Execute()

Update Security System



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
	updateSecuritySystemRequest := *openapiclient.NewUpdateSecuritySystemRequest("sys1", "sys1Disp") // UpdateSecuritySystemRequest | Request payload for updating a Security System.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SecuritySystemsAPI.UpdateSecuritySystem(context.Background()).UpdateSecuritySystemRequest(updateSecuritySystemRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecuritySystemsAPI.UpdateSecuritySystem``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateSecuritySystem`: CreateSecuritySystem200Response
	fmt.Fprintf(os.Stdout, "Response from `SecuritySystemsAPI.UpdateSecuritySystem`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSecuritySystemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateSecuritySystemRequest** | [**UpdateSecuritySystemRequest**](UpdateSecuritySystemRequest.md) | Request payload for updating a Security System. | 

### Return type

[**CreateSecuritySystem200Response**](CreateSecuritySystem200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

