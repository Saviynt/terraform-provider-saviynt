# EntraIDConnectionAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UpdateUserJSON** | Pointer to **string** |  | [optional] 
**MICROSOFT_GRAPH_ENDPOINT** | Pointer to **string** |  | [optional] 
**ENDPOINTS_FILTER** | Pointer to **string** |  | [optional] 
**ImportUserJSON** | Pointer to **string** |  | [optional] 
**ConnectionType** | Pointer to **string** |  | [optional] 
**EnableAccountJSON** | Pointer to **string** |  | [optional] 
**ConnectionJSON** | Pointer to **string** |  | [optional] 
**CLIENT_ID** | Pointer to **string** |  | [optional] 
**DeleteGroupJSON** | Pointer to **string** |  | [optional] 
**ConfigJSON** | Pointer to **string** |  | [optional] 
**ACCESS_TOKEN** | Pointer to **string** |  | [optional] 
**AddAccessJSON** | Pointer to **string** |  | [optional] 
**CreateChannelJSON** | Pointer to **string** |  | [optional] 
**UpdateAccountJSON** | Pointer to **string** |  | [optional] 
**IsTimeoutSupported** | Pointer to **bool** |  | [optional] 
**RemoveServicePrincipalJSON** | Pointer to **string** |  | [optional] 
**IMPORT_DEPTH** | Pointer to **string** |  | [optional] 
**CreateAccountJSON** | Pointer to **string** |  | [optional] 
**PAM_CONFIG** | Pointer to **string** |  | [optional] 
**UpdateServicePrincipalJSON** | Pointer to **string** |  | [optional] 
**AZURE_MANAGEMENT_ENDPOINT** | Pointer to **string** |  | [optional] 
**ENTITLEMENT_ATTRIBUTE** | Pointer to **string** |  | [optional] 
**ACCOUNTS_FILTER** | Pointer to **string** |  | [optional] 
**WINDOWS_CONNECTOR_JSON** | Pointer to **string** |  | [optional] 
**DELTATOKENSJSON** | Pointer to **string** |  | [optional] 
**AZURE_MGMT_ACCESS_TOKEN** | Pointer to **string** |  | [optional] 
**CreateTeamJSON** | Pointer to **string** |  | [optional] 
**ConnectionTimeoutConfig** | Pointer to [**ConnectionTimeoutConfig**](ConnectionTimeoutConfig.md) |  | [optional] 
**ENHANCEDDIRECTORYROLES** | Pointer to **string** |  | [optional] 
**STATUS_THRESHOLD_CONFIG** | Pointer to **string** |  | [optional] 
**ACCOUNT_IMPORT_FIELDS** | Pointer to **string** |  | [optional] 
**RemoveAccountJSON** | Pointer to **string** |  | [optional] 
**ChangePassJSON** | Pointer to **string** |  | [optional] 
**CLIENT_SECRET** | Pointer to **string** |  | [optional] 
**ENTITLEMENT_FILTER_JSON** | Pointer to **string** |  | [optional] 
**SERVICE_ACCOUNT_ATTRIBUTES** | Pointer to **string** |  | [optional] 
**AddAccessToEntitlementJSON** | Pointer to **string** |  | [optional] 
**AUTHENTICATION_ENDPOINT** | Pointer to **string** |  | [optional] 
**CreateServicePrincipalJSON** | Pointer to **string** |  | [optional] 
**MODIFYUSERDATAJSON** | Pointer to **string** |  | [optional] 
**IsTimeoutConfigValidated** | Pointer to **bool** |  | [optional] 
**RemoveAccessJSON** | Pointer to **string** |  | [optional] 
**CREATEUSERS** | Pointer to **string** |  | [optional] 
**RemoveAccessFromEntitlementJSON** | Pointer to **string** |  | [optional] 
**DisableAccountJSON** | Pointer to **string** |  | [optional] 
**CREATE_NEW_ENDPOINTS** | Pointer to **string** |  | [optional] 
**MANAGED_ACCOUNT_TYPE** | Pointer to **string** |  | [optional] 
**ACCOUNT_ATTRIBUTES** | Pointer to **string** |  | [optional] 
**AAD_TENANT_ID** | Pointer to **string** |  | [optional] 
**UpdateGroupJSON** | Pointer to **string** |  | [optional] 
**CreateGroupJSON** | Pointer to **string** |  | [optional] 

## Methods

### NewEntraIDConnectionAttributes

`func NewEntraIDConnectionAttributes() *EntraIDConnectionAttributes`

NewEntraIDConnectionAttributes instantiates a new EntraIDConnectionAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEntraIDConnectionAttributesWithDefaults

`func NewEntraIDConnectionAttributesWithDefaults() *EntraIDConnectionAttributes`

NewEntraIDConnectionAttributesWithDefaults instantiates a new EntraIDConnectionAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUpdateUserJSON

`func (o *EntraIDConnectionAttributes) GetUpdateUserJSON() string`

GetUpdateUserJSON returns the UpdateUserJSON field if non-nil, zero value otherwise.

### GetUpdateUserJSONOk

`func (o *EntraIDConnectionAttributes) GetUpdateUserJSONOk() (*string, bool)`

GetUpdateUserJSONOk returns a tuple with the UpdateUserJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateUserJSON

`func (o *EntraIDConnectionAttributes) SetUpdateUserJSON(v string)`

SetUpdateUserJSON sets UpdateUserJSON field to given value.

### HasUpdateUserJSON

`func (o *EntraIDConnectionAttributes) HasUpdateUserJSON() bool`

HasUpdateUserJSON returns a boolean if a field has been set.

### GetMICROSOFT_GRAPH_ENDPOINT

`func (o *EntraIDConnectionAttributes) GetMICROSOFT_GRAPH_ENDPOINT() string`

GetMICROSOFT_GRAPH_ENDPOINT returns the MICROSOFT_GRAPH_ENDPOINT field if non-nil, zero value otherwise.

### GetMICROSOFT_GRAPH_ENDPOINTOk

`func (o *EntraIDConnectionAttributes) GetMICROSOFT_GRAPH_ENDPOINTOk() (*string, bool)`

GetMICROSOFT_GRAPH_ENDPOINTOk returns a tuple with the MICROSOFT_GRAPH_ENDPOINT field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMICROSOFT_GRAPH_ENDPOINT

`func (o *EntraIDConnectionAttributes) SetMICROSOFT_GRAPH_ENDPOINT(v string)`

SetMICROSOFT_GRAPH_ENDPOINT sets MICROSOFT_GRAPH_ENDPOINT field to given value.

