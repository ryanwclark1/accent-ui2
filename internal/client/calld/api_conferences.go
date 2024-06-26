/*
accent-calld

Control your calls from a REST API

API version: 1.0.0
Contact: help@accentvoice.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package calld

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)

type ConferencesAPI interface {

	/*
		KickParticipant Kick participant from a conference

		**Required ACL:** `calld.conferences.{conference_id}.participants.{participant_id}.delete`

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param conferenceId Unique identifier of the conference
		@param participantId Unique identifier of the participant
		@return ConferencesAPIKickParticipantRequest
	*/
	KickParticipant(ctx context.Context, conferenceId string, participantId string) ConferencesAPIKickParticipantRequest

	// KickParticipantExecute executes the request
	KickParticipantExecute(r ConferencesAPIKickParticipantRequest) (*http.Response, error)

	/*
		ListConferenceParticipants List participants of a conference

		**Required ACL:** `calld.conferences.{conference_id}.participants.read`

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param conferenceId Unique identifier of the conference
		@return ConferencesAPIListConferenceParticipantsRequest
	*/
	ListConferenceParticipants(ctx context.Context, conferenceId string) ConferencesAPIListConferenceParticipantsRequest

	// ListConferenceParticipantsExecute executes the request
	//  @return ParticipantList
	ListConferenceParticipantsExecute(r ConferencesAPIListConferenceParticipantsRequest) (*ParticipantList, *http.Response, error)

	/*
		ListUserConferenceParticipants List participants of a conference as a user

		**Required ACL:** `calld.users.me.conferences.{conference_id}.participants.read`

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param conferenceId Unique identifier of the conference
		@return ConferencesAPIListUserConferenceParticipantsRequest
	*/
	ListUserConferenceParticipants(ctx context.Context, conferenceId string) ConferencesAPIListUserConferenceParticipantsRequest

	// ListUserConferenceParticipantsExecute executes the request
	//  @return ParticipantList
	ListUserConferenceParticipantsExecute(r ConferencesAPIListUserConferenceParticipantsRequest) (*ParticipantList, *http.Response, error)

	/*
		MuteParticipant Mute a participant in a conference

		**Required ACL:** `calld.conferences.{conference_id}.participants.{participant_id}.mute.update`

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param conferenceId Unique identifier of the conference
		@param participantId Unique identifier of the participant
		@return ConferencesAPIMuteParticipantRequest
	*/
	MuteParticipant(ctx context.Context, conferenceId string, participantId string) ConferencesAPIMuteParticipantRequest

	// MuteParticipantExecute executes the request
	MuteParticipantExecute(r ConferencesAPIMuteParticipantRequest) (*http.Response, error)

	/*
		StartConferenceRecording Record a conference

		**Required ACL:** `calld.conferences.{conference_id}.record.create`

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param conferenceId Unique identifier of the conference
		@return ConferencesAPIStartConferenceRecordingRequest
	*/
	StartConferenceRecording(ctx context.Context, conferenceId string) ConferencesAPIStartConferenceRecordingRequest

	// StartConferenceRecordingExecute executes the request
	StartConferenceRecordingExecute(r ConferencesAPIStartConferenceRecordingRequest) (*http.Response, error)

	/*
		StopConferenceRecording Stop recording a conference

		**Required ACL:** `calld.conferences.{conference_id}.record.delete`

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param conferenceId Unique identifier of the conference
		@return ConferencesAPIStopConferenceRecordingRequest
	*/
	StopConferenceRecording(ctx context.Context, conferenceId string) ConferencesAPIStopConferenceRecordingRequest

	// StopConferenceRecordingExecute executes the request
	StopConferenceRecordingExecute(r ConferencesAPIStopConferenceRecordingRequest) (*http.Response, error)

	/*
		UnmuteParticipant Unmute a participant in a conference

		**Required ACL:** `calld.conferences.{conference_id}.participants.{participant_id}.unmute.update`

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param conferenceId Unique identifier of the conference
		@param participantId Unique identifier of the participant
		@return ConferencesAPIUnmuteParticipantRequest
	*/
	UnmuteParticipant(ctx context.Context, conferenceId string, participantId string) ConferencesAPIUnmuteParticipantRequest

	// UnmuteParticipantExecute executes the request
	UnmuteParticipantExecute(r ConferencesAPIUnmuteParticipantRequest) (*http.Response, error)
}

// ConferencesAPIService ConferencesAPI service
type ConferencesAPIService service

type ConferencesAPIKickParticipantRequest struct {
	ctx           context.Context
	ApiService    ConferencesAPI
	conferenceId  string
	participantId string
}

func (r ConferencesAPIKickParticipantRequest) Execute() (*http.Response, error) {
	return r.ApiService.KickParticipantExecute(r)
}

