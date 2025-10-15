// Copyright (c) 2025 Saviynt Inc.
// SPDX-License-Identifier: MPL-2.0

package client

import (
	"context"
	"net/http"
	"strings"

	openapi "github.com/saviynt/saviynt-api-go-client/securitysystems"
)

// SecuritySystemOperationsInterface defines the interface for security system operations
type SecuritySystemOperationsInterface interface {
	CreateSecuritySystem(ctx context.Context, req openapi.CreateSecuritySystemRequest) (*openapi.CreateSecuritySystem200Response, *http.Response, error)
	UpdateSecuritySystem(ctx context.Context, req openapi.UpdateSecuritySystemRequest) (*openapi.CreateSecuritySystem200Response, *http.Response, error)
	GetSecuritySystems(ctx context.Context, systemname string) (*openapi.GetSecuritySystems200Response, *http.Response, error)
	GetSecuritySystemsRequest(ctx context.Context) openapi.ApiGetSecuritySystemsRequest
}

// SecuritySystemOperationsWrapper wraps the actual security system operations to implement the interface
type SecuritySystemOperationsWrapper struct {
	client *openapi.APIClient
}

func (w *SecuritySystemOperationsWrapper) CreateSecuritySystem(ctx context.Context, req openapi.CreateSecuritySystemRequest) (*openapi.CreateSecuritySystem200Response, *http.Response, error) {
	return w.client.SecuritySystemsAPI.CreateSecuritySystem(ctx).CreateSecuritySystemRequest(req).Execute()
}

func (w *SecuritySystemOperationsWrapper) UpdateSecuritySystem(ctx context.Context, req openapi.UpdateSecuritySystemRequest) (*openapi.CreateSecuritySystem200Response, *http.Response, error) {
	return w.client.SecuritySystemsAPI.UpdateSecuritySystem(ctx).UpdateSecuritySystemRequest(req).Execute()
}

func (w *SecuritySystemOperationsWrapper) GetSecuritySystems(ctx context.Context, systemname string) (*openapi.GetSecuritySystems200Response, *http.Response, error) {
	return w.client.SecuritySystemsAPI.GetSecuritySystems(ctx).Systemname(systemname).Execute()
}

func (w *SecuritySystemOperationsWrapper) GetSecuritySystemsRequest(ctx context.Context) openapi.ApiGetSecuritySystemsRequest {
	return w.client.SecuritySystemsAPI.GetSecuritySystems(ctx)
}

// SecuritySystemFactoryInterface defines the interface for creating security system operations
type SecuritySystemFactoryInterface interface {
	CreateSecuritySystemOperations(baseURL, token string) SecuritySystemOperationsInterface
}

// DefaultSecuritySystemFactory implements the SecuritySystemFactoryInterface
type DefaultSecuritySystemFactory struct{}

func (f *DefaultSecuritySystemFactory) CreateSecuritySystemOperations(baseURL, token string) SecuritySystemOperationsInterface {
	cfg := openapi.NewConfiguration()
	apiBaseURL := strings.TrimPrefix(strings.TrimPrefix(baseURL, "https://"), "http://")
	cfg.Host = apiBaseURL
	cfg.Scheme = "https"
	cfg.AddDefaultHeader("Authorization", "Bearer "+token)
	cfg.HTTPClient = http.DefaultClient
	apiClient := openapi.NewAPIClient(cfg)
	return &SecuritySystemOperationsWrapper{client: apiClient}
}
