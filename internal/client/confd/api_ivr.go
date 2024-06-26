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

type IvrAPI interface {

	/*
		CreateIvr Create IVR

		**Required ACL:** `confd.ivr.create`

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@return IvrAPICreateIvrRequest
	*/
	CreateIvr(ctx context.Context) IvrAPICreateIvrRequest

	// CreateIvrExecute executes the request
	//  @return Ivr
	CreateIvrExecute(r IvrAPICreateIvrRequest) (*Ivr, *http.Response, error)

	/*
		DeleteIvr Delete IVR

		**Required ACL:** `confd.ivr.{ivr_id}.delete`

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param ivrId IVR's ID
		@return IvrAPIDeleteIvrRequest
	*/
	DeleteIvr(ctx context.Context, ivrId int32) IvrAPIDeleteIvrRequest

	// DeleteIvrExecute executes the request
	DeleteIvrExecute(r IvrAPIDeleteIvrRequest) (*http.Response, error)

	/*
		GetIvr Get IVR

		**Required ACL:** `confd.ivr.{ivr_id}.read`

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param ivrId IVR's ID
		@return IvrAPIGetIvrRequest
	*/
	GetIvr(ctx context.Context, ivrId int32) IvrAPIGetIvrRequest

	// GetIvrExecute executes the request
	//  @return Ivr
	GetIvrExecute(r IvrAPIGetIvrRequest) (*Ivr, *http.Response, error)

	/*
		ListIvr List IVR

		**Required ACL:** `confd.ivr.read`

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@return IvrAPIListIvrRequest
	*/
	ListIvr(ctx context.Context) IvrAPIListIvrRequest

	// ListIvrExecute executes the request
	//  @return IvrItems
	ListIvrExecute(r IvrAPIListIvrRequest) (*IvrItems, *http.Response, error)

	/*
		UpdateIvr Update IVR

		**Required ACL:** `confd.ivr.{ivr_id}.update`

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param ivrId IVR's ID
		@return IvrAPIUpdateIvrRequest
	*/
	UpdateIvr(ctx context.Context, ivrId int32) IvrAPIUpdateIvrRequest

	// UpdateIvrExecute executes the request
	UpdateIvrExecute(r IvrAPIUpdateIvrRequest) (*http.Response, error)
}

// IvrAPIService IvrAPI service
type IvrAPIService service

type IvrAPICreateIvrRequest struct {
	ctx          context.Context
	ApiService   IvrAPI
	body         *Ivr
	accentTenant *string
}

// IVR to create
func (r IvrAPICreateIvrRequest) Body(body Ivr) IvrAPICreateIvrRequest {
	r.body = &body
	return r
}

// The tenant&#39;s UUID, defining the ownership of a given resource.
func (r IvrAPICreateIvrRequest) AccentTenant(accentTenant string) IvrAPICreateIvrRequest {
	r.accentTenant = &accentTenant
	return r
}

func (r IvrAPICreateIvrRequest) Execute() (*Ivr, *http.Response, error) {
	return r.ApiService.CreateIvrExecute(r)
}

/*
CreateIvr Create IVR

**Required ACL:** `confd.ivr.create`

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return IvrAPICreateIvrRequest
*/
func (a *IvrAPIService) CreateIvr(ctx context.Context) IvrAPICreateIvrRequest {
	return IvrAPICreateIvrRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return Ivr
func (a *IvrAPIService) CreateIvrExecute(r IvrAPICreateIvrRequest) (*Ivr, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *Ivr
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "IvrAPIService.CreateIvr")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/ivr"

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
	if r.accentTenant != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "Accent-Tenant", r.accentTenant, "")
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

type IvrAPIDeleteIvrRequest struct {
	ctx          context.Context
	ApiService   IvrAPI
	ivrId        int32
	accentTenant *string
}

// The tenant&#39;s UUID, defining the ownership of a given resource.
func (r IvrAPIDeleteIvrRequest) AccentTenant(accentTenant string) IvrAPIDeleteIvrRequest {
	r.accentTenant = &accentTenant
	return r
}

func (r IvrAPIDeleteIvrRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteIvrExecute(r)
}

/*
DeleteIvr Delete IVR

**Required ACL:** `confd.ivr.{ivr_id}.delete`

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param ivrId IVR's ID
	@return IvrAPIDeleteIvrRequest
*/
func (a *IvrAPIService) DeleteIvr(ctx context.Context, ivrId int32) IvrAPIDeleteIvrRequest {
	return IvrAPIDeleteIvrRequest{
		ApiService: a,
		ctx:        ctx,
		ivrId:      ivrId,
	}
}

// Execute executes the request
func (a *IvrAPIService) DeleteIvrExecute(r IvrAPIDeleteIvrRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "IvrAPIService.DeleteIvr")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/ivr/{ivr_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"ivr_id"+"}", url.PathEscape(parameterValueToString(r.ivrId, "ivrId")), -1)

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

type IvrAPIGetIvrRequest struct {
	ctx          context.Context
	ApiService   IvrAPI
	ivrId        int32
	accentTenant *string
}

