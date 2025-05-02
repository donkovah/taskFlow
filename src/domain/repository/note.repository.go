package repository

import (
	"be/src/domain/models"
	"context"
)

type NoteRepository interface {
	GetNotes(ctx context.Context) ([]models.Note, error)
	GetNote(ctx context.Context, id string) (*models.Note, error)
	CreateNote(ctx context.Context, note *models.Note) (*models.Note, error)
	UpdateNote(ctx context.Context, note *models.Note) (*models.Note, error)
	DeleteNote(ctx context.Context, id string) error
}
