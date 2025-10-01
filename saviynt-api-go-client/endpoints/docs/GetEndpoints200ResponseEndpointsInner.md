# GetEndpoints200ResponseEndpointsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** | Description for the endpoint. | [optional] 
**StatusForUniqueAccount** | Pointer to **string** |  | [optional] 
**Requestowner** | Pointer to **string** |  | [optional] 
**Requestable** | Pointer to **string** |  | [optional] 
**PrimaryAccountType** | Pointer to **string** | Specify the primary account type for the endpoint. | [optional] 
**AccountTypeNoPasswordChange** | Pointer to **string** | Specify the account type for which password change is not allowed. | [optional] 
**ServiceAccountNameRule** | Pointer to **string** |  | [optional] 
**AccountNameValidatorRegex** | Pointer to **string** |  | [optional] 
**AllowChangePasswordSqlquery** | Pointer to **string** |  | [optional] 
**ParentAccountPattern** | Pointer to **string** |  | [optional] 
**OwnerType** | Pointer to **string** | Specify the owner type for the endpoint. An endpoint can be owned by a User or Usergroup. | [optional] 
**Securitysystem** | Pointer to **string** | Specify the Security system for which you want to create an endpoint. | [optional] 
**Endpointname** | Pointer to **string** | Specify a name for the endpoint. Provide a logical name that will help you easily identify it. | [optional] 
**UpdatedBy** | Pointer to **string** | The username or identifier of the user who last updated this resource. | [optional] 
**Accessquery** | Pointer to **string** | Specify the query to filter the access and display of the endpoint to specific users. If you do not define a query, the endpoint is displayed for all users. | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**DisplayName** | Pointer to **string** | Enter a user-friendly display name for the endpoint that will be displayed in the user interface. Display Name can be different from Endpoint Name. | [optional] 
**UpdateDate** | Pointer to **string** |  | [optional] 
**AllowRemoveAllRoleOnRequest** | Pointer to **string** | Specify true to displays the Remove All Roles option in the Request page that can be used to remove all the roles. | [optional] 
**RoleTypeAsJson** | Pointer to **string** |  | [optional] 
**EntsWithNewAccount** | Pointer to **string** |  | [optional] 
**ConnectionconfigAsJson** | Pointer to **string** |  | [optional] 
**Connectionconfig** | Pointer to **string** | Use this configuration for processing the add access tasks and remove access tasks for AD and LDAP Connectors. | [optional] 
**AccountNameRule** | Pointer to **string** | Specify rule to generate an account name for this endpoint while creating a new account. | [optional] 
**ChangePasswordAccessQuery** | Pointer to **string** | Specify query to restrict the access for changing the account password of the endpoint. | [optional] 
**Disableaccountrequest** | Pointer to **string** |  | [optional] 
**PluginConfigs** | Pointer to **string** |  The Plugin Configuration drives the functionality of the Saviynt SmartAssist (Browserplugin). | [optional] 
**DisableaccountrequestServiceAccount** | Pointer to **string** |  | [optional] 
**Requestableapplication** | Pointer to **string** |  | [optional] 
**CreatedFrom** | Pointer to **string** |  | [optional] 
**CreatedBy** | Pointer to **string** |  | [optional] 
**CreateDate** | Pointer to **string** |  | [optional] 
**ParentEndpoint** | Pointer to **string** |  | [optional] 
**BaseLineConfig** | Pointer to **string** |  | [optional] 
**Requestownertype** | Pointer to **string** |  | [optional] 
**CreateEntTaskforRemoveAcc** | Pointer to **string** | If this is set to true, remove Access tasks will be created for entitlements (account entitlements and their dependent entitlements) when a user requests for removing an account. | [optional] 
**EnableCopyAccess** | Pointer to **string** | Specify true to display the Copy Access from User option in the Request pages. | [optional] 
**AccountTypeNoDeprovision** | Pointer to **string** |  | [optional] 
**EndpointConfig** | Pointer to **string** | Option to copy data in Step 3 of the service account request will be enabled. | [optional] 
**Taskemailtemplates** | Pointer to **string** |  | [optional] 
**Ownerkey** | Pointer to **string** |  | [optional] 
**ServiceAccountAccessQuery** | Pointer to **string** | Specify the query to filter the access and display of the endpoint for specific users while managing service accounts. | [optional] 
**UserAccountCorrelationRule** | Pointer to **string** |  Specify rule to map users in EIC with the accounts during account import. | [optional] 
**StatusConfig** | Pointer to **string** | Enable the State and Status options (Enable, Disable, Lock, Unlock) that would be available to request for a user and service accounts. | [optional] 
**CustomProperty1** | Pointer to **string** | Custom Property 1 | [optional] 
**CustomProperty2** | Pointer to **string** | Custom Property 2 | [optional] 
**CustomProperty3** | Pointer to **string** | Custom Property 3 | [optional] 
**CustomProperty4** | Pointer to **string** | Custom Property 4 | [optional] 
**CustomProperty5** | Pointer to **string** | Custom Property 5 | [optional] 
**CustomProperty6** | Pointer to **string** | Custom Property 6 | [optional] 
**CustomProperty7** | Pointer to **string** | Custom Property 7 | [optional] 
**CustomProperty8** | Pointer to **string** | Custom Property 8 | [optional] 
**CustomProperty9** | Pointer to **string** | Custom Property 9 | [optional] 
**CustomProperty10** | Pointer to **string** | Custom Property 10 | [optional] 
**CustomProperty11** | Pointer to **string** | Custom Property 11 | [optional] 
**CustomProperty12** | Pointer to **string** | Custom Property 12 | [optional] 
**CustomProperty13** | Pointer to **string** | Custom Property 13 | [optional] 
**CustomProperty14** | Pointer to **string** | Custom Property 14 | [optional] 
**CustomProperty15** | Pointer to **string** | Custom Property 15 | [optional] 
**CustomProperty16** | Pointer to **string** | Custom Property 16 | [optional] 
**CustomProperty17** | Pointer to **string** | Custom Property 17 | [optional] 
**CustomProperty18** | Pointer to **string** | Custom Property 18 | [optional] 
**CustomProperty19** | Pointer to **string** | Custom Property 19 | [optional] 
**CustomProperty20** | Pointer to **string** | Custom Property 20 | [optional] 
**CustomProperty21** | Pointer to **string** | Custom Property 21 | [optional] 
**CustomProperty22** | Pointer to **string** | Custom Property 22 | [optional] 
**CustomProperty23** | Pointer to **string** | Custom Property 23 | [optional] 
**CustomProperty24** | Pointer to **string** | Custom Property 24 | [optional] 
**CustomProperty25** | Pointer to **string** | Custom Property 25 | [optional] 
**CustomProperty26** | Pointer to **string** | Custom Property 26 | [optional] 
**CustomProperty27** | Pointer to **string** | Custom Property 27 | [optional] 
**CustomProperty28** | Pointer to **string** | Custom Property 28 | [optional] 
**CustomProperty29** | Pointer to **string** | Custom Property 29 | [optional] 
**CustomProperty30** | Pointer to **string** | Custom Property 30 | [optional] 
**Customproperty31** | Pointer to **string** | Custom Property 31 | [optional] 
**Customproperty32** | Pointer to **string** | Custom Property 32 | [optional] 
**Customproperty33** | Pointer to **string** | Custom Property 33 | [optional] 
**Customproperty34** | Pointer to **string** | Custom Property 34 | [optional] 
**Customproperty35** | Pointer to **string** | Custom Property 35 | [optional] 
**Customproperty36** | Pointer to **string** | Custom Property 36 | [optional] 
**Customproperty37** | Pointer to **string** | Custom Property 37 | [optional] 
**Customproperty38** | Pointer to **string** | Custom Property 38 | [optional] 
**Customproperty39** | Pointer to **string** | Custom Property 39 | [optional] 
**Customproperty40** | Pointer to **string** | Custom Property 40 | [optional] 
**Customproperty41** | Pointer to **string** | Custom Property 41 | [optional] 
**Customproperty42** | Pointer to **string** | Custom Property 42 | [optional] 
**Customproperty43** | Pointer to **string** | Custom Property 43 | [optional] 
**Customproperty44** | Pointer to **string** | Custom Property 44 | [optional] 
**Customproperty45** | Pointer to **string** | Custom Property 45 | [optional] 
**AccountCustomProperty1Label** | Pointer to **string** | Custom Property 1 Label | [optional] 
**AccountCustomProperty2Label** | Pointer to **string** | Custom Property 2 Label | [optional] 
**AccountCustomProperty3Label** | Pointer to **string** | Custom Property 3 Label | [optional] 
**AccountCustomProperty4Label** | Pointer to **string** | Custom Property 4 Label | [optional] 
**AccountCustomProperty5Label** | Pointer to **string** | Custom Property 5 Label | [optional] 
**AccountCustomProperty6Label** | Pointer to **string** | Custom Property 6 Label | [optional] 
**AccountCustomProperty7Label** | Pointer to **string** | Custom Property 7 Label | [optional] 
**AccountCustomProperty8Label** | Pointer to **string** | Custom Property 8 Label | [optional] 
**AccountCustomProperty9Label** | Pointer to **string** | Custom Property 9 Label | [optional] 
**AccountCustomProperty10Label** | Pointer to **string** | Custom Property 10 Label | [optional] 
**AccountCustomProperty11Label** | Pointer to **string** | Custom Property 11 Label | [optional] 
**AccountCustomProperty12Label** | Pointer to **string** | Custom Property 12 Label | [optional] 
**AccountCustomProperty13Label** | Pointer to **string** | Custom Property 13 Label | [optional] 
**AccountCustomProperty14Label** | Pointer to **string** | Custom Property 14 Label | [optional] 
**AccountCustomProperty15Label** | Pointer to **string** | Custom Property 15 Label | [optional] 
**AccountCustomProperty16Label** | Pointer to **string** | Custom Property 16 Label | [optional] 
**AccountCustomProperty17Label** | Pointer to **string** | Custom Property 17 Label | [optional] 
**AccountCustomProperty18Label** | Pointer to **string** | Custom Property 18 Label | [optional] 
**AccountCustomProperty19Label** | Pointer to **string** | Custom Property 19 Label | [optional] 
**AccountCustomProperty20Label** | Pointer to **string** | Custom Property 20 Label | [optional] 
**AccountCustomProperty21Label** | Pointer to **string** | Custom Property 21 Label | [optional] 
**AccountCustomProperty22Label** | Pointer to **string** | Custom Property 22 Label | [optional] 
**AccountCustomProperty23Label** | Pointer to **string** | Custom Property 23 Label | [optional] 
**AccountCustomProperty24Label** | Pointer to **string** | Custom Property 24 Label | [optional] 
**AccountCustomProperty25Label** | Pointer to **string** | Custom Property 25 Label | [optional] 
**AccountCustomProperty26Label** | Pointer to **string** | Custom Property 26 Label | [optional] 
**AccountCustomProperty27Label** | Pointer to **string** | Custom Property 27 Label | [optional] 
**AccountCustomProperty28Label** | Pointer to **string** | Custom Property 28 Label | [optional] 
**AccountCustomProperty29Label** | Pointer to **string** | Custom Property 29 Label | [optional] 
**AccountCustomProperty30Label** | Pointer to **string** | Custom Property 30 Label | [optional] 
**Customproperty31Label** | Pointer to **string** | Custom Property 31 Label | [optional] 
**Customproperty32Label** | Pointer to **string** | Custom Property 32 Label | [optional] 
**Customproperty33Label** | Pointer to **string** | Custom Property 33 Label | [optional] 
**Customproperty34Label** | Pointer to **string** | Custom Property 34 Label | [optional] 
**Customproperty35Label** | Pointer to **string** | Custom Property 35 Label | [optional] 
**Customproperty36Label** | Pointer to **string** | Custom Property 36 Label | [optional] 
**Customproperty37Label** | Pointer to **string** | Custom Property 37 Label | [optional] 
**Customproperty38Label** | Pointer to **string** | Custom Property 38 Label | [optional] 
**Customproperty39Label** | Pointer to **string** | Custom Property 39 Label | [optional] 
**Customproperty40Label** | Pointer to **string** | Custom Property 40 Label | [optional] 
**Customproperty41Label** | Pointer to **string** | Custom Property 41 Label | [optional] 
**Customproperty42Label** | Pointer to **string** | Custom Property 42 Label | [optional] 
**Customproperty43Label** | Pointer to **string** | Custom Property 43 Label | [optional] 
**Customproperty44Label** | Pointer to **string** | Custom Property 44 Label | [optional] 
**Customproperty45Label** | Pointer to **string** | Custom Property 45 Label | [optional] 
**Customproperty46Label** | Pointer to **string** | Custom Property 46 Label | [optional] 
**Customproperty47Label** | Pointer to **string** | Custom Property 47 Label | [optional] 
**Customproperty48Label** | Pointer to **string** | Custom Property 48 Label | [optional] 
**Customproperty49Label** | Pointer to **string** | Custom Property 49 Label | [optional] 
**Customproperty50Label** | Pointer to **string** | Custom Property 50 Label | [optional] 
**Customproperty51Label** | Pointer to **string** | Custom Property 51 Label | [optional] 
**Customproperty52Label** | Pointer to **string** | Custom Property 52 Label | [optional] 
**Customproperty53Label** | Pointer to **string** | Custom Property 53 Label | [optional] 
**Customproperty54Label** | Pointer to **string** | Custom Property 54 Label | [optional] 
**Customproperty55Label** | Pointer to **string** | Custom Property 55 Label | [optional] 
**Customproperty56Label** | Pointer to **string** | Custom Property 56 Label | [optional] 
**Customproperty57Label** | Pointer to **string** | Custom Property 57 Label | [optional] 
**Customproperty58Label** | Pointer to **string** | Custom Property 58 Label | [optional] 
**Customproperty59Label** | Pointer to **string** | Custom Property 59 Label | [optional] 
**Customproperty60Label** | Pointer to **string** | Custom Property 60 Label | [optional] 

## Methods

### NewGetEndpoints200ResponseEndpointsInner

`func NewGetEndpoints200ResponseEndpointsInner() *GetEndpoints200ResponseEndpointsInner`

NewGetEndpoints200ResponseEndpointsInner instantiates a new GetEndpoints200ResponseEndpointsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetEndpoints200ResponseEndpointsInnerWithDefaults

`func NewGetEndpoints200ResponseEndpointsInnerWithDefaults() *GetEndpoints200ResponseEndpointsInner`

NewGetEndpoints200ResponseEndpointsInnerWithDefaults instantiates a new GetEndpoints200ResponseEndpointsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetEndpoints200ResponseEndpointsInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetEndpoints200ResponseEndpointsInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetEndpoints200ResponseEndpointsInner) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *GetEndpoints200ResponseEndpointsInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDescription

