// Copyright (c) 2025 Saviynt Inc.
// SPDX-License-Identifier: MPL-2.0

// Package errorsutil provides standardized error handling utilities for Saviynt Terraform provider.
// It includes error code management, message standardization, sensitive data sanitization,
// and operation context tracking for all connector resources.
package errorsutil

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// ConnectorType represents the type of connector for error context
type ConnectorType string

const (
	ConnectorTypeAD         ConnectorType = "AD"
	ConnectorTypeREST       ConnectorType = "REST"
	ConnectorTypeADSI       ConnectorType = "ADSI"
	ConnectorTypeDB         ConnectorType = "DB"
	ConnectorTypeEntraID    ConnectorType = "ENTRAID"
	ConnectorTypeSAP        ConnectorType = "SAP"
	ConnectorTypeSalesforce ConnectorType = "SALESFORCE"
	ConnectorTypeWorkday    ConnectorType = "WORKDAY"
	ConnectorTypeUnix       ConnectorType = "UNIX"
	ConnectorTypeGithubREST ConnectorType = "GITHUBREST"
	ConnectorTypeOkta       ConnectorType = "OKTA"
)

// ErrorCategory represents different categories of errors
type ErrorCategory string

const (
	CategoryConfiguration   ErrorCategory = "CONFIG"
	CategoryBusinessLogic   ErrorCategory = "BUSINESS"
	CategoryAPIOperation    ErrorCategory = "API"
	CategoryStateManagement ErrorCategory = "STATE"
)

// StandardError represents a standardized error with context
type StandardError struct {
	Code           string
	Message        string
	Operation      string
	ConnectorType  ConnectorType
	ConnectionName string
	Category       ErrorCategory
	OriginalError  error
}

// Error implements the error interface
func (se *StandardError) Error() string {
	if se.OriginalError != nil {
		return fmt.Sprintf("[%s] %s during %s operation for %s connection '%s': %v",
			se.Code, se.Message, se.Operation, se.ConnectorType, se.ConnectionName, se.OriginalError)
	}
	return fmt.Sprintf("[%s] %s during %s operation for %s connection '%s'",
		se.Code, se.Message, se.Operation, se.ConnectorType, se.ConnectionName)
}

// Unwrap returns the original error for error unwrapping
func (se *StandardError) Unwrap() error {
	return se.OriginalError
}

// ErrorCodeGenerator generates standardized error codes
type ErrorCodeGenerator struct {
	connectorType ConnectorType
}

// NewErrorCodeGenerator creates a new error code generator for a specific connector type
func NewErrorCodeGenerator(connectorType ConnectorType) *ErrorCodeGenerator {
	return &ErrorCodeGenerator{
		connectorType: connectorType,
	}
}

// GenerateErrorCode generates a standardized error code
func (ecg *ErrorCodeGenerator) GenerateErrorCode(category ErrorCategory, sequence int) string {
	var categoryCode string
	switch category {
	case CategoryConfiguration:
		categoryCode = "00"
	case CategoryBusinessLogic:
		categoryCode = "10"
	case CategoryAPIOperation:
		categoryCode = "20"
	case CategoryStateManagement:
		categoryCode = "30"
	default:
		categoryCode = "99"
	}

	return fmt.Sprintf("%s_CONN_%s%d", ecg.connectorType, categoryCode, sequence)
}

// Common error codes that can be used across all connectors
const (
	// Configuration errors (001-009)
	ErrProviderConfig   = "PROVIDER_CONFIG"
	ErrPlanExtraction   = "PLAN_EXTRACTION"
	ErrConfigExtraction = "CONFIG_EXTRACTION"
	ErrStateExtraction  = "STATE_EXTRACTION"

	// Business logic errors (101-109)
	ErrDuplicateName = "DUPLICATE_NAME"
	ErrNameImmutable = "NAME_IMMUTABLE"
	ErrInvalidConfig = "INVALID_CONFIG"

	// API operation errors (201-209)
	ErrCreateFailed = "CREATE_FAILED"
	ErrReadFailed   = "READ_FAILED"
	ErrUpdateFailed = "UPDATE_FAILED"
	ErrDeleteFailed = "DELETE_FAILED"
	ErrAPIError     = "API_ERROR"

	// State management errors (301-309)
	ErrStateUpdate = "STATE_UPDATE"
	ErrStateRead   = "STATE_READ"
)

