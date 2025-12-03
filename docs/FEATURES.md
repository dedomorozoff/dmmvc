# Feature Toggles

DMMVC supports modular feature management, allowing you to enable or disable framework features based on your application needs. This helps reduce dependencies, improve performance, and simplify deployment.

## Available Features

| Feature | Default | Description |
|---------|---------|-------------|
| `WEBSOCKET_ENABLE` | `false` | WebSocket support for real-time communication |
| `REDIS_ENABLE` | `false` | Redis caching for improved performance |
| `SWAGGER_ENABLE` | `false` | Swagger API documentation |
| `FILE_UPLOAD_ENABLE` | `false` | File upload and image processing |
| `I18N_ENABLE` | `false` | Internationalization (i18n) support |
| `QUEUE_ENABLE` | `false` | Background task queue with Redis |
| `EMAIL_ENABLE` | `false` | Email sending via SMTP |

## Configuration

Features are configured via environment variables in your `.env` file:

```env
# Enable/disable framework features (true/false)
WEBSOCKET_ENABLE=true
REDIS_ENABLE=false
SWAGGER_ENABLE=true
FILE_UPLOAD_ENABLE=true
I18N_ENABLE=true
QUEUE_ENABLE=false
EMAIL_ENABLE=false
```

## How It Works

### 1. Feature Configuration

The `internal/config/features.go` package reads environment variables and provides a centralized configuration:

```go
features := config.GetFeatures()

if features.WebSocket {
    // WebSocket is enabled
}
```

### 2. Conditional Initialization

In `cmd/server/main.go`, features are initialized only when enabled:

```go
// Redis is only initialized if REDIS_ENABLE=true
if features.Redis {
    if err := cache.Connect(); err != nil {
        logger.Log.Warn("Redis not available")
    }
}
```

### 3. Conditional Routes

In `internal/routes/routes.go`, routes are registered only for enabled features:

```go
// Swagger routes only available if SWAGGER_ENABLE=true
if features.Swagger {
    r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
```

## Use Cases

### Minimal Setup (Core Only)

For a minimal application with just authentication and database:

```env
WEBSOCKET_ENABLE=false
REDIS_ENABLE=false
SWAGGER_ENABLE=false
FILE_UPLOAD_ENABLE=false
I18N_ENABLE=false
QUEUE_ENABLE=false
EMAIL_ENABLE=false
```

### Development Setup

For development with full documentation:

```env
WEBSOCKET_ENABLE=true
REDIS_ENABLE=false
SWAGGER_ENABLE=true
FILE_UPLOAD_ENABLE=true
I18N_ENABLE=true
QUEUE_ENABLE=false
EMAIL_ENABLE=false
```

### Production Setup

For production with caching and background jobs:

```env
WEBSOCKET_ENABLE=true
REDIS_ENABLE=true
SWAGGER_ENABLE=false
FILE_UPLOAD_ENABLE=true
I18N_ENABLE=true
QUEUE_ENABLE=true
EMAIL_ENABLE=true
```

## Benefits

1. **Reduced Dependencies**: Disabled features don't require their dependencies to be configured
2. **Faster Startup**: Skip initialization of unused features
3. **Cleaner Logs**: No warnings about unconfigured features you don't need
4. **Smaller Attack Surface**: Disabled features don't expose routes or functionality
5. **Flexible Deployment**: Different environments can enable different features

## Feature Dependencies

Some features depend on others:

- **Queue** requires **Redis** (both should be enabled together)
- **Email** can work with **Queue** for async sending (optional)

## Checking Feature Status

On startup, the server logs which features are enabled:

```
INFO Feature configuration:
INFO   WebSocket: true
INFO   Redis: false
INFO   Swagger: true
INFO   File Upload: true
INFO   I18n: true
INFO   Queue: false
INFO   Email: false
```

## Adding New Features

To add a new feature toggle:

1. Add the field to `Features` struct in `internal/config/features.go`
2. Initialize it in `InitFeatures()` function
3. Add conditional initialization in `cmd/server/main.go`
4. Add conditional routes in `internal/routes/routes.go`
5. Update this documentation

Example:

```go
// In internal/config/features.go
type Features struct {
    // ... existing features
    MyNewFeature bool
}

// In InitFeatures()
features = &Features{
    // ... existing features
    MyNewFeature: getBoolEnv("MY_NEW_FEATURE_ENABLE", false),
}
```
