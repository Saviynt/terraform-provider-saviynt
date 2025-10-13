# ADSIConnectionAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ImportNestedMembership** | Pointer to **string** |  | [optional] 
**PASSWDPOLICYJSON** | Pointer to **string** |  | [optional] 
**CREATEACCOUNTJSON** | Pointer to **string** |  | [optional] 
**ENDPOINTS_FILTER** | Pointer to **string** |  | [optional] 
**DISABLEACCOUNTJSON** | Pointer to **string** |  | [optional] 
**REMOVEACCESSENTITLEMENTJSON** | Pointer to **string** |  | [optional] 
**GroupSearchBaseDN** | Pointer to **string** |  | [optional] 
**ConnectionType** | Pointer to **string** |  | [optional] 
**STATUSKEYJSON** | Pointer to **string** |  | [optional] 
**DEFAULT_USER_ROLE** | Pointer to **string** |  | [optional] 
**FOREST_DETAILS** | Pointer to **string** |  | [optional] 
**USERNAME** | Pointer to **string** |  | [optional] 
**UPDATESERVICEACCOUNTJSON** | Pointer to **string** |  | [optional] 
**ADDACCESSJSON** | Pointer to **string** |  | [optional] 
**CREATESERVICEACCOUNTJSON** | Pointer to **string** |  | [optional] 
**ACCOUNTNAMERULE** | Pointer to **string** |  | [optional] 
**CONNECTION_URL** | Pointer to **string** |  | [optional] 
**IsTimeoutSupported** | Pointer to **bool** |  | [optional] 
**CreateUpdateMappings** | Pointer to **string** |  | [optional] 
**ACCOUNT_ATTRIBUTE** | Pointer to **string** |  | [optional] 
**PASSWORD** | Pointer to **string** |  | [optional] 
**PAM_CONFIG** | Pointer to **string** |  | [optional] 
**PAGE_SIZE** | Pointer to **string** |  | [optional] 
**SEARCHFILTER** | Pointer to **string** |  | [optional] 
**UPDATEGROUPJSON** | Pointer to **string** |  | [optional] 
**CREATEGROUPJSON** | Pointer to **string** |  | [optional] 
**ENTITLEMENT_ATTRIBUTE** | Pointer to **string** |  | [optional] 
**CHECKFORUNIQUE** | Pointer to **string** |  | [optional] 
**REMOVESERVICEACCOUNTJSON** | Pointer to **string** |  | [optional] 
**ConnectionTimeoutConfig** | Pointer to [**ConnectionTimeoutConfig**](ConnectionTimeoutConfig.md) |  | [optional] 
**UPDATEUSERJSON** | Pointer to **string** |  | [optional] 
**URL** | Pointer to **string** |  | [optional] 
**MOVEACCOUNTJSON** | Pointer to **string** |  | [optional] 
**CUSTOMCONFIGJSON** | Pointer to **string** |  | [optional] 
**STATUS_THRESHOLD_CONFIG** | Pointer to **string** |  | [optional] 
**GroupImportMapping** | Pointer to **string** |  | [optional] 
**PROVISIONING_URL** | Pointer to **string** |  | [optional] 
**REMOVEGROUPJSON** | Pointer to **string** |  | [optional] 
**REMOVEACCESSJSON** | Pointer to **string** |  | [optional] 
**IMPORTDATACOOKIES** | Pointer to **string** |  | [optional] 
**RESETANDCHANGEPASSWRDJSON** | Pointer to **string** |  | [optional] 
**USER_ATTRIBUTE** | Pointer to **string** |  | [optional] 
**ADDACCESSENTITLEMENTJSON** | Pointer to **string** |  | [optional] 
**MODIFYUSERDATAJSON** | Pointer to **string** |  | [optional] 
**IsTimeoutConfigValidated** | Pointer to **bool** |  | [optional] 
**ENABLEGROUPMANAGEMENT** | Pointer to **string** |  | [optional] 
**ENABLEACCOUNTJSON** | Pointer to **string** |  | [optional] 
**FORESTLIST** | Pointer to **string** |  | [optional] 
**OBJECTFILTER** | Pointer to **string** |  | [optional] 
**UPDATEACCOUNTJSON** | Pointer to **string** |  | [optional] 
**REMOVEACCOUNTJSON** | Pointer to **string** |  | [optional] 

## Methods

### NewADSIConnectionAttributes

`func NewADSIConnectionAttributes() *ADSIConnectionAttributes`

NewADSIConnectionAttributes instantiates a new ADSIConnectionAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewADSIConnectionAttributesWithDefaults

`func NewADSIConnectionAttributesWithDefaults() *ADSIConnectionAttributes`

NewADSIConnectionAttributesWithDefaults instantiates a new ADSIConnectionAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetImportNestedMembership

`func (o *ADSIConnectionAttributes) GetImportNestedMembership() string`

GetImportNestedMembership returns the ImportNestedMembership field if non-nil, zero value otherwise.

### GetImportNestedMembershipOk

`func (o *ADSIConnectionAttributes) GetImportNestedMembershipOk() (*string, bool)`

GetImportNestedMembershipOk returns a tuple with the ImportNestedMembership field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportNestedMembership

`func (o *ADSIConnectionAttributes) SetImportNestedMembership(v string)`

SetImportNestedMembership sets ImportNestedMembership field to given value.

### HasImportNestedMembership

`func (o *ADSIConnectionAttributes) HasImportNestedMembership() bool`

HasImportNestedMembership returns a boolean if a field has been set.

### GetPASSWDPOLICYJSON

`func (o *ADSIConnectionAttributes) GetPASSWDPOLICYJSON() string`

GetPASSWDPOLICYJSON returns the PASSWDPOLICYJSON field if non-nil, zero value otherwise.

### GetPASSWDPOLICYJSONOk

`func (o *ADSIConnectionAttributes) GetPASSWDPOLICYJSONOk() (*string, bool)`

GetPASSWDPOLICYJSONOk returns a tuple with the PASSWDPOLICYJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPASSWDPOLICYJSON

`func (o *ADSIConnectionAttributes) SetPASSWDPOLICYJSON(v string)`

SetPASSWDPOLICYJSON sets PASSWDPOLICYJSON field to given value.

### HasPASSWDPOLICYJSON

`func (o *ADSIConnectionAttributes) HasPASSWDPOLICYJSON() bool`

HasPASSWDPOLICYJSON returns a boolean if a field has been set.

