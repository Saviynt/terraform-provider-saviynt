# RESTConnector

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConnectionJSON** | Pointer to **string** |  | [optional] 
**ImportUserJSON** | Pointer to **string** | Property for ImportUserJSON | [optional] 
**ImportAccountEntJSON** | Pointer to **string** |  | [optional] 
**STATUS_THRESHOLD_CONFIG** | Pointer to **string** |  | [optional] 
**CreateAccountJSON** | Pointer to **string** |  | [optional] 
**UpdateAccountJSON** | Pointer to **string** |  | [optional] 
**EnableAccountJSON** | Pointer to **string** |  | [optional] 
**DisableAccountJSON** | Pointer to **string** |  | [optional] 
**AddAccessJSON** | Pointer to **string** |  | [optional] 
**RemoveAccessJSON** | Pointer to **string** |  | [optional] 
**UpdateUserJSON** | Pointer to **string** | Property for UpdateUserJSON | [optional] 
**ChangePassJSON** | Pointer to **string** |  | [optional] 
**RemoveAccountJSON** | Pointer to **string** |  | [optional] 
**TicketStatusJSON** | Pointer to **string** |  | [optional] 
**CreateTicketJSON** | Pointer to **string** |  | [optional] 
**ENDPOINTS_FILTER** | Pointer to **string** |  | [optional] 
**PasswdPolicyJSON** | Pointer to **string** |  | [optional] 
**ConfigJSON** | Pointer to **string** | We can use this attribute to define the provisioningLimit,showLogs and connectionTimeoutConfig. | [optional] 
**AddFFIDAccessJSON** | Pointer to **string** |  | [optional] 
**RemoveFFIDAccessJSON** | Pointer to **string** |  | [optional] 
**MODIFYUSERDATAJSON** | Pointer to **string** |  | [optional] 
**SendOtpJSON** | Pointer to **string** |  | [optional] 
**ValidateOtpJSON** | Pointer to **string** |  | [optional] 
**PAM_CONFIG** | Pointer to **string** |  | [optional] 

## Methods

### NewRESTConnector

`func NewRESTConnector() *RESTConnector`

NewRESTConnector instantiates a new RESTConnector object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRESTConnectorWithDefaults

`func NewRESTConnectorWithDefaults() *RESTConnector`

NewRESTConnectorWithDefaults instantiates a new RESTConnector object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConnectionJSON

`func (o *RESTConnector) GetConnectionJSON() string`

GetConnectionJSON returns the ConnectionJSON field if non-nil, zero value otherwise.

### GetConnectionJSONOk

`func (o *RESTConnector) GetConnectionJSONOk() (*string, bool)`

GetConnectionJSONOk returns a tuple with the ConnectionJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionJSON

`func (o *RESTConnector) SetConnectionJSON(v string)`

SetConnectionJSON sets ConnectionJSON field to given value.

### HasConnectionJSON

`func (o *RESTConnector) HasConnectionJSON() bool`

HasConnectionJSON returns a boolean if a field has been set.

### GetImportUserJSON

`func (o *RESTConnector) GetImportUserJSON() string`

GetImportUserJSON returns the ImportUserJSON field if non-nil, zero value otherwise.

### GetImportUserJSONOk

`func (o *RESTConnector) GetImportUserJSONOk() (*string, bool)`

GetImportUserJSONOk returns a tuple with the ImportUserJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportUserJSON

`func (o *RESTConnector) SetImportUserJSON(v string)`

SetImportUserJSON sets ImportUserJSON field to given value.

### HasImportUserJSON

`func (o *RESTConnector) HasImportUserJSON() bool`

HasImportUserJSON returns a boolean if a field has been set.

### GetImportAccountEntJSON

`func (o *RESTConnector) GetImportAccountEntJSON() string`

GetImportAccountEntJSON returns the ImportAccountEntJSON field if non-nil, zero value otherwise.

### GetImportAccountEntJSONOk

`func (o *RESTConnector) GetImportAccountEntJSONOk() (*string, bool)`

GetImportAccountEntJSONOk returns a tuple with the ImportAccountEntJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportAccountEntJSON

