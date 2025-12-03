# Обновление Зависимостей

Это руководство объясняет, как поддерживать зависимости фреймворка DMMVC в актуальном состоянии.

## Быстрое Обновление

Чтобы обновить все зависимости до последних версий:

```bash
# Обновить все зависимости
go get -u ./...

# Очистить и загрузить недостающие зависимости
go mod tidy

# Проверить, что всё компилируется
go build -o dmmvc cmd/server/main.go
```

## Проверка Обновлений

### Проверить устаревшие пакеты

```bash
# Список всех зависимостей
go list -m all

# Проверить доступные обновления
go list -u -m all
```

### Проверить конкретный пакет

```bash
go list -m -u github.com/gin-gonic/gin
```

## Обновление Конкретных Пакетов

### Обновить один пакет

```bash
go get -u github.com/gin-gonic/gin@latest
```

### Обновить до конкретной версии

```bash
go get github.com/gin-gonic/gin@v1.9.0
```

### Обновить мажорную версию

```bash
go get github.com/gin-gonic/gin@v2
```

## Тестирование После Обновлений

После обновления зависимостей всегда тестируйте приложение:

```bash
# Запустить тесты
go test ./...

# Собрать приложение
go build -o dmmvc cmd/server/main.go

# Запустить сервер
./dmmvc
# или
go run cmd/server/main.go
```

## Частые Проблемы

### Критические Изменения

Обновления мажорных версий могут включать критические изменения. Всегда проверяйте:
- Changelog/release notes пакета
- Руководства по миграции
- Документацию API

### Конфликты Зависимостей

Если возникают конфликты:

```bash
# Удалить go.sum и попробовать снова
rm go.sum
go mod tidy

# Или вернуться к рабочему состоянию
git checkout go.mod go.sum
```

### Директория Vendor

Если используете vendoring:

```bash
# Обновить директорию vendor
go mod vendor

# Проверить директорию vendor
go mod verify
```

## Рекомендуемый График Обновлений

- **Обновления безопасности**: Немедленно
- **Минорные обновления**: Ежемесячно
- **Мажорные обновления**: Ежеквартально (с тестированием)

## Ключевые Зависимости

### Основной Фреймворк
- `github.com/gin-gonic/gin` - Веб-фреймворк
- `gorm.io/gorm` - ORM
- `github.com/gin-contrib/sessions` - Управление сессиями

### Драйверы Баз Данных
- `github.com/glebarez/sqlite` - SQLite (чистый Go)
- `gorm.io/driver/mysql` - MySQL
- `gorm.io/driver/postgres` - PostgreSQL

### Опциональные Функции
- `github.com/gorilla/websocket` - Поддержка WebSocket
- `github.com/redis/go-redis/v9` - Redis клиент
- `github.com/swaggo/gin-swagger` - Swagger документация
- `github.com/hibiken/asynq` - Очередь задач
- `github.com/nfnt/resize` - Обработка изображений

### Утилиты
- `github.com/joho/godotenv` - Переменные окружения
- `github.com/sirupsen/logrus` - Логирование
- `golang.org/x/crypto` - Криптография

## Автоматические Обновления

### Использование Dependabot (GitHub)

Создайте `.github/dependabot.yml`:

```yaml
version: 2
updates:
  - package-ecosystem: "gomod"
    directory: "/"
    schedule:
      interval: "weekly"
    open-pull-requests-limit: 10
```

### Использование Renovate

Создайте `renovate.json`:

```json
{
  "extends": ["config:base"],
  "packageRules": [
    {
      "matchUpdateTypes": ["minor", "patch"],
      "automerge": true
    }
  ]
}
```

## Откат

Если обновление вызывает проблемы:

```bash
# Вернуться к предыдущим версиям
git checkout go.mod go.sum

# Или вручную отредактировать go.mod и выполнить
go mod tidy
```

## Лучшие Практики

1. **Читайте changelog** перед обновлением
2. **Тщательно тестируйте** после обновлений
3. **Обновляйте регулярно**, чтобы избежать больших скачков
4. **Храните go.sum** в системе контроля версий
5. **Документируйте критические изменения** в вашем CHANGELOG
6. **Используйте семантическое версионирование** для своих пакетов

## Ресурсы

- [Go Modules Reference](https://go.dev/ref/mod)
- [Go Module Versioning](https://go.dev/doc/modules/version-numbers)
- [Semantic Versioning](https://semver.org/)
