# EntraIDConnector

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CLIENT_ID** | **string** |  | 
**CLIENT_SECRET** | **string** | Property for CLIENT_SECRET | 
**ACCESS_TOKEN** | Pointer to **string** | Property for ACCESS_TOKEN | [optional] 
**AAD_TENANT_ID** | **string** | Property for AAD_TENANT_ID | 
**AZURE_MGMT_ACCESS_TOKEN** | Pointer to **string** | Property for AZURE_MGMT_ACCESS_TOKEN | [optional] 
**AUTHENTICATION_ENDPOINT** | Pointer to **string** | Property for AUTHENTICATION_ENDPOINT | [optional] 
**MICROSOFT_GRAPH_ENDPOINT** | Pointer to **string** | Property for MICROSOFT_GRAPH_ENDPOINT | [optional] 
**AZURE_MANAGEMENT_ENDPOINT** | Pointer to **string** | Property for AZURE_MANAGEMENT_ENDPOINT | [optional] 
**ImportUserJSON** | Pointer to **string** | Property for ImportUserJSON | [optional] 
**CREATEUSERS** | Pointer to **string** | Property for CREATEUSERS | [optional] 
**WINDOWS_CONNECTOR_JSON** | Pointer to **string** | Property for WINDOWS_CONNECTOR_JSON | [optional] 
**CREATE_NEW_ENDPOINTS** | Pointer to **string** | Property for CREATE_NEW_ENDPOINTS | [optional] 
**MANAGED_ACCOUNT_TYPE** | Pointer to **string** | Property for MANAGED_ACCOUNT_TYPE | [optional] 
**ACCOUNT_ATTRIBUTES** | Pointer to **string** | Property for ACCOUNT_ATTRIBUTES | [optional] 
**SERVICE_ACCOUNT_ATTRIBUTES** | Pointer to **string** | Property for SERVICE_ACCOUNT_ATTRIBUTES | [optional] 
**DELTATOKENSJSON** | Pointer to **string** | Property for DELTATOKENSJSON | [optional] 
**ACCOUNT_IMPORT_FIELDS** | Pointer to **string** | Property for ACCOUNT_IMPORT_FIELDS | [optional] 
**IMPORT_DEPTH** | Pointer to **string** | Property for IMPORT_DEPTH | [optional] 
**ENTITLEMENT_ATTRIBUTE** | Pointer to **string** | Property for ENTITLEMENT_ATTRIBUTE | [optional] 
**CreateAccountJSON** | Pointer to **string** | Property for CreateAccountJSON | [optional] 
**UpdateAccountJSON** | Pointer to **string** | Property for UpdateAccountJSON | [optional] 
**EnableAccountJSON** | Pointer to **string** | Property for EnableAccountJSON | [optional] 
**DisableAccountJSON** | Pointer to **string** | Property for DisableAccountJSON | [optional] 
**AddAccessJSON** | Pointer to **string** | Property for AddAccessJSON | [optional] 
**RemoveAccessJSON** | Pointer to **string** | Property for RemoveAccessJSON | [optional] 
**UpdateUserJSON** | Pointer to **string** | Property for UpdateUserJSON | [optional] 
**ChangePassJSON** | Pointer to **string** | Property for ChangePassJSON | [optional] 
**RemoveAccountJSON** | Pointer to **string** | Property for RemoveAccountJSON | [optional] 
**ConnectionJSON** | Pointer to **string** | Property for ConnectionJSON | [optional] 
**CreateGroupJSON** | Pointer to **string** | Property for CreateGroupJSON | [optional] 
**UpdateGroupJSON** | Pointer to **string** | Property for UpdateGroupJSON | [optional] 
**AddAccessToEntitlementJSON** | Pointer to **string** | Property for AddAccessToEntitlementJSON | [optional] 
**RemoveAccessFromEntitlementJSON** | Pointer to **string** | Property for RemoveAccessFromEntitlementJSON | [optional] 
**DeleteGroupJSON** | Pointer to **string** | Property for DeleteGroupJSON | [optional] 
**CreateServicePrincipalJSON** | Pointer to **string** | Property for CreateServicePrincipalJSON | [optional] 
**UpdateServicePrincipalJSON** | Pointer to **string** | Property for UpdateServicePrincipalJSON | [optional] 
**RemoveServicePrincipalJSON** | Pointer to **string** | Property for RemoveServicePrincipalJSON | [optional] 
**ENTITLEMENT_FILTER_JSON** | Pointer to **string** | Property for ENTITLEMENT_FILTER_JSON | [optional] 
**CreateTeamJSON** | Pointer to **string** | Property for CreateTeamJSON | [optional] 
**CreateChannelJSON** | Pointer to **string** | Property for CreateChannelJSON | [optional] 
**STATUS_THRESHOLD_CONFIG** | Pointer to **string** | Property for STATUS_THRESHOLD_CONFIG | [optional] 
**ACCOUNTS_FILTER** | Pointer to **string** | Property for ACCOUNTS_FILTER | [optional] 
**PAM_CONFIG** | Pointer to **string** | Property for PAM_CONFIG | [optional] 
**ENDPOINTS_FILTER** | Pointer to **string** | Property for ENDPOINTS_FILTER | [optional] 
**ConfigJSON** | Pointer to **string** | Property for ConfigJSON | [optional] 
**MODIFYUSERDATAJSON** | Pointer to **string** | Property for MODIFYUSERDATAJSON | [optional] 
**ENHANCEDDIRECTORYROLES** | Pointer to **string** | Property for ENHANCEDDIRECTORYROLES | [optional] 

## Methods

### NewEntraIDConnector

`func NewEntraIDConnector(cLIENTID string, cLIENTSECRET string, aADTENANTID string, ) *EntraIDConnector`

NewEntraIDConnector instantiates a new EntraIDConnector object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEntraIDConnectorWithDefaults

`func NewEntraIDConnectorWithDefaults() *EntraIDConnector`

