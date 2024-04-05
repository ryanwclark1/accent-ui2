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

// DevicesAPIService is a service that implements the logic for the DevicesAPIServicer
// This service should implement the business logic for every endpoint for the DevicesAPI API.
// Include any external packages or services that will be required by this service.
type DevicesAPIService struct {
}

// NewDevicesAPIService creates a default api service
func NewDevicesAPIService() DevicesAPIServicer {
	return &DevicesAPIService{}
}

// AssignUnallocatedDeviceTenant - Assign unallocated device tenant
func (s *DevicesAPIService) AssignUnallocatedDeviceTenant(ctx context.Context, deviceId string, accentTenant string) (ImplResponse, error) {
	// TODO - update AssignUnallocatedDeviceTenant with the required logic for this service method.
	// Add api_devices_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(204, {}) or use other options such as http.Ok ...
	// return Response(204, nil),nil

	// TODO: Uncomment the next line to return response Response(400, []string{}) or use other options such as http.Ok ...
	// return Response(400, []string{}), nil

	// TODO: Uncomment the next line to return response Response(404, []string{}) or use other options such as http.Ok ...
	// return Response(404, []string{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("AssignUnallocatedDeviceTenant method not implemented")
}

// AssociateLineDevice - Associate line and device
func (s *DevicesAPIService) AssociateLineDevice(ctx context.Context, lineId int32, deviceId string, accentTenant string) (ImplResponse, error) {
	// TODO - update AssociateLineDevice with the required logic for this service method.
	// Add api_devices_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(204, {}) or use other options such as http.Ok ...
	// return Response(204, nil),nil

	// TODO: Uncomment the next line to return response Response(400, []string{}) or use other options such as http.Ok ...
	// return Response(400, []string{}), nil

	// TODO: Uncomment the next line to return response Response(404, []string{}) or use other options such as http.Ok ...
	// return Response(404, []string{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("AssociateLineDevice method not implemented")
}

// CreateDevice - Create device
func (s *DevicesAPIService) CreateDevice(ctx context.Context, accentTenant string, body Device) (ImplResponse, error) {
	// TODO - update CreateDevice with the required logic for this service method.
	// Add api_devices_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(201, Device{}) or use other options such as http.Ok ...
	// return Response(201, Device{}), nil

	// TODO: Uncomment the next line to return response Response(400, []string{}) or use other options such as http.Ok ...
	// return Response(400, []string{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("CreateDevice method not implemented")
}

// DeleteDevice - Delete device
func (s *DevicesAPIService) DeleteDevice(ctx context.Context, deviceId string, accentTenant string) (ImplResponse, error) {
	// TODO - update DeleteDevice with the required logic for this service method.
	// Add api_devices_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(204, {}) or use other options such as http.Ok ...
	// return Response(204, nil),nil

	// TODO: Uncomment the next line to return response Response(400, []string{}) or use other options such as http.Ok ...
	// return Response(400, []string{}), nil

	// TODO: Uncomment the next line to return response Response(404, []string{}) or use other options such as http.Ok ...
	// return Response(404, []string{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("DeleteDevice method not implemented")
}

// DissociateLineDevice - Dissociate line and device
func (s *DevicesAPIService) DissociateLineDevice(ctx context.Context, lineId int32, deviceId string, accentTenant string) (ImplResponse, error) {
	// TODO - update DissociateLineDevice with the required logic for this service method.
	// Add api_devices_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(204, {}) or use other options such as http.Ok ...
	// return Response(204, nil),nil

	// TODO: Uncomment the next line to return response Response(400, []string{}) or use other options such as http.Ok ...
	// return Response(400, []string{}), nil

	// TODO: Uncomment the next line to return response Response(404, []string{}) or use other options such as http.Ok ...
	// return Response(404, []string{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("DissociateLineDevice method not implemented")
}

// GetDevice - Get device
func (s *DevicesAPIService) GetDevice(ctx context.Context, deviceId string, accentTenant string) (ImplResponse, error) {
	// TODO - update GetDevice with the required logic for this service method.
	// Add api_devices_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(200, Device{}) or use other options such as http.Ok ...
	// return Response(200, Device{}), nil

	// TODO: Uncomment the next line to return response Response(404, []string{}) or use other options such as http.Ok ...
	// return Response(404, []string{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("GetDevice method not implemented")
}

