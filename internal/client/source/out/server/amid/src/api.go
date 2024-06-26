/*
 * accent-amid
 *
 * Send AMI actions to Asterisk, providing token based authentication.
 *
 * API version: 0.1
 * Contact: help@accentvoice.io
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package amidserver

import (
	"context"
	"net/http"
)



// ActionAPIRouter defines the required methods for binding the api requests to a responses for the ActionAPI
// The ActionAPIRouter implementation should parse necessary information from the http request,
// pass the data to a ActionAPIServicer to perform the required actions, then write the service results to the http response.
type ActionAPIRouter interface { 
	ActionAsteriskManager(http.ResponseWriter, *http.Request)
}
// CommandAPIRouter defines the required methods for binding the api requests to a responses for the CommandAPI
// The CommandAPIRouter implementation should parse necessary information from the http request,
// pass the data to a CommandAPIServicer to perform the required actions, then write the service results to the http response.
type CommandAPIRouter interface { 
	CommandAsteriskManager(http.ResponseWriter, *http.Request)
}
// StatusAPIRouter defines the required methods for binding the api requests to a responses for the StatusAPI
// The StatusAPIRouter implementation should parse necessary information from the http request,
// pass the data to a StatusAPIServicer to perform the required actions, then write the service results to the http response.
type StatusAPIRouter interface { 
	StatusSummary(http.ResponseWriter, *http.Request)
}


// ActionAPIServicer defines the api actions for the ActionAPI service
// This interface intended to stay up to date with the openapi yaml used to generate it,
// while the service implementation can be ignored with the .openapi-generator-ignore file
// and updated with the logic required for the API.
type ActionAPIServicer interface { 
	ActionAsteriskManager(context.Context, string, map[string]interface{}) (ImplResponse, error)
}


// CommandAPIServicer defines the api actions for the CommandAPI service
// This interface intended to stay up to date with the openapi yaml used to generate it,
// while the service implementation can be ignored with the .openapi-generator-ignore file
// and updated with the logic required for the API.
type CommandAPIServicer interface { 
	CommandAsteriskManager(context.Context, Command) (ImplResponse, error)
}


// StatusAPIServicer defines the api actions for the StatusAPI service
// This interface intended to stay up to date with the openapi yaml used to generate it,
// while the service implementation can be ignored with the .openapi-generator-ignore file
// and updated with the logic required for the API.
type StatusAPIServicer interface { 
	StatusSummary(context.Context) (ImplResponse, error)
}
