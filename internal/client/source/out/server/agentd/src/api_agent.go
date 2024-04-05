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

// AgentAPIController binds http requests to an api service and writes the service results to the http response
type AgentAPIController struct {
	service      AgentAPIServicer
	errorHandler ErrorHandler
}

// AgentAPIOption for how the controller is set up.
type AgentAPIOption func(*AgentAPIController)

// WithAgentAPIErrorHandler inject ErrorHandler into controller
func WithAgentAPIErrorHandler(h ErrorHandler) AgentAPIOption {
	return func(c *AgentAPIController) {
		c.errorHandler = h
	}
}

// NewAgentAPIController creates a default api controller
func NewAgentAPIController(s AgentAPIServicer, opts ...AgentAPIOption) Router {
	controller := &AgentAPIController{
		service:      s,
		errorHandler: DefaultErrorHandler,
	}

	for _, opt := range opts {
		opt(controller)
	}

	return controller
}

// Routes returns all the api routes for the AgentAPIController
func (c *AgentAPIController) Routes() Routes {
	return Routes{
		"AddAgentById": Route{
			strings.ToUpper("Post"),
			"/1.0/agents/by-id/{agent_id}/add",
			c.AddAgentById,
		},
		"GetAgentById": Route{
			strings.ToUpper("Get"),
			"/1.0/agents/by-id/{agent_id}",
			c.GetAgentById,
		},
		"GetAgentByNumber": Route{
			strings.ToUpper("Get"),
			"/1.0/agents/by-number/{agent_number}",
			c.GetAgentByNumber,
		},
		"GetUserAgent": Route{
			strings.ToUpper("Get"),
			"/1.0/users/me/agents",
			c.GetUserAgent,
		},
		"LoginAgentById": Route{
			strings.ToUpper("Post"),
			"/1.0/agents/by-id/{agent_id}/login",
			c.LoginAgentById,
		},
		"LoginAgentByNumber": Route{
			strings.ToUpper("Post"),
			"/1.0/agents/by-number/{agent_number}/login",
			c.LoginAgentByNumber,
		},
		"LoginUserAgent": Route{
			strings.ToUpper("Post"),
			"/1.0/users/me/agents/login",
			c.LoginUserAgent,
		},
		"LogoffAgentById": Route{
			strings.ToUpper("Post"),
			"/1.0/agents/by-id/{agent_id}/logoff",
			c.LogoffAgentById,
		},
		"LogoffAgentByNumber": Route{
			strings.ToUpper("Post"),
			"/1.0/agents/by-number/{agent_number}/logoff",
			c.LogoffAgentByNumber,
		},
		"LogoffUserAgent": Route{
			strings.ToUpper("Post"),
			"/1.0/users/me/agents/logoff",
			c.LogoffUserAgent,
		},
		"PauseAgentByNumber": Route{
			strings.ToUpper("Post"),
			"/1.0/agents/by-number/{agent_number}/pause",
			c.PauseAgentByNumber,
		},
		"PauseUserAgent": Route{
			strings.ToUpper("Post"),
			"/1.0/users/me/agents/pause",
			c.PauseUserAgent,
		},
		"RemoveAgentById": Route{
			strings.ToUpper("Post"),
			"/1.0/agents/by-id/{agent_id}/remove",
			c.RemoveAgentById,
		},
		"UnpauseAgentByNumber": Route{
			strings.ToUpper("Post"),
			"/1.0/agents/by-number/{agent_number}/unpause",
			c.UnpauseAgentByNumber,
		},
		"UnpauseUserAgent": Route{
			strings.ToUpper("Post"),
			"/1.0/users/me/agents/unpause",
			c.UnpauseUserAgent,
		},
	}
}

