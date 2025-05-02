package persistence

import (
	"be/src/domain/models"
	"be/src/domain/repository"
	"context"
	"errors"

	"gorm.io/gorm"
)

type NoteRepositoryImpl struct {
	db *gorm.DB
}

func NewNoteRepository(db *gorm.DB) repository.NoteRepository {
	return &NoteRepositoryImpl{db: db}
}

func (r *NoteRepositoryImpl) GetNotes(ctx context.Context) ([]models.Note, error) {
	var projects []models.Note
	result := r.db.WithContext(ctx).Find(&projects)
	if result.Error != nil {
		return nil, result.Error
	}
	return projects, nil
}

func (r *NoteRepositoryImpl) GetNote(ctx context.Context, id string) (*models.Note, error) {
	var project models.Note
	result := r.db.WithContext(ctx).First(&project, id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	if result.Error != nil {
		return nil, result.Error
	}
	return &project, nil
}

func (r *NoteRepositoryImpl) CreateNote(ctx context.Context, project *models.Note) (*models.Note, error) {
	result := r.db.WithContext(ctx).Create(project)
	if result.Error != nil {
		return nil, result.Error
	}
	return project, nil
}

func (r *NoteRepositoryImpl) UpdateNote(ctx context.Context, project *models.Note) (*models.Note, error) {
	result := r.db.WithContext(ctx).Save(project)
	if result.Error != nil {
		return nil, result.Error
	}
	return project, nil
}

func (r *NoteRepositoryImpl) DeleteNote(ctx context.Context, id string) error {
	result := r.db.WithContext(ctx).Delete(&models.Note{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
