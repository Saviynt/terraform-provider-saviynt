// Copyright (c) 2025 Saviynt Inc.
// SPDX-License-Identifier: MPL-2.0

terraform {
  required_providers {
    saviynt = {
      source  = "saviynt/saviynt"
      version = "x.x.x"
    }
  }
}

# Option 1: OAuth2 Token Exchange (Entra ID / M2M) — highest priority
provider "saviynt" {
  server_url    = "https://example.saviyntcloud.com"
  subject_token = var.entra_access_token # Entra ID access token
  scope         = "terraformtesting"     # Saviynt ExternalConnection name
}

# Option 2: Direct Bearer Token — second priority
# provider "saviynt" {
#   server_url    = "https://example.saviyntcloud.com"
#   access_token  = var.saviynt_access_token
#   refresh_token = var.saviynt_refresh_token # optional: enables auto-refresh when access token expires
# }

# Option 3: Username + Password — fallback
# provider "saviynt" {
#   server_url = "https://example.saviyntcloud.com"
#   username   = var.saviynt_username
#   password   = var.saviynt_password
# }