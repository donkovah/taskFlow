package service

import (
	"be/src/domain/models"
	"be/src/domain/repository"
	"context"
)

type CommentService struct {
	repo repository.CommentRepository
}

func NewCommentService(repo repository.CommentRepository) *CommentService {
	return &CommentService{repo: repo}
}

func (s *CommentService) GetComments(ctx context.Context) ([]models.Comment, error) {
	return s.repo.GetComments(ctx)
}

func (s *CommentService) GetComment(ctx context.Context, id string) (*models.Comment, error) {
	return s.repo.GetComment(ctx, id)
}

func (s *CommentService) CreateComment(ctx context.Context, project *models.Comment) (*models.Comment, error) {
	return s.repo.CreateComment(ctx, project)
}

func (s *CommentService) UpdateComment(ctx context.Context, project *models.Comment) (*models.Comment, error) {
	return s.repo.UpdateComment(ctx, project)
}

func (s *CommentService) DeleteComment(ctx context.Context, id string) error {
	return s.repo.DeleteComment(ctx, id)
}
