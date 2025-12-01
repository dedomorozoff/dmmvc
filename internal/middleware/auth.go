package middleware

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// AuthRequired проверяет авторизацию пользователя
func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		userID := session.Get("user_id")

		if userID == nil {
			c.Redirect(http.StatusFound, "/login")
			c.Abort()
			return
		}

		// Добавляем данные пользователя в контекст
		username := session.Get("username")
		role := session.Get("role")
		
		c.Set("user_id", userID)
		c.Set("username", username)
		c.Set("role", role)

		c.Next()
	}
}

// InjectUserData добавляет данные пользователя в шаблоны
func InjectUserData() gin.HandlerFunc {
	return func(c *gin.Context) {
		username, _ := c.Get("username")
		role, _ := c.Get("role")

		c.Set("template_username", username)
		c.Set("template_role", role)

		c.Next()
	}
}
