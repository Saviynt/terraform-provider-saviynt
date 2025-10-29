# UpdateDynamicAttributesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attributename** | **string** | Specify the dynamic attribute name to be used for filtering out and calling the respective attribute to be displayed. | 
**Requesttype** | Pointer to **string** | Type of request | [optional] 
**Attributetype** | Pointer to **string** | Specify the attribute type based on which you want to filter and display the dynamic attributes. | [optional] 
**Attributegroup** | Pointer to **string** | Attribute Group provides you an option to group or categorize and display the dynamic attributes in the Access Requests request form. | [optional] 
**Orderindex** | Pointer to **string** | Specify the sequence where you want to display the dynamic attributes | [optional] 
**Attributelable** | Pointer to **string** | pecify the name of the attribute, which you want to display in the Access Requests request form. | [optional] 
**Accountscolumn** | Pointer to **string** |  | [optional] 
**Hideoncreate** | Pointer to **string** |  | [optional] 
**Actionstring** | Pointer to **string** |  | [optional] 
**Editable** | Pointer to **string** |  | [optional] 
**Hideonupdate** | Pointer to **string** |  | [optional] 
**Actiontoperformwhenparentattributechanges** | Pointer to **string** |  | [optional] 
**Defaultvalue** | Pointer to **string** |  | [optional] 
**Required** | Pointer to **string** |  | [optional] 
**Regex** | Pointer to **string** |  | [optional] 
**Attributevalue** | Pointer to **string** |  | [optional] 
**Showonchild** | Pointer to **string** |  | [optional] 
**Parentattribute** | Pointer to **string** |  | [optional] 
**Descriptionascsv** | Pointer to **string** |  | [optional] 

## Methods

### NewUpdateDynamicAttributesInner

`func NewUpdateDynamicAttributesInner(attributename string, ) *UpdateDynamicAttributesInner`

NewUpdateDynamicAttributesInner instantiates a new UpdateDynamicAttributesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateDynamicAttributesInnerWithDefaults

`func NewUpdateDynamicAttributesInnerWithDefaults() *UpdateDynamicAttributesInner`

NewUpdateDynamicAttributesInnerWithDefaults instantiates a new UpdateDynamicAttributesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttributename

`func (o *UpdateDynamicAttributesInner) GetAttributename() string`

GetAttributename returns the Attributename field if non-nil, zero value otherwise.

### GetAttributenameOk

`func (o *UpdateDynamicAttributesInner) GetAttributenameOk() (*string, bool)`

GetAttributenameOk returns a tuple with the Attributename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributename

`func (o *UpdateDynamicAttributesInner) SetAttributename(v string)`

SetAttributename sets Attributename field to given value.


### GetRequesttype

`func (o *UpdateDynamicAttributesInner) GetRequesttype() string`

GetRequesttype returns the Requesttype field if non-nil, zero value otherwise.

### GetRequesttypeOk

`func (o *UpdateDynamicAttributesInner) GetRequesttypeOk() (*string, bool)`

GetRequesttypeOk returns a tuple with the Requesttype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequesttype

`func (o *UpdateDynamicAttributesInner) SetRequesttype(v string)`

SetRequesttype sets Requesttype field to given value.

### HasRequesttype

`func (o *UpdateDynamicAttributesInner) HasRequesttype() bool`

HasRequesttype returns a boolean if a field has been set.

### GetAttributetype

`func (o *UpdateDynamicAttributesInner) GetAttributetype() string`

GetAttributetype returns the Attributetype field if non-nil, zero value otherwise.

### GetAttributetypeOk

`func (o *UpdateDynamicAttributesInner) GetAttributetypeOk() (*string, bool)`

GetAttributetypeOk returns a tuple with the Attributetype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributetype

`func (o *UpdateDynamicAttributesInner) SetAttributetype(v string)`

SetAttributetype sets Attributetype field to given value.

### HasAttributetype

`func (o *UpdateDynamicAttributesInner) HasAttributetype() bool`

HasAttributetype returns a boolean if a field has been set.

### GetAttributegroup

`func (o *UpdateDynamicAttributesInner) GetAttributegroup() string`

GetAttributegroup returns the Attributegroup field if non-nil, zero value otherwise.

### GetAttributegroupOk

`func (o *UpdateDynamicAttributesInner) GetAttributegroupOk() (*string, bool)`

GetAttributegroupOk returns a tuple with the Attributegroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributegroup

`func (o *UpdateDynamicAttributesInner) SetAttributegroup(v string)`

SetAttributegroup sets Attributegroup field to given value.

### HasAttributegroup

`func (o *UpdateDynamicAttributesInner) HasAttributegroup() bool`

HasAttributegroup returns a boolean if a field has been set.

### GetOrderindex

`func (o *UpdateDynamicAttributesInner) GetOrderindex() string`

GetOrderindex returns the Orderindex field if non-nil, zero value otherwise.

### GetOrderindexOk

`func (o *UpdateDynamicAttributesInner) GetOrderindexOk() (*string, bool)`

