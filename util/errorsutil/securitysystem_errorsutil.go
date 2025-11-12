// Copyright (c) 2025 Saviynt Inc.
// SPDX-License-Identifier: MPL-2.0

package errorsutil

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// SecuritySystemErrorCategory represents different categories of errors
type SecuritySystemErrorCategory string

const (
	SecuritySystemCategoryConfiguration   SecuritySystemErrorCategory = "CONFIG"
	SecuritySystemCategoryBusinessLogic   SecuritySystemErrorCategory = "BUSINESS"
	SecuritySystemCategoryAPIOperation    SecuritySystemErrorCategory = "API"
	SecuritySystemCategoryStateManagement SecuritySystemErrorCategory = "STATE"
)

// SecuritySystemStandardError represents a standardized error with context
type SecuritySystemStandardError struct {
	Code           string
	Message        string
	Operation      string
	SecuritySystem string
	Category       SecuritySystemErrorCategory
	OriginalError  error
}

// Error implements the error interface
func (se *SecuritySystemStandardError) Error() string {
	if se.OriginalError != nil {
		return fmt.Sprintf("[%s] %s during %s operation for security system '%s': %v",
			se.Code, se.Message, se.Operation, se.SecuritySystem, se.OriginalError)
	}
	return fmt.Sprintf("[%s] %s during %s operation for security system '%s'",
		se.Code, se.Message, se.Operation, se.SecuritySystem)
}

// Unwrap returns the original error for error unwrapping
func (se *SecuritySystemStandardError) Unwrap() error {
	return se.OriginalError
}

// SecuritySystemErrorCodeGenerator generates standardized error codes
type SecuritySystemErrorCodeGenerator struct{}

// NewSecuritySystemErrorCodeGenerator creates a new error code generator for a specific connector type
func NewSecuritySystemErrorCodeGenerator() *SecuritySystemErrorCodeGenerator {
	return &SecuritySystemErrorCodeGenerator{}
}

// GenerateErrorCode generates a standardized error code
func (ecg *SecuritySystemErrorCodeGenerator) GenerateErrorCode(category SecuritySystemErrorCategory, sequence int) string {
	var categoryCode string
	switch category {
	case SecuritySystemCategoryConfiguration:
		categoryCode = "00"
	case SecuritySystemCategoryBusinessLogic:
		categoryCode = "10"
	case SecuritySystemCategoryAPIOperation:
		categoryCode = "20"
	case SecuritySystemCategoryStateManagement:
		categoryCode = "30"
	default:
		categoryCode = "99"
	}

	return fmt.Sprintf("SECURITY_SYSTEM_%s%d", categoryCode, sequence)
}

// Common error codes that can be used across all connectors
const (
	// Configuration errors (001-009)
	ErrSecuritySystemProviderConfig    = "PROVIDER_CONFIG"
	ErrSecuritySystemPlanExtraction    = "PLAN_EXTRACTION"
	ErrSecuritySystemConfigExtraction  = "CONFIG_EXTRACTION"
	ErrSecuritySystemStateExtraction   = "STATE_EXTRACTION"
	ErrSecuritySystemMissingIdentifier = "MISSING_IDENTIFIER"

	// Business logic errors (101-109)
	ErrSecuritySystemDuplicateName = "DUPLICATE_NAME"
	ErrSecuritySystemNameImmutable = "NAME_IMMUTABLE"
	ErrSecuritySystemInvalidConfig = "INVALID_CONFIG"

	// API operation errors (201-209)
	ErrSecuritySystemCreateFailed = "CREATE_FAILED"
	ErrSecuritySystemReadFailed   = "READ_FAILED"
	ErrSecuritySystemUpdateFailed = "UPDATE_FAILED"
	ErrSecuritySystemDeleteFailed = "DELETE_FAILED"
	ErrSecuritySystemAPIError     = "API_ERROR"

	// State management errors (301-309)
	ErrSecuritySystemStateUpdate = "STATE_UPDATE"
	ErrSecuritySystemStateRead   = "STATE_READ"
)

// Common error messages
var securitySystemCommonErrorMessages = map[string]string{
	ErrSecuritySystemProviderConfig:    "Failed to configure provider",
	ErrSecuritySystemPlanExtraction:    "Failed to extract plan from Terraform request",
	ErrSecuritySystemConfigExtraction:  "Failed to extract configuration from Terraform request",
	ErrSecuritySystemStateExtraction:   "Failed to extract state from Terraform request",
	ErrSecuritySystemMissingIdentifier: "Missing required identifier for operation",
	ErrSecuritySystemDuplicateName:     "Security System with this name already exists",
	ErrSecuritySystemNameImmutable:     "Security System name cannot be modified after creation",
	ErrSecuritySystemInvalidConfig:     "Invalid configuration provided",
	ErrSecuritySystemCreateFailed:      "Failed to create Security System",
	ErrSecuritySystemReadFailed:        "Failed to read Security System details",
	ErrSecuritySystemUpdateFailed:      "Failed to update Security System",
	ErrSecuritySystemDeleteFailed:      "Failed to delete Security System",
	ErrSecuritySystemAPIError:          "API operation returned an error",
	ErrSecuritySystemStateUpdate:       "Failed to update Terraform state",
	ErrSecuritySystemStateRead:         "Failed to read Terraform state",
}

