// Copyright (c) 2025 Saviynt Inc.
// SPDX-License-Identifier: MPL-2.0

package rolesutil

import (
	"fmt"
	"time"
)

func RoleNameUsingTimeStamp(base string) string {
	ts := time.Now().Format("20060102150405")
	return fmt.Sprintf("%s_%s", base, ts)
}