NewEntraIDConnectorWithDefaults instantiates a new EntraIDConnector object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCLIENT_ID

`func (o *EntraIDConnector) GetCLIENT_ID() string`

GetCLIENT_ID returns the CLIENT_ID field if non-nil, zero value otherwise.

### GetCLIENT_IDOk

`func (o *EntraIDConnector) GetCLIENT_IDOk() (*string, bool)`

GetCLIENT_IDOk returns a tuple with the CLIENT_ID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCLIENT_ID

`func (o *EntraIDConnector) SetCLIENT_ID(v string)`

SetCLIENT_ID sets CLIENT_ID field to given value.


### GetCLIENT_SECRET

`func (o *EntraIDConnector) GetCLIENT_SECRET() string`

GetCLIENT_SECRET returns the CLIENT_SECRET field if non-nil, zero value otherwise.

### GetCLIENT_SECRETOk

`func (o *EntraIDConnector) GetCLIENT_SECRETOk() (*string, bool)`

GetCLIENT_SECRETOk returns a tuple with the CLIENT_SECRET field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCLIENT_SECRET

`func (o *EntraIDConnector) SetCLIENT_SECRET(v string)`

SetCLIENT_SECRET sets CLIENT_SECRET field to given value.


### GetACCESS_TOKEN

`func (o *EntraIDConnector) GetACCESS_TOKEN() string`

GetACCESS_TOKEN returns the ACCESS_TOKEN field if non-nil, zero value otherwise.

### GetACCESS_TOKENOk

`func (o *EntraIDConnector) GetACCESS_TOKENOk() (*string, bool)`

GetACCESS_TOKENOk returns a tuple with the ACCESS_TOKEN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetACCESS_TOKEN

`func (o *EntraIDConnector) SetACCESS_TOKEN(v string)`

SetACCESS_TOKEN sets ACCESS_TOKEN field to given value.

### HasACCESS_TOKEN

`func (o *EntraIDConnector) HasACCESS_TOKEN() bool`

HasACCESS_TOKEN returns a boolean if a field has been set.

### GetAAD_TENANT_ID

`func (o *EntraIDConnector) GetAAD_TENANT_ID() string`

GetAAD_TENANT_ID returns the AAD_TENANT_ID field if non-nil, zero value otherwise.

### GetAAD_TENANT_IDOk

`func (o *EntraIDConnector) GetAAD_TENANT_IDOk() (*string, bool)`

GetAAD_TENANT_IDOk returns a tuple with the AAD_TENANT_ID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAAD_TENANT_ID

`func (o *EntraIDConnector) SetAAD_TENANT_ID(v string)`

SetAAD_TENANT_ID sets AAD_TENANT_ID field to given value.


### GetAZURE_MGMT_ACCESS_TOKEN

`func (o *EntraIDConnector) GetAZURE_MGMT_ACCESS_TOKEN() string`

GetAZURE_MGMT_ACCESS_TOKEN returns the AZURE_MGMT_ACCESS_TOKEN field if non-nil, zero value otherwise.

### GetAZURE_MGMT_ACCESS_TOKENOk

`func (o *EntraIDConnector) GetAZURE_MGMT_ACCESS_TOKENOk() (*string, bool)`

GetAZURE_MGMT_ACCESS_TOKENOk returns a tuple with the AZURE_MGMT_ACCESS_TOKEN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAZURE_MGMT_ACCESS_TOKEN

`func (o *EntraIDConnector) SetAZURE_MGMT_ACCESS_TOKEN(v string)`

SetAZURE_MGMT_ACCESS_TOKEN sets AZURE_MGMT_ACCESS_TOKEN field to given value.

### HasAZURE_MGMT_ACCESS_TOKEN

`func (o *EntraIDConnector) HasAZURE_MGMT_ACCESS_TOKEN() bool`

HasAZURE_MGMT_ACCESS_TOKEN returns a boolean if a field has been set.

### GetAUTHENTICATION_ENDPOINT

`func (o *EntraIDConnector) GetAUTHENTICATION_ENDPOINT() string`

GetAUTHENTICATION_ENDPOINT returns the AUTHENTICATION_ENDPOINT field if non-nil, zero value otherwise.

### GetAUTHENTICATION_ENDPOINTOk

`func (o *EntraIDConnector) GetAUTHENTICATION_ENDPOINTOk() (*string, bool)`

GetAUTHENTICATION_ENDPOINTOk returns a tuple with the AUTHENTICATION_ENDPOINT field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAUTHENTICATION_ENDPOINT

`func (o *EntraIDConnector) SetAUTHENTICATION_ENDPOINT(v string)`

SetAUTHENTICATION_ENDPOINT sets AUTHENTICATION_ENDPOINT field to given value.

### HasAUTHENTICATION_ENDPOINT

`func (o *EntraIDConnector) HasAUTHENTICATION_ENDPOINT() bool`

HasAUTHENTICATION_ENDPOINT returns a boolean if a field has been set.

### GetMICROSOFT_GRAPH_ENDPOINT

`func (o *EntraIDConnector) GetMICROSOFT_GRAPH_ENDPOINT() string`

GetMICROSOFT_GRAPH_ENDPOINT returns the MICROSOFT_GRAPH_ENDPOINT field if non-nil, zero value otherwise.

### GetMICROSOFT_GRAPH_ENDPOINTOk

`func (o *EntraIDConnector) GetMICROSOFT_GRAPH_ENDPOINTOk() (*string, bool)`

GetMICROSOFT_GRAPH_ENDPOINTOk returns a tuple with the MICROSOFT_GRAPH_ENDPOINT field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMICROSOFT_GRAPH_ENDPOINT

`func (o *EntraIDConnector) SetMICROSOFT_GRAPH_ENDPOINT(v string)`

SetMICROSOFT_GRAPH_ENDPOINT sets MICROSOFT_GRAPH_ENDPOINT field to given value.

