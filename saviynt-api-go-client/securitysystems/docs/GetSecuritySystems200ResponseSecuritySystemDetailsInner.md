# GetSecuritySystems200ResponseSecuritySystemDetailsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Connectionparameters** | Pointer to **string** | Query to filter the access and display of the endpoint to specific users. If you do not define a query, the endpoint is displayed for all users | [optional] 
**AccessAddWorkflow** | Pointer to **string** | Specify the workflow to be used for approvals for an access request, which can be for an account, entitlements, role, and so on | [optional] 
**DefaultSystem** | Pointer to **string** | Specify true to set the security system as the Default System. Following which accounts search will only be searched and displayed for this security system. | [optional] 
**ProvisioningConnection** | Pointer to **string** | You can use a separate connection to an endpoint where you are performing provisioning or deprovisioning. Based on your requirement, you can specify a separate connection where you want to perform provisioning and de-provisioning. | [optional] 
**UpdateDate** | Pointer to **string** | Last updated date for the security system. | [optional] 
**DisplayName** | Pointer to **string** | Specify a user-friendly display name that is shown on the the user interface. | [optional] 
**ConnectionType** | Pointer to **string** | Specify a connection type to view all the connection in EIC for the connection type. | [optional] 
**Instantprovision** | Pointer to **string** | If &#x60;true&#x60;, prevents users from submitting duplicate provisioning requests. | [optional] 
**Connectionname** | Pointer to **string** | Select the connection name for performing reconciliation of identity objects from third-party application. | [optional] 
**Hostname** | Pointer to **string** | Security system for which you want to create an endpoint. | [optional] 
**FirefighteridWorkflow** | Pointer to **string** | Specifies the workflow for handling firefighter ID requests. | [optional] 
**FirefighteridRequestAccessWorkflow** | Pointer to **string** | Defines the workflow for requesting access to firefighter IDs. | [optional] 
**Provisioningcomments** | Pointer to **string** | Specify relevant comments for performing provisioning. | [optional] 
**ProposedAccountOwnersworkflow** | Pointer to **string** | Defines the workflow for assigning proposed account owners. | [optional] 
**ServiceDeskConnection** | Pointer to **string** | Specify the Service Desk Connection used for integration with a ticketing system, which can be a disconnected system too. | [optional] 
**Connection** | Pointer to **string** |  | [optional] 
**Useopenconnector** | Pointer to **string** | Specify true to enable the connectivity with any system over the open-source connectors such as REST. | [optional] 
**ReconApplication** | Pointer to **string** | Specify true to importing data from respective endpoint associated to security system. | [optional] 
**CreateDate** | Pointer to **string** |  | [optional] 
**Endpoints** | Pointer to **string** |  | [optional] 
**UpdatedBy** | Pointer to **string** |  | [optional] 
**AutomatedProvisioning** | Pointer to **string** | Specify true to enable automated provisioning. | [optional] 
**Systemname** | Pointer to **string** | Specify the security system name. | [optional] 
**AccessRemoveWorkflow** | Pointer to **string** | Specify the workflow to be used when access has to be revoked, which can be for an account, entitlement, or any other de-provisioning task. | [optional] 
**AddServiceAccountWorkflow** | Pointer to **string** | Workflow for adding a service account. | [optional] 
**RemoveServiceAccountWorkflow** | Pointer to **string** | Workflow for removing a service account. | [optional] 
**PolicyRule** | Pointer to **string** | Use this setting to assign the password policy for the security system. | [optional] 
**PolicyRuleServiceAccount** | Pointer to **string** | Defines the password policy for service account passwords. | [optional] 
**CreatedFrom** | Pointer to **string** |  | [optional] 
**CreatedBy** | Pointer to **string** |  | [optional] 
**ExternalRiskConnectionJson** | Pointer to **string** | Contains JSON configuration for external risk connections and is applicable only for few connections like SAP | [optional] 
**ProvisioningTries** | Pointer to **string** | Number of provisioning tries. | [optional] 
**Port** | Pointer to **string** | Description for the endpoint | [optional] 
**InherentSODReportFields** | Pointer to **[]string** |  You can use this option used to filter out columns in SOD. | [optional] 
**Status** | Pointer to **string** |  | [optional] 

## Methods

### NewGetSecuritySystems200ResponseSecuritySystemDetailsInner

