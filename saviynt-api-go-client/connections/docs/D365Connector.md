# D365Connector

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BASEURL** | **string** |  | 
**TENANT_ID** | **string** |  | 
**LOGIN_URL** | **string** |  | 
**CLIENT_ID** | **string** |  | 
**CLIENT_SECRET** | **string** |  | 
**USER_FILTER** | Pointer to **string** | Property for USER_FILTER | [optional] 
**USER_IMPORT_MAPPING** | Pointer to **string** | Property for USER_IMPORT_MAPPING | [optional] 
**ACCOUNT_IMPORT_MAPPING** | Pointer to **string** |  | [optional] 
**ORGANIZATION_FILTER** | Pointer to **string** |  | [optional] 
**STATUS_THRESHOLD_CONFIG** | Pointer to **string** |  | [optional] 
**ConfigJSON** | Pointer to **string** |  | [optional] 
**SCOPE** | Pointer to **string** | If present - 2.0 API will be used. Accepts space separated values. If blank 1.0 API is used | [optional] 
**CreateAccountJSON** | Pointer to **string** | JSON to specify the Field Value which will be used to Create the New Account | [optional] 
**UpdateAccountJSON** | Pointer to **string** | JSON to specify the Field Value which will be used to Update existing Account | [optional] 
**EnableAccountJSON** | Pointer to **string** | JSON to specify the different attributes to be checked and action to be performed for enabling a disabled account | [optional] 
**DisableAccountJSON** | Pointer to **string** | JSON to specify the different attributes to be checked and action to be performed for disabling a enabled account | [optional] 
**AddAccessJSON** | Pointer to **string** | JSON to ADD Access to an account | [optional] 
**RemoveAccessJSON** | Pointer to **string** | JSON to REMOVE Access from an account | [optional] 
**RemoveAccountJSON** | Pointer to **string** | JSON to specify the different attributes to be checked and action to be performed for deleting/suspending an account | [optional] 

## Methods

### NewD365Connector

`func NewD365Connector(bASEURL string, tENANTID string, lOGINURL string, cLIENTID string, cLIENTSECRET string, ) *D365Connector`

NewD365Connector instantiates a new D365Connector object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewD365ConnectorWithDefaults

`func NewD365ConnectorWithDefaults() *D365Connector`

NewD365ConnectorWithDefaults instantiates a new D365Connector object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBASEURL

`func (o *D365Connector) GetBASEURL() string`

GetBASEURL returns the BASEURL field if non-nil, zero value otherwise.

### GetBASEURLOk

`func (o *D365Connector) GetBASEURLOk() (*string, bool)`

GetBASEURLOk returns a tuple with the BASEURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBASEURL

`func (o *D365Connector) SetBASEURL(v string)`

SetBASEURL sets BASEURL field to given value.


### GetTENANT_ID

`func (o *D365Connector) GetTENANT_ID() string`

GetTENANT_ID returns the TENANT_ID field if non-nil, zero value otherwise.

### GetTENANT_IDOk

`func (o *D365Connector) GetTENANT_IDOk() (*string, bool)`

GetTENANT_IDOk returns a tuple with the TENANT_ID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTENANT_ID

`func (o *D365Connector) SetTENANT_ID(v string)`

SetTENANT_ID sets TENANT_ID field to given value.


### GetLOGIN_URL

`func (o *D365Connector) GetLOGIN_URL() string`

GetLOGIN_URL returns the LOGIN_URL field if non-nil, zero value otherwise.

### GetLOGIN_URLOk

`func (o *D365Connector) GetLOGIN_URLOk() (*string, bool)`

GetLOGIN_URLOk returns a tuple with the LOGIN_URL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLOGIN_URL

`func (o *D365Connector) SetLOGIN_URL(v string)`

SetLOGIN_URL sets LOGIN_URL field to given value.


### GetCLIENT_ID

`func (o *D365Connector) GetCLIENT_ID() string`

GetCLIENT_ID returns the CLIENT_ID field if non-nil, zero value otherwise.