### GetCREATEACCOUNTJSON

`func (o *ADSIConnectionAttributes) GetCREATEACCOUNTJSON() string`

GetCREATEACCOUNTJSON returns the CREATEACCOUNTJSON field if non-nil, zero value otherwise.

### GetCREATEACCOUNTJSONOk

`func (o *ADSIConnectionAttributes) GetCREATEACCOUNTJSONOk() (*string, bool)`

GetCREATEACCOUNTJSONOk returns a tuple with the CREATEACCOUNTJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCREATEACCOUNTJSON

`func (o *ADSIConnectionAttributes) SetCREATEACCOUNTJSON(v string)`

SetCREATEACCOUNTJSON sets CREATEACCOUNTJSON field to given value.

### HasCREATEACCOUNTJSON

`func (o *ADSIConnectionAttributes) HasCREATEACCOUNTJSON() bool`

HasCREATEACCOUNTJSON returns a boolean if a field has been set.

### GetENDPOINTS_FILTER

`func (o *ADSIConnectionAttributes) GetENDPOINTS_FILTER() string`

GetENDPOINTS_FILTER returns the ENDPOINTS_FILTER field if non-nil, zero value otherwise.

### GetENDPOINTS_FILTEROk

`func (o *ADSIConnectionAttributes) GetENDPOINTS_FILTEROk() (*string, bool)`

GetENDPOINTS_FILTEROk returns a tuple with the ENDPOINTS_FILTER field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetENDPOINTS_FILTER

`func (o *ADSIConnectionAttributes) SetENDPOINTS_FILTER(v string)`

SetENDPOINTS_FILTER sets ENDPOINTS_FILTER field to given value.

### HasENDPOINTS_FILTER

`func (o *ADSIConnectionAttributes) HasENDPOINTS_FILTER() bool`

HasENDPOINTS_FILTER returns a boolean if a field has been set.

### GetDISABLEACCOUNTJSON

`func (o *ADSIConnectionAttributes) GetDISABLEACCOUNTJSON() string`

GetDISABLEACCOUNTJSON returns the DISABLEACCOUNTJSON field if non-nil, zero value otherwise.

### GetDISABLEACCOUNTJSONOk

`func (o *ADSIConnectionAttributes) GetDISABLEACCOUNTJSONOk() (*string, bool)`

GetDISABLEACCOUNTJSONOk returns a tuple with the DISABLEACCOUNTJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDISABLEACCOUNTJSON

`func (o *ADSIConnectionAttributes) SetDISABLEACCOUNTJSON(v string)`

SetDISABLEACCOUNTJSON sets DISABLEACCOUNTJSON field to given value.

### HasDISABLEACCOUNTJSON

`func (o *ADSIConnectionAttributes) HasDISABLEACCOUNTJSON() bool`

HasDISABLEACCOUNTJSON returns a boolean if a field has been set.

### GetREMOVEACCESSENTITLEMENTJSON

`func (o *ADSIConnectionAttributes) GetREMOVEACCESSENTITLEMENTJSON() string`

GetREMOVEACCESSENTITLEMENTJSON returns the REMOVEACCESSENTITLEMENTJSON field if non-nil, zero value otherwise.

### GetREMOVEACCESSENTITLEMENTJSONOk

`func (o *ADSIConnectionAttributes) GetREMOVEACCESSENTITLEMENTJSONOk() (*string, bool)`

GetREMOVEACCESSENTITLEMENTJSONOk returns a tuple with the REMOVEACCESSENTITLEMENTJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetREMOVEACCESSENTITLEMENTJSON

`func (o *ADSIConnectionAttributes) SetREMOVEACCESSENTITLEMENTJSON(v string)`

SetREMOVEACCESSENTITLEMENTJSON sets REMOVEACCESSENTITLEMENTJSON field to given value.

### HasREMOVEACCESSENTITLEMENTJSON

`func (o *ADSIConnectionAttributes) HasREMOVEACCESSENTITLEMENTJSON() bool`

HasREMOVEACCESSENTITLEMENTJSON returns a boolean if a field has been set.

### GetGroupSearchBaseDN

`func (o *ADSIConnectionAttributes) GetGroupSearchBaseDN() string`

GetGroupSearchBaseDN returns the GroupSearchBaseDN field if non-nil, zero value otherwise.

### GetGroupSearchBaseDNOk

`func (o *ADSIConnectionAttributes) GetGroupSearchBaseDNOk() (*string, bool)`

GetGroupSearchBaseDNOk returns a tuple with the GroupSearchBaseDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupSearchBaseDN

`func (o *ADSIConnectionAttributes) SetGroupSearchBaseDN(v string)`

SetGroupSearchBaseDN sets GroupSearchBaseDN field to given value.

### HasGroupSearchBaseDN

`func (o *ADSIConnectionAttributes) HasGroupSearchBaseDN() bool`

HasGroupSearchBaseDN returns a boolean if a field has been set.

### GetConnectionType

`func (o *ADSIConnectionAttributes) GetConnectionType() string`

GetConnectionType returns the ConnectionType field if non-nil, zero value otherwise.

### GetConnectionTypeOk

`func (o *ADSIConnectionAttributes) GetConnectionTypeOk() (*string, bool)`

GetConnectionTypeOk returns a tuple with the ConnectionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionType

`func (o *ADSIConnectionAttributes) SetConnectionType(v string)`

SetConnectionType sets ConnectionType field to given value.

### HasConnectionType

`func (o *ADSIConnectionAttributes) HasConnectionType() bool`

HasConnectionType returns a boolean if a field has been set.

### GetSTATUSKEYJSON

`func (o *ADSIConnectionAttributes) GetSTATUSKEYJSON() string`

GetSTATUSKEYJSON returns the STATUSKEYJSON field if non-nil, zero value otherwise.

### GetSTATUSKEYJSONOk

`func (o *ADSIConnectionAttributes) GetSTATUSKEYJSONOk() (*string, bool)`

GetSTATUSKEYJSONOk returns a tuple with the STATUSKEYJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSTATUSKEYJSON

`func (o *ADSIConnectionAttributes) SetSTATUSKEYJSON(v string)`

SetSTATUSKEYJSON sets STATUSKEYJSON field to given value.

### HasSTATUSKEYJSON

`func (o *ADSIConnectionAttributes) HasSTATUSKEYJSON() bool`

HasSTATUSKEYJSON returns a boolean if a field has been set.

