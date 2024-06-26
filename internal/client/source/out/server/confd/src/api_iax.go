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

// IaxAPIController binds http requests to an api service and writes the service results to the http response
type IaxAPIController struct {
	service      IaxAPIServicer
	errorHandler ErrorHandler
}

// IaxAPIOption for how the controller is set up.
type IaxAPIOption func(*IaxAPIController)

// WithIaxAPIErrorHandler inject ErrorHandler into controller
func WithIaxAPIErrorHandler(h ErrorHandler) IaxAPIOption {
	return func(c *IaxAPIController) {
		c.errorHandler = h
	}
}

// NewIaxAPIController creates a default api controller
func NewIaxAPIController(s IaxAPIServicer, opts ...IaxAPIOption) Router {
	controller := &IaxAPIController{
		service:      s,
		errorHandler: DefaultErrorHandler,
	}

	for _, opt := range opts {
		opt(controller)
	}

	return controller
}

// Routes returns all the api routes for the IaxAPIController
func (c *IaxAPIController) Routes() Routes {
	return Routes{
		"AssociateTrunkEndpointIax": Route{
			strings.ToUpper("Put"),
			"/1.1/trunks/{trunk_id}/endpoints/iax/{iax_id}",
			c.AssociateTrunkEndpointIax,
		},
		"AssociateTrunkRegisterIax": Route{
			strings.ToUpper("Put"),
			"/1.1/trunks/{trunk_id}/registers/iax/{iax_id}",
			c.AssociateTrunkRegisterIax,
		},
		"CreateEndpointIax": Route{
			strings.ToUpper("Post"),
			"/1.1/endpoints/iax",
			c.CreateEndpointIax,
		},
		"CreateRegisterIax": Route{
			strings.ToUpper("Post"),
			"/1.1/registers/iax",
			c.CreateRegisterIax,
		},
		"DeleteEndpointIax": Route{
			strings.ToUpper("Delete"),
			"/1.1/endpoints/iax/{iax_id}",
			c.DeleteEndpointIax,
		},
		"DeleteRegisterIax": Route{
			strings.ToUpper("Delete"),
			"/1.1/registers/iax/{register_iax_id}",
			c.DeleteRegisterIax,
		},
		"DissociateTrunkEndpointIax": Route{
			strings.ToUpper("Delete"),
			"/1.1/trunks/{trunk_id}/endpoints/iax/{iax_id}",
			c.DissociateTrunkEndpointIax,
		},
		"DissociateTrunkRegisterIax": Route{
			strings.ToUpper("Delete"),
			"/1.1/trunks/{trunk_id}/registers/iax/{iax_id}",
			c.DissociateTrunkRegisterIax,
		},
		"GetEndpointIax": Route{
			strings.ToUpper("Get"),
			"/1.1/endpoints/iax/{iax_id}",
			c.GetEndpointIax,
		},
		"GetRegisterIax": Route{
			strings.ToUpper("Get"),
			"/1.1/registers/iax/{register_iax_id}",
			c.GetRegisterIax,
		},
		"ListAsteriskIaxCallnumberlimits": Route{
			strings.ToUpper("Get"),
			"/1.1/asterisk/iax/callnumberlimits",
			c.ListAsteriskIaxCallnumberlimits,
		},
		"ListEndpointsIax": Route{
			strings.ToUpper("Get"),
			"/1.1/endpoints/iax",
			c.ListEndpointsIax,
		},
		"ListRegistersIax": Route{
			strings.ToUpper("Get"),
			"/1.1/registers/iax",
			c.ListRegistersIax,
		},
		"UpdateAsteriskIaxCallnumberlimits": Route{
			strings.ToUpper("Put"),
			"/1.1/asterisk/iax/callnumberlimits",
			c.UpdateAsteriskIaxCallnumberlimits,
		},
		"UpdateEndpointIax": Route{
			strings.ToUpper("Put"),
			"/1.1/endpoints/iax/{iax_id}",
			c.UpdateEndpointIax,
		},
		"UpdateRegisterIax": Route{
			strings.ToUpper("Put"),
			"/1.1/registers/iax/{register_iax_id}",
			c.UpdateRegisterIax,
		},
	}
}

