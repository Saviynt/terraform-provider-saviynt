# TriggerItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Triggername** | **string** | Unique name of the trigger | 
**Jobname** | **string** | Name of the job associated with the trigger | 
**Jobgroup** | **string** | Name of the job group associated with the trigger | 
**Triggergroup** | Pointer to **string** | Group classification for the trigger | [optional] 
**Cronexpression** | **string** | Cron expression defining the schedule for the trigger | 
**Securitysystems** | Pointer to **[]string** |  | [optional] 
**Tasktypes** | Pointer to **string** |  | [optional] 
**ValueMap** | Pointer to [**AccountsImportFullJobAllOfValueMap**](AccountsImportFullJobAllOfValueMap.md) |  | [optional] 

## Methods

### NewTriggerItem

`func NewTriggerItem(triggername string, jobname string, jobgroup string, cronexpression string, ) *TriggerItem`

NewTriggerItem instantiates a new TriggerItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTriggerItemWithDefaults

`func NewTriggerItemWithDefaults() *TriggerItem`

NewTriggerItemWithDefaults instantiates a new TriggerItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTriggername

`func (o *TriggerItem) GetTriggername() string`

GetTriggername returns the Triggername field if non-nil, zero value otherwise.

### GetTriggernameOk

`func (o *TriggerItem) GetTriggernameOk() (*string, bool)`

GetTriggernameOk returns a tuple with the Triggername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggername

`func (o *TriggerItem) SetTriggername(v string)`

SetTriggername sets Triggername field to given value.


### GetJobname

`func (o *TriggerItem) GetJobname() string`

GetJobname returns the Jobname field if non-nil, zero value otherwise.

### GetJobnameOk

`func (o *TriggerItem) GetJobnameOk() (*string, bool)`

GetJobnameOk returns a tuple with the Jobname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobname

`func (o *TriggerItem) SetJobname(v string)`

SetJobname sets Jobname field to given value.


### GetJobgroup

`func (o *TriggerItem) GetJobgroup() string`

GetJobgroup returns the Jobgroup field if non-nil, zero value otherwise.

### GetJobgroupOk

`func (o *TriggerItem) GetJobgroupOk() (*string, bool)`

GetJobgroupOk returns a tuple with the Jobgroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobgroup

`func (o *TriggerItem) SetJobgroup(v string)`

SetJobgroup sets Jobgroup field to given value.


### GetTriggergroup

`func (o *TriggerItem) GetTriggergroup() string`

GetTriggergroup returns the Triggergroup field if non-nil, zero value otherwise.

### GetTriggergroupOk

`func (o *TriggerItem) GetTriggergroupOk() (*string, bool)`

GetTriggergroupOk returns a tuple with the Triggergroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggergroup

`func (o *TriggerItem) SetTriggergroup(v string)`

SetTriggergroup sets Triggergroup field to given value.

### HasTriggergroup

`func (o *TriggerItem) HasTriggergroup() bool`

HasTriggergroup returns a boolean if a field has been set.

### GetCronexpression

`func (o *TriggerItem) GetCronexpression() string`

GetCronexpression returns the Cronexpression field if non-nil, zero value otherwise.

### GetCronexpressionOk

`func (o *TriggerItem) GetCronexpressionOk() (*string, bool)`

GetCronexpressionOk returns a tuple with the Cronexpression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCronexpression

`func (o *TriggerItem) SetCronexpression(v string)`

SetCronexpression sets Cronexpression field to given value.


### GetSecuritysystems

`func (o *TriggerItem) GetSecuritysystems() []string`

GetSecuritysystems returns the Securitysystems field if non-nil, zero value otherwise.

### GetSecuritysystemsOk

`func (o *TriggerItem) GetSecuritysystemsOk() (*[]string, bool)`

GetSecuritysystemsOk returns a tuple with the Securitysystems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecuritysystems

`func (o *TriggerItem) SetSecuritysystems(v []string)`

SetSecuritysystems sets Securitysystems field to given value.

### HasSecuritysystems

`func (o *TriggerItem) HasSecuritysystems() bool`

HasSecuritysystems returns a boolean if a field has been set.

### GetTasktypes

`func (o *TriggerItem) GetTasktypes() string`

GetTasktypes returns the Tasktypes field if non-nil, zero value otherwise.

### GetTasktypesOk

`func (o *TriggerItem) GetTasktypesOk() (*string, bool)`

GetTasktypesOk returns a tuple with the Tasktypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTasktypes

`func (o *TriggerItem) SetTasktypes(v string)`

SetTasktypes sets Tasktypes field to given value.

### HasTasktypes

`func (o *TriggerItem) HasTasktypes() bool`

HasTasktypes returns a boolean if a field has been set.

### GetValueMap

`func (o *TriggerItem) GetValueMap() AccountsImportFullJobAllOfValueMap`

GetValueMap returns the ValueMap field if non-nil, zero value otherwise.

### GetValueMapOk

`func (o *TriggerItem) GetValueMapOk() (*AccountsImportFullJobAllOfValueMap, bool)`

GetValueMapOk returns a tuple with the ValueMap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueMap

`func (o *TriggerItem) SetValueMap(v AccountsImportFullJobAllOfValueMap)`

SetValueMap sets ValueMap field to given value.

### HasValueMap

`func (o *TriggerItem) HasValueMap() bool`

HasValueMap returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


