# UNIXConnector

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HOST_NAME** | **string** | Property for HOST_NAME | 
**PORT_NUMBER** | **string** | Property for PORT_NUMBER | 
**USERNAME** | **string** | Property for USERNAME | 
**PASSWORD** | Pointer to **string** | Property for PASSWORD | [optional] 
**GROUPS_FILE** | Pointer to **string** | Property for GROUPS_FILE | [optional] 
**ACCOUNTS_FILE** | Pointer to **string** | Property for ACCOUNTS_FILE | [optional] 
**SHADOW_FILE** | Pointer to **string** | Property for SHADOW_FILE | [optional] 
**PROVISION_ACCOUNT_COMMAND** | Pointer to **string** | Property for PROVISION_ACCOUNT_COMMAND | [optional] 
**DEPROVISION_ACCOUNT_COMMAND** | Pointer to **string** | Property for DEPROVISION_ACCOUNT_COMMAND | [optional] 
**ADD_ACCESS_COMMAND** | Pointer to **string** | Property for ADD_ACCESS_COMMAND | [optional] 
**REMOVE_ACCESS_COMMAND** | Pointer to **string** | Property for REMOVE_ACCESS_COMMAND | [optional] 
**CHANGE_PASSWRD_JSON** | Pointer to **string** | Property for CHANGE_PASSWRD_JSON | [optional] 
**PEM_KEY_FILE** | Pointer to **string** | Property for PEM_KEY_FILE | [optional] 
**ENABLE_ACCOUNT_COMMAND** | Pointer to **string** | Property for ENABLE_ACCOUNT_COMMAND | [optional] 
**DISABLE_ACCOUNT_COMMAND** | Pointer to **string** | Property for DISABLE_ACCOUNT_COMMAND | [optional] 
**ACCOUNT_ENTITLEMENT_MAPPING_COMMAND** | Pointer to **string** | Property for ACCOUNT_ENTITLEMENT_MAPPING_COMMAND | [optional] 
**PASSPHRASE** | Pointer to **string** | Property for PASSPHRASE | [optional] 
**UPDATE_ACCOUNT_COMMAND** | Pointer to **string** | Property for UPDATE_ACCOUNT_COMMAND | [optional] 
**CREATE_GROUP_COMMAND** | Pointer to **string** | Property for CREATE_GROUP_COMMAND | [optional] 
**DELETE_GROUP_COMMAND** | Pointer to **string** | Property for DELETE_GROUP_COMMAND | [optional] 
**ADD_GROUP_OWNER_COMMAND** | Pointer to **string** | Property for ADD_GROUP_OWNER_COMMAND | [optional] 
**ADD_PRIMARY_GROUP_COMMAND** | Pointer to **string** | Property for ADD_PRIMARY_GROUP_COMMAND | [optional] 
**FIREFIGHTERID_GRANT_ACCESS_COMMAND** | Pointer to **string** | Property for FIREFIGHTERID_GRANT_ACCESS_COMMAND | [optional] 
**FIREFIGHTERID_REVOKE_ACCESS_COMMAND** | Pointer to **string** | Property for FIREFIGHTERID_REVOKE_ACCESS_COMMAND | [optional] 
**INACTIVE_LOCK_ACCOUNT** | Pointer to **string** | Property for INACTIVE_LOCK_ACCOUNT | [optional] 
**STATUS_THRESHOLD_CONFIG** | Pointer to **string** | Property for STATUS_THRESHOLD_CONFIG | [optional] 
**CUSTOM_CONFIG_JSON** | Pointer to **string** | Property for CUSTOM_CONFIG_JSON | [optional] 
**SSH_KEY** | Pointer to **string** | Property for SSH_KEY | [optional] 
**LOCK_ACCOUNT_COMMAND** | Pointer to **string** | Property for LOCK_ACCOUNT_COMMAND | [optional] 
**UNLOCK_ACCOUNT_COMMAND** | Pointer to **string** | Property for UNLOCK_ACCOUNT_COMMAND | [optional] 
**PassThroughConnectionDetails** | Pointer to **string** | Property for PassThroughConnectionDetails | [optional] 
**SSHPassThroughPassword** | Pointer to **string** | Property for SSHPassThroughPassword | [optional] 
**SSHPassThroughSSHKEY** | Pointer to **string** | Property for SSHPassThroughSSHKEY | [optional] 
**SSHPassThroughPassphrase** | Pointer to **string** | Property for SSHPassThroughPassphrase | [optional] 

## Methods

### NewUNIXConnector

`func NewUNIXConnector(hOSTNAME string, pORTNUMBER string, uSERNAME string, ) *UNIXConnector`

NewUNIXConnector instantiates a new UNIXConnector object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUNIXConnectorWithDefaults

