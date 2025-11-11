# CreateEnterpriseRoleRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Roletype** | **string** | Use the roletype as &#39;ENTERPRISE&#39; to Create Enterprise Role Request. | 
**RoleName** | **string** |  | 
**Requestor** | **string** |  | 
**Owner** | [**[]CreateRoleOwnerPayload**](CreateRoleOwnerPayload.md) |  | 
**Customproperty1** | Pointer to **string** |  | [optional] 
**Customproperty2** | Pointer to **string** |  | [optional] 
**Customproperty3** | Pointer to **string** |  | [optional] 
**Customproperty4** | Pointer to **string** |  | [optional] 
**Customproperty5** | Pointer to **string** |  | [optional] 
**Customproperty6** | Pointer to **string** |  | [optional] 
**Customproperty7** | Pointer to **string** |  | [optional] 
**Customproperty8** | Pointer to **string** |  | [optional] 
**Customproperty9** | Pointer to **string** |  | [optional] 
**Customproperty10** | Pointer to **string** |  | [optional] 
**Customproperty11** | Pointer to **string** |  | [optional] 
**Customproperty12** | Pointer to **string** |  | [optional] 
**Customproperty13** | Pointer to **string** |  | [optional] 
**Customproperty14** | Pointer to **string** |  | [optional] 
**Customproperty15** | Pointer to **string** |  | [optional] 
**Customproperty16** | Pointer to **string** |  | [optional] 
**Customproperty17** | Pointer to **string** |  | [optional] 
**Customproperty18** | Pointer to **string** |  | [optional] 
**Customproperty19** | Pointer to **string** |  | [optional] 
**Customproperty20** | Pointer to **string** |  | [optional] 
**Customproperty21** | Pointer to **string** |  | [optional] 
**Customproperty22** | Pointer to **string** |  | [optional] 
**Customproperty23** | Pointer to **string** |  | [optional] 
**Customproperty24** | Pointer to **string** |  | [optional] 
**Customproperty25** | Pointer to **string** |  | [optional] 
**Customproperty26** | Pointer to **string** |  | [optional] 
**Customproperty27** | Pointer to **string** |  | [optional] 
**Customproperty28** | Pointer to **string** |  | [optional] 
**Customproperty29** | Pointer to **string** |  | [optional] 
**Customproperty30** | Pointer to **string** |  | [optional] 
**Customproperty31** | Pointer to **string** |  | [optional] 
**Customproperty32** | Pointer to **string** |  | [optional] 
**Customproperty33** | Pointer to **string** |  | [optional] 
**Customproperty34** | Pointer to **string** |  | [optional] 
**Customproperty35** | Pointer to **string** |  | [optional] 
**Customproperty36** | Pointer to **string** |  | [optional] 
**Customproperty37** | Pointer to **string** |  | [optional] 
**Customproperty38** | Pointer to **string** |  | [optional] 
**Customproperty39** | Pointer to **string** |  | [optional] 
**Customproperty40** | Pointer to **string** |  | [optional] 
**Customproperty41** | Pointer to **string** |  | [optional] 
**Customproperty42** | Pointer to **string** |  | [optional] 
**Customproperty43** | Pointer to **string** |  | [optional] 
**Customproperty44** | Pointer to **string** |  | [optional] 
**Customproperty45** | Pointer to **string** |  | [optional] 
**Customproperty46** | Pointer to **string** |  | [optional] 
**Customproperty47** | Pointer to **string** |  | [optional] 
**Customproperty48** | Pointer to **string** |  | [optional] 
**Customproperty49** | Pointer to **string** |  | [optional] 
**Customproperty50** | Pointer to **string** |  | [optional] 
**Customproperty51** | Pointer to **string** |  | [optional] 
**Customproperty52** | Pointer to **string** |  | [optional] 
**Customproperty53** | Pointer to **string** |  | [optional] 
**Customproperty54** | Pointer to **string** |  | [optional] 
**Customproperty55** | Pointer to **string** |  | [optional] 
**Customproperty56** | Pointer to **string** |  | [optional] 
**Customproperty57** | Pointer to **string** |  | [optional] 
**Customproperty58** | Pointer to **string** |  | [optional] 
**Customproperty59** | Pointer to **string** |  | [optional] 
**Customproperty60** | Pointer to **string** |  | [optional] 
**Endpointname** | Pointer to **string** |  | [optional] 
**Defaulttimeframe** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Displayname** | Pointer to **string** |  | [optional] 
**Glossary** | Pointer to **string** |  | [optional] 
**Risk** | Pointer to **string** |  | [optional] 
**Level** | Pointer to **string** |  | [optional] 
**Soxcritical** | Pointer to **string** |  | [optional] 
**Syscritical** | Pointer to **string** |  | [optional] 
**Priviliged** | Pointer to **string** |  | [optional] 
**Confidentiality** | Pointer to **string** |  | [optional] 
**Requestable** | Pointer to **string** |  | [optional] 
**ShowDynamicAttrs** | Pointer to **string** |  | [optional] 
**Checksod** | Pointer to **string** |  | [optional] 

## Methods

### NewCreateEnterpriseRoleRequest

`func NewCreateEnterpriseRoleRequest(roletype string, roleName string, requestor string, owner []CreateRoleOwnerPayload, ) *CreateEnterpriseRoleRequest`

NewCreateEnterpriseRoleRequest instantiates a new CreateEnterpriseRoleRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateEnterpriseRoleRequestWithDefaults

`func NewCreateEnterpriseRoleRequestWithDefaults() *CreateEnterpriseRoleRequest`

NewCreateEnterpriseRoleRequestWithDefaults instantiates a new CreateEnterpriseRoleRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRoletype

`func (o *CreateEnterpriseRoleRequest) GetRoletype() string`

GetRoletype returns the Roletype field if non-nil, zero value otherwise.

### GetRoletypeOk

`func (o *CreateEnterpriseRoleRequest) GetRoletypeOk() (*string, bool)`

GetRoletypeOk returns a tuple with the Roletype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoletype

`func (o *CreateEnterpriseRoleRequest) SetRoletype(v string)`

SetRoletype sets Roletype field to given value.


### GetRoleName

`func (o *CreateEnterpriseRoleRequest) GetRoleName() string`

GetRoleName returns the RoleName field if non-nil, zero value otherwise.

### GetRoleNameOk

`func (o *CreateEnterpriseRoleRequest) GetRoleNameOk() (*string, bool)`

GetRoleNameOk returns a tuple with the RoleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleName

`func (o *CreateEnterpriseRoleRequest) SetRoleName(v string)`

SetRoleName sets RoleName field to given value.


### GetRequestor

`func (o *CreateEnterpriseRoleRequest) GetRequestor() string`

GetRequestor returns the Requestor field if non-nil, zero value otherwise.

### GetRequestorOk

`func (o *CreateEnterpriseRoleRequest) GetRequestorOk() (*string, bool)`

GetRequestorOk returns a tuple with the Requestor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestor

`func (o *CreateEnterpriseRoleRequest) SetRequestor(v string)`

SetRequestor sets Requestor field to given value.


### GetOwner