### GetCLIENT_IDOk

`func (o *D365Connector) GetCLIENT_IDOk() (*string, bool)`

GetCLIENT_IDOk returns a tuple with the CLIENT_ID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCLIENT_ID

`func (o *D365Connector) SetCLIENT_ID(v string)`

SetCLIENT_ID sets CLIENT_ID field to given value.


### GetCLIENT_SECRET

`func (o *D365Connector) GetCLIENT_SECRET() string`

GetCLIENT_SECRET returns the CLIENT_SECRET field if non-nil, zero value otherwise.

### GetCLIENT_SECRETOk

`func (o *D365Connector) GetCLIENT_SECRETOk() (*string, bool)`

GetCLIENT_SECRETOk returns a tuple with the CLIENT_SECRET field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCLIENT_SECRET

`func (o *D365Connector) SetCLIENT_SECRET(v string)`

SetCLIENT_SECRET sets CLIENT_SECRET field to given value.


### GetUSER_FILTER

`func (o *D365Connector) GetUSER_FILTER() string`

GetUSER_FILTER returns the USER_FILTER field if non-nil, zero value otherwise.

### GetUSER_FILTEROk

`func (o *D365Connector) GetUSER_FILTEROk() (*string, bool)`

GetUSER_FILTEROk returns a tuple with the USER_FILTER field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUSER_FILTER

`func (o *D365Connector) SetUSER_FILTER(v string)`

SetUSER_FILTER sets USER_FILTER field to given value.

### HasUSER_FILTER

`func (o *D365Connector) HasUSER_FILTER() bool`

HasUSER_FILTER returns a boolean if a field has been set.

### GetUSER_IMPORT_MAPPING

`func (o *D365Connector) GetUSER_IMPORT_MAPPING() string`

GetUSER_IMPORT_MAPPING returns the USER_IMPORT_MAPPING field if non-nil, zero value otherwise.

### GetUSER_IMPORT_MAPPINGOk

`func (o *D365Connector) GetUSER_IMPORT_MAPPINGOk() (*string, bool)`

GetUSER_IMPORT_MAPPINGOk returns a tuple with the USER_IMPORT_MAPPING field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUSER_IMPORT_MAPPING

`func (o *D365Connector) SetUSER_IMPORT_MAPPING(v string)`

SetUSER_IMPORT_MAPPING sets USER_IMPORT_MAPPING field to given value.

### HasUSER_IMPORT_MAPPING

`func (o *D365Connector) HasUSER_IMPORT_MAPPING() bool`

HasUSER_IMPORT_MAPPING returns a boolean if a field has been set.

### GetACCOUNT_IMPORT_MAPPING

`func (o *D365Connector) GetACCOUNT_IMPORT_MAPPING() string`

GetACCOUNT_IMPORT_MAPPING returns the ACCOUNT_IMPORT_MAPPING field if non-nil, zero value otherwise.

### GetACCOUNT_IMPORT_MAPPINGOk

`func (o *D365Connector) GetACCOUNT_IMPORT_MAPPINGOk() (*string, bool)`

GetACCOUNT_IMPORT_MAPPINGOk returns a tuple with the ACCOUNT_IMPORT_MAPPING field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetACCOUNT_IMPORT_MAPPING

`func (o *D365Connector) SetACCOUNT_IMPORT_MAPPING(v string)`

SetACCOUNT_IMPORT_MAPPING sets ACCOUNT_IMPORT_MAPPING field to given value.

### HasACCOUNT_IMPORT_MAPPING

`func (o *D365Connector) HasACCOUNT_IMPORT_MAPPING() bool`

HasACCOUNT_IMPORT_MAPPING returns a boolean if a field has been set.

### GetORGANIZATION_FILTER

`func (o *D365Connector) GetORGANIZATION_FILTER() string`

GetORGANIZATION_FILTER returns the ORGANIZATION_FILTER field if non-nil, zero value otherwise.

### GetORGANIZATION_FILTEROk

