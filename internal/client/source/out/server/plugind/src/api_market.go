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
	"encoding/json"
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5"
)

// MarketAPIController binds http requests to an api service and writes the service results to the http response
type MarketAPIController struct {
	service      MarketAPIServicer
	errorHandler ErrorHandler
}

// MarketAPIOption for how the controller is set up.
type MarketAPIOption func(*MarketAPIController)

// WithMarketAPIErrorHandler inject ErrorHandler into controller
func WithMarketAPIErrorHandler(h ErrorHandler) MarketAPIOption {
	return func(c *MarketAPIController) {
		c.errorHandler = h
	}
}

// NewMarketAPIController creates a default api controller
func NewMarketAPIController(s MarketAPIServicer, opts ...MarketAPIOption) Router {
	controller := &MarketAPIController{
		service:      s,
		errorHandler: DefaultErrorHandler,
	}

	for _, opt := range opts {
		opt(controller)
	}

	return controller
}

// Routes returns all the api routes for the MarketAPIController
func (c *MarketAPIController) Routes() Routes {
	return Routes{
		"GetMarket": Route{
			strings.ToUpper("Get"),
			"/0.2/market",
			c.GetMarket,
		},
		"GetMarketPlugin": Route{
			strings.ToUpper("Get"),
			"/0.2/market/{namespace}/{name}",
			c.GetMarketPlugin,
		},
	}
}

// GetMarket - List plugins available on the configured market
func (c *MarketAPIController) GetMarket(w http.ResponseWriter, r *http.Request) {
	query, err := parseQuery(r.URL.RawQuery)
	if err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	var limitParam int32
	if query.Has("limit") {
		param, err := parseNumericParameter[int32](
			query.Get("limit"),
			WithParse[int32](parseInt32),
		)
		if err != nil {
			c.errorHandler(w, r, &ParsingError{Err: err}, nil)
			return
		}

		limitParam = param
	} else {
	}
	var offsetParam int32
	if query.Has("offset") {
		param, err := parseNumericParameter[int32](
			query.Get("offset"),
			WithParse[int32](parseInt32),
		)
		if err != nil {
			c.errorHandler(w, r, &ParsingError{Err: err}, nil)
			return
		}

		offsetParam = param
	} else {
	}
	var orderParam string
	if query.Has("order") {
		param := query.Get("order")

		orderParam = param
	} else {
	}
	var directionParam string
	if query.Has("direction") {
		param := query.Get("direction")

		directionParam = param
	} else {
	}
	var searchParam string
	if query.Has("search") {
		param := query.Get("search")

		searchParam = param
	} else {
	}
	var namespaceParam string
	if query.Has("namespace") {
		param := query.Get("namespace")

		namespaceParam = param
	} else {
	}
	var nameParam string
	if query.Has("name") {
		param := query.Get("name")

		nameParam = param
	} else {
	}
	var installedParam bool
	if query.Has("installed") {
		param, err := parseBoolParameter(
			query.Get("installed"),
			WithParse[bool](parseBool),
		)
		if err != nil {
			c.errorHandler(w, r, &ParsingError{Err: err}, nil)
			return
		}

		installedParam = param
	} else {
	}
	result, err := c.service.GetMarket(r.Context(), limitParam, offsetParam, orderParam, directionParam, searchParam, namespaceParam, nameParam, installedParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, result.Headers, w)
}

// GetMarketPlugin - Fetch the information about a plugin from the market
func (c *MarketAPIController) GetMarketPlugin(w http.ResponseWriter, r *http.Request) {
	namespaceParam := chi.URLParam(r, "namespace")
	if namespaceParam == "" {
		c.errorHandler(w, r, &RequiredError{"namespace"}, nil)
		return
	}
	nameParam := chi.URLParam(r, "name")
	if nameParam == "" {
		c.errorHandler(w, r, &RequiredError{"name"}, nil)
		return
	}
	result, err := c.service.GetMarketPlugin(r.Context(), namespaceParam, nameParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, result.Headers, w)
}
