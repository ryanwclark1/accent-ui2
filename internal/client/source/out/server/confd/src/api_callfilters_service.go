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

// CallfiltersAPIService is a service that implements the logic for the CallfiltersAPIServicer
// This service should implement the business logic for every endpoint for the CallfiltersAPI API.
// Include any external packages or services that will be required by this service.
type CallfiltersAPIService struct {
}

// NewCallfiltersAPIService creates a default api service
func NewCallfiltersAPIService() CallfiltersAPIServicer {
	return &CallfiltersAPIService{}
}

// CreateCallfilter - Create call filter
func (s *CallfiltersAPIService) CreateCallfilter(ctx context.Context, body CallFilter, accentTenant string) (ImplResponse, error) {
	// TODO - update CreateCallfilter with the required logic for this service method.
	// Add api_callfilters_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(201, CallFilter{}) or use other options such as http.Ok ...
	// return Response(201, CallFilter{}), nil

	// TODO: Uncomment the next line to return response Response(400, []string{}) or use other options such as http.Ok ...
	// return Response(400, []string{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("CreateCallfilter method not implemented")
}

// DeleteCallfilter - Delete call filter
func (s *CallfiltersAPIService) DeleteCallfilter(ctx context.Context, callfilterId int32, accentTenant string) (ImplResponse, error) {
	// TODO - update DeleteCallfilter with the required logic for this service method.
	// Add api_callfilters_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(204, {}) or use other options such as http.Ok ...
	// return Response(204, nil),nil

	// TODO: Uncomment the next line to return response Response(400, []string{}) or use other options such as http.Ok ...
	// return Response(400, []string{}), nil

	// TODO: Uncomment the next line to return response Response(404, []string{}) or use other options such as http.Ok ...
	// return Response(404, []string{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("DeleteCallfilter method not implemented")
}

// GetCallFilterFallback - List all fallbacks for call filter
func (s *CallfiltersAPIService) GetCallFilterFallback(ctx context.Context, callfilterId int32) (ImplResponse, error) {
	// TODO - update GetCallFilterFallback with the required logic for this service method.
	// Add api_callfilters_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(200, CallFilterFallbacks{}) or use other options such as http.Ok ...
	// return Response(200, CallFilterFallbacks{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("GetCallFilterFallback method not implemented")
}

// GetCallfilter - Get call filter
func (s *CallfiltersAPIService) GetCallfilter(ctx context.Context, callfilterId int32, accentTenant string) (ImplResponse, error) {
	// TODO - update GetCallfilter with the required logic for this service method.
	// Add api_callfilters_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(200, CallFilter{}) or use other options such as http.Ok ...
	// return Response(200, CallFilter{}), nil

	// TODO: Uncomment the next line to return response Response(404, []string{}) or use other options such as http.Ok ...
	// return Response(404, []string{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("GetCallfilter method not implemented")
}

// ListCallFilters - List call filters
func (s *CallfiltersAPIService) ListCallFilters(ctx context.Context, accentTenant string, recurse bool, order string, direction string, limit int32, offset int32, search string) (ImplResponse, error) {
	// TODO - update ListCallFilters with the required logic for this service method.
	// Add api_callfilters_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(200, CallFilterItems{}) or use other options such as http.Ok ...
	// return Response(200, CallFilterItems{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("ListCallFilters method not implemented")
}

// UpdateCallFilterCallerUsers - Update call filter and recipients
func (s *CallfiltersAPIService) UpdateCallFilterCallerUsers(ctx context.Context, callfilterId int32, body CallFilterRecipientUsersUuid) (ImplResponse, error) {
	// TODO - update UpdateCallFilterCallerUsers with the required logic for this service method.
	// Add api_callfilters_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(204, {}) or use other options such as http.Ok ...
	// return Response(204, nil),nil

	// TODO: Uncomment the next line to return response Response(400, []string{}) or use other options such as http.Ok ...
	// return Response(400, []string{}), nil

	// TODO: Uncomment the next line to return response Response(404, []string{}) or use other options such as http.Ok ...
	// return Response(404, []string{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("UpdateCallFilterCallerUsers method not implemented")
}

// UpdateCallFilterFallback - Update call filter&#39;s fallbacks
func (s *CallfiltersAPIService) UpdateCallFilterFallback(ctx context.Context, callfilterId int32, body CallFilterFallbacks) (ImplResponse, error) {
	// TODO - update UpdateCallFilterFallback with the required logic for this service method.
	// Add api_callfilters_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(204, {}) or use other options such as http.Ok ...
	// return Response(204, nil),nil

	// TODO: Uncomment the next line to return response Response(400, []string{}) or use other options such as http.Ok ...
	// return Response(400, []string{}), nil

	// TODO: Uncomment the next line to return response Response(404, []string{}) or use other options such as http.Ok ...
	// return Response(404, []string{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("UpdateCallFilterFallback method not implemented")
}

// UpdateCallFilterMemberUsers - Update call filter and surrogates
func (s *CallfiltersAPIService) UpdateCallFilterMemberUsers(ctx context.Context, callfilterId int32, body UsersUuid) (ImplResponse, error) {
	// TODO - update UpdateCallFilterMemberUsers with the required logic for this service method.
	// Add api_callfilters_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(204, {}) or use other options such as http.Ok ...
	// return Response(204, nil),nil

	// TODO: Uncomment the next line to return response Response(400, []string{}) or use other options such as http.Ok ...
	// return Response(400, []string{}), nil

	// TODO: Uncomment the next line to return response Response(404, []string{}) or use other options such as http.Ok ...
	// return Response(404, []string{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("UpdateCallFilterMemberUsers method not implemented")
}

// UpdateCallfilter - Update call filter
func (s *CallfiltersAPIService) UpdateCallfilter(ctx context.Context, callfilterId int32, body CallFilter, accentTenant string) (ImplResponse, error) {
	// TODO - update UpdateCallfilter with the required logic for this service method.
	// Add api_callfilters_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(204, {}) or use other options such as http.Ok ...
	// return Response(204, nil),nil

	// TODO: Uncomment the next line to return response Response(400, []string{}) or use other options such as http.Ok ...
	// return Response(400, []string{}), nil

	// TODO: Uncomment the next line to return response Response(404, []string{}) or use other options such as http.Ok ...
	// return Response(404, []string{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("UpdateCallfilter method not implemented")
}
