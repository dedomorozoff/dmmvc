[English](README.md) | **Русский** | [Шпаргалка](CHEATSHEET.ru.md)

# DMMVC - Lightweight MVC Web Framework

**DMMVC** — минималистичный MVC веб-фреймворк на Go, готовый для создания любых веб-приложений.

> **Документация**: [docs/](docs/)

## Особенности

### Архитектура
- **MVC паттерн** - Model-View-Controller
- **Модульная структура** - Расширяемая архитектура
- **Middleware система** - Обработка запросов
- **Роутинг** - Система маршрутов

### Безопасность
- **Аутентификация** - Система авторизации
- **Сессии** - Безопасное управление сессиями
- **Хеширование паролей** - bcrypt для безопасного хранения
- **CSRF защита** - Готова к интеграции

### База данных
- **GORM ORM** - ORM для работы с БД
- **Миграции** - Автоматическое создание структуры БД
- **Поддержка SQLite** - Для быстрого старта
- **Поддержка MySQL** - Для production окружения
- **Поддержка PostgreSQL** - Мощная реляционная БД

### Шаблоны
- **Go Templates** - Шаблонизатор
- **Layouts** - Система layouts для переиспользования
- **Partials** - Компоненты для модульности
- **Статические файлы** - Автоматическая раздача CSS/JS

### Логирование
- **Logrus** - Структурированное логирование
- **Request logging** - Автоматическое логирование запросов
- **Error tracking** - Отслеживание ошибок
- **Panic recovery** - Обработка паник

## Технологический стек

- **Backend**: Go 1.20+
- **Web Framework**: Gin
- **ORM**: GORM
- **Database**: SQLite / MySQL / PostgreSQL
- **Logger**: Logrus
- **Sessions**: gorilla/sessions

## Быстрый старт

### Установка

#### Вариант 1: Автоматическая установка (рекомендуется)

**Linux/macOS:**
```bash
# Клонируйте репозиторий
git clone https://github.com/dedomorozoff/dmmvc
cd dmmvc

# Сделайте скрипт исполняемым и запустите
chmod +x scripts/install.sh
./scripts/install.sh
```

**Windows:**
```bash
# Клонируйте репозиторий
git clone https://github.com/dedomorozoff/dmmvc
cd dmmvc

# Запустите установочный скрипт
scripts\install.bat
```

#### Вариант 2: Установка через Go

```bash
# Установить CLI глобально
go install github.com/dedomorozoff/dmmvc/cmd/cli@latest

# Или локально
git clone https://github.com/dedomorozoff/dmmvc
cd dmmvc
make install-go
```

#### Вариант 3: Ручная установка

```bash
# Клонируйте репозиторий
git clone https://github.com/dedomorozoff/dmmvc
cd dmmvc

# Установите зависимости
go mod tidy

# Соберите CLI
make build

# Установите глобально (опционально)
make install

# Запустите сервер
go run cmd/server/main.go
```

### Создание нового проекта

После установки CLI вы можете создать новый проект:

**Linux/macOS:**
```bash
# Создать новый проект
chmod +x scripts/create-project.sh
./scripts/create-project.sh my-app

# Перейти в проект
cd my-app

# Установить зависимости
go mod tidy

# Запустить сервер
go run cmd/server/main.go
```

**Windows:**
```bash
# Создать новый проект
scripts\create-project.bat my-app

# Перейти в проект
cd my-app

# Установить зависимости
go mod tidy

# Запустить сервер
go run cmd/server/main.go
```

### Помощник разработки (Linux/macOS)

Используйте скрипт `dev.sh` для упрощения разработки:

```bash
# Показать доступные команды
./scripts/dev.sh help

# Установить зависимости
./scripts/dev.sh install

# Собрать CLI и сервер
./scripts/dev.sh build

# Запустить сервер разработки
./scripts/dev.sh run

# Запустить тесты
./scripts/dev.sh test

# Форматировать код
./scripts/dev.sh fmt

# Генерировать Swagger документацию
./scripts/dev.sh swagger
```

### Первый вход

Откройте браузер: **http://localhost:8080**

