# UNIXConnectionAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GROUPS_FILE** | Pointer to **string** |  | [optional] 
**SSH_KEY** | Pointer to **string** |  | [optional] 
**ACCOUNT_ENTITLEMENT_MAPPING_COMMAND** | Pointer to **string** |  | [optional] 
**REMOVE_ACCESS_COMMAND** | Pointer to **string** |  | [optional] 
**PEM_KEY_FILE** | Pointer to **string** |  | [optional] 
**PassThroughConnectionDetails** | Pointer to **string** |  | [optional] 
**DISABLE_ACCOUNT_COMMAND** | Pointer to **string** |  | [optional] 
**ConnectionTimeoutConfig** | Pointer to [**ConnectionTimeoutConfig**](ConnectionTimeoutConfig.md) |  | [optional] 
**PORT_NUMBER** | Pointer to **string** |  | [optional] 
**ConnectionType** | Pointer to **string** |  | [optional] 
**CREATE_GROUP_COMMAND** | Pointer to **string** |  | [optional] 
**ACCOUNTS_FILE** | Pointer to **string** |  | [optional] 
**PASSPHRASE** | Pointer to **string** |  | [optional] 
**DELETE_GROUP_COMMAND** | Pointer to **string** |  | [optional] 
**HOST_NAME** | Pointer to **string** |  | [optional] 
**ADD_GROUP_OWNER_COMMAND** | Pointer to **string** |  | [optional] 
**STATUS_THRESHOLD_CONFIG** | Pointer to **string** |  | [optional] 
**USERNAME** | Pointer to **string** |  | [optional] 
**INACTIVE_LOCK_ACCOUNT** | Pointer to **string** |  | [optional] 
**ADD_ACCESS_COMMAND** | Pointer to **string** |  | [optional] 
**UPDATE_ACCOUNT_COMMAND** | Pointer to **string** |  | [optional] 
**SSHPassThroughPassphrase** | Pointer to **string** |  | [optional] 
**SHADOW_FILE** | Pointer to **string** |  | [optional] 
**IsTimeoutSupported** | Pointer to **bool** |  | [optional] 
**SSHPassThroughSSHKEY** | Pointer to **string** |  | [optional] 
**PROVISION_ACCOUNT_COMMAND** | Pointer to **string** |  | [optional] 
**FIREFIGHTERID_GRANT_ACCESS_COMMAND** | Pointer to **string** |  | [optional] 
**UNLOCK_ACCOUNT_COMMAND** | Pointer to **string** |  | [optional] 
**DEPROVISION_ACCOUNT_COMMAND** | Pointer to **string** |  | [optional] 
**CHANGE_PASSWRD_JSON** | Pointer to **string** |  | [optional] 
**SSHPassThroughPassword** | Pointer to **string** |  | [optional] 
**FIREFIGHTERID_REVOKE_ACCESS_COMMAND** | Pointer to **string** |  | [optional] 
**ADD_PRIMARY_GROUP_COMMAND** | Pointer to **string** |  | [optional] 
**IsTimeoutConfigValidated** | Pointer to **bool** |  | [optional] 
**LOCK_ACCOUNT_COMMAND** | Pointer to **string** |  | [optional] 
**PASSWORD** | Pointer to **string** |  | [optional] 
**CUSTOM_CONFIG_JSON** | Pointer to **string** |  | [optional] 
**ENABLE_ACCOUNT_COMMAND** | Pointer to **string** |  | [optional] 
**SERVER_TYPE** | Pointer to **string** |  | [optional] 

## Methods

### NewUNIXConnectionAttributes

`func NewUNIXConnectionAttributes() *UNIXConnectionAttributes`

NewUNIXConnectionAttributes instantiates a new UNIXConnectionAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUNIXConnectionAttributesWithDefaults

`func NewUNIXConnectionAttributesWithDefaults() *UNIXConnectionAttributes`

NewUNIXConnectionAttributesWithDefaults instantiates a new UNIXConnectionAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGROUPS_FILE

`func (o *UNIXConnectionAttributes) GetGROUPS_FILE() string`

GetGROUPS_FILE returns the GROUPS_FILE field if non-nil, zero value otherwise.

### GetGROUPS_FILEOk

`func (o *UNIXConnectionAttributes) GetGROUPS_FILEOk() (*string, bool)`

GetGROUPS_FILEOk returns a tuple with the GROUPS_FILE field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGROUPS_FILE

`func (o *UNIXConnectionAttributes) SetGROUPS_FILE(v string)`

SetGROUPS_FILE sets GROUPS_FILE field to given value.

### HasGROUPS_FILE

`func (o *UNIXConnectionAttributes) HasGROUPS_FILE() bool`

HasGROUPS_FILE returns a boolean if a field has been set.

### GetSSH_KEY

`func (o *UNIXConnectionAttributes) GetSSH_KEY() string`

GetSSH_KEY returns the SSH_KEY field if non-nil, zero value otherwise.

### GetSSH_KEYOk

`func (o *UNIXConnectionAttributes) GetSSH_KEYOk() (*string, bool)`

GetSSH_KEYOk returns a tuple with the SSH_KEY field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSSH_KEY

`func (o *UNIXConnectionAttributes) SetSSH_KEY(v string)`

SetSSH_KEY sets SSH_KEY field to given value.

### HasSSH_KEY

`func (o *UNIXConnectionAttributes) HasSSH_KEY() bool`