### HasMICROSOFT_GRAPH_ENDPOINT

`func (o *EntraIDConnector) HasMICROSOFT_GRAPH_ENDPOINT() bool`

HasMICROSOFT_GRAPH_ENDPOINT returns a boolean if a field has been set.

### GetAZURE_MANAGEMENT_ENDPOINT

`func (o *EntraIDConnector) GetAZURE_MANAGEMENT_ENDPOINT() string`

GetAZURE_MANAGEMENT_ENDPOINT returns the AZURE_MANAGEMENT_ENDPOINT field if non-nil, zero value otherwise.

### GetAZURE_MANAGEMENT_ENDPOINTOk

`func (o *EntraIDConnector) GetAZURE_MANAGEMENT_ENDPOINTOk() (*string, bool)`

GetAZURE_MANAGEMENT_ENDPOINTOk returns a tuple with the AZURE_MANAGEMENT_ENDPOINT field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAZURE_MANAGEMENT_ENDPOINT

`func (o *EntraIDConnector) SetAZURE_MANAGEMENT_ENDPOINT(v string)`

SetAZURE_MANAGEMENT_ENDPOINT sets AZURE_MANAGEMENT_ENDPOINT field to given value.

### HasAZURE_MANAGEMENT_ENDPOINT

`func (o *EntraIDConnector) HasAZURE_MANAGEMENT_ENDPOINT() bool`

HasAZURE_MANAGEMENT_ENDPOINT returns a boolean if a field has been set.

### GetImportUserJSON

`func (o *EntraIDConnector) GetImportUserJSON() string`

GetImportUserJSON returns the ImportUserJSON field if non-nil, zero value otherwise.

### GetImportUserJSONOk

`func (o *EntraIDConnector) GetImportUserJSONOk() (*string, bool)`

GetImportUserJSONOk returns a tuple with the ImportUserJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportUserJSON

`func (o *EntraIDConnector) SetImportUserJSON(v string)`

SetImportUserJSON sets ImportUserJSON field to given value.

### HasImportUserJSON

`func (o *EntraIDConnector) HasImportUserJSON() bool`

HasImportUserJSON returns a boolean if a field has been set.

### GetCREATEUSERS

`func (o *EntraIDConnector) GetCREATEUSERS() string`

GetCREATEUSERS returns the CREATEUSERS field if non-nil, zero value otherwise.

### GetCREATEUSERSOk

`func (o *EntraIDConnector) GetCREATEUSERSOk() (*string, bool)`

GetCREATEUSERSOk returns a tuple with the CREATEUSERS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCREATEUSERS

`func (o *EntraIDConnector) SetCREATEUSERS(v string)`

SetCREATEUSERS sets CREATEUSERS field to given value.

### HasCREATEUSERS

`func (o *EntraIDConnector) HasCREATEUSERS() bool`

HasCREATEUSERS returns a boolean if a field has been set.

### GetWINDOWS_CONNECTOR_JSON

`func (o *EntraIDConnector) GetWINDOWS_CONNECTOR_JSON() string`

GetWINDOWS_CONNECTOR_JSON returns the WINDOWS_CONNECTOR_JSON field if non-nil, zero value otherwise.

### GetWINDOWS_CONNECTOR_JSONOk

`func (o *EntraIDConnector) GetWINDOWS_CONNECTOR_JSONOk() (*string, bool)`

GetWINDOWS_CONNECTOR_JSONOk returns a tuple with the WINDOWS_CONNECTOR_JSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWINDOWS_CONNECTOR_JSON

`func (o *EntraIDConnector) SetWINDOWS_CONNECTOR_JSON(v string)`

SetWINDOWS_CONNECTOR_JSON sets WINDOWS_CONNECTOR_JSON field to given value.

### HasWINDOWS_CONNECTOR_JSON

`func (o *EntraIDConnector) HasWINDOWS_CONNECTOR_JSON() bool`

HasWINDOWS_CONNECTOR_JSON returns a boolean if a field has been set.

### GetCREATE_NEW_ENDPOINTS

`func (o *EntraIDConnector) GetCREATE_NEW_ENDPOINTS() string`

GetCREATE_NEW_ENDPOINTS returns the CREATE_NEW_ENDPOINTS field if non-nil, zero value otherwise.

### GetCREATE_NEW_ENDPOINTSOk

`func (o *EntraIDConnector) GetCREATE_NEW_ENDPOINTSOk() (*string, bool)`

GetCREATE_NEW_ENDPOINTSOk returns a tuple with the CREATE_NEW_ENDPOINTS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCREATE_NEW_ENDPOINTS

`func (o *EntraIDConnector) SetCREATE_NEW_ENDPOINTS(v string)`

SetCREATE_NEW_ENDPOINTS sets CREATE_NEW_ENDPOINTS field to given value.

### HasCREATE_NEW_ENDPOINTS

`func (o *EntraIDConnector) HasCREATE_NEW_ENDPOINTS() bool`

HasCREATE_NEW_ENDPOINTS returns a boolean if a field has been set.

### GetMANAGED_ACCOUNT_TYPE

`func (o *EntraIDConnector) GetMANAGED_ACCOUNT_TYPE() string`

GetMANAGED_ACCOUNT_TYPE returns the MANAGED_ACCOUNT_TYPE field if non-nil, zero value otherwise.

### GetMANAGED_ACCOUNT_TYPEOk

`func (o *EntraIDConnector) GetMANAGED_ACCOUNT_TYPEOk() (*string, bool)`

GetMANAGED_ACCOUNT_TYPEOk returns a tuple with the MANAGED_ACCOUNT_TYPE field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMANAGED_ACCOUNT_TYPE

`func (o *EntraIDConnector) SetMANAGED_ACCOUNT_TYPE(v string)`

SetMANAGED_ACCOUNT_TYPE sets MANAGED_ACCOUNT_TYPE field to given value.

### HasMANAGED_ACCOUNT_TYPE

