/*
accent-provd

Provisioning application REST API

API version: 0.2
Contact: help@accentvoice.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package provd

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)

type ConfigsAPI interface {

	/*
		DeleteCfgMgrConfigsConfigId Delete a configuration

		**Required ACL:** `provd.cfg_mgr.configs.{config_id}.delete`

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param configId Configuration ID
		@return ApiDeleteCfgMgrConfigsConfigIdRequest
	*/
	DeleteCfgMgrConfigsConfigId(ctx context.Context, configId string) ApiDeleteCfgMgrConfigsConfigIdRequest

	// DeleteCfgMgrConfigsConfigIdExecute executes the request
	DeleteCfgMgrConfigsConfigIdExecute(r ApiDeleteCfgMgrConfigsConfigIdRequest) (*http.Response, error)

	/*
		GetCfgMgr Get the Config Manager resource

		**Required ACL:** `provd.cfg_mgr.read` The configuration manager resource represents the entry point to the accent-provd configuration REST API

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@return ApiGetCfgMgrRequest
	*/
	GetCfgMgr(ctx context.Context) ApiGetCfgMgrRequest

	// GetCfgMgrExecute executes the request
	//  @return LinksObject
	GetCfgMgrExecute(r ApiGetCfgMgrRequest) (*LinksObject, *http.Response, error)

	/*
		GetCfgMgrConfig Get a configuration

		**Required ACL:** `provd.cfg_mgr.configs.{config_id}.read`

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param configId Configuration ID
		@return ApiGetCfgMgrConfigRequest
	*/
	GetCfgMgrConfig(ctx context.Context, configId string) ApiGetCfgMgrConfigRequest

	// GetCfgMgrConfigExecute executes the request
	//  @return ConfigObject
	GetCfgMgrConfigExecute(r ApiGetCfgMgrConfigRequest) (*ConfigObject, *http.Response, error)

	/*
		GetCfgMgrConfigs List and find configurations

		**Required ACL:** `provd.cfg_mgr.configs.read`

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@return ApiGetCfgMgrConfigsRequest
	*/
	GetCfgMgrConfigs(ctx context.Context) ApiGetCfgMgrConfigsRequest

	// GetCfgMgrConfigsExecute executes the request
	//  @return ConfigsObject
	GetCfgMgrConfigsExecute(r ApiGetCfgMgrConfigsRequest) (*ConfigsObject, *http.Response, error)

	/*
		GetCfgMgrRawConfig Get a raw configuration

		**Required ACL:** `provd.cfg_mgr.configs.{config_id}.raw.read`

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param configId Configuration ID
		@return ApiGetCfgMgrRawConfigRequest
	*/
	GetCfgMgrRawConfig(ctx context.Context, configId string) ApiGetCfgMgrRawConfigRequest

	// GetCfgMgrRawConfigExecute executes the request
	//  @return RawConfigurationObject
	GetCfgMgrRawConfigExecute(r ApiGetCfgMgrRawConfigRequest) (*RawConfigurationObject, *http.Response, error)

	/*
		PostCfgMgrAutocreate Create an autocreate configuration

		**Required ACL:** `provd.cfg_mgr.autocreate.create` Create a new config based on the config that has the autocreate role

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@return ApiPostCfgMgrAutocreateRequest
	*/
	PostCfgMgrAutocreate(ctx context.Context) ApiPostCfgMgrAutocreateRequest

	// PostCfgMgrAutocreateExecute executes the request
	//  @return IdObject
	PostCfgMgrAutocreateExecute(r ApiPostCfgMgrAutocreateRequest) (*IdObject, *http.Response, error)

	/*
		PostCfgMgrConfigs Create a configuration

		**Required ACL:** `provd.cfg_mgr.configs.create`

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@return ApiPostCfgMgrConfigsRequest
	*/
	PostCfgMgrConfigs(ctx context.Context) ApiPostCfgMgrConfigsRequest

	// PostCfgMgrConfigsExecute executes the request
	//  @return IdObject
	PostCfgMgrConfigsExecute(r ApiPostCfgMgrConfigsRequest) (*IdObject, *http.Response, error)

	/*
		PutCfgMgrConfig Update a configuration

		**Required ACL:** `provd.cfg_mgr.configs.{config_id}.update`

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param configId Configuration ID
		@return ApiPutCfgMgrConfigRequest
	*/
	PutCfgMgrConfig(ctx context.Context, configId string) ApiPutCfgMgrConfigRequest

	// PutCfgMgrConfigExecute executes the request
	PutCfgMgrConfigExecute(r ApiPutCfgMgrConfigRequest) (*http.Response, error)
}