HasSSH_KEY returns a boolean if a field has been set.

### GetACCOUNT_ENTITLEMENT_MAPPING_COMMAND

`func (o *UNIXConnectionAttributes) GetACCOUNT_ENTITLEMENT_MAPPING_COMMAND() string`

GetACCOUNT_ENTITLEMENT_MAPPING_COMMAND returns the ACCOUNT_ENTITLEMENT_MAPPING_COMMAND field if non-nil, zero value otherwise.

### GetACCOUNT_ENTITLEMENT_MAPPING_COMMANDOk

`func (o *UNIXConnectionAttributes) GetACCOUNT_ENTITLEMENT_MAPPING_COMMANDOk() (*string, bool)`

GetACCOUNT_ENTITLEMENT_MAPPING_COMMANDOk returns a tuple with the ACCOUNT_ENTITLEMENT_MAPPING_COMMAND field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetACCOUNT_ENTITLEMENT_MAPPING_COMMAND

`func (o *UNIXConnectionAttributes) SetACCOUNT_ENTITLEMENT_MAPPING_COMMAND(v string)`

SetACCOUNT_ENTITLEMENT_MAPPING_COMMAND sets ACCOUNT_ENTITLEMENT_MAPPING_COMMAND field to given value.

### HasACCOUNT_ENTITLEMENT_MAPPING_COMMAND

`func (o *UNIXConnectionAttributes) HasACCOUNT_ENTITLEMENT_MAPPING_COMMAND() bool`

HasACCOUNT_ENTITLEMENT_MAPPING_COMMAND returns a boolean if a field has been set.

### GetREMOVE_ACCESS_COMMAND

`func (o *UNIXConnectionAttributes) GetREMOVE_ACCESS_COMMAND() string`

GetREMOVE_ACCESS_COMMAND returns the REMOVE_ACCESS_COMMAND field if non-nil, zero value otherwise.

### GetREMOVE_ACCESS_COMMANDOk

`func (o *UNIXConnectionAttributes) GetREMOVE_ACCESS_COMMANDOk() (*string, bool)`

GetREMOVE_ACCESS_COMMANDOk returns a tuple with the REMOVE_ACCESS_COMMAND field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetREMOVE_ACCESS_COMMAND

`func (o *UNIXConnectionAttributes) SetREMOVE_ACCESS_COMMAND(v string)`

SetREMOVE_ACCESS_COMMAND sets REMOVE_ACCESS_COMMAND field to given value.

### HasREMOVE_ACCESS_COMMAND

`func (o *UNIXConnectionAttributes) HasREMOVE_ACCESS_COMMAND() bool`

HasREMOVE_ACCESS_COMMAND returns a boolean if a field has been set.

### GetPEM_KEY_FILE

`func (o *UNIXConnectionAttributes) GetPEM_KEY_FILE() string`

GetPEM_KEY_FILE returns the PEM_KEY_FILE field if non-nil, zero value otherwise.

### GetPEM_KEY_FILEOk

`func (o *UNIXConnectionAttributes) GetPEM_KEY_FILEOk() (*string, bool)`

GetPEM_KEY_FILEOk returns a tuple with the PEM_KEY_FILE field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPEM_KEY_FILE

`func (o *UNIXConnectionAttributes) SetPEM_KEY_FILE(v string)`

SetPEM_KEY_FILE sets PEM_KEY_FILE field to given value.

### HasPEM_KEY_FILE

`func (o *UNIXConnectionAttributes) HasPEM_KEY_FILE() bool`

HasPEM_KEY_FILE returns a boolean if a field has been set.

### GetPassThroughConnectionDetails

`func (o *UNIXConnectionAttributes) GetPassThroughConnectionDetails() string`

GetPassThroughConnectionDetails returns the PassThroughConnectionDetails field if non-nil, zero value otherwise.

### GetPassThroughConnectionDetailsOk

`func (o *UNIXConnectionAttributes) GetPassThroughConnectionDetailsOk() (*string, bool)`

GetPassThroughConnectionDetailsOk returns a tuple with the PassThroughConnectionDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassThroughConnectionDetails

`func (o *UNIXConnectionAttributes) SetPassThroughConnectionDetails(v string)`

SetPassThroughConnectionDetails sets PassThroughConnectionDetails field to given value.

### HasPassThroughConnectionDetails

`func (o *UNIXConnectionAttributes) HasPassThroughConnectionDetails() bool`

HasPassThroughConnectionDetails returns a boolean if a field has been set.

### GetDISABLE_ACCOUNT_COMMAND

`func (o *UNIXConnectionAttributes) GetDISABLE_ACCOUNT_COMMAND() string`

GetDISABLE_ACCOUNT_COMMAND returns the DISABLE_ACCOUNT_COMMAND field if non-nil, zero value otherwise.

### GetDISABLE_ACCOUNT_COMMANDOk

`func (o *UNIXConnectionAttributes) GetDISABLE_ACCOUNT_COMMANDOk() (*string, bool)`

GetDISABLE_ACCOUNT_COMMANDOk returns a tuple with the DISABLE_ACCOUNT_COMMAND field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDISABLE_ACCOUNT_COMMAND

`func (o *UNIXConnectionAttributes) SetDISABLE_ACCOUNT_COMMAND(v string)`

SetDISABLE_ACCOUNT_COMMAND sets DISABLE_ACCOUNT_COMMAND field to given value.

### HasDISABLE_ACCOUNT_COMMAND

