package models

import (
	"time"

	"github.com/google/uuid"
)

type Task struct {
	BaseModel
	Title       string     `gorm:"not null"`
	Description string     `gorm:"not null"`
	ProjectID   uuid.UUID  `gorm:"not null"`
	UserID      *uuid.UUID `gorm:"null"`
	Status      Status     `gorm:"type:varchar(20);not null"`
	CompletedAt time.Time  `gorm:"null"`
	Deadline    *time.Time
	Project     *Project
	User        *User
}
