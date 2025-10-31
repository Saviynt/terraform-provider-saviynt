# GetRoleDetailsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RoleKey** | Pointer to **int32** |  | [optional] 
**Updatedate** | Pointer to **string** |  | [optional] 
**Roletype** | Pointer to **string** |  | [optional] 
**Version** | Pointer to [**GetRoleDetailsResponseVersion**](GetRoleDetailsResponseVersion.md) |  | [optional] 
**RoleName** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Glossary** | Pointer to **string** |  | [optional] 
**Priviliged** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**ShowDynamicAttrs** | Pointer to **string** |  | [optional] 
**DefaultTimeFrameHrs** | Pointer to **string** |  | [optional] 
**MaxTimeFrameHrs** | Pointer to **string** |  | [optional] 
**Confidentiality** | Pointer to **string** |  | [optional] 
**Soxcritical** | Pointer to **string** |  | [optional] 
**Syscritical** | Pointer to **string** |  | [optional] 
**Requestable** | Pointer to **string** |  | [optional] 
**Displayname** | Pointer to **string** |  | [optional] 
**Updateuser** | Pointer to **string** |  | [optional] 
**EntitlementValueKey** | Pointer to **string** |  | [optional] 
**RoleState** | Pointer to **string** |  | [optional] 
**Endpointkey** | Pointer to **string** |  | [optional] 
**LastReviewedCampaignName** | Pointer to **string** |  | [optional] 
**Risk** | Pointer to **string** |  | [optional] 
**LastReviewedBy** | Pointer to **string** |  | [optional] 
**Owner** | Pointer to [**GetRoleDetailsResponseOwner**](GetRoleDetailsResponseOwner.md) |  | [optional] 
**UserDetails** | Pointer to [**[]GetRoleDetailsResponseUserDetailsInner**](GetRoleDetailsResponseUserDetailsInner.md) |  | [optional] 
**EntitlementDetails** | Pointer to [**[]GetEntitlementDetailsResponse**](GetEntitlementDetailsResponse.md) |  | [optional] 
**CustomProperty1** | Pointer to **string** |  | [optional] 
**CustomProperty2** | Pointer to **string** |  | [optional] 
**CustomProperty3** | Pointer to **string** |  | [optional] 
**CustomProperty4** | Pointer to **string** |  | [optional] 
**CustomProperty5** | Pointer to **string** |  | [optional] 
**CustomProperty6** | Pointer to **string** |  | [optional] 
**CustomProperty7** | Pointer to **string** |  | [optional] 
**CustomProperty8** | Pointer to **string** |  | [optional] 
**CustomProperty9** | Pointer to **string** |  | [optional] 
**CustomProperty10** | Pointer to **string** |  | [optional] 
**CustomProperty11** | Pointer to **string** |  | [optional] 
**CustomProperty12** | Pointer to **string** |  | [optional] 
**CustomProperty13** | Pointer to **string** |  | [optional] 
**CustomProperty14** | Pointer to **string** |  | [optional] 
**CustomProperty15** | Pointer to **string** |  | [optional] 
**CustomProperty16** | Pointer to **string** |  | [optional] 
**CustomProperty17** | Pointer to **string** |  | [optional] 
**CustomProperty18** | Pointer to **string** |  | [optional] 
**CustomProperty19** | Pointer to **string** |  | [optional] 
**CustomProperty20** | Pointer to **string** |  | [optional] 
**CustomProperty21** | Pointer to **string** |  | [optional] 
**CustomProperty22** | Pointer to **string** |  | [optional] 
**CustomProperty23** | Pointer to **string** |  | [optional] 
**CustomProperty24** | Pointer to **string** |  | [optional] 
**CustomProperty25** | Pointer to **string** |  | [optional] 
**CustomProperty26** | Pointer to **string** |  | [optional] 
**CustomProperty27** | Pointer to **string** |  | [optional] 
**CustomProperty28** | Pointer to **string** |  | [optional] 
**CustomProperty29** | Pointer to **string** |  | [optional] 
**CustomProperty30** | Pointer to **string** |  | [optional] 
**CustomProperty31** | Pointer to **string** |  | [optional] 
**CustomProperty32** | Pointer to **string** |  | [optional] 
**CustomProperty33** | Pointer to **string** |  | [optional] 
**CustomProperty34** | Pointer to **string** |  | [optional] 
**CustomProperty35** | Pointer to **string** |  | [optional] 
**CustomProperty36** | Pointer to **string** |  | [optional] 
**CustomProperty37** | Pointer to **string** |  | [optional] 
**CustomProperty38** | Pointer to **string** |  | [optional] 
**CustomProperty39** | Pointer to **string** |  | [optional] 
**CustomProperty40** | Pointer to **string** |  | [optional] 
**CustomProperty41** | Pointer to **string** |  | [optional] 
**CustomProperty42** | Pointer to **string** |  | [optional] 
**CustomProperty43** | Pointer to **string** |  | [optional] 
**CustomProperty44** | Pointer to **string** |  | [optional] 
**CustomProperty45** | Pointer to **string** |  | [optional] 
**CustomProperty46** | Pointer to **string** |  | [optional] 
**CustomProperty47** | Pointer to **string** |  | [optional] 
**CustomProperty48** | Pointer to **string** |  | [optional] 
**CustomProperty49** | Pointer to **string** |  | [optional] 
**CustomProperty50** | Pointer to **string** |  | [optional] 
**CustomProperty51** | Pointer to **string** |  | [optional] 
**CustomProperty52** | Pointer to **string** |  | [optional] 
**CustomProperty53** | Pointer to **string** |  | [optional] 
**CustomProperty54** | Pointer to **string** |  | [optional] 
**CustomProperty55** | Pointer to **string** |  | [optional] 
**CustomProperty56** | Pointer to **string** |  | [optional] 
**CustomProperty57** | Pointer to **string** |  | [optional] 
**CustomProperty58** | Pointer to **string** |  | [optional] 
**CustomProperty59** | Pointer to **string** |  | [optional] 
**CustomProperty60** | Pointer to **string** |  | [optional] 

## Methods

### NewGetRoleDetailsResponse

`func NewGetRoleDetailsResponse() *GetRoleDetailsResponse`

NewGetRoleDetailsResponse instantiates a new GetRoleDetailsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetRoleDetailsResponseWithDefaults

`func NewGetRoleDetailsResponseWithDefaults() *GetRoleDetailsResponse`

NewGetRoleDetailsResponseWithDefaults instantiates a new GetRoleDetailsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRoleKey

`func (o *GetRoleDetailsResponse) GetRoleKey() int32`

GetRoleKey returns the RoleKey field if non-nil, zero value otherwise.

### GetRoleKeyOk

`func (o *GetRoleDetailsResponse) GetRoleKeyOk() (*int32, bool)`

GetRoleKeyOk returns a tuple with the RoleKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleKey

`func (o *GetRoleDetailsResponse) SetRoleKey(v int32)`

SetRoleKey sets RoleKey field to given value.

### HasRoleKey

`func (o *GetRoleDetailsResponse) HasRoleKey() bool`

HasRoleKey returns a boolean if a field has been set.

### GetUpdatedate

`func (o *GetRoleDetailsResponse) GetUpdatedate() string`

GetUpdatedate returns the Updatedate field if non-nil, zero value otherwise.

### GetUpdatedateOk

`func (o *GetRoleDetailsResponse) GetUpdatedateOk() (*string, bool)`

GetUpdatedateOk returns a tuple with the Updatedate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedate

`func (o *GetRoleDetailsResponse) SetUpdatedate(v string)`

SetUpdatedate sets Updatedate field to given value.

### HasUpdatedate

`func (o *GetRoleDetailsResponse) HasUpdatedate() bool`

HasUpdatedate returns a boolean if a field has been set.

### GetRoletype

`func (o *GetRoleDetailsResponse) GetRoletype() string`

GetRoletype returns the Roletype field if non-nil, zero value otherwise.

### GetRoletypeOk

`func (o *GetRoleDetailsResponse) GetRoletypeOk() (*string, bool)`

GetRoletypeOk returns a tuple with the Roletype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoletype

`func (o *GetRoleDetailsResponse) SetRoletype(v string)`

SetRoletype sets Roletype field to given value.

### HasRoletype

`func (o *GetRoleDetailsResponse) HasRoletype() bool`

HasRoletype returns a boolean if a field has been set.

### GetVersion

