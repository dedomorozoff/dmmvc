package controllers

import (
	"dmmvc/internal/email"
	"dmmvc/internal/queue"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hibiken/asynq"
)

// SendEmailDirect отправляет email напрямую (синхронно)
// @Summary Отправить email напрямую
// @Description Отправляет email синхронно (блокирует запрос)
// @Tags email
// @Accept json
// @Produce json
// @Param email body SendEmailRequest true "Email data"
// @Success 200 {object} APIResponse
// @Failure 400 {object} APIResponse
// @Router /api/email/send [post]
// @Security SessionAuth
func SendEmailDirect(c *gin.Context) {
	var req SendEmailRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, APIResponse{
			Success: false,
			Error:   err.Error(),
		})
		return
	}

	if !email.IsEnabled() {
		c.JSON(http.StatusServiceUnavailable, APIResponse{
			Success: false,
			Error:   "Email service is not configured",
		})
		return
	}

	// Отправка email
	if err := email.Send(req.To, req.Subject, req.Body); err != nil {
		c.JSON(http.StatusInternalServerError, APIResponse{
			Success: false,
			Error:   "Failed to send email: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, APIResponse{
		Success: true,
		Message: "Email sent successfully",
	})
}

// SendEmailAsync отправляет email асинхронно через очередь
// @Summary Отправить email асинхронно
// @Description Добавляет задачу отправки email в очередь для фоновой обработки
// @Tags email
// @Accept json
// @Produce json
// @Param email body SendEmailRequest true "Email data"
// @Success 200 {object} APIResponse
// @Failure 400 {object} APIResponse
// @Router /api/email/send/async [post]
// @Security SessionAuth
func SendEmailAsync(c *gin.Context) {
	var req SendEmailRequest
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
		Message: "Email queued for sending",
	})
}

// SendWelcomeEmail отправляет приветственное письмо
// @Summary Отправить приветственное письмо
// @Description Отправляет приветственное письмо новому пользователю
// @Tags email
// @Accept json
// @Produce json
// @Param email body WelcomeEmailRequest true "Email data"
// @Success 200 {object} APIResponse
// @Failure 400 {object} APIResponse
// @Router /api/email/welcome [post]
// @Security SessionAuth
func SendWelcomeEmail(c *gin.Context) {
	var req WelcomeEmailRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, APIResponse{
			Success: false,
			Error:   err.Error(),
		})
		return
	}

	if !email.IsEnabled() {
		c.JSON(http.StatusServiceUnavailable, APIResponse{
			Success: false,
			Error:   "Email service is not configured",
		})
		return
	}

	if err := email.WelcomeEmail(req.To, req.Username); err != nil {
		c.JSON(http.StatusInternalServerError, APIResponse{
			Success: false,
			Error:   "Failed to send welcome email",
		})
		return
	}

	c.JSON(http.StatusOK, APIResponse{
		Success: true,
		Message: "Welcome email sent successfully",
	})
}

// SendPasswordResetEmail отправляет письмо для сброса пароля
// @Summary Отправить письмо для сброса пароля
// @Description Отправляет письмо со ссылкой для сброса пароля
// @Tags email
// @Accept json
// @Produce json
// @Param email body PasswordResetEmailRequest true "Email data"
// @Success 200 {object} APIResponse
// @Failure 400 {object} APIResponse
// @Router /api/email/password-reset [post]
// @Security SessionAuth
func SendPasswordResetEmail(c *gin.Context) {
	var req PasswordResetEmailRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, APIResponse{
			Success: false,
			Error:   err.Error(),
		})
		return
	}

	if !email.IsEnabled() {
		c.JSON(http.StatusServiceUnavailable, APIResponse{
			Success: false,
			Error:   "Email service is not configured",
		})
		return
	}

	if err := email.PasswordResetEmail(req.To, req.ResetLink); err != nil {
		c.JSON(http.StatusInternalServerError, APIResponse{
			Success: false,
			Error:   "Failed to send password reset email",
		})
		return
	}

	c.JSON(http.StatusOK, APIResponse{
		Success: true,
		Message: "Password reset email sent successfully",
	})
}

// EmailStatus возвращает статус email сервиса
// @Summary Статус email сервиса
// @Description Проверяет, настроен ли email сервис
// @Tags email
// @Accept json
// @Produce json
// @Success 200 {object} APIResponse
// @Router /api/email/status [get]
// @Security SessionAuth
func EmailStatus(c *gin.Context) {
	enabled := email.IsEnabled()
	message := "Email service is disabled"
	if enabled {
		message = "Email service is enabled"
	}
	
	c.JSON(http.StatusOK, APIResponse{
		Success: enabled,
		Message: message,
		Data: gin.H{
			"enabled": enabled,
		},
	})
}

// SendEmailRequest структура запроса для отправки email
type SendEmailRequest struct {
	To      string `json:"to" binding:"required,email" example:"user@example.com"`
	Subject string `json:"subject" binding:"required" example:"Hello"`
	Body    string `json:"body" binding:"required" example:"<h1>Hello World</h1>"`
}

// WelcomeEmailRequest структура запроса для приветственного письма
type WelcomeEmailRequest struct {
	To       string `json:"to" binding:"required,email" example:"user@example.com"`
	Username string `json:"username" binding:"required" example:"john_doe"`
}

// PasswordResetEmailRequest структура запроса для сброса пароля
type PasswordResetEmailRequest struct {
	To        string `json:"to" binding:"required,email" example:"user@example.com"`
	ResetLink string `json:"reset_link" binding:"required" example:"https://example.com/reset?token=abc123"`
}
