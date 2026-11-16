// Copyright (c) 2025 Saviynt Inc.
// SPDX-License-Identifier: MPL-2.0

resource "saviynt_application_data_import_job_resource" "example" {
  jobs = [
    {
      trigger_name        = "app_data_import_trigger_1" # required
      job_group           = "DATA"                      # required
      cron_expression     = "0 0 2 * * ?"               # required
      trigger_group       = "GRAILS_JOBS"               # optional
      security_system     = "sample-101"                # optional
      accounts_or_access  = "access"                    # optional
      external_conn       = "4750"                      # optional
      full_or_incremental = "full"                      # optional
      import_tables       = "USR02"                     # optional, SAP example
    },
    {
      trigger_name        = "app_data_import_trigger_2" # required
      job_group           = "DATA"                      # required
      cron_expression     = "0 0 3 * * ?"               # required
      trigger_group       = "GRAILS_JOBS"               # optional
      security_system     = "sample-101"                # optional
      accounts_or_access  = "access"                    # optional
      external_conn       = "4750"                      # optional
      full_or_incremental = "full"                      # optional
    },
    {
      # Single-threaded import: routes to the backend `SapImportJob` class
      # (equivalent to picking "ApplicationDataImportJob (Single Threaded)" in the UI).
      trigger_name        = "app_data_import_trigger_single_threaded"
      job_group           = "DATA"
      cron_expression     = "0 0 2 * * ?"
      trigger_group       = "GRAILS_JOBS"
      security_system     = "sample-101"
      accounts_or_access  = "accounts"
      external_conn       = "4750"
      full_or_incremental = "full"
      single_threaded     = true
    }
  ]
}
