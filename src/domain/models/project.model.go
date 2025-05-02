package models

import "github.com/google/uuid"

type Project struct {
	BaseModel
	Name        string     `json:"name" validate:"required,min=2,max=100" gorm:"not null"`
	Description string     `json:"description" validate:"max=500"`
	UserID      *uuid.UUID `gorm:"null"`
	Tasks       []Task     `gorm:"foreignKey:ProjectID;references:ID"`
	User        User       `gorm:"references:ID"`
}