### HasMICROSOFT_GRAPH_ENDPOINT

`func (o *EntraIDConnectionAttributes) HasMICROSOFT_GRAPH_ENDPOINT() bool`

HasMICROSOFT_GRAPH_ENDPOINT returns a boolean if a field has been set.

### GetENDPOINTS_FILTER

`func (o *EntraIDConnectionAttributes) GetENDPOINTS_FILTER() string`

GetENDPOINTS_FILTER returns the ENDPOINTS_FILTER field if non-nil, zero value otherwise.

### GetENDPOINTS_FILTEROk

`func (o *EntraIDConnectionAttributes) GetENDPOINTS_FILTEROk() (*string, bool)`

GetENDPOINTS_FILTEROk returns a tuple with the ENDPOINTS_FILTER field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetENDPOINTS_FILTER

`func (o *EntraIDConnectionAttributes) SetENDPOINTS_FILTER(v string)`

SetENDPOINTS_FILTER sets ENDPOINTS_FILTER field to given value.

### HasENDPOINTS_FILTER

`func (o *EntraIDConnectionAttributes) HasENDPOINTS_FILTER() bool`

HasENDPOINTS_FILTER returns a boolean if a field has been set.

### GetImportUserJSON

`func (o *EntraIDConnectionAttributes) GetImportUserJSON() string`

GetImportUserJSON returns the ImportUserJSON field if non-nil, zero value otherwise.

### GetImportUserJSONOk

`func (o *EntraIDConnectionAttributes) GetImportUserJSONOk() (*string, bool)`

GetImportUserJSONOk returns a tuple with the ImportUserJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportUserJSON

`func (o *EntraIDConnectionAttributes) SetImportUserJSON(v string)`

SetImportUserJSON sets ImportUserJSON field to given value.

### HasImportUserJSON

`func (o *EntraIDConnectionAttributes) HasImportUserJSON() bool`

HasImportUserJSON returns a boolean if a field has been set.

### GetConnectionType

`func (o *EntraIDConnectionAttributes) GetConnectionType() string`

GetConnectionType returns the ConnectionType field if non-nil, zero value otherwise.

### GetConnectionTypeOk

`func (o *EntraIDConnectionAttributes) GetConnectionTypeOk() (*string, bool)`

GetConnectionTypeOk returns a tuple with the ConnectionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionType

`func (o *EntraIDConnectionAttributes) SetConnectionType(v string)`

SetConnectionType sets ConnectionType field to given value.

### HasConnectionType

`func (o *EntraIDConnectionAttributes) HasConnectionType() bool`

HasConnectionType returns a boolean if a field has been set.

### GetEnableAccountJSON

`func (o *EntraIDConnectionAttributes) GetEnableAccountJSON() string`

GetEnableAccountJSON returns the EnableAccountJSON field if non-nil, zero value otherwise.

### GetEnableAccountJSONOk

`func (o *EntraIDConnectionAttributes) GetEnableAccountJSONOk() (*string, bool)`

GetEnableAccountJSONOk returns a tuple with the EnableAccountJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableAccountJSON

`func (o *EntraIDConnectionAttributes) SetEnableAccountJSON(v string)`

SetEnableAccountJSON sets EnableAccountJSON field to given value.

### HasEnableAccountJSON

`func (o *EntraIDConnectionAttributes) HasEnableAccountJSON() bool`

HasEnableAccountJSON returns a boolean if a field has been set.

### GetConnectionJSON

`func (o *EntraIDConnectionAttributes) GetConnectionJSON() string`

GetConnectionJSON returns the ConnectionJSON field if non-nil, zero value otherwise.

### GetConnectionJSONOk

`func (o *EntraIDConnectionAttributes) GetConnectionJSONOk() (*string, bool)`

GetConnectionJSONOk returns a tuple with the ConnectionJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionJSON

`func (o *EntraIDConnectionAttributes) SetConnectionJSON(v string)`

SetConnectionJSON sets ConnectionJSON field to given value.

### HasConnectionJSON

`func (o *EntraIDConnectionAttributes) HasConnectionJSON() bool`

HasConnectionJSON returns a boolean if a field has been set.

### GetCLIENT_ID

`func (o *EntraIDConnectionAttributes) GetCLIENT_ID() string`

GetCLIENT_ID returns the CLIENT_ID field if non-nil, zero value otherwise.

### GetCLIENT_IDOk

`func (o *EntraIDConnectionAttributes) GetCLIENT_IDOk() (*string, bool)`

GetCLIENT_IDOk returns a tuple with the CLIENT_ID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCLIENT_ID

`func (o *EntraIDConnectionAttributes) SetCLIENT_ID(v string)`

SetCLIENT_ID sets CLIENT_ID field to given value.

### HasCLIENT_ID

`func (o *EntraIDConnectionAttributes) HasCLIENT_ID() bool`

HasCLIENT_ID returns a boolean if a field has been set.

### GetDeleteGroupJSON

`func (o *EntraIDConnectionAttributes) GetDeleteGroupJSON() string`

GetDeleteGroupJSON returns the DeleteGroupJSON field if non-nil, zero value otherwise.

### GetDeleteGroupJSONOk

`func (o *EntraIDConnectionAttributes) GetDeleteGroupJSONOk() (*string, bool)`

GetDeleteGroupJSONOk returns a tuple with the DeleteGroupJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteGroupJSON

`func (o *EntraIDConnectionAttributes) SetDeleteGroupJSON(v string)`

SetDeleteGroupJSON sets DeleteGroupJSON field to given value.

### HasDeleteGroupJSON

`func (o *EntraIDConnectionAttributes) HasDeleteGroupJSON() bool`

HasDeleteGroupJSON returns a boolean if a field has been set.

### GetConfigJSON

`func (o *EntraIDConnectionAttributes) GetConfigJSON() string`

GetConfigJSON returns the ConfigJSON field if non-nil, zero value otherwise.

### GetConfigJSONOk

`func (o *EntraIDConnectionAttributes) GetConfigJSONOk() (*string, bool)`

GetConfigJSONOk returns a tuple with the ConfigJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigJSON

`func (o *EntraIDConnectionAttributes) SetConfigJSON(v string)`

SetConfigJSON sets ConfigJSON field to given value.

### HasConfigJSON

`func (o *EntraIDConnectionAttributes) HasConfigJSON() bool`

HasConfigJSON returns a boolean if a field has been set.

### GetACCESS_TOKEN

`func (o *EntraIDConnectionAttributes) GetACCESS_TOKEN() string`

