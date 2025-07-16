// Copyright (c) Saviynt Inc.
// SPDX-License-Identifier: MPL-2.0

terraform {
  required_providers {
    saviynt = {
      source  = "saviynt/saviynt"
      version = "x.x.x"
    }
  }
}

provider "saviynt" {
  server_url = "https://example.saviyntcloud.com"
  username   = "username"
  password   = "password"
}