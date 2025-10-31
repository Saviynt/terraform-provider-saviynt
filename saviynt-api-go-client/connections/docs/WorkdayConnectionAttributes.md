# WorkdayConnectionAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**USE_OAUTH** | Pointer to **string** |  | [optional] 
**USER_IMPORT_MAPPING** | Pointer to **string** |  | [optional] 
**ACCOUNTS_LAST_IMPORT_TIME** | Pointer to **string** |  | [optional] 
**STATUS_KEY_JSON** | Pointer to **string** |  | [optional] 
**ConnectionTimeoutConfig** | Pointer to [**ConnectionTimeoutConfig**](ConnectionTimeoutConfig.md) |  | [optional] 
**ConnectionType** | Pointer to **string** |  | [optional] 
**RAAS_MAPPING_JSON** | Pointer to **string** |  | [optional] 
**ACCOUNT_IMPORT_PAYLOAD** | Pointer to **string** |  | [optional] 
**UPDATE_ACCOUNT_PAYLOAD** | Pointer to **string** |  | [optional] 
**CLIENT_ID** | Pointer to **string** |  | [optional] 
**STATUS_THRESHOLD_CONFIG** | Pointer to **string** |  | [optional] 
**USERNAME** | Pointer to **string** |  | [optional] 
**ACCESS_IMPORT_LIST** | Pointer to **string** |  | [optional] 
**IsTimeoutSupported** | Pointer to **bool** |  | [optional] 
**ACCOUNT_IMPORT_MAPPING** | Pointer to **string** |  | [optional] 
**CLIENT_SECRET** | Pointer to **string** |  | [optional] 
**ORGROLE_IMPORT_PAYLOAD** | Pointer to **string** |  | [optional] 
**ASSIGN_ORGROLE_PAYLOAD** | Pointer to **string** |  | [optional] 
**ACCESS_IMPORT_MAPPING** | Pointer to **string** |  | [optional] 
**API_VERSION** | Pointer to **string** |  | [optional] 
**REMOVE_ORGROLE_PAYLOAD** | Pointer to **string** |  | [optional] 
**INCLUDE_REFERENCE_DESCRIPTORS** | Pointer to **string** |  | [optional] 
**REFRESH_TOKEN** | Pointer to **string** |  | [optional] 
**MODIFYUSERDATAJSON** | Pointer to **string** |  | [optional] 
**IsTimeoutConfigValidated** | Pointer to **bool** |  | [optional] 
**USEX509AUTHFORSOAP** | Pointer to **string** |  | [optional] 
**REPORT_OWNER** | Pointer to **string** |  | [optional] 
**X509KEY** | Pointer to **string** |  | [optional] 
**CUSTOM_CONFIG** | Pointer to **string** |  | [optional] 
**USERATTRIBUTEJSON** | Pointer to **string** |  | [optional] 
**X509CERT** | Pointer to **string** |  | [optional] 
**PASSWORD** | Pointer to **string** |  | [optional] 
**USER_IMPORT_PAYLOAD** | Pointer to **string** |  | [optional] 
**PAM_CONFIG** | Pointer to **string** |  | [optional] 
**ACCESS_LAST_IMPORT_TIME** | Pointer to **string** |  | [optional] 
**USERS_LAST_IMPORT_TIME** | Pointer to **string** |  | [optional] 
**UPDATE_USER_PAYLOAD** | Pointer to **string** |  | [optional] 
**PAGE_SIZE** | Pointer to **string** |  | [optional] 
**TENANT_NAME** | Pointer to **string** |  | [optional] 
**USE_ENHANCED_ORGROLE** | Pointer to **string** |  | [optional] 
**CREATE_ACCOUNT_PAYLOAD** | Pointer to **string** |  | [optional] 
**BASE_URL** | Pointer to **string** |  | [optional] 

## Methods

### NewWorkdayConnectionAttributes

`func NewWorkdayConnectionAttributes() *WorkdayConnectionAttributes`

NewWorkdayConnectionAttributes instantiates a new WorkdayConnectionAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkdayConnectionAttributesWithDefaults

`func NewWorkdayConnectionAttributesWithDefaults() *WorkdayConnectionAttributes`

NewWorkdayConnectionAttributesWithDefaults instantiates a new WorkdayConnectionAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUSE_OAUTH

`func (o *WorkdayConnectionAttributes) GetUSE_OAUTH() string`

GetUSE_OAUTH returns the USE_OAUTH field if non-nil, zero value otherwise.

### GetUSE_OAUTHOk

`func (o *WorkdayConnectionAttributes) GetUSE_OAUTHOk() (*string, bool)`

GetUSE_OAUTHOk returns a tuple with the USE_OAUTH field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUSE_OAUTH

`func (o *WorkdayConnectionAttributes) SetUSE_OAUTH(v string)`

SetUSE_OAUTH sets USE_OAUTH field to given value.

### HasUSE_OAUTH

`func (o *WorkdayConnectionAttributes) HasUSE_OAUTH() bool`

HasUSE_OAUTH returns a boolean if a field has been set.

### GetUSER_IMPORT_MAPPING

`func (o *WorkdayConnectionAttributes) GetUSER_IMPORT_MAPPING() string`

GetUSER_IMPORT_MAPPING returns the USER_IMPORT_MAPPING field if non-nil, zero value otherwise.

### GetUSER_IMPORT_MAPPINGOk

`func (o *WorkdayConnectionAttributes) GetUSER_IMPORT_MAPPINGOk() (*string, bool)`

GetUSER_IMPORT_MAPPINGOk returns a tuple with the USER_IMPORT_MAPPING field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUSER_IMPORT_MAPPING

