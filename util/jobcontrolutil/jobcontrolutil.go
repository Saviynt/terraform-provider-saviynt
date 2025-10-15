// Copyright (c) 2025 Saviynt Inc.
// SPDX-License-Identifier: MPL-2.0

package jobcontrolutil

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

// MergeJobControlResourceAttributes merges base job control attributes with specific job type attributes
func MergeJobResourceAttributes(base, specific map[string]schema.Attribute) map[string]schema.Attribute {
	merged := make(map[string]schema.Attribute)
	// Add base attributes
	for k, v := range base {
		merged[k] = v
	}
	// Add specific attributes (will override base if same key exists)
	for k, v := range specific {
		merged[k] = v
	}
	return merged
}
