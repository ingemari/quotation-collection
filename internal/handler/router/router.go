package router

import (
	"log/slog"
	"quotation-collection/internal/handler"
	"quotation-collection/internal/repository"
	"quotation-collection/internal/service"

	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v5/pgxpool"
)

func SetupRouter(db *pgxpool.Pool, logger *slog.Logger) *mux.Router {
	router := mux.NewRouter()

	//// repo
	quoteRepo := repository.NewQuoteRepository(db, logger)
	//// service
	quoteService := service.NewQuoteService(quoteRepo, logger)
	//// handler
	quoteHandler := handler.NewQuoteHandler(quoteService, logger)

	// open routes
	router.HandleFunc("/quotes", quoteHandler.HandleCreateQuote).Methods("POST")
	//router.HandleFunc("/quotes", quoteHandler.HandleRegister).Methods("GET")
	//router.HandleFunc("/quotes/random", quoteHandler.HandleRegister).Methods("GET")
	//router.HandleFunc("/quotes/{id}", quoteHandler.HandleRegister).Methods("DELETE")

	return router
}
