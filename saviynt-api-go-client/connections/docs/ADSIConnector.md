# ADSIConnector

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**URL** | **string** | Primary/root domain URL list (comma Separated) | 
**USERNAME** | **string** | Service account username | 
**ENABLEGROUPMANAGEMENT** | Pointer to **string** | True implies group management capability limited to AD only. | [optional] 
**IMPORTDATACOOKIES** | Pointer to **string** |  | [optional] 
**MOVEACCOUNTJSON** | Pointer to **string** | JSON to specify movement action logic explicitely. If you have defined &#39;moveObjectToOU&#39; parameter in your update/enable/disable actions implicitely, then no need to define this explicitely here | [optional] 
**PasswdPolicyJSON** | Pointer to **string** | Optional Password policy JSON, in case you have no password policy defined/assigned in your Security System. | [optional] 
**CreateUpdateMappings** | Pointer to **string** | Mapping used during group creation to specify which attribute of a group maps to which attribute on Saviynt. | [optional] 
**PASSWORD** | **string** | Service account password | 
**CONNECTION_URL** | **string** | ADSI remote agent Connection URL | 
**PROVISIONING_URL** | Pointer to **string** | ADSI remote agent Provisioning URL | [optional] 
**FORESTLIST** | **string** | Forest List (Comma Separated) which we need to manage | 
**DEFAULT_USER_ROLE** | Pointer to **string** | Default SAV Role to be assigned to all the new users that gets imported via User Import | [optional] 
**UPDATEUSERJSON** | Pointer to **string** | Specify the attribute Value which will be used to Update existing User | [optional] 
**ENDPOINTS_FILTER** | Pointer to **string** | Provide the configuration to create Child Endpoints and import associated accounts and entitlements | [optional] 
**SEARCHFILTER** | Pointer to **string** | Account Search Filter to specify the starting point of the directory from where the accounts needs to be imported. You can have multiple BaseDNs here separated by ###. | [optional] 
**OBJECTFILTER** | Pointer to **string** | Object Filter is used to filter the objects that will be returned.This filter will be same for all domains. | [optional] 
**ACCOUNT_ATTRIBUTE** | Pointer to **string** | Map EIC and AD attributes for account import (AD attributes must be in lower case) | [optional] 
**STATUS_THRESHOLD_CONFIG** | Pointer to **string** | Account status and threshold related config | [optional] 
**ENTITLEMENT_ATTRIBUTE** | Pointer to **string** | Account attribute that contains group membership | [optional] 
**USER_ATTRIBUTE** | Pointer to **string** | Map EIC and AD attributes for user import (AD attributes must be in lower case) | [optional] 
**GroupSearchBaseDN** | Pointer to **string** | Group Search Filter to specify the starting point of the directory from where the groups needs to be imported. You can have multiple BaseDNs here separated by ###. | [optional] 
**CHECKFORUNIQUE** | Pointer to **string** | Evaluate the uniqueness of an attribute | [optional] 
**STATUSKEYJSON** | Pointer to **string** | JSON configuration to specify Users status | [optional] 
**GroupImportMapping** | Pointer to **string** | Map AD group attribute to EIC entitlement attribute for import | [optional] 
**ImportNestedMembership** | Pointer to **string** | Specify if you want the connector to import all indirect or nested membership of an account or a group during access import | [optional] 
**PAGE_SIZE** | Pointer to **string** | Page size defines the number of objects to be returned from each AD operation. | [optional] 
**ACCOUNTNAMERULE** | Pointer to **string** | Rule to generate account name. | [optional] 
**CREATEACCOUNTJSON** | Pointer to **string** | Specify the attributes values which will be used to Create the New Account. | [optional] 
**UPDATEACCOUNTJSON** | Pointer to **string** | Specify the attributes values which will be used to Update existing Account. | [optional] 
**ENABLEACCOUNTJSON** | Pointer to **string** | Specify the actions and attribute updates to be performed for enabling an account. | [optional] 
**DISABLEACCOUNTJSON** | Pointer to **string** | Specify the actions and attributes updates to be performed for disabling an account. | [optional] 
**REMOVEACCOUNTJSON** | Pointer to **string** | Specify the actions to be performed for deleting an account. | [optional] 
**ADDACCESSJSON** | Pointer to **string** | Configuration to ADD Access (cross domain/forest group membership) to an account. | [optional] 
**REMOVEACCESSJSON** | Pointer to **string** | Configuration to REMOVE Access (cross domain/forest group membership) to an account. | [optional] 
**RESETANDCHANGEPASSWRDJSON** | Pointer to **string** | Configuration to Reset and Change Password. | [optional] 
**CREATEGROUPJSON** | Pointer to **string** | Configuration to Create a Group | [optional] 
**UPDATEGROUPJSON** | Pointer to **string** | Configuration to Update a Group | [optional] 
**REMOVEGROUPJSON** | Pointer to **string** | Configuration to Delete a Group | [optional] 
**ADDACCESSENTITLEMENTJSON** | Pointer to **string** | Configuration to Add nested group hierarchy | [optional] 
**CUSTOMCONFIGJSON** | Pointer to **string** |  | [optional] 
**REMOVEACCESSENTITLEMENTJSON** | Pointer to **string** | Configuration to Remove nested group hierarchy | [optional] 
**CREATESERVICEACCOUNTJSON** | Pointer to **string** | Specify the Field Value which will be used to Create the New Service Account. | [optional] 
**UPDATESERVICEACCOUNTJSON** | Pointer to **string** | Specify the Field Value which will be used to update the existing Service Account. | [optional] 
**REMOVESERVICEACCOUNTJSON** | Pointer to **string** | Specify the actions to be performed while deleting a service account. | [optional] 
**PAM_CONFIG** | Pointer to **string** | JSON to specify Bootstrap Config. | [optional] 
**MODIFYUSERDATAJSON** | Pointer to **string** | Specify this parameter to transform the data during user import. | [optional] 