`func (o *RESTConnector) SetImportAccountEntJSON(v string)`

SetImportAccountEntJSON sets ImportAccountEntJSON field to given value.

### HasImportAccountEntJSON

`func (o *RESTConnector) HasImportAccountEntJSON() bool`

HasImportAccountEntJSON returns a boolean if a field has been set.

### GetSTATUS_THRESHOLD_CONFIG

`func (o *RESTConnector) GetSTATUS_THRESHOLD_CONFIG() string`

GetSTATUS_THRESHOLD_CONFIG returns the STATUS_THRESHOLD_CONFIG field if non-nil, zero value otherwise.

### GetSTATUS_THRESHOLD_CONFIGOk

`func (o *RESTConnector) GetSTATUS_THRESHOLD_CONFIGOk() (*string, bool)`

GetSTATUS_THRESHOLD_CONFIGOk returns a tuple with the STATUS_THRESHOLD_CONFIG field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSTATUS_THRESHOLD_CONFIG

`func (o *RESTConnector) SetSTATUS_THRESHOLD_CONFIG(v string)`

SetSTATUS_THRESHOLD_CONFIG sets STATUS_THRESHOLD_CONFIG field to given value.

### HasSTATUS_THRESHOLD_CONFIG

`func (o *RESTConnector) HasSTATUS_THRESHOLD_CONFIG() bool`

HasSTATUS_THRESHOLD_CONFIG returns a boolean if a field has been set.

### GetCreateAccountJSON

`func (o *RESTConnector) GetCreateAccountJSON() string`

GetCreateAccountJSON returns the CreateAccountJSON field if non-nil, zero value otherwise.

### GetCreateAccountJSONOk

`func (o *RESTConnector) GetCreateAccountJSONOk() (*string, bool)`

GetCreateAccountJSONOk returns a tuple with the CreateAccountJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateAccountJSON

`func (o *RESTConnector) SetCreateAccountJSON(v string)`

SetCreateAccountJSON sets CreateAccountJSON field to given value.

### HasCreateAccountJSON

`func (o *RESTConnector) HasCreateAccountJSON() bool`

HasCreateAccountJSON returns a boolean if a field has been set.

### GetUpdateAccountJSON

`func (o *RESTConnector) GetUpdateAccountJSON() string`

GetUpdateAccountJSON returns the UpdateAccountJSON field if non-nil, zero value otherwise.

### GetUpdateAccountJSONOk

`func (o *RESTConnector) GetUpdateAccountJSONOk() (*string, bool)`

GetUpdateAccountJSONOk returns a tuple with the UpdateAccountJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateAccountJSON

`func (o *RESTConnector) SetUpdateAccountJSON(v string)`

SetUpdateAccountJSON sets UpdateAccountJSON field to given value.

### HasUpdateAccountJSON

`func (o *RESTConnector) HasUpdateAccountJSON() bool`

HasUpdateAccountJSON returns a boolean if a field has been set.

### GetEnableAccountJSON

`func (o *RESTConnector) GetEnableAccountJSON() string`

GetEnableAccountJSON returns the EnableAccountJSON field if non-nil, zero value otherwise.

### GetEnableAccountJSONOk

`func (o *RESTConnector) GetEnableAccountJSONOk() (*string, bool)`

GetEnableAccountJSONOk returns a tuple with the EnableAccountJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableAccountJSON

`func (o *RESTConnector) SetEnableAccountJSON(v string)`

SetEnableAccountJSON sets EnableAccountJSON field to given value.

### HasEnableAccountJSON

`func (o *RESTConnector) HasEnableAccountJSON() bool`

HasEnableAccountJSON returns a boolean if a field has been set.

### GetDisableAccountJSON

`func (o *RESTConnector) GetDisableAccountJSON() string`

GetDisableAccountJSON returns the DisableAccountJSON field if non-nil, zero value otherwise.

### GetDisableAccountJSONOk

`func (o *RESTConnector) GetDisableAccountJSONOk() (*string, bool)`

GetDisableAccountJSONOk returns a tuple with the DisableAccountJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableAccountJSON

`func (o *RESTConnector) SetDisableAccountJSON(v string)`

