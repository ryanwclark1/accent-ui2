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

// ForwardsAPIController binds http requests to an api service and writes the service results to the http response
type ForwardsAPIController struct {
	service      ForwardsAPIServicer
	errorHandler ErrorHandler
}

// ForwardsAPIOption for how the controller is set up.
type ForwardsAPIOption func(*ForwardsAPIController)

// WithForwardsAPIErrorHandler inject ErrorHandler into controller
func WithForwardsAPIErrorHandler(h ErrorHandler) ForwardsAPIOption {
	return func(c *ForwardsAPIController) {
		c.errorHandler = h
	}
}

// NewForwardsAPIController creates a default api controller
func NewForwardsAPIController(s ForwardsAPIServicer, opts ...ForwardsAPIOption) Router {
	controller := &ForwardsAPIController{
		service:      s,
		errorHandler: DefaultErrorHandler,
	}

	for _, opt := range opts {
		opt(controller)
	}

	return controller
}

// Routes returns all the api routes for the ForwardsAPIController
func (c *ForwardsAPIController) Routes() Routes {
	return Routes{
		"GetUserForward": Route{
			strings.ToUpper("Get"),
			"/1.1/users/{user_id}/forwards/{forward_name}",
			c.GetUserForward,
		},
		"ListUserForwards": Route{
			strings.ToUpper("Get"),
			"/1.1/users/{user_id}/forwards",
			c.ListUserForwards,
		},
		"UpdateUserForward": Route{
			strings.ToUpper("Put"),
			"/1.1/users/{user_id}/forwards/{forward_name}",
			c.UpdateUserForward,
		},
		"UpdateUserForwards": Route{
			strings.ToUpper("Put"),
			"/1.1/users/{user_id}/forwards",
			c.UpdateUserForwards,
		},
	}
}

// GetUserForward - Get forward for a user
func (c *ForwardsAPIController) GetUserForward(w http.ResponseWriter, r *http.Request) {
	userIdParam := chi.URLParam(r, "user_id")
	if userIdParam == "" {
		c.errorHandler(w, r, &RequiredError{"user_id"}, nil)
		return
	}
	forwardNameParam := chi.URLParam(r, "forward_name")
	if forwardNameParam == "" {
		c.errorHandler(w, r, &RequiredError{"forward_name"}, nil)
		return
	}
	result, err := c.service.GetUserForward(r.Context(), userIdParam, forwardNameParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, result.Headers, w)
}

// ListUserForwards - List forwards for a user
func (c *ForwardsAPIController) ListUserForwards(w http.ResponseWriter, r *http.Request) {
	userIdParam := chi.URLParam(r, "user_id")
	if userIdParam == "" {
		c.errorHandler(w, r, &RequiredError{"user_id"}, nil)
		return
	}
	result, err := c.service.ListUserForwards(r.Context(), userIdParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, result.Headers, w)
}

// UpdateUserForward - Update a forward for a user
func (c *ForwardsAPIController) UpdateUserForward(w http.ResponseWriter, r *http.Request) {
	userIdParam := chi.URLParam(r, "user_id")
	if userIdParam == "" {
		c.errorHandler(w, r, &RequiredError{"user_id"}, nil)
		return
	}
	forwardNameParam := chi.URLParam(r, "forward_name")
	if forwardNameParam == "" {
		c.errorHandler(w, r, &RequiredError{"forward_name"}, nil)
		return
	}
	bodyParam := UserForward{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&bodyParam); err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	if err := AssertUserForwardRequired(bodyParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	if err := AssertUserForwardConstraints(bodyParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	result, err := c.service.UpdateUserForward(r.Context(), userIdParam, forwardNameParam, bodyParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, result.Headers, w)
}

// UpdateUserForwards - Update all forwards for a user
func (c *ForwardsAPIController) UpdateUserForwards(w http.ResponseWriter, r *http.Request) {
	userIdParam := chi.URLParam(r, "user_id")
	if userIdParam == "" {
		c.errorHandler(w, r, &RequiredError{"user_id"}, nil)
		return
	}
	bodyParam := UserForwards{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&bodyParam); err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	if err := AssertUserForwardsRequired(bodyParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	if err := AssertUserForwardsConstraints(bodyParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	result, err := c.service.UpdateUserForwards(r.Context(), userIdParam, bodyParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, result.Headers, w)
}