`func (o *GetRoleDetailsResponse) GetVersion() GetRoleDetailsResponseVersion`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *GetRoleDetailsResponse) GetVersionOk() (*GetRoleDetailsResponseVersion, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *GetRoleDetailsResponse) SetVersion(v GetRoleDetailsResponseVersion)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *GetRoleDetailsResponse) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetRoleName

`func (o *GetRoleDetailsResponse) GetRoleName() string`

GetRoleName returns the RoleName field if non-nil, zero value otherwise.

### GetRoleNameOk

`func (o *GetRoleDetailsResponse) GetRoleNameOk() (*string, bool)`

GetRoleNameOk returns a tuple with the RoleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleName

`func (o *GetRoleDetailsResponse) SetRoleName(v string)`

SetRoleName sets RoleName field to given value.

### HasRoleName

`func (o *GetRoleDetailsResponse) HasRoleName() bool`

HasRoleName returns a boolean if a field has been set.

### GetDescription

`func (o *GetRoleDetailsResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GetRoleDetailsResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GetRoleDetailsResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GetRoleDetailsResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetGlossary

`func (o *GetRoleDetailsResponse) GetGlossary() string`

GetGlossary returns the Glossary field if non-nil, zero value otherwise.

### GetGlossaryOk

`func (o *GetRoleDetailsResponse) GetGlossaryOk() (*string, bool)`

GetGlossaryOk returns a tuple with the Glossary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlossary

`func (o *GetRoleDetailsResponse) SetGlossary(v string)`

SetGlossary sets Glossary field to given value.

### HasGlossary

`func (o *GetRoleDetailsResponse) HasGlossary() bool`

HasGlossary returns a boolean if a field has been set.

### GetPriviliged

`func (o *GetRoleDetailsResponse) GetPriviliged() string`

GetPriviliged returns the Priviliged field if non-nil, zero value otherwise.

### GetPriviligedOk

`func (o *GetRoleDetailsResponse) GetPriviligedOk() (*string, bool)`

GetPriviligedOk returns a tuple with the Priviliged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriviliged

`func (o *GetRoleDetailsResponse) SetPriviliged(v string)`

SetPriviliged sets Priviliged field to given value.

### HasPriviliged

`func (o *GetRoleDetailsResponse) HasPriviliged() bool`

HasPriviliged returns a boolean if a field has been set.

### GetStatus

`func (o *GetRoleDetailsResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetRoleDetailsResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetRoleDetailsResponse) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetRoleDetailsResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetShowDynamicAttrs

`func (o *GetRoleDetailsResponse) GetShowDynamicAttrs() string`

GetShowDynamicAttrs returns the ShowDynamicAttrs field if non-nil, zero value otherwise.

### GetShowDynamicAttrsOk

`func (o *GetRoleDetailsResponse) GetShowDynamicAttrsOk() (*string, bool)`

GetShowDynamicAttrsOk returns a tuple with the ShowDynamicAttrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowDynamicAttrs

`func (o *GetRoleDetailsResponse) SetShowDynamicAttrs(v string)`

SetShowDynamicAttrs sets ShowDynamicAttrs field to given value.

### HasShowDynamicAttrs

`func (o *GetRoleDetailsResponse) HasShowDynamicAttrs() bool`

HasShowDynamicAttrs returns a boolean if a field has been set.

### GetDefaultTimeFrameHrs

`func (o *GetRoleDetailsResponse) GetDefaultTimeFrameHrs() string`

GetDefaultTimeFrameHrs returns the DefaultTimeFrameHrs field if non-nil, zero value otherwise.

### GetDefaultTimeFrameHrsOk

`func (o *GetRoleDetailsResponse) GetDefaultTimeFrameHrsOk() (*string, bool)`

GetDefaultTimeFrameHrsOk returns a tuple with the DefaultTimeFrameHrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultTimeFrameHrs

`func (o *GetRoleDetailsResponse) SetDefaultTimeFrameHrs(v string)`

SetDefaultTimeFrameHrs sets DefaultTimeFrameHrs field to given value.

### HasDefaultTimeFrameHrs

`func (o *GetRoleDetailsResponse) HasDefaultTimeFrameHrs() bool`

HasDefaultTimeFrameHrs returns a boolean if a field has been set.

### GetMaxTimeFrameHrs

`func (o *GetRoleDetailsResponse) GetMaxTimeFrameHrs() string`

GetMaxTimeFrameHrs returns the MaxTimeFrameHrs field if non-nil, zero value otherwise.

### GetMaxTimeFrameHrsOk

`func (o *GetRoleDetailsResponse) GetMaxTimeFrameHrsOk() (*string, bool)`

GetMaxTimeFrameHrsOk returns a tuple with the MaxTimeFrameHrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxTimeFrameHrs

`func (o *GetRoleDetailsResponse) SetMaxTimeFrameHrs(v string)`

SetMaxTimeFrameHrs sets MaxTimeFrameHrs field to given value.

### HasMaxTimeFrameHrs

`func (o *GetRoleDetailsResponse) HasMaxTimeFrameHrs() bool`

HasMaxTimeFrameHrs returns a boolean if a field has been set.

### GetConfidentiality

`func (o *GetRoleDetailsResponse) GetConfidentiality() string`

GetConfidentiality returns the Confidentiality field if non-nil, zero value otherwise.

### GetConfidentialityOk

`func (o *GetRoleDetailsResponse) GetConfidentialityOk() (*string, bool)`

GetConfidentialityOk returns a tuple with the Confidentiality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfidentiality

`func (o *GetRoleDetailsResponse) SetConfidentiality(v string)`

SetConfidentiality sets Confidentiality field to given value.

### HasConfidentiality

`func (o *GetRoleDetailsResponse) HasConfidentiality() bool`

HasConfidentiality returns a boolean if a field has been set.

### GetSoxcritical

`func (o *GetRoleDetailsResponse) GetSoxcritical() string`

GetSoxcritical returns the Soxcritical field if non-nil, zero value otherwise.

### GetSoxcriticalOk

`func (o *GetRoleDetailsResponse) GetSoxcriticalOk() (*string, bool)`

GetSoxcriticalOk returns a tuple with the Soxcritical field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoxcritical

`func (o *GetRoleDetailsResponse) SetSoxcritical(v string)`

SetSoxcritical sets Soxcritical field to given value.

### HasSoxcritical

`func (o *GetRoleDetailsResponse) HasSoxcritical() bool`

HasSoxcritical returns a boolean if a field has been set.

### GetSyscritical

`func (o *GetRoleDetailsResponse) GetSyscritical() string`

GetSyscritical returns the Syscritical field if non-nil, zero value otherwise.

### GetSyscriticalOk

`func (o *GetRoleDetailsResponse) GetSyscriticalOk() (*string, bool)`

GetSyscriticalOk returns a tuple with the Syscritical field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyscritical

`func (o *GetRoleDetailsResponse) SetSyscritical(v string)`

SetSyscritical sets Syscritical field to given value.

### HasSyscritical

`func (o *GetRoleDetailsResponse) HasSyscritical() bool`

HasSyscritical returns a boolean if a field has been set.

### GetRequestable

`func (o *GetRoleDetailsResponse) GetRequestable() string`

GetRequestable returns the Requestable field if non-nil, zero value otherwise.

### GetRequestableOk

`func (o *GetRoleDetailsResponse) GetRequestableOk() (*string, bool)`

GetRequestableOk returns a tuple with the Requestable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestable

`func (o *GetRoleDetailsResponse) SetRequestable(v string)`

SetRequestable sets Requestable field to given value.

### HasRequestable

`func (o *GetRoleDetailsResponse) HasRequestable() bool`

HasRequestable returns a boolean if a field has been set.

### GetDisplayname

`func (o *GetRoleDetailsResponse) GetDisplayname() string`

GetDisplayname returns the Displayname field if non-nil, zero value otherwise.

### GetDisplaynameOk

`func (o *GetRoleDetailsResponse) GetDisplaynameOk() (*string, bool)`

GetDisplaynameOk returns a tuple with the Displayname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayname

`func (o *GetRoleDetailsResponse) SetDisplayname(v string)`

SetDisplayname sets Displayname field to given value.

### HasDisplayname

`func (o *GetRoleDetailsResponse) HasDisplayname() bool`

HasDisplayname returns a boolean if a field has been set.

### GetUpdateuser

`func (o *GetRoleDetailsResponse) GetUpdateuser() string`