## Methods

### NewADSIConnector

`func NewADSIConnector(uRL string, uSERNAME string, pASSWORD string, cONNECTIONURL string, fORESTLIST string, ) *ADSIConnector`

NewADSIConnector instantiates a new ADSIConnector object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewADSIConnectorWithDefaults

`func NewADSIConnectorWithDefaults() *ADSIConnector`

NewADSIConnectorWithDefaults instantiates a new ADSIConnector object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetURL

`func (o *ADSIConnector) GetURL() string`

GetURL returns the URL field if non-nil, zero value otherwise.

### GetURLOk

`func (o *ADSIConnector) GetURLOk() (*string, bool)`

GetURLOk returns a tuple with the URL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetURL

`func (o *ADSIConnector) SetURL(v string)`

SetURL sets URL field to given value.


### GetUSERNAME

`func (o *ADSIConnector) GetUSERNAME() string`

GetUSERNAME returns the USERNAME field if non-nil, zero value otherwise.

### GetUSERNAMEOk

`func (o *ADSIConnector) GetUSERNAMEOk() (*string, bool)`

GetUSERNAMEOk returns a tuple with the USERNAME field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUSERNAME

`func (o *ADSIConnector) SetUSERNAME(v string)`

SetUSERNAME sets USERNAME field to given value.


### GetENABLEGROUPMANAGEMENT

`func (o *ADSIConnector) GetENABLEGROUPMANAGEMENT() string`

GetENABLEGROUPMANAGEMENT returns the ENABLEGROUPMANAGEMENT field if non-nil, zero value otherwise.

### GetENABLEGROUPMANAGEMENTOk

`func (o *ADSIConnector) GetENABLEGROUPMANAGEMENTOk() (*string, bool)`

GetENABLEGROUPMANAGEMENTOk returns a tuple with the ENABLEGROUPMANAGEMENT field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetENABLEGROUPMANAGEMENT

`func (o *ADSIConnector) SetENABLEGROUPMANAGEMENT(v string)`

SetENABLEGROUPMANAGEMENT sets ENABLEGROUPMANAGEMENT field to given value.

### HasENABLEGROUPMANAGEMENT

`func (o *ADSIConnector) HasENABLEGROUPMANAGEMENT() bool`

HasENABLEGROUPMANAGEMENT returns a boolean if a field has been set.

### GetIMPORTDATACOOKIES

`func (o *ADSIConnector) GetIMPORTDATACOOKIES() string`

GetIMPORTDATACOOKIES returns the IMPORTDATACOOKIES field if non-nil, zero value otherwise.

### GetIMPORTDATACOOKIESOk

`func (o *ADSIConnector) GetIMPORTDATACOOKIESOk() (*string, bool)`

GetIMPORTDATACOOKIESOk returns a tuple with the IMPORTDATACOOKIES field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIMPORTDATACOOKIES

`func (o *ADSIConnector) SetIMPORTDATACOOKIES(v string)`

SetIMPORTDATACOOKIES sets IMPORTDATACOOKIES field to given value.

### HasIMPORTDATACOOKIES

`func (o *ADSIConnector) HasIMPORTDATACOOKIES() bool`

HasIMPORTDATACOOKIES returns a boolean if a field has been set.

### GetMOVEACCOUNTJSON

`func (o *ADSIConnector) GetMOVEACCOUNTJSON() string`

GetMOVEACCOUNTJSON returns the MOVEACCOUNTJSON field if non-nil, zero value otherwise.

### GetMOVEACCOUNTJSONOk

`func (o *ADSIConnector) GetMOVEACCOUNTJSONOk() (*string, bool)`

GetMOVEACCOUNTJSONOk returns a tuple with the MOVEACCOUNTJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMOVEACCOUNTJSON

`func (o *ADSIConnector) SetMOVEACCOUNTJSON(v string)`

SetMOVEACCOUNTJSON sets MOVEACCOUNTJSON field to given value.

### HasMOVEACCOUNTJSON

`func (o *ADSIConnector) HasMOVEACCOUNTJSON() bool`

HasMOVEACCOUNTJSON returns a boolean if a field has been set.

### GetPasswdPolicyJSON

`func (o *ADSIConnector) GetPasswdPolicyJSON() string`

GetPasswdPolicyJSON returns the PasswdPolicyJSON field if non-nil, zero value otherwise.

### GetPasswdPolicyJSONOk

`func (o *ADSIConnector) GetPasswdPolicyJSONOk() (*string, bool)`

GetPasswdPolicyJSONOk returns a tuple with the PasswdPolicyJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswdPolicyJSON

`func (o *ADSIConnector) SetPasswdPolicyJSON(v string)`

SetPasswdPolicyJSON sets PasswdPolicyJSON field to given value.

### HasPasswdPolicyJSON

`func (o *ADSIConnector) HasPasswdPolicyJSON() bool`

HasPasswdPolicyJSON returns a boolean if a field has been set.

### GetCreateUpdateMappings

`func (o *ADSIConnector) GetCreateUpdateMappings() string`

GetCreateUpdateMappings returns the CreateUpdateMappings field if non-nil, zero value otherwise.

### GetCreateUpdateMappingsOk

`func (o *ADSIConnector) GetCreateUpdateMappingsOk() (*string, bool)`

GetCreateUpdateMappingsOk returns a tuple with the CreateUpdateMappings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateUpdateMappings

`func (o *ADSIConnector) SetCreateUpdateMappings(v string)`

