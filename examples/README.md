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
- [saviynt_unix_connection_resource](./resources/saviynt_unix_connection_resource) For managing Unix connections.
- [saviynt_workday_connection_resource](./resources/saviynt_workday_connection_resource) For managing Workday connections.

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
- [saviynt_unix_connection_datasource](./data-sources/saviynt_unix_connection_datasource) Retrieves a Unix connection.
- [saviynt_workday_connection_datasource](./data-sources/saviynt_workday_connection_datasource) Retrieves a Workday connection.
