package persistence

import (
	"be/src/domain/models"
	"be/src/domain/repository"
	"context"
	"errors"

	"gorm.io/gorm"
)

type TimelineRepositoryImpl struct {
	db *gorm.DB
}

func NewTimelineRepository(db *gorm.DB) repository.TimelineRepository {
	return &TimelineRepositoryImpl{db: db}
}

func (r *TimelineRepositoryImpl) GetTimelines(ctx context.Context) ([]models.Timeline, error) {
	var timelines []models.Timeline
	result := r.db.WithContext(ctx).Find(&timelines)
	if result.Error != nil {
		return nil, result.Error
	}
	return timelines, nil
}

func (r *TimelineRepositoryImpl) GetTimeline(ctx context.Context, id string) (*models.Timeline, error) {
	var timeline models.Timeline
	result := r.db.WithContext(ctx).First(&timeline, id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	if result.Error != nil {
		return nil, result.Error
	}
	return &timeline, nil
}

func (r *TimelineRepositoryImpl) CreateTimeline(ctx context.Context, timeline *models.Timeline) (*models.Timeline, error) {
	result := r.db.WithContext(ctx).Create(timeline)
	if result.Error != nil {
		return nil, result.Error
	}
	return timeline, nil
}

func (r *TimelineRepositoryImpl) UpdateTimeline(ctx context.Context, timeline *models.Timeline) (*models.Timeline, error) {
	result := r.db.WithContext(ctx).Save(timeline)
	if result.Error != nil {
		return nil, result.Error
	}
	return timeline, nil
}

func (r *TimelineRepositoryImpl) DeleteTimeline(ctx context.Context, id string) error {
	result := r.db.WithContext(ctx).Delete(&models.Timeline{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
