package queue

import (
	"os"

	"github.com/hibiken/asynq"
	"github.com/sirupsen/logrus"
)

var (
	Client *asynq.Client
)

// InitClient инициализирует клиент очереди задач
func InitClient() error {
	redisURL := os.Getenv("REDIS_URL")
	if redisURL == "" {
		redisURL = "localhost:6379"
	}

	redisPassword := os.Getenv("REDIS_PASSWORD")

	Client = asynq.NewClient(asynq.RedisClientOpt{
		Addr:     redisURL,
		Password: redisPassword,
	})

	logrus.Info("Task queue client initialized")
	return nil
}

// Close закрывает клиент очереди
func Close() error {
	if Client != nil {
		return Client.Close()
	}
	return nil
}

// IsEnabled проверяет, включена ли очередь задач
func IsEnabled() bool {
	return Client != nil
}