SetCreateUpdateMappings sets CreateUpdateMappings field to given value.

### HasCreateUpdateMappings

`func (o *ADSIConnector) HasCreateUpdateMappings() bool`

HasCreateUpdateMappings returns a boolean if a field has been set.

### GetPASSWORD

`func (o *ADSIConnector) GetPASSWORD() string`

GetPASSWORD returns the PASSWORD field if non-nil, zero value otherwise.

### GetPASSWORDOk

`func (o *ADSIConnector) GetPASSWORDOk() (*string, bool)`

GetPASSWORDOk returns a tuple with the PASSWORD field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPASSWORD

`func (o *ADSIConnector) SetPASSWORD(v string)`

SetPASSWORD sets PASSWORD field to given value.


### GetCONNECTION_URL

`func (o *ADSIConnector) GetCONNECTION_URL() string`

GetCONNECTION_URL returns the CONNECTION_URL field if non-nil, zero value otherwise.

### GetCONNECTION_URLOk

`func (o *ADSIConnector) GetCONNECTION_URLOk() (*string, bool)`

GetCONNECTION_URLOk returns a tuple with the CONNECTION_URL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCONNECTION_URL

`func (o *ADSIConnector) SetCONNECTION_URL(v string)`

SetCONNECTION_URL sets CONNECTION_URL field to given value.


### GetPROVISIONING_URL

`func (o *ADSIConnector) GetPROVISIONING_URL() string`

GetPROVISIONING_URL returns the PROVISIONING_URL field if non-nil, zero value otherwise.

### GetPROVISIONING_URLOk

`func (o *ADSIConnector) GetPROVISIONING_URLOk() (*string, bool)`

GetPROVISIONING_URLOk returns a tuple with the PROVISIONING_URL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPROVISIONING_URL

`func (o *ADSIConnector) SetPROVISIONING_URL(v string)`

SetPROVISIONING_URL sets PROVISIONING_URL field to given value.

### HasPROVISIONING_URL

`func (o *ADSIConnector) HasPROVISIONING_URL() bool`

HasPROVISIONING_URL returns a boolean if a field has been set.

### GetFORESTLIST

`func (o *ADSIConnector) GetFORESTLIST() string`

GetFORESTLIST returns the FORESTLIST field if non-nil, zero value otherwise.

### GetFORESTLISTOk

`func (o *ADSIConnector) GetFORESTLISTOk() (*string, bool)`

GetFORESTLISTOk returns a tuple with the FORESTLIST field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFORESTLIST

`func (o *ADSIConnector) SetFORESTLIST(v string)`

SetFORESTLIST sets FORESTLIST field to given value.


### GetDEFAULT_USER_ROLE

`func (o *ADSIConnector) GetDEFAULT_USER_ROLE() string`

GetDEFAULT_USER_ROLE returns the DEFAULT_USER_ROLE field if non-nil, zero value otherwise.

### GetDEFAULT_USER_ROLEOk

`func (o *ADSIConnector) GetDEFAULT_USER_ROLEOk() (*string, bool)`

GetDEFAULT_USER_ROLEOk returns a tuple with the DEFAULT_USER_ROLE field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDEFAULT_USER_ROLE

`func (o *ADSIConnector) SetDEFAULT_USER_ROLE(v string)`

SetDEFAULT_USER_ROLE sets DEFAULT_USER_ROLE field to given value.

### HasDEFAULT_USER_ROLE

`func (o *ADSIConnector) HasDEFAULT_USER_ROLE() bool`

HasDEFAULT_USER_ROLE returns a boolean if a field has been set.

### GetUPDATEUSERJSON

`func (o *ADSIConnector) GetUPDATEUSERJSON() string`

GetUPDATEUSERJSON returns the UPDATEUSERJSON field if non-nil, zero value otherwise.

### GetUPDATEUSERJSONOk

`func (o *ADSIConnector) GetUPDATEUSERJSONOk() (*string, bool)`

GetUPDATEUSERJSONOk returns a tuple with the UPDATEUSERJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUPDATEUSERJSON

`func (o *ADSIConnector) SetUPDATEUSERJSON(v string)`

SetUPDATEUSERJSON sets UPDATEUSERJSON field to given value.

### HasUPDATEUSERJSON

`func (o *ADSIConnector) HasUPDATEUSERJSON() bool`

HasUPDATEUSERJSON returns a boolean if a field has been set.

### GetENDPOINTS_FILTER

`func (o *ADSIConnector) GetENDPOINTS_FILTER() string`

GetENDPOINTS_FILTER returns the ENDPOINTS_FILTER field if non-nil, zero value otherwise.

### GetENDPOINTS_FILTEROk

`func (o *ADSIConnector) GetENDPOINTS_FILTEROk() (*string, bool)`

GetENDPOINTS_FILTEROk returns a tuple with the ENDPOINTS_FILTER field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetENDPOINTS_FILTER

`func (o *ADSIConnector) SetENDPOINTS_FILTER(v string)`

SetENDPOINTS_FILTER sets ENDPOINTS_FILTER field to given value.

### HasENDPOINTS_FILTER

`func (o *ADSIConnector) HasENDPOINTS_FILTER() bool`

HasENDPOINTS_FILTER returns a boolean if a field has been set.

### GetSEARCHFILTER

`func (o *ADSIConnector) GetSEARCHFILTER() string`

GetSEARCHFILTER returns the SEARCHFILTER field if non-nil, zero value otherwise.

### GetSEARCHFILTEROk

`func (o *ADSIConnector) GetSEARCHFILTEROk() (*string, bool)`

GetSEARCHFILTEROk returns a tuple with the SEARCHFILTER field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEARCHFILTER

`func (o *ADSIConnector) SetSEARCHFILTER(v string)`

