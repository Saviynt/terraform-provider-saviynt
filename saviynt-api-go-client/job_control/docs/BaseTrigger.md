# BaseTrigger

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Triggername** | **string** | Unique name of the trigger | 
**Jobname** | **string** | Name of the job associated with the trigger | 
**Jobgroup** | **string** | Name of the job group associated with the trigger | 
**Triggergroup** | Pointer to **string** | Group classification for the trigger | [optional] 
**Cronexpression** | **string** | Cron expression defining the schedule for the trigger | 

## Methods

### NewBaseTrigger

`func NewBaseTrigger(triggername string, jobname string, jobgroup string, cronexpression string, ) *BaseTrigger`

NewBaseTrigger instantiates a new BaseTrigger object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBaseTriggerWithDefaults

`func NewBaseTriggerWithDefaults() *BaseTrigger`

NewBaseTriggerWithDefaults instantiates a new BaseTrigger object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTriggername

`func (o *BaseTrigger) GetTriggername() string`

GetTriggername returns the Triggername field if non-nil, zero value otherwise.

### GetTriggernameOk

`func (o *BaseTrigger) GetTriggernameOk() (*string, bool)`

GetTriggernameOk returns a tuple with the Triggername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggername

`func (o *BaseTrigger) SetTriggername(v string)`

SetTriggername sets Triggername field to given value.


### GetJobname

`func (o *BaseTrigger) GetJobname() string`

GetJobname returns the Jobname field if non-nil, zero value otherwise.

### GetJobnameOk

`func (o *BaseTrigger) GetJobnameOk() (*string, bool)`

GetJobnameOk returns a tuple with the Jobname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobname

`func (o *BaseTrigger) SetJobname(v string)`

SetJobname sets Jobname field to given value.


### GetJobgroup

`func (o *BaseTrigger) GetJobgroup() string`

GetJobgroup returns the Jobgroup field if non-nil, zero value otherwise.

### GetJobgroupOk

`func (o *BaseTrigger) GetJobgroupOk() (*string, bool)`

GetJobgroupOk returns a tuple with the Jobgroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobgroup

`func (o *BaseTrigger) SetJobgroup(v string)`

SetJobgroup sets Jobgroup field to given value.


### GetTriggergroup

`func (o *BaseTrigger) GetTriggergroup() string`

GetTriggergroup returns the Triggergroup field if non-nil, zero value otherwise.

### GetTriggergroupOk

`func (o *BaseTrigger) GetTriggergroupOk() (*string, bool)`

GetTriggergroupOk returns a tuple with the Triggergroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggergroup

`func (o *BaseTrigger) SetTriggergroup(v string)`

SetTriggergroup sets Triggergroup field to given value.

### HasTriggergroup

`func (o *BaseTrigger) HasTriggergroup() bool`

HasTriggergroup returns a boolean if a field has been set.

### GetCronexpression

`func (o *BaseTrigger) GetCronexpression() string`

GetCronexpression returns the Cronexpression field if non-nil, zero value otherwise.

### GetCronexpressionOk

`func (o *BaseTrigger) GetCronexpressionOk() (*string, bool)`

GetCronexpressionOk returns a tuple with the Cronexpression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCronexpression

`func (o *BaseTrigger) SetCronexpression(v string)`

SetCronexpression sets Cronexpression field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


