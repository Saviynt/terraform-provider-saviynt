# SalesforceConnectionAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsTimeoutSupported** | Pointer to **bool** |  | [optional] 
**CLIENT_SECRET** | Pointer to **string** |  | [optional] 
**OBJECT_TO_BE_IMPORTED** | Pointer to **string** |  | [optional] 
**FEATURE_LICENSE_JSON** | Pointer to **string** |  | [optional] 
**CREATEACCOUNTJSON** | Pointer to **string** |  | [optional] 
**REDIRECT_URI** | Pointer to **string** |  | [optional] 
**REFRESH_TOKEN** | Pointer to **string** |  | [optional] 
**ConnectionTimeoutConfig** | Pointer to [**ConnectionTimeoutConfig**](ConnectionTimeoutConfig.md) |  | [optional] 
**MODIFYACCOUNTJSON** | Pointer to **string** |  | [optional] 
**ConnectionType** | Pointer to **string** |  | [optional] 
**IsTimeoutConfigValidated** | Pointer to **bool** |  | [optional] 
**CLIENT_ID** | Pointer to **string** |  | [optional] 
**PAM_CONFIG** | Pointer to **string** |  | [optional] 
**CUSTOMCONFIGJSON** | Pointer to **string** |  | [optional] 
**FIELD_MAPPING_JSON** | Pointer to **string** |  | [optional] 
**STATUS_THRESHOLD_CONFIG** | Pointer to **string** |  | [optional] 
**ACCOUNT_FIELD_QUERY** | Pointer to **string** |  | [optional] 
**CUSTOM_CREATEACCOUNT_URL** | Pointer to **string** |  | [optional] 
**ACCOUNT_FILTER_QUERY** | Pointer to **string** |  | [optional] 
**INSTANCE_URL** | Pointer to **string** |  | [optional] 

## Methods

### NewSalesforceConnectionAttributes

`func NewSalesforceConnectionAttributes() *SalesforceConnectionAttributes`

NewSalesforceConnectionAttributes instantiates a new SalesforceConnectionAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSalesforceConnectionAttributesWithDefaults

`func NewSalesforceConnectionAttributesWithDefaults() *SalesforceConnectionAttributes`

NewSalesforceConnectionAttributesWithDefaults instantiates a new SalesforceConnectionAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsTimeoutSupported

`func (o *SalesforceConnectionAttributes) GetIsTimeoutSupported() bool`

GetIsTimeoutSupported returns the IsTimeoutSupported field if non-nil, zero value otherwise.

### GetIsTimeoutSupportedOk

`func (o *SalesforceConnectionAttributes) GetIsTimeoutSupportedOk() (*bool, bool)`

GetIsTimeoutSupportedOk returns a tuple with the IsTimeoutSupported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsTimeoutSupported

`func (o *SalesforceConnectionAttributes) SetIsTimeoutSupported(v bool)`

SetIsTimeoutSupported sets IsTimeoutSupported field to given value.

### HasIsTimeoutSupported

`func (o *SalesforceConnectionAttributes) HasIsTimeoutSupported() bool`

HasIsTimeoutSupported returns a boolean if a field has been set.

### GetCLIENT_SECRET

`func (o *SalesforceConnectionAttributes) GetCLIENT_SECRET() string`

GetCLIENT_SECRET returns the CLIENT_SECRET field if non-nil, zero value otherwise.

### GetCLIENT_SECRETOk

`func (o *SalesforceConnectionAttributes) GetCLIENT_SECRETOk() (*string, bool)`

GetCLIENT_SECRETOk returns a tuple with the CLIENT_SECRET field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCLIENT_SECRET

`func (o *SalesforceConnectionAttributes) SetCLIENT_SECRET(v string)`

SetCLIENT_SECRET sets CLIENT_SECRET field to given value.

### HasCLIENT_SECRET

`func (o *SalesforceConnectionAttributes) HasCLIENT_SECRET() bool`

HasCLIENT_SECRET returns a boolean if a field has been set.

### GetOBJECT_TO_BE_IMPORTED

`func (o *SalesforceConnectionAttributes) GetOBJECT_TO_BE_IMPORTED() string`

GetOBJECT_TO_BE_IMPORTED returns the OBJECT_TO_BE_IMPORTED field if non-nil, zero value otherwise.

