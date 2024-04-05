/*
 * accent-amid
 *
 * Send AMI actions to Asterisk, providing token based authentication.
 *
 * API version: 0.1
 * Contact: help@accentvoice.io
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package amidserver

import (
	"context"
	"errors"
	"net/http"
)

// ActionAPIService is a service that implements the logic for the ActionAPIServicer
// This service should implement the business logic for every endpoint for the ActionAPI API.
// Include any external packages or services that will be required by this service.
type ActionAPIService struct {
}

// NewActionAPIService creates a default api service
func NewActionAPIService() ActionAPIServicer {
	return &ActionAPIService{}
}

// ActionAsteriskManager - AMI action
func (s *ActionAPIService) ActionAsteriskManager(ctx context.Context, action string, actionArguments map[string]interface{}) (ImplResponse, error) {
	// TODO - update ActionAsteriskManager with the required logic for this service method.
	// Add api_action_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(200, Response{}) or use other options such as http.Ok ...
	// return Response(200, Response{}), nil

	// TODO: Uncomment the next line to return response Response(401, Error{}) or use other options such as http.Ok ...
	// return Response(401, Error{}), nil

	// TODO: Uncomment the next line to return response Response(503, Error{}) or use other options such as http.Ok ...
	// return Response(503, Error{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("ActionAsteriskManager method not implemented")
}