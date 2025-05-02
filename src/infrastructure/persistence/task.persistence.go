package persistence

import (
	"be/src/domain/models"
	"be/src/domain/repository"
	"context"
	"errors"

	"gorm.io/gorm"
)

type TaskRepositoryImpl struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) repository.TaskRepository {
	return &TaskRepositoryImpl{db: db}
}

func (r *TaskRepositoryImpl) GetTasks(ctx context.Context) ([]models.Task, error) {
	var tasks []models.Task
	result := r.db.WithContext(ctx).Find(&tasks)
	if result.Error != nil {
		return nil, result.Error
	}
	return tasks, nil
}

func (r *TaskRepositoryImpl) GetTask(ctx context.Context, id string) (*models.Task, error) {
	var task models.Task
	result := r.db.WithContext(ctx).First(&task, "id = ?", id)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	if result.Error != nil {
		return nil, result.Error
	}

	return &task, nil
}

func (r *TaskRepositoryImpl) CreateTask(ctx context.Context, task *models.Task) (*models.Task, error) {
	result := r.db.WithContext(ctx).Create(task)
	if result.Error != nil {
		return nil, result.Error
	}

	return task, nil
}

func (r *TaskRepositoryImpl) UpdateTask(ctx context.Context, project *models.Task) (*models.Task, error) {
	result := r.db.WithContext(ctx).Save(project)
	if result.Error != nil {
		return nil, result.Error
	}
	return project, nil
}

func (r *TaskRepositoryImpl) DeleteTask(ctx context.Context, id string) error {
	result := r.db.WithContext(ctx).Delete(&models.Task{}, "id = ?", id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
