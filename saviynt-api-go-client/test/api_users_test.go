// Copyright (c) 2025 Saviynt Inc.
// SPDX-License-Identifier: MPL-2.0

/*
Saviynt Users API

Testing UsersAPIService

*/

package test

import (
	"context"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_users_UsersAPIService(t *testing.T) {
	apiClient, _, skipTests, skipMsg, err := client()
	require.Nil(t, err)

	t.Run("Test_UsersAPIService_GetUser", func(t *testing.T) {
		if skipTests && strings.TrimSpace(skipMsg) != "" {
			t.Skip(skipMsg)
		} else if skipTests {
			t.Skip(MsgSkipTest)
		}

		resp, httpRes, err := apiClient.Users.
			GetUser(context.Background()).
			Execute()

		require.Nil(t, err)
		require.NotNil(t, httpRes)
		assert.Equal(t, 200, httpRes.StatusCode)
		require.NotNil(t, resp)
		assert.Greater(t, len(resp.Userdetails), 0)
	})
}
