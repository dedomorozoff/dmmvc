package config

import (
	"os"
	"strconv"
)

// Features содержит флаги включения/выключения функций фреймворка
type Features struct {
	WebSocket   bool
	Redis       bool
	Swagger     bool
	FileUpload  bool
	I18n        bool
	Queue       bool
	Email       bool
}

var features *Features

// InitFeatures инициализирует конфигурацию функций из переменных окружения
func InitFeatures() *Features {
	if features != nil {
		return features
	}

	features = &Features{
		WebSocket:  getBoolEnv("WEBSOCKET_ENABLE", false),
		Redis:      getBoolEnv("REDIS_ENABLE", false),
		Swagger:    getBoolEnv("SWAGGER_ENABLE", false),
		FileUpload: getBoolEnv("FILE_UPLOAD_ENABLE", false),
		I18n:       getBoolEnv("I18N_ENABLE", false),
		Queue:      getBoolEnv("QUEUE_ENABLE", false),
		Email:      getBoolEnv("EMAIL_ENABLE", false),
	}

	return features
}

// GetFeatures возвращает текущую конфигурацию функций
func GetFeatures() *Features {
	if features == nil {
		return InitFeatures()
	}
	return features
}

// getBoolEnv читает boolean значение из переменной окружения
func getBoolEnv(key string, defaultValue bool) bool {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}

	boolValue, err := strconv.ParseBool(value)
	if err != nil {
		return defaultValue
	}

	return boolValue
}
