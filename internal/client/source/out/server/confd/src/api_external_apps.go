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

// ExternalAppsAPIController binds http requests to an api service and writes the service results to the http response
type ExternalAppsAPIController struct {
	service      ExternalAppsAPIServicer
	errorHandler ErrorHandler
}

// ExternalAppsAPIOption for how the controller is set up.
type ExternalAppsAPIOption func(*ExternalAppsAPIController)

// WithExternalAppsAPIErrorHandler inject ErrorHandler into controller
func WithExternalAppsAPIErrorHandler(h ErrorHandler) ExternalAppsAPIOption {
	return func(c *ExternalAppsAPIController) {
		c.errorHandler = h
	}
}

// NewExternalAppsAPIController creates a default api controller
func NewExternalAppsAPIController(s ExternalAppsAPIServicer, opts ...ExternalAppsAPIOption) Router {
	controller := &ExternalAppsAPIController{
		service:      s,
		errorHandler: DefaultErrorHandler,
	}

	for _, opt := range opts {
		opt(controller)
	}

	return controller
}

// Routes returns all the api routes for the ExternalAppsAPIController
func (c *ExternalAppsAPIController) Routes() Routes {
	return Routes{
		"CreateExternalApp": Route{
			strings.ToUpper("Post"),
			"/1.1/external/apps/{app_name}",
			c.CreateExternalApp,
		},
		"CreateUserExternalApp": Route{
			strings.ToUpper("Post"),
			"/1.1/users/{user_uuid}/external/apps/{app_name}",
			c.CreateUserExternalApp,
		},
		"DeleteExternalApp": Route{
			strings.ToUpper("Delete"),
			"/1.1/external/apps/{app_name}",
			c.DeleteExternalApp,
		},
		"DeleteUserExternalApp": Route{
			strings.ToUpper("Delete"),
			"/1.1/users/{user_uuid}/external/apps/{app_name}",
			c.DeleteUserExternalApp,
		},
		"GetExternalApp": Route{
			strings.ToUpper("Get"),
			"/1.1/external/apps/{app_name}",
			c.GetExternalApp,
		},
		"GetUserExternalApp": Route{
			strings.ToUpper("Get"),
			"/1.1/users/{user_uuid}/external/apps/{app_name}",
			c.GetUserExternalApp,
		},
		"ListExternalApps": Route{
			strings.ToUpper("Get"),
			"/1.1/external/apps",
			c.ListExternalApps,
		},
		"ListUserExternalApps": Route{
			strings.ToUpper("Get"),
			"/1.1/users/{user_uuid}/external/apps",
			c.ListUserExternalApps,
		},
		"UpdateExternalApp": Route{
			strings.ToUpper("Put"),
			"/1.1/external/apps/{app_name}",
			c.UpdateExternalApp,
		},
		"UpdateUserExternalApp": Route{
			strings.ToUpper("Put"),
			"/1.1/users/{user_uuid}/external/apps/{app_name}",
			c.UpdateUserExternalApp,
		},
	}
}

