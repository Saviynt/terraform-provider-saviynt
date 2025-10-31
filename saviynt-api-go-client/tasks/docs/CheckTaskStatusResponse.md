# CheckTaskStatusResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **string** | Can be null or one of the following:  1. &#x60;New&#x60; 2. &#x60;In Progress&#x60; 3. &#x60;Complete&#x60; 4. &#x60;Discontinued&#x60; 5. &#x60;Pending Create&#x60; 6. &#x60;Pending Provision&#x60; 7. &#x60;Provisioning Failed&#x60; 8. &#x60;Error&#x60; 9. &#x60;No Action Required&#x60; | [optional] 
**UpdateDate** | Pointer to **string** |  | [optional] 
**Comments** | Pointer to **string** |  | [optional] 
**ProvisioningComments** | Pointer to **string** |  | [optional] 
**UpdateUser** | Pointer to **string** |  | [optional] 
**ProvisioningMetadata** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewCheckTaskStatusResponse

`func NewCheckTaskStatusResponse() *CheckTaskStatusResponse`

NewCheckTaskStatusResponse instantiates a new CheckTaskStatusResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCheckTaskStatusResponseWithDefaults

`func NewCheckTaskStatusResponseWithDefaults() *CheckTaskStatusResponse`

NewCheckTaskStatusResponseWithDefaults instantiates a new CheckTaskStatusResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *CheckTaskStatusResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CheckTaskStatusResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CheckTaskStatusResponse) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CheckTaskStatusResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUpdateDate

`func (o *CheckTaskStatusResponse) GetUpdateDate() string`

GetUpdateDate returns the UpdateDate field if non-nil, zero value otherwise.

### GetUpdateDateOk

`func (o *CheckTaskStatusResponse) GetUpdateDateOk() (*string, bool)`

GetUpdateDateOk returns a tuple with the UpdateDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateDate

`func (o *CheckTaskStatusResponse) SetUpdateDate(v string)`

SetUpdateDate sets UpdateDate field to given value.

### HasUpdateDate

`func (o *CheckTaskStatusResponse) HasUpdateDate() bool`

HasUpdateDate returns a boolean if a field has been set.

### GetComments

`func (o *CheckTaskStatusResponse) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *CheckTaskStatusResponse) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *CheckTaskStatusResponse) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *CheckTaskStatusResponse) HasComments() bool`

HasComments returns a boolean if a field has been set.

### GetProvisioningComments

`func (o *CheckTaskStatusResponse) GetProvisioningComments() string`

GetProvisioningComments returns the ProvisioningComments field if non-nil, zero value otherwise.

### GetProvisioningCommentsOk

`func (o *CheckTaskStatusResponse) GetProvisioningCommentsOk() (*string, bool)`

GetProvisioningCommentsOk returns a tuple with the ProvisioningComments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisioningComments

`func (o *CheckTaskStatusResponse) SetProvisioningComments(v string)`

SetProvisioningComments sets ProvisioningComments field to given value.

### HasProvisioningComments

`func (o *CheckTaskStatusResponse) HasProvisioningComments() bool`

HasProvisioningComments returns a boolean if a field has been set.

### GetUpdateUser

`func (o *CheckTaskStatusResponse) GetUpdateUser() string`

GetUpdateUser returns the UpdateUser field if non-nil, zero value otherwise.

### GetUpdateUserOk

`func (o *CheckTaskStatusResponse) GetUpdateUserOk() (*string, bool)`

GetUpdateUserOk returns a tuple with the UpdateUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateUser

`func (o *CheckTaskStatusResponse) SetUpdateUser(v string)`

SetUpdateUser sets UpdateUser field to given value.

### HasUpdateUser

`func (o *CheckTaskStatusResponse) HasUpdateUser() bool`

HasUpdateUser returns a boolean if a field has been set.

### GetProvisioningMetadata

`func (o *CheckTaskStatusResponse) GetProvisioningMetadata() map[string]interface{}`

GetProvisioningMetadata returns the ProvisioningMetadata field if non-nil, zero value otherwise.

### GetProvisioningMetadataOk

`func (o *CheckTaskStatusResponse) GetProvisioningMetadataOk() (*map[string]interface{}, bool)`

GetProvisioningMetadataOk returns a tuple with the ProvisioningMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisioningMetadata

`func (o *CheckTaskStatusResponse) SetProvisioningMetadata(v map[string]interface{})`

SetProvisioningMetadata sets ProvisioningMetadata field to given value.

### HasProvisioningMetadata

`func (o *CheckTaskStatusResponse) HasProvisioningMetadata() bool`

HasProvisioningMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