`func (o *WorkdayConnectionAttributes) SetUSER_IMPORT_MAPPING(v string)`

SetUSER_IMPORT_MAPPING sets USER_IMPORT_MAPPING field to given value.

### HasUSER_IMPORT_MAPPING

`func (o *WorkdayConnectionAttributes) HasUSER_IMPORT_MAPPING() bool`

HasUSER_IMPORT_MAPPING returns a boolean if a field has been set.

### GetACCOUNTS_LAST_IMPORT_TIME

`func (o *WorkdayConnectionAttributes) GetACCOUNTS_LAST_IMPORT_TIME() string`

GetACCOUNTS_LAST_IMPORT_TIME returns the ACCOUNTS_LAST_IMPORT_TIME field if non-nil, zero value otherwise.

### GetACCOUNTS_LAST_IMPORT_TIMEOk

`func (o *WorkdayConnectionAttributes) GetACCOUNTS_LAST_IMPORT_TIMEOk() (*string, bool)`

GetACCOUNTS_LAST_IMPORT_TIMEOk returns a tuple with the ACCOUNTS_LAST_IMPORT_TIME field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetACCOUNTS_LAST_IMPORT_TIME

`func (o *WorkdayConnectionAttributes) SetACCOUNTS_LAST_IMPORT_TIME(v string)`

SetACCOUNTS_LAST_IMPORT_TIME sets ACCOUNTS_LAST_IMPORT_TIME field to given value.

### HasACCOUNTS_LAST_IMPORT_TIME

`func (o *WorkdayConnectionAttributes) HasACCOUNTS_LAST_IMPORT_TIME() bool`

HasACCOUNTS_LAST_IMPORT_TIME returns a boolean if a field has been set.

### GetSTATUS_KEY_JSON

`func (o *WorkdayConnectionAttributes) GetSTATUS_KEY_JSON() string`

GetSTATUS_KEY_JSON returns the STATUS_KEY_JSON field if non-nil, zero value otherwise.

### GetSTATUS_KEY_JSONOk

`func (o *WorkdayConnectionAttributes) GetSTATUS_KEY_JSONOk() (*string, bool)`

GetSTATUS_KEY_JSONOk returns a tuple with the STATUS_KEY_JSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSTATUS_KEY_JSON

`func (o *WorkdayConnectionAttributes) SetSTATUS_KEY_JSON(v string)`

SetSTATUS_KEY_JSON sets STATUS_KEY_JSON field to given value.

### HasSTATUS_KEY_JSON

`func (o *WorkdayConnectionAttributes) HasSTATUS_KEY_JSON() bool`

HasSTATUS_KEY_JSON returns a boolean if a field has been set.

### GetConnectionTimeoutConfig

`func (o *WorkdayConnectionAttributes) GetConnectionTimeoutConfig() ConnectionTimeoutConfig`

GetConnectionTimeoutConfig returns the ConnectionTimeoutConfig field if non-nil, zero value otherwise.

### GetConnectionTimeoutConfigOk

`func (o *WorkdayConnectionAttributes) GetConnectionTimeoutConfigOk() (*ConnectionTimeoutConfig, bool)`

GetConnectionTimeoutConfigOk returns a tuple with the ConnectionTimeoutConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionTimeoutConfig

`func (o *WorkdayConnectionAttributes) SetConnectionTimeoutConfig(v ConnectionTimeoutConfig)`

SetConnectionTimeoutConfig sets ConnectionTimeoutConfig field to given value.

### HasConnectionTimeoutConfig

`func (o *WorkdayConnectionAttributes) HasConnectionTimeoutConfig() bool`

HasConnectionTimeoutConfig returns a boolean if a field has been set.

### GetConnectionType

`func (o *WorkdayConnectionAttributes) GetConnectionType() string`

GetConnectionType returns the ConnectionType field if non-nil, zero value otherwise.

### GetConnectionTypeOk

`func (o *WorkdayConnectionAttributes) GetConnectionTypeOk() (*string, bool)`

GetConnectionTypeOk returns a tuple with the ConnectionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionType

`func (o *WorkdayConnectionAttributes) SetConnectionType(v string)`

SetConnectionType sets ConnectionType field to given value.

### HasConnectionType

`func (o *WorkdayConnectionAttributes) HasConnectionType() bool`

HasConnectionType returns a boolean if a field has been set.

### GetRAAS_MAPPING_JSON

`func (o *WorkdayConnectionAttributes) GetRAAS_MAPPING_JSON() string`

GetRAAS_MAPPING_JSON returns the RAAS_MAPPING_JSON field if non-nil, zero value otherwise.

### GetRAAS_MAPPING_JSONOk

`func (o *WorkdayConnectionAttributes) GetRAAS_MAPPING_JSONOk() (*string, bool)`

GetRAAS_MAPPING_JSONOk returns a tuple with the RAAS_MAPPING_JSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRAAS_MAPPING_JSON

`func (o *WorkdayConnectionAttributes) SetRAAS_MAPPING_JSON(v string)`

SetRAAS_MAPPING_JSON sets RAAS_MAPPING_JSON field to given value.

### HasRAAS_MAPPING_JSON

`func (o *WorkdayConnectionAttributes) HasRAAS_MAPPING_JSON() bool`

HasRAAS_MAPPING_JSON returns a boolean if a field has been set.

### GetACCOUNT_IMPORT_PAYLOAD

`func (o *WorkdayConnectionAttributes) GetACCOUNT_IMPORT_PAYLOAD() string`

GetACCOUNT_IMPORT_PAYLOAD returns the ACCOUNT_IMPORT_PAYLOAD field if non-nil, zero value otherwise.

