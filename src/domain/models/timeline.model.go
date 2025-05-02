package models

import "github.com/google/uuid"

type Timeline struct {
	BaseModel
	Name        string `gorm:"not null"`
	Description string
	UserID      uuid.UUID `gorm:"not null"`
	Comment     []Comment
	User        User
}
