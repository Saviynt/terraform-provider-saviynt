// Copyright (c) 2025 Saviynt Inc.
// SPDX-License-Identifier: MPL-2.0

// auth_retry_helper.go provides enhanced authentication retry logic for handling
// 401 errors with automatic token refresh and limited retry attempts.

package provider

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strings"

	openapi "github.com/saviynt/saviynt-api-go-client/utility"
)

// makeAuthenticatedRequestWithRetry handles API calls with 401 retry logic only
func (p *SaviyntProvider) makeAuthenticatedRequestWithRetry(ctx context.Context, requestFunc func(token string) error) error {
	// Use current token (no proactive refresh)
	p.tokenMutex.RLock()
	token := p.accessToken
	p.tokenMutex.RUnlock()

	err := requestFunc(token)

	// Only refresh token if we get 401 error
	if err != nil && is401Error(err) {
		return p.retryWithTokenRefresh(ctx, requestFunc, 2)
	}

	return err
}

// retryWithTokenRefresh attempts token refresh and retry up to maxRetries times for 401 errors
func (p *SaviyntProvider) retryWithTokenRefresh(ctx context.Context, requestFunc func(token string) error, maxRetries int) error {
	for attempt := 1; attempt <= maxRetries; attempt++ {
		log.Printf("[DEBUG] Received 401 error, attempting token refresh (attempt %d/%d)...", attempt, maxRetries)

		// Call refresh token API with grant_type and refresh_token
		if refreshErr := p.callRefreshTokenAPI(ctx); refreshErr != nil {
			return fmt.Errorf("authentication failed and token refresh failed on attempt %d: %w", attempt, refreshErr)
		}

		// Get the new token and retry
		p.tokenMutex.RLock()
		newToken := p.accessToken
		p.tokenMutex.RUnlock()

		err := requestFunc(newToken)
		if err == nil {
			log.Printf("[DEBUG] Request succeeded after token refresh attempt %d", attempt)
			return nil
		}

		// Check if it's still a 401 error
		if !is401Error(err) {
			return err
		}

		if attempt == maxRetries {
			return fmt.Errorf("authentication failed after %d token refresh attempts: %w", maxRetries, err)
		}
	}

	return fmt.Errorf("unexpected error in retry logic")
}

// callRefreshTokenAPI calls the utility API with grant_type and refresh_token
func (p *SaviyntProvider) callRefreshTokenAPI(ctx context.Context) error {
	p.tokenMutex.Lock()
	defer p.tokenMutex.Unlock()

	log.Printf("[DEBUG] Calling refresh token API...")

	// Create API configuration
	cfg := openapi.NewConfiguration()
	apiBaseURL := strings.TrimPrefix(strings.TrimPrefix(p.client.APIBaseURL(), "https://"), "http://")
	cfg.Host = apiBaseURL
	cfg.Scheme = "https"
	cfg.HTTPClient = &http.Client{}

	apiClient := openapi.NewAPIClient(cfg)

	// Make refresh token request with grant_type and refresh_token
	refreshReq := apiClient.UtilityAPI.AccessToken(ctx).
		GrantType("refresh_token").
		RefreshToken(p.refreshToken)

	tokenResp, _, err := refreshReq.Execute()
	if err != nil {
		log.Printf("[ERROR] Failed to refresh access token: %v", err)
		return fmt.Errorf("failed to refresh access token: %w", err)
	}

	if tokenResp.AccessToken == nil || *tokenResp.AccessToken == "" {
		return fmt.Errorf("received empty access token from refresh")
	}

	// Update token information (no expiry tracking needed)
	p.accessToken = *tokenResp.AccessToken
	if tokenResp.RefreshToken != nil && *tokenResp.RefreshToken != "" {
		p.refreshToken = *tokenResp.RefreshToken
	}

	log.Printf("[DEBUG] Access token refreshed successfully")
	return nil
}

// is401Error checks if the error indicates a 401 status
func is401Error(err error) bool {
	if err == nil {
		return false
	}
	errStr := strings.ToLower(err.Error())
	return strings.Contains(errStr, "401")
}

// AuthenticatedAPICallWithRetry is the main function resources should use for API calls with retry logic
func (p *SaviyntProvider) AuthenticatedAPICallWithRetry(ctx context.Context, operation string, apiCall func(token string) error) error {
	log.Printf("[DEBUG] Making authenticated API call with retry for operation: %s", operation)

	return p.makeAuthenticatedRequestWithRetry(ctx, apiCall)
}
