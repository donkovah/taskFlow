package interfaces

import (
	"be/src/domain/models"
	"context"
)

type AuthService interface {
	Login(ctx context.Context, email, password string) (*models.LoginResponse, error)
	Register(ctx context.Context, user *models.User) (*models.User, error)
	Logout(ctx context.Context, token string) error
}
