// Copyright (c) 2025 Saviynt Inc.
// SPDX-License-Identifier: MPL-2.0

resource "saviynt_user_import_job_resource" "example" {
  jobs = [
    {
      trigger_name                           = "user_import_trigger_1" # required
      job_name                               = "UserImportJob"         # required
      job_group                              = "DATA"                  # required
      cron_expression                        = "0 0 2 * * ?"           # required
      trigger_group                          = "GRAILS_JOBS"           # optional
      external_conn                          = "4750"                  # required
      full_or_incremental                    = "full"                  # optional
      user_not_in_feed_action                = "INACTIVATE"            # optional
      user_operations_allowed                = "CREATE,UPDATE"         # optional
      zero_day_provisioning                  = "true"                  # optional
      generate_system_username               = "true"                  # optional
      generate_email                         = "true"                  # optional
      check_rules                            = "true"                  # optional
      build_user_map                         = "true"                  # optional
      user_threshold                         = "1000"                  # optional
      on_failure                             = "CONTINUE"              # optional
      zero_day_limit                         = "100"                   # optional
      term_user_limit                        = "50"                    # optional
      import_sav_connect                     = "false"                 # optional
      export_to_sav_cloud                    = "false"                 # optional
      user_reconciliation_field              = "firstname"             # optional
      user_default_sav_role                  = "ROLE_ADMIN"            # optional
      user_status_config                     = "active"                # optional
      endpoints_to_associate_orphan_accounts = "sample-101,sample-102" # optional
    },
    {
      trigger_name                           = "user_import_trigger_2" # required
      job_name                               = "UserImportJob"         # required
      job_group                              = "DEFAULT"               # required
      cron_expression                        = "0 0 3 * * ?"           # required
      trigger_group                          = "DEFAULT"               # optional
      external_conn                          = "4750"                  # required
      full_or_incremental                    = "full"                  # optional
      user_not_in_feed_action                = "ACTIVATE"              # optional
      user_operations_allowed                = "CREATE"                # optional
      zero_day_provisioning                  = "true"                  # optional
      generate_system_username               = "true"                  # optional
      generate_email                         = "true"                  # optional
      check_rules                            = "true"                  # optional
      build_user_map                         = "true"                  # optional
      user_threshold                         = "1000"                  # optional
      on_failure                             = "CONTINUE"              # optional
      zero_day_limit                         = "100"                   # optional
      term_user_limit                        = "50"                    # optional
      import_sav_connect                     = "false"                 # optional
      export_to_sav_cloud                    = "false"                 # optional
      user_reconciliation_field              = "firstname"             # optional
      user_default_sav_role                  = "ROLE_ADMIN"            # optional
      user_status_config                     = "active"                # optional
      endpoints_to_associate_orphan_accounts = "sample-101,sample-102" # optional
    }
  ]
}
