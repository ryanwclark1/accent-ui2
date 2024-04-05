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
	"encoding/json"
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5"
)

// ConferencesAPIController binds http requests to an api service and writes the service results to the http response
type ConferencesAPIController struct {
	service      ConferencesAPIServicer
	errorHandler ErrorHandler
}

// ConferencesAPIOption for how the controller is set up.
type ConferencesAPIOption func(*ConferencesAPIController)

// WithConferencesAPIErrorHandler inject ErrorHandler into controller
func WithConferencesAPIErrorHandler(h ErrorHandler) ConferencesAPIOption {
	return func(c *ConferencesAPIController) {
		c.errorHandler = h
	}
}

// NewConferencesAPIController creates a default api controller
func NewConferencesAPIController(s ConferencesAPIServicer, opts ...ConferencesAPIOption) Router {
	controller := &ConferencesAPIController{
		service:      s,
		errorHandler: DefaultErrorHandler,
	}

	for _, opt := range opts {
		opt(controller)
	}

	return controller
}

// Routes returns all the api routes for the ConferencesAPIController
func (c *ConferencesAPIController) Routes() Routes {
	return Routes{
		"KickParticipant": Route{
			strings.ToUpper("Delete"),
			"/1.0/conferences/{conference_id}/participants/{participant_id}",
			c.KickParticipant,
		},
		"ListConferenceParticipants": Route{
			strings.ToUpper("Get"),
			"/1.0/conferences/{conference_id}/participants",
			c.ListConferenceParticipants,
		},
		"ListUserConferenceParticipants": Route{
			strings.ToUpper("Get"),
			"/1.0/users/me/conferences/{conference_id}/participants",
			c.ListUserConferenceParticipants,
		},
		"MuteParticipant": Route{
			strings.ToUpper("Put"),
			"/1.0/conferences/{conference_id}/participants/{participant_id}/mute",
			c.MuteParticipant,
		},
		"StartConferenceRecording": Route{
			strings.ToUpper("Post"),
			"/1.0/conferences/{conference_id}/record",
			c.StartConferenceRecording,
		},
		"StopConferenceRecording": Route{
			strings.ToUpper("Delete"),
			"/1.0/conferences/{conference_id}/record",
			c.StopConferenceRecording,
		},
		"UnmuteParticipant": Route{
			strings.ToUpper("Put"),
			"/1.0/conferences/{conference_id}/participants/{participant_id}/unmute",
			c.UnmuteParticipant,
		},
	}
}

// KickParticipant - Kick participant from a conference
func (c *ConferencesAPIController) KickParticipant(w http.ResponseWriter, r *http.Request) {
	conferenceIdParam := chi.URLParam(r, "conference_id")
	if conferenceIdParam == "" {
		c.errorHandler(w, r, &RequiredError{"conference_id"}, nil)
		return
	}
	participantIdParam := chi.URLParam(r, "participant_id")
	if participantIdParam == "" {
		c.errorHandler(w, r, &RequiredError{"participant_id"}, nil)
		return
	}
	result, err := c.service.KickParticipant(r.Context(), conferenceIdParam, participantIdParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, result.Headers, w)
}

// ListConferenceParticipants - List participants of a conference
func (c *ConferencesAPIController) ListConferenceParticipants(w http.ResponseWriter, r *http.Request) {
	conferenceIdParam := chi.URLParam(r, "conference_id")
	if conferenceIdParam == "" {
		c.errorHandler(w, r, &RequiredError{"conference_id"}, nil)
		return
	}
	result, err := c.service.ListConferenceParticipants(r.Context(), conferenceIdParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, result.Headers, w)
}

// ListUserConferenceParticipants - List participants of a conference as a user
func (c *ConferencesAPIController) ListUserConferenceParticipants(w http.ResponseWriter, r *http.Request) {
	conferenceIdParam := chi.URLParam(r, "conference_id")
	if conferenceIdParam == "" {
		c.errorHandler(w, r, &RequiredError{"conference_id"}, nil)
		return
	}
	result, err := c.service.ListUserConferenceParticipants(r.Context(), conferenceIdParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, result.Headers, w)
}

// MuteParticipant - Mute a participant in a conference
func (c *ConferencesAPIController) MuteParticipant(w http.ResponseWriter, r *http.Request) {
	conferenceIdParam := chi.URLParam(r, "conference_id")
	if conferenceIdParam == "" {
		c.errorHandler(w, r, &RequiredError{"conference_id"}, nil)
		return
	}
	participantIdParam := chi.URLParam(r, "participant_id")
	if participantIdParam == "" {
		c.errorHandler(w, r, &RequiredError{"participant_id"}, nil)
		return
	}
	result, err := c.service.MuteParticipant(r.Context(), conferenceIdParam, participantIdParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, result.Headers, w)
}

// StartConferenceRecording - Record a conference
func (c *ConferencesAPIController) StartConferenceRecording(w http.ResponseWriter, r *http.Request) {
	conferenceIdParam := chi.URLParam(r, "conference_id")
	if conferenceIdParam == "" {
		c.errorHandler(w, r, &RequiredError{"conference_id"}, nil)
		return
	}
	result, err := c.service.StartConferenceRecording(r.Context(), conferenceIdParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, result.Headers, w)
}

// StopConferenceRecording - Stop recording a conference
func (c *ConferencesAPIController) StopConferenceRecording(w http.ResponseWriter, r *http.Request) {
	conferenceIdParam := chi.URLParam(r, "conference_id")
	if conferenceIdParam == "" {
		c.errorHandler(w, r, &RequiredError{"conference_id"}, nil)
		return
	}
	result, err := c.service.StopConferenceRecording(r.Context(), conferenceIdParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, result.Headers, w)
}

// UnmuteParticipant - Unmute a participant in a conference
func (c *ConferencesAPIController) UnmuteParticipant(w http.ResponseWriter, r *http.Request) {
	conferenceIdParam := chi.URLParam(r, "conference_id")
	if conferenceIdParam == "" {
		c.errorHandler(w, r, &RequiredError{"conference_id"}, nil)
		return
	}
	participantIdParam := chi.URLParam(r, "participant_id")
	if participantIdParam == "" {
		c.errorHandler(w, r, &RequiredError{"participant_id"}, nil)
		return
	}
	result, err := c.service.UnmuteParticipant(r.Context(), conferenceIdParam, participantIdParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, result.Headers, w)
}