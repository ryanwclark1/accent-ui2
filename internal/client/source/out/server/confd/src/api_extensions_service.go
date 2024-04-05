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

// ExtensionsAPIService is a service that implements the logic for the ExtensionsAPIServicer
// This service should implement the business logic for every endpoint for the ExtensionsAPI API.
// Include any external packages or services that will be required by this service.
type ExtensionsAPIService struct {
}

// NewExtensionsAPIService creates a default api service
func NewExtensionsAPIService() ExtensionsAPIServicer {
	return &ExtensionsAPIService{}
}

// AssociateConferenceExtension - Associate conference and extension
func (s *ExtensionsAPIService) AssociateConferenceExtension(ctx context.Context, conferenceId int32, extensionId int32) (ImplResponse, error) {
	// TODO - update AssociateConferenceExtension with the required logic for this service method.
	// Add api_extensions_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(204, {}) or use other options such as http.Ok ...
	// return Response(204, nil),nil

	// TODO: Uncomment the next line to return response Response(400, []string{}) or use other options such as http.Ok ...
	// return Response(400, []string{}), nil

	// TODO: Uncomment the next line to return response Response(404, []string{}) or use other options such as http.Ok ...
	// return Response(404, []string{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("AssociateConferenceExtension method not implemented")
}

// AssociateGroupExtension - Associate group and extension
func (s *ExtensionsAPIService) AssociateGroupExtension(ctx context.Context, groupUuid string, extensionId int32) (ImplResponse, error) {
	// TODO - update AssociateGroupExtension with the required logic for this service method.
	// Add api_extensions_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(204, {}) or use other options such as http.Ok ...
	// return Response(204, nil),nil

	// TODO: Uncomment the next line to return response Response(400, []string{}) or use other options such as http.Ok ...
	// return Response(400, []string{}), nil

	// TODO: Uncomment the next line to return response Response(404, []string{}) or use other options such as http.Ok ...
	// return Response(404, []string{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("AssociateGroupExtension method not implemented")
}

// AssociateIncallExtension - Associate incall and extension
func (s *ExtensionsAPIService) AssociateIncallExtension(ctx context.Context, incallId int32, extensionId int32) (ImplResponse, error) {
	// TODO - update AssociateIncallExtension with the required logic for this service method.
	// Add api_extensions_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(204, {}) or use other options such as http.Ok ...
	// return Response(204, nil),nil

	// TODO: Uncomment the next line to return response Response(400, []string{}) or use other options such as http.Ok ...
	// return Response(400, []string{}), nil

	// TODO: Uncomment the next line to return response Response(404, []string{}) or use other options such as http.Ok ...
	// return Response(404, []string{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("AssociateIncallExtension method not implemented")
}

// AssociateLineExtension - Associate line and extension
func (s *ExtensionsAPIService) AssociateLineExtension(ctx context.Context, lineId int32, extensionId int32) (ImplResponse, error) {
	// TODO - update AssociateLineExtension with the required logic for this service method.
	// Add api_extensions_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(204, {}) or use other options such as http.Ok ...
	// return Response(204, nil),nil

	// TODO: Uncomment the next line to return response Response(400, []string{}) or use other options such as http.Ok ...
	// return Response(400, []string{}), nil

	// TODO: Uncomment the next line to return response Response(404, []string{}) or use other options such as http.Ok ...
	// return Response(404, []string{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("AssociateLineExtension method not implemented")
}

// AssociateOutcallExtension - Associate outcall and extension
func (s *ExtensionsAPIService) AssociateOutcallExtension(ctx context.Context, outcallId int32, extensionId int32, body OutcallExtension) (ImplResponse, error) {
	// TODO - update AssociateOutcallExtension with the required logic for this service method.
	// Add api_extensions_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(204, {}) or use other options such as http.Ok ...
	// return Response(204, nil),nil

	// TODO: Uncomment the next line to return response Response(400, []string{}) or use other options such as http.Ok ...
	// return Response(400, []string{}), nil

	// TODO: Uncomment the next line to return response Response(404, []string{}) or use other options such as http.Ok ...
	// return Response(404, []string{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("AssociateOutcallExtension method not implemented")
}