`func NewUNIXConnectorWithDefaults() *UNIXConnector`

NewUNIXConnectorWithDefaults instantiates a new UNIXConnector object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHOST_NAME

`func (o *UNIXConnector) GetHOST_NAME() string`

GetHOST_NAME returns the HOST_NAME field if non-nil, zero value otherwise.

### GetHOST_NAMEOk

`func (o *UNIXConnector) GetHOST_NAMEOk() (*string, bool)`

GetHOST_NAMEOk returns a tuple with the HOST_NAME field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHOST_NAME

`func (o *UNIXConnector) SetHOST_NAME(v string)`

SetHOST_NAME sets HOST_NAME field to given value.


### GetPORT_NUMBER

`func (o *UNIXConnector) GetPORT_NUMBER() string`

GetPORT_NUMBER returns the PORT_NUMBER field if non-nil, zero value otherwise.

### GetPORT_NUMBEROk

`func (o *UNIXConnector) GetPORT_NUMBEROk() (*string, bool)`

GetPORT_NUMBEROk returns a tuple with the PORT_NUMBER field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPORT_NUMBER

`func (o *UNIXConnector) SetPORT_NUMBER(v string)`

SetPORT_NUMBER sets PORT_NUMBER field to given value.


### GetUSERNAME

`func (o *UNIXConnector) GetUSERNAME() string`

GetUSERNAME returns the USERNAME field if non-nil, zero value otherwise.

### GetUSERNAMEOk

`func (o *UNIXConnector) GetUSERNAMEOk() (*string, bool)`

GetUSERNAMEOk returns a tuple with the USERNAME field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUSERNAME

`func (o *UNIXConnector) SetUSERNAME(v string)`

SetUSERNAME sets USERNAME field to given value.


### GetPASSWORD

`func (o *UNIXConnector) GetPASSWORD() string`

GetPASSWORD returns the PASSWORD field if non-nil, zero value otherwise.

### GetPASSWORDOk

`func (o *UNIXConnector) GetPASSWORDOk() (*string, bool)`

GetPASSWORDOk returns a tuple with the PASSWORD field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPASSWORD

`func (o *UNIXConnector) SetPASSWORD(v string)`

SetPASSWORD sets PASSWORD field to given value.

### HasPASSWORD

`func (o *UNIXConnector) HasPASSWORD() bool`

HasPASSWORD returns a boolean if a field has been set.

### GetGROUPS_FILE

`func (o *UNIXConnector) GetGROUPS_FILE() string`

GetGROUPS_FILE returns the GROUPS_FILE field if non-nil, zero value otherwise.

### GetGROUPS_FILEOk

`func (o *UNIXConnector) GetGROUPS_FILEOk() (*string, bool)`

GetGROUPS_FILEOk returns a tuple with the GROUPS_FILE field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGROUPS_FILE

`func (o *UNIXConnector) SetGROUPS_FILE(v string)`

SetGROUPS_FILE sets GROUPS_FILE field to given value.

### HasGROUPS_FILE

`func (o *UNIXConnector) HasGROUPS_FILE() bool`

HasGROUPS_FILE returns a boolean if a field has been set.

### GetACCOUNTS_FILE

`func (o *UNIXConnector) GetACCOUNTS_FILE() string`

GetACCOUNTS_FILE returns the ACCOUNTS_FILE field if non-nil, zero value otherwise.

### GetACCOUNTS_FILEOk

`func (o *UNIXConnector) GetACCOUNTS_FILEOk() (*string, bool)`

GetACCOUNTS_FILEOk returns a tuple with the ACCOUNTS_FILE field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetACCOUNTS_FILE

`func (o *UNIXConnector) SetACCOUNTS_FILE(v string)`

SetACCOUNTS_FILE sets ACCOUNTS_FILE field to given value.

### HasACCOUNTS_FILE

`func (o *UNIXConnector) HasACCOUNTS_FILE() bool`

HasACCOUNTS_FILE returns a boolean if a field has been set.

### GetSHADOW_FILE

`func (o *UNIXConnector) GetSHADOW_FILE() string`

GetSHADOW_FILE returns the SHADOW_FILE field if non-nil, zero value otherwise.

### GetSHADOW_FILEOk

`func (o *UNIXConnector) GetSHADOW_FILEOk() (*string, bool)`

GetSHADOW_FILEOk returns a tuple with the SHADOW_FILE field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSHADOW_FILE

`func (o *UNIXConnector) SetSHADOW_FILE(v string)`

SetSHADOW_FILE sets SHADOW_FILE field to given value.

### HasSHADOW_FILE

`func (o *UNIXConnector) HasSHADOW_FILE() bool`

HasSHADOW_FILE returns a boolean if a field has been set.

### GetPROVISION_ACCOUNT_COMMAND

