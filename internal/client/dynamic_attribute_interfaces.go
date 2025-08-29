// Copyright (c) 2025 Saviynt Inc.
// SPDX-License-Identifier: MPL-2.0

package client

import (
	"context"
	"net/http"
	"strings"

	openapi "github.com/saviynt/saviynt-api-go-client/dynamicattributes"
	endpoint "github.com/saviynt/saviynt-api-go-client/endpoints"
)

// DynamicAttributeOperationsInterface defines the interface for dynamic attribute operations
// This interface is used by the dynamic attribute resource for dependency injection
type DynamicAttributeOperationsInterface interface {
	CreateDynamicAttribute(ctx context.Context, req openapi.CreateDynamicAttributeRequest) (*openapi.CreateOrUpdateOrDeleteDynamicAttributeResponse, *http.Response, error)
	FetchDynamicAttribute(ctx context.Context, endpointName string) (*openapi.FetchDynamicAttributesResponse, *http.Response, error)
	UpdateDynamicAttribute(ctx context.Context, req openapi.UpdateDynamicAttributeRequest) (*openapi.CreateOrUpdateOrDeleteDynamicAttributeResponse, *http.Response, error)
	DeleteDynamicAttribute(ctx context.Context, req openapi.DeleteDynamicAttributeRequest) (*openapi.CreateOrUpdateOrDeleteDynamicAttributeResponse, *http.Response, error)
}

// DynamicAttributeOperationsWrapper wraps the actual dynamic attribute operations to implement the interface
type DynamicAttributeOperationsWrapper struct {
	client *openapi.APIClient
}

func (w *DynamicAttributeOperationsWrapper) CreateDynamicAttribute(ctx context.Context, req openapi.CreateDynamicAttributeRequest) (*openapi.CreateOrUpdateOrDeleteDynamicAttributeResponse, *http.Response, error) {
	return w.client.DynamicAttributesAPI.CreateDynamicAttribute(ctx).CreateDynamicAttributeRequest(req).Execute()
}

func (w *DynamicAttributeOperationsWrapper) FetchDynamicAttribute(ctx context.Context, endpointName string) (*openapi.FetchDynamicAttributesResponse, *http.Response, error) {
	fetchReq := w.client.DynamicAttributesAPI.FetchDynamicAttribute(ctx)
	if endpointName != "" {
		fetchReq = fetchReq.Endpoint([]string{endpointName})
	}
	return fetchReq.Execute()
}

func (w *DynamicAttributeOperationsWrapper) UpdateDynamicAttribute(ctx context.Context, req openapi.UpdateDynamicAttributeRequest) (*openapi.CreateOrUpdateOrDeleteDynamicAttributeResponse, *http.Response, error) {
	return w.client.DynamicAttributesAPI.UpdateDynamicAttribute(ctx).UpdateDynamicAttributeRequest(req).Execute()
}

func (w *DynamicAttributeOperationsWrapper) DeleteDynamicAttribute(ctx context.Context, req openapi.DeleteDynamicAttributeRequest) (*openapi.CreateOrUpdateOrDeleteDynamicAttributeResponse, *http.Response, error) {
	return w.client.DynamicAttributesAPI.DeleteDynamicAttribute(ctx).DeleteDynamicAttributeRequest(req).Execute()
}

// DynamicAttributeFactoryInterface defines the interface for creating dynamic attribute operations
// This factory is used by the dynamic attribute resource for dependency injection
type DynamicAttributeFactoryInterface interface {
	CreateDynamicAttributeOperations(baseURL, token string) DynamicAttributeOperationsInterface
	CreateEndpointOperations(baseURL, token string) EndpointOperationsInterface
}

// DefaultDynamicAttributeFactory implements the DynamicAttributeFactoryInterface
type DefaultDynamicAttributeFactory struct{}

func (f *DefaultDynamicAttributeFactory) CreateDynamicAttributeOperations(baseURL, token string) DynamicAttributeOperationsInterface {
	cfg := openapi.NewConfiguration()
	apiBaseURL := strings.TrimPrefix(strings.TrimPrefix(baseURL, "https://"), "http://")
	cfg.Host = apiBaseURL
	cfg.Scheme = "https"
	cfg.AddDefaultHeader("Authorization", "Bearer "+token)
	cfg.HTTPClient = http.DefaultClient
	apiClient := openapi.NewAPIClient(cfg)
	return &DynamicAttributeOperationsWrapper{client: apiClient}
}

func (f *DefaultDynamicAttributeFactory) CreateEndpointOperations(baseURL, token string) EndpointOperationsInterface {
	cfg := endpoint.NewConfiguration()
	apiBaseURL := strings.TrimPrefix(strings.TrimPrefix(baseURL, "https://"), "http://")
	cfg.Host = apiBaseURL
	cfg.Scheme = "https"
	cfg.AddDefaultHeader("Authorization", "Bearer "+token)
	cfg.HTTPClient = http.DefaultClient
	apiClient := endpoint.NewAPIClient(cfg)
	return &EndpointOperationsWrapper{client: apiClient}
}