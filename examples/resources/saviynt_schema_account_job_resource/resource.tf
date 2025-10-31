// Copyright (c) 2025 Saviynt Inc.
// SPDX-License-Identifier: MPL-2.0

resource "saviynt_schema_account_job_resource" "example" {
  jobs = [
    {
      name              = "schema_account_trigger_1"           # required
      job_group         = "Schema"                             # required
      group             = "GRAILS_JOBS"                        # required
      cron_exp          = "0 0 2 * * ?"                        # required
      schema_file_names = "ConnScrumACCOUNTS.sav,ACCOUNTS.sav" # optional
    },
    {
      name              = "schema_account_trigger_2"     # required
      job_group         = "Schema"                       # required
      group             = "GRAILS_JOBS"                  # required
      cron_exp          = "0 0 3 * * ?"                  # required
      schema_file_names = "Accounts.sav,ep_accounts.sav" # optional
    }
  ]
}
