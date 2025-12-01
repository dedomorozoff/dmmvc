[English](CACHE.md) | **Русский**

# Redis Кеширование

DMMVC включает встроенную поддержку Redis кеширования для улучшения производительности приложения.

## Возможности

- **Автоматическое кеширование** - Cache middleware для GET запросов
- **Ручное кеширование** - Прямые операции с кешем
- **Опционально** - Работает без Redis (graceful degradation)
- **Простой API** - Легкие в использовании функции кеша

## Быстрый старт

### 1. Установка Redis

**Windows:**
```bash
# Используя Chocolatey
choco install redis-64

# Или скачайте с: https://github.com/microsoftarchive/redis/releases
```

**Linux:**
```bash
sudo apt-get install redis-server
```

**macOS:**
```bash
brew install redis
```

**Docker:**
```bash
docker run -d -p 6379:6379 redis:alpine
```

### 2. Настройка Redis

Отредактируйте файл `.env`:

```env
# Redis Cache Settings
REDIS_URL=localhost:6379
REDIS_PASSWORD=
REDIS_ENABLED=true
```

### 3. Запуск сервера

```bash
make run
```

Если Redis недоступен, приложение будет работать без кеширования.

## Использование

### Ручное кеширование

```go
import (
    "dmmvc/internal/cache"
    "time"
)

// Сохранить в кеш
cache.Set("key", "value", 5*time.Minute)

// Получить из кеша
value, err := cache.Get("key")
if err == nil {
    // Использовать закешированное значение
}

// Удалить из кеша
cache.Delete("key")

// Проверить существование ключа
if cache.Exists("key") {
    // Ключ существует
}
```

### JSON кеширование

```go
// Сохранить структуру в кеш
user := models.User{Username: "john"}
cache.SetJSON("user:1", user, 10*time.Minute)

// Получить структуру из кеша
var user models.User
err := cache.GetJSON("user:1", &user)
```

### Cache Middleware

Применить cache middleware к маршрутам:

```go
import (
    "dmmvc/internal/middleware"
    "time"
)

// Кешировать на 5 минут
r.GET("/api/data", middleware.CacheMiddleware(5*time.Minute), controllers.GetData)
```

Middleware:
- Кеширует только GET запросы
- Кеширует только успешные ответы (200 OK)
- Добавляет заголовок `X-Cache: HIT` или `X-Cache: MISS`

## Пример контроллера

```go
func GetUsers(c *gin.Context) {
    cacheKey := "users:list"

    // Попытка получить из кеша
    if cache.IsEnabled() {
        cachedData, err := cache.Get(cacheKey)
        if err == nil {
            var users []models.User
            json.Unmarshal([]byte(cachedData), &users)
            c.JSON(200, users)
            return
        }
    }

    // Получение из БД
    var users []models.User
    database.DB.Find(&users)

    // Сохранение в кеш на 5 минут
    if cache.IsEnabled() {
        data, _ := json.Marshal(users)
        cache.Set(cacheKey, string(data), 5*time.Minute)
    }

    c.JSON(200, users)
}
```

## Инвалидация кеша

Очистка кеша при изменении данных:

```go
func UpdateUser(c *gin.Context) {
    // Обновление пользователя в БД
    database.DB.Save(&user)

    // Очистка связанного кеша
    cache.Delete("users:list")
    cache.Delete("user:" + user.ID)

    c.JSON(200, user)
}
```

## API эндпоинты

DMMVC включает примеры cache эндпоинтов:

- `GET /api/users/cached` - Получить пользователей с кешированием
- `POST /api/cache/clear` - Очистить кеш пользователей
- `GET /api/cache/stats` - Статистика кеша

## Функции кеша

### Базовые операции

```go
// Установить значение с истечением
cache.Set(key string, value interface{}, expiration time.Duration) error

// Получить значение
cache.Get(key string) (string, error)

// Удалить значение
cache.Delete(key string) error

// Проверить существование
cache.Exists(key string) bool

// Очистить весь кеш (осторожно!)
cache.Clear() error
```

### JSON операции

```go
// Установить JSON
cache.SetJSON(key string, value interface{}, expiration time.Duration) error

// Получить JSON
cache.GetJSON(key string, dest interface{}) error
```

### Утилиты

```go
// Проверить, включен ли кеш
cache.IsEnabled() bool
```

## Лучшие практики

1. **Используйте подходящий TTL** - Устанавливайте время истечения на основе частоты обновления данных
2. **Кешируйте дорогие операции** - Запросы к БД, API вызовы, вычисления
3. **Инвалидируйте при обновлениях** - Очищайте кеш при изменении данных
4. **Используйте ключи разумно** - Используйте описательные, иерархические ключи
5. **Обрабатывайте промахи кеша** - Всегда имейте fallback к БД
6. **Мониторьте использование кеша** - Отслеживайте hit/miss rates

## Паттерны ключей кеша

```go
// Данные пользователей
"user:{id}"
"users:list"
"users:active"

// Посты
"post:{id}"
"posts:recent"
"posts:user:{user_id}"

// Сессии
"session:{token}"

// API ответы
"api:endpoint:{params}"
```

## Production конфигурация

```env
# Production Redis настройки
REDIS_URL=redis.example.com:6379
REDIS_PASSWORD=your-secure-password
REDIS_ENABLED=true

# Используйте Redis Sentinel для высокой доступности
# REDIS_SENTINEL_MASTER=mymaster
# REDIS_SENTINEL_ADDRS=sentinel1:26379,sentinel2:26379
```

## Docker Compose

```yaml
version: '3.8'
services:
  app:
    build: .
    ports:
      - "8080:8080"
    environment:
      - REDIS_URL=redis:6379
    depends_on:
      - redis

  redis:
    image: redis:alpine
    ports:
      - "6379:6379"
    volumes:
      - redis_data:/data

volumes:
  redis_data:
```

## Решение проблем

### Redis не подключается

Проверьте, запущен ли Redis:
```bash
redis-cli ping
# Должен вернуть: PONG
```

### Кеш не работает

1. Проверьте конфигурацию `.env`
2. Проверьте подключение Redis в логах
3. Убедитесь, что `REDIS_ENABLED=true`

### Очистить весь кеш

```bash
redis-cli FLUSHDB
```

## Советы по производительности

1. **Используйте connection pooling** - Уже настроен в go-redis
2. **Устанавливайте подходящие таймауты** - Предотвращайте медленные запросы
3. **Используйте pipelining** - Для множественных операций
4. **Мониторьте использование памяти** - Установите maxmemory policy
5. **Используйте Redis persistence** - RDB или AOF для сохранности данных

## Ресурсы

- [Документация Redis](https://redis.io/documentation)
- [Документация go-redis](https://redis.uptrace.dev/)
- [Redis Best Practices](https://redis.io/docs/manual/patterns/)

## Следующие шаги

- Реализовать стратегии прогрева кеша
- Добавить метрики и мониторинг кеша
- Использовать Redis для хранения сессий
- Реализовать распределенные блокировки
- Добавить pub/sub для real-time функций
