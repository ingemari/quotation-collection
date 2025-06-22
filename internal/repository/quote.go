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
		r.logger.Error("Failed to create quote", "err", err)
		return model.Quote{}, err
	}

	createdQuote := entities.EntityToQuote(ent)
	return createdQuote, nil
}

func (r *QuoteRepository) GetAllQuotes(ctx context.Context) ([]model.Quote, error) {
	query := `SELECT id, author, quote, created_at FROM quotes`

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		r.logger.Error("GetAllQuotes: query failed", "err", err)
		return nil, err
	}
	defer rows.Close()

	var quotes []model.Quote

	for rows.Next() {
		var e entities.Quote
		if err := rows.Scan(&e.ID, &e.Author, &e.Quote, &e.CreatedAt); err != nil {
			r.logger.Error("GetAllQuotes: scan failed", "err", err)
			continue
		}
		quotes = append(quotes, entities.EntityToQuote(e))
	}

	return quotes, nil
}

func (r *QuoteRepository) GetRandomQuote(ctx context.Context) (model.Quote, error) {
	var ent entities.Quote

	query := `
		SELECT id, author, quote, created_at
		FROM quotes
		ORDER BY RANDOM()
		LIMIT 1
	`

	err := r.db.QueryRow(ctx, query).Scan(&ent.ID, &ent.Author, &ent.Quote, &ent.CreatedAt)
	if err != nil {
		r.logger.Error("GetRandomQuote failed", "err", err)
		return model.Quote{}, err
	}

	quote := entities.EntityToQuote(ent)

	return quote, nil
}

func (r *QuoteRepository) GetQuotesByAuthor(ctx context.Context, author string) ([]model.Quote, error) {
	query := `
		SELECT id, author, quote, created_at
		FROM quotes
		WHERE author = $1
	`

	rows, err := r.db.Query(ctx, query, author)
	if err != nil {
		r.logger.Error("GetQuotesByAuthor query failed", slog.Any("err", err))
		return nil, err
	}
	defer rows.Close()

	var quotes []model.Quote
	for rows.Next() {
		var e entities.Quote
		if err := rows.Scan(&e.ID, &e.Author, &e.Quote, &e.CreatedAt); err != nil {
			r.logger.Error("Scan failed in GetQuotesByAuthor", "err", err)
			continue
		}
		quotes = append(quotes, entities.EntityToQuote(e))
	}

	return quotes, nil
}
