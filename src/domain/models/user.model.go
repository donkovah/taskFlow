package models

type User struct {
	BaseModel
	Username string `gorm:"unique_index;not null"`
	Email    string `gorm:"unique_index;not null"`
	Password string `gorm:"not null"`
	Role     string `json:"role"`
}
