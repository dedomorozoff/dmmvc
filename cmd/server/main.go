package main

import (
	"dmmvc/internal/cache"
	"dmmvc/internal/config"
	"dmmvc/internal/database"
	"dmmvc/internal/email"
	"dmmvc/internal/i18n"
	"dmmvc/internal/logger"
	"dmmvc/internal/models"
	"dmmvc/internal/queue"
	"dmmvc/internal/routes"
	"dmmvc/internal/upload"
	"log"
	"os"

	"github.com/joho/godotenv"
)

// @title DMMVC API
// @version 1.4.0
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

	// Инициализация конфигурации функций
	features := config.InitFeatures()
	logger.Log.Info("Feature configuration:")
	logger.Log.Infof("  WebSocket: %v", features.WebSocket)
	logger.Log.Infof("  Redis: %v", features.Redis)
	logger.Log.Infof("  Swagger: %v", features.Swagger)
	logger.Log.Infof("  File Upload: %v", features.FileUpload)
	logger.Log.Infof("  I18n: %v", features.I18n)
	logger.Log.Infof("  Queue: %v", features.Queue)
	logger.Log.Infof("  Email: %v", features.Email)

	// Подключение к базе данных
	database.Connect()
	
	// Миграция моделей
	database.Migrate(&models.User{})
	
	// Создание администратора по умолчанию
	database.SeedAdmin()
	
	// Создание демо пользователей (опционально)
	database.SeedDemoUsers()

	// Подключение к Redis (если включено)
	if features.Redis {
		if err := cache.Connect(); err != nil {
			logger.Log.Warn("Redis not available, caching disabled")
		} else {
			logger.Log.Info("Redis cache enabled")
		}
	} else {
		logger.Log.Info("Redis cache disabled")
	}

	// Инициализация клиента очереди задач (если включено)
	if features.Queue {
		if err := queue.InitClient(); err != nil {
			logger.Log.Warn("Task queue client initialization failed")
		} else {
			logger.Log.Info("Task queue enabled")
		}
	} else {
		logger.Log.Info("Task queue disabled")
	}

	// Инициализация email сервиса (если включено)
	if features.Email {
		if err := email.Init(); err != nil {
			logger.Log.Warn("Email service not configured")
		} else {
			logger.Log.Info("Email service enabled")
		}
	} else {
		logger.Log.Info("Email service disabled")
	}

	// Инициализация сервиса загрузки файлов (если включено)
	if features.FileUpload {
		upload.Init()
		logger.Log.Info("File upload enabled")
	} else {
		logger.Log.Info("File upload disabled")
	}

	// Инициализация i18n (если включено)
	if features.I18n {
		i18nInstance := i18n.GetInstance()
		if err := i18nInstance.LoadTranslations("locales"); err != nil {
			logger.Log.Warn("Failed to load translations: ", err)
		} else {
			defaultLocale := i18nInstance.GetDefaultLocale()
			availableLocales := i18nInstance.GetAvailableLocales()
			logger.Log.Info("Translations loaded successfully")
			logger.Log.Infof("Default locale: %s", defaultLocale)
			logger.Log.Infof("Available locales: %v", availableLocales)
		}
	} else {
		logger.Log.Info("I18n disabled")
	}

	// Запуск worker для обработки задач (если включено)
	if features.Queue && os.Getenv("QUEUE_WORKER_ENABLED") == "true" {
		if err := queue.StartWorker(); err != nil {
			logger.Log.Warn("Task queue worker failed to start")
		} else {
			logger.Log.Info("Task queue worker started")
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