GetACCESS_TOKEN returns the ACCESS_TOKEN field if non-nil, zero value otherwise.

### GetACCESS_TOKENOk

`func (o *EntraIDConnectionAttributes) GetACCESS_TOKENOk() (*string, bool)`

GetACCESS_TOKENOk returns a tuple with the ACCESS_TOKEN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetACCESS_TOKEN

`func (o *EntraIDConnectionAttributes) SetACCESS_TOKEN(v string)`

SetACCESS_TOKEN sets ACCESS_TOKEN field to given value.

### HasACCESS_TOKEN

`func (o *EntraIDConnectionAttributes) HasACCESS_TOKEN() bool`

HasACCESS_TOKEN returns a boolean if a field has been set.

### GetAddAccessJSON

`func (o *EntraIDConnectionAttributes) GetAddAccessJSON() string`

GetAddAccessJSON returns the AddAccessJSON field if non-nil, zero value otherwise.

### GetAddAccessJSONOk

`func (o *EntraIDConnectionAttributes) GetAddAccessJSONOk() (*string, bool)`

GetAddAccessJSONOk returns a tuple with the AddAccessJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddAccessJSON

`func (o *EntraIDConnectionAttributes) SetAddAccessJSON(v string)`

SetAddAccessJSON sets AddAccessJSON field to given value.

### HasAddAccessJSON

`func (o *EntraIDConnectionAttributes) HasAddAccessJSON() bool`

HasAddAccessJSON returns a boolean if a field has been set.

### GetCreateChannelJSON

`func (o *EntraIDConnectionAttributes) GetCreateChannelJSON() string`

GetCreateChannelJSON returns the CreateChannelJSON field if non-nil, zero value otherwise.

### GetCreateChannelJSONOk

`func (o *EntraIDConnectionAttributes) GetCreateChannelJSONOk() (*string, bool)`

GetCreateChannelJSONOk returns a tuple with the CreateChannelJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateChannelJSON

`func (o *EntraIDConnectionAttributes) SetCreateChannelJSON(v string)`

SetCreateChannelJSON sets CreateChannelJSON field to given value.

### HasCreateChannelJSON

`func (o *EntraIDConnectionAttributes) HasCreateChannelJSON() bool`

HasCreateChannelJSON returns a boolean if a field has been set.

### GetUpdateAccountJSON

`func (o *EntraIDConnectionAttributes) GetUpdateAccountJSON() string`

GetUpdateAccountJSON returns the UpdateAccountJSON field if non-nil, zero value otherwise.

### GetUpdateAccountJSONOk

`func (o *EntraIDConnectionAttributes) GetUpdateAccountJSONOk() (*string, bool)`

GetUpdateAccountJSONOk returns a tuple with the UpdateAccountJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateAccountJSON

`func (o *EntraIDConnectionAttributes) SetUpdateAccountJSON(v string)`

SetUpdateAccountJSON sets UpdateAccountJSON field to given value.

### HasUpdateAccountJSON

`func (o *EntraIDConnectionAttributes) HasUpdateAccountJSON() bool`

HasUpdateAccountJSON returns a boolean if a field has been set.

### GetIsTimeoutSupported

`func (o *EntraIDConnectionAttributes) GetIsTimeoutSupported() bool`

GetIsTimeoutSupported returns the IsTimeoutSupported field if non-nil, zero value otherwise.

### GetIsTimeoutSupportedOk

`func (o *EntraIDConnectionAttributes) GetIsTimeoutSupportedOk() (*bool, bool)`

GetIsTimeoutSupportedOk returns a tuple with the IsTimeoutSupported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsTimeoutSupported

`func (o *EntraIDConnectionAttributes) SetIsTimeoutSupported(v bool)`

SetIsTimeoutSupported sets IsTimeoutSupported field to given value.

### HasIsTimeoutSupported

`func (o *EntraIDConnectionAttributes) HasIsTimeoutSupported() bool`

HasIsTimeoutSupported returns a boolean if a field has been set.

### GetRemoveServicePrincipalJSON

`func (o *EntraIDConnectionAttributes) GetRemoveServicePrincipalJSON() string`

GetRemoveServicePrincipalJSON returns the RemoveServicePrincipalJSON field if non-nil, zero value otherwise.

### GetRemoveServicePrincipalJSONOk

`func (o *EntraIDConnectionAttributes) GetRemoveServicePrincipalJSONOk() (*string, bool)`

GetRemoveServicePrincipalJSONOk returns a tuple with the RemoveServicePrincipalJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoveServicePrincipalJSON

`func (o *EntraIDConnectionAttributes) SetRemoveServicePrincipalJSON(v string)`

SetRemoveServicePrincipalJSON sets RemoveServicePrincipalJSON field to given value.

### HasRemoveServicePrincipalJSON

`func (o *EntraIDConnectionAttributes) HasRemoveServicePrincipalJSON() bool`

HasRemoveServicePrincipalJSON returns a boolean if a field has been set.

### GetIMPORT_DEPTH

`func (o *EntraIDConnectionAttributes) GetIMPORT_DEPTH() string`

GetIMPORT_DEPTH returns the IMPORT_DEPTH field if non-nil, zero value otherwise.

### GetIMPORT_DEPTHOk

`func (o *EntraIDConnectionAttributes) GetIMPORT_DEPTHOk() (*string, bool)`

GetIMPORT_DEPTHOk returns a tuple with the IMPORT_DEPTH field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIMPORT_DEPTH

`func (o *EntraIDConnectionAttributes) SetIMPORT_DEPTH(v string)`

SetIMPORT_DEPTH sets IMPORT_DEPTH field to given value.

### HasIMPORT_DEPTH

`func (o *EntraIDConnectionAttributes) HasIMPORT_DEPTH() bool`

HasIMPORT_DEPTH returns a boolean if a field has been set.

### GetCreateAccountJSON

`func (o *EntraIDConnectionAttributes) GetCreateAccountJSON() string`

GetCreateAccountJSON returns the CreateAccountJSON field if non-nil, zero value otherwise.

### GetCreateAccountJSONOk

`func (o *EntraIDConnectionAttributes) GetCreateAccountJSONOk() (*string, bool)`

GetCreateAccountJSONOk returns a tuple with the CreateAccountJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateAccountJSON

`func (o *EntraIDConnectionAttributes) SetCreateAccountJSON(v string)`

SetCreateAccountJSON sets CreateAccountJSON field to given value.

### HasCreateAccountJSON

`func (o *EntraIDConnectionAttributes) HasCreateAccountJSON() bool`

