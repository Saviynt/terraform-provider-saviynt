# UpdateSecuritySystemRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Systemname** | **string** | The security system name to update. | 
**DisplayName** | **string** | Specify a user-friendly display name that is shown on the the user interface. | 
**Hostname** | Pointer to **string** | Security system for which you want to create an endpoint. | [optional] 
**Port** | Pointer to **string** | Description for the endpoint. | [optional] 
**AccessAddWorkflow** | Pointer to **string** | Specify the workflow to be used for approvals for an access request, which can be for an account, entitlements, role, and so on | [optional] 
**AccessRemoveWorkflow** | Pointer to **string** | Specify the workflow to be used when access has to be revoked, which can be for an account, entitlement, or any other de-provisioning task. | [optional] 
**AddServiceAccountWorkflow** | Pointer to **string** | Workflow for adding a service account. | [optional] 
**RemoveServiceAccountWorkflow** | Pointer to **string** | Workflow for removing a service account. | [optional] 
**Connectionparameters** | Pointer to **string** | Query to filter the access and display of the endpoint to specific users. If you do not define a query, the endpoint is displayed for all users | [optional] 
**ProposedAccountOwnersworkflow** | Pointer to **string** | Query to filter the access and display of the endpoint to specific users. If you do not define a query, the endpoint is displayed for all users | [optional] 
**FirefighteridWorkflow** | Pointer to **string** | firefighteridWorkflow | [optional] 
**FirefighteridRequestAccessWorkflow** | Pointer to **string** | firefighteridRequestAccessWorkflow | [optional] 
**AutomatedProvisioning** | Pointer to **string** | Specify true to enable automated provisioning. | [optional] 
**Useopenconnector** | Pointer to **bool** | Specify true to enable the connectivity with any system over the open-source connectors such as REST. | [optional] 
**DefaultSystem** | Pointer to **string** | Specify true to set the security system as the Default System. Following which accounts search will only be searched and displayed for this security system. | [optional] 
**ReconApplication** | Pointer to **string** | Specify true to importing data from respective endpoint associated to security system. | [optional] 
**Instantprovision** | Pointer to **string** | Use this flag to prevent users from raising duplicate requests for the same applications. | [optional] 
**ProvisioningTries** | Pointer to **string** | Specify the number of tries to be used for performing provisioning / de-provisioning to the third-party application. You can specify provisioningTries between 1 to 20 based on your requirement. | [optional] 
**ExternalRiskConnectionJson** | Pointer to **string** | Contains JSON configuration for external risk connections and is applicable only for few connections like SAP | [optional] 
**PolicyRule** | Pointer to **string** | Use this setting to assign the password policy for the security system. | [optional] 
**PolicyRuleServiceAccount** | Pointer to **string** | Use this setting to assign the password policy which will be used to set the service account passwords for the security system. | [optional] 
**Connectionname** | Pointer to **string** | Select the connection name for performing reconciliation of identity objects from third-party application. | [optional] 
**ProvisioningConnection** | Pointer to **string** | You can use a separate connection to an endpoint where you are performing provisioning or deprovisioning. Based on your requirement, you can specify a separate connection where you want to perform provisioning and de-provisioning. | [optional] 
**ServiceDeskConnection** | Pointer to **string** | Specify the Service Desk Connection used for integration with a ticketing system, which can be a disconnected system too. | [optional] 
**Provisioningcomments** | Pointer to **string** | Specify relevant comments for performing provisioning. | [optional] 
**InherentSODReportFields** | Pointer to **[]string** |  You can use this option used to filter out columns in SOD. | [optional] 

## Methods

### NewUpdateSecuritySystemRequest

`func NewUpdateSecuritySystemRequest(systemname string, displayName string, ) *UpdateSecuritySystemRequest`

NewUpdateSecuritySystemRequest instantiates a new UpdateSecuritySystemRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateSecuritySystemRequestWithDefaults

`func NewUpdateSecuritySystemRequestWithDefaults() *UpdateSecuritySystemRequest`

