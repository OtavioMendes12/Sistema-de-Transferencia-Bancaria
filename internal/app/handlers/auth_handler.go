package handlers

import (
	"bank-transfer-system/internal/app/config"
	"encoding/json"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type AuthHandler struct {
	Config config.Config
}

type LoginRequest struct {
	UserID string `json:"user_id"`
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var loginReq LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&loginReq); err != nil {
		http.Error(w, "Dados inválidos", http.StatusBadRequest)
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": loginReq.UserID,
		"exp":     time.Now().Add(time.Hour * 1).Unix(),
	})

	tokenString, err := token.SignedString([]byte(h.Config.JWTSecret))
	if err != nil {
		http.Error(w, "Erro ao gerar token", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"token": tokenString})
}
