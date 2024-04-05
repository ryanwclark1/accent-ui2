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

// CallpickupsAPIController binds http requests to an api service and writes the service results to the http response
type CallpickupsAPIController struct {
	service      CallpickupsAPIServicer
	errorHandler ErrorHandler
}

// CallpickupsAPIOption for how the controller is set up.
type CallpickupsAPIOption func(*CallpickupsAPIController)

// WithCallpickupsAPIErrorHandler inject ErrorHandler into controller
func WithCallpickupsAPIErrorHandler(h ErrorHandler) CallpickupsAPIOption {
	return func(c *CallpickupsAPIController) {
		c.errorHandler = h
	}
}

// NewCallpickupsAPIController creates a default api controller
func NewCallpickupsAPIController(s CallpickupsAPIServicer, opts ...CallpickupsAPIOption) Router {
	controller := &CallpickupsAPIController{
		service:      s,
		errorHandler: DefaultErrorHandler,
	}

	for _, opt := range opts {
		opt(controller)
	}

	return controller
}

// Routes returns all the api routes for the CallpickupsAPIController
func (c *CallpickupsAPIController) Routes() Routes {
	return Routes{
		"CreateCallpickup": Route{
			strings.ToUpper("Post"),
			"/1.1/callpickups",
			c.CreateCallpickup,
		},
		"DeleteCallpickup": Route{
			strings.ToUpper("Delete"),
			"/1.1/callpickups/{callpickup_id}",
			c.DeleteCallpickup,
		},
		"GetCallpickup": Route{
			strings.ToUpper("Get"),
			"/1.1/callpickups/{callpickup_id}",
			c.GetCallpickup,
		},
		"ListCallPickups": Route{
			strings.ToUpper("Get"),
			"/1.1/callpickups",
			c.ListCallPickups,
		},
		"UpdateCallPickupInterceptorGroups": Route{
			strings.ToUpper("Put"),
			"/1.1/callpickups/{callpickup_id}/interceptors/groups",
			c.UpdateCallPickupInterceptorGroups,
		},
		"UpdateCallPickupInterceptorUsers": Route{
			strings.ToUpper("Put"),
			"/1.1/callpickups/{callpickup_id}/interceptors/users",
			c.UpdateCallPickupInterceptorUsers,
		},
		"UpdateCallPickupTargetGroups": Route{
			strings.ToUpper("Put"),
			"/1.1/callpickups/{callpickup_id}/targets/groups",
			c.UpdateCallPickupTargetGroups,
		},
		"UpdateCallPickupTargetUsers": Route{
			strings.ToUpper("Put"),
			"/1.1/callpickups/{callpickup_id}/targets/users",
			c.UpdateCallPickupTargetUsers,
		},
		"UpdateCallpickup": Route{
			strings.ToUpper("Put"),
			"/1.1/callpickups/{callpickup_id}",
			c.UpdateCallpickup,
		},
	}
}

