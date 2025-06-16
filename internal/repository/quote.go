package repository

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"log/slog"
	"quotation-collection/internal/model"
	"quotation-collection/internal/repository/entities"
)

type QuoteRepository struct {
	db     *pgxpool.Pool
	logger *slog.Logger
}

func NewQuoteRepository(db *pgxpool.Pool, logger *slog.Logger) *QuoteRepository {
	return &QuoteRepository{db: db, logger: logger}
}

func (r *QuoteRepository) CreateQuote(ctx context.Context, quote model.Quote) (model.Quote, error) {
	ent := entities.QuoteToEntity(quote)

	query := `
	INSERT INTO quotes (author, quote)
	VALUES ($1, $2)
	RETURNING id, created_at
	`

	err := r.db.QueryRow(ctx, query, ent.Author, ent.Quote).Scan(&ent.ID, &ent.CreatedAt)
	if err != nil {
		r.logger.Error("Failed to create user (already exist)")
		return model.Quote{}, err
	}

	createdQuote := entities.EntityToQuote(ent)
	return createdQuote, nil
}
