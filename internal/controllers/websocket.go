package controllers

import (
	"net/http"

	"dmmvc/internal/i18n"
	"dmmvc/internal/websocket"

	"github.com/gin-gonic/gin"
	ws "github.com/gorilla/websocket"
)

var upgrader = ws.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true // В production настроить проверку origin
	},
}

// WebSocketHandler обрабатывает WebSocket соединения
func WebSocketHandler(hub *websocket.Hub) gin.HandlerFunc {
	return func(c *gin.Context) {
		conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to upgrade connection"})
			return
		}

		client := &websocket.Client{
			Hub:  hub,
			Conn: &websocket.Conn{Conn: conn},
			Send: make(chan []byte, 256),
			ID:   c.Query("id"),
		}

		hub.Register <- client

		go client.WritePump()
		go client.ReadPump()
	}
}

// WebSocketDemo страница демонстрации WebSocket
func WebSocketDemo(c *gin.Context) {
	locale := i18n.GetLocale(c)
	c.HTML(http.StatusOK, "pages/websocket.html", gin.H{
		"title":  i18nT(c, "websocket.title"),
		"locale": i18nLocale(c),
		"T":      i18n.TFunc(locale),
	})
}