`func (o *GetEndpoints200ResponseEndpointsInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GetEndpoints200ResponseEndpointsInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GetEndpoints200ResponseEndpointsInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GetEndpoints200ResponseEndpointsInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetStatusForUniqueAccount

`func (o *GetEndpoints200ResponseEndpointsInner) GetStatusForUniqueAccount() string`

GetStatusForUniqueAccount returns the StatusForUniqueAccount field if non-nil, zero value otherwise.

### GetStatusForUniqueAccountOk

`func (o *GetEndpoints200ResponseEndpointsInner) GetStatusForUniqueAccountOk() (*string, bool)`

GetStatusForUniqueAccountOk returns a tuple with the StatusForUniqueAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusForUniqueAccount

`func (o *GetEndpoints200ResponseEndpointsInner) SetStatusForUniqueAccount(v string)`

SetStatusForUniqueAccount sets StatusForUniqueAccount field to given value.

### HasStatusForUniqueAccount

`func (o *GetEndpoints200ResponseEndpointsInner) HasStatusForUniqueAccount() bool`

HasStatusForUniqueAccount returns a boolean if a field has been set.

### GetRequestowner

`func (o *GetEndpoints200ResponseEndpointsInner) GetRequestowner() string`

GetRequestowner returns the Requestowner field if non-nil, zero value otherwise.

### GetRequestownerOk

`func (o *GetEndpoints200ResponseEndpointsInner) GetRequestownerOk() (*string, bool)`

GetRequestownerOk returns a tuple with the Requestowner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestowner

`func (o *GetEndpoints200ResponseEndpointsInner) SetRequestowner(v string)`

SetRequestowner sets Requestowner field to given value.

### HasRequestowner

`func (o *GetEndpoints200ResponseEndpointsInner) HasRequestowner() bool`

HasRequestowner returns a boolean if a field has been set.

### GetRequestable

`func (o *GetEndpoints200ResponseEndpointsInner) GetRequestable() string`

GetRequestable returns the Requestable field if non-nil, zero value otherwise.

### GetRequestableOk

`func (o *GetEndpoints200ResponseEndpointsInner) GetRequestableOk() (*string, bool)`

GetRequestableOk returns a tuple with the Requestable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestable

`func (o *GetEndpoints200ResponseEndpointsInner) SetRequestable(v string)`

SetRequestable sets Requestable field to given value.

### HasRequestable

`func (o *GetEndpoints200ResponseEndpointsInner) HasRequestable() bool`

HasRequestable returns a boolean if a field has been set.

### GetPrimaryAccountType

`func (o *GetEndpoints200ResponseEndpointsInner) GetPrimaryAccountType() string`

GetPrimaryAccountType returns the PrimaryAccountType field if non-nil, zero value otherwise.

### GetPrimaryAccountTypeOk

`func (o *GetEndpoints200ResponseEndpointsInner) GetPrimaryAccountTypeOk() (*string, bool)`

GetPrimaryAccountTypeOk returns a tuple with the PrimaryAccountType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryAccountType

`func (o *GetEndpoints200ResponseEndpointsInner) SetPrimaryAccountType(v string)`

SetPrimaryAccountType sets PrimaryAccountType field to given value.

### HasPrimaryAccountType

`func (o *GetEndpoints200ResponseEndpointsInner) HasPrimaryAccountType() bool`

HasPrimaryAccountType returns a boolean if a field has been set.

### GetAccountTypeNoPasswordChange

`func (o *GetEndpoints200ResponseEndpointsInner) GetAccountTypeNoPasswordChange() string`

GetAccountTypeNoPasswordChange returns the AccountTypeNoPasswordChange field if non-nil, zero value otherwise.

### GetAccountTypeNoPasswordChangeOk

`func (o *GetEndpoints200ResponseEndpointsInner) GetAccountTypeNoPasswordChangeOk() (*string, bool)`

GetAccountTypeNoPasswordChangeOk returns a tuple with the AccountTypeNoPasswordChange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountTypeNoPasswordChange

`func (o *GetEndpoints200ResponseEndpointsInner) SetAccountTypeNoPasswordChange(v string)`

SetAccountTypeNoPasswordChange sets AccountTypeNoPasswordChange field to given value.

### HasAccountTypeNoPasswordChange

`func (o *GetEndpoints200ResponseEndpointsInner) HasAccountTypeNoPasswordChange() bool`

HasAccountTypeNoPasswordChange returns a boolean if a field has been set.

### GetServiceAccountNameRule

`func (o *GetEndpoints200ResponseEndpointsInner) GetServiceAccountNameRule() string`

GetServiceAccountNameRule returns the ServiceAccountNameRule field if non-nil, zero value otherwise.

### GetServiceAccountNameRuleOk

`func (o *GetEndpoints200ResponseEndpointsInner) GetServiceAccountNameRuleOk() (*string, bool)`

GetServiceAccountNameRuleOk returns a tuple with the ServiceAccountNameRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAccountNameRule

`func (o *GetEndpoints200ResponseEndpointsInner) SetServiceAccountNameRule(v string)`

SetServiceAccountNameRule sets ServiceAccountNameRule field to given value.

### HasServiceAccountNameRule

`func (o *GetEndpoints200ResponseEndpointsInner) HasServiceAccountNameRule() bool`

HasServiceAccountNameRule returns a boolean if a field has been set.

### GetAccountNameValidatorRegex

`func (o *GetEndpoints200ResponseEndpointsInner) GetAccountNameValidatorRegex() string`

GetAccountNameValidatorRegex returns the AccountNameValidatorRegex field if non-nil, zero value otherwise.

### GetAccountNameValidatorRegexOk

`func (o *GetEndpoints200ResponseEndpointsInner) GetAccountNameValidatorRegexOk() (*string, bool)`

GetAccountNameValidatorRegexOk returns a tuple with the AccountNameValidatorRegex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountNameValidatorRegex

`func (o *GetEndpoints200ResponseEndpointsInner) SetAccountNameValidatorRegex(v string)`

SetAccountNameValidatorRegex sets AccountNameValidatorRegex field to given value.

### HasAccountNameValidatorRegex

`func (o *GetEndpoints200ResponseEndpointsInner) HasAccountNameValidatorRegex() bool`

HasAccountNameValidatorRegex returns a boolean if a field has been set.

### GetAllowChangePasswordSqlquery

`func (o *GetEndpoints200ResponseEndpointsInner) GetAllowChangePasswordSqlquery() string`

GetAllowChangePasswordSqlquery returns the AllowChangePasswordSqlquery field if non-nil, zero value otherwise.

### GetAllowChangePasswordSqlqueryOk

`func (o *GetEndpoints200ResponseEndpointsInner) GetAllowChangePasswordSqlqueryOk() (*string, bool)`

GetAllowChangePasswordSqlqueryOk returns a tuple with the AllowChangePasswordSqlquery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowChangePasswordSqlquery

`func (o *GetEndpoints200ResponseEndpointsInner) SetAllowChangePasswordSqlquery(v string)`

SetAllowChangePasswordSqlquery sets AllowChangePasswordSqlquery field to given value.

### HasAllowChangePasswordSqlquery

`func (o *GetEndpoints200ResponseEndpointsInner) HasAllowChangePasswordSqlquery() bool`

HasAllowChangePasswordSqlquery returns a boolean if a field has been set.

### GetParentAccountPattern

`func (o *GetEndpoints200ResponseEndpointsInner) GetParentAccountPattern() string`

GetParentAccountPattern returns the ParentAccountPattern field if non-nil, zero value otherwise.

### GetParentAccountPatternOk

`func (o *GetEndpoints200ResponseEndpointsInner) GetParentAccountPatternOk() (*string, bool)`

GetParentAccountPatternOk returns a tuple with the ParentAccountPattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentAccountPattern

`func (o *GetEndpoints200ResponseEndpointsInner) SetParentAccountPattern(v string)`

SetParentAccountPattern sets ParentAccountPattern field to given value.

### HasParentAccountPattern

`func (o *GetEndpoints200ResponseEndpointsInner) HasParentAccountPattern() bool`

HasParentAccountPattern returns a boolean if a field has been set.

### GetOwnerType

`func (o *GetEndpoints200ResponseEndpointsInner) GetOwnerType() string`

GetOwnerType returns the OwnerType field if non-nil, zero value otherwise.

### GetOwnerTypeOk

`func (o *GetEndpoints200ResponseEndpointsInner) GetOwnerTypeOk() (*string, bool)`

GetOwnerTypeOk returns a tuple with the OwnerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerType

`func (o *GetEndpoints200ResponseEndpointsInner) SetOwnerType(v string)`

SetOwnerType sets OwnerType field to given value.

### HasOwnerType

`func (o *GetEndpoints200ResponseEndpointsInner) HasOwnerType() bool`

HasOwnerType returns a boolean if a field has been set.

### GetSecuritysystem

`func (o *GetEndpoints200ResponseEndpointsInner) GetSecuritysystem() string`

GetSecuritysystem returns the Securitysystem field if non-nil, zero value otherwise.

### GetSecuritysystemOk

`func (o *GetEndpoints200ResponseEndpointsInner) GetSecuritysystemOk() (*string, bool)`

GetSecuritysystemOk returns a tuple with the Securitysystem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecuritysystem

`func (o *GetEndpoints200ResponseEndpointsInner) SetSecuritysystem(v string)`

SetSecuritysystem sets Securitysystem field to given value.

### HasSecuritysystem

`func (o *GetEndpoints200ResponseEndpointsInner) HasSecuritysystem() bool`

HasSecuritysystem returns a boolean if a field has been set.

### GetEndpointname

`func (o *GetEndpoints200ResponseEndpointsInner) GetEndpointname() string`

GetEndpointname returns the Endpointname field if non-nil, zero value otherwise.

### GetEndpointnameOk

`func (o *GetEndpoints200ResponseEndpointsInner) GetEndpointnameOk() (*string, bool)`

GetEndpointnameOk returns a tuple with the Endpointname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpointname

`func (o *GetEndpoints200ResponseEndpointsInner) SetEndpointname(v string)`

SetEndpointname sets Endpointname field to given value.

### HasEndpointname

`func (o *GetEndpoints200ResponseEndpointsInner) HasEndpointname() bool`

HasEndpointname returns a boolean if a field has been set.

### GetUpdatedBy

`func (o *GetEndpoints200ResponseEndpointsInner) GetUpdatedBy() string`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *GetEndpoints200ResponseEndpointsInner) GetUpdatedByOk() (*string, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *GetEndpoints200ResponseEndpointsInner) SetUpdatedBy(v string)`

SetUpdatedBy sets UpdatedBy field to given value.

### HasUpdatedBy

`func (o *GetEndpoints200ResponseEndpointsInner) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.

### GetAccessquery

`func (o *GetEndpoints200ResponseEndpointsInner) GetAccessquery() string`

GetAccessquery returns the Accessquery field if non-nil, zero value otherwise.

### GetAccessqueryOk

`func (o *GetEndpoints200ResponseEndpointsInner) GetAccessqueryOk() (*string, bool)`

GetAccessqueryOk returns a tuple with the Accessquery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessquery

`func (o *GetEndpoints200ResponseEndpointsInner) SetAccessquery(v string)`

SetAccessquery sets Accessquery field to given value.

### HasAccessquery

`func (o *GetEndpoints200ResponseEndpointsInner) HasAccessquery() bool`

HasAccessquery returns a boolean if a field has been set.

### GetStatus

`func (o *GetEndpoints200ResponseEndpointsInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetEndpoints200ResponseEndpointsInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetEndpoints200ResponseEndpointsInner) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetEndpoints200ResponseEndpointsInner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetDisplayName

`func (o *GetEndpoints200ResponseEndpointsInner) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *GetEndpoints200ResponseEndpointsInner) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *GetEndpoints200ResponseEndpointsInner) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *GetEndpoints200ResponseEndpointsInner) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetUpdateDate

`func (o *GetEndpoints200ResponseEndpointsInner) GetUpdateDate() string`

GetUpdateDate returns the UpdateDate field if non-nil, zero value otherwise.

### GetUpdateDateOk

`func (o *GetEndpoints200ResponseEndpointsInner) GetUpdateDateOk() (*string, bool)`

GetUpdateDateOk returns a tuple with the UpdateDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateDate

`func (o *GetEndpoints200ResponseEndpointsInner) SetUpdateDate(v string)`

SetUpdateDate sets UpdateDate field to given value.

### HasUpdateDate

`func (o *GetEndpoints200ResponseEndpointsInner) HasUpdateDate() bool`

HasUpdateDate returns a boolean if a field has been set.

### GetAllowRemoveAllRoleOnRequest

`func (o *GetEndpoints200ResponseEndpointsInner) GetAllowRemoveAllRoleOnRequest() string`

GetAllowRemoveAllRoleOnRequest returns the AllowRemoveAllRoleOnRequest field if non-nil, zero value otherwise.

### GetAllowRemoveAllRoleOnRequestOk

`func (o *GetEndpoints200ResponseEndpointsInner) GetAllowRemoveAllRoleOnRequestOk() (*string, bool)`

GetAllowRemoveAllRoleOnRequestOk returns a tuple with the AllowRemoveAllRoleOnRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowRemoveAllRoleOnRequest

`func (o *GetEndpoints200ResponseEndpointsInner) SetAllowRemoveAllRoleOnRequest(v string)`

SetAllowRemoveAllRoleOnRequest sets AllowRemoveAllRoleOnRequest field to given value.

### HasAllowRemoveAllRoleOnRequest

`func (o *GetEndpoints200ResponseEndpointsInner) HasAllowRemoveAllRoleOnRequest() bool`

HasAllowRemoveAllRoleOnRequest returns a boolean if a field has been set.

### GetRoleTypeAsJson

`func (o *GetEndpoints200ResponseEndpointsInner) GetRoleTypeAsJson() string`

GetRoleTypeAsJson returns the RoleTypeAsJson field if non-nil, zero value otherwise.

### GetRoleTypeAsJsonOk

`func (o *GetEndpoints200ResponseEndpointsInner) GetRoleTypeAsJsonOk() (*string, bool)`

GetRoleTypeAsJsonOk returns a tuple with the RoleTypeAsJson field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleTypeAsJson

`func (o *GetEndpoints200ResponseEndpointsInner) SetRoleTypeAsJson(v string)`

SetRoleTypeAsJson sets RoleTypeAsJson field to given value.

### HasRoleTypeAsJson

`func (o *GetEndpoints200ResponseEndpointsInner) HasRoleTypeAsJson() bool`

HasRoleTypeAsJson returns a boolean if a field has been set.

### GetEntsWithNewAccount

`func (o *GetEndpoints200ResponseEndpointsInner) GetEntsWithNewAccount() string`

GetEntsWithNewAccount returns the EntsWithNewAccount field if non-nil, zero value otherwise.

### GetEntsWithNewAccountOk

`func (o *GetEndpoints200ResponseEndpointsInner) GetEntsWithNewAccountOk() (*string, bool)`

GetEntsWithNewAccountOk returns a tuple with the EntsWithNewAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntsWithNewAccount

`func (o *GetEndpoints200ResponseEndpointsInner) SetEntsWithNewAccount(v string)`

SetEntsWithNewAccount sets EntsWithNewAccount field to given value.

### HasEntsWithNewAccount

`func (o *GetEndpoints200ResponseEndpointsInner) HasEntsWithNewAccount() bool`

HasEntsWithNewAccount returns a boolean if a field has been set.

### GetConnectionconfigAsJson

`func (o *GetEndpoints200ResponseEndpointsInner) GetConnectionconfigAsJson() string`

GetConnectionconfigAsJson returns the ConnectionconfigAsJson field if non-nil, zero value otherwise.

### GetConnectionconfigAsJsonOk

`func (o *GetEndpoints200ResponseEndpointsInner) GetConnectionconfigAsJsonOk() (*string, bool)`

GetConnectionconfigAsJsonOk returns a tuple with the ConnectionconfigAsJson field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionconfigAsJson

`func (o *GetEndpoints200ResponseEndpointsInner) SetConnectionconfigAsJson(v string)`

SetConnectionconfigAsJson sets ConnectionconfigAsJson field to given value.

### HasConnectionconfigAsJson

`func (o *GetEndpoints200ResponseEndpointsInner) HasConnectionconfigAsJson() bool`

HasConnectionconfigAsJson returns a boolean if a field has been set.

### GetConnectionconfig

`func (o *GetEndpoints200ResponseEndpointsInner) GetConnectionconfig() string`

GetConnectionconfig returns the Connectionconfig field if non-nil, zero value otherwise.

### GetConnectionconfigOk

`func (o *GetEndpoints200ResponseEndpointsInner) GetConnectionconfigOk() (*string, bool)`

GetConnectionconfigOk returns a tuple with the Connectionconfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionconfig

`func (o *GetEndpoints200ResponseEndpointsInner) SetConnectionconfig(v string)`

SetConnectionconfig sets Connectionconfig field to given value.

### HasConnectionconfig

`func (o *GetEndpoints200ResponseEndpointsInner) HasConnectionconfig() bool`

HasConnectionconfig returns a boolean if a field has been set.

### GetAccountNameRule

`func (o *GetEndpoints200ResponseEndpointsInner) GetAccountNameRule() string`

GetAccountNameRule returns the AccountNameRule field if non-nil, zero value otherwise.

### GetAccountNameRuleOk

`func (o *GetEndpoints200ResponseEndpointsInner) GetAccountNameRuleOk() (*string, bool)`

GetAccountNameRuleOk returns a tuple with the AccountNameRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountNameRule

`func (o *GetEndpoints200ResponseEndpointsInner) SetAccountNameRule(v string)`

SetAccountNameRule sets AccountNameRule field to given value.

### HasAccountNameRule

`func (o *GetEndpoints200ResponseEndpointsInner) HasAccountNameRule() bool`

HasAccountNameRule returns a boolean if a field has been set.

### GetChangePasswordAccessQuery

`func (o *GetEndpoints200ResponseEndpointsInner) GetChangePasswordAccessQuery() string`

GetChangePasswordAccessQuery returns the ChangePasswordAccessQuery field if non-nil, zero value otherwise.

### GetChangePasswordAccessQueryOk

`func (o *GetEndpoints200ResponseEndpointsInner) GetChangePasswordAccessQueryOk() (*string, bool)`

GetChangePasswordAccessQueryOk returns a tuple with the ChangePasswordAccessQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangePasswordAccessQuery

`func (o *GetEndpoints200ResponseEndpointsInner) SetChangePasswordAccessQuery(v string)`

SetChangePasswordAccessQuery sets ChangePasswordAccessQuery field to given value.

### HasChangePasswordAccessQuery

`func (o *GetEndpoints200ResponseEndpointsInner) HasChangePasswordAccessQuery() bool`

HasChangePasswordAccessQuery returns a boolean if a field has been set.

### GetDisableaccountrequest

`func (o *GetEndpoints200ResponseEndpointsInner) GetDisableaccountrequest() string`

GetDisableaccountrequest returns the Disableaccountrequest field if non-nil, zero value otherwise.

### GetDisableaccountrequestOk

`func (o *GetEndpoints200ResponseEndpointsInner) GetDisableaccountrequestOk() (*string, bool)`

GetDisableaccountrequestOk returns a tuple with the Disableaccountrequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableaccountrequest

`func (o *GetEndpoints200ResponseEndpointsInner) SetDisableaccountrequest(v string)`

SetDisableaccountrequest sets Disableaccountrequest field to given value.

### HasDisableaccountrequest

`func (o *GetEndpoints200ResponseEndpointsInner) HasDisableaccountrequest() bool`

HasDisableaccountrequest returns a boolean if a field has been set.

### GetPluginConfigs

`func (o *GetEndpoints200ResponseEndpointsInner) GetPluginConfigs() string`

GetPluginConfigs returns the PluginConfigs field if non-nil, zero value otherwise.

### GetPluginConfigsOk

`func (o *GetEndpoints200ResponseEndpointsInner) GetPluginConfigsOk() (*string, bool)`

GetPluginConfigsOk returns a tuple with the PluginConfigs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPluginConfigs

`func (o *GetEndpoints200ResponseEndpointsInner) SetPluginConfigs(v string)`

SetPluginConfigs sets PluginConfigs field to given value.

### HasPluginConfigs

`func (o *GetEndpoints200ResponseEndpointsInner) HasPluginConfigs() bool`

HasPluginConfigs returns a boolean if a field has been set.

### GetDisableaccountrequestServiceAccount

`func (o *GetEndpoints200ResponseEndpointsInner) GetDisableaccountrequestServiceAccount() string`

GetDisableaccountrequestServiceAccount returns the DisableaccountrequestServiceAccount field if non-nil, zero value otherwise.

### GetDisableaccountrequestServiceAccountOk

`func (o *GetEndpoints200ResponseEndpointsInner) GetDisableaccountrequestServiceAccountOk() (*string, bool)`

GetDisableaccountrequestServiceAccountOk returns a tuple with the DisableaccountrequestServiceAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableaccountrequestServiceAccount

`func (o *GetEndpoints200ResponseEndpointsInner) SetDisableaccountrequestServiceAccount(v string)`

SetDisableaccountrequestServiceAccount sets DisableaccountrequestServiceAccount field to given value.

### HasDisableaccountrequestServiceAccount

`func (o *GetEndpoints200ResponseEndpointsInner) HasDisableaccountrequestServiceAccount() bool`

HasDisableaccountrequestServiceAccount returns a boolean if a field has been set.

### GetRequestableapplication

`func (o *GetEndpoints200ResponseEndpointsInner) GetRequestableapplication() string`

GetRequestableapplication returns the Requestableapplication field if non-nil, zero value otherwise.

### GetRequestableapplicationOk

`func (o *GetEndpoints200ResponseEndpointsInner) GetRequestableapplicationOk() (*string, bool)`

GetRequestableapplicationOk returns a tuple with the Requestableapplication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestableapplication

`func (o *GetEndpoints200ResponseEndpointsInner) SetRequestableapplication(v string)`

SetRequestableapplication sets Requestableapplication field to given value.

### HasRequestableapplication

`func (o *GetEndpoints200ResponseEndpointsInner) HasRequestableapplication() bool`

HasRequestableapplication returns a boolean if a field has been set.

### GetCreatedFrom

`func (o *GetEndpoints200ResponseEndpointsInner) GetCreatedFrom() string`

GetCreatedFrom returns the CreatedFrom field if non-nil, zero value otherwise.

### GetCreatedFromOk

`func (o *GetEndpoints200ResponseEndpointsInner) GetCreatedFromOk() (*string, bool)`

GetCreatedFromOk returns a tuple with the CreatedFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedFrom

`func (o *GetEndpoints200ResponseEndpointsInner) SetCreatedFrom(v string)`

SetCreatedFrom sets CreatedFrom field to given value.

### HasCreatedFrom

`func (o *GetEndpoints200ResponseEndpointsInner) HasCreatedFrom() bool`

HasCreatedFrom returns a boolean if a field has been set.

### GetCreatedBy

`func (o *GetEndpoints200ResponseEndpointsInner) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *GetEndpoints200ResponseEndpointsInner) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *GetEndpoints200ResponseEndpointsInner) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *GetEndpoints200ResponseEndpointsInner) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetCreateDate

