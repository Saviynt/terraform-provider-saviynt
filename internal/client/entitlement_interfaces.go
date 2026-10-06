// Copyright (c) 2025 Saviynt Inc.
// SPDX-License-Identifier: MPL-2.0

package client

import (
	"context"
	"net/http"

	openapi "github.com/saviynt/saviynt-api-go-client/entitlements"
)

// EntitlementTypeOperationsInterface defines the interface for entitlement type operations
type EntitlementOperationsInterface interface {
	CreateUpdateEntitlement(ctx context.Context, req openapi.CreateUpdateEntitlementRequest) (*openapi.CreateOrUpdateEntitlementResponse, *http.Response, error)
	GetEntitlements(ctx context.Context, req openapi.GetEntitlementRequest) (*openapi.GetEntitlementResponse, *http.Response, error)
}

// EntitlementTypeOperationsWrapper wraps the actual entitlement type operations to implement the interface
type EntitlementOperationsWrapper struct {
	client *openapi.APIClient
}

// NewEntitlementOperationsWrapper creates a wrapper from an already-configured APIClient.
// Primarily used in tests to point the real injection logic at a test HTTP server.
func NewEntitlementOperationsWrapper(apiClient *openapi.APIClient) *EntitlementOperationsWrapper {
	return &EntitlementOperationsWrapper{client: apiClient}
}

func (w *EntitlementOperationsWrapper) CreateUpdateEntitlement(ctx context.Context, req openapi.CreateUpdateEntitlementRequest) (*openapi.CreateOrUpdateEntitlementResponse, *http.Response, error) {
	return w.client.EntitlementAPI.CreateUpdateEntitlement(ctx).CreateUpdateEntitlementRequest(req).Execute()
}

func (w *EntitlementOperationsWrapper) GetEntitlements(ctx context.Context, req openapi.GetEntitlementRequest) (*openapi.GetEntitlementResponse, *http.Response, error) {
	// Always request backend field names regardless of server-side readlabels config.
	req.SetReadlabels("false")
	return w.client.EntitlementAPI.GetEntitlements(ctx).GetEntitlementRequest(req).Execute()
}

// EntitlementTypeFactoryInterface defines the interface for creating entitlement type operations
type EntitlementFactoryInterface interface {
	CreateEntitlementOperations(baseURL, token string) EntitlementOperationsInterface
}

// DefaultEntitlementTypeFactory implements the EntitlementTypeFactoryInterface
type DefaultEntitlementFactory struct{}

func (f *DefaultEntitlementFactory) CreateEntitlementOperations(baseURL, token string) EntitlementOperationsInterface {
	cfg := openapi.NewConfiguration()
	cfg.Servers = openapi.ServerConfigurations{{URL: baseURL}}
	cfg.AddDefaultHeader("Authorization", "Bearer "+token)
	cfg.HTTPClient = http.DefaultClient
	apiClient := openapi.NewAPIClient(cfg)
	return &EntitlementOperationsWrapper{client: apiClient}
}
