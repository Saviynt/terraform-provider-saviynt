# \EndpointsAPI

All URIs are relative to *http://localhost:3000*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateEndpoint**](EndpointsAPI.md#CreateEndpoint) | **Post** /ECM/api/v5/createEndpoint | Create endpoint
[**GetEndpoints**](EndpointsAPI.md#GetEndpoints) | **Post** /ECM/api/v5/getEndpoints | Get list of endpoints
[**UpdateEndpoint**](EndpointsAPI.md#UpdateEndpoint) | **Put** /ECM/api/v5/updateEndpoint | Update endpoint



## CreateEndpoint

> UpdateEndpoint200Response CreateEndpoint(ctx).CreateEndpointRequest(createEndpointRequest).Execute()

Create endpoint

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
	createEndpointRequest := *openapiclient.NewCreateEndpointRequest("Create-an-endpoint", "new-endpoint", "connectiontest") // CreateEndpointRequest | Request payload for creating an endpoint.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EndpointsAPI.CreateEndpoint(context.Background()).CreateEndpointRequest(createEndpointRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EndpointsAPI.CreateEndpoint``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateEndpoint`: UpdateEndpoint200Response
	fmt.Fprintf(os.Stdout, "Response from `EndpointsAPI.CreateEndpoint`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateEndpointRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createEndpointRequest** | [**CreateEndpointRequest**](CreateEndpointRequest.md) | Request payload for creating an endpoint. | 

### Return type

[**UpdateEndpoint200Response**](UpdateEndpoint200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEndpoints

> GetEndpoints200Response GetEndpoints(ctx).GetEndpointsRequest(getEndpointsRequest).Execute()

Get list of endpoints

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
	getEndpointsRequest := *openapiclient.NewGetEndpointsRequest() // GetEndpointsRequest | Request payload for listing enpoints.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EndpointsAPI.GetEndpoints(context.Background()).GetEndpointsRequest(getEndpointsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EndpointsAPI.GetEndpoints``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEndpoints`: GetEndpoints200Response
	fmt.Fprintf(os.Stdout, "Response from `EndpointsAPI.GetEndpoints`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetEndpointsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **getEndpointsRequest** | [**GetEndpointsRequest**](GetEndpointsRequest.md) | Request payload for listing enpoints. | 

### Return type

[**GetEndpoints200Response**](GetEndpoints200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateEndpoint

> UpdateEndpoint200Response UpdateEndpoint(ctx).UpdateEndpointRequest(updateEndpointRequest).Execute()

Update endpoint

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
	updateEndpointRequest := *openapiclient.NewUpdateEndpointRequest("Create-an-endpoint") // UpdateEndpointRequest | Request payload for updating an endpoint.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EndpointsAPI.UpdateEndpoint(context.Background()).UpdateEndpointRequest(updateEndpointRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EndpointsAPI.UpdateEndpoint``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateEndpoint`: UpdateEndpoint200Response
	fmt.Fprintf(os.Stdout, "Response from `EndpointsAPI.UpdateEndpoint`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateEndpointRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateEndpointRequest** | [**UpdateEndpointRequest**](UpdateEndpointRequest.md) | Request payload for updating an endpoint. | 

### Return type

[**UpdateEndpoint200Response**](UpdateEndpoint200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