GetUpdateuser returns the Updateuser field if non-nil, zero value otherwise.

### GetUpdateuserOk

`func (o *GetRoleDetailsResponse) GetUpdateuserOk() (*string, bool)`

GetUpdateuserOk returns a tuple with the Updateuser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateuser

`func (o *GetRoleDetailsResponse) SetUpdateuser(v string)`

SetUpdateuser sets Updateuser field to given value.

### HasUpdateuser

`func (o *GetRoleDetailsResponse) HasUpdateuser() bool`

HasUpdateuser returns a boolean if a field has been set.

### GetEntitlementValueKey

`func (o *GetRoleDetailsResponse) GetEntitlementValueKey() string`

GetEntitlementValueKey returns the EntitlementValueKey field if non-nil, zero value otherwise.

### GetEntitlementValueKeyOk

`func (o *GetRoleDetailsResponse) GetEntitlementValueKeyOk() (*string, bool)`

GetEntitlementValueKeyOk returns a tuple with the EntitlementValueKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitlementValueKey

`func (o *GetRoleDetailsResponse) SetEntitlementValueKey(v string)`

SetEntitlementValueKey sets EntitlementValueKey field to given value.

### HasEntitlementValueKey

`func (o *GetRoleDetailsResponse) HasEntitlementValueKey() bool`

HasEntitlementValueKey returns a boolean if a field has been set.

### GetRoleState

`func (o *GetRoleDetailsResponse) GetRoleState() string`

GetRoleState returns the RoleState field if non-nil, zero value otherwise.

### GetRoleStateOk

`func (o *GetRoleDetailsResponse) GetRoleStateOk() (*string, bool)`

GetRoleStateOk returns a tuple with the RoleState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleState

`func (o *GetRoleDetailsResponse) SetRoleState(v string)`

SetRoleState sets RoleState field to given value.

### HasRoleState

`func (o *GetRoleDetailsResponse) HasRoleState() bool`

HasRoleState returns a boolean if a field has been set.

### GetEndpointkey

`func (o *GetRoleDetailsResponse) GetEndpointkey() string`

GetEndpointkey returns the Endpointkey field if non-nil, zero value otherwise.

### GetEndpointkeyOk

`func (o *GetRoleDetailsResponse) GetEndpointkeyOk() (*string, bool)`

GetEndpointkeyOk returns a tuple with the Endpointkey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpointkey

`func (o *GetRoleDetailsResponse) SetEndpointkey(v string)`

SetEndpointkey sets Endpointkey field to given value.

### HasEndpointkey

`func (o *GetRoleDetailsResponse) HasEndpointkey() bool`

HasEndpointkey returns a boolean if a field has been set.

### GetLastReviewedCampaignName

`func (o *GetRoleDetailsResponse) GetLastReviewedCampaignName() string`

GetLastReviewedCampaignName returns the LastReviewedCampaignName field if non-nil, zero value otherwise.

### GetLastReviewedCampaignNameOk

`func (o *GetRoleDetailsResponse) GetLastReviewedCampaignNameOk() (*string, bool)`

GetLastReviewedCampaignNameOk returns a tuple with the LastReviewedCampaignName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastReviewedCampaignName

`func (o *GetRoleDetailsResponse) SetLastReviewedCampaignName(v string)`

SetLastReviewedCampaignName sets LastReviewedCampaignName field to given value.

### HasLastReviewedCampaignName

`func (o *GetRoleDetailsResponse) HasLastReviewedCampaignName() bool`

HasLastReviewedCampaignName returns a boolean if a field has been set.

### GetRisk

`func (o *GetRoleDetailsResponse) GetRisk() string`

GetRisk returns the Risk field if non-nil, zero value otherwise.

### GetRiskOk

`func (o *GetRoleDetailsResponse) GetRiskOk() (*string, bool)`

GetRiskOk returns a tuple with the Risk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRisk

`func (o *GetRoleDetailsResponse) SetRisk(v string)`

SetRisk sets Risk field to given value.

### HasRisk

`func (o *GetRoleDetailsResponse) HasRisk() bool`

HasRisk returns a boolean if a field has been set.

### GetLastReviewedBy

`func (o *GetRoleDetailsResponse) GetLastReviewedBy() string`

GetLastReviewedBy returns the LastReviewedBy field if non-nil, zero value otherwise.

### GetLastReviewedByOk

`func (o *GetRoleDetailsResponse) GetLastReviewedByOk() (*string, bool)`

GetLastReviewedByOk returns a tuple with the LastReviewedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastReviewedBy

`func (o *GetRoleDetailsResponse) SetLastReviewedBy(v string)`

SetLastReviewedBy sets LastReviewedBy field to given value.

### HasLastReviewedBy

`func (o *GetRoleDetailsResponse) HasLastReviewedBy() bool`

HasLastReviewedBy returns a boolean if a field has been set.

### GetOwner

`func (o *GetRoleDetailsResponse) GetOwner() GetRoleDetailsResponseOwner`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *GetRoleDetailsResponse) GetOwnerOk() (*GetRoleDetailsResponseOwner, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *GetRoleDetailsResponse) SetOwner(v GetRoleDetailsResponseOwner)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *GetRoleDetailsResponse) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetUserDetails

`func (o *GetRoleDetailsResponse) GetUserDetails() []GetRoleDetailsResponseUserDetailsInner`

GetUserDetails returns the UserDetails field if non-nil, zero value otherwise.

### GetUserDetailsOk

`func (o *GetRoleDetailsResponse) GetUserDetailsOk() (*[]GetRoleDetailsResponseUserDetailsInner, bool)`

GetUserDetailsOk returns a tuple with the UserDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserDetails

`func (o *GetRoleDetailsResponse) SetUserDetails(v []GetRoleDetailsResponseUserDetailsInner)`

SetUserDetails sets UserDetails field to given value.

### HasUserDetails

`func (o *GetRoleDetailsResponse) HasUserDetails() bool`

HasUserDetails returns a boolean if a field has been set.

### GetEntitlementDetails

`func (o *GetRoleDetailsResponse) GetEntitlementDetails() []GetEntitlementDetailsResponse`

GetEntitlementDetails returns the EntitlementDetails field if non-nil, zero value otherwise.

### GetEntitlementDetailsOk

`func (o *GetRoleDetailsResponse) GetEntitlementDetailsOk() (*[]GetEntitlementDetailsResponse, bool)`

GetEntitlementDetailsOk returns a tuple with the EntitlementDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitlementDetails

`func (o *GetRoleDetailsResponse) SetEntitlementDetails(v []GetEntitlementDetailsResponse)`

SetEntitlementDetails sets EntitlementDetails field to given value.

### HasEntitlementDetails

`func (o *GetRoleDetailsResponse) HasEntitlementDetails() bool`

HasEntitlementDetails returns a boolean if a field has been set.

### GetCustomProperty1

`func (o *GetRoleDetailsResponse) GetCustomProperty1() string`

GetCustomProperty1 returns the CustomProperty1 field if non-nil, zero value otherwise.

### GetCustomProperty1Ok

`func (o *GetRoleDetailsResponse) GetCustomProperty1Ok() (*string, bool)`

GetCustomProperty1Ok returns a tuple with the CustomProperty1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomProperty1

`func (o *GetRoleDetailsResponse) SetCustomProperty1(v string)`

SetCustomProperty1 sets CustomProperty1 field to given value.

### HasCustomProperty1

`func (o *GetRoleDetailsResponse) HasCustomProperty1() bool`

HasCustomProperty1 returns a boolean if a field has been set.

### GetCustomProperty2

`func (o *GetRoleDetailsResponse) GetCustomProperty2() string`

GetCustomProperty2 returns the CustomProperty2 field if non-nil, zero value otherwise.

### GetCustomProperty2Ok

`func (o *GetRoleDetailsResponse) GetCustomProperty2Ok() (*string, bool)`

GetCustomProperty2Ok returns a tuple with the CustomProperty2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomProperty2

`func (o *GetRoleDetailsResponse) SetCustomProperty2(v string)`

SetCustomProperty2 sets CustomProperty2 field to given value.

### HasCustomProperty2

`func (o *GetRoleDetailsResponse) HasCustomProperty2() bool`

HasCustomProperty2 returns a boolean if a field has been set.

### GetCustomProperty3

`func (o *GetRoleDetailsResponse) GetCustomProperty3() string`

GetCustomProperty3 returns the CustomProperty3 field if non-nil, zero value otherwise.