### GetOBJECT_TO_BE_IMPORTEDOk

`func (o *SalesforceConnectionAttributes) GetOBJECT_TO_BE_IMPORTEDOk() (*string, bool)`

GetOBJECT_TO_BE_IMPORTEDOk returns a tuple with the OBJECT_TO_BE_IMPORTED field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOBJECT_TO_BE_IMPORTED

`func (o *SalesforceConnectionAttributes) SetOBJECT_TO_BE_IMPORTED(v string)`

SetOBJECT_TO_BE_IMPORTED sets OBJECT_TO_BE_IMPORTED field to given value.

### HasOBJECT_TO_BE_IMPORTED

`func (o *SalesforceConnectionAttributes) HasOBJECT_TO_BE_IMPORTED() bool`

HasOBJECT_TO_BE_IMPORTED returns a boolean if a field has been set.

### GetFEATURE_LICENSE_JSON

`func (o *SalesforceConnectionAttributes) GetFEATURE_LICENSE_JSON() string`

GetFEATURE_LICENSE_JSON returns the FEATURE_LICENSE_JSON field if non-nil, zero value otherwise.

### GetFEATURE_LICENSE_JSONOk

`func (o *SalesforceConnectionAttributes) GetFEATURE_LICENSE_JSONOk() (*string, bool)`

GetFEATURE_LICENSE_JSONOk returns a tuple with the FEATURE_LICENSE_JSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFEATURE_LICENSE_JSON

`func (o *SalesforceConnectionAttributes) SetFEATURE_LICENSE_JSON(v string)`

SetFEATURE_LICENSE_JSON sets FEATURE_LICENSE_JSON field to given value.

### HasFEATURE_LICENSE_JSON

`func (o *SalesforceConnectionAttributes) HasFEATURE_LICENSE_JSON() bool`

HasFEATURE_LICENSE_JSON returns a boolean if a field has been set.

### GetCREATEACCOUNTJSON

`func (o *SalesforceConnectionAttributes) GetCREATEACCOUNTJSON() string`

GetCREATEACCOUNTJSON returns the CREATEACCOUNTJSON field if non-nil, zero value otherwise.

### GetCREATEACCOUNTJSONOk

`func (o *SalesforceConnectionAttributes) GetCREATEACCOUNTJSONOk() (*string, bool)`

GetCREATEACCOUNTJSONOk returns a tuple with the CREATEACCOUNTJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCREATEACCOUNTJSON

`func (o *SalesforceConnectionAttributes) SetCREATEACCOUNTJSON(v string)`

SetCREATEACCOUNTJSON sets CREATEACCOUNTJSON field to given value.

### HasCREATEACCOUNTJSON

`func (o *SalesforceConnectionAttributes) HasCREATEACCOUNTJSON() bool`

HasCREATEACCOUNTJSON returns a boolean if a field has been set.

### GetREDIRECT_URI

`func (o *SalesforceConnectionAttributes) GetREDIRECT_URI() string`

GetREDIRECT_URI returns the REDIRECT_URI field if non-nil, zero value otherwise.

### GetREDIRECT_URIOk

`func (o *SalesforceConnectionAttributes) GetREDIRECT_URIOk() (*string, bool)`

GetREDIRECT_URIOk returns a tuple with the REDIRECT_URI field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetREDIRECT_URI

`func (o *SalesforceConnectionAttributes) SetREDIRECT_URI(v string)`

SetREDIRECT_URI sets REDIRECT_URI field to given value.

### HasREDIRECT_URI

`func (o *SalesforceConnectionAttributes) HasREDIRECT_URI() bool`

HasREDIRECT_URI returns a boolean if a field has been set.

### GetREFRESH_TOKEN

`func (o *SalesforceConnectionAttributes) GetREFRESH_TOKEN() string`

GetREFRESH_TOKEN returns the REFRESH_TOKEN field if non-nil, zero value otherwise.

### GetREFRESH_TOKENOk

`func (o *SalesforceConnectionAttributes) GetREFRESH_TOKENOk() (*string, bool)`

GetREFRESH_TOKENOk returns a tuple with the REFRESH_TOKEN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetREFRESH_TOKEN

`func (o *SalesforceConnectionAttributes) SetREFRESH_TOKEN(v string)`