// AddAgentById - Add agent to a queue.
func (c *AgentAPIController) AddAgentById(w http.ResponseWriter, r *http.Request) {
	agentIdParam, err := parseNumericParameter[int32](
		chi.URLParam(r, "agent_id"),
		WithRequire[int32](parseInt32),
	)
	if err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	bodyParam := Queue{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&bodyParam); err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	if err := AssertQueueRequired(bodyParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	if err := AssertQueueConstraints(bodyParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	accentTenantParam := r.Header.Get("Accent-Tenant")
	result, err := c.service.AddAgentById(r.Context(), agentIdParam, bodyParam, accentTenantParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, result.Headers, w)
}

// GetAgentById - Get agent status.
func (c *AgentAPIController) GetAgentById(w http.ResponseWriter, r *http.Request) {
	agentIdParam, err := parseNumericParameter[int32](
		chi.URLParam(r, "agent_id"),
		WithRequire[int32](parseInt32),
	)
	if err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	accentTenantParam := r.Header.Get("Accent-Tenant")
	result, err := c.service.GetAgentById(r.Context(), agentIdParam, accentTenantParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, result.Headers, w)
}

// GetAgentByNumber - Get agent status.
func (c *AgentAPIController) GetAgentByNumber(w http.ResponseWriter, r *http.Request) {
	agentNumberParam := chi.URLParam(r, "agent_number")
	if agentNumberParam == "" {
		c.errorHandler(w, r, &RequiredError{"agent_number"}, nil)
		return
	}
	accentTenantParam := r.Header.Get("Accent-Tenant")
	result, err := c.service.GetAgentByNumber(r.Context(), agentNumberParam, accentTenantParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, result.Headers, w)
}

// GetUserAgent - Get agent status of the user holding the authentication token.
func (c *AgentAPIController) GetUserAgent(w http.ResponseWriter, r *http.Request) {
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

// LoginAgentById - Log an agent.
func (c *AgentAPIController) LoginAgentById(w http.ResponseWriter, r *http.Request) {
	agentIdParam, err := parseNumericParameter[int32](
		chi.URLParam(r, "agent_id"),
		WithRequire[int32](parseInt32),
	)
	if err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	bodyParam := LoginInfo{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&bodyParam); err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	if err := AssertLoginInfoRequired(bodyParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	if err := AssertLoginInfoConstraints(bodyParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	accentTenantParam := r.Header.Get("Accent-Tenant")
	result, err := c.service.LoginAgentById(r.Context(), agentIdParam, bodyParam, accentTenantParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, result.Headers, w)
}

// LoginAgentByNumber - Log an agent.
func (c *AgentAPIController) LoginAgentByNumber(w http.ResponseWriter, r *http.Request) {
	agentNumberParam := chi.URLParam(r, "agent_number")
	if agentNumberParam == "" {
		c.errorHandler(w, r, &RequiredError{"agent_number"}, nil)
		return
	}
	bodyParam := LoginInfo{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&bodyParam); err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	if err := AssertLoginInfoRequired(bodyParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	if err := AssertLoginInfoConstraints(bodyParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	accentTenantParam := r.Header.Get("Accent-Tenant")
	result, err := c.service.LoginAgentByNumber(r.Context(), agentNumberParam, bodyParam, accentTenantParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, result.Headers, w)
}

// LoginUserAgent - Log the agent of the user holding the authentication token
func (c *AgentAPIController) LoginUserAgent(w http.ResponseWriter, r *http.Request) {
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

// LogoffAgentById - Logoff an agent.
func (c *AgentAPIController) LogoffAgentById(w http.ResponseWriter, r *http.Request) {
	agentIdParam, err := parseNumericParameter[int32](
		chi.URLParam(r, "agent_id"),
		WithRequire[int32](parseInt32),
	)
	if err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	accentTenantParam := r.Header.Get("Accent-Tenant")
	result, err := c.service.LogoffAgentById(r.Context(), agentIdParam, accentTenantParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, result.Headers, w)
}

// LogoffAgentByNumber - Logoff an agent.
func (c *AgentAPIController) LogoffAgentByNumber(w http.ResponseWriter, r *http.Request) {
	agentNumberParam := chi.URLParam(r, "agent_number")
	if agentNumberParam == "" {
		c.errorHandler(w, r, &RequiredError{"agent_number"}, nil)
		return
	}
	accentTenantParam := r.Header.Get("Accent-Tenant")
	result, err := c.service.LogoffAgentByNumber(r.Context(), agentNumberParam, accentTenantParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, result.Headers, w)
}

// LogoffUserAgent - Logoff the agent of the user holding the authentication token
func (c *AgentAPIController) LogoffUserAgent(w http.ResponseWriter, r *http.Request) {
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

// PauseAgentByNumber - Pause an agent.
func (c *AgentAPIController) PauseAgentByNumber(w http.ResponseWriter, r *http.Request) {
	agentNumberParam := chi.URLParam(r, "agent_number")
	if agentNumberParam == "" {
		c.errorHandler(w, r, &RequiredError{"agent_number"}, nil)
		return
	}
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
	result, err := c.service.PauseAgentByNumber(r.Context(), agentNumberParam, accentTenantParam, bodyParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, result.Headers, w)
}

// PauseUserAgent - Pause the agent of the user holding the authentication token
func (c *AgentAPIController) PauseUserAgent(w http.ResponseWriter, r *http.Request) {
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

// RemoveAgentById - Remove agent from a queue.
func (c *AgentAPIController) RemoveAgentById(w http.ResponseWriter, r *http.Request) {
	agentIdParam, err := parseNumericParameter[int32](
		chi.URLParam(r, "agent_id"),
		WithRequire[int32](parseInt32),
	)
	if err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	bodyParam := Queue{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&bodyParam); err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	if err := AssertQueueRequired(bodyParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	if err := AssertQueueConstraints(bodyParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	accentTenantParam := r.Header.Get("Accent-Tenant")
	result, err := c.service.RemoveAgentById(r.Context(), agentIdParam, bodyParam, accentTenantParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, result.Headers, w)
}

// UnpauseAgentByNumber - Unpause an agent.
func (c *AgentAPIController) UnpauseAgentByNumber(w http.ResponseWriter, r *http.Request) {
	agentNumberParam := chi.URLParam(r, "agent_number")
	if agentNumberParam == "" {
		c.errorHandler(w, r, &RequiredError{"agent_number"}, nil)
		return
	}
	accentTenantParam := r.Header.Get("Accent-Tenant")
	result, err := c.service.UnpauseAgentByNumber(r.Context(), agentNumberParam, accentTenantParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, result.Headers, w)
}

// UnpauseUserAgent - Unpause the agent of the user holding the authentication token
func (c *AgentAPIController) UnpauseUserAgent(w http.ResponseWriter, r *http.Request) {
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
