# CreateUpdatePrivilegeRequestPrivilege

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attributename** | Pointer to **string** | Attribute name for the privilege | [optional] 
**Attributetype** | Pointer to **string** | Type of the attribute/privilege | [optional] 
**Orderindex** | Pointer to **string** | Order index | [optional] 
**Defaultvalue** | Pointer to **string** | Default value for the privilege. | [optional] 
**Attributeconfig** | Pointer to **string** | Configuration type for the attribute | [optional] 
**Label** | Pointer to **string** | Label for the privilege | [optional] 
**Attributegroup** | Pointer to **string** | Attribute group | [optional] 
**Parentattribute** | Pointer to **string** | Parent attribute for the given privilege | [optional] 
**Childaction** | Pointer to **string** | Child action | [optional] 
**Description** | Pointer to **string** | Description for the privilege | [optional] 
**Required** | Pointer to **bool** | Is required | [optional] 
**Requestable** | Pointer to **bool** | Is requestable | [optional] 
**Hideoncreate** | Pointer to **bool** | Hide on create | [optional] 
**Hideonupd** | Pointer to **bool** | Hide on update | [optional] 
**Actionstring** | Pointer to **string** | Action string | [optional] 

## Methods

### NewCreateUpdatePrivilegeRequestPrivilege

`func NewCreateUpdatePrivilegeRequestPrivilege() *CreateUpdatePrivilegeRequestPrivilege`

NewCreateUpdatePrivilegeRequestPrivilege instantiates a new CreateUpdatePrivilegeRequestPrivilege object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateUpdatePrivilegeRequestPrivilegeWithDefaults

`func NewCreateUpdatePrivilegeRequestPrivilegeWithDefaults() *CreateUpdatePrivilegeRequestPrivilege`

NewCreateUpdatePrivilegeRequestPrivilegeWithDefaults instantiates a new CreateUpdatePrivilegeRequestPrivilege object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttributename

`func (o *CreateUpdatePrivilegeRequestPrivilege) GetAttributename() string`

GetAttributename returns the Attributename field if non-nil, zero value otherwise.

### GetAttributenameOk

`func (o *CreateUpdatePrivilegeRequestPrivilege) GetAttributenameOk() (*string, bool)`

GetAttributenameOk returns a tuple with the Attributename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributename

`func (o *CreateUpdatePrivilegeRequestPrivilege) SetAttributename(v string)`

SetAttributename sets Attributename field to given value.

### HasAttributename

`func (o *CreateUpdatePrivilegeRequestPrivilege) HasAttributename() bool`

HasAttributename returns a boolean if a field has been set.

### GetAttributetype

`func (o *CreateUpdatePrivilegeRequestPrivilege) GetAttributetype() string`

GetAttributetype returns the Attributetype field if non-nil, zero value otherwise.

### GetAttributetypeOk

`func (o *CreateUpdatePrivilegeRequestPrivilege) GetAttributetypeOk() (*string, bool)`

GetAttributetypeOk returns a tuple with the Attributetype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributetype

`func (o *CreateUpdatePrivilegeRequestPrivilege) SetAttributetype(v string)`

SetAttributetype sets Attributetype field to given value.

### HasAttributetype

`func (o *CreateUpdatePrivilegeRequestPrivilege) HasAttributetype() bool`

HasAttributetype returns a boolean if a field has been set.

### GetOrderindex

`func (o *CreateUpdatePrivilegeRequestPrivilege) GetOrderindex() string`

GetOrderindex returns the Orderindex field if non-nil, zero value otherwise.

### GetOrderindexOk

`func (o *CreateUpdatePrivilegeRequestPrivilege) GetOrderindexOk() (*string, bool)`

GetOrderindexOk returns a tuple with the Orderindex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderindex

`func (o *CreateUpdatePrivilegeRequestPrivilege) SetOrderindex(v string)`

SetOrderindex sets Orderindex field to given value.

### HasOrderindex

