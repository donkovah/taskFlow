package service

import (
	"be/src/domain/models"
	"be/src/domain/repository"
	"context"
)

type TaskService struct {
	repo repository.TaskRepository
}

func NewTaskService(repo repository.TaskRepository) *TaskService {
	return &TaskService{repo: repo}
}

func (s *TaskService) GetTasks(ctx context.Context) ([]models.Task, error) {
	return s.repo.GetTasks(ctx)
}

func (s *TaskService) GetTask(ctx context.Context, id string) (*models.Task, error) {
	return s.repo.GetTask(ctx, id)
}

func (s *TaskService) CreateTask(ctx context.Context, project *models.Task) (*models.Task, error) {
	return s.repo.CreateTask(ctx, project)
}

func (s *TaskService) UpdateTask(ctx context.Context, project *models.Task) (*models.Task, error) {
	return s.repo.UpdateTask(ctx, project)
}

func (s *TaskService) DeleteTask(ctx context.Context, id string) error {
	return s.repo.DeleteTask(ctx, id)
}