NewUpdateSecuritySystemRequestWithDefaults instantiates a new UpdateSecuritySystemRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSystemname

`func (o *UpdateSecuritySystemRequest) GetSystemname() string`

GetSystemname returns the Systemname field if non-nil, zero value otherwise.

### GetSystemnameOk

`func (o *UpdateSecuritySystemRequest) GetSystemnameOk() (*string, bool)`

GetSystemnameOk returns a tuple with the Systemname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemname

`func (o *UpdateSecuritySystemRequest) SetSystemname(v string)`

SetSystemname sets Systemname field to given value.


### GetDisplayName

`func (o *UpdateSecuritySystemRequest) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *UpdateSecuritySystemRequest) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *UpdateSecuritySystemRequest) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetHostname

`func (o *UpdateSecuritySystemRequest) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *UpdateSecuritySystemRequest) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *UpdateSecuritySystemRequest) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *UpdateSecuritySystemRequest) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### GetPort

`func (o *UpdateSecuritySystemRequest) GetPort() string`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *UpdateSecuritySystemRequest) GetPortOk() (*string, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *UpdateSecuritySystemRequest) SetPort(v string)`

SetPort sets Port field to given value.

### HasPort

`func (o *UpdateSecuritySystemRequest) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetAccessAddWorkflow

`func (o *UpdateSecuritySystemRequest) GetAccessAddWorkflow() string`

GetAccessAddWorkflow returns the AccessAddWorkflow field if non-nil, zero value otherwise.

### GetAccessAddWorkflowOk

`func (o *UpdateSecuritySystemRequest) GetAccessAddWorkflowOk() (*string, bool)`

GetAccessAddWorkflowOk returns a tuple with the AccessAddWorkflow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessAddWorkflow

`func (o *UpdateSecuritySystemRequest) SetAccessAddWorkflow(v string)`

SetAccessAddWorkflow sets AccessAddWorkflow field to given value.

### HasAccessAddWorkflow

`func (o *UpdateSecuritySystemRequest) HasAccessAddWorkflow() bool`

HasAccessAddWorkflow returns a boolean if a field has been set.

### GetAccessRemoveWorkflow

`func (o *UpdateSecuritySystemRequest) GetAccessRemoveWorkflow() string`

GetAccessRemoveWorkflow returns the AccessRemoveWorkflow field if non-nil, zero value otherwise.

### GetAccessRemoveWorkflowOk

`func (o *UpdateSecuritySystemRequest) GetAccessRemoveWorkflowOk() (*string, bool)`

GetAccessRemoveWorkflowOk returns a tuple with the AccessRemoveWorkflow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessRemoveWorkflow

`func (o *UpdateSecuritySystemRequest) SetAccessRemoveWorkflow(v string)`

SetAccessRemoveWorkflow sets AccessRemoveWorkflow field to given value.

### HasAccessRemoveWorkflow

`func (o *UpdateSecuritySystemRequest) HasAccessRemoveWorkflow() bool`

HasAccessRemoveWorkflow returns a boolean if a field has been set.

### GetAddServiceAccountWorkflow

`func (o *UpdateSecuritySystemRequest) GetAddServiceAccountWorkflow() string`

GetAddServiceAccountWorkflow returns the AddServiceAccountWorkflow field if non-nil, zero value otherwise.

### GetAddServiceAccountWorkflowOk

`func (o *UpdateSecuritySystemRequest) GetAddServiceAccountWorkflowOk() (*string, bool)`

GetAddServiceAccountWorkflowOk returns a tuple with the AddServiceAccountWorkflow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddServiceAccountWorkflow

`func (o *UpdateSecuritySystemRequest) SetAddServiceAccountWorkflow(v string)`

SetAddServiceAccountWorkflow sets AddServiceAccountWorkflow field to given value.

### HasAddServiceAccountWorkflow

`func (o *UpdateSecuritySystemRequest) HasAddServiceAccountWorkflow() bool`