`func (o *CreateUpdatePrivilegeRequestPrivilege) HasOrderindex() bool`

HasOrderindex returns a boolean if a field has been set.

### GetDefaultvalue

`func (o *CreateUpdatePrivilegeRequestPrivilege) GetDefaultvalue() string`

GetDefaultvalue returns the Defaultvalue field if non-nil, zero value otherwise.

### GetDefaultvalueOk

`func (o *CreateUpdatePrivilegeRequestPrivilege) GetDefaultvalueOk() (*string, bool)`

GetDefaultvalueOk returns a tuple with the Defaultvalue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultvalue

`func (o *CreateUpdatePrivilegeRequestPrivilege) SetDefaultvalue(v string)`

SetDefaultvalue sets Defaultvalue field to given value.

### HasDefaultvalue

`func (o *CreateUpdatePrivilegeRequestPrivilege) HasDefaultvalue() bool`

HasDefaultvalue returns a boolean if a field has been set.

### GetAttributeconfig

`func (o *CreateUpdatePrivilegeRequestPrivilege) GetAttributeconfig() string`

GetAttributeconfig returns the Attributeconfig field if non-nil, zero value otherwise.

### GetAttributeconfigOk

`func (o *CreateUpdatePrivilegeRequestPrivilege) GetAttributeconfigOk() (*string, bool)`

GetAttributeconfigOk returns a tuple with the Attributeconfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributeconfig

`func (o *CreateUpdatePrivilegeRequestPrivilege) SetAttributeconfig(v string)`

SetAttributeconfig sets Attributeconfig field to given value.

### HasAttributeconfig

`func (o *CreateUpdatePrivilegeRequestPrivilege) HasAttributeconfig() bool`

HasAttributeconfig returns a boolean if a field has been set.

### GetLabel

`func (o *CreateUpdatePrivilegeRequestPrivilege) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *CreateUpdatePrivilegeRequestPrivilege) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *CreateUpdatePrivilegeRequestPrivilege) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *CreateUpdatePrivilegeRequestPrivilege) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetAttributegroup

`func (o *CreateUpdatePrivilegeRequestPrivilege) GetAttributegroup() string`

GetAttributegroup returns the Attributegroup field if non-nil, zero value otherwise.

### GetAttributegroupOk

`func (o *CreateUpdatePrivilegeRequestPrivilege) GetAttributegroupOk() (*string, bool)`

GetAttributegroupOk returns a tuple with the Attributegroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributegroup

`func (o *CreateUpdatePrivilegeRequestPrivilege) SetAttributegroup(v string)`

SetAttributegroup sets Attributegroup field to given value.

### HasAttributegroup

`func (o *CreateUpdatePrivilegeRequestPrivilege) HasAttributegroup() bool`

HasAttributegroup returns a boolean if a field has been set.

### GetParentattribute

`func (o *CreateUpdatePrivilegeRequestPrivilege) GetParentattribute() string`

GetParentattribute returns the Parentattribute field if non-nil, zero value otherwise.

### GetParentattributeOk

`func (o *CreateUpdatePrivilegeRequestPrivilege) GetParentattributeOk() (*string, bool)`

GetParentattributeOk returns a tuple with the Parentattribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentattribute

`func (o *CreateUpdatePrivilegeRequestPrivilege) SetParentattribute(v string)`

SetParentattribute sets Parentattribute field to given value.

### HasParentattribute

`func (o *CreateUpdatePrivilegeRequestPrivilege) HasParentattribute() bool`

HasParentattribute returns a boolean if a field has been set.

### GetChildaction

`func (o *CreateUpdatePrivilegeRequestPrivilege) GetChildaction() string`

GetChildaction returns the Childaction field if non-nil, zero value otherwise.

### GetChildactionOk

`func (o *CreateUpdatePrivilegeRequestPrivilege) GetChildactionOk() (*string, bool)`

GetChildactionOk returns a tuple with the Childaction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildaction

`func (o *CreateUpdatePrivilegeRequestPrivilege) SetChildaction(v string)`

SetChildaction sets Childaction field to given value.

### HasChildaction

`func (o *CreateUpdatePrivilegeRequestPrivilege) HasChildaction() bool`

HasChildaction returns a boolean if a field has been set.

### GetDescription

`func (o *CreateUpdatePrivilegeRequestPrivilege) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateUpdatePrivilegeRequestPrivilege) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateUpdatePrivilegeRequestPrivilege) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateUpdatePrivilegeRequestPrivilege) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetRequired

