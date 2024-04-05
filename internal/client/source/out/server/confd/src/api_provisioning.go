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

// ProvisioningAPIController binds http requests to an api service and writes the service results to the http response
type ProvisioningAPIController struct {
	service      ProvisioningAPIServicer
	errorHandler ErrorHandler
}

// ProvisioningAPIOption for how the controller is set up.
type ProvisioningAPIOption func(*ProvisioningAPIController)

// WithProvisioningAPIErrorHandler inject ErrorHandler into controller
func WithProvisioningAPIErrorHandler(h ErrorHandler) ProvisioningAPIOption {
	return func(c *ProvisioningAPIController) {
		c.errorHandler = h
	}
}

// NewProvisioningAPIController creates a default api controller
func NewProvisioningAPIController(s ProvisioningAPIServicer, opts ...ProvisioningAPIOption) Router {
	controller := &ProvisioningAPIController{
		service:      s,
		errorHandler: DefaultErrorHandler,
	}

	for _, opt := range opts {
		opt(controller)
	}

	return controller
}

// Routes returns all the api routes for the ProvisioningAPIController
func (c *ProvisioningAPIController) Routes() Routes {
	return Routes{
		"GetProvisioningNetworking": Route{
			strings.ToUpper("Get"),
			"/1.1/provisioning/networking",
			c.GetProvisioningNetworking,
		},
		"UpdateProvisioningNetworking": Route{
			strings.ToUpper("Put"),
			"/1.1/provisioning/networking",
			c.UpdateProvisioningNetworking,
		},
	}
}

// GetProvisioningNetworking - Get Provisioning Networking configuration
func (c *ProvisioningAPIController) GetProvisioningNetworking(w http.ResponseWriter, r *http.Request) {
	result, err := c.service.GetProvisioningNetworking(r.Context())
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, result.Headers, w)
}

// UpdateProvisioningNetworking - Update Provisioning Networking configuration
func (c *ProvisioningAPIController) UpdateProvisioningNetworking(w http.ResponseWriter, r *http.Request) {
	bodyParam := ProvisioningNetworking{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&bodyParam); err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	if err := AssertProvisioningNetworkingRequired(bodyParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	if err := AssertProvisioningNetworkingConstraints(bodyParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	result, err := c.service.UpdateProvisioningNetworking(r.Context(), bodyParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, result.Headers, w)
}
