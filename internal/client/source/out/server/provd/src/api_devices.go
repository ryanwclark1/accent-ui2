/*
 * accent-provd
 *
 * Provisioning application REST API
 *
 * API version: 0.2
 * Contact: help@accentvoice.io
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package provdserver

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5"
)

// DevicesAPIController binds http requests to an api service and writes the service results to the http response
type DevicesAPIController struct {
	service      DevicesAPIServicer
	errorHandler ErrorHandler
}

// DevicesAPIOption for how the controller is set up.
type DevicesAPIOption func(*DevicesAPIController)

// WithDevicesAPIErrorHandler inject ErrorHandler into controller
func WithDevicesAPIErrorHandler(h ErrorHandler) DevicesAPIOption {
	return func(c *DevicesAPIController) {
		c.errorHandler = h
	}
}

// NewDevicesAPIController creates a default api controller
func NewDevicesAPIController(s DevicesAPIServicer, opts ...DevicesAPIOption) Router {
	controller := &DevicesAPIController{
		service:      s,
		errorHandler: DefaultErrorHandler,
	}

	for _, opt := range opts {
		opt(controller)
	}

	return controller
}

// Routes returns all the api routes for the DevicesAPIController
func (c *DevicesAPIController) Routes() Routes {
	return Routes{
		"DeleteDevMgrDevice": Route{
			strings.ToUpper("Delete"),
			"/0.2/dev_mgr/devices/{device_id}",
			c.DeleteDevMgrDevice,
		},
		"DeleteDevMgrSynchronize": Route{
			strings.ToUpper("Delete"),
			"/0.2/dev_mgr/synchronize/{operation_id}",
			c.DeleteDevMgrSynchronize,
		},
		"DevMgrDevicesPost": Route{
			strings.ToUpper("Post"),
			"/0.2/dev_mgr/devices",
			c.DevMgrDevicesPost,
		},
		"GetDevMgr": Route{
			strings.ToUpper("Get"),
			"/0.2/dev_mgr",
			c.GetDevMgr,
		},
		"GetDevMgrDevice": Route{
			strings.ToUpper("Get"),
			"/0.2/dev_mgr/devices/{device_id}",
			c.GetDevMgrDevice,
		},
		"GetDevMgrDevices": Route{
			strings.ToUpper("Get"),
			"/0.2/dev_mgr/devices",
			c.GetDevMgrDevices,
		},
		"GetDevMgrSynchronize": Route{
			strings.ToUpper("Get"),
			"/0.2/dev_mgr/synchronize/{operation_id}",
			c.GetDevMgrSynchronize,
		},
		"PostDevMgrDhcpinfo": Route{
			strings.ToUpper("Post"),
			"/0.2/dev_mgr/dhcpinfo",
			c.PostDevMgrDhcpinfo,
		},
		"PostDevMgrReconfigure": Route{
			strings.ToUpper("Post"),
			"/0.2/dev_mgr/reconfigure",
			c.PostDevMgrReconfigure,
		},
		"PostDevMgrSynchronize": Route{
			strings.ToUpper("Post"),
			"/0.2/dev_mgr/synchronize",
			c.PostDevMgrSynchronize,
		},
		"PutDevMgrDevice": Route{
			strings.ToUpper("Put"),
			"/0.2/dev_mgr/devices/{device_id}",
			c.PutDevMgrDevice,
		},
	}
}

// DeleteDevMgrDevice - Delete a device
func (c *DevicesAPIController) DeleteDevMgrDevice(w http.ResponseWriter, r *http.Request) {
	deviceIdParam := chi.URLParam(r, "device_id")
	if deviceIdParam == "" {
		c.errorHandler(w, r, &RequiredError{"device_id"}, nil)
		return
	}
	accentTenantParam := r.Header.Get("Accent-Tenant")
	result, err := c.service.DeleteDevMgrDevice(r.Context(), deviceIdParam, accentTenantParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, result.Headers, w)
}

// DeleteDevMgrSynchronize - Delete the Operation In Progress
func (c *DevicesAPIController) DeleteDevMgrSynchronize(w http.ResponseWriter, r *http.Request) {
	operationIdParam := chi.URLParam(r, "operation_id")
	if operationIdParam == "" {
		c.errorHandler(w, r, &RequiredError{"operation_id"}, nil)
		return
	}
	result, err := c.service.DeleteDevMgrSynchronize(r.Context(), operationIdParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, result.Headers, w)
}

// DevMgrDevicesPost - Create a device
func (c *DevicesAPIController) DevMgrDevicesPost(w http.ResponseWriter, r *http.Request) {
	accentTenantParam := r.Header.Get("Accent-Tenant")
	deviceParam := DeviceObject{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&deviceParam); err != nil && !errors.Is(err, io.EOF) {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	if err := AssertDeviceObjectRequired(deviceParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	if err := AssertDeviceObjectConstraints(deviceParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	result, err := c.service.DevMgrDevicesPost(r.Context(), accentTenantParam, deviceParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, result.Headers, w)
}

// GetDevMgr - Get the Device Manager resource
func (c *DevicesAPIController) GetDevMgr(w http.ResponseWriter, r *http.Request) {
	result, err := c.service.GetDevMgr(r.Context())
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, result.Headers, w)
}

// GetDevMgrDevice - Get a device by ID
func (c *DevicesAPIController) GetDevMgrDevice(w http.ResponseWriter, r *http.Request) {
	deviceIdParam := chi.URLParam(r, "device_id")
	if deviceIdParam == "" {
		c.errorHandler(w, r, &RequiredError{"device_id"}, nil)
		return
	}
	accentTenantParam := r.Header.Get("Accent-Tenant")
	result, err := c.service.GetDevMgrDevice(r.Context(), deviceIdParam, accentTenantParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, result.Headers, w)
}

// GetDevMgrDevices - List and find devices
func (c *DevicesAPIController) GetDevMgrDevices(w http.ResponseWriter, r *http.Request) {
	query, err := parseQuery(r.URL.RawQuery)
	if err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	var qParam string
	if query.Has("q") {
		param := query.Get("q")

		qParam = param
	} else {
	}
	var fieldsParam string
	if query.Has("fields") {
		param := query.Get("fields")

		fieldsParam = param
	} else {
	}
	var skipParam int32
	if query.Has("skip") {
		param, err := parseNumericParameter[int32](
			query.Get("skip"),
			WithParse[int32](parseInt32),
		)
		if err != nil {
			c.errorHandler(w, r, &ParsingError{Err: err}, nil)
			return
		}

		skipParam = param
	} else {
	}
	var sortParam string
	if query.Has("sort") {
		param := query.Get("sort")

		sortParam = param
	} else {
	}
	var sortOrdParam string
	if query.Has("sort_ord") {
		param := query.Get("sort_ord")

		sortOrdParam = param
	} else {
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
	result, err := c.service.GetDevMgrDevices(r.Context(), qParam, fieldsParam, skipParam, sortParam, sortOrdParam, accentTenantParam, recurseParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, result.Headers, w)
}

// GetDevMgrSynchronize - Get the status of a synchronize Operation In Progress
func (c *DevicesAPIController) GetDevMgrSynchronize(w http.ResponseWriter, r *http.Request) {
	operationIdParam := chi.URLParam(r, "operation_id")
	if operationIdParam == "" {
		c.errorHandler(w, r, &RequiredError{"operation_id"}, nil)
		return
	}
	result, err := c.service.GetDevMgrSynchronize(r.Context(), operationIdParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, result.Headers, w)
}

// PostDevMgrDhcpinfo - Push DHCP request information
func (c *DevicesAPIController) PostDevMgrDhcpinfo(w http.ResponseWriter, r *http.Request) {
	accentTenantParam := r.Header.Get("Accent-Tenant")
	bodyParam := DhcpInfoObject{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&bodyParam); err != nil && !errors.Is(err, io.EOF) {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	if err := AssertDhcpInfoObjectRequired(bodyParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	if err := AssertDhcpInfoObjectConstraints(bodyParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	result, err := c.service.PostDevMgrDhcpinfo(r.Context(), accentTenantParam, bodyParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, result.Headers, w)
}

// PostDevMgrReconfigure - Reconfigure a device
func (c *DevicesAPIController) PostDevMgrReconfigure(w http.ResponseWriter, r *http.Request) {
	accentTenantParam := r.Header.Get("Accent-Tenant")
	bodyParam := IdObject{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&bodyParam); err != nil && !errors.Is(err, io.EOF) {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	if err := AssertIdObjectRequired(bodyParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	if err := AssertIdObjectConstraints(bodyParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	result, err := c.service.PostDevMgrReconfigure(r.Context(), accentTenantParam, bodyParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, result.Headers, w)
}

// PostDevMgrSynchronize - Synchronize a device
func (c *DevicesAPIController) PostDevMgrSynchronize(w http.ResponseWriter, r *http.Request) {
	accentTenantParam := r.Header.Get("Accent-Tenant")
	bodyParam := IdObject{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&bodyParam); err != nil && !errors.Is(err, io.EOF) {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	if err := AssertIdObjectRequired(bodyParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	if err := AssertIdObjectConstraints(bodyParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	result, err := c.service.PostDevMgrSynchronize(r.Context(), accentTenantParam, bodyParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, result.Headers, w)
}

// PutDevMgrDevice - Update a device
func (c *DevicesAPIController) PutDevMgrDevice(w http.ResponseWriter, r *http.Request) {
	deviceIdParam := chi.URLParam(r, "device_id")
	if deviceIdParam == "" {
		c.errorHandler(w, r, &RequiredError{"device_id"}, nil)
		return
	}
	accentTenantParam := r.Header.Get("Accent-Tenant")
	deviceParam := DeviceObject{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&deviceParam); err != nil && !errors.Is(err, io.EOF) {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	if err := AssertDeviceObjectRequired(deviceParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	if err := AssertDeviceObjectConstraints(deviceParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	result, err := c.service.PutDevMgrDevice(r.Context(), deviceIdParam, accentTenantParam, deviceParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, result.Headers, w)
}
