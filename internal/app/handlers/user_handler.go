package handlers

import (
	"bank-transfer-system/internal/core/transfer"
	"bank-transfer-system/internal/core/user"
	"context"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

type UserHandler struct {
	Service         user.Service
	TransferService transfer.Service
}

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

	// Realizar a transferência
	err := h.Service.Transfer(context.Background(), transferDTO)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Salvar no histórico de transferências
	transferHistory := transfer.TransferHistory{
		FromID:   transferDTO.FromID,
		ToID:     transferDTO.ToID,
		Amount:   transferDTO.Amount,
		Currency: "BRL", // Por exemplo, podemos ajustar isso conforme necessário
	}

	err = h.TransferService.SaveTransfer(context.Background(), &transferHistory)
	if err != nil {
		http.Error(w, "Erro ao salvar histórico de transferência", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Transferência realizada com sucesso"))
}

func (h *UserHandler) GetTransfers(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID := vars["id"]

	if userID == "" {
		http.Error(w, "ID do usuário não fornecido", http.StatusBadRequest)
		return
	}

	transfers, err := h.TransferService.GetTransfersByUserID(context.Background(), userID)
	if err != nil {
		http.Error(w, "Erro ao buscar histórico de transferências", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(transfers)
}
