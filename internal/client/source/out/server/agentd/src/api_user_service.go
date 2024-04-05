/*
 * accent-agentd
 *
 * Agent service
 *
 * API version: 1.0
 * Contact: help@accentvoice.io
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package agentdserver

import (
	"context"
	"errors"
	"net/http"
)

// UserAPIService is a service that implements the logic for the UserAPIServicer
// This service should implement the business logic for every endpoint for the UserAPI API.
// Include any external packages or services that will be required by this service.
type UserAPIService struct {
}

// NewUserAPIService creates a default api service
func NewUserAPIService() UserAPIServicer {
	return &UserAPIService{}
}

// GetUserAgent - Get agent status of the user holding the authentication token.
func (s *UserAPIService) GetUserAgent(ctx context.Context, accentTenant string) (ImplResponse, error) {
	// TODO - update GetUserAgent with the required logic for this service method.
	// Add api_user_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(200, AgentStatus{}) or use other options such as http.Ok ...
	// return Response(200, AgentStatus{}), nil

	// TODO: Uncomment the next line to return response Response(404, Error{}) or use other options such as http.Ok ...
	// return Response(404, Error{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("GetUserAgent method not implemented")
}

// LoginUserAgent - Log the agent of the user holding the authentication token
func (s *UserAPIService) LoginUserAgent(ctx context.Context, body UserAgentLoginInfo, accentTenant string) (ImplResponse, error) {
	// TODO - update LoginUserAgent with the required logic for this service method.
	// Add api_user_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(204, {}) or use other options such as http.Ok ...
	// return Response(204, nil),nil

	// TODO: Uncomment the next line to return response Response(400, Error{}) or use other options such as http.Ok ...
	// return Response(400, Error{}), nil

	// TODO: Uncomment the next line to return response Response(409, Error{}) or use other options such as http.Ok ...
	// return Response(409, Error{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("LoginUserAgent method not implemented")
}

// LogoffUserAgent - Logoff the agent of the user holding the authentication token
func (s *UserAPIService) LogoffUserAgent(ctx context.Context, accentTenant string) (ImplResponse, error) {
	// TODO - update LogoffUserAgent with the required logic for this service method.
	// Add api_user_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(204, {}) or use other options such as http.Ok ...
	// return Response(204, nil),nil

	// TODO: Uncomment the next line to return response Response(400, Error{}) or use other options such as http.Ok ...
	// return Response(400, Error{}), nil

	// TODO: Uncomment the next line to return response Response(409, Error{}) or use other options such as http.Ok ...
	// return Response(409, Error{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("LogoffUserAgent method not implemented")
}

// PauseUserAgent - Pause the agent of the user holding the authentication token
func (s *UserAPIService) PauseUserAgent(ctx context.Context, accentTenant string, body AgentPauseReason) (ImplResponse, error) {
	// TODO - update PauseUserAgent with the required logic for this service method.
	// Add api_user_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(204, {}) or use other options such as http.Ok ...
	// return Response(204, nil),nil

	// TODO: Uncomment the next line to return response Response(400, Error{}) or use other options such as http.Ok ...
	// return Response(400, Error{}), nil

	// TODO: Uncomment the next line to return response Response(409, Error{}) or use other options such as http.Ok ...
	// return Response(409, Error{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("PauseUserAgent method not implemented")
}

// UnpauseUserAgent - Unpause the agent of the user holding the authentication token
func (s *UserAPIService) UnpauseUserAgent(ctx context.Context, accentTenant string) (ImplResponse, error) {
	// TODO - update UnpauseUserAgent with the required logic for this service method.
	// Add api_user_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(204, {}) or use other options such as http.Ok ...
	// return Response(204, nil),nil

	// TODO: Uncomment the next line to return response Response(400, Error{}) or use other options such as http.Ok ...
	// return Response(400, Error{}), nil

	// TODO: Uncomment the next line to return response Response(409, Error{}) or use other options such as http.Ok ...
	// return Response(409, Error{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("UnpauseUserAgent method not implemented")
}