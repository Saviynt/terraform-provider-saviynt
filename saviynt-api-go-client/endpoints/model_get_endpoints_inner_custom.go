// Copyright (c) 2025 Saviynt Inc.
// SPDX-License-Identifier: MPL-2.0

// This file is NOT generated. It extends GetEndpoints200ResponseEndpointsInner
// with a custom UnmarshalJSON that handles both JSON key formats for
// custom properties 1–30:
//   - Spaced:    "Custom Property 1"   (some server configurations)
//   - Lowercase: "customproperty1"     (other server configurations)
//
// The generated struct tags (model_get_endpoints_200_response_endpoints_inner.go)
// only cover the spaced format. Without this file, CP1–CP30 silently deserialise
// to nil when the server returns lowercase keys, causing permanent state data
// loss in the Terraform provider (manifests as "Provider produced inconsistent
// result after apply" errors).
//
// Ticket: TER-283
//
// WARNING: If the generated struct (model_get_endpoints_200_response_endpoints_inner.go)
// is regenerated via openapi-generator, this file must be preserved — it will
// NOT be overwritten. If the generated struct fields are renamed, this file
// must be updated manually (compilation will fail, making the breakage visible).
package endpoints

import (
	"encoding/json"
	"fmt"
)

// UnmarshalJSON implements a two-pass JSON decode for GetEndpoints200ResponseEndpointsInner.
//
// Pass 1 (via alias type): runs the default generated field matching, which
// handles all fields including CP1–CP30 with spaced keys ("Custom Property N")
// and CP31–CP45 with lowercase keys ("custompropertyN").
//
// Pass 2 (fallback): for CP1–CP30 fields that are still nil after pass 1
// (meaning the server returned lowercase keys "customproperty1"…"customproperty30"
// instead of the spaced format), re-checks the raw JSON map and fills them in.
func (o *GetEndpoints200ResponseEndpointsInner) UnmarshalJSON(data []byte) error {
	// Use an alias type to invoke the default generated unmarshal logic
	// without triggering infinite recursion on this method.
	type Alias GetEndpoints200ResponseEndpointsInner
	aux := &Alias{}
	if err := json.Unmarshal(data, aux); err != nil {
		return err
	}
	*o = GetEndpoints200ResponseEndpointsInner(*aux)

	// For CP1–CP30, the generated tags use the spaced format ("Custom Property N").
	// If the server returned lowercase keys ("custompropertyN"), those fields will
	// be nil after the default unmarshal above. Parse the raw map and fill them in.
	var raw map[string]json.RawMessage
	if err := json.Unmarshal(data, &raw); err != nil {
		return nil
	}

	fallback := func(current *string, index int) *string {
		if current != nil {
			return current
		}
		key := fmt.Sprintf("customproperty%d", index)
		v, ok := raw[key]
		if !ok {
			return nil
		}
		var s string
		if err := json.Unmarshal(v, &s); err != nil {
			return nil
		}
		return &s
	}

	o.CustomProperty1 = fallback(o.CustomProperty1, 1)
	o.CustomProperty2 = fallback(o.CustomProperty2, 2)
	o.CustomProperty3 = fallback(o.CustomProperty3, 3)
	o.CustomProperty4 = fallback(o.CustomProperty4, 4)
	o.CustomProperty5 = fallback(o.CustomProperty5, 5)
	o.CustomProperty6 = fallback(o.CustomProperty6, 6)
	o.CustomProperty7 = fallback(o.CustomProperty7, 7)
	o.CustomProperty8 = fallback(o.CustomProperty8, 8)
	o.CustomProperty9 = fallback(o.CustomProperty9, 9)
	o.CustomProperty10 = fallback(o.CustomProperty10, 10)
	o.CustomProperty11 = fallback(o.CustomProperty11, 11)
	o.CustomProperty12 = fallback(o.CustomProperty12, 12)
	o.CustomProperty13 = fallback(o.CustomProperty13, 13)
	o.CustomProperty14 = fallback(o.CustomProperty14, 14)
	o.CustomProperty15 = fallback(o.CustomProperty15, 15)
	o.CustomProperty16 = fallback(o.CustomProperty16, 16)
	o.CustomProperty17 = fallback(o.CustomProperty17, 17)
	o.CustomProperty18 = fallback(o.CustomProperty18, 18)
	o.CustomProperty19 = fallback(o.CustomProperty19, 19)
	o.CustomProperty20 = fallback(o.CustomProperty20, 20)
	o.CustomProperty21 = fallback(o.CustomProperty21, 21)
	o.CustomProperty22 = fallback(o.CustomProperty22, 22)
	o.CustomProperty23 = fallback(o.CustomProperty23, 23)
	o.CustomProperty24 = fallback(o.CustomProperty24, 24)
	o.CustomProperty25 = fallback(o.CustomProperty25, 25)
	o.CustomProperty26 = fallback(o.CustomProperty26, 26)
	o.CustomProperty27 = fallback(o.CustomProperty27, 27)
	o.CustomProperty28 = fallback(o.CustomProperty28, 28)
	o.CustomProperty29 = fallback(o.CustomProperty29, 29)
	o.CustomProperty30 = fallback(o.CustomProperty30, 30)

	return nil
}
