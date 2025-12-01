**English** | [Русский](README.ru.md) | [Navigation](NAVIGATION.md)

# DMMVC - Lightweight MVC Web Framework

**DMMVC** is a minimalist MVC web framework in Go, ready for building any web application.

> **Documentation**: [docs/](docs/) | **Navigation**: [NAVIGATION.md](NAVIGATION.md)

## Features

### Architecture
- **MVC Pattern** - Model-View-Controller
- **Modular Structure** - Extensible architecture
- **Middleware System** - Request processing
- **Routing** - Route system

### Security
- **Authentication** - Authorization system
- **Sessions** - Secure session management
- **Password Hashing** - bcrypt for secure storage
- **CSRF Protection** - Ready for integration

### Database
- **GORM ORM** - ORM for database interaction
- **Migrations** - Automatic database structure creation
- **SQLite Support** - For quick start
- **MySQL Support** - For production environment
- **PostgreSQL Support** - Powerful relational database

### Templates
- **Go Templates** - Template engine
- **Layouts** - Layout system for reusability
- **Partials** - Components for modularity
- **Static Files** - Automatic CSS/JS serving

### Logging
- **Logrus** - Structured logging
- **Request logging** - Automatic request logging
- **Error tracking** - Error tracking
- **Panic recovery** - Panic handling

## Tech Stack

- **Backend**: Go 1.20+
- **Web Framework**: Gin
- **ORM**: GORM
- **Database**: SQLite / MySQL / PostgreSQL
- **Logger**: Logrus
- **Sessions**: gorilla/sessions

## Quick Start

### Installation

```bash
# Clone the repository
git clone https://github.com/dedomorozoff/dmmvc
cd dmmvc

# Install dependencies
go mod tidy

# Run the server
go run cmd/server/main.go
```

### First Login

Open browser: **http://localhost:8080**

Default credentials:
- **Username**: `admin`
- **Password**: `admin`

**Important**: Change the password after the first login!

## Project Structure

```
dmmvc/
├── cmd/
│   ├── cli/                 # CLI tool
│   └── server/              # Application entry point
├── internal/
│   ├── controllers/         # HTTP controllers
│   ├── database/            # Database connection
│   ├── logger/              # Logging
│   ├── middleware/          # Middleware
│   ├── models/              # Data models
│   └── routes/              # Routes
├── static/
│   ├── css/                 # Styles
│   └── js/                  # JavaScript
├── templates/
│   ├── layouts/             # Layouts
│   ├── partials/            # Components
│   └── pages/               # Pages
├── docs/                    # Documentation
├── docker/                  # Docker configuration
├── scripts/                 # Utilities
├── .env.example             # Configuration example
├── Dockerfile               # Docker image
├── Makefile                 # Build commands
└── README.md                # Documentation
```

**Full documentation**: [docs/](docs/)

## Configuration

Create a `.env` file in the project root:

```env
# Server Settings
PORT=8080
GIN_MODE=debug

# Database Settings
DB_TYPE=sqlite
DB_DSN=dmmvc.db

# For MySQL:
# DB_TYPE=mysql
# DB_DSN=user:password@tcp(localhost:3306)/dmmvc?charset=utf8mb4&parseTime=True&loc=Local

# For PostgreSQL:
# DB_TYPE=postgres
# DB_DSN=host=localhost user=postgres password=postgres dbname=dmmvc port=5432 sslmode=disable

# Security Settings
SESSION_SECRET=your-super-secret-key-change-this-in-production

# Logging Settings
LOG_LEVEL=info
LOG_FILE=dmmvc.log

# Development Settings
DEBUG=true
```

## Creating a New Controller

```go
package controllers

import (
    "net/http"
    "github.com/gin-gonic/gin"
)

func MyController(c *gin.Context) {
    c.HTML(http.StatusOK, "pages/mypage.html", gin.H{
        "title": "My Page",
        "data": "Some data",
    })
}
```

## Creating a New Model

```go
package models

import "gorm.io/gorm"

type MyModel struct {
    gorm.Model
    Name        string `json:"name"`
    Description string `json:"description"`
}
```

## Adding a Route

```go
// In internal/routes/routes.go
authorized.GET("/mypage", controllers.MyController)
```

