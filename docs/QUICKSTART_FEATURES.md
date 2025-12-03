# Quick Start: Feature Toggles

This guide shows how to quickly configure DMMVC features for different scenarios.

## Scenarios

### 1. Minimal API (No UI Features)

Perfect for REST APIs without WebSocket, file uploads, or i18n:

```env
WEBSOCKET_ENABLE=false
REDIS_ENABLE=false
SWAGGER_ENABLE=true
FILE_UPLOAD_ENABLE=false
I18N_ENABLE=false
QUEUE_ENABLE=false
EMAIL_ENABLE=false
```

**What you get:**
- Basic authentication and sessions
- Database (SQLite/MySQL/PostgreSQL)
- REST API endpoints
- Swagger documentation

### 2. Simple Web App

For traditional web applications:

```env
WEBSOCKET_ENABLE=false
REDIS_ENABLE=false
SWAGGER_ENABLE=false
FILE_UPLOAD_ENABLE=true
I18N_ENABLE=true
QUEUE_ENABLE=false
EMAIL_ENABLE=false
```

**What you get:**
- User authentication
- File uploads
- Multi-language support
- Template rendering

### 3. Real-time Application

For chat apps, dashboards, or live updates:

```env
WEBSOCKET_ENABLE=true
REDIS_ENABLE=true
SWAGGER_ENABLE=true
FILE_UPLOAD_ENABLE=true
I18N_ENABLE=true
QUEUE_ENABLE=false
EMAIL_ENABLE=false
```

**What you get:**
- WebSocket support
- Redis caching
- All basic features

### 4. Full-Featured Production

For production apps with all features:

```env
WEBSOCKET_ENABLE=true
REDIS_ENABLE=true
SWAGGER_ENABLE=false
FILE_UPLOAD_ENABLE=true
I18N_ENABLE=true
QUEUE_ENABLE=true
EMAIL_ENABLE=true
```

**What you get:**
- All features enabled
- Background job processing
- Email notifications
- Production-ready setup

## Step-by-Step Setup

### 1. Copy Environment File

```bash
cp .env.example .env
```

### 2. Edit Feature Flags

Open `.env` and set your desired features:

```env
# Example: Enable only what you need
WEBSOCKET_ENABLE=true
REDIS_ENABLE=false
SWAGGER_ENABLE=true
FILE_UPLOAD_ENABLE=true
I18N_ENABLE=true
QUEUE_ENABLE=false
EMAIL_ENABLE=false
```

### 3. Configure Required Services

Only configure services for enabled features:

**If REDIS_ENABLE=true:**
```env
REDIS_URL=localhost:6379
REDIS_PASSWORD=
```

**If QUEUE_ENABLE=true:**
```env
QUEUE_WORKER_ENABLED=true
QUEUE_CONCURRENCY=10
```

**If EMAIL_ENABLE=true:**
```env
SMTP_HOST=smtp.gmail.com
SMTP_PORT=587
SMTP_USERNAME=your-email@gmail.com
SMTP_PASSWORD=your-app-password
```

**If FILE_UPLOAD_ENABLE=true:**
```env
UPLOAD_DIR=./uploads
MAX_FILE_SIZE=10485760
```

### 4. Start the Server

```bash
go run cmd/server/main.go
```

### 5. Check Logs

The server will log which features are enabled:

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

## Testing Features

### Test WebSocket (if enabled)

Visit: `http://localhost:8080/websocket`

### Test Swagger (if enabled)

Visit: `http://localhost:8080/swagger/index.html`

### Test File Upload (if enabled)

Visit: `http://localhost:8080/upload`

### Test i18n (if enabled)

Visit: `http://localhost:8080/i18n`

## Troubleshooting

### Feature Not Working

1. Check `.env` file - is the feature enabled?
2. Check logs - did the feature initialize successfully?
3. Check dependencies - is Redis running (for Redis/Queue features)?

### Routes Not Available

If routes return 404:
- Feature is disabled in `.env`
- Restart the server after changing `.env`

### Warnings in Logs

Warnings like "Redis not available" are normal when features are disabled. To remove them, ensure the feature is set to `false` in `.env`.

## Next Steps

- Read full documentation: [FEATURES.md](FEATURES.md)
- Configure specific features:
  - [WebSocket](WEBSOCKET.md)
  - [Redis Cache](CACHE.md)
  - [File Upload](UPLOAD.md)
  - [i18n](I18N.md)
  - [Queue](QUEUE.md)
  - [Email](EMAIL.md)
