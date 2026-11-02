// Copyright (c) 2025 Saviynt Inc.
// SPDX-License-Identifier: MPL-2.0

# Import an entitlement using the format: endpoint_name:entitlement_type:entitlement_value
#
# Note: endpoint_name and entitlement_type must not contain colons.
#       Only entitlement_value supports colons (e.g. AWS ARNs, Azure resource IDs).
#
# Plain value:
#   terraform import saviynt_entitlement_resource.test_entitlement sample-endpoint:sample-ent-type:sample-ent
#
# Value containing colons (e.g. an AWS ARN):
#   terraform import saviynt_entitlement_resource.test_entitlement sample-endpoint:sample-ent-type:arn:aws:iam::123456789012:role/MyRole

resource "saviynt_entitlement_resource" "test_entitlement" {
  # Required attributes
  endpoint          = "sample-endpoint"
  entitlement_type  = "sample-ent-type"
  entitlement_value = "sample-ent"

  # Optional attributes
  displayname          = "sample-ent-display-name"
  description          = "Test entitlement created via Terraform for comprehensive testing"
  entitlement_glossary = "Test entitlement for access management"

  risk            = 3
  status          = 1
  soxcritical     = 3
  syscritical     = 1
  privileged      = 2
  confidentiality = 2
  priority        = 10

  module = "IT"
  access = "Update"

  entitlement_owners = {
    rank_1 = ["user1"]
    rank_2 = ["user1", "user2"]
  }

  entitlement_map = [
    {
      entitlement_value         = "sample_ent_1"
      entitlement_type          = "sample_ent_type_2"
      endpoint                  = "sample_endpoint2"
      request_filter            = false
      exclude_entitlement       = true
      add_dependent_task        = true
      remove_dependent_ent_task = false
    },
    {
      entitlement_value         = "sample_ent_2"
      entitlement_type          = "sample_ent_type_3"
      endpoint                  = "sample_endpoint3"
      request_filter            = false
      exclude_entitlement       = true
      add_dependent_task        = true
      remove_dependent_ent_task = false
    }
  ]

  customproperty1  = "Test Custom Property 1"
  customproperty2  = "Test Custom Property 2"
  customproperty4  = "Test Custom Property 4"
  customproperty21 = "Test custom property 21"
}