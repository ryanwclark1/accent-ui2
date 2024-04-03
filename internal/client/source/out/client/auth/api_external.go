/*
accent-auth

Accent's authentication service

API version: 0.1
Contact: help@accentvoice.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package auth

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)

type ExternalAPI interface {

	/*
		DeleteExternalAuthConfig Delete the client id and client secret

		**Required ACL**: `auth.{auth_type}.external.config.delete`

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param authType External auth type name
		@return ExternalAPIDeleteExternalAuthConfigRequest
	*/
	DeleteExternalAuthConfig(ctx context.Context, authType string) ExternalAPIDeleteExternalAuthConfigRequest

	// DeleteExternalAuthConfigExecute executes the request
	DeleteExternalAuthConfigExecute(r ExternalAPIDeleteExternalAuthConfigRequest) (*http.Response, error)

	/*
		GetExternalAuthConfig Retrieve the client id and client secret

		**Required ACL**: `auth.{auth_type}.external.config.read`

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param authType External auth type name
		@return ExternalAPIGetExternalAuthConfigRequest
	*/
	GetExternalAuthConfig(ctx context.Context, authType string) ExternalAPIGetExternalAuthConfigRequest

	// GetExternalAuthConfigExecute executes the request
	//  @return ExternalConfig
	GetExternalAuthConfigExecute(r ExternalAPIGetExternalAuthConfigRequest) (*ExternalConfig, *http.Response, error)

	/*
		GetExternalAuthUsers Retrieves the list of connected users to this external source

		**Required ACL**: `auth.{auth_type}.external.users`

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param authType External auth type name
		@return ExternalAPIGetExternalAuthUsersRequest
	*/
	GetExternalAuthUsers(ctx context.Context, authType string) ExternalAPIGetExternalAuthUsersRequest

	// GetExternalAuthUsersExecute executes the request
	//  @return ExternalAuthUserList
	GetExternalAuthUsersExecute(r ExternalAPIGetExternalAuthUsersRequest) (*ExternalAuthUserList, *http.Response, error)

	/*
		GetUserExternalAuth Retrieves the list of the users external auth data

		**Required ACL**: `auth.users.{user_uuid}.external.read` This list should not contain any sensible information

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param userUuid The UUID of the user
		@return ExternalAPIGetUserExternalAuthRequest
	*/
	GetUserExternalAuth(ctx context.Context, userUuid string) ExternalAPIGetUserExternalAuthRequest

	// GetUserExternalAuthExecute executes the request
	//  @return ExternalAuthList
	GetUserExternalAuthExecute(r ExternalAPIGetUserExternalAuthRequest) (*ExternalAuthList, *http.Response, error)

	/*
		PostExternalAuthConfig Add configuration for the given auth_type

		**Required ACL**: `auth.{auth_type}.external.config.create`

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param authType External auth type name
		@return ExternalAPIPostExternalAuthConfigRequest
	*/
	PostExternalAuthConfig(ctx context.Context, authType string) ExternalAPIPostExternalAuthConfigRequest

	// PostExternalAuthConfigExecute executes the request
	PostExternalAuthConfigExecute(r ExternalAPIPostExternalAuthConfigRequest) (*http.Response, error)

	/*
		UpdateExternalAuthConfig Update configuration for the given auth_type

		**Required ACL**: `auth.{auth_type}.external.config.edit`

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param authType External auth type name
		@return ExternalAPIUpdateExternalAuthConfigRequest
	*/
	UpdateExternalAuthConfig(ctx context.Context, authType string) ExternalAPIUpdateExternalAuthConfigRequest

	// UpdateExternalAuthConfigExecute executes the request
	UpdateExternalAuthConfigExecute(r ExternalAPIUpdateExternalAuthConfigRequest) (*http.Response, error)
}

// ExternalAPIService ExternalAPI service
type ExternalAPIService service

type ExternalAPIDeleteExternalAuthConfigRequest struct {
	ctx          context.Context
	ApiService   ExternalAPI
	authType     string
	accentTenant *string
}

// The tenant&#39;s UUID, defining the ownership of a given resource.
func (r ExternalAPIDeleteExternalAuthConfigRequest) AccentTenant(accentTenant string) ExternalAPIDeleteExternalAuthConfigRequest {
	r.accentTenant = &accentTenant
	return r
}

