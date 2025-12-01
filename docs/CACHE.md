**English** | [Русский](CACHE.ru.md)

# Redis Caching

DMMVC includes built-in support for Redis caching to improve application performance.

## Features

- **Automatic caching** - Cache middleware for GET requests
- **Manual caching** - Direct cache operations
- **Optional** - Works without Redis (graceful degradation)
- **Simple API** - Easy to use cache functions

## Quick Start

### 1. Install Redis

**Windows:**
```bash
# Using Chocolatey
choco install redis-64

# Or download from: https://github.com/microsoftarchive/redis/releases
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

### 2. Configure Redis

Edit `.env` file:

```env
# Redis Cache Settings
REDIS_URL=localhost:6379
REDIS_PASSWORD=
REDIS_ENABLED=true
```

### 3. Start Server

```bash
make run
```

If Redis is not available, the application will work without caching.

## Usage

### Manual Caching

```go
import (
    "dmmvc/internal/cache"
    "time"
)

// Save to cache
cache.Set("key", "value", 5*time.Minute)

// Get from cache
value, err := cache.Get("key")
if err == nil {
    // Use cached value
}

// Delete from cache
cache.Delete("key")

// Check if key exists
if cache.Exists("key") {
    // Key exists
}
```

### JSON Caching

```go
// Save struct to cache
user := models.User{Username: "john"}
cache.SetJSON("user:1", user, 10*time.Minute)

// Get struct from cache
var user models.User
err := cache.GetJSON("user:1", &user)
```

### Cache Middleware

Apply cache middleware to routes:

```go
import (
    "dmmvc/internal/middleware"
    "time"
)

// Cache for 5 minutes
r.GET("/api/data", middleware.CacheMiddleware(5*time.Minute), controllers.GetData)
```

The middleware:
- Only caches GET requests
- Only caches successful responses (200 OK)
- Adds `X-Cache: HIT` or `X-Cache: MISS` header

## Example Controller

```go
func GetUsers(c *gin.Context) {
    cacheKey := "users:list"

    // Try to get from cache
    if cache.IsEnabled() {
        cachedData, err := cache.Get(cacheKey)
        if err == nil {
            var users []models.User
            json.Unmarshal([]byte(cachedData), &users)
            c.JSON(200, users)
            return
        }
    }

    // Get from database
    var users []models.User
    database.DB.Find(&users)

    // Save to cache for 5 minutes
    if cache.IsEnabled() {
        data, _ := json.Marshal(users)
        cache.Set(cacheKey, string(data), 5*time.Minute)
    }

    c.JSON(200, users)
}
```

## Cache Invalidation

Clear cache when data changes:

```go
func UpdateUser(c *gin.Context) {
    // Update user in database
    database.DB.Save(&user)

    // Clear related cache
    cache.Delete("users:list")
    cache.Delete("user:" + user.ID)

    c.JSON(200, user)
}
```

## API Endpoints

DMMVC includes example cache endpoints:

- `GET /api/users/cached` - Get users with caching
- `POST /api/cache/clear` - Clear user cache
- `GET /api/cache/stats` - Cache statistics

## Cache Functions

### Basic Operations

```go
// Set value with expiration
cache.Set(key string, value interface{}, expiration time.Duration) error

// Get value
cache.Get(key string) (string, error)

// Delete value
cache.Delete(key string) error

// Check existence
cache.Exists(key string) bool

// Clear all cache (use carefully!)
cache.Clear() error
```

### JSON Operations

```go
// Set JSON
cache.SetJSON(key string, value interface{}, expiration time.Duration) error

// Get JSON
cache.GetJSON(key string, dest interface{}) error
```

### Utility

```go
// Check if cache is enabled
cache.IsEnabled() bool
```

## Best Practices

1. **Use appropriate TTL** - Set expiration based on data update frequency
2. **Cache expensive operations** - Database queries, API calls, computations
3. **Invalidate on updates** - Clear cache when data changes
4. **Use cache keys wisely** - Use descriptive, hierarchical keys
5. **Handle cache misses** - Always have fallback to database
6. **Monitor cache usage** - Track hit/miss rates

## Cache Key Patterns

```go
// User data
"user:{id}"
"users:list"
"users:active"

// Posts
"post:{id}"
"posts:recent"
"posts:user:{user_id}"

// Sessions
"session:{token}"

// API responses
"api:endpoint:{params}"
```

## Production Configuration

```env
# Production Redis settings
REDIS_URL=redis.example.com:6379
REDIS_PASSWORD=your-secure-password
REDIS_ENABLED=true

# Use Redis Sentinel for high availability
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

## Troubleshooting

### Redis not connecting

Check if Redis is running:
```bash
redis-cli ping
# Should return: PONG
```

### Cache not working

1. Check `.env` configuration
2. Verify Redis connection in logs
3. Ensure `REDIS_ENABLED=true`

### Clear all cache

```bash
redis-cli FLUSHDB
```

## Performance Tips

1. **Use connection pooling** - Already configured in go-redis
2. **Set appropriate timeouts** - Prevent slow queries
3. **Use pipelining** - For multiple operations
4. **Monitor memory usage** - Set maxmemory policy
5. **Use Redis persistence** - RDB or AOF for data durability

## Resources

- [Redis Documentation](https://redis.io/documentation)
- [go-redis Documentation](https://redis.uptrace.dev/)
- [Redis Best Practices](https://redis.io/docs/manual/patterns/)

## Next Steps

- Implement cache warming strategies
- Add cache metrics and monitoring
- Use Redis for session storage
- Implement distributed locking
- Add pub/sub for real-time features
