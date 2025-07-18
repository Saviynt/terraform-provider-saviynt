---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "saviynt_security_system_datasource Data Source - saviynt"
subcategory: ""
description: |-
  Datasource to retrieve all security systems
---

# saviynt_security_system_datasource (Data Source)

Datasource to retrieve all security systems

## Example Usage

```terraform
// Copyright (c) 2025 Saviynt Inc.
// SPDX-License-Identifier: MPL-2.0

# List all the security systems
data "saviynt_security_system_datasource" "all" {}


# Get a security system by its name
data "saviynt_security_system_datasource" "by_name" {
  systemname = "sample"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `connection_type` (String) Owner of the endpoint. If ownerType is User, specify the username of the owner. If ownerType is Usergroup, sepecify the name of the User group.
- `connectionname` (String) Owner type of the endpoint. It could be User or Usergroup.
- `max` (Number) Name for the security system that will be displayed in the user interface.
- `offset` (Number) Security system for which you want to create an endpoint.
- `systemname` (String) Name of the security systeme.

### Read-Only

- `display_count` (Number) The number of items currently displayed (e.g., on the current page or view).
- `error_code` (String) An error code where '0' signifies success and '1' signifies an unsuccessful operation.
- `msg` (String) A message indicating the outcome of the operation.
- `results` (Attributes List) List of security systems retrieved (see [below for nested schema](#nestedatt--results))
- `total_count` (Number) The total number of items available in the dataset, irrespective of the current display settings.

<a id="nestedatt--results"></a>
### Nested Schema for `results`

Read-Only:

- `access_add_workflow` (String) Specify the workflow used for approvals for an access request (account, entitlements, role, etc.).
- `access_remove_workflow` (String) Workflow used when revoking access from accounts, entitlements, or performing other de-provisioning tasks.
- `add_service_account_workflow` (String) Workflow for adding a service account.
- `automated_provisioning` (String) Enables automated provisioning if set to true.
- `connection` (String) Primary connection used by the security system.
- `connection_parameters` (String) Query or parameters to restrict endpoint access to specific users.
- `connection_type_1` (String) Specify a connection type to view all connections in EIC for the connection type.
- `connectionname1` (String) Name of connection used for reconciling identity objects from third-party applications.
- `create_date` (String) Timestamp indicating when the security system was created.
- `created_by` (String) Identifier of the user who created the security system.
- `created_from` (String) Origin or method through which the security system was created.
- `default_system` (String) Sets this security system as the default system for account searches when set to true.
- `display_name` (String) Specify a user-friendly display name that is shown on the the user interface.
- `endpoints` (String) Endpoints associated with the security system.
- `external_risk_connection_json` (String) JSON configuration for external risk connections (e.g., SAP).
- `firefighterid_request_access_workflow` (String) Workflow for requesting access to firefighter IDs.
- `firefighterid_workflow` (String) Workflow for handling firefighter ID requests.
- `hostname` (String) Security system for which you want to create an endpoint.
- `manage_entity` (String) Indicates if entity management is enabled for the security system.
- `persistent_data` (String) Indicates whether persistent data storage is enabled for the security system.
- `port` (String) Port information or description for the endpoint.
- `proposed_account_owners_workflow` (String) Workflow for assigning proposed account owners.
- `provisioning_comments` (String) Comments relevant to provisioning actions.
- `provisioning_connection` (String) Dedicated connection for provisioning and de-provisioning tasks.
- `provisioning_tries` (String) Number of attempts allowed for provisioning actions.
- `recon_application` (String) Enables importing data from endpoints associated with the security system.
- `remove_service_account_workflow` (String) Workflow for removing a service account.
- `service_desk_connection` (String) Connection to service desk or ticketing system integration.
- `status` (String) Current status of the security system (e.g., enabled, disabled).
- `systemname1` (String) Specify the security system name.
- `update_date` (String) Timestamp indicating the last update to the security system.
- `updated_by` (String) Identifier of the user who last updated the security system.
- `use_open_connector` (String) Enables connectivity using open-source connectors such as REST if set to true.
