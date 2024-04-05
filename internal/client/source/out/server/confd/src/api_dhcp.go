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

// DhcpAPIController binds http requests to an api service and writes the service results to the http response
type DhcpAPIController struct {
	service      DhcpAPIServicer
	errorHandler ErrorHandler
}

// DhcpAPIOption for how the controller is set up.
type DhcpAPIOption func(*DhcpAPIController)

// WithDhcpAPIErrorHandler inject ErrorHandler into controller
func WithDhcpAPIErrorHandler(h ErrorHandler) DhcpAPIOption {
	return func(c *DhcpAPIController) {
		c.errorHandler = h
	}
}

// NewDhcpAPIController creates a default api controller
func NewDhcpAPIController(s DhcpAPIServicer, opts ...DhcpAPIOption) Router {
	controller := &DhcpAPIController{
		service:      s,
		errorHandler: DefaultErrorHandler,
	}

	for _, opt := range opts {
		opt(controller)
	}

	return controller
}

// Routes returns all the api routes for the DhcpAPIController
func (c *DhcpAPIController) Routes() Routes {
	return Routes{
		"GetDhcp": Route{
			strings.ToUpper("Get"),
			"/1.1/dhcp",
			c.GetDhcp,
		},
		"UpdateDhcp": Route{
			strings.ToUpper("Put"),
			"/1.1/dhcp",
			c.UpdateDhcp,
		},
	}
}

// GetDhcp - Get DHCP configuration
func (c *DhcpAPIController) GetDhcp(w http.ResponseWriter, r *http.Request) {
	result, err := c.service.GetDhcp(r.Context())
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, result.Headers, w)
}

// UpdateDhcp - Update DHCP configuration
func (c *DhcpAPIController) UpdateDhcp(w http.ResponseWriter, r *http.Request) {
	bodyParam := Dhcp{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&bodyParam); err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	if err := AssertDhcpRequired(bodyParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	if err := AssertDhcpConstraints(bodyParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	result, err := c.service.UpdateDhcp(r.Context(), bodyParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, result.Headers, w)
}
