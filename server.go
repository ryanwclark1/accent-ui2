package main

import (
	"encoding/json"
	"net/http"
	"github.com/julienschmidt/httprouter"
	"github.com/myusername/myproject/internal/client/auth"
)

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func AuthHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	decoder := json.NewDecoder(r.Body)
	var creds LoginRequest
	err := decoder.Decode(&creds)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	authClient := auth.NewAuthClient("http://api.accentvoice.com/")
	response, err := authClient.Authenticate(&auth.Credentials{Username: creds.Username, Password: creds.Password})
	if err != nil || !response.IsValid {
		http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		return
	}

	// Serve the main page after successful login
	// Placeholder for main page serving logic
}

func main() {
	router := httprouter.New()
	router.POST("/auth", AuthHandler)

	http.ListenAndServe(":8080", router)
}