// AssociateTrunkEndpointIax - Associate trunk and IAX endpoint
func (c *IaxAPIController) AssociateTrunkEndpointIax(w http.ResponseWriter, r *http.Request) {
	trunkIdParam, err := parseNumericParameter[int32](
		chi.URLParam(r, "trunk_id"),
		WithRequire[int32](parseInt32),
	)
	if err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	iaxIdParam, err := parseNumericParameter[int32](
		chi.URLParam(r, "iax_id"),
		WithRequire[int32](parseInt32),
	)
	if err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	result, err := c.service.AssociateTrunkEndpointIax(r.Context(), trunkIdParam, iaxIdParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, result.Headers, w)
}

// AssociateTrunkRegisterIax - Associate trunk and IAX register
func (c *IaxAPIController) AssociateTrunkRegisterIax(w http.ResponseWriter, r *http.Request) {
	trunkIdParam, err := parseNumericParameter[int32](
		chi.URLParam(r, "trunk_id"),
		WithRequire[int32](parseInt32),
	)
	if err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	iaxIdParam, err := parseNumericParameter[int32](
		chi.URLParam(r, "iax_id"),
		WithRequire[int32](parseInt32),
	)
	if err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	result, err := c.service.AssociateTrunkRegisterIax(r.Context(), trunkIdParam, iaxIdParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, result.Headers, w)
}

