# FetchJobMetadataRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Jobname** | **string** | Name of the job to fetch metadata for | 
**Triggername** | Pointer to **string** | Name of the trigger to fetch metadata for | [optional] 
**Jobgroup** | Pointer to **string** | Group of the job to fetch metadata for | [optional] 
**Jobstartdate** | Pointer to **string** | Start date for the job run | [optional] 

## Methods

### NewFetchJobMetadataRequest

`func NewFetchJobMetadataRequest(jobname string, ) *FetchJobMetadataRequest`

NewFetchJobMetadataRequest instantiates a new FetchJobMetadataRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFetchJobMetadataRequestWithDefaults

`func NewFetchJobMetadataRequestWithDefaults() *FetchJobMetadataRequest`

NewFetchJobMetadataRequestWithDefaults instantiates a new FetchJobMetadataRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJobname

`func (o *FetchJobMetadataRequest) GetJobname() string`

GetJobname returns the Jobname field if non-nil, zero value otherwise.

### GetJobnameOk

`func (o *FetchJobMetadataRequest) GetJobnameOk() (*string, bool)`

GetJobnameOk returns a tuple with the Jobname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobname

`func (o *FetchJobMetadataRequest) SetJobname(v string)`

SetJobname sets Jobname field to given value.


### GetTriggername

`func (o *FetchJobMetadataRequest) GetTriggername() string`

GetTriggername returns the Triggername field if non-nil, zero value otherwise.

### GetTriggernameOk

`func (o *FetchJobMetadataRequest) GetTriggernameOk() (*string, bool)`

GetTriggernameOk returns a tuple with the Triggername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggername

`func (o *FetchJobMetadataRequest) SetTriggername(v string)`

SetTriggername sets Triggername field to given value.

### HasTriggername

`func (o *FetchJobMetadataRequest) HasTriggername() bool`

HasTriggername returns a boolean if a field has been set.

### GetJobgroup

`func (o *FetchJobMetadataRequest) GetJobgroup() string`

GetJobgroup returns the Jobgroup field if non-nil, zero value otherwise.

### GetJobgroupOk

`func (o *FetchJobMetadataRequest) GetJobgroupOk() (*string, bool)`

GetJobgroupOk returns a tuple with the Jobgroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobgroup

`func (o *FetchJobMetadataRequest) SetJobgroup(v string)`

SetJobgroup sets Jobgroup field to given value.

### HasJobgroup

`func (o *FetchJobMetadataRequest) HasJobgroup() bool`

HasJobgroup returns a boolean if a field has been set.

### GetJobstartdate

`func (o *FetchJobMetadataRequest) GetJobstartdate() string`

GetJobstartdate returns the Jobstartdate field if non-nil, zero value otherwise.

### GetJobstartdateOk

`func (o *FetchJobMetadataRequest) GetJobstartdateOk() (*string, bool)`

GetJobstartdateOk returns a tuple with the Jobstartdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobstartdate

`func (o *FetchJobMetadataRequest) SetJobstartdate(v string)`

SetJobstartdate sets Jobstartdate field to given value.

### HasJobstartdate

`func (o *FetchJobMetadataRequest) HasJobstartdate() bool`

HasJobstartdate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


