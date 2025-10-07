# UpdateTaskRequestInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Taskid** | **string** |  | 
**Status** | **string** | Use the following numeric status codes to represent the following status. These string status values are returned in the &#x60;checkTaskStatus&#x60; endpoint. 1: New 2: In Progress 3: Complete 4: Discontinued 5: Pending Create 6: Pending Provision 7: Provisioning Failed 8: Error 9: No Action Required | 

## Methods

### NewUpdateTaskRequestInfo

`func NewUpdateTaskRequestInfo(taskid string, status string, ) *UpdateTaskRequestInfo`

NewUpdateTaskRequestInfo instantiates a new UpdateTaskRequestInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateTaskRequestInfoWithDefaults

`func NewUpdateTaskRequestInfoWithDefaults() *UpdateTaskRequestInfo`

NewUpdateTaskRequestInfoWithDefaults instantiates a new UpdateTaskRequestInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTaskid

`func (o *UpdateTaskRequestInfo) GetTaskid() string`

GetTaskid returns the Taskid field if non-nil, zero value otherwise.

### GetTaskidOk

`func (o *UpdateTaskRequestInfo) GetTaskidOk() (*string, bool)`

GetTaskidOk returns a tuple with the Taskid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskid

`func (o *UpdateTaskRequestInfo) SetTaskid(v string)`

SetTaskid sets Taskid field to given value.


### GetStatus

`func (o *UpdateTaskRequestInfo) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *UpdateTaskRequestInfo) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *UpdateTaskRequestInfo) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


