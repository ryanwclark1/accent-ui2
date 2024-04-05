/*
 * accent-setupd
 *
 * Initialize Accent Engine from a REST API
 *
 * API version: 1.0.0
 * Contact: help@accentvoice.io
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package setupdserver

import (
	"context"
	"errors"
	"net/http"
)

// StatusAPIService is a service that implements the logic for the StatusAPIServicer
// This service should implement the business logic for every endpoint for the StatusAPI API.
// Include any external packages or services that will be required by this service.
type StatusAPIService struct {
}

// NewStatusAPIService creates a default api service
func NewStatusAPIService() StatusAPIServicer {
	return &StatusAPIService{}
}

// GetStatus - Print infos about internal status of accent-setupd
func (s *StatusAPIService) GetStatus(ctx context.Context) (ImplResponse, error) {
	// TODO - update GetStatus with the required logic for this service method.
	// Add api_status_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(200, StatusSummary{}) or use other options such as http.Ok ...
	// return Response(200, StatusSummary{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("GetStatus method not implemented")
}
