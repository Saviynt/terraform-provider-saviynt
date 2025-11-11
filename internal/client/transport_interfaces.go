// Copyright (c) 2025 Saviynt Inc.
// SPDX-License-Identifier: MPL-2.0

package client

import (
	"context"
	"net/http"
	"strings"

	openapi "github.com/saviynt/saviynt-api-go-client/transports"
)

// TransportOperationsInterface defines the interface for transport operations
type TransportOperationsInterface interface {
	ExportTransportPackage(ctx context.Context, req openapi.ExportTransportPackageRequest) (*openapi.ExportTransportPackageResponse, *http.Response, error)
	ImportTransportPackage(ctx context.Context, req openapi.ImportTransportPackageRequest) (*openapi.ImportTransportPackageResponse, *http.Response, error)
	TransportPackageStatus(ctx context.Context, req openapi.TransportPackageStatusRequest) (*openapi.TransportPackageStatusResponse, *http.Response, error)
}

// TransportOperationsWrapper wraps the actual transport operations to implement the interface
type TransportOperationsWrapper struct {
	client *openapi.APIClient
}

func (w *TransportOperationsWrapper) ExportTransportPackage(ctx context.Context, req openapi.ExportTransportPackageRequest) (*openapi.ExportTransportPackageResponse, *http.Response, error) {
	return w.client.TransportsAPI.ExportTransportPackage(ctx).ExportTransportPackageRequest(req).Execute()
}

func (w *TransportOperationsWrapper) ImportTransportPackage(ctx context.Context, req openapi.ImportTransportPackageRequest) (*openapi.ImportTransportPackageResponse, *http.Response, error) {
	return w.client.TransportsAPI.ImportTransportPackage(ctx).ImportTransportPackageRequest(req).Execute()
}

func (w *TransportOperationsWrapper) TransportPackageStatus(ctx context.Context, req openapi.TransportPackageStatusRequest) (*openapi.TransportPackageStatusResponse, *http.Response, error) {
	return w.client.TransportsAPI.TransportPackageStatus(ctx).TransportPackageStatusRequest(req).Execute()
}

// TransportFactoryInterface defines the interface for creating transport operations
type TransportFactoryInterface interface {
	CreateTransportOperations(baseURL, token string) TransportOperationsInterface
}

// DefaultTransportFactory implements the TransportFactoryInterface
type DefaultTransportFactory struct{}

func (f *DefaultTransportFactory) CreateTransportOperations(baseURL, token string) TransportOperationsInterface {
	cfg := openapi.NewConfiguration()
	apiBaseURL := strings.TrimPrefix(strings.TrimPrefix(baseURL, "https://"), "http://")
	cfg.Host = apiBaseURL
	cfg.Scheme = "https"
	cfg.AddDefaultHeader("Authorization", "Bearer "+token)
	cfg.HTTPClient = http.DefaultClient
	apiClient := openapi.NewAPIClient(cfg)
	return &TransportOperationsWrapper{client: apiClient}
}