### GetCustomProperty3Ok

`func (o *GetRoleDetailsResponse) GetCustomProperty3Ok() (*string, bool)`

GetCustomProperty3Ok returns a tuple with the CustomProperty3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomProperty3

`func (o *GetRoleDetailsResponse) SetCustomProperty3(v string)`

SetCustomProperty3 sets CustomProperty3 field to given value.

### HasCustomProperty3

`func (o *GetRoleDetailsResponse) HasCustomProperty3() bool`

HasCustomProperty3 returns a boolean if a field has been set.

### GetCustomProperty4

`func (o *GetRoleDetailsResponse) GetCustomProperty4() string`

GetCustomProperty4 returns the CustomProperty4 field if non-nil, zero value otherwise.

### GetCustomProperty4Ok

`func (o *GetRoleDetailsResponse) GetCustomProperty4Ok() (*string, bool)`

GetCustomProperty4Ok returns a tuple with the CustomProperty4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomProperty4

`func (o *GetRoleDetailsResponse) SetCustomProperty4(v string)`

SetCustomProperty4 sets CustomProperty4 field to given value.

### HasCustomProperty4

`func (o *GetRoleDetailsResponse) HasCustomProperty4() bool`

HasCustomProperty4 returns a boolean if a field has been set.

### GetCustomProperty5

`func (o *GetRoleDetailsResponse) GetCustomProperty5() string`

GetCustomProperty5 returns the CustomProperty5 field if non-nil, zero value otherwise.

### GetCustomProperty5Ok

`func (o *GetRoleDetailsResponse) GetCustomProperty5Ok() (*string, bool)`

GetCustomProperty5Ok returns a tuple with the CustomProperty5 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomProperty5

`func (o *GetRoleDetailsResponse) SetCustomProperty5(v string)`

SetCustomProperty5 sets CustomProperty5 field to given value.

### HasCustomProperty5

`func (o *GetRoleDetailsResponse) HasCustomProperty5() bool`

HasCustomProperty5 returns a boolean if a field has been set.

### GetCustomProperty6

`func (o *GetRoleDetailsResponse) GetCustomProperty6() string`

GetCustomProperty6 returns the CustomProperty6 field if non-nil, zero value otherwise.

### GetCustomProperty6Ok

`func (o *GetRoleDetailsResponse) GetCustomProperty6Ok() (*string, bool)`

GetCustomProperty6Ok returns a tuple with the CustomProperty6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomProperty6

`func (o *GetRoleDetailsResponse) SetCustomProperty6(v string)`

SetCustomProperty6 sets CustomProperty6 field to given value.

### HasCustomProperty6

`func (o *GetRoleDetailsResponse) HasCustomProperty6() bool`

HasCustomProperty6 returns a boolean if a field has been set.

### GetCustomProperty7

`func (o *GetRoleDetailsResponse) GetCustomProperty7() string`

GetCustomProperty7 returns the CustomProperty7 field if non-nil, zero value otherwise.

### GetCustomProperty7Ok

`func (o *GetRoleDetailsResponse) GetCustomProperty7Ok() (*string, bool)`

GetCustomProperty7Ok returns a tuple with the CustomProperty7 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomProperty7

`func (o *GetRoleDetailsResponse) SetCustomProperty7(v string)`

SetCustomProperty7 sets CustomProperty7 field to given value.

### HasCustomProperty7

`func (o *GetRoleDetailsResponse) HasCustomProperty7() bool`

HasCustomProperty7 returns a boolean if a field has been set.

### GetCustomProperty8

`func (o *GetRoleDetailsResponse) GetCustomProperty8() string`

GetCustomProperty8 returns the CustomProperty8 field if non-nil, zero value otherwise.

### GetCustomProperty8Ok

`func (o *GetRoleDetailsResponse) GetCustomProperty8Ok() (*string, bool)`

GetCustomProperty8Ok returns a tuple with the CustomProperty8 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomProperty8

`func (o *GetRoleDetailsResponse) SetCustomProperty8(v string)`

SetCustomProperty8 sets CustomProperty8 field to given value.

### HasCustomProperty8

`func (o *GetRoleDetailsResponse) HasCustomProperty8() bool`

HasCustomProperty8 returns a boolean if a field has been set.

### GetCustomProperty9

`func (o *GetRoleDetailsResponse) GetCustomProperty9() string`

GetCustomProperty9 returns the CustomProperty9 field if non-nil, zero value otherwise.

### GetCustomProperty9Ok

`func (o *GetRoleDetailsResponse) GetCustomProperty9Ok() (*string, bool)`

GetCustomProperty9Ok returns a tuple with the CustomProperty9 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomProperty9

`func (o *GetRoleDetailsResponse) SetCustomProperty9(v string)`

SetCustomProperty9 sets CustomProperty9 field to given value.

### HasCustomProperty9

`func (o *GetRoleDetailsResponse) HasCustomProperty9() bool`

HasCustomProperty9 returns a boolean if a field has been set.

### GetCustomProperty10

`func (o *GetRoleDetailsResponse) GetCustomProperty10() string`

GetCustomProperty10 returns the CustomProperty10 field if non-nil, zero value otherwise.

### GetCustomProperty10Ok

`func (o *GetRoleDetailsResponse) GetCustomProperty10Ok() (*string, bool)`

GetCustomProperty10Ok returns a tuple with the CustomProperty10 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomProperty10

`func (o *GetRoleDetailsResponse) SetCustomProperty10(v string)`

SetCustomProperty10 sets CustomProperty10 field to given value.

### HasCustomProperty10

`func (o *GetRoleDetailsResponse) HasCustomProperty10() bool`

HasCustomProperty10 returns a boolean if a field has been set.

### GetCustomProperty11

`func (o *GetRoleDetailsResponse) GetCustomProperty11() string`

GetCustomProperty11 returns the CustomProperty11 field if non-nil, zero value otherwise.

### GetCustomProperty11Ok

`func (o *GetRoleDetailsResponse) GetCustomProperty11Ok() (*string, bool)`

GetCustomProperty11Ok returns a tuple with the CustomProperty11 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomProperty11

`func (o *GetRoleDetailsResponse) SetCustomProperty11(v string)`

SetCustomProperty11 sets CustomProperty11 field to given value.

### HasCustomProperty11

`func (o *GetRoleDetailsResponse) HasCustomProperty11() bool`

HasCustomProperty11 returns a boolean if a field has been set.

### GetCustomProperty12

`func (o *GetRoleDetailsResponse) GetCustomProperty12() string`

GetCustomProperty12 returns the CustomProperty12 field if non-nil, zero value otherwise.

### GetCustomProperty12Ok

`func (o *GetRoleDetailsResponse) GetCustomProperty12Ok() (*string, bool)`

GetCustomProperty12Ok returns a tuple with the CustomProperty12 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomProperty12

`func (o *GetRoleDetailsResponse) SetCustomProperty12(v string)`

SetCustomProperty12 sets CustomProperty12 field to given value.

### HasCustomProperty12

`func (o *GetRoleDetailsResponse) HasCustomProperty12() bool`

HasCustomProperty12 returns a boolean if a field has been set.

### GetCustomProperty13

`func (o *GetRoleDetailsResponse) GetCustomProperty13() string`

GetCustomProperty13 returns the CustomProperty13 field if non-nil, zero value otherwise.

### GetCustomProperty13Ok

`func (o *GetRoleDetailsResponse) GetCustomProperty13Ok() (*string, bool)`

GetCustomProperty13Ok returns a tuple with the CustomProperty13 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomProperty13

`func (o *GetRoleDetailsResponse) SetCustomProperty13(v string)`

SetCustomProperty13 sets CustomProperty13 field to given value.

### HasCustomProperty13

`func (o *GetRoleDetailsResponse) HasCustomProperty13() bool`

HasCustomProperty13 returns a boolean if a field has been set.

### GetCustomProperty14

`func (o *GetRoleDetailsResponse) GetCustomProperty14() string`

GetCustomProperty14 returns the CustomProperty14 field if non-nil, zero value otherwise.

### GetCustomProperty14Ok

`func (o *GetRoleDetailsResponse) GetCustomProperty14Ok() (*string, bool)`

GetCustomProperty14Ok returns a tuple with the CustomProperty14 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomProperty14

`func (o *GetRoleDetailsResponse) SetCustomProperty14(v string)`

SetCustomProperty14 sets CustomProperty14 field to given value.

### HasCustomProperty14

