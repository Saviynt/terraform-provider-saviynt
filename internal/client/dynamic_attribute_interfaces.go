// Copyright (c) 2025 Saviynt Inc.
// SPDX-License-Identifier: MPL-2.0

package client

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"

	openapi "github.com/saviynt/saviynt-api-go-client/dynamicattributes"
	endpoint "github.com/saviynt/saviynt-api-go-client/endpoints"
)

// DynamicAttributeOperationsInterface defines the interface for dynamic attribute operations
// This interface is used by the dynamic attribute resource for dependency injection
type DynamicAttributeOperationsInterface interface {
	CreateDynamicAttribute(ctx context.Context, req openapi.CreateDynamicAttributeRequest) (*openapi.CreateOrUpdateOrDeleteDynamicAttributeResponse, *http.Response, error)
	FetchDynamicAttribute(ctx context.Context, endpointName string) (*openapi.FetchDynamicAttributesResponse, *http.Response, error)
	FetchDynamicAttributesForDataSource(ctx context.Context, securitySystems, endpoints, dynamicAttributes, requestTypes []string, loggedInUser, offset, max *string) (*openapi.FetchDynamicAttributesResponse, *http.Response, error)
	UpdateDynamicAttribute(ctx context.Context, req openapi.UpdateDynamicAttributeRequest) (*openapi.CreateOrUpdateOrDeleteDynamicAttributeResponse, *http.Response, error)
	DeleteDynamicAttribute(ctx context.Context, req openapi.DeleteDynamicAttributeRequest) (*openapi.CreateOrUpdateOrDeleteDynamicAttributeResponse, *http.Response, error)
}

// DynamicAttributeOperationsWrapper wraps the actual dynamic attribute operations to implement the interface
type DynamicAttributeOperationsWrapper struct {
	client *openapi.APIClient
}

// NewDynamicAttributeOperationsWrapper creates a wrapper from an already-configured APIClient.
// Primarily used in tests to point the real pagination logic at a test HTTP server.
func NewDynamicAttributeOperationsWrapper(apiClient *openapi.APIClient) *DynamicAttributeOperationsWrapper {
	return &DynamicAttributeOperationsWrapper{client: apiClient}
}

func (w *DynamicAttributeOperationsWrapper) CreateDynamicAttribute(ctx context.Context, req openapi.CreateDynamicAttributeRequest) (*openapi.CreateOrUpdateOrDeleteDynamicAttributeResponse, *http.Response, error) {
	return w.client.DynamicAttributesAPI.CreateDynamicAttribute(ctx).CreateDynamicAttributeRequest(req).Execute()
}

