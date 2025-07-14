# \DelegatedAdministrationAPI

All URIs are relative to *http://localhost:3000*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDelegate**](DelegatedAdministrationAPI.md#CreateDelegate) | **Post** /ECM/api/v5/createDelegate | Create Delegate
[**DeleteDelegate**](DelegatedAdministrationAPI.md#DeleteDelegate) | **Post** /ECM/api/v5/deleteDelegate | Delete Delegate
[**EditDelegate**](DelegatedAdministrationAPI.md#EditDelegate) | **Post** /ECM/api/v5/editDelegate | Edit Delegate
[**FetchExistingDelegatesList**](DelegatedAdministrationAPI.md#FetchExistingDelegatesList) | **Post** /ECM/api/v5/fetchDelegatesList | Fetch Existing Delegates List
[**GetDelegateUserList**](DelegatedAdministrationAPI.md#GetDelegateUserList) | **Get** /ECM/api/v5/getDelegateUserList | Get Delegate User List



## CreateDelegate

> CreateDelegateResponse CreateDelegate(ctx).CreateDelegateRequest(createDelegateRequest).Execute()

Create Delegate



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
	createDelegateRequest := *openapiclient.NewCreateDelegateRequest("UserName_example", "Name_example", "Delegateusername_example", "Delegatestartdate_example", "Delegateenddate_example") // CreateDelegateRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DelegatedAdministrationAPI.CreateDelegate(context.Background()).CreateDelegateRequest(createDelegateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DelegatedAdministrationAPI.CreateDelegate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateDelegate`: CreateDelegateResponse
	fmt.Fprintf(os.Stdout, "Response from `DelegatedAdministrationAPI.CreateDelegate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateDelegateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createDelegateRequest** | [**CreateDelegateRequest**](CreateDelegateRequest.md) |  | 

### Return type

[**CreateDelegateResponse**](CreateDelegateResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteDelegate

> Response DeleteDelegate(ctx).UserName(userName).Key(key).Execute()

Delete Delegate



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
	userName := "userName_example" // string | this is the user who is deleting the delegate
	key := "key_example" // string | the is the delegatekey

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DelegatedAdministrationAPI.DeleteDelegate(context.Background()).UserName(userName).Key(key).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DelegatedAdministrationAPI.DeleteDelegate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteDelegate`: Response
	fmt.Fprintf(os.Stdout, "Response from `DelegatedAdministrationAPI.DeleteDelegate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDelegateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userName** | **string** | this is the user who is deleting the delegate | 
 **key** | **string** | the is the delegatekey | 

### Return type

[**Response**](Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EditDelegate

> Response EditDelegate(ctx).EditDelegateRequest(editDelegateRequest).Execute()

Edit Delegate



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
	editDelegateRequest := *openapiclient.NewEditDelegateRequest("Key_example", "UserName_example", "Name_example", "Delegateusername_example", "Delegatestartdate_example", "Delegateenddate_example") // EditDelegateRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DelegatedAdministrationAPI.EditDelegate(context.Background()).EditDelegateRequest(editDelegateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DelegatedAdministrationAPI.EditDelegate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EditDelegate`: Response
	fmt.Fprintf(os.Stdout, "Response from `DelegatedAdministrationAPI.EditDelegate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEditDelegateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **editDelegateRequest** | [**EditDelegateRequest**](EditDelegateRequest.md) |  | 

### Return type

[**Response**](Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchExistingDelegatesList

> FetchExistingDelegatesListResponse FetchExistingDelegatesList(ctx).FetchExistingDelegatesListRequest(fetchExistingDelegatesListRequest).Execute()

Fetch Existing Delegates List



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
	fetchExistingDelegatesListRequest := *openapiclient.NewFetchExistingDelegatesListRequest("UserName_example") // FetchExistingDelegatesListRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DelegatedAdministrationAPI.FetchExistingDelegatesList(context.Background()).FetchExistingDelegatesListRequest(fetchExistingDelegatesListRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DelegatedAdministrationAPI.FetchExistingDelegatesList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FetchExistingDelegatesList`: FetchExistingDelegatesListResponse
	fmt.Fprintf(os.Stdout, "Response from `DelegatedAdministrationAPI.FetchExistingDelegatesList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFetchExistingDelegatesListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fetchExistingDelegatesListRequest** | [**FetchExistingDelegatesListRequest**](FetchExistingDelegatesListRequest.md) |  | 

### Return type

[**FetchExistingDelegatesListResponse**](FetchExistingDelegatesListResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDelegateUserList

> GetDelegateUserListResponse GetDelegateUserList(ctx).GetDelegateUserListRequest(getDelegateUserListRequest).Execute()

Get Delegate User List



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
	getDelegateUserListRequest := *openapiclient.NewGetDelegateUserListRequest("Parentusername_example") // GetDelegateUserListRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DelegatedAdministrationAPI.GetDelegateUserList(context.Background()).GetDelegateUserListRequest(getDelegateUserListRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DelegatedAdministrationAPI.GetDelegateUserList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDelegateUserList`: GetDelegateUserListResponse
	fmt.Fprintf(os.Stdout, "Response from `DelegatedAdministrationAPI.GetDelegateUserList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetDelegateUserListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **getDelegateUserListRequest** | [**GetDelegateUserListRequest**](GetDelegateUserListRequest.md) |  | 

### Return type

[**GetDelegateUserListResponse**](GetDelegateUserListResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

