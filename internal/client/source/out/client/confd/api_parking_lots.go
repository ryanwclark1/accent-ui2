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

type ParkingLotsAPI interface {

	/*
		AssociateParkingLotExtension Associate parking_lot and extension

		**Required ACL:** `confd.parkinglots.{parking_lot_id}.extensions.{extension_id}.update`

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param parkingLotId Parking Lot's ID
		@param extensionId
		@return ApiAssociateParkingLotExtensionRequest
	*/
	AssociateParkingLotExtension(ctx context.Context, parkingLotId int32, extensionId int32) ApiAssociateParkingLotExtensionRequest

	// AssociateParkingLotExtensionExecute executes the request
	AssociateParkingLotExtensionExecute(r ApiAssociateParkingLotExtensionRequest) (*http.Response, error)

	/*
		CreateParkingLot Create parking lot

		**Required ACL:** `confd.parkinglots.create`

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@return ApiCreateParkingLotRequest
	*/
	CreateParkingLot(ctx context.Context) ApiCreateParkingLotRequest

	// CreateParkingLotExecute executes the request
	//  @return ParkingLot
	CreateParkingLotExecute(r ApiCreateParkingLotRequest) (*ParkingLot, *http.Response, error)

	/*
		DeleteParkingLot Delete parking lot

		**Required ACL:** `confd.parkinglots.{parking_lot_id}.delete`

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param parkingLotId Parking Lot's ID
		@return ApiDeleteParkingLotRequest
	*/
	DeleteParkingLot(ctx context.Context, parkingLotId int32) ApiDeleteParkingLotRequest

	// DeleteParkingLotExecute executes the request
	DeleteParkingLotExecute(r ApiDeleteParkingLotRequest) (*http.Response, error)

	/*
		DissociateParkingLotExtension Dissociate parking lot and extension

		**Required ACL:** `confd.parkinglots.{parking_lot_id}.extensions.{extension_id}.delete`

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param parkingLotId Parking Lot's ID
		@param extensionId
		@return ApiDissociateParkingLotExtensionRequest
	*/
	DissociateParkingLotExtension(ctx context.Context, parkingLotId int32, extensionId int32) ApiDissociateParkingLotExtensionRequest

	// DissociateParkingLotExtensionExecute executes the request
	DissociateParkingLotExtensionExecute(r ApiDissociateParkingLotExtensionRequest) (*http.Response, error)

	/*
		GetParkingLot Get parking lot

		**Required ACL:** `confd.parkinglots.{parking_lot_id}.read`

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param parkingLotId Parking Lot's ID
		@return ApiGetParkingLotRequest
	*/
	GetParkingLot(ctx context.Context, parkingLotId int32) ApiGetParkingLotRequest

	// GetParkingLotExecute executes the request
	//  @return ParkingLot
	GetParkingLotExecute(r ApiGetParkingLotRequest) (*ParkingLot, *http.Response, error)

	/*
		ListParkingLots List parking lots

		**Required ACL:** `confd.parkinglots.read`

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@return ApiListParkingLotsRequest
	*/
	ListParkingLots(ctx context.Context) ApiListParkingLotsRequest

	// ListParkingLotsExecute executes the request
	//  @return ParkingLotItems
	ListParkingLotsExecute(r ApiListParkingLotsRequest) (*ParkingLotItems, *http.Response, error)

	/*
		UpdateParkingLot Update parking lot

		**Required ACL:** `confd.parkinglots.{parking_lot_id}.update`

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param parkingLotId Parking Lot's ID
		@return ApiUpdateParkingLotRequest
	*/
	UpdateParkingLot(ctx context.Context, parkingLotId int32) ApiUpdateParkingLotRequest

	// UpdateParkingLotExecute executes the request
	UpdateParkingLotExecute(r ApiUpdateParkingLotRequest) (*http.Response, error)
}

// ParkingLotsAPIService ParkingLotsAPI service
type ParkingLotsAPIService service

type ApiAssociateParkingLotExtensionRequest struct {
	ctx          context.Context
	ApiService   ParkingLotsAPI
	parkingLotId int32
	extensionId  int32
}

func (r ApiAssociateParkingLotExtensionRequest) Execute() (*http.Response, error) {
	return r.ApiService.AssociateParkingLotExtensionExecute(r)
}

/*
AssociateParkingLotExtension Associate parking_lot and extension

**Required ACL:** `confd.parkinglots.{parking_lot_id}.extensions.{extension_id}.update`

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param parkingLotId Parking Lot's ID
	@param extensionId
	@return ApiAssociateParkingLotExtensionRequest
*/
func (a *ParkingLotsAPIService) AssociateParkingLotExtension(ctx context.Context, parkingLotId int32, extensionId int32) ApiAssociateParkingLotExtensionRequest {
	return ApiAssociateParkingLotExtensionRequest{
		ApiService:   a,
		ctx:          ctx,
		parkingLotId: parkingLotId,
		extensionId:  extensionId,
	}
}

