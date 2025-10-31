# ADConnector

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**URL** | Pointer to **string** | LDAP or target system URL. | [optional] 
**USERNAME** | Pointer to **string** | LDAP or system admin username. | [optional] 
**PASSWORD** | **string** | Set the Password | 
**LDAP_OR_AD** | Pointer to **string** | Type of Endpoint - LDAP or AD, default value is AD | [optional] 
**ENTITLEMENT_ATTRIBUTE** | Pointer to **string** | Attribute used for entitlements. | [optional] 
**CHECKFORUNIQUE** | Pointer to **string** | Attribute uniqueness validation rule. | [optional] 
**GroupSearchBaseDN** | Pointer to **string** | Base DN for group search in LDAP. | [optional] 
**CreateUpdateMappings** | Pointer to **string** | Mapping used during group creation and updation. | [optional] 
**INCREMENTAL_CONFIG** | Pointer to **string** | Configuration for Incremental Import if the Incremental import is based on ChangeNumber. | [optional] 
**MAX_CHANGENUMBER** | Pointer to **string** | Maximum value of the changeNumber. | [optional] 
**READ_OPERATIONAL_ATTRIBUTES** | Pointer to **string** | Read operational attributes from LDAP. | [optional] 
**BASE** | Pointer to **string** | LDAP base DN. | [optional] 
**DC_LOCATOR** | Pointer to **string** | Windows connector connection name to connect and get domain controller data. | [optional] 
**STATUS_THRESHOLD_CONFIG** | Pointer to **string** | JSON configuration for account status thresholds. | [optional] 
**REMOVEACCOUNTACTION** | Pointer to **string** | Action performed on account removal. | [optional] 
**ACCOUNT_ATTRIBUTE** | Pointer to **string** | Map LDAP user attribute to EIC account attribute for import | [optional] 
**ACCOUNTNAMERULE** | Pointer to **string** | Rule to generate account name. | [optional] 
**ADVSEARCH** | Pointer to **string** |  | [optional] 
**SETDEFAULTPAGESIZE** | Pointer to **string** | Set default page size for LDAP queries. | [optional] 
**RESETANDCHANGEPASSWRDJSON** | Pointer to **string** | JSON actions for password reset/change. | [optional] 
**REUSEINACTIVEACCOUNT** | Pointer to **string** | Reuse inactive accounts setting. | [optional] 
**IMPORTJSON** | Pointer to **string** | JSON import configurations. | [optional] 
**SUPPORTEMPTYSTRING** | Pointer to **string** | Set to TRUE to send an empty value or null string during provisioning. | [optional] 
**ENABLEACCOUNTJSON** | Pointer to **string** | JSON configuration to enable account actions. | [optional] 
**PAGE_SIZE** | Pointer to **string** | LDAP pagination page size. | [optional] 
**USER_ATTRIBUTE** | Pointer to **string** | Map LDAP user attribute to EIC user attribute for import. | [optional] 
**DEFAULT_USER_ROLE** | Pointer to **string** | Default SAV Role to be assigned to all the new users that gets imported via User Import | [optional] 
**SEARCHFILTER** | Pointer to **string** | LDAP search filter. | [optional] 
**ENDPOINTS_FILTER** | Pointer to **string** | Provide the configuration to create Child Endpoints and import associated accounts and entitlements | [optional] 
**CREATEACCOUNTJSON** | Pointer to **string** | To create an account, specify the value of the CREATEACCOUNTJSON parameter. | [optional] 
**UPDATEACCOUNTJSON** | Pointer to **string** | To update an existing account, specify the value of the UPDATEACCOUNTJSON parameter. | [optional] 
**REUSEACCOUNTJSON** | Pointer to **string** | To retain the suspended user accounts, specify the value of the REUSEACCOUNTJSON parameter. This is a mandatory attribute. | [optional] 
**ENFORCE_TREE_DELETION** | Pointer to **string** | Enforce tree deletion in LDAP. | [optional] 
**ADVANCE_FILTER_JSON** | Pointer to **string** |  | [optional] 
**FILTER** | Pointer to **string** | Filters the objects that will be returned. | [optional] 
**OBJECTFILTER** | Pointer to **string** | LDAP object filter. | [optional] 
**UPDATEUSERJSON** | Pointer to **string** | To update an existing user, specify the value of the UPDATEUSERJSON parameter. | [optional] 
**Saveconnection** | Pointer to **string** | Flag to save connection permanently. | [optional] 
**Systemname** | Pointer to **string** | Associated system name. | [optional] 
**SETRANDOMPASSWORD** | Pointer to **string** | Set random password option. | [optional] 
**PASSWORD_MIN_LENGTH** | Pointer to **string** | Minimum password length. | [optional] 
**PASSWORD_MAX_LENGTH** | Pointer to **string** | Maximum password length. | [optional] 
**PASSWORD_NOOFCAPSALPHA** | Pointer to **string** | Specify the number of special characters required for the random password. | [optional] 
**PASSWORD_NOOFSPLCHARS** | Pointer to **string** | Number of special characters in password. | [optional] 
**PASSWORD_NOOFDIGITS** | Pointer to **string** | Specify the Number of digits required for the random password. | [optional] 
**GroupImportMapping** | Pointer to **string** | Map LDAP group attribute to EIC entitlement attribute for import | [optional] 
**UNLOCKACCOUNTJSON** | Pointer to **string** | JSON configuration to unlock accounts. | [optional] 
**STATUSKEYJSON** | Pointer to **string** | JSON configuration for account statuses. | [optional] 
**ENABLEGROUPMANAGEMENT** | Pointer to **string** | Specify TRUE to enable Group Management for AD. | [optional] 
**DISABLEACCOUNTJSON** | Pointer to **string** | JSON to specify the different attributes to be checked and action to be performed for disabling an active account. | [optional] 
**MODIFYUSERDATAJSON** | Pointer to **string** | Specify this parameter to use the inline processor for transforming the data during user import. | [optional] 
**ORG_BASE** | Pointer to **string** | Organization BASE for provision job. | [optional] 
**ORGANIZATION_ATTRIBUTE** | Pointer to **string** | Organization Attributes. | [optional] 
**ORGIMPORTJSON** | Pointer to **string** |  | [optional] 
**CREATEORGJSON** | Pointer to **string** | JSON to specify different attributes for Organization Create Provisioning Job. | [optional] 
**UPDATEORGJSON** | Pointer to **string** | JSON to specify different attributes for Organization Update Provisioning Job. | [optional] 
**ConfigJSON** | Pointer to **string** | We can use this attribute to define the connectionTimeoutConfig. | [optional] 
**LAST_IMPORT_TIME** | Pointer to **string** |  | [optional] 
**PAM_CONFIG** | Pointer to **string** | JSON to specify Bootstrap Config. | [optional] 

## Methods

### NewADConnector

`func NewADConnector(pASSWORD string, ) *ADConnector`

NewADConnector instantiates a new ADConnector object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewADConnectorWithDefaults

`func NewADConnectorWithDefaults() *ADConnector`