SetDisableAccountJSON sets DisableAccountJSON field to given value.

### HasDisableAccountJSON

`func (o *RESTConnector) HasDisableAccountJSON() bool`

HasDisableAccountJSON returns a boolean if a field has been set.

### GetAddAccessJSON

`func (o *RESTConnector) GetAddAccessJSON() string`

GetAddAccessJSON returns the AddAccessJSON field if non-nil, zero value otherwise.

### GetAddAccessJSONOk

`func (o *RESTConnector) GetAddAccessJSONOk() (*string, bool)`

GetAddAccessJSONOk returns a tuple with the AddAccessJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddAccessJSON

`func (o *RESTConnector) SetAddAccessJSON(v string)`

SetAddAccessJSON sets AddAccessJSON field to given value.

### HasAddAccessJSON

`func (o *RESTConnector) HasAddAccessJSON() bool`

HasAddAccessJSON returns a boolean if a field has been set.

### GetRemoveAccessJSON

`func (o *RESTConnector) GetRemoveAccessJSON() string`

GetRemoveAccessJSON returns the RemoveAccessJSON field if non-nil, zero value otherwise.

### GetRemoveAccessJSONOk

`func (o *RESTConnector) GetRemoveAccessJSONOk() (*string, bool)`

GetRemoveAccessJSONOk returns a tuple with the RemoveAccessJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoveAccessJSON

`func (o *RESTConnector) SetRemoveAccessJSON(v string)`

SetRemoveAccessJSON sets RemoveAccessJSON field to given value.

### HasRemoveAccessJSON

`func (o *RESTConnector) HasRemoveAccessJSON() bool`

HasRemoveAccessJSON returns a boolean if a field has been set.

### GetUpdateUserJSON

`func (o *RESTConnector) GetUpdateUserJSON() string`

GetUpdateUserJSON returns the UpdateUserJSON field if non-nil, zero value otherwise.

### GetUpdateUserJSONOk

`func (o *RESTConnector) GetUpdateUserJSONOk() (*string, bool)`

GetUpdateUserJSONOk returns a tuple with the UpdateUserJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateUserJSON

`func (o *RESTConnector) SetUpdateUserJSON(v string)`

SetUpdateUserJSON sets UpdateUserJSON field to given value.

### HasUpdateUserJSON

`func (o *RESTConnector) HasUpdateUserJSON() bool`

HasUpdateUserJSON returns a boolean if a field has been set.

### GetChangePassJSON

`func (o *RESTConnector) GetChangePassJSON() string`

GetChangePassJSON returns the ChangePassJSON field if non-nil, zero value otherwise.

### GetChangePassJSONOk

`func (o *RESTConnector) GetChangePassJSONOk() (*string, bool)`

GetChangePassJSONOk returns a tuple with the ChangePassJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangePassJSON

`func (o *RESTConnector) SetChangePassJSON(v string)`

SetChangePassJSON sets ChangePassJSON field to given value.

### HasChangePassJSON

`func (o *RESTConnector) HasChangePassJSON() bool`

HasChangePassJSON returns a boolean if a field has been set.

### GetRemoveAccountJSON

`func (o *RESTConnector) GetRemoveAccountJSON() string`

GetRemoveAccountJSON returns the RemoveAccountJSON field if non-nil, zero value otherwise.

### GetRemoveAccountJSONOk

`func (o *RESTConnector) GetRemoveAccountJSONOk() (*string, bool)`

GetRemoveAccountJSONOk returns a tuple with the RemoveAccountJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoveAccountJSON

`func (o *RESTConnector) SetRemoveAccountJSON(v string)`

SetRemoveAccountJSON sets RemoveAccountJSON field to given value.

### HasRemoveAccountJSON

`func (o *RESTConnector) HasRemoveAccountJSON() bool`

HasRemoveAccountJSON returns a boolean if a field has been set.

### GetTicketStatusJSON

`func (o *RESTConnector) GetTicketStatusJSON() string`

GetTicketStatusJSON returns the TicketStatusJSON field if non-nil, zero value otherwise.

### GetTicketStatusJSONOk

`func (o *RESTConnector) GetTicketStatusJSONOk() (*string, bool)`

