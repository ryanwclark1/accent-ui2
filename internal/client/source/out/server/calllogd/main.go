/*
 * accent-call-logd
 *
 * Consult call logs from a REST API
 *
 * API version: 1.0.0
 * Contact: help@accentvoice.io
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package main

import (
	"log"
	"net/http"

	calllogdserver "github.com/ryanwclark/accent-voice/src"
)

func main() {
	log.Printf("Server started")

	AgentStatisticsAPIService := calllogdserver.NewAgentStatisticsAPIService()
	AgentStatisticsAPIController := calllogdserver.NewAgentStatisticsAPIController(AgentStatisticsAPIService)

	CdrAPIService := calllogdserver.NewCdrAPIService()
	CdrAPIController := calllogdserver.NewCdrAPIController(CdrAPIService)

	ConfigAPIService := calllogdserver.NewConfigAPIService()
	ConfigAPIController := calllogdserver.NewConfigAPIController(ConfigAPIService)

	ExportsAPIService := calllogdserver.NewExportsAPIService()
	ExportsAPIController := calllogdserver.NewExportsAPIController(ExportsAPIService)

	QueueStatisticsAPIService := calllogdserver.NewQueueStatisticsAPIService()
	QueueStatisticsAPIController := calllogdserver.NewQueueStatisticsAPIController(QueueStatisticsAPIService)

	RetentionAPIService := calllogdserver.NewRetentionAPIService()
	RetentionAPIController := calllogdserver.NewRetentionAPIController(RetentionAPIService)

	StatusAPIService := calllogdserver.NewStatusAPIService()
	StatusAPIController := calllogdserver.NewStatusAPIController(StatusAPIService)

	UsersAPIService := calllogdserver.NewUsersAPIService()
	UsersAPIController := calllogdserver.NewUsersAPIController(UsersAPIService)

	router := calllogdserver.NewRouter(AgentStatisticsAPIController, CdrAPIController, ConfigAPIController, ExportsAPIController, QueueStatisticsAPIController, RetentionAPIController, StatusAPIController, UsersAPIController)

	log.Fatal(http.ListenAndServe(":8080", router))
}
