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
	"encoding/json"
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5"
)

// PresencesAPIController binds http requests to an api service and writes the service results to the http response
type PresencesAPIController struct {
	service      PresencesAPIServicer
	errorHandler ErrorHandler
}

// PresencesAPIOption for how the controller is set up.
type PresencesAPIOption func(*PresencesAPIController)

// WithPresencesAPIErrorHandler inject ErrorHandler into controller
func WithPresencesAPIErrorHandler(h ErrorHandler) PresencesAPIOption {
	return func(c *PresencesAPIController) {
		c.errorHandler = h
	}
}

// NewPresencesAPIController creates a default api controller
func NewPresencesAPIController(s PresencesAPIServicer, opts ...PresencesAPIOption) Router {
	controller := &PresencesAPIController{
		service:      s,
		errorHandler: DefaultErrorHandler,
	}

	for _, opt := range opts {
		opt(controller)
	}

	return controller
}

// Routes returns all the api routes for the PresencesAPIController
func (c *PresencesAPIController) Routes() Routes {
	return Routes{
		"GetUserPresence": Route{
			strings.ToUpper("Get"),
			"/1.0/users/{user_uuid}/presences",
			c.GetUserPresence,
		},
		"ListPresences": Route{
			strings.ToUpper("Get"),
			"/1.0/users/presences",
			c.ListPresences,
		},
		"UpdateUserPresence": Route{
			strings.ToUpper("Put"),
			"/1.0/users/{user_uuid}/presences",
			c.UpdateUserPresence,
		},
	}
}

// GetUserPresence - Get user presence
func (c *PresencesAPIController) GetUserPresence(w http.ResponseWriter, r *http.Request) {
	userUuidParam := chi.URLParam(r, "user_uuid")
	if userUuidParam == "" {
		c.errorHandler(w, r, &RequiredError{"user_uuid"}, nil)
		return
	}
	accentTenantParam := r.Header.Get("Accent-Tenant")
	result, err := c.service.GetUserPresence(r.Context(), userUuidParam, accentTenantParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, result.Headers, w)
}

// ListPresences - List presences
func (c *PresencesAPIController) ListPresences(w http.ResponseWriter, r *http.Request) {
	query, err := parseQuery(r.URL.RawQuery)
	if err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	accentTenantParam := r.Header.Get("Accent-Tenant")
	var recurseParam bool
	if query.Has("recurse") {
		param, err := parseBoolParameter(
			query.Get("recurse"),
			WithParse[bool](parseBool),
		)
		if err != nil {
			c.errorHandler(w, r, &ParsingError{Err: err}, nil)
			return
		}

		recurseParam = param
	} else {
		var param bool = false
		recurseParam = param
	}
	var userUuidParam []string
	if query.Has("user_uuid") {
		userUuidParam = strings.Split(query.Get("user_uuid"), ",")
	}
	result, err := c.service.ListPresences(r.Context(), accentTenantParam, recurseParam, userUuidParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, result.Headers, w)
}

// UpdateUserPresence - Update user presence
func (c *PresencesAPIController) UpdateUserPresence(w http.ResponseWriter, r *http.Request) {
	userUuidParam := chi.URLParam(r, "user_uuid")
	if userUuidParam == "" {
		c.errorHandler(w, r, &RequiredError{"user_uuid"}, nil)
		return
	}
	bodyParam := Presence{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&bodyParam); err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	if err := AssertPresenceRequired(bodyParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	if err := AssertPresenceConstraints(bodyParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	accentTenantParam := r.Header.Get("Accent-Tenant")
	result, err := c.service.UpdateUserPresence(r.Context(), userUuidParam, bodyParam, accentTenantParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, result.Headers, w)
}
