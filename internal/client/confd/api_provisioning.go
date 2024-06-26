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
)

type ProvisioningAPI interface {

	/*
		GetProvisioningNetworking Get Provisioning Networking configuration

		**Required ACL:** `confd.provisioning.networking.read`

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@return ProvisioningAPIGetProvisioningNetworkingRequest
	*/
	GetProvisioningNetworking(ctx context.Context) ProvisioningAPIGetProvisioningNetworkingRequest

	// GetProvisioningNetworkingExecute executes the request
	//  @return ProvisioningNetworking
	GetProvisioningNetworkingExecute(r ProvisioningAPIGetProvisioningNetworkingRequest) (*ProvisioningNetworking, *http.Response, error)

	/*
		UpdateProvisioningNetworking Update Provisioning Networking configuration

		**Required ACL:** `confd.provisioning.networking.update`

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@return ProvisioningAPIUpdateProvisioningNetworkingRequest
	*/
	UpdateProvisioningNetworking(ctx context.Context) ProvisioningAPIUpdateProvisioningNetworkingRequest

	// UpdateProvisioningNetworkingExecute executes the request
	UpdateProvisioningNetworkingExecute(r ProvisioningAPIUpdateProvisioningNetworkingRequest) (*http.Response, error)
}

// ProvisioningAPIService ProvisioningAPI service
type ProvisioningAPIService service

type ProvisioningAPIGetProvisioningNetworkingRequest struct {
	ctx        context.Context
	ApiService ProvisioningAPI
}

func (r ProvisioningAPIGetProvisioningNetworkingRequest) Execute() (*ProvisioningNetworking, *http.Response, error) {
	return r.ApiService.GetProvisioningNetworkingExecute(r)
}

/*
GetProvisioningNetworking Get Provisioning Networking configuration

**Required ACL:** `confd.provisioning.networking.read`

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ProvisioningAPIGetProvisioningNetworkingRequest
*/
func (a *ProvisioningAPIService) GetProvisioningNetworking(ctx context.Context) ProvisioningAPIGetProvisioningNetworkingRequest {
	return ProvisioningAPIGetProvisioningNetworkingRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ProvisioningNetworking
func (a *ProvisioningAPIService) GetProvisioningNetworkingExecute(r ProvisioningAPIGetProvisioningNetworkingRequest) (*ProvisioningNetworking, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ProvisioningNetworking
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ProvisioningAPIService.GetProvisioningNetworking")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/provisioning/networking"

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

type ProvisioningAPIUpdateProvisioningNetworkingRequest struct {
	ctx        context.Context
	ApiService ProvisioningAPI
	body       *ProvisioningNetworking
}

func (r ProvisioningAPIUpdateProvisioningNetworkingRequest) Body(body ProvisioningNetworking) ProvisioningAPIUpdateProvisioningNetworkingRequest {
	r.body = &body
	return r
}

func (r ProvisioningAPIUpdateProvisioningNetworkingRequest) Execute() (*http.Response, error) {
	return r.ApiService.UpdateProvisioningNetworkingExecute(r)
}

/*
UpdateProvisioningNetworking Update Provisioning Networking configuration

**Required ACL:** `confd.provisioning.networking.update`

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ProvisioningAPIUpdateProvisioningNetworkingRequest
*/
func (a *ProvisioningAPIService) UpdateProvisioningNetworking(ctx context.Context) ProvisioningAPIUpdateProvisioningNetworkingRequest {
	return ProvisioningAPIUpdateProvisioningNetworkingRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
func (a *ProvisioningAPIService) UpdateProvisioningNetworkingExecute(r ProvisioningAPIUpdateProvisioningNetworkingRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodPut
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ProvisioningAPIService.UpdateProvisioningNetworking")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/provisioning/networking"

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
	localVarHTTPHeaderAccepts := []string{}

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
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}
