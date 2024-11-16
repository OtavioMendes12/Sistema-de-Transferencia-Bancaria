package routes

import (
	"bank-transfer-system/internal/app/handlers"
	"bank-transfer-system/internal/infra/http"
	"github.com/gorilla/mux"
)

func SetupRoutes(authHandler *handlers.AuthHandler, userHandler *handlers.UserHandler, jwtSecret string) *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/login", authHandler.Login).Methods("POST")

	// Rotas protegidas
	protected := router.NewRoute().Subrouter()
	protected.Use(http.JWTMiddleware(jwtSecret))

	protected.HandleFunc("/users", userHandler.GetUsers).Methods("GET")
	protected.HandleFunc("/users/{id}/transfers", userHandler.GetTransfers).Methods("GET")
	protected.HandleFunc("/transfer", userHandler.Transfer).Methods("POST")

	return router
}
