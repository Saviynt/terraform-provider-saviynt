// Copyright (c) 2025 Saviynt Inc.
// SPDX-License-Identifier: MPL-2.0

package jobcontrolutil

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// JobControlErrorResponse represents the structure of job control API error responses
type JobControlErrorResponse struct {
	Msg       string                   `json:"msg"`
	ErrorCode string                   `json:"errorCode"`
	Triggers  []JobControlTriggerError `json:"triggers,omitempty"`
}

// JobControlTriggerError represents individual trigger errors in the response
type JobControlTriggerError struct {
	Msg            string                 `json:"msg"`
	ValueMap       map[string]interface{} `json:"valueMap,omitempty"`
	CronExpression string                 `json:"cronexpression,omitempty"`
	TriggerName    string                 `json:"triggername,omitempty"`
	JobName        string                 `json:"jobname,omitempty"`
	JobGroup       string                 `json:"jobgroup,omitempty"`
}

// JobControlHandleHTTPError handles HTTP errors specifically for job control operations
// It handles two types of 412 error responses:
// 1. Detailed trigger errors (with triggers array)
// 2. Simple error messages (msg + errorCode only)
func JobControlHandleHTTPError(ctx context.Context, httpResp *http.Response, err error, operation string, diagnostics *diag.Diagnostics) bool {
	if err != nil {
		if httpResp != nil && httpResp.StatusCode == http.StatusPreconditionFailed {
			tflog.Error(ctx, fmt.Sprintf("HTTP 412 error during %s", operation), map[string]interface{}{
				"status":      httpResp.Status,
				"status_code": httpResp.StatusCode,
			})

			var errorResp JobControlErrorResponse
			if decodeErr := json.NewDecoder(httpResp.Body).Decode(&errorResp); decodeErr != nil {
				diagnostics.AddError("Failed to decode error response", decodeErr.Error())
				return true
			}

			// Handle detailed trigger errors (Type 1: with triggers array)
			if len(errorResp.Triggers) > 0 {
				for _, trigger := range errorResp.Triggers {
					diagnostics.AddError(
						"Job Control Trigger Error",
						fmt.Sprintf("Trigger '%s': %s", trigger.TriggerName, trigger.Msg),
					)
				}
				return true
			}

			// Handle simple error messages (Type 2: msg + errorCode only)
			if errorResp.Msg != "" {
				diagnostics.AddError(
					"Job Control Error",
					fmt.Sprintf("Error during %s: %s", operation, errorResp.Msg),
				)
				return true
			}

			// Fallback for unknown 412 format
			diagnostics.AddError(
				"HTTP Error",
				fmt.Sprintf("HTTP 412 error during %s: %s", operation, httpResp.Status),
			)
			return true

		} else if httpResp != nil && httpResp.StatusCode != http.StatusOK {
			tflog.Error(ctx, fmt.Sprintf("HTTP error during %s", operation), map[string]interface{}{
				"status":      httpResp.Status,
				"status_code": httpResp.StatusCode,
				"error":       err.Error(),
			})
			diagnostics.AddError(
				"HTTP Error",
				fmt.Sprintf("HTTP %d error during %s: %s - %v", httpResp.StatusCode, operation, httpResp.Status, err),
			)
			return true
		} else {
			diagnostics.AddError(
				"Request Error",
				fmt.Sprintf("Error during %s: %v", operation, err),
			)
			return true
		}
	}
	return false
}

// JobControlHandleAPIError handles API error responses from job control operations
func JobControlHandleAPIError(ctx context.Context, errorCode *string, message *string, operation string, diagnostics *diag.Diagnostics) bool {
	if errorCode != nil && *errorCode != "0" {
		if message != nil && *message != "" {
			tflog.Error(ctx, fmt.Sprintf("Job Control API error during %s", operation), map[string]interface{}{
				"error_code": *errorCode,
				"message":    *message,
			})
			diagnostics.AddError(
				"Job Control API Error",
				fmt.Sprintf("API error during %s: %s", operation, *message),
			)
			return true
		} else {
			tflog.Error(ctx, fmt.Sprintf("Job Control API error during %s", operation), map[string]interface{}{
				"error_code": *errorCode,
			})
			diagnostics.AddError(
				"Job Control API Error",
				fmt.Sprintf("API error during %s: Error Code %s", operation, *errorCode),
			)
			return true
		}
	}
	return false
}
