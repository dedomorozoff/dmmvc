package models

import (
	"gorm.io/gorm"
)

// User модель пользователя
// @Description Модель пользователя системы
type User struct {
	gorm.Model
	Username string `gorm:"uniqueIndex;not null" json:"username" example:"admin"`
	Email    string `gorm:"uniqueIndex;not null" json:"email" example:"admin@example.com"`
	Password string `gorm:"not null" json:"-"`
	Role     string `gorm:"default:'user'" json:"role" example:"user"` // admin, user
}
