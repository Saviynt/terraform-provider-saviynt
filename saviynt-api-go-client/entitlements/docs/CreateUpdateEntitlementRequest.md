# CreateUpdateEntitlementRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Endpoint** | **string** | Name of the endpoint for the entitlement | 
**Entitlementtype** | **string** | Entitlement type for the entitlement | 
**EntitlementValue** | **string** | Value of the entitlement | 
**Attributes** | Pointer to **string** | Attributes for entitlement values | [optional] 
**EntitlementID** | Pointer to **string** | Unique identifier for entitlement. Used to update the entitlement if multiple entitlements with same entitlement_value are present under the same entitlementtype and endpoint. | [optional] 
**NewentitlementValue** | Pointer to **string** | New value for the entitlement if updating an existing one | [optional] 
**Entitlementcasecheck** | Pointer to **string** | If true, entitlement value search will be case sensitive during create or update. | [optional] 
**EntitlementValuekey** | Pointer to **string** | Key for the entitlement value | [optional] 
**UpdatedentitlementValue** | Pointer to **string** | New value for entitlement_value | [optional] 
**Displayname** | Pointer to **string** | Display name of the entitlement. | [optional] 
**Risk** | Pointer to **int32** | Indicates the risk score or level of the entitlement. | [optional] 
**Status** | Pointer to **int32** | Status of the entitlement (e.g., active/inactive). | [optional] 
**Soxcritical** | Pointer to **int32** | SOX criticality flag. | [optional] 
**Syscritical** | Pointer to **int32** | System criticality flag. | [optional] 
**EntitlementGlossary** | Pointer to **string** | Glossary term or explanation for the entitlement. | [optional] 
**Priviliged** | Pointer to **int32** | Indicates if the entitlement is privileged. | [optional] 
**Module** | Pointer to **string** | Functional module the entitlement belongs to. | [optional] 
**Access** | Pointer to **string** | Access type or permission level (e.g., Read-Only). | [optional] 
**Priority** | Pointer to **int32** | Priority level of the entitlement. | [optional] 
**Description** | Pointer to **string** | Description of the entitlement. | [optional] 
**Confidentiality** | Pointer to **int32** | Confidentiality classification level. | [optional] 
**Customproperty1** | Pointer to **string** | Custom property 1 value. | [optional] 
**Customproperty2** | Pointer to **string** | Custom property 2 value. | [optional] 
**Customproperty3** | Pointer to **string** | Custom property 3 value. | [optional] 
**Customproperty4** | Pointer to **string** | Custom property 4 value. | [optional] 
**Customproperty5** | Pointer to **string** | Custom property 5 value. | [optional] 
**Customproperty6** | Pointer to **string** | Custom property 6 value. | [optional] 
**Customproperty7** | Pointer to **string** | Custom property 7 value. | [optional] 
**Customproperty8** | Pointer to **string** | Custom property 8 value. | [optional] 
**Customproperty9** | Pointer to **string** | Custom property 9 value. | [optional] 
**Customproperty10** | Pointer to **string** | Custom property 10 value. | [optional] 
**Customproperty11** | Pointer to **string** | Custom property 11 value. | [optional] 
**Customproperty12** | Pointer to **string** | Custom property 12 value. | [optional] 
**Customproperty13** | Pointer to **string** | Custom property 13 value. | [optional] 
**Customproperty14** | Pointer to **string** | Custom property 14 value. | [optional] 
**Customproperty15** | Pointer to **string** | Custom property 15 value. | [optional] 
**Customproperty16** | Pointer to **string** | Custom property 16 value. | [optional] 
**Customproperty17** | Pointer to **string** | Custom property 17 value. | [optional] 
**Customproperty18** | Pointer to **string** | Custom property 18 value. | [optional] 
**Customproperty19** | Pointer to **string** | Custom property 19 value. | [optional] 
**Customproperty20** | Pointer to **string** | Custom property 20 value. | [optional] 
**Customproperty21** | Pointer to **string** | Custom property 21 value. | [optional] 
**Customproperty22** | Pointer to **string** | Custom property 22 value. | [optional] 
**Customproperty23** | Pointer to **string** | Custom property 23 value. | [optional] 
**Customproperty24** | Pointer to **string** | Custom property 24 value. | [optional] 
**Customproperty25** | Pointer to **string** | Custom property 25 value. | [optional] 
**Customproperty26** | Pointer to **string** | Custom property 26 value. | [optional] 
**Customproperty27** | Pointer to **string** | Custom property 27 value. | [optional] 
**Customproperty28** | Pointer to **string** | Custom property 28 value. | [optional] 
**Customproperty29** | Pointer to **string** | Custom property 29 value. | [optional] 
**Customproperty30** | Pointer to **string** | Custom property 30 value. | [optional] 
**Customproperty31** | Pointer to **string** | Custom property 31 value. | [optional] 
**Customproperty32** | Pointer to **string** | Custom property 32 value. | [optional] 
**Customproperty33** | Pointer to **string** | Custom property 33 value. | [optional] 
**Customproperty34** | Pointer to **string** | Custom property 34 value. | [optional] 
**Customproperty35** | Pointer to **string** | Custom property 35 value. | [optional] 
**Customproperty36** | Pointer to **string** | Custom property 36 value. | [optional] 
**Customproperty37** | Pointer to **string** | Custom property 37 value. | [optional] 
**Customproperty38** | Pointer to **string** | Custom property 38 value. | [optional] 
**Customproperty39** | Pointer to **string** | Custom property 39 value. | [optional] 
**Customproperty40** | Pointer to **string** | Custom property 40 value. | [optional] 
**Entitlementowner1** | Pointer to **[]string** | Primary entitlement owners. | [optional] 
**Entitlementowner2** | Pointer to **[]string** | Secondary entitlement owners. | [optional] 
**Entitlementowner3** | Pointer to **[]string** |  | [optional] 
**Entitlementowner4** | Pointer to **[]string** |  | [optional] 
**Entitlementowner5** | Pointer to **[]string** |  | [optional] 
**Entitlementowner6** | Pointer to **[]string** |  | [optional] 
**Entitlementowner7** | Pointer to **[]string** |  | [optional] 
**Entitlementowner8** | Pointer to **[]string** |  | [optional] 
**Entitlementowner9** | Pointer to **[]string** |  | [optional] 
**Entitlementowner10** | Pointer to **[]string** |  | [optional] 
**Entitlementowner11** | Pointer to **[]string** |  | [optional] 
**Entitlementowner12** | Pointer to **[]string** |  | [optional] 
**Entitlementowner13** | Pointer to **[]string** |  | [optional] 
**Entitlementowner14** | Pointer to **[]string** |  | [optional] 
**Entitlementowner15** | Pointer to **[]string** |  | [optional] 
**Entitlementowner16** | Pointer to **[]string** |  | [optional] 
**Entitlementowner17** | Pointer to **[]string** |  | [optional] 
**Entitlementowner18** | Pointer to **[]string** |  | [optional] 
**Entitlementowner19** | Pointer to **[]string** |  | [optional] 
**Entitlementowner20** | Pointer to **[]string** |  | [optional] 
**Entitlementowner21** | Pointer to **[]string** |  | [optional] 
**Entitlementowner22** | Pointer to **[]string** |  | [optional] 
**Entitlementowner23** | Pointer to **[]string** |  | [optional] 
**Entitlementowner24** | Pointer to **[]string** |  | [optional] 
**Entitlementowner25** | Pointer to **[]string** |  | [optional] 
**Entitlementowner26** | Pointer to **[]string** |  | [optional] 
**Entitlementowner27** | Pointer to **[]string** |  | [optional] 
**Entitlementmap** | Pointer to [**[]CreateUpdateEntitlementRequestEntitlementmapInner**](CreateUpdateEntitlementRequestEntitlementmapInner.md) |  | [optional] 

## Methods

### NewCreateUpdateEntitlementRequest

`func NewCreateUpdateEntitlementRequest(endpoint string, entitlementtype string, entitlementValue string, ) *CreateUpdateEntitlementRequest`

NewCreateUpdateEntitlementRequest instantiates a new CreateUpdateEntitlementRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateUpdateEntitlementRequestWithDefaults

`func NewCreateUpdateEntitlementRequestWithDefaults() *CreateUpdateEntitlementRequest`

NewCreateUpdateEntitlementRequestWithDefaults instantiates a new CreateUpdateEntitlementRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEndpoint

`func (o *CreateUpdateEntitlementRequest) GetEndpoint() string`

GetEndpoint returns the Endpoint field if non-nil, zero value otherwise.

### GetEndpointOk

`func (o *CreateUpdateEntitlementRequest) GetEndpointOk() (*string, bool)`

GetEndpointOk returns a tuple with the Endpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoint

`func (o *CreateUpdateEntitlementRequest) SetEndpoint(v string)`

SetEndpoint sets Endpoint field to given value.


### GetEntitlementtype

`func (o *CreateUpdateEntitlementRequest) GetEntitlementtype() string`

GetEntitlementtype returns the Entitlementtype field if non-nil, zero value otherwise.

### GetEntitlementtypeOk

`func (o *CreateUpdateEntitlementRequest) GetEntitlementtypeOk() (*string, bool)`