`func (o *EntraIDConnector) HasMANAGED_ACCOUNT_TYPE() bool`

HasMANAGED_ACCOUNT_TYPE returns a boolean if a field has been set.

### GetACCOUNT_ATTRIBUTES

`func (o *EntraIDConnector) GetACCOUNT_ATTRIBUTES() string`

GetACCOUNT_ATTRIBUTES returns the ACCOUNT_ATTRIBUTES field if non-nil, zero value otherwise.

### GetACCOUNT_ATTRIBUTESOk

`func (o *EntraIDConnector) GetACCOUNT_ATTRIBUTESOk() (*string, bool)`

GetACCOUNT_ATTRIBUTESOk returns a tuple with the ACCOUNT_ATTRIBUTES field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetACCOUNT_ATTRIBUTES

`func (o *EntraIDConnector) SetACCOUNT_ATTRIBUTES(v string)`

SetACCOUNT_ATTRIBUTES sets ACCOUNT_ATTRIBUTES field to given value.

### HasACCOUNT_ATTRIBUTES

`func (o *EntraIDConnector) HasACCOUNT_ATTRIBUTES() bool`

HasACCOUNT_ATTRIBUTES returns a boolean if a field has been set.

### GetSERVICE_ACCOUNT_ATTRIBUTES

`func (o *EntraIDConnector) GetSERVICE_ACCOUNT_ATTRIBUTES() string`

GetSERVICE_ACCOUNT_ATTRIBUTES returns the SERVICE_ACCOUNT_ATTRIBUTES field if non-nil, zero value otherwise.

### GetSERVICE_ACCOUNT_ATTRIBUTESOk

`func (o *EntraIDConnector) GetSERVICE_ACCOUNT_ATTRIBUTESOk() (*string, bool)`

GetSERVICE_ACCOUNT_ATTRIBUTESOk returns a tuple with the SERVICE_ACCOUNT_ATTRIBUTES field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSERVICE_ACCOUNT_ATTRIBUTES

`func (o *EntraIDConnector) SetSERVICE_ACCOUNT_ATTRIBUTES(v string)`

SetSERVICE_ACCOUNT_ATTRIBUTES sets SERVICE_ACCOUNT_ATTRIBUTES field to given value.

### HasSERVICE_ACCOUNT_ATTRIBUTES

`func (o *EntraIDConnector) HasSERVICE_ACCOUNT_ATTRIBUTES() bool`

HasSERVICE_ACCOUNT_ATTRIBUTES returns a boolean if a field has been set.

### GetDELTATOKENSJSON

`func (o *EntraIDConnector) GetDELTATOKENSJSON() string`

GetDELTATOKENSJSON returns the DELTATOKENSJSON field if non-nil, zero value otherwise.

### GetDELTATOKENSJSONOk

`func (o *EntraIDConnector) GetDELTATOKENSJSONOk() (*string, bool)`

GetDELTATOKENSJSONOk returns a tuple with the DELTATOKENSJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDELTATOKENSJSON

`func (o *EntraIDConnector) SetDELTATOKENSJSON(v string)`

SetDELTATOKENSJSON sets DELTATOKENSJSON field to given value.

### HasDELTATOKENSJSON

`func (o *EntraIDConnector) HasDELTATOKENSJSON() bool`

HasDELTATOKENSJSON returns a boolean if a field has been set.

### GetACCOUNT_IMPORT_FIELDS

`func (o *EntraIDConnector) GetACCOUNT_IMPORT_FIELDS() string`

GetACCOUNT_IMPORT_FIELDS returns the ACCOUNT_IMPORT_FIELDS field if non-nil, zero value otherwise.

### GetACCOUNT_IMPORT_FIELDSOk

`func (o *EntraIDConnector) GetACCOUNT_IMPORT_FIELDSOk() (*string, bool)`

GetACCOUNT_IMPORT_FIELDSOk returns a tuple with the ACCOUNT_IMPORT_FIELDS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetACCOUNT_IMPORT_FIELDS

`func (o *EntraIDConnector) SetACCOUNT_IMPORT_FIELDS(v string)`

SetACCOUNT_IMPORT_FIELDS sets ACCOUNT_IMPORT_FIELDS field to given value.

### HasACCOUNT_IMPORT_FIELDS

`func (o *EntraIDConnector) HasACCOUNT_IMPORT_FIELDS() bool`

HasACCOUNT_IMPORT_FIELDS returns a boolean if a field has been set.

### GetIMPORT_DEPTH

`func (o *EntraIDConnector) GetIMPORT_DEPTH() string`

GetIMPORT_DEPTH returns the IMPORT_DEPTH field if non-nil, zero value otherwise.

### GetIMPORT_DEPTHOk

`func (o *EntraIDConnector) GetIMPORT_DEPTHOk() (*string, bool)`

GetIMPORT_DEPTHOk returns a tuple with the IMPORT_DEPTH field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIMPORT_DEPTH

`func (o *EntraIDConnector) SetIMPORT_DEPTH(v string)`

SetIMPORT_DEPTH sets IMPORT_DEPTH field to given value.

### HasIMPORT_DEPTH

`func (o *EntraIDConnector) HasIMPORT_DEPTH() bool`

HasIMPORT_DEPTH returns a boolean if a field has been set.

### GetENTITLEMENT_ATTRIBUTE

`func (o *EntraIDConnector) GetENTITLEMENT_ATTRIBUTE() string`

GetENTITLEMENT_ATTRIBUTE returns the ENTITLEMENT_ATTRIBUTE field if non-nil, zero value otherwise.

### GetENTITLEMENT_ATTRIBUTEOk

`func (o *EntraIDConnector) GetENTITLEMENT_ATTRIBUTEOk() (*string, bool)`

GetENTITLEMENT_ATTRIBUTEOk returns a tuple with the ENTITLEMENT_ATTRIBUTE field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetENTITLEMENT_ATTRIBUTE

