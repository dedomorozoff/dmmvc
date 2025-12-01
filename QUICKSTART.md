**English** | [Русский](QUICKSTART.ru.md)

# DMMVC Quick Start

## Installation and Run

### 1. Install Dependencies

```bash
go mod tidy
```

### 2. Run Server

```bash
go run cmd/server/main.go
```

### 3. Open in Browser

```
http://localhost:8080
```

### 4. Login

- **Username**: `admin`
- **Password**: `admin`

**Important**: Change the password after the first login!

## Project Structure

```
dmmvc/
├── cmd/server/main.go          # Entry point
├── internal/
│   ├── controllers/            # Controllers (HTTP handlers)
│   ├── database/               # Database operations
│   ├── logger/                 # Logging
│   ├── middleware/             # Middleware
│   ├── models/                 # Data models
│   └── routes/                 # Routes
├── static/                     # Static files
│   ├── css/style.css
│   └── js/app.js
├── templates/                  # HTML templates
│   ├── layouts/base.html
│   ├── partials/
│   └── pages/
├── .env                        # Configuration
└── go.mod                      # Dependencies
```

## Creating New Functionality

### 1. Create Model

```go
// internal/models/post.go
package models

import "gorm.io/gorm"

type Post struct {
    gorm.Model
    Title   string `json:"title"`
    Content string `json:"content"`
    UserID  uint   `json:"user_id"`
}
```

### 2. Add Migration

```go
// cmd/server/main.go
database.Migrate(&models.User{}, &models.Post{})
```

### 3. Create Controller

```go
// internal/controllers/post_controller.go
package controllers

import (
    "net/http"
    "github.com/gin-gonic/gin"
)

func PostList(c *gin.Context) {
    c.HTML(http.StatusOK, "pages/posts/list.html", gin.H{
        "title": "Posts",
    })
}
```

### 4. Add Route

```go
// internal/routes/routes.go
authorized.GET("/posts", controllers.PostList)
```

### 5. Create Template

```html
<!-- templates/pages/posts/list.html -->
{{define "pages/posts/list.html"}}
{{template "layouts/base.html" .}}
{{define "content"}}
{{template "partials/header.html" .}}

<main class="main">
    <div class="container">
        <h1>Posts List</h1>
    </div>
</main>

{{template "partials/footer.html" .}}
{{end}}
{{end}}
```

## Useful Commands

```bash
# Run in development mode
go run cmd/server/main.go

# Build binary
go build -o dmmvc cmd/server/main.go

# Run binary
./dmmvc

# Run tests
go test ./...
```

## Database Configuration

### SQLite (default)

```env
DB_TYPE=sqlite
DB_DSN=dmmvc.db
```

### MySQL

```env
DB_TYPE=mysql
DB_DSN=user:password@tcp(localhost:3306)/dmmvc?charset=utf8mb4&parseTime=True&loc=Local
```

## Done!

Now you have a fully working MVC framework ready for building any web application!
