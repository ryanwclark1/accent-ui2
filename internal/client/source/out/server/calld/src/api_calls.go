/*
 * accent-calld
 *
 * Control your calls from a REST API
 *
 * API version: 1.0.0
 * Contact: help@accentvoice.io
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package calldserver

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5"
)

// CallsAPIController binds http requests to an api service and writes the service results to the http response
type CallsAPIController struct {
	service      CallsAPIServicer
	errorHandler ErrorHandler
}

// CallsAPIOption for how the controller is set up.
type CallsAPIOption func(*CallsAPIController)

// WithCallsAPIErrorHandler inject ErrorHandler into controller
func WithCallsAPIErrorHandler(h ErrorHandler) CallsAPIOption {
	return func(c *CallsAPIController) {
		c.errorHandler = h
	}
}

// NewCallsAPIController creates a default api controller
func NewCallsAPIController(s CallsAPIServicer, opts ...CallsAPIOption) Router {
	controller := &CallsAPIController{
		service:      s,
		errorHandler: DefaultErrorHandler,
	}

	for _, opt := range opts {
		opt(controller)
	}

	return controller
}

// Routes returns all the api routes for the CallsAPIController
func (c *CallsAPIController) Routes() Routes {
	return Routes{
		"AnswerCall": Route{
			strings.ToUpper("Put"),
			"/1.0/calls/{call_id}/answer",
			c.AnswerCall,
		},
		"AnswerUserCall": Route{
			strings.ToUpper("Put"),
			"/1.0/users/me/calls/{call_id}/answer",
			c.AnswerUserCall,
		},
		"ConnectCallToUser": Route{
			strings.ToUpper("Put"),
			"/1.0/calls/{call_id}/user/{user_uuid}",
			c.ConnectCallToUser,
		},
		"CreateCall": Route{
			strings.ToUpper("Post"),
			"/1.0/calls",
			c.CreateCall,
		},
		"CreateUserCall": Route{
			strings.ToUpper("Post"),
			"/1.0/users/me/calls",
			c.CreateUserCall,
		},
		"DeleteCall": Route{
			strings.ToUpper("Delete"),
			"/1.0/calls/{call_id}",
			c.DeleteCall,
		},
		"GetCall": Route{
			strings.ToUpper("Get"),
			"/1.0/calls/{call_id}",
			c.GetCall,
		},
		"HangupUserCall": Route{
			strings.ToUpper("Delete"),
			"/1.0/users/me/calls/{call_id}",
			c.HangupUserCall,
		},
		"HoldCall": Route{
			strings.ToUpper("Put"),
			"/1.0/calls/{call_id}/hold/start",
			c.HoldCall,
		},
		"HoldUserCall": Route{
			strings.ToUpper("Put"),
			"/1.0/users/me/calls/{call_id}/hold/start",
			c.HoldUserCall,
		},
		"ListCalls": Route{
			strings.ToUpper("Get"),
			"/1.0/calls",
			c.ListCalls,
		},
		"ListUserCalls": Route{
			strings.ToUpper("Get"),
			"/1.0/users/me/calls",
			c.ListUserCalls,
		},
		"MuteCall": Route{
			strings.ToUpper("Put"),
			"/1.0/calls/{call_id}/mute/start",
			c.MuteCall,
		},
		"MuteUserCall": Route{
			strings.ToUpper("Put"),
			"/1.0/users/me/calls/{call_id}/mute/start",
			c.MuteUserCall,
		},
		"SendCallDTMF": Route{
			strings.ToUpper("Put"),
			"/1.0/calls/{call_id}/dtmf",
			c.SendCallDTMF,
		},
		"SendUserDTMF": Route{
			strings.ToUpper("Put"),
			"/1.0/users/me/calls/{call_id}/dtmf",
			c.SendUserDTMF,
		},
		"StartCurrentUserRecording": Route{
			strings.ToUpper("Put"),
			"/1.0/users/me/calls/{call_id}/record/start",
			c.StartCurrentUserRecording,
		},
		"StartRecording": Route{
			strings.ToUpper("Put"),
			"/1.0/calls/{call_id}/record/start",
			c.StartRecording,
		},
		"StopCurrentUserRecording": Route{
			strings.ToUpper("Put"),
			"/1.0/users/me/calls/{call_id}/record/stop",
			c.StopCurrentUserRecording,
		},
		"StopRecording": Route{
			strings.ToUpper("Put"),
			"/1.0/calls/{call_id}/record/stop",
			c.StopRecording,
		},
		"UnholdCall": Route{
			strings.ToUpper("Put"),
			"/1.0/calls/{call_id}/hold/stop",
			c.UnholdCall,
		},
		"UnholdUserCall": Route{
			strings.ToUpper("Put"),
			"/1.0/users/me/calls/{call_id}/hold/stop",
			c.UnholdUserCall,
		},
		"UnmuteCall": Route{
			strings.ToUpper("Put"),
			"/1.0/calls/{call_id}/mute/stop",
			c.UnmuteCall,
		},
		"UnmuteUserCall": Route{
			strings.ToUpper("Put"),
			"/1.0/users/me/calls/{call_id}/mute/stop",
			c.UnmuteUserCall,
		},
	}
}

// AnswerCall - Answer a call
func (c *CallsAPIController) AnswerCall(w http.ResponseWriter, r *http.Request) {
	callIdParam := chi.URLParam(r, "call_id")
	if callIdParam == "" {
		c.errorHandler(w, r, &RequiredError{"call_id"}, nil)
		return
	}
	result, err := c.service.AnswerCall(r.Context(), callIdParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}

// AnswerUserCall - Answer a call from user
func (c *CallsAPIController) AnswerUserCall(w http.ResponseWriter, r *http.Request) {
	callIdParam := chi.URLParam(r, "call_id")
	if callIdParam == "" {
		c.errorHandler(w, r, &RequiredError{"call_id"}, nil)
		return
	}
	result, err := c.service.AnswerUserCall(r.Context(), callIdParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}

// ConnectCallToUser - Connect a call to a user
func (c *CallsAPIController) ConnectCallToUser(w http.ResponseWriter, r *http.Request) {
	callIdParam := chi.URLParam(r, "call_id")
	if callIdParam == "" {
		c.errorHandler(w, r, &RequiredError{"call_id"}, nil)
		return
	}
	userUuidParam := chi.URLParam(r, "user_uuid")
	if userUuidParam == "" {
		c.errorHandler(w, r, &RequiredError{"user_uuid"}, nil)
		return
	}
	bodyParam := ConnectCallToUserRequest{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&bodyParam); err != nil && !errors.Is(err, io.EOF) {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	if err := AssertConnectCallToUserRequestRequired(bodyParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	if err := AssertConnectCallToUserRequestConstraints(bodyParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	result, err := c.service.ConnectCallToUser(r.Context(), callIdParam, userUuidParam, bodyParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}

// CreateCall - Make a new call
func (c *CallsAPIController) CreateCall(w http.ResponseWriter, r *http.Request) {
	bodyParam := CallRequest{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&bodyParam); err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	if err := AssertCallRequestRequired(bodyParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	if err := AssertCallRequestConstraints(bodyParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	result, err := c.service.CreateCall(r.Context(), bodyParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}

// CreateUserCall - Make a new call from a user
func (c *CallsAPIController) CreateUserCall(w http.ResponseWriter, r *http.Request) {
	bodyParam := UserCallRequest{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&bodyParam); err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	if err := AssertUserCallRequestRequired(bodyParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	if err := AssertUserCallRequestConstraints(bodyParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	result, err := c.service.CreateUserCall(r.Context(), bodyParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}

// DeleteCall - Hangup a call
func (c *CallsAPIController) DeleteCall(w http.ResponseWriter, r *http.Request) {
	callIdParam := chi.URLParam(r, "call_id")
	if callIdParam == "" {
		c.errorHandler(w, r, &RequiredError{"call_id"}, nil)
		return
	}
	result, err := c.service.DeleteCall(r.Context(), callIdParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}

// GetCall - Show a call
func (c *CallsAPIController) GetCall(w http.ResponseWriter, r *http.Request) {
	callIdParam := chi.URLParam(r, "call_id")
	if callIdParam == "" {
		c.errorHandler(w, r, &RequiredError{"call_id"}, nil)
		return
	}
	result, err := c.service.GetCall(r.Context(), callIdParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}

// HangupUserCall - Hangup a call from a user
func (c *CallsAPIController) HangupUserCall(w http.ResponseWriter, r *http.Request) {
	callIdParam := chi.URLParam(r, "call_id")
	if callIdParam == "" {
		c.errorHandler(w, r, &RequiredError{"call_id"}, nil)
		return
	}
	result, err := c.service.HangupUserCall(r.Context(), callIdParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}

// HoldCall - Hold a call
func (c *CallsAPIController) HoldCall(w http.ResponseWriter, r *http.Request) {
	callIdParam := chi.URLParam(r, "call_id")
	if callIdParam == "" {
		c.errorHandler(w, r, &RequiredError{"call_id"}, nil)
		return
	}
	result, err := c.service.HoldCall(r.Context(), callIdParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}

// HoldUserCall - Hold a call from user
func (c *CallsAPIController) HoldUserCall(w http.ResponseWriter, r *http.Request) {
	callIdParam := chi.URLParam(r, "call_id")
	if callIdParam == "" {
		c.errorHandler(w, r, &RequiredError{"call_id"}, nil)
		return
	}
	result, err := c.service.HoldUserCall(r.Context(), callIdParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}

// ListCalls - List calls
func (c *CallsAPIController) ListCalls(w http.ResponseWriter, r *http.Request) {
	query, err := parseQuery(r.URL.RawQuery)
	if err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	var applicationParam string
	if query.Has("application") {
		param := query.Get("application")

		applicationParam = param
	} else {
	}
	var applicationInstanceParam string
	if query.Has("application_instance") {
		param := query.Get("application_instance")

		applicationInstanceParam = param
	} else {
	}
	result, err := c.service.ListCalls(r.Context(), applicationParam, applicationInstanceParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}

// ListUserCalls - List calls of a user
func (c *CallsAPIController) ListUserCalls(w http.ResponseWriter, r *http.Request) {
	query, err := parseQuery(r.URL.RawQuery)
	if err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	var applicationParam string
	if query.Has("application") {
		param := query.Get("application")

		applicationParam = param
	} else {
	}
	var applicationInstanceParam string
	if query.Has("application_instance") {
		param := query.Get("application_instance")

		applicationInstanceParam = param
	} else {
	}
	result, err := c.service.ListUserCalls(r.Context(), applicationParam, applicationInstanceParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}

// MuteCall - Mute a call
func (c *CallsAPIController) MuteCall(w http.ResponseWriter, r *http.Request) {
	callIdParam := chi.URLParam(r, "call_id")
	if callIdParam == "" {
		c.errorHandler(w, r, &RequiredError{"call_id"}, nil)
		return
	}
	result, err := c.service.MuteCall(r.Context(), callIdParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}

// MuteUserCall - Mute a call from user
func (c *CallsAPIController) MuteUserCall(w http.ResponseWriter, r *http.Request) {
	callIdParam := chi.URLParam(r, "call_id")
	if callIdParam == "" {
		c.errorHandler(w, r, &RequiredError{"call_id"}, nil)
		return
	}
	result, err := c.service.MuteUserCall(r.Context(), callIdParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}

// SendCallDTMF - Simulate a user pressing DTMF keys
func (c *CallsAPIController) SendCallDTMF(w http.ResponseWriter, r *http.Request) {
	query, err := parseQuery(r.URL.RawQuery)
	if err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	callIdParam := chi.URLParam(r, "call_id")
	if callIdParam == "" {
		c.errorHandler(w, r, &RequiredError{"call_id"}, nil)
		return
	}
	var digitsParam string
	if query.Has("digits") {
		param := query.Get("digits")

		digitsParam = param
	} else {
		c.errorHandler(w, r, &RequiredError{Field: "digits"}, nil)
		return
	}
	result, err := c.service.SendCallDTMF(r.Context(), callIdParam, digitsParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}

// SendUserDTMF - Simulate a user pressing DTMF keys
func (c *CallsAPIController) SendUserDTMF(w http.ResponseWriter, r *http.Request) {
	query, err := parseQuery(r.URL.RawQuery)
	if err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	callIdParam := chi.URLParam(r, "call_id")
	if callIdParam == "" {
		c.errorHandler(w, r, &RequiredError{"call_id"}, nil)
		return
	}
	var digitsParam string
	if query.Has("digits") {
		param := query.Get("digits")

		digitsParam = param
	} else {
		c.errorHandler(w, r, &RequiredError{Field: "digits"}, nil)
		return
	}
	result, err := c.service.SendUserDTMF(r.Context(), callIdParam, digitsParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}

// StartCurrentUserRecording - Start recording a call
func (c *CallsAPIController) StartCurrentUserRecording(w http.ResponseWriter, r *http.Request) {
	callIdParam := chi.URLParam(r, "call_id")
	if callIdParam == "" {
		c.errorHandler(w, r, &RequiredError{"call_id"}, nil)
		return
	}
	result, err := c.service.StartCurrentUserRecording(r.Context(), callIdParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}

// StartRecording - Start recording a call
func (c *CallsAPIController) StartRecording(w http.ResponseWriter, r *http.Request) {
	callIdParam := chi.URLParam(r, "call_id")
	if callIdParam == "" {
		c.errorHandler(w, r, &RequiredError{"call_id"}, nil)
		return
	}
	result, err := c.service.StartRecording(r.Context(), callIdParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}

// StopCurrentUserRecording - Stop recording a call
func (c *CallsAPIController) StopCurrentUserRecording(w http.ResponseWriter, r *http.Request) {
	callIdParam := chi.URLParam(r, "call_id")
	if callIdParam == "" {
		c.errorHandler(w, r, &RequiredError{"call_id"}, nil)
		return
	}
	result, err := c.service.StopCurrentUserRecording(r.Context(), callIdParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}

// StopRecording - Stop recording a call
func (c *CallsAPIController) StopRecording(w http.ResponseWriter, r *http.Request) {
	callIdParam := chi.URLParam(r, "call_id")
	if callIdParam == "" {
		c.errorHandler(w, r, &RequiredError{"call_id"}, nil)
		return
	}
	result, err := c.service.StopRecording(r.Context(), callIdParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}

// UnholdCall - Unhold a call
func (c *CallsAPIController) UnholdCall(w http.ResponseWriter, r *http.Request) {
	callIdParam := chi.URLParam(r, "call_id")
	if callIdParam == "" {
		c.errorHandler(w, r, &RequiredError{"call_id"}, nil)
		return
	}
	result, err := c.service.UnholdCall(r.Context(), callIdParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}

// UnholdUserCall - Unhold a call from user
func (c *CallsAPIController) UnholdUserCall(w http.ResponseWriter, r *http.Request) {
	callIdParam := chi.URLParam(r, "call_id")
	if callIdParam == "" {
		c.errorHandler(w, r, &RequiredError{"call_id"}, nil)
		return
	}
	result, err := c.service.UnholdUserCall(r.Context(), callIdParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}

// UnmuteCall - Unmute a call
func (c *CallsAPIController) UnmuteCall(w http.ResponseWriter, r *http.Request) {
	callIdParam := chi.URLParam(r, "call_id")
	if callIdParam == "" {
		c.errorHandler(w, r, &RequiredError{"call_id"}, nil)
		return
	}
	result, err := c.service.UnmuteCall(r.Context(), callIdParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}

// UnmuteUserCall - Unmute a call from user
func (c *CallsAPIController) UnmuteUserCall(w http.ResponseWriter, r *http.Request) {
	callIdParam := chi.URLParam(r, "call_id")
	if callIdParam == "" {
		c.errorHandler(w, r, &RequiredError{"call_id"}, nil)
		return
	}
	result, err := c.service.UnmuteUserCall(r.Context(), callIdParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}
