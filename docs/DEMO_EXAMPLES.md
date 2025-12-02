# Demo Data Examples / Примеры использования демо данных

## Testing with Demo Users / Тестирование с демо пользователями

### Login as Admin / Вход как администратор

```bash
curl -X POST http://localhost:8080/login \
  -H "Content-Type: application/x-www-form-urlencoded" \
  -d "username=admin&password=admin"
```

### Login as English Demo User / Вход как английский демо пользователь

```bash
curl -X POST http://localhost:8080/login \
  -H "Content-Type: application/x-www-form-urlencoded" \
  -d "username=john_doe&password=password123"
```

### Login as Russian Demo User / Вход как русский демо пользователь

```bash
curl -X POST http://localhost:8080/login \
  -H "Content-Type: application/x-www-form-urlencoded" \
  -d "username=ivan_ivanov&password=password123"
```

## API Examples with i18n / Примеры API с i18n

### Get Users List (English) / Получить список пользователей (английский)

```bash
curl -H "Accept-Language: en" \
     -H "Cookie: session=YOUR_SESSION" \
     http://localhost:8080/api/users
```

Response:
```json
{
  "success": true,
  "data": [
    {
      "id": 1,
      "username": "admin",
      "email": "admin@example.com",
      "role": "admin"
    },
    {
      "id": 2,
      "username": "john_doe",
      "email": "john@example.com",
      "role": "user"
    }
  ]
}
```

### Get Users List (Russian) / Получить список пользователей (русский)

```bash
curl -H "Accept-Language: ru" \
     -H "Cookie: session=YOUR_SESSION" \
     http://localhost:8080/api/users
```

### Create User with English Response / Создать пользователя с английским ответом

```bash
curl -X POST http://localhost:8080/api/users \
  -H "Content-Type: application/json" \
  -H "Accept-Language: en" \
  -H "Cookie: session=YOUR_SESSION" \
  -d '{
    "username": "new_user",
    "email": "new@example.com",
    "password": "password123"
  }'
```

Response:
```json
{
  "success": true,
  "message": "User created successfully",
  "data": {
    "id": 7,
    "username": "new_user",
    "email": "new@example.com",
    "role": "user"
  }
}
```

### Create User with Russian Response / Создать пользователя с русским ответом

```bash
curl -X POST http://localhost:8080/api/users \
  -H "Content-Type: application/json" \
  -H "Accept-Language: ru" \
  -H "Cookie: session=YOUR_SESSION" \
  -d '{
    "username": "новый_пользователь",
    "email": "новый@example.ru",
    "password": "password123"
  }'
```

Response:
```json
{
  "success": true,
  "message": "Пользователь успешно создан",
  "data": {
    "id": 8,
    "username": "новый_пользователь",
    "email": "новый@example.ru",
    "role": "user"
  }
}
```

## Email Examples / Примеры email

### Send Welcome Email (English) / Отправить приветственное письмо (английский)

```bash
curl -X POST http://localhost:8080/api/email/welcome \
  -H "Content-Type: application/json" \
  -H "Accept-Language: en" \
  -H "Cookie: session=YOUR_SESSION" \
  -d '{
    "to": "john@example.com",
    "username": "john_doe"
  }'
```

Response:
```json
{
  "success": true,
  "message": "Welcome email sent successfully"
}
```

### Send Welcome Email (Russian) / Отправить приветственное письмо (русский)

```bash
curl -X POST http://localhost:8080/api/email/welcome \
  -H "Content-Type: application/json" \
  -H "Accept-Language: ru" \
  -H "Cookie: session=YOUR_SESSION" \
  -d '{
    "to": "ivan@example.ru",
    "username": "ivan_ivanov"
  }'
```

Response:
```json
{
  "success": true,
  "message": "Приветственное письмо успешно отправлено"
}
```

## Queue Examples / Примеры очередей

### Enqueue Email Task (English) / Добавить задачу email в очередь (английский)

