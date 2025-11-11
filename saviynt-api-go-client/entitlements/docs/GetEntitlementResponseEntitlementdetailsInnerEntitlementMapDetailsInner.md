# GetEntitlementResponseEntitlementdetailsInnerEntitlementMapDetailsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PrimaryEntType** | Pointer to **string** | Type of the primary entitlement. | [optional] 
**RequestFilter** | Pointer to **bool** | Request filter for the mapping. | [optional] 
**AddDependentTask** | Pointer to **bool** | Task to add dependent entitlement. | [optional] 
**RemoveDependentEntTask** | Pointer to **bool** | Task to remove dependent entitlement. | [optional] 
**ExcludeEntitlement** | Pointer to **bool** | Entitlements to exclude. | [optional] 
**Primary** | Pointer to **string** | Primary entitlement identifier. | [optional] 
**PrimaryEntKey** | Pointer to **string** | Key of the primary entitlement. | [optional] 
**ExportPrimary** | Pointer to **string** | Export identifier for primary entitlement. | [optional] 
**Description** | Pointer to **string** | Description of the entitlement mapping. | [optional] 

## Methods

### NewGetEntitlementResponseEntitlementdetailsInnerEntitlementMapDetailsInner

`func NewGetEntitlementResponseEntitlementdetailsInnerEntitlementMapDetailsInner() *GetEntitlementResponseEntitlementdetailsInnerEntitlementMapDetailsInner`

NewGetEntitlementResponseEntitlementdetailsInnerEntitlementMapDetailsInner instantiates a new GetEntitlementResponseEntitlementdetailsInnerEntitlementMapDetailsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetEntitlementResponseEntitlementdetailsInnerEntitlementMapDetailsInnerWithDefaults

`func NewGetEntitlementResponseEntitlementdetailsInnerEntitlementMapDetailsInnerWithDefaults() *GetEntitlementResponseEntitlementdetailsInnerEntitlementMapDetailsInner`

NewGetEntitlementResponseEntitlementdetailsInnerEntitlementMapDetailsInnerWithDefaults instantiates a new GetEntitlementResponseEntitlementdetailsInnerEntitlementMapDetailsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPrimaryEntType

`func (o *GetEntitlementResponseEntitlementdetailsInnerEntitlementMapDetailsInner) GetPrimaryEntType() string`

GetPrimaryEntType returns the PrimaryEntType field if non-nil, zero value otherwise.

### GetPrimaryEntTypeOk

`func (o *GetEntitlementResponseEntitlementdetailsInnerEntitlementMapDetailsInner) GetPrimaryEntTypeOk() (*string, bool)`

GetPrimaryEntTypeOk returns a tuple with the PrimaryEntType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryEntType

`func (o *GetEntitlementResponseEntitlementdetailsInnerEntitlementMapDetailsInner) SetPrimaryEntType(v string)`

SetPrimaryEntType sets PrimaryEntType field to given value.

### HasPrimaryEntType

`func (o *GetEntitlementResponseEntitlementdetailsInnerEntitlementMapDetailsInner) HasPrimaryEntType() bool`

HasPrimaryEntType returns a boolean if a field has been set.

### GetRequestFilter

`func (o *GetEntitlementResponseEntitlementdetailsInnerEntitlementMapDetailsInner) GetRequestFilter() bool`

GetRequestFilter returns the RequestFilter field if non-nil, zero value otherwise.

### GetRequestFilterOk

`func (o *GetEntitlementResponseEntitlementdetailsInnerEntitlementMapDetailsInner) GetRequestFilterOk() (*bool, bool)`

GetRequestFilterOk returns a tuple with the RequestFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestFilter

`func (o *GetEntitlementResponseEntitlementdetailsInnerEntitlementMapDetailsInner) SetRequestFilter(v bool)`

SetRequestFilter sets RequestFilter field to given value.

### HasRequestFilter

`func (o *GetEntitlementResponseEntitlementdetailsInnerEntitlementMapDetailsInner) HasRequestFilter() bool`

HasRequestFilter returns a boolean if a field has been set.

### GetAddDependentTask

`func (o *GetEntitlementResponseEntitlementdetailsInnerEntitlementMapDetailsInner) GetAddDependentTask() bool`

GetAddDependentTask returns the AddDependentTask field if non-nil, zero value otherwise.

### GetAddDependentTaskOk

`func (o *GetEntitlementResponseEntitlementdetailsInnerEntitlementMapDetailsInner) GetAddDependentTaskOk() (*bool, bool)`

GetAddDependentTaskOk returns a tuple with the AddDependentTask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddDependentTask

`func (o *GetEntitlementResponseEntitlementdetailsInnerEntitlementMapDetailsInner) SetAddDependentTask(v bool)`

SetAddDependentTask sets AddDependentTask field to given value.

### HasAddDependentTask

`func (o *GetEntitlementResponseEntitlementdetailsInnerEntitlementMapDetailsInner) HasAddDependentTask() bool`

HasAddDependentTask returns a boolean if a field has been set.

### GetRemoveDependentEntTask

`func (o *GetEntitlementResponseEntitlementdetailsInnerEntitlementMapDetailsInner) GetRemoveDependentEntTask() bool`