SetREFRESH_TOKEN sets REFRESH_TOKEN field to given value.

### HasREFRESH_TOKEN

`func (o *SalesforceConnectionAttributes) HasREFRESH_TOKEN() bool`

HasREFRESH_TOKEN returns a boolean if a field has been set.

### GetConnectionTimeoutConfig

`func (o *SalesforceConnectionAttributes) GetConnectionTimeoutConfig() ConnectionTimeoutConfig`

GetConnectionTimeoutConfig returns the ConnectionTimeoutConfig field if non-nil, zero value otherwise.

### GetConnectionTimeoutConfigOk

`func (o *SalesforceConnectionAttributes) GetConnectionTimeoutConfigOk() (*ConnectionTimeoutConfig, bool)`

GetConnectionTimeoutConfigOk returns a tuple with the ConnectionTimeoutConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionTimeoutConfig

`func (o *SalesforceConnectionAttributes) SetConnectionTimeoutConfig(v ConnectionTimeoutConfig)`

SetConnectionTimeoutConfig sets ConnectionTimeoutConfig field to given value.

### HasConnectionTimeoutConfig

`func (o *SalesforceConnectionAttributes) HasConnectionTimeoutConfig() bool`

HasConnectionTimeoutConfig returns a boolean if a field has been set.

### GetMODIFYACCOUNTJSON

`func (o *SalesforceConnectionAttributes) GetMODIFYACCOUNTJSON() string`

GetMODIFYACCOUNTJSON returns the MODIFYACCOUNTJSON field if non-nil, zero value otherwise.

### GetMODIFYACCOUNTJSONOk

`func (o *SalesforceConnectionAttributes) GetMODIFYACCOUNTJSONOk() (*string, bool)`

GetMODIFYACCOUNTJSONOk returns a tuple with the MODIFYACCOUNTJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMODIFYACCOUNTJSON

`func (o *SalesforceConnectionAttributes) SetMODIFYACCOUNTJSON(v string)`

SetMODIFYACCOUNTJSON sets MODIFYACCOUNTJSON field to given value.

### HasMODIFYACCOUNTJSON

`func (o *SalesforceConnectionAttributes) HasMODIFYACCOUNTJSON() bool`

HasMODIFYACCOUNTJSON returns a boolean if a field has been set.

### GetConnectionType

`func (o *SalesforceConnectionAttributes) GetConnectionType() string`

GetConnectionType returns the ConnectionType field if non-nil, zero value otherwise.

### GetConnectionTypeOk

`func (o *SalesforceConnectionAttributes) GetConnectionTypeOk() (*string, bool)`

GetConnectionTypeOk returns a tuple with the ConnectionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionType

`func (o *SalesforceConnectionAttributes) SetConnectionType(v string)`

SetConnectionType sets ConnectionType field to given value.

### HasConnectionType

`func (o *SalesforceConnectionAttributes) HasConnectionType() bool`

HasConnectionType returns a boolean if a field has been set.

### GetIsTimeoutConfigValidated

`func (o *SalesforceConnectionAttributes) GetIsTimeoutConfigValidated() bool`

GetIsTimeoutConfigValidated returns the IsTimeoutConfigValidated field if non-nil, zero value otherwise.

### GetIsTimeoutConfigValidatedOk

`func (o *SalesforceConnectionAttributes) GetIsTimeoutConfigValidatedOk() (*bool, bool)`

GetIsTimeoutConfigValidatedOk returns a tuple with the IsTimeoutConfigValidated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsTimeoutConfigValidated

`func (o *SalesforceConnectionAttributes) SetIsTimeoutConfigValidated(v bool)`

SetIsTimeoutConfigValidated sets IsTimeoutConfigValidated field to given value.

### HasIsTimeoutConfigValidated

`func (o *SalesforceConnectionAttributes) HasIsTimeoutConfigValidated() bool`

HasIsTimeoutConfigValidated returns a boolean if a field has been set.

### GetCLIENT_ID

`func (o *SalesforceConnectionAttributes) GetCLIENT_ID() string`

GetCLIENT_ID returns the CLIENT_ID field if non-nil, zero value otherwise.

### GetCLIENT_IDOk

`func (o *SalesforceConnectionAttributes) GetCLIENT_IDOk() (*string, bool)`

