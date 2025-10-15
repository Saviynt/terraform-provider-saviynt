# CheckJobStatusRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Triggername** | Pointer to **string** | Name of the trigger to check status for | [optional] 
**Jobname** | **string** | Name of the job to check status for | 
**Jobgroup** | **string** | Group of the job to check status for | 

## Methods

### NewCheckJobStatusRequest

`func NewCheckJobStatusRequest(jobname string, jobgroup string, ) *CheckJobStatusRequest`

NewCheckJobStatusRequest instantiates a new CheckJobStatusRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCheckJobStatusRequestWithDefaults

`func NewCheckJobStatusRequestWithDefaults() *CheckJobStatusRequest`

NewCheckJobStatusRequestWithDefaults instantiates a new CheckJobStatusRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTriggername

`func (o *CheckJobStatusRequest) GetTriggername() string`

GetTriggername returns the Triggername field if non-nil, zero value otherwise.

### GetTriggernameOk

`func (o *CheckJobStatusRequest) GetTriggernameOk() (*string, bool)`

GetTriggernameOk returns a tuple with the Triggername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggername

`func (o *CheckJobStatusRequest) SetTriggername(v string)`

SetTriggername sets Triggername field to given value.

### HasTriggername

`func (o *CheckJobStatusRequest) HasTriggername() bool`

HasTriggername returns a boolean if a field has been set.

### GetJobname

`func (o *CheckJobStatusRequest) GetJobname() string`

GetJobname returns the Jobname field if non-nil, zero value otherwise.

### GetJobnameOk

`func (o *CheckJobStatusRequest) GetJobnameOk() (*string, bool)`

GetJobnameOk returns a tuple with the Jobname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobname

`func (o *CheckJobStatusRequest) SetJobname(v string)`

SetJobname sets Jobname field to given value.


### GetJobgroup

`func (o *CheckJobStatusRequest) GetJobgroup() string`

GetJobgroup returns the Jobgroup field if non-nil, zero value otherwise.

### GetJobgroupOk

`func (o *CheckJobStatusRequest) GetJobgroupOk() (*string, bool)`

GetJobgroupOk returns a tuple with the Jobgroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobgroup

`func (o *CheckJobStatusRequest) SetJobgroup(v string)`

SetJobgroup sets Jobgroup field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


