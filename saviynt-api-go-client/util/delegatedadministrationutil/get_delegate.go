// Copyright (c) 2025 Saviynt Inc.
// SPDX-License-Identifier: MPL-2.0

package delegatedadministrationutil

import (
	"context"
	"errors"
	"fmt"
	"math"
	"strconv"

	saviyntapigoclient "github.com/saviynt/saviynt-api-go-client"
	"github.com/saviynt/saviynt-api-go-client/delegatedadministration"
)

func GetDelegate(ctx context.Context, api *delegatedadministration.DelegatedAdministrationAPIService, username, key string) (*delegatedadministration.Delegate, error) {
	if api == nil {
		return nil, errors.New("delegatedadministration API service cannot be nil")
	}
	max := int32(50)
	offset := int32(0)
	totalCount := int32(0)

	for {
		req := delegatedadministration.FetchExistingDelegatesListRequest{
			UserName: username,
			Offset:   saviyntapigoclient.Pointer(offset),
			Max:      saviyntapigoclient.Pointer(max),
		}

		resp, httpRes, err := api.
			FetchExistingDelegatesList(ctx).
			FetchExistingDelegatesListRequest(req).
			Execute()
		if err != nil {
			return nil, err
		} else if httpRes.StatusCode != 200 {
			return nil, fmt.Errorf("http response status code is not successful (%d)", httpRes.StatusCode)
		} else {
			for _, d := range resp.DelegateList {
				if d.Delegatekey == key {
					return &d, nil
				}
			}
			if resp.TotalCount != nil {
				if totalCountInt, err := strconv.Atoi(*resp.TotalCount); err != nil {
					return nil, err
				} else if i32, err := IntConvIntToInt32(totalCountInt); err != nil {
					return nil, err
				} else {
					totalCount = i32
				}
			}
			if offset+max >= totalCount {
				break
			}
			offset += max
		}
	}
	return nil, fmt.Errorf("delegtekey not found (%s)", key)
}

func IntConvIntToInt32(i int) (int32, error) {
	if i < math.MinInt32 || i > math.MaxInt32 {
		return 0, fmt.Errorf("value %d out of int32 range", i)
	} else {
		return int32(i), nil
	}
}
