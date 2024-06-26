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

// CdrAPIController binds http requests to an api service and writes the service results to the http response
type CdrAPIController struct {
	service      CdrAPIServicer
	errorHandler ErrorHandler
}

// CdrAPIOption for how the controller is set up.
type CdrAPIOption func(*CdrAPIController)

// WithCdrAPIErrorHandler inject ErrorHandler into controller
func WithCdrAPIErrorHandler(h ErrorHandler) CdrAPIOption {
	return func(c *CdrAPIController) {
		c.errorHandler = h
	}
}

// NewCdrAPIController creates a default api controller
func NewCdrAPIController(s CdrAPIServicer, opts ...CdrAPIOption) Router {
	controller := &CdrAPIController{
		service:      s,
		errorHandler: DefaultErrorHandler,
	}

	for _, opt := range opts {
		opt(controller)
	}

	return controller
}

// Routes returns all the api routes for the CdrAPIController
func (c *CdrAPIController) Routes() Routes {
	return Routes{
		"CreateCDRRecordingsMediaExport": Route{
			strings.ToUpper("Post"),
			"/1.0/cdr/recordings/media/export",
			c.CreateCDRRecordingsMediaExport,
		},
		"DeleteCDRRecordingMedia": Route{
			strings.ToUpper("Delete"),
			"/1.0/cdr/{cdr_id}/recordings/{recording_uuid}/media",
			c.DeleteCDRRecordingMedia,
		},
		"DeleteCDRRecordingsMedia": Route{
			strings.ToUpper("Delete"),
			"/1.0/cdr/recordings/media",
			c.DeleteCDRRecordingsMedia,
		},
		"GetCDR": Route{
			strings.ToUpper("Get"),
			"/1.0/cdr/{cdr_id}",
			c.GetCDR,
		},
		"GetCDRRecordingMedia": Route{
			strings.ToUpper("Get"),
			"/1.0/cdr/{cdr_id}/recordings/{recording_uuid}/media",
			c.GetCDRRecordingMedia,
		},
		"GetCDRs": Route{
			strings.ToUpper("Get"),
			"/1.0/cdr",
			c.GetCDRs,
		},
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

// CreateCDRRecordingsMediaExport - Create an export for the recording media of multiple CDRs
func (c *CdrAPIController) CreateCDRRecordingsMediaExport(w http.ResponseWriter, r *http.Request) {
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

// DeleteCDRRecordingMedia - Delete a recording media
func (c *CdrAPIController) DeleteCDRRecordingMedia(w http.ResponseWriter, r *http.Request) {
	cdrIdParam, err := parseNumericParameter[int32](
		chi.URLParam(r, "cdr_id"),
		WithRequire[int32](parseInt32),
	)
	if err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	recordingUuidParam, err := parseNumericParameter[int32](
		chi.URLParam(r, "recording_uuid"),
		WithRequire[int32](parseInt32),
	)
	if err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	result, err := c.service.DeleteCDRRecordingMedia(r.Context(), cdrIdParam, recordingUuidParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, result.Headers, w)
}

// DeleteCDRRecordingsMedia - Delete multiple CDRs recording media
func (c *CdrAPIController) DeleteCDRRecordingsMedia(w http.ResponseWriter, r *http.Request) {
	bodyParam := DeleteCdrRecordingsMediaRequest{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&bodyParam); err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	if err := AssertDeleteCdrRecordingsMediaRequestRequired(bodyParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	if err := AssertDeleteCdrRecordingsMediaRequestConstraints(bodyParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	result, err := c.service.DeleteCDRRecordingsMedia(r.Context(), bodyParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, result.Headers, w)
}

// GetCDR - Get a CDR by ID
func (c *CdrAPIController) GetCDR(w http.ResponseWriter, r *http.Request) {
	cdrIdParam, err := parseNumericParameter[int32](
		chi.URLParam(r, "cdr_id"),
		WithRequire[int32](parseInt32),
	)
	if err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	result, err := c.service.GetCDR(r.Context(), cdrIdParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, result.Headers, w)
}

// GetCDRRecordingMedia - Get a recording media
func (c *CdrAPIController) GetCDRRecordingMedia(w http.ResponseWriter, r *http.Request) {
	cdrIdParam, err := parseNumericParameter[int32](
		chi.URLParam(r, "cdr_id"),
		WithRequire[int32](parseInt32),
	)
	if err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	recordingUuidParam, err := parseNumericParameter[int32](
		chi.URLParam(r, "recording_uuid"),
		WithRequire[int32](parseInt32),
	)
	if err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	result, err := c.service.GetCDRRecordingMedia(r.Context(), cdrIdParam, recordingUuidParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, result.Headers, w)
}

// GetCDRs - List CDR
func (c *CdrAPIController) GetCDRs(w http.ResponseWriter, r *http.Request) {
	query, err := parseQuery(r.URL.RawQuery)
	if err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	accentTenantParam := r.Header.Get("Accent-Tenant")
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
	result, err := c.service.GetCDRs(r.Context(), accentTenantParam, fromParam, untilParam, limitParam, offsetParam, orderParam, directionParam, searchParam, callDirectionParam, numberParam, tagsParam, userUuidParam, fromIdParam, recurseParam, distinctParam, recordedParam, formatParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, result.Headers, w)
}

// GetCurrentUserCDR - List CDR of the authenticated user
func (c *CdrAPIController) GetCurrentUserCDR(w http.ResponseWriter, r *http.Request) {
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
	EncodeJSONResponse(result.Body, &result.Code, result.Headers, w)
}

// GetUserCDR - List CDR of the given user
func (c *CdrAPIController) GetUserCDR(w http.ResponseWriter, r *http.Request) {
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
	EncodeJSONResponse(result.Body, &result.Code, result.Headers, w)
}
