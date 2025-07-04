// Copyright (c) Saviynt Inc.

data "saviynt_ad_connection_datasource" "ad" {
  connection_key = "123"
  # connection_name="sample"          #Either one can be used
}