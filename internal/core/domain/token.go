package domain

import (
	"github.com/google/uuid"
)

type Token struct {
	ID        uuid.UUID
	UserID    uint64
	Token     string
	CreatedAt string
	UpdatedAt string
}