`func NewGetSecuritySystems200ResponseSecuritySystemDetailsInner() *GetSecuritySystems200ResponseSecuritySystemDetailsInner`

NewGetSecuritySystems200ResponseSecuritySystemDetailsInner instantiates a new GetSecuritySystems200ResponseSecuritySystemDetailsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetSecuritySystems200ResponseSecuritySystemDetailsInnerWithDefaults

`func NewGetSecuritySystems200ResponseSecuritySystemDetailsInnerWithDefaults() *GetSecuritySystems200ResponseSecuritySystemDetailsInner`

NewGetSecuritySystems200ResponseSecuritySystemDetailsInnerWithDefaults instantiates a new GetSecuritySystems200ResponseSecuritySystemDetailsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConnectionparameters

`func (o *GetSecuritySystems200ResponseSecuritySystemDetailsInner) GetConnectionparameters() string`

GetConnectionparameters returns the Connectionparameters field if non-nil, zero value otherwise.

### GetConnectionparametersOk

`func (o *GetSecuritySystems200ResponseSecuritySystemDetailsInner) GetConnectionparametersOk() (*string, bool)`

GetConnectionparametersOk returns a tuple with the Connectionparameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionparameters

`func (o *GetSecuritySystems200ResponseSecuritySystemDetailsInner) SetConnectionparameters(v string)`

SetConnectionparameters sets Connectionparameters field to given value.

### HasConnectionparameters

`func (o *GetSecuritySystems200ResponseSecuritySystemDetailsInner) HasConnectionparameters() bool`

HasConnectionparameters returns a boolean if a field has been set.

### GetAccessAddWorkflow

`func (o *GetSecuritySystems200ResponseSecuritySystemDetailsInner) GetAccessAddWorkflow() string`

GetAccessAddWorkflow returns the AccessAddWorkflow field if non-nil, zero value otherwise.

### GetAccessAddWorkflowOk

`func (o *GetSecuritySystems200ResponseSecuritySystemDetailsInner) GetAccessAddWorkflowOk() (*string, bool)`

GetAccessAddWorkflowOk returns a tuple with the AccessAddWorkflow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessAddWorkflow

`func (o *GetSecuritySystems200ResponseSecuritySystemDetailsInner) SetAccessAddWorkflow(v string)`

SetAccessAddWorkflow sets AccessAddWorkflow field to given value.

### HasAccessAddWorkflow

`func (o *GetSecuritySystems200ResponseSecuritySystemDetailsInner) HasAccessAddWorkflow() bool`

HasAccessAddWorkflow returns a boolean if a field has been set.

### GetDefaultSystem

`func (o *GetSecuritySystems200ResponseSecuritySystemDetailsInner) GetDefaultSystem() string`

GetDefaultSystem returns the DefaultSystem field if non-nil, zero value otherwise.

### GetDefaultSystemOk

`func (o *GetSecuritySystems200ResponseSecuritySystemDetailsInner) GetDefaultSystemOk() (*string, bool)`

GetDefaultSystemOk returns a tuple with the DefaultSystem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultSystem

`func (o *GetSecuritySystems200ResponseSecuritySystemDetailsInner) SetDefaultSystem(v string)`

SetDefaultSystem sets DefaultSystem field to given value.

### HasDefaultSystem

`func (o *GetSecuritySystems200ResponseSecuritySystemDetailsInner) HasDefaultSystem() bool`

HasDefaultSystem returns a boolean if a field has been set.

### GetProvisioningConnection

`func (o *GetSecuritySystems200ResponseSecuritySystemDetailsInner) GetProvisioningConnection() string`

GetProvisioningConnection returns the ProvisioningConnection field if non-nil, zero value otherwise.

### GetProvisioningConnectionOk

`func (o *GetSecuritySystems200ResponseSecuritySystemDetailsInner) GetProvisioningConnectionOk() (*string, bool)`

GetProvisioningConnectionOk returns a tuple with the ProvisioningConnection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisioningConnection

`func (o *GetSecuritySystems200ResponseSecuritySystemDetailsInner) SetProvisioningConnection(v string)`

SetProvisioningConnection sets ProvisioningConnection field to given value.

### HasProvisioningConnection

`func (o *GetSecuritySystems200ResponseSecuritySystemDetailsInner) HasProvisioningConnection() bool`