GetRemoveDependentEntTask returns the RemoveDependentEntTask field if non-nil, zero value otherwise.

### GetRemoveDependentEntTaskOk

`func (o *GetEntitlementResponseEntitlementdetailsInnerEntitlementMapDetailsInner) GetRemoveDependentEntTaskOk() (*bool, bool)`

GetRemoveDependentEntTaskOk returns a tuple with the RemoveDependentEntTask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoveDependentEntTask

`func (o *GetEntitlementResponseEntitlementdetailsInnerEntitlementMapDetailsInner) SetRemoveDependentEntTask(v bool)`

SetRemoveDependentEntTask sets RemoveDependentEntTask field to given value.

### HasRemoveDependentEntTask

`func (o *GetEntitlementResponseEntitlementdetailsInnerEntitlementMapDetailsInner) HasRemoveDependentEntTask() bool`

HasRemoveDependentEntTask returns a boolean if a field has been set.

### GetExcludeEntitlement

`func (o *GetEntitlementResponseEntitlementdetailsInnerEntitlementMapDetailsInner) GetExcludeEntitlement() bool`

GetExcludeEntitlement returns the ExcludeEntitlement field if non-nil, zero value otherwise.

### GetExcludeEntitlementOk

`func (o *GetEntitlementResponseEntitlementdetailsInnerEntitlementMapDetailsInner) GetExcludeEntitlementOk() (*bool, bool)`

GetExcludeEntitlementOk returns a tuple with the ExcludeEntitlement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeEntitlement

`func (o *GetEntitlementResponseEntitlementdetailsInnerEntitlementMapDetailsInner) SetExcludeEntitlement(v bool)`

SetExcludeEntitlement sets ExcludeEntitlement field to given value.

### HasExcludeEntitlement

`func (o *GetEntitlementResponseEntitlementdetailsInnerEntitlementMapDetailsInner) HasExcludeEntitlement() bool`

HasExcludeEntitlement returns a boolean if a field has been set.

### GetPrimary

`func (o *GetEntitlementResponseEntitlementdetailsInnerEntitlementMapDetailsInner) GetPrimary() string`

GetPrimary returns the Primary field if non-nil, zero value otherwise.

### GetPrimaryOk

`func (o *GetEntitlementResponseEntitlementdetailsInnerEntitlementMapDetailsInner) GetPrimaryOk() (*string, bool)`

GetPrimaryOk returns a tuple with the Primary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimary

`func (o *GetEntitlementResponseEntitlementdetailsInnerEntitlementMapDetailsInner) SetPrimary(v string)`

SetPrimary sets Primary field to given value.

### HasPrimary

`func (o *GetEntitlementResponseEntitlementdetailsInnerEntitlementMapDetailsInner) HasPrimary() bool`

HasPrimary returns a boolean if a field has been set.

### GetPrimaryEntKey

`func (o *GetEntitlementResponseEntitlementdetailsInnerEntitlementMapDetailsInner) GetPrimaryEntKey() string`

GetPrimaryEntKey returns the PrimaryEntKey field if non-nil, zero value otherwise.

### GetPrimaryEntKeyOk

`func (o *GetEntitlementResponseEntitlementdetailsInnerEntitlementMapDetailsInner) GetPrimaryEntKeyOk() (*string, bool)`

GetPrimaryEntKeyOk returns a tuple with the PrimaryEntKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryEntKey

`func (o *GetEntitlementResponseEntitlementdetailsInnerEntitlementMapDetailsInner) SetPrimaryEntKey(v string)`

SetPrimaryEntKey sets PrimaryEntKey field to given value.

### HasPrimaryEntKey

`func (o *GetEntitlementResponseEntitlementdetailsInnerEntitlementMapDetailsInner) HasPrimaryEntKey() bool`

HasPrimaryEntKey returns a boolean if a field has been set.

### GetExportPrimary

`func (o *GetEntitlementResponseEntitlementdetailsInnerEntitlementMapDetailsInner) GetExportPrimary() string`

GetExportPrimary returns the ExportPrimary field if non-nil, zero value otherwise.

### GetExportPrimaryOk

`func (o *GetEntitlementResponseEntitlementdetailsInnerEntitlementMapDetailsInner) GetExportPrimaryOk() (*string, bool)`

GetExportPrimaryOk returns a tuple with the ExportPrimary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExportPrimary

`func (o *GetEntitlementResponseEntitlementdetailsInnerEntitlementMapDetailsInner) SetExportPrimary(v string)`

SetExportPrimary sets ExportPrimary field to given value.

### HasExportPrimary

`func (o *GetEntitlementResponseEntitlementdetailsInnerEntitlementMapDetailsInner) HasExportPrimary() bool`

HasExportPrimary returns a boolean if a field has been set.

### GetDescription

`func (o *GetEntitlementResponseEntitlementdetailsInnerEntitlementMapDetailsInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GetEntitlementResponseEntitlementdetailsInnerEntitlementMapDetailsInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GetEntitlementResponseEntitlementdetailsInnerEntitlementMapDetailsInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GetEntitlementResponseEntitlementdetailsInnerEntitlementMapDetailsInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