`func (o *GetRoleDetailsResponse) HasCustomProperty14() bool`

HasCustomProperty14 returns a boolean if a field has been set.

### GetCustomProperty15

`func (o *GetRoleDetailsResponse) GetCustomProperty15() string`

GetCustomProperty15 returns the CustomProperty15 field if non-nil, zero value otherwise.

### GetCustomProperty15Ok

`func (o *GetRoleDetailsResponse) GetCustomProperty15Ok() (*string, bool)`

GetCustomProperty15Ok returns a tuple with the CustomProperty15 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomProperty15

`func (o *GetRoleDetailsResponse) SetCustomProperty15(v string)`

SetCustomProperty15 sets CustomProperty15 field to given value.

### HasCustomProperty15

`func (o *GetRoleDetailsResponse) HasCustomProperty15() bool`

HasCustomProperty15 returns a boolean if a field has been set.

### GetCustomProperty16

`func (o *GetRoleDetailsResponse) GetCustomProperty16() string`

GetCustomProperty16 returns the CustomProperty16 field if non-nil, zero value otherwise.

### GetCustomProperty16Ok

`func (o *GetRoleDetailsResponse) GetCustomProperty16Ok() (*string, bool)`

GetCustomProperty16Ok returns a tuple with the CustomProperty16 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomProperty16

`func (o *GetRoleDetailsResponse) SetCustomProperty16(v string)`

SetCustomProperty16 sets CustomProperty16 field to given value.

### HasCustomProperty16

`func (o *GetRoleDetailsResponse) HasCustomProperty16() bool`

HasCustomProperty16 returns a boolean if a field has been set.

### GetCustomProperty17

`func (o *GetRoleDetailsResponse) GetCustomProperty17() string`

GetCustomProperty17 returns the CustomProperty17 field if non-nil, zero value otherwise.

### GetCustomProperty17Ok

`func (o *GetRoleDetailsResponse) GetCustomProperty17Ok() (*string, bool)`

GetCustomProperty17Ok returns a tuple with the CustomProperty17 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomProperty17

`func (o *GetRoleDetailsResponse) SetCustomProperty17(v string)`

SetCustomProperty17 sets CustomProperty17 field to given value.

### HasCustomProperty17

`func (o *GetRoleDetailsResponse) HasCustomProperty17() bool`

HasCustomProperty17 returns a boolean if a field has been set.

### GetCustomProperty18

`func (o *GetRoleDetailsResponse) GetCustomProperty18() string`

GetCustomProperty18 returns the CustomProperty18 field if non-nil, zero value otherwise.

### GetCustomProperty18Ok

`func (o *GetRoleDetailsResponse) GetCustomProperty18Ok() (*string, bool)`

GetCustomProperty18Ok returns a tuple with the CustomProperty18 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomProperty18

`func (o *GetRoleDetailsResponse) SetCustomProperty18(v string)`

SetCustomProperty18 sets CustomProperty18 field to given value.

### HasCustomProperty18

`func (o *GetRoleDetailsResponse) HasCustomProperty18() bool`

HasCustomProperty18 returns a boolean if a field has been set.

### GetCustomProperty19

`func (o *GetRoleDetailsResponse) GetCustomProperty19() string`

GetCustomProperty19 returns the CustomProperty19 field if non-nil, zero value otherwise.

### GetCustomProperty19Ok

`func (o *GetRoleDetailsResponse) GetCustomProperty19Ok() (*string, bool)`

GetCustomProperty19Ok returns a tuple with the CustomProperty19 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomProperty19

`func (o *GetRoleDetailsResponse) SetCustomProperty19(v string)`

SetCustomProperty19 sets CustomProperty19 field to given value.

### HasCustomProperty19

`func (o *GetRoleDetailsResponse) HasCustomProperty19() bool`

HasCustomProperty19 returns a boolean if a field has been set.

### GetCustomProperty20

`func (o *GetRoleDetailsResponse) GetCustomProperty20() string`

GetCustomProperty20 returns the CustomProperty20 field if non-nil, zero value otherwise.

### GetCustomProperty20Ok

`func (o *GetRoleDetailsResponse) GetCustomProperty20Ok() (*string, bool)`

GetCustomProperty20Ok returns a tuple with the CustomProperty20 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomProperty20

`func (o *GetRoleDetailsResponse) SetCustomProperty20(v string)`

SetCustomProperty20 sets CustomProperty20 field to given value.

### HasCustomProperty20

`func (o *GetRoleDetailsResponse) HasCustomProperty20() bool`

HasCustomProperty20 returns a boolean if a field has been set.

### GetCustomProperty21

`func (o *GetRoleDetailsResponse) GetCustomProperty21() string`

GetCustomProperty21 returns the CustomProperty21 field if non-nil, zero value otherwise.

### GetCustomProperty21Ok

`func (o *GetRoleDetailsResponse) GetCustomProperty21Ok() (*string, bool)`

GetCustomProperty21Ok returns a tuple with the CustomProperty21 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomProperty21

`func (o *GetRoleDetailsResponse) SetCustomProperty21(v string)`

SetCustomProperty21 sets CustomProperty21 field to given value.

### HasCustomProperty21

`func (o *GetRoleDetailsResponse) HasCustomProperty21() bool`

HasCustomProperty21 returns a boolean if a field has been set.

### GetCustomProperty22

`func (o *GetRoleDetailsResponse) GetCustomProperty22() string`

GetCustomProperty22 returns the CustomProperty22 field if non-nil, zero value otherwise.

### GetCustomProperty22Ok

`func (o *GetRoleDetailsResponse) GetCustomProperty22Ok() (*string, bool)`

GetCustomProperty22Ok returns a tuple with the CustomProperty22 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomProperty22

`func (o *GetRoleDetailsResponse) SetCustomProperty22(v string)`

SetCustomProperty22 sets CustomProperty22 field to given value.

### HasCustomProperty22

`func (o *GetRoleDetailsResponse) HasCustomProperty22() bool`

HasCustomProperty22 returns a boolean if a field has been set.

### GetCustomProperty23

`func (o *GetRoleDetailsResponse) GetCustomProperty23() string`

GetCustomProperty23 returns the CustomProperty23 field if non-nil, zero value otherwise.

### GetCustomProperty23Ok

`func (o *GetRoleDetailsResponse) GetCustomProperty23Ok() (*string, bool)`

GetCustomProperty23Ok returns a tuple with the CustomProperty23 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomProperty23

`func (o *GetRoleDetailsResponse) SetCustomProperty23(v string)`

SetCustomProperty23 sets CustomProperty23 field to given value.

### HasCustomProperty23

`func (o *GetRoleDetailsResponse) HasCustomProperty23() bool`

HasCustomProperty23 returns a boolean if a field has been set.

### GetCustomProperty24

`func (o *GetRoleDetailsResponse) GetCustomProperty24() string`

GetCustomProperty24 returns the CustomProperty24 field if non-nil, zero value otherwise.

### GetCustomProperty24Ok

`func (o *GetRoleDetailsResponse) GetCustomProperty24Ok() (*string, bool)`

GetCustomProperty24Ok returns a tuple with the CustomProperty24 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomProperty24

`func (o *GetRoleDetailsResponse) SetCustomProperty24(v string)`

SetCustomProperty24 sets CustomProperty24 field to given value.

### HasCustomProperty24

`func (o *GetRoleDetailsResponse) HasCustomProperty24() bool`

HasCustomProperty24 returns a boolean if a field has been set.

### GetCustomProperty25

`func (o *GetRoleDetailsResponse) GetCustomProperty25() string`

GetCustomProperty25 returns the CustomProperty25 field if non-nil, zero value otherwise.

### GetCustomProperty25Ok

`func (o *GetRoleDetailsResponse) GetCustomProperty25Ok() (*string, bool)`

GetCustomProperty25Ok returns a tuple with the CustomProperty25 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomProperty25

`func (o *GetRoleDetailsResponse) SetCustomProperty25(v string)`

SetCustomProperty25 sets CustomProperty25 field to given value.

### HasCustomProperty25

`func (o *GetRoleDetailsResponse) HasCustomProperty25() bool`

HasCustomProperty25 returns a boolean if a field has been set.

### GetCustomProperty26

`func (o *GetRoleDetailsResponse) GetCustomProperty26() string`

GetCustomProperty26 returns the CustomProperty26 field if non-nil, zero value otherwise.

### GetCustomProperty26Ok