`func (o *GetEndpoints200ResponseEndpointsInner) GetCreateDate() string`

GetCreateDate returns the CreateDate field if non-nil, zero value otherwise.

### GetCreateDateOk

`func (o *GetEndpoints200ResponseEndpointsInner) GetCreateDateOk() (*string, bool)`

GetCreateDateOk returns a tuple with the CreateDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateDate

`func (o *GetEndpoints200ResponseEndpointsInner) SetCreateDate(v string)`

SetCreateDate sets CreateDate field to given value.

### HasCreateDate

`func (o *GetEndpoints200ResponseEndpointsInner) HasCreateDate() bool`

HasCreateDate returns a boolean if a field has been set.

### GetParentEndpoint

`func (o *GetEndpoints200ResponseEndpointsInner) GetParentEndpoint() string`

GetParentEndpoint returns the ParentEndpoint field if non-nil, zero value otherwise.

### GetParentEndpointOk

`func (o *GetEndpoints200ResponseEndpointsInner) GetParentEndpointOk() (*string, bool)`

GetParentEndpointOk returns a tuple with the ParentEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentEndpoint

`func (o *GetEndpoints200ResponseEndpointsInner) SetParentEndpoint(v string)`

SetParentEndpoint sets ParentEndpoint field to given value.

### HasParentEndpoint

`func (o *GetEndpoints200ResponseEndpointsInner) HasParentEndpoint() bool`

HasParentEndpoint returns a boolean if a field has been set.

### GetBaseLineConfig

`func (o *GetEndpoints200ResponseEndpointsInner) GetBaseLineConfig() string`

GetBaseLineConfig returns the BaseLineConfig field if non-nil, zero value otherwise.

### GetBaseLineConfigOk

`func (o *GetEndpoints200ResponseEndpointsInner) GetBaseLineConfigOk() (*string, bool)`

GetBaseLineConfigOk returns a tuple with the BaseLineConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseLineConfig

`func (o *GetEndpoints200ResponseEndpointsInner) SetBaseLineConfig(v string)`

SetBaseLineConfig sets BaseLineConfig field to given value.

### HasBaseLineConfig

`func (o *GetEndpoints200ResponseEndpointsInner) HasBaseLineConfig() bool`

HasBaseLineConfig returns a boolean if a field has been set.

### GetRequestownertype

`func (o *GetEndpoints200ResponseEndpointsInner) GetRequestownertype() string`

GetRequestownertype returns the Requestownertype field if non-nil, zero value otherwise.

### GetRequestownertypeOk

`func (o *GetEndpoints200ResponseEndpointsInner) GetRequestownertypeOk() (*string, bool)`

GetRequestownertypeOk returns a tuple with the Requestownertype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestownertype

`func (o *GetEndpoints200ResponseEndpointsInner) SetRequestownertype(v string)`

SetRequestownertype sets Requestownertype field to given value.

### HasRequestownertype

`func (o *GetEndpoints200ResponseEndpointsInner) HasRequestownertype() bool`

HasRequestownertype returns a boolean if a field has been set.

### GetCreateEntTaskforRemoveAcc

`func (o *GetEndpoints200ResponseEndpointsInner) GetCreateEntTaskforRemoveAcc() string`

GetCreateEntTaskforRemoveAcc returns the CreateEntTaskforRemoveAcc field if non-nil, zero value otherwise.

### GetCreateEntTaskforRemoveAccOk

`func (o *GetEndpoints200ResponseEndpointsInner) GetCreateEntTaskforRemoveAccOk() (*string, bool)`

GetCreateEntTaskforRemoveAccOk returns a tuple with the CreateEntTaskforRemoveAcc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateEntTaskforRemoveAcc

`func (o *GetEndpoints200ResponseEndpointsInner) SetCreateEntTaskforRemoveAcc(v string)`

SetCreateEntTaskforRemoveAcc sets CreateEntTaskforRemoveAcc field to given value.

### HasCreateEntTaskforRemoveAcc

`func (o *GetEndpoints200ResponseEndpointsInner) HasCreateEntTaskforRemoveAcc() bool`

HasCreateEntTaskforRemoveAcc returns a boolean if a field has been set.

### GetEnableCopyAccess

`func (o *GetEndpoints200ResponseEndpointsInner) GetEnableCopyAccess() string`

GetEnableCopyAccess returns the EnableCopyAccess field if non-nil, zero value otherwise.

### GetEnableCopyAccessOk

`func (o *GetEndpoints200ResponseEndpointsInner) GetEnableCopyAccessOk() (*string, bool)`

GetEnableCopyAccessOk returns a tuple with the EnableCopyAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableCopyAccess

`func (o *GetEndpoints200ResponseEndpointsInner) SetEnableCopyAccess(v string)`

SetEnableCopyAccess sets EnableCopyAccess field to given value.

### HasEnableCopyAccess

`func (o *GetEndpoints200ResponseEndpointsInner) HasEnableCopyAccess() bool`

HasEnableCopyAccess returns a boolean if a field has been set.

### GetAccountTypeNoDeprovision

`func (o *GetEndpoints200ResponseEndpointsInner) GetAccountTypeNoDeprovision() string`

GetAccountTypeNoDeprovision returns the AccountTypeNoDeprovision field if non-nil, zero value otherwise.

### GetAccountTypeNoDeprovisionOk

`func (o *GetEndpoints200ResponseEndpointsInner) GetAccountTypeNoDeprovisionOk() (*string, bool)`

GetAccountTypeNoDeprovisionOk returns a tuple with the AccountTypeNoDeprovision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountTypeNoDeprovision

`func (o *GetEndpoints200ResponseEndpointsInner) SetAccountTypeNoDeprovision(v string)`

SetAccountTypeNoDeprovision sets AccountTypeNoDeprovision field to given value.

### HasAccountTypeNoDeprovision

`func (o *GetEndpoints200ResponseEndpointsInner) HasAccountTypeNoDeprovision() bool`

HasAccountTypeNoDeprovision returns a boolean if a field has been set.

### GetEndpointConfig

`func (o *GetEndpoints200ResponseEndpointsInner) GetEndpointConfig() string`

GetEndpointConfig returns the EndpointConfig field if non-nil, zero value otherwise.

### GetEndpointConfigOk

`func (o *GetEndpoints200ResponseEndpointsInner) GetEndpointConfigOk() (*string, bool)`

GetEndpointConfigOk returns a tuple with the EndpointConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpointConfig

`func (o *GetEndpoints200ResponseEndpointsInner) SetEndpointConfig(v string)`

SetEndpointConfig sets EndpointConfig field to given value.

### HasEndpointConfig

`func (o *GetEndpoints200ResponseEndpointsInner) HasEndpointConfig() bool`

HasEndpointConfig returns a boolean if a field has been set.

### GetTaskemailtemplates

`func (o *GetEndpoints200ResponseEndpointsInner) GetTaskemailtemplates() string`

GetTaskemailtemplates returns the Taskemailtemplates field if non-nil, zero value otherwise.

### GetTaskemailtemplatesOk

`func (o *GetEndpoints200ResponseEndpointsInner) GetTaskemailtemplatesOk() (*string, bool)`

GetTaskemailtemplatesOk returns a tuple with the Taskemailtemplates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskemailtemplates

`func (o *GetEndpoints200ResponseEndpointsInner) SetTaskemailtemplates(v string)`

SetTaskemailtemplates sets Taskemailtemplates field to given value.

### HasTaskemailtemplates

`func (o *GetEndpoints200ResponseEndpointsInner) HasTaskemailtemplates() bool`

HasTaskemailtemplates returns a boolean if a field has been set.

### GetOwnerkey

`func (o *GetEndpoints200ResponseEndpointsInner) GetOwnerkey() string`

GetOwnerkey returns the Ownerkey field if non-nil, zero value otherwise.

### GetOwnerkeyOk

`func (o *GetEndpoints200ResponseEndpointsInner) GetOwnerkeyOk() (*string, bool)`

GetOwnerkeyOk returns a tuple with the Ownerkey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerkey

`func (o *GetEndpoints200ResponseEndpointsInner) SetOwnerkey(v string)`

SetOwnerkey sets Ownerkey field to given value.

### HasOwnerkey

`func (o *GetEndpoints200ResponseEndpointsInner) HasOwnerkey() bool`

HasOwnerkey returns a boolean if a field has been set.

### GetServiceAccountAccessQuery

`func (o *GetEndpoints200ResponseEndpointsInner) GetServiceAccountAccessQuery() string`

GetServiceAccountAccessQuery returns the ServiceAccountAccessQuery field if non-nil, zero value otherwise.

### GetServiceAccountAccessQueryOk

`func (o *GetEndpoints200ResponseEndpointsInner) GetServiceAccountAccessQueryOk() (*string, bool)`

GetServiceAccountAccessQueryOk returns a tuple with the ServiceAccountAccessQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAccountAccessQuery

`func (o *GetEndpoints200ResponseEndpointsInner) SetServiceAccountAccessQuery(v string)`

SetServiceAccountAccessQuery sets ServiceAccountAccessQuery field to given value.

### HasServiceAccountAccessQuery

`func (o *GetEndpoints200ResponseEndpointsInner) HasServiceAccountAccessQuery() bool`

HasServiceAccountAccessQuery returns a boolean if a field has been set.

### GetUserAccountCorrelationRule

`func (o *GetEndpoints200ResponseEndpointsInner) GetUserAccountCorrelationRule() string`

GetUserAccountCorrelationRule returns the UserAccountCorrelationRule field if non-nil, zero value otherwise.

### GetUserAccountCorrelationRuleOk

`func (o *GetEndpoints200ResponseEndpointsInner) GetUserAccountCorrelationRuleOk() (*string, bool)`

GetUserAccountCorrelationRuleOk returns a tuple with the UserAccountCorrelationRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAccountCorrelationRule

`func (o *GetEndpoints200ResponseEndpointsInner) SetUserAccountCorrelationRule(v string)`

SetUserAccountCorrelationRule sets UserAccountCorrelationRule field to given value.

### HasUserAccountCorrelationRule

`func (o *GetEndpoints200ResponseEndpointsInner) HasUserAccountCorrelationRule() bool`

HasUserAccountCorrelationRule returns a boolean if a field has been set.

### GetStatusConfig

`func (o *GetEndpoints200ResponseEndpointsInner) GetStatusConfig() string`

GetStatusConfig returns the StatusConfig field if non-nil, zero value otherwise.

### GetStatusConfigOk

`func (o *GetEndpoints200ResponseEndpointsInner) GetStatusConfigOk() (*string, bool)`

GetStatusConfigOk returns a tuple with the StatusConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusConfig

`func (o *GetEndpoints200ResponseEndpointsInner) SetStatusConfig(v string)`

SetStatusConfig sets StatusConfig field to given value.

### HasStatusConfig

`func (o *GetEndpoints200ResponseEndpointsInner) HasStatusConfig() bool`

HasStatusConfig returns a boolean if a field has been set.

### GetCustomProperty1

`func (o *GetEndpoints200ResponseEndpointsInner) GetCustomProperty1() string`

GetCustomProperty1 returns the CustomProperty1 field if non-nil, zero value otherwise.

### GetCustomProperty1Ok

`func (o *GetEndpoints200ResponseEndpointsInner) GetCustomProperty1Ok() (*string, bool)`

GetCustomProperty1Ok returns a tuple with the CustomProperty1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomProperty1

`func (o *GetEndpoints200ResponseEndpointsInner) SetCustomProperty1(v string)`

SetCustomProperty1 sets CustomProperty1 field to given value.

### HasCustomProperty1

`func (o *GetEndpoints200ResponseEndpointsInner) HasCustomProperty1() bool`

HasCustomProperty1 returns a boolean if a field has been set.

### GetCustomProperty2

`func (o *GetEndpoints200ResponseEndpointsInner) GetCustomProperty2() string`

GetCustomProperty2 returns the CustomProperty2 field if non-nil, zero value otherwise.

### GetCustomProperty2Ok

`func (o *GetEndpoints200ResponseEndpointsInner) GetCustomProperty2Ok() (*string, bool)`

GetCustomProperty2Ok returns a tuple with the CustomProperty2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomProperty2

`func (o *GetEndpoints200ResponseEndpointsInner) SetCustomProperty2(v string)`

SetCustomProperty2 sets CustomProperty2 field to given value.

### HasCustomProperty2

`func (o *GetEndpoints200ResponseEndpointsInner) HasCustomProperty2() bool`

HasCustomProperty2 returns a boolean if a field has been set.

### GetCustomProperty3

`func (o *GetEndpoints200ResponseEndpointsInner) GetCustomProperty3() string`

GetCustomProperty3 returns the CustomProperty3 field if non-nil, zero value otherwise.

### GetCustomProperty3Ok

`func (o *GetEndpoints200ResponseEndpointsInner) GetCustomProperty3Ok() (*string, bool)`

GetCustomProperty3Ok returns a tuple with the CustomProperty3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomProperty3

`func (o *GetEndpoints200ResponseEndpointsInner) SetCustomProperty3(v string)`

SetCustomProperty3 sets CustomProperty3 field to given value.

### HasCustomProperty3

`func (o *GetEndpoints200ResponseEndpointsInner) HasCustomProperty3() bool`

HasCustomProperty3 returns a boolean if a field has been set.

### GetCustomProperty4

`func (o *GetEndpoints200ResponseEndpointsInner) GetCustomProperty4() string`

GetCustomProperty4 returns the CustomProperty4 field if non-nil, zero value otherwise.

### GetCustomProperty4Ok

`func (o *GetEndpoints200ResponseEndpointsInner) GetCustomProperty4Ok() (*string, bool)`

