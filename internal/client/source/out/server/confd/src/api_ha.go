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

// HaAPIController binds http requests to an api service and writes the service results to the http response
type HaAPIController struct {
	service      HaAPIServicer
	errorHandler ErrorHandler
}

// HaAPIOption for how the controller is set up.
type HaAPIOption func(*HaAPIController)

// WithHaAPIErrorHandler inject ErrorHandler into controller
func WithHaAPIErrorHandler(h ErrorHandler) HaAPIOption {
	return func(c *HaAPIController) {
		c.errorHandler = h
	}
}

// NewHaAPIController creates a default api controller
func NewHaAPIController(s HaAPIServicer, opts ...HaAPIOption) Router {
	controller := &HaAPIController{
		service:      s,
		errorHandler: DefaultErrorHandler,
	}

	for _, opt := range opts {
		opt(controller)
	}

	return controller
}

// Routes returns all the api routes for the HaAPIController
func (c *HaAPIController) Routes() Routes {
	return Routes{
		"GetHa": Route{
			strings.ToUpper("Get"),
			"/1.1/ha",
			c.GetHa,
		},
		"UpdateHa": Route{
			strings.ToUpper("Put"),
			"/1.1/ha",
			c.UpdateHa,
		},
	}
}

// GetHa - Get High Availability configuration
func (c *HaAPIController) GetHa(w http.ResponseWriter, r *http.Request) {
	result, err := c.service.GetHa(r.Context())
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, result.Headers, w)
}

// UpdateHa - Update High Availability configuration
func (c *HaAPIController) UpdateHa(w http.ResponseWriter, r *http.Request) {
	bodyParam := Ha{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&bodyParam); err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	if err := AssertHaRequired(bodyParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	if err := AssertHaConstraints(bodyParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	result, err := c.service.UpdateHa(r.Context(), bodyParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, result.Headers, w)
}
