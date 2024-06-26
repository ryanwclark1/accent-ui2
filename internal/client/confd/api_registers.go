/*
accent-confd

Confd exposes an API for managing core resources on a Accent server such as users, extensions, devices, voicemails, queues, etc. Resources can be associated together to provide additional functionality. For example: By associating a voicemail with a user, calls will automatically fallback on to the voicemail when the user cannot answer.  Implementation notes ====================  Errors ------  Responses containing errors will have a status code in the 400 or 500 class. A list of error messages will be returned in the body of the response as a JSON-encoded array:  ~~~ [     \"Input error - User not found\",     \"Resource error - User not associated to a line\" ] ~~~   Updating resources via PUT --------------------------  When updating a resource, all fields become optional. In other words, only values that have been changed need to be sent to the server. Please note that this behavior may change in future versions of the API.

API version: 1.1
Contact: help@accentvoice.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package confd

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)

type RegistersAPI interface {

	/*
		AssociateTrunkRegisterIax Associate trunk and IAX register

		**Required ACL:** `confd.trunks.{trunk_id}.registers.iax.{iax_id}.update`

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param trunkId Trunk's ID
		@param iaxId
		@return RegistersAPIAssociateTrunkRegisterIaxRequest
	*/
	AssociateTrunkRegisterIax(ctx context.Context, trunkId int32, iaxId int32) RegistersAPIAssociateTrunkRegisterIaxRequest

	// AssociateTrunkRegisterIaxExecute executes the request
	AssociateTrunkRegisterIaxExecute(r RegistersAPIAssociateTrunkRegisterIaxRequest) (*http.Response, error)

	/*
		CreateRegisterIax Create register_iax

		**Required ACL:** `confd.registers.iax.create`

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@return RegistersAPICreateRegisterIaxRequest
	*/
	CreateRegisterIax(ctx context.Context) RegistersAPICreateRegisterIaxRequest

	// CreateRegisterIaxExecute executes the request
	//  @return RegisterIAX
	CreateRegisterIaxExecute(r RegistersAPICreateRegisterIaxRequest) (*RegisterIAX, *http.Response, error)

	/*
		DeleteRegisterIax Delete register IAX

		**Required ACL:** `confd.registers.iax.{register_iax_id}.delete`

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param registerIaxId Register IAX's ID
		@return RegistersAPIDeleteRegisterIaxRequest
	*/
	DeleteRegisterIax(ctx context.Context, registerIaxId int32) RegistersAPIDeleteRegisterIaxRequest

	// DeleteRegisterIaxExecute executes the request
	DeleteRegisterIaxExecute(r RegistersAPIDeleteRegisterIaxRequest) (*http.Response, error)

	/*
		DissociateTrunkRegisterIax Dissociate trunk and IAX register

		**Required ACL:** `confd.trunks.{trunk_id}.registers.iax.{iax_id}.delete`

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param trunkId Trunk's ID
		@param iaxId
		@return RegistersAPIDissociateTrunkRegisterIaxRequest
	*/
	DissociateTrunkRegisterIax(ctx context.Context, trunkId int32, iaxId int32) RegistersAPIDissociateTrunkRegisterIaxRequest

	// DissociateTrunkRegisterIaxExecute executes the request
	DissociateTrunkRegisterIaxExecute(r RegistersAPIDissociateTrunkRegisterIaxRequest) (*http.Response, error)

	/*
		GetRegisterIax Get register IAX

		**Required ACL:** `confd.registers.iax.{register_iax_id}.read`

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param registerIaxId Register IAX's ID
		@return RegistersAPIGetRegisterIaxRequest
	*/
	GetRegisterIax(ctx context.Context, registerIaxId int32) RegistersAPIGetRegisterIaxRequest

	// GetRegisterIaxExecute executes the request
	//  @return RegisterIAX
	GetRegisterIaxExecute(r RegistersAPIGetRegisterIaxRequest) (*RegisterIAX, *http.Response, error)

	/*
		ListRegistersIax List registers iax

		**Required ACL:** `confd.registers.iax.read`

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@return RegistersAPIListRegistersIaxRequest
	*/
	ListRegistersIax(ctx context.Context) RegistersAPIListRegistersIaxRequest

	// ListRegistersIaxExecute executes the request
	//  @return RegisterIAXItems
	ListRegistersIaxExecute(r RegistersAPIListRegistersIaxRequest) (*RegisterIAXItems, *http.Response, error)

	/*
		UpdateRegisterIax Update register IAX

		**Required ACL:** `confd.registers.iax.{register_iax_id}.update`

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param registerIaxId Register IAX's ID
		@return RegistersAPIUpdateRegisterIaxRequest
	*/
	UpdateRegisterIax(ctx context.Context, registerIaxId int32) RegistersAPIUpdateRegisterIaxRequest

	// UpdateRegisterIaxExecute executes the request
	UpdateRegisterIaxExecute(r RegistersAPIUpdateRegisterIaxRequest) (*http.Response, error)
}

