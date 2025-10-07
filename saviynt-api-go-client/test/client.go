// Copyright (c) 2025 Saviynt Inc.
// SPDX-License-Identifier: MPL-2.0

package test

import (
	"context"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"

	saviyntapigoclient "github.com/saviynt/saviynt-api-go-client"
)

const (
	EnvSaviyntTestCredentials = "SAVIYNT_TEST_CREDENTIALS" // #nosec G101
	EnvSaviyntTestData        = "SAVIYNT_TEST_DATA"        // #nosec G101

	MsgSkipTest                         = "skip test"
	MsgSkipTestUnimplemented            = "skip test: unimplemented"
	MsgSkipTestErrorInstantiatingClient = "skip test: error instantiating client (%s)"
	MsgSkipTestCredentialsNotSet        = "skip test: credentials env var not set" // #nosec G101
	MsgSkipTestClientNotConfigured      = "skip test: client not configured"
	MsgSkipTestPrereqNotSet             = "skip test: test pre-req not set from (%s)"
)

func skipTestMessage(msg string) string {
	if msg = strings.TrimSpace(msg); msg == "" {
		return MsgSkipTest
	} else {
		return MsgSkipTest + ": " + msg
	}
}

func client() (*saviyntapigoclient.Client, saviyntapigoclient.Credentials, bool, string, error) {
	creds := saviyntapigoclient.Credentials{}
	if v := strings.TrimSpace(os.Getenv(EnvSaviyntTestCredentials)); v == "" {
		return nil, creds, true, MsgSkipTestCredentialsNotSet, nil
	} else if clt, creds, err := saviyntapigoclient.NewClientPasswordEnv(context.Background(),
		EnvSaviyntTestCredentials); err != nil {
		skipMsgInner := fmt.Sprintf(MsgSkipTestErrorInstantiatingClient, err.Error())
		if err.Error() == "saviynt token response status (404)" {
			skipMsgInner = "saviynt token response status (404)" + ": serverURL format must be in the format of 'https://myidentity.saviyntcloud.com' without a URL path"
			err = errors.New(skipMsgInner)
		}
		return nil, creds, true, fmt.Sprintf(MsgSkipTestErrorInstantiatingClient, skipMsgInner), err
	} else if clt == nil {
		skipMsgInner := "client not generated"
		skipMsg := fmt.Sprintf(MsgSkipTestErrorInstantiatingClient, skipMsgInner)
		return nil, creds, true, skipMsg, errors.New(skipMsgInner)
	} else {
		return clt, creds, false, "", nil
	}
}

func PrintBody(r io.Reader) error {
	if b, err := io.ReadAll(r); err != nil {
		return err
	} else {
		fmt.Println(string(b))
		return nil
	}
}
