# CreateSecuritySystemRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Systemname** | **string** | Name of the security system. | 
**DisplayName** | **string** | Specify a user-friendly display name that is shown on the the user interface. | 
**Hostname** | Pointer to **string** | Security system for which you want to create an endpoint. | [optional] 
**Port** | Pointer to **string** | Description for the endpoint. | [optional] 
**AccessAddWorkflow** | Pointer to **string** | Specify the workflow to be used for approvals for an access request, which can be for an account, entitlements, role, and so on | [optional] 
**AccessRemoveWorkflow** | Pointer to **string** | Specify the workflow to be used when access has to be revoked, which can be for an account, entitlement, or any other de-provisioning task. | [optional] 
**AddServiceAccountWorkflow** | Pointer to **string** | Workflow for adding a service account. | [optional] 
**RemoveServiceAccountWorkflow** | Pointer to **string** | Workflow for removing a service account. | [optional] 
**Connectionparameters** | Pointer to **string** | Query to filter the access and display of the endpoint to specific users. If you do not define a query, the endpoint is displayed for all users | [optional] 
**AutomatedProvisioning** | Pointer to **string** | Specify true to enable automated provisioning. | [optional] 
**Useopenconnector** | Pointer to **bool** | Specify true to enable the connectivity with any system over the open-source connectors such as REST. | [optional] 
**ReconApplication** | Pointer to **string** | Specify true to importing data from respective endpoint associated to security system. | [optional] 
**ProvisioningTries** | Pointer to **string** | Number of provisioning tries. | [optional] 
**Instantprovision** | Pointer to **string** | Specify true to prevent users from raising duplicate requests for the same applications. | [optional] 
**Provisioningcomments** | Pointer to **string** | Specify relevant comments for performing provisioning. | [optional] 

## Methods

### NewCreateSecuritySystemRequest

`func NewCreateSecuritySystemRequest(systemname string, displayName string, ) *CreateSecuritySystemRequest`

NewCreateSecuritySystemRequest instantiates a new CreateSecuritySystemRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateSecuritySystemRequestWithDefaults

`func NewCreateSecuritySystemRequestWithDefaults() *CreateSecuritySystemRequest`

NewCreateSecuritySystemRequestWithDefaults instantiates a new CreateSecuritySystemRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSystemname

`func (o *CreateSecuritySystemRequest) GetSystemname() string`

GetSystemname returns the Systemname field if non-nil, zero value otherwise.

### GetSystemnameOk

`func (o *CreateSecuritySystemRequest) GetSystemnameOk() (*string, bool)`

GetSystemnameOk returns a tuple with the Systemname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemname

`func (o *CreateSecuritySystemRequest) SetSystemname(v string)`

SetSystemname sets Systemname field to given value.


### GetDisplayName

`func (o *CreateSecuritySystemRequest) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *CreateSecuritySystemRequest) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *CreateSecuritySystemRequest) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetHostname

`func (o *CreateSecuritySystemRequest) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *CreateSecuritySystemRequest) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *CreateSecuritySystemRequest) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *CreateSecuritySystemRequest) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### GetPort

`func (o *CreateSecuritySystemRequest) GetPort() string`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *CreateSecuritySystemRequest) GetPortOk() (*string, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *CreateSecuritySystemRequest) SetPort(v string)`

SetPort sets Port field to given value.

### HasPort

`func (o *CreateSecuritySystemRequest) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetAccessAddWorkflow

`func (o *CreateSecuritySystemRequest) GetAccessAddWorkflow() string`

GetAccessAddWorkflow returns the AccessAddWorkflow field if non-nil, zero value otherwise.

### GetAccessAddWorkflowOk

`func (o *CreateSecuritySystemRequest) GetAccessAddWorkflowOk() (*string, bool)`

GetAccessAddWorkflowOk returns a tuple with the AccessAddWorkflow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessAddWorkflow

`func (o *CreateSecuritySystemRequest) SetAccessAddWorkflow(v string)`

SetAccessAddWorkflow sets AccessAddWorkflow field to given value.

### HasAccessAddWorkflow

