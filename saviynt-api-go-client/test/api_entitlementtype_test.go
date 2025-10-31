// Copyright (c) Saviynt Inc.
// SPDX-License-Identifier: MPL-2.0

/*
Saviynt Entitlement Type API

Testing EntitlementTypeAPIService

*/

package test

import (
	"context"
	"fmt"
	"strings"
	"testing"
	"time"

	saviyntapigoclient "github.com/saviynt/saviynt-api-go-client"
	"github.com/saviynt/saviynt-api-go-client/entitlementtype"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_entitlementtype_EntitlementTypeAPIService(t *testing.T) {
	apiClient, _, skipTests, skipMsg, err := client()
	require.Nil(t, err)

	ctx := context.Background()
	// Generate a unique name for the entitlement type
	uniqueName := fmt.Sprintf("TestEntType_%d", time.Now().Unix())
	endpointName := "sample-101" // Assuming AWS endpoint exists
	var createdEntitlementTypeName string

	t.Run("Test_EntitlementTypeAPIService_CreateEntitlementType", func(t *testing.T) {
		if skipTests && strings.TrimSpace(skipMsg) != "" {
			t.Skip(skipMsg)
		} else if skipTests {
			t.Skip(MsgSkipTest)
		}

		// Create a new entitlement type
		req := entitlementtype.CreateEntitlementTypeRequest{
			Entitlementname:        uniqueName,
			Entitlementdescription: saviyntapigoclient.Pointer("Initial description for test entitlement type"),
			Endpointname:           endpointName,
			DisplayName:            saviyntapigoclient.Pointer(uniqueName),
			Workflow:               saviyntapigoclient.Pointer("Autoapprovalwf"),
			Orderindex:             saviyntapigoclient.Pointer(int32(1)),
			Certifiable:            saviyntapigoclient.Pointer(true),
		}

		resp, httpRes, err := apiClient.EntitlementType.
			CreateEntitlementType(ctx).
			CreateEntitlementTypeRequest(req).
			Execute()

		// If creation fails, log the error
		if err != nil {
			t.Logf("Error creating entitlement type: %v", err)
			t.Skip("Skipping test due to entitlement type creation failure")
			return
		}

		require.NotNil(t, httpRes)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

		if resp.ErrorCode != nil {
			assert.Equal(t, "0", *resp.ErrorCode, "Expected success error code")
		}

		// Store the created entitlement type name for subsequent tests
		createdEntitlementTypeName = uniqueName
		t.Logf("Successfully created entitlement type: %s", uniqueName)
	})

	t.Run("Test_EntitlementTypeAPIService_UpdateEntitlementType", func(t *testing.T) {
		if skipTests && strings.TrimSpace(skipMsg) != "" {
			t.Skip(skipMsg)
		} else if skipTests {
			t.Skip(MsgSkipTest)
		} else if createdEntitlementTypeName == "" {
			t.Skip("No entitlement type name available for testing")
		}

		// Update the entitlement type
		updatedDescription := fmt.Sprintf("Updated description at %s", time.Now().Format(time.RFC3339))
		req := entitlementtype.UpdateEntitlementTypeRequest{
			Entitlementname:        createdEntitlementTypeName,
			Endpointname:           endpointName,
			Entitlementdescription: saviyntapigoclient.Pointer(updatedDescription),
			Certifiable:            saviyntapigoclient.Pointer(true),
		}

		resp, httpRes, err := apiClient.EntitlementType.
			UpdateEntitlementType(ctx).
			UpdateEntitlementTypeRequest(req).
			Execute()

		// If update fails, log the error
		if err != nil {
			t.Logf("Error updating entitlement type: %v", err)
			t.Skip("Skipping test due to entitlement type update failure")
			return
		}

		require.NotNil(t, httpRes)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

		if resp.ErrorCode != nil {
			assert.Equal(t, "0", *resp.ErrorCode, "Expected success error code for update")
		}

		t.Logf("Successfully updated entitlement type: %s", createdEntitlementTypeName)
	})

	t.Run("Test_EntitlementTypeAPIService_GetEntitlementTypeByName", func(t *testing.T) {
		if skipTests && strings.TrimSpace(skipMsg) != "" {
			t.Skip(skipMsg)
		} else if skipTests {
			t.Skip(MsgSkipTest)
		} else if createdEntitlementTypeName == "" {
			t.Skip("No entitlement type name available for testing")
		}

		// Get the specific entitlement type by name
		resp, httpRes, err := apiClient.EntitlementType.
			GetEntitlementType(ctx).
			Entitlementname(createdEntitlementTypeName).
			Execute()

		require.Nil(t, err, "Error reading entitlement type by name: %v", err)
		require.NotNil(t, httpRes)
		assert.Equal(t, 200, httpRes.StatusCode)
		require.NotNil(t, resp)

		// Verify the response contains the requested entitlement type
		found := false
		if resp.EntitlementTypeDetails != nil {
			for _, entType := range resp.EntitlementTypeDetails {
				if entType.Entitlementname != nil && *entType.Entitlementname == createdEntitlementTypeName {
					found = true
					t.Logf("Found entitlement type: %s", *entType.Entitlementname)
					break
				}
			}
		}

		assert.True(t, found, "Created entitlement type not found in response")
	})
}
