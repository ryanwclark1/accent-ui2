/*
accent-plugind

Accent's plugin management service

API version: 0.2
Contact: help@accentvoice.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package plugind

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)

type PluginAPI interface {

	/*
		GetMarket List plugins available on the configured market

		**Required ACL:** `plugind.market.read` Allow the administrator to get a list of available plugins

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@return PluginAPIGetMarketRequest
	*/
	GetMarket(ctx context.Context) PluginAPIGetMarketRequest

	// GetMarketExecute executes the request
	//  @return GetMarketResult
	GetMarketExecute(r PluginAPIGetMarketRequest) (*GetMarketResult, *http.Response, error)

	/*
		GetMarketPlugin Fetch the information about a plugin from the market

		**Required ACL:** `plugind.market.read` Allow the administrator to view a plugins information from the market. ---

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param namespace The plugin's namespace
		@param name The plugin's name
		@return PluginAPIGetMarketPluginRequest
	*/
	GetMarketPlugin(ctx context.Context, namespace string, name string) PluginAPIGetMarketPluginRequest

	// GetMarketPluginExecute executes the request
	//  @return MarketPluginList
	GetMarketPluginExecute(r PluginAPIGetMarketPluginRequest) (*MarketPluginList, *http.Response, error)

	/*
		GetPlugin Fetch the information about a plugin that has been installed

		**Required ACL:** `plugind.plugins.{namespace}.{name}.read` Allow the administrator to view a plugins metadata file. ---

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param namespace The plugin's namespace
		@param name The plugin's name
		@return PluginAPIGetPluginRequest
	*/
	GetPlugin(ctx context.Context, namespace string, name string) PluginAPIGetPluginRequest

	// GetPluginExecute executes the request
	//  @return PluginMetadata
	GetPluginExecute(r PluginAPIGetPluginRequest) (*PluginMetadata, *http.Response, error)

	/*
		GetPlugins List installed plugins

		**Required ACL:** `plugind.plugins.read` Allow the administrator to get a list of all installed plugins

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@return PluginAPIGetPluginsRequest
	*/
	GetPlugins(ctx context.Context) PluginAPIGetPluginsRequest

	// GetPluginsExecute executes the request
	//  @return GetPluginsResult
	GetPluginsExecute(r PluginAPIGetPluginsRequest) (*GetPluginsResult, *http.Response, error)

	/*
		InstallPlugin Install a plugin

		**Required ACL:** `plugind.plugins.create` Allow the administrator to install a plugin on the server.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@return PluginAPIInstallPluginRequest
	*/
	InstallPlugin(ctx context.Context) PluginAPIInstallPluginRequest

	// InstallPluginExecute executes the request
	//  @return InstallResponse
	InstallPluginExecute(r PluginAPIInstallPluginRequest) (*InstallResponse, *http.Response, error)

	/*
		UninstallPlugin Uninstall a plugin

		**Required ACL:** `plugind.plugins.{namespace}.{name}.delete` Allow the administrator to uninstall a plugin. ---

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param namespace The plugin's namespace
		@param name The plugin's name
		@return PluginAPIUninstallPluginRequest
	*/
	UninstallPlugin(ctx context.Context, namespace string, name string) PluginAPIUninstallPluginRequest

	// UninstallPluginExecute executes the request
	UninstallPluginExecute(r PluginAPIUninstallPluginRequest) (*http.Response, error)
}

// PluginAPIService PluginAPI service
type PluginAPIService service

type PluginAPIGetMarketRequest struct {
	ctx        context.Context
	ApiService PluginAPI
	limit      *int32
	offset     *int32
	order      *string
	direction  *string
	search     *string
	namespace  *string
	name       *string
	installed  *bool
}

// Maximum number of items to return in the list
func (r PluginAPIGetMarketRequest) Limit(limit int32) PluginAPIGetMarketRequest {
	r.limit = &limit
	return r
}

// Number of items to skip over in the list. Useful for pagination.
func (r PluginAPIGetMarketRequest) Offset(offset int32) PluginAPIGetMarketRequest {
	r.offset = &offset
	return r
}

// Name of the field to use for sorting the list of items returned.
func (r PluginAPIGetMarketRequest) Order(order string) PluginAPIGetMarketRequest {
	r.order = &order
	return r
}

