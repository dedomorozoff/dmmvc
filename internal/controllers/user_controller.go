package controllers

import (
	"dmmvc/internal/database"
	"dmmvc/internal/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// UserList отображает список пользователей
func UserList(c *gin.Context) {
	var users []models.User
	database.DB.Find(&users)
	features, _ := c.Get("features")

	c.HTML(http.StatusOK, "pages/users/list.html", gin.H{
		"title":    "Users",
		"users":    users,
		"features": features,
	})
}

// UserCreate отображает форму создания пользователя
func UserCreate(c *gin.Context) {
	features, _ := c.Get("features")
	
	c.HTML(http.StatusOK, "pages/users/create.html", gin.H{
		"title":    "Create User",
		"features": features,
	})
}

// UserStore сохраняет нового пользователя
func UserStore(c *gin.Context) {
	username := c.PostForm("username")
	email := c.PostForm("email")
	password := c.PostForm("password")
	role := c.PostForm("role")
	features, _ := c.Get("features")

	// Хеширование пароля
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "pages/users/create.html", gin.H{
			"title":    "Create User",
			"error":    "Error creating user",
			"features": features,
		})
		return
	}

	user := models.User{
		Username: username,
		Email:    email,
		Password: string(hashedPassword),
		Role:     role,
	}

	if err := database.DB.Create(&user).Error; err != nil {
		c.HTML(http.StatusInternalServerError, "pages/users/create.html", gin.H{
			"title":    "Create User",
			"error":    "Error creating user: " + err.Error(),
			"features": features,
		})
		return
	}

	c.Redirect(http.StatusFound, "/admin/users")
}

// UserEdit отображает форму редактирования пользователя
func UserEdit(c *gin.Context) {
	id := c.Param("id")
	var user models.User
	features, _ := c.Get("features")

	if err := database.DB.First(&user, id).Error; err != nil {
		c.Redirect(http.StatusFound, "/admin/users")
		return
	}

	c.HTML(http.StatusOK, "pages/users/edit.html", gin.H{
		"title":    "Edit User",
		"user":     user,
		"features": features,
	})
}

// UserUpdate обновляет пользователя
func UserUpdate(c *gin.Context) {
	id := c.Param("id")
	var user models.User
	features, _ := c.Get("features")

	if err := database.DB.First(&user, id).Error; err != nil {
		c.Redirect(http.StatusFound, "/admin/users")
		return
	}

	user.Username = c.PostForm("username")
	user.Email = c.PostForm("email")
	user.Role = c.PostForm("role")

	// Обновление пароля только если указан новый
	newPassword := c.PostForm("password")
	if newPassword != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
		if err != nil {
			c.HTML(http.StatusInternalServerError, "pages/users/edit.html", gin.H{
				"title":    "Edit User",
				"user":     user,
				"error":    "Error updating password",
				"features": features,
			})
			return
		}
		user.Password = string(hashedPassword)
	}

	if err := database.DB.Save(&user).Error; err != nil {
		c.HTML(http.StatusInternalServerError, "pages/users/edit.html", gin.H{
			"title":    "Edit User",
			"user":     user,
			"error":    "Error updating user: " + err.Error(),
			"features": features,
		})
		return
	}

	c.Redirect(http.StatusFound, "/admin/users")
}

// UserDelete удаляет пользователя
func UserDelete(c *gin.Context) {
	id := c.Param("id")
	idInt, _ := strconv.Atoi(id)

	// Запрещаем удаление самого себя
	userID, _ := c.Get("user_id")
	if uint(idInt) == userID.(uint) {
		c.Redirect(http.StatusFound, "/admin/users")
		return
	}

	database.DB.Delete(&models.User{}, id)
	c.Redirect(http.StatusFound, "/admin/users")
}
