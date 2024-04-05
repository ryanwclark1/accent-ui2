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

// PagingsAPIController binds http requests to an api service and writes the service results to the http response
type PagingsAPIController struct {
	service      PagingsAPIServicer
	errorHandler ErrorHandler
}

// PagingsAPIOption for how the controller is set up.
type PagingsAPIOption func(*PagingsAPIController)

// WithPagingsAPIErrorHandler inject ErrorHandler into controller
func WithPagingsAPIErrorHandler(h ErrorHandler) PagingsAPIOption {
	return func(c *PagingsAPIController) {
		c.errorHandler = h
	}
}

// NewPagingsAPIController creates a default api controller
func NewPagingsAPIController(s PagingsAPIServicer, opts ...PagingsAPIOption) Router {
	controller := &PagingsAPIController{
		service:      s,
		errorHandler: DefaultErrorHandler,
	}

	for _, opt := range opts {
		opt(controller)
	}

	return controller
}

// Routes returns all the api routes for the PagingsAPIController
func (c *PagingsAPIController) Routes() Routes {
	return Routes{
		"CreatePaging": Route{
			strings.ToUpper("Post"),
			"/1.1/pagings",
			c.CreatePaging,
		},
		"DeletePaging": Route{
			strings.ToUpper("Delete"),
			"/1.1/pagings/{paging_id}",
			c.DeletePaging,
		},
		"GetPaging": Route{
			strings.ToUpper("Get"),
			"/1.1/pagings/{paging_id}",
			c.GetPaging,
		},
		"ListPagings": Route{
			strings.ToUpper("Get"),
			"/1.1/pagings",
			c.ListPagings,
		},
		"UpdatePaging": Route{
			strings.ToUpper("Put"),
			"/1.1/pagings/{paging_id}",
			c.UpdatePaging,
		},
		"UpdatePagingCallerUsers": Route{
			strings.ToUpper("Put"),
			"/1.1/pagings/{paging_id}/callers/users",
			c.UpdatePagingCallerUsers,
		},
		"UpdatePagingMemberUsers": Route{
			strings.ToUpper("Put"),
			"/1.1/pagings/{paging_id}/members/users",
			c.UpdatePagingMemberUsers,
		},
	}
}

// CreatePaging - Create paging
func (c *PagingsAPIController) CreatePaging(w http.ResponseWriter, r *http.Request) {
	bodyParam := Paging{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&bodyParam); err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	if err := AssertPagingRequired(bodyParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	if err := AssertPagingConstraints(bodyParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	accentTenantParam := r.Header.Get("Accent-Tenant")
	result, err := c.service.CreatePaging(r.Context(), bodyParam, accentTenantParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}

// DeletePaging - Delete paging
func (c *PagingsAPIController) DeletePaging(w http.ResponseWriter, r *http.Request) {
	pagingIdParam, err := parseNumericParameter[int32](
		chi.URLParam(r, "paging_id"),
		WithRequire[int32](parseInt32),
	)
	if err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	accentTenantParam := r.Header.Get("Accent-Tenant")
	result, err := c.service.DeletePaging(r.Context(), pagingIdParam, accentTenantParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}

// GetPaging - Get paging
func (c *PagingsAPIController) GetPaging(w http.ResponseWriter, r *http.Request) {
	pagingIdParam, err := parseNumericParameter[int32](
		chi.URLParam(r, "paging_id"),
		WithRequire[int32](parseInt32),
	)
	if err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	accentTenantParam := r.Header.Get("Accent-Tenant")
	result, err := c.service.GetPaging(r.Context(), pagingIdParam, accentTenantParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}

// ListPagings - List paging
func (c *PagingsAPIController) ListPagings(w http.ResponseWriter, r *http.Request) {
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
	var searchParam string
	if query.Has("search") {
		param := query.Get("search")

		searchParam = param
	} else {
	}
	result, err := c.service.ListPagings(r.Context(), accentTenantParam, recurseParam, orderParam, directionParam, limitParam, offsetParam, searchParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}

// UpdatePaging - Update paging
func (c *PagingsAPIController) UpdatePaging(w http.ResponseWriter, r *http.Request) {
	pagingIdParam, err := parseNumericParameter[int32](
		chi.URLParam(r, "paging_id"),
		WithRequire[int32](parseInt32),
	)
	if err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	bodyParam := Paging{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&bodyParam); err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	if err := AssertPagingRequired(bodyParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	if err := AssertPagingConstraints(bodyParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	accentTenantParam := r.Header.Get("Accent-Tenant")
	result, err := c.service.UpdatePaging(r.Context(), pagingIdParam, bodyParam, accentTenantParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}

// UpdatePagingCallerUsers - Update paging and callers
func (c *PagingsAPIController) UpdatePagingCallerUsers(w http.ResponseWriter, r *http.Request) {
	pagingIdParam, err := parseNumericParameter[int32](
		chi.URLParam(r, "paging_id"),
		WithRequire[int32](parseInt32),
	)
	if err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	bodyParam := UsersUuid{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&bodyParam); err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	if err := AssertUsersUuidRequired(bodyParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	if err := AssertUsersUuidConstraints(bodyParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	result, err := c.service.UpdatePagingCallerUsers(r.Context(), pagingIdParam, bodyParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}

// UpdatePagingMemberUsers - Update paging and members
func (c *PagingsAPIController) UpdatePagingMemberUsers(w http.ResponseWriter, r *http.Request) {
	pagingIdParam, err := parseNumericParameter[int32](
		chi.URLParam(r, "paging_id"),
		WithRequire[int32](parseInt32),
	)
	if err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	bodyParam := UsersUuid{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&bodyParam); err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	if err := AssertUsersUuidRequired(bodyParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	if err := AssertUsersUuidConstraints(bodyParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	result, err := c.service.UpdatePagingMemberUsers(r.Context(), pagingIdParam, bodyParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}