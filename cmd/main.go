package main

import (
	"bank-transfer-system/internal/app/config"
	"bank-transfer-system/internal/app/handlers"
	"bank-transfer-system/internal/app/routes"
	"bank-transfer-system/internal/core/currency"
	"bank-transfer-system/internal/core/transfer"
	"bank-transfer-system/internal/core/user"
	"bank-transfer-system/internal/infra/db"
	"bank-transfer-system/internal/infra/http"
)

func main() {

	cfg := config.LoadConfig()

	db.ConnectDatabase(cfg.MongoURI)

	userRepo := db.NewUserRepository(cfg.Database, cfg.CollectionName)
	currencyRepo := db.NewCurrencyRepository(cfg.Database, "CurrencyRates")
	currencyService := currency.NewService(currencyRepo)
	userService := user.NewService(userRepo, currencyService)
	transferRepo := db.NewTransferRepository(cfg.Database, "TransferHistory")
	transferService := transfer.NewService(transferRepo)
	userHandler := &handlers.UserHandler{
		Service:         userService,
		TransferService: transferService,
	}

	authHandler := &handlers.AuthHandler{Config: cfg}

	router := routes.SetupRoutes(authHandler, userHandler, cfg.JWTSecret)

	http.StartServer(router, cfg.ServerPort)
}
