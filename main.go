package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/julienschmidt/httprouter"
)

func AuthHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// In production, you'd make a request to the client auth service here.
	// This is simplified for demonstration.
	username := r.FormValue("username")
	password := r.FormValue("password")
	if username == "test" && password == "password" {
		fmt.Fprintln(w, "Login successful, welcome!")
	} else {
		http.Error(w, "Invalid login", http.StatusUnauthorized)
	}
}

func main() {
	router := httprouter.New()
	router.POST("/auth", AuthHandler)

	log.Fatal(http.ListenAndServe(":8080", router))
}