# UpdateEndpointRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Endpointname** | **string** | Specify a name for the endpoint. Provide a logical name that will help you easily identify it. | 
**PrimaryAccountType** | Pointer to **string** |  | [optional] 
**AccountTypeNoPasswordChange** | Pointer to **string** |  | [optional] 
**DisplayName** | Pointer to **string** | Enter a user-friendly display name for the endpoint that will be displayed in the user interface. Display Name can be different from Endpoint Name. | [optional] 
**Securitysystem** | Pointer to **string** | Specify the Security system for which you want to create an endpoint. | [optional] 
**Description** | Pointer to **string** | Specify a description for the endpoint. | [optional] 
**OwnerType** | Pointer to **string** | Specify the owner type for the endpoint. An endpoint can be owned by a User or Usergroup. | [optional] 
**Owner** | Pointer to **string** | Specify the owner of the endpoint. If the ownerType is User, then specify the username of the owner, and If it is is Usergroup then specify the name of the user group. | [optional] 
**ResourceOwnerType** | Pointer to **string** | Specify the resource owner type for the endpoint. An endpoint can be owned by a User or Usergroup. | [optional] 
**ResourceOwner** | Pointer to **string** |  Specify the resource owner of the endpoint. If the resourceOwnerType is User, then specify the username of the owner and If it is Usergroup, specify the name of the user group. | [optional] 
**Accessquery** | Pointer to **string** | Specify the query to filter the access and display of the endpoint to specific users. If you do not define a query, the endpoint is displayed for all users. | [optional] 
**EnableCopyAccess** | Pointer to **string** | Specify true to display the Copy Access from User option in the Request pages. | [optional] 
**CreateEntTaskforRemoveAccEnd** | Pointer to **string** | If this is set to true, remove Access tasks will be created for entitlements (account entitlements and their dependent entitlements) when a user requests for removing an account. | [optional] 
**DisableNewAccountRequestIfAccountExists** | Pointer to **string** | Specify true to disable users from requesting additional accounts on applications where they already have active accounts. | [optional] 
**DisableRemoveAccount** | Pointer to **string** | Specify true to disable users from removing their existing application accounts. | [optional] 
**DisableModifyAccount** | Pointer to **string** | Specify true to disable users from modifying their application accounts. | [optional] 
**UserAccountCorrelationRule** | Pointer to **string** |  Specify rule to map users in EIC with the accounts during account import. | [optional] 
**CreateEntTaskforRemoveAcc** | Pointer to **string** | If this is set to true, remove Access tasks will be created for entitlements (account entitlements and their dependent entitlements) when a user requests for removing an account | [optional] 
**Connectionconfig** | Pointer to **string** | Use this configuration for processing the add access tasks and remove access tasks for AD and LDAP Connectors. | [optional] 
**BlockInflightRequest** | Pointer to **string** | Specify true to prevent users from raising duplicate requests for the same applications. | [optional] 
**Outofbandaction** | Pointer to **string** | Use this parameter to determine if you need to remove the accesses which were granted outside Saviynt. | [optional] 
**RequestableRoleType** | Pointer to [**[]UpdateEndpointRequestRequestableRoleTypeInner**](UpdateEndpointRequestRequestableRoleTypeInner.md) |  | [optional] 
**Taskemailtemplates** | Pointer to [**[]CreateEndpointRequestEmailTemplateInner**](CreateEndpointRequestEmailTemplateInner.md) |  | [optional] 
**MappedEndpoints** | Pointer to [**[]UpdateEndpointRequestMappedEndpointsInner**](UpdateEndpointRequestMappedEndpointsInner.md) |  | [optional] 
**AccountNameRule** | Pointer to **string** | Specify rule to generate an account name for this endpoint while creating a new account. | [optional] 
**AllowChangePasswordSqlquery** | Pointer to **string** | SQL query to configure the accounts for which you can change passwords. | [optional] 
**Requestable** | Pointer to **string** | Is this endpoint requestable. | [optional] 
**ParentAccountPattern** | Pointer to **string** | Specify the parent and child relationship for the Active Directory endpoint. The specified value is used to filter the parent and child objects in the Request Access tile. | [optional] 
**ServiceAccountNameRule** | Pointer to **string** | Rule to generate a name for this endpoint while creating a new service account. | [optional] 
**ServiceAccountAccessQuery** | Pointer to **string** |  | [optional] 
**ChangePasswordAccessQuery** | Pointer to **string** | Specify query to restrict the access for changing the account password of the endpoint. | [optional] 
**AccountNameValidatorRegex** | Pointer to **string** | Specify the regular expression which will be used to validate the account name either generated by the rule or provided manually. | [optional] 
**StatusConfig** | Pointer to **string** | Enable the State and Status options (Enable, Disable, Lock, Unlock) that would be available to request for a user and service accounts. | [optional] 
**PluginConfigs** | Pointer to **string** | The Plugin Configuration drives the functionality of the Saviynt SmartAssist (Browserplugin). | [optional] 
**EndpointConfig** | Pointer to **string** | Option to copy data in Step 3 of the service account request will be enabled. | [optional] 
**Customproperty1** | Pointer to **string** | Custom Property 1 | [optional] 
**Customproperty2** | Pointer to **string** | Custom Property 2 | [optional] 
**Customproperty3** | Pointer to **string** | Custom Property 3 | [optional] 
**Customproperty4** | Pointer to **string** | Custom Property 4 | [optional] 
**Customproperty5** | Pointer to **string** | Custom Property 5 | [optional] 
**Customproperty6** | Pointer to **string** | Custom Property 6 | [optional] 
**Customproperty7** | Pointer to **string** | Custom Property 7 | [optional] 
**Customproperty8** | Pointer to **string** | Custom Property 8 | [optional] 
**Customproperty9** | Pointer to **string** | Custom Property 9 | [optional] 
**Customproperty10** | Pointer to **string** | Custom Property 10 | [optional] 
**Customproperty11** | Pointer to **string** | Custom Property 11 | [optional] 
**Customproperty12** | Pointer to **string** | Custom Property 12 | [optional] 
**Customproperty13** | Pointer to **string** | Custom Property 13 | [optional] 
**Customproperty14** | Pointer to **string** | Custom Property 14 | [optional] 
**Customproperty15** | Pointer to **string** | Custom Property 15 | [optional] 
**Customproperty16** | Pointer to **string** | Custom Property 16 | [optional] 
**Customproperty17** | Pointer to **string** | Custom Property 17 | [optional] 
**Customproperty18** | Pointer to **string** | Custom Property 18 | [optional] 
**Customproperty19** | Pointer to **string** | Custom Property 19 | [optional] 
**Customproperty20** | Pointer to **string** | Custom Property 20 | [optional] 
**Customproperty21** | Pointer to **string** | Custom Property 21 | [optional] 
**Customproperty22** | Pointer to **string** | Custom Property 22 | [optional] 
**Customproperty23** | Pointer to **string** | Custom Property 23 | [optional] 
**Customproperty24** | Pointer to **string** | Custom Property 24 | [optional] 
**Customproperty25** | Pointer to **string** | Custom Property 25 | [optional] 
**Customproperty26** | Pointer to **string** | Custom Property 26 | [optional] 
**Customproperty27** | Pointer to **string** | Custom Property 27 | [optional] 
**Customproperty28** | Pointer to **string** | Custom Property 28 | [optional] 
**Customproperty29** | Pointer to **string** | Custom Property 29 | [optional] 
**Customproperty30** | Pointer to **string** | Custom Property 30 | [optional] 
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
**Customproperty1Label** | Pointer to **string** | Label for the custom property 1 of accounts of this endpoint | [optional] 
**Customproperty2Label** | Pointer to **string** | Label for the custom property 2 of accounts of this endpoint | [optional] 
**Customproperty3Label** | Pointer to **string** | Label for the custom property 3 of accounts of this endpoint | [optional] 
**Customproperty4Label** | Pointer to **string** | Label for the custom property 4 of accounts of this endpoint | [optional] 
**Customproperty5Label** | Pointer to **string** | Label for the custom property 5 of accounts of this endpoint | [optional] 
**Customproperty6Label** | Pointer to **string** | Label for the custom property 6 of accounts of this endpoint | [optional] 
**Customproperty7Label** | Pointer to **string** | Label for the custom property 7 of accounts of this endpoint | [optional] 
**Customproperty8Label** | Pointer to **string** | Label for the custom property 8 of accounts of this endpoint | [optional] 
**Customproperty9Label** | Pointer to **string** | Label for the custom property 9 of accounts of this endpoint | [optional] 
**Customproperty10Label** | Pointer to **string** | Label for the custom property 10 of accounts of this endpoint | [optional] 
**Customproperty11Label** | Pointer to **string** | Label for the custom property 11 of accounts of this endpoint | [optional] 
**Customproperty12Label** | Pointer to **string** | Label for the custom property 12 of accounts of this endpoint | [optional] 
**Customproperty13Label** | Pointer to **string** | Label for the custom property 13 of accounts of this endpoint | [optional] 
**Customproperty14Label** | Pointer to **string** | Label for the custom property 14 of accounts of this endpoint | [optional] 
**Customproperty15Label** | Pointer to **string** | Label for the custom property 15 of accounts of this endpoint | [optional] 
**Customproperty16Label** | Pointer to **string** | Label for the custom property 16 of accounts of this endpoint | [optional] 
**Customproperty17Label** | Pointer to **string** | Label for the custom property 17 of accounts of this endpoint | [optional] 
**Customproperty18Label** | Pointer to **string** | Label for the custom property 18 of accounts of this endpoint | [optional] 
**Customproperty19Label** | Pointer to **string** | Label for the custom property 19 of accounts of this endpoint | [optional] 
**Customproperty20Label** | Pointer to **string** | Label for the custom property 20 of accounts of this endpoint | [optional] 
**Customproperty21Label** | Pointer to **string** | Label for the custom property 21 of accounts of this endpoint | [optional] 
**Customproperty22Label** | Pointer to **string** | Label for the custom property 22 of accounts of this endpoint | [optional] 
**Customproperty23Label** | Pointer to **string** | Label for the custom property 23 of accounts of this endpoint | [optional] 
**Customproperty24Label** | Pointer to **string** | Label for the custom property 24 of accounts of this endpoint | [optional] 
**Customproperty25Label** | Pointer to **string** | Label for the custom property 25 of accounts of this endpoint | [optional] 
**Customproperty26Label** | Pointer to **string** | Label for the custom property 26 of accounts of this endpoint | [optional] 
**Customproperty27Label** | Pointer to **string** | Label for the custom property 27 of accounts of this endpoint | [optional] 
**Customproperty28Label** | Pointer to **string** | Label for the custom property 28 of accounts of this endpoint | [optional] 
**Customproperty29Label** | Pointer to **string** | Label for the custom property 29 of accounts of this endpoint | [optional] 
**Customproperty30Label** | Pointer to **string** | Label for the custom property 30 of accounts of this endpoint | [optional] 
**Customproperty31Label** | Pointer to **string** | Label for the custom property 31 of accounts of this endpoint | [optional] 
**Customproperty32Label** | Pointer to **string** | Label for the custom property 32 of accounts of this endpoint | [optional] 
**Customproperty33Label** | Pointer to **string** | Label for the custom property 33 of accounts of this endpoint | [optional] 
**Customproperty34Label** | Pointer to **string** | Label for the custom property 34 of accounts of this endpoint | [optional] 
**Customproperty35Label** | Pointer to **string** | Label for the custom property 35 of accounts of this endpoint | [optional] 
**Customproperty36Label** | Pointer to **string** | Label for the custom property 36 of accounts of this endpoint | [optional] 
**Customproperty37Label** | Pointer to **string** | Label for the custom property 37 of accounts of this endpoint | [optional] 
**Customproperty38Label** | Pointer to **string** | Label for the custom property 38 of accounts of this endpoint | [optional] 
**Customproperty39Label** | Pointer to **string** | Label for the custom property 39 of accounts of this endpoint | [optional] 
**Customproperty40Label** | Pointer to **string** | Label for the custom property 40 of accounts of this endpoint | [optional] 
**Customproperty41Label** | Pointer to **string** | Label for the custom property 41 of accounts of this endpoint | [optional] 
**Customproperty42Label** | Pointer to **string** | Label for the custom property 42 of accounts of this endpoint | [optional] 
**Customproperty43Label** | Pointer to **string** | Label for the custom property 43 of accounts of this endpoint | [optional] 
**Customproperty44Label** | Pointer to **string** | Label for the custom property 44 of accounts of this endpoint | [optional] 
**Customproperty45Label** | Pointer to **string** | Label for the custom property 45 of accounts of this endpoint | [optional] 
**Customproperty46Label** | Pointer to **string** | Label for the custom property 46 of accounts of this endpoint | [optional] 
**Customproperty47Label** | Pointer to **string** | Label for the custom property 47 of accounts of this endpoint | [optional] 
**Customproperty48Label** | Pointer to **string** | Label for the custom property 48 of accounts of this endpoint | [optional] 
**Customproperty49Label** | Pointer to **string** | Label for the custom property 49 of accounts of this endpoint | [optional] 
**Customproperty50Label** | Pointer to **string** | Label for the custom property 50 of accounts of this endpoint | [optional] 
**Customproperty51Label** | Pointer to **string** | Label for the custom property 51 of accounts of this endpoint | [optional] 
**Customproperty52Label** | Pointer to **string** | Label for the custom property 52 of accounts of this endpoint | [optional] 
**Customproperty53Label** | Pointer to **string** | Label for the custom property 53 of accounts of this endpoint | [optional] 
**Customproperty54Label** | Pointer to **string** | Label for the custom property 54 of accounts of this endpoint | [optional] 
**Customproperty55Label** | Pointer to **string** | Label for the custom property 55 of accounts of this endpoint | [optional] 
**Customproperty56Label** | Pointer to **string** | Label for the custom property 56 of accounts of this endpoint | [optional] 
**Customproperty57Label** | Pointer to **string** | Label for the custom property 57 of accounts of this endpoint | [optional] 
**Customproperty58Label** | Pointer to **string** | Label for the custom property 58 of accounts of this endpoint | [optional] 
**Customproperty59Label** | Pointer to **string** | Label for the custom property 59 of accounts of this endpoint | [optional] 
**Customproperty60Label** | Pointer to **string** | Label for the custom property 60 of accounts of this endpoint | [optional] 
**AllowRemoveAllRoleOnRequest** | Pointer to **string** | Specify true to displays the Remove All Roles option in the Request page that can be used to remove all the roles. | [optional] 

## Methods

### NewUpdateEndpointRequest

`func NewUpdateEndpointRequest(endpointname string, ) *UpdateEndpointRequest`

NewUpdateEndpointRequest instantiates a new UpdateEndpointRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateEndpointRequestWithDefaults

