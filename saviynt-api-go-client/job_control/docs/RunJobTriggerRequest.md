# RunJobTriggerRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Jobname** | **string** | Name of the job to run | 
**Triggername** | **string** | Name of the trigger to run | 
**Jobgroup** | **string** | Group of the job to run | 

## Methods

### NewRunJobTriggerRequest

`func NewRunJobTriggerRequest(jobname string, triggername string, jobgroup string, ) *RunJobTriggerRequest`

NewRunJobTriggerRequest instantiates a new RunJobTriggerRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRunJobTriggerRequestWithDefaults

`func NewRunJobTriggerRequestWithDefaults() *RunJobTriggerRequest`

NewRunJobTriggerRequestWithDefaults instantiates a new RunJobTriggerRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJobname

`func (o *RunJobTriggerRequest) GetJobname() string`

GetJobname returns the Jobname field if non-nil, zero value otherwise.

### GetJobnameOk

`func (o *RunJobTriggerRequest) GetJobnameOk() (*string, bool)`

GetJobnameOk returns a tuple with the Jobname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobname

`func (o *RunJobTriggerRequest) SetJobname(v string)`

SetJobname sets Jobname field to given value.


### GetTriggername

`func (o *RunJobTriggerRequest) GetTriggername() string`

GetTriggername returns the Triggername field if non-nil, zero value otherwise.

### GetTriggernameOk

`func (o *RunJobTriggerRequest) GetTriggernameOk() (*string, bool)`

GetTriggernameOk returns a tuple with the Triggername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggername

`func (o *RunJobTriggerRequest) SetTriggername(v string)`

SetTriggername sets Triggername field to given value.


### GetJobgroup

`func (o *RunJobTriggerRequest) GetJobgroup() string`

GetJobgroup returns the Jobgroup field if non-nil, zero value otherwise.

### GetJobgroupOk

`func (o *RunJobTriggerRequest) GetJobgroupOk() (*string, bool)`

GetJobgroupOk returns a tuple with the Jobgroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobgroup

`func (o *RunJobTriggerRequest) SetJobgroup(v string)`

SetJobgroup sets Jobgroup field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


