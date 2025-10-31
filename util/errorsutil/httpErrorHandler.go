// Copyright (c) Saviynt Inc.
// SPDX-License-Identifier: MPL-2.0

package errorsutil

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func HandleHTTPError(httpResp *http.Response, originalErr error, operation string) error {
	if httpResp != nil && httpResp.StatusCode != http.StatusOK {
		log.Printf("[DEBUG] HTTP error for %s operation status: %s\n", operation, httpResp.Status)
		var errorResp map[string]interface{}
		if decodeErr := json.NewDecoder(httpResp.Body).Decode(&errorResp); decodeErr == nil {
			if msg, exists := errorResp["msg"]; exists {
				if errorCode, codeExists := errorResp["errorCode"]; codeExists {
					return fmt.Errorf("%v - ErrorCode: %v, Msg: %v", originalErr, errorCode, msg)
				} else {
					return fmt.Errorf("%v - Msg: %v", originalErr, msg)
				}
			}
		}
	}
	return originalErr
}
