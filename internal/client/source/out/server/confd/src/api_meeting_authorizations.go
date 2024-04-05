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
	"encoding/json"
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5"
)

// MeetingAuthorizationsAPIController binds http requests to an api service and writes the service results to the http response
type MeetingAuthorizationsAPIController struct {
	service      MeetingAuthorizationsAPIServicer
	errorHandler ErrorHandler
}

// MeetingAuthorizationsAPIOption for how the controller is set up.
type MeetingAuthorizationsAPIOption func(*MeetingAuthorizationsAPIController)

// WithMeetingAuthorizationsAPIErrorHandler inject ErrorHandler into controller
func WithMeetingAuthorizationsAPIErrorHandler(h ErrorHandler) MeetingAuthorizationsAPIOption {
	return func(c *MeetingAuthorizationsAPIController) {
		c.errorHandler = h
	}
}

// NewMeetingAuthorizationsAPIController creates a default api controller
func NewMeetingAuthorizationsAPIController(s MeetingAuthorizationsAPIServicer, opts ...MeetingAuthorizationsAPIOption) Router {
	controller := &MeetingAuthorizationsAPIController{
		service:      s,
		errorHandler: DefaultErrorHandler,
	}

	for _, opt := range opts {
		opt(controller)
	}

	return controller
}

// Routes returns all the api routes for the MeetingAuthorizationsAPIController
func (c *MeetingAuthorizationsAPIController) Routes() Routes {
	return Routes{
		"CreateGuestMeetingAuthorization": Route{
			strings.ToUpper("Post"),
			"/1.1/guests/{guest_uuid}/meetings/{meeting_uuid}/authorizations",
			c.CreateGuestMeetingAuthorization,
		},
		"DeleteUserMeetingAuthorization": Route{
			strings.ToUpper("Delete"),
			"/1.1/users/me/meetings/{meeting_uuid}/authorizations/{authorization_uuid}",
			c.DeleteUserMeetingAuthorization,
		},
		"GetGuestMeetingAuthorization": Route{
			strings.ToUpper("Get"),
			"/1.1/guests/{guest_uuid}/meetings/{meeting_uuid}/authorizations/{authorization_uuid}",
			c.GetGuestMeetingAuthorization,
		},
		"GetUserMeetingAuthorization": Route{
			strings.ToUpper("Get"),
			"/1.1/users/me/meetings/{meeting_uuid}/authorizations/{authorization_uuid}",
			c.GetUserMeetingAuthorization,
		},
		"ListUserMeetingAuthorizations": Route{
			strings.ToUpper("Get"),
			"/1.1/user/me/meetings/{meeting_uuid}/authorizations",
			c.ListUserMeetingAuthorizations,
		},
		"PutUserMeetingAuthorizationAccept": Route{
			strings.ToUpper("Put"),
			"/1.1/user/me/meetings/{meeting_uuid}/authorizations/{authorization_uuid}/accept",
			c.PutUserMeetingAuthorizationAccept,
		},
		"PutUserMeetingAuthorizationReject": Route{
			strings.ToUpper("Put"),
			"/1.1/user/me/meetings/{meeting_uuid}/authorizations/{authorization_uuid}/reject",
			c.PutUserMeetingAuthorizationReject,
		},
	}
}

