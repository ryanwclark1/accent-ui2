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

type RegistrarsAPI interface {

	/*
		CreateRegistrar Create registrar

		**Required ACL:** `confd.registrars.create`

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@return RegistrarsAPICreateRegistrarRequest
	*/
	CreateRegistrar(ctx context.Context) RegistrarsAPICreateRegistrarRequest

	// CreateRegistrarExecute executes the request
	//  @return Registrar
	CreateRegistrarExecute(r RegistrarsAPICreateRegistrarRequest) (*Registrar, *http.Response, error)

	/*
		DeleteRegistrar Delete registrar

		**Required ACL:** `confd.registrars.{registrar_id}.delete`

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param registrarId Registrar ID
		@return RegistrarsAPIDeleteRegistrarRequest
	*/
	DeleteRegistrar(ctx context.Context, registrarId string) RegistrarsAPIDeleteRegistrarRequest

	// DeleteRegistrarExecute executes the request
	DeleteRegistrarExecute(r RegistrarsAPIDeleteRegistrarRequest) (*http.Response, error)

	/*
		GetRegistrar Get registrar

		**Required ACL:** `confd.registrars.{registrar_id}.read`

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param registrarId Registrar ID
		@return RegistrarsAPIGetRegistrarRequest
	*/
	GetRegistrar(ctx context.Context, registrarId string) RegistrarsAPIGetRegistrarRequest

	// GetRegistrarExecute executes the request
	//  @return Registrar
	GetRegistrarExecute(r RegistrarsAPIGetRegistrarRequest) (*Registrar, *http.Response, error)

	/*
		GetRegistrars Get registrars

		**Required ACL:** `confd.registrars.read`

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@return RegistrarsAPIGetRegistrarsRequest
	*/
	GetRegistrars(ctx context.Context) RegistrarsAPIGetRegistrarsRequest

	// GetRegistrarsExecute executes the request
	//  @return RegistrarItems
	GetRegistrarsExecute(r RegistrarsAPIGetRegistrarsRequest) (*RegistrarItems, *http.Response, error)

	/*
		UpdateRegistrar Update registrar

		**Required ACL:** `confd.registrars.{registrar_id}.update`

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param registrarId Registrar ID
		@return RegistrarsAPIUpdateRegistrarRequest
	*/
	UpdateRegistrar(ctx context.Context, registrarId string) RegistrarsAPIUpdateRegistrarRequest

	// UpdateRegistrarExecute executes the request
	UpdateRegistrarExecute(r RegistrarsAPIUpdateRegistrarRequest) (*http.Response, error)
}

// RegistrarsAPIService RegistrarsAPI service
type RegistrarsAPIService service

type RegistrarsAPICreateRegistrarRequest struct {
	ctx        context.Context
	ApiService RegistrarsAPI
	body       *Registrar
}

// Registrar to create
func (r RegistrarsAPICreateRegistrarRequest) Body(body Registrar) RegistrarsAPICreateRegistrarRequest {
	r.body = &body
	return r
}

func (r RegistrarsAPICreateRegistrarRequest) Execute() (*Registrar, *http.Response, error) {
	return r.ApiService.CreateRegistrarExecute(r)
}