NewADConnectorWithDefaults instantiates a new ADConnector object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetURL

`func (o *ADConnector) GetURL() string`

GetURL returns the URL field if non-nil, zero value otherwise.

### GetURLOk

`func (o *ADConnector) GetURLOk() (*string, bool)`

GetURLOk returns a tuple with the URL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetURL

`func (o *ADConnector) SetURL(v string)`

SetURL sets URL field to given value.

### HasURL

`func (o *ADConnector) HasURL() bool`

HasURL returns a boolean if a field has been set.

### GetUSERNAME

`func (o *ADConnector) GetUSERNAME() string`

GetUSERNAME returns the USERNAME field if non-nil, zero value otherwise.

### GetUSERNAMEOk

`func (o *ADConnector) GetUSERNAMEOk() (*string, bool)`

GetUSERNAMEOk returns a tuple with the USERNAME field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUSERNAME

`func (o *ADConnector) SetUSERNAME(v string)`

SetUSERNAME sets USERNAME field to given value.

### HasUSERNAME

`func (o *ADConnector) HasUSERNAME() bool`

HasUSERNAME returns a boolean if a field has been set.

### GetPASSWORD

`func (o *ADConnector) GetPASSWORD() string`

GetPASSWORD returns the PASSWORD field if non-nil, zero value otherwise.

### GetPASSWORDOk

`func (o *ADConnector) GetPASSWORDOk() (*string, bool)`

GetPASSWORDOk returns a tuple with the PASSWORD field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPASSWORD

`func (o *ADConnector) SetPASSWORD(v string)`

SetPASSWORD sets PASSWORD field to given value.


### GetLDAP_OR_AD

`func (o *ADConnector) GetLDAP_OR_AD() string`

GetLDAP_OR_AD returns the LDAP_OR_AD field if non-nil, zero value otherwise.

### GetLDAP_OR_ADOk

`func (o *ADConnector) GetLDAP_OR_ADOk() (*string, bool)`

GetLDAP_OR_ADOk returns a tuple with the LDAP_OR_AD field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLDAP_OR_AD

`func (o *ADConnector) SetLDAP_OR_AD(v string)`

SetLDAP_OR_AD sets LDAP_OR_AD field to given value.

### HasLDAP_OR_AD

`func (o *ADConnector) HasLDAP_OR_AD() bool`

HasLDAP_OR_AD returns a boolean if a field has been set.

### GetENTITLEMENT_ATTRIBUTE

`func (o *ADConnector) GetENTITLEMENT_ATTRIBUTE() string`

GetENTITLEMENT_ATTRIBUTE returns the ENTITLEMENT_ATTRIBUTE field if non-nil, zero value otherwise.

### GetENTITLEMENT_ATTRIBUTEOk

`func (o *ADConnector) GetENTITLEMENT_ATTRIBUTEOk() (*string, bool)`

GetENTITLEMENT_ATTRIBUTEOk returns a tuple with the ENTITLEMENT_ATTRIBUTE field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetENTITLEMENT_ATTRIBUTE

`func (o *ADConnector) SetENTITLEMENT_ATTRIBUTE(v string)`

SetENTITLEMENT_ATTRIBUTE sets ENTITLEMENT_ATTRIBUTE field to given value.

### HasENTITLEMENT_ATTRIBUTE

`func (o *ADConnector) HasENTITLEMENT_ATTRIBUTE() bool`

HasENTITLEMENT_ATTRIBUTE returns a boolean if a field has been set.

### GetCHECKFORUNIQUE

`func (o *ADConnector) GetCHECKFORUNIQUE() string`

GetCHECKFORUNIQUE returns the CHECKFORUNIQUE field if non-nil, zero value otherwise.

### GetCHECKFORUNIQUEOk

`func (o *ADConnector) GetCHECKFORUNIQUEOk() (*string, bool)`

GetCHECKFORUNIQUEOk returns a tuple with the CHECKFORUNIQUE field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCHECKFORUNIQUE

`func (o *ADConnector) SetCHECKFORUNIQUE(v string)`

SetCHECKFORUNIQUE sets CHECKFORUNIQUE field to given value.

### HasCHECKFORUNIQUE

`func (o *ADConnector) HasCHECKFORUNIQUE() bool`

HasCHECKFORUNIQUE returns a boolean if a field has been set.

### GetGroupSearchBaseDN

`func (o *ADConnector) GetGroupSearchBaseDN() string`

GetGroupSearchBaseDN returns the GroupSearchBaseDN field if non-nil, zero value otherwise.

### GetGroupSearchBaseDNOk

`func (o *ADConnector) GetGroupSearchBaseDNOk() (*string, bool)`

GetGroupSearchBaseDNOk returns a tuple with the GroupSearchBaseDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupSearchBaseDN

`func (o *ADConnector) SetGroupSearchBaseDN(v string)`

SetGroupSearchBaseDN sets GroupSearchBaseDN field to given value.

### HasGroupSearchBaseDN

`func (o *ADConnector) HasGroupSearchBaseDN() bool`

HasGroupSearchBaseDN returns a boolean if a field has been set.

### GetCreateUpdateMappings

`func (o *ADConnector) GetCreateUpdateMappings() string`

GetCreateUpdateMappings returns the CreateUpdateMappings field if non-nil, zero value otherwise.

### GetCreateUpdateMappingsOk

`func (o *ADConnector) GetCreateUpdateMappingsOk() (*string, bool)`

GetCreateUpdateMappingsOk returns a tuple with the CreateUpdateMappings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateUpdateMappings

`func (o *ADConnector) SetCreateUpdateMappings(v string)`

SetCreateUpdateMappings sets CreateUpdateMappings field to given value.

### HasCreateUpdateMappings

`func (o *ADConnector) HasCreateUpdateMappings() bool`

HasCreateUpdateMappings returns a boolean if a field has been set.

### GetINCREMENTAL_CONFIG

`func (o *ADConnector) GetINCREMENTAL_CONFIG() string`

GetINCREMENTAL_CONFIG returns the INCREMENTAL_CONFIG field if non-nil, zero value otherwise.

### GetINCREMENTAL_CONFIGOk

`func (o *ADConnector) GetINCREMENTAL_CONFIGOk() (*string, bool)`

GetINCREMENTAL_CONFIGOk returns a tuple with the INCREMENTAL_CONFIG field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetINCREMENTAL_CONFIG

`func (o *ADConnector) SetINCREMENTAL_CONFIG(v string)`

SetINCREMENTAL_CONFIG sets INCREMENTAL_CONFIG field to given value.

### HasINCREMENTAL_CONFIG

`func (o *ADConnector) HasINCREMENTAL_CONFIG() bool`

HasINCREMENTAL_CONFIG returns a boolean if a field has been set.

### GetMAX_CHANGENUMBER

`func (o *ADConnector) GetMAX_CHANGENUMBER() string`

GetMAX_CHANGENUMBER returns the MAX_CHANGENUMBER field if non-nil, zero value otherwise.

### GetMAX_CHANGENUMBEROk