### GetACCOUNT_IMPORT_PAYLOADOk

`func (o *WorkdayConnectionAttributes) GetACCOUNT_IMPORT_PAYLOADOk() (*string, bool)`

GetACCOUNT_IMPORT_PAYLOADOk returns a tuple with the ACCOUNT_IMPORT_PAYLOAD field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetACCOUNT_IMPORT_PAYLOAD

`func (o *WorkdayConnectionAttributes) SetACCOUNT_IMPORT_PAYLOAD(v string)`

SetACCOUNT_IMPORT_PAYLOAD sets ACCOUNT_IMPORT_PAYLOAD field to given value.

### HasACCOUNT_IMPORT_PAYLOAD

`func (o *WorkdayConnectionAttributes) HasACCOUNT_IMPORT_PAYLOAD() bool`

HasACCOUNT_IMPORT_PAYLOAD returns a boolean if a field has been set.

### GetUPDATE_ACCOUNT_PAYLOAD

`func (o *WorkdayConnectionAttributes) GetUPDATE_ACCOUNT_PAYLOAD() string`

GetUPDATE_ACCOUNT_PAYLOAD returns the UPDATE_ACCOUNT_PAYLOAD field if non-nil, zero value otherwise.

### GetUPDATE_ACCOUNT_PAYLOADOk

`func (o *WorkdayConnectionAttributes) GetUPDATE_ACCOUNT_PAYLOADOk() (*string, bool)`

GetUPDATE_ACCOUNT_PAYLOADOk returns a tuple with the UPDATE_ACCOUNT_PAYLOAD field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUPDATE_ACCOUNT_PAYLOAD

`func (o *WorkdayConnectionAttributes) SetUPDATE_ACCOUNT_PAYLOAD(v string)`

SetUPDATE_ACCOUNT_PAYLOAD sets UPDATE_ACCOUNT_PAYLOAD field to given value.

### HasUPDATE_ACCOUNT_PAYLOAD

`func (o *WorkdayConnectionAttributes) HasUPDATE_ACCOUNT_PAYLOAD() bool`

HasUPDATE_ACCOUNT_PAYLOAD returns a boolean if a field has been set.

### GetCLIENT_ID

`func (o *WorkdayConnectionAttributes) GetCLIENT_ID() string`

GetCLIENT_ID returns the CLIENT_ID field if non-nil, zero value otherwise.

### GetCLIENT_IDOk

`func (o *WorkdayConnectionAttributes) GetCLIENT_IDOk() (*string, bool)`

GetCLIENT_IDOk returns a tuple with the CLIENT_ID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCLIENT_ID

`func (o *WorkdayConnectionAttributes) SetCLIENT_ID(v string)`

SetCLIENT_ID sets CLIENT_ID field to given value.

### HasCLIENT_ID

`func (o *WorkdayConnectionAttributes) HasCLIENT_ID() bool`

HasCLIENT_ID returns a boolean if a field has been set.

### GetSTATUS_THRESHOLD_CONFIG

`func (o *WorkdayConnectionAttributes) GetSTATUS_THRESHOLD_CONFIG() string`

GetSTATUS_THRESHOLD_CONFIG returns the STATUS_THRESHOLD_CONFIG field if non-nil, zero value otherwise.

### GetSTATUS_THRESHOLD_CONFIGOk

`func (o *WorkdayConnectionAttributes) GetSTATUS_THRESHOLD_CONFIGOk() (*string, bool)`

GetSTATUS_THRESHOLD_CONFIGOk returns a tuple with the STATUS_THRESHOLD_CONFIG field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSTATUS_THRESHOLD_CONFIG

`func (o *WorkdayConnectionAttributes) SetSTATUS_THRESHOLD_CONFIG(v string)`

SetSTATUS_THRESHOLD_CONFIG sets STATUS_THRESHOLD_CONFIG field to given value.

### HasSTATUS_THRESHOLD_CONFIG

`func (o *WorkdayConnectionAttributes) HasSTATUS_THRESHOLD_CONFIG() bool`

HasSTATUS_THRESHOLD_CONFIG returns a boolean if a field has been set.

### GetUSERNAME

`func (o *WorkdayConnectionAttributes) GetUSERNAME() string`

GetUSERNAME returns the USERNAME field if non-nil, zero value otherwise.

### GetUSERNAMEOk

`func (o *WorkdayConnectionAttributes) GetUSERNAMEOk() (*string, bool)`

GetUSERNAMEOk returns a tuple with the USERNAME field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUSERNAME

`func (o *WorkdayConnectionAttributes) SetUSERNAME(v string)`

SetUSERNAME sets USERNAME field to given value.

### HasUSERNAME

`func (o *WorkdayConnectionAttributes) HasUSERNAME() bool`

HasUSERNAME returns a boolean if a field has been set.

### GetACCESS_IMPORT_LIST

`func (o *WorkdayConnectionAttributes) GetACCESS_IMPORT_LIST() string`

GetACCESS_IMPORT_LIST returns the ACCESS_IMPORT_LIST field if non-nil, zero value otherwise.

### GetACCESS_IMPORT_LISTOk

`func (o *WorkdayConnectionAttributes) GetACCESS_IMPORT_LISTOk() (*string, bool)`

GetACCESS_IMPORT_LISTOk returns a tuple with the ACCESS_IMPORT_LIST field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetACCESS_IMPORT_LIST

`func (o *WorkdayConnectionAttributes) SetACCESS_IMPORT_LIST(v string)`

SetACCESS_IMPORT_LIST sets ACCESS_IMPORT_LIST field to given value.

### HasACCESS_IMPORT_LIST

