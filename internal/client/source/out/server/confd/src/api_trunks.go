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

// TrunksAPIController binds http requests to an api service and writes the service results to the http response
type TrunksAPIController struct {
	service      TrunksAPIServicer
	errorHandler ErrorHandler
}

// TrunksAPIOption for how the controller is set up.
type TrunksAPIOption func(*TrunksAPIController)

// WithTrunksAPIErrorHandler inject ErrorHandler into controller
func WithTrunksAPIErrorHandler(h ErrorHandler) TrunksAPIOption {
	return func(c *TrunksAPIController) {
		c.errorHandler = h
	}
}

// NewTrunksAPIController creates a default api controller
func NewTrunksAPIController(s TrunksAPIServicer, opts ...TrunksAPIOption) Router {
	controller := &TrunksAPIController{
		service:      s,
		errorHandler: DefaultErrorHandler,
	}

	for _, opt := range opts {
		opt(controller)
	}

	return controller
}

// Routes returns all the api routes for the TrunksAPIController
func (c *TrunksAPIController) Routes() Routes {
	return Routes{
		"AssociateOutcallTrunks": Route{
			strings.ToUpper("Put"),
			"/1.1/outcalls/{outcall_id}/trunks",
			c.AssociateOutcallTrunks,
		},
		"AssociateTrunkEndpointCustom": Route{
			strings.ToUpper("Put"),
			"/1.1/trunks/{trunk_id}/endpoints/custom/{custom_id}",
			c.AssociateTrunkEndpointCustom,
		},
		"AssociateTrunkEndpointIax": Route{
			strings.ToUpper("Put"),
			"/1.1/trunks/{trunk_id}/endpoints/iax/{iax_id}",
			c.AssociateTrunkEndpointIax,
		},
		"AssociateTrunkEndpointSip": Route{
			strings.ToUpper("Put"),
			"/1.1/trunks/{trunk_id}/endpoints/sip/{sip_uuid}",
			c.AssociateTrunkEndpointSip,
		},
		"AssociateTrunkRegisterIax": Route{
			strings.ToUpper("Put"),
			"/1.1/trunks/{trunk_id}/registers/iax/{iax_id}",
			c.AssociateTrunkRegisterIax,
		},
		"CreateTrunk": Route{
			strings.ToUpper("Post"),
			"/1.1/trunks",
			c.CreateTrunk,
		},
		"DeleteTrunk": Route{
			strings.ToUpper("Delete"),
			"/1.1/trunks/{trunk_id}",
			c.DeleteTrunk,
		},
		"DissociateTrunkEndpointCustom": Route{
			strings.ToUpper("Delete"),
			"/1.1/trunks/{trunk_id}/endpoints/custom/{custom_id}",
			c.DissociateTrunkEndpointCustom,
		},
		"DissociateTrunkEndpointIax": Route{
			strings.ToUpper("Delete"),
			"/1.1/trunks/{trunk_id}/endpoints/iax/{iax_id}",
			c.DissociateTrunkEndpointIax,
		},
		"DissociateTrunkEndpointSip": Route{
			strings.ToUpper("Delete"),
			"/1.1/trunks/{trunk_id}/endpoints/sip/{sip_uuid}",
			c.DissociateTrunkEndpointSip,
		},
		"DissociateTrunkRegisterIax": Route{
			strings.ToUpper("Delete"),
			"/1.1/trunks/{trunk_id}/registers/iax/{iax_id}",
			c.DissociateTrunkRegisterIax,
		},
		"GetTrunk": Route{
			strings.ToUpper("Get"),
			"/1.1/trunks/{trunk_id}",
			c.GetTrunk,
		},
		"ListTrunks": Route{
			strings.ToUpper("Get"),
			"/1.1/trunks",
			c.ListTrunks,
		},
		"UpdateTrunk": Route{
			strings.ToUpper("Put"),
			"/1.1/trunks/{trunk_id}",
			c.UpdateTrunk,
		},
	}
}

