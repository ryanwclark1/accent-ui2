package main

import (
	"embed"
	"net/http"
)

//go:embed public/*
var staticFiles embed.FS

func main() {
	http.Handle("/", http.FileServer(http.FS(staticFiles)))
	// Setup routes for the REST APIs
	http.HandleFunc("/api/agent", agentHandler)
	http.HandleFunc("/api/amid", amidHandler)
	http.HandleFunc("/api/auth", authHandler)
	http.HandleFunc("/api/call", callHandler)
	http.HandleFunc("/api/chat", chatHandler)
	http.HandleFunc("/api/conf", confHandler)
	http.HandleFunc("/api/dird", dirdHandler)
	http.HandleFunc("/api/phone", phoneHandler)
	http.HandleFunc("/api/plugin", pluginHandler)
	http.HandleFunc("/api/websocket", websocketHandler)

	// Start server
	if err := http.ListenAndServe(":8080", nil); err != nil {
	    panic(err)
	}
}

// Placeholder handlers for the API endpoints
func agentHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Agent API Endpoint"))
}
func amidHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("AMID API Endpoint"))
}
func authHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Auth API Endpoint"))
}
func callHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Call API Endpoint"))
}
func chatHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Chat API Endpoint"))
}
func confHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Conf API Endpoint"))
}
func dirdHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Dird API Endpoint"))
}
func phoneHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Phone API Endpoint"))
}
func pluginHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Plugin API Endpoint"))
}
func websocketHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("WebSocket API Endpoint"))
}