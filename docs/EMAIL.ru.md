[English](EMAIL.md) | **Русский**

# Отправка Email

DMMVC включает встроенную поддержку отправки email через SMTP с HTML шаблонами.

## Возможности

- **SMTP поддержка** - Отправка email через любой SMTP сервер
- **HTML шаблоны** - Готовые шаблоны писем
- **Асинхронная отправка** - Отправка в фоновом режиме через очередь задач
- **TLS/SSL поддержка** - Безопасная передача email
- **Система шаблонов** - Легко настраиваемые шаблоны писем

## Быстрый старт

### 1. Настройка SMTP

Отредактируйте файл `.env`:

```env
# Пример для Gmail
SMTP_HOST=smtp.gmail.com
SMTP_PORT=587
SMTP_USERNAME=your-email@gmail.com
SMTP_PASSWORD=your-app-password
SMTP_FROM=noreply@example.com
SMTP_USE_TLS=true
```

### 2. Запуск сервера

```bash
make run
```

## SMTP провайдеры

### Gmail

```env
SMTP_HOST=smtp.gmail.com
SMTP_PORT=587
SMTP_USERNAME=your-email@gmail.com
SMTP_PASSWORD=your-app-password
```

**Примечание**: Используйте [App Password](https://support.google.com/accounts/answer/185833) вместо обычного пароля.

### SendGrid

```env
SMTP_HOST=smtp.sendgrid.net
SMTP_PORT=587
SMTP_USERNAME=apikey
SMTP_PASSWORD=your-sendgrid-api-key
```

### Mailgun

```env
SMTP_HOST=smtp.mailgun.org
SMTP_PORT=587
SMTP_USERNAME=postmaster@your-domain.mailgun.org
SMTP_PASSWORD=your-mailgun-password
```

### Amazon SES

```env
SMTP_HOST=email-smtp.us-east-1.amazonaws.com
SMTP_PORT=587
SMTP_USERNAME=your-ses-smtp-username
SMTP_PASSWORD=your-ses-smtp-password
```

### Пользовательский SMTP

```env
SMTP_HOST=mail.example.com
SMTP_PORT=587
SMTP_USERNAME=user@example.com
SMTP_PASSWORD=your-password
SMTP_FROM=noreply@example.com
SMTP_USE_TLS=true
```

## Использование

### Отправка простого email

```go
import "dmmvc/internal/email"

err := email.Send(
    "user@example.com",
    "Привет",
    "<h1>Привет, мир!</h1><p>Это тестовое письмо.</p>",
)
```

### Отправка нескольким получателям

```go
recipients := []string{
    "user1@example.com",
    "user2@example.com",
    "user3@example.com",
}

err := email.SendMultiple(recipients, "Тема", "Текст")
```

### Асинхронная отправка (в фоне)

```go
import (
    "dmmvc/internal/queue"
    "github.com/hibiken/asynq"
)

// Создать задачу
task, _ := queue.NewEmailDeliveryTask(
    "user@example.com",
    "Тема",
    "Текст",
)

// Добавить в очередь для фоновой обработки
queue.EnqueueTask(task, asynq.Queue("default"))
```

## Встроенные шаблоны

### Приветственное письмо

```go
import "dmmvc/internal/email"

err := email.WelcomeEmail("user@example.com", "Иван Иванов")
```

### Письмо для сброса пароля

```go
resetLink := "https://example.com/reset?token=abc123"
err := email.PasswordResetEmail("user@example.com", resetLink)
```

### Уведомление

```go
err := email.NotificationEmail(
    "user@example.com",
    "Новое сообщение",
    "У вас новое сообщение от поддержки.",
)
```

## Пользовательские шаблоны

Создайте свои собственные шаблоны email:

```go
// internal/email/templates.go

func CustomEmail(to, name string) error {
    subject := "Пользовательское письмо"
    body := renderCustomTemplate(name)
    return Send(to, subject, body)
}

func renderCustomTemplate(name string) string {
    tmpl := `
<!DOCTYPE html>
<html>
<head>
    <meta charset="UTF-8">
    <style>
        body { font-family: Arial, sans-serif; }
        .container { max-width: 600px; margin: 0 auto; }
    </style>
</head>
<body>
    <div class="container">
        <h1>Привет, {{.Name}}!</h1>
        <p>Это пользовательский шаблон письма.</p>
    </div>
</body>
</html>
`
    t := template.Must(template.New("custom").Parse(tmpl))
    var buf bytes.Buffer
    t.Execute(&buf, map[string]string{"Name": name})
    return buf.String()
}
```

## API эндпоинты

DMMVC включает примеры email эндпоинтов:

- `POST /api/email/send` - Отправить email напрямую (синхронно)
- `POST /api/email/send/async` - Отправить email через очередь (асинхронно)
- `POST /api/email/welcome` - Отправить приветственное письмо
- `POST /api/email/password-reset` - Отправить письмо для сброса пароля
- `GET /api/email/status` - Проверить статус email сервиса

### Пример запроса

```bash
curl -X POST http://localhost:8080/api/email/send \
  -H "Content-Type: application/json" \
  -d '{
    "to": "user@example.com",
    "subject": "Привет",
    "body": "<h1>Привет, мир</h1>"
  }'
```

## Пример контроллера

```go
func RegisterUser(c *gin.Context) {
    // Создать пользователя
    user := models.User{
        Username: "ivan",
        Email:    "ivan@example.com",
    }
    database.DB.Create(&user)
    
    // Отправить приветственное письмо асинхронно
    task, _ := queue.NewEmailDeliveryTask(
        user.Email,
        "Добро пожаловать!",
        "Спасибо за регистрацию!",
    )
    queue.EnqueueTask(task)
    
    c.JSON(200, gin.H{"message": "Пользователь зарегистрирован"})
}
```

## Email с вложениями

Для вложений используйте библиотеку `gomail`:

```bash
go get gopkg.in/gomail.v2
```

```go
import "gopkg.in/gomail.v2"

func SendWithAttachment(to, subject, body, filePath string) error {
    m := gomail.NewMessage()
    m.SetHeader("From", config.From)
    m.SetHeader("To", to)
    m.SetHeader("Subject", subject)
    m.SetBody("text/html", body)
    m.Attach(filePath)
    
    d := gomail.NewDialer(config.Host, config.Port, config.Username, config.Password)
    return d.DialAndSend(m)
}
```

## Тестирование

### Тест конфигурации email

```go
func TestEmailConfig(c *gin.Context) {
    if !email.IsEnabled() {
        c.JSON(500, gin.H{"error": "Email не настроен"})
        return
    }
    
    err := email.Send(
        "test@example.com",
        "Тестовое письмо",
        "<h1>Тест</h1><p>Если вы получили это, email работает!</p>",
    )
    
    if err != nil {
        c.JSON(500, gin.H{"error": err.Error()})
        return
    }
    
    c.JSON(200, gin.H{"message": "Тестовое письмо отправлено"})
}
```

## Лучшие практики

1. **Используйте асинхронную отправку** - Не блокируйте HTTP запросы
2. **Валидируйте email адреса** - Проверяйте формат перед отправкой
3. **Обрабатывайте ошибки корректно** - Логируйте ошибки, повторяйте при необходимости
4. **Используйте шаблоны** - Единообразный брендинг и стиль
5. **Тестируйте тщательно** - Тестируйте с разными провайдерами
6. **Мониторьте доставку** - Отслеживайте отказы и ошибки
7. **Соблюдайте лимиты** - Не спамьте
8. **Используйте правильный from адрес** - Улучшает доставляемость

## Обработка ошибок

```go
err := email.Send(to, subject, body)
if err != nil {
    logrus.Errorf("Не удалось отправить email: %v", err)
    
    // Логика повтора
    task, _ := queue.NewEmailDeliveryTask(to, subject, body)
    queue.EnqueueTaskIn(task, 5*time.Minute)
}
```

## Production конфигурация

```env
# Используйте выделенный email сервис
SMTP_HOST=smtp.sendgrid.net
SMTP_PORT=587
SMTP_USERNAME=apikey
SMTP_PASSWORD=your-api-key
SMTP_FROM=noreply@yourdomain.com
SMTP_USE_TLS=true

# Включите worker для асинхронной отправки
QUEUE_WORKER_ENABLED=true
```

## Безопасность

1. **Используйте App Passwords** - Не используйте основной пароль
2. **Включите TLS** - Всегда используйте зашифрованное соединение
3. **Храните учетные данные безопасно** - Используйте переменные окружения
4. **Валидируйте ввод** - Предотвращайте email injection
5. **Rate limiting** - Предотвращайте злоупотребления

## Решение проблем

### Email не отправляется

1. Проверьте SMTP конфигурацию в `.env`
2. Проверьте правильность учетных данных
3. Проверьте настройки firewall/сети
4. Просмотрите логи на наличие ошибок

### Gmail ошибка "Less secure app"

Используйте [App Password](https://support.google.com/accounts/answer/185833) вместо обычного пароля.

### Connection timeout

1. Проверьте SMTP host и port
2. Проверьте сетевое подключение
3. Попробуйте другой порт (25, 465, 587)

### Authentication failed

1. Проверьте username и password
2. Проверьте, включена ли 2FA (используйте app password)
3. Убедитесь, что аккаунт не заблокирован

## Ресурсы

- [Gmail SMTP настройки](https://support.google.com/mail/answer/7126229)
- [Документация SendGrid](https://docs.sendgrid.com/)
- [Документация Mailgun](https://documentation.mailgun.com/)
- [Документация Amazon SES](https://docs.aws.amazon.com/ses/)

## Следующие шаги

- Настроить email шаблоны
- Настроить очередь email
- Добавить отслеживание email
- Реализовать обработку отказов
- Настроить email аналитику
