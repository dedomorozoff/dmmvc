package main

import (
	"dmmvc/internal/database"
	"dmmvc/internal/logger"
	"dmmvc/internal/models"
	"dmmvc/internal/routes"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	// Загрузка переменных окружения
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using default/env vars")
	}

	// Инициализация логгера
	logger.Init()
	logger.Log.Info("Starting DMMVC server...")

	// Подключение к базе данных
	database.Connect()
	
	// Миграция моделей
	database.Migrate(&models.User{})
	
	// Создание администратора по умолчанию
	database.SeedAdmin()

	// Настройка роутов
	r := routes.SetupRouter()

	// Получение порта из переменных окружения
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Запуск сервера
	logger.Log.Info("Server started on http://localhost:" + port)
	r.Run(":" + port)
}
