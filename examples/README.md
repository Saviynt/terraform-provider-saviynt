# Configuration Examples

Below are example configurations to help guide you through your Terraform journey with **Saviynt**.

> **Note:** This is not an exhaustive list — only common resources and use cases are currently included. Additional examples and documentation will be added as the project evolves.

---

## Supported Resources
- [saviynt_security_system_resource](./resources/saviynt_security_system_resource) Manages lifecycle (create, update, read) of security systems. Supports workflows, connectors, password policies and more.
- [saviynt_endpoint_resource](./resources/saviynt_endpoint_resource) For managing endpoints definitions used by security systems.
- [saviynt_ad_connection_resource](./resources/saviynt_ad_connection_resource) For managing Active Directory (AD) connections.
- [saviynt_adsi_connection_resource](./resources/saviynt_adsi_connection_resource) For managing ADSI connections.
- [saviynt_db_connection_resource](./resources/saviynt_db_connection_resource) For managing DataBase connections.
- [saviynt_entraid_connection_resource](./resources/saviynt_entraid_connection_resource) For managing EntraID(AzureAD) connections.
- [saviynt_github_rest_connection_resource](./resources/saviynt_github_rest_connection_resource) For managing Github REST connections.
- [saviynt_rest_connection_resource](./resources/saviynt_rest_connection_resource) For managing REST connections.
- [saviynt_salesforce_connection_resource](./resources/saviynt_salesforce_connection_resource) For managing Salesforce connections.
- [saviynt_sap_connection_resource](./resources/saviynt_sap_connection_resource) For managing SAP connections.
- [saviynt_sftp_connection_resource](./resources/saviynt_sftp_connection_resource) For managing SFTP connections.
- [saviynt_unix_connection_resource](./resources/saviynt_unix_connection_resource) For managing Unix connections.
- [saviynt_workday_connection_resource](./resources/saviynt_workday_connection_resource) For managing Workday connections.
- [saviynt_workday_soap_connection_resource](./resources/saviynt_workday_soap_connection_resource) For managing Workday SOAP connections.
- [saviynt_dynamic_attribute_resource](./resources/saviynt_dynamic_attribute_resource) Manages lifecycle (create, update, read, delete) of dynamic attribute
- [saviynt_okta_connection_resource](./resources/saviynt_okta_connection_resource) For managing Okta connections.
- [saviynt_entitlement_type_resource](./resources/saviynt_entitlement_type_resource) For managing entitlement types.
- [saviynt_entitlement_resource](./resources/saviynt_entitlement_resource) For managing entitlements.
- [saviynt_privilege_resource](./resources/saviynt_privilege_resource) For managing privileges.
- [saviynt_enterprise_roles_resource](./resources/saviynt_enterprise_roles_resource) For managing enterprise roles.
- [saviynt_application_data_import_job_resource](./resources/saviynt_application_data_import_job_resource) For managing Application Data Import Job triggers with bulk operations support.
- [saviynt_ws_retry_job_resource](./resources/saviynt_ws_retry_job_resource) For managing WS Retry Job triggers with bulk operations support.
- [saviynt_ws_retry_blocking_job_resource](./resources/saviynt_ws_retry_blocking_job_resource) For managing WS Blocking Retry Job triggers with bulk operations support.
- [saviynt_user_import_job_resource](./resources/saviynt_user_import_job_resource) For managing User Import Job triggers with bulk operations support.
- [saviynt_ecm_job_resource](./resources/saviynt_ecm_job_resource) For managing ECM Job triggers with bulk operations support.
- [saviynt_ecm_sap_user_job_resource](./resources/saviynt_ecm_sap_user_job_resource) For managing ECM SAP User Job triggers with bulk operations support.
- [saviynt_accounts_import_full_job_resource](./resources/saviynt_accounts_import_full_job_resource) For managing Accounts Import Full Job triggers with bulk operations support.
- [saviynt_accounts_import_incremental_job_resource](./resources/saviynt_accounts_import_incremental_job_resource) For managing Accounts Import Incremental Job triggers with bulk operations support.
- [saviynt_schema_role_job_resource](./resources/saviynt_schema_role_job_resource) For managing Schema Role Job triggers with bulk operations support.
- [saviynt_schema_account_job_resource](./resources/saviynt_schema_account_job_resource) For managing Schema Account Job triggers with bulk operations support.
- [saviynt_schema_user_job_resource](./resources/saviynt_schema_user_job_resource) For managing Schema User Job triggers with bulk operations support.
- [saviynt_file_transfer_job_resource](./resources/saviynt_file_transfer_job_resource) For managing File Transfer Job triggers with bulk operations support.
- [saviynt_job_control_resource](./resources/saviynt_job_control_resource) For managing Job Control triggers with bulk operations support.
- [saviynt_file_upload_resource](./resources/saviynt_file_upload_resource) For managing file uploads to Saviynt.
- [saviynt_export_transport_package_resource](./resources/saviynt_export_transport_package_resource) For managing export transport packages.
- [saviynt_import_transport_package_resource](./resources/saviynt_import_transport_package_resource) For managing import transport packages.

