# GetPrivilegeDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Hideoncreate** | Pointer to **bool** |  | [optional] 
**ActionString** | Pointer to **string** |  | [optional] 
**Attributegroup** | Pointer to **string** |  | [optional] 
**Childaction** | Pointer to **string** |  | [optional] 
**Orderindex** | Pointer to **string** |  | [optional] 
**Requestablerequired** | Pointer to **bool** |  | [optional] 
**Editable** | Pointer to **bool** |  | [optional] 
**EntitlementsColumn** | Pointer to **string** |  | [optional] 
**Defaultvalue** | Pointer to **string** |  | [optional] 
**Hideonupd** | Pointer to **bool** |  | [optional] 
**AttributeType** | Pointer to **string** |  | [optional] 
**Label** | Pointer to **string** |  | [optional] 
**AttributeConfig** | Pointer to **string** |  | [optional] 
**Descriptionascsv** | Pointer to **string** |  | [optional] 
**Required** | Pointer to **bool** |  | [optional] 
**Regex** | Pointer to **string** |  | [optional] 
**Updatedate** | Pointer to **string** |  | [optional] 
**EntitlementTypes** | Pointer to [**GetPrivilegeDetailEntitlementTypes**](GetPrivilegeDetailEntitlementTypes.md) |  | [optional] 
**AttributeValues** | Pointer to **string** |  | [optional] 
**Parentattribute** | Pointer to **string** |  | [optional] 
**Attribute** | Pointer to **string** |  | [optional] 
**Sqlquery** | Pointer to **string** |  | [optional] 
**Updateuser** | Pointer to **string** |  | [optional] 

## Methods

### NewGetPrivilegeDetail

`func NewGetPrivilegeDetail() *GetPrivilegeDetail`

NewGetPrivilegeDetail instantiates a new GetPrivilegeDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetPrivilegeDetailWithDefaults

`func NewGetPrivilegeDetailWithDefaults() *GetPrivilegeDetail`

NewGetPrivilegeDetailWithDefaults instantiates a new GetPrivilegeDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHideoncreate

`func (o *GetPrivilegeDetail) GetHideoncreate() bool`

GetHideoncreate returns the Hideoncreate field if non-nil, zero value otherwise.

### GetHideoncreateOk

`func (o *GetPrivilegeDetail) GetHideoncreateOk() (*bool, bool)`

GetHideoncreateOk returns a tuple with the Hideoncreate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHideoncreate

`func (o *GetPrivilegeDetail) SetHideoncreate(v bool)`

SetHideoncreate sets Hideoncreate field to given value.

### HasHideoncreate

`func (o *GetPrivilegeDetail) HasHideoncreate() bool`

HasHideoncreate returns a boolean if a field has been set.

### GetActionString

`func (o *GetPrivilegeDetail) GetActionString() string`

GetActionString returns the ActionString field if non-nil, zero value otherwise.

### GetActionStringOk

`func (o *GetPrivilegeDetail) GetActionStringOk() (*string, bool)`

GetActionStringOk returns a tuple with the ActionString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionString

`func (o *GetPrivilegeDetail) SetActionString(v string)`

SetActionString sets ActionString field to given value.

### HasActionString

`func (o *GetPrivilegeDetail) HasActionString() bool`

HasActionString returns a boolean if a field has been set.

### GetAttributegroup

`func (o *GetPrivilegeDetail) GetAttributegroup() string`

GetAttributegroup returns the Attributegroup field if non-nil, zero value otherwise.

### GetAttributegroupOk

`func (o *GetPrivilegeDetail) GetAttributegroupOk() (*string, bool)`

GetAttributegroupOk returns a tuple with the Attributegroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributegroup

`func (o *GetPrivilegeDetail) SetAttributegroup(v string)`

SetAttributegroup sets Attributegroup field to given value.

### HasAttributegroup

`func (o *GetPrivilegeDetail) HasAttributegroup() bool`

HasAttributegroup returns a boolean if a field has been set.

### GetChildaction

`func (o *GetPrivilegeDetail) GetChildaction() string`

GetChildaction returns the Childaction field if non-nil, zero value otherwise.

### GetChildactionOk