Учетные данные по умолчанию:
- **Логин**: `admin`
- **Пароль**: `admin`

**Важно**: Смените пароль после первого входа!

## Структура проекта

```
dmmvc/
├── cmd/
│   ├── cli/                 # CLI инструмент
│   └── server/              # Точка входа приложения
├── internal/
│   ├── controllers/         # HTTP контроллеры
│   ├── database/            # Подключение к БД
│   ├── logger/              # Логирование
│   ├── middleware/          # Middleware
│   ├── models/              # Модели данных
│   └── routes/              # Маршруты
├── static/
│   ├── css/                 # Стили
│   └── js/                  # JavaScript
├── templates/
│   ├── layouts/             # Layouts
│   ├── partials/            # Компоненты
│   └── pages/               # Страницы
├── docs/                    # Документация
├── docker/                  # Docker конфигурация
├── scripts/                 # Утилиты
├── .env.example             # Пример конфигурации
├── Dockerfile               # Docker образ
├── Makefile                 # Команды сборки
└── README.md                # Документация
```

**Полная документация**: [docs/](docs/)

## Конфигурация

Создайте файл `.env` в корне проекта:

```env
# Server Settings
PORT=8080
GIN_MODE=debug

# Database Settings
DB_TYPE=sqlite
DB_DSN=dmmvc.db

# For MySQL:
# DB_TYPE=mysql
# DB_DSN=user:password@tcp(localhost:3306)/dmmvc?charset=utf8mb4&parseTime=True&loc=Local

# For PostgreSQL:
# DB_TYPE=postgres
# DB_DSN=host=localhost user=postgres password=postgres dbname=dmmvc port=5432 sslmode=disable

# Security Settings
SESSION_SECRET=your-super-secret-key-change-this-in-production

# Logging Settings
LOG_LEVEL=info
LOG_FILE=dmmvc.log

# Development Settings
DEBUG=true
```

## Создание нового контроллера

```go
package controllers

import (
    "net/http"
    "github.com/gin-gonic/gin"
)

func MyController(c *gin.Context) {
    c.HTML(http.StatusOK, "pages/mypage.html", gin.H{
        "title": "My Page",
        "data": "Some data",
    })
}
```

## Создание новой модели

```go
package models

import "gorm.io/gorm"

type MyModel struct {
    gorm.Model
    Name        string `json:"name"`
    Description string `json:"description"`
}
```

## Добавление маршрута

```go
// В файле internal/routes/routes.go
authorized.GET("/mypage", controllers.MyController)
```

## Создание шаблона

```html
{{define "pages/mypage.html"}}
{{template "layouts/base.html" .}}
{{define "content"}}
<div class="container">
    <h1>{{.title}}</h1>
    <p>{{.data}}</p>
</div>
{{end}}
{{end}}
```

## Middleware

### Создание собственного middleware

```go
func MyMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        // Код до обработки запроса
        c.Next()
        // Код после обработки запроса
    }
}
```

### Использование middleware

```go
// Глобально
r.Use(MyMiddleware())

// Для группы маршрутов
authorized := r.Group("/")
authorized.Use(MyMiddleware())
```

## Работа с базой данных

```go
import "dmmvc/internal/database"

// Создание записи
user := models.User{Username: "john", Email: "john@example.com"}
database.DB.Create(&user)

// Поиск
var user models.User
database.DB.First(&user, 1) // По ID
database.DB.Where("username = ?", "john").First(&user)

// Обновление
database.DB.Model(&user).Update("Email", "newemail@example.com")

// Удаление
database.DB.Delete(&user)
```

## Тестирование

```bash
# Запуск тестов
go test ./...

# С покрытием
go test -cover ./...
```

## Сборка для production

```bash
# Сборка бинарника
go build -o dmmvc cmd/server/main.go

# Запуск
./dmmvc
```

## Docker

```dockerfile
FROM golang:1.20-alpine
WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o dmmvc cmd/server/main.go
EXPOSE 8080
CMD ["./dmmvc"]
```

## Лицензия

MIT License

## Расширенные возможности

