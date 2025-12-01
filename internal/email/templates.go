package email

import (
	"bytes"
	"html/template"
)

// WelcomeEmail отправляет приветственное письмо
func WelcomeEmail(to, username string) error {
	subject := "Welcome to DMMVC!"
	body := renderWelcomeTemplate(username)
	return Send(to, subject, body)
}

// PasswordResetEmail отправляет письмо для сброса пароля
func PasswordResetEmail(to, resetLink string) error {
	subject := "Password Reset Request"
	body := renderPasswordResetTemplate(resetLink)
	return Send(to, subject, body)
}

// NotificationEmail отправляет уведомление
func NotificationEmail(to, title, message string) error {
	subject := title
	body := renderNotificationTemplate(title, message)
	return Send(to, subject, body)
}

// renderWelcomeTemplate рендерит шаблон приветственного письма
func renderWelcomeTemplate(username string) string {
	tmpl := `
<!DOCTYPE html>
<html>
<head>
    <meta charset="UTF-8">
    <style>
        body { font-family: Arial, sans-serif; line-height: 1.6; color: #333; }
        .container { max-width: 600px; margin: 0 auto; padding: 20px; }
        .header { background: #4CAF50; color: white; padding: 20px; text-align: center; }
        .content { padding: 20px; background: #f9f9f9; }
        .footer { padding: 20px; text-align: center; font-size: 12px; color: #666; }
    </style>
</head>
<body>
    <div class="container">
        <div class="header">
            <h1>Welcome to DMMVC!</h1>
        </div>
        <div class="content">
            <h2>Hello, {{.Username}}!</h2>
            <p>Thank you for joining DMMVC. We're excited to have you on board!</p>
            <p>You can now start building amazing web applications with our lightweight MVC framework.</p>
            <p>If you have any questions, feel free to reach out to our support team.</p>
        </div>
        <div class="footer">
            <p>&copy; 2024 DMMVC. All rights reserved.</p>
        </div>
    </div>
</body>
</html>
`
	t := template.Must(template.New("welcome").Parse(tmpl))
	var buf bytes.Buffer
	t.Execute(&buf, map[string]string{"Username": username})
	return buf.String()
}

// renderPasswordResetTemplate рендерит шаблон сброса пароля
func renderPasswordResetTemplate(resetLink string) string {
	tmpl := `
<!DOCTYPE html>
<html>
<head>
    <meta charset="UTF-8">
    <style>
        body { font-family: Arial, sans-serif; line-height: 1.6; color: #333; }
        .container { max-width: 600px; margin: 0 auto; padding: 20px; }
        .header { background: #2196F3; color: white; padding: 20px; text-align: center; }
        .content { padding: 20px; background: #f9f9f9; }
        .button { display: inline-block; padding: 12px 24px; background: #2196F3; color: white; text-decoration: none; border-radius: 4px; margin: 20px 0; }
        .footer { padding: 20px; text-align: center; font-size: 12px; color: #666; }
    </style>
</head>
<body>
    <div class="container">
        <div class="header">
            <h1>Password Reset Request</h1>
        </div>
        <div class="content">
            <p>You requested to reset your password.</p>
            <p>Click the button below to reset your password:</p>
            <a href="{{.ResetLink}}" class="button">Reset Password</a>
            <p>If you didn't request this, please ignore this email.</p>
            <p>This link will expire in 24 hours.</p>
        </div>
        <div class="footer">
            <p>&copy; 2024 DMMVC. All rights reserved.</p>
        </div>
    </div>
</body>
</html>
`
	t := template.Must(template.New("reset").Parse(tmpl))
	var buf bytes.Buffer
	t.Execute(&buf, map[string]string{"ResetLink": resetLink})
	return buf.String()
}

// renderNotificationTemplate рендерит шаблон уведомления
func renderNotificationTemplate(title, message string) string {
	tmpl := `
<!DOCTYPE html>
<html>
<head>
    <meta charset="UTF-8">
    <style>
        body { font-family: Arial, sans-serif; line-height: 1.6; color: #333; }
        .container { max-width: 600px; margin: 0 auto; padding: 20px; }
        .header { background: #FF9800; color: white; padding: 20px; text-align: center; }
        .content { padding: 20px; background: #f9f9f9; }
        .footer { padding: 20px; text-align: center; font-size: 12px; color: #666; }
    </style>
</head>
<body>
    <div class="container">
        <div class="header">
            <h1>{{.Title}}</h1>
        </div>
        <div class="content">
            <p>{{.Message}}</p>
        </div>
        <div class="footer">
            <p>&copy; 2024 DMMVC. All rights reserved.</p>
        </div>
    </div>
</body>
</html>
`
	t := template.Must(template.New("notification").Parse(tmpl))
	var buf bytes.Buffer
	t.Execute(&buf, map[string]interface{}{
		"Title":   title,
		"Message": message,
	})
	return buf.String()
}