HasProvisioningConnection returns a boolean if a field has been set.

### GetUpdateDate

`func (o *GetSecuritySystems200ResponseSecuritySystemDetailsInner) GetUpdateDate() string`

GetUpdateDate returns the UpdateDate field if non-nil, zero value otherwise.

### GetUpdateDateOk

`func (o *GetSecuritySystems200ResponseSecuritySystemDetailsInner) GetUpdateDateOk() (*string, bool)`

GetUpdateDateOk returns a tuple with the UpdateDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateDate

`func (o *GetSecuritySystems200ResponseSecuritySystemDetailsInner) SetUpdateDate(v string)`

SetUpdateDate sets UpdateDate field to given value.

### HasUpdateDate

`func (o *GetSecuritySystems200ResponseSecuritySystemDetailsInner) HasUpdateDate() bool`

HasUpdateDate returns a boolean if a field has been set.

### GetDisplayName

`func (o *GetSecuritySystems200ResponseSecuritySystemDetailsInner) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *GetSecuritySystems200ResponseSecuritySystemDetailsInner) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *GetSecuritySystems200ResponseSecuritySystemDetailsInner) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *GetSecuritySystems200ResponseSecuritySystemDetailsInner) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetConnectionType

`func (o *GetSecuritySystems200ResponseSecuritySystemDetailsInner) GetConnectionType() string`

GetConnectionType returns the ConnectionType field if non-nil, zero value otherwise.

### GetConnectionTypeOk

`func (o *GetSecuritySystems200ResponseSecuritySystemDetailsInner) GetConnectionTypeOk() (*string, bool)`

GetConnectionTypeOk returns a tuple with the ConnectionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionType

`func (o *GetSecuritySystems200ResponseSecuritySystemDetailsInner) SetConnectionType(v string)`

SetConnectionType sets ConnectionType field to given value.

### HasConnectionType

`func (o *GetSecuritySystems200ResponseSecuritySystemDetailsInner) HasConnectionType() bool`

HasConnectionType returns a boolean if a field has been set.

### GetInstantprovision

`func (o *GetSecuritySystems200ResponseSecuritySystemDetailsInner) GetInstantprovision() string`

GetInstantprovision returns the Instantprovision field if non-nil, zero value otherwise.

### GetInstantprovisionOk

`func (o *GetSecuritySystems200ResponseSecuritySystemDetailsInner) GetInstantprovisionOk() (*string, bool)`

GetInstantprovisionOk returns a tuple with the Instantprovision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstantprovision

`func (o *GetSecuritySystems200ResponseSecuritySystemDetailsInner) SetInstantprovision(v string)`

SetInstantprovision sets Instantprovision field to given value.

### HasInstantprovision

`func (o *GetSecuritySystems200ResponseSecuritySystemDetailsInner) HasInstantprovision() bool`

HasInstantprovision returns a boolean if a field has been set.

### GetConnectionname

`func (o *GetSecuritySystems200ResponseSecuritySystemDetailsInner) GetConnectionname() string`

GetConnectionname returns the Connectionname field if non-nil, zero value otherwise.

### GetConnectionnameOk

`func (o *GetSecuritySystems200ResponseSecuritySystemDetailsInner) GetConnectionnameOk() (*string, bool)`

GetConnectionnameOk returns a tuple with the Connectionname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionname

`func (o *GetSecuritySystems200ResponseSecuritySystemDetailsInner) SetConnectionname(v string)`

SetConnectionname sets Connectionname field to given value.

### HasConnectionname

`func (o *GetSecuritySystems200ResponseSecuritySystemDetailsInner) HasConnectionname() bool`

HasConnectionname returns a boolean if a field has been set.

### GetHostname

`func (o *GetSecuritySystems200ResponseSecuritySystemDetailsInner) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *GetSecuritySystems200ResponseSecuritySystemDetailsInner) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *GetSecuritySystems200ResponseSecuritySystemDetailsInner) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *GetSecuritySystems200ResponseSecuritySystemDetailsInner) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### GetFirefighteridWorkflow

`func (o *GetSecuritySystems200ResponseSecuritySystemDetailsInner) GetFirefighteridWorkflow() string`

GetFirefighteridWorkflow returns the FirefighteridWorkflow field if non-nil, zero value otherwise.