// GetDeviceLineAssociation - List lines associated to device
func (s *DevicesAPIService) GetDeviceLineAssociation(ctx context.Context, deviceId string, accentTenant string) (ImplResponse, error) {
	// TODO - update GetDeviceLineAssociation with the required logic for this service method.
	// Add api_devices_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(200, LineDeviceItems{}) or use other options such as http.Ok ...
	// return Response(200, LineDeviceItems{}), nil

	// TODO: Uncomment the next line to return response Response(404, []string{}) or use other options such as http.Ok ...
	// return Response(404, []string{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("GetDeviceLineAssociation method not implemented")
}

// GetLineDevice - Get Device associated to Line
func (s *DevicesAPIService) GetLineDevice(ctx context.Context, lineId int32, accentTenant string) (ImplResponse, error) {
	// TODO - update GetLineDevice with the required logic for this service method.
	// Add api_devices_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(200, LineDevice{}) or use other options such as http.Ok ...
	// return Response(200, LineDevice{}), nil

	// TODO: Uncomment the next line to return response Response(400, []string{}) or use other options such as http.Ok ...
	// return Response(400, []string{}), nil

	// TODO: Uncomment the next line to return response Response(404, []string{}) or use other options such as http.Ok ...
	// return Response(404, []string{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("GetLineDevice method not implemented")
}

// ListDevices - List devices
func (s *DevicesAPIService) ListDevices(ctx context.Context, accentTenant string, recurse bool, order string, direction string, limit int32, offset int32, search string) (ImplResponse, error) {
	// TODO - update ListDevices with the required logic for this service method.
	// Add api_devices_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(200, DeviceItems{}) or use other options such as http.Ok ...
	// return Response(200, DeviceItems{}), nil

	// TODO: Uncomment the next line to return response Response(400, []string{}) or use other options such as http.Ok ...
	// return Response(400, []string{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("ListDevices method not implemented")
}

// ListUnallocatedDevices - List unallocated devices
func (s *DevicesAPIService) ListUnallocatedDevices(ctx context.Context, order string, direction string, limit int32, offset int32, search string) (ImplResponse, error) {
	// TODO - update ListUnallocatedDevices with the required logic for this service method.
	// Add api_devices_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(200, DeviceItems{}) or use other options such as http.Ok ...
	// return Response(200, DeviceItems{}), nil

	// TODO: Uncomment the next line to return response Response(400, []string{}) or use other options such as http.Ok ...
	// return Response(400, []string{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("ListUnallocatedDevices method not implemented")
}

// ResetDeviceAutoprov - Reset device to autoprov
func (s *DevicesAPIService) ResetDeviceAutoprov(ctx context.Context, deviceId string, accentTenant string) (ImplResponse, error) {
	// TODO - update ResetDeviceAutoprov with the required logic for this service method.
	// Add api_devices_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(204, {}) or use other options such as http.Ok ...
	// return Response(204, nil),nil

	// TODO: Uncomment the next line to return response Response(400, []string{}) or use other options such as http.Ok ...
	// return Response(400, []string{}), nil

	// TODO: Uncomment the next line to return response Response(404, []string{}) or use other options such as http.Ok ...
	// return Response(404, []string{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("ResetDeviceAutoprov method not implemented")
}

// SynchronizeDevice - Synchronize device
func (s *DevicesAPIService) SynchronizeDevice(ctx context.Context, deviceId string, accentTenant string) (ImplResponse, error) {
	// TODO - update SynchronizeDevice with the required logic for this service method.
	// Add api_devices_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(204, {}) or use other options such as http.Ok ...
	// return Response(204, nil),nil

	// TODO: Uncomment the next line to return response Response(400, []string{}) or use other options such as http.Ok ...
	// return Response(400, []string{}), nil

	// TODO: Uncomment the next line to return response Response(404, []string{}) or use other options such as http.Ok ...
	// return Response(404, []string{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("SynchronizeDevice method not implemented")
}

// UpdateDevice - Update device
func (s *DevicesAPIService) UpdateDevice(ctx context.Context, deviceId string, body Device, accentTenant string) (ImplResponse, error) {
	// TODO - update UpdateDevice with the required logic for this service method.
	// Add api_devices_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(204, {}) or use other options such as http.Ok ...
	// return Response(204, nil),nil

	// TODO: Uncomment the next line to return response Response(400, []string{}) or use other options such as http.Ok ...
	// return Response(400, []string{}), nil

	// TODO: Uncomment the next line to return response Response(404, []string{}) or use other options such as http.Ok ...
	// return Response(404, []string{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("UpdateDevice method not implemented")
}
