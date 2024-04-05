/*
 * accent-chatd
 *
 * Control your message and presence from a REST API
 *
 * API version: 1.0.0
 * Contact: help@accentvoice.io
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package main

import (
	"log"
	"net/http"

	chatdserver "github.com/ryanwclark/accent-voice/src"
)

func main() {
	log.Printf("Server started")

	ConfigAPIService := chatdserver.NewConfigAPIService()
	ConfigAPIController := chatdserver.NewConfigAPIController(ConfigAPIService)

	MessagesAPIService := chatdserver.NewMessagesAPIService()
	MessagesAPIController := chatdserver.NewMessagesAPIController(MessagesAPIService)

	MicrosoftAPIService := chatdserver.NewMicrosoftAPIService()
	MicrosoftAPIController := chatdserver.NewMicrosoftAPIController(MicrosoftAPIService)

	PresencesAPIService := chatdserver.NewPresencesAPIService()
	PresencesAPIController := chatdserver.NewPresencesAPIController(PresencesAPIService)

	RoomsAPIService := chatdserver.NewRoomsAPIService()
	RoomsAPIController := chatdserver.NewRoomsAPIController(RoomsAPIService)

	StatusAPIService := chatdserver.NewStatusAPIService()
	StatusAPIController := chatdserver.NewStatusAPIController(StatusAPIService)

	router := chatdserver.NewRouter(ConfigAPIController, MessagesAPIController, MicrosoftAPIController, PresencesAPIController, RoomsAPIController, StatusAPIController)

	log.Fatal(http.ListenAndServe(":8080", router))
}
