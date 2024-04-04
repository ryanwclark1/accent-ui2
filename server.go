package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	
	mux.HandleFunc("/auth", func(w http.ResponseWriter, r *http.Request) {
		// Use client auth service to validate user credentials
		// This is a placeholder for the actual authentication logic
		// which would involve calling the internal/client/auth service and handling the response
		username := r.URL.Query().Get("username")
		password := r.URL.Query().Get("password")
		if username == "admin" && password == "password" {
			// Redirect to main page on successful login
			w.Header().Set("Location", "/main")
			w.WriteHeader(http.StatusFound)
		} else {
			// Inform the user of incorrect login details
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprint(w, "Incorrect username or password")
		}
	})
	
	// Serve static pages using the templ library
	// This is a placeholder for static files serving
	mux.Handle("/", http.FileServer(http.Dir("./static")))
	
	log.Fatal(http.ListenAndServe(":8080", mux))
}