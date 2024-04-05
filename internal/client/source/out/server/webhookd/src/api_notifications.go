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

// NotificationsAPIController binds http requests to an api service and writes the service results to the http response
type NotificationsAPIController struct {
	service      NotificationsAPIServicer
	errorHandler ErrorHandler
}

// NotificationsAPIOption for how the controller is set up.
type NotificationsAPIOption func(*NotificationsAPIController)

// WithNotificationsAPIErrorHandler inject ErrorHandler into controller
func WithNotificationsAPIErrorHandler(h ErrorHandler) NotificationsAPIOption {
	return func(c *NotificationsAPIController) {
		c.errorHandler = h
	}
}

// NewNotificationsAPIController creates a default api controller
func NewNotificationsAPIController(s NotificationsAPIServicer, opts ...NotificationsAPIOption) Router {
	controller := &NotificationsAPIController{
		service:      s,
		errorHandler: DefaultErrorHandler,
	}

	for _, opt := range opts {
		opt(controller)
	}

	return controller
}

// Routes returns all the api routes for the NotificationsAPIController
func (c *NotificationsAPIController) Routes() Routes {
	return Routes{
		"PostMobileNotification": Route{
			strings.ToUpper("Post"),
			"/1.0/mobile/notifications",
			c.PostMobileNotification,
		},
	}
}

// PostMobileNotification - Send a push notification to a user
func (c *NotificationsAPIController) PostMobileNotification(w http.ResponseWriter, r *http.Request) {
	bodyParam := Notification{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&bodyParam); err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	if err := AssertNotificationRequired(bodyParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	if err := AssertNotificationConstraints(bodyParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	accentTenantParam := r.Header.Get("Accent-Tenant")
	result, err := c.service.PostMobileNotification(r.Context(), bodyParam, accentTenantParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, result.Headers, w)
}