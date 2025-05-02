package persistence

import (
	"be/src/domain/models"
	"be/src/domain/repository"
	"context"
	"errors"

	"gorm.io/gorm"
)

type CommentRepositoryImpl struct {
	db *gorm.DB
}

func NewCommentRepository(db *gorm.DB) repository.CommentRepository {
	return &CommentRepositoryImpl{db: db}
}

func (r *CommentRepositoryImpl) GetComments(ctx context.Context) ([]models.Comment, error) {
	var comments []models.Comment
	result := r.db.WithContext(ctx).Find(&comments)
	if result.Error != nil {
		return nil, result.Error
	}
	return comments, nil
}

func (r *CommentRepositoryImpl) GetComment(ctx context.Context, id string) (*models.Comment, error) {
	var comment models.Comment
	result := r.db.WithContext(ctx).First(&comment, id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	if result.Error != nil {
		return nil, result.Error
	}
	return &comment, nil
}

func (r *CommentRepositoryImpl) CreateComment(ctx context.Context, comment *models.Comment) (*models.Comment, error) {
	result := r.db.WithContext(ctx).Create(comment)
	if result.Error != nil {
		return nil, result.Error
	}
	return comment, nil
}

func (r *CommentRepositoryImpl) UpdateComment(ctx context.Context, comment *models.Comment) (*models.Comment, error) {
	result := r.db.WithContext(ctx).Save(comment)
	if result.Error != nil {
		return nil, result.Error
	}
	return comment, nil
}

func (r *CommentRepositoryImpl) DeleteComment(ctx context.Context, id string) error {
	result := r.db.WithContext(ctx).Delete(&models.Comment{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
