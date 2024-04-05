/*
 * accent-plugind
 *
 * Accent's plugin management service
 *
 * API version: 0.2
 * Contact: help@accentvoice.io
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package plugindserver

import (
	"context"
	"errors"
	"net/http"
)

// PluginAPIService is a service that implements the logic for the PluginAPIServicer
// This service should implement the business logic for every endpoint for the PluginAPI API.
// Include any external packages or services that will be required by this service.
type PluginAPIService struct {
}

// NewPluginAPIService creates a default api service
func NewPluginAPIService() PluginAPIServicer {
	return &PluginAPIService{}
}

// GetMarket - List plugins available on the configured market
func (s *PluginAPIService) GetMarket(ctx context.Context, limit int32, offset int32, order string, direction string, search string, namespace string, name string, installed bool) (ImplResponse, error) {
	// TODO - update GetMarket with the required logic for this service method.
	// Add api_plugin_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(200, GetMarketResult{}) or use other options such as http.Ok ...
	// return Response(200, GetMarketResult{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("GetMarket method not implemented")
}

// GetMarketPlugin - Fetch the information about a plugin from the market
func (s *PluginAPIService) GetMarketPlugin(ctx context.Context, namespace string, name string) (ImplResponse, error) {
	// TODO - update GetMarketPlugin with the required logic for this service method.
	// Add api_plugin_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(200, MarketPluginList{}) or use other options such as http.Ok ...
	// return Response(200, MarketPluginList{}), nil

	// TODO: Uncomment the next line to return response Response(400, Error{}) or use other options such as http.Ok ...
	// return Response(400, Error{}), nil

	// TODO: Uncomment the next line to return response Response(404, Error{}) or use other options such as http.Ok ...
	// return Response(404, Error{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("GetMarketPlugin method not implemented")
}

// GetPlugin - Fetch the information about a plugin that has been installed
func (s *PluginAPIService) GetPlugin(ctx context.Context, namespace string, name string) (ImplResponse, error) {
	// TODO - update GetPlugin with the required logic for this service method.
	// Add api_plugin_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(200, PluginMetadata{}) or use other options such as http.Ok ...
	// return Response(200, PluginMetadata{}), nil

	// TODO: Uncomment the next line to return response Response(400, Error{}) or use other options such as http.Ok ...
	// return Response(400, Error{}), nil

	// TODO: Uncomment the next line to return response Response(404, Error{}) or use other options such as http.Ok ...
	// return Response(404, Error{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("GetPlugin method not implemented")
}

// GetPlugins - List installed plugins
func (s *PluginAPIService) GetPlugins(ctx context.Context) (ImplResponse, error) {
	// TODO - update GetPlugins with the required logic for this service method.
	// Add api_plugin_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(200, GetPluginsResult{}) or use other options such as http.Ok ...
	// return Response(200, GetPluginsResult{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("GetPlugins method not implemented")
}

// InstallPlugin - Install a plugin
func (s *PluginAPIService) InstallPlugin(ctx context.Context, body PluginInstallParameters, reinstall bool) (ImplResponse, error) {
	// TODO - update InstallPlugin with the required logic for this service method.
	// Add api_plugin_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(200, InstallResponse{}) or use other options such as http.Ok ...
	// return Response(200, InstallResponse{}), nil

	// TODO: Uncomment the next line to return response Response(400, Error{}) or use other options such as http.Ok ...
	// return Response(400, Error{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("InstallPlugin method not implemented")
}

// UninstallPlugin - Uninstall a plugin
func (s *PluginAPIService) UninstallPlugin(ctx context.Context, namespace string, name string) (ImplResponse, error) {
	// TODO - update UninstallPlugin with the required logic for this service method.
	// Add api_plugin_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(204, {}) or use other options such as http.Ok ...
	// return Response(204, nil),nil

	// TODO: Uncomment the next line to return response Response(400, Error{}) or use other options such as http.Ok ...
	// return Response(400, Error{}), nil

	// TODO: Uncomment the next line to return response Response(404, Error{}) or use other options such as http.Ok ...
	// return Response(404, Error{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("UninstallPlugin method not implemented")
}
