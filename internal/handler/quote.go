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
		h.logger.Error("Некорректный запрос")
		return
	}

	quote := mapper.CreateReqToQuote(req)

	quote, err := h.quoteService.CreateQuote(r.Context(), quote)
	if err != nil {
		h.logger.Error("Failed create qoute")
		http.Error(w, "Ошибка добавления цитаты", http.StatusInternalServerError)
	}

	resp := mapper.QuoteToCreateResp(quote)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		h.logger.Error("Ошибка при отправке ответа: ", err)
	}
}