HasAddServiceAccountWorkflow returns a boolean if a field has been set.

### GetRemoveServiceAccountWorkflow

`func (o *UpdateSecuritySystemRequest) GetRemoveServiceAccountWorkflow() string`

GetRemoveServiceAccountWorkflow returns the RemoveServiceAccountWorkflow field if non-nil, zero value otherwise.

### GetRemoveServiceAccountWorkflowOk

`func (o *UpdateSecuritySystemRequest) GetRemoveServiceAccountWorkflowOk() (*string, bool)`

GetRemoveServiceAccountWorkflowOk returns a tuple with the RemoveServiceAccountWorkflow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoveServiceAccountWorkflow

`func (o *UpdateSecuritySystemRequest) SetRemoveServiceAccountWorkflow(v string)`

SetRemoveServiceAccountWorkflow sets RemoveServiceAccountWorkflow field to given value.

### HasRemoveServiceAccountWorkflow

`func (o *UpdateSecuritySystemRequest) HasRemoveServiceAccountWorkflow() bool`

HasRemoveServiceAccountWorkflow returns a boolean if a field has been set.

### GetConnectionparameters

`func (o *UpdateSecuritySystemRequest) GetConnectionparameters() string`

GetConnectionparameters returns the Connectionparameters field if non-nil, zero value otherwise.

### GetConnectionparametersOk

`func (o *UpdateSecuritySystemRequest) GetConnectionparametersOk() (*string, bool)`

GetConnectionparametersOk returns a tuple with the Connectionparameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionparameters

`func (o *UpdateSecuritySystemRequest) SetConnectionparameters(v string)`

SetConnectionparameters sets Connectionparameters field to given value.

### HasConnectionparameters

`func (o *UpdateSecuritySystemRequest) HasConnectionparameters() bool`

HasConnectionparameters returns a boolean if a field has been set.

### GetProposedAccountOwnersworkflow

`func (o *UpdateSecuritySystemRequest) GetProposedAccountOwnersworkflow() string`

GetProposedAccountOwnersworkflow returns the ProposedAccountOwnersworkflow field if non-nil, zero value otherwise.

### GetProposedAccountOwnersworkflowOk

`func (o *UpdateSecuritySystemRequest) GetProposedAccountOwnersworkflowOk() (*string, bool)`

GetProposedAccountOwnersworkflowOk returns a tuple with the ProposedAccountOwnersworkflow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProposedAccountOwnersworkflow

`func (o *UpdateSecuritySystemRequest) SetProposedAccountOwnersworkflow(v string)`

SetProposedAccountOwnersworkflow sets ProposedAccountOwnersworkflow field to given value.

### HasProposedAccountOwnersworkflow

`func (o *UpdateSecuritySystemRequest) HasProposedAccountOwnersworkflow() bool`

HasProposedAccountOwnersworkflow returns a boolean if a field has been set.

### GetFirefighteridWorkflow

`func (o *UpdateSecuritySystemRequest) GetFirefighteridWorkflow() string`

GetFirefighteridWorkflow returns the FirefighteridWorkflow field if non-nil, zero value otherwise.

### GetFirefighteridWorkflowOk

`func (o *UpdateSecuritySystemRequest) GetFirefighteridWorkflowOk() (*string, bool)`

GetFirefighteridWorkflowOk returns a tuple with the FirefighteridWorkflow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirefighteridWorkflow

`func (o *UpdateSecuritySystemRequest) SetFirefighteridWorkflow(v string)`

SetFirefighteridWorkflow sets FirefighteridWorkflow field to given value.

### HasFirefighteridWorkflow

`func (o *UpdateSecuritySystemRequest) HasFirefighteridWorkflow() bool`

HasFirefighteridWorkflow returns a boolean if a field has been set.

### GetFirefighteridRequestAccessWorkflow

`func (o *UpdateSecuritySystemRequest) GetFirefighteridRequestAccessWorkflow() string`

