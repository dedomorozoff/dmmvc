# Переключатели Функций

DMMVC поддерживает модульное управление функциями, позволяя включать или отключать возможности фреймворка в зависимости от потребностей вашего приложения. Это помогает уменьшить зависимости, улучшить производительность и упростить развертывание.

## Доступные Функции

| Функция | По умолчанию | Описание |
|---------|--------------|----------|
| `WEBSOCKET_ENABLE` | `true` | Поддержка WebSocket для связи в реальном времени |
| `REDIS_ENABLE` | `false` | Кэширование Redis для улучшения производительности |
| `SWAGGER_ENABLE` | `true` | Документация API Swagger |
| `FILE_UPLOAD_ENABLE` | `true` | Загрузка файлов и обработка изображений |
| `I18N_ENABLE` | `true` | Поддержка интернационализации (i18n) |
| `QUEUE_ENABLE` | `false` | Очередь фоновых задач с Redis |
| `EMAIL_ENABLE` | `false` | Отправка email через SMTP |

## Конфигурация

Функции настраиваются через переменные окружения в файле `.env`:

```env
# Включить/выключить функции фреймворка (true/false)
WEBSOCKET_ENABLE=true
REDIS_ENABLE=false
SWAGGER_ENABLE=true
FILE_UPLOAD_ENABLE=true
I18N_ENABLE=true
QUEUE_ENABLE=false
EMAIL_ENABLE=false
```

## Как Это Работает

### 1. Конфигурация Функций

Пакет `internal/config/features.go` читает переменные окружения и предоставляет централизованную конфигурацию:

```go
features := config.GetFeatures()

if features.WebSocket {
    // WebSocket включен
}
```

### 2. Условная Инициализация

В `cmd/server/main.go` функции инициализируются только когда включены:

```go
// Redis инициализируется только если REDIS_ENABLE=true
if features.Redis {
    if err := cache.Connect(); err != nil {
        logger.Log.Warn("Redis недоступен")
    }
}
```

### 3. Условные Маршруты

В `internal/routes/routes.go` маршруты регистрируются только для включенных функций:

```go
// Маршруты Swagger доступны только если SWAGGER_ENABLE=true
if features.Swagger {
    r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
```

## Примеры Использования

### Минимальная Настройка (Только Ядро)

Для минимального приложения только с аутентификацией и базой данных:

```env
WEBSOCKET_ENABLE=false
REDIS_ENABLE=false
SWAGGER_ENABLE=false
FILE_UPLOAD_ENABLE=false
I18N_ENABLE=false
QUEUE_ENABLE=false
EMAIL_ENABLE=false
```

### Настройка для Разработки

Для разработки с полной документацией:

```env
WEBSOCKET_ENABLE=true
REDIS_ENABLE=false
SWAGGER_ENABLE=true
FILE_UPLOAD_ENABLE=true
I18N_ENABLE=true
QUEUE_ENABLE=false
EMAIL_ENABLE=false
```

### Настройка для Продакшена

Для продакшена с кэшированием и фоновыми задачами:

```env
WEBSOCKET_ENABLE=true
REDIS_ENABLE=true
SWAGGER_ENABLE=false
FILE_UPLOAD_ENABLE=true
I18N_ENABLE=true
QUEUE_ENABLE=true
EMAIL_ENABLE=true
```

## Преимущества

1. **Уменьшенные Зависимости**: Отключенные функции не требуют настройки своих зависимостей
2. **Быстрый Запуск**: Пропуск инициализации неиспользуемых функций
3. **Чистые Логи**: Нет предупреждений о ненастроенных функциях, которые вам не нужны
4. **Меньшая Поверхность Атаки**: Отключенные функции не открывают маршруты или функциональность
5. **Гибкое Развертывание**: Разные окружения могут включать разные функции

## Зависимости Функций

Некоторые функции зависят от других:

- **Queue** требует **Redis** (обе должны быть включены вместе)
- **Email** может работать с **Queue** для асинхронной отправки (опционально)

## Проверка Статуса Функций

При запуске сервер логирует, какие функции включены:

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

## Добавление Новых Функций

Чтобы добавить новый переключатель функции:

1. Добавьте поле в структуру `Features` в `internal/config/features.go`
2. Инициализируйте его в функции `InitFeatures()`
3. Добавьте условную инициализацию в `cmd/server/main.go`
4. Добавьте условные маршруты в `internal/routes/routes.go`
5. Обновите эту документацию

Пример:

```go
// В internal/config/features.go
type Features struct {
    // ... существующие функции
    MyNewFeature bool
}

// В InitFeatures()
features = &Features{
    // ... существующие функции
    MyNewFeature: getBoolEnv("MY_NEW_FEATURE_ENABLE", false),
}
```
