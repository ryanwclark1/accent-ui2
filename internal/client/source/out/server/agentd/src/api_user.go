/*
 * accent-agentd
 *
 * Agent service
 *
 * API version: 1.0
 * Contact: help@accentvoice.io
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package agentdserver

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5"
)

// UserAPIController binds http requests to an api service and writes the service results to the http response
type UserAPIController struct {
	service      UserAPIServicer
	errorHandler ErrorHandler
}

// UserAPIOption for how the controller is set up.
type UserAPIOption func(*UserAPIController)

// WithUserAPIErrorHandler inject ErrorHandler into controller
func WithUserAPIErrorHandler(h ErrorHandler) UserAPIOption {
	return func(c *UserAPIController) {
		c.errorHandler = h
	}
}

// NewUserAPIController creates a default api controller
func NewUserAPIController(s UserAPIServicer, opts ...UserAPIOption) Router {
	controller := &UserAPIController{
		service:      s,
		errorHandler: DefaultErrorHandler,
	}

	for _, opt := range opts {
		opt(controller)
	}

	return controller
}

// Routes returns all the api routes for the UserAPIController
func (c *UserAPIController) Routes() Routes {
	return Routes{
		"GetUserAgent": Route{
			strings.ToUpper("Get"),
			"/1.0/users/me/agents",
			c.GetUserAgent,
		},
		"LoginUserAgent": Route{
			strings.ToUpper("Post"),
			"/1.0/users/me/agents/login",
			c.LoginUserAgent,
		},
		"LogoffUserAgent": Route{
			strings.ToUpper("Post"),
			"/1.0/users/me/agents/logoff",
			c.LogoffUserAgent,
		},
		"PauseUserAgent": Route{
			strings.ToUpper("Post"),
			"/1.0/users/me/agents/pause",
			c.PauseUserAgent,
		},
		"UnpauseUserAgent": Route{
			strings.ToUpper("Post"),
			"/1.0/users/me/agents/unpause",
			c.UnpauseUserAgent,
		},
	}
}

// GetUserAgent - Get agent status of the user holding the authentication token.
func (c *UserAPIController) GetUserAgent(w http.ResponseWriter, r *http.Request) {
	accentTenantParam := r.Header.Get("Accent-Tenant")
	result, err := c.service.GetUserAgent(r.Context(), accentTenantParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, result.Headers, w)
}

// LoginUserAgent - Log the agent of the user holding the authentication token
func (c *UserAPIController) LoginUserAgent(w http.ResponseWriter, r *http.Request) {
	bodyParam := UserAgentLoginInfo{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&bodyParam); err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	if err := AssertUserAgentLoginInfoRequired(bodyParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	if err := AssertUserAgentLoginInfoConstraints(bodyParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	accentTenantParam := r.Header.Get("Accent-Tenant")
	result, err := c.service.LoginUserAgent(r.Context(), bodyParam, accentTenantParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, result.Headers, w)
}

// LogoffUserAgent - Logoff the agent of the user holding the authentication token
func (c *UserAPIController) LogoffUserAgent(w http.ResponseWriter, r *http.Request) {
	accentTenantParam := r.Header.Get("Accent-Tenant")
	result, err := c.service.LogoffUserAgent(r.Context(), accentTenantParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, result.Headers, w)
}

// PauseUserAgent - Pause the agent of the user holding the authentication token
func (c *UserAPIController) PauseUserAgent(w http.ResponseWriter, r *http.Request) {
	accentTenantParam := r.Header.Get("Accent-Tenant")
	bodyParam := AgentPauseReason{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&bodyParam); err != nil && !errors.Is(err, io.EOF) {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	if err := AssertAgentPauseReasonRequired(bodyParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	if err := AssertAgentPauseReasonConstraints(bodyParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	result, err := c.service.PauseUserAgent(r.Context(), accentTenantParam, bodyParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, result.Headers, w)
}

// UnpauseUserAgent - Unpause the agent of the user holding the authentication token
func (c *UserAPIController) UnpauseUserAgent(w http.ResponseWriter, r *http.Request) {
	accentTenantParam := r.Header.Get("Accent-Tenant")
	result, err := c.service.UnpauseUserAgent(r.Context(), accentTenantParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, result.Headers, w)
}