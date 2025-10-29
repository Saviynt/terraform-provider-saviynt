// Copyright (c) 2025 Saviynt Inc.
// SPDX-License-Identifier: MPL-2.0

package rolesutil

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ validator.Set = &ownerNameLimitValidator{}

type ownerNameLimitValidator struct {
}

func (v ownerNameLimitValidator) Description(ctx context.Context) string {
	return "Ensures the same owner_name does not appear more than 5 times"
}

func (v ownerNameLimitValidator) MarkdownDescription(ctx context.Context) string {
	return v.Description(ctx)
}

func (v ownerNameLimitValidator) ValidateSet(ctx context.Context, req validator.SetRequest, resp *validator.SetResponse) {
	if req.ConfigValue.IsNull() || req.ConfigValue.IsUnknown() {
		return
	}

	// Extract owners from the set
	var owners []struct {
		OwnerName types.String `tfsdk:"owner_name"`
		Rank      types.String `tfsdk:"rank"`
	}

	diags := req.ConfigValue.ElementsAs(ctx, &owners, false)
	resp.Diagnostics.Append(diags...)
	if diags.HasError() {
		return
	}

	// Count occurrences per owner_name
	ownerCount := make(map[string]int)
	for _, owner := range owners {
		ownerName := owner.OwnerName.ValueString()
		ownerCount[ownerName]++
	}

	// Check if any owner_name appears more than 5 times
	for ownerName, count := range ownerCount {
		if count > 5 {
			resp.Diagnostics.AddAttributeError(
				req.Path,
				"Too many occurrences for owner",
				fmt.Sprintf("Owner '%s' appears %d times, but maximum allowed is 5", ownerName, count),
			)
		}
	}
}

// Factory function
func OwnerNameAddLimit() validator.Set {
	return &ownerNameLimitValidator{}
}

// Child role validator
var _ validator.Set = &childRoleValidator{}

type childRoleValidator struct{}

func (v childRoleValidator) Description(ctx context.Context) string {
	return "Validates child roles and shows version warning"
}

func (v childRoleValidator) MarkdownDescription(ctx context.Context) string {
	return v.Description(ctx)
}

func (v childRoleValidator) ValidateSet(ctx context.Context, req validator.SetRequest, resp *validator.SetResponse) {
	if req.ConfigValue.IsNull() || req.ConfigValue.IsUnknown() {
		return
	}

	// Extract child roles from the set
	var childRoles []struct {
		RoleName types.String `tfsdk:"role_name"`
	}

	diags := req.ConfigValue.ElementsAs(ctx, &childRoles, false)
	resp.Diagnostics.Append(diags...)
	if diags.HasError() {
		return
	}

	// If child_roles length > 0, show version warning
	if len(childRoles) > 0 {
		resp.Diagnostics.AddAttributeWarning(
			req.Path,
			"Version Compatibility Warning",
			"The 'child_roles' attribute is only supported in Saviynt version 25.Brisbane and later. "+
				"Please ensure your Saviynt instance is running version 25.Brisbane or this attribute may not work as expected.",
		)
	}
}

// Factory function for child role validator
func ChildRoleValidator() validator.Set {
	return &childRoleValidator{}
}

// Owner rank validator
var _ validator.Set = &ownerRankValidator{}

type ownerRankValidator struct{}

func (v ownerRankValidator) Description(ctx context.Context) string {
	return "Validates that owner ranks are between 1-27"
}

func (v ownerRankValidator) MarkdownDescription(ctx context.Context) string {
	return v.Description(ctx)
}

func (v ownerRankValidator) ValidateSet(ctx context.Context, req validator.SetRequest, resp *validator.SetResponse) {
	if req.ConfigValue.IsNull() || req.ConfigValue.IsUnknown() {
		return
	}

	// Extract owners from the set
	var owners []struct {
		OwnerName types.String `tfsdk:"owner_name"`
		Rank      types.String `tfsdk:"rank"`
	}

	diags := req.ConfigValue.ElementsAs(ctx, &owners, false)
	resp.Diagnostics.Append(diags...)
	if diags.HasError() {
		return
	}

	// Validate each rank
	for _, owner := range owners {
		ownerName := owner.OwnerName.ValueString()
		rankStr := owner.Rank.ValueString()

		if rankStr != "" {
			var rank int
			if _, err := fmt.Sscanf(rankStr, "%d", &rank); err != nil {
				resp.Diagnostics.AddAttributeError(
					req.Path,
					"Invalid rank format",
					fmt.Sprintf("Rank '%s' for owner '%s' must be a number between 1-27, where 1 is highest priority", rankStr, ownerName),
				)
				continue
			}

			if rank < 1 || rank > 27 {
				resp.Diagnostics.AddAttributeError(
					req.Path,
					"Invalid rank value",
					fmt.Sprintf("Rank '%d' for owner '%s' must be between 1-27, where 1 is highest priority", rank, ownerName),
				)
			}
		}
	}
}

// Factory function for owner rank validator
func OwnerRankValidator() validator.Set {
	return &ownerRankValidator{}
}