// The tenant&#39;s UUID, defining the ownership of a given resource.
func (r IvrAPIGetIvrRequest) AccentTenant(accentTenant string) IvrAPIGetIvrRequest {
	r.accentTenant = &accentTenant
	return r
}

func (r IvrAPIGetIvrRequest) Execute() (*Ivr, *http.Response, error) {
	return r.ApiService.GetIvrExecute(r)
}

/*
GetIvr Get IVR

**Required ACL:** `confd.ivr.{ivr_id}.read`

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param ivrId IVR's ID
	@return IvrAPIGetIvrRequest
*/
func (a *IvrAPIService) GetIvr(ctx context.Context, ivrId int32) IvrAPIGetIvrRequest {
	return IvrAPIGetIvrRequest{
		ApiService: a,
		ctx:        ctx,
		ivrId:      ivrId,
	}
}

// Execute executes the request
//
//	@return Ivr
func (a *IvrAPIService) GetIvrExecute(r IvrAPIGetIvrRequest) (*Ivr, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *Ivr
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "IvrAPIService.GetIvr")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/ivr/{ivr_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"ivr_id"+"}", url.PathEscape(parameterValueToString(r.ivrId, "ivrId")), -1)

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

type IvrAPIListIvrRequest struct {
	ctx          context.Context
	ApiService   IvrAPI
	accentTenant *string
	recurse      *bool
	order        *string
	direction    *string
	limit        *int32
	offset       *int32
	search       *string
}

// The tenant&#39;s UUID, defining the ownership of a given resource.
func (r IvrAPIListIvrRequest) AccentTenant(accentTenant string) IvrAPIListIvrRequest {
	r.accentTenant = &accentTenant
	return r
}

// Should the query include sub-tenants
func (r IvrAPIListIvrRequest) Recurse(recurse bool) IvrAPIListIvrRequest {
	r.recurse = &recurse
	return r
}

// Name of the field to use for sorting the list of items returned.
func (r IvrAPIListIvrRequest) Order(order string) IvrAPIListIvrRequest {
	r.order = &order
	return r
}

// Sort list of items in &#39;asc&#39; (ascending) or &#39;desc&#39; (descending) order
func (r IvrAPIListIvrRequest) Direction(direction string) IvrAPIListIvrRequest {
	r.direction = &direction
	return r
}

// Maximum number of items to return in the list
func (r IvrAPIListIvrRequest) Limit(limit int32) IvrAPIListIvrRequest {
	r.limit = &limit
	return r
}

// Number of items to skip over in the list. Useful for pagination.
func (r IvrAPIListIvrRequest) Offset(offset int32) IvrAPIListIvrRequest {
	r.offset = &offset
	return r
}

// Search term for filtering a list of items. Only items with a field containing the search term will be returned.
func (r IvrAPIListIvrRequest) Search(search string) IvrAPIListIvrRequest {
	r.search = &search
	return r
}

func (r IvrAPIListIvrRequest) Execute() (*IvrItems, *http.Response, error) {
	return r.ApiService.ListIvrExecute(r)
}

/*
ListIvr List IVR

**Required ACL:** `confd.ivr.read`

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return IvrAPIListIvrRequest
*/
func (a *IvrAPIService) ListIvr(ctx context.Context) IvrAPIListIvrRequest {
	return IvrAPIListIvrRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return IvrItems
func (a *IvrAPIService) ListIvrExecute(r IvrAPIListIvrRequest) (*IvrItems, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *IvrItems
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "IvrAPIService.ListIvr")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/ivr"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.recurse != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "recurse", r.recurse, "")
	} else {
		var defaultValue bool = false
		r.recurse = &defaultValue
	}
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

type IvrAPIUpdateIvrRequest struct {
	ctx          context.Context
	ApiService   IvrAPI
	body         *Ivr
	ivrId        int32
	accentTenant *string
}

func (r IvrAPIUpdateIvrRequest) Body(body Ivr) IvrAPIUpdateIvrRequest {
	r.body = &body
	return r
}

// The tenant&#39;s UUID, defining the ownership of a given resource.
func (r IvrAPIUpdateIvrRequest) AccentTenant(accentTenant string) IvrAPIUpdateIvrRequest {
	r.accentTenant = &accentTenant
	return r
}

func (r IvrAPIUpdateIvrRequest) Execute() (*http.Response, error) {
	return r.ApiService.UpdateIvrExecute(r)
}

/*
UpdateIvr Update IVR

**Required ACL:** `confd.ivr.{ivr_id}.update`

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param ivrId IVR's ID
	@return IvrAPIUpdateIvrRequest
*/
func (a *IvrAPIService) UpdateIvr(ctx context.Context, ivrId int32) IvrAPIUpdateIvrRequest {
	return IvrAPIUpdateIvrRequest{
		ApiService: a,
		ctx:        ctx,
		ivrId:      ivrId,
	}
}

// Execute executes the request
func (a *IvrAPIService) UpdateIvrExecute(r IvrAPIUpdateIvrRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodPut
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "IvrAPIService.UpdateIvr")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/ivr/{ivr_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"ivr_id"+"}", url.PathEscape(parameterValueToString(r.ivrId, "ivrId")), -1)

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
	if r.accentTenant != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "Accent-Tenant", r.accentTenant, "")
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
