// Copyright (c) Saviynt Inc.
// SPDX-License-Identifier: MPL-2.0

package delegatedadministrationutil

import (
	"strings"
	"time"

	"github.com/saviynt/saviynt-api-go-client/delegatedadministration"
)

type DelegateUsers []delegatedadministration.DelegateUser

func (dels DelegateUsers) Usernames(exclUsernames []string) []string {
	var users []string
	exclMap := map[string]int{}
	for _, exclUsername := range exclUsernames {
		exclMap[exclUsername]++
	}
	for _, del := range dels {
		if _, ok := exclMap[del.Username]; ok {
			continue
		}
		users = append(users, del.Username)
	}
	return users
}

type Delegates []delegatedadministration.Delegate

func (dels Delegates) Keys() []string {
	var keys []string
	for _, del := range dels {
		keys = append(keys, del.Delegatekey)
	}
	return keys
}

type TimeRange struct {
	InputLayout  string
	OutputLayout string
	Min          *time.Time
	Max          *time.Time
}

func (tr *TimeRange) AddTime(t time.Time) {
	if tr.Min == nil || t.Before(*tr.Min) {
		tr.Min = &t
	}
	if tr.Max == nil || t.After(*tr.Max) {
		tr.Max = &t
	}
}

func (tr *TimeRange) MaxOrDefault(def time.Time) time.Time {
	if tr.Max == nil {
		return def
	} else if tr.Max.After(def) {
		return *tr.Max
	} else {
		return def
	}
}

func (tr *TimeRange) MinOrDefault(def time.Time) time.Time {
	if tr.Min == nil {
		return def
	} else if tr.Min.Before(def) {
		return *tr.Min
	} else {
		return def
	}
}

func (tr *TimeRange) AddParseTime(layout string, value string) error {
	if layout == "" {
		layout = tr.InputLayout
	}
	if t, err := time.Parse(layout, value); err != nil {
		return err
	} else {
		tr.AddTime(t)
		return nil
	}
}

func (dels Delegates) MaxTimeRange(dateLayout string) (TimeRange, error) {
	tr := TimeRange{
		InputLayout: dateLayout,
	}

	for _, del := range dels {
		sta := strings.TrimSpace(del.Startdate)
		if sta != "" {
			if err := tr.AddParseTime("", sta); err != nil {
				return tr, err
			}
		}
		end := strings.TrimSpace(del.Enddate)
		if end != "" {
			if err := tr.AddParseTime("", end); err != nil {
				return tr, err
			}
		}
	}

	return tr, nil
}