`func (o *UNIXConnectionAttributes) HasDISABLE_ACCOUNT_COMMAND() bool`

HasDISABLE_ACCOUNT_COMMAND returns a boolean if a field has been set.

### GetConnectionTimeoutConfig

`func (o *UNIXConnectionAttributes) GetConnectionTimeoutConfig() ConnectionTimeoutConfig`

GetConnectionTimeoutConfig returns the ConnectionTimeoutConfig field if non-nil, zero value otherwise.

### GetConnectionTimeoutConfigOk

`func (o *UNIXConnectionAttributes) GetConnectionTimeoutConfigOk() (*ConnectionTimeoutConfig, bool)`

GetConnectionTimeoutConfigOk returns a tuple with the ConnectionTimeoutConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionTimeoutConfig

`func (o *UNIXConnectionAttributes) SetConnectionTimeoutConfig(v ConnectionTimeoutConfig)`

SetConnectionTimeoutConfig sets ConnectionTimeoutConfig field to given value.

### HasConnectionTimeoutConfig

`func (o *UNIXConnectionAttributes) HasConnectionTimeoutConfig() bool`

HasConnectionTimeoutConfig returns a boolean if a field has been set.

### GetPORT_NUMBER

`func (o *UNIXConnectionAttributes) GetPORT_NUMBER() string`

GetPORT_NUMBER returns the PORT_NUMBER field if non-nil, zero value otherwise.

### GetPORT_NUMBEROk

`func (o *UNIXConnectionAttributes) GetPORT_NUMBEROk() (*string, bool)`

GetPORT_NUMBEROk returns a tuple with the PORT_NUMBER field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPORT_NUMBER

`func (o *UNIXConnectionAttributes) SetPORT_NUMBER(v string)`

SetPORT_NUMBER sets PORT_NUMBER field to given value.

### HasPORT_NUMBER

`func (o *UNIXConnectionAttributes) HasPORT_NUMBER() bool`

HasPORT_NUMBER returns a boolean if a field has been set.

### GetConnectionType

`func (o *UNIXConnectionAttributes) GetConnectionType() string`

GetConnectionType returns the ConnectionType field if non-nil, zero value otherwise.

### GetConnectionTypeOk

`func (o *UNIXConnectionAttributes) GetConnectionTypeOk() (*string, bool)`

GetConnectionTypeOk returns a tuple with the ConnectionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionType

`func (o *UNIXConnectionAttributes) SetConnectionType(v string)`

SetConnectionType sets ConnectionType field to given value.

### HasConnectionType

`func (o *UNIXConnectionAttributes) HasConnectionType() bool`

HasConnectionType returns a boolean if a field has been set.

### GetCREATE_GROUP_COMMAND

`func (o *UNIXConnectionAttributes) GetCREATE_GROUP_COMMAND() string`

GetCREATE_GROUP_COMMAND returns the CREATE_GROUP_COMMAND field if non-nil, zero value otherwise.

### GetCREATE_GROUP_COMMANDOk

`func (o *UNIXConnectionAttributes) GetCREATE_GROUP_COMMANDOk() (*string, bool)`

GetCREATE_GROUP_COMMANDOk returns a tuple with the CREATE_GROUP_COMMAND field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCREATE_GROUP_COMMAND

`func (o *UNIXConnectionAttributes) SetCREATE_GROUP_COMMAND(v string)`

SetCREATE_GROUP_COMMAND sets CREATE_GROUP_COMMAND field to given value.

### HasCREATE_GROUP_COMMAND

`func (o *UNIXConnectionAttributes) HasCREATE_GROUP_COMMAND() bool`

HasCREATE_GROUP_COMMAND returns a boolean if a field has been set.

### GetACCOUNTS_FILE

`func (o *UNIXConnectionAttributes) GetACCOUNTS_FILE() string`

GetACCOUNTS_FILE returns the ACCOUNTS_FILE field if non-nil, zero value otherwise.

### GetACCOUNTS_FILEOk

`func (o *UNIXConnectionAttributes) GetACCOUNTS_FILEOk() (*string, bool)`

GetACCOUNTS_FILEOk returns a tuple with the ACCOUNTS_FILE field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetACCOUNTS_FILE

`func (o *UNIXConnectionAttributes) SetACCOUNTS_FILE(v string)`

SetACCOUNTS_FILE sets ACCOUNTS_FILE field to given value.

### HasACCOUNTS_FILE

`func (o *UNIXConnectionAttributes) HasACCOUNTS_FILE() bool`

HasACCOUNTS_FILE returns a boolean if a field has been set.

### GetPASSPHRASE

`func (o *UNIXConnectionAttributes) GetPASSPHRASE() string`

GetPASSPHRASE returns the PASSPHRASE field if non-nil, zero value otherwise.

### GetPASSPHRASEOk

`func (o *UNIXConnectionAttributes) GetPASSPHRASEOk() (*string, bool)`

GetPASSPHRASEOk returns a tuple with the PASSPHRASE field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPASSPHRASE

`func (o *UNIXConnectionAttributes) SetPASSPHRASE(v string)`

SetPASSPHRASE sets PASSPHRASE field to given value.

### HasPASSPHRASE

`func (o *UNIXConnectionAttributes) HasPASSPHRASE() bool`

HasPASSPHRASE returns a boolean if a field has been set.

### GetDELETE_GROUP_COMMAND

`func (o *UNIXConnectionAttributes) GetDELETE_GROUP_COMMAND() string`

