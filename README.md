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

- Terraform version `>= 1.8+`
- Go programming language `>= 1.21+`
- Saviynt Identity Cloud instance and credentials.

---

##  Features

Following resources are available for management: 
- Security System
- Endpoint
- Connections
- Dynamic Attributes

Following connectors are available:
- Active Directory(AD)
- REST
- ADSI
- Database(DB)
- EntraID(AzureAD)
- Salesforce
- Workday
- Unix
- Github REST

Ephemeral resources available:
- [File ephemeral resource](#ephemeral-file-credential-resource)
- [Env ephemeral resource](#ephemeral-env-credential-resource)
---

##  Documentation

Check out the [Latest Saviynt Provider Docs](https://registry.terraform.io/providers/saviynt/saviynt/latest/docs) to know more.

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
## Ephemeral File Credential Resource

The **Ephemeral File Credential Resource** is a transient Terraform resource that provides temporary, in-memory credentials to other connector resources by reading values from a local JSON file at apply time. This allows secure and flexible provisioning without persisting sensitive data in the Terraform state.

### Supported Connectors

The following connectors are supported and can consume credentials provided by this resource:

- **AD**: `username`, `password`
- **ADSI**: `username`, `password`
- **DB**: `username`, `password`, `change_pass_json`
- **EntraId**: `client_id`, `client_secret`, `access_token`,`azure_mgmt_access_token`, `windows_connector_json`, `change_pass_json`, `connection_json`
- **Github REST**: `connection_json`, `access_tokens`
- **REST**: `connection_json`, `change_pass_json`
- **Salesforce**: `client_id`, `client_secret`, `refresh_token`
- **SAP**: `password`, `prov_password`
- **Unix**: `username`, `password`, `change_password_json`,  `passphrase`, `ssh_key`, `ssh_pass_through_password`, `ssh_pass_through_sshkey`, `ssh_pass_through_passphrase`
- **Workday**: `username`, `password`, `client_id`, `client_secret`, `refresh_token`

### Usage

The ephemeral credential resource reads from a local JSON file structured with the required fields for the target connector. These fields are then dynamically injected into the respective connector resources during the `apply` phase.

> **Note:** This resource is ephemeral and does not store any state. It is designed for use cases where credentials must remain local and transient.

## Ephemeral Env Credential Resource

The **Ephemeral Env Credential Resource** is a transient Terraform resource that provides temporary, in-memory credentials to other connector resources by reading values from the environment variables at apply time. This allows secure and flexible provisioning without persisting sensitive data in the Terraform state.

### Supported Connectors

The following connectors are supported and can consume credentials provided by this resource:

- **AD**: `username`, `password`
- **ADSI**: `username`, `password`
- **DB**: `username`, `password`, `change_pass_json`
- **EntraId**: `client_id`, `client_secret`, `access_token`,`azure_mgmt_access_token`, `windows_connector_json`, `change_pass_json`, `connection_json`
- **Github REST**: `connection_json`, `access_tokens`
- **REST**: `connection_json`, `change_pass_json`
- **Salesforce**: `client_id`, `client_secret`, `refresh_token`
- **SAP**: `password`, `prov_password`
- **Unix**: `username`, `password`, `change_password_json`,  `passphrase`, `ssh_key`, `ssh_pass_through_password`, `ssh_pass_through_sshkey`, `ssh_pass_through_passphrase`
- **Workday**: `username`, `password`, `client_id`, `client_secret`, `refresh_token`

### Usage

The ephemeral credential resource reads from a local JSON file structured with the required fields for the target connector. These fields are then dynamically injected into the respective connector resources during the `apply` phase.

> **Note:** This resource is ephemeral and does not store any state. It is designed for use cases where credentials must remain local and transient.
## Security Considerations

- Ensure the credential file is secured and not committed to version control.
- Avoid using this resource in long-lived plans, as it relies on local files that may change or expire.

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

## Known Limitations

### Note: This provider is for Saviynt ECM v24.4 only.

The following limitations are present in the latest version of the provider. These are being prioritized for resolution in the upcoming release alongside new feature additions:

### 1. All Resource objects
 - `terraform destroy` is not supported for all resources except for dynamic attributes.

### 2. Endpoints

- **State management is not supported** for the following attributes:
  - `Owner`
  - `ResourceOwner`

- For`saviynt_endpoint_resource.requestable_role_types.request_option`, the supported values for proper state tracking are:
  - `DropdownSingle`
  - `Table`
  - `TableOnlyAdd`

- The following settings are **not currently configurable via Terraform**:
  - `mapped_endpoints`
  - `show_on` attribute in 'requestable_role_types`
  - `Disable Remove Service Account`
  - `Disable Modify Service Account`
  - `Disable New Account Request if Account Exists`

### 3. Connections
- `description` field can't be set from Terraform currently.
- **State management** is not supported for the following attributes due to their sensitive nature:
  - **AD**: `password`
  - **ADSI**: `password`
  - **DB**: `password`, `change_pass_json`
  - **EntraId**: `access_token`, `azure_mgmt_access_token`, `client_secret`, `windows_connector_json`, `connection_json`
  - **Github REST**: `connection_json`, `access_tokens`
  - **REST**: `connection_json`
  - **Salesforce**: `client_secret`, `refresh_token`
  - **Unix**: `password`, `passphrase`, `ssh_key`, `ssh_pass_through_password`, `ssh_pass_through_sshkey`, `ssh_pass_through_passphrase`
  - **Workday**: `password`, `client_secret`, `refresh_token`
- The following fields are **not currently configurable via Terraform**:
  - **Github REST**: `Status_Threshold_Config`, `Pam_Config`
  - **Workday**: `orgrole_import_payload`

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
