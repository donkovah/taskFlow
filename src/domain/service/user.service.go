package service

import (
	"be/src/domain/models"
	"be/src/domain/repository"
	"context"
)

type UserService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) GetUsers(ctx context.Context) ([]models.User, error) {
	return s.repo.GetUsers(ctx)
}

func (s *UserService) GetUser(ctx context.Context, id string) (*models.User, error) {
	return s.repo.GetUser(ctx, id)
}

func (s *UserService) CreateUser(ctx context.Context, project *models.User) (*models.User, error) {
	return s.repo.CreateUser(ctx, project)
}

func (s *UserService) UpdateUser(ctx context.Context, project *models.User) (*models.User, error) {
	return s.repo.UpdateUser(ctx, project)
}

func (s *UserService) DeleteUser(ctx context.Context, id string) error {
	return s.repo.DeleteUser(ctx, id)
}