GetTicketStatusJSONOk returns a tuple with the TicketStatusJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTicketStatusJSON

`func (o *RESTConnector) SetTicketStatusJSON(v string)`

SetTicketStatusJSON sets TicketStatusJSON field to given value.

### HasTicketStatusJSON

`func (o *RESTConnector) HasTicketStatusJSON() bool`

HasTicketStatusJSON returns a boolean if a field has been set.

### GetCreateTicketJSON

`func (o *RESTConnector) GetCreateTicketJSON() string`

GetCreateTicketJSON returns the CreateTicketJSON field if non-nil, zero value otherwise.

### GetCreateTicketJSONOk

`func (o *RESTConnector) GetCreateTicketJSONOk() (*string, bool)`

GetCreateTicketJSONOk returns a tuple with the CreateTicketJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTicketJSON

`func (o *RESTConnector) SetCreateTicketJSON(v string)`

SetCreateTicketJSON sets CreateTicketJSON field to given value.

### HasCreateTicketJSON

`func (o *RESTConnector) HasCreateTicketJSON() bool`

HasCreateTicketJSON returns a boolean if a field has been set.

### GetENDPOINTS_FILTER

`func (o *RESTConnector) GetENDPOINTS_FILTER() string`

GetENDPOINTS_FILTER returns the ENDPOINTS_FILTER field if non-nil, zero value otherwise.

### GetENDPOINTS_FILTEROk

`func (o *RESTConnector) GetENDPOINTS_FILTEROk() (*string, bool)`

GetENDPOINTS_FILTEROk returns a tuple with the ENDPOINTS_FILTER field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetENDPOINTS_FILTER

`func (o *RESTConnector) SetENDPOINTS_FILTER(v string)`

SetENDPOINTS_FILTER sets ENDPOINTS_FILTER field to given value.

### HasENDPOINTS_FILTER

`func (o *RESTConnector) HasENDPOINTS_FILTER() bool`

HasENDPOINTS_FILTER returns a boolean if a field has been set.

### GetPasswdPolicyJSON

`func (o *RESTConnector) GetPasswdPolicyJSON() string`

GetPasswdPolicyJSON returns the PasswdPolicyJSON field if non-nil, zero value otherwise.

### GetPasswdPolicyJSONOk

`func (o *RESTConnector) GetPasswdPolicyJSONOk() (*string, bool)`

GetPasswdPolicyJSONOk returns a tuple with the PasswdPolicyJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswdPolicyJSON

`func (o *RESTConnector) SetPasswdPolicyJSON(v string)`

SetPasswdPolicyJSON sets PasswdPolicyJSON field to given value.

### HasPasswdPolicyJSON

`func (o *RESTConnector) HasPasswdPolicyJSON() bool`

HasPasswdPolicyJSON returns a boolean if a field has been set.

### GetConfigJSON

`func (o *RESTConnector) GetConfigJSON() string`

GetConfigJSON returns the ConfigJSON field if non-nil, zero value otherwise.

### GetConfigJSONOk

`func (o *RESTConnector) GetConfigJSONOk() (*string, bool)`

GetConfigJSONOk returns a tuple with the ConfigJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigJSON

`func (o *RESTConnector) SetConfigJSON(v string)`

SetConfigJSON sets ConfigJSON field to given value.

### HasConfigJSON

`func (o *RESTConnector) HasConfigJSON() bool`

HasConfigJSON returns a boolean if a field has been set.

### GetAddFFIDAccessJSON

`func (o *RESTConnector) GetAddFFIDAccessJSON() string`

GetAddFFIDAccessJSON returns the AddFFIDAccessJSON field if non-nil, zero value otherwise.

### GetAddFFIDAccessJSONOk

`func (o *RESTConnector) GetAddFFIDAccessJSONOk() (*string, bool)`

GetAddFFIDAccessJSONOk returns a tuple with the AddFFIDAccessJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddFFIDAccessJSON

`func (o *RESTConnector) SetAddFFIDAccessJSON(v string)`

SetAddFFIDAccessJSON sets AddFFIDAccessJSON field to given value.

### HasAddFFIDAccessJSON

