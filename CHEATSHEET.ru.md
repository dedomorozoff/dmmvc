# DMMVC Шпаргалка

Быстрый справочник по основным задачам DMMVC.

## Установка

```bash
# Linux/macOS
git clone https://github.com/dedomorozoff/dmmvc
cd dmmvc
chmod +x scripts/install.sh
./scripts/install.sh

# Windows
git clone https://github.com/dedomorozoff/dmmvc
cd dmmvc
scripts\install.bat

# Через go install
go install github.com/dedomorozoff/dmmvc/cmd/cli@latest
```

## Быстрый старт

```bash
# Linux/macOS
./scripts/quickstart.sh

# Windows
scripts\quickstart.bat

# Вручную
go mod tidy
go run cmd/server/main.go
```

## Создание нового проекта

```bash
# Linux/macOS
./scripts/create-project.sh my-app

# Windows
scripts\create-project.bat my-app

cd my-app
go mod tidy
go run cmd/server/main.go
```

## Команды CLI

```bash
# Показать справку
dmmvc --help

# Создать CRUD
dmmvc make:crud Product
dmmvc make:crud User --fields="name:string,email:string,age:int"

# Создать контроллер
dmmvc make:controller About
dmmvc make:controller api/Users

# Создать модель
dmmvc make:model Category
dmmvc make:model Post --migration

# Создать представление
dmmvc make:view products/index
dmmvc make:view users/profile

# Список ресурсов
dmmvc list
dmmvc list controllers
dmmvc list models
```

## Разработка

```bash
# Запустить сервер
go run cmd/server/main.go

# Собрать
go build -o server cmd/server/main.go

# Запустить тесты
go test ./...
go test -v ./...
go test -cover ./...

# Форматировать код
go fmt ./...

# Проверить код
go vet ./...

# Обновить зависимости
go get -u ./...
go mod tidy
```

## Использование dev.sh (Linux/macOS)

```bash
# Установить зависимости
./scripts/dev.sh install

# Собрать CLI и сервер
./scripts/dev.sh build

# Запустить сервер
./scripts/dev.sh run

# Запустить тесты
./scripts/dev.sh test

# Форматировать код
./scripts/dev.sh fmt

# Проверить код
./scripts/dev.sh lint

# Генерировать Swagger
./scripts/dev.sh swagger

# Очистить
./scripts/dev.sh clean
```

## База данных

```bash
# SQLite (по умолчанию)
DB_TYPE=sqlite
DB_DSN=dmmvc.db

# MySQL
DB_TYPE=mysql
DB_DSN=user:password@tcp(localhost:3306)/dbname?charset=utf8mb4&parseTime=True

# PostgreSQL
DB_TYPE=postgres
DB_DSN=host=localhost user=postgres password=postgres dbname=dmmvc port=5432 sslmode=disable
```

## Операции GORM

```go
// Создать
user := models.User{Username: "john", Email: "john@example.com"}
database.DB.Create(&user)

// Найти по ID
var user models.User
database.DB.First(&user, 1)

// Найти по условию
database.DB.Where("username = ?", "john").First(&user)

// Найти все
var users []models.User
database.DB.Find(&users)

// Обновить
database.DB.Model(&user).Update("Email", "new@example.com")
database.DB.Model(&user).Updates(models.User{Email: "new@example.com", Age: 30})

// Удалить
database.DB.Delete(&user)
database.DB.Where("age < ?", 18).Delete(&models.User{})
```

## Маршруты

```go
// Публичные маршруты
r.GET("/", controllers.Index)
r.POST("/api/data", controllers.CreateData)

// Защищенные маршруты
authorized := r.Group("/")
authorized.Use(middleware.AuthRequired())
{
    authorized.GET("/dashboard", controllers.Dashboard)
    authorized.POST("/profile", controllers.UpdateProfile)
}

// API маршруты
api := r.Group("/api")
{
    api.GET("/users", controllers.GetUsers)
    api.POST("/users", controllers.CreateUser)
    api.PUT("/users/:id", controllers.UpdateUser)
    api.DELETE("/users/:id", controllers.DeleteUser)
}
```

## Контроллеры

```go
package controllers

import (
    "net/http"
    "github.com/gin-gonic/gin"
)

// HTML ответ
func Index(c *gin.Context) {
    c.HTML(http.StatusOK, "pages/index.html", gin.H{
        "title": "Главная",
        "data": "Данные",
    })
}

// JSON ответ
func GetUsers(c *gin.Context) {
    users := []User{} // получить из БД
    c.JSON(http.StatusOK, gin.H{
        "users": users,
    })
}

// Получить параметр URL
func GetUser(c *gin.Context) {
    id := c.Param("id")
    // ...
}

// Получить query параметр
func Search(c *gin.Context) {
    query := c.Query("q")
    // ...
}

// Привязать JSON
func CreateUser(c *gin.Context) {
    var user User
    if err := c.ShouldBindJSON(&user); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    // ...
}
```

