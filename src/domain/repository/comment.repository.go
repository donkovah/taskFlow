package repository

import (
	"be/src/domain/models"
	"context"
)

type CommentRepository interface {
	GetComments(ctx context.Context) ([]models.Comment, error)
	GetComment(ctx context.Context, id string) (*models.Comment, error)
	CreateComment(ctx context.Context, comment *models.Comment) (*models.Comment, error)
	UpdateComment(ctx context.Context, comment *models.Comment) (*models.Comment, error)
	DeleteComment(ctx context.Context, id string) error
}