`func NewUpdateEndpointRequestWithDefaults() *UpdateEndpointRequest`

NewUpdateEndpointRequestWithDefaults instantiates a new UpdateEndpointRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEndpointname

`func (o *UpdateEndpointRequest) GetEndpointname() string`

GetEndpointname returns the Endpointname field if non-nil, zero value otherwise.

### GetEndpointnameOk

`func (o *UpdateEndpointRequest) GetEndpointnameOk() (*string, bool)`

GetEndpointnameOk returns a tuple with the Endpointname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpointname

`func (o *UpdateEndpointRequest) SetEndpointname(v string)`

SetEndpointname sets Endpointname field to given value.


### GetPrimaryAccountType

`func (o *UpdateEndpointRequest) GetPrimaryAccountType() string`

GetPrimaryAccountType returns the PrimaryAccountType field if non-nil, zero value otherwise.

### GetPrimaryAccountTypeOk

`func (o *UpdateEndpointRequest) GetPrimaryAccountTypeOk() (*string, bool)`

GetPrimaryAccountTypeOk returns a tuple with the PrimaryAccountType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryAccountType

`func (o *UpdateEndpointRequest) SetPrimaryAccountType(v string)`

SetPrimaryAccountType sets PrimaryAccountType field to given value.

### HasPrimaryAccountType

`func (o *UpdateEndpointRequest) HasPrimaryAccountType() bool`

HasPrimaryAccountType returns a boolean if a field has been set.

### GetAccountTypeNoPasswordChange

`func (o *UpdateEndpointRequest) GetAccountTypeNoPasswordChange() string`

GetAccountTypeNoPasswordChange returns the AccountTypeNoPasswordChange field if non-nil, zero value otherwise.

### GetAccountTypeNoPasswordChangeOk

`func (o *UpdateEndpointRequest) GetAccountTypeNoPasswordChangeOk() (*string, bool)`

GetAccountTypeNoPasswordChangeOk returns a tuple with the AccountTypeNoPasswordChange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountTypeNoPasswordChange

`func (o *UpdateEndpointRequest) SetAccountTypeNoPasswordChange(v string)`

SetAccountTypeNoPasswordChange sets AccountTypeNoPasswordChange field to given value.

### HasAccountTypeNoPasswordChange

`func (o *UpdateEndpointRequest) HasAccountTypeNoPasswordChange() bool`

HasAccountTypeNoPasswordChange returns a boolean if a field has been set.

### GetDisplayName

