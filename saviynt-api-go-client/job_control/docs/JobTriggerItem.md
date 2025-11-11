# JobTriggerItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Unique name of the trigger | 
**JobName** | **string** | Name of the job associated with the trigger | 
**JobGroup** | **string** | Name of the job group associated with the trigger | 
**Group** | **string** | Group classification | 
**CronExp** | **string** | Cron expression defining the schedule for the trigger | 
**ValueMap** | Pointer to [**FileTransferJobAllOfValueMap**](FileTransferJobAllOfValueMap.md) |  | [optional] 

## Methods

### NewJobTriggerItem

`func NewJobTriggerItem(name string, jobName string, jobGroup string, group string, cronExp string, ) *JobTriggerItem`

NewJobTriggerItem instantiates a new JobTriggerItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJobTriggerItemWithDefaults

`func NewJobTriggerItemWithDefaults() *JobTriggerItem`

NewJobTriggerItemWithDefaults instantiates a new JobTriggerItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *JobTriggerItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *JobTriggerItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *JobTriggerItem) SetName(v string)`

SetName sets Name field to given value.


### GetJobName

`func (o *JobTriggerItem) GetJobName() string`

GetJobName returns the JobName field if non-nil, zero value otherwise.

### GetJobNameOk

`func (o *JobTriggerItem) GetJobNameOk() (*string, bool)`

GetJobNameOk returns a tuple with the JobName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobName

`func (o *JobTriggerItem) SetJobName(v string)`

SetJobName sets JobName field to given value.


### GetJobGroup

`func (o *JobTriggerItem) GetJobGroup() string`

GetJobGroup returns the JobGroup field if non-nil, zero value otherwise.

### GetJobGroupOk

`func (o *JobTriggerItem) GetJobGroupOk() (*string, bool)`

GetJobGroupOk returns a tuple with the JobGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobGroup

`func (o *JobTriggerItem) SetJobGroup(v string)`

SetJobGroup sets JobGroup field to given value.


### GetGroup

`func (o *JobTriggerItem) GetGroup() string`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *JobTriggerItem) GetGroupOk() (*string, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *JobTriggerItem) SetGroup(v string)`

SetGroup sets Group field to given value.


### GetCronExp

`func (o *JobTriggerItem) GetCronExp() string`

GetCronExp returns the CronExp field if non-nil, zero value otherwise.

### GetCronExpOk

`func (o *JobTriggerItem) GetCronExpOk() (*string, bool)`

GetCronExpOk returns a tuple with the CronExp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCronExp

`func (o *JobTriggerItem) SetCronExp(v string)`

SetCronExp sets CronExp field to given value.


### GetValueMap

`func (o *JobTriggerItem) GetValueMap() FileTransferJobAllOfValueMap`

GetValueMap returns the ValueMap field if non-nil, zero value otherwise.

### GetValueMapOk

`func (o *JobTriggerItem) GetValueMapOk() (*FileTransferJobAllOfValueMap, bool)`

GetValueMapOk returns a tuple with the ValueMap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueMap

`func (o *JobTriggerItem) SetValueMap(v FileTransferJobAllOfValueMap)`

SetValueMap sets ValueMap field to given value.

### HasValueMap

`func (o *JobTriggerItem) HasValueMap() bool`

HasValueMap returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