### GetDEFAULT_USER_ROLE

`func (o *ADSIConnectionAttributes) GetDEFAULT_USER_ROLE() string`

GetDEFAULT_USER_ROLE returns the DEFAULT_USER_ROLE field if non-nil, zero value otherwise.

### GetDEFAULT_USER_ROLEOk

`func (o *ADSIConnectionAttributes) GetDEFAULT_USER_ROLEOk() (*string, bool)`

GetDEFAULT_USER_ROLEOk returns a tuple with the DEFAULT_USER_ROLE field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDEFAULT_USER_ROLE

`func (o *ADSIConnectionAttributes) SetDEFAULT_USER_ROLE(v string)`

SetDEFAULT_USER_ROLE sets DEFAULT_USER_ROLE field to given value.

### HasDEFAULT_USER_ROLE

`func (o *ADSIConnectionAttributes) HasDEFAULT_USER_ROLE() bool`

HasDEFAULT_USER_ROLE returns a boolean if a field has been set.

### GetFOREST_DETAILS

`func (o *ADSIConnectionAttributes) GetFOREST_DETAILS() string`

GetFOREST_DETAILS returns the FOREST_DETAILS field if non-nil, zero value otherwise.

### GetFOREST_DETAILSOk

`func (o *ADSIConnectionAttributes) GetFOREST_DETAILSOk() (*string, bool)`

GetFOREST_DETAILSOk returns a tuple with the FOREST_DETAILS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFOREST_DETAILS

`func (o *ADSIConnectionAttributes) SetFOREST_DETAILS(v string)`

SetFOREST_DETAILS sets FOREST_DETAILS field to given value.

### HasFOREST_DETAILS

`func (o *ADSIConnectionAttributes) HasFOREST_DETAILS() bool`

HasFOREST_DETAILS returns a boolean if a field has been set.

### GetUSERNAME

`func (o *ADSIConnectionAttributes) GetUSERNAME() string`

GetUSERNAME returns the USERNAME field if non-nil, zero value otherwise.

### GetUSERNAMEOk

`func (o *ADSIConnectionAttributes) GetUSERNAMEOk() (*string, bool)`

GetUSERNAMEOk returns a tuple with the USERNAME field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUSERNAME

`func (o *ADSIConnectionAttributes) SetUSERNAME(v string)`

SetUSERNAME sets USERNAME field to given value.

### HasUSERNAME

`func (o *ADSIConnectionAttributes) HasUSERNAME() bool`

HasUSERNAME returns a boolean if a field has been set.

### GetUPDATESERVICEACCOUNTJSON

`func (o *ADSIConnectionAttributes) GetUPDATESERVICEACCOUNTJSON() string`

GetUPDATESERVICEACCOUNTJSON returns the UPDATESERVICEACCOUNTJSON field if non-nil, zero value otherwise.

### GetUPDATESERVICEACCOUNTJSONOk

`func (o *ADSIConnectionAttributes) GetUPDATESERVICEACCOUNTJSONOk() (*string, bool)`

GetUPDATESERVICEACCOUNTJSONOk returns a tuple with the UPDATESERVICEACCOUNTJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUPDATESERVICEACCOUNTJSON

`func (o *ADSIConnectionAttributes) SetUPDATESERVICEACCOUNTJSON(v string)`

SetUPDATESERVICEACCOUNTJSON sets UPDATESERVICEACCOUNTJSON field to given value.

### HasUPDATESERVICEACCOUNTJSON

`func (o *ADSIConnectionAttributes) HasUPDATESERVICEACCOUNTJSON() bool`

HasUPDATESERVICEACCOUNTJSON returns a boolean if a field has been set.

### GetADDACCESSJSON

`func (o *ADSIConnectionAttributes) GetADDACCESSJSON() string`

GetADDACCESSJSON returns the ADDACCESSJSON field if non-nil, zero value otherwise.

### GetADDACCESSJSONOk

`func (o *ADSIConnectionAttributes) GetADDACCESSJSONOk() (*string, bool)`

GetADDACCESSJSONOk returns a tuple with the ADDACCESSJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetADDACCESSJSON

`func (o *ADSIConnectionAttributes) SetADDACCESSJSON(v string)`

SetADDACCESSJSON sets ADDACCESSJSON field to given value.

### HasADDACCESSJSON

`func (o *ADSIConnectionAttributes) HasADDACCESSJSON() bool`

HasADDACCESSJSON returns a boolean if a field has been set.

### GetCREATESERVICEACCOUNTJSON

`func (o *ADSIConnectionAttributes) GetCREATESERVICEACCOUNTJSON() string`

GetCREATESERVICEACCOUNTJSON returns the CREATESERVICEACCOUNTJSON field if non-nil, zero value otherwise.

### GetCREATESERVICEACCOUNTJSONOk

`func (o *ADSIConnectionAttributes) GetCREATESERVICEACCOUNTJSONOk() (*string, bool)`

GetCREATESERVICEACCOUNTJSONOk returns a tuple with the CREATESERVICEACCOUNTJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCREATESERVICEACCOUNTJSON

`func (o *ADSIConnectionAttributes) SetCREATESERVICEACCOUNTJSON(v string)`

SetCREATESERVICEACCOUNTJSON sets CREATESERVICEACCOUNTJSON field to given value.

### HasCREATESERVICEACCOUNTJSON

`func (o *ADSIConnectionAttributes) HasCREATESERVICEACCOUNTJSON() bool`

HasCREATESERVICEACCOUNTJSON returns a boolean if a field has been set.

### GetACCOUNTNAMERULE

`func (o *ADSIConnectionAttributes) GetACCOUNTNAMERULE() string`

GetACCOUNTNAMERULE returns the ACCOUNTNAMERULE field if non-nil, zero value otherwise.

### GetACCOUNTNAMERULEOk

`func (o *ADSIConnectionAttributes) GetACCOUNTNAMERULEOk() (*string, bool)`

GetACCOUNTNAMERULEOk returns a tuple with the ACCOUNTNAMERULE field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetACCOUNTNAMERULE

`func (o *ADSIConnectionAttributes) SetACCOUNTNAMERULE(v string)`

SetACCOUNTNAMERULE sets ACCOUNTNAMERULE field to given value.

### HasACCOUNTNAMERULE

`func (o *ADSIConnectionAttributes) HasACCOUNTNAMERULE() bool`

