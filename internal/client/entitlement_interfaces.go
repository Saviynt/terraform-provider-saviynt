// Copyright (c) 2025 Saviynt Inc.
// SPDX-License-Identifier: MPL-2.0

package client

import (
	"context"
	"net/http"
	"strings"

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

func (w *EntitlementOperationsWrapper) CreateUpdateEntitlement(ctx context.Context, req openapi.CreateUpdateEntitlementRequest) (*openapi.CreateOrUpdateEntitlementResponse, *http.Response, error) {
	return w.client.EntitlementAPI.CreateUpdateEntitlement(ctx).CreateUpdateEntitlementRequest(req).Execute()
}

func (w *EntitlementOperationsWrapper) GetEntitlements(ctx context.Context, req openapi.GetEntitlementRequest) (*openapi.GetEntitlementResponse, *http.Response, error) {
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
	apiBaseURL := strings.TrimPrefix(strings.TrimPrefix(baseURL, "https://"), "http://")
	cfg.Host = apiBaseURL
	cfg.Scheme = "https"
	cfg.AddDefaultHeader("Authorization", "Bearer "+token)
	cfg.HTTPClient = http.DefaultClient
	apiClient := openapi.NewAPIClient(cfg)
	return &EntitlementOperationsWrapper{client: apiClient}
}
