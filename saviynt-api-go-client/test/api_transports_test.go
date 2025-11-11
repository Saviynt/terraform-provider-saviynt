// Copyright (c) 2025 Saviynt Inc.
// SPDX-License-Identifier: MPL-2.0

/*
Saviynt Transports API

Testing TransportsAPIService
*/

package test

import (
	"context"
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/saviynt/saviynt-api-go-client/transports"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_transports_TransportsAPIService(t *testing.T) {
	apiClient, _, skipTests, skipMsg, err := client()
	require.Nil(t, err)

	ctx := context.Background()
	exportPath := "/saviynt_shared/testexport/transportPackage"
	var exportedFileName string
	var importRequestId string

	t.Run("Test TransportsAPIService ExportTransportPackage", func(t *testing.T) {
		if skipTests && strings.TrimSpace(skipMsg) != "" {
			t.Skip(skipMsg)
		} else if skipTests {
			t.Skip(MsgSkipTest)
		}

		exportReq := transports.ExportTransportPackageRequest{
			Updateuser:       transports.PtrString("admin"),
			Transportowner:   transports.PtrString("true"),
			Transportmembers: transports.PtrString("true"),
			Exportonline:     "false",
			Exportpath:       exportPath,
			Objectstoexport: transports.ExportTransportPackageRequestObjectstoexport{
				SavRoles:      []string{"ROLE_ADMIN"},
				EmailTemplate: []string{"Account Password Expiry Email"},
			},
			Businessjustification: transports.PtrString("API test export"),
		}

		resp, httpRes, err := apiClient.Transports.ExportTransportPackage(ctx).
			ExportTransportPackageRequest(exportReq).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
		assert.NotNil(t, resp.FileName)
		assert.NotNil(t, resp.Msg)

		exportedFileName = *resp.FileName
		fmt.Printf("Export initiated: %s, FileName: %s\n", *resp.Msg, exportedFileName)
	})

	t.Run("Test TransportsAPIService TransportPackageStatus - Export", func(t *testing.T) {
		if skipTests && strings.TrimSpace(skipMsg) != "" {
			t.Skip(skipMsg)
		} else if skipTests {
			t.Skip(MsgSkipTest)
		}

		if exportedFileName == "" {
			t.Skip("No exported file to check status for")
		}

		statusReq := transports.TransportPackageStatusRequest{
			Operation: "export",
			Filename:  exportedFileName,
		}

		resp, httpRes, err := apiClient.Transports.TransportPackageStatus(ctx).
			TransportPackageStatusRequest(statusReq).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
		assert.NotNil(t, resp.Msg)

		fmt.Printf("Export status: %s\n", *resp.Msg)
	})

	t.Run("Test TransportsAPIService ImportTransportPackage", func(t *testing.T) {
		if skipTests && strings.TrimSpace(skipMsg) != "" {
			t.Skip(skipMsg)
		} else if skipTests {
			t.Skip(MsgSkipTest)
		}

		if exportedFileName == "" {
			t.Skip("No exported file to import")
		}

		// Wait for export to complete
		fmt.Println("Waiting 30 seconds for export to complete...")
		time.Sleep(30 * time.Second)

		importReq := transports.ImportTransportPackageRequest{
			Updateuser:            transports.PtrString("admin"),
			Packagetoimport:       fmt.Sprintf("%s/%s", exportPath, exportedFileName),
			Businessjustification: transports.PtrString("API test import"),
		}

		resp, httpRes, err := apiClient.Transports.ImportTransportPackage(ctx).
			ImportTransportPackageRequest(importReq).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
		assert.NotNil(t, resp.Msg)
		assert.NotNil(t, resp.RequestId)

		importRequestId = *resp.RequestId
		fmt.Printf("Import initiated: %s, RequestId: %s\n", *resp.Msg, importRequestId)
	})

	t.Run("Test TransportsAPIService TransportPackageStatus - Import", func(t *testing.T) {
		if skipTests && strings.TrimSpace(skipMsg) != "" {
			t.Skip(skipMsg)
		} else if skipTests {
			t.Skip(MsgSkipTest)
		}

		if exportedFileName == "" || importRequestId == "" {
			t.Skip("No import operation to check status for")
		}

		statusReq := transports.TransportPackageStatusRequest{
			Operation: "import",
			Filename:  exportedFileName,
			Requestid: &importRequestId,
		}

		resp, httpRes, err := apiClient.Transports.TransportPackageStatus(ctx).
			TransportPackageStatusRequest(statusReq).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
		assert.NotNil(t, resp.Msg)

		fmt.Printf("Import status: %s\n", *resp.Msg)
	})
}