HasCreateAccountJSON returns a boolean if a field has been set.

### GetPAM_CONFIG

`func (o *EntraIDConnectionAttributes) GetPAM_CONFIG() string`

GetPAM_CONFIG returns the PAM_CONFIG field if non-nil, zero value otherwise.

### GetPAM_CONFIGOk

`func (o *EntraIDConnectionAttributes) GetPAM_CONFIGOk() (*string, bool)`

GetPAM_CONFIGOk returns a tuple with the PAM_CONFIG field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPAM_CONFIG

`func (o *EntraIDConnectionAttributes) SetPAM_CONFIG(v string)`

SetPAM_CONFIG sets PAM_CONFIG field to given value.

### HasPAM_CONFIG

`func (o *EntraIDConnectionAttributes) HasPAM_CONFIG() bool`

HasPAM_CONFIG returns a boolean if a field has been set.

### GetUpdateServicePrincipalJSON

`func (o *EntraIDConnectionAttributes) GetUpdateServicePrincipalJSON() string`

GetUpdateServicePrincipalJSON returns the UpdateServicePrincipalJSON field if non-nil, zero value otherwise.

### GetUpdateServicePrincipalJSONOk

`func (o *EntraIDConnectionAttributes) GetUpdateServicePrincipalJSONOk() (*string, bool)`

GetUpdateServicePrincipalJSONOk returns a tuple with the UpdateServicePrincipalJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateServicePrincipalJSON

`func (o *EntraIDConnectionAttributes) SetUpdateServicePrincipalJSON(v string)`

SetUpdateServicePrincipalJSON sets UpdateServicePrincipalJSON field to given value.

### HasUpdateServicePrincipalJSON

`func (o *EntraIDConnectionAttributes) HasUpdateServicePrincipalJSON() bool`

HasUpdateServicePrincipalJSON returns a boolean if a field has been set.

### GetAZURE_MANAGEMENT_ENDPOINT

`func (o *EntraIDConnectionAttributes) GetAZURE_MANAGEMENT_ENDPOINT() string`

GetAZURE_MANAGEMENT_ENDPOINT returns the AZURE_MANAGEMENT_ENDPOINT field if non-nil, zero value otherwise.

### GetAZURE_MANAGEMENT_ENDPOINTOk

`func (o *EntraIDConnectionAttributes) GetAZURE_MANAGEMENT_ENDPOINTOk() (*string, bool)`

GetAZURE_MANAGEMENT_ENDPOINTOk returns a tuple with the AZURE_MANAGEMENT_ENDPOINT field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAZURE_MANAGEMENT_ENDPOINT

`func (o *EntraIDConnectionAttributes) SetAZURE_MANAGEMENT_ENDPOINT(v string)`

SetAZURE_MANAGEMENT_ENDPOINT sets AZURE_MANAGEMENT_ENDPOINT field to given value.

### HasAZURE_MANAGEMENT_ENDPOINT

`func (o *EntraIDConnectionAttributes) HasAZURE_MANAGEMENT_ENDPOINT() bool`

HasAZURE_MANAGEMENT_ENDPOINT returns a boolean if a field has been set.

### GetENTITLEMENT_ATTRIBUTE

`func (o *EntraIDConnectionAttributes) GetENTITLEMENT_ATTRIBUTE() string`

GetENTITLEMENT_ATTRIBUTE returns the ENTITLEMENT_ATTRIBUTE field if non-nil, zero value otherwise.

### GetENTITLEMENT_ATTRIBUTEOk

`func (o *EntraIDConnectionAttributes) GetENTITLEMENT_ATTRIBUTEOk() (*string, bool)`

GetENTITLEMENT_ATTRIBUTEOk returns a tuple with the ENTITLEMENT_ATTRIBUTE field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetENTITLEMENT_ATTRIBUTE

`func (o *EntraIDConnectionAttributes) SetENTITLEMENT_ATTRIBUTE(v string)`

SetENTITLEMENT_ATTRIBUTE sets ENTITLEMENT_ATTRIBUTE field to given value.

### HasENTITLEMENT_ATTRIBUTE

`func (o *EntraIDConnectionAttributes) HasENTITLEMENT_ATTRIBUTE() bool`

HasENTITLEMENT_ATTRIBUTE returns a boolean if a field has been set.

### GetACCOUNTS_FILTER

`func (o *EntraIDConnectionAttributes) GetACCOUNTS_FILTER() string`

GetACCOUNTS_FILTER returns the ACCOUNTS_FILTER field if non-nil, zero value otherwise.

### GetACCOUNTS_FILTEROk

`func (o *EntraIDConnectionAttributes) GetACCOUNTS_FILTEROk() (*string, bool)`

GetACCOUNTS_FILTEROk returns a tuple with the ACCOUNTS_FILTER field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetACCOUNTS_FILTER

`func (o *EntraIDConnectionAttributes) SetACCOUNTS_FILTER(v string)`

SetACCOUNTS_FILTER sets ACCOUNTS_FILTER field to given value.

### HasACCOUNTS_FILTER

`func (o *EntraIDConnectionAttributes) HasACCOUNTS_FILTER() bool`

HasACCOUNTS_FILTER returns a boolean if a field has been set.

### GetWINDOWS_CONNECTOR_JSON

`func (o *EntraIDConnectionAttributes) GetWINDOWS_CONNECTOR_JSON() string`

GetWINDOWS_CONNECTOR_JSON returns the WINDOWS_CONNECTOR_JSON field if non-nil, zero value otherwise.

### GetWINDOWS_CONNECTOR_JSONOk

`func (o *EntraIDConnectionAttributes) GetWINDOWS_CONNECTOR_JSONOk() (*string, bool)`

GetWINDOWS_CONNECTOR_JSONOk returns a tuple with the WINDOWS_CONNECTOR_JSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWINDOWS_CONNECTOR_JSON

`func (o *EntraIDConnectionAttributes) SetWINDOWS_CONNECTOR_JSON(v string)`

SetWINDOWS_CONNECTOR_JSON sets WINDOWS_CONNECTOR_JSON field to given value.

### HasWINDOWS_CONNECTOR_JSON

`func (o *EntraIDConnectionAttributes) HasWINDOWS_CONNECTOR_JSON() bool`

HasWINDOWS_CONNECTOR_JSON returns a boolean if a field has been set.

### GetDELTATOKENSJSON

`func (o *EntraIDConnectionAttributes) GetDELTATOKENSJSON() string`

GetDELTATOKENSJSON returns the DELTATOKENSJSON field if non-nil, zero value otherwise.

