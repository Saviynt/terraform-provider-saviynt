# \RolesAPI

All URIs are relative to *http://localhost:3000*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Addrole**](RolesAPI.md#Addrole) | **Post** /ECM/api/v5/addrole | This API adds role to user.
[**CreateEnterpriseRoleRequest**](RolesAPI.md#CreateEnterpriseRoleRequest) | **Post** /ECM/api/v5/createEnterpriseRoleRequest | This API call can be used for creating a new role and assigning an owner to the role
[**GetFireFighterRoles**](RolesAPI.md#GetFireFighterRoles) | **Post** /ECM/api/v5/getFireFighterRoles | Get FireFighter Roles
[**GetRoles**](RolesAPI.md#GetRoles) | **Post** /ECM/api/v5/getRoles | This API can be used to get the list of all the roles
[**Removerole**](RolesAPI.md#Removerole) | **Post** /ECM/api/v5/removerole | This API removes role from user.
[**UpdateEnterpriseRoleRequest**](RolesAPI.md#UpdateEnterpriseRoleRequest) | **Post** /ECM/api/v5/updateEnterpriseRoleRequest | This API call can be used for creating a new role and assigning an owner to the role



## Addrole

> AddOrRemoveRoleResponse Addrole(ctx).AddOrRemoveRoleRequest(addOrRemoveRoleRequest).Execute()

This API adds role to user.

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
	addOrRemoveRoleRequest := *openapiclient.NewAddOrRemoveRoleRequest("johndoe", "Fire Fighter") // AddOrRemoveRoleRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RolesAPI.Addrole(context.Background()).AddOrRemoveRoleRequest(addOrRemoveRoleRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RolesAPI.Addrole``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Addrole`: AddOrRemoveRoleResponse
	fmt.Fprintf(os.Stdout, "Response from `RolesAPI.Addrole`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddroleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **addOrRemoveRoleRequest** | [**AddOrRemoveRoleRequest**](AddOrRemoveRoleRequest.md) |  | 

### Return type

[**AddOrRemoveRoleResponse**](AddOrRemoveRoleResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateEnterpriseRoleRequest

> CreateEnterpriseRoleResponse CreateEnterpriseRoleRequest(ctx).CreateEnterpriseRoleRequest(createEnterpriseRoleRequest).Execute()

This API call can be used for creating a new role and assigning an owner to the role

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
	createEnterpriseRoleRequest := *openapiclient.NewCreateEnterpriseRoleRequest("Roletype_example", "RoleName_example", "Requestor_example", []openapiclient.CreateRoleOwnerPayload{*openapiclient.NewCreateRoleOwnerPayload()}) // CreateEnterpriseRoleRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RolesAPI.CreateEnterpriseRoleRequest(context.Background()).CreateEnterpriseRoleRequest(createEnterpriseRoleRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RolesAPI.CreateEnterpriseRoleRequest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateEnterpriseRoleRequest`: CreateEnterpriseRoleResponse
	fmt.Fprintf(os.Stdout, "Response from `RolesAPI.CreateEnterpriseRoleRequest`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateEnterpriseRoleRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createEnterpriseRoleRequest** | [**CreateEnterpriseRoleRequest**](CreateEnterpriseRoleRequest.md) |  | 

### Return type

[**CreateEnterpriseRoleResponse**](CreateEnterpriseRoleResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFireFighterRoles

> []GetFireFighterRole GetFireFighterRoles(ctx).Execute()

Get FireFighter Roles

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
	resp, r, err := apiClient.RolesAPI.GetFireFighterRoles(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RolesAPI.GetFireFighterRoles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFireFighterRoles`: []GetFireFighterRole
	fmt.Fprintf(os.Stdout, "Response from `RolesAPI.GetFireFighterRoles`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetFireFighterRolesRequest struct via the builder pattern


### Return type

[**[]GetFireFighterRole**](GetFireFighterRole.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRoles

> GetRolesResponse GetRoles(ctx).GetRolesRequest(getRolesRequest).Execute()

This API can be used to get the list of all the roles

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
	getRolesRequest := *openapiclient.NewGetRolesRequest() // GetRolesRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RolesAPI.GetRoles(context.Background()).GetRolesRequest(getRolesRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RolesAPI.GetRoles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRoles`: GetRolesResponse
	fmt.Fprintf(os.Stdout, "Response from `RolesAPI.GetRoles`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetRolesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **getRolesRequest** | [**GetRolesRequest**](GetRolesRequest.md) |  | 

### Return type

[**GetRolesResponse**](GetRolesResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Removerole

> AddOrRemoveRoleResponse Removerole(ctx).AddOrRemoveRoleRequest(addOrRemoveRoleRequest).Execute()

This API removes role from user.

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
	addOrRemoveRoleRequest := *openapiclient.NewAddOrRemoveRoleRequest("johndoe", "Fire Fighter") // AddOrRemoveRoleRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RolesAPI.Removerole(context.Background()).AddOrRemoveRoleRequest(addOrRemoveRoleRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RolesAPI.Removerole``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Removerole`: AddOrRemoveRoleResponse
	fmt.Fprintf(os.Stdout, "Response from `RolesAPI.Removerole`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRemoveroleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **addOrRemoveRoleRequest** | [**AddOrRemoveRoleRequest**](AddOrRemoveRoleRequest.md) |  | 

### Return type

[**AddOrRemoveRoleResponse**](AddOrRemoveRoleResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateEnterpriseRoleRequest

> UpdateEnterpriseRoleResponse UpdateEnterpriseRoleRequest(ctx).UpdateEnterpriseRoleRequest(updateEnterpriseRoleRequest).Execute()

This API call can be used for creating a new role and assigning an owner to the role

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
	updateEnterpriseRoleRequest := *openapiclient.NewUpdateEnterpriseRoleRequest("Roletype_example", "RoleName_example") // UpdateEnterpriseRoleRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RolesAPI.UpdateEnterpriseRoleRequest(context.Background()).UpdateEnterpriseRoleRequest(updateEnterpriseRoleRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RolesAPI.UpdateEnterpriseRoleRequest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateEnterpriseRoleRequest`: UpdateEnterpriseRoleResponse
	fmt.Fprintf(os.Stdout, "Response from `RolesAPI.UpdateEnterpriseRoleRequest`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateEnterpriseRoleRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateEnterpriseRoleRequest** | [**UpdateEnterpriseRoleRequest**](UpdateEnterpriseRoleRequest.md) |  | 

### Return type

[**UpdateEnterpriseRoleResponse**](UpdateEnterpriseRoleResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