`func (o *GetPrivilegeDetail) GetChildactionOk() (*string, bool)`

GetChildactionOk returns a tuple with the Childaction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildaction

`func (o *GetPrivilegeDetail) SetChildaction(v string)`

SetChildaction sets Childaction field to given value.

### HasChildaction

`func (o *GetPrivilegeDetail) HasChildaction() bool`

HasChildaction returns a boolean if a field has been set.

### GetOrderindex

`func (o *GetPrivilegeDetail) GetOrderindex() string`

GetOrderindex returns the Orderindex field if non-nil, zero value otherwise.

### GetOrderindexOk

`func (o *GetPrivilegeDetail) GetOrderindexOk() (*string, bool)`

GetOrderindexOk returns a tuple with the Orderindex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderindex

`func (o *GetPrivilegeDetail) SetOrderindex(v string)`

SetOrderindex sets Orderindex field to given value.

### HasOrderindex

`func (o *GetPrivilegeDetail) HasOrderindex() bool`

HasOrderindex returns a boolean if a field has been set.

### GetRequestablerequired

`func (o *GetPrivilegeDetail) GetRequestablerequired() bool`

GetRequestablerequired returns the Requestablerequired field if non-nil, zero value otherwise.

### GetRequestablerequiredOk

`func (o *GetPrivilegeDetail) GetRequestablerequiredOk() (*bool, bool)`

GetRequestablerequiredOk returns a tuple with the Requestablerequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestablerequired

`func (o *GetPrivilegeDetail) SetRequestablerequired(v bool)`

SetRequestablerequired sets Requestablerequired field to given value.

### HasRequestablerequired

`func (o *GetPrivilegeDetail) HasRequestablerequired() bool`

HasRequestablerequired returns a boolean if a field has been set.

### GetEditable

`func (o *GetPrivilegeDetail) GetEditable() bool`

GetEditable returns the Editable field if non-nil, zero value otherwise.

### GetEditableOk

`func (o *GetPrivilegeDetail) GetEditableOk() (*bool, bool)`

GetEditableOk returns a tuple with the Editable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditable

`func (o *GetPrivilegeDetail) SetEditable(v bool)`

SetEditable sets Editable field to given value.

### HasEditable

`func (o *GetPrivilegeDetail) HasEditable() bool`

HasEditable returns a boolean if a field has been set.

### GetEntitlementsColumn

`func (o *GetPrivilegeDetail) GetEntitlementsColumn() string`

GetEntitlementsColumn returns the EntitlementsColumn field if non-nil, zero value otherwise.

### GetEntitlementsColumnOk

`func (o *GetPrivilegeDetail) GetEntitlementsColumnOk() (*string, bool)`

GetEntitlementsColumnOk returns a tuple with the EntitlementsColumn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitlementsColumn

`func (o *GetPrivilegeDetail) SetEntitlementsColumn(v string)`

SetEntitlementsColumn sets EntitlementsColumn field to given value.

### HasEntitlementsColumn

`func (o *GetPrivilegeDetail) HasEntitlementsColumn() bool`

HasEntitlementsColumn returns a boolean if a field has been set.

### GetDefaultvalue

`func (o *GetPrivilegeDetail) GetDefaultvalue() string`

GetDefaultvalue returns the Defaultvalue field if non-nil, zero value otherwise.

### GetDefaultvalueOk

`func (o *GetPrivilegeDetail) GetDefaultvalueOk() (*string, bool)`

GetDefaultvalueOk returns a tuple with the Defaultvalue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultvalue

`func (o *GetPrivilegeDetail) SetDefaultvalue(v string)`

SetDefaultvalue sets Defaultvalue field to given value.

### HasDefaultvalue

`func (o *GetPrivilegeDetail) HasDefaultvalue() bool`

HasDefaultvalue returns a boolean if a field has been set.

### GetHideonupd

`func (o *GetPrivilegeDetail) GetHideonupd() bool`

GetHideonupd returns the Hideonupd field if non-nil, zero value otherwise.

### GetHideonupdOk

`func (o *GetPrivilegeDetail) GetHideonupdOk() (*bool, bool)`

