package service

import (
	"be/src/domain/models"
	"be/src/domain/repository"
	"context"
)

type NoteService struct {
	repo repository.NoteRepository
}

func NewNoteService(repo repository.NoteRepository) *NoteService {
	return &NoteService{repo: repo}
}

func (s *NoteService) GetNotes(ctx context.Context) ([]models.Note, error) {
	return s.repo.GetNotes(ctx)
}

func (s *NoteService) GetNote(ctx context.Context, id string) (*models.Note, error) {
	return s.repo.GetNote(ctx, id)
}

func (s *NoteService) CreateNote(ctx context.Context, project *models.Note) (*models.Note, error) {
	return s.repo.CreateNote(ctx, project)
}

func (s *NoteService) UpdateNote(ctx context.Context, project *models.Note) (*models.Note, error) {
	return s.repo.UpdateNote(ctx, project)
}

func (s *NoteService) DeleteNote(ctx context.Context, id string) error {
	return s.repo.DeleteNote(ctx, id)
}