// RegistersAPIService RegistersAPI service
type RegistersAPIService service

type RegistersAPIAssociateTrunkRegisterIaxRequest struct {
	ctx        context.Context
	ApiService RegistersAPI
	trunkId    int32
	iaxId      int32
}

func (r RegistersAPIAssociateTrunkRegisterIaxRequest) Execute() (*http.Response, error) {
	return r.ApiService.AssociateTrunkRegisterIaxExecute(r)
}

/*
AssociateTrunkRegisterIax Associate trunk and IAX register

**Required ACL:** `confd.trunks.{trunk_id}.registers.iax.{iax_id}.update`

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param trunkId Trunk's ID
	@param iaxId
	@return RegistersAPIAssociateTrunkRegisterIaxRequest
*/
func (a *RegistersAPIService) AssociateTrunkRegisterIax(ctx context.Context, trunkId int32, iaxId int32) RegistersAPIAssociateTrunkRegisterIaxRequest {
	return RegistersAPIAssociateTrunkRegisterIaxRequest{
		ApiService: a,
		ctx:        ctx,
		trunkId:    trunkId,
		iaxId:      iaxId,
	}
}

// Execute executes the request
func (a *RegistersAPIService) AssociateTrunkRegisterIaxExecute(r RegistersAPIAssociateTrunkRegisterIaxRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodPut
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RegistersAPIService.AssociateTrunkRegisterIax")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/trunks/{trunk_id}/registers/iax/{iax_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"trunk_id"+"}", url.PathEscape(parameterValueToString(r.trunkId, "trunkId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"iax_id"+"}", url.PathEscape(parameterValueToString(r.iaxId, "iaxId")), -1)

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
			var v []string
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
			var v []string
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

type RegistersAPICreateRegisterIaxRequest struct {
	ctx        context.Context
	ApiService RegistersAPI
	body       *RegisterIAX
}

// Register iax to create
func (r RegistersAPICreateRegisterIaxRequest) Body(body RegisterIAX) RegistersAPICreateRegisterIaxRequest {
	r.body = &body
	return r
}

func (r RegistersAPICreateRegisterIaxRequest) Execute() (*RegisterIAX, *http.Response, error) {
	return r.ApiService.CreateRegisterIaxExecute(r)
}