GetFirefighteridRequestAccessWorkflow returns the FirefighteridRequestAccessWorkflow field if non-nil, zero value otherwise.

### GetFirefighteridRequestAccessWorkflowOk

`func (o *UpdateSecuritySystemRequest) GetFirefighteridRequestAccessWorkflowOk() (*string, bool)`

GetFirefighteridRequestAccessWorkflowOk returns a tuple with the FirefighteridRequestAccessWorkflow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirefighteridRequestAccessWorkflow

`func (o *UpdateSecuritySystemRequest) SetFirefighteridRequestAccessWorkflow(v string)`

SetFirefighteridRequestAccessWorkflow sets FirefighteridRequestAccessWorkflow field to given value.

### HasFirefighteridRequestAccessWorkflow

`func (o *UpdateSecuritySystemRequest) HasFirefighteridRequestAccessWorkflow() bool`

HasFirefighteridRequestAccessWorkflow returns a boolean if a field has been set.

### GetAutomatedProvisioning

`func (o *UpdateSecuritySystemRequest) GetAutomatedProvisioning() string`

GetAutomatedProvisioning returns the AutomatedProvisioning field if non-nil, zero value otherwise.

### GetAutomatedProvisioningOk

`func (o *UpdateSecuritySystemRequest) GetAutomatedProvisioningOk() (*string, bool)`

GetAutomatedProvisioningOk returns a tuple with the AutomatedProvisioning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutomatedProvisioning

`func (o *UpdateSecuritySystemRequest) SetAutomatedProvisioning(v string)`

SetAutomatedProvisioning sets AutomatedProvisioning field to given value.

### HasAutomatedProvisioning

`func (o *UpdateSecuritySystemRequest) HasAutomatedProvisioning() bool`

HasAutomatedProvisioning returns a boolean if a field has been set.

### GetUseopenconnector

`func (o *UpdateSecuritySystemRequest) GetUseopenconnector() bool`

GetUseopenconnector returns the Useopenconnector field if non-nil, zero value otherwise.

### GetUseopenconnectorOk

`func (o *UpdateSecuritySystemRequest) GetUseopenconnectorOk() (*bool, bool)`

GetUseopenconnectorOk returns a tuple with the Useopenconnector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseopenconnector

`func (o *UpdateSecuritySystemRequest) SetUseopenconnector(v bool)`

SetUseopenconnector sets Useopenconnector field to given value.

### HasUseopenconnector

`func (o *UpdateSecuritySystemRequest) HasUseopenconnector() bool`

HasUseopenconnector returns a boolean if a field has been set.

### GetDefaultSystem

`func (o *UpdateSecuritySystemRequest) GetDefaultSystem() string`

GetDefaultSystem returns the DefaultSystem field if non-nil, zero value otherwise.

### GetDefaultSystemOk

`func (o *UpdateSecuritySystemRequest) GetDefaultSystemOk() (*string, bool)`

GetDefaultSystemOk returns a tuple with the DefaultSystem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultSystem

`func (o *UpdateSecuritySystemRequest) SetDefaultSystem(v string)`

SetDefaultSystem sets DefaultSystem field to given value.

### HasDefaultSystem

`func (o *UpdateSecuritySystemRequest) HasDefaultSystem() bool`

HasDefaultSystem returns a boolean if a field has been set.

### GetReconApplication

`func (o *UpdateSecuritySystemRequest) GetReconApplication() string`

GetReconApplication returns the ReconApplication field if non-nil, zero value otherwise.

### GetReconApplicationOk

`func (o *UpdateSecuritySystemRequest) GetReconApplicationOk() (*string, bool)`

GetReconApplicationOk returns a tuple with the ReconApplication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReconApplication

`func (o *UpdateSecuritySystemRequest) SetReconApplication(v string)`

SetReconApplication sets ReconApplication field to given value.

### HasReconApplication

`func (o *UpdateSecuritySystemRequest) HasReconApplication() bool`

HasReconApplication returns a boolean if a field has been set.

### GetInstantprovision