`func (o *UNIXConnector) GetPROVISION_ACCOUNT_COMMAND() string`

GetPROVISION_ACCOUNT_COMMAND returns the PROVISION_ACCOUNT_COMMAND field if non-nil, zero value otherwise.

### GetPROVISION_ACCOUNT_COMMANDOk

`func (o *UNIXConnector) GetPROVISION_ACCOUNT_COMMANDOk() (*string, bool)`

GetPROVISION_ACCOUNT_COMMANDOk returns a tuple with the PROVISION_ACCOUNT_COMMAND field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPROVISION_ACCOUNT_COMMAND

`func (o *UNIXConnector) SetPROVISION_ACCOUNT_COMMAND(v string)`

SetPROVISION_ACCOUNT_COMMAND sets PROVISION_ACCOUNT_COMMAND field to given value.

### HasPROVISION_ACCOUNT_COMMAND

`func (o *UNIXConnector) HasPROVISION_ACCOUNT_COMMAND() bool`

HasPROVISION_ACCOUNT_COMMAND returns a boolean if a field has been set.

### GetDEPROVISION_ACCOUNT_COMMAND

`func (o *UNIXConnector) GetDEPROVISION_ACCOUNT_COMMAND() string`

GetDEPROVISION_ACCOUNT_COMMAND returns the DEPROVISION_ACCOUNT_COMMAND field if non-nil, zero value otherwise.

### GetDEPROVISION_ACCOUNT_COMMANDOk

`func (o *UNIXConnector) GetDEPROVISION_ACCOUNT_COMMANDOk() (*string, bool)`

GetDEPROVISION_ACCOUNT_COMMANDOk returns a tuple with the DEPROVISION_ACCOUNT_COMMAND field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDEPROVISION_ACCOUNT_COMMAND

`func (o *UNIXConnector) SetDEPROVISION_ACCOUNT_COMMAND(v string)`

SetDEPROVISION_ACCOUNT_COMMAND sets DEPROVISION_ACCOUNT_COMMAND field to given value.

### HasDEPROVISION_ACCOUNT_COMMAND

`func (o *UNIXConnector) HasDEPROVISION_ACCOUNT_COMMAND() bool`

HasDEPROVISION_ACCOUNT_COMMAND returns a boolean if a field has been set.

### GetADD_ACCESS_COMMAND

`func (o *UNIXConnector) GetADD_ACCESS_COMMAND() string`

GetADD_ACCESS_COMMAND returns the ADD_ACCESS_COMMAND field if non-nil, zero value otherwise.

### GetADD_ACCESS_COMMANDOk

`func (o *UNIXConnector) GetADD_ACCESS_COMMANDOk() (*string, bool)`

GetADD_ACCESS_COMMANDOk returns a tuple with the ADD_ACCESS_COMMAND field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetADD_ACCESS_COMMAND

`func (o *UNIXConnector) SetADD_ACCESS_COMMAND(v string)`

SetADD_ACCESS_COMMAND sets ADD_ACCESS_COMMAND field to given value.

### HasADD_ACCESS_COMMAND

`func (o *UNIXConnector) HasADD_ACCESS_COMMAND() bool`

HasADD_ACCESS_COMMAND returns a boolean if a field has been set.

### GetREMOVE_ACCESS_COMMAND

`func (o *UNIXConnector) GetREMOVE_ACCESS_COMMAND() string`

GetREMOVE_ACCESS_COMMAND returns the REMOVE_ACCESS_COMMAND field if non-nil, zero value otherwise.

### GetREMOVE_ACCESS_COMMANDOk

`func (o *UNIXConnector) GetREMOVE_ACCESS_COMMANDOk() (*string, bool)`

GetREMOVE_ACCESS_COMMANDOk returns a tuple with the REMOVE_ACCESS_COMMAND field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetREMOVE_ACCESS_COMMAND

`func (o *UNIXConnector) SetREMOVE_ACCESS_COMMAND(v string)`

SetREMOVE_ACCESS_COMMAND sets REMOVE_ACCESS_COMMAND field to given value.

### HasREMOVE_ACCESS_COMMAND

`func (o *UNIXConnector) HasREMOVE_ACCESS_COMMAND() bool`

HasREMOVE_ACCESS_COMMAND returns a boolean if a field has been set.

### GetCHANGE_PASSWRD_JSON

`func (o *UNIXConnector) GetCHANGE_PASSWRD_JSON() string`

GetCHANGE_PASSWRD_JSON returns the CHANGE_PASSWRD_JSON field if non-nil, zero value otherwise.

### GetCHANGE_PASSWRD_JSONOk

`func (o *UNIXConnector) GetCHANGE_PASSWRD_JSONOk() (*string, bool)`

