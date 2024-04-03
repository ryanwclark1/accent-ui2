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
)

type ConfigAPI interface {

	/*
		GetConfig Show the current configuration

		**Required ACL:** `auth.config.read`

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@return ConfigAPIGetConfigRequest
	*/
	GetConfig(ctx context.Context) ConfigAPIGetConfigRequest

	// GetConfigExecute executes the request
	GetConfigExecute(r ConfigAPIGetConfigRequest) (*http.Response, error)

	/*
		PatchConfig Update the current configuration.

		**Required ACL:** `auth.config.update` Changes are not persistent across service restart.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@return ConfigAPIPatchConfigRequest
	*/
	PatchConfig(ctx context.Context) ConfigAPIPatchConfigRequest

	// PatchConfigExecute executes the request
	PatchConfigExecute(r ConfigAPIPatchConfigRequest) (*http.Response, error)
}

// ConfigAPIService ConfigAPI service
type ConfigAPIService service

type ConfigAPIGetConfigRequest struct {
	ctx        context.Context
	ApiService ConfigAPI
}

func (r ConfigAPIGetConfigRequest) Execute() (*http.Response, error) {
	return r.ApiService.GetConfigExecute(r)
}

/*
GetConfig Show the current configuration

**Required ACL:** `auth.config.read`

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ConfigAPIGetConfigRequest
*/
func (a *ConfigAPIService) GetConfig(ctx context.Context) ConfigAPIGetConfigRequest {
	return ConfigAPIGetConfigRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
func (a *ConfigAPIService) GetConfigExecute(r ConfigAPIGetConfigRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodGet
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ConfigAPIService.GetConfig")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/config"

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

type ConfigAPIPatchConfigRequest struct {
	ctx         context.Context
	ApiService  ConfigAPI
	configPatch *[]ConfigPatchItem
}

// See https://en.wikipedia.org/wiki/JSON_Patch.
func (r ConfigAPIPatchConfigRequest) ConfigPatch(configPatch []ConfigPatchItem) ConfigAPIPatchConfigRequest {
	r.configPatch = &configPatch
	return r
}

func (r ConfigAPIPatchConfigRequest) Execute() (*http.Response, error) {
	return r.ApiService.PatchConfigExecute(r)
}

/*
PatchConfig Update the current configuration.

**Required ACL:** `auth.config.update` Changes are not persistent across service restart.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ConfigAPIPatchConfigRequest
*/
func (a *ConfigAPIService) PatchConfig(ctx context.Context) ConfigAPIPatchConfigRequest {
	return ConfigAPIPatchConfigRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
func (a *ConfigAPIService) PatchConfigExecute(r ConfigAPIPatchConfigRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodPatch
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ConfigAPIService.PatchConfig")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/config"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.configPatch == nil {
		return nil, reportError("configPatch is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

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
	// body params
	localVarPostBody = r.configPatch
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
