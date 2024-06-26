/*
 * accent-calld
 *
 * Control your calls from a REST API
 *
 * API version: 1.0.0
 * Contact: help@accentvoice.io
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package calldserver

import (
	"context"
	"errors"
	"net/http"
	"os"
)

// FaxesAPIService is a service that implements the logic for the FaxesAPIServicer
// This service should implement the business logic for every endpoint for the FaxesAPI API.
// Include any external packages or services that will be required by this service.
type FaxesAPIService struct {
}

// NewFaxesAPIService creates a default api service
func NewFaxesAPIService() FaxesAPIServicer {
	return &FaxesAPIService{}
}

// SendFax - Send a fax
func (s *FaxesAPIService) SendFax(ctx context.Context, context string, extension string, faxContent *os.File, callerId string, ivrExtension string, waitTime int32) (ImplResponse, error) {
	// TODO - update SendFax with the required logic for this service method.
	// Add api_faxes_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(201, Fax{}) or use other options such as http.Ok ...
	// return Response(201, Fax{}), nil

	// TODO: Uncomment the next line to return response Response(400, Error{}) or use other options such as http.Ok ...
	// return Response(400, Error{}), nil

	// TODO: Uncomment the next line to return response Response(503, Error{}) or use other options such as http.Ok ...
	// return Response(503, Error{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("SendFax method not implemented")
}

// SendUserFax - Send a fax as the user detected from the token
func (s *FaxesAPIService) SendUserFax(ctx context.Context, extension string, faxContent *os.File, callerId string, ivrExtension string, waitTime int32) (ImplResponse, error) {
	// TODO - update SendUserFax with the required logic for this service method.
	// Add api_faxes_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(201, Fax{}) or use other options such as http.Ok ...
	// return Response(201, Fax{}), nil

	// TODO: Uncomment the next line to return response Response(400, Error{}) or use other options such as http.Ok ...
	// return Response(400, Error{}), nil

	// TODO: Uncomment the next line to return response Response(503, Error{}) or use other options such as http.Ok ...
	// return Response(503, Error{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("SendUserFax method not implemented")
}
