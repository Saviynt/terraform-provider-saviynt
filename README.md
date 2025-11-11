[![Release build for Saviynt Terraform Provider](https://github.com/saviynt/terraform-provider-saviynt/actions/workflows/release.yml/badge.svg)](https://github.com/saviynt/terraform-provider-saviynt/actions/workflows/release.yml)
<br/><br/>

<a href="https://terraform.io">
    <picture>
        <source media="(prefers-color-scheme: dark)" srcset="assets/hashicorp-terraform-dark.svg">
        <source media="(prefers-color-scheme: light)" srcset="assets/hashicorp-terraform-light.svg">
        <img alt="Terraform logo" title="Terraform" height="60" src="assets/hashicorp-terraform-dark.svg">
    </picture>
</a>

<a href="https://saviynt.com/">
    <img src="assets/s-platform-icon-01.svg" alt="Saviynt logo" title="Saviynt" height="75" />
</a>

# Terraform Provider for Saviynt

The Saviynt Terraform provider empowers you to leverage Terraform's declarative Infrastructure-as-Code (IaC) capabilities to provision, configure, and manage resources within the Saviynt Identity Cloud.

New to Terraform? Check out the [official Terraform introduction by HashiCorp](https://developer.hashicorp.com/terraform/intro) to get up to speed with the basics.

---

##  Requirements

- Terraform version `>= 1.11+`
- Saviynt Identity Cloud instance and credentials.
- **Note**: Write-only attributes (e.g., `password_wo`, `client_secret_wo`) require Terraform version `>= 1.11+`

---

##  Features

Following resources are available for management: 
- [Security System](docs/resources/security_system_resource.md)
- [Endpoint](docs/resources/endpoint_resource.md)
- [Dynamic Attribute](docs/resources/dynamic_attribute_resource.md)
- [Entitlement Type](docs/resources/entitlement_type_resource.md)
- [Enterprise Role](docs/resources/enterprise_roles_resource.md)
- [Entitlements](docs/resources/entitlement_resource.md)
- [Privileges](docs/resources/privilege_resource.md)
- [File Upload](docs/resources/file_upload_resource.md)
- Connections
  - [Active Directory(AD)](docs/resources/ad_connection_resource.md)
  - [REST](docs/resources/rest_connection_resource.md)
  - [ADSI](docs/resources/adsi_connection_resource.md)
  - [Database(DB)](docs/resources/db_connection_resource.md)
  - [EntraID(AzureAD)](docs/resources/entraid_connection_resource.md)
  - [SAP](docs/resources/sap_connection_resource.md)
  - [Salesforce](docs/resources/salesforce_connection_resource.md)
  - [Workday](docs/resources/workday_connection_resource.md)
  - [Workday SOAP](docs/resources/workday_soap_connection_resource.md)
  - [Unix](docs/resources/unix_connection_resource.md)
  - [Github REST](docs/resources/github_rest_connection_resource.md)
  - [Okta](docs/resources/okta_connection_resource.md)
  - [SFTP](docs/resources/sftp_connection_resource.md)
- Jobs
  - [Job Control](docs/resources/job_control_resource.md)
  - [Application Data Import Job](docs/resources/application_data_import_job_resource.md)
  - [WS Retry Job](docs/resources/ws_retry_job_resource.md)
  - [WS Retry Blocking Job](docs/resources/ws_retry_blocking_job_resource.md)
  - [User Import Job](docs/resources/user_import_job_resource.md)
  - [ECM Job](docs/resources/ecm_job_resource.md)
  - [ECM SAP User Job](docs/resources/ecm_sap_user_job_resource.md)
  - [Accounts Import Full Job](docs/resources/accounts_import_full_job_resource.md)
  - [Accounts Import Incremental Job](docs/resources/accounts_import_incremental_job_resource.md)
  - [Schema Role Job](docs/resources/schema_role_job_resource.md)
  - [Schema Account Job](docs/resources/schema_account_job_resource.md)
  - [Schema User Job](docs/resources/schema_user_job_resource.md)
  - [File Transfer Job](docs/resources/file_transfer_job_resource.md)
- Transport Packages
  - [Export Transport Package](docs/resources/export_transport_package_resource.md)
  - [Import Transport Package](docs/resources/import_transport_package_resource.md)
- Ephemerals
  - [File ephemeral resource](docs/ephemeral-resources/file_connector_ephemeral_resource.md)
  - [Env ephemeral resource](docs/ephemeral-resources/env_ephemeral_resource.md)

---

##  Documentation

Check out the [Latest Saviynt Provider Docs](https://registry.terraform.io/providers/saviynt/saviynt/latest/docs) to know more.

---

### Supported Saviynt Versions by Provider

| Supported Saviynt EIC Versions | Terraform Provider Version |
| -------------------------- | ------------------------------ |
| `25.C` | Latest Version: `v0.3.0`<br> Supported Version(s): `v0.2.13` - `v0.3.0`|
| `25.B` | Latest Version: `v0.3.0`<br> Supported Version(s): `v0.2.8` - `v0.3.0`|
| `25.A` | Latest Version: `v0.3.0`<br> Supported Version(s): `v0.2.8` - `v0.3.0`|
| `24.10` | Latest Version: `v0.3.0`<br> Supported Version(s): `v0.2.8` - `v0.3.0`|

--- 

### Attribute Compatibility by EIC Version
The table below shows attributes that are supported in newer versions of Saviynt EIC. If using an older version of Saviynt, some attributes may not work.
Check the table to see which attributes are supported in your version before using them.

| Connector                | Attribute(s) Added                                                                                                    | Present in 25.C | Present in 25.B | Present in 25.A | Present in 24.10 |
| ------------------------ | --------------------------------------------------------------------------------------------------------------------- | ------------------ | --------------- | --------------- | ---------------- |
| **Workday Connector**    | `orgrole_import_payload`                                                                          | Yes                | Yes              | No             | No              |
| **REST Connector**       | `application_discovery_json`, `create_entitlement_json`, `delete_entitlement_json`, `update_entitlement_json`         | Yes                | Yes             | No              | No               |
| **REST Connector**       | `app_type`                                                                                         | Yes                | No              | No              | No               |
| **DB Connector**         | `create_entitlement_json`, `delete_entitlement_json`, `entitlement_exist_json`, `update_entitlement_json`             | Yes                | Yes             | No              | No               |
| **SAP Connector**        | `role_default_date`                                                                                 | Yes                | No              | No              | No               |
| **Unix Connector**       | `server_type`                                                                                      | Yes                | No              | No              | No               |
| **GithubREST Connector** | `status_threshold_config`                                                                           | Yes                | Yes              | Yes              | No               |
| **Security System** | `instant_provisioning`                                                                           | Yes                | Yes              | No              | No               |
| **Entitlement Type** | `enable_entitlement_to_role_sync`                                                                           | Yes                | Yes              | Yes              | No               |
| **Enterprise Role** | `child_roles`                                                                           | Yes                | Yes              | No              | No               |

---

## Write-Only Attributes Management

### Overview

Connection resources provide two approaches for handling sensitive attributes like passwords and tokens:

1. **Sensitive Attributes** (e.g., `password`) - Marked as `sensitive = true`, stored in state but hidden during plan/apply
2. **Write-Only Attributes** (e.g., `password_wo`) - Never stored in state, require `wo_version` for updates

### Sensitive vs Write-Only Attributes
| Aspect               | Sensitive Attributes                                                                           | Write-Only Attributes                                                                                  |
|----------------------|-----------------------------------------------------------------------------------------------|------------------------------------------------------------------------------------------------------|
| **State Storage**    | The value **is stored** in the Terraform state file but is encrypted and redacted in outputs. | The value is **never persisted** to the state file. It is sent to the API and then discarded.          |
| **Plan/Apply Display**| Shows as `(sensitive)` in `plan` and `apply` outputs to prevent exposure.                      | Shows as `(sensitive)` because the value is known but should not be displayed.                         |
| **Update Mechanism** | To update the value, you change the attribute directly in your Terraform configuration.        | Requires a separate, non-sensitive value (e.g., `wo_version`) to be changed (e.g., incremented) to trigger an update. |
| **Drift Detection**  | **Supported.** Terraform can detect if the remote resource's value differs from the state.      | **Not supported.** Because the value is not stored in the state, Terraform cannot detect drift.        |


### Usage Options

#### Option 1: Sensitive Attributes (Standard)
```hcl
resource "saviynt_ad_connection_resource" "example" {
  connection_name = "My_AD_Connection"
  username        = "admin"
  password        = var.db_password  # Sensitive attribute
}
```

#### Option 2: Write-Only Attributes (Maximum Security)
```hcl
resource "saviynt_ad_connection_resource" "example" {
  connection_name = "My_AD_Connection"
  username        = "admin"
  password_wo     = var.db_password  # Write-only attribute
  wo_version      = "v1.1"           # Required for updates
}
```

### Write-Only Attributes by Connector

| Connector | Sensitive Attributes | Write-Only Equivalents |
|-----------|---------------------|------------------------|
| **AD/ADSI** | `password` | `password_wo` |
| **Database** | `password`, `change_pass_json` | `password_wo`, `change_pass_json_wo` |
| **EntraID** | `client_secret`, `access_token`, `connection_json` | `client_secret_wo`, `access_token_wo`, `connection_json_wo` |
| **REST/Github** | `connection_json`, `access_tokens` | `connection_json_wo`, `access_tokens_wo` |
| **Salesforce** | `client_secret`, `refresh_token` | `client_secret_wo`, `refresh_token_wo` |
| **SAP** | `password`, `prov_password` | `password_wo`, `prov_password_wo` |
| **SFTP** | `auth_credential_value`, `passphrase` | `auth_credential_value_wo`, `passphrase_wo` |
| **Unix** | `password`, `passphrase`, `ssh_key` | `password_wo`, `passphrase_wo`, `ssh_key_wo` |
| **Workday** | `password`, `client_secret`, `refresh_token` | `password_wo`, `client_secret_wo`, `refresh_token_wo` |
| **Workday SOAP** | `password`, `change_pass_json`, `connection_json` | `password_wo`, `change_pass_json_wo`, `connection_json_wo` |
| **Okta** | `auth_token` | `auth_token_wo` |

### The `wo_version` Mechanism

When using write-only attributes, the `wo_version` acts as a trigger:

1. **Initial Creation**: Set `wo_version` to any value (e.g., "v1.0")
2. **Updating Credentials**: 
   - Modify the write-only attribute value
   - Change `wo_version` (e.g., "v1.0" → "v1.1")
   - Run `terraform apply`

```hcl
resource "saviynt_db_connection_resource" "database" {
  connection_name = "Production_DB"
  username        = "db_admin"
  password_wo     = var.new_db_password
  wo_version      = "v2.0"  # Increment when password changes
}
```

### Choosing the Right Approach

**Use Sensitive Attributes when:**
- You need drift detection for credentials
- You want simpler credential management

**Use Write-Only Attributes when:**
- Maximum security is required (no state storage)
- Compliance requires credentials never be persisted
- You can manage the `wo_version` trigger mechanism

### Best Practices

1. **Consistent Approach**: Choose one method per environment and stick to it
2. **Variable Usage**: Store sensitive values in Terraform variables:
   ```hcl
   variable "db_password" {
     description = "Database password"
     type        = string
     sensitive   = true
   }
   ```
3. **Version Naming**: Use meaningful `wo_version` identifiers:
   - Semantic: "v1.0", "v1.1", "v2.0"
   - Descriptive: "password-rotation-jan"

4. **Documentation**: Document credential rotation in your infrastructure code

### Important Notes

- **Mutual Exclusivity**: Use either sensitive OR write-only attributes, not both for the same credential
- **State Behavior**: While both attributes don't offer state management, write only attributes don't get stored in the state file as well.
- **Ephemeral Resources**: Consider using [ephemeral resources](#feature-ephemeral-file-credential-resource) for temporary credential management
- **Vault Integration**: Use vault connections for enhanced security when available

## Importing Existing Resources

The Saviynt Terraform provider supports importing existing resources from your Saviynt EIC instance into Terraform state. This allows you to manage existing infrastructure with Terraform without recreating resources.

### Import Methods

There are two ways to import resources:

#### Method 1: Using `terraform import` Command

The traditional approach using the command line:

```bash
terraform import <resource_type>.<resource_name> <import_id>
```

**Example:**
```bash
terraform import saviynt_security_system_resource.example SYSTEM1
```

#### Method 2: Using Import Blocks (Recommended)

The modern approach using import blocks in your Terraform configuration. This method also supports automatic configuration generation.

**Step 1:** Add an import block to your `.tf` file:
```hcl
import {
  to = saviynt_security_system_resource.example
  id = "SYSTEM1"
}
```

**Step 2:** Generate configuration automatically:
```bash
terraform plan -var-file=dev.tfvars -generate-config-out=generated.tf
```

This command will:
- Create a `generated.tf` file with the resource configuration
- Show you what will be imported
- Generate all the necessary attributes based on the current state in Saviynt

**Step 3:** Review and customize the generated configuration:
```bash
# Review the generated configuration
cat generated.tf

# Move relevant parts to your main configuration files
# Remove the import block after successful import
```

**Step 4:** Apply the import:
```bash
terraform apply -var-file=dev.tfvars
```

### Import ID Reference

Each resource type requires a specific import ID format:

| Resource Type        | Import ID Format   | Example                                                        |
| -------------------- | ----------------- | -------------------------------------------------------------- |
| Security System      | `systemname`     | `terraform import saviynt_security_system_resource.example SYSTEM1`     |
| Endpoint             | `endpoint_name`    | `terraform import saviynt_endpoint_resource.example ENDPOINT1`          |
| Connection Resources | `connection_name` | `terraform import saviynt_ad_connection_resource.example AD_CONN1`      |
| Dynamic Attributes   | `endpoint`        | `terraform import saviynt_dynamic_attribute_resource.example ENDPOINT1` |
| Entitlement Type     | `endpoint_name:entitlement_name` | `terraform import saviynt_entitlement_type_resource.example ENDPOINT1:ENTTYPE1` |
| Enterprise Role     | `role_name` | `terraform import saviynt_enterprise_roles_resource.example role_name` |
| Entitlement     | `endpoint:entitlement_type:entitlement_value` | `terraform import saviynt_entitlement_resource.example ENDPOINT1:ENTTYPE1:ENT1` |
| Privilege     | `endpoint:entitlement_type` | `terraform import saviynt_privilege_resource.example ENDPOINT1:ENTTYPE1` |

### Example: Complete Import Workflow

**1. Add import block to main.tf:**
```bash
cat >> main.tf << EOF
import {
  to = saviynt_security_system_resource.prod_system
  id = "PROD_SYSTEM_01"
}
EOF
```

**2. Generate configuration:**
```bash
terraform plan -var-file=prod.tfvars -generate-config-out=generated.tf
```
**2. Generate configuration:**
```bash
terraform plan -var-file=prod.tfvars -generate-config-out=generated.tf
```
> **Note:**<br>
> **Role User Task Management:**  
> When users are added to roles, Saviynt creates tasks that require manual completion. After successful user additions:  
> 1. Navigate to **Pending Tasks** section in the Saviynt UI  
> 2. Approve/complete the user assignment tasks manually  
> 3. After task completion, running `terraform plan` may show drift due to state changes  
> 4. Use `terraform import` to sync the current state:  
>    ```
>    terraform import saviynt_enterprise_roles_resource.resource_name "role-name"
>   

**3. Review and move configuration:**
```bash
cat generated.tf
# Move relevant parts to your organized .tf files
```

**3. Review and move configuration:**
```bash
cat generated.tf
# Move relevant parts to your organized .tf files
```

**4. Remove import block and apply:**
```bash
# Edit main.tf to remove the import block
terraform apply -var-file=prod.tfvars
```

### Best Practices for Importing

1. **Use Import Blocks**: Prefer import blocks over the `terraform import` command for better workflow and automatic config generation.

2. **Generate Configuration First**: Always use `-generate-config-out` to create initial configuration, then customize as needed.

3. **Review Generated Config**: The generated configuration includes all current attributes. Remove any that you don't want Terraform to manage.

4. **Handle Sensitive Data**: Some attributes (like passwords) won't be imported due to security. You'll need to set these manually or use ephemeral resources.

5. **Test in Non-Production**: Always test imports in a development environment first.

6. **Backup State**: Create a backup of your Terraform state before importing:
   ```bash
   cp terraform.tfstate terraform.tfstate.backup
   ```

### Troubleshooting Import Issues

- **Invalid Import ID**: Ensure you're using the correct format for each resource type
- **Resource Not Found**: Verify the resource exists in Saviynt with the exact name/ID
- **Permission Issues**: Ensure your Saviynt user has read access to the resource
- **State Conflicts**: If importing fails due to existing state, use `terraform state rm` to remove conflicting resources first

---

## Getting started

Before installing the provider, ensure that you have the following dependencies installed:

### **1. Install Terraform**  
Terraform is required to use this provider. Install Terraform using one of the following methods:

#### **For macOS (using Homebrew)**
```sh
brew tap hashicorp/tap
brew install hashicorp/tap/terraform
```

#### **For Windows (using chocolatey)**
```sh
choco install terraform
```

#### **For Manual installation or other platforms**
Visit [Terraform Installation](https://developer.hashicorp.com/terraform/install) for installation instructions.

<!-- ### 2. Install Go

#### **For macOS (using Homebrew)**
```sh
brew install go
```
#### **For Windows (using chocolatey)**
```sh
choco install golang
```
#### **For Manual installation or other platforms**
Visit [Go Setup](https://go.dev/doc/install) for installation instructions.

### 3. Finding the GOBIN Folder Path

#### **For macOS**

To check the `GOBIN` path, run the following command in your terminal:

```sh
go env GOBIN
```

If it doesn't return anything, Go will use the default: `$GOPATH/bin`.

To explicitly set `GOBIN`, you can update your shell configuration file (e.g., `~/.zshrc`, `~/.bashrc`, etc.). Below are steps for `~/.zshrc`:

1. Open the file in your default editor:
   ```sh
   open ~/.zshrc
   ```

2. Add the following lines at the end of the file:
   ```sh
   export GOBIN=$HOME/go/bin
   export PATH=$PATH:$GOBIN
   ```

3. Apply the changes:
   ```sh
   source ~/.zshrc
   ```

4. Confirm the value:
   ```sh
   go env GOBIN
   ```

You should now see the path as `$HOME/go/bin`.

---

#### **For Windows**

To check the current `GOBIN` value, run the following in **Command Prompt** or **PowerShell**:

```sh
go env GOBIN
```

If it's empty, Go defaults to `%GOPATH%\bin`. To check `GOPATH`, run:

```sh
go env GOPATH
```

> The default path is usually: `C:\Users\<YourUsername>\go\bin`

To explicitly set `GOBIN`, follow these steps:

1. Open the **Start Menu** and search for **"Environment Variables"**.
2. Click **"Edit the system environment variables"**.
3. In the **System Properties** window, click **"Environment Variables…"**.
4. Under **User variables**, click **"New…"** and enter:
   - **Variable name**: `GOBIN`
   - **Variable value**: `C:\Users\<YourUsername>\go\bin` (or your desired path)
5. Add `GOBIN` to your system `Path`:
   - Under **User variables**, select the `Path` variable and click **Edit**.
   - Click **New** and add: `C:\Users\<YourUsername>\go\bin`
6. Click **OK** to save and exit all dialogs.
7. Restart your terminal or system.

To verify:

```sh
go env GOBIN
```

You should now see the configured GOBIN path.


> **Note:** Save the GOBIN path for later use. -->


<!-- ### 4. Download the Binary
Copy the provider binary from provider directory to the Go bin directory: 

```sh
cp provider/terraform-provider-saviynt_v0.1.3 <GOBIN PATH>/terraform-provider-saviynt
chmod +x GOBIN/terraform-provider-saviynt

```
Replace `<GOBIN PATH>` with your actual GOBIN path where the go bin folder is located. -->

<!-- ### 4. Download the Binary

Inside the `provider` directory, you will find multiple `.zip` files for different operating systems (e.g., macOS, Windows, Linux). Choose the appropriate binary for your OS, extract it, and copy the provider binary to your Go bin directory.

For example, on macOS

```sh
# Unzip the appropriate binary
unzip terraform-provider-saviynt_v0.1.6_darwin_amd64.zip -d provider/

# Copy the binary to your GOBIN directory
cp provider/terraform-provider-saviynt_v0.1.6 <GOBIN PATH>/terraform-provider-saviynt

# Make it executable
chmod +x <GOBIN PATH>/terraform-provider-saviynt
```

> **Note:** Replace `<GOBIN PATH>` with your actual GOBIN path. If you're unsure, run:
```sh
go env GOBIN
```

### - macOS Security Warning Workaround

When using the downloaded Terraform provider binary on macOS, you might encounter a security warning like:

> `"Apple is not able to verify that it is free from malware that could harm your Mac or compromise your privacy. Don’t open this unless you are certain it is from a trustworthy source.`

This happens because macOS restricts the execution of unsigned binaries.  
To work around this, you can follow either of the options below:

####  Option 1: Allow via System Settings

1. Try running the provider binary once to trigger the security warning.
2. Open **System Settings** → **Privacy & Security**.
3. Scroll down to the **Security** section.
4. You’ll see a message similar to:
   > `"terraform-provider-saviynt" was blocked from use because it is not from an identified developer.`
5. Click **"Allow Anyway"**.
6. Re-run your Terraform command.
7. If prompted again, click **"Open"** to allow execution.

####  Option 2: Allow via Terminal

You can also manually remove the quarantine attribute using the Terminal:

```sh
xattr -d com.apple.quarantine <path-to-binary>/terraform-provider-saviynt
``` -->

<!-- ### 5. Configure `.terraformrc` or `terraform.rc`

Create the file at:

- **macOS/Linux**: `~/.terraformrc`
- **Windows**: `%APPDATA%\terraform.rc`

```hcl
provider_installation {
  dev_overrides {
    "<PROVIDER SOURCE PATH>" = "<GOBIN PATH>"
  }
  direct {}
}
```

Note: If there is an error in Windows while running Terraform later, the user can append the `"<PROVIDER SOURCE PATH>" = "<GOBIN PATH>"` with `"<PROVIDER SOURCE PATH>" = "<GOBIN PATH>/terraform-provider-saviynt"`, while replacing the `<PROVIDER SOURCE PATH>` and `<GOBIN PATH>` with the respective paths. -->


### 2. Getting Started with Terraform

Follow the steps below to start using the Saviynt Terraform Provider:

---

#### **Step 1: Create a Terraform Project Folder**

```sh
mkdir saviynt-terraform-demo
cd saviynt-terraform-demo
```

---

#### **Step 2: Initialize a Terraform Configuration File**

Create a file named `main.tf` and define your provider and resources:

````hcl
terraform {
  required_providers {
    saviynt = {
      source = "saviynt/saviynt"
      version = "x.x.x"
    }
  }
}

provider "saviynt" {
  server_url  = "https://example.saviyntcloud.com"
  username   = "username"
  password   = "password"
}
````
<!-- 
Replace the `<PROVIDER SOURCE PATH>` with your provider path. The configuration should look similar to `registry.terraform.io/local/saviynt`. -->

---

#### **Step 3: Define Input Variables**

Create a file called `variables.tf` to declare your input variables:

```
variable "server_url" {
  description = "Saviynt instance base URL"
  type        = string
}

variable "username" {
  description = "Username"
  type        = string
}

variable "password" {
  description = "Password"
  type        = string
  sensitive   = true
}
```
<!-- 
> You can refer to a sample `variables.tf` file in the `resources/connections/` folder for guidance. -->

---

#### **Step 4: Create a `terraform.tfvars` File**

This file contains the actual values for the declared variables:

```hcl
server_url   = "https://example.saviyntcloud.com"
username = "username"
password = "password"
```

> This file is automatically used by Terraform during plan and apply.

You can also name the file `prod.tfvars`, `dev.tfvars`, etc. and explicitly reference it:

```sh
terraform apply -var-file="terraform.tfvars"
```

> Make sure to add `terraform.tfvars` to your `.gitignore` if it contains sensitive information:

```sh
echo "terraform.tfvars" >> .gitignore
```

**Important:** When using `tfvars`, you must refer to the variables using `var.<variable_name>` syntax in your `.tf` files.  
For example, in your provider configuration or resource definitions:

```
provider "saviynt" {
  server_url = var.server_url
  username   = var.username
  password = var.password
}
```

---

#### **Step 5: Validate & Apply Configuration**

Initialise the provider:

```sh
terraform init
```

Validate the syntactic correctness of the .tf file:

```sh
terraform validate
```

Plan the changes:

```sh
terraform plan
```

Apply the changes:

```sh
terraform apply -var-file=terraform.tfvars
```

That's it! You've now set up and run your first configuration using the Saviynt Terraform Provider.

---

##  Usage

Here's an example of defining and managing a resource:

```
resource "saviynt_security_system_resource" "sample" {
  systemname          = "sample_security_system"
  display_name        = "sample security system"
  hostname            = "sample.system.com"
  port                = "443"
  access_add_workflow = "sample_workflow"
}
```
Here's an example of using the data source block:
```
data "saviynt_security_systems_datasource" "all" {
  connection_type = "REST"
  max             = 10
  offset          = 0
}

output "systems" {
  value = data.saviynt_security_systems_datasource.all.results
}
```

You can find the starter templates to define each supported resource type in the ```examples/``` folder. To know the differnt types of arguments that can be passed for each resource, user can refer to the ```docs/``` folder.

For inputs that require JSON config, you can give the values as in the given example:
```sh
create_account_json = jsonencode({
    "cn" : "$${cn}",
    "displayname" : "$${user.displayname}",
    "givenname" : "$${user.firstname}",
    "mail" : "$${user.email}",
    "name" : "$${user.displayname}",
    "objectClass" : ["top", "person", "organizationalPerson", "user"],
    "userAccountControl" : "544",
    "sAMAccountName" : "$${task.accountName}",
    "sn" : "$${user.lastname}",
    "title" : "$${user.title}"
  })
```
As in the above example, to pass special characters like `$`, we have to use `$$` instead and for json data, use the `jsonencode()` function to properly pass the data using terraform.

For mutliline string inputs, use `trimspace` to pass the values. Below is an example:
```sh
account_import_payload = trimspace(
    <<EOF
    <bsvc:Get_Workers_Request bsvc:version="$${API_VERSION}">
      <bsvc:Request_Criteria>
          <bsvc:Exclude_Inactive_Workers>false</bsvc:Exclude_Inactive_Workers>
          <bsvc:Exclude_Employees>false</bsvc:Exclude_Employees>
          <bsvc:Exclude_Contingent_Workers>false</bsvc:Exclude_Contingent_Workers>
          $${INCREMENTAL_IMPORT_CRITERIA}
      </bsvc:Request_Criteria>
      <bsvc:Response_Filter><bsvc:Page>$${PAGE_NUMBER}</bsvc:Page><bsvc:Count>$${PAGE_SIZE}</bsvc:Count></bsvc:Response_Filter>
      <bsvc:Response_Group>
          <bsvc:Include_Reference>true</bsvc:Include_Reference>
          <bsvc:Include_Personal_Information>true</bsvc:Include_Personal_Information>
      </bsvc:Response_Group>
  </bsvc:Get_Workers_Request>
  EOF
)
```

---

## Credential Management Best Practices
To ensure secure handling of sensitive credentials, follow these best practices:
1. **Use Vault-backed Secrets**   : Externalize sensitive values using a secure secrets manager such as HashiCorp Vault. This avoids hardcoding secrets in .tf files or storing them in Terraform state.
2. **Prefer Environment Variables**   : If Vault integration is not available, use environment variables (e.g., TF_VAR_password) to pass secrets instead of storing them in source files.
3. **Use Ephemeral Resources for Sensitive Data**
   Handle all sensitive data using **ephemeral resources** — those created only when needed and destroyed immediately afterward. Avoid long-term storage of secrets in:
   - Terraform state files (`terraform.tfstate`)
   - Version control systems (e.g., Git)
   - Local plaintext configuration files

---

## Feature: Ephemeral File Credential Resource

The **Ephemeral File Credential Resource** is a transient Terraform resource that provides temporary, in-memory credentials to other connector resources by reading values from a local JSON file at apply time. This allows secure and flexible provisioning without persisting sensitive data in the Terraform state.

### Supported Connectors

The following connectors are supported and can consume credentials provided by this resource:

- **AD**: `password`
- **ADSI**: `password`
- **DB**: `password`, `change_pass_json`
- **EntraId**: `client_secret`, `access_token`, `azure_mgmt_access_token`, `windows_connector_json`, `connection_json`
- **Github REST**: `connection_json`, `access_tokens`
- **REST**: `connection_json`
- **Salesforce**: `client_secret`, `refresh_token`
- **SAP**: `password`, `prov_password`
- **SFTP**: `auth_credential_value`, `passphrase`
- **Unix**: `password`, `passphrase`, `ssh_key`, `ssh_pass_through_password`, `ssh_pass_through_sshkey`, `ssh_pass_through_passphrase`
- **Workday**: `password`, `client_secret`, `refresh_token`
- **Workday SOAP**: `password`, `change_pass_json`, `connection_json`
- **Okta**: `auth_token`

### Usage

The ephemeral credential resource reads from a local JSON file structured with the required fields for the target connector. These fields are then dynamically injected into the respective connector resources during the `apply` phase.

> **Note:** This resource is ephemeral and does not store any state. It is designed for use cases where credentials must remain local and transient.

### Security Considerations

- Ensure the credential file is secured and not committed to version control.
- Avoid using this resource in long-lived plans, as it relies on local files that may change or expire.

## Feature: Ephemeral Env Credential Resource

The **Ephemeral Env Credential Resource** is a transient Terraform resource that provides temporary, in-memory credentials to other connector resources by reading values from the environment variables at apply time. This allows secure and flexible provisioning without persisting sensitive data in the Terraform state.

### Supported Connectors

The following connectors are supported and can consume credentials provided by this resource:

- **AD**: `password`
- **ADSI**: `password`
- **DB**: `password`, `change_pass_json`
- **EntraId**: `client_secret`, `access_token`, `azure_mgmt_access_token`, `windows_connector_json`, `connection_json`
- **Github REST**: `connection_json`, `access_tokens`
- **REST**: `connection_json`
- **Salesforce**: `client_secret`, `refresh_token`
- **SAP**: `password`, `prov_password`
- **SFTP**: `auth_credential_value`, `passphrase`
- **Unix**: `password`, `passphrase`, `ssh_key`, `ssh_pass_through_password`, `ssh_pass_through_sshkey`, `ssh_pass_through_passphrase`
- **Workday**: `password`, `client_secret`, `refresh_token`
- **Workday SOAP**: `password`, `change_pass_json`, `connection_json`
- **Okta**: `auth_token`

### Usage

The ephemeral credential resource reads from environment variables. These fields are then dynamically injected into the respective connector resources during the `apply` phase.

> **Note:** This resource is ephemeral and does not store any state. It is designed for use cases where credentials must remain local and transient.

## Feature: `authenticate` Toggle for All Data Source

The Saviynt datasources now have a required boolean flag `authenticate` to control the visibility of sensitive data.

### Purpose

This feature is designed to help prevent potential sensitive data from appearing in Terraform state files or CLI output during `plan` and `apply` when a datasource is called.

### Behavior

- When `authenticate = false`:
  - The provider will **omit** all the attributes from state file.
- When `authenticate = true`:
  - All attributes will be returned as usual with sensitive still not visible.

### Example Usage

```hcl
data "saviynt_rest_connection_datasource" "example" {
  connection_name = "Terraform_REST_Connector"
  authenticate    = false
}
```

---

<!-- ##  Examples

Examples are available for all resources. Follow the following steps to try out the examples

1. Uncomment the code block corresponding to the object for which you want to try an operation (say create ad connection) in [provider.tf](provider.tf)
2. Navigate to the file corresponding to the resource that you uncommented (the uncommented code block contains the path) and update the values.
3. Review the changes using the following
   ```
   terraform plan
   ```
5. If everything works fine, apply the changes using the following
   ```
   terraform apply -var-file=<main tf file>
   ``` -->

<!-- --- -->
## Known Limitations

The following limitations are present in the latest version of the provider. These are being prioritized for resolution in the upcoming release alongside new feature additions:

### 1. All Resource objects
 - `terraform destroy` is not supported for resources such as:
    - Security System
    - Endpoint
    - Connectors
    - Enterprise Role
    - Entitlement
    - Entitlement Type
    - Job Control
    - Transport Packages (Export/Import)
    - File Upload

### 2. Endpoints
- **State management is not supported** for the following attributes:
  - `Owner`
  - `ResourceOwner`
  - `OutOfBandAccess`

- For `resource_owner_type` and `owner_type`, the allowed values are:
  - `User`
  - `Usergroup`

- The `MappedEndpoints` field **cannot be configured during endpoint creation**; it must be managed after the endpoint is created.

- The `RequestableRoleType` attribute **can only be set during updates**, since the role must be assigned to the endpoint beforehand.

- For `saviynt_endpoint_resource.requestable_role_types.request_option`, the supported values for proper state tracking are:
  - `DropdownSingle`
  - `Table`
  - `TableOnlyAdd`

- The following service account settings are **not currently configurable via Terraform**:
  - `Disable Remove Service Account`
  - `Disable Modify Service Account`
  - `Disable New Account Request if Account Exists`

### 3. Connections
- **State management** is not supported for the following attributes due to their sensitive nature:
  - **AD**: `password`
  - **ADSI**: `password`
  - **DB**: `password`, `change_pass_json`
  - **EntraId**: `access_token`, `azure_mgmt_access_token`, `client_secret`, `windows_connector_json`, `connection_json`
  - **Github REST**: `connection_json`, `access_tokens`
  - **REST**: `connection_json`
  - **Salesforce**: `client_secret`, `refresh_token`
  - **SAP**: `password`, `prov_password`
  - **SFTP**: `auth_credential_value`, `passphrase`
  - **Unix**: `password`, `passphrase`, `ssh_key`, `ssh_pass_through_password`, `ssh_pass_through_sshkey`, `ssh_pass_through_passphrase`
  - **Workday**: `password`, `client_secret`, `refresh_token`
  - **Workday SOAP**: `password`, `change_pass_json`, `connection_json`
  - **Okta**: `auth_token`
- **SFTP Connection**: Requires manual configuration of "Connector Version" field in Saviynt UI after creation (see [Troubleshooting Guide](#9-sftp-connection-post-creation-configuration))


### 4. Dynamic Attributes
- For `saviynt_dynamic_attribute_resource.dynamic_attributes.attribute_type`, the supported values for proper state tracking are:
  - `NUMBER`
  - `STRING`
  - `ENUM`
  - `BOOLEAN`
  - `MULTIPLE SELECT FROM LIST`
  - `MULTIPLE SELECT FROM SQL QUERY`
  - `SINGLE SELECT FROM SQL QUERY`
  - `PASSWORD`
  - `LARGE TEXT`
  - `CHECK BOX`
  - `DATE`

### 5. Entitlement Types

- **State management is not supported** for the following attributes:
  - `start_date_in_revoke_request`
  - `start_end_date_in_request`
  - `allow_remove_all_entitlement_in_request`

- For `request_option`, users need to set values using the following mapping:
  - `"SHOW_BUT_NOTREUESTABLESINGLE"` → `"Request Form NotRequestable Single"`
  - `"SHOW_BUT_NOTREUESTABLEMULTIPLE"` → `"Request Form NotRequestable Multiple"`
  - `"NONE"` → `"Request Form None"`
  - `"SINGLE"` → `"Request Form Single"`
  - `"MULTIPLE"` → `"Request Form Multiple"`
  - `"TABLE"` → `"Request Form Table"`
  - `"FREEFORMTEXT"` → `"Request Form Free From Text"`
  - `"TABLENOREMOVE"` → `"Request Form Table No Remove"`
  - `"RADIOBUTN"` → `"Request Form Radio Button"`
  - `"CHECKBOXN"` → `"Request Form CheckBox"`
  - `"READONLYTABLE"` → `"Request Form Read Only Table"`

- For `show_ent_type_on`, use the following values:
  - `"0"` → All requests
  - `"1"` → Standard Account Requests
  - `"2"` → Service Account Requests

- For `hierarchy_required`, use the following values:
  - `"0"` → Not required
  - `"1"` → Required
  
### 6. Role 

- Ranks of owners(`owners.rank`) should not be updated from Terraform as the owners get removed and added again which might affect any related workflows.

- **State management is not supported** for the following attributes:
  - `check_sod`
  - `level`
  - `requestor`

- **6 types of roles available for role datasource:**
  - `"ENABLER"`
  - `"TRANSACTIONAL"`
  - `"FIREFIGHTER"`
  - `"ENTERPRISE"`
  - `"APPLICATION"`
  - `"ENTITLEMENT"`

- **6 types of criticality levels** (for `sox_critical`, `sys_critical`, `privileged`, `confidentiality`, `risk`):
  - `"None"`
  - `"Very Low"`
  - `"Low"`
  - `"Medium"`
  - `"High"`
  - `"Critical"`

- **2 types of status:**
  - `"Active"`
  - `"Inactive"`

### 7. Entitlement
- `confidentiality` is the attribute to set the field for `Financial` on the UI.

- **6 types of criticality levels** use the following values for `sox_critical`, `sys_critical`, `privileged`, `confidentiality`, `risk`:
  - `0` → `"None"`
  - `1` → `"Very Low"`
  - `2` → `"Low"`
  - `3` → `"Medium"`
  - `4` → `"High"`
  - `5` → `"Critical"`

- **2 types of status:** use the following values for status:
  - `0` → `"None"`
  - `1` → `"Active"`
  - `2` → `"Inactive"`
  - `3` → `"Decommission Active"`
  - `4` → `"Decommission Active"`

- **Values for module:**
    - `None`
    - `IT`
    - `Business`
    - `N/A`

- **Values for access:**
    - `None`
    - `Read-Only`
    - `Update`
    - `Delete`

### 8. Jobs
- **State management is not supported** for job resources created or modified outside of Terraform (e.g., via Saviynt UI)
- **Import functionality is not available** for existing job triggers - jobs must be created through Terraform
- **Manual job modifications** made in Saviynt UI will not be detected by Terraform 
- Users can only **create**, **update**, and **delete** job triggers through Terraform for the time being

### 9. Transport Packages
- **Export and Import Transport Packages**: Exporting or importing transport packages to or from local storage is not supported. Packages are exported within the EIC environment, and the import path must reference a location in EIC rather than a local directory.

---


## Troubleshooting Guide – Breaking Changes, API Errors & Configuration Tips

This section outlines recent changes that may cause issues during Terraform runs and provides guidance on how to resolve them.

### 1. Breaking Change: Removal of connection_type Attribute
The connection_type attribute has been removed from all connector resources.
**Impact**:
If your Terraform configuration includes this attribute, it will now result in an error during `terraform plan` or `apply`.

**Example Error**:
```
╷
│ Error: Unsupported argument
│
│   on provider.tf line 33, in resource "saviynt_ad_connection_resource" "AD1":
│   33:   connection_type = "AD"
│
│ An argument named "connection_type" is not expected here.
```

**Action Required**:
Manually remove all instances of connection_type from your Terraform resource blocks related to connector resources.

### 2. Security Notice: API Access Restriction Based on SAV Role
If the **Restrict API access based on SAV Role** option is enabled (`Settings > API` in ECM), you might encounter the following error during resource creation:

**Example Error**:
```
│ Error: API Read Failed In Create Block
│
│  with saviynt_endpoint_resource.example,
│  on generated.tf line 5, in resource "saviynt_endpoint_resource" "example":
│   5: resource "saviynt_endpoint_resource" "example" {
│
│ Error: 412 Precondition Failed
```
Recommendation:
Ensure that **Restrict API access based on SAV Role** is disabled for successful provisioning via Terraform.

### 3. API Compatibility Warning: `readlabels` Settings
Changes to the following settings under `Settings > Configuration Files>externalconfig.properties` can affect the structure of API responses:
```properties
users.readlabels=true
endpoints.readlabels=true
entitlements.readlabels=true
roles.readlabels=true
```
**Impact**:
- When these properties are set to `false`, API responses return machine-friendly field names like:
```json
"customproperty10": "Project",
"customproperty12": "Role"
```
- When set to `true` (default and recommended), responses return human-readable keys:
```json
"Custom Property 10": "Project",
"Custom Property 11": "Team"
```

**Recommendation**:
Keep all `readlabels` settings set to `true` to ensure compatibility with Terraform provider expectations and avoid field mapping issues.

---


### 4. User Operation Failures in Role Resources

**Issue**: When creating roles with users that don't exist or are inactive, you may see warnings about user operation failures.

**Symptoms**:
```
Warning: User Operation Failed During Role Creation

Role 'my-role' was created successfully, but some users failed to be added: 
failed to add 1/3 users to role my-role: User nonexistent_user: 
[{HTTP error while adding user: user nonexistent_user inactive or not found HTTP Error}]

The role exists and successful users have been added. Please fix the failed 
users and run terraform apply again to add the remaining users.
```

**Resolution Options**:

**Option 1: Fix Users and Reapply (Recommended)**
1. Verify user existence in Saviynt UI
2. **Add missing users** to Saviynt or **fix usernames** in Terraform config
3. Run terraform apply again to add the corrected users:
   ```bash
   terraform apply
   ```

**Option 2: If You Encounter "Delete Not Supported" Error**
If you accidentally tried to remove users from config and got the delete error:
```bash
terraform untaint saviynt_enterprise_roles_resource.resource_name
terraform apply
```

**Option 3: Reset State for Clean Retry**
```bash
terraform untaint saviynt_enterprise_roles_resource.resource_name
terraform apply
```

**Option 4: Import Existing Role State**
```bash
terraform state rm saviynt_enterprise_roles_resource.resource_name
terraform import saviynt_enterprise_roles_resource.resource_name "role-name"
terraform apply
```

**Prevention**:
- Validate usernames in Saviynt UI before adding to Terraform config
- Use data sources to verify user existence before role creation
- Check user status (active/inactive) in Saviynt

**Note**: The provider handles user failures gracefully - roles are created successfully even if some users fail, allowing incremental fixes.

### 5. Resource Dependency Issues

**Issue**: When creating multiple resources where one depends on another, you may encounter errors if resources are created in the wrong order.

**Example Error**:
```
╷
│ Error: API Create Failed In Create Block
│
│   with saviynt_endpoint_resource.EP1,
│   on provider.tf line 164, in resource "saviynt_endpoint_resource" "EP1":
│  164: resource "saviynt_endpoint_resource" "EP1" {
│
│ API error In CreateEndpoint Block: systemname not found
```

**Resolution**:
Use the `depends_on` meta-argument to explicitly define resource dependencies:

```hcl
resource "saviynt_security_system_resource" "system1" {
  systemname   = "TF_Security_System"
  display_name = "TF_Security_System"
  # ... other attributes
}

resource "saviynt_endpoint_resource" "EP1" {
  endpoint_name    = "TF_Endpoint"
  security_system  = "TF_Security_System"
  # ... other attributes
  
  depends_on = [saviynt_security_system_resource.system1]
}
```

**Note**: For more information on meta-arguments available in Terraform, see the [official Terraform documentation on meta-arguments](https://developer.hashicorp.com/terraform/language/meta-arguments).

### 6. Transport Package Version Attributes

**Issue**: You need to force a transport package operation to run again even when there are no configuration changes.

**Background**: Transport package resources include version attributes (`export_package_version` and `import_package_version`) that serve as triggers to force API calls when needed.

**Use Cases**:
- **Re-import for testing**: When you want to test the same package import multiple times
- **Force refresh**: When you suspect the previous operation didn't complete successfully
- **Scheduled operations**: When you want to trigger regular exports/imports as part of maintenance

**Solution**:

**For Export Transport Package:**
```hcl
resource "saviynt_export_transport_package_resource" "example" {
  export_online          = "false"
  export_path            = "/saviynt_shared/exports"
  sav_roles              = ["ROLE_ADMIN"]
  export_package_version = "1.1"  # Change this to force re-export
}
```

**For Import Transport Package:**
```hcl
resource "saviynt_import_transport_package_resource" "example" {
  package_path           = "/saviynt_shared/package.zip"
  import_package_version = "2.0"  # Change this to force re-import
}
```

**Best Practices**:
- Use semantic versioning (1.0, 1.1, 2.0) for clarity
- Document the reason for version changes in comments
- Increment version when you want to force the operation to run again
- Use meaningful version identifiers like "maintenance-jan-2024"

### 7. File Upload Version Attributes

**Issue**: You need to force file upload to run again even when there are no configuration changes.

**Background**: File upload resources include version attributes (`file_version`) that serve as triggers to force file upload when needed.

**Use Cases**:
- **Re-upload for testing**: When you want to test the same file upload multiple times
- **Force refresh**: When the file content has changed but the file path remains the same
- **File updates**: When you need to upload a modified version of the same file

**Solution**:

**For File Upload Resource:**
```hcl
resource "saviynt_file_upload_resource" "example" {
  file_path     = "/path/to/file.csv"
  path_location = "Datafiles"
  file_version  = "1.1"  # Change this to force re-upload
}
```

**Best Practices**:
- Use semantic versioning (1.0, 1.1, 2.0) for clarity
- Document the reason for version changes in comments
- Increment version when you want to force the upload to run again
- Use meaningful version identifiers like "data-update-jan-2024"

### 8. Job Control Version Attributes

**Issue**: You need to force job execution to run again even when there are no configuration changes.

**Background**: Job control resources include version attributes (`run_job_version`) that serve as triggers to force job execution when needed.

**Use Cases**:
- **Re-run for testing**: When you want to test the same job execution multiple times
- **Force refresh**: When you suspect the previous job execution didn't complete successfully

**Solution**:

**For Job Control Resource:**
```hcl
resource "saviynt_job_control_resource" "example" {
  run_jobs = [
    {
      job_name        = "WSRetryJob"
      trigger_name    = "ws_retry_trigger_prod"
      job_group       = "Utility"
      run_job_version = "1.1"  # Change this to force re-run
    },
    {
      job_name        = "UserImportJob"
      trigger_name    = "user_import_trigger"
      job_group       = "Import"
      run_job_version = "2.0"  # Change this to force re-run
    }
  ]
}
```

**Best Practices**:
- Use semantic versioning (1.0, 1.1, 2.0) for clarity
- Document the reason for version changes in comments
- Increment version when you want to force the operation to run again
- Use meaningful version identifiers like "maintenance-jan-2024"

### 9. SFTP Connection Post-Creation Configuration

**Issue**: SFTP connections require manual configuration of the `Select Connector version` field after creation via Terraform.

**Background**: The SFTP connector has a mandatory field called "Connector Version" that cannot be set during the initial Terraform resource creation and must be configured manually in the Saviynt UI.

**Required Action**:
After successfully creating an SFTP connection resource with Terraform:

1. **Navigate to Saviynt UI**: Go to **Admin** → **Identity Repository** → **Connections**
2. **Find Your Connection**: Locate the SFTP connection created by Terraform
3. **Edit Connection**: Click on the connection name to open the configuration
4. **Set Connector Version**: Configure the "Select Connector version" field with the value `SFTPConnector::1.0`
5. **Save Changes**: Save and test the connection configuration

**Example Workflow**:
```bash
# 1. Create SFTP connection via Terraform
terraform apply

# 2. Manually configure Connector Version in Saviynt UI
# 3. Connection is now fully functional
```

**Important Notes**:
- This is a **one-time configuration** per SFTP connection resource
- The connection will not function properly until the Connector Version is set
- This manual step is required due to API limitations for the SFTP connector
- Future updates to the connection via Terraform will not affect the Connector Version field

___

### Summary
| Change Area              | Description                                          | Action Required                                      |
|--------------------------|------------------------------------------------------|------------------------------------------------------|
| `connection_type` Removal| Deprecated from connector resources                 | Remove from your resource configuration              |
| SAV Role Restriction     | May result in 412 errors if enabled                | Adjust SAV role or disable restriction               |
| `readlabels` Settings    | Alters field naming in API response                | Keep values as `true`                                |
| User Operation Failures  | Role creation succeeds but user assignment may fail | Validate users exist and are active before applying |
| Resource Dependencies    | Errors when resources created in wrong order        | Use `depends_on` meta-argument for explicit dependencies |
| Transport Package Versions | Version attributes force API calls when unchanged | Increment version to trigger re-export/re-import operations |
| File Upload Versions    | Version attributes force file upload when unchanged | Increment `file_version` to trigger file re-upload |
| Job Control Versions    | Version attributes force job execution when unchanged | Increment `run_job_version` to trigger job re-execution |
| SFTP Connector Version  | Manual UI configuration required after creation     | Set `Select Connector version` field in Saviynt UI after Terraform creation |

---

##  Contributing

> 👋 **Hey Developer!**
>
> We’re glad you’re here and excited that you're interested in contributing. Right now, we’re in the middle of setting up some core processes — like contribution guidelines, issue templates, and workflows — to make sure everything runs smoothly for everyone.
>
> While we’re not quite ready for external contributions just yet, we’re getting close! Hang tight, and keep an eye on this space — we’ll be opening things up soon, and we’d love to have you onboard when we do.

---

##  License

Licensed under Mozilla Public License 2.0. Refer to [LICENSE](LICENSE) for full license details.

---