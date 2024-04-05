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

// SoundsAPIController binds http requests to an api service and writes the service results to the http response
type SoundsAPIController struct {
	service      SoundsAPIServicer
	errorHandler ErrorHandler
}

// SoundsAPIOption for how the controller is set up.
type SoundsAPIOption func(*SoundsAPIController)

// WithSoundsAPIErrorHandler inject ErrorHandler into controller
func WithSoundsAPIErrorHandler(h ErrorHandler) SoundsAPIOption {
	return func(c *SoundsAPIController) {
		c.errorHandler = h
	}
}

// NewSoundsAPIController creates a default api controller
func NewSoundsAPIController(s SoundsAPIServicer, opts ...SoundsAPIOption) Router {
	controller := &SoundsAPIController{
		service:      s,
		errorHandler: DefaultErrorHandler,
	}

	for _, opt := range opts {
		opt(controller)
	}

	return controller
}

// Routes returns all the api routes for the SoundsAPIController
func (c *SoundsAPIController) Routes() Routes {
	return Routes{
		"CreateSounds": Route{
			strings.ToUpper("Post"),
			"/1.1/sounds",
			c.CreateSounds,
		},
		"DeleteSounds": Route{
			strings.ToUpper("Delete"),
			"/1.1/sounds/{sound_category}",
			c.DeleteSounds,
		},
		"DeleteSoundsFiles": Route{
			strings.ToUpper("Delete"),
			"/1.1/sounds/{sound_category}/files/{sound_filename}",
			c.DeleteSoundsFiles,
		},
		"GetSounds": Route{
			strings.ToUpper("Get"),
			"/1.1/sounds/{sound_category}",
			c.GetSounds,
		},
		"GetSoundsFiles": Route{
			strings.ToUpper("Get"),
			"/1.1/sounds/{sound_category}/files/{sound_filename}",
			c.GetSoundsFiles,
		},
		"ListSounds": Route{
			strings.ToUpper("Get"),
			"/1.1/sounds",
			c.ListSounds,
		},
		"ListSoundsLanguages": Route{
			strings.ToUpper("Get"),
			"/1.1/sounds/languages",
			c.ListSoundsLanguages,
		},
		"UpdateSoundsFiles": Route{
			strings.ToUpper("Put"),
			"/1.1/sounds/{sound_category}/files/{sound_filename}",
			c.UpdateSoundsFiles,
		},
	}
}

