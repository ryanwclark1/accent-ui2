/*
 * accent-webhookd
 *
 * Control your webhooks from a REST API
 *
 * API version: 1.0.0
 * Contact: help@accentvoice.io
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package webhookdserver

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5"
)

// StatusAPIController binds http requests to an api service and writes the service results to the http response
type StatusAPIController struct {
	service      StatusAPIServicer
	errorHandler ErrorHandler
}

// StatusAPIOption for how the controller is set up.
type StatusAPIOption func(*StatusAPIController)

// WithStatusAPIErrorHandler inject ErrorHandler into controller
func WithStatusAPIErrorHandler(h ErrorHandler) StatusAPIOption {
	return func(c *StatusAPIController) {
		c.errorHandler = h
	}
}

// NewStatusAPIController creates a default api controller
func NewStatusAPIController(s StatusAPIServicer, opts ...StatusAPIOption) Router {
	controller := &StatusAPIController{
		service:      s,
		errorHandler: DefaultErrorHandler,
	}

	for _, opt := range opts {
		opt(controller)
	}

	return controller
}

// Routes returns all the api routes for the StatusAPIController
func (c *StatusAPIController) Routes() Routes {
	return Routes{
		"GetStatus": Route{
			strings.ToUpper("Get"),
			"/1.0/status",
			c.GetStatus,
		},
	}
}

// GetStatus - Print infos about internal status of accent-webhookd
func (c *StatusAPIController) GetStatus(w http.ResponseWriter, r *http.Request) {
	result, err := c.service.GetStatus(r.Context())
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, result.Headers, w)
}
