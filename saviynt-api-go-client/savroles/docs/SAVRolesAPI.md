# \SAVRolesAPI

All URIs are relative to *http://localhost:3000*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAllSAVRoles**](SAVRolesAPI.md#GetAllSAVRoles) | **Get** /ECMv6/api/userms/savroles | Get All SAV Roles
[**GetSAVRoleUsers**](SAVRolesAPI.md#GetSAVRoleUsers) | **Get** /ECMv6/api/userms/savroles/{savRoleName}/users | Get Users Associated with a Particular SAV Role



## GetAllSAVRoles

> GetAllSAVRolesResponse GetAllSAVRoles(ctx).Execute()

Get All SAV Roles



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
	resp, r, err := apiClient.SAVRolesAPI.GetAllSAVRoles(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SAVRolesAPI.GetAllSAVRoles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAllSAVRoles`: GetAllSAVRolesResponse
	fmt.Fprintf(os.Stdout, "Response from `SAVRolesAPI.GetAllSAVRoles`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAllSAVRolesRequest struct via the builder pattern


### Return type

[**GetAllSAVRolesResponse**](GetAllSAVRolesResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSAVRoleUsers

> GetSAVRoleUsersResponse GetSAVRoleUsers(ctx, savRoleName).Limit(limit).Offset(offset).Execute()

Get Users Associated with a Particular SAV Role



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
	limit := "limit_example" // string | Specifies the number of retrieved results (default to "0")
	offset := "offset_example" // string | Specifies the number of rows of the result to skip before any rows are retrieved, and must be used with the `limit` parameter (default to "1000")
	savRoleName := "savRoleName_example" // string | The `ROLENAME` field in geAllSAVRoles API

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SAVRolesAPI.GetSAVRoleUsers(context.Background(), savRoleName).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SAVRolesAPI.GetSAVRoleUsers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSAVRoleUsers`: GetSAVRoleUsersResponse
	fmt.Fprintf(os.Stdout, "Response from `SAVRolesAPI.GetSAVRoleUsers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**savRoleName** | **string** | The &#x60;ROLENAME&#x60; field in geAllSAVRoles API | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSAVRoleUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **string** | Specifies the number of retrieved results | [default to &quot;0&quot;]
 **offset** | **string** | Specifies the number of rows of the result to skip before any rows are retrieved, and must be used with the &#x60;limit&#x60; parameter | [default to &quot;1000&quot;]


### Return type

[**GetSAVRoleUsersResponse**](GetSAVRoleUsersResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

