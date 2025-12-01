package controllers

import (
	"dmmvc/internal/queue"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/hibiken/asynq"
)

// EnqueueEmailTask добавляет задачу отправки email в очередь
// @Summary Добавить задачу отправки email
// @Description Добавляет задачу отправки email в очередь для фоновой обработки
// @Tags queue
// @Accept json
// @Produce json
// @Param email body EmailTaskRequest true "Email data"
// @Success 200 {object} APIResponse
// @Failure 400 {object} APIResponse
// @Router /api/queue/email [post]
// @Security SessionAuth
func EnqueueEmailTask(c *gin.Context) {
	var req EmailTaskRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, APIResponse{
			Success: false,
			Error:   err.Error(),
		})
		return
	}

	// Создание задачи
	task, err := queue.NewEmailDeliveryTask(req.To, req.Subject, req.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, APIResponse{
			Success: false,
			Error:   "Failed to create task",
		})
		return
	}

	// Добавление в очередь
	if err := queue.EnqueueTask(task, asynq.Queue("default")); err != nil {
		c.JSON(http.StatusInternalServerError, APIResponse{
			Success: false,
			Error:   "Failed to enqueue task",
		})
		return
	}

	c.JSON(http.StatusOK, APIResponse{
		Success: true,
		Message: "Email task enqueued successfully",
	})
}

// EnqueueDelayedTask добавляет задачу с задержкой
// @Summary Добавить отложенную задачу
// @Description Добавляет задачу отправки email с задержкой
// @Tags queue
// @Accept json
// @Produce json
// @Param email body EmailTaskRequest true "Email data"
// @Param delay query int false "Delay in seconds" default(60)
// @Success 200 {object} APIResponse
// @Failure 400 {object} APIResponse
// @Router /api/queue/email/delayed [post]
// @Security SessionAuth
func EnqueueDelayedTask(c *gin.Context) {
	var req EmailTaskRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, APIResponse{
			Success: false,
			Error:   err.Error(),
		})
		return
	}

	delay := 60 // default 60 seconds
	if d := c.Query("delay"); d != "" {
		// Parse delay from query
	}

	task, err := queue.NewEmailDeliveryTask(req.To, req.Subject, req.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, APIResponse{
			Success: false,
			Error:   "Failed to create task",
		})
		return
	}

	if err := queue.EnqueueTaskIn(task, time.Duration(delay)*time.Second); err != nil {
		c.JSON(http.StatusInternalServerError, APIResponse{
			Success: false,
			Error:   "Failed to enqueue task",
		})
		return
	}

	c.JSON(http.StatusOK, APIResponse{
		Success: true,
		Message: "Delayed email task enqueued successfully",
	})
}

// EnqueueImageTask добавляет задачу обработки изображения
// @Summary Добавить задачу обработки изображения
// @Description Добавляет задачу изменения размера изображения в очередь
// @Tags queue
// @Accept json
// @Produce json
// @Param image body ImageTaskRequest true "Image data"
// @Success 200 {object} APIResponse
// @Failure 400 {object} APIResponse
// @Router /api/queue/image [post]
// @Security SessionAuth
func EnqueueImageTask(c *gin.Context) {
	var req ImageTaskRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, APIResponse{
			Success: false,
			Error:   err.Error(),
		})
		return
	}

	task, err := queue.NewImageResizeTask(req.SourcePath, req.TargetPath, req.Width, req.Height)
	if err != nil {
		c.JSON(http.StatusInternalServerError, APIResponse{
			Success: false,
			Error:   "Failed to create task",
		})
		return
	}

	if err := queue.EnqueueTask(task, asynq.Queue("default")); err != nil {
		c.JSON(http.StatusInternalServerError, APIResponse{
			Success: false,
			Error:   "Failed to enqueue task",
		})
		return
	}

	c.JSON(http.StatusOK, APIResponse{
		Success: true,
		Message: "Image task enqueued successfully",
	})
}

// QueueStats возвращает статистику очереди
// @Summary Статистика очереди
// @Description Получить статистику очереди задач
// @Tags queue
// @Accept json
// @Produce json
// @Success 200 {object} APIResponse
// @Router /api/queue/stats [get]
// @Security SessionAuth
func QueueStats(c *gin.Context) {
	if !queue.IsEnabled() {
		c.JSON(http.StatusOK, APIResponse{
			Success: false,
			Message: "Task queue is disabled",
		})
		return
	}

	c.JSON(http.StatusOK, APIResponse{
		Success: true,
		Message: "Task queue is enabled",
		Data: gin.H{
			"status": "running",
		},
	})
}

// EmailTaskRequest структура запроса для email задачи
type EmailTaskRequest struct {
	To      string `json:"to" binding:"required,email" example:"user@example.com"`
	Subject string `json:"subject" binding:"required" example:"Welcome!"`
	Body    string `json:"body" binding:"required" example:"Welcome to our service!"`
}

// ImageTaskRequest структура запроса для задачи обработки изображения
type ImageTaskRequest struct {
	SourcePath string `json:"source_path" binding:"required" example:"/uploads/image.jpg"`
	TargetPath string `json:"target_path" binding:"required" example:"/uploads/thumb.jpg"`
	Width      int    `json:"width" binding:"required" example:"300"`
	Height     int    `json:"height" binding:"required" example:"200"`
}
