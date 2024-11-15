package main

import (
	"bank-transfer-system/internal/app/config"
	"bank-transfer-system/internal/app/handlers"
	"bank-transfer-system/internal/app/routes"
	"bank-transfer-system/internal/core/user"
	"bank-transfer-system/internal/infra/db"
	"bank-transfer-system/internal/infra/http"
)

func main() {
	// Carregar configurações
	cfg := config.LoadConfig()

	// Conectar ao banco de dados MongoDB
	db.ConnectDatabase(cfg.MongoURI)

	userRepo := db.NewUserRepository(cfg.Database, cfg.CollectionName)
	userService := user.NewService(userRepo)
	userHandler := &handlers.UserHandler{Service: userService}

	// Configurar as rotas
	router := routes.SetupRoutes(userHandler)

	// Iniciar o servidor
	http.StartServer(router, cfg.ServerPort)
}
