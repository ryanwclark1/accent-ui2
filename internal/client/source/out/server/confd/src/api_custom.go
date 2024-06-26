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

// CustomAPIController binds http requests to an api service and writes the service results to the http response
type CustomAPIController struct {
	service      CustomAPIServicer
	errorHandler ErrorHandler
}

// CustomAPIOption for how the controller is set up.
type CustomAPIOption func(*CustomAPIController)

// WithCustomAPIErrorHandler inject ErrorHandler into controller
func WithCustomAPIErrorHandler(h ErrorHandler) CustomAPIOption {
	return func(c *CustomAPIController) {
		c.errorHandler = h
	}
}

// NewCustomAPIController creates a default api controller
func NewCustomAPIController(s CustomAPIServicer, opts ...CustomAPIOption) Router {
	controller := &CustomAPIController{
		service:      s,
		errorHandler: DefaultErrorHandler,
	}

	for _, opt := range opts {
		opt(controller)
	}

	return controller
}

// Routes returns all the api routes for the CustomAPIController
func (c *CustomAPIController) Routes() Routes {
	return Routes{
		"AssociateLineEndpointCustom": Route{
			strings.ToUpper("Put"),
			"/1.1/lines/{line_id}/endpoints/custom/{custom_id}",
			c.AssociateLineEndpointCustom,
		},
		"AssociateTrunkEndpointCustom": Route{
			strings.ToUpper("Put"),
			"/1.1/trunks/{trunk_id}/endpoints/custom/{custom_id}",
			c.AssociateTrunkEndpointCustom,
		},
		"CreateEndpointCustom": Route{
			strings.ToUpper("Post"),
			"/1.1/endpoints/custom",
			c.CreateEndpointCustom,
		},
		"DeleteEndpointCustom": Route{
			strings.ToUpper("Delete"),
			"/1.1/endpoints/custom/{custom_id}",
			c.DeleteEndpointCustom,
		},
		"DissociateLineEndpointCustom": Route{
			strings.ToUpper("Delete"),
			"/1.1/lines/{line_id}/endpoints/custom/{custom_id}",
			c.DissociateLineEndpointCustom,
		},
		"DissociateTrunkEndpointCustom": Route{
			strings.ToUpper("Delete"),
			"/1.1/trunks/{trunk_id}/endpoints/custom/{custom_id}",
			c.DissociateTrunkEndpointCustom,
		},
		"GetEndpointCustom": Route{
			strings.ToUpper("Get"),
			"/1.1/endpoints/custom/{custom_id}",
			c.GetEndpointCustom,
		},
		"ListEndpointsCustom": Route{
			strings.ToUpper("Get"),
			"/1.1/endpoints/custom",
			c.ListEndpointsCustom,
		},
		"UpdateEndpointCustom": Route{
			strings.ToUpper("Put"),
			"/1.1/endpoints/custom/{custom_id}",
			c.UpdateEndpointCustom,
		},
	}
}

// AssociateLineEndpointCustom - Associate line and Custom endpoint
func (c *CustomAPIController) AssociateLineEndpointCustom(w http.ResponseWriter, r *http.Request) {
	lineIdParam, err := parseNumericParameter[int32](
		chi.URLParam(r, "line_id"),
		WithRequire[int32](parseInt32),
	)
	if err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	customIdParam, err := parseNumericParameter[int32](
		chi.URLParam(r, "custom_id"),
		WithRequire[int32](parseInt32),
	)
	if err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	result, err := c.service.AssociateLineEndpointCustom(r.Context(), lineIdParam, customIdParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, result.Headers, w)
}

// AssociateTrunkEndpointCustom - Associate trunk and Custom endpoint
func (c *CustomAPIController) AssociateTrunkEndpointCustom(w http.ResponseWriter, r *http.Request) {
	trunkIdParam, err := parseNumericParameter[int32](
		chi.URLParam(r, "trunk_id"),
		WithRequire[int32](parseInt32),
	)
	if err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	customIdParam, err := parseNumericParameter[int32](
		chi.URLParam(r, "custom_id"),
		WithRequire[int32](parseInt32),
	)
	if err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	result, err := c.service.AssociateTrunkEndpointCustom(r.Context(), trunkIdParam, customIdParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, result.Headers, w)
}

