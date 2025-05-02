package controllers

import (
	"be/src/domain/models"
	"be/src/domain/service"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TimelineController struct {
	service *service.TimelineService
}

func NewTimelineController(service *service.TimelineService) *TimelineController {
	return &TimelineController{service: service}
}
func (tc *TimelineController) GetTimeline(c *gin.Context) {
	id := c.Param("id")
	timeline, err := tc.service.GetTimeline(context.Background(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get timeline"})
	}
	c.JSON(http.StatusOK, timeline)
}

func (tc TimelineController) GetTimelines(c *gin.Context) {
	timelines, err := tc.service.GetTimelines(context.Background())
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Failed to fetch timeline"})
	}
	c.JSON(http.StatusOK, timelines)
}

func (tc TimelineController) CreateTimeline(c *gin.Context) {
	var timeline models.Timeline
	if err := c.ShouldBindJSON(&timeline); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	createdTimeline, err := tc.service.CreateTimeline(context.Background(), &timeline)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create timeline"})
	}

	c.JSON(http.StatusOK, createdTimeline)
}

func (tc TimelineController) UpdateTimeline(c *gin.Context) {
	id := c.Param("id")
	var timelineBody *models.Timeline

	if err := c.ShouldBindJSON(&timelineBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	timeline, err := tc.service.GetTimeline(context.Background(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch timeline"})
	}
	timeline.Name = timelineBody.Name
	timeline.Description = timelineBody.Description

	updatedTimeline, err := tc.service.UpdateTimeline(context.Background(), timeline)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update timeline"})
	}

	c.JSON(http.StatusOK, updatedTimeline)
}

func (ts TimelineController) DeleteTimeline(c *gin.Context) {
	id := c.Param(("id"))
	err := ts.service.DeleteTimeline(context.Background(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete timeline"})
	}
	c.JSON(http.StatusNoContent, nil)
}
