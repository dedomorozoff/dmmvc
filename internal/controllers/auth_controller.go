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

	features, _ := c.Get("features")
	data := gin.H{
		"title":    "Login",
		"features": features,
	}
	
	// Add i18n support if enabled
	addI18nData(c, data)
	
	c.HTML(http.StatusOK, "pages/login.html", data)
}

// LoginPost обрабатывает форму входа
func LoginPost(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	features, _ := c.Get("features")

	var user models.User
	if err := database.DB.Where("username = ?", username).First(&user).Error; err != nil {
		data := gin.H{
			"title":    "Login",
			"error":    "Invalid username or password",
			"features": features,
		}
		addI18nData(c, data)
		c.HTML(http.StatusUnauthorized, "pages/login.html", data)
		return
	}

	// Проверка пароля
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		data := gin.H{
			"title":    "Login",
			"error":    "Invalid username or password",
			"features": features,
		}
		addI18nData(c, data)
		c.HTML(http.StatusUnauthorized, "pages/login.html", data)
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
