/*
 * accent-auth
 *
 * Accent's authentication service
 *
 * API version: 0.1
 * Contact: help@accentvoice.io
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package authserver

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5"
)

// AdminAPIController binds http requests to an api service and writes the service results to the http response
type AdminAPIController struct {
	service      AdminAPIServicer
	errorHandler ErrorHandler
}

// AdminAPIOption for how the controller is set up.
type AdminAPIOption func(*AdminAPIController)

// WithAdminAPIErrorHandler inject ErrorHandler into controller
func WithAdminAPIErrorHandler(h ErrorHandler) AdminAPIOption {
	return func(c *AdminAPIController) {
		c.errorHandler = h
	}
}

// NewAdminAPIController creates a default api controller
func NewAdminAPIController(s AdminAPIServicer, opts ...AdminAPIOption) Router {
	controller := &AdminAPIController{
		service:      s,
		errorHandler: DefaultErrorHandler,
	}

	for _, opt := range opts {
		opt(controller)
	}

	return controller
}

// Routes returns all the api routes for the AdminAPIController
func (c *AdminAPIController) Routes() Routes {
	return Routes{
		"UpdateAllUserEmails": Route{
			strings.ToUpper("Put"),
			"/0.1/admin/users/{user_uuid}/emails",
			c.UpdateAllUserEmails,
		},
	}
}

// UpdateAllUserEmails - Update email addresses
func (c *AdminAPIController) UpdateAllUserEmails(w http.ResponseWriter, r *http.Request) {
	userUuidParam := chi.URLParam(r, "user_uuid")
	if userUuidParam == "" {
		c.errorHandler(w, r, &RequiredError{"user_uuid"}, nil)
		return
	}
	bodyParam := AdminUserEmailList{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&bodyParam); err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	if err := AssertAdminUserEmailListRequired(bodyParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	if err := AssertAdminUserEmailListConstraints(bodyParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	result, err := c.service.UpdateAllUserEmails(r.Context(), userUuidParam, bodyParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}