// Sort list of items in &#39;asc&#39; (ascending) or &#39;desc&#39; (descending) order
func (r PluginAPIGetMarketRequest) Direction(direction string) PluginAPIGetMarketRequest {
	r.direction = &direction
	return r
}

// Filter list of items, this search is applied to all fields and is not strict (eba matches Sébastien)
func (r PluginAPIGetMarketRequest) Search(search string) PluginAPIGetMarketRequest {
	r.search = &search
	return r
}

// Search by namespace
func (r PluginAPIGetMarketRequest) Namespace(namespace string) PluginAPIGetMarketRequest {
	r.namespace = &namespace
	return r
}

// Search by name
func (r PluginAPIGetMarketRequest) Name(name string) PluginAPIGetMarketRequest {
	r.name = &name
	return r
}

// Filter installed plugins
func (r PluginAPIGetMarketRequest) Installed(installed bool) PluginAPIGetMarketRequest {
	r.installed = &installed
	return r
}

func (r PluginAPIGetMarketRequest) Execute() (*GetMarketResult, *http.Response, error) {
	return r.ApiService.GetMarketExecute(r)
}

/*
GetMarket List plugins available on the configured market

**Required ACL:** `plugind.market.read` Allow the administrator to get a list of available plugins

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return PluginAPIGetMarketRequest
*/
func (a *PluginAPIService) GetMarket(ctx context.Context) PluginAPIGetMarketRequest {
	return PluginAPIGetMarketRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GetMarketResult
func (a *PluginAPIService) GetMarketExecute(r PluginAPIGetMarketRequest) (*GetMarketResult, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GetMarketResult
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PluginAPIService.GetMarket")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/market"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "")
	}
	if r.offset != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "offset", r.offset, "")
	}
	if r.order != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "order", r.order, "")
	}
	if r.direction != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "direction", r.direction, "")
	}
	if r.search != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "search", r.search, "")
	}
	if r.namespace != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "namespace", r.namespace, "")
	}
	if r.name != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "name", r.name, "")
	}
	if r.installed != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "installed", r.installed, "")
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

type PluginAPIGetMarketPluginRequest struct {
	ctx        context.Context
	ApiService PluginAPI
	namespace  string
	name       string
}

func (r PluginAPIGetMarketPluginRequest) Execute() (*MarketPluginList, *http.Response, error) {
	return r.ApiService.GetMarketPluginExecute(r)
}

/*
GetMarketPlugin Fetch the information about a plugin from the market

**Required ACL:** `plugind.market.read` Allow the administrator to view a plugins information from the market. ---

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param namespace The plugin's namespace
	@param name The plugin's name
	@return PluginAPIGetMarketPluginRequest
*/
func (a *PluginAPIService) GetMarketPlugin(ctx context.Context, namespace string, name string) PluginAPIGetMarketPluginRequest {
	return PluginAPIGetMarketPluginRequest{
		ApiService: a,
		ctx:        ctx,
		namespace:  namespace,
		name:       name,
	}
}

