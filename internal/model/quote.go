package model

import (
	"time"

	"github.com/google/uuid"
)

type Quote struct {
	ID        uuid.UUID
	Author    string
	Quote     string
	CreatedAt time.Time
}
