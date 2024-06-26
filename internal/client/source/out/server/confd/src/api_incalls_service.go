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

// IncallsAPIService is a service that implements the logic for the IncallsAPIServicer
// This service should implement the business logic for every endpoint for the IncallsAPI API.
// Include any external packages or services that will be required by this service.
type IncallsAPIService struct {
}

// NewIncallsAPIService creates a default api service
func NewIncallsAPIService() IncallsAPIServicer {
	return &IncallsAPIService{}
}

// AssociateIncallExtension - Associate incall and extension
func (s *IncallsAPIService) AssociateIncallExtension(ctx context.Context, incallId int32, extensionId int32) (ImplResponse, error) {
	// TODO - update AssociateIncallExtension with the required logic for this service method.
	// Add api_incalls_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(204, {}) or use other options such as http.Ok ...
	// return Response(204, nil),nil

	// TODO: Uncomment the next line to return response Response(400, []string{}) or use other options such as http.Ok ...
	// return Response(400, []string{}), nil

	// TODO: Uncomment the next line to return response Response(404, []string{}) or use other options such as http.Ok ...
	// return Response(404, []string{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("AssociateIncallExtension method not implemented")
}

// AssociateIncallSchedule - Associate incall and schedule
func (s *IncallsAPIService) AssociateIncallSchedule(ctx context.Context, incallId int32, scheduleId int32, accentTenant string) (ImplResponse, error) {
	// TODO - update AssociateIncallSchedule with the required logic for this service method.
	// Add api_incalls_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(204, {}) or use other options such as http.Ok ...
	// return Response(204, nil),nil

	// TODO: Uncomment the next line to return response Response(400, []string{}) or use other options such as http.Ok ...
	// return Response(400, []string{}), nil

	// TODO: Uncomment the next line to return response Response(404, []string{}) or use other options such as http.Ok ...
	// return Response(404, []string{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("AssociateIncallSchedule method not implemented")
}

// CreateIncall - Create incoming call
func (s *IncallsAPIService) CreateIncall(ctx context.Context, body Incall, accentTenant string) (ImplResponse, error) {
	// TODO - update CreateIncall with the required logic for this service method.
	// Add api_incalls_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(201, Incall{}) or use other options such as http.Ok ...
	// return Response(201, Incall{}), nil

	// TODO: Uncomment the next line to return response Response(400, []string{}) or use other options such as http.Ok ...
	// return Response(400, []string{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("CreateIncall method not implemented")
}

// DeleteIncall - Delete incoming call
func (s *IncallsAPIService) DeleteIncall(ctx context.Context, incallId int32, accentTenant string) (ImplResponse, error) {
	// TODO - update DeleteIncall with the required logic for this service method.
	// Add api_incalls_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(204, {}) or use other options such as http.Ok ...
	// return Response(204, nil),nil

	// TODO: Uncomment the next line to return response Response(400, []string{}) or use other options such as http.Ok ...
	// return Response(400, []string{}), nil

	// TODO: Uncomment the next line to return response Response(404, []string{}) or use other options such as http.Ok ...
	// return Response(404, []string{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("DeleteIncall method not implemented")
}

// DissociateIncallExtension - Dissociate incall and extension
func (s *IncallsAPIService) DissociateIncallExtension(ctx context.Context, incallId int32, extensionId int32) (ImplResponse, error) {
	// TODO - update DissociateIncallExtension with the required logic for this service method.
	// Add api_incalls_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(204, {}) or use other options such as http.Ok ...
	// return Response(204, nil),nil

	// TODO: Uncomment the next line to return response Response(400, []string{}) or use other options such as http.Ok ...
	// return Response(400, []string{}), nil

	// TODO: Uncomment the next line to return response Response(404, []string{}) or use other options such as http.Ok ...
	// return Response(404, []string{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("DissociateIncallExtension method not implemented")
}

// DissociateIncallSchedule - Dissociate incall and schedule
func (s *IncallsAPIService) DissociateIncallSchedule(ctx context.Context, incallId int32, scheduleId int32, accentTenant string) (ImplResponse, error) {
	// TODO - update DissociateIncallSchedule with the required logic for this service method.
	// Add api_incalls_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(204, {}) or use other options such as http.Ok ...
	// return Response(204, nil),nil

	// TODO: Uncomment the next line to return response Response(400, []string{}) or use other options such as http.Ok ...
	// return Response(400, []string{}), nil

	// TODO: Uncomment the next line to return response Response(404, []string{}) or use other options such as http.Ok ...
	// return Response(404, []string{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("DissociateIncallSchedule method not implemented")
}

// GetIncall - Get incoming call
func (s *IncallsAPIService) GetIncall(ctx context.Context, incallId int32, accentTenant string) (ImplResponse, error) {
	// TODO - update GetIncall with the required logic for this service method.
	// Add api_incalls_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(200, Incall{}) or use other options such as http.Ok ...
	// return Response(200, Incall{}), nil

	// TODO: Uncomment the next line to return response Response(404, []string{}) or use other options such as http.Ok ...
	// return Response(404, []string{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("GetIncall method not implemented")
}

// ListIncalls - List incoming calls
func (s *IncallsAPIService) ListIncalls(ctx context.Context, accentTenant string, recurse bool, order string, direction string, limit int32, offset int32, search string) (ImplResponse, error) {
	// TODO - update ListIncalls with the required logic for this service method.
	// Add api_incalls_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(200, IncallItems{}) or use other options such as http.Ok ...
	// return Response(200, IncallItems{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("ListIncalls method not implemented")
}

// UpdateIncall - Update incoming call
func (s *IncallsAPIService) UpdateIncall(ctx context.Context, incallId int32, body Incall, accentTenant string) (ImplResponse, error) {
	// TODO - update UpdateIncall with the required logic for this service method.
	// Add api_incalls_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(204, {}) or use other options such as http.Ok ...
	// return Response(204, nil),nil

	// TODO: Uncomment the next line to return response Response(400, []string{}) or use other options such as http.Ok ...
	// return Response(400, []string{}), nil

	// TODO: Uncomment the next line to return response Response(404, []string{}) or use other options such as http.Ok ...
	// return Response(404, []string{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("UpdateIncall method not implemented")
}