/*
CreateRegisterIax Create register_iax

**Required ACL:** `confd.registers.iax.create`

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return RegistersAPICreateRegisterIaxRequest
*/
func (a *RegistersAPIService) CreateRegisterIax(ctx context.Context) RegistersAPICreateRegisterIaxRequest {
	return RegistersAPICreateRegisterIaxRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return RegisterIAX
func (a *RegistersAPIService) CreateRegisterIaxExecute(r RegistersAPICreateRegisterIaxRequest) (*RegisterIAX, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *RegisterIAX
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RegistersAPIService.CreateRegisterIax")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/registers/iax"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.body == nil {
		return localVarReturnValue, nil, reportError("body is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

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
			var v []string
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

type RegistersAPIDeleteRegisterIaxRequest struct {
	ctx           context.Context
	ApiService    RegistersAPI
	registerIaxId int32
}

func (r RegistersAPIDeleteRegisterIaxRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteRegisterIaxExecute(r)
}

/*
DeleteRegisterIax Delete register IAX

**Required ACL:** `confd.registers.iax.{register_iax_id}.delete`

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param registerIaxId Register IAX's ID
	@return RegistersAPIDeleteRegisterIaxRequest
*/
func (a *RegistersAPIService) DeleteRegisterIax(ctx context.Context, registerIaxId int32) RegistersAPIDeleteRegisterIaxRequest {
	return RegistersAPIDeleteRegisterIaxRequest{
		ApiService:    a,
		ctx:           ctx,
		registerIaxId: registerIaxId,
	}
}

// Execute executes the request
func (a *RegistersAPIService) DeleteRegisterIaxExecute(r RegistersAPIDeleteRegisterIaxRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RegistersAPIService.DeleteRegisterIax")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/registers/iax/{register_iax_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"register_iax_id"+"}", url.PathEscape(parameterValueToString(r.registerIaxId, "registerIaxId")), -1)

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
			var v []string
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
			var v []string
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

type RegistersAPIDissociateTrunkRegisterIaxRequest struct {
	ctx        context.Context
	ApiService RegistersAPI
	trunkId    int32
	iaxId      int32
}

func (r RegistersAPIDissociateTrunkRegisterIaxRequest) Execute() (*http.Response, error) {
	return r.ApiService.DissociateTrunkRegisterIaxExecute(r)
}

/*
DissociateTrunkRegisterIax Dissociate trunk and IAX register

**Required ACL:** `confd.trunks.{trunk_id}.registers.iax.{iax_id}.delete`

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param trunkId Trunk's ID
	@param iaxId
	@return RegistersAPIDissociateTrunkRegisterIaxRequest
*/
func (a *RegistersAPIService) DissociateTrunkRegisterIax(ctx context.Context, trunkId int32, iaxId int32) RegistersAPIDissociateTrunkRegisterIaxRequest {
	return RegistersAPIDissociateTrunkRegisterIaxRequest{
		ApiService: a,
		ctx:        ctx,
		trunkId:    trunkId,
		iaxId:      iaxId,
	}
}

// Execute executes the request
func (a *RegistersAPIService) DissociateTrunkRegisterIaxExecute(r RegistersAPIDissociateTrunkRegisterIaxRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RegistersAPIService.DissociateTrunkRegisterIax")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/trunks/{trunk_id}/registers/iax/{iax_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"trunk_id"+"}", url.PathEscape(parameterValueToString(r.trunkId, "trunkId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"iax_id"+"}", url.PathEscape(parameterValueToString(r.iaxId, "iaxId")), -1)

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
			var v []string
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
			var v []string
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

type RegistersAPIGetRegisterIaxRequest struct {
	ctx           context.Context
	ApiService    RegistersAPI
	registerIaxId int32
}

func (r RegistersAPIGetRegisterIaxRequest) Execute() (*RegisterIAX, *http.Response, error) {
	return r.ApiService.GetRegisterIaxExecute(r)
}

/*
GetRegisterIax Get register IAX

**Required ACL:** `confd.registers.iax.{register_iax_id}.read`

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param registerIaxId Register IAX's ID
	@return RegistersAPIGetRegisterIaxRequest
*/
func (a *RegistersAPIService) GetRegisterIax(ctx context.Context, registerIaxId int32) RegistersAPIGetRegisterIaxRequest {
	return RegistersAPIGetRegisterIaxRequest{
		ApiService:    a,
		ctx:           ctx,
		registerIaxId: registerIaxId,
	}
}

// Execute executes the request
//
//	@return RegisterIAX
func (a *RegistersAPIService) GetRegisterIaxExecute(r RegistersAPIGetRegisterIaxRequest) (*RegisterIAX, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *RegisterIAX
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RegistersAPIService.GetRegisterIax")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/registers/iax/{register_iax_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"register_iax_id"+"}", url.PathEscape(parameterValueToString(r.registerIaxId, "registerIaxId")), -1)

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
			var v []string
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

type RegistersAPIListRegistersIaxRequest struct {
	ctx        context.Context
	ApiService RegistersAPI
	order      *string
	direction  *string
	limit      *int32
	offset     *int32
	search     *string
}

// Name of the field to use for sorting the list of items returned.
func (r RegistersAPIListRegistersIaxRequest) Order(order string) RegistersAPIListRegistersIaxRequest {
	r.order = &order
	return r
}

// Sort list of items in &#39;asc&#39; (ascending) or &#39;desc&#39; (descending) order
func (r RegistersAPIListRegistersIaxRequest) Direction(direction string) RegistersAPIListRegistersIaxRequest {
	r.direction = &direction
	return r
}

// Maximum number of items to return in the list
func (r RegistersAPIListRegistersIaxRequest) Limit(limit int32) RegistersAPIListRegistersIaxRequest {
	r.limit = &limit
	return r
}

// Number of items to skip over in the list. Useful for pagination.
func (r RegistersAPIListRegistersIaxRequest) Offset(offset int32) RegistersAPIListRegistersIaxRequest {
	r.offset = &offset
	return r
}

// Search term for filtering a list of items. Only items with a field containing the search term will be returned.
func (r RegistersAPIListRegistersIaxRequest) Search(search string) RegistersAPIListRegistersIaxRequest {
	r.search = &search
	return r
}

func (r RegistersAPIListRegistersIaxRequest) Execute() (*RegisterIAXItems, *http.Response, error) {
	return r.ApiService.ListRegistersIaxExecute(r)
}

/*
ListRegistersIax List registers iax

**Required ACL:** `confd.registers.iax.read`

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return RegistersAPIListRegistersIaxRequest
*/
func (a *RegistersAPIService) ListRegistersIax(ctx context.Context) RegistersAPIListRegistersIaxRequest {
	return RegistersAPIListRegistersIaxRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return RegisterIAXItems
func (a *RegistersAPIService) ListRegistersIaxExecute(r RegistersAPIListRegistersIaxRequest) (*RegisterIAXItems, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *RegisterIAXItems
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RegistersAPIService.ListRegistersIax")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/registers/iax"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.order != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "order", r.order, "")
	}
	if r.direction != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "direction", r.direction, "")
	}
	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "")
	}
	if r.offset != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "offset", r.offset, "")
	}
	if r.search != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "search", r.search, "")
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