`func (o *ADConnector) GetMAX_CHANGENUMBEROk() (*string, bool)`

GetMAX_CHANGENUMBEROk returns a tuple with the MAX_CHANGENUMBER field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMAX_CHANGENUMBER

`func (o *ADConnector) SetMAX_CHANGENUMBER(v string)`

SetMAX_CHANGENUMBER sets MAX_CHANGENUMBER field to given value.

### HasMAX_CHANGENUMBER

`func (o *ADConnector) HasMAX_CHANGENUMBER() bool`

HasMAX_CHANGENUMBER returns a boolean if a field has been set.

### GetREAD_OPERATIONAL_ATTRIBUTES

`func (o *ADConnector) GetREAD_OPERATIONAL_ATTRIBUTES() string`

GetREAD_OPERATIONAL_ATTRIBUTES returns the READ_OPERATIONAL_ATTRIBUTES field if non-nil, zero value otherwise.

### GetREAD_OPERATIONAL_ATTRIBUTESOk

`func (o *ADConnector) GetREAD_OPERATIONAL_ATTRIBUTESOk() (*string, bool)`

GetREAD_OPERATIONAL_ATTRIBUTESOk returns a tuple with the READ_OPERATIONAL_ATTRIBUTES field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetREAD_OPERATIONAL_ATTRIBUTES

`func (o *ADConnector) SetREAD_OPERATIONAL_ATTRIBUTES(v string)`

SetREAD_OPERATIONAL_ATTRIBUTES sets READ_OPERATIONAL_ATTRIBUTES field to given value.

### HasREAD_OPERATIONAL_ATTRIBUTES

`func (o *ADConnector) HasREAD_OPERATIONAL_ATTRIBUTES() bool`

HasREAD_OPERATIONAL_ATTRIBUTES returns a boolean if a field has been set.

### GetBASE

`func (o *ADConnector) GetBASE() string`

GetBASE returns the BASE field if non-nil, zero value otherwise.

### GetBASEOk

`func (o *ADConnector) GetBASEOk() (*string, bool)`

GetBASEOk returns a tuple with the BASE field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBASE

`func (o *ADConnector) SetBASE(v string)`

SetBASE sets BASE field to given value.

### HasBASE

`func (o *ADConnector) HasBASE() bool`

HasBASE returns a boolean if a field has been set.

### GetDC_LOCATOR

`func (o *ADConnector) GetDC_LOCATOR() string`

GetDC_LOCATOR returns the DC_LOCATOR field if non-nil, zero value otherwise.

### GetDC_LOCATOROk

`func (o *ADConnector) GetDC_LOCATOROk() (*string, bool)`

GetDC_LOCATOROk returns a tuple with the DC_LOCATOR field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDC_LOCATOR

`func (o *ADConnector) SetDC_LOCATOR(v string)`

SetDC_LOCATOR sets DC_LOCATOR field to given value.

### HasDC_LOCATOR

`func (o *ADConnector) HasDC_LOCATOR() bool`

HasDC_LOCATOR returns a boolean if a field has been set.

### GetSTATUS_THRESHOLD_CONFIG

`func (o *ADConnector) GetSTATUS_THRESHOLD_CONFIG() string`

GetSTATUS_THRESHOLD_CONFIG returns the STATUS_THRESHOLD_CONFIG field if non-nil, zero value otherwise.

### GetSTATUS_THRESHOLD_CONFIGOk

`func (o *ADConnector) GetSTATUS_THRESHOLD_CONFIGOk() (*string, bool)`

GetSTATUS_THRESHOLD_CONFIGOk returns a tuple with the STATUS_THRESHOLD_CONFIG field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSTATUS_THRESHOLD_CONFIG

`func (o *ADConnector) SetSTATUS_THRESHOLD_CONFIG(v string)`

SetSTATUS_THRESHOLD_CONFIG sets STATUS_THRESHOLD_CONFIG field to given value.

### HasSTATUS_THRESHOLD_CONFIG

`func (o *ADConnector) HasSTATUS_THRESHOLD_CONFIG() bool`

HasSTATUS_THRESHOLD_CONFIG returns a boolean if a field has been set.

### GetREMOVEACCOUNTACTION

`func (o *ADConnector) GetREMOVEACCOUNTACTION() string`

GetREMOVEACCOUNTACTION returns the REMOVEACCOUNTACTION field if non-nil, zero value otherwise.

### GetREMOVEACCOUNTACTIONOk

`func (o *ADConnector) GetREMOVEACCOUNTACTIONOk() (*string, bool)`

GetREMOVEACCOUNTACTIONOk returns a tuple with the REMOVEACCOUNTACTION field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetREMOVEACCOUNTACTION

`func (o *ADConnector) SetREMOVEACCOUNTACTION(v string)`

SetREMOVEACCOUNTACTION sets REMOVEACCOUNTACTION field to given value.

### HasREMOVEACCOUNTACTION

`func (o *ADConnector) HasREMOVEACCOUNTACTION() bool`

HasREMOVEACCOUNTACTION returns a boolean if a field has been set.

### GetACCOUNT_ATTRIBUTE

`func (o *ADConnector) GetACCOUNT_ATTRIBUTE() string`

GetACCOUNT_ATTRIBUTE returns the ACCOUNT_ATTRIBUTE field if non-nil, zero value otherwise.

### GetACCOUNT_ATTRIBUTEOk

`func (o *ADConnector) GetACCOUNT_ATTRIBUTEOk() (*string, bool)`

GetACCOUNT_ATTRIBUTEOk returns a tuple with the ACCOUNT_ATTRIBUTE field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetACCOUNT_ATTRIBUTE

`func (o *ADConnector) SetACCOUNT_ATTRIBUTE(v string)`

SetACCOUNT_ATTRIBUTE sets ACCOUNT_ATTRIBUTE field to given value.

### HasACCOUNT_ATTRIBUTE

`func (o *ADConnector) HasACCOUNT_ATTRIBUTE() bool`

HasACCOUNT_ATTRIBUTE returns a boolean if a field has been set.

### GetACCOUNTNAMERULE

`func (o *ADConnector) GetACCOUNTNAMERULE() string`

GetACCOUNTNAMERULE returns the ACCOUNTNAMERULE field if non-nil, zero value otherwise.

### GetACCOUNTNAMERULEOk

`func (o *ADConnector) GetACCOUNTNAMERULEOk() (*string, bool)`

GetACCOUNTNAMERULEOk returns a tuple with the ACCOUNTNAMERULE field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetACCOUNTNAMERULE

`func (o *ADConnector) SetACCOUNTNAMERULE(v string)`

SetACCOUNTNAMERULE sets ACCOUNTNAMERULE field to given value.

### HasACCOUNTNAMERULE

`func (o *ADConnector) HasACCOUNTNAMERULE() bool`

HasACCOUNTNAMERULE returns a boolean if a field has been set.

### GetADVSEARCH

`func (o *ADConnector) GetADVSEARCH() string`

GetADVSEARCH returns the ADVSEARCH field if non-nil, zero value otherwise.

