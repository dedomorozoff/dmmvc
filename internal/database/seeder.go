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

// SeedDemoUsers создает демо пользователей на английском и русском
func SeedDemoUsers() {
	var count int64
	DB.Model(&models.User{}).Count(&count)

	// Создаем демо пользователей только если есть только админ
	if count <= 1 {
		demoUsers := []struct {
			username string
			email    string
			password string
			role     string
		}{
			{"john_doe", "john@example.com", "password123", "user"},
			{"jane_smith", "jane@example.com", "password123", "user"},
			{"bob_johnson", "bob@example.com", "password123", "user"},
			{"ivan_ivanov", "ivan@example.ru", "password123", "user"},
			{"maria_petrova", "maria@example.ru", "password123", "user"},
			{"alexey_sidorov", "alexey@example.ru", "password123", "user"},
		}

		for _, demo := range demoUsers {
			hashedPassword, err := bcrypt.GenerateFromPassword([]byte(demo.password), bcrypt.DefaultCost)
			if err != nil {
				log.Printf("Failed to hash password for %s: %v", demo.username, err)
				continue
			}

			user := models.User{
				Username: demo.username,
				Email:    demo.email,
				Password: string(hashedPassword),
				Role:     demo.role,
			}

			if err := DB.Create(&user).Error; err != nil {
				log.Printf("Failed to create demo user %s: %v", demo.username, err)
			} else {
				log.Printf("Demo user created: %s (password: %s)", demo.username, demo.password)
			}
		}
	}
}