GetCustomProperty4Ok returns a tuple with the CustomProperty4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomProperty4

`func (o *GetEndpoints200ResponseEndpointsInner) SetCustomProperty4(v string)`

SetCustomProperty4 sets CustomProperty4 field to given value.

### HasCustomProperty4

`func (o *GetEndpoints200ResponseEndpointsInner) HasCustomProperty4() bool`

HasCustomProperty4 returns a boolean if a field has been set.

### GetCustomProperty5

`func (o *GetEndpoints200ResponseEndpointsInner) GetCustomProperty5() string`

GetCustomProperty5 returns the CustomProperty5 field if non-nil, zero value otherwise.

### GetCustomProperty5Ok

`func (o *GetEndpoints200ResponseEndpointsInner) GetCustomProperty5Ok() (*string, bool)`

GetCustomProperty5Ok returns a tuple with the CustomProperty5 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomProperty5

`func (o *GetEndpoints200ResponseEndpointsInner) SetCustomProperty5(v string)`

SetCustomProperty5 sets CustomProperty5 field to given value.

### HasCustomProperty5

`func (o *GetEndpoints200ResponseEndpointsInner) HasCustomProperty5() bool`

HasCustomProperty5 returns a boolean if a field has been set.

### GetCustomProperty6

`func (o *GetEndpoints200ResponseEndpointsInner) GetCustomProperty6() string`

GetCustomProperty6 returns the CustomProperty6 field if non-nil, zero value otherwise.

### GetCustomProperty6Ok

`func (o *GetEndpoints200ResponseEndpointsInner) GetCustomProperty6Ok() (*string, bool)`

GetCustomProperty6Ok returns a tuple with the CustomProperty6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomProperty6

`func (o *GetEndpoints200ResponseEndpointsInner) SetCustomProperty6(v string)`

SetCustomProperty6 sets CustomProperty6 field to given value.

### HasCustomProperty6

`func (o *GetEndpoints200ResponseEndpointsInner) HasCustomProperty6() bool`

HasCustomProperty6 returns a boolean if a field has been set.

### GetCustomProperty7

`func (o *GetEndpoints200ResponseEndpointsInner) GetCustomProperty7() string`

GetCustomProperty7 returns the CustomProperty7 field if non-nil, zero value otherwise.

### GetCustomProperty7Ok

`func (o *GetEndpoints200ResponseEndpointsInner) GetCustomProperty7Ok() (*string, bool)`

GetCustomProperty7Ok returns a tuple with the CustomProperty7 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomProperty7

`func (o *GetEndpoints200ResponseEndpointsInner) SetCustomProperty7(v string)`

SetCustomProperty7 sets CustomProperty7 field to given value.

### HasCustomProperty7

`func (o *GetEndpoints200ResponseEndpointsInner) HasCustomProperty7() bool`

HasCustomProperty7 returns a boolean if a field has been set.

### GetCustomProperty8

`func (o *GetEndpoints200ResponseEndpointsInner) GetCustomProperty8() string`

GetCustomProperty8 returns the CustomProperty8 field if non-nil, zero value otherwise.

### GetCustomProperty8Ok

`func (o *GetEndpoints200ResponseEndpointsInner) GetCustomProperty8Ok() (*string, bool)`

GetCustomProperty8Ok returns a tuple with the CustomProperty8 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomProperty8

`func (o *GetEndpoints200ResponseEndpointsInner) SetCustomProperty8(v string)`

SetCustomProperty8 sets CustomProperty8 field to given value.

### HasCustomProperty8

`func (o *GetEndpoints200ResponseEndpointsInner) HasCustomProperty8() bool`

HasCustomProperty8 returns a boolean if a field has been set.

### GetCustomProperty9

`func (o *GetEndpoints200ResponseEndpointsInner) GetCustomProperty9() string`

GetCustomProperty9 returns the CustomProperty9 field if non-nil, zero value otherwise.

### GetCustomProperty9Ok

`func (o *GetEndpoints200ResponseEndpointsInner) GetCustomProperty9Ok() (*string, bool)`

GetCustomProperty9Ok returns a tuple with the CustomProperty9 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomProperty9

`func (o *GetEndpoints200ResponseEndpointsInner) SetCustomProperty9(v string)`

SetCustomProperty9 sets CustomProperty9 field to given value.

### HasCustomProperty9

`func (o *GetEndpoints200ResponseEndpointsInner) HasCustomProperty9() bool`

HasCustomProperty9 returns a boolean if a field has been set.

### GetCustomProperty10

`func (o *GetEndpoints200ResponseEndpointsInner) GetCustomProperty10() string`

GetCustomProperty10 returns the CustomProperty10 field if non-nil, zero value otherwise.

### GetCustomProperty10Ok

`func (o *GetEndpoints200ResponseEndpointsInner) GetCustomProperty10Ok() (*string, bool)`

GetCustomProperty10Ok returns a tuple with the CustomProperty10 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomProperty10

`func (o *GetEndpoints200ResponseEndpointsInner) SetCustomProperty10(v string)`

SetCustomProperty10 sets CustomProperty10 field to given value.

### HasCustomProperty10

`func (o *GetEndpoints200ResponseEndpointsInner) HasCustomProperty10() bool`

HasCustomProperty10 returns a boolean if a field has been set.

### GetCustomProperty11

`func (o *GetEndpoints200ResponseEndpointsInner) GetCustomProperty11() string`

GetCustomProperty11 returns the CustomProperty11 field if non-nil, zero value otherwise.

### GetCustomProperty11Ok

`func (o *GetEndpoints200ResponseEndpointsInner) GetCustomProperty11Ok() (*string, bool)`

GetCustomProperty11Ok returns a tuple with the CustomProperty11 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomProperty11

`func (o *GetEndpoints200ResponseEndpointsInner) SetCustomProperty11(v string)`

SetCustomProperty11 sets CustomProperty11 field to given value.

### HasCustomProperty11

`func (o *GetEndpoints200ResponseEndpointsInner) HasCustomProperty11() bool`

HasCustomProperty11 returns a boolean if a field has been set.

### GetCustomProperty12

`func (o *GetEndpoints200ResponseEndpointsInner) GetCustomProperty12() string`

GetCustomProperty12 returns the CustomProperty12 field if non-nil, zero value otherwise.

### GetCustomProperty12Ok

`func (o *GetEndpoints200ResponseEndpointsInner) GetCustomProperty12Ok() (*string, bool)`

GetCustomProperty12Ok returns a tuple with the CustomProperty12 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomProperty12

`func (o *GetEndpoints200ResponseEndpointsInner) SetCustomProperty12(v string)`

SetCustomProperty12 sets CustomProperty12 field to given value.

### HasCustomProperty12

`func (o *GetEndpoints200ResponseEndpointsInner) HasCustomProperty12() bool`

HasCustomProperty12 returns a boolean if a field has been set.

### GetCustomProperty13

`func (o *GetEndpoints200ResponseEndpointsInner) GetCustomProperty13() string`

GetCustomProperty13 returns the CustomProperty13 field if non-nil, zero value otherwise.

### GetCustomProperty13Ok

`func (o *GetEndpoints200ResponseEndpointsInner) GetCustomProperty13Ok() (*string, bool)`

GetCustomProperty13Ok returns a tuple with the CustomProperty13 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomProperty13

`func (o *GetEndpoints200ResponseEndpointsInner) SetCustomProperty13(v string)`

SetCustomProperty13 sets CustomProperty13 field to given value.

### HasCustomProperty13

`func (o *GetEndpoints200ResponseEndpointsInner) HasCustomProperty13() bool`

HasCustomProperty13 returns a boolean if a field has been set.

### GetCustomProperty14

`func (o *GetEndpoints200ResponseEndpointsInner) GetCustomProperty14() string`

GetCustomProperty14 returns the CustomProperty14 field if non-nil, zero value otherwise.

### GetCustomProperty14Ok

`func (o *GetEndpoints200ResponseEndpointsInner) GetCustomProperty14Ok() (*string, bool)`

GetCustomProperty14Ok returns a tuple with the CustomProperty14 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomProperty14

`func (o *GetEndpoints200ResponseEndpointsInner) SetCustomProperty14(v string)`

SetCustomProperty14 sets CustomProperty14 field to given value.

### HasCustomProperty14

`func (o *GetEndpoints200ResponseEndpointsInner) HasCustomProperty14() bool`

HasCustomProperty14 returns a boolean if a field has been set.

### GetCustomProperty15

`func (o *GetEndpoints200ResponseEndpointsInner) GetCustomProperty15() string`

GetCustomProperty15 returns the CustomProperty15 field if non-nil, zero value otherwise.

### GetCustomProperty15Ok

`func (o *GetEndpoints200ResponseEndpointsInner) GetCustomProperty15Ok() (*string, bool)`

GetCustomProperty15Ok returns a tuple with the CustomProperty15 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomProperty15

`func (o *GetEndpoints200ResponseEndpointsInner) SetCustomProperty15(v string)`

SetCustomProperty15 sets CustomProperty15 field to given value.

### HasCustomProperty15

`func (o *GetEndpoints200ResponseEndpointsInner) HasCustomProperty15() bool`

HasCustomProperty15 returns a boolean if a field has been set.

### GetCustomProperty16

`func (o *GetEndpoints200ResponseEndpointsInner) GetCustomProperty16() string`

GetCustomProperty16 returns the CustomProperty16 field if non-nil, zero value otherwise.

### GetCustomProperty16Ok

`func (o *GetEndpoints200ResponseEndpointsInner) GetCustomProperty16Ok() (*string, bool)`

GetCustomProperty16Ok returns a tuple with the CustomProperty16 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomProperty16

`func (o *GetEndpoints200ResponseEndpointsInner) SetCustomProperty16(v string)`

SetCustomProperty16 sets CustomProperty16 field to given value.

### HasCustomProperty16

`func (o *GetEndpoints200ResponseEndpointsInner) HasCustomProperty16() bool`

HasCustomProperty16 returns a boolean if a field has been set.

### GetCustomProperty17

`func (o *GetEndpoints200ResponseEndpointsInner) GetCustomProperty17() string`

GetCustomProperty17 returns the CustomProperty17 field if non-nil, zero value otherwise.

### GetCustomProperty17Ok

`func (o *GetEndpoints200ResponseEndpointsInner) GetCustomProperty17Ok() (*string, bool)`

GetCustomProperty17Ok returns a tuple with the CustomProperty17 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomProperty17

`func (o *GetEndpoints200ResponseEndpointsInner) SetCustomProperty17(v string)`

SetCustomProperty17 sets CustomProperty17 field to given value.

### HasCustomProperty17

`func (o *GetEndpoints200ResponseEndpointsInner) HasCustomProperty17() bool`

HasCustomProperty17 returns a boolean if a field has been set.

### GetCustomProperty18

`func (o *GetEndpoints200ResponseEndpointsInner) GetCustomProperty18() string`

GetCustomProperty18 returns the CustomProperty18 field if non-nil, zero value otherwise.

### GetCustomProperty18Ok

`func (o *GetEndpoints200ResponseEndpointsInner) GetCustomProperty18Ok() (*string, bool)`

GetCustomProperty18Ok returns a tuple with the CustomProperty18 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomProperty18

`func (o *GetEndpoints200ResponseEndpointsInner) SetCustomProperty18(v string)`

SetCustomProperty18 sets CustomProperty18 field to given value.

### HasCustomProperty18

`func (o *GetEndpoints200ResponseEndpointsInner) HasCustomProperty18() bool`

HasCustomProperty18 returns a boolean if a field has been set.

### GetCustomProperty19

`func (o *GetEndpoints200ResponseEndpointsInner) GetCustomProperty19() string`

GetCustomProperty19 returns the CustomProperty19 field if non-nil, zero value otherwise.

### GetCustomProperty19Ok

`func (o *GetEndpoints200ResponseEndpointsInner) GetCustomProperty19Ok() (*string, bool)`

GetCustomProperty19Ok returns a tuple with the CustomProperty19 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomProperty19

`func (o *GetEndpoints200ResponseEndpointsInner) SetCustomProperty19(v string)`

SetCustomProperty19 sets CustomProperty19 field to given value.

### HasCustomProperty19

`func (o *GetEndpoints200ResponseEndpointsInner) HasCustomProperty19() bool`

HasCustomProperty19 returns a boolean if a field has been set.

### GetCustomProperty20

`func (o *GetEndpoints200ResponseEndpointsInner) GetCustomProperty20() string`

GetCustomProperty20 returns the CustomProperty20 field if non-nil, zero value otherwise.

### GetCustomProperty20Ok

`func (o *GetEndpoints200ResponseEndpointsInner) GetCustomProperty20Ok() (*string, bool)`

GetCustomProperty20Ok returns a tuple with the CustomProperty20 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomProperty20

`func (o *GetEndpoints200ResponseEndpointsInner) SetCustomProperty20(v string)`

SetCustomProperty20 sets CustomProperty20 field to given value.

### HasCustomProperty20

`func (o *GetEndpoints200ResponseEndpointsInner) HasCustomProperty20() bool`

HasCustomProperty20 returns a boolean if a field has been set.

### GetCustomProperty21

`func (o *GetEndpoints200ResponseEndpointsInner) GetCustomProperty21() string`

GetCustomProperty21 returns the CustomProperty21 field if non-nil, zero value otherwise.

### GetCustomProperty21Ok

`func (o *GetEndpoints200ResponseEndpointsInner) GetCustomProperty21Ok() (*string, bool)`

GetCustomProperty21Ok returns a tuple with the CustomProperty21 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomProperty21

`func (o *GetEndpoints200ResponseEndpointsInner) SetCustomProperty21(v string)`

SetCustomProperty21 sets CustomProperty21 field to given value.

### HasCustomProperty21

`func (o *GetEndpoints200ResponseEndpointsInner) HasCustomProperty21() bool`

HasCustomProperty21 returns a boolean if a field has been set.

### GetCustomProperty22

`func (o *GetEndpoints200ResponseEndpointsInner) GetCustomProperty22() string`

GetCustomProperty22 returns the CustomProperty22 field if non-nil, zero value otherwise.

### GetCustomProperty22Ok

`func (o *GetEndpoints200ResponseEndpointsInner) GetCustomProperty22Ok() (*string, bool)`

GetCustomProperty22Ok returns a tuple with the CustomProperty22 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomProperty22

`func (o *GetEndpoints200ResponseEndpointsInner) SetCustomProperty22(v string)`

SetCustomProperty22 sets CustomProperty22 field to given value.

### HasCustomProperty22

`func (o *GetEndpoints200ResponseEndpointsInner) HasCustomProperty22() bool`

HasCustomProperty22 returns a boolean if a field has been set.

### GetCustomProperty23

`func (o *GetEndpoints200ResponseEndpointsInner) GetCustomProperty23() string`

GetCustomProperty23 returns the CustomProperty23 field if non-nil, zero value otherwise.

### GetCustomProperty23Ok

`func (o *GetEndpoints200ResponseEndpointsInner) GetCustomProperty23Ok() (*string, bool)`

GetCustomProperty23Ok returns a tuple with the CustomProperty23 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomProperty23

`func (o *GetEndpoints200ResponseEndpointsInner) SetCustomProperty23(v string)`

SetCustomProperty23 sets CustomProperty23 field to given value.

### HasCustomProperty23

`func (o *GetEndpoints200ResponseEndpointsInner) HasCustomProperty23() bool`

HasCustomProperty23 returns a boolean if a field has been set.

### GetCustomProperty24

`func (o *GetEndpoints200ResponseEndpointsInner) GetCustomProperty24() string`

GetCustomProperty24 returns the CustomProperty24 field if non-nil, zero value otherwise.

### GetCustomProperty24Ok

`func (o *GetEndpoints200ResponseEndpointsInner) GetCustomProperty24Ok() (*string, bool)`

GetCustomProperty24Ok returns a tuple with the CustomProperty24 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomProperty24

`func (o *GetEndpoints200ResponseEndpointsInner) SetCustomProperty24(v string)`

SetCustomProperty24 sets CustomProperty24 field to given value.

### HasCustomProperty24

`func (o *GetEndpoints200ResponseEndpointsInner) HasCustomProperty24() bool`

HasCustomProperty24 returns a boolean if a field has been set.

### GetCustomProperty25

`func (o *GetEndpoints200ResponseEndpointsInner) GetCustomProperty25() string`

GetCustomProperty25 returns the CustomProperty25 field if non-nil, zero value otherwise.

### GetCustomProperty25Ok

`func (o *GetEndpoints200ResponseEndpointsInner) GetCustomProperty25Ok() (*string, bool)`

GetCustomProperty25Ok returns a tuple with the CustomProperty25 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomProperty25

`func (o *GetEndpoints200ResponseEndpointsInner) SetCustomProperty25(v string)`

SetCustomProperty25 sets CustomProperty25 field to given value.

### HasCustomProperty25

`func (o *GetEndpoints200ResponseEndpointsInner) HasCustomProperty25() bool`

HasCustomProperty25 returns a boolean if a field has been set.

### GetCustomProperty26

`func (o *GetEndpoints200ResponseEndpointsInner) GetCustomProperty26() string`

GetCustomProperty26 returns the CustomProperty26 field if non-nil, zero value otherwise.

### GetCustomProperty26Ok