`func (o *UpdateSecuritySystemRequest) GetInstantprovision() string`

GetInstantprovision returns the Instantprovision field if non-nil, zero value otherwise.

### GetInstantprovisionOk

`func (o *UpdateSecuritySystemRequest) GetInstantprovisionOk() (*string, bool)`

GetInstantprovisionOk returns a tuple with the Instantprovision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstantprovision

`func (o *UpdateSecuritySystemRequest) SetInstantprovision(v string)`

SetInstantprovision sets Instantprovision field to given value.

### HasInstantprovision

`func (o *UpdateSecuritySystemRequest) HasInstantprovision() bool`

HasInstantprovision returns a boolean if a field has been set.

### GetProvisioningTries

`func (o *UpdateSecuritySystemRequest) GetProvisioningTries() string`

GetProvisioningTries returns the ProvisioningTries field if non-nil, zero value otherwise.

### GetProvisioningTriesOk

`func (o *UpdateSecuritySystemRequest) GetProvisioningTriesOk() (*string, bool)`

GetProvisioningTriesOk returns a tuple with the ProvisioningTries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisioningTries

`func (o *UpdateSecuritySystemRequest) SetProvisioningTries(v string)`

SetProvisioningTries sets ProvisioningTries field to given value.

### HasProvisioningTries

`func (o *UpdateSecuritySystemRequest) HasProvisioningTries() bool`

HasProvisioningTries returns a boolean if a field has been set.

### GetExternalRiskConnectionJson

`func (o *UpdateSecuritySystemRequest) GetExternalRiskConnectionJson() string`

GetExternalRiskConnectionJson returns the ExternalRiskConnectionJson field if non-nil, zero value otherwise.

### GetExternalRiskConnectionJsonOk

`func (o *UpdateSecuritySystemRequest) GetExternalRiskConnectionJsonOk() (*string, bool)`

GetExternalRiskConnectionJsonOk returns a tuple with the ExternalRiskConnectionJson field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalRiskConnectionJson

`func (o *UpdateSecuritySystemRequest) SetExternalRiskConnectionJson(v string)`

SetExternalRiskConnectionJson sets ExternalRiskConnectionJson field to given value.

### HasExternalRiskConnectionJson

`func (o *UpdateSecuritySystemRequest) HasExternalRiskConnectionJson() bool`

HasExternalRiskConnectionJson returns a boolean if a field has been set.

### GetPolicyRule

`func (o *UpdateSecuritySystemRequest) GetPolicyRule() string`

GetPolicyRule returns the PolicyRule field if non-nil, zero value otherwise.

### GetPolicyRuleOk

`func (o *UpdateSecuritySystemRequest) GetPolicyRuleOk() (*string, bool)`

GetPolicyRuleOk returns a tuple with the PolicyRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyRule

`func (o *UpdateSecuritySystemRequest) SetPolicyRule(v string)`

SetPolicyRule sets PolicyRule field to given value.

### HasPolicyRule

`func (o *UpdateSecuritySystemRequest) HasPolicyRule() bool`

HasPolicyRule returns a boolean if a field has been set.

### GetPolicyRuleServiceAccount

`func (o *UpdateSecuritySystemRequest) GetPolicyRuleServiceAccount() string`

GetPolicyRuleServiceAccount returns the PolicyRuleServiceAccount field if non-nil, zero value otherwise.

### GetPolicyRuleServiceAccountOk

`func (o *UpdateSecuritySystemRequest) GetPolicyRuleServiceAccountOk() (*string, bool)`

GetPolicyRuleServiceAccountOk returns a tuple with the PolicyRuleServiceAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyRuleServiceAccount

`func (o *UpdateSecuritySystemRequest) SetPolicyRuleServiceAccount(v string)`

SetPolicyRuleServiceAccount sets PolicyRuleServiceAccount field to given value.

### HasPolicyRuleServiceAccount

`func (o *UpdateSecuritySystemRequest) HasPolicyRuleServiceAccount() bool`