### GetFirefighteridWorkflowOk

`func (o *GetSecuritySystems200ResponseSecuritySystemDetailsInner) GetFirefighteridWorkflowOk() (*string, bool)`

GetFirefighteridWorkflowOk returns a tuple with the FirefighteridWorkflow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirefighteridWorkflow

`func (o *GetSecuritySystems200ResponseSecuritySystemDetailsInner) SetFirefighteridWorkflow(v string)`

SetFirefighteridWorkflow sets FirefighteridWorkflow field to given value.

### HasFirefighteridWorkflow

`func (o *GetSecuritySystems200ResponseSecuritySystemDetailsInner) HasFirefighteridWorkflow() bool`

HasFirefighteridWorkflow returns a boolean if a field has been set.

### GetFirefighteridRequestAccessWorkflow

`func (o *GetSecuritySystems200ResponseSecuritySystemDetailsInner) GetFirefighteridRequestAccessWorkflow() string`

GetFirefighteridRequestAccessWorkflow returns the FirefighteridRequestAccessWorkflow field if non-nil, zero value otherwise.

### GetFirefighteridRequestAccessWorkflowOk

`func (o *GetSecuritySystems200ResponseSecuritySystemDetailsInner) GetFirefighteridRequestAccessWorkflowOk() (*string, bool)`

GetFirefighteridRequestAccessWorkflowOk returns a tuple with the FirefighteridRequestAccessWorkflow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirefighteridRequestAccessWorkflow

`func (o *GetSecuritySystems200ResponseSecuritySystemDetailsInner) SetFirefighteridRequestAccessWorkflow(v string)`

SetFirefighteridRequestAccessWorkflow sets FirefighteridRequestAccessWorkflow field to given value.

### HasFirefighteridRequestAccessWorkflow

`func (o *GetSecuritySystems200ResponseSecuritySystemDetailsInner) HasFirefighteridRequestAccessWorkflow() bool`

HasFirefighteridRequestAccessWorkflow returns a boolean if a field has been set.

### GetProvisioningcomments

`func (o *GetSecuritySystems200ResponseSecuritySystemDetailsInner) GetProvisioningcomments() string`

GetProvisioningcomments returns the Provisioningcomments field if non-nil, zero value otherwise.

### GetProvisioningcommentsOk

`func (o *GetSecuritySystems200ResponseSecuritySystemDetailsInner) GetProvisioningcommentsOk() (*string, bool)`

GetProvisioningcommentsOk returns a tuple with the Provisioningcomments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisioningcomments

`func (o *GetSecuritySystems200ResponseSecuritySystemDetailsInner) SetProvisioningcomments(v string)`

SetProvisioningcomments sets Provisioningcomments field to given value.

### HasProvisioningcomments

`func (o *GetSecuritySystems200ResponseSecuritySystemDetailsInner) HasProvisioningcomments() bool`

HasProvisioningcomments returns a boolean if a field has been set.

### GetProposedAccountOwnersworkflow

`func (o *GetSecuritySystems200ResponseSecuritySystemDetailsInner) GetProposedAccountOwnersworkflow() string`

GetProposedAccountOwnersworkflow returns the ProposedAccountOwnersworkflow field if non-nil, zero value otherwise.

### GetProposedAccountOwnersworkflowOk

`func (o *GetSecuritySystems200ResponseSecuritySystemDetailsInner) GetProposedAccountOwnersworkflowOk() (*string, bool)`

GetProposedAccountOwnersworkflowOk returns a tuple with the ProposedAccountOwnersworkflow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProposedAccountOwnersworkflow

`func (o *GetSecuritySystems200ResponseSecuritySystemDetailsInner) SetProposedAccountOwnersworkflow(v string)`

SetProposedAccountOwnersworkflow sets ProposedAccountOwnersworkflow field to given value.

### HasProposedAccountOwnersworkflow

`func (o *GetSecuritySystems200ResponseSecuritySystemDetailsInner) HasProposedAccountOwnersworkflow() bool`

HasProposedAccountOwnersworkflow returns a boolean if a field has been set.

### GetServiceDeskConnection

`func (o *GetSecuritySystems200ResponseSecuritySystemDetailsInner) GetServiceDeskConnection() string`

GetServiceDeskConnection returns the ServiceDeskConnection field if non-nil, zero value otherwise.

