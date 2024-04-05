/*
 * accent-setupd
 *
 * Initialize Accent Engine from a REST API
 *
 * API version: 1.0.0
 * Contact: help@accentvoice.io
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package setupdserver

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5"
)

// SetupAPIController binds http requests to an api service and writes the service results to the http response
type SetupAPIController struct {
	service      SetupAPIServicer
	errorHandler ErrorHandler
}

// SetupAPIOption for how the controller is set up.
type SetupAPIOption func(*SetupAPIController)

// WithSetupAPIErrorHandler inject ErrorHandler into controller
func WithSetupAPIErrorHandler(h ErrorHandler) SetupAPIOption {
	return func(c *SetupAPIController) {
		c.errorHandler = h
	}
}

// NewSetupAPIController creates a default api controller
func NewSetupAPIController(s SetupAPIServicer, opts ...SetupAPIOption) Router {
	controller := &SetupAPIController{
		service:      s,
		errorHandler: DefaultErrorHandler,
	}

	for _, opt := range opts {
		opt(controller)
	}

	return controller
}

// Routes returns all the api routes for the SetupAPIController
func (c *SetupAPIController) Routes() Routes {
	return Routes{
		"Create": Route{
			strings.ToUpper("Post"),
			"/1.0/setup",
			c.Create,
		},
	}
}

// Create - Setup the Accent Engine
func (c *SetupAPIController) Create(w http.ResponseWriter, r *http.Request) {
	bodyParam := SetupRequest{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&bodyParam); err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	if err := AssertSetupRequestRequired(bodyParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	if err := AssertSetupRequestConstraints(bodyParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	result, err := c.service.Create(r.Context(), bodyParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}