### GetADVSEARCHOk

`func (o *ADConnector) GetADVSEARCHOk() (*string, bool)`

GetADVSEARCHOk returns a tuple with the ADVSEARCH field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetADVSEARCH

`func (o *ADConnector) SetADVSEARCH(v string)`

SetADVSEARCH sets ADVSEARCH field to given value.

### HasADVSEARCH

`func (o *ADConnector) HasADVSEARCH() bool`

HasADVSEARCH returns a boolean if a field has been set.

### GetSETDEFAULTPAGESIZE

`func (o *ADConnector) GetSETDEFAULTPAGESIZE() string`

GetSETDEFAULTPAGESIZE returns the SETDEFAULTPAGESIZE field if non-nil, zero value otherwise.

### GetSETDEFAULTPAGESIZEOk

`func (o *ADConnector) GetSETDEFAULTPAGESIZEOk() (*string, bool)`

GetSETDEFAULTPAGESIZEOk returns a tuple with the SETDEFAULTPAGESIZE field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSETDEFAULTPAGESIZE

`func (o *ADConnector) SetSETDEFAULTPAGESIZE(v string)`

SetSETDEFAULTPAGESIZE sets SETDEFAULTPAGESIZE field to given value.

### HasSETDEFAULTPAGESIZE

`func (o *ADConnector) HasSETDEFAULTPAGESIZE() bool`

HasSETDEFAULTPAGESIZE returns a boolean if a field has been set.

### GetRESETANDCHANGEPASSWRDJSON

`func (o *ADConnector) GetRESETANDCHANGEPASSWRDJSON() string`

GetRESETANDCHANGEPASSWRDJSON returns the RESETANDCHANGEPASSWRDJSON field if non-nil, zero value otherwise.

### GetRESETANDCHANGEPASSWRDJSONOk

`func (o *ADConnector) GetRESETANDCHANGEPASSWRDJSONOk() (*string, bool)`

GetRESETANDCHANGEPASSWRDJSONOk returns a tuple with the RESETANDCHANGEPASSWRDJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRESETANDCHANGEPASSWRDJSON

`func (o *ADConnector) SetRESETANDCHANGEPASSWRDJSON(v string)`

SetRESETANDCHANGEPASSWRDJSON sets RESETANDCHANGEPASSWRDJSON field to given value.

### HasRESETANDCHANGEPASSWRDJSON

`func (o *ADConnector) HasRESETANDCHANGEPASSWRDJSON() bool`

HasRESETANDCHANGEPASSWRDJSON returns a boolean if a field has been set.

### GetREUSEINACTIVEACCOUNT

`func (o *ADConnector) GetREUSEINACTIVEACCOUNT() string`

GetREUSEINACTIVEACCOUNT returns the REUSEINACTIVEACCOUNT field if non-nil, zero value otherwise.

### GetREUSEINACTIVEACCOUNTOk

`func (o *ADConnector) GetREUSEINACTIVEACCOUNTOk() (*string, bool)`

GetREUSEINACTIVEACCOUNTOk returns a tuple with the REUSEINACTIVEACCOUNT field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetREUSEINACTIVEACCOUNT

`func (o *ADConnector) SetREUSEINACTIVEACCOUNT(v string)`

SetREUSEINACTIVEACCOUNT sets REUSEINACTIVEACCOUNT field to given value.

### HasREUSEINACTIVEACCOUNT

`func (o *ADConnector) HasREUSEINACTIVEACCOUNT() bool`

HasREUSEINACTIVEACCOUNT returns a boolean if a field has been set.

### GetIMPORTJSON

`func (o *ADConnector) GetIMPORTJSON() string`

GetIMPORTJSON returns the IMPORTJSON field if non-nil, zero value otherwise.

### GetIMPORTJSONOk

`func (o *ADConnector) GetIMPORTJSONOk() (*string, bool)`

GetIMPORTJSONOk returns a tuple with the IMPORTJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIMPORTJSON

`func (o *ADConnector) SetIMPORTJSON(v string)`

SetIMPORTJSON sets IMPORTJSON field to given value.

### HasIMPORTJSON

`func (o *ADConnector) HasIMPORTJSON() bool`

HasIMPORTJSON returns a boolean if a field has been set.

### GetSUPPORTEMPTYSTRING

`func (o *ADConnector) GetSUPPORTEMPTYSTRING() string`

GetSUPPORTEMPTYSTRING returns the SUPPORTEMPTYSTRING field if non-nil, zero value otherwise.

### GetSUPPORTEMPTYSTRINGOk

`func (o *ADConnector) GetSUPPORTEMPTYSTRINGOk() (*string, bool)`

GetSUPPORTEMPTYSTRINGOk returns a tuple with the SUPPORTEMPTYSTRING field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSUPPORTEMPTYSTRING

`func (o *ADConnector) SetSUPPORTEMPTYSTRING(v string)`

SetSUPPORTEMPTYSTRING sets SUPPORTEMPTYSTRING field to given value.

### HasSUPPORTEMPTYSTRING

`func (o *ADConnector) HasSUPPORTEMPTYSTRING() bool`

HasSUPPORTEMPTYSTRING returns a boolean if a field has been set.

### GetENABLEACCOUNTJSON

`func (o *ADConnector) GetENABLEACCOUNTJSON() string`

GetENABLEACCOUNTJSON returns the ENABLEACCOUNTJSON field if non-nil, zero value otherwise.

### GetENABLEACCOUNTJSONOk

`func (o *ADConnector) GetENABLEACCOUNTJSONOk() (*string, bool)`

GetENABLEACCOUNTJSONOk returns a tuple with the ENABLEACCOUNTJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetENABLEACCOUNTJSON

`func (o *ADConnector) SetENABLEACCOUNTJSON(v string)`

SetENABLEACCOUNTJSON sets ENABLEACCOUNTJSON field to given value.

### HasENABLEACCOUNTJSON

`func (o *ADConnector) HasENABLEACCOUNTJSON() bool`

HasENABLEACCOUNTJSON returns a boolean if a field has been set.

### GetPAGE_SIZE

`func (o *ADConnector) GetPAGE_SIZE() string`

GetPAGE_SIZE returns the PAGE_SIZE field if non-nil, zero value otherwise.

### GetPAGE_SIZEOk

`func (o *ADConnector) GetPAGE_SIZEOk() (*string, bool)`

GetPAGE_SIZEOk returns a tuple with the PAGE_SIZE field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPAGE_SIZE

`func (o *ADConnector) SetPAGE_SIZE(v string)`

SetPAGE_SIZE sets PAGE_SIZE field to given value.

### HasPAGE_SIZE

`func (o *ADConnector) HasPAGE_SIZE() bool`

HasPAGE_SIZE returns a boolean if a field has been set.

### GetUSER_ATTRIBUTE

`func (o *ADConnector) GetUSER_ATTRIBUTE() string`

GetUSER_ATTRIBUTE returns the USER_ATTRIBUTE field if non-nil, zero value otherwise.

