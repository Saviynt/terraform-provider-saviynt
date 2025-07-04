// Copyright (c) 2025 Saviynt Inc.
// SPDX-License-Identifier: MPL-2.0

variable "HOST_NAME" {
  type        = string
  description = "The hostname of the Unix server"
}
variable "PORT_NUMBER" {
  type        = string
  description = "The port number for the Unix server"
}
variable "USERNAME" {
  type        = string
  description = "The username for the Unix server"
}
variable "PASSWORD" {
  type        = string
  description = "The password for the Unix server"
  sensitive   = true
}
variable "PEM_KEY_FILE" {
  type        = string
  description = "The path to the PEM key file"
}
variable "PASSPHRASE" {
  type        = string
  description = "The value used in conjunction with the PEM Key file."
}
variable "SSH_KEY" {
  type        = string
  description = "SSH key used to connect to Unix serv"
}

resource "saviynt_unix_connection_resource" "example" {
  connection_type = "Unix"
  connection_name = "Terraform_Unix_Connector"
  host_name       = var.HOST_NAME
  port_number     = var.PORT_NUMBER
  username        = var.USERNAME
  password        = var.PASSWORD

  groups_file   = "sudo cat /etc/group"
  accounts_file = jsonencode({ unix = "sudo cat /etc/passwd" })
  shadow_file   = jsonencode({ unix = "sudo cat /etc/shadow" })

  provision_account_command = "sudo useradd $${username} -p $${password} -c \"$${user?.country}/$${user?.employeeType}/$${user?.employeeid}/$${user?.lastname}.$${user?.firstname}/$${user?.email}\" -g users"

  deprovision_account_command = "sudo userdel $${accountName}"

  change_password_json = jsonencode({
    command        = "echo '$${account.name}:$${newPassword}' | sudo /usr/sbin/chpasswd"
    changePassword = true
    changeKey      = true
  })


  # pem_key_file = var.PEM_KEY_FILE         ## This is required for passwordless authetication

  enable_account_command = "sudo usermod -U $${accountName}"

  disable_account_command = "sudo usermod -L $${accountName}"

  account_entitlement_mapping_command = "sudo cat /etc/group"

  # passphrase = var.PASSPHRASE            ## This is required for passwordless authetication

  update_account_command = "sudo usermod -c \"$${user?.country}/$${user?.employeeType}/$${user?.employeeid}/$${user?.lastname}.$${user?.firstname}/$${user?.email}\" $${username}"

  create_group_command = "sudo groupadd $${groupName}"

  delete_group_command = "sudo groupdel  $${groupName}"

  add_group_owner_command = "sudo gpasswd -A $${username} $${groupName}"

  add_primary_group_command = "sudo gpasswd -A $${username} $${groupName}"

  fire_fighter_id_grant_access_command = "sudo usermod -c \"$${user?.country}/$${user?.lastname}.$${user?.firstname}/$${user?.email}\" $${username}"

  fire_fighter_id_revoke_access_command = "sudo usermod -c \"$${user?.country}/$${user?.firstname}\" $${username}"

  inactive_lock_account = "true"

  status_threshold_config = jsonencode({
    statusAndThresholdConfig = {
      accountThresholdValue         = 25
      accountEntThresholdValue      = 3
      entThresholdValue             = 10
      deleteAccEntForActiveAccounts = true
      statusColumn                  = "customproperty1"
      deleteLinks                   = false
      correlateInactiveAccounts     = true
      inactivateAccountsNotInFile   = false
      inactivateEntsNotInFeed       = true
      activeStatus                  = ["1", "Active"]
      lockedStatusColumn            = "userlock"
      lockedStatusMapping = {
        Locked   = ["1"]
        Unlocked = ["0"]
      }
    }
  })


  custom_config_json = jsonencode({
    connectionTimeout = 10
    readTimeout       = 50
    writeTimeout      = 50
    retryWait         = 2
    retryCount        = 3
  })

  ssh_key = <<EOKEY
  -----BEGIN RSA PRIVATE KEY-----
  .....
  -----END RSA PRIVATE KEY-----
  EOKEY

  lock_account_command   = "sudo usermod -L $${accountName}"
  unlock_account_command = "sudo usermod -U $${accountName}"
}