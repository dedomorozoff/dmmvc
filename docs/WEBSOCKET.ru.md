# WebSocket в DMMVC

DMMVC поддерживает WebSocket для двусторонней связи в реальном времени между клиентом и сервером.

## Возможности

- **Двусторонняя связь** - Обмен сообщениями в реальном времени
- **Hub система** - Управление множественными соединениями
- **Автоматическое переподключение** - Восстановление соединения при разрыве
- **Broadcast** - Отправка сообщений всем подключенным клиентам
- **Ping/Pong** - Автоматическая проверка соединения

## Быстрый старт

### 1. Демо страница

Откройте в браузере: `http://localhost:8080/websocket`

### 2. Использование в коде

```go
// Получить Hub из контекста
hub := c.MustGet("hub").(*websocket.Hub)

// Отправить сообщение всем клиентам
hub.Broadcast([]byte("Hello, everyone!"))

// Получить количество подключенных клиентов
count := hub.GetClients()
```

## Архитектура

### Hub

Hub управляет всеми WebSocket соединениями:

```go
type Hub struct {
    clients    map[*Client]bool
    broadcast  chan []byte
    register   chan *Client
    unregister chan *Client
}
```

### Client

Client представляет одно WebSocket соединение:

```go
type Client struct {
    Hub  *Hub
    Conn *Conn
    Send chan []byte
    ID   string
}
```

## Примеры использования

### Чат приложение

```go
func ChatHandler(hub *websocket.Hub) gin.HandlerFunc {
    return func(c *gin.Context) {
        username := c.Query("username")
        
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

        client.Hub.register <- client

        // Уведомить всех о новом пользователе
        hub.Broadcast([]byte(username + " joined the chat"))

        go client.WritePump()
        go client.ReadPump()
    }
}
```

### Уведомления в реальном времени

```go
func SendNotification(hub *websocket.Hub, message string) {
    notification := map[string]interface{}{
        "type":    "notification",
        "message": message,
        "time":    time.Now(),
    }
    
    data, _ := json.Marshal(notification)
    hub.Broadcast(data)
}
```

### Live обновления данных

```go
func DataUpdateHandler(hub *websocket.Hub) {
    ticker := time.NewTicker(5 * time.Second)
    defer ticker.Stop()

    for range ticker.C {
        data := fetchLatestData()
        jsonData, _ := json.Marshal(data)
        hub.Broadcast(jsonData)
    }
}
```

## Клиентская часть

### Подключение

```javascript
const ws = new WebSocket('ws://localhost:8080/ws');

ws.onopen = function() {
    console.log('Connected');
};

ws.onmessage = function(event) {
    console.log('Message:', event.data);
};

ws.onerror = function(error) {
    console.error('Error:', error);
};

ws.onclose = function() {
    console.log('Disconnected');
};
```

### Отправка сообщений

```javascript
function sendMessage(message) {
    if (ws.readyState === WebSocket.OPEN) {
        ws.send(message);
    }
}
```

### Автоматическое переподключение

```javascript
function connect() {
    const ws = new WebSocket('ws://localhost:8080/ws');
    
    ws.onclose = function() {
        setTimeout(connect, 3000);
    };
}
```

## Конфигурация

### Настройки таймаутов

```go
const (
    writeWait      = 10 * time.Second  // Таймаут записи
    pongWait       = 60 * time.Second  // Таймаут pong
    pingPeriod     = 54 * time.Second  // Период ping
    maxMessageSize = 512               // Макс размер сообщения
)
```

### CORS настройки

```go
var upgrader = ws.Upgrader{
    CheckOrigin: func(r *http.Request) bool {
        // В production проверять origin
        origin := r.Header.Get("Origin")
        return origin == "https://yourdomain.com"
    },
}
```

## Безопасность

### Аутентификация

```go
func WebSocketHandler(hub *websocket.Hub) gin.HandlerFunc {
    return func(c *gin.Context) {
        // Проверка токена
        token := c.Query("token")
        if !validateToken(token) {
            c.JSON(401, gin.H{"error": "Unauthorized"})
            return
        }
        
        // Upgrade соединения
        conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
        // ...
    }
}
```

### Валидация сообщений

```go
func (c *Client) ReadPump() {
    for {
        _, message, err := c.Conn.ReadMessage()
        if err != nil {
            break
        }
        
        // Валидация сообщения
        if !isValidMessage(message) {
            continue
        }
        
        c.Hub.broadcast <- message
    }
}
```

## Тестирование

### Тест подключения

```bash
# Используя websocat
websocat ws://localhost:8080/ws

# Используя wscat
wscat -c ws://localhost:8080/ws
```

### Нагрузочное тестирование

```go
func TestWebSocketLoad(t *testing.T) {
    hub := websocket.NewHub()
    go hub.Run()
    
    // Создать 1000 клиентов
    for i := 0; i < 1000; i++ {
        go connectClient(hub)
    }
    
    // Отправить сообщения
    for i := 0; i < 100; i++ {
        hub.Broadcast([]byte("test message"))
    }
}
```

## Производительность

### Оптимизация

- Используйте буферизованные каналы
- Ограничьте размер сообщений
- Настройте таймауты
- Используйте пулы для переиспользования объектов

### Мониторинг

```go
func (h *Hub) GetStats() map[string]interface{} {
    return map[string]interface{}{
        "clients":   h.GetClients(),
        "broadcast": len(h.broadcast),
        "register":  len(h.register),
    }
}
```

## Примеры использования

### Игровой сервер

```go
type GameHub struct {
    *websocket.Hub
    rooms map[string][]*websocket.Client
}

func (g *GameHub) JoinRoom(client *websocket.Client, roomID string) {
    g.rooms[roomID] = append(g.rooms[roomID], client)
}

func (g *GameHub) BroadcastToRoom(roomID string, message []byte) {
    for _, client := range g.rooms[roomID] {
        client.Send <- message
    }
}
```

### Мониторинг системы

```go
func SystemMonitor(hub *websocket.Hub) {
    ticker := time.NewTicker(1 * time.Second)
    defer ticker.Stop()

    for range ticker.C {
        stats := getSystemStats()
        data, _ := json.Marshal(stats)
        hub.Broadcast(data)
    }
}
```

## Troubleshooting

### Проблема: Соединение разрывается

**Решение**: Проверьте настройки таймаутов и ping/pong

### Проблема: Сообщения не доставляются

**Решение**: Проверьте размер буфера канала Send

### Проблема: Высокая нагрузка на CPU

**Решение**: Оптимизируйте частоту отправки сообщений

## Дополнительные ресурсы

- [Gorilla WebSocket документация](https://github.com/gorilla/websocket)
- [WebSocket RFC 6455](https://tools.ietf.org/html/rfc6455)
- [MDN WebSocket API](https://developer.mozilla.org/en-US/docs/Web/API/WebSocket)
