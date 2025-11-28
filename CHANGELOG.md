## 0.3.1 (released)

BUG FIXES:

* **resource/saviynt_entitlement_resource:** Fixed entitlement map update functionality
  - Fixed issue where existing entitlement maps could not be updated
  - Previously only supported add/remove operations, now correctly handles field-level updates

ENHANCEMENTS:

* **resource/saviynt_entitlement_resource:** Added support for additional entitlement map fields
  - Added support for boolean fields: `request_filter`, `exclude_entitlement`, `add_dependent_task`, `remove_dependent_ent_task`

* **resource/saviynt_sftp_connection_resource:** Updated connection type configuration
  - Changed connection type from `SFTP` to `SFTPFileTransfer` to align with official Saviynt EIC documentation
  - **Note:** Ensure your EIC instance has the `SFTPFileTransfer` connection type configured for this resource to work properly
  - For setup instructions, see: [Saviynt Documentation](https://docs.saviyntcloud.com/bundle/SFTP-Certified-25/page/Content/Configuring-the-Integration-for-File-Transfer.htm#creating_a_connection_type)

## 0.3.0 (released)

FEATURES:

* **New Resource:** `saviynt_export_transport_package_resource` - Added support for exporting transport packages from Saviynt
  - Export configurations, roles, and other Saviynt objects to transport packages
  - Version-controlled exports with `export_package_version` trigger attribute
  - Support for online/offline export modes and selective object export
  
* **New Resource:** `saviynt_import_transport_package_resource` - Added support for importing transport packages into Saviynt
  - Import configurations and objects from transport package files
  - Version-controlled imports with `import_package_version` trigger attribute
  - Comprehensive import validation and error handling

* **New Resource:** `saviynt_file_upload_resource` - Added support for uploading files to Saviynt
  - Upload CSV/SAV files to Datafiles/SAV directory
  - Version-controlled uploads with `file_version` trigger attribute for re-upload scenarios
  - Support for various file types and upload locations

* **New Resource:** `saviynt_sftp_connection_resource` - Added support for managing SFTP connectors
  - Full CRUI lifecycle management (Create, Read, Update, Import)
  - Write-only attribute support for sensitive credentials (`auth_credential_value_wo`, `passphrase_wo`)
  - Support for key-based and password-based authentication methods
  - **Note**: Requires manual configuration of "Select Connector version" field in Saviynt UI after creation (see [Troubleshooting Guide](README.md#9-sftp-connection-post-creation-configuration))
  
* **New Data Source:** `saviynt_sftp_connection_datasource` - Added support for reading SFTP connector configurations
  - Authentication control with `authenticate` flag for sensitive data visibility
  - Comprehensive attribute validation and state management

* **New Resource:** `saviynt_file_transfer_job_resource` - Added support for managing file transfer jobs
  - Configure automated file transfer operations between systems
  - Support for upload and download operations with validation
  - Integration with SFTP and other file transfer protocols

* **New Resource:** `saviynt_job_control_resource` - Added support for executing Saviynt jobs
  - Execute jobs on-demand through Terraform operations
  - Version-controlled job execution with `run_job_version` trigger attribute
  - Support for multiple job types and job groups
  - Real-time job execution status and response messages

ENHANCEMENTS:

* **Connection Description Support** - Added support for setting connection descriptions through Terraform
  - All connection resources now support the `description` attribute for documentation and management purposes
  - Previously, connection descriptions could only be set through the Saviynt UI
  - Enables better infrastructure-as-code practices with comprehensive resource documentation

## 0.2.13 (Released)

FEATURES:

* **New Resource:** `saviynt_workday_soap_connection_resource` - Added support for managing Workday SOAP connectors with comprehensive configuration options
  - Full CRUI lifecycle management (Create, Read, Update, Import)
  - Write-only attribute support for sensitive data (`password_wo`, `change_pass_json_wo`, `connection_json_wo`)
* **New Data Source:** `saviynt_workday_soap_connection_datasource` - Added support for reading Workday SOAP connector configurations

ENHANCEMENTS:

* **Enhanced Connection Handling for Compulsory Attributes** - Improved support for connections with mandatory attributes (AD, DB, etc.)
  - Removed required validation for `password` and `password_wo` attributes to enable smoother Terraform operations
  - Enhanced import functionality - existing connections can now be imported without requiring credential updates

* **Version Compatibility with EIC Releases** - Added automatic validation for attribute compatibility with Saviynt EIC versions
  - Validates that resource attributes are supported in your Saviynt EIC version (25.B, 25.A, 24.10)
  - Provides clear error messages when using unsupported attributes for your EIC version
  - Supports REST, DB, Workday, GitHub REST, Security System, Entitlement Type, and Enterprise Role resources
  - Prevents deployment failures by catching version incompatibilities during planning phase

* **New 25.Chicago.EA Attributes Support** - Added support for new attributes available in Saviynt EIC 25.Chicago.EA
  - **REST Connector**: Added `app_type` attribute for CUA configuration
  - **SAP Connector**: Added `role_default_date` attribute for default end date configuration
  - **Unix Connector**: Added `server_type` attribute for server type specification
  - All new attributes include proper version validation and are available in both resource and datasource implementations
  
## 0.2.12 (Released)

FEATURES:

* **New Resource:** Added support for the following jobs:
  - Application Data Import Job: `saviynt_application_data_import_job_resource`
  - WS Retry Job: `saviynt_ws_retry_job_resource`
  - WS Retry Blocking Job: `saviynt_ws_retry_blocking_job_resource`
  - User Import Job: `saviynt_user_import_job_resource`
  - ECM Job: `saviynt_ecm_job_resource`
  - ECM SAP User Job: `saviynt_ecm_sap_user_job_resource`
  - Accounts Import Full Job: `saviynt_accounts_import_full_job_resource`
  - Accounts Import Incremental Job: `saviynt_accounts_import_incremental_job_resource`
  - Schema Role Job: `saviynt_schema_role_job_resource`
  - Schema Account Job: `saviynt_schema_account_job_resource`
  - Schema User Job: `saviynt_schema_user_job_resource`

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