GetDELETE_GROUP_COMMAND returns the DELETE_GROUP_COMMAND field if non-nil, zero value otherwise.

### GetDELETE_GROUP_COMMANDOk

`func (o *UNIXConnectionAttributes) GetDELETE_GROUP_COMMANDOk() (*string, bool)`

GetDELETE_GROUP_COMMANDOk returns a tuple with the DELETE_GROUP_COMMAND field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDELETE_GROUP_COMMAND

`func (o *UNIXConnectionAttributes) SetDELETE_GROUP_COMMAND(v string)`

SetDELETE_GROUP_COMMAND sets DELETE_GROUP_COMMAND field to given value.

### HasDELETE_GROUP_COMMAND

`func (o *UNIXConnectionAttributes) HasDELETE_GROUP_COMMAND() bool`

HasDELETE_GROUP_COMMAND returns a boolean if a field has been set.

### GetHOST_NAME

`func (o *UNIXConnectionAttributes) GetHOST_NAME() string`

GetHOST_NAME returns the HOST_NAME field if non-nil, zero value otherwise.

### GetHOST_NAMEOk

`func (o *UNIXConnectionAttributes) GetHOST_NAMEOk() (*string, bool)`

GetHOST_NAMEOk returns a tuple with the HOST_NAME field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHOST_NAME

`func (o *UNIXConnectionAttributes) SetHOST_NAME(v string)`

SetHOST_NAME sets HOST_NAME field to given value.

### HasHOST_NAME

`func (o *UNIXConnectionAttributes) HasHOST_NAME() bool`

HasHOST_NAME returns a boolean if a field has been set.

### GetADD_GROUP_OWNER_COMMAND

`func (o *UNIXConnectionAttributes) GetADD_GROUP_OWNER_COMMAND() string`

GetADD_GROUP_OWNER_COMMAND returns the ADD_GROUP_OWNER_COMMAND field if non-nil, zero value otherwise.

### GetADD_GROUP_OWNER_COMMANDOk

`func (o *UNIXConnectionAttributes) GetADD_GROUP_OWNER_COMMANDOk() (*string, bool)`

GetADD_GROUP_OWNER_COMMANDOk returns a tuple with the ADD_GROUP_OWNER_COMMAND field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetADD_GROUP_OWNER_COMMAND

`func (o *UNIXConnectionAttributes) SetADD_GROUP_OWNER_COMMAND(v string)`

SetADD_GROUP_OWNER_COMMAND sets ADD_GROUP_OWNER_COMMAND field to given value.

### HasADD_GROUP_OWNER_COMMAND

`func (o *UNIXConnectionAttributes) HasADD_GROUP_OWNER_COMMAND() bool`

HasADD_GROUP_OWNER_COMMAND returns a boolean if a field has been set.

### GetSTATUS_THRESHOLD_CONFIG

`func (o *UNIXConnectionAttributes) GetSTATUS_THRESHOLD_CONFIG() string`

GetSTATUS_THRESHOLD_CONFIG returns the STATUS_THRESHOLD_CONFIG field if non-nil, zero value otherwise.

### GetSTATUS_THRESHOLD_CONFIGOk

`func (o *UNIXConnectionAttributes) GetSTATUS_THRESHOLD_CONFIGOk() (*string, bool)`

GetSTATUS_THRESHOLD_CONFIGOk returns a tuple with the STATUS_THRESHOLD_CONFIG field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSTATUS_THRESHOLD_CONFIG

`func (o *UNIXConnectionAttributes) SetSTATUS_THRESHOLD_CONFIG(v string)`

SetSTATUS_THRESHOLD_CONFIG sets STATUS_THRESHOLD_CONFIG field to given value.

### HasSTATUS_THRESHOLD_CONFIG

`func (o *UNIXConnectionAttributes) HasSTATUS_THRESHOLD_CONFIG() bool`

HasSTATUS_THRESHOLD_CONFIG returns a boolean if a field has been set.

### GetUSERNAME

`func (o *UNIXConnectionAttributes) GetUSERNAME() string`

GetUSERNAME returns the USERNAME field if non-nil, zero value otherwise.

### GetUSERNAMEOk

`func (o *UNIXConnectionAttributes) GetUSERNAMEOk() (*string, bool)`

GetUSERNAMEOk returns a tuple with the USERNAME field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUSERNAME

`func (o *UNIXConnectionAttributes) SetUSERNAME(v string)`

SetUSERNAME sets USERNAME field to given value.

### HasUSERNAME

`func (o *UNIXConnectionAttributes) HasUSERNAME() bool`

HasUSERNAME returns a boolean if a field has been set.

### GetINACTIVE_LOCK_ACCOUNT

`func (o *UNIXConnectionAttributes) GetINACTIVE_LOCK_ACCOUNT() string`

GetINACTIVE_LOCK_ACCOUNT returns the INACTIVE_LOCK_ACCOUNT field if non-nil, zero value otherwise.

### GetINACTIVE_LOCK_ACCOUNTOk

`func (o *UNIXConnectionAttributes) GetINACTIVE_LOCK_ACCOUNTOk() (*string, bool)`

GetINACTIVE_LOCK_ACCOUNTOk returns a tuple with the INACTIVE_LOCK_ACCOUNT field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetINACTIVE_LOCK_ACCOUNT

`func (o *UNIXConnectionAttributes) SetINACTIVE_LOCK_ACCOUNT(v string)`

