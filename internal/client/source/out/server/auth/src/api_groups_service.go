/*
 * accent-auth
 *
 * Accent's authentication service
 *
 * API version: 0.1
 * Contact: help@accentvoice.io
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package authserver

import (
	"context"
	"errors"
	"net/http"
)

// GroupsAPIService is a service that implements the logic for the GroupsAPIServicer
// This service should implement the business logic for every endpoint for the GroupsAPI API.
// Include any external packages or services that will be required by this service.
type GroupsAPIService struct {
}

// NewGroupsAPIService creates a default api service
func NewGroupsAPIService() GroupsAPIServicer {
	return &GroupsAPIService{}
}

// AddGroupPolicy - Associate a group to a policy
func (s *GroupsAPIService) AddGroupPolicy(ctx context.Context, groupUuid string, policyUuid string) (ImplResponse, error) {
	// TODO - update AddGroupPolicy with the required logic for this service method.
	// Add api_groups_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(204, {}) or use other options such as http.Ok ...
	// return Response(204, nil),nil

	// TODO: Uncomment the next line to return response Response(404, Error{}) or use other options such as http.Ok ...
	// return Response(404, Error{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("AddGroupPolicy method not implemented")
}

// AddGroupUser - Associate a group to a user
func (s *GroupsAPIService) AddGroupUser(ctx context.Context, groupUuid string, userUuid string) (ImplResponse, error) {
	// TODO - update AddGroupUser with the required logic for this service method.
	// Add api_groups_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(204, {}) or use other options such as http.Ok ...
	// return Response(204, nil),nil

	// TODO: Uncomment the next line to return response Response(404, Error{}) or use other options such as http.Ok ...
	// return Response(404, Error{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("AddGroupUser method not implemented")
}

// DeleteGroup - Delete a group
func (s *GroupsAPIService) DeleteGroup(ctx context.Context, groupUuid string, accentTenant string) (ImplResponse, error) {
	// TODO - update DeleteGroup with the required logic for this service method.
	// Add api_groups_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(204, {}) or use other options such as http.Ok ...
	// return Response(204, nil),nil

	// TODO: Uncomment the next line to return response Response(401, Error{}) or use other options such as http.Ok ...
	// return Response(401, Error{}), nil

	// TODO: Uncomment the next line to return response Response(404, Error{}) or use other options such as http.Ok ...
	// return Response(404, Error{}), nil

	// TODO: Uncomment the next line to return response Response(500, Error{}) or use other options such as http.Ok ...
	// return Response(500, Error{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("DeleteGroup method not implemented")
}

// DeleteGroupPolicy - Dissociate a policy from a group
func (s *GroupsAPIService) DeleteGroupPolicy(ctx context.Context, groupUuid string, policyUuid string) (ImplResponse, error) {
	// TODO - update DeleteGroupPolicy with the required logic for this service method.
	// Add api_groups_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(204, {}) or use other options such as http.Ok ...
	// return Response(204, nil),nil

	// TODO: Uncomment the next line to return response Response(404, Error{}) or use other options such as http.Ok ...
	// return Response(404, Error{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("DeleteGroupPolicy method not implemented")
}

// EditGroup - Modify a group
func (s *GroupsAPIService) EditGroup(ctx context.Context, groupUuid string, body GroupPut, accentTenant string) (ImplResponse, error) {
	// TODO - update EditGroup with the required logic for this service method.
	// Add api_groups_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(200, GroupResult{}) or use other options such as http.Ok ...
	// return Response(200, GroupResult{}), nil

	// TODO: Uncomment the next line to return response Response(401, Error{}) or use other options such as http.Ok ...
	// return Response(401, Error{}), nil

	// TODO: Uncomment the next line to return response Response(409, Error{}) or use other options such as http.Ok ...
	// return Response(409, Error{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("EditGroup method not implemented")
}

// GetGroup - Retrieves the details of a group
func (s *GroupsAPIService) GetGroup(ctx context.Context, groupUuid string, accentTenant string) (ImplResponse, error) {
	// TODO - update GetGroup with the required logic for this service method.
	// Add api_groups_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(200, GroupResult{}) or use other options such as http.Ok ...
	// return Response(200, GroupResult{}), nil

	// TODO: Uncomment the next line to return response Response(401, Error{}) or use other options such as http.Ok ...
	// return Response(401, Error{}), nil

	// TODO: Uncomment the next line to return response Response(404, Error{}) or use other options such as http.Ok ...
	// return Response(404, Error{}), nil

	// TODO: Uncomment the next line to return response Response(500, Error{}) or use other options such as http.Ok ...
	// return Response(500, Error{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("GetGroup method not implemented")
}

// GetGroupPolicies - Retrieves the list of policies associated to a group
func (s *GroupsAPIService) GetGroupPolicies(ctx context.Context, groupUuid string, order string, direction string, limit int32, offset int32, search string) (ImplResponse, error) {
	// TODO - update GetGroupPolicies with the required logic for this service method.
	// Add api_groups_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(200, GetPoliciesResult{}) or use other options such as http.Ok ...
	// return Response(200, GetPoliciesResult{}), nil

	// TODO: Uncomment the next line to return response Response(401, Error{}) or use other options such as http.Ok ...
	// return Response(401, Error{}), nil

	// TODO: Uncomment the next line to return response Response(404, Error{}) or use other options such as http.Ok ...
	// return Response(404, Error{}), nil

	// TODO: Uncomment the next line to return response Response(500, Error{}) or use other options such as http.Ok ...
	// return Response(500, Error{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("GetGroupPolicies method not implemented")
}

// GetGroupUsers - Retrieves the list of users associated to a group
func (s *GroupsAPIService) GetGroupUsers(ctx context.Context, groupUuid string, order string, direction string, limit int32, offset int32, search string) (ImplResponse, error) {
	// TODO - update GetGroupUsers with the required logic for this service method.
	// Add api_groups_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(200, UserList{}) or use other options such as http.Ok ...
	// return Response(200, UserList{}), nil

	// TODO: Uncomment the next line to return response Response(401, Error{}) or use other options such as http.Ok ...
	// return Response(401, Error{}), nil

	// TODO: Uncomment the next line to return response Response(404, Error{}) or use other options such as http.Ok ...
	// return Response(404, Error{}), nil

	// TODO: Uncomment the next line to return response Response(500, Error{}) or use other options such as http.Ok ...
	// return Response(500, Error{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("GetGroupUsers method not implemented")
}

// RemoveGroupUser - Dissociate a user from a group
func (s *GroupsAPIService) RemoveGroupUser(ctx context.Context, groupUuid string, userUuid string) (ImplResponse, error) {
	// TODO - update RemoveGroupUser with the required logic for this service method.
	// Add api_groups_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(204, {}) or use other options such as http.Ok ...
	// return Response(204, nil),nil

	// TODO: Uncomment the next line to return response Response(404, Error{}) or use other options such as http.Ok ...
	// return Response(404, Error{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("RemoveGroupUser method not implemented")
}
