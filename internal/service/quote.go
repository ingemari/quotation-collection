package service

import (
	"context"
	"log/slog"
	"quotation-collection/internal/model"
)

type QuoteRepository interface {
	CreateQuote(ctx context.Context, quote model.Quote) (model.Quote, error)
}

type QuoteService struct {
	quoteRepo QuoteRepository
	logger    *slog.Logger
}

func NewQuoteService(pr QuoteRepository, logger *slog.Logger) *QuoteService {
	return &QuoteService{quoteRepo: pr, logger: logger}
}

func (s *QuoteService) CreateQuote(ctx context.Context, quote model.Quote) (model.Quote, error) {
	return s.quoteRepo.CreateQuote(ctx, quote)
}