`func (o *CreateSecuritySystemRequest) HasAccessAddWorkflow() bool`

HasAccessAddWorkflow returns a boolean if a field has been set.

### GetAccessRemoveWorkflow

`func (o *CreateSecuritySystemRequest) GetAccessRemoveWorkflow() string`

GetAccessRemoveWorkflow returns the AccessRemoveWorkflow field if non-nil, zero value otherwise.

### GetAccessRemoveWorkflowOk

`func (o *CreateSecuritySystemRequest) GetAccessRemoveWorkflowOk() (*string, bool)`

GetAccessRemoveWorkflowOk returns a tuple with the AccessRemoveWorkflow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessRemoveWorkflow

`func (o *CreateSecuritySystemRequest) SetAccessRemoveWorkflow(v string)`

SetAccessRemoveWorkflow sets AccessRemoveWorkflow field to given value.

### HasAccessRemoveWorkflow

`func (o *CreateSecuritySystemRequest) HasAccessRemoveWorkflow() bool`

HasAccessRemoveWorkflow returns a boolean if a field has been set.

### GetAddServiceAccountWorkflow

`func (o *CreateSecuritySystemRequest) GetAddServiceAccountWorkflow() string`

GetAddServiceAccountWorkflow returns the AddServiceAccountWorkflow field if non-nil, zero value otherwise.

### GetAddServiceAccountWorkflowOk

`func (o *CreateSecuritySystemRequest) GetAddServiceAccountWorkflowOk() (*string, bool)`

GetAddServiceAccountWorkflowOk returns a tuple with the AddServiceAccountWorkflow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddServiceAccountWorkflow

`func (o *CreateSecuritySystemRequest) SetAddServiceAccountWorkflow(v string)`

SetAddServiceAccountWorkflow sets AddServiceAccountWorkflow field to given value.

### HasAddServiceAccountWorkflow

`func (o *CreateSecuritySystemRequest) HasAddServiceAccountWorkflow() bool`

HasAddServiceAccountWorkflow returns a boolean if a field has been set.

### GetRemoveServiceAccountWorkflow

`func (o *CreateSecuritySystemRequest) GetRemoveServiceAccountWorkflow() string`

GetRemoveServiceAccountWorkflow returns the RemoveServiceAccountWorkflow field if non-nil, zero value otherwise.

### GetRemoveServiceAccountWorkflowOk

`func (o *CreateSecuritySystemRequest) GetRemoveServiceAccountWorkflowOk() (*string, bool)`

GetRemoveServiceAccountWorkflowOk returns a tuple with the RemoveServiceAccountWorkflow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoveServiceAccountWorkflow

`func (o *CreateSecuritySystemRequest) SetRemoveServiceAccountWorkflow(v string)`

SetRemoveServiceAccountWorkflow sets RemoveServiceAccountWorkflow field to given value.

### HasRemoveServiceAccountWorkflow

`func (o *CreateSecuritySystemRequest) HasRemoveServiceAccountWorkflow() bool`

HasRemoveServiceAccountWorkflow returns a boolean if a field has been set.

### GetConnectionparameters

`func (o *CreateSecuritySystemRequest) GetConnectionparameters() string`

GetConnectionparameters returns the Connectionparameters field if non-nil, zero value otherwise.

### GetConnectionparametersOk

`func (o *CreateSecuritySystemRequest) GetConnectionparametersOk() (*string, bool)`

GetConnectionparametersOk returns a tuple with the Connectionparameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionparameters

`func (o *CreateSecuritySystemRequest) SetConnectionparameters(v string)`

SetConnectionparameters sets Connectionparameters field to given value.

### HasConnectionparameters

`func (o *CreateSecuritySystemRequest) HasConnectionparameters() bool`

HasConnectionparameters returns a boolean if a field has been set.

### GetAutomatedProvisioning

`func (o *CreateSecuritySystemRequest) GetAutomatedProvisioning() string`

GetAutomatedProvisioning returns the AutomatedProvisioning field if non-nil, zero value otherwise.

### GetAutomatedProvisioningOk

`func (o *CreateSecuritySystemRequest) GetAutomatedProvisioningOk() (*string, bool)`

