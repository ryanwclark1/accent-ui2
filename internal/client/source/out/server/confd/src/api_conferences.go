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
		"AssociateConferenceExtension": Route{
			strings.ToUpper("Put"),
			"/1.1/conferences/{conference_id}/extensions/{extension_id}",
			c.AssociateConferenceExtension,
		},
		"CreateConference": Route{
			strings.ToUpper("Post"),
			"/1.1/conferences",
			c.CreateConference,
		},
		"DeleteConference": Route{
			strings.ToUpper("Delete"),
			"/1.1/conferences/{conference_id}",
			c.DeleteConference,
		},
		"DissociateConferenceExtension": Route{
			strings.ToUpper("Delete"),
			"/1.1/conferences/{conference_id}/extensions/{extension_id}",
			c.DissociateConferenceExtension,
		},
		"GetConference": Route{
			strings.ToUpper("Get"),
			"/1.1/conferences/{conference_id}",
			c.GetConference,
		},
		"ListAsteriskConfbridgeAccentDefaultBridge": Route{
			strings.ToUpper("Get"),
			"/1.1/asterisk/confbridge/accent_default_bridge",
			c.ListAsteriskConfbridgeAccentDefaultBridge,
		},
		"ListAsteriskConfbridgeAccentDefaultUser": Route{
			strings.ToUpper("Get"),
			"/1.1/asterisk/confbridge/accent_default_user",
			c.ListAsteriskConfbridgeAccentDefaultUser,
		},
		"ListConferences": Route{
			strings.ToUpper("Get"),
			"/1.1/conferences",
			c.ListConferences,
		},
		"UpdateAsteriskConfbridgeAccentDefaultBridge": Route{
			strings.ToUpper("Put"),
			"/1.1/asterisk/confbridge/accent_default_bridge",
			c.UpdateAsteriskConfbridgeAccentDefaultBridge,
		},
		"UpdateAsteriskConfbridgeAccentDefaultUser": Route{
			strings.ToUpper("Put"),
			"/1.1/asterisk/confbridge/accent_default_user",
			c.UpdateAsteriskConfbridgeAccentDefaultUser,
		},
		"UpdateConference": Route{
			strings.ToUpper("Put"),
			"/1.1/conferences/{conference_id}",
			c.UpdateConference,
		},
	}
}

