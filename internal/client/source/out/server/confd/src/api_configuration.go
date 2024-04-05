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

// ConfigurationAPIController binds http requests to an api service and writes the service results to the http response
type ConfigurationAPIController struct {
	service      ConfigurationAPIServicer
	errorHandler ErrorHandler
}

// ConfigurationAPIOption for how the controller is set up.
type ConfigurationAPIOption func(*ConfigurationAPIController)

// WithConfigurationAPIErrorHandler inject ErrorHandler into controller
func WithConfigurationAPIErrorHandler(h ErrorHandler) ConfigurationAPIOption {
	return func(c *ConfigurationAPIController) {
		c.errorHandler = h
	}
}

// NewConfigurationAPIController creates a default api controller
func NewConfigurationAPIController(s ConfigurationAPIServicer, opts ...ConfigurationAPIOption) Router {
	controller := &ConfigurationAPIController{
		service:      s,
		errorHandler: DefaultErrorHandler,
	}

	for _, opt := range opts {
		opt(controller)
	}

	return controller
}

// Routes returns all the api routes for the ConfigurationAPIController
func (c *ConfigurationAPIController) Routes() Routes {
	return Routes{
		"GetConfiguration": Route{
			strings.ToUpper("Get"),
			"/1.1/configuration/live_reload",
			c.GetConfiguration,
		},
		"UpdateConfiguration": Route{
			strings.ToUpper("Put"),
			"/1.1/configuration/live_reload",
			c.UpdateConfiguration,
		},
	}
}

// GetConfiguration - Get live reload status
func (c *ConfigurationAPIController) GetConfiguration(w http.ResponseWriter, r *http.Request) {
	result, err := c.service.GetConfiguration(r.Context())
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, result.Headers, w)
}

// UpdateConfiguration - Update live reload status
func (c *ConfigurationAPIController) UpdateConfiguration(w http.ResponseWriter, r *http.Request) {
	bodyParam := LiveReload{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&bodyParam); err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	if err := AssertLiveReloadRequired(bodyParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	if err := AssertLiveReloadConstraints(bodyParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	result, err := c.service.UpdateConfiguration(r.Context(), bodyParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, result.Headers, w)
}
