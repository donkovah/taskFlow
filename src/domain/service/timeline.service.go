package service

import (
	"be/src/domain/models"
	"be/src/domain/repository"
	"context"
)

type TimelineService struct {
	repo repository.TimelineRepository
}

func NewTimelineService(repo repository.TimelineRepository) *TimelineService {
	return &TimelineService{repo: repo}
}

func (s *TimelineService) GetTimelines(ctx context.Context) ([]models.Timeline, error) {
	return s.repo.GetTimelines(ctx)
}

func (s *TimelineService) GetTimeline(ctx context.Context, id string) (*models.Timeline, error) {
	return s.repo.GetTimeline(ctx, id)
}

func (s *TimelineService) CreateTimeline(ctx context.Context, project *models.Timeline) (*models.Timeline, error) {
	return s.repo.CreateTimeline(ctx, project)
}

func (s *TimelineService) UpdateTimeline(ctx context.Context, project *models.Timeline) (*models.Timeline, error) {
	return s.repo.UpdateTimeline(ctx, project)
}

func (s *TimelineService) DeleteTimeline(ctx context.Context, id string) error {
	return s.repo.DeleteTimeline(ctx, id)
}
