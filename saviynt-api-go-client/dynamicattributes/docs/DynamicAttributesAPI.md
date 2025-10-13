# \DynamicAttributesAPI

All URIs are relative to *http://localhost:3000*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDynamicAttribute**](DynamicAttributesAPI.md#CreateDynamicAttribute) | **Post** /ECM/api/v5/createDynamicAttribute | Create a dynamic attribute
[**DeleteDynamicAttribute**](DynamicAttributesAPI.md#DeleteDynamicAttribute) | **Delete** /ECM/api/v5/deleteDynamicAttribute | Delete a dynamic attribute
[**FetchDynamicAttribute**](DynamicAttributesAPI.md#FetchDynamicAttribute) | **Get** /ECM/api/v5/fetchDynamicAttribute | Fetch the dynamic attributes based on a given filter value or all
[**UpdateDynamicAttribute**](DynamicAttributesAPI.md#UpdateDynamicAttribute) | **Put** /ECM/api/v5/updateDynamicAttribute | Update a dynamic attribute



## CreateDynamicAttribute

> CreateOrUpdateOrDeleteDynamicAttributeResponse CreateDynamicAttribute(ctx).CreateDynamicAttributeRequest(createDynamicAttributeRequest).Execute()

Create a dynamic attribute

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
	createDynamicAttributeRequest := *openapiclient.NewCreateDynamicAttributeRequest("sample-system", "sample-endpoint", "Updateuser_example", []openapiclient.CreateDynamicAttributesInner{*openapiclient.NewCreateDynamicAttributesInner("sample-attribute", "SERVICE ACCOUNT")}) // CreateDynamicAttributeRequest | Request payload for creating a dynamic attribute.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DynamicAttributesAPI.CreateDynamicAttribute(context.Background()).CreateDynamicAttributeRequest(createDynamicAttributeRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DynamicAttributesAPI.CreateDynamicAttribute``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateDynamicAttribute`: CreateOrUpdateOrDeleteDynamicAttributeResponse
	fmt.Fprintf(os.Stdout, "Response from `DynamicAttributesAPI.CreateDynamicAttribute`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateDynamicAttributeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createDynamicAttributeRequest** | [**CreateDynamicAttributeRequest**](CreateDynamicAttributeRequest.md) | Request payload for creating a dynamic attribute. | 

### Return type

[**CreateOrUpdateOrDeleteDynamicAttributeResponse**](CreateOrUpdateOrDeleteDynamicAttributeResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteDynamicAttribute

> CreateOrUpdateOrDeleteDynamicAttributeResponse DeleteDynamicAttribute(ctx).DeleteDynamicAttributeRequest(deleteDynamicAttributeRequest).Execute()

Delete a dynamic attribute

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
	deleteDynamicAttributeRequest := *openapiclient.NewDeleteDynamicAttributeRequest("System1", "System1", "admin-user", []string{"Dynamicattributes_example"}) // DeleteDynamicAttributeRequest | Request payload for deleting dynamic attributes.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DynamicAttributesAPI.DeleteDynamicAttribute(context.Background()).DeleteDynamicAttributeRequest(deleteDynamicAttributeRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DynamicAttributesAPI.DeleteDynamicAttribute``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteDynamicAttribute`: CreateOrUpdateOrDeleteDynamicAttributeResponse
	fmt.Fprintf(os.Stdout, "Response from `DynamicAttributesAPI.DeleteDynamicAttribute`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDynamicAttributeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deleteDynamicAttributeRequest** | [**DeleteDynamicAttributeRequest**](DeleteDynamicAttributeRequest.md) | Request payload for deleting dynamic attributes. | 

### Return type

[**CreateOrUpdateOrDeleteDynamicAttributeResponse**](CreateOrUpdateOrDeleteDynamicAttributeResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchDynamicAttribute

> FetchDynamicAttributesResponse FetchDynamicAttribute(ctx).Securitysystem(securitysystem).Endpoint(endpoint).Dynamicattributes(dynamicattributes).Requesttype(requesttype).Offset(offset).Max(max).Loggedinuser(loggedinuser).Execute()

Fetch the dynamic attributes based on a given filter value or all

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
	securitysystem := []string{"Inner_example"} // []string | List of security systems to filter (optional)
	endpoint := []string{"Inner_example"} // []string | List of endpoints to filter (optional)
	dynamicattributes := []string{"Inner_example"} // []string | List of dynamic attribute names (optional)
	requesttype := []string{"Inner_example"} // []string | Types of request (ACCOUNT, etc.) (optional)
	offset := "offset_example" // string | Pagination offset (optional)
	max := "max_example" // string | Maximum number of results (optional)
	loggedinuser := "loggedinuser_example" // string | Username of the logged-in user (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DynamicAttributesAPI.FetchDynamicAttribute(context.Background()).Securitysystem(securitysystem).Endpoint(endpoint).Dynamicattributes(dynamicattributes).Requesttype(requesttype).Offset(offset).Max(max).Loggedinuser(loggedinuser).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DynamicAttributesAPI.FetchDynamicAttribute``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FetchDynamicAttribute`: FetchDynamicAttributesResponse
	fmt.Fprintf(os.Stdout, "Response from `DynamicAttributesAPI.FetchDynamicAttribute`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFetchDynamicAttributeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **securitysystem** | **[]string** | List of security systems to filter | 
 **endpoint** | **[]string** | List of endpoints to filter | 
 **dynamicattributes** | **[]string** | List of dynamic attribute names | 
 **requesttype** | **[]string** | Types of request (ACCOUNT, etc.) | 
 **offset** | **string** | Pagination offset | 
 **max** | **string** | Maximum number of results | 
 **loggedinuser** | **string** | Username of the logged-in user | 

### Return type

[**FetchDynamicAttributesResponse**](FetchDynamicAttributesResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDynamicAttribute

> CreateOrUpdateOrDeleteDynamicAttributeResponse UpdateDynamicAttribute(ctx).UpdateDynamicAttributeRequest(updateDynamicAttributeRequest).Execute()

Update a dynamic attribute

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
	updateDynamicAttributeRequest := *openapiclient.NewUpdateDynamicAttributeRequest("sample-system", "sample-endpoint", "username", []openapiclient.UpdateDynamicAttributesInner{*openapiclient.NewUpdateDynamicAttributesInner("sample-attribute")}) // UpdateDynamicAttributeRequest | Request payload for updating a dynamic attribute.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DynamicAttributesAPI.UpdateDynamicAttribute(context.Background()).UpdateDynamicAttributeRequest(updateDynamicAttributeRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DynamicAttributesAPI.UpdateDynamicAttribute``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateDynamicAttribute`: CreateOrUpdateOrDeleteDynamicAttributeResponse
	fmt.Fprintf(os.Stdout, "Response from `DynamicAttributesAPI.UpdateDynamicAttribute`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDynamicAttributeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateDynamicAttributeRequest** | [**UpdateDynamicAttributeRequest**](UpdateDynamicAttributeRequest.md) | Request payload for updating a dynamic attribute. | 

### Return type

[**CreateOrUpdateOrDeleteDynamicAttributeResponse**](CreateOrUpdateOrDeleteDynamicAttributeResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

