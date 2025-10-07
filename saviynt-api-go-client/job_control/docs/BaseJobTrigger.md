# BaseJobTrigger

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Unique name of the trigger | 
**JobName** | **string** | Name of the job associated with the trigger | 
**JobGroup** | **string** | Name of the job group associated with the trigger | 
**Group** | **string** | Group classification | 
**CronExp** | **string** | Cron expression defining the schedule for the trigger | 

## Methods

### NewBaseJobTrigger

`func NewBaseJobTrigger(name string, jobName string, jobGroup string, group string, cronExp string, ) *BaseJobTrigger`

NewBaseJobTrigger instantiates a new BaseJobTrigger object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBaseJobTriggerWithDefaults

`func NewBaseJobTriggerWithDefaults() *BaseJobTrigger`

NewBaseJobTriggerWithDefaults instantiates a new BaseJobTrigger object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *BaseJobTrigger) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BaseJobTrigger) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BaseJobTrigger) SetName(v string)`

SetName sets Name field to given value.


### GetJobName

`func (o *BaseJobTrigger) GetJobName() string`

GetJobName returns the JobName field if non-nil, zero value otherwise.

### GetJobNameOk

`func (o *BaseJobTrigger) GetJobNameOk() (*string, bool)`

GetJobNameOk returns a tuple with the JobName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobName

`func (o *BaseJobTrigger) SetJobName(v string)`

SetJobName sets JobName field to given value.


### GetJobGroup

`func (o *BaseJobTrigger) GetJobGroup() string`

GetJobGroup returns the JobGroup field if non-nil, zero value otherwise.

### GetJobGroupOk

`func (o *BaseJobTrigger) GetJobGroupOk() (*string, bool)`

GetJobGroupOk returns a tuple with the JobGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobGroup

`func (o *BaseJobTrigger) SetJobGroup(v string)`

SetJobGroup sets JobGroup field to given value.


### GetGroup

`func (o *BaseJobTrigger) GetGroup() string`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *BaseJobTrigger) GetGroupOk() (*string, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *BaseJobTrigger) SetGroup(v string)`

SetGroup sets Group field to given value.


### GetCronExp

`func (o *BaseJobTrigger) GetCronExp() string`

GetCronExp returns the CronExp field if non-nil, zero value otherwise.

### GetCronExpOk

`func (o *BaseJobTrigger) GetCronExpOk() (*string, bool)`

GetCronExpOk returns a tuple with the CronExp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCronExp

`func (o *BaseJobTrigger) SetCronExp(v string)`

SetCronExp sets CronExp field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