// CreateEndpointIax - Create IAX endpoint
func (c *IaxAPIController) CreateEndpointIax(w http.ResponseWriter, r *http.Request) {
	accentTenantParam := r.Header.Get("Accent-Tenant")
	bodyParam := EndpointIax{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&bodyParam); err != nil && !errors.Is(err, io.EOF) {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	if err := AssertEndpointIaxRequired(bodyParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	if err := AssertEndpointIaxConstraints(bodyParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	result, err := c.service.CreateEndpointIax(r.Context(), accentTenantParam, bodyParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, result.Headers, w)
}

// CreateRegisterIax - Create register_iax
func (c *IaxAPIController) CreateRegisterIax(w http.ResponseWriter, r *http.Request) {
	bodyParam := RegisterIax{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&bodyParam); err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	if err := AssertRegisterIaxRequired(bodyParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	if err := AssertRegisterIaxConstraints(bodyParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	result, err := c.service.CreateRegisterIax(r.Context(), bodyParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, result.Headers, w)
}

// DeleteEndpointIax - Delete IAX Endpoint
func (c *IaxAPIController) DeleteEndpointIax(w http.ResponseWriter, r *http.Request) {
	iaxIdParam, err := parseNumericParameter[int32](
		chi.URLParam(r, "iax_id"),
		WithRequire[int32](parseInt32),
	)
	if err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	accentTenantParam := r.Header.Get("Accent-Tenant")
	result, err := c.service.DeleteEndpointIax(r.Context(), iaxIdParam, accentTenantParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, result.Headers, w)
}

// DeleteRegisterIax - Delete register IAX
func (c *IaxAPIController) DeleteRegisterIax(w http.ResponseWriter, r *http.Request) {
	registerIaxIdParam, err := parseNumericParameter[int32](
		chi.URLParam(r, "register_iax_id"),
		WithRequire[int32](parseInt32),
	)
	if err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	result, err := c.service.DeleteRegisterIax(r.Context(), registerIaxIdParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, result.Headers, w)
}

// DissociateTrunkEndpointIax - Dissociate trunk and IAX endpoint
func (c *IaxAPIController) DissociateTrunkEndpointIax(w http.ResponseWriter, r *http.Request) {
	trunkIdParam, err := parseNumericParameter[int32](
		chi.URLParam(r, "trunk_id"),
		WithRequire[int32](parseInt32),
	)
	if err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	iaxIdParam, err := parseNumericParameter[int32](
		chi.URLParam(r, "iax_id"),
		WithRequire[int32](parseInt32),
	)
	if err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	result, err := c.service.DissociateTrunkEndpointIax(r.Context(), trunkIdParam, iaxIdParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, result.Headers, w)
}

// DissociateTrunkRegisterIax - Dissociate trunk and IAX register
func (c *IaxAPIController) DissociateTrunkRegisterIax(w http.ResponseWriter, r *http.Request) {
	trunkIdParam, err := parseNumericParameter[int32](
		chi.URLParam(r, "trunk_id"),
		WithRequire[int32](parseInt32),
	)
	if err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	iaxIdParam, err := parseNumericParameter[int32](
		chi.URLParam(r, "iax_id"),
		WithRequire[int32](parseInt32),
	)
	if err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	result, err := c.service.DissociateTrunkRegisterIax(r.Context(), trunkIdParam, iaxIdParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, result.Headers, w)
}

// GetEndpointIax - Get IAX Endpoint
func (c *IaxAPIController) GetEndpointIax(w http.ResponseWriter, r *http.Request) {
	iaxIdParam, err := parseNumericParameter[int32](
		chi.URLParam(r, "iax_id"),
		WithRequire[int32](parseInt32),
	)
	if err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	accentTenantParam := r.Header.Get("Accent-Tenant")
	result, err := c.service.GetEndpointIax(r.Context(), iaxIdParam, accentTenantParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, result.Headers, w)
}

// GetRegisterIax - Get register IAX
func (c *IaxAPIController) GetRegisterIax(w http.ResponseWriter, r *http.Request) {
	registerIaxIdParam, err := parseNumericParameter[int32](
		chi.URLParam(r, "register_iax_id"),
		WithRequire[int32](parseInt32),
	)
	if err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	result, err := c.service.GetRegisterIax(r.Context(), registerIaxIdParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, result.Headers, w)
}

// ListAsteriskIaxCallnumberlimits - List IAX callnumberlimits options
func (c *IaxAPIController) ListAsteriskIaxCallnumberlimits(w http.ResponseWriter, r *http.Request) {
	result, err := c.service.ListAsteriskIaxCallnumberlimits(r.Context())
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, result.Headers, w)
}

// ListEndpointsIax - List IAX endpoints
func (c *IaxAPIController) ListEndpointsIax(w http.ResponseWriter, r *http.Request) {
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
	result, err := c.service.ListEndpointsIax(r.Context(), accentTenantParam, recurseParam, orderParam, directionParam, limitParam, offsetParam, searchParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, result.Headers, w)
}

// ListRegistersIax - List registers iax
func (c *IaxAPIController) ListRegistersIax(w http.ResponseWriter, r *http.Request) {
	query, err := parseQuery(r.URL.RawQuery)
	if err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
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
	result, err := c.service.ListRegistersIax(r.Context(), orderParam, directionParam, limitParam, offsetParam, searchParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, result.Headers, w)
}

// UpdateAsteriskIaxCallnumberlimits - Update IAX callnumberlimits option
func (c *IaxAPIController) UpdateAsteriskIaxCallnumberlimits(w http.ResponseWriter, r *http.Request) {
	bodyParam := IaxCallNumberLimitss{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&bodyParam); err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	if err := AssertIaxCallNumberLimitssRequired(bodyParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	if err := AssertIaxCallNumberLimitssConstraints(bodyParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	result, err := c.service.UpdateAsteriskIaxCallnumberlimits(r.Context(), bodyParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, result.Headers, w)
}

// UpdateEndpointIax - Update IAX Endpoint
func (c *IaxAPIController) UpdateEndpointIax(w http.ResponseWriter, r *http.Request) {
	iaxIdParam, err := parseNumericParameter[int32](
		chi.URLParam(r, "iax_id"),
		WithRequire[int32](parseInt32),
	)
	if err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	bodyParam := EndpointIax{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&bodyParam); err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	if err := AssertEndpointIaxRequired(bodyParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	if err := AssertEndpointIaxConstraints(bodyParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	accentTenantParam := r.Header.Get("Accent-Tenant")
	result, err := c.service.UpdateEndpointIax(r.Context(), iaxIdParam, bodyParam, accentTenantParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, result.Headers, w)
}

// UpdateRegisterIax - Update register IAX
func (c *IaxAPIController) UpdateRegisterIax(w http.ResponseWriter, r *http.Request) {
	registerIaxIdParam, err := parseNumericParameter[int32](
		chi.URLParam(r, "register_iax_id"),
		WithRequire[int32](parseInt32),
	)
	if err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	bodyParam := RegisterIax{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&bodyParam); err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	if err := AssertRegisterIaxRequired(bodyParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	if err := AssertRegisterIaxConstraints(bodyParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	result, err := c.service.UpdateRegisterIax(r.Context(), registerIaxIdParam, bodyParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, result.Headers, w)
}
