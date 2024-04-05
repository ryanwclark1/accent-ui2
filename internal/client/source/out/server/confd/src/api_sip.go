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

// SipAPIController binds http requests to an api service and writes the service results to the http response
type SipAPIController struct {
	service      SipAPIServicer
	errorHandler ErrorHandler
}

// SipAPIOption for how the controller is set up.
type SipAPIOption func(*SipAPIController)

// WithSipAPIErrorHandler inject ErrorHandler into controller
func WithSipAPIErrorHandler(h ErrorHandler) SipAPIOption {
	return func(c *SipAPIController) {
		c.errorHandler = h
	}
}

// NewSipAPIController creates a default api controller
func NewSipAPIController(s SipAPIServicer, opts ...SipAPIOption) Router {
	controller := &SipAPIController{
		service:      s,
		errorHandler: DefaultErrorHandler,
	}

	for _, opt := range opts {
		opt(controller)
	}

	return controller
}

// Routes returns all the api routes for the SipAPIController
func (c *SipAPIController) Routes() Routes {
	return Routes{
		"AssociateLineEndpointSip": Route{
			strings.ToUpper("Put"),
			"/1.1/lines/{line_id}/endpoints/sip/{sip_uuid}",
			c.AssociateLineEndpointSip,
		},
		"AssociateTrunkEndpointSip": Route{
			strings.ToUpper("Put"),
			"/1.1/trunks/{trunk_id}/endpoints/sip/{sip_uuid}",
			c.AssociateTrunkEndpointSip,
		},
		"CreateEndpointSip": Route{
			strings.ToUpper("Post"),
			"/1.1/endpoints/sip",
			c.CreateEndpointSip,
		},
		"CreateEndpointSipTemplate": Route{
			strings.ToUpper("Post"),
			"/1.1/endpoints/sip/templates",
			c.CreateEndpointSipTemplate,
		},
		"CreateSipTransport": Route{
			strings.ToUpper("Post"),
			"/1.1/sip/transports",
			c.CreateSipTransport,
		},
		"DeleteEndpointSip": Route{
			strings.ToUpper("Delete"),
			"/1.1/endpoints/sip/{sip_uuid}",
			c.DeleteEndpointSip,
		},
		"DeleteEndpointSipTemplate": Route{
			strings.ToUpper("Delete"),
			"/1.1/endpoints/sip/templates/{template_uuid}",
			c.DeleteEndpointSipTemplate,
		},
		"DeleteSipTransport": Route{
			strings.ToUpper("Delete"),
			"/1.1/sip/transports/{transport_uuid}",
			c.DeleteSipTransport,
		},
		"DissociateLineEndpointSip": Route{
			strings.ToUpper("Delete"),
			"/1.1/lines/{line_id}/endpoints/sip/{sip_uuid}",
			c.DissociateLineEndpointSip,
		},
		"DissociateTrunkEndpointSip": Route{
			strings.ToUpper("Delete"),
			"/1.1/trunks/{trunk_id}/endpoints/sip/{sip_uuid}",
			c.DissociateTrunkEndpointSip,
		},
		"GetEndpointSip": Route{
			strings.ToUpper("Get"),
			"/1.1/endpoints/sip/{sip_uuid}",
			c.GetEndpointSip,
		},
		"GetEndpointSipTemplate": Route{
			strings.ToUpper("Get"),
			"/1.1/endpoints/sip/templates/{template_uuid}",
			c.GetEndpointSipTemplate,
		},
		"GetSipTransport": Route{
			strings.ToUpper("Get"),
			"/1.1/sip/transports/{transport_uuid}",
			c.GetSipTransport,
		},
		"ListAsteriskPjsipGlobal": Route{
			strings.ToUpper("Get"),
			"/1.1/asterisk/pjsip/global",
			c.ListAsteriskPjsipGlobal,
		},
		"ListAsteriskPjsipSystem": Route{
			strings.ToUpper("Get"),
			"/1.1/asterisk/pjsip/system",
			c.ListAsteriskPjsipSystem,
		},
		"ListEndpointsSip": Route{
			strings.ToUpper("Get"),
			"/1.1/endpoints/sip",
			c.ListEndpointsSip,
		},
		"ListEndpointsSipTemplates": Route{
			strings.ToUpper("Get"),
			"/1.1/endpoints/sip/templates",
			c.ListEndpointsSipTemplates,
		},
		"ListSipTransports": Route{
			strings.ToUpper("Get"),
			"/1.1/sip/transports",
			c.ListSipTransports,
		},
		"ShowPjsipDoc": Route{
			strings.ToUpper("Get"),
			"/1.1/asterisk/pjsip/doc",
			c.ShowPjsipDoc,
		},
		"UpdateAsteriskPjsipGlobal": Route{
			strings.ToUpper("Put"),
			"/1.1/asterisk/pjsip/global",
			c.UpdateAsteriskPjsipGlobal,
		},
		"UpdateAsteriskPjsipSystem": Route{
			strings.ToUpper("Put"),
			"/1.1/asterisk/pjsip/system",
			c.UpdateAsteriskPjsipSystem,
		},
		"UpdateEndpointSip": Route{
			strings.ToUpper("Put"),
			"/1.1/endpoints/sip/{sip_uuid}",
			c.UpdateEndpointSip,
		},
		"UpdateEndpointSipTemplate": Route{
			strings.ToUpper("Put"),
			"/1.1/endpoints/sip/templates/{template_uuid}",
			c.UpdateEndpointSipTemplate,
		},
		"UpdateSipTransport": Route{
			strings.ToUpper("Put"),
			"/1.1/sip/transports/{transport_uuid}",
			c.UpdateSipTransport,
		},
	}
}