### GetUSER_ATTRIBUTEOk

`func (o *ADConnector) GetUSER_ATTRIBUTEOk() (*string, bool)`

GetUSER_ATTRIBUTEOk returns a tuple with the USER_ATTRIBUTE field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUSER_ATTRIBUTE

`func (o *ADConnector) SetUSER_ATTRIBUTE(v string)`

SetUSER_ATTRIBUTE sets USER_ATTRIBUTE field to given value.

### HasUSER_ATTRIBUTE

`func (o *ADConnector) HasUSER_ATTRIBUTE() bool`

HasUSER_ATTRIBUTE returns a boolean if a field has been set.

### GetDEFAULT_USER_ROLE

`func (o *ADConnector) GetDEFAULT_USER_ROLE() string`

GetDEFAULT_USER_ROLE returns the DEFAULT_USER_ROLE field if non-nil, zero value otherwise.

### GetDEFAULT_USER_ROLEOk

`func (o *ADConnector) GetDEFAULT_USER_ROLEOk() (*string, bool)`

GetDEFAULT_USER_ROLEOk returns a tuple with the DEFAULT_USER_ROLE field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDEFAULT_USER_ROLE

`func (o *ADConnector) SetDEFAULT_USER_ROLE(v string)`

SetDEFAULT_USER_ROLE sets DEFAULT_USER_ROLE field to given value.

### HasDEFAULT_USER_ROLE

`func (o *ADConnector) HasDEFAULT_USER_ROLE() bool`

HasDEFAULT_USER_ROLE returns a boolean if a field has been set.

### GetSEARCHFILTER

`func (o *ADConnector) GetSEARCHFILTER() string`

GetSEARCHFILTER returns the SEARCHFILTER field if non-nil, zero value otherwise.

### GetSEARCHFILTEROk

`func (o *ADConnector) GetSEARCHFILTEROk() (*string, bool)`

GetSEARCHFILTEROk returns a tuple with the SEARCHFILTER field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEARCHFILTER

`func (o *ADConnector) SetSEARCHFILTER(v string)`

SetSEARCHFILTER sets SEARCHFILTER field to given value.

### HasSEARCHFILTER

`func (o *ADConnector) HasSEARCHFILTER() bool`

HasSEARCHFILTER returns a boolean if a field has been set.

### GetENDPOINTS_FILTER

`func (o *ADConnector) GetENDPOINTS_FILTER() string`

GetENDPOINTS_FILTER returns the ENDPOINTS_FILTER field if non-nil, zero value otherwise.

### GetENDPOINTS_FILTEROk

`func (o *ADConnector) GetENDPOINTS_FILTEROk() (*string, bool)`

GetENDPOINTS_FILTEROk returns a tuple with the ENDPOINTS_FILTER field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetENDPOINTS_FILTER

`func (o *ADConnector) SetENDPOINTS_FILTER(v string)`

SetENDPOINTS_FILTER sets ENDPOINTS_FILTER field to given value.

### HasENDPOINTS_FILTER

`func (o *ADConnector) HasENDPOINTS_FILTER() bool`

HasENDPOINTS_FILTER returns a boolean if a field has been set.

### GetCREATEACCOUNTJSON

`func (o *ADConnector) GetCREATEACCOUNTJSON() string`

GetCREATEACCOUNTJSON returns the CREATEACCOUNTJSON field if non-nil, zero value otherwise.

### GetCREATEACCOUNTJSONOk

`func (o *ADConnector) GetCREATEACCOUNTJSONOk() (*string, bool)`

GetCREATEACCOUNTJSONOk returns a tuple with the CREATEACCOUNTJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCREATEACCOUNTJSON

`func (o *ADConnector) SetCREATEACCOUNTJSON(v string)`

SetCREATEACCOUNTJSON sets CREATEACCOUNTJSON field to given value.

### HasCREATEACCOUNTJSON

`func (o *ADConnector) HasCREATEACCOUNTJSON() bool`

HasCREATEACCOUNTJSON returns a boolean if a field has been set.

### GetUPDATEACCOUNTJSON

`func (o *ADConnector) GetUPDATEACCOUNTJSON() string`

GetUPDATEACCOUNTJSON returns the UPDATEACCOUNTJSON field if non-nil, zero value otherwise.

### GetUPDATEACCOUNTJSONOk

`func (o *ADConnector) GetUPDATEACCOUNTJSONOk() (*string, bool)`

GetUPDATEACCOUNTJSONOk returns a tuple with the UPDATEACCOUNTJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUPDATEACCOUNTJSON

`func (o *ADConnector) SetUPDATEACCOUNTJSON(v string)`

SetUPDATEACCOUNTJSON sets UPDATEACCOUNTJSON field to given value.

### HasUPDATEACCOUNTJSON

`func (o *ADConnector) HasUPDATEACCOUNTJSON() bool`

HasUPDATEACCOUNTJSON returns a boolean if a field has been set.

### GetREUSEACCOUNTJSON

`func (o *ADConnector) GetREUSEACCOUNTJSON() string`

GetREUSEACCOUNTJSON returns the REUSEACCOUNTJSON field if non-nil, zero value otherwise.

### GetREUSEACCOUNTJSONOk

`func (o *ADConnector) GetREUSEACCOUNTJSONOk() (*string, bool)`

GetREUSEACCOUNTJSONOk returns a tuple with the REUSEACCOUNTJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetREUSEACCOUNTJSON

`func (o *ADConnector) SetREUSEACCOUNTJSON(v string)`

SetREUSEACCOUNTJSON sets REUSEACCOUNTJSON field to given value.

### HasREUSEACCOUNTJSON

`func (o *ADConnector) HasREUSEACCOUNTJSON() bool`

HasREUSEACCOUNTJSON returns a boolean if a field has been set.

### GetENFORCE_TREE_DELETION

`func (o *ADConnector) GetENFORCE_TREE_DELETION() string`

GetENFORCE_TREE_DELETION returns the ENFORCE_TREE_DELETION field if non-nil, zero value otherwise.

### GetENFORCE_TREE_DELETIONOk

`func (o *ADConnector) GetENFORCE_TREE_DELETIONOk() (*string, bool)`

GetENFORCE_TREE_DELETIONOk returns a tuple with the ENFORCE_TREE_DELETION field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetENFORCE_TREE_DELETION

`func (o *ADConnector) SetENFORCE_TREE_DELETION(v string)`

SetENFORCE_TREE_DELETION sets ENFORCE_TREE_DELETION field to given value.

### HasENFORCE_TREE_DELETION

`func (o *ADConnector) HasENFORCE_TREE_DELETION() bool`

HasENFORCE_TREE_DELETION returns a boolean if a field has been set.

### GetADVANCE_FILTER_JSON

`func (o *ADConnector) GetADVANCE_FILTER_JSON() string`

GetADVANCE_FILTER_JSON returns the ADVANCE_FILTER_JSON field if non-nil, zero value otherwise.

### GetADVANCE_FILTER_JSONOk

