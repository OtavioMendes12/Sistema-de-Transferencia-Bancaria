package handlers

import (
	"bank-transfer-system/internal/core/user"
	"context"
	"encoding/json"
	"net/http"
)

// UserHandler estrutura para os controladores de usuários
type UserHandler struct {
	Service user.Service
}

// GetUsers retorna a lista de usuários
func (h *UserHandler) GetUsers(w http.ResponseWriter, r *http.Request) {
	users, err := h.Service.GetUsers()
	if err != nil {
		http.Error(w, "Erro ao buscar usuários", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

// CreateUser cria um novo usuário
func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var dto user.User
	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		http.Error(w, "Dados inválidos", http.StatusBadRequest)
		return
	}
	err := h.Service.AddUser(&dto)
	if err != nil {
		http.Error(w, "Erro ao criar usuário", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(dto)
}

func (h *UserHandler) Transfer(w http.ResponseWriter, r *http.Request) {
	var transferDTO user.TransferDTO
	if err := json.NewDecoder(r.Body).Decode(&transferDTO); err != nil {
		http.Error(w, "Dados inválidos", http.StatusBadRequest)
		return
	}

	err := h.Service.Transfer(context.Background(), transferDTO)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Transferência realizada com sucesso"))
}
