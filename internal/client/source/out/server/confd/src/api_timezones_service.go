/*
 * accent-confd
 *
 * Confd exposes an API for managing core resources on a Accent server such as users, extensions, devices, voicemails, queues, etc. Resources can be associated together to provide additional functionality. For example: By associating a voicemail with a user, calls will automatically fallback on to the voicemail when the user cannot answer.  Implementation notes ====================  Errors ------  Responses containing errors will have a status code in the 400 or 500 class. A list of error messages will be returned in the body of the response as a JSON-encoded array:  ~~~ [     \"Input error - User not found\",     \"Resource error - User not associated to a line\" ] ~~~   Updating resources via PUT --------------------------  When updating a resource, all fields become optional. In other words, only values that have been changed need to be sent to the server. Please note that this behavior may change in future versions of the API.
 *
 * API version: 1.1
 * Contact: help@accentvoice.io
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package confdserver

import (
	"context"
	"errors"
	"net/http"
)

// TimezonesAPIService is a service that implements the logic for the TimezonesAPIServicer
// This service should implement the business logic for every endpoint for the TimezonesAPI API.
// Include any external packages or services that will be required by this service.
type TimezonesAPIService struct {
}

// NewTimezonesAPIService creates a default api service
func NewTimezonesAPIService() TimezonesAPIServicer {
	return &TimezonesAPIService{}
}

// ListTimezones - List all available timezones
func (s *TimezonesAPIService) ListTimezones(ctx context.Context) (ImplResponse, error) {
	// TODO - update ListTimezones with the required logic for this service method.
	// Add api_timezones_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(200, TimezoneItems{}) or use other options such as http.Ok ...
	// return Response(200, TimezoneItems{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("ListTimezones method not implemented")
}