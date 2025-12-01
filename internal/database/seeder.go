package database

import (
	"dmmvc/internal/models"
	"log"

	"golang.org/x/crypto/bcrypt"
)

// SeedAdmin создает администратора по умолчанию, если его нет
func SeedAdmin() {
	var count int64
	DB.Model(&models.User{}).Count(&count)

	if count == 0 {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte("admin"), bcrypt.DefaultCost)
		if err != nil {
			log.Fatal("Failed to hash password:", err)
		}

		admin := models.User{
			Username: "admin",
			Email:    "admin@example.com",
			Password: string(hashedPassword),
			Role:     "admin",
		}

		if err := DB.Create(&admin).Error; err != nil {
			log.Fatal("Failed to create admin user:", err)
		}

		log.Println("Admin user created successfully (username: admin, password: admin)")
		log.Println("⚠️  Please change the default password!")
	}
}
