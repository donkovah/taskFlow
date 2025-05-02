package models

import "github.com/google/uuid"

type Comment struct {
	BaseModel
	Name       string    `gorm:"not null"`
	UserID     uuid.UUID `gorm:"not null"`
	TimelineID int
	Timeline   Timeline
	User       User
}
