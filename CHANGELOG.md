## 0.2.11 (Released)

FEATURES:

* **New Resource:** `saviynt_entitlement_resource` - Added support for managing entitlements with comprehensive Create, Update and Read operations
* **New Data Source:** `saviynt_entitlement_datasource` - Added support for reading entitlement configurations and metadata
* **New Resource:** `saviynt_privilege_resource` - Added support for managing privileges with full lifecycle management
* **New Data Source:** `saviynt_privilege_datasource` - Added support for reading privilege configurations
* **New Resource:** `saviynt_enterprise_roles_resource` - Added support for managing enterprise roles with user assignments and hierarchical structures
* **New Data Source:** `saviynt_roles_datasource` - Added support for reading role configurations across all role types

ENHANCEMENTS:

* **Write-Only Attributes:** Added write-only variants for all sensitive connection attributes with `_wo` suffix for enhanced security
  - Sensitive data is never stored in Terraform state when using write-only attributes
  - Supported across all connector types (AD, REST, DB, EntraID, SAP, Salesforce, Workday, Unix, GitHub REST, Okta, ADSI)
* **Version Control for Write-Only Updates:** Added `wo_version` attribute to connection resources to enable updates of write-only attributes
* **Enhanced Ephemeral Resources:** Updated ephemeral resources to support all write-only attributes for improved credential management
  - File-based ephemeral resource supports loading write-only credentials from JSON files
  - Environment-based ephemeral resource supports loading write-only credentials from environment variables

SECURITY:

* Improved credential security with write-only attributes that prevent sensitive data persistence in state files
* Enhanced support for external credential management through ephemeral resources

## 0.2.10 (Released)

FEATURES:

* **New Resource:** `saviynt_okta_connection_resource` - Added support for managing Okta connectors
* **New Data Source:** `saviynt_okta_connection_datasource` - Added support for reading Okta connector configurations
* **New Resource:** `saviynt_entitlement_type_resource` - Added support for managing entitlement types
* **New Data Source:** `saviynt_entitlement_type_datasource` - Added support for reading entitlement type configurations

ENHANCEMENTS:

* Updated documentation with Okta connector examples and usage guidelines
* Added comprehensive example configurations for Okta connector resource
* Updated documentation with entitlement type examples and usage guidelines
* Added comprehensive example configurations for entitlement type resource