GetCHANGE_PASSWRD_JSONOk returns a tuple with the CHANGE_PASSWRD_JSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCHANGE_PASSWRD_JSON

`func (o *UNIXConnector) SetCHANGE_PASSWRD_JSON(v string)`

SetCHANGE_PASSWRD_JSON sets CHANGE_PASSWRD_JSON field to given value.

### HasCHANGE_PASSWRD_JSON

`func (o *UNIXConnector) HasCHANGE_PASSWRD_JSON() bool`

HasCHANGE_PASSWRD_JSON returns a boolean if a field has been set.

### GetPEM_KEY_FILE

`func (o *UNIXConnector) GetPEM_KEY_FILE() string`

GetPEM_KEY_FILE returns the PEM_KEY_FILE field if non-nil, zero value otherwise.

### GetPEM_KEY_FILEOk

`func (o *UNIXConnector) GetPEM_KEY_FILEOk() (*string, bool)`

GetPEM_KEY_FILEOk returns a tuple with the PEM_KEY_FILE field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPEM_KEY_FILE

`func (o *UNIXConnector) SetPEM_KEY_FILE(v string)`

SetPEM_KEY_FILE sets PEM_KEY_FILE field to given value.

### HasPEM_KEY_FILE

`func (o *UNIXConnector) HasPEM_KEY_FILE() bool`

HasPEM_KEY_FILE returns a boolean if a field has been set.

### GetENABLE_ACCOUNT_COMMAND

`func (o *UNIXConnector) GetENABLE_ACCOUNT_COMMAND() string`

GetENABLE_ACCOUNT_COMMAND returns the ENABLE_ACCOUNT_COMMAND field if non-nil, zero value otherwise.

### GetENABLE_ACCOUNT_COMMANDOk

`func (o *UNIXConnector) GetENABLE_ACCOUNT_COMMANDOk() (*string, bool)`

GetENABLE_ACCOUNT_COMMANDOk returns a tuple with the ENABLE_ACCOUNT_COMMAND field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetENABLE_ACCOUNT_COMMAND

`func (o *UNIXConnector) SetENABLE_ACCOUNT_COMMAND(v string)`

SetENABLE_ACCOUNT_COMMAND sets ENABLE_ACCOUNT_COMMAND field to given value.

### HasENABLE_ACCOUNT_COMMAND

`func (o *UNIXConnector) HasENABLE_ACCOUNT_COMMAND() bool`

HasENABLE_ACCOUNT_COMMAND returns a boolean if a field has been set.

### GetDISABLE_ACCOUNT_COMMAND

`func (o *UNIXConnector) GetDISABLE_ACCOUNT_COMMAND() string`

GetDISABLE_ACCOUNT_COMMAND returns the DISABLE_ACCOUNT_COMMAND field if non-nil, zero value otherwise.

### GetDISABLE_ACCOUNT_COMMANDOk

`func (o *UNIXConnector) GetDISABLE_ACCOUNT_COMMANDOk() (*string, bool)`

GetDISABLE_ACCOUNT_COMMANDOk returns a tuple with the DISABLE_ACCOUNT_COMMAND field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDISABLE_ACCOUNT_COMMAND

`func (o *UNIXConnector) SetDISABLE_ACCOUNT_COMMAND(v string)`

SetDISABLE_ACCOUNT_COMMAND sets DISABLE_ACCOUNT_COMMAND field to given value.

### HasDISABLE_ACCOUNT_COMMAND

`func (o *UNIXConnector) HasDISABLE_ACCOUNT_COMMAND() bool`

HasDISABLE_ACCOUNT_COMMAND returns a boolean if a field has been set.

### GetACCOUNT_ENTITLEMENT_MAPPING_COMMAND

`func (o *UNIXConnector) GetACCOUNT_ENTITLEMENT_MAPPING_COMMAND() string`

GetACCOUNT_ENTITLEMENT_MAPPING_COMMAND returns the ACCOUNT_ENTITLEMENT_MAPPING_COMMAND field if non-nil, zero value otherwise.

### GetACCOUNT_ENTITLEMENT_MAPPING_COMMANDOk

`func (o *UNIXConnector) GetACCOUNT_ENTITLEMENT_MAPPING_COMMANDOk() (*string, bool)`

GetACCOUNT_ENTITLEMENT_MAPPING_COMMANDOk returns a tuple with the ACCOUNT_ENTITLEMENT_MAPPING_COMMAND field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetACCOUNT_ENTITLEMENT_MAPPING_COMMAND

`func (o *UNIXConnector) SetACCOUNT_ENTITLEMENT_MAPPING_COMMAND(v string)`

SetACCOUNT_ENTITLEMENT_MAPPING_COMMAND sets ACCOUNT_ENTITLEMENT_MAPPING_COMMAND field to given value.