func (r ExternalAPIDeleteExternalAuthConfigRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteExternalAuthConfigExecute(r)
}

/*
DeleteExternalAuthConfig Delete the client id and client secret

**Required ACL**: `auth.{auth_type}.external.config.delete`

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param authType External auth type name
	@return ExternalAPIDeleteExternalAuthConfigRequest
*/
func (a *ExternalAPIService) DeleteExternalAuthConfig(ctx context.Context, authType string) ExternalAPIDeleteExternalAuthConfigRequest {
	return ExternalAPIDeleteExternalAuthConfigRequest{
		ApiService: a,
		ctx:        ctx,
		authType:   authType,
	}
}

// Execute executes the request
func (a *ExternalAPIService) DeleteExternalAuthConfigExecute(r ExternalAPIDeleteExternalAuthConfigRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ExternalAPIService.DeleteExternalAuthConfig")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/external/{auth_type}/config"
	localVarPath = strings.Replace(localVarPath, "{"+"auth_type"+"}", url.PathEscape(parameterValueToString(r.authType, "authType")), -1)

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
		if localVarHTTPResponse.StatusCode == 401 {
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
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ExternalAPIGetExternalAuthConfigRequest struct {
	ctx          context.Context
	ApiService   ExternalAPI
	authType     string
	accentTenant *string
}

// The tenant&#39;s UUID, defining the ownership of a given resource.
func (r ExternalAPIGetExternalAuthConfigRequest) AccentTenant(accentTenant string) ExternalAPIGetExternalAuthConfigRequest {
	r.accentTenant = &accentTenant
	return r
}

func (r ExternalAPIGetExternalAuthConfigRequest) Execute() (*ExternalConfig, *http.Response, error) {
	return r.ApiService.GetExternalAuthConfigExecute(r)
}

/*
GetExternalAuthConfig Retrieve the client id and client secret

**Required ACL**: `auth.{auth_type}.external.config.read`

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param authType External auth type name
	@return ExternalAPIGetExternalAuthConfigRequest
*/
func (a *ExternalAPIService) GetExternalAuthConfig(ctx context.Context, authType string) ExternalAPIGetExternalAuthConfigRequest {
	return ExternalAPIGetExternalAuthConfigRequest{
		ApiService: a,
		ctx:        ctx,
		authType:   authType,
	}
}

// Execute executes the request
//
//	@return ExternalConfig
func (a *ExternalAPIService) GetExternalAuthConfigExecute(r ExternalAPIGetExternalAuthConfigRequest) (*ExternalConfig, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ExternalConfig
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ExternalAPIService.GetExternalAuthConfig")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/external/{auth_type}/config"
	localVarPath = strings.Replace(localVarPath, "{"+"auth_type"+"}", url.PathEscape(parameterValueToString(r.authType, "authType")), -1)

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
		if localVarHTTPResponse.StatusCode == 401 {
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
		if localVarHTTPResponse.StatusCode == 404 {
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

type ExternalAPIGetExternalAuthUsersRequest struct {
	ctx          context.Context
	ApiService   ExternalAPI
	authType     string
	accentTenant *string
	recurse      *bool
	limit        *int32
	offset       *int32
}

// The tenant&#39;s UUID, defining the ownership of a given resource.
func (r ExternalAPIGetExternalAuthUsersRequest) AccentTenant(accentTenant string) ExternalAPIGetExternalAuthUsersRequest {
	r.accentTenant = &accentTenant
	return r
}

// Should the query include sub-tenants
func (r ExternalAPIGetExternalAuthUsersRequest) Recurse(recurse bool) ExternalAPIGetExternalAuthUsersRequest {
	r.recurse = &recurse
	return r
}

// The limit defines the number of individual objects that are returned
func (r ExternalAPIGetExternalAuthUsersRequest) Limit(limit int32) ExternalAPIGetExternalAuthUsersRequest {
	r.limit = &limit
	return r
}

// The offset defines the offsets the start by the number specified
func (r ExternalAPIGetExternalAuthUsersRequest) Offset(offset int32) ExternalAPIGetExternalAuthUsersRequest {
	r.offset = &offset
	return r
}

func (r ExternalAPIGetExternalAuthUsersRequest) Execute() (*ExternalAuthUserList, *http.Response, error) {
	return r.ApiService.GetExternalAuthUsersExecute(r)
}

/*
GetExternalAuthUsers Retrieves the list of connected users to this external source

**Required ACL**: `auth.{auth_type}.external.users`

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param authType External auth type name
	@return ExternalAPIGetExternalAuthUsersRequest
*/
func (a *ExternalAPIService) GetExternalAuthUsers(ctx context.Context, authType string) ExternalAPIGetExternalAuthUsersRequest {
	return ExternalAPIGetExternalAuthUsersRequest{
		ApiService: a,
		ctx:        ctx,
		authType:   authType,
	}
}

// Execute executes the request
//
//	@return ExternalAuthUserList
func (a *ExternalAPIService) GetExternalAuthUsersExecute(r ExternalAPIGetExternalAuthUsersRequest) (*ExternalAuthUserList, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ExternalAuthUserList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ExternalAPIService.GetExternalAuthUsers")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/external/{auth_type}/users"
	localVarPath = strings.Replace(localVarPath, "{"+"auth_type"+"}", url.PathEscape(parameterValueToString(r.authType, "authType")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.recurse != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "recurse", r.recurse, "")
	} else {
		var defaultValue bool = false
		r.recurse = &defaultValue
	}
	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "")
	}
	if r.offset != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "offset", r.offset, "")
	} else {
		var defaultValue int32 = 0
		r.offset = &defaultValue
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
		if localVarHTTPResponse.StatusCode == 401 {
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
		if localVarHTTPResponse.StatusCode == 404 {
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

type ExternalAPIGetUserExternalAuthRequest struct {
	ctx        context.Context
	ApiService ExternalAPI
	userUuid   string
	order      *string
	direction  *string
	limit      *int32
	offset     *int32
	search     *string
}

// Name of the field to use for sorting the list of items returned.
func (r ExternalAPIGetUserExternalAuthRequest) Order(order string) ExternalAPIGetUserExternalAuthRequest {
	r.order = &order
	return r
}

// Sort list of items in &#39;asc&#39; (ascending) or &#39;desc&#39; (descending) order
func (r ExternalAPIGetUserExternalAuthRequest) Direction(direction string) ExternalAPIGetUserExternalAuthRequest {
	r.direction = &direction
	return r
}

// The limit defines the number of individual objects that are returned
func (r ExternalAPIGetUserExternalAuthRequest) Limit(limit int32) ExternalAPIGetUserExternalAuthRequest {
	r.limit = &limit
	return r
}

// The offset defines the offsets the start by the number specified
func (r ExternalAPIGetUserExternalAuthRequest) Offset(offset int32) ExternalAPIGetUserExternalAuthRequest {
	r.offset = &offset
	return r
}

// Search term for filtering a list of items. Only items with a field containing the search term will be returned.
func (r ExternalAPIGetUserExternalAuthRequest) Search(search string) ExternalAPIGetUserExternalAuthRequest {
	r.search = &search
	return r
}

func (r ExternalAPIGetUserExternalAuthRequest) Execute() (*ExternalAuthList, *http.Response, error) {
	return r.ApiService.GetUserExternalAuthExecute(r)
}

/*
GetUserExternalAuth Retrieves the list of the users external auth data

**Required ACL**: `auth.users.{user_uuid}.external.read` This list should not contain any sensible information

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param userUuid The UUID of the user
	@return ExternalAPIGetUserExternalAuthRequest
*/
func (a *ExternalAPIService) GetUserExternalAuth(ctx context.Context, userUuid string) ExternalAPIGetUserExternalAuthRequest {
	return ExternalAPIGetUserExternalAuthRequest{
		ApiService: a,
		ctx:        ctx,
		userUuid:   userUuid,
	}
}

// Execute executes the request
//
//	@return ExternalAuthList
func (a *ExternalAPIService) GetUserExternalAuthExecute(r ExternalAPIGetUserExternalAuthRequest) (*ExternalAuthList, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ExternalAuthList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ExternalAPIService.GetUserExternalAuth")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/users/{user_uuid}/external"
	localVarPath = strings.Replace(localVarPath, "{"+"user_uuid"+"}", url.PathEscape(parameterValueToString(r.userUuid, "userUuid")), -1)

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
	} else {
		var defaultValue int32 = 0
		r.offset = &defaultValue
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

type ExternalAPIPostExternalAuthConfigRequest struct {
	ctx          context.Context
	ApiService   ExternalAPI
	config       *ExternalConfig
	authType     string
	accentTenant *string
}

// JSON object holding configuration for the given authentication type
func (r ExternalAPIPostExternalAuthConfigRequest) Config(config ExternalConfig) ExternalAPIPostExternalAuthConfigRequest {
	r.config = &config
	return r
}

// The tenant&#39;s UUID, defining the ownership of a given resource.
func (r ExternalAPIPostExternalAuthConfigRequest) AccentTenant(accentTenant string) ExternalAPIPostExternalAuthConfigRequest {
	r.accentTenant = &accentTenant
	return r
}

func (r ExternalAPIPostExternalAuthConfigRequest) Execute() (*http.Response, error) {
	return r.ApiService.PostExternalAuthConfigExecute(r)
}

/*
PostExternalAuthConfig Add configuration for the given auth_type

**Required ACL**: `auth.{auth_type}.external.config.create`

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param authType External auth type name
	@return ExternalAPIPostExternalAuthConfigRequest
*/
func (a *ExternalAPIService) PostExternalAuthConfig(ctx context.Context, authType string) ExternalAPIPostExternalAuthConfigRequest {
	return ExternalAPIPostExternalAuthConfigRequest{
		ApiService: a,
		ctx:        ctx,
		authType:   authType,
	}
}

// Execute executes the request
func (a *ExternalAPIService) PostExternalAuthConfigExecute(r ExternalAPIPostExternalAuthConfigRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodPost
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ExternalAPIService.PostExternalAuthConfig")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/external/{auth_type}/config"
	localVarPath = strings.Replace(localVarPath, "{"+"auth_type"+"}", url.PathEscape(parameterValueToString(r.authType, "authType")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.config == nil {
		return nil, reportError("config is required and must be specified")
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
	localVarPostBody = r.config
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
		if localVarHTTPResponse.StatusCode == 401 {
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
		if localVarHTTPResponse.StatusCode == 409 {
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

type ExternalAPIUpdateExternalAuthConfigRequest struct {
	ctx          context.Context
	ApiService   ExternalAPI
	config       *ExternalConfig
	authType     string
	accentTenant *string
}

// JSON object holding configuration for the given authentication type
func (r ExternalAPIUpdateExternalAuthConfigRequest) Config(config ExternalConfig) ExternalAPIUpdateExternalAuthConfigRequest {
	r.config = &config
	return r
}

// The tenant&#39;s UUID, defining the ownership of a given resource.
func (r ExternalAPIUpdateExternalAuthConfigRequest) AccentTenant(accentTenant string) ExternalAPIUpdateExternalAuthConfigRequest {
	r.accentTenant = &accentTenant
	return r
}

func (r ExternalAPIUpdateExternalAuthConfigRequest) Execute() (*http.Response, error) {
	return r.ApiService.UpdateExternalAuthConfigExecute(r)
}

/*
UpdateExternalAuthConfig Update configuration for the given auth_type

**Required ACL**: `auth.{auth_type}.external.config.edit`

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param authType External auth type name
	@return ExternalAPIUpdateExternalAuthConfigRequest
*/
func (a *ExternalAPIService) UpdateExternalAuthConfig(ctx context.Context, authType string) ExternalAPIUpdateExternalAuthConfigRequest {
	return ExternalAPIUpdateExternalAuthConfigRequest{
		ApiService: a,
		ctx:        ctx,
		authType:   authType,
	}
}

// Execute executes the request
func (a *ExternalAPIService) UpdateExternalAuthConfigExecute(r ExternalAPIUpdateExternalAuthConfigRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodPut
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ExternalAPIService.UpdateExternalAuthConfig")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/external/{auth_type}/config"
	localVarPath = strings.Replace(localVarPath, "{"+"auth_type"+"}", url.PathEscape(parameterValueToString(r.authType, "authType")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.config == nil {
		return nil, reportError("config is required and must be specified")
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
	localVarPostBody = r.config
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
		if localVarHTTPResponse.StatusCode == 401 {
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
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}
