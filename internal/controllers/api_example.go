package controllers

import (
	"dmmvc/internal/database"
	"dmmvc/internal/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// APIResponse стандартный формат ответа API
type APIResponse struct {
	Success bool        `json:"success" example:"true"`
	Message string      `json:"message,omitempty" example:"Operation successful"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty" example:"Error message"`
}

// APIUserList godoc
// @Summary Список пользователей
// @Description Получить список всех пользователей
// @Tags users
// @Accept json
// @Produce json
// @Success 200 {object} APIResponse{data=[]models.User}
// @Failure 500 {object} APIResponse
// @Router /api/users [get]
// @Security SessionAuth
func APIUserList(c *gin.Context) {
	var users []models.User
	if err := database.DB.Find(&users).Error; err != nil {
		c.JSON(http.StatusInternalServerError, APIResponse{
			Success: false,
			Error:   "Failed to fetch users",
		})
		return
	}

	c.JSON(http.StatusOK, APIResponse{
		Success: true,
		Data:    users,
	})
}

// APIUserGet godoc
// @Summary Получить пользователя
// @Description Получить пользователя по ID
// @Tags users
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} APIResponse{data=models.User}
// @Failure 404 {object} APIResponse
// @Router /api/users/{id} [get]
// @Security SessionAuth
func APIUserGet(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var user models.User

	if err := database.DB.First(&user, id).Error; err != nil {
		c.JSON(http.StatusNotFound, APIResponse{
			Success: false,
			Error:   "User not found",
		})
		return
	}

	c.JSON(http.StatusOK, APIResponse{
		Success: true,
		Data:    user,
	})
}

// UserCreateRequest структура запроса создания пользователя
type UserCreateRequest struct {
	Username string `json:"username" binding:"required" example:"john_doe"`
	Email    string `json:"email" binding:"required,email" example:"john@example.com"`
	Password string `json:"password" binding:"required,min=6" example:"password123"`
}

// APIUserCreate godoc
// @Summary Создать пользователя
// @Description Создать нового пользователя
// @Tags users
// @Accept json
// @Produce json
// @Param user body UserCreateRequest true "User data"
// @Success 201 {object} APIResponse{data=models.User}
// @Failure 400 {object} APIResponse
// @Router /api/users [post]
// @Security SessionAuth
func APIUserCreate(c *gin.Context) {
	var req UserCreateRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, APIResponse{
			Success: false,
			Error:   err.Error(),
		})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, APIResponse{
			Success: false,
			Error:   "Failed to hash password",
		})
		return
	}

	user := models.User{
		Username: req.Username,
		Email:    req.Email,
		Password: string(hashedPassword),
	}

	if err := database.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, APIResponse{
			Success: false,
			Error:   "Failed to create user",
		})
		return
	}

	c.JSON(http.StatusCreated, APIResponse{
		Success: true,
		Message: "User created successfully",
		Data:    user,
	})
}

// APIUserDelete godoc
// @Summary Удалить пользователя
// @Description Удалить пользователя по ID
// @Tags users
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} APIResponse
// @Failure 404 {object} APIResponse
// @Router /api/users/{id} [delete]
// @Security SessionAuth
func APIUserDelete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var user models.User

	if err := database.DB.First(&user, id).Error; err != nil {
		c.JSON(http.StatusNotFound, APIResponse{
			Success: false,
			Error:   "User not found",
		})
		return
	}

	database.DB.Delete(&user)

	c.JSON(http.StatusOK, APIResponse{
		Success: true,
		Message: "User deleted successfully",
	})
}