// Execute executes the request
//
//	@return MarketPluginList
func (a *PluginAPIService) GetMarketPluginExecute(r PluginAPIGetMarketPluginRequest) (*MarketPluginList, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *MarketPluginList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PluginAPIService.GetMarketPlugin")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/market/{namespace}/{name}"
	localVarPath = strings.Replace(localVarPath, "{"+"namespace"+"}", url.PathEscape(parameterValueToString(r.namespace, "namespace")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", url.PathEscape(parameterValueToString(r.name, "name")), -1)

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

type PluginAPIGetPluginRequest struct {
	ctx        context.Context
	ApiService PluginAPI
	namespace  string
	name       string
}

func (r PluginAPIGetPluginRequest) Execute() (*PluginMetadata, *http.Response, error) {
	return r.ApiService.GetPluginExecute(r)
}

/*
GetPlugin Fetch the information about a plugin that has been installed

**Required ACL:** `plugind.plugins.{namespace}.{name}.read` Allow the administrator to view a plugins metadata file. ---

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param namespace The plugin's namespace
	@param name The plugin's name
	@return PluginAPIGetPluginRequest
*/
func (a *PluginAPIService) GetPlugin(ctx context.Context, namespace string, name string) PluginAPIGetPluginRequest {
	return PluginAPIGetPluginRequest{
		ApiService: a,
		ctx:        ctx,
		namespace:  namespace,
		name:       name,
	}
}

// Execute executes the request
//
//	@return PluginMetadata
func (a *PluginAPIService) GetPluginExecute(r PluginAPIGetPluginRequest) (*PluginMetadata, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *PluginMetadata
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PluginAPIService.GetPlugin")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/plugins/{namespace}/{name}"
	localVarPath = strings.Replace(localVarPath, "{"+"namespace"+"}", url.PathEscape(parameterValueToString(r.namespace, "namespace")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", url.PathEscape(parameterValueToString(r.name, "name")), -1)

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

type PluginAPIGetPluginsRequest struct {
	ctx        context.Context
	ApiService PluginAPI
}

func (r PluginAPIGetPluginsRequest) Execute() (*GetPluginsResult, *http.Response, error) {
	return r.ApiService.GetPluginsExecute(r)
}

/*
GetPlugins List installed plugins

**Required ACL:** `plugind.plugins.read` Allow the administrator to get a list of all installed plugins

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return PluginAPIGetPluginsRequest
*/
func (a *PluginAPIService) GetPlugins(ctx context.Context) PluginAPIGetPluginsRequest {
	return PluginAPIGetPluginsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GetPluginsResult
func (a *PluginAPIService) GetPluginsExecute(r PluginAPIGetPluginsRequest) (*GetPluginsResult, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GetPluginsResult
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PluginAPIService.GetPlugins")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/plugins"

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

type PluginAPIInstallPluginRequest struct {
	ctx        context.Context
	ApiService PluginAPI
	body       *PluginInstallParameters
	reinstall  *bool
}

// The plugins&#39; installation parameters
func (r PluginAPIInstallPluginRequest) Body(body PluginInstallParameters) PluginAPIInstallPluginRequest {
	r.body = &body
	return r
}

// With this option the plugin will be reinstalled if it is already installed
func (r PluginAPIInstallPluginRequest) Reinstall(reinstall bool) PluginAPIInstallPluginRequest {
	r.reinstall = &reinstall
	return r
}

func (r PluginAPIInstallPluginRequest) Execute() (*InstallResponse, *http.Response, error) {
	return r.ApiService.InstallPluginExecute(r)
}

/*
InstallPlugin Install a plugin

**Required ACL:** `plugind.plugins.create` Allow the administrator to install a plugin on the server.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return PluginAPIInstallPluginRequest
*/
func (a *PluginAPIService) InstallPlugin(ctx context.Context) PluginAPIInstallPluginRequest {
	return PluginAPIInstallPluginRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return InstallResponse
func (a *PluginAPIService) InstallPluginExecute(r PluginAPIInstallPluginRequest) (*InstallResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *InstallResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PluginAPIService.InstallPlugin")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/plugins"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.body == nil {
		return localVarReturnValue, nil, reportError("body is required and must be specified")
	}

	if r.reinstall != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "reinstall", r.reinstall, "")
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

type PluginAPIUninstallPluginRequest struct {
	ctx        context.Context
	ApiService PluginAPI
	namespace  string
	name       string
}

func (r PluginAPIUninstallPluginRequest) Execute() (*http.Response, error) {
	return r.ApiService.UninstallPluginExecute(r)
}

/*
UninstallPlugin Uninstall a plugin

**Required ACL:** `plugind.plugins.{namespace}.{name}.delete` Allow the administrator to uninstall a plugin. ---

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param namespace The plugin's namespace
	@param name The plugin's name
	@return PluginAPIUninstallPluginRequest
*/
func (a *PluginAPIService) UninstallPlugin(ctx context.Context, namespace string, name string) PluginAPIUninstallPluginRequest {
	return PluginAPIUninstallPluginRequest{
		ApiService: a,
		ctx:        ctx,
		namespace:  namespace,
		name:       name,
	}
}

// Execute executes the request
func (a *PluginAPIService) UninstallPluginExecute(r PluginAPIUninstallPluginRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PluginAPIService.UninstallPlugin")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/plugins/{namespace}/{name}"
	localVarPath = strings.Replace(localVarPath, "{"+"namespace"+"}", url.PathEscape(parameterValueToString(r.namespace, "namespace")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", url.PathEscape(parameterValueToString(r.name, "name")), -1)

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
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}