### GetDELTATOKENSJSONOk

`func (o *EntraIDConnectionAttributes) GetDELTATOKENSJSONOk() (*string, bool)`

GetDELTATOKENSJSONOk returns a tuple with the DELTATOKENSJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDELTATOKENSJSON

`func (o *EntraIDConnectionAttributes) SetDELTATOKENSJSON(v string)`

SetDELTATOKENSJSON sets DELTATOKENSJSON field to given value.

### HasDELTATOKENSJSON

`func (o *EntraIDConnectionAttributes) HasDELTATOKENSJSON() bool`

HasDELTATOKENSJSON returns a boolean if a field has been set.

### GetAZURE_MGMT_ACCESS_TOKEN

`func (o *EntraIDConnectionAttributes) GetAZURE_MGMT_ACCESS_TOKEN() string`

GetAZURE_MGMT_ACCESS_TOKEN returns the AZURE_MGMT_ACCESS_TOKEN field if non-nil, zero value otherwise.

### GetAZURE_MGMT_ACCESS_TOKENOk

`func (o *EntraIDConnectionAttributes) GetAZURE_MGMT_ACCESS_TOKENOk() (*string, bool)`

GetAZURE_MGMT_ACCESS_TOKENOk returns a tuple with the AZURE_MGMT_ACCESS_TOKEN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAZURE_MGMT_ACCESS_TOKEN

`func (o *EntraIDConnectionAttributes) SetAZURE_MGMT_ACCESS_TOKEN(v string)`

SetAZURE_MGMT_ACCESS_TOKEN sets AZURE_MGMT_ACCESS_TOKEN field to given value.

### HasAZURE_MGMT_ACCESS_TOKEN

`func (o *EntraIDConnectionAttributes) HasAZURE_MGMT_ACCESS_TOKEN() bool`

HasAZURE_MGMT_ACCESS_TOKEN returns a boolean if a field has been set.

### GetCreateTeamJSON

`func (o *EntraIDConnectionAttributes) GetCreateTeamJSON() string`

GetCreateTeamJSON returns the CreateTeamJSON field if non-nil, zero value otherwise.

### GetCreateTeamJSONOk

`func (o *EntraIDConnectionAttributes) GetCreateTeamJSONOk() (*string, bool)`

GetCreateTeamJSONOk returns a tuple with the CreateTeamJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTeamJSON

`func (o *EntraIDConnectionAttributes) SetCreateTeamJSON(v string)`

SetCreateTeamJSON sets CreateTeamJSON field to given value.

### HasCreateTeamJSON

`func (o *EntraIDConnectionAttributes) HasCreateTeamJSON() bool`

HasCreateTeamJSON returns a boolean if a field has been set.

### GetConnectionTimeoutConfig

`func (o *EntraIDConnectionAttributes) GetConnectionTimeoutConfig() ConnectionTimeoutConfig`

GetConnectionTimeoutConfig returns the ConnectionTimeoutConfig field if non-nil, zero value otherwise.

### GetConnectionTimeoutConfigOk

`func (o *EntraIDConnectionAttributes) GetConnectionTimeoutConfigOk() (*ConnectionTimeoutConfig, bool)`

GetConnectionTimeoutConfigOk returns a tuple with the ConnectionTimeoutConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionTimeoutConfig

`func (o *EntraIDConnectionAttributes) SetConnectionTimeoutConfig(v ConnectionTimeoutConfig)`

SetConnectionTimeoutConfig sets ConnectionTimeoutConfig field to given value.

### HasConnectionTimeoutConfig

`func (o *EntraIDConnectionAttributes) HasConnectionTimeoutConfig() bool`

HasConnectionTimeoutConfig returns a boolean if a field has been set.

### GetENHANCEDDIRECTORYROLES

`func (o *EntraIDConnectionAttributes) GetENHANCEDDIRECTORYROLES() string`

GetENHANCEDDIRECTORYROLES returns the ENHANCEDDIRECTORYROLES field if non-nil, zero value otherwise.

### GetENHANCEDDIRECTORYROLESOk

`func (o *EntraIDConnectionAttributes) GetENHANCEDDIRECTORYROLESOk() (*string, bool)`

GetENHANCEDDIRECTORYROLESOk returns a tuple with the ENHANCEDDIRECTORYROLES field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetENHANCEDDIRECTORYROLES

`func (o *EntraIDConnectionAttributes) SetENHANCEDDIRECTORYROLES(v string)`

SetENHANCEDDIRECTORYROLES sets ENHANCEDDIRECTORYROLES field to given value.

### HasENHANCEDDIRECTORYROLES

`func (o *EntraIDConnectionAttributes) HasENHANCEDDIRECTORYROLES() bool`

HasENHANCEDDIRECTORYROLES returns a boolean if a field has been set.

### GetSTATUS_THRESHOLD_CONFIG

`func (o *EntraIDConnectionAttributes) GetSTATUS_THRESHOLD_CONFIG() string`

GetSTATUS_THRESHOLD_CONFIG returns the STATUS_THRESHOLD_CONFIG field if non-nil, zero value otherwise.

### GetSTATUS_THRESHOLD_CONFIGOk

`func (o *EntraIDConnectionAttributes) GetSTATUS_THRESHOLD_CONFIGOk() (*string, bool)`

GetSTATUS_THRESHOLD_CONFIGOk returns a tuple with the STATUS_THRESHOLD_CONFIG field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSTATUS_THRESHOLD_CONFIG

`func (o *EntraIDConnectionAttributes) SetSTATUS_THRESHOLD_CONFIG(v string)`

SetSTATUS_THRESHOLD_CONFIG sets STATUS_THRESHOLD_CONFIG field to given value.

### HasSTATUS_THRESHOLD_CONFIG

`func (o *EntraIDConnectionAttributes) HasSTATUS_THRESHOLD_CONFIG() bool`

HasSTATUS_THRESHOLD_CONFIG returns a boolean if a field has been set.

### GetACCOUNT_IMPORT_FIELDS

`func (o *EntraIDConnectionAttributes) GetACCOUNT_IMPORT_FIELDS() string`

GetACCOUNT_IMPORT_FIELDS returns the ACCOUNT_IMPORT_FIELDS field if non-nil, zero value otherwise.

### GetACCOUNT_IMPORT_FIELDSOk

`func (o *EntraIDConnectionAttributes) GetACCOUNT_IMPORT_FIELDSOk() (*string, bool)`

GetACCOUNT_IMPORT_FIELDSOk returns a tuple with the ACCOUNT_IMPORT_FIELDS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetACCOUNT_IMPORT_FIELDS

