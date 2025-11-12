# \JobControlAPI

All URIs are relative to *http://localhost:3000*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CheckJobStatus**](JobControlAPI.md#CheckJobStatus) | **Post** /ECM/api/v5/checkJobStatus | This API is used to fetch the status of any job other that Data Import Job.
[**CreateOrUpdateTrigger**](JobControlAPI.md#CreateOrUpdateTrigger) | **Post** /ECM/api/v5/createUpdateTrigger | This API call can be used for create and update a trigger for a particular jobgroup in EIC.
[**CreateTrigger**](JobControlAPI.md#CreateTrigger) | **Post** /ECM/api/v5/createTriggers | This API call can be used for create and update a trigger for a particular jobgroup in EIC.
[**DeleteTrigger**](JobControlAPI.md#DeleteTrigger) | **Post** /ECM/api/v5/deleteTrigger | This API call can be used to delete a trigger for a particular \&quot;jobgroup\&quot; in SSM.
[**FetchJobMetadata**](JobControlAPI.md#FetchJobMetadata) | **Post** /ECM/api/v5/fetchJobMetadata | This API call return job metadata for the last run of a job in SSM.
[**PauseAllJobs**](JobControlAPI.md#PauseAllJobs) | **Put** /ECM/api/v5/jobs/pause-all | Use this API to pause all running jobs.
[**PauseJob**](JobControlAPI.md#PauseJob) | **Put** /ECM/api/v5/jobs/pause | Use this API to pause a selected running job.
[**PauseResumeJobs**](JobControlAPI.md#PauseResumeJobs) | **Post** /ECM/api/v5/resumePauseJobs | The resumePauseJobs API enables you to pause jobs based on their job type and job name.When a job is paused, its status is displayed as Paused on the Job Control Panel page.
[**ResumeAllJobs**](JobControlAPI.md#ResumeAllJobs) | **Put** /ECM/api/v5/jobs/resume-all | Use this API to resume all paused jobs.
[**ResumeJob**](JobControlAPI.md#ResumeJob) | **Put** /ECM/api/v5/jobs/resume | Use this API to resume a selected pause job.
[**RunJobTrigger**](JobControlAPI.md#RunJobTrigger) | **Post** /ECM/api/v5/runJobTrigger | This API call can be used to run a job trigger in SSM.



## CheckJobStatus

> CheckJobStatusResponse CheckJobStatus(ctx).CheckJobStatusRequest(checkJobStatusRequest).Execute()

This API is used to fetch the status of any job other that Data Import Job.

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
	checkJobStatusRequest := *openapiclient.NewCheckJobStatusRequest("EcmJob", "ecmGroup") // CheckJobStatusRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.JobControlAPI.CheckJobStatus(context.Background()).CheckJobStatusRequest(checkJobStatusRequest).Execute()
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
 **checkJobStatusRequest** | [**CheckJobStatusRequest**](CheckJobStatusRequest.md) |  | 

### Return type

[**CheckJobStatusResponse**](CheckJobStatusResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateOrUpdateTrigger

> CreateOrUpdateTriggersResponse CreateOrUpdateTrigger(ctx).CreateOrUpdateTriggersRequest(createOrUpdateTriggersRequest).Execute()

This API call can be used for create and update a trigger for a particular jobgroup in EIC.

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
	createOrUpdateTriggersRequest := *openapiclient.NewCreateOrUpdateTriggersRequest([]openapiclient.TriggerItem{openapiclient.TriggerItem{AccountsImportFullJob: openapiclient.NewAccountsImportFullJob("MyTrigger_001", "WSRetryJob", "utility", "0 0 2 * * ?")}}) // CreateOrUpdateTriggersRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.JobControlAPI.CreateOrUpdateTrigger(context.Background()).CreateOrUpdateTriggersRequest(createOrUpdateTriggersRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `JobControlAPI.CreateOrUpdateTrigger``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateOrUpdateTrigger`: CreateOrUpdateTriggersResponse
	fmt.Fprintf(os.Stdout, "Response from `JobControlAPI.CreateOrUpdateTrigger`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrUpdateTriggerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createOrUpdateTriggersRequest** | [**CreateOrUpdateTriggersRequest**](CreateOrUpdateTriggersRequest.md) |  | 

### Return type

[**CreateOrUpdateTriggersResponse**](CreateOrUpdateTriggersResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateTrigger

> CreateTriggersResponse CreateTrigger(ctx).JobTriggerRequest(jobTriggerRequest).Execute()

This API call can be used for create and update a trigger for a particular jobgroup in EIC.

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
	jobTriggerRequest := []openapiclient.JobTriggerRequest{*openapiclient.NewJobTriggerRequest([]openapiclient.JobTriggerItem{openapiclient.JobTriggerItem{AccountsImportIncrementalJob: openapiclient.NewAccountsImportIncrementalJob("TestSB71", "SchemaUserJob", "Schema", "GRAILS_JOBS", "0 33 14 * * ? 2060")}})} // []JobTriggerRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.JobControlAPI.CreateTrigger(context.Background()).JobTriggerRequest(jobTriggerRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `JobControlAPI.CreateTrigger``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateTrigger`: CreateTriggersResponse
	fmt.Fprintf(os.Stdout, "Response from `JobControlAPI.CreateTrigger`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateTriggerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **jobTriggerRequest** | [**[]JobTriggerRequest**](JobTriggerRequest.md) |  | 

### Return type

[**CreateTriggersResponse**](CreateTriggersResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteTrigger

> DeleteTriggerResponse DeleteTrigger(ctx).DeleteTriggerRequest(deleteTriggerRequest).Execute()

This API call can be used to delete a trigger for a particular \"jobgroup\" in SSM.

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
	deleteTriggerRequest := *openapiclient.NewDeleteTriggerRequest("MyJobGroup", "MyTrigger", "MyJobGroup") // DeleteTriggerRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.JobControlAPI.DeleteTrigger(context.Background()).DeleteTriggerRequest(deleteTriggerRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `JobControlAPI.DeleteTrigger``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteTrigger`: DeleteTriggerResponse
	fmt.Fprintf(os.Stdout, "Response from `JobControlAPI.DeleteTrigger`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTriggerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deleteTriggerRequest** | [**DeleteTriggerRequest**](DeleteTriggerRequest.md) |  | 

### Return type

[**DeleteTriggerResponse**](DeleteTriggerResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchJobMetadata

> FetchJobMetadataResponse FetchJobMetadata(ctx).FetchJobMetadataRequest(fetchJobMetadataRequest).Execute()

This API call return job metadata for the last run of a job in SSM.

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
	fetchJobMetadataRequest := *openapiclient.NewFetchJobMetadataRequest("DataImportJob") // FetchJobMetadataRequest | 

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


## PauseAllJobs

> PauseResumeJobsResponse PauseAllJobs(ctx).Execute()

Use this API to pause all running jobs.

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
	resp, r, err := apiClient.JobControlAPI.PauseAllJobs(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `JobControlAPI.PauseAllJobs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PauseAllJobs`: PauseResumeJobsResponse
	fmt.Fprintf(os.Stdout, "Response from `JobControlAPI.PauseAllJobs`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPauseAllJobsRequest struct via the builder pattern


### Return type

[**PauseResumeJobsResponse**](PauseResumeJobsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PauseJob

> PauseResumeJobsResponse PauseJob(ctx).PauseResumeJobRequest(pauseResumeJobRequest).Execute()

Use this API to pause a selected running job.

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
	pauseResumeJobRequest := *openapiclient.NewPauseResumeJobRequest("Job_Name") // PauseResumeJobRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.JobControlAPI.PauseJob(context.Background()).PauseResumeJobRequest(pauseResumeJobRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `JobControlAPI.PauseJob``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PauseJob`: PauseResumeJobsResponse
	fmt.Fprintf(os.Stdout, "Response from `JobControlAPI.PauseJob`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPauseJobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pauseResumeJobRequest** | [**PauseResumeJobRequest**](PauseResumeJobRequest.md) |  | 

### Return type

[**PauseResumeJobsResponse**](PauseResumeJobsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PauseResumeJobs

> string PauseResumeJobs(ctx).PauseResumeJobsRequest(pauseResumeJobsRequest).Execute()

The resumePauseJobs API enables you to pause jobs based on their job type and job name.When a job is paused, its status is displayed as Paused on the Job Control Panel page.

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
	pauseResumeJobsRequest := *openapiclient.NewPauseResumeJobsRequest("PAUSE") // PauseResumeJobsRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.JobControlAPI.PauseResumeJobs(context.Background()).PauseResumeJobsRequest(pauseResumeJobsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `JobControlAPI.PauseResumeJobs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PauseResumeJobs`: string
	fmt.Fprintf(os.Stdout, "Response from `JobControlAPI.PauseResumeJobs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPauseResumeJobsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pauseResumeJobsRequest** | [**PauseResumeJobsRequest**](PauseResumeJobsRequest.md) |  | 

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


## ResumeAllJobs

> PauseResumeJobsResponse ResumeAllJobs(ctx).Execute()

Use this API to resume all paused jobs.

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
	resp, r, err := apiClient.JobControlAPI.ResumeAllJobs(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `JobControlAPI.ResumeAllJobs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ResumeAllJobs`: PauseResumeJobsResponse
	fmt.Fprintf(os.Stdout, "Response from `JobControlAPI.ResumeAllJobs`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiResumeAllJobsRequest struct via the builder pattern


### Return type

[**PauseResumeJobsResponse**](PauseResumeJobsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ResumeJob

> PauseResumeJobsResponse ResumeJob(ctx).PauseResumeJobRequest(pauseResumeJobRequest).Execute()

Use this API to resume a selected pause job.

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
	pauseResumeJobRequest := *openapiclient.NewPauseResumeJobRequest("Job_Name") // PauseResumeJobRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.JobControlAPI.ResumeJob(context.Background()).PauseResumeJobRequest(pauseResumeJobRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `JobControlAPI.ResumeJob``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ResumeJob`: PauseResumeJobsResponse
	fmt.Fprintf(os.Stdout, "Response from `JobControlAPI.ResumeJob`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiResumeJobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pauseResumeJobRequest** | [**PauseResumeJobRequest**](PauseResumeJobRequest.md) |  | 

### Return type

[**PauseResumeJobsResponse**](PauseResumeJobsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RunJobTrigger

> RunJobTriggerResponse RunJobTrigger(ctx).RunJobTriggerRequest(runJobTriggerRequest).Execute()

This API call can be used to run a job trigger in SSM.

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
	runJobTriggerRequest := *openapiclient.NewRunJobTriggerRequest("DataImportJob", "DataImportTrigger", "DataImportGroup") // RunJobTriggerRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.JobControlAPI.RunJobTrigger(context.Background()).RunJobTriggerRequest(runJobTriggerRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `JobControlAPI.RunJobTrigger``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RunJobTrigger`: RunJobTriggerResponse
	fmt.Fprintf(os.Stdout, "Response from `JobControlAPI.RunJobTrigger`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRunJobTriggerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **runJobTriggerRequest** | [**RunJobTriggerRequest**](RunJobTriggerRequest.md) |  | 

### Return type

[**RunJobTriggerResponse**](RunJobTriggerResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

