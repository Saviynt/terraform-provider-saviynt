// Copyright (c) 2025 Saviynt Inc.
// SPDX-License-Identifier: MPL-2.0

package client

import (
	"context"
	"net/http"
	"strings"

	openapi "github.com/saviynt/saviynt-api-go-client/connections"
)

// ConnectionOperationsInterface defines the generic interface for all connection operations
// This interface is used by all connection resources (AD, ADSI, REST, DB, etc.)
type ConnectionOperationsInterface interface {
	GetConnectionDetails(ctx context.Context, connectionName string) (*openapi.GetConnectionDetailsResponse, *http.Response, error)
	CreateOrUpdateConnection(ctx context.Context, req openapi.CreateOrUpdateRequest) (*openapi.CreateOrUpdateResponse, *http.Response, error)
	GetConnectionDetailsDataSource(ctx context.Context, connectionParam openapi.GetConnectionDetailsRequest) (*openapi.GetConnectionDetailsResponse, *http.Response, error)
	GetConnectionsDataSource(ctx context.Context, req openapi.GetConnectionsRequest) (*openapi.GetConnectionsResponse, *http.Response, error)
}

// ConnectionOperationsWrapper wraps the actual connection operations to implement the interface
type ConnectionOperationsWrapper struct {
	client *openapi.APIClient
}

func (w *ConnectionOperationsWrapper) GetConnectionDetails(ctx context.Context, connectionName string) (*openapi.GetConnectionDetailsResponse, *http.Response, error) {
	reqParams := openapi.GetConnectionDetailsRequest{}
	reqParams.SetConnectionname(connectionName)
	return w.client.ConnectionsAPI.GetConnectionDetails(ctx).GetConnectionDetailsRequest(reqParams).Execute()
}

func (w *ConnectionOperationsWrapper) CreateOrUpdateConnection(ctx context.Context, req openapi.CreateOrUpdateRequest) (*openapi.CreateOrUpdateResponse, *http.Response, error) {
	return w.client.ConnectionsAPI.CreateOrUpdate(ctx).CreateOrUpdateRequest(req).Execute()
}

func (w *ConnectionOperationsWrapper) GetConnectionDetailsDataSource(ctx context.Context, connectionParam openapi.GetConnectionDetailsRequest) (*openapi.GetConnectionDetailsResponse, *http.Response, error) {
	return w.client.ConnectionsAPI.GetConnectionDetails(ctx).GetConnectionDetailsRequest(connectionParam).Execute()
}

func (w *ConnectionOperationsWrapper) GetConnectionsDataSource(ctx context.Context, req openapi.GetConnectionsRequest) (*openapi.GetConnectionsResponse, *http.Response, error) {
	return w.client.ConnectionsAPI.GetConnections(ctx).GetConnectionsRequest(req).Execute()
}

// ConnectionFactoryInterface defines the interface for creating connection operations
// This factory is used by all connection resources for dependency injection
type ConnectionFactoryInterface interface {
	CreateConnectionOperations(baseURL, token string) ConnectionOperationsInterface
}

// DefaultConnectionFactory implements the ConnectionFactoryInterface
type DefaultConnectionFactory struct{}

func (f *DefaultConnectionFactory) CreateConnectionOperations(baseURL, token string) ConnectionOperationsInterface {
	cfg := openapi.NewConfiguration()
	apiBaseURL := strings.TrimPrefix(strings.TrimPrefix(baseURL, "https://"), "http://")
	cfg.Host = apiBaseURL
	cfg.Scheme = "https"
	cfg.AddDefaultHeader("Authorization", "Bearer "+token)
	cfg.HTTPClient = http.DefaultClient
	apiClient := openapi.NewAPIClient(cfg)
	return &ConnectionOperationsWrapper{client: apiClient}
}