SetSEARCHFILTER sets SEARCHFILTER field to given value.

### HasSEARCHFILTER

`func (o *ADSIConnector) HasSEARCHFILTER() bool`

HasSEARCHFILTER returns a boolean if a field has been set.

### GetOBJECTFILTER

`func (o *ADSIConnector) GetOBJECTFILTER() string`

GetOBJECTFILTER returns the OBJECTFILTER field if non-nil, zero value otherwise.

### GetOBJECTFILTEROk

`func (o *ADSIConnector) GetOBJECTFILTEROk() (*string, bool)`

GetOBJECTFILTEROk returns a tuple with the OBJECTFILTER field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOBJECTFILTER

`func (o *ADSIConnector) SetOBJECTFILTER(v string)`

SetOBJECTFILTER sets OBJECTFILTER field to given value.

### HasOBJECTFILTER

`func (o *ADSIConnector) HasOBJECTFILTER() bool`

HasOBJECTFILTER returns a boolean if a field has been set.

### GetACCOUNT_ATTRIBUTE

`func (o *ADSIConnector) GetACCOUNT_ATTRIBUTE() string`

GetACCOUNT_ATTRIBUTE returns the ACCOUNT_ATTRIBUTE field if non-nil, zero value otherwise.

### GetACCOUNT_ATTRIBUTEOk

`func (o *ADSIConnector) GetACCOUNT_ATTRIBUTEOk() (*string, bool)`

GetACCOUNT_ATTRIBUTEOk returns a tuple with the ACCOUNT_ATTRIBUTE field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetACCOUNT_ATTRIBUTE

`func (o *ADSIConnector) SetACCOUNT_ATTRIBUTE(v string)`

SetACCOUNT_ATTRIBUTE sets ACCOUNT_ATTRIBUTE field to given value.

### HasACCOUNT_ATTRIBUTE

`func (o *ADSIConnector) HasACCOUNT_ATTRIBUTE() bool`

HasACCOUNT_ATTRIBUTE returns a boolean if a field has been set.

### GetSTATUS_THRESHOLD_CONFIG

`func (o *ADSIConnector) GetSTATUS_THRESHOLD_CONFIG() string`

GetSTATUS_THRESHOLD_CONFIG returns the STATUS_THRESHOLD_CONFIG field if non-nil, zero value otherwise.

### GetSTATUS_THRESHOLD_CONFIGOk

`func (o *ADSIConnector) GetSTATUS_THRESHOLD_CONFIGOk() (*string, bool)`

GetSTATUS_THRESHOLD_CONFIGOk returns a tuple with the STATUS_THRESHOLD_CONFIG field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSTATUS_THRESHOLD_CONFIG

`func (o *ADSIConnector) SetSTATUS_THRESHOLD_CONFIG(v string)`

SetSTATUS_THRESHOLD_CONFIG sets STATUS_THRESHOLD_CONFIG field to given value.

### HasSTATUS_THRESHOLD_CONFIG

`func (o *ADSIConnector) HasSTATUS_THRESHOLD_CONFIG() bool`

HasSTATUS_THRESHOLD_CONFIG returns a boolean if a field has been set.

### GetENTITLEMENT_ATTRIBUTE

`func (o *ADSIConnector) GetENTITLEMENT_ATTRIBUTE() string`

GetENTITLEMENT_ATTRIBUTE returns the ENTITLEMENT_ATTRIBUTE field if non-nil, zero value otherwise.

### GetENTITLEMENT_ATTRIBUTEOk

`func (o *ADSIConnector) GetENTITLEMENT_ATTRIBUTEOk() (*string, bool)`

GetENTITLEMENT_ATTRIBUTEOk returns a tuple with the ENTITLEMENT_ATTRIBUTE field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetENTITLEMENT_ATTRIBUTE

`func (o *ADSIConnector) SetENTITLEMENT_ATTRIBUTE(v string)`

SetENTITLEMENT_ATTRIBUTE sets ENTITLEMENT_ATTRIBUTE field to given value.

### HasENTITLEMENT_ATTRIBUTE

`func (o *ADSIConnector) HasENTITLEMENT_ATTRIBUTE() bool`

HasENTITLEMENT_ATTRIBUTE returns a boolean if a field has been set.

### GetUSER_ATTRIBUTE

`func (o *ADSIConnector) GetUSER_ATTRIBUTE() string`

GetUSER_ATTRIBUTE returns the USER_ATTRIBUTE field if non-nil, zero value otherwise.

### GetUSER_ATTRIBUTEOk

`func (o *ADSIConnector) GetUSER_ATTRIBUTEOk() (*string, bool)`

GetUSER_ATTRIBUTEOk returns a tuple with the USER_ATTRIBUTE field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUSER_ATTRIBUTE

`func (o *ADSIConnector) SetUSER_ATTRIBUTE(v string)`

SetUSER_ATTRIBUTE sets USER_ATTRIBUTE field to given value.

### HasUSER_ATTRIBUTE

`func (o *ADSIConnector) HasUSER_ATTRIBUTE() bool`

HasUSER_ATTRIBUTE returns a boolean if a field has been set.

### GetGroupSearchBaseDN

`func (o *ADSIConnector) GetGroupSearchBaseDN() string`

GetGroupSearchBaseDN returns the GroupSearchBaseDN field if non-nil, zero value otherwise.

### GetGroupSearchBaseDNOk

`func (o *ADSIConnector) GetGroupSearchBaseDNOk() (*string, bool)`

GetGroupSearchBaseDNOk returns a tuple with the GroupSearchBaseDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupSearchBaseDN

`func (o *ADSIConnector) SetGroupSearchBaseDN(v string)`

