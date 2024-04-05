/*
 * accent-provd
 *
 * Provisioning application REST API
 *
 * API version: 0.2
 * Contact: help@accentvoice.io
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package main

import (
	"log"
	"net/http"

	provdserver "github.com/ryanwclark/accent-voice/src"
)

func main() {
	log.Printf("Server started")

	ConfigsAPIService := provdserver.NewConfigsAPIService()
	ConfigsAPIController := provdserver.NewConfigsAPIController(ConfigsAPIService)

	DevicesAPIService := provdserver.NewDevicesAPIService()
	DevicesAPIController := provdserver.NewDevicesAPIController(DevicesAPIService)

	PluginsAPIService := provdserver.NewPluginsAPIService()
	PluginsAPIController := provdserver.NewPluginsAPIController(PluginsAPIService)

	ProvdAPIService := provdserver.NewProvdAPIService()
	ProvdAPIController := provdserver.NewProvdAPIController(ProvdAPIService)

	StatusAPIService := provdserver.NewStatusAPIService()
	StatusAPIController := provdserver.NewStatusAPIController(StatusAPIService)

	router := provdserver.NewRouter(ConfigsAPIController, DevicesAPIController, PluginsAPIController, ProvdAPIController, StatusAPIController)

	log.Fatal(http.ListenAndServe(":8080", router))
}