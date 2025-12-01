[English](CLI.md) | **Русский** | [📚 Docs](README.md)

# DMMVC CLI Tool

Инструмент командной строки для генерации кода в DMMVC фреймворке.

## Установка

### Сборка CLI

```bash
make build
```

Это создаст исполняемый файл `dmmvc.exe` в корне проекта.

### Глобальная установка

```bash
make install
```

После этого команда `dmmvc` будет доступна глобально.

## Использование

### Основные команды

#### Создание контроллера

```bash
# Простой контроллер
dmmvc make:controller Product

# Resource контроллер с CRUD методами
dmmvc make:controller Product --resource
```

#### Создание модели

```bash
# Простая модель
dmmvc make:model Product

# Модель с подсказкой по миграции
dmmvc make:model Product --migration
```

#### Создание middleware

```bash
dmmvc make:middleware RateLimit
```

#### Создание страницы

```bash
dmmvc make:page about
```

#### Создание полного CRUD

```bash
# Создает модель, контроллер и все страницы для CRUD
dmmvc make:crud Product
```

Эта команда создаст:
- Модель: `internal/models/product.go`
- Resource контроллер: `internal/controllers/product_controller.go`
- Страницы:
  - `templates/pages/product/index.html` - список
  - `templates/pages/product/show.html` - просмотр
  - `templates/pages/product/create.html` - создание
  - `templates/pages/product/edit.html` - редактирование

#### Список ресурсов

```bash
# Показать все контроллеры, модели, middleware и страницы
dmmvc list
```

### Справка

```bash
# Показать справку
dmmvc --help
dmmvc -h

# Показать версию
dmmvc --version
dmmvc -v
```

## Примеры использования

### Пример 1: Создание простого контроллера

```bash
dmmvc make:controller About
```

Создаст файл `internal/controllers/about_controller.go`:

```go
package controllers

import (
    "net/http"
    "github.com/gin-gonic/gin"
)

func AboutController(c *gin.Context) {
    c.HTML(http.StatusOK, "pages/about.html", gin.H{
        "title": "About",
    })
}
```

### Пример 2: Создание CRUD для продуктов

```bash
# 1. Создаем полный CRUD
dmmvc make:crud Product

# 2. Редактируем модель, добавляем поля
# internal/models/product.go
type Product struct {
    gorm.Model
    Name        string  `json:"name" gorm:"not null"`
    Description string  `json:"description"`
    Price       float64 `json:"price" gorm:"not null"`
}

# 3. Добавляем миграцию в internal/database/db.go
db.AutoMigrate(&models.Product{})

# 4. Добавляем маршруты в internal/routes/routes.go
authorized.GET("/product", controllers.ProductControllerIndex)
authorized.GET("/product/:id", controllers.ProductControllerShow)
authorized.GET("/product/create", controllers.ProductControllerCreate)
authorized.POST("/product", controllers.ProductControllerStore)
authorized.GET("/product/:id/edit", controllers.ProductControllerEdit)
authorized.POST("/product/:id", controllers.ProductControllerUpdate)
authorized.POST("/product/:id/delete", controllers.ProductControllerDelete)

# 5. Настраиваем шаблоны в templates/pages/product/
```

### Пример 3: Создание middleware

```bash
dmmvc make:middleware RateLimit
```

Создаст файл `internal/middleware/rate_limit.go`:

```go
package middleware

import "github.com/gin-gonic/gin"

func RateLimit() gin.HandlerFunc {
    return func(c *gin.Context) {
        // Before request
        c.Next()
        // After request
    }
}
```

## Соглашения об именовании

### Контроллеры
- Входное имя: `Product` или `ProductController`
- Имя файла: `product_controller.go`
- Имя функции: `ProductController` (простой) или `ProductControllerIndex`, `ProductControllerShow` и т.д. (resource)

### Модели
- Входное имя: `Product`
- Имя файла: `product.go`
- Имя структуры: `Product`

### Middleware
- Входное имя: `RateLimit`
- Имя файла: `rate_limit.go`
- Имя функции: `RateLimit`

### Страницы
- Входное имя: `about`
- Имя файла: `about.html`
- Путь: `templates/pages/about.html`

## Структура файлов после генерации

```
dmmvc/
├── internal/
│   ├── controllers/
│   │   ├── product_controller.go    # make:controller Product --resource
│   │   └── about_controller.go      # make:controller About
│   ├── models/
│   │   └── product.go               # make:model Product
│   └── middleware/
│       └── rate_limit.go            # make:middleware RateLimit
└── templates/
    └── pages/
        ├── about.html               # make:page about
        └── product/                 # make:crud Product
            ├── index.html
            ├── show.html
            ├── create.html
            └── edit.html
```

## Советы

1. **Используйте make:crud** для быстрого создания полного CRUD функционала
2. **Используйте --resource** при создании контроллера, если планируете CRUD операции
3. **Используйте --migration** при создании модели, чтобы не забыть добавить миграцию
4. **Используйте dmmvc list** чтобы увидеть все созданные ресурсы
5. После генерации кода всегда проверяйте и настраивайте под свои нужды

## Разработка CLI

Если вы хотите добавить новые команды в CLI:

1. Откройте `cmd/cli/main.go`
2. Добавьте новый case в switch statement
3. Создайте функцию-обработчик
4. Обновите `printUsage()`
5. Пересоберите: `make build`

## Troubleshooting

### CLI не найден после установки

Убедитесь, что `%GOPATH%\bin` добавлен в PATH:

```bash
echo %PATH%
```

### Файл уже существует

CLI не перезаписывает существующие файлы. Если нужно пересоздать файл, удалите его вручную.

### Ошибки при создании файлов

Убедитесь, что вы запускаете CLI из корня проекта DMMVC.
