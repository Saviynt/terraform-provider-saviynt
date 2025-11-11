// Copyright (c) 2025 Saviynt Inc.
// SPDX-License-Identifier: MPL-2.0

resource "saviynt_file_transfer_job_resource" "example" {
  jobs = [
    {
      name                    = "file_transfer_trigger_1" # required
      job_group               = "ECFConnector"            # required
      group                   = "GRAILS_JOBS"             # required
      cron_exp                = "0 0 2 * * ?"             # required
      external_connection_key = "7442"                    # required
      file_transfer_action    = "UPLOAD"                  # required
    },
    {
      name                    = "file_transfer_trigger_2" # required
      job_group               = "ECFConnector"            # required
      group                   = "GRAILS_JOBS"             # required
      cron_exp                = "0 0 3 * * ?"             # required
      external_connection_key = "5781"                    # required
      file_transfer_action    = "DOWNLOAD"                # required
    }
  ]
}
