# DeleteTriggerRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Jobname** | **string** |  | 
**Triggername** | Pointer to **string** |  | [optional] 
**Jobgroup** | Pointer to **string** |  | [optional] 

## Methods

### NewDeleteTriggerRequest

`func NewDeleteTriggerRequest(jobname string, ) *DeleteTriggerRequest`

NewDeleteTriggerRequest instantiates a new DeleteTriggerRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeleteTriggerRequestWithDefaults

`func NewDeleteTriggerRequestWithDefaults() *DeleteTriggerRequest`

NewDeleteTriggerRequestWithDefaults instantiates a new DeleteTriggerRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJobname

`func (o *DeleteTriggerRequest) GetJobname() string`

GetJobname returns the Jobname field if non-nil, zero value otherwise.

### GetJobnameOk

`func (o *DeleteTriggerRequest) GetJobnameOk() (*string, bool)`

GetJobnameOk returns a tuple with the Jobname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobname

`func (o *DeleteTriggerRequest) SetJobname(v string)`

SetJobname sets Jobname field to given value.


### GetTriggername

`func (o *DeleteTriggerRequest) GetTriggername() string`

GetTriggername returns the Triggername field if non-nil, zero value otherwise.

### GetTriggernameOk

`func (o *DeleteTriggerRequest) GetTriggernameOk() (*string, bool)`

GetTriggernameOk returns a tuple with the Triggername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggername

`func (o *DeleteTriggerRequest) SetTriggername(v string)`

SetTriggername sets Triggername field to given value.

### HasTriggername

`func (o *DeleteTriggerRequest) HasTriggername() bool`

HasTriggername returns a boolean if a field has been set.

### GetJobgroup

`func (o *DeleteTriggerRequest) GetJobgroup() string`

GetJobgroup returns the Jobgroup field if non-nil, zero value otherwise.

### GetJobgroupOk

`func (o *DeleteTriggerRequest) GetJobgroupOk() (*string, bool)`

GetJobgroupOk returns a tuple with the Jobgroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobgroup

`func (o *DeleteTriggerRequest) SetJobgroup(v string)`

SetJobgroup sets Jobgroup field to given value.

### HasJobgroup

`func (o *DeleteTriggerRequest) HasJobgroup() bool`

HasJobgroup returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


