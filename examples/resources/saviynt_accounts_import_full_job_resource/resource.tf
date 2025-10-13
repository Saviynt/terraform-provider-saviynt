// Copyright (c) 2025 Saviynt Inc.
// SPDX-License-Identifier: MPL-2.0

resource "saviynt_accounts_import_full_job_resource" "example" {
  jobs = [
    {
      trigger_name    = "accounts_import_trigger_1" # required
      job_name        = "AccountsImportFullJob"     # required
      job_group       = "DATABASE"                  # required
      cron_expression = "0 0 2 * * ?"               # required
      trigger_group   = "GRAILS_JOBS"               # optional
      connection_name = "AD_Connection_1"           # required
    },
    {
      trigger_name    = "accounts_import_trigger_2" # required
      job_name        = "AccountsImportFullJob"     # required
      job_group       = "DATABASE"                  # required
      cron_expression = "0 0 3 * * ?"               # required
      trigger_group   = "GRAILS_JOBS"               # optional
      connection_name = "REST_Connection_1"         # required
    }
  ]
}
