# DMMVC - Lightweight MVC Web Framework

**DMMVC** — минималистичный MVC веб-фреймворк на Go, готовый для создания любых веб-приложений.

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
- **Database**: SQLite / MySQL
- **Logger**: Logrus
- **Sessions**: gorilla/sessions

## Быстрый старт

### Установка

```bash
# Клонируйте репозиторий
git clone https://github.com/dedomorozoff/dmmvc
cd dmmvc

# Установите зависимости
go mod tidy

# Запустите сервер
go run cmd/server/main.go
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
│   └── server/              # Точка входа приложения
│       └── main.go
├── internal/
│   ├── controllers/         # HTTP контроллеры
│   │   ├── auth_controller.go
│   │   ├── home_controller.go
│   │   └── user_controller.go
│   ├── database/            # Подключение к БД
│   │   ├── db.go
│   │   └── seeder.go
│   ├── logger/              # Логирование
│   │   └── logger.go
│   ├── middleware/          # Middleware
│   │   ├── auth.go
│   │   └── logger.go
│   ├── models/              # Модели данных
│   │   └── user.go
│   └── routes/              # Маршруты
│       └── routes.go
├── static/
│   ├── css/                 # Стили
│   │   └── style.css
│   └── js/                  # JavaScript
│       └── app.js
├── templates/
│   ├── layouts/             # Layouts
│   │   └── base.html
│   ├── partials/            # Переиспользуемые компоненты
│   │   ├── header.html
│   │   ├── footer.html
│   │   └── sidebar.html
│   └── pages/               # Страницы
│       ├── home.html
│       ├── login.html
│       └── dashboard.html
├── .env.example             # Пример конфигурации
├── go.mod                   # Go модули
└── README.md                # Документация
```

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

## Roadmap

- [ ] CLI инструмент для генерации кода
- [ ] Поддержка PostgreSQL
- [ ] WebSocket поддержка
- [ ] API документация (Swagger)
- [ ] Кеширование (Redis)
- [ ] Очереди задач
- [ ] Email отправка
- [ ] File upload helper
- [ ] Локализация (i18n)
