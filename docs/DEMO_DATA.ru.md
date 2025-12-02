# Демо данные

Этот документ описывает демо данные, доступные во фреймворке DMMVC.

## Демо пользователи

Фреймворк автоматически создает демо пользователей при первом запуске с поддержкой двух языков (английского и русского).

### Администратор

- **Имя пользователя**: `admin`
- **Email**: `admin@example.com`
- **Пароль**: `admin`
- **Роль**: `admin`

⚠️ **Важно**: Измените пароль по умолчанию после первого входа!

### Английские демо пользователи

| Имя пользователя | Email | Пароль | Роль |
|------------------|-------|--------|------|
| john_doe | john@example.com | password123 | user |
| jane_smith | jane@example.com | password123 | user |
| bob_johnson | bob@example.com | password123 | user |

### Русские демо пользователи

| Имя пользователя | Email | Пароль | Роль |
|------------------|-------|--------|------|
| ivan_ivanov | ivan@example.ru | password123 | user |
| maria_petrova | maria@example.ru | password123 | user |
| alexey_sidorov | alexey@example.ru | password123 | user |

## API сообщения

Все ответы API поддерживают интернационализацию (i18n) на основе заголовка `Accept-Language` или параметра запроса `lang`.

### Поддерживаемые языки

- **Английский** (`en`) - По умолчанию
- **Русский** (`ru`)

### Примеры API вызовов

#### Английский ответ

```bash
curl -H "Accept-Language: en" http://localhost:8080/api/users
```

Ответ:
```json
{
  "success": true,
  "data": [...]
}
```

#### Русский ответ

```bash
curl -H "Accept-Language: ru" http://localhost:8080/api/users
```

Ответ:
```json
{
  "success": true,
  "data": [...]
}
```

### Использование параметра запроса

```bash
# Английский
curl http://localhost:8080/api/users?lang=en

# Русский
curl http://localhost:8080/api/users?lang=ru
```

## Демо шаблоны email

Примеры тем и текстов email доступны на обоих языках:

### Приветственное письмо

**Английский**:
- Тема: "Welcome!"
- Текст: "Welcome to our service!"

**Русский**:
- Тема: "Добро пожаловать!"
- Текст: "Добро пожаловать в наш сервис!"

### Тестовый email

**Английский**:
- Тема: "Hello"
- Текст: "<h1>Hello World</h1>"

**Русский**:
- Тема: "Привет"
- Текст: "<h1>Привет, мир!</h1>"

## Отключение демо данных

Чтобы отключить автоматическое создание демо пользователей, закомментируйте следующую строку в `cmd/server/main.go`:

```go
// database.SeedDemoUsers()
```

## Ключи переводов

Все ключи переводов определены в:

- `locales/en.json` - Английские переводы
- `locales/ru.json` - Русские переводы

### Ключи переводов API

- `api.user.*` - Сообщения управления пользователями
- `api.email.*` - Сообщения email сервиса
- `api.queue.*` - Сообщения сервиса очередей
- `demo.*` - Переводы демо данных

Для получения дополнительной информации о i18n см. [I18N.md](I18N.md) и [I18N.ru.md](I18N.ru.md).

## Примеры использования в коде

### Использование переводов в контроллерах

```go
import "dmmvc/internal/i18n"

func MyController(c *gin.Context) {
    message := i18n.T(c, "api.user.created")
    c.JSON(http.StatusOK, gin.H{
        "message": message,
    })
}
```

### Добавление новых переводов

1. Добавьте ключ в `locales/en.json`:
```json
{
  "my.new.key": "My new message"
}
```

2. Добавьте перевод в `locales/ru.json`:
```json
{
  "my.new.key": "Мое новое сообщение"
}
```

3. Используйте в коде:
```go
message := i18n.T(c, "my.new.key")
```

## Тестирование переводов

### Через браузер

Измените язык в настройках браузера или используйте параметр URL:

```
http://localhost:8080/?lang=ru
http://localhost:8080/?lang=en
```

### Через API

Используйте заголовок `Accept-Language`:

```bash
# Русский
curl -H "Accept-Language: ru" http://localhost:8080/api/users

# Английский
curl -H "Accept-Language: en" http://localhost:8080/api/users
```

Или параметр запроса `lang`:

```bash
curl http://localhost:8080/api/users?lang=ru
```
