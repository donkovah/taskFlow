package repository

import (
	"be/src/domain/models"
	"context"
)

type TaskRepository interface {
	GetTasks(ctx context.Context) ([]models.Task, error)
	GetTask(ctx context.Context, id string) (*models.Task, error)
	CreateTask(ctx context.Context, task *models.Task) (*models.Task, error)
	UpdateTask(ctx context.Context, task *models.Task) (*models.Task, error)
	DeleteTask(ctx context.Context, id string) error
}