// AssociateConferenceExtension - Associate conference and extension
func (c *ConferencesAPIController) AssociateConferenceExtension(w http.ResponseWriter, r *http.Request) {
	conferenceIdParam, err := parseNumericParameter[int32](
		chi.URLParam(r, "conference_id"),
		WithRequire[int32](parseInt32),
	)
	if err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	extensionIdParam, err := parseNumericParameter[int32](
		chi.URLParam(r, "extension_id"),
		WithRequire[int32](parseInt32),
	)
	if err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	result, err := c.service.AssociateConferenceExtension(r.Context(), conferenceIdParam, extensionIdParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}

// CreateConference - Create conference
func (c *ConferencesAPIController) CreateConference(w http.ResponseWriter, r *http.Request) {
	accentTenantParam := r.Header.Get("Accent-Tenant")
	bodyParam := Conference{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&bodyParam); err != nil && !errors.Is(err, io.EOF) {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	if err := AssertConferenceRequired(bodyParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	if err := AssertConferenceConstraints(bodyParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	result, err := c.service.CreateConference(r.Context(), accentTenantParam, bodyParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}

// DeleteConference - Delete conference
func (c *ConferencesAPIController) DeleteConference(w http.ResponseWriter, r *http.Request) {
	conferenceIdParam, err := parseNumericParameter[int32](
		chi.URLParam(r, "conference_id"),
		WithRequire[int32](parseInt32),
	)
	if err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	accentTenantParam := r.Header.Get("Accent-Tenant")
	result, err := c.service.DeleteConference(r.Context(), conferenceIdParam, accentTenantParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}

// DissociateConferenceExtension - Dissociate conference and extension
func (c *ConferencesAPIController) DissociateConferenceExtension(w http.ResponseWriter, r *http.Request) {
	conferenceIdParam, err := parseNumericParameter[int32](
		chi.URLParam(r, "conference_id"),
		WithRequire[int32](parseInt32),
	)
	if err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	extensionIdParam, err := parseNumericParameter[int32](
		chi.URLParam(r, "extension_id"),
		WithRequire[int32](parseInt32),
	)
	if err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	result, err := c.service.DissociateConferenceExtension(r.Context(), conferenceIdParam, extensionIdParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}

// GetConference - Get conference
func (c *ConferencesAPIController) GetConference(w http.ResponseWriter, r *http.Request) {
	conferenceIdParam, err := parseNumericParameter[int32](
		chi.URLParam(r, "conference_id"),
		WithRequire[int32](parseInt32),
	)
	if err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	accentTenantParam := r.Header.Get("Accent-Tenant")
	result, err := c.service.GetConference(r.Context(), conferenceIdParam, accentTenantParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}

// ListAsteriskConfbridgeAccentDefaultBridge - List ConfBridge accent_default_bridge options
func (c *ConferencesAPIController) ListAsteriskConfbridgeAccentDefaultBridge(w http.ResponseWriter, r *http.Request) {
	result, err := c.service.ListAsteriskConfbridgeAccentDefaultBridge(r.Context())
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}

// ListAsteriskConfbridgeAccentDefaultUser - List ConfBridge accent_default_user options
func (c *ConferencesAPIController) ListAsteriskConfbridgeAccentDefaultUser(w http.ResponseWriter, r *http.Request) {
	result, err := c.service.ListAsteriskConfbridgeAccentDefaultUser(r.Context())
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}

// ListConferences - List conference
func (c *ConferencesAPIController) ListConferences(w http.ResponseWriter, r *http.Request) {
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
	var orderParam string
	if query.Has("order") {
		param := query.Get("order")

		orderParam = param
	} else {
	}
	var directionParam string
	if query.Has("direction") {
		param := query.Get("direction")

		directionParam = param
	} else {
	}
	var limitParam int32
	if query.Has("limit") {
		param, err := parseNumericParameter[int32](
			query.Get("limit"),
			WithParse[int32](parseInt32),
		)
		if err != nil {
			c.errorHandler(w, r, &ParsingError{Err: err}, nil)
			return
		}

		limitParam = param
	} else {
	}
	var offsetParam int32
	if query.Has("offset") {
		param, err := parseNumericParameter[int32](
			query.Get("offset"),
			WithParse[int32](parseInt32),
		)
		if err != nil {
			c.errorHandler(w, r, &ParsingError{Err: err}, nil)
			return
		}

		offsetParam = param
	} else {
	}
	var searchParam string
	if query.Has("search") {
		param := query.Get("search")

		searchParam = param
	} else {
	}
	result, err := c.service.ListConferences(r.Context(), accentTenantParam, recurseParam, orderParam, directionParam, limitParam, offsetParam, searchParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}

// UpdateAsteriskConfbridgeAccentDefaultBridge - Update ConfBridge accent_default_bridge option
func (c *ConferencesAPIController) UpdateAsteriskConfbridgeAccentDefaultBridge(w http.ResponseWriter, r *http.Request) {
	bodyParam := ConfBridgeConfiguration{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&bodyParam); err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	if err := AssertConfBridgeConfigurationRequired(bodyParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	if err := AssertConfBridgeConfigurationConstraints(bodyParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	result, err := c.service.UpdateAsteriskConfbridgeAccentDefaultBridge(r.Context(), bodyParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}

// UpdateAsteriskConfbridgeAccentDefaultUser - Update ConfBridge accent_default_user option
func (c *ConferencesAPIController) UpdateAsteriskConfbridgeAccentDefaultUser(w http.ResponseWriter, r *http.Request) {
	bodyParam := ConfBridgeConfiguration{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&bodyParam); err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	if err := AssertConfBridgeConfigurationRequired(bodyParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	if err := AssertConfBridgeConfigurationConstraints(bodyParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	result, err := c.service.UpdateAsteriskConfbridgeAccentDefaultUser(r.Context(), bodyParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}

// UpdateConference - Update conference
func (c *ConferencesAPIController) UpdateConference(w http.ResponseWriter, r *http.Request) {
	conferenceIdParam, err := parseNumericParameter[int32](
		chi.URLParam(r, "conference_id"),
		WithRequire[int32](parseInt32),
	)
	if err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	bodyParam := Conference{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&bodyParam); err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	if err := AssertConferenceRequired(bodyParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	if err := AssertConferenceConstraints(bodyParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	accentTenantParam := r.Header.Get("Accent-Tenant")
	result, err := c.service.UpdateConference(r.Context(), conferenceIdParam, bodyParam, accentTenantParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}