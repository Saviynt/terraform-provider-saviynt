# \MTLSAuthenticationAPI

All URIs are relative to *http://localhost:3000*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteKeyStore**](MTLSAuthenticationAPI.md#DeleteKeyStore) | **Delete** /ECM/api/v5/deleteKeyStoreAlias/{alias} | Delete KeyStore
[**GetKeyStoreCertificateDetails**](MTLSAuthenticationAPI.md#GetKeyStoreCertificateDetails) | **Get** /ECM/api/v5/getKeyStoreCertificateDetails | Get KeyStore Details
[**UploadKeyStore**](MTLSAuthenticationAPI.md#UploadKeyStore) | **Post** /ECM/api/v5/uploadKeyStore | Upload KeyStore



## DeleteKeyStore

> DeleteKeyStore(ctx, alias).Execute()

Delete KeyStore



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
	alias := "alias_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.MTLSAuthenticationAPI.DeleteKeyStore(context.Background(), alias).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MTLSAuthenticationAPI.DeleteKeyStore``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**alias** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteKeyStoreRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetKeyStoreCertificateDetails

> GetKeyStoreCertificateDetailsResponse GetKeyStoreCertificateDetails(ctx).Execute()

Get KeyStore Details



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
	resp, r, err := apiClient.MTLSAuthenticationAPI.GetKeyStoreCertificateDetails(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MTLSAuthenticationAPI.GetKeyStoreCertificateDetails``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetKeyStoreCertificateDetails`: GetKeyStoreCertificateDetailsResponse
	fmt.Fprintf(os.Stdout, "Response from `MTLSAuthenticationAPI.GetKeyStoreCertificateDetails`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetKeyStoreCertificateDetailsRequest struct via the builder pattern


### Return type

[**GetKeyStoreCertificateDetailsResponse**](GetKeyStoreCertificateDetailsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UploadKeyStore

> UploadKeyStoreResponse UploadKeyStore(ctx).KeyStoreFile(keyStoreFile).KeyStorePassword(keyStorePassword).Execute()

Upload KeyStore



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
	keyStoreFile := os.NewFile(1234, "some_file") // *os.File |  (optional)
	keyStorePassword := "keyStorePassword_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MTLSAuthenticationAPI.UploadKeyStore(context.Background()).KeyStoreFile(keyStoreFile).KeyStorePassword(keyStorePassword).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MTLSAuthenticationAPI.UploadKeyStore``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UploadKeyStore`: UploadKeyStoreResponse
	fmt.Fprintf(os.Stdout, "Response from `MTLSAuthenticationAPI.UploadKeyStore`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUploadKeyStoreRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **keyStoreFile** | ***os.File** |  | 
 **keyStorePassword** | **string** |  | 

### Return type

[**UploadKeyStoreResponse**](UploadKeyStoreResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

