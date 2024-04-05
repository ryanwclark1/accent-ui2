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

// UsersAPIController binds http requests to an api service and writes the service results to the http response
type UsersAPIController struct {
	service      UsersAPIServicer
	errorHandler ErrorHandler
}

// UsersAPIOption for how the controller is set up.
type UsersAPIOption func(*UsersAPIController)

// WithUsersAPIErrorHandler inject ErrorHandler into controller
func WithUsersAPIErrorHandler(h ErrorHandler) UsersAPIOption {
	return func(c *UsersAPIController) {
		c.errorHandler = h
	}
}

// NewUsersAPIController creates a default api controller
func NewUsersAPIController(s UsersAPIServicer, opts ...UsersAPIOption) Router {
	controller := &UsersAPIController{
		service:      s,
		errorHandler: DefaultErrorHandler,
	}

	for _, opt := range opts {
		opt(controller)
	}

	return controller
}

// Routes returns all the api routes for the UsersAPIController
func (c *UsersAPIController) Routes() Routes {
	return Routes{
		"GetUserSubscription": Route{
			strings.ToUpper("Get"),
			"/1.0/users/me/subscriptions/{subscription_uuid}",
			c.GetUserSubscription,
		},
		"UserCreate": Route{
			strings.ToUpper("Post"),
			"/1.0/users/me/subscriptions",
			c.UserCreate,
		},
		"UserDelete": Route{
			strings.ToUpper("Delete"),
			"/1.0/users/me/subscriptions/{subscription_uuid}",
			c.UserDelete,
		},
		"UserList": Route{
			strings.ToUpper("Get"),
			"/1.0/users/me/subscriptions",
			c.UserList,
		},
	}
}

// GetUserSubscription - Get a user subscription
func (c *UsersAPIController) GetUserSubscription(w http.ResponseWriter, r *http.Request) {
	subscriptionUuidParam := chi.URLParam(r, "subscription_uuid")
	if subscriptionUuidParam == "" {
		c.errorHandler(w, r, &RequiredError{"subscription_uuid"}, nil)
		return
	}
	result, err := c.service.GetUserSubscription(r.Context(), subscriptionUuidParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}

// UserCreate - Subscribe to a HTTP callback (webhook) as a user
func (c *UsersAPIController) UserCreate(w http.ResponseWriter, r *http.Request) {
	bodyParam := UserSubscriptionRequest{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&bodyParam); err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	if err := AssertUserSubscriptionRequestRequired(bodyParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	if err := AssertUserSubscriptionRequestConstraints(bodyParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	result, err := c.service.UserCreate(r.Context(), bodyParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}

// UserDelete - Delete a user subscription
func (c *UsersAPIController) UserDelete(w http.ResponseWriter, r *http.Request) {
	subscriptionUuidParam := chi.URLParam(r, "subscription_uuid")
	if subscriptionUuidParam == "" {
		c.errorHandler(w, r, &RequiredError{"subscription_uuid"}, nil)
		return
	}
	result, err := c.service.UserDelete(r.Context(), subscriptionUuidParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}

// UserList - List subscriptions of a user to HTTP callbacks
func (c *UsersAPIController) UserList(w http.ResponseWriter, r *http.Request) {
	query, err := parseQuery(r.URL.RawQuery)
	if err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	var searchMetadataParam string
	if query.Has("search_metadata") {
		param := query.Get("search_metadata")

		searchMetadataParam = param
	} else {
	}
	result, err := c.service.UserList(r.Context(), searchMetadataParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}
