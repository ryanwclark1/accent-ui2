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

// ExportsAPIController binds http requests to an api service and writes the service results to the http response
type ExportsAPIController struct {
	service      ExportsAPIServicer
	errorHandler ErrorHandler
}

// ExportsAPIOption for how the controller is set up.
type ExportsAPIOption func(*ExportsAPIController)

// WithExportsAPIErrorHandler inject ErrorHandler into controller
func WithExportsAPIErrorHandler(h ErrorHandler) ExportsAPIOption {
	return func(c *ExportsAPIController) {
		c.errorHandler = h
	}
}

// NewExportsAPIController creates a default api controller
func NewExportsAPIController(s ExportsAPIServicer, opts ...ExportsAPIOption) Router {
	controller := &ExportsAPIController{
		service:      s,
		errorHandler: DefaultErrorHandler,
	}

	for _, opt := range opts {
		opt(controller)
	}

	return controller
}

// Routes returns all the api routes for the ExportsAPIController
func (c *ExportsAPIController) Routes() Routes {
	return Routes{
		"CreateCDRRecordingsMediaExport": Route{
			strings.ToUpper("Post"),
			"/1.0/cdr/recordings/media/export",
			c.CreateCDRRecordingsMediaExport,
		},
		"GetExport": Route{
			strings.ToUpper("Get"),
			"/1.0/exports/{export_uuid}",
			c.GetExport,
		},
		"GetExportDownload": Route{
			strings.ToUpper("Get"),
			"/1.0/exports/{export_uuid}/download",
			c.GetExportDownload,
		},
	}
}

// CreateCDRRecordingsMediaExport - Create an export for the recording media of multiple CDRs
func (c *ExportsAPIController) CreateCDRRecordingsMediaExport(w http.ResponseWriter, r *http.Request) {
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
	var tagsParam []string
	if query.Has("tags") {
		tagsParam = strings.Split(query.Get("tags"), ",")
	}
	var userUuidParam []string
	if query.Has("user_uuid") {
		userUuidParam = strings.Split(query.Get("user_uuid"), ",")
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
	accentTenantParam := r.Header.Get("Accent-Tenant")
	var emailParam string
	if query.Has("email") {
		param := query.Get("email")

		emailParam = param
	} else {
	}
	bodyParam := CreateCdrRecordingsMediaExportRequest{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&bodyParam); err != nil && !errors.Is(err, io.EOF) {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	if err := AssertCreateCdrRecordingsMediaExportRequestRequired(bodyParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	if err := AssertCreateCdrRecordingsMediaExportRequestConstraints(bodyParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	result, err := c.service.CreateCDRRecordingsMediaExport(r.Context(), fromParam, untilParam, searchParam, callDirectionParam, numberParam, tagsParam, userUuidParam, fromIdParam, recurseParam, accentTenantParam, emailParam, bodyParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, result.Headers, w)
}

// GetExport - Get an export by the given UUID
func (c *ExportsAPIController) GetExport(w http.ResponseWriter, r *http.Request) {
	exportUuidParam := chi.URLParam(r, "export_uuid")
	if exportUuidParam == "" {
		c.errorHandler(w, r, &RequiredError{"export_uuid"}, nil)
		return
	}
	result, err := c.service.GetExport(r.Context(), exportUuidParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, result.Headers, w)
}

// GetExportDownload - Download an export as a ZIP archive by the given UUID
func (c *ExportsAPIController) GetExportDownload(w http.ResponseWriter, r *http.Request) {
	exportUuidParam := chi.URLParam(r, "export_uuid")
	if exportUuidParam == "" {
		c.errorHandler(w, r, &RequiredError{"export_uuid"}, nil)
		return
	}
	result, err := c.service.GetExportDownload(r.Context(), exportUuidParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, result.Headers, w)
}
