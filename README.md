**English** | [Русский](README.ru.md)

# DMMVC - Lightweight MVC Web Framework

**DMMVC** is a minimalist MVC web framework in Go, ready for building any web application.

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
- **Database**: SQLite / MySQL
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
│   └── server/              # Application entry point
│       └── main.go
├── internal/
│   ├── controllers/         # HTTP controllers
│   │   ├── auth_controller.go
│   │   ├── home_controller.go
│   │   └── user_controller.go
│   ├── database/            # Database connection
│   │   ├── db.go
│   │   └── seeder.go
│   ├── logger/              # Logging
│   │   └── logger.go
│   ├── middleware/          # Middleware
│   │   ├── auth.go
│   │   └── logger.go
│   ├── models/              # Data models
│   │   └── user.go
│   └── routes/              # Routes
│       └── routes.go
├── static/
│   ├── css/                 # Styles
│   │   └── style.css
│   └── js/                  # JavaScript
│       └── app.js
├── templates/
│   ├── layouts/             # Layouts
│   │   └── base.html
│   ├── partials/            # Reusable components
│   │   ├── header.html
│   │   ├── footer.html
│   │   └── sidebar.html
│   └── pages/               # Pages
│       ├── home.html
│       ├── login.html
│       └── dashboard.html
├── .env.example             # Configuration example
├── go.mod                   # Go modules
└── README.md                # Documentation
```

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

## CLI Tool

DMMVC includes a powerful CLI tool for code generation!

### Quick Start with CLI

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

📖 **Full documentation**: [CLI.md](CLI.md)

## Roadmap

- [x] CLI tool for code generation ✅
- [ ] PostgreSQL support
- [ ] WebSocket support
- [ ] API documentation (Swagger)
- [ ] Caching (Redis)
- [ ] Task queues
- [ ] Email sending
- [ ] File upload helper
- [ ] Localization (i18n)
