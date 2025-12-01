package main

import (
	"fmt"
	"log"
	"os"

	"github.com/glebarez/sqlite"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using environment variables")
	}

	dbType := os.Getenv("DB_TYPE")
	dsn := os.Getenv("DB_DSN")

	if dbType == "" {
		dbType = "sqlite"
	}

	fmt.Printf("Testing database connection...\n")
	fmt.Printf("Database Type: %s\n", dbType)
	fmt.Printf("DSN: %s\n\n", maskPassword(dsn))

	var db *gorm.DB
	var err error

	switch dbType {
	case "mysql":
		if dsn == "" {
			dsn = "user:password@tcp(127.0.0.1:3306)/dmmvc?charset=utf8mb4&parseTime=True&loc=Local"
		}
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		fmt.Println("Connecting to MySQL...")

	case "postgres", "postgresql":
		if dsn == "" {
			dsn = "host=localhost user=postgres password=postgres dbname=dmmvc port=5432 sslmode=disable"
		}
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		fmt.Println("Connecting to PostgreSQL...")

	default:
		if dsn == "" {
			dsn = "test.db"
		}
		db, err = gorm.Open(sqlite.Open(dsn), &gorm.Config{})
		fmt.Println("Connecting to SQLite...")
	}

	if err != nil {
		log.Fatalf("❌ Failed to connect to database: %v\n", err)
	}

	// Test connection
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("❌ Failed to get database instance: %v\n", err)
	}

	if err := sqlDB.Ping(); err != nil {
		log.Fatalf("❌ Failed to ping database: %v\n", err)
	}

	// Get database info
	var version string
	switch dbType {
	case "mysql":
		db.Raw("SELECT VERSION()").Scan(&version)
	case "postgres", "postgresql":
		db.Raw("SELECT version()").Scan(&version)
	case "sqlite":
		db.Raw("SELECT sqlite_version()").Scan(&version)
	}

	fmt.Println("\n✅ Database connection successful!")
	fmt.Printf("Database Version: %s\n", version)

	// Get connection stats
	stats := sqlDB.Stats()
	fmt.Printf("\nConnection Stats:\n")
	fmt.Printf("  Open Connections: %d\n", stats.OpenConnections)
	fmt.Printf("  In Use: %d\n", stats.InUse)
	fmt.Printf("  Idle: %d\n", stats.Idle)

	// Close connection
	sqlDB.Close()
	fmt.Println("\n✅ Test completed successfully!")
}

// maskPassword masks password in DSN for display
func maskPassword(dsn string) string {
	if dsn == "" {
		return "(empty)"
	}
	// Simple masking - just show first 20 chars
	if len(dsn) > 50 {
		return dsn[:20] + "..." + dsn[len(dsn)-10:]
	}
	return dsn
}