// Common error messages
var commonErrorMessages = map[string]string{
	ErrProviderConfig:   "Failed to configure provider",
	ErrPlanExtraction:   "Failed to extract plan from Terraform request",
	ErrConfigExtraction: "Failed to extract configuration from Terraform request",
	ErrStateExtraction:  "Failed to extract state from Terraform request",
	ErrDuplicateName:    "Connection with this name already exists",
	ErrNameImmutable:    "Connection name cannot be modified after creation",
	ErrInvalidConfig:    "Invalid configuration provided",
	ErrCreateFailed:     "Failed to create connection",
	ErrReadFailed:       "Failed to read connection details",
	ErrUpdateFailed:     "Failed to update connection",
	ErrDeleteFailed:     "Failed to delete connection",
	ErrAPIError:         "API operation returned an error",
	ErrStateUpdate:      "Failed to update Terraform state",
	ErrStateRead:        "Failed to read Terraform state",
}

// GetErrorMessage returns a standardized error message for the given error code
func GetErrorMessage(errorCode string) string {
	// First check if it's a common error code
	if msg, exists := commonErrorMessages[errorCode]; exists {
		return msg
	}

	// Extract connector type from error code for specific messages
	var connectorType string
	if strings.HasPrefix(errorCode, "AD_CONN_") {
		connectorType = "AD"
	} else if strings.HasPrefix(errorCode, "REST_CONN_") {
		connectorType = "REST"
	} else if strings.HasPrefix(errorCode, "ADSI_CONN_") {
		connectorType = "ADSI"
	} else if strings.HasPrefix(errorCode, "DB_CONN_") {
		connectorType = "DB"
	} else if strings.HasPrefix(errorCode, "ENTRAID_CONN_") {
		connectorType = "EntraID"
	} else if strings.HasPrefix(errorCode, "SAP_CONN_") {
		connectorType = "SAP"
	} else if strings.HasPrefix(errorCode, "SALESFORCE_CONN_") {
		connectorType = "Salesforce"
	} else if strings.HasPrefix(errorCode, "WORKDAY_CONN_") {
		connectorType = "Workday"
	} else if strings.HasPrefix(errorCode, "UNIX_CONN_") {
		connectorType = "Unix"
	} else if strings.HasPrefix(errorCode, "GITHUBREST_CONN_") {
		connectorType = "GithubREST"
	} else if strings.HasPrefix(errorCode, "OKTA_CONN_") {
		connectorType = "Okta"
	}

	// Then check if it's a connector-specific error code and map to specific message
	switch {
	case strings.Contains(errorCode, "_CONN_001"):
		return fmt.Sprintf("Failed to configure %s connection provider", connectorType)
	case strings.Contains(errorCode, "_CONN_002"):
		return commonErrorMessages[ErrPlanExtraction]
	case strings.Contains(errorCode, "_CONN_003"):
		return commonErrorMessages[ErrConfigExtraction]
	case strings.Contains(errorCode, "_CONN_004"):
		return commonErrorMessages[ErrStateExtraction]
	case strings.Contains(errorCode, "_CONN_101"):
		return fmt.Sprintf("%s connection with this name already exists (Please import or use a different name)", connectorType)
	case strings.Contains(errorCode, "_CONN_102"):
		return fmt.Sprintf("%s connection name cannot be modified after creation", connectorType)
	case strings.Contains(errorCode, "_CONN_201"):
		return fmt.Sprintf("Failed to create %s connection", connectorType)
	case strings.Contains(errorCode, "_CONN_202"):
		return fmt.Sprintf("Failed to read %s connection", connectorType)
	case strings.Contains(errorCode, "_CONN_203"):
		return fmt.Sprintf("Failed to update %s connection", connectorType)
	case strings.Contains(errorCode, "_CONN_204"):
		return fmt.Sprintf("%s connection API operation returned an error", connectorType)
	case strings.Contains(errorCode, "_CONN_301"):
		return commonErrorMessages[ErrStateUpdate]
	default:
		return "Unknown error occurred"
	}
}