SetINACTIVE_LOCK_ACCOUNT sets INACTIVE_LOCK_ACCOUNT field to given value.

### HasINACTIVE_LOCK_ACCOUNT

`func (o *UNIXConnectionAttributes) HasINACTIVE_LOCK_ACCOUNT() bool`

HasINACTIVE_LOCK_ACCOUNT returns a boolean if a field has been set.

### GetADD_ACCESS_COMMAND

`func (o *UNIXConnectionAttributes) GetADD_ACCESS_COMMAND() string`

GetADD_ACCESS_COMMAND returns the ADD_ACCESS_COMMAND field if non-nil, zero value otherwise.

### GetADD_ACCESS_COMMANDOk

`func (o *UNIXConnectionAttributes) GetADD_ACCESS_COMMANDOk() (*string, bool)`

GetADD_ACCESS_COMMANDOk returns a tuple with the ADD_ACCESS_COMMAND field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetADD_ACCESS_COMMAND

`func (o *UNIXConnectionAttributes) SetADD_ACCESS_COMMAND(v string)`

SetADD_ACCESS_COMMAND sets ADD_ACCESS_COMMAND field to given value.

### HasADD_ACCESS_COMMAND

`func (o *UNIXConnectionAttributes) HasADD_ACCESS_COMMAND() bool`

HasADD_ACCESS_COMMAND returns a boolean if a field has been set.

### GetUPDATE_ACCOUNT_COMMAND

`func (o *UNIXConnectionAttributes) GetUPDATE_ACCOUNT_COMMAND() string`

GetUPDATE_ACCOUNT_COMMAND returns the UPDATE_ACCOUNT_COMMAND field if non-nil, zero value otherwise.

### GetUPDATE_ACCOUNT_COMMANDOk

`func (o *UNIXConnectionAttributes) GetUPDATE_ACCOUNT_COMMANDOk() (*string, bool)`

GetUPDATE_ACCOUNT_COMMANDOk returns a tuple with the UPDATE_ACCOUNT_COMMAND field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUPDATE_ACCOUNT_COMMAND

`func (o *UNIXConnectionAttributes) SetUPDATE_ACCOUNT_COMMAND(v string)`

SetUPDATE_ACCOUNT_COMMAND sets UPDATE_ACCOUNT_COMMAND field to given value.

### HasUPDATE_ACCOUNT_COMMAND

`func (o *UNIXConnectionAttributes) HasUPDATE_ACCOUNT_COMMAND() bool`

HasUPDATE_ACCOUNT_COMMAND returns a boolean if a field has been set.

### GetSSHPassThroughPassphrase

`func (o *UNIXConnectionAttributes) GetSSHPassThroughPassphrase() string`

GetSSHPassThroughPassphrase returns the SSHPassThroughPassphrase field if non-nil, zero value otherwise.

### GetSSHPassThroughPassphraseOk

`func (o *UNIXConnectionAttributes) GetSSHPassThroughPassphraseOk() (*string, bool)`

GetSSHPassThroughPassphraseOk returns a tuple with the SSHPassThroughPassphrase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSSHPassThroughPassphrase

`func (o *UNIXConnectionAttributes) SetSSHPassThroughPassphrase(v string)`

SetSSHPassThroughPassphrase sets SSHPassThroughPassphrase field to given value.

### HasSSHPassThroughPassphrase

`func (o *UNIXConnectionAttributes) HasSSHPassThroughPassphrase() bool`

HasSSHPassThroughPassphrase returns a boolean if a field has been set.

### GetSHADOW_FILE

`func (o *UNIXConnectionAttributes) GetSHADOW_FILE() string`

GetSHADOW_FILE returns the SHADOW_FILE field if non-nil, zero value otherwise.

### GetSHADOW_FILEOk

`func (o *UNIXConnectionAttributes) GetSHADOW_FILEOk() (*string, bool)`

GetSHADOW_FILEOk returns a tuple with the SHADOW_FILE field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSHADOW_FILE

`func (o *UNIXConnectionAttributes) SetSHADOW_FILE(v string)`

SetSHADOW_FILE sets SHADOW_FILE field to given value.

### HasSHADOW_FILE

`func (o *UNIXConnectionAttributes) HasSHADOW_FILE() bool`

HasSHADOW_FILE returns a boolean if a field has been set.

### GetIsTimeoutSupported

`func (o *UNIXConnectionAttributes) GetIsTimeoutSupported() bool`

GetIsTimeoutSupported returns the IsTimeoutSupported field if non-nil, zero value otherwise.

### GetIsTimeoutSupportedOk

`func (o *UNIXConnectionAttributes) GetIsTimeoutSupportedOk() (*bool, bool)`

GetIsTimeoutSupportedOk returns a tuple with the IsTimeoutSupported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsTimeoutSupported

`func (o *UNIXConnectionAttributes) SetIsTimeoutSupported(v bool)`

SetIsTimeoutSupported sets IsTimeoutSupported field to given value.

### HasIsTimeoutSupported

`func (o *UNIXConnectionAttributes) HasIsTimeoutSupported() bool`

HasIsTimeoutSupported returns a boolean if a field has been set.

### GetSSHPassThroughSSHKEY

`func (o *UNIXConnectionAttributes) GetSSHPassThroughSSHKEY() string`

GetSSHPassThroughSSHKEY returns the SSHPassThroughSSHKEY field if non-nil, zero value otherwise.

