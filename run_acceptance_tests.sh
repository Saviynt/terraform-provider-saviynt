#!/bin/bash
# Copyright (c) HashiCorp, Inc.
# SPDX-License-Identifier: MPL-2.0


# Set required environment variables for Saviynt provider
# Replace these placeholder values with your actual Saviynt credentials
export TF_ACC=1
export SAVIYNT_URL="https://your-saviynt-instance.saviyntcloud.com"
export SAVIYNT_USERNAME="your-username"
export SAVIYNT_PASSWORD="your-password"

# Run all acceptance tests
# The -timeout flag is set to 120 minutes to allow for long-running tests
go test ./internal/provider -v -timeout 120m
