package models

import (
	"gorm.io/gorm"
)

// User модель пользователя
type User struct {
	gorm.Model
	Username string `gorm:"uniqueIndex;not null" json:"username"`
	Email    string `gorm:"uniqueIndex;not null" json:"email"`
	Password string `gorm:"not null" json:"-"`
	Role     string `gorm:"default:'user'" json:"role"` // admin, user
}
