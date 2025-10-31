// Copyright (c) 2025 Saviynt Inc.
// SPDX-License-Identifier: MPL-2.0

package filepathutil

import (
	"fmt"
	"os"
)

func ExistsNonZero(filename string) error {
	_, err := os.Stat(filename)
	if err != nil && os.IsNotExist(err) {
		return fmt.Errorf("filename (%s) does not exist", filename)
	} else if err != nil {
		return fmt.Errorf("error checking filename (%s)", filename)
	} else {
		return nil
	}
}