## Creating a Template

```html
{{define "pages/mypage.html"}}
{{template "layouts/base.html" .}}
{{define "content"}}
<div class="container">
    <h1>{{.title}}</h1>
    <p>{{.data}}</p>
</div>
{{end}}
{{end}}
```

## Middleware

### Creating Custom Middleware

```go
func MyMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        // Code before request processing
        c.Next()
        // Code after request processing
    }
}
```

### Using Middleware

```go
// Global
r.Use(MyMiddleware())

// For route group
authorized := r.Group("/")
authorized.Use(MyMiddleware())
```

## Database Operations

```go
import "dmmvc/internal/database"

// Create record
user := models.User{Username: "john", Email: "john@example.com"}
database.DB.Create(&user)

// Find
var user models.User
database.DB.First(&user, 1) // By ID
database.DB.Where("username = ?", "john").First(&user)

// Update
database.DB.Model(&user).Update("Email", "newemail@example.com")

// Delete
database.DB.Delete(&user)
```

## Testing

```bash
# Run tests
go test ./...

# With coverage
go test -cover ./...
```

## Build for Production

```bash
# Build binary
go build -o dmmvc cmd/server/main.go

# Run
./dmmvc
```

## Docker

```dockerfile
FROM golang:1.20-alpine
WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o dmmvc cmd/server/main.go
EXPOSE 8080
CMD ["./dmmvc"]
```

## License

MIT License

## Advanced Features

### CLI Tool

DMMVC includes a powerful CLI tool for code generation:

```bash
# Build CLI
make build

# Create CRUD for products
dmmvc make:crud Product

# Create controller
dmmvc make:controller About

# Create model
dmmvc make:model Category --migration

# List all resources
dmmvc list
```

**Full documentation**: [docs/CLI.md](docs/CLI.md)

### Database Support

DMMVC supports three types of databases:

- **SQLite** - For development and small projects
- **MySQL** - Popular relational database
- **PostgreSQL** - Powerful open-source database

**Documentation**: [docs/POSTGRESQL.md](docs/POSTGRESQL.md)

### WebSocket Support

Real-time bidirectional communication:

```go
// WebSocket endpoint available at /ws
// Demo page at /websocket
```

**Documentation**: [docs/WEBSOCKET.md](docs/WEBSOCKET.md)

### API Documentation (Swagger)

Automatic API documentation generation:

```bash
# Generate Swagger docs
make swagger

# Access Swagger UI
# http://localhost:8080/swagger/index.html
```

**Documentation**: [docs/SWAGGER.md](docs/SWAGGER.md)

### Redis Caching

Improve performance with Redis caching:

```go
// Manual caching
cache.Set("key", "value", 5*time.Minute)
value, _ := cache.Get("key")

// Cache middleware
r.GET("/api/data", middleware.CacheMiddleware(5*time.Minute), handler)
```

**Documentation**: [docs/CACHE.md](docs/CACHE.md)

### Task Queue

Process background jobs asynchronously:

```go
// Create and enqueue task
task, _ := queue.NewEmailDeliveryTask("user@example.com", "Subject", "Body")
queue.EnqueueTask(task)

// Delayed task
queue.EnqueueTaskIn(task, 5*time.Minute)
```

**Documentation**: [docs/QUEUE.md](docs/QUEUE.md)

## Roadmap

- [x] CLI tool for code generation
- [x] PostgreSQL support
- [x] WebSocket support
- [x] API documentation (Swagger)
- [x] Caching (Redis)
- [x] Task queues
- [ ] Email sending
- [ ] File upload helper
- [ ] Localization (i18n)

## Documentation

- [CLI Tool](docs/CLI.md) - Code generation commands
- [PostgreSQL](docs/POSTGRESQL.md) - PostgreSQL setup guide
- [WebSocket](docs/WEBSOCKET.md) - Real-time communication
- [Swagger API](docs/SWAGGER.md) - API documentation
- [Redis Cache](docs/CACHE.md) - Caching with Redis
- [Task Queue](docs/QUEUE.md) - Background job processing
- [Examples](docs/EXAMPLES.md) - Usage examples
- [Deployment](docs/DEPLOYMENT.md) - Production deployment