`func (o *GetRoleDetailsResponse) GetCustomProperty26Ok() (*string, bool)`

GetCustomProperty26Ok returns a tuple with the CustomProperty26 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomProperty26

`func (o *GetRoleDetailsResponse) SetCustomProperty26(v string)`

SetCustomProperty26 sets CustomProperty26 field to given value.

### HasCustomProperty26

`func (o *GetRoleDetailsResponse) HasCustomProperty26() bool`

HasCustomProperty26 returns a boolean if a field has been set.

### GetCustomProperty27

`func (o *GetRoleDetailsResponse) GetCustomProperty27() string`

GetCustomProperty27 returns the CustomProperty27 field if non-nil, zero value otherwise.

### GetCustomProperty27Ok

`func (o *GetRoleDetailsResponse) GetCustomProperty27Ok() (*string, bool)`

GetCustomProperty27Ok returns a tuple with the CustomProperty27 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomProperty27

`func (o *GetRoleDetailsResponse) SetCustomProperty27(v string)`

SetCustomProperty27 sets CustomProperty27 field to given value.

### HasCustomProperty27

`func (o *GetRoleDetailsResponse) HasCustomProperty27() bool`

HasCustomProperty27 returns a boolean if a field has been set.

### GetCustomProperty28

`func (o *GetRoleDetailsResponse) GetCustomProperty28() string`

GetCustomProperty28 returns the CustomProperty28 field if non-nil, zero value otherwise.

### GetCustomProperty28Ok

`func (o *GetRoleDetailsResponse) GetCustomProperty28Ok() (*string, bool)`

GetCustomProperty28Ok returns a tuple with the CustomProperty28 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomProperty28

`func (o *GetRoleDetailsResponse) SetCustomProperty28(v string)`

SetCustomProperty28 sets CustomProperty28 field to given value.

### HasCustomProperty28

`func (o *GetRoleDetailsResponse) HasCustomProperty28() bool`

HasCustomProperty28 returns a boolean if a field has been set.

### GetCustomProperty29

`func (o *GetRoleDetailsResponse) GetCustomProperty29() string`

GetCustomProperty29 returns the CustomProperty29 field if non-nil, zero value otherwise.

### GetCustomProperty29Ok

`func (o *GetRoleDetailsResponse) GetCustomProperty29Ok() (*string, bool)`

GetCustomProperty29Ok returns a tuple with the CustomProperty29 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomProperty29

`func (o *GetRoleDetailsResponse) SetCustomProperty29(v string)`

SetCustomProperty29 sets CustomProperty29 field to given value.

### HasCustomProperty29

`func (o *GetRoleDetailsResponse) HasCustomProperty29() bool`

HasCustomProperty29 returns a boolean if a field has been set.

### GetCustomProperty30

`func (o *GetRoleDetailsResponse) GetCustomProperty30() string`

GetCustomProperty30 returns the CustomProperty30 field if non-nil, zero value otherwise.

### GetCustomProperty30Ok

`func (o *GetRoleDetailsResponse) GetCustomProperty30Ok() (*string, bool)`

GetCustomProperty30Ok returns a tuple with the CustomProperty30 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomProperty30

`func (o *GetRoleDetailsResponse) SetCustomProperty30(v string)`

SetCustomProperty30 sets CustomProperty30 field to given value.

### HasCustomProperty30

`func (o *GetRoleDetailsResponse) HasCustomProperty30() bool`

HasCustomProperty30 returns a boolean if a field has been set.

### GetCustomProperty31

`func (o *GetRoleDetailsResponse) GetCustomProperty31() string`

GetCustomProperty31 returns the CustomProperty31 field if non-nil, zero value otherwise.

### GetCustomProperty31Ok

`func (o *GetRoleDetailsResponse) GetCustomProperty31Ok() (*string, bool)`

GetCustomProperty31Ok returns a tuple with the CustomProperty31 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomProperty31

`func (o *GetRoleDetailsResponse) SetCustomProperty31(v string)`

SetCustomProperty31 sets CustomProperty31 field to given value.

### HasCustomProperty31

`func (o *GetRoleDetailsResponse) HasCustomProperty31() bool`

HasCustomProperty31 returns a boolean if a field has been set.

### GetCustomProperty32

`func (o *GetRoleDetailsResponse) GetCustomProperty32() string`

GetCustomProperty32 returns the CustomProperty32 field if non-nil, zero value otherwise.

### GetCustomProperty32Ok

`func (o *GetRoleDetailsResponse) GetCustomProperty32Ok() (*string, bool)`

GetCustomProperty32Ok returns a tuple with the CustomProperty32 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomProperty32

`func (o *GetRoleDetailsResponse) SetCustomProperty32(v string)`

SetCustomProperty32 sets CustomProperty32 field to given value.

### HasCustomProperty32

`func (o *GetRoleDetailsResponse) HasCustomProperty32() bool`

HasCustomProperty32 returns a boolean if a field has been set.

### GetCustomProperty33

`func (o *GetRoleDetailsResponse) GetCustomProperty33() string`

GetCustomProperty33 returns the CustomProperty33 field if non-nil, zero value otherwise.

### GetCustomProperty33Ok

`func (o *GetRoleDetailsResponse) GetCustomProperty33Ok() (*string, bool)`

GetCustomProperty33Ok returns a tuple with the CustomProperty33 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomProperty33

`func (o *GetRoleDetailsResponse) SetCustomProperty33(v string)`

SetCustomProperty33 sets CustomProperty33 field to given value.

### HasCustomProperty33

`func (o *GetRoleDetailsResponse) HasCustomProperty33() bool`

HasCustomProperty33 returns a boolean if a field has been set.

### GetCustomProperty34

`func (o *GetRoleDetailsResponse) GetCustomProperty34() string`

GetCustomProperty34 returns the CustomProperty34 field if non-nil, zero value otherwise.

### GetCustomProperty34Ok

`func (o *GetRoleDetailsResponse) GetCustomProperty34Ok() (*string, bool)`

GetCustomProperty34Ok returns a tuple with the CustomProperty34 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomProperty34

`func (o *GetRoleDetailsResponse) SetCustomProperty34(v string)`

SetCustomProperty34 sets CustomProperty34 field to given value.

### HasCustomProperty34

`func (o *GetRoleDetailsResponse) HasCustomProperty34() bool`

HasCustomProperty34 returns a boolean if a field has been set.

### GetCustomProperty35

`func (o *GetRoleDetailsResponse) GetCustomProperty35() string`

GetCustomProperty35 returns the CustomProperty35 field if non-nil, zero value otherwise.

### GetCustomProperty35Ok

`func (o *GetRoleDetailsResponse) GetCustomProperty35Ok() (*string, bool)`

GetCustomProperty35Ok returns a tuple with the CustomProperty35 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomProperty35

`func (o *GetRoleDetailsResponse) SetCustomProperty35(v string)`

SetCustomProperty35 sets CustomProperty35 field to given value.

### HasCustomProperty35

`func (o *GetRoleDetailsResponse) HasCustomProperty35() bool`

HasCustomProperty35 returns a boolean if a field has been set.

### GetCustomProperty36

`func (o *GetRoleDetailsResponse) GetCustomProperty36() string`

GetCustomProperty36 returns the CustomProperty36 field if non-nil, zero value otherwise.

### GetCustomProperty36Ok

`func (o *GetRoleDetailsResponse) GetCustomProperty36Ok() (*string, bool)`

GetCustomProperty36Ok returns a tuple with the CustomProperty36 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomProperty36

`func (o *GetRoleDetailsResponse) SetCustomProperty36(v string)`

SetCustomProperty36 sets CustomProperty36 field to given value.

### HasCustomProperty36

`func (o *GetRoleDetailsResponse) HasCustomProperty36() bool`

HasCustomProperty36 returns a boolean if a field has been set.

### GetCustomProperty37

`func (o *GetRoleDetailsResponse) GetCustomProperty37() string`

GetCustomProperty37 returns the CustomProperty37 field if non-nil, zero value otherwise.

### GetCustomProperty37Ok

`func (o *GetRoleDetailsResponse) GetCustomProperty37Ok() (*string, bool)`

GetCustomProperty37Ok returns a tuple with the CustomProperty37 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomProperty37

`func (o *GetRoleDetailsResponse) SetCustomProperty37(v string)`

SetCustomProperty37 sets CustomProperty37 field to given value.

### HasCustomProperty37

`func (o *GetRoleDetailsResponse) HasCustomProperty37() bool`

