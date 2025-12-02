# DMMVC Cheat Sheet

Quick reference for common DMMVC tasks.

## Installation

```bash
# Linux/macOS
git clone https://github.com/dedomorozoff/dmmvc
cd dmmvc
chmod +x scripts/install.sh
./scripts/install.sh

# Windows
git clone https://github.com/dedomorozoff/dmmvc
cd dmmvc
scripts\install.bat

# Via go install
go install github.com/dedomorozoff/dmmvc/cmd/cli@latest
```

## Quick Start

```bash
# Linux/macOS
./scripts/quickstart.sh

# Windows
scripts\quickstart.bat

# Manual
go mod tidy
go run cmd/server/main.go
```

## Create New Project

```bash
# Linux/macOS
./scripts/create-project.sh my-app

# Windows
scripts\create-project.bat my-app

cd my-app
go mod tidy
go run cmd/server/main.go
```

## CLI Commands

```bash
# Show help
dmmvc --help

# Create CRUD
dmmvc make:crud Product
dmmvc make:crud User --fields="name:string,email:string,age:int"

# Create controller
dmmvc make:controller About
dmmvc make:controller api/Users

# Create model
dmmvc make:model Category
dmmvc make:model Post --migration

# Create view
dmmvc make:view products/index
dmmvc make:view users/profile

# List resources
dmmvc list
dmmvc list controllers
dmmvc list models
```

## Development

```bash
# Run server
go run cmd/server/main.go

# Build
go build -o server cmd/server/main.go

# Run tests
go test ./...
go test -v ./...
go test -cover ./...

# Format code
go fmt ./...

# Vet code
go vet ./...

# Update dependencies
go get -u ./...
go mod tidy
```

## Using dev.sh (Linux/macOS)

```bash
# Install dependencies
./scripts/dev.sh install

# Build CLI and server
./scripts/dev.sh build

# Run server
./scripts/dev.sh run

# Run tests
./scripts/dev.sh test

# Format code
./scripts/dev.sh fmt

# Lint code
./scripts/dev.sh lint

# Generate Swagger
./scripts/dev.sh swagger

# Clean
./scripts/dev.sh clean
```

## Database

```bash
# SQLite (default)
DB_TYPE=sqlite
DB_DSN=dmmvc.db

# MySQL
DB_TYPE=mysql
DB_DSN=user:password@tcp(localhost:3306)/dbname?charset=utf8mb4&parseTime=True

# PostgreSQL
DB_TYPE=postgres
DB_DSN=host=localhost user=postgres password=postgres dbname=dmmvc port=5432 sslmode=disable
```

## GORM Operations

```go
// Create
user := models.User{Username: "john", Email: "john@example.com"}
database.DB.Create(&user)

// Find by ID
var user models.User
database.DB.First(&user, 1)

// Find by condition
database.DB.Where("username = ?", "john").First(&user)

// Find all
var users []models.User
database.DB.Find(&users)

// Update
database.DB.Model(&user).Update("Email", "new@example.com")
database.DB.Model(&user).Updates(models.User{Email: "new@example.com", Age: 30})

// Delete
database.DB.Delete(&user)
database.DB.Where("age < ?", 18).Delete(&models.User{})
```

## Routes

```go
// Public routes
r.GET("/", controllers.Index)
r.POST("/api/data", controllers.CreateData)

// Authenticated routes
authorized := r.Group("/")
authorized.Use(middleware.AuthRequired())
{
    authorized.GET("/dashboard", controllers.Dashboard)
    authorized.POST("/profile", controllers.UpdateProfile)
}

// API routes
api := r.Group("/api")
{
    api.GET("/users", controllers.GetUsers)
    api.POST("/users", controllers.CreateUser)
    api.PUT("/users/:id", controllers.UpdateUser)
    api.DELETE("/users/:id", controllers.DeleteUser)
}
```

## Controllers

