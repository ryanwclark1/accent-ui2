/*
 * accent-chatd
 *
 * Control your message and presence from a REST API
 *
 * API version: 1.0.0
 * Contact: help@accentvoice.io
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package chatdserver

import (
	"context"
	"errors"
	"net/http"
)

// MicrosoftAPIService is a service that implements the logic for the MicrosoftAPIServicer
// This service should implement the business logic for every endpoint for the MicrosoftAPI API.
// Include any external packages or services that will be required by this service.
type MicrosoftAPIService struct {
}

// NewMicrosoftAPIService creates a default api service
func NewMicrosoftAPIService() MicrosoftAPIServicer {
	return &MicrosoftAPIService{}
}

// UpdateTeamsPresence - Receive presence information from Microsoft Teams
func (s *MicrosoftAPIService) UpdateTeamsPresence(ctx context.Context, userUuid string) (ImplResponse, error) {
	// TODO - update UpdateTeamsPresence with the required logic for this service method.
	// Add api_microsoft_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(200, Object{}) or use other options such as http.Ok ...
	// return Response(200, Object{}), nil

	// TODO: Uncomment the next line to return response Response(404, Object{}) or use other options such as http.Ok ...
	// return Response(404, Object{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("UpdateTeamsPresence method not implemented")
}