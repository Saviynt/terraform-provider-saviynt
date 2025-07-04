# \JobControlAPI

All URIs are relative to *http://localhost:3000*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CheckJobStatus**](JobControlAPI.md#CheckJobStatus) | **Post** /ECM/api/v5/checkJobStatus | Check Job Status
[**CreateTriggers**](JobControlAPI.md#CreateTriggers) | **Post** /ECM/api/v5/createTriggers | Create Triggers
[**CreateUpdateTrigger**](JobControlAPI.md#CreateUpdateTrigger) | **Post** /ECM/api/v5/createUpdateTrigger | Create and Update Trigger
[**DeleteTrigger**](JobControlAPI.md#DeleteTrigger) | **Post** /ECM/api/v5/deleteTrigger | Delete Trigger
[**FetchJobMetadata**](JobControlAPI.md#FetchJobMetadata) | **Post** /ECM/api/v5/fetchJobMetadata | Fetch Job Metadata
[**ResumePauseJobs**](JobControlAPI.md#ResumePauseJobs) | **Post** /ECM/api/v5/resumePauseJobs | Resume Pause Jobs
[**RunJobTrigger**](JobControlAPI.md#RunJobTrigger) | **Post** /ECM/api/v5/runJobTrigger | Run Job Trigger



## CheckJobStatus

> CheckJobStatusResponse CheckJobStatus(ctx).Jobgroup(jobgroup).Jobname(jobname).Execute()

Check Job Status



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
	jobgroup := "jobgroup_example" // string | 
	jobname := "jobname_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.JobControlAPI.CheckJobStatus(context.Background()).Jobgroup(jobgroup).Jobname(jobname).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `JobControlAPI.CheckJobStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CheckJobStatus`: CheckJobStatusResponse
	fmt.Fprintf(os.Stdout, "Response from `JobControlAPI.CheckJobStatus`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCheckJobStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **jobgroup** | **string** |  | 
 **jobname** | **string** |  | 

### Return type

[**CheckJobStatusResponse**](CheckJobStatusResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateTriggers

> CreateTriggers(ctx).Execute()

Create Triggers



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
	r, err := apiClient.JobControlAPI.CreateTriggers(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `JobControlAPI.CreateTriggers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateTriggersRequest struct via the builder pattern


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


## CreateUpdateTrigger

> CreateUpdateTrigger(ctx).Execute()

Create and Update Trigger



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
	r, err := apiClient.JobControlAPI.CreateUpdateTrigger(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `JobControlAPI.CreateUpdateTrigger``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateUpdateTriggerRequest struct via the builder pattern


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


## DeleteTrigger

> DeleteTrigger(ctx).Execute()

Delete Trigger



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
	r, err := apiClient.JobControlAPI.DeleteTrigger(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `JobControlAPI.DeleteTrigger``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTriggerRequest struct via the builder pattern


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


## FetchJobMetadata

> FetchJobMetadataResponse FetchJobMetadata(ctx).FetchJobMetadataRequest(fetchJobMetadataRequest).Execute()

Fetch Job Metadata



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
	fetchJobMetadataRequest := *openapiclient.NewFetchJobMetadataRequest("Jobname_example") // FetchJobMetadataRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.JobControlAPI.FetchJobMetadata(context.Background()).FetchJobMetadataRequest(fetchJobMetadataRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `JobControlAPI.FetchJobMetadata``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FetchJobMetadata`: FetchJobMetadataResponse
	fmt.Fprintf(os.Stdout, "Response from `JobControlAPI.FetchJobMetadata`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFetchJobMetadataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fetchJobMetadataRequest** | [**FetchJobMetadataRequest**](FetchJobMetadataRequest.md) |  | 

### Return type

[**FetchJobMetadataResponse**](FetchJobMetadataResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ResumePauseJobs

> string ResumePauseJobs(ctx).ResumePauseJobsRequest(resumePauseJobsRequest).Execute()

Resume Pause Jobs



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
	resumePauseJobsRequest := *openapiclient.NewResumePauseJobsRequest("Action_example") // ResumePauseJobsRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.JobControlAPI.ResumePauseJobs(context.Background()).ResumePauseJobsRequest(resumePauseJobsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `JobControlAPI.ResumePauseJobs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ResumePauseJobs`: string
	fmt.Fprintf(os.Stdout, "Response from `JobControlAPI.ResumePauseJobs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiResumePauseJobsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **resumePauseJobsRequest** | [**ResumePauseJobsRequest**](ResumePauseJobsRequest.md) |  | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RunJobTrigger

> RunJobTrigger(ctx).Execute()

Run Job Trigger



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
	r, err := apiClient.JobControlAPI.RunJobTrigger(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `JobControlAPI.RunJobTrigger``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiRunJobTriggerRequest struct via the builder pattern


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