### CLI Инструмент

DMMVC включает мощный CLI инструмент для генерации кода:

```bash
# Собрать CLI
make build

# Создать CRUD для продуктов
dmmvc make:crud Product

# Создать контроллер
dmmvc make:controller About

# Создать модель
dmmvc make:model Category --migration

# Показать все ресурсы
dmmvc list
```

**Полная документация**: [docs/CLI.ru.md](docs/CLI.ru.md)

### Поддержка баз данных

DMMVC поддерживает три типа баз данных:

- **SQLite** - Для разработки и небольших проектов
- **MySQL** - Популярная реляционная БД
- **PostgreSQL** - Мощная open-source БД

**Документация**: [docs/POSTGRESQL.ru.md](docs/POSTGRESQL.ru.md)

### WebSocket поддержка

Реал-тайм двунаправленная коммуникация:

```go
// WebSocket эндпоинт доступен по /ws
// Демо страница по /websocket
```

**Документация**: [docs/WEBSOCKET.ru.md](docs/WEBSOCKET.ru.md)

### API Документация (Swagger)

Автоматическая генерация документации API:

```bash
# Генерация Swagger документации
make swagger

# Доступ к Swagger UI
# http://localhost:8080/swagger/index.html
```

**Документация**: [docs/SWAGGER.ru.md](docs/SWAGGER.ru.md)

### Redis Кеширование

Улучшение производительности с Redis кешированием:

```go
// Ручное кеширование
cache.Set("key", "value", 5*time.Minute)
value, _ := cache.Get("key")

// Cache middleware
r.GET("/api/data", middleware.CacheMiddleware(5*time.Minute), handler)
```

**Документация**: [docs/CACHE.ru.md](docs/CACHE.ru.md)

### Очереди задач

Асинхронная обработка фоновых задач:

```go
// Создать и добавить задачу в очередь
task, _ := queue.NewEmailDeliveryTask("user@example.com", "Тема", "Текст")
queue.EnqueueTask(task)

// Отложенная задача
queue.EnqueueTaskIn(task, 5*time.Minute)
```

**Документация**: [docs/QUEUE.ru.md](docs/QUEUE.ru.md)

### Отправка Email

Отправка email через SMTP с HTML шаблонами:

```go
// Отправить email напрямую
email.Send("user@example.com", "Тема", "<h1>Привет</h1>")

// Отправить с шаблоном
email.WelcomeEmail("user@example.com", "Иван Иванов")
```

**Документация**: [docs/EMAIL.ru.md](docs/EMAIL.ru.md)

### Загрузка файлов

Загрузка и обработка файлов с изменением размера изображений:

```go
// Загрузить файл
fileInfo, _ := upload.UploadFile(file)

// Создать миниатюру
upload.CreateThumbnail(fileInfo.Path, 300, 300)
```

**Документация**: [docs/UPLOAD.ru.md](docs/UPLOAD.ru.md)

## Документация

### Начало работы
- [Установка](docs/INSTALLATION.ru.md) - Подробное руководство по установке
- [CLI Инструмент](docs/CLI.ru.md) - Команды генерации кода
- [Примеры](docs/EXAMPLES.ru.md) - Примеры использования

### Возможности
- [PostgreSQL](docs/POSTGRESQL.ru.md) - Настройка PostgreSQL
- [WebSocket](docs/WEBSOCKET.ru.md) - Реал-тайм коммуникация
- [Swagger API](docs/SWAGGER.ru.md) - API документация
- [Redis Cache](docs/CACHE.ru.md) - Кеширование с Redis
- [Очереди задач](docs/QUEUE.ru.md) - Фоновая обработка задач
- [Отправка Email](docs/EMAIL.ru.md) - SMTP доставка email
- [Загрузка файлов](docs/UPLOAD.ru.md) - Загрузка и обработка файлов
- [Интернационализация](docs/I18N.ru.md) - Поддержка нескольких языков

### Развертывание
- [Развертывание](docs/DEPLOYMENT.ru.md) - Production развертывание
- [Скрипты](scripts/README.md) - Утилиты для разработки
