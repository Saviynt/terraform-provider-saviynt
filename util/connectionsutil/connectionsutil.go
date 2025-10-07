// Copyright (c) 2025 Saviynt Inc.
// SPDX-License-Identifier: MPL-2.0

package connectionsutil

import (
	datasource "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	resource "github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

func MergeResourceAttributes(a, b map[string]resource.Attribute) map[string]resource.Attribute {
	out := make(map[string]resource.Attribute, len(a)+len(b))
	for k, v := range a {
		out[k] = v
	}
	for k, v := range b {
		out[k] = v
	}
	return out
}

func MergeDataSourceAttributes(a, b map[string]datasource.Attribute) map[string]datasource.Attribute {
	out := make(map[string]datasource.Attribute, len(a)+len(b))
	for k, v := range a {
		out[k] = v
	}
	for k, v := range b {
		out[k] = v
	}
	return out
}