### GetSSHPassThroughSSHKEYOk

`func (o *UNIXConnectionAttributes) GetSSHPassThroughSSHKEYOk() (*string, bool)`

GetSSHPassThroughSSHKEYOk returns a tuple with the SSHPassThroughSSHKEY field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSSHPassThroughSSHKEY

`func (o *UNIXConnectionAttributes) SetSSHPassThroughSSHKEY(v string)`

SetSSHPassThroughSSHKEY sets SSHPassThroughSSHKEY field to given value.

### HasSSHPassThroughSSHKEY

`func (o *UNIXConnectionAttributes) HasSSHPassThroughSSHKEY() bool`

HasSSHPassThroughSSHKEY returns a boolean if a field has been set.

### GetPROVISION_ACCOUNT_COMMAND

`func (o *UNIXConnectionAttributes) GetPROVISION_ACCOUNT_COMMAND() string`

GetPROVISION_ACCOUNT_COMMAND returns the PROVISION_ACCOUNT_COMMAND field if non-nil, zero value otherwise.

### GetPROVISION_ACCOUNT_COMMANDOk

`func (o *UNIXConnectionAttributes) GetPROVISION_ACCOUNT_COMMANDOk() (*string, bool)`

GetPROVISION_ACCOUNT_COMMANDOk returns a tuple with the PROVISION_ACCOUNT_COMMAND field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPROVISION_ACCOUNT_COMMAND

`func (o *UNIXConnectionAttributes) SetPROVISION_ACCOUNT_COMMAND(v string)`

SetPROVISION_ACCOUNT_COMMAND sets PROVISION_ACCOUNT_COMMAND field to given value.

### HasPROVISION_ACCOUNT_COMMAND

`func (o *UNIXConnectionAttributes) HasPROVISION_ACCOUNT_COMMAND() bool`

HasPROVISION_ACCOUNT_COMMAND returns a boolean if a field has been set.

### GetFIREFIGHTERID_GRANT_ACCESS_COMMAND

`func (o *UNIXConnectionAttributes) GetFIREFIGHTERID_GRANT_ACCESS_COMMAND() string`

GetFIREFIGHTERID_GRANT_ACCESS_COMMAND returns the FIREFIGHTERID_GRANT_ACCESS_COMMAND field if non-nil, zero value otherwise.

### GetFIREFIGHTERID_GRANT_ACCESS_COMMANDOk

`func (o *UNIXConnectionAttributes) GetFIREFIGHTERID_GRANT_ACCESS_COMMANDOk() (*string, bool)`

GetFIREFIGHTERID_GRANT_ACCESS_COMMANDOk returns a tuple with the FIREFIGHTERID_GRANT_ACCESS_COMMAND field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFIREFIGHTERID_GRANT_ACCESS_COMMAND

`func (o *UNIXConnectionAttributes) SetFIREFIGHTERID_GRANT_ACCESS_COMMAND(v string)`

SetFIREFIGHTERID_GRANT_ACCESS_COMMAND sets FIREFIGHTERID_GRANT_ACCESS_COMMAND field to given value.

### HasFIREFIGHTERID_GRANT_ACCESS_COMMAND

`func (o *UNIXConnectionAttributes) HasFIREFIGHTERID_GRANT_ACCESS_COMMAND() bool`

HasFIREFIGHTERID_GRANT_ACCESS_COMMAND returns a boolean if a field has been set.

### GetUNLOCK_ACCOUNT_COMMAND

`func (o *UNIXConnectionAttributes) GetUNLOCK_ACCOUNT_COMMAND() string`

GetUNLOCK_ACCOUNT_COMMAND returns the UNLOCK_ACCOUNT_COMMAND field if non-nil, zero value otherwise.

### GetUNLOCK_ACCOUNT_COMMANDOk

`func (o *UNIXConnectionAttributes) GetUNLOCK_ACCOUNT_COMMANDOk() (*string, bool)`

GetUNLOCK_ACCOUNT_COMMANDOk returns a tuple with the UNLOCK_ACCOUNT_COMMAND field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUNLOCK_ACCOUNT_COMMAND

`func (o *UNIXConnectionAttributes) SetUNLOCK_ACCOUNT_COMMAND(v string)`

SetUNLOCK_ACCOUNT_COMMAND sets UNLOCK_ACCOUNT_COMMAND field to given value.

### HasUNLOCK_ACCOUNT_COMMAND

`func (o *UNIXConnectionAttributes) HasUNLOCK_ACCOUNT_COMMAND() bool`

HasUNLOCK_ACCOUNT_COMMAND returns a boolean if a field has been set.

### GetDEPROVISION_ACCOUNT_COMMAND

`func (o *UNIXConnectionAttributes) GetDEPROVISION_ACCOUNT_COMMAND() string`

GetDEPROVISION_ACCOUNT_COMMAND returns the DEPROVISION_ACCOUNT_COMMAND field if non-nil, zero value otherwise.

### GetDEPROVISION_ACCOUNT_COMMANDOk

`func (o *UNIXConnectionAttributes) GetDEPROVISION_ACCOUNT_COMMANDOk() (*string, bool)`

GetDEPROVISION_ACCOUNT_COMMANDOk returns a tuple with the DEPROVISION_ACCOUNT_COMMAND field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDEPROVISION_ACCOUNT_COMMAND

`func (o *UNIXConnectionAttributes) SetDEPROVISION_ACCOUNT_COMMAND(v string)`