// CreateExternalApp - Create external app
func (c *ExternalAppsAPIController) CreateExternalApp(w http.ResponseWriter, r *http.Request) {
	appNameParam := chi.URLParam(r, "app_name")
	if appNameParam == "" {
		c.errorHandler(w, r, &RequiredError{"app_name"}, nil)
		return
	}
	bodyParam := ExternalApp{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&bodyParam); err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	if err := AssertExternalAppRequired(bodyParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	if err := AssertExternalAppConstraints(bodyParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	accentTenantParam := r.Header.Get("Accent-Tenant")
	result, err := c.service.CreateExternalApp(r.Context(), appNameParam, bodyParam, accentTenantParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, result.Headers, w)
}

// CreateUserExternalApp - Create user external app
func (c *ExternalAppsAPIController) CreateUserExternalApp(w http.ResponseWriter, r *http.Request) {
	userUuidParam := chi.URLParam(r, "user_uuid")
	if userUuidParam == "" {
		c.errorHandler(w, r, &RequiredError{"user_uuid"}, nil)
		return
	}
	appNameParam := chi.URLParam(r, "app_name")
	if appNameParam == "" {
		c.errorHandler(w, r, &RequiredError{"app_name"}, nil)
		return
	}
	bodyParam := UserExternalApp{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&bodyParam); err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	if err := AssertUserExternalAppRequired(bodyParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	if err := AssertUserExternalAppConstraints(bodyParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	accentTenantParam := r.Header.Get("Accent-Tenant")
	result, err := c.service.CreateUserExternalApp(r.Context(), userUuidParam, appNameParam, bodyParam, accentTenantParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, result.Headers, w)
}

// DeleteExternalApp - Delete external app
func (c *ExternalAppsAPIController) DeleteExternalApp(w http.ResponseWriter, r *http.Request) {
	appNameParam := chi.URLParam(r, "app_name")
	if appNameParam == "" {
		c.errorHandler(w, r, &RequiredError{"app_name"}, nil)
		return
	}
	accentTenantParam := r.Header.Get("Accent-Tenant")
	result, err := c.service.DeleteExternalApp(r.Context(), appNameParam, accentTenantParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, result.Headers, w)
}

// DeleteUserExternalApp - Delete user external app
func (c *ExternalAppsAPIController) DeleteUserExternalApp(w http.ResponseWriter, r *http.Request) {
	userUuidParam := chi.URLParam(r, "user_uuid")
	if userUuidParam == "" {
		c.errorHandler(w, r, &RequiredError{"user_uuid"}, nil)
		return
	}
	appNameParam := chi.URLParam(r, "app_name")
	if appNameParam == "" {
		c.errorHandler(w, r, &RequiredError{"app_name"}, nil)
		return
	}
	accentTenantParam := r.Header.Get("Accent-Tenant")
	result, err := c.service.DeleteUserExternalApp(r.Context(), userUuidParam, appNameParam, accentTenantParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, result.Headers, w)
}

// GetExternalApp - Get external app
func (c *ExternalAppsAPIController) GetExternalApp(w http.ResponseWriter, r *http.Request) {
	appNameParam := chi.URLParam(r, "app_name")
	if appNameParam == "" {
		c.errorHandler(w, r, &RequiredError{"app_name"}, nil)
		return
	}
	accentTenantParam := r.Header.Get("Accent-Tenant")
	result, err := c.service.GetExternalApp(r.Context(), appNameParam, accentTenantParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, result.Headers, w)
}

// GetUserExternalApp - Get user external app
func (c *ExternalAppsAPIController) GetUserExternalApp(w http.ResponseWriter, r *http.Request) {
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
	appNameParam := chi.URLParam(r, "app_name")
	if appNameParam == "" {
		c.errorHandler(w, r, &RequiredError{"app_name"}, nil)
		return
	}
	accentTenantParam := r.Header.Get("Accent-Tenant")
	var viewParam string
	if query.Has("view") {
		param := query.Get("view")

		viewParam = param
	} else {
	}
	result, err := c.service.GetUserExternalApp(r.Context(), userUuidParam, appNameParam, accentTenantParam, viewParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, result.Headers, w)
}

// ListExternalApps - List external apps
func (c *ExternalAppsAPIController) ListExternalApps(w http.ResponseWriter, r *http.Request) {
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
	result, err := c.service.ListExternalApps(r.Context(), accentTenantParam, recurseParam, orderParam, directionParam, limitParam, offsetParam, searchParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, result.Headers, w)
}

// ListUserExternalApps - List user external apps
func (c *ExternalAppsAPIController) ListUserExternalApps(w http.ResponseWriter, r *http.Request) {
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
	accentTenantParam := r.Header.Get("Accent-Tenant")
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
	var viewParam string
	if query.Has("view") {
		param := query.Get("view")

		viewParam = param
	} else {
	}
	result, err := c.service.ListUserExternalApps(r.Context(), userUuidParam, accentTenantParam, orderParam, directionParam, limitParam, offsetParam, searchParam, viewParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, result.Headers, w)
}

// UpdateExternalApp - Update external app
func (c *ExternalAppsAPIController) UpdateExternalApp(w http.ResponseWriter, r *http.Request) {
	appNameParam := chi.URLParam(r, "app_name")
	if appNameParam == "" {
		c.errorHandler(w, r, &RequiredError{"app_name"}, nil)
		return
	}
	bodyParam := ExternalApp{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&bodyParam); err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	if err := AssertExternalAppRequired(bodyParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	if err := AssertExternalAppConstraints(bodyParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	accentTenantParam := r.Header.Get("Accent-Tenant")
	result, err := c.service.UpdateExternalApp(r.Context(), appNameParam, bodyParam, accentTenantParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, result.Headers, w)
}

// UpdateUserExternalApp - Update user external app
func (c *ExternalAppsAPIController) UpdateUserExternalApp(w http.ResponseWriter, r *http.Request) {
	userUuidParam := chi.URLParam(r, "user_uuid")
	if userUuidParam == "" {
		c.errorHandler(w, r, &RequiredError{"user_uuid"}, nil)
		return
	}
	appNameParam := chi.URLParam(r, "app_name")
	if appNameParam == "" {
		c.errorHandler(w, r, &RequiredError{"app_name"}, nil)
		return
	}
	bodyParam := UserExternalApp{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&bodyParam); err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	if err := AssertUserExternalAppRequired(bodyParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	if err := AssertUserExternalAppConstraints(bodyParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	accentTenantParam := r.Header.Get("Accent-Tenant")
	result, err := c.service.UpdateUserExternalApp(r.Context(), userUuidParam, appNameParam, bodyParam, accentTenantParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, result.Headers, w)
}
