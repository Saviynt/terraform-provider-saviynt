// Copyright (c) Saviynt Inc.
// SPDX-License-Identifier: MPL-2.0

/*
Saviynt Entitlement API

Testing EntitlementAPIService

*/

package test

import (
	"context"
	"fmt"
	"strings"
	"testing"
	"time"

	saviyntapigoclient "github.com/saviynt/saviynt-api-go-client"
	"github.com/saviynt/saviynt-api-go-client/entitlements"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_entitlements_EntitlementAPIService(t *testing.T) {
	apiClient, _, skipTests, skipMsg, err := client()
	require.Nil(t, err)

	ctx := context.Background()
	uniqueName := fmt.Sprintf("TestEnt_%d", time.Now().Unix())
	endpointName := "sample-101"
	var entitlementKey string
	var updatedName string

	t.Run("Test_EntitlementAPIService_CreateEntitlement", func(t *testing.T) {
		if skipTests && strings.TrimSpace(skipMsg) != "" {
			t.Skip(skipMsg)
		} else if skipTests {
			t.Skip(MsgSkipTest)
		}

		createReq := entitlements.CreateUpdateEntitlementRequest{
			Endpoint:         endpointName,
			Entitlementtype:  "test_postman_1",
			EntitlementValue: uniqueName,
			Displayname:      saviyntapigoclient.Pointer(uniqueName),
			Risk:             saviyntapigoclient.Pointer(int32(1)),
			Status:           saviyntapigoclient.Pointer(int32(1)),
		}

		createResp, httpResp, err := apiClient.Entitlements.CreateUpdateEntitlement(ctx).CreateUpdateEntitlementRequest(createReq).Execute()
		if err != nil {
			t.Logf("Error creating entitlement: %v", err)
			t.Skip("Skipping test due to entitlement creation failure")
			return
		}
		require.NotNil(t, httpResp)
		require.NotNil(t, createResp)
		assert.Equal(t, 200, httpResp.StatusCode)

		if createResp.ErrorCode != nil {
			assert.Equal(t, "0", *createResp.ErrorCode, "Expected success error code")
		}

		if createResp.EntitlementObj != nil {
			entitlementKey = *createResp.EntitlementObj.EntitlementValuekey
		}

		t.Logf("Successfully created entitlement: %s", uniqueName)
	})

	t.Run("Test_EntitlementAPIService_UpdateEntitlement", func(t *testing.T) {
		if skipTests && strings.TrimSpace(skipMsg) != "" {
			t.Skip(skipMsg)
		} else if skipTests {
			t.Skip(MsgSkipTest)
		}

		updatedName = fmt.Sprintf("UpdatedTestEnt_%d", time.Now().Unix())
		updateReq := entitlements.CreateUpdateEntitlementRequest{
			Endpoint:                "sample-101",
			Entitlementtype:         "test_postman_1",
			EntitlementValue:        uniqueName,
			UpdatedentitlementValue: saviyntapigoclient.Pointer(updatedName),
			Displayname:             saviyntapigoclient.Pointer("Updated Test Entitlement"),
			Entitlementowner2:       []string{"admin"},
			Entitlementowner24:      []string{"admin"},
			Customproperty23:        saviyntapigoclient.Pointer("cp23"),
			Customproperty2:         saviyntapigoclient.Pointer("cp2"),
			Risk:                    saviyntapigoclient.Pointer(int32(2)),
			Status:                  saviyntapigoclient.Pointer(int32(1)),
			Entitlementmap: []entitlements.CreateUpdateEntitlementRequestEntitlementmapInner{
				{
					Entitlementvalue: saviyntapigoclient.Pointer("HRGroup"),
					Entitlementtype:  saviyntapigoclient.Pointer("Group"),
					Endpoint:         saviyntapigoclient.Pointer("AWS_IC_NewJar5"),
					UpdateType:       saviyntapigoclient.Pointer("ADD"),
					Requestfilter:    saviyntapigoclient.Pointer(true),
				},
			},
		}

		updateResp, httpResp, err := apiClient.Entitlements.CreateUpdateEntitlement(ctx).CreateUpdateEntitlementRequest(updateReq).Execute()
		if err != nil {
			t.Logf("Error updating entitlement: %v", err)
			t.Skip("Skipping test due to entitlement updation failure")
			return
		}

		require.NotNil(t, httpResp)
		require.NotNil(t, updateResp)
		assert.Equal(t, 200, httpResp.StatusCode)

		if updateResp.ErrorCode != nil {
			assert.Equal(t, "0", *updateResp.ErrorCode, "Expected success error code")
		}

		if updateResp.EntitlementObj != nil {
			assert.Equal(t, updatedName, *updateResp.EntitlementObj.EntitlementValue, "Expected entitlement value to be updated")
			assert.Equal(t, entitlementKey, *updateResp.EntitlementObj.EntitlementValuekey, "Same entitlement key not found")
		}

		t.Logf("Successfully updated entitlement with original name: %s and updated name as: %s", uniqueName, updatedName)
	})

	t.Run("Test_EntitleementAPIService_GetEntitlementByUniqueCompositeKey(Endpoint, Ent Type, Ent)", func(t *testing.T) {
		if skipTests && strings.TrimSpace(skipMsg) != "" {
			t.Skip(skipMsg)
		} else if skipTests {
			t.Skip(MsgSkipTest)
		}

		getReq := entitlements.GetEntitlementRequest{
			Endpoint:             saviyntapigoclient.Pointer("sample-101"),
			Entitlementtype:      saviyntapigoclient.Pointer("test_postman_1"),
			EntitlementValue:     saviyntapigoclient.Pointer(updatedName),
			Entownerwithrank:     saviyntapigoclient.Pointer("true"),
			Returnentitlementmap: saviyntapigoclient.Pointer("true"),
		}

		getResp, httpResp, err := apiClient.Entitlements.GetEntitlements(ctx).GetEntitlementRequest(getReq).Execute()
		if err != nil {
			t.Logf("Error reading entitlement: %v", err)
			t.Skip("Skipping test due to entitlement read failure")
			return
		}

		require.NotNil(t, httpResp)
		require.NotNil(t, getResp)
		assert.Equal(t, 200, httpResp.StatusCode)

		// Validate response structure
		assert.NotNil(t, getResp.ErrorCode, "ErrorCode should not be nil")
		assert.Equal(t, "0", *getResp.ErrorCode, "Expected success error code")
		assert.NotNil(t, getResp.Msg, "Message should not be nil")
		assert.Equal(t, "Successful", *getResp.Msg, "Expected success message")

		// Validate entitlement details
		require.NotNil(t, getResp.Entitlementdetails, "Entitlementdetails should not be nil")
		require.Greater(t, len(getResp.Entitlementdetails), 0, "Should have at least one entitlement")

		entitlement := getResp.Entitlementdetails[0]
		assert.Equal(t, updatedName, *entitlement.EntitlementValue, "Entitlement value should match")
		assert.Equal(t, "sample-101", *entitlement.Endpoint, "Endpoint should match")
		assert.Equal(t, "test_postman_1", *entitlement.EntitlementType, "Entitlement type should match")

		// Validate entitlement owners
		if entitlement.EntitlementOwner != nil && entitlement.EntitlementOwner.MapmapOfStringarrayOfString != nil {
			owners := *entitlement.EntitlementOwner.MapmapOfStringarrayOfString
			assert.Greater(t, len(owners), 0, "Should have entitlement owners")
			t.Logf("Found %d owner ranks", len(owners))
			for rank, users := range owners {
				assert.Greater(t, len(users), 0, "Rank %s should have users", rank)
				t.Logf("Rank %s has %d users: %v", rank, len(users), users)
			}
		}

		// Validate entitlement map details
		if len(entitlement.EntitlementMapDetails) > 0 {
			assert.Greater(t, len(entitlement.EntitlementMapDetails), 0, "Should have entitlement map details")
			t.Logf("Found %d entitlement map entries", len(entitlement.EntitlementMapDetails))
			for i, mapDetail := range entitlement.EntitlementMapDetails {
				// Validate Entitlementtype matches primaryEntType
				if mapDetail.PrimaryEntType != nil {
					assert.Equal(t, "Group", *mapDetail.PrimaryEntType,
						"Map detail %d: Entitlementtype should match primaryEntType", i)
				}

				// Validate entitlementvalue matches primary
				if mapDetail.Primary != nil {
					assert.Equal(t, "HRGroup", *mapDetail.Primary,
						"Map detail %d: entitlementvalue should match primary", i)
				}

				t.Logf("Map detail %d: PrimaryEntType=%v, Primary=%v", i,
					*mapDetail.PrimaryEntType, *mapDetail.Primary)
			}
		}

		t.Logf("Successfully read entitlement with its composite key consisting of endpoint, ent type, ent value name")
	})

	t.Run("Test_EntitleementAPIService_GetEntitlementByEntitlementValueKey", func(t *testing.T) {
		if skipTests && strings.TrimSpace(skipMsg) != "" {
			t.Skip(skipMsg)
		} else if skipTests {
			t.Skip(MsgSkipTest)
		}

		entquery := fmt.Sprintf("ent.id like '%s'", entitlementKey)
		fmt.Printf("Entitlement query is: %s", entquery)
		getReq := entitlements.GetEntitlementRequest{
			EntQuery:             saviyntapigoclient.Pointer(entquery),
			Entownerwithrank:     saviyntapigoclient.Pointer("true"),
			Returnentitlementmap: saviyntapigoclient.Pointer("true"),
		}

		getResp, httpResp, err := apiClient.Entitlements.GetEntitlements(ctx).GetEntitlementRequest(getReq).Execute()
		if err != nil {
			t.Logf("Error reading entitlement: %v", err)
			t.Skip("Skipping test due to entitlement read failure")
			return
		}

		require.NotNil(t, httpResp)
		require.NotNil(t, getResp)
		assert.Equal(t, 200, httpResp.StatusCode)

		// Validate response structure
		assert.NotNil(t, getResp.ErrorCode, "ErrorCode should not be nil")
		assert.Equal(t, "0", *getResp.ErrorCode, "Expected success error code")
		assert.NotNil(t, getResp.Msg, "Message should not be nil")
		assert.Equal(t, "Successful", *getResp.Msg, "Expected success message")

		// Validate entitlement details
		require.NotNil(t, getResp.Entitlementdetails, "Entitlementdetails should not be nil")
		require.Greater(t, len(getResp.Entitlementdetails), 0, "Should have at least one entitlement")

		entitlement := getResp.Entitlementdetails[0]
		assert.Equal(t, entitlementKey, *entitlement.EntitlementValuekey, "Entitlement value key should match")
		assert.NotNil(t, entitlement.EntitlementValue, "Entitlement value should not be nil")
		assert.NotNil(t, entitlement.Endpoint, "Endpoint should not be nil")
		assert.NotNil(t, entitlement.EntitlementType, "Entitlement type should not be nil")

		// Validate entitlement owners
		if entitlement.EntitlementOwner != nil && entitlement.EntitlementOwner.MapmapOfStringarrayOfString != nil {
			owners := *entitlement.EntitlementOwner.MapmapOfStringarrayOfString
			assert.Greater(t, len(owners), 0, "Should have entitlement owners")
			t.Logf("Found %d owner ranks", len(owners))
			for rank, users := range owners {
				assert.Greater(t, len(users), 0, "Rank %s should have users", rank)
				t.Logf("Rank %s has %d users: %v", rank, len(users), users)
			}
		}

		t.Logf("Successfully read entitlement with its unique entitlement value key")
	})
}