`func (o *EntraIDConnector) SetENTITLEMENT_ATTRIBUTE(v string)`

SetENTITLEMENT_ATTRIBUTE sets ENTITLEMENT_ATTRIBUTE field to given value.

### HasENTITLEMENT_ATTRIBUTE

`func (o *EntraIDConnector) HasENTITLEMENT_ATTRIBUTE() bool`

HasENTITLEMENT_ATTRIBUTE returns a boolean if a field has been set.

### GetCreateAccountJSON

`func (o *EntraIDConnector) GetCreateAccountJSON() string`

GetCreateAccountJSON returns the CreateAccountJSON field if non-nil, zero value otherwise.

### GetCreateAccountJSONOk

`func (o *EntraIDConnector) GetCreateAccountJSONOk() (*string, bool)`

GetCreateAccountJSONOk returns a tuple with the CreateAccountJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateAccountJSON

`func (o *EntraIDConnector) SetCreateAccountJSON(v string)`

SetCreateAccountJSON sets CreateAccountJSON field to given value.

### HasCreateAccountJSON

`func (o *EntraIDConnector) HasCreateAccountJSON() bool`

HasCreateAccountJSON returns a boolean if a field has been set.

### GetUpdateAccountJSON

`func (o *EntraIDConnector) GetUpdateAccountJSON() string`

GetUpdateAccountJSON returns the UpdateAccountJSON field if non-nil, zero value otherwise.

### GetUpdateAccountJSONOk

`func (o *EntraIDConnector) GetUpdateAccountJSONOk() (*string, bool)`

GetUpdateAccountJSONOk returns a tuple with the UpdateAccountJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateAccountJSON

`func (o *EntraIDConnector) SetUpdateAccountJSON(v string)`

SetUpdateAccountJSON sets UpdateAccountJSON field to given value.

### HasUpdateAccountJSON

`func (o *EntraIDConnector) HasUpdateAccountJSON() bool`

HasUpdateAccountJSON returns a boolean if a field has been set.

### GetEnableAccountJSON

`func (o *EntraIDConnector) GetEnableAccountJSON() string`

GetEnableAccountJSON returns the EnableAccountJSON field if non-nil, zero value otherwise.

### GetEnableAccountJSONOk

`func (o *EntraIDConnector) GetEnableAccountJSONOk() (*string, bool)`

GetEnableAccountJSONOk returns a tuple with the EnableAccountJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableAccountJSON

`func (o *EntraIDConnector) SetEnableAccountJSON(v string)`

SetEnableAccountJSON sets EnableAccountJSON field to given value.

### HasEnableAccountJSON

`func (o *EntraIDConnector) HasEnableAccountJSON() bool`

HasEnableAccountJSON returns a boolean if a field has been set.

### GetDisableAccountJSON

`func (o *EntraIDConnector) GetDisableAccountJSON() string`

GetDisableAccountJSON returns the DisableAccountJSON field if non-nil, zero value otherwise.

### GetDisableAccountJSONOk

`func (o *EntraIDConnector) GetDisableAccountJSONOk() (*string, bool)`

GetDisableAccountJSONOk returns a tuple with the DisableAccountJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableAccountJSON

`func (o *EntraIDConnector) SetDisableAccountJSON(v string)`

SetDisableAccountJSON sets DisableAccountJSON field to given value.

### HasDisableAccountJSON

`func (o *EntraIDConnector) HasDisableAccountJSON() bool`

HasDisableAccountJSON returns a boolean if a field has been set.

### GetAddAccessJSON

`func (o *EntraIDConnector) GetAddAccessJSON() string`

GetAddAccessJSON returns the AddAccessJSON field if non-nil, zero value otherwise.

### GetAddAccessJSONOk

`func (o *EntraIDConnector) GetAddAccessJSONOk() (*string, bool)`

GetAddAccessJSONOk returns a tuple with the AddAccessJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddAccessJSON

`func (o *EntraIDConnector) SetAddAccessJSON(v string)`

SetAddAccessJSON sets AddAccessJSON field to given value.

### HasAddAccessJSON

`func (o *EntraIDConnector) HasAddAccessJSON() bool`

HasAddAccessJSON returns a boolean if a field has been set.

### GetRemoveAccessJSON

`func (o *EntraIDConnector) GetRemoveAccessJSON() string`

GetRemoveAccessJSON returns the RemoveAccessJSON field if non-nil, zero value otherwise.

### GetRemoveAccessJSONOk

`func (o *EntraIDConnector) GetRemoveAccessJSONOk() (*string, bool)`

GetRemoveAccessJSONOk returns a tuple with the RemoveAccessJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoveAccessJSON

`func (o *EntraIDConnector) SetRemoveAccessJSON(v string)`

SetRemoveAccessJSON sets RemoveAccessJSON field to given value.

### HasRemoveAccessJSON

`func (o *EntraIDConnector) HasRemoveAccessJSON() bool`

HasRemoveAccessJSON returns a boolean if a field has been set.

### GetUpdateUserJSON

`func (o *EntraIDConnector) GetUpdateUserJSON() string`

GetUpdateUserJSON returns the UpdateUserJSON field if non-nil, zero value otherwise.

### GetUpdateUserJSONOk

`func (o *EntraIDConnector) GetUpdateUserJSONOk() (*string, bool)`

GetUpdateUserJSONOk returns a tuple with the UpdateUserJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateUserJSON

`func (o *EntraIDConnector) SetUpdateUserJSON(v string)`

SetUpdateUserJSON sets UpdateUserJSON field to given value.

### HasUpdateUserJSON

`func (o *EntraIDConnector) HasUpdateUserJSON() bool`

HasUpdateUserJSON returns a boolean if a field has been set.

### GetChangePassJSON

`func (o *EntraIDConnector) GetChangePassJSON() string`

GetChangePassJSON returns the ChangePassJSON field if non-nil, zero value otherwise.

### GetChangePassJSONOk