`func (o *UpdateEndpointRequest) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *UpdateEndpointRequest) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *UpdateEndpointRequest) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *UpdateEndpointRequest) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetSecuritysystem

`func (o *UpdateEndpointRequest) GetSecuritysystem() string`

GetSecuritysystem returns the Securitysystem field if non-nil, zero value otherwise.

### GetSecuritysystemOk

`func (o *UpdateEndpointRequest) GetSecuritysystemOk() (*string, bool)`

GetSecuritysystemOk returns a tuple with the Securitysystem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecuritysystem

`func (o *UpdateEndpointRequest) SetSecuritysystem(v string)`

SetSecuritysystem sets Securitysystem field to given value.

### HasSecuritysystem

`func (o *UpdateEndpointRequest) HasSecuritysystem() bool`

HasSecuritysystem returns a boolean if a field has been set.

### GetDescription

`func (o *UpdateEndpointRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateEndpointRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateEndpointRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateEndpointRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetOwnerType

`func (o *UpdateEndpointRequest) GetOwnerType() string`

GetOwnerType returns the OwnerType field if non-nil, zero value otherwise.

### GetOwnerTypeOk

`func (o *UpdateEndpointRequest) GetOwnerTypeOk() (*string, bool)`

GetOwnerTypeOk returns a tuple with the OwnerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerType

`func (o *UpdateEndpointRequest) SetOwnerType(v string)`

SetOwnerType sets OwnerType field to given value.

### HasOwnerType

`func (o *UpdateEndpointRequest) HasOwnerType() bool`

HasOwnerType returns a boolean if a field has been set.

### GetOwner

`func (o *UpdateEndpointRequest) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *UpdateEndpointRequest) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *UpdateEndpointRequest) SetOwner(v string)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *UpdateEndpointRequest) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetResourceOwnerType

`func (o *UpdateEndpointRequest) GetResourceOwnerType() string`

GetResourceOwnerType returns the ResourceOwnerType field if non-nil, zero value otherwise.

### GetResourceOwnerTypeOk

`func (o *UpdateEndpointRequest) GetResourceOwnerTypeOk() (*string, bool)`

GetResourceOwnerTypeOk returns a tuple with the ResourceOwnerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceOwnerType

`func (o *UpdateEndpointRequest) SetResourceOwnerType(v string)`

SetResourceOwnerType sets ResourceOwnerType field to given value.

### HasResourceOwnerType

`func (o *UpdateEndpointRequest) HasResourceOwnerType() bool`

HasResourceOwnerType returns a boolean if a field has been set.

### GetResourceOwner

`func (o *UpdateEndpointRequest) GetResourceOwner() string`

GetResourceOwner returns the ResourceOwner field if non-nil, zero value otherwise.

### GetResourceOwnerOk

`func (o *UpdateEndpointRequest) GetResourceOwnerOk() (*string, bool)`

GetResourceOwnerOk returns a tuple with the ResourceOwner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceOwner

`func (o *UpdateEndpointRequest) SetResourceOwner(v string)`

SetResourceOwner sets ResourceOwner field to given value.

### HasResourceOwner

`func (o *UpdateEndpointRequest) HasResourceOwner() bool`

HasResourceOwner returns a boolean if a field has been set.

### GetAccessquery

`func (o *UpdateEndpointRequest) GetAccessquery() string`

GetAccessquery returns the Accessquery field if non-nil, zero value otherwise.

### GetAccessqueryOk

`func (o *UpdateEndpointRequest) GetAccessqueryOk() (*string, bool)`

GetAccessqueryOk returns a tuple with the Accessquery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessquery

`func (o *UpdateEndpointRequest) SetAccessquery(v string)`

SetAccessquery sets Accessquery field to given value.

### HasAccessquery

`func (o *UpdateEndpointRequest) HasAccessquery() bool`

HasAccessquery returns a boolean if a field has been set.

### GetEnableCopyAccess

`func (o *UpdateEndpointRequest) GetEnableCopyAccess() string`

GetEnableCopyAccess returns the EnableCopyAccess field if non-nil, zero value otherwise.

### GetEnableCopyAccessOk

`func (o *UpdateEndpointRequest) GetEnableCopyAccessOk() (*string, bool)`

GetEnableCopyAccessOk returns a tuple with the EnableCopyAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableCopyAccess

`func (o *UpdateEndpointRequest) SetEnableCopyAccess(v string)`

SetEnableCopyAccess sets EnableCopyAccess field to given value.

### HasEnableCopyAccess

`func (o *UpdateEndpointRequest) HasEnableCopyAccess() bool`

HasEnableCopyAccess returns a boolean if a field has been set.

### GetCreateEntTaskforRemoveAccEnd

`func (o *UpdateEndpointRequest) GetCreateEntTaskforRemoveAccEnd() string`

GetCreateEntTaskforRemoveAccEnd returns the CreateEntTaskforRemoveAccEnd field if non-nil, zero value otherwise.

### GetCreateEntTaskforRemoveAccEndOk

`func (o *UpdateEndpointRequest) GetCreateEntTaskforRemoveAccEndOk() (*string, bool)`

GetCreateEntTaskforRemoveAccEndOk returns a tuple with the CreateEntTaskforRemoveAccEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateEntTaskforRemoveAccEnd

`func (o *UpdateEndpointRequest) SetCreateEntTaskforRemoveAccEnd(v string)`

SetCreateEntTaskforRemoveAccEnd sets CreateEntTaskforRemoveAccEnd field to given value.

### HasCreateEntTaskforRemoveAccEnd

`func (o *UpdateEndpointRequest) HasCreateEntTaskforRemoveAccEnd() bool`

HasCreateEntTaskforRemoveAccEnd returns a boolean if a field has been set.

### GetDisableNewAccountRequestIfAccountExists

`func (o *UpdateEndpointRequest) GetDisableNewAccountRequestIfAccountExists() string`

GetDisableNewAccountRequestIfAccountExists returns the DisableNewAccountRequestIfAccountExists field if non-nil, zero value otherwise.

### GetDisableNewAccountRequestIfAccountExistsOk

`func (o *UpdateEndpointRequest) GetDisableNewAccountRequestIfAccountExistsOk() (*string, bool)`

GetDisableNewAccountRequestIfAccountExistsOk returns a tuple with the DisableNewAccountRequestIfAccountExists field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableNewAccountRequestIfAccountExists

`func (o *UpdateEndpointRequest) SetDisableNewAccountRequestIfAccountExists(v string)`

SetDisableNewAccountRequestIfAccountExists sets DisableNewAccountRequestIfAccountExists field to given value.

### HasDisableNewAccountRequestIfAccountExists

`func (o *UpdateEndpointRequest) HasDisableNewAccountRequestIfAccountExists() bool`

HasDisableNewAccountRequestIfAccountExists returns a boolean if a field has been set.

### GetDisableRemoveAccount

`func (o *UpdateEndpointRequest) GetDisableRemoveAccount() string`

GetDisableRemoveAccount returns the DisableRemoveAccount field if non-nil, zero value otherwise.

### GetDisableRemoveAccountOk

`func (o *UpdateEndpointRequest) GetDisableRemoveAccountOk() (*string, bool)`

GetDisableRemoveAccountOk returns a tuple with the DisableRemoveAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableRemoveAccount

`func (o *UpdateEndpointRequest) SetDisableRemoveAccount(v string)`

SetDisableRemoveAccount sets DisableRemoveAccount field to given value.

### HasDisableRemoveAccount

`func (o *UpdateEndpointRequest) HasDisableRemoveAccount() bool`

HasDisableRemoveAccount returns a boolean if a field has been set.

### GetDisableModifyAccount

`func (o *UpdateEndpointRequest) GetDisableModifyAccount() string`

GetDisableModifyAccount returns the DisableModifyAccount field if non-nil, zero value otherwise.

### GetDisableModifyAccountOk

`func (o *UpdateEndpointRequest) GetDisableModifyAccountOk() (*string, bool)`

GetDisableModifyAccountOk returns a tuple with the DisableModifyAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableModifyAccount

`func (o *UpdateEndpointRequest) SetDisableModifyAccount(v string)`

SetDisableModifyAccount sets DisableModifyAccount field to given value.

### HasDisableModifyAccount

`func (o *UpdateEndpointRequest) HasDisableModifyAccount() bool`

HasDisableModifyAccount returns a boolean if a field has been set.

### GetUserAccountCorrelationRule

`func (o *UpdateEndpointRequest) GetUserAccountCorrelationRule() string`

GetUserAccountCorrelationRule returns the UserAccountCorrelationRule field if non-nil, zero value otherwise.

### GetUserAccountCorrelationRuleOk

`func (o *UpdateEndpointRequest) GetUserAccountCorrelationRuleOk() (*string, bool)`

GetUserAccountCorrelationRuleOk returns a tuple with the UserAccountCorrelationRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAccountCorrelationRule

`func (o *UpdateEndpointRequest) SetUserAccountCorrelationRule(v string)`

SetUserAccountCorrelationRule sets UserAccountCorrelationRule field to given value.

### HasUserAccountCorrelationRule

`func (o *UpdateEndpointRequest) HasUserAccountCorrelationRule() bool`

HasUserAccountCorrelationRule returns a boolean if a field has been set.

### GetCreateEntTaskforRemoveAcc

`func (o *UpdateEndpointRequest) GetCreateEntTaskforRemoveAcc() string`

GetCreateEntTaskforRemoveAcc returns the CreateEntTaskforRemoveAcc field if non-nil, zero value otherwise.

### GetCreateEntTaskforRemoveAccOk

`func (o *UpdateEndpointRequest) GetCreateEntTaskforRemoveAccOk() (*string, bool)`

GetCreateEntTaskforRemoveAccOk returns a tuple with the CreateEntTaskforRemoveAcc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateEntTaskforRemoveAcc

`func (o *UpdateEndpointRequest) SetCreateEntTaskforRemoveAcc(v string)`

SetCreateEntTaskforRemoveAcc sets CreateEntTaskforRemoveAcc field to given value.

### HasCreateEntTaskforRemoveAcc

`func (o *UpdateEndpointRequest) HasCreateEntTaskforRemoveAcc() bool`

HasCreateEntTaskforRemoveAcc returns a boolean if a field has been set.

### GetConnectionconfig

`func (o *UpdateEndpointRequest) GetConnectionconfig() string`

GetConnectionconfig returns the Connectionconfig field if non-nil, zero value otherwise.

### GetConnectionconfigOk

`func (o *UpdateEndpointRequest) GetConnectionconfigOk() (*string, bool)`

GetConnectionconfigOk returns a tuple with the Connectionconfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionconfig

`func (o *UpdateEndpointRequest) SetConnectionconfig(v string)`

SetConnectionconfig sets Connectionconfig field to given value.

### HasConnectionconfig

`func (o *UpdateEndpointRequest) HasConnectionconfig() bool`

HasConnectionconfig returns a boolean if a field has been set.

### GetBlockInflightRequest

`func (o *UpdateEndpointRequest) GetBlockInflightRequest() string`

GetBlockInflightRequest returns the BlockInflightRequest field if non-nil, zero value otherwise.

### GetBlockInflightRequestOk

`func (o *UpdateEndpointRequest) GetBlockInflightRequestOk() (*string, bool)`

GetBlockInflightRequestOk returns a tuple with the BlockInflightRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockInflightRequest

`func (o *UpdateEndpointRequest) SetBlockInflightRequest(v string)`

SetBlockInflightRequest sets BlockInflightRequest field to given value.

### HasBlockInflightRequest

`func (o *UpdateEndpointRequest) HasBlockInflightRequest() bool`

HasBlockInflightRequest returns a boolean if a field has been set.

### GetOutofbandaction

`func (o *UpdateEndpointRequest) GetOutofbandaction() string`

GetOutofbandaction returns the Outofbandaction field if non-nil, zero value otherwise.

### GetOutofbandactionOk

`func (o *UpdateEndpointRequest) GetOutofbandactionOk() (*string, bool)`

GetOutofbandactionOk returns a tuple with the Outofbandaction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutofbandaction

`func (o *UpdateEndpointRequest) SetOutofbandaction(v string)`

SetOutofbandaction sets Outofbandaction field to given value.

### HasOutofbandaction

`func (o *UpdateEndpointRequest) HasOutofbandaction() bool`

HasOutofbandaction returns a boolean if a field has been set.

### GetRequestableRoleType

`func (o *UpdateEndpointRequest) GetRequestableRoleType() []UpdateEndpointRequestRequestableRoleTypeInner`

GetRequestableRoleType returns the RequestableRoleType field if non-nil, zero value otherwise.

### GetRequestableRoleTypeOk

`func (o *UpdateEndpointRequest) GetRequestableRoleTypeOk() (*[]UpdateEndpointRequestRequestableRoleTypeInner, bool)`

GetRequestableRoleTypeOk returns a tuple with the RequestableRoleType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestableRoleType

`func (o *UpdateEndpointRequest) SetRequestableRoleType(v []UpdateEndpointRequestRequestableRoleTypeInner)`

SetRequestableRoleType sets RequestableRoleType field to given value.

### HasRequestableRoleType

`func (o *UpdateEndpointRequest) HasRequestableRoleType() bool`

HasRequestableRoleType returns a boolean if a field has been set.

### GetTaskemailtemplates

`func (o *UpdateEndpointRequest) GetTaskemailtemplates() []CreateEndpointRequestEmailTemplateInner`

GetTaskemailtemplates returns the Taskemailtemplates field if non-nil, zero value otherwise.

### GetTaskemailtemplatesOk

`func (o *UpdateEndpointRequest) GetTaskemailtemplatesOk() (*[]CreateEndpointRequestEmailTemplateInner, bool)`

GetTaskemailtemplatesOk returns a tuple with the Taskemailtemplates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskemailtemplates

`func (o *UpdateEndpointRequest) SetTaskemailtemplates(v []CreateEndpointRequestEmailTemplateInner)`

SetTaskemailtemplates sets Taskemailtemplates field to given value.

### HasTaskemailtemplates

`func (o *UpdateEndpointRequest) HasTaskemailtemplates() bool`

HasTaskemailtemplates returns a boolean if a field has been set.

### GetMappedEndpoints

`func (o *UpdateEndpointRequest) GetMappedEndpoints() []UpdateEndpointRequestMappedEndpointsInner`

GetMappedEndpoints returns the MappedEndpoints field if non-nil, zero value otherwise.

### GetMappedEndpointsOk

`func (o *UpdateEndpointRequest) GetMappedEndpointsOk() (*[]UpdateEndpointRequestMappedEndpointsInner, bool)`

GetMappedEndpointsOk returns a tuple with the MappedEndpoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMappedEndpoints

`func (o *UpdateEndpointRequest) SetMappedEndpoints(v []UpdateEndpointRequestMappedEndpointsInner)`

SetMappedEndpoints sets MappedEndpoints field to given value.

### HasMappedEndpoints

`func (o *UpdateEndpointRequest) HasMappedEndpoints() bool`

HasMappedEndpoints returns a boolean if a field has been set.

### GetAccountNameRule

`func (o *UpdateEndpointRequest) GetAccountNameRule() string`

GetAccountNameRule returns the AccountNameRule field if non-nil, zero value otherwise.

### GetAccountNameRuleOk

`func (o *UpdateEndpointRequest) GetAccountNameRuleOk() (*string, bool)`

GetAccountNameRuleOk returns a tuple with the AccountNameRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountNameRule

`func (o *UpdateEndpointRequest) SetAccountNameRule(v string)`

SetAccountNameRule sets AccountNameRule field to given value.

### HasAccountNameRule

`func (o *UpdateEndpointRequest) HasAccountNameRule() bool`

HasAccountNameRule returns a boolean if a field has been set.

### GetAllowChangePasswordSqlquery

`func (o *UpdateEndpointRequest) GetAllowChangePasswordSqlquery() string`

GetAllowChangePasswordSqlquery returns the AllowChangePasswordSqlquery field if non-nil, zero value otherwise.

### GetAllowChangePasswordSqlqueryOk

`func (o *UpdateEndpointRequest) GetAllowChangePasswordSqlqueryOk() (*string, bool)`

GetAllowChangePasswordSqlqueryOk returns a tuple with the AllowChangePasswordSqlquery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowChangePasswordSqlquery

`func (o *UpdateEndpointRequest) SetAllowChangePasswordSqlquery(v string)`

SetAllowChangePasswordSqlquery sets AllowChangePasswordSqlquery field to given value.

### HasAllowChangePasswordSqlquery

`func (o *UpdateEndpointRequest) HasAllowChangePasswordSqlquery() bool`

HasAllowChangePasswordSqlquery returns a boolean if a field has been set.

### GetRequestable

`func (o *UpdateEndpointRequest) GetRequestable() string`

GetRequestable returns the Requestable field if non-nil, zero value otherwise.

### GetRequestableOk

`func (o *UpdateEndpointRequest) GetRequestableOk() (*string, bool)`

GetRequestableOk returns a tuple with the Requestable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestable

`func (o *UpdateEndpointRequest) SetRequestable(v string)`

SetRequestable sets Requestable field to given value.

### HasRequestable

`func (o *UpdateEndpointRequest) HasRequestable() bool`

HasRequestable returns a boolean if a field has been set.

### GetParentAccountPattern

`func (o *UpdateEndpointRequest) GetParentAccountPattern() string`

GetParentAccountPattern returns the ParentAccountPattern field if non-nil, zero value otherwise.

### GetParentAccountPatternOk

`func (o *UpdateEndpointRequest) GetParentAccountPatternOk() (*string, bool)`

GetParentAccountPatternOk returns a tuple with the ParentAccountPattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentAccountPattern

`func (o *UpdateEndpointRequest) SetParentAccountPattern(v string)`

SetParentAccountPattern sets ParentAccountPattern field to given value.

### HasParentAccountPattern

`func (o *UpdateEndpointRequest) HasParentAccountPattern() bool`

HasParentAccountPattern returns a boolean if a field has been set.

### GetServiceAccountNameRule

`func (o *UpdateEndpointRequest) GetServiceAccountNameRule() string`

GetServiceAccountNameRule returns the ServiceAccountNameRule field if non-nil, zero value otherwise.

### GetServiceAccountNameRuleOk

`func (o *UpdateEndpointRequest) GetServiceAccountNameRuleOk() (*string, bool)`

GetServiceAccountNameRuleOk returns a tuple with the ServiceAccountNameRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAccountNameRule

`func (o *UpdateEndpointRequest) SetServiceAccountNameRule(v string)`

SetServiceAccountNameRule sets ServiceAccountNameRule field to given value.

### HasServiceAccountNameRule

`func (o *UpdateEndpointRequest) HasServiceAccountNameRule() bool`

HasServiceAccountNameRule returns a boolean if a field has been set.

### GetServiceAccountAccessQuery

`func (o *UpdateEndpointRequest) GetServiceAccountAccessQuery() string`

GetServiceAccountAccessQuery returns the ServiceAccountAccessQuery field if non-nil, zero value otherwise.

### GetServiceAccountAccessQueryOk

`func (o *UpdateEndpointRequest) GetServiceAccountAccessQueryOk() (*string, bool)`

GetServiceAccountAccessQueryOk returns a tuple with the ServiceAccountAccessQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAccountAccessQuery

`func (o *UpdateEndpointRequest) SetServiceAccountAccessQuery(v string)`

SetServiceAccountAccessQuery sets ServiceAccountAccessQuery field to given value.

### HasServiceAccountAccessQuery

`func (o *UpdateEndpointRequest) HasServiceAccountAccessQuery() bool`

HasServiceAccountAccessQuery returns a boolean if a field has been set.

### GetChangePasswordAccessQuery

`func (o *UpdateEndpointRequest) GetChangePasswordAccessQuery() string`

GetChangePasswordAccessQuery returns the ChangePasswordAccessQuery field if non-nil, zero value otherwise.

### GetChangePasswordAccessQueryOk

`func (o *UpdateEndpointRequest) GetChangePasswordAccessQueryOk() (*string, bool)`

GetChangePasswordAccessQueryOk returns a tuple with the ChangePasswordAccessQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangePasswordAccessQuery

`func (o *UpdateEndpointRequest) SetChangePasswordAccessQuery(v string)`

SetChangePasswordAccessQuery sets ChangePasswordAccessQuery field to given value.

### HasChangePasswordAccessQuery

`func (o *UpdateEndpointRequest) HasChangePasswordAccessQuery() bool`

HasChangePasswordAccessQuery returns a boolean if a field has been set.

### GetAccountNameValidatorRegex

`func (o *UpdateEndpointRequest) GetAccountNameValidatorRegex() string`

GetAccountNameValidatorRegex returns the AccountNameValidatorRegex field if non-nil, zero value otherwise.

### GetAccountNameValidatorRegexOk

`func (o *UpdateEndpointRequest) GetAccountNameValidatorRegexOk() (*string, bool)`

GetAccountNameValidatorRegexOk returns a tuple with the AccountNameValidatorRegex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountNameValidatorRegex

`func (o *UpdateEndpointRequest) SetAccountNameValidatorRegex(v string)`

SetAccountNameValidatorRegex sets AccountNameValidatorRegex field to given value.

### HasAccountNameValidatorRegex

`func (o *UpdateEndpointRequest) HasAccountNameValidatorRegex() bool`

HasAccountNameValidatorRegex returns a boolean if a field has been set.

### GetStatusConfig

`func (o *UpdateEndpointRequest) GetStatusConfig() string`

GetStatusConfig returns the StatusConfig field if non-nil, zero value otherwise.

### GetStatusConfigOk

`func (o *UpdateEndpointRequest) GetStatusConfigOk() (*string, bool)`

GetStatusConfigOk returns a tuple with the StatusConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusConfig

`func (o *UpdateEndpointRequest) SetStatusConfig(v string)`

SetStatusConfig sets StatusConfig field to given value.

### HasStatusConfig

`func (o *UpdateEndpointRequest) HasStatusConfig() bool`

HasStatusConfig returns a boolean if a field has been set.

### GetPluginConfigs

`func (o *UpdateEndpointRequest) GetPluginConfigs() string`

GetPluginConfigs returns the PluginConfigs field if non-nil, zero value otherwise.

### GetPluginConfigsOk

`func (o *UpdateEndpointRequest) GetPluginConfigsOk() (*string, bool)`

GetPluginConfigsOk returns a tuple with the PluginConfigs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPluginConfigs

`func (o *UpdateEndpointRequest) SetPluginConfigs(v string)`

SetPluginConfigs sets PluginConfigs field to given value.

### HasPluginConfigs

`func (o *UpdateEndpointRequest) HasPluginConfigs() bool`

HasPluginConfigs returns a boolean if a field has been set.

### GetEndpointConfig

`func (o *UpdateEndpointRequest) GetEndpointConfig() string`

GetEndpointConfig returns the EndpointConfig field if non-nil, zero value otherwise.

### GetEndpointConfigOk

`func (o *UpdateEndpointRequest) GetEndpointConfigOk() (*string, bool)`

GetEndpointConfigOk returns a tuple with the EndpointConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpointConfig

`func (o *UpdateEndpointRequest) SetEndpointConfig(v string)`

SetEndpointConfig sets EndpointConfig field to given value.

### HasEndpointConfig

`func (o *UpdateEndpointRequest) HasEndpointConfig() bool`

HasEndpointConfig returns a boolean if a field has been set.

### GetCustomproperty1

`func (o *UpdateEndpointRequest) GetCustomproperty1() string`

GetCustomproperty1 returns the Customproperty1 field if non-nil, zero value otherwise.

### GetCustomproperty1Ok

`func (o *UpdateEndpointRequest) GetCustomproperty1Ok() (*string, bool)`

GetCustomproperty1Ok returns a tuple with the Customproperty1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty1

`func (o *UpdateEndpointRequest) SetCustomproperty1(v string)`

SetCustomproperty1 sets Customproperty1 field to given value.

### HasCustomproperty1

`func (o *UpdateEndpointRequest) HasCustomproperty1() bool`

HasCustomproperty1 returns a boolean if a field has been set.

### GetCustomproperty2

`func (o *UpdateEndpointRequest) GetCustomproperty2() string`

GetCustomproperty2 returns the Customproperty2 field if non-nil, zero value otherwise.

### GetCustomproperty2Ok

`func (o *UpdateEndpointRequest) GetCustomproperty2Ok() (*string, bool)`

GetCustomproperty2Ok returns a tuple with the Customproperty2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty2

`func (o *UpdateEndpointRequest) SetCustomproperty2(v string)`

SetCustomproperty2 sets Customproperty2 field to given value.

### HasCustomproperty2

`func (o *UpdateEndpointRequest) HasCustomproperty2() bool`

HasCustomproperty2 returns a boolean if a field has been set.

### GetCustomproperty3

`func (o *UpdateEndpointRequest) GetCustomproperty3() string`

GetCustomproperty3 returns the Customproperty3 field if non-nil, zero value otherwise.

### GetCustomproperty3Ok

`func (o *UpdateEndpointRequest) GetCustomproperty3Ok() (*string, bool)`

GetCustomproperty3Ok returns a tuple with the Customproperty3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty3

`func (o *UpdateEndpointRequest) SetCustomproperty3(v string)`

SetCustomproperty3 sets Customproperty3 field to given value.

### HasCustomproperty3

`func (o *UpdateEndpointRequest) HasCustomproperty3() bool`

HasCustomproperty3 returns a boolean if a field has been set.

### GetCustomproperty4

`func (o *UpdateEndpointRequest) GetCustomproperty4() string`

GetCustomproperty4 returns the Customproperty4 field if non-nil, zero value otherwise.

### GetCustomproperty4Ok

`func (o *UpdateEndpointRequest) GetCustomproperty4Ok() (*string, bool)`

GetCustomproperty4Ok returns a tuple with the Customproperty4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty4

`func (o *UpdateEndpointRequest) SetCustomproperty4(v string)`

SetCustomproperty4 sets Customproperty4 field to given value.

### HasCustomproperty4

`func (o *UpdateEndpointRequest) HasCustomproperty4() bool`

HasCustomproperty4 returns a boolean if a field has been set.

### GetCustomproperty5

`func (o *UpdateEndpointRequest) GetCustomproperty5() string`

GetCustomproperty5 returns the Customproperty5 field if non-nil, zero value otherwise.

### GetCustomproperty5Ok

`func (o *UpdateEndpointRequest) GetCustomproperty5Ok() (*string, bool)`

GetCustomproperty5Ok returns a tuple with the Customproperty5 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty5

`func (o *UpdateEndpointRequest) SetCustomproperty5(v string)`

SetCustomproperty5 sets Customproperty5 field to given value.

### HasCustomproperty5

`func (o *UpdateEndpointRequest) HasCustomproperty5() bool`

HasCustomproperty5 returns a boolean if a field has been set.

### GetCustomproperty6

`func (o *UpdateEndpointRequest) GetCustomproperty6() string`

GetCustomproperty6 returns the Customproperty6 field if non-nil, zero value otherwise.

### GetCustomproperty6Ok

`func (o *UpdateEndpointRequest) GetCustomproperty6Ok() (*string, bool)`

GetCustomproperty6Ok returns a tuple with the Customproperty6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty6

`func (o *UpdateEndpointRequest) SetCustomproperty6(v string)`

SetCustomproperty6 sets Customproperty6 field to given value.

### HasCustomproperty6

`func (o *UpdateEndpointRequest) HasCustomproperty6() bool`

HasCustomproperty6 returns a boolean if a field has been set.

### GetCustomproperty7

`func (o *UpdateEndpointRequest) GetCustomproperty7() string`

GetCustomproperty7 returns the Customproperty7 field if non-nil, zero value otherwise.

### GetCustomproperty7Ok

`func (o *UpdateEndpointRequest) GetCustomproperty7Ok() (*string, bool)`

GetCustomproperty7Ok returns a tuple with the Customproperty7 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty7

`func (o *UpdateEndpointRequest) SetCustomproperty7(v string)`

SetCustomproperty7 sets Customproperty7 field to given value.

### HasCustomproperty7

`func (o *UpdateEndpointRequest) HasCustomproperty7() bool`

HasCustomproperty7 returns a boolean if a field has been set.

### GetCustomproperty8

`func (o *UpdateEndpointRequest) GetCustomproperty8() string`

GetCustomproperty8 returns the Customproperty8 field if non-nil, zero value otherwise.

### GetCustomproperty8Ok

`func (o *UpdateEndpointRequest) GetCustomproperty8Ok() (*string, bool)`

GetCustomproperty8Ok returns a tuple with the Customproperty8 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty8

`func (o *UpdateEndpointRequest) SetCustomproperty8(v string)`

SetCustomproperty8 sets Customproperty8 field to given value.

### HasCustomproperty8

`func (o *UpdateEndpointRequest) HasCustomproperty8() bool`

HasCustomproperty8 returns a boolean if a field has been set.

### GetCustomproperty9

`func (o *UpdateEndpointRequest) GetCustomproperty9() string`

GetCustomproperty9 returns the Customproperty9 field if non-nil, zero value otherwise.

### GetCustomproperty9Ok

`func (o *UpdateEndpointRequest) GetCustomproperty9Ok() (*string, bool)`

GetCustomproperty9Ok returns a tuple with the Customproperty9 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty9

`func (o *UpdateEndpointRequest) SetCustomproperty9(v string)`

SetCustomproperty9 sets Customproperty9 field to given value.

### HasCustomproperty9

`func (o *UpdateEndpointRequest) HasCustomproperty9() bool`

HasCustomproperty9 returns a boolean if a field has been set.

### GetCustomproperty10

`func (o *UpdateEndpointRequest) GetCustomproperty10() string`

GetCustomproperty10 returns the Customproperty10 field if non-nil, zero value otherwise.

### GetCustomproperty10Ok

`func (o *UpdateEndpointRequest) GetCustomproperty10Ok() (*string, bool)`

GetCustomproperty10Ok returns a tuple with the Customproperty10 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty10

`func (o *UpdateEndpointRequest) SetCustomproperty10(v string)`

SetCustomproperty10 sets Customproperty10 field to given value.

### HasCustomproperty10

`func (o *UpdateEndpointRequest) HasCustomproperty10() bool`

HasCustomproperty10 returns a boolean if a field has been set.

### GetCustomproperty11

`func (o *UpdateEndpointRequest) GetCustomproperty11() string`

GetCustomproperty11 returns the Customproperty11 field if non-nil, zero value otherwise.

### GetCustomproperty11Ok

`func (o *UpdateEndpointRequest) GetCustomproperty11Ok() (*string, bool)`

GetCustomproperty11Ok returns a tuple with the Customproperty11 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty11

`func (o *UpdateEndpointRequest) SetCustomproperty11(v string)`

SetCustomproperty11 sets Customproperty11 field to given value.

### HasCustomproperty11

`func (o *UpdateEndpointRequest) HasCustomproperty11() bool`

HasCustomproperty11 returns a boolean if a field has been set.

### GetCustomproperty12

`func (o *UpdateEndpointRequest) GetCustomproperty12() string`

GetCustomproperty12 returns the Customproperty12 field if non-nil, zero value otherwise.

### GetCustomproperty12Ok

`func (o *UpdateEndpointRequest) GetCustomproperty12Ok() (*string, bool)`

GetCustomproperty12Ok returns a tuple with the Customproperty12 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty12

`func (o *UpdateEndpointRequest) SetCustomproperty12(v string)`

SetCustomproperty12 sets Customproperty12 field to given value.

### HasCustomproperty12

`func (o *UpdateEndpointRequest) HasCustomproperty12() bool`

HasCustomproperty12 returns a boolean if a field has been set.

### GetCustomproperty13

`func (o *UpdateEndpointRequest) GetCustomproperty13() string`

GetCustomproperty13 returns the Customproperty13 field if non-nil, zero value otherwise.

### GetCustomproperty13Ok

`func (o *UpdateEndpointRequest) GetCustomproperty13Ok() (*string, bool)`

GetCustomproperty13Ok returns a tuple with the Customproperty13 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty13

`func (o *UpdateEndpointRequest) SetCustomproperty13(v string)`

SetCustomproperty13 sets Customproperty13 field to given value.

### HasCustomproperty13

`func (o *UpdateEndpointRequest) HasCustomproperty13() bool`

HasCustomproperty13 returns a boolean if a field has been set.

### GetCustomproperty14

`func (o *UpdateEndpointRequest) GetCustomproperty14() string`

GetCustomproperty14 returns the Customproperty14 field if non-nil, zero value otherwise.

### GetCustomproperty14Ok

`func (o *UpdateEndpointRequest) GetCustomproperty14Ok() (*string, bool)`

GetCustomproperty14Ok returns a tuple with the Customproperty14 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty14

`func (o *UpdateEndpointRequest) SetCustomproperty14(v string)`

SetCustomproperty14 sets Customproperty14 field to given value.

### HasCustomproperty14

`func (o *UpdateEndpointRequest) HasCustomproperty14() bool`

HasCustomproperty14 returns a boolean if a field has been set.

### GetCustomproperty15

`func (o *UpdateEndpointRequest) GetCustomproperty15() string`

GetCustomproperty15 returns the Customproperty15 field if non-nil, zero value otherwise.

### GetCustomproperty15Ok

`func (o *UpdateEndpointRequest) GetCustomproperty15Ok() (*string, bool)`

GetCustomproperty15Ok returns a tuple with the Customproperty15 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty15

`func (o *UpdateEndpointRequest) SetCustomproperty15(v string)`

SetCustomproperty15 sets Customproperty15 field to given value.

### HasCustomproperty15

`func (o *UpdateEndpointRequest) HasCustomproperty15() bool`

HasCustomproperty15 returns a boolean if a field has been set.

### GetCustomproperty16

`func (o *UpdateEndpointRequest) GetCustomproperty16() string`

GetCustomproperty16 returns the Customproperty16 field if non-nil, zero value otherwise.

### GetCustomproperty16Ok

`func (o *UpdateEndpointRequest) GetCustomproperty16Ok() (*string, bool)`

GetCustomproperty16Ok returns a tuple with the Customproperty16 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty16

`func (o *UpdateEndpointRequest) SetCustomproperty16(v string)`

SetCustomproperty16 sets Customproperty16 field to given value.

### HasCustomproperty16

`func (o *UpdateEndpointRequest) HasCustomproperty16() bool`

HasCustomproperty16 returns a boolean if a field has been set.

### GetCustomproperty17

`func (o *UpdateEndpointRequest) GetCustomproperty17() string`

GetCustomproperty17 returns the Customproperty17 field if non-nil, zero value otherwise.

### GetCustomproperty17Ok

`func (o *UpdateEndpointRequest) GetCustomproperty17Ok() (*string, bool)`

GetCustomproperty17Ok returns a tuple with the Customproperty17 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty17

`func (o *UpdateEndpointRequest) SetCustomproperty17(v string)`

SetCustomproperty17 sets Customproperty17 field to given value.

### HasCustomproperty17

`func (o *UpdateEndpointRequest) HasCustomproperty17() bool`

HasCustomproperty17 returns a boolean if a field has been set.

### GetCustomproperty18

`func (o *UpdateEndpointRequest) GetCustomproperty18() string`

GetCustomproperty18 returns the Customproperty18 field if non-nil, zero value otherwise.

### GetCustomproperty18Ok

`func (o *UpdateEndpointRequest) GetCustomproperty18Ok() (*string, bool)`

GetCustomproperty18Ok returns a tuple with the Customproperty18 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty18

`func (o *UpdateEndpointRequest) SetCustomproperty18(v string)`

SetCustomproperty18 sets Customproperty18 field to given value.

### HasCustomproperty18

`func (o *UpdateEndpointRequest) HasCustomproperty18() bool`

HasCustomproperty18 returns a boolean if a field has been set.

### GetCustomproperty19

`func (o *UpdateEndpointRequest) GetCustomproperty19() string`

GetCustomproperty19 returns the Customproperty19 field if non-nil, zero value otherwise.

### GetCustomproperty19Ok

`func (o *UpdateEndpointRequest) GetCustomproperty19Ok() (*string, bool)`

GetCustomproperty19Ok returns a tuple with the Customproperty19 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty19

`func (o *UpdateEndpointRequest) SetCustomproperty19(v string)`

SetCustomproperty19 sets Customproperty19 field to given value.

### HasCustomproperty19

`func (o *UpdateEndpointRequest) HasCustomproperty19() bool`

HasCustomproperty19 returns a boolean if a field has been set.

### GetCustomproperty20

`func (o *UpdateEndpointRequest) GetCustomproperty20() string`

GetCustomproperty20 returns the Customproperty20 field if non-nil, zero value otherwise.

### GetCustomproperty20Ok

`func (o *UpdateEndpointRequest) GetCustomproperty20Ok() (*string, bool)`

GetCustomproperty20Ok returns a tuple with the Customproperty20 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty20

`func (o *UpdateEndpointRequest) SetCustomproperty20(v string)`

SetCustomproperty20 sets Customproperty20 field to given value.

### HasCustomproperty20

`func (o *UpdateEndpointRequest) HasCustomproperty20() bool`

HasCustomproperty20 returns a boolean if a field has been set.

### GetCustomproperty21

`func (o *UpdateEndpointRequest) GetCustomproperty21() string`

GetCustomproperty21 returns the Customproperty21 field if non-nil, zero value otherwise.

### GetCustomproperty21Ok

`func (o *UpdateEndpointRequest) GetCustomproperty21Ok() (*string, bool)`

GetCustomproperty21Ok returns a tuple with the Customproperty21 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty21

`func (o *UpdateEndpointRequest) SetCustomproperty21(v string)`

SetCustomproperty21 sets Customproperty21 field to given value.

### HasCustomproperty21

`func (o *UpdateEndpointRequest) HasCustomproperty21() bool`

HasCustomproperty21 returns a boolean if a field has been set.

### GetCustomproperty22

`func (o *UpdateEndpointRequest) GetCustomproperty22() string`

GetCustomproperty22 returns the Customproperty22 field if non-nil, zero value otherwise.

### GetCustomproperty22Ok

`func (o *UpdateEndpointRequest) GetCustomproperty22Ok() (*string, bool)`

GetCustomproperty22Ok returns a tuple with the Customproperty22 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty22

`func (o *UpdateEndpointRequest) SetCustomproperty22(v string)`

SetCustomproperty22 sets Customproperty22 field to given value.

### HasCustomproperty22

`func (o *UpdateEndpointRequest) HasCustomproperty22() bool`

HasCustomproperty22 returns a boolean if a field has been set.

### GetCustomproperty23

`func (o *UpdateEndpointRequest) GetCustomproperty23() string`

GetCustomproperty23 returns the Customproperty23 field if non-nil, zero value otherwise.

### GetCustomproperty23Ok

`func (o *UpdateEndpointRequest) GetCustomproperty23Ok() (*string, bool)`

GetCustomproperty23Ok returns a tuple with the Customproperty23 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty23

`func (o *UpdateEndpointRequest) SetCustomproperty23(v string)`

SetCustomproperty23 sets Customproperty23 field to given value.

### HasCustomproperty23

`func (o *UpdateEndpointRequest) HasCustomproperty23() bool`

HasCustomproperty23 returns a boolean if a field has been set.

### GetCustomproperty24

`func (o *UpdateEndpointRequest) GetCustomproperty24() string`

GetCustomproperty24 returns the Customproperty24 field if non-nil, zero value otherwise.

### GetCustomproperty24Ok

`func (o *UpdateEndpointRequest) GetCustomproperty24Ok() (*string, bool)`

GetCustomproperty24Ok returns a tuple with the Customproperty24 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty24

`func (o *UpdateEndpointRequest) SetCustomproperty24(v string)`

SetCustomproperty24 sets Customproperty24 field to given value.

### HasCustomproperty24

`func (o *UpdateEndpointRequest) HasCustomproperty24() bool`

HasCustomproperty24 returns a boolean if a field has been set.

### GetCustomproperty25

`func (o *UpdateEndpointRequest) GetCustomproperty25() string`

GetCustomproperty25 returns the Customproperty25 field if non-nil, zero value otherwise.

### GetCustomproperty25Ok

`func (o *UpdateEndpointRequest) GetCustomproperty25Ok() (*string, bool)`

GetCustomproperty25Ok returns a tuple with the Customproperty25 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty25

`func (o *UpdateEndpointRequest) SetCustomproperty25(v string)`

SetCustomproperty25 sets Customproperty25 field to given value.

### HasCustomproperty25

`func (o *UpdateEndpointRequest) HasCustomproperty25() bool`

HasCustomproperty25 returns a boolean if a field has been set.

### GetCustomproperty26

`func (o *UpdateEndpointRequest) GetCustomproperty26() string`

GetCustomproperty26 returns the Customproperty26 field if non-nil, zero value otherwise.

### GetCustomproperty26Ok

`func (o *UpdateEndpointRequest) GetCustomproperty26Ok() (*string, bool)`

GetCustomproperty26Ok returns a tuple with the Customproperty26 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty26

`func (o *UpdateEndpointRequest) SetCustomproperty26(v string)`

SetCustomproperty26 sets Customproperty26 field to given value.

### HasCustomproperty26

`func (o *UpdateEndpointRequest) HasCustomproperty26() bool`

HasCustomproperty26 returns a boolean if a field has been set.

### GetCustomproperty27

`func (o *UpdateEndpointRequest) GetCustomproperty27() string`

GetCustomproperty27 returns the Customproperty27 field if non-nil, zero value otherwise.

### GetCustomproperty27Ok

`func (o *UpdateEndpointRequest) GetCustomproperty27Ok() (*string, bool)`

GetCustomproperty27Ok returns a tuple with the Customproperty27 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty27

`func (o *UpdateEndpointRequest) SetCustomproperty27(v string)`

SetCustomproperty27 sets Customproperty27 field to given value.

### HasCustomproperty27

`func (o *UpdateEndpointRequest) HasCustomproperty27() bool`

HasCustomproperty27 returns a boolean if a field has been set.

### GetCustomproperty28

`func (o *UpdateEndpointRequest) GetCustomproperty28() string`

GetCustomproperty28 returns the Customproperty28 field if non-nil, zero value otherwise.

### GetCustomproperty28Ok

`func (o *UpdateEndpointRequest) GetCustomproperty28Ok() (*string, bool)`

GetCustomproperty28Ok returns a tuple with the Customproperty28 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty28

`func (o *UpdateEndpointRequest) SetCustomproperty28(v string)`

SetCustomproperty28 sets Customproperty28 field to given value.

### HasCustomproperty28

`func (o *UpdateEndpointRequest) HasCustomproperty28() bool`

HasCustomproperty28 returns a boolean if a field has been set.

### GetCustomproperty29

`func (o *UpdateEndpointRequest) GetCustomproperty29() string`

GetCustomproperty29 returns the Customproperty29 field if non-nil, zero value otherwise.

### GetCustomproperty29Ok

`func (o *UpdateEndpointRequest) GetCustomproperty29Ok() (*string, bool)`

GetCustomproperty29Ok returns a tuple with the Customproperty29 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty29

`func (o *UpdateEndpointRequest) SetCustomproperty29(v string)`

SetCustomproperty29 sets Customproperty29 field to given value.

### HasCustomproperty29

`func (o *UpdateEndpointRequest) HasCustomproperty29() bool`

HasCustomproperty29 returns a boolean if a field has been set.

### GetCustomproperty30

`func (o *UpdateEndpointRequest) GetCustomproperty30() string`

GetCustomproperty30 returns the Customproperty30 field if non-nil, zero value otherwise.

### GetCustomproperty30Ok

`func (o *UpdateEndpointRequest) GetCustomproperty30Ok() (*string, bool)`

GetCustomproperty30Ok returns a tuple with the Customproperty30 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty30

`func (o *UpdateEndpointRequest) SetCustomproperty30(v string)`

SetCustomproperty30 sets Customproperty30 field to given value.

### HasCustomproperty30

`func (o *UpdateEndpointRequest) HasCustomproperty30() bool`

HasCustomproperty30 returns a boolean if a field has been set.

### GetCustomproperty31

`func (o *UpdateEndpointRequest) GetCustomproperty31() string`

GetCustomproperty31 returns the Customproperty31 field if non-nil, zero value otherwise.

### GetCustomproperty31Ok

`func (o *UpdateEndpointRequest) GetCustomproperty31Ok() (*string, bool)`

GetCustomproperty31Ok returns a tuple with the Customproperty31 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty31

`func (o *UpdateEndpointRequest) SetCustomproperty31(v string)`

SetCustomproperty31 sets Customproperty31 field to given value.

### HasCustomproperty31

`func (o *UpdateEndpointRequest) HasCustomproperty31() bool`

HasCustomproperty31 returns a boolean if a field has been set.

### GetCustomproperty32

`func (o *UpdateEndpointRequest) GetCustomproperty32() string`

GetCustomproperty32 returns the Customproperty32 field if non-nil, zero value otherwise.

### GetCustomproperty32Ok

`func (o *UpdateEndpointRequest) GetCustomproperty32Ok() (*string, bool)`

GetCustomproperty32Ok returns a tuple with the Customproperty32 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty32

`func (o *UpdateEndpointRequest) SetCustomproperty32(v string)`

SetCustomproperty32 sets Customproperty32 field to given value.

### HasCustomproperty32

`func (o *UpdateEndpointRequest) HasCustomproperty32() bool`

HasCustomproperty32 returns a boolean if a field has been set.

### GetCustomproperty33

`func (o *UpdateEndpointRequest) GetCustomproperty33() string`

GetCustomproperty33 returns the Customproperty33 field if non-nil, zero value otherwise.

### GetCustomproperty33Ok

`func (o *UpdateEndpointRequest) GetCustomproperty33Ok() (*string, bool)`

GetCustomproperty33Ok returns a tuple with the Customproperty33 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty33

`func (o *UpdateEndpointRequest) SetCustomproperty33(v string)`

SetCustomproperty33 sets Customproperty33 field to given value.

### HasCustomproperty33

`func (o *UpdateEndpointRequest) HasCustomproperty33() bool`

HasCustomproperty33 returns a boolean if a field has been set.

### GetCustomproperty34

`func (o *UpdateEndpointRequest) GetCustomproperty34() string`

GetCustomproperty34 returns the Customproperty34 field if non-nil, zero value otherwise.

### GetCustomproperty34Ok

`func (o *UpdateEndpointRequest) GetCustomproperty34Ok() (*string, bool)`

GetCustomproperty34Ok returns a tuple with the Customproperty34 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty34

`func (o *UpdateEndpointRequest) SetCustomproperty34(v string)`

SetCustomproperty34 sets Customproperty34 field to given value.

### HasCustomproperty34

`func (o *UpdateEndpointRequest) HasCustomproperty34() bool`

HasCustomproperty34 returns a boolean if a field has been set.

### GetCustomproperty35

`func (o *UpdateEndpointRequest) GetCustomproperty35() string`

GetCustomproperty35 returns the Customproperty35 field if non-nil, zero value otherwise.

### GetCustomproperty35Ok

`func (o *UpdateEndpointRequest) GetCustomproperty35Ok() (*string, bool)`

GetCustomproperty35Ok returns a tuple with the Customproperty35 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty35

`func (o *UpdateEndpointRequest) SetCustomproperty35(v string)`

SetCustomproperty35 sets Customproperty35 field to given value.

### HasCustomproperty35

`func (o *UpdateEndpointRequest) HasCustomproperty35() bool`

HasCustomproperty35 returns a boolean if a field has been set.

### GetCustomproperty36

`func (o *UpdateEndpointRequest) GetCustomproperty36() string`

GetCustomproperty36 returns the Customproperty36 field if non-nil, zero value otherwise.

### GetCustomproperty36Ok

`func (o *UpdateEndpointRequest) GetCustomproperty36Ok() (*string, bool)`

GetCustomproperty36Ok returns a tuple with the Customproperty36 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty36

`func (o *UpdateEndpointRequest) SetCustomproperty36(v string)`

SetCustomproperty36 sets Customproperty36 field to given value.

### HasCustomproperty36

`func (o *UpdateEndpointRequest) HasCustomproperty36() bool`

HasCustomproperty36 returns a boolean if a field has been set.

### GetCustomproperty37

`func (o *UpdateEndpointRequest) GetCustomproperty37() string`

GetCustomproperty37 returns the Customproperty37 field if non-nil, zero value otherwise.

### GetCustomproperty37Ok

`func (o *UpdateEndpointRequest) GetCustomproperty37Ok() (*string, bool)`

GetCustomproperty37Ok returns a tuple with the Customproperty37 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty37

`func (o *UpdateEndpointRequest) SetCustomproperty37(v string)`

SetCustomproperty37 sets Customproperty37 field to given value.

### HasCustomproperty37

`func (o *UpdateEndpointRequest) HasCustomproperty37() bool`

HasCustomproperty37 returns a boolean if a field has been set.

### GetCustomproperty38

`func (o *UpdateEndpointRequest) GetCustomproperty38() string`

GetCustomproperty38 returns the Customproperty38 field if non-nil, zero value otherwise.

### GetCustomproperty38Ok

`func (o *UpdateEndpointRequest) GetCustomproperty38Ok() (*string, bool)`

GetCustomproperty38Ok returns a tuple with the Customproperty38 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty38

`func (o *UpdateEndpointRequest) SetCustomproperty38(v string)`

SetCustomproperty38 sets Customproperty38 field to given value.

### HasCustomproperty38

`func (o *UpdateEndpointRequest) HasCustomproperty38() bool`

HasCustomproperty38 returns a boolean if a field has been set.

### GetCustomproperty39

`func (o *UpdateEndpointRequest) GetCustomproperty39() string`

GetCustomproperty39 returns the Customproperty39 field if non-nil, zero value otherwise.

### GetCustomproperty39Ok

`func (o *UpdateEndpointRequest) GetCustomproperty39Ok() (*string, bool)`

GetCustomproperty39Ok returns a tuple with the Customproperty39 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty39

`func (o *UpdateEndpointRequest) SetCustomproperty39(v string)`

SetCustomproperty39 sets Customproperty39 field to given value.

### HasCustomproperty39

`func (o *UpdateEndpointRequest) HasCustomproperty39() bool`

HasCustomproperty39 returns a boolean if a field has been set.

### GetCustomproperty40

`func (o *UpdateEndpointRequest) GetCustomproperty40() string`

GetCustomproperty40 returns the Customproperty40 field if non-nil, zero value otherwise.

### GetCustomproperty40Ok

`func (o *UpdateEndpointRequest) GetCustomproperty40Ok() (*string, bool)`

GetCustomproperty40Ok returns a tuple with the Customproperty40 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty40

`func (o *UpdateEndpointRequest) SetCustomproperty40(v string)`

SetCustomproperty40 sets Customproperty40 field to given value.

### HasCustomproperty40

`func (o *UpdateEndpointRequest) HasCustomproperty40() bool`

HasCustomproperty40 returns a boolean if a field has been set.

### GetCustomproperty41

`func (o *UpdateEndpointRequest) GetCustomproperty41() string`

GetCustomproperty41 returns the Customproperty41 field if non-nil, zero value otherwise.

### GetCustomproperty41Ok

`func (o *UpdateEndpointRequest) GetCustomproperty41Ok() (*string, bool)`

GetCustomproperty41Ok returns a tuple with the Customproperty41 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty41

`func (o *UpdateEndpointRequest) SetCustomproperty41(v string)`

SetCustomproperty41 sets Customproperty41 field to given value.

### HasCustomproperty41

`func (o *UpdateEndpointRequest) HasCustomproperty41() bool`

HasCustomproperty41 returns a boolean if a field has been set.

### GetCustomproperty42

`func (o *UpdateEndpointRequest) GetCustomproperty42() string`

GetCustomproperty42 returns the Customproperty42 field if non-nil, zero value otherwise.

### GetCustomproperty42Ok

`func (o *UpdateEndpointRequest) GetCustomproperty42Ok() (*string, bool)`

GetCustomproperty42Ok returns a tuple with the Customproperty42 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty42

`func (o *UpdateEndpointRequest) SetCustomproperty42(v string)`

SetCustomproperty42 sets Customproperty42 field to given value.

### HasCustomproperty42

`func (o *UpdateEndpointRequest) HasCustomproperty42() bool`

HasCustomproperty42 returns a boolean if a field has been set.

### GetCustomproperty43

`func (o *UpdateEndpointRequest) GetCustomproperty43() string`

GetCustomproperty43 returns the Customproperty43 field if non-nil, zero value otherwise.

### GetCustomproperty43Ok

`func (o *UpdateEndpointRequest) GetCustomproperty43Ok() (*string, bool)`

GetCustomproperty43Ok returns a tuple with the Customproperty43 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty43

`func (o *UpdateEndpointRequest) SetCustomproperty43(v string)`

SetCustomproperty43 sets Customproperty43 field to given value.

### HasCustomproperty43

`func (o *UpdateEndpointRequest) HasCustomproperty43() bool`

HasCustomproperty43 returns a boolean if a field has been set.

### GetCustomproperty44

`func (o *UpdateEndpointRequest) GetCustomproperty44() string`

GetCustomproperty44 returns the Customproperty44 field if non-nil, zero value otherwise.

### GetCustomproperty44Ok

`func (o *UpdateEndpointRequest) GetCustomproperty44Ok() (*string, bool)`

GetCustomproperty44Ok returns a tuple with the Customproperty44 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty44

`func (o *UpdateEndpointRequest) SetCustomproperty44(v string)`

SetCustomproperty44 sets Customproperty44 field to given value.

### HasCustomproperty44

`func (o *UpdateEndpointRequest) HasCustomproperty44() bool`

HasCustomproperty44 returns a boolean if a field has been set.

### GetCustomproperty45

`func (o *UpdateEndpointRequest) GetCustomproperty45() string`

GetCustomproperty45 returns the Customproperty45 field if non-nil, zero value otherwise.

### GetCustomproperty45Ok

`func (o *UpdateEndpointRequest) GetCustomproperty45Ok() (*string, bool)`

GetCustomproperty45Ok returns a tuple with the Customproperty45 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty45

`func (o *UpdateEndpointRequest) SetCustomproperty45(v string)`

SetCustomproperty45 sets Customproperty45 field to given value.

### HasCustomproperty45

`func (o *UpdateEndpointRequest) HasCustomproperty45() bool`

HasCustomproperty45 returns a boolean if a field has been set.

### GetCustomproperty1Label

`func (o *UpdateEndpointRequest) GetCustomproperty1Label() string`

GetCustomproperty1Label returns the Customproperty1Label field if non-nil, zero value otherwise.

### GetCustomproperty1LabelOk

`func (o *UpdateEndpointRequest) GetCustomproperty1LabelOk() (*string, bool)`

GetCustomproperty1LabelOk returns a tuple with the Customproperty1Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty1Label

`func (o *UpdateEndpointRequest) SetCustomproperty1Label(v string)`

SetCustomproperty1Label sets Customproperty1Label field to given value.

### HasCustomproperty1Label

`func (o *UpdateEndpointRequest) HasCustomproperty1Label() bool`

HasCustomproperty1Label returns a boolean if a field has been set.

### GetCustomproperty2Label

`func (o *UpdateEndpointRequest) GetCustomproperty2Label() string`

GetCustomproperty2Label returns the Customproperty2Label field if non-nil, zero value otherwise.

### GetCustomproperty2LabelOk

`func (o *UpdateEndpointRequest) GetCustomproperty2LabelOk() (*string, bool)`

GetCustomproperty2LabelOk returns a tuple with the Customproperty2Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty2Label

`func (o *UpdateEndpointRequest) SetCustomproperty2Label(v string)`

SetCustomproperty2Label sets Customproperty2Label field to given value.

### HasCustomproperty2Label

`func (o *UpdateEndpointRequest) HasCustomproperty2Label() bool`

HasCustomproperty2Label returns a boolean if a field has been set.

### GetCustomproperty3Label

`func (o *UpdateEndpointRequest) GetCustomproperty3Label() string`

GetCustomproperty3Label returns the Customproperty3Label field if non-nil, zero value otherwise.

### GetCustomproperty3LabelOk

`func (o *UpdateEndpointRequest) GetCustomproperty3LabelOk() (*string, bool)`

GetCustomproperty3LabelOk returns a tuple with the Customproperty3Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty3Label

`func (o *UpdateEndpointRequest) SetCustomproperty3Label(v string)`

SetCustomproperty3Label sets Customproperty3Label field to given value.

### HasCustomproperty3Label

`func (o *UpdateEndpointRequest) HasCustomproperty3Label() bool`

HasCustomproperty3Label returns a boolean if a field has been set.

### GetCustomproperty4Label

`func (o *UpdateEndpointRequest) GetCustomproperty4Label() string`

GetCustomproperty4Label returns the Customproperty4Label field if non-nil, zero value otherwise.

### GetCustomproperty4LabelOk

`func (o *UpdateEndpointRequest) GetCustomproperty4LabelOk() (*string, bool)`

GetCustomproperty4LabelOk returns a tuple with the Customproperty4Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty4Label

`func (o *UpdateEndpointRequest) SetCustomproperty4Label(v string)`

SetCustomproperty4Label sets Customproperty4Label field to given value.

### HasCustomproperty4Label

`func (o *UpdateEndpointRequest) HasCustomproperty4Label() bool`

HasCustomproperty4Label returns a boolean if a field has been set.

### GetCustomproperty5Label

`func (o *UpdateEndpointRequest) GetCustomproperty5Label() string`

GetCustomproperty5Label returns the Customproperty5Label field if non-nil, zero value otherwise.

### GetCustomproperty5LabelOk

`func (o *UpdateEndpointRequest) GetCustomproperty5LabelOk() (*string, bool)`

GetCustomproperty5LabelOk returns a tuple with the Customproperty5Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty5Label

`func (o *UpdateEndpointRequest) SetCustomproperty5Label(v string)`

SetCustomproperty5Label sets Customproperty5Label field to given value.

### HasCustomproperty5Label

`func (o *UpdateEndpointRequest) HasCustomproperty5Label() bool`

HasCustomproperty5Label returns a boolean if a field has been set.

### GetCustomproperty6Label

`func (o *UpdateEndpointRequest) GetCustomproperty6Label() string`

GetCustomproperty6Label returns the Customproperty6Label field if non-nil, zero value otherwise.

### GetCustomproperty6LabelOk

`func (o *UpdateEndpointRequest) GetCustomproperty6LabelOk() (*string, bool)`

GetCustomproperty6LabelOk returns a tuple with the Customproperty6Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty6Label

`func (o *UpdateEndpointRequest) SetCustomproperty6Label(v string)`

SetCustomproperty6Label sets Customproperty6Label field to given value.

### HasCustomproperty6Label

`func (o *UpdateEndpointRequest) HasCustomproperty6Label() bool`

HasCustomproperty6Label returns a boolean if a field has been set.

### GetCustomproperty7Label

`func (o *UpdateEndpointRequest) GetCustomproperty7Label() string`

GetCustomproperty7Label returns the Customproperty7Label field if non-nil, zero value otherwise.

### GetCustomproperty7LabelOk

`func (o *UpdateEndpointRequest) GetCustomproperty7LabelOk() (*string, bool)`

GetCustomproperty7LabelOk returns a tuple with the Customproperty7Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty7Label

`func (o *UpdateEndpointRequest) SetCustomproperty7Label(v string)`

SetCustomproperty7Label sets Customproperty7Label field to given value.

### HasCustomproperty7Label

`func (o *UpdateEndpointRequest) HasCustomproperty7Label() bool`

HasCustomproperty7Label returns a boolean if a field has been set.

### GetCustomproperty8Label

`func (o *UpdateEndpointRequest) GetCustomproperty8Label() string`

GetCustomproperty8Label returns the Customproperty8Label field if non-nil, zero value otherwise.

### GetCustomproperty8LabelOk

`func (o *UpdateEndpointRequest) GetCustomproperty8LabelOk() (*string, bool)`

GetCustomproperty8LabelOk returns a tuple with the Customproperty8Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty8Label

`func (o *UpdateEndpointRequest) SetCustomproperty8Label(v string)`

SetCustomproperty8Label sets Customproperty8Label field to given value.

### HasCustomproperty8Label

`func (o *UpdateEndpointRequest) HasCustomproperty8Label() bool`

HasCustomproperty8Label returns a boolean if a field has been set.

### GetCustomproperty9Label

`func (o *UpdateEndpointRequest) GetCustomproperty9Label() string`

GetCustomproperty9Label returns the Customproperty9Label field if non-nil, zero value otherwise.

### GetCustomproperty9LabelOk

`func (o *UpdateEndpointRequest) GetCustomproperty9LabelOk() (*string, bool)`

GetCustomproperty9LabelOk returns a tuple with the Customproperty9Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty9Label

`func (o *UpdateEndpointRequest) SetCustomproperty9Label(v string)`

SetCustomproperty9Label sets Customproperty9Label field to given value.

### HasCustomproperty9Label

`func (o *UpdateEndpointRequest) HasCustomproperty9Label() bool`

HasCustomproperty9Label returns a boolean if a field has been set.

### GetCustomproperty10Label

`func (o *UpdateEndpointRequest) GetCustomproperty10Label() string`

GetCustomproperty10Label returns the Customproperty10Label field if non-nil, zero value otherwise.

### GetCustomproperty10LabelOk

`func (o *UpdateEndpointRequest) GetCustomproperty10LabelOk() (*string, bool)`

GetCustomproperty10LabelOk returns a tuple with the Customproperty10Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty10Label

`func (o *UpdateEndpointRequest) SetCustomproperty10Label(v string)`

SetCustomproperty10Label sets Customproperty10Label field to given value.

### HasCustomproperty10Label

`func (o *UpdateEndpointRequest) HasCustomproperty10Label() bool`

HasCustomproperty10Label returns a boolean if a field has been set.

### GetCustomproperty11Label

`func (o *UpdateEndpointRequest) GetCustomproperty11Label() string`

GetCustomproperty11Label returns the Customproperty11Label field if non-nil, zero value otherwise.

### GetCustomproperty11LabelOk

`func (o *UpdateEndpointRequest) GetCustomproperty11LabelOk() (*string, bool)`

GetCustomproperty11LabelOk returns a tuple with the Customproperty11Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty11Label

`func (o *UpdateEndpointRequest) SetCustomproperty11Label(v string)`

SetCustomproperty11Label sets Customproperty11Label field to given value.

### HasCustomproperty11Label

`func (o *UpdateEndpointRequest) HasCustomproperty11Label() bool`

HasCustomproperty11Label returns a boolean if a field has been set.

### GetCustomproperty12Label

`func (o *UpdateEndpointRequest) GetCustomproperty12Label() string`

GetCustomproperty12Label returns the Customproperty12Label field if non-nil, zero value otherwise.

### GetCustomproperty12LabelOk

`func (o *UpdateEndpointRequest) GetCustomproperty12LabelOk() (*string, bool)`

GetCustomproperty12LabelOk returns a tuple with the Customproperty12Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty12Label

`func (o *UpdateEndpointRequest) SetCustomproperty12Label(v string)`

SetCustomproperty12Label sets Customproperty12Label field to given value.

### HasCustomproperty12Label

`func (o *UpdateEndpointRequest) HasCustomproperty12Label() bool`

HasCustomproperty12Label returns a boolean if a field has been set.

### GetCustomproperty13Label

`func (o *UpdateEndpointRequest) GetCustomproperty13Label() string`

GetCustomproperty13Label returns the Customproperty13Label field if non-nil, zero value otherwise.

### GetCustomproperty13LabelOk

`func (o *UpdateEndpointRequest) GetCustomproperty13LabelOk() (*string, bool)`

GetCustomproperty13LabelOk returns a tuple with the Customproperty13Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty13Label

`func (o *UpdateEndpointRequest) SetCustomproperty13Label(v string)`

SetCustomproperty13Label sets Customproperty13Label field to given value.

### HasCustomproperty13Label

`func (o *UpdateEndpointRequest) HasCustomproperty13Label() bool`

HasCustomproperty13Label returns a boolean if a field has been set.

### GetCustomproperty14Label

`func (o *UpdateEndpointRequest) GetCustomproperty14Label() string`

GetCustomproperty14Label returns the Customproperty14Label field if non-nil, zero value otherwise.

### GetCustomproperty14LabelOk

`func (o *UpdateEndpointRequest) GetCustomproperty14LabelOk() (*string, bool)`

GetCustomproperty14LabelOk returns a tuple with the Customproperty14Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty14Label

`func (o *UpdateEndpointRequest) SetCustomproperty14Label(v string)`

SetCustomproperty14Label sets Customproperty14Label field to given value.

### HasCustomproperty14Label

`func (o *UpdateEndpointRequest) HasCustomproperty14Label() bool`

HasCustomproperty14Label returns a boolean if a field has been set.

### GetCustomproperty15Label

`func (o *UpdateEndpointRequest) GetCustomproperty15Label() string`

GetCustomproperty15Label returns the Customproperty15Label field if non-nil, zero value otherwise.

### GetCustomproperty15LabelOk

`func (o *UpdateEndpointRequest) GetCustomproperty15LabelOk() (*string, bool)`

GetCustomproperty15LabelOk returns a tuple with the Customproperty15Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty15Label

`func (o *UpdateEndpointRequest) SetCustomproperty15Label(v string)`

SetCustomproperty15Label sets Customproperty15Label field to given value.

### HasCustomproperty15Label

`func (o *UpdateEndpointRequest) HasCustomproperty15Label() bool`

HasCustomproperty15Label returns a boolean if a field has been set.

### GetCustomproperty16Label

`func (o *UpdateEndpointRequest) GetCustomproperty16Label() string`

GetCustomproperty16Label returns the Customproperty16Label field if non-nil, zero value otherwise.

### GetCustomproperty16LabelOk

`func (o *UpdateEndpointRequest) GetCustomproperty16LabelOk() (*string, bool)`

GetCustomproperty16LabelOk returns a tuple with the Customproperty16Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty16Label

`func (o *UpdateEndpointRequest) SetCustomproperty16Label(v string)`

SetCustomproperty16Label sets Customproperty16Label field to given value.

### HasCustomproperty16Label

`func (o *UpdateEndpointRequest) HasCustomproperty16Label() bool`

HasCustomproperty16Label returns a boolean if a field has been set.

### GetCustomproperty17Label

`func (o *UpdateEndpointRequest) GetCustomproperty17Label() string`

GetCustomproperty17Label returns the Customproperty17Label field if non-nil, zero value otherwise.

### GetCustomproperty17LabelOk

`func (o *UpdateEndpointRequest) GetCustomproperty17LabelOk() (*string, bool)`

GetCustomproperty17LabelOk returns a tuple with the Customproperty17Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty17Label

`func (o *UpdateEndpointRequest) SetCustomproperty17Label(v string)`

SetCustomproperty17Label sets Customproperty17Label field to given value.

### HasCustomproperty17Label

`func (o *UpdateEndpointRequest) HasCustomproperty17Label() bool`

HasCustomproperty17Label returns a boolean if a field has been set.

### GetCustomproperty18Label

`func (o *UpdateEndpointRequest) GetCustomproperty18Label() string`

GetCustomproperty18Label returns the Customproperty18Label field if non-nil, zero value otherwise.

### GetCustomproperty18LabelOk

`func (o *UpdateEndpointRequest) GetCustomproperty18LabelOk() (*string, bool)`

GetCustomproperty18LabelOk returns a tuple with the Customproperty18Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty18Label

`func (o *UpdateEndpointRequest) SetCustomproperty18Label(v string)`

SetCustomproperty18Label sets Customproperty18Label field to given value.

### HasCustomproperty18Label

`func (o *UpdateEndpointRequest) HasCustomproperty18Label() bool`

HasCustomproperty18Label returns a boolean if a field has been set.

### GetCustomproperty19Label

`func (o *UpdateEndpointRequest) GetCustomproperty19Label() string`

GetCustomproperty19Label returns the Customproperty19Label field if non-nil, zero value otherwise.

### GetCustomproperty19LabelOk

`func (o *UpdateEndpointRequest) GetCustomproperty19LabelOk() (*string, bool)`

GetCustomproperty19LabelOk returns a tuple with the Customproperty19Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty19Label

`func (o *UpdateEndpointRequest) SetCustomproperty19Label(v string)`

SetCustomproperty19Label sets Customproperty19Label field to given value.

### HasCustomproperty19Label

`func (o *UpdateEndpointRequest) HasCustomproperty19Label() bool`

HasCustomproperty19Label returns a boolean if a field has been set.

### GetCustomproperty20Label

`func (o *UpdateEndpointRequest) GetCustomproperty20Label() string`

GetCustomproperty20Label returns the Customproperty20Label field if non-nil, zero value otherwise.

### GetCustomproperty20LabelOk

`func (o *UpdateEndpointRequest) GetCustomproperty20LabelOk() (*string, bool)`

GetCustomproperty20LabelOk returns a tuple with the Customproperty20Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty20Label

`func (o *UpdateEndpointRequest) SetCustomproperty20Label(v string)`

SetCustomproperty20Label sets Customproperty20Label field to given value.

### HasCustomproperty20Label

`func (o *UpdateEndpointRequest) HasCustomproperty20Label() bool`

HasCustomproperty20Label returns a boolean if a field has been set.

### GetCustomproperty21Label

`func (o *UpdateEndpointRequest) GetCustomproperty21Label() string`

GetCustomproperty21Label returns the Customproperty21Label field if non-nil, zero value otherwise.

### GetCustomproperty21LabelOk

`func (o *UpdateEndpointRequest) GetCustomproperty21LabelOk() (*string, bool)`

GetCustomproperty21LabelOk returns a tuple with the Customproperty21Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty21Label

`func (o *UpdateEndpointRequest) SetCustomproperty21Label(v string)`

SetCustomproperty21Label sets Customproperty21Label field to given value.

### HasCustomproperty21Label

`func (o *UpdateEndpointRequest) HasCustomproperty21Label() bool`

HasCustomproperty21Label returns a boolean if a field has been set.

### GetCustomproperty22Label

`func (o *UpdateEndpointRequest) GetCustomproperty22Label() string`

GetCustomproperty22Label returns the Customproperty22Label field if non-nil, zero value otherwise.

### GetCustomproperty22LabelOk

`func (o *UpdateEndpointRequest) GetCustomproperty22LabelOk() (*string, bool)`

GetCustomproperty22LabelOk returns a tuple with the Customproperty22Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty22Label

`func (o *UpdateEndpointRequest) SetCustomproperty22Label(v string)`

SetCustomproperty22Label sets Customproperty22Label field to given value.

### HasCustomproperty22Label

`func (o *UpdateEndpointRequest) HasCustomproperty22Label() bool`

HasCustomproperty22Label returns a boolean if a field has been set.

### GetCustomproperty23Label

`func (o *UpdateEndpointRequest) GetCustomproperty23Label() string`

GetCustomproperty23Label returns the Customproperty23Label field if non-nil, zero value otherwise.

### GetCustomproperty23LabelOk

`func (o *UpdateEndpointRequest) GetCustomproperty23LabelOk() (*string, bool)`

GetCustomproperty23LabelOk returns a tuple with the Customproperty23Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty23Label

`func (o *UpdateEndpointRequest) SetCustomproperty23Label(v string)`

SetCustomproperty23Label sets Customproperty23Label field to given value.

### HasCustomproperty23Label

`func (o *UpdateEndpointRequest) HasCustomproperty23Label() bool`

HasCustomproperty23Label returns a boolean if a field has been set.

### GetCustomproperty24Label

`func (o *UpdateEndpointRequest) GetCustomproperty24Label() string`

GetCustomproperty24Label returns the Customproperty24Label field if non-nil, zero value otherwise.

### GetCustomproperty24LabelOk

`func (o *UpdateEndpointRequest) GetCustomproperty24LabelOk() (*string, bool)`

GetCustomproperty24LabelOk returns a tuple with the Customproperty24Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty24Label

`func (o *UpdateEndpointRequest) SetCustomproperty24Label(v string)`

SetCustomproperty24Label sets Customproperty24Label field to given value.

### HasCustomproperty24Label

`func (o *UpdateEndpointRequest) HasCustomproperty24Label() bool`

HasCustomproperty24Label returns a boolean if a field has been set.

### GetCustomproperty25Label

`func (o *UpdateEndpointRequest) GetCustomproperty25Label() string`

GetCustomproperty25Label returns the Customproperty25Label field if non-nil, zero value otherwise.

### GetCustomproperty25LabelOk

`func (o *UpdateEndpointRequest) GetCustomproperty25LabelOk() (*string, bool)`

GetCustomproperty25LabelOk returns a tuple with the Customproperty25Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty25Label

`func (o *UpdateEndpointRequest) SetCustomproperty25Label(v string)`

SetCustomproperty25Label sets Customproperty25Label field to given value.

### HasCustomproperty25Label

`func (o *UpdateEndpointRequest) HasCustomproperty25Label() bool`

HasCustomproperty25Label returns a boolean if a field has been set.

### GetCustomproperty26Label

`func (o *UpdateEndpointRequest) GetCustomproperty26Label() string`

GetCustomproperty26Label returns the Customproperty26Label field if non-nil, zero value otherwise.

### GetCustomproperty26LabelOk

`func (o *UpdateEndpointRequest) GetCustomproperty26LabelOk() (*string, bool)`

GetCustomproperty26LabelOk returns a tuple with the Customproperty26Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty26Label

`func (o *UpdateEndpointRequest) SetCustomproperty26Label(v string)`

SetCustomproperty26Label sets Customproperty26Label field to given value.

### HasCustomproperty26Label

`func (o *UpdateEndpointRequest) HasCustomproperty26Label() bool`

HasCustomproperty26Label returns a boolean if a field has been set.

### GetCustomproperty27Label

`func (o *UpdateEndpointRequest) GetCustomproperty27Label() string`

GetCustomproperty27Label returns the Customproperty27Label field if non-nil, zero value otherwise.

### GetCustomproperty27LabelOk

`func (o *UpdateEndpointRequest) GetCustomproperty27LabelOk() (*string, bool)`

GetCustomproperty27LabelOk returns a tuple with the Customproperty27Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty27Label

`func (o *UpdateEndpointRequest) SetCustomproperty27Label(v string)`

SetCustomproperty27Label sets Customproperty27Label field to given value.

### HasCustomproperty27Label

`func (o *UpdateEndpointRequest) HasCustomproperty27Label() bool`

HasCustomproperty27Label returns a boolean if a field has been set.

### GetCustomproperty28Label

`func (o *UpdateEndpointRequest) GetCustomproperty28Label() string`

GetCustomproperty28Label returns the Customproperty28Label field if non-nil, zero value otherwise.

### GetCustomproperty28LabelOk

`func (o *UpdateEndpointRequest) GetCustomproperty28LabelOk() (*string, bool)`

GetCustomproperty28LabelOk returns a tuple with the Customproperty28Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty28Label

`func (o *UpdateEndpointRequest) SetCustomproperty28Label(v string)`

SetCustomproperty28Label sets Customproperty28Label field to given value.

### HasCustomproperty28Label

`func (o *UpdateEndpointRequest) HasCustomproperty28Label() bool`

HasCustomproperty28Label returns a boolean if a field has been set.

### GetCustomproperty29Label

`func (o *UpdateEndpointRequest) GetCustomproperty29Label() string`

GetCustomproperty29Label returns the Customproperty29Label field if non-nil, zero value otherwise.

### GetCustomproperty29LabelOk

`func (o *UpdateEndpointRequest) GetCustomproperty29LabelOk() (*string, bool)`

GetCustomproperty29LabelOk returns a tuple with the Customproperty29Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty29Label

`func (o *UpdateEndpointRequest) SetCustomproperty29Label(v string)`

SetCustomproperty29Label sets Customproperty29Label field to given value.

### HasCustomproperty29Label

`func (o *UpdateEndpointRequest) HasCustomproperty29Label() bool`

HasCustomproperty29Label returns a boolean if a field has been set.

### GetCustomproperty30Label

`func (o *UpdateEndpointRequest) GetCustomproperty30Label() string`

GetCustomproperty30Label returns the Customproperty30Label field if non-nil, zero value otherwise.

### GetCustomproperty30LabelOk

`func (o *UpdateEndpointRequest) GetCustomproperty30LabelOk() (*string, bool)`

GetCustomproperty30LabelOk returns a tuple with the Customproperty30Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty30Label

`func (o *UpdateEndpointRequest) SetCustomproperty30Label(v string)`

SetCustomproperty30Label sets Customproperty30Label field to given value.

### HasCustomproperty30Label

`func (o *UpdateEndpointRequest) HasCustomproperty30Label() bool`

HasCustomproperty30Label returns a boolean if a field has been set.

### GetCustomproperty31Label

`func (o *UpdateEndpointRequest) GetCustomproperty31Label() string`

GetCustomproperty31Label returns the Customproperty31Label field if non-nil, zero value otherwise.

### GetCustomproperty31LabelOk

`func (o *UpdateEndpointRequest) GetCustomproperty31LabelOk() (*string, bool)`

GetCustomproperty31LabelOk returns a tuple with the Customproperty31Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty31Label

`func (o *UpdateEndpointRequest) SetCustomproperty31Label(v string)`

SetCustomproperty31Label sets Customproperty31Label field to given value.

### HasCustomproperty31Label

`func (o *UpdateEndpointRequest) HasCustomproperty31Label() bool`

HasCustomproperty31Label returns a boolean if a field has been set.

### GetCustomproperty32Label

`func (o *UpdateEndpointRequest) GetCustomproperty32Label() string`

GetCustomproperty32Label returns the Customproperty32Label field if non-nil, zero value otherwise.

### GetCustomproperty32LabelOk

`func (o *UpdateEndpointRequest) GetCustomproperty32LabelOk() (*string, bool)`

GetCustomproperty32LabelOk returns a tuple with the Customproperty32Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty32Label

`func (o *UpdateEndpointRequest) SetCustomproperty32Label(v string)`

SetCustomproperty32Label sets Customproperty32Label field to given value.

### HasCustomproperty32Label

`func (o *UpdateEndpointRequest) HasCustomproperty32Label() bool`

HasCustomproperty32Label returns a boolean if a field has been set.

### GetCustomproperty33Label

`func (o *UpdateEndpointRequest) GetCustomproperty33Label() string`

GetCustomproperty33Label returns the Customproperty33Label field if non-nil, zero value otherwise.

### GetCustomproperty33LabelOk

`func (o *UpdateEndpointRequest) GetCustomproperty33LabelOk() (*string, bool)`

GetCustomproperty33LabelOk returns a tuple with the Customproperty33Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty33Label

`func (o *UpdateEndpointRequest) SetCustomproperty33Label(v string)`

SetCustomproperty33Label sets Customproperty33Label field to given value.

### HasCustomproperty33Label

`func (o *UpdateEndpointRequest) HasCustomproperty33Label() bool`

HasCustomproperty33Label returns a boolean if a field has been set.

### GetCustomproperty34Label

`func (o *UpdateEndpointRequest) GetCustomproperty34Label() string`

GetCustomproperty34Label returns the Customproperty34Label field if non-nil, zero value otherwise.

### GetCustomproperty34LabelOk

`func (o *UpdateEndpointRequest) GetCustomproperty34LabelOk() (*string, bool)`

GetCustomproperty34LabelOk returns a tuple with the Customproperty34Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty34Label

`func (o *UpdateEndpointRequest) SetCustomproperty34Label(v string)`

SetCustomproperty34Label sets Customproperty34Label field to given value.

### HasCustomproperty34Label

`func (o *UpdateEndpointRequest) HasCustomproperty34Label() bool`

HasCustomproperty34Label returns a boolean if a field has been set.

### GetCustomproperty35Label

`func (o *UpdateEndpointRequest) GetCustomproperty35Label() string`

GetCustomproperty35Label returns the Customproperty35Label field if non-nil, zero value otherwise.

### GetCustomproperty35LabelOk

`func (o *UpdateEndpointRequest) GetCustomproperty35LabelOk() (*string, bool)`

GetCustomproperty35LabelOk returns a tuple with the Customproperty35Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty35Label

`func (o *UpdateEndpointRequest) SetCustomproperty35Label(v string)`

SetCustomproperty35Label sets Customproperty35Label field to given value.

### HasCustomproperty35Label

`func (o *UpdateEndpointRequest) HasCustomproperty35Label() bool`

HasCustomproperty35Label returns a boolean if a field has been set.

### GetCustomproperty36Label

`func (o *UpdateEndpointRequest) GetCustomproperty36Label() string`

GetCustomproperty36Label returns the Customproperty36Label field if non-nil, zero value otherwise.

### GetCustomproperty36LabelOk

`func (o *UpdateEndpointRequest) GetCustomproperty36LabelOk() (*string, bool)`

GetCustomproperty36LabelOk returns a tuple with the Customproperty36Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty36Label

`func (o *UpdateEndpointRequest) SetCustomproperty36Label(v string)`

SetCustomproperty36Label sets Customproperty36Label field to given value.

### HasCustomproperty36Label

`func (o *UpdateEndpointRequest) HasCustomproperty36Label() bool`

HasCustomproperty36Label returns a boolean if a field has been set.

### GetCustomproperty37Label

`func (o *UpdateEndpointRequest) GetCustomproperty37Label() string`

GetCustomproperty37Label returns the Customproperty37Label field if non-nil, zero value otherwise.

### GetCustomproperty37LabelOk

`func (o *UpdateEndpointRequest) GetCustomproperty37LabelOk() (*string, bool)`

GetCustomproperty37LabelOk returns a tuple with the Customproperty37Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty37Label

`func (o *UpdateEndpointRequest) SetCustomproperty37Label(v string)`

SetCustomproperty37Label sets Customproperty37Label field to given value.

### HasCustomproperty37Label

`func (o *UpdateEndpointRequest) HasCustomproperty37Label() bool`

HasCustomproperty37Label returns a boolean if a field has been set.

### GetCustomproperty38Label

`func (o *UpdateEndpointRequest) GetCustomproperty38Label() string`

GetCustomproperty38Label returns the Customproperty38Label field if non-nil, zero value otherwise.

### GetCustomproperty38LabelOk

`func (o *UpdateEndpointRequest) GetCustomproperty38LabelOk() (*string, bool)`

GetCustomproperty38LabelOk returns a tuple with the Customproperty38Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty38Label

`func (o *UpdateEndpointRequest) SetCustomproperty38Label(v string)`

SetCustomproperty38Label sets Customproperty38Label field to given value.

### HasCustomproperty38Label

`func (o *UpdateEndpointRequest) HasCustomproperty38Label() bool`

HasCustomproperty38Label returns a boolean if a field has been set.

### GetCustomproperty39Label

`func (o *UpdateEndpointRequest) GetCustomproperty39Label() string`

GetCustomproperty39Label returns the Customproperty39Label field if non-nil, zero value otherwise.

### GetCustomproperty39LabelOk

`func (o *UpdateEndpointRequest) GetCustomproperty39LabelOk() (*string, bool)`

GetCustomproperty39LabelOk returns a tuple with the Customproperty39Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty39Label

`func (o *UpdateEndpointRequest) SetCustomproperty39Label(v string)`

SetCustomproperty39Label sets Customproperty39Label field to given value.

### HasCustomproperty39Label

`func (o *UpdateEndpointRequest) HasCustomproperty39Label() bool`

HasCustomproperty39Label returns a boolean if a field has been set.

### GetCustomproperty40Label

`func (o *UpdateEndpointRequest) GetCustomproperty40Label() string`

GetCustomproperty40Label returns the Customproperty40Label field if non-nil, zero value otherwise.

### GetCustomproperty40LabelOk

`func (o *UpdateEndpointRequest) GetCustomproperty40LabelOk() (*string, bool)`

GetCustomproperty40LabelOk returns a tuple with the Customproperty40Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty40Label

`func (o *UpdateEndpointRequest) SetCustomproperty40Label(v string)`

SetCustomproperty40Label sets Customproperty40Label field to given value.

### HasCustomproperty40Label

`func (o *UpdateEndpointRequest) HasCustomproperty40Label() bool`

HasCustomproperty40Label returns a boolean if a field has been set.

### GetCustomproperty41Label

`func (o *UpdateEndpointRequest) GetCustomproperty41Label() string`

GetCustomproperty41Label returns the Customproperty41Label field if non-nil, zero value otherwise.

### GetCustomproperty41LabelOk

`func (o *UpdateEndpointRequest) GetCustomproperty41LabelOk() (*string, bool)`

GetCustomproperty41LabelOk returns a tuple with the Customproperty41Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty41Label

`func (o *UpdateEndpointRequest) SetCustomproperty41Label(v string)`

SetCustomproperty41Label sets Customproperty41Label field to given value.

### HasCustomproperty41Label

`func (o *UpdateEndpointRequest) HasCustomproperty41Label() bool`

HasCustomproperty41Label returns a boolean if a field has been set.

### GetCustomproperty42Label

`func (o *UpdateEndpointRequest) GetCustomproperty42Label() string`

GetCustomproperty42Label returns the Customproperty42Label field if non-nil, zero value otherwise.

### GetCustomproperty42LabelOk

`func (o *UpdateEndpointRequest) GetCustomproperty42LabelOk() (*string, bool)`

GetCustomproperty42LabelOk returns a tuple with the Customproperty42Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty42Label

`func (o *UpdateEndpointRequest) SetCustomproperty42Label(v string)`

SetCustomproperty42Label sets Customproperty42Label field to given value.

### HasCustomproperty42Label

`func (o *UpdateEndpointRequest) HasCustomproperty42Label() bool`

HasCustomproperty42Label returns a boolean if a field has been set.

### GetCustomproperty43Label

`func (o *UpdateEndpointRequest) GetCustomproperty43Label() string`

GetCustomproperty43Label returns the Customproperty43Label field if non-nil, zero value otherwise.

### GetCustomproperty43LabelOk

`func (o *UpdateEndpointRequest) GetCustomproperty43LabelOk() (*string, bool)`

GetCustomproperty43LabelOk returns a tuple with the Customproperty43Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty43Label

`func (o *UpdateEndpointRequest) SetCustomproperty43Label(v string)`

SetCustomproperty43Label sets Customproperty43Label field to given value.

### HasCustomproperty43Label

`func (o *UpdateEndpointRequest) HasCustomproperty43Label() bool`

HasCustomproperty43Label returns a boolean if a field has been set.

### GetCustomproperty44Label

`func (o *UpdateEndpointRequest) GetCustomproperty44Label() string`

GetCustomproperty44Label returns the Customproperty44Label field if non-nil, zero value otherwise.

### GetCustomproperty44LabelOk

`func (o *UpdateEndpointRequest) GetCustomproperty44LabelOk() (*string, bool)`

GetCustomproperty44LabelOk returns a tuple with the Customproperty44Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty44Label

`func (o *UpdateEndpointRequest) SetCustomproperty44Label(v string)`

SetCustomproperty44Label sets Customproperty44Label field to given value.

### HasCustomproperty44Label

`func (o *UpdateEndpointRequest) HasCustomproperty44Label() bool`

HasCustomproperty44Label returns a boolean if a field has been set.

### GetCustomproperty45Label

`func (o *UpdateEndpointRequest) GetCustomproperty45Label() string`

GetCustomproperty45Label returns the Customproperty45Label field if non-nil, zero value otherwise.

### GetCustomproperty45LabelOk

`func (o *UpdateEndpointRequest) GetCustomproperty45LabelOk() (*string, bool)`

GetCustomproperty45LabelOk returns a tuple with the Customproperty45Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty45Label

`func (o *UpdateEndpointRequest) SetCustomproperty45Label(v string)`

SetCustomproperty45Label sets Customproperty45Label field to given value.

### HasCustomproperty45Label

`func (o *UpdateEndpointRequest) HasCustomproperty45Label() bool`

HasCustomproperty45Label returns a boolean if a field has been set.

### GetCustomproperty46Label

`func (o *UpdateEndpointRequest) GetCustomproperty46Label() string`

GetCustomproperty46Label returns the Customproperty46Label field if non-nil, zero value otherwise.

### GetCustomproperty46LabelOk

`func (o *UpdateEndpointRequest) GetCustomproperty46LabelOk() (*string, bool)`

GetCustomproperty46LabelOk returns a tuple with the Customproperty46Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty46Label

`func (o *UpdateEndpointRequest) SetCustomproperty46Label(v string)`

SetCustomproperty46Label sets Customproperty46Label field to given value.

### HasCustomproperty46Label

`func (o *UpdateEndpointRequest) HasCustomproperty46Label() bool`

HasCustomproperty46Label returns a boolean if a field has been set.

### GetCustomproperty47Label

`func (o *UpdateEndpointRequest) GetCustomproperty47Label() string`

GetCustomproperty47Label returns the Customproperty47Label field if non-nil, zero value otherwise.

### GetCustomproperty47LabelOk

`func (o *UpdateEndpointRequest) GetCustomproperty47LabelOk() (*string, bool)`

GetCustomproperty47LabelOk returns a tuple with the Customproperty47Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty47Label

`func (o *UpdateEndpointRequest) SetCustomproperty47Label(v string)`

SetCustomproperty47Label sets Customproperty47Label field to given value.

### HasCustomproperty47Label

`func (o *UpdateEndpointRequest) HasCustomproperty47Label() bool`

HasCustomproperty47Label returns a boolean if a field has been set.

### GetCustomproperty48Label

`func (o *UpdateEndpointRequest) GetCustomproperty48Label() string`

GetCustomproperty48Label returns the Customproperty48Label field if non-nil, zero value otherwise.

### GetCustomproperty48LabelOk

`func (o *UpdateEndpointRequest) GetCustomproperty48LabelOk() (*string, bool)`

GetCustomproperty48LabelOk returns a tuple with the Customproperty48Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty48Label

`func (o *UpdateEndpointRequest) SetCustomproperty48Label(v string)`

SetCustomproperty48Label sets Customproperty48Label field to given value.

### HasCustomproperty48Label

`func (o *UpdateEndpointRequest) HasCustomproperty48Label() bool`

HasCustomproperty48Label returns a boolean if a field has been set.

### GetCustomproperty49Label

`func (o *UpdateEndpointRequest) GetCustomproperty49Label() string`

GetCustomproperty49Label returns the Customproperty49Label field if non-nil, zero value otherwise.

### GetCustomproperty49LabelOk

`func (o *UpdateEndpointRequest) GetCustomproperty49LabelOk() (*string, bool)`

GetCustomproperty49LabelOk returns a tuple with the Customproperty49Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty49Label

`func (o *UpdateEndpointRequest) SetCustomproperty49Label(v string)`

SetCustomproperty49Label sets Customproperty49Label field to given value.

### HasCustomproperty49Label

`func (o *UpdateEndpointRequest) HasCustomproperty49Label() bool`

HasCustomproperty49Label returns a boolean if a field has been set.

### GetCustomproperty50Label

`func (o *UpdateEndpointRequest) GetCustomproperty50Label() string`

GetCustomproperty50Label returns the Customproperty50Label field if non-nil, zero value otherwise.

### GetCustomproperty50LabelOk

`func (o *UpdateEndpointRequest) GetCustomproperty50LabelOk() (*string, bool)`

GetCustomproperty50LabelOk returns a tuple with the Customproperty50Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty50Label

`func (o *UpdateEndpointRequest) SetCustomproperty50Label(v string)`

SetCustomproperty50Label sets Customproperty50Label field to given value.

### HasCustomproperty50Label

`func (o *UpdateEndpointRequest) HasCustomproperty50Label() bool`

HasCustomproperty50Label returns a boolean if a field has been set.

### GetCustomproperty51Label

`func (o *UpdateEndpointRequest) GetCustomproperty51Label() string`

GetCustomproperty51Label returns the Customproperty51Label field if non-nil, zero value otherwise.

### GetCustomproperty51LabelOk

`func (o *UpdateEndpointRequest) GetCustomproperty51LabelOk() (*string, bool)`

GetCustomproperty51LabelOk returns a tuple with the Customproperty51Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty51Label

`func (o *UpdateEndpointRequest) SetCustomproperty51Label(v string)`

SetCustomproperty51Label sets Customproperty51Label field to given value.

### HasCustomproperty51Label

`func (o *UpdateEndpointRequest) HasCustomproperty51Label() bool`

HasCustomproperty51Label returns a boolean if a field has been set.

### GetCustomproperty52Label

`func (o *UpdateEndpointRequest) GetCustomproperty52Label() string`

GetCustomproperty52Label returns the Customproperty52Label field if non-nil, zero value otherwise.

### GetCustomproperty52LabelOk

`func (o *UpdateEndpointRequest) GetCustomproperty52LabelOk() (*string, bool)`

GetCustomproperty52LabelOk returns a tuple with the Customproperty52Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty52Label

`func (o *UpdateEndpointRequest) SetCustomproperty52Label(v string)`

SetCustomproperty52Label sets Customproperty52Label field to given value.

### HasCustomproperty52Label

`func (o *UpdateEndpointRequest) HasCustomproperty52Label() bool`

HasCustomproperty52Label returns a boolean if a field has been set.

### GetCustomproperty53Label

`func (o *UpdateEndpointRequest) GetCustomproperty53Label() string`

GetCustomproperty53Label returns the Customproperty53Label field if non-nil, zero value otherwise.

### GetCustomproperty53LabelOk

`func (o *UpdateEndpointRequest) GetCustomproperty53LabelOk() (*string, bool)`

GetCustomproperty53LabelOk returns a tuple with the Customproperty53Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty53Label

`func (o *UpdateEndpointRequest) SetCustomproperty53Label(v string)`

SetCustomproperty53Label sets Customproperty53Label field to given value.

### HasCustomproperty53Label

`func (o *UpdateEndpointRequest) HasCustomproperty53Label() bool`

HasCustomproperty53Label returns a boolean if a field has been set.

### GetCustomproperty54Label

`func (o *UpdateEndpointRequest) GetCustomproperty54Label() string`

GetCustomproperty54Label returns the Customproperty54Label field if non-nil, zero value otherwise.

### GetCustomproperty54LabelOk

`func (o *UpdateEndpointRequest) GetCustomproperty54LabelOk() (*string, bool)`

GetCustomproperty54LabelOk returns a tuple with the Customproperty54Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty54Label

`func (o *UpdateEndpointRequest) SetCustomproperty54Label(v string)`

SetCustomproperty54Label sets Customproperty54Label field to given value.

### HasCustomproperty54Label

`func (o *UpdateEndpointRequest) HasCustomproperty54Label() bool`

HasCustomproperty54Label returns a boolean if a field has been set.

### GetCustomproperty55Label

`func (o *UpdateEndpointRequest) GetCustomproperty55Label() string`

GetCustomproperty55Label returns the Customproperty55Label field if non-nil, zero value otherwise.

### GetCustomproperty55LabelOk

`func (o *UpdateEndpointRequest) GetCustomproperty55LabelOk() (*string, bool)`

GetCustomproperty55LabelOk returns a tuple with the Customproperty55Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty55Label

`func (o *UpdateEndpointRequest) SetCustomproperty55Label(v string)`

SetCustomproperty55Label sets Customproperty55Label field to given value.

### HasCustomproperty55Label

`func (o *UpdateEndpointRequest) HasCustomproperty55Label() bool`

HasCustomproperty55Label returns a boolean if a field has been set.

### GetCustomproperty56Label

`func (o *UpdateEndpointRequest) GetCustomproperty56Label() string`

GetCustomproperty56Label returns the Customproperty56Label field if non-nil, zero value otherwise.

### GetCustomproperty56LabelOk

`func (o *UpdateEndpointRequest) GetCustomproperty56LabelOk() (*string, bool)`

GetCustomproperty56LabelOk returns a tuple with the Customproperty56Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty56Label

`func (o *UpdateEndpointRequest) SetCustomproperty56Label(v string)`

SetCustomproperty56Label sets Customproperty56Label field to given value.

### HasCustomproperty56Label

`func (o *UpdateEndpointRequest) HasCustomproperty56Label() bool`

HasCustomproperty56Label returns a boolean if a field has been set.

### GetCustomproperty57Label

`func (o *UpdateEndpointRequest) GetCustomproperty57Label() string`

GetCustomproperty57Label returns the Customproperty57Label field if non-nil, zero value otherwise.

### GetCustomproperty57LabelOk

`func (o *UpdateEndpointRequest) GetCustomproperty57LabelOk() (*string, bool)`

GetCustomproperty57LabelOk returns a tuple with the Customproperty57Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty57Label

`func (o *UpdateEndpointRequest) SetCustomproperty57Label(v string)`

SetCustomproperty57Label sets Customproperty57Label field to given value.

### HasCustomproperty57Label

`func (o *UpdateEndpointRequest) HasCustomproperty57Label() bool`

HasCustomproperty57Label returns a boolean if a field has been set.

### GetCustomproperty58Label

`func (o *UpdateEndpointRequest) GetCustomproperty58Label() string`

GetCustomproperty58Label returns the Customproperty58Label field if non-nil, zero value otherwise.

### GetCustomproperty58LabelOk

`func (o *UpdateEndpointRequest) GetCustomproperty58LabelOk() (*string, bool)`

GetCustomproperty58LabelOk returns a tuple with the Customproperty58Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty58Label

`func (o *UpdateEndpointRequest) SetCustomproperty58Label(v string)`

SetCustomproperty58Label sets Customproperty58Label field to given value.

### HasCustomproperty58Label

`func (o *UpdateEndpointRequest) HasCustomproperty58Label() bool`

HasCustomproperty58Label returns a boolean if a field has been set.

### GetCustomproperty59Label

`func (o *UpdateEndpointRequest) GetCustomproperty59Label() string`

GetCustomproperty59Label returns the Customproperty59Label field if non-nil, zero value otherwise.

### GetCustomproperty59LabelOk

`func (o *UpdateEndpointRequest) GetCustomproperty59LabelOk() (*string, bool)`

GetCustomproperty59LabelOk returns a tuple with the Customproperty59Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty59Label

`func (o *UpdateEndpointRequest) SetCustomproperty59Label(v string)`

SetCustomproperty59Label sets Customproperty59Label field to given value.

### HasCustomproperty59Label

`func (o *UpdateEndpointRequest) HasCustomproperty59Label() bool`

HasCustomproperty59Label returns a boolean if a field has been set.

### GetCustomproperty60Label

`func (o *UpdateEndpointRequest) GetCustomproperty60Label() string`

GetCustomproperty60Label returns the Customproperty60Label field if non-nil, zero value otherwise.

### GetCustomproperty60LabelOk

`func (o *UpdateEndpointRequest) GetCustomproperty60LabelOk() (*string, bool)`

GetCustomproperty60LabelOk returns a tuple with the Customproperty60Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomproperty60Label

`func (o *UpdateEndpointRequest) SetCustomproperty60Label(v string)`

SetCustomproperty60Label sets Customproperty60Label field to given value.

### HasCustomproperty60Label

`func (o *UpdateEndpointRequest) HasCustomproperty60Label() bool`

HasCustomproperty60Label returns a boolean if a field has been set.

### GetAllowRemoveAllRoleOnRequest

`func (o *UpdateEndpointRequest) GetAllowRemoveAllRoleOnRequest() string`

GetAllowRemoveAllRoleOnRequest returns the AllowRemoveAllRoleOnRequest field if non-nil, zero value otherwise.

### GetAllowRemoveAllRoleOnRequestOk

`func (o *UpdateEndpointRequest) GetAllowRemoveAllRoleOnRequestOk() (*string, bool)`

GetAllowRemoveAllRoleOnRequestOk returns a tuple with the AllowRemoveAllRoleOnRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowRemoveAllRoleOnRequest

`func (o *UpdateEndpointRequest) SetAllowRemoveAllRoleOnRequest(v string)`

SetAllowRemoveAllRoleOnRequest sets AllowRemoveAllRoleOnRequest field to given value.

### HasAllowRemoveAllRoleOnRequest

`func (o *UpdateEndpointRequest) HasAllowRemoveAllRoleOnRequest() bool`

HasAllowRemoveAllRoleOnRequest returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


