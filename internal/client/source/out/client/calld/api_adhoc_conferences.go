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

type AdhocConferencesAPI interface {

	/*
		AddParticipantToAdhocConference Add a participant into an adhoc conference

		**Required ACL:** `calld.users.me.conferences.adhoc.participants.update`

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param conferenceId ID of the adhoc conference
		@param callId ID of the call
		@return AdhocConferencesAPIAddParticipantToAdhocConferenceRequest
	*/
	AddParticipantToAdhocConference(ctx context.Context, conferenceId string, callId string) AdhocConferencesAPIAddParticipantToAdhocConferenceRequest

	// AddParticipantToAdhocConferenceExecute executes the request
	AddParticipantToAdhocConferenceExecute(r AdhocConferencesAPIAddParticipantToAdhocConferenceRequest) (*http.Response, error)

	/*
		CreateAdhocConference Create an adhoc conference

		**Required ACL:** `calld.users.me.conferences.adhoc.create`. An adhoc conference allows a user to merge multiple calls in one conversation. It acts like a conference room, but has no dedicated extension. The user creating the adhoc conference acts as the owner of the conference and controls who enters or leaves the conference. The conference will be destroyed when the owner leaves the conference.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@return AdhocConferencesAPICreateAdhocConferenceRequest
	*/
	CreateAdhocConference(ctx context.Context) AdhocConferencesAPICreateAdhocConferenceRequest

	// CreateAdhocConferenceExecute executes the request
	//  @return AdhocConference
	CreateAdhocConferenceExecute(r AdhocConferencesAPICreateAdhocConferenceRequest) (*AdhocConference, *http.Response, error)

	/*
		DeleteAdhocConference Delete an adhoc conference

		**Required ACL:** `calld.users.me.conferences.adhoc.delete`. All calls in the adhoc conference will be hungup.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param conferenceId ID of the adhoc conference
		@return AdhocConferencesAPIDeleteAdhocConferenceRequest
	*/
	DeleteAdhocConference(ctx context.Context, conferenceId string) AdhocConferencesAPIDeleteAdhocConferenceRequest

	// DeleteAdhocConferenceExecute executes the request
	//  @return AdhocConference
	DeleteAdhocConferenceExecute(r AdhocConferencesAPIDeleteAdhocConferenceRequest) (*AdhocConference, *http.Response, error)

	/*
		RemoveParticipantFromAdhocConference Remove a participant from an adhoc conference

		**Required ACL:** `calld.users.me.conferences.adhoc.participants.delete`

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param conferenceId ID of the adhoc conference
		@param callId ID of the call
		@return AdhocConferencesAPIRemoveParticipantFromAdhocConferenceRequest
	*/
	RemoveParticipantFromAdhocConference(ctx context.Context, conferenceId string, callId string) AdhocConferencesAPIRemoveParticipantFromAdhocConferenceRequest

	// RemoveParticipantFromAdhocConferenceExecute executes the request
	RemoveParticipantFromAdhocConferenceExecute(r AdhocConferencesAPIRemoveParticipantFromAdhocConferenceRequest) (*http.Response, error)
}

// AdhocConferencesAPIService AdhocConferencesAPI service
type AdhocConferencesAPIService service

type AdhocConferencesAPIAddParticipantToAdhocConferenceRequest struct {
	ctx          context.Context
	ApiService   AdhocConferencesAPI
	conferenceId string
	callId       string
}

func (r AdhocConferencesAPIAddParticipantToAdhocConferenceRequest) Execute() (*http.Response, error) {
	return r.ApiService.AddParticipantToAdhocConferenceExecute(r)
}

/*
AddParticipantToAdhocConference Add a participant into an adhoc conference

**Required ACL:** `calld.users.me.conferences.adhoc.participants.update`

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param conferenceId ID of the adhoc conference
	@param callId ID of the call
	@return AdhocConferencesAPIAddParticipantToAdhocConferenceRequest
*/
func (a *AdhocConferencesAPIService) AddParticipantToAdhocConference(ctx context.Context, conferenceId string, callId string) AdhocConferencesAPIAddParticipantToAdhocConferenceRequest {
	return AdhocConferencesAPIAddParticipantToAdhocConferenceRequest{
		ApiService:   a,
		ctx:          ctx,
		conferenceId: conferenceId,
		callId:       callId,
	}
}

// Execute executes the request
func (a *AdhocConferencesAPIService) AddParticipantToAdhocConferenceExecute(r AdhocConferencesAPIAddParticipantToAdhocConferenceRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodPut
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AdhocConferencesAPIService.AddParticipantToAdhocConference")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/users/me/conferences/adhoc/{conference_id}/participants/{call_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"conference_id"+"}", url.PathEscape(parameterValueToString(r.conferenceId, "conferenceId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"call_id"+"}", url.PathEscape(parameterValueToString(r.callId, "callId")), -1)

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

type AdhocConferencesAPICreateAdhocConferenceRequest struct {
	ctx        context.Context
	ApiService AdhocConferencesAPI
	body       *AdhocConferenceCreation
}