`func (o *EntraIDConnectionAttributes) SetACCOUNT_IMPORT_FIELDS(v string)`

SetACCOUNT_IMPORT_FIELDS sets ACCOUNT_IMPORT_FIELDS field to given value.

### HasACCOUNT_IMPORT_FIELDS

`func (o *EntraIDConnectionAttributes) HasACCOUNT_IMPORT_FIELDS() bool`

HasACCOUNT_IMPORT_FIELDS returns a boolean if a field has been set.

### GetRemoveAccountJSON

`func (o *EntraIDConnectionAttributes) GetRemoveAccountJSON() string`

GetRemoveAccountJSON returns the RemoveAccountJSON field if non-nil, zero value otherwise.

### GetRemoveAccountJSONOk

`func (o *EntraIDConnectionAttributes) GetRemoveAccountJSONOk() (*string, bool)`

GetRemoveAccountJSONOk returns a tuple with the RemoveAccountJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoveAccountJSON

`func (o *EntraIDConnectionAttributes) SetRemoveAccountJSON(v string)`

SetRemoveAccountJSON sets RemoveAccountJSON field to given value.

### HasRemoveAccountJSON

`func (o *EntraIDConnectionAttributes) HasRemoveAccountJSON() bool`

HasRemoveAccountJSON returns a boolean if a field has been set.

### GetChangePassJSON

`func (o *EntraIDConnectionAttributes) GetChangePassJSON() string`

GetChangePassJSON returns the ChangePassJSON field if non-nil, zero value otherwise.

### GetChangePassJSONOk

`func (o *EntraIDConnectionAttributes) GetChangePassJSONOk() (*string, bool)`

GetChangePassJSONOk returns a tuple with the ChangePassJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangePassJSON

`func (o *EntraIDConnectionAttributes) SetChangePassJSON(v string)`

SetChangePassJSON sets ChangePassJSON field to given value.

### HasChangePassJSON

`func (o *EntraIDConnectionAttributes) HasChangePassJSON() bool`

HasChangePassJSON returns a boolean if a field has been set.

### GetCLIENT_SECRET

`func (o *EntraIDConnectionAttributes) GetCLIENT_SECRET() string`

GetCLIENT_SECRET returns the CLIENT_SECRET field if non-nil, zero value otherwise.

### GetCLIENT_SECRETOk

`func (o *EntraIDConnectionAttributes) GetCLIENT_SECRETOk() (*string, bool)`

GetCLIENT_SECRETOk returns a tuple with the CLIENT_SECRET field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCLIENT_SECRET

`func (o *EntraIDConnectionAttributes) SetCLIENT_SECRET(v string)`

SetCLIENT_SECRET sets CLIENT_SECRET field to given value.

### HasCLIENT_SECRET

`func (o *EntraIDConnectionAttributes) HasCLIENT_SECRET() bool`

HasCLIENT_SECRET returns a boolean if a field has been set.

### GetENTITLEMENT_FILTER_JSON

`func (o *EntraIDConnectionAttributes) GetENTITLEMENT_FILTER_JSON() string`

GetENTITLEMENT_FILTER_JSON returns the ENTITLEMENT_FILTER_JSON field if non-nil, zero value otherwise.

### GetENTITLEMENT_FILTER_JSONOk

`func (o *EntraIDConnectionAttributes) GetENTITLEMENT_FILTER_JSONOk() (*string, bool)`

GetENTITLEMENT_FILTER_JSONOk returns a tuple with the ENTITLEMENT_FILTER_JSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetENTITLEMENT_FILTER_JSON

`func (o *EntraIDConnectionAttributes) SetENTITLEMENT_FILTER_JSON(v string)`

SetENTITLEMENT_FILTER_JSON sets ENTITLEMENT_FILTER_JSON field to given value.

### HasENTITLEMENT_FILTER_JSON

`func (o *EntraIDConnectionAttributes) HasENTITLEMENT_FILTER_JSON() bool`

HasENTITLEMENT_FILTER_JSON returns a boolean if a field has been set.

### GetSERVICE_ACCOUNT_ATTRIBUTES

`func (o *EntraIDConnectionAttributes) GetSERVICE_ACCOUNT_ATTRIBUTES() string`

GetSERVICE_ACCOUNT_ATTRIBUTES returns the SERVICE_ACCOUNT_ATTRIBUTES field if non-nil, zero value otherwise.

### GetSERVICE_ACCOUNT_ATTRIBUTESOk

`func (o *EntraIDConnectionAttributes) GetSERVICE_ACCOUNT_ATTRIBUTESOk() (*string, bool)`

GetSERVICE_ACCOUNT_ATTRIBUTESOk returns a tuple with the SERVICE_ACCOUNT_ATTRIBUTES field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSERVICE_ACCOUNT_ATTRIBUTES

`func (o *EntraIDConnectionAttributes) SetSERVICE_ACCOUNT_ATTRIBUTES(v string)`

SetSERVICE_ACCOUNT_ATTRIBUTES sets SERVICE_ACCOUNT_ATTRIBUTES field to given value.

### HasSERVICE_ACCOUNT_ATTRIBUTES

`func (o *EntraIDConnectionAttributes) HasSERVICE_ACCOUNT_ATTRIBUTES() bool`

HasSERVICE_ACCOUNT_ATTRIBUTES returns a boolean if a field has been set.

### GetAddAccessToEntitlementJSON

`func (o *EntraIDConnectionAttributes) GetAddAccessToEntitlementJSON() string`

GetAddAccessToEntitlementJSON returns the AddAccessToEntitlementJSON field if non-nil, zero value otherwise.

### GetAddAccessToEntitlementJSONOk

`func (o *EntraIDConnectionAttributes) GetAddAccessToEntitlementJSONOk() (*string, bool)`

GetAddAccessToEntitlementJSONOk returns a tuple with the AddAccessToEntitlementJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddAccessToEntitlementJSON

`func (o *EntraIDConnectionAttributes) SetAddAccessToEntitlementJSON(v string)`

SetAddAccessToEntitlementJSON sets AddAccessToEntitlementJSON field to given value.

### HasAddAccessToEntitlementJSON

`func (o *EntraIDConnectionAttributes) HasAddAccessToEntitlementJSON() bool`

HasAddAccessToEntitlementJSON returns a boolean if a field has been set.

### GetAUTHENTICATION_ENDPOINT

`func (o *EntraIDConnectionAttributes) GetAUTHENTICATION_ENDPOINT() string`

