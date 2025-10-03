// Copyright (c) Saviynt Inc.
// SPDX-License-Identifier: MPL-2.0

package connectionsutil

import (
	"context"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ExactlyOneOfValidator struct {
	Attrs []path.Expression
}

func (v ExactlyOneOfValidator) Description(_ context.Context) string {
	return fmt.Sprintf("Exactly one of %s must be specified", attrNames(v.Attrs))
}

func (v ExactlyOneOfValidator) MarkdownDescription(_ context.Context) string {
	return v.Description(context.Background())
}

func (v ExactlyOneOfValidator) ValidateResource(
    ctx context.Context,
    req resource.ValidateConfigRequest,
    resp *resource.ValidateConfigResponse,
) {
	var count int
	var setAttrs []string
	var firstAttrPath path.Path

	for i, attrExpr := range v.Attrs {
		attrPath := attrExpr.Resolve()

		matches, diags := req.Config.PathMatches(ctx, attrPath)
		resp.Diagnostics.Append(diags...)
		if diags.HasError() || len(matches) == 0 {
			continue
		}
		if i == 0 {
			firstAttrPath = matches[0]
		}

		val := types.String{}
		diags = req.Config.GetAttribute(ctx, matches[0], &val)
		resp.Diagnostics.Append(diags...)
		if diags.HasError() {
			continue
		}

		if !val.IsNull() && !val.IsUnknown() {
			count++
			setAttrs = append(setAttrs, matches[0].String())
		}
	}

	if count == 0 {
		resp.Diagnostics.AddAttributeError(
			firstAttrPath,
			"Missing Required Attribute",
			fmt.Sprintf("One of the following attributes must be set: %s", attrNames(v.Attrs)),
		)
	} else if count > 1 {
		resp.Diagnostics.AddAttributeError(
			firstAttrPath,
			"Conflicting Attributes",
			fmt.Sprintf("Only one of the following attributes can be set: %s. You set: %s", attrNames(v.Attrs), strings.Join(setAttrs, ", ")),
		)
	}
}

type AtMostOneOfValidator struct {
	Attrs []path.Expression
}

func (v AtMostOneOfValidator) Description(_ context.Context) string {
	return fmt.Sprintf("At most one of %s can be specified", attrNames(v.Attrs))
}

func (v AtMostOneOfValidator) MarkdownDescription(_ context.Context) string {
	return v.Description(context.Background())
}

func (v AtMostOneOfValidator) ValidateResource(
    ctx context.Context,
    req resource.ValidateConfigRequest,
    resp *resource.ValidateConfigResponse,
) {
	var count int
	var setAttrs []string
	var firstAttrPath path.Path

	for i, attrExpr := range v.Attrs {
		attrPath := attrExpr.Resolve()

		matches, diags := req.Config.PathMatches(ctx, attrPath)
		resp.Diagnostics.Append(diags...)
		if diags.HasError() || len(matches) == 0 {
			continue
		}
		if i == 0 {
			firstAttrPath = matches[0]
		}

		val := types.String{}
		diags = req.Config.GetAttribute(ctx, matches[0], &val)
		resp.Diagnostics.Append(diags...)
		if diags.HasError() {
			continue
		}

		if !val.IsNull() && !val.IsUnknown() {
			count++
			setAttrs = append(setAttrs, matches[0].String())
		}
	}

	if count > 1 {
		resp.Diagnostics.AddAttributeError(
			firstAttrPath,
			"Conflicting Attributes",
			fmt.Sprintf("Only one of the following attributes can be set: %s. You set: %s", attrNames(v.Attrs), strings.Join(setAttrs, ", ")),
		)
	}
}


func attrNames(exprs []path.Expression) string {
	names := make([]string, len(exprs))
	for i, e := range exprs {
		names[i] = fmt.Sprintf("%q", e.String())
	}
	return strings.Join(names, ", ")
}