// AssociateParkingLotExtension - Associate parking_lot and extension
func (s *ExtensionsAPIService) AssociateParkingLotExtension(ctx context.Context, parkingLotId int32, extensionId int32) (ImplResponse, error) {
	// TODO - update AssociateParkingLotExtension with the required logic for this service method.
	// Add api_extensions_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(204, {}) or use other options such as http.Ok ...
	// return Response(204, nil),nil

	// TODO: Uncomment the next line to return response Response(400, []string{}) or use other options such as http.Ok ...
	// return Response(400, []string{}), nil

	// TODO: Uncomment the next line to return response Response(404, []string{}) or use other options such as http.Ok ...
	// return Response(404, []string{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("AssociateParkingLotExtension method not implemented")
}

// AssociateQueueExtension - Associate queue and extension
func (s *ExtensionsAPIService) AssociateQueueExtension(ctx context.Context, queueId int32, extensionId int32, accentTenant string) (ImplResponse, error) {
	// TODO - update AssociateQueueExtension with the required logic for this service method.
	// Add api_extensions_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(204, {}) or use other options such as http.Ok ...
	// return Response(204, nil),nil

	// TODO: Uncomment the next line to return response Response(400, []string{}) or use other options such as http.Ok ...
	// return Response(400, []string{}), nil

	// TODO: Uncomment the next line to return response Response(404, []string{}) or use other options such as http.Ok ...
	// return Response(404, []string{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("AssociateQueueExtension method not implemented")
}

// CreateExtension - Create extension
func (s *ExtensionsAPIService) CreateExtension(ctx context.Context, body Extension, accentTenant string) (ImplResponse, error) {
	// TODO - update CreateExtension with the required logic for this service method.
	// Add api_extensions_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(201, Extension{}) or use other options such as http.Ok ...
	// return Response(201, Extension{}), nil

	// TODO: Uncomment the next line to return response Response(400, []string{}) or use other options such as http.Ok ...
	// return Response(400, []string{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("CreateExtension method not implemented")
}

// CreateLineExtension - Create extension
func (s *ExtensionsAPIService) CreateLineExtension(ctx context.Context, lineId int32, body Extension, accentTenant string) (ImplResponse, error) {
	// TODO - update CreateLineExtension with the required logic for this service method.
	// Add api_extensions_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(201, Extension{}) or use other options such as http.Ok ...
	// return Response(201, Extension{}), nil

	// TODO: Uncomment the next line to return response Response(400, []string{}) or use other options such as http.Ok ...
	// return Response(400, []string{}), nil

	// TODO: Uncomment the next line to return response Response(404, Object{}) or use other options such as http.Ok ...
	// return Response(404, Object{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("CreateLineExtension method not implemented")
}

// DeleteExtension - Delete extension
func (s *ExtensionsAPIService) DeleteExtension(ctx context.Context, extensionId int32) (ImplResponse, error) {
	// TODO - update DeleteExtension with the required logic for this service method.
	// Add api_extensions_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(204, {}) or use other options such as http.Ok ...
	// return Response(204, nil),nil

	// TODO: Uncomment the next line to return response Response(404, []string{}) or use other options such as http.Ok ...
	// return Response(404, []string{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("DeleteExtension method not implemented")
}

// DissociateConferenceExtension - Dissociate conference and extension
func (s *ExtensionsAPIService) DissociateConferenceExtension(ctx context.Context, conferenceId int32, extensionId int32) (ImplResponse, error) {
	// TODO - update DissociateConferenceExtension with the required logic for this service method.
	// Add api_extensions_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(204, {}) or use other options such as http.Ok ...
	// return Response(204, nil),nil

	// TODO: Uncomment the next line to return response Response(400, []string{}) or use other options such as http.Ok ...
	// return Response(400, []string{}), nil

	// TODO: Uncomment the next line to return response Response(404, []string{}) or use other options such as http.Ok ...
	// return Response(404, []string{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("DissociateConferenceExtension method not implemented")
}

