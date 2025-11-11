// Copyright (c) 2025 Saviynt Inc.
// SPDX-License-Identifier: MPL-2.0

package client

import (
	"context"
	"net/http"
	"strings"

	endpoint "github.com/saviynt/saviynt-api-go-client/endpoints"
	openapi "github.com/saviynt/saviynt-api-go-client/roles"
)

// RoleOperationsInterface defines the interface for role operations
// This interface is used by the roles resource for dependency injection
type RoleOperationsInterface interface {
	CreateEnterpriseRole(ctx context.Context, req openapi.CreateEnterpriseRoleRequest) (*openapi.CreateEnterpriseRoleResponse, *http.Response, error)
	GetRoles(ctx context.Context, req openapi.GetRolesRequest) (*openapi.GetRolesResponse, *http.Response, error)
	UpdateEnterpriseRole(ctx context.Context, req openapi.UpdateEnterpriseRoleRequest) (*openapi.UpdateEnterpriseRoleResponse, *http.Response, error)
	AddUserToRole(ctx context.Context, userName string, roleName string) (*openapi.AddOrRemoveRoleResponse, *http.Response, error)
	RemoveUserFromRole(ctx context.Context, userName string, roleName string) (*openapi.AddOrRemoveRoleResponse, *http.Response, error)
}

// RoleOperationsWrapper wraps the actual role operations to implement the interface
type RoleOperationsWrapper struct {
	client *openapi.APIClient
}

func (w *RoleOperationsWrapper) CreateEnterpriseRole(ctx context.Context, req openapi.CreateEnterpriseRoleRequest) (*openapi.CreateEnterpriseRoleResponse, *http.Response, error) {
	return w.client.RolesAPI.CreateEnterpriseRoleRequest(ctx).CreateEnterpriseRoleRequest(req).Execute()
}

func (w *RoleOperationsWrapper) GetRoles(ctx context.Context, req openapi.GetRolesRequest) (*openapi.GetRolesResponse, *http.Response, error) {
	return w.client.RolesAPI.GetRoles(ctx).GetRolesRequest(req).Execute()
}

func (w *RoleOperationsWrapper) UpdateEnterpriseRole(ctx context.Context, req openapi.UpdateEnterpriseRoleRequest) (*openapi.UpdateEnterpriseRoleResponse, *http.Response, error) {
	return w.client.RolesAPI.UpdateEnterpriseRoleRequest(ctx).UpdateEnterpriseRoleRequest(req).Execute()
}

func (w *RoleOperationsWrapper) AddUserToRole(ctx context.Context, userName string, roleName string) (*openapi.AddOrRemoveRoleResponse, *http.Response, error) {
	req := openapi.AddOrRemoveRoleRequest{
		Username: userName,
		Rolename: roleName,
	}
	return w.client.RolesAPI.Addrole(ctx).AddOrRemoveRoleRequest(req).Execute()
}

func (w *RoleOperationsWrapper) RemoveUserFromRole(ctx context.Context, userName string, roleName string) (*openapi.AddOrRemoveRoleResponse, *http.Response, error) {
	req := openapi.AddOrRemoveRoleRequest{
		Username: userName,
		Rolename: roleName,
	}
	return w.client.RolesAPI.Removerole(ctx).AddOrRemoveRoleRequest(req).Execute()
}

// RoleFactoryInterface defines the interface for creating role operations
// This factory is used by the roles resource for dependency injection
type RoleFactoryInterface interface {
	CreateRoleOperations(baseURL, token string) RoleOperationsInterface
	CreateEndpointOperations(baseURL, token string) EndpointOperationsInterface
}

// DefaultRoleFactory implements the RoleFactoryInterface
type DefaultRoleFactory struct{}

func (f *DefaultRoleFactory) CreateRoleOperations(baseURL, token string) RoleOperationsInterface {
	cfg := openapi.NewConfiguration()
	apiBaseURL := strings.TrimPrefix(strings.TrimPrefix(baseURL, "https://"), "http://")
	cfg.Host = apiBaseURL
	cfg.Scheme = "https"
	cfg.AddDefaultHeader("Authorization", "Bearer "+token)
	cfg.HTTPClient = http.DefaultClient
	apiClient := openapi.NewAPIClient(cfg)
	return &RoleOperationsWrapper{client: apiClient}
}

func (f *DefaultRoleFactory) CreateEndpointOperations(baseURL, token string) EndpointOperationsInterface {
	cfg := endpoint.NewConfiguration()
	apiBaseURL := strings.TrimPrefix(strings.TrimPrefix(baseURL, "https://"), "http://")
	cfg.Host = apiBaseURL
	cfg.Scheme = "https"
	cfg.AddDefaultHeader("Authorization", "Bearer "+token)
	cfg.HTTPClient = http.DefaultClient
	apiClient := endpoint.NewAPIClient(cfg)
	return &EndpointOperationsWrapper{client: apiClient}
}