SetDEPROVISION_ACCOUNT_COMMAND sets DEPROVISION_ACCOUNT_COMMAND field to given value.

### HasDEPROVISION_ACCOUNT_COMMAND

`func (o *UNIXConnectionAttributes) HasDEPROVISION_ACCOUNT_COMMAND() bool`

HasDEPROVISION_ACCOUNT_COMMAND returns a boolean if a field has been set.

### GetCHANGE_PASSWRD_JSON

`func (o *UNIXConnectionAttributes) GetCHANGE_PASSWRD_JSON() string`

GetCHANGE_PASSWRD_JSON returns the CHANGE_PASSWRD_JSON field if non-nil, zero value otherwise.

### GetCHANGE_PASSWRD_JSONOk

`func (o *UNIXConnectionAttributes) GetCHANGE_PASSWRD_JSONOk() (*string, bool)`

GetCHANGE_PASSWRD_JSONOk returns a tuple with the CHANGE_PASSWRD_JSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCHANGE_PASSWRD_JSON

`func (o *UNIXConnectionAttributes) SetCHANGE_PASSWRD_JSON(v string)`

SetCHANGE_PASSWRD_JSON sets CHANGE_PASSWRD_JSON field to given value.

### HasCHANGE_PASSWRD_JSON

`func (o *UNIXConnectionAttributes) HasCHANGE_PASSWRD_JSON() bool`

HasCHANGE_PASSWRD_JSON returns a boolean if a field has been set.

### GetSSHPassThroughPassword

`func (o *UNIXConnectionAttributes) GetSSHPassThroughPassword() string`

GetSSHPassThroughPassword returns the SSHPassThroughPassword field if non-nil, zero value otherwise.

### GetSSHPassThroughPasswordOk

`func (o *UNIXConnectionAttributes) GetSSHPassThroughPasswordOk() (*string, bool)`

GetSSHPassThroughPasswordOk returns a tuple with the SSHPassThroughPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSSHPassThroughPassword

`func (o *UNIXConnectionAttributes) SetSSHPassThroughPassword(v string)`

SetSSHPassThroughPassword sets SSHPassThroughPassword field to given value.

### HasSSHPassThroughPassword

`func (o *UNIXConnectionAttributes) HasSSHPassThroughPassword() bool`

HasSSHPassThroughPassword returns a boolean if a field has been set.

### GetFIREFIGHTERID_REVOKE_ACCESS_COMMAND

`func (o *UNIXConnectionAttributes) GetFIREFIGHTERID_REVOKE_ACCESS_COMMAND() string`

GetFIREFIGHTERID_REVOKE_ACCESS_COMMAND returns the FIREFIGHTERID_REVOKE_ACCESS_COMMAND field if non-nil, zero value otherwise.

### GetFIREFIGHTERID_REVOKE_ACCESS_COMMANDOk

`func (o *UNIXConnectionAttributes) GetFIREFIGHTERID_REVOKE_ACCESS_COMMANDOk() (*string, bool)`

GetFIREFIGHTERID_REVOKE_ACCESS_COMMANDOk returns a tuple with the FIREFIGHTERID_REVOKE_ACCESS_COMMAND field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFIREFIGHTERID_REVOKE_ACCESS_COMMAND

`func (o *UNIXConnectionAttributes) SetFIREFIGHTERID_REVOKE_ACCESS_COMMAND(v string)`

SetFIREFIGHTERID_REVOKE_ACCESS_COMMAND sets FIREFIGHTERID_REVOKE_ACCESS_COMMAND field to given value.

### HasFIREFIGHTERID_REVOKE_ACCESS_COMMAND

`func (o *UNIXConnectionAttributes) HasFIREFIGHTERID_REVOKE_ACCESS_COMMAND() bool`

HasFIREFIGHTERID_REVOKE_ACCESS_COMMAND returns a boolean if a field has been set.

### GetADD_PRIMARY_GROUP_COMMAND

`func (o *UNIXConnectionAttributes) GetADD_PRIMARY_GROUP_COMMAND() string`

GetADD_PRIMARY_GROUP_COMMAND returns the ADD_PRIMARY_GROUP_COMMAND field if non-nil, zero value otherwise.

### GetADD_PRIMARY_GROUP_COMMANDOk

`func (o *UNIXConnectionAttributes) GetADD_PRIMARY_GROUP_COMMANDOk() (*string, bool)`

GetADD_PRIMARY_GROUP_COMMANDOk returns a tuple with the ADD_PRIMARY_GROUP_COMMAND field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetADD_PRIMARY_GROUP_COMMAND

`func (o *UNIXConnectionAttributes) SetADD_PRIMARY_GROUP_COMMAND(v string)`

SetADD_PRIMARY_GROUP_COMMAND sets ADD_PRIMARY_GROUP_COMMAND field to given value.

### HasADD_PRIMARY_GROUP_COMMAND

`func (o *UNIXConnectionAttributes) HasADD_PRIMARY_GROUP_COMMAND() bool`

HasADD_PRIMARY_GROUP_COMMAND returns a boolean if a field has been set.

### GetIsTimeoutConfigValidated

`func (o *UNIXConnectionAttributes) GetIsTimeoutConfigValidated() bool`

GetIsTimeoutConfigValidated returns the IsTimeoutConfigValidated field if non-nil, zero value otherwise.

### GetIsTimeoutConfigValidatedOk

