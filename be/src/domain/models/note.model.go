package models

type Note struct {
	BaseModel
	Name        string `gorm:"not null"`
	Description string
}