`func (o *WorkdayConnectionAttributes) HasACCESS_IMPORT_LIST() bool`

HasACCESS_IMPORT_LIST returns a boolean if a field has been set.

### GetIsTimeoutSupported

`func (o *WorkdayConnectionAttributes) GetIsTimeoutSupported() bool`

GetIsTimeoutSupported returns the IsTimeoutSupported field if non-nil, zero value otherwise.

### GetIsTimeoutSupportedOk

`func (o *WorkdayConnectionAttributes) GetIsTimeoutSupportedOk() (*bool, bool)`

GetIsTimeoutSupportedOk returns a tuple with the IsTimeoutSupported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsTimeoutSupported

`func (o *WorkdayConnectionAttributes) SetIsTimeoutSupported(v bool)`

SetIsTimeoutSupported sets IsTimeoutSupported field to given value.

### HasIsTimeoutSupported

`func (o *WorkdayConnectionAttributes) HasIsTimeoutSupported() bool`

HasIsTimeoutSupported returns a boolean if a field has been set.

### GetACCOUNT_IMPORT_MAPPING

`func (o *WorkdayConnectionAttributes) GetACCOUNT_IMPORT_MAPPING() string`

GetACCOUNT_IMPORT_MAPPING returns the ACCOUNT_IMPORT_MAPPING field if non-nil, zero value otherwise.

### GetACCOUNT_IMPORT_MAPPINGOk

`func (o *WorkdayConnectionAttributes) GetACCOUNT_IMPORT_MAPPINGOk() (*string, bool)`

GetACCOUNT_IMPORT_MAPPINGOk returns a tuple with the ACCOUNT_IMPORT_MAPPING field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetACCOUNT_IMPORT_MAPPING

`func (o *WorkdayConnectionAttributes) SetACCOUNT_IMPORT_MAPPING(v string)`

SetACCOUNT_IMPORT_MAPPING sets ACCOUNT_IMPORT_MAPPING field to given value.

### HasACCOUNT_IMPORT_MAPPING

`func (o *WorkdayConnectionAttributes) HasACCOUNT_IMPORT_MAPPING() bool`

HasACCOUNT_IMPORT_MAPPING returns a boolean if a field has been set.

### GetCLIENT_SECRET

`func (o *WorkdayConnectionAttributes) GetCLIENT_SECRET() string`

GetCLIENT_SECRET returns the CLIENT_SECRET field if non-nil, zero value otherwise.

### GetCLIENT_SECRETOk

`func (o *WorkdayConnectionAttributes) GetCLIENT_SECRETOk() (*string, bool)`

GetCLIENT_SECRETOk returns a tuple with the CLIENT_SECRET field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCLIENT_SECRET

`func (o *WorkdayConnectionAttributes) SetCLIENT_SECRET(v string)`

SetCLIENT_SECRET sets CLIENT_SECRET field to given value.

### HasCLIENT_SECRET

`func (o *WorkdayConnectionAttributes) HasCLIENT_SECRET() bool`

HasCLIENT_SECRET returns a boolean if a field has been set.

### GetORGROLE_IMPORT_PAYLOAD

`func (o *WorkdayConnectionAttributes) GetORGROLE_IMPORT_PAYLOAD() string`

GetORGROLE_IMPORT_PAYLOAD returns the ORGROLE_IMPORT_PAYLOAD field if non-nil, zero value otherwise.

### GetORGROLE_IMPORT_PAYLOADOk

`func (o *WorkdayConnectionAttributes) GetORGROLE_IMPORT_PAYLOADOk() (*string, bool)`

GetORGROLE_IMPORT_PAYLOADOk returns a tuple with the ORGROLE_IMPORT_PAYLOAD field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetORGROLE_IMPORT_PAYLOAD

`func (o *WorkdayConnectionAttributes) SetORGROLE_IMPORT_PAYLOAD(v string)`

SetORGROLE_IMPORT_PAYLOAD sets ORGROLE_IMPORT_PAYLOAD field to given value.

### HasORGROLE_IMPORT_PAYLOAD

`func (o *WorkdayConnectionAttributes) HasORGROLE_IMPORT_PAYLOAD() bool`

HasORGROLE_IMPORT_PAYLOAD returns a boolean if a field has been set.

### GetASSIGN_ORGROLE_PAYLOAD

`func (o *WorkdayConnectionAttributes) GetASSIGN_ORGROLE_PAYLOAD() string`

GetASSIGN_ORGROLE_PAYLOAD returns the ASSIGN_ORGROLE_PAYLOAD field if non-nil, zero value otherwise.

### GetASSIGN_ORGROLE_PAYLOADOk

`func (o *WorkdayConnectionAttributes) GetASSIGN_ORGROLE_PAYLOADOk() (*string, bool)`

GetASSIGN_ORGROLE_PAYLOADOk returns a tuple with the ASSIGN_ORGROLE_PAYLOAD field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetASSIGN_ORGROLE_PAYLOAD

`func (o *WorkdayConnectionAttributes) SetASSIGN_ORGROLE_PAYLOAD(v string)`

SetASSIGN_ORGROLE_PAYLOAD sets ASSIGN_ORGROLE_PAYLOAD field to given value.

### HasASSIGN_ORGROLE_PAYLOAD

`func (o *WorkdayConnectionAttributes) HasASSIGN_ORGROLE_PAYLOAD() bool`

HasASSIGN_ORGROLE_PAYLOAD returns a boolean if a field has been set.

### GetACCESS_IMPORT_MAPPING

`func (o *WorkdayConnectionAttributes) GetACCESS_IMPORT_MAPPING() string`

GetACCESS_IMPORT_MAPPING returns the ACCESS_IMPORT_MAPPING field if non-nil, zero value otherwise.