`func (o *CreateEnterpriseRoleRequest) GetOwner() []CreateRoleOwnerPayload`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *CreateEnterpriseRoleRequest) GetOwnerOk() (*[]CreateRoleOwnerPayload, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *CreateEnterpriseRoleRequest) SetOwner(v []CreateRoleOwnerPayload)`

SetOwner sets Owner field to given value.


### GetCustomproperty1

`func (o *CreateEnterpriseRoleRequest) GetCustomproperty1() string`

GetCustomproperty1 returns the Customproperty1 field if non-nil, zero value otherwise.

### GetCustomproperty1Ok

`func (o *CreateEnterpriseRoleRequest) GetCustomproperty1Ok() (*string, bool)`

GetCustomproperty1Ok returns a tuple with the Customproperty1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty1

`func (o *CreateEnterpriseRoleRequest) SetCustomproperty1(v string)`

SetCustomproperty1 sets Customproperty1 field to given value.

### HasCustomproperty1

`func (o *CreateEnterpriseRoleRequest) HasCustomproperty1() bool`

HasCustomproperty1 returns a boolean if a field has been set.

### GetCustomproperty2

`func (o *CreateEnterpriseRoleRequest) GetCustomproperty2() string`

GetCustomproperty2 returns the Customproperty2 field if non-nil, zero value otherwise.

### GetCustomproperty2Ok

`func (o *CreateEnterpriseRoleRequest) GetCustomproperty2Ok() (*string, bool)`

GetCustomproperty2Ok returns a tuple with the Customproperty2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty2

`func (o *CreateEnterpriseRoleRequest) SetCustomproperty2(v string)`

SetCustomproperty2 sets Customproperty2 field to given value.

### HasCustomproperty2

`func (o *CreateEnterpriseRoleRequest) HasCustomproperty2() bool`

HasCustomproperty2 returns a boolean if a field has been set.

### GetCustomproperty3

`func (o *CreateEnterpriseRoleRequest) GetCustomproperty3() string`

GetCustomproperty3 returns the Customproperty3 field if non-nil, zero value otherwise.

### GetCustomproperty3Ok

`func (o *CreateEnterpriseRoleRequest) GetCustomproperty3Ok() (*string, bool)`

GetCustomproperty3Ok returns a tuple with the Customproperty3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty3

`func (o *CreateEnterpriseRoleRequest) SetCustomproperty3(v string)`

SetCustomproperty3 sets Customproperty3 field to given value.

### HasCustomproperty3

`func (o *CreateEnterpriseRoleRequest) HasCustomproperty3() bool`

HasCustomproperty3 returns a boolean if a field has been set.

### GetCustomproperty4

`func (o *CreateEnterpriseRoleRequest) GetCustomproperty4() string`

GetCustomproperty4 returns the Customproperty4 field if non-nil, zero value otherwise.

### GetCustomproperty4Ok

`func (o *CreateEnterpriseRoleRequest) GetCustomproperty4Ok() (*string, bool)`

GetCustomproperty4Ok returns a tuple with the Customproperty4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty4

`func (o *CreateEnterpriseRoleRequest) SetCustomproperty4(v string)`

SetCustomproperty4 sets Customproperty4 field to given value.

### HasCustomproperty4

`func (o *CreateEnterpriseRoleRequest) HasCustomproperty4() bool`

HasCustomproperty4 returns a boolean if a field has been set.

### GetCustomproperty5

`func (o *CreateEnterpriseRoleRequest) GetCustomproperty5() string`

GetCustomproperty5 returns the Customproperty5 field if non-nil, zero value otherwise.

### GetCustomproperty5Ok

`func (o *CreateEnterpriseRoleRequest) GetCustomproperty5Ok() (*string, bool)`

GetCustomproperty5Ok returns a tuple with the Customproperty5 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty5

`func (o *CreateEnterpriseRoleRequest) SetCustomproperty5(v string)`

SetCustomproperty5 sets Customproperty5 field to given value.

### HasCustomproperty5

`func (o *CreateEnterpriseRoleRequest) HasCustomproperty5() bool`

HasCustomproperty5 returns a boolean if a field has been set.

### GetCustomproperty6

`func (o *CreateEnterpriseRoleRequest) GetCustomproperty6() string`

GetCustomproperty6 returns the Customproperty6 field if non-nil, zero value otherwise.

### GetCustomproperty6Ok

`func (o *CreateEnterpriseRoleRequest) GetCustomproperty6Ok() (*string, bool)`

GetCustomproperty6Ok returns a tuple with the Customproperty6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty6

`func (o *CreateEnterpriseRoleRequest) SetCustomproperty6(v string)`

SetCustomproperty6 sets Customproperty6 field to given value.

### HasCustomproperty6

`func (o *CreateEnterpriseRoleRequest) HasCustomproperty6() bool`

HasCustomproperty6 returns a boolean if a field has been set.

### GetCustomproperty7

`func (o *CreateEnterpriseRoleRequest) GetCustomproperty7() string`

GetCustomproperty7 returns the Customproperty7 field if non-nil, zero value otherwise.

### GetCustomproperty7Ok

`func (o *CreateEnterpriseRoleRequest) GetCustomproperty7Ok() (*string, bool)`

GetCustomproperty7Ok returns a tuple with the Customproperty7 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty7

`func (o *CreateEnterpriseRoleRequest) SetCustomproperty7(v string)`

SetCustomproperty7 sets Customproperty7 field to given value.

### HasCustomproperty7

`func (o *CreateEnterpriseRoleRequest) HasCustomproperty7() bool`

HasCustomproperty7 returns a boolean if a field has been set.

### GetCustomproperty8

`func (o *CreateEnterpriseRoleRequest) GetCustomproperty8() string`

GetCustomproperty8 returns the Customproperty8 field if non-nil, zero value otherwise.

### GetCustomproperty8Ok

`func (o *CreateEnterpriseRoleRequest) GetCustomproperty8Ok() (*string, bool)`

GetCustomproperty8Ok returns a tuple with the Customproperty8 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty8

`func (o *CreateEnterpriseRoleRequest) SetCustomproperty8(v string)`

SetCustomproperty8 sets Customproperty8 field to given value.

### HasCustomproperty8

`func (o *CreateEnterpriseRoleRequest) HasCustomproperty8() bool`

HasCustomproperty8 returns a boolean if a field has been set.

### GetCustomproperty9

`func (o *CreateEnterpriseRoleRequest) GetCustomproperty9() string`

GetCustomproperty9 returns the Customproperty9 field if non-nil, zero value otherwise.

### GetCustomproperty9Ok

`func (o *CreateEnterpriseRoleRequest) GetCustomproperty9Ok() (*string, bool)`

GetCustomproperty9Ok returns a tuple with the Customproperty9 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty9

`func (o *CreateEnterpriseRoleRequest) SetCustomproperty9(v string)`

SetCustomproperty9 sets Customproperty9 field to given value.

### HasCustomproperty9

`func (o *CreateEnterpriseRoleRequest) HasCustomproperty9() bool`

HasCustomproperty9 returns a boolean if a field has been set.

### GetCustomproperty10

`func (o *CreateEnterpriseRoleRequest) GetCustomproperty10() string`

GetCustomproperty10 returns the Customproperty10 field if non-nil, zero value otherwise.

### GetCustomproperty10Ok

`func (o *CreateEnterpriseRoleRequest) GetCustomproperty10Ok() (*string, bool)`

GetCustomproperty10Ok returns a tuple with the Customproperty10 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty10

`func (o *CreateEnterpriseRoleRequest) SetCustomproperty10(v string)`

SetCustomproperty10 sets Customproperty10 field to given value.

### HasCustomproperty10

`func (o *CreateEnterpriseRoleRequest) HasCustomproperty10() bool`

HasCustomproperty10 returns a boolean if a field has been set.

### GetCustomproperty11

`func (o *CreateEnterpriseRoleRequest) GetCustomproperty11() string`

GetCustomproperty11 returns the Customproperty11 field if non-nil, zero value otherwise.

### GetCustomproperty11Ok

`func (o *CreateEnterpriseRoleRequest) GetCustomproperty11Ok() (*string, bool)`

GetCustomproperty11Ok returns a tuple with the Customproperty11 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty11

`func (o *CreateEnterpriseRoleRequest) SetCustomproperty11(v string)`

SetCustomproperty11 sets Customproperty11 field to given value.

### HasCustomproperty11

`func (o *CreateEnterpriseRoleRequest) HasCustomproperty11() bool`

HasCustomproperty11 returns a boolean if a field has been set.

### GetCustomproperty12

`func (o *CreateEnterpriseRoleRequest) GetCustomproperty12() string`

GetCustomproperty12 returns the Customproperty12 field if non-nil, zero value otherwise.

### GetCustomproperty12Ok

`func (o *CreateEnterpriseRoleRequest) GetCustomproperty12Ok() (*string, bool)`

GetCustomproperty12Ok returns a tuple with the Customproperty12 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty12

`func (o *CreateEnterpriseRoleRequest) SetCustomproperty12(v string)`

SetCustomproperty12 sets Customproperty12 field to given value.

### HasCustomproperty12

`func (o *CreateEnterpriseRoleRequest) HasCustomproperty12() bool`

HasCustomproperty12 returns a boolean if a field has been set.

### GetCustomproperty13

`func (o *CreateEnterpriseRoleRequest) GetCustomproperty13() string`

GetCustomproperty13 returns the Customproperty13 field if non-nil, zero value otherwise.

### GetCustomproperty13Ok

`func (o *CreateEnterpriseRoleRequest) GetCustomproperty13Ok() (*string, bool)`

GetCustomproperty13Ok returns a tuple with the Customproperty13 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty13

`func (o *CreateEnterpriseRoleRequest) SetCustomproperty13(v string)`

SetCustomproperty13 sets Customproperty13 field to given value.

### HasCustomproperty13

`func (o *CreateEnterpriseRoleRequest) HasCustomproperty13() bool`

HasCustomproperty13 returns a boolean if a field has been set.

### GetCustomproperty14

`func (o *CreateEnterpriseRoleRequest) GetCustomproperty14() string`

GetCustomproperty14 returns the Customproperty14 field if non-nil, zero value otherwise.

### GetCustomproperty14Ok

`func (o *CreateEnterpriseRoleRequest) GetCustomproperty14Ok() (*string, bool)`

GetCustomproperty14Ok returns a tuple with the Customproperty14 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty14

`func (o *CreateEnterpriseRoleRequest) SetCustomproperty14(v string)`

SetCustomproperty14 sets Customproperty14 field to given value.

### HasCustomproperty14

`func (o *CreateEnterpriseRoleRequest) HasCustomproperty14() bool`

HasCustomproperty14 returns a boolean if a field has been set.

### GetCustomproperty15

`func (o *CreateEnterpriseRoleRequest) GetCustomproperty15() string`

GetCustomproperty15 returns the Customproperty15 field if non-nil, zero value otherwise.

### GetCustomproperty15Ok

`func (o *CreateEnterpriseRoleRequest) GetCustomproperty15Ok() (*string, bool)`

GetCustomproperty15Ok returns a tuple with the Customproperty15 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty15

`func (o *CreateEnterpriseRoleRequest) SetCustomproperty15(v string)`

SetCustomproperty15 sets Customproperty15 field to given value.

### HasCustomproperty15

`func (o *CreateEnterpriseRoleRequest) HasCustomproperty15() bool`

HasCustomproperty15 returns a boolean if a field has been set.

### GetCustomproperty16

`func (o *CreateEnterpriseRoleRequest) GetCustomproperty16() string`

GetCustomproperty16 returns the Customproperty16 field if non-nil, zero value otherwise.

### GetCustomproperty16Ok

`func (o *CreateEnterpriseRoleRequest) GetCustomproperty16Ok() (*string, bool)`

GetCustomproperty16Ok returns a tuple with the Customproperty16 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty16

`func (o *CreateEnterpriseRoleRequest) SetCustomproperty16(v string)`

SetCustomproperty16 sets Customproperty16 field to given value.

### HasCustomproperty16

`func (o *CreateEnterpriseRoleRequest) HasCustomproperty16() bool`

HasCustomproperty16 returns a boolean if a field has been set.

### GetCustomproperty17

`func (o *CreateEnterpriseRoleRequest) GetCustomproperty17() string`

GetCustomproperty17 returns the Customproperty17 field if non-nil, zero value otherwise.

### GetCustomproperty17Ok

`func (o *CreateEnterpriseRoleRequest) GetCustomproperty17Ok() (*string, bool)`

GetCustomproperty17Ok returns a tuple with the Customproperty17 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty17

`func (o *CreateEnterpriseRoleRequest) SetCustomproperty17(v string)`

SetCustomproperty17 sets Customproperty17 field to given value.

### HasCustomproperty17

`func (o *CreateEnterpriseRoleRequest) HasCustomproperty17() bool`

HasCustomproperty17 returns a boolean if a field has been set.

### GetCustomproperty18

`func (o *CreateEnterpriseRoleRequest) GetCustomproperty18() string`

GetCustomproperty18 returns the Customproperty18 field if non-nil, zero value otherwise.

### GetCustomproperty18Ok

`func (o *CreateEnterpriseRoleRequest) GetCustomproperty18Ok() (*string, bool)`

GetCustomproperty18Ok returns a tuple with the Customproperty18 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty18

`func (o *CreateEnterpriseRoleRequest) SetCustomproperty18(v string)`

SetCustomproperty18 sets Customproperty18 field to given value.

### HasCustomproperty18

`func (o *CreateEnterpriseRoleRequest) HasCustomproperty18() bool`

HasCustomproperty18 returns a boolean if a field has been set.

### GetCustomproperty19

`func (o *CreateEnterpriseRoleRequest) GetCustomproperty19() string`

GetCustomproperty19 returns the Customproperty19 field if non-nil, zero value otherwise.

### GetCustomproperty19Ok

`func (o *CreateEnterpriseRoleRequest) GetCustomproperty19Ok() (*string, bool)`

GetCustomproperty19Ok returns a tuple with the Customproperty19 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty19

`func (o *CreateEnterpriseRoleRequest) SetCustomproperty19(v string)`

SetCustomproperty19 sets Customproperty19 field to given value.

### HasCustomproperty19

`func (o *CreateEnterpriseRoleRequest) HasCustomproperty19() bool`

HasCustomproperty19 returns a boolean if a field has been set.

### GetCustomproperty20

`func (o *CreateEnterpriseRoleRequest) GetCustomproperty20() string`

GetCustomproperty20 returns the Customproperty20 field if non-nil, zero value otherwise.

### GetCustomproperty20Ok

`func (o *CreateEnterpriseRoleRequest) GetCustomproperty20Ok() (*string, bool)`

GetCustomproperty20Ok returns a tuple with the Customproperty20 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty20

`func (o *CreateEnterpriseRoleRequest) SetCustomproperty20(v string)`

SetCustomproperty20 sets Customproperty20 field to given value.

### HasCustomproperty20

`func (o *CreateEnterpriseRoleRequest) HasCustomproperty20() bool`

HasCustomproperty20 returns a boolean if a field has been set.

### GetCustomproperty21

`func (o *CreateEnterpriseRoleRequest) GetCustomproperty21() string`

GetCustomproperty21 returns the Customproperty21 field if non-nil, zero value otherwise.

### GetCustomproperty21Ok

`func (o *CreateEnterpriseRoleRequest) GetCustomproperty21Ok() (*string, bool)`

GetCustomproperty21Ok returns a tuple with the Customproperty21 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty21

`func (o *CreateEnterpriseRoleRequest) SetCustomproperty21(v string)`

SetCustomproperty21 sets Customproperty21 field to given value.

### HasCustomproperty21

`func (o *CreateEnterpriseRoleRequest) HasCustomproperty21() bool`

HasCustomproperty21 returns a boolean if a field has been set.

### GetCustomproperty22

`func (o *CreateEnterpriseRoleRequest) GetCustomproperty22() string`

GetCustomproperty22 returns the Customproperty22 field if non-nil, zero value otherwise.

### GetCustomproperty22Ok

`func (o *CreateEnterpriseRoleRequest) GetCustomproperty22Ok() (*string, bool)`

GetCustomproperty22Ok returns a tuple with the Customproperty22 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty22

`func (o *CreateEnterpriseRoleRequest) SetCustomproperty22(v string)`

SetCustomproperty22 sets Customproperty22 field to given value.

### HasCustomproperty22

`func (o *CreateEnterpriseRoleRequest) HasCustomproperty22() bool`

HasCustomproperty22 returns a boolean if a field has been set.

### GetCustomproperty23

`func (o *CreateEnterpriseRoleRequest) GetCustomproperty23() string`

GetCustomproperty23 returns the Customproperty23 field if non-nil, zero value otherwise.

### GetCustomproperty23Ok

`func (o *CreateEnterpriseRoleRequest) GetCustomproperty23Ok() (*string, bool)`

GetCustomproperty23Ok returns a tuple with the Customproperty23 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty23

`func (o *CreateEnterpriseRoleRequest) SetCustomproperty23(v string)`

SetCustomproperty23 sets Customproperty23 field to given value.

### HasCustomproperty23

`func (o *CreateEnterpriseRoleRequest) HasCustomproperty23() bool`

HasCustomproperty23 returns a boolean if a field has been set.

### GetCustomproperty24

`func (o *CreateEnterpriseRoleRequest) GetCustomproperty24() string`

GetCustomproperty24 returns the Customproperty24 field if non-nil, zero value otherwise.

### GetCustomproperty24Ok

`func (o *CreateEnterpriseRoleRequest) GetCustomproperty24Ok() (*string, bool)`

GetCustomproperty24Ok returns a tuple with the Customproperty24 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty24

`func (o *CreateEnterpriseRoleRequest) SetCustomproperty24(v string)`

SetCustomproperty24 sets Customproperty24 field to given value.

### HasCustomproperty24

`func (o *CreateEnterpriseRoleRequest) HasCustomproperty24() bool`

HasCustomproperty24 returns a boolean if a field has been set.

### GetCustomproperty25

`func (o *CreateEnterpriseRoleRequest) GetCustomproperty25() string`

GetCustomproperty25 returns the Customproperty25 field if non-nil, zero value otherwise.

### GetCustomproperty25Ok

`func (o *CreateEnterpriseRoleRequest) GetCustomproperty25Ok() (*string, bool)`

GetCustomproperty25Ok returns a tuple with the Customproperty25 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty25

`func (o *CreateEnterpriseRoleRequest) SetCustomproperty25(v string)`

SetCustomproperty25 sets Customproperty25 field to given value.

### HasCustomproperty25

`func (o *CreateEnterpriseRoleRequest) HasCustomproperty25() bool`

HasCustomproperty25 returns a boolean if a field has been set.

### GetCustomproperty26

`func (o *CreateEnterpriseRoleRequest) GetCustomproperty26() string`

GetCustomproperty26 returns the Customproperty26 field if non-nil, zero value otherwise.

### GetCustomproperty26Ok

`func (o *CreateEnterpriseRoleRequest) GetCustomproperty26Ok() (*string, bool)`

GetCustomproperty26Ok returns a tuple with the Customproperty26 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty26

`func (o *CreateEnterpriseRoleRequest) SetCustomproperty26(v string)`

SetCustomproperty26 sets Customproperty26 field to given value.

### HasCustomproperty26

`func (o *CreateEnterpriseRoleRequest) HasCustomproperty26() bool`

HasCustomproperty26 returns a boolean if a field has been set.

### GetCustomproperty27

`func (o *CreateEnterpriseRoleRequest) GetCustomproperty27() string`

GetCustomproperty27 returns the Customproperty27 field if non-nil, zero value otherwise.

### GetCustomproperty27Ok

`func (o *CreateEnterpriseRoleRequest) GetCustomproperty27Ok() (*string, bool)`

GetCustomproperty27Ok returns a tuple with the Customproperty27 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty27

`func (o *CreateEnterpriseRoleRequest) SetCustomproperty27(v string)`

SetCustomproperty27 sets Customproperty27 field to given value.

### HasCustomproperty27

`func (o *CreateEnterpriseRoleRequest) HasCustomproperty27() bool`

HasCustomproperty27 returns a boolean if a field has been set.

### GetCustomproperty28

`func (o *CreateEnterpriseRoleRequest) GetCustomproperty28() string`

GetCustomproperty28 returns the Customproperty28 field if non-nil, zero value otherwise.

### GetCustomproperty28Ok

`func (o *CreateEnterpriseRoleRequest) GetCustomproperty28Ok() (*string, bool)`

GetCustomproperty28Ok returns a tuple with the Customproperty28 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty28

`func (o *CreateEnterpriseRoleRequest) SetCustomproperty28(v string)`

SetCustomproperty28 sets Customproperty28 field to given value.

### HasCustomproperty28

`func (o *CreateEnterpriseRoleRequest) HasCustomproperty28() bool`

HasCustomproperty28 returns a boolean if a field has been set.

### GetCustomproperty29

`func (o *CreateEnterpriseRoleRequest) GetCustomproperty29() string`

GetCustomproperty29 returns the Customproperty29 field if non-nil, zero value otherwise.

### GetCustomproperty29Ok

`func (o *CreateEnterpriseRoleRequest) GetCustomproperty29Ok() (*string, bool)`

GetCustomproperty29Ok returns a tuple with the Customproperty29 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty29

`func (o *CreateEnterpriseRoleRequest) SetCustomproperty29(v string)`

SetCustomproperty29 sets Customproperty29 field to given value.

### HasCustomproperty29

`func (o *CreateEnterpriseRoleRequest) HasCustomproperty29() bool`

HasCustomproperty29 returns a boolean if a field has been set.

### GetCustomproperty30

`func (o *CreateEnterpriseRoleRequest) GetCustomproperty30() string`

GetCustomproperty30 returns the Customproperty30 field if non-nil, zero value otherwise.

### GetCustomproperty30Ok

`func (o *CreateEnterpriseRoleRequest) GetCustomproperty30Ok() (*string, bool)`

GetCustomproperty30Ok returns a tuple with the Customproperty30 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty30

`func (o *CreateEnterpriseRoleRequest) SetCustomproperty30(v string)`

SetCustomproperty30 sets Customproperty30 field to given value.

### HasCustomproperty30

`func (o *CreateEnterpriseRoleRequest) HasCustomproperty30() bool`

HasCustomproperty30 returns a boolean if a field has been set.

### GetCustomproperty31

`func (o *CreateEnterpriseRoleRequest) GetCustomproperty31() string`

GetCustomproperty31 returns the Customproperty31 field if non-nil, zero value otherwise.

### GetCustomproperty31Ok

`func (o *CreateEnterpriseRoleRequest) GetCustomproperty31Ok() (*string, bool)`

GetCustomproperty31Ok returns a tuple with the Customproperty31 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty31

`func (o *CreateEnterpriseRoleRequest) SetCustomproperty31(v string)`

SetCustomproperty31 sets Customproperty31 field to given value.

### HasCustomproperty31

`func (o *CreateEnterpriseRoleRequest) HasCustomproperty31() bool`

HasCustomproperty31 returns a boolean if a field has been set.

### GetCustomproperty32

`func (o *CreateEnterpriseRoleRequest) GetCustomproperty32() string`

GetCustomproperty32 returns the Customproperty32 field if non-nil, zero value otherwise.

### GetCustomproperty32Ok

`func (o *CreateEnterpriseRoleRequest) GetCustomproperty32Ok() (*string, bool)`

GetCustomproperty32Ok returns a tuple with the Customproperty32 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty32

`func (o *CreateEnterpriseRoleRequest) SetCustomproperty32(v string)`

SetCustomproperty32 sets Customproperty32 field to given value.

### HasCustomproperty32

`func (o *CreateEnterpriseRoleRequest) HasCustomproperty32() bool`

HasCustomproperty32 returns a boolean if a field has been set.

### GetCustomproperty33

`func (o *CreateEnterpriseRoleRequest) GetCustomproperty33() string`

GetCustomproperty33 returns the Customproperty33 field if non-nil, zero value otherwise.

### GetCustomproperty33Ok

`func (o *CreateEnterpriseRoleRequest) GetCustomproperty33Ok() (*string, bool)`

GetCustomproperty33Ok returns a tuple with the Customproperty33 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty33

`func (o *CreateEnterpriseRoleRequest) SetCustomproperty33(v string)`

SetCustomproperty33 sets Customproperty33 field to given value.

### HasCustomproperty33

`func (o *CreateEnterpriseRoleRequest) HasCustomproperty33() bool`

HasCustomproperty33 returns a boolean if a field has been set.

### GetCustomproperty34

`func (o *CreateEnterpriseRoleRequest) GetCustomproperty34() string`

GetCustomproperty34 returns the Customproperty34 field if non-nil, zero value otherwise.

### GetCustomproperty34Ok

`func (o *CreateEnterpriseRoleRequest) GetCustomproperty34Ok() (*string, bool)`

GetCustomproperty34Ok returns a tuple with the Customproperty34 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty34

`func (o *CreateEnterpriseRoleRequest) SetCustomproperty34(v string)`

SetCustomproperty34 sets Customproperty34 field to given value.

### HasCustomproperty34

`func (o *CreateEnterpriseRoleRequest) HasCustomproperty34() bool`

HasCustomproperty34 returns a boolean if a field has been set.

### GetCustomproperty35

`func (o *CreateEnterpriseRoleRequest) GetCustomproperty35() string`

GetCustomproperty35 returns the Customproperty35 field if non-nil, zero value otherwise.

### GetCustomproperty35Ok

`func (o *CreateEnterpriseRoleRequest) GetCustomproperty35Ok() (*string, bool)`

GetCustomproperty35Ok returns a tuple with the Customproperty35 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty35

`func (o *CreateEnterpriseRoleRequest) SetCustomproperty35(v string)`

SetCustomproperty35 sets Customproperty35 field to given value.

### HasCustomproperty35

`func (o *CreateEnterpriseRoleRequest) HasCustomproperty35() bool`

HasCustomproperty35 returns a boolean if a field has been set.

### GetCustomproperty36

`func (o *CreateEnterpriseRoleRequest) GetCustomproperty36() string`

GetCustomproperty36 returns the Customproperty36 field if non-nil, zero value otherwise.

### GetCustomproperty36Ok

`func (o *CreateEnterpriseRoleRequest) GetCustomproperty36Ok() (*string, bool)`

GetCustomproperty36Ok returns a tuple with the Customproperty36 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty36

`func (o *CreateEnterpriseRoleRequest) SetCustomproperty36(v string)`

SetCustomproperty36 sets Customproperty36 field to given value.

### HasCustomproperty36

`func (o *CreateEnterpriseRoleRequest) HasCustomproperty36() bool`

HasCustomproperty36 returns a boolean if a field has been set.

### GetCustomproperty37

`func (o *CreateEnterpriseRoleRequest) GetCustomproperty37() string`

GetCustomproperty37 returns the Customproperty37 field if non-nil, zero value otherwise.

### GetCustomproperty37Ok

`func (o *CreateEnterpriseRoleRequest) GetCustomproperty37Ok() (*string, bool)`

GetCustomproperty37Ok returns a tuple with the Customproperty37 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty37

`func (o *CreateEnterpriseRoleRequest) SetCustomproperty37(v string)`

SetCustomproperty37 sets Customproperty37 field to given value.

### HasCustomproperty37

`func (o *CreateEnterpriseRoleRequest) HasCustomproperty37() bool`

HasCustomproperty37 returns a boolean if a field has been set.

### GetCustomproperty38

`func (o *CreateEnterpriseRoleRequest) GetCustomproperty38() string`

GetCustomproperty38 returns the Customproperty38 field if non-nil, zero value otherwise.

### GetCustomproperty38Ok

`func (o *CreateEnterpriseRoleRequest) GetCustomproperty38Ok() (*string, bool)`

GetCustomproperty38Ok returns a tuple with the Customproperty38 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty38

`func (o *CreateEnterpriseRoleRequest) SetCustomproperty38(v string)`

SetCustomproperty38 sets Customproperty38 field to given value.

### HasCustomproperty38

`func (o *CreateEnterpriseRoleRequest) HasCustomproperty38() bool`

HasCustomproperty38 returns a boolean if a field has been set.

### GetCustomproperty39

`func (o *CreateEnterpriseRoleRequest) GetCustomproperty39() string`

GetCustomproperty39 returns the Customproperty39 field if non-nil, zero value otherwise.

### GetCustomproperty39Ok

`func (o *CreateEnterpriseRoleRequest) GetCustomproperty39Ok() (*string, bool)`

GetCustomproperty39Ok returns a tuple with the Customproperty39 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty39

`func (o *CreateEnterpriseRoleRequest) SetCustomproperty39(v string)`

SetCustomproperty39 sets Customproperty39 field to given value.

### HasCustomproperty39

`func (o *CreateEnterpriseRoleRequest) HasCustomproperty39() bool`

HasCustomproperty39 returns a boolean if a field has been set.

### GetCustomproperty40

`func (o *CreateEnterpriseRoleRequest) GetCustomproperty40() string`

GetCustomproperty40 returns the Customproperty40 field if non-nil, zero value otherwise.

### GetCustomproperty40Ok

`func (o *CreateEnterpriseRoleRequest) GetCustomproperty40Ok() (*string, bool)`

GetCustomproperty40Ok returns a tuple with the Customproperty40 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty40

`func (o *CreateEnterpriseRoleRequest) SetCustomproperty40(v string)`

SetCustomproperty40 sets Customproperty40 field to given value.

### HasCustomproperty40

`func (o *CreateEnterpriseRoleRequest) HasCustomproperty40() bool`

HasCustomproperty40 returns a boolean if a field has been set.

### GetCustomproperty41

`func (o *CreateEnterpriseRoleRequest) GetCustomproperty41() string`

GetCustomproperty41 returns the Customproperty41 field if non-nil, zero value otherwise.

### GetCustomproperty41Ok

`func (o *CreateEnterpriseRoleRequest) GetCustomproperty41Ok() (*string, bool)`

GetCustomproperty41Ok returns a tuple with the Customproperty41 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty41

`func (o *CreateEnterpriseRoleRequest) SetCustomproperty41(v string)`

SetCustomproperty41 sets Customproperty41 field to given value.

### HasCustomproperty41

`func (o *CreateEnterpriseRoleRequest) HasCustomproperty41() bool`

HasCustomproperty41 returns a boolean if a field has been set.

### GetCustomproperty42

`func (o *CreateEnterpriseRoleRequest) GetCustomproperty42() string`

GetCustomproperty42 returns the Customproperty42 field if non-nil, zero value otherwise.

### GetCustomproperty42Ok

`func (o *CreateEnterpriseRoleRequest) GetCustomproperty42Ok() (*string, bool)`

GetCustomproperty42Ok returns a tuple with the Customproperty42 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty42

`func (o *CreateEnterpriseRoleRequest) SetCustomproperty42(v string)`

SetCustomproperty42 sets Customproperty42 field to given value.

### HasCustomproperty42

`func (o *CreateEnterpriseRoleRequest) HasCustomproperty42() bool`

HasCustomproperty42 returns a boolean if a field has been set.

### GetCustomproperty43

`func (o *CreateEnterpriseRoleRequest) GetCustomproperty43() string`

GetCustomproperty43 returns the Customproperty43 field if non-nil, zero value otherwise.

### GetCustomproperty43Ok

`func (o *CreateEnterpriseRoleRequest) GetCustomproperty43Ok() (*string, bool)`

GetCustomproperty43Ok returns a tuple with the Customproperty43 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty43

`func (o *CreateEnterpriseRoleRequest) SetCustomproperty43(v string)`

SetCustomproperty43 sets Customproperty43 field to given value.

### HasCustomproperty43

`func (o *CreateEnterpriseRoleRequest) HasCustomproperty43() bool`

HasCustomproperty43 returns a boolean if a field has been set.

### GetCustomproperty44

`func (o *CreateEnterpriseRoleRequest) GetCustomproperty44() string`

GetCustomproperty44 returns the Customproperty44 field if non-nil, zero value otherwise.

### GetCustomproperty44Ok

`func (o *CreateEnterpriseRoleRequest) GetCustomproperty44Ok() (*string, bool)`

GetCustomproperty44Ok returns a tuple with the Customproperty44 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty44

`func (o *CreateEnterpriseRoleRequest) SetCustomproperty44(v string)`

SetCustomproperty44 sets Customproperty44 field to given value.

### HasCustomproperty44

`func (o *CreateEnterpriseRoleRequest) HasCustomproperty44() bool`

HasCustomproperty44 returns a boolean if a field has been set.

### GetCustomproperty45

`func (o *CreateEnterpriseRoleRequest) GetCustomproperty45() string`

GetCustomproperty45 returns the Customproperty45 field if non-nil, zero value otherwise.

### GetCustomproperty45Ok

`func (o *CreateEnterpriseRoleRequest) GetCustomproperty45Ok() (*string, bool)`

GetCustomproperty45Ok returns a tuple with the Customproperty45 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty45

`func (o *CreateEnterpriseRoleRequest) SetCustomproperty45(v string)`

SetCustomproperty45 sets Customproperty45 field to given value.

### HasCustomproperty45

`func (o *CreateEnterpriseRoleRequest) HasCustomproperty45() bool`

HasCustomproperty45 returns a boolean if a field has been set.

### GetCustomproperty46

`func (o *CreateEnterpriseRoleRequest) GetCustomproperty46() string`

GetCustomproperty46 returns the Customproperty46 field if non-nil, zero value otherwise.

### GetCustomproperty46Ok

`func (o *CreateEnterpriseRoleRequest) GetCustomproperty46Ok() (*string, bool)`

GetCustomproperty46Ok returns a tuple with the Customproperty46 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty46

`func (o *CreateEnterpriseRoleRequest) SetCustomproperty46(v string)`

SetCustomproperty46 sets Customproperty46 field to given value.

### HasCustomproperty46

`func (o *CreateEnterpriseRoleRequest) HasCustomproperty46() bool`

HasCustomproperty46 returns a boolean if a field has been set.

### GetCustomproperty47

`func (o *CreateEnterpriseRoleRequest) GetCustomproperty47() string`

GetCustomproperty47 returns the Customproperty47 field if non-nil, zero value otherwise.

### GetCustomproperty47Ok

`func (o *CreateEnterpriseRoleRequest) GetCustomproperty47Ok() (*string, bool)`

GetCustomproperty47Ok returns a tuple with the Customproperty47 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty47

`func (o *CreateEnterpriseRoleRequest) SetCustomproperty47(v string)`

SetCustomproperty47 sets Customproperty47 field to given value.

### HasCustomproperty47

`func (o *CreateEnterpriseRoleRequest) HasCustomproperty47() bool`

HasCustomproperty47 returns a boolean if a field has been set.

### GetCustomproperty48

`func (o *CreateEnterpriseRoleRequest) GetCustomproperty48() string`

GetCustomproperty48 returns the Customproperty48 field if non-nil, zero value otherwise.

### GetCustomproperty48Ok

`func (o *CreateEnterpriseRoleRequest) GetCustomproperty48Ok() (*string, bool)`

GetCustomproperty48Ok returns a tuple with the Customproperty48 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty48

`func (o *CreateEnterpriseRoleRequest) SetCustomproperty48(v string)`

SetCustomproperty48 sets Customproperty48 field to given value.

### HasCustomproperty48

`func (o *CreateEnterpriseRoleRequest) HasCustomproperty48() bool`

HasCustomproperty48 returns a boolean if a field has been set.

### GetCustomproperty49

`func (o *CreateEnterpriseRoleRequest) GetCustomproperty49() string`

GetCustomproperty49 returns the Customproperty49 field if non-nil, zero value otherwise.

### GetCustomproperty49Ok

`func (o *CreateEnterpriseRoleRequest) GetCustomproperty49Ok() (*string, bool)`

GetCustomproperty49Ok returns a tuple with the Customproperty49 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty49

`func (o *CreateEnterpriseRoleRequest) SetCustomproperty49(v string)`

SetCustomproperty49 sets Customproperty49 field to given value.

### HasCustomproperty49

`func (o *CreateEnterpriseRoleRequest) HasCustomproperty49() bool`

HasCustomproperty49 returns a boolean if a field has been set.

### GetCustomproperty50

`func (o *CreateEnterpriseRoleRequest) GetCustomproperty50() string`

GetCustomproperty50 returns the Customproperty50 field if non-nil, zero value otherwise.

### GetCustomproperty50Ok

`func (o *CreateEnterpriseRoleRequest) GetCustomproperty50Ok() (*string, bool)`

GetCustomproperty50Ok returns a tuple with the Customproperty50 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty50

`func (o *CreateEnterpriseRoleRequest) SetCustomproperty50(v string)`

SetCustomproperty50 sets Customproperty50 field to given value.

### HasCustomproperty50

`func (o *CreateEnterpriseRoleRequest) HasCustomproperty50() bool`

HasCustomproperty50 returns a boolean if a field has been set.

### GetCustomproperty51

`func (o *CreateEnterpriseRoleRequest) GetCustomproperty51() string`

GetCustomproperty51 returns the Customproperty51 field if non-nil, zero value otherwise.

### GetCustomproperty51Ok

`func (o *CreateEnterpriseRoleRequest) GetCustomproperty51Ok() (*string, bool)`

GetCustomproperty51Ok returns a tuple with the Customproperty51 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty51

`func (o *CreateEnterpriseRoleRequest) SetCustomproperty51(v string)`

SetCustomproperty51 sets Customproperty51 field to given value.

### HasCustomproperty51

`func (o *CreateEnterpriseRoleRequest) HasCustomproperty51() bool`

HasCustomproperty51 returns a boolean if a field has been set.

### GetCustomproperty52

`func (o *CreateEnterpriseRoleRequest) GetCustomproperty52() string`

GetCustomproperty52 returns the Customproperty52 field if non-nil, zero value otherwise.

### GetCustomproperty52Ok

`func (o *CreateEnterpriseRoleRequest) GetCustomproperty52Ok() (*string, bool)`

GetCustomproperty52Ok returns a tuple with the Customproperty52 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty52

`func (o *CreateEnterpriseRoleRequest) SetCustomproperty52(v string)`

SetCustomproperty52 sets Customproperty52 field to given value.

### HasCustomproperty52

`func (o *CreateEnterpriseRoleRequest) HasCustomproperty52() bool`

HasCustomproperty52 returns a boolean if a field has been set.

### GetCustomproperty53

`func (o *CreateEnterpriseRoleRequest) GetCustomproperty53() string`

GetCustomproperty53 returns the Customproperty53 field if non-nil, zero value otherwise.

### GetCustomproperty53Ok

`func (o *CreateEnterpriseRoleRequest) GetCustomproperty53Ok() (*string, bool)`

GetCustomproperty53Ok returns a tuple with the Customproperty53 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty53

`func (o *CreateEnterpriseRoleRequest) SetCustomproperty53(v string)`

SetCustomproperty53 sets Customproperty53 field to given value.

### HasCustomproperty53

`func (o *CreateEnterpriseRoleRequest) HasCustomproperty53() bool`

HasCustomproperty53 returns a boolean if a field has been set.

### GetCustomproperty54

`func (o *CreateEnterpriseRoleRequest) GetCustomproperty54() string`

GetCustomproperty54 returns the Customproperty54 field if non-nil, zero value otherwise.

### GetCustomproperty54Ok

`func (o *CreateEnterpriseRoleRequest) GetCustomproperty54Ok() (*string, bool)`

GetCustomproperty54Ok returns a tuple with the Customproperty54 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty54

`func (o *CreateEnterpriseRoleRequest) SetCustomproperty54(v string)`

SetCustomproperty54 sets Customproperty54 field to given value.

### HasCustomproperty54

`func (o *CreateEnterpriseRoleRequest) HasCustomproperty54() bool`

HasCustomproperty54 returns a boolean if a field has been set.

### GetCustomproperty55

`func (o *CreateEnterpriseRoleRequest) GetCustomproperty55() string`

GetCustomproperty55 returns the Customproperty55 field if non-nil, zero value otherwise.

### GetCustomproperty55Ok

`func (o *CreateEnterpriseRoleRequest) GetCustomproperty55Ok() (*string, bool)`

GetCustomproperty55Ok returns a tuple with the Customproperty55 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty55

`func (o *CreateEnterpriseRoleRequest) SetCustomproperty55(v string)`

SetCustomproperty55 sets Customproperty55 field to given value.

### HasCustomproperty55

`func (o *CreateEnterpriseRoleRequest) HasCustomproperty55() bool`

HasCustomproperty55 returns a boolean if a field has been set.

### GetCustomproperty56

`func (o *CreateEnterpriseRoleRequest) GetCustomproperty56() string`

GetCustomproperty56 returns the Customproperty56 field if non-nil, zero value otherwise.

### GetCustomproperty56Ok

`func (o *CreateEnterpriseRoleRequest) GetCustomproperty56Ok() (*string, bool)`

GetCustomproperty56Ok returns a tuple with the Customproperty56 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty56

`func (o *CreateEnterpriseRoleRequest) SetCustomproperty56(v string)`

SetCustomproperty56 sets Customproperty56 field to given value.

### HasCustomproperty56

`func (o *CreateEnterpriseRoleRequest) HasCustomproperty56() bool`

HasCustomproperty56 returns a boolean if a field has been set.

### GetCustomproperty57

`func (o *CreateEnterpriseRoleRequest) GetCustomproperty57() string`

GetCustomproperty57 returns the Customproperty57 field if non-nil, zero value otherwise.

### GetCustomproperty57Ok

`func (o *CreateEnterpriseRoleRequest) GetCustomproperty57Ok() (*string, bool)`

GetCustomproperty57Ok returns a tuple with the Customproperty57 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty57

`func (o *CreateEnterpriseRoleRequest) SetCustomproperty57(v string)`

SetCustomproperty57 sets Customproperty57 field to given value.

### HasCustomproperty57

`func (o *CreateEnterpriseRoleRequest) HasCustomproperty57() bool`

HasCustomproperty57 returns a boolean if a field has been set.

### GetCustomproperty58

`func (o *CreateEnterpriseRoleRequest) GetCustomproperty58() string`

GetCustomproperty58 returns the Customproperty58 field if non-nil, zero value otherwise.

### GetCustomproperty58Ok

`func (o *CreateEnterpriseRoleRequest) GetCustomproperty58Ok() (*string, bool)`

GetCustomproperty58Ok returns a tuple with the Customproperty58 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty58

`func (o *CreateEnterpriseRoleRequest) SetCustomproperty58(v string)`

SetCustomproperty58 sets Customproperty58 field to given value.

### HasCustomproperty58

`func (o *CreateEnterpriseRoleRequest) HasCustomproperty58() bool`

HasCustomproperty58 returns a boolean if a field has been set.

### GetCustomproperty59

`func (o *CreateEnterpriseRoleRequest) GetCustomproperty59() string`

GetCustomproperty59 returns the Customproperty59 field if non-nil, zero value otherwise.

### GetCustomproperty59Ok

`func (o *CreateEnterpriseRoleRequest) GetCustomproperty59Ok() (*string, bool)`

GetCustomproperty59Ok returns a tuple with the Customproperty59 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty59

`func (o *CreateEnterpriseRoleRequest) SetCustomproperty59(v string)`

SetCustomproperty59 sets Customproperty59 field to given value.

### HasCustomproperty59

`func (o *CreateEnterpriseRoleRequest) HasCustomproperty59() bool`

HasCustomproperty59 returns a boolean if a field has been set.

### GetCustomproperty60

`func (o *CreateEnterpriseRoleRequest) GetCustomproperty60() string`

GetCustomproperty60 returns the Customproperty60 field if non-nil, zero value otherwise.

### GetCustomproperty60Ok

`func (o *CreateEnterpriseRoleRequest) GetCustomproperty60Ok() (*string, bool)`

GetCustomproperty60Ok returns a tuple with the Customproperty60 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty60

`func (o *CreateEnterpriseRoleRequest) SetCustomproperty60(v string)`

SetCustomproperty60 sets Customproperty60 field to given value.

### HasCustomproperty60

`func (o *CreateEnterpriseRoleRequest) HasCustomproperty60() bool`

HasCustomproperty60 returns a boolean if a field has been set.

### GetEndpointname

`func (o *CreateEnterpriseRoleRequest) GetEndpointname() string`

GetEndpointname returns the Endpointname field if non-nil, zero value otherwise.

### GetEndpointnameOk

`func (o *CreateEnterpriseRoleRequest) GetEndpointnameOk() (*string, bool)`

GetEndpointnameOk returns a tuple with the Endpointname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpointname

`func (o *CreateEnterpriseRoleRequest) SetEndpointname(v string)`

SetEndpointname sets Endpointname field to given value.

### HasEndpointname

`func (o *CreateEnterpriseRoleRequest) HasEndpointname() bool`

HasEndpointname returns a boolean if a field has been set.

### GetDefaulttimeframe

`func (o *CreateEnterpriseRoleRequest) GetDefaulttimeframe() string`

GetDefaulttimeframe returns the Defaulttimeframe field if non-nil, zero value otherwise.

### GetDefaulttimeframeOk

`func (o *CreateEnterpriseRoleRequest) GetDefaulttimeframeOk() (*string, bool)`

GetDefaulttimeframeOk returns a tuple with the Defaulttimeframe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaulttimeframe

`func (o *CreateEnterpriseRoleRequest) SetDefaulttimeframe(v string)`

SetDefaulttimeframe sets Defaulttimeframe field to given value.

### HasDefaulttimeframe

`func (o *CreateEnterpriseRoleRequest) HasDefaulttimeframe() bool`

HasDefaulttimeframe returns a boolean if a field has been set.

### GetDescription

`func (o *CreateEnterpriseRoleRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateEnterpriseRoleRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateEnterpriseRoleRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateEnterpriseRoleRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDisplayname

`func (o *CreateEnterpriseRoleRequest) GetDisplayname() string`

GetDisplayname returns the Displayname field if non-nil, zero value otherwise.

### GetDisplaynameOk

`func (o *CreateEnterpriseRoleRequest) GetDisplaynameOk() (*string, bool)`

GetDisplaynameOk returns a tuple with the Displayname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayname

`func (o *CreateEnterpriseRoleRequest) SetDisplayname(v string)`

SetDisplayname sets Displayname field to given value.

### HasDisplayname

`func (o *CreateEnterpriseRoleRequest) HasDisplayname() bool`

HasDisplayname returns a boolean if a field has been set.

### GetGlossary

`func (o *CreateEnterpriseRoleRequest) GetGlossary() string`

GetGlossary returns the Glossary field if non-nil, zero value otherwise.

### GetGlossaryOk

`func (o *CreateEnterpriseRoleRequest) GetGlossaryOk() (*string, bool)`

GetGlossaryOk returns a tuple with the Glossary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlossary

`func (o *CreateEnterpriseRoleRequest) SetGlossary(v string)`

SetGlossary sets Glossary field to given value.

### HasGlossary

`func (o *CreateEnterpriseRoleRequest) HasGlossary() bool`

HasGlossary returns a boolean if a field has been set.

### GetRisk

`func (o *CreateEnterpriseRoleRequest) GetRisk() string`

GetRisk returns the Risk field if non-nil, zero value otherwise.

### GetRiskOk

`func (o *CreateEnterpriseRoleRequest) GetRiskOk() (*string, bool)`

GetRiskOk returns a tuple with the Risk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRisk

`func (o *CreateEnterpriseRoleRequest) SetRisk(v string)`

SetRisk sets Risk field to given value.

### HasRisk

`func (o *CreateEnterpriseRoleRequest) HasRisk() bool`

HasRisk returns a boolean if a field has been set.

### GetLevel

`func (o *CreateEnterpriseRoleRequest) GetLevel() string`

GetLevel returns the Level field if non-nil, zero value otherwise.

### GetLevelOk

`func (o *CreateEnterpriseRoleRequest) GetLevelOk() (*string, bool)`

GetLevelOk returns a tuple with the Level field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevel

`func (o *CreateEnterpriseRoleRequest) SetLevel(v string)`

SetLevel sets Level field to given value.

### HasLevel

`func (o *CreateEnterpriseRoleRequest) HasLevel() bool`

HasLevel returns a boolean if a field has been set.

### GetSoxcritical

`func (o *CreateEnterpriseRoleRequest) GetSoxcritical() string`

GetSoxcritical returns the Soxcritical field if non-nil, zero value otherwise.

### GetSoxcriticalOk

`func (o *CreateEnterpriseRoleRequest) GetSoxcriticalOk() (*string, bool)`

GetSoxcriticalOk returns a tuple with the Soxcritical field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoxcritical

`func (o *CreateEnterpriseRoleRequest) SetSoxcritical(v string)`

SetSoxcritical sets Soxcritical field to given value.

### HasSoxcritical

`func (o *CreateEnterpriseRoleRequest) HasSoxcritical() bool`

HasSoxcritical returns a boolean if a field has been set.

### GetSyscritical

`func (o *CreateEnterpriseRoleRequest) GetSyscritical() string`

GetSyscritical returns the Syscritical field if non-nil, zero value otherwise.

### GetSyscriticalOk

`func (o *CreateEnterpriseRoleRequest) GetSyscriticalOk() (*string, bool)`

GetSyscriticalOk returns a tuple with the Syscritical field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyscritical

`func (o *CreateEnterpriseRoleRequest) SetSyscritical(v string)`

SetSyscritical sets Syscritical field to given value.

### HasSyscritical

`func (o *CreateEnterpriseRoleRequest) HasSyscritical() bool`

HasSyscritical returns a boolean if a field has been set.

### GetPriviliged

`func (o *CreateEnterpriseRoleRequest) GetPriviliged() string`

GetPriviliged returns the Priviliged field if non-nil, zero value otherwise.

### GetPriviligedOk

`func (o *CreateEnterpriseRoleRequest) GetPriviligedOk() (*string, bool)`

GetPriviligedOk returns a tuple with the Priviliged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriviliged

`func (o *CreateEnterpriseRoleRequest) SetPriviliged(v string)`

SetPriviliged sets Priviliged field to given value.

### HasPriviliged

`func (o *CreateEnterpriseRoleRequest) HasPriviliged() bool`

HasPriviliged returns a boolean if a field has been set.

### GetConfidentiality

`func (o *CreateEnterpriseRoleRequest) GetConfidentiality() string`

GetConfidentiality returns the Confidentiality field if non-nil, zero value otherwise.

### GetConfidentialityOk

`func (o *CreateEnterpriseRoleRequest) GetConfidentialityOk() (*string, bool)`

GetConfidentialityOk returns a tuple with the Confidentiality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfidentiality

`func (o *CreateEnterpriseRoleRequest) SetConfidentiality(v string)`

SetConfidentiality sets Confidentiality field to given value.

### HasConfidentiality

`func (o *CreateEnterpriseRoleRequest) HasConfidentiality() bool`

HasConfidentiality returns a boolean if a field has been set.

### GetRequestable

`func (o *CreateEnterpriseRoleRequest) GetRequestable() string`

GetRequestable returns the Requestable field if non-nil, zero value otherwise.

### GetRequestableOk

`func (o *CreateEnterpriseRoleRequest) GetRequestableOk() (*string, bool)`

GetRequestableOk returns a tuple with the Requestable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestable

`func (o *CreateEnterpriseRoleRequest) SetRequestable(v string)`

SetRequestable sets Requestable field to given value.

### HasRequestable

`func (o *CreateEnterpriseRoleRequest) HasRequestable() bool`

HasRequestable returns a boolean if a field has been set.

### GetShowDynamicAttrs

`func (o *CreateEnterpriseRoleRequest) GetShowDynamicAttrs() string`

GetShowDynamicAttrs returns the ShowDynamicAttrs field if non-nil, zero value otherwise.

### GetShowDynamicAttrsOk

`func (o *CreateEnterpriseRoleRequest) GetShowDynamicAttrsOk() (*string, bool)`

GetShowDynamicAttrsOk returns a tuple with the ShowDynamicAttrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowDynamicAttrs

`func (o *CreateEnterpriseRoleRequest) SetShowDynamicAttrs(v string)`

SetShowDynamicAttrs sets ShowDynamicAttrs field to given value.

### HasShowDynamicAttrs

`func (o *CreateEnterpriseRoleRequest) HasShowDynamicAttrs() bool`

HasShowDynamicAttrs returns a boolean if a field has been set.

### GetChecksod

`func (o *CreateEnterpriseRoleRequest) GetChecksod() string`

GetChecksod returns the Checksod field if non-nil, zero value otherwise.

### GetChecksodOk

`func (o *CreateEnterpriseRoleRequest) GetChecksodOk() (*string, bool)`

GetChecksodOk returns a tuple with the Checksod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChecksod

`func (o *CreateEnterpriseRoleRequest) SetChecksod(v string)`

SetChecksod sets Checksod field to given value.

### HasChecksod

`func (o *CreateEnterpriseRoleRequest) HasChecksod() bool`

HasChecksod returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