`func (o *D365Connector) GetORGANIZATION_FILTEROk() (*string, bool)`

GetORGANIZATION_FILTEROk returns a tuple with the ORGANIZATION_FILTER field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetORGANIZATION_FILTER

`func (o *D365Connector) SetORGANIZATION_FILTER(v string)`

SetORGANIZATION_FILTER sets ORGANIZATION_FILTER field to given value.

### HasORGANIZATION_FILTER

`func (o *D365Connector) HasORGANIZATION_FILTER() bool`

HasORGANIZATION_FILTER returns a boolean if a field has been set.

### GetSTATUS_THRESHOLD_CONFIG

`func (o *D365Connector) GetSTATUS_THRESHOLD_CONFIG() string`

GetSTATUS_THRESHOLD_CONFIG returns the STATUS_THRESHOLD_CONFIG field if non-nil, zero value otherwise.

### GetSTATUS_THRESHOLD_CONFIGOk

`func (o *D365Connector) GetSTATUS_THRESHOLD_CONFIGOk() (*string, bool)`

GetSTATUS_THRESHOLD_CONFIGOk returns a tuple with the STATUS_THRESHOLD_CONFIG field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSTATUS_THRESHOLD_CONFIG

`func (o *D365Connector) SetSTATUS_THRESHOLD_CONFIG(v string)`

SetSTATUS_THRESHOLD_CONFIG sets STATUS_THRESHOLD_CONFIG field to given value.

### HasSTATUS_THRESHOLD_CONFIG

`func (o *D365Connector) HasSTATUS_THRESHOLD_CONFIG() bool`

HasSTATUS_THRESHOLD_CONFIG returns a boolean if a field has been set.

### GetConfigJSON

`func (o *D365Connector) GetConfigJSON() string`

GetConfigJSON returns the ConfigJSON field if non-nil, zero value otherwise.

### GetConfigJSONOk

`func (o *D365Connector) GetConfigJSONOk() (*string, bool)`

GetConfigJSONOk returns a tuple with the ConfigJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigJSON

`func (o *D365Connector) SetConfigJSON(v string)`

SetConfigJSON sets ConfigJSON field to given value.

### HasConfigJSON

`func (o *D365Connector) HasConfigJSON() bool`

HasConfigJSON returns a boolean if a field has been set.

### GetSCOPE

`func (o *D365Connector) GetSCOPE() string`

GetSCOPE returns the SCOPE field if non-nil, zero value otherwise.

### GetSCOPEOk

`func (o *D365Connector) GetSCOPEOk() (*string, bool)`

GetSCOPEOk returns a tuple with the SCOPE field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSCOPE

`func (o *D365Connector) SetSCOPE(v string)`

SetSCOPE sets SCOPE field to given value.

### HasSCOPE

`func (o *D365Connector) HasSCOPE() bool`

HasSCOPE returns a boolean if a field has been set.

### GetCreateAccountJSON

`func (o *D365Connector) GetCreateAccountJSON() string`

GetCreateAccountJSON returns the CreateAccountJSON field if non-nil, zero value otherwise.

### GetCreateAccountJSONOk

`func (o *D365Connector) GetCreateAccountJSONOk() (*string, bool)`

GetCreateAccountJSONOk returns a tuple with the CreateAccountJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateAccountJSON

`func (o *D365Connector) SetCreateAccountJSON(v string)`

SetCreateAccountJSON sets CreateAccountJSON field to given value.

### HasCreateAccountJSON

`func (o *D365Connector) HasCreateAccountJSON() bool`

HasCreateAccountJSON returns a boolean if a field has been set.

### GetUpdateAccountJSON

`func (o *D365Connector) GetUpdateAccountJSON() string`

GetUpdateAccountJSON returns the UpdateAccountJSON field if non-nil, zero value otherwise.

### GetUpdateAccountJSONOk

`func (o *D365Connector) GetUpdateAccountJSONOk() (*string, bool)`

GetUpdateAccountJSONOk returns a tuple with the UpdateAccountJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateAccountJSON

`func (o *D365Connector) SetUpdateAccountJSON(v string)`

SetUpdateAccountJSON sets UpdateAccountJSON field to given value.

### HasUpdateAccountJSON

`func (o *D365Connector) HasUpdateAccountJSON() bool`

HasUpdateAccountJSON returns a boolean if a field has been set.

### GetEnableAccountJSON

`func (o *D365Connector) GetEnableAccountJSON() string`

GetEnableAccountJSON returns the EnableAccountJSON field if non-nil, zero value otherwise.

### GetEnableAccountJSONOk

`func (o *D365Connector) GetEnableAccountJSONOk() (*string, bool)`

GetEnableAccountJSONOk returns a tuple with the EnableAccountJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableAccountJSON

`func (o *D365Connector) SetEnableAccountJSON(v string)`

SetEnableAccountJSON sets EnableAccountJSON field to given value.

### HasEnableAccountJSON

`func (o *D365Connector) HasEnableAccountJSON() bool`

HasEnableAccountJSON returns a boolean if a field has been set.

### GetDisableAccountJSON

`func (o *D365Connector) GetDisableAccountJSON() string`

GetDisableAccountJSON returns the DisableAccountJSON field if non-nil, zero value otherwise.

### GetDisableAccountJSONOk

`func (o *D365Connector) GetDisableAccountJSONOk() (*string, bool)`

GetDisableAccountJSONOk returns a tuple with the DisableAccountJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableAccountJSON

`func (o *D365Connector) SetDisableAccountJSON(v string)`

SetDisableAccountJSON sets DisableAccountJSON field to given value.

### HasDisableAccountJSON

`func (o *D365Connector) HasDisableAccountJSON() bool`

HasDisableAccountJSON returns a boolean if a field has been set.

### GetAddAccessJSON

`func (o *D365Connector) GetAddAccessJSON() string`

GetAddAccessJSON returns the AddAccessJSON field if non-nil, zero value otherwise.

### GetAddAccessJSONOk

`func (o *D365Connector) GetAddAccessJSONOk() (*string, bool)`

GetAddAccessJSONOk returns a tuple with the AddAccessJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddAccessJSON

`func (o *D365Connector) SetAddAccessJSON(v string)`

SetAddAccessJSON sets AddAccessJSON field to given value.

### HasAddAccessJSON

`func (o *D365Connector) HasAddAccessJSON() bool`

HasAddAccessJSON returns a boolean if a field has been set.

### GetRemoveAccessJSON

`func (o *D365Connector) GetRemoveAccessJSON() string`

GetRemoveAccessJSON returns the RemoveAccessJSON field if non-nil, zero value otherwise.

### GetRemoveAccessJSONOk

`func (o *D365Connector) GetRemoveAccessJSONOk() (*string, bool)`

GetRemoveAccessJSONOk returns a tuple with the RemoveAccessJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoveAccessJSON

`func (o *D365Connector) SetRemoveAccessJSON(v string)`

SetRemoveAccessJSON sets RemoveAccessJSON field to given value.

### HasRemoveAccessJSON

`func (o *D365Connector) HasRemoveAccessJSON() bool`

HasRemoveAccessJSON returns a boolean if a field has been set.

### GetRemoveAccountJSON

`func (o *D365Connector) GetRemoveAccountJSON() string`

GetRemoveAccountJSON returns the RemoveAccountJSON field if non-nil, zero value otherwise.

### GetRemoveAccountJSONOk

`func (o *D365Connector) GetRemoveAccountJSONOk() (*string, bool)`

GetRemoveAccountJSONOk returns a tuple with the RemoveAccountJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoveAccountJSON

`func (o *D365Connector) SetRemoveAccountJSON(v string)`

SetRemoveAccountJSON sets RemoveAccountJSON field to given value.

### HasRemoveAccountJSON

`func (o *D365Connector) HasRemoveAccountJSON() bool`

HasRemoveAccountJSON returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