### HasACCOUNT_ENTITLEMENT_MAPPING_COMMAND

`func (o *UNIXConnector) HasACCOUNT_ENTITLEMENT_MAPPING_COMMAND() bool`

HasACCOUNT_ENTITLEMENT_MAPPING_COMMAND returns a boolean if a field has been set.

### GetPASSPHRASE

`func (o *UNIXConnector) GetPASSPHRASE() string`

GetPASSPHRASE returns the PASSPHRASE field if non-nil, zero value otherwise.

### GetPASSPHRASEOk

`func (o *UNIXConnector) GetPASSPHRASEOk() (*string, bool)`

GetPASSPHRASEOk returns a tuple with the PASSPHRASE field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPASSPHRASE

`func (o *UNIXConnector) SetPASSPHRASE(v string)`

SetPASSPHRASE sets PASSPHRASE field to given value.

### HasPASSPHRASE

`func (o *UNIXConnector) HasPASSPHRASE() bool`

HasPASSPHRASE returns a boolean if a field has been set.

### GetUPDATE_ACCOUNT_COMMAND

`func (o *UNIXConnector) GetUPDATE_ACCOUNT_COMMAND() string`

GetUPDATE_ACCOUNT_COMMAND returns the UPDATE_ACCOUNT_COMMAND field if non-nil, zero value otherwise.

### GetUPDATE_ACCOUNT_COMMANDOk

`func (o *UNIXConnector) GetUPDATE_ACCOUNT_COMMANDOk() (*string, bool)`

GetUPDATE_ACCOUNT_COMMANDOk returns a tuple with the UPDATE_ACCOUNT_COMMAND field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUPDATE_ACCOUNT_COMMAND

`func (o *UNIXConnector) SetUPDATE_ACCOUNT_COMMAND(v string)`

SetUPDATE_ACCOUNT_COMMAND sets UPDATE_ACCOUNT_COMMAND field to given value.

### HasUPDATE_ACCOUNT_COMMAND

`func (o *UNIXConnector) HasUPDATE_ACCOUNT_COMMAND() bool`

HasUPDATE_ACCOUNT_COMMAND returns a boolean if a field has been set.

### GetCREATE_GROUP_COMMAND

`func (o *UNIXConnector) GetCREATE_GROUP_COMMAND() string`

GetCREATE_GROUP_COMMAND returns the CREATE_GROUP_COMMAND field if non-nil, zero value otherwise.

### GetCREATE_GROUP_COMMANDOk

`func (o *UNIXConnector) GetCREATE_GROUP_COMMANDOk() (*string, bool)`

GetCREATE_GROUP_COMMANDOk returns a tuple with the CREATE_GROUP_COMMAND field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCREATE_GROUP_COMMAND

`func (o *UNIXConnector) SetCREATE_GROUP_COMMAND(v string)`

SetCREATE_GROUP_COMMAND sets CREATE_GROUP_COMMAND field to given value.

### HasCREATE_GROUP_COMMAND

`func (o *UNIXConnector) HasCREATE_GROUP_COMMAND() bool`

HasCREATE_GROUP_COMMAND returns a boolean if a field has been set.

### GetDELETE_GROUP_COMMAND

`func (o *UNIXConnector) GetDELETE_GROUP_COMMAND() string`

GetDELETE_GROUP_COMMAND returns the DELETE_GROUP_COMMAND field if non-nil, zero value otherwise.

### GetDELETE_GROUP_COMMANDOk

`func (o *UNIXConnector) GetDELETE_GROUP_COMMANDOk() (*string, bool)`

GetDELETE_GROUP_COMMANDOk returns a tuple with the DELETE_GROUP_COMMAND field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDELETE_GROUP_COMMAND

`func (o *UNIXConnector) SetDELETE_GROUP_COMMAND(v string)`

SetDELETE_GROUP_COMMAND sets DELETE_GROUP_COMMAND field to given value.

### HasDELETE_GROUP_COMMAND

`func (o *UNIXConnector) HasDELETE_GROUP_COMMAND() bool`

HasDELETE_GROUP_COMMAND returns a boolean if a field has been set.

### GetADD_GROUP_OWNER_COMMAND

`func (o *UNIXConnector) GetADD_GROUP_OWNER_COMMAND() string`

GetADD_GROUP_OWNER_COMMAND returns the ADD_GROUP_OWNER_COMMAND field if non-nil, zero value otherwise.

### GetADD_GROUP_OWNER_COMMANDOk

`func (o *UNIXConnector) GetADD_GROUP_OWNER_COMMANDOk() (*string, bool)`

GetADD_GROUP_OWNER_COMMANDOk returns a tuple with the ADD_GROUP_OWNER_COMMAND field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetADD_GROUP_OWNER_COMMAND

