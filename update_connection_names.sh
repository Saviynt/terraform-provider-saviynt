#!/bin/bash
# Copyright (c) 2025 Saviynt Inc.
# SPDX-License-Identifier: MPL-2.0


# Function to generate a random string
generate_random_name() {
  local prefix=$1
  local random_suffix=$(cat /dev/urandom | LC_ALL=C tr -dc 'a-zA-Z0-9' | fold -w 8 | head -n 1)
  echo "${prefix}_${random_suffix}"
}

# Function to update connection_name in JSON files
update_connection_name() {
  local file=$1
  local connector_type=$2
  
  # Generate random names for each section
  local create_name=$(generate_random_name "Terraform_${connector_type}")
  local update_name=$create_name
  local update_connection_name=$(generate_random_name "Terraform_${connector_type}")
  
  # Use jq to update the JSON file
  # First, create a temporary file with the updated content
  jq --arg create_name "$create_name" \
     --arg update_name "$update_name" \
     --arg update_connection_name "$update_connection_name" \
     '.create.connection_name = $create_name | 
      .update.connection_name = $update_name | 
      if has("update_connection_name") then .update_connection_name.connection_name = $update_connection_name else . end |
      if has("update_connection_type") then .update_connection_type.connection_name = $update_name else . end' \
     "$file" > "${file}.tmp"
  
  # Replace the original file with the updated one
  mv "${file}.tmp" "$file"
  
  echo "Updated $file with new connection names"
}

# Process each connection test file
update_connection_name "./internal/provider/ad_connection_test_data.json" "AD"
update_connection_name "./internal/provider/adsi_connection_test_data.json" "ADSI"
update_connection_name "./internal/provider/db_connection_test_data.json" "DB"
update_connection_name "./internal/provider/entra_id_connection_test_data.json" "EntraID"
update_connection_name "./internal/provider/github_rest_connection_test_data.json" "GithubREST"
update_connection_name "./internal/provider/rest_connection_test_data.json" "REST"
update_connection_name "./internal/provider/salesforce_connection_test_data.json" "Salesforce"
update_connection_name "./internal/provider/sap_connection_test_data.json" "SAP"
update_connection_name "./internal/provider/unix_connection_test_data.json" "Unix"
update_connection_name "./internal/provider/workday_connection_test_data.json" "Workday"

echo "All connection test files have been updated with random connection names."