### GetServiceDeskConnectionOk

`func (o *GetSecuritySystems200ResponseSecuritySystemDetailsInner) GetServiceDeskConnectionOk() (*string, bool)`

GetServiceDeskConnectionOk returns a tuple with the ServiceDeskConnection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceDeskConnection

`func (o *GetSecuritySystems200ResponseSecuritySystemDetailsInner) SetServiceDeskConnection(v string)`

SetServiceDeskConnection sets ServiceDeskConnection field to given value.

### HasServiceDeskConnection

`func (o *GetSecuritySystems200ResponseSecuritySystemDetailsInner) HasServiceDeskConnection() bool`

HasServiceDeskConnection returns a boolean if a field has been set.

### GetConnection

`func (o *GetSecuritySystems200ResponseSecuritySystemDetailsInner) GetConnection() string`

GetConnection returns the Connection field if non-nil, zero value otherwise.

### GetConnectionOk

`func (o *GetSecuritySystems200ResponseSecuritySystemDetailsInner) GetConnectionOk() (*string, bool)`

GetConnectionOk returns a tuple with the Connection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnection

`func (o *GetSecuritySystems200ResponseSecuritySystemDetailsInner) SetConnection(v string)`

SetConnection sets Connection field to given value.

### HasConnection

`func (o *GetSecuritySystems200ResponseSecuritySystemDetailsInner) HasConnection() bool`

HasConnection returns a boolean if a field has been set.

### GetUseopenconnector

`func (o *GetSecuritySystems200ResponseSecuritySystemDetailsInner) GetUseopenconnector() string`

GetUseopenconnector returns the Useopenconnector field if non-nil, zero value otherwise.

### GetUseopenconnectorOk

`func (o *GetSecuritySystems200ResponseSecuritySystemDetailsInner) GetUseopenconnectorOk() (*string, bool)`

GetUseopenconnectorOk returns a tuple with the Useopenconnector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseopenconnector

`func (o *GetSecuritySystems200ResponseSecuritySystemDetailsInner) SetUseopenconnector(v string)`

SetUseopenconnector sets Useopenconnector field to given value.

### HasUseopenconnector

`func (o *GetSecuritySystems200ResponseSecuritySystemDetailsInner) HasUseopenconnector() bool`

HasUseopenconnector returns a boolean if a field has been set.

### GetReconApplication

`func (o *GetSecuritySystems200ResponseSecuritySystemDetailsInner) GetReconApplication() string`

GetReconApplication returns the ReconApplication field if non-nil, zero value otherwise.

### GetReconApplicationOk

`func (o *GetSecuritySystems200ResponseSecuritySystemDetailsInner) GetReconApplicationOk() (*string, bool)`

GetReconApplicationOk returns a tuple with the ReconApplication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReconApplication

`func (o *GetSecuritySystems200ResponseSecuritySystemDetailsInner) SetReconApplication(v string)`

SetReconApplication sets ReconApplication field to given value.

### HasReconApplication

`func (o *GetSecuritySystems200ResponseSecuritySystemDetailsInner) HasReconApplication() bool`

HasReconApplication returns a boolean if a field has been set.

### GetCreateDate

`func (o *GetSecuritySystems200ResponseSecuritySystemDetailsInner) GetCreateDate() string`

GetCreateDate returns the CreateDate field if non-nil, zero value otherwise.

### GetCreateDateOk

`func (o *GetSecuritySystems200ResponseSecuritySystemDetailsInner) GetCreateDateOk() (*string, bool)`

GetCreateDateOk returns a tuple with the CreateDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateDate

`func (o *GetSecuritySystems200ResponseSecuritySystemDetailsInner) SetCreateDate(v string)`

SetCreateDate sets CreateDate field to given value.

### HasCreateDate

`func (o *GetSecuritySystems200ResponseSecuritySystemDetailsInner) HasCreateDate() bool`

HasCreateDate returns a boolean if a field has been set.

### GetEndpoints

`func (o *GetSecuritySystems200ResponseSecuritySystemDetailsInner) GetEndpoints() string`

GetEndpoints returns the Endpoints field if non-nil, zero value otherwise.

### GetEndpointsOk

`func (o *GetSecuritySystems200ResponseSecuritySystemDetailsInner) GetEndpointsOk() (*string, bool)`

