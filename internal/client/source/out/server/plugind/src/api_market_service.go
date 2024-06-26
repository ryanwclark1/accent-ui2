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

// MarketAPIService is a service that implements the logic for the MarketAPIServicer
// This service should implement the business logic for every endpoint for the MarketAPI API.
// Include any external packages or services that will be required by this service.
type MarketAPIService struct {
}

// NewMarketAPIService creates a default api service
func NewMarketAPIService() MarketAPIServicer {
	return &MarketAPIService{}
}

// GetMarket - List plugins available on the configured market
func (s *MarketAPIService) GetMarket(ctx context.Context, limit int32, offset int32, order string, direction string, search string, namespace string, name string, installed bool) (ImplResponse, error) {
	// TODO - update GetMarket with the required logic for this service method.
	// Add api_market_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(200, GetMarketResult{}) or use other options such as http.Ok ...
	// return Response(200, GetMarketResult{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("GetMarket method not implemented")
}

// GetMarketPlugin - Fetch the information about a plugin from the market
func (s *MarketAPIService) GetMarketPlugin(ctx context.Context, namespace string, name string) (ImplResponse, error) {
	// TODO - update GetMarketPlugin with the required logic for this service method.
	// Add api_market_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(200, MarketPluginList{}) or use other options such as http.Ok ...
	// return Response(200, MarketPluginList{}), nil

	// TODO: Uncomment the next line to return response Response(400, Error{}) or use other options such as http.Ok ...
	// return Response(400, Error{}), nil

	// TODO: Uncomment the next line to return response Response(404, Error{}) or use other options such as http.Ok ...
	// return Response(404, Error{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("GetMarketPlugin method not implemented")
}