// CreateStandardError creates a standardized error with code and context
func CreateStandardError(connectorType ConnectorType, errorCode, operation, connectionName string, originalErr error) *StandardError {
	baseMsg := GetErrorMessage(errorCode)

	// Determine category from error code
	var category ErrorCategory
	switch {
	case strings.Contains(errorCode, "_00"):
		category = CategoryConfiguration
	case strings.Contains(errorCode, "_10"):
		category = CategoryBusinessLogic
	case strings.Contains(errorCode, "_20"):
		category = CategoryAPIOperation
	case strings.Contains(errorCode, "_30"):
		category = CategoryStateManagement
	default:
		category = CategoryConfiguration
	}

	return &StandardError{
		Code:           errorCode,
		Message:        baseMsg,
		Operation:      operation,
		ConnectorType:  connectorType,
		ConnectionName: connectionName,
		Category:       category,
		OriginalError:  originalErr,
	}
}

// SanitizeError removes sensitive information from error messages
func SanitizeError(err error) string {
	if err == nil {
		return ""
	}

	errStr := err.Error()

	// List of sensitive keywords to redact
	sensitiveKeywords := []string{
		"password", "Password", "PASSWORD",
		"token", "Token", "TOKEN",
		"secret", "Secret", "SECRET",
		"key", "Key", "KEY",
		"credential", "Credential", "CREDENTIAL",
		"auth", "Auth", "AUTH",
		"bearer", "Bearer", "BEARER",
	}

	// Replace sensitive information
	for _, keyword := range sensitiveKeywords {
		if strings.Contains(strings.ToLower(errStr), strings.ToLower(keyword)) {
			return "Error contains sensitive information - details redacted"
		}
	}

	return errStr
}

// SanitizeMessage removes sensitive information from API response messages
func SanitizeMessage(msg *string) string {
	if msg == nil {
		return ""
	}

	msgStr := *msg
	sensitiveKeywords := []string{
		"password", "Password", "PASSWORD",
		"token", "Token", "TOKEN",
		"secret", "Secret", "SECRET",
		"auth", "Auth", "AUTH",
		"bearer", "Bearer", "BEARER",
	}

	for _, keyword := range sensitiveKeywords {
		if strings.Contains(strings.ToLower(msgStr), strings.ToLower(keyword)) {
			return "Message contains sensitive information - details redacted"
		}
	}

	return msgStr
}

// OperationContext holds contextual information for operations
type OperationContext struct {
	CorrelationID  string
	Operation      string
	ConnectionName string
	ConnectorType  ConnectorType
	StartTime      time.Time
	Resource       string
}

// CreateOperationContext creates a new operation context with correlation ID
func CreateOperationContext(connectorType ConnectorType, operation, connectionName string) *OperationContext {
	return &OperationContext{
		CorrelationID:  uuid.New().String(),
		Operation:      operation,
		ConnectionName: connectionName,
		ConnectorType:  connectorType,
		StartTime:      time.Now(),
		Resource:       fmt.Sprintf("%s_connection", strings.ToLower(string(connectorType))),
	}
}

