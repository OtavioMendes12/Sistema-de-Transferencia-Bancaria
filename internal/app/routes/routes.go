package routes

import (
	"bank-transfer-system/internal/app/handlers"
	"github.com/gorilla/mux"
)

func SetupRoutes(userHandler *handlers.UserHandler) *mux.Router {
	router := mux.NewRouter()

	// Rotas relacionadas a usuários
	router.HandleFunc("/users", userHandler.GetUsers).Methods("GET")
	router.HandleFunc("/users", userHandler.CreateUser).Methods("POST")
	router.HandleFunc("/transfer", userHandler.Transfer).Methods("POST")

	return router
}
