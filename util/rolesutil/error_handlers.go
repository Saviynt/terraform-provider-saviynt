// Copyright (c) 2025 Saviynt Inc.
// SPDX-License-Identifier: MPL-2.0

package rolesutil

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// RoleHandleHTTPError handles HTTP errors and decodes error responses
// This is a shared utility function used by both roles resource and datasource
func RoleHandleHTTPError(ctx context.Context, httpResp *http.Response, err error, operation string, diagnostics *diag.Diagnostics) bool {
	if err != nil {
		if httpResp != nil && httpResp.StatusCode == http.StatusPreconditionFailed {
			tflog.Error(ctx, fmt.Sprintf("HTTP error during %s", operation), map[string]interface{}{
				"status":      httpResp.Status,
				"status_code": httpResp.StatusCode,
			})
			var fetchResp map[string]interface{}
			if decodeErr := json.NewDecoder(httpResp.Body).Decode(&fetchResp); decodeErr != nil {
				diagnostics.AddError("Failed to decode error response", decodeErr.Error())
				return true
			}
			// Safe message extraction
			message := "Unknown error"
			if msg, exists := fetchResp["message"]; exists && msg != nil {
				message = fmt.Sprintf("%v", msg)
			}

			diagnostics.AddError(
				"HTTP Error",
				fmt.Sprintf("HTTP error while %s: %s", operation, message),
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
				fmt.Sprintf("HTTP %d error while %s: %s - %v", httpResp.StatusCode, operation, httpResp.Status, err),
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

// RoleHandleAPIError handles API error responses from Saviynt
// This is a shared utility function used by both roles resource and datasource
func RoleHandleAPIError(ctx context.Context, errorCode *string, message *string, operation string, diagnostics *diag.Diagnostics) bool {
	if errorCode != nil && *errorCode != "0" {
		if message != nil && *message != "" {
			tflog.Error(ctx, fmt.Sprintf("API error during %s", operation), map[string]interface{}{
				"error_code": *errorCode,
				"message":    *message,
			})
			diagnostics.AddError(
				"API Error",
				fmt.Sprintf("API error while %s: %s", operation, *message),
			)
			return true
		} else {
			tflog.Error(ctx, fmt.Sprintf("API error during %s", operation), map[string]interface{}{
				"error_code": *errorCode,
			})
			diagnostics.AddError(
				"API Error",
				fmt.Sprintf("API error while %s: %s", operation, *errorCode),
			)
			return true
		}
	}
	return false
}
