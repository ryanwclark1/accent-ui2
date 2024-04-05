/*
 * accent-auth
 *
 * Accent's authentication service
 *
 * API version: 0.1
 * Contact: help@accentvoice.io
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package authserver

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5"
)

// UsersAPIController binds http requests to an api service and writes the service results to the http response
type UsersAPIController struct {
	service      UsersAPIServicer
	errorHandler ErrorHandler
}

// UsersAPIOption for how the controller is set up.
type UsersAPIOption func(*UsersAPIController)

// WithUsersAPIErrorHandler inject ErrorHandler into controller
func WithUsersAPIErrorHandler(h ErrorHandler) UsersAPIOption {
	return func(c *UsersAPIController) {
		c.errorHandler = h
	}
}

// NewUsersAPIController creates a default api controller
func NewUsersAPIController(s UsersAPIServicer, opts ...UsersAPIOption) Router {
	controller := &UsersAPIController{
		service:      s,
		errorHandler: DefaultErrorHandler,
	}

	for _, opt := range opts {
		opt(controller)
	}

	return controller
}

// Routes returns all the api routes for the UsersAPIController
func (c *UsersAPIController) Routes() Routes {
	return Routes{
		"AddGroupUser": Route{
			strings.ToUpper("Put"),
			"/0.1/groups/{group_uuid}/users/{user_uuid}",
			c.AddGroupUser,
		},
		"AddUserPolicy": Route{
			strings.ToUpper("Put"),
			"/0.1/users/{user_uuid}/policies/{policy_uuid}",
			c.AddUserPolicy,
		},
		"ChangeUserPassword": Route{
			strings.ToUpper("Put"),
			"/0.1/users/{user_uuid}/password",
			c.ChangeUserPassword,
		},
		"CreateUser": Route{
			strings.ToUpper("Post"),
			"/0.1/users",
			c.CreateUser,
		},
		"DeleteRefreshTokens": Route{
			strings.ToUpper("Delete"),
			"/0.1/users/{user_uuid_or_me}/tokens/{client_id}",
			c.DeleteRefreshTokens,
		},
		"DeleteUser": Route{
			strings.ToUpper("Delete"),
			"/0.1/users/{user_uuid}",
			c.DeleteUser,
		},
		"DeleteUserPolicy": Route{
			strings.ToUpper("Delete"),
			"/0.1/users/{user_uuid}/policies/{policy_uuid}",
			c.DeleteUserPolicy,
		},
		"GetNewEmailConfirmation": Route{
			strings.ToUpper("Get"),
			"/0.1/users/{user_uuid}/emails/{email_uuid}/confirm",
			c.GetNewEmailConfirmation,
		},
		"GetUser": Route{
			strings.ToUpper("Get"),
			"/0.1/users/{user_uuid}",
			c.GetUser,
		},
		"GetUserExternalAuth": Route{
			strings.ToUpper("Get"),
			"/0.1/users/{user_uuid}/external",
			c.GetUserExternalAuth,
		},
		"GetUserGroups": Route{
			strings.ToUpper("Get"),
			"/0.1/users/{user_uuid}/groups",
			c.GetUserGroups,
		},
		"GetUserPolicies": Route{
			strings.ToUpper("Get"),
			"/0.1/users/{user_uuid}/policies",
			c.GetUserPolicies,
		},
		"GetUserSessions": Route{
			strings.ToUpper("Get"),
			"/0.1/users/{user_uuid}/sessions",
			c.GetUserSessions,
		},
		"GetUserTokens": Route{
			strings.ToUpper("Get"),
			"/0.1/users/{user_uuid_or_me}/tokens",
			c.GetUserTokens,
		},
		"GetUsers": Route{
			strings.ToUpper("Get"),
			"/0.1/users",
			c.GetUsers,
		},
		"RegisterUser": Route{
			strings.ToUpper("Post"),
			"/0.1/users/register",
			c.RegisterUser,
		},
		"RemoveGroupUser": Route{
			strings.ToUpper("Delete"),
			"/0.1/groups/{group_uuid}/users/{user_uuid}",
			c.RemoveGroupUser,
		},
		"ResetPassword": Route{
			strings.ToUpper("Get"),
			"/0.1/users/password/reset",
			c.ResetPassword,
		},
		"ResetPasswordChange": Route{
			strings.ToUpper("Post"),
			"/0.1/users/password/reset",
			c.ResetPasswordChange,
		},
		"UpdateAllUserEmails": Route{
			strings.ToUpper("Put"),
			"/0.1/admin/users/{user_uuid}/emails",
			c.UpdateAllUserEmails,
		},
		"UpdateUser": Route{
			strings.ToUpper("Put"),
			"/0.1/users/{user_uuid}",
			c.UpdateUser,
		},
		"UpdateUserEmails": Route{
			strings.ToUpper("Put"),
			"/0.1/users/{user_uuid}/emails",
			c.UpdateUserEmails,
		},
		"UserDeleteSession": Route{
			strings.ToUpper("Delete"),
			"/0.1/users/{user_uuid}/sessions/{session_uuid}",
			c.UserDeleteSession,
		},
	}
}

// AddGroupUser - Associate a group to a user
func (c *UsersAPIController) AddGroupUser(w http.ResponseWriter, r *http.Request) {
	groupUuidParam := chi.URLParam(r, "group_uuid")
	if groupUuidParam == "" {
		c.errorHandler(w, r, &RequiredError{"group_uuid"}, nil)
		return
	}
	userUuidParam := chi.URLParam(r, "user_uuid")
	if userUuidParam == "" {
		c.errorHandler(w, r, &RequiredError{"user_uuid"}, nil)
		return
	}
	result, err := c.service.AddGroupUser(r.Context(), groupUuidParam, userUuidParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, result.Headers, w)
}

// AddUserPolicy - Associate a policy to a user
func (c *UsersAPIController) AddUserPolicy(w http.ResponseWriter, r *http.Request) {
	policyUuidParam := chi.URLParam(r, "policy_uuid")
	if policyUuidParam == "" {
		c.errorHandler(w, r, &RequiredError{"policy_uuid"}, nil)
		return
	}
	userUuidParam := chi.URLParam(r, "user_uuid")
	if userUuidParam == "" {
		c.errorHandler(w, r, &RequiredError{"user_uuid"}, nil)
		return
	}
	result, err := c.service.AddUserPolicy(r.Context(), policyUuidParam, userUuidParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, result.Headers, w)
}

// ChangeUserPassword - Change the user's password
func (c *UsersAPIController) ChangeUserPassword(w http.ResponseWriter, r *http.Request) {
	userUuidParam := chi.URLParam(r, "user_uuid")
	if userUuidParam == "" {
		c.errorHandler(w, r, &RequiredError{"user_uuid"}, nil)
		return
	}
	bodyParam := PasswordChange{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&bodyParam); err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	if err := AssertPasswordChangeRequired(bodyParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	if err := AssertPasswordChangeConstraints(bodyParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	result, err := c.service.ChangeUserPassword(r.Context(), userUuidParam, bodyParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, result.Headers, w)
}

// CreateUser - Create a user
func (c *UsersAPIController) CreateUser(w http.ResponseWriter, r *http.Request) {
	accentTenantParam := r.Header.Get("Accent-Tenant")
	bodyParam := UserCreate{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&bodyParam); err != nil && !errors.Is(err, io.EOF) {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	if err := AssertUserCreateRequired(bodyParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	if err := AssertUserCreateConstraints(bodyParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	result, err := c.service.CreateUser(r.Context(), accentTenantParam, bodyParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, result.Headers, w)
}

// DeleteRefreshTokens - Delete a user's refresh token
func (c *UsersAPIController) DeleteRefreshTokens(w http.ResponseWriter, r *http.Request) {
	userUuidOrMeParam := chi.URLParam(r, "user_uuid_or_me")
	if userUuidOrMeParam == "" {
		c.errorHandler(w, r, &RequiredError{"user_uuid_or_me"}, nil)
		return
	}
	clientIdParam := chi.URLParam(r, "client_id")
	if clientIdParam == "" {
		c.errorHandler(w, r, &RequiredError{"client_id"}, nil)
		return
	}
	result, err := c.service.DeleteRefreshTokens(r.Context(), userUuidOrMeParam, clientIdParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, result.Headers, w)
}

// DeleteUser - Delete a user
func (c *UsersAPIController) DeleteUser(w http.ResponseWriter, r *http.Request) {
	userUuidParam := chi.URLParam(r, "user_uuid")
	if userUuidParam == "" {
		c.errorHandler(w, r, &RequiredError{"user_uuid"}, nil)
		return
	}
	result, err := c.service.DeleteUser(r.Context(), userUuidParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, result.Headers, w)
}

// DeleteUserPolicy - Dissociate a policy from a user
func (c *UsersAPIController) DeleteUserPolicy(w http.ResponseWriter, r *http.Request) {
	policyUuidParam := chi.URLParam(r, "policy_uuid")
	if policyUuidParam == "" {
		c.errorHandler(w, r, &RequiredError{"policy_uuid"}, nil)
		return
	}
	userUuidParam := chi.URLParam(r, "user_uuid")
	if userUuidParam == "" {
		c.errorHandler(w, r, &RequiredError{"user_uuid"}, nil)
		return
	}
	result, err := c.service.DeleteUserPolicy(r.Context(), policyUuidParam, userUuidParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, result.Headers, w)
}

// GetNewEmailConfirmation - Ask a new confirmation email
func (c *UsersAPIController) GetNewEmailConfirmation(w http.ResponseWriter, r *http.Request) {
	userUuidParam := chi.URLParam(r, "user_uuid")
	if userUuidParam == "" {
		c.errorHandler(w, r, &RequiredError{"user_uuid"}, nil)
		return
	}
	emailUuidParam := chi.URLParam(r, "email_uuid")
	if emailUuidParam == "" {
		c.errorHandler(w, r, &RequiredError{"email_uuid"}, nil)
		return
	}
	result, err := c.service.GetNewEmailConfirmation(r.Context(), userUuidParam, emailUuidParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, result.Headers, w)
}

// GetUser - Retrieves the details of a user
func (c *UsersAPIController) GetUser(w http.ResponseWriter, r *http.Request) {
	userUuidParam := chi.URLParam(r, "user_uuid")
	if userUuidParam == "" {
		c.errorHandler(w, r, &RequiredError{"user_uuid"}, nil)
		return
	}
	result, err := c.service.GetUser(r.Context(), userUuidParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, result.Headers, w)
}

// GetUserExternalAuth - Retrieves the list of the users external auth data
func (c *UsersAPIController) GetUserExternalAuth(w http.ResponseWriter, r *http.Request) {
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
		var param int32 = 0
		offsetParam = param
	}
	var searchParam string
	if query.Has("search") {
		param := query.Get("search")

		searchParam = param
	} else {
	}
	result, err := c.service.GetUserExternalAuth(r.Context(), userUuidParam, orderParam, directionParam, limitParam, offsetParam, searchParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, result.Headers, w)
}

// GetUserGroups - Retrieves the list of groups associated to a user
func (c *UsersAPIController) GetUserGroups(w http.ResponseWriter, r *http.Request) {
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
		var param int32 = 0
		offsetParam = param
	}
	var searchParam string
	if query.Has("search") {
		param := query.Get("search")

		searchParam = param
	} else {
	}
	result, err := c.service.GetUserGroups(r.Context(), userUuidParam, orderParam, directionParam, limitParam, offsetParam, searchParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, result.Headers, w)
}

// GetUserPolicies - Retrieves the list of policies associated to a user
func (c *UsersAPIController) GetUserPolicies(w http.ResponseWriter, r *http.Request) {
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
		var param int32 = 0
		offsetParam = param
	}
	var searchParam string
	if query.Has("search") {
		param := query.Get("search")

		searchParam = param
	} else {
	}
	result, err := c.service.GetUserPolicies(r.Context(), userUuidParam, orderParam, directionParam, limitParam, offsetParam, searchParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, result.Headers, w)
}

// GetUserSessions - Retrieves the list of sessions associated to a user
func (c *UsersAPIController) GetUserSessions(w http.ResponseWriter, r *http.Request) {
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
		var param int32 = 0
		offsetParam = param
	}
	result, err := c.service.GetUserSessions(r.Context(), userUuidParam, accentTenantParam, limitParam, offsetParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, result.Headers, w)
}

// GetUserTokens - Retrieve a user's refresh token list
func (c *UsersAPIController) GetUserTokens(w http.ResponseWriter, r *http.Request) {
	query, err := parseQuery(r.URL.RawQuery)
	if err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	userUuidOrMeParam := chi.URLParam(r, "user_uuid_or_me")
	if userUuidOrMeParam == "" {
		c.errorHandler(w, r, &RequiredError{"user_uuid_or_me"}, nil)
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
		var param int32 = 0
		offsetParam = param
	}
	var searchParam string
	if query.Has("search") {
		param := query.Get("search")

		searchParam = param
	} else {
	}
	result, err := c.service.GetUserTokens(r.Context(), userUuidOrMeParam, accentTenantParam, orderParam, directionParam, limitParam, offsetParam, searchParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, result.Headers, w)
}

// GetUsers - Retrieves the list of users
func (c *UsersAPIController) GetUsers(w http.ResponseWriter, r *http.Request) {
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
		var param int32 = 0
		offsetParam = param
	}
	var searchParam string
	if query.Has("search") {
		param := query.Get("search")

		searchParam = param
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
	var hasPolicySlugParam string
	if query.Has("has_policy_slug") {
		param := query.Get("has_policy_slug")

		hasPolicySlugParam = param
	} else {
	}
	var hasPolicyUuidParam string
	if query.Has("has_policy_uuid") {
		param := query.Get("has_policy_uuid")

		hasPolicyUuidParam = param
	} else {
	}
	var policySlugParam string
	if query.Has("policy_slug") {
		param := query.Get("policy_slug")

		policySlugParam = param
	} else {
	}
	var policyUuidParam string
	if query.Has("policy_uuid") {
		param := query.Get("policy_uuid")

		policyUuidParam = param
	} else {
	}
	result, err := c.service.GetUsers(r.Context(), orderParam, directionParam, limitParam, offsetParam, searchParam, accentTenantParam, recurseParam, hasPolicySlugParam, hasPolicyUuidParam, policySlugParam, policyUuidParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, result.Headers, w)
}

// RegisterUser - Create a user
func (c *UsersAPIController) RegisterUser(w http.ResponseWriter, r *http.Request) {
	bodyParam := UserRegister{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&bodyParam); err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	if err := AssertUserRegisterRequired(bodyParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	if err := AssertUserRegisterConstraints(bodyParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	result, err := c.service.RegisterUser(r.Context(), bodyParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, result.Headers, w)
}

// RemoveGroupUser - Dissociate a user from a group
func (c *UsersAPIController) RemoveGroupUser(w http.ResponseWriter, r *http.Request) {
	groupUuidParam := chi.URLParam(r, "group_uuid")
	if groupUuidParam == "" {
		c.errorHandler(w, r, &RequiredError{"group_uuid"}, nil)
		return
	}
	userUuidParam := chi.URLParam(r, "user_uuid")
	if userUuidParam == "" {
		c.errorHandler(w, r, &RequiredError{"user_uuid"}, nil)
		return
	}
	result, err := c.service.RemoveGroupUser(r.Context(), groupUuidParam, userUuidParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, result.Headers, w)
}

// ResetPassword - Reset the user password
func (c *UsersAPIController) ResetPassword(w http.ResponseWriter, r *http.Request) {
	query, err := parseQuery(r.URL.RawQuery)
	if err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	var usernameParam string
	if query.Has("username") {
		param := query.Get("username")

		usernameParam = param
	} else {
	}
	var emailParam string
	if query.Has("email") {
		param := query.Get("email")

		emailParam = param
	} else {
	}
	var loginParam string
	if query.Has("login") {
		param := query.Get("login")

		loginParam = param
	} else {
	}
	result, err := c.service.ResetPassword(r.Context(), usernameParam, emailParam, loginParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, result.Headers, w)
}

// ResetPasswordChange - Set the user password
func (c *UsersAPIController) ResetPasswordChange(w http.ResponseWriter, r *http.Request) {
	query, err := parseQuery(r.URL.RawQuery)
	if err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	var userUuidParam string
	if query.Has("user_uuid") {
		param := query.Get("user_uuid")

		userUuidParam = param
	} else {
		c.errorHandler(w, r, &RequiredError{Field: "user_uuid"}, nil)
		return
	}
	bodyParam := PostPasswordReset{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&bodyParam); err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	if err := AssertPostPasswordResetRequired(bodyParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	if err := AssertPostPasswordResetConstraints(bodyParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	result, err := c.service.ResetPasswordChange(r.Context(), userUuidParam, bodyParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, result.Headers, w)
}

// UpdateAllUserEmails - Update email addresses
func (c *UsersAPIController) UpdateAllUserEmails(w http.ResponseWriter, r *http.Request) {
	userUuidParam := chi.URLParam(r, "user_uuid")
	if userUuidParam == "" {
		c.errorHandler(w, r, &RequiredError{"user_uuid"}, nil)
		return
	}
	bodyParam := AdminUserEmailList{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&bodyParam); err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	if err := AssertAdminUserEmailListRequired(bodyParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	if err := AssertAdminUserEmailListConstraints(bodyParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	result, err := c.service.UpdateAllUserEmails(r.Context(), userUuidParam, bodyParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, result.Headers, w)
}

// UpdateUser - Update an existing user
func (c *UsersAPIController) UpdateUser(w http.ResponseWriter, r *http.Request) {
	userUuidParam := chi.URLParam(r, "user_uuid")
	if userUuidParam == "" {
		c.errorHandler(w, r, &RequiredError{"user_uuid"}, nil)
		return
	}
	bodyParam := UserEdit{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&bodyParam); err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	if err := AssertUserEditRequired(bodyParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	if err := AssertUserEditConstraints(bodyParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	result, err := c.service.UpdateUser(r.Context(), userUuidParam, bodyParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, result.Headers, w)
}

// UpdateUserEmails - Update email addresses
func (c *UsersAPIController) UpdateUserEmails(w http.ResponseWriter, r *http.Request) {
	userUuidParam := chi.URLParam(r, "user_uuid")
	if userUuidParam == "" {
		c.errorHandler(w, r, &RequiredError{"user_uuid"}, nil)
		return
	}
	bodyParam := UserEmailList{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&bodyParam); err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	if err := AssertUserEmailListRequired(bodyParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	if err := AssertUserEmailListConstraints(bodyParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	result, err := c.service.UpdateUserEmails(r.Context(), userUuidParam, bodyParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, result.Headers, w)
}

// UserDeleteSession - Delete a session
func (c *UsersAPIController) UserDeleteSession(w http.ResponseWriter, r *http.Request) {
	userUuidParam := chi.URLParam(r, "user_uuid")
	if userUuidParam == "" {
		c.errorHandler(w, r, &RequiredError{"user_uuid"}, nil)
		return
	}
	sessionUuidParam := chi.URLParam(r, "session_uuid")
	if sessionUuidParam == "" {
		c.errorHandler(w, r, &RequiredError{"session_uuid"}, nil)
		return
	}
	result, err := c.service.UserDeleteSession(r.Context(), userUuidParam, sessionUuidParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, result.Headers, w)
}