// ConfigsAPIService ConfigsAPI service
type ConfigsAPIService service

type ApiDeleteCfgMgrConfigsConfigIdRequest struct {
	ctx        context.Context
	ApiService ConfigsAPI
	configId   string
}

func (r ApiDeleteCfgMgrConfigsConfigIdRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteCfgMgrConfigsConfigIdExecute(r)
}

/*
DeleteCfgMgrConfigsConfigId Delete a configuration

**Required ACL:** `provd.cfg_mgr.configs.{config_id}.delete`

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param configId Configuration ID
	@return ApiDeleteCfgMgrConfigsConfigIdRequest
*/
func (a *ConfigsAPIService) DeleteCfgMgrConfigsConfigId(ctx context.Context, configId string) ApiDeleteCfgMgrConfigsConfigIdRequest {
	return ApiDeleteCfgMgrConfigsConfigIdRequest{
		ApiService: a,
		ctx:        ctx,
		configId:   configId,
	}
}

// Execute executes the request
func (a *ConfigsAPIService) DeleteCfgMgrConfigsConfigIdExecute(r ApiDeleteCfgMgrConfigsConfigIdRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ConfigsAPIService.DeleteCfgMgrConfigsConfigId")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/cfg_mgr/configs/{config_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"config_id"+"}", url.PathEscape(parameterValueToString(r.configId, "configId")), -1)

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
	localVarHTTPHeaderAccepts := []string{"application/vnd.accent.provd+json"}

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
		if localVarHTTPResponse.StatusCode == 403 {
			var v string
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
			var v string
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

type ApiGetCfgMgrRequest struct {
	ctx        context.Context
	ApiService ConfigsAPI
}

func (r ApiGetCfgMgrRequest) Execute() (*LinksObject, *http.Response, error) {
	return r.ApiService.GetCfgMgrExecute(r)
}

/*
GetCfgMgr Get the Config Manager resource

**Required ACL:** `provd.cfg_mgr.read` The configuration manager resource represents the entry point to the accent-provd configuration REST API

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiGetCfgMgrRequest
*/
func (a *ConfigsAPIService) GetCfgMgr(ctx context.Context) ApiGetCfgMgrRequest {
	return ApiGetCfgMgrRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return LinksObject
func (a *ConfigsAPIService) GetCfgMgrExecute(r ApiGetCfgMgrRequest) (*LinksObject, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *LinksObject
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ConfigsAPIService.GetCfgMgr")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/cfg_mgr"

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
	localVarHTTPHeaderAccepts := []string{"application/vnd.accent.provd+json", "links"}

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

type ApiGetCfgMgrConfigRequest struct {
	ctx        context.Context
	ApiService ConfigsAPI
	configId   string
}

func (r ApiGetCfgMgrConfigRequest) Execute() (*ConfigObject, *http.Response, error) {
	return r.ApiService.GetCfgMgrConfigExecute(r)
}

/*
GetCfgMgrConfig Get a configuration

**Required ACL:** `provd.cfg_mgr.configs.{config_id}.read`

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param configId Configuration ID
	@return ApiGetCfgMgrConfigRequest
*/
func (a *ConfigsAPIService) GetCfgMgrConfig(ctx context.Context, configId string) ApiGetCfgMgrConfigRequest {
	return ApiGetCfgMgrConfigRequest{
		ApiService: a,
		ctx:        ctx,
		configId:   configId,
	}
}

// Execute executes the request
//
//	@return ConfigObject
func (a *ConfigsAPIService) GetCfgMgrConfigExecute(r ApiGetCfgMgrConfigRequest) (*ConfigObject, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ConfigObject
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ConfigsAPIService.GetCfgMgrConfig")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/cfg_mgr/configs/{config_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"config_id"+"}", url.PathEscape(parameterValueToString(r.configId, "configId")), -1)

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
	localVarHTTPHeaderAccepts := []string{"application/vnd.accent.provd+json"}

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
			var v string
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

type ApiGetCfgMgrConfigsRequest struct {
	ctx        context.Context
	ApiService ConfigsAPI
	q          *string
	fields     *string
	skip       *int32
	sort       *string
	sortOrd    *string
}

// A selector, encoded in JSON, describing which entries should be returned. All entries are returned if not specified.  Example: &#x60;{\&quot;ip\&quot;:\&quot;10.34.1.110\&quot;}&#x60;
func (r ApiGetCfgMgrConfigsRequest) Q(q string) ApiGetCfgMgrConfigsRequest {
	r.q = &q
	return r
}

// A list of fields, separated by comma.  Example: &#x60;mac,ip&#x60;
func (r ApiGetCfgMgrConfigsRequest) Fields(fields string) ApiGetCfgMgrConfigsRequest {
	r.fields = &fields
	return r
}

// An integer specifing the number of entries to skip.  Example: 10
func (r ApiGetCfgMgrConfigsRequest) Skip(skip int32) ApiGetCfgMgrConfigsRequest {
	r.skip = &skip
	return r
}

// The key on which to sort the results.  Example: &#x60;id&#x60;
func (r ApiGetCfgMgrConfigsRequest) Sort(sort string) ApiGetCfgMgrConfigsRequest {
	r.sort = &sort
	return r
}

// The order of sort
func (r ApiGetCfgMgrConfigsRequest) SortOrd(sortOrd string) ApiGetCfgMgrConfigsRequest {
	r.sortOrd = &sortOrd
	return r
}

func (r ApiGetCfgMgrConfigsRequest) Execute() (*ConfigsObject, *http.Response, error) {
	return r.ApiService.GetCfgMgrConfigsExecute(r)
}

/*
GetCfgMgrConfigs List and find configurations

**Required ACL:** `provd.cfg_mgr.configs.read`

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiGetCfgMgrConfigsRequest
*/
func (a *ConfigsAPIService) GetCfgMgrConfigs(ctx context.Context) ApiGetCfgMgrConfigsRequest {
	return ApiGetCfgMgrConfigsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ConfigsObject
func (a *ConfigsAPIService) GetCfgMgrConfigsExecute(r ApiGetCfgMgrConfigsRequest) (*ConfigsObject, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ConfigsObject
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ConfigsAPIService.GetCfgMgrConfigs")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/cfg_mgr/configs"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.q != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "q", r.q, "")
	}
	if r.fields != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "fields", r.fields, "")
	}
	if r.skip != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "skip", r.skip, "")
	}
	if r.sort != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "sort", r.sort, "")
	}
	if r.sortOrd != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "sort_ord", r.sortOrd, "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/vnd.accent.provd+json"}

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