SetGroupSearchBaseDN sets GroupSearchBaseDN field to given value.

### HasGroupSearchBaseDN

`func (o *ADSIConnector) HasGroupSearchBaseDN() bool`

HasGroupSearchBaseDN returns a boolean if a field has been set.

### GetCHECKFORUNIQUE

`func (o *ADSIConnector) GetCHECKFORUNIQUE() string`

GetCHECKFORUNIQUE returns the CHECKFORUNIQUE field if non-nil, zero value otherwise.

### GetCHECKFORUNIQUEOk

`func (o *ADSIConnector) GetCHECKFORUNIQUEOk() (*string, bool)`

GetCHECKFORUNIQUEOk returns a tuple with the CHECKFORUNIQUE field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCHECKFORUNIQUE

`func (o *ADSIConnector) SetCHECKFORUNIQUE(v string)`

SetCHECKFORUNIQUE sets CHECKFORUNIQUE field to given value.

### HasCHECKFORUNIQUE

`func (o *ADSIConnector) HasCHECKFORUNIQUE() bool`

HasCHECKFORUNIQUE returns a boolean if a field has been set.

### GetSTATUSKEYJSON

`func (o *ADSIConnector) GetSTATUSKEYJSON() string`

GetSTATUSKEYJSON returns the STATUSKEYJSON field if non-nil, zero value otherwise.

### GetSTATUSKEYJSONOk

`func (o *ADSIConnector) GetSTATUSKEYJSONOk() (*string, bool)`

GetSTATUSKEYJSONOk returns a tuple with the STATUSKEYJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSTATUSKEYJSON

`func (o *ADSIConnector) SetSTATUSKEYJSON(v string)`

SetSTATUSKEYJSON sets STATUSKEYJSON field to given value.

### HasSTATUSKEYJSON

`func (o *ADSIConnector) HasSTATUSKEYJSON() bool`

HasSTATUSKEYJSON returns a boolean if a field has been set.

### GetGroupImportMapping

`func (o *ADSIConnector) GetGroupImportMapping() string`

GetGroupImportMapping returns the GroupImportMapping field if non-nil, zero value otherwise.

### GetGroupImportMappingOk

`func (o *ADSIConnector) GetGroupImportMappingOk() (*string, bool)`

GetGroupImportMappingOk returns a tuple with the GroupImportMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupImportMapping

`func (o *ADSIConnector) SetGroupImportMapping(v string)`

SetGroupImportMapping sets GroupImportMapping field to given value.

### HasGroupImportMapping

`func (o *ADSIConnector) HasGroupImportMapping() bool`

HasGroupImportMapping returns a boolean if a field has been set.

### GetImportNestedMembership

`func (o *ADSIConnector) GetImportNestedMembership() string`

GetImportNestedMembership returns the ImportNestedMembership field if non-nil, zero value otherwise.

### GetImportNestedMembershipOk

`func (o *ADSIConnector) GetImportNestedMembershipOk() (*string, bool)`

GetImportNestedMembershipOk returns a tuple with the ImportNestedMembership field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportNestedMembership

`func (o *ADSIConnector) SetImportNestedMembership(v string)`

SetImportNestedMembership sets ImportNestedMembership field to given value.

### HasImportNestedMembership

`func (o *ADSIConnector) HasImportNestedMembership() bool`

HasImportNestedMembership returns a boolean if a field has been set.

### GetPAGE_SIZE

`func (o *ADSIConnector) GetPAGE_SIZE() string`

GetPAGE_SIZE returns the PAGE_SIZE field if non-nil, zero value otherwise.

### GetPAGE_SIZEOk

`func (o *ADSIConnector) GetPAGE_SIZEOk() (*string, bool)`

GetPAGE_SIZEOk returns a tuple with the PAGE_SIZE field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPAGE_SIZE

`func (o *ADSIConnector) SetPAGE_SIZE(v string)`

SetPAGE_SIZE sets PAGE_SIZE field to given value.

### HasPAGE_SIZE

`func (o *ADSIConnector) HasPAGE_SIZE() bool`

HasPAGE_SIZE returns a boolean if a field has been set.

### GetACCOUNTNAMERULE

`func (o *ADSIConnector) GetACCOUNTNAMERULE() string`

GetACCOUNTNAMERULE returns the ACCOUNTNAMERULE field if non-nil, zero value otherwise.

### GetACCOUNTNAMERULEOk

`func (o *ADSIConnector) GetACCOUNTNAMERULEOk() (*string, bool)`

GetACCOUNTNAMERULEOk returns a tuple with the ACCOUNTNAMERULE field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetACCOUNTNAMERULE

`func (o *ADSIConnector) SetACCOUNTNAMERULE(v string)`

SetACCOUNTNAMERULE sets ACCOUNTNAMERULE field to given value.

### HasACCOUNTNAMERULE

`func (o *ADSIConnector) HasACCOUNTNAMERULE() bool`

HasACCOUNTNAMERULE returns a boolean if a field has been set.

### GetCREATEACCOUNTJSON

`func (o *ADSIConnector) GetCREATEACCOUNTJSON() string`

GetCREATEACCOUNTJSON returns the CREATEACCOUNTJSON field if non-nil, zero value otherwise.

### GetCREATEACCOUNTJSONOk

`func (o *ADSIConnector) GetCREATEACCOUNTJSONOk() (*string, bool)`

GetCREATEACCOUNTJSONOk returns a tuple with the CREATEACCOUNTJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCREATEACCOUNTJSON

`func (o *ADSIConnector) SetCREATEACCOUNTJSON(v string)`

SetCREATEACCOUNTJSON sets CREATEACCOUNTJSON field to given value.

### HasCREATEACCOUNTJSON