GetAutomatedProvisioningOk returns a tuple with the AutomatedProvisioning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutomatedProvisioning

`func (o *CreateSecuritySystemRequest) SetAutomatedProvisioning(v string)`

SetAutomatedProvisioning sets AutomatedProvisioning field to given value.

### HasAutomatedProvisioning

`func (o *CreateSecuritySystemRequest) HasAutomatedProvisioning() bool`

HasAutomatedProvisioning returns a boolean if a field has been set.

### GetUseopenconnector

`func (o *CreateSecuritySystemRequest) GetUseopenconnector() bool`

GetUseopenconnector returns the Useopenconnector field if non-nil, zero value otherwise.

### GetUseopenconnectorOk

`func (o *CreateSecuritySystemRequest) GetUseopenconnectorOk() (*bool, bool)`

GetUseopenconnectorOk returns a tuple with the Useopenconnector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseopenconnector

`func (o *CreateSecuritySystemRequest) SetUseopenconnector(v bool)`

SetUseopenconnector sets Useopenconnector field to given value.

### HasUseopenconnector

`func (o *CreateSecuritySystemRequest) HasUseopenconnector() bool`

HasUseopenconnector returns a boolean if a field has been set.

### GetReconApplication

`func (o *CreateSecuritySystemRequest) GetReconApplication() string`

GetReconApplication returns the ReconApplication field if non-nil, zero value otherwise.

### GetReconApplicationOk

`func (o *CreateSecuritySystemRequest) GetReconApplicationOk() (*string, bool)`

GetReconApplicationOk returns a tuple with the ReconApplication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReconApplication

`func (o *CreateSecuritySystemRequest) SetReconApplication(v string)`

SetReconApplication sets ReconApplication field to given value.

### HasReconApplication

`func (o *CreateSecuritySystemRequest) HasReconApplication() bool`

HasReconApplication returns a boolean if a field has been set.

### GetProvisioningTries

`func (o *CreateSecuritySystemRequest) GetProvisioningTries() string`

GetProvisioningTries returns the ProvisioningTries field if non-nil, zero value otherwise.

### GetProvisioningTriesOk

`func (o *CreateSecuritySystemRequest) GetProvisioningTriesOk() (*string, bool)`

GetProvisioningTriesOk returns a tuple with the ProvisioningTries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisioningTries

`func (o *CreateSecuritySystemRequest) SetProvisioningTries(v string)`

SetProvisioningTries sets ProvisioningTries field to given value.

### HasProvisioningTries

`func (o *CreateSecuritySystemRequest) HasProvisioningTries() bool`

HasProvisioningTries returns a boolean if a field has been set.

### GetInstantprovision

`func (o *CreateSecuritySystemRequest) GetInstantprovision() string`

GetInstantprovision returns the Instantprovision field if non-nil, zero value otherwise.

### GetInstantprovisionOk

`func (o *CreateSecuritySystemRequest) GetInstantprovisionOk() (*string, bool)`

GetInstantprovisionOk returns a tuple with the Instantprovision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstantprovision

`func (o *CreateSecuritySystemRequest) SetInstantprovision(v string)`

SetInstantprovision sets Instantprovision field to given value.

### HasInstantprovision

`func (o *CreateSecuritySystemRequest) HasInstantprovision() bool`

HasInstantprovision returns a boolean if a field has been set.

### GetProvisioningcomments

`func (o *CreateSecuritySystemRequest) GetProvisioningcomments() string`

GetProvisioningcomments returns the Provisioningcomments field if non-nil, zero value otherwise.

### GetProvisioningcommentsOk

`func (o *CreateSecuritySystemRequest) GetProvisioningcommentsOk() (*string, bool)`

GetProvisioningcommentsOk returns a tuple with the Provisioningcomments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisioningcomments

`func (o *CreateSecuritySystemRequest) SetProvisioningcomments(v string)`

SetProvisioningcomments sets Provisioningcomments field to given value.

### HasProvisioningcomments

`func (o *CreateSecuritySystemRequest) HasProvisioningcomments() bool`

HasProvisioningcomments returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


