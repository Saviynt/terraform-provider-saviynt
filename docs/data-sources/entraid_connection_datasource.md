---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "saviynt_entraid_connection_datasource Data Source - saviynt"
subcategory: ""
description: |-
  Retrieve the details for a given EntraID (AzureAD) connector by its name or key
---

# saviynt_entraid_connection_datasource (Data Source)

Retrieve the details for a given EntraID (AzureAD) connector by its name or key

## Example Usage

```terraform
// Copyright (c) 2025 Saviynt Inc.
// SPDX-License-Identifier: MPL-2.0

data "saviynt_entraid_connection_datasource" "example" {
  connection_key = "123"
  # connection_name="sample"          #Either one can be used
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `authenticate` (Boolean) If false, do not store connection_attributes in state

### Optional

- `connection_key` (Number) The key of the connection.
- `connection_name` (String) The name of the connection.

### Read-Only

- `connection_attributes` (Attributes) (see [below for nested schema](#nestedatt--connection_attributes))
- `connection_type` (String)
- `created_by` (String)
- `created_on` (String)
- `default_sav_roles` (String)
- `description` (String)
- `email_template` (String)
- `error_code` (Number)
- `id` (String) Resource ID.
- `msg` (String)
- `status` (Number)
- `updated_by` (String)

<a id="nestedatt--connection_attributes"></a>
### Nested Schema for `connection_attributes`

Read-Only:

- `aad_tenant_id` (String)
- `account_attributes` (String)
- `account_import_fields` (String)
- `accounts_filter` (String)
- `add_access_json` (String)
- `authentication_endpoint` (String)
- `azure_management_endpoint` (String)
- `config_json` (String)
- `connection_type` (String)
- `create_account_json` (String)
- `create_channel_json` (String)
- `create_group_json` (String)
- `create_new_endpoints` (String)
- `create_team_json` (String)
- `createusers` (String)
- `delete_group_json` (String)
- `deltatokens_json` (String)
- `disable_account_json` (String)
- `enable_account_json` (String)
- `endpoints_filter` (String)
- `entitlement_attribute` (String)
- `entitlement_filter_json` (String)
- `import_user_json` (String)
- `is_timeout_config_validated` (Boolean)
- `is_timeout_supported` (Boolean)
- `microsoft_graph_endpoint` (String)
- `modifyuserdatajson` (String)
- `pam_config` (String)
- `remove_access_json` (String)
- `remove_account_json` (String)
- `status_threshold_config` (String)
- `update_account_json` (String)
- `update_group_json` (String)
- `update_user_json` (String)