`func (o *EntraIDConnector) GetChangePassJSONOk() (*string, bool)`

GetChangePassJSONOk returns a tuple with the ChangePassJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangePassJSON

`func (o *EntraIDConnector) SetChangePassJSON(v string)`

SetChangePassJSON sets ChangePassJSON field to given value.

### HasChangePassJSON

`func (o *EntraIDConnector) HasChangePassJSON() bool`

HasChangePassJSON returns a boolean if a field has been set.

### GetRemoveAccountJSON

`func (o *EntraIDConnector) GetRemoveAccountJSON() string`

GetRemoveAccountJSON returns the RemoveAccountJSON field if non-nil, zero value otherwise.

### GetRemoveAccountJSONOk

`func (o *EntraIDConnector) GetRemoveAccountJSONOk() (*string, bool)`

GetRemoveAccountJSONOk returns a tuple with the RemoveAccountJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoveAccountJSON

`func (o *EntraIDConnector) SetRemoveAccountJSON(v string)`

SetRemoveAccountJSON sets RemoveAccountJSON field to given value.

### HasRemoveAccountJSON

`func (o *EntraIDConnector) HasRemoveAccountJSON() bool`

HasRemoveAccountJSON returns a boolean if a field has been set.

### GetConnectionJSON

`func (o *EntraIDConnector) GetConnectionJSON() string`

GetConnectionJSON returns the ConnectionJSON field if non-nil, zero value otherwise.

### GetConnectionJSONOk

`func (o *EntraIDConnector) GetConnectionJSONOk() (*string, bool)`

GetConnectionJSONOk returns a tuple with the ConnectionJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionJSON

`func (o *EntraIDConnector) SetConnectionJSON(v string)`

SetConnectionJSON sets ConnectionJSON field to given value.

### HasConnectionJSON

`func (o *EntraIDConnector) HasConnectionJSON() bool`

HasConnectionJSON returns a boolean if a field has been set.

### GetCreateGroupJSON

`func (o *EntraIDConnector) GetCreateGroupJSON() string`

GetCreateGroupJSON returns the CreateGroupJSON field if non-nil, zero value otherwise.

### GetCreateGroupJSONOk

`func (o *EntraIDConnector) GetCreateGroupJSONOk() (*string, bool)`

GetCreateGroupJSONOk returns a tuple with the CreateGroupJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateGroupJSON

`func (o *EntraIDConnector) SetCreateGroupJSON(v string)`

SetCreateGroupJSON sets CreateGroupJSON field to given value.

### HasCreateGroupJSON

`func (o *EntraIDConnector) HasCreateGroupJSON() bool`

HasCreateGroupJSON returns a boolean if a field has been set.

### GetUpdateGroupJSON

`func (o *EntraIDConnector) GetUpdateGroupJSON() string`

GetUpdateGroupJSON returns the UpdateGroupJSON field if non-nil, zero value otherwise.

### GetUpdateGroupJSONOk

`func (o *EntraIDConnector) GetUpdateGroupJSONOk() (*string, bool)`

GetUpdateGroupJSONOk returns a tuple with the UpdateGroupJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateGroupJSON

`func (o *EntraIDConnector) SetUpdateGroupJSON(v string)`

SetUpdateGroupJSON sets UpdateGroupJSON field to given value.

### HasUpdateGroupJSON

`func (o *EntraIDConnector) HasUpdateGroupJSON() bool`

HasUpdateGroupJSON returns a boolean if a field has been set.

### GetAddAccessToEntitlementJSON

`func (o *EntraIDConnector) GetAddAccessToEntitlementJSON() string`

GetAddAccessToEntitlementJSON returns the AddAccessToEntitlementJSON field if non-nil, zero value otherwise.

### GetAddAccessToEntitlementJSONOk

`func (o *EntraIDConnector) GetAddAccessToEntitlementJSONOk() (*string, bool)`

GetAddAccessToEntitlementJSONOk returns a tuple with the AddAccessToEntitlementJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddAccessToEntitlementJSON

`func (o *EntraIDConnector) SetAddAccessToEntitlementJSON(v string)`

SetAddAccessToEntitlementJSON sets AddAccessToEntitlementJSON field to given value.

### HasAddAccessToEntitlementJSON

`func (o *EntraIDConnector) HasAddAccessToEntitlementJSON() bool`

HasAddAccessToEntitlementJSON returns a boolean if a field has been set.

### GetRemoveAccessFromEntitlementJSON

`func (o *EntraIDConnector) GetRemoveAccessFromEntitlementJSON() string`

GetRemoveAccessFromEntitlementJSON returns the RemoveAccessFromEntitlementJSON field if non-nil, zero value otherwise.

### GetRemoveAccessFromEntitlementJSONOk

`func (o *EntraIDConnector) GetRemoveAccessFromEntitlementJSONOk() (*string, bool)`

GetRemoveAccessFromEntitlementJSONOk returns a tuple with the RemoveAccessFromEntitlementJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoveAccessFromEntitlementJSON

`func (o *EntraIDConnector) SetRemoveAccessFromEntitlementJSON(v string)`

SetRemoveAccessFromEntitlementJSON sets RemoveAccessFromEntitlementJSON field to given value.

### HasRemoveAccessFromEntitlementJSON

`func (o *EntraIDConnector) HasRemoveAccessFromEntitlementJSON() bool`

HasRemoveAccessFromEntitlementJSON returns a boolean if a field has been set.

### GetDeleteGroupJSON

`func (o *EntraIDConnector) GetDeleteGroupJSON() string`

GetDeleteGroupJSON returns the DeleteGroupJSON field if non-nil, zero value otherwise.

### GetDeleteGroupJSONOk

`func (o *EntraIDConnector) GetDeleteGroupJSONOk() (*string, bool)`

GetDeleteGroupJSONOk returns a tuple with the DeleteGroupJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteGroupJSON

