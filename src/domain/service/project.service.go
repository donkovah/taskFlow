package service

import (
	"be/src/domain/models"
	"be/src/domain/repository"
	"context"
)

type ProjectService struct {
	repo repository.ProjectRepository
}

func NewProjectService(repo repository.ProjectRepository) *ProjectService {
	return &ProjectService{repo: repo}
}

func (s *ProjectService) GetProjects(ctx context.Context) ([]models.Project, error) {
	return s.repo.GetProjects(ctx)
}

func (s *ProjectService) GetProject(ctx context.Context, id string) (*models.Project, error) {
	return s.repo.GetProject(ctx, id)
}

func (s *ProjectService) CreateProject(ctx context.Context, project *models.Project) (*models.Project, error) {
	return s.repo.CreateProject(ctx, project)
}

func (s *ProjectService) UpdateProject(ctx context.Context, project *models.Project) (*models.Project, error) {
	return s.repo.UpdateProject(ctx, project)
}

func (s *ProjectService) DeleteProject(ctx context.Context, id string) error {
	return s.repo.DeleteProject(ctx, id)
}