`func (o *ADConnector) GetADVANCE_FILTER_JSONOk() (*string, bool)`

GetADVANCE_FILTER_JSONOk returns a tuple with the ADVANCE_FILTER_JSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetADVANCE_FILTER_JSON

`func (o *ADConnector) SetADVANCE_FILTER_JSON(v string)`

SetADVANCE_FILTER_JSON sets ADVANCE_FILTER_JSON field to given value.

### HasADVANCE_FILTER_JSON

`func (o *ADConnector) HasADVANCE_FILTER_JSON() bool`

HasADVANCE_FILTER_JSON returns a boolean if a field has been set.

### GetFILTER

`func (o *ADConnector) GetFILTER() string`

GetFILTER returns the FILTER field if non-nil, zero value otherwise.

### GetFILTEROk

`func (o *ADConnector) GetFILTEROk() (*string, bool)`

GetFILTEROk returns a tuple with the FILTER field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFILTER

`func (o *ADConnector) SetFILTER(v string)`

SetFILTER sets FILTER field to given value.

### HasFILTER

`func (o *ADConnector) HasFILTER() bool`

HasFILTER returns a boolean if a field has been set.

### GetOBJECTFILTER

`func (o *ADConnector) GetOBJECTFILTER() string`

GetOBJECTFILTER returns the OBJECTFILTER field if non-nil, zero value otherwise.

### GetOBJECTFILTEROk

`func (o *ADConnector) GetOBJECTFILTEROk() (*string, bool)`

GetOBJECTFILTEROk returns a tuple with the OBJECTFILTER field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOBJECTFILTER

`func (o *ADConnector) SetOBJECTFILTER(v string)`

SetOBJECTFILTER sets OBJECTFILTER field to given value.

### HasOBJECTFILTER

`func (o *ADConnector) HasOBJECTFILTER() bool`

HasOBJECTFILTER returns a boolean if a field has been set.

### GetUPDATEUSERJSON

`func (o *ADConnector) GetUPDATEUSERJSON() string`

GetUPDATEUSERJSON returns the UPDATEUSERJSON field if non-nil, zero value otherwise.

### GetUPDATEUSERJSONOk

`func (o *ADConnector) GetUPDATEUSERJSONOk() (*string, bool)`

GetUPDATEUSERJSONOk returns a tuple with the UPDATEUSERJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUPDATEUSERJSON

`func (o *ADConnector) SetUPDATEUSERJSON(v string)`

SetUPDATEUSERJSON sets UPDATEUSERJSON field to given value.

### HasUPDATEUSERJSON

`func (o *ADConnector) HasUPDATEUSERJSON() bool`

HasUPDATEUSERJSON returns a boolean if a field has been set.

### GetSaveconnection

`func (o *ADConnector) GetSaveconnection() string`

GetSaveconnection returns the Saveconnection field if non-nil, zero value otherwise.

### GetSaveconnectionOk

`func (o *ADConnector) GetSaveconnectionOk() (*string, bool)`

GetSaveconnectionOk returns a tuple with the Saveconnection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaveconnection

`func (o *ADConnector) SetSaveconnection(v string)`

SetSaveconnection sets Saveconnection field to given value.

### HasSaveconnection

`func (o *ADConnector) HasSaveconnection() bool`

HasSaveconnection returns a boolean if a field has been set.

### GetSystemname

`func (o *ADConnector) GetSystemname() string`

GetSystemname returns the Systemname field if non-nil, zero value otherwise.

### GetSystemnameOk

`func (o *ADConnector) GetSystemnameOk() (*string, bool)`

GetSystemnameOk returns a tuple with the Systemname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemname

`func (o *ADConnector) SetSystemname(v string)`

SetSystemname sets Systemname field to given value.

### HasSystemname

`func (o *ADConnector) HasSystemname() bool`

HasSystemname returns a boolean if a field has been set.

### GetSETRANDOMPASSWORD

`func (o *ADConnector) GetSETRANDOMPASSWORD() string`

GetSETRANDOMPASSWORD returns the SETRANDOMPASSWORD field if non-nil, zero value otherwise.

### GetSETRANDOMPASSWORDOk

`func (o *ADConnector) GetSETRANDOMPASSWORDOk() (*string, bool)`

GetSETRANDOMPASSWORDOk returns a tuple with the SETRANDOMPASSWORD field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSETRANDOMPASSWORD

`func (o *ADConnector) SetSETRANDOMPASSWORD(v string)`

SetSETRANDOMPASSWORD sets SETRANDOMPASSWORD field to given value.

### HasSETRANDOMPASSWORD

`func (o *ADConnector) HasSETRANDOMPASSWORD() bool`

HasSETRANDOMPASSWORD returns a boolean if a field has been set.

### GetPASSWORD_MIN_LENGTH

`func (o *ADConnector) GetPASSWORD_MIN_LENGTH() string`

GetPASSWORD_MIN_LENGTH returns the PASSWORD_MIN_LENGTH field if non-nil, zero value otherwise.

### GetPASSWORD_MIN_LENGTHOk

`func (o *ADConnector) GetPASSWORD_MIN_LENGTHOk() (*string, bool)`

GetPASSWORD_MIN_LENGTHOk returns a tuple with the PASSWORD_MIN_LENGTH field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPASSWORD_MIN_LENGTH

`func (o *ADConnector) SetPASSWORD_MIN_LENGTH(v string)`

SetPASSWORD_MIN_LENGTH sets PASSWORD_MIN_LENGTH field to given value.

### HasPASSWORD_MIN_LENGTH

`func (o *ADConnector) HasPASSWORD_MIN_LENGTH() bool`

HasPASSWORD_MIN_LENGTH returns a boolean if a field has been set.

### GetPASSWORD_MAX_LENGTH

`func (o *ADConnector) GetPASSWORD_MAX_LENGTH() string`

GetPASSWORD_MAX_LENGTH returns the PASSWORD_MAX_LENGTH field if non-nil, zero value otherwise.

### GetPASSWORD_MAX_LENGTHOk

`func (o *ADConnector) GetPASSWORD_MAX_LENGTHOk() (*string, bool)`

GetPASSWORD_MAX_LENGTHOk returns a tuple with the PASSWORD_MAX_LENGTH field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPASSWORD_MAX_LENGTH

`func (o *ADConnector) SetPASSWORD_MAX_LENGTH(v string)`

SetPASSWORD_MAX_LENGTH sets PASSWORD_MAX_LENGTH field to given value.

### HasPASSWORD_MAX_LENGTH

`func (o *ADConnector) HasPASSWORD_MAX_LENGTH() bool`

HasPASSWORD_MAX_LENGTH returns a boolean if a field has been set.

### GetPASSWORD_NOOFCAPSALPHA

`func (o *ADConnector) GetPASSWORD_NOOFCAPSALPHA() string`

GetPASSWORD_NOOFCAPSALPHA returns the PASSWORD_NOOFCAPSALPHA field if non-nil, zero value otherwise.

### GetPASSWORD_NOOFCAPSALPHAOk

