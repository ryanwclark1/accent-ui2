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

// WizardAPIService is a service that implements the logic for the WizardAPIServicer
// This service should implement the business logic for every endpoint for the WizardAPI API.
// Include any external packages or services that will be required by this service.
type WizardAPIService struct {
}

// NewWizardAPIService creates a default api service
func NewWizardAPIService() WizardAPIServicer {
	return &WizardAPIService{}
}

// GetWizardDiscover - Get wizard discover
func (s *WizardAPIService) GetWizardDiscover(ctx context.Context) (ImplResponse, error) {
	// TODO - update GetWizardDiscover with the required logic for this service method.
	// Add api_wizard_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(200, WizardDiscover{}) or use other options such as http.Ok ...
	// return Response(200, WizardDiscover{}), nil

	// TODO: Uncomment the next line to return response Response(403, []string{}) or use other options such as http.Ok ...
	// return Response(403, []string{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("GetWizardDiscover method not implemented")
}

// GetWizardStatus - Get wizard status
func (s *WizardAPIService) GetWizardStatus(ctx context.Context) (ImplResponse, error) {
	// TODO - update GetWizardStatus with the required logic for this service method.
	// Add api_wizard_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(200, WizardConfigured{}) or use other options such as http.Ok ...
	// return Response(200, WizardConfigured{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("GetWizardStatus method not implemented")
}

// PassWizard - Pass the wizard
func (s *WizardAPIService) PassWizard(ctx context.Context, body Wizard) (ImplResponse, error) {
	// TODO - update PassWizard with the required logic for this service method.
	// Add api_wizard_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(201, Wizard{}) or use other options such as http.Ok ...
	// return Response(201, Wizard{}), nil

	// TODO: Uncomment the next line to return response Response(400, []string{}) or use other options such as http.Ok ...
	// return Response(400, []string{}), nil

	// TODO: Uncomment the next line to return response Response(403, []string{}) or use other options such as http.Ok ...
	// return Response(403, []string{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("PassWizard method not implemented")
}
