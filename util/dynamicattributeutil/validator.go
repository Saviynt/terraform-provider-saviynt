/*
 * Copyright (c) 2025 Saviynt Inc.
 * All Rights Reserved.
 *
 * This software is the confidential and proprietary information of
 * Saviynt Inc. ("Confidential Information"). You shall not disclose,
 * use, or distribute such Confidential Information except in accordance
 * with the terms of the license agreement you entered into with Saviynt.
 *
 * SAVIYNT MAKES NO REPRESENTATIONS OR WARRANTIES ABOUT THE SUITABILITY OF
 * THE SOFTWARE, EITHER EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO
 * THE IMPLIED WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR
 * PURPOSE, OR NON-INFRINGEMENT.
 */

package dynamicattributeutil

import (
	"context"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
)

var _ validator.String = &attributeValueDisallowedForAttributeTypesValidator{}

type attributeValueDisallowedForAttributeTypesValidator struct {
}

func (v attributeValueDisallowedForAttributeTypesValidator) Description(ctx context.Context) string {
	return "Ensures attribute_value is not set for disallowed attribute_type values like LABEL, DATE, etc."
}

func (v attributeValueDisallowedForAttributeTypesValidator) MarkdownDescription(ctx context.Context) string {
	return v.Description(ctx)
}

func (v attributeValueDisallowedForAttributeTypesValidator) ValidateString(ctx context.Context, req validator.StringRequest, resp *validator.StringResponse) {
	if req.ConfigValue.IsNull() || req.ConfigValue.IsUnknown() {
		return
	}

	siblingPath := req.Path.ParentPath().AtName("attribute_type")

	var attrTypeVal attr.Value
	diags := req.Config.GetAttribute(ctx, siblingPath, &attrTypeVal)
	resp.Diagnostics.Append(diags...)
	if diags.HasError() || attrTypeVal.IsNull() || attrTypeVal.IsUnknown() {
		return
	}

	var attrType string
	diags = tfsdk.ValueAs(ctx, attrTypeVal, &attrType)
	resp.Diagnostics.Append(diags...)
	if diags.HasError() {
		return
	}

	switch strings.ToUpper(attrType) {
	case "NUMBER", "PASSWORD", "STRING", "BOOLEAN", "LARGE TEXT", "DATE":
		resp.Diagnostics.AddAttributeError(
			req.Path,
			"Invalid attribute_value",
			fmt.Sprintf("attribute_value must not be set when attribute_type is '%s'", attrType),
		)
	}
}
// Factory function with built-in disallowed types
func AttributeValueDisallowedForCertainAttributeTypes() validator.String {
	return &attributeValueDisallowedForAttributeTypesValidator{}
}

// --------------------------------------------------------------
var _ validator.String = &descriptionDisallowedForAttributeTypesValidator{}

type descriptionDisallowedForAttributeTypesValidator struct {
}

func (v descriptionDisallowedForAttributeTypesValidator) Description(ctx context.Context) string {
	return "Ensures description is not set for disallowed attribute_type like SQL ENUM etc"
}

func (v descriptionDisallowedForAttributeTypesValidator) MarkdownDescription(ctx context.Context) string {
	return v.Description(ctx)
}

func (v descriptionDisallowedForAttributeTypesValidator) ValidateString(ctx context.Context, req validator.StringRequest, resp *validator.StringResponse) {
	if req.ConfigValue.IsNull() || req.ConfigValue.IsUnknown() {
		return
	}

	siblingPath := req.Path.ParentPath().AtName("attribute_type")

	var attrTypeVal attr.Value
	diags := req.Config.GetAttribute(ctx, siblingPath, &attrTypeVal)
	resp.Diagnostics.Append(diags...)
	if diags.HasError() || attrTypeVal.IsNull() || attrTypeVal.IsUnknown() {
		return
	}

	var attrType string
	diags = tfsdk.ValueAs(ctx, attrTypeVal, &attrType)
	resp.Diagnostics.Append(diags...)
	if diags.HasError() {
		return
	}

	switch strings.ToUpper(attrType) {
	case "SQL ENUM", "SQL MULTISELECT":
		resp.Diagnostics.AddAttributeError(
			req.Path,
			"Invalid description",
			fmt.Sprintf("description must not be set when attribute_type is '%s'", attrType),
		)
	}
}