HasACCOUNTNAMERULE returns a boolean if a field has been set.

### GetCONNECTION_URL

`func (o *ADSIConnectionAttributes) GetCONNECTION_URL() string`

GetCONNECTION_URL returns the CONNECTION_URL field if non-nil, zero value otherwise.

### GetCONNECTION_URLOk

`func (o *ADSIConnectionAttributes) GetCONNECTION_URLOk() (*string, bool)`

GetCONNECTION_URLOk returns a tuple with the CONNECTION_URL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCONNECTION_URL

`func (o *ADSIConnectionAttributes) SetCONNECTION_URL(v string)`

SetCONNECTION_URL sets CONNECTION_URL field to given value.

### HasCONNECTION_URL

`func (o *ADSIConnectionAttributes) HasCONNECTION_URL() bool`

HasCONNECTION_URL returns a boolean if a field has been set.

### GetIsTimeoutSupported

`func (o *ADSIConnectionAttributes) GetIsTimeoutSupported() bool`

GetIsTimeoutSupported returns the IsTimeoutSupported field if non-nil, zero value otherwise.

### GetIsTimeoutSupportedOk

`func (o *ADSIConnectionAttributes) GetIsTimeoutSupportedOk() (*bool, bool)`

GetIsTimeoutSupportedOk returns a tuple with the IsTimeoutSupported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsTimeoutSupported

`func (o *ADSIConnectionAttributes) SetIsTimeoutSupported(v bool)`

SetIsTimeoutSupported sets IsTimeoutSupported field to given value.

### HasIsTimeoutSupported

`func (o *ADSIConnectionAttributes) HasIsTimeoutSupported() bool`

HasIsTimeoutSupported returns a boolean if a field has been set.

### GetCreateUpdateMappings

`func (o *ADSIConnectionAttributes) GetCreateUpdateMappings() string`

GetCreateUpdateMappings returns the CreateUpdateMappings field if non-nil, zero value otherwise.

### GetCreateUpdateMappingsOk

`func (o *ADSIConnectionAttributes) GetCreateUpdateMappingsOk() (*string, bool)`

GetCreateUpdateMappingsOk returns a tuple with the CreateUpdateMappings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateUpdateMappings

`func (o *ADSIConnectionAttributes) SetCreateUpdateMappings(v string)`

SetCreateUpdateMappings sets CreateUpdateMappings field to given value.

### HasCreateUpdateMappings

`func (o *ADSIConnectionAttributes) HasCreateUpdateMappings() bool`

HasCreateUpdateMappings returns a boolean if a field has been set.

### GetACCOUNT_ATTRIBUTE

`func (o *ADSIConnectionAttributes) GetACCOUNT_ATTRIBUTE() string`

GetACCOUNT_ATTRIBUTE returns the ACCOUNT_ATTRIBUTE field if non-nil, zero value otherwise.

### GetACCOUNT_ATTRIBUTEOk

`func (o *ADSIConnectionAttributes) GetACCOUNT_ATTRIBUTEOk() (*string, bool)`

GetACCOUNT_ATTRIBUTEOk returns a tuple with the ACCOUNT_ATTRIBUTE field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetACCOUNT_ATTRIBUTE

`func (o *ADSIConnectionAttributes) SetACCOUNT_ATTRIBUTE(v string)`

SetACCOUNT_ATTRIBUTE sets ACCOUNT_ATTRIBUTE field to given value.

### HasACCOUNT_ATTRIBUTE

`func (o *ADSIConnectionAttributes) HasACCOUNT_ATTRIBUTE() bool`

HasACCOUNT_ATTRIBUTE returns a boolean if a field has been set.

### GetPASSWORD

`func (o *ADSIConnectionAttributes) GetPASSWORD() string`

GetPASSWORD returns the PASSWORD field if non-nil, zero value otherwise.

### GetPASSWORDOk

`func (o *ADSIConnectionAttributes) GetPASSWORDOk() (*string, bool)`

GetPASSWORDOk returns a tuple with the PASSWORD field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPASSWORD

`func (o *ADSIConnectionAttributes) SetPASSWORD(v string)`

SetPASSWORD sets PASSWORD field to given value.

### HasPASSWORD

`func (o *ADSIConnectionAttributes) HasPASSWORD() bool`

HasPASSWORD returns a boolean if a field has been set.

### GetPAM_CONFIG

`func (o *ADSIConnectionAttributes) GetPAM_CONFIG() string`

GetPAM_CONFIG returns the PAM_CONFIG field if non-nil, zero value otherwise.

### GetPAM_CONFIGOk

`func (o *ADSIConnectionAttributes) GetPAM_CONFIGOk() (*string, bool)`

GetPAM_CONFIGOk returns a tuple with the PAM_CONFIG field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPAM_CONFIG

`func (o *ADSIConnectionAttributes) SetPAM_CONFIG(v string)`

SetPAM_CONFIG sets PAM_CONFIG field to given value.

### HasPAM_CONFIG

`func (o *ADSIConnectionAttributes) HasPAM_CONFIG() bool`

HasPAM_CONFIG returns a boolean if a field has been set.

### GetPAGE_SIZE

`func (o *ADSIConnectionAttributes) GetPAGE_SIZE() string`

GetPAGE_SIZE returns the PAGE_SIZE field if non-nil, zero value otherwise.

### GetPAGE_SIZEOk

`func (o *ADSIConnectionAttributes) GetPAGE_SIZEOk() (*string, bool)`

GetPAGE_SIZEOk returns a tuple with the PAGE_SIZE field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPAGE_SIZE

`func (o *ADSIConnectionAttributes) SetPAGE_SIZE(v string)`

SetPAGE_SIZE sets PAGE_SIZE field to given value.

### HasPAGE_SIZE

`func (o *ADSIConnectionAttributes) HasPAGE_SIZE() bool`

HasPAGE_SIZE returns a boolean if a field has been set.

### GetSEARCHFILTER

`func (o *ADSIConnectionAttributes) GetSEARCHFILTER() string`

GetSEARCHFILTER returns the SEARCHFILTER field if non-nil, zero value otherwise.

### GetSEARCHFILTEROk

`func (o *ADSIConnectionAttributes) GetSEARCHFILTEROk() (*string, bool)`

GetSEARCHFILTEROk returns a tuple with the SEARCHFILTER field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEARCHFILTER

`func (o *ADSIConnectionAttributes) SetSEARCHFILTER(v string)`