// CreateGuestMeetingAuthorization - Request guest authorization to enter a meeting
func (c *MeetingAuthorizationsAPIController) CreateGuestMeetingAuthorization(w http.ResponseWriter, r *http.Request) {
	guestUuidParam := chi.URLParam(r, "guest_uuid")
	if guestUuidParam == "" {
		c.errorHandler(w, r, &RequiredError{"guest_uuid"}, nil)
		return
	}
	meetingUuidParam := chi.URLParam(r, "meeting_uuid")
	if meetingUuidParam == "" {
		c.errorHandler(w, r, &RequiredError{"meeting_uuid"}, nil)
		return
	}
	bodyParam := MeetingAuthorizationRequest{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&bodyParam); err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	if err := AssertMeetingAuthorizationRequestRequired(bodyParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	if err := AssertMeetingAuthorizationRequestConstraints(bodyParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	result, err := c.service.CreateGuestMeetingAuthorization(r.Context(), guestUuidParam, meetingUuidParam, bodyParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}

// DeleteUserMeetingAuthorization - Delete the guest authorization to enter a meeting
func (c *MeetingAuthorizationsAPIController) DeleteUserMeetingAuthorization(w http.ResponseWriter, r *http.Request) {
	guestUuidParam := chi.URLParam(r, "guest_uuid")
	if guestUuidParam == "" {
		c.errorHandler(w, r, &RequiredError{"guest_uuid"}, nil)
		return
	}
	meetingUuidParam := chi.URLParam(r, "meeting_uuid")
	if meetingUuidParam == "" {
		c.errorHandler(w, r, &RequiredError{"meeting_uuid"}, nil)
		return
	}
	authorizationUuidParam := chi.URLParam(r, "authorization_uuid")
	if authorizationUuidParam == "" {
		c.errorHandler(w, r, &RequiredError{"authorization_uuid"}, nil)
		return
	}
	result, err := c.service.DeleteUserMeetingAuthorization(r.Context(), guestUuidParam, meetingUuidParam, authorizationUuidParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}

// GetGuestMeetingAuthorization - Read the guest authorization to enter a meeting
func (c *MeetingAuthorizationsAPIController) GetGuestMeetingAuthorization(w http.ResponseWriter, r *http.Request) {
	guestUuidParam := chi.URLParam(r, "guest_uuid")
	if guestUuidParam == "" {
		c.errorHandler(w, r, &RequiredError{"guest_uuid"}, nil)
		return
	}
	meetingUuidParam := chi.URLParam(r, "meeting_uuid")
	if meetingUuidParam == "" {
		c.errorHandler(w, r, &RequiredError{"meeting_uuid"}, nil)
		return
	}
	authorizationUuidParam := chi.URLParam(r, "authorization_uuid")
	if authorizationUuidParam == "" {
		c.errorHandler(w, r, &RequiredError{"authorization_uuid"}, nil)
		return
	}
	result, err := c.service.GetGuestMeetingAuthorization(r.Context(), guestUuidParam, meetingUuidParam, authorizationUuidParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}

// GetUserMeetingAuthorization - Read the guest authorization to enter a meeting
func (c *MeetingAuthorizationsAPIController) GetUserMeetingAuthorization(w http.ResponseWriter, r *http.Request) {
	guestUuidParam := chi.URLParam(r, "guest_uuid")
	if guestUuidParam == "" {
		c.errorHandler(w, r, &RequiredError{"guest_uuid"}, nil)
		return
	}
	meetingUuidParam := chi.URLParam(r, "meeting_uuid")
	if meetingUuidParam == "" {
		c.errorHandler(w, r, &RequiredError{"meeting_uuid"}, nil)
		return
	}
	authorizationUuidParam := chi.URLParam(r, "authorization_uuid")
	if authorizationUuidParam == "" {
		c.errorHandler(w, r, &RequiredError{"authorization_uuid"}, nil)
		return
	}
	result, err := c.service.GetUserMeetingAuthorization(r.Context(), guestUuidParam, meetingUuidParam, authorizationUuidParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}

// ListUserMeetingAuthorizations - List all guest authorization requests of a meeting
func (c *MeetingAuthorizationsAPIController) ListUserMeetingAuthorizations(w http.ResponseWriter, r *http.Request) {
	meetingUuidParam := chi.URLParam(r, "meeting_uuid")
	if meetingUuidParam == "" {
		c.errorHandler(w, r, &RequiredError{"meeting_uuid"}, nil)
		return
	}
	result, err := c.service.ListUserMeetingAuthorizations(r.Context(), meetingUuidParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}

// PutUserMeetingAuthorizationAccept - Accept a guest authorization request
func (c *MeetingAuthorizationsAPIController) PutUserMeetingAuthorizationAccept(w http.ResponseWriter, r *http.Request) {
	meetingUuidParam := chi.URLParam(r, "meeting_uuid")
	if meetingUuidParam == "" {
		c.errorHandler(w, r, &RequiredError{"meeting_uuid"}, nil)
		return
	}
	authorizationUuidParam := chi.URLParam(r, "authorization_uuid")
	if authorizationUuidParam == "" {
		c.errorHandler(w, r, &RequiredError{"authorization_uuid"}, nil)
		return
	}
	result, err := c.service.PutUserMeetingAuthorizationAccept(r.Context(), meetingUuidParam, authorizationUuidParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}

// PutUserMeetingAuthorizationReject - Reject a guest authorization request
func (c *MeetingAuthorizationsAPIController) PutUserMeetingAuthorizationReject(w http.ResponseWriter, r *http.Request) {
	meetingUuidParam := chi.URLParam(r, "meeting_uuid")
	if meetingUuidParam == "" {
		c.errorHandler(w, r, &RequiredError{"meeting_uuid"}, nil)
		return
	}
	authorizationUuidParam := chi.URLParam(r, "authorization_uuid")
	if authorizationUuidParam == "" {
		c.errorHandler(w, r, &RequiredError{"authorization_uuid"}, nil)
		return
	}
	result, err := c.service.PutUserMeetingAuthorizationReject(r.Context(), meetingUuidParam, authorizationUuidParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}
