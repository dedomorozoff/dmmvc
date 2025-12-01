package queue

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/hibiken/asynq"
	"github.com/sirupsen/logrus"
)

// Task types
const (
	TypeEmailDelivery  = "email:deliver"
	TypeImageResize    = "image:resize"
	TypeDataExport     = "data:export"
	TypeNotification   = "notification:send"
	TypeCleanup        = "cleanup:old_data"
)

// EmailDeliveryPayload структура для задачи отправки email
type EmailDeliveryPayload struct {
	To      string `json:"to"`
	Subject string `json:"subject"`
	Body    string `json:"body"`
}

// ImageResizePayload структура для задачи изменения размера изображения
type ImageResizePayload struct {
	SourcePath string `json:"source_path"`
	TargetPath string `json:"target_path"`
	Width      int    `json:"width"`
	Height     int    `json:"height"`
}

// NewEmailDeliveryTask создает новую задачу отправки email
func NewEmailDeliveryTask(to, subject, body string) (*asynq.Task, error) {
	payload, err := json.Marshal(EmailDeliveryPayload{
		To:      to,
		Subject: subject,
		Body:    body,
	})
	if err != nil {
		return nil, err
	}
	return asynq.NewTask(TypeEmailDelivery, payload), nil
}

// NewImageResizeTask создает новую задачу изменения размера изображения
func NewImageResizeTask(sourcePath, targetPath string, width, height int) (*asynq.Task, error) {
	payload, err := json.Marshal(ImageResizePayload{
		SourcePath: sourcePath,
		TargetPath: targetPath,
		Width:      width,
		Height:     height,
	})
	if err != nil {
		return nil, err
	}
	return asynq.NewTask(TypeImageResize, payload), nil
}

// EnqueueTask добавляет задачу в очередь
func EnqueueTask(task *asynq.Task, opts ...asynq.Option) error {
	if !IsEnabled() {
		logrus.Warn("Task queue is disabled, task not enqueued")
		return nil
	}

	info, err := Client.Enqueue(task, opts...)
	if err != nil {
		return err
	}

	logrus.Infof("Task enqueued: id=%s queue=%s", info.ID, info.Queue)
	return nil
}

// EnqueueTaskIn добавляет задачу в очередь с задержкой
func EnqueueTaskIn(task *asynq.Task, delay time.Duration, opts ...asynq.Option) error {
	if !IsEnabled() {
		logrus.Warn("Task queue is disabled, task not enqueued")
		return nil
	}

	info, err := Client.Enqueue(task, append(opts, asynq.ProcessIn(delay))...)
	if err != nil {
		return err
	}

	logrus.Infof("Task scheduled: id=%s queue=%s delay=%v", info.ID, info.Queue, delay)
	return nil
}

// EnqueueTaskAt добавляет задачу в очередь на определенное время
func EnqueueTaskAt(task *asynq.Task, processAt time.Time, opts ...asynq.Option) error {
	if !IsEnabled() {
		logrus.Warn("Task queue is disabled, task not enqueued")
		return nil
	}

	info, err := Client.Enqueue(task, append(opts, asynq.ProcessAt(processAt))...)
	if err != nil {
		return err
	}

	logrus.Infof("Task scheduled: id=%s queue=%s at=%v", info.ID, info.Queue, processAt)
	return nil
}

// HandleEmailDeliveryTask обработчик задачи отправки email
func HandleEmailDeliveryTask(ctx context.Context, t *asynq.Task) error {
	var p EmailDeliveryPayload
	if err := json.Unmarshal(t.Payload(), &p); err != nil {
		return fmt.Errorf("json.Unmarshal failed: %v: %w", err, asynq.SkipRetry)
	}

	logrus.Infof("Sending email to %s: %s", p.To, p.Subject)
	
	// Импортируем email пакет для отправки
	// import "dmmvc/internal/email"
	// if err := email.Send(p.To, p.Subject, p.Body); err != nil {
	//     return err
	// }
	
	// Временная симуляция отправки (удалить после настройки SMTP)
	time.Sleep(2 * time.Second)
	
	logrus.Infof("Email sent successfully to %s", p.To)
	return nil
}

// HandleImageResizeTask обработчик задачи изменения размера изображения
func HandleImageResizeTask(ctx context.Context, t *asynq.Task) error {
	var p ImageResizePayload
	if err := json.Unmarshal(t.Payload(), &p); err != nil {
		return fmt.Errorf("json.Unmarshal failed: %v: %w", err, asynq.SkipRetry)
	}

	logrus.Infof("Resizing image: %s -> %s (%dx%d)", p.SourcePath, p.TargetPath, p.Width, p.Height)
	
	// TODO: Реализовать изменение размера изображения
	// Здесь должна быть логика обработки изображения
	
	time.Sleep(1 * time.Second) // Симуляция обработки
	
	logrus.Infof("Image resized successfully: %s", p.TargetPath)
	return nil
}