// AssociateOutcallTrunks - Associate outcall and trunks
func (c *TrunksAPIController) AssociateOutcallTrunks(w http.ResponseWriter, r *http.Request) {
	outcallIdParam, err := parseNumericParameter[int32](
		chi.URLParam(r, "outcall_id"),
		WithRequire[int32](parseInt32),
	)
	if err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	bodyParam := TrunksId{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&bodyParam); err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	if err := AssertTrunksIdRequired(bodyParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	if err := AssertTrunksIdConstraints(bodyParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	result, err := c.service.AssociateOutcallTrunks(r.Context(), outcallIdParam, bodyParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, result.Headers, w)
}

// AssociateTrunkEndpointCustom - Associate trunk and Custom endpoint
func (c *TrunksAPIController) AssociateTrunkEndpointCustom(w http.ResponseWriter, r *http.Request) {
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

// AssociateTrunkEndpointIax - Associate trunk and IAX endpoint
func (c *TrunksAPIController) AssociateTrunkEndpointIax(w http.ResponseWriter, r *http.Request) {
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

// AssociateTrunkEndpointSip - Associate trunk and SIP endpoint
func (c *TrunksAPIController) AssociateTrunkEndpointSip(w http.ResponseWriter, r *http.Request) {
	trunkIdParam, err := parseNumericParameter[int32](
		chi.URLParam(r, "trunk_id"),
		WithRequire[int32](parseInt32),
	)
	if err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	sipUuidParam := chi.URLParam(r, "sip_uuid")
	if sipUuidParam == "" {
		c.errorHandler(w, r, &RequiredError{"sip_uuid"}, nil)
		return
	}
	result, err := c.service.AssociateTrunkEndpointSip(r.Context(), trunkIdParam, sipUuidParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, result.Headers, w)
}

// AssociateTrunkRegisterIax - Associate trunk and IAX register
func (c *TrunksAPIController) AssociateTrunkRegisterIax(w http.ResponseWriter, r *http.Request) {
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

// CreateTrunk - Create trunk
func (c *TrunksAPIController) CreateTrunk(w http.ResponseWriter, r *http.Request) {
	bodyParam := Trunk{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&bodyParam); err != nil && !errors.Is(err, io.EOF) {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	if err := AssertTrunkRequired(bodyParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	if err := AssertTrunkConstraints(bodyParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	result, err := c.service.CreateTrunk(r.Context(), bodyParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, result.Headers, w)
}

// DeleteTrunk - Delete trunk
func (c *TrunksAPIController) DeleteTrunk(w http.ResponseWriter, r *http.Request) {
	trunkIdParam, err := parseNumericParameter[int32](
		chi.URLParam(r, "trunk_id"),
		WithRequire[int32](parseInt32),
	)
	if err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	accentTenantParam := r.Header.Get("Accent-Tenant")
	result, err := c.service.DeleteTrunk(r.Context(), trunkIdParam, accentTenantParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, result.Headers, w)
}

// DissociateTrunkEndpointCustom - Dissociate trunk and Custom endpoint
func (c *TrunksAPIController) DissociateTrunkEndpointCustom(w http.ResponseWriter, r *http.Request) {
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

// DissociateTrunkEndpointIax - Dissociate trunk and IAX endpoint
func (c *TrunksAPIController) DissociateTrunkEndpointIax(w http.ResponseWriter, r *http.Request) {
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

// DissociateTrunkEndpointSip - Dissociate trunk and SIP endpoint
func (c *TrunksAPIController) DissociateTrunkEndpointSip(w http.ResponseWriter, r *http.Request) {
	trunkIdParam, err := parseNumericParameter[int32](
		chi.URLParam(r, "trunk_id"),
		WithRequire[int32](parseInt32),
	)
	if err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	sipUuidParam := chi.URLParam(r, "sip_uuid")
	if sipUuidParam == "" {
		c.errorHandler(w, r, &RequiredError{"sip_uuid"}, nil)
		return
	}
	result, err := c.service.DissociateTrunkEndpointSip(r.Context(), trunkIdParam, sipUuidParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, result.Headers, w)
}

// DissociateTrunkRegisterIax - Dissociate trunk and IAX register
func (c *TrunksAPIController) DissociateTrunkRegisterIax(w http.ResponseWriter, r *http.Request) {
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

// GetTrunk - Get trunk
func (c *TrunksAPIController) GetTrunk(w http.ResponseWriter, r *http.Request) {
	trunkIdParam, err := parseNumericParameter[int32](
		chi.URLParam(r, "trunk_id"),
		WithRequire[int32](parseInt32),
	)
	if err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	accentTenantParam := r.Header.Get("Accent-Tenant")
	result, err := c.service.GetTrunk(r.Context(), trunkIdParam, accentTenantParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, result.Headers, w)
}

// ListTrunks - List trunks
func (c *TrunksAPIController) ListTrunks(w http.ResponseWriter, r *http.Request) {
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
	result, err := c.service.ListTrunks(r.Context(), accentTenantParam, recurseParam, orderParam, directionParam, limitParam, offsetParam, searchParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, result.Headers, w)
}

// UpdateTrunk - Update trunk
func (c *TrunksAPIController) UpdateTrunk(w http.ResponseWriter, r *http.Request) {
	trunkIdParam, err := parseNumericParameter[int32](
		chi.URLParam(r, "trunk_id"),
		WithRequire[int32](parseInt32),
	)
	if err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	bodyParam := Trunk{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&bodyParam); err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	if err := AssertTrunkRequired(bodyParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	if err := AssertTrunkConstraints(bodyParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	accentTenantParam := r.Header.Get("Accent-Tenant")
	result, err := c.service.UpdateTrunk(r.Context(), trunkIdParam, bodyParam, accentTenantParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, result.Headers, w)
}