GetCLIENT_IDOk returns a tuple with the CLIENT_ID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCLIENT_ID

`func (o *SalesforceConnectionAttributes) SetCLIENT_ID(v string)`

SetCLIENT_ID sets CLIENT_ID field to given value.

### HasCLIENT_ID

`func (o *SalesforceConnectionAttributes) HasCLIENT_ID() bool`

HasCLIENT_ID returns a boolean if a field has been set.

### GetPAM_CONFIG

`func (o *SalesforceConnectionAttributes) GetPAM_CONFIG() string`

GetPAM_CONFIG returns the PAM_CONFIG field if non-nil, zero value otherwise.

### GetPAM_CONFIGOk

`func (o *SalesforceConnectionAttributes) GetPAM_CONFIGOk() (*string, bool)`

GetPAM_CONFIGOk returns a tuple with the PAM_CONFIG field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPAM_CONFIG

`func (o *SalesforceConnectionAttributes) SetPAM_CONFIG(v string)`

SetPAM_CONFIG sets PAM_CONFIG field to given value.

### HasPAM_CONFIG

`func (o *SalesforceConnectionAttributes) HasPAM_CONFIG() bool`

HasPAM_CONFIG returns a boolean if a field has been set.

### GetCUSTOMCONFIGJSON

`func (o *SalesforceConnectionAttributes) GetCUSTOMCONFIGJSON() string`

GetCUSTOMCONFIGJSON returns the CUSTOMCONFIGJSON field if non-nil, zero value otherwise.

### GetCUSTOMCONFIGJSONOk

`func (o *SalesforceConnectionAttributes) GetCUSTOMCONFIGJSONOk() (*string, bool)`

GetCUSTOMCONFIGJSONOk returns a tuple with the CUSTOMCONFIGJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCUSTOMCONFIGJSON

`func (o *SalesforceConnectionAttributes) SetCUSTOMCONFIGJSON(v string)`

SetCUSTOMCONFIGJSON sets CUSTOMCONFIGJSON field to given value.

### HasCUSTOMCONFIGJSON

`func (o *SalesforceConnectionAttributes) HasCUSTOMCONFIGJSON() bool`

HasCUSTOMCONFIGJSON returns a boolean if a field has been set.

### GetFIELD_MAPPING_JSON

`func (o *SalesforceConnectionAttributes) GetFIELD_MAPPING_JSON() string`

GetFIELD_MAPPING_JSON returns the FIELD_MAPPING_JSON field if non-nil, zero value otherwise.

### GetFIELD_MAPPING_JSONOk

`func (o *SalesforceConnectionAttributes) GetFIELD_MAPPING_JSONOk() (*string, bool)`

GetFIELD_MAPPING_JSONOk returns a tuple with the FIELD_MAPPING_JSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFIELD_MAPPING_JSON

`func (o *SalesforceConnectionAttributes) SetFIELD_MAPPING_JSON(v string)`

SetFIELD_MAPPING_JSON sets FIELD_MAPPING_JSON field to given value.

### HasFIELD_MAPPING_JSON

`func (o *SalesforceConnectionAttributes) HasFIELD_MAPPING_JSON() bool`

HasFIELD_MAPPING_JSON returns a boolean if a field has been set.

### GetSTATUS_THRESHOLD_CONFIG

`func (o *SalesforceConnectionAttributes) GetSTATUS_THRESHOLD_CONFIG() string`

GetSTATUS_THRESHOLD_CONFIG returns the STATUS_THRESHOLD_CONFIG field if non-nil, zero value otherwise.

### GetSTATUS_THRESHOLD_CONFIGOk

`func (o *SalesforceConnectionAttributes) GetSTATUS_THRESHOLD_CONFIGOk() (*string, bool)`

GetSTATUS_THRESHOLD_CONFIGOk returns a tuple with the STATUS_THRESHOLD_CONFIG field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSTATUS_THRESHOLD_CONFIG

`func (o *SalesforceConnectionAttributes) SetSTATUS_THRESHOLD_CONFIG(v string)`

SetSTATUS_THRESHOLD_CONFIG sets STATUS_THRESHOLD_CONFIG field to given value.

### HasSTATUS_THRESHOLD_CONFIG

`func (o *SalesforceConnectionAttributes) HasSTATUS_THRESHOLD_CONFIG() bool`

