/*
 * accent-call-logd
 *
 * Consult call logs from a REST API
 *
 * API version: 1.0.0
 * Contact: help@accentvoice.io
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package calllogdserver

import (
	"encoding/json"
	"net/http"
	"strings"
	"time"

	"github.com/go-chi/chi/v5"
)

// AgentStatisticsAPIController binds http requests to an api service and writes the service results to the http response
type AgentStatisticsAPIController struct {
	service      AgentStatisticsAPIServicer
	errorHandler ErrorHandler
}

// AgentStatisticsAPIOption for how the controller is set up.
type AgentStatisticsAPIOption func(*AgentStatisticsAPIController)

// WithAgentStatisticsAPIErrorHandler inject ErrorHandler into controller
func WithAgentStatisticsAPIErrorHandler(h ErrorHandler) AgentStatisticsAPIOption {
	return func(c *AgentStatisticsAPIController) {
		c.errorHandler = h
	}
}

// NewAgentStatisticsAPIController creates a default api controller
func NewAgentStatisticsAPIController(s AgentStatisticsAPIServicer, opts ...AgentStatisticsAPIOption) Router {
	controller := &AgentStatisticsAPIController{
		service:      s,
		errorHandler: DefaultErrorHandler,
	}

	for _, opt := range opts {
		opt(controller)
	}

	return controller
}

// Routes returns all the api routes for the AgentStatisticsAPIController
func (c *AgentStatisticsAPIController) Routes() Routes {
	return Routes{
		"GetAgentStatistics": Route{
			strings.ToUpper("Get"),
			"/1.0/agents/{agent_id}/statistics",
			c.GetAgentStatistics,
		},
		"GetAgentsStatistics": Route{
			strings.ToUpper("Get"),
			"/1.0/agents/statistics",
			c.GetAgentsStatistics,
		},
	}
}

// GetAgentStatistics - Statistics for a specific agent
func (c *AgentStatisticsAPIController) GetAgentStatistics(w http.ResponseWriter, r *http.Request) {
	query, err := parseQuery(r.URL.RawQuery)
	if err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	agentIdParam, err := parseNumericParameter[int32](
		chi.URLParam(r, "agent_id"),
		WithRequire[int32](parseInt32),
	)
	if err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	accentTenantParam := r.Header.Get("Accent-Tenant")
	var fromParam time.Time
	if query.Has("from") {
		param, err := parseTime(query.Get("from"))
		if err != nil {
			c.errorHandler(w, r, &ParsingError{Err: err}, nil)
			return
		}

		fromParam = param
	} else {
	}
	var untilParam time.Time
	if query.Has("until") {
		param, err := parseTime(query.Get("until"))
		if err != nil {
			c.errorHandler(w, r, &ParsingError{Err: err}, nil)
			return
		}

		untilParam = param
	} else {
	}
	var intervalParam string
	if query.Has("interval") {
		param := query.Get("interval")

		intervalParam = param
	} else {
	}
	var dayStartTimeParam string
	if query.Has("day_start_time") {
		param := query.Get("day_start_time")

		dayStartTimeParam = param
	} else {
	}
	var dayEndTimeParam string
	if query.Has("day_end_time") {
		param := query.Get("day_end_time")

		dayEndTimeParam = param
	} else {
	}
	weekDaysParam, err := parseNumericArrayParameter[int32](
		query.Get("week_days"), ",", false,
		WithParse[int32](parseInt32),
	)
	if err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	var timezoneParam string
	if query.Has("timezone") {
		param := query.Get("timezone")

		timezoneParam = param
	} else {
		param := UTC
		timezoneParam = param
	}
	result, err := c.service.GetAgentStatistics(r.Context(), agentIdParam, accentTenantParam, fromParam, untilParam, intervalParam, dayStartTimeParam, dayEndTimeParam, weekDaysParam, timezoneParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, result.Headers, w)
}

// GetAgentsStatistics - Statistics for all agents
func (c *AgentStatisticsAPIController) GetAgentsStatistics(w http.ResponseWriter, r *http.Request) {
	query, err := parseQuery(r.URL.RawQuery)
	if err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	accentTenantParam := r.Header.Get("Accent-Tenant")
	var fromParam time.Time
	if query.Has("from") {
		param, err := parseTime(query.Get("from"))
		if err != nil {
			c.errorHandler(w, r, &ParsingError{Err: err}, nil)
			return
		}

		fromParam = param
	} else {
	}
	var untilParam time.Time
	if query.Has("until") {
		param, err := parseTime(query.Get("until"))
		if err != nil {
			c.errorHandler(w, r, &ParsingError{Err: err}, nil)
			return
		}

		untilParam = param
	} else {
	}
	var dayStartTimeParam string
	if query.Has("day_start_time") {
		param := query.Get("day_start_time")

		dayStartTimeParam = param
	} else {
	}
	var dayEndTimeParam string
	if query.Has("day_end_time") {
		param := query.Get("day_end_time")

		dayEndTimeParam = param
	} else {
	}
	weekDaysParam, err := parseNumericArrayParameter[int32](
		query.Get("week_days"), ",", false,
		WithParse[int32](parseInt32),
	)
	if err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	var timezoneParam string
	if query.Has("timezone") {
		param := query.Get("timezone")

		timezoneParam = param
	} else {
		param := UTC
		timezoneParam = param
	}
	result, err := c.service.GetAgentsStatistics(r.Context(), accentTenantParam, fromParam, untilParam, dayStartTimeParam, dayEndTimeParam, weekDaysParam, timezoneParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, result.Headers, w)
}
