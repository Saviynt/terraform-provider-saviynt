# \FileDirectoryAPI

All URIs are relative to *http://localhost:3000*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UploadNewFile**](FileDirectoryAPI.md#UploadNewFile) | **Post** /ECM/api/v5/uploadSchemaFile | Upload File



## UploadNewFile

> UploadSchemaFileResponse UploadNewFile(ctx).File(file).PathLocation(pathLocation).Execute()

Upload File



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
	file := os.NewFile(1234, "some_file") // *os.File | the file to upload
	pathLocation := "pathLocation_example" // string | Should be set to `Datafiles` to upload to `job.ecm.imp.file.path` in `InternalConfig.groovy`, or `SAV` to upload to `job.ecm.savfile.path` in `InternalConfig.groovy`. 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FileDirectoryAPI.UploadNewFile(context.Background()).File(file).PathLocation(pathLocation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FileDirectoryAPI.UploadNewFile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UploadNewFile`: UploadSchemaFileResponse
	fmt.Fprintf(os.Stdout, "Response from `FileDirectoryAPI.UploadNewFile`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUploadNewFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **file** | ***os.File** | the file to upload | 
 **pathLocation** | **string** | Should be set to &#x60;Datafiles&#x60; to upload to &#x60;job.ecm.imp.file.path&#x60; in &#x60;InternalConfig.groovy&#x60;, or &#x60;SAV&#x60; to upload to &#x60;job.ecm.savfile.path&#x60; in &#x60;InternalConfig.groovy&#x60;.  | 

### Return type

[**UploadSchemaFileResponse**](UploadSchemaFileResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