// CreateCallpickup - Create call pickup
func (c *CallpickupsAPIController) CreateCallpickup(w http.ResponseWriter, r *http.Request) {
	bodyParam := CallPickup{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&bodyParam); err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	if err := AssertCallPickupRequired(bodyParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	if err := AssertCallPickupConstraints(bodyParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	accentTenantParam := r.Header.Get("Accent-Tenant")
	result, err := c.service.CreateCallpickup(r.Context(), bodyParam, accentTenantParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, result.Headers, w)
}

// DeleteCallpickup - Delete call pickup
func (c *CallpickupsAPIController) DeleteCallpickup(w http.ResponseWriter, r *http.Request) {
	callpickupIdParam, err := parseNumericParameter[int32](
		chi.URLParam(r, "callpickup_id"),
		WithRequire[int32](parseInt32),
	)
	if err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	accentTenantParam := r.Header.Get("Accent-Tenant")
	result, err := c.service.DeleteCallpickup(r.Context(), callpickupIdParam, accentTenantParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, result.Headers, w)
}

// GetCallpickup - Get call pickup
func (c *CallpickupsAPIController) GetCallpickup(w http.ResponseWriter, r *http.Request) {
	callpickupIdParam, err := parseNumericParameter[int32](
		chi.URLParam(r, "callpickup_id"),
		WithRequire[int32](parseInt32),
	)
	if err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	accentTenantParam := r.Header.Get("Accent-Tenant")
	result, err := c.service.GetCallpickup(r.Context(), callpickupIdParam, accentTenantParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, result.Headers, w)
}

// ListCallPickups - List call pickups
func (c *CallpickupsAPIController) ListCallPickups(w http.ResponseWriter, r *http.Request) {
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
	result, err := c.service.ListCallPickups(r.Context(), accentTenantParam, recurseParam, orderParam, directionParam, limitParam, offsetParam, searchParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, result.Headers, w)
}

// UpdateCallPickupInterceptorGroups - Update call pickup and interceptors
func (c *CallpickupsAPIController) UpdateCallPickupInterceptorGroups(w http.ResponseWriter, r *http.Request) {
	callpickupIdParam, err := parseNumericParameter[int32](
		chi.URLParam(r, "callpickup_id"),
		WithRequire[int32](parseInt32),
	)
	if err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	bodyParam := GroupsId{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&bodyParam); err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	if err := AssertGroupsIdRequired(bodyParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	if err := AssertGroupsIdConstraints(bodyParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	accentTenantParam := r.Header.Get("Accent-Tenant")
	result, err := c.service.UpdateCallPickupInterceptorGroups(r.Context(), callpickupIdParam, bodyParam, accentTenantParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, result.Headers, w)
}

// UpdateCallPickupInterceptorUsers - Update call pickup and interceptors
func (c *CallpickupsAPIController) UpdateCallPickupInterceptorUsers(w http.ResponseWriter, r *http.Request) {
	callpickupIdParam, err := parseNumericParameter[int32](
		chi.URLParam(r, "callpickup_id"),
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
	accentTenantParam := r.Header.Get("Accent-Tenant")
	result, err := c.service.UpdateCallPickupInterceptorUsers(r.Context(), callpickupIdParam, bodyParam, accentTenantParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, result.Headers, w)
}

// UpdateCallPickupTargetGroups - Update call pickup and targets
func (c *CallpickupsAPIController) UpdateCallPickupTargetGroups(w http.ResponseWriter, r *http.Request) {
	callpickupIdParam, err := parseNumericParameter[int32](
		chi.URLParam(r, "callpickup_id"),
		WithRequire[int32](parseInt32),
	)
	if err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	bodyParam := GroupsId{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&bodyParam); err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	if err := AssertGroupsIdRequired(bodyParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	if err := AssertGroupsIdConstraints(bodyParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	accentTenantParam := r.Header.Get("Accent-Tenant")
	result, err := c.service.UpdateCallPickupTargetGroups(r.Context(), callpickupIdParam, bodyParam, accentTenantParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, result.Headers, w)
}

// UpdateCallPickupTargetUsers - Update call pickup and targets
func (c *CallpickupsAPIController) UpdateCallPickupTargetUsers(w http.ResponseWriter, r *http.Request) {
	callpickupIdParam, err := parseNumericParameter[int32](
		chi.URLParam(r, "callpickup_id"),
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
	accentTenantParam := r.Header.Get("Accent-Tenant")
	result, err := c.service.UpdateCallPickupTargetUsers(r.Context(), callpickupIdParam, bodyParam, accentTenantParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, result.Headers, w)
}

// UpdateCallpickup - Update call pickup
func (c *CallpickupsAPIController) UpdateCallpickup(w http.ResponseWriter, r *http.Request) {
	callpickupIdParam, err := parseNumericParameter[int32](
		chi.URLParam(r, "callpickup_id"),
		WithRequire[int32](parseInt32),
	)
	if err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	bodyParam := CallPickup{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&bodyParam); err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	if err := AssertCallPickupRequired(bodyParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	if err := AssertCallPickupConstraints(bodyParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	accentTenantParam := r.Header.Get("Accent-Tenant")
	result, err := c.service.UpdateCallpickup(r.Context(), callpickupIdParam, bodyParam, accentTenantParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, result.Headers, w)
}
