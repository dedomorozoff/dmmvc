package middleware

import (
	"dmmvc/internal/config"

	"github.com/gin-gonic/gin"
)

// InjectFeatures добавляет информацию о включенных функциях в контекст
func InjectFeatures() gin.HandlerFunc {
	return func(c *gin.Context) {
		features := config.GetFeatures()
		c.Set("features", features)
		c.Next()
	}
}
