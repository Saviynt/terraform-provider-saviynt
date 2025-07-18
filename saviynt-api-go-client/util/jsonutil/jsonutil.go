// Copyright (c) 2025 Saviynt Inc.
// SPDX-License-Identifier: MPL-2.0

package jsonutil

import (
	"encoding/json"
	"reflect"
	"sort"
	"strings"
)

type Fields []Field

func (fs Fields) JSONTagNames(sortAsc bool) []string {
	var names []string
	for _, f := range fs {
		if t, ok := f.Tags["json"]; ok {
			if t.Name != "" {
				names = append(names, t.Name)
			}
		}
	}
	if sortAsc {
		sort.Strings(names)
	}
	return names
}

type Field struct {
	Name string
	Tags map[string]Tag
}

type Tag struct {
	Name string
}

func ParseFields(s any) Fields {
	fields := Fields{}
	t := reflect.TypeOf(s)

	for i := range t.NumField() {
		field := t.Field(i)
		jsonTag := strings.TrimSpace(field.Tag.Get("json"))

		f := Field{
			Name: field.Name,
			Tags: map[string]Tag{},
		}

		if jsonTag == "" {
			f.Tags["json"] = Tag{Name: field.Name}
		} else {
			if jsonTag = parseJSONTag(jsonTag); jsonTag == "" {
				f.Tags["json"] = Tag{Name: field.Name}
			} else if jsonTag == "-," {
				f.Tags["json"] = Tag{Name: "-"}
			} else if jsonTag != "-" {
				f.Tags["json"] = Tag{Name: jsonTag}
			}
		}
		fields = append(fields, f)
	}
	return fields
}

func parseJSONTag(tag string) string {
	if commaIdx := len(tag); commaIdx > 0 {
		if idx := strings.Index(tag, ","); idx != -1 {
			return tag[:idx]
		}
	}
	return tag
}

func JSONFilterKeys(b []byte, fieldNamesIncl []string, fieldNamesValuesIncl map[string]any) ([]byte, error) {
	msa := map[string]any{}
	err := json.Unmarshal(b, &msa)
	if err != nil {
		return []byte{}, err
	}
	if len(msa) == 0 {
		return b, nil
	}
	incl := map[string]int{}
	for _, f := range fieldNamesIncl {
		incl[f]++
	}
	out := map[string]any{}
	for k, v := range msa {
		if _, ok := incl[k]; ok {
			out[k] = v
		}
	}
	for k, v := range fieldNamesValuesIncl {
		out[k] = v
	}
	return json.Marshal(out)
}
