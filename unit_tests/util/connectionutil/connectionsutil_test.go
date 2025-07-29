// Copyright (c) 2025 Saviynt Inc.
// SPDX-License-Identifier: MPL-2.0

package connectionsutil

import (
	"testing"

	datasource "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	resource "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/stretchr/testify/assert"

	"terraform-provider-Saviynt/util/connectionsutil"
)

func TestMergeResourceAttributes(t *testing.T) {
	tests := []struct {
		name     string
		inputA   map[string]resource.Attribute
		inputB   map[string]resource.Attribute
		expected map[string]resource.Attribute
	}{
		{
			name: "merge two non-empty maps",
			inputA: map[string]resource.Attribute{
				"attr1": resource.StringAttribute{
					Required:    true,
					Description: "First attribute",
				},
				"attr2": resource.Int64Attribute{
					Optional:    true,
					Description: "Second attribute",
				},
			},
			inputB: map[string]resource.Attribute{
				"attr3": resource.BoolAttribute{
					Computed:    true,
					Description: "Third attribute",
				},
				"attr4": resource.StringAttribute{
					Optional:    true,
					Description: "Fourth attribute",
				},
			},
			expected: map[string]resource.Attribute{
				"attr1": resource.StringAttribute{
					Required:    true,
					Description: "First attribute",
				},
				"attr2": resource.Int64Attribute{
					Optional:    true,
					Description: "Second attribute",
				},
				"attr3": resource.BoolAttribute{
					Computed:    true,
					Description: "Third attribute",
				},
				"attr4": resource.StringAttribute{
					Optional:    true,
					Description: "Fourth attribute",
				},
			},
		},
		{
			name:   "merge with empty first map",
			inputA: map[string]resource.Attribute{},
			inputB: map[string]resource.Attribute{
				"attr1": resource.StringAttribute{
					Required:    true,
					Description: "Only attribute",
				},
			},
			expected: map[string]resource.Attribute{
				"attr1": resource.StringAttribute{
					Required:    true,
					Description: "Only attribute",
				},
			},
		},
		{
			name: "merge with empty second map",
			inputA: map[string]resource.Attribute{
				"attr1": resource.StringAttribute{
					Required:    true,
					Description: "Only attribute",
				},
			},
			inputB: map[string]resource.Attribute{},
			expected: map[string]resource.Attribute{
				"attr1": resource.StringAttribute{
					Required:    true,
					Description: "Only attribute",
				},
			},
		},
		{
			name:     "merge two empty maps",
			inputA:   map[string]resource.Attribute{},
			inputB:   map[string]resource.Attribute{},
			expected: map[string]resource.Attribute{},
		},
		{
			name:   "merge with nil first map",
			inputA: nil,
			inputB: map[string]resource.Attribute{
				"attr1": resource.StringAttribute{
					Required:    true,
					Description: "Only attribute",
				},
			},
			expected: map[string]resource.Attribute{
				"attr1": resource.StringAttribute{
					Required:    true,
					Description: "Only attribute",
				},
			},
		},
		{
			name: "merge with nil second map",
			inputA: map[string]resource.Attribute{
				"attr1": resource.StringAttribute{
					Required:    true,
					Description: "Only attribute",
				},
			},
			inputB: nil,
			expected: map[string]resource.Attribute{
				"attr1": resource.StringAttribute{
					Required:    true,
					Description: "Only attribute",
				},
			},
		},
		{
			name:     "merge two nil maps",
			inputA:   nil,
			inputB:   nil,
			expected: map[string]resource.Attribute{},
		},
		{
			name: "merge with overlapping keys - second map overwrites first",
			inputA: map[string]resource.Attribute{
				"attr1": resource.StringAttribute{
					Required:    true,
					Description: "First version",
				},
				"attr2": resource.Int64Attribute{
					Optional:    true,
					Description: "Unique to A",
				},
			},
			inputB: map[string]resource.Attribute{
				"attr1": resource.StringAttribute{
					Optional:    true,
					Description: "Second version - overwrites first",
				},
				"attr3": resource.BoolAttribute{
					Computed:    true,
					Description: "Unique to B",
				},
			},
			expected: map[string]resource.Attribute{
				"attr1": resource.StringAttribute{
					Optional:    true,
					Description: "Second version - overwrites first",
				},
				"attr2": resource.Int64Attribute{
					Optional:    true,
					Description: "Unique to A",
				},
				"attr3": resource.BoolAttribute{
					Computed:    true,
					Description: "Unique to B",
				},
			},
		},
		{
			name: "merge with different attribute types",
			inputA: map[string]resource.Attribute{
				"string_attr": resource.StringAttribute{
					Required: true,
				},
				"list_attr": resource.ListAttribute{
					ElementType: types.StringType,
					Optional:    true,
				},
			},
			inputB: map[string]resource.Attribute{
				"set_attr": resource.SetAttribute{
					ElementType: types.StringType,
					Computed:    true,
				},
				"object_attr": resource.SingleNestedAttribute{
					Optional: true,
					Attributes: map[string]resource.Attribute{
						"nested": resource.StringAttribute{
							Required: true,
						},
					},
				},
			},
			expected: map[string]resource.Attribute{
				"string_attr": resource.StringAttribute{
					Required: true,
				},
				"list_attr": resource.ListAttribute{
					ElementType: types.StringType,
					Optional:    true,
				},
				"set_attr": resource.SetAttribute{
					ElementType: types.StringType,
					Computed:    true,
				},
				"object_attr": resource.SingleNestedAttribute{
					Optional: true,
					Attributes: map[string]resource.Attribute{
						"nested": resource.StringAttribute{
							Required: true,
						},
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := connectionsutil.MergeResourceAttributes(tt.inputA, tt.inputB)

			// Check that the result has the expected number of attributes
			assert.Equal(t, len(tt.expected), len(result))

			// Check each expected attribute exists and matches
			for key, expectedAttr := range tt.expected {
				actualAttr, exists := result[key]
				assert.True(t, exists, "Expected attribute %s to exist in result", key)

				// Compare attribute types and properties
				assert.Equal(t, expectedAttr, actualAttr, "Attribute %s should match expected value", key)
			}

			// Ensure no unexpected attributes exist
			for key := range result {
				_, exists := tt.expected[key]
				assert.True(t, exists, "Unexpected attribute %s found in result", key)
			}
		})
	}
}

func TestMergeDataSourceAttributes(t *testing.T) {
	tests := []struct {
		name     string
		inputA   map[string]datasource.Attribute
		inputB   map[string]datasource.Attribute
		expected map[string]datasource.Attribute
	}{
		{
			name: "merge two non-empty maps",
			inputA: map[string]datasource.Attribute{
				"attr1": datasource.StringAttribute{
					Required:    true,
					Description: "First attribute",
				},
				"attr2": datasource.Int64Attribute{
					Optional:    true,
					Description: "Second attribute",
				},
			},
			inputB: map[string]datasource.Attribute{
				"attr3": datasource.BoolAttribute{
					Computed:    true,
					Description: "Third attribute",
				},
				"attr4": datasource.StringAttribute{
					Optional:    true,
					Description: "Fourth attribute",
				},
			},
			expected: map[string]datasource.Attribute{
				"attr1": datasource.StringAttribute{
					Required:    true,
					Description: "First attribute",
				},
				"attr2": datasource.Int64Attribute{
					Optional:    true,
					Description: "Second attribute",
				},
				"attr3": datasource.BoolAttribute{
					Computed:    true,
					Description: "Third attribute",
				},
				"attr4": datasource.StringAttribute{
					Optional:    true,
					Description: "Fourth attribute",
				},
			},
		},
		{
			name:   "merge with empty first map",
			inputA: map[string]datasource.Attribute{},
			inputB: map[string]datasource.Attribute{
				"attr1": datasource.StringAttribute{
					Required:    true,
					Description: "Only attribute",
				},
			},
			expected: map[string]datasource.Attribute{
				"attr1": datasource.StringAttribute{
					Required:    true,
					Description: "Only attribute",
				},
			},
		},
		{
			name: "merge with empty second map",
			inputA: map[string]datasource.Attribute{
				"attr1": datasource.StringAttribute{
					Required:    true,
					Description: "Only attribute",
				},
			},
			inputB: map[string]datasource.Attribute{},
			expected: map[string]datasource.Attribute{
				"attr1": datasource.StringAttribute{
					Required:    true,
					Description: "Only attribute",
				},
			},
		},
		{
			name:     "merge two empty maps",
			inputA:   map[string]datasource.Attribute{},
			inputB:   map[string]datasource.Attribute{},
			expected: map[string]datasource.Attribute{},
		},
		{
			name:   "merge with nil first map",
			inputA: nil,
			inputB: map[string]datasource.Attribute{
				"attr1": datasource.StringAttribute{
					Required:    true,
					Description: "Only attribute",
				},
			},
			expected: map[string]datasource.Attribute{
				"attr1": datasource.StringAttribute{
					Required:    true,
					Description: "Only attribute",
				},
			},
		},
		{
			name: "merge with nil second map",
			inputA: map[string]datasource.Attribute{
				"attr1": datasource.StringAttribute{
					Required:    true,
					Description: "Only attribute",
				},
			},
			inputB: nil,
			expected: map[string]datasource.Attribute{
				"attr1": datasource.StringAttribute{
					Required:    true,
					Description: "Only attribute",
				},
			},
		},
		{
			name:     "merge two nil maps",
			inputA:   nil,
			inputB:   nil,
			expected: map[string]datasource.Attribute{},
		},
		{
			name: "merge with overlapping keys - second map overwrites first",
			inputA: map[string]datasource.Attribute{
				"attr1": datasource.StringAttribute{
					Required:    true,
					Description: "First version",
				},
				"attr2": datasource.Int64Attribute{
					Optional:    true,
					Description: "Unique to A",
				},
			},
			inputB: map[string]datasource.Attribute{
				"attr1": datasource.StringAttribute{
					Optional:    true,
					Description: "Second version - overwrites first",
				},
				"attr3": datasource.BoolAttribute{
					Computed:    true,
					Description: "Unique to B",
				},
			},
			expected: map[string]datasource.Attribute{
				"attr1": datasource.StringAttribute{
					Optional:    true,
					Description: "Second version - overwrites first",
				},
				"attr2": datasource.Int64Attribute{
					Optional:    true,
					Description: "Unique to A",
				},
				"attr3": datasource.BoolAttribute{
					Computed:    true,
					Description: "Unique to B",
				},
			},
		},
		{
			name: "merge with different attribute types",
			inputA: map[string]datasource.Attribute{
				"string_attr": datasource.StringAttribute{
					Required: true,
				},
				"list_attr": datasource.ListAttribute{
					ElementType: types.StringType,
					Optional:    true,
				},
			},
			inputB: map[string]datasource.Attribute{
				"set_attr": datasource.SetAttribute{
					ElementType: types.StringType,
					Computed:    true,
				},
				"object_attr": datasource.SingleNestedAttribute{
					Optional: true,
					Attributes: map[string]datasource.Attribute{
						"nested": datasource.StringAttribute{
							Required: true,
						},
					},
				},
			},
			expected: map[string]datasource.Attribute{
				"string_attr": datasource.StringAttribute{
					Required: true,
				},
				"list_attr": datasource.ListAttribute{
					ElementType: types.StringType,
					Optional:    true,
				},
				"set_attr": datasource.SetAttribute{
					ElementType: types.StringType,
					Computed:    true,
				},
				"object_attr": datasource.SingleNestedAttribute{
					Optional: true,
					Attributes: map[string]datasource.Attribute{
						"nested": datasource.StringAttribute{
							Required: true,
						},
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := connectionsutil.MergeDataSourceAttributes(tt.inputA, tt.inputB)

			// Check that the result has the expected number of attributes
			assert.Equal(t, len(tt.expected), len(result))

			// Check each expected attribute exists and matches
			for key, expectedAttr := range tt.expected {
				actualAttr, exists := result[key]
				assert.True(t, exists, "Expected attribute %s to exist in result", key)

				// Compare attribute types and properties
				assert.Equal(t, expectedAttr, actualAttr, "Attribute %s should match expected value", key)
			}

			// Ensure no unexpected attributes exist
			for key := range result {
				_, exists := tt.expected[key]
				assert.True(t, exists, "Unexpected attribute %s found in result", key)
			}
		})
	}
}
