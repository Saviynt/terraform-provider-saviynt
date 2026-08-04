// Copyright (c) 2025 Saviynt Inc.
// SPDX-License-Identifier: MPL-2.0

// Package validators provides custom Terraform schema validators for the
// Saviynt provider.
package validators

import (
	"context"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

// NoWhitespaceOnly returns a validator that rejects values consisting entirely
// of whitespace characters (spaces, tabs, newlines, etc.).
//
// Empty string "" is allowed — this validator only rejects values where
// strings.TrimSpace(value) == "" AND value != "".
//
// Rationale: Saviynt normalises whitespace-only custom property values to ""
// internally. If a user sets custom_property1 = " ", Saviynt stores "" and
// returns it as omitted (omitempty). This causes a perpetual plan diff because
// the config holds " " while the API always returns "". Rejecting whitespace-only
// values at plan time surfaces the problem immediately with a clear error rather
// than a confusing inconsistency error after apply.
func NoWhitespaceOnly() validator.String {
	return noWhitespaceOnlyValidator{}
}

type noWhitespaceOnlyValidator struct{}

func (v noWhitespaceOnlyValidator) Description(_ context.Context) string {
	return "Value must not consist entirely of whitespace. Use an empty string \"\" to represent an empty value."
}

func (v noWhitespaceOnlyValidator) MarkdownDescription(_ context.Context) string {
	return "Value must not consist entirely of whitespace. Use an empty string `\"\"` to represent an empty value."
}

func (v noWhitespaceOnlyValidator) ValidateString(_ context.Context, req validator.StringRequest, resp *validator.StringResponse) {
	if req.ConfigValue.IsNull() || req.ConfigValue.IsUnknown() {
		return
	}
	val := req.ConfigValue.ValueString()
	if val != "" && strings.TrimSpace(val) == "" {
		resp.Diagnostics.AddAttributeError(
			req.Path,
			"Whitespace-only value not allowed",
			"Value must not be whitespace-only. Use \"\" to set the field to empty.",
		)
	}
}
