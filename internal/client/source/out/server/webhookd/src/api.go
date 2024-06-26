/*
 * accent-webhookd
 *
 * Control your webhooks from a REST API
 *
 * API version: 1.0.0
 * Contact: help@accentvoice.io
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package webhookdserver

import (
	"context"
	"net/http"
)



// ConfigAPIRouter defines the required methods for binding the api requests to a responses for the ConfigAPI
// The ConfigAPIRouter implementation should parse necessary information from the http request,
// pass the data to a ConfigAPIServicer to perform the required actions, then write the service results to the http response.
type ConfigAPIRouter interface { 
	GetConfig(http.ResponseWriter, *http.Request)
	PatchConfig(http.ResponseWriter, *http.Request)
}
// NotificationsAPIRouter defines the required methods for binding the api requests to a responses for the NotificationsAPI
// The NotificationsAPIRouter implementation should parse necessary information from the http request,
// pass the data to a NotificationsAPIServicer to perform the required actions, then write the service results to the http response.
type NotificationsAPIRouter interface { 
	PostMobileNotification(http.ResponseWriter, *http.Request)
}
// StatusAPIRouter defines the required methods for binding the api requests to a responses for the StatusAPI
// The StatusAPIRouter implementation should parse necessary information from the http request,
// pass the data to a StatusAPIServicer to perform the required actions, then write the service results to the http response.
type StatusAPIRouter interface { 
	GetStatus(http.ResponseWriter, *http.Request)
}
// SubscriptionsAPIRouter defines the required methods for binding the api requests to a responses for the SubscriptionsAPI
// The SubscriptionsAPIRouter implementation should parse necessary information from the http request,
// pass the data to a SubscriptionsAPIServicer to perform the required actions, then write the service results to the http response.
type SubscriptionsAPIRouter interface { 
	Create(http.ResponseWriter, *http.Request)
	Delete(http.ResponseWriter, *http.Request)
	Edit(http.ResponseWriter, *http.Request)
	Get(http.ResponseWriter, *http.Request)
	GetLogs(http.ResponseWriter, *http.Request)
	GetSubscriptionsServices(http.ResponseWriter, *http.Request)
	GetUserSubscription(http.ResponseWriter, *http.Request)
	List(http.ResponseWriter, *http.Request)
	UserCreate(http.ResponseWriter, *http.Request)
	UserDelete(http.ResponseWriter, *http.Request)
	UserList(http.ResponseWriter, *http.Request)
}
// UsersAPIRouter defines the required methods for binding the api requests to a responses for the UsersAPI
// The UsersAPIRouter implementation should parse necessary information from the http request,
// pass the data to a UsersAPIServicer to perform the required actions, then write the service results to the http response.
type UsersAPIRouter interface { 
	GetUserSubscription(http.ResponseWriter, *http.Request)
	UserCreate(http.ResponseWriter, *http.Request)
	UserDelete(http.ResponseWriter, *http.Request)
	UserList(http.ResponseWriter, *http.Request)
}


// ConfigAPIServicer defines the api actions for the ConfigAPI service
// This interface intended to stay up to date with the openapi yaml used to generate it,
// while the service implementation can be ignored with the .openapi-generator-ignore file
// and updated with the logic required for the API.
type ConfigAPIServicer interface { 
	GetConfig(context.Context) (ImplResponse, error)
	PatchConfig(context.Context, []ConfigPatchItem) (ImplResponse, error)
}


// NotificationsAPIServicer defines the api actions for the NotificationsAPI service
// This interface intended to stay up to date with the openapi yaml used to generate it,
// while the service implementation can be ignored with the .openapi-generator-ignore file
// and updated with the logic required for the API.
type NotificationsAPIServicer interface { 
	PostMobileNotification(context.Context, Notification, string) (ImplResponse, error)
}


// StatusAPIServicer defines the api actions for the StatusAPI service
// This interface intended to stay up to date with the openapi yaml used to generate it,
// while the service implementation can be ignored with the .openapi-generator-ignore file
// and updated with the logic required for the API.
type StatusAPIServicer interface { 
	GetStatus(context.Context) (ImplResponse, error)
}


// SubscriptionsAPIServicer defines the api actions for the SubscriptionsAPI service
// This interface intended to stay up to date with the openapi yaml used to generate it,
// while the service implementation can be ignored with the .openapi-generator-ignore file
// and updated with the logic required for the API.
type SubscriptionsAPIServicer interface { 
	Create(context.Context, SubscriptionRequest) (ImplResponse, error)
	Delete(context.Context, string) (ImplResponse, error)
	Edit(context.Context, string, SubscriptionRequest) (ImplResponse, error)
	Get(context.Context, string) (ImplResponse, error)
	GetLogs(context.Context, string) (ImplResponse, error)
	GetSubscriptionsServices(context.Context) (ImplResponse, error)
	GetUserSubscription(context.Context, string) (ImplResponse, error)
	List(context.Context, string, bool, string) (ImplResponse, error)
	UserCreate(context.Context, UserSubscriptionRequest) (ImplResponse, error)
	UserDelete(context.Context, string) (ImplResponse, error)
	UserList(context.Context, string) (ImplResponse, error)
}


// UsersAPIServicer defines the api actions for the UsersAPI service
// This interface intended to stay up to date with the openapi yaml used to generate it,
// while the service implementation can be ignored with the .openapi-generator-ignore file
// and updated with the logic required for the API.
type UsersAPIServicer interface { 
	GetUserSubscription(context.Context, string) (ImplResponse, error)
	UserCreate(context.Context, UserSubscriptionRequest) (ImplResponse, error)
	UserDelete(context.Context, string) (ImplResponse, error)
	UserList(context.Context, string) (ImplResponse, error)
}
