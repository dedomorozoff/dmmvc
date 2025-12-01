package email

import (
	"crypto/tls"
	"fmt"
	"net/smtp"
	"os"
	"strconv"

	"github.com/sirupsen/logrus"
)

// Config конфигурация SMTP
type Config struct {
	Host     string
	Port     int
	Username string
	Password string
	From     string
	UseTLS   bool
}

var (
	config *Config
)

// Init инициализирует email конфигурацию
func Init() error {
	port, _ := strconv.Atoi(os.Getenv("SMTP_PORT"))
	if port == 0 {
		port = 587 // default SMTP port
	}

	useTLS := os.Getenv("SMTP_USE_TLS") != "false"

	config = &Config{
		Host:     os.Getenv("SMTP_HOST"),
		Port:     port,
		Username: os.Getenv("SMTP_USERNAME"),
		Password: os.Getenv("SMTP_PASSWORD"),
		From:     os.Getenv("SMTP_FROM"),
		UseTLS:   useTLS,
	}

	if config.Host == "" {
		logrus.Warn("SMTP not configured, email sending disabled")
		return fmt.Errorf("SMTP_HOST not set")
	}

	logrus.Info("Email service initialized")
	return nil
}

// IsEnabled проверяет, настроен ли email
func IsEnabled() bool {
	return config != nil && config.Host != ""
}

// Send отправляет email
func Send(to, subject, body string) error {
	if !IsEnabled() {
		logrus.Warn("Email service is disabled")
		return fmt.Errorf("email service not configured")
	}

	return SendWithConfig(config, to, subject, body)
}

// SendWithConfig отправляет email с указанной конфигурацией
func SendWithConfig(cfg *Config, to, subject, body string) error {
	from := cfg.From
	if from == "" {
		from = cfg.Username
	}

	// Формирование сообщения
	message := []byte(
		"From: " + from + "\r\n" +
			"To: " + to + "\r\n" +
			"Subject: " + subject + "\r\n" +
			"MIME-Version: 1.0\r\n" +
			"Content-Type: text/html; charset=UTF-8\r\n" +
			"\r\n" +
			body + "\r\n",
	)

	// Аутентификация
	auth := smtp.PlainAuth("", cfg.Username, cfg.Password, cfg.Host)

	// Адрес сервера
	addr := fmt.Sprintf("%s:%d", cfg.Host, cfg.Port)

	// Отправка с TLS
	if cfg.UseTLS {
		return sendWithTLS(addr, auth, from, []string{to}, message)
	}

	// Отправка без TLS
	return smtp.SendMail(addr, auth, from, []string{to}, message)
}

// sendWithTLS отправляет email с TLS
func sendWithTLS(addr string, auth smtp.Auth, from string, to []string, msg []byte) error {
	// Подключение к серверу
	client, err := smtp.Dial(addr)
	if err != nil {
		return err
	}
	defer client.Close()

	// STARTTLS
	tlsConfig := &tls.Config{
		ServerName: addr[:len(addr)-4], // Remove :port
	}
	if err = client.StartTLS(tlsConfig); err != nil {
		return err
	}

	// Аутентификация
	if err = client.Auth(auth); err != nil {
		return err
	}

	// Отправитель
	if err = client.Mail(from); err != nil {
		return err
	}

	// Получатели
	for _, recipient := range to {
		if err = client.Rcpt(recipient); err != nil {
			return err
		}
	}

	// Данные
	w, err := client.Data()
	if err != nil {
		return err
	}

	_, err = w.Write(msg)
	if err != nil {
		return err
	}

	err = w.Close()
	if err != nil {
		return err
	}

	return client.Quit()
}

// SendMultiple отправляет email нескольким получателям
func SendMultiple(to []string, subject, body string) error {
	if !IsEnabled() {
		return fmt.Errorf("email service not configured")
	}

	for _, recipient := range to {
		if err := Send(recipient, subject, body); err != nil {
			logrus.Errorf("Failed to send email to %s: %v", recipient, err)
			return err
		}
	}

	return nil
}
