**English** | [Русский](EMAIL.ru.md)

# Email Sending

DMMVC includes built-in support for sending emails via SMTP with HTML templates.

## Features

- **SMTP support** - Send emails via any SMTP server
- **HTML templates** - Pre-built email templates
- **Async sending** - Send emails in background via task queue
- **TLS/SSL support** - Secure email transmission
- **Template system** - Easy to customize email templates

## Quick Start

### 1. Configure SMTP

Edit `.env` file:

```env
# Gmail example
SMTP_HOST=smtp.gmail.com
SMTP_PORT=587
SMTP_USERNAME=your-email@gmail.com
SMTP_PASSWORD=your-app-password
SMTP_FROM=noreply@example.com
SMTP_USE_TLS=true
```

### 2. Start Server

```bash
make run
```

## SMTP Providers

### Gmail

```env
SMTP_HOST=smtp.gmail.com
SMTP_PORT=587
SMTP_USERNAME=your-email@gmail.com
SMTP_PASSWORD=your-app-password
```

**Note**: Use [App Password](https://support.google.com/accounts/answer/185833) instead of your regular password.

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

### Custom SMTP

```env
SMTP_HOST=mail.example.com
SMTP_PORT=587
SMTP_USERNAME=user@example.com
SMTP_PASSWORD=your-password
SMTP_FROM=noreply@example.com
SMTP_USE_TLS=true
```

## Usage

### Send Simple Email

```go
import "dmmvc/internal/email"

err := email.Send(
    "user@example.com",
    "Hello",
    "<h1>Hello World!</h1><p>This is a test email.</p>",
)
```

### Send to Multiple Recipients

```go
recipients := []string{
    "user1@example.com",
    "user2@example.com",
    "user3@example.com",
}

err := email.SendMultiple(recipients, "Subject", "Body")
```

### Send Async (Background)

```go
import (
    "dmmvc/internal/queue"
    "github.com/hibiken/asynq"
)

// Create task
task, _ := queue.NewEmailDeliveryTask(
    "user@example.com",
    "Subject",
    "Body",
)

// Enqueue for background processing
queue.EnqueueTask(task, asynq.Queue("default"))
```

## Built-in Templates

### Welcome Email

```go
import "dmmvc/internal/email"

err := email.WelcomeEmail("user@example.com", "John Doe")
```

### Password Reset Email

```go
resetLink := "https://example.com/reset?token=abc123"
err := email.PasswordResetEmail("user@example.com", resetLink)
```

### Notification Email

```go
err := email.NotificationEmail(
    "user@example.com",
    "New Message",
    "You have a new message from support.",
)
```

## Custom Templates

Create your own email templates:

```go
// internal/email/templates.go

func CustomEmail(to, name string) error {
    subject := "Custom Email"
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
        <h1>Hello, {{.Name}}!</h1>
        <p>This is a custom email template.</p>
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

## API Endpoints

DMMVC includes example email endpoints:

- `POST /api/email/send` - Send email directly (sync)
- `POST /api/email/send/async` - Send email via queue (async)
- `POST /api/email/welcome` - Send welcome email
- `POST /api/email/password-reset` - Send password reset email
- `GET /api/email/status` - Check email service status

### Example Request

```bash
curl -X POST http://localhost:8080/api/email/send \
  -H "Content-Type: application/json" \
  -d '{
    "to": "user@example.com",
    "subject": "Hello",
    "body": "<h1>Hello World</h1>"
  }'
```

## Controller Example

```go
func RegisterUser(c *gin.Context) {
    // Create user
    user := models.User{
        Username: "john",
        Email:    "john@example.com",
    }
    database.DB.Create(&user)
    
    // Send welcome email asynchronously
    task, _ := queue.NewEmailDeliveryTask(
        user.Email,
        "Welcome!",
        "Thank you for registering!",
    )
    queue.EnqueueTask(task)
    
    c.JSON(200, gin.H{"message": "User registered"})
}
```

## Email with Attachments

For attachments, use a library like `gomail`:

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

## Testing

### Test Email Configuration

```go
func TestEmailConfig(c *gin.Context) {
    if !email.IsEnabled() {
        c.JSON(500, gin.H{"error": "Email not configured"})
        return
    }
    
    err := email.Send(
        "test@example.com",
        "Test Email",
        "<h1>Test</h1><p>If you receive this, email is working!</p>",
    )
    
    if err != nil {
        c.JSON(500, gin.H{"error": err.Error()})
        return
    }
    
    c.JSON(200, gin.H{"message": "Test email sent"})
}
```

## Best Practices

1. **Use async sending** - Don't block HTTP requests
2. **Validate email addresses** - Check format before sending
3. **Handle failures gracefully** - Log errors, retry if needed
4. **Use templates** - Consistent branding and styling
5. **Test thoroughly** - Test with different providers
6. **Monitor delivery** - Track bounces and failures
7. **Respect rate limits** - Don't spam
8. **Use proper from address** - Improve deliverability

## Error Handling

```go
err := email.Send(to, subject, body)
if err != nil {
    logrus.Errorf("Failed to send email: %v", err)
    
    // Retry logic
    task, _ := queue.NewEmailDeliveryTask(to, subject, body)
    queue.EnqueueTaskIn(task, 5*time.Minute)
}
```

## Production Configuration

```env
# Use dedicated email service
SMTP_HOST=smtp.sendgrid.net
SMTP_PORT=587
SMTP_USERNAME=apikey
SMTP_PASSWORD=your-api-key
SMTP_FROM=noreply@yourdomain.com
SMTP_USE_TLS=true

# Enable worker for async sending
QUEUE_WORKER_ENABLED=true
```

## Security

1. **Use App Passwords** - Don't use your main password
2. **Enable TLS** - Always use encrypted connection
3. **Store credentials securely** - Use environment variables
4. **Validate input** - Prevent email injection
5. **Rate limiting** - Prevent abuse

## Troubleshooting

### Email not sending

1. Check SMTP configuration in `.env`
2. Verify credentials are correct
3. Check firewall/network settings
4. Review logs for errors

### Gmail "Less secure app" error

Use [App Password](https://support.google.com/accounts/answer/185833) instead of regular password.

### Connection timeout

1. Check SMTP host and port
2. Verify network connectivity
3. Try different port (25, 465, 587)

### Authentication failed

1. Verify username and password
2. Check if 2FA is enabled (use app password)
3. Ensure account is not locked

## Resources

- [Gmail SMTP Settings](https://support.google.com/mail/answer/7126229)
- [SendGrid Documentation](https://docs.sendgrid.com/)
- [Mailgun Documentation](https://documentation.mailgun.com/)
- [Amazon SES Documentation](https://docs.aws.amazon.com/ses/)

## Next Steps

- Set up email templates
- Configure email queue
- Add email tracking
- Implement bounce handling
- Set up email analytics
