# CreateEntitlementTypeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Entitlementname** | **string** | Name of the entitlement type | 
**DisplayName** | Pointer to **string** | Display name for the entitlement. | [optional] 
**Endpointname** | **string** | Name of the endpoint with which the entitlement type is associated | 
**Entitlementdescription** | Pointer to **string** | Description for the entitlement type | [optional] 
**Workflow** | Pointer to **string** | Workflow associated with the entitlement type | [optional] 
**Orderindex** | Pointer to **int32** | Index to determine the order of processing or display. | [optional] 
**EnableEntitlementToRoleSync** | Pointer to **bool** | Enable entitlement to role sync | [optional] 
**Certifiable** | Pointer to **bool** | Indicates if the entitlement is certifiable. | [optional] 
**ShowEntTypeOn** | Pointer to **string** | Show entitlement type on | [optional] 
**AvailableQueryServiceAccount** | Pointer to **string** | Available query service account | [optional] 
**SelectedQueryServiceAccount** | Pointer to **string** | Selected query service account | [optional] 
**RequestDatesConfJson** | Pointer to **string** | Configuration in JSON for handling request dates. | [optional] 
**Requiredinrequest** | Pointer to **bool** | Required in request | [optional] 
**Requiredinservicerequest** | Pointer to **bool** | Required in service request | [optional] 
**Recon** | Pointer to **bool** | Recon | [optional] 
**Showonchild** | Pointer to **bool** | Show on child | [optional] 
**Hiearchyrequired** | Pointer to **int32** | Flag indicating if a hierarchy is required. | [optional] 
**EnableProvisioningPriority** | Pointer to **bool** | Enable provisioning priority | [optional] 
**ExcludeRuleAssgnEntsInRequest** | Pointer to **bool** | Exclude Entitlements Assigned via Rule while Request | [optional] 
**Customproperty1** | Pointer to **string** | Custom property 1 | [optional] 
**Customproperty2** | Pointer to **string** | Custom property 2 | [optional] 
**Customproperty3** | Pointer to **string** | Custom property 3 | [optional] 
**Customproperty4** | Pointer to **string** | Custom property 4 | [optional] 
**Customproperty5** | Pointer to **string** | Custom property 5 | [optional] 
**Customproperty6** | Pointer to **string** | Custom property 6 | [optional] 
**Customproperty7** | Pointer to **string** | Custom property 7 | [optional] 
**Customproperty8** | Pointer to **string** | Custom property 8 | [optional] 
**Customproperty9** | Pointer to **string** | Custom property 9 | [optional] 
**Customproperty10** | Pointer to **string** | Custom property 10 | [optional] 
**Customproperty11** | Pointer to **string** | Custom property 11 | [optional] 
**Customproperty12** | Pointer to **string** | Custom property 12 | [optional] 
**Customproperty13** | Pointer to **string** | Custom property 13 | [optional] 
**Customproperty14** | Pointer to **string** | Custom property 14 | [optional] 
**Customproperty15** | Pointer to **string** | Custom property 15 | [optional] 
**Customproperty16** | Pointer to **string** | Custom property 16 | [optional] 
**Customproperty17** | Pointer to **string** | Custom property 17 | [optional] 
**Customproperty18** | Pointer to **string** | Custom property 18 | [optional] 
**Customproperty19** | Pointer to **string** | Custom property 19 | [optional] 
**Customproperty20** | Pointer to **string** | Custom property 20 | [optional] 
**Customproperty21** | Pointer to **string** | Custom property 21 | [optional] 
**Customproperty22** | Pointer to **string** | Custom property 22 | [optional] 
**Customproperty23** | Pointer to **string** | Custom property 23 | [optional] 
**Customproperty24** | Pointer to **string** | Custom property 24 | [optional] 
**Customproperty25** | Pointer to **string** | Custom property 25 | [optional] 
**Customproperty26** | Pointer to **string** | Custom property 26 | [optional] 
**Customproperty27** | Pointer to **string** | Custom property 27 | [optional] 
**Customproperty28** | Pointer to **string** | Custom property 28 | [optional] 
**Customproperty29** | Pointer to **string** | Custom property 29 | [optional] 
**Customproperty30** | Pointer to **string** | Custom property 30 | [optional] 
**Customproperty31** | Pointer to **string** | Custom property 31 | [optional] 
**Customproperty32** | Pointer to **string** | Custom property 32 | [optional] 
**Customproperty33** | Pointer to **string** | Custom property 33 | [optional] 
**Customproperty34** | Pointer to **string** | Custom property 34 | [optional] 
**Customproperty35** | Pointer to **string** | Custom property 35 | [optional] 
**Customproperty36** | Pointer to **string** | Custom property 36 | [optional] 
**Customproperty37** | Pointer to **string** | Custom property 37 | [optional] 
**Customproperty38** | Pointer to **string** | Custom property 38 | [optional] 
**Customproperty39** | Pointer to **string** | Custom property 39 | [optional] 
**Customproperty40** | Pointer to **string** | Custom property 40 | [optional] 

## Methods

### NewCreateEntitlementTypeRequest

`func NewCreateEntitlementTypeRequest(entitlementname string, endpointname string, ) *CreateEntitlementTypeRequest`

NewCreateEntitlementTypeRequest instantiates a new CreateEntitlementTypeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateEntitlementTypeRequestWithDefaults

`func NewCreateEntitlementTypeRequestWithDefaults() *CreateEntitlementTypeRequest`

NewCreateEntitlementTypeRequestWithDefaults instantiates a new CreateEntitlementTypeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntitlementname

`func (o *CreateEntitlementTypeRequest) GetEntitlementname() string`

GetEntitlementname returns the Entitlementname field if non-nil, zero value otherwise.

### GetEntitlementnameOk

`func (o *CreateEntitlementTypeRequest) GetEntitlementnameOk() (*string, bool)`

GetEntitlementnameOk returns a tuple with the Entitlementname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitlementname

`func (o *CreateEntitlementTypeRequest) SetEntitlementname(v string)`

SetEntitlementname sets Entitlementname field to given value.


### GetDisplayName

