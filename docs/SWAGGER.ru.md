[English](SWAGGER.md) | **Русский**

# Swagger API Документация

DMMVC включает встроенную поддержку Swagger/OpenAPI документации, позволяя автоматически генерировать интерактивную документацию API.

## Быстрый старт

### 1. Генерация документации

```bash
make swagger
```

Или вручную:

```bash
swag init -g cmd/server/main.go -o docs/swagger --parseDependency --parseInternal
```

### 2. Запуск сервера

```bash
make run
```

### 3. Доступ к Swagger UI

Откройте в браузере: **http://localhost:8080/swagger/index.html**

## Документирование вашего API

### Базовый пример контроллера

```go
// GetUser godoc
// @Summary Получить пользователя по ID
// @Description Получить детали пользователя по его ID
// @Tags users
// @Accept json
// @Produce json
// @Param id path int true "ID пользователя"
// @Success 200 {object} APIResponse{data=models.User}
// @Failure 404 {object} APIResponse
// @Router /api/users/{id} [get]
// @Security SessionAuth
func GetUser(c *gin.Context) {
    // Ваша реализация
}
```

### Теги аннотаций

- `@Summary` - Краткое описание
- `@Description` - Подробное описание
- `@Tags` - Группировка эндпоинтов по тегам
- `@Accept` - Тип контента запроса (json, xml и т.д.)
- `@Produce` - Тип контента ответа
- `@Param` - Определение параметра
- `@Success` - Успешный ответ
- `@Failure` - Ответ с ошибкой
- `@Router` - Путь маршрута и HTTP метод
- `@Security` - Схема безопасности

### Типы параметров

```go
// Параметр пути
// @Param id path int true "ID пользователя"

// Параметр запроса
// @Param page query int false "Номер страницы"

// Параметр тела запроса
// @Param user body UserCreateRequest true "Данные пользователя"

// Параметр заголовка
// @Param Authorization header string true "Bearer токен"
```

### Примеры ответов

```go
// Простой ответ
// @Success 200 {object} models.User

// Ответ с вложенными данными
// @Success 200 {object} APIResponse{data=models.User}

// Ответ с массивом
// @Success 200 {object} APIResponse{data=[]models.User}

// Несколько кодов статуса
// @Success 200 {object} APIResponse
// @Failure 400 {object} APIResponse
// @Failure 404 {object} APIResponse
// @Failure 500 {object} APIResponse
```

## Структура ответа API

Стандартный формат ответа API:

```go
type APIResponse struct {
    Success bool        `json:"success" example:"true"`
    Message string      `json:"message,omitempty" example:"Operation successful"`
    Data    interface{} `json:"data,omitempty"`
    Error   string      `json:"error,omitempty" example:"Error message"`
}
```

## Документирование моделей

Добавьте Swagger аннотации к вашим моделям:

```go
// User модель
// @Description Информация об учетной записи пользователя
type User struct {
    ID       uint   `json:"id" example:"1"`
    Username string `json:"username" example:"john_doe"`
    Email    string `json:"email" example:"john@example.com"`
    Role     string `json:"role" example:"user"`
}
```

## Модели запросов

Документируйте структуры запросов:

```go
type UserCreateRequest struct {
    Username string `json:"username" binding:"required" example:"john_doe"`
    Email    string `json:"email" binding:"required,email" example:"john@example.com"`
    Password string `json:"password" binding:"required,min=6" example:"password123"`
}
```

## Схемы безопасности

DMMVC использует сессионную аутентификацию по умолчанию:

```go
// @securityDefinitions.apikey SessionAuth
// @in cookie
// @name session_token
```

Чтобы требовать аутентификацию для эндпоинта:

```go
// @Security SessionAuth
```

## Общая информация об API

Настраивается в `cmd/server/main.go`:

```go
// @title DMMVC API
// @version 1.0
// @description API легковесного MVC веб-фреймворка
// @termsOfService http://swagger.io/terms/

// @contact.name Поддержка API
// @contact.email support@dmmvc.com

// @license.name MIT
// @license.url https://opensource.org/licenses/MIT

// @host localhost:8080
// @BasePath /
```

## Примеры API эндпоинтов

DMMVC включает примеры API эндпоинтов в `internal/controllers/api_example.go`:

- `GET /api/users` - Список всех пользователей
- `GET /api/users/:id` - Получить пользователя по ID
- `POST /api/users` - Создать нового пользователя
- `DELETE /api/users/:id` - Удалить пользователя

## Тестирование API

Используйте Swagger UI для тестирования вашего API:

1. Откройте http://localhost:8080/swagger/index.html
2. Кликните на эндпоинт
3. Нажмите "Try it out"
4. Заполните параметры
5. Нажмите "Execute"

## Регенерация документации

После внесения изменений в ваш API:

```bash
make swagger
```

Документация будет автоматически обновлена.

## Кастомизация

### Кастомная тема

Отредактируйте `internal/routes/routes.go`:

```go
url := ginSwagger.URL("http://localhost:8080/swagger/doc.json")
r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
```

### Скрыть Swagger в продакшене

```go
if os.Getenv("GIN_MODE") != "release" {
    r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
```

## Лучшие практики

1. **Документируйте все публичные API** - Каждый API эндпоинт должен иметь Swagger аннотации
2. **Используйте примеры** - Добавляйте теги `example` для помощи пользователям
3. **Группируйте по тегам** - Используйте `@Tags` для логической организации эндпоинтов
4. **Документируйте ошибки** - Включайте все возможные ответы с ошибками
5. **Держите в актуальном состоянии** - Запускайте `make swagger` после изменений API
6. **Используйте стандартные ответы** - Придерживайтесь структуры APIResponse для консистентности

## Решение проблем

### Документация не обновляется

```bash
# Очистить и регенерировать
rm -rf docs/swagger
make swagger
```

### Swagger UI не загружается

Проверьте, что импорт присутствует в `internal/routes/routes.go`:

```go
_ "dmmvc/docs/swagger"
```

### Ошибки определения типов

Используйте флаги `--parseDependency --parseInternal`:

```bash
swag init -g cmd/server/main.go -o docs/swagger --parseDependency --parseInternal
```

## Ресурсы

- [Документация Swaggo](https://github.com/swaggo/swag)
- [Спецификация OpenAPI](https://swagger.io/specification/)
- [Swagger UI](https://swagger.io/tools/swagger-ui/)

## Следующие шаги

- Добавить больше API эндпоинтов
- Документировать существующие контроллеры
- Кастомизировать форматы ответов
- Добавить версионирование API
- Экспортировать в Postman/Insomnia