GetOrderindexOk returns a tuple with the Orderindex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderindex

`func (o *UpdateDynamicAttributesInner) SetOrderindex(v string)`

SetOrderindex sets Orderindex field to given value.

### HasOrderindex

`func (o *UpdateDynamicAttributesInner) HasOrderindex() bool`

HasOrderindex returns a boolean if a field has been set.

### GetAttributelable

`func (o *UpdateDynamicAttributesInner) GetAttributelable() string`

GetAttributelable returns the Attributelable field if non-nil, zero value otherwise.

### GetAttributelableOk

`func (o *UpdateDynamicAttributesInner) GetAttributelableOk() (*string, bool)`

GetAttributelableOk returns a tuple with the Attributelable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributelable

`func (o *UpdateDynamicAttributesInner) SetAttributelable(v string)`

SetAttributelable sets Attributelable field to given value.

### HasAttributelable

`func (o *UpdateDynamicAttributesInner) HasAttributelable() bool`

HasAttributelable returns a boolean if a field has been set.

### GetAccountscolumn

`func (o *UpdateDynamicAttributesInner) GetAccountscolumn() string`

GetAccountscolumn returns the Accountscolumn field if non-nil, zero value otherwise.

### GetAccountscolumnOk

`func (o *UpdateDynamicAttributesInner) GetAccountscolumnOk() (*string, bool)`

GetAccountscolumnOk returns a tuple with the Accountscolumn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountscolumn

`func (o *UpdateDynamicAttributesInner) SetAccountscolumn(v string)`

SetAccountscolumn sets Accountscolumn field to given value.

### HasAccountscolumn

`func (o *UpdateDynamicAttributesInner) HasAccountscolumn() bool`

HasAccountscolumn returns a boolean if a field has been set.

### GetHideoncreate

`func (o *UpdateDynamicAttributesInner) GetHideoncreate() string`

GetHideoncreate returns the Hideoncreate field if non-nil, zero value otherwise.

### GetHideoncreateOk

`func (o *UpdateDynamicAttributesInner) GetHideoncreateOk() (*string, bool)`

GetHideoncreateOk returns a tuple with the Hideoncreate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHideoncreate

`func (o *UpdateDynamicAttributesInner) SetHideoncreate(v string)`

SetHideoncreate sets Hideoncreate field to given value.

### HasHideoncreate

`func (o *UpdateDynamicAttributesInner) HasHideoncreate() bool`

HasHideoncreate returns a boolean if a field has been set.

### GetActionstring

`func (o *UpdateDynamicAttributesInner) GetActionstring() string`

GetActionstring returns the Actionstring field if non-nil, zero value otherwise.

### GetActionstringOk

`func (o *UpdateDynamicAttributesInner) GetActionstringOk() (*string, bool)`

GetActionstringOk returns a tuple with the Actionstring field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionstring

`func (o *UpdateDynamicAttributesInner) SetActionstring(v string)`

SetActionstring sets Actionstring field to given value.

### HasActionstring

`func (o *UpdateDynamicAttributesInner) HasActionstring() bool`

HasActionstring returns a boolean if a field has been set.

### GetEditable

`func (o *UpdateDynamicAttributesInner) GetEditable() string`

GetEditable returns the Editable field if non-nil, zero value otherwise.

### GetEditableOk

`func (o *UpdateDynamicAttributesInner) GetEditableOk() (*string, bool)`

GetEditableOk returns a tuple with the Editable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditable

`func (o *UpdateDynamicAttributesInner) SetEditable(v string)`

SetEditable sets Editable field to given value.

### HasEditable

`func (o *UpdateDynamicAttributesInner) HasEditable() bool`

HasEditable returns a boolean if a field has been set.

### GetHideonupdate

`func (o *UpdateDynamicAttributesInner) GetHideonupdate() string`

GetHideonupdate returns the Hideonupdate field if non-nil, zero value otherwise.

### GetHideonupdateOk

`func (o *UpdateDynamicAttributesInner) GetHideonupdateOk() (*string, bool)`

GetHideonupdateOk returns a tuple with the Hideonupdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHideonupdate

`func (o *UpdateDynamicAttributesInner) SetHideonupdate(v string)`

SetHideonupdate sets Hideonupdate field to given value.

### HasHideonupdate

`func (o *UpdateDynamicAttributesInner) HasHideonupdate() bool`

HasHideonupdate returns a boolean if a field has been set.

### GetActiontoperformwhenparentattributechanges

`func (o *UpdateDynamicAttributesInner) GetActiontoperformwhenparentattributechanges() string`

GetActiontoperformwhenparentattributechanges returns the Actiontoperformwhenparentattributechanges field if non-nil, zero value otherwise.

### GetActiontoperformwhenparentattributechangesOk

`func (o *UpdateDynamicAttributesInner) GetActiontoperformwhenparentattributechangesOk() (*string, bool)`

GetActiontoperformwhenparentattributechangesOk returns a tuple with the Actiontoperformwhenparentattributechanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiontoperformwhenparentattributechanges

`func (o *UpdateDynamicAttributesInner) SetActiontoperformwhenparentattributechanges(v string)`

SetActiontoperformwhenparentattributechanges sets Actiontoperformwhenparentattributechanges field to given value.

### HasActiontoperformwhenparentattributechanges

`func (o *UpdateDynamicAttributesInner) HasActiontoperformwhenparentattributechanges() bool`

HasActiontoperformwhenparentattributechanges returns a boolean if a field has been set.

### GetDefaultvalue

`func (o *UpdateDynamicAttributesInner) GetDefaultvalue() string`

GetDefaultvalue returns the Defaultvalue field if non-nil, zero value otherwise.

### GetDefaultvalueOk

`func (o *UpdateDynamicAttributesInner) GetDefaultvalueOk() (*string, bool)`

GetDefaultvalueOk returns a tuple with the Defaultvalue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultvalue

`func (o *UpdateDynamicAttributesInner) SetDefaultvalue(v string)`

SetDefaultvalue sets Defaultvalue field to given value.

### HasDefaultvalue

`func (o *UpdateDynamicAttributesInner) HasDefaultvalue() bool`

HasDefaultvalue returns a boolean if a field has been set.

### GetRequired

`func (o *UpdateDynamicAttributesInner) GetRequired() string`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *UpdateDynamicAttributesInner) GetRequiredOk() (*string, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *UpdateDynamicAttributesInner) SetRequired(v string)`

SetRequired sets Required field to given value.

### HasRequired

`func (o *UpdateDynamicAttributesInner) HasRequired() bool`

HasRequired returns a boolean if a field has been set.

### GetRegex

`func (o *UpdateDynamicAttributesInner) GetRegex() string`

GetRegex returns the Regex field if non-nil, zero value otherwise.

### GetRegexOk

`func (o *UpdateDynamicAttributesInner) GetRegexOk() (*string, bool)`

GetRegexOk returns a tuple with the Regex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegex

`func (o *UpdateDynamicAttributesInner) SetRegex(v string)`

SetRegex sets Regex field to given value.

### HasRegex

`func (o *UpdateDynamicAttributesInner) HasRegex() bool`

HasRegex returns a boolean if a field has been set.

### GetAttributevalue

`func (o *UpdateDynamicAttributesInner) GetAttributevalue() string`

GetAttributevalue returns the Attributevalue field if non-nil, zero value otherwise.

### GetAttributevalueOk

`func (o *UpdateDynamicAttributesInner) GetAttributevalueOk() (*string, bool)`

GetAttributevalueOk returns a tuple with the Attributevalue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributevalue

`func (o *UpdateDynamicAttributesInner) SetAttributevalue(v string)`

SetAttributevalue sets Attributevalue field to given value.

### HasAttributevalue

`func (o *UpdateDynamicAttributesInner) HasAttributevalue() bool`

HasAttributevalue returns a boolean if a field has been set.

### GetShowonchild

`func (o *UpdateDynamicAttributesInner) GetShowonchild() string`

GetShowonchild returns the Showonchild field if non-nil, zero value otherwise.

### GetShowonchildOk

`func (o *UpdateDynamicAttributesInner) GetShowonchildOk() (*string, bool)`

GetShowonchildOk returns a tuple with the Showonchild field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowonchild

`func (o *UpdateDynamicAttributesInner) SetShowonchild(v string)`

SetShowonchild sets Showonchild field to given value.

### HasShowonchild

`func (o *UpdateDynamicAttributesInner) HasShowonchild() bool`

HasShowonchild returns a boolean if a field has been set.

### GetParentattribute

`func (o *UpdateDynamicAttributesInner) GetParentattribute() string`

GetParentattribute returns the Parentattribute field if non-nil, zero value otherwise.

### GetParentattributeOk

`func (o *UpdateDynamicAttributesInner) GetParentattributeOk() (*string, bool)`

GetParentattributeOk returns a tuple with the Parentattribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentattribute

`func (o *UpdateDynamicAttributesInner) SetParentattribute(v string)`

SetParentattribute sets Parentattribute field to given value.

### HasParentattribute

`func (o *UpdateDynamicAttributesInner) HasParentattribute() bool`

HasParentattribute returns a boolean if a field has been set.

### GetDescriptionascsv

`func (o *UpdateDynamicAttributesInner) GetDescriptionascsv() string`

GetDescriptionascsv returns the Descriptionascsv field if non-nil, zero value otherwise.

### GetDescriptionascsvOk

`func (o *UpdateDynamicAttributesInner) GetDescriptionascsvOk() (*string, bool)`

GetDescriptionascsvOk returns a tuple with the Descriptionascsv field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescriptionascsv

`func (o *UpdateDynamicAttributesInner) SetDescriptionascsv(v string)`

SetDescriptionascsv sets Descriptionascsv field to given value.

### HasDescriptionascsv

`func (o *UpdateDynamicAttributesInner) HasDescriptionascsv() bool`

HasDescriptionascsv returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