SetSEARCHFILTER sets SEARCHFILTER field to given value.

### HasSEARCHFILTER

`func (o *ADSIConnectionAttributes) HasSEARCHFILTER() bool`

HasSEARCHFILTER returns a boolean if a field has been set.

### GetUPDATEGROUPJSON

`func (o *ADSIConnectionAttributes) GetUPDATEGROUPJSON() string`

GetUPDATEGROUPJSON returns the UPDATEGROUPJSON field if non-nil, zero value otherwise.

### GetUPDATEGROUPJSONOk

`func (o *ADSIConnectionAttributes) GetUPDATEGROUPJSONOk() (*string, bool)`

GetUPDATEGROUPJSONOk returns a tuple with the UPDATEGROUPJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUPDATEGROUPJSON

`func (o *ADSIConnectionAttributes) SetUPDATEGROUPJSON(v string)`

SetUPDATEGROUPJSON sets UPDATEGROUPJSON field to given value.

### HasUPDATEGROUPJSON

`func (o *ADSIConnectionAttributes) HasUPDATEGROUPJSON() bool`

HasUPDATEGROUPJSON returns a boolean if a field has been set.

### GetCREATEGROUPJSON

`func (o *ADSIConnectionAttributes) GetCREATEGROUPJSON() string`

GetCREATEGROUPJSON returns the CREATEGROUPJSON field if non-nil, zero value otherwise.

### GetCREATEGROUPJSONOk

`func (o *ADSIConnectionAttributes) GetCREATEGROUPJSONOk() (*string, bool)`

GetCREATEGROUPJSONOk returns a tuple with the CREATEGROUPJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCREATEGROUPJSON

`func (o *ADSIConnectionAttributes) SetCREATEGROUPJSON(v string)`

SetCREATEGROUPJSON sets CREATEGROUPJSON field to given value.

### HasCREATEGROUPJSON

`func (o *ADSIConnectionAttributes) HasCREATEGROUPJSON() bool`

HasCREATEGROUPJSON returns a boolean if a field has been set.

### GetENTITLEMENT_ATTRIBUTE

`func (o *ADSIConnectionAttributes) GetENTITLEMENT_ATTRIBUTE() string`

GetENTITLEMENT_ATTRIBUTE returns the ENTITLEMENT_ATTRIBUTE field if non-nil, zero value otherwise.

### GetENTITLEMENT_ATTRIBUTEOk

`func (o *ADSIConnectionAttributes) GetENTITLEMENT_ATTRIBUTEOk() (*string, bool)`

GetENTITLEMENT_ATTRIBUTEOk returns a tuple with the ENTITLEMENT_ATTRIBUTE field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetENTITLEMENT_ATTRIBUTE

`func (o *ADSIConnectionAttributes) SetENTITLEMENT_ATTRIBUTE(v string)`

SetENTITLEMENT_ATTRIBUTE sets ENTITLEMENT_ATTRIBUTE field to given value.

### HasENTITLEMENT_ATTRIBUTE

`func (o *ADSIConnectionAttributes) HasENTITLEMENT_ATTRIBUTE() bool`

HasENTITLEMENT_ATTRIBUTE returns a boolean if a field has been set.

### GetCHECKFORUNIQUE

`func (o *ADSIConnectionAttributes) GetCHECKFORUNIQUE() string`

GetCHECKFORUNIQUE returns the CHECKFORUNIQUE field if non-nil, zero value otherwise.

### GetCHECKFORUNIQUEOk

`func (o *ADSIConnectionAttributes) GetCHECKFORUNIQUEOk() (*string, bool)`

GetCHECKFORUNIQUEOk returns a tuple with the CHECKFORUNIQUE field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCHECKFORUNIQUE

`func (o *ADSIConnectionAttributes) SetCHECKFORUNIQUE(v string)`

SetCHECKFORUNIQUE sets CHECKFORUNIQUE field to given value.

### HasCHECKFORUNIQUE

`func (o *ADSIConnectionAttributes) HasCHECKFORUNIQUE() bool`

HasCHECKFORUNIQUE returns a boolean if a field has been set.

### GetREMOVESERVICEACCOUNTJSON

`func (o *ADSIConnectionAttributes) GetREMOVESERVICEACCOUNTJSON() string`

GetREMOVESERVICEACCOUNTJSON returns the REMOVESERVICEACCOUNTJSON field if non-nil, zero value otherwise.

### GetREMOVESERVICEACCOUNTJSONOk

`func (o *ADSIConnectionAttributes) GetREMOVESERVICEACCOUNTJSONOk() (*string, bool)`

GetREMOVESERVICEACCOUNTJSONOk returns a tuple with the REMOVESERVICEACCOUNTJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetREMOVESERVICEACCOUNTJSON

`func (o *ADSIConnectionAttributes) SetREMOVESERVICEACCOUNTJSON(v string)`

SetREMOVESERVICEACCOUNTJSON sets REMOVESERVICEACCOUNTJSON field to given value.

### HasREMOVESERVICEACCOUNTJSON

`func (o *ADSIConnectionAttributes) HasREMOVESERVICEACCOUNTJSON() bool`

HasREMOVESERVICEACCOUNTJSON returns a boolean if a field has been set.

### GetConnectionTimeoutConfig

`func (o *ADSIConnectionAttributes) GetConnectionTimeoutConfig() ConnectionTimeoutConfig`

GetConnectionTimeoutConfig returns the ConnectionTimeoutConfig field if non-nil, zero value otherwise.

### GetConnectionTimeoutConfigOk

`func (o *ADSIConnectionAttributes) GetConnectionTimeoutConfigOk() (*ConnectionTimeoutConfig, bool)`

GetConnectionTimeoutConfigOk returns a tuple with the ConnectionTimeoutConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionTimeoutConfig

`func (o *ADSIConnectionAttributes) SetConnectionTimeoutConfig(v ConnectionTimeoutConfig)`

SetConnectionTimeoutConfig sets ConnectionTimeoutConfig field to given value.

### HasConnectionTimeoutConfig

`func (o *ADSIConnectionAttributes) HasConnectionTimeoutConfig() bool`

HasConnectionTimeoutConfig returns a boolean if a field has been set.

### GetUPDATEUSERJSON

`func (o *ADSIConnectionAttributes) GetUPDATEUSERJSON() string`

GetUPDATEUSERJSON returns the UPDATEUSERJSON field if non-nil, zero value otherwise.