`func (o *UNIXConnectionAttributes) GetIsTimeoutConfigValidatedOk() (*bool, bool)`

GetIsTimeoutConfigValidatedOk returns a tuple with the IsTimeoutConfigValidated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsTimeoutConfigValidated

`func (o *UNIXConnectionAttributes) SetIsTimeoutConfigValidated(v bool)`

SetIsTimeoutConfigValidated sets IsTimeoutConfigValidated field to given value.

### HasIsTimeoutConfigValidated

`func (o *UNIXConnectionAttributes) HasIsTimeoutConfigValidated() bool`

HasIsTimeoutConfigValidated returns a boolean if a field has been set.

### GetLOCK_ACCOUNT_COMMAND

`func (o *UNIXConnectionAttributes) GetLOCK_ACCOUNT_COMMAND() string`

GetLOCK_ACCOUNT_COMMAND returns the LOCK_ACCOUNT_COMMAND field if non-nil, zero value otherwise.

### GetLOCK_ACCOUNT_COMMANDOk

`func (o *UNIXConnectionAttributes) GetLOCK_ACCOUNT_COMMANDOk() (*string, bool)`

GetLOCK_ACCOUNT_COMMANDOk returns a tuple with the LOCK_ACCOUNT_COMMAND field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLOCK_ACCOUNT_COMMAND

`func (o *UNIXConnectionAttributes) SetLOCK_ACCOUNT_COMMAND(v string)`

SetLOCK_ACCOUNT_COMMAND sets LOCK_ACCOUNT_COMMAND field to given value.

### HasLOCK_ACCOUNT_COMMAND

`func (o *UNIXConnectionAttributes) HasLOCK_ACCOUNT_COMMAND() bool`

HasLOCK_ACCOUNT_COMMAND returns a boolean if a field has been set.

### GetPASSWORD

`func (o *UNIXConnectionAttributes) GetPASSWORD() string`

GetPASSWORD returns the PASSWORD field if non-nil, zero value otherwise.

### GetPASSWORDOk

`func (o *UNIXConnectionAttributes) GetPASSWORDOk() (*string, bool)`

GetPASSWORDOk returns a tuple with the PASSWORD field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPASSWORD

`func (o *UNIXConnectionAttributes) SetPASSWORD(v string)`

SetPASSWORD sets PASSWORD field to given value.

### HasPASSWORD

`func (o *UNIXConnectionAttributes) HasPASSWORD() bool`

HasPASSWORD returns a boolean if a field has been set.

### GetCUSTOM_CONFIG_JSON

`func (o *UNIXConnectionAttributes) GetCUSTOM_CONFIG_JSON() string`

GetCUSTOM_CONFIG_JSON returns the CUSTOM_CONFIG_JSON field if non-nil, zero value otherwise.

### GetCUSTOM_CONFIG_JSONOk

`func (o *UNIXConnectionAttributes) GetCUSTOM_CONFIG_JSONOk() (*string, bool)`

GetCUSTOM_CONFIG_JSONOk returns a tuple with the CUSTOM_CONFIG_JSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCUSTOM_CONFIG_JSON

`func (o *UNIXConnectionAttributes) SetCUSTOM_CONFIG_JSON(v string)`

SetCUSTOM_CONFIG_JSON sets CUSTOM_CONFIG_JSON field to given value.

### HasCUSTOM_CONFIG_JSON

`func (o *UNIXConnectionAttributes) HasCUSTOM_CONFIG_JSON() bool`

HasCUSTOM_CONFIG_JSON returns a boolean if a field has been set.

### GetENABLE_ACCOUNT_COMMAND

`func (o *UNIXConnectionAttributes) GetENABLE_ACCOUNT_COMMAND() string`

GetENABLE_ACCOUNT_COMMAND returns the ENABLE_ACCOUNT_COMMAND field if non-nil, zero value otherwise.

### GetENABLE_ACCOUNT_COMMANDOk

`func (o *UNIXConnectionAttributes) GetENABLE_ACCOUNT_COMMANDOk() (*string, bool)`

GetENABLE_ACCOUNT_COMMANDOk returns a tuple with the ENABLE_ACCOUNT_COMMAND field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetENABLE_ACCOUNT_COMMAND

`func (o *UNIXConnectionAttributes) SetENABLE_ACCOUNT_COMMAND(v string)`

SetENABLE_ACCOUNT_COMMAND sets ENABLE_ACCOUNT_COMMAND field to given value.

### HasENABLE_ACCOUNT_COMMAND

`func (o *UNIXConnectionAttributes) HasENABLE_ACCOUNT_COMMAND() bool`

HasENABLE_ACCOUNT_COMMAND returns a boolean if a field has been set.

### GetSERVER_TYPE

`func (o *UNIXConnectionAttributes) GetSERVER_TYPE() string`

GetSERVER_TYPE returns the SERVER_TYPE field if non-nil, zero value otherwise.

### GetSERVER_TYPEOk

`func (o *UNIXConnectionAttributes) GetSERVER_TYPEOk() (*string, bool)`

GetSERVER_TYPEOk returns a tuple with the SERVER_TYPE field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSERVER_TYPE

`func (o *UNIXConnectionAttributes) SetSERVER_TYPE(v string)`

SetSERVER_TYPE sets SERVER_TYPE field to given value.

### HasSERVER_TYPE

`func (o *UNIXConnectionAttributes) HasSERVER_TYPE() bool`

HasSERVER_TYPE returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


