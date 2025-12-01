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

## Example 9: WebSocket Chat

### Create Chat Room

```go
// internal/controllers/chat.go
package controllers

import (
    "dmmvc/internal/websocket"
    "encoding/json"
    "net/http"
    "time"
    
    "github.com/gin-gonic/gin"
    ws "github.com/gorilla/websocket"
)

type ChatMessage struct {
    Username  string    `json:"username"`
    Message   string    `json:"message"`
    Timestamp time.Time `json:"timestamp"`
}

var upgrader = ws.Upgrader{
    CheckOrigin: func(r *http.Request) bool {
        return true
    },
}

func ChatHandler(hub *websocket.Hub) gin.HandlerFunc {
    return func(c *gin.Context) {
        username := c.Query("username")
        if username == "" {
            username = "Anonymous"
        }
        
        conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
        if err != nil {
            return
        }

        client := &websocket.Client{
            Hub:  hub,
            Conn: &websocket.Conn{Conn: conn},
            Send: make(chan []byte, 256),
            ID:   username,
        }

        // Notify about new user
        joinMsg := ChatMessage{
            Username:  "System",
            Message:   username + " joined the chat",
            Timestamp: time.Now(),
        }
        data, _ := json.Marshal(joinMsg)
        hub.Broadcast(data)

        client.Hub.register <- client

        go client.WritePump()
        go client.ReadPump()
    }
}

func ChatPage(c *gin.Context) {
    c.HTML(http.StatusOK, "pages/chat.html", gin.H{
        "title": "Chat Room",
    })
}
```

### Chat Template

```html
<!-- templates/pages/chat.html -->
{{define "pages/chat.html"}}
{{template "layouts/base.html" .}}
{{define "content"}}
<div class="container mt-5">
    <div class="card">
        <div class="card-header">
            <h3>Chat Room</h3>
            <span id="online" class="badge badge-success">0 online</span>
        </div>
        <div class="card-body">
            <div id="messages" style="height: 400px; overflow-y: auto;"></div>
            <div class="input-group mt-3">
                <input type="text" id="username" class="form-control" placeholder="Your name">
                <input type="text" id="message" class="form-control" placeholder="Message">
                <button id="send" class="btn btn-primary">Send</button>
            </div>
        </div>
    </div>
</div>

<script>
let ws;
const username = prompt("Enter your name:") || "Anonymous";

function connect() {
    ws = new WebSocket(`ws://${window.location.host}/chat?username=${username}`);
    
    ws.onmessage = function(event) {
        const msg = JSON.parse(event.data);
        addMessage(msg);
    };
}

function addMessage(msg) {
    const div = document.createElement('div');
    div.className = 'alert alert-info';
    div.innerHTML = `<strong>${msg.username}</strong>: ${msg.message}`;
    document.getElementById('messages').appendChild(div);
}

document.getElementById('send').onclick = function() {
    const message = document.getElementById('message').value;
    if (message && ws.readyState === WebSocket.OPEN) {
        const msg = {
            username: username,
            message: message,
            timestamp: new Date()
        };
        ws.send(JSON.stringify(msg));
        document.getElementById('message').value = '';
    }
};

connect();
</script>
{{end}}
{{end}}
```

### Add Route

```go
// internal/routes/routes.go
hub := websocket.NewHub()
go hub.Run()

r.GET("/chat", controllers.ChatPage)
r.GET("/ws/chat", controllers.ChatHandler(hub))
```


## Example 8: Internationalization (i18n)

### Using in Controllers

```go
package controllers

import (
    "dmmvc/internal/i18n"
    "net/http"
    "github.com/gin-gonic/gin"
)

func WelcomePage(c *gin.Context) {
    username := "John"
    
    c.HTML(http.StatusOK, "pages/welcome.html", gin.H{
        "title": i18n.T(c, "home.welcome"),
        "greeting": i18n.T(c, "dashboard.welcome", username),
    })
}

func APIGreeting(c *gin.Context) {
    name := c.Query("name")
    if name == "" {
        name = "Guest"
    }
    
    c.JSON(http.StatusOK, gin.H{
        "message": i18n.T(c, "dashboard.welcome", name),
        "locale": i18n.GetLocale(c),
    })
}
```

### Using in Templates

```html
{{define "pages/welcome.html"}}
{{template "partials/base_head.html" .}}
{{template "partials/header.html" .}}

<main class="main">
    <div class="container">
        <h1>{{t "home.welcome"}}</h1>
        <p>{{t "home.subtitle"}}</p>
        
        <div class="buttons">
            <a href="/dashboard" class="btn btn-primary">
                {{t "home.get_started"}}
            </a>
        </div>
        
        <p>{{t "common.loading"}}</p>
    </div>
</main>

{{template "partials/footer.html" .}}
{{template "partials/base_foot.html" .}}
{{end}}
```

### Switching Language via API

```javascript
// Switch to Russian
async function switchToRussian() {
    const response = await fetch('/api/locale', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify({ locale: 'ru' })
    });
    
    if (response.ok) {
        window.location.reload();
    }
}

// Get available languages
async function getAvailableLanguages() {
    const response = await fetch('/api/locales');
    const data = await response.json();
    console.log('Available:', data.data.locales);
    console.log('Current:', data.data.current);
}
```

### Adding a New Language

1. Create translation file `locales/es.json`:

```json
{
  "app.name": "DMMVC",
  "home.welcome": "Welcome to DMMVC",
  "nav.home": "Home",
  "nav.dashboard": "Dashboard"
}
```

2. Update `internal/i18n/i18n.go`:

```go
const (
    LocaleEN Locale = "en"
    LocaleRU Locale = "ru"
    LocaleES Locale = "es"
)

func (i *I18n) LoadTranslations(dir string) error {
    locales := []Locale{LocaleEN, LocaleRU, LocaleES}
    // ...
}
```

3. Update middleware to support new language:

```go
func parseLocale(lang string) Locale {
    if len(lang) >= 2 {
        switch lang[:2] {
        case "en":
            return LocaleEN
        case "ru":
            return LocaleRU
        case "es":
            return LocaleES
        }
    }
    return ""
}
```

See [i18n Documentation](I18N.md) for more details.
