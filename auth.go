package main

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"html/template"
)

type LoginForm struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type AuthResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.html"))
}

func AuthHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var form LoginForm
	err := json.NewDecoder(r.Body).Decode(&form)
	if err != nil {
		log.Printf("Error decoding login form: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	data := bytes.NewBuffer(nil)
	json.NewEncoder(data).Encode(form)

	resp, err := http.Post("http://api.accentvoice.com/api/auth/0.1/login", "application/json", data)
	if err != nil {
		log.Printf("Error contacting auth service: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Error reading response body: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var authResp AuthResponse
	err = json.Unmarshal(body, &authResp)
	if err != nil {
		log.Printf("Error unmarshaling response: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if authResp.Success {
		tpl.ExecuteTemplate(w, "main_page.html", nil)
	} else {
		tpl.ExecuteTemplate(w, "login.html", authResp)
	}
}

func main() {
	router := httprouter.New()
	router.POST("/auth", AuthHandler)

	log.Fatal(http.ListenAndServe(":8080", router))
}