// GetSecuritySystemErrorMessage returns a standardized error message for the given error code
func GetSecuritySystemErrorMessage(errorCode string) string {
	if msg, exists := securitySystemCommonErrorMessages[errorCode]; exists {
		return msg
	}
	return "Unknown error occurred"
}

// CreateSecuritySystemStandardError creates a standardized error with code and context
func CreateSecuritySystemStandardError(errorCode, operation, securitySystem string, originalErr error) *SecuritySystemStandardError {
	baseMsg := GetSecuritySystemErrorMessage(errorCode)

	// Determine category from error code
	var category SecuritySystemErrorCategory
	switch {
	case strings.Contains(errorCode, "_00"):
		category = SecuritySystemCategoryConfiguration
	case strings.Contains(errorCode, "_10"):
		category = SecuritySystemCategoryBusinessLogic
	case strings.Contains(errorCode, "_20"):
		category = SecuritySystemCategoryAPIOperation
	case strings.Contains(errorCode, "_30"):
		category = SecuritySystemCategoryStateManagement
	default:
		category = SecuritySystemCategoryConfiguration
	}

	return &SecuritySystemStandardError{
		Code:           errorCode,
		Message:        baseMsg,
		Operation:      operation,
		SecuritySystem: securitySystem,
		Category:       category,
		OriginalError:  originalErr,
	}
}

// SecuritySystemOperationContext holds contextual information for operations
type SecuritySystemOperationContext struct {
	CorrelationID  string
	Operation      string
	SecuritySystem string
	StartTime      time.Time
	Resource       string
}

// CreateSecuritySystemOperationContext creates a new operation context with correlation ID
func CreateSecuritySystemOperationContext(operation, securitySystem string) *SecuritySystemOperationContext {
	return &SecuritySystemOperationContext{
		CorrelationID:  uuid.New().String(),
		Operation:      operation,
		SecuritySystem: securitySystem,
		StartTime:      time.Now(),
		Resource:       "security_system",
	}
}

// AddContextToLogger adds operation context to the logger
func (oc *SecuritySystemOperationContext) AddContextToLogger(ctx context.Context) context.Context {
	ctx = tflog.SetField(ctx, "correlation_id", oc.CorrelationID)
	ctx = tflog.SetField(ctx, "resource", oc.Resource)
	ctx = tflog.SetField(ctx, "operation", oc.Operation)
	if oc.SecuritySystem != "" {
		ctx = tflog.SetField(ctx, "security_system", oc.SecuritySystem)
	}
	return ctx
}

// GetBaseLogFields returns common log fields for the operation
func (oc *SecuritySystemOperationContext) GetBaseLogFields() map[string]interface{} {
	fields := map[string]interface{}{
		"correlation_id": oc.CorrelationID,
		"resource":       oc.Resource,
		"operation":      oc.Operation,
	}
	if oc.SecuritySystem != "" {
		fields["security_system"] = oc.SecuritySystem
	}
	return fields
}

// LogOperationStart logs the start of an operation
func (oc *SecuritySystemOperationContext) LogOperationStart(ctx context.Context, message string, additionalFields ...map[string]interface{}) {
	fields := oc.GetBaseLogFields()
	fields["start_time"] = oc.StartTime.Format(time.RFC3339)

	// Merge additional fields
	for _, additional := range additionalFields {
		for k, v := range additional {
			fields[k] = v
		}
	}

	tflog.Info(ctx, message, fields)
}

// LogOperationEnd logs the completion of an operation with duration
func (oc *SecuritySystemOperationContext) LogOperationEnd(ctx context.Context, message string, additionalFields ...map[string]interface{}) {
	duration := time.Since(oc.StartTime)
	fields := oc.GetBaseLogFields()
	fields["duration_ms"] = duration.Milliseconds()
	fields["end_time"] = time.Now().Format(time.RFC3339)

	// Merge additional fields
	for _, additional := range additionalFields {
		for k, v := range additional {
			fields[k] = v
		}
	}

	tflog.Info(ctx, message, fields)
}

// LogOperationError logs an error during operation with context
func (oc *SecuritySystemOperationContext) LogOperationError(ctx context.Context, message, errorCode string, err error, additionalFields ...map[string]interface{}) {
	duration := time.Since(oc.StartTime)
	fields := oc.GetBaseLogFields()
	fields["duration_ms"] = duration.Milliseconds()
	fields["error_code"] = errorCode
	fields["error"] = err.Error()

	// Merge additional fields
	for _, additional := range additionalFields {
		for k, v := range additional {
			fields[k] = v
		}
	}

	tflog.Error(ctx, message, fields)
}