`func (o *EntraIDConnector) SetDeleteGroupJSON(v string)`

SetDeleteGroupJSON sets DeleteGroupJSON field to given value.

### HasDeleteGroupJSON

`func (o *EntraIDConnector) HasDeleteGroupJSON() bool`

HasDeleteGroupJSON returns a boolean if a field has been set.

### GetCreateServicePrincipalJSON

`func (o *EntraIDConnector) GetCreateServicePrincipalJSON() string`

GetCreateServicePrincipalJSON returns the CreateServicePrincipalJSON field if non-nil, zero value otherwise.

### GetCreateServicePrincipalJSONOk

`func (o *EntraIDConnector) GetCreateServicePrincipalJSONOk() (*string, bool)`

GetCreateServicePrincipalJSONOk returns a tuple with the CreateServicePrincipalJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateServicePrincipalJSON

`func (o *EntraIDConnector) SetCreateServicePrincipalJSON(v string)`

SetCreateServicePrincipalJSON sets CreateServicePrincipalJSON field to given value.

### HasCreateServicePrincipalJSON

`func (o *EntraIDConnector) HasCreateServicePrincipalJSON() bool`

HasCreateServicePrincipalJSON returns a boolean if a field has been set.

### GetUpdateServicePrincipalJSON

`func (o *EntraIDConnector) GetUpdateServicePrincipalJSON() string`

GetUpdateServicePrincipalJSON returns the UpdateServicePrincipalJSON field if non-nil, zero value otherwise.

### GetUpdateServicePrincipalJSONOk

`func (o *EntraIDConnector) GetUpdateServicePrincipalJSONOk() (*string, bool)`

GetUpdateServicePrincipalJSONOk returns a tuple with the UpdateServicePrincipalJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateServicePrincipalJSON

`func (o *EntraIDConnector) SetUpdateServicePrincipalJSON(v string)`

SetUpdateServicePrincipalJSON sets UpdateServicePrincipalJSON field to given value.

### HasUpdateServicePrincipalJSON

`func (o *EntraIDConnector) HasUpdateServicePrincipalJSON() bool`

HasUpdateServicePrincipalJSON returns a boolean if a field has been set.

### GetRemoveServicePrincipalJSON

`func (o *EntraIDConnector) GetRemoveServicePrincipalJSON() string`

GetRemoveServicePrincipalJSON returns the RemoveServicePrincipalJSON field if non-nil, zero value otherwise.

### GetRemoveServicePrincipalJSONOk

`func (o *EntraIDConnector) GetRemoveServicePrincipalJSONOk() (*string, bool)`

GetRemoveServicePrincipalJSONOk returns a tuple with the RemoveServicePrincipalJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoveServicePrincipalJSON

`func (o *EntraIDConnector) SetRemoveServicePrincipalJSON(v string)`

SetRemoveServicePrincipalJSON sets RemoveServicePrincipalJSON field to given value.

### HasRemoveServicePrincipalJSON

`func (o *EntraIDConnector) HasRemoveServicePrincipalJSON() bool`

HasRemoveServicePrincipalJSON returns a boolean if a field has been set.

### GetENTITLEMENT_FILTER_JSON

`func (o *EntraIDConnector) GetENTITLEMENT_FILTER_JSON() string`

GetENTITLEMENT_FILTER_JSON returns the ENTITLEMENT_FILTER_JSON field if non-nil, zero value otherwise.

### GetENTITLEMENT_FILTER_JSONOk

`func (o *EntraIDConnector) GetENTITLEMENT_FILTER_JSONOk() (*string, bool)`

GetENTITLEMENT_FILTER_JSONOk returns a tuple with the ENTITLEMENT_FILTER_JSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetENTITLEMENT_FILTER_JSON

`func (o *EntraIDConnector) SetENTITLEMENT_FILTER_JSON(v string)`

SetENTITLEMENT_FILTER_JSON sets ENTITLEMENT_FILTER_JSON field to given value.

### HasENTITLEMENT_FILTER_JSON

`func (o *EntraIDConnector) HasENTITLEMENT_FILTER_JSON() bool`

HasENTITLEMENT_FILTER_JSON returns a boolean if a field has been set.

### GetCreateTeamJSON

`func (o *EntraIDConnector) GetCreateTeamJSON() string`

GetCreateTeamJSON returns the CreateTeamJSON field if non-nil, zero value otherwise.

### GetCreateTeamJSONOk

`func (o *EntraIDConnector) GetCreateTeamJSONOk() (*string, bool)`

GetCreateTeamJSONOk returns a tuple with the CreateTeamJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTeamJSON

`func (o *EntraIDConnector) SetCreateTeamJSON(v string)`

SetCreateTeamJSON sets CreateTeamJSON field to given value.

### HasCreateTeamJSON

`func (o *EntraIDConnector) HasCreateTeamJSON() bool`

HasCreateTeamJSON returns a boolean if a field has been set.

### GetCreateChannelJSON

`func (o *EntraIDConnector) GetCreateChannelJSON() string`

GetCreateChannelJSON returns the CreateChannelJSON field if non-nil, zero value otherwise.

### GetCreateChannelJSONOk

`func (o *EntraIDConnector) GetCreateChannelJSONOk() (*string, bool)`

GetCreateChannelJSONOk returns a tuple with the CreateChannelJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateChannelJSON

`func (o *EntraIDConnector) SetCreateChannelJSON(v string)`

SetCreateChannelJSON sets CreateChannelJSON field to given value.

### HasCreateChannelJSON

`func (o *EntraIDConnector) HasCreateChannelJSON() bool`

HasCreateChannelJSON returns a boolean if a field has been set.

### GetSTATUS_THRESHOLD_CONFIG

`func (o *EntraIDConnector) GetSTATUS_THRESHOLD_CONFIG() string`

GetSTATUS_THRESHOLD_CONFIG returns the STATUS_THRESHOLD_CONFIG field if non-nil, zero value otherwise.