`func (o *CreateEntitlementTypeRequest) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *CreateEntitlementTypeRequest) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *CreateEntitlementTypeRequest) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *CreateEntitlementTypeRequest) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetEndpointname

`func (o *CreateEntitlementTypeRequest) GetEndpointname() string`

GetEndpointname returns the Endpointname field if non-nil, zero value otherwise.

### GetEndpointnameOk

`func (o *CreateEntitlementTypeRequest) GetEndpointnameOk() (*string, bool)`

GetEndpointnameOk returns a tuple with the Endpointname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpointname

`func (o *CreateEntitlementTypeRequest) SetEndpointname(v string)`

SetEndpointname sets Endpointname field to given value.


### GetEntitlementdescription

`func (o *CreateEntitlementTypeRequest) GetEntitlementdescription() string`

GetEntitlementdescription returns the Entitlementdescription field if non-nil, zero value otherwise.

### GetEntitlementdescriptionOk

`func (o *CreateEntitlementTypeRequest) GetEntitlementdescriptionOk() (*string, bool)`

GetEntitlementdescriptionOk returns a tuple with the Entitlementdescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitlementdescription

`func (o *CreateEntitlementTypeRequest) SetEntitlementdescription(v string)`

SetEntitlementdescription sets Entitlementdescription field to given value.

### HasEntitlementdescription

`func (o *CreateEntitlementTypeRequest) HasEntitlementdescription() bool`

HasEntitlementdescription returns a boolean if a field has been set.

### GetWorkflow

`func (o *CreateEntitlementTypeRequest) GetWorkflow() string`

GetWorkflow returns the Workflow field if non-nil, zero value otherwise.

### GetWorkflowOk

`func (o *CreateEntitlementTypeRequest) GetWorkflowOk() (*string, bool)`

GetWorkflowOk returns a tuple with the Workflow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflow

`func (o *CreateEntitlementTypeRequest) SetWorkflow(v string)`

SetWorkflow sets Workflow field to given value.

### HasWorkflow

`func (o *CreateEntitlementTypeRequest) HasWorkflow() bool`

HasWorkflow returns a boolean if a field has been set.

### GetOrderindex

`func (o *CreateEntitlementTypeRequest) GetOrderindex() int32`

GetOrderindex returns the Orderindex field if non-nil, zero value otherwise.

### GetOrderindexOk

`func (o *CreateEntitlementTypeRequest) GetOrderindexOk() (*int32, bool)`

GetOrderindexOk returns a tuple with the Orderindex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderindex

`func (o *CreateEntitlementTypeRequest) SetOrderindex(v int32)`

SetOrderindex sets Orderindex field to given value.

### HasOrderindex

`func (o *CreateEntitlementTypeRequest) HasOrderindex() bool`

HasOrderindex returns a boolean if a field has been set.

### GetEnableEntitlementToRoleSync

`func (o *CreateEntitlementTypeRequest) GetEnableEntitlementToRoleSync() bool`

GetEnableEntitlementToRoleSync returns the EnableEntitlementToRoleSync field if non-nil, zero value otherwise.

### GetEnableEntitlementToRoleSyncOk

`func (o *CreateEntitlementTypeRequest) GetEnableEntitlementToRoleSyncOk() (*bool, bool)`

GetEnableEntitlementToRoleSyncOk returns a tuple with the EnableEntitlementToRoleSync field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableEntitlementToRoleSync

`func (o *CreateEntitlementTypeRequest) SetEnableEntitlementToRoleSync(v bool)`

SetEnableEntitlementToRoleSync sets EnableEntitlementToRoleSync field to given value.

### HasEnableEntitlementToRoleSync

`func (o *CreateEntitlementTypeRequest) HasEnableEntitlementToRoleSync() bool`

HasEnableEntitlementToRoleSync returns a boolean if a field has been set.

### GetCertifiable

`func (o *CreateEntitlementTypeRequest) GetCertifiable() bool`

GetCertifiable returns the Certifiable field if non-nil, zero value otherwise.

### GetCertifiableOk

`func (o *CreateEntitlementTypeRequest) GetCertifiableOk() (*bool, bool)`

GetCertifiableOk returns a tuple with the Certifiable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertifiable

`func (o *CreateEntitlementTypeRequest) SetCertifiable(v bool)`

SetCertifiable sets Certifiable field to given value.

### HasCertifiable

`func (o *CreateEntitlementTypeRequest) HasCertifiable() bool`

HasCertifiable returns a boolean if a field has been set.

### GetShowEntTypeOn

`func (o *CreateEntitlementTypeRequest) GetShowEntTypeOn() string`

GetShowEntTypeOn returns the ShowEntTypeOn field if non-nil, zero value otherwise.

### GetShowEntTypeOnOk

`func (o *CreateEntitlementTypeRequest) GetShowEntTypeOnOk() (*string, bool)`

GetShowEntTypeOnOk returns a tuple with the ShowEntTypeOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowEntTypeOn

`func (o *CreateEntitlementTypeRequest) SetShowEntTypeOn(v string)`

SetShowEntTypeOn sets ShowEntTypeOn field to given value.

### HasShowEntTypeOn

`func (o *CreateEntitlementTypeRequest) HasShowEntTypeOn() bool`

HasShowEntTypeOn returns a boolean if a field has been set.

### GetAvailableQueryServiceAccount

`func (o *CreateEntitlementTypeRequest) GetAvailableQueryServiceAccount() string`

GetAvailableQueryServiceAccount returns the AvailableQueryServiceAccount field if non-nil, zero value otherwise.

### GetAvailableQueryServiceAccountOk

`func (o *CreateEntitlementTypeRequest) GetAvailableQueryServiceAccountOk() (*string, bool)`

GetAvailableQueryServiceAccountOk returns a tuple with the AvailableQueryServiceAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableQueryServiceAccount

`func (o *CreateEntitlementTypeRequest) SetAvailableQueryServiceAccount(v string)`

SetAvailableQueryServiceAccount sets AvailableQueryServiceAccount field to given value.

### HasAvailableQueryServiceAccount

`func (o *CreateEntitlementTypeRequest) HasAvailableQueryServiceAccount() bool`

HasAvailableQueryServiceAccount returns a boolean if a field has been set.

### GetSelectedQueryServiceAccount

`func (o *CreateEntitlementTypeRequest) GetSelectedQueryServiceAccount() string`

GetSelectedQueryServiceAccount returns the SelectedQueryServiceAccount field if non-nil, zero value otherwise.

### GetSelectedQueryServiceAccountOk

`func (o *CreateEntitlementTypeRequest) GetSelectedQueryServiceAccountOk() (*string, bool)`

GetSelectedQueryServiceAccountOk returns a tuple with the SelectedQueryServiceAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectedQueryServiceAccount

`func (o *CreateEntitlementTypeRequest) SetSelectedQueryServiceAccount(v string)`

SetSelectedQueryServiceAccount sets SelectedQueryServiceAccount field to given value.

### HasSelectedQueryServiceAccount

`func (o *CreateEntitlementTypeRequest) HasSelectedQueryServiceAccount() bool`

HasSelectedQueryServiceAccount returns a boolean if a field has been set.

### GetRequestDatesConfJson

`func (o *CreateEntitlementTypeRequest) GetRequestDatesConfJson() string`

GetRequestDatesConfJson returns the RequestDatesConfJson field if non-nil, zero value otherwise.

### GetRequestDatesConfJsonOk

`func (o *CreateEntitlementTypeRequest) GetRequestDatesConfJsonOk() (*string, bool)`

GetRequestDatesConfJsonOk returns a tuple with the RequestDatesConfJson field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestDatesConfJson

`func (o *CreateEntitlementTypeRequest) SetRequestDatesConfJson(v string)`

SetRequestDatesConfJson sets RequestDatesConfJson field to given value.

### HasRequestDatesConfJson

`func (o *CreateEntitlementTypeRequest) HasRequestDatesConfJson() bool`

HasRequestDatesConfJson returns a boolean if a field has been set.

### GetRequiredinrequest

`func (o *CreateEntitlementTypeRequest) GetRequiredinrequest() bool`

GetRequiredinrequest returns the Requiredinrequest field if non-nil, zero value otherwise.

### GetRequiredinrequestOk

`func (o *CreateEntitlementTypeRequest) GetRequiredinrequestOk() (*bool, bool)`

GetRequiredinrequestOk returns a tuple with the Requiredinrequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredinrequest

`func (o *CreateEntitlementTypeRequest) SetRequiredinrequest(v bool)`

SetRequiredinrequest sets Requiredinrequest field to given value.

### HasRequiredinrequest

`func (o *CreateEntitlementTypeRequest) HasRequiredinrequest() bool`

HasRequiredinrequest returns a boolean if a field has been set.

### GetRequiredinservicerequest

`func (o *CreateEntitlementTypeRequest) GetRequiredinservicerequest() bool`

GetRequiredinservicerequest returns the Requiredinservicerequest field if non-nil, zero value otherwise.

### GetRequiredinservicerequestOk

`func (o *CreateEntitlementTypeRequest) GetRequiredinservicerequestOk() (*bool, bool)`

GetRequiredinservicerequestOk returns a tuple with the Requiredinservicerequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredinservicerequest

`func (o *CreateEntitlementTypeRequest) SetRequiredinservicerequest(v bool)`

SetRequiredinservicerequest sets Requiredinservicerequest field to given value.

### HasRequiredinservicerequest

`func (o *CreateEntitlementTypeRequest) HasRequiredinservicerequest() bool`

HasRequiredinservicerequest returns a boolean if a field has been set.

### GetRecon

`func (o *CreateEntitlementTypeRequest) GetRecon() bool`

GetRecon returns the Recon field if non-nil, zero value otherwise.

### GetReconOk

`func (o *CreateEntitlementTypeRequest) GetReconOk() (*bool, bool)`

GetReconOk returns a tuple with the Recon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecon

`func (o *CreateEntitlementTypeRequest) SetRecon(v bool)`

SetRecon sets Recon field to given value.

### HasRecon

`func (o *CreateEntitlementTypeRequest) HasRecon() bool`

HasRecon returns a boolean if a field has been set.

### GetShowonchild

`func (o *CreateEntitlementTypeRequest) GetShowonchild() bool`

GetShowonchild returns the Showonchild field if non-nil, zero value otherwise.

### GetShowonchildOk

`func (o *CreateEntitlementTypeRequest) GetShowonchildOk() (*bool, bool)`

GetShowonchildOk returns a tuple with the Showonchild field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowonchild

`func (o *CreateEntitlementTypeRequest) SetShowonchild(v bool)`

SetShowonchild sets Showonchild field to given value.

### HasShowonchild

`func (o *CreateEntitlementTypeRequest) HasShowonchild() bool`

HasShowonchild returns a boolean if a field has been set.

### GetHiearchyrequired

`func (o *CreateEntitlementTypeRequest) GetHiearchyrequired() int32`

GetHiearchyrequired returns the Hiearchyrequired field if non-nil, zero value otherwise.

### GetHiearchyrequiredOk

`func (o *CreateEntitlementTypeRequest) GetHiearchyrequiredOk() (*int32, bool)`

GetHiearchyrequiredOk returns a tuple with the Hiearchyrequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHiearchyrequired

`func (o *CreateEntitlementTypeRequest) SetHiearchyrequired(v int32)`

SetHiearchyrequired sets Hiearchyrequired field to given value.

### HasHiearchyrequired

`func (o *CreateEntitlementTypeRequest) HasHiearchyrequired() bool`

HasHiearchyrequired returns a boolean if a field has been set.

### GetEnableProvisioningPriority

`func (o *CreateEntitlementTypeRequest) GetEnableProvisioningPriority() bool`

GetEnableProvisioningPriority returns the EnableProvisioningPriority field if non-nil, zero value otherwise.

### GetEnableProvisioningPriorityOk

`func (o *CreateEntitlementTypeRequest) GetEnableProvisioningPriorityOk() (*bool, bool)`

GetEnableProvisioningPriorityOk returns a tuple with the EnableProvisioningPriority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableProvisioningPriority

`func (o *CreateEntitlementTypeRequest) SetEnableProvisioningPriority(v bool)`

SetEnableProvisioningPriority sets EnableProvisioningPriority field to given value.

### HasEnableProvisioningPriority

`func (o *CreateEntitlementTypeRequest) HasEnableProvisioningPriority() bool`

HasEnableProvisioningPriority returns a boolean if a field has been set.

### GetExcludeRuleAssgnEntsInRequest

`func (o *CreateEntitlementTypeRequest) GetExcludeRuleAssgnEntsInRequest() bool`

GetExcludeRuleAssgnEntsInRequest returns the ExcludeRuleAssgnEntsInRequest field if non-nil, zero value otherwise.

### GetExcludeRuleAssgnEntsInRequestOk

`func (o *CreateEntitlementTypeRequest) GetExcludeRuleAssgnEntsInRequestOk() (*bool, bool)`

GetExcludeRuleAssgnEntsInRequestOk returns a tuple with the ExcludeRuleAssgnEntsInRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeRuleAssgnEntsInRequest

`func (o *CreateEntitlementTypeRequest) SetExcludeRuleAssgnEntsInRequest(v bool)`

SetExcludeRuleAssgnEntsInRequest sets ExcludeRuleAssgnEntsInRequest field to given value.

### HasExcludeRuleAssgnEntsInRequest

`func (o *CreateEntitlementTypeRequest) HasExcludeRuleAssgnEntsInRequest() bool`

HasExcludeRuleAssgnEntsInRequest returns a boolean if a field has been set.

### GetCustomproperty1

`func (o *CreateEntitlementTypeRequest) GetCustomproperty1() string`

GetCustomproperty1 returns the Customproperty1 field if non-nil, zero value otherwise.

### GetCustomproperty1Ok

`func (o *CreateEntitlementTypeRequest) GetCustomproperty1Ok() (*string, bool)`

GetCustomproperty1Ok returns a tuple with the Customproperty1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty1

`func (o *CreateEntitlementTypeRequest) SetCustomproperty1(v string)`

SetCustomproperty1 sets Customproperty1 field to given value.

### HasCustomproperty1

`func (o *CreateEntitlementTypeRequest) HasCustomproperty1() bool`

HasCustomproperty1 returns a boolean if a field has been set.

### GetCustomproperty2

`func (o *CreateEntitlementTypeRequest) GetCustomproperty2() string`

GetCustomproperty2 returns the Customproperty2 field if non-nil, zero value otherwise.

### GetCustomproperty2Ok

`func (o *CreateEntitlementTypeRequest) GetCustomproperty2Ok() (*string, bool)`

GetCustomproperty2Ok returns a tuple with the Customproperty2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty2

`func (o *CreateEntitlementTypeRequest) SetCustomproperty2(v string)`

SetCustomproperty2 sets Customproperty2 field to given value.

### HasCustomproperty2

`func (o *CreateEntitlementTypeRequest) HasCustomproperty2() bool`

HasCustomproperty2 returns a boolean if a field has been set.

### GetCustomproperty3

`func (o *CreateEntitlementTypeRequest) GetCustomproperty3() string`

GetCustomproperty3 returns the Customproperty3 field if non-nil, zero value otherwise.

### GetCustomproperty3Ok

`func (o *CreateEntitlementTypeRequest) GetCustomproperty3Ok() (*string, bool)`

GetCustomproperty3Ok returns a tuple with the Customproperty3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty3

`func (o *CreateEntitlementTypeRequest) SetCustomproperty3(v string)`

SetCustomproperty3 sets Customproperty3 field to given value.

### HasCustomproperty3

`func (o *CreateEntitlementTypeRequest) HasCustomproperty3() bool`

HasCustomproperty3 returns a boolean if a field has been set.

### GetCustomproperty4

`func (o *CreateEntitlementTypeRequest) GetCustomproperty4() string`

GetCustomproperty4 returns the Customproperty4 field if non-nil, zero value otherwise.

### GetCustomproperty4Ok

`func (o *CreateEntitlementTypeRequest) GetCustomproperty4Ok() (*string, bool)`

GetCustomproperty4Ok returns a tuple with the Customproperty4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty4

`func (o *CreateEntitlementTypeRequest) SetCustomproperty4(v string)`

SetCustomproperty4 sets Customproperty4 field to given value.

### HasCustomproperty4

`func (o *CreateEntitlementTypeRequest) HasCustomproperty4() bool`

HasCustomproperty4 returns a boolean if a field has been set.

### GetCustomproperty5

`func (o *CreateEntitlementTypeRequest) GetCustomproperty5() string`

GetCustomproperty5 returns the Customproperty5 field if non-nil, zero value otherwise.

### GetCustomproperty5Ok

`func (o *CreateEntitlementTypeRequest) GetCustomproperty5Ok() (*string, bool)`

GetCustomproperty5Ok returns a tuple with the Customproperty5 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty5

`func (o *CreateEntitlementTypeRequest) SetCustomproperty5(v string)`

SetCustomproperty5 sets Customproperty5 field to given value.

### HasCustomproperty5

`func (o *CreateEntitlementTypeRequest) HasCustomproperty5() bool`

HasCustomproperty5 returns a boolean if a field has been set.

### GetCustomproperty6

`func (o *CreateEntitlementTypeRequest) GetCustomproperty6() string`

GetCustomproperty6 returns the Customproperty6 field if non-nil, zero value otherwise.

### GetCustomproperty6Ok

`func (o *CreateEntitlementTypeRequest) GetCustomproperty6Ok() (*string, bool)`

GetCustomproperty6Ok returns a tuple with the Customproperty6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty6

`func (o *CreateEntitlementTypeRequest) SetCustomproperty6(v string)`

SetCustomproperty6 sets Customproperty6 field to given value.

### HasCustomproperty6

`func (o *CreateEntitlementTypeRequest) HasCustomproperty6() bool`

HasCustomproperty6 returns a boolean if a field has been set.

### GetCustomproperty7

`func (o *CreateEntitlementTypeRequest) GetCustomproperty7() string`

GetCustomproperty7 returns the Customproperty7 field if non-nil, zero value otherwise.

### GetCustomproperty7Ok

`func (o *CreateEntitlementTypeRequest) GetCustomproperty7Ok() (*string, bool)`

GetCustomproperty7Ok returns a tuple with the Customproperty7 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty7

`func (o *CreateEntitlementTypeRequest) SetCustomproperty7(v string)`

SetCustomproperty7 sets Customproperty7 field to given value.

### HasCustomproperty7

`func (o *CreateEntitlementTypeRequest) HasCustomproperty7() bool`

HasCustomproperty7 returns a boolean if a field has been set.

### GetCustomproperty8

`func (o *CreateEntitlementTypeRequest) GetCustomproperty8() string`

GetCustomproperty8 returns the Customproperty8 field if non-nil, zero value otherwise.

### GetCustomproperty8Ok

`func (o *CreateEntitlementTypeRequest) GetCustomproperty8Ok() (*string, bool)`

GetCustomproperty8Ok returns a tuple with the Customproperty8 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty8

`func (o *CreateEntitlementTypeRequest) SetCustomproperty8(v string)`

SetCustomproperty8 sets Customproperty8 field to given value.

### HasCustomproperty8

`func (o *CreateEntitlementTypeRequest) HasCustomproperty8() bool`

HasCustomproperty8 returns a boolean if a field has been set.

### GetCustomproperty9

`func (o *CreateEntitlementTypeRequest) GetCustomproperty9() string`

GetCustomproperty9 returns the Customproperty9 field if non-nil, zero value otherwise.

### GetCustomproperty9Ok

`func (o *CreateEntitlementTypeRequest) GetCustomproperty9Ok() (*string, bool)`

GetCustomproperty9Ok returns a tuple with the Customproperty9 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty9

`func (o *CreateEntitlementTypeRequest) SetCustomproperty9(v string)`

SetCustomproperty9 sets Customproperty9 field to given value.

### HasCustomproperty9

`func (o *CreateEntitlementTypeRequest) HasCustomproperty9() bool`

HasCustomproperty9 returns a boolean if a field has been set.

### GetCustomproperty10

`func (o *CreateEntitlementTypeRequest) GetCustomproperty10() string`

GetCustomproperty10 returns the Customproperty10 field if non-nil, zero value otherwise.

### GetCustomproperty10Ok

`func (o *CreateEntitlementTypeRequest) GetCustomproperty10Ok() (*string, bool)`

GetCustomproperty10Ok returns a tuple with the Customproperty10 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty10

`func (o *CreateEntitlementTypeRequest) SetCustomproperty10(v string)`

SetCustomproperty10 sets Customproperty10 field to given value.

### HasCustomproperty10

`func (o *CreateEntitlementTypeRequest) HasCustomproperty10() bool`

HasCustomproperty10 returns a boolean if a field has been set.

### GetCustomproperty11

`func (o *CreateEntitlementTypeRequest) GetCustomproperty11() string`

GetCustomproperty11 returns the Customproperty11 field if non-nil, zero value otherwise.

### GetCustomproperty11Ok

`func (o *CreateEntitlementTypeRequest) GetCustomproperty11Ok() (*string, bool)`

GetCustomproperty11Ok returns a tuple with the Customproperty11 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty11

`func (o *CreateEntitlementTypeRequest) SetCustomproperty11(v string)`

SetCustomproperty11 sets Customproperty11 field to given value.

### HasCustomproperty11

`func (o *CreateEntitlementTypeRequest) HasCustomproperty11() bool`

HasCustomproperty11 returns a boolean if a field has been set.

### GetCustomproperty12

`func (o *CreateEntitlementTypeRequest) GetCustomproperty12() string`

GetCustomproperty12 returns the Customproperty12 field if non-nil, zero value otherwise.

### GetCustomproperty12Ok

`func (o *CreateEntitlementTypeRequest) GetCustomproperty12Ok() (*string, bool)`

GetCustomproperty12Ok returns a tuple with the Customproperty12 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty12

`func (o *CreateEntitlementTypeRequest) SetCustomproperty12(v string)`

SetCustomproperty12 sets Customproperty12 field to given value.

### HasCustomproperty12

`func (o *CreateEntitlementTypeRequest) HasCustomproperty12() bool`

HasCustomproperty12 returns a boolean if a field has been set.

### GetCustomproperty13

`func (o *CreateEntitlementTypeRequest) GetCustomproperty13() string`

GetCustomproperty13 returns the Customproperty13 field if non-nil, zero value otherwise.

### GetCustomproperty13Ok

`func (o *CreateEntitlementTypeRequest) GetCustomproperty13Ok() (*string, bool)`

GetCustomproperty13Ok returns a tuple with the Customproperty13 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty13

`func (o *CreateEntitlementTypeRequest) SetCustomproperty13(v string)`

SetCustomproperty13 sets Customproperty13 field to given value.

### HasCustomproperty13

`func (o *CreateEntitlementTypeRequest) HasCustomproperty13() bool`

HasCustomproperty13 returns a boolean if a field has been set.

### GetCustomproperty14

`func (o *CreateEntitlementTypeRequest) GetCustomproperty14() string`

GetCustomproperty14 returns the Customproperty14 field if non-nil, zero value otherwise.

### GetCustomproperty14Ok

`func (o *CreateEntitlementTypeRequest) GetCustomproperty14Ok() (*string, bool)`

GetCustomproperty14Ok returns a tuple with the Customproperty14 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty14

`func (o *CreateEntitlementTypeRequest) SetCustomproperty14(v string)`

SetCustomproperty14 sets Customproperty14 field to given value.

### HasCustomproperty14

`func (o *CreateEntitlementTypeRequest) HasCustomproperty14() bool`

HasCustomproperty14 returns a boolean if a field has been set.

### GetCustomproperty15

`func (o *CreateEntitlementTypeRequest) GetCustomproperty15() string`

GetCustomproperty15 returns the Customproperty15 field if non-nil, zero value otherwise.

### GetCustomproperty15Ok

`func (o *CreateEntitlementTypeRequest) GetCustomproperty15Ok() (*string, bool)`

GetCustomproperty15Ok returns a tuple with the Customproperty15 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty15

`func (o *CreateEntitlementTypeRequest) SetCustomproperty15(v string)`

SetCustomproperty15 sets Customproperty15 field to given value.

### HasCustomproperty15

`func (o *CreateEntitlementTypeRequest) HasCustomproperty15() bool`

HasCustomproperty15 returns a boolean if a field has been set.

### GetCustomproperty16

`func (o *CreateEntitlementTypeRequest) GetCustomproperty16() string`

GetCustomproperty16 returns the Customproperty16 field if non-nil, zero value otherwise.

### GetCustomproperty16Ok

`func (o *CreateEntitlementTypeRequest) GetCustomproperty16Ok() (*string, bool)`

GetCustomproperty16Ok returns a tuple with the Customproperty16 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty16

`func (o *CreateEntitlementTypeRequest) SetCustomproperty16(v string)`

SetCustomproperty16 sets Customproperty16 field to given value.

### HasCustomproperty16

`func (o *CreateEntitlementTypeRequest) HasCustomproperty16() bool`

HasCustomproperty16 returns a boolean if a field has been set.

### GetCustomproperty17

`func (o *CreateEntitlementTypeRequest) GetCustomproperty17() string`

GetCustomproperty17 returns the Customproperty17 field if non-nil, zero value otherwise.

### GetCustomproperty17Ok

`func (o *CreateEntitlementTypeRequest) GetCustomproperty17Ok() (*string, bool)`

GetCustomproperty17Ok returns a tuple with the Customproperty17 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty17

`func (o *CreateEntitlementTypeRequest) SetCustomproperty17(v string)`

SetCustomproperty17 sets Customproperty17 field to given value.

### HasCustomproperty17

`func (o *CreateEntitlementTypeRequest) HasCustomproperty17() bool`

HasCustomproperty17 returns a boolean if a field has been set.

### GetCustomproperty18

`func (o *CreateEntitlementTypeRequest) GetCustomproperty18() string`

GetCustomproperty18 returns the Customproperty18 field if non-nil, zero value otherwise.

### GetCustomproperty18Ok

`func (o *CreateEntitlementTypeRequest) GetCustomproperty18Ok() (*string, bool)`

GetCustomproperty18Ok returns a tuple with the Customproperty18 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty18

`func (o *CreateEntitlementTypeRequest) SetCustomproperty18(v string)`

SetCustomproperty18 sets Customproperty18 field to given value.

### HasCustomproperty18

`func (o *CreateEntitlementTypeRequest) HasCustomproperty18() bool`

HasCustomproperty18 returns a boolean if a field has been set.

### GetCustomproperty19

`func (o *CreateEntitlementTypeRequest) GetCustomproperty19() string`

GetCustomproperty19 returns the Customproperty19 field if non-nil, zero value otherwise.

### GetCustomproperty19Ok

`func (o *CreateEntitlementTypeRequest) GetCustomproperty19Ok() (*string, bool)`

GetCustomproperty19Ok returns a tuple with the Customproperty19 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty19

`func (o *CreateEntitlementTypeRequest) SetCustomproperty19(v string)`

SetCustomproperty19 sets Customproperty19 field to given value.

### HasCustomproperty19

`func (o *CreateEntitlementTypeRequest) HasCustomproperty19() bool`

HasCustomproperty19 returns a boolean if a field has been set.

### GetCustomproperty20

`func (o *CreateEntitlementTypeRequest) GetCustomproperty20() string`

GetCustomproperty20 returns the Customproperty20 field if non-nil, zero value otherwise.

### GetCustomproperty20Ok

`func (o *CreateEntitlementTypeRequest) GetCustomproperty20Ok() (*string, bool)`

GetCustomproperty20Ok returns a tuple with the Customproperty20 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty20

`func (o *CreateEntitlementTypeRequest) SetCustomproperty20(v string)`

SetCustomproperty20 sets Customproperty20 field to given value.

### HasCustomproperty20

`func (o *CreateEntitlementTypeRequest) HasCustomproperty20() bool`

HasCustomproperty20 returns a boolean if a field has been set.

### GetCustomproperty21

`func (o *CreateEntitlementTypeRequest) GetCustomproperty21() string`

GetCustomproperty21 returns the Customproperty21 field if non-nil, zero value otherwise.

### GetCustomproperty21Ok

`func (o *CreateEntitlementTypeRequest) GetCustomproperty21Ok() (*string, bool)`

GetCustomproperty21Ok returns a tuple with the Customproperty21 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty21

`func (o *CreateEntitlementTypeRequest) SetCustomproperty21(v string)`

SetCustomproperty21 sets Customproperty21 field to given value.

### HasCustomproperty21

`func (o *CreateEntitlementTypeRequest) HasCustomproperty21() bool`

HasCustomproperty21 returns a boolean if a field has been set.

### GetCustomproperty22

`func (o *CreateEntitlementTypeRequest) GetCustomproperty22() string`

GetCustomproperty22 returns the Customproperty22 field if non-nil, zero value otherwise.

### GetCustomproperty22Ok

`func (o *CreateEntitlementTypeRequest) GetCustomproperty22Ok() (*string, bool)`

GetCustomproperty22Ok returns a tuple with the Customproperty22 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty22

`func (o *CreateEntitlementTypeRequest) SetCustomproperty22(v string)`

SetCustomproperty22 sets Customproperty22 field to given value.

### HasCustomproperty22

`func (o *CreateEntitlementTypeRequest) HasCustomproperty22() bool`

HasCustomproperty22 returns a boolean if a field has been set.

### GetCustomproperty23

`func (o *CreateEntitlementTypeRequest) GetCustomproperty23() string`

GetCustomproperty23 returns the Customproperty23 field if non-nil, zero value otherwise.

### GetCustomproperty23Ok

`func (o *CreateEntitlementTypeRequest) GetCustomproperty23Ok() (*string, bool)`

GetCustomproperty23Ok returns a tuple with the Customproperty23 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty23

`func (o *CreateEntitlementTypeRequest) SetCustomproperty23(v string)`

SetCustomproperty23 sets Customproperty23 field to given value.

### HasCustomproperty23

`func (o *CreateEntitlementTypeRequest) HasCustomproperty23() bool`

HasCustomproperty23 returns a boolean if a field has been set.

### GetCustomproperty24

`func (o *CreateEntitlementTypeRequest) GetCustomproperty24() string`

GetCustomproperty24 returns the Customproperty24 field if non-nil, zero value otherwise.

### GetCustomproperty24Ok

`func (o *CreateEntitlementTypeRequest) GetCustomproperty24Ok() (*string, bool)`

GetCustomproperty24Ok returns a tuple with the Customproperty24 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty24

`func (o *CreateEntitlementTypeRequest) SetCustomproperty24(v string)`

SetCustomproperty24 sets Customproperty24 field to given value.

### HasCustomproperty24

`func (o *CreateEntitlementTypeRequest) HasCustomproperty24() bool`

HasCustomproperty24 returns a boolean if a field has been set.

### GetCustomproperty25

`func (o *CreateEntitlementTypeRequest) GetCustomproperty25() string`

GetCustomproperty25 returns the Customproperty25 field if non-nil, zero value otherwise.

### GetCustomproperty25Ok

`func (o *CreateEntitlementTypeRequest) GetCustomproperty25Ok() (*string, bool)`

GetCustomproperty25Ok returns a tuple with the Customproperty25 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty25

`func (o *CreateEntitlementTypeRequest) SetCustomproperty25(v string)`

SetCustomproperty25 sets Customproperty25 field to given value.

### HasCustomproperty25

`func (o *CreateEntitlementTypeRequest) HasCustomproperty25() bool`

HasCustomproperty25 returns a boolean if a field has been set.

### GetCustomproperty26

`func (o *CreateEntitlementTypeRequest) GetCustomproperty26() string`

GetCustomproperty26 returns the Customproperty26 field if non-nil, zero value otherwise.

### GetCustomproperty26Ok

`func (o *CreateEntitlementTypeRequest) GetCustomproperty26Ok() (*string, bool)`

GetCustomproperty26Ok returns a tuple with the Customproperty26 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty26

`func (o *CreateEntitlementTypeRequest) SetCustomproperty26(v string)`

SetCustomproperty26 sets Customproperty26 field to given value.

### HasCustomproperty26

`func (o *CreateEntitlementTypeRequest) HasCustomproperty26() bool`

HasCustomproperty26 returns a boolean if a field has been set.

### GetCustomproperty27

`func (o *CreateEntitlementTypeRequest) GetCustomproperty27() string`

GetCustomproperty27 returns the Customproperty27 field if non-nil, zero value otherwise.

### GetCustomproperty27Ok

`func (o *CreateEntitlementTypeRequest) GetCustomproperty27Ok() (*string, bool)`

GetCustomproperty27Ok returns a tuple with the Customproperty27 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty27

`func (o *CreateEntitlementTypeRequest) SetCustomproperty27(v string)`

SetCustomproperty27 sets Customproperty27 field to given value.

### HasCustomproperty27

`func (o *CreateEntitlementTypeRequest) HasCustomproperty27() bool`

HasCustomproperty27 returns a boolean if a field has been set.

### GetCustomproperty28

`func (o *CreateEntitlementTypeRequest) GetCustomproperty28() string`

GetCustomproperty28 returns the Customproperty28 field if non-nil, zero value otherwise.

### GetCustomproperty28Ok

`func (o *CreateEntitlementTypeRequest) GetCustomproperty28Ok() (*string, bool)`

GetCustomproperty28Ok returns a tuple with the Customproperty28 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty28

`func (o *CreateEntitlementTypeRequest) SetCustomproperty28(v string)`

SetCustomproperty28 sets Customproperty28 field to given value.

### HasCustomproperty28

`func (o *CreateEntitlementTypeRequest) HasCustomproperty28() bool`

HasCustomproperty28 returns a boolean if a field has been set.

### GetCustomproperty29

`func (o *CreateEntitlementTypeRequest) GetCustomproperty29() string`

GetCustomproperty29 returns the Customproperty29 field if non-nil, zero value otherwise.

### GetCustomproperty29Ok

`func (o *CreateEntitlementTypeRequest) GetCustomproperty29Ok() (*string, bool)`

GetCustomproperty29Ok returns a tuple with the Customproperty29 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty29

`func (o *CreateEntitlementTypeRequest) SetCustomproperty29(v string)`

SetCustomproperty29 sets Customproperty29 field to given value.

### HasCustomproperty29

`func (o *CreateEntitlementTypeRequest) HasCustomproperty29() bool`

HasCustomproperty29 returns a boolean if a field has been set.

### GetCustomproperty30

`func (o *CreateEntitlementTypeRequest) GetCustomproperty30() string`

GetCustomproperty30 returns the Customproperty30 field if non-nil, zero value otherwise.

### GetCustomproperty30Ok

`func (o *CreateEntitlementTypeRequest) GetCustomproperty30Ok() (*string, bool)`

GetCustomproperty30Ok returns a tuple with the Customproperty30 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty30

`func (o *CreateEntitlementTypeRequest) SetCustomproperty30(v string)`

SetCustomproperty30 sets Customproperty30 field to given value.

### HasCustomproperty30

`func (o *CreateEntitlementTypeRequest) HasCustomproperty30() bool`

HasCustomproperty30 returns a boolean if a field has been set.

### GetCustomproperty31

`func (o *CreateEntitlementTypeRequest) GetCustomproperty31() string`

GetCustomproperty31 returns the Customproperty31 field if non-nil, zero value otherwise.

### GetCustomproperty31Ok

`func (o *CreateEntitlementTypeRequest) GetCustomproperty31Ok() (*string, bool)`

GetCustomproperty31Ok returns a tuple with the Customproperty31 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty31

`func (o *CreateEntitlementTypeRequest) SetCustomproperty31(v string)`

SetCustomproperty31 sets Customproperty31 field to given value.

### HasCustomproperty31

`func (o *CreateEntitlementTypeRequest) HasCustomproperty31() bool`

HasCustomproperty31 returns a boolean if a field has been set.

### GetCustomproperty32

`func (o *CreateEntitlementTypeRequest) GetCustomproperty32() string`

GetCustomproperty32 returns the Customproperty32 field if non-nil, zero value otherwise.

### GetCustomproperty32Ok

`func (o *CreateEntitlementTypeRequest) GetCustomproperty32Ok() (*string, bool)`

GetCustomproperty32Ok returns a tuple with the Customproperty32 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty32

`func (o *CreateEntitlementTypeRequest) SetCustomproperty32(v string)`

SetCustomproperty32 sets Customproperty32 field to given value.

### HasCustomproperty32

`func (o *CreateEntitlementTypeRequest) HasCustomproperty32() bool`

HasCustomproperty32 returns a boolean if a field has been set.

### GetCustomproperty33

`func (o *CreateEntitlementTypeRequest) GetCustomproperty33() string`

GetCustomproperty33 returns the Customproperty33 field if non-nil, zero value otherwise.

### GetCustomproperty33Ok

`func (o *CreateEntitlementTypeRequest) GetCustomproperty33Ok() (*string, bool)`

GetCustomproperty33Ok returns a tuple with the Customproperty33 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty33

`func (o *CreateEntitlementTypeRequest) SetCustomproperty33(v string)`

SetCustomproperty33 sets Customproperty33 field to given value.

### HasCustomproperty33

`func (o *CreateEntitlementTypeRequest) HasCustomproperty33() bool`

HasCustomproperty33 returns a boolean if a field has been set.

### GetCustomproperty34

`func (o *CreateEntitlementTypeRequest) GetCustomproperty34() string`

GetCustomproperty34 returns the Customproperty34 field if non-nil, zero value otherwise.

### GetCustomproperty34Ok

`func (o *CreateEntitlementTypeRequest) GetCustomproperty34Ok() (*string, bool)`

GetCustomproperty34Ok returns a tuple with the Customproperty34 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty34

`func (o *CreateEntitlementTypeRequest) SetCustomproperty34(v string)`

SetCustomproperty34 sets Customproperty34 field to given value.

### HasCustomproperty34

`func (o *CreateEntitlementTypeRequest) HasCustomproperty34() bool`

HasCustomproperty34 returns a boolean if a field has been set.

### GetCustomproperty35

`func (o *CreateEntitlementTypeRequest) GetCustomproperty35() string`

GetCustomproperty35 returns the Customproperty35 field if non-nil, zero value otherwise.

### GetCustomproperty35Ok

`func (o *CreateEntitlementTypeRequest) GetCustomproperty35Ok() (*string, bool)`

GetCustomproperty35Ok returns a tuple with the Customproperty35 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty35

`func (o *CreateEntitlementTypeRequest) SetCustomproperty35(v string)`

SetCustomproperty35 sets Customproperty35 field to given value.

### HasCustomproperty35

`func (o *CreateEntitlementTypeRequest) HasCustomproperty35() bool`

HasCustomproperty35 returns a boolean if a field has been set.

### GetCustomproperty36

`func (o *CreateEntitlementTypeRequest) GetCustomproperty36() string`

GetCustomproperty36 returns the Customproperty36 field if non-nil, zero value otherwise.

### GetCustomproperty36Ok

`func (o *CreateEntitlementTypeRequest) GetCustomproperty36Ok() (*string, bool)`

GetCustomproperty36Ok returns a tuple with the Customproperty36 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty36

`func (o *CreateEntitlementTypeRequest) SetCustomproperty36(v string)`

SetCustomproperty36 sets Customproperty36 field to given value.

### HasCustomproperty36

`func (o *CreateEntitlementTypeRequest) HasCustomproperty36() bool`

HasCustomproperty36 returns a boolean if a field has been set.

### GetCustomproperty37

`func (o *CreateEntitlementTypeRequest) GetCustomproperty37() string`

GetCustomproperty37 returns the Customproperty37 field if non-nil, zero value otherwise.

### GetCustomproperty37Ok

`func (o *CreateEntitlementTypeRequest) GetCustomproperty37Ok() (*string, bool)`

GetCustomproperty37Ok returns a tuple with the Customproperty37 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty37

`func (o *CreateEntitlementTypeRequest) SetCustomproperty37(v string)`

SetCustomproperty37 sets Customproperty37 field to given value.

### HasCustomproperty37

`func (o *CreateEntitlementTypeRequest) HasCustomproperty37() bool`

HasCustomproperty37 returns a boolean if a field has been set.

### GetCustomproperty38

`func (o *CreateEntitlementTypeRequest) GetCustomproperty38() string`

GetCustomproperty38 returns the Customproperty38 field if non-nil, zero value otherwise.

### GetCustomproperty38Ok

`func (o *CreateEntitlementTypeRequest) GetCustomproperty38Ok() (*string, bool)`

GetCustomproperty38Ok returns a tuple with the Customproperty38 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty38

`func (o *CreateEntitlementTypeRequest) SetCustomproperty38(v string)`

SetCustomproperty38 sets Customproperty38 field to given value.

### HasCustomproperty38

`func (o *CreateEntitlementTypeRequest) HasCustomproperty38() bool`

HasCustomproperty38 returns a boolean if a field has been set.

### GetCustomproperty39

`func (o *CreateEntitlementTypeRequest) GetCustomproperty39() string`

GetCustomproperty39 returns the Customproperty39 field if non-nil, zero value otherwise.

### GetCustomproperty39Ok

`func (o *CreateEntitlementTypeRequest) GetCustomproperty39Ok() (*string, bool)`

GetCustomproperty39Ok returns a tuple with the Customproperty39 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty39

`func (o *CreateEntitlementTypeRequest) SetCustomproperty39(v string)`

SetCustomproperty39 sets Customproperty39 field to given value.

### HasCustomproperty39

`func (o *CreateEntitlementTypeRequest) HasCustomproperty39() bool`

HasCustomproperty39 returns a boolean if a field has been set.

### GetCustomproperty40

`func (o *CreateEntitlementTypeRequest) GetCustomproperty40() string`

GetCustomproperty40 returns the Customproperty40 field if non-nil, zero value otherwise.

### GetCustomproperty40Ok

`func (o *CreateEntitlementTypeRequest) GetCustomproperty40Ok() (*string, bool)`

GetCustomproperty40Ok returns a tuple with the Customproperty40 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty40

`func (o *CreateEntitlementTypeRequest) SetCustomproperty40(v string)`

SetCustomproperty40 sets Customproperty40 field to given value.

### HasCustomproperty40

`func (o *CreateEntitlementTypeRequest) HasCustomproperty40() bool`

HasCustomproperty40 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


