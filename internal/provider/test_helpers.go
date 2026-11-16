// Copyright (c) 2025 Saviynt Inc.
// SPDX-License-Identifier: MPL-2.0

// test_helpers.go exposes a minimal set of constructors so unit tests in the
// `unit_tests` tree can build a SaviyntProvider with controlled token state
// without touching the real authentication flow. These helpers are intentionally
// tiny and exist only to enable testing of the auth retry logic.

package provider

import (
	"context"

	s "github.com/saviynt/saviynt-api-go-client"
	"golang.org/x/oauth2"
)

// NewProviderWithTokens constructs a SaviyntProvider for tests, pre-populating
// the access/refresh tokens. The client is left nil, so only code paths that
// don't dereference p.client (e.g. the "no refresh token" guard) are safe to
// exercise.
func NewProviderWithTokens(accessToken, refreshToken string) *SaviyntProvider {
	return &SaviyntProvider{
		accessToken:  accessToken,
		refreshToken: refreshToken,
	}
}

// NewProviderWithClient constructs a SaviyntProvider with a non-nil client
// so Configure handlers can be driven past the `client == nil` guard. The
// client is not wired to anything real; any operation that actually touches
// the network will fail, but Configure only stores the pointer so this is
// sufficient for coverage of the happy path.
func NewProviderWithClient(accessToken string) *SaviyntProvider {
	ctx := context.Background()
	username := "test-user"
	tok := &oauth2.Token{AccessToken: accessToken}
	c := s.NewClientToken(ctx, "https://example.test", &username, tok)
	return &SaviyntProvider{
		accessToken:  accessToken,
		refreshToken: "test-refresh-token",
		client:       c,
	}
}

// UpdateModelFromCreateResponseForTest exposes the unexported
// updateModelFromCreateResponse helper on RolesResource for unit tests.
func (r *RolesResource) UpdateModelFromCreateResponseForTest(plan *RolesResourceModel) {
	r.updateModelFromCreateResponse(plan)
}

// SetCustomPropertiesInPlanForTest exposes the unexported helper for tests.
func (r *RolesResource) SetCustomPropertiesInPlanForTest(plan *RolesResourceModel) {
	r.setCustomPropertiesInPlan(plan)
}
