// Copyright (c) 2025 Saviynt Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/types"
)

func TestResolveProviderConfigValue_ServerURLEnv(t *testing.T) {
	t.Setenv(envSaviyntServerURL, "https://tenant.saviyntcloud.com")

	value, ok := resolveProviderConfigValue(types.StringNull(), envSaviyntServerURL)
	if !ok {
		t.Fatalf("expected %s to be used when server_url is not set", envSaviyntServerURL)
	}

	if value != "https://tenant.saviyntcloud.com" {
		t.Fatalf("unexpected server_url value from env: got %q", value)
	}
}

func TestResolveProviderConfigValue_UsernameEnv(t *testing.T) {
	t.Setenv(envSaviyntUsername, "terraform-user")

	value, ok := resolveProviderConfigValue(types.StringNull(), envSaviyntUsername)
	if !ok {
		t.Fatalf("expected %s to be used when username is not set", envSaviyntUsername)
	}

	if value != "terraform-user" {
		t.Fatalf("unexpected username value from env: got %q", value)
	}
}

func TestResolveProviderConfigValue_PasswordEnv(t *testing.T) {
	t.Setenv(envSaviyntPassword, "super-secret")

	value, ok := resolveProviderConfigValue(types.StringNull(), envSaviyntPassword)
	if !ok {
		t.Fatalf("expected %s to be used when password is not set", envSaviyntPassword)
	}

	if value != "super-secret" {
		t.Fatalf("unexpected password value from env: got %q", value)
	}
}

func TestResolveProviderConfigValue_ConfigWinsOverEnv(t *testing.T) {
	t.Setenv(envSaviyntUsername, "env-user")

	value, ok := resolveProviderConfigValue(types.StringValue("configured-user"), envSaviyntUsername)
	if !ok {
		t.Fatal("expected configured value to be used")
	}

	if value != "configured-user" {
		t.Fatalf("expected configured value to take precedence, got %q", value)
	}
}
