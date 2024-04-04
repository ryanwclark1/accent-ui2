package auth

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type AuthClient struct {
	BaseURL string
}

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type AuthResponse struct {
	IsValid bool   `json:"isValid"`
	Message string `json:"message"`
}

func NewAuthClient(baseURL string) *AuthClient {
	return &AuthClient{BaseURL: baseURL}
}

func (c *AuthClient) Authenticate(credentials *Credentials) (*AuthResponse, error) {
	reqBody, err := json.Marshal(credentials)
	if err != nil {
		return nil, err
	}

	resp, err := http.Post(c.BaseURL+"api/auth/0.1/validate", "application/json", bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var authResp AuthResponse
	if err := json.NewDecoder(resp.Body).Decode(&authResp); err != nil {
		return nil, err
	}

	return &authResp, nil
}
