# Feature Toggles Cheatsheet

Quick reference for enabling/disabling DMMVC features.

## Environment Variables

```env
# WebSocket - Real-time communication
WEBSOCKET_ENABLE=true

# Redis - Caching layer
REDIS_ENABLE=false

# Swagger - API documentation
SWAGGER_ENABLE=true

# File Upload - File handling & image processing
FILE_UPLOAD_ENABLE=true

# i18n - Multi-language support
I18N_ENABLE=true

# Queue - Background job processing (requires Redis)
QUEUE_ENABLE=false

# Email - SMTP email sending
EMAIL_ENABLE=false
```

## Quick Presets

### Minimal (API only)
```env
WEBSOCKET_ENABLE=false
REDIS_ENABLE=false
SWAGGER_ENABLE=true
FILE_UPLOAD_ENABLE=false
I18N_ENABLE=false
QUEUE_ENABLE=false
EMAIL_ENABLE=false
```

### Development (Full features + docs)
```env
WEBSOCKET_ENABLE=true
REDIS_ENABLE=false
SWAGGER_ENABLE=true
FILE_UPLOAD_ENABLE=true
I18N_ENABLE=true
QUEUE_ENABLE=false
EMAIL_ENABLE=false
```

### Production (Optimized)
```env
WEBSOCKET_ENABLE=true
REDIS_ENABLE=true
SWAGGER_ENABLE=false
FILE_UPLOAD_ENABLE=true
I18N_ENABLE=true
QUEUE_ENABLE=true
EMAIL_ENABLE=true
```

## Feature Routes

| Feature | Routes |
|---------|--------|
| WebSocket | `/ws`, `/websocket` |
| Swagger | `/swagger/*` |
| File Upload | `/upload`, `/api/upload/*` |
| i18n | `/i18n`, `/api/locale`, `/api/locales` |
| Redis | `/api/users/cached`, `/api/cache/*` |
| Queue | `/api/queue/*` |
| Email | `/api/email/*` |

## Dependencies

- **Queue** requires **Redis** (both must be enabled)
- **Email** works better with **Queue** for async sending

## Checking Status

Look for this in server logs:
```
INFO Feature configuration:
INFO   WebSocket: true
INFO   Redis: false
...
```
