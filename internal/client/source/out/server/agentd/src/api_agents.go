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

// AgentsAPIController binds http requests to an api service and writes the service results to the http response
type AgentsAPIController struct {
	service      AgentsAPIServicer
	errorHandler ErrorHandler
}

// AgentsAPIOption for how the controller is set up.
type AgentsAPIOption func(*AgentsAPIController)

// WithAgentsAPIErrorHandler inject ErrorHandler into controller
func WithAgentsAPIErrorHandler(h ErrorHandler) AgentsAPIOption {
	return func(c *AgentsAPIController) {
		c.errorHandler = h
	}
}

// NewAgentsAPIController creates a default api controller
func NewAgentsAPIController(s AgentsAPIServicer, opts ...AgentsAPIOption) Router {
	controller := &AgentsAPIController{
		service:      s,
		errorHandler: DefaultErrorHandler,
	}

	for _, opt := range opts {
		opt(controller)
	}

	return controller
}

// Routes returns all the api routes for the AgentsAPIController
func (c *AgentsAPIController) Routes() Routes {
	return Routes{
		"GetAgents": Route{
			strings.ToUpper("Get"),
			"/1.0/agents",
			c.GetAgents,
		},
		"LogoffAgents": Route{
			strings.ToUpper("Post"),
			"/1.0/agents/logoff",
			c.LogoffAgents,
		},
		"RelogAgents": Route{
			strings.ToUpper("Post"),
			"/1.0/agents/relog",
			c.RelogAgents,
		},
	}
}

// GetAgents - Get the status of all agents.
func (c *AgentsAPIController) GetAgents(w http.ResponseWriter, r *http.Request) {
	query, err := parseQuery(r.URL.RawQuery)
	if err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	accentTenantParam := r.Header.Get("Accent-Tenant")
	var recurseParam bool
	if query.Has("recurse") {
		param, err := parseBoolParameter(
			query.Get("recurse"),
			WithParse[bool](parseBool),
		)
		if err != nil {
			c.errorHandler(w, r, &ParsingError{Err: err}, nil)
			return
		}

		recurseParam = param
	} else {
		var param bool = false
		recurseParam = param
	}
	result, err := c.service.GetAgents(r.Context(), accentTenantParam, recurseParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, result.Headers, w)
}

// LogoffAgents - Logoff all agents.
func (c *AgentsAPIController) LogoffAgents(w http.ResponseWriter, r *http.Request) {
	accentTenantParam := r.Header.Get("Accent-Tenant")
	result, err := c.service.LogoffAgents(r.Context(), accentTenantParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, result.Headers, w)
}

// RelogAgents - Relog all agents.
func (c *AgentsAPIController) RelogAgents(w http.ResponseWriter, r *http.Request) {
	accentTenantParam := r.Header.Get("Accent-Tenant")
	result, err := c.service.RelogAgents(r.Context(), accentTenantParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, result.Headers, w)
}
