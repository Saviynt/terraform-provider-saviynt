# BaseConnector

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConnectionName** | **string** | Name of the connection | 
**Connectiontype** | **string** | Connection type (e.g., &#39;AD&#39; for Active Directory). | 
**Description** | Pointer to **string** | Description for the connection. | [optional] 
**Defaultsavroles** | Pointer to **string** | Default SAV roles for managing the connection. | [optional] 
**EmailTemplate** | Pointer to **string** | Email template for notifications. | [optional] 
**SslCertificate** | Pointer to **string** | SSL certificates to secure the connection. | [optional] 
**VaultConnection** | Pointer to **string** | Specifies the type of vault connection being used (e.g., Hashicorp, AWS Secrets Manager). | [optional] 
**VaultConfiguration** | Pointer to **string** | JSON string specifying vault configuration | [optional] 
**Saveinvault** | Pointer to **string** | Flag indicating whether the encrypted attribute should be saved in the configured vault. | [optional] 

## Methods

### NewBaseConnector

`func NewBaseConnector(connectionName string, connectiontype string, ) *BaseConnector`

NewBaseConnector instantiates a new BaseConnector object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBaseConnectorWithDefaults

`func NewBaseConnectorWithDefaults() *BaseConnector`

NewBaseConnectorWithDefaults instantiates a new BaseConnector object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConnectionName

`func (o *BaseConnector) GetConnectionName() string`

GetConnectionName returns the ConnectionName field if non-nil, zero value otherwise.

### GetConnectionNameOk

`func (o *BaseConnector) GetConnectionNameOk() (*string, bool)`

GetConnectionNameOk returns a tuple with the ConnectionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionName

`func (o *BaseConnector) SetConnectionName(v string)`

SetConnectionName sets ConnectionName field to given value.


### GetConnectiontype

`func (o *BaseConnector) GetConnectiontype() string`

GetConnectiontype returns the Connectiontype field if non-nil, zero value otherwise.

### GetConnectiontypeOk

`func (o *BaseConnector) GetConnectiontypeOk() (*string, bool)`

GetConnectiontypeOk returns a tuple with the Connectiontype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectiontype

`func (o *BaseConnector) SetConnectiontype(v string)`

SetConnectiontype sets Connectiontype field to given value.


### GetDescription

`func (o *BaseConnector) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BaseConnector) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BaseConnector) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BaseConnector) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDefaultsavroles

`func (o *BaseConnector) GetDefaultsavroles() string`

GetDefaultsavroles returns the Defaultsavroles field if non-nil, zero value otherwise.

### GetDefaultsavrolesOk

`func (o *BaseConnector) GetDefaultsavrolesOk() (*string, bool)`

GetDefaultsavrolesOk returns a tuple with the Defaultsavroles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultsavroles

`func (o *BaseConnector) SetDefaultsavroles(v string)`

SetDefaultsavroles sets Defaultsavroles field to given value.

### HasDefaultsavroles

`func (o *BaseConnector) HasDefaultsavroles() bool`

HasDefaultsavroles returns a boolean if a field has been set.

### GetEmailTemplate

`func (o *BaseConnector) GetEmailTemplate() string`

GetEmailTemplate returns the EmailTemplate field if non-nil, zero value otherwise.

### GetEmailTemplateOk

`func (o *BaseConnector) GetEmailTemplateOk() (*string, bool)`

GetEmailTemplateOk returns a tuple with the EmailTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailTemplate

`func (o *BaseConnector) SetEmailTemplate(v string)`

SetEmailTemplate sets EmailTemplate field to given value.

### HasEmailTemplate

`func (o *BaseConnector) HasEmailTemplate() bool`

HasEmailTemplate returns a boolean if a field has been set.

### GetSslCertificate

`func (o *BaseConnector) GetSslCertificate() string`

GetSslCertificate returns the SslCertificate field if non-nil, zero value otherwise.

### GetSslCertificateOk

`func (o *BaseConnector) GetSslCertificateOk() (*string, bool)`

GetSslCertificateOk returns a tuple with the SslCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslCertificate

`func (o *BaseConnector) SetSslCertificate(v string)`

SetSslCertificate sets SslCertificate field to given value.

### HasSslCertificate

`func (o *BaseConnector) HasSslCertificate() bool`

HasSslCertificate returns a boolean if a field has been set.

### GetVaultConnection

`func (o *BaseConnector) GetVaultConnection() string`

GetVaultConnection returns the VaultConnection field if non-nil, zero value otherwise.

### GetVaultConnectionOk

`func (o *BaseConnector) GetVaultConnectionOk() (*string, bool)`

GetVaultConnectionOk returns a tuple with the VaultConnection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVaultConnection

`func (o *BaseConnector) SetVaultConnection(v string)`

SetVaultConnection sets VaultConnection field to given value.

### HasVaultConnection

`func (o *BaseConnector) HasVaultConnection() bool`

HasVaultConnection returns a boolean if a field has been set.

### GetVaultConfiguration

`func (o *BaseConnector) GetVaultConfiguration() string`

GetVaultConfiguration returns the VaultConfiguration field if non-nil, zero value otherwise.

### GetVaultConfigurationOk

`func (o *BaseConnector) GetVaultConfigurationOk() (*string, bool)`

GetVaultConfigurationOk returns a tuple with the VaultConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVaultConfiguration

`func (o *BaseConnector) SetVaultConfiguration(v string)`

SetVaultConfiguration sets VaultConfiguration field to given value.

### HasVaultConfiguration

`func (o *BaseConnector) HasVaultConfiguration() bool`

HasVaultConfiguration returns a boolean if a field has been set.

### GetSaveinvault

`func (o *BaseConnector) GetSaveinvault() string`

GetSaveinvault returns the Saveinvault field if non-nil, zero value otherwise.

### GetSaveinvaultOk

`func (o *BaseConnector) GetSaveinvaultOk() (*string, bool)`

GetSaveinvaultOk returns a tuple with the Saveinvault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaveinvault

`func (o *BaseConnector) SetSaveinvault(v string)`

SetSaveinvault sets Saveinvault field to given value.

### HasSaveinvault

`func (o *BaseConnector) HasSaveinvault() bool`

HasSaveinvault returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