// FetchDynamicAttribute retrieves all dynamic attributes for the given endpoint by paginating
// through the API until every attribute is collected. The Saviynt API defaults max to 50 when
// the parameter is omitted, which would silently truncate state for endpoints with more than 50
// attributes. This method uses a page size of 500 and loops until the accumulated count reaches
// the totalcount advertised by the API, or until an empty page is returned as a safety stop.
// The caller receives a single merged FetchDynamicAttributesResponse as if the API had no limit.
func (w *DynamicAttributeOperationsWrapper) FetchDynamicAttribute(ctx context.Context, endpointName string) (*openapi.FetchDynamicAttributesResponse, *http.Response, error) {
	const pageSize = 500

	var (
		allAttrs   []openapi.FetchDynamicAttributeResponseInner
		lastResp   *http.Response
		finalMsg   *string
		finalECode *string
		offset     = 0
	)

	for {
		fetchReq := w.client.DynamicAttributesAPI.FetchDynamicAttribute(ctx).
			Max(strconv.Itoa(pageSize)).
			Offset(strconv.Itoa(offset))

		if endpointName != "" {
			fetchReq = fetchReq.Endpoint([]string{endpointName})
		}

		page, httpResp, err := fetchReq.Execute()
		lastResp = httpResp
		if err != nil {
			return nil, lastResp, fmt.Errorf("fetching dynamic attributes (offset %d): %w", offset, err)
		}

		// Capture msg/errorcode from every page so the final merged response
		// reflects the last page's status, not just the first.
		if page != nil {
			finalMsg = page.Msg
			finalECode = page.Errorcode
		}

		// No attributes returned — either we have everything or the endpoint has none.
		if page == nil ||
			page.Dynamicattributes == nil ||
			page.Dynamicattributes.ArrayOfFetchDynamicAttributeResponseInner == nil ||
			len(*page.Dynamicattributes.ArrayOfFetchDynamicAttributeResponseInner) == 0 {
			log.Printf("[DEBUG] FetchDynamicAttribute: empty page at offset %d, stopping pagination", offset)
			break
		}

		pageAttrs := *page.Dynamicattributes.ArrayOfFetchDynamicAttributeResponseInner
		allAttrs = append(allAttrs, pageAttrs...)
		log.Printf("[DEBUG] FetchDynamicAttribute: fetched %d attributes at offset %d (running total: %d)",
			len(pageAttrs), offset, len(allAttrs))

		// Stop if we have collected everything the API knows about.
		if page.Totalcount != nil && int32(len(allAttrs)) >= *page.Totalcount {
			log.Printf("[DEBUG] FetchDynamicAttribute: collected all %d attributes (totalcount=%d)", len(allAttrs), *page.Totalcount)
			break
		}

		// Stop if the page was smaller than a full page AND totalcount doesn't indicate
		// there are more — this handles APIs that omit totalcount.
		if len(pageAttrs) < pageSize && (page.Totalcount == nil || int32(len(allAttrs)) >= *page.Totalcount) {
			log.Printf("[DEBUG] FetchDynamicAttribute: partial page (%d < %d), stopping pagination", len(pageAttrs), pageSize)
			break
		}

		offset += len(pageAttrs)
	}

	// Build a merged response that looks like a single API call to callers.
	totalCount := int32(len(allAttrs))
	merged := &openapi.FetchDynamicAttributesResponse{
		Msg:          finalMsg,
		Errorcode:    finalECode,
		Displaycount: &totalCount,
		Totalcount:   &totalCount,
	}
	if len(allAttrs) > 0 {
		dynAttrs := openapi.FetchDynamicAttributesResponseDynamicattributes{
			ArrayOfFetchDynamicAttributeResponseInner: &allAttrs,
		}
		merged.Dynamicattributes = &dynAttrs
	}

	return merged, lastResp, nil
}

func (w *DynamicAttributeOperationsWrapper) FetchDynamicAttributesForDataSource(ctx context.Context, securitySystems, endpoints, dynamicAttributes, requestTypes []string, loggedInUser, offset, max *string) (*openapi.FetchDynamicAttributesResponse, *http.Response, error) {
	fetchReq := w.client.DynamicAttributesAPI.FetchDynamicAttribute(ctx)

	if securitySystems != nil {
		fetchReq = fetchReq.Securitysystem(securitySystems)
	}
	if endpoints != nil {
		fetchReq = fetchReq.Endpoint(endpoints)
	}
	if dynamicAttributes != nil {
		fetchReq = fetchReq.Dynamicattributes(dynamicAttributes)
	}
	if requestTypes != nil {
		fetchReq = fetchReq.Requesttype(requestTypes)
	}
	if offset != nil {
		fetchReq = fetchReq.Offset(*offset)
	}
	if max != nil {
		fetchReq = fetchReq.Max(*max)
	}
	if loggedInUser != nil {
		fetchReq = fetchReq.Loggedinuser(*loggedInUser)
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
	cfg.Servers = openapi.ServerConfigurations{{URL: baseURL}}
	cfg.AddDefaultHeader("Authorization", "Bearer "+token)
	cfg.HTTPClient = http.DefaultClient
	apiClient := openapi.NewAPIClient(cfg)
	return &DynamicAttributeOperationsWrapper{client: apiClient}
}

func (f *DefaultDynamicAttributeFactory) CreateEndpointOperations(baseURL, token string) EndpointOperationsInterface {
	cfg := endpoint.NewConfiguration()
	cfg.Servers = endpoint.ServerConfigurations{{URL: baseURL}}
	cfg.AddDefaultHeader("Authorization", "Bearer "+token)
	cfg.HTTPClient = http.DefaultClient
	apiClient := endpoint.NewAPIClient(cfg)
	return &EndpointOperationsWrapper{client: apiClient}
}
