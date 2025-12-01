package controllers

import (
	"dmmvc/internal/database"
	"dmmvc/internal/models"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// LoginPage отображает страницу входа
func LoginPage(c *gin.Context) {
	// Если пользователь уже авторизован, перенаправляем на dashboard
	session := sessions.Default(c)
	if session.Get("user_id") != nil {
		c.Redirect(http.StatusFound, "/dashboard")
		return
	}

	c.HTML(http.StatusOK, "pages/login.html", gin.H{
		"title": "Вход",
	})
}

// LoginPost обрабатывает форму входа
func LoginPost(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	var user models.User
	if err := database.DB.Where("username = ?", username).First(&user).Error; err != nil {
		c.HTML(http.StatusUnauthorized, "pages/login.html", gin.H{
			"title": "Вход",
			"error": "Неверное имя пользователя или пароль",
		})
		return
	}

	// Проверка пароля
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		c.HTML(http.StatusUnauthorized, "pages/login.html", gin.H{
			"title": "Вход",
			"error": "Неверное имя пользователя или пароль",
		})
		return
	}

	// Сохранение сессии
	session := sessions.Default(c)
	session.Set("user_id", user.ID)
	session.Set("username", user.Username)
	session.Set("role", user.Role)
	session.Save()

	c.Redirect(http.StatusFound, "/dashboard")
}

// Logout выход из системы
func Logout(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	session.Save()

	c.Redirect(http.StatusFound, "/login")
}