`func (o *ADConnector) GetPASSWORD_NOOFCAPSALPHAOk() (*string, bool)`

GetPASSWORD_NOOFCAPSALPHAOk returns a tuple with the PASSWORD_NOOFCAPSALPHA field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPASSWORD_NOOFCAPSALPHA

`func (o *ADConnector) SetPASSWORD_NOOFCAPSALPHA(v string)`

SetPASSWORD_NOOFCAPSALPHA sets PASSWORD_NOOFCAPSALPHA field to given value.

### HasPASSWORD_NOOFCAPSALPHA

`func (o *ADConnector) HasPASSWORD_NOOFCAPSALPHA() bool`

HasPASSWORD_NOOFCAPSALPHA returns a boolean if a field has been set.

### GetPASSWORD_NOOFSPLCHARS

`func (o *ADConnector) GetPASSWORD_NOOFSPLCHARS() string`

GetPASSWORD_NOOFSPLCHARS returns the PASSWORD_NOOFSPLCHARS field if non-nil, zero value otherwise.

### GetPASSWORD_NOOFSPLCHARSOk

`func (o *ADConnector) GetPASSWORD_NOOFSPLCHARSOk() (*string, bool)`

GetPASSWORD_NOOFSPLCHARSOk returns a tuple with the PASSWORD_NOOFSPLCHARS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPASSWORD_NOOFSPLCHARS

`func (o *ADConnector) SetPASSWORD_NOOFSPLCHARS(v string)`

SetPASSWORD_NOOFSPLCHARS sets PASSWORD_NOOFSPLCHARS field to given value.

### HasPASSWORD_NOOFSPLCHARS

`func (o *ADConnector) HasPASSWORD_NOOFSPLCHARS() bool`

HasPASSWORD_NOOFSPLCHARS returns a boolean if a field has been set.

### GetPASSWORD_NOOFDIGITS

`func (o *ADConnector) GetPASSWORD_NOOFDIGITS() string`

GetPASSWORD_NOOFDIGITS returns the PASSWORD_NOOFDIGITS field if non-nil, zero value otherwise.

### GetPASSWORD_NOOFDIGITSOk

`func (o *ADConnector) GetPASSWORD_NOOFDIGITSOk() (*string, bool)`

GetPASSWORD_NOOFDIGITSOk returns a tuple with the PASSWORD_NOOFDIGITS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPASSWORD_NOOFDIGITS

`func (o *ADConnector) SetPASSWORD_NOOFDIGITS(v string)`

SetPASSWORD_NOOFDIGITS sets PASSWORD_NOOFDIGITS field to given value.

### HasPASSWORD_NOOFDIGITS

`func (o *ADConnector) HasPASSWORD_NOOFDIGITS() bool`

HasPASSWORD_NOOFDIGITS returns a boolean if a field has been set.

### GetGroupImportMapping

`func (o *ADConnector) GetGroupImportMapping() string`

GetGroupImportMapping returns the GroupImportMapping field if non-nil, zero value otherwise.

### GetGroupImportMappingOk

`func (o *ADConnector) GetGroupImportMappingOk() (*string, bool)`

GetGroupImportMappingOk returns a tuple with the GroupImportMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupImportMapping

`func (o *ADConnector) SetGroupImportMapping(v string)`

SetGroupImportMapping sets GroupImportMapping field to given value.

### HasGroupImportMapping

`func (o *ADConnector) HasGroupImportMapping() bool`

HasGroupImportMapping returns a boolean if a field has been set.

### GetUNLOCKACCOUNTJSON

`func (o *ADConnector) GetUNLOCKACCOUNTJSON() string`

GetUNLOCKACCOUNTJSON returns the UNLOCKACCOUNTJSON field if non-nil, zero value otherwise.

### GetUNLOCKACCOUNTJSONOk

`func (o *ADConnector) GetUNLOCKACCOUNTJSONOk() (*string, bool)`

GetUNLOCKACCOUNTJSONOk returns a tuple with the UNLOCKACCOUNTJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUNLOCKACCOUNTJSON

`func (o *ADConnector) SetUNLOCKACCOUNTJSON(v string)`

SetUNLOCKACCOUNTJSON sets UNLOCKACCOUNTJSON field to given value.

### HasUNLOCKACCOUNTJSON

`func (o *ADConnector) HasUNLOCKACCOUNTJSON() bool`

HasUNLOCKACCOUNTJSON returns a boolean if a field has been set.

### GetSTATUSKEYJSON

`func (o *ADConnector) GetSTATUSKEYJSON() string`

GetSTATUSKEYJSON returns the STATUSKEYJSON field if non-nil, zero value otherwise.

### GetSTATUSKEYJSONOk

`func (o *ADConnector) GetSTATUSKEYJSONOk() (*string, bool)`

GetSTATUSKEYJSONOk returns a tuple with the STATUSKEYJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSTATUSKEYJSON

`func (o *ADConnector) SetSTATUSKEYJSON(v string)`

SetSTATUSKEYJSON sets STATUSKEYJSON field to given value.

### HasSTATUSKEYJSON

`func (o *ADConnector) HasSTATUSKEYJSON() bool`

HasSTATUSKEYJSON returns a boolean if a field has been set.

### GetENABLEGROUPMANAGEMENT

`func (o *ADConnector) GetENABLEGROUPMANAGEMENT() string`

GetENABLEGROUPMANAGEMENT returns the ENABLEGROUPMANAGEMENT field if non-nil, zero value otherwise.

### GetENABLEGROUPMANAGEMENTOk

`func (o *ADConnector) GetENABLEGROUPMANAGEMENTOk() (*string, bool)`

GetENABLEGROUPMANAGEMENTOk returns a tuple with the ENABLEGROUPMANAGEMENT field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetENABLEGROUPMANAGEMENT

`func (o *ADConnector) SetENABLEGROUPMANAGEMENT(v string)`

SetENABLEGROUPMANAGEMENT sets ENABLEGROUPMANAGEMENT field to given value.

### HasENABLEGROUPMANAGEMENT

`func (o *ADConnector) HasENABLEGROUPMANAGEMENT() bool`

HasENABLEGROUPMANAGEMENT returns a boolean if a field has been set.

### GetDISABLEACCOUNTJSON

`func (o *ADConnector) GetDISABLEACCOUNTJSON() string`

GetDISABLEACCOUNTJSON returns the DISABLEACCOUNTJSON field if non-nil, zero value otherwise.

### GetDISABLEACCOUNTJSONOk

`func (o *ADConnector) GetDISABLEACCOUNTJSONOk() (*string, bool)`

GetDISABLEACCOUNTJSONOk returns a tuple with the DISABLEACCOUNTJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDISABLEACCOUNTJSON

`func (o *ADConnector) SetDISABLEACCOUNTJSON(v string)`

SetDISABLEACCOUNTJSON sets DISABLEACCOUNTJSON field to given value.

### HasDISABLEACCOUNTJSON

`func (o *ADConnector) HasDISABLEACCOUNTJSON() bool`

HasDISABLEACCOUNTJSON returns a boolean if a field has been set.

