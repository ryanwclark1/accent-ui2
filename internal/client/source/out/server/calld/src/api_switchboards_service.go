/*
 * accent-calld
 *
 * Control your calls from a REST API
 *
 * API version: 1.0.0
 * Contact: help@accentvoice.io
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package calldserver

import (
	"context"
	"errors"
	"net/http"
)

// SwitchboardsAPIService is a service that implements the logic for the SwitchboardsAPIServicer
// This service should implement the business logic for every endpoint for the SwitchboardsAPI API.
// Include any external packages or services that will be required by this service.
type SwitchboardsAPIService struct {
}

// NewSwitchboardsAPIService creates a default api service
func NewSwitchboardsAPIService() SwitchboardsAPIServicer {
	return &SwitchboardsAPIService{}
}

// AnswerHeldCall - Answer the specified held call
func (s *SwitchboardsAPIService) AnswerHeldCall(ctx context.Context, switchboardUuid string, callId string, accentTenant string, lineId int32) (ImplResponse, error) {
	// TODO - update AnswerHeldCall with the required logic for this service method.
	// Add api_switchboards_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(200, CallId{}) or use other options such as http.Ok ...
	// return Response(200, CallId{}), nil

	// TODO: Uncomment the next line to return response Response(400, Error{}) or use other options such as http.Ok ...
	// return Response(400, Error{}), nil

	// TODO: Uncomment the next line to return response Response(404, Error{}) or use other options such as http.Ok ...
	// return Response(404, Error{}), nil

	// TODO: Uncomment the next line to return response Response(503, Error{}) or use other options such as http.Ok ...
	// return Response(503, Error{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("AnswerHeldCall method not implemented")
}

// AnswerQueuedCall - Answer the specified queued call
func (s *SwitchboardsAPIService) AnswerQueuedCall(ctx context.Context, switchboardUuid string, callId string, accentTenant string, lineId int32) (ImplResponse, error) {
	// TODO - update AnswerQueuedCall with the required logic for this service method.
	// Add api_switchboards_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(200, CallId{}) or use other options such as http.Ok ...
	// return Response(200, CallId{}), nil

	// TODO: Uncomment the next line to return response Response(400, Error{}) or use other options such as http.Ok ...
	// return Response(400, Error{}), nil

	// TODO: Uncomment the next line to return response Response(404, Error{}) or use other options such as http.Ok ...
	// return Response(404, Error{}), nil

	// TODO: Uncomment the next line to return response Response(503, Error{}) or use other options such as http.Ok ...
	// return Response(503, Error{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("AnswerQueuedCall method not implemented")
}

// HoldSwitchboardCall - Put the specified call on hold in the switchboard
func (s *SwitchboardsAPIService) HoldSwitchboardCall(ctx context.Context, switchboardUuid string, callId string, accentTenant string) (ImplResponse, error) {
	// TODO - update HoldSwitchboardCall with the required logic for this service method.
	// Add api_switchboards_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(204, {}) or use other options such as http.Ok ...
	// return Response(204, nil),nil

	// TODO: Uncomment the next line to return response Response(400, Error{}) or use other options such as http.Ok ...
	// return Response(400, Error{}), nil

	// TODO: Uncomment the next line to return response Response(404, Error{}) or use other options such as http.Ok ...
	// return Response(404, Error{}), nil

	// TODO: Uncomment the next line to return response Response(503, Error{}) or use other options such as http.Ok ...
	// return Response(503, Error{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("HoldSwitchboardCall method not implemented")
}

// ListSwitchboardHeldCalls - List calls held in the switchboard
func (s *SwitchboardsAPIService) ListSwitchboardHeldCalls(ctx context.Context, switchboardUuid string, accentTenant string) (ImplResponse, error) {
	// TODO - update ListSwitchboardHeldCalls with the required logic for this service method.
	// Add api_switchboards_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(200, SwitchboardHeldCalls{}) or use other options such as http.Ok ...
	// return Response(200, SwitchboardHeldCalls{}), nil

	// TODO: Uncomment the next line to return response Response(404, Error{}) or use other options such as http.Ok ...
	// return Response(404, Error{}), nil

	// TODO: Uncomment the next line to return response Response(503, Error{}) or use other options such as http.Ok ...
	// return Response(503, Error{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("ListSwitchboardHeldCalls method not implemented")
}

// ListSwitchboardQueuedCalls - List calls queued in the switchboard
func (s *SwitchboardsAPIService) ListSwitchboardQueuedCalls(ctx context.Context, switchboardUuid string, accentTenant string) (ImplResponse, error) {
	// TODO - update ListSwitchboardQueuedCalls with the required logic for this service method.
	// Add api_switchboards_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(200, SwitchboardQueuedCalls{}) or use other options such as http.Ok ...
	// return Response(200, SwitchboardQueuedCalls{}), nil

	// TODO: Uncomment the next line to return response Response(404, Error{}) or use other options such as http.Ok ...
	// return Response(404, Error{}), nil

	// TODO: Uncomment the next line to return response Response(503, Error{}) or use other options such as http.Ok ...
	// return Response(503, Error{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("ListSwitchboardQueuedCalls method not implemented")
}