### GetUPDATEUSERJSONOk

`func (o *ADSIConnectionAttributes) GetUPDATEUSERJSONOk() (*string, bool)`

GetUPDATEUSERJSONOk returns a tuple with the UPDATEUSERJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUPDATEUSERJSON

`func (o *ADSIConnectionAttributes) SetUPDATEUSERJSON(v string)`

SetUPDATEUSERJSON sets UPDATEUSERJSON field to given value.

### HasUPDATEUSERJSON

`func (o *ADSIConnectionAttributes) HasUPDATEUSERJSON() bool`

HasUPDATEUSERJSON returns a boolean if a field has been set.

### GetURL

`func (o *ADSIConnectionAttributes) GetURL() string`

GetURL returns the URL field if non-nil, zero value otherwise.

### GetURLOk

`func (o *ADSIConnectionAttributes) GetURLOk() (*string, bool)`

GetURLOk returns a tuple with the URL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetURL

`func (o *ADSIConnectionAttributes) SetURL(v string)`

SetURL sets URL field to given value.

### HasURL

`func (o *ADSIConnectionAttributes) HasURL() bool`

HasURL returns a boolean if a field has been set.

### GetMOVEACCOUNTJSON

`func (o *ADSIConnectionAttributes) GetMOVEACCOUNTJSON() string`

GetMOVEACCOUNTJSON returns the MOVEACCOUNTJSON field if non-nil, zero value otherwise.

### GetMOVEACCOUNTJSONOk

`func (o *ADSIConnectionAttributes) GetMOVEACCOUNTJSONOk() (*string, bool)`

GetMOVEACCOUNTJSONOk returns a tuple with the MOVEACCOUNTJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMOVEACCOUNTJSON

`func (o *ADSIConnectionAttributes) SetMOVEACCOUNTJSON(v string)`

SetMOVEACCOUNTJSON sets MOVEACCOUNTJSON field to given value.

### HasMOVEACCOUNTJSON

`func (o *ADSIConnectionAttributes) HasMOVEACCOUNTJSON() bool`

HasMOVEACCOUNTJSON returns a boolean if a field has been set.

### GetCUSTOMCONFIGJSON

`func (o *ADSIConnectionAttributes) GetCUSTOMCONFIGJSON() string`

GetCUSTOMCONFIGJSON returns the CUSTOMCONFIGJSON field if non-nil, zero value otherwise.

### GetCUSTOMCONFIGJSONOk

`func (o *ADSIConnectionAttributes) GetCUSTOMCONFIGJSONOk() (*string, bool)`

GetCUSTOMCONFIGJSONOk returns a tuple with the CUSTOMCONFIGJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCUSTOMCONFIGJSON

`func (o *ADSIConnectionAttributes) SetCUSTOMCONFIGJSON(v string)`

SetCUSTOMCONFIGJSON sets CUSTOMCONFIGJSON field to given value.

### HasCUSTOMCONFIGJSON

`func (o *ADSIConnectionAttributes) HasCUSTOMCONFIGJSON() bool`

HasCUSTOMCONFIGJSON returns a boolean if a field has been set.

### GetSTATUS_THRESHOLD_CONFIG

`func (o *ADSIConnectionAttributes) GetSTATUS_THRESHOLD_CONFIG() string`

GetSTATUS_THRESHOLD_CONFIG returns the STATUS_THRESHOLD_CONFIG field if non-nil, zero value otherwise.

### GetSTATUS_THRESHOLD_CONFIGOk

`func (o *ADSIConnectionAttributes) GetSTATUS_THRESHOLD_CONFIGOk() (*string, bool)`

GetSTATUS_THRESHOLD_CONFIGOk returns a tuple with the STATUS_THRESHOLD_CONFIG field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSTATUS_THRESHOLD_CONFIG

`func (o *ADSIConnectionAttributes) SetSTATUS_THRESHOLD_CONFIG(v string)`

SetSTATUS_THRESHOLD_CONFIG sets STATUS_THRESHOLD_CONFIG field to given value.

### HasSTATUS_THRESHOLD_CONFIG

`func (o *ADSIConnectionAttributes) HasSTATUS_THRESHOLD_CONFIG() bool`

HasSTATUS_THRESHOLD_CONFIG returns a boolean if a field has been set.

### GetGroupImportMapping

`func (o *ADSIConnectionAttributes) GetGroupImportMapping() string`

GetGroupImportMapping returns the GroupImportMapping field if non-nil, zero value otherwise.

### GetGroupImportMappingOk

`func (o *ADSIConnectionAttributes) GetGroupImportMappingOk() (*string, bool)`

GetGroupImportMappingOk returns a tuple with the GroupImportMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupImportMapping

`func (o *ADSIConnectionAttributes) SetGroupImportMapping(v string)`

SetGroupImportMapping sets GroupImportMapping field to given value.

### HasGroupImportMapping

`func (o *ADSIConnectionAttributes) HasGroupImportMapping() bool`

HasGroupImportMapping returns a boolean if a field has been set.

### GetPROVISIONING_URL

`func (o *ADSIConnectionAttributes) GetPROVISIONING_URL() string`

GetPROVISIONING_URL returns the PROVISIONING_URL field if non-nil, zero value otherwise.

### GetPROVISIONING_URLOk

`func (o *ADSIConnectionAttributes) GetPROVISIONING_URLOk() (*string, bool)`

GetPROVISIONING_URLOk returns a tuple with the PROVISIONING_URL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPROVISIONING_URL

`func (o *ADSIConnectionAttributes) SetPROVISIONING_URL(v string)`

SetPROVISIONING_URL sets PROVISIONING_URL field to given value.

### HasPROVISIONING_URL

`func (o *ADSIConnectionAttributes) HasPROVISIONING_URL() bool`

HasPROVISIONING_URL returns a boolean if a field has been set.

### GetREMOVEGROUPJSON

`func (o *ADSIConnectionAttributes) GetREMOVEGROUPJSON() string`

GetREMOVEGROUPJSON returns the REMOVEGROUPJSON field if non-nil, zero value otherwise.

### GetREMOVEGROUPJSONOk

`func (o *ADSIConnectionAttributes) GetREMOVEGROUPJSONOk() (*string, bool)`

GetREMOVEGROUPJSONOk returns a tuple with the REMOVEGROUPJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetREMOVEGROUPJSON

`func (o *ADSIConnectionAttributes) SetREMOVEGROUPJSON(v string)`

SetREMOVEGROUPJSON sets REMOVEGROUPJSON field to given value.

### HasREMOVEGROUPJSON

`func (o *ADSIConnectionAttributes) HasREMOVEGROUPJSON() bool`

HasREMOVEGROUPJSON returns a boolean if a field has been set.

### GetREMOVEACCESSJSON

`func (o *ADSIConnectionAttributes) GetREMOVEACCESSJSON() string`

GetREMOVEACCESSJSON returns the REMOVEACCESSJSON field if non-nil, zero value otherwise.

### GetREMOVEACCESSJSONOk

`func (o *ADSIConnectionAttributes) GetREMOVEACCESSJSONOk() (*string, bool)`

GetREMOVEACCESSJSONOk returns a tuple with the REMOVEACCESSJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetREMOVEACCESSJSON

`func (o *ADSIConnectionAttributes) SetREMOVEACCESSJSON(v string)`

SetREMOVEACCESSJSON sets REMOVEACCESSJSON field to given value.

### HasREMOVEACCESSJSON

`func (o *ADSIConnectionAttributes) HasREMOVEACCESSJSON() bool`

HasREMOVEACCESSJSON returns a boolean if a field has been set.

### GetIMPORTDATACOOKIES

`func (o *ADSIConnectionAttributes) GetIMPORTDATACOOKIES() string`

GetIMPORTDATACOOKIES returns the IMPORTDATACOOKIES field if non-nil, zero value otherwise.

### GetIMPORTDATACOOKIESOk

`func (o *ADSIConnectionAttributes) GetIMPORTDATACOOKIESOk() (*string, bool)`

GetIMPORTDATACOOKIESOk returns a tuple with the IMPORTDATACOOKIES field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIMPORTDATACOOKIES

`func (o *ADSIConnectionAttributes) SetIMPORTDATACOOKIES(v string)`

SetIMPORTDATACOOKIES sets IMPORTDATACOOKIES field to given value.

### HasIMPORTDATACOOKIES

`func (o *ADSIConnectionAttributes) HasIMPORTDATACOOKIES() bool`

HasIMPORTDATACOOKIES returns a boolean if a field has been set.

### GetRESETANDCHANGEPASSWRDJSON

`func (o *ADSIConnectionAttributes) GetRESETANDCHANGEPASSWRDJSON() string`

GetRESETANDCHANGEPASSWRDJSON returns the RESETANDCHANGEPASSWRDJSON field if non-nil, zero value otherwise.

### GetRESETANDCHANGEPASSWRDJSONOk

`func (o *ADSIConnectionAttributes) GetRESETANDCHANGEPASSWRDJSONOk() (*string, bool)`

GetRESETANDCHANGEPASSWRDJSONOk returns a tuple with the RESETANDCHANGEPASSWRDJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRESETANDCHANGEPASSWRDJSON

`func (o *ADSIConnectionAttributes) SetRESETANDCHANGEPASSWRDJSON(v string)`

SetRESETANDCHANGEPASSWRDJSON sets RESETANDCHANGEPASSWRDJSON field to given value.

### HasRESETANDCHANGEPASSWRDJSON

`func (o *ADSIConnectionAttributes) HasRESETANDCHANGEPASSWRDJSON() bool`

HasRESETANDCHANGEPASSWRDJSON returns a boolean if a field has been set.

### GetUSER_ATTRIBUTE

`func (o *ADSIConnectionAttributes) GetUSER_ATTRIBUTE() string`

GetUSER_ATTRIBUTE returns the USER_ATTRIBUTE field if non-nil, zero value otherwise.

### GetUSER_ATTRIBUTEOk

`func (o *ADSIConnectionAttributes) GetUSER_ATTRIBUTEOk() (*string, bool)`

GetUSER_ATTRIBUTEOk returns a tuple with the USER_ATTRIBUTE field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUSER_ATTRIBUTE

`func (o *ADSIConnectionAttributes) SetUSER_ATTRIBUTE(v string)`

SetUSER_ATTRIBUTE sets USER_ATTRIBUTE field to given value.

### HasUSER_ATTRIBUTE

`func (o *ADSIConnectionAttributes) HasUSER_ATTRIBUTE() bool`

HasUSER_ATTRIBUTE returns a boolean if a field has been set.

### GetADDACCESSENTITLEMENTJSON

`func (o *ADSIConnectionAttributes) GetADDACCESSENTITLEMENTJSON() string`

GetADDACCESSENTITLEMENTJSON returns the ADDACCESSENTITLEMENTJSON field if non-nil, zero value otherwise.

### GetADDACCESSENTITLEMENTJSONOk

`func (o *ADSIConnectionAttributes) GetADDACCESSENTITLEMENTJSONOk() (*string, bool)`

GetADDACCESSENTITLEMENTJSONOk returns a tuple with the ADDACCESSENTITLEMENTJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetADDACCESSENTITLEMENTJSON

`func (o *ADSIConnectionAttributes) SetADDACCESSENTITLEMENTJSON(v string)`

SetADDACCESSENTITLEMENTJSON sets ADDACCESSENTITLEMENTJSON field to given value.

### HasADDACCESSENTITLEMENTJSON

`func (o *ADSIConnectionAttributes) HasADDACCESSENTITLEMENTJSON() bool`

HasADDACCESSENTITLEMENTJSON returns a boolean if a field has been set.

### GetMODIFYUSERDATAJSON

`func (o *ADSIConnectionAttributes) GetMODIFYUSERDATAJSON() string`

GetMODIFYUSERDATAJSON returns the MODIFYUSERDATAJSON field if non-nil, zero value otherwise.

### GetMODIFYUSERDATAJSONOk

`func (o *ADSIConnectionAttributes) GetMODIFYUSERDATAJSONOk() (*string, bool)`

GetMODIFYUSERDATAJSONOk returns a tuple with the MODIFYUSERDATAJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMODIFYUSERDATAJSON

`func (o *ADSIConnectionAttributes) SetMODIFYUSERDATAJSON(v string)`

SetMODIFYUSERDATAJSON sets MODIFYUSERDATAJSON field to given value.

### HasMODIFYUSERDATAJSON

`func (o *ADSIConnectionAttributes) HasMODIFYUSERDATAJSON() bool`

HasMODIFYUSERDATAJSON returns a boolean if a field has been set.