`func (o *UNIXConnector) SetADD_GROUP_OWNER_COMMAND(v string)`

SetADD_GROUP_OWNER_COMMAND sets ADD_GROUP_OWNER_COMMAND field to given value.

### HasADD_GROUP_OWNER_COMMAND

`func (o *UNIXConnector) HasADD_GROUP_OWNER_COMMAND() bool`

HasADD_GROUP_OWNER_COMMAND returns a boolean if a field has been set.

### GetADD_PRIMARY_GROUP_COMMAND

`func (o *UNIXConnector) GetADD_PRIMARY_GROUP_COMMAND() string`

GetADD_PRIMARY_GROUP_COMMAND returns the ADD_PRIMARY_GROUP_COMMAND field if non-nil, zero value otherwise.

### GetADD_PRIMARY_GROUP_COMMANDOk

`func (o *UNIXConnector) GetADD_PRIMARY_GROUP_COMMANDOk() (*string, bool)`

GetADD_PRIMARY_GROUP_COMMANDOk returns a tuple with the ADD_PRIMARY_GROUP_COMMAND field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetADD_PRIMARY_GROUP_COMMAND

`func (o *UNIXConnector) SetADD_PRIMARY_GROUP_COMMAND(v string)`

SetADD_PRIMARY_GROUP_COMMAND sets ADD_PRIMARY_GROUP_COMMAND field to given value.

### HasADD_PRIMARY_GROUP_COMMAND

`func (o *UNIXConnector) HasADD_PRIMARY_GROUP_COMMAND() bool`

HasADD_PRIMARY_GROUP_COMMAND returns a boolean if a field has been set.

### GetFIREFIGHTERID_GRANT_ACCESS_COMMAND

`func (o *UNIXConnector) GetFIREFIGHTERID_GRANT_ACCESS_COMMAND() string`

GetFIREFIGHTERID_GRANT_ACCESS_COMMAND returns the FIREFIGHTERID_GRANT_ACCESS_COMMAND field if non-nil, zero value otherwise.

### GetFIREFIGHTERID_GRANT_ACCESS_COMMANDOk

`func (o *UNIXConnector) GetFIREFIGHTERID_GRANT_ACCESS_COMMANDOk() (*string, bool)`

GetFIREFIGHTERID_GRANT_ACCESS_COMMANDOk returns a tuple with the FIREFIGHTERID_GRANT_ACCESS_COMMAND field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFIREFIGHTERID_GRANT_ACCESS_COMMAND

`func (o *UNIXConnector) SetFIREFIGHTERID_GRANT_ACCESS_COMMAND(v string)`

SetFIREFIGHTERID_GRANT_ACCESS_COMMAND sets FIREFIGHTERID_GRANT_ACCESS_COMMAND field to given value.

### HasFIREFIGHTERID_GRANT_ACCESS_COMMAND

`func (o *UNIXConnector) HasFIREFIGHTERID_GRANT_ACCESS_COMMAND() bool`

HasFIREFIGHTERID_GRANT_ACCESS_COMMAND returns a boolean if a field has been set.

### GetFIREFIGHTERID_REVOKE_ACCESS_COMMAND

`func (o *UNIXConnector) GetFIREFIGHTERID_REVOKE_ACCESS_COMMAND() string`

GetFIREFIGHTERID_REVOKE_ACCESS_COMMAND returns the FIREFIGHTERID_REVOKE_ACCESS_COMMAND field if non-nil, zero value otherwise.

### GetFIREFIGHTERID_REVOKE_ACCESS_COMMANDOk

`func (o *UNIXConnector) GetFIREFIGHTERID_REVOKE_ACCESS_COMMANDOk() (*string, bool)`

GetFIREFIGHTERID_REVOKE_ACCESS_COMMANDOk returns a tuple with the FIREFIGHTERID_REVOKE_ACCESS_COMMAND field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFIREFIGHTERID_REVOKE_ACCESS_COMMAND

`func (o *UNIXConnector) SetFIREFIGHTERID_REVOKE_ACCESS_COMMAND(v string)`

SetFIREFIGHTERID_REVOKE_ACCESS_COMMAND sets FIREFIGHTERID_REVOKE_ACCESS_COMMAND field to given value.

### HasFIREFIGHTERID_REVOKE_ACCESS_COMMAND

`func (o *UNIXConnector) HasFIREFIGHTERID_REVOKE_ACCESS_COMMAND() bool`

HasFIREFIGHTERID_REVOKE_ACCESS_COMMAND returns a boolean if a field has been set.

### GetINACTIVE_LOCK_ACCOUNT

`func (o *UNIXConnector) GetINACTIVE_LOCK_ACCOUNT() string`