type ApiGetCfgMgrRawConfigRequest struct {
	ctx        context.Context
	ApiService ConfigsAPI
	configId   string
}

func (r ApiGetCfgMgrRawConfigRequest) Execute() (*RawConfigurationObject, *http.Response, error) {
	return r.ApiService.GetCfgMgrRawConfigExecute(r)
}

/*
GetCfgMgrRawConfig Get a raw configuration

**Required ACL:** `provd.cfg_mgr.configs.{config_id}.raw.read`

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param configId Configuration ID
	@return ApiGetCfgMgrRawConfigRequest
*/
func (a *ConfigsAPIService) GetCfgMgrRawConfig(ctx context.Context, configId string) ApiGetCfgMgrRawConfigRequest {
	return ApiGetCfgMgrRawConfigRequest{
		ApiService: a,
		ctx:        ctx,
		configId:   configId,
	}
}

// Execute executes the request
//
//	@return RawConfigurationObject
func (a *ConfigsAPIService) GetCfgMgrRawConfigExecute(r ApiGetCfgMgrRawConfigRequest) (*RawConfigurationObject, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *RawConfigurationObject
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ConfigsAPIService.GetCfgMgrRawConfig")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/cfg_mgr/configs/{config_id}/raw"
	localVarPath = strings.Replace(localVarPath, "{"+"config_id"+"}", url.PathEscape(parameterValueToString(r.configId, "configId")), -1)

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
	localVarHTTPHeaderAccepts := []string{"application/vnd.accent.provd+json"}

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
			var v string
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

type ApiPostCfgMgrAutocreateRequest struct {
	ctx        context.Context
	ApiService ConfigsAPI
	body       *map[string]interface{}
}

// Empty object body
func (r ApiPostCfgMgrAutocreateRequest) Body(body map[string]interface{}) ApiPostCfgMgrAutocreateRequest {
	r.body = &body
	return r
}

func (r ApiPostCfgMgrAutocreateRequest) Execute() (*IdObject, *http.Response, error) {
	return r.ApiService.PostCfgMgrAutocreateExecute(r)
}

/*
PostCfgMgrAutocreate Create an autocreate configuration

**Required ACL:** `provd.cfg_mgr.autocreate.create` Create a new config based on the config that has the autocreate role

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiPostCfgMgrAutocreateRequest
*/
func (a *ConfigsAPIService) PostCfgMgrAutocreate(ctx context.Context) ApiPostCfgMgrAutocreateRequest {
	return ApiPostCfgMgrAutocreateRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return IdObject
func (a *ConfigsAPIService) PostCfgMgrAutocreateExecute(r ApiPostCfgMgrAutocreateRequest) (*IdObject, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *IdObject
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ConfigsAPIService.PostCfgMgrAutocreate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/cfg_mgr/autocreate"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/vnd.accent.provd+json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/vnd.accent.provd+json"}

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

type ApiPostCfgMgrConfigsRequest struct {
	ctx        context.Context
	ApiService ConfigsAPI
	body       *ConfigObject
}

// Body of a configuration parameter
func (r ApiPostCfgMgrConfigsRequest) Body(body ConfigObject) ApiPostCfgMgrConfigsRequest {
	r.body = &body
	return r
}

func (r ApiPostCfgMgrConfigsRequest) Execute() (*IdObject, *http.Response, error) {
	return r.ApiService.PostCfgMgrConfigsExecute(r)
}

/*
PostCfgMgrConfigs Create a configuration

**Required ACL:** `provd.cfg_mgr.configs.create`

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiPostCfgMgrConfigsRequest
*/
func (a *ConfigsAPIService) PostCfgMgrConfigs(ctx context.Context) ApiPostCfgMgrConfigsRequest {
	return ApiPostCfgMgrConfigsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return IdObject
func (a *ConfigsAPIService) PostCfgMgrConfigsExecute(r ApiPostCfgMgrConfigsRequest) (*IdObject, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *IdObject
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ConfigsAPIService.PostCfgMgrConfigs")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/cfg_mgr/configs"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/vnd.accent.provd+json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/vnd.accent.provd+json"}

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

type ApiPutCfgMgrConfigRequest struct {
	ctx        context.Context
	ApiService ConfigsAPI
	configId   string
	body       *ConfigObject
}

// Body of a configuration parameter
func (r ApiPutCfgMgrConfigRequest) Body(body ConfigObject) ApiPutCfgMgrConfigRequest {
	r.body = &body
	return r
}

func (r ApiPutCfgMgrConfigRequest) Execute() (*http.Response, error) {
	return r.ApiService.PutCfgMgrConfigExecute(r)
}

/*
PutCfgMgrConfig Update a configuration

**Required ACL:** `provd.cfg_mgr.configs.{config_id}.update`

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param configId Configuration ID
	@return ApiPutCfgMgrConfigRequest
*/
func (a *ConfigsAPIService) PutCfgMgrConfig(ctx context.Context, configId string) ApiPutCfgMgrConfigRequest {
	return ApiPutCfgMgrConfigRequest{
		ApiService: a,
		ctx:        ctx,
		configId:   configId,
	}
}

// Execute executes the request
func (a *ConfigsAPIService) PutCfgMgrConfigExecute(r ApiPutCfgMgrConfigRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodPut
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ConfigsAPIService.PutCfgMgrConfig")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/cfg_mgr/configs/{config_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"config_id"+"}", url.PathEscape(parameterValueToString(r.configId, "configId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/vnd.accent.provd+json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/vnd.accent.provd+json"}

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
		if localVarHTTPResponse.StatusCode == 404 {
			var v string
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v string
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
