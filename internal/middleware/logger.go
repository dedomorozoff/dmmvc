package middleware

import (
	"dmmvc/internal/logger"
	"time"

	"github.com/gin-gonic/gin"
)

// RequestLogger логирует все входящие запросы
func RequestLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		startTime := time.Now()

		c.Next()

		duration := time.Since(startTime)

		logger.Log.WithFields(map[string]interface{}{
			"method":   c.Request.Method,
			"path":     c.Request.URL.Path,
			"status":   c.Writer.Status(),
			"duration": duration,
			"ip":       c.ClientIP(),
		}).Info("Request processed")
	}
}
