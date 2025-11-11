# UpdateTasksRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TaskKeyToUpdate** | [**[]UpdateTaskRequestInfo**](UpdateTaskRequestInfo.md) |  | 
**Updateuser** | Pointer to **string** |  | [optional] 

## Methods

### NewUpdateTasksRequest

`func NewUpdateTasksRequest(taskKeyToUpdate []UpdateTaskRequestInfo, ) *UpdateTasksRequest`

NewUpdateTasksRequest instantiates a new UpdateTasksRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateTasksRequestWithDefaults

`func NewUpdateTasksRequestWithDefaults() *UpdateTasksRequest`

NewUpdateTasksRequestWithDefaults instantiates a new UpdateTasksRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTaskKeyToUpdate

`func (o *UpdateTasksRequest) GetTaskKeyToUpdate() []UpdateTaskRequestInfo`

GetTaskKeyToUpdate returns the TaskKeyToUpdate field if non-nil, zero value otherwise.

### GetTaskKeyToUpdateOk

`func (o *UpdateTasksRequest) GetTaskKeyToUpdateOk() (*[]UpdateTaskRequestInfo, bool)`

GetTaskKeyToUpdateOk returns a tuple with the TaskKeyToUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskKeyToUpdate

`func (o *UpdateTasksRequest) SetTaskKeyToUpdate(v []UpdateTaskRequestInfo)`

SetTaskKeyToUpdate sets TaskKeyToUpdate field to given value.


### GetUpdateuser

`func (o *UpdateTasksRequest) GetUpdateuser() string`

GetUpdateuser returns the Updateuser field if non-nil, zero value otherwise.

### GetUpdateuserOk

`func (o *UpdateTasksRequest) GetUpdateuserOk() (*string, bool)`

GetUpdateuserOk returns a tuple with the Updateuser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateuser

`func (o *UpdateTasksRequest) SetUpdateuser(v string)`

SetUpdateuser sets Updateuser field to given value.

### HasUpdateuser

`func (o *UpdateTasksRequest) HasUpdateuser() bool`

HasUpdateuser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


