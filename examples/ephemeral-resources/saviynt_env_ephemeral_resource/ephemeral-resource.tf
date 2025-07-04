/*
 * Copyright (c) 2025 Saviynt Inc.
 * All Rights Reserved.
 *
 * This software is the confidential and proprietary information of
 * Saviynt Inc. ("Confidential Information"). You shall not disclose,
 * use, or distribute such Confidential Information except in accordance
 * with the terms of the license agreement you entered into with Saviynt.
 *
 * SAVIYNT MAKES NO REPRESENTATIONS OR WARRANTIES ABOUT THE SUITABILITY OF
 * THE SOFTWARE, EITHER EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO
 * THE IMPLIED WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR
 * PURPOSE, OR NON-INFRINGEMENT.
 */

ephemeral "saviynt_env_ephemeral_resource" "example" {}

# Can be used as below:
resource "saviynt_ad_connection_resource" "example" {
  connection_type = "AD"
  connection_name = "Terraform_AD_Connector"
  url             = format("%s://%s:%d", var.LDAP_PROTOCOL, var.IP_ADDRESS, var.LDAP_PORT)
  password        = ephemeral.saviynt_env_ephemeral_resource.example.svnt_password
  username        = ephemeral.saviynt_env_ephemeral_resource.example.svnt_username
}