// DissociateGroupExtension - Dissociate group and extension
func (s *ExtensionsAPIService) DissociateGroupExtension(ctx context.Context, groupUuid string, extensionId int32) (ImplResponse, error) {
	// TODO - update DissociateGroupExtension with the required logic for this service method.
	// Add api_extensions_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(204, {}) or use other options such as http.Ok ...
	// return Response(204, nil),nil

	// TODO: Uncomment the next line to return response Response(400, []string{}) or use other options such as http.Ok ...
	// return Response(400, []string{}), nil

	// TODO: Uncomment the next line to return response Response(404, []string{}) or use other options such as http.Ok ...
	// return Response(404, []string{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("DissociateGroupExtension method not implemented")
}

// DissociateIncallExtension - Dissociate incall and extension
func (s *ExtensionsAPIService) DissociateIncallExtension(ctx context.Context, incallId int32, extensionId int32) (ImplResponse, error) {
	// TODO - update DissociateIncallExtension with the required logic for this service method.
	// Add api_extensions_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(204, {}) or use other options such as http.Ok ...
	// return Response(204, nil),nil

	// TODO: Uncomment the next line to return response Response(400, []string{}) or use other options such as http.Ok ...
	// return Response(400, []string{}), nil

	// TODO: Uncomment the next line to return response Response(404, []string{}) or use other options such as http.Ok ...
	// return Response(404, []string{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("DissociateIncallExtension method not implemented")
}

// DissociateLineExtension - Dissociate line and extension
func (s *ExtensionsAPIService) DissociateLineExtension(ctx context.Context, lineId int32, extensionId int32) (ImplResponse, error) {
	// TODO - update DissociateLineExtension with the required logic for this service method.
	// Add api_extensions_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(204, {}) or use other options such as http.Ok ...
	// return Response(204, nil),nil

	// TODO: Uncomment the next line to return response Response(400, []string{}) or use other options such as http.Ok ...
	// return Response(400, []string{}), nil

	// TODO: Uncomment the next line to return response Response(404, []string{}) or use other options such as http.Ok ...
	// return Response(404, []string{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("DissociateLineExtension method not implemented")
}

// DissociateOutcallExtension - Dissociate outcall and extension
func (s *ExtensionsAPIService) DissociateOutcallExtension(ctx context.Context, outcallId int32, extensionId int32) (ImplResponse, error) {
	// TODO - update DissociateOutcallExtension with the required logic for this service method.
	// Add api_extensions_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(204, {}) or use other options such as http.Ok ...
	// return Response(204, nil),nil

	// TODO: Uncomment the next line to return response Response(400, []string{}) or use other options such as http.Ok ...
	// return Response(400, []string{}), nil

	// TODO: Uncomment the next line to return response Response(404, []string{}) or use other options such as http.Ok ...
	// return Response(404, []string{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("DissociateOutcallExtension method not implemented")
}

// DissociateParkingLotExtension - Dissociate parking lot and extension
func (s *ExtensionsAPIService) DissociateParkingLotExtension(ctx context.Context, parkingLotId int32, extensionId int32) (ImplResponse, error) {
	// TODO - update DissociateParkingLotExtension with the required logic for this service method.
	// Add api_extensions_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(204, {}) or use other options such as http.Ok ...
	// return Response(204, nil),nil

	// TODO: Uncomment the next line to return response Response(400, []string{}) or use other options such as http.Ok ...
	// return Response(400, []string{}), nil

	// TODO: Uncomment the next line to return response Response(404, []string{}) or use other options such as http.Ok ...
	// return Response(404, []string{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("DissociateParkingLotExtension method not implemented")
}