### GetACCESS_IMPORT_MAPPINGOk

`func (o *WorkdayConnectionAttributes) GetACCESS_IMPORT_MAPPINGOk() (*string, bool)`

GetACCESS_IMPORT_MAPPINGOk returns a tuple with the ACCESS_IMPORT_MAPPING field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetACCESS_IMPORT_MAPPING

`func (o *WorkdayConnectionAttributes) SetACCESS_IMPORT_MAPPING(v string)`

SetACCESS_IMPORT_MAPPING sets ACCESS_IMPORT_MAPPING field to given value.

### HasACCESS_IMPORT_MAPPING

`func (o *WorkdayConnectionAttributes) HasACCESS_IMPORT_MAPPING() bool`

HasACCESS_IMPORT_MAPPING returns a boolean if a field has been set.

### GetAPI_VERSION

`func (o *WorkdayConnectionAttributes) GetAPI_VERSION() string`

GetAPI_VERSION returns the API_VERSION field if non-nil, zero value otherwise.

### GetAPI_VERSIONOk

`func (o *WorkdayConnectionAttributes) GetAPI_VERSIONOk() (*string, bool)`

GetAPI_VERSIONOk returns a tuple with the API_VERSION field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAPI_VERSION

`func (o *WorkdayConnectionAttributes) SetAPI_VERSION(v string)`

SetAPI_VERSION sets API_VERSION field to given value.

### HasAPI_VERSION

`func (o *WorkdayConnectionAttributes) HasAPI_VERSION() bool`

HasAPI_VERSION returns a boolean if a field has been set.

### GetREMOVE_ORGROLE_PAYLOAD

`func (o *WorkdayConnectionAttributes) GetREMOVE_ORGROLE_PAYLOAD() string`

GetREMOVE_ORGROLE_PAYLOAD returns the REMOVE_ORGROLE_PAYLOAD field if non-nil, zero value otherwise.

### GetREMOVE_ORGROLE_PAYLOADOk

`func (o *WorkdayConnectionAttributes) GetREMOVE_ORGROLE_PAYLOADOk() (*string, bool)`

GetREMOVE_ORGROLE_PAYLOADOk returns a tuple with the REMOVE_ORGROLE_PAYLOAD field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetREMOVE_ORGROLE_PAYLOAD

`func (o *WorkdayConnectionAttributes) SetREMOVE_ORGROLE_PAYLOAD(v string)`

SetREMOVE_ORGROLE_PAYLOAD sets REMOVE_ORGROLE_PAYLOAD field to given value.

### HasREMOVE_ORGROLE_PAYLOAD

`func (o *WorkdayConnectionAttributes) HasREMOVE_ORGROLE_PAYLOAD() bool`

HasREMOVE_ORGROLE_PAYLOAD returns a boolean if a field has been set.

### GetINCLUDE_REFERENCE_DESCRIPTORS

`func (o *WorkdayConnectionAttributes) GetINCLUDE_REFERENCE_DESCRIPTORS() string`

GetINCLUDE_REFERENCE_DESCRIPTORS returns the INCLUDE_REFERENCE_DESCRIPTORS field if non-nil, zero value otherwise.

### GetINCLUDE_REFERENCE_DESCRIPTORSOk

`func (o *WorkdayConnectionAttributes) GetINCLUDE_REFERENCE_DESCRIPTORSOk() (*string, bool)`

GetINCLUDE_REFERENCE_DESCRIPTORSOk returns a tuple with the INCLUDE_REFERENCE_DESCRIPTORS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetINCLUDE_REFERENCE_DESCRIPTORS

`func (o *WorkdayConnectionAttributes) SetINCLUDE_REFERENCE_DESCRIPTORS(v string)`

SetINCLUDE_REFERENCE_DESCRIPTORS sets INCLUDE_REFERENCE_DESCRIPTORS field to given value.

### HasINCLUDE_REFERENCE_DESCRIPTORS

`func (o *WorkdayConnectionAttributes) HasINCLUDE_REFERENCE_DESCRIPTORS() bool`

HasINCLUDE_REFERENCE_DESCRIPTORS returns a boolean if a field has been set.

### GetREFRESH_TOKEN

`func (o *WorkdayConnectionAttributes) GetREFRESH_TOKEN() string`

GetREFRESH_TOKEN returns the REFRESH_TOKEN field if non-nil, zero value otherwise.

### GetREFRESH_TOKENOk

`func (o *WorkdayConnectionAttributes) GetREFRESH_TOKENOk() (*string, bool)`

GetREFRESH_TOKENOk returns a tuple with the REFRESH_TOKEN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetREFRESH_TOKEN

`func (o *WorkdayConnectionAttributes) SetREFRESH_TOKEN(v string)`

SetREFRESH_TOKEN sets REFRESH_TOKEN field to given value.

### HasREFRESH_TOKEN

`func (o *WorkdayConnectionAttributes) HasREFRESH_TOKEN() bool`

HasREFRESH_TOKEN returns a boolean if a field has been set.

### GetMODIFYUSERDATAJSON

`func (o *WorkdayConnectionAttributes) GetMODIFYUSERDATAJSON() string`

GetMODIFYUSERDATAJSON returns the MODIFYUSERDATAJSON field if non-nil, zero value otherwise.

### GetMODIFYUSERDATAJSONOk

`func (o *WorkdayConnectionAttributes) GetMODIFYUSERDATAJSONOk() (*string, bool)`

GetMODIFYUSERDATAJSONOk returns a tuple with the MODIFYUSERDATAJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMODIFYUSERDATAJSON

