schema_version = 1

project {
  # SPDX license identifier
  license = "MPL-2.0"

  # Custom copyright holder
  copyright_holder = "Saviynt Inc."

  # Year to appear in the header
  copyright_year = 2025

  # Ignore files and folders where license headers should not be added
  header_ignore = [
    "examples/**",
    ".github/ISSUE_TEMPLATE/*.yml",
    ".golangci.yml",
    ".goreleaser.yml",
  ]
}