GetAUTHENTICATION_ENDPOINT returns the AUTHENTICATION_ENDPOINT field if non-nil, zero value otherwise.

### GetAUTHENTICATION_ENDPOINTOk

`func (o *EntraIDConnectionAttributes) GetAUTHENTICATION_ENDPOINTOk() (*string, bool)`

GetAUTHENTICATION_ENDPOINTOk returns a tuple with the AUTHENTICATION_ENDPOINT field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAUTHENTICATION_ENDPOINT

`func (o *EntraIDConnectionAttributes) SetAUTHENTICATION_ENDPOINT(v string)`

SetAUTHENTICATION_ENDPOINT sets AUTHENTICATION_ENDPOINT field to given value.

### HasAUTHENTICATION_ENDPOINT

`func (o *EntraIDConnectionAttributes) HasAUTHENTICATION_ENDPOINT() bool`

HasAUTHENTICATION_ENDPOINT returns a boolean if a field has been set.

### GetCreateServicePrincipalJSON

`func (o *EntraIDConnectionAttributes) GetCreateServicePrincipalJSON() string`

GetCreateServicePrincipalJSON returns the CreateServicePrincipalJSON field if non-nil, zero value otherwise.

### GetCreateServicePrincipalJSONOk

`func (o *EntraIDConnectionAttributes) GetCreateServicePrincipalJSONOk() (*string, bool)`

GetCreateServicePrincipalJSONOk returns a tuple with the CreateServicePrincipalJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateServicePrincipalJSON

`func (o *EntraIDConnectionAttributes) SetCreateServicePrincipalJSON(v string)`

SetCreateServicePrincipalJSON sets CreateServicePrincipalJSON field to given value.

### HasCreateServicePrincipalJSON

`func (o *EntraIDConnectionAttributes) HasCreateServicePrincipalJSON() bool`

HasCreateServicePrincipalJSON returns a boolean if a field has been set.

### GetMODIFYUSERDATAJSON

`func (o *EntraIDConnectionAttributes) GetMODIFYUSERDATAJSON() string`

GetMODIFYUSERDATAJSON returns the MODIFYUSERDATAJSON field if non-nil, zero value otherwise.

### GetMODIFYUSERDATAJSONOk

`func (o *EntraIDConnectionAttributes) GetMODIFYUSERDATAJSONOk() (*string, bool)`

GetMODIFYUSERDATAJSONOk returns a tuple with the MODIFYUSERDATAJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMODIFYUSERDATAJSON

`func (o *EntraIDConnectionAttributes) SetMODIFYUSERDATAJSON(v string)`

SetMODIFYUSERDATAJSON sets MODIFYUSERDATAJSON field to given value.

### HasMODIFYUSERDATAJSON

`func (o *EntraIDConnectionAttributes) HasMODIFYUSERDATAJSON() bool`

HasMODIFYUSERDATAJSON returns a boolean if a field has been set.

### GetIsTimeoutConfigValidated

`func (o *EntraIDConnectionAttributes) GetIsTimeoutConfigValidated() bool`

GetIsTimeoutConfigValidated returns the IsTimeoutConfigValidated field if non-nil, zero value otherwise.

### GetIsTimeoutConfigValidatedOk

`func (o *EntraIDConnectionAttributes) GetIsTimeoutConfigValidatedOk() (*bool, bool)`

GetIsTimeoutConfigValidatedOk returns a tuple with the IsTimeoutConfigValidated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsTimeoutConfigValidated

`func (o *EntraIDConnectionAttributes) SetIsTimeoutConfigValidated(v bool)`

SetIsTimeoutConfigValidated sets IsTimeoutConfigValidated field to given value.

### HasIsTimeoutConfigValidated

`func (o *EntraIDConnectionAttributes) HasIsTimeoutConfigValidated() bool`

HasIsTimeoutConfigValidated returns a boolean if a field has been set.

### GetRemoveAccessJSON

`func (o *EntraIDConnectionAttributes) GetRemoveAccessJSON() string`

GetRemoveAccessJSON returns the RemoveAccessJSON field if non-nil, zero value otherwise.

### GetRemoveAccessJSONOk

`func (o *EntraIDConnectionAttributes) GetRemoveAccessJSONOk() (*string, bool)`

GetRemoveAccessJSONOk returns a tuple with the RemoveAccessJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoveAccessJSON

`func (o *EntraIDConnectionAttributes) SetRemoveAccessJSON(v string)`

SetRemoveAccessJSON sets RemoveAccessJSON field to given value.

### HasRemoveAccessJSON

`func (o *EntraIDConnectionAttributes) HasRemoveAccessJSON() bool`

HasRemoveAccessJSON returns a boolean if a field has been set.

### GetCREATEUSERS

`func (o *EntraIDConnectionAttributes) GetCREATEUSERS() string`

GetCREATEUSERS returns the CREATEUSERS field if non-nil, zero value otherwise.

### GetCREATEUSERSOk

`func (o *EntraIDConnectionAttributes) GetCREATEUSERSOk() (*string, bool)`

GetCREATEUSERSOk returns a tuple with the CREATEUSERS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCREATEUSERS

`func (o *EntraIDConnectionAttributes) SetCREATEUSERS(v string)`

SetCREATEUSERS sets CREATEUSERS field to given value.

### HasCREATEUSERS

`func (o *EntraIDConnectionAttributes) HasCREATEUSERS() bool`

HasCREATEUSERS returns a boolean if a field has been set.

### GetRemoveAccessFromEntitlementJSON

`func (o *EntraIDConnectionAttributes) GetRemoveAccessFromEntitlementJSON() string`

GetRemoveAccessFromEntitlementJSON returns the RemoveAccessFromEntitlementJSON field if non-nil, zero value otherwise.

### GetRemoveAccessFromEntitlementJSONOk

`func (o *EntraIDConnectionAttributes) GetRemoveAccessFromEntitlementJSONOk() (*string, bool)`

GetRemoveAccessFromEntitlementJSONOk returns a tuple with the RemoveAccessFromEntitlementJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoveAccessFromEntitlementJSON

`func (o *EntraIDConnectionAttributes) SetRemoveAccessFromEntitlementJSON(v string)`

SetRemoveAccessFromEntitlementJSON sets RemoveAccessFromEntitlementJSON field to given value.

### HasRemoveAccessFromEntitlementJSON

`func (o *EntraIDConnectionAttributes) HasRemoveAccessFromEntitlementJSON() bool`

HasRemoveAccessFromEntitlementJSON returns a boolean if a field has been set.

### GetDisableAccountJSON

