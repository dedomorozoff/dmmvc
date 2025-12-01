package logger

import (
	"os"

	"github.com/sirupsen/logrus"
)

var Log *logrus.Logger

// Init инициализирует логгер
func Init() {
	Log = logrus.New()

	// Настройка уровня логирования
	level := os.Getenv("LOG_LEVEL")
	switch level {
	case "debug":
		Log.SetLevel(logrus.DebugLevel)
	case "warn":
		Log.SetLevel(logrus.WarnLevel)
	case "error":
		Log.SetLevel(logrus.ErrorLevel)
	default:
		Log.SetLevel(logrus.InfoLevel)
	}

	// Настройка формата
	Log.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})

	// Настройка вывода в файл
	logFile := os.Getenv("LOG_FILE")
	if logFile != "" {
		file, err := os.OpenFile(logFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err == nil {
			Log.SetOutput(file)
		} else {
			Log.Error("Failed to log to file, using default stderr")
		}
	}
}
