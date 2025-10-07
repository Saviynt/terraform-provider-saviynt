# \TransportAPI

All URIs are relative to *http://localhost:3000*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ExportTransportPackage**](TransportAPI.md#ExportTransportPackage) | **Post** /ECM/api/v5/exportTransportPackage | Export Transport Package
[**ImportTransportPackage**](TransportAPI.md#ImportTransportPackage) | **Post** /ECM/api/v5/importTransportPackage | Import Transport Package
[**TransportPackageStatus**](TransportAPI.md#TransportPackageStatus) | **Get** /ECM/api/v5/transportPackageStatus | Transport Package Status



## ExportTransportPackage

> ExportTransportPackageResponse ExportTransportPackage(ctx).ExportTransportPackageRequest(exportTransportPackageRequest).Execute()

Export Transport Package



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
	exportTransportPackageRequest := *openapiclient.NewExportTransportPackageRequest("Exportonline_example", *openapiclient.NewObjectsToExport()) // ExportTransportPackageRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TransportAPI.ExportTransportPackage(context.Background()).ExportTransportPackageRequest(exportTransportPackageRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TransportAPI.ExportTransportPackage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ExportTransportPackage`: ExportTransportPackageResponse
	fmt.Fprintf(os.Stdout, "Response from `TransportAPI.ExportTransportPackage`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiExportTransportPackageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **exportTransportPackageRequest** | [**ExportTransportPackageRequest**](ExportTransportPackageRequest.md) |  | 

### Return type

[**ExportTransportPackageResponse**](ExportTransportPackageResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ImportTransportPackage

> ImportTransportPackageResponse ImportTransportPackage(ctx).ImportTransportPackageRequest(importTransportPackageRequest).Execute()

Import Transport Package



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
	importTransportPackageRequest := *openapiclient.NewImportTransportPackageRequest("Packagetoimport_example") // ImportTransportPackageRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TransportAPI.ImportTransportPackage(context.Background()).ImportTransportPackageRequest(importTransportPackageRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TransportAPI.ImportTransportPackage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ImportTransportPackage`: ImportTransportPackageResponse
	fmt.Fprintf(os.Stdout, "Response from `TransportAPI.ImportTransportPackage`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiImportTransportPackageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **importTransportPackageRequest** | [**ImportTransportPackageRequest**](ImportTransportPackageRequest.md) |  | 

### Return type

[**ImportTransportPackageResponse**](ImportTransportPackageResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TransportPackageStatus

> TransportPackageStatusResponse TransportPackageStatus(ctx).TransportPackageStatusRequest(transportPackageStatusRequest).Execute()

Transport Package Status



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
	transportPackageStatusRequest := *openapiclient.NewTransportPackageStatusRequest("Operation_example", "Filename_example") // TransportPackageStatusRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TransportAPI.TransportPackageStatus(context.Background()).TransportPackageStatusRequest(transportPackageStatusRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TransportAPI.TransportPackageStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TransportPackageStatus`: TransportPackageStatusResponse
	fmt.Fprintf(os.Stdout, "Response from `TransportAPI.TransportPackageStatus`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTransportPackageStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **transportPackageStatusRequest** | [**TransportPackageStatusRequest**](TransportPackageStatusRequest.md) |  | 

### Return type

[**TransportPackageStatusResponse**](TransportPackageStatusResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

