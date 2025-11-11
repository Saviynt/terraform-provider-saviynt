// Copyright (c) 2025 Saviynt Inc.
// SPDX-License-Identifier: MPL-2.0

resource "saviynt_job_control_resource" "example" {
  run_jobs = [
    {
      job_name        = "WSRetryJob"
      trigger_name    = "ws_retry_trigger_example"
      job_group       = "Utility"
      run_job_version = "v1.0"
    },
    {
      job_name        = "UserImportJob"
      trigger_name    = "user_import_trigger"
      job_group       = "Import"
      run_job_version = "v1.0"
    }
  ]
}