// DissociateQueueExtension - Dissociate queue and extension
func (s *ExtensionsAPIService) DissociateQueueExtension(ctx context.Context, queueId int32, extensionId int32, accentTenant string) (ImplResponse, error) {
	// TODO - update DissociateQueueExtension with the required logic for this service method.
	// Add api_extensions_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(204, {}) or use other options such as http.Ok ...
	// return Response(204, nil),nil

	// TODO: Uncomment the next line to return response Response(400, []string{}) or use other options such as http.Ok ...
	// return Response(400, []string{}), nil

	// TODO: Uncomment the next line to return response Response(404, []string{}) or use other options such as http.Ok ...
	// return Response(404, []string{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("DissociateQueueExtension method not implemented")
}

// GetExtension - Get extension
func (s *ExtensionsAPIService) GetExtension(ctx context.Context, extensionId int32) (ImplResponse, error) {
	// TODO - update GetExtension with the required logic for this service method.
	// Add api_extensions_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(200, Extension{}) or use other options such as http.Ok ...
	// return Response(200, Extension{}), nil

	// TODO: Uncomment the next line to return response Response(404, []string{}) or use other options such as http.Ok ...
	// return Response(404, []string{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("GetExtension method not implemented")
}

// GetExtensionFeature - Get extension feature
func (s *ExtensionsAPIService) GetExtensionFeature(ctx context.Context, extensionUuid string) (ImplResponse, error) {
	// TODO - update GetExtensionFeature with the required logic for this service method.
	// Add api_extensions_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(200, ExtensionFeature{}) or use other options such as http.Ok ...
	// return Response(200, ExtensionFeature{}), nil

	// TODO: Uncomment the next line to return response Response(404, []string{}) or use other options such as http.Ok ...
	// return Response(404, []string{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("GetExtensionFeature method not implemented")
}

// ListExtensions - List extensions
func (s *ExtensionsAPIService) ListExtensions(ctx context.Context, accentTenant string, recurse bool, order string, direction string, limit int32, offset int32, search string, type_ string, exten string, context string) (ImplResponse, error) {
	// TODO - update ListExtensions with the required logic for this service method.
	// Add api_extensions_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(200, ExtensionItems{}) or use other options such as http.Ok ...
	// return Response(200, ExtensionItems{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("ListExtensions method not implemented")
}

// ListExtensionsFeatures - List extensions features
func (s *ExtensionsAPIService) ListExtensionsFeatures(ctx context.Context, order string, direction string, limit int32, offset int32, search string) (ImplResponse, error) {
	// TODO - update ListExtensionsFeatures with the required logic for this service method.
	// Add api_extensions_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(200, ExtensionFeatureItems{}) or use other options such as http.Ok ...
	// return Response(200, ExtensionFeatureItems{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("ListExtensionsFeatures method not implemented")
}

// UpdateExtension - Update extension
func (s *ExtensionsAPIService) UpdateExtension(ctx context.Context, extensionId int32, body Extension) (ImplResponse, error) {
	// TODO - update UpdateExtension with the required logic for this service method.
	// Add api_extensions_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(204, {}) or use other options such as http.Ok ...
	// return Response(204, nil),nil

	// TODO: Uncomment the next line to return response Response(400, []string{}) or use other options such as http.Ok ...
	// return Response(400, []string{}), nil

	// TODO: Uncomment the next line to return response Response(404, []string{}) or use other options such as http.Ok ...
	// return Response(404, []string{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("UpdateExtension method not implemented")
}

// UpdateExtensionFeature - Update extension
func (s *ExtensionsAPIService) UpdateExtensionFeature(ctx context.Context, extensionUuid string, body ExtensionFeature) (ImplResponse, error) {
	// TODO - update UpdateExtensionFeature with the required logic for this service method.
	// Add api_extensions_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(204, {}) or use other options such as http.Ok ...
	// return Response(204, nil),nil

	// TODO: Uncomment the next line to return response Response(400, []string{}) or use other options such as http.Ok ...
	// return Response(400, []string{}), nil

	// TODO: Uncomment the next line to return response Response(404, []string{}) or use other options such as http.Ok ...
	// return Response(404, []string{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("UpdateExtensionFeature method not implemented")
}
