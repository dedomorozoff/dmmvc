[English](QUICKSTART.md) | **Русский**

# Быстрый старт DMMVC

## Установка и запуск

### 1. Установка зависимостей

```bash
go mod tidy
```

### 2. Запуск сервера

```bash
go run cmd/server/main.go
```

### 3. Открыть в браузере

```
http://localhost:8080
```

### 4. Вход в систему

- **Логин**: `admin`
- **Пароль**: `admin`

**Важно**: Смените пароль после первого входа!

## Структура проекта

```
dmmvc/
├── cmd/server/main.go          # Точка входа
├── internal/
│   ├── controllers/            # Контроллеры (HTTP обработчики)
│   ├── database/               # Работа с БД
│   ├── logger/                 # Логирование
│   ├── middleware/             # Middleware
│   ├── models/                 # Модели данных
│   └── routes/                 # Маршруты
├── static/                     # Статические файлы
│   ├── css/style.css
│   └── js/app.js
├── templates/                  # HTML шаблоны
│   ├── layouts/base.html
│   ├── partials/
│   └── pages/
├── .env                        # Конфигурация
└── go.mod                      # Зависимости
```

## Создание нового функционала

### 1. Создать модель

```go
// internal/models/post.go
package models

import "gorm.io/gorm"

type Post struct {
    gorm.Model
    Title   string `json:"title"`
    Content string `json:"content"`
    UserID  uint   `json:"user_id"`
}
```

### 2. Добавить миграцию

```go
// cmd/server/main.go
database.Migrate(&models.User{}, &models.Post{})
```

### 3. Создать контроллер

```go
// internal/controllers/post_controller.go
package controllers

import (
    "net/http"
    "github.com/gin-gonic/gin"
)

func PostList(c *gin.Context) {
    c.HTML(http.StatusOK, "pages/posts/list.html", gin.H{
        "title": "Посты",
    })
}
```

### 4. Добавить маршрут

```go
// internal/routes/routes.go
authorized.GET("/posts", controllers.PostList)
```

### 5. Создать шаблон

```html
<!-- templates/pages/posts/list.html -->
{{define "pages/posts/list.html"}}
{{template "layouts/base.html" .}}
{{define "content"}}
{{template "partials/header.html" .}}

<main class="main">
    <div class="container">
        <h1>Список постов</h1>
    </div>
</main>

{{template "partials/footer.html" .}}
{{end}}
{{end}}
```

## Полезные команды

```bash
# Запуск в режиме разработки
go run cmd/server/main.go

# Сборка бинарника
go build -o dmmvc cmd/server/main.go

# Запуск бинарника
./dmmvc

# Запуск тестов
go test ./...
```

## Настройка базы данных

### SQLite (по умолчанию)

```env
DB_TYPE=sqlite
DB_DSN=dmmvc.db
```

### MySQL

```env
DB_TYPE=mysql
DB_DSN=user:password@tcp(localhost:3306)/dmmvc?charset=utf8mb4&parseTime=True&loc=Local
```

## Готово!

Теперь у вас есть полностью рабочий MVC фреймворк, готовый к созданию любого веб-приложения!