// Factory function with built-in disallowed types
func DescriptionDisallowedForCertainAttributeTypes() validator.String {
	return &descriptionDisallowedForAttributeTypesValidator{}
}

// --------------------------------------------------------------
var _ validator.String = &defaultValueDisallowedForAttributeTypesValidator{}

type defaultValueDisallowedForAttributeTypesValidator struct {
}

func (v defaultValueDisallowedForAttributeTypesValidator) Description(ctx context.Context) string {
	return "Ensures default value is not set for disallowed attribute_types"
}

func (v defaultValueDisallowedForAttributeTypesValidator) MarkdownDescription(ctx context.Context) string {
	return v.Description(ctx)
}

func (v defaultValueDisallowedForAttributeTypesValidator) ValidateString(ctx context.Context, req validator.StringRequest, resp *validator.StringResponse) {
	if req.ConfigValue.IsNull() || req.ConfigValue.IsUnknown() {
		return
	}

	siblingPath := req.Path.ParentPath().AtName("attribute_type")

	var attrTypeVal attr.Value
	diags := req.Config.GetAttribute(ctx, siblingPath, &attrTypeVal)
	resp.Diagnostics.Append(diags...)
	if diags.HasError() || attrTypeVal.IsNull() || attrTypeVal.IsUnknown() {
		return
	}

	var attrType string
	diags = tfsdk.ValueAs(ctx, attrTypeVal, &attrType)
	resp.Diagnostics.Append(diags...)
	if diags.HasError() {
		return
	}
	if strings.ToUpper(attrType) == "BOOLEAN" {
		resp.Diagnostics.AddAttributeError(
			req.Path,
			"Default value setting not allowed",
			fmt.Sprintf("default value is currently not configurable from Terraform when attribute_type is '%s'", attrType),
		)
	}

	if strings.ToUpper(attrType) == "PASSWORD" {
		resp.Diagnostics.AddAttributeError(
			req.Path,
			"Invalid default value",
			fmt.Sprintf("default value must not be set when attribute_type is '%s'", attrType),
		)
	}
}

// Factory function with built-in disallowed types
func DefaultValueDisallowedForCertainAttributeTypes() validator.String {
	return &defaultValueDisallowedForAttributeTypesValidator{}
}

// --------------------------------------------------------------
var _ validator.String = &regexDisallowedForAttributeTypesValidator{}

type regexDisallowedForAttributeTypesValidator struct {
}

func (v regexDisallowedForAttributeTypesValidator) Description(ctx context.Context) string {
	return "Ensures regex is not set for disallowed attribute_type."
}

func (v regexDisallowedForAttributeTypesValidator) MarkdownDescription(ctx context.Context) string {
	return v.Description(ctx)
}

func (v regexDisallowedForAttributeTypesValidator) ValidateString(ctx context.Context, req validator.StringRequest, resp *validator.StringResponse) {
	if req.ConfigValue.IsNull() || req.ConfigValue.IsUnknown() {
		return
	}

	siblingPath := req.Path.ParentPath().AtName("attribute_type")

	var attrTypeVal attr.Value
	diags := req.Config.GetAttribute(ctx, siblingPath, &attrTypeVal)
	resp.Diagnostics.Append(diags...)
	if diags.HasError() || attrTypeVal.IsNull() || attrTypeVal.IsUnknown() {
		return
	}

	var attrType string
	diags = tfsdk.ValueAs(ctx, attrTypeVal, &attrType)
	resp.Diagnostics.Append(diags...)
	if diags.HasError() {
		return
	}

	switch strings.ToUpper(attrType) {
	case "NUMBER", "STRING", "PASSWORD", "LARGE TEXT":
		// Allowed types – do nothing
	default:
		resp.Diagnostics.AddAttributeError(
			req.Path,
			"Invalid regex",
			fmt.Sprintf("regex must not be set when attribute_type is '%s'", attrType),
		)
	}

}

// Factory function with built-in disallowed types
func RegexDisallowedForCertainAttributeTypes() validator.String {
	return &regexDisallowedForAttributeTypesValidator{}
}