```bash
curl -X POST http://localhost:8080/api/queue/email \
  -H "Content-Type: application/json" \
  -H "Accept-Language: en" \
  -H "Cookie: session=YOUR_SESSION" \
  -d '{
    "to": "john@example.com",
    "subject": "Welcome!",
    "body": "Welcome to our service!"
  }'
```

Response:
```json
{
  "success": true,
  "message": "Email task enqueued successfully"
}
```

### Enqueue Email Task (Russian) / Добавить задачу email в очередь (русский)

```bash
curl -X POST http://localhost:8080/api/queue/email \
  -H "Content-Type: application/json" \
  -H "Accept-Language: ru" \
  -H "Cookie: session=YOUR_SESSION" \
  -d '{
    "to": "ivan@example.ru",
    "subject": "Добро пожаловать!",
    "body": "Добро пожаловать в наш сервис!"
  }'
```

Response:
```json
{
  "success": true,
  "message": "Задача отправки email добавлена в очередь"
}
```

## Error Handling Examples / Примеры обработки ошибок

### User Not Found (English) / Пользователь не найден (английский)

```bash
curl -H "Accept-Language: en" \
     -H "Cookie: session=YOUR_SESSION" \
     http://localhost:8080/api/users/999
```

Response:
```json
{
  "success": false,
  "error": "User not found"
}
```

### User Not Found (Russian) / Пользователь не найден (русский)

```bash
curl -H "Accept-Language: ru" \
     -H "Cookie: session=YOUR_SESSION" \
     http://localhost:8080/api/users/999
```

Response:
```json
{
  "success": false,
  "error": "Пользователь не найден"
}
```

## Testing Language Switching / Тестирование переключения языков

### Test Home Page / Тест главной страницы

```bash
# English
curl http://localhost:8080/?lang=en | grep "Welcome to DMMVC"

# Russian
curl http://localhost:8080/?lang=ru | grep "Добро пожаловать в DMMVC"
```

### Test with Different Browsers / Тест с разными браузерами

1. **Chrome/Edge**: Settings → Languages → Add Russian/English
2. **Firefox**: Preferences → Language → Choose language
3. **Safari**: Preferences → General → Preferred languages

## Database Queries / Запросы к базе данных

### View All Demo Users / Просмотр всех демо пользователей

```sql
SELECT id, username, email, role FROM users;
```

Expected output:
```
id | username        | email                | role
---+-----------------+----------------------+------
1  | admin           | admin@example.com    | admin
2  | john_doe        | john@example.com     | user
3  | jane_smith      | jane@example.com     | user
4  | bob_johnson     | bob@example.com      | user
5  | ivan_ivanov     | ivan@example.ru      | user
6  | maria_petrova   | maria@example.ru     | user
7  | alexey_sidorov  | alexey@example.ru    | user
```

### Filter English Users / Фильтр английских пользователей

```sql
SELECT * FROM users WHERE email LIKE '%@example.com';
```

### Filter Russian Users / Фильтр русских пользователей

```sql
SELECT * FROM users WHERE email LIKE '%@example.ru';
```

## Integration Testing / Интеграционное тестирование

### Test Script Example / Пример тестового скрипта

```bash
#!/bin/bash

echo "Testing English API..."
EN_RESPONSE=$(curl -s -H "Accept-Language: en" http://localhost:8080/api/email/status)
echo $EN_RESPONSE | grep "Email service"

echo "Testing Russian API..."
RU_RESPONSE=$(curl -s -H "Accept-Language: ru" http://localhost:8080/api/email/status)
echo $RU_RESPONSE | grep "Email сервис"

echo "All tests passed!"
```

## Related Documentation / Связанная документация

- [Demo Data](DEMO_DATA.md) - Full demo data documentation
- [Language Switching](LANGUAGE_SWITCHING.md) - How to switch languages
- [i18n Guide](I18N.md) - Internationalization guide
- [API Documentation](SWAGGER.md) - Complete API reference

---

**Note**: Replace `YOUR_SESSION` with actual session cookie value obtained after login.

**Примечание**: Замените `YOUR_SESSION` на фактическое значение cookie сессии, полученное после входа.
