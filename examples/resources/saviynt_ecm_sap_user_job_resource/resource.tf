// Copyright (c) 2025 Saviynt Inc.
// SPDX-License-Identifier: MPL-2.0

resource "saviynt_ecm_sap_user_job_resource" "example" {
  jobs = [
    {
      trigger_name    = "ecm_sap_user_trigger_1" # required
      job_group       = "ecmGroup"               # required
      cron_expression = "0 0 2 * * ?"            # required
      trigger_group   = "GRAILS_JOBS"            # optional
      on_failure      = "continue"               # optional
    },
    {
      trigger_name    = "ecm_sap_user_trigger_2" # required
      job_group       = "ecmGroup"               # required
      cron_expression = "0 0 3 * * ?"            # required
      trigger_group   = "GRAILS_JOBS"            # optional
      on_failure      = "continue"               # optional
    }
  ]
}
