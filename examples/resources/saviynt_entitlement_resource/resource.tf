// Copyright (c) 2025 Saviynt Inc.
// SPDX-License-Identifier: MPL-2.0

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
      entitlement_value = "sample_ent_1"
      entitlement_type  = "sample_ent_type_2"
      endpoint          = "sample_endpoint2"
    },
    {
      entitlement_value = "sample_ent_2"
      entitlement_type  = "sample_ent_type_3"
      endpoint          = "sample_endpoint3"
    }
  ]

  customproperty1  = "Test Custom Property 1"
  customproperty2  = "Test Custom Property 2"
  customproperty4  = "Test Custom Property 4"
  customproperty21 = "Test custom property 21"
}