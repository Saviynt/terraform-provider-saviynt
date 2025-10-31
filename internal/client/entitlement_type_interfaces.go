// Copyright (c) 2025 Saviynt Inc.
// SPDX-License-Identifier: MPL-2.0

package client

import (
	"context"
	"net/http"
	"strings"

	openapi "github.com/saviynt/saviynt-api-go-client/entitlementtype"
)

// EntitlementTypeOperationsInterface defines the interface for entitlement type operations
type EntitlementTypeOperationsInterface interface {
	CreateEntitlementType(ctx context.Context, req openapi.CreateEntitlementTypeRequest) (*openapi.CreateOrUpdateEntitlementTypeResponse, *http.Response, error)
	UpdateEntitlementType(ctx context.Context, req openapi.UpdateEntitlementTypeRequest) (*openapi.CreateOrUpdateEntitlementTypeResponse, *http.Response, error)
	GetEntitlementType(ctx context.Context, entitlementname, max, offset, endpointname string) (*openapi.GetEntitlementTypeResponse, *http.Response, error)
}

// EntitlementTypeOperationsWrapper wraps the actual entitlement type operations to implement the interface
type EntitlementTypeOperationsWrapper struct {
	client *openapi.APIClient
}

func (w *EntitlementTypeOperationsWrapper) CreateEntitlementType(ctx context.Context, req openapi.CreateEntitlementTypeRequest) (*openapi.CreateOrUpdateEntitlementTypeResponse, *http.Response, error) {
	return w.client.EntitlementTypeAPI.CreateEntitlementType(ctx).CreateEntitlementTypeRequest(req).Execute()
}

func (w *EntitlementTypeOperationsWrapper) UpdateEntitlementType(ctx context.Context, req openapi.UpdateEntitlementTypeRequest) (*openapi.CreateOrUpdateEntitlementTypeResponse, *http.Response, error) {
	return w.client.EntitlementTypeAPI.UpdateEntitlementType(ctx).UpdateEntitlementTypeRequest(req).Execute()
}

func (w *EntitlementTypeOperationsWrapper) GetEntitlementType(ctx context.Context, entitlementname, max, offset, endpointname string) (*openapi.GetEntitlementTypeResponse, *http.Response, error) {
	req := w.client.EntitlementTypeAPI.GetEntitlementType(ctx)

	if entitlementname != "" {
		req = req.Entitlementname(entitlementname)
	}
	if max != "" {
		req = req.Max(max)
	}
	if offset != "" {
		req = req.Offset(offset)
	}
	if endpointname != "" {
		req = req.Endpointname(endpointname)
	}

	return req.Execute()
}

// EntitlementTypeFactoryInterface defines the interface for creating entitlement type operations
type EntitlementTypeFactoryInterface interface {
	CreateEntitlementTypeOperations(baseURL, token string) EntitlementTypeOperationsInterface
}

// DefaultEntitlementTypeFactory implements the EntitlementTypeFactoryInterface
type DefaultEntitlementTypeFactory struct{}

func (f *DefaultEntitlementTypeFactory) CreateEntitlementTypeOperations(baseURL, token string) EntitlementTypeOperationsInterface {
	cfg := openapi.NewConfiguration()
	apiBaseURL := strings.TrimPrefix(strings.TrimPrefix(baseURL, "https://"), "http://")
	cfg.Host = apiBaseURL
	cfg.Scheme = "https"
	cfg.AddDefaultHeader("Authorization", "Bearer "+token)
	cfg.HTTPClient = http.DefaultClient
	apiClient := openapi.NewAPIClient(cfg)
	return &EntitlementTypeOperationsWrapper{client: apiClient}
}