HasCustomProperty37 returns a boolean if a field has been set.

### GetCustomProperty38

`func (o *GetRoleDetailsResponse) GetCustomProperty38() string`

GetCustomProperty38 returns the CustomProperty38 field if non-nil, zero value otherwise.

### GetCustomProperty38Ok

`func (o *GetRoleDetailsResponse) GetCustomProperty38Ok() (*string, bool)`

GetCustomProperty38Ok returns a tuple with the CustomProperty38 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomProperty38

`func (o *GetRoleDetailsResponse) SetCustomProperty38(v string)`

SetCustomProperty38 sets CustomProperty38 field to given value.

### HasCustomProperty38

`func (o *GetRoleDetailsResponse) HasCustomProperty38() bool`

HasCustomProperty38 returns a boolean if a field has been set.

### GetCustomProperty39

`func (o *GetRoleDetailsResponse) GetCustomProperty39() string`

GetCustomProperty39 returns the CustomProperty39 field if non-nil, zero value otherwise.

### GetCustomProperty39Ok

`func (o *GetRoleDetailsResponse) GetCustomProperty39Ok() (*string, bool)`

GetCustomProperty39Ok returns a tuple with the CustomProperty39 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomProperty39

`func (o *GetRoleDetailsResponse) SetCustomProperty39(v string)`

SetCustomProperty39 sets CustomProperty39 field to given value.

### HasCustomProperty39

`func (o *GetRoleDetailsResponse) HasCustomProperty39() bool`

HasCustomProperty39 returns a boolean if a field has been set.

### GetCustomProperty40

`func (o *GetRoleDetailsResponse) GetCustomProperty40() string`

GetCustomProperty40 returns the CustomProperty40 field if non-nil, zero value otherwise.

### GetCustomProperty40Ok

`func (o *GetRoleDetailsResponse) GetCustomProperty40Ok() (*string, bool)`

GetCustomProperty40Ok returns a tuple with the CustomProperty40 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomProperty40

`func (o *GetRoleDetailsResponse) SetCustomProperty40(v string)`

SetCustomProperty40 sets CustomProperty40 field to given value.

### HasCustomProperty40

`func (o *GetRoleDetailsResponse) HasCustomProperty40() bool`

HasCustomProperty40 returns a boolean if a field has been set.

### GetCustomProperty41

`func (o *GetRoleDetailsResponse) GetCustomProperty41() string`

GetCustomProperty41 returns the CustomProperty41 field if non-nil, zero value otherwise.

### GetCustomProperty41Ok

`func (o *GetRoleDetailsResponse) GetCustomProperty41Ok() (*string, bool)`

GetCustomProperty41Ok returns a tuple with the CustomProperty41 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomProperty41

`func (o *GetRoleDetailsResponse) SetCustomProperty41(v string)`

SetCustomProperty41 sets CustomProperty41 field to given value.

### HasCustomProperty41

`func (o *GetRoleDetailsResponse) HasCustomProperty41() bool`

HasCustomProperty41 returns a boolean if a field has been set.

### GetCustomProperty42

`func (o *GetRoleDetailsResponse) GetCustomProperty42() string`

GetCustomProperty42 returns the CustomProperty42 field if non-nil, zero value otherwise.

### GetCustomProperty42Ok

`func (o *GetRoleDetailsResponse) GetCustomProperty42Ok() (*string, bool)`

GetCustomProperty42Ok returns a tuple with the CustomProperty42 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomProperty42

`func (o *GetRoleDetailsResponse) SetCustomProperty42(v string)`

SetCustomProperty42 sets CustomProperty42 field to given value.

### HasCustomProperty42

`func (o *GetRoleDetailsResponse) HasCustomProperty42() bool`

HasCustomProperty42 returns a boolean if a field has been set.

### GetCustomProperty43

`func (o *GetRoleDetailsResponse) GetCustomProperty43() string`

GetCustomProperty43 returns the CustomProperty43 field if non-nil, zero value otherwise.

### GetCustomProperty43Ok

`func (o *GetRoleDetailsResponse) GetCustomProperty43Ok() (*string, bool)`

GetCustomProperty43Ok returns a tuple with the CustomProperty43 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomProperty43

`func (o *GetRoleDetailsResponse) SetCustomProperty43(v string)`

SetCustomProperty43 sets CustomProperty43 field to given value.

### HasCustomProperty43

`func (o *GetRoleDetailsResponse) HasCustomProperty43() bool`

HasCustomProperty43 returns a boolean if a field has been set.

### GetCustomProperty44

`func (o *GetRoleDetailsResponse) GetCustomProperty44() string`

GetCustomProperty44 returns the CustomProperty44 field if non-nil, zero value otherwise.

### GetCustomProperty44Ok

`func (o *GetRoleDetailsResponse) GetCustomProperty44Ok() (*string, bool)`

GetCustomProperty44Ok returns a tuple with the CustomProperty44 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomProperty44

`func (o *GetRoleDetailsResponse) SetCustomProperty44(v string)`

SetCustomProperty44 sets CustomProperty44 field to given value.

### HasCustomProperty44

`func (o *GetRoleDetailsResponse) HasCustomProperty44() bool`

HasCustomProperty44 returns a boolean if a field has been set.

### GetCustomProperty45

`func (o *GetRoleDetailsResponse) GetCustomProperty45() string`

GetCustomProperty45 returns the CustomProperty45 field if non-nil, zero value otherwise.

### GetCustomProperty45Ok

`func (o *GetRoleDetailsResponse) GetCustomProperty45Ok() (*string, bool)`

GetCustomProperty45Ok returns a tuple with the CustomProperty45 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomProperty45

`func (o *GetRoleDetailsResponse) SetCustomProperty45(v string)`

SetCustomProperty45 sets CustomProperty45 field to given value.

### HasCustomProperty45

`func (o *GetRoleDetailsResponse) HasCustomProperty45() bool`

HasCustomProperty45 returns a boolean if a field has been set.

### GetCustomProperty46

`func (o *GetRoleDetailsResponse) GetCustomProperty46() string`

GetCustomProperty46 returns the CustomProperty46 field if non-nil, zero value otherwise.

### GetCustomProperty46Ok

`func (o *GetRoleDetailsResponse) GetCustomProperty46Ok() (*string, bool)`

GetCustomProperty46Ok returns a tuple with the CustomProperty46 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomProperty46

`func (o *GetRoleDetailsResponse) SetCustomProperty46(v string)`

SetCustomProperty46 sets CustomProperty46 field to given value.

### HasCustomProperty46

`func (o *GetRoleDetailsResponse) HasCustomProperty46() bool`

HasCustomProperty46 returns a boolean if a field has been set.

### GetCustomProperty47

`func (o *GetRoleDetailsResponse) GetCustomProperty47() string`

GetCustomProperty47 returns the CustomProperty47 field if non-nil, zero value otherwise.

### GetCustomProperty47Ok

`func (o *GetRoleDetailsResponse) GetCustomProperty47Ok() (*string, bool)`

GetCustomProperty47Ok returns a tuple with the CustomProperty47 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomProperty47

`func (o *GetRoleDetailsResponse) SetCustomProperty47(v string)`

SetCustomProperty47 sets CustomProperty47 field to given value.

### HasCustomProperty47

`func (o *GetRoleDetailsResponse) HasCustomProperty47() bool`

HasCustomProperty47 returns a boolean if a field has been set.

### GetCustomProperty48

`func (o *GetRoleDetailsResponse) GetCustomProperty48() string`

GetCustomProperty48 returns the CustomProperty48 field if non-nil, zero value otherwise.

### GetCustomProperty48Ok

`func (o *GetRoleDetailsResponse) GetCustomProperty48Ok() (*string, bool)`

GetCustomProperty48Ok returns a tuple with the CustomProperty48 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomProperty48

`func (o *GetRoleDetailsResponse) SetCustomProperty48(v string)`

SetCustomProperty48 sets CustomProperty48 field to given value.

### HasCustomProperty48

`func (o *GetRoleDetailsResponse) HasCustomProperty48() bool`

HasCustomProperty48 returns a boolean if a field has been set.

### GetCustomProperty49

`func (o *GetRoleDetailsResponse) GetCustomProperty49() string`

GetCustomProperty49 returns the CustomProperty49 field if non-nil, zero value otherwise.

### GetCustomProperty49Ok

`func (o *GetRoleDetailsResponse) GetCustomProperty49Ok() (*string, bool)`

