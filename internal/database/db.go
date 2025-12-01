package database

import (
	"log"
	"os"

	"github.com/glebarez/sqlite"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

// Connect устанавливает соединение с базой данных
func Connect() {
	var err error
	dbType := os.Getenv("DB_TYPE")
	dsn := os.Getenv("DB_DSN")

	if dbType == "mysql" {
		if dsn == "" {
			dsn = "user:password@tcp(127.0.0.1:3306)/dmmvc?charset=utf8mb4&parseTime=True&loc=Local"
		}
		DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	} else {
		// По умолчанию SQLite
		if dsn == "" {
			dsn = "dmmvc.db"
		}
		DB, err = gorm.Open(sqlite.Open(dsn), &gorm.Config{})
	}

	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	log.Println("Database connected successfully")
}

// Migrate выполняет миграции для переданных моделей
func Migrate(dst ...interface{}) {
	if err := DB.AutoMigrate(dst...); err != nil {
		log.Fatal("Failed to migrate database:", err)
	}
	log.Println("Database migration completed")
}