GetEndpointsOk returns a tuple with the Endpoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoints

`func (o *GetSecuritySystems200ResponseSecuritySystemDetailsInner) SetEndpoints(v string)`

SetEndpoints sets Endpoints field to given value.

### HasEndpoints

`func (o *GetSecuritySystems200ResponseSecuritySystemDetailsInner) HasEndpoints() bool`

HasEndpoints returns a boolean if a field has been set.

### GetUpdatedBy

`func (o *GetSecuritySystems200ResponseSecuritySystemDetailsInner) GetUpdatedBy() string`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *GetSecuritySystems200ResponseSecuritySystemDetailsInner) GetUpdatedByOk() (*string, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *GetSecuritySystems200ResponseSecuritySystemDetailsInner) SetUpdatedBy(v string)`

SetUpdatedBy sets UpdatedBy field to given value.

### HasUpdatedBy

`func (o *GetSecuritySystems200ResponseSecuritySystemDetailsInner) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.

### GetAutomatedProvisioning

`func (o *GetSecuritySystems200ResponseSecuritySystemDetailsInner) GetAutomatedProvisioning() string`

GetAutomatedProvisioning returns the AutomatedProvisioning field if non-nil, zero value otherwise.

### GetAutomatedProvisioningOk

`func (o *GetSecuritySystems200ResponseSecuritySystemDetailsInner) GetAutomatedProvisioningOk() (*string, bool)`

GetAutomatedProvisioningOk returns a tuple with the AutomatedProvisioning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutomatedProvisioning

`func (o *GetSecuritySystems200ResponseSecuritySystemDetailsInner) SetAutomatedProvisioning(v string)`

SetAutomatedProvisioning sets AutomatedProvisioning field to given value.

### HasAutomatedProvisioning

`func (o *GetSecuritySystems200ResponseSecuritySystemDetailsInner) HasAutomatedProvisioning() bool`

HasAutomatedProvisioning returns a boolean if a field has been set.

### GetSystemname

`func (o *GetSecuritySystems200ResponseSecuritySystemDetailsInner) GetSystemname() string`

GetSystemname returns the Systemname field if non-nil, zero value otherwise.

### GetSystemnameOk

`func (o *GetSecuritySystems200ResponseSecuritySystemDetailsInner) GetSystemnameOk() (*string, bool)`

GetSystemnameOk returns a tuple with the Systemname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemname

`func (o *GetSecuritySystems200ResponseSecuritySystemDetailsInner) SetSystemname(v string)`

SetSystemname sets Systemname field to given value.

### HasSystemname

`func (o *GetSecuritySystems200ResponseSecuritySystemDetailsInner) HasSystemname() bool`

HasSystemname returns a boolean if a field has been set.

### GetAccessRemoveWorkflow

`func (o *GetSecuritySystems200ResponseSecuritySystemDetailsInner) GetAccessRemoveWorkflow() string`

GetAccessRemoveWorkflow returns the AccessRemoveWorkflow field if non-nil, zero value otherwise.

### GetAccessRemoveWorkflowOk

`func (o *GetSecuritySystems200ResponseSecuritySystemDetailsInner) GetAccessRemoveWorkflowOk() (*string, bool)`

GetAccessRemoveWorkflowOk returns a tuple with the AccessRemoveWorkflow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessRemoveWorkflow

`func (o *GetSecuritySystems200ResponseSecuritySystemDetailsInner) SetAccessRemoveWorkflow(v string)`

SetAccessRemoveWorkflow sets AccessRemoveWorkflow field to given value.

### HasAccessRemoveWorkflow

`func (o *GetSecuritySystems200ResponseSecuritySystemDetailsInner) HasAccessRemoveWorkflow() bool`

HasAccessRemoveWorkflow returns a boolean if a field has been set.

### GetAddServiceAccountWorkflow

`func (o *GetSecuritySystems200ResponseSecuritySystemDetailsInner) GetAddServiceAccountWorkflow() string`

GetAddServiceAccountWorkflow returns the AddServiceAccountWorkflow field if non-nil, zero value otherwise.

### GetAddServiceAccountWorkflowOk

`func (o *GetSecuritySystems200ResponseSecuritySystemDetailsInner) GetAddServiceAccountWorkflowOk() (*string, bool)`