`func (o *GetEndpoints200ResponseEndpointsInner) GetCustomProperty26Ok() (*string, bool)`

GetCustomProperty26Ok returns a tuple with the CustomProperty26 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomProperty26

`func (o *GetEndpoints200ResponseEndpointsInner) SetCustomProperty26(v string)`

SetCustomProperty26 sets CustomProperty26 field to given value.

### HasCustomProperty26

`func (o *GetEndpoints200ResponseEndpointsInner) HasCustomProperty26() bool`

HasCustomProperty26 returns a boolean if a field has been set.

### GetCustomProperty27

`func (o *GetEndpoints200ResponseEndpointsInner) GetCustomProperty27() string`

GetCustomProperty27 returns the CustomProperty27 field if non-nil, zero value otherwise.

### GetCustomProperty27Ok

`func (o *GetEndpoints200ResponseEndpointsInner) GetCustomProperty27Ok() (*string, bool)`

GetCustomProperty27Ok returns a tuple with the CustomProperty27 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomProperty27

`func (o *GetEndpoints200ResponseEndpointsInner) SetCustomProperty27(v string)`

SetCustomProperty27 sets CustomProperty27 field to given value.

### HasCustomProperty27

`func (o *GetEndpoints200ResponseEndpointsInner) HasCustomProperty27() bool`

HasCustomProperty27 returns a boolean if a field has been set.

### GetCustomProperty28

`func (o *GetEndpoints200ResponseEndpointsInner) GetCustomProperty28() string`

GetCustomProperty28 returns the CustomProperty28 field if non-nil, zero value otherwise.

### GetCustomProperty28Ok

`func (o *GetEndpoints200ResponseEndpointsInner) GetCustomProperty28Ok() (*string, bool)`

GetCustomProperty28Ok returns a tuple with the CustomProperty28 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomProperty28

`func (o *GetEndpoints200ResponseEndpointsInner) SetCustomProperty28(v string)`

SetCustomProperty28 sets CustomProperty28 field to given value.

### HasCustomProperty28

`func (o *GetEndpoints200ResponseEndpointsInner) HasCustomProperty28() bool`

HasCustomProperty28 returns a boolean if a field has been set.

### GetCustomProperty29

`func (o *GetEndpoints200ResponseEndpointsInner) GetCustomProperty29() string`

GetCustomProperty29 returns the CustomProperty29 field if non-nil, zero value otherwise.

### GetCustomProperty29Ok

`func (o *GetEndpoints200ResponseEndpointsInner) GetCustomProperty29Ok() (*string, bool)`

GetCustomProperty29Ok returns a tuple with the CustomProperty29 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomProperty29

`func (o *GetEndpoints200ResponseEndpointsInner) SetCustomProperty29(v string)`

SetCustomProperty29 sets CustomProperty29 field to given value.

### HasCustomProperty29

`func (o *GetEndpoints200ResponseEndpointsInner) HasCustomProperty29() bool`

HasCustomProperty29 returns a boolean if a field has been set.

### GetCustomProperty30

`func (o *GetEndpoints200ResponseEndpointsInner) GetCustomProperty30() string`

GetCustomProperty30 returns the CustomProperty30 field if non-nil, zero value otherwise.

### GetCustomProperty30Ok

`func (o *GetEndpoints200ResponseEndpointsInner) GetCustomProperty30Ok() (*string, bool)`

GetCustomProperty30Ok returns a tuple with the CustomProperty30 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomProperty30

`func (o *GetEndpoints200ResponseEndpointsInner) SetCustomProperty30(v string)`

SetCustomProperty30 sets CustomProperty30 field to given value.

### HasCustomProperty30

`func (o *GetEndpoints200ResponseEndpointsInner) HasCustomProperty30() bool`

HasCustomProperty30 returns a boolean if a field has been set.

### GetCustomproperty31

`func (o *GetEndpoints200ResponseEndpointsInner) GetCustomproperty31() string`

GetCustomproperty31 returns the Customproperty31 field if non-nil, zero value otherwise.

### GetCustomproperty31Ok

`func (o *GetEndpoints200ResponseEndpointsInner) GetCustomproperty31Ok() (*string, bool)`

GetCustomproperty31Ok returns a tuple with the Customproperty31 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty31

`func (o *GetEndpoints200ResponseEndpointsInner) SetCustomproperty31(v string)`

SetCustomproperty31 sets Customproperty31 field to given value.

### HasCustomproperty31

`func (o *GetEndpoints200ResponseEndpointsInner) HasCustomproperty31() bool`

HasCustomproperty31 returns a boolean if a field has been set.

### GetCustomproperty32

`func (o *GetEndpoints200ResponseEndpointsInner) GetCustomproperty32() string`

GetCustomproperty32 returns the Customproperty32 field if non-nil, zero value otherwise.

### GetCustomproperty32Ok

`func (o *GetEndpoints200ResponseEndpointsInner) GetCustomproperty32Ok() (*string, bool)`

GetCustomproperty32Ok returns a tuple with the Customproperty32 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty32

`func (o *GetEndpoints200ResponseEndpointsInner) SetCustomproperty32(v string)`

SetCustomproperty32 sets Customproperty32 field to given value.

### HasCustomproperty32

`func (o *GetEndpoints200ResponseEndpointsInner) HasCustomproperty32() bool`

HasCustomproperty32 returns a boolean if a field has been set.

### GetCustomproperty33

`func (o *GetEndpoints200ResponseEndpointsInner) GetCustomproperty33() string`

GetCustomproperty33 returns the Customproperty33 field if non-nil, zero value otherwise.

### GetCustomproperty33Ok

`func (o *GetEndpoints200ResponseEndpointsInner) GetCustomproperty33Ok() (*string, bool)`

GetCustomproperty33Ok returns a tuple with the Customproperty33 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty33

`func (o *GetEndpoints200ResponseEndpointsInner) SetCustomproperty33(v string)`

SetCustomproperty33 sets Customproperty33 field to given value.

### HasCustomproperty33

`func (o *GetEndpoints200ResponseEndpointsInner) HasCustomproperty33() bool`

HasCustomproperty33 returns a boolean if a field has been set.

### GetCustomproperty34

`func (o *GetEndpoints200ResponseEndpointsInner) GetCustomproperty34() string`

GetCustomproperty34 returns the Customproperty34 field if non-nil, zero value otherwise.

### GetCustomproperty34Ok

`func (o *GetEndpoints200ResponseEndpointsInner) GetCustomproperty34Ok() (*string, bool)`

GetCustomproperty34Ok returns a tuple with the Customproperty34 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty34

`func (o *GetEndpoints200ResponseEndpointsInner) SetCustomproperty34(v string)`

SetCustomproperty34 sets Customproperty34 field to given value.

### HasCustomproperty34

`func (o *GetEndpoints200ResponseEndpointsInner) HasCustomproperty34() bool`

HasCustomproperty34 returns a boolean if a field has been set.

### GetCustomproperty35

`func (o *GetEndpoints200ResponseEndpointsInner) GetCustomproperty35() string`

GetCustomproperty35 returns the Customproperty35 field if non-nil, zero value otherwise.

### GetCustomproperty35Ok

`func (o *GetEndpoints200ResponseEndpointsInner) GetCustomproperty35Ok() (*string, bool)`

GetCustomproperty35Ok returns a tuple with the Customproperty35 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty35

`func (o *GetEndpoints200ResponseEndpointsInner) SetCustomproperty35(v string)`

SetCustomproperty35 sets Customproperty35 field to given value.

### HasCustomproperty35

`func (o *GetEndpoints200ResponseEndpointsInner) HasCustomproperty35() bool`

HasCustomproperty35 returns a boolean if a field has been set.

### GetCustomproperty36

`func (o *GetEndpoints200ResponseEndpointsInner) GetCustomproperty36() string`

GetCustomproperty36 returns the Customproperty36 field if non-nil, zero value otherwise.

### GetCustomproperty36Ok

`func (o *GetEndpoints200ResponseEndpointsInner) GetCustomproperty36Ok() (*string, bool)`

GetCustomproperty36Ok returns a tuple with the Customproperty36 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty36

`func (o *GetEndpoints200ResponseEndpointsInner) SetCustomproperty36(v string)`

SetCustomproperty36 sets Customproperty36 field to given value.

### HasCustomproperty36

`func (o *GetEndpoints200ResponseEndpointsInner) HasCustomproperty36() bool`

HasCustomproperty36 returns a boolean if a field has been set.

### GetCustomproperty37

`func (o *GetEndpoints200ResponseEndpointsInner) GetCustomproperty37() string`

GetCustomproperty37 returns the Customproperty37 field if non-nil, zero value otherwise.

### GetCustomproperty37Ok

`func (o *GetEndpoints200ResponseEndpointsInner) GetCustomproperty37Ok() (*string, bool)`

GetCustomproperty37Ok returns a tuple with the Customproperty37 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty37

`func (o *GetEndpoints200ResponseEndpointsInner) SetCustomproperty37(v string)`

SetCustomproperty37 sets Customproperty37 field to given value.

### HasCustomproperty37

`func (o *GetEndpoints200ResponseEndpointsInner) HasCustomproperty37() bool`

HasCustomproperty37 returns a boolean if a field has been set.

### GetCustomproperty38

`func (o *GetEndpoints200ResponseEndpointsInner) GetCustomproperty38() string`

GetCustomproperty38 returns the Customproperty38 field if non-nil, zero value otherwise.

### GetCustomproperty38Ok

`func (o *GetEndpoints200ResponseEndpointsInner) GetCustomproperty38Ok() (*string, bool)`

GetCustomproperty38Ok returns a tuple with the Customproperty38 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty38

`func (o *GetEndpoints200ResponseEndpointsInner) SetCustomproperty38(v string)`

SetCustomproperty38 sets Customproperty38 field to given value.

### HasCustomproperty38

`func (o *GetEndpoints200ResponseEndpointsInner) HasCustomproperty38() bool`

HasCustomproperty38 returns a boolean if a field has been set.

### GetCustomproperty39

`func (o *GetEndpoints200ResponseEndpointsInner) GetCustomproperty39() string`

GetCustomproperty39 returns the Customproperty39 field if non-nil, zero value otherwise.

### GetCustomproperty39Ok

`func (o *GetEndpoints200ResponseEndpointsInner) GetCustomproperty39Ok() (*string, bool)`

GetCustomproperty39Ok returns a tuple with the Customproperty39 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty39

`func (o *GetEndpoints200ResponseEndpointsInner) SetCustomproperty39(v string)`

SetCustomproperty39 sets Customproperty39 field to given value.

### HasCustomproperty39

`func (o *GetEndpoints200ResponseEndpointsInner) HasCustomproperty39() bool`

HasCustomproperty39 returns a boolean if a field has been set.

### GetCustomproperty40

`func (o *GetEndpoints200ResponseEndpointsInner) GetCustomproperty40() string`

GetCustomproperty40 returns the Customproperty40 field if non-nil, zero value otherwise.

### GetCustomproperty40Ok

`func (o *GetEndpoints200ResponseEndpointsInner) GetCustomproperty40Ok() (*string, bool)`

GetCustomproperty40Ok returns a tuple with the Customproperty40 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty40

`func (o *GetEndpoints200ResponseEndpointsInner) SetCustomproperty40(v string)`

SetCustomproperty40 sets Customproperty40 field to given value.

### HasCustomproperty40

`func (o *GetEndpoints200ResponseEndpointsInner) HasCustomproperty40() bool`

HasCustomproperty40 returns a boolean if a field has been set.

### GetCustomproperty41

`func (o *GetEndpoints200ResponseEndpointsInner) GetCustomproperty41() string`

GetCustomproperty41 returns the Customproperty41 field if non-nil, zero value otherwise.

### GetCustomproperty41Ok

`func (o *GetEndpoints200ResponseEndpointsInner) GetCustomproperty41Ok() (*string, bool)`

GetCustomproperty41Ok returns a tuple with the Customproperty41 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty41

`func (o *GetEndpoints200ResponseEndpointsInner) SetCustomproperty41(v string)`

SetCustomproperty41 sets Customproperty41 field to given value.

### HasCustomproperty41

`func (o *GetEndpoints200ResponseEndpointsInner) HasCustomproperty41() bool`

HasCustomproperty41 returns a boolean if a field has been set.

### GetCustomproperty42

`func (o *GetEndpoints200ResponseEndpointsInner) GetCustomproperty42() string`

GetCustomproperty42 returns the Customproperty42 field if non-nil, zero value otherwise.

### GetCustomproperty42Ok

`func (o *GetEndpoints200ResponseEndpointsInner) GetCustomproperty42Ok() (*string, bool)`

GetCustomproperty42Ok returns a tuple with the Customproperty42 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty42

`func (o *GetEndpoints200ResponseEndpointsInner) SetCustomproperty42(v string)`

SetCustomproperty42 sets Customproperty42 field to given value.

### HasCustomproperty42

`func (o *GetEndpoints200ResponseEndpointsInner) HasCustomproperty42() bool`

HasCustomproperty42 returns a boolean if a field has been set.

### GetCustomproperty43

`func (o *GetEndpoints200ResponseEndpointsInner) GetCustomproperty43() string`

GetCustomproperty43 returns the Customproperty43 field if non-nil, zero value otherwise.

### GetCustomproperty43Ok

`func (o *GetEndpoints200ResponseEndpointsInner) GetCustomproperty43Ok() (*string, bool)`

GetCustomproperty43Ok returns a tuple with the Customproperty43 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty43

`func (o *GetEndpoints200ResponseEndpointsInner) SetCustomproperty43(v string)`

SetCustomproperty43 sets Customproperty43 field to given value.

### HasCustomproperty43

`func (o *GetEndpoints200ResponseEndpointsInner) HasCustomproperty43() bool`

HasCustomproperty43 returns a boolean if a field has been set.

### GetCustomproperty44

`func (o *GetEndpoints200ResponseEndpointsInner) GetCustomproperty44() string`

GetCustomproperty44 returns the Customproperty44 field if non-nil, zero value otherwise.

### GetCustomproperty44Ok

`func (o *GetEndpoints200ResponseEndpointsInner) GetCustomproperty44Ok() (*string, bool)`

GetCustomproperty44Ok returns a tuple with the Customproperty44 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty44

`func (o *GetEndpoints200ResponseEndpointsInner) SetCustomproperty44(v string)`

SetCustomproperty44 sets Customproperty44 field to given value.

### HasCustomproperty44

`func (o *GetEndpoints200ResponseEndpointsInner) HasCustomproperty44() bool`

HasCustomproperty44 returns a boolean if a field has been set.

### GetCustomproperty45

`func (o *GetEndpoints200ResponseEndpointsInner) GetCustomproperty45() string`

GetCustomproperty45 returns the Customproperty45 field if non-nil, zero value otherwise.

### GetCustomproperty45Ok

`func (o *GetEndpoints200ResponseEndpointsInner) GetCustomproperty45Ok() (*string, bool)`

GetCustomproperty45Ok returns a tuple with the Customproperty45 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty45

`func (o *GetEndpoints200ResponseEndpointsInner) SetCustomproperty45(v string)`

SetCustomproperty45 sets Customproperty45 field to given value.

### HasCustomproperty45

`func (o *GetEndpoints200ResponseEndpointsInner) HasCustomproperty45() bool`

HasCustomproperty45 returns a boolean if a field has been set.

### GetAccountCustomProperty1Label

`func (o *GetEndpoints200ResponseEndpointsInner) GetAccountCustomProperty1Label() string`

GetAccountCustomProperty1Label returns the AccountCustomProperty1Label field if non-nil, zero value otherwise.

### GetAccountCustomProperty1LabelOk

`func (o *GetEndpoints200ResponseEndpointsInner) GetAccountCustomProperty1LabelOk() (*string, bool)`

GetAccountCustomProperty1LabelOk returns a tuple with the AccountCustomProperty1Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountCustomProperty1Label

`func (o *GetEndpoints200ResponseEndpointsInner) SetAccountCustomProperty1Label(v string)`

SetAccountCustomProperty1Label sets AccountCustomProperty1Label field to given value.

### HasAccountCustomProperty1Label

`func (o *GetEndpoints200ResponseEndpointsInner) HasAccountCustomProperty1Label() bool`

HasAccountCustomProperty1Label returns a boolean if a field has been set.

### GetAccountCustomProperty2Label

`func (o *GetEndpoints200ResponseEndpointsInner) GetAccountCustomProperty2Label() string`

GetAccountCustomProperty2Label returns the AccountCustomProperty2Label field if non-nil, zero value otherwise.

### GetAccountCustomProperty2LabelOk

`func (o *GetEndpoints200ResponseEndpointsInner) GetAccountCustomProperty2LabelOk() (*string, bool)`

GetAccountCustomProperty2LabelOk returns a tuple with the AccountCustomProperty2Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountCustomProperty2Label

`func (o *GetEndpoints200ResponseEndpointsInner) SetAccountCustomProperty2Label(v string)`

SetAccountCustomProperty2Label sets AccountCustomProperty2Label field to given value.

### HasAccountCustomProperty2Label

`func (o *GetEndpoints200ResponseEndpointsInner) HasAccountCustomProperty2Label() bool`

