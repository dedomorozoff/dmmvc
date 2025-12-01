# WebSocket in DMMVC

DMMVC supports WebSocket for real-time bidirectional communication between client and server.

## Features

- **Bidirectional Communication** - Real-time message exchange
- **Hub System** - Manage multiple connections
- **Auto-reconnection** - Restore connection on disconnect
- **Broadcast** - Send messages to all connected clients
- **Ping/Pong** - Automatic connection health checks

## Quick Start

### 1. Demo Page

Open in browser: `http://localhost:8080/websocket`

### 2. Usage in Code

```go
// Get Hub from context
hub := c.MustGet("hub").(*websocket.Hub)

// Broadcast message to all clients
hub.Broadcast([]byte("Hello, everyone!"))

// Get connected clients count
count := hub.GetClients()
```

## Architecture

### Hub

Hub manages all WebSocket connections:

```go
type Hub struct {
    clients    map[*Client]bool
    broadcast  chan []byte
    register   chan *Client
    unregister chan *Client
}
```

### Client

Client represents a single WebSocket connection:

```go
type Client struct {
    Hub  *Hub
    Conn *Conn
    Send chan []byte
    ID   string
}
```

## Usage Examples

### Chat Application

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

        // Notify all about new user
        hub.Broadcast([]byte(username + " joined the chat"))

        go client.WritePump()
        go client.ReadPump()
    }
}
```

### Real-time Notifications

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

### Live Data Updates

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

## Client Side

### Connection

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

### Sending Messages

```javascript
function sendMessage(message) {
    if (ws.readyState === WebSocket.OPEN) {
        ws.send(message);
    }
}
```

### Auto-reconnection

```javascript
function connect() {
    const ws = new WebSocket('ws://localhost:8080/ws');
    
    ws.onclose = function() {
        setTimeout(connect, 3000);
    };
}
```

## Configuration

### Timeout Settings

```go
const (
    writeWait      = 10 * time.Second  // Write timeout
    pongWait       = 60 * time.Second  // Pong timeout
    pingPeriod     = 54 * time.Second  // Ping period
    maxMessageSize = 512               // Max message size
)
```

### CORS Settings

```go
var upgrader = ws.Upgrader{
    CheckOrigin: func(r *http.Request) bool {
        // Check origin in production
        origin := r.Header.Get("Origin")
        return origin == "https://yourdomain.com"
    },
}
```

## Security

### Authentication

```go
func WebSocketHandler(hub *websocket.Hub) gin.HandlerFunc {
    return func(c *gin.Context) {
        // Check token
        token := c.Query("token")
        if !validateToken(token) {
            c.JSON(401, gin.H{"error": "Unauthorized"})
            return
        }
        
        // Upgrade connection
        conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
        // ...
    }
}
```

### Message Validation

```go
func (c *Client) ReadPump() {
    for {
        _, message, err := c.Conn.ReadMessage()
        if err != nil {
            break
        }
        
        // Validate message
        if !isValidMessage(message) {
            continue
        }
        
        c.Hub.broadcast <- message
    }
}
```

## Testing

### Connection Test

```bash
# Using websocat
websocat ws://localhost:8080/ws

# Using wscat
wscat -c ws://localhost:8080/ws
```

### Load Testing

```go
func TestWebSocketLoad(t *testing.T) {
    hub := websocket.NewHub()
    go hub.Run()
    
    // Create 1000 clients
    for i := 0; i < 1000; i++ {
        go connectClient(hub)
    }
    
    // Send messages
    for i := 0; i < 100; i++ {
        hub.Broadcast([]byte("test message"))
    }
}
```

## Performance

### Optimization

- Use buffered channels
- Limit message size
- Configure timeouts
- Use pools for object reuse

### Monitoring

```go
func (h *Hub) GetStats() map[string]interface{} {
    return map[string]interface{}{
        "clients":   h.GetClients(),
        "broadcast": len(h.broadcast),
        "register":  len(h.register),
    }
}
```

## Use Cases

### Game Server

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

### System Monitoring

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

### Issue: Connection drops

**Solution**: Check timeout settings and ping/pong

### Issue: Messages not delivered

**Solution**: Check Send channel buffer size

### Issue: High CPU usage

**Solution**: Optimize message sending frequency

## Additional Resources

- [Gorilla WebSocket Documentation](https://github.com/gorilla/websocket)
- [WebSocket RFC 6455](https://tools.ietf.org/html/rfc6455)
- [MDN WebSocket API](https://developer.mozilla.org/en-US/docs/Web/API/WebSocket)
