package queue

import (
	"os"

	"github.com/hibiken/asynq"
	"github.com/sirupsen/logrus"
)

var (
	Server *asynq.Server
)

// StartWorker запускает worker для обработки задач
func StartWorker() error {
	redisURL := os.Getenv("REDIS_URL")
	if redisURL == "" {
		redisURL = "localhost:6379"
	}

	redisPassword := os.Getenv("REDIS_PASSWORD")

	Server = asynq.NewServer(
		asynq.RedisClientOpt{
			Addr:     redisURL,
			Password: redisPassword,
		},
		asynq.Config{
			Concurrency: 10, // Количество одновременно обрабатываемых задач
			Queues: map[string]int{
				"critical": 6, // Приоритет 6
				"default":  3, // Приоритет 3
				"low":      1, // Приоритет 1
			},
		},
	)

	// Регистрация обработчиков задач
	mux := asynq.NewServeMux()
	mux.HandleFunc(TypeEmailDelivery, HandleEmailDeliveryTask)
	mux.HandleFunc(TypeImageResize, HandleImageResizeTask)
	// Добавьте другие обработчики здесь

	logrus.Info("Starting task queue worker...")
	
	// Запуск worker в отдельной горутине
	go func() {
		if err := Server.Run(mux); err != nil {
			logrus.Fatalf("Could not run task queue worker: %v", err)
		}
	}()

	logrus.Info("Task queue worker started")
	return nil
}

// StopWorker останавливает worker
func StopWorker() {
	if Server != nil {
		Server.Shutdown()
		logrus.Info("Task queue worker stopped")
	}
}