`func (o *RESTConnector) HasAddFFIDAccessJSON() bool`

HasAddFFIDAccessJSON returns a boolean if a field has been set.

### GetRemoveFFIDAccessJSON

`func (o *RESTConnector) GetRemoveFFIDAccessJSON() string`

GetRemoveFFIDAccessJSON returns the RemoveFFIDAccessJSON field if non-nil, zero value otherwise.

### GetRemoveFFIDAccessJSONOk

`func (o *RESTConnector) GetRemoveFFIDAccessJSONOk() (*string, bool)`

GetRemoveFFIDAccessJSONOk returns a tuple with the RemoveFFIDAccessJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoveFFIDAccessJSON

`func (o *RESTConnector) SetRemoveFFIDAccessJSON(v string)`

SetRemoveFFIDAccessJSON sets RemoveFFIDAccessJSON field to given value.

### HasRemoveFFIDAccessJSON

`func (o *RESTConnector) HasRemoveFFIDAccessJSON() bool`

HasRemoveFFIDAccessJSON returns a boolean if a field has been set.

### GetMODIFYUSERDATAJSON

`func (o *RESTConnector) GetMODIFYUSERDATAJSON() string`

GetMODIFYUSERDATAJSON returns the MODIFYUSERDATAJSON field if non-nil, zero value otherwise.

### GetMODIFYUSERDATAJSONOk

`func (o *RESTConnector) GetMODIFYUSERDATAJSONOk() (*string, bool)`

GetMODIFYUSERDATAJSONOk returns a tuple with the MODIFYUSERDATAJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMODIFYUSERDATAJSON

`func (o *RESTConnector) SetMODIFYUSERDATAJSON(v string)`

SetMODIFYUSERDATAJSON sets MODIFYUSERDATAJSON field to given value.

### HasMODIFYUSERDATAJSON

`func (o *RESTConnector) HasMODIFYUSERDATAJSON() bool`

HasMODIFYUSERDATAJSON returns a boolean if a field has been set.

### GetSendOtpJSON

`func (o *RESTConnector) GetSendOtpJSON() string`

GetSendOtpJSON returns the SendOtpJSON field if non-nil, zero value otherwise.

### GetSendOtpJSONOk

`func (o *RESTConnector) GetSendOtpJSONOk() (*string, bool)`

GetSendOtpJSONOk returns a tuple with the SendOtpJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendOtpJSON

`func (o *RESTConnector) SetSendOtpJSON(v string)`

SetSendOtpJSON sets SendOtpJSON field to given value.

### HasSendOtpJSON

`func (o *RESTConnector) HasSendOtpJSON() bool`

HasSendOtpJSON returns a boolean if a field has been set.

### GetValidateOtpJSON

`func (o *RESTConnector) GetValidateOtpJSON() string`

GetValidateOtpJSON returns the ValidateOtpJSON field if non-nil, zero value otherwise.

### GetValidateOtpJSONOk

`func (o *RESTConnector) GetValidateOtpJSONOk() (*string, bool)`

GetValidateOtpJSONOk returns a tuple with the ValidateOtpJSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidateOtpJSON

`func (o *RESTConnector) SetValidateOtpJSON(v string)`

SetValidateOtpJSON sets ValidateOtpJSON field to given value.

### HasValidateOtpJSON

`func (o *RESTConnector) HasValidateOtpJSON() bool`

HasValidateOtpJSON returns a boolean if a field has been set.

### GetPAM_CONFIG

`func (o *RESTConnector) GetPAM_CONFIG() string`

GetPAM_CONFIG returns the PAM_CONFIG field if non-nil, zero value otherwise.

### GetPAM_CONFIGOk

`func (o *RESTConnector) GetPAM_CONFIGOk() (*string, bool)`

GetPAM_CONFIGOk returns a tuple with the PAM_CONFIG field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPAM_CONFIG

`func (o *RESTConnector) SetPAM_CONFIG(v string)`

SetPAM_CONFIG sets PAM_CONFIG field to given value.

### HasPAM_CONFIG

`func (o *RESTConnector) HasPAM_CONFIG() bool`

HasPAM_CONFIG returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


