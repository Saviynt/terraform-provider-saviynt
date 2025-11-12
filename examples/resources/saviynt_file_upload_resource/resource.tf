// Copyright (c) 2025 Saviynt Inc.
// SPDX-License-Identifier: MPL-2.0

resource "saviynt_file_upload_resource" "f" {
  file_path     = "/path/to/users.csv"
  path_location = "Datafiles"

  #Optional
  file_version = "v1.1"
}