GetINACTIVE_LOCK_ACCOUNT returns the INACTIVE_LOCK_ACCOUNT field if non-nil, zero value otherwise.

### GetINACTIVE_LOCK_ACCOUNTOk

`func (o *UNIXConnector) GetINACTIVE_LOCK_ACCOUNTOk() (*string, bool)`

GetINACTIVE_LOCK_ACCOUNTOk returns a tuple with the INACTIVE_LOCK_ACCOUNT field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetINACTIVE_LOCK_ACCOUNT

`func (o *UNIXConnector) SetINACTIVE_LOCK_ACCOUNT(v string)`

SetINACTIVE_LOCK_ACCOUNT sets INACTIVE_LOCK_ACCOUNT field to given value.

### HasINACTIVE_LOCK_ACCOUNT

`func (o *UNIXConnector) HasINACTIVE_LOCK_ACCOUNT() bool`

HasINACTIVE_LOCK_ACCOUNT returns a boolean if a field has been set.

### GetSTATUS_THRESHOLD_CONFIG

`func (o *UNIXConnector) GetSTATUS_THRESHOLD_CONFIG() string`

GetSTATUS_THRESHOLD_CONFIG returns the STATUS_THRESHOLD_CONFIG field if non-nil, zero value otherwise.

### GetSTATUS_THRESHOLD_CONFIGOk

`func (o *UNIXConnector) GetSTATUS_THRESHOLD_CONFIGOk() (*string, bool)`

GetSTATUS_THRESHOLD_CONFIGOk returns a tuple with the STATUS_THRESHOLD_CONFIG field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSTATUS_THRESHOLD_CONFIG

`func (o *UNIXConnector) SetSTATUS_THRESHOLD_CONFIG(v string)`

SetSTATUS_THRESHOLD_CONFIG sets STATUS_THRESHOLD_CONFIG field to given value.

### HasSTATUS_THRESHOLD_CONFIG

`func (o *UNIXConnector) HasSTATUS_THRESHOLD_CONFIG() bool`

HasSTATUS_THRESHOLD_CONFIG returns a boolean if a field has been set.

### GetCUSTOM_CONFIG_JSON

`func (o *UNIXConnector) GetCUSTOM_CONFIG_JSON() string`

GetCUSTOM_CONFIG_JSON returns the CUSTOM_CONFIG_JSON field if non-nil, zero value otherwise.

### GetCUSTOM_CONFIG_JSONOk

`func (o *UNIXConnector) GetCUSTOM_CONFIG_JSONOk() (*string, bool)`

GetCUSTOM_CONFIG_JSONOk returns a tuple with the CUSTOM_CONFIG_JSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCUSTOM_CONFIG_JSON

`func (o *UNIXConnector) SetCUSTOM_CONFIG_JSON(v string)`

SetCUSTOM_CONFIG_JSON sets CUSTOM_CONFIG_JSON field to given value.

### HasCUSTOM_CONFIG_JSON

`func (o *UNIXConnector) HasCUSTOM_CONFIG_JSON() bool`

HasCUSTOM_CONFIG_JSON returns a boolean if a field has been set.

### GetSSH_KEY

`func (o *UNIXConnector) GetSSH_KEY() string`

GetSSH_KEY returns the SSH_KEY field if non-nil, zero value otherwise.

### GetSSH_KEYOk

`func (o *UNIXConnector) GetSSH_KEYOk() (*string, bool)`

GetSSH_KEYOk returns a tuple with the SSH_KEY field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSSH_KEY

`func (o *UNIXConnector) SetSSH_KEY(v string)`

SetSSH_KEY sets SSH_KEY field to given value.

### HasSSH_KEY

`func (o *UNIXConnector) HasSSH_KEY() bool`

HasSSH_KEY returns a boolean if a field has been set.

### GetLOCK_ACCOUNT_COMMAND

`func (o *UNIXConnector) GetLOCK_ACCOUNT_COMMAND() string`

GetLOCK_ACCOUNT_COMMAND returns the LOCK_ACCOUNT_COMMAND field if non-nil, zero value otherwise.

### GetLOCK_ACCOUNT_COMMANDOk

`func (o *UNIXConnector) GetLOCK_ACCOUNT_COMMANDOk() (*string, bool)`

GetLOCK_ACCOUNT_COMMANDOk returns a tuple with the LOCK_ACCOUNT_COMMAND field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLOCK_ACCOUNT_COMMAND

`func (o *UNIXConnector) SetLOCK_ACCOUNT_COMMAND(v string)`

SetLOCK_ACCOUNT_COMMAND sets LOCK_ACCOUNT_COMMAND field to given value.

### HasLOCK_ACCOUNT_COMMAND

`func (o *UNIXConnector) HasLOCK_ACCOUNT_COMMAND() bool`

