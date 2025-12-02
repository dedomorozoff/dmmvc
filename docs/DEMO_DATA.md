# Demo Data / Демо данные

This document describes the demo data available in DMMVC framework.

Этот документ описывает демо данные, доступные во фреймворке DMMVC.

## Demo Users / Демо пользователи

The framework automatically creates demo users on first run with bilingual support (English and Russian).

Фреймворк автоматически создает демо пользователей при первом запуске с поддержкой двух языков (английского и русского).

### Admin User / Администратор

- **Username**: `admin`
- **Email**: `admin@example.com`
- **Password**: `admin`
- **Role**: `admin`

⚠️ **Important**: Change the default password after first login!

⚠️ **Важно**: Измените пароль по умолчанию после первого входа!

### English Demo Users / Английские демо пользователи

| Username | Email | Password | Role |
|----------|-------|----------|------|
| john_doe | john@example.com | password123 | user |
| jane_smith | jane@example.com | password123 | user |
| bob_johnson | bob@example.com | password123 | user |

### Russian Demo Users / Русские демо пользователи

| Username | Email | Password | Role |
|----------|-------|----------|------|
| ivan_ivanov | ivan@example.ru | password123 | user |
| maria_petrova | maria@example.ru | password123 | user |
| alexey_sidorov | alexey@example.ru | password123 | user |

## API Messages / API сообщения

All API responses support internationalization (i18n) based on the `Accept-Language` header or `lang` query parameter.

Все ответы API поддерживают интернационализацию (i18n) на основе заголовка `Accept-Language` или параметра запроса `lang`.

### Supported Languages / Поддерживаемые языки

- **English** (`en`) - Default
- **Russian** (`ru`)

### Example API Calls / Примеры API вызовов

#### English Response / Английский ответ

```bash
curl -H "Accept-Language: en" http://localhost:8080/api/users
```

Response:
```json
{
  "success": true,
  "data": [...]
}
```

#### Russian Response / Русский ответ

```bash
curl -H "Accept-Language: ru" http://localhost:8080/api/users
```

Response:
```json
{
  "success": true,
  "data": [...]
}
```

### Using Query Parameter / Использование параметра запроса

```bash
# English
curl http://localhost:8080/api/users?lang=en

# Russian
curl http://localhost:8080/api/users?lang=ru
```

## Demo Email Templates / Демо шаблоны email

Example email subjects and bodies are available in both languages:

Примеры тем и текстов email доступны на обоих языках:

### Welcome Email / Приветственное письмо

**English**:
- Subject: "Welcome!"
- Body: "Welcome to our service!"

**Russian**:
- Subject: "Добро пожаловать!"
- Body: "Добро пожаловать в наш сервис!"

### Test Email / Тестовый email

**English**:
- Subject: "Hello"
- Body: "<h1>Hello World</h1>"

**Russian**:
- Subject: "Привет"
- Body: "<h1>Привет, мир!</h1>"

## Disabling Demo Data / Отключение демо данных

To disable automatic demo user creation, comment out the following line in `cmd/server/main.go`:

Чтобы отключить автоматическое создание демо пользователей, закомментируйте следующую строку в `cmd/server/main.go`:

```go
// database.SeedDemoUsers()
```

## Translation Keys / Ключи переводов

All translation keys are defined in:

Все ключи переводов определены в:

- `locales/en.json` - English translations
- `locales/ru.json` - Russian translations

### API Translation Keys / Ключи переводов API

- `api.user.*` - User management messages
- `api.email.*` - Email service messages
- `api.queue.*` - Queue service messages
- `demo.*` - Demo data translations

For more information about i18n, see [I18N.md](I18N.md) and [I18N.ru.md](I18N.ru.md).

Для получения дополнительной информации о i18n см. [I18N.md](I18N.md) и [I18N.ru.md](I18N.ru.md).