`func (o *WorkdayConnectionAttributes) SetMODIFYUSERDATAJSON(v string)`

SetMODIFYUSERDATAJSON sets MODIFYUSERDATAJSON field to given value.

### HasMODIFYUSERDATAJSON

`func (o *WorkdayConnectionAttributes) HasMODIFYUSERDATAJSON() bool`

HasMODIFYUSERDATAJSON returns a boolean if a field has been set.

### GetIsTimeoutConfigValidated

`func (o *WorkdayConnectionAttributes) GetIsTimeoutConfigValidated() bool`

GetIsTimeoutConfigValidated returns the IsTimeoutConfigValidated field if non-nil, zero value otherwise.

### GetIsTimeoutConfigValidatedOk

`func (o *WorkdayConnectionAttributes) GetIsTimeoutConfigValidatedOk() (*bool, bool)`

GetIsTimeoutConfigValidatedOk returns a tuple with the IsTimeoutConfigValidated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsTimeoutConfigValidated

`func (o *WorkdayConnectionAttributes) SetIsTimeoutConfigValidated(v bool)`

SetIsTimeoutConfigValidated sets IsTimeoutConfigValidated field to given value.

### HasIsTimeoutConfigValidated

`func (o *WorkdayConnectionAttributes) HasIsTimeoutConfigValidated() bool`

HasIsTimeoutConfigValidated returns a boolean if a field has been set.

### GetUSEX509AUTHFORSOAP

`func (o *WorkdayConnectionAttributes) GetUSEX509AUTHFORSOAP() string`

GetUSEX509AUTHFORSOAP returns the USEX509AUTHFORSOAP field if non-nil, zero value otherwise.

### GetUSEX509AUTHFORSOAPOk

`func (o *WorkdayConnectionAttributes) GetUSEX509AUTHFORSOAPOk() (*string, bool)`

GetUSEX509AUTHFORSOAPOk returns a tuple with the USEX509AUTHFORSOAP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUSEX509AUTHFORSOAP

`func (o *WorkdayConnectionAttributes) SetUSEX509AUTHFORSOAP(v string)`

SetUSEX509AUTHFORSOAP sets USEX509AUTHFORSOAP field to given value.

### HasUSEX509AUTHFORSOAP

`func (o *WorkdayConnectionAttributes) HasUSEX509AUTHFORSOAP() bool`

HasUSEX509AUTHFORSOAP returns a boolean if a field has been set.

### GetREPORT_OWNER

`func (o *WorkdayConnectionAttributes) GetREPORT_OWNER() string`

GetREPORT_OWNER returns the REPORT_OWNER field if non-nil, zero value otherwise.

### GetREPORT_OWNEROk

`func (o *WorkdayConnectionAttributes) GetREPORT_OWNEROk() (*string, bool)`

GetREPORT_OWNEROk returns a tuple with the REPORT_OWNER field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetREPORT_OWNER

`func (o *WorkdayConnectionAttributes) SetREPORT_OWNER(v string)`

SetREPORT_OWNER sets REPORT_OWNER field to given value.

### HasREPORT_OWNER

`func (o *WorkdayConnectionAttributes) HasREPORT_OWNER() bool`

HasREPORT_OWNER returns a boolean if a field has been set.

### GetX509KEY

`func (o *WorkdayConnectionAttributes) GetX509KEY() string`

GetX509KEY returns the X509KEY field if non-nil, zero value otherwise.

### GetX509KEYOk

`func (o *WorkdayConnectionAttributes) GetX509KEYOk() (*string, bool)`

GetX509KEYOk returns a tuple with the X509KEY field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetX509KEY

`func (o *WorkdayConnectionAttributes) SetX509KEY(v string)`

SetX509KEY sets X509KEY field to given value.

### HasX509KEY

`func (o *WorkdayConnectionAttributes) HasX509KEY() bool`

HasX509KEY returns a boolean if a field has been set.

### GetCUSTOM_CONFIG

`func (o *WorkdayConnectionAttributes) GetCUSTOM_CONFIG() string`

GetCUSTOM_CONFIG returns the CUSTOM_CONFIG field if non-nil, zero value otherwise.

### GetCUSTOM_CONFIGOk

`func (o *WorkdayConnectionAttributes) GetCUSTOM_CONFIGOk() (*string, bool)`

GetCUSTOM_CONFIGOk returns a tuple with the CUSTOM_CONFIG field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCUSTOM_CONFIG

`func (o *WorkdayConnectionAttributes) SetCUSTOM_CONFIG(v string)`

SetCUSTOM_CONFIG sets CUSTOM_CONFIG field to given value.

### HasCUSTOM_CONFIG

`func (o *WorkdayConnectionAttributes) HasCUSTOM_CONFIG() bool`

HasCUSTOM_CONFIG returns a boolean if a field has been set.

### GetUSERATTRIBUTEJSON

`func (o *WorkdayConnectionAttributes) GetUSERATTRIBUTEJSON() string`

GetUSERATTRIBUTEJSON returns the USERATTRIBUTEJSON field if non-nil, zero value otherwise.

### GetUSERATTRIBUTEJSONOk

`func (o *WorkdayConnectionAttributes) GetUSERATTRIBUTEJSONOk() (*string, bool)`

GetUSERATTRIBUTEJSONOk returns a tuple with the USERATTRIBUTEJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUSERATTRIBUTEJSON

`func (o *WorkdayConnectionAttributes) SetUSERATTRIBUTEJSON(v string)`

SetUSERATTRIBUTEJSON sets USERATTRIBUTEJSON field to given value.

### HasUSERATTRIBUTEJSON

`func (o *WorkdayConnectionAttributes) HasUSERATTRIBUTEJSON() bool`

