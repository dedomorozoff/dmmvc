package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// HomePage отображает главную страницу
func HomePage(c *gin.Context) {
	c.HTML(http.StatusOK, "pages/home.html", gin.H{
		"title": "Главная",
	})
}

// DashboardPage отображает панель управления
func DashboardPage(c *gin.Context) {
	username, _ := c.Get("username")
	role, _ := c.Get("role")

	c.HTML(http.StatusOK, "pages/dashboard.html", gin.H{
		"title":    "Панель управления",
		"username": username,
		"role":     role,
	})
}

// ProfilePage отображает профиль пользователя
func ProfilePage(c *gin.Context) {
	username, _ := c.Get("username")
	role, _ := c.Get("role")

	c.HTML(http.StatusOK, "pages/profile.html", gin.H{
		"title":    "Профиль",
		"username": username,
		"role":     role,
	})
}