// CreateSounds - Create sound category
func (c *SoundsAPIController) CreateSounds(w http.ResponseWriter, r *http.Request) {
	bodyParam := Sound{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&bodyParam); err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	if err := AssertSoundRequired(bodyParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	if err := AssertSoundConstraints(bodyParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	accentTenantParam := r.Header.Get("Accent-Tenant")
	result, err := c.service.CreateSounds(r.Context(), bodyParam, accentTenantParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}

// DeleteSounds - Delete sound category
func (c *SoundsAPIController) DeleteSounds(w http.ResponseWriter, r *http.Request) {
	soundCategoryParam := chi.URLParam(r, "sound_category")
	if soundCategoryParam == "" {
		c.errorHandler(w, r, &RequiredError{"sound_category"}, nil)
		return
	}
	accentTenantParam := r.Header.Get("Accent-Tenant")
	result, err := c.service.DeleteSounds(r.Context(), soundCategoryParam, accentTenantParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}

// DeleteSoundsFiles - Delete audio file
func (c *SoundsAPIController) DeleteSoundsFiles(w http.ResponseWriter, r *http.Request) {
	query, err := parseQuery(r.URL.RawQuery)
	if err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	soundCategoryParam := chi.URLParam(r, "sound_category")
	if soundCategoryParam == "" {
		c.errorHandler(w, r, &RequiredError{"sound_category"}, nil)
		return
	}
	soundFilenameParam := chi.URLParam(r, "sound_filename")
	if soundFilenameParam == "" {
		c.errorHandler(w, r, &RequiredError{"sound_filename"}, nil)
		return
	}
	accentTenantParam := r.Header.Get("Accent-Tenant")
	var languageParam string
	if query.Has("language") {
		param := query.Get("language")

		languageParam = param
	} else {
	}
	var formatParam string
	if query.Has("format") {
		param := query.Get("format")

		formatParam = param
	} else {
	}
	result, err := c.service.DeleteSoundsFiles(r.Context(), soundCategoryParam, soundFilenameParam, accentTenantParam, languageParam, formatParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}

// GetSounds - Get sound category
func (c *SoundsAPIController) GetSounds(w http.ResponseWriter, r *http.Request) {
	soundCategoryParam := chi.URLParam(r, "sound_category")
	if soundCategoryParam == "" {
		c.errorHandler(w, r, &RequiredError{"sound_category"}, nil)
		return
	}
	accentTenantParam := r.Header.Get("Accent-Tenant")
	result, err := c.service.GetSounds(r.Context(), soundCategoryParam, accentTenantParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}

// GetSoundsFiles - Get audio file
func (c *SoundsAPIController) GetSoundsFiles(w http.ResponseWriter, r *http.Request) {
	query, err := parseQuery(r.URL.RawQuery)
	if err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	soundCategoryParam := chi.URLParam(r, "sound_category")
	if soundCategoryParam == "" {
		c.errorHandler(w, r, &RequiredError{"sound_category"}, nil)
		return
	}
	soundFilenameParam := chi.URLParam(r, "sound_filename")
	if soundFilenameParam == "" {
		c.errorHandler(w, r, &RequiredError{"sound_filename"}, nil)
		return
	}
	accentTenantParam := r.Header.Get("Accent-Tenant")
	var languageParam string
	if query.Has("language") {
		param := query.Get("language")

		languageParam = param
	} else {
	}
	var formatParam string
	if query.Has("format") {
		param := query.Get("format")

		formatParam = param
	} else {
	}
	result, err := c.service.GetSoundsFiles(r.Context(), soundCategoryParam, soundFilenameParam, accentTenantParam, languageParam, formatParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}

// ListSounds - List sound categories
func (c *SoundsAPIController) ListSounds(w http.ResponseWriter, r *http.Request) {
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
	result, err := c.service.ListSounds(r.Context(), accentTenantParam, recurseParam, orderParam, directionParam, limitParam, offsetParam, searchParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}

// ListSoundsLanguages - List all languages for sounds
func (c *SoundsAPIController) ListSoundsLanguages(w http.ResponseWriter, r *http.Request) {
	result, err := c.service.ListSoundsLanguages(r.Context())
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}

// UpdateSoundsFiles - Add or update audio file
func (c *SoundsAPIController) UpdateSoundsFiles(w http.ResponseWriter, r *http.Request) {
	query, err := parseQuery(r.URL.RawQuery)
	if err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	soundCategoryParam := chi.URLParam(r, "sound_category")
	if soundCategoryParam == "" {
		c.errorHandler(w, r, &RequiredError{"sound_category"}, nil)
		return
	}
	soundFilenameParam := chi.URLParam(r, "sound_filename")
	if soundFilenameParam == "" {
		c.errorHandler(w, r, &RequiredError{"sound_filename"}, nil)
		return
	}
	bodyParam := map[string]interface{}{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&bodyParam); err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	accentTenantParam := r.Header.Get("Accent-Tenant")
	var languageParam string
	if query.Has("language") {
		param := query.Get("language")

		languageParam = param
	} else {
	}
	var formatParam string
	if query.Has("format") {
		param := query.Get("format")

		formatParam = param
	} else {
	}
	result, err := c.service.UpdateSoundsFiles(r.Context(), soundCategoryParam, soundFilenameParam, bodyParam, accentTenantParam, languageParam, formatParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}
