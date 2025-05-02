package models

import "github.com/google/uuid"

type Note struct {
	BaseModel
	Name        string `gorm:"not null"`
	Description string
	UserID      uuid.UUID `gorm:"not null"`
	User        User
}