HasAccountCustomProperty2Label returns a boolean if a field has been set.

### GetAccountCustomProperty3Label

`func (o *GetEndpoints200ResponseEndpointsInner) GetAccountCustomProperty3Label() string`

GetAccountCustomProperty3Label returns the AccountCustomProperty3Label field if non-nil, zero value otherwise.

### GetAccountCustomProperty3LabelOk

`func (o *GetEndpoints200ResponseEndpointsInner) GetAccountCustomProperty3LabelOk() (*string, bool)`

GetAccountCustomProperty3LabelOk returns a tuple with the AccountCustomProperty3Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountCustomProperty3Label

`func (o *GetEndpoints200ResponseEndpointsInner) SetAccountCustomProperty3Label(v string)`

SetAccountCustomProperty3Label sets AccountCustomProperty3Label field to given value.

### HasAccountCustomProperty3Label

`func (o *GetEndpoints200ResponseEndpointsInner) HasAccountCustomProperty3Label() bool`

HasAccountCustomProperty3Label returns a boolean if a field has been set.

### GetAccountCustomProperty4Label

`func (o *GetEndpoints200ResponseEndpointsInner) GetAccountCustomProperty4Label() string`

GetAccountCustomProperty4Label returns the AccountCustomProperty4Label field if non-nil, zero value otherwise.

### GetAccountCustomProperty4LabelOk

`func (o *GetEndpoints200ResponseEndpointsInner) GetAccountCustomProperty4LabelOk() (*string, bool)`

GetAccountCustomProperty4LabelOk returns a tuple with the AccountCustomProperty4Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountCustomProperty4Label

`func (o *GetEndpoints200ResponseEndpointsInner) SetAccountCustomProperty4Label(v string)`

SetAccountCustomProperty4Label sets AccountCustomProperty4Label field to given value.

### HasAccountCustomProperty4Label

`func (o *GetEndpoints200ResponseEndpointsInner) HasAccountCustomProperty4Label() bool`

HasAccountCustomProperty4Label returns a boolean if a field has been set.

### GetAccountCustomProperty5Label

`func (o *GetEndpoints200ResponseEndpointsInner) GetAccountCustomProperty5Label() string`

GetAccountCustomProperty5Label returns the AccountCustomProperty5Label field if non-nil, zero value otherwise.

### GetAccountCustomProperty5LabelOk

`func (o *GetEndpoints200ResponseEndpointsInner) GetAccountCustomProperty5LabelOk() (*string, bool)`

GetAccountCustomProperty5LabelOk returns a tuple with the AccountCustomProperty5Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountCustomProperty5Label

`func (o *GetEndpoints200ResponseEndpointsInner) SetAccountCustomProperty5Label(v string)`

SetAccountCustomProperty5Label sets AccountCustomProperty5Label field to given value.

### HasAccountCustomProperty5Label

`func (o *GetEndpoints200ResponseEndpointsInner) HasAccountCustomProperty5Label() bool`

HasAccountCustomProperty5Label returns a boolean if a field has been set.

### GetAccountCustomProperty6Label

`func (o *GetEndpoints200ResponseEndpointsInner) GetAccountCustomProperty6Label() string`

GetAccountCustomProperty6Label returns the AccountCustomProperty6Label field if non-nil, zero value otherwise.

### GetAccountCustomProperty6LabelOk

`func (o *GetEndpoints200ResponseEndpointsInner) GetAccountCustomProperty6LabelOk() (*string, bool)`

GetAccountCustomProperty6LabelOk returns a tuple with the AccountCustomProperty6Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountCustomProperty6Label

`func (o *GetEndpoints200ResponseEndpointsInner) SetAccountCustomProperty6Label(v string)`

SetAccountCustomProperty6Label sets AccountCustomProperty6Label field to given value.

### HasAccountCustomProperty6Label

`func (o *GetEndpoints200ResponseEndpointsInner) HasAccountCustomProperty6Label() bool`

HasAccountCustomProperty6Label returns a boolean if a field has been set.

### GetAccountCustomProperty7Label

`func (o *GetEndpoints200ResponseEndpointsInner) GetAccountCustomProperty7Label() string`

GetAccountCustomProperty7Label returns the AccountCustomProperty7Label field if non-nil, zero value otherwise.

### GetAccountCustomProperty7LabelOk

`func (o *GetEndpoints200ResponseEndpointsInner) GetAccountCustomProperty7LabelOk() (*string, bool)`

GetAccountCustomProperty7LabelOk returns a tuple with the AccountCustomProperty7Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountCustomProperty7Label

`func (o *GetEndpoints200ResponseEndpointsInner) SetAccountCustomProperty7Label(v string)`

SetAccountCustomProperty7Label sets AccountCustomProperty7Label field to given value.

### HasAccountCustomProperty7Label

`func (o *GetEndpoints200ResponseEndpointsInner) HasAccountCustomProperty7Label() bool`

HasAccountCustomProperty7Label returns a boolean if a field has been set.

### GetAccountCustomProperty8Label

`func (o *GetEndpoints200ResponseEndpointsInner) GetAccountCustomProperty8Label() string`

GetAccountCustomProperty8Label returns the AccountCustomProperty8Label field if non-nil, zero value otherwise.

### GetAccountCustomProperty8LabelOk

`func (o *GetEndpoints200ResponseEndpointsInner) GetAccountCustomProperty8LabelOk() (*string, bool)`

GetAccountCustomProperty8LabelOk returns a tuple with the AccountCustomProperty8Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountCustomProperty8Label

`func (o *GetEndpoints200ResponseEndpointsInner) SetAccountCustomProperty8Label(v string)`

SetAccountCustomProperty8Label sets AccountCustomProperty8Label field to given value.

### HasAccountCustomProperty8Label

`func (o *GetEndpoints200ResponseEndpointsInner) HasAccountCustomProperty8Label() bool`

HasAccountCustomProperty8Label returns a boolean if a field has been set.

### GetAccountCustomProperty9Label

`func (o *GetEndpoints200ResponseEndpointsInner) GetAccountCustomProperty9Label() string`

GetAccountCustomProperty9Label returns the AccountCustomProperty9Label field if non-nil, zero value otherwise.

### GetAccountCustomProperty9LabelOk

`func (o *GetEndpoints200ResponseEndpointsInner) GetAccountCustomProperty9LabelOk() (*string, bool)`

GetAccountCustomProperty9LabelOk returns a tuple with the AccountCustomProperty9Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountCustomProperty9Label

`func (o *GetEndpoints200ResponseEndpointsInner) SetAccountCustomProperty9Label(v string)`

SetAccountCustomProperty9Label sets AccountCustomProperty9Label field to given value.

### HasAccountCustomProperty9Label

`func (o *GetEndpoints200ResponseEndpointsInner) HasAccountCustomProperty9Label() bool`

HasAccountCustomProperty9Label returns a boolean if a field has been set.

### GetAccountCustomProperty10Label

`func (o *GetEndpoints200ResponseEndpointsInner) GetAccountCustomProperty10Label() string`

GetAccountCustomProperty10Label returns the AccountCustomProperty10Label field if non-nil, zero value otherwise.

### GetAccountCustomProperty10LabelOk

`func (o *GetEndpoints200ResponseEndpointsInner) GetAccountCustomProperty10LabelOk() (*string, bool)`

GetAccountCustomProperty10LabelOk returns a tuple with the AccountCustomProperty10Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountCustomProperty10Label

`func (o *GetEndpoints200ResponseEndpointsInner) SetAccountCustomProperty10Label(v string)`

SetAccountCustomProperty10Label sets AccountCustomProperty10Label field to given value.

### HasAccountCustomProperty10Label

`func (o *GetEndpoints200ResponseEndpointsInner) HasAccountCustomProperty10Label() bool`

HasAccountCustomProperty10Label returns a boolean if a field has been set.

### GetAccountCustomProperty11Label

`func (o *GetEndpoints200ResponseEndpointsInner) GetAccountCustomProperty11Label() string`

GetAccountCustomProperty11Label returns the AccountCustomProperty11Label field if non-nil, zero value otherwise.

### GetAccountCustomProperty11LabelOk

`func (o *GetEndpoints200ResponseEndpointsInner) GetAccountCustomProperty11LabelOk() (*string, bool)`

GetAccountCustomProperty11LabelOk returns a tuple with the AccountCustomProperty11Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountCustomProperty11Label

`func (o *GetEndpoints200ResponseEndpointsInner) SetAccountCustomProperty11Label(v string)`

SetAccountCustomProperty11Label sets AccountCustomProperty11Label field to given value.

### HasAccountCustomProperty11Label

`func (o *GetEndpoints200ResponseEndpointsInner) HasAccountCustomProperty11Label() bool`

HasAccountCustomProperty11Label returns a boolean if a field has been set.

### GetAccountCustomProperty12Label

`func (o *GetEndpoints200ResponseEndpointsInner) GetAccountCustomProperty12Label() string`

GetAccountCustomProperty12Label returns the AccountCustomProperty12Label field if non-nil, zero value otherwise.

### GetAccountCustomProperty12LabelOk

`func (o *GetEndpoints200ResponseEndpointsInner) GetAccountCustomProperty12LabelOk() (*string, bool)`

GetAccountCustomProperty12LabelOk returns a tuple with the AccountCustomProperty12Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountCustomProperty12Label

`func (o *GetEndpoints200ResponseEndpointsInner) SetAccountCustomProperty12Label(v string)`

SetAccountCustomProperty12Label sets AccountCustomProperty12Label field to given value.

### HasAccountCustomProperty12Label

`func (o *GetEndpoints200ResponseEndpointsInner) HasAccountCustomProperty12Label() bool`

HasAccountCustomProperty12Label returns a boolean if a field has been set.

### GetAccountCustomProperty13Label

`func (o *GetEndpoints200ResponseEndpointsInner) GetAccountCustomProperty13Label() string`

GetAccountCustomProperty13Label returns the AccountCustomProperty13Label field if non-nil, zero value otherwise.

### GetAccountCustomProperty13LabelOk

`func (o *GetEndpoints200ResponseEndpointsInner) GetAccountCustomProperty13LabelOk() (*string, bool)`

GetAccountCustomProperty13LabelOk returns a tuple with the AccountCustomProperty13Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountCustomProperty13Label

`func (o *GetEndpoints200ResponseEndpointsInner) SetAccountCustomProperty13Label(v string)`

SetAccountCustomProperty13Label sets AccountCustomProperty13Label field to given value.

### HasAccountCustomProperty13Label

`func (o *GetEndpoints200ResponseEndpointsInner) HasAccountCustomProperty13Label() bool`

HasAccountCustomProperty13Label returns a boolean if a field has been set.

### GetAccountCustomProperty14Label

`func (o *GetEndpoints200ResponseEndpointsInner) GetAccountCustomProperty14Label() string`

GetAccountCustomProperty14Label returns the AccountCustomProperty14Label field if non-nil, zero value otherwise.

### GetAccountCustomProperty14LabelOk

`func (o *GetEndpoints200ResponseEndpointsInner) GetAccountCustomProperty14LabelOk() (*string, bool)`

GetAccountCustomProperty14LabelOk returns a tuple with the AccountCustomProperty14Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountCustomProperty14Label

`func (o *GetEndpoints200ResponseEndpointsInner) SetAccountCustomProperty14Label(v string)`

SetAccountCustomProperty14Label sets AccountCustomProperty14Label field to given value.

### HasAccountCustomProperty14Label

`func (o *GetEndpoints200ResponseEndpointsInner) HasAccountCustomProperty14Label() bool`

HasAccountCustomProperty14Label returns a boolean if a field has been set.

### GetAccountCustomProperty15Label

`func (o *GetEndpoints200ResponseEndpointsInner) GetAccountCustomProperty15Label() string`

GetAccountCustomProperty15Label returns the AccountCustomProperty15Label field if non-nil, zero value otherwise.

### GetAccountCustomProperty15LabelOk

`func (o *GetEndpoints200ResponseEndpointsInner) GetAccountCustomProperty15LabelOk() (*string, bool)`

GetAccountCustomProperty15LabelOk returns a tuple with the AccountCustomProperty15Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountCustomProperty15Label

`func (o *GetEndpoints200ResponseEndpointsInner) SetAccountCustomProperty15Label(v string)`

SetAccountCustomProperty15Label sets AccountCustomProperty15Label field to given value.

### HasAccountCustomProperty15Label

`func (o *GetEndpoints200ResponseEndpointsInner) HasAccountCustomProperty15Label() bool`

HasAccountCustomProperty15Label returns a boolean if a field has been set.

### GetAccountCustomProperty16Label

`func (o *GetEndpoints200ResponseEndpointsInner) GetAccountCustomProperty16Label() string`

GetAccountCustomProperty16Label returns the AccountCustomProperty16Label field if non-nil, zero value otherwise.

### GetAccountCustomProperty16LabelOk

`func (o *GetEndpoints200ResponseEndpointsInner) GetAccountCustomProperty16LabelOk() (*string, bool)`

GetAccountCustomProperty16LabelOk returns a tuple with the AccountCustomProperty16Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountCustomProperty16Label

`func (o *GetEndpoints200ResponseEndpointsInner) SetAccountCustomProperty16Label(v string)`

SetAccountCustomProperty16Label sets AccountCustomProperty16Label field to given value.

### HasAccountCustomProperty16Label

`func (o *GetEndpoints200ResponseEndpointsInner) HasAccountCustomProperty16Label() bool`

HasAccountCustomProperty16Label returns a boolean if a field has been set.

### GetAccountCustomProperty17Label

`func (o *GetEndpoints200ResponseEndpointsInner) GetAccountCustomProperty17Label() string`

GetAccountCustomProperty17Label returns the AccountCustomProperty17Label field if non-nil, zero value otherwise.

### GetAccountCustomProperty17LabelOk

`func (o *GetEndpoints200ResponseEndpointsInner) GetAccountCustomProperty17LabelOk() (*string, bool)`

GetAccountCustomProperty17LabelOk returns a tuple with the AccountCustomProperty17Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountCustomProperty17Label

`func (o *GetEndpoints200ResponseEndpointsInner) SetAccountCustomProperty17Label(v string)`

SetAccountCustomProperty17Label sets AccountCustomProperty17Label field to given value.

### HasAccountCustomProperty17Label

`func (o *GetEndpoints200ResponseEndpointsInner) HasAccountCustomProperty17Label() bool`

HasAccountCustomProperty17Label returns a boolean if a field has been set.

### GetAccountCustomProperty18Label

`func (o *GetEndpoints200ResponseEndpointsInner) GetAccountCustomProperty18Label() string`

GetAccountCustomProperty18Label returns the AccountCustomProperty18Label field if non-nil, zero value otherwise.

### GetAccountCustomProperty18LabelOk

`func (o *GetEndpoints200ResponseEndpointsInner) GetAccountCustomProperty18LabelOk() (*string, bool)`

GetAccountCustomProperty18LabelOk returns a tuple with the AccountCustomProperty18Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountCustomProperty18Label

`func (o *GetEndpoints200ResponseEndpointsInner) SetAccountCustomProperty18Label(v string)`

SetAccountCustomProperty18Label sets AccountCustomProperty18Label field to given value.

### HasAccountCustomProperty18Label

`func (o *GetEndpoints200ResponseEndpointsInner) HasAccountCustomProperty18Label() bool`

HasAccountCustomProperty18Label returns a boolean if a field has been set.

### GetAccountCustomProperty19Label

`func (o *GetEndpoints200ResponseEndpointsInner) GetAccountCustomProperty19Label() string`

GetAccountCustomProperty19Label returns the AccountCustomProperty19Label field if non-nil, zero value otherwise.

### GetAccountCustomProperty19LabelOk

`func (o *GetEndpoints200ResponseEndpointsInner) GetAccountCustomProperty19LabelOk() (*string, bool)`

GetAccountCustomProperty19LabelOk returns a tuple with the AccountCustomProperty19Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountCustomProperty19Label

`func (o *GetEndpoints200ResponseEndpointsInner) SetAccountCustomProperty19Label(v string)`

SetAccountCustomProperty19Label sets AccountCustomProperty19Label field to given value.

### HasAccountCustomProperty19Label

`func (o *GetEndpoints200ResponseEndpointsInner) HasAccountCustomProperty19Label() bool`

HasAccountCustomProperty19Label returns a boolean if a field has been set.

### GetAccountCustomProperty20Label

`func (o *GetEndpoints200ResponseEndpointsInner) GetAccountCustomProperty20Label() string`

GetAccountCustomProperty20Label returns the AccountCustomProperty20Label field if non-nil, zero value otherwise.

### GetAccountCustomProperty20LabelOk

`func (o *GetEndpoints200ResponseEndpointsInner) GetAccountCustomProperty20LabelOk() (*string, bool)`

GetAccountCustomProperty20LabelOk returns a tuple with the AccountCustomProperty20Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountCustomProperty20Label

`func (o *GetEndpoints200ResponseEndpointsInner) SetAccountCustomProperty20Label(v string)`

SetAccountCustomProperty20Label sets AccountCustomProperty20Label field to given value.

### HasAccountCustomProperty20Label

`func (o *GetEndpoints200ResponseEndpointsInner) HasAccountCustomProperty20Label() bool`

HasAccountCustomProperty20Label returns a boolean if a field has been set.