// AddContextToLogger adds operation context to the logger
func (oc *OperationContext) AddContextToLogger(ctx context.Context) context.Context {
	ctx = tflog.SetField(ctx, "correlation_id", oc.CorrelationID)
	ctx = tflog.SetField(ctx, "resource", oc.Resource)
	ctx = tflog.SetField(ctx, "operation", oc.Operation)
	ctx = tflog.SetField(ctx, "connector_type", string(oc.ConnectorType))
	if oc.ConnectionName != "" {
		ctx = tflog.SetField(ctx, "connection_name", oc.ConnectionName)
	}
	return ctx
}

// GetBaseLogFields returns common log fields for the operation
func (oc *OperationContext) GetBaseLogFields() map[string]interface{} {
	fields := map[string]interface{}{
		"correlation_id": oc.CorrelationID,
		"resource":       oc.Resource,
		"operation":      oc.Operation,
		"connector_type": string(oc.ConnectorType),
	}
	if oc.ConnectionName != "" {
		fields["connection_name"] = oc.ConnectionName
	}
	return fields
}

// LogOperationStart logs the start of an operation
func (oc *OperationContext) LogOperationStart(ctx context.Context, message string, additionalFields ...map[string]interface{}) {
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
func (oc *OperationContext) LogOperationEnd(ctx context.Context, message string, additionalFields ...map[string]interface{}) {
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
func (oc *OperationContext) LogOperationError(ctx context.Context, message, errorCode string, err error, additionalFields ...map[string]interface{}) {
	duration := time.Since(oc.StartTime)
	fields := oc.GetBaseLogFields()
	fields["duration_ms"] = duration.Milliseconds()
	fields["error_code"] = errorCode
	fields["error"] = SanitizeError(err)

	// Merge additional fields
	for _, additional := range additionalFields {
		for k, v := range additional {
			fields[k] = v
		}
	}

	tflog.Error(ctx, message, fields)
}

// ConnectorErrorCodes provides error code generation for specific connectors
type ConnectorErrorCodes struct {
	generator *ErrorCodeGenerator
}

// NewConnectorErrorCodes creates a new connector error codes instance
func NewConnectorErrorCodes(connectorType ConnectorType) *ConnectorErrorCodes {
	return &ConnectorErrorCodes{
		generator: NewErrorCodeGenerator(connectorType),
	}
}

// Configuration error codes
func (cec *ConnectorErrorCodes) ProviderConfig() string {
	return cec.generator.GenerateErrorCode(CategoryConfiguration, 1)
}

func (cec *ConnectorErrorCodes) PlanExtraction() string {
	return cec.generator.GenerateErrorCode(CategoryConfiguration, 2)
}

func (cec *ConnectorErrorCodes) ConfigExtraction() string {
	return cec.generator.GenerateErrorCode(CategoryConfiguration, 3)
}

func (cec *ConnectorErrorCodes) StateExtraction() string {
	return cec.generator.GenerateErrorCode(CategoryConfiguration, 4)
}

// Business logic error codes
func (cec *ConnectorErrorCodes) DuplicateName() string {
	return cec.generator.GenerateErrorCode(CategoryBusinessLogic, 1)
}

func (cec *ConnectorErrorCodes) NameImmutable() string {
	return cec.generator.GenerateErrorCode(CategoryBusinessLogic, 2)
}

// API operation error codes
func (cec *ConnectorErrorCodes) CreateFailed() string {
	return cec.generator.GenerateErrorCode(CategoryAPIOperation, 1)
}

func (cec *ConnectorErrorCodes) ReadFailed() string {
	return cec.generator.GenerateErrorCode(CategoryAPIOperation, 2)
}

func (cec *ConnectorErrorCodes) UpdateFailed() string {
	return cec.generator.GenerateErrorCode(CategoryAPIOperation, 3)
}

func (cec *ConnectorErrorCodes) APIError() string {
	return cec.generator.GenerateErrorCode(CategoryAPIOperation, 4)
}

// State management error codes
func (cec *ConnectorErrorCodes) StateUpdate() string {
	return cec.generator.GenerateErrorCode(CategoryStateManagement, 1)
}