`func (o *ADSIConnector) HasCREATEACCOUNTJSON() bool`

HasCREATEACCOUNTJSON returns a boolean if a field has been set.

### GetUPDATEACCOUNTJSON

`func (o *ADSIConnector) GetUPDATEACCOUNTJSON() string`

GetUPDATEACCOUNTJSON returns the UPDATEACCOUNTJSON field if non-nil, zero value otherwise.

### GetUPDATEACCOUNTJSONOk

`func (o *ADSIConnector) GetUPDATEACCOUNTJSONOk() (*string, bool)`

GetUPDATEACCOUNTJSONOk returns a tuple with the UPDATEACCOUNTJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUPDATEACCOUNTJSON

`func (o *ADSIConnector) SetUPDATEACCOUNTJSON(v string)`

SetUPDATEACCOUNTJSON sets UPDATEACCOUNTJSON field to given value.

### HasUPDATEACCOUNTJSON

`func (o *ADSIConnector) HasUPDATEACCOUNTJSON() bool`

HasUPDATEACCOUNTJSON returns a boolean if a field has been set.

### GetENABLEACCOUNTJSON

`func (o *ADSIConnector) GetENABLEACCOUNTJSON() string`

GetENABLEACCOUNTJSON returns the ENABLEACCOUNTJSON field if non-nil, zero value otherwise.

### GetENABLEACCOUNTJSONOk

`func (o *ADSIConnector) GetENABLEACCOUNTJSONOk() (*string, bool)`

GetENABLEACCOUNTJSONOk returns a tuple with the ENABLEACCOUNTJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetENABLEACCOUNTJSON

`func (o *ADSIConnector) SetENABLEACCOUNTJSON(v string)`

SetENABLEACCOUNTJSON sets ENABLEACCOUNTJSON field to given value.

### HasENABLEACCOUNTJSON

`func (o *ADSIConnector) HasENABLEACCOUNTJSON() bool`

HasENABLEACCOUNTJSON returns a boolean if a field has been set.

### GetDISABLEACCOUNTJSON

`func (o *ADSIConnector) GetDISABLEACCOUNTJSON() string`

GetDISABLEACCOUNTJSON returns the DISABLEACCOUNTJSON field if non-nil, zero value otherwise.

### GetDISABLEACCOUNTJSONOk

`func (o *ADSIConnector) GetDISABLEACCOUNTJSONOk() (*string, bool)`

GetDISABLEACCOUNTJSONOk returns a tuple with the DISABLEACCOUNTJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDISABLEACCOUNTJSON

`func (o *ADSIConnector) SetDISABLEACCOUNTJSON(v string)`

SetDISABLEACCOUNTJSON sets DISABLEACCOUNTJSON field to given value.

### HasDISABLEACCOUNTJSON

`func (o *ADSIConnector) HasDISABLEACCOUNTJSON() bool`

HasDISABLEACCOUNTJSON returns a boolean if a field has been set.

### GetREMOVEACCOUNTJSON

`func (o *ADSIConnector) GetREMOVEACCOUNTJSON() string`

GetREMOVEACCOUNTJSON returns the REMOVEACCOUNTJSON field if non-nil, zero value otherwise.

### GetREMOVEACCOUNTJSONOk

`func (o *ADSIConnector) GetREMOVEACCOUNTJSONOk() (*string, bool)`

GetREMOVEACCOUNTJSONOk returns a tuple with the REMOVEACCOUNTJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetREMOVEACCOUNTJSON

`func (o *ADSIConnector) SetREMOVEACCOUNTJSON(v string)`

SetREMOVEACCOUNTJSON sets REMOVEACCOUNTJSON field to given value.

### HasREMOVEACCOUNTJSON

`func (o *ADSIConnector) HasREMOVEACCOUNTJSON() bool`

HasREMOVEACCOUNTJSON returns a boolean if a field has been set.

### GetADDACCESSJSON

`func (o *ADSIConnector) GetADDACCESSJSON() string`

GetADDACCESSJSON returns the ADDACCESSJSON field if non-nil, zero value otherwise.

### GetADDACCESSJSONOk

`func (o *ADSIConnector) GetADDACCESSJSONOk() (*string, bool)`

GetADDACCESSJSONOk returns a tuple with the ADDACCESSJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetADDACCESSJSON

`func (o *ADSIConnector) SetADDACCESSJSON(v string)`

SetADDACCESSJSON sets ADDACCESSJSON field to given value.

### HasADDACCESSJSON

`func (o *ADSIConnector) HasADDACCESSJSON() bool`

HasADDACCESSJSON returns a boolean if a field has been set.

### GetREMOVEACCESSJSON

`func (o *ADSIConnector) GetREMOVEACCESSJSON() string`

GetREMOVEACCESSJSON returns the REMOVEACCESSJSON field if non-nil, zero value otherwise.

### GetREMOVEACCESSJSONOk

`func (o *ADSIConnector) GetREMOVEACCESSJSONOk() (*string, bool)`

GetREMOVEACCESSJSONOk returns a tuple with the REMOVEACCESSJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetREMOVEACCESSJSON

`func (o *ADSIConnector) SetREMOVEACCESSJSON(v string)`

SetREMOVEACCESSJSON sets REMOVEACCESSJSON field to given value.

### HasREMOVEACCESSJSON

`func (o *ADSIConnector) HasREMOVEACCESSJSON() bool`

HasREMOVEACCESSJSON returns a boolean if a field has been set.

### GetRESETANDCHANGEPASSWRDJSON

`func (o *ADSIConnector) GetRESETANDCHANGEPASSWRDJSON() string`

GetRESETANDCHANGEPASSWRDJSON returns the RESETANDCHANGEPASSWRDJSON field if non-nil, zero value otherwise.

