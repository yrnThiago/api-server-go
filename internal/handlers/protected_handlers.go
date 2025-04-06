package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/yrnThiago/api-server-go/internal/keys"
)

type Response struct {
	Message string `json:"message"`
}
type ProtectedHandler struct{}

func NewProtectedHandlers() *ProtectedHandler {
	return &ProtectedHandler{}
}

func (h *ProtectedHandler) TestCtx(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	userID, ok := ctx.Value(keys.UserIDKey).(string)
	if !ok {
		http.Error(w, "userID not found in context", http.StatusUnauthorized)
		return
	}
	fmt.Println("user id: ", userID)

	err := json.NewEncoder(w).Encode(Response{userID})
	if err != nil {
		fmt.Println("Error while parsing json")
		http.Error(w, "Parsing JSON", http.StatusInternalServerError)
	}
}
