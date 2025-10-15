// Copyright (c) 2025 Saviynt Inc.
// SPDX-License-Identifier: MPL-2.0

resource "saviynt_ws_retry_job_resource" "example" {
  jobs = [
    {
      trigger_name     = "ws_retry_trigger_1" # required
      job_group        = "utility"            # required
      cron_expression  = "0 0 2 * * ?"        # required
      trigger_group    = "GRAILS_JOBS"        # optional
      security_systems = ["sample-101"]       # optional
      task_types       = "1,2,3"              # optional
    },
    {
      trigger_name     = "ws_retry_trigger_2" # required
      job_group        = "utility"            # required
      cron_expression  = "0 0 3 * * ?"        # required
      trigger_group    = "GRAILS_JOBS"        # optional
      security_systems = ["sample-101"]       # optional
      task_types       = "1,2,3"              # optional
    }
  ]
}