// AssociateLineEndpointSip - Associate line and SIP endpoint
func (c *SipAPIController) AssociateLineEndpointSip(w http.ResponseWriter, r *http.Request) {
	lineIdParam, err := parseNumericParameter[int32](
		chi.URLParam(r, "line_id"),
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
	result, err := c.service.AssociateLineEndpointSip(r.Context(), lineIdParam, sipUuidParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}

// AssociateTrunkEndpointSip - Associate trunk and SIP endpoint
func (c *SipAPIController) AssociateTrunkEndpointSip(w http.ResponseWriter, r *http.Request) {
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
	EncodeJSONResponse(result.Body, &result.Code, w)
}

// CreateEndpointSip - Create a SIP endpoint
func (c *SipAPIController) CreateEndpointSip(w http.ResponseWriter, r *http.Request) {
	accentTenantParam := r.Header.Get("Accent-Tenant")
	bodyParam := EndpointSip{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&bodyParam); err != nil && !errors.Is(err, io.EOF) {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	if err := AssertEndpointSipRequired(bodyParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	if err := AssertEndpointSipConstraints(bodyParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	result, err := c.service.CreateEndpointSip(r.Context(), accentTenantParam, bodyParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}

// CreateEndpointSipTemplate - Create a SIP endpoint template
func (c *SipAPIController) CreateEndpointSipTemplate(w http.ResponseWriter, r *http.Request) {
	accentTenantParam := r.Header.Get("Accent-Tenant")
	bodyParam := EndpointSip{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&bodyParam); err != nil && !errors.Is(err, io.EOF) {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	if err := AssertEndpointSipRequired(bodyParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	if err := AssertEndpointSipConstraints(bodyParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	result, err := c.service.CreateEndpointSipTemplate(r.Context(), accentTenantParam, bodyParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}

// CreateSipTransport - Create SIP transport
func (c *SipAPIController) CreateSipTransport(w http.ResponseWriter, r *http.Request) {
	bodyParam := SipTransport{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&bodyParam); err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	if err := AssertSipTransportRequired(bodyParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	if err := AssertSipTransportConstraints(bodyParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	result, err := c.service.CreateSipTransport(r.Context(), bodyParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}

// DeleteEndpointSip - Delete SIP Endpoint
func (c *SipAPIController) DeleteEndpointSip(w http.ResponseWriter, r *http.Request) {
	sipUuidParam := chi.URLParam(r, "sip_uuid")
	if sipUuidParam == "" {
		c.errorHandler(w, r, &RequiredError{"sip_uuid"}, nil)
		return
	}
	accentTenantParam := r.Header.Get("Accent-Tenant")
	result, err := c.service.DeleteEndpointSip(r.Context(), sipUuidParam, accentTenantParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}

// DeleteEndpointSipTemplate - Delete SIP Endpoint Template
func (c *SipAPIController) DeleteEndpointSipTemplate(w http.ResponseWriter, r *http.Request) {
	templateUuidParam := chi.URLParam(r, "template_uuid")
	if templateUuidParam == "" {
		c.errorHandler(w, r, &RequiredError{"template_uuid"}, nil)
		return
	}
	accentTenantParam := r.Header.Get("Accent-Tenant")
	result, err := c.service.DeleteEndpointSipTemplate(r.Context(), templateUuidParam, accentTenantParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}

// DeleteSipTransport - Delete SIP transport
func (c *SipAPIController) DeleteSipTransport(w http.ResponseWriter, r *http.Request) {
	query, err := parseQuery(r.URL.RawQuery)
	if err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	transportUuidParam := chi.URLParam(r, "transport_uuid")
	if transportUuidParam == "" {
		c.errorHandler(w, r, &RequiredError{"transport_uuid"}, nil)
		return
	}
	var fallbackParam string
	if query.Has("fallback") {
		param := query.Get("fallback")

		fallbackParam = param
	} else {
	}
	result, err := c.service.DeleteSipTransport(r.Context(), transportUuidParam, fallbackParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}

// DissociateLineEndpointSip - Dissociate line and SIP endpoint
func (c *SipAPIController) DissociateLineEndpointSip(w http.ResponseWriter, r *http.Request) {
	lineIdParam, err := parseNumericParameter[int32](
		chi.URLParam(r, "line_id"),
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
	result, err := c.service.DissociateLineEndpointSip(r.Context(), lineIdParam, sipUuidParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}

// DissociateTrunkEndpointSip - Dissociate trunk and SIP endpoint
func (c *SipAPIController) DissociateTrunkEndpointSip(w http.ResponseWriter, r *http.Request) {
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
	EncodeJSONResponse(result.Body, &result.Code, w)
}

// GetEndpointSip - Get SIP Endpoint
func (c *SipAPIController) GetEndpointSip(w http.ResponseWriter, r *http.Request) {
	query, err := parseQuery(r.URL.RawQuery)
	if err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	sipUuidParam := chi.URLParam(r, "sip_uuid")
	if sipUuidParam == "" {
		c.errorHandler(w, r, &RequiredError{"sip_uuid"}, nil)
		return
	}
	accentTenantParam := r.Header.Get("Accent-Tenant")
	var viewParam string
	if query.Has("view") {
		param := query.Get("view")

		viewParam = param
	} else {
	}
	result, err := c.service.GetEndpointSip(r.Context(), sipUuidParam, accentTenantParam, viewParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}

// GetEndpointSipTemplate - Get SIP Endpoint template
func (c *SipAPIController) GetEndpointSipTemplate(w http.ResponseWriter, r *http.Request) {
	templateUuidParam := chi.URLParam(r, "template_uuid")
	if templateUuidParam == "" {
		c.errorHandler(w, r, &RequiredError{"template_uuid"}, nil)
		return
	}
	accentTenantParam := r.Header.Get("Accent-Tenant")
	result, err := c.service.GetEndpointSipTemplate(r.Context(), templateUuidParam, accentTenantParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}

// GetSipTransport - Get SIP transport
func (c *SipAPIController) GetSipTransport(w http.ResponseWriter, r *http.Request) {
	transportUuidParam := chi.URLParam(r, "transport_uuid")
	if transportUuidParam == "" {
		c.errorHandler(w, r, &RequiredError{"transport_uuid"}, nil)
		return
	}
	result, err := c.service.GetSipTransport(r.Context(), transportUuidParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}

// ListAsteriskPjsipGlobal - List of PJSIP options for the `global` section
func (c *SipAPIController) ListAsteriskPjsipGlobal(w http.ResponseWriter, r *http.Request) {
	result, err := c.service.ListAsteriskPjsipGlobal(r.Context())
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}

// ListAsteriskPjsipSystem - List of PJSIP options for the `system` section
func (c *SipAPIController) ListAsteriskPjsipSystem(w http.ResponseWriter, r *http.Request) {
	result, err := c.service.ListAsteriskPjsipSystem(r.Context())
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}

// ListEndpointsSip - List SIP endpoints
func (c *SipAPIController) ListEndpointsSip(w http.ResponseWriter, r *http.Request) {
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
	result, err := c.service.ListEndpointsSip(r.Context(), accentTenantParam, recurseParam, orderParam, directionParam, limitParam, offsetParam, searchParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}

// ListEndpointsSipTemplates - List SIP endpoints templates
func (c *SipAPIController) ListEndpointsSipTemplates(w http.ResponseWriter, r *http.Request) {
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
	result, err := c.service.ListEndpointsSipTemplates(r.Context(), accentTenantParam, recurseParam, orderParam, directionParam, limitParam, offsetParam, searchParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}

// ListSipTransports - List all configured SIP transports
func (c *SipAPIController) ListSipTransports(w http.ResponseWriter, r *http.Request) {
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
	result, err := c.service.ListSipTransports(r.Context(), orderParam, directionParam, limitParam, offsetParam, searchParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}

// ShowPjsipDoc - List all PJSIP configuration options
func (c *SipAPIController) ShowPjsipDoc(w http.ResponseWriter, r *http.Request) {
	result, err := c.service.ShowPjsipDoc(r.Context())
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}

// UpdateAsteriskPjsipGlobal - Update PJSIP section options
func (c *SipAPIController) UpdateAsteriskPjsipGlobal(w http.ResponseWriter, r *http.Request) {
	bodyParam := PjsipGlobal{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&bodyParam); err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	if err := AssertPjsipGlobalRequired(bodyParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	if err := AssertPjsipGlobalConstraints(bodyParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	result, err := c.service.UpdateAsteriskPjsipGlobal(r.Context(), bodyParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}

// UpdateAsteriskPjsipSystem - Update PJSIP section options
func (c *SipAPIController) UpdateAsteriskPjsipSystem(w http.ResponseWriter, r *http.Request) {
	bodyParam := PjsipSystem{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&bodyParam); err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	if err := AssertPjsipSystemRequired(bodyParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	if err := AssertPjsipSystemConstraints(bodyParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	result, err := c.service.UpdateAsteriskPjsipSystem(r.Context(), bodyParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}

// UpdateEndpointSip - Update SIP Endpoint
func (c *SipAPIController) UpdateEndpointSip(w http.ResponseWriter, r *http.Request) {
	sipUuidParam := chi.URLParam(r, "sip_uuid")
	if sipUuidParam == "" {
		c.errorHandler(w, r, &RequiredError{"sip_uuid"}, nil)
		return
	}
	bodyParam := EndpointSip{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&bodyParam); err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	if err := AssertEndpointSipRequired(bodyParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	if err := AssertEndpointSipConstraints(bodyParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	accentTenantParam := r.Header.Get("Accent-Tenant")
	result, err := c.service.UpdateEndpointSip(r.Context(), sipUuidParam, bodyParam, accentTenantParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}

// UpdateEndpointSipTemplate - Update SIP Endpoint Template
func (c *SipAPIController) UpdateEndpointSipTemplate(w http.ResponseWriter, r *http.Request) {
	templateUuidParam := chi.URLParam(r, "template_uuid")
	if templateUuidParam == "" {
		c.errorHandler(w, r, &RequiredError{"template_uuid"}, nil)
		return
	}
	bodyParam := EndpointSip{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&bodyParam); err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	if err := AssertEndpointSipRequired(bodyParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	if err := AssertEndpointSipConstraints(bodyParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	accentTenantParam := r.Header.Get("Accent-Tenant")
	result, err := c.service.UpdateEndpointSipTemplate(r.Context(), templateUuidParam, bodyParam, accentTenantParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}

// UpdateSipTransport - Update SIP transport
func (c *SipAPIController) UpdateSipTransport(w http.ResponseWriter, r *http.Request) {
	transportUuidParam := chi.URLParam(r, "transport_uuid")
	if transportUuidParam == "" {
		c.errorHandler(w, r, &RequiredError{"transport_uuid"}, nil)
		return
	}
	bodyParam := SipTransport{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&bodyParam); err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	if err := AssertSipTransportRequired(bodyParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	if err := AssertSipTransportConstraints(bodyParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	result, err := c.service.UpdateSipTransport(r.Context(), transportUuidParam, bodyParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}
