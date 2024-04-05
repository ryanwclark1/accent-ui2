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
		"GetCurrentUserCDR": Route{
			strings.ToUpper("Get"),
			"/1.0/users/me/cdr",
			c.GetCurrentUserCDR,
		},
		"GetUserCDR": Route{
			strings.ToUpper("Get"),
			"/1.0/users/{user_uuid}/cdr",
			c.GetUserCDR,
		},
	}
}

// GetCurrentUserCDR - List CDR of the authenticated user
func (c *UsersAPIController) GetCurrentUserCDR(w http.ResponseWriter, r *http.Request) {
	query, err := parseQuery(r.URL.RawQuery)
	if err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
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
	var limitParam int32
	if query.Has("limit") {
		param, err := parseNumericParameter[int32](
			query.Get("limit"),
			WithParse[int32](parseInt32),
		)
		if err != nil {
			c.errorHandler(w, r, &ParsingError{Err: err}, nil)
			return
		}

		limitParam = param
	} else {
	}
	var offsetParam int32
	if query.Has("offset") {
		param, err := parseNumericParameter[int32](
			query.Get("offset"),
			WithParse[int32](parseInt32),
		)
		if err != nil {
			c.errorHandler(w, r, &ParsingError{Err: err}, nil)
			return
		}

		offsetParam = param
	} else {
	}
	var orderParam string
	if query.Has("order") {
		param := query.Get("order")

		orderParam = param
	} else {
	}
	var directionParam string
	if query.Has("direction") {
		param := query.Get("direction")

		directionParam = param
	} else {
	}
	var searchParam string
	if query.Has("search") {
		param := query.Get("search")

		searchParam = param
	} else {
	}
	var callDirectionParam string
	if query.Has("call_direction") {
		param := query.Get("call_direction")

		callDirectionParam = param
	} else {
	}
	var numberParam string
	if query.Has("number") {
		param := query.Get("number")

		numberParam = param
	} else {
	}
	var fromIdParam int32
	if query.Has("from_id") {
		param, err := parseNumericParameter[int32](
			query.Get("from_id"),
			WithParse[int32](parseInt32),
		)
		if err != nil {
			c.errorHandler(w, r, &ParsingError{Err: err}, nil)
			return
		}

		fromIdParam = param
	} else {
	}
	var userUuidParam []string
	if query.Has("user_uuid") {
		userUuidParam = strings.Split(query.Get("user_uuid"), ",")
	}
	var distinctParam string
	if query.Has("distinct") {
		param := query.Get("distinct")

		distinctParam = param
	} else {
	}
	var recordedParam bool
	if query.Has("recorded") {
		param, err := parseBoolParameter(
			query.Get("recorded"),
			WithParse[bool](parseBool),
		)
		if err != nil {
			c.errorHandler(w, r, &ParsingError{Err: err}, nil)
			return
		}

		recordedParam = param
	} else {
	}
	var formatParam string
	if query.Has("format") {
		param := query.Get("format")

		formatParam = param
	} else {
	}
	result, err := c.service.GetCurrentUserCDR(r.Context(), fromParam, untilParam, limitParam, offsetParam, orderParam, directionParam, searchParam, callDirectionParam, numberParam, fromIdParam, userUuidParam, distinctParam, recordedParam, formatParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}

// GetUserCDR - List CDR of the given user
func (c *UsersAPIController) GetUserCDR(w http.ResponseWriter, r *http.Request) {
	query, err := parseQuery(r.URL.RawQuery)
	if err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	userUuidParam := chi.URLParam(r, "user_uuid")
	if userUuidParam == "" {
		c.errorHandler(w, r, &RequiredError{"user_uuid"}, nil)
		return
	}
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
	var limitParam int32
	if query.Has("limit") {
		param, err := parseNumericParameter[int32](
			query.Get("limit"),
			WithParse[int32](parseInt32),
		)
		if err != nil {
			c.errorHandler(w, r, &ParsingError{Err: err}, nil)
			return
		}

		limitParam = param
	} else {
	}
	var offsetParam int32
	if query.Has("offset") {
		param, err := parseNumericParameter[int32](
			query.Get("offset"),
			WithParse[int32](parseInt32),
		)
		if err != nil {
			c.errorHandler(w, r, &ParsingError{Err: err}, nil)
			return
		}

		offsetParam = param
	} else {
	}
	var orderParam string
	if query.Has("order") {
		param := query.Get("order")

		orderParam = param
	} else {
	}
	var directionParam string
	if query.Has("direction") {
		param := query.Get("direction")

		directionParam = param
	} else {
	}
	var searchParam string
	if query.Has("search") {
		param := query.Get("search")

		searchParam = param
	} else {
	}
	var callDirectionParam string
	if query.Has("call_direction") {
		param := query.Get("call_direction")

		callDirectionParam = param
	} else {
	}
	var numberParam string
	if query.Has("number") {
		param := query.Get("number")

		numberParam = param
	} else {
	}
	var fromIdParam int32
	if query.Has("from_id") {
		param, err := parseNumericParameter[int32](
			query.Get("from_id"),
			WithParse[int32](parseInt32),
		)
		if err != nil {
			c.errorHandler(w, r, &ParsingError{Err: err}, nil)
			return
		}

		fromIdParam = param
	} else {
	}
	var distinctParam string
	if query.Has("distinct") {
		param := query.Get("distinct")

		distinctParam = param
	} else {
	}
	var recordedParam bool
	if query.Has("recorded") {
		param, err := parseBoolParameter(
			query.Get("recorded"),
			WithParse[bool](parseBool),
		)
		if err != nil {
			c.errorHandler(w, r, &ParsingError{Err: err}, nil)
			return
		}

		recordedParam = param
	} else {
	}
	var formatParam string
	if query.Has("format") {
		param := query.Get("format")

		formatParam = param
	} else {
	}
	result, err := c.service.GetUserCDR(r.Context(), userUuidParam, fromParam, untilParam, limitParam, offsetParam, orderParam, directionParam, searchParam, callDirectionParam, numberParam, fromIdParam, distinctParam, recordedParam, formatParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}