HasPolicyRuleServiceAccount returns a boolean if a field has been set.

### GetConnectionname

`func (o *UpdateSecuritySystemRequest) GetConnectionname() string`

GetConnectionname returns the Connectionname field if non-nil, zero value otherwise.

### GetConnectionnameOk

`func (o *UpdateSecuritySystemRequest) GetConnectionnameOk() (*string, bool)`

GetConnectionnameOk returns a tuple with the Connectionname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionname

`func (o *UpdateSecuritySystemRequest) SetConnectionname(v string)`

SetConnectionname sets Connectionname field to given value.

### HasConnectionname

`func (o *UpdateSecuritySystemRequest) HasConnectionname() bool`

HasConnectionname returns a boolean if a field has been set.

### GetProvisioningConnection

`func (o *UpdateSecuritySystemRequest) GetProvisioningConnection() string`

GetProvisioningConnection returns the ProvisioningConnection field if non-nil, zero value otherwise.

### GetProvisioningConnectionOk

`func (o *UpdateSecuritySystemRequest) GetProvisioningConnectionOk() (*string, bool)`

GetProvisioningConnectionOk returns a tuple with the ProvisioningConnection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisioningConnection

`func (o *UpdateSecuritySystemRequest) SetProvisioningConnection(v string)`

SetProvisioningConnection sets ProvisioningConnection field to given value.

### HasProvisioningConnection

`func (o *UpdateSecuritySystemRequest) HasProvisioningConnection() bool`

HasProvisioningConnection returns a boolean if a field has been set.

### GetServiceDeskConnection

`func (o *UpdateSecuritySystemRequest) GetServiceDeskConnection() string`

GetServiceDeskConnection returns the ServiceDeskConnection field if non-nil, zero value otherwise.

### GetServiceDeskConnectionOk

`func (o *UpdateSecuritySystemRequest) GetServiceDeskConnectionOk() (*string, bool)`

GetServiceDeskConnectionOk returns a tuple with the ServiceDeskConnection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceDeskConnection

`func (o *UpdateSecuritySystemRequest) SetServiceDeskConnection(v string)`

SetServiceDeskConnection sets ServiceDeskConnection field to given value.

### HasServiceDeskConnection

`func (o *UpdateSecuritySystemRequest) HasServiceDeskConnection() bool`

HasServiceDeskConnection returns a boolean if a field has been set.

### GetProvisioningcomments

`func (o *UpdateSecuritySystemRequest) GetProvisioningcomments() string`

GetProvisioningcomments returns the Provisioningcomments field if non-nil, zero value otherwise.

### GetProvisioningcommentsOk

`func (o *UpdateSecuritySystemRequest) GetProvisioningcommentsOk() (*string, bool)`

GetProvisioningcommentsOk returns a tuple with the Provisioningcomments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisioningcomments

`func (o *UpdateSecuritySystemRequest) SetProvisioningcomments(v string)`

SetProvisioningcomments sets Provisioningcomments field to given value.

### HasProvisioningcomments

`func (o *UpdateSecuritySystemRequest) HasProvisioningcomments() bool`

HasProvisioningcomments returns a boolean if a field has been set.

### GetInherentSODReportFields

`func (o *UpdateSecuritySystemRequest) GetInherentSODReportFields() []string`

GetInherentSODReportFields returns the InherentSODReportFields field if non-nil, zero value otherwise.

### GetInherentSODReportFieldsOk

`func (o *UpdateSecuritySystemRequest) GetInherentSODReportFieldsOk() (*[]string, bool)`

GetInherentSODReportFieldsOk returns a tuple with the InherentSODReportFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInherentSODReportFields

`func (o *UpdateSecuritySystemRequest) SetInherentSODReportFields(v []string)`

SetInherentSODReportFields sets InherentSODReportFields field to given value.

### HasInherentSODReportFields

`func (o *UpdateSecuritySystemRequest) HasInherentSODReportFields() bool`

HasInherentSODReportFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