// Parameters of the conference calls
func (r AdhocConferencesAPICreateAdhocConferenceRequest) Body(body AdhocConferenceCreation) AdhocConferencesAPICreateAdhocConferenceRequest {
	r.body = &body
	return r
}

func (r AdhocConferencesAPICreateAdhocConferenceRequest) Execute() (*AdhocConference, *http.Response, error) {
	return r.ApiService.CreateAdhocConferenceExecute(r)
}

/*
CreateAdhocConference Create an adhoc conference

**Required ACL:** `calld.users.me.conferences.adhoc.create`. An adhoc conference allows a user to merge multiple calls in one conversation. It acts like a conference room, but has no dedicated extension. The user creating the adhoc conference acts as the owner of the conference and controls who enters or leaves the conference. The conference will be destroyed when the owner leaves the conference.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return AdhocConferencesAPICreateAdhocConferenceRequest
*/
func (a *AdhocConferencesAPIService) CreateAdhocConference(ctx context.Context) AdhocConferencesAPICreateAdhocConferenceRequest {
	return AdhocConferencesAPICreateAdhocConferenceRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return AdhocConference
func (a *AdhocConferencesAPIService) CreateAdhocConferenceExecute(r AdhocConferencesAPICreateAdhocConferenceRequest) (*AdhocConference, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *AdhocConference
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AdhocConferencesAPIService.CreateAdhocConference")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/users/me/conferences/adhoc"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.body == nil {
		return localVarReturnValue, nil, reportError("body is required and must be specified")
	}

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
	// body params
	localVarPostBody = r.body
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
		if localVarHTTPResponse.StatusCode == 400 {
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
		if localVarHTTPResponse.StatusCode == 409 {
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

type AdhocConferencesAPIDeleteAdhocConferenceRequest struct {
	ctx          context.Context
	ApiService   AdhocConferencesAPI
	conferenceId string
}

func (r AdhocConferencesAPIDeleteAdhocConferenceRequest) Execute() (*AdhocConference, *http.Response, error) {
	return r.ApiService.DeleteAdhocConferenceExecute(r)
}

/*
DeleteAdhocConference Delete an adhoc conference

**Required ACL:** `calld.users.me.conferences.adhoc.delete`. All calls in the adhoc conference will be hungup.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param conferenceId ID of the adhoc conference
	@return AdhocConferencesAPIDeleteAdhocConferenceRequest
*/
func (a *AdhocConferencesAPIService) DeleteAdhocConference(ctx context.Context, conferenceId string) AdhocConferencesAPIDeleteAdhocConferenceRequest {
	return AdhocConferencesAPIDeleteAdhocConferenceRequest{
		ApiService:   a,
		ctx:          ctx,
		conferenceId: conferenceId,
	}
}

// Execute executes the request
//
//	@return AdhocConference
func (a *AdhocConferencesAPIService) DeleteAdhocConferenceExecute(r AdhocConferencesAPIDeleteAdhocConferenceRequest) (*AdhocConference, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodDelete
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *AdhocConference
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AdhocConferencesAPIService.DeleteAdhocConference")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/users/me/conferences/adhoc/{conference_id}"
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

type AdhocConferencesAPIRemoveParticipantFromAdhocConferenceRequest struct {
	ctx          context.Context
	ApiService   AdhocConferencesAPI
	conferenceId string
	callId       string
}

func (r AdhocConferencesAPIRemoveParticipantFromAdhocConferenceRequest) Execute() (*http.Response, error) {
	return r.ApiService.RemoveParticipantFromAdhocConferenceExecute(r)
}

/*
RemoveParticipantFromAdhocConference Remove a participant from an adhoc conference

**Required ACL:** `calld.users.me.conferences.adhoc.participants.delete`

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param conferenceId ID of the adhoc conference
	@param callId ID of the call
	@return AdhocConferencesAPIRemoveParticipantFromAdhocConferenceRequest
*/
func (a *AdhocConferencesAPIService) RemoveParticipantFromAdhocConference(ctx context.Context, conferenceId string, callId string) AdhocConferencesAPIRemoveParticipantFromAdhocConferenceRequest {
	return AdhocConferencesAPIRemoveParticipantFromAdhocConferenceRequest{
		ApiService:   a,
		ctx:          ctx,
		conferenceId: conferenceId,
		callId:       callId,
	}
}

// Execute executes the request
func (a *AdhocConferencesAPIService) RemoveParticipantFromAdhocConferenceExecute(r AdhocConferencesAPIRemoveParticipantFromAdhocConferenceRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AdhocConferencesAPIService.RemoveParticipantFromAdhocConference")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/users/me/conferences/adhoc/{conference_id}/participants/{call_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"conference_id"+"}", url.PathEscape(parameterValueToString(r.conferenceId, "conferenceId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"call_id"+"}", url.PathEscape(parameterValueToString(r.callId, "callId")), -1)

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
