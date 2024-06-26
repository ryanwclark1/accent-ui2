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

// MeetingAuthorizationsAPIService is a service that implements the logic for the MeetingAuthorizationsAPIServicer
// This service should implement the business logic for every endpoint for the MeetingAuthorizationsAPI API.
// Include any external packages or services that will be required by this service.
type MeetingAuthorizationsAPIService struct {
}

// NewMeetingAuthorizationsAPIService creates a default api service
func NewMeetingAuthorizationsAPIService() MeetingAuthorizationsAPIServicer {
	return &MeetingAuthorizationsAPIService{}
}

// CreateGuestMeetingAuthorization - Request guest authorization to enter a meeting
func (s *MeetingAuthorizationsAPIService) CreateGuestMeetingAuthorization(ctx context.Context, guestUuid string, meetingUuid string, body MeetingAuthorizationRequest) (ImplResponse, error) {
	// TODO - update CreateGuestMeetingAuthorization with the required logic for this service method.
	// Add api_meeting_authorizations_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(201, MeetingAuthorization{}) or use other options such as http.Ok ...
	// return Response(201, MeetingAuthorization{}), nil

	// TODO: Uncomment the next line to return response Response(400, []string{}) or use other options such as http.Ok ...
	// return Response(400, []string{}), nil

	// TODO: Uncomment the next line to return response Response(404, []string{}) or use other options such as http.Ok ...
	// return Response(404, []string{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("CreateGuestMeetingAuthorization method not implemented")
}

// DeleteUserMeetingAuthorization - Delete the guest authorization to enter a meeting
func (s *MeetingAuthorizationsAPIService) DeleteUserMeetingAuthorization(ctx context.Context, guestUuid string, meetingUuid string, authorizationUuid string) (ImplResponse, error) {
	// TODO - update DeleteUserMeetingAuthorization with the required logic for this service method.
	// Add api_meeting_authorizations_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(204, {}) or use other options such as http.Ok ...
	// return Response(204, nil),nil

	// TODO: Uncomment the next line to return response Response(400, []string{}) or use other options such as http.Ok ...
	// return Response(400, []string{}), nil

	// TODO: Uncomment the next line to return response Response(404, []string{}) or use other options such as http.Ok ...
	// return Response(404, []string{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("DeleteUserMeetingAuthorization method not implemented")
}

// GetGuestMeetingAuthorization - Read the guest authorization to enter a meeting
func (s *MeetingAuthorizationsAPIService) GetGuestMeetingAuthorization(ctx context.Context, guestUuid string, meetingUuid string, authorizationUuid string) (ImplResponse, error) {
	// TODO - update GetGuestMeetingAuthorization with the required logic for this service method.
	// Add api_meeting_authorizations_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(200, MeetingAuthorization{}) or use other options such as http.Ok ...
	// return Response(200, MeetingAuthorization{}), nil

	// TODO: Uncomment the next line to return response Response(400, []string{}) or use other options such as http.Ok ...
	// return Response(400, []string{}), nil

	// TODO: Uncomment the next line to return response Response(404, []string{}) or use other options such as http.Ok ...
	// return Response(404, []string{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("GetGuestMeetingAuthorization method not implemented")
}

// GetUserMeetingAuthorization - Read the guest authorization to enter a meeting
func (s *MeetingAuthorizationsAPIService) GetUserMeetingAuthorization(ctx context.Context, guestUuid string, meetingUuid string, authorizationUuid string) (ImplResponse, error) {
	// TODO - update GetUserMeetingAuthorization with the required logic for this service method.
	// Add api_meeting_authorizations_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(200, MeetingAuthorization{}) or use other options such as http.Ok ...
	// return Response(200, MeetingAuthorization{}), nil

	// TODO: Uncomment the next line to return response Response(400, []string{}) or use other options such as http.Ok ...
	// return Response(400, []string{}), nil

	// TODO: Uncomment the next line to return response Response(404, []string{}) or use other options such as http.Ok ...
	// return Response(404, []string{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("GetUserMeetingAuthorization method not implemented")
}

// ListUserMeetingAuthorizations - List all guest authorization requests of a meeting
func (s *MeetingAuthorizationsAPIService) ListUserMeetingAuthorizations(ctx context.Context, meetingUuid string) (ImplResponse, error) {
	// TODO - update ListUserMeetingAuthorizations with the required logic for this service method.
	// Add api_meeting_authorizations_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(200, MeetingAuthorizationItems{}) or use other options such as http.Ok ...
	// return Response(200, MeetingAuthorizationItems{}), nil

	// TODO: Uncomment the next line to return response Response(404, []string{}) or use other options such as http.Ok ...
	// return Response(404, []string{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("ListUserMeetingAuthorizations method not implemented")
}

// PutUserMeetingAuthorizationAccept - Accept a guest authorization request
func (s *MeetingAuthorizationsAPIService) PutUserMeetingAuthorizationAccept(ctx context.Context, meetingUuid string, authorizationUuid string) (ImplResponse, error) {
	// TODO - update PutUserMeetingAuthorizationAccept with the required logic for this service method.
	// Add api_meeting_authorizations_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(200, MeetingAuthorization{}) or use other options such as http.Ok ...
	// return Response(200, MeetingAuthorization{}), nil

	// TODO: Uncomment the next line to return response Response(400, []string{}) or use other options such as http.Ok ...
	// return Response(400, []string{}), nil

	// TODO: Uncomment the next line to return response Response(404, []string{}) or use other options such as http.Ok ...
	// return Response(404, []string{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("PutUserMeetingAuthorizationAccept method not implemented")
}

// PutUserMeetingAuthorizationReject - Reject a guest authorization request
func (s *MeetingAuthorizationsAPIService) PutUserMeetingAuthorizationReject(ctx context.Context, meetingUuid string, authorizationUuid string) (ImplResponse, error) {
	// TODO - update PutUserMeetingAuthorizationReject with the required logic for this service method.
	// Add api_meeting_authorizations_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(200, MeetingAuthorization{}) or use other options such as http.Ok ...
	// return Response(200, MeetingAuthorization{}), nil

	// TODO: Uncomment the next line to return response Response(400, []string{}) or use other options such as http.Ok ...
	// return Response(400, []string{}), nil

	// TODO: Uncomment the next line to return response Response(404, []string{}) or use other options such as http.Ok ...
	// return Response(404, []string{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("PutUserMeetingAuthorizationReject method not implemented")
}
