package handler

import (
	"context"
	"encoding/json"
	"log/slog"
	"net/http"
	"quotation-collection/internal/handler/dto"
	"quotation-collection/internal/handler/mapper"
	"quotation-collection/internal/model"
)

type QuoteService interface {
	CreateQuote(ctx context.Context, quote model.Quote) (model.Quote, error)
	GetAllQuotes(ctx context.Context) ([]model.Quote, error)
	GetRandomQuote(ctx context.Context) (model.Quote, error)
	GetQuotesByAuthor(ctx context.Context, author string) ([]model.Quote, error)
}

type QuoteHandler struct {
	quoteService QuoteService
	logger       *slog.Logger
}

func NewQuoteHandler(as QuoteService, logger *slog.Logger) *QuoteHandler {
	return &QuoteHandler{quoteService: as, logger: logger}
}

func (h *QuoteHandler) HandleCreateQuote(w http.ResponseWriter, r *http.Request) {
	var req dto.CreateQuoteReq
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Некорректный запрос", http.StatusBadRequest)
		h.logger.Error("Некорректный запрос", "err", err)
		return
	}

	quote := mapper.CreateReqToQuote(req)

	quote, err := h.quoteService.CreateQuote(r.Context(), quote)
	if err != nil {
		h.logger.Error("Failed create qoute", "err", err)
		http.Error(w, "Ошибка добавления цитаты", http.StatusInternalServerError)
		return
	}

	resp := mapper.QuoteToQuoteResp(quote)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		h.logger.Error("Ошибка при отправке ответа: ", "err", err)
	}
}

func (h *QuoteHandler) HandleGetAllQuotes(w http.ResponseWriter, r *http.Request) {
	quotes, err := h.quoteService.GetAllQuotes(r.Context())
	if err != nil {
		h.logger.Error("Failed to get all quotes", "err", err)
		http.Error(w, "Ошибка при получении цитат", http.StatusInternalServerError)
		return
	}

	resp := mapper.QuotesToListResp(quotes)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		h.logger.Error("Ошибка при отправке ответа", "err", err)
	}
}

func (h *QuoteHandler) HandleGetRandomQuote(w http.ResponseWriter, r *http.Request) {
	quote, err := h.quoteService.GetRandomQuote(r.Context())
	if err != nil {
		h.logger.Error("Failed to get random quote", "err", err)
		http.Error(w, "Цитаты не найдены", http.StatusNotFound)
		return
	}

	resp := mapper.QuoteToQuoteResp(quote)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		h.logger.Error("Ошибка при отправке ответа", "err", err)
	}
}

func (h *QuoteHandler) HandleGetQuotesByAuthor(w http.ResponseWriter, r *http.Request) {
	author := r.URL.Query().Get("author")
	//if author == "" {
	//	http.Error(w, "Параметр author обязателен", http.StatusBadRequest)
	//	return
	//}

	quotes, err := h.quoteService.GetQuotesByAuthor(r.Context(), author)
	if err != nil {
		h.logger.Error("Failed to get quotes by author", "err", err)
		http.Error(w, "Ошибка при получении цитат", http.StatusInternalServerError)
		return
	}

	resp := mapper.QuotesToListResp(quotes)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		h.logger.Error("Ошибка при отправке ответа", slog.Any("err", err))
	}
}