GetHideonupdOk returns a tuple with the Hideonupd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHideonupd

`func (o *GetPrivilegeDetail) SetHideonupd(v bool)`

SetHideonupd sets Hideonupd field to given value.

### HasHideonupd

`func (o *GetPrivilegeDetail) HasHideonupd() bool`

HasHideonupd returns a boolean if a field has been set.

### GetAttributeType

`func (o *GetPrivilegeDetail) GetAttributeType() string`

GetAttributeType returns the AttributeType field if non-nil, zero value otherwise.

### GetAttributeTypeOk

`func (o *GetPrivilegeDetail) GetAttributeTypeOk() (*string, bool)`

GetAttributeTypeOk returns a tuple with the AttributeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributeType

`func (o *GetPrivilegeDetail) SetAttributeType(v string)`

SetAttributeType sets AttributeType field to given value.

### HasAttributeType

`func (o *GetPrivilegeDetail) HasAttributeType() bool`

HasAttributeType returns a boolean if a field has been set.

### GetLabel

`func (o *GetPrivilegeDetail) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *GetPrivilegeDetail) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *GetPrivilegeDetail) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *GetPrivilegeDetail) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetAttributeConfig

`func (o *GetPrivilegeDetail) GetAttributeConfig() string`

GetAttributeConfig returns the AttributeConfig field if non-nil, zero value otherwise.

### GetAttributeConfigOk

`func (o *GetPrivilegeDetail) GetAttributeConfigOk() (*string, bool)`

GetAttributeConfigOk returns a tuple with the AttributeConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributeConfig

`func (o *GetPrivilegeDetail) SetAttributeConfig(v string)`

SetAttributeConfig sets AttributeConfig field to given value.

### HasAttributeConfig

`func (o *GetPrivilegeDetail) HasAttributeConfig() bool`

HasAttributeConfig returns a boolean if a field has been set.

### GetDescriptionascsv

`func (o *GetPrivilegeDetail) GetDescriptionascsv() string`

GetDescriptionascsv returns the Descriptionascsv field if non-nil, zero value otherwise.

### GetDescriptionascsvOk

`func (o *GetPrivilegeDetail) GetDescriptionascsvOk() (*string, bool)`

GetDescriptionascsvOk returns a tuple with the Descriptionascsv field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescriptionascsv

`func (o *GetPrivilegeDetail) SetDescriptionascsv(v string)`

SetDescriptionascsv sets Descriptionascsv field to given value.

### HasDescriptionascsv

`func (o *GetPrivilegeDetail) HasDescriptionascsv() bool`

HasDescriptionascsv returns a boolean if a field has been set.

### GetRequired

`func (o *GetPrivilegeDetail) GetRequired() bool`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *GetPrivilegeDetail) GetRequiredOk() (*bool, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *GetPrivilegeDetail) SetRequired(v bool)`

SetRequired sets Required field to given value.

### HasRequired

`func (o *GetPrivilegeDetail) HasRequired() bool`

HasRequired returns a boolean if a field has been set.

### GetRegex

`func (o *GetPrivilegeDetail) GetRegex() string`

GetRegex returns the Regex field if non-nil, zero value otherwise.

### GetRegexOk

`func (o *GetPrivilegeDetail) GetRegexOk() (*string, bool)`

GetRegexOk returns a tuple with the Regex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegex

`func (o *GetPrivilegeDetail) SetRegex(v string)`

SetRegex sets Regex field to given value.

### HasRegex

`func (o *GetPrivilegeDetail) HasRegex() bool`

HasRegex returns a boolean if a field has been set.

### GetUpdatedate

`func (o *GetPrivilegeDetail) GetUpdatedate() string`

GetUpdatedate returns the Updatedate field if non-nil, zero value otherwise.

### GetUpdatedateOk

`func (o *GetPrivilegeDetail) GetUpdatedateOk() (*string, bool)`

GetUpdatedateOk returns a tuple with the Updatedate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedate

`func (o *GetPrivilegeDetail) SetUpdatedate(v string)`

SetUpdatedate sets Updatedate field to given value.

### HasUpdatedate

`func (o *GetPrivilegeDetail) HasUpdatedate() bool`