/*
KickParticipant Kick participant from a conference

**Required ACL:** `calld.conferences.{conference_id}.participants.{participant_id}.delete`

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param conferenceId Unique identifier of the conference
	@param participantId Unique identifier of the participant
	@return ConferencesAPIKickParticipantRequest
*/
func (a *ConferencesAPIService) KickParticipant(ctx context.Context, conferenceId string, participantId string) ConferencesAPIKickParticipantRequest {
	return ConferencesAPIKickParticipantRequest{
		ApiService:    a,
		ctx:           ctx,
		conferenceId:  conferenceId,
		participantId: participantId,
	}
}

// Execute executes the request
func (a *ConferencesAPIService) KickParticipantExecute(r ConferencesAPIKickParticipantRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ConferencesAPIService.KickParticipant")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/conferences/{conference_id}/participants/{participant_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"conference_id"+"}", url.PathEscape(parameterValueToString(r.conferenceId, "conferenceId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"participant_id"+"}", url.PathEscape(parameterValueToString(r.participantId, "participantId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["accent_auth_token"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["X-Auth-Token"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 503 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ConferencesAPIListConferenceParticipantsRequest struct {
	ctx          context.Context
	ApiService   ConferencesAPI
	conferenceId string
}

func (r ConferencesAPIListConferenceParticipantsRequest) Execute() (*ParticipantList, *http.Response, error) {
	return r.ApiService.ListConferenceParticipantsExecute(r)
}

/*
ListConferenceParticipants List participants of a conference

**Required ACL:** `calld.conferences.{conference_id}.participants.read`

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param conferenceId Unique identifier of the conference
	@return ConferencesAPIListConferenceParticipantsRequest
*/
func (a *ConferencesAPIService) ListConferenceParticipants(ctx context.Context, conferenceId string) ConferencesAPIListConferenceParticipantsRequest {
	return ConferencesAPIListConferenceParticipantsRequest{
		ApiService:   a,
		ctx:          ctx,
		conferenceId: conferenceId,
	}
}

// Execute executes the request
//
//	@return ParticipantList
func (a *ConferencesAPIService) ListConferenceParticipantsExecute(r ConferencesAPIListConferenceParticipantsRequest) (*ParticipantList, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ParticipantList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ConferencesAPIService.ListConferenceParticipants")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/conferences/{conference_id}/participants"
	localVarPath = strings.Replace(localVarPath, "{"+"conference_id"+"}", url.PathEscape(parameterValueToString(r.conferenceId, "conferenceId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["accent_auth_token"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["X-Auth-Token"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 503 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ConferencesAPIListUserConferenceParticipantsRequest struct {
	ctx          context.Context
	ApiService   ConferencesAPI
	conferenceId string
}

func (r ConferencesAPIListUserConferenceParticipantsRequest) Execute() (*ParticipantList, *http.Response, error) {
	return r.ApiService.ListUserConferenceParticipantsExecute(r)
}

/*
ListUserConferenceParticipants List participants of a conference as a user

**Required ACL:** `calld.users.me.conferences.{conference_id}.participants.read`

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param conferenceId Unique identifier of the conference
	@return ConferencesAPIListUserConferenceParticipantsRequest
*/
func (a *ConferencesAPIService) ListUserConferenceParticipants(ctx context.Context, conferenceId string) ConferencesAPIListUserConferenceParticipantsRequest {
	return ConferencesAPIListUserConferenceParticipantsRequest{
		ApiService:   a,
		ctx:          ctx,
		conferenceId: conferenceId,
	}
}

// Execute executes the request
//
//	@return ParticipantList
func (a *ConferencesAPIService) ListUserConferenceParticipantsExecute(r ConferencesAPIListUserConferenceParticipantsRequest) (*ParticipantList, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ParticipantList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ConferencesAPIService.ListUserConferenceParticipants")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/users/me/conferences/{conference_id}/participants"
	localVarPath = strings.Replace(localVarPath, "{"+"conference_id"+"}", url.PathEscape(parameterValueToString(r.conferenceId, "conferenceId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["accent_auth_token"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["X-Auth-Token"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 503 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ConferencesAPIMuteParticipantRequest struct {
	ctx           context.Context
	ApiService    ConferencesAPI
	conferenceId  string
	participantId string
}

func (r ConferencesAPIMuteParticipantRequest) Execute() (*http.Response, error) {
	return r.ApiService.MuteParticipantExecute(r)
}

/*
MuteParticipant Mute a participant in a conference

**Required ACL:** `calld.conferences.{conference_id}.participants.{participant_id}.mute.update`

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param conferenceId Unique identifier of the conference
	@param participantId Unique identifier of the participant
	@return ConferencesAPIMuteParticipantRequest
*/
func (a *ConferencesAPIService) MuteParticipant(ctx context.Context, conferenceId string, participantId string) ConferencesAPIMuteParticipantRequest {
	return ConferencesAPIMuteParticipantRequest{
		ApiService:    a,
		ctx:           ctx,
		conferenceId:  conferenceId,
		participantId: participantId,
	}
}

// Execute executes the request
func (a *ConferencesAPIService) MuteParticipantExecute(r ConferencesAPIMuteParticipantRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodPut
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ConferencesAPIService.MuteParticipant")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/conferences/{conference_id}/participants/{participant_id}/mute"
	localVarPath = strings.Replace(localVarPath, "{"+"conference_id"+"}", url.PathEscape(parameterValueToString(r.conferenceId, "conferenceId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"participant_id"+"}", url.PathEscape(parameterValueToString(r.participantId, "participantId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["accent_auth_token"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["X-Auth-Token"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 503 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ConferencesAPIStartConferenceRecordingRequest struct {
	ctx          context.Context
	ApiService   ConferencesAPI
	conferenceId string
}

func (r ConferencesAPIStartConferenceRecordingRequest) Execute() (*http.Response, error) {
	return r.ApiService.StartConferenceRecordingExecute(r)
}

/*
StartConferenceRecording Record a conference

**Required ACL:** `calld.conferences.{conference_id}.record.create`

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param conferenceId Unique identifier of the conference
	@return ConferencesAPIStartConferenceRecordingRequest
*/
func (a *ConferencesAPIService) StartConferenceRecording(ctx context.Context, conferenceId string) ConferencesAPIStartConferenceRecordingRequest {
	return ConferencesAPIStartConferenceRecordingRequest{
		ApiService:   a,
		ctx:          ctx,
		conferenceId: conferenceId,
	}
}

// Execute executes the request
func (a *ConferencesAPIService) StartConferenceRecordingExecute(r ConferencesAPIStartConferenceRecordingRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodPost
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ConferencesAPIService.StartConferenceRecording")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/conferences/{conference_id}/record"
	localVarPath = strings.Replace(localVarPath, "{"+"conference_id"+"}", url.PathEscape(parameterValueToString(r.conferenceId, "conferenceId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["accent_auth_token"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["X-Auth-Token"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 503 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ConferencesAPIStopConferenceRecordingRequest struct {
	ctx          context.Context
	ApiService   ConferencesAPI
	conferenceId string
}

func (r ConferencesAPIStopConferenceRecordingRequest) Execute() (*http.Response, error) {
	return r.ApiService.StopConferenceRecordingExecute(r)
}

/*
StopConferenceRecording Stop recording a conference

**Required ACL:** `calld.conferences.{conference_id}.record.delete`

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param conferenceId Unique identifier of the conference
	@return ConferencesAPIStopConferenceRecordingRequest
*/
func (a *ConferencesAPIService) StopConferenceRecording(ctx context.Context, conferenceId string) ConferencesAPIStopConferenceRecordingRequest {
	return ConferencesAPIStopConferenceRecordingRequest{
		ApiService:   a,
		ctx:          ctx,
		conferenceId: conferenceId,
	}
}

// Execute executes the request
func (a *ConferencesAPIService) StopConferenceRecordingExecute(r ConferencesAPIStopConferenceRecordingRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ConferencesAPIService.StopConferenceRecording")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/conferences/{conference_id}/record"
	localVarPath = strings.Replace(localVarPath, "{"+"conference_id"+"}", url.PathEscape(parameterValueToString(r.conferenceId, "conferenceId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["accent_auth_token"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["X-Auth-Token"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 503 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ConferencesAPIUnmuteParticipantRequest struct {
	ctx           context.Context
	ApiService    ConferencesAPI
	conferenceId  string
	participantId string
}

func (r ConferencesAPIUnmuteParticipantRequest) Execute() (*http.Response, error) {
	return r.ApiService.UnmuteParticipantExecute(r)
}

/*
UnmuteParticipant Unmute a participant in a conference

**Required ACL:** `calld.conferences.{conference_id}.participants.{participant_id}.unmute.update`

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param conferenceId Unique identifier of the conference
	@param participantId Unique identifier of the participant
	@return ConferencesAPIUnmuteParticipantRequest
*/
func (a *ConferencesAPIService) UnmuteParticipant(ctx context.Context, conferenceId string, participantId string) ConferencesAPIUnmuteParticipantRequest {
	return ConferencesAPIUnmuteParticipantRequest{
		ApiService:    a,
		ctx:           ctx,
		conferenceId:  conferenceId,
		participantId: participantId,
	}
}

// Execute executes the request
func (a *ConferencesAPIService) UnmuteParticipantExecute(r ConferencesAPIUnmuteParticipantRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodPut
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ConferencesAPIService.UnmuteParticipant")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/conferences/{conference_id}/participants/{participant_id}/unmute"
	localVarPath = strings.Replace(localVarPath, "{"+"conference_id"+"}", url.PathEscape(parameterValueToString(r.conferenceId, "conferenceId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"participant_id"+"}", url.PathEscape(parameterValueToString(r.participantId, "participantId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["accent_auth_token"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["X-Auth-Token"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 503 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}