HasSTATUS_THRESHOLD_CONFIG returns a boolean if a field has been set.

### GetACCOUNT_FIELD_QUERY

`func (o *SalesforceConnectionAttributes) GetACCOUNT_FIELD_QUERY() string`

GetACCOUNT_FIELD_QUERY returns the ACCOUNT_FIELD_QUERY field if non-nil, zero value otherwise.

### GetACCOUNT_FIELD_QUERYOk

`func (o *SalesforceConnectionAttributes) GetACCOUNT_FIELD_QUERYOk() (*string, bool)`

GetACCOUNT_FIELD_QUERYOk returns a tuple with the ACCOUNT_FIELD_QUERY field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetACCOUNT_FIELD_QUERY

`func (o *SalesforceConnectionAttributes) SetACCOUNT_FIELD_QUERY(v string)`

SetACCOUNT_FIELD_QUERY sets ACCOUNT_FIELD_QUERY field to given value.

### HasACCOUNT_FIELD_QUERY

`func (o *SalesforceConnectionAttributes) HasACCOUNT_FIELD_QUERY() bool`

HasACCOUNT_FIELD_QUERY returns a boolean if a field has been set.

### GetCUSTOM_CREATEACCOUNT_URL

`func (o *SalesforceConnectionAttributes) GetCUSTOM_CREATEACCOUNT_URL() string`

GetCUSTOM_CREATEACCOUNT_URL returns the CUSTOM_CREATEACCOUNT_URL field if non-nil, zero value otherwise.

### GetCUSTOM_CREATEACCOUNT_URLOk

`func (o *SalesforceConnectionAttributes) GetCUSTOM_CREATEACCOUNT_URLOk() (*string, bool)`

GetCUSTOM_CREATEACCOUNT_URLOk returns a tuple with the CUSTOM_CREATEACCOUNT_URL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCUSTOM_CREATEACCOUNT_URL

`func (o *SalesforceConnectionAttributes) SetCUSTOM_CREATEACCOUNT_URL(v string)`

SetCUSTOM_CREATEACCOUNT_URL sets CUSTOM_CREATEACCOUNT_URL field to given value.

### HasCUSTOM_CREATEACCOUNT_URL

`func (o *SalesforceConnectionAttributes) HasCUSTOM_CREATEACCOUNT_URL() bool`

HasCUSTOM_CREATEACCOUNT_URL returns a boolean if a field has been set.

### GetACCOUNT_FILTER_QUERY

`func (o *SalesforceConnectionAttributes) GetACCOUNT_FILTER_QUERY() string`

GetACCOUNT_FILTER_QUERY returns the ACCOUNT_FILTER_QUERY field if non-nil, zero value otherwise.

### GetACCOUNT_FILTER_QUERYOk

`func (o *SalesforceConnectionAttributes) GetACCOUNT_FILTER_QUERYOk() (*string, bool)`

GetACCOUNT_FILTER_QUERYOk returns a tuple with the ACCOUNT_FILTER_QUERY field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetACCOUNT_FILTER_QUERY

`func (o *SalesforceConnectionAttributes) SetACCOUNT_FILTER_QUERY(v string)`

SetACCOUNT_FILTER_QUERY sets ACCOUNT_FILTER_QUERY field to given value.

### HasACCOUNT_FILTER_QUERY

`func (o *SalesforceConnectionAttributes) HasACCOUNT_FILTER_QUERY() bool`

HasACCOUNT_FILTER_QUERY returns a boolean if a field has been set.

### GetINSTANCE_URL

`func (o *SalesforceConnectionAttributes) GetINSTANCE_URL() string`

GetINSTANCE_URL returns the INSTANCE_URL field if non-nil, zero value otherwise.

### GetINSTANCE_URLOk

`func (o *SalesforceConnectionAttributes) GetINSTANCE_URLOk() (*string, bool)`

GetINSTANCE_URLOk returns a tuple with the INSTANCE_URL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetINSTANCE_URL

`func (o *SalesforceConnectionAttributes) SetINSTANCE_URL(v string)`

SetINSTANCE_URL sets INSTANCE_URL field to given value.

### HasINSTANCE_URL

`func (o *SalesforceConnectionAttributes) HasINSTANCE_URL() bool`

HasINSTANCE_URL returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


