/*
accent-agentd

Agent service

API version: 1.0
Contact: help@accentvoice.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package agentd

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
)

type AgentsAPI interface {

	/*
		GetAgents Get the status of all agents.

		**Required ACL:** `agentd.agents.read`

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@return AgentsAPIGetAgentsRequest
	*/
	GetAgents(ctx context.Context) AgentsAPIGetAgentsRequest

	// GetAgentsExecute executes the request
	//  @return AgentStatus
	GetAgentsExecute(r AgentsAPIGetAgentsRequest) (*AgentStatus, *http.Response, error)

	/*
		LogoffAgents Logoff all agents.

		**Required ACL:** `agentd.agents.logoff.create`

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@return AgentsAPILogoffAgentsRequest
	*/
	LogoffAgents(ctx context.Context) AgentsAPILogoffAgentsRequest

	// LogoffAgentsExecute executes the request
	LogoffAgentsExecute(r AgentsAPILogoffAgentsRequest) (*http.Response, error)

	/*
		RelogAgents Relog all agents.

		**Required ACL:** `agentd.agents.relog.create`

	Relog all agents which are currently logged.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@return AgentsAPIRelogAgentsRequest
	*/
	RelogAgents(ctx context.Context) AgentsAPIRelogAgentsRequest

	// RelogAgentsExecute executes the request
	RelogAgentsExecute(r AgentsAPIRelogAgentsRequest) (*http.Response, error)
}

// AgentsAPIService AgentsAPI service
type AgentsAPIService service

type AgentsAPIGetAgentsRequest struct {
	ctx          context.Context
	ApiService   AgentsAPI
	accentTenant *string
	recurse      *bool
}

// The tenant&#39;s UUID, defining the ownership of a given resource.
func (r AgentsAPIGetAgentsRequest) AccentTenant(accentTenant string) AgentsAPIGetAgentsRequest {
	r.accentTenant = &accentTenant
	return r
}

// Should the query include sub-tenants
func (r AgentsAPIGetAgentsRequest) Recurse(recurse bool) AgentsAPIGetAgentsRequest {
	r.recurse = &recurse
	return r
}

func (r AgentsAPIGetAgentsRequest) Execute() (*AgentStatus, *http.Response, error) {
	return r.ApiService.GetAgentsExecute(r)
}

/*
GetAgents Get the status of all agents.

**Required ACL:** `agentd.agents.read`

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return AgentsAPIGetAgentsRequest
*/
func (a *AgentsAPIService) GetAgents(ctx context.Context) AgentsAPIGetAgentsRequest {
	return AgentsAPIGetAgentsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return AgentStatus
func (a *AgentsAPIService) GetAgentsExecute(r AgentsAPIGetAgentsRequest) (*AgentStatus, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *AgentStatus
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AgentsAPIService.GetAgents")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/agents"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.recurse != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "recurse", r.recurse, "")
	} else {
		var defaultValue bool = false
		r.recurse = &defaultValue
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
	if r.accentTenant != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "Accent-Tenant", r.accentTenant, "")
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

type AgentsAPILogoffAgentsRequest struct {
	ctx          context.Context
	ApiService   AgentsAPI
	accentTenant *string
}

// The tenant&#39;s UUID, defining the ownership of a given resource.
func (r AgentsAPILogoffAgentsRequest) AccentTenant(accentTenant string) AgentsAPILogoffAgentsRequest {
	r.accentTenant = &accentTenant
	return r
}

func (r AgentsAPILogoffAgentsRequest) Execute() (*http.Response, error) {
	return r.ApiService.LogoffAgentsExecute(r)
}

/*
LogoffAgents Logoff all agents.

**Required ACL:** `agentd.agents.logoff.create`

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return AgentsAPILogoffAgentsRequest
*/
func (a *AgentsAPIService) LogoffAgents(ctx context.Context) AgentsAPILogoffAgentsRequest {
	return AgentsAPILogoffAgentsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
func (a *AgentsAPIService) LogoffAgentsExecute(r AgentsAPILogoffAgentsRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodPost
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AgentsAPIService.LogoffAgents")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/agents/logoff"

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
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.accentTenant != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "Accent-Tenant", r.accentTenant, "")
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
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type AgentsAPIRelogAgentsRequest struct {
	ctx          context.Context
	ApiService   AgentsAPI
	accentTenant *string
}

// The tenant&#39;s UUID, defining the ownership of a given resource.
func (r AgentsAPIRelogAgentsRequest) AccentTenant(accentTenant string) AgentsAPIRelogAgentsRequest {
	r.accentTenant = &accentTenant
	return r
}

func (r AgentsAPIRelogAgentsRequest) Execute() (*http.Response, error) {
	return r.ApiService.RelogAgentsExecute(r)
}

/*
RelogAgents Relog all agents.

**Required ACL:** `agentd.agents.relog.create`

Relog all agents which are currently logged.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return AgentsAPIRelogAgentsRequest
*/
func (a *AgentsAPIService) RelogAgents(ctx context.Context) AgentsAPIRelogAgentsRequest {
	return AgentsAPIRelogAgentsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
func (a *AgentsAPIService) RelogAgentsExecute(r AgentsAPIRelogAgentsRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodPost
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AgentsAPIService.RelogAgents")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/agents/relog"

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
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.accentTenant != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "Accent-Tenant", r.accentTenant, "")
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
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}
