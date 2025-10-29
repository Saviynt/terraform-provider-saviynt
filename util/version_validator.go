// Copyright (c) 2025 Saviynt Inc.
// SPDX-License-Identifier: MPL-2.0

package util

import (
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/diag"
)

// AttributeVersionCompatibility defines which attributes are supported in which versions
type AttributeVersionCompatibility struct {
	AttributeName    string
	SupportedIn25B   bool
	SupportedIn25A   bool
	SupportedIn24_10 bool
	ResourceType     string
}

// VersionCompatibilityMap contains the compatibility matrix from README.md
var VersionCompatibilityMap = []AttributeVersionCompatibility{
	// Workday Connector
	{"ORGROLE_IMPORT_PAYLOAD", true, false, false, "Workday"},

	// REST Connector
	{"ApplicationDiscoveryJSON", true, false, false, "REST"},
	{"CreateEntitlementJSON", true, false, false, "REST"},
	{"DeleteEntitlementJSON", true, false, false, "REST"},
	{"UpdateEntitlementJSON", true, false, false, "REST"},

	// DB Connector
	{"CREATEENTITLEMENTJSON", true, false, false, "DB"},
	{"DELETEENTITLEMENTJSON", true, false, false, "DB"},
	{"ENTITLEMENTEXISTJSON", true, false, false, "DB"},
	{"UPDATEENTITLEMENTJSON", true, false, false, "DB"},

	// GitHub REST Connector
	{"status_threshold_config", true, true, false, "GitHubREST"},

	// Security System
	{"instant_provisioning", true, false, false, "SecuritySystem"},

	// Entitlement Type
	{"enable_entitlement_to_role_sync", true, true, false, "EntitlementType"},

	// Enterprise Role
	{"child_roles", true, false, false, "EnterpriseRole"},
}

// ValidateAttributeCompatibility checks if an attribute is supported in the given Saviynt version
func ValidateAttributeCompatibility(saviyntVersion, resourceType, attributeName string, attributeValue interface{}, diags *diag.Diagnostics) {
	// Skip validation if version is empty or attribute value is empty/null
	if saviyntVersion == "" || isEmptyValue(attributeValue) {
		return
	}

	// Check if version is supported
	if !isSupportedVersion(saviyntVersion) {
		diags.AddWarning(
			"Unsupported Saviynt Version Detected",
			fmt.Sprintf("Saviynt version '%s' is not officially supported. "+
				"This provider supports versions: 25.Brisbane (25.B), 25.Amsterdam (25.A), and 24.10. "+
				"Some features may not work as expected. "+
				"Please check compatibility at: "+
				"https://registry.terraform.io/providers/saviynt/saviynt/latest/docs#supported-saviynt-versions-by-provider",
				saviyntVersion),
		)
		return // Skip attribute validation for unsupported versions
	}

	// Find the attribute in compatibility map
	for _, compat := range VersionCompatibilityMap {
		if compat.AttributeName == attributeName && compat.ResourceType == resourceType {
			isSupported := isAttributeSupported(saviyntVersion, compat)

			if !isSupported {
				diags.AddError(
					fmt.Sprintf("Unsupported Attribute for Saviynt Version %s", saviyntVersion),
					fmt.Sprintf("The attribute '%s' is not supported in Saviynt version %s for %s resource. "+
						"Please check the compatibility matrix at: "+
						"https://registry.terraform.io/providers/saviynt/saviynt/latest/docs#attribute-compatibility-by-eic-version",
						attributeName, saviyntVersion, resourceType),
				)
			}
			return
		}
	}
}

// isSupportedVersion checks if the version is one of the officially supported versions
func isSupportedVersion(version string) bool {
	version = strings.ToLower(strings.TrimSpace(version))

	// Check for supported versions
	supportedVersions := []string{"25.brisbane", "25.b", "25.amsterdam", "25.a", "24.10"}

	for _, supported := range supportedVersions {
		if strings.Contains(version, supported) {
			return true
		}
	}

	return false
}

// isAttributeSupported checks if attribute is supported in the given version
func isAttributeSupported(version string, compat AttributeVersionCompatibility) bool {
	version = strings.ToLower(strings.TrimSpace(version))

	// Check for 25.Brisbane (25.B)
	if strings.Contains(version, "25.brisbane") || strings.Contains(version, "25.b") {
		return compat.SupportedIn25B
	}

	// Check for 25.Amsterdam (25.A)
	if strings.Contains(version, "25.amsterdam") || strings.Contains(version, "25.a") {
		return compat.SupportedIn25A
	}

	// Check for 24.10
	if strings.Contains(version, "24.10") {
		return compat.SupportedIn24_10
	}

	// Default to true for unknown versions to avoid blocking
	return true
}

// isEmptyValue checks if the attribute value is empty/null
func isEmptyValue(value interface{}) bool {
	if value == nil {
		return true
	}

	switch v := value.(type) {
	case string:
		return strings.TrimSpace(v) == ""
	case *string:
		return v == nil || strings.TrimSpace(*v) == ""
	default:
		return false
	}
}
