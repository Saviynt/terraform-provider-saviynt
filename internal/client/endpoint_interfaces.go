// Copyright (c) 2025 Saviynt Inc.
// SPDX-License-Identifier: MPL-2.0

package client

import (
	"context"
	"net/http"
	"strings"

	openapi "github.com/saviynt/saviynt-api-go-client/endpoints"
)

// EndpointOperationsInterface defines the interface for endpoint operations
type EndpointOperationsInterface interface {
	CreateEndpoint(ctx context.Context, req openapi.CreateEndpointRequest) (*openapi.UpdateEndpoint200Response, *http.Response, error)
	UpdateEndpoint(ctx context.Context, req openapi.UpdateEndpointRequest) (*openapi.UpdateEndpoint200Response, *http.Response, error)
	GetEndpoints(ctx context.Context, req openapi.GetEndpointsRequest) (*openapi.GetEndpoints200Response, *http.Response, error)
}

// EndpointOperationsWrapper wraps the actual endpoint operations to implement the interface
type EndpointOperationsWrapper struct {
	client *openapi.APIClient
}

func (w *EndpointOperationsWrapper) CreateEndpoint(ctx context.Context, req openapi.CreateEndpointRequest) (*openapi.UpdateEndpoint200Response, *http.Response, error) {
	return w.client.EndpointsAPI.CreateEndpoint(ctx).CreateEndpointRequest(req).Execute()
}

func (w *EndpointOperationsWrapper) UpdateEndpoint(ctx context.Context, req openapi.UpdateEndpointRequest) (*openapi.UpdateEndpoint200Response, *http.Response, error) {
	return w.client.EndpointsAPI.UpdateEndpoint(ctx).UpdateEndpointRequest(req).Execute()
}

func (w *EndpointOperationsWrapper) GetEndpoints(ctx context.Context, req openapi.GetEndpointsRequest) (*openapi.GetEndpoints200Response, *http.Response, error) {
	return w.client.EndpointsAPI.GetEndpoints(ctx).GetEndpointsRequest(req).Execute()
}

// EndpointFactoryInterface defines the interface for creating endpoint operations
type EndpointFactoryInterface interface {
	CreateEndpointOperations(baseURL, token string) EndpointOperationsInterface
}

// DefaultEndpointFactory implements the EndpointFactoryInterface
type DefaultEndpointFactory struct{}

func (f *DefaultEndpointFactory) CreateEndpointOperations(baseURL, token string) EndpointOperationsInterface {
	cfg := openapi.NewConfiguration()
	apiBaseURL := strings.TrimPrefix(strings.TrimPrefix(baseURL, "https://"), "http://")
	cfg.Host = apiBaseURL
	cfg.Scheme = "https"
	cfg.AddDefaultHeader("Authorization", "Bearer "+token)
	cfg.HTTPClient = http.DefaultClient
	apiClient := openapi.NewAPIClient(cfg)
	return &EndpointOperationsWrapper{client: apiClient}
}