HasLOCK_ACCOUNT_COMMAND returns a boolean if a field has been set.

### GetUNLOCK_ACCOUNT_COMMAND

`func (o *UNIXConnector) GetUNLOCK_ACCOUNT_COMMAND() string`

GetUNLOCK_ACCOUNT_COMMAND returns the UNLOCK_ACCOUNT_COMMAND field if non-nil, zero value otherwise.

### GetUNLOCK_ACCOUNT_COMMANDOk

`func (o *UNIXConnector) GetUNLOCK_ACCOUNT_COMMANDOk() (*string, bool)`

GetUNLOCK_ACCOUNT_COMMANDOk returns a tuple with the UNLOCK_ACCOUNT_COMMAND field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUNLOCK_ACCOUNT_COMMAND

`func (o *UNIXConnector) SetUNLOCK_ACCOUNT_COMMAND(v string)`

SetUNLOCK_ACCOUNT_COMMAND sets UNLOCK_ACCOUNT_COMMAND field to given value.

### HasUNLOCK_ACCOUNT_COMMAND

`func (o *UNIXConnector) HasUNLOCK_ACCOUNT_COMMAND() bool`

HasUNLOCK_ACCOUNT_COMMAND returns a boolean if a field has been set.

### GetPassThroughConnectionDetails

`func (o *UNIXConnector) GetPassThroughConnectionDetails() string`

GetPassThroughConnectionDetails returns the PassThroughConnectionDetails field if non-nil, zero value otherwise.

### GetPassThroughConnectionDetailsOk

`func (o *UNIXConnector) GetPassThroughConnectionDetailsOk() (*string, bool)`

GetPassThroughConnectionDetailsOk returns a tuple with the PassThroughConnectionDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassThroughConnectionDetails

`func (o *UNIXConnector) SetPassThroughConnectionDetails(v string)`

SetPassThroughConnectionDetails sets PassThroughConnectionDetails field to given value.

### HasPassThroughConnectionDetails

`func (o *UNIXConnector) HasPassThroughConnectionDetails() bool`

HasPassThroughConnectionDetails returns a boolean if a field has been set.

### GetSSHPassThroughPassword

`func (o *UNIXConnector) GetSSHPassThroughPassword() string`

GetSSHPassThroughPassword returns the SSHPassThroughPassword field if non-nil, zero value otherwise.

### GetSSHPassThroughPasswordOk

`func (o *UNIXConnector) GetSSHPassThroughPasswordOk() (*string, bool)`

GetSSHPassThroughPasswordOk returns a tuple with the SSHPassThroughPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSSHPassThroughPassword

`func (o *UNIXConnector) SetSSHPassThroughPassword(v string)`

SetSSHPassThroughPassword sets SSHPassThroughPassword field to given value.

### HasSSHPassThroughPassword

`func (o *UNIXConnector) HasSSHPassThroughPassword() bool`

HasSSHPassThroughPassword returns a boolean if a field has been set.

### GetSSHPassThroughSSHKEY

`func (o *UNIXConnector) GetSSHPassThroughSSHKEY() string`

GetSSHPassThroughSSHKEY returns the SSHPassThroughSSHKEY field if non-nil, zero value otherwise.

### GetSSHPassThroughSSHKEYOk

`func (o *UNIXConnector) GetSSHPassThroughSSHKEYOk() (*string, bool)`

GetSSHPassThroughSSHKEYOk returns a tuple with the SSHPassThroughSSHKEY field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSSHPassThroughSSHKEY

`func (o *UNIXConnector) SetSSHPassThroughSSHKEY(v string)`

SetSSHPassThroughSSHKEY sets SSHPassThroughSSHKEY field to given value.

### HasSSHPassThroughSSHKEY

`func (o *UNIXConnector) HasSSHPassThroughSSHKEY() bool`

HasSSHPassThroughSSHKEY returns a boolean if a field has been set.

### GetSSHPassThroughPassphrase

`func (o *UNIXConnector) GetSSHPassThroughPassphrase() string`

GetSSHPassThroughPassphrase returns the SSHPassThroughPassphrase field if non-nil, zero value otherwise.

### GetSSHPassThroughPassphraseOk

`func (o *UNIXConnector) GetSSHPassThroughPassphraseOk() (*string, bool)`

GetSSHPassThroughPassphraseOk returns a tuple with the SSHPassThroughPassphrase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSSHPassThroughPassphrase

`func (o *UNIXConnector) SetSSHPassThroughPassphrase(v string)`

SetSSHPassThroughPassphrase sets SSHPassThroughPassphrase field to given value.

### HasSSHPassThroughPassphrase

`func (o *UNIXConnector) HasSSHPassThroughPassphrase() bool`

HasSSHPassThroughPassphrase returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


