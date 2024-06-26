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

// EmailsAPIController binds http requests to an api service and writes the service results to the http response
type EmailsAPIController struct {
	service      EmailsAPIServicer
	errorHandler ErrorHandler
}

// EmailsAPIOption for how the controller is set up.
type EmailsAPIOption func(*EmailsAPIController)

// WithEmailsAPIErrorHandler inject ErrorHandler into controller
func WithEmailsAPIErrorHandler(h ErrorHandler) EmailsAPIOption {
	return func(c *EmailsAPIController) {
		c.errorHandler = h
	}
}

// NewEmailsAPIController creates a default api controller
func NewEmailsAPIController(s EmailsAPIServicer, opts ...EmailsAPIOption) Router {
	controller := &EmailsAPIController{
		service:      s,
		errorHandler: DefaultErrorHandler,
	}

	for _, opt := range opts {
		opt(controller)
	}

	return controller
}

// Routes returns all the api routes for the EmailsAPIController
func (c *EmailsAPIController) Routes() Routes {
	return Routes{
		"GetEmailConfirm": Route{
			strings.ToUpper("Get"),
			"/0.1/emails/{email_uuid}/confirm",
			c.GetEmailConfirm,
		},
		"GetNewEmailConfirmation": Route{
			strings.ToUpper("Get"),
			"/0.1/users/{user_uuid}/emails/{email_uuid}/confirm",
			c.GetNewEmailConfirmation,
		},
		"PutEmailConfirm": Route{
			strings.ToUpper("Put"),
			"/0.1/emails/{email_uuid}/confirm",
			c.PutEmailConfirm,
		},
		"UpdateAllUserEmails": Route{
			strings.ToUpper("Put"),
			"/0.1/admin/users/{user_uuid}/emails",
			c.UpdateAllUserEmails,
		},
		"UpdateUserEmails": Route{
			strings.ToUpper("Put"),
			"/0.1/users/{user_uuid}/emails",
			c.UpdateUserEmails,
		},
	}
}

// GetEmailConfirm - Confirm an email address
func (c *EmailsAPIController) GetEmailConfirm(w http.ResponseWriter, r *http.Request) {
	query, err := parseQuery(r.URL.RawQuery)
	if err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	emailUuidParam := chi.URLParam(r, "email_uuid")
	if emailUuidParam == "" {
		c.errorHandler(w, r, &RequiredError{"email_uuid"}, nil)
		return
	}
	var tokenParam string
	if query.Has("token") {
		param := query.Get("token")

		tokenParam = param
	} else {
		c.errorHandler(w, r, &RequiredError{Field: "token"}, nil)
		return
	}
	result, err := c.service.GetEmailConfirm(r.Context(), emailUuidParam, tokenParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, result.Headers, w)
}

// GetNewEmailConfirmation - Ask a new confirmation email
func (c *EmailsAPIController) GetNewEmailConfirmation(w http.ResponseWriter, r *http.Request) {
	userUuidParam := chi.URLParam(r, "user_uuid")
	if userUuidParam == "" {
		c.errorHandler(w, r, &RequiredError{"user_uuid"}, nil)
		return
	}
	emailUuidParam := chi.URLParam(r, "email_uuid")
	if emailUuidParam == "" {
		c.errorHandler(w, r, &RequiredError{"email_uuid"}, nil)
		return
	}
	result, err := c.service.GetNewEmailConfirmation(r.Context(), userUuidParam, emailUuidParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, result.Headers, w)
}

// PutEmailConfirm - Confirm an email address
func (c *EmailsAPIController) PutEmailConfirm(w http.ResponseWriter, r *http.Request) {
	emailUuidParam := chi.URLParam(r, "email_uuid")
	if emailUuidParam == "" {
		c.errorHandler(w, r, &RequiredError{"email_uuid"}, nil)
		return
	}
	result, err := c.service.PutEmailConfirm(r.Context(), emailUuidParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, result.Headers, w)
}

// UpdateAllUserEmails - Update email addresses
func (c *EmailsAPIController) UpdateAllUserEmails(w http.ResponseWriter, r *http.Request) {
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
	EncodeJSONResponse(result.Body, &result.Code, result.Headers, w)
}

// UpdateUserEmails - Update email addresses
func (c *EmailsAPIController) UpdateUserEmails(w http.ResponseWriter, r *http.Request) {
	userUuidParam := chi.URLParam(r, "user_uuid")
	if userUuidParam == "" {
		c.errorHandler(w, r, &RequiredError{"user_uuid"}, nil)
		return
	}
	bodyParam := UserEmailList{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&bodyParam); err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	if err := AssertUserEmailListRequired(bodyParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	if err := AssertUserEmailListConstraints(bodyParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	result, err := c.service.UpdateUserEmails(r.Context(), userUuidParam, bodyParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, result.Headers, w)
}
