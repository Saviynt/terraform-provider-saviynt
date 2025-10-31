# \PrivilegeAPI

All URIs are relative to *http://localhost:3000*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreatePrivilege**](PrivilegeAPI.md#CreatePrivilege) | **Post** /ECM/api/v5/createPrivilege | Create a privilege
[**DeletePrivilege**](PrivilegeAPI.md#DeletePrivilege) | **Delete** /ECM/api/v5/deletePrivilege | Delete a privilege
[**GetPrivilege**](PrivilegeAPI.md#GetPrivilege) | **Post** /ECM/api/v5/getListofPrivileges | Get a list of privileges
[**UpdatePrivilege**](PrivilegeAPI.md#UpdatePrivilege) | **Put** /ECM/api/v5/updatePrivilege | Update a privilege



## CreatePrivilege

> CreateUpdatePrivilegeResponse CreatePrivilege(ctx).CreateUpdatePrivilegeRequest(createUpdatePrivilegeRequest).Execute()

Create a privilege

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
	createUpdatePrivilegeRequest := *openapiclient.NewCreateUpdatePrivilegeRequest("sample_ss", "sample_endpoint", "sample_ent_type", *openapiclient.NewCreateUpdatePrivilegeRequestPrivilege()) // CreateUpdatePrivilegeRequest | Request payload for creating a privilege

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PrivilegeAPI.CreatePrivilege(context.Background()).CreateUpdatePrivilegeRequest(createUpdatePrivilegeRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PrivilegeAPI.CreatePrivilege``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreatePrivilege`: CreateUpdatePrivilegeResponse
	fmt.Fprintf(os.Stdout, "Response from `PrivilegeAPI.CreatePrivilege`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreatePrivilegeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createUpdatePrivilegeRequest** | [**CreateUpdatePrivilegeRequest**](CreateUpdatePrivilegeRequest.md) | Request payload for creating a privilege | 

### Return type

[**CreateUpdatePrivilegeResponse**](CreateUpdatePrivilegeResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeletePrivilege

> DeletePrivilegeResponse DeletePrivilege(ctx).DeletePrivilegeRequest(deletePrivilegeRequest).Execute()

Delete a privilege

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
	deletePrivilegeRequest := *openapiclient.NewDeletePrivilegeRequest("Securitysystem_example", "Endpoint_example", "Entitlementtype_example", "Privilege_example") // DeletePrivilegeRequest | Request payload for deleting a privilege

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PrivilegeAPI.DeletePrivilege(context.Background()).DeletePrivilegeRequest(deletePrivilegeRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PrivilegeAPI.DeletePrivilege``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeletePrivilege`: DeletePrivilegeResponse
	fmt.Fprintf(os.Stdout, "Response from `PrivilegeAPI.DeletePrivilege`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeletePrivilegeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deletePrivilegeRequest** | [**DeletePrivilegeRequest**](DeletePrivilegeRequest.md) | Request payload for deleting a privilege | 

### Return type

[**DeletePrivilegeResponse**](DeletePrivilegeResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPrivilege

> GetPrivilegeListResponse GetPrivilege(ctx).GetPrivilegeListRequest(getPrivilegeListRequest).Execute()

Get a list of privileges

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
	getPrivilegeListRequest := *openapiclient.NewGetPrivilegeListRequest("Endpoint_example") // GetPrivilegeListRequest | Request payload for getting the list of privileges

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PrivilegeAPI.GetPrivilege(context.Background()).GetPrivilegeListRequest(getPrivilegeListRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PrivilegeAPI.GetPrivilege``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPrivilege`: GetPrivilegeListResponse
	fmt.Fprintf(os.Stdout, "Response from `PrivilegeAPI.GetPrivilege`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPrivilegeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **getPrivilegeListRequest** | [**GetPrivilegeListRequest**](GetPrivilegeListRequest.md) | Request payload for getting the list of privileges | 

### Return type

[**GetPrivilegeListResponse**](GetPrivilegeListResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdatePrivilege

> CreateUpdatePrivilegeResponse UpdatePrivilege(ctx).CreateUpdatePrivilegeRequest(createUpdatePrivilegeRequest).Execute()

Update a privilege

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
	createUpdatePrivilegeRequest := *openapiclient.NewCreateUpdatePrivilegeRequest("sample_ss", "sample_endpoint", "sample_ent_type", *openapiclient.NewCreateUpdatePrivilegeRequestPrivilege()) // CreateUpdatePrivilegeRequest | Request payload for updating a privilege

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PrivilegeAPI.UpdatePrivilege(context.Background()).CreateUpdatePrivilegeRequest(createUpdatePrivilegeRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PrivilegeAPI.UpdatePrivilege``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdatePrivilege`: CreateUpdatePrivilegeResponse
	fmt.Fprintf(os.Stdout, "Response from `PrivilegeAPI.UpdatePrivilege`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePrivilegeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createUpdatePrivilegeRequest** | [**CreateUpdatePrivilegeRequest**](CreateUpdatePrivilegeRequest.md) | Request payload for updating a privilege | 

### Return type

[**CreateUpdatePrivilegeResponse**](CreateUpdatePrivilegeResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

