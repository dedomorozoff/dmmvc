package main

import (
	"dmmvc/internal/cache"
	"dmmvc/internal/database"
	"dmmvc/internal/email"
	"dmmvc/internal/logger"
	"dmmvc/internal/models"
	"dmmvc/internal/queue"
	"dmmvc/internal/routes"
	"log"
	"os"

	"github.com/joho/godotenv"
)

// @title DMMVC API
// @version 1.0
// @description Lightweight MVC Web Framework API
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.email support@dmmvc.com

// @license.name MIT
// @license.url https://opensource.org/licenses/MIT

// @host localhost:8080
// @BasePath /

// @securityDefinitions.apikey SessionAuth
// @in cookie
// @name session_token

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

	// Подключение к Redis (опционально)
	if err := cache.Connect(); err != nil {
		logger.Log.Warn("Redis not available, caching disabled")
	}

	// Инициализация клиента очереди задач (опционально)
	if err := queue.InitClient(); err != nil {
		logger.Log.Warn("Task queue client initialization failed")
	}

	// Инициализация email сервиса (опционально)
	if err := email.Init(); err != nil {
		logger.Log.Warn("Email service not configured")
	}

	// Запуск worker для обработки задач (опционально)
	if os.Getenv("QUEUE_WORKER_ENABLED") == "true" {
		if err := queue.StartWorker(); err != nil {
			logger.Log.Warn("Task queue worker failed to start")
		}
	}

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