GetEntitlementtypeOk returns a tuple with the Entitlementtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitlementtype

`func (o *CreateUpdateEntitlementRequest) SetEntitlementtype(v string)`

SetEntitlementtype sets Entitlementtype field to given value.


### GetEntitlementValue

`func (o *CreateUpdateEntitlementRequest) GetEntitlementValue() string`

GetEntitlementValue returns the EntitlementValue field if non-nil, zero value otherwise.

### GetEntitlementValueOk

`func (o *CreateUpdateEntitlementRequest) GetEntitlementValueOk() (*string, bool)`

GetEntitlementValueOk returns a tuple with the EntitlementValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitlementValue

`func (o *CreateUpdateEntitlementRequest) SetEntitlementValue(v string)`

SetEntitlementValue sets EntitlementValue field to given value.


### GetAttributes

`func (o *CreateUpdateEntitlementRequest) GetAttributes() string`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *CreateUpdateEntitlementRequest) GetAttributesOk() (*string, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *CreateUpdateEntitlementRequest) SetAttributes(v string)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *CreateUpdateEntitlementRequest) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetEntitlementID

`func (o *CreateUpdateEntitlementRequest) GetEntitlementID() string`

GetEntitlementID returns the EntitlementID field if non-nil, zero value otherwise.

### GetEntitlementIDOk

`func (o *CreateUpdateEntitlementRequest) GetEntitlementIDOk() (*string, bool)`

GetEntitlementIDOk returns a tuple with the EntitlementID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitlementID

`func (o *CreateUpdateEntitlementRequest) SetEntitlementID(v string)`

SetEntitlementID sets EntitlementID field to given value.

### HasEntitlementID

`func (o *CreateUpdateEntitlementRequest) HasEntitlementID() bool`

HasEntitlementID returns a boolean if a field has been set.

### GetNewentitlementValue

`func (o *CreateUpdateEntitlementRequest) GetNewentitlementValue() string`

GetNewentitlementValue returns the NewentitlementValue field if non-nil, zero value otherwise.

### GetNewentitlementValueOk

`func (o *CreateUpdateEntitlementRequest) GetNewentitlementValueOk() (*string, bool)`

GetNewentitlementValueOk returns a tuple with the NewentitlementValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewentitlementValue

`func (o *CreateUpdateEntitlementRequest) SetNewentitlementValue(v string)`

SetNewentitlementValue sets NewentitlementValue field to given value.

### HasNewentitlementValue

`func (o *CreateUpdateEntitlementRequest) HasNewentitlementValue() bool`

HasNewentitlementValue returns a boolean if a field has been set.

### GetEntitlementcasecheck

`func (o *CreateUpdateEntitlementRequest) GetEntitlementcasecheck() string`

GetEntitlementcasecheck returns the Entitlementcasecheck field if non-nil, zero value otherwise.

### GetEntitlementcasecheckOk

`func (o *CreateUpdateEntitlementRequest) GetEntitlementcasecheckOk() (*string, bool)`

GetEntitlementcasecheckOk returns a tuple with the Entitlementcasecheck field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitlementcasecheck

`func (o *CreateUpdateEntitlementRequest) SetEntitlementcasecheck(v string)`

SetEntitlementcasecheck sets Entitlementcasecheck field to given value.

### HasEntitlementcasecheck

`func (o *CreateUpdateEntitlementRequest) HasEntitlementcasecheck() bool`

HasEntitlementcasecheck returns a boolean if a field has been set.

### GetEntitlementValuekey

`func (o *CreateUpdateEntitlementRequest) GetEntitlementValuekey() string`

GetEntitlementValuekey returns the EntitlementValuekey field if non-nil, zero value otherwise.

### GetEntitlementValuekeyOk

`func (o *CreateUpdateEntitlementRequest) GetEntitlementValuekeyOk() (*string, bool)`

GetEntitlementValuekeyOk returns a tuple with the EntitlementValuekey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitlementValuekey

`func (o *CreateUpdateEntitlementRequest) SetEntitlementValuekey(v string)`

SetEntitlementValuekey sets EntitlementValuekey field to given value.

### HasEntitlementValuekey

`func (o *CreateUpdateEntitlementRequest) HasEntitlementValuekey() bool`

HasEntitlementValuekey returns a boolean if a field has been set.

### GetUpdatedentitlementValue

`func (o *CreateUpdateEntitlementRequest) GetUpdatedentitlementValue() string`

GetUpdatedentitlementValue returns the UpdatedentitlementValue field if non-nil, zero value otherwise.

### GetUpdatedentitlementValueOk

`func (o *CreateUpdateEntitlementRequest) GetUpdatedentitlementValueOk() (*string, bool)`

GetUpdatedentitlementValueOk returns a tuple with the UpdatedentitlementValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedentitlementValue

`func (o *CreateUpdateEntitlementRequest) SetUpdatedentitlementValue(v string)`

SetUpdatedentitlementValue sets UpdatedentitlementValue field to given value.

### HasUpdatedentitlementValue

`func (o *CreateUpdateEntitlementRequest) HasUpdatedentitlementValue() bool`

HasUpdatedentitlementValue returns a boolean if a field has been set.

### GetDisplayname

`func (o *CreateUpdateEntitlementRequest) GetDisplayname() string`

GetDisplayname returns the Displayname field if non-nil, zero value otherwise.

### GetDisplaynameOk

`func (o *CreateUpdateEntitlementRequest) GetDisplaynameOk() (*string, bool)`

GetDisplaynameOk returns a tuple with the Displayname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayname

`func (o *CreateUpdateEntitlementRequest) SetDisplayname(v string)`

SetDisplayname sets Displayname field to given value.

### HasDisplayname

`func (o *CreateUpdateEntitlementRequest) HasDisplayname() bool`

HasDisplayname returns a boolean if a field has been set.

### GetRisk

`func (o *CreateUpdateEntitlementRequest) GetRisk() int32`

GetRisk returns the Risk field if non-nil, zero value otherwise.

### GetRiskOk

`func (o *CreateUpdateEntitlementRequest) GetRiskOk() (*int32, bool)`

GetRiskOk returns a tuple with the Risk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRisk

`func (o *CreateUpdateEntitlementRequest) SetRisk(v int32)`

SetRisk sets Risk field to given value.

### HasRisk

`func (o *CreateUpdateEntitlementRequest) HasRisk() bool`

HasRisk returns a boolean if a field has been set.

### GetStatus

`func (o *CreateUpdateEntitlementRequest) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CreateUpdateEntitlementRequest) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CreateUpdateEntitlementRequest) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CreateUpdateEntitlementRequest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSoxcritical

`func (o *CreateUpdateEntitlementRequest) GetSoxcritical() int32`

GetSoxcritical returns the Soxcritical field if non-nil, zero value otherwise.

### GetSoxcriticalOk

`func (o *CreateUpdateEntitlementRequest) GetSoxcriticalOk() (*int32, bool)`

GetSoxcriticalOk returns a tuple with the Soxcritical field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoxcritical

`func (o *CreateUpdateEntitlementRequest) SetSoxcritical(v int32)`

SetSoxcritical sets Soxcritical field to given value.

### HasSoxcritical

`func (o *CreateUpdateEntitlementRequest) HasSoxcritical() bool`

HasSoxcritical returns a boolean if a field has been set.

### GetSyscritical

`func (o *CreateUpdateEntitlementRequest) GetSyscritical() int32`

GetSyscritical returns the Syscritical field if non-nil, zero value otherwise.

### GetSyscriticalOk

`func (o *CreateUpdateEntitlementRequest) GetSyscriticalOk() (*int32, bool)`

GetSyscriticalOk returns a tuple with the Syscritical field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyscritical

`func (o *CreateUpdateEntitlementRequest) SetSyscritical(v int32)`

SetSyscritical sets Syscritical field to given value.

### HasSyscritical

`func (o *CreateUpdateEntitlementRequest) HasSyscritical() bool`

HasSyscritical returns a boolean if a field has been set.

### GetEntitlementGlossary

`func (o *CreateUpdateEntitlementRequest) GetEntitlementGlossary() string`

GetEntitlementGlossary returns the EntitlementGlossary field if non-nil, zero value otherwise.

### GetEntitlementGlossaryOk

`func (o *CreateUpdateEntitlementRequest) GetEntitlementGlossaryOk() (*string, bool)`

GetEntitlementGlossaryOk returns a tuple with the EntitlementGlossary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitlementGlossary

`func (o *CreateUpdateEntitlementRequest) SetEntitlementGlossary(v string)`

SetEntitlementGlossary sets EntitlementGlossary field to given value.

### HasEntitlementGlossary

`func (o *CreateUpdateEntitlementRequest) HasEntitlementGlossary() bool`

HasEntitlementGlossary returns a boolean if a field has been set.

### GetPriviliged

`func (o *CreateUpdateEntitlementRequest) GetPriviliged() int32`

GetPriviliged returns the Priviliged field if non-nil, zero value otherwise.

### GetPriviligedOk

`func (o *CreateUpdateEntitlementRequest) GetPriviligedOk() (*int32, bool)`

GetPriviligedOk returns a tuple with the Priviliged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriviliged

`func (o *CreateUpdateEntitlementRequest) SetPriviliged(v int32)`

SetPriviliged sets Priviliged field to given value.

### HasPriviliged

`func (o *CreateUpdateEntitlementRequest) HasPriviliged() bool`

HasPriviliged returns a boolean if a field has been set.

### GetModule

`func (o *CreateUpdateEntitlementRequest) GetModule() string`

GetModule returns the Module field if non-nil, zero value otherwise.

### GetModuleOk

`func (o *CreateUpdateEntitlementRequest) GetModuleOk() (*string, bool)`

GetModuleOk returns a tuple with the Module field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModule

`func (o *CreateUpdateEntitlementRequest) SetModule(v string)`

SetModule sets Module field to given value.

### HasModule

`func (o *CreateUpdateEntitlementRequest) HasModule() bool`

HasModule returns a boolean if a field has been set.

### GetAccess

`func (o *CreateUpdateEntitlementRequest) GetAccess() string`

GetAccess returns the Access field if non-nil, zero value otherwise.

### GetAccessOk

`func (o *CreateUpdateEntitlementRequest) GetAccessOk() (*string, bool)`

GetAccessOk returns a tuple with the Access field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccess

`func (o *CreateUpdateEntitlementRequest) SetAccess(v string)`

SetAccess sets Access field to given value.

### HasAccess

`func (o *CreateUpdateEntitlementRequest) HasAccess() bool`

HasAccess returns a boolean if a field has been set.

### GetPriority

`func (o *CreateUpdateEntitlementRequest) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *CreateUpdateEntitlementRequest) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *CreateUpdateEntitlementRequest) SetPriority(v int32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *CreateUpdateEntitlementRequest) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetDescription

`func (o *CreateUpdateEntitlementRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateUpdateEntitlementRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateUpdateEntitlementRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateUpdateEntitlementRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetConfidentiality

`func (o *CreateUpdateEntitlementRequest) GetConfidentiality() int32`

GetConfidentiality returns the Confidentiality field if non-nil, zero value otherwise.

### GetConfidentialityOk

`func (o *CreateUpdateEntitlementRequest) GetConfidentialityOk() (*int32, bool)`

GetConfidentialityOk returns a tuple with the Confidentiality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfidentiality

`func (o *CreateUpdateEntitlementRequest) SetConfidentiality(v int32)`

SetConfidentiality sets Confidentiality field to given value.

### HasConfidentiality

`func (o *CreateUpdateEntitlementRequest) HasConfidentiality() bool`

HasConfidentiality returns a boolean if a field has been set.

### GetCustomproperty1

`func (o *CreateUpdateEntitlementRequest) GetCustomproperty1() string`

GetCustomproperty1 returns the Customproperty1 field if non-nil, zero value otherwise.

### GetCustomproperty1Ok

`func (o *CreateUpdateEntitlementRequest) GetCustomproperty1Ok() (*string, bool)`

GetCustomproperty1Ok returns a tuple with the Customproperty1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty1

`func (o *CreateUpdateEntitlementRequest) SetCustomproperty1(v string)`

SetCustomproperty1 sets Customproperty1 field to given value.

### HasCustomproperty1

`func (o *CreateUpdateEntitlementRequest) HasCustomproperty1() bool`

HasCustomproperty1 returns a boolean if a field has been set.

### GetCustomproperty2

`func (o *CreateUpdateEntitlementRequest) GetCustomproperty2() string`

GetCustomproperty2 returns the Customproperty2 field if non-nil, zero value otherwise.

### GetCustomproperty2Ok

`func (o *CreateUpdateEntitlementRequest) GetCustomproperty2Ok() (*string, bool)`

GetCustomproperty2Ok returns a tuple with the Customproperty2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty2

`func (o *CreateUpdateEntitlementRequest) SetCustomproperty2(v string)`

SetCustomproperty2 sets Customproperty2 field to given value.

### HasCustomproperty2

`func (o *CreateUpdateEntitlementRequest) HasCustomproperty2() bool`

HasCustomproperty2 returns a boolean if a field has been set.

### GetCustomproperty3

`func (o *CreateUpdateEntitlementRequest) GetCustomproperty3() string`

GetCustomproperty3 returns the Customproperty3 field if non-nil, zero value otherwise.

### GetCustomproperty3Ok

`func (o *CreateUpdateEntitlementRequest) GetCustomproperty3Ok() (*string, bool)`

GetCustomproperty3Ok returns a tuple with the Customproperty3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty3

`func (o *CreateUpdateEntitlementRequest) SetCustomproperty3(v string)`

SetCustomproperty3 sets Customproperty3 field to given value.

### HasCustomproperty3

`func (o *CreateUpdateEntitlementRequest) HasCustomproperty3() bool`

HasCustomproperty3 returns a boolean if a field has been set.

### GetCustomproperty4

`func (o *CreateUpdateEntitlementRequest) GetCustomproperty4() string`

GetCustomproperty4 returns the Customproperty4 field if non-nil, zero value otherwise.

### GetCustomproperty4Ok

`func (o *CreateUpdateEntitlementRequest) GetCustomproperty4Ok() (*string, bool)`

GetCustomproperty4Ok returns a tuple with the Customproperty4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty4

`func (o *CreateUpdateEntitlementRequest) SetCustomproperty4(v string)`

SetCustomproperty4 sets Customproperty4 field to given value.

### HasCustomproperty4

`func (o *CreateUpdateEntitlementRequest) HasCustomproperty4() bool`

HasCustomproperty4 returns a boolean if a field has been set.

### GetCustomproperty5

`func (o *CreateUpdateEntitlementRequest) GetCustomproperty5() string`

GetCustomproperty5 returns the Customproperty5 field if non-nil, zero value otherwise.

### GetCustomproperty5Ok

`func (o *CreateUpdateEntitlementRequest) GetCustomproperty5Ok() (*string, bool)`

GetCustomproperty5Ok returns a tuple with the Customproperty5 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty5

`func (o *CreateUpdateEntitlementRequest) SetCustomproperty5(v string)`

SetCustomproperty5 sets Customproperty5 field to given value.

### HasCustomproperty5

`func (o *CreateUpdateEntitlementRequest) HasCustomproperty5() bool`

HasCustomproperty5 returns a boolean if a field has been set.

### GetCustomproperty6

`func (o *CreateUpdateEntitlementRequest) GetCustomproperty6() string`

GetCustomproperty6 returns the Customproperty6 field if non-nil, zero value otherwise.

### GetCustomproperty6Ok

`func (o *CreateUpdateEntitlementRequest) GetCustomproperty6Ok() (*string, bool)`

GetCustomproperty6Ok returns a tuple with the Customproperty6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty6

`func (o *CreateUpdateEntitlementRequest) SetCustomproperty6(v string)`

SetCustomproperty6 sets Customproperty6 field to given value.

### HasCustomproperty6

`func (o *CreateUpdateEntitlementRequest) HasCustomproperty6() bool`

HasCustomproperty6 returns a boolean if a field has been set.

### GetCustomproperty7

`func (o *CreateUpdateEntitlementRequest) GetCustomproperty7() string`

GetCustomproperty7 returns the Customproperty7 field if non-nil, zero value otherwise.

### GetCustomproperty7Ok

`func (o *CreateUpdateEntitlementRequest) GetCustomproperty7Ok() (*string, bool)`

GetCustomproperty7Ok returns a tuple with the Customproperty7 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty7

`func (o *CreateUpdateEntitlementRequest) SetCustomproperty7(v string)`

SetCustomproperty7 sets Customproperty7 field to given value.

### HasCustomproperty7

`func (o *CreateUpdateEntitlementRequest) HasCustomproperty7() bool`

HasCustomproperty7 returns a boolean if a field has been set.

### GetCustomproperty8

`func (o *CreateUpdateEntitlementRequest) GetCustomproperty8() string`

GetCustomproperty8 returns the Customproperty8 field if non-nil, zero value otherwise.

### GetCustomproperty8Ok

`func (o *CreateUpdateEntitlementRequest) GetCustomproperty8Ok() (*string, bool)`

GetCustomproperty8Ok returns a tuple with the Customproperty8 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty8

`func (o *CreateUpdateEntitlementRequest) SetCustomproperty8(v string)`

SetCustomproperty8 sets Customproperty8 field to given value.

### HasCustomproperty8

`func (o *CreateUpdateEntitlementRequest) HasCustomproperty8() bool`

HasCustomproperty8 returns a boolean if a field has been set.

### GetCustomproperty9

`func (o *CreateUpdateEntitlementRequest) GetCustomproperty9() string`

GetCustomproperty9 returns the Customproperty9 field if non-nil, zero value otherwise.

### GetCustomproperty9Ok

`func (o *CreateUpdateEntitlementRequest) GetCustomproperty9Ok() (*string, bool)`

GetCustomproperty9Ok returns a tuple with the Customproperty9 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty9

`func (o *CreateUpdateEntitlementRequest) SetCustomproperty9(v string)`

SetCustomproperty9 sets Customproperty9 field to given value.

### HasCustomproperty9

`func (o *CreateUpdateEntitlementRequest) HasCustomproperty9() bool`

HasCustomproperty9 returns a boolean if a field has been set.

### GetCustomproperty10

`func (o *CreateUpdateEntitlementRequest) GetCustomproperty10() string`

GetCustomproperty10 returns the Customproperty10 field if non-nil, zero value otherwise.

### GetCustomproperty10Ok

`func (o *CreateUpdateEntitlementRequest) GetCustomproperty10Ok() (*string, bool)`

GetCustomproperty10Ok returns a tuple with the Customproperty10 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty10

`func (o *CreateUpdateEntitlementRequest) SetCustomproperty10(v string)`

SetCustomproperty10 sets Customproperty10 field to given value.

### HasCustomproperty10

`func (o *CreateUpdateEntitlementRequest) HasCustomproperty10() bool`

HasCustomproperty10 returns a boolean if a field has been set.

### GetCustomproperty11

`func (o *CreateUpdateEntitlementRequest) GetCustomproperty11() string`

GetCustomproperty11 returns the Customproperty11 field if non-nil, zero value otherwise.

### GetCustomproperty11Ok

`func (o *CreateUpdateEntitlementRequest) GetCustomproperty11Ok() (*string, bool)`

GetCustomproperty11Ok returns a tuple with the Customproperty11 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty11

`func (o *CreateUpdateEntitlementRequest) SetCustomproperty11(v string)`

SetCustomproperty11 sets Customproperty11 field to given value.

### HasCustomproperty11

`func (o *CreateUpdateEntitlementRequest) HasCustomproperty11() bool`

HasCustomproperty11 returns a boolean if a field has been set.

### GetCustomproperty12

`func (o *CreateUpdateEntitlementRequest) GetCustomproperty12() string`

GetCustomproperty12 returns the Customproperty12 field if non-nil, zero value otherwise.

### GetCustomproperty12Ok

`func (o *CreateUpdateEntitlementRequest) GetCustomproperty12Ok() (*string, bool)`

GetCustomproperty12Ok returns a tuple with the Customproperty12 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty12

`func (o *CreateUpdateEntitlementRequest) SetCustomproperty12(v string)`

SetCustomproperty12 sets Customproperty12 field to given value.

### HasCustomproperty12

`func (o *CreateUpdateEntitlementRequest) HasCustomproperty12() bool`

HasCustomproperty12 returns a boolean if a field has been set.

### GetCustomproperty13

`func (o *CreateUpdateEntitlementRequest) GetCustomproperty13() string`

GetCustomproperty13 returns the Customproperty13 field if non-nil, zero value otherwise.

### GetCustomproperty13Ok

`func (o *CreateUpdateEntitlementRequest) GetCustomproperty13Ok() (*string, bool)`

GetCustomproperty13Ok returns a tuple with the Customproperty13 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty13

`func (o *CreateUpdateEntitlementRequest) SetCustomproperty13(v string)`

SetCustomproperty13 sets Customproperty13 field to given value.

### HasCustomproperty13

`func (o *CreateUpdateEntitlementRequest) HasCustomproperty13() bool`

HasCustomproperty13 returns a boolean if a field has been set.

### GetCustomproperty14

`func (o *CreateUpdateEntitlementRequest) GetCustomproperty14() string`

GetCustomproperty14 returns the Customproperty14 field if non-nil, zero value otherwise.

### GetCustomproperty14Ok

`func (o *CreateUpdateEntitlementRequest) GetCustomproperty14Ok() (*string, bool)`

GetCustomproperty14Ok returns a tuple with the Customproperty14 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty14

`func (o *CreateUpdateEntitlementRequest) SetCustomproperty14(v string)`

SetCustomproperty14 sets Customproperty14 field to given value.

### HasCustomproperty14

`func (o *CreateUpdateEntitlementRequest) HasCustomproperty14() bool`

HasCustomproperty14 returns a boolean if a field has been set.

### GetCustomproperty15

`func (o *CreateUpdateEntitlementRequest) GetCustomproperty15() string`

GetCustomproperty15 returns the Customproperty15 field if non-nil, zero value otherwise.

### GetCustomproperty15Ok

`func (o *CreateUpdateEntitlementRequest) GetCustomproperty15Ok() (*string, bool)`

GetCustomproperty15Ok returns a tuple with the Customproperty15 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty15

`func (o *CreateUpdateEntitlementRequest) SetCustomproperty15(v string)`

SetCustomproperty15 sets Customproperty15 field to given value.

### HasCustomproperty15

`func (o *CreateUpdateEntitlementRequest) HasCustomproperty15() bool`

HasCustomproperty15 returns a boolean if a field has been set.

### GetCustomproperty16

`func (o *CreateUpdateEntitlementRequest) GetCustomproperty16() string`

GetCustomproperty16 returns the Customproperty16 field if non-nil, zero value otherwise.

### GetCustomproperty16Ok

`func (o *CreateUpdateEntitlementRequest) GetCustomproperty16Ok() (*string, bool)`

GetCustomproperty16Ok returns a tuple with the Customproperty16 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty16

`func (o *CreateUpdateEntitlementRequest) SetCustomproperty16(v string)`

SetCustomproperty16 sets Customproperty16 field to given value.

### HasCustomproperty16

`func (o *CreateUpdateEntitlementRequest) HasCustomproperty16() bool`

HasCustomproperty16 returns a boolean if a field has been set.

### GetCustomproperty17

`func (o *CreateUpdateEntitlementRequest) GetCustomproperty17() string`

GetCustomproperty17 returns the Customproperty17 field if non-nil, zero value otherwise.

### GetCustomproperty17Ok

`func (o *CreateUpdateEntitlementRequest) GetCustomproperty17Ok() (*string, bool)`

GetCustomproperty17Ok returns a tuple with the Customproperty17 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty17

`func (o *CreateUpdateEntitlementRequest) SetCustomproperty17(v string)`

SetCustomproperty17 sets Customproperty17 field to given value.

### HasCustomproperty17

`func (o *CreateUpdateEntitlementRequest) HasCustomproperty17() bool`

HasCustomproperty17 returns a boolean if a field has been set.

### GetCustomproperty18

`func (o *CreateUpdateEntitlementRequest) GetCustomproperty18() string`

GetCustomproperty18 returns the Customproperty18 field if non-nil, zero value otherwise.

### GetCustomproperty18Ok

`func (o *CreateUpdateEntitlementRequest) GetCustomproperty18Ok() (*string, bool)`

GetCustomproperty18Ok returns a tuple with the Customproperty18 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty18

`func (o *CreateUpdateEntitlementRequest) SetCustomproperty18(v string)`

SetCustomproperty18 sets Customproperty18 field to given value.

### HasCustomproperty18

`func (o *CreateUpdateEntitlementRequest) HasCustomproperty18() bool`

HasCustomproperty18 returns a boolean if a field has been set.

### GetCustomproperty19

`func (o *CreateUpdateEntitlementRequest) GetCustomproperty19() string`

GetCustomproperty19 returns the Customproperty19 field if non-nil, zero value otherwise.

### GetCustomproperty19Ok

`func (o *CreateUpdateEntitlementRequest) GetCustomproperty19Ok() (*string, bool)`

GetCustomproperty19Ok returns a tuple with the Customproperty19 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty19

`func (o *CreateUpdateEntitlementRequest) SetCustomproperty19(v string)`

SetCustomproperty19 sets Customproperty19 field to given value.

### HasCustomproperty19

`func (o *CreateUpdateEntitlementRequest) HasCustomproperty19() bool`

HasCustomproperty19 returns a boolean if a field has been set.

### GetCustomproperty20

`func (o *CreateUpdateEntitlementRequest) GetCustomproperty20() string`

GetCustomproperty20 returns the Customproperty20 field if non-nil, zero value otherwise.

### GetCustomproperty20Ok

`func (o *CreateUpdateEntitlementRequest) GetCustomproperty20Ok() (*string, bool)`

GetCustomproperty20Ok returns a tuple with the Customproperty20 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty20

`func (o *CreateUpdateEntitlementRequest) SetCustomproperty20(v string)`

SetCustomproperty20 sets Customproperty20 field to given value.

### HasCustomproperty20

`func (o *CreateUpdateEntitlementRequest) HasCustomproperty20() bool`

HasCustomproperty20 returns a boolean if a field has been set.

### GetCustomproperty21

`func (o *CreateUpdateEntitlementRequest) GetCustomproperty21() string`

GetCustomproperty21 returns the Customproperty21 field if non-nil, zero value otherwise.

### GetCustomproperty21Ok

`func (o *CreateUpdateEntitlementRequest) GetCustomproperty21Ok() (*string, bool)`

GetCustomproperty21Ok returns a tuple with the Customproperty21 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty21

`func (o *CreateUpdateEntitlementRequest) SetCustomproperty21(v string)`

SetCustomproperty21 sets Customproperty21 field to given value.

### HasCustomproperty21

`func (o *CreateUpdateEntitlementRequest) HasCustomproperty21() bool`

HasCustomproperty21 returns a boolean if a field has been set.

### GetCustomproperty22

`func (o *CreateUpdateEntitlementRequest) GetCustomproperty22() string`

GetCustomproperty22 returns the Customproperty22 field if non-nil, zero value otherwise.

### GetCustomproperty22Ok

`func (o *CreateUpdateEntitlementRequest) GetCustomproperty22Ok() (*string, bool)`

GetCustomproperty22Ok returns a tuple with the Customproperty22 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty22

`func (o *CreateUpdateEntitlementRequest) SetCustomproperty22(v string)`

SetCustomproperty22 sets Customproperty22 field to given value.

### HasCustomproperty22

`func (o *CreateUpdateEntitlementRequest) HasCustomproperty22() bool`

HasCustomproperty22 returns a boolean if a field has been set.

### GetCustomproperty23

`func (o *CreateUpdateEntitlementRequest) GetCustomproperty23() string`

GetCustomproperty23 returns the Customproperty23 field if non-nil, zero value otherwise.

### GetCustomproperty23Ok

`func (o *CreateUpdateEntitlementRequest) GetCustomproperty23Ok() (*string, bool)`

GetCustomproperty23Ok returns a tuple with the Customproperty23 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty23

`func (o *CreateUpdateEntitlementRequest) SetCustomproperty23(v string)`

SetCustomproperty23 sets Customproperty23 field to given value.

### HasCustomproperty23

`func (o *CreateUpdateEntitlementRequest) HasCustomproperty23() bool`

HasCustomproperty23 returns a boolean if a field has been set.

### GetCustomproperty24

`func (o *CreateUpdateEntitlementRequest) GetCustomproperty24() string`

GetCustomproperty24 returns the Customproperty24 field if non-nil, zero value otherwise.

### GetCustomproperty24Ok

`func (o *CreateUpdateEntitlementRequest) GetCustomproperty24Ok() (*string, bool)`

GetCustomproperty24Ok returns a tuple with the Customproperty24 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty24

`func (o *CreateUpdateEntitlementRequest) SetCustomproperty24(v string)`

SetCustomproperty24 sets Customproperty24 field to given value.

### HasCustomproperty24

`func (o *CreateUpdateEntitlementRequest) HasCustomproperty24() bool`

HasCustomproperty24 returns a boolean if a field has been set.

### GetCustomproperty25

`func (o *CreateUpdateEntitlementRequest) GetCustomproperty25() string`

GetCustomproperty25 returns the Customproperty25 field if non-nil, zero value otherwise.

### GetCustomproperty25Ok

`func (o *CreateUpdateEntitlementRequest) GetCustomproperty25Ok() (*string, bool)`

GetCustomproperty25Ok returns a tuple with the Customproperty25 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty25

`func (o *CreateUpdateEntitlementRequest) SetCustomproperty25(v string)`

SetCustomproperty25 sets Customproperty25 field to given value.

### HasCustomproperty25

`func (o *CreateUpdateEntitlementRequest) HasCustomproperty25() bool`

HasCustomproperty25 returns a boolean if a field has been set.

### GetCustomproperty26

`func (o *CreateUpdateEntitlementRequest) GetCustomproperty26() string`

GetCustomproperty26 returns the Customproperty26 field if non-nil, zero value otherwise.

### GetCustomproperty26Ok

`func (o *CreateUpdateEntitlementRequest) GetCustomproperty26Ok() (*string, bool)`

GetCustomproperty26Ok returns a tuple with the Customproperty26 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty26

`func (o *CreateUpdateEntitlementRequest) SetCustomproperty26(v string)`

SetCustomproperty26 sets Customproperty26 field to given value.

### HasCustomproperty26

`func (o *CreateUpdateEntitlementRequest) HasCustomproperty26() bool`

HasCustomproperty26 returns a boolean if a field has been set.

### GetCustomproperty27

`func (o *CreateUpdateEntitlementRequest) GetCustomproperty27() string`

GetCustomproperty27 returns the Customproperty27 field if non-nil, zero value otherwise.

### GetCustomproperty27Ok

`func (o *CreateUpdateEntitlementRequest) GetCustomproperty27Ok() (*string, bool)`

GetCustomproperty27Ok returns a tuple with the Customproperty27 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty27

`func (o *CreateUpdateEntitlementRequest) SetCustomproperty27(v string)`

SetCustomproperty27 sets Customproperty27 field to given value.

### HasCustomproperty27

`func (o *CreateUpdateEntitlementRequest) HasCustomproperty27() bool`

HasCustomproperty27 returns a boolean if a field has been set.

### GetCustomproperty28

`func (o *CreateUpdateEntitlementRequest) GetCustomproperty28() string`

GetCustomproperty28 returns the Customproperty28 field if non-nil, zero value otherwise.

### GetCustomproperty28Ok

`func (o *CreateUpdateEntitlementRequest) GetCustomproperty28Ok() (*string, bool)`

GetCustomproperty28Ok returns a tuple with the Customproperty28 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty28

`func (o *CreateUpdateEntitlementRequest) SetCustomproperty28(v string)`

SetCustomproperty28 sets Customproperty28 field to given value.

### HasCustomproperty28

`func (o *CreateUpdateEntitlementRequest) HasCustomproperty28() bool`

HasCustomproperty28 returns a boolean if a field has been set.

### GetCustomproperty29

`func (o *CreateUpdateEntitlementRequest) GetCustomproperty29() string`

GetCustomproperty29 returns the Customproperty29 field if non-nil, zero value otherwise.

### GetCustomproperty29Ok

`func (o *CreateUpdateEntitlementRequest) GetCustomproperty29Ok() (*string, bool)`

GetCustomproperty29Ok returns a tuple with the Customproperty29 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty29

`func (o *CreateUpdateEntitlementRequest) SetCustomproperty29(v string)`

SetCustomproperty29 sets Customproperty29 field to given value.

### HasCustomproperty29

`func (o *CreateUpdateEntitlementRequest) HasCustomproperty29() bool`

HasCustomproperty29 returns a boolean if a field has been set.

### GetCustomproperty30

`func (o *CreateUpdateEntitlementRequest) GetCustomproperty30() string`

GetCustomproperty30 returns the Customproperty30 field if non-nil, zero value otherwise.

### GetCustomproperty30Ok

`func (o *CreateUpdateEntitlementRequest) GetCustomproperty30Ok() (*string, bool)`

GetCustomproperty30Ok returns a tuple with the Customproperty30 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty30

`func (o *CreateUpdateEntitlementRequest) SetCustomproperty30(v string)`

SetCustomproperty30 sets Customproperty30 field to given value.

### HasCustomproperty30

`func (o *CreateUpdateEntitlementRequest) HasCustomproperty30() bool`

HasCustomproperty30 returns a boolean if a field has been set.

### GetCustomproperty31

`func (o *CreateUpdateEntitlementRequest) GetCustomproperty31() string`

GetCustomproperty31 returns the Customproperty31 field if non-nil, zero value otherwise.

### GetCustomproperty31Ok

`func (o *CreateUpdateEntitlementRequest) GetCustomproperty31Ok() (*string, bool)`

GetCustomproperty31Ok returns a tuple with the Customproperty31 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty31

`func (o *CreateUpdateEntitlementRequest) SetCustomproperty31(v string)`

SetCustomproperty31 sets Customproperty31 field to given value.

### HasCustomproperty31

`func (o *CreateUpdateEntitlementRequest) HasCustomproperty31() bool`

HasCustomproperty31 returns a boolean if a field has been set.

### GetCustomproperty32

`func (o *CreateUpdateEntitlementRequest) GetCustomproperty32() string`

GetCustomproperty32 returns the Customproperty32 field if non-nil, zero value otherwise.

### GetCustomproperty32Ok

`func (o *CreateUpdateEntitlementRequest) GetCustomproperty32Ok() (*string, bool)`

GetCustomproperty32Ok returns a tuple with the Customproperty32 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty32

`func (o *CreateUpdateEntitlementRequest) SetCustomproperty32(v string)`

SetCustomproperty32 sets Customproperty32 field to given value.

### HasCustomproperty32

`func (o *CreateUpdateEntitlementRequest) HasCustomproperty32() bool`

HasCustomproperty32 returns a boolean if a field has been set.

### GetCustomproperty33

`func (o *CreateUpdateEntitlementRequest) GetCustomproperty33() string`

GetCustomproperty33 returns the Customproperty33 field if non-nil, zero value otherwise.

### GetCustomproperty33Ok

`func (o *CreateUpdateEntitlementRequest) GetCustomproperty33Ok() (*string, bool)`

GetCustomproperty33Ok returns a tuple with the Customproperty33 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty33

`func (o *CreateUpdateEntitlementRequest) SetCustomproperty33(v string)`

SetCustomproperty33 sets Customproperty33 field to given value.

### HasCustomproperty33

`func (o *CreateUpdateEntitlementRequest) HasCustomproperty33() bool`

HasCustomproperty33 returns a boolean if a field has been set.

### GetCustomproperty34

`func (o *CreateUpdateEntitlementRequest) GetCustomproperty34() string`

GetCustomproperty34 returns the Customproperty34 field if non-nil, zero value otherwise.

### GetCustomproperty34Ok

`func (o *CreateUpdateEntitlementRequest) GetCustomproperty34Ok() (*string, bool)`

GetCustomproperty34Ok returns a tuple with the Customproperty34 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty34

`func (o *CreateUpdateEntitlementRequest) SetCustomproperty34(v string)`

SetCustomproperty34 sets Customproperty34 field to given value.

### HasCustomproperty34

`func (o *CreateUpdateEntitlementRequest) HasCustomproperty34() bool`

HasCustomproperty34 returns a boolean if a field has been set.

### GetCustomproperty35

`func (o *CreateUpdateEntitlementRequest) GetCustomproperty35() string`

GetCustomproperty35 returns the Customproperty35 field if non-nil, zero value otherwise.

### GetCustomproperty35Ok

`func (o *CreateUpdateEntitlementRequest) GetCustomproperty35Ok() (*string, bool)`

GetCustomproperty35Ok returns a tuple with the Customproperty35 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty35

`func (o *CreateUpdateEntitlementRequest) SetCustomproperty35(v string)`

SetCustomproperty35 sets Customproperty35 field to given value.

### HasCustomproperty35

`func (o *CreateUpdateEntitlementRequest) HasCustomproperty35() bool`

HasCustomproperty35 returns a boolean if a field has been set.

### GetCustomproperty36

`func (o *CreateUpdateEntitlementRequest) GetCustomproperty36() string`

GetCustomproperty36 returns the Customproperty36 field if non-nil, zero value otherwise.

### GetCustomproperty36Ok

`func (o *CreateUpdateEntitlementRequest) GetCustomproperty36Ok() (*string, bool)`

GetCustomproperty36Ok returns a tuple with the Customproperty36 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty36

`func (o *CreateUpdateEntitlementRequest) SetCustomproperty36(v string)`

SetCustomproperty36 sets Customproperty36 field to given value.

### HasCustomproperty36

`func (o *CreateUpdateEntitlementRequest) HasCustomproperty36() bool`

HasCustomproperty36 returns a boolean if a field has been set.

### GetCustomproperty37

`func (o *CreateUpdateEntitlementRequest) GetCustomproperty37() string`

GetCustomproperty37 returns the Customproperty37 field if non-nil, zero value otherwise.

### GetCustomproperty37Ok

`func (o *CreateUpdateEntitlementRequest) GetCustomproperty37Ok() (*string, bool)`

GetCustomproperty37Ok returns a tuple with the Customproperty37 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty37

`func (o *CreateUpdateEntitlementRequest) SetCustomproperty37(v string)`

SetCustomproperty37 sets Customproperty37 field to given value.

### HasCustomproperty37

`func (o *CreateUpdateEntitlementRequest) HasCustomproperty37() bool`

HasCustomproperty37 returns a boolean if a field has been set.

### GetCustomproperty38

`func (o *CreateUpdateEntitlementRequest) GetCustomproperty38() string`

GetCustomproperty38 returns the Customproperty38 field if non-nil, zero value otherwise.

### GetCustomproperty38Ok

`func (o *CreateUpdateEntitlementRequest) GetCustomproperty38Ok() (*string, bool)`

GetCustomproperty38Ok returns a tuple with the Customproperty38 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty38

`func (o *CreateUpdateEntitlementRequest) SetCustomproperty38(v string)`

SetCustomproperty38 sets Customproperty38 field to given value.

### HasCustomproperty38

`func (o *CreateUpdateEntitlementRequest) HasCustomproperty38() bool`

HasCustomproperty38 returns a boolean if a field has been set.

### GetCustomproperty39

`func (o *CreateUpdateEntitlementRequest) GetCustomproperty39() string`

GetCustomproperty39 returns the Customproperty39 field if non-nil, zero value otherwise.

### GetCustomproperty39Ok

`func (o *CreateUpdateEntitlementRequest) GetCustomproperty39Ok() (*string, bool)`

GetCustomproperty39Ok returns a tuple with the Customproperty39 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty39

`func (o *CreateUpdateEntitlementRequest) SetCustomproperty39(v string)`

SetCustomproperty39 sets Customproperty39 field to given value.

### HasCustomproperty39

`func (o *CreateUpdateEntitlementRequest) HasCustomproperty39() bool`

HasCustomproperty39 returns a boolean if a field has been set.

### GetCustomproperty40

`func (o *CreateUpdateEntitlementRequest) GetCustomproperty40() string`

GetCustomproperty40 returns the Customproperty40 field if non-nil, zero value otherwise.

### GetCustomproperty40Ok

`func (o *CreateUpdateEntitlementRequest) GetCustomproperty40Ok() (*string, bool)`

GetCustomproperty40Ok returns a tuple with the Customproperty40 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty40

`func (o *CreateUpdateEntitlementRequest) SetCustomproperty40(v string)`

SetCustomproperty40 sets Customproperty40 field to given value.

### HasCustomproperty40

`func (o *CreateUpdateEntitlementRequest) HasCustomproperty40() bool`

HasCustomproperty40 returns a boolean if a field has been set.

### GetEntitlementowner1

`func (o *CreateUpdateEntitlementRequest) GetEntitlementowner1() []string`

GetEntitlementowner1 returns the Entitlementowner1 field if non-nil, zero value otherwise.

### GetEntitlementowner1Ok

`func (o *CreateUpdateEntitlementRequest) GetEntitlementowner1Ok() (*[]string, bool)`

GetEntitlementowner1Ok returns a tuple with the Entitlementowner1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitlementowner1

`func (o *CreateUpdateEntitlementRequest) SetEntitlementowner1(v []string)`

SetEntitlementowner1 sets Entitlementowner1 field to given value.

### HasEntitlementowner1

`func (o *CreateUpdateEntitlementRequest) HasEntitlementowner1() bool`

HasEntitlementowner1 returns a boolean if a field has been set.

### GetEntitlementowner2

`func (o *CreateUpdateEntitlementRequest) GetEntitlementowner2() []string`

GetEntitlementowner2 returns the Entitlementowner2 field if non-nil, zero value otherwise.

### GetEntitlementowner2Ok

`func (o *CreateUpdateEntitlementRequest) GetEntitlementowner2Ok() (*[]string, bool)`

GetEntitlementowner2Ok returns a tuple with the Entitlementowner2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitlementowner2

`func (o *CreateUpdateEntitlementRequest) SetEntitlementowner2(v []string)`

SetEntitlementowner2 sets Entitlementowner2 field to given value.

### HasEntitlementowner2

`func (o *CreateUpdateEntitlementRequest) HasEntitlementowner2() bool`

HasEntitlementowner2 returns a boolean if a field has been set.

### GetEntitlementowner3

`func (o *CreateUpdateEntitlementRequest) GetEntitlementowner3() []string`

GetEntitlementowner3 returns the Entitlementowner3 field if non-nil, zero value otherwise.

### GetEntitlementowner3Ok

`func (o *CreateUpdateEntitlementRequest) GetEntitlementowner3Ok() (*[]string, bool)`

GetEntitlementowner3Ok returns a tuple with the Entitlementowner3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitlementowner3

`func (o *CreateUpdateEntitlementRequest) SetEntitlementowner3(v []string)`

SetEntitlementowner3 sets Entitlementowner3 field to given value.

### HasEntitlementowner3

`func (o *CreateUpdateEntitlementRequest) HasEntitlementowner3() bool`

HasEntitlementowner3 returns a boolean if a field has been set.

### GetEntitlementowner4

`func (o *CreateUpdateEntitlementRequest) GetEntitlementowner4() []string`

GetEntitlementowner4 returns the Entitlementowner4 field if non-nil, zero value otherwise.

### GetEntitlementowner4Ok

`func (o *CreateUpdateEntitlementRequest) GetEntitlementowner4Ok() (*[]string, bool)`

GetEntitlementowner4Ok returns a tuple with the Entitlementowner4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitlementowner4

`func (o *CreateUpdateEntitlementRequest) SetEntitlementowner4(v []string)`

SetEntitlementowner4 sets Entitlementowner4 field to given value.

### HasEntitlementowner4

`func (o *CreateUpdateEntitlementRequest) HasEntitlementowner4() bool`

HasEntitlementowner4 returns a boolean if a field has been set.

### GetEntitlementowner5

`func (o *CreateUpdateEntitlementRequest) GetEntitlementowner5() []string`

GetEntitlementowner5 returns the Entitlementowner5 field if non-nil, zero value otherwise.

### GetEntitlementowner5Ok

`func (o *CreateUpdateEntitlementRequest) GetEntitlementowner5Ok() (*[]string, bool)`

GetEntitlementowner5Ok returns a tuple with the Entitlementowner5 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitlementowner5

`func (o *CreateUpdateEntitlementRequest) SetEntitlementowner5(v []string)`

SetEntitlementowner5 sets Entitlementowner5 field to given value.

### HasEntitlementowner5

`func (o *CreateUpdateEntitlementRequest) HasEntitlementowner5() bool`

HasEntitlementowner5 returns a boolean if a field has been set.

### GetEntitlementowner6

`func (o *CreateUpdateEntitlementRequest) GetEntitlementowner6() []string`

GetEntitlementowner6 returns the Entitlementowner6 field if non-nil, zero value otherwise.

### GetEntitlementowner6Ok

`func (o *CreateUpdateEntitlementRequest) GetEntitlementowner6Ok() (*[]string, bool)`

GetEntitlementowner6Ok returns a tuple with the Entitlementowner6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitlementowner6

`func (o *CreateUpdateEntitlementRequest) SetEntitlementowner6(v []string)`

SetEntitlementowner6 sets Entitlementowner6 field to given value.

### HasEntitlementowner6

`func (o *CreateUpdateEntitlementRequest) HasEntitlementowner6() bool`

HasEntitlementowner6 returns a boolean if a field has been set.

### GetEntitlementowner7

`func (o *CreateUpdateEntitlementRequest) GetEntitlementowner7() []string`

GetEntitlementowner7 returns the Entitlementowner7 field if non-nil, zero value otherwise.

### GetEntitlementowner7Ok

`func (o *CreateUpdateEntitlementRequest) GetEntitlementowner7Ok() (*[]string, bool)`

GetEntitlementowner7Ok returns a tuple with the Entitlementowner7 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitlementowner7

`func (o *CreateUpdateEntitlementRequest) SetEntitlementowner7(v []string)`

SetEntitlementowner7 sets Entitlementowner7 field to given value.

### HasEntitlementowner7

`func (o *CreateUpdateEntitlementRequest) HasEntitlementowner7() bool`

HasEntitlementowner7 returns a boolean if a field has been set.

### GetEntitlementowner8

`func (o *CreateUpdateEntitlementRequest) GetEntitlementowner8() []string`

GetEntitlementowner8 returns the Entitlementowner8 field if non-nil, zero value otherwise.

### GetEntitlementowner8Ok

`func (o *CreateUpdateEntitlementRequest) GetEntitlementowner8Ok() (*[]string, bool)`

GetEntitlementowner8Ok returns a tuple with the Entitlementowner8 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitlementowner8

`func (o *CreateUpdateEntitlementRequest) SetEntitlementowner8(v []string)`

SetEntitlementowner8 sets Entitlementowner8 field to given value.

### HasEntitlementowner8

`func (o *CreateUpdateEntitlementRequest) HasEntitlementowner8() bool`

HasEntitlementowner8 returns a boolean if a field has been set.

### GetEntitlementowner9

`func (o *CreateUpdateEntitlementRequest) GetEntitlementowner9() []string`

GetEntitlementowner9 returns the Entitlementowner9 field if non-nil, zero value otherwise.

### GetEntitlementowner9Ok

`func (o *CreateUpdateEntitlementRequest) GetEntitlementowner9Ok() (*[]string, bool)`

GetEntitlementowner9Ok returns a tuple with the Entitlementowner9 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitlementowner9

`func (o *CreateUpdateEntitlementRequest) SetEntitlementowner9(v []string)`

SetEntitlementowner9 sets Entitlementowner9 field to given value.

### HasEntitlementowner9

`func (o *CreateUpdateEntitlementRequest) HasEntitlementowner9() bool`

HasEntitlementowner9 returns a boolean if a field has been set.

### GetEntitlementowner10

`func (o *CreateUpdateEntitlementRequest) GetEntitlementowner10() []string`

GetEntitlementowner10 returns the Entitlementowner10 field if non-nil, zero value otherwise.

### GetEntitlementowner10Ok

`func (o *CreateUpdateEntitlementRequest) GetEntitlementowner10Ok() (*[]string, bool)`

GetEntitlementowner10Ok returns a tuple with the Entitlementowner10 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitlementowner10

`func (o *CreateUpdateEntitlementRequest) SetEntitlementowner10(v []string)`

SetEntitlementowner10 sets Entitlementowner10 field to given value.

### HasEntitlementowner10

`func (o *CreateUpdateEntitlementRequest) HasEntitlementowner10() bool`

HasEntitlementowner10 returns a boolean if a field has been set.

### GetEntitlementowner11

`func (o *CreateUpdateEntitlementRequest) GetEntitlementowner11() []string`

GetEntitlementowner11 returns the Entitlementowner11 field if non-nil, zero value otherwise.

### GetEntitlementowner11Ok

`func (o *CreateUpdateEntitlementRequest) GetEntitlementowner11Ok() (*[]string, bool)`

GetEntitlementowner11Ok returns a tuple with the Entitlementowner11 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitlementowner11

`func (o *CreateUpdateEntitlementRequest) SetEntitlementowner11(v []string)`

SetEntitlementowner11 sets Entitlementowner11 field to given value.

### HasEntitlementowner11

`func (o *CreateUpdateEntitlementRequest) HasEntitlementowner11() bool`

HasEntitlementowner11 returns a boolean if a field has been set.

### GetEntitlementowner12

`func (o *CreateUpdateEntitlementRequest) GetEntitlementowner12() []string`

GetEntitlementowner12 returns the Entitlementowner12 field if non-nil, zero value otherwise.

### GetEntitlementowner12Ok

`func (o *CreateUpdateEntitlementRequest) GetEntitlementowner12Ok() (*[]string, bool)`

GetEntitlementowner12Ok returns a tuple with the Entitlementowner12 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitlementowner12

`func (o *CreateUpdateEntitlementRequest) SetEntitlementowner12(v []string)`

SetEntitlementowner12 sets Entitlementowner12 field to given value.

### HasEntitlementowner12

`func (o *CreateUpdateEntitlementRequest) HasEntitlementowner12() bool`

HasEntitlementowner12 returns a boolean if a field has been set.

### GetEntitlementowner13

`func (o *CreateUpdateEntitlementRequest) GetEntitlementowner13() []string`

GetEntitlementowner13 returns the Entitlementowner13 field if non-nil, zero value otherwise.

### GetEntitlementowner13Ok

`func (o *CreateUpdateEntitlementRequest) GetEntitlementowner13Ok() (*[]string, bool)`

GetEntitlementowner13Ok returns a tuple with the Entitlementowner13 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitlementowner13

`func (o *CreateUpdateEntitlementRequest) SetEntitlementowner13(v []string)`

SetEntitlementowner13 sets Entitlementowner13 field to given value.

### HasEntitlementowner13

`func (o *CreateUpdateEntitlementRequest) HasEntitlementowner13() bool`

HasEntitlementowner13 returns a boolean if a field has been set.

### GetEntitlementowner14

`func (o *CreateUpdateEntitlementRequest) GetEntitlementowner14() []string`

GetEntitlementowner14 returns the Entitlementowner14 field if non-nil, zero value otherwise.

### GetEntitlementowner14Ok

`func (o *CreateUpdateEntitlementRequest) GetEntitlementowner14Ok() (*[]string, bool)`

GetEntitlementowner14Ok returns a tuple with the Entitlementowner14 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitlementowner14

`func (o *CreateUpdateEntitlementRequest) SetEntitlementowner14(v []string)`

SetEntitlementowner14 sets Entitlementowner14 field to given value.

### HasEntitlementowner14

`func (o *CreateUpdateEntitlementRequest) HasEntitlementowner14() bool`

HasEntitlementowner14 returns a boolean if a field has been set.

### GetEntitlementowner15

`func (o *CreateUpdateEntitlementRequest) GetEntitlementowner15() []string`

GetEntitlementowner15 returns the Entitlementowner15 field if non-nil, zero value otherwise.

### GetEntitlementowner15Ok

`func (o *CreateUpdateEntitlementRequest) GetEntitlementowner15Ok() (*[]string, bool)`

GetEntitlementowner15Ok returns a tuple with the Entitlementowner15 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitlementowner15

`func (o *CreateUpdateEntitlementRequest) SetEntitlementowner15(v []string)`

SetEntitlementowner15 sets Entitlementowner15 field to given value.

### HasEntitlementowner15

`func (o *CreateUpdateEntitlementRequest) HasEntitlementowner15() bool`

HasEntitlementowner15 returns a boolean if a field has been set.

### GetEntitlementowner16

`func (o *CreateUpdateEntitlementRequest) GetEntitlementowner16() []string`

GetEntitlementowner16 returns the Entitlementowner16 field if non-nil, zero value otherwise.

### GetEntitlementowner16Ok

`func (o *CreateUpdateEntitlementRequest) GetEntitlementowner16Ok() (*[]string, bool)`

GetEntitlementowner16Ok returns a tuple with the Entitlementowner16 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitlementowner16

`func (o *CreateUpdateEntitlementRequest) SetEntitlementowner16(v []string)`

SetEntitlementowner16 sets Entitlementowner16 field to given value.

### HasEntitlementowner16

`func (o *CreateUpdateEntitlementRequest) HasEntitlementowner16() bool`

HasEntitlementowner16 returns a boolean if a field has been set.

### GetEntitlementowner17

`func (o *CreateUpdateEntitlementRequest) GetEntitlementowner17() []string`

GetEntitlementowner17 returns the Entitlementowner17 field if non-nil, zero value otherwise.

### GetEntitlementowner17Ok

`func (o *CreateUpdateEntitlementRequest) GetEntitlementowner17Ok() (*[]string, bool)`

GetEntitlementowner17Ok returns a tuple with the Entitlementowner17 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitlementowner17

`func (o *CreateUpdateEntitlementRequest) SetEntitlementowner17(v []string)`

SetEntitlementowner17 sets Entitlementowner17 field to given value.

### HasEntitlementowner17

`func (o *CreateUpdateEntitlementRequest) HasEntitlementowner17() bool`

HasEntitlementowner17 returns a boolean if a field has been set.

### GetEntitlementowner18

`func (o *CreateUpdateEntitlementRequest) GetEntitlementowner18() []string`

GetEntitlementowner18 returns the Entitlementowner18 field if non-nil, zero value otherwise.

### GetEntitlementowner18Ok

`func (o *CreateUpdateEntitlementRequest) GetEntitlementowner18Ok() (*[]string, bool)`

GetEntitlementowner18Ok returns a tuple with the Entitlementowner18 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitlementowner18

`func (o *CreateUpdateEntitlementRequest) SetEntitlementowner18(v []string)`

SetEntitlementowner18 sets Entitlementowner18 field to given value.

### HasEntitlementowner18

`func (o *CreateUpdateEntitlementRequest) HasEntitlementowner18() bool`

HasEntitlementowner18 returns a boolean if a field has been set.

### GetEntitlementowner19

`func (o *CreateUpdateEntitlementRequest) GetEntitlementowner19() []string`

GetEntitlementowner19 returns the Entitlementowner19 field if non-nil, zero value otherwise.

### GetEntitlementowner19Ok

`func (o *CreateUpdateEntitlementRequest) GetEntitlementowner19Ok() (*[]string, bool)`

GetEntitlementowner19Ok returns a tuple with the Entitlementowner19 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitlementowner19

`func (o *CreateUpdateEntitlementRequest) SetEntitlementowner19(v []string)`

SetEntitlementowner19 sets Entitlementowner19 field to given value.

### HasEntitlementowner19

`func (o *CreateUpdateEntitlementRequest) HasEntitlementowner19() bool`

HasEntitlementowner19 returns a boolean if a field has been set.

### GetEntitlementowner20

`func (o *CreateUpdateEntitlementRequest) GetEntitlementowner20() []string`

GetEntitlementowner20 returns the Entitlementowner20 field if non-nil, zero value otherwise.

### GetEntitlementowner20Ok

`func (o *CreateUpdateEntitlementRequest) GetEntitlementowner20Ok() (*[]string, bool)`

GetEntitlementowner20Ok returns a tuple with the Entitlementowner20 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitlementowner20

`func (o *CreateUpdateEntitlementRequest) SetEntitlementowner20(v []string)`

SetEntitlementowner20 sets Entitlementowner20 field to given value.

### HasEntitlementowner20

`func (o *CreateUpdateEntitlementRequest) HasEntitlementowner20() bool`

HasEntitlementowner20 returns a boolean if a field has been set.

### GetEntitlementowner21

`func (o *CreateUpdateEntitlementRequest) GetEntitlementowner21() []string`

GetEntitlementowner21 returns the Entitlementowner21 field if non-nil, zero value otherwise.

### GetEntitlementowner21Ok

`func (o *CreateUpdateEntitlementRequest) GetEntitlementowner21Ok() (*[]string, bool)`

GetEntitlementowner21Ok returns a tuple with the Entitlementowner21 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitlementowner21

`func (o *CreateUpdateEntitlementRequest) SetEntitlementowner21(v []string)`

SetEntitlementowner21 sets Entitlementowner21 field to given value.

### HasEntitlementowner21

`func (o *CreateUpdateEntitlementRequest) HasEntitlementowner21() bool`

HasEntitlementowner21 returns a boolean if a field has been set.

### GetEntitlementowner22

`func (o *CreateUpdateEntitlementRequest) GetEntitlementowner22() []string`

GetEntitlementowner22 returns the Entitlementowner22 field if non-nil, zero value otherwise.

### GetEntitlementowner22Ok

`func (o *CreateUpdateEntitlementRequest) GetEntitlementowner22Ok() (*[]string, bool)`

GetEntitlementowner22Ok returns a tuple with the Entitlementowner22 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitlementowner22

`func (o *CreateUpdateEntitlementRequest) SetEntitlementowner22(v []string)`

SetEntitlementowner22 sets Entitlementowner22 field to given value.

### HasEntitlementowner22

`func (o *CreateUpdateEntitlementRequest) HasEntitlementowner22() bool`

HasEntitlementowner22 returns a boolean if a field has been set.

### GetEntitlementowner23

`func (o *CreateUpdateEntitlementRequest) GetEntitlementowner23() []string`

GetEntitlementowner23 returns the Entitlementowner23 field if non-nil, zero value otherwise.

### GetEntitlementowner23Ok

`func (o *CreateUpdateEntitlementRequest) GetEntitlementowner23Ok() (*[]string, bool)`

GetEntitlementowner23Ok returns a tuple with the Entitlementowner23 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitlementowner23

`func (o *CreateUpdateEntitlementRequest) SetEntitlementowner23(v []string)`

SetEntitlementowner23 sets Entitlementowner23 field to given value.

### HasEntitlementowner23

`func (o *CreateUpdateEntitlementRequest) HasEntitlementowner23() bool`

HasEntitlementowner23 returns a boolean if a field has been set.

### GetEntitlementowner24

`func (o *CreateUpdateEntitlementRequest) GetEntitlementowner24() []string`

GetEntitlementowner24 returns the Entitlementowner24 field if non-nil, zero value otherwise.

### GetEntitlementowner24Ok

`func (o *CreateUpdateEntitlementRequest) GetEntitlementowner24Ok() (*[]string, bool)`

GetEntitlementowner24Ok returns a tuple with the Entitlementowner24 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitlementowner24

`func (o *CreateUpdateEntitlementRequest) SetEntitlementowner24(v []string)`

SetEntitlementowner24 sets Entitlementowner24 field to given value.

### HasEntitlementowner24

`func (o *CreateUpdateEntitlementRequest) HasEntitlementowner24() bool`

HasEntitlementowner24 returns a boolean if a field has been set.

### GetEntitlementowner25

`func (o *CreateUpdateEntitlementRequest) GetEntitlementowner25() []string`

GetEntitlementowner25 returns the Entitlementowner25 field if non-nil, zero value otherwise.

### GetEntitlementowner25Ok

`func (o *CreateUpdateEntitlementRequest) GetEntitlementowner25Ok() (*[]string, bool)`

GetEntitlementowner25Ok returns a tuple with the Entitlementowner25 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitlementowner25

`func (o *CreateUpdateEntitlementRequest) SetEntitlementowner25(v []string)`

SetEntitlementowner25 sets Entitlementowner25 field to given value.

### HasEntitlementowner25

`func (o *CreateUpdateEntitlementRequest) HasEntitlementowner25() bool`

HasEntitlementowner25 returns a boolean if a field has been set.

### GetEntitlementowner26

`func (o *CreateUpdateEntitlementRequest) GetEntitlementowner26() []string`

GetEntitlementowner26 returns the Entitlementowner26 field if non-nil, zero value otherwise.

### GetEntitlementowner26Ok

`func (o *CreateUpdateEntitlementRequest) GetEntitlementowner26Ok() (*[]string, bool)`

GetEntitlementowner26Ok returns a tuple with the Entitlementowner26 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitlementowner26

`func (o *CreateUpdateEntitlementRequest) SetEntitlementowner26(v []string)`

SetEntitlementowner26 sets Entitlementowner26 field to given value.

### HasEntitlementowner26

`func (o *CreateUpdateEntitlementRequest) HasEntitlementowner26() bool`

HasEntitlementowner26 returns a boolean if a field has been set.

### GetEntitlementowner27

`func (o *CreateUpdateEntitlementRequest) GetEntitlementowner27() []string`

GetEntitlementowner27 returns the Entitlementowner27 field if non-nil, zero value otherwise.

### GetEntitlementowner27Ok

`func (o *CreateUpdateEntitlementRequest) GetEntitlementowner27Ok() (*[]string, bool)`

GetEntitlementowner27Ok returns a tuple with the Entitlementowner27 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitlementowner27

`func (o *CreateUpdateEntitlementRequest) SetEntitlementowner27(v []string)`

SetEntitlementowner27 sets Entitlementowner27 field to given value.

### HasEntitlementowner27

`func (o *CreateUpdateEntitlementRequest) HasEntitlementowner27() bool`

HasEntitlementowner27 returns a boolean if a field has been set.

### GetEntitlementmap

`func (o *CreateUpdateEntitlementRequest) GetEntitlementmap() []CreateUpdateEntitlementRequestEntitlementmapInner`

GetEntitlementmap returns the Entitlementmap field if non-nil, zero value otherwise.

### GetEntitlementmapOk

`func (o *CreateUpdateEntitlementRequest) GetEntitlementmapOk() (*[]CreateUpdateEntitlementRequestEntitlementmapInner, bool)`

GetEntitlementmapOk returns a tuple with the Entitlementmap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitlementmap

`func (o *CreateUpdateEntitlementRequest) SetEntitlementmap(v []CreateUpdateEntitlementRequestEntitlementmapInner)`

SetEntitlementmap sets Entitlementmap field to given value.

### HasEntitlementmap

`func (o *CreateUpdateEntitlementRequest) HasEntitlementmap() bool`

HasEntitlementmap returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


