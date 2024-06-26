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

// QueueStatisticsAPIController binds http requests to an api service and writes the service results to the http response
type QueueStatisticsAPIController struct {
	service      QueueStatisticsAPIServicer
	errorHandler ErrorHandler
}

// QueueStatisticsAPIOption for how the controller is set up.
type QueueStatisticsAPIOption func(*QueueStatisticsAPIController)

// WithQueueStatisticsAPIErrorHandler inject ErrorHandler into controller
func WithQueueStatisticsAPIErrorHandler(h ErrorHandler) QueueStatisticsAPIOption {
	return func(c *QueueStatisticsAPIController) {
		c.errorHandler = h
	}
}

// NewQueueStatisticsAPIController creates a default api controller
func NewQueueStatisticsAPIController(s QueueStatisticsAPIServicer, opts ...QueueStatisticsAPIOption) Router {
	controller := &QueueStatisticsAPIController{
		service:      s,
		errorHandler: DefaultErrorHandler,
	}

	for _, opt := range opts {
		opt(controller)
	}

	return controller
}

// Routes returns all the api routes for the QueueStatisticsAPIController
func (c *QueueStatisticsAPIController) Routes() Routes {
	return Routes{
		"GetQueueQoSStatistics": Route{
			strings.ToUpper("Get"),
			"/1.0/queues/{queue_id}/statistics/qos",
			c.GetQueueQoSStatistics,
		},
		"GetQueueStatistics": Route{
			strings.ToUpper("Get"),
			"/1.0/queues/{queue_id}/statistics",
			c.GetQueueStatistics,
		},
		"GetQueuesStatistics": Route{
			strings.ToUpper("Get"),
			"/1.0/queues/statistics",
			c.GetQueuesStatistics,
		},
	}
}

// GetQueueQoSStatistics - QoS statistics for a specific queue
func (c *QueueStatisticsAPIController) GetQueueQoSStatistics(w http.ResponseWriter, r *http.Request) {
	query, err := parseQuery(r.URL.RawQuery)
	if err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	queueIdParam, err := parseNumericParameter[int32](
		chi.URLParam(r, "queue_id"),
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
	qosThresholdsParam, err := parseNumericArrayParameter[int32](
		query.Get("qos_thresholds"), ",", false,
		WithParse[int32](parseInt32),
	)
	if err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
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
	result, err := c.service.GetQueueQoSStatistics(r.Context(), queueIdParam, accentTenantParam, fromParam, untilParam, intervalParam, qosThresholdsParam, dayStartTimeParam, dayEndTimeParam, weekDaysParam, timezoneParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, result.Headers, w)
}

// GetQueueStatistics - Statistics for a specific queue
func (c *QueueStatisticsAPIController) GetQueueStatistics(w http.ResponseWriter, r *http.Request) {
	query, err := parseQuery(r.URL.RawQuery)
	if err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	queueIdParam, err := parseNumericParameter[int32](
		chi.URLParam(r, "queue_id"),
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
	var qosThresholdParam int32
	if query.Has("qos_threshold") {
		param, err := parseNumericParameter[int32](
			query.Get("qos_threshold"),
			WithParse[int32](parseInt32),
		)
		if err != nil {
			c.errorHandler(w, r, &ParsingError{Err: err}, nil)
			return
		}

		qosThresholdParam = param
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
	result, err := c.service.GetQueueStatistics(r.Context(), queueIdParam, accentTenantParam, fromParam, untilParam, intervalParam, qosThresholdParam, dayStartTimeParam, dayEndTimeParam, weekDaysParam, timezoneParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, result.Headers, w)
}

// GetQueuesStatistics - Statistics for all queues
func (c *QueueStatisticsAPIController) GetQueuesStatistics(w http.ResponseWriter, r *http.Request) {
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
	var qosThresholdParam int32
	if query.Has("qos_threshold") {
		param, err := parseNumericParameter[int32](
			query.Get("qos_threshold"),
			WithParse[int32](parseInt32),
		)
		if err != nil {
			c.errorHandler(w, r, &ParsingError{Err: err}, nil)
			return
		}

		qosThresholdParam = param
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
	result, err := c.service.GetQueuesStatistics(r.Context(), accentTenantParam, fromParam, untilParam, qosThresholdParam, dayStartTimeParam, dayEndTimeParam, weekDaysParam, timezoneParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, result.Headers, w)
}
