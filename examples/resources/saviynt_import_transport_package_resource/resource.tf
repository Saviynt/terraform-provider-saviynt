// Copyright (c) 2025 Saviynt Inc.
// SPDX-License-Identifier: MPL-2.0

resource "saviynt_import_transport_package_resource" "example" {
  package_path           = "/saviynt_shared/transport_packages/my_package.zip" # required - Path to the transport package file (zip file created by export transport)
  update_user            = "admin"                                             # optional - User performing the import
  business_justification = "Importing configuration changes for Q1 release"    # optional - Business reason for import
  import_package_version = "1.0"                                               # optional - Version identifier for the package
}

# Example with minimal configuration
resource "saviynt_import_transport_package_resource" "minimal" {
  package_path = "/saviynt_shared/transport_packages/minimal_package.zip" # zip file created by export transport
}