type RegistersAPIUpdateRegisterIaxRequest struct {
	ctx           context.Context
	ApiService    RegistersAPI
	body          *RegisterIAX
	registerIaxId int32
}

func (r RegistersAPIUpdateRegisterIaxRequest) Body(body RegisterIAX) RegistersAPIUpdateRegisterIaxRequest {
	r.body = &body
	return r
}

func (r RegistersAPIUpdateRegisterIaxRequest) Execute() (*http.Response, error) {
	return r.ApiService.UpdateRegisterIaxExecute(r)
}

/*
UpdateRegisterIax Update register IAX

**Required ACL:** `confd.registers.iax.{register_iax_id}.update`

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param registerIaxId Register IAX's ID
	@return RegistersAPIUpdateRegisterIaxRequest
*/
func (a *RegistersAPIService) UpdateRegisterIax(ctx context.Context, registerIaxId int32) RegistersAPIUpdateRegisterIaxRequest {
	return RegistersAPIUpdateRegisterIaxRequest{
		ApiService:    a,
		ctx:           ctx,
		registerIaxId: registerIaxId,
	}
}

// Execute executes the request
func (a *RegistersAPIService) UpdateRegisterIaxExecute(r RegistersAPIUpdateRegisterIaxRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodPut
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RegistersAPIService.UpdateRegisterIax")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/registers/iax/{register_iax_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"register_iax_id"+"}", url.PathEscape(parameterValueToString(r.registerIaxId, "registerIaxId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.body == nil {
		return nil, reportError("body is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

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
			var v []string
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
			var v []string
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
