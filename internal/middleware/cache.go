package middleware

import (
	"crypto/md5"
	"dmmvc/internal/cache"
	"encoding/hex"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// CacheMiddleware кеширует GET запросы
func CacheMiddleware(duration time.Duration) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Кешируем только GET запросы
		if c.Request.Method != "GET" {
			c.Next()
			return
		}

		// Если кеш отключен, пропускаем
		if !cache.IsEnabled() {
			c.Next()
			return
		}

		// Генерируем ключ кеша на основе URL
		cacheKey := generateCacheKey(c.Request.URL.String())

		// Проверяем наличие в кеше
		cachedResponse, err := cache.Get(cacheKey)
		if err == nil && cachedResponse != "" {
			c.Header("X-Cache", "HIT")
			c.String(http.StatusOK, cachedResponse)
			c.Abort()
			return
		}

		// Создаем writer для перехвата ответа
		writer := &responseWriter{
			ResponseWriter: c.Writer,
			body:           []byte{},
		}
		c.Writer = writer

		c.Next()

		// Кешируем только успешные ответы
		if c.Writer.Status() == http.StatusOK {
			cache.Set(cacheKey, string(writer.body), duration)
			c.Header("X-Cache", "MISS")
		}
	}
}

// responseWriter перехватывает ответ для кеширования
type responseWriter struct {
	gin.ResponseWriter
	body []byte
}

func (w *responseWriter) Write(b []byte) (int, error) {
	w.body = append(w.body, b...)
	return w.ResponseWriter.Write(b)
}

// generateCacheKey генерирует ключ кеша
func generateCacheKey(url string) string {
	hash := md5.Sum([]byte(url))
	return "cache:" + hex.EncodeToString(hash[:])
}
