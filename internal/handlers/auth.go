package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type AuthHandler struct{}

func NewAuthHandlers() *AuthHandler {
	return &AuthHandler{}
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	err := json.NewEncoder(w).Encode(TestResponse{"user logged in"})
	if err != nil {
		fmt.Println("Error while parsing json")
		http.Error(w, "Parsing JSON", http.StatusInternalServerError)
	}
}