### GetSTATUS_THRESHOLD_CONFIGOk

`func (o *EntraIDConnector) GetSTATUS_THRESHOLD_CONFIGOk() (*string, bool)`

GetSTATUS_THRESHOLD_CONFIGOk returns a tuple with the STATUS_THRESHOLD_CONFIG field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSTATUS_THRESHOLD_CONFIG

`func (o *EntraIDConnector) SetSTATUS_THRESHOLD_CONFIG(v string)`

SetSTATUS_THRESHOLD_CONFIG sets STATUS_THRESHOLD_CONFIG field to given value.

### HasSTATUS_THRESHOLD_CONFIG

`func (o *EntraIDConnector) HasSTATUS_THRESHOLD_CONFIG() bool`

HasSTATUS_THRESHOLD_CONFIG returns a boolean if a field has been set.

### GetACCOUNTS_FILTER

`func (o *EntraIDConnector) GetACCOUNTS_FILTER() string`

GetACCOUNTS_FILTER returns the ACCOUNTS_FILTER field if non-nil, zero value otherwise.

### GetACCOUNTS_FILTEROk

`func (o *EntraIDConnector) GetACCOUNTS_FILTEROk() (*string, bool)`

GetACCOUNTS_FILTEROk returns a tuple with the ACCOUNTS_FILTER field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetACCOUNTS_FILTER

`func (o *EntraIDConnector) SetACCOUNTS_FILTER(v string)`

SetACCOUNTS_FILTER sets ACCOUNTS_FILTER field to given value.

### HasACCOUNTS_FILTER

`func (o *EntraIDConnector) HasACCOUNTS_FILTER() bool`

HasACCOUNTS_FILTER returns a boolean if a field has been set.

### GetPAM_CONFIG

`func (o *EntraIDConnector) GetPAM_CONFIG() string`

GetPAM_CONFIG returns the PAM_CONFIG field if non-nil, zero value otherwise.

### GetPAM_CONFIGOk

`func (o *EntraIDConnector) GetPAM_CONFIGOk() (*string, bool)`

GetPAM_CONFIGOk returns a tuple with the PAM_CONFIG field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPAM_CONFIG

`func (o *EntraIDConnector) SetPAM_CONFIG(v string)`

SetPAM_CONFIG sets PAM_CONFIG field to given value.

### HasPAM_CONFIG

`func (o *EntraIDConnector) HasPAM_CONFIG() bool`

HasPAM_CONFIG returns a boolean if a field has been set.

### GetENDPOINTS_FILTER

`func (o *EntraIDConnector) GetENDPOINTS_FILTER() string`

GetENDPOINTS_FILTER returns the ENDPOINTS_FILTER field if non-nil, zero value otherwise.

### GetENDPOINTS_FILTEROk

`func (o *EntraIDConnector) GetENDPOINTS_FILTEROk() (*string, bool)`

GetENDPOINTS_FILTEROk returns a tuple with the ENDPOINTS_FILTER field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetENDPOINTS_FILTER

`func (o *EntraIDConnector) SetENDPOINTS_FILTER(v string)`

SetENDPOINTS_FILTER sets ENDPOINTS_FILTER field to given value.

### HasENDPOINTS_FILTER

`func (o *EntraIDConnector) HasENDPOINTS_FILTER() bool`

HasENDPOINTS_FILTER returns a boolean if a field has been set.

### GetConfigJSON

`func (o *EntraIDConnector) GetConfigJSON() string`

GetConfigJSON returns the ConfigJSON field if non-nil, zero value otherwise.

### GetConfigJSONOk

`func (o *EntraIDConnector) GetConfigJSONOk() (*string, bool)`

GetConfigJSONOk returns a tuple with the ConfigJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigJSON

`func (o *EntraIDConnector) SetConfigJSON(v string)`

SetConfigJSON sets ConfigJSON field to given value.

### HasConfigJSON

`func (o *EntraIDConnector) HasConfigJSON() bool`

HasConfigJSON returns a boolean if a field has been set.

### GetMODIFYUSERDATAJSON

`func (o *EntraIDConnector) GetMODIFYUSERDATAJSON() string`

GetMODIFYUSERDATAJSON returns the MODIFYUSERDATAJSON field if non-nil, zero value otherwise.

### GetMODIFYUSERDATAJSONOk

`func (o *EntraIDConnector) GetMODIFYUSERDATAJSONOk() (*string, bool)`

GetMODIFYUSERDATAJSONOk returns a tuple with the MODIFYUSERDATAJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMODIFYUSERDATAJSON

`func (o *EntraIDConnector) SetMODIFYUSERDATAJSON(v string)`

SetMODIFYUSERDATAJSON sets MODIFYUSERDATAJSON field to given value.

### HasMODIFYUSERDATAJSON

`func (o *EntraIDConnector) HasMODIFYUSERDATAJSON() bool`

HasMODIFYUSERDATAJSON returns a boolean if a field has been set.

### GetENHANCEDDIRECTORYROLES

`func (o *EntraIDConnector) GetENHANCEDDIRECTORYROLES() string`

GetENHANCEDDIRECTORYROLES returns the ENHANCEDDIRECTORYROLES field if non-nil, zero value otherwise.

### GetENHANCEDDIRECTORYROLESOk

`func (o *EntraIDConnector) GetENHANCEDDIRECTORYROLESOk() (*string, bool)`

GetENHANCEDDIRECTORYROLESOk returns a tuple with the ENHANCEDDIRECTORYROLES field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetENHANCEDDIRECTORYROLES

`func (o *EntraIDConnector) SetENHANCEDDIRECTORYROLES(v string)`

SetENHANCEDDIRECTORYROLES sets ENHANCEDDIRECTORYROLES field to given value.

### HasENHANCEDDIRECTORYROLES

`func (o *EntraIDConnector) HasENHANCEDDIRECTORYROLES() bool`

HasENHANCEDDIRECTORYROLES returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