### GetMODIFYUSERDATAJSON

`func (o *ADConnector) GetMODIFYUSERDATAJSON() string`

GetMODIFYUSERDATAJSON returns the MODIFYUSERDATAJSON field if non-nil, zero value otherwise.

### GetMODIFYUSERDATAJSONOk

`func (o *ADConnector) GetMODIFYUSERDATAJSONOk() (*string, bool)`

GetMODIFYUSERDATAJSONOk returns a tuple with the MODIFYUSERDATAJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMODIFYUSERDATAJSON

`func (o *ADConnector) SetMODIFYUSERDATAJSON(v string)`

SetMODIFYUSERDATAJSON sets MODIFYUSERDATAJSON field to given value.

### HasMODIFYUSERDATAJSON

`func (o *ADConnector) HasMODIFYUSERDATAJSON() bool`

HasMODIFYUSERDATAJSON returns a boolean if a field has been set.

### GetORG_BASE

`func (o *ADConnector) GetORG_BASE() string`

GetORG_BASE returns the ORG_BASE field if non-nil, zero value otherwise.

### GetORG_BASEOk

`func (o *ADConnector) GetORG_BASEOk() (*string, bool)`

GetORG_BASEOk returns a tuple with the ORG_BASE field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetORG_BASE

`func (o *ADConnector) SetORG_BASE(v string)`

SetORG_BASE sets ORG_BASE field to given value.

### HasORG_BASE

`func (o *ADConnector) HasORG_BASE() bool`

HasORG_BASE returns a boolean if a field has been set.

### GetORGANIZATION_ATTRIBUTE

`func (o *ADConnector) GetORGANIZATION_ATTRIBUTE() string`

GetORGANIZATION_ATTRIBUTE returns the ORGANIZATION_ATTRIBUTE field if non-nil, zero value otherwise.

### GetORGANIZATION_ATTRIBUTEOk

`func (o *ADConnector) GetORGANIZATION_ATTRIBUTEOk() (*string, bool)`

GetORGANIZATION_ATTRIBUTEOk returns a tuple with the ORGANIZATION_ATTRIBUTE field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetORGANIZATION_ATTRIBUTE

`func (o *ADConnector) SetORGANIZATION_ATTRIBUTE(v string)`

SetORGANIZATION_ATTRIBUTE sets ORGANIZATION_ATTRIBUTE field to given value.

### HasORGANIZATION_ATTRIBUTE

`func (o *ADConnector) HasORGANIZATION_ATTRIBUTE() bool`

HasORGANIZATION_ATTRIBUTE returns a boolean if a field has been set.

### GetORGIMPORTJSON

`func (o *ADConnector) GetORGIMPORTJSON() string`

GetORGIMPORTJSON returns the ORGIMPORTJSON field if non-nil, zero value otherwise.

### GetORGIMPORTJSONOk

`func (o *ADConnector) GetORGIMPORTJSONOk() (*string, bool)`

GetORGIMPORTJSONOk returns a tuple with the ORGIMPORTJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetORGIMPORTJSON

`func (o *ADConnector) SetORGIMPORTJSON(v string)`

SetORGIMPORTJSON sets ORGIMPORTJSON field to given value.

### HasORGIMPORTJSON

`func (o *ADConnector) HasORGIMPORTJSON() bool`

HasORGIMPORTJSON returns a boolean if a field has been set.

### GetCREATEORGJSON

`func (o *ADConnector) GetCREATEORGJSON() string`

GetCREATEORGJSON returns the CREATEORGJSON field if non-nil, zero value otherwise.

### GetCREATEORGJSONOk

`func (o *ADConnector) GetCREATEORGJSONOk() (*string, bool)`

GetCREATEORGJSONOk returns a tuple with the CREATEORGJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCREATEORGJSON

`func (o *ADConnector) SetCREATEORGJSON(v string)`

SetCREATEORGJSON sets CREATEORGJSON field to given value.

### HasCREATEORGJSON

`func (o *ADConnector) HasCREATEORGJSON() bool`

HasCREATEORGJSON returns a boolean if a field has been set.

### GetUPDATEORGJSON

`func (o *ADConnector) GetUPDATEORGJSON() string`

GetUPDATEORGJSON returns the UPDATEORGJSON field if non-nil, zero value otherwise.

### GetUPDATEORGJSONOk

`func (o *ADConnector) GetUPDATEORGJSONOk() (*string, bool)`

GetUPDATEORGJSONOk returns a tuple with the UPDATEORGJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUPDATEORGJSON

`func (o *ADConnector) SetUPDATEORGJSON(v string)`

SetUPDATEORGJSON sets UPDATEORGJSON field to given value.

### HasUPDATEORGJSON

`func (o *ADConnector) HasUPDATEORGJSON() bool`

HasUPDATEORGJSON returns a boolean if a field has been set.

### GetConfigJSON

`func (o *ADConnector) GetConfigJSON() string`

GetConfigJSON returns the ConfigJSON field if non-nil, zero value otherwise.

### GetConfigJSONOk

`func (o *ADConnector) GetConfigJSONOk() (*string, bool)`

GetConfigJSONOk returns a tuple with the ConfigJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigJSON

`func (o *ADConnector) SetConfigJSON(v string)`

SetConfigJSON sets ConfigJSON field to given value.

### HasConfigJSON

`func (o *ADConnector) HasConfigJSON() bool`

HasConfigJSON returns a boolean if a field has been set.

### GetLAST_IMPORT_TIME

`func (o *ADConnector) GetLAST_IMPORT_TIME() string`

GetLAST_IMPORT_TIME returns the LAST_IMPORT_TIME field if non-nil, zero value otherwise.

### GetLAST_IMPORT_TIMEOk

`func (o *ADConnector) GetLAST_IMPORT_TIMEOk() (*string, bool)`

GetLAST_IMPORT_TIMEOk returns a tuple with the LAST_IMPORT_TIME field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLAST_IMPORT_TIME

`func (o *ADConnector) SetLAST_IMPORT_TIME(v string)`

SetLAST_IMPORT_TIME sets LAST_IMPORT_TIME field to given value.

### HasLAST_IMPORT_TIME

`func (o *ADConnector) HasLAST_IMPORT_TIME() bool`

HasLAST_IMPORT_TIME returns a boolean if a field has been set.

### GetPAM_CONFIG

`func (o *ADConnector) GetPAM_CONFIG() string`

GetPAM_CONFIG returns the PAM_CONFIG field if non-nil, zero value otherwise.

### GetPAM_CONFIGOk

`func (o *ADConnector) GetPAM_CONFIGOk() (*string, bool)`

GetPAM_CONFIGOk returns a tuple with the PAM_CONFIG field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPAM_CONFIG

`func (o *ADConnector) SetPAM_CONFIG(v string)`

SetPAM_CONFIG sets PAM_CONFIG field to given value.

### HasPAM_CONFIG

`func (o *ADConnector) HasPAM_CONFIG() bool`

HasPAM_CONFIG returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


