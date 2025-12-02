#!/bin/bash
# DMMVC Project Creation Script

set -e

if [ -z "$1" ]; then
    echo "Usage: ./create-project.sh [project-name]"
    echo "Example: ./create-project.sh my-app"
    exit 1
fi

PROJECT_NAME="$1"
PROJECT_DIR="$(pwd)/$PROJECT_NAME"

echo "========================================"
echo "Creating new DMMVC project: $PROJECT_NAME"
echo "========================================"
echo ""

# Check if directory exists
if [ -d "$PROJECT_DIR" ]; then
    echo "[ERROR] Directory '$PROJECT_NAME' already exists"
    exit 1
fi

echo "[1/5] Creating project directory..."
mkdir -p "$PROJECT_DIR"
cd "$PROJECT_DIR"
echo ""

echo "[2/5] Initializing Go module..."
go mod init "$PROJECT_NAME"
echo ""

echo "[3/5] Installing dependencies..."
go get github.com/gin-gonic/gin@latest
go get github.com/joho/godotenv@latest
go get gorm.io/gorm@latest
go get github.com/glebarez/sqlite@latest
go get go.uber.org/zap@latest
echo ""

echo "[4/5] Creating project structure..."
mkdir -p cmd/server
mkdir -p internal/controllers
mkdir -p internal/database
mkdir -p internal/logger
mkdir -p internal/models
mkdir -p static/css
mkdir -p static/js
mkdir -p templates/layouts
mkdir -p templates/pages
mkdir -p templates/partials
echo ""

echo "[5/5] Creating initial files..."

# Create main.go
cat > cmd/server/main.go << EOF
package main

import (
    "${PROJECT_NAME}/internal/controllers"
    "${PROJECT_NAME}/internal/database"
    "${PROJECT_NAME}/internal/logger"
    "github.com/gin-gonic/gin"
    "github.com/joho/godotenv"
    "log"
    "os"
)

func main() {
    // Load .env
    if err := godotenv.Load(); err != nil {
        log.Println("No .env file found")
    }

    // Initialize logger
    logger.Init()

    // Initialize database
    database.Connect()

    // Setup Gin
    r := gin.Default()

    // Load templates
    r.LoadHTMLGlob("templates/**/*")
    r.Static("/static", "./static")

    // Setup routes
    r.GET("/", controllers.HomeHandler)

    // Start server
    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }
    log.Printf("Server starting on port %s", port)
    if err := r.Run(":" + port); err != nil {
        log.Fatal(err)
    }
}
EOF

# Create database package
cat > internal/database/database.go << EOF
package database

import (
    "log"
    "os"
    "github.com/glebarez/sqlite"
    "gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
    dbType := os.Getenv("DB_TYPE")
    dbDSN := os.Getenv("DB_DSN")

    if dbType == "" {
        dbType = "sqlite"
    }
    if dbDSN == "" {
        dbDSN = "${PROJECT_NAME}.db"
    }

    var err error
    // Use pure Go SQLite driver (no CGO required)
    DB, err = gorm.Open(sqlite.Open(dbDSN), &gorm.Config{})
    if err != nil {
        log.Fatal("Failed to connect to database:", err)
    }
    log.Println("Database connected")
}
EOF

# Create logger package
cat > internal/logger/logger.go << 'EOF'
package logger

import (
    "go.uber.org/zap"
    "log"
)

var Log *zap.SugaredLogger

func Init() {
    logger, err := zap.NewDevelopment()
    if err != nil {
        log.Fatal("Failed to initialize logger:", err)
    }
    Log = logger.Sugar()
}
EOF

# Create home controller
cat > internal/controllers/home.go << 'EOF'
package controllers

import (
    "github.com/gin-gonic/gin"
    "net/http"
)

func HomeHandler(c *gin.Context) {
    c.HTML(http.StatusOK, "home.html", gin.H{
        "title": "Welcome",
    })
}
EOF

# Create base layout
cat > templates/layouts/base.html << 'EOF'
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>{{ .title }}</title>
    <link rel="stylesheet" href="/static/css/style.css">
</head>
<body>
    {{ template "content" . }}
</body>
</html>
EOF

# Create home page
cat > templates/pages/home.html << EOF
{{ define "content" }}
<div class="container">
    <h1>Welcome to ${PROJECT_NAME}</h1>
    <p>Your DMMVC application is running!</p>
</div>
{{ end }}
EOF

# Create basic CSS
cat > static/css/style.css << 'EOF'
body {
    font-family: Arial, sans-serif;
    margin: 0;
    padding: 20px;
    background-color: #f5f5f5;
}

.container {
    max-width: 800px;
    margin: 0 auto;
    background: white;
    padding: 40px;
    border-radius: 8px;
    box-shadow: 0 2px 4px rgba(0,0,0,0.1);
}

h1 {
    color: #333;
}
EOF

# Create .env
cat > .env << EOF
PORT=8080
GIN_MODE=debug
DB_TYPE=sqlite
DB_DSN=${PROJECT_NAME}.db
SESSION_SECRET=change-this-secret-key
LOG_LEVEL=info
LOG_FILE=${PROJECT_NAME}.log
DEBUG=true
EOF

# Create .gitignore
cat > .gitignore << 'EOF'
# Binaries
*.exe
*.dll
*.so
*.dylib
dmmvc
server

# Test files
*.test

# Output
*.out

# Database
*.db

# Logs
*.log

# Environment
.env

# IDE
.vscode/
.idea/
*.swp
*.swo
*~

# Uploads
uploads/

# Temp
tmp/

# macOS
.DS_Store

# Linux
*~
EOF

# Create README
cat > README.md << EOF
# $PROJECT_NAME

DMMVC-based web application

## Quick Start

\`\`\`bash
# Install dependencies
go mod tidy

# Run server
go run cmd/server/main.go
\`\`\`

Open http://localhost:8080

## Development

\`\`\`bash
# Run tests
go test ./...

# Build
go build -o server cmd/server/main.go

# Run binary
./server
\`\`\`

## Documentation

- [DMMVC Documentation](https://github.com/dedomorozoff/dmmvc)
EOF

echo ""
echo "========================================"
echo "Project created successfully!"
echo "========================================"
echo ""
echo "Next steps:"
echo "  cd $PROJECT_NAME"
echo "  go mod tidy"
echo "  go run cmd/server/main.go"
echo ""
echo "Happy coding!"