```go
package controllers

import (
    "net/http"
    "github.com/gin-gonic/gin"
)

// HTML response
func Index(c *gin.Context) {
    c.HTML(http.StatusOK, "pages/index.html", gin.H{
        "title": "Home",
        "data": "Some data",
    })
}

// JSON response
func GetUsers(c *gin.Context) {
    users := []User{} // fetch from DB
    c.JSON(http.StatusOK, gin.H{
        "users": users,
    })
}

// Get URL parameter
func GetUser(c *gin.Context) {
    id := c.Param("id")
    // ...
}

// Get query parameter
func Search(c *gin.Context) {
    query := c.Query("q")
    // ...
}

// Bind JSON
func CreateUser(c *gin.Context) {
    var user User
    if err := c.ShouldBindJSON(&user); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    // ...
}
```

## Middleware

```go
// Custom middleware
func MyMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        // Before request
        c.Next()
        // After request
    }
}

// Use globally
r.Use(MyMiddleware())

// Use for group
authorized := r.Group("/")
authorized.Use(middleware.AuthRequired())
```

## Templates

```html
{{define "pages/mypage.html"}}
{{template "layouts/base.html" .}}
{{define "content"}}
<div class="container">
    <h1>{{.title}}</h1>
    
    {{/* Variables */}}
    <p>{{.message}}</p>
    
    {{/* Conditionals */}}
    {{if .user}}
        <p>Welcome, {{.user.Name}}</p>
    {{else}}
        <p>Please login</p>
    {{end}}
    
    {{/* Loops */}}
    {{range .items}}
        <div>{{.Name}}</div>
    {{end}}
    
    {{/* Include partial */}}
    {{template "partials/header.html" .}}
</div>
{{end}}
{{end}}
```

## Sessions

```go
// Get session
session := sessions.Default(c)

// Set value
session.Set("user_id", user.ID)
session.Save()

// Get value
userID := session.Get("user_id")

// Delete value
session.Delete("user_id")
session.Save()

// Clear session
session.Clear()
session.Save()
```

## Docker

```bash
# Build image
docker build -t dmmvc .
./scripts/docker-build.sh

# Run container
docker run -p 8080:8080 dmmvc

# Run with .env
docker run -p 8080:8080 --env-file .env dmmvc

# Run in background
docker run -d -p 8080:8080 --name dmmvc-app dmmvc

# View logs
docker logs dmmvc-app

# Stop container
docker stop dmmvc-app

# Docker Compose
docker-compose up
docker-compose up -d
docker-compose down
```

## Testing

```bash
# Run all tests
go test ./...

# Verbose
go test -v ./...

# With coverage
go test -cover ./...
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out

# Specific package
go test ./internal/controllers

# Using script
./scripts/test.sh
./scripts/test.sh --coverage
./scripts/test.sh --verbose
./scripts/test.sh --bench
```

## Swagger

```bash
# Generate docs
swag init -g cmd/server/main.go -o docs/swagger
make swagger
./scripts/dev.sh swagger

# Access UI
# http://localhost:8080/swagger/index.html
```

## Environment Variables

```env
# Server
PORT=8080
GIN_MODE=debug|release

# Database
DB_TYPE=sqlite|mysql|postgres
DB_DSN=connection_string

# Security
SESSION_SECRET=your-secret-key

# Logging
LOG_LEVEL=debug|info|warn|error
LOG_FILE=app.log

# Development
DEBUG=true|false
```

## Common Issues

```bash
# Command not found: dmmvc
export PATH=$PATH:$(go env GOPATH)/bin  # Linux/macOS
set PATH=%PATH%;%GOPATH%\bin            # Windows

# Permission denied (Linux/macOS)
chmod +x scripts/*.sh

# Module errors
go clean -modcache
go mod download
go mod tidy

# Port already in use
# Change PORT in .env or kill process using port
lsof -ti:8080 | xargs kill -9  # Linux/macOS
netstat -ano | findstr :8080   # Windows
```

## Useful Links

- [Documentation](docs/)
- [Installation Guide](docs/INSTALLATION.md)
- [CLI Reference](docs/CLI.md)
- [Examples](docs/EXAMPLES.md)
- [Scripts](scripts/README.md)

## Quick Reference Card

```
Installation:  ./scripts/install.sh | scripts\install.bat
New Project:   ./scripts/create-project.sh my-app
Quick Start:   ./scripts/quickstart.sh
Run Server:    go run cmd/server/main.go
Build CLI:     go build -o dmmvc cmd/cli/main.go
Run Tests:     go test ./...
Generate CRUD: dmmvc make:crud Product
Format Code:   go fmt ./...
```
