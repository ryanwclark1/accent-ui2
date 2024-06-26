/*
 * accent-amid
 *
 * Send AMI actions to Asterisk, providing token based authentication.
 *
 * API version: 0.1
 * Contact: help@accentvoice.io
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package main

import (
	"log"
	"net/http"

	amidserver "github.com/ryanwclark/accent-voice/src"
)

func main() {
	log.Printf("Server started")

	ActionAPIService := amidserver.NewActionAPIService()
	ActionAPIController := amidserver.NewActionAPIController(ActionAPIService)

	CommandAPIService := amidserver.NewCommandAPIService()
	CommandAPIController := amidserver.NewCommandAPIController(CommandAPIService)

	StatusAPIService := amidserver.NewStatusAPIService()
	StatusAPIController := amidserver.NewStatusAPIController(StatusAPIService)

	router := amidserver.NewRouter(ActionAPIController, CommandAPIController, StatusAPIController)

	log.Fatal(http.ListenAndServe(":8080", router))
}
