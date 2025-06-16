package entities

import (
	"github.com/google/uuid"
	"quotation-collection/internal/model"
	"time"
)

type Quote struct {
	ID        uuid.UUID `db:"id"`
	Author    string    `db:"author"`
	Quote     string    `db:"quote"`
	CreatedAt time.Time `db:"created_at"`
}

func QuoteToEntity(q model.Quote) Quote {
	return Quote{
		ID:        q.ID,
		Author:    q.Author,
		Quote:     q.Quote,
		CreatedAt: q.CreatedAt,
	}
}

func EntityToQuote(e Quote) model.Quote {
	return model.Quote{
		ID:        e.ID,
		Author:    e.Author,
		Quote:     e.Quote,
		CreatedAt: e.CreatedAt,
	}
}
