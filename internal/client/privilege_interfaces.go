// Copyright (c) 2025 Saviynt Inc.
// SPDX-License-Identifier: MPL-2.0

package client

import (
	"context"
	"net/http"
	"strings"

	endpoint "github.com/saviynt/saviynt-api-go-client/endpoints"
	openapi "github.com/saviynt/saviynt-api-go-client/privileges"
)

type PrivilegeOperationInterface interface {
	CreatePrivilege(ctx context.Context, req openapi.CreateUpdatePrivilegeRequest) (*openapi.CreateUpdatePrivilegeResponse, *http.Response, error)
	GetPrivilege(ctx context.Context, req openapi.GetPrivilegeListRequest) (*openapi.GetPrivilegeListResponse, *http.Response, error)
	UpdatePrivilege(ctx context.Context, req openapi.CreateUpdatePrivilegeRequest) (*openapi.CreateUpdatePrivilegeResponse, *http.Response, error)
	DeletePrivilege(ctx context.Context, req openapi.DeletePrivilegeRequest) (*openapi.DeletePrivilegeResponse, *http.Response, error)
}

type PrivilegeOperationsWrapper struct {
	client *openapi.APIClient
}

func (p *PrivilegeOperationsWrapper) CreatePrivilege(ctx context.Context, req openapi.CreateUpdatePrivilegeRequest) (*openapi.CreateUpdatePrivilegeResponse, *http.Response, error) {
	return p.client.PrivilegeAPI.CreatePrivilege(ctx).CreateUpdatePrivilegeRequest(req).Execute()
}

func (p *PrivilegeOperationsWrapper) GetPrivilege(ctx context.Context, req openapi.GetPrivilegeListRequest) (*openapi.GetPrivilegeListResponse, *http.Response, error) {
	return p.client.PrivilegeAPI.GetPrivilege(ctx).GetPrivilegeListRequest(req).Execute()
}

func (p *PrivilegeOperationsWrapper) UpdatePrivilege(ctx context.Context, req openapi.CreateUpdatePrivilegeRequest) (*openapi.CreateUpdatePrivilegeResponse, *http.Response, error) {
	return p.client.PrivilegeAPI.UpdatePrivilege(ctx).CreateUpdatePrivilegeRequest(req).Execute()
}

func (p *PrivilegeOperationsWrapper) DeletePrivilege(ctx context.Context, req openapi.DeletePrivilegeRequest) (*openapi.DeletePrivilegeResponse, *http.Response, error) {
	return p.client.PrivilegeAPI.DeletePrivilege(ctx).DeletePrivilegeRequest(req).Execute()
}

type PrivilegeFactoryInterface interface {
	CreatePrivilegeOperations(baseURL, token string) PrivilegeOperationInterface
	CreateEndpointOperations(baseURL, token string) EndpointOperationsInterface
}

type DefaultPrivilegeFactory struct{}

func (f *DefaultPrivilegeFactory) CreatePrivilegeOperations(baseURL, token string) PrivilegeOperationInterface {
	cfg := openapi.NewConfiguration()
	apiBaseURL := strings.TrimPrefix(strings.TrimPrefix(baseURL, "https://"), "http://")
	cfg.Host = apiBaseURL
	cfg.Scheme = "https"
	cfg.AddDefaultHeader("Authorization", "Bearer "+token)
	cfg.HTTPClient = http.DefaultClient
	apiClient := openapi.NewAPIClient(cfg)
	return &PrivilegeOperationsWrapper{client: apiClient}
}

func (f *DefaultPrivilegeFactory) CreateEndpointOperations(baseURL, token string) EndpointOperationsInterface {
	cfg := endpoint.NewConfiguration()
	apiBaseURL := strings.TrimPrefix(strings.TrimPrefix(baseURL, "https://"), "http://")
	cfg.Host = apiBaseURL
	cfg.Scheme = "https"
	cfg.AddDefaultHeader("Authorization", "Bearer "+token)
	cfg.HTTPClient = http.DefaultClient
	apiClient := endpoint.NewAPIClient(cfg)
	return &EndpointOperationsWrapper{client: apiClient}
}