### GetAccountCustomProperty21Label

`func (o *GetEndpoints200ResponseEndpointsInner) GetAccountCustomProperty21Label() string`

GetAccountCustomProperty21Label returns the AccountCustomProperty21Label field if non-nil, zero value otherwise.

### GetAccountCustomProperty21LabelOk

`func (o *GetEndpoints200ResponseEndpointsInner) GetAccountCustomProperty21LabelOk() (*string, bool)`

GetAccountCustomProperty21LabelOk returns a tuple with the AccountCustomProperty21Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountCustomProperty21Label

`func (o *GetEndpoints200ResponseEndpointsInner) SetAccountCustomProperty21Label(v string)`

SetAccountCustomProperty21Label sets AccountCustomProperty21Label field to given value.

### HasAccountCustomProperty21Label

`func (o *GetEndpoints200ResponseEndpointsInner) HasAccountCustomProperty21Label() bool`

HasAccountCustomProperty21Label returns a boolean if a field has been set.

### GetAccountCustomProperty22Label

`func (o *GetEndpoints200ResponseEndpointsInner) GetAccountCustomProperty22Label() string`

GetAccountCustomProperty22Label returns the AccountCustomProperty22Label field if non-nil, zero value otherwise.

### GetAccountCustomProperty22LabelOk

`func (o *GetEndpoints200ResponseEndpointsInner) GetAccountCustomProperty22LabelOk() (*string, bool)`

GetAccountCustomProperty22LabelOk returns a tuple with the AccountCustomProperty22Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountCustomProperty22Label

`func (o *GetEndpoints200ResponseEndpointsInner) SetAccountCustomProperty22Label(v string)`

SetAccountCustomProperty22Label sets AccountCustomProperty22Label field to given value.

### HasAccountCustomProperty22Label

`func (o *GetEndpoints200ResponseEndpointsInner) HasAccountCustomProperty22Label() bool`

HasAccountCustomProperty22Label returns a boolean if a field has been set.

### GetAccountCustomProperty23Label

`func (o *GetEndpoints200ResponseEndpointsInner) GetAccountCustomProperty23Label() string`

GetAccountCustomProperty23Label returns the AccountCustomProperty23Label field if non-nil, zero value otherwise.

### GetAccountCustomProperty23LabelOk

`func (o *GetEndpoints200ResponseEndpointsInner) GetAccountCustomProperty23LabelOk() (*string, bool)`

GetAccountCustomProperty23LabelOk returns a tuple with the AccountCustomProperty23Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountCustomProperty23Label

`func (o *GetEndpoints200ResponseEndpointsInner) SetAccountCustomProperty23Label(v string)`

SetAccountCustomProperty23Label sets AccountCustomProperty23Label field to given value.

### HasAccountCustomProperty23Label

`func (o *GetEndpoints200ResponseEndpointsInner) HasAccountCustomProperty23Label() bool`

HasAccountCustomProperty23Label returns a boolean if a field has been set.

### GetAccountCustomProperty24Label

`func (o *GetEndpoints200ResponseEndpointsInner) GetAccountCustomProperty24Label() string`

GetAccountCustomProperty24Label returns the AccountCustomProperty24Label field if non-nil, zero value otherwise.

### GetAccountCustomProperty24LabelOk

`func (o *GetEndpoints200ResponseEndpointsInner) GetAccountCustomProperty24LabelOk() (*string, bool)`

GetAccountCustomProperty24LabelOk returns a tuple with the AccountCustomProperty24Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountCustomProperty24Label

`func (o *GetEndpoints200ResponseEndpointsInner) SetAccountCustomProperty24Label(v string)`

SetAccountCustomProperty24Label sets AccountCustomProperty24Label field to given value.

### HasAccountCustomProperty24Label

`func (o *GetEndpoints200ResponseEndpointsInner) HasAccountCustomProperty24Label() bool`

HasAccountCustomProperty24Label returns a boolean if a field has been set.

### GetAccountCustomProperty25Label

`func (o *GetEndpoints200ResponseEndpointsInner) GetAccountCustomProperty25Label() string`

GetAccountCustomProperty25Label returns the AccountCustomProperty25Label field if non-nil, zero value otherwise.

### GetAccountCustomProperty25LabelOk

`func (o *GetEndpoints200ResponseEndpointsInner) GetAccountCustomProperty25LabelOk() (*string, bool)`

GetAccountCustomProperty25LabelOk returns a tuple with the AccountCustomProperty25Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountCustomProperty25Label

`func (o *GetEndpoints200ResponseEndpointsInner) SetAccountCustomProperty25Label(v string)`

SetAccountCustomProperty25Label sets AccountCustomProperty25Label field to given value.

### HasAccountCustomProperty25Label

`func (o *GetEndpoints200ResponseEndpointsInner) HasAccountCustomProperty25Label() bool`

HasAccountCustomProperty25Label returns a boolean if a field has been set.

### GetAccountCustomProperty26Label

`func (o *GetEndpoints200ResponseEndpointsInner) GetAccountCustomProperty26Label() string`

GetAccountCustomProperty26Label returns the AccountCustomProperty26Label field if non-nil, zero value otherwise.

### GetAccountCustomProperty26LabelOk

`func (o *GetEndpoints200ResponseEndpointsInner) GetAccountCustomProperty26LabelOk() (*string, bool)`

GetAccountCustomProperty26LabelOk returns a tuple with the AccountCustomProperty26Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountCustomProperty26Label

`func (o *GetEndpoints200ResponseEndpointsInner) SetAccountCustomProperty26Label(v string)`

SetAccountCustomProperty26Label sets AccountCustomProperty26Label field to given value.

### HasAccountCustomProperty26Label

`func (o *GetEndpoints200ResponseEndpointsInner) HasAccountCustomProperty26Label() bool`

HasAccountCustomProperty26Label returns a boolean if a field has been set.

### GetAccountCustomProperty27Label

`func (o *GetEndpoints200ResponseEndpointsInner) GetAccountCustomProperty27Label() string`

GetAccountCustomProperty27Label returns the AccountCustomProperty27Label field if non-nil, zero value otherwise.

### GetAccountCustomProperty27LabelOk

`func (o *GetEndpoints200ResponseEndpointsInner) GetAccountCustomProperty27LabelOk() (*string, bool)`

GetAccountCustomProperty27LabelOk returns a tuple with the AccountCustomProperty27Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountCustomProperty27Label

`func (o *GetEndpoints200ResponseEndpointsInner) SetAccountCustomProperty27Label(v string)`

SetAccountCustomProperty27Label sets AccountCustomProperty27Label field to given value.

### HasAccountCustomProperty27Label

`func (o *GetEndpoints200ResponseEndpointsInner) HasAccountCustomProperty27Label() bool`

HasAccountCustomProperty27Label returns a boolean if a field has been set.

### GetAccountCustomProperty28Label

`func (o *GetEndpoints200ResponseEndpointsInner) GetAccountCustomProperty28Label() string`

GetAccountCustomProperty28Label returns the AccountCustomProperty28Label field if non-nil, zero value otherwise.

### GetAccountCustomProperty28LabelOk

`func (o *GetEndpoints200ResponseEndpointsInner) GetAccountCustomProperty28LabelOk() (*string, bool)`

GetAccountCustomProperty28LabelOk returns a tuple with the AccountCustomProperty28Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountCustomProperty28Label

`func (o *GetEndpoints200ResponseEndpointsInner) SetAccountCustomProperty28Label(v string)`

SetAccountCustomProperty28Label sets AccountCustomProperty28Label field to given value.

### HasAccountCustomProperty28Label

`func (o *GetEndpoints200ResponseEndpointsInner) HasAccountCustomProperty28Label() bool`

HasAccountCustomProperty28Label returns a boolean if a field has been set.

### GetAccountCustomProperty29Label

`func (o *GetEndpoints200ResponseEndpointsInner) GetAccountCustomProperty29Label() string`

GetAccountCustomProperty29Label returns the AccountCustomProperty29Label field if non-nil, zero value otherwise.

### GetAccountCustomProperty29LabelOk

`func (o *GetEndpoints200ResponseEndpointsInner) GetAccountCustomProperty29LabelOk() (*string, bool)`

GetAccountCustomProperty29LabelOk returns a tuple with the AccountCustomProperty29Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountCustomProperty29Label

`func (o *GetEndpoints200ResponseEndpointsInner) SetAccountCustomProperty29Label(v string)`

SetAccountCustomProperty29Label sets AccountCustomProperty29Label field to given value.

### HasAccountCustomProperty29Label

`func (o *GetEndpoints200ResponseEndpointsInner) HasAccountCustomProperty29Label() bool`

HasAccountCustomProperty29Label returns a boolean if a field has been set.

### GetAccountCustomProperty30Label

`func (o *GetEndpoints200ResponseEndpointsInner) GetAccountCustomProperty30Label() string`

GetAccountCustomProperty30Label returns the AccountCustomProperty30Label field if non-nil, zero value otherwise.

### GetAccountCustomProperty30LabelOk

`func (o *GetEndpoints200ResponseEndpointsInner) GetAccountCustomProperty30LabelOk() (*string, bool)`

GetAccountCustomProperty30LabelOk returns a tuple with the AccountCustomProperty30Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountCustomProperty30Label

`func (o *GetEndpoints200ResponseEndpointsInner) SetAccountCustomProperty30Label(v string)`

SetAccountCustomProperty30Label sets AccountCustomProperty30Label field to given value.

### HasAccountCustomProperty30Label

`func (o *GetEndpoints200ResponseEndpointsInner) HasAccountCustomProperty30Label() bool`

HasAccountCustomProperty30Label returns a boolean if a field has been set.

### GetCustomproperty31Label

`func (o *GetEndpoints200ResponseEndpointsInner) GetCustomproperty31Label() string`

GetCustomproperty31Label returns the Customproperty31Label field if non-nil, zero value otherwise.

### GetCustomproperty31LabelOk

`func (o *GetEndpoints200ResponseEndpointsInner) GetCustomproperty31LabelOk() (*string, bool)`

GetCustomproperty31LabelOk returns a tuple with the Customproperty31Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty31Label

`func (o *GetEndpoints200ResponseEndpointsInner) SetCustomproperty31Label(v string)`

SetCustomproperty31Label sets Customproperty31Label field to given value.

### HasCustomproperty31Label

`func (o *GetEndpoints200ResponseEndpointsInner) HasCustomproperty31Label() bool`

HasCustomproperty31Label returns a boolean if a field has been set.

### GetCustomproperty32Label

`func (o *GetEndpoints200ResponseEndpointsInner) GetCustomproperty32Label() string`

GetCustomproperty32Label returns the Customproperty32Label field if non-nil, zero value otherwise.

### GetCustomproperty32LabelOk

`func (o *GetEndpoints200ResponseEndpointsInner) GetCustomproperty32LabelOk() (*string, bool)`

GetCustomproperty32LabelOk returns a tuple with the Customproperty32Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty32Label

`func (o *GetEndpoints200ResponseEndpointsInner) SetCustomproperty32Label(v string)`

SetCustomproperty32Label sets Customproperty32Label field to given value.

### HasCustomproperty32Label

`func (o *GetEndpoints200ResponseEndpointsInner) HasCustomproperty32Label() bool`

HasCustomproperty32Label returns a boolean if a field has been set.

### GetCustomproperty33Label

`func (o *GetEndpoints200ResponseEndpointsInner) GetCustomproperty33Label() string`

GetCustomproperty33Label returns the Customproperty33Label field if non-nil, zero value otherwise.

### GetCustomproperty33LabelOk

`func (o *GetEndpoints200ResponseEndpointsInner) GetCustomproperty33LabelOk() (*string, bool)`

GetCustomproperty33LabelOk returns a tuple with the Customproperty33Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty33Label

`func (o *GetEndpoints200ResponseEndpointsInner) SetCustomproperty33Label(v string)`

SetCustomproperty33Label sets Customproperty33Label field to given value.

### HasCustomproperty33Label

`func (o *GetEndpoints200ResponseEndpointsInner) HasCustomproperty33Label() bool`

HasCustomproperty33Label returns a boolean if a field has been set.

### GetCustomproperty34Label

`func (o *GetEndpoints200ResponseEndpointsInner) GetCustomproperty34Label() string`

GetCustomproperty34Label returns the Customproperty34Label field if non-nil, zero value otherwise.

### GetCustomproperty34LabelOk

`func (o *GetEndpoints200ResponseEndpointsInner) GetCustomproperty34LabelOk() (*string, bool)`

GetCustomproperty34LabelOk returns a tuple with the Customproperty34Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty34Label

`func (o *GetEndpoints200ResponseEndpointsInner) SetCustomproperty34Label(v string)`

SetCustomproperty34Label sets Customproperty34Label field to given value.

### HasCustomproperty34Label

`func (o *GetEndpoints200ResponseEndpointsInner) HasCustomproperty34Label() bool`

HasCustomproperty34Label returns a boolean if a field has been set.

### GetCustomproperty35Label

`func (o *GetEndpoints200ResponseEndpointsInner) GetCustomproperty35Label() string`

GetCustomproperty35Label returns the Customproperty35Label field if non-nil, zero value otherwise.

### GetCustomproperty35LabelOk

`func (o *GetEndpoints200ResponseEndpointsInner) GetCustomproperty35LabelOk() (*string, bool)`

GetCustomproperty35LabelOk returns a tuple with the Customproperty35Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty35Label

`func (o *GetEndpoints200ResponseEndpointsInner) SetCustomproperty35Label(v string)`

SetCustomproperty35Label sets Customproperty35Label field to given value.

### HasCustomproperty35Label

`func (o *GetEndpoints200ResponseEndpointsInner) HasCustomproperty35Label() bool`

HasCustomproperty35Label returns a boolean if a field has been set.

### GetCustomproperty36Label

`func (o *GetEndpoints200ResponseEndpointsInner) GetCustomproperty36Label() string`

GetCustomproperty36Label returns the Customproperty36Label field if non-nil, zero value otherwise.

### GetCustomproperty36LabelOk

`func (o *GetEndpoints200ResponseEndpointsInner) GetCustomproperty36LabelOk() (*string, bool)`

GetCustomproperty36LabelOk returns a tuple with the Customproperty36Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty36Label

`func (o *GetEndpoints200ResponseEndpointsInner) SetCustomproperty36Label(v string)`

SetCustomproperty36Label sets Customproperty36Label field to given value.

### HasCustomproperty36Label

`func (o *GetEndpoints200ResponseEndpointsInner) HasCustomproperty36Label() bool`

HasCustomproperty36Label returns a boolean if a field has been set.

### GetCustomproperty37Label

`func (o *GetEndpoints200ResponseEndpointsInner) GetCustomproperty37Label() string`

GetCustomproperty37Label returns the Customproperty37Label field if non-nil, zero value otherwise.

### GetCustomproperty37LabelOk

`func (o *GetEndpoints200ResponseEndpointsInner) GetCustomproperty37LabelOk() (*string, bool)`

GetCustomproperty37LabelOk returns a tuple with the Customproperty37Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty37Label

`func (o *GetEndpoints200ResponseEndpointsInner) SetCustomproperty37Label(v string)`

SetCustomproperty37Label sets Customproperty37Label field to given value.

### HasCustomproperty37Label

`func (o *GetEndpoints200ResponseEndpointsInner) HasCustomproperty37Label() bool`

HasCustomproperty37Label returns a boolean if a field has been set.

### GetCustomproperty38Label

`func (o *GetEndpoints200ResponseEndpointsInner) GetCustomproperty38Label() string`

GetCustomproperty38Label returns the Customproperty38Label field if non-nil, zero value otherwise.

### GetCustomproperty38LabelOk

`func (o *GetEndpoints200ResponseEndpointsInner) GetCustomproperty38LabelOk() (*string, bool)`

GetCustomproperty38LabelOk returns a tuple with the Customproperty38Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty38Label

`func (o *GetEndpoints200ResponseEndpointsInner) SetCustomproperty38Label(v string)`

SetCustomproperty38Label sets Customproperty38Label field to given value.

### HasCustomproperty38Label

`func (o *GetEndpoints200ResponseEndpointsInner) HasCustomproperty38Label() bool`

HasCustomproperty38Label returns a boolean if a field has been set.

### GetCustomproperty39Label

`func (o *GetEndpoints200ResponseEndpointsInner) GetCustomproperty39Label() string`

GetCustomproperty39Label returns the Customproperty39Label field if non-nil, zero value otherwise.

### GetCustomproperty39LabelOk

`func (o *GetEndpoints200ResponseEndpointsInner) GetCustomproperty39LabelOk() (*string, bool)`

GetCustomproperty39LabelOk returns a tuple with the Customproperty39Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty39Label

`func (o *GetEndpoints200ResponseEndpointsInner) SetCustomproperty39Label(v string)`

SetCustomproperty39Label sets Customproperty39Label field to given value.

### HasCustomproperty39Label

`func (o *GetEndpoints200ResponseEndpointsInner) HasCustomproperty39Label() bool`

HasCustomproperty39Label returns a boolean if a field has been set.

### GetCustomproperty40Label

`func (o *GetEndpoints200ResponseEndpointsInner) GetCustomproperty40Label() string`

GetCustomproperty40Label returns the Customproperty40Label field if non-nil, zero value otherwise.