GetAddServiceAccountWorkflowOk returns a tuple with the AddServiceAccountWorkflow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddServiceAccountWorkflow

`func (o *GetSecuritySystems200ResponseSecuritySystemDetailsInner) SetAddServiceAccountWorkflow(v string)`

SetAddServiceAccountWorkflow sets AddServiceAccountWorkflow field to given value.

### HasAddServiceAccountWorkflow

`func (o *GetSecuritySystems200ResponseSecuritySystemDetailsInner) HasAddServiceAccountWorkflow() bool`

HasAddServiceAccountWorkflow returns a boolean if a field has been set.

### GetRemoveServiceAccountWorkflow

`func (o *GetSecuritySystems200ResponseSecuritySystemDetailsInner) GetRemoveServiceAccountWorkflow() string`

GetRemoveServiceAccountWorkflow returns the RemoveServiceAccountWorkflow field if non-nil, zero value otherwise.

### GetRemoveServiceAccountWorkflowOk

`func (o *GetSecuritySystems200ResponseSecuritySystemDetailsInner) GetRemoveServiceAccountWorkflowOk() (*string, bool)`

GetRemoveServiceAccountWorkflowOk returns a tuple with the RemoveServiceAccountWorkflow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoveServiceAccountWorkflow

`func (o *GetSecuritySystems200ResponseSecuritySystemDetailsInner) SetRemoveServiceAccountWorkflow(v string)`

SetRemoveServiceAccountWorkflow sets RemoveServiceAccountWorkflow field to given value.

### HasRemoveServiceAccountWorkflow

`func (o *GetSecuritySystems200ResponseSecuritySystemDetailsInner) HasRemoveServiceAccountWorkflow() bool`

HasRemoveServiceAccountWorkflow returns a boolean if a field has been set.

### GetPolicyRule

`func (o *GetSecuritySystems200ResponseSecuritySystemDetailsInner) GetPolicyRule() string`

GetPolicyRule returns the PolicyRule field if non-nil, zero value otherwise.

### GetPolicyRuleOk

`func (o *GetSecuritySystems200ResponseSecuritySystemDetailsInner) GetPolicyRuleOk() (*string, bool)`

GetPolicyRuleOk returns a tuple with the PolicyRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyRule

`func (o *GetSecuritySystems200ResponseSecuritySystemDetailsInner) SetPolicyRule(v string)`

SetPolicyRule sets PolicyRule field to given value.

### HasPolicyRule

`func (o *GetSecuritySystems200ResponseSecuritySystemDetailsInner) HasPolicyRule() bool`

HasPolicyRule returns a boolean if a field has been set.

### GetPolicyRuleServiceAccount

`func (o *GetSecuritySystems200ResponseSecuritySystemDetailsInner) GetPolicyRuleServiceAccount() string`

GetPolicyRuleServiceAccount returns the PolicyRuleServiceAccount field if non-nil, zero value otherwise.

### GetPolicyRuleServiceAccountOk

`func (o *GetSecuritySystems200ResponseSecuritySystemDetailsInner) GetPolicyRuleServiceAccountOk() (*string, bool)`

GetPolicyRuleServiceAccountOk returns a tuple with the PolicyRuleServiceAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyRuleServiceAccount

`func (o *GetSecuritySystems200ResponseSecuritySystemDetailsInner) SetPolicyRuleServiceAccount(v string)`

SetPolicyRuleServiceAccount sets PolicyRuleServiceAccount field to given value.

### HasPolicyRuleServiceAccount

`func (o *GetSecuritySystems200ResponseSecuritySystemDetailsInner) HasPolicyRuleServiceAccount() bool`

HasPolicyRuleServiceAccount returns a boolean if a field has been set.

### GetCreatedFrom

`func (o *GetSecuritySystems200ResponseSecuritySystemDetailsInner) GetCreatedFrom() string`

GetCreatedFrom returns the CreatedFrom field if non-nil, zero value otherwise.

### GetCreatedFromOk

`func (o *GetSecuritySystems200ResponseSecuritySystemDetailsInner) GetCreatedFromOk() (*string, bool)`

GetCreatedFromOk returns a tuple with the CreatedFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedFrom

`func (o *GetSecuritySystems200ResponseSecuritySystemDetailsInner) SetCreatedFrom(v string)`