GetCustomProperty49Ok returns a tuple with the CustomProperty49 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomProperty49

`func (o *GetRoleDetailsResponse) SetCustomProperty49(v string)`

SetCustomProperty49 sets CustomProperty49 field to given value.

### HasCustomProperty49

`func (o *GetRoleDetailsResponse) HasCustomProperty49() bool`

HasCustomProperty49 returns a boolean if a field has been set.

### GetCustomProperty50

`func (o *GetRoleDetailsResponse) GetCustomProperty50() string`

GetCustomProperty50 returns the CustomProperty50 field if non-nil, zero value otherwise.

### GetCustomProperty50Ok

`func (o *GetRoleDetailsResponse) GetCustomProperty50Ok() (*string, bool)`

GetCustomProperty50Ok returns a tuple with the CustomProperty50 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomProperty50

`func (o *GetRoleDetailsResponse) SetCustomProperty50(v string)`

SetCustomProperty50 sets CustomProperty50 field to given value.

### HasCustomProperty50

`func (o *GetRoleDetailsResponse) HasCustomProperty50() bool`

HasCustomProperty50 returns a boolean if a field has been set.

### GetCustomProperty51

`func (o *GetRoleDetailsResponse) GetCustomProperty51() string`

GetCustomProperty51 returns the CustomProperty51 field if non-nil, zero value otherwise.

### GetCustomProperty51Ok

`func (o *GetRoleDetailsResponse) GetCustomProperty51Ok() (*string, bool)`

GetCustomProperty51Ok returns a tuple with the CustomProperty51 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomProperty51

`func (o *GetRoleDetailsResponse) SetCustomProperty51(v string)`

SetCustomProperty51 sets CustomProperty51 field to given value.

### HasCustomProperty51

`func (o *GetRoleDetailsResponse) HasCustomProperty51() bool`

HasCustomProperty51 returns a boolean if a field has been set.

### GetCustomProperty52

`func (o *GetRoleDetailsResponse) GetCustomProperty52() string`

GetCustomProperty52 returns the CustomProperty52 field if non-nil, zero value otherwise.

### GetCustomProperty52Ok

`func (o *GetRoleDetailsResponse) GetCustomProperty52Ok() (*string, bool)`

GetCustomProperty52Ok returns a tuple with the CustomProperty52 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomProperty52

`func (o *GetRoleDetailsResponse) SetCustomProperty52(v string)`

SetCustomProperty52 sets CustomProperty52 field to given value.

### HasCustomProperty52

`func (o *GetRoleDetailsResponse) HasCustomProperty52() bool`

HasCustomProperty52 returns a boolean if a field has been set.

### GetCustomProperty53

`func (o *GetRoleDetailsResponse) GetCustomProperty53() string`

GetCustomProperty53 returns the CustomProperty53 field if non-nil, zero value otherwise.

### GetCustomProperty53Ok

`func (o *GetRoleDetailsResponse) GetCustomProperty53Ok() (*string, bool)`

GetCustomProperty53Ok returns a tuple with the CustomProperty53 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomProperty53

`func (o *GetRoleDetailsResponse) SetCustomProperty53(v string)`

SetCustomProperty53 sets CustomProperty53 field to given value.

### HasCustomProperty53

`func (o *GetRoleDetailsResponse) HasCustomProperty53() bool`

HasCustomProperty53 returns a boolean if a field has been set.

### GetCustomProperty54

`func (o *GetRoleDetailsResponse) GetCustomProperty54() string`

GetCustomProperty54 returns the CustomProperty54 field if non-nil, zero value otherwise.

### GetCustomProperty54Ok

`func (o *GetRoleDetailsResponse) GetCustomProperty54Ok() (*string, bool)`

GetCustomProperty54Ok returns a tuple with the CustomProperty54 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomProperty54

`func (o *GetRoleDetailsResponse) SetCustomProperty54(v string)`

SetCustomProperty54 sets CustomProperty54 field to given value.

### HasCustomProperty54

`func (o *GetRoleDetailsResponse) HasCustomProperty54() bool`

HasCustomProperty54 returns a boolean if a field has been set.

### GetCustomProperty55

`func (o *GetRoleDetailsResponse) GetCustomProperty55() string`

GetCustomProperty55 returns the CustomProperty55 field if non-nil, zero value otherwise.

### GetCustomProperty55Ok

`func (o *GetRoleDetailsResponse) GetCustomProperty55Ok() (*string, bool)`

GetCustomProperty55Ok returns a tuple with the CustomProperty55 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomProperty55

`func (o *GetRoleDetailsResponse) SetCustomProperty55(v string)`

SetCustomProperty55 sets CustomProperty55 field to given value.

### HasCustomProperty55

`func (o *GetRoleDetailsResponse) HasCustomProperty55() bool`

HasCustomProperty55 returns a boolean if a field has been set.

### GetCustomProperty56

`func (o *GetRoleDetailsResponse) GetCustomProperty56() string`

GetCustomProperty56 returns the CustomProperty56 field if non-nil, zero value otherwise.

### GetCustomProperty56Ok

`func (o *GetRoleDetailsResponse) GetCustomProperty56Ok() (*string, bool)`

GetCustomProperty56Ok returns a tuple with the CustomProperty56 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomProperty56

`func (o *GetRoleDetailsResponse) SetCustomProperty56(v string)`

SetCustomProperty56 sets CustomProperty56 field to given value.

### HasCustomProperty56

`func (o *GetRoleDetailsResponse) HasCustomProperty56() bool`

HasCustomProperty56 returns a boolean if a field has been set.

### GetCustomProperty57

`func (o *GetRoleDetailsResponse) GetCustomProperty57() string`

GetCustomProperty57 returns the CustomProperty57 field if non-nil, zero value otherwise.

### GetCustomProperty57Ok

`func (o *GetRoleDetailsResponse) GetCustomProperty57Ok() (*string, bool)`

GetCustomProperty57Ok returns a tuple with the CustomProperty57 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomProperty57

`func (o *GetRoleDetailsResponse) SetCustomProperty57(v string)`

SetCustomProperty57 sets CustomProperty57 field to given value.

### HasCustomProperty57

`func (o *GetRoleDetailsResponse) HasCustomProperty57() bool`

HasCustomProperty57 returns a boolean if a field has been set.

### GetCustomProperty58

`func (o *GetRoleDetailsResponse) GetCustomProperty58() string`

GetCustomProperty58 returns the CustomProperty58 field if non-nil, zero value otherwise.

### GetCustomProperty58Ok

`func (o *GetRoleDetailsResponse) GetCustomProperty58Ok() (*string, bool)`

GetCustomProperty58Ok returns a tuple with the CustomProperty58 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomProperty58

`func (o *GetRoleDetailsResponse) SetCustomProperty58(v string)`

SetCustomProperty58 sets CustomProperty58 field to given value.

### HasCustomProperty58

`func (o *GetRoleDetailsResponse) HasCustomProperty58() bool`

HasCustomProperty58 returns a boolean if a field has been set.

### GetCustomProperty59

`func (o *GetRoleDetailsResponse) GetCustomProperty59() string`

GetCustomProperty59 returns the CustomProperty59 field if non-nil, zero value otherwise.

### GetCustomProperty59Ok

`func (o *GetRoleDetailsResponse) GetCustomProperty59Ok() (*string, bool)`

GetCustomProperty59Ok returns a tuple with the CustomProperty59 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomProperty59

`func (o *GetRoleDetailsResponse) SetCustomProperty59(v string)`

SetCustomProperty59 sets CustomProperty59 field to given value.

### HasCustomProperty59

`func (o *GetRoleDetailsResponse) HasCustomProperty59() bool`

HasCustomProperty59 returns a boolean if a field has been set.

### GetCustomProperty60

`func (o *GetRoleDetailsResponse) GetCustomProperty60() string`

GetCustomProperty60 returns the CustomProperty60 field if non-nil, zero value otherwise.

### GetCustomProperty60Ok

`func (o *GetRoleDetailsResponse) GetCustomProperty60Ok() (*string, bool)`

GetCustomProperty60Ok returns a tuple with the CustomProperty60 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomProperty60

`func (o *GetRoleDetailsResponse) SetCustomProperty60(v string)`

SetCustomProperty60 sets CustomProperty60 field to given value.

### HasCustomProperty60

`func (o *GetRoleDetailsResponse) HasCustomProperty60() bool`

HasCustomProperty60 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