// Execute executes the request
func (a *ParkingLotsAPIService) AssociateParkingLotExtensionExecute(r ApiAssociateParkingLotExtensionRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodPut
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ParkingLotsAPIService.AssociateParkingLotExtension")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/parkinglots/{parking_lot_id}/extensions/{extension_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"parking_lot_id"+"}", url.PathEscape(parameterValueToString(r.parkingLotId, "parkingLotId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"extension_id"+"}", url.PathEscape(parameterValueToString(r.extensionId, "extensionId")), -1)

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

type ApiCreateParkingLotRequest struct {
	ctx          context.Context
	ApiService   ParkingLotsAPI
	body         *ParkingLot
	accentTenant *string
}

// Parking lot to create
func (r ApiCreateParkingLotRequest) Body(body ParkingLot) ApiCreateParkingLotRequest {
	r.body = &body
	return r
}

// The tenant&#39;s UUID, defining the ownership of a given resource.
func (r ApiCreateParkingLotRequest) AccentTenant(accentTenant string) ApiCreateParkingLotRequest {
	r.accentTenant = &accentTenant
	return r
}

func (r ApiCreateParkingLotRequest) Execute() (*ParkingLot, *http.Response, error) {
	return r.ApiService.CreateParkingLotExecute(r)
}

/*
CreateParkingLot Create parking lot

**Required ACL:** `confd.parkinglots.create`

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiCreateParkingLotRequest
*/
func (a *ParkingLotsAPIService) CreateParkingLot(ctx context.Context) ApiCreateParkingLotRequest {
	return ApiCreateParkingLotRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ParkingLot
func (a *ParkingLotsAPIService) CreateParkingLotExecute(r ApiCreateParkingLotRequest) (*ParkingLot, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ParkingLot
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ParkingLotsAPIService.CreateParkingLot")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/parkinglots"

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

type ApiDeleteParkingLotRequest struct {
	ctx          context.Context
	ApiService   ParkingLotsAPI
	parkingLotId int32
	accentTenant *string
}

// The tenant&#39;s UUID, defining the ownership of a given resource.
func (r ApiDeleteParkingLotRequest) AccentTenant(accentTenant string) ApiDeleteParkingLotRequest {
	r.accentTenant = &accentTenant
	return r
}

func (r ApiDeleteParkingLotRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteParkingLotExecute(r)
}

/*
DeleteParkingLot Delete parking lot

**Required ACL:** `confd.parkinglots.{parking_lot_id}.delete`

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param parkingLotId Parking Lot's ID
	@return ApiDeleteParkingLotRequest
*/
func (a *ParkingLotsAPIService) DeleteParkingLot(ctx context.Context, parkingLotId int32) ApiDeleteParkingLotRequest {
	return ApiDeleteParkingLotRequest{
		ApiService:   a,
		ctx:          ctx,
		parkingLotId: parkingLotId,
	}
}

// Execute executes the request
func (a *ParkingLotsAPIService) DeleteParkingLotExecute(r ApiDeleteParkingLotRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ParkingLotsAPIService.DeleteParkingLot")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/parkinglots/{parking_lot_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"parking_lot_id"+"}", url.PathEscape(parameterValueToString(r.parkingLotId, "parkingLotId")), -1)

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

type ApiDissociateParkingLotExtensionRequest struct {
	ctx          context.Context
	ApiService   ParkingLotsAPI
	parkingLotId int32
	extensionId  int32
}

func (r ApiDissociateParkingLotExtensionRequest) Execute() (*http.Response, error) {
	return r.ApiService.DissociateParkingLotExtensionExecute(r)
}

/*
DissociateParkingLotExtension Dissociate parking lot and extension

**Required ACL:** `confd.parkinglots.{parking_lot_id}.extensions.{extension_id}.delete`

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param parkingLotId Parking Lot's ID
	@param extensionId
	@return ApiDissociateParkingLotExtensionRequest
*/
func (a *ParkingLotsAPIService) DissociateParkingLotExtension(ctx context.Context, parkingLotId int32, extensionId int32) ApiDissociateParkingLotExtensionRequest {
	return ApiDissociateParkingLotExtensionRequest{
		ApiService:   a,
		ctx:          ctx,
		parkingLotId: parkingLotId,
		extensionId:  extensionId,
	}
}

// Execute executes the request
func (a *ParkingLotsAPIService) DissociateParkingLotExtensionExecute(r ApiDissociateParkingLotExtensionRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ParkingLotsAPIService.DissociateParkingLotExtension")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/parkinglots/{parking_lot_id}/extensions/{extension_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"parking_lot_id"+"}", url.PathEscape(parameterValueToString(r.parkingLotId, "parkingLotId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"extension_id"+"}", url.PathEscape(parameterValueToString(r.extensionId, "extensionId")), -1)

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

type ApiGetParkingLotRequest struct {
	ctx          context.Context
	ApiService   ParkingLotsAPI
	parkingLotId int32
	accentTenant *string
}

// The tenant&#39;s UUID, defining the ownership of a given resource.
func (r ApiGetParkingLotRequest) AccentTenant(accentTenant string) ApiGetParkingLotRequest {
	r.accentTenant = &accentTenant
	return r
}

func (r ApiGetParkingLotRequest) Execute() (*ParkingLot, *http.Response, error) {
	return r.ApiService.GetParkingLotExecute(r)
}

/*
GetParkingLot Get parking lot

**Required ACL:** `confd.parkinglots.{parking_lot_id}.read`

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param parkingLotId Parking Lot's ID
	@return ApiGetParkingLotRequest
*/
func (a *ParkingLotsAPIService) GetParkingLot(ctx context.Context, parkingLotId int32) ApiGetParkingLotRequest {
	return ApiGetParkingLotRequest{
		ApiService:   a,
		ctx:          ctx,
		parkingLotId: parkingLotId,
	}
}

// Execute executes the request
//
//	@return ParkingLot
func (a *ParkingLotsAPIService) GetParkingLotExecute(r ApiGetParkingLotRequest) (*ParkingLot, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ParkingLot
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ParkingLotsAPIService.GetParkingLot")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/parkinglots/{parking_lot_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"parking_lot_id"+"}", url.PathEscape(parameterValueToString(r.parkingLotId, "parkingLotId")), -1)

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

type ApiListParkingLotsRequest struct {
	ctx          context.Context
	ApiService   ParkingLotsAPI
	accentTenant *string
	recurse      *bool
	order        *string
	direction    *string
	limit        *int32
	offset       *int32
	search       *string
}

// The tenant&#39;s UUID, defining the ownership of a given resource.
func (r ApiListParkingLotsRequest) AccentTenant(accentTenant string) ApiListParkingLotsRequest {
	r.accentTenant = &accentTenant
	return r
}

// Should the query include sub-tenants
func (r ApiListParkingLotsRequest) Recurse(recurse bool) ApiListParkingLotsRequest {
	r.recurse = &recurse
	return r
}

// Name of the field to use for sorting the list of items returned.
func (r ApiListParkingLotsRequest) Order(order string) ApiListParkingLotsRequest {
	r.order = &order
	return r
}

// Sort list of items in &#39;asc&#39; (ascending) or &#39;desc&#39; (descending) order
func (r ApiListParkingLotsRequest) Direction(direction string) ApiListParkingLotsRequest {
	r.direction = &direction
	return r
}

// Maximum number of items to return in the list
func (r ApiListParkingLotsRequest) Limit(limit int32) ApiListParkingLotsRequest {
	r.limit = &limit
	return r
}

// Number of items to skip over in the list. Useful for pagination.
func (r ApiListParkingLotsRequest) Offset(offset int32) ApiListParkingLotsRequest {
	r.offset = &offset
	return r
}

// Search term for filtering a list of items. Only items with a field containing the search term will be returned.
func (r ApiListParkingLotsRequest) Search(search string) ApiListParkingLotsRequest {
	r.search = &search
	return r
}

func (r ApiListParkingLotsRequest) Execute() (*ParkingLotItems, *http.Response, error) {
	return r.ApiService.ListParkingLotsExecute(r)
}

/*
ListParkingLots List parking lots

**Required ACL:** `confd.parkinglots.read`

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiListParkingLotsRequest
*/
func (a *ParkingLotsAPIService) ListParkingLots(ctx context.Context) ApiListParkingLotsRequest {
	return ApiListParkingLotsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ParkingLotItems
func (a *ParkingLotsAPIService) ListParkingLotsExecute(r ApiListParkingLotsRequest) (*ParkingLotItems, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ParkingLotItems
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ParkingLotsAPIService.ListParkingLots")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/parkinglots"

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

type ApiUpdateParkingLotRequest struct {
	ctx          context.Context
	ApiService   ParkingLotsAPI
	body         *ParkingLot
	parkingLotId int32
	accentTenant *string
}

func (r ApiUpdateParkingLotRequest) Body(body ParkingLot) ApiUpdateParkingLotRequest {
	r.body = &body
	return r
}

// The tenant&#39;s UUID, defining the ownership of a given resource.
func (r ApiUpdateParkingLotRequest) AccentTenant(accentTenant string) ApiUpdateParkingLotRequest {
	r.accentTenant = &accentTenant
	return r
}

func (r ApiUpdateParkingLotRequest) Execute() (*http.Response, error) {
	return r.ApiService.UpdateParkingLotExecute(r)
}

/*
UpdateParkingLot Update parking lot

**Required ACL:** `confd.parkinglots.{parking_lot_id}.update`

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param parkingLotId Parking Lot's ID
	@return ApiUpdateParkingLotRequest
*/
func (a *ParkingLotsAPIService) UpdateParkingLot(ctx context.Context, parkingLotId int32) ApiUpdateParkingLotRequest {
	return ApiUpdateParkingLotRequest{
		ApiService:   a,
		ctx:          ctx,
		parkingLotId: parkingLotId,
	}
}

// Execute executes the request
func (a *ParkingLotsAPIService) UpdateParkingLotExecute(r ApiUpdateParkingLotRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodPut
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ParkingLotsAPIService.UpdateParkingLot")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/parkinglots/{parking_lot_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"parking_lot_id"+"}", url.PathEscape(parameterValueToString(r.parkingLotId, "parkingLotId")), -1)

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