SetCreatedFrom sets CreatedFrom field to given value.

### HasCreatedFrom

`func (o *GetSecuritySystems200ResponseSecuritySystemDetailsInner) HasCreatedFrom() bool`

HasCreatedFrom returns a boolean if a field has been set.

### GetCreatedBy

`func (o *GetSecuritySystems200ResponseSecuritySystemDetailsInner) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *GetSecuritySystems200ResponseSecuritySystemDetailsInner) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *GetSecuritySystems200ResponseSecuritySystemDetailsInner) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *GetSecuritySystems200ResponseSecuritySystemDetailsInner) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetExternalRiskConnectionJson

`func (o *GetSecuritySystems200ResponseSecuritySystemDetailsInner) GetExternalRiskConnectionJson() string`

GetExternalRiskConnectionJson returns the ExternalRiskConnectionJson field if non-nil, zero value otherwise.

### GetExternalRiskConnectionJsonOk

`func (o *GetSecuritySystems200ResponseSecuritySystemDetailsInner) GetExternalRiskConnectionJsonOk() (*string, bool)`

GetExternalRiskConnectionJsonOk returns a tuple with the ExternalRiskConnectionJson field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalRiskConnectionJson

`func (o *GetSecuritySystems200ResponseSecuritySystemDetailsInner) SetExternalRiskConnectionJson(v string)`

SetExternalRiskConnectionJson sets ExternalRiskConnectionJson field to given value.

### HasExternalRiskConnectionJson

`func (o *GetSecuritySystems200ResponseSecuritySystemDetailsInner) HasExternalRiskConnectionJson() bool`

HasExternalRiskConnectionJson returns a boolean if a field has been set.

### GetProvisioningTries

`func (o *GetSecuritySystems200ResponseSecuritySystemDetailsInner) GetProvisioningTries() string`

GetProvisioningTries returns the ProvisioningTries field if non-nil, zero value otherwise.

### GetProvisioningTriesOk

`func (o *GetSecuritySystems200ResponseSecuritySystemDetailsInner) GetProvisioningTriesOk() (*string, bool)`

GetProvisioningTriesOk returns a tuple with the ProvisioningTries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisioningTries

`func (o *GetSecuritySystems200ResponseSecuritySystemDetailsInner) SetProvisioningTries(v string)`

SetProvisioningTries sets ProvisioningTries field to given value.

### HasProvisioningTries

`func (o *GetSecuritySystems200ResponseSecuritySystemDetailsInner) HasProvisioningTries() bool`

HasProvisioningTries returns a boolean if a field has been set.

### GetPort

`func (o *GetSecuritySystems200ResponseSecuritySystemDetailsInner) GetPort() string`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *GetSecuritySystems200ResponseSecuritySystemDetailsInner) GetPortOk() (*string, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *GetSecuritySystems200ResponseSecuritySystemDetailsInner) SetPort(v string)`

SetPort sets Port field to given value.

### HasPort

`func (o *GetSecuritySystems200ResponseSecuritySystemDetailsInner) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetInherentSODReportFields

`func (o *GetSecuritySystems200ResponseSecuritySystemDetailsInner) GetInherentSODReportFields() []string`

GetInherentSODReportFields returns the InherentSODReportFields field if non-nil, zero value otherwise.

### GetInherentSODReportFieldsOk

`func (o *GetSecuritySystems200ResponseSecuritySystemDetailsInner) GetInherentSODReportFieldsOk() (*[]string, bool)`

GetInherentSODReportFieldsOk returns a tuple with the InherentSODReportFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInherentSODReportFields

`func (o *GetSecuritySystems200ResponseSecuritySystemDetailsInner) SetInherentSODReportFields(v []string)`

SetInherentSODReportFields sets InherentSODReportFields field to given value.

### HasInherentSODReportFields

`func (o *GetSecuritySystems200ResponseSecuritySystemDetailsInner) HasInherentSODReportFields() bool`

HasInherentSODReportFields returns a boolean if a field has been set.

### GetStatus

`func (o *GetSecuritySystems200ResponseSecuritySystemDetailsInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetSecuritySystems200ResponseSecuritySystemDetailsInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetSecuritySystems200ResponseSecuritySystemDetailsInner) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetSecuritySystems200ResponseSecuritySystemDetailsInner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


