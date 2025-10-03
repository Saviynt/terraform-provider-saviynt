// Copyright (c) 2025 Saviynt Inc.
// SPDX-License-Identifier: MPL-2.0

package client

import (
	"context"

	s "github.com/saviynt/saviynt-api-go-client"
)

// SaviyntClientInterface defines the interface for the Saviynt client
// This is used across all resources to get the base API URL
type SaviyntClientInterface interface {
	APIBaseURL() string
}

// SaviyntClientWrapper wraps the actual client to implement the interface
type SaviyntClientWrapper struct {
	Client *s.Client
}

func (w *SaviyntClientWrapper) APIBaseURL() string {
	return w.Client.APIBaseURL()
}

// SaviyntProviderInterface defines the interface for provider operations needed by resources/datasources
type SaviyntProviderInterface interface {
	AuthenticatedAPICallWithRetry(ctx context.Context, operation string, apiCall func(token string) error) error
}

// SaviyntProviderWrapper wraps the actual provider to implement the interface
type SaviyntProviderWrapper struct {
	Provider SaviyntProviderInterface
}

func (w *SaviyntProviderWrapper) AuthenticatedAPICallWithRetry(ctx context.Context, operation string, apiCall func(token string) error) error {
	return w.Provider.AuthenticatedAPICallWithRetry(ctx, operation, apiCall)
}
