package repository

import (
	"be/src/domain/models"
	"context"
)

type ProjectRepository interface {
	GetProjects(ctx context.Context) ([]models.Project, error)
	GetProject(ctx context.Context, id string) (*models.Project, error)
	CreateProject(ctx context.Context, project *models.Project) (*models.Project, error)
	UpdateProject(ctx context.Context, project *models.Project) (*models.Project, error)
	DeleteProject(ctx context.Context, id string) error
}
