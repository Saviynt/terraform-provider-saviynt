// Copyright (c) 2025 Saviynt Inc.
// SPDX-License-Identifier: MPL-2.0

package client

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	openapi "github.com/saviynt/saviynt-api-go-client/filedirectory"
)

// FileOperationsInterface defines the interface for file operations
type FileOperationsInterface interface {
	UploadSchemaFile(ctx context.Context, file *os.File, pathLocation string) (*openapi.UploadSchemaFileResponse, *http.Response, error)
}

// FileOperationsWrapper wraps the actual file operations to implement the interface
type FileOperationsWrapper struct {
	client *openapi.APIClient
}

func (w *FileOperationsWrapper) UploadSchemaFile(ctx context.Context, file *os.File, pathLocation string) (*openapi.UploadSchemaFileResponse, *http.Response, error) {
	// Validate file extension and set appropriate path location
	fileName := file.Name()
	ext := strings.ToLower(filepath.Ext(fileName))

	var validatedPathLocation string
	switch ext {
	case ".csv":
		validatedPathLocation = "Datafiles"
	case ".sav":
		validatedPathLocation = "SAV"
	default:
		return nil, nil, fmt.Errorf("unsupported file extension '%s'. Only .csv and .sav files are allowed", ext)
	}

	// Override pathLocation if provided but validate it matches expected value
	if pathLocation != "" && pathLocation != validatedPathLocation {
		return nil, nil, fmt.Errorf("invalid path_location '%s' for file extension '%s'. Expected '%s'", pathLocation, ext, validatedPathLocation)
	}

	req := w.client.FileDirectoryAPI.UploadNewFile(ctx).File(file).PathLocation(validatedPathLocation)
	return req.Execute()
}

// FileFactoryInterface defines the interface for creating file operations
type FileFactoryInterface interface {
	CreateFileOperations(baseURL, token string) FileOperationsInterface
}

// DefaultFileFactory implements the FileFactoryInterface
type DefaultFileFactory struct {
	ProviderConfig interface{}
}

func (f *DefaultFileFactory) CreateFileOperations(baseURL, token string) FileOperationsInterface {
	cfg := openapi.NewConfiguration()
	apiBaseURL := strings.TrimPrefix(strings.TrimPrefix(baseURL, "https://"), "http://")
	cfg.Host = apiBaseURL
	cfg.Scheme = "https"
	cfg.AddDefaultHeader("Authorization", "Bearer "+token)
	cfg.HTTPClient = http.DefaultClient
	apiClient := openapi.NewAPIClient(cfg)
	return &FileOperationsWrapper{client: apiClient}
}