/*
CreateRegistrar Create registrar

**Required ACL:** `confd.registrars.create`

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return RegistrarsAPICreateRegistrarRequest
*/
func (a *RegistrarsAPIService) CreateRegistrar(ctx context.Context) RegistrarsAPICreateRegistrarRequest {
	return RegistrarsAPICreateRegistrarRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return Registrar
func (a *RegistrarsAPIService) CreateRegistrarExecute(r RegistrarsAPICreateRegistrarRequest) (*Registrar, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *Registrar
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RegistrarsAPIService.CreateRegistrar")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/registrars"

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

type RegistrarsAPIDeleteRegistrarRequest struct {
	ctx         context.Context
	ApiService  RegistrarsAPI
	registrarId string
}

func (r RegistrarsAPIDeleteRegistrarRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteRegistrarExecute(r)
}

/*
DeleteRegistrar Delete registrar

**Required ACL:** `confd.registrars.{registrar_id}.delete`

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param registrarId Registrar ID
	@return RegistrarsAPIDeleteRegistrarRequest
*/
func (a *RegistrarsAPIService) DeleteRegistrar(ctx context.Context, registrarId string) RegistrarsAPIDeleteRegistrarRequest {
	return RegistrarsAPIDeleteRegistrarRequest{
		ApiService:  a,
		ctx:         ctx,
		registrarId: registrarId,
	}
}

// Execute executes the request
func (a *RegistrarsAPIService) DeleteRegistrarExecute(r RegistrarsAPIDeleteRegistrarRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RegistrarsAPIService.DeleteRegistrar")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/registrars/{registrar_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"registrar_id"+"}", url.PathEscape(parameterValueToString(r.registrarId, "registrarId")), -1)

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

type RegistrarsAPIGetRegistrarRequest struct {
	ctx         context.Context
	ApiService  RegistrarsAPI
	registrarId string
}

func (r RegistrarsAPIGetRegistrarRequest) Execute() (*Registrar, *http.Response, error) {
	return r.ApiService.GetRegistrarExecute(r)
}

/*
GetRegistrar Get registrar

**Required ACL:** `confd.registrars.{registrar_id}.read`

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param registrarId Registrar ID
	@return RegistrarsAPIGetRegistrarRequest
*/
func (a *RegistrarsAPIService) GetRegistrar(ctx context.Context, registrarId string) RegistrarsAPIGetRegistrarRequest {
	return RegistrarsAPIGetRegistrarRequest{
		ApiService:  a,
		ctx:         ctx,
		registrarId: registrarId,
	}
}

// Execute executes the request
//
//	@return Registrar
func (a *RegistrarsAPIService) GetRegistrarExecute(r RegistrarsAPIGetRegistrarRequest) (*Registrar, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *Registrar
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RegistrarsAPIService.GetRegistrar")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/registrars/{registrar_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"registrar_id"+"}", url.PathEscape(parameterValueToString(r.registrarId, "registrarId")), -1)

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

type RegistrarsAPIGetRegistrarsRequest struct {
	ctx        context.Context
	ApiService RegistrarsAPI
	order      *string
	direction  *string
	limit      *int32
	offset     *int32
	search     *string
}

// Name of the field to use for sorting the list of items returned.
func (r RegistrarsAPIGetRegistrarsRequest) Order(order string) RegistrarsAPIGetRegistrarsRequest {
	r.order = &order
	return r
}

// Sort list of items in &#39;asc&#39; (ascending) or &#39;desc&#39; (descending) order
func (r RegistrarsAPIGetRegistrarsRequest) Direction(direction string) RegistrarsAPIGetRegistrarsRequest {
	r.direction = &direction
	return r
}

// Maximum number of items to return in the list
func (r RegistrarsAPIGetRegistrarsRequest) Limit(limit int32) RegistrarsAPIGetRegistrarsRequest {
	r.limit = &limit
	return r
}

// Number of items to skip over in the list. Useful for pagination.
func (r RegistrarsAPIGetRegistrarsRequest) Offset(offset int32) RegistrarsAPIGetRegistrarsRequest {
	r.offset = &offset
	return r
}

// Search term for filtering a list of items. Only items with a field containing the search term will be returned.
func (r RegistrarsAPIGetRegistrarsRequest) Search(search string) RegistrarsAPIGetRegistrarsRequest {
	r.search = &search
	return r
}

func (r RegistrarsAPIGetRegistrarsRequest) Execute() (*RegistrarItems, *http.Response, error) {
	return r.ApiService.GetRegistrarsExecute(r)
}

/*
GetRegistrars Get registrars

**Required ACL:** `confd.registrars.read`

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return RegistrarsAPIGetRegistrarsRequest
*/
func (a *RegistrarsAPIService) GetRegistrars(ctx context.Context) RegistrarsAPIGetRegistrarsRequest {
	return RegistrarsAPIGetRegistrarsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return RegistrarItems
func (a *RegistrarsAPIService) GetRegistrarsExecute(r RegistrarsAPIGetRegistrarsRequest) (*RegistrarItems, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *RegistrarItems
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RegistrarsAPIService.GetRegistrars")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/registrars"

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

type RegistrarsAPIUpdateRegistrarRequest struct {
	ctx         context.Context
	ApiService  RegistrarsAPI
	body        *Registrar
	registrarId string
}

// Registrar body
func (r RegistrarsAPIUpdateRegistrarRequest) Body(body Registrar) RegistrarsAPIUpdateRegistrarRequest {
	r.body = &body
	return r
}

func (r RegistrarsAPIUpdateRegistrarRequest) Execute() (*http.Response, error) {
	return r.ApiService.UpdateRegistrarExecute(r)
}

/*
UpdateRegistrar Update registrar

**Required ACL:** `confd.registrars.{registrar_id}.update`

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param registrarId Registrar ID
	@return RegistrarsAPIUpdateRegistrarRequest
*/
func (a *RegistrarsAPIService) UpdateRegistrar(ctx context.Context, registrarId string) RegistrarsAPIUpdateRegistrarRequest {
	return RegistrarsAPIUpdateRegistrarRequest{
		ApiService:  a,
		ctx:         ctx,
		registrarId: registrarId,
	}
}

// Execute executes the request
func (a *RegistrarsAPIService) UpdateRegistrarExecute(r RegistrarsAPIUpdateRegistrarRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodPut
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RegistrarsAPIService.UpdateRegistrar")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/registrars/{registrar_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"registrar_id"+"}", url.PathEscape(parameterValueToString(r.registrarId, "registrarId")), -1)

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
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}
