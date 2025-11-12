# PauseResumeJobsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | **string** | Action to perform on the job (Pause or Resume) | 
**Triggername** | Pointer to **string** | Name of the trigger to pause or resume | [optional] 
**Jobname** | Pointer to **string** | Name of the job to pause or resume | [optional] 

## Methods

### NewPauseResumeJobsRequest

`func NewPauseResumeJobsRequest(action string, ) *PauseResumeJobsRequest`

NewPauseResumeJobsRequest instantiates a new PauseResumeJobsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPauseResumeJobsRequestWithDefaults

`func NewPauseResumeJobsRequestWithDefaults() *PauseResumeJobsRequest`

NewPauseResumeJobsRequestWithDefaults instantiates a new PauseResumeJobsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *PauseResumeJobsRequest) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *PauseResumeJobsRequest) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *PauseResumeJobsRequest) SetAction(v string)`

SetAction sets Action field to given value.


### GetTriggername

`func (o *PauseResumeJobsRequest) GetTriggername() string`

GetTriggername returns the Triggername field if non-nil, zero value otherwise.

### GetTriggernameOk

`func (o *PauseResumeJobsRequest) GetTriggernameOk() (*string, bool)`

GetTriggernameOk returns a tuple with the Triggername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggername

`func (o *PauseResumeJobsRequest) SetTriggername(v string)`

SetTriggername sets Triggername field to given value.

### HasTriggername

`func (o *PauseResumeJobsRequest) HasTriggername() bool`

HasTriggername returns a boolean if a field has been set.

### GetJobname

`func (o *PauseResumeJobsRequest) GetJobname() string`

GetJobname returns the Jobname field if non-nil, zero value otherwise.

### GetJobnameOk

`func (o *PauseResumeJobsRequest) GetJobnameOk() (*string, bool)`

GetJobnameOk returns a tuple with the Jobname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobname

`func (o *PauseResumeJobsRequest) SetJobname(v string)`

SetJobname sets Jobname field to given value.

### HasJobname

`func (o *PauseResumeJobsRequest) HasJobname() bool`

HasJobname returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


