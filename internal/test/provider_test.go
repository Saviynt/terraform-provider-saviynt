// Copyright (c) Saviynt Inc.
// SPDX-License-Identifier: MPL-2.0

// acceptance_test_utils.go provides the scaffolding required for Terraform acceptance testing
// of the Saviynt provider using Terraform CLI protocol v6 (gRPC). It defines factory functions
// for creating in-process provider servers and a pre-check helper to ensure necessary
// environment variables are set before tests run.
package testing

import (
	"os"
	"testing"
	"terraform-provider-Saviynt/internal/provider"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-testing/echoprovider"
)

// TestAccProtoV6ProviderFactories is used to instantiate a provider during acceptance testing.
// The factory function is called for each Terraform CLI command to create a provider
// server that the CLI can connect to and interact with.
var TestAccProtoV6ProviderFactories = map[string]func() (tfprotov6.ProviderServer, error){
	"saviynt": providerserver.NewProtocol6WithError(provider.New("test")()),
}

// TestAccProtoV6ProviderFactoriesWithEcho includes the echo provider alongside the scaffolding provider.
// It allows for testing assertions on data returned by an ephemeral resource during Open.
// The echoprovider is used to arrange tests by echoing ephemeral data into the Terraform state.
// This lets the data be referenced in test assertions with state checks.
var TestAccProtoV6ProviderFactoriesWithEcho = map[string]func() (tfprotov6.ProviderServer, error){
	"saviynt": providerserver.NewProtocol6WithError(provider.New("test")()),
	"echo":    echoprovider.NewProviderServer(),
}

func TestAccPreCheck(t *testing.T) {
	requiredEnvVars := []string{
		"SAVIYNT_URL",
		"SAVIYNT_USERNAME",
		"SAVIYNT_PASSWORD",
	}

	for _, envVar := range requiredEnvVars {
		if v := os.Getenv(envVar); v == "" {
			t.Fatalf("Environment variable %s must be set for acceptance tests", envVar)
		}
	}
}
