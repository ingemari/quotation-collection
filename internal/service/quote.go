package service

import (
	"context"
	"log/slog"
	"quotation-collection/internal/model"
)

type QuoteRepository interface {
	CreateQuote(ctx context.Context, quote model.Quote) (model.Quote, error)
	GetAllQuotes(ctx context.Context) ([]model.Quote, error)
	GetRandomQuote(ctx context.Context) (model.Quote, error)
	GetQuotesByAuthor(ctx context.Context, author string) ([]model.Quote, error)
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

func (s *QuoteService) GetAllQuotes(ctx context.Context) ([]model.Quote, error) {
	return s.quoteRepo.GetAllQuotes(ctx)
}

func (s *QuoteService) GetRandomQuote(ctx context.Context) (model.Quote, error) {
	return s.quoteRepo.GetRandomQuote(ctx)
}

func (s *QuoteService) GetQuotesByAuthor(ctx context.Context, author string) ([]model.Quote, error) {
	return s.quoteRepo.GetQuotesByAuthor(ctx, author)
}
