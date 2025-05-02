package repository

import (
	"be/src/domain/models"
	"context"
)

type TimelineRepository interface {
	GetTimelines(ctx context.Context) ([]models.Timeline, error)
	GetTimeline(ctx context.Context, id string) (*models.Timeline, error)
	CreateTimeline(ctx context.Context, timeline *models.Timeline) (*models.Timeline, error)
	UpdateTimeline(ctx context.Context, timeline *models.Timeline) (*models.Timeline, error)
	DeleteTimeline(ctx context.Context, id string) error
}