`func (o *CreateUpdatePrivilegeRequestPrivilege) GetRequired() bool`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *CreateUpdatePrivilegeRequestPrivilege) GetRequiredOk() (*bool, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *CreateUpdatePrivilegeRequestPrivilege) SetRequired(v bool)`

SetRequired sets Required field to given value.

### HasRequired

`func (o *CreateUpdatePrivilegeRequestPrivilege) HasRequired() bool`

HasRequired returns a boolean if a field has been set.

### GetRequestable

`func (o *CreateUpdatePrivilegeRequestPrivilege) GetRequestable() bool`

GetRequestable returns the Requestable field if non-nil, zero value otherwise.

### GetRequestableOk

`func (o *CreateUpdatePrivilegeRequestPrivilege) GetRequestableOk() (*bool, bool)`

GetRequestableOk returns a tuple with the Requestable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestable

`func (o *CreateUpdatePrivilegeRequestPrivilege) SetRequestable(v bool)`

SetRequestable sets Requestable field to given value.

### HasRequestable

`func (o *CreateUpdatePrivilegeRequestPrivilege) HasRequestable() bool`

HasRequestable returns a boolean if a field has been set.

### GetHideoncreate

`func (o *CreateUpdatePrivilegeRequestPrivilege) GetHideoncreate() bool`

GetHideoncreate returns the Hideoncreate field if non-nil, zero value otherwise.

### GetHideoncreateOk

`func (o *CreateUpdatePrivilegeRequestPrivilege) GetHideoncreateOk() (*bool, bool)`

GetHideoncreateOk returns a tuple with the Hideoncreate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHideoncreate

`func (o *CreateUpdatePrivilegeRequestPrivilege) SetHideoncreate(v bool)`

SetHideoncreate sets Hideoncreate field to given value.

### HasHideoncreate

`func (o *CreateUpdatePrivilegeRequestPrivilege) HasHideoncreate() bool`

HasHideoncreate returns a boolean if a field has been set.

### GetHideonupd

`func (o *CreateUpdatePrivilegeRequestPrivilege) GetHideonupd() bool`

GetHideonupd returns the Hideonupd field if non-nil, zero value otherwise.

### GetHideonupdOk

`func (o *CreateUpdatePrivilegeRequestPrivilege) GetHideonupdOk() (*bool, bool)`

GetHideonupdOk returns a tuple with the Hideonupd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHideonupd

`func (o *CreateUpdatePrivilegeRequestPrivilege) SetHideonupd(v bool)`

SetHideonupd sets Hideonupd field to given value.

### HasHideonupd

`func (o *CreateUpdatePrivilegeRequestPrivilege) HasHideonupd() bool`

HasHideonupd returns a boolean if a field has been set.

### GetActionstring

`func (o *CreateUpdatePrivilegeRequestPrivilege) GetActionstring() string`

GetActionstring returns the Actionstring field if non-nil, zero value otherwise.

### GetActionstringOk

`func (o *CreateUpdatePrivilegeRequestPrivilege) GetActionstringOk() (*string, bool)`

GetActionstringOk returns a tuple with the Actionstring field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionstring

`func (o *CreateUpdatePrivilegeRequestPrivilege) SetActionstring(v string)`

SetActionstring sets Actionstring field to given value.

### HasActionstring

`func (o *CreateUpdatePrivilegeRequestPrivilege) HasActionstring() bool`

HasActionstring returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


