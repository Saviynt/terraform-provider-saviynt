# \TasksAPI

All URIs are relative to *http://localhost:3000*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CheckTaskStatus**](TasksAPI.md#CheckTaskStatus) | **Post** /ECM/api/v5/checkTaskStatus | Check Task Status
[**UpdateTasks**](TasksAPI.md#UpdateTasks) | **Post** /ECM/api/v5/updateTasks | Update Tasks



## CheckTaskStatus

> CheckTaskStatusResponse CheckTaskStatus(ctx).Taskid(taskid).Execute()

Check Task Status



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
	taskid := "taskid_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TasksAPI.CheckTaskStatus(context.Background()).Taskid(taskid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TasksAPI.CheckTaskStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CheckTaskStatus`: CheckTaskStatusResponse
	fmt.Fprintf(os.Stdout, "Response from `TasksAPI.CheckTaskStatus`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCheckTaskStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **taskid** | **string** |  | 

### Return type

[**CheckTaskStatusResponse**](CheckTaskStatusResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateTasks

> map[string]UpdateTaskResponseInfo UpdateTasks(ctx).UpdateTasksRequest(updateTasksRequest).Execute()

Update Tasks



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
	updateTasksRequest := *openapiclient.NewUpdateTasksRequest([]openapiclient.UpdateTaskRequestInfo{*openapiclient.NewUpdateTaskRequestInfo("Taskid_example", "Status_example")}) // UpdateTasksRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TasksAPI.UpdateTasks(context.Background()).UpdateTasksRequest(updateTasksRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TasksAPI.UpdateTasks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateTasks`: map[string]UpdateTaskResponseInfo
	fmt.Fprintf(os.Stdout, "Response from `TasksAPI.UpdateTasks`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateTasksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateTasksRequest** | [**UpdateTasksRequest**](UpdateTasksRequest.md) |  | 

### Return type

[**map[string]UpdateTaskResponseInfo**](UpdateTaskResponseInfo.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

