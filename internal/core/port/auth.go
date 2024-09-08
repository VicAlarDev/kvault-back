package port

import (
	"context"

	"github.com/VicAlarDev/kvault-back/internal/core/domain"
)

type TokenService interface {
	GenerateToken(user *domain.User) (string, error)
	Validate(token string) (*domain.Token, error)
}

type AuthService interface {
	// Login authenticates a user by email and password and returns a token
	Login(ctx context.Context, email, password string) (string, error)
}
