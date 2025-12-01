package database

import (
	"fmt"
	"log"
	"os"

	"github.com/glebarez/sqlite"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

// Connect устанавливает соединение с базой данных
func Connect() {
	var err error
	dbType := os.Getenv("DB_TYPE")
	dsn := os.Getenv("DB_DSN")

	switch dbType {
	case "mysql":
		if dsn == "" {
			dsn = "user:password@tcp(127.0.0.1:3306)/dmmvc?charset=utf8mb4&parseTime=True&loc=Local"
		}
		DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		log.Println("Connecting to MySQL database...")

	case "postgres", "postgresql":
		if dsn == "" {
			dsn = "host=localhost user=postgres password=postgres dbname=dmmvc port=5432 sslmode=disable"
		}
		DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		log.Println("Connecting to PostgreSQL database...")

	default:
		// По умолчанию SQLite
		if dsn == "" {
			dsn = "dmmvc.db"
		}
		DB, err = gorm.Open(sqlite.Open(dsn), &gorm.Config{})
		log.Println("Connecting to SQLite database...")
	}

	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	log.Printf("Database connected successfully (type: %s)\n", getDBType(dbType))
}

// getDBType возвращает читаемое имя типа БД
func getDBType(dbType string) string {
	switch dbType {
	case "mysql":
		return "MySQL"
	case "postgres", "postgresql":
		return "PostgreSQL"
	default:
		return "SQLite"
	}
}

// GetDBInfo возвращает информацию о подключенной БД
func GetDBInfo() string {
	dbType := os.Getenv("DB_TYPE")
	switch dbType {
	case "mysql":
		return "MySQL"
	case "postgres", "postgresql":
		return "PostgreSQL"
	default:
		return "SQLite"
	}
}

// TestConnection проверяет соединение с БД
func TestConnection() error {
	sqlDB, err := DB.DB()
	if err != nil {
		return fmt.Errorf("failed to get database instance: %w", err)
	}

	if err := sqlDB.Ping(); err != nil {
		return fmt.Errorf("failed to ping database: %w", err)
	}

	return nil
}

// Migrate выполняет миграции для переданных моделей
func Migrate(dst ...interface{}) {
	if err := DB.AutoMigrate(dst...); err != nil {
		log.Fatal("Failed to migrate database:", err)
	}
	log.Println("Database migration completed")
}
