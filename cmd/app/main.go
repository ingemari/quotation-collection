package main

import (
	"log/slog"
	"quotation-collection/internal/config"
	"quotation-collection/internal/db"
	"quotation-collection/internal/handler/router"
	"quotation-collection/internal/middleware/logs"
	"quotation-collection/internal/server"
	"time"
)

func main() {
	logger := logs.SetupLogger()
	slog.SetDefault(logger)

	cfg := config.LoadConfig()
	logger.Info("Configuration loaded", "config", cfg)

	db := db.InitDB(config.MakeDSN(*cfg))
	defer db.Close()
	logger.Info("Database connection established")

	router := router.SetupRouter(db, logger)

	server.Run(
		logger,
		router,
		cfg.ServerPort,
		30*time.Second,
	)
}