// CreateEndpointCustom - Create Custom endpoint
func (c *CustomAPIController) CreateEndpointCustom(w http.ResponseWriter, r *http.Request) {
	bodyParam := EndpointCustom{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&bodyParam); err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	if err := AssertEndpointCustomRequired(bodyParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	if err := AssertEndpointCustomConstraints(bodyParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	accentTenantParam := r.Header.Get("Accent-Tenant")
	result, err := c.service.CreateEndpointCustom(r.Context(), bodyParam, accentTenantParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, result.Headers, w)
}

// DeleteEndpointCustom - Delete Custom Endpoint
func (c *CustomAPIController) DeleteEndpointCustom(w http.ResponseWriter, r *http.Request) {
	customIdParam, err := parseNumericParameter[int32](
		chi.URLParam(r, "custom_id"),
		WithRequire[int32](parseInt32),
	)
	if err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	accentTenantParam := r.Header.Get("Accent-Tenant")
	result, err := c.service.DeleteEndpointCustom(r.Context(), customIdParam, accentTenantParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, result.Headers, w)
}

// DissociateLineEndpointCustom - Dissociate line and Custom endpoint
func (c *CustomAPIController) DissociateLineEndpointCustom(w http.ResponseWriter, r *http.Request) {
	lineIdParam, err := parseNumericParameter[int32](
		chi.URLParam(r, "line_id"),
		WithRequire[int32](parseInt32),
	)
	if err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	customIdParam, err := parseNumericParameter[int32](
		chi.URLParam(r, "custom_id"),
		WithRequire[int32](parseInt32),
	)
	if err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	result, err := c.service.DissociateLineEndpointCustom(r.Context(), lineIdParam, customIdParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, result.Headers, w)
}

// DissociateTrunkEndpointCustom - Dissociate trunk and Custom endpoint
func (c *CustomAPIController) DissociateTrunkEndpointCustom(w http.ResponseWriter, r *http.Request) {
	trunkIdParam, err := parseNumericParameter[int32](
		chi.URLParam(r, "trunk_id"),
		WithRequire[int32](parseInt32),
	)
	if err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	customIdParam, err := parseNumericParameter[int32](
		chi.URLParam(r, "custom_id"),
		WithRequire[int32](parseInt32),
	)
	if err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	result, err := c.service.DissociateTrunkEndpointCustom(r.Context(), trunkIdParam, customIdParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, result.Headers, w)
}

// GetEndpointCustom - Get Custom Endpoint
func (c *CustomAPIController) GetEndpointCustom(w http.ResponseWriter, r *http.Request) {
	customIdParam, err := parseNumericParameter[int32](
		chi.URLParam(r, "custom_id"),
		WithRequire[int32](parseInt32),
	)
	if err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	accentTenantParam := r.Header.Get("Accent-Tenant")
	result, err := c.service.GetEndpointCustom(r.Context(), customIdParam, accentTenantParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, result.Headers, w)
}

// ListEndpointsCustom - List Custom endpoints
func (c *CustomAPIController) ListEndpointsCustom(w http.ResponseWriter, r *http.Request) {
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
	result, err := c.service.ListEndpointsCustom(r.Context(), accentTenantParam, recurseParam, orderParam, directionParam, limitParam, offsetParam, searchParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, result.Headers, w)
}

// UpdateEndpointCustom - Update Custom Endpoint
func (c *CustomAPIController) UpdateEndpointCustom(w http.ResponseWriter, r *http.Request) {
	customIdParam, err := parseNumericParameter[int32](
		chi.URLParam(r, "custom_id"),
		WithRequire[int32](parseInt32),
	)
	if err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	bodyParam := EndpointCustom{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&bodyParam); err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	if err := AssertEndpointCustomRequired(bodyParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	if err := AssertEndpointCustomConstraints(bodyParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	accentTenantParam := r.Header.Get("Accent-Tenant")
	result, err := c.service.UpdateEndpointCustom(r.Context(), customIdParam, bodyParam, accentTenantParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, result.Headers, w)
}
