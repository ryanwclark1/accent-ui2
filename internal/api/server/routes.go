package server

import (
	"embed"
	"net/http"
)


//go:embed static
var static embed.FS

func (s *Server) Routes() error {
	// s.r.Handle("GET /static/*", s.mw(s.HandleAssets(static)))
	s.r.Handle("GET /favicon.ico", s.HandleFavicon(static))

	s.r.HandleFunc("/api/agent", agentHandler)
	s.r.HandleFunc("/api/amid", amidHandler)
	s.r.HandleFunc("/api/auth", authHandler)
	s.r.HandleFunc("/api/call", callHandler)
	s.r.HandleFunc("/api/chat", chatHandler)
	s.r.HandleFunc("/api/conf", confHandler)
	s.r.HandleFunc("/api/dird", dirdHandler)
	s.r.HandleFunc("/api/phone", phoneHandler)
	s.r.HandleFunc("/api/plugin", pluginHandler)
	s.r.HandleFunc("/api/websocket", websocketHandler)

	return nil
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
