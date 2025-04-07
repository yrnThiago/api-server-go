package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/yrnThiago/api-server-go/internal/config"
	"github.com/yrnThiago/api-server-go/internal/utils"
)

type AuthHandler struct{}

func NewAuthHandlers() *AuthHandler {
	return &AuthHandler{}
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	userAuthorization, err := utils.GenerateJWT()
	if err != nil {
		fmt.Println("Error creating a new tokens", err)
	}
	cookie := http.Cookie{}
	cookie.Name = config.Env.COOKIE_NAME
	cookie.Value = userAuthorization
	cookie.Expires = time.Now().Add(365 * 24 * time.Hour)
	cookie.Secure = false
	cookie.HttpOnly = true
	cookie.Path = "/"

	http.SetCookie(w, &cookie)

	err = json.NewEncoder(w).Encode(TestResponse{"user logged in"})
	if err != nil {
		fmt.Println("Error while parsing json")
		http.Error(w, "Parsing JSON", http.StatusInternalServerError)
	}
}