## Middleware

```go
// Собственный middleware
func MyMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        // До запроса
        c.Next()
        // После запроса
    }
}

// Использовать глобально
r.Use(MyMiddleware())

// Использовать для группы
authorized := r.Group("/")
authorized.Use(middleware.AuthRequired())
```

## Шаблоны

```html
{{define "pages/mypage.html"}}
{{template "layouts/base.html" .}}
{{define "content"}}
<div class="container">
    <h1>{{.title}}</h1>
    
    {{/* Переменные */}}
    <p>{{.message}}</p>
    
    {{/* Условия */}}
    {{if .user}}
        <p>Добро пожаловать, {{.user.Name}}</p>
    {{else}}
        <p>Пожалуйста, войдите</p>
    {{end}}
    
    {{/* Циклы */}}
    {{range .items}}
        <div>{{.Name}}</div>
    {{end}}
    
    {{/* Включить partial */}}
    {{template "partials/header.html" .}}
</div>
{{end}}
{{end}}
```

## Сессии

```go
// Получить сессию
session := sessions.Default(c)

// Установить значение
session.Set("user_id", user.ID)
session.Save()

// Получить значение
userID := session.Get("user_id")

// Удалить значение
session.Delete("user_id")
session.Save()

// Очистить сессию
session.Clear()
session.Save()
```

## Docker

```bash
# Собрать образ
docker build -t dmmvc .
./scripts/docker-build.sh

# Запустить контейнер
docker run -p 8080:8080 dmmvc

# Запустить с .env
docker run -p 8080:8080 --env-file .env dmmvc

# Запустить в фоне
docker run -d -p 8080:8080 --name dmmvc-app dmmvc

# Просмотр логов
docker logs dmmvc-app

# Остановить контейнер
docker stop dmmvc-app

# Docker Compose
docker-compose up
docker-compose up -d
docker-compose down
```

## Тестирование

```bash
# Запустить все тесты
go test ./...

# Подробный вывод
go test -v ./...

# С покрытием
go test -cover ./...
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out

# Конкретный пакет
go test ./internal/controllers

# Используя скрипт
./scripts/test.sh
./scripts/test.sh --coverage
./scripts/test.sh --verbose
./scripts/test.sh --bench
```

## Swagger

```bash
# Генерировать документацию
swag init -g cmd/server/main.go -o docs/swagger
make swagger
./scripts/dev.sh swagger

# Доступ к UI
# http://localhost:8080/swagger/index.html
```

## Переменные окружения

```env
# Сервер
PORT=8080
GIN_MODE=debug|release

# База данных
DB_TYPE=sqlite|mysql|postgres
DB_DSN=строка_подключения

# Безопасность
SESSION_SECRET=ваш-секретный-ключ

# Логирование
LOG_LEVEL=debug|info|warn|error
LOG_FILE=app.log

# Разработка
DEBUG=true|false
```

## Частые проблемы

```bash
# Команда не найдена: dmmvc
export PATH=$PATH:$(go env GOPATH)/bin  # Linux/macOS
set PATH=%PATH%;%GOPATH%\bin            # Windows

# Отказано в доступе (Linux/macOS)
chmod +x scripts/*.sh

# Ошибки модулей
go clean -modcache
go mod download
go mod tidy

# Порт уже используется
# Измените PORT в .env или завершите процесс
lsof -ti:8080 | xargs kill -9  # Linux/macOS
netstat -ano | findstr :8080   # Windows
```

## Полезные ссылки

- [Документация](docs/)
- [Руководство по установке](docs/INSTALLATION.ru.md)
- [Справочник CLI](docs/CLI.ru.md)
- [Примеры](docs/EXAMPLES.ru.md)
- [Скрипты](scripts/README.md)

## Карточка быстрого доступа

```
Установка:      ./scripts/install.sh | scripts\install.bat
Новый проект:   ./scripts/create-project.sh my-app
Быстрый старт:  ./scripts/quickstart.sh
Запуск сервера: go run cmd/server/main.go
Сборка CLI:     go build -o dmmvc cmd/cli/main.go
Запуск тестов:  go test ./...
Генерация CRUD: dmmvc make:crud Product
Форматирование: go fmt ./...
```
