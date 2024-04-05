/*
 * accent-confd
 *
 * Confd exposes an API for managing core resources on a Accent server such as users, extensions, devices, voicemails, queues, etc. Resources can be associated together to provide additional functionality. For example: By associating a voicemail with a user, calls will automatically fallback on to the voicemail when the user cannot answer.  Implementation notes ====================  Errors ------  Responses containing errors will have a status code in the 400 or 500 class. A list of error messages will be returned in the body of the response as a JSON-encoded array:  ~~~ [     \"Input error - User not found\",     \"Resource error - User not associated to a line\" ] ~~~   Updating resources via PUT --------------------------  When updating a resource, all fields become optional. In other words, only values that have been changed need to be sent to the server. Please note that this behavior may change in future versions of the API.
 *
 * API version: 1.1
 * Contact: help@accentvoice.io
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package confdserver

import (
	"context"
	"errors"
	"net/http"
)

// PagingsAPIService is a service that implements the logic for the PagingsAPIServicer
// This service should implement the business logic for every endpoint for the PagingsAPI API.
// Include any external packages or services that will be required by this service.
type PagingsAPIService struct {
}

// NewPagingsAPIService creates a default api service
func NewPagingsAPIService() PagingsAPIServicer {
	return &PagingsAPIService{}
}

// CreatePaging - Create paging
func (s *PagingsAPIService) CreatePaging(ctx context.Context, body Paging, accentTenant string) (ImplResponse, error) {
	// TODO - update CreatePaging with the required logic for this service method.
	// Add api_pagings_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(201, Paging{}) or use other options such as http.Ok ...
	// return Response(201, Paging{}), nil

	// TODO: Uncomment the next line to return response Response(400, []string{}) or use other options such as http.Ok ...
	// return Response(400, []string{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("CreatePaging method not implemented")
}

// DeletePaging - Delete paging
func (s *PagingsAPIService) DeletePaging(ctx context.Context, pagingId int32, accentTenant string) (ImplResponse, error) {
	// TODO - update DeletePaging with the required logic for this service method.
	// Add api_pagings_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(204, {}) or use other options such as http.Ok ...
	// return Response(204, nil),nil

	// TODO: Uncomment the next line to return response Response(400, []string{}) or use other options such as http.Ok ...
	// return Response(400, []string{}), nil

	// TODO: Uncomment the next line to return response Response(404, []string{}) or use other options such as http.Ok ...
	// return Response(404, []string{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("DeletePaging method not implemented")
}

// GetPaging - Get paging
func (s *PagingsAPIService) GetPaging(ctx context.Context, pagingId int32, accentTenant string) (ImplResponse, error) {
	// TODO - update GetPaging with the required logic for this service method.
	// Add api_pagings_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(200, Paging{}) or use other options such as http.Ok ...
	// return Response(200, Paging{}), nil

	// TODO: Uncomment the next line to return response Response(404, []string{}) or use other options such as http.Ok ...
	// return Response(404, []string{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("GetPaging method not implemented")
}

// ListPagings - List paging
func (s *PagingsAPIService) ListPagings(ctx context.Context, accentTenant string, recurse bool, order string, direction string, limit int32, offset int32, search string) (ImplResponse, error) {
	// TODO - update ListPagings with the required logic for this service method.
	// Add api_pagings_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(200, PagingItems{}) or use other options such as http.Ok ...
	// return Response(200, PagingItems{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("ListPagings method not implemented")
}

// UpdatePaging - Update paging
func (s *PagingsAPIService) UpdatePaging(ctx context.Context, pagingId int32, body Paging, accentTenant string) (ImplResponse, error) {
	// TODO - update UpdatePaging with the required logic for this service method.
	// Add api_pagings_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(204, {}) or use other options such as http.Ok ...
	// return Response(204, nil),nil

	// TODO: Uncomment the next line to return response Response(400, []string{}) or use other options such as http.Ok ...
	// return Response(400, []string{}), nil

	// TODO: Uncomment the next line to return response Response(404, []string{}) or use other options such as http.Ok ...
	// return Response(404, []string{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("UpdatePaging method not implemented")
}

// UpdatePagingCallerUsers - Update paging and callers
func (s *PagingsAPIService) UpdatePagingCallerUsers(ctx context.Context, pagingId int32, body UsersUuid) (ImplResponse, error) {
	// TODO - update UpdatePagingCallerUsers with the required logic for this service method.
	// Add api_pagings_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(204, {}) or use other options such as http.Ok ...
	// return Response(204, nil),nil

	// TODO: Uncomment the next line to return response Response(400, []string{}) or use other options such as http.Ok ...
	// return Response(400, []string{}), nil

	// TODO: Uncomment the next line to return response Response(404, []string{}) or use other options such as http.Ok ...
	// return Response(404, []string{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("UpdatePagingCallerUsers method not implemented")
}

// UpdatePagingMemberUsers - Update paging and members
func (s *PagingsAPIService) UpdatePagingMemberUsers(ctx context.Context, pagingId int32, body UsersUuid) (ImplResponse, error) {
	// TODO - update UpdatePagingMemberUsers with the required logic for this service method.
	// Add api_pagings_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(204, {}) or use other options such as http.Ok ...
	// return Response(204, nil),nil

	// TODO: Uncomment the next line to return response Response(400, []string{}) or use other options such as http.Ok ...
	// return Response(400, []string{}), nil

	// TODO: Uncomment the next line to return response Response(404, []string{}) or use other options such as http.Ok ...
	// return Response(404, []string{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("UpdatePagingMemberUsers method not implemented")
}