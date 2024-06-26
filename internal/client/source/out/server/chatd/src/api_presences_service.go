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

// PresencesAPIService is a service that implements the logic for the PresencesAPIServicer
// This service should implement the business logic for every endpoint for the PresencesAPI API.
// Include any external packages or services that will be required by this service.
type PresencesAPIService struct {
}

// NewPresencesAPIService creates a default api service
func NewPresencesAPIService() PresencesAPIServicer {
	return &PresencesAPIService{}
}

// GetUserPresence - Get user presence
func (s *PresencesAPIService) GetUserPresence(ctx context.Context, userUuid string, accentTenant string) (ImplResponse, error) {
	// TODO - update GetUserPresence with the required logic for this service method.
	// Add api_presences_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(200, Presence{}) or use other options such as http.Ok ...
	// return Response(200, Presence{}), nil

	// TODO: Uncomment the next line to return response Response(404, ApiError{}) or use other options such as http.Ok ...
	// return Response(404, ApiError{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("GetUserPresence method not implemented")
}

// ListPresences - List presences
func (s *PresencesAPIService) ListPresences(ctx context.Context, accentTenant string, recurse bool, userUuid []string) (ImplResponse, error) {
	// TODO - update ListPresences with the required logic for this service method.
	// Add api_presences_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(200, PresenceList{}) or use other options such as http.Ok ...
	// return Response(200, PresenceList{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("ListPresences method not implemented")
}

// UpdateUserPresence - Update user presence
func (s *PresencesAPIService) UpdateUserPresence(ctx context.Context, userUuid string, body Presence, accentTenant string) (ImplResponse, error) {
	// TODO - update UpdateUserPresence with the required logic for this service method.
	// Add api_presences_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(204, {}) or use other options such as http.Ok ...
	// return Response(204, nil),nil

	// TODO: Uncomment the next line to return response Response(400, ApiError{}) or use other options such as http.Ok ...
	// return Response(400, ApiError{}), nil

	// TODO: Uncomment the next line to return response Response(404, ApiError{}) or use other options such as http.Ok ...
	// return Response(404, ApiError{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("UpdateUserPresence method not implemented")
}