### GetIsTimeoutConfigValidated

`func (o *ADSIConnectionAttributes) GetIsTimeoutConfigValidated() bool`

GetIsTimeoutConfigValidated returns the IsTimeoutConfigValidated field if non-nil, zero value otherwise.

### GetIsTimeoutConfigValidatedOk

`func (o *ADSIConnectionAttributes) GetIsTimeoutConfigValidatedOk() (*bool, bool)`

GetIsTimeoutConfigValidatedOk returns a tuple with the IsTimeoutConfigValidated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsTimeoutConfigValidated

`func (o *ADSIConnectionAttributes) SetIsTimeoutConfigValidated(v bool)`

SetIsTimeoutConfigValidated sets IsTimeoutConfigValidated field to given value.

### HasIsTimeoutConfigValidated

`func (o *ADSIConnectionAttributes) HasIsTimeoutConfigValidated() bool`

HasIsTimeoutConfigValidated returns a boolean if a field has been set.

### GetENABLEGROUPMANAGEMENT

`func (o *ADSIConnectionAttributes) GetENABLEGROUPMANAGEMENT() string`

GetENABLEGROUPMANAGEMENT returns the ENABLEGROUPMANAGEMENT field if non-nil, zero value otherwise.

### GetENABLEGROUPMANAGEMENTOk

`func (o *ADSIConnectionAttributes) GetENABLEGROUPMANAGEMENTOk() (*string, bool)`

GetENABLEGROUPMANAGEMENTOk returns a tuple with the ENABLEGROUPMANAGEMENT field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetENABLEGROUPMANAGEMENT

`func (o *ADSIConnectionAttributes) SetENABLEGROUPMANAGEMENT(v string)`

SetENABLEGROUPMANAGEMENT sets ENABLEGROUPMANAGEMENT field to given value.

### HasENABLEGROUPMANAGEMENT

`func (o *ADSIConnectionAttributes) HasENABLEGROUPMANAGEMENT() bool`

HasENABLEGROUPMANAGEMENT returns a boolean if a field has been set.

### GetENABLEACCOUNTJSON

`func (o *ADSIConnectionAttributes) GetENABLEACCOUNTJSON() string`

GetENABLEACCOUNTJSON returns the ENABLEACCOUNTJSON field if non-nil, zero value otherwise.

### GetENABLEACCOUNTJSONOk

`func (o *ADSIConnectionAttributes) GetENABLEACCOUNTJSONOk() (*string, bool)`

GetENABLEACCOUNTJSONOk returns a tuple with the ENABLEACCOUNTJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetENABLEACCOUNTJSON

`func (o *ADSIConnectionAttributes) SetENABLEACCOUNTJSON(v string)`

SetENABLEACCOUNTJSON sets ENABLEACCOUNTJSON field to given value.

### HasENABLEACCOUNTJSON

`func (o *ADSIConnectionAttributes) HasENABLEACCOUNTJSON() bool`

HasENABLEACCOUNTJSON returns a boolean if a field has been set.

### GetFORESTLIST

`func (o *ADSIConnectionAttributes) GetFORESTLIST() string`

GetFORESTLIST returns the FORESTLIST field if non-nil, zero value otherwise.

### GetFORESTLISTOk

`func (o *ADSIConnectionAttributes) GetFORESTLISTOk() (*string, bool)`

GetFORESTLISTOk returns a tuple with the FORESTLIST field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFORESTLIST

`func (o *ADSIConnectionAttributes) SetFORESTLIST(v string)`

SetFORESTLIST sets FORESTLIST field to given value.

### HasFORESTLIST

`func (o *ADSIConnectionAttributes) HasFORESTLIST() bool`

HasFORESTLIST returns a boolean if a field has been set.

### GetOBJECTFILTER

`func (o *ADSIConnectionAttributes) GetOBJECTFILTER() string`

GetOBJECTFILTER returns the OBJECTFILTER field if non-nil, zero value otherwise.

### GetOBJECTFILTEROk

`func (o *ADSIConnectionAttributes) GetOBJECTFILTEROk() (*string, bool)`

GetOBJECTFILTEROk returns a tuple with the OBJECTFILTER field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOBJECTFILTER

`func (o *ADSIConnectionAttributes) SetOBJECTFILTER(v string)`

SetOBJECTFILTER sets OBJECTFILTER field to given value.

### HasOBJECTFILTER

`func (o *ADSIConnectionAttributes) HasOBJECTFILTER() bool`

HasOBJECTFILTER returns a boolean if a field has been set.

### GetUPDATEACCOUNTJSON

`func (o *ADSIConnectionAttributes) GetUPDATEACCOUNTJSON() string`

GetUPDATEACCOUNTJSON returns the UPDATEACCOUNTJSON field if non-nil, zero value otherwise.

### GetUPDATEACCOUNTJSONOk

`func (o *ADSIConnectionAttributes) GetUPDATEACCOUNTJSONOk() (*string, bool)`

GetUPDATEACCOUNTJSONOk returns a tuple with the UPDATEACCOUNTJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUPDATEACCOUNTJSON

`func (o *ADSIConnectionAttributes) SetUPDATEACCOUNTJSON(v string)`

SetUPDATEACCOUNTJSON sets UPDATEACCOUNTJSON field to given value.

### HasUPDATEACCOUNTJSON

`func (o *ADSIConnectionAttributes) HasUPDATEACCOUNTJSON() bool`

HasUPDATEACCOUNTJSON returns a boolean if a field has been set.

### GetREMOVEACCOUNTJSON

`func (o *ADSIConnectionAttributes) GetREMOVEACCOUNTJSON() string`

GetREMOVEACCOUNTJSON returns the REMOVEACCOUNTJSON field if non-nil, zero value otherwise.

### GetREMOVEACCOUNTJSONOk

`func (o *ADSIConnectionAttributes) GetREMOVEACCOUNTJSONOk() (*string, bool)`

GetREMOVEACCOUNTJSONOk returns a tuple with the REMOVEACCOUNTJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetREMOVEACCOUNTJSON

`func (o *ADSIConnectionAttributes) SetREMOVEACCOUNTJSON(v string)`

SetREMOVEACCOUNTJSON sets REMOVEACCOUNTJSON field to given value.

### HasREMOVEACCOUNTJSON

`func (o *ADSIConnectionAttributes) HasREMOVEACCOUNTJSON() bool`

HasREMOVEACCOUNTJSON returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


