package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type TestResponse struct {
	Message string `json:"message"`
}

type HealthHandler struct{}

func NewHealthHandlers() *HealthHandler {
	return &HealthHandler{}
}

func (h *HealthHandler) Ping(w http.ResponseWriter, r *http.Request) {
	err := json.NewEncoder(w).Encode(TestResponse{"pong"})
	if err != nil {
		fmt.Println("Error while parsing json")
		http.Error(w, "Parsing JSON", http.StatusInternalServerError)
	}
}