`func (o *EntraIDConnectionAttributes) GetDisableAccountJSON() string`

GetDisableAccountJSON returns the DisableAccountJSON field if non-nil, zero value otherwise.

### GetDisableAccountJSONOk

`func (o *EntraIDConnectionAttributes) GetDisableAccountJSONOk() (*string, bool)`

GetDisableAccountJSONOk returns a tuple with the DisableAccountJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableAccountJSON

`func (o *EntraIDConnectionAttributes) SetDisableAccountJSON(v string)`

SetDisableAccountJSON sets DisableAccountJSON field to given value.

### HasDisableAccountJSON

`func (o *EntraIDConnectionAttributes) HasDisableAccountJSON() bool`

HasDisableAccountJSON returns a boolean if a field has been set.

### GetCREATE_NEW_ENDPOINTS

`func (o *EntraIDConnectionAttributes) GetCREATE_NEW_ENDPOINTS() string`

GetCREATE_NEW_ENDPOINTS returns the CREATE_NEW_ENDPOINTS field if non-nil, zero value otherwise.

### GetCREATE_NEW_ENDPOINTSOk

`func (o *EntraIDConnectionAttributes) GetCREATE_NEW_ENDPOINTSOk() (*string, bool)`

GetCREATE_NEW_ENDPOINTSOk returns a tuple with the CREATE_NEW_ENDPOINTS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCREATE_NEW_ENDPOINTS

`func (o *EntraIDConnectionAttributes) SetCREATE_NEW_ENDPOINTS(v string)`

SetCREATE_NEW_ENDPOINTS sets CREATE_NEW_ENDPOINTS field to given value.

### HasCREATE_NEW_ENDPOINTS

`func (o *EntraIDConnectionAttributes) HasCREATE_NEW_ENDPOINTS() bool`

HasCREATE_NEW_ENDPOINTS returns a boolean if a field has been set.

### GetMANAGED_ACCOUNT_TYPE

`func (o *EntraIDConnectionAttributes) GetMANAGED_ACCOUNT_TYPE() string`

GetMANAGED_ACCOUNT_TYPE returns the MANAGED_ACCOUNT_TYPE field if non-nil, zero value otherwise.

### GetMANAGED_ACCOUNT_TYPEOk

`func (o *EntraIDConnectionAttributes) GetMANAGED_ACCOUNT_TYPEOk() (*string, bool)`

GetMANAGED_ACCOUNT_TYPEOk returns a tuple with the MANAGED_ACCOUNT_TYPE field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMANAGED_ACCOUNT_TYPE

`func (o *EntraIDConnectionAttributes) SetMANAGED_ACCOUNT_TYPE(v string)`

SetMANAGED_ACCOUNT_TYPE sets MANAGED_ACCOUNT_TYPE field to given value.

### HasMANAGED_ACCOUNT_TYPE

`func (o *EntraIDConnectionAttributes) HasMANAGED_ACCOUNT_TYPE() bool`

HasMANAGED_ACCOUNT_TYPE returns a boolean if a field has been set.

### GetACCOUNT_ATTRIBUTES

`func (o *EntraIDConnectionAttributes) GetACCOUNT_ATTRIBUTES() string`

GetACCOUNT_ATTRIBUTES returns the ACCOUNT_ATTRIBUTES field if non-nil, zero value otherwise.

### GetACCOUNT_ATTRIBUTESOk

`func (o *EntraIDConnectionAttributes) GetACCOUNT_ATTRIBUTESOk() (*string, bool)`

GetACCOUNT_ATTRIBUTESOk returns a tuple with the ACCOUNT_ATTRIBUTES field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetACCOUNT_ATTRIBUTES

`func (o *EntraIDConnectionAttributes) SetACCOUNT_ATTRIBUTES(v string)`

SetACCOUNT_ATTRIBUTES sets ACCOUNT_ATTRIBUTES field to given value.

### HasACCOUNT_ATTRIBUTES

`func (o *EntraIDConnectionAttributes) HasACCOUNT_ATTRIBUTES() bool`

HasACCOUNT_ATTRIBUTES returns a boolean if a field has been set.

### GetAAD_TENANT_ID

`func (o *EntraIDConnectionAttributes) GetAAD_TENANT_ID() string`

GetAAD_TENANT_ID returns the AAD_TENANT_ID field if non-nil, zero value otherwise.

### GetAAD_TENANT_IDOk

`func (o *EntraIDConnectionAttributes) GetAAD_TENANT_IDOk() (*string, bool)`

GetAAD_TENANT_IDOk returns a tuple with the AAD_TENANT_ID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAAD_TENANT_ID

`func (o *EntraIDConnectionAttributes) SetAAD_TENANT_ID(v string)`

SetAAD_TENANT_ID sets AAD_TENANT_ID field to given value.

### HasAAD_TENANT_ID

`func (o *EntraIDConnectionAttributes) HasAAD_TENANT_ID() bool`

HasAAD_TENANT_ID returns a boolean if a field has been set.

### GetUpdateGroupJSON

`func (o *EntraIDConnectionAttributes) GetUpdateGroupJSON() string`

GetUpdateGroupJSON returns the UpdateGroupJSON field if non-nil, zero value otherwise.

### GetUpdateGroupJSONOk

`func (o *EntraIDConnectionAttributes) GetUpdateGroupJSONOk() (*string, bool)`

GetUpdateGroupJSONOk returns a tuple with the UpdateGroupJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateGroupJSON

`func (o *EntraIDConnectionAttributes) SetUpdateGroupJSON(v string)`

SetUpdateGroupJSON sets UpdateGroupJSON field to given value.

### HasUpdateGroupJSON

`func (o *EntraIDConnectionAttributes) HasUpdateGroupJSON() bool`

HasUpdateGroupJSON returns a boolean if a field has been set.

### GetCreateGroupJSON

`func (o *EntraIDConnectionAttributes) GetCreateGroupJSON() string`

GetCreateGroupJSON returns the CreateGroupJSON field if non-nil, zero value otherwise.

### GetCreateGroupJSONOk

`func (o *EntraIDConnectionAttributes) GetCreateGroupJSONOk() (*string, bool)`

GetCreateGroupJSONOk returns a tuple with the CreateGroupJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateGroupJSON

`func (o *EntraIDConnectionAttributes) SetCreateGroupJSON(v string)`

SetCreateGroupJSON sets CreateGroupJSON field to given value.

### HasCreateGroupJSON

`func (o *EntraIDConnectionAttributes) HasCreateGroupJSON() bool`

HasCreateGroupJSON returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


