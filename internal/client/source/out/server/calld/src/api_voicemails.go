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

// VoicemailsAPIController binds http requests to an api service and writes the service results to the http response
type VoicemailsAPIController struct {
	service      VoicemailsAPIServicer
	errorHandler ErrorHandler
}

// VoicemailsAPIOption for how the controller is set up.
type VoicemailsAPIOption func(*VoicemailsAPIController)

// WithVoicemailsAPIErrorHandler inject ErrorHandler into controller
func WithVoicemailsAPIErrorHandler(h ErrorHandler) VoicemailsAPIOption {
	return func(c *VoicemailsAPIController) {
		c.errorHandler = h
	}
}

// NewVoicemailsAPIController creates a default api controller
func NewVoicemailsAPIController(s VoicemailsAPIServicer, opts ...VoicemailsAPIOption) Router {
	controller := &VoicemailsAPIController{
		service:      s,
		errorHandler: DefaultErrorHandler,
	}

	for _, opt := range opts {
		opt(controller)
	}

	return controller
}

// Routes returns all the api routes for the VoicemailsAPIController
func (c *VoicemailsAPIController) Routes() Routes {
	return Routes{
		"CheckUserVoicemailGreeting": Route{
			strings.ToUpper("Head"),
			"/1.0/users/me/voicemails/greetings/{greeting}",
			c.CheckUserVoicemailGreeting,
		},
		"CheckVoicemailGreeting": Route{
			strings.ToUpper("Head"),
			"/1.0/voicemails/{voicemail_id}/greetings/{greeting}",
			c.CheckVoicemailGreeting,
		},
		"CopyUserVoicemailGreeting": Route{
			strings.ToUpper("Post"),
			"/1.0/users/me/voicemails/greetings/{greeting}/copy",
			c.CopyUserVoicemailGreeting,
		},
		"CopyVoicemailGreeting": Route{
			strings.ToUpper("Post"),
			"/1.0/voicemails/{voicemail_id}/greetings/{greeting}/copy",
			c.CopyVoicemailGreeting,
		},
		"CreateUserVoicemailGreeting": Route{
			strings.ToUpper("Post"),
			"/1.0/users/me/voicemails/greetings/{greeting}",
			c.CreateUserVoicemailGreeting,
		},
		"CreateVoicemailGreeting": Route{
			strings.ToUpper("Post"),
			"/1.0/voicemails/{voicemail_id}/greetings/{greeting}",
			c.CreateVoicemailGreeting,
		},
		"DeleteUserVoicemailGreeting": Route{
			strings.ToUpper("Delete"),
			"/1.0/users/me/voicemails/greetings/{greeting}",
			c.DeleteUserVoicemailGreeting,
		},
		"DeleteUserVoicemailMessage": Route{
			strings.ToUpper("Delete"),
			"/1.0/users/me/voicemails/messages/{message_id}",
			c.DeleteUserVoicemailMessage,
		},
		"DeleteVoicemailGreeting": Route{
			strings.ToUpper("Delete"),
			"/1.0/voicemails/{voicemail_id}/greetings/{greeting}",
			c.DeleteVoicemailGreeting,
		},
		"DeleteVoicemailMessage": Route{
			strings.ToUpper("Delete"),
			"/1.0/voicemails/{voicemail_id}/messages/{message_id}",
			c.DeleteVoicemailMessage,
		},
		"GetUserVoicemailFolder": Route{
			strings.ToUpper("Get"),
			"/1.0/users/me/voicemails/folders/{folder_id}",
			c.GetUserVoicemailFolder,
		},
		"GetUserVoicemailGreeting": Route{
			strings.ToUpper("Get"),
			"/1.0/users/me/voicemails/greetings/{greeting}",
			c.GetUserVoicemailGreeting,
		},
		"GetUserVoicemailMessage": Route{
			strings.ToUpper("Get"),
			"/1.0/users/me/voicemails/messages/{message_id}",
			c.GetUserVoicemailMessage,
		},
		"GetUserVoicemailMessageRecording": Route{
			strings.ToUpper("Get"),
			"/1.0/users/me/voicemails/messages/{message_id}/recording",
			c.GetUserVoicemailMessageRecording,
		},
		"GetVoicemail": Route{
			strings.ToUpper("Get"),
			"/1.0/voicemails/{voicemail_id}",
			c.GetVoicemail,
		},
		"GetVoicemailFolder": Route{
			strings.ToUpper("Get"),
			"/1.0/voicemails/{voicemail_id}/folders/{folder_id}",
			c.GetVoicemailFolder,
		},
		"GetVoicemailGreeting": Route{
			strings.ToUpper("Get"),
			"/1.0/voicemails/{voicemail_id}/greetings/{greeting}",
			c.GetVoicemailGreeting,
		},
		"GetVoicemailMessage": Route{
			strings.ToUpper("Get"),
			"/1.0/voicemails/{voicemail_id}/messages/{message_id}",
			c.GetVoicemailMessage,
		},
		"GetVoicemailMessageRecording": Route{
			strings.ToUpper("Get"),
			"/1.0/voicemails/{voicemail_id}/messages/{message_id}/recording",
			c.GetVoicemailMessageRecording,
		},
		"ListUserVoicemails": Route{
			strings.ToUpper("Get"),
			"/1.0/users/me/voicemails",
			c.ListUserVoicemails,
		},
		"UpdateUserVoicemailGreeting": Route{
			strings.ToUpper("Put"),
			"/1.0/users/me/voicemails/greetings/{greeting}",
			c.UpdateUserVoicemailGreeting,
		},
		"UpdateUserVoicemailMessage": Route{
			strings.ToUpper("Put"),
			"/1.0/users/me/voicemails/messages/{message_id}",
			c.UpdateUserVoicemailMessage,
		},
		"UpdateVoicemailGreeting": Route{
			strings.ToUpper("Put"),
			"/1.0/voicemails/{voicemail_id}/greetings/{greeting}",
			c.UpdateVoicemailGreeting,
		},
		"UpdateVoicemailMessage": Route{
			strings.ToUpper("Put"),
			"/1.0/voicemails/{voicemail_id}/messages/{message_id}",
			c.UpdateVoicemailMessage,
		},
	}
}