HasUpdatedate returns a boolean if a field has been set.

### GetEntitlementTypes

`func (o *GetPrivilegeDetail) GetEntitlementTypes() GetPrivilegeDetailEntitlementTypes`

GetEntitlementTypes returns the EntitlementTypes field if non-nil, zero value otherwise.

### GetEntitlementTypesOk

`func (o *GetPrivilegeDetail) GetEntitlementTypesOk() (*GetPrivilegeDetailEntitlementTypes, bool)`

GetEntitlementTypesOk returns a tuple with the EntitlementTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitlementTypes

`func (o *GetPrivilegeDetail) SetEntitlementTypes(v GetPrivilegeDetailEntitlementTypes)`

SetEntitlementTypes sets EntitlementTypes field to given value.

### HasEntitlementTypes

`func (o *GetPrivilegeDetail) HasEntitlementTypes() bool`

HasEntitlementTypes returns a boolean if a field has been set.

### GetAttributeValues

`func (o *GetPrivilegeDetail) GetAttributeValues() string`

GetAttributeValues returns the AttributeValues field if non-nil, zero value otherwise.

### GetAttributeValuesOk

`func (o *GetPrivilegeDetail) GetAttributeValuesOk() (*string, bool)`

GetAttributeValuesOk returns a tuple with the AttributeValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributeValues

`func (o *GetPrivilegeDetail) SetAttributeValues(v string)`

SetAttributeValues sets AttributeValues field to given value.

### HasAttributeValues

`func (o *GetPrivilegeDetail) HasAttributeValues() bool`

HasAttributeValues returns a boolean if a field has been set.

### GetParentattribute

`func (o *GetPrivilegeDetail) GetParentattribute() string`

GetParentattribute returns the Parentattribute field if non-nil, zero value otherwise.

### GetParentattributeOk

`func (o *GetPrivilegeDetail) GetParentattributeOk() (*string, bool)`

GetParentattributeOk returns a tuple with the Parentattribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentattribute

`func (o *GetPrivilegeDetail) SetParentattribute(v string)`

SetParentattribute sets Parentattribute field to given value.

### HasParentattribute

`func (o *GetPrivilegeDetail) HasParentattribute() bool`

HasParentattribute returns a boolean if a field has been set.

### GetAttribute

`func (o *GetPrivilegeDetail) GetAttribute() string`

GetAttribute returns the Attribute field if non-nil, zero value otherwise.

### GetAttributeOk

`func (o *GetPrivilegeDetail) GetAttributeOk() (*string, bool)`

GetAttributeOk returns a tuple with the Attribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttribute

`func (o *GetPrivilegeDetail) SetAttribute(v string)`

SetAttribute sets Attribute field to given value.

### HasAttribute

`func (o *GetPrivilegeDetail) HasAttribute() bool`

HasAttribute returns a boolean if a field has been set.

### GetSqlquery

`func (o *GetPrivilegeDetail) GetSqlquery() string`

GetSqlquery returns the Sqlquery field if non-nil, zero value otherwise.

### GetSqlqueryOk

`func (o *GetPrivilegeDetail) GetSqlqueryOk() (*string, bool)`

GetSqlqueryOk returns a tuple with the Sqlquery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSqlquery

`func (o *GetPrivilegeDetail) SetSqlquery(v string)`

SetSqlquery sets Sqlquery field to given value.

### HasSqlquery

`func (o *GetPrivilegeDetail) HasSqlquery() bool`

HasSqlquery returns a boolean if a field has been set.

### GetUpdateuser

`func (o *GetPrivilegeDetail) GetUpdateuser() string`

GetUpdateuser returns the Updateuser field if non-nil, zero value otherwise.

### GetUpdateuserOk

`func (o *GetPrivilegeDetail) GetUpdateuserOk() (*string, bool)`

GetUpdateuserOk returns a tuple with the Updateuser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateuser

`func (o *GetPrivilegeDetail) SetUpdateuser(v string)`

SetUpdateuser sets Updateuser field to given value.

### HasUpdateuser

`func (o *GetPrivilegeDetail) HasUpdateuser() bool`

HasUpdateuser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


