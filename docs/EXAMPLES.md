**English** | [Русский](EXAMPLES.ru.md)

# DMMVC Usage Examples

## Example 1: Creating a Blog

### 1. Create Post Model

```go
// internal/models/post.go
package models

import "gorm.io/gorm"

type Post struct {
    gorm.Model
    Title       string `gorm:"not null" json:"title"`
    Content     string `gorm:"type:text" json:"content"`
    Published   bool   `gorm:"default:false" json:"published"`
    UserID      uint   `json:"user_id"`
    User        User   `gorm:"foreignKey:UserID"`
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
    "dmmvc/internal/database"
    "dmmvc/internal/models"
    "net/http"
    "github.com/gin-gonic/gin"
)

func PostList(c *gin.Context) {
    var posts []models.Post
    database.DB.Preload("User").Find(&posts)
    
    c.HTML(http.StatusOK, "pages/posts/list.html", gin.H{
        "title": "Blog",
        "posts": posts,
    })
}

func PostCreate(c *gin.Context) {
    c.HTML(http.StatusOK, "pages/posts/create.html", gin.H{
        "title": "Create Post",
    })
}

func PostStore(c *gin.Context) {
    userID, _ := c.Get("user_id")
    
    post := models.Post{
        Title:   c.PostForm("title"),
        Content: c.PostForm("content"),
        UserID:  userID.(uint),
    }
    
    database.DB.Create(&post)
    c.Redirect(http.StatusFound, "/posts")
}
```

### 4. Add Routes

```go
// internal/routes/routes.go
authorized.GET("/posts", controllers.PostList)
authorized.GET("/posts/create", controllers.PostCreate)
authorized.POST("/posts", controllers.PostStore)
```

### 5. Create Templates

```html
<!-- templates/pages/posts/list.html -->
{{define "pages/posts/list.html"}}
{{template "layouts/base.html" .}}
{{define "content"}}
{{template "partials/header.html" .}}

<main class="main">
    <div class="container">
        <div class="page-header">
            <h1 class="page-title">Blog</h1>
            <a href="/posts/create" class="btn btn-primary">Create Post</a>
        </div>

        <div class="posts-grid">
            {{range .posts}}
            <article class="post-card">
                <h2>{{.Title}}</h2>
                <p>{{.Content}}</p>
                <div class="post-meta">
                    <span>Author: {{.User.Username}}</span>
                    <span>{{.CreatedAt.Format "02.01.2006"}}</span>
                </div>
            </article>
            {{end}}
        </div>
    </div>
</main>

{{template "partials/footer.html" .}}
{{end}}
{{end}}
```

## Example 2: API Endpoint

### Create REST API

```go
// internal/controllers/api_controller.go
package controllers

import (
    "dmmvc/internal/database"
    "dmmvc/internal/models"
    "net/http"
    "github.com/gin-gonic/gin"
)

func APIPostList(c *gin.Context) {
    var posts []models.Post
    database.DB.Preload("User").Find(&posts)
    
    c.JSON(http.StatusOK, gin.H{
        "success": true,
        "data": posts,
    })
}

func APIPostCreate(c *gin.Context) {
    var post models.Post
    
    if err := c.ShouldBindJSON(&post); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "success": false,
            "error": err.Error(),
        })
        return
    }
    
    userID, _ := c.Get("user_id")
    post.UserID = userID.(uint)
    
    database.DB.Create(&post)
    
    c.JSON(http.StatusCreated, gin.H{
        "success": true,
        "data": post,
    })
}
```

### Add API Routes

```go
// internal/routes/routes.go
api := authorized.Group("/api")
{
    api.GET("/posts", controllers.APIPostList)
    api.POST("/posts", controllers.APIPostCreate)
}
```

## Example 3: Role Check Middleware

```go
// internal/middleware/role.go
package middleware

import (
    "net/http"
    "github.com/gin-gonic/gin"
)

func RequireRole(role string) gin.HandlerFunc {
    return func(c *gin.Context) {
        userRole, exists := c.Get("role")
        
        if !exists || userRole != role {
            c.HTML(http.StatusForbidden, "pages/403.html", gin.H{
                "title": "Access Denied",
            })
            c.Abort()
            return
        }
        
        c.Next()
    }
}
```

### Usage

```go
// internal/routes/routes.go
admin := authorized.Group("/admin")
admin.Use(middleware.RequireRole("admin"))
{
    admin.GET("/users", controllers.UserList)
}
```

## Example 4: Pagination

```go
// internal/controllers/post_controller.go
func PostList(c *gin.Context) {
    page := c.DefaultQuery("page", "1")
    pageSize := 10
    
    var posts []models.Post
    var total int64
    
    database.DB.Model(&models.Post{}).Count(&total)
    
    offset := (atoi(page) - 1) * pageSize
    database.DB.Preload("User").
        Offset(offset).
        Limit(pageSize).
        Find(&posts)
    
    c.HTML(http.StatusOK, "pages/posts/list.html", gin.H{
        "title": "Blog",
        "posts": posts,
        "page": page,
        "totalPages": (total + int64(pageSize) - 1) / int64(pageSize),
    })
}
```

## Example 5: File Upload

```go
// internal/controllers/upload_controller.go
func UploadFile(c *gin.Context) {
    file, err := c.FormFile("file")
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "error": "File not found",
        })
        return
    }
    
    filename := filepath.Base(file.Filename)
    if err := c.SaveUploadedFile(file, "./uploads/"+filename); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
            "error": "Error saving file",
        })
        return
    }
    
    c.JSON(http.StatusOK, gin.H{
        "success": true,
        "filename": filename,
    })
}
```

## Example 6: Data Validation

```go
// internal/models/post.go
type PostInput struct {
    Title   string `json:"title" binding:"required,min=3,max=100"`
    Content string `json:"content" binding:"required,min=10"`
}

// internal/controllers/post_controller.go
func PostStore(c *gin.Context) {
    var input PostInput
    
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "error": err.Error(),
        })
        return
    }
    
    // Create post...
}
```

## Example 7: Sending Email

```go
// internal/services/email.go
package services

import (
    "net/smtp"
)

func SendEmail(to, subject, body string) error {
    from := "noreply@example.com"
    password := "your-password"
    
    msg := "From: " + from + "\n" +
        "To: " + to + "\n" +
        "Subject: " + subject + "\n\n" +
        body
    
    err := smtp.SendMail("smtp.gmail.com:587",
        smtp.PlainAuth("", from, password, "smtp.gmail.com"),
        from, []string{to}, []byte(msg))
    
    return err
}
```

## Example 8: Caching

```go
// internal/middleware/cache.go
package middleware

import (
    "time"
    "github.com/gin-gonic/gin"
)

var cache = make(map[string]string)

func CacheMiddleware(duration time.Duration) gin.HandlerFunc {
    return func(c *gin.Context) {
        key := c.Request.URL.Path
        
        if cached, found := cache[key]; found {
            c.String(200, cached)
            c.Abort()
            return
        }
        
        c.Next()
        
        // Save to cache
        if c.Writer.Status() == 200 {
            // Cache implementation
        }
    }
}
```

---

These examples show how easy it is to extend the DMMVC framework to build various web applications!
