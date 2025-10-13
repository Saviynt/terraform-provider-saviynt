// Copyright (c) 2025 Saviynt Inc.
// SPDX-License-Identifier: MPL-2.0

resource "saviynt_accounts_import_incremental_job_resource" "example" {
  jobs = [
    {
      name            = "accounts_incremental_trigger_1" # required
      job_name        = "AccountsImportIncrementalJob"   # required
      job_group       = "DATABASE"                       # required
      group           = "GRAILS_JOBS"                    # required
      cron_exp        = "0 0 2 * * ?"                    # required
      connection_name = "AD_Connection_1"                # optional
    },
    {
      name            = "accounts_incremental_trigger_2" # required
      job_name        = "AccountsImportIncrementalJob"   # required
      job_group       = "DATABASE"                       # required
      group           = "GRAILS_JOBS"                    # required
      cron_exp        = "0 0 3 * * ?"                    # required
      connection_name = "REST_Connection_1"              # optional
    }
  ]
}
