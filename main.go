package main

import (
	"html/template"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/main", mainPageHandler)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		os.Exit(1)
	}
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		// Here, integrate with the internal/client/auth for authentication.
		// If successful, redirect to the main page.
		// If failed, reload the login page with error message.
	} else {
		tmpl, err := template.ParseFiles("login.html")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		tmpl.Execute(w, nil)
	}
}

func mainPageHandler(w http.ResponseWriter, r *http.Request) {
	// Check for user's authentication
	// Serve main page if authenticated
	// Serve login page with error if not authenticated
}