# Быстрый Старт: Переключатели Функций

Это руководство показывает, как быстро настроить функции DMMVC для различных сценариев.

## Сценарии

### 1. Минимальный API (Без UI Функций)

Идеально для REST API без WebSocket, загрузки файлов или i18n:

```env
WEBSOCKET_ENABLE=false
REDIS_ENABLE=false
SWAGGER_ENABLE=true
FILE_UPLOAD_ENABLE=false
I18N_ENABLE=false
QUEUE_ENABLE=false
EMAIL_ENABLE=false
```

**Что вы получаете:**
- Базовая аутентификация и сессии
- База данных (SQLite/MySQL/PostgreSQL)
- REST API эндпоинты
- Документация Swagger

### 2. Простое Веб-Приложение

Для традиционных веб-приложений:

```env
WEBSOCKET_ENABLE=false
REDIS_ENABLE=false
SWAGGER_ENABLE=false
FILE_UPLOAD_ENABLE=true
I18N_ENABLE=true
QUEUE_ENABLE=false
EMAIL_ENABLE=false
```

**Что вы получаете:**
- Аутентификация пользователей
- Загрузка файлов
- Поддержка нескольких языков
- Рендеринг шаблонов

### 3. Приложение Реального Времени

Для чатов, дашбордов или живых обновлений:

```env
WEBSOCKET_ENABLE=true
REDIS_ENABLE=true
SWAGGER_ENABLE=true
FILE_UPLOAD_ENABLE=true
I18N_ENABLE=true
QUEUE_ENABLE=false
EMAIL_ENABLE=false
```

**Что вы получаете:**
- Поддержка WebSocket
- Кэширование Redis
- Все базовые функции

### 4. Полнофункциональный Продакшен

Для продакшен приложений со всеми функциями:

```env
WEBSOCKET_ENABLE=true
REDIS_ENABLE=true
SWAGGER_ENABLE=false
FILE_UPLOAD_ENABLE=true
I18N_ENABLE=true
QUEUE_ENABLE=true
EMAIL_ENABLE=true
```

**Что вы получаете:**
- Все функции включены
- Обработка фоновых задач
- Email уведомления
- Готовая к продакшену настройка

## Пошаговая Настройка

### 1. Скопируйте Файл Окружения

```bash
cp .env.example .env
```

### 2. Отредактируйте Флаги Функций

Откройте `.env` и установите нужные функции:

```env
# Пример: Включите только то, что нужно
WEBSOCKET_ENABLE=true
REDIS_ENABLE=false
SWAGGER_ENABLE=true
FILE_UPLOAD_ENABLE=true
I18N_ENABLE=true
QUEUE_ENABLE=false
EMAIL_ENABLE=false
```

### 3. Настройте Необходимые Сервисы

Настраивайте только сервисы для включенных функций:

**Если REDIS_ENABLE=true:**
```env
REDIS_URL=localhost:6379
REDIS_PASSWORD=
```

**Если QUEUE_ENABLE=true:**
```env
QUEUE_WORKER_ENABLED=true
QUEUE_CONCURRENCY=10
```

**Если EMAIL_ENABLE=true:**
```env
SMTP_HOST=smtp.gmail.com
SMTP_PORT=587
SMTP_USERNAME=your-email@gmail.com
SMTP_PASSWORD=your-app-password
```

**Если FILE_UPLOAD_ENABLE=true:**
```env
UPLOAD_DIR=./uploads
MAX_FILE_SIZE=10485760
```

### 4. Запустите Сервер

```bash
go run cmd/server/main.go
```

### 5. Проверьте Логи

Сервер залогирует, какие функции включены:

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

## Тестирование Функций

### Тест WebSocket (если включен)

Посетите: `http://localhost:8080/websocket`

### Тест Swagger (если включен)

Посетите: `http://localhost:8080/swagger/index.html`

### Тест Загрузки Файлов (если включена)

Посетите: `http://localhost:8080/upload`

### Тест i18n (если включен)

Посетите: `http://localhost:8080/i18n`

## Устранение Неполадок

### Функция Не Работает

1. Проверьте файл `.env` - включена ли функция?
2. Проверьте логи - успешно ли инициализировалась функция?
3. Проверьте зависимости - запущен ли Redis (для функций Redis/Queue)?

### Маршруты Недоступны

Если маршруты возвращают 404:
- Функция отключена в `.env`
- Перезапустите сервер после изменения `.env`

### Предупреждения в Логах

Предупреждения типа "Redis not available" нормальны, когда функции отключены. Чтобы их убрать, убедитесь, что функция установлена в `false` в `.env`.

## Следующие Шаги

- Прочитайте полную документацию: [FEATURES.ru.md](FEATURES.ru.md)
- Настройте конкретные функции:
  - [WebSocket](WEBSOCKET.ru.md)
  - [Redis Cache](CACHE.ru.md)
  - [File Upload](UPLOAD.ru.md)
  - [i18n](I18N.ru.md)
  - [Queue](QUEUE.ru.md)
  - [Email](EMAIL.ru.md)