## Supported Data Sources

- [saviynt_security_system_datasource](./data-sources/saviynt_security_system_datasource) Retrieves a list of configured security systems filtered by systemname, connection_type, etc.
- [saviynt_endpoints_datasource](./data-sources/saviynt_endpoints_datasource) Retrieves a list of endpoints.
- [saviynt_connection_datasource](./data-sources/saviynt_connections_datasource) Retrieves a list of connections.
- [saviynt_ad_connection_datasource](./data-sources/saviynt_ad_connection_datasource) Retrieves an AD connection.
- [saviynt_adsi_connection_datasource](./data-sources/saviynt_adsi_connection_datasource) Retrieves an ADSI connection.
- [saviynt_db_connection_datasource](./data-sources/saviynt_db_connection_datasource) Retrieves an DB connection.
- [saviynt_entraid_connection_datasource](./data-sources/saviynt_entraid_connection_datasource) Retrieves an EntraID(AzureAD) connection.
- [saviynt_github_rest_connection_datasource](./data-sources/saviynt_github_rest_connection_datasource) Retrieves an Github REST connection.
- [saviynt_rest_connection_datasource](./data-sources/saviynt_rest_connection_datasource) Retrieves an REST connection.
- [saviynt_salesforce_connection_datasource](./data-sources/saviynt_salesforce_connection_datasource) Retrieves a Salesforce connection.
- [saviynt_sap_connection_datasource](./data-sources/saviynt_sap_connection_datasource) Retrieves a SAP connection.
- [saviynt_sftp_connection_datasource](./data-sources/saviynt_sftp_connection_datasource) Retrieves an SFTP connection.
- [saviynt_unix_connection_datasource](./data-sources/saviynt_unix_connection_datasource) Retrieves a Unix connection.
- [saviynt_workday_connection_datasource](./data-sources/saviynt_workday_connection_datasource) Retrieves a Workday connection.
- [saviynt_workday_soap_connection_datasource](./data-sources/saviynt_workday_soap_connection_datasource) Retrieves a Workday SOAP connection.
- [saviynt_dynamic_attribute_datasource](./data-sources/saviynt_dynamic_attribute_datasource) Retrieves Dynamic attribute.
- [saviynt_okta_connection_datasource](./data-sources/saviynt_okta_connection_datasource) Retrieves a Okta connection.
- [saviynt_entitlement_type_datasource](./data-sources/saviynt_entitlement_type_datasource) Retrieves entitlement types.
- [saviynt_entitlement_datasource](./data-sources/saviynt_entitlement_datasource) Retrieves entitlements.
- [saviynt_privilege_datasource](./data-sources/saviynt_privilege_datasource) Retrieves privileges.
- [saviynt_roles_datasource](./data-sources/saviynt_role_datasource) Retrieves roles with filtering options.
