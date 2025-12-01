[English](EXAMPLES.md) | **Русский**

# Примеры использования DMMVC

## Пример 1: Создание блога

### 1. Создаем модель Post

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

### 2. Добавляем миграцию

```go
// cmd/server/main.go
database.Migrate(&models.User{}, &models.Post{})
```

### 3. Создаем контроллер

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
        "title": "Блог",
        "posts": posts,
    })
}

func PostCreate(c *gin.Context) {
    c.HTML(http.StatusOK, "pages/posts/create.html", gin.H{
        "title": "Создать пост",
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

### 4. Добавляем маршруты

```go
// internal/routes/routes.go
authorized.GET("/posts", controllers.PostList)
authorized.GET("/posts/create", controllers.PostCreate)
authorized.POST("/posts", controllers.PostStore)
```

### 5. Создаем шаблоны

```html
<!-- templates/pages/posts/list.html -->
{{define "pages/posts/list.html"}}
{{template "layouts/base.html" .}}
{{define "content"}}
{{template "partials/header.html" .}}

<main class="main">
    <div class="container">
        <div class="page-header">
            <h1 class="page-title">Блог</h1>
            <a href="/posts/create" class="btn btn-primary">Создать пост</a>
        </div>

        <div class="posts-grid">
            {{range .posts}}
            <article class="post-card">
                <h2>{{.Title}}</h2>
                <p>{{.Content}}</p>
                <div class="post-meta">
                    <span>Автор: {{.User.Username}}</span>
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

## Пример 2: API Endpoint

### Создание REST API

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

### Добавляем API маршруты

```go
// internal/routes/routes.go
api := authorized.Group("/api")
{
    api.GET("/posts", controllers.APIPostList)
    api.POST("/posts", controllers.APIPostCreate)
}
```

## Пример 3: Middleware для проверки роли

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
                "title": "Доступ запрещен",
            })
            c.Abort()
            return
        }
        
        c.Next()
    }
}
```

### Использование

```go
// internal/routes/routes.go
admin := authorized.Group("/admin")
admin.Use(middleware.RequireRole("admin"))
{
    admin.GET("/users", controllers.UserList)
}
```

## Пример 4: Пагинация

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
        "title": "Блог",
        "posts": posts,
        "page": page,
        "totalPages": (total + int64(pageSize) - 1) / int64(pageSize),
    })
}
```

## Пример 5: Загрузка файлов

```go
// internal/controllers/upload_controller.go
func UploadFile(c *gin.Context) {
    file, err := c.FormFile("file")
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "error": "Файл не найден",
        })
        return
    }
    
    filename := filepath.Base(file.Filename)
    if err := c.SaveUploadedFile(file, "./uploads/"+filename); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
            "error": "Ошибка сохранения файла",
        })
        return
    }
    
    c.JSON(http.StatusOK, gin.H{
        "success": true,
        "filename": filename,
    })
}
```

## Пример 6: Валидация данных

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
    
    // Создание поста...
}
```

## Пример 7: Отправка Email

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

## Пример 8: Кеширование

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
        
        // Сохранение в кеш
        if c.Writer.Status() == 200 {
            // Реализация кеширования
        }
    }
}
```

---

Эти примеры показывают, как легко расширять DMMVC фреймворк для создания различных веб-приложений!

## Пример 9: WebSocket Чат

### Создание чат-комнаты

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
            username = "Аноним"
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

        // Уведомление о новом пользователе
        joinMsg := ChatMessage{
            Username:  "Система",
            Message:   username + " присоединился к чату",
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
        "title": "Чат",
    })
}
```

### Шаблон чата

```html
<!-- templates/pages/chat.html -->
{{define "pages/chat.html"}}
{{template "layouts/base.html" .}}
{{define "content"}}
<div class="container mt-5">
    <div class="card">
        <div class="card-header">
            <h3>Чат</h3>
            <span id="online" class="badge badge-success">0 онлайн</span>
        </div>
        <div class="card-body">
            <div id="messages" style="height: 400px; overflow-y: auto;"></div>
            <div class="input-group mt-3">
                <input type="text" id="username" class="form-control" placeholder="Ваше имя">
                <input type="text" id="message" class="form-control" placeholder="Сообщение">
                <button id="send" class="btn btn-primary">Отправить</button>
            </div>
        </div>
    </div>
</div>

<script>
let ws;
const username = prompt("Введите ваше имя:") || "Аноним";

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

### Добавление маршрута

```go
// internal/routes/routes.go
hub := websocket.NewHub()
go hub.Run()

r.GET("/chat", controllers.ChatPage)
r.GET("/ws/chat", controllers.ChatHandler(hub))
```


## Пример 8: Интернационализация (i18n)

### Использование в контроллерах

```go
package controllers

import (
    "dmmvc/internal/i18n"
    "net/http"
    "github.com/gin-gonic/gin"
)

func WelcomePage(c *gin.Context) {
    username := "Иван"
    
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

### Использование в шаблонах

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

### Переключение языка через API

```javascript
// Переключить на русский
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

// Получить доступные языки
async function getAvailableLanguages() {
    const response = await fetch('/api/locales');
    const data = await response.json();
    console.log('Available:', data.data.locales);
    console.log('Current:', data.data.current);
}
```

### Добавление нового языка

1. Создайте файл перевода `locales/es.json`:

```json
{
  "app.name": "DMMVC",
  "home.welcome": "Bienvenido a DMMVC",
  "nav.home": "Inicio",
  "nav.dashboard": "Panel"
}
```

2. Обновите `internal/i18n/i18n.go`:

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

3. Обновите middleware для поддержки нового языка:

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

Подробнее см. [Документацию по i18n](I18N.ru.md)