HasUSERATTRIBUTEJSON returns a boolean if a field has been set.

### GetX509CERT

`func (o *WorkdayConnectionAttributes) GetX509CERT() string`

GetX509CERT returns the X509CERT field if non-nil, zero value otherwise.

### GetX509CERTOk

`func (o *WorkdayConnectionAttributes) GetX509CERTOk() (*string, bool)`

GetX509CERTOk returns a tuple with the X509CERT field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetX509CERT

`func (o *WorkdayConnectionAttributes) SetX509CERT(v string)`

SetX509CERT sets X509CERT field to given value.

### HasX509CERT

`func (o *WorkdayConnectionAttributes) HasX509CERT() bool`

HasX509CERT returns a boolean if a field has been set.

### GetPASSWORD

`func (o *WorkdayConnectionAttributes) GetPASSWORD() string`

GetPASSWORD returns the PASSWORD field if non-nil, zero value otherwise.

### GetPASSWORDOk

`func (o *WorkdayConnectionAttributes) GetPASSWORDOk() (*string, bool)`

GetPASSWORDOk returns a tuple with the PASSWORD field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPASSWORD

`func (o *WorkdayConnectionAttributes) SetPASSWORD(v string)`

SetPASSWORD sets PASSWORD field to given value.

### HasPASSWORD

`func (o *WorkdayConnectionAttributes) HasPASSWORD() bool`

HasPASSWORD returns a boolean if a field has been set.

### GetUSER_IMPORT_PAYLOAD

`func (o *WorkdayConnectionAttributes) GetUSER_IMPORT_PAYLOAD() string`

GetUSER_IMPORT_PAYLOAD returns the USER_IMPORT_PAYLOAD field if non-nil, zero value otherwise.

### GetUSER_IMPORT_PAYLOADOk

`func (o *WorkdayConnectionAttributes) GetUSER_IMPORT_PAYLOADOk() (*string, bool)`

GetUSER_IMPORT_PAYLOADOk returns a tuple with the USER_IMPORT_PAYLOAD field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUSER_IMPORT_PAYLOAD

`func (o *WorkdayConnectionAttributes) SetUSER_IMPORT_PAYLOAD(v string)`

SetUSER_IMPORT_PAYLOAD sets USER_IMPORT_PAYLOAD field to given value.

### HasUSER_IMPORT_PAYLOAD

`func (o *WorkdayConnectionAttributes) HasUSER_IMPORT_PAYLOAD() bool`

HasUSER_IMPORT_PAYLOAD returns a boolean if a field has been set.

### GetPAM_CONFIG

`func (o *WorkdayConnectionAttributes) GetPAM_CONFIG() string`

GetPAM_CONFIG returns the PAM_CONFIG field if non-nil, zero value otherwise.

### GetPAM_CONFIGOk

`func (o *WorkdayConnectionAttributes) GetPAM_CONFIGOk() (*string, bool)`

GetPAM_CONFIGOk returns a tuple with the PAM_CONFIG field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPAM_CONFIG

`func (o *WorkdayConnectionAttributes) SetPAM_CONFIG(v string)`

SetPAM_CONFIG sets PAM_CONFIG field to given value.

### HasPAM_CONFIG

`func (o *WorkdayConnectionAttributes) HasPAM_CONFIG() bool`

HasPAM_CONFIG returns a boolean if a field has been set.

### GetACCESS_LAST_IMPORT_TIME

`func (o *WorkdayConnectionAttributes) GetACCESS_LAST_IMPORT_TIME() string`

GetACCESS_LAST_IMPORT_TIME returns the ACCESS_LAST_IMPORT_TIME field if non-nil, zero value otherwise.

### GetACCESS_LAST_IMPORT_TIMEOk

`func (o *WorkdayConnectionAttributes) GetACCESS_LAST_IMPORT_TIMEOk() (*string, bool)`

GetACCESS_LAST_IMPORT_TIMEOk returns a tuple with the ACCESS_LAST_IMPORT_TIME field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetACCESS_LAST_IMPORT_TIME

`func (o *WorkdayConnectionAttributes) SetACCESS_LAST_IMPORT_TIME(v string)`

SetACCESS_LAST_IMPORT_TIME sets ACCESS_LAST_IMPORT_TIME field to given value.

### HasACCESS_LAST_IMPORT_TIME

`func (o *WorkdayConnectionAttributes) HasACCESS_LAST_IMPORT_TIME() bool`

HasACCESS_LAST_IMPORT_TIME returns a boolean if a field has been set.

### GetUSERS_LAST_IMPORT_TIME

`func (o *WorkdayConnectionAttributes) GetUSERS_LAST_IMPORT_TIME() string`

GetUSERS_LAST_IMPORT_TIME returns the USERS_LAST_IMPORT_TIME field if non-nil, zero value otherwise.

### GetUSERS_LAST_IMPORT_TIMEOk

`func (o *WorkdayConnectionAttributes) GetUSERS_LAST_IMPORT_TIMEOk() (*string, bool)`

GetUSERS_LAST_IMPORT_TIMEOk returns a tuple with the USERS_LAST_IMPORT_TIME field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUSERS_LAST_IMPORT_TIME

`func (o *WorkdayConnectionAttributes) SetUSERS_LAST_IMPORT_TIME(v string)`

SetUSERS_LAST_IMPORT_TIME sets USERS_LAST_IMPORT_TIME field to given value.

### HasUSERS_LAST_IMPORT_TIME

`func (o *WorkdayConnectionAttributes) HasUSERS_LAST_IMPORT_TIME() bool`

HasUSERS_LAST_IMPORT_TIME returns a boolean if a field has been set.

