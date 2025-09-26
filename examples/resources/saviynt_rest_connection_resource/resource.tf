// Copyright (c) 2025 Saviynt Inc.
// SPDX-License-Identifier: MPL-2.0

resource "saviynt_rest_connection_resource" "example" {
  connection_name = "Terraform_Rest_Connector"

  # Using file() function to read JSON content from external files
  connection_json         = file("${path.module}/json/connection.json")
  import_user_json        = file("${path.module}/json/import_user.json")
  import_account_ent_json = file("${path.module}/json/import_account_ent.json")
  status_threshold_config = file("${path.module}/json/status_threshold_config.json")
  create_account_json     = file("${path.module}/json/create_account.json")

  # Other JSON configurations are still inline with jsonencode
  # These could be extracted to files in a similar way if needed
  update_account_json = "{\"dateFormat\":\"yyyy-MM-dd'T'HH:mm:ssXXX\",\"responseColsToPropsMap\":{\"displayName\":\"call1.message.user.name~#~char\"},\"call\":[{\"name\":\"Role\",\"connection\":\"acctAuth\",\"url\":\"@HOSTNAME@/api/v2/users/${account.accountID}\",\"httpMethod\":\"PUT\",\"httpParams\":\"{\\\"user\\\":{\\\"name\\\":\\\"${user.firstname} ${user.lastname}\\\"}}\",\"httpHeaders\":{\"Authorization\":\"${access_token}\",\"Accept\":\"application/json\"},\"httpContentType\":\"application/json\",\"successResponses\":{\"statusCode\":[200,201]}}]}"

  enable_account_json = "{\"call\":[{\"name\":\"call1\",\"connection\":\"acctAuth\",\"url\":\"@HOSTNAME@/api/v2/users\",\"httpMethod\":\"PUT\",\"httpParams\":\"{\\\"user\\\":{\\\"suspended\\\":\\\"false\\\"}}\",\"httpHeaders\":{\"Authorization\":\"${access_token}\",\"Accept\":\"application/json\"},\"httpContentType\":\"application/json\",\"successResponses\":{\"statusCode\":[200,201]}}]}"

  disable_account_json = "{\"call\":[{\"name\":\"call1\",\"connection\":\"acctAuth\",\"url\":\"@HOSTNAME@/api/v2/users\",\"httpMethod\":\"PUT\",\"httpParams\":\"{\\\"user\\\":{\\\"suspended\\\":\\\"true\\\"}}\",\"httpHeaders\":{\"Authorization\":\"${access_token}\",\"Accept\":\"application/json\"},\"httpContentType\":\"application/json\",\"successResponses\":{\"statusCode\":[200,201]}}]}"
}
