package repository

import (
	"be/src/domain/models"
	"context"
)

type UserRepository interface {
	GetUsers(ctx context.Context) ([]models.User, error)
	GetUser(ctx context.Context, id string) (*models.User, error)
	CreateUser(ctx context.Context, timeline *models.User) (*models.User, error)
	UpdateUser(ctx context.Context, timeline *models.User) (*models.User, error)
	DeleteUser(ctx context.Context, id string) error
	GetUserByEmail(ctx context.Context, email string) (*models.User, error)
}
