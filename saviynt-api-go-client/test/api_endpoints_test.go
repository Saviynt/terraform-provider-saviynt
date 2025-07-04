/*
 * Copyright (c) 2025 Saviynt Inc.
 * All Rights Reserved.
 *
 * This software is the confidential and proprietary information of
 * Saviynt Inc. ("Confidential Information"). You shall not disclose,
 * use, or distribute such Confidential Information except in accordance
 * with the terms of the license agreement you entered into with Saviynt.
 *
 * SAVIYNT MAKES NO REPRESENTATIONS OR WARRANTIES ABOUT THE SUITABILITY OF
 * THE SOFTWARE, EITHER EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO
 * THE IMPLIED WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR
 * PURPOSE, OR NON-INFRINGEMENT.
 */

package test

import (
	"context"
	"strings"
	"testing"
	"time"

	"github.com/saviynt/saviynt-api-go-client/endpoints"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Test_endpoints_EndpointsAPIService tests the Endpoints APIs.
func Test_endpoints_EndpointsAPIService(t *testing.T) {
	apiClient, _, skipTests, skipMsg, err := client()
	require.Nil(t, err)

	ctx := context.Background()

	testSecuritySystemName := ""
	testCreateEndpointName := "saviynt-api-go-client Test on " + time.Now().UTC().Format(time.RFC3339)
	testCreateEndpointSuccess := false
	testUpdateEndpointSuccess := false
	updatedDesc := ""

	t.Run("Test_EndpointsAPIService_GetEndpoints", func(t *testing.T) {
		if skipTests && strings.TrimSpace(skipMsg) != "" {
			t.Skip(skipMsg)
		} else if skipTests {
			t.Skip(MsgSkipTest)
		}

		resp, httpRes, err := apiClient.Endpoints.
			GetEndpoints(ctx).
			GetEndpointsRequest(endpoints.GetEndpointsRequest{}).
			Execute()

		require.Nil(t, err)
		require.NotNil(t, httpRes)
		assert.Equal(t, 200, httpRes.StatusCode)
		require.NotNil(t, resp)
		assert.Equal(t, "0", *resp.ErrorCode)
		assert.Equal(t, "Success", *resp.Message)
	})

	t.Run("Test_EndpointsAPIService_CreateEndpoint", func(t *testing.T) {
		if skipTests && strings.TrimSpace(skipMsg) != "" {
			t.Skip(skipMsg)
		} else if skipTests {
			t.Skip(MsgSkipTest)
		} else if testSecuritySystemName == "" {
			t.Skip(skipTestMessage("prerequisite not found, security system name not set"))
		}

		req := endpoints.CreateEndpointRequest{
			Endpointname:   testCreateEndpointName,
			DisplayName:    testCreateEndpointName,
			Securitysystem: testSecuritySystemName,
			Description:    &testCreateEndpointName,
		}

		resp, httpRes, err := apiClient.Endpoints.
			CreateEndpoint(ctx).
			CreateEndpointRequest(req).
			Execute()

		require.Nil(t, err)
		require.NotNil(t, httpRes)
		assert.Equal(t, 200, httpRes.StatusCode)
		require.NotNil(t, resp)
		assert.Equal(t, "0", *resp.ErrorCode)
	})

	t.Run("Test_EndpointsAPIService_GetEndpoints_CreatedEndpoint", func(t *testing.T) {
		if skipTests && strings.TrimSpace(skipMsg) != "" {
			t.Skip(skipMsg)
		} else if skipTests {
			t.Skip(MsgSkipTest)
		} else if !testCreateEndpointSuccess {
			t.Skip(skipTestMessage("prerequisite not found, no created endpoint"))
		}

		resp, httpRes, err := apiClient.Endpoints.
			GetEndpoints(ctx).
			GetEndpointsRequest(endpoints.GetEndpointsRequest{
				Endpointname: &testCreateEndpointName,
			}).
			Execute()

		require.Nil(t, err)
		require.NotNil(t, httpRes)
		assert.Equal(t, 200, httpRes.StatusCode)
		require.NotNil(t, resp)
		assert.Equal(t, "0", *resp.ErrorCode)
	})

	t.Run("Test_EndpointsAPIService_UpdateEndpoint", func(t *testing.T) {
		if skipTests && strings.TrimSpace(skipMsg) != "" {
			t.Skip(skipMsg)
		} else if skipTests {
			t.Skip(MsgSkipTest)
		} else if !testCreateEndpointSuccess {
			t.Skip(skipTestMessage("prerequisite not found, no created endpoint"))
		}

		updatedDesc = "Updated at " + time.Now().UTC().Format(time.RFC3339)

		req := endpoints.UpdateEndpointRequest{
			Endpointname: testCreateEndpointName,
			Description:  &updatedDesc,
		}

		resp, httpRes, err := apiClient.Endpoints.
			UpdateEndpoint(ctx).
			UpdateEndpointRequest(req).
			Execute()

		require.Nil(t, err)
		require.NotNil(t, httpRes)
		assert.Equal(t, 200, httpRes.StatusCode)
		require.NotNil(t, resp)
		assert.Equal(t, "0", *resp.ErrorCode)
	})

	t.Run("Test_EndpointsAPIService_GetEndpoints_UpdatedEndpoint", func(t *testing.T) {
		if skipTests && strings.TrimSpace(skipMsg) != "" {
			t.Skip(skipMsg)
		} else if skipTests {
			t.Skip(MsgSkipTest)
		} else if !testUpdateEndpointSuccess {
			t.Skip(skipTestMessage("prerequisite not found, no updated endpoint"))
		}

		resp, httpRes, err := apiClient.Endpoints.
			GetEndpoints(ctx).
			GetEndpointsRequest(endpoints.GetEndpointsRequest{
				Endpointname: &testCreateEndpointName,
			}).
			Execute()

		require.Nil(t, err)
		require.NotNil(t, httpRes)
		assert.Equal(t, 200, httpRes.StatusCode)
		require.NotNil(t, resp)
		assert.Equal(t, "0", *resp.ErrorCode)
	})
}