### GetCustomproperty40LabelOk

`func (o *GetEndpoints200ResponseEndpointsInner) GetCustomproperty40LabelOk() (*string, bool)`

GetCustomproperty40LabelOk returns a tuple with the Customproperty40Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty40Label

`func (o *GetEndpoints200ResponseEndpointsInner) SetCustomproperty40Label(v string)`

SetCustomproperty40Label sets Customproperty40Label field to given value.

### HasCustomproperty40Label

`func (o *GetEndpoints200ResponseEndpointsInner) HasCustomproperty40Label() bool`

HasCustomproperty40Label returns a boolean if a field has been set.

### GetCustomproperty41Label

`func (o *GetEndpoints200ResponseEndpointsInner) GetCustomproperty41Label() string`

GetCustomproperty41Label returns the Customproperty41Label field if non-nil, zero value otherwise.

### GetCustomproperty41LabelOk

`func (o *GetEndpoints200ResponseEndpointsInner) GetCustomproperty41LabelOk() (*string, bool)`

GetCustomproperty41LabelOk returns a tuple with the Customproperty41Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty41Label

`func (o *GetEndpoints200ResponseEndpointsInner) SetCustomproperty41Label(v string)`

SetCustomproperty41Label sets Customproperty41Label field to given value.

### HasCustomproperty41Label

`func (o *GetEndpoints200ResponseEndpointsInner) HasCustomproperty41Label() bool`

HasCustomproperty41Label returns a boolean if a field has been set.

### GetCustomproperty42Label

`func (o *GetEndpoints200ResponseEndpointsInner) GetCustomproperty42Label() string`

GetCustomproperty42Label returns the Customproperty42Label field if non-nil, zero value otherwise.

### GetCustomproperty42LabelOk

`func (o *GetEndpoints200ResponseEndpointsInner) GetCustomproperty42LabelOk() (*string, bool)`

GetCustomproperty42LabelOk returns a tuple with the Customproperty42Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty42Label

`func (o *GetEndpoints200ResponseEndpointsInner) SetCustomproperty42Label(v string)`

SetCustomproperty42Label sets Customproperty42Label field to given value.

### HasCustomproperty42Label

`func (o *GetEndpoints200ResponseEndpointsInner) HasCustomproperty42Label() bool`

HasCustomproperty42Label returns a boolean if a field has been set.

### GetCustomproperty43Label

`func (o *GetEndpoints200ResponseEndpointsInner) GetCustomproperty43Label() string`

GetCustomproperty43Label returns the Customproperty43Label field if non-nil, zero value otherwise.

### GetCustomproperty43LabelOk

`func (o *GetEndpoints200ResponseEndpointsInner) GetCustomproperty43LabelOk() (*string, bool)`

GetCustomproperty43LabelOk returns a tuple with the Customproperty43Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty43Label

`func (o *GetEndpoints200ResponseEndpointsInner) SetCustomproperty43Label(v string)`

SetCustomproperty43Label sets Customproperty43Label field to given value.

### HasCustomproperty43Label

`func (o *GetEndpoints200ResponseEndpointsInner) HasCustomproperty43Label() bool`

HasCustomproperty43Label returns a boolean if a field has been set.

### GetCustomproperty44Label

`func (o *GetEndpoints200ResponseEndpointsInner) GetCustomproperty44Label() string`

GetCustomproperty44Label returns the Customproperty44Label field if non-nil, zero value otherwise.

### GetCustomproperty44LabelOk

`func (o *GetEndpoints200ResponseEndpointsInner) GetCustomproperty44LabelOk() (*string, bool)`

GetCustomproperty44LabelOk returns a tuple with the Customproperty44Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty44Label

`func (o *GetEndpoints200ResponseEndpointsInner) SetCustomproperty44Label(v string)`

SetCustomproperty44Label sets Customproperty44Label field to given value.

### HasCustomproperty44Label

`func (o *GetEndpoints200ResponseEndpointsInner) HasCustomproperty44Label() bool`

HasCustomproperty44Label returns a boolean if a field has been set.

### GetCustomproperty45Label

`func (o *GetEndpoints200ResponseEndpointsInner) GetCustomproperty45Label() string`

GetCustomproperty45Label returns the Customproperty45Label field if non-nil, zero value otherwise.

### GetCustomproperty45LabelOk

`func (o *GetEndpoints200ResponseEndpointsInner) GetCustomproperty45LabelOk() (*string, bool)`

GetCustomproperty45LabelOk returns a tuple with the Customproperty45Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty45Label

`func (o *GetEndpoints200ResponseEndpointsInner) SetCustomproperty45Label(v string)`

SetCustomproperty45Label sets Customproperty45Label field to given value.

### HasCustomproperty45Label

`func (o *GetEndpoints200ResponseEndpointsInner) HasCustomproperty45Label() bool`

HasCustomproperty45Label returns a boolean if a field has been set.

### GetCustomproperty46Label

`func (o *GetEndpoints200ResponseEndpointsInner) GetCustomproperty46Label() string`

GetCustomproperty46Label returns the Customproperty46Label field if non-nil, zero value otherwise.

### GetCustomproperty46LabelOk

`func (o *GetEndpoints200ResponseEndpointsInner) GetCustomproperty46LabelOk() (*string, bool)`

GetCustomproperty46LabelOk returns a tuple with the Customproperty46Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty46Label

`func (o *GetEndpoints200ResponseEndpointsInner) SetCustomproperty46Label(v string)`

SetCustomproperty46Label sets Customproperty46Label field to given value.

### HasCustomproperty46Label

`func (o *GetEndpoints200ResponseEndpointsInner) HasCustomproperty46Label() bool`

HasCustomproperty46Label returns a boolean if a field has been set.

### GetCustomproperty47Label

`func (o *GetEndpoints200ResponseEndpointsInner) GetCustomproperty47Label() string`

GetCustomproperty47Label returns the Customproperty47Label field if non-nil, zero value otherwise.

### GetCustomproperty47LabelOk

`func (o *GetEndpoints200ResponseEndpointsInner) GetCustomproperty47LabelOk() (*string, bool)`

GetCustomproperty47LabelOk returns a tuple with the Customproperty47Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty47Label

`func (o *GetEndpoints200ResponseEndpointsInner) SetCustomproperty47Label(v string)`

SetCustomproperty47Label sets Customproperty47Label field to given value.

### HasCustomproperty47Label

`func (o *GetEndpoints200ResponseEndpointsInner) HasCustomproperty47Label() bool`

HasCustomproperty47Label returns a boolean if a field has been set.

### GetCustomproperty48Label

`func (o *GetEndpoints200ResponseEndpointsInner) GetCustomproperty48Label() string`

GetCustomproperty48Label returns the Customproperty48Label field if non-nil, zero value otherwise.

### GetCustomproperty48LabelOk

`func (o *GetEndpoints200ResponseEndpointsInner) GetCustomproperty48LabelOk() (*string, bool)`

GetCustomproperty48LabelOk returns a tuple with the Customproperty48Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty48Label

`func (o *GetEndpoints200ResponseEndpointsInner) SetCustomproperty48Label(v string)`

SetCustomproperty48Label sets Customproperty48Label field to given value.

### HasCustomproperty48Label

`func (o *GetEndpoints200ResponseEndpointsInner) HasCustomproperty48Label() bool`

HasCustomproperty48Label returns a boolean if a field has been set.

### GetCustomproperty49Label

`func (o *GetEndpoints200ResponseEndpointsInner) GetCustomproperty49Label() string`

GetCustomproperty49Label returns the Customproperty49Label field if non-nil, zero value otherwise.

### GetCustomproperty49LabelOk

`func (o *GetEndpoints200ResponseEndpointsInner) GetCustomproperty49LabelOk() (*string, bool)`

GetCustomproperty49LabelOk returns a tuple with the Customproperty49Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty49Label

`func (o *GetEndpoints200ResponseEndpointsInner) SetCustomproperty49Label(v string)`

SetCustomproperty49Label sets Customproperty49Label field to given value.

### HasCustomproperty49Label

`func (o *GetEndpoints200ResponseEndpointsInner) HasCustomproperty49Label() bool`

HasCustomproperty49Label returns a boolean if a field has been set.

### GetCustomproperty50Label

`func (o *GetEndpoints200ResponseEndpointsInner) GetCustomproperty50Label() string`

GetCustomproperty50Label returns the Customproperty50Label field if non-nil, zero value otherwise.

### GetCustomproperty50LabelOk

`func (o *GetEndpoints200ResponseEndpointsInner) GetCustomproperty50LabelOk() (*string, bool)`

GetCustomproperty50LabelOk returns a tuple with the Customproperty50Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty50Label

`func (o *GetEndpoints200ResponseEndpointsInner) SetCustomproperty50Label(v string)`

SetCustomproperty50Label sets Customproperty50Label field to given value.

### HasCustomproperty50Label

`func (o *GetEndpoints200ResponseEndpointsInner) HasCustomproperty50Label() bool`

HasCustomproperty50Label returns a boolean if a field has been set.

### GetCustomproperty51Label

`func (o *GetEndpoints200ResponseEndpointsInner) GetCustomproperty51Label() string`

GetCustomproperty51Label returns the Customproperty51Label field if non-nil, zero value otherwise.

### GetCustomproperty51LabelOk

`func (o *GetEndpoints200ResponseEndpointsInner) GetCustomproperty51LabelOk() (*string, bool)`

GetCustomproperty51LabelOk returns a tuple with the Customproperty51Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty51Label

`func (o *GetEndpoints200ResponseEndpointsInner) SetCustomproperty51Label(v string)`

SetCustomproperty51Label sets Customproperty51Label field to given value.

### HasCustomproperty51Label

`func (o *GetEndpoints200ResponseEndpointsInner) HasCustomproperty51Label() bool`

HasCustomproperty51Label returns a boolean if a field has been set.

### GetCustomproperty52Label

`func (o *GetEndpoints200ResponseEndpointsInner) GetCustomproperty52Label() string`

GetCustomproperty52Label returns the Customproperty52Label field if non-nil, zero value otherwise.

### GetCustomproperty52LabelOk

`func (o *GetEndpoints200ResponseEndpointsInner) GetCustomproperty52LabelOk() (*string, bool)`

GetCustomproperty52LabelOk returns a tuple with the Customproperty52Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty52Label

`func (o *GetEndpoints200ResponseEndpointsInner) SetCustomproperty52Label(v string)`

SetCustomproperty52Label sets Customproperty52Label field to given value.

### HasCustomproperty52Label

`func (o *GetEndpoints200ResponseEndpointsInner) HasCustomproperty52Label() bool`

HasCustomproperty52Label returns a boolean if a field has been set.

### GetCustomproperty53Label

`func (o *GetEndpoints200ResponseEndpointsInner) GetCustomproperty53Label() string`

GetCustomproperty53Label returns the Customproperty53Label field if non-nil, zero value otherwise.

### GetCustomproperty53LabelOk

`func (o *GetEndpoints200ResponseEndpointsInner) GetCustomproperty53LabelOk() (*string, bool)`

GetCustomproperty53LabelOk returns a tuple with the Customproperty53Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty53Label

`func (o *GetEndpoints200ResponseEndpointsInner) SetCustomproperty53Label(v string)`

SetCustomproperty53Label sets Customproperty53Label field to given value.

### HasCustomproperty53Label

`func (o *GetEndpoints200ResponseEndpointsInner) HasCustomproperty53Label() bool`

HasCustomproperty53Label returns a boolean if a field has been set.

### GetCustomproperty54Label

`func (o *GetEndpoints200ResponseEndpointsInner) GetCustomproperty54Label() string`

GetCustomproperty54Label returns the Customproperty54Label field if non-nil, zero value otherwise.

### GetCustomproperty54LabelOk

`func (o *GetEndpoints200ResponseEndpointsInner) GetCustomproperty54LabelOk() (*string, bool)`

GetCustomproperty54LabelOk returns a tuple with the Customproperty54Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty54Label

`func (o *GetEndpoints200ResponseEndpointsInner) SetCustomproperty54Label(v string)`

SetCustomproperty54Label sets Customproperty54Label field to given value.

### HasCustomproperty54Label

`func (o *GetEndpoints200ResponseEndpointsInner) HasCustomproperty54Label() bool`

HasCustomproperty54Label returns a boolean if a field has been set.

### GetCustomproperty55Label

`func (o *GetEndpoints200ResponseEndpointsInner) GetCustomproperty55Label() string`

GetCustomproperty55Label returns the Customproperty55Label field if non-nil, zero value otherwise.

### GetCustomproperty55LabelOk

`func (o *GetEndpoints200ResponseEndpointsInner) GetCustomproperty55LabelOk() (*string, bool)`

GetCustomproperty55LabelOk returns a tuple with the Customproperty55Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty55Label

`func (o *GetEndpoints200ResponseEndpointsInner) SetCustomproperty55Label(v string)`

SetCustomproperty55Label sets Customproperty55Label field to given value.

### HasCustomproperty55Label

`func (o *GetEndpoints200ResponseEndpointsInner) HasCustomproperty55Label() bool`

HasCustomproperty55Label returns a boolean if a field has been set.

### GetCustomproperty56Label

`func (o *GetEndpoints200ResponseEndpointsInner) GetCustomproperty56Label() string`

GetCustomproperty56Label returns the Customproperty56Label field if non-nil, zero value otherwise.

### GetCustomproperty56LabelOk

`func (o *GetEndpoints200ResponseEndpointsInner) GetCustomproperty56LabelOk() (*string, bool)`

GetCustomproperty56LabelOk returns a tuple with the Customproperty56Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty56Label

`func (o *GetEndpoints200ResponseEndpointsInner) SetCustomproperty56Label(v string)`

SetCustomproperty56Label sets Customproperty56Label field to given value.

### HasCustomproperty56Label

`func (o *GetEndpoints200ResponseEndpointsInner) HasCustomproperty56Label() bool`

HasCustomproperty56Label returns a boolean if a field has been set.

### GetCustomproperty57Label

`func (o *GetEndpoints200ResponseEndpointsInner) GetCustomproperty57Label() string`

GetCustomproperty57Label returns the Customproperty57Label field if non-nil, zero value otherwise.

### GetCustomproperty57LabelOk

`func (o *GetEndpoints200ResponseEndpointsInner) GetCustomproperty57LabelOk() (*string, bool)`

GetCustomproperty57LabelOk returns a tuple with the Customproperty57Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty57Label

`func (o *GetEndpoints200ResponseEndpointsInner) SetCustomproperty57Label(v string)`

SetCustomproperty57Label sets Customproperty57Label field to given value.

### HasCustomproperty57Label

`func (o *GetEndpoints200ResponseEndpointsInner) HasCustomproperty57Label() bool`

HasCustomproperty57Label returns a boolean if a field has been set.

### GetCustomproperty58Label

`func (o *GetEndpoints200ResponseEndpointsInner) GetCustomproperty58Label() string`

GetCustomproperty58Label returns the Customproperty58Label field if non-nil, zero value otherwise.

### GetCustomproperty58LabelOk

`func (o *GetEndpoints200ResponseEndpointsInner) GetCustomproperty58LabelOk() (*string, bool)`

GetCustomproperty58LabelOk returns a tuple with the Customproperty58Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty58Label

`func (o *GetEndpoints200ResponseEndpointsInner) SetCustomproperty58Label(v string)`

SetCustomproperty58Label sets Customproperty58Label field to given value.

### HasCustomproperty58Label

`func (o *GetEndpoints200ResponseEndpointsInner) HasCustomproperty58Label() bool`

HasCustomproperty58Label returns a boolean if a field has been set.

### GetCustomproperty59Label

`func (o *GetEndpoints200ResponseEndpointsInner) GetCustomproperty59Label() string`

GetCustomproperty59Label returns the Customproperty59Label field if non-nil, zero value otherwise.

### GetCustomproperty59LabelOk

`func (o *GetEndpoints200ResponseEndpointsInner) GetCustomproperty59LabelOk() (*string, bool)`

GetCustomproperty59LabelOk returns a tuple with the Customproperty59Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty59Label

`func (o *GetEndpoints200ResponseEndpointsInner) SetCustomproperty59Label(v string)`

SetCustomproperty59Label sets Customproperty59Label field to given value.

### HasCustomproperty59Label

`func (o *GetEndpoints200ResponseEndpointsInner) HasCustomproperty59Label() bool`

HasCustomproperty59Label returns a boolean if a field has been set.

### GetCustomproperty60Label

`func (o *GetEndpoints200ResponseEndpointsInner) GetCustomproperty60Label() string`

GetCustomproperty60Label returns the Customproperty60Label field if non-nil, zero value otherwise.

### GetCustomproperty60LabelOk

`func (o *GetEndpoints200ResponseEndpointsInner) GetCustomproperty60LabelOk() (*string, bool)`

GetCustomproperty60LabelOk returns a tuple with the Customproperty60Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty60Label

`func (o *GetEndpoints200ResponseEndpointsInner) SetCustomproperty60Label(v string)`

SetCustomproperty60Label sets Customproperty60Label field to given value.

### HasCustomproperty60Label

`func (o *GetEndpoints200ResponseEndpointsInner) HasCustomproperty60Label() bool`

HasCustomproperty60Label returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


