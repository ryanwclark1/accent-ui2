package server

import (
	"embed"
	"io/fs"
	"log/slog"
	"net/http"

)

func (s *Server) HandleAssets(assets embed.FS) http.Handler {
	contentAssets, err := fs.Sub(fs.FS(assets), "static")
	if err != nil {
		slog.Info("HandleAssets: failed to load assets: %v", err)
	}
	return http.StripPrefix("/static/", http.FileServerFS(contentAssets))
}

func (s *Server) HandleFavicon(assets embed.FS) http.Handler {
	b, err := assets.ReadFile("static/favicon.ico")
	if err != nil {
		slog.Info("HandleFavicon: failed to read favicon.ico: %v", err)
	}
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "image/x-icon")
		w.Write(b)
	})
}