### GetRESETANDCHANGEPASSWRDJSONOk

`func (o *ADSIConnector) GetRESETANDCHANGEPASSWRDJSONOk() (*string, bool)`

GetRESETANDCHANGEPASSWRDJSONOk returns a tuple with the RESETANDCHANGEPASSWRDJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRESETANDCHANGEPASSWRDJSON

`func (o *ADSIConnector) SetRESETANDCHANGEPASSWRDJSON(v string)`

SetRESETANDCHANGEPASSWRDJSON sets RESETANDCHANGEPASSWRDJSON field to given value.

### HasRESETANDCHANGEPASSWRDJSON

`func (o *ADSIConnector) HasRESETANDCHANGEPASSWRDJSON() bool`

HasRESETANDCHANGEPASSWRDJSON returns a boolean if a field has been set.

### GetCREATEGROUPJSON

`func (o *ADSIConnector) GetCREATEGROUPJSON() string`

GetCREATEGROUPJSON returns the CREATEGROUPJSON field if non-nil, zero value otherwise.

### GetCREATEGROUPJSONOk

`func (o *ADSIConnector) GetCREATEGROUPJSONOk() (*string, bool)`

GetCREATEGROUPJSONOk returns a tuple with the CREATEGROUPJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCREATEGROUPJSON

`func (o *ADSIConnector) SetCREATEGROUPJSON(v string)`

SetCREATEGROUPJSON sets CREATEGROUPJSON field to given value.

### HasCREATEGROUPJSON

`func (o *ADSIConnector) HasCREATEGROUPJSON() bool`

HasCREATEGROUPJSON returns a boolean if a field has been set.

### GetUPDATEGROUPJSON

`func (o *ADSIConnector) GetUPDATEGROUPJSON() string`

GetUPDATEGROUPJSON returns the UPDATEGROUPJSON field if non-nil, zero value otherwise.

### GetUPDATEGROUPJSONOk

`func (o *ADSIConnector) GetUPDATEGROUPJSONOk() (*string, bool)`

GetUPDATEGROUPJSONOk returns a tuple with the UPDATEGROUPJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUPDATEGROUPJSON

`func (o *ADSIConnector) SetUPDATEGROUPJSON(v string)`

SetUPDATEGROUPJSON sets UPDATEGROUPJSON field to given value.

### HasUPDATEGROUPJSON

`func (o *ADSIConnector) HasUPDATEGROUPJSON() bool`

HasUPDATEGROUPJSON returns a boolean if a field has been set.

### GetREMOVEGROUPJSON

`func (o *ADSIConnector) GetREMOVEGROUPJSON() string`

GetREMOVEGROUPJSON returns the REMOVEGROUPJSON field if non-nil, zero value otherwise.

### GetREMOVEGROUPJSONOk

`func (o *ADSIConnector) GetREMOVEGROUPJSONOk() (*string, bool)`

GetREMOVEGROUPJSONOk returns a tuple with the REMOVEGROUPJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetREMOVEGROUPJSON

`func (o *ADSIConnector) SetREMOVEGROUPJSON(v string)`

SetREMOVEGROUPJSON sets REMOVEGROUPJSON field to given value.

### HasREMOVEGROUPJSON

`func (o *ADSIConnector) HasREMOVEGROUPJSON() bool`

HasREMOVEGROUPJSON returns a boolean if a field has been set.

### GetADDACCESSENTITLEMENTJSON

`func (o *ADSIConnector) GetADDACCESSENTITLEMENTJSON() string`

GetADDACCESSENTITLEMENTJSON returns the ADDACCESSENTITLEMENTJSON field if non-nil, zero value otherwise.

### GetADDACCESSENTITLEMENTJSONOk

`func (o *ADSIConnector) GetADDACCESSENTITLEMENTJSONOk() (*string, bool)`

GetADDACCESSENTITLEMENTJSONOk returns a tuple with the ADDACCESSENTITLEMENTJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetADDACCESSENTITLEMENTJSON

`func (o *ADSIConnector) SetADDACCESSENTITLEMENTJSON(v string)`

SetADDACCESSENTITLEMENTJSON sets ADDACCESSENTITLEMENTJSON field to given value.

### HasADDACCESSENTITLEMENTJSON

`func (o *ADSIConnector) HasADDACCESSENTITLEMENTJSON() bool`

HasADDACCESSENTITLEMENTJSON returns a boolean if a field has been set.

### GetCUSTOMCONFIGJSON

`func (o *ADSIConnector) GetCUSTOMCONFIGJSON() string`

GetCUSTOMCONFIGJSON returns the CUSTOMCONFIGJSON field if non-nil, zero value otherwise.

### GetCUSTOMCONFIGJSONOk

`func (o *ADSIConnector) GetCUSTOMCONFIGJSONOk() (*string, bool)`

GetCUSTOMCONFIGJSONOk returns a tuple with the CUSTOMCONFIGJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCUSTOMCONFIGJSON

`func (o *ADSIConnector) SetCUSTOMCONFIGJSON(v string)`

SetCUSTOMCONFIGJSON sets CUSTOMCONFIGJSON field to given value.

### HasCUSTOMCONFIGJSON

`func (o *ADSIConnector) HasCUSTOMCONFIGJSON() bool`

HasCUSTOMCONFIGJSON returns a boolean if a field has been set.

### GetREMOVEACCESSENTITLEMENTJSON

`func (o *ADSIConnector) GetREMOVEACCESSENTITLEMENTJSON() string`

GetREMOVEACCESSENTITLEMENTJSON returns the REMOVEACCESSENTITLEMENTJSON field if non-nil, zero value otherwise.

