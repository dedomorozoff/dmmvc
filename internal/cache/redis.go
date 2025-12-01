package cache

import (
	"context"
	"os"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
)

var (
	Client *redis.Client
	ctx    = context.Background()
)

// Connect подключается к Redis
func Connect() error {
	redisURL := os.Getenv("REDIS_URL")
	if redisURL == "" {
		redisURL = "localhost:6379"
	}

	redisPassword := os.Getenv("REDIS_PASSWORD")
	redisDB := 0 // default DB

	Client = redis.NewClient(&redis.Options{
		Addr:     redisURL,
		Password: redisPassword,
		DB:       redisDB,
	})

	// Проверка подключения
	_, err := Client.Ping(ctx).Result()
	if err != nil {
		logrus.Warn("Redis connection failed, caching disabled: ", err)
		return err
	}

	logrus.Info("Redis connected successfully")
	return nil
}

// Set сохраняет значение в кеш
func Set(key string, value interface{}, expiration time.Duration) error {
	if Client == nil {
		return nil // Кеш отключен
	}
	return Client.Set(ctx, key, value, expiration).Err()
}

// Get получает значение из кеша
func Get(key string) (string, error) {
	if Client == nil {
		return "", redis.Nil // Кеш отключен
	}
	return Client.Get(ctx, key).Result()
}

// Delete удаляет значение из кеша
func Delete(key string) error {
	if Client == nil {
		return nil
	}
	return Client.Del(ctx, key).Err()
}

// Exists проверяет существование ключа
func Exists(key string) bool {
	if Client == nil {
		return false
	}
	val, err := Client.Exists(ctx, key).Result()
	return err == nil && val > 0
}

// Clear очищает весь кеш (осторожно!)
func Clear() error {
	if Client == nil {
		return nil
	}
	return Client.FlushDB(ctx).Err()
}

// SetJSON сохраняет JSON в кеш
func SetJSON(key string, value interface{}, expiration time.Duration) error {
	if Client == nil {
		return nil
	}
	return Client.Set(ctx, key, value, expiration).Err()
}

// GetJSON получает JSON из кеша
func GetJSON(key string, dest interface{}) error {
	if Client == nil {
		return redis.Nil
	}
	return Client.Get(ctx, key).Scan(dest)
}

// IsEnabled проверяет, включен ли кеш
func IsEnabled() bool {
	return Client != nil
}