// CheckUserVoicemailGreeting - Check if greeting exists
func (c *VoicemailsAPIController) CheckUserVoicemailGreeting(w http.ResponseWriter, r *http.Request) {
	greetingParam := chi.URLParam(r, "greeting")
	if greetingParam == "" {
		c.errorHandler(w, r, &RequiredError{"greeting"}, nil)
		return
	}
	result, err := c.service.CheckUserVoicemailGreeting(r.Context(), greetingParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, result.Headers, w)
}

// CheckVoicemailGreeting - Check if greeting exists
func (c *VoicemailsAPIController) CheckVoicemailGreeting(w http.ResponseWriter, r *http.Request) {
	voicemailIdParam, err := parseNumericParameter[int32](
		chi.URLParam(r, "voicemail_id"),
		WithRequire[int32](parseInt32),
	)
	if err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	greetingParam := chi.URLParam(r, "greeting")
	if greetingParam == "" {
		c.errorHandler(w, r, &RequiredError{"greeting"}, nil)
		return
	}
	result, err := c.service.CheckVoicemailGreeting(r.Context(), voicemailIdParam, greetingParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, result.Headers, w)
}

// CopyUserVoicemailGreeting - Copy a custom greeting
func (c *VoicemailsAPIController) CopyUserVoicemailGreeting(w http.ResponseWriter, r *http.Request) {
	greetingParam := chi.URLParam(r, "greeting")
	if greetingParam == "" {
		c.errorHandler(w, r, &RequiredError{"greeting"}, nil)
		return
	}
	bodyParam := GreetingCopy{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&bodyParam); err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	if err := AssertGreetingCopyRequired(bodyParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	if err := AssertGreetingCopyConstraints(bodyParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	result, err := c.service.CopyUserVoicemailGreeting(r.Context(), greetingParam, bodyParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, result.Headers, w)
}

// CopyVoicemailGreeting - Copy a custom greeting
func (c *VoicemailsAPIController) CopyVoicemailGreeting(w http.ResponseWriter, r *http.Request) {
	voicemailIdParam, err := parseNumericParameter[int32](
		chi.URLParam(r, "voicemail_id"),
		WithRequire[int32](parseInt32),
	)
	if err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	greetingParam := chi.URLParam(r, "greeting")
	if greetingParam == "" {
		c.errorHandler(w, r, &RequiredError{"greeting"}, nil)
		return
	}
	bodyParam := GreetingCopy{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&bodyParam); err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	if err := AssertGreetingCopyRequired(bodyParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	if err := AssertGreetingCopyConstraints(bodyParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	result, err := c.service.CopyVoicemailGreeting(r.Context(), voicemailIdParam, greetingParam, bodyParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, result.Headers, w)
}

// CreateUserVoicemailGreeting - Create a custom greeting
func (c *VoicemailsAPIController) CreateUserVoicemailGreeting(w http.ResponseWriter, r *http.Request) {
	greetingParam := chi.URLParam(r, "greeting")
	if greetingParam == "" {
		c.errorHandler(w, r, &RequiredError{"greeting"}, nil)
		return
	}
	bodyParam := map[string]interface{}{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&bodyParam); err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	result, err := c.service.CreateUserVoicemailGreeting(r.Context(), greetingParam, bodyParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, result.Headers, w)
}

// CreateVoicemailGreeting - Create a custom greeting
func (c *VoicemailsAPIController) CreateVoicemailGreeting(w http.ResponseWriter, r *http.Request) {
	voicemailIdParam, err := parseNumericParameter[int32](
		chi.URLParam(r, "voicemail_id"),
		WithRequire[int32](parseInt32),
	)
	if err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	greetingParam := chi.URLParam(r, "greeting")
	if greetingParam == "" {
		c.errorHandler(w, r, &RequiredError{"greeting"}, nil)
		return
	}
	bodyParam := map[string]interface{}{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&bodyParam); err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	result, err := c.service.CreateVoicemailGreeting(r.Context(), voicemailIdParam, greetingParam, bodyParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, result.Headers, w)
}

// DeleteUserVoicemailGreeting - Delete a custom greeting
func (c *VoicemailsAPIController) DeleteUserVoicemailGreeting(w http.ResponseWriter, r *http.Request) {
	greetingParam := chi.URLParam(r, "greeting")
	if greetingParam == "" {
		c.errorHandler(w, r, &RequiredError{"greeting"}, nil)
		return
	}
	result, err := c.service.DeleteUserVoicemailGreeting(r.Context(), greetingParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, result.Headers, w)
}

// DeleteUserVoicemailMessage - Delete a mesage
func (c *VoicemailsAPIController) DeleteUserVoicemailMessage(w http.ResponseWriter, r *http.Request) {
	messageIdParam := chi.URLParam(r, "message_id")
	if messageIdParam == "" {
		c.errorHandler(w, r, &RequiredError{"message_id"}, nil)
		return
	}
	result, err := c.service.DeleteUserVoicemailMessage(r.Context(), messageIdParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, result.Headers, w)
}

// DeleteVoicemailGreeting - Delete a custom greeting
func (c *VoicemailsAPIController) DeleteVoicemailGreeting(w http.ResponseWriter, r *http.Request) {
	voicemailIdParam, err := parseNumericParameter[int32](
		chi.URLParam(r, "voicemail_id"),
		WithRequire[int32](parseInt32),
	)
	if err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	greetingParam := chi.URLParam(r, "greeting")
	if greetingParam == "" {
		c.errorHandler(w, r, &RequiredError{"greeting"}, nil)
		return
	}
	result, err := c.service.DeleteVoicemailGreeting(r.Context(), voicemailIdParam, greetingParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, result.Headers, w)
}

// DeleteVoicemailMessage - Delete a mesage
func (c *VoicemailsAPIController) DeleteVoicemailMessage(w http.ResponseWriter, r *http.Request) {
	voicemailIdParam, err := parseNumericParameter[int32](
		chi.URLParam(r, "voicemail_id"),
		WithRequire[int32](parseInt32),
	)
	if err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	messageIdParam := chi.URLParam(r, "message_id")
	if messageIdParam == "" {
		c.errorHandler(w, r, &RequiredError{"message_id"}, nil)
		return
	}
	result, err := c.service.DeleteVoicemailMessage(r.Context(), voicemailIdParam, messageIdParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, result.Headers, w)
}

// GetUserVoicemailFolder - Get details of a folder
func (c *VoicemailsAPIController) GetUserVoicemailFolder(w http.ResponseWriter, r *http.Request) {
	folderIdParam, err := parseNumericParameter[int32](
		chi.URLParam(r, "folder_id"),
		WithRequire[int32](parseInt32),
	)
	if err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	result, err := c.service.GetUserVoicemailFolder(r.Context(), folderIdParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, result.Headers, w)
}

// GetUserVoicemailGreeting - Get a custom greeting
func (c *VoicemailsAPIController) GetUserVoicemailGreeting(w http.ResponseWriter, r *http.Request) {
	greetingParam := chi.URLParam(r, "greeting")
	if greetingParam == "" {
		c.errorHandler(w, r, &RequiredError{"greeting"}, nil)
		return
	}
	result, err := c.service.GetUserVoicemailGreeting(r.Context(), greetingParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, result.Headers, w)
}

// GetUserVoicemailMessage - Get a message
func (c *VoicemailsAPIController) GetUserVoicemailMessage(w http.ResponseWriter, r *http.Request) {
	messageIdParam := chi.URLParam(r, "message_id")
	if messageIdParam == "" {
		c.errorHandler(w, r, &RequiredError{"message_id"}, nil)
		return
	}
	result, err := c.service.GetUserVoicemailMessage(r.Context(), messageIdParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, result.Headers, w)
}

// GetUserVoicemailMessageRecording - Get a message's recording
func (c *VoicemailsAPIController) GetUserVoicemailMessageRecording(w http.ResponseWriter, r *http.Request) {
	query, err := parseQuery(r.URL.RawQuery)
	if err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	messageIdParam := chi.URLParam(r, "message_id")
	if messageIdParam == "" {
		c.errorHandler(w, r, &RequiredError{"message_id"}, nil)
		return
	}
	var tokenParam string
	if query.Has("token") {
		param := query.Get("token")

		tokenParam = param
	} else {
	}
	var downloadParam string
	if query.Has("download") {
		param := query.Get("download")

		downloadParam = param
	} else {
	}
	result, err := c.service.GetUserVoicemailMessageRecording(r.Context(), messageIdParam, tokenParam, downloadParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, result.Headers, w)
}

// GetVoicemail - Get details of a voicemail
func (c *VoicemailsAPIController) GetVoicemail(w http.ResponseWriter, r *http.Request) {
	voicemailIdParam, err := parseNumericParameter[int32](
		chi.URLParam(r, "voicemail_id"),
		WithRequire[int32](parseInt32),
	)
	if err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	result, err := c.service.GetVoicemail(r.Context(), voicemailIdParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, result.Headers, w)
}

// GetVoicemailFolder - Get details of a folder
func (c *VoicemailsAPIController) GetVoicemailFolder(w http.ResponseWriter, r *http.Request) {
	voicemailIdParam, err := parseNumericParameter[int32](
		chi.URLParam(r, "voicemail_id"),
		WithRequire[int32](parseInt32),
	)
	if err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	folderIdParam, err := parseNumericParameter[int32](
		chi.URLParam(r, "folder_id"),
		WithRequire[int32](parseInt32),
	)
	if err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	result, err := c.service.GetVoicemailFolder(r.Context(), voicemailIdParam, folderIdParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, result.Headers, w)
}

// GetVoicemailGreeting - Get a custom greeting
func (c *VoicemailsAPIController) GetVoicemailGreeting(w http.ResponseWriter, r *http.Request) {
	voicemailIdParam, err := parseNumericParameter[int32](
		chi.URLParam(r, "voicemail_id"),
		WithRequire[int32](parseInt32),
	)
	if err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	greetingParam := chi.URLParam(r, "greeting")
	if greetingParam == "" {
		c.errorHandler(w, r, &RequiredError{"greeting"}, nil)
		return
	}
	result, err := c.service.GetVoicemailGreeting(r.Context(), voicemailIdParam, greetingParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, result.Headers, w)
}

// GetVoicemailMessage - Get a message
func (c *VoicemailsAPIController) GetVoicemailMessage(w http.ResponseWriter, r *http.Request) {
	voicemailIdParam, err := parseNumericParameter[int32](
		chi.URLParam(r, "voicemail_id"),
		WithRequire[int32](parseInt32),
	)
	if err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	messageIdParam := chi.URLParam(r, "message_id")
	if messageIdParam == "" {
		c.errorHandler(w, r, &RequiredError{"message_id"}, nil)
		return
	}
	result, err := c.service.GetVoicemailMessage(r.Context(), voicemailIdParam, messageIdParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, result.Headers, w)
}

// GetVoicemailMessageRecording - Get a message's recording
func (c *VoicemailsAPIController) GetVoicemailMessageRecording(w http.ResponseWriter, r *http.Request) {
	query, err := parseQuery(r.URL.RawQuery)
	if err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	voicemailIdParam, err := parseNumericParameter[int32](
		chi.URLParam(r, "voicemail_id"),
		WithRequire[int32](parseInt32),
	)
	if err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	messageIdParam := chi.URLParam(r, "message_id")
	if messageIdParam == "" {
		c.errorHandler(w, r, &RequiredError{"message_id"}, nil)
		return
	}
	var tokenParam string
	if query.Has("token") {
		param := query.Get("token")

		tokenParam = param
	} else {
	}
	var downloadParam string
	if query.Has("download") {
		param := query.Get("download")

		downloadParam = param
	} else {
	}
	result, err := c.service.GetVoicemailMessageRecording(r.Context(), voicemailIdParam, messageIdParam, tokenParam, downloadParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, result.Headers, w)
}

// ListUserVoicemails - Get details of the voicemail of the authenticated user
func (c *VoicemailsAPIController) ListUserVoicemails(w http.ResponseWriter, r *http.Request) {
	result, err := c.service.ListUserVoicemails(r.Context())
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, result.Headers, w)
}

// UpdateUserVoicemailGreeting - Update a custom greeting
func (c *VoicemailsAPIController) UpdateUserVoicemailGreeting(w http.ResponseWriter, r *http.Request) {
	greetingParam := chi.URLParam(r, "greeting")
	if greetingParam == "" {
		c.errorHandler(w, r, &RequiredError{"greeting"}, nil)
		return
	}
	bodyParam := map[string]interface{}{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&bodyParam); err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	result, err := c.service.UpdateUserVoicemailGreeting(r.Context(), greetingParam, bodyParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, result.Headers, w)
}

// UpdateUserVoicemailMessage - Update a message
func (c *VoicemailsAPIController) UpdateUserVoicemailMessage(w http.ResponseWriter, r *http.Request) {
	messageIdParam := chi.URLParam(r, "message_id")
	if messageIdParam == "" {
		c.errorHandler(w, r, &RequiredError{"message_id"}, nil)
		return
	}
	bodyParam := VoicemailMessageUpdate{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&bodyParam); err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	if err := AssertVoicemailMessageUpdateRequired(bodyParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	if err := AssertVoicemailMessageUpdateConstraints(bodyParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	result, err := c.service.UpdateUserVoicemailMessage(r.Context(), messageIdParam, bodyParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, result.Headers, w)
}

// UpdateVoicemailGreeting - Update a custom greeting
func (c *VoicemailsAPIController) UpdateVoicemailGreeting(w http.ResponseWriter, r *http.Request) {
	voicemailIdParam, err := parseNumericParameter[int32](
		chi.URLParam(r, "voicemail_id"),
		WithRequire[int32](parseInt32),
	)
	if err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	greetingParam := chi.URLParam(r, "greeting")
	if greetingParam == "" {
		c.errorHandler(w, r, &RequiredError{"greeting"}, nil)
		return
	}
	bodyParam := map[string]interface{}{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&bodyParam); err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	result, err := c.service.UpdateVoicemailGreeting(r.Context(), voicemailIdParam, greetingParam, bodyParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, result.Headers, w)
}

// UpdateVoicemailMessage - Update a message
func (c *VoicemailsAPIController) UpdateVoicemailMessage(w http.ResponseWriter, r *http.Request) {
	voicemailIdParam, err := parseNumericParameter[int32](
		chi.URLParam(r, "voicemail_id"),
		WithRequire[int32](parseInt32),
	)
	if err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	messageIdParam := chi.URLParam(r, "message_id")
	if messageIdParam == "" {
		c.errorHandler(w, r, &RequiredError{"message_id"}, nil)
		return
	}
	bodyParam := VoicemailMessageUpdate{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&bodyParam); err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	if err := AssertVoicemailMessageUpdateRequired(bodyParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	if err := AssertVoicemailMessageUpdateConstraints(bodyParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	result, err := c.service.UpdateVoicemailMessage(r.Context(), voicemailIdParam, messageIdParam, bodyParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, result.Headers, w)
}