### GetUPDATE_USER_PAYLOAD

`func (o *WorkdayConnectionAttributes) GetUPDATE_USER_PAYLOAD() string`

GetUPDATE_USER_PAYLOAD returns the UPDATE_USER_PAYLOAD field if non-nil, zero value otherwise.

### GetUPDATE_USER_PAYLOADOk

`func (o *WorkdayConnectionAttributes) GetUPDATE_USER_PAYLOADOk() (*string, bool)`

GetUPDATE_USER_PAYLOADOk returns a tuple with the UPDATE_USER_PAYLOAD field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUPDATE_USER_PAYLOAD

`func (o *WorkdayConnectionAttributes) SetUPDATE_USER_PAYLOAD(v string)`

SetUPDATE_USER_PAYLOAD sets UPDATE_USER_PAYLOAD field to given value.

### HasUPDATE_USER_PAYLOAD

`func (o *WorkdayConnectionAttributes) HasUPDATE_USER_PAYLOAD() bool`

HasUPDATE_USER_PAYLOAD returns a boolean if a field has been set.

### GetPAGE_SIZE

`func (o *WorkdayConnectionAttributes) GetPAGE_SIZE() string`

GetPAGE_SIZE returns the PAGE_SIZE field if non-nil, zero value otherwise.

### GetPAGE_SIZEOk

`func (o *WorkdayConnectionAttributes) GetPAGE_SIZEOk() (*string, bool)`

GetPAGE_SIZEOk returns a tuple with the PAGE_SIZE field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPAGE_SIZE

`func (o *WorkdayConnectionAttributes) SetPAGE_SIZE(v string)`

SetPAGE_SIZE sets PAGE_SIZE field to given value.

### HasPAGE_SIZE

`func (o *WorkdayConnectionAttributes) HasPAGE_SIZE() bool`

HasPAGE_SIZE returns a boolean if a field has been set.

### GetTENANT_NAME

`func (o *WorkdayConnectionAttributes) GetTENANT_NAME() string`

GetTENANT_NAME returns the TENANT_NAME field if non-nil, zero value otherwise.

### GetTENANT_NAMEOk

`func (o *WorkdayConnectionAttributes) GetTENANT_NAMEOk() (*string, bool)`

GetTENANT_NAMEOk returns a tuple with the TENANT_NAME field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTENANT_NAME

`func (o *WorkdayConnectionAttributes) SetTENANT_NAME(v string)`

SetTENANT_NAME sets TENANT_NAME field to given value.

### HasTENANT_NAME

`func (o *WorkdayConnectionAttributes) HasTENANT_NAME() bool`

HasTENANT_NAME returns a boolean if a field has been set.

### GetUSE_ENHANCED_ORGROLE

`func (o *WorkdayConnectionAttributes) GetUSE_ENHANCED_ORGROLE() string`

GetUSE_ENHANCED_ORGROLE returns the USE_ENHANCED_ORGROLE field if non-nil, zero value otherwise.

### GetUSE_ENHANCED_ORGROLEOk

`func (o *WorkdayConnectionAttributes) GetUSE_ENHANCED_ORGROLEOk() (*string, bool)`

GetUSE_ENHANCED_ORGROLEOk returns a tuple with the USE_ENHANCED_ORGROLE field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUSE_ENHANCED_ORGROLE

`func (o *WorkdayConnectionAttributes) SetUSE_ENHANCED_ORGROLE(v string)`

SetUSE_ENHANCED_ORGROLE sets USE_ENHANCED_ORGROLE field to given value.

### HasUSE_ENHANCED_ORGROLE

`func (o *WorkdayConnectionAttributes) HasUSE_ENHANCED_ORGROLE() bool`

HasUSE_ENHANCED_ORGROLE returns a boolean if a field has been set.

### GetCREATE_ACCOUNT_PAYLOAD

`func (o *WorkdayConnectionAttributes) GetCREATE_ACCOUNT_PAYLOAD() string`

GetCREATE_ACCOUNT_PAYLOAD returns the CREATE_ACCOUNT_PAYLOAD field if non-nil, zero value otherwise.

### GetCREATE_ACCOUNT_PAYLOADOk

`func (o *WorkdayConnectionAttributes) GetCREATE_ACCOUNT_PAYLOADOk() (*string, bool)`

GetCREATE_ACCOUNT_PAYLOADOk returns a tuple with the CREATE_ACCOUNT_PAYLOAD field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCREATE_ACCOUNT_PAYLOAD

`func (o *WorkdayConnectionAttributes) SetCREATE_ACCOUNT_PAYLOAD(v string)`

SetCREATE_ACCOUNT_PAYLOAD sets CREATE_ACCOUNT_PAYLOAD field to given value.

### HasCREATE_ACCOUNT_PAYLOAD

`func (o *WorkdayConnectionAttributes) HasCREATE_ACCOUNT_PAYLOAD() bool`

HasCREATE_ACCOUNT_PAYLOAD returns a boolean if a field has been set.

### GetBASE_URL

`func (o *WorkdayConnectionAttributes) GetBASE_URL() string`

GetBASE_URL returns the BASE_URL field if non-nil, zero value otherwise.

### GetBASE_URLOk

`func (o *WorkdayConnectionAttributes) GetBASE_URLOk() (*string, bool)`

GetBASE_URLOk returns a tuple with the BASE_URL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBASE_URL

`func (o *WorkdayConnectionAttributes) SetBASE_URL(v string)`

SetBASE_URL sets BASE_URL field to given value.

### HasBASE_URL

`func (o *WorkdayConnectionAttributes) HasBASE_URL() bool`

HasBASE_URL returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


