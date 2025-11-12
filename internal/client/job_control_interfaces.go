// Copyright (c) 2025 Saviynt Inc.
// SPDX-License-Identifier: MPL-2.0

package client

import (
	"context"
	"net/http"
	"strings"

	openapi "github.com/saviynt/saviynt-api-go-client/job_control"
)

// JobControlOperationsInterface defines the interface for job control operations
// This interface is used by job control resources for dependency injection
type JobControlOperationsInterface interface {
	CreateOrUpdateTriggers(ctx context.Context, req openapi.CreateOrUpdateTriggersRequest) (*openapi.CreateOrUpdateTriggersResponse, *http.Response, error)
	CreateTrigger(ctx context.Context, req []openapi.JobTriggerRequest) (*openapi.CreateTriggersResponse, *http.Response, error)
	DeleteTrigger(ctx context.Context, req openapi.DeleteTriggerRequest) (*openapi.DeleteTriggerResponse, *http.Response, error)
	RunJobTrigger(ctx context.Context, req openapi.RunJobTriggerRequest) (*openapi.RunJobTriggerResponse, *http.Response, error)
	CheckJobStatus(ctx context.Context, req openapi.CheckJobStatusRequest) (*openapi.CheckJobStatusResponse, *http.Response, error)
	PauseResumeJobs(ctx context.Context, req openapi.PauseResumeJobsRequest) (string, *http.Response, error)
	PauseAllJobs(ctx context.Context) (*openapi.PauseResumeJobsResponse, *http.Response, error)
	ResumeAllJobs(ctx context.Context) (*openapi.PauseResumeJobsResponse, *http.Response, error)
	PauseJob(ctx context.Context, req openapi.PauseResumeJobRequest) (*openapi.PauseResumeJobsResponse, *http.Response, error)
	ResumeJob(ctx context.Context, req openapi.PauseResumeJobRequest) (*openapi.PauseResumeJobsResponse, *http.Response, error)
	FetchJobMetadata(ctx context.Context, req openapi.FetchJobMetadataRequest) (*openapi.FetchJobMetadataResponse, *http.Response, error)
}

// JobControlOperationsWrapper wraps the actual job control operations to implement the interface
type JobControlOperationsWrapper struct {
	client *openapi.APIClient
}

func (w *JobControlOperationsWrapper) CreateOrUpdateTriggers(ctx context.Context, req openapi.CreateOrUpdateTriggersRequest) (*openapi.CreateOrUpdateTriggersResponse, *http.Response, error) {
	return w.client.JobControlAPI.CreateOrUpdateTrigger(ctx).CreateOrUpdateTriggersRequest(req).Execute()
}

func (w *JobControlOperationsWrapper) DeleteTrigger(ctx context.Context, req openapi.DeleteTriggerRequest) (*openapi.DeleteTriggerResponse, *http.Response, error) {
	return w.client.JobControlAPI.DeleteTrigger(ctx).DeleteTriggerRequest(req).Execute()
}

func (w *JobControlOperationsWrapper) RunJobTrigger(ctx context.Context, req openapi.RunJobTriggerRequest) (*openapi.RunJobTriggerResponse, *http.Response, error) {
	return w.client.JobControlAPI.RunJobTrigger(ctx).RunJobTriggerRequest(req).Execute()
}

func (w *JobControlOperationsWrapper) CheckJobStatus(ctx context.Context, req openapi.CheckJobStatusRequest) (*openapi.CheckJobStatusResponse, *http.Response, error) {
	return w.client.JobControlAPI.CheckJobStatus(ctx).CheckJobStatusRequest(req).Execute()
}

func (w *JobControlOperationsWrapper) PauseResumeJobs(ctx context.Context, req openapi.PauseResumeJobsRequest) (string, *http.Response, error) {
	return w.client.JobControlAPI.PauseResumeJobs(ctx).PauseResumeJobsRequest(req).Execute()
}

func (w *JobControlOperationsWrapper) CreateTrigger(ctx context.Context, req []openapi.JobTriggerRequest) (*openapi.CreateTriggersResponse, *http.Response, error) {
	return w.client.JobControlAPI.CreateTrigger(ctx).JobTriggerRequest(req).Execute()
}

func (w *JobControlOperationsWrapper) PauseAllJobs(ctx context.Context) (*openapi.PauseResumeJobsResponse, *http.Response, error) {
	return w.client.JobControlAPI.PauseAllJobs(ctx).Execute()
}

func (w *JobControlOperationsWrapper) ResumeAllJobs(ctx context.Context) (*openapi.PauseResumeJobsResponse, *http.Response, error) {
	return w.client.JobControlAPI.ResumeAllJobs(ctx).Execute()
}

func (w *JobControlOperationsWrapper) PauseJob(ctx context.Context, req openapi.PauseResumeJobRequest) (*openapi.PauseResumeJobsResponse, *http.Response, error) {
	return w.client.JobControlAPI.PauseJob(ctx).PauseResumeJobRequest(req).Execute()
}

func (w *JobControlOperationsWrapper) ResumeJob(ctx context.Context, req openapi.PauseResumeJobRequest) (*openapi.PauseResumeJobsResponse, *http.Response, error) {
	return w.client.JobControlAPI.ResumeJob(ctx).PauseResumeJobRequest(req).Execute()
}

func (w *JobControlOperationsWrapper) FetchJobMetadata(ctx context.Context, req openapi.FetchJobMetadataRequest) (*openapi.FetchJobMetadataResponse, *http.Response, error) {
	return w.client.JobControlAPI.FetchJobMetadata(ctx).FetchJobMetadataRequest(req).Execute()
}

// JobControlFactoryInterface defines the interface for creating job control operations
// This factory is used by job control resources for dependency injection
type JobControlFactoryInterface interface {
	CreateJobControlOperations(baseURL, token string) JobControlOperationsInterface
}

// DefaultJobControlFactory implements the JobControlFactoryInterface
type DefaultJobControlFactory struct{}

func (f *DefaultJobControlFactory) CreateJobControlOperations(baseURL, token string) JobControlOperationsInterface {
	cfg := openapi.NewConfiguration()
	apiBaseURL := strings.TrimPrefix(strings.TrimPrefix(baseURL, "https://"), "http://")
	cfg.Host = apiBaseURL
	cfg.Scheme = "https"
	cfg.AddDefaultHeader("Authorization", "Bearer "+token)
	cfg.HTTPClient = http.DefaultClient
	apiClient := openapi.NewAPIClient(cfg)
	return &JobControlOperationsWrapper{client: apiClient}
}
