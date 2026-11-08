// Copyright (c) 2025 Saviynt Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"context"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
)

// connectionResourcesForTest returns every connection resource registered by the provider.
func connectionResourcesForTest(t *testing.T) map[string]resource.Resource {
	t.Helper()
	out := map[string]resource.Resource{}
	for _, f := range (&SaviyntProvider{}).Resources(context.Background()) {
		r := f()
		var mr resource.MetadataResponse
		r.Metadata(context.Background(), resource.MetadataRequest{ProviderTypeName: "saviynt"}, &mr)
		if strings.HasSuffix(mr.TypeName, "_connection_resource") {
			out[mr.TypeName] = r
		}
	}
	return out
}

func hasModifier(mods []planmodifier.String, want planmodifier.String) bool {
	for _, m := range mods {
		if m.Description(context.Background()) == want.Description(context.Background()) {
			return true
		}
	}
	return false
}

func hasUseStateForUnknown(mods []planmodifier.String) bool {
	return hasModifier(mods, stringplanmodifier.UseStateForUnknown())
}

func hasUnsetKeepsPriorState(mods []planmodifier.String) bool {
	return hasModifier(mods, UnsetKeepsPriorState())
}

func TestConnectionResources_KeepPriorStateForUnset(t *testing.T) {
	resources := connectionResourcesForTest(t)
	if len(resources) != 13 {
		t.Fatalf("expected 13 connection resources, found %d", len(resources))
	}

	for name, r := range resources {
		t.Run(name, func(t *testing.T) {
			var sr resource.SchemaResponse
			r.Schema(context.Background(), resource.SchemaRequest{}, &sr)
			if sr.Diagnostics.HasError() {
				t.Fatalf("schema diagnostics: %v", sr.Diagnostics)
			}

			covered := 0
			for attrName, attr := range sr.Schema.Attributes {
				switch a := attr.(type) {
				case schema.StringAttribute:
					switch {
					case a.Optional && a.Computed && !a.Sensitive:
						if !hasUnsetKeepsPriorState(a.PlanModifiers) {
							t.Errorf("%s: optional+computed attribute has no UnsetKeepsPriorState", attrName)
						}
						covered++
					case a.Optional && a.Computed && a.Sensitive:
						if len(a.PlanModifiers) != 0 {
							t.Errorf("%s: sensitive attribute must not plan its prior value", attrName)
						}
					case attrName == "id":
						if !hasUseStateForUnknown(a.PlanModifiers) {
							t.Errorf("id has no UseStateForUnknown")
						}
					case attrName == "msg" || attrName == "error_code":
						if len(a.PlanModifiers) != 0 {
							t.Errorf("%s changes on every apply and must stay unknown in plans", attrName)
						}
					case !a.Computed:
						if len(a.PlanModifiers) != 0 {
							t.Errorf("%s: non-computed attribute got a plan modifier", attrName)
						}
					}
				case schema.Int64Attribute:
					if attrName == "connection_key" && len(a.PlanModifiers) != 1 {
						t.Errorf("connection_key has no UseStateForUnknown")
					}
				}
			}
			if covered == 0 {
				t.Errorf("no optional+computed attributes covered")
			}
		})
	}
}

func TestWithPriorStateForUnset_KeepsExistingModifiers(t *testing.T) {
	custom := stringplanmodifier.RequiresReplace()
	attrs := map[string]schema.Attribute{
		"has_modifier": schema.StringAttribute{Optional: true, Computed: true, PlanModifiers: []planmodifier.String{custom}},
		"plain":        schema.StringAttribute{Optional: true},
		"computed":     schema.StringAttribute{Computed: true},
		"secret":       schema.StringAttribute{Optional: true, Computed: true, Sensitive: true},
	}
	got := WithPriorStateForUnset(attrs)

	if mods := got["has_modifier"].(schema.StringAttribute).PlanModifiers; len(mods) != 1 || hasUnsetKeepsPriorState(mods) {
		t.Errorf("existing plan modifiers must be left untouched, got %v", mods)
	}
	for _, n := range []string{"plain", "computed", "secret"} {
		if mods := got[n].(schema.StringAttribute).PlanModifiers; len(mods) != 0 {
			t.Errorf("%s: unexpected plan modifiers %v", n, mods)
		}
	}
}

func TestUnsetKeepsPriorState(t *testing.T) {
	objType := tftypes.Object{AttributeTypes: map[string]tftypes.Type{"a": tftypes.String}}
	existing := tftypes.NewValue(objType, map[string]tftypes.Value{"a": tftypes.NewValue(tftypes.String, "x")})
	absent := tftypes.NewValue(objType, nil)

	cases := map[string]struct {
		stateRaw tftypes.Value
		state    types.String
		config   types.String
		plan     types.String
		want     types.String
	}{
		"create keeps unknown":            {absent, types.StringNull(), types.StringNull(), types.StringUnknown(), types.StringUnknown()},
		"unset keeps prior value":         {existing, types.StringValue(`{"k":1}`), types.StringNull(), types.StringUnknown(), types.StringValue(`{"k":1}`)},
		"unset keeps prior null":          {existing, types.StringNull(), types.StringNull(), types.StringUnknown(), types.StringNull()},
		"configured value is not touched": {existing, types.StringValue("old"), types.StringValue("new"), types.StringValue("new"), types.StringValue("new")},
		"unknown config stays unknown":    {existing, types.StringValue("old"), types.StringUnknown(), types.StringUnknown(), types.StringUnknown()},
	}
	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			req := planmodifier.StringRequest{
				State:       tfsdk.State{Raw: tc.stateRaw},
				Plan:        tfsdk.Plan{Raw: existing},
				StateValue:  tc.state,
				ConfigValue: tc.config,
				PlanValue:   tc.plan,
			}
			resp := &planmodifier.StringResponse{PlanValue: tc.plan}
			UnsetKeepsPriorState().PlanModifyString(context.Background(), req, resp)
			if !resp.PlanValue.Equal(tc.want) {
				t.Errorf("plan value = %s, want %s", resp.PlanValue, tc.want)
			}
		})
	}
}
