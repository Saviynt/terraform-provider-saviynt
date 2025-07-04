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

/*
Saviynt mTLS Authentication API

Testing MTLSAuthenticationAPIService

*/

package test

import (
	"context"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_mtlsauthentication_MTLSAuthenticationAPIService(t *testing.T) {
	apiClient, _, skipTests, skipMsg, err := client()
	require.Nil(t, err)

	ctx := context.Background()

	// Get Initial KeyStore State

	keyStoreAliasesInitial := map[string]int{}
	t.Run("Test_MTLSAuthenticationAPIService_GetKeyStoreCertificateDetails_Before_Upload", func(t *testing.T) {
		if skipTests && strings.TrimSpace(skipMsg) != "" {
			t.Skip(skipMsg)
		} else if skipTests {
			t.Skip(MsgSkipTest)
		}

		resp, httpRes, err := apiClient.MTLSAuthentication.
			GetKeyStoreCertificateDetails(ctx).
			Execute()

		require.Nil(t, err)
		require.NotNil(t, httpRes)
		assert.Equal(t, 200, httpRes.StatusCode)
		require.NotNil(t, resp)

		for _, k := range resp.CertificateDetails {
			if k.Alias != nil && *k.Alias != "" {
				keyStoreAliasesInitial[*k.Alias]++
			}
		}
	})

	t.Run("Test_MTLSAuthenticationAPIService_UploadKeyStore", func(t *testing.T) {
		if skipTests && strings.TrimSpace(skipMsg) != "" {
			t.Skip(skipMsg)
		} else if skipTests {
			t.Skip(MsgSkipTest)
		}

		filename := "testdata/mtlsauthentication_pki_key.p12"
		filenamePassword := "12345" // #nosec G101

		f, err := os.Open(filename)
		require.Nil(t, err)
		defer f.Close()

		resp, httpRes, err := apiClient.MTLSAuthentication.
			UploadKeyStore(ctx).
			KeyStoreFile(f).
			KeyStorePassword(filenamePassword).
			Execute()

		require.Nil(t, err)
		require.NotNil(t, httpRes)
		assert.Equal(t, 200, httpRes.StatusCode)
		require.NotNil(t, resp)
	})

	keyStoreAlias := ""
	keyStoreAliasesUpdated := map[string]int{}
	keyStoreAliasNew := []string{}

	// Check KeyStoreCertificates and ensure there is 1 more than previously existing.
	t.Run("Test_MTLSAuthenticationAPIService_GetKeyStoreCertificateDetails_After_Upload", func(t *testing.T) {
		if skipTests && strings.TrimSpace(skipMsg) != "" {
			t.Skip(skipMsg)
		} else if skipTests {
			t.Skip(MsgSkipTest)
		}

		resp, httpRes, err := apiClient.MTLSAuthentication.
			GetKeyStoreCertificateDetails(ctx).
			Execute()

		require.Nil(t, err)
		require.NotNil(t, httpRes)
		assert.Equal(t, 200, httpRes.StatusCode)
		require.NotNil(t, resp)
		assert.Equal(t, len(resp.CertificateDetails), len(keyStoreAliasesInitial)+1)
		keyStoreAlias = *resp.CertificateDetails[0].Alias
		for _, k := range resp.CertificateDetails {
			if k.Alias != nil && *k.Alias != "" {
				keyStoreAliasesUpdated[*k.Alias]++
			}
		}
		for k := range keyStoreAliasesUpdated {
			if _, ok := keyStoreAliasesInitial[k]; !ok {
				keyStoreAliasNew = append(keyStoreAliasNew, k)
			}
		}
		assert.Equal(t, len(keyStoreAliasNew), 1)
		keyStoreAlias = keyStoreAliasNew[0]
	})

	t.Run("Test_MTLSAuthenticationAPIService_DeleteKeyStore", func(t *testing.T) {
		if skipTests && strings.TrimSpace(skipMsg) != "" {
			t.Skip(skipMsg)
		} else if skipTests {
			t.Skip(MsgSkipTest)
		} else if keyStoreAlias == "" {
			t.Skipf(MsgSkipTestPrereqNotSet, "`Test_MTLSAuthenticationAPIService_GetKeyStoreCertificateDetails_After_Upload`")
		}

		assert.NotEqual(t, keyStoreAlias, "")

		httpRes, err := apiClient.MTLSAuthentication.
			DeleteKeyStore(ctx, keyStoreAlias).
			Execute()

		require.Nil(t, err)
		require.NotNil(t, httpRes)
		assert.Equal(t, 200, httpRes.StatusCode)
	})
}
