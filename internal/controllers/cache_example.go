package controllers

import (
	"dmmvc/internal/cache"
	"dmmvc/internal/database"
	"dmmvc/internal/models"
	"encoding/json"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// CachedUserList возвращает список пользователей с кешированием
func CachedUserList(c *gin.Context) {
	cacheKey := "users:list"

	// Попытка получить из кеша
	if cache.IsEnabled() {
		cachedData, err := cache.Get(cacheKey)
		if err == nil {
			var users []models.User
			if json.Unmarshal([]byte(cachedData), &users) == nil {
				c.JSON(http.StatusOK, APIResponse{
					Success: true,
					Data:    users,
					Message: "From cache",
				})
				return
			}
		}
	}

	// Получение из БД
	var users []models.User
	if err := database.DB.Find(&users).Error; err != nil {
		c.JSON(http.StatusInternalServerError, APIResponse{
			Success: false,
			Error:   "Failed to fetch users",
		})
		return
	}

	// Сохранение в кеш на 5 минут
	if cache.IsEnabled() {
		if data, err := json.Marshal(users); err == nil {
			cache.Set(cacheKey, string(data), 5*time.Minute)
		}
	}

	c.JSON(http.StatusOK, APIResponse{
		Success: true,
		Data:    users,
		Message: "From database",
	})
}

// ClearUserCache очищает кеш пользователей
func ClearUserCache(c *gin.Context) {
	if !cache.IsEnabled() {
		c.JSON(http.StatusOK, APIResponse{
			Success: false,
			Message: "Cache is disabled",
		})
		return
	}

	cache.Delete("users:list")

	c.JSON(http.StatusOK, APIResponse{
		Success: true,
		Message: "User cache cleared",
	})
}

// CacheStats возвращает статистику кеша
func CacheStats(c *gin.Context) {
	if !cache.IsEnabled() {
		c.JSON(http.StatusOK, gin.H{
			"enabled": false,
			"message": "Redis cache is disabled",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"enabled": true,
		"message": "Redis cache is enabled",
	})
}
