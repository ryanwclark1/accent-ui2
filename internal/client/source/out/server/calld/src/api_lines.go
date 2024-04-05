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

// LinesAPIController binds http requests to an api service and writes the service results to the http response
type LinesAPIController struct {
	service      LinesAPIServicer
	errorHandler ErrorHandler
}

// LinesAPIOption for how the controller is set up.
type LinesAPIOption func(*LinesAPIController)

// WithLinesAPIErrorHandler inject ErrorHandler into controller
func WithLinesAPIErrorHandler(h ErrorHandler) LinesAPIOption {
	return func(c *LinesAPIController) {
		c.errorHandler = h
	}
}

// NewLinesAPIController creates a default api controller
func NewLinesAPIController(s LinesAPIServicer, opts ...LinesAPIOption) Router {
	controller := &LinesAPIController{
		service:      s,
		errorHandler: DefaultErrorHandler,
	}

	for _, opt := range opts {
		opt(controller)
	}

	return controller
}

// Routes returns all the api routes for the LinesAPIController
func (c *LinesAPIController) Routes() Routes {
	return Routes{
		"ListLines": Route{
			strings.ToUpper("Get"),
			"/1.0/lines",
			c.ListLines,
		},
	}
}

// ListLines - List line endpoint statuses
func (c *LinesAPIController) ListLines(w http.ResponseWriter, r *http.Request) {
	accentTenantParam := r.Header.Get("Accent-Tenant")
	result, err := c.service.ListLines(r.Context(), accentTenantParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}