### GetREMOVEACCESSENTITLEMENTJSONOk

`func (o *ADSIConnector) GetREMOVEACCESSENTITLEMENTJSONOk() (*string, bool)`

GetREMOVEACCESSENTITLEMENTJSONOk returns a tuple with the REMOVEACCESSENTITLEMENTJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetREMOVEACCESSENTITLEMENTJSON

`func (o *ADSIConnector) SetREMOVEACCESSENTITLEMENTJSON(v string)`

SetREMOVEACCESSENTITLEMENTJSON sets REMOVEACCESSENTITLEMENTJSON field to given value.

### HasREMOVEACCESSENTITLEMENTJSON

`func (o *ADSIConnector) HasREMOVEACCESSENTITLEMENTJSON() bool`

HasREMOVEACCESSENTITLEMENTJSON returns a boolean if a field has been set.

### GetCREATESERVICEACCOUNTJSON

`func (o *ADSIConnector) GetCREATESERVICEACCOUNTJSON() string`

GetCREATESERVICEACCOUNTJSON returns the CREATESERVICEACCOUNTJSON field if non-nil, zero value otherwise.

### GetCREATESERVICEACCOUNTJSONOk

`func (o *ADSIConnector) GetCREATESERVICEACCOUNTJSONOk() (*string, bool)`

GetCREATESERVICEACCOUNTJSONOk returns a tuple with the CREATESERVICEACCOUNTJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCREATESERVICEACCOUNTJSON

`func (o *ADSIConnector) SetCREATESERVICEACCOUNTJSON(v string)`

SetCREATESERVICEACCOUNTJSON sets CREATESERVICEACCOUNTJSON field to given value.

### HasCREATESERVICEACCOUNTJSON

`func (o *ADSIConnector) HasCREATESERVICEACCOUNTJSON() bool`

HasCREATESERVICEACCOUNTJSON returns a boolean if a field has been set.

### GetUPDATESERVICEACCOUNTJSON

`func (o *ADSIConnector) GetUPDATESERVICEACCOUNTJSON() string`

GetUPDATESERVICEACCOUNTJSON returns the UPDATESERVICEACCOUNTJSON field if non-nil, zero value otherwise.

### GetUPDATESERVICEACCOUNTJSONOk

`func (o *ADSIConnector) GetUPDATESERVICEACCOUNTJSONOk() (*string, bool)`

GetUPDATESERVICEACCOUNTJSONOk returns a tuple with the UPDATESERVICEACCOUNTJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUPDATESERVICEACCOUNTJSON

`func (o *ADSIConnector) SetUPDATESERVICEACCOUNTJSON(v string)`

SetUPDATESERVICEACCOUNTJSON sets UPDATESERVICEACCOUNTJSON field to given value.

### HasUPDATESERVICEACCOUNTJSON

`func (o *ADSIConnector) HasUPDATESERVICEACCOUNTJSON() bool`

HasUPDATESERVICEACCOUNTJSON returns a boolean if a field has been set.

### GetREMOVESERVICEACCOUNTJSON

`func (o *ADSIConnector) GetREMOVESERVICEACCOUNTJSON() string`

GetREMOVESERVICEACCOUNTJSON returns the REMOVESERVICEACCOUNTJSON field if non-nil, zero value otherwise.

### GetREMOVESERVICEACCOUNTJSONOk

`func (o *ADSIConnector) GetREMOVESERVICEACCOUNTJSONOk() (*string, bool)`

GetREMOVESERVICEACCOUNTJSONOk returns a tuple with the REMOVESERVICEACCOUNTJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetREMOVESERVICEACCOUNTJSON

`func (o *ADSIConnector) SetREMOVESERVICEACCOUNTJSON(v string)`

SetREMOVESERVICEACCOUNTJSON sets REMOVESERVICEACCOUNTJSON field to given value.

### HasREMOVESERVICEACCOUNTJSON

`func (o *ADSIConnector) HasREMOVESERVICEACCOUNTJSON() bool`

HasREMOVESERVICEACCOUNTJSON returns a boolean if a field has been set.

### GetPAM_CONFIG

`func (o *ADSIConnector) GetPAM_CONFIG() string`

GetPAM_CONFIG returns the PAM_CONFIG field if non-nil, zero value otherwise.

### GetPAM_CONFIGOk

`func (o *ADSIConnector) GetPAM_CONFIGOk() (*string, bool)`

GetPAM_CONFIGOk returns a tuple with the PAM_CONFIG field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPAM_CONFIG

`func (o *ADSIConnector) SetPAM_CONFIG(v string)`

SetPAM_CONFIG sets PAM_CONFIG field to given value.

### HasPAM_CONFIG

`func (o *ADSIConnector) HasPAM_CONFIG() bool`

HasPAM_CONFIG returns a boolean if a field has been set.

### GetMODIFYUSERDATAJSON

`func (o *ADSIConnector) GetMODIFYUSERDATAJSON() string`

GetMODIFYUSERDATAJSON returns the MODIFYUSERDATAJSON field if non-nil, zero value otherwise.

### GetMODIFYUSERDATAJSONOk

`func (o *ADSIConnector) GetMODIFYUSERDATAJSONOk() (*string, bool)`

GetMODIFYUSERDATAJSONOk returns a tuple with the MODIFYUSERDATAJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMODIFYUSERDATAJSON

`func (o *ADSIConnector) SetMODIFYUSERDATAJSON(v string)`

SetMODIFYUSERDATAJSON sets MODIFYUSERDATAJSON field to given value.

### HasMODIFYUSERDATAJSON

`func (o *ADSIConnector) HasMODIFYUSERDATAJSON() bool`

HasMODIFYUSERDATAJSON returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


