/*
 * accent-chatd
 *
 * Control your message and presence from a REST API
 *
 * API version: 1.0.0
 * Contact: help@accentvoice.io
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package chatdserver

import (
	"encoding/json"
	"net/http"
	"strings"
	"time"

	"github.com/go-chi/chi/v5"
)

// MessagesAPIController binds http requests to an api service and writes the service results to the http response
type MessagesAPIController struct {
	service      MessagesAPIServicer
	errorHandler ErrorHandler
}

// MessagesAPIOption for how the controller is set up.
type MessagesAPIOption func(*MessagesAPIController)

// WithMessagesAPIErrorHandler inject ErrorHandler into controller
func WithMessagesAPIErrorHandler(h ErrorHandler) MessagesAPIOption {
	return func(c *MessagesAPIController) {
		c.errorHandler = h
	}
}

// NewMessagesAPIController creates a default api controller
func NewMessagesAPIController(s MessagesAPIServicer, opts ...MessagesAPIOption) Router {
	controller := &MessagesAPIController{
		service:      s,
		errorHandler: DefaultErrorHandler,
	}

	for _, opt := range opts {
		opt(controller)
	}

	return controller
}

// Routes returns all the api routes for the MessagesAPIController
func (c *MessagesAPIController) Routes() Routes {
	return Routes{
		"CreateRoomMessage": Route{
			strings.ToUpper("Post"),
			"/1.0/users/me/rooms/{room_uuid}/messages",
			c.CreateRoomMessage,
		},
		"ListRoomMessage": Route{
			strings.ToUpper("Get"),
			"/1.0/users/me/rooms/{room_uuid}/messages",
			c.ListRoomMessage,
		},
		"ListRoomsMessages": Route{
			strings.ToUpper("Get"),
			"/1.0/users/me/rooms/messages",
			c.ListRoomsMessages,
		},
	}
}

// CreateRoomMessage - Create room messages
func (c *MessagesAPIController) CreateRoomMessage(w http.ResponseWriter, r *http.Request) {
	roomUuidParam := chi.URLParam(r, "room_uuid")
	if roomUuidParam == "" {
		c.errorHandler(w, r, &RequiredError{"room_uuid"}, nil)
		return
	}
	bodyParam := UserMessagePost{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&bodyParam); err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	if err := AssertUserMessagePostRequired(bodyParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	if err := AssertUserMessagePostConstraints(bodyParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	result, err := c.service.CreateRoomMessage(r.Context(), roomUuidParam, bodyParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, result.Headers, w)
}

// ListRoomMessage - List room messages
func (c *MessagesAPIController) ListRoomMessage(w http.ResponseWriter, r *http.Request) {
	query, err := parseQuery(r.URL.RawQuery)
	if err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	roomUuidParam := chi.URLParam(r, "room_uuid")
	if roomUuidParam == "" {
		c.errorHandler(w, r, &RequiredError{"room_uuid"}, nil)
		return
	}
	var fromDateParam time.Time
	if query.Has("from_date") {
		param, err := parseTime(query.Get("from_date"))
		if err != nil {
			c.errorHandler(w, r, &ParsingError{Err: err}, nil)
			return
		}

		fromDateParam = param
	} else {
	}
	var directionParam string
	if query.Has("direction") {
		param := query.Get("direction")

		directionParam = param
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
	var orderParam string
	if query.Has("order") {
		param := query.Get("order")

		orderParam = param
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
	var searchParam string
	if query.Has("search") {
		param := query.Get("search")

		searchParam = param
	} else {
	}
	result, err := c.service.ListRoomMessage(r.Context(), roomUuidParam, fromDateParam, directionParam, limitParam, orderParam, offsetParam, searchParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, result.Headers, w)
}

// ListRoomsMessages - List rooms messages
func (c *MessagesAPIController) ListRoomsMessages(w http.ResponseWriter, r *http.Request) {
	query, err := parseQuery(r.URL.RawQuery)
	if err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	var directionParam string
	if query.Has("direction") {
		param := query.Get("direction")

		directionParam = param
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
	var orderParam string
	if query.Has("order") {
		param := query.Get("order")

		orderParam = param
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
	var searchParam string
	if query.Has("search") {
		param := query.Get("search")

		searchParam = param
	} else {
	}
	var distinctParam string
	if query.Has("distinct") {
		param := query.Get("distinct")

		distinctParam = param
	} else {
	}
	result, err := c.service.ListRoomsMessages(r.Context(), directionParam, limitParam, orderParam, offsetParam, searchParam, distinctParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, result.